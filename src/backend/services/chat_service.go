/*
   Copyright (c) [2021] IT.SOS
   study-notes is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package services

import (
	"time"

	"context"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/869413421/chatgpt-web/pkg/types"
	"github.com/it-sos/blog/config"
	"github.com/kataras/golog"

	gogpt "github.com/sashabaranov/go-openai"
	"golang.org/x/net/proxy"
)

var chatModels = []string{gogpt.GPT432K0314, gogpt.GPT4, gogpt.GPT40314, gogpt.GPT432K, gogpt.GPT3Dot5Turbo, gogpt.GPT3Dot5Turbo0301}

// ChatService chat ai 通讯
type ChatService interface {
	// Completion 获取ai回复
	Completion(askMessage []string) (string, error)
}

type chatService struct{}

func (c *chatService) Completion(askMessage []string) (replyMessage string, err error) {
	var request gogpt.ChatCompletionRequest

	gptConfig := gogpt.DefaultConfig(config.C.GetChatAPIKey())
	var proxy = config.C.GetProxy()
	if proxy != "" {
		transport := &http.Transport{}

		if strings.HasPrefix(config.C.GetProxy(), "socks5h://") {
			// 创建一个 DialContext 对象，并设置代理服务器
			dialContext, err := newDialContext(proxy[10:])
			if err != nil {
				panic(err)
			}
			transport.DialContext = dialContext
		} else {
			// 创建一个 HTTP Transport 对象，并设置代理服务器
			proxyUrl, err := url.Parse(proxy)
			if err != nil {
				panic(err)
			}
			transport.Proxy = http.ProxyURL(proxyUrl)
		}
		// 创建一个 HTTP 客户端，并将 Transport 对象设置为其 Transport 字段
		gptConfig.HTTPClient = &http.Client{
			Transport: transport,
		}

	}

	// 自定义gptConfig.BaseURL
	if config.C.GetChatApiUrl() != "" {
		gptConfig.BaseURL = config.C.GetChatApiUrl()
	}
	var ctx = context.Background()

	client := gogpt.NewClientWithConfig(gptConfig)
	newMessage := append([]gogpt.ChatCompletionMessage{
		{Role: "system", Content: config.C.GetChatBotDesc()},
		{Role: "user", Content: askMessage[0]},
	}, request.Messages...)
	request.Messages = newMessage
	golog.Info(request.Messages)

	var model = config.C.GetChatModel()

	if types.Contains(chatModels, config.C.GetChatModel()) {
		request.Model = model
		var resp gogpt.ChatCompletionResponse
		resp, err = client.CreateChatCompletion(ctx, request)
		if err != nil {
			return
		}
		replyMessage = resp.Choices[0].Message.Content
		golog.Infof("in model choices: %v", resp.Choices)
	} else {
		prompt := ""
		for _, msg := range askMessage {
			prompt += msg + "/n"
		}
		prompt = strings.Trim(prompt, "/n")

		golog.Infof("request prompt is %s", prompt)
		req := gogpt.ChatCompletionRequest{
			Model:            config.C.GetChatModel(),
			MaxTokens:        config.C.GetChatMaxTokens(),
			TopP:             config.C.GetChatTopp(),
			FrequencyPenalty: config.C.GetChatFrequencyPenalty(),
			PresencePenalty:  config.C.GetChatPresencePenalty(),
			Messages:         []gogpt.ChatCompletionMessage{{Role: "user", Content: prompt}},
		}

		var resp gogpt.ChatCompletionResponse
		resp, err = client.CreateChatCompletion(ctx, req)
		if err != nil {
			return
		}
		replyMessage = strings.ReplaceAll(resp.Choices[0].Message.Content, "\n", "<br/>")
	}
	return
}

type dialContextFunc func(ctx context.Context, network, address string) (net.Conn, error)

func newDialContext(socks5 string) (dialContextFunc, error) {
	baseDialer := &net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 60 * time.Second,
	}

	if socks5 != "" {
		// split socks5 proxy string [username:password@]host:port
		var auth *proxy.Auth = nil

		if strings.Contains(socks5, "@") {
			proxyInfo := strings.SplitN(socks5, "@", 2)
			proxyUser := strings.Split(proxyInfo[0], ":")
			if len(proxyUser) == 2 {
				auth = &proxy.Auth{
					User:     proxyUser[0],
					Password: proxyUser[1],
				}
			}
			socks5 = proxyInfo[1]
		}

		dialSocksProxy, err := proxy.SOCKS5("tcp", socks5, auth, baseDialer)
		if err != nil {
			return nil, err
		}

		contextDialer, ok := dialSocksProxy.(proxy.ContextDialer)
		if !ok {
			return nil, err
		}

		return contextDialer.DialContext, nil
	} else {
		return baseDialer.DialContext, nil
	}
}

// SChatService chat service
var SChatService ChatService = &chatService{}
