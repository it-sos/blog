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
	"testing"

	"github.com/it-sos/golibs/utils/crypt"
)

func Test_authService_Register(t *testing.T) {
	u := "admin"
	p := crypt.Md5("123456")
	if err := SAuthService.Register(u, p); err != nil {
		t.Error("create account fail.", err)
	}

	token, err := SAuthService.Login(u, p, true)
	if err != nil {
		t.Error("login faill.", err)
	}

	if !SAuthService.IsLogin(token) {
		t.Error("validate login fail.")
	}

	//uu, _ := repositories.RUser.SelectByAccount(u)
	//repositories.RUser.DeleteByUid(uu.Id)

}
