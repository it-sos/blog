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
	"strconv"
)

type CategoryService interface {
	// GetNameById 通过分类id获取分类名
	GetNameById(cid uint) string
	// GetTopicAndTag 通过文章id获取专题与标签list id
	GetTopicAndTag(aid uint) (topic, tag []string)
	// GetArticleListIds 获取文章列表id
	GetArticleListIds(cid uint) (aid []string)
	// GetTopicAndTagName 获取标签专题名list
	GetTopicAndTagName(aid uint) (topic, tag []string)
}

type categoryService struct {
}

func (c categoryService) GetArticleListIds(cid uint) (aid []string) {
	aids := repositories.RCategoryRel.SelectManyByCid(cid)
	for _, v := range aids {
		aid = append(aid, strconv.Itoa(int(v.Aid)))
	}
	return
}

func (c categoryService) GetNameById(cid uint) string {
	cache := caches.CCategory.Id(cid)
	if cache.Exists() {
		return cache.Get()
	}
	category, has := repositories.RCategory.Select(&datamodels.Category{Id: cid})
	if has {
		cache.Set(category.Name)
	}
	return category.Name
}

// GetTopicAndTag 专题标签id
func (c categoryService) GetTopicAndTag(aid uint) (topic, tag []string) {
	tags := caches.CCategoryRel.Id(aid, repositories.CategoryRelTypeTag)
	topics := caches.CCategoryRel.Id(aid, repositories.CategoryRelTypeTopic)
	if !(tags.Exists() && topics.Exists()) {
		for _, v := range repositories.RCategoryRel.SelectManyByAid(aid) {
			if v.Type == repositories.CategoryRelTypeTag {
				tags.Add(v.Cid)
			} else {
				topics.Add(v.Cid)
			}
		}
	}
	topic = topics.Get()
	tag = tags.Get()
	return
}

// GetTopicAndTagName 专题标签名称
func (c categoryService) GetTopicAndTagName(aid uint) (topic, tag []string) {
	topicIds, tagIds := c.GetTopicAndTag(aid)
	for _, v := range topicIds {
		cid, _ := strconv.Atoi(v)
		topic = append(topic, c.GetNameById(uint(cid)))
	}
	for _, v := range tagIds {
		cid, _ := strconv.Atoi(v)
		tag = append(tag, c.GetNameById(uint(cid)))
	}
	return
}

var SCategory CategoryService = &categoryService{}
