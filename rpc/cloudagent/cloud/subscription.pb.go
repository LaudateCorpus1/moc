// Code generated by protoc-gen-go. DO NOT EDIT.
// source: subscription.proto

package cloud

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc-proto/rpc/common"
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

type SubscriptionRequest struct {
	Subscriptions        []*Subscription  `protobuf:"bytes,1,rep,name=Subscriptions,proto3" json:"Subscriptions,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SubscriptionRequest) Reset()         { *m = SubscriptionRequest{} }
func (m *SubscriptionRequest) String() string { return proto.CompactTextString(m) }
func (*SubscriptionRequest) ProtoMessage()    {}
func (*SubscriptionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{0}
}

func (m *SubscriptionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionRequest.Unmarshal(m, b)
}
func (m *SubscriptionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionRequest.Marshal(b, m, deterministic)
}
func (m *SubscriptionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionRequest.Merge(m, src)
}
func (m *SubscriptionRequest) XXX_Size() int {
	return xxx_messageInfo_SubscriptionRequest.Size(m)
}
func (m *SubscriptionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionRequest proto.InternalMessageInfo

func (m *SubscriptionRequest) GetSubscriptions() []*Subscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *SubscriptionRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type SubscriptionResponse struct {
	Subscriptions        []*Subscription     `protobuf:"bytes,1,rep,name=Subscriptions,proto3" json:"Subscriptions,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SubscriptionResponse) Reset()         { *m = SubscriptionResponse{} }
func (m *SubscriptionResponse) String() string { return proto.CompactTextString(m) }
func (*SubscriptionResponse) ProtoMessage()    {}
func (*SubscriptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{1}
}

func (m *SubscriptionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionResponse.Unmarshal(m, b)
}
func (m *SubscriptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionResponse.Marshal(b, m, deterministic)
}
func (m *SubscriptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionResponse.Merge(m, src)
}
func (m *SubscriptionResponse) XXX_Size() int {
	return xxx_messageInfo_SubscriptionResponse.Size(m)
}
func (m *SubscriptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionResponse proto.InternalMessageInfo

func (m *SubscriptionResponse) GetSubscriptions() []*Subscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *SubscriptionResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *SubscriptionResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Subscription struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Status               *common.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f8ad1a64b2bad6, []int{2}
}

func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subscription) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Subscription) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscriptionRequest)(nil), "moc.cloudagent.subscription.SubscriptionRequest")
	proto.RegisterType((*SubscriptionResponse)(nil), "moc.cloudagent.subscription.SubscriptionResponse")
	proto.RegisterType((*Subscription)(nil), "moc.cloudagent.subscription.Subscription")
}

func init() { proto.RegisterFile("subscription.proto", fileDescriptor_c4f8ad1a64b2bad6) }

var fileDescriptor_c4f8ad1a64b2bad6 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xcf, 0x4e, 0xe3, 0x30,
	0x10, 0xc6, 0x37, 0xed, 0x6e, 0xa4, 0x3a, 0x6d, 0xa5, 0xf5, 0xf6, 0x10, 0x65, 0x25, 0x54, 0x85,
	0x4b, 0x39, 0xe0, 0x40, 0xe0, 0x05, 0xa8, 0xc4, 0x81, 0x53, 0x25, 0x17, 0x81, 0xc4, 0x2d, 0x71,
	0xdd, 0x10, 0x11, 0x67, 0x8c, 0xff, 0x80, 0xb8, 0xf3, 0x18, 0xbc, 0x04, 0x6f, 0x88, 0xea, 0xb4,
	0x90, 0x5c, 0x10, 0x48, 0xdc, 0x6c, 0xcf, 0xf7, 0xfd, 0x66, 0xe6, 0x33, 0xc2, 0xda, 0xe6, 0x9a,
	0xa9, 0x52, 0x9a, 0x12, 0x6a, 0x22, 0x15, 0x18, 0xc0, 0xff, 0x05, 0x30, 0xc2, 0x2a, 0xb0, 0xab,
	0xac, 0xe0, 0xb5, 0x21, 0x6d, 0x49, 0xb4, 0x57, 0x00, 0x14, 0x15, 0x4f, 0x9c, 0x34, 0xb7, 0xeb,
	0xe4, 0x51, 0x65, 0x52, 0x72, 0xa5, 0x1b, 0x73, 0x34, 0x64, 0x20, 0xc4, 0x0e, 0x15, 0xbf, 0x78,
	0xe8, 0xdf, 0xb2, 0x65, 0xa7, 0xfc, 0xde, 0x72, 0x6d, 0xf0, 0x02, 0x8d, 0xda, 0xcf, 0x3a, 0xf4,
	0xa6, 0xfd, 0x59, 0x90, 0x1e, 0x90, 0x4f, 0x5a, 0x93, 0x0e, 0xa8, 0xeb, 0xc7, 0xa7, 0x68, 0xb4,
	0x90, 0x5c, 0x65, 0x9b, 0xdb, 0xe5, 0x93, 0xe4, 0x61, 0x6f, 0xea, 0xcd, 0xc6, 0xe9, 0xd8, 0x01,
	0xdf, 0x2b, 0xb4, 0x2b, 0x8a, 0x5f, 0x3d, 0x34, 0xe9, 0x8e, 0xa7, 0x25, 0xd4, 0x9a, 0xff, 0xfc,
	0x7c, 0x29, 0xf2, 0x29, 0xd7, 0xb6, 0x32, 0x6e, 0xb0, 0x20, 0x8d, 0x48, 0x93, 0x23, 0xd9, 0xe5,
	0x48, 0xe6, 0x00, 0xd5, 0x55, 0x56, 0x59, 0x4e, 0xb7, 0x4a, 0x3c, 0x41, 0x7f, 0xce, 0x95, 0x02,
	0x15, 0xf6, 0xa7, 0xde, 0x6c, 0x40, 0x9b, 0x4b, 0x7c, 0x8d, 0x86, 0x6d, 0x34, 0xc6, 0xe8, 0x77,
	0x9d, 0x09, 0x1e, 0x7a, 0x4e, 0xe4, 0xce, 0x78, 0x8c, 0x7a, 0xe5, 0xca, 0x75, 0x1a, 0xd0, 0x5e,
	0xb9, 0xc2, 0xfb, 0xc8, 0xd7, 0x26, 0x33, 0x56, 0x3b, 0x54, 0x90, 0x06, 0x6e, 0x8f, 0xa5, 0x7b,
	0xa2, 0xdb, 0x52, 0xfa, 0xec, 0xa1, 0xbf, 0x6d, 0xf2, 0xd9, 0x66, 0x43, 0x0c, 0xc8, 0xbf, 0xa8,
	0x1f, 0xe0, 0x8e, 0xe3, 0xa3, 0xaf, 0x2f, 0xdf, 0xfc, 0x72, 0x74, 0xfc, 0x0d, 0x47, 0x13, 0x7c,
	0xfc, 0x6b, 0x9e, 0xdc, 0x1c, 0x16, 0xa5, 0xb9, 0xb5, 0x39, 0x61, 0x20, 0x12, 0x51, 0x32, 0x05,
	0x1a, 0xd6, 0x26, 0x11, 0xc0, 0x12, 0x25, 0x59, 0xf2, 0x81, 0x6b, 0x8e, 0xb9, 0xef, 0x22, 0x3c,
	0x79, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xb5, 0xe6, 0xfc, 0xcb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SubscriptionAgentClient is the client API for SubscriptionAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscriptionAgentClient interface {
	Invoke(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error)
}

type subscriptionAgentClient struct {
	cc *grpc.ClientConn
}

func NewSubscriptionAgentClient(cc *grpc.ClientConn) SubscriptionAgentClient {
	return &subscriptionAgentClient{cc}
}

func (c *subscriptionAgentClient) Invoke(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error) {
	out := new(SubscriptionResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.subscription.SubscriptionAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionAgentServer is the server API for SubscriptionAgent service.
type SubscriptionAgentServer interface {
	Invoke(context.Context, *SubscriptionRequest) (*SubscriptionResponse, error)
}

// UnimplementedSubscriptionAgentServer can be embedded to have forward compatible implementations.
type UnimplementedSubscriptionAgentServer struct {
}

func (*UnimplementedSubscriptionAgentServer) Invoke(ctx context.Context, req *SubscriptionRequest) (*SubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterSubscriptionAgentServer(s *grpc.Server, srv SubscriptionAgentServer) {
	s.RegisterService(&_SubscriptionAgent_serviceDesc, srv)
}

func _SubscriptionAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.subscription.SubscriptionAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionAgentServer).Invoke(ctx, req.(*SubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubscriptionAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.subscription.SubscriptionAgent",
	HandlerType: (*SubscriptionAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _SubscriptionAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription.proto",
}
