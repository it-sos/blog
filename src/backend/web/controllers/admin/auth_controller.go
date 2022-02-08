/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package admin

import (
	"gitee.com/itsos/studynotes/cerrors"
	"gitee.com/itsos/studynotes/models/vo"
	"gitee.com/itsos/studynotes/services"
	"github.com/kataras/iris/v12"
)

type AuthController struct {
	Ctx iris.Context
}

// PostLogin
// @Tags 博客后台接口
// @Summary 登录验证
// @Description 校验成功返回token
// @Accept json
// @Produce json
// @Param body body vo.LoginParamsVO true "json参数"
// @Success 200 {string} string ""
// @Failure 400 {object} cerrors.Errors "error"
// @Router /admin/login [Post]
func (c *AuthController) PostLogin() (token string, err error) {
	params := new(vo.LoginParamsVO)
	err = c.Ctx.ReadJSON(params)
	if err != nil {
		err = cerrors.Errorf("parameter_error", err)
		return
	}
	token, err = services.SAuthService.Login(params.Account, params.Password, params.LoginFree)
	return
}
