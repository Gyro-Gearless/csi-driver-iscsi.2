// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google.golang.org/appengine/internal/remote_api/remote_api.proto

package remote_api

import (
	proto "github.com/golang/protobuf/proto"
	fmt "fmt"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

var (
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RpcError_ErrorCode int32

const (
	RpcError_UNKNOWN             RpcError_ErrorCode = 0
	RpcError_CALL_NOT_FOUND      RpcError_ErrorCode = 1
	RpcError_PARSE_ERROR         RpcError_ErrorCode = 2
	RpcError_SECURITY_VIOLATION  RpcError_ErrorCode = 3
	RpcError_OVER_QUOTA          RpcError_ErrorCode = 4
	RpcError_REQUEST_TOO_LARGE   RpcError_ErrorCode = 5
	RpcError_CAPABILITY_DISABLED RpcError_ErrorCode = 6
	RpcError_FEATURE_DISABLED    RpcError_ErrorCode = 7
	RpcError_BAD_REQUEST         RpcError_ErrorCode = 8
	RpcError_RESPONSE_TOO_LARGE  RpcError_ErrorCode = 9
	RpcError_CANCELLED           RpcError_ErrorCode = 10
	RpcError_REPLAY_ERROR        RpcError_ErrorCode = 11
	RpcError_DEADLINE_EXCEEDED   RpcError_ErrorCode = 12
)

var RpcError_ErrorCode_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CALL_NOT_FOUND",
	2:  "PARSE_ERROR",
	3:  "SECURITY_VIOLATION",
	4:  "OVER_QUOTA",
	5:  "REQUEST_TOO_LARGE",
	6:  "CAPABILITY_DISABLED",
	7:  "FEATURE_DISABLED",
	8:  "BAD_REQUEST",
	9:  "RESPONSE_TOO_LARGE",
	10: "CANCELLED",
	11: "REPLAY_ERROR",
	12: "DEADLINE_EXCEEDED",
}

var RpcError_ErrorCode_value = map[string]int32{
	"UNKNOWN":             0,
	"CALL_NOT_FOUND":      1,
	"PARSE_ERROR":         2,
	"SECURITY_VIOLATION":  3,
	"OVER_QUOTA":          4,
	"REQUEST_TOO_LARGE":   5,
	"CAPABILITY_DISABLED": 6,
	"FEATURE_DISABLED":    7,
	"BAD_REQUEST":         8,
	"RESPONSE_TOO_LARGE":  9,
	"CANCELLED":           10,
	"REPLAY_ERROR":        11,
	"DEADLINE_EXCEEDED":   12,
}

func (x RpcError_ErrorCode) Enum() *RpcError_ErrorCode {
	p := new(RpcError_ErrorCode)
	*p = x
	return p
}

func (x RpcError_ErrorCode) String() string {
	return proto.EnumName(RpcError_ErrorCode_name, int32(x))
}

func (x *RpcError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcError_ErrorCode_value, data, "RpcError_ErrorCode")
	if err != nil {
		return err
	}
	*x = RpcError_ErrorCode(value)
	return nil
}

func (RpcError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_remote_api_1978114ec33a273d, []int{2, 0}
}

type Request struct {
	ServiceName          *string  `protobuf:"bytes,2,req,name=service_name,json=serviceName" json:"service_name,omitempty"`
	Method               *string  `protobuf:"bytes,3,req,name=method" json:"method,omitempty"`
	Request              []byte   `protobuf:"bytes,4,req,name=request" json:"request,omitempty"`
	RequestId            *string  `protobuf:"bytes,5,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_api_1978114ec33a273d, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}

func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}

func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}

func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}

func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Request) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *Request) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Request) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

type ApplicationError struct {
	Code                 *int32   `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Detail               *string  `protobuf:"bytes,2,req,name=detail" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplicationError) Reset()         { *m = ApplicationError{} }
func (m *ApplicationError) String() string { return proto.CompactTextString(m) }
func (*ApplicationError) ProtoMessage()    {}
func (*ApplicationError) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_api_1978114ec33a273d, []int{1}
}

func (m *ApplicationError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationError.Unmarshal(m, b)
}

func (m *ApplicationError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationError.Marshal(b, m, deterministic)
}

func (dst *ApplicationError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationError.Merge(dst, src)
}

func (m *ApplicationError) XXX_Size() int {
	return xxx_messageInfo_ApplicationError.Size(m)
}

func (m *ApplicationError) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationError.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationError proto.InternalMessageInfo

func (m *ApplicationError) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *ApplicationError) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

type RpcError struct {
	Code                 *int32   `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Detail               *string  `protobuf:"bytes,2,opt,name=detail" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcError) Reset()         { *m = RpcError{} }
func (m *RpcError) String() string { return proto.CompactTextString(m) }
func (*RpcError) ProtoMessage()    {}
func (*RpcError) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_api_1978114ec33a273d, []int{2}
}

func (m *RpcError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcError.Unmarshal(m, b)
}

func (m *RpcError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcError.Marshal(b, m, deterministic)
}

func (dst *RpcError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcError.Merge(dst, src)
}

func (m *RpcError) XXX_Size() int {
	return xxx_messageInfo_RpcError.Size(m)
}

func (m *RpcError) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcError.DiscardUnknown(m)
}

var xxx_messageInfo_RpcError proto.InternalMessageInfo

func (m *RpcError) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *RpcError) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

type Response struct {
	Response             []byte            `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Exception            []byte            `protobuf:"bytes,2,opt,name=exception" json:"exception,omitempty"`
	ApplicationError     *ApplicationError `protobuf:"bytes,3,opt,name=application_error,json=applicationError" json:"application_error,omitempty"`
	JavaException        []byte            `protobuf:"bytes,4,opt,name=java_exception,json=javaException" json:"java_exception,omitempty"`
	RpcError             *RpcError         `protobuf:"bytes,5,opt,name=rpc_error,json=rpcError" json:"rpc_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_api_1978114ec33a273d, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}

func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}

func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}

func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}

func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Response) GetException() []byte {
	if m != nil {
		return m.Exception
	}
	return nil
}

func (m *Response) GetApplicationError() *ApplicationError {
	if m != nil {
		return m.ApplicationError
	}
	return nil
}

func (m *Response) GetJavaException() []byte {
	if m != nil {
		return m.JavaException
	}
	return nil
}

func (m *Response) GetRpcError() *RpcError {
	if m != nil {
		return m.RpcError
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "remote_api.Request")
	proto.RegisterType((*ApplicationError)(nil), "remote_api.ApplicationError")
	proto.RegisterType((*RpcError)(nil), "remote_api.RpcError")
	proto.RegisterType((*Response)(nil), "remote_api.Response")
}

func init() {
	proto.RegisterFile("google.golang.org/appengine/internal/remote_api/remote_api.proto", fileDescriptor_remote_api_1978114ec33a273d)
}

var fileDescriptor_remote_api_1978114ec33a273d = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x51, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb1, 0x9b, 0x34, 0xf1, 0xc4, 0x2d, 0xdb, 0xa5, 0x14, 0x0b, 0x15, 0x29, 0x44, 0x42,
	0xca, 0x53, 0x2a, 0x38, 0x00, 0x62, 0x63, 0x6f, 0x91, 0x85, 0x65, 0xa7, 0x6b, 0xbb, 0x50, 0x5e,
	0x56, 0x2b, 0x67, 0x65, 0x8c, 0x12, 0xaf, 0xd9, 0x98, 0x8a, 0x17, 0x6e, 0xc0, 0xb5, 0x38, 0x0c,
	0xb7, 0x40, 0x36, 0x6e, 0x63, 0xf5, 0x89, 0xb7, 0x7f, 0x7e, 0x7b, 0xe6, 0x1b, 0xcd, 0xcc, 0xc2,
	0xbb, 0x5c, 0xa9, 0x7c, 0x23, 0x17, 0xb9, 0xda, 0x88, 0x32, 0x5f, 0x28, 0x9d, 0x5f, 0x88, 0xaa,
	0x92, 0x65, 0x5e, 0x94, 0xf2, 0xa2, 0x28, 0x6b, 0xa9, 0x4b, 0xb1, 0xb9, 0xd0, 0x72, 0xab, 0x6a,
	0xc9, 0x45, 0x55, 0xf4, 0xe4, 0xa2, 0xd2, 0xaa, 0x56, 0x18, 0xf6, 0xce, 0xec, 0x27, 0x8c, 0x98,
	0xfc, 0xf6, 0x5d, 0xee, 0x6a, 0xfc, 0x12, 0xec, 0x9d, 0xd4, 0xb7, 0x45, 0x26, 0x79, 0x29, 0xb6,
	0xd2, 0x31, 0xa7, 0xe6, 0xdc, 0x62, 0x93, 0xce, 0x0b, 0xc5, 0x56, 0xe2, 0x33, 0x38, 0xdc, 0xca,
	0xfa, 0x8b, 0x5a, 0x3b, 0x07, 0xed, 0xc7, 0x2e, 0xc2, 0x0e, 0x8c, 0xf4, 0xbf, 0x2a, 0xce, 0x60,
	0x6a, 0xce, 0x6d, 0x76, 0x17, 0xe2, 0x17, 0x00, 0x9d, 0xe4, 0xc5, 0xda, 0x19, 0x4e, 0x8d, 0xb9,
	0xc5, 0xac, 0xce, 0xf1, 0xd7, 0xb3, 0xb7, 0x80, 0x48, 0x55, 0x6d, 0x8a, 0x4c, 0xd4, 0x85, 0x2a,
	0xa9, 0xd6, 0x4a, 0x63, 0x0c, 0x83, 0x4c, 0xad, 0xa5, 0x63, 0x4c, 0xcd, 0xf9, 0x90, 0xb5, 0xba,
	0x01, 0xaf, 0x65, 0x2d, 0x8a, 0x4d, 0xd7, 0x55, 0x17, 0xcd, 0x7e, 0x9b, 0x30, 0x66, 0x55, 0xf6,
	0x7f, 0x89, 0x46, 0x2f, 0xf1, 0x97, 0x09, 0x56, 0x9b, 0xe5, 0x36, 0x7f, 0x4d, 0x60, 0x94, 0x86,
	0x1f, 0xc2, 0xe8, 0x63, 0x88, 0x1e, 0x61, 0x0c, 0xc7, 0x2e, 0x09, 0x02, 0x1e, 0x46, 0x09, 0xbf,
	0x8c, 0xd2, 0xd0, 0x43, 0x06, 0x7e, 0x0c, 0x93, 0x15, 0x61, 0x31, 0xe5, 0x94, 0xb1, 0x88, 0x21,
	0x13, 0x9f, 0x01, 0x8e, 0xa9, 0x9b, 0x32, 0x3f, 0xb9, 0xe1, 0xd7, 0x7e, 0x14, 0x90, 0xc4, 0x8f,
	0x42, 0x74, 0x80, 0x8f, 0x01, 0xa2, 0x6b, 0xca, 0xf8, 0x55, 0x1a, 0x25, 0x04, 0x0d, 0xf0, 0x53,
	0x38, 0x61, 0xf4, 0x2a, 0xa5, 0x71, 0xc2, 0x93, 0x28, 0xe2, 0x01, 0x61, 0xef, 0x29, 0x1a, 0xe2,
	0x67, 0xf0, 0xc4, 0x25, 0x2b, 0xb2, 0xf4, 0x83, 0xa6, 0x80, 0xe7, 0xc7, 0x64, 0x19, 0x50, 0x0f,
	0x1d, 0xe2, 0x53, 0x40, 0x97, 0x94, 0x24, 0x29, 0xa3, 0x7b, 0x77, 0xd4, 0xe0, 0x97, 0xc4, 0xe3,
	0x5d, 0x25, 0x34, 0x6e, 0xf0, 0x8c, 0xc6, 0xab, 0x28, 0x8c, 0x69, 0xaf, 0xae, 0x85, 0x8f, 0xc0,
	0x72, 0x49, 0xe8, 0xd2, 0xa0, 0xc9, 0x03, 0x8c, 0xc0, 0x66, 0x74, 0x15, 0x90, 0x9b, 0xae, 0xef,
	0x49, 0xd3, 0x8f, 0x47, 0x89, 0x17, 0xf8, 0x21, 0xe5, 0xf4, 0x93, 0x4b, 0xa9, 0x47, 0x3d, 0x64,
	0xcf, 0xfe, 0x18, 0x30, 0x66, 0x72, 0x57, 0xa9, 0x72, 0x27, 0xf1, 0x73, 0x18, 0xeb, 0x4e, 0x3b,
	0xc6, 0xd4, 0x98, 0xdb, 0xec, 0x3e, 0xc6, 0xe7, 0x60, 0xc9, 0x1f, 0x99, 0xac, 0x9a, 0x75, 0xb5,
	0x23, 0xb5, 0xd9, 0xde, 0xc0, 0x3e, 0x9c, 0x88, 0xfd, 0x3a, 0xb9, 0x6c, 0x06, 0xec, 0x1c, 0x4c,
	0x8d, 0xf9, 0xe4, 0xcd, 0xf9, 0xa2, 0x77, 0x87, 0x0f, 0x77, 0xce, 0x90, 0x78, 0x78, 0x05, 0xaf,
	0xe0, 0xf8, 0xab, 0xb8, 0x15, 0x7c, 0x4f, 0x1b, 0xb4, 0xb4, 0xa3, 0xc6, 0xa5, 0xf7, 0xc4, 0xd7,
	0x60, 0xe9, 0x2a, 0xeb, 0x48, 0xc3, 0x96, 0x74, 0xda, 0x27, 0xdd, 0x1d, 0x07, 0x1b, 0xeb, 0x4e,
	0x2d, 0xed, 0xcf, 0xbd, 0x07, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0x38, 0xd1, 0x0f, 0x22, 0x4f,
	0x03, 0x00, 0x00,
}
