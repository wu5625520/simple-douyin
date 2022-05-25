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
	"simple-douyin/kitex_gen/videoproto"
)

type Video struct {
	VideoId  int64  `json:"video_id"`
	UserID   int64  `json:"user_id"`
	Title    string `json:"title"`
	PlayAddr string `json:"play_addr"`
	CoverUrl string `json:"cover_url"`
}

// 所有方法都是空方法，直接返回成功
// CreateNote create video info
func CreateVideo(ctx context.Context, notes []*Video) error {
	return nil
}

//  根据用户id查视频
func MGetVideoByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	example := &Video{
		VideoId:  1111,
		UserID:   1223,
		Title:    "test_Title",
		PlayAddr: "test_PlayAddr",
		CoverUrl: "test_CoverUrl",
	}
	var res []*Video
	res = append(res, example)
	return res, nil
}

// 返回视频点赞数
func GetLikeCount(vid int64) (int64, error) {
	return 666, nil
}

// 返回视频评论数
func GetCommentCount(vid int64) (int64, error) {
	return 777, nil
}

// 返回是否点赞数
func IsFavorite(vid int64, uid int64) (bool, error) {
	return true, nil
}

// 根据时间戳返回最近count个视频,还需要返回next time
func MGetVideoByTime(ctx context.Context, LatestTime int64, Count int64) ([]*Video, int64, error) {
	example := &Video{
		VideoId:  1111,
		UserID:   1223,
		Title:    "test_Title",
		PlayAddr: "test_PlayAddr",
		CoverUrl: "test_CoverUrl",
	}
	var res []*Video
	res = append(res, example)
	example2 := &Video{
		VideoId:  2222,
		UserID:   3333,
		Title:    "test_Title2",
		PlayAddr: "test_PlayAddr2",
		CoverUrl: "test_CoverUrl2",
	}
	res = append(res, example2)
	return res, 999, nil
}

// 点赞视频
func LikeVideo(ctx context.Context, UserId int64, VideoId int64) error {
	return nil
}

// 取消点赞视频
func UnLikeVideo(ctx context.Context, UserId int64, VideoId int64) error {
	return nil
}

// 新增评论,需要da层返回评论详情
func CreateComment(ctx context.Context, UserId int64, VideoId int64, Content string) (*videoproto.CommentInfo, error) {
	example := &videoproto.CommentInfo{
		CommentId:  8888,
		UserId:     1223,
		Content:    "test_comment",
		CreateDate: "test_create_data",
	}
	return example, nil
}

// 删除评论
func DeleteComment(ctx context.Context, CommentId int64) error {
	return nil
}

// 查询评论,需要da层返回评论详情,有可能有多条评论
func GetComment(ctx context.Context, AppUserId int64, VideoId int64) ([]*videoproto.CommentInfo, error) {
	example := &videoproto.CommentInfo{
		CommentId:  4444,
		UserId:     1255523,
		Content:    "test_get_comment",
		CreateDate: "test_create_data",
	}
	var comments []*videoproto.CommentInfo
	comments = append(comments, example)
	example2 := &videoproto.CommentInfo{
		CommentId:  4444,
		UserId:     1255523,
		Content:    "test_get_comment",
		CreateDate: "test_create_data",
	}
	comments = append(comments, example2)
	return comments, nil
}
