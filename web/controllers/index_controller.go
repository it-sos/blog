/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package controllers

import (
	"gitee.com/itsos/golibs/global/variable"
	"gitee.com/itsos/golibs/utils"
	"gitee.com/itsos/studynotes/errors"
	"gitee.com/itsos/studynotes/models/vo"
	"gitee.com/itsos/studynotes/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"mime/multipart"
)

type IndexController struct {
	S   services.PhotoService
	Ctx iris.Context
}

var (
	uploadDir = variable.BasePath + "/web/public/photos"
)

// PostPhotos
// @Tags 图库管理
// @Summary 上传照片
// @Description 上传照片文件
// @Accept  multipart/form-data
// @Produce json
// @Param file formData file true "request file data"
// @Success 200 {object} vo.QrcodeVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Router /photos [post]
func (c *IndexController) PostPhotos() (vo.QrcodeVO, error) {
	file, _, err := c.Ctx.UploadFormFiles(uploadDir, beforeSave)
	if err != nil {
		return vo.QrcodeVO{}, errors.Error("upload_err")
	}
	return vo.QrcodeVO{Id: file[0].Filename}, nil
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) bool {
	file.Filename = utils.Rand(32) + ".gif"
	return true
}

// GetQrcode
// @Tags 图库管理
// @Summary 展示二维码
// @Description 通过Qrcode ID返回二维码图片流
// @Accept json
// @Produce json
// @Param id query string true "Qrcode Id"
// @Success 200 {string} string "图片流"
// @Failure 400 {object} errors.Errors "error"
// @Router /qrcode [get]
func (c *IndexController) GetQrcode() (mvc.Result, error) {
	id := c.Ctx.FormValue("id")

	if id == "" {
		return nil, errors.Error("id_param_err")
	}
	return mvc.Response{
		ContentType: "image/png",
		Content:     c.S.GetQrcode(id),
	}, nil
}
