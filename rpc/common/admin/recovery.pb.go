// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin/recovery/recovery.proto

package admin

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Operation int32

const (
	Operation_BACKUP  Operation = 0
	Operation_RESTORE Operation = 1
)

var Operation_name = map[int32]string{
	0: "BACKUP",
	1: "RESTORE",
}

var Operation_value = map[string]int32{
	"BACKUP":  0,
	"RESTORE": 1,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}

func (Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6c726908b67e934, []int{0}
}

type RecoveryRequest struct {
	OperationType Operation `protobuf:"varint,1,opt,name=OperationType,proto3,enum=moc.common.admin.Operation" json:"OperationType,omitempty"`
	// Path to back to or restore from
	Path                 string   `protobuf:"bytes,2,opt,name=Path,proto3" json:"Path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecoveryRequest) Reset()         { *m = RecoveryRequest{} }
func (m *RecoveryRequest) String() string { return proto.CompactTextString(m) }
func (*RecoveryRequest) ProtoMessage()    {}
func (*RecoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6c726908b67e934, []int{0}
}

func (m *RecoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecoveryRequest.Unmarshal(m, b)
}
func (m *RecoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecoveryRequest.Marshal(b, m, deterministic)
}
func (m *RecoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoveryRequest.Merge(m, src)
}
func (m *RecoveryRequest) XXX_Size() int {
	return xxx_messageInfo_RecoveryRequest.Size(m)
}
func (m *RecoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecoveryRequest proto.InternalMessageInfo

func (m *RecoveryRequest) GetOperationType() Operation {
	if m != nil {
		return m.OperationType
	}
	return Operation_BACKUP
}

func (m *RecoveryRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type RecoveryResponse struct {
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RecoveryResponse) Reset()         { *m = RecoveryResponse{} }
func (m *RecoveryResponse) String() string { return proto.CompactTextString(m) }
func (*RecoveryResponse) ProtoMessage()    {}
func (*RecoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6c726908b67e934, []int{1}
}

func (m *RecoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecoveryResponse.Unmarshal(m, b)
}
func (m *RecoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecoveryResponse.Marshal(b, m, deterministic)
}
func (m *RecoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoveryResponse.Merge(m, src)
}
func (m *RecoveryResponse) XXX_Size() int {
	return xxx_messageInfo_RecoveryResponse.Size(m)
}
func (m *RecoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecoveryResponse proto.InternalMessageInfo

func (m *RecoveryResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *RecoveryResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("moc.common.admin.Operation", Operation_name, Operation_value)
	proto.RegisterType((*RecoveryRequest)(nil), "moc.common.admin.RecoveryRequest")
	proto.RegisterType((*RecoveryResponse)(nil), "moc.common.admin.RecoveryResponse")
}

func init() { proto.RegisterFile("admin/recovery/recovery.proto", fileDescriptor_d6c726908b67e934) }

var fileDescriptor_d6c726908b67e934 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x5f, 0x4f, 0xc2, 0x30,
	0x14, 0xc5, 0x99, 0x7f, 0x66, 0xb8, 0x04, 0x5d, 0x1a, 0x1f, 0x08, 0x46, 0x83, 0x8b, 0x0f, 0xa8,
	0x49, 0x9b, 0xcc, 0x4f, 0x00, 0x86, 0x07, 0xe3, 0x03, 0x58, 0xd1, 0x07, 0xe3, 0xcb, 0x36, 0x2f,
	0x63, 0x71, 0xdb, 0xad, 0x5d, 0x87, 0xe1, 0xdb, 0x1b, 0x3b, 0xc0, 0x88, 0x89, 0x6f, 0xa7, 0xe9,
	0xef, 0x9c, 0x9e, 0x53, 0x38, 0x0d, 0xdf, 0xf2, 0xb4, 0x10, 0x1a, 0x63, 0x5a, 0xa0, 0x5e, 0x6e,
	0x04, 0x57, 0x9a, 0x0c, 0x31, 0x2f, 0xa7, 0x98, 0xc7, 0x94, 0xe7, 0x54, 0x70, 0x4b, 0x76, 0xcf,
	0x12, 0xa2, 0x24, 0x43, 0x61, 0xef, 0xa3, 0x6a, 0x26, 0x3e, 0x75, 0xa8, 0x14, 0xea, 0xb2, 0x76,
	0xf8, 0x73, 0x38, 0x92, 0xab, 0x0c, 0x89, 0x1f, 0x15, 0x96, 0x86, 0x0d, 0xa0, 0x3d, 0x56, 0xa8,
	0x43, 0x93, 0x52, 0x31, 0x5d, 0x2a, 0xec, 0x38, 0x3d, 0xa7, 0x7f, 0x18, 0x9c, 0xf0, 0xed, 0x70,
	0xbe, 0xc1, 0xe4, 0x6f, 0x07, 0x63, 0xb0, 0x37, 0x09, 0xcd, 0xbc, 0xb3, 0xd3, 0x73, 0xfa, 0x4d,
	0x69, 0xb5, 0xff, 0x0a, 0xde, 0xcf, 0x4b, 0xa5, 0xa2, 0xa2, 0x44, 0x16, 0x80, 0x2b, 0xb1, 0xac,
	0x32, 0x63, 0xc9, 0x56, 0xd0, 0xe5, 0x75, 0x5d, 0xbe, 0xae, 0xcb, 0x87, 0x44, 0xd9, 0x73, 0x98,
	0x55, 0x28, 0x57, 0x24, 0x3b, 0x86, 0xfd, 0x91, 0xd6, 0xa4, 0x3b, 0xbb, 0x36, 0xbc, 0x3e, 0x5c,
	0x5d, 0x40, 0x73, 0x53, 0x81, 0x01, 0xb8, 0xc3, 0xc1, 0xed, 0xfd, 0xd3, 0xc4, 0x6b, 0xb0, 0x16,
	0x1c, 0xc8, 0xd1, 0xe3, 0x74, 0x2c, 0x47, 0x9e, 0x13, 0x44, 0xd0, 0x5e, 0x77, 0x18, 0x24, 0x58,
	0x18, 0xf6, 0x00, 0xee, 0x5d, 0xb1, 0xa0, 0x77, 0x64, 0xe7, 0x7f, 0xe7, 0x6d, 0x7d, 0x4c, 0xd7,
	0xff, 0x0f, 0xa9, 0x17, 0xf9, 0x8d, 0xe1, 0xf5, 0xcb, 0x65, 0x92, 0x9a, 0x79, 0x15, 0x7d, 0x93,
	0x22, 0x4f, 0x63, 0x4d, 0x25, 0xcd, 0x8c, 0xc8, 0x29, 0x16, 0x5a, 0xc5, 0xa2, 0xf6, 0x0b, 0xeb,
	0x8f, 0x5c, 0x3b, 0xf4, 0xe6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xac, 0xe1, 0xf5, 0xd8, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecoveryAgentClient is the client API for RecoveryAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecoveryAgentClient interface {
	Invoke(ctx context.Context, in *RecoveryRequest, opts ...grpc.CallOption) (*RecoveryResponse, error)
}

type recoveryAgentClient struct {
	cc *grpc.ClientConn
}

func NewRecoveryAgentClient(cc *grpc.ClientConn) RecoveryAgentClient {
	return &recoveryAgentClient{cc}
}

func (c *recoveryAgentClient) Invoke(ctx context.Context, in *RecoveryRequest, opts ...grpc.CallOption) (*RecoveryResponse, error) {
	out := new(RecoveryResponse)
	err := c.cc.Invoke(ctx, "/moc.common.admin.RecoveryAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecoveryAgentServer is the server API for RecoveryAgent service.
type RecoveryAgentServer interface {
	Invoke(context.Context, *RecoveryRequest) (*RecoveryResponse, error)
}

// UnimplementedRecoveryAgentServer can be embedded to have forward compatible implementations.
type UnimplementedRecoveryAgentServer struct {
}

func (*UnimplementedRecoveryAgentServer) Invoke(ctx context.Context, req *RecoveryRequest) (*RecoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterRecoveryAgentServer(s *grpc.Server, srv RecoveryAgentServer) {
	s.RegisterService(&_RecoveryAgent_serviceDesc, srv)
}

func _RecoveryAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecoveryAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.common.admin.RecoveryAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecoveryAgentServer).Invoke(ctx, req.(*RecoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecoveryAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.common.admin.RecoveryAgent",
	HandlerType: (*RecoveryAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _RecoveryAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/recovery/recovery.proto",
}
