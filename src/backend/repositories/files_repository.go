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
)

type FilesRepository interface {
	// Insert 新增
	Insert(p *datamodels.Files) (id uint)
	// Delete 删除
	Delete(id uint) bool
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

func (c filesRepository) Delete(id uint) bool {
	affected, err := c.db.ID(id).Delete(new(datamodels.Category))
	if err != nil {
		panic(err)
	}
	return affected > 0
}

var RFiles FilesRepository = &filesRepository{mysql.NewMysql()}
