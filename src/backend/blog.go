/*
   Copyright (c) [2021] itsos
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
               http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package blog

import (
	"net"
	"os"

	"gitee.com/itsos/blog/config"
	"gitee.com/itsos/blog/web/routes"
	"gitee.com/itsos/golibs/v2/framework/iris/bootstrap"
	"gitee.com/itsos/golibs/v2/framework/iris/middleware/auth"
	"gitee.com/itsos/golibs/v2/framework/iris/middleware/identity"
	"github.com/kataras/iris/v12"
)

func NewApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("blog", "peng.yu@qjfu.cn")
	app.Bootstrap()
	app.Configure(identity.Identity, auth.CheckSign, routes.Routes)
	return app
}

func Listen() {
	NewApp().Listen(":"+config.C.GetPort(), iris.WithOptimizations)
}

//ListenSock socket 方式
//socat -d -d TCP-LISTEN:8080,fork,bind=127.0.0.1 UNIX:/tmp/blog.sock
//curl http://localhost:8080
func ListenSock() {
	app := NewApp()
	socketFile := config.C.GetSock()
	if errOs := os.Remove(socketFile); errOs != nil && !os.IsNotExist(errOs) {
		app.Logger().Fatal(errOs)
	}
	l, err := net.Listen("unix", socketFile)
	if err != nil {
		app.Logger().Fatal(err)
	}
	if err = os.Chmod(socketFile, 0666|os.ModeSocket); err != nil {
		app.Logger().Fatal(err)
	}
	app.ListenSock(l, iris.WithOptimizations)
}
