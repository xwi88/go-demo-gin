// Package model define related struct
package model

// UniRequest request
type UniRequest struct {
	RequestID string `json:"request_id" form:"request_id"`
}

type SimpleRequest struct {
	UniRequest
	Name string `json:"name" form:"name"`
}
