// Code generated by protoc-gen-go. DO NOT EDIT.
// source: certificate.proto

package security

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

type CertificateType int32

const (
	CertificateType_Client CertificateType = 0
	CertificateType_Server CertificateType = 1
)

var CertificateType_name = map[int32]string{
	0: "Client",
	1: "Server",
}

var CertificateType_value = map[string]int32{
	"Client": 0,
	"Server": 1,
}

func (x CertificateType) String() string {
	return proto.EnumName(CertificateType_name, int32(x))
}

func (CertificateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{0}
}

type CertificateRequest struct {
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CertificateRequest) Reset()         { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()    {}
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{0}
}

func (m *CertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequest.Unmarshal(m, b)
}
func (m *CertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequest.Marshal(b, m, deterministic)
}
func (m *CertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequest.Merge(m, src)
}
func (m *CertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateRequest.Size(m)
}
func (m *CertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequest proto.InternalMessageInfo

func (m *CertificateRequest) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type CertificateResponse struct {
	Certificates         []*Certificate      `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{1}
}

func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateResponse.Unmarshal(m, b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
}
func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return xxx_messageInfo_CertificateResponse.Size(m)
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *CertificateResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CertificateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Certificate struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	NotBefore            int64           `protobuf:"varint,3,opt,name=notBefore,proto3" json:"notBefore,omitempty"`
	NotAfter             int64           `protobuf:"varint,4,opt,name=notAfter,proto3" json:"notAfter,omitempty"`
	Certificate          []byte          `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Status               *common.Status  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Type                 CertificateType `protobuf:"varint,7,opt,name=type,proto3,enum=moc.cloudagent.security.CertificateType" json:"type,omitempty"`
	GroupName            string          `protobuf:"bytes,8,opt,name=groupName,proto3" json:"groupName,omitempty"`
	VaultName            string          `protobuf:"bytes,9,opt,name=vaultName,proto3" json:"vaultName,omitempty"`
	LocationName         string          `protobuf:"bytes,10,opt,name=locationName,proto3" json:"locationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{2}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Certificate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Certificate) GetNotBefore() int64 {
	if m != nil {
		return m.NotBefore
	}
	return 0
}

func (m *Certificate) GetNotAfter() int64 {
	if m != nil {
		return m.NotAfter
	}
	return 0
}

func (m *Certificate) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *Certificate) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Certificate) GetType() CertificateType {
	if m != nil {
		return m.Type
	}
	return CertificateType_Client
}

func (m *Certificate) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *Certificate) GetVaultName() string {
	if m != nil {
		return m.VaultName
	}
	return ""
}

func (m *Certificate) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func init() {
	proto.RegisterEnum("moc.cloudagent.security.CertificateType", CertificateType_name, CertificateType_value)
	proto.RegisterType((*CertificateRequest)(nil), "moc.cloudagent.security.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "moc.cloudagent.security.CertificateResponse")
	proto.RegisterType((*Certificate)(nil), "moc.cloudagent.security.Certificate")
}

func init() { proto.RegisterFile("certificate.proto", fileDescriptor_c0d34c34dd33be4b) }

var fileDescriptor_c0d34c34dd33be4b = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x93, 0xd6, 0x34, 0x93, 0x28, 0x84, 0x05, 0x09, 0x2b, 0x42, 0xc8, 0x0a, 0x1c, 0xcc,
	0x87, 0x6c, 0x64, 0xae, 0x5c, 0x9a, 0x80, 0xe0, 0x04, 0x92, 0x0b, 0x1c, 0x38, 0x20, 0x39, 0x9b,
	0xb1, 0xb1, 0x64, 0x7b, 0xcc, 0xee, 0x6c, 0x51, 0x7e, 0x12, 0xff, 0x85, 0x03, 0x3f, 0x09, 0x79,
	0xdd, 0x36, 0x0e, 0x12, 0x52, 0x0e, 0xf4, 0xb6, 0xfb, 0xe6, 0xcd, 0x7b, 0x3b, 0xcf, 0x63, 0xb8,
	0x23, 0x51, 0x71, 0x91, 0x15, 0x32, 0x65, 0x0c, 0x1b, 0x45, 0x4c, 0xe2, 0x7e, 0x45, 0x32, 0x94,
	0x25, 0x99, 0x4d, 0x9a, 0x63, 0xcd, 0xa1, 0x46, 0x69, 0x54, 0xc1, 0xdb, 0xf9, 0xc3, 0x9c, 0x28,
	0x2f, 0x31, 0xb2, 0xb4, 0xb5, 0xc9, 0xa2, 0x1f, 0x2a, 0x6d, 0x1a, 0x54, 0xba, 0x6b, 0x9c, 0x4f,
	0x24, 0x55, 0x15, 0xd5, 0xdd, 0x6d, 0xf1, 0x15, 0xc4, 0x6a, 0xa7, 0x9d, 0xe0, 0x77, 0x83, 0x9a,
	0xc5, 0x3b, 0x98, 0xf4, 0x50, 0xed, 0x39, 0xfe, 0x30, 0x18, 0xc7, 0x8f, 0xc3, 0x7f, 0x78, 0x86,
	0x7d, 0x89, 0xbd, 0xce, 0xc5, 0x4f, 0x07, 0xee, 0xee, 0x19, 0xe8, 0x86, 0x6a, 0x8d, 0xff, 0xcf,
	0x41, 0xc4, 0xe0, 0x26, 0xa8, 0x4d, 0xc9, 0xde, 0xc0, 0x77, 0x82, 0x71, 0x3c, 0x0f, 0xbb, 0x00,
	0xc2, 0xab, 0x00, 0xc2, 0x25, 0x51, 0xf9, 0x39, 0x2d, 0x0d, 0x26, 0x97, 0x4c, 0x71, 0x0f, 0x4e,
	0xde, 0x28, 0x45, 0xca, 0x1b, 0xfa, 0x4e, 0x30, 0x4a, 0xba, 0xcb, 0xe2, 0xf7, 0x00, 0xc6, 0x3d,
	0x69, 0x21, 0xe0, 0xb8, 0x4e, 0x2b, 0xf4, 0x1c, 0x4b, 0xb2, 0x67, 0x31, 0x85, 0x41, 0xb1, 0xb1,
	0x4e, 0xa3, 0x64, 0x50, 0x6c, 0xc4, 0x03, 0x18, 0xd5, 0xc4, 0x4b, 0xcc, 0x48, 0xa1, 0x55, 0x1b,
	0x26, 0x3b, 0x40, 0xcc, 0xe1, 0xb4, 0x26, 0x3e, 0xcb, 0x18, 0x95, 0x77, 0x6c, 0x8b, 0xd7, 0x77,
	0xe1, 0xc3, 0xb8, 0xf7, 0x55, 0xbd, 0x13, 0xdf, 0x09, 0x26, 0x49, 0x1f, 0x12, 0x8f, 0xc0, 0xd5,
	0x9c, 0xb2, 0xd1, 0x9e, 0x6b, 0x27, 0x1b, 0xdb, 0x74, 0xce, 0x2d, 0x94, 0x5c, 0x96, 0xc4, 0x2b,
	0x38, 0xe6, 0x6d, 0x83, 0xde, 0x2d, 0xdf, 0x09, 0xa6, 0x71, 0x70, 0x48, 0x80, 0x1f, 0xb7, 0x0d,
	0x26, 0xb6, 0xab, 0x7d, 0x7e, 0xae, 0xc8, 0x34, 0xef, 0xdb, 0x39, 0x4f, 0xed, 0x54, 0x3b, 0xa0,
	0xad, 0x5e, 0xa4, 0xa6, 0x64, 0x5b, 0x1d, 0x75, 0xd5, 0x6b, 0x40, 0x2c, 0x60, 0x52, 0x92, 0x4c,
	0xb9, 0xa0, 0xda, 0x12, 0xc0, 0x12, 0xf6, 0xb0, 0xa7, 0x4f, 0xe0, 0xf6, 0x5f, 0xc6, 0x02, 0xc0,
	0x5d, 0x95, 0x05, 0xd6, 0x3c, 0x3b, 0x6a, 0xcf, 0xe7, 0xa8, 0x2e, 0x50, 0xcd, 0x9c, 0xf8, 0xd7,
	0x00, 0x66, 0x3d, 0xee, 0x59, 0xfb, 0x7c, 0x51, 0xc1, 0x74, 0xa5, 0x30, 0x65, 0xfc, 0xa0, 0x3e,
	0x35, 0x9b, 0x36, 0x94, 0x67, 0x07, 0xad, 0x48, 0xb7, 0xc7, 0xf3, 0xe7, 0x87, 0x91, 0xbb, 0x9d,
	0x5c, 0x1c, 0x89, 0x35, 0x0c, 0xdf, 0x22, 0xdf, 0xac, 0x07, 0x82, 0xfb, 0x1a, 0x4b, 0xbc, 0xe1,
	0x51, 0x96, 0xf1, 0x97, 0x17, 0x79, 0xc1, 0xdf, 0xcc, 0x3a, 0x94, 0x54, 0x45, 0x55, 0x21, 0x15,
	0x69, 0xca, 0x38, 0xaa, 0x48, 0x46, 0xaa, 0x91, 0xd1, 0x4e, 0x29, 0xba, 0x52, 0x5a, 0xbb, 0xf6,
	0x97, 0x79, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xcf, 0xbe, 0x93, 0x6f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CertificateAgentClient is the client API for CertificateAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateAgentClient interface {
	CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type certificateAgentClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAgentClient(cc *grpc.ClientConn) CertificateAgentClient {
	return &certificateAgentClient{cc}
}

func (c *certificateAgentClient) CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateAgentServer is the server API for CertificateAgent service.
type CertificateAgentServer interface {
	CreateOrUpdate(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Get(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Delete(context.Context, *CertificateRequest) (*CertificateResponse, error)
}

// UnimplementedCertificateAgentServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateAgentServer struct {
}

func (*UnimplementedCertificateAgentServer) CreateOrUpdate(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdate not implemented")
}
func (*UnimplementedCertificateAgentServer) Get(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCertificateAgentServer) Delete(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCertificateAgentServer(s *grpc.Server, srv CertificateAgentServer) {
	s.RegisterService(&_CertificateAgent_serviceDesc, srv)
}

func _CertificateAgent_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/CreateOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Get(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Delete(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.security.CertificateAgent",
	HandlerType: (*CertificateAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _CertificateAgent_CreateOrUpdate_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CertificateAgent_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CertificateAgent_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "certificate.proto",
}
