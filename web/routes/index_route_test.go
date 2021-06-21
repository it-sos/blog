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
	"gitee.com/itsos/study-notes/tests"
	"github.com/kataras/iris/v12/httptest"
	"os"
	"testing"
)

func TestGetQrcode(t *testing.T) {
	e := httptest.New(t, tests.App(Configure))
	t.Log(e.GET("/qrcode").WithQuery("id", "1").Expect().Status(httptest.StatusOK).Body())
}

func TestCreatePhoto(t *testing.T) {
	e := httptest.New(t, tests.App(Configure))
	fh, err := os.Open("mm.gif")
	if err != nil {
		t.Fatal(err)
	}
	defer fh.Close()
	t.Log(e.POST("/photos").
		WithMultipart().
		WithFile("file", "mm.gif", fh).
		Expect().Status(httptest.StatusOK).Body())
}
