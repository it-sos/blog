/*
   Copyright (c) [2021] IT.SOS
   study-notes is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package services

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/it-sos/golibs/db/es"
	"github.com/kataras/golog"
)

// "id": {
//     "type": "long"
// },
// "is_state": {
//     "type": "long"
// },
// "is_del": {
//     "type": "long"
// },
// "utime": {
//     "type": "date",
//     "format": "strict_date_optional_time||epoch_millis"
// },
// "ctime": {
//     "type": "date",
//     "format": "strict_date_optional_time||epoch_millis"
// },
// "title": {
//     "type": "text",
//     "analyzer": "ik_max_word",
//     "search_analyzer": "ik_smart"
// },
// "intro": {
//     "type": "text",
//     "analyzer": "ik_max_word",
//     "search_analyzer": "ik_smart"
// },
// "data": {
//     "type": "text",
//     "analyzer": "ik_max_word",
//     "search_analyzer": "ik_smart"
// }

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

func EsSync(esData EsData, wg *sync.WaitGroup) (err error) {
	ej, _ := json.Marshal(esData)
	es := es.NewEs()
	res, err := es.Index("canal", strings.NewReader(string(ej)),
		es.Index.WithDocumentID(strconv.Itoa(int(esData.Id))),
		es.Index.WithRefresh("true"),
	)
	if err != nil {
		golog.Errorf("Error getting response: %s", err)
		return
	}
	if wg != nil {
		defer wg.Done()
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err = json.NewDecoder(res.Body).Decode(&e); err != nil {
			golog.Errorf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			golog.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
		return
	}

	var r map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		golog.Errorf("Error parsing the response body: %s", err)
		return
	}
	if len(r) > 0 {
		golog.Info(r)
	} else {
		golog.Error(esData.Id)
	}
	return
}
