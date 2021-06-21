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
	"fmt"
	"gitee.com/itsos/studynotes/config"
	"github.com/armon/circbuf"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
)

var c = config.C

type PhotoService interface {
	GetQrcode(id string) []byte
}

func NewPhotoService() PhotoService {
	return &photoService{}
}

type photoService struct{}

func (p photoService) GetQrcode(id string) []byte {
	url := fmt.Sprintf(c.GetQUrl(), id)
	qrCode, _ := qr.Encode(url, qr.M, qr.Unicode)
	qrCode, _ = barcode.Scale(qrCode, c.GetQWidth(), c.GetQHeight())
	buf, _ := circbuf.NewBuffer(c.GetQBuffer())
	defer buf.Reset()
	png.Encode(buf, qrCode)
	return buf.Bytes()
}
