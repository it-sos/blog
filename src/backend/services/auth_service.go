/*
   Copyright (c) [2021] IT.SOS
   study-notes is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package services

import (
	"gitee.com/itsos/golibs/v2/framework/iris/services"
	"gitee.com/itsos/golibs/v2/utils"
	"gitee.com/itsos/golibs/v2/utils/crypt"
	"gitee.com/itsos/studynotes/cerrors"
	"gitee.com/itsos/studynotes/repositories"
	"strconv"
	"time"
)

// AuthService 认证模块
type AuthService interface {
	// IsLogin 获取登录状态
	IsLogin(token string) bool
	// Login 登录
	Login(account, password string, loginFree bool) (string, error)
	// Register 注册
	Register(account, password string) error
}

type authService struct {
	ru repositories.UserRepository
}

func (a authService) Login(account string, password string, loginFree bool) (token string, err error) {
	if !a.ru.ExistAccount(account) {
		err = cerrors.Error("user_auth_fail")
		return
	}

	u, _ := a.ru.SelectByAccount(account)
	password = crypt.Md5(password + u.Salt)
	if u.Password != password {
		err = cerrors.Error("user_auth_fail")
		return
	}

	var lifetime time.Duration = 0
	// 免登7天，否则4小时
	if !loginFree {
		lifetime = time.Hour * 4
	}

	token = services.GetToken(strconv.Itoa(int(u.Id)), lifetime)
	return
}

func (a authService) Register(account, password string) (err error) {
	if a.ru.ExistAccount(account) {
		err = cerrors.Error("user_account_exits")
		return
	}
	salt := utils.Rand(32, utils.RandMix)
	password = crypt.Md5(password + salt)
	a.ru.Insert(account, password, salt)
	return err
}

func (a authService) IsLogin(token string) bool {
	_, err := services.GetLoginId(token)
	return err == nil
}

var SAuthService AuthService = &authService{ru: repositories.RUser}
