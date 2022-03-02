package main

import (
	"encoding/json"
	"fmt"
	"gitee.com/itsos/golibs/v2/db/es"
	"gitee.com/itsos/golibs/v2/db/mysql"
	"gitee.com/itsos/studynotes/datamodels"
	"github.com/kataras/golog"
	"log"
	"strconv"
	"strings"
	"time"
)

type EsData struct {
	Id      uint      `json:"id"`       // 文章表ID
	Title   string    `json:"title"`    // 标题
	Intro   string    `json:"intro"`    // 简介
	IsState uint8     `json:"is_state"` // 状态1私密；2公开
	IsDel   uint8     `json:"is_del"`   // 状态1未删除；2已删除
	Data    string    `json:"data"`     // 文章内容
	Utime   time.Time `json:"utime"`
	Ctime   time.Time `json:"ctime"`
}

func main() {
	db := mysql.NewMysql()
	article := make([]datamodels.Article, 0)
	err := db.Where("is_del=?", 1).Desc("utime").Find(&article)
	if err != nil {
		panic(err)
	}

	content := datamodels.ArticleContent{}
	esData := EsData{}
	for _, v := range article {
		db.Where("aid=?", v.Id).Find(&content)
		esData.Id = v.Id
		esData.Title = v.Title
		esData.Intro = v.Intro
		esData.IsState = v.IsState
		esData.Utime = v.Utime
		esData.Ctime = v.Ctime
		esData.IsDel = v.IsDel
		esData.Data = content.Data
		esSync(esData)
	}
}

var r map[string]interface{}

func esSync(esData EsData) {
	//        "id": {
	//            "type": "long"
	//        },
	//        "is_state": {
	//            "type": "long"
	//        },
	//        "is_del": {
	//            "type": "long"
	//        },
	//        "utime": {
	//            "type": "date",
	//            "format": "strict_date_optional_time||epoch_millis"
	//        },
	//        "ctime": {
	//            "type": "date",
	//            "format": "strict_date_optional_time||epoch_millis"
	//        },
	//        "title": {
	//            "type": "text",
	//            "analyzer": "ik_max_word",
	//            "search_analyzer": "ik_smart"
	//        },
	//        "intro": {
	//            "type": "text",
	//            "analyzer": "ik_max_word",
	//            "search_analyzer": "ik_smart"
	//        },
	//        "data": {
	//            "type": "text",
	//            "analyzer": "ik_max_word",
	//            "search_analyzer": "ik_smart"
	//        }
	ej, _ := json.Marshal(esData)
	golog.Info()
	es := es.NewEs()
	res, err := es.Index("canal", strings.NewReader(string(ej)),
		es.Index.WithDocumentID(strconv.Itoa(int(esData.Id))),
		es.Index.WithRefresh("true"),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	fmt.Println(r)
}
