// Code generated by protoc-gen-go. DO NOT EDIT.
// source: insonmnia.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NetworkType int32

const (
	NetworkType_NO_NETWORK NetworkType = 0
	NetworkType_OUTBOUND   NetworkType = 1
	NetworkType_INCOMING   NetworkType = 2
)

var NetworkType_name = map[int32]string{
	0: "NO_NETWORK",
	1: "OUTBOUND",
	2: "INCOMING",
}
var NetworkType_value = map[string]int32{
	"NO_NETWORK": 0,
	"OUTBOUND":   1,
	"INCOMING":   2,
}

func (x NetworkType) String() string {
	return proto.EnumName(NetworkType_name, int32(x))
}
func (NetworkType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type GPUCount int32

const (
	GPUCount_NO_GPU       GPUCount = 0
	GPUCount_SINGLE_GPU   GPUCount = 1
	GPUCount_MULTIPLE_GPU GPUCount = 2
)

var GPUCount_name = map[int32]string{
	0: "NO_GPU",
	1: "SINGLE_GPU",
	2: "MULTIPLE_GPU",
}
var GPUCount_value = map[string]int32{
	"NO_GPU":       0,
	"SINGLE_GPU":   1,
	"MULTIPLE_GPU": 2,
}

func (x GPUCount) String() string {
	return proto.EnumName(GPUCount_name, int32(x))
}
func (GPUCount) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type TaskStatusReply_Status int32

const (
	TaskStatusReply_UNKNOWN  TaskStatusReply_Status = 0
	TaskStatusReply_SPOOLING TaskStatusReply_Status = 1
	TaskStatusReply_SPAWNING TaskStatusReply_Status = 2
	TaskStatusReply_RUNNING  TaskStatusReply_Status = 3
	TaskStatusReply_FINISHED TaskStatusReply_Status = 4
	TaskStatusReply_BROKEN   TaskStatusReply_Status = 5
)

var TaskStatusReply_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SPOOLING",
	2: "SPAWNING",
	3: "RUNNING",
	4: "FINISHED",
	5: "BROKEN",
}
var TaskStatusReply_Status_value = map[string]int32{
	"UNKNOWN":  0,
	"SPOOLING": 1,
	"SPAWNING": 2,
	"RUNNING":  3,
	"FINISHED": 4,
	"BROKEN":   5,
}

func (x TaskStatusReply_Status) String() string {
	return proto.EnumName(TaskStatusReply_Status_name, int32(x))
}
func (TaskStatusReply_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{9, 0} }

type TaskLogsRequest_Type int32

const (
	TaskLogsRequest_STDOUT TaskLogsRequest_Type = 0
	TaskLogsRequest_STDERR TaskLogsRequest_Type = 1
	TaskLogsRequest_BOTH   TaskLogsRequest_Type = 2
)

var TaskLogsRequest_Type_name = map[int32]string{
	0: "STDOUT",
	1: "STDERR",
	2: "BOTH",
}
var TaskLogsRequest_Type_value = map[string]int32{
	"STDOUT": 0,
	"STDERR": 1,
	"BOTH":   2,
}

func (x TaskLogsRequest_Type) String() string {
	return proto.EnumName(TaskLogsRequest_Type_name, int32(x))
}
func (TaskLogsRequest_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{13, 0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type ID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ID) Reset()                    { *m = ID{} }
func (m *ID) String() string            { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()               {}
func (*ID) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TaskID struct {
	// Id is task ID itself
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// HubID is ID of hub on which task if running
	HubAddr string `protobuf:"bytes,2,opt,name=hubAddr" json:"hubAddr,omitempty"`
}

func (m *TaskID) Reset()                    { *m = TaskID{} }
func (m *TaskID) String() string            { return proto.CompactTextString(m) }
func (*TaskID) ProtoMessage()               {}
func (*TaskID) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *TaskID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskID) GetHubAddr() string {
	if m != nil {
		return m.HubAddr
	}
	return ""
}

type PingReply struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *PingReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type CPUUsage struct {
	Total uint64 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
}

func (m *CPUUsage) Reset()                    { *m = CPUUsage{} }
func (m *CPUUsage) String() string            { return proto.CompactTextString(m) }
func (*CPUUsage) ProtoMessage()               {}
func (*CPUUsage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *CPUUsage) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type MemoryUsage struct {
	MaxUsage uint64 `protobuf:"varint,1,opt,name=maxUsage" json:"maxUsage,omitempty"`
}

func (m *MemoryUsage) Reset()                    { *m = MemoryUsage{} }
func (m *MemoryUsage) String() string            { return proto.CompactTextString(m) }
func (*MemoryUsage) ProtoMessage()               {}
func (*MemoryUsage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *MemoryUsage) GetMaxUsage() uint64 {
	if m != nil {
		return m.MaxUsage
	}
	return 0
}

type NetworkUsage struct {
	TxBytes   uint64 `protobuf:"varint,1,opt,name=txBytes" json:"txBytes,omitempty"`
	RxBytes   uint64 `protobuf:"varint,2,opt,name=rxBytes" json:"rxBytes,omitempty"`
	TxPackets uint64 `protobuf:"varint,3,opt,name=txPackets" json:"txPackets,omitempty"`
	RxPackets uint64 `protobuf:"varint,4,opt,name=rxPackets" json:"rxPackets,omitempty"`
	TxErrors  uint64 `protobuf:"varint,5,opt,name=txErrors" json:"txErrors,omitempty"`
	RxErrors  uint64 `protobuf:"varint,6,opt,name=rxErrors" json:"rxErrors,omitempty"`
	TxDropped uint64 `protobuf:"varint,7,opt,name=txDropped" json:"txDropped,omitempty"`
	RxDropped uint64 `protobuf:"varint,8,opt,name=rxDropped" json:"rxDropped,omitempty"`
}

func (m *NetworkUsage) Reset()                    { *m = NetworkUsage{} }
func (m *NetworkUsage) String() string            { return proto.CompactTextString(m) }
func (*NetworkUsage) ProtoMessage()               {}
func (*NetworkUsage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *NetworkUsage) GetTxBytes() uint64 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

func (m *NetworkUsage) GetRxBytes() uint64 {
	if m != nil {
		return m.RxBytes
	}
	return 0
}

func (m *NetworkUsage) GetTxPackets() uint64 {
	if m != nil {
		return m.TxPackets
	}
	return 0
}

func (m *NetworkUsage) GetRxPackets() uint64 {
	if m != nil {
		return m.RxPackets
	}
	return 0
}

func (m *NetworkUsage) GetTxErrors() uint64 {
	if m != nil {
		return m.TxErrors
	}
	return 0
}

func (m *NetworkUsage) GetRxErrors() uint64 {
	if m != nil {
		return m.RxErrors
	}
	return 0
}

func (m *NetworkUsage) GetTxDropped() uint64 {
	if m != nil {
		return m.TxDropped
	}
	return 0
}

func (m *NetworkUsage) GetRxDropped() uint64 {
	if m != nil {
		return m.RxDropped
	}
	return 0
}

type ResourceUsage struct {
	Cpu     *CPUUsage                `protobuf:"bytes,1,opt,name=cpu" json:"cpu,omitempty"`
	Memory  *MemoryUsage             `protobuf:"bytes,2,opt,name=memory" json:"memory,omitempty"`
	Network map[string]*NetworkUsage `protobuf:"bytes,3,rep,name=network" json:"network,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ResourceUsage) Reset()                    { *m = ResourceUsage{} }
func (m *ResourceUsage) String() string            { return proto.CompactTextString(m) }
func (*ResourceUsage) ProtoMessage()               {}
func (*ResourceUsage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *ResourceUsage) GetCpu() *CPUUsage {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *ResourceUsage) GetMemory() *MemoryUsage {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *ResourceUsage) GetNetwork() map[string]*NetworkUsage {
	if m != nil {
		return m.Network
	}
	return nil
}

type InfoReply struct {
	Usage        map[string]*ResourceUsage `protobuf:"bytes,1,rep,name=usage" json:"usage,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Name         string                    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Capabilities *Capabilities             `protobuf:"bytes,3,opt,name=capabilities" json:"capabilities,omitempty"`
}

func (m *InfoReply) Reset()                    { *m = InfoReply{} }
func (m *InfoReply) String() string            { return proto.CompactTextString(m) }
func (*InfoReply) ProtoMessage()               {}
func (*InfoReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *InfoReply) GetUsage() map[string]*ResourceUsage {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *InfoReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InfoReply) GetCapabilities() *Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

type TaskStatusReply struct {
	Status             TaskStatusReply_Status `protobuf:"varint,1,opt,name=status,enum=sonm.TaskStatusReply_Status" json:"status,omitempty"`
	ImageName          string                 `protobuf:"bytes,2,opt,name=imageName" json:"imageName,omitempty"`
	Ports              string                 `protobuf:"bytes,3,opt,name=ports" json:"ports,omitempty"`
	Uptime             uint64                 `protobuf:"varint,4,opt,name=uptime" json:"uptime,omitempty"`
	Usage              *ResourceUsage         `protobuf:"bytes,5,opt,name=usage" json:"usage,omitempty"`
	AvailableResources *AvailableResources    `protobuf:"bytes,6,opt,name=availableResources" json:"availableResources,omitempty"`
	MinerID            string                 `protobuf:"bytes,7,opt,name=minerID" json:"minerID,omitempty"`
}

func (m *TaskStatusReply) Reset()                    { *m = TaskStatusReply{} }
func (m *TaskStatusReply) String() string            { return proto.CompactTextString(m) }
func (*TaskStatusReply) ProtoMessage()               {}
func (*TaskStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *TaskStatusReply) GetStatus() TaskStatusReply_Status {
	if m != nil {
		return m.Status
	}
	return TaskStatusReply_UNKNOWN
}

func (m *TaskStatusReply) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *TaskStatusReply) GetPorts() string {
	if m != nil {
		return m.Ports
	}
	return ""
}

func (m *TaskStatusReply) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *TaskStatusReply) GetUsage() *ResourceUsage {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *TaskStatusReply) GetAvailableResources() *AvailableResources {
	if m != nil {
		return m.AvailableResources
	}
	return nil
}

func (m *TaskStatusReply) GetMinerID() string {
	if m != nil {
		return m.MinerID
	}
	return ""
}

type AvailableResources struct {
	NumCPUs            int64  `protobuf:"varint,1,opt,name=numCPUs" json:"numCPUs,omitempty"`
	NumGPUs            int64  `protobuf:"varint,2,opt,name=numGPUs" json:"numGPUs,omitempty"`
	Memory             uint64 `protobuf:"varint,3,opt,name=memory" json:"memory,omitempty"`
	CPUPeriod          int64  `protobuf:"varint,4,opt,name=CPUPeriod" json:"CPUPeriod,omitempty"`
	CPUQuota           int64  `protobuf:"varint,5,opt,name=CPUQuota" json:"CPUQuota,omitempty"`
	CPURealtimePeriod  int64  `protobuf:"varint,6,opt,name=CPURealtimePeriod" json:"CPURealtimePeriod,omitempty"`
	CPURealtimeRuntime int64  `protobuf:"varint,7,opt,name=CPURealtimeRuntime" json:"CPURealtimeRuntime,omitempty"`
	CpusetCpus         string `protobuf:"bytes,8,opt,name=CpusetCpus" json:"CpusetCpus,omitempty"`
	CpusetMems         string `protobuf:"bytes,9,opt,name=CpusetMems" json:"CpusetMems,omitempty"`
	DiskQuota          int64  `protobuf:"varint,10,opt,name=DiskQuota" json:"DiskQuota,omitempty"`
	KernelMemory       int64  `protobuf:"varint,11,opt,name=KernelMemory" json:"KernelMemory,omitempty"`
	MemoryReservation  int64  `protobuf:"varint,12,opt,name=MemoryReservation" json:"MemoryReservation,omitempty"`
	MemorySwap         int64  `protobuf:"varint,13,opt,name=MemorySwap" json:"MemorySwap,omitempty"`
	PidsLimit          int64  `protobuf:"varint,14,opt,name=PidsLimit" json:"PidsLimit,omitempty"`
	Cgroup             string `protobuf:"bytes,15,opt,name=cgroup" json:"cgroup,omitempty"`
	CgroupParent       string `protobuf:"bytes,16,opt,name=cgroupParent" json:"cgroupParent,omitempty"`
}

func (m *AvailableResources) Reset()                    { *m = AvailableResources{} }
func (m *AvailableResources) String() string            { return proto.CompactTextString(m) }
func (*AvailableResources) ProtoMessage()               {}
func (*AvailableResources) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

func (m *AvailableResources) GetNumCPUs() int64 {
	if m != nil {
		return m.NumCPUs
	}
	return 0
}

func (m *AvailableResources) GetNumGPUs() int64 {
	if m != nil {
		return m.NumGPUs
	}
	return 0
}

func (m *AvailableResources) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *AvailableResources) GetCPUPeriod() int64 {
	if m != nil {
		return m.CPUPeriod
	}
	return 0
}

func (m *AvailableResources) GetCPUQuota() int64 {
	if m != nil {
		return m.CPUQuota
	}
	return 0
}

func (m *AvailableResources) GetCPURealtimePeriod() int64 {
	if m != nil {
		return m.CPURealtimePeriod
	}
	return 0
}

func (m *AvailableResources) GetCPURealtimeRuntime() int64 {
	if m != nil {
		return m.CPURealtimeRuntime
	}
	return 0
}

func (m *AvailableResources) GetCpusetCpus() string {
	if m != nil {
		return m.CpusetCpus
	}
	return ""
}

func (m *AvailableResources) GetCpusetMems() string {
	if m != nil {
		return m.CpusetMems
	}
	return ""
}

func (m *AvailableResources) GetDiskQuota() int64 {
	if m != nil {
		return m.DiskQuota
	}
	return 0
}

func (m *AvailableResources) GetKernelMemory() int64 {
	if m != nil {
		return m.KernelMemory
	}
	return 0
}

func (m *AvailableResources) GetMemoryReservation() int64 {
	if m != nil {
		return m.MemoryReservation
	}
	return 0
}

func (m *AvailableResources) GetMemorySwap() int64 {
	if m != nil {
		return m.MemorySwap
	}
	return 0
}

func (m *AvailableResources) GetPidsLimit() int64 {
	if m != nil {
		return m.PidsLimit
	}
	return 0
}

func (m *AvailableResources) GetCgroup() string {
	if m != nil {
		return m.Cgroup
	}
	return ""
}

func (m *AvailableResources) GetCgroupParent() string {
	if m != nil {
		return m.CgroupParent
	}
	return ""
}

type StatusMapReply struct {
	Statuses map[string]*TaskStatusReply `protobuf:"bytes,1,rep,name=statuses" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *StatusMapReply) Reset()                    { *m = StatusMapReply{} }
func (m *StatusMapReply) String() string            { return proto.CompactTextString(m) }
func (*StatusMapReply) ProtoMessage()               {}
func (*StatusMapReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{11} }

func (m *StatusMapReply) GetStatuses() map[string]*TaskStatusReply {
	if m != nil {
		return m.Statuses
	}
	return nil
}

type ContainerRestartPolicy struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	MaximumRetryCount uint32 `protobuf:"varint,2,opt,name=maximumRetryCount" json:"maximumRetryCount,omitempty"`
}

func (m *ContainerRestartPolicy) Reset()                    { *m = ContainerRestartPolicy{} }
func (m *ContainerRestartPolicy) String() string            { return proto.CompactTextString(m) }
func (*ContainerRestartPolicy) ProtoMessage()               {}
func (*ContainerRestartPolicy) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{12} }

func (m *ContainerRestartPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerRestartPolicy) GetMaximumRetryCount() uint32 {
	if m != nil {
		return m.MaximumRetryCount
	}
	return 0
}

type TaskLogsRequest struct {
	Type          TaskLogsRequest_Type `protobuf:"varint,1,opt,name=type,enum=sonm.TaskLogsRequest_Type" json:"type,omitempty"`
	Id            string               `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Since         string               `protobuf:"bytes,3,opt,name=since" json:"since,omitempty"`
	AddTimestamps bool                 `protobuf:"varint,4,opt,name=addTimestamps" json:"addTimestamps,omitempty"`
	Follow        bool                 `protobuf:"varint,5,opt,name=Follow" json:"Follow,omitempty"`
	Tail          string               `protobuf:"bytes,6,opt,name=Tail" json:"Tail,omitempty"`
	Details       bool                 `protobuf:"varint,7,opt,name=Details" json:"Details,omitempty"`
	HubAddr       string               `protobuf:"bytes,8,opt,name=hubAddr" json:"hubAddr,omitempty"`
}

func (m *TaskLogsRequest) Reset()                    { *m = TaskLogsRequest{} }
func (m *TaskLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskLogsRequest) ProtoMessage()               {}
func (*TaskLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{13} }

func (m *TaskLogsRequest) GetType() TaskLogsRequest_Type {
	if m != nil {
		return m.Type
	}
	return TaskLogsRequest_STDOUT
}

func (m *TaskLogsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskLogsRequest) GetSince() string {
	if m != nil {
		return m.Since
	}
	return ""
}

func (m *TaskLogsRequest) GetAddTimestamps() bool {
	if m != nil {
		return m.AddTimestamps
	}
	return false
}

func (m *TaskLogsRequest) GetFollow() bool {
	if m != nil {
		return m.Follow
	}
	return false
}

func (m *TaskLogsRequest) GetTail() string {
	if m != nil {
		return m.Tail
	}
	return ""
}

func (m *TaskLogsRequest) GetDetails() bool {
	if m != nil {
		return m.Details
	}
	return false
}

func (m *TaskLogsRequest) GetHubAddr() string {
	if m != nil {
		return m.HubAddr
	}
	return ""
}

type TaskLogsChunk struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TaskLogsChunk) Reset()                    { *m = TaskLogsChunk{} }
func (m *TaskLogsChunk) String() string            { return proto.CompactTextString(m) }
func (*TaskLogsChunk) ProtoMessage()               {}
func (*TaskLogsChunk) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{14} }

func (m *TaskLogsChunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DiscoverHubRequest struct {
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *DiscoverHubRequest) Reset()                    { *m = DiscoverHubRequest{} }
func (m *DiscoverHubRequest) String() string            { return proto.CompactTextString(m) }
func (*DiscoverHubRequest) ProtoMessage()               {}
func (*DiscoverHubRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{15} }

func (m *DiscoverHubRequest) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type TaskResourceRequirements struct {
	// CPUCores specifies the number of CPU cores required.
	CPUCores uint64 `protobuf:"varint,1,opt,name=CPUCores" json:"CPUCores,omitempty"`
	// MaxMemory specifies the maximum memory in bytes required.
	MaxMemory int64 `protobuf:"varint,2,opt,name=maxMemory" json:"maxMemory,omitempty"`
	// GPUCount Describes whether a task requires GPU support.
	GPUSupport GPUCount `protobuf:"varint,3,opt,name=GPUSupport,enum=sonm.GPUCount" json:"GPUSupport,omitempty"`
}

func (m *TaskResourceRequirements) Reset()                    { *m = TaskResourceRequirements{} }
func (m *TaskResourceRequirements) String() string            { return proto.CompactTextString(m) }
func (*TaskResourceRequirements) ProtoMessage()               {}
func (*TaskResourceRequirements) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{16} }

func (m *TaskResourceRequirements) GetCPUCores() uint64 {
	if m != nil {
		return m.CPUCores
	}
	return 0
}

func (m *TaskResourceRequirements) GetMaxMemory() int64 {
	if m != nil {
		return m.MaxMemory
	}
	return 0
}

func (m *TaskResourceRequirements) GetGPUSupport() GPUCount {
	if m != nil {
		return m.GPUSupport
	}
	return GPUCount_NO_GPU
}

type Chunk struct {
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{17} }

func (m *Chunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

type Progress struct {
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
}

func (m *Progress) Reset()                    { *m = Progress{} }
func (m *Progress) String() string            { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()               {}
func (*Progress) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{18} }

func (m *Progress) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "sonm.Empty")
	proto.RegisterType((*ID)(nil), "sonm.ID")
	proto.RegisterType((*TaskID)(nil), "sonm.TaskID")
	proto.RegisterType((*PingReply)(nil), "sonm.PingReply")
	proto.RegisterType((*CPUUsage)(nil), "sonm.CPUUsage")
	proto.RegisterType((*MemoryUsage)(nil), "sonm.MemoryUsage")
	proto.RegisterType((*NetworkUsage)(nil), "sonm.NetworkUsage")
	proto.RegisterType((*ResourceUsage)(nil), "sonm.ResourceUsage")
	proto.RegisterType((*InfoReply)(nil), "sonm.InfoReply")
	proto.RegisterType((*TaskStatusReply)(nil), "sonm.TaskStatusReply")
	proto.RegisterType((*AvailableResources)(nil), "sonm.AvailableResources")
	proto.RegisterType((*StatusMapReply)(nil), "sonm.StatusMapReply")
	proto.RegisterType((*ContainerRestartPolicy)(nil), "sonm.ContainerRestartPolicy")
	proto.RegisterType((*TaskLogsRequest)(nil), "sonm.TaskLogsRequest")
	proto.RegisterType((*TaskLogsChunk)(nil), "sonm.TaskLogsChunk")
	proto.RegisterType((*DiscoverHubRequest)(nil), "sonm.DiscoverHubRequest")
	proto.RegisterType((*TaskResourceRequirements)(nil), "sonm.TaskResourceRequirements")
	proto.RegisterType((*Chunk)(nil), "sonm.Chunk")
	proto.RegisterType((*Progress)(nil), "sonm.Progress")
	proto.RegisterEnum("sonm.NetworkType", NetworkType_name, NetworkType_value)
	proto.RegisterEnum("sonm.GPUCount", GPUCount_name, GPUCount_value)
	proto.RegisterEnum("sonm.TaskStatusReply_Status", TaskStatusReply_Status_name, TaskStatusReply_Status_value)
	proto.RegisterEnum("sonm.TaskLogsRequest_Type", TaskLogsRequest_Type_name, TaskLogsRequest_Type_value)
}

func init() { proto.RegisterFile("insonmnia.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 1311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0x4d, 0x6f, 0x1b, 0x37,
	0x13, 0x8e, 0xbe, 0xa5, 0x91, 0x6c, 0x2b, 0x7c, 0xfd, 0x06, 0x0b, 0x21, 0x0d, 0x8c, 0x4d, 0x0f,
	0x4e, 0x1a, 0x08, 0x81, 0x5b, 0x14, 0x69, 0x0e, 0x05, 0x6c, 0x49, 0xb1, 0x05, 0x5b, 0xab, 0x2d,
	0xa5, 0x45, 0x8a, 0x5e, 0x02, 0x5a, 0x62, 0x9d, 0x85, 0xb5, 0x1f, 0xe5, 0x72, 0x1d, 0xab, 0xe7,
	0xfe, 0x87, 0xde, 0x8b, 0xfe, 0xa8, 0xfe, 0x84, 0x9e, 0x7b, 0xee, 0xa1, 0x20, 0x87, 0xbb, 0x5a,
	0xd5, 0x42, 0x2f, 0x0b, 0x3e, 0xf3, 0x0c, 0x97, 0x33, 0x43, 0xce, 0x43, 0xc2, 0x81, 0x1f, 0x26,
	0x51, 0x18, 0x84, 0x3e, 0xeb, 0xc7, 0x22, 0x92, 0x11, 0xa9, 0x2a, 0xd8, 0x23, 0x0b, 0x16, 0xb3,
	0x6b, 0x7f, 0xe5, 0x4b, 0x9f, 0x27, 0xc8, 0xf4, 0x0e, 0xa4, 0x1f, 0xf0, 0x44, 0xb2, 0x20, 0x46,
	0x83, 0xdd, 0x80, 0xda, 0x28, 0x88, 0xe5, 0xda, 0x3e, 0x84, 0xf2, 0x78, 0x48, 0xf6, 0xa1, 0xec,
	0x2f, 0xad, 0xd2, 0x51, 0xe9, 0xb8, 0x45, 0xcb, 0xfe, 0xd2, 0x3e, 0x81, 0xfa, 0x9c, 0x25, 0xb7,
	0x0f, 0x19, 0x62, 0x41, 0xe3, 0x63, 0x7a, 0x7d, 0xba, 0x5c, 0x0a, 0xab, 0xac, 0x8d, 0x19, 0xb4,
	0x9f, 0x43, 0xcb, 0xf5, 0xc3, 0x1b, 0xca, 0xe3, 0xd5, 0x9a, 0x3c, 0x81, 0x7a, 0x22, 0x99, 0x4c,
	0x13, 0x33, 0xd5, 0x20, 0xfb, 0x08, 0x9a, 0x03, 0xd7, 0xf3, 0x12, 0x76, 0xc3, 0xc9, 0x21, 0xd4,
	0x64, 0x24, 0xd9, 0x4a, 0xbb, 0x54, 0x29, 0x02, 0xfb, 0x05, 0xb4, 0x27, 0x3c, 0x88, 0xc4, 0x1a,
	0x9d, 0x7a, 0xd0, 0x0c, 0xd8, 0xbd, 0x1e, 0x1b, 0xbf, 0x1c, 0xdb, 0x7f, 0x95, 0xa0, 0xe3, 0x70,
	0xf9, 0x29, 0x12, 0xb7, 0xe8, 0x6c, 0x41, 0x43, 0xde, 0x9f, 0xad, 0x25, 0x4f, 0x8c, 0x6f, 0x06,
	0x15, 0x23, 0x0c, 0x53, 0x46, 0xc6, 0x40, 0xf2, 0x14, 0x5a, 0xf2, 0xde, 0x65, 0x8b, 0x5b, 0x2e,
	0x13, 0xab, 0xa2, 0xb9, 0x8d, 0x41, 0xb1, 0x22, 0x67, 0xab, 0xc8, 0xe6, 0x06, 0x15, 0x9c, 0xbc,
	0x1f, 0x09, 0x11, 0x89, 0xc4, 0xaa, 0x61, 0x70, 0x19, 0x56, 0x9c, 0xc8, 0xb8, 0x3a, 0x72, 0x19,
	0xc6, 0x35, 0x87, 0x22, 0x8a, 0x63, 0xbe, 0xb4, 0x1a, 0xd9, 0x9a, 0xc6, 0x80, 0x6b, 0x66, 0x6c,
	0x33, 0x5b, 0xd3, 0x18, 0xec, 0x3f, 0x4b, 0xb0, 0x47, 0x79, 0x12, 0xa5, 0x62, 0xc1, 0x31, 0xeb,
	0x23, 0xa8, 0x2c, 0xe2, 0x54, 0x67, 0xdc, 0x3e, 0xd9, 0xef, 0xab, 0x43, 0xd0, 0xcf, 0x8a, 0x4c,
	0x15, 0x45, 0x5e, 0x40, 0x3d, 0xd0, 0x35, 0xd5, 0xc9, 0xb7, 0x4f, 0x1e, 0xa3, 0x53, 0xa1, 0xce,
	0xd4, 0x38, 0x90, 0xb7, 0xd0, 0x08, 0xb1, 0xa4, 0x56, 0xe5, 0xa8, 0x72, 0xdc, 0x3e, 0x39, 0x42,
	0xdf, 0xad, 0x25, 0xfb, 0xa6, 0xea, 0xa3, 0x50, 0x8a, 0x35, 0xcd, 0x26, 0xf4, 0x9c, 0x7c, 0x3b,
	0x34, 0x41, 0xba, 0x50, 0xb9, 0xe5, 0x6b, 0x73, 0x02, 0xd4, 0x90, 0x1c, 0x43, 0xed, 0x8e, 0xad,
	0x52, 0x6e, 0xe2, 0x20, 0xf8, 0xef, 0xe2, 0x1e, 0x52, 0x74, 0x78, 0x5b, 0x7e, 0x53, 0xb2, 0xff,
	0x28, 0x41, 0x6b, 0x1c, 0xfe, 0x18, 0xe1, 0x91, 0x7a, 0x0d, 0xb5, 0xd4, 0x1c, 0x03, 0x15, 0x57,
	0x0f, 0xe7, 0xe6, 0x7c, 0x5f, 0x4f, 0xc7, 0x88, 0xd0, 0x91, 0x10, 0xa8, 0x86, 0x2c, 0xe0, 0xe6,
	0xa0, 0xea, 0x31, 0xf9, 0x1a, 0x3a, 0xc5, 0xfe, 0xd0, 0x3b, 0x9e, 0x07, 0x32, 0x28, 0x30, 0x74,
	0xcb, 0xaf, 0x37, 0x01, 0xd8, 0x2c, 0xb0, 0x23, 0xb3, 0x17, 0xdb, 0x99, 0xfd, 0x6f, 0x47, 0xd5,
	0x8a, 0xa9, 0xfd, 0x5d, 0x86, 0x03, 0xd5, 0x61, 0x33, 0xdd, 0x16, 0x98, 0xe0, 0x57, 0x5b, 0x3d,
	0xb3, 0x7f, 0xf2, 0x14, 0xff, 0xf1, 0x2f, 0xb7, 0xbe, 0x19, 0x1b, 0x5f, 0x75, 0x5a, 0xfc, 0x80,
	0xdd, 0x70, 0x67, 0x93, 0xe9, 0xc6, 0xa0, 0x7a, 0x2c, 0x8e, 0x84, 0x39, 0xd9, 0x2d, 0x8a, 0x40,
	0x75, 0x67, 0x1a, 0x2b, 0x49, 0x30, 0x47, 0xda, 0x20, 0x95, 0x04, 0x96, 0xb8, 0xf6, 0x1f, 0x49,
	0x60, 0x6d, 0x2f, 0x80, 0xb0, 0x3b, 0xe6, 0xaf, 0xd8, 0xf5, 0x8a, 0x67, 0x0e, 0x78, 0xd0, 0xdb,
	0x27, 0x16, 0xce, 0x3b, 0x7d, 0xc0, 0xd3, 0x1d, 0x73, 0x54, 0x6b, 0x06, 0x7e, 0xc8, 0xc5, 0x78,
	0xa8, 0x5b, 0xa1, 0x45, 0x33, 0x68, 0x7f, 0x0f, 0x75, 0x4c, 0x96, 0xb4, 0xa1, 0xe1, 0x39, 0x97,
	0xce, 0xf4, 0xbd, 0xd3, 0x7d, 0x44, 0x3a, 0xd0, 0x9c, 0xb9, 0xd3, 0xe9, 0xd5, 0xd8, 0x39, 0xef,
	0x96, 0x10, 0x9d, 0xbe, 0x77, 0x14, 0x2a, 0x2b, 0x47, 0xea, 0x39, 0x1a, 0x54, 0x14, 0xf5, 0x6e,
	0xec, 0x8c, 0x67, 0x17, 0xa3, 0x61, 0xb7, 0x4a, 0x00, 0xea, 0x67, 0x74, 0x7a, 0x39, 0x72, 0xba,
	0x35, 0xfb, 0xb7, 0x2a, 0x90, 0xd3, 0x9d, 0xa1, 0x84, 0x69, 0x30, 0x70, 0x3d, 0xdc, 0x82, 0x0a,
	0xcd, 0xa0, 0x61, 0xce, 0x15, 0x53, 0xce, 0x19, 0x05, 0x55, 0x2d, 0x4d, 0x6f, 0xa1, 0x78, 0x64,
	0x8d, 0xf4, 0x14, 0x5a, 0x03, 0xd7, 0x73, 0xb9, 0xf0, 0xa3, 0xa5, 0x2e, 0x73, 0x85, 0x6e, 0x0c,
	0x4a, 0x1d, 0x06, 0xae, 0xf7, 0x5d, 0x1a, 0x49, 0xa6, 0x8b, 0x5d, 0xa1, 0x39, 0x26, 0xaf, 0xe0,
	0xf1, 0xc0, 0xf5, 0x28, 0x67, 0x2b, 0xb5, 0x29, 0xe6, 0x0f, 0x75, 0xed, 0xf4, 0x90, 0x20, 0x7d,
	0x20, 0x05, 0x23, 0x4d, 0x43, 0xbd, 0xaf, 0x0d, 0xed, 0xbe, 0x83, 0x21, 0xcf, 0x00, 0x06, 0x71,
	0x9a, 0x70, 0xa9, 0xbe, 0x5a, 0x5e, 0x5a, 0xb4, 0x60, 0xd9, 0xf0, 0x13, 0x1e, 0x24, 0x56, 0xab,
	0xc8, 0x2b, 0x8b, 0xca, 0x6b, 0xe8, 0x27, 0xb7, 0x18, 0x3a, 0x60, 0x5e, 0xb9, 0x81, 0xd8, 0xd0,
	0xb9, 0xe4, 0x22, 0xe4, 0x2b, 0xd4, 0x16, 0xab, 0xad, 0x1d, 0xb6, 0x6c, 0x2a, 0x3f, 0x1c, 0x51,
	0x9e, 0x70, 0x71, 0xc7, 0xa4, 0x1f, 0x85, 0x56, 0x07, 0xf3, 0x7b, 0x40, 0xa8, 0x78, 0xd0, 0x38,
	0xfb, 0xc4, 0x62, 0x6b, 0x4f, 0xbb, 0x15, 0x2c, 0x2a, 0x1e, 0xd7, 0x5f, 0x26, 0x57, 0x7e, 0xe0,
	0x4b, 0x6b, 0x1f, 0xe3, 0xc9, 0x0d, 0x6a, 0x77, 0x16, 0x37, 0x22, 0x4a, 0x63, 0xeb, 0x00, 0xef,
	0x21, 0x44, 0x2a, 0x4e, 0x1c, 0xb9, 0x4c, 0xf0, 0x50, 0x5a, 0x5d, 0xcd, 0x6e, 0xd9, 0xec, 0xdf,
	0x4b, 0xb0, 0x8f, 0xe7, 0x6f, 0xc2, 0x62, 0x6c, 0xd1, 0x6f, 0xa1, 0x89, 0x6d, 0xa7, 0x6f, 0x18,
	0x25, 0x43, 0x36, 0x9e, 0xf5, 0x6d, 0x3f, 0x03, 0x79, 0x82, 0x72, 0x94, 0xcf, 0xe9, 0x51, 0xd8,
	0xdb, 0xa2, 0x76, 0x08, 0xc9, 0x17, 0xdb, 0x42, 0xf2, 0xff, 0x9d, 0x22, 0x50, 0x94, 0x92, 0x1f,
	0xe0, 0xc9, 0x20, 0x0a, 0x25, 0x53, 0x4d, 0x43, 0xd5, 0x25, 0x2f, 0xa4, 0x1b, 0xad, 0xfc, 0xc5,
	0x3a, 0xd7, 0xbf, 0x52, 0x41, 0xff, 0x5e, 0xc1, 0xe3, 0x80, 0xdd, 0xfb, 0x41, 0x1a, 0x50, 0x2e,
	0xc5, 0x7a, 0x10, 0xa5, 0xa1, 0xd4, 0x4b, 0xed, 0xd1, 0x87, 0x84, 0xfd, 0xab, 0x91, 0xa9, 0xab,
	0xe8, 0x26, 0xa1, 0xfc, 0xa7, 0x94, 0x27, 0x92, 0xf4, 0xa1, 0x2a, 0xd7, 0x31, 0x37, 0x22, 0xd5,
	0xdb, 0xc4, 0x57, 0x70, 0xea, 0xcf, 0xd7, 0x31, 0xa7, 0xda, 0xcf, 0xbc, 0x20, 0xca, 0xf9, 0x0b,
	0xe2, 0x10, 0x6a, 0x89, 0x1f, 0x2e, 0x78, 0x26, 0x49, 0x1a, 0x90, 0xcf, 0x61, 0x8f, 0x2d, 0x97,
	0xf3, 0xec, 0x99, 0x82, 0x97, 0x6d, 0x93, 0x6e, 0x1b, 0xd5, 0x76, 0xbe, 0x8b, 0x56, 0xab, 0xe8,
	0x93, 0x6e, 0x9a, 0x26, 0x35, 0x48, 0x65, 0x3a, 0x67, 0xfe, 0x4a, 0x77, 0x49, 0x8b, 0xea, 0xb1,
	0x6a, 0xd9, 0x21, 0x97, 0xcc, 0x5f, 0x25, 0xba, 0x1b, 0x9a, 0x34, 0x83, 0xc5, 0x37, 0x4c, 0x73,
	0xfb, 0x0d, 0x73, 0x0c, 0x55, 0x15, 0xb9, 0xd2, 0x8a, 0xd9, 0x7c, 0x38, 0xf5, 0xe6, 0xdd, 0x47,
	0x66, 0x3c, 0xa2, 0xb4, 0x5b, 0x22, 0x4d, 0xa8, 0x9e, 0x4d, 0xe7, 0x17, 0xdd, 0xb2, 0xfd, 0x1c,
	0xf6, 0xb2, 0x9c, 0x07, 0x1f, 0xd3, 0xf0, 0x56, 0x85, 0xb0, 0x64, 0x92, 0xe9, 0xb2, 0x74, 0xa8,
	0x1e, 0xdb, 0xaf, 0x81, 0x0c, 0xfd, 0x64, 0x11, 0xdd, 0x71, 0x71, 0x91, 0x5e, 0x67, 0x05, 0xec,
	0x41, 0x93, 0x87, 0xcb, 0x38, 0xf2, 0x43, 0x69, 0xb6, 0x26, 0xc7, 0xf6, 0x2f, 0x25, 0xb0, 0xd4,
	0x7f, 0x33, 0x4d, 0x52, 0x73, 0x7c, 0xc1, 0x03, 0x1e, 0xe2, 0x73, 0x63, 0xe0, 0x7a, 0x83, 0x48,
	0xe4, 0xef, 0x9b, 0x1c, 0xab, 0x36, 0x08, 0xd8, 0xfd, 0x64, 0x73, 0xcb, 0x57, 0xe8, 0xc6, 0x40,
	0xfa, 0x00, 0xe7, 0xae, 0x37, 0x4b, 0x63, 0xa5, 0xff, 0xba, 0xf0, 0xfb, 0xd9, 0x4b, 0xe1, 0x5c,
	0xfd, 0x21, 0x0d, 0x25, 0x2d, 0x78, 0xd8, 0x9f, 0x41, 0x0d, 0xb3, 0x3a, 0x84, 0xda, 0x42, 0x0d,
	0x4c, 0x5a, 0x08, 0xec, 0x67, 0xd0, 0x74, 0x45, 0x74, 0x23, 0x78, 0x92, 0xa8, 0xbc, 0x13, 0xff,
	0x67, 0x6e, 0x04, 0x53, 0x8f, 0x5f, 0x7e, 0x03, 0x6d, 0x73, 0xa7, 0xcf, 0xf1, 0x04, 0x80, 0x33,
	0xfd, 0xe0, 0x8c, 0xe6, 0xef, 0xa7, 0xf4, 0x12, 0x05, 0x7c, 0xea, 0xcd, 0xcf, 0xa6, 0x9e, 0x33,
	0x44, 0x01, 0x1f, 0x3b, 0x83, 0xe9, 0x44, 0x0b, 0xf8, 0xcb, 0x37, 0xd0, 0xcc, 0x22, 0x52, 0x95,
	0x77, 0xa6, 0x1f, 0xce, 0x5d, 0xaf, 0xfb, 0x48, 0xfd, 0x63, 0x36, 0x76, 0xce, 0xaf, 0x46, 0x1a,
	0x97, 0x48, 0x17, 0x3a, 0x13, 0xef, 0x6a, 0x3e, 0x76, 0x8d, 0xa5, 0x7c, 0x5d, 0xd7, 0x2f, 0xdb,
	0x2f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x01, 0x6b, 0xe8, 0xe1, 0x17, 0x0b, 0x00, 0x00,
}
