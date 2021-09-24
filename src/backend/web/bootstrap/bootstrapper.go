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
	"gitee.com/itsos/golibs/v2/errors"
	"gitee.com/itsos/golibs/v2/global/consts"
	"gitee.com/itsos/golibs/v2/global/variable"
	"gitee.com/itsos/studynotes/config"
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"net"
	"strconv"
	"strings"
	"time"
)

var c = config.C

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
	Sessions     *sessions.Sessions
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

func (b *Bootstrapper) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + strings.ToUpper(b.AppName),
		Expires:  expires,
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
	})
	b.Sessions.UseDatabase(redis.New(redis.Config{
		Addr:     c.GetRedis().GetHost() + ":" + strconv.Itoa(c.GetRedis().GetPort()),
		Database: strconv.Itoa(c.GetRedis().GetDb()),
	}))
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
			"code":    res.GetErrCode(),
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

func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	// 设置golang views目录
	b.SetupViews(variable.BasePath + "/web/views")

	// 设置session
	hashKey := securecookie.GenerateRandomKey(64)
	blockKey := securecookie.GenerateRandomKey(32)
	b.SetupSessions(10*time.Minute, hashKey, blockKey)

	// 设置错误捕获
	b.SetupErrorHandlers()

	// static files
	//b.Favicon(StaticAssets + Favicon)
	router.DefaultDirOptions.IndexName = "/index.html"
	b.HandleDir("public", iris.Dir(StaticAssets), router.DefaultDirOptions)
	b.Use(recover.New())
	b.Use(logger.New())
	b.Use(b.Sessions.Handler())

	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}

func (b *Bootstrapper) ListenSock(l net.Listener, cfgs ...iris.Configurator) {
	b.Run(iris.Listener(l), cfgs...)
}
