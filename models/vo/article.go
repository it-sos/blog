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
	ArticlesVO struct {
		Time int `json:"time"` // 时间量，1天前
		*datamodels.Article
	}

	ArticleVO struct {
		AccesssTimes int `json:"accesss_times"` // 访问次数
		*datamodels.Article
	}

	TopicVO struct {
		Title   string // 专题名
		Article []ArticleVO
	}

	ArticleContentVO struct {
		Topic []TopicVO
	}
)
