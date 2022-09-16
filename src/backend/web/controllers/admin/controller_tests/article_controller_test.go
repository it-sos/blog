/*
   Copyright (c) [2021] IT.SOS
   study-notes is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package controller_tests

import (
	"github.com/it-sos/blog"
	"github.com/it-sos/golibs/v2/db/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

// 增删改查 - 单元测试
func TestArticleController_Article(t *testing.T) {
	e := httptest.New(t, blog.NewApp().Application)
	t.Run("创建文章", func(t *testing.T) {
		r := e.POST("/admin/article").
			WithHeader("Accept", "application/json").
			WithJSON(map[string]interface{}{
				"id":         0,
				"title":      "创建文章",
				"intro":      "文章简介",
				"content":    "文章内容",
				"is_state":   1,
				"is_encrypt": 1,
			}).
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("修改文章", func(t *testing.T) {
		r := e.POST("/admin/article").
			WithHeader("Accept", "application/json").
			WithJSON(map[string]interface{}{
				"id":         1,
				"title":      "修改文章",
				"intro":      "文章简介",
				"content":    "文章内容",
				"is_state":   1,
				"is_encrypt": 1,
			}).
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("获取文章", func(t *testing.T) {
		r := e.GET("/admin/article").WithQuery("id", 1).Expect().Status(iris.StatusOK)
		t.Log(r.JSON().Object().ValueEqual("title", "修改文章"))
	})

	t.Run("删除文章", func(t *testing.T) {
		r := e.DELETE("/admin/article").WithQuery("id", 1).Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("获取文章", func(t *testing.T) {
		r := e.GET("/admin/article").
			WithQuery("id", 1).
			WithHeader("accept", "application/json").
			Expect().Status(iris.StatusBadRequest)
		r.JSON().Object().ValueEqual("code", 4001002)
	})

	mysql.NewMysql().Exec("truncate table article")
	mysql.NewMysql().Exec("truncate table article_content")
}
