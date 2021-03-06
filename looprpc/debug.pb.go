// Code generated by protoc-gen-go. DO NOT EDIT.
// source: debug.proto

package looprpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ForceAutoLoopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForceAutoLoopRequest) Reset()         { *m = ForceAutoLoopRequest{} }
func (m *ForceAutoLoopRequest) String() string { return proto.CompactTextString(m) }
func (*ForceAutoLoopRequest) ProtoMessage()    {}
func (*ForceAutoLoopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{0}
}

func (m *ForceAutoLoopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForceAutoLoopRequest.Unmarshal(m, b)
}
func (m *ForceAutoLoopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForceAutoLoopRequest.Marshal(b, m, deterministic)
}
func (m *ForceAutoLoopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForceAutoLoopRequest.Merge(m, src)
}
func (m *ForceAutoLoopRequest) XXX_Size() int {
	return xxx_messageInfo_ForceAutoLoopRequest.Size(m)
}
func (m *ForceAutoLoopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForceAutoLoopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForceAutoLoopRequest proto.InternalMessageInfo

type ForceAutoLoopResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForceAutoLoopResponse) Reset()         { *m = ForceAutoLoopResponse{} }
func (m *ForceAutoLoopResponse) String() string { return proto.CompactTextString(m) }
func (*ForceAutoLoopResponse) ProtoMessage()    {}
func (*ForceAutoLoopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d9d361be58531fb, []int{1}
}

func (m *ForceAutoLoopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForceAutoLoopResponse.Unmarshal(m, b)
}
func (m *ForceAutoLoopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForceAutoLoopResponse.Marshal(b, m, deterministic)
}
func (m *ForceAutoLoopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForceAutoLoopResponse.Merge(m, src)
}
func (m *ForceAutoLoopResponse) XXX_Size() int {
	return xxx_messageInfo_ForceAutoLoopResponse.Size(m)
}
func (m *ForceAutoLoopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ForceAutoLoopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ForceAutoLoopResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ForceAutoLoopRequest)(nil), "looprpc.ForceAutoLoopRequest")
	proto.RegisterType((*ForceAutoLoopResponse)(nil), "looprpc.ForceAutoLoopResponse")
}

func init() { proto.RegisterFile("debug.proto", fileDescriptor_8d9d361be58531fb) }

var fileDescriptor_8d9d361be58531fb = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x49, 0x4d, 0x2a,
	0x4d, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0xc9, 0xcf, 0x2f, 0x28, 0x2a, 0x48,
	0x56, 0x12, 0xe3, 0x12, 0x71, 0xcb, 0x2f, 0x4a, 0x4e, 0x75, 0x2c, 0x2d, 0xc9, 0xf7, 0xc9, 0xcf,
	0x2f, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x12, 0xe7, 0x12, 0x45, 0x13, 0x2f, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x35, 0x8a, 0xe4, 0x62, 0x75, 0x01, 0x19, 0x24, 0x14, 0xc0, 0xc5, 0x8b,
	0xa2, 0x42, 0x48, 0x56, 0x0f, 0x6a, 0xa8, 0x1e, 0x36, 0x13, 0xa5, 0xe4, 0x70, 0x49, 0x43, 0x0c,
	0x56, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x39, 0x7a, 0x3c,
	0x6b, 0xaa, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DebugClient is the client API for Debug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DebugClient interface {
	//
	//ForceAutoLoop is intended for *testing purposes only* and will not work on
	//mainnet. This endpoint ticks our autoloop timer, triggering automated
	//dispatch of a swap if one is suggested.
	ForceAutoLoop(ctx context.Context, in *ForceAutoLoopRequest, opts ...grpc.CallOption) (*ForceAutoLoopResponse, error)
}

type debugClient struct {
	cc *grpc.ClientConn
}

func NewDebugClient(cc *grpc.ClientConn) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) ForceAutoLoop(ctx context.Context, in *ForceAutoLoopRequest, opts ...grpc.CallOption) (*ForceAutoLoopResponse, error) {
	out := new(ForceAutoLoopResponse)
	err := c.cc.Invoke(ctx, "/looprpc.Debug/ForceAutoLoop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugServer is the server API for Debug service.
type DebugServer interface {
	//
	//ForceAutoLoop is intended for *testing purposes only* and will not work on
	//mainnet. This endpoint ticks our autoloop timer, triggering automated
	//dispatch of a swap if one is suggested.
	ForceAutoLoop(context.Context, *ForceAutoLoopRequest) (*ForceAutoLoopResponse, error)
}

// UnimplementedDebugServer can be embedded to have forward compatible implementations.
type UnimplementedDebugServer struct {
}

func (*UnimplementedDebugServer) ForceAutoLoop(ctx context.Context, req *ForceAutoLoopRequest) (*ForceAutoLoopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceAutoLoop not implemented")
}

func RegisterDebugServer(s *grpc.Server, srv DebugServer) {
	s.RegisterService(&_Debug_serviceDesc, srv)
}

func _Debug_ForceAutoLoop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForceAutoLoopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).ForceAutoLoop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.Debug/ForceAutoLoop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).ForceAutoLoop(ctx, req.(*ForceAutoLoopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Debug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "looprpc.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForceAutoLoop",
			Handler:    _Debug_ForceAutoLoop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "debug.proto",
}
