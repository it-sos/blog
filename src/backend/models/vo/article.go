/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package vo

import "gitee.com/itsos/studynotes/datamodels"

type (
	// ArticleEditVO 后台文章详情
	ArticleEditVO struct {
		Id        uint     `json:"id"`
		Title     string   `json:"title"`      // 文章标题
		Intro     string   `json:"intro"`      // 文章简介
		Content   string   `json:"content"`    // 文章内容
		IsState   uint8    `json:"is_state"`   // 状态1私有2公开
		IsEncrypt uint8    `json:"is_encrypt"` // 是否加密1明文2密文
		Tags      []string `json:"tags"`       // 标签ID列表
		Topics    []string `json:"topics"`     // 专题ID列表
	}

	// ArticleParamsVO 后台文章详情
	ArticleParamsVO struct {
		Id        uint   `json:"id"`
		Title     string `json:"title"`      // 文章标题
		Intro     string `json:"intro"`      // 文章简介
		Content   string `json:"content"`    // 文章内容
		IsState   uint8  `json:"is_state"`   // 状态1私有2公开
		IsEncrypt uint8  `json:"is_encrypt"` // 是否加密1明文2密文
	}

	// ArticleVO 首页文章列表vo
	ArticleVO struct {
		Article  datamodels.Article `json:"article"`
		Duration string             `json:"duration"` // 持续时间
		Topics   []string           `json:"topics"`
		Tags     []string           `json:"tags"`
	}

	// ArticleAccessTimesVO  文章及访问次数
	ArticleAccessTimesVO struct {
		Title       string `json:"title"`        // 文章标题
		AccessTimes int    `json:"access_times"` // 访问次数
	}

	// TopicVO 专题下文章列表及访问次数
	TopicVO struct {
		Title   string                 `json:"title"`   // 专题名
		Article []ArticleAccessTimesVO `json:"article"` // 文章列表
	}

	// TagVO 标签下文章列表及访问次数
	TagVO struct {
		Title   string                 `json:"title"`   // 专题名
		Article []ArticleAccessTimesVO `json:"article"` // 文章列表
	}

	// NavigationVO 导航
	NavigationVO struct {
		PrevTitle string `json:"prev_title"` // 上一文章title
		NextTitle string `json:"next_title"` // 下一文章title
	}

	// ArticleContentVO 文章内容及专题列表
	ArticleContentVO struct {
		Navigation     NavigationVO              `json:"navigations"`     // 导航
		Article        ArticleVO                 `json:"article"`         // 文章信息
		ArticleContent datamodels.ArticleContent `json:"article_content"` // 文章内容
		Topics         []TopicVO                 `json:"topics"`          // 专题文章列表
		Tags           []TagVO                   `json:"tags"`            // 标签文章列表
	}
)
