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
	"bytes"
	"encoding/json"
	"fmt"
	"gitee.com/itsos/golibs/v2/db/es"
	"gitee.com/itsos/golibs/v2/utils"
	"gitee.com/itsos/golibs/v2/utils/validate"
	"gitee.com/itsos/studynotes/caches"
	"gitee.com/itsos/studynotes/cerrors"
	"gitee.com/itsos/studynotes/datamodels"
	"gitee.com/itsos/studynotes/models/vo"
	"gitee.com/itsos/studynotes/repositories"
	"github.com/kataras/golog"
	"golang.org/x/net/context"
	"golang.org/x/net/html"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// es 搜索返回数据结构
type (
	EsSource struct {
		IsState uint8     `json:"is_state"`
		Intro   string    `json:"intro"`
		Title   string    `json:"title"`
		Utime   time.Time `json:"utime"`
		Ctime   time.Time `json:"ctime"`
	}
	EsHighlight struct {
		Title []string `json:"title"`
		Intro []string `json:"intro"`
	}
	EsSubHits struct {
		Id        string      `json:"_id"`
		Source    EsSource    `json:"_source"`
		Highlight EsHighlight `json:"highlight"`
	}
	EsTotal struct {
		Value int `json:"value"`
	}
	EsHits struct {
		Total   EsTotal     `json:"total"`
		SubHits []EsSubHits `json:"hits"`
	}
	EsResultDTO struct {
		Hits EsHits `json:"hits"`
	}
)

// ArticleService 标记为「后台」的方法按登录状态下的逻辑处理
type ArticleService interface {
	// GetRank 获取前访问前50的文章列表
	GetRank(isLogin bool) []vo.ArticleAccessTimesVO
	// GetListPage 获取最新文章列表
	GetListPage(isLogin bool, page int, size int, keyword string) []vo.ArticleVO
	// GetContent 获取文章详情
	GetContent(isLogin bool, title string) vo.ArticleContentVO

	// SaveArticle 保存文章「后台」
	SaveArticle(vo vo.ArticleParamsVO) (id uint, err error)
	// DeleteArticle 删除文章&内容「后台」
	DeleteArticle(id uint) error
	// GetArticleAndContent 查询文章及相关信息「后台」
	GetArticleAndContent(id uint) (art vo.ArticleEditVO, err error)
	// GetArticleList 获取文章列表 [后台]
	GetArticleList(page int, size int, keyword string) []vo.ArticleListVO
	// EsSync 同步es
	EsSync(id uint)
}

var SArticle ArticleService = &articleService{
	article:     repositories.RArticle,
	content:     repositories.RArticleContent,
	accessTimes: caches.CAccessTimes}

type articleService struct {
	article     repositories.ArticleRepository
	content     repositories.ArticleContentRepository
	accessTimes caches.AccessTimes
}

func (a articleService) EsSync(id uint) {
	v, has := a.article.GetInfoById(id)
	if !has {
		golog.Errorf("article not found: [id=%v]", id)
		return
	}
	content, _ := a.content.Select(&datamodels.ArticleContent{
		Aid: v.Id,
	})
	esData := EsData{}
	esData.Id = v.Id
	esData.Title = v.Title
	esData.Intro = v.Intro
	esData.IsState = v.IsState
	esData.Utime = v.Utime
	esData.Ctime = v.Ctime
	esData.IsDel = v.IsDel
	esData.Data = content.Data
	EsSync(esData, nil)
}

// GetArticleList 获取文章列表 [后台]
func (a articleService) GetArticleList(page int, size int, keyword string) []vo.ArticleListVO {
	article := a.getListPage(true, page, size, keyword)
	articleVO := make([]vo.ArticleListVO, 0)
	if len(article) > 0 {
		for _, v := range article {
			articleVO = append(articleVO, vo.ArticleListVO{
				Id:         v.Id,
				Title:      v.Title,
				TitleMatch: v.TitleMatch,
				Duration:   a.getDuration(v),
				IsState:    v.IsState,
			})
		}
	}
	return articleVO
}

func (a articleService) GetArticleAndContent(id uint) (info vo.ArticleEditVO, err error) {
	article, has := a.article.GetInfoById(id)
	if !has {
		err = cerrors.Error("article_notfound_err")
		return
	}
	content, _ := a.content.Select(&datamodels.ArticleContent{
		Aid: article.Id,
	})

	// 解密处理
	if article.IsEncrypt == repositories.IsEncrypt {
		// todo 通过用户私钥解密 intro、content 字段
	}

	info = vo.ArticleEditVO{
		Id:        id,
		Title:     article.Title,
		Intro:     article.Intro,
		IsState:   article.IsState,
		IsEncrypt: article.IsEncrypt,
	}

	info.Content = html.UnescapeString(content.Data)
	info.Topics, info.Tags = SCategory.GetTopicAndTag(id)
	return info, nil
}

func (a articleService) SaveArticle(vo vo.ArticleParamsVO) (id uint, err error) {
	// 验证标题是否存在
	if a.article.TitleExists(vo.Title) {
		var title string
		if vo.Id > 0 {
			info, _ := a.article.GetInfoById(vo.Id)
			title = info.Title
		}
		if vo.Id < 1 || title != "" && title != vo.Title {
			err = cerrors.Error("article_exists_err")
			return
		}
	}
	vo.Content = html.EscapeString(vo.Content)
	vo.Intro = html.EscapeString(vo.Intro)

	if vo.IsEncrypt == repositories.IsEncrypt {
		// todo  此处通过公钥加密处理 vo.Content,vo.Intro
	}

	article := datamodels.Article{
		Title:     vo.Title,
		Intro:     vo.Intro,
		IsState:   vo.IsState,
		IsEncrypt: vo.IsEncrypt,
		IsDel:     repositories.NotDeleted,
	}

	if vo.Id > 0 {
		// 更新
		a.article.UpdateTrans(vo.Id, &article, vo.Content)
	} else {
		// 新增
		vo.Id = a.article.InsertTrans(&article, vo.Content)
	}
	id = vo.Id
	a.EsSync(id)
	return
}

func (a articleService) DeleteArticle(id uint) (err error) {
	caches.CCategoryRel.Id(id, repositories.CategoryTypeTag)
	caches.CCategoryRel.Id(id, repositories.CategoryTypeTopic)
	if !a.article.SoftDelete(id) {
		err = cerrors.Error("article_remove_err")
	}
	a.EsSync(id)
	return
}

func (a articleService) GetRank(isLogin bool) []vo.ArticleAccessTimesVO {
	// 获取访问量的前50条
	rank := a.accessTimes.Rank(50)
	return a.getArticleList(isLogin, rank)
}

func (a articleService) GetListPage(isLogin bool, page int, size int, keyword string) []vo.ArticleVO {
	article := a.getListPage(isLogin, page, size, keyword)
	articleVO := make([]vo.ArticleVO, 0)
	if len(article) > 0 {
		for _, v := range article {
			v.Intro = html.UnescapeString(v.Intro)
			articleVO = append(articleVO, a.getArticle(v))
		}
	}
	return articleVO
}

// GetContent 文章内容
func (a articleService) GetContent(isLogin bool, title string) vo.ArticleContentVO {
	article, has := a.article.Select(a.getAuthorize(isLogin), &datamodels.Article{Title: title})
	if !has {
		return vo.ArticleContentVO{}
	}
	content, _ := a.content.Select(&datamodels.ArticleContent{
		Aid: article.Id,
	})
	a.accessTimes.Id(article.Id).Incr()
	content.Data = html.UnescapeString(content.Data)
	prevTitle, nextTitle := a.article.Navigation(a.getAuthorize(isLogin), article.Utime)
	topics, tags := a.getTopicAndTagArticle(isLogin, article.Id)
	return vo.ArticleContentVO{
		Navigation: vo.NavigationVO{
			PrevTitle: prevTitle,
			NextTitle: nextTitle,
		},
		Article:        a.getArticle(article),
		ArticleContent: content,
		Topics:         topics,
		Tags:           tags,
	}
}

// 获取专题和标签文章列表
func (a articleService) getTopicAndTagArticle(isLogin bool, id uint) (topics []vo.TopicVO, tags []vo.TagVO) {
	topic, tag := SCategory.GetTopicAndTag(id)

	for _, v := range topic {
		topics = append(
			topics,
			vo.TopicVO{
				Title:   SCategory.GetNameById(v),
				Article: a.getArticleList(isLogin, SCategory.GetArticleListIds(v)),
			},
		)
	}
	for _, v := range tag {
		tags = append(
			tags,
			vo.TagVO{
				Title:   SCategory.GetNameById(v),
				Article: a.getArticleList(isLogin, SCategory.GetArticleListIds(v)),
			},
		)
	}
	return
}

// 获取文章信息
func (a articleService) getArticle(v datamodels.Article) vo.ArticleVO {
	topics, tags := SCategory.GetTopicAndTagName(v.Id)
	return vo.ArticleVO{
		Article:  v,
		Duration: a.getDuration(v),
		Topics:   topics,
		Tags:     tags}
}

// 获取查看文章的权限
func (a articleService) getAuthorize(isLogin bool) (state []uint8) {
	state = []uint8{repositories.IsStatePublic}
	if isLogin {
		state = append(state, repositories.IsStatePrivate)
	}
	return
}

// 获取小栏目的文章列表
func (a articleService) getArticleList(isLogin bool, ids []string) []vo.ArticleAccessTimesVO {
	articleVO := make([]vo.ArticleAccessTimesVO, 0)
	if len(ids) > 0 {
		article := a.article.SelectManyByIds(a.getAuthorize(isLogin), ids)
		if len(article) > 0 {
			for _, v := range article {
				articleVO = append(articleVO, vo.ArticleAccessTimesVO{
					Title:       v.Title,
					AccessTimes: a.accessTimes.Id(v.Id).Get()})
			}
		}
	}
	return articleVO
}

// search 搜索关键词
func (a articleService) search(isLogin bool, keyword string, offset, size int) *EsResultDTO {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"highlight": map[string]interface{}{
			"pre_tags":  []string{"<font color='#F56C6C'><i><b>"},
			"post_tags": []string{"</b></i></font>"},
			"fields": map[string]interface{}{
				"title": map[string]interface{}{},
				"intro": map[string]interface{}{},
				"data":  map[string]interface{}{},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		panic("Error encoding query:" + err.Error())
	}
	// Perform the search request.
	isState := strings.Replace(strings.Trim(fmt.Sprint(a.getAuthorize(isLogin)), "[]"), " ", " OR ", 1)
	newEs := es.NewEs()
	res, err := newEs.Search(
		newEs.Search.WithContext(context.Background()),
		newEs.Search.WithIndex("canal"),
		newEs.Search.WithFrom(offset),
		newEs.Search.WithSize(size),
		//newEs.Search.WithSourceExcludes("is_del"),
		newEs.Search.WithAnalyzer("ik_smart"),
		// https://lucene.apache.org/core/2_9_4/queryparsersyntax.html
		newEs.Search.WithQuery(fmt.Sprintf("%s AND is_del:%d AND is_state:(%s)", keyword, repositories.NotDeleted, isState)),
		newEs.Search.WithDefaultOperator("and"),
		newEs.Search.WithBody(&buf),
		newEs.Search.WithErrorTrace(),
		newEs.Search.WithPretty(),
	)

	if err != nil {
		panic("Error getting response: " + err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("read body fail." + err.Error())
	}
	esResultDTO := &EsResultDTO{}

	if err = json.Unmarshal(body, esResultDTO); err != nil {
		print(err.Error() + res.String())
	}
	return esResultDTO
}

// 获取文章列表（分页）
func (a articleService) getListPage(isLogin bool, page int, size int, keyword string) []datamodels.Article {
	page = validate.IntRange(page, 1, 100000)
	size = validate.IntRange(size, 1, 100000)
	offset := (page - 1) * size

	var article []datamodels.Article
	if keyword != "" {
		es := a.search(isLogin, keyword, offset, size)
		for _, v := range es.Hits.SubHits {
			var titleMatch, introMatch string
			if len(v.Highlight.Title) > 0 {
				titleMatch = v.Highlight.Title[0]
			}
			if len(v.Highlight.Intro) > 0 {
				introMatch = v.Highlight.Intro[0]
			}
			id, _ := strconv.Atoi(v.Id)
			article = append(article, datamodels.Article{
				Id:         uint(id),
				Title:      v.Source.Title,
				TitleMatch: titleMatch,
				Intro:      v.Source.Intro,
				IntroMatch: introMatch,
				IsState:    v.Source.IsState,
				Utime:      v.Source.Utime,
				Ctime:      v.Source.Ctime,
			})
		}
	} else {
		article = a.article.SelectMany(a.getAuthorize(isLogin), offset, size)
	}
	return article
}

func (a articleService) getDuration(v datamodels.Article) string {
	u := utils.TimeDuration(v.Utime)
	c := utils.TimeDuration(v.Ctime)
	if u == c {
		return c
	}
	return u + " ◷ " + c + " ◶"
}
