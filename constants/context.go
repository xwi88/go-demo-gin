// Package constants context
package constants

type contextKey struct {
	name string
}

const (
	ContextBusinessLog        = "context_bus_log"
	ContextRequestOrigin      = "context_request_origin"       // 用于记录原始体请求信息
	ContextRequestOriginBody  = "context_request_origin_body"  // 用于记录原始体请求信息 http body
	ContextRequestOriginQuery = "context_request_origin_query" // 用于记录原始请求参数信息 query|form
	ContextRequestUserAgent   = "context_request_user_agent"   // 用于记录原始请求参数信息 user_agent
	ContextRequestCookies     = "context_request_cookies"      // 用于记录原始请求参数信息 cookies
	ContextResponseData       = "context_response_data"        // 响应信息
	ContextRequestRateLimit   = "context_request_rate_limit"   // 用于记录请求 rate limit 信息
	ContextRequestAppTokenMD5 = "context_app_token_md5"        // app_token encode with md5
	ContextRequestChannel     = "context_app_token_channel"    // app_token 对应渠道

	RequestID        = "request_id"
	RequestTimestamp = "timestamp"

	// 自定义header key, nginx 转发会忽略带 _ 的请求头，最好使用 - 替换 _
	HeaderKeyRequestID = "request-id"
	HeaderKeyTimestamp = "timestamp"

	// HTTP 标准 header key
	HTTPHeaderKeyContentType      = "Content-Type"
	HTTPHeaderValApplicationProto = "application/x-protobuf"
	HTTPHeaderValApplicationJSON  = "application/json"

	DefaultAPIVersion = "v1"
)

// RequestIDContextKey ...
var RequestIDContextKey = &contextKey{RequestID}
