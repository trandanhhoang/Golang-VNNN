// Code generated by protoc-gen-go. DO NOT EDIT.
// source: weather.proto

package weather

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

type WeatherRespond struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherRespond) Reset()         { *m = WeatherRespond{} }
func (m *WeatherRespond) String() string { return proto.CompactTextString(m) }
func (*WeatherRespond) ProtoMessage()    {}
func (*WeatherRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{0}
}

func (m *WeatherRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherRespond.Unmarshal(m, b)
}
func (m *WeatherRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherRespond.Marshal(b, m, deterministic)
}
func (m *WeatherRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherRespond.Merge(m, src)
}
func (m *WeatherRespond) XXX_Size() int {
	return xxx_messageInfo_WeatherRespond.Size(m)
}
func (m *WeatherRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherRespond.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherRespond proto.InternalMessageInfo

func (m *WeatherRespond) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WeatherRespond) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type WeatherRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeatherRequest) Reset()         { *m = WeatherRequest{} }
func (m *WeatherRequest) String() string { return proto.CompactTextString(m) }
func (*WeatherRequest) ProtoMessage()    {}
func (*WeatherRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_231dcd72b885f4be, []int{1}
}

func (m *WeatherRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeatherRequest.Unmarshal(m, b)
}
func (m *WeatherRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeatherRequest.Marshal(b, m, deterministic)
}
func (m *WeatherRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeatherRequest.Merge(m, src)
}
func (m *WeatherRequest) XXX_Size() int {
	return xxx_messageInfo_WeatherRequest.Size(m)
}
func (m *WeatherRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WeatherRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WeatherRequest proto.InternalMessageInfo

func (m *WeatherRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*WeatherRespond)(nil), "weather.WeatherRespond")
	proto.RegisterType((*WeatherRequest)(nil), "weather.WeatherRequest")
}

func init() { proto.RegisterFile("weather.proto", fileDescriptor_231dcd72b885f4be) }

var fileDescriptor_231dcd72b885f4be = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4f, 0x4d, 0x2c,
	0xc9, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x2c, 0xb8,
	0xf8, 0xc2, 0x21, 0xcc, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x1b, 0x24, 0x96, 0x92, 0x58,
	0x92, 0x28, 0xc1, 0x04, 0x11, 0x03, 0xb1, 0x95, 0x54, 0x90, 0x74, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x60, 0xd3, 0x69, 0x14, 0x04, 0x57, 0x15, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xc0,
	0xc5, 0xe5, 0x9e, 0x5a, 0x02, 0x15, 0x14, 0x12, 0xd7, 0x83, 0x39, 0x0c, 0xd5, 0x30, 0x29, 0x2c,
	0x12, 0x60, 0xf7, 0x29, 0x31, 0x24, 0xb1, 0x81, 0xfd, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0xf3, 0x80, 0xca, 0xd4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WeatherServiceClient interface {
	GetWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherRespond, error)
}

type weatherServiceClient struct {
	cc *grpc.ClientConn
}

func NewWeatherServiceClient(cc *grpc.ClientConn) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) GetWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherRespond, error) {
	out := new(WeatherRespond)
	err := c.cc.Invoke(ctx, "/weather.WeatherService/GetWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServiceServer is the server API for WeatherService service.
type WeatherServiceServer interface {
	GetWeather(context.Context, *WeatherRequest) (*WeatherRespond, error)
}

// UnimplementedWeatherServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWeatherServiceServer struct {
}

func (*UnimplementedWeatherServiceServer) GetWeather(ctx context.Context, req *WeatherRequest) (*WeatherRespond, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeather not implemented")
}

func RegisterWeatherServiceServer(s *grpc.Server, srv WeatherServiceServer) {
	s.RegisterService(&_WeatherService_serviceDesc, srv)
}

func _WeatherService_GetWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.WeatherService/GetWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetWeather(ctx, req.(*WeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WeatherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weather.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWeather",
			Handler:    _WeatherService_GetWeather_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weather.proto",
}
