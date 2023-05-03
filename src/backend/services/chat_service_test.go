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

import "testing"

func Test_chatService_Completion(t *testing.T) {
	type args struct {
		askMessage []string
	}
	tests := []struct {
		name             string
		c                ChatService
		args             args
		wantReplyMessage string
		wantErr          bool
	}{
		{
			"test",
			SChatService,
			args{
				[]string{"架构师如何学习？"},
			},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &chatService{}
			gotReplyMessage, err := c.Completion(tt.args.askMessage)
			if (err != nil) != tt.wantErr {
				t.Errorf("chatService.Completion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotReplyMessage)
			// if gotReplyMessage != tt.wantReplyMessage {
			// 	t.Errorf("chatService.Completion() = %v, want %v", gotReplyMessage, tt.wantReplyMessage)
			// }
		})
	}
}
