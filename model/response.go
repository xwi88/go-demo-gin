// Package model define related struct
package model

// UniResponse response
type UniResponse struct {
	RequestID string      `json:"request_id,omitempty"` // 传入或系统生成的 request_id
	Code      int         `json:"code"`
	Data      interface{} `json:"data,omitempty"`
	MSG       string      `json:"msg"`
}
