// Package handlers route handler
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kdpujie/log4go"

	"github.com/xwi88/go-demo-gin/constants"
	"github.com/xwi88/go-demo-gin/model"
	"github.com/xwi88/go-demo-gin/pb"
)

// writeResponseJSON json 格式数据返回
func writeResponseJSON(c *gin.Context, bl *pb.BusLog, w http.ResponseWriter, rsp *model.UniResponse) (int, error) {
	w.Header().Set(constants.HTTPHeaderKeyContentType, constants.HTTPHeaderValApplicationJSON)
	respData, err := json.Marshal(rsp)
	if err != nil {
		if bl.Response == nil {
			bl.Response = new(pb.Response)
		}
		bl.Response.Code = int32(rsp.Code)
		bl.Response.RequestId = rsp.RequestID
		bl.Response.Message = rsp.MSG
		c.Set(constants.ContextBusinessLog, bl)
		return http.StatusInternalServerError, err
	}
	if _, exist := c.Get(constants.ContextResponseData); !exist {
		c.Set(constants.ContextResponseData, rsp) // 填充返回消息到 context, 记录日志
	}
	return w.Write(respData)
}

// writeResponseJsonDefault
func writeResponseJsonDefault(c *gin.Context, bl *pb.BusLog, w http.ResponseWriter, rsp *model.UniResponse) {
	log4go.Error("response code:%v, msg:%v", rsp.Code, rsp.MSG)
	_, err := writeResponseJSON(c, bl, w, rsp)
	if err != nil {
		log4go.Error("write response err:%v", err.Error())
	}
	return
}

// writeResponseJsonWithRequestID
func writeResponseJsonWithRequestID(c *gin.Context, bl *pb.BusLog, w http.ResponseWriter, requestID string, rsp *model.UniResponse) {
	log4go.Error("request[%v], code:%v, msg:%v", requestID, rsp.Code, rsp.MSG)
	_, err := writeResponseJSON(c, bl, w, rsp)
	if err != nil {
		log4go.Error("request[%v], write response err:%v", requestID, err.Error())
	}
	return
}

func requestFilterMethod(r *http.Request, rsp *model.UniResponse, method string) bool {
	if r.Method != http.MethodPost {
		rsp.MSG = fmt.Sprintf("请求的资源仅支持 [%v] 方法", method)
		rsp.Code = constants.RequestMethodNotAllowed
		return false
	}
	return true
}
