/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package repositories

import (
	_ "gitee.com/itsos/golibs/tests/testsdb"
	"gitee.com/itsos/studynotes/datamodels"
	"testing"
)

var article = &datamodels.Article{}

func connect() ArticleRepository {
	return NewArticleRepository()
}

func Test_articleRepository_InsertOrUpdate(t *testing.T) {
	t.Log(connect().Insert(article))
}

func Test_articleRepository_SelectMany(t *testing.T) {
	t.Log(connect().SelectMany([]int8{IS_STATE_PUBLIC, IS_STATE_PRIVATE}, 0, 3))
}

func Test_articleRepository_Select(t *testing.T) {
	article.Id = 1
	t.Log(connect().Select(article))
}
