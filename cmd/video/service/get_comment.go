package service

import (
	"context"

	"simple-douyin/cmd/video/dal/db"
	"simple-douyin/kitex_gen/videoproto"
)

type GetCommentService struct {
	ctx context.Context
}

//
func NewGetCommentService(ctx context.Context) *GetCommentService {
	return &GetCommentService{ctx: ctx}
}

//
func (s *GetCommentService) GetComment(req *videoproto.GetCommentsReq) ([]*videoproto.CommentInfo, error) {

	// 查询评论,需要da层返回评论详情,有可能有多条评论
	return db.GetComment(s.ctx, req.AppUserId, req.VideoId)
}
