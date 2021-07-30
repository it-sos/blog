/*
   Copyright (c) [2021] itsos
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
               http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package studynotes

import (
	"gitee.com/itsos/studynotes/config"
	"gitee.com/itsos/studynotes/web/bootstrap"
	"gitee.com/itsos/studynotes/web/middleware/identity"
	"gitee.com/itsos/studynotes/web/routes"
	"github.com/kataras/iris/v12"
	"net"
	"os"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("studynotes", "peng.yu@qjfu.cn")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func Listen() {
	newApp().Listen(":"+config.C.GetPort(), iris.WithOptimizations)
}

// ListenSock socket 方式
// socat -d -d TCP-LISTEN:8080,fork,bind=127.0.0.1 UNIX:/tmp/studynote.sock
// curl http://localhost:8080
func ListenSock() {
	app := newApp()
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
