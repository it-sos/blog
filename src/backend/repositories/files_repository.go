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
	"github.com/it-sos/blog/datamodels"
	"github.com/it-sos/golibs/v2/db/mysql"
)

type FilesRepository interface {
	// Insert 新增
	Insert(p *datamodels.Files) (id uint)
	// Delete 删除
	Delete(fileName string) bool
	// ExistName 名称存在性判断
	ExistName(aid uint, name string) bool
	// SelectManyByAid 查询多条
	SelectManyByAid(aid uint) []datamodels.Files
}

type filesRepository struct {
	db mysql.GoLibMysql
}

func (c filesRepository) Insert(p *datamodels.Files) (id uint) {
	_, err := c.db.Insert(p)
	if err != nil {
		panic(err)
	}
	return p.Id
}

func (c filesRepository) SelectManyByAid(aid uint) (results []datamodels.Files) {
	results = make([]datamodels.Files, 0)
	err := c.db.Where("aid=?", aid).Desc("id").Find(&results)
	if err != nil {
		panic(err)
	}
	return
}

func (c filesRepository) Delete(fileName string) bool {
	affected, err := c.db.Where("file=?", fileName).Delete(new(datamodels.Files))
	if err != nil {
		panic(err)
	}
	return affected > 0
}

func (c filesRepository) ExistName(aid uint, name string) bool {
	file := datamodels.Files{Name: name, Aid: aid}
	isExits, err := c.db.Exist(&file)
	if err != nil {
		panic(err)
	}
	return isExits
}

var RFiles FilesRepository = &filesRepository{mysql.NewMysql()}
