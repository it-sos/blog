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
	"time"
)

type AuthTimes interface {
	Key(k string) AuthTimesCmd
}

type AuthTimesCmd interface {
	Get() int64
	Incr() int64
	Decr() int64
	Clear() bool
	expire()
}

type authTimes struct{}

type authTimesCmd struct {
	k string
}

func (a *authTimesCmd) Get() int64 {
	times, _ := db.Rdb.Get(context.Background(), a.k).Int64()
	return times
}

var NAuthTimes = NewAuthTimes()

func NewAuthTimes() AuthTimes {
	return &authTimes{}
}

const prefixTimes = "auth_times_%s"

func (a *authTimes) Key(k string) AuthTimesCmd {
	return &authTimesCmd{fmt.Sprintf(prefixTimes, k)}
}

const ttlTimes = 3 * time.Hour

func (a *authTimesCmd) expire() {
	if err := db.Rdb.Expire(context.Background(), a.k, ttlTimes).Err(); err != nil {
		panic(err)
	}
}

func (a *authTimesCmd) Decr() int64 {
	decr, err := db.Rdb.Decr(context.Background(), a.k).Result()
	if err != nil {
		panic(err)
	}
	a.expire()
	return decr
}

func (a *authTimesCmd) Clear() bool {
	err := db.Rdb.Del(context.Background(), a.k).Err()
	if err != nil {
		panic(err)
	}
	return true
}

func (a *authTimesCmd) Incr() int64 {
	incr, err := db.Rdb.Incr(context.Background(), a.k).Result()
	if err != nil {
		panic(err)
	}
	a.expire()
	return incr
}
