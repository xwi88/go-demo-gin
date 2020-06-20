// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: bus_log.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BusLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        string    `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	RequestTimestamp int64     `protobuf:"varint,2,opt,name=request_timestamp,json=requestTimestamp,proto3" json:"request_timestamp,omitempty"`
	CostMsec         float64   `protobuf:"fixed64,3,opt,name=cost_msec,json=costMsec,proto3" json:"cost_msec,omitempty"`
	HttpStatusCode   int32     `protobuf:"varint,4,opt,name=http_status_code,json=httpStatusCode,proto3" json:"http_status_code,omitempty"`
	Request          *Request  `protobuf:"bytes,5,opt,name=request,proto3" json:"request,omitempty"`
	Response         *Response `protobuf:"bytes,6,opt,name=response,proto3" json:"response,omitempty"`
	RequestTime      string    `protobuf:"bytes,7,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	ResponseTime     string    `protobuf:"bytes,8,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
}

func (x *BusLog) Reset() {
	*x = BusLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bus_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusLog) ProtoMessage() {}

func (x *BusLog) ProtoReflect() protoreflect.Message {
	mi := &file_bus_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusLog.ProtoReflect.Descriptor instead.
func (*BusLog) Descriptor() ([]byte, []int) {
	return file_bus_log_proto_rawDescGZIP(), []int{0}
}

func (x *BusLog) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *BusLog) GetRequestTimestamp() int64 {
	if x != nil {
		return x.RequestTimestamp
	}
	return 0
}

func (x *BusLog) GetCostMsec() float64 {
	if x != nil {
		return x.CostMsec
	}
	return 0
}

func (x *BusLog) GetHttpStatusCode() int32 {
	if x != nil {
		return x.HttpStatusCode
	}
	return 0
}

func (x *BusLog) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *BusLog) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *BusLog) GetRequestTime() string {
	if x != nil {
		return x.RequestTime
	}
	return ""
}

func (x *BusLog) GetResponseTime() string {
	if x != nil {
		return x.ResponseTime
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId  string                   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Host       string                   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	ClientIp   string                   `protobuf:"bytes,3,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	RemoteUrl  string                   `protobuf:"bytes,4,opt,name=remote_url,json=remoteUrl,proto3" json:"remote_url,omitempty"`
	Method     string                   `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	Header     map[string]*ListOfString `protobuf:"bytes,6,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Url        *RequestURL              `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Body       *RequestBody             `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	BodyOrigin []byte                   `protobuf:"bytes,9,opt,name=body_origin,json=bodyOrigin,proto3" json:"body_origin,omitempty"`
	TokenMd5   string                   `protobuf:"bytes,10,opt,name=token_md5,json=tokenMd5,proto3" json:"token_md5,omitempty"` // token_md5
	Channel    string                   `protobuf:"bytes,11,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bus_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_bus_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_bus_log_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *Request) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Request) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *Request) GetRemoteUrl() string {
	if x != nil {
		return x.RemoteUrl
	}
	return ""
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetHeader() map[string]*ListOfString {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Request) GetUrl() *RequestURL {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *Request) GetBody() *RequestBody {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Request) GetBodyOrigin() []byte {
	if x != nil {
		return x.BodyOrigin
	}
	return nil
}

func (x *Request) GetTokenMd5() string {
	if x != nil {
		return x.TokenMd5
	}
	return ""
}

func (x *Request) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type ListOfString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []string `protobuf:"bytes,1,rep,name=Item,proto3" json:"Item,omitempty"`
}

func (x *ListOfString) Reset() {
	*x = ListOfString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bus_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfString) ProtoMessage() {}

func (x *ListOfString) ProtoReflect() protoreflect.Message {
	mi := &file_bus_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfString.ProtoReflect.Descriptor instead.
func (*ListOfString) Descriptor() ([]byte, []int) {
	return file_bus_log_proto_rawDescGZIP(), []int{2}
}

func (x *ListOfString) GetItem() []string {
	if x != nil {
		return x.Item
	}
	return nil
}

type RequestURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	RawQuery string `protobuf:"bytes,2,opt,name=raw_query,json=rawQuery,proto3" json:"raw_query,omitempty"`
}

func (x *RequestURL) Reset() {
	*x = RequestURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bus_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestURL) ProtoMessage() {}

func (x *RequestURL) ProtoReflect() protoreflect.Message {
	mi := &file_bus_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestURL.ProtoReflect.Descriptor instead.
func (*RequestURL) Descriptor() ([]byte, []int) {
	return file_bus_log_proto_rawDescGZIP(), []int{3}
}

func (x *RequestURL) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RequestURL) GetRawQuery() string {
	if x != nil {
		return x.RawQuery
	}
	return ""
}

type RequestBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestBody) Reset() {
	*x = RequestBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bus_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBody) ProtoMessage() {}

func (x *RequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_bus_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBody.ProtoReflect.Descriptor instead.
func (*RequestBody) Descriptor() ([]byte, []int) {
	return file_bus_log_proto_rawDescGZIP(), []int{4}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Code      int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data      []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bus_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_bus_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_bus_log_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_bus_log_proto protoreflect.FileDescriptor

var file_bus_log_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x75, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x78, 0x77, 0x69, 0x38, 0x38, 0x2e, 0x6c, 0x6f, 0x67, 0x22, 0xc2, 0x02, 0x0a, 0x06, 0x42,
	0x75, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x73, 0x65, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x4d, 0x73, 0x65, 0x63, 0x12, 0x28,
	0x0a, 0x10, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x68, 0x74, 0x74, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x78, 0x77, 0x69, 0x38,
	0x38, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x78, 0x77, 0x69, 0x38, 0x38,
	0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xc9, 0x03, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x78, 0x77, 0x69, 0x38, 0x38, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x78, 0x77, 0x69, 0x38, 0x38, 0x2e,
	0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x78, 0x77, 0x69, 0x38, 0x38, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x64, 0x35, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x52, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x78, 0x77, 0x69, 0x38, 0x38, 0x2e,
	0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x22, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x49,
	0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x22,
	0x3d, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x0d,
	0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x6b, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bus_log_proto_rawDescOnce sync.Once
	file_bus_log_proto_rawDescData = file_bus_log_proto_rawDesc
)

func file_bus_log_proto_rawDescGZIP() []byte {
	file_bus_log_proto_rawDescOnce.Do(func() {
		file_bus_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_bus_log_proto_rawDescData)
	})
	return file_bus_log_proto_rawDescData
}

var file_bus_log_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bus_log_proto_goTypes = []interface{}{
	(*BusLog)(nil),       // 0: xwi88.log.BusLog
	(*Request)(nil),      // 1: xwi88.log.Request
	(*ListOfString)(nil), // 2: xwi88.log.ListOfString
	(*RequestURL)(nil),   // 3: xwi88.log.RequestURL
	(*RequestBody)(nil),  // 4: xwi88.log.RequestBody
	(*Response)(nil),     // 5: xwi88.log.Response
	nil,                  // 6: xwi88.log.Request.HeaderEntry
}
var file_bus_log_proto_depIdxs = []int32{
	1, // 0: xwi88.log.BusLog.request:type_name -> xwi88.log.Request
	5, // 1: xwi88.log.BusLog.response:type_name -> xwi88.log.Response
	6, // 2: xwi88.log.Request.header:type_name -> xwi88.log.Request.HeaderEntry
	3, // 3: xwi88.log.Request.url:type_name -> xwi88.log.RequestURL
	4, // 4: xwi88.log.Request.body:type_name -> xwi88.log.RequestBody
	2, // 5: xwi88.log.Request.HeaderEntry.value:type_name -> xwi88.log.ListOfString
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_bus_log_proto_init() }
func file_bus_log_proto_init() {
	if File_bus_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bus_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bus_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bus_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfString); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bus_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestURL); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bus_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBody); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bus_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bus_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bus_log_proto_goTypes,
		DependencyIndexes: file_bus_log_proto_depIdxs,
		MessageInfos:      file_bus_log_proto_msgTypes,
	}.Build()
	File_bus_log_proto = out.File
	file_bus_log_proto_rawDesc = nil
	file_bus_log_proto_goTypes = nil
	file_bus_log_proto_depIdxs = nil
}