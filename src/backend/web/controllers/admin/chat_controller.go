package admin

import (
	"time"

	"github.com/it-sos/blog/cerrors"
	"github.com/it-sos/blog/services"
	"github.com/kataras/iris/v12"
)

// ChatController 首页控制器
type ChatController struct {
	Ctx       iris.Context
	StartTime time.Time
}

// PostChatCompletion search 内容
// @Tags 博客后台接口
// @Summary ai搜索
// @Description ai搜索
// @Accept json
// @Produce json
// @Param keyword query string true "搜索内容"
// @Success 200 {object} string ""
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/chat/completion [post]
func (c *ChatController) PostChatCompletion() (message string, err error) {
	keyword := c.Ctx.FormValues()["keyword"]
	if len(keyword) == 0 {
		err = cerrors.Error("search_keyword_err")
		return
	}
	message, err = services.SChatService.Completion(keyword)
	return
}
