package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kdpujie/log4go"
	"github.com/xwi88/version"

	"github.com/xwi88/go-demo-gin/constants"
	"github.com/xwi88/go-demo-gin/model"
	"github.com/xwi88/go-demo-gin/pb"
)

// VersionHandler version
func VersionHandler(c *gin.Context) {
	rsp := &model.UniResponse{
		Code: constants.ResponseCodeOK,
		MSG:  constants.ResponseMSG[constants.ResponseCodeOK],
	}
	rsp.Data = version.Get()

	// FIXME ...
	requestID := c.GetString(constants.RequestID)
	var bl *pb.BusLog
	blData, exist := c.Get(constants.ContextBusinessLog)
	if !exist {
		bl = new(pb.BusLog)
	} else {
		switch _data := blData.(type) {
		case *pb.BusLog:
			bl = _data
		}
	}
	if bl != nil && len(bl.RequestId) == 0 {
		bl.RequestId = requestID
	}
	// FIXME ...

	w := c.Writer
	errCode, err := writeResponseJSON(c, bl, w, rsp)
	if err != nil {
		log4go.Error("response err_code:%v, err:%v", errCode, err.Error())
	}
}
