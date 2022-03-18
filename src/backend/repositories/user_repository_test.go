/*
   Copyright (c) [2021] IT.SOS
   study-notes is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package repositories

import (
	"gitee.com/itsos/golibs/v2/db/mysql"
	"testing"
)

func Test_userRepository_Account(t *testing.T) {
	t.Run("用户信息", func(t *testing.T) {
		account := "itsos"
		password := "123456"
		salt := "jadflajsdf8979asfasdf`"
		if RUser.ExistAccount(account) {
			t.Errorf("account is exists, expected=false actual=true")
			return
		}

		if RUser.Insert(account, password, salt) != 1 {
			t.Errorf("expected=1")
			return
		}

		user, exits := RUser.SelectByAccount(account)
		if user.Account != account || !exits {
			t.Errorf("account not exits.")
			return
		}

		user, exits = RUser.SelectByUid(user.Id)
		if user.Account != account || !exits {
			t.Errorf("uid not exits.")
			return
		}
	})

	t.Run("清空用户表", func(t *testing.T) {
		mysql.NewMysql().Exec("truncate table user")
	})
}

func BenchmarkUserRepository_DeleteByUid(b *testing.B) {
	
}
