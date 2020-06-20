package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kdpujie/log4go"
	"github.com/xwi88/kit4go/utils"

	"github.com/xwi88/go-demo-gin/constants"
	"github.com/xwi88/go-demo-gin/model"
	"github.com/xwi88/go-demo-gin/pb"
)

// Logger 日志拦截，只要进入具体handler, 则必存在 request|response
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		if method == http.MethodHead && path == "/" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		bl := new(pb.BusLog)

		recordOriginRequest(c, bl) // 记录原始请求信息

		c.Next()

		// c := c1.Copy()
		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		// request_id 未传入的话，会自动生成一个
		requestID := c.GetString(constants.RequestID)
		// 请求服务器内花费时间: ms
		costMills := float64(latency.Microseconds()) / 1000

		requestAPIData := model.RequestAPIData{
			HTTPStatus:       statusCode,
			RequestID:        requestID,
			RequestTime:      start,
			ResponseTime:     end,
			RequestTimestamp: start.Unix(),
			CostMills:        utils.Round(costMills, 6),
		}
		rateLimit, exist := c.Get(constants.ContextRequestRateLimit)
		if exist {
			if requestAPIData.ExtraMSG == nil {
				requestAPIData.ExtraMSG = make(map[string]interface{})
			}
			requestAPIData.ExtraMSG["rate_limit"] = rateLimit
		}
		requestData, exist := c.Get(constants.ContextRequestOrigin)
		if exist {
			switch _data := requestData.(type) {
			case model.RequestLog:
				requestAPIData.Request = _data
			default:
				log4go.Error("request data type err, with request_id(%v)",
					requestID)
			}
		}
		responseData, exist2 := c.Get(constants.ContextResponseData)
		uniResponseData := new(model.UniResponse)
		if exist2 {
			requestAPIData.Response = responseData

			var responseDataMap map[string]interface{}
			responseDataMapB, _ := json.Marshal(responseData)
			_ = json.Unmarshal(responseDataMapB, &responseDataMap)
			uniResponseData.RequestID = requestID
			uniResponseData.Code = int(responseDataMap["code"].(float64))
			uniResponseData.MSG = responseDataMap["msg"].(string)
			uniResponseData.Data = responseDataMap["data"]
		}

		lB, _ := json.Marshal(requestAPIData)
		log4go.Info("%v", string(lB))
		if statusCode == http.StatusOK {
			log4go.Info("[eval-api] | %3d | %13v | %15s | %7s  %s | %s",
				statusCode, latency, clientIP, method, path, requestID)
		} else {
			log4go.Error("[eval-api] | %3d | %13v | %15s | %7s  %s | %s",
				statusCode, latency, clientIP, method, path, requestID)
		}

		if method == http.MethodPost && path == "/api/v1/" {
			tokenMD5 := c.GetString(constants.ContextRequestAppTokenMD5)
			channel := c.GetString(constants.ContextRequestChannel)
			blData, exist3 := c.Get(constants.ContextBusinessLog)
			if exist3 {
				switch _data := blData.(type) {
				case *pb.BusLog:
					bl = _data
				default:
					log4go.Error("eval log type err, with request_id(%v)",
						requestID)
				}
			}
			bl.HttpStatusCode = int32(statusCode)
			bl.RequestId = requestID
			bl.Request.Channel = channel
			bl.Request.TokenMd5 = tokenMD5
			bl.CostMsec = utils.Round(costMills, 6)
			bl.RequestTimestamp = start.Unix()
			bl.RequestTime = start.Format(constants.LayoutDateTimeISO8601ZoneP8MidMills)
			bl.ResponseTime = end.Format(constants.LayoutDateTimeISO8601ZoneP8MidMills)
			if bl.Response == nil {
				bl.Response = new(pb.Response)
				bl.Response.Message = uniResponseData.MSG
				bl.Response.Code = int32(uniResponseData.Code)
				bl.Response.RequestId = uniResponseData.RequestID
			}

			// busLogByte, err := proto.Marshal(bl)
			// if err != nil {
			// 	log4go.Error(err.Error())
			// }
			// kp := resources.GetProducer()
			// kp.SendPBMessage(busLogByte)

		} else {
			bl = nil
		}
	}
}

func recordOriginRequest(c *gin.Context, bl *pb.BusLog) {
	data, err := c.GetRawData()
	if err != nil {
		log4go.Error("record origin request data err: %v", err.Error())
		return
	}
	// copy rawData and reSet to request.Body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	r := c.Request

	var bodyMap map[string]interface{}
	err = json.Unmarshal(data, &bodyMap)
	requestLog := model.RequestLog{
		Host:       r.Host,
		Method:     r.Method,
		ClientIP:   c.ClientIP(),
		URL:        r.URL,
		Header:     r.Header,
		Body:       bodyMap,
		Form:       r.Form,
		PostForm:   r.PostForm,
		RemoteAddr: r.RemoteAddr,
	}
	c.Set(constants.ContextRequestOrigin, requestLog)

	if bl.Request == nil {
		bl.Request = new(pb.Request)
	}
	if bl.Request.Url == nil {
		bl.Request.Url = new(pb.RequestURL)
	}
	bl.Request.BodyOrigin = data
	bl.Request.Host = r.Host
	bl.Request.Method = r.Method
	bl.Request.ClientIp = c.ClientIP()
	bl.Request.RemoteUrl = r.RemoteAddr
	bl.Request.Url.Path = r.URL.Path
	bl.Request.Url.RawQuery = r.URL.RawQuery
	if len(r.Header) != 0 {
		if len(bl.Request.Header) == 0 {
			bl.Request.Header = make(map[string]*pb.ListOfString)
		}
		for headerK, headerV := range r.Header {
			bl.Request.Header[headerK] = &pb.ListOfString{Item: headerV}
		}
	}
	c.Set(constants.ContextBusinessLog, bl)
}
