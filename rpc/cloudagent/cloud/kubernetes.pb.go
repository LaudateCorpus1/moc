// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kubernetes.proto

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

type NodeType int32

const (
	NodeType_ControlPlane  NodeType = 0
	NodeType_LinuxWorker   NodeType = 1
	NodeType_WindowsWorker NodeType = 2
	NodeType_LoadBalancer  NodeType = 3
)

var NodeType_name = map[int32]string{
	0: "ControlPlane",
	1: "LinuxWorker",
	2: "WindowsWorker",
	3: "LoadBalancer",
}

var NodeType_value = map[string]int32{
	"ControlPlane":  0,
	"LinuxWorker":   1,
	"WindowsWorker": 2,
	"LoadBalancer":  3,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}

func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{0}
}

type ManagementStrategyType int32

const (
	ManagementStrategyType_Pivoted  ManagementStrategyType = 0
	ManagementStrategyType_Distinct ManagementStrategyType = 1
)

var ManagementStrategyType_name = map[int32]string{
	0: "Pivoted",
	1: "Distinct",
}

var ManagementStrategyType_value = map[string]int32{
	"Pivoted":  0,
	"Distinct": 1,
}

func (x ManagementStrategyType) String() string {
	return proto.EnumName(ManagementStrategyType_name, int32(x))
}

func (ManagementStrategyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{1}
}

type KubernetesRequest struct {
	Kubernetess          []*Kubernetes    `protobuf:"bytes,1,rep,name=Kubernetess,proto3" json:"Kubernetess,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *KubernetesRequest) Reset()         { *m = KubernetesRequest{} }
func (m *KubernetesRequest) String() string { return proto.CompactTextString(m) }
func (*KubernetesRequest) ProtoMessage()    {}
func (*KubernetesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{0}
}

func (m *KubernetesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesRequest.Unmarshal(m, b)
}
func (m *KubernetesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesRequest.Marshal(b, m, deterministic)
}
func (m *KubernetesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesRequest.Merge(m, src)
}
func (m *KubernetesRequest) XXX_Size() int {
	return xxx_messageInfo_KubernetesRequest.Size(m)
}
func (m *KubernetesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesRequest proto.InternalMessageInfo

func (m *KubernetesRequest) GetKubernetess() []*Kubernetes {
	if m != nil {
		return m.Kubernetess
	}
	return nil
}

func (m *KubernetesRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type KubernetesResponse struct {
	Kubernetess          []*Kubernetes       `protobuf:"bytes,1,rep,name=Kubernetess,proto3" json:"Kubernetess,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *KubernetesResponse) Reset()         { *m = KubernetesResponse{} }
func (m *KubernetesResponse) String() string { return proto.CompactTextString(m) }
func (*KubernetesResponse) ProtoMessage()    {}
func (*KubernetesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{1}
}

func (m *KubernetesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesResponse.Unmarshal(m, b)
}
func (m *KubernetesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesResponse.Marshal(b, m, deterministic)
}
func (m *KubernetesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesResponse.Merge(m, src)
}
func (m *KubernetesResponse) XXX_Size() int {
	return xxx_messageInfo_KubernetesResponse.Size(m)
}
func (m *KubernetesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesResponse proto.InternalMessageInfo

func (m *KubernetesResponse) GetKubernetess() []*Kubernetes {
	if m != nil {
		return m.Kubernetess
	}
	return nil
}

func (m *KubernetesResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *KubernetesResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type StorageConfiguration struct {
	Csi                  string   `protobuf:"bytes,1,opt,name=csi,proto3" json:"csi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StorageConfiguration) Reset()         { *m = StorageConfiguration{} }
func (m *StorageConfiguration) String() string { return proto.CompactTextString(m) }
func (*StorageConfiguration) ProtoMessage()    {}
func (*StorageConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{2}
}

func (m *StorageConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageConfiguration.Unmarshal(m, b)
}
func (m *StorageConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageConfiguration.Marshal(b, m, deterministic)
}
func (m *StorageConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageConfiguration.Merge(m, src)
}
func (m *StorageConfiguration) XXX_Size() int {
	return xxx_messageInfo_StorageConfiguration.Size(m)
}
func (m *StorageConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_StorageConfiguration proto.InternalMessageInfo

func (m *StorageConfiguration) GetCsi() string {
	if m != nil {
		return m.Csi
	}
	return ""
}

type NetworkConfiguration struct {
	Cni         string `protobuf:"bytes,1,opt,name=cni,proto3" json:"cni,omitempty"`
	PodCidr     string `protobuf:"bytes,2,opt,name=podCidr,proto3" json:"podCidr,omitempty"`
	ClusterCidr string `protobuf:"bytes,3,opt,name=clusterCidr,proto3" json:"clusterCidr,omitempty"`
	// TODO: merge controlplane cidr and network
	ControlPlaneCidr     string   `protobuf:"bytes,4,opt,name=controlPlaneCidr,proto3" json:"controlPlaneCidr,omitempty"`
	Virtualnetwork       string   `protobuf:"bytes,5,opt,name=virtualnetwork,proto3" json:"virtualnetwork,omitempty"`
	LoadBalancerVip      string   `protobuf:"bytes,6,opt,name=loadBalancerVip,proto3" json:"loadBalancerVip,omitempty"`
	LoadBalancerMac      string   `protobuf:"bytes,7,opt,name=loadBalancerMac,proto3" json:"loadBalancerMac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkConfiguration) Reset()         { *m = NetworkConfiguration{} }
func (m *NetworkConfiguration) String() string { return proto.CompactTextString(m) }
func (*NetworkConfiguration) ProtoMessage()    {}
func (*NetworkConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{3}
}

func (m *NetworkConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkConfiguration.Unmarshal(m, b)
}
func (m *NetworkConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkConfiguration.Marshal(b, m, deterministic)
}
func (m *NetworkConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkConfiguration.Merge(m, src)
}
func (m *NetworkConfiguration) XXX_Size() int {
	return xxx_messageInfo_NetworkConfiguration.Size(m)
}
func (m *NetworkConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkConfiguration proto.InternalMessageInfo

func (m *NetworkConfiguration) GetCni() string {
	if m != nil {
		return m.Cni
	}
	return ""
}

func (m *NetworkConfiguration) GetPodCidr() string {
	if m != nil {
		return m.PodCidr
	}
	return ""
}

func (m *NetworkConfiguration) GetClusterCidr() string {
	if m != nil {
		return m.ClusterCidr
	}
	return ""
}

func (m *NetworkConfiguration) GetControlPlaneCidr() string {
	if m != nil {
		return m.ControlPlaneCidr
	}
	return ""
}

func (m *NetworkConfiguration) GetVirtualnetwork() string {
	if m != nil {
		return m.Virtualnetwork
	}
	return ""
}

func (m *NetworkConfiguration) GetLoadBalancerVip() string {
	if m != nil {
		return m.LoadBalancerVip
	}
	return ""
}

func (m *NetworkConfiguration) GetLoadBalancerMac() string {
	if m != nil {
		return m.LoadBalancerMac
	}
	return ""
}

type NodePoolConfiguration struct {
	NodeType             NodeType `protobuf:"varint,1,opt,name=NodeType,proto3,enum=moc.cloudagent.kubernetes.NodeType" json:"NodeType,omitempty"`
	Imagereference       string   `protobuf:"bytes,2,opt,name=imagereference,proto3" json:"imagereference,omitempty"`
	Replicas             int32    `protobuf:"varint,3,opt,name=replicas,proto3" json:"replicas,omitempty"`
	VMSize               string   `protobuf:"bytes,4,opt,name=VMSize,proto3" json:"VMSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodePoolConfiguration) Reset()         { *m = NodePoolConfiguration{} }
func (m *NodePoolConfiguration) String() string { return proto.CompactTextString(m) }
func (*NodePoolConfiguration) ProtoMessage()    {}
func (*NodePoolConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{4}
}

func (m *NodePoolConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePoolConfiguration.Unmarshal(m, b)
}
func (m *NodePoolConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePoolConfiguration.Marshal(b, m, deterministic)
}
func (m *NodePoolConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePoolConfiguration.Merge(m, src)
}
func (m *NodePoolConfiguration) XXX_Size() int {
	return xxx_messageInfo_NodePoolConfiguration.Size(m)
}
func (m *NodePoolConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePoolConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_NodePoolConfiguration proto.InternalMessageInfo

func (m *NodePoolConfiguration) GetNodeType() NodeType {
	if m != nil {
		return m.NodeType
	}
	return NodeType_ControlPlane
}

func (m *NodePoolConfiguration) GetImagereference() string {
	if m != nil {
		return m.Imagereference
	}
	return ""
}

func (m *NodePoolConfiguration) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *NodePoolConfiguration) GetVMSize() string {
	if m != nil {
		return m.VMSize
	}
	return ""
}

type SSHPublicKey struct {
	KeyData              string   `protobuf:"bytes,1,opt,name=keyData,proto3" json:"keyData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSHPublicKey) Reset()         { *m = SSHPublicKey{} }
func (m *SSHPublicKey) String() string { return proto.CompactTextString(m) }
func (*SSHPublicKey) ProtoMessage()    {}
func (*SSHPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{5}
}

func (m *SSHPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSHPublicKey.Unmarshal(m, b)
}
func (m *SSHPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSHPublicKey.Marshal(b, m, deterministic)
}
func (m *SSHPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSHPublicKey.Merge(m, src)
}
func (m *SSHPublicKey) XXX_Size() int {
	return xxx_messageInfo_SSHPublicKey.Size(m)
}
func (m *SSHPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SSHPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_SSHPublicKey proto.InternalMessageInfo

func (m *SSHPublicKey) GetKeyData() string {
	if m != nil {
		return m.KeyData
	}
	return ""
}

type ComputeConfiguration struct {
	Cri                  string                   `protobuf:"bytes,1,opt,name=cri,proto3" json:"cri,omitempty"`
	PublicKey            *SSHPublicKey            `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	NodePools            []*NodePoolConfiguration `protobuf:"bytes,3,rep,name=NodePools,proto3" json:"NodePools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ComputeConfiguration) Reset()         { *m = ComputeConfiguration{} }
func (m *ComputeConfiguration) String() string { return proto.CompactTextString(m) }
func (*ComputeConfiguration) ProtoMessage()    {}
func (*ComputeConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{6}
}

func (m *ComputeConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeConfiguration.Unmarshal(m, b)
}
func (m *ComputeConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeConfiguration.Marshal(b, m, deterministic)
}
func (m *ComputeConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeConfiguration.Merge(m, src)
}
func (m *ComputeConfiguration) XXX_Size() int {
	return xxx_messageInfo_ComputeConfiguration.Size(m)
}
func (m *ComputeConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeConfiguration proto.InternalMessageInfo

func (m *ComputeConfiguration) GetCri() string {
	if m != nil {
		return m.Cri
	}
	return ""
}

func (m *ComputeConfiguration) GetPublicKey() *SSHPublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *ComputeConfiguration) GetNodePools() []*NodePoolConfiguration {
	if m != nil {
		return m.NodePools
	}
	return nil
}

type ClusterConfiguration struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterConfiguration) Reset()         { *m = ClusterConfiguration{} }
func (m *ClusterConfiguration) String() string { return proto.CompactTextString(m) }
func (*ClusterConfiguration) ProtoMessage()    {}
func (*ClusterConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{7}
}

func (m *ClusterConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterConfiguration.Unmarshal(m, b)
}
func (m *ClusterConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterConfiguration.Marshal(b, m, deterministic)
}
func (m *ClusterConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterConfiguration.Merge(m, src)
}
func (m *ClusterConfiguration) XXX_Size() int {
	return xxx_messageInfo_ClusterConfiguration.Size(m)
}
func (m *ClusterConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterConfiguration proto.InternalMessageInfo

func (m *ClusterConfiguration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ClusterAPIConfiguration struct {
	ConfigurationEndpoint         string   `protobuf:"bytes,1,opt,name=ConfigurationEndpoint,proto3" json:"ConfigurationEndpoint,omitempty"`
	InfrastructureProviderVersion string   `protobuf:"bytes,2,opt,name=InfrastructureProviderVersion,proto3" json:"InfrastructureProviderVersion,omitempty"`
	BootstrapProviderVersion      string   `protobuf:"bytes,3,opt,name=BootstrapProviderVersion,proto3" json:"BootstrapProviderVersion,omitempty"`
	ControlPlaneProviderVersion   string   `protobuf:"bytes,4,opt,name=ControlPlaneProviderVersion,proto3" json:"ControlPlaneProviderVersion,omitempty"`
	CoreProviderVersion           string   `protobuf:"bytes,5,opt,name=CoreProviderVersion,proto3" json:"CoreProviderVersion,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *ClusterAPIConfiguration) Reset()         { *m = ClusterAPIConfiguration{} }
func (m *ClusterAPIConfiguration) String() string { return proto.CompactTextString(m) }
func (*ClusterAPIConfiguration) ProtoMessage()    {}
func (*ClusterAPIConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{8}
}

func (m *ClusterAPIConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterAPIConfiguration.Unmarshal(m, b)
}
func (m *ClusterAPIConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterAPIConfiguration.Marshal(b, m, deterministic)
}
func (m *ClusterAPIConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterAPIConfiguration.Merge(m, src)
}
func (m *ClusterAPIConfiguration) XXX_Size() int {
	return xxx_messageInfo_ClusterAPIConfiguration.Size(m)
}
func (m *ClusterAPIConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterAPIConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterAPIConfiguration proto.InternalMessageInfo

func (m *ClusterAPIConfiguration) GetConfigurationEndpoint() string {
	if m != nil {
		return m.ConfigurationEndpoint
	}
	return ""
}

func (m *ClusterAPIConfiguration) GetInfrastructureProviderVersion() string {
	if m != nil {
		return m.InfrastructureProviderVersion
	}
	return ""
}

func (m *ClusterAPIConfiguration) GetBootstrapProviderVersion() string {
	if m != nil {
		return m.BootstrapProviderVersion
	}
	return ""
}

func (m *ClusterAPIConfiguration) GetControlPlaneProviderVersion() string {
	if m != nil {
		return m.ControlPlaneProviderVersion
	}
	return ""
}

func (m *ClusterAPIConfiguration) GetCoreProviderVersion() string {
	if m != nil {
		return m.CoreProviderVersion
	}
	return ""
}

type ContainerRegistry struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerRegistry) Reset()         { *m = ContainerRegistry{} }
func (m *ContainerRegistry) String() string { return proto.CompactTextString(m) }
func (*ContainerRegistry) ProtoMessage()    {}
func (*ContainerRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{9}
}

func (m *ContainerRegistry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerRegistry.Unmarshal(m, b)
}
func (m *ContainerRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerRegistry.Marshal(b, m, deterministic)
}
func (m *ContainerRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerRegistry.Merge(m, src)
}
func (m *ContainerRegistry) XXX_Size() int {
	return xxx_messageInfo_ContainerRegistry.Size(m)
}
func (m *ContainerRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerRegistry proto.InternalMessageInfo

func (m *ContainerRegistry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerRegistry) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ContainerRegistry) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Kubernetes struct {
	Name                 string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Status               *common.Status           `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Cluster              *ClusterConfiguration    `protobuf:"bytes,5,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Network              *NetworkConfiguration    `protobuf:"bytes,6,opt,name=network,proto3" json:"network,omitempty"`
	Storage              *StorageConfiguration    `protobuf:"bytes,7,opt,name=storage,proto3" json:"storage,omitempty"`
	Compute              *ComputeConfiguration    `protobuf:"bytes,8,opt,name=compute,proto3" json:"compute,omitempty"`
	GroupName            string                   `protobuf:"bytes,9,opt,name=groupName,proto3" json:"groupName,omitempty"`
	ManagementStrategy   ManagementStrategyType   `protobuf:"varint,10,opt,name=managementStrategy,proto3,enum=moc.cloudagent.kubernetes.ManagementStrategyType" json:"managementStrategy,omitempty"`
	LocationName         string                   `protobuf:"bytes,11,opt,name=locationName,proto3" json:"locationName,omitempty"`
	KubeConfig           []byte                   `protobuf:"bytes,12,opt,name=kubeConfig,proto3" json:"kubeConfig,omitempty"`
	CapiConfig           *ClusterAPIConfiguration `protobuf:"bytes,13,opt,name=capiConfig,proto3" json:"capiConfig,omitempty"`
	ContainerRegistry    *ContainerRegistry       `protobuf:"bytes,14,opt,name=containerRegistry,proto3" json:"containerRegistry,omitempty"`
	DeploymentManifest   []byte                   `protobuf:"bytes,15,opt,name=deploymentManifest,proto3" json:"deploymentManifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Kubernetes) Reset()         { *m = Kubernetes{} }
func (m *Kubernetes) String() string { return proto.CompactTextString(m) }
func (*Kubernetes) ProtoMessage()    {}
func (*Kubernetes) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{10}
}

func (m *Kubernetes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Kubernetes.Unmarshal(m, b)
}
func (m *Kubernetes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Kubernetes.Marshal(b, m, deterministic)
}
func (m *Kubernetes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Kubernetes.Merge(m, src)
}
func (m *Kubernetes) XXX_Size() int {
	return xxx_messageInfo_Kubernetes.Size(m)
}
func (m *Kubernetes) XXX_DiscardUnknown() {
	xxx_messageInfo_Kubernetes.DiscardUnknown(m)
}

var xxx_messageInfo_Kubernetes proto.InternalMessageInfo

func (m *Kubernetes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Kubernetes) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Kubernetes) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Kubernetes) GetCluster() *ClusterConfiguration {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *Kubernetes) GetNetwork() *NetworkConfiguration {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Kubernetes) GetStorage() *StorageConfiguration {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *Kubernetes) GetCompute() *ComputeConfiguration {
	if m != nil {
		return m.Compute
	}
	return nil
}

func (m *Kubernetes) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *Kubernetes) GetManagementStrategy() ManagementStrategyType {
	if m != nil {
		return m.ManagementStrategy
	}
	return ManagementStrategyType_Pivoted
}

func (m *Kubernetes) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *Kubernetes) GetKubeConfig() []byte {
	if m != nil {
		return m.KubeConfig
	}
	return nil
}

func (m *Kubernetes) GetCapiConfig() *ClusterAPIConfiguration {
	if m != nil {
		return m.CapiConfig
	}
	return nil
}

func (m *Kubernetes) GetContainerRegistry() *ContainerRegistry {
	if m != nil {
		return m.ContainerRegistry
	}
	return nil
}

func (m *Kubernetes) GetDeploymentManifest() []byte {
	if m != nil {
		return m.DeploymentManifest
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.cloudagent.kubernetes.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("moc.cloudagent.kubernetes.ManagementStrategyType", ManagementStrategyType_name, ManagementStrategyType_value)
	proto.RegisterType((*KubernetesRequest)(nil), "moc.cloudagent.kubernetes.KubernetesRequest")
	proto.RegisterType((*KubernetesResponse)(nil), "moc.cloudagent.kubernetes.KubernetesResponse")
	proto.RegisterType((*StorageConfiguration)(nil), "moc.cloudagent.kubernetes.StorageConfiguration")
	proto.RegisterType((*NetworkConfiguration)(nil), "moc.cloudagent.kubernetes.NetworkConfiguration")
	proto.RegisterType((*NodePoolConfiguration)(nil), "moc.cloudagent.kubernetes.NodePoolConfiguration")
	proto.RegisterType((*SSHPublicKey)(nil), "moc.cloudagent.kubernetes.SSHPublicKey")
	proto.RegisterType((*ComputeConfiguration)(nil), "moc.cloudagent.kubernetes.ComputeConfiguration")
	proto.RegisterType((*ClusterConfiguration)(nil), "moc.cloudagent.kubernetes.ClusterConfiguration")
	proto.RegisterType((*ClusterAPIConfiguration)(nil), "moc.cloudagent.kubernetes.ClusterAPIConfiguration")
	proto.RegisterType((*ContainerRegistry)(nil), "moc.cloudagent.kubernetes.ContainerRegistry")
	proto.RegisterType((*Kubernetes)(nil), "moc.cloudagent.kubernetes.Kubernetes")
}

func init() { proto.RegisterFile("kubernetes.proto", fileDescriptor_40204d9320c6ada8) }

var fileDescriptor_40204d9320c6ada8 = []byte{
	// 1060 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x6e, 0xdb, 0xc6,
	0x13, 0x36, 0xe5, 0x58, 0xb6, 0x46, 0xb2, 0x2d, 0xef, 0xcf, 0xc9, 0x8f, 0x75, 0xdb, 0xc0, 0x60,
	0xd0, 0x56, 0x30, 0x12, 0xc9, 0x55, 0x72, 0xea, 0xa5, 0x8d, 0x65, 0xa3, 0x35, 0x12, 0xbb, 0x02,
	0xd5, 0x3a, 0x40, 0x2e, 0xc5, 0x9a, 0x1c, 0xb1, 0x0b, 0x93, 0xbb, 0xec, 0xee, 0xd2, 0xae, 0xf2,
	0x0c, 0x3d, 0xf5, 0x21, 0xda, 0x17, 0xe8, 0xbd, 0xa7, 0xbe, 0x57, 0xc1, 0x25, 0x29, 0x52, 0x7f,
	0x22, 0xe7, 0xd0, 0x1b, 0x67, 0xe6, 0xfb, 0x3e, 0xcd, 0xcc, 0xce, 0x8e, 0x16, 0xda, 0x37, 0xc9,
	0x35, 0x4a, 0x8e, 0x1a, 0x55, 0x37, 0x96, 0x42, 0x0b, 0xf2, 0x51, 0x24, 0xbc, 0xae, 0x17, 0x8a,
	0xc4, 0xa7, 0x01, 0x72, 0xdd, 0x2d, 0x01, 0x07, 0x8f, 0x03, 0x21, 0x82, 0x10, 0x7b, 0x06, 0x78,
	0x9d, 0x8c, 0x7b, 0x77, 0x92, 0xc6, 0x31, 0xca, 0x9c, 0x7a, 0xd0, 0xf2, 0x44, 0x14, 0x09, 0x9e,
	0x59, 0xce, 0xef, 0x16, 0xec, 0xbd, 0x9a, 0x92, 0x5d, 0xfc, 0x25, 0x41, 0xa5, 0xc9, 0xb7, 0xd0,
	0x2c, 0x9d, 0xca, 0xb6, 0x0e, 0xd7, 0x3b, 0xcd, 0xfe, 0x67, 0xdd, 0xf7, 0xfe, 0x68, 0xb7, 0x22,
	0x51, 0x65, 0x92, 0x17, 0xb0, 0xfd, 0x7d, 0x8c, 0x92, 0x6a, 0x26, 0xf8, 0x0f, 0x93, 0x18, 0xed,
	0xda, 0xa1, 0xd5, 0xd9, 0xe9, 0xef, 0x18, 0xa9, 0x69, 0xc4, 0x9d, 0x05, 0x39, 0x7f, 0x5a, 0x40,
	0xaa, 0x49, 0xa9, 0x58, 0x70, 0x85, 0xff, 0x5d, 0x56, 0x7d, 0xa8, 0xbb, 0xa8, 0x92, 0x50, 0x9b,
	0x74, 0x9a, 0xfd, 0x83, 0x6e, 0xd6, 0xb3, 0x6e, 0xd1, 0xb3, 0xee, 0x89, 0x10, 0xe1, 0x15, 0x0d,
	0x13, 0x74, 0x73, 0x24, 0xd9, 0x87, 0x8d, 0x33, 0x29, 0x85, 0xb4, 0xd7, 0x0f, 0xad, 0x4e, 0xc3,
	0xcd, 0x0c, 0xa7, 0x03, 0xfb, 0x23, 0x2d, 0x24, 0x0d, 0x70, 0x20, 0xf8, 0x98, 0x05, 0x49, 0x56,
	0x05, 0x69, 0xc3, 0xba, 0xa7, 0x98, 0x6d, 0x19, 0x6c, 0xfa, 0xe9, 0xfc, 0x56, 0x83, 0xfd, 0x4b,
	0xd4, 0x77, 0x42, 0xde, 0x2c, 0x42, 0x79, 0x09, 0xe5, 0x8c, 0xd8, 0xb0, 0x19, 0x0b, 0x7f, 0xc0,
	0x7c, 0x69, 0xf2, 0x6b, 0xb8, 0x85, 0x49, 0x0e, 0xa1, 0xe9, 0x85, 0x89, 0xd2, 0x28, 0x4d, 0x34,
	0x4b, 0xa5, 0xea, 0x22, 0x47, 0xd0, 0xf6, 0x04, 0xd7, 0x52, 0x84, 0xc3, 0x90, 0x72, 0x34, 0xb0,
	0x07, 0x06, 0xb6, 0xe0, 0x27, 0x9f, 0xc3, 0xce, 0x2d, 0x93, 0x3a, 0xa1, 0x21, 0xcf, 0x12, 0xb3,
	0x37, 0x0c, 0x72, 0xce, 0x4b, 0x3a, 0xb0, 0x1b, 0x0a, 0xea, 0x9f, 0xd0, 0x90, 0x72, 0x0f, 0xe5,
	0x15, 0x8b, 0xed, 0xba, 0x01, 0xce, 0xbb, 0xe7, 0x91, 0x17, 0xd4, 0xb3, 0x37, 0x17, 0x91, 0x17,
	0xd4, 0x73, 0xfe, 0xb2, 0xe0, 0xe1, 0xa5, 0xf0, 0x71, 0x28, 0x44, 0x38, 0xdb, 0x8f, 0xaf, 0x61,
	0x2b, 0x0d, 0x98, 0x69, 0xb1, 0xcc, 0xb4, 0x3c, 0x59, 0x71, 0xc4, 0x05, 0xd4, 0x9d, 0x92, 0xd2,
	0xb2, 0x58, 0x44, 0x03, 0x94, 0x38, 0x46, 0x89, 0xdc, 0xc3, 0xbc, 0x8b, 0x73, 0x5e, 0x72, 0x00,
	0x5b, 0x12, 0xe3, 0x90, 0x79, 0x54, 0x99, 0x4e, 0x6e, 0xb8, 0x53, 0x9b, 0x3c, 0x82, 0xfa, 0xd5,
	0xc5, 0x88, 0xbd, 0xc3, 0xbc, 0x79, 0xb9, 0xe5, 0x74, 0xa0, 0x35, 0x1a, 0x7d, 0x37, 0x4c, 0xae,
	0x43, 0xe6, 0xbd, 0xc2, 0x49, 0x7a, 0x54, 0x37, 0x38, 0x39, 0xa5, 0x9a, 0xe6, 0x07, 0x58, 0x98,
	0xce, 0xdf, 0x16, 0xec, 0x0f, 0x44, 0x14, 0x27, 0x7a, 0xc9, 0x68, 0xc8, 0xf2, 0xbc, 0x25, 0x23,
	0x67, 0xd0, 0x88, 0x0b, 0xc5, 0x7c, 0x22, 0xbf, 0x58, 0x51, 0x72, 0x35, 0x01, 0xb7, 0x64, 0x92,
	0x4b, 0x68, 0x14, 0x1d, 0x4d, 0x0b, 0x4a, 0x2f, 0xc7, 0xf1, 0x3d, 0x9d, 0x5b, 0xe8, 0xbe, 0x5b,
	0x4a, 0x38, 0xc7, 0xb0, 0x3f, 0xc8, 0x27, 0x6b, 0xa6, 0x00, 0x1b, 0x36, 0x6f, 0x51, 0x2a, 0x26,
	0x78, 0x51, 0x73, 0x6e, 0x3a, 0xff, 0xd4, 0xe0, 0xff, 0x39, 0xe5, 0xe5, 0xf0, 0x7c, 0x96, 0xf5,
	0x02, 0x1e, 0xce, 0x38, 0xce, 0xb8, 0x1f, 0x0b, 0xc6, 0x75, 0xae, 0xb1, 0x3c, 0x48, 0x4e, 0xe1,
	0xd3, 0x73, 0x3e, 0x96, 0x54, 0x69, 0x99, 0x78, 0x3a, 0x91, 0x38, 0x94, 0xe2, 0x96, 0xf9, 0x28,
	0xaf, 0xf2, 0x0c, 0xb2, 0xa3, 0x5d, 0x0d, 0x22, 0x5f, 0x81, 0x7d, 0x22, 0x84, 0x56, 0x5a, 0xd2,
	0x78, 0x5e, 0x20, 0xbb, 0x43, 0xef, 0x8d, 0x93, 0x6f, 0xe0, 0xe3, 0x41, 0xe5, 0xe2, 0xcc, 0xd3,
	0xb3, 0xf1, 0x58, 0x05, 0x21, 0xc7, 0xf0, 0xbf, 0x81, 0x58, 0xcc, 0x3c, 0xbb, 0x6b, 0xcb, 0x42,
	0xce, 0x4f, 0xb0, 0x97, 0x0a, 0x52, 0xc6, 0x51, 0xba, 0x18, 0x30, 0xa5, 0xe5, 0x84, 0x10, 0x78,
	0x70, 0x49, 0x23, 0xcc, 0xfb, 0x65, 0xbe, 0xd3, 0x11, 0xfe, 0x51, 0xa1, 0xe4, 0xa9, 0x3f, 0xeb,
	0xc4, 0xd4, 0x4e, 0x63, 0x43, 0xaa, 0xd4, 0x9d, 0x90, 0x7e, 0x5e, 0xe4, 0xd4, 0x76, 0xfe, 0xa8,
	0x03, 0x94, 0x0b, 0x31, 0x95, 0xe6, 0x15, 0x69, 0x43, 0xdf, 0x81, 0x1a, 0xf3, 0x73, 0xd1, 0x1a,
	0xf3, 0xc9, 0x13, 0xa8, 0x2b, 0x4d, 0x75, 0xa2, 0x4c, 0xc9, 0xcd, 0x7e, 0xd3, 0x8c, 0xd6, 0xc8,
	0xb8, 0xdc, 0x3c, 0x44, 0xce, 0x61, 0x33, 0x5f, 0x46, 0xa6, 0xbc, 0x66, 0xbf, 0xb7, 0x62, 0x00,
	0x97, 0x0d, 0x97, 0x5b, 0xf0, 0x53, 0xa9, 0x62, 0x2b, 0xd5, 0xef, 0x95, 0x5a, 0xb6, 0x58, 0xdd,
	0x82, 0x9f, 0x4a, 0xa9, 0x6c, 0x49, 0x9b, 0x6d, 0xb4, 0x5a, 0x6a, 0xd9, 0x3a, 0x77, 0x0b, 0xbe,
	0x29, 0x30, 0xbb, 0xd4, 0xf6, 0xd6, 0xfd, 0x05, 0x2e, 0xb9, 0xfe, 0x6e, 0xc1, 0x27, 0x9f, 0x40,
	0x23, 0x90, 0x22, 0x89, 0xcd, 0xa1, 0x36, 0x4c, 0x9f, 0x4b, 0x07, 0xa1, 0x40, 0x22, 0xca, 0x69,
	0x80, 0x11, 0x72, 0x3d, 0xd2, 0x92, 0x6a, 0x0c, 0x26, 0x36, 0x98, 0x7d, 0xf8, 0xe5, 0x8a, 0xdf,
	0xbc, 0x58, 0x20, 0x99, 0xed, 0xb8, 0x44, 0x8c, 0x38, 0xd0, 0x0a, 0x85, 0x67, 0xb2, 0x32, 0x39,
	0x34, 0x4d, 0x0e, 0x33, 0x3e, 0xf2, 0x18, 0x20, 0x15, 0xcf, 0x4a, 0xb0, 0x5b, 0x87, 0x56, 0xa7,
	0xe5, 0x56, 0x3c, 0xc4, 0x05, 0xf0, 0x68, 0xcc, 0xf2, 0xf8, 0xb6, 0x69, 0x49, 0xff, 0xfe, 0x33,
	0x9f, 0xdf, 0x0e, 0x6e, 0x45, 0x85, 0xbc, 0x85, 0x3d, 0x6f, 0x7e, 0xfa, 0xed, 0x1d, 0x23, 0xfd,
	0x74, 0x65, 0xb7, 0xe7, 0x38, 0xee, 0xa2, 0x0c, 0xe9, 0x02, 0xf1, 0x31, 0x0e, 0xc5, 0x24, 0xed,
	0xc4, 0x05, 0xe5, 0x6c, 0x8c, 0x4a, 0xdb, 0xbb, 0xa6, 0xae, 0x25, 0x91, 0x23, 0xb7, 0xfc, 0x33,
	0x22, 0x6d, 0x68, 0x55, 0xaf, 0x79, 0x7b, 0x8d, 0xec, 0x42, 0xf3, 0x35, 0xe3, 0xc9, 0xaf, 0x6f,
	0x84, 0xbc, 0x41, 0xd9, 0xb6, 0xc8, 0x1e, 0x6c, 0xbf, 0x61, 0xdc, 0x17, 0x77, 0x2a, 0x77, 0xd5,
	0x52, 0xd6, 0xeb, 0xca, 0x7f, 0x5f, 0x7b, 0xfd, 0xe8, 0x39, 0x3c, 0x5a, 0x7e, 0x4a, 0xa4, 0x09,
	0x9b, 0x43, 0x76, 0x2b, 0x34, 0xfa, 0xed, 0x35, 0xd2, 0x82, 0xad, 0x53, 0xa6, 0x34, 0xe3, 0x9e,
	0x6e, 0x5b, 0xfd, 0x77, 0xb0, 0x5b, 0x5e, 0xd8, 0x97, 0x69, 0xed, 0x24, 0x80, 0xfa, 0x39, 0xbf,
	0x15, 0x37, 0x48, 0x9e, 0x7e, 0xd8, 0x1b, 0x28, 0x7b, 0xdc, 0x1d, 0x3c, 0xfb, 0x40, 0x74, 0xf6,
	0xea, 0x72, 0xd6, 0x4e, 0x7a, 0x6f, 0x9f, 0x05, 0x4c, 0xff, 0x9c, 0x5c, 0x77, 0x3d, 0x11, 0xf5,
	0x22, 0xe6, 0x49, 0xa1, 0xc4, 0x58, 0xf7, 0x22, 0xe1, 0xf5, 0x64, 0xec, 0xf5, 0x4a, 0xa9, 0xec,
	0xf3, 0xba, 0x6e, 0xde, 0x51, 0xcf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x81, 0xba, 0x1e,
	0xb8, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KubernetesAgentClient is the client API for KubernetesAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KubernetesAgentClient interface {
	Invoke(ctx context.Context, in *KubernetesRequest, opts ...grpc.CallOption) (*KubernetesResponse, error)
}

type kubernetesAgentClient struct {
	cc *grpc.ClientConn
}

func NewKubernetesAgentClient(cc *grpc.ClientConn) KubernetesAgentClient {
	return &kubernetesAgentClient{cc}
}

func (c *kubernetesAgentClient) Invoke(ctx context.Context, in *KubernetesRequest, opts ...grpc.CallOption) (*KubernetesResponse, error) {
	out := new(KubernetesResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.kubernetes.KubernetesAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesAgentServer is the server API for KubernetesAgent service.
type KubernetesAgentServer interface {
	Invoke(context.Context, *KubernetesRequest) (*KubernetesResponse, error)
}

// UnimplementedKubernetesAgentServer can be embedded to have forward compatible implementations.
type UnimplementedKubernetesAgentServer struct {
}

func (*UnimplementedKubernetesAgentServer) Invoke(ctx context.Context, req *KubernetesRequest) (*KubernetesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterKubernetesAgentServer(s *grpc.Server, srv KubernetesAgentServer) {
	s.RegisterService(&_KubernetesAgent_serviceDesc, srv)
}

func _KubernetesAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubernetesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.kubernetes.KubernetesAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesAgentServer).Invoke(ctx, req.(*KubernetesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KubernetesAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.kubernetes.KubernetesAgent",
	HandlerType: (*KubernetesAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _KubernetesAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubernetes.proto",
}
