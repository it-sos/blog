package controller_tests

import (
	"encoding/base64"
	"gitee.com/itsos/studynotes"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"os"
	"testing"
)

func TestFilesController_Files(t *testing.T) {
	e := httptest.New(t, studynotes.NewApp().Application)
	png := "/tmp/john.png"
	pngData := "iVBORw0KGgoAAAANSUhEUgAAAAIAAAACCAYAAABytg0kAAAAFElEQVQImWMQFRX9Lyoq+p8BxgAAJ04E+dJh/q8AAAAASUVORK5CYII="

	t.Run("上传文件", func(t *testing.T) {
		f, _ := os.Create(png)
		defer f.Close()
		decoded, _ := base64.StdEncoding.DecodeString(pngData)
		f.Write(decoded)

		r := e.POST("/admin/files").
			WithMultipart().
			WithFile("file", png).
			WithForm(map[string]interface{}{
				"aid": "1",
			}).
			Expect().Status(iris.StatusOK)
		t.Log(r.Body())
	})
}
