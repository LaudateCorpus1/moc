// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_baremetalhostagent.proto

package baremetalhostagent

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type BareMetalHostSSHPublicKey struct {
	Keydata              string   `protobuf:"bytes,1,opt,name=keydata,proto3" json:"keydata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BareMetalHostSSHPublicKey) Reset()         { *m = BareMetalHostSSHPublicKey{} }
func (m *BareMetalHostSSHPublicKey) String() string { return proto.CompactTextString(m) }
func (*BareMetalHostSSHPublicKey) ProtoMessage()    {}
func (*BareMetalHostSSHPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{0}
}

func (m *BareMetalHostSSHPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostSSHPublicKey.Unmarshal(m, b)
}
func (m *BareMetalHostSSHPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostSSHPublicKey.Marshal(b, m, deterministic)
}
func (m *BareMetalHostSSHPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostSSHPublicKey.Merge(m, src)
}
func (m *BareMetalHostSSHPublicKey) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostSSHPublicKey.Size(m)
}
func (m *BareMetalHostSSHPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostSSHPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostSSHPublicKey proto.InternalMessageInfo

func (m *BareMetalHostSSHPublicKey) GetKeydata() string {
	if m != nil {
		return m.Keydata
	}
	return ""
}

type BareMetalHostDisk struct {
	DiskName             string   `protobuf:"bytes,1,opt,name=diskName,proto3" json:"diskName,omitempty"`
	DiskSizeGB           int32    `protobuf:"varint,2,opt,name=diskSizeGB,proto3" json:"diskSizeGB,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BareMetalHostDisk) Reset()         { *m = BareMetalHostDisk{} }
func (m *BareMetalHostDisk) String() string { return proto.CompactTextString(m) }
func (*BareMetalHostDisk) ProtoMessage()    {}
func (*BareMetalHostDisk) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{1}
}

func (m *BareMetalHostDisk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostDisk.Unmarshal(m, b)
}
func (m *BareMetalHostDisk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostDisk.Marshal(b, m, deterministic)
}
func (m *BareMetalHostDisk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostDisk.Merge(m, src)
}
func (m *BareMetalHostDisk) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostDisk.Size(m)
}
func (m *BareMetalHostDisk) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostDisk.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostDisk proto.InternalMessageInfo

func (m *BareMetalHostDisk) GetDiskName() string {
	if m != nil {
		return m.DiskName
	}
	return ""
}

func (m *BareMetalHostDisk) GetDiskSizeGB() int32 {
	if m != nil {
		return m.DiskSizeGB
	}
	return 0
}

type BareMetalHostStorageConfiguration struct {
	Disks                []*BareMetalHostDisk `protobuf:"bytes,1,rep,name=disks,proto3" json:"disks,omitempty"`
	ImageReference       string               `protobuf:"bytes,2,opt,name=imageReference,proto3" json:"imageReference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BareMetalHostStorageConfiguration) Reset()         { *m = BareMetalHostStorageConfiguration{} }
func (m *BareMetalHostStorageConfiguration) String() string { return proto.CompactTextString(m) }
func (*BareMetalHostStorageConfiguration) ProtoMessage()    {}
func (*BareMetalHostStorageConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{2}
}

func (m *BareMetalHostStorageConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostStorageConfiguration.Unmarshal(m, b)
}
func (m *BareMetalHostStorageConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostStorageConfiguration.Marshal(b, m, deterministic)
}
func (m *BareMetalHostStorageConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostStorageConfiguration.Merge(m, src)
}
func (m *BareMetalHostStorageConfiguration) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostStorageConfiguration.Size(m)
}
func (m *BareMetalHostStorageConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostStorageConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostStorageConfiguration proto.InternalMessageInfo

func (m *BareMetalHostStorageConfiguration) GetDisks() []*BareMetalHostDisk {
	if m != nil {
		return m.Disks
	}
	return nil
}

func (m *BareMetalHostStorageConfiguration) GetImageReference() string {
	if m != nil {
		return m.ImageReference
	}
	return ""
}

type BareMetalHostUserConfiguration struct {
	Username             string          `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string          `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Usertype             common.UserType `protobuf:"varint,3,opt,name=usertype,proto3,enum=moc.UserType" json:"usertype,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BareMetalHostUserConfiguration) Reset()         { *m = BareMetalHostUserConfiguration{} }
func (m *BareMetalHostUserConfiguration) String() string { return proto.CompactTextString(m) }
func (*BareMetalHostUserConfiguration) ProtoMessage()    {}
func (*BareMetalHostUserConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{3}
}

func (m *BareMetalHostUserConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostUserConfiguration.Unmarshal(m, b)
}
func (m *BareMetalHostUserConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostUserConfiguration.Marshal(b, m, deterministic)
}
func (m *BareMetalHostUserConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostUserConfiguration.Merge(m, src)
}
func (m *BareMetalHostUserConfiguration) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostUserConfiguration.Size(m)
}
func (m *BareMetalHostUserConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostUserConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostUserConfiguration proto.InternalMessageInfo

func (m *BareMetalHostUserConfiguration) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *BareMetalHostUserConfiguration) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *BareMetalHostUserConfiguration) GetUsertype() common.UserType {
	if m != nil {
		return m.Usertype
	}
	return common.UserType_ROOT
}

type BareMetalHostRDPConfiguration struct {
	DisableRDP           bool     `protobuf:"varint,1,opt,name=disableRDP,proto3" json:"disableRDP,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BareMetalHostRDPConfiguration) Reset()         { *m = BareMetalHostRDPConfiguration{} }
func (m *BareMetalHostRDPConfiguration) String() string { return proto.CompactTextString(m) }
func (*BareMetalHostRDPConfiguration) ProtoMessage()    {}
func (*BareMetalHostRDPConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{4}
}

func (m *BareMetalHostRDPConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostRDPConfiguration.Unmarshal(m, b)
}
func (m *BareMetalHostRDPConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostRDPConfiguration.Marshal(b, m, deterministic)
}
func (m *BareMetalHostRDPConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostRDPConfiguration.Merge(m, src)
}
func (m *BareMetalHostRDPConfiguration) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostRDPConfiguration.Size(m)
}
func (m *BareMetalHostRDPConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostRDPConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostRDPConfiguration proto.InternalMessageInfo

func (m *BareMetalHostRDPConfiguration) GetDisableRDP() bool {
	if m != nil {
		return m.DisableRDP
	}
	return false
}

type BareMetalHostLinuxConfiguration struct {
	DisablePasswordAuthentication bool     `protobuf:"varint,1,opt,name=disablePasswordAuthentication,proto3" json:"disablePasswordAuthentication,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *BareMetalHostLinuxConfiguration) Reset()         { *m = BareMetalHostLinuxConfiguration{} }
func (m *BareMetalHostLinuxConfiguration) String() string { return proto.CompactTextString(m) }
func (*BareMetalHostLinuxConfiguration) ProtoMessage()    {}
func (*BareMetalHostLinuxConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{5}
}

func (m *BareMetalHostLinuxConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostLinuxConfiguration.Unmarshal(m, b)
}
func (m *BareMetalHostLinuxConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostLinuxConfiguration.Marshal(b, m, deterministic)
}
func (m *BareMetalHostLinuxConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostLinuxConfiguration.Merge(m, src)
}
func (m *BareMetalHostLinuxConfiguration) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostLinuxConfiguration.Size(m)
}
func (m *BareMetalHostLinuxConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostLinuxConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostLinuxConfiguration proto.InternalMessageInfo

func (m *BareMetalHostLinuxConfiguration) GetDisablePasswordAuthentication() bool {
	if m != nil {
		return m.DisablePasswordAuthentication
	}
	return false
}

type BareMetalHostWindowsConfiguration struct {
	EnableAutomaticUpdates bool                           `protobuf:"varint,1,opt,name=enableAutomaticUpdates,proto3" json:"enableAutomaticUpdates,omitempty"`
	TimeZone               string                         `protobuf:"bytes,2,opt,name=timeZone,proto3" json:"timeZone,omitempty"`
	RDPConfiguration       *BareMetalHostRDPConfiguration `protobuf:"bytes,3,opt,name=RDPConfiguration,proto3" json:"RDPConfiguration,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                       `json:"-"`
	XXX_unrecognized       []byte                         `json:"-"`
	XXX_sizecache          int32                          `json:"-"`
}

func (m *BareMetalHostWindowsConfiguration) Reset()         { *m = BareMetalHostWindowsConfiguration{} }
func (m *BareMetalHostWindowsConfiguration) String() string { return proto.CompactTextString(m) }
func (*BareMetalHostWindowsConfiguration) ProtoMessage()    {}
func (*BareMetalHostWindowsConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{6}
}

func (m *BareMetalHostWindowsConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostWindowsConfiguration.Unmarshal(m, b)
}
func (m *BareMetalHostWindowsConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostWindowsConfiguration.Marshal(b, m, deterministic)
}
func (m *BareMetalHostWindowsConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostWindowsConfiguration.Merge(m, src)
}
func (m *BareMetalHostWindowsConfiguration) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostWindowsConfiguration.Size(m)
}
func (m *BareMetalHostWindowsConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostWindowsConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostWindowsConfiguration proto.InternalMessageInfo

func (m *BareMetalHostWindowsConfiguration) GetEnableAutomaticUpdates() bool {
	if m != nil {
		return m.EnableAutomaticUpdates
	}
	return false
}

func (m *BareMetalHostWindowsConfiguration) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

func (m *BareMetalHostWindowsConfiguration) GetRDPConfiguration() *BareMetalHostRDPConfiguration {
	if m != nil {
		return m.RDPConfiguration
	}
	return nil
}

type BareMetalHostOperatingSystemConfiguration struct {
	ComputerName  string                            `protobuf:"bytes,1,opt,name=computerName,proto3" json:"computerName,omitempty"`
	Administrator *BareMetalHostUserConfiguration   `protobuf:"bytes,2,opt,name=administrator,proto3" json:"administrator,omitempty"`
	Users         []*BareMetalHostUserConfiguration `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	CustomData    string                            `protobuf:"bytes,4,opt,name=customData,proto3" json:"customData,omitempty"`
	Publickeys    []*BareMetalHostSSHPublicKey      `protobuf:"bytes,5,rep,name=publickeys,proto3" json:"publickeys,omitempty"`
	Ostype        common.OperatingSystemType        `protobuf:"varint,6,opt,name=ostype,proto3,enum=moc.OperatingSystemType" json:"ostype,omitempty"`
	// bootstrap engine can be cloud-init, Windows answer files, ...
	OsBootstrapEngine    common.OperatingSystemBootstrapEngine `protobuf:"varint,7,opt,name=osBootstrapEngine,proto3,enum=moc.OperatingSystemBootstrapEngine" json:"osBootstrapEngine,omitempty"`
	LinuxConfiguration   *BareMetalHostLinuxConfiguration      `protobuf:"bytes,8,opt,name=linuxConfiguration,proto3" json:"linuxConfiguration,omitempty"`
	WindowsConfiguration *BareMetalHostWindowsConfiguration    `protobuf:"bytes,9,opt,name=windowsConfiguration,proto3" json:"windowsConfiguration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *BareMetalHostOperatingSystemConfiguration) Reset() {
	*m = BareMetalHostOperatingSystemConfiguration{}
}
func (m *BareMetalHostOperatingSystemConfiguration) String() string {
	return proto.CompactTextString(m)
}
func (*BareMetalHostOperatingSystemConfiguration) ProtoMessage() {}
func (*BareMetalHostOperatingSystemConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{7}
}

func (m *BareMetalHostOperatingSystemConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostOperatingSystemConfiguration.Unmarshal(m, b)
}
func (m *BareMetalHostOperatingSystemConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostOperatingSystemConfiguration.Marshal(b, m, deterministic)
}
func (m *BareMetalHostOperatingSystemConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostOperatingSystemConfiguration.Merge(m, src)
}
func (m *BareMetalHostOperatingSystemConfiguration) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostOperatingSystemConfiguration.Size(m)
}
func (m *BareMetalHostOperatingSystemConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostOperatingSystemConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostOperatingSystemConfiguration proto.InternalMessageInfo

func (m *BareMetalHostOperatingSystemConfiguration) GetComputerName() string {
	if m != nil {
		return m.ComputerName
	}
	return ""
}

func (m *BareMetalHostOperatingSystemConfiguration) GetAdministrator() *BareMetalHostUserConfiguration {
	if m != nil {
		return m.Administrator
	}
	return nil
}

func (m *BareMetalHostOperatingSystemConfiguration) GetUsers() []*BareMetalHostUserConfiguration {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *BareMetalHostOperatingSystemConfiguration) GetCustomData() string {
	if m != nil {
		return m.CustomData
	}
	return ""
}

func (m *BareMetalHostOperatingSystemConfiguration) GetPublickeys() []*BareMetalHostSSHPublicKey {
	if m != nil {
		return m.Publickeys
	}
	return nil
}

func (m *BareMetalHostOperatingSystemConfiguration) GetOstype() common.OperatingSystemType {
	if m != nil {
		return m.Ostype
	}
	return common.OperatingSystemType_WINDOWS
}

func (m *BareMetalHostOperatingSystemConfiguration) GetOsBootstrapEngine() common.OperatingSystemBootstrapEngine {
	if m != nil {
		return m.OsBootstrapEngine
	}
	return common.OperatingSystemBootstrapEngine_CLOUD_INIT
}

func (m *BareMetalHostOperatingSystemConfiguration) GetLinuxConfiguration() *BareMetalHostLinuxConfiguration {
	if m != nil {
		return m.LinuxConfiguration
	}
	return nil
}

func (m *BareMetalHostOperatingSystemConfiguration) GetWindowsConfiguration() *BareMetalHostWindowsConfiguration {
	if m != nil {
		return m.WindowsConfiguration
	}
	return nil
}

type BareMetalHost struct {
	Name                 string                                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                                     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Storage              *BareMetalHostStorageConfiguration         `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	Os                   *BareMetalHostOperatingSystemConfiguration `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	Status               *common.Status                             `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *common.Entity                             `protobuf:"bytes,6,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags                               `protobuf:"bytes,7,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *BareMetalHost) Reset()         { *m = BareMetalHost{} }
func (m *BareMetalHost) String() string { return proto.CompactTextString(m) }
func (*BareMetalHost) ProtoMessage()    {}
func (*BareMetalHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{8}
}

func (m *BareMetalHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHost.Unmarshal(m, b)
}
func (m *BareMetalHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHost.Marshal(b, m, deterministic)
}
func (m *BareMetalHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHost.Merge(m, src)
}
func (m *BareMetalHost) XXX_Size() int {
	return xxx_messageInfo_BareMetalHost.Size(m)
}
func (m *BareMetalHost) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHost.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHost proto.InternalMessageInfo

func (m *BareMetalHost) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BareMetalHost) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BareMetalHost) GetStorage() *BareMetalHostStorageConfiguration {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *BareMetalHost) GetOs() *BareMetalHostOperatingSystemConfiguration {
	if m != nil {
		return m.Os
	}
	return nil
}

func (m *BareMetalHost) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BareMetalHost) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *BareMetalHost) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type BareMetalHostRequest struct {
	BareMetalHost        *BareMetalHost `protobuf:"bytes,1,opt,name=bareMetalHost,proto3" json:"bareMetalHost,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BareMetalHostRequest) Reset()         { *m = BareMetalHostRequest{} }
func (m *BareMetalHostRequest) String() string { return proto.CompactTextString(m) }
func (*BareMetalHostRequest) ProtoMessage()    {}
func (*BareMetalHostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{9}
}

func (m *BareMetalHostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostRequest.Unmarshal(m, b)
}
func (m *BareMetalHostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostRequest.Marshal(b, m, deterministic)
}
func (m *BareMetalHostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostRequest.Merge(m, src)
}
func (m *BareMetalHostRequest) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostRequest.Size(m)
}
func (m *BareMetalHostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostRequest proto.InternalMessageInfo

func (m *BareMetalHostRequest) GetBareMetalHost() *BareMetalHost {
	if m != nil {
		return m.BareMetalHost
	}
	return nil
}

type BareMetalHostResponse struct {
	BareMetalHost        *BareMetalHost      `protobuf:"bytes,1,opt,name=bareMetalHost,proto3" json:"bareMetalHost,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BareMetalHostResponse) Reset()         { *m = BareMetalHostResponse{} }
func (m *BareMetalHostResponse) String() string { return proto.CompactTextString(m) }
func (*BareMetalHostResponse) ProtoMessage()    {}
func (*BareMetalHostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{10}
}

func (m *BareMetalHostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostResponse.Unmarshal(m, b)
}
func (m *BareMetalHostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostResponse.Marshal(b, m, deterministic)
}
func (m *BareMetalHostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostResponse.Merge(m, src)
}
func (m *BareMetalHostResponse) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostResponse.Size(m)
}
func (m *BareMetalHostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostResponse proto.InternalMessageInfo

func (m *BareMetalHostResponse) GetBareMetalHost() *BareMetalHost {
	if m != nil {
		return m.BareMetalHost
	}
	return nil
}

func (m *BareMetalHostResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *BareMetalHostResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*BareMetalHostSSHPublicKey)(nil), "moc.baremetalhostagent.BareMetalHostSSHPublicKey")
	proto.RegisterType((*BareMetalHostDisk)(nil), "moc.baremetalhostagent.BareMetalHostDisk")
	proto.RegisterType((*BareMetalHostStorageConfiguration)(nil), "moc.baremetalhostagent.BareMetalHostStorageConfiguration")
	proto.RegisterType((*BareMetalHostUserConfiguration)(nil), "moc.baremetalhostagent.BareMetalHostUserConfiguration")
	proto.RegisterType((*BareMetalHostRDPConfiguration)(nil), "moc.baremetalhostagent.BareMetalHostRDPConfiguration")
	proto.RegisterType((*BareMetalHostLinuxConfiguration)(nil), "moc.baremetalhostagent.BareMetalHostLinuxConfiguration")
	proto.RegisterType((*BareMetalHostWindowsConfiguration)(nil), "moc.baremetalhostagent.BareMetalHostWindowsConfiguration")
	proto.RegisterType((*BareMetalHostOperatingSystemConfiguration)(nil), "moc.baremetalhostagent.BareMetalHostOperatingSystemConfiguration")
	proto.RegisterType((*BareMetalHost)(nil), "moc.baremetalhostagent.BareMetalHost")
	proto.RegisterType((*BareMetalHostRequest)(nil), "moc.baremetalhostagent.BareMetalHostRequest")
	proto.RegisterType((*BareMetalHostResponse)(nil), "moc.baremetalhostagent.BareMetalHostResponse")
}

func init() { proto.RegisterFile("moc_baremetalhostagent.proto", fileDescriptor_254a3462c20ef4c7) }

var fileDescriptor_254a3462c20ef4c7 = []byte{
	// 895 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x51, 0x6f, 0x1b, 0x45,
	0x10, 0xc6, 0x4e, 0xec, 0xc4, 0x63, 0x12, 0xd1, 0x55, 0x28, 0x87, 0x45, 0x42, 0xb8, 0x0a, 0x94,
	0x48, 0x60, 0xd3, 0x43, 0x2d, 0xe2, 0xa9, 0x8a, 0x71, 0x44, 0xa5, 0x16, 0x9a, 0xac, 0x5b, 0x90,
	0x2a, 0x24, 0xb4, 0x3e, 0x4f, 0x2e, 0xab, 0xf8, 0x76, 0x8f, 0xdd, 0x3d, 0x05, 0x23, 0xf1, 0xc4,
	0x2b, 0xbf, 0x85, 0x5f, 0xc1, 0x1b, 0xff, 0x80, 0x5f, 0x83, 0x76, 0xef, 0x6c, 0xdd, 0xda, 0x26,
	0x3a, 0xa1, 0x3e, 0xd9, 0x3b, 0x33, 0xdf, 0xf7, 0xcd, 0xce, 0xce, 0x8c, 0x0e, 0x3e, 0x48, 0x65,
	0xfc, 0xd3, 0x84, 0x29, 0x4c, 0xd1, 0xb0, 0xd9, 0xb5, 0xd4, 0x86, 0x25, 0x28, 0x4c, 0x3f, 0x53,
	0xd2, 0x48, 0x72, 0x3f, 0x95, 0x71, 0x7f, 0xdd, 0xdb, 0x3b, 0x4a, 0xa4, 0x4c, 0x66, 0x38, 0x70,
	0x51, 0x93, 0xfc, 0x6a, 0x70, 0xab, 0x58, 0x96, 0xa1, 0xd2, 0x05, 0xae, 0xf7, 0x9e, 0x65, 0x8d,
	0x65, 0x9a, 0x4a, 0x51, 0xfe, 0x94, 0x8e, 0x23, 0xdf, 0x91, 0xe5, 0x06, 0xab, 0xfe, 0xf0, 0x11,
	0xbc, 0x3f, 0x64, 0x0a, 0xbf, 0xb5, 0x72, 0x4f, 0xa5, 0x36, 0xe3, 0xf1, 0xd3, 0x8b, 0x7c, 0x32,
	0xe3, 0xf1, 0x33, 0x9c, 0x93, 0x00, 0x76, 0x6e, 0x70, 0x3e, 0x65, 0x86, 0x05, 0x8d, 0xe3, 0xc6,
	0x49, 0x87, 0x2e, 0x8e, 0xe1, 0x0b, 0xb8, 0xe7, 0xc1, 0x46, 0x5c, 0xdf, 0x90, 0x1e, 0xec, 0x4e,
	0xb9, 0xbe, 0xf9, 0x8e, 0xa5, 0x58, 0xc6, 0x2f, 0xcf, 0xe4, 0x08, 0xc0, 0xfe, 0x1f, 0xf3, 0x5f,
	0xf1, 0x9b, 0x61, 0xd0, 0x3c, 0x6e, 0x9c, 0xb4, 0x68, 0xc5, 0x12, 0xfe, 0xd1, 0x80, 0x8f, 0xfc,
	0x44, 0x8c, 0x54, 0x2c, 0xc1, 0xaf, 0xa5, 0xb8, 0xe2, 0x49, 0xae, 0x98, 0xe1, 0x52, 0x90, 0x27,
	0xd0, 0xb2, 0x18, 0x1d, 0x34, 0x8e, 0xb7, 0x4e, 0xba, 0xd1, 0x69, 0x7f, 0x73, 0xb9, 0xfa, 0x6b,
	0xb9, 0xd1, 0x02, 0x47, 0x3e, 0x81, 0x7d, 0x9e, 0xb2, 0x04, 0x29, 0x5e, 0xa1, 0x42, 0x11, 0xa3,
	0x4b, 0xa5, 0x43, 0x57, 0xac, 0xe1, 0xef, 0x0d, 0x38, 0xf2, 0x48, 0x5e, 0x69, 0x54, 0x7e, 0x2e,
	0x3d, 0xd8, 0xcd, 0x35, 0x2a, 0x51, 0xb9, 0xed, 0xe2, 0x6c, 0x7d, 0x19, 0xd3, 0xfa, 0x56, 0xaa,
	0x69, 0x29, 0xb0, 0x3c, 0x93, 0xd3, 0x02, 0x67, 0xe6, 0x19, 0x06, 0x5b, 0xc7, 0x8d, 0x93, 0xfd,
	0x68, 0xcf, 0x5d, 0xc3, 0x2a, 0xbc, 0x9c, 0x67, 0x48, 0x97, 0xee, 0xf0, 0x09, 0x1c, 0x7a, 0x49,
	0xd0, 0xd1, 0x85, 0x9f, 0x43, 0x51, 0x55, 0x36, 0x99, 0x21, 0x1d, 0x5d, 0xb8, 0x2c, 0x76, 0x69,
	0xc5, 0x12, 0x26, 0xf0, 0xa1, 0x47, 0xf0, 0x9c, 0x8b, 0xfc, 0x17, 0x9f, 0x62, 0x04, 0x87, 0x25,
	0xe0, 0xa2, 0xcc, 0xf0, 0x2c, 0x37, 0xd7, 0x28, 0x0c, 0x8f, 0x5d, 0x40, 0xc9, 0x7a, 0x77, 0x50,
	0xf8, 0xcf, 0xea, 0xf3, 0xfd, 0xc0, 0xc5, 0x54, 0xde, 0x6a, 0x5f, 0xeb, 0x31, 0xdc, 0x47, 0x61,
	0x59, 0xce, 0x72, 0x23, 0x53, 0x66, 0x78, 0xfc, 0x2a, 0x9b, 0x32, 0x83, 0xba, 0x14, 0xf9, 0x0f,
	0xaf, 0x2d, 0xa7, 0xe1, 0x29, 0xbe, 0x96, 0x62, 0xf1, 0x5e, 0xcb, 0x33, 0x61, 0xf0, 0xce, 0x6a,
	0x59, 0x5c, 0x59, 0xbb, 0xd1, 0xa3, 0x5a, 0xdd, 0xb1, 0x0a, 0xa6, 0x6b, 0x74, 0xe1, 0xdf, 0x2d,
	0x38, 0xf5, 0x30, 0x2f, 0x32, 0xb4, 0x2e, 0x91, 0x8c, 0xe7, 0xda, 0x60, 0xea, 0x5f, 0x32, 0x84,
	0xb7, 0xcb, 0x41, 0x53, 0x95, 0x49, 0xf0, 0x6c, 0xe4, 0x47, 0xd8, 0x63, 0xd3, 0x94, 0x0b, 0xae,
	0x8d, 0x62, 0x46, 0x2a, 0x77, 0xab, 0x6e, 0xf4, 0xb8, 0x56, 0xc6, 0x6b, 0xad, 0x48, 0x7d, 0x32,
	0xf2, 0x1c, 0x5a, 0xb6, 0x85, 0x74, 0xb0, 0xe5, 0xa6, 0xe4, 0xff, 0xb2, 0x16, 0x24, 0xb6, 0xc7,
	0xe2, 0x5c, 0x1b, 0x99, 0x8e, 0xec, 0x1e, 0xd8, 0x76, 0xb7, 0xa9, 0x58, 0xc8, 0x25, 0x40, 0xe6,
	0x36, 0xc6, 0x0d, 0xce, 0x75, 0xd0, 0x72, 0x92, 0x0f, 0x6b, 0x49, 0x56, 0x77, 0x0d, 0xad, 0x90,
	0x90, 0xcf, 0xa1, 0x2d, 0xb5, 0x1b, 0x90, 0xb6, 0x1b, 0x90, 0xc0, 0xd1, 0xad, 0x54, 0xdd, 0xcd,
	0x4a, 0x19, 0x47, 0x2e, 0xe1, 0x9e, 0xd4, 0x43, 0x29, 0x8d, 0xad, 0x41, 0x76, 0x2e, 0x12, 0x2e,
	0x30, 0xd8, 0x71, 0xe0, 0x07, 0x9b, 0xc0, 0x2b, 0xa1, 0x74, 0x1d, 0x4d, 0x12, 0x20, 0xb3, 0xb5,
	0x71, 0x09, 0x76, 0xdd, 0x43, 0x7d, 0x59, 0xeb, 0x7e, 0xeb, 0xd3, 0x46, 0x37, 0x50, 0x92, 0x14,
	0x0e, 0x6e, 0x37, 0x4c, 0x4b, 0xd0, 0x71, 0x52, 0x5f, 0xd5, 0x92, 0xda, 0x34, 0x6e, 0x74, 0x23,
	0x6d, 0xf8, 0x57, 0x13, 0xf6, 0x3c, 0x2c, 0x21, 0xb0, 0x5d, 0xd9, 0x62, 0xee, 0x3f, 0xd9, 0x87,
	0x26, 0x5f, 0xec, 0xae, 0x26, 0x9f, 0x92, 0x31, 0xec, 0xe8, 0x62, 0x23, 0x97, 0xd3, 0x55, 0x2f,
	0xaf, 0x4d, 0x5b, 0x9c, 0x2e, 0x98, 0xc8, 0x25, 0x34, 0xa5, 0x76, 0x2d, 0xd5, 0x8d, 0xce, 0x6a,
	0xf1, 0xdd, 0x35, 0x79, 0xb4, 0x29, 0x35, 0x79, 0x00, 0x6d, 0x6d, 0x98, 0xc9, 0x6d, 0x27, 0x5a,
	0xda, 0xae, 0xa3, 0x1d, 0x3b, 0x13, 0x2d, 0x5d, 0x36, 0xc8, 0xee, 0x2e, 0x33, 0x77, 0xfd, 0xb5,
	0x08, 0x3a, 0x77, 0x26, 0x5a, 0xba, 0xc8, 0x21, 0x6c, 0x1b, 0x96, 0x68, 0xd7, 0x45, 0xdd, 0xa8,
	0xe3, 0x42, 0x5e, 0xb2, 0x44, 0x53, 0x67, 0x0e, 0x63, 0x38, 0xf0, 0xf7, 0x08, 0xfe, 0x9c, 0xa3,
	0x36, 0xe4, 0x19, 0xec, 0x4d, 0xaa, 0x76, 0x57, 0xd5, 0x6e, 0xf4, 0x71, 0xbd, 0x65, 0xe4, 0x63,
	0xc3, 0x3f, 0x1b, 0xf0, 0xee, 0x8a, 0x8a, 0xce, 0xa4, 0xd0, 0xf8, 0x46, 0x65, 0x48, 0x04, 0x6d,
	0x8a, 0x3a, 0x9f, 0x99, 0x72, 0x0f, 0xf5, 0xfa, 0xc5, 0xe7, 0x46, 0x7f, 0xf1, 0xb9, 0xd1, 0x1f,
	0x4a, 0x39, 0xfb, 0x9e, 0xcd, 0x72, 0xa4, 0x65, 0x24, 0x39, 0x80, 0xd6, 0xb9, 0x52, 0x52, 0xb9,
	0x76, 0xe8, 0xd0, 0xe2, 0x10, 0xfd, 0x06, 0xc4, 0x53, 0x3a, 0xb3, 0xe2, 0x24, 0x81, 0x76, 0xb1,
	0xca, 0xc9, 0xa7, 0xf5, 0xf2, 0x2b, 0x6a, 0xd9, 0xfb, 0xac, 0x66, 0x74, 0x51, 0x93, 0xf0, 0xad,
	0xe1, 0xc3, 0xd7, 0x83, 0x84, 0x9b, 0xeb, 0x7c, 0xd2, 0x8f, 0x65, 0x3a, 0x48, 0x79, 0xac, 0xa4,
	0x96, 0x57, 0x66, 0x90, 0xca, 0x78, 0xa0, 0xb2, 0x78, 0xb0, 0x4e, 0x35, 0x69, 0xbb, 0x3b, 0x7e,
	0xf1, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xfb, 0xd5, 0x13, 0x98, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BareMetalHostAgentClient is the client API for BareMetalHostAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BareMetalHostAgentClient interface {
	Update(ctx context.Context, in *BareMetalHostRequest, opts ...grpc.CallOption) (*BareMetalHostResponse, error)
}

type bareMetalHostAgentClient struct {
	cc *grpc.ClientConn
}

func NewBareMetalHostAgentClient(cc *grpc.ClientConn) BareMetalHostAgentClient {
	return &bareMetalHostAgentClient{cc}
}

func (c *bareMetalHostAgentClient) Update(ctx context.Context, in *BareMetalHostRequest, opts ...grpc.CallOption) (*BareMetalHostResponse, error) {
	out := new(BareMetalHostResponse)
	err := c.cc.Invoke(ctx, "/moc.baremetalhostagent.BareMetalHostAgent/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BareMetalHostAgentServer is the server API for BareMetalHostAgent service.
type BareMetalHostAgentServer interface {
	Update(context.Context, *BareMetalHostRequest) (*BareMetalHostResponse, error)
}

// UnimplementedBareMetalHostAgentServer can be embedded to have forward compatible implementations.
type UnimplementedBareMetalHostAgentServer struct {
}

func (*UnimplementedBareMetalHostAgentServer) Update(ctx context.Context, req *BareMetalHostRequest) (*BareMetalHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterBareMetalHostAgentServer(s *grpc.Server, srv BareMetalHostAgentServer) {
	s.RegisterService(&_BareMetalHostAgent_serviceDesc, srv)
}

func _BareMetalHostAgent_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BareMetalHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareMetalHostAgentServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.baremetalhostagent.BareMetalHostAgent/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareMetalHostAgentServer).Update(ctx, req.(*BareMetalHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BareMetalHostAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.baremetalhostagent.BareMetalHostAgent",
	HandlerType: (*BareMetalHostAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _BareMetalHostAgent_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_baremetalhostagent.proto",
}