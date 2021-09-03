/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package routes

import (
	"gitee.com/itsos/studynotes/web/bootstrap"
	"gitee.com/itsos/studynotes/web/controllers"
	"gitee.com/itsos/studynotes/web/controllers/admin"
	"gitee.com/itsos/studynotes/web/middleware/auth"
	"github.com/kataras/iris/v12/mvc"
	"time"
)

func Configure(b *bootstrap.Bootstrapper) {
	mvc.Configure(
		b.Party("/"),
		func(app *mvc.Application) {
			app.Handle(new(controllers.IndexController))
		},
	).Register(b.Sessions.Start, time.Now())

	mvc.Configure(
		b.Party("/admin", auth.Secret),
		func(app *mvc.Application) {
			app.Handle(new(admin.ArticleController)).
				Handle(new(admin.CategoryController))
		},
	).Register(b.Sessions.Start, time.Now())
}
