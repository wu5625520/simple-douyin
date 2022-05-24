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

package handlers

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"

	"simple-douyin/kitex_gen/videoproto"

	"simple-douyin/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

// CreateNote create note info
func CreateVideo(c *gin.Context) {
	var videoVar VideoParam
	if err := c.ShouldBind(&videoVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	// claims := jwt.ExtractClaims(c)
	// userID := int64(claims[constants.IdentityKey].(float64))
	videoBaseInfo := &videoproto.VideoBaseInfo{
		UserId:    videoVar.UserId,
		PlayAddr:  videoVar.PlayAddr,
		CoverAddr: videoVar.CoverAddr,
		Title:     videoVar.Title,
	}

	err := rpc.CreateVideo(context.Background(), &videoproto.CreateVideoReq{
		VideoBaseInfo: videoBaseInfo,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
