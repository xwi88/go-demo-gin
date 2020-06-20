package constants

const (
	LayoutDateTimeISO8601ZoneP8Mid      = "2006-01-02T15:04:05+0800"
	LayoutDateTimeISO8601ZoneP8MidMills = "2006-01-02T15:04:05.000+0800"
	LayoutDateTimeISO8601ZoneP8MidMicro = "2006-01-02T15:04:05.000000+0800"

	TimestampFormatForLog = LayoutDateTimeISO8601ZoneP8MidMills

	// request_id 正则验证
	RegexRequestID = "^[a-zA-Z0-9]{32}$"
)
