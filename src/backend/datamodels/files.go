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

// Files 文件列表
type Files struct {
	Id    uint      `json:"id" readonly:"true" example:"1" xorm:"notnull autoincr pk id"`  // 文件表ID
	Aid   uint      `json:"aid" xorm:"index notnull comment('文章id')"`                      // 文章id
	Name  string    `json:"name" xorm:"varchar(128) index notnull comment('文件名')"`         // 文件名
	File  string    `json:"file" xorm:"varchar(255) notnull default '' comment('文件存储地址')"` // 文件存储地址
	Utime time.Time `json:"utime" readonly:"true" xorm:"index notnull updated"`
	Ctime time.Time `json:"ctime" readonly:"true" xorm:"notnull created"`
}
