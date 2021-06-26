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
	"gitee.com/itsos/studynotes/datamodels"
)

type CategoryRelRepository interface {
	// Insert 新增
	Insert(p *datamodels.CategoryRel) (id uint)
	// Update 更新
	Update(p *datamodels.CategoryRel) (id uint)
	// Select 查询文章详细
	Select(p *datamodels.CategoryRel) (bool, datamodels.CategoryRel)
	// SelectMany 查询文章列表
	SelectMany(state []uint8, offset int, limit int) (results []datamodels.Article)
	SelectManyByIds(ids []string) []datamodels.Article
}

type categoryRelRepository struct {
}
