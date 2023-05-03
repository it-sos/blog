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
	"testing"

	"github.com/it-sos/blog"
	"github.com/it-sos/golibs/db/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func TestCategoryController_CategoryTopic(t *testing.T) {

	e := httptest.New(t, blog.NewApp().Application)

	t.Run("新增专题", func(t *testing.T) {
		r := e.POST("/api/admin/category/topic").
			WithHeader("Accept", "application/json").
			WithForm(map[string]interface{}{
				"name": "创建专题",
			}).
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("创建重复的专题", func(t *testing.T) {
		r := e.POST("/api/admin/category/topic").
			WithHeader("Accept", "application/json").
			WithForm(map[string]interface{}{
				"name": "创建专题",
			}).
			Expect().Status(iris.StatusBadRequest)
		t.Log(r.JSON().Object().ValueEqual("code", 4002001))
	})

	t.Run("修改专题", func(t *testing.T) {
		r := e.PUT("/api/admin/category/topic").
			WithHeader("Accept", "application/json").
			WithForm(map[string]interface{}{
				"name": "修改专题",
				"id":   "1",
			}).
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("查询专题列表", func(t *testing.T) {
		r := e.GET("/api/admin/category/topics").
			WithHeader("Accept", "application/json").
			WithQuery("id", "1").
			Expect().Status(iris.StatusOK)
		r.JSON().Array().Element(0).Object().ValueEqual("name", "修改专题")
	})

	t.Run("新增专题文章绑定", func(t *testing.T) {
		r := e.POST("/api/admin/category/bindtopic").
			WithQuery("id", "1").
			WithQuery("aid", "1").
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("获取绑定次数", func(t *testing.T) {
		r := e.GET("/api/admin/category/bindartcount").
			WithQuery("id", "1").
			Expect().Status(iris.StatusOK)
		r.JSON().Number().Equal(1)
	})

	t.Run("解除专题绑定关系", func(t *testing.T) {
		r := e.DELETE("/api/admin/category/relations").
			WithQuery("id", "1").
			WithQuery("aid", "1").
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("删除专题", func(t *testing.T) {
		r := e.DELETE("/api/admin/category/topic").
			WithQuery("id", "1").
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("删除不存在专题", func(t *testing.T) {
		r := e.DELETE("/api/admin/category/topic").
			WithHeader("Accept", "application/json").
			WithQuery("id", "1").
			Expect().Status(iris.StatusBadRequest)
		r.JSON().Object().ValueEqual("code", 4002002)
	})

	t.Run("清除表数据", func(t *testing.T) {
		mysql.NewMysql().Exec("truncate table category")
		mysql.NewMysql().Exec("truncate table category_rel")
	})
}

func TestCategoryController_CategoryTag(t *testing.T) {

	e := httptest.New(t, blog.NewApp().Application)

	t.Run("新增标签", func(t *testing.T) {
		r := e.POST("/api/admin/category/tag").
			WithHeader("Accept", "application/json").
			WithForm(map[string]interface{}{
				"name": "创建标签",
			}).
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("创建重复的标签", func(t *testing.T) {
		r := e.POST("/api/admin/category/tag").
			WithHeader("Accept", "application/json").
			WithForm(map[string]interface{}{
				"name": "创建标签",
			}).
			Expect().Status(iris.StatusBadRequest)
		t.Log(r.JSON().Object().ValueEqual("code", 4002001))
	})

	t.Run("修改标签", func(t *testing.T) {
		r := e.PUT("/api/admin/category/tag").
			WithHeader("Accept", "application/json").
			WithForm(map[string]interface{}{
				"name": "修改标签",
				"id":   "1",
			}).
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("查询标签列表", func(t *testing.T) {
		r := e.GET("/api/admin/category/tags").
			WithHeader("Accept", "application/json").
			WithQuery("id", "1").
			Expect().Status(iris.StatusOK)
		r.JSON().Array().Element(0).Object().ValueEqual("name", "修改标签")
	})

	t.Run("新增标签文章绑定", func(t *testing.T) {
		r := e.POST("/api/admin/category/bindtag").
			WithQuery("id", "1").
			WithQuery("aid", "1").
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("获取绑定次数", func(t *testing.T) {
		r := e.GET("/api/admin/category/bindartcount").
			WithQuery("id", "1").
			Expect().Status(iris.StatusOK)
		r.JSON().Number().Equal(1)
	})

	t.Run("解除标签绑定关系", func(t *testing.T) {
		r := e.DELETE("/api/admin/category/relations").
			WithQuery("id", "1").
			WithQuery("aid", "1").
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("删除标签", func(t *testing.T) {
		r := e.DELETE("/api/admin/category/tag").
			WithQuery("id", "1").
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("删除不存在标签", func(t *testing.T) {
		r := e.DELETE("/api/admin/category/tag").
			WithHeader("Accept", "application/json").
			WithQuery("id", "1").
			Expect().Status(iris.StatusBadRequest)
		r.JSON().Object().ValueEqual("code", 4003002)
	})

	t.Run("清除表数据", func(t *testing.T) {
		mysql.NewMysql().Exec("truncate table category")
		mysql.NewMysql().Exec("truncate table category_rel")
	})
}
