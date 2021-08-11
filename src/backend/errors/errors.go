/*
   Copyright (c) [2021] IT.SOS
   kn is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package errors

import (
	"encoding/json"
	"errors"
)

type Errors struct {
	ErrCode int    `json:"errCode" example:"4000000"`
	Message string `json:"message"`
}

var errCodeList = map[string]Errors{
	"param_err": {4002001, "参数异常"},
}

func Error(key string) error {
	ret, _ := json.Marshal(errCodeList[key])
	return errors.New(string(ret))
}
