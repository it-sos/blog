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
	"gitee.com/itsos/blog/datamodels"
	"testing"
	"time"
)

func TestArticleRepository_Navigation(t *testing.T) {
	state := []uint8{IsStatePublic, IsStatePrivate}
	times, _ := time.Parse("2006-01-02 15:04:05", "2017-09-22 16:39:57")
	t.Log(RArticle.Navigation(state, times))
}

func TestArticleRepository_Select(t *testing.T) {
	state := []uint8{IsStatePublic, IsStatePrivate}
	t.Log(RArticle.Select(state, &datamodels.Article{}))
}

func TestArticleRepository_InsertTrans(t *testing.T) {
	t.Log(RArticle.InsertTrans(&datamodels.Article{Title: "hello"}, "hello world."))
}

func TestArticleRepository_UpdateTrans(t *testing.T) {
	RArticle.UpdateTrans(1, &datamodels.Article{Title: "hello"}, "hello man.")
}

func Test_articleRepository_TitleExists(t *testing.T) {
	t.Log(RArticle.TitleExists("此行为标题，固定样式为H2"))
}
