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
	// ArticleVO 首页文章列表vo
	ArticleVO struct {
		datamodels.Article
		Duration string `json:"duration"` // 持续时间
		Topics   []string
		Tags     []string
	}

	// ArticleAccessTimesVO  文章及访问次数
	ArticleAccessTimesVO struct {
		Title       string `json:"title"`        // 文章标题
		AccessTimes int    `json:"access_times"` // 访问次数
	}

	// TopicVO 专题下文章列表及访问次数
	TopicVO struct {
		Title   string                 // 专题名
		Article []ArticleAccessTimesVO // 文章列表
	}

	// ArticleContentVO 文章内容及专题列表
	ArticleContentVO struct {
		Article ArticleVO
		Topic   []TopicVO
	}
)
