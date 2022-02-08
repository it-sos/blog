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
	"fmt"
	"gitee.com/itsos/golibs/v2/utils"
	"gitee.com/itsos/golibs/v2/utils/crypt"
	"gitee.com/itsos/golibs/v2/utils/crypt/aes"
	"gitee.com/itsos/studynotes/cerrors"
	"gitee.com/itsos/studynotes/config"
	"gitee.com/itsos/studynotes/repositories"
	"log"
	"strconv"
	"strings"
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
	GetUserId(token string) (string, error)
	GetToken(userId string) string
}

type authService struct {
	ru repositories.UserRepository
}

func (a authService) Login(account string, password string, loginFree bool) (token string, err error) {
	//a.ru.
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
	return false
}

func (u authService) GetUserId(token string) (userId string, err error) {
	if token == "" {
		err = cerrors.Error("unauthorized_access")
		return
	}
	deToken, err := aes.JavaDecryptCBC(token, config.C.GetCryptAesToken())
	if err != nil {
		log.Panicf("token aes decode fail: %v, data: %v", err, token)
	}
	tokenSplit := strings.Split(deToken, "|")
	if len(tokenSplit) != 2 {
		log.Panicf("token is error: %v", deToken)
	}
	unix, err := strconv.ParseInt(tokenSplit[1], 10, 64)
	if err != nil {
		log.Panicf("token unix time parseInt err: %v, data: %v", err, tokenSplit[1])
	}
	tts := time.Unix(unix, 0)
	// token有效时间一周
	tts = tts.Add(time.Hour * 24 * 7)
	if time.Now().After(tts) {
		err = cerrors.Error("unauthorized_access")
	} else {
		userId = tokenSplit[0]
	}
	return
}

func (u authService) GetToken(userId string) string {
	tk := fmt.Sprintf("%s|%d", userId, time.Now().Unix())
	token, err := aes.JavaEncryptCBC(tk, config.C.GetCryptAesToken())
	if err != nil {
		log.Panicf("token aes encode fail: %v, data: %v", err, tk)
	}
	return token
}

var SAuthService AuthService = &authService{ru: repositories.RUser}
