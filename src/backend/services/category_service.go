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
	"gitee.com/itsos/studynotes/errors"
	"gitee.com/itsos/studynotes/repositories"
	"strconv"
)

// CategoryService 标记为「后台」的方法按登录状态下的逻辑处理
type CategoryService interface {
	// GetNameById 通过分类id获取分类名
	GetNameById(cid uint) string
	// GetTopicAndTag 通过文章id获取专题与标签list id
	GetTopicAndTag(aid uint) (topic, tag []string)
	// GetArticleListIds 获取文章列表id
	GetArticleListIds(cid uint) (aid []string)
	// GetTopicAndTagName 获取标签专题名list
	GetTopicAndTagName(aid uint) (topic, tag []string)

	// Unbind 解除绑定关系「后台」
	Unbind(id uint, aid uint)
	// NewTopic 新增专题「后台」
	NewTopic(name string, aid uint) uint
	// DeleteTopic 删除专题「后台」
	DeleteTopic(id uint)
	// UpdateTopic 更新专题「后台」
	UpdateTopic(id uint, name string)
	// GetTopicList 查询专题列表「后台」
	GetTopicList() []datamodels.Category
	// NewTag 新增标签「后台」
	NewTag(name string, aid uint) uint
	// DeleteTag 删除标签「后台」
	DeleteTag(id uint)
	// UpdateTag 更新标签「后台」
	UpdateTag(id uint, name string)
	// GetTagList 查询标签列表「后台」
	GetTagList() []datamodels.Category
	// GetBindArtCount 查询绑定的文章条数「后台」
	GetBindArtCount(id uint) uint
}

type categoryService struct {
	cyr repositories.CategoryRelRepository
	cy  repositories.CategoryRepository
}

func (c categoryService) Unbind(id uint, aid uint) {
	c.cyr.DeleteByAidAndCid(aid, id)
	caches.CCategoryRel.Id(aid, repositories.CategoryTypeTopic).Remove(id)
}

func (c categoryService) GetBindArtCount(id uint) uint {
	return c.cyr.GetCountByCid(id)
}

func (c categoryService) NewTopic(name string, aid uint) uint {
	if c.cy.ExistName(name, repositories.CategoryTypeTopic) {
		panic(errors.Error("topic_exists_err"))
	}
	id := c.cy.Insert(&datamodels.Category{
		Name: name,
		Type: repositories.CategoryTypeTopic,
	})
	c.cyr.Insert(&datamodels.CategoryRel{
		Aid:  aid,
		Cid:  id,
		Type: repositories.CategoryTypeTopic,
	})
	caches.CCategoryRel.Id(aid, repositories.CategoryTypeTopic).Add(id)
	return id
}

func (c categoryService) DeleteTopic(id uint) {
	if !c.cy.Delete(id) {
		panic(errors.Error("topic_remove_err"))
	}
	// 移除相关cache
	caches.CCategory.Id(id).Remove()
	aids := c.cyr.SelectManyByCid(id)
	for _, v := range aids {
		caches.CCategoryRel.Id(v.Aid, repositories.CategoryTypeTopic).Remove(id)
	}
	// 删除绑定关系
	c.cyr.DeleteByCid(id)
}

func (c categoryService) UpdateTopic(id uint, name string) {
	if c.cy.ExistName(name, repositories.CategoryTypeTopic) {
		panic(errors.Error("topic_exists_err"))
	}
	if c.cy.Update(id, &datamodels.Category{
		Name: name,
	}) {
		caches.CCategory.Id(id).Set(name)
	}
}

func (c categoryService) GetTopicList() []datamodels.Category {
	return c.cy.SelectManyByType(repositories.CategoryTypeTopic)
}

func (c categoryService) NewTag(name string, aid uint) uint {
	if c.cy.ExistName(name, repositories.CategoryTypeTag) {
		panic(errors.Error("topic_exists_err"))
	}
	id := c.cy.Insert(&datamodels.Category{
		Name: name,
		Type: repositories.CategoryTypeTag,
	})
	c.cyr.Insert(&datamodels.CategoryRel{
		Aid:  aid,
		Cid:  id,
		Type: repositories.CategoryTypeTag,
	})
	caches.CCategoryRel.Id(aid, repositories.CategoryTypeTag).Add(id)
	return id
}

func (c categoryService) DeleteTag(id uint) {
	if !c.cy.Delete(id) {
		panic(errors.Error("tag_remove_err"))
	}
	// 移除与之相关的 cache
	caches.CCategory.Id(id).Remove()
	aids := c.cyr.SelectManyByCid(id)
	for _, v := range aids {
		caches.CCategoryRel.Id(v.Aid, repositories.CategoryTypeTag).Remove(id)
	}
	// 删除绑定关系
	c.cyr.DeleteByCid(id)
}

func (c categoryService) UpdateTag(id uint, name string) {
	if c.cy.ExistName(name, repositories.CategoryTypeTag) {
		panic(errors.Error("tag_exists_err"))
	}
	if c.cy.Update(id, &datamodels.Category{
		Name: name,
	}) {
		caches.CCategory.Id(id).Set(name)
	}
}

func (c categoryService) GetTagList() []datamodels.Category {
	return c.cy.SelectManyByType(repositories.CategoryTypeTag)
}

func (c categoryService) GetArticleListIds(cid uint) (aid []string) {
	aids := c.cyr.SelectManyByCid(cid)
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
	category, has := c.cy.Select(&datamodels.Category{Id: cid})
	if has {
		cache.Set(category.Name)
	}
	return category.Name
}

// GetTopicAndTag 专题标签id
func (c categoryService) GetTopicAndTag(aid uint) (topic, tag []string) {
	tags := caches.CCategoryRel.Id(aid, repositories.CategoryTypeTag)
	topics := caches.CCategoryRel.Id(aid, repositories.CategoryTypeTopic)
	if !(tags.Exists() && topics.Exists()) {
		for _, v := range c.cyr.SelectManyByAid(aid) {
			if v.Type == repositories.CategoryTypeTag {
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

var SCategory CategoryService = &categoryService{
	cyr: repositories.RCategoryRel,
	cy:  repositories.RCategory,
}
