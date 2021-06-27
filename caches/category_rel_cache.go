/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package caches

import (
	"fmt"
	"gitee.com/itsos/golibs/db"
	"golang.org/x/net/context"
)

type CategoryRel interface {
	Id(aid uint, types uint8) CategoryRelCmd
}

type categoryRel struct{}

type CategoryRelCmd interface {
	Get() []string
	Add(v uint)
	Exists() bool
}

type categoryRelCmd struct {
	aidType string
}

func (a *categoryRelCmd) Exists() bool {
	return db.Rdb.Exists(context.Background(), a.aidType).Val() > 0
}

func (a *categoryRelCmd) Get() []string {
	return db.Rdb.SMembers(context.Background(), a.aidType).Val()
}

func (a *categoryRelCmd) Add(v uint) {
	_, err := db.Rdb.SAdd(context.Background(), a.aidType, v).Result()
	if err != nil {
		panic(err)
	}
}

const categoryRelPrefix = "%d_%d"

func (a *categoryRel) Id(aid uint, types uint8) CategoryRelCmd {
	return &categoryRelCmd{fmt.Sprintf(categoryRelPrefix, aid, types)}
}

var CCategoryRel CategoryRel = &categoryRel{}
