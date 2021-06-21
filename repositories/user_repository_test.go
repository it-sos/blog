/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package repositories

import (
	_ "gitee.com/itsos/golibs/tests/testsdb"
	"gitee.com/itsos/studynotes/datamodels"
	"testing"
)

var user = &datamodels.User{Name: "peng.yu"}

func connect() UserRepository {
	return NewUserRepository()
}

func Test_userRepository_InsertOrUpdate(t *testing.T) {
	t.Log(connect().Insert(user))
}

func Test_userRepository_SelectMany(t *testing.T) {
	t.Log(connect().SelectMany(user, 0, 3))
}

func Test_userRepository_Select(t *testing.T) {
	t.Log(connect().Select(user))
}
