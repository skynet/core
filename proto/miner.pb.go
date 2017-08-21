// Code generated by protoc-gen-go. DO NOT EDIT.
// source: miner.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NATType int32

const (
	NATType_NONE                   NATType = 0
	NATType_BLOCKED                NATType = 1
	NATType_FULL                   NATType = 2
	NATType_SYMMETRIC              NATType = 3
	NATType_RESTRICTED             NATType = 4
	NATType_PORT_RESTRICTED        NATType = 5
	NATType_SYMMETRIC_UDP_FIREWALL NATType = 6
	NATType_UNKNOWN                NATType = 7
)

var NATType_name = map[int32]string{
	0: "NONE",
	1: "BLOCKED",
	2: "FULL",
	3: "SYMMETRIC",
	4: "RESTRICTED",
	5: "PORT_RESTRICTED",
	6: "SYMMETRIC_UDP_FIREWALL",
	7: "UNKNOWN",
}
var NATType_value = map[string]int32{
	"NONE":                   0,
	"BLOCKED":                1,
	"FULL":                   2,
	"SYMMETRIC":              3,
	"RESTRICTED":             4,
	"PORT_RESTRICTED":        5,
	"SYMMETRIC_UDP_FIREWALL": 6,
	"UNKNOWN":                7,
}

func (x NATType) String() string {
	return proto.EnumName(NATType_name, int32(x))
}
func (NATType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type MinerInfoRequest struct {
}

func (m *MinerInfoRequest) Reset()                    { *m = MinerInfoRequest{} }
func (m *MinerInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerInfoRequest) ProtoMessage()               {}
func (*MinerInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type MinerHandshakeRequest struct {
	Hub string `protobuf:"bytes,1,opt,name=hub" json:"hub,omitempty"`
}

func (m *MinerHandshakeRequest) Reset()                    { *m = MinerHandshakeRequest{} }
func (m *MinerHandshakeRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerHandshakeRequest) ProtoMessage()               {}
func (*MinerHandshakeRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *MinerHandshakeRequest) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

type MinerHandshakeReply struct {
	Miner        string        `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
	Capabilities *Capabilities `protobuf:"bytes,2,opt,name=capabilities" json:"capabilities,omitempty"`
	NatType      NATType       `protobuf:"varint,3,opt,name=natType,enum=sonm.NATType" json:"natType,omitempty"`
}

func (m *MinerHandshakeReply) Reset()                    { *m = MinerHandshakeReply{} }
func (m *MinerHandshakeReply) String() string            { return proto.CompactTextString(m) }
func (*MinerHandshakeReply) ProtoMessage()               {}
func (*MinerHandshakeReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *MinerHandshakeReply) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *MinerHandshakeReply) GetCapabilities() *Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

func (m *MinerHandshakeReply) GetNatType() NATType {
	if m != nil {
		return m.NatType
	}
	return NATType_NONE
}

type MinerStartRequest struct {
	Id            string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Registry      string                  `protobuf:"bytes,2,opt,name=registry" json:"registry,omitempty"`
	Image         string                  `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Auth          string                  `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	RestartPolicy *ContainerRestartPolicy `protobuf:"bytes,5,opt,name=restartPolicy" json:"restartPolicy,omitempty"`
	Resources     *ContainerResources     `protobuf:"bytes,6,opt,name=resources" json:"resources,omitempty"`
	PublicKeyData string                  `protobuf:"bytes,7,opt,name=PublicKeyData,json=publicKeyData" json:"PublicKeyData,omitempty"`
}

func (m *MinerStartRequest) Reset()                    { *m = MinerStartRequest{} }
func (m *MinerStartRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerStartRequest) ProtoMessage()               {}
func (*MinerStartRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *MinerStartRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MinerStartRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *MinerStartRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *MinerStartRequest) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *MinerStartRequest) GetRestartPolicy() *ContainerRestartPolicy {
	if m != nil {
		return m.RestartPolicy
	}
	return nil
}

func (m *MinerStartRequest) GetResources() *ContainerResources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *MinerStartRequest) GetPublicKeyData() string {
	if m != nil {
		return m.PublicKeyData
	}
	return ""
}

type MinerStartReply struct {
	Container string                          `protobuf:"bytes,1,opt,name=container" json:"container,omitempty"`
	Ports     map[string]*MinerStartReplyPort `protobuf:"bytes,2,rep,name=Ports,json=ports" json:"Ports,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MinerStartReply) Reset()                    { *m = MinerStartReply{} }
func (m *MinerStartReply) String() string            { return proto.CompactTextString(m) }
func (*MinerStartReply) ProtoMessage()               {}
func (*MinerStartReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *MinerStartReply) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *MinerStartReply) GetPorts() map[string]*MinerStartReplyPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

type MinerStartReplyPort struct {
	IP   string `protobuf:"bytes,1,opt,name=IP,json=iP" json:"IP,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
}

func (m *MinerStartReplyPort) Reset()                    { *m = MinerStartReplyPort{} }
func (m *MinerStartReplyPort) String() string            { return proto.CompactTextString(m) }
func (*MinerStartReplyPort) ProtoMessage()               {}
func (*MinerStartReplyPort) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4, 0} }

func (m *MinerStartReplyPort) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *MinerStartReplyPort) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type MinerStatusMapRequest struct {
}

func (m *MinerStatusMapRequest) Reset()                    { *m = MinerStatusMapRequest{} }
func (m *MinerStatusMapRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerStatusMapRequest) ProtoMessage()               {}
func (*MinerStatusMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func init() {
	proto.RegisterType((*MinerInfoRequest)(nil), "sonm.MinerInfoRequest")
	proto.RegisterType((*MinerHandshakeRequest)(nil), "sonm.MinerHandshakeRequest")
	proto.RegisterType((*MinerHandshakeReply)(nil), "sonm.MinerHandshakeReply")
	proto.RegisterType((*MinerStartRequest)(nil), "sonm.MinerStartRequest")
	proto.RegisterType((*MinerStartReply)(nil), "sonm.MinerStartReply")
	proto.RegisterType((*MinerStartReplyPort)(nil), "sonm.MinerStartReply.port")
	proto.RegisterType((*MinerStatusMapRequest)(nil), "sonm.MinerStatusMapRequest")
	proto.RegisterEnum("sonm.NATType", NATType_name, NATType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Miner service

type MinerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Info(ctx context.Context, in *MinerInfoRequest, opts ...grpc.CallOption) (*InfoReply, error)
	Handshake(ctx context.Context, in *MinerHandshakeRequest, opts ...grpc.CallOption) (*MinerHandshakeReply, error)
	Start(ctx context.Context, in *MinerStartRequest, opts ...grpc.CallOption) (*MinerStartReply, error)
	Stop(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error)
	TasksStatus(ctx context.Context, opts ...grpc.CallOption) (Miner_TasksStatusClient, error)
	TaskDetails(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error)
	TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Miner_TaskLogsClient, error)
}

type minerClient struct {
	cc *grpc.ClientConn
}

func NewMinerClient(cc *grpc.ClientConn) MinerClient {
	return &minerClient{cc}
}

func (c *minerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Info(ctx context.Context, in *MinerInfoRequest, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Handshake(ctx context.Context, in *MinerHandshakeRequest, opts ...grpc.CallOption) (*MinerHandshakeReply, error) {
	out := new(MinerHandshakeReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Handshake", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Start(ctx context.Context, in *MinerStartRequest, opts ...grpc.CallOption) (*MinerStartReply, error) {
	out := new(MinerStartReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Stop(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error) {
	out := new(StopTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) TasksStatus(ctx context.Context, opts ...grpc.CallOption) (Miner_TasksStatusClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Miner_serviceDesc.Streams[0], c.cc, "/sonm.Miner/TasksStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &minerTasksStatusClient{stream}
	return x, nil
}

type Miner_TasksStatusClient interface {
	Send(*MinerStatusMapRequest) error
	Recv() (*StatusMapReply, error)
	grpc.ClientStream
}

type minerTasksStatusClient struct {
	grpc.ClientStream
}

func (x *minerTasksStatusClient) Send(m *MinerStatusMapRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *minerTasksStatusClient) Recv() (*StatusMapReply, error) {
	m := new(StatusMapReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *minerClient) TaskDetails(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error) {
	out := new(TaskStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/TaskDetails", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Miner_TaskLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Miner_serviceDesc.Streams[1], c.cc, "/sonm.Miner/TaskLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &minerTaskLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Miner_TaskLogsClient interface {
	Recv() (*TaskLogsChunk, error)
	grpc.ClientStream
}

type minerTaskLogsClient struct {
	grpc.ClientStream
}

func (x *minerTaskLogsClient) Recv() (*TaskLogsChunk, error) {
	m := new(TaskLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Miner service

type MinerServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Info(context.Context, *MinerInfoRequest) (*InfoReply, error)
	Handshake(context.Context, *MinerHandshakeRequest) (*MinerHandshakeReply, error)
	Start(context.Context, *MinerStartRequest) (*MinerStartReply, error)
	Stop(context.Context, *StopTaskRequest) (*StopTaskReply, error)
	TasksStatus(Miner_TasksStatusServer) error
	TaskDetails(context.Context, *TaskStatusRequest) (*TaskStatusReply, error)
	TaskLogs(*TaskLogsRequest, Miner_TaskLogsServer) error
}

func RegisterMinerServer(s *grpc.Server, srv MinerServer) {
	s.RegisterService(&_Miner_serviceDesc, srv)
}

func _Miner_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Info(ctx, req.(*MinerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerHandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Handshake(ctx, req.(*MinerHandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Start(ctx, req.(*MinerStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Stop(ctx, req.(*StopTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_TasksStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MinerServer).TasksStatus(&minerTasksStatusServer{stream})
}

type Miner_TasksStatusServer interface {
	Send(*StatusMapReply) error
	Recv() (*MinerStatusMapRequest, error)
	grpc.ServerStream
}

type minerTasksStatusServer struct {
	grpc.ServerStream
}

func (x *minerTasksStatusServer) Send(m *StatusMapReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *minerTasksStatusServer) Recv() (*MinerStatusMapRequest, error) {
	m := new(MinerStatusMapRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Miner_TaskDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).TaskDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/TaskDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).TaskDetails(ctx, req.(*TaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_TaskLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MinerServer).TaskLogs(m, &minerTaskLogsServer{stream})
}

type Miner_TaskLogsServer interface {
	Send(*TaskLogsChunk) error
	grpc.ServerStream
}

type minerTaskLogsServer struct {
	grpc.ServerStream
}

func (x *minerTaskLogsServer) Send(m *TaskLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

var _Miner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Miner",
	HandlerType: (*MinerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Miner_Ping_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Miner_Info_Handler,
		},
		{
			MethodName: "Handshake",
			Handler:    _Miner_Handshake_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Miner_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Miner_Stop_Handler,
		},
		{
			MethodName: "TaskDetails",
			Handler:    _Miner_TaskDetails_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TasksStatus",
			Handler:       _Miner_TasksStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TaskLogs",
			Handler:       _Miner_TaskLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "miner.proto",
}

func init() { proto.RegisterFile("miner.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x54, 0x5d, 0x6f, 0xda, 0x4a,
	0x10, 0xc5, 0x60, 0x87, 0x30, 0x5c, 0x82, 0x33, 0xf9, 0xf2, 0xf5, 0xcd, 0x03, 0xb2, 0xae, 0x54,
	0x1a, 0x55, 0x28, 0xa5, 0x55, 0xd4, 0xe6, 0xa5, 0x4a, 0xc0, 0x51, 0x51, 0xf8, 0xb0, 0x0c, 0x51,
	0xd4, 0xa7, 0x68, 0x21, 0x2e, 0x58, 0x10, 0xdb, 0xb5, 0xd7, 0x95, 0xfc, 0x0b, 0xfa, 0xd0, 0xfe,
	0xcd, 0x3e, 0xf4, 0x5f, 0x54, 0xbb, 0x6b, 0x13, 0x87, 0xe4, 0x89, 0x99, 0x33, 0xe7, 0x8c, 0xcf,
	0xce, 0xec, 0x02, 0xd5, 0x07, 0xd7, 0x73, 0xc2, 0x56, 0x10, 0xfa, 0xd4, 0x47, 0x39, 0xf2, 0xbd,
	0x07, 0xbd, 0xee, 0x7a, 0xec, 0xd7, 0x73, 0x89, 0x80, 0x75, 0x9c, 0x91, 0x80, 0x4c, 0xdd, 0x95,
	0x4b, 0x5d, 0x27, 0x12, 0x98, 0x81, 0xa0, 0x0e, 0x98, 0xb2, 0xe7, 0x7d, 0xf5, 0x6d, 0xe7, 0x5b,
	0xec, 0x44, 0xd4, 0x78, 0x0d, 0x07, 0x1c, 0xfb, 0x4c, 0xbc, 0xfb, 0x68, 0x41, 0x96, 0x4e, 0x5a,
	0x40, 0x15, 0x4a, 0x8b, 0x78, 0xaa, 0x49, 0x0d, 0xa9, 0x59, 0xb1, 0x59, 0x68, 0xfc, 0x92, 0x60,
	0x6f, 0x93, 0x1b, 0xac, 0x12, 0xdc, 0x07, 0x85, 0x1b, 0x4a, 0xb9, 0x22, 0xc1, 0x33, 0xf8, 0x27,
	0x6f, 0x41, 0x2b, 0x36, 0xa4, 0x66, 0xb5, 0x8d, 0x2d, 0x66, 0xb3, 0xd5, 0xc9, 0x55, 0xec, 0x27,
	0x3c, 0x7c, 0x05, 0x65, 0x8f, 0xd0, 0x49, 0x12, 0x38, 0x5a, 0xa9, 0x21, 0x35, 0x77, 0xda, 0x35,
	0x21, 0x19, 0x5e, 0x4c, 0x18, 0x68, 0x67, 0x55, 0xe3, 0x67, 0x11, 0x76, 0xb9, 0x9d, 0x31, 0x25,
	0x21, 0xcd, 0x6c, 0xef, 0x40, 0xd1, 0xbd, 0x4f, 0x9d, 0x14, 0xdd, 0x7b, 0xd4, 0x61, 0x3b, 0x74,
	0xe6, 0x6e, 0x44, 0xc3, 0x84, 0x5b, 0xa8, 0xd8, 0xeb, 0x9c, 0x19, 0x77, 0x1f, 0xc8, 0x5c, 0x7c,
	0xa8, 0x62, 0x8b, 0x04, 0x11, 0x64, 0x12, 0xd3, 0x85, 0x26, 0x73, 0x90, 0xc7, 0x78, 0x09, 0xb5,
	0xd0, 0x89, 0xd8, 0x77, 0x2c, 0x7f, 0xe5, 0xce, 0x12, 0x4d, 0xe1, 0xa7, 0x39, 0x4e, 0x4f, 0xe3,
	0x7b, 0x94, 0x30, 0x27, 0x76, 0x9e, 0x63, 0x3f, 0x95, 0xe0, 0x19, 0x54, 0x42, 0x27, 0xf2, 0xe3,
	0x70, 0xe6, 0x44, 0xda, 0x16, 0xd7, 0x6b, 0xcf, 0xf5, 0xa2, 0x6e, 0x3f, 0x52, 0xf1, 0x7f, 0xa8,
	0x59, 0xf1, 0x74, 0xe5, 0xce, 0xae, 0x9d, 0xa4, 0x4b, 0x28, 0xd1, 0xca, 0xdc, 0x58, 0x2d, 0xc8,
	0x83, 0xc6, 0x1f, 0x09, 0xea, 0xf9, 0x69, 0xb0, 0xc5, 0x1c, 0x43, 0x65, 0x96, 0xb5, 0x4e, 0x47,
	0xf2, 0x08, 0xe0, 0x19, 0x28, 0x96, 0x1f, 0x52, 0xb6, 0x99, 0x52, 0xb3, 0xda, 0x6e, 0x08, 0x2f,
	0x1b, 0x3d, 0x5a, 0x9c, 0x62, 0x7a, 0x34, 0x4c, 0x6c, 0x25, 0x60, 0xb1, 0x7e, 0x02, 0x32, 0x0b,
	0xd8, 0xa4, 0x7b, 0xd6, 0x7a, 0xd2, 0x16, 0x9b, 0x1b, 0xc3, 0xd3, 0x29, 0xf3, 0x58, 0x9f, 0x00,
	0x3c, 0x36, 0x60, 0x57, 0x6a, 0xe9, 0x24, 0xd9, 0x95, 0x5a, 0x3a, 0x09, 0x9e, 0x82, 0xf2, 0x9d,
	0xac, 0x62, 0x27, 0xbd, 0x1d, 0xfa, 0xcb, 0x1e, 0x58, 0x2b, 0x5b, 0x10, 0xcf, 0x8b, 0x1f, 0x24,
	0xe3, 0x28, 0xbd, 0xb3, 0x63, 0x4a, 0x68, 0x1c, 0x0d, 0x48, 0x90, 0x2e, 0xff, 0xe4, 0x87, 0x04,
	0xe5, 0xf4, 0x9e, 0xe0, 0x36, 0xc8, 0xc3, 0xd1, 0xd0, 0x54, 0x0b, 0x58, 0x85, 0xf2, 0x65, 0x7f,
	0xd4, 0xb9, 0x36, 0xbb, 0xaa, 0xc4, 0xe0, 0xab, 0x9b, 0x7e, 0x5f, 0x2d, 0x62, 0x0d, 0x2a, 0xe3,
	0x2f, 0x83, 0x81, 0x39, 0xb1, 0x7b, 0x1d, 0xb5, 0x84, 0x3b, 0x00, 0xb6, 0x39, 0x66, 0xc9, 0xc4,
	0xec, 0xaa, 0x32, 0xee, 0x41, 0xdd, 0x1a, 0xd9, 0x93, 0xbb, 0x1c, 0xa8, 0xa0, 0x0e, 0x87, 0x6b,
	0xcd, 0xdd, 0x4d, 0xd7, 0xba, 0xbb, 0xea, 0xd9, 0xe6, 0xed, 0x45, 0xbf, 0xaf, 0x6e, 0xb1, 0xcf,
	0xdc, 0x0c, 0xaf, 0x87, 0xa3, 0xdb, 0xa1, 0x5a, 0x6e, 0xff, 0x2e, 0x81, 0xc2, 0x3d, 0xe2, 0x1b,
	0x90, 0x2d, 0xd7, 0x9b, 0xe3, 0xae, 0x38, 0x1b, 0x8b, 0x53, 0xbb, 0x7a, 0x3d, 0x0f, 0x05, 0xab,
	0xc4, 0x28, 0xe0, 0x5b, 0x90, 0xd9, 0xeb, 0xc4, 0xc3, 0xdc, 0x24, 0x72, 0xcf, 0x35, 0x93, 0x08,
	0x48, 0x48, 0x4c, 0xa8, 0xac, 0x1f, 0x24, 0xfe, 0x97, 0xd3, 0x6d, 0x3e, 0x69, 0xfd, 0xdf, 0x97,
	0x8b, 0xa2, 0xcd, 0x47, 0x50, 0xf8, 0xc8, 0xf1, 0xe8, 0xf9, 0x12, 0x84, 0xfc, 0xe0, 0xc5, 0xed,
	0x18, 0x05, 0x7c, 0x0f, 0xf2, 0x98, 0xfa, 0x01, 0xa6, 0x04, 0x16, 0x4f, 0x48, 0xb4, 0xcc, 0x74,
	0x7b, 0x9b, 0xb0, 0x50, 0x5d, 0x41, 0x95, 0xa5, 0x91, 0xd8, 0xe2, 0x13, 0xe7, 0x9b, 0x8b, 0xd5,
	0xf7, 0xb3, 0x16, 0x6b, 0x9c, 0xf7, 0x68, 0x4a, 0xa7, 0x12, 0x7e, 0x12, 0x7d, 0xba, 0x0e, 0x25,
	0xee, 0x2a, 0xca, 0xec, 0x33, 0x48, 0xd0, 0x37, 0xec, 0xe7, 0x0b, 0xc2, 0xc8, 0x39, 0x6c, 0x33,
	0xb0, 0xef, 0xcf, 0x23, 0xcc, 0x91, 0x58, 0xbe, 0x71, 0x84, 0x0c, 0xee, 0x2c, 0x62, 0x6f, 0x69,
	0x14, 0x4e, 0xa5, 0xe9, 0x16, 0xff, 0x67, 0x7d, 0xf7, 0x37, 0x00, 0x00, 0xff, 0xff, 0xca, 0x51,
	0x90, 0xf7, 0x93, 0x05, 0x00, 0x00,
}
