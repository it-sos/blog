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
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func TestPostCompletion(t *testing.T) {

	e := httptest.New(t, blog.NewApp().Application)

	t.Run("发消息", func(t *testing.T) {
		r := e.POST("/api/admin/chat/completion").
			WithHeader("Accept", "application/json").
			WithForm(map[string]interface{}{
				"keyword": "如何成为架构师",
			}).
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})
}
