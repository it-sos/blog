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
	"gitee.com/itsos/golibs/v2/db/mysql"
	"gitee.com/itsos/studynotes/datamodels"
	"log"
)

type UserRepository interface {
	// Insert 新增
	Insert(account, password, salt string) uint
	// DeleteByUid 删除
	DeleteByUid(uid uint) bool
	// ExistAccount 用户存在性校验
	ExistAccount(account string) bool
	// SelectByAccount 查询单条
	SelectByAccount(account string) (datamodels.User, bool)
	// SelectByUid 查询单条
	SelectByUid(uid uint) (datamodels.User, bool)
}

type userRepository struct {
	db mysql.GoLibMysql
}

func (u userRepository) Insert(account, password, salt string) uint {
	user := &datamodels.User{
		Account:  account,
		Password: password,
		Salt:     salt,
	}
	_, err := u.db.Insert(user)
	if err != nil {
		panic(err)
	}
	return user.Id
}

func (u userRepository) DeleteByUid(uid uint) bool {
	user := &datamodels.User{Id: uid}
	i, err := u.db.Delete(user)
	if err != nil {
		log.Panicf("delete user fail. error:%v", err)
	}
	return i > 0
}

func (u userRepository) ExistAccount(account string) bool {
	user := &datamodels.User{Account: account}
	isExits, err := u.db.Exist(user)
	if err != nil {
		panic(err)
	}
	return isExits
}

func (u userRepository) SelectByAccount(account string) (datamodels.User, bool) {
	user := &datamodels.User{Account: account}
	has, err := u.db.Get(user)
	if err != nil {
		panic(err)
	}
	return *user, has
}

func (u userRepository) SelectByUid(uid uint) (datamodels.User, bool) {
	user := &datamodels.User{Id: uid}
	has, err := u.db.Get(user)
	if err != nil {
		panic(err)
	}
	return *user, has
}

var RUser UserRepository = &userRepository{mysql.NewMysql()}
