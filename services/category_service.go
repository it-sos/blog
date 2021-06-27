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
	"gitee.com/itsos/studynotes/caches"
	"gitee.com/itsos/studynotes/datamodels"
	"gitee.com/itsos/studynotes/repositories"
)

type CategoryService interface {
	// GetNameById 通过分类id获取分类名
	GetNameById(id uint) string
	// GetTopicAndTag 通过文章id获取专题与标签
	GetTopicAndTag(aid uint) (topic, tag []string)
}

type categoryService struct {
}

func (c categoryService) GetNameById(id uint) string {
	cache := caches.CCategory.Id(id)
	if cache.Exists() {
		return cache.Get()
	}
	category := repositories.RCategory.SelectMany()
	name := ""
	if len(category) > 0 {
		for _, v := range category {
			caches.CCategory.Id(v.Id).Set(v.Name)
			if v.Id == id {
				name = v.Name
			}
		}
	}
	return name
}

func (c categoryService) GetTopicAndTag(aid uint) (topic, tag []string) {
	tags := caches.CCategoryRel.Id(aid, repositories.CategoryRelTypeTag)
	topics := caches.CCategoryRel.Id(aid, repositories.CategoryRelTypeTopic)
	if tags.Exists() || topics.Exists() {
		tag = tags.Get()
		topic = topics.Get()
		return
	}
	for _, v := range repositories.RCategoryRel.SelectMany(&datamodels.CategoryRel{Aid: aid}) {
		caches.CCategoryRel.Id(aid, v.Type).Add(v.Cid)
	}
	// todo
	return nil, nil
}

var SCategory CategoryService = &categoryService{}
