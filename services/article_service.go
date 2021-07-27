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
	"gitee.com/itsos/golibs/db/es"
	"gitee.com/itsos/golibs/utils"
	"gitee.com/itsos/golibs/utils/validate"
	"gitee.com/itsos/studynotes/caches"
	"gitee.com/itsos/studynotes/datamodels"
	"gitee.com/itsos/studynotes/models/vo"
	"gitee.com/itsos/studynotes/repositories"
	"golang.org/x/net/context"
	"golang.org/x/net/html"
	"io/ioutil"
	"strconv"
	"strings"
)

type ArticleService interface {
	// GetRank 获取前访问前50的文章列表
	GetRank(isLogin bool) []vo.ArticleAccessTimesVO
	// GetListPage 获取最新文章列表
	GetListPage(isLogin bool, page int, size int, keyword string) []vo.ArticleVO
	// GetContent 获取文章详情
	GetContent(isLogin bool, title string) vo.ArticleContentVO
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

//{
//  "took" : 4,
//  "timed_out" : false,
//  "_shards" : {
//    "total" : 1,
//    "successful" : 1,
//    "skipped" : 0,
//    "failed" : 0
//  },
//  "hits" : {
//    "total" : {
//      "value" : 3,
//      "relation" : "eq"
//    },
//    "max_score" : 9.455517,
//    "hits" : [
//      {
//        "_index" : "canal",
//        "_type" : "_doc",
//        "_id" : "160",
//        "_score" : 9.455517,
//        "_source" : {
//          "is_state" : 2,
//          "data" : "&lt;h3&gt;一、简介&lt;/h3&gt;&lt;p&gt;PHP 用来管理依赖（dependency）关系的工具。你可以在自己的项目中声明所依赖的外部工具库（libraries），Composer 会帮你安装 这些依赖的库文件。&lt;br/&gt;&lt;/p&gt;&lt;h3&gt;二、安装&lt;/h3&gt;&lt;p&gt;1、php方式安装&lt;br/&gt;&lt;/p&gt;&lt;pre class=&quot;brush:bash;toolbar:false&quot;&gt;$&amp;nbsp;php&amp;nbsp;-r&amp;nbsp;&amp;quot;copy(&amp;#39;https://install.phpcomposer.com/installer&amp;#39;,&amp;nbsp;&amp;#39;composer-setup.php&amp;#39;);&amp;quot;\n$&amp;nbsp;php&amp;nbsp;composer-setup.php\n$&amp;nbsp;php&amp;nbsp;-r&amp;nbsp;&amp;quot;unlink(&amp;#39;composer-setup.php&amp;#39;);&amp;quot;\n#&amp;nbsp;升级最近版本\n$&amp;nbsp;composer&amp;nbsp;selfupdate&lt;/pre&gt;&lt;p&gt;2、linux方式安装&lt;/p&gt;&lt;pre&gt;curl&amp;nbsp;-sS&amp;nbsp;https://getcomposer.org/installer&amp;nbsp;|&amp;nbsp;php&lt;/pre&gt;&lt;h3&gt;三、使用&lt;/h3&gt;&lt;p&gt;&lt;a href=&quot;http://docs.phpcomposer.com/01-basic-usage.html&quot; target=&quot;_blank&quot; title=&quot;基本使用&quot;&gt;基本使用&lt;/a&gt;&lt;/p&gt;&lt;h3&gt;四、镜像&lt;/h3&gt;&lt;p&gt;&lt;a href=&quot;https://pkg.phpcomposer.com/&quot; target=&quot;_blank&quot; title=&quot;中国镜像&quot;&gt;中国镜像&lt;/a&gt;&lt;br/&gt;&lt;/p&gt;&lt;pre class=&quot;brush:bash;toolbar:false&quot;&gt;$&amp;nbsp;composer&amp;nbsp;config&amp;nbsp;repo.packagist&amp;nbsp;composer&amp;nbsp;https://packagist.phpcomposer.com&lt;/pre&gt;&lt;p&gt;&lt;br/&gt;&lt;/p&gt;",
//          "intro" : "一、简介PHP 用来管理依赖（dependency）关系的工具。你可以在自己的项目中声明所依赖的外部工具库（libraries），Composer 会帮你安装 这些依赖的库文件。二、安装1、php方式安装$&nbsp;php&nbsp;-r&nbsp;&quot;copy(&#39;https://install.phpcomposer.com/installer&#39;,&nbsp;&#39;",
//          "title" : "composer php项目依赖管理工具"
//        },
//        "highlight" : {
//          "data" : [
//            "amp;#39;);&amp;quot;\n#&amp;nbsp;升级最近版本\n$&amp;nbsp;composer&amp;nbsp;selfupdate&lt;/pre&gt;&lt;p&gt;2、<font-size='red'>linux</font>",
//            "方式安装&lt;/p&gt;&lt;pre&gt;<font-size='red'>curl</font>&amp;nbsp;-sS&amp;nbsp;https://getcomposer.org/installer&amp;nbsp;|&amp;"
//          ]
//        }
//      },
// es 搜索返回数据结构
type (
	EsSource struct {
		IsState int    `json:"is_state"`
		Intro   string `json:"intro"`
		Title   string `json:"title"`
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

func (a articleService) GetRank(isLogin bool) []vo.ArticleAccessTimesVO {
	// 获取访问量的前50条
	rank := a.accessTimes.Rank(50)
	return a.getArticleList(isLogin, rank)
}

func (a articleService) GetListPage(isLogin bool, page int, size int, keyword string) []vo.ArticleVO {
	page = validate.IntRange(page, 1, 100000)
	size = validate.IntRange(size, 1, 100000)
	offset := (page - 1) * size
	if keyword != "" {
		a.search(isLogin, keyword, offset, size)
		return nil
	}
	article := a.article.SelectMany(a.getAuthorize(isLogin), offset, size)
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
		cid, _ := strconv.Atoi(v)
		topics = append(
			topics,
			vo.TopicVO{
				Title:   SCategory.GetNameById(uint(cid)),
				Article: a.getArticleList(isLogin, SCategory.GetArticleListIds(uint(cid))),
			},
		)
	}
	for _, v := range tag {
		cid, _ := strconv.Atoi(v)
		tags = append(
			tags,
			vo.TagVO{
				Title:   SCategory.GetNameById(uint(cid)),
				Article: a.getArticleList(isLogin, SCategory.GetArticleListIds(uint(cid))),
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
		Duration: utils.TimeDuration(v.Utime),
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
			"pre_tags":  []string{"<font-size='red'>"},
			"post_tags": []string{"</font>"},
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
		newEs.Search.WithSourceExcludes("is_del"),
		newEs.Search.WithAnalyzer("ik_smart"),
		// https://lucene.apache.org/core/2_9_4/queryparsersyntax.html
		newEs.Search.WithQuery(fmt.Sprintf("%s AND is_del:0 AND is_state:(%s)", keyword, isState)),
		newEs.Search.WithDefaultOperator("and"),
		newEs.Search.WithBody(&buf),
		// newEs.Search.WithErrorTrace(),
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
		panic(err)
	}
	return esResultDTO
}
