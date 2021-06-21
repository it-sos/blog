/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package bootstrap

import (
	"gitee.com/itsos/golibs/config"
	"gitee.com/itsos/golibs/db"
	"gitee.com/itsos/golibs/errors"
	"gitee.com/itsos/golibs/global/consts"
	"gitee.com/itsos/golibs/global/variable"
	config2 "gitee.com/itsos/studynotes/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"time"
)

var c = config2.C

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

// New returns a new Bootstrapper.
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// SetupViews loads the templates.
func (b *Bootstrapper) SetupViews(viewsDir string) {
	b.RegisterView(iris.HTML(viewsDir, ".html").
		Layout("shared/layout.html").
		Reload(c.GetActive() != consts.EnvProduct))
}

// SetupErrorHandlers `(context.StatusCodeNotSuccessful`,  which defaults to >=400 (but you can change it).
func (b *Bootstrapper) SetupErrorHandlers() {
	b.APIBuilder.ConfigureContainer().Container.GetErrorHandler = func(*context.Context) hero.ErrorHandler {
		return hero.ErrorHandlerFunc(func(ctx *context.Context, err error) {
			if err != hero.ErrStopExecution {
				if status := ctx.GetStatusCode(); status == 0 || !context.StatusCodeNotSuccessful(status) {
					ctx.StatusCode(hero.DefaultErrStatusCode)
				}
				if isOutJson(ctx) {
					ctx.ContentType(context.ContentJSONHeaderValue)
				}
				_, _ = ctx.WriteString(err.Error())
			}

			ctx.StopExecution()
		})
	}

	b.OnAnyErrorCode(func(ctx iris.Context) {
		res := errors.Errors{}
		res.SetErrCode(ctx.GetStatusCode())
		res.SetMessage(iris.StatusText(ctx.GetStatusCode()))

		if isOutJson(ctx) {
			ctx.JSON(res)
			return
		}

		err := iris.Map{
			"errCode": res.GetErrCode(),
			"message": res.GetMessage(),
		}
		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

func isOutJson(ctx iris.Context) bool {
	return ctx.IsAjax() ||
		ctx.URLParamExists("json") ||
		ctx.GetHeader("Accept") == context.ContentJSONHeaderValue
}

var (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = variable.BasePath + "/web/public/"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	//Favicon = "favicon.ico"
)

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// 初始化配置和存储连接等
func (b *Bootstrapper) initialization() {
	config.C.Init()
	db.Init()
}

func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.initialization()

	b.SetupViews(variable.BasePath + "/web/views")
	b.SetupErrorHandlers()

	b.Use(func(ctx *context.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Next()
	})

	// static files
	//b.Favicon(StaticAssets + Favicon)
	router.DefaultDirOptions.IndexName = "/default.html"
	b.HandleDir("public", iris.Dir(StaticAssets), router.DefaultDirOptions)

	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
