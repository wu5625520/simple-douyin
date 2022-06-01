package video

import (
	"simple-douyin/cmd/api/forms"
	"simple-douyin/cmd/api/respond"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/kitex_gen/videoproto"
	"simple-douyin/pkg/constants"

	"github.com/gin-gonic/gin"
)

func LikeList(c *gin.Context) {
	appUserID := c.GetInt64(constants.IdentityKey)
	param := new(forms.LikeListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	req := &videoproto.GetLikeVideosReq{
		AppUserId: appUserID,
		UserId:    param.UserId,
	}
	videos, err := rpc.GetLikeVideos(c, req)
	if err != nil {
		respond.Error(c, err)
		return
	}
	n := len(videos)
	authors := make([]*userproto.UserInfo, n)
	for i := 0; i < n; i++ {
		subReq := &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    videos[i].VideoBaseInfo.UserId,
		}
		authors[i], err = rpc.GetUser(c, subReq)
		if err != nil {
			respond.Error(c, err)
			return
		}
	}
	resp := &respond.LikeListResp{
		BaseResp:  respond.Success,
		VideoList: respond.PackVideos(videos, authors),
	}
	respond.Send(c, resp)
}
