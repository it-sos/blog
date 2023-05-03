package index_controller_test

import (
	"testing"

	"github.com/it-sos/blog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func TestArticleList(t *testing.T) {
	t.Run("获取文章列表", func(t *testing.T) {
		e := httptest.New(t, blog.NewApp().Application).
			GET("/api/article/list").Expect().Status(iris.StatusOK)
		t.Log(e.Body())
	})
}
