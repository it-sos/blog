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
	"gitee.com/itsos/golibs/v2/db/redis"
	"golang.org/x/net/context"
)

type Category interface {
	Id(k uint) CategoryCmd
}

type category struct{}

type CategoryCmd interface {
	Get() string
	Set(v string)
	Exists() bool
	Remove() bool
}

type categoryCmd struct {
	k  string
	db redis.GoLibRedis
}

func (a *categoryCmd) Remove() bool {
	return a.db.HDel(context.Background(), categoryRoot, a.k).Val() > 0
}

func (a *categoryCmd) Exists() bool {
	return a.db.HExists(context.Background(), categoryRoot, a.k).Val()
}

func (a *categoryCmd) Get() string {
	return a.db.HGet(context.Background(), categoryRoot, a.k).Val()
}

func (a *categoryCmd) Set(v string) {
	_, err := a.db.HSet(context.Background(), categoryRoot, a.k, v).Result()
	if err != nil {
		panic(err)
	}
}

const categoryPrefix = "%d"
const categoryRoot = "category"

func (a *category) Id(k uint) CategoryCmd {
	return &categoryCmd{fmt.Sprintf(categoryPrefix, k), redis.NewRedis()}
}

// CCategory cache分类id与分类名
var CCategory Category = &category{}
