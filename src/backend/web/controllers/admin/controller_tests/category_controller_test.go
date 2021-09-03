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
	"gitee.com/itsos/golibs/v2/db/mysql"
	"gitee.com/itsos/studynotes"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestCategoryController_CategoryTopic(t *testing.T) {

	e := httptest.New(t, studynotes.NewApp().Application)

	t.Run("新增专题与绑定文章", func(t *testing.T) {
		r := e.POST("/admin/category/topic").
			WithHeader("Accept", "application/json").
			WithJSON(map[string]interface{}{
				"name": "创建专题",
				"aid":  "1",
			}).
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("创建重复的专题", func(t *testing.T) {
		r := e.POST("/admin/category/topic").
			WithHeader("Accept", "application/json").
			WithJSON(map[string]interface{}{
				"name": "创建专题",
				"aid":  "1",
			}).
			Expect().Status(iris.StatusBadRequest)
		t.Log(r.JSON().Object().ValueEqual("code", 4002001))
	})

	t.Run("修改专题", func(t *testing.T) {
		r := e.PUT("/admin/category/topic").
			WithHeader("Accept", "application/json").
			WithForm(map[string]interface{}{
				"name": "修改专题",
				"id":   "1",
			}).
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("查询专题列表", func(t *testing.T) {
		r := e.GET("/admin/category/topics").
			WithHeader("Accept", "application/json").
			WithQuery("id", "1").
			Expect().Status(iris.StatusOK)
		r.JSON().Array().Element(0).Object().ValueEqual("name", "修改专题")
	})

	t.Run("获取绑定次数", func(t *testing.T) {
		r := e.GET("/admin/category/bindartcount").
			WithQuery("id", "1").
			Expect().Status(iris.StatusOK)
		r.JSON().Number().Equal(1)
	})

	t.Run("解除专题绑定关系", func(t *testing.T) {
		r := e.DELETE("/admin/category/relations").
			WithQuery("id", "1").
			WithQuery("aid", "1").
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("删除专题", func(t *testing.T) {
		r := e.DELETE("/admin/category/topic").
			WithQuery("id", "1").
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})

	t.Run("删除不存在专题", func(t *testing.T) {
		r := e.DELETE("/admin/category/topic").
			WithHeader("Accept", "application/json").
			WithQuery("id", "1").
			Expect().Status(iris.StatusBadRequest)
		r.JSON().Object().ValueEqual("code", 4002002)
	})

	mysql.NewMysql().Exec("truncate table category")
	mysql.NewMysql().Exec("truncate table category_rel")
}
