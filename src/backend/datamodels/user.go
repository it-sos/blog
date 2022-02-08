/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package datamodels

import "time"

// User 用户信息
type User struct {
	Id       uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"` // uid
	Account  string    `json:"account" xorm:"varchar(128) index notnull comment('标题')"`      // 账号
	Password string    `json:"password" xorm:"varchar(32) notnull comment('密码')"`            // 密码
	Salt     string    `json:"salt" xorm:"varchar(32) notnull comment('盐')"`                 // 盐
	Utime    time.Time `json:"utime" readonly:"true" xorm:"notnull updated"`
	Ctime    time.Time `json:"ctime" readonly:"true" xorm:"notnull created"`
}
