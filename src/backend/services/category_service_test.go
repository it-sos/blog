/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package services

import (
	"gitee.com/itsos/golibs/v2/db/mysql"
	"testing"
)

func Test_categoryService_BindTag(t *testing.T) {
	type vo struct {
		cid      uint
		aid      uint
		expected interface{}
	}

	to := vo{
		1,
		1,
		nil,
	}

	t.Run("绑定解绑cache检测", func(t *testing.T) {
		SCategory.BindTag(to.cid, to.aid)
		err := SCategory.Unbind(to.cid, to.aid)
		if err != to.expected {
			t.Errorf("unbind() error = %v, expected = %v", err, nil)
			return
		}
		_, tag := SCategory.GetTopicAndTag(to.aid)
		if len(tag) > 0 {
			t.Errorf("cache clear fail. GetTopicAndTag() = %v, expected = %s", tag, "[]")
		}
	})

	t.Run("清空分类表", func(t *testing.T) {
		mysql.NewMysql().Exec("truncate table category")
		mysql.NewMysql().Exec("truncate table category_rel")
	})
}
