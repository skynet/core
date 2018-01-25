// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketplace.proto

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

type GetOrdersRequest struct {
	// Order keeps slot and type for searching
	Order *Order `protobuf:"bytes,1,opt,name=order" json:"order,omitempty"`
	// Count describe how namy results must be returned (order by price)
	Count uint64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *GetOrdersRequest) Reset()                    { *m = GetOrdersRequest{} }
func (m *GetOrdersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOrdersRequest) ProtoMessage()               {}
func (*GetOrdersRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *GetOrdersRequest) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *GetOrdersRequest) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetOrdersReply struct {
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *GetOrdersReply) Reset()                    { *m = GetOrdersReply{} }
func (m *GetOrdersReply) String() string            { return proto.CompactTextString(m) }
func (*GetOrdersReply) ProtoMessage()               {}
func (*GetOrdersReply) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *GetOrdersReply) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

// GetProcessingReply represent Node's local
// orders processing status
type GetProcessingReply struct {
	Orders map[string]*GetProcessingReply_ProcessedOrder `protobuf:"bytes,2,rep,name=orders" json:"orders,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GetProcessingReply) Reset()                    { *m = GetProcessingReply{} }
func (m *GetProcessingReply) String() string            { return proto.CompactTextString(m) }
func (*GetProcessingReply) ProtoMessage()               {}
func (*GetProcessingReply) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *GetProcessingReply) GetOrders() map[string]*GetProcessingReply_ProcessedOrder {
	if m != nil {
		return m.Orders
	}
	return nil
}

type GetProcessingReply_ProcessedOrder struct {
	Id        string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Status    uint32     `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	Timestamp *Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Extra     string     `protobuf:"bytes,4,opt,name=extra" json:"extra,omitempty"`
}

func (m *GetProcessingReply_ProcessedOrder) Reset()         { *m = GetProcessingReply_ProcessedOrder{} }
func (m *GetProcessingReply_ProcessedOrder) String() string { return proto.CompactTextString(m) }
func (*GetProcessingReply_ProcessedOrder) ProtoMessage()    {}
func (*GetProcessingReply_ProcessedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{2, 0}
}

func (m *GetProcessingReply_ProcessedOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetProcessingReply_ProcessedOrder) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetProcessingReply_ProcessedOrder) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *GetProcessingReply_ProcessedOrder) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

type TouchOrdersRequest struct {
	IDs []string `protobuf:"bytes,1,rep,name=IDs" json:"IDs,omitempty"`
}

func (m *TouchOrdersRequest) Reset()                    { *m = TouchOrdersRequest{} }
func (m *TouchOrdersRequest) String() string            { return proto.CompactTextString(m) }
func (*TouchOrdersRequest) ProtoMessage()               {}
func (*TouchOrdersRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *TouchOrdersRequest) GetIDs() []string {
	if m != nil {
		return m.IDs
	}
	return nil
}

func init() {
	proto.RegisterType((*GetOrdersRequest)(nil), "sonm.GetOrdersRequest")
	proto.RegisterType((*GetOrdersReply)(nil), "sonm.GetOrdersReply")
	proto.RegisterType((*GetProcessingReply)(nil), "sonm.GetProcessingReply")
	proto.RegisterType((*GetProcessingReply_ProcessedOrder)(nil), "sonm.GetProcessingReply.ProcessedOrder")
	proto.RegisterType((*TouchOrdersRequest)(nil), "sonm.TouchOrdersRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Market service

type MarketClient interface {
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersReply, error)
	GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error)
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	CancelOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Empty, error)
	TouchOrders(ctx context.Context, in *TouchOrdersRequest, opts ...grpc.CallOption) (*Empty, error)
	// GetProcessing returns currently processing orders
	// not required in Marketplace service
	// must be implemented on Node
	GetProcessing(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetProcessingReply, error)
}

type marketClient struct {
	cc *grpc.ClientConn
}

func NewMarketClient(cc *grpc.ClientConn) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersReply, error) {
	out := new(GetOrdersReply)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrderByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/CreateOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CancelOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.Market/CancelOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) TouchOrders(ctx context.Context, in *TouchOrdersRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.Market/TouchOrders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetProcessing(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetProcessingReply, error) {
	out := new(GetProcessingReply)
	err := grpc.Invoke(ctx, "/sonm.Market/GetProcessing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Market service

type MarketServer interface {
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersReply, error)
	GetOrderByID(context.Context, *ID) (*Order, error)
	CreateOrder(context.Context, *Order) (*Order, error)
	CancelOrder(context.Context, *Order) (*Empty, error)
	TouchOrders(context.Context, *TouchOrdersRequest) (*Empty, error)
	// GetProcessing returns currently processing orders
	// not required in Marketplace service
	// must be implemented on Node
	GetProcessing(context.Context, *Empty) (*GetProcessingReply, error)
}

func RegisterMarketServer(s *grpc.Server, srv MarketServer) {
	s.RegisterService(&_Market_serviceDesc, srv)
}

func _Market_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrderByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CancelOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_TouchOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TouchOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).TouchOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/TouchOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).TouchOrders(ctx, req.(*TouchOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetProcessing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetProcessing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetProcessing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetProcessing(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Market_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrders",
			Handler:    _Market_GetOrders_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _Market_GetOrderByID_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Market_CreateOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Market_CancelOrder_Handler,
		},
		{
			MethodName: "TouchOrders",
			Handler:    _Market_TouchOrders_Handler,
		},
		{
			MethodName: "GetProcessing",
			Handler:    _Market_GetProcessing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketplace.proto",
}

func init() { proto.RegisterFile("marketplace.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x9d, 0x0f, 0xe1, 0x31, 0x4d, 0xc3, 0xa8, 0xaa, 0x2c, 0x9f, 0x82, 0x41, 0xb4, 0x1c,
	0xc8, 0x21, 0x08, 0x54, 0xf1, 0x71, 0x81, 0x54, 0x55, 0x84, 0x10, 0x68, 0xd5, 0x3f, 0xb0, 0xb1,
	0x47, 0x60, 0xd5, 0x5f, 0xec, 0xae, 0x11, 0x3e, 0xf0, 0xb3, 0xe0, 0xf7, 0x21, 0xef, 0xda, 0xc6,
	0x4e, 0x9a, 0xdb, 0xce, 0x9b, 0xf7, 0xde, 0x64, 0xde, 0xc4, 0xf0, 0x28, 0xe5, 0xe2, 0x8e, 0x54,
	0x91, 0xf0, 0x90, 0x56, 0x85, 0xc8, 0x55, 0x8e, 0x13, 0x99, 0x67, 0xa9, 0xef, 0xec, 0xe2, 0xc8,
	0x00, 0xfe, 0x69, 0x9c, 0xd5, 0x50, 0x16, 0xf3, 0x16, 0x50, 0x71, 0x4a, 0x52, 0xf1, 0xb4, 0x30,
	0x40, 0xf0, 0x09, 0x16, 0x37, 0xa4, 0xbe, 0x88, 0x88, 0x84, 0x64, 0xf4, 0xa3, 0x24, 0xa9, 0xf0,
	0x31, 0x4c, 0xf3, 0x1a, 0xf0, 0xac, 0xa5, 0x75, 0xe9, 0xae, 0xdd, 0x55, 0xed, 0xb1, 0xd2, 0x1c,
	0x66, 0x3a, 0x78, 0x06, 0xd3, 0x30, 0x2f, 0x33, 0xe5, 0xd9, 0x4b, 0xeb, 0x72, 0xc2, 0x4c, 0x11,
	0xbc, 0x82, 0x79, 0xcf, 0xac, 0x48, 0x2a, 0x7c, 0x02, 0x33, 0x2d, 0x90, 0x9e, 0xb5, 0x1c, 0xef,
	0x7b, 0x35, 0xad, 0xe0, 0x8f, 0x0d, 0x78, 0x43, 0xea, 0xab, 0xc8, 0x43, 0x92, 0x32, 0xce, 0xbe,
	0x19, 0xed, 0xbb, 0x4e, 0x6b, 0x6b, 0xed, 0x53, 0xa3, 0x3d, 0x64, 0x1a, 0x3b, 0x79, 0x9d, 0x29,
	0x51, 0xb5, 0xa6, 0xfe, 0x6f, 0x98, 0x37, 0x34, 0x8a, 0x74, 0x1f, 0xe7, 0x60, 0xc7, 0x91, 0xde,
	0xc9, 0x61, 0x76, 0x1c, 0xe1, 0x39, 0xcc, 0xa4, 0xe2, 0xaa, 0x94, 0x7a, 0x89, 0x13, 0xd6, 0x54,
	0xf8, 0x02, 0x9c, 0x2e, 0x25, 0x6f, 0xac, 0x23, 0x38, 0x35, 0xa3, 0x6f, 0x5b, 0x98, 0xfd, 0x67,
	0xd4, 0x51, 0xd0, 0x2f, 0x25, 0xb8, 0x37, 0xd1, 0xce, 0xa6, 0xf0, 0x77, 0xe0, 0xf6, 0x7e, 0x15,
	0x2e, 0x60, 0x7c, 0x47, 0x55, 0x33, 0xbc, 0x7e, 0xe2, 0x7b, 0x98, 0xfe, 0xe4, 0x49, 0x49, 0x7a,
	0xb8, 0xbb, 0xbe, 0x38, 0xba, 0xdc, 0x70, 0x0b, 0x66, 0x54, 0x6f, 0xec, 0x2b, 0x2b, 0x78, 0x06,
	0x78, 0x9b, 0x97, 0xe1, 0xf7, 0xe1, 0xf5, 0x16, 0x30, 0xde, 0x6e, 0x4c, 0xde, 0x0e, 0xab, 0x9f,
	0xeb, 0xbf, 0x36, 0xcc, 0x3e, 0xeb, 0x3f, 0x0b, 0xbe, 0x05, 0xa7, 0xbb, 0x10, 0x9e, 0x77, 0x33,
	0x07, 0x0e, 0xfe, 0xd9, 0x01, 0x5e, 0x24, 0x55, 0x30, 0xc2, 0x0b, 0x78, 0xd8, 0x62, 0x1f, 0xaa,
	0xed, 0x06, 0x1f, 0x18, 0xde, 0x76, 0xe3, 0xf7, 0xcf, 0x1a, 0x8c, 0xf0, 0x39, 0xb8, 0x1f, 0x05,
	0x71, 0x45, 0x26, 0xf8, 0x7e, 0xf7, 0x3e, 0x2a, 0xcf, 0x42, 0x4a, 0x8e, 0x53, 0xaf, 0xd3, 0x42,
	0xd5, 0xe3, 0x5f, 0x83, 0xdb, 0x5b, 0x17, 0xbd, 0xe6, 0x26, 0x07, 0x09, 0xec, 0xeb, 0xae, 0xe0,
	0x64, 0x10, 0x2b, 0xf6, 0xfb, 0xbe, 0x77, 0x2c, 0xf8, 0x60, 0xb4, 0x9b, 0xe9, 0x6f, 0xe4, 0xe5,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0x07, 0xbc, 0xeb, 0x6b, 0x03, 0x00, 0x00,
}
