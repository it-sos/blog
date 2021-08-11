/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package services

import (
	"testing"
)

func Test_GetRange(t *testing.T) {
	t.Log(SArticle.GetRank(true))
}

func Test_GetListPage(t *testing.T) {
	t.Log(SArticle.GetListPage(true, 1, 10, "linux curl"))
}

func TestArticleService_GetContent(t *testing.T) {
	s := SArticle.GetContent(false, "@Valid 参数验证之自定义注解实现复杂验证逻辑")
	t.Log(s.Article)
	t.Log(s.Navigation)
	t.Log(s.ArticleContent.Data)
	t.Log(s.Tags)
	t.Log(s.Topics)
}

func BenchmarkArticleService_GetContent(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		SArticle.GetContent(false, "@Valid 参数验证之自定义注解实现复杂验证逻辑")
	}
}
