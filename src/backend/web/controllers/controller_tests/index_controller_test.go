package index_controller_test

import (
	"github.com/it-sos/blog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestArticleList(t *testing.T) {
	t.Run("获取文章列表", func(t *testing.T) {
		e := httptest.New(t, blog.NewApp().Application).
			GET("/article/list").Expect().Status(iris.StatusOK)
		t.Log(e.Body())
	})
}
