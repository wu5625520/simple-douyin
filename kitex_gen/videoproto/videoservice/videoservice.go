// Code generated by Kitex v0.3.1. DO NOT EDIT.

package videoservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
	"simple-douyin/kitex_gen/videoproto"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*videoproto.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateVideo":       kitex.NewMethodInfo(createVideoHandler, newCreateVideoArgs, newCreateVideoResult, false),
		"GetVideosByUserId": kitex.NewMethodInfo(getVideosByUserIdHandler, newGetVideosByUserIdArgs, newGetVideosByUserIdResult, false),
		"GetVideosByTime":   kitex.NewMethodInfo(getVideosByTimeHandler, newGetVideosByTimeArgs, newGetVideosByTimeResult, false),
		"LikeVideo":         kitex.NewMethodInfo(likeVideoHandler, newLikeVideoArgs, newLikeVideoResult, false),
		"UnLikeVideo":       kitex.NewMethodInfo(unLikeVideoHandler, newUnLikeVideoArgs, newUnLikeVideoResult, false),
		"GetWhetherBeLiked": kitex.NewMethodInfo(getWhetherBeLikedHandler, newGetWhetherBeLikedArgs, newGetWhetherBeLikedResult, false),
		"GetLikesCount":     kitex.NewMethodInfo(getLikesCountHandler, newGetLikesCountArgs, newGetLikesCountResult, false),
		"CreateComment":     kitex.NewMethodInfo(createCommentHandler, newCreateCommentArgs, newCreateCommentResult, false),
		"DeleteComment":     kitex.NewMethodInfo(deleteCommentHandler, newDeleteCommentArgs, newDeleteCommentResult, false),
		"GetComments":       kitex.NewMethodInfo(getCommentsHandler, newGetCommentsArgs, newGetCommentsResult, false),
		"GetCommentsCount":  kitex.NewMethodInfo(getCommentsCountHandler, newGetCommentsCountArgs, newGetCommentsCountResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func createVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoproto.CreateVideoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoproto.VideoService).CreateVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateVideoArgs:
		success, err := handler.(videoproto.VideoService).CreateVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateVideoResult)
		realResult.Success = success
	}
	return nil
}
func newCreateVideoArgs() interface{} {
	return &CreateVideoArgs{}
}

func newCreateVideoResult() interface{} {
	return &CreateVideoResult{}
}

type CreateVideoArgs struct {
	Req *videoproto.CreateVideoReq
}

func (p *CreateVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateVideoArgs) Unmarshal(in []byte) error {
	msg := new(videoproto.CreateVideoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateVideoArgs_Req_DEFAULT *videoproto.CreateVideoReq

func (p *CreateVideoArgs) GetReq() *videoproto.CreateVideoReq {
	if !p.IsSetReq() {
		return CreateVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateVideoResult struct {
	Success *videoproto.CreateVideoResp
}

var CreateVideoResult_Success_DEFAULT *videoproto.CreateVideoResp

func (p *CreateVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateVideoResult) Unmarshal(in []byte) error {
	msg := new(videoproto.CreateVideoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateVideoResult) GetSuccess() *videoproto.CreateVideoResp {
	if !p.IsSetSuccess() {
		return CreateVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoproto.CreateVideoResp)
}

func (p *CreateVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getVideosByUserIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoproto.GetVideosByUserIdReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoproto.VideoService).GetVideosByUserId(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetVideosByUserIdArgs:
		success, err := handler.(videoproto.VideoService).GetVideosByUserId(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetVideosByUserIdResult)
		realResult.Success = success
	}
	return nil
}
func newGetVideosByUserIdArgs() interface{} {
	return &GetVideosByUserIdArgs{}
}

func newGetVideosByUserIdResult() interface{} {
	return &GetVideosByUserIdResult{}
}

type GetVideosByUserIdArgs struct {
	Req *videoproto.GetVideosByUserIdReq
}

func (p *GetVideosByUserIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetVideosByUserIdArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetVideosByUserIdArgs) Unmarshal(in []byte) error {
	msg := new(videoproto.GetVideosByUserIdReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetVideosByUserIdArgs_Req_DEFAULT *videoproto.GetVideosByUserIdReq

func (p *GetVideosByUserIdArgs) GetReq() *videoproto.GetVideosByUserIdReq {
	if !p.IsSetReq() {
		return GetVideosByUserIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetVideosByUserIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetVideosByUserIdResult struct {
	Success *videoproto.GetVideosByUserIdResp
}

var GetVideosByUserIdResult_Success_DEFAULT *videoproto.GetVideosByUserIdResp

func (p *GetVideosByUserIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetVideosByUserIdResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetVideosByUserIdResult) Unmarshal(in []byte) error {
	msg := new(videoproto.GetVideosByUserIdResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetVideosByUserIdResult) GetSuccess() *videoproto.GetVideosByUserIdResp {
	if !p.IsSetSuccess() {
		return GetVideosByUserIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetVideosByUserIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoproto.GetVideosByUserIdResp)
}

func (p *GetVideosByUserIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getVideosByTimeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoproto.GetVideosByTimeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoproto.VideoService).GetVideosByTime(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetVideosByTimeArgs:
		success, err := handler.(videoproto.VideoService).GetVideosByTime(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetVideosByTimeResult)
		realResult.Success = success
	}
	return nil
}
func newGetVideosByTimeArgs() interface{} {
	return &GetVideosByTimeArgs{}
}

func newGetVideosByTimeResult() interface{} {
	return &GetVideosByTimeResult{}
}

type GetVideosByTimeArgs struct {
	Req *videoproto.GetVideosByTimeReq
}

func (p *GetVideosByTimeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetVideosByTimeArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetVideosByTimeArgs) Unmarshal(in []byte) error {
	msg := new(videoproto.GetVideosByTimeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetVideosByTimeArgs_Req_DEFAULT *videoproto.GetVideosByTimeReq

func (p *GetVideosByTimeArgs) GetReq() *videoproto.GetVideosByTimeReq {
	if !p.IsSetReq() {
		return GetVideosByTimeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetVideosByTimeArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetVideosByTimeResult struct {
	Success *videoproto.GetVideosByTimeResp
}

var GetVideosByTimeResult_Success_DEFAULT *videoproto.GetVideosByTimeResp

func (p *GetVideosByTimeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetVideosByTimeResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetVideosByTimeResult) Unmarshal(in []byte) error {
	msg := new(videoproto.GetVideosByTimeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetVideosByTimeResult) GetSuccess() *videoproto.GetVideosByTimeResp {
	if !p.IsSetSuccess() {
		return GetVideosByTimeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetVideosByTimeResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoproto.GetVideosByTimeResp)
}

func (p *GetVideosByTimeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func likeVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoproto.LikeVideoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoproto.VideoService).LikeVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *LikeVideoArgs:
		success, err := handler.(videoproto.VideoService).LikeVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LikeVideoResult)
		realResult.Success = success
	}
	return nil
}
func newLikeVideoArgs() interface{} {
	return &LikeVideoArgs{}
}

func newLikeVideoResult() interface{} {
	return &LikeVideoResult{}
}

type LikeVideoArgs struct {
	Req *videoproto.LikeVideoReq
}

func (p *LikeVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in LikeVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *LikeVideoArgs) Unmarshal(in []byte) error {
	msg := new(videoproto.LikeVideoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LikeVideoArgs_Req_DEFAULT *videoproto.LikeVideoReq

func (p *LikeVideoArgs) GetReq() *videoproto.LikeVideoReq {
	if !p.IsSetReq() {
		return LikeVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LikeVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

type LikeVideoResult struct {
	Success *videoproto.LikeVideoResp
}

var LikeVideoResult_Success_DEFAULT *videoproto.LikeVideoResp

func (p *LikeVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in LikeVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *LikeVideoResult) Unmarshal(in []byte) error {
	msg := new(videoproto.LikeVideoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LikeVideoResult) GetSuccess() *videoproto.LikeVideoResp {
	if !p.IsSetSuccess() {
		return LikeVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LikeVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoproto.LikeVideoResp)
}

func (p *LikeVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func unLikeVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoproto.UnLikeVideoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoproto.VideoService).UnLikeVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UnLikeVideoArgs:
		success, err := handler.(videoproto.VideoService).UnLikeVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UnLikeVideoResult)
		realResult.Success = success
	}
	return nil
}
func newUnLikeVideoArgs() interface{} {
	return &UnLikeVideoArgs{}
}

func newUnLikeVideoResult() interface{} {
	return &UnLikeVideoResult{}
}

type UnLikeVideoArgs struct {
	Req *videoproto.UnLikeVideoReq
}

func (p *UnLikeVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UnLikeVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UnLikeVideoArgs) Unmarshal(in []byte) error {
	msg := new(videoproto.UnLikeVideoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UnLikeVideoArgs_Req_DEFAULT *videoproto.UnLikeVideoReq

func (p *UnLikeVideoArgs) GetReq() *videoproto.UnLikeVideoReq {
	if !p.IsSetReq() {
		return UnLikeVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UnLikeVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

type UnLikeVideoResult struct {
	Success *videoproto.UnLikeVideoResp
}

var UnLikeVideoResult_Success_DEFAULT *videoproto.UnLikeVideoResp

func (p *UnLikeVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UnLikeVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UnLikeVideoResult) Unmarshal(in []byte) error {
	msg := new(videoproto.UnLikeVideoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnLikeVideoResult) GetSuccess() *videoproto.UnLikeVideoResp {
	if !p.IsSetSuccess() {
		return UnLikeVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UnLikeVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoproto.UnLikeVideoResp)
}

func (p *UnLikeVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getWhetherBeLikedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoproto.GetWhetherBeLikedReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoproto.VideoService).GetWhetherBeLiked(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetWhetherBeLikedArgs:
		success, err := handler.(videoproto.VideoService).GetWhetherBeLiked(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetWhetherBeLikedResult)
		realResult.Success = success
	}
	return nil
}
func newGetWhetherBeLikedArgs() interface{} {
	return &GetWhetherBeLikedArgs{}
}

func newGetWhetherBeLikedResult() interface{} {
	return &GetWhetherBeLikedResult{}
}

type GetWhetherBeLikedArgs struct {
	Req *videoproto.GetWhetherBeLikedReq
}

func (p *GetWhetherBeLikedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetWhetherBeLikedArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetWhetherBeLikedArgs) Unmarshal(in []byte) error {
	msg := new(videoproto.GetWhetherBeLikedReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetWhetherBeLikedArgs_Req_DEFAULT *videoproto.GetWhetherBeLikedReq

func (p *GetWhetherBeLikedArgs) GetReq() *videoproto.GetWhetherBeLikedReq {
	if !p.IsSetReq() {
		return GetWhetherBeLikedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetWhetherBeLikedArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetWhetherBeLikedResult struct {
	Success *videoproto.GetWhetherBeLikedResp
}

var GetWhetherBeLikedResult_Success_DEFAULT *videoproto.GetWhetherBeLikedResp

func (p *GetWhetherBeLikedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetWhetherBeLikedResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetWhetherBeLikedResult) Unmarshal(in []byte) error {
	msg := new(videoproto.GetWhetherBeLikedResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetWhetherBeLikedResult) GetSuccess() *videoproto.GetWhetherBeLikedResp {
	if !p.IsSetSuccess() {
		return GetWhetherBeLikedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetWhetherBeLikedResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoproto.GetWhetherBeLikedResp)
}

func (p *GetWhetherBeLikedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getLikesCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoproto.GetLikesCountReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoproto.VideoService).GetLikesCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetLikesCountArgs:
		success, err := handler.(videoproto.VideoService).GetLikesCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetLikesCountResult)
		realResult.Success = success
	}
	return nil
}
func newGetLikesCountArgs() interface{} {
	return &GetLikesCountArgs{}
}

func newGetLikesCountResult() interface{} {
	return &GetLikesCountResult{}
}

type GetLikesCountArgs struct {
	Req *videoproto.GetLikesCountReq
}

func (p *GetLikesCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetLikesCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetLikesCountArgs) Unmarshal(in []byte) error {
	msg := new(videoproto.GetLikesCountReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetLikesCountArgs_Req_DEFAULT *videoproto.GetLikesCountReq

func (p *GetLikesCountArgs) GetReq() *videoproto.GetLikesCountReq {
	if !p.IsSetReq() {
		return GetLikesCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLikesCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetLikesCountResult struct {
	Success *videoproto.GetLikesCountResp
}

var GetLikesCountResult_Success_DEFAULT *videoproto.GetLikesCountResp

func (p *GetLikesCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetLikesCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetLikesCountResult) Unmarshal(in []byte) error {
	msg := new(videoproto.GetLikesCountResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLikesCountResult) GetSuccess() *videoproto.GetLikesCountResp {
	if !p.IsSetSuccess() {
		return GetLikesCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLikesCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoproto.GetLikesCountResp)
}

func (p *GetLikesCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func createCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoproto.CreateCommentReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoproto.VideoService).CreateComment(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateCommentArgs:
		success, err := handler.(videoproto.VideoService).CreateComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateCommentResult)
		realResult.Success = success
	}
	return nil
}
func newCreateCommentArgs() interface{} {
	return &CreateCommentArgs{}
}

func newCreateCommentResult() interface{} {
	return &CreateCommentResult{}
}

type CreateCommentArgs struct {
	Req *videoproto.CreateCommentReq
}

func (p *CreateCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateCommentArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateCommentArgs) Unmarshal(in []byte) error {
	msg := new(videoproto.CreateCommentReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateCommentArgs_Req_DEFAULT *videoproto.CreateCommentReq

func (p *CreateCommentArgs) GetReq() *videoproto.CreateCommentReq {
	if !p.IsSetReq() {
		return CreateCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateCommentResult struct {
	Success *videoproto.CreateCommentResp
}

var CreateCommentResult_Success_DEFAULT *videoproto.CreateCommentResp

func (p *CreateCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateCommentResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateCommentResult) Unmarshal(in []byte) error {
	msg := new(videoproto.CreateCommentResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateCommentResult) GetSuccess() *videoproto.CreateCommentResp {
	if !p.IsSetSuccess() {
		return CreateCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoproto.CreateCommentResp)
}

func (p *CreateCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoproto.DeleteCommentReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoproto.VideoService).DeleteComment(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteCommentArgs:
		success, err := handler.(videoproto.VideoService).DeleteComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteCommentResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteCommentArgs() interface{} {
	return &DeleteCommentArgs{}
}

func newDeleteCommentResult() interface{} {
	return &DeleteCommentResult{}
}

type DeleteCommentArgs struct {
	Req *videoproto.DeleteCommentReq
}

func (p *DeleteCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteCommentArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteCommentArgs) Unmarshal(in []byte) error {
	msg := new(videoproto.DeleteCommentReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteCommentArgs_Req_DEFAULT *videoproto.DeleteCommentReq

func (p *DeleteCommentArgs) GetReq() *videoproto.DeleteCommentReq {
	if !p.IsSetReq() {
		return DeleteCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteCommentResult struct {
	Success *videoproto.DeleteCommentResp
}

var DeleteCommentResult_Success_DEFAULT *videoproto.DeleteCommentResp

func (p *DeleteCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteCommentResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteCommentResult) Unmarshal(in []byte) error {
	msg := new(videoproto.DeleteCommentResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteCommentResult) GetSuccess() *videoproto.DeleteCommentResp {
	if !p.IsSetSuccess() {
		return DeleteCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoproto.DeleteCommentResp)
}

func (p *DeleteCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getCommentsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoproto.GetCommentsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoproto.VideoService).GetComments(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetCommentsArgs:
		success, err := handler.(videoproto.VideoService).GetComments(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCommentsResult)
		realResult.Success = success
	}
	return nil
}
func newGetCommentsArgs() interface{} {
	return &GetCommentsArgs{}
}

func newGetCommentsResult() interface{} {
	return &GetCommentsResult{}
}

type GetCommentsArgs struct {
	Req *videoproto.GetCommentsReq
}

func (p *GetCommentsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCommentsArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetCommentsArgs) Unmarshal(in []byte) error {
	msg := new(videoproto.GetCommentsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCommentsArgs_Req_DEFAULT *videoproto.GetCommentsReq

func (p *GetCommentsArgs) GetReq() *videoproto.GetCommentsReq {
	if !p.IsSetReq() {
		return GetCommentsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCommentsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetCommentsResult struct {
	Success *videoproto.GetCommentsResp
}

var GetCommentsResult_Success_DEFAULT *videoproto.GetCommentsResp

func (p *GetCommentsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCommentsResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetCommentsResult) Unmarshal(in []byte) error {
	msg := new(videoproto.GetCommentsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCommentsResult) GetSuccess() *videoproto.GetCommentsResp {
	if !p.IsSetSuccess() {
		return GetCommentsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCommentsResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoproto.GetCommentsResp)
}

func (p *GetCommentsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getCommentsCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videoproto.GetCommentsCountReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videoproto.VideoService).GetCommentsCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetCommentsCountArgs:
		success, err := handler.(videoproto.VideoService).GetCommentsCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCommentsCountResult)
		realResult.Success = success
	}
	return nil
}
func newGetCommentsCountArgs() interface{} {
	return &GetCommentsCountArgs{}
}

func newGetCommentsCountResult() interface{} {
	return &GetCommentsCountResult{}
}

type GetCommentsCountArgs struct {
	Req *videoproto.GetCommentsCountReq
}

func (p *GetCommentsCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCommentsCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetCommentsCountArgs) Unmarshal(in []byte) error {
	msg := new(videoproto.GetCommentsCountReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCommentsCountArgs_Req_DEFAULT *videoproto.GetCommentsCountReq

func (p *GetCommentsCountArgs) GetReq() *videoproto.GetCommentsCountReq {
	if !p.IsSetReq() {
		return GetCommentsCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCommentsCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetCommentsCountResult struct {
	Success *videoproto.GetCommentsCountResp
}

var GetCommentsCountResult_Success_DEFAULT *videoproto.GetCommentsCountResp

func (p *GetCommentsCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCommentsCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetCommentsCountResult) Unmarshal(in []byte) error {
	msg := new(videoproto.GetCommentsCountResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCommentsCountResult) GetSuccess() *videoproto.GetCommentsCountResp {
	if !p.IsSetSuccess() {
		return GetCommentsCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCommentsCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*videoproto.GetCommentsCountResp)
}

func (p *GetCommentsCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateVideo(ctx context.Context, Req *videoproto.CreateVideoReq) (r *videoproto.CreateVideoResp, err error) {
	var _args CreateVideoArgs
	_args.Req = Req
	var _result CreateVideoResult
	if err = p.c.Call(ctx, "CreateVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideosByUserId(ctx context.Context, Req *videoproto.GetVideosByUserIdReq) (r *videoproto.GetVideosByUserIdResp, err error) {
	var _args GetVideosByUserIdArgs
	_args.Req = Req
	var _result GetVideosByUserIdResult
	if err = p.c.Call(ctx, "GetVideosByUserId", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideosByTime(ctx context.Context, Req *videoproto.GetVideosByTimeReq) (r *videoproto.GetVideosByTimeResp, err error) {
	var _args GetVideosByTimeArgs
	_args.Req = Req
	var _result GetVideosByTimeResult
	if err = p.c.Call(ctx, "GetVideosByTime", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) LikeVideo(ctx context.Context, Req *videoproto.LikeVideoReq) (r *videoproto.LikeVideoResp, err error) {
	var _args LikeVideoArgs
	_args.Req = Req
	var _result LikeVideoResult
	if err = p.c.Call(ctx, "LikeVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UnLikeVideo(ctx context.Context, Req *videoproto.UnLikeVideoReq) (r *videoproto.UnLikeVideoResp, err error) {
	var _args UnLikeVideoArgs
	_args.Req = Req
	var _result UnLikeVideoResult
	if err = p.c.Call(ctx, "UnLikeVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetWhetherBeLiked(ctx context.Context, Req *videoproto.GetWhetherBeLikedReq) (r *videoproto.GetWhetherBeLikedResp, err error) {
	var _args GetWhetherBeLikedArgs
	_args.Req = Req
	var _result GetWhetherBeLikedResult
	if err = p.c.Call(ctx, "GetWhetherBeLiked", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetLikesCount(ctx context.Context, Req *videoproto.GetLikesCountReq) (r *videoproto.GetLikesCountResp, err error) {
	var _args GetLikesCountArgs
	_args.Req = Req
	var _result GetLikesCountResult
	if err = p.c.Call(ctx, "GetLikesCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateComment(ctx context.Context, Req *videoproto.CreateCommentReq) (r *videoproto.CreateCommentResp, err error) {
	var _args CreateCommentArgs
	_args.Req = Req
	var _result CreateCommentResult
	if err = p.c.Call(ctx, "CreateComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteComment(ctx context.Context, Req *videoproto.DeleteCommentReq) (r *videoproto.DeleteCommentResp, err error) {
	var _args DeleteCommentArgs
	_args.Req = Req
	var _result DeleteCommentResult
	if err = p.c.Call(ctx, "DeleteComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetComments(ctx context.Context, Req *videoproto.GetCommentsReq) (r *videoproto.GetCommentsResp, err error) {
	var _args GetCommentsArgs
	_args.Req = Req
	var _result GetCommentsResult
	if err = p.c.Call(ctx, "GetComments", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCommentsCount(ctx context.Context, Req *videoproto.GetCommentsCountReq) (r *videoproto.GetCommentsCountResp, err error) {
	var _args GetCommentsCountArgs
	_args.Req = Req
	var _result GetCommentsCountResult
	if err = p.c.Call(ctx, "GetCommentsCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
