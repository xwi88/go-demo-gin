// Package model define related struct
package model

import (
	"net/http"
	"net/url"
	"time"
)

// RequestLog request detail log
type RequestLog struct {
	Host     string `json:"host"`
	ClientIP string `json:"client_ip"` // 调用方ip
	Method   string `json:"method"`    // 请求方法
	// include User-Agent, Referer, Cookie(所有Cookie) and so on
	Header     http.Header `json:"header"`
	URL        *url.URL    `json:"url"`
	Body       interface{} `json:"body,omitempty"` //  请求数据 body
	Form       url.Values  `json:"form,omitempty"`
	PostForm   url.Values  `json:"post_form,omitempty"`
	RemoteAddr string      `json:"remote_addr,omitempty"`
}

type RequestAPIData struct {
	HTTPStatus       int                    `json:"http_status"` // http status code
	CostMills        float64                `json:"cost_mills"`  // 接口耗时 ms
	RequestID        string                 `json:"request_id"`  // 请求标识
	RequestTime      time.Time              `json:"request_time"`
	ResponseTime     time.Time              `json:"response_time"`
	RequestTimestamp int64                  `json:"request_timestamp"` // 请求时间 s
	Request          RequestLog             `json:"request"`
	Response         interface{}            `json:"response"` // 返回数据
	ExtraMSG         map[string]interface{} `json:"extra_msg,omitempty"`
}
