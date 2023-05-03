/*
   Copyright (c) [2021] itsos
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
               http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package config

import (
	"reflect"

	"github.com/it-sos/golibs/config"
	"github.com/it-sos/golibs/utils/reflects"
	"github.com/spf13/viper"
)

type ConfigurationReadOnly interface {
	config.ConfigurationReadOnly
	GetProxy() string
	GetChatAPIKey() string
	GetChatApiUrl() string
	GetChatBotDesc() string
	GetChatModel() string
	GetChatMaxTokens() int
	GetChatTemperature() float32
	GetChatTopp() float32
	GetChatFrequencyPenalty() float32
	GetChatPresencePenalty() float32
}

type Configuration struct {
	*config.Configuration

	HttpProxy            string `yaml:"http.proxy"`
	ChatApiKey           string `yaml:"chat.api_key"`
	ChatApiUrl           string `yaml:"chat.api_url"`
	ChatBotDesc          string `yaml:"chat.bot_desc"`
	ChatModel            string `yaml:"chat.model"`
	ChatMaxTokens        string `yaml:"chat.max_tokens"`
	ChatTemperature      string `yaml:"chat.temperature"`
	ChatTopp             string `yaml:"chat.top_p"`
	ChatFrequencyPenalty string `yaml:"chat.frequency_penalty"`
	ChatPresencePenalty  string `yaml:"chat.presence_penalty"`
}

func (c *Configuration) GetChatApiUrl() string {
	return viper.GetString(c.ChatApiUrl)
}

func (c *Configuration) GetChatBotDesc() string {
	return viper.GetString(c.ChatBotDesc)
}

func (c *Configuration) GetChatModel() string {
	return viper.GetString(c.ChatModel)
}

func (c *Configuration) GetChatMaxTokens() int {
	return viper.GetInt(c.ChatMaxTokens)
}

func (c *Configuration) GetChatTemperature() float32 {
	return float32(viper.GetFloat64(c.ChatTemperature))
}

func (c *Configuration) GetChatTopp() float32 {
	return float32(viper.GetFloat64(c.ChatTopp))
}

func (c *Configuration) GetChatFrequencyPenalty() float32 {
	return float32(viper.GetFloat64(c.ChatFrequencyPenalty))
}

func (c *Configuration) GetChatPresencePenalty() float32 {
	return float32(viper.GetFloat64(c.ChatPresencePenalty))
}

// GetProxy 获取代理地址
func (c *Configuration) GetProxy() string {
	return viper.GetString(c.HttpProxy)
}

// GetChatAPIKey 获取 chat api key
func (c *Configuration) GetChatAPIKey() string {
	return viper.GetString(c.ChatApiKey)
}

func covertConfiguration() *Configuration {
	c := &Configuration{
		Configuration: config.Config,
	}
	t := reflect.ValueOf(c).Elem()
	reflects.TagToValueFlip(t, reflects.YAML)
	return c
}

// C 配置
var C ConfigurationReadOnly = covertConfiguration()
