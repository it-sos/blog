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

// AuthService 认证模块
type AuthService interface {
	// IsLogin 获取登录状态
	IsLogin(token string) bool
	// Login 登录
	Login(username string, password string) string
	// Register 注册
	Register(username, password string)
}

type authService struct {
}

func (a authService) Login(username string, password string) string {
	panic("implement me")
}

func (a authService) Register(username, password string) {
	panic("implement me")
}

func (a authService) IsLogin(token string) bool {
	return false
}

var SAuthService AuthService = &authService{}
