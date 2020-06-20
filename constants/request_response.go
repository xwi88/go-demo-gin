package constants

import (
	"net/http"
)

const (
	RequestMethodNotAllowed = http.StatusMethodNotAllowed
)

const (
	ResponseCodeOK = 0
)

// ResponseMSG response code message map
var ResponseMSG = map[int]string{
	ResponseCodeOK: "成功",
}
