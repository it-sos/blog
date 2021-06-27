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

type AccessTimes interface {
	Id(k uint) AccessTimesCmd
	Rank(num int) []string
}

type accessTimes struct{}

type AccessTimesCmd interface {
	Get() int
	Incr()
}

type accessTimesCmd struct {
	k string
}

func (a *accessTimesCmd) Get() int {
	return int(db.Rdb.ZScore(ctx, root, a.k).Val())
}

func (a *accessTimesCmd) Incr() {
	err := db.Rdb.ZIncrBy(ctx, root, 1, a.k).Err()
	if err != nil {
		panic(err)
	}
}

const prefixTimes = "%d"
const root = "access_times_page"

var ctx = context.Background()

func (a *accessTimes) Id(k uint) AccessTimesCmd {
	return &accessTimesCmd{fmt.Sprintf(prefixTimes, k)}
}

func (a *accessTimes) Rank(num int) []string {
	return db.Rdb.ZRevRange(ctx, root, 0, int64(num-1)).Val()
}

var CAccessTimes AccessTimes = &accessTimes{}
