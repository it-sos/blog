package routes

import (
	"gitee.com/itsos/studynotes/tests"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestConfigure(t *testing.T) {
	t.Log(httptest.New(t, tests.App(Configure)).
		GET("/article/list").Expect().Status(httptest.StatusOK).Body())
}
