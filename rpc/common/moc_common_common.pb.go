// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_common_common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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
	Operation_GET    Operation = 0
	Operation_POST   Operation = 1
	Operation_DELETE Operation = 2
	Operation_UPDATE Operation = 3
)

var Operation_name = map[int32]string{
	0: "GET",
	1: "POST",
	2: "DELETE",
	3: "UPDATE",
}

var Operation_value = map[string]int32{
	"GET":    0,
	"POST":   1,
	"DELETE": 2,
	"UPDATE": 3,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}

func (Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{0}
}

type ProvisionState int32

const (
	ProvisionState_UNKNOWN            ProvisionState = 0
	ProvisionState_CREATING           ProvisionState = 1
	ProvisionState_CREATED            ProvisionState = 2
	ProvisionState_CREATE_FAILED      ProvisionState = 3
	ProvisionState_DELETING           ProvisionState = 4
	ProvisionState_DELETE_FAILED      ProvisionState = 5
	ProvisionState_DELETED            ProvisionState = 6
	ProvisionState_UPDATING           ProvisionState = 7
	ProvisionState_UPDATE_FAILED      ProvisionState = 8
	ProvisionState_UPDATED            ProvisionState = 9
	ProvisionState_PROVISIONING       ProvisionState = 10
	ProvisionState_PROVISIONED        ProvisionState = 11
	ProvisionState_PROVISION_FAILED   ProvisionState = 12
	ProvisionState_DEPROVISIONING     ProvisionState = 13
	ProvisionState_DEPROVISIONED      ProvisionState = 14
	ProvisionState_DEPROVISION_FAILED ProvisionState = 15
	ProvisionState_DELETE_PENDING     ProvisionState = 16
)

var ProvisionState_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CREATING",
	2:  "CREATED",
	3:  "CREATE_FAILED",
	4:  "DELETING",
	5:  "DELETE_FAILED",
	6:  "DELETED",
	7:  "UPDATING",
	8:  "UPDATE_FAILED",
	9:  "UPDATED",
	10: "PROVISIONING",
	11: "PROVISIONED",
	12: "PROVISION_FAILED",
	13: "DEPROVISIONING",
	14: "DEPROVISIONED",
	15: "DEPROVISION_FAILED",
	16: "DELETE_PENDING",
}

var ProvisionState_value = map[string]int32{
	"UNKNOWN":            0,
	"CREATING":           1,
	"CREATED":            2,
	"CREATE_FAILED":      3,
	"DELETING":           4,
	"DELETE_FAILED":      5,
	"DELETED":            6,
	"UPDATING":           7,
	"UPDATE_FAILED":      8,
	"UPDATED":            9,
	"PROVISIONING":       10,
	"PROVISIONED":        11,
	"PROVISION_FAILED":   12,
	"DEPROVISIONING":     13,
	"DEPROVISIONED":      14,
	"DEPROVISION_FAILED": 15,
	"DELETE_PENDING":     16,
}

func (x ProvisionState) String() string {
	return proto.EnumName(ProvisionState_name, int32(x))
}

func (ProvisionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{1}
}

type HighAvailabilityState int32

const (
	HighAvailabilityState_UNKNOWN_HA_STATE HighAvailabilityState = 0
	HighAvailabilityState_STABLE           HighAvailabilityState = 1
	HighAvailabilityState_PENDING          HighAvailabilityState = 2
)

var HighAvailabilityState_name = map[int32]string{
	0: "UNKNOWN_HA_STATE",
	1: "STABLE",
	2: "PENDING",
}

var HighAvailabilityState_value = map[string]int32{
	"UNKNOWN_HA_STATE": 0,
	"STABLE":           1,
	"PENDING":          2,
}

func (x HighAvailabilityState) String() string {
	return proto.EnumName(HighAvailabilityState_name, int32(x))
}

func (HighAvailabilityState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{2}
}

type HealthState int32

const (
	HealthState_NOTKNOWN HealthState = 0
	HealthState_OK       HealthState = 1
	HealthState_WARNING  HealthState = 2
	HealthState_CRITICAL HealthState = 3
	// The entity went missing from the platform
	HealthState_MISSING  HealthState = 4
	HealthState_DEGRADED HealthState = 5
	// The entity went missing from the agent
	HealthState_NOTFOUND HealthState = 6
)

var HealthState_name = map[int32]string{
	0: "NOTKNOWN",
	1: "OK",
	2: "WARNING",
	3: "CRITICAL",
	4: "MISSING",
	5: "DEGRADED",
	6: "NOTFOUND",
}

var HealthState_value = map[string]int32{
	"NOTKNOWN": 0,
	"OK":       1,
	"WARNING":  2,
	"CRITICAL": 3,
	"MISSING":  4,
	"DEGRADED": 5,
	"NOTFOUND": 6,
}

func (x HealthState) String() string {
	return proto.EnumName(HealthState_name, int32(x))
}

func (HealthState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{3}
}

type ClientType int32

const (
	ClientType_CONTROLPLANE   ClientType = 0
	ClientType_EXTERNALCLIENT ClientType = 1
	ClientType_NODE           ClientType = 2
	ClientType_ADMIN          ClientType = 3
	ClientType_BAREMETAL      ClientType = 4
	ClientType_LOADBALANCER   ClientType = 5
)

var ClientType_name = map[int32]string{
	0: "CONTROLPLANE",
	1: "EXTERNALCLIENT",
	2: "NODE",
	3: "ADMIN",
	4: "BAREMETAL",
	5: "LOADBALANCER",
}

var ClientType_value = map[string]int32{
	"CONTROLPLANE":   0,
	"EXTERNALCLIENT": 1,
	"NODE":           2,
	"ADMIN":          3,
	"BAREMETAL":      4,
	"LOADBALANCER":   5,
}

func (x ClientType) String() string {
	return proto.EnumName(ClientType_name, int32(x))
}

func (ClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{4}
}

type AuthenticationType int32

const (
	AuthenticationType_SELFSIGNED AuthenticationType = 0
	AuthenticationType_CASIGNED   AuthenticationType = 1
)

var AuthenticationType_name = map[int32]string{
	0: "SELFSIGNED",
	1: "CASIGNED",
}

var AuthenticationType_value = map[string]int32{
	"SELFSIGNED": 0,
	"CASIGNED":   1,
}

func (x AuthenticationType) String() string {
	return proto.EnumName(AuthenticationType_name, int32(x))
}

func (AuthenticationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{5}
}

type ProviderType int32

const (
	ProviderType_AnyProvider            ProviderType = 0
	ProviderType_VirtualMachine         ProviderType = 1
	ProviderType_VirtualMachineScaleSet ProviderType = 2
	ProviderType_LoadBalancer           ProviderType = 3
	ProviderType_VirtualNetwork         ProviderType = 4
	ProviderType_VirtualHardDisk        ProviderType = 5
	ProviderType_GalleryImage           ProviderType = 6
	ProviderType_VirtualMachineImage    ProviderType = 7
	ProviderType_NetworkInterface       ProviderType = 8
	ProviderType_Certificate            ProviderType = 9
	ProviderType_Key                    ProviderType = 10
	ProviderType_Secret                 ProviderType = 11
	ProviderType_KeyVault               ProviderType = 12
	ProviderType_Identity               ProviderType = 13
	ProviderType_Role                   ProviderType = 14
	ProviderType_RoleAssignment         ProviderType = 15
	ProviderType_Kubernetes             ProviderType = 16
	ProviderType_Cluster                ProviderType = 17
	ProviderType_ControlPlane           ProviderType = 18
	ProviderType_Group                  ProviderType = 19
	ProviderType_Node                   ProviderType = 20
	ProviderType_Location               ProviderType = 21
	ProviderType_StorageContainer       ProviderType = 22
	ProviderType_StorageFile            ProviderType = 23
	ProviderType_StorageDirectory       ProviderType = 24
	ProviderType_Subscription           ProviderType = 25
	ProviderType_VipPool                ProviderType = 26
	ProviderType_MacPool                ProviderType = 27
	ProviderType_EtcdCluster            ProviderType = 28
	ProviderType_EtcdServer             ProviderType = 29
	ProviderType_BareMetalMachine       ProviderType = 30
	ProviderType_CredentialMonitor      ProviderType = 31
	ProviderType_Logging                ProviderType = 32
	ProviderType_Recovery               ProviderType = 33
	ProviderType_Debug                  ProviderType = 34
)

var ProviderType_name = map[int32]string{
	0:  "AnyProvider",
	1:  "VirtualMachine",
	2:  "VirtualMachineScaleSet",
	3:  "LoadBalancer",
	4:  "VirtualNetwork",
	5:  "VirtualHardDisk",
	6:  "GalleryImage",
	7:  "VirtualMachineImage",
	8:  "NetworkInterface",
	9:  "Certificate",
	10: "Key",
	11: "Secret",
	12: "KeyVault",
	13: "Identity",
	14: "Role",
	15: "RoleAssignment",
	16: "Kubernetes",
	17: "Cluster",
	18: "ControlPlane",
	19: "Group",
	20: "Node",
	21: "Location",
	22: "StorageContainer",
	23: "StorageFile",
	24: "StorageDirectory",
	25: "Subscription",
	26: "VipPool",
	27: "MacPool",
	28: "EtcdCluster",
	29: "EtcdServer",
	30: "BareMetalMachine",
	31: "CredentialMonitor",
	32: "Logging",
	33: "Recovery",
	34: "Debug",
}

var ProviderType_value = map[string]int32{
	"AnyProvider":            0,
	"VirtualMachine":         1,
	"VirtualMachineScaleSet": 2,
	"LoadBalancer":           3,
	"VirtualNetwork":         4,
	"VirtualHardDisk":        5,
	"GalleryImage":           6,
	"VirtualMachineImage":    7,
	"NetworkInterface":       8,
	"Certificate":            9,
	"Key":                    10,
	"Secret":                 11,
	"KeyVault":               12,
	"Identity":               13,
	"Role":                   14,
	"RoleAssignment":         15,
	"Kubernetes":             16,
	"Cluster":                17,
	"ControlPlane":           18,
	"Group":                  19,
	"Node":                   20,
	"Location":               21,
	"StorageContainer":       22,
	"StorageFile":            23,
	"StorageDirectory":       24,
	"Subscription":           25,
	"VipPool":                26,
	"MacPool":                27,
	"EtcdCluster":            28,
	"EtcdServer":             29,
	"BareMetalMachine":       30,
	"CredentialMonitor":      31,
	"Logging":                32,
	"Recovery":               33,
	"Debug":                  34,
}

func (x ProviderType) String() string {
	return proto.EnumName(ProviderType_name, int32(x))
}

func (ProviderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{6}
}

type ImageSource int32

const (
	ImageSource_LOCAL_SOURCE ImageSource = 0
	ImageSource_SFS_SOURCE   ImageSource = 1
	ImageSource_HTTP_SOURCE  ImageSource = 2
	ImageSource_CLONE_SOURCE ImageSource = 3
)

var ImageSource_name = map[int32]string{
	0: "LOCAL_SOURCE",
	1: "SFS_SOURCE",
	2: "HTTP_SOURCE",
	3: "CLONE_SOURCE",
}

var ImageSource_value = map[string]int32{
	"LOCAL_SOURCE": 0,
	"SFS_SOURCE":   1,
	"HTTP_SOURCE":  2,
	"CLONE_SOURCE": 3,
}

func (x ImageSource) String() string {
	return proto.EnumName(ImageSource_name, int32(x))
}

func (ImageSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{7}
}

type Error struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type ProvisionStatus struct {
	CurrentState         ProvisionState `protobuf:"varint,1,opt,name=currentState,proto3,enum=moc.ProvisionState" json:"currentState,omitempty"`
	PreviousState        ProvisionState `protobuf:"varint,2,opt,name=previousState,proto3,enum=moc.ProvisionState" json:"previousState,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProvisionStatus) Reset()         { *m = ProvisionStatus{} }
func (m *ProvisionStatus) String() string { return proto.CompactTextString(m) }
func (*ProvisionStatus) ProtoMessage()    {}
func (*ProvisionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{1}
}

func (m *ProvisionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProvisionStatus.Unmarshal(m, b)
}
func (m *ProvisionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProvisionStatus.Marshal(b, m, deterministic)
}
func (m *ProvisionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvisionStatus.Merge(m, src)
}
func (m *ProvisionStatus) XXX_Size() int {
	return xxx_messageInfo_ProvisionStatus.Size(m)
}
func (m *ProvisionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvisionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ProvisionStatus proto.InternalMessageInfo

func (m *ProvisionStatus) GetCurrentState() ProvisionState {
	if m != nil {
		return m.CurrentState
	}
	return ProvisionState_UNKNOWN
}

func (m *ProvisionStatus) GetPreviousState() ProvisionState {
	if m != nil {
		return m.PreviousState
	}
	return ProvisionState_UNKNOWN
}

type Health struct {
	CurrentState         HealthState `protobuf:"varint,1,opt,name=currentState,proto3,enum=moc.HealthState" json:"currentState,omitempty"`
	PreviousState        HealthState `protobuf:"varint,2,opt,name=previousState,proto3,enum=moc.HealthState" json:"previousState,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Health) Reset()         { *m = Health{} }
func (m *Health) String() string { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()    {}
func (*Health) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{2}
}

func (m *Health) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Health.Unmarshal(m, b)
}
func (m *Health) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Health.Marshal(b, m, deterministic)
}
func (m *Health) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Health.Merge(m, src)
}
func (m *Health) XXX_Size() int {
	return xxx_messageInfo_Health.Size(m)
}
func (m *Health) XXX_DiscardUnknown() {
	xxx_messageInfo_Health.DiscardUnknown(m)
}

var xxx_messageInfo_Health proto.InternalMessageInfo

func (m *Health) GetCurrentState() HealthState {
	if m != nil {
		return m.CurrentState
	}
	return HealthState_NOTKNOWN
}

func (m *Health) GetPreviousState() HealthState {
	if m != nil {
		return m.PreviousState
	}
	return HealthState_NOTKNOWN
}

type Version struct {
	Number               string   `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{3}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type Status struct {
	Health               *Health          `protobuf:"bytes,1,opt,name=health,proto3" json:"health,omitempty"`
	ProvisioningStatus   *ProvisionStatus `protobuf:"bytes,2,opt,name=provisioningStatus,proto3" json:"provisioningStatus,omitempty"`
	LastError            *Error           `protobuf:"bytes,3,opt,name=lastError,proto3" json:"lastError,omitempty"`
	Version              *Version         `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{4}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetHealth() *Health {
	if m != nil {
		return m.Health
	}
	return nil
}

func (m *Status) GetProvisioningStatus() *ProvisionStatus {
	if m != nil {
		return m.ProvisioningStatus
	}
	return nil
}

func (m *Status) GetLastError() *Error {
	if m != nil {
		return m.LastError
	}
	return nil
}

func (m *Status) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type Entity struct {
	IsPlaceholder        bool     `protobuf:"varint,1,opt,name=IsPlaceholder,proto3" json:"IsPlaceholder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{5}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetIsPlaceholder() bool {
	if m != nil {
		return m.IsPlaceholder
	}
	return false
}

type Tag struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{6}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Tags struct {
	Tags                 []*Tag   `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tags) Reset()         { *m = Tags{} }
func (m *Tags) String() string { return proto.CompactTextString(m) }
func (*Tags) ProtoMessage()    {}
func (*Tags) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f78e46755eb93, []int{7}
}

func (m *Tags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tags.Unmarshal(m, b)
}
func (m *Tags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tags.Marshal(b, m, deterministic)
}
func (m *Tags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tags.Merge(m, src)
}
func (m *Tags) XXX_Size() int {
	return xxx_messageInfo_Tags.Size(m)
}
func (m *Tags) XXX_DiscardUnknown() {
	xxx_messageInfo_Tags.DiscardUnknown(m)
}

var xxx_messageInfo_Tags proto.InternalMessageInfo

func (m *Tags) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

var E_Sensitive = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50001,
	Name:          "moc.sensitive",
	Tag:           "varint,50001,opt,name=sensitive",
	Filename:      "moc_common_common.proto",
}

func init() {
	proto.RegisterEnum("moc.Operation", Operation_name, Operation_value)
	proto.RegisterEnum("moc.ProvisionState", ProvisionState_name, ProvisionState_value)
	proto.RegisterEnum("moc.HighAvailabilityState", HighAvailabilityState_name, HighAvailabilityState_value)
	proto.RegisterEnum("moc.HealthState", HealthState_name, HealthState_value)
	proto.RegisterEnum("moc.ClientType", ClientType_name, ClientType_value)
	proto.RegisterEnum("moc.AuthenticationType", AuthenticationType_name, AuthenticationType_value)
	proto.RegisterEnum("moc.ProviderType", ProviderType_name, ProviderType_value)
	proto.RegisterEnum("moc.ImageSource", ImageSource_name, ImageSource_value)
	proto.RegisterType((*Error)(nil), "moc.Error")
	proto.RegisterType((*ProvisionStatus)(nil), "moc.ProvisionStatus")
	proto.RegisterType((*Health)(nil), "moc.Health")
	proto.RegisterType((*Version)(nil), "moc.Version")
	proto.RegisterType((*Status)(nil), "moc.Status")
	proto.RegisterType((*Entity)(nil), "moc.Entity")
	proto.RegisterType((*Tag)(nil), "moc.Tag")
	proto.RegisterType((*Tags)(nil), "moc.Tags")
	proto.RegisterExtension(E_Sensitive)
}

func init() { proto.RegisterFile("moc_common_common.proto", fileDescriptor_681f78e46755eb93) }

var fileDescriptor_681f78e46755eb93 = []byte{
	// 1258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0xdd, 0x6e, 0xdb, 0xc6,
	0x12, 0x8e, 0x7e, 0x2c, 0x5b, 0x23, 0xff, 0x6c, 0xd6, 0x8e, 0xa3, 0x93, 0x93, 0x9c, 0xe3, 0x28,
	0x69, 0x61, 0x08, 0xa8, 0x0c, 0xb8, 0x6d, 0x8a, 0x16, 0xe8, 0x05, 0x2d, 0xd2, 0x36, 0x61, 0x9a,
	0x14, 0x48, 0xda, 0x29, 0x7a, 0x63, 0xac, 0xa8, 0x31, 0x45, 0x84, 0xe2, 0x0a, 0xcb, 0xa5, 0x0a,
	0x3d, 0x40, 0xfb, 0x42, 0x7d, 0x86, 0x3e, 0x40, 0x6f, 0xfb, 0x34, 0xc5, 0x2e, 0xc5, 0xc4, 0x6a,
	0xdc, 0x2b, 0xed, 0x7c, 0xfb, 0xcd, 0xdf, 0x37, 0x43, 0x91, 0xf0, 0x7c, 0xc6, 0xa3, 0xbb, 0x88,
	0xcf, 0x66, 0x3c, 0x5b, 0xfd, 0x0c, 0xe6, 0x82, 0x4b, 0x4e, 0x1b, 0x33, 0x1e, 0xbd, 0x38, 0x8a,
	0x39, 0x8f, 0x53, 0x3c, 0xd1, 0xd0, 0xb8, 0xb8, 0x3f, 0x99, 0x60, 0x1e, 0x89, 0x64, 0x2e, 0xb9,
	0x28, 0x69, 0xbd, 0x6f, 0x61, 0xc3, 0x12, 0x82, 0x0b, 0xda, 0x85, 0xcd, 0x6b, 0xcc, 0x73, 0x16,
	0x63, 0xb7, 0x76, 0x54, 0x3b, 0x6e, 0xfb, 0x95, 0x49, 0x29, 0x34, 0x87, 0x7c, 0x82, 0xdd, 0xfa,
	0x51, 0xed, 0x78, 0xc3, 0xd7, 0xe7, 0xde, 0xaf, 0x35, 0xd8, 0x1b, 0x09, 0xbe, 0x48, 0xf2, 0x84,
	0x67, 0x81, 0x64, 0xb2, 0xc8, 0xe9, 0x77, 0xb0, 0x1d, 0x15, 0x42, 0x60, 0x26, 0x15, 0x50, 0x86,
	0xd9, 0x3d, 0xdd, 0x1f, 0xcc, 0x78, 0x34, 0x58, 0xe3, 0xa2, 0xbf, 0x46, 0xa4, 0xdf, 0xc3, 0xce,
	0x5c, 0xe0, 0x22, 0xe1, 0x45, 0x5e, 0x7a, 0xd6, 0xff, 0xdd, 0x73, 0x9d, 0xd9, 0x5b, 0x40, 0xeb,
	0x12, 0x59, 0x2a, 0xa7, 0xf4, 0x9b, 0x47, 0xb3, 0x13, 0x1d, 0xa3, 0xa4, 0x3c, 0x96, 0xfa, 0xdd,
	0xe3, 0xa9, 0x3f, 0x77, 0xfb, 0x47, 0xde, 0xd7, 0xb0, 0x79, 0x8b, 0x42, 0x95, 0x45, 0x0f, 0xa1,
	0x95, 0x15, 0xb3, 0x31, 0x8a, 0x95, 0x6e, 0x2b, 0xab, 0xf7, 0x47, 0x0d, 0x5a, 0x2b, 0x65, 0xde,
	0x40, 0x6b, 0xaa, 0x63, 0x69, 0x4a, 0xe7, 0xb4, 0xf3, 0x20, 0xbc, 0xbf, 0xba, 0xa2, 0x26, 0xd0,
	0x79, 0xd5, 0x6b, 0x92, 0xc5, 0xa5, 0xab, 0xae, 0xa7, 0x73, 0x7a, 0xf0, 0xb9, 0x14, 0x45, 0xee,
	0x3f, 0xc2, 0xa7, 0xc7, 0xd0, 0x4e, 0x59, 0x2e, 0xf5, 0x4c, 0xbb, 0x0d, 0xed, 0x0c, 0xda, 0x59,
	0x23, 0xfe, 0xa7, 0x4b, 0xfa, 0x25, 0x6c, 0x2e, 0xca, 0x16, 0xba, 0x4d, 0xcd, 0xdb, 0xd6, 0xbc,
	0x55, 0x5b, 0x7e, 0x75, 0xd9, 0x1b, 0x40, 0xcb, 0xca, 0x64, 0x22, 0x97, 0xf4, 0x2d, 0xec, 0xd8,
	0xf9, 0x28, 0x65, 0x11, 0x4e, 0x79, 0x3a, 0x59, 0x35, 0xbc, 0xe5, 0xaf, 0x83, 0xbd, 0xaf, 0xa0,
	0x11, 0xb2, 0x98, 0x12, 0x68, 0x7c, 0xc0, 0xe5, 0x4a, 0x13, 0x75, 0xa4, 0x07, 0xb0, 0xb1, 0x60,
	0x69, 0x51, 0x6a, 0xdc, 0xf6, 0x4b, 0xa3, 0xf7, 0x16, 0x9a, 0x21, 0x8b, 0x73, 0xfa, 0x12, 0x9a,
	0x92, 0xc5, 0x79, 0xb7, 0x76, 0xd4, 0x38, 0xee, 0x9c, 0x6e, 0xe9, 0x5a, 0x42, 0x16, 0xfb, 0x1a,
	0xed, 0xbf, 0x83, 0xb6, 0x37, 0x47, 0xc1, 0xa4, 0x52, 0x7c, 0x13, 0x1a, 0x17, 0x56, 0x48, 0x9e,
	0xd0, 0x2d, 0x68, 0x8e, 0xbc, 0x20, 0x24, 0x35, 0x0a, 0xd0, 0x32, 0x2d, 0xc7, 0x0a, 0x2d, 0x52,
	0x57, 0xe7, 0x9b, 0x91, 0x69, 0x84, 0x16, 0x69, 0xf4, 0x7f, 0xaf, 0xc3, 0xee, 0xfa, 0x06, 0xd1,
	0x0e, 0x6c, 0xde, 0xb8, 0x57, 0xae, 0xf7, 0xde, 0x25, 0x4f, 0xe8, 0x36, 0x6c, 0x0d, 0x7d, 0xcb,
	0x08, 0x6d, 0xf7, 0x82, 0xd4, 0xd4, 0x95, 0xb6, 0x2c, 0x93, 0xd4, 0xe9, 0x53, 0xd8, 0x29, 0x8d,
	0xbb, 0x73, 0xc3, 0x76, 0x2c, 0x93, 0x34, 0x14, 0x5b, 0x67, 0x51, 0xec, 0xa6, 0x22, 0x94, 0x39,
	0x2b, 0xc2, 0x86, 0x0a, 0x50, 0x42, 0x26, 0x69, 0x29, 0xb6, 0xae, 0x43, 0xb1, 0x37, 0x15, 0xbb,
	0xac, 0xaa, 0x62, 0x6f, 0xe9, 0x4a, 0x34, 0x64, 0x92, 0x36, 0x25, 0xb0, 0x3d, 0xf2, 0xbd, 0x5b,
	0x3b, 0xb0, 0x3d, 0x57, 0x79, 0x00, 0xdd, 0x83, 0xce, 0x47, 0xc4, 0x32, 0x49, 0x87, 0x1e, 0x00,
	0xf9, 0x08, 0x54, 0x51, 0xb6, 0x29, 0x85, 0x5d, 0xd3, 0x5a, 0x73, 0xdd, 0x29, 0x4b, 0x7b, 0xe8,
	0xbc, 0x4b, 0x0f, 0x81, 0x3e, 0x80, 0x2a, 0xf7, 0xbd, 0xd2, 0x5d, 0x77, 0x31, 0xb2, 0x5c, 0x53,
	0xb9, 0x93, 0xfe, 0x39, 0x3c, 0xbb, 0x4c, 0xe2, 0xa9, 0xb1, 0x60, 0x49, 0xca, 0xc6, 0x49, 0x9a,
	0xc8, 0x65, 0xa9, 0xdd, 0x01, 0x90, 0x95, 0x76, 0x77, 0x97, 0xc6, 0x5d, 0x10, 0x2a, 0x91, 0x9f,
	0x28, 0xc1, 0x83, 0xd0, 0x38, 0x73, 0xac, 0x52, 0xc2, 0x2a, 0x4e, 0xbd, 0x1f, 0x43, 0xe7, 0xc1,
	0x33, 0xa4, 0x04, 0x71, 0xbd, 0xb0, 0x92, 0xbe, 0x05, 0x75, 0xef, 0xaa, 0xf4, 0x78, 0x6f, 0xf8,
	0xba, 0xf0, 0x7a, 0x39, 0x0f, 0x3b, 0xb4, 0x87, 0x86, 0x43, 0x1a, 0xea, 0xea, 0xda, 0x0e, 0x82,
	0x52, 0x6e, 0x2d, 0xfe, 0x85, 0x6f, 0x98, 0x5a, 0xe9, 0x32, 0xd6, 0xb9, 0x77, 0xe3, 0x9a, 0xa4,
	0xd5, 0x9f, 0x02, 0x0c, 0xd3, 0x04, 0x33, 0x19, 0x2e, 0xe7, 0xa8, 0xa4, 0x1c, 0x7a, 0x6e, 0xe8,
	0x7b, 0xce, 0xc8, 0x31, 0x5c, 0x55, 0x21, 0x85, 0x5d, 0xeb, 0xa7, 0xd0, 0xf2, 0x5d, 0xc3, 0x19,
	0x3a, 0xb6, 0xe5, 0xaa, 0x95, 0xd9, 0x82, 0xa6, 0xeb, 0x99, 0x6a, 0x61, 0xda, 0xb0, 0x61, 0x98,
	0xd7, 0xb6, 0x4b, 0x1a, 0x74, 0x07, 0xda, 0x67, 0x86, 0x6f, 0x5d, 0x5b, 0xa1, 0xe1, 0x90, 0xa6,
	0x8a, 0xe4, 0x78, 0x86, 0x79, 0x66, 0x38, 0x86, 0x3b, 0xb4, 0x7c, 0xb2, 0xd1, 0x3f, 0x05, 0x6a,
	0x14, 0x72, 0x8a, 0x99, 0x4c, 0x22, 0xbd, 0x8d, 0x3a, 0xe3, 0x2e, 0x40, 0x60, 0x39, 0xe7, 0x81,
	0x7d, 0xa1, 0xc4, 0x2e, 0xd7, 0xca, 0x58, 0x59, 0xb5, 0xfe, 0x5f, 0x4d, 0xd8, 0xd6, 0x4b, 0x38,
	0x41, 0xa1, 0xe9, 0x7b, 0xd0, 0x31, 0xb2, 0x65, 0x05, 0x95, 0xf5, 0xdd, 0x26, 0x42, 0x16, 0x2c,
	0xbd, 0x66, 0xd1, 0x34, 0xc9, 0x90, 0xd4, 0xe8, 0x0b, 0x38, 0x5c, 0xc7, 0x82, 0x88, 0xa5, 0x18,
	0xa0, 0x24, 0x75, 0x5d, 0x17, 0x67, 0x93, 0x33, 0x96, 0xb2, 0x2c, 0x42, 0x41, 0x1a, 0x0f, 0x22,
	0xb8, 0x28, 0x7f, 0xe1, 0xe2, 0x03, 0x69, 0xd2, 0x7d, 0xd8, 0x5b, 0x61, 0x97, 0x4c, 0x4c, 0xcc,
	0x24, 0xff, 0x40, 0x36, 0x94, 0xeb, 0x05, 0x4b, 0x53, 0x14, 0x4b, 0x7b, 0xc6, 0x62, 0x24, 0x2d,
	0xfa, 0x1c, 0xf6, 0xd7, 0x13, 0x95, 0x17, 0x9b, 0x6a, 0xda, 0xab, 0x60, 0x76, 0x26, 0x51, 0xdc,
	0xb3, 0x08, 0xc9, 0x96, 0x2a, 0x7e, 0x88, 0x42, 0x26, 0xf7, 0x4a, 0x00, 0x24, 0x6d, 0xf5, 0x38,
	0x5e, 0xe1, 0x92, 0x80, 0xde, 0x03, 0x8c, 0x04, 0x4a, 0xd2, 0x51, 0x0a, 0x5c, 0xe1, 0xf2, 0x96,
	0x15, 0xa9, 0x24, 0xdb, 0xca, 0xb2, 0x27, 0xa8, 0xff, 0x45, 0xc8, 0x8e, 0x52, 0xde, 0xe7, 0x29,
	0x92, 0x5d, 0x55, 0xb5, 0x3a, 0x19, 0x79, 0x9e, 0xc4, 0xd9, 0x0c, 0x33, 0x49, 0xf6, 0x94, 0x96,
	0x57, 0xc5, 0x18, 0x45, 0x86, 0x12, 0x73, 0x42, 0xf4, 0x43, 0x99, 0x16, 0xb9, 0x44, 0x41, 0x9e,
	0xea, 0xd1, 0xf2, 0x4c, 0x0a, 0x9e, 0x8e, 0x52, 0x96, 0x21, 0xa1, 0x6a, 0x78, 0x17, 0x82, 0x17,
	0x73, 0xb2, 0xaf, 0x27, 0xca, 0x27, 0x48, 0x0e, 0x54, 0x3e, 0x87, 0x97, 0xf3, 0x21, 0xcf, 0x54,
	0x1f, 0x81, 0xe4, 0x82, 0xc5, 0xa8, 0x7c, 0x59, 0x92, 0xa1, 0x20, 0x87, 0xaa, 0x8f, 0x15, 0x7a,
	0x9e, 0xa4, 0x48, 0x9e, 0x3f, 0xa0, 0x99, 0x89, 0xc0, 0x48, 0x72, 0xb1, 0x24, 0x5d, 0x95, 0x31,
	0x28, 0xc6, 0xe5, 0x5b, 0x53, 0x85, 0xfb, 0x8f, 0x2a, 0xe8, 0x36, 0x99, 0x8f, 0x38, 0x4f, 0xc9,
	0x0b, 0xbd, 0xa2, 0x2c, 0xd2, 0xc6, 0x7f, 0x55, 0x48, 0x4b, 0x46, 0x93, 0xaa, 0xdc, 0x97, 0xaa,
	0x17, 0x05, 0x04, 0x28, 0x16, 0x28, 0xc8, 0x2b, 0x95, 0xe2, 0x8c, 0x09, 0xbc, 0x46, 0xf9, 0x69,
	0xd2, 0xff, 0xa3, 0xcf, 0xe0, 0xe9, 0x50, 0xa0, 0xd6, 0x87, 0xa5, 0xd7, 0x3c, 0x4b, 0x24, 0x17,
	0xe4, 0xff, 0x2a, 0xb4, 0xc3, 0xe3, 0x38, 0xc9, 0x62, 0x72, 0xa4, 0x3a, 0xf2, 0x31, 0xe2, 0x0b,
	0x14, 0x4b, 0xf2, 0x5a, 0x35, 0x6d, 0xe2, 0xb8, 0x88, 0x49, 0xaf, 0xef, 0x43, 0x47, 0xcf, 0x2b,
	0xe0, 0x85, 0x88, 0xb0, 0xdc, 0xd8, 0xa1, 0xe1, 0xdc, 0x05, 0xde, 0x8d, 0x3f, 0x54, 0xbb, 0xaf,
	0x76, 0xf3, 0x3c, 0xa8, 0xec, 0x9a, 0x2a, 0xf2, 0x32, 0x0c, 0x47, 0x15, 0xa0, 0x97, 0x69, 0xe8,
	0x78, 0xae, 0x55, 0x21, 0x8d, 0x1f, 0x7e, 0x84, 0x76, 0x8e, 0x59, 0x9e, 0xc8, 0x64, 0x81, 0xf4,
	0xd5, 0xa0, 0xfc, 0x88, 0x18, 0x54, 0x1f, 0x11, 0x83, 0xf3, 0x04, 0xd3, 0x89, 0xa7, 0xe5, 0xc8,
	0xbb, 0x7f, 0xfe, 0xd6, 0xd0, 0x6f, 0x81, 0x4f, 0x1e, 0x67, 0x5f, 0xfc, 0xfc, 0x26, 0x4e, 0xe4,
	0xb4, 0x18, 0x0f, 0x22, 0x3e, 0x3b, 0x99, 0x25, 0x91, 0xe0, 0x39, 0xbf, 0x97, 0x27, 0x33, 0x1e,
	0x9d, 0x88, 0x79, 0x74, 0x52, 0x7e, 0xa7, 0x8c, 0x5b, 0x3a, 0xe0, 0xd7, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xca, 0x02, 0x8d, 0xb5, 0xc3, 0x08, 0x00, 0x00,
}
