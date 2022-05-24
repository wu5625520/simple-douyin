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

package db

import (
	"context"
)

type Video struct {
	UserID   int64  `json:"user_id"`
	Title    string `json:"title"`
	PlayAddr string `json:"play_addr"`
	CoverUrl string `json:"cover_url"`
}

// 所有方法都是空方法，直接返回成功
// CreateNote create note info
func CreateNote(ctx context.Context, notes []*Video) error {
	return nil
}

// DeleteNote delete note info
