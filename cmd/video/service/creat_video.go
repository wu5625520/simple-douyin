// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service

import (
	"context"

	// "github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo"
	"simple-douyin/cmd/video/dal/db"
	"simple-douyin/kitex_gen/videoproto"
	// "github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/note/dal/db"
)

type CreateVideoService struct {
	ctx context.Context
}

// NewCreateNoteService new CreateNoteService
func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

// 这个是dal层定义的
// type Video struct {
// 	UserID   int64  `json:"user_id"`
// 	Title    string `json:"title"`
// 	PlayAddr string `json:"play_addr"`
// 	CoverUrl string `json:"cover_url"`
// }

// CreateNote create note info
func (s *CreateVideoService) CreateVideo(req *videoproto.CreateVideoReq) error {
	VideoModel := &db.Video{
		UserID:   req.VideoBaseInfo.UserId,
		Title:    req.VideoBaseInfo.Title,
		PlayAddr: req.VideoBaseInfo.PlayAddr,
		CoverUrl: req.VideoBaseInfo.CoverAddr,
	}
	// 这里需要根据dal的返回结果进行处理
	return db.CreateNote(s.ctx, []*db.Video{VideoModel})
}
