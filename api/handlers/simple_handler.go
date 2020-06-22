package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kdpujie/log4go"
	"github.com/xwi88/kit4go/datetime"

	"github.com/xwi88/go-demo-gin/constants"
	"github.com/xwi88/go-demo-gin/model"
	"github.com/xwi88/go-demo-gin/pb"
)

// SimpleHandler handler
func SimpleHandler(c *gin.Context) {
	requestID := c.GetString(constants.RequestID)
	rsp := &model.UniResponse{
		Code: constants.ResponseCodeOK,
		MSG:  constants.ResponseMSG[constants.ResponseCodeOK],
	}
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
	// contentType := c.GetHeader(constant.HTTPHeaderKeyContentType)
	ctx := context.WithValue(c, constants.RequestID, requestID)
	defer ctx.Done()

	w := c.Writer
	var requestData model.SimpleRequest
	err := c.BindQuery(&requestData)
	if err != nil {
		log4go.Error("[SimpleHandler] error:%v", err.Error())
		rsp.Code = constants.ResponseCodeOK
		rsp.MSG = constants.ResponseMSG[rsp.Code]
		writeResponseJsonDefault(c, bl, w, rsp)
		return
	}
	log4go.Debug("[SimpleHandler] origin request data:%+v", requestData)
	rsp.Data = datetime.NowTimestampStr(datetime.LayoutDateTimeISO8601ZoneP8Mid)
	errCode, err := writeResponseJSON(c, bl, w, rsp)
	if err != nil {
		log4go.Error("response err_code:%v, err:%v", errCode, err.Error())
	}
}
