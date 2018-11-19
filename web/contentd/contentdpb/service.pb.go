// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package contentdpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NewTweetRequest struct {
	Content              string   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	PosterUID            string   `protobuf:"bytes,2,opt,name=PosterUID,proto3" json:"PosterUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewTweetRequest) Reset()         { *m = NewTweetRequest{} }
func (m *NewTweetRequest) String() string { return proto.CompactTextString(m) }
func (*NewTweetRequest) ProtoMessage()    {}
func (*NewTweetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *NewTweetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTweetRequest.Unmarshal(m, b)
}
func (m *NewTweetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTweetRequest.Marshal(b, m, deterministic)
}
func (m *NewTweetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTweetRequest.Merge(m, src)
}
func (m *NewTweetRequest) XXX_Size() int {
	return xxx_messageInfo_NewTweetRequest.Size(m)
}
func (m *NewTweetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTweetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewTweetRequest proto.InternalMessageInfo

func (m *NewTweetRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *NewTweetRequest) GetPosterUID() string {
	if m != nil {
		return m.PosterUID
	}
	return ""
}

type NewTweetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewTweetResponse) Reset()         { *m = NewTweetResponse{} }
func (m *NewTweetResponse) String() string { return proto.CompactTextString(m) }
func (*NewTweetResponse) ProtoMessage()    {}
func (*NewTweetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *NewTweetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTweetResponse.Unmarshal(m, b)
}
func (m *NewTweetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTweetResponse.Marshal(b, m, deterministic)
}
func (m *NewTweetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTweetResponse.Merge(m, src)
}
func (m *NewTweetResponse) XXX_Size() int {
	return xxx_messageInfo_NewTweetResponse.Size(m)
}
func (m *NewTweetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTweetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewTweetResponse proto.InternalMessageInfo

type GetTweetRequest struct {
	TID                  string   `protobuf:"bytes,1,opt,name=TID,proto3" json:"TID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTweetRequest) Reset()         { *m = GetTweetRequest{} }
func (m *GetTweetRequest) String() string { return proto.CompactTextString(m) }
func (*GetTweetRequest) ProtoMessage()    {}
func (*GetTweetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *GetTweetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTweetRequest.Unmarshal(m, b)
}
func (m *GetTweetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTweetRequest.Marshal(b, m, deterministic)
}
func (m *GetTweetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTweetRequest.Merge(m, src)
}
func (m *GetTweetRequest) XXX_Size() int {
	return xxx_messageInfo_GetTweetRequest.Size(m)
}
func (m *GetTweetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTweetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTweetRequest proto.InternalMessageInfo

func (m *GetTweetRequest) GetTID() string {
	if m != nil {
		return m.TID
	}
	return ""
}

type GetTweetResponse struct {
	Tweet                *Tweet   `protobuf:"bytes,1,opt,name=Tweet,proto3" json:"Tweet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTweetResponse) Reset()         { *m = GetTweetResponse{} }
func (m *GetTweetResponse) String() string { return proto.CompactTextString(m) }
func (*GetTweetResponse) ProtoMessage()    {}
func (*GetTweetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *GetTweetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTweetResponse.Unmarshal(m, b)
}
func (m *GetTweetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTweetResponse.Marshal(b, m, deterministic)
}
func (m *GetTweetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTweetResponse.Merge(m, src)
}
func (m *GetTweetResponse) XXX_Size() int {
	return xxx_messageInfo_GetTweetResponse.Size(m)
}
func (m *GetTweetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTweetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTweetResponse proto.InternalMessageInfo

func (m *GetTweetResponse) GetTweet() *Tweet {
	if m != nil {
		return m.Tweet
	}
	return nil
}

type GetTweetsByUserRequest struct {
	UID                  []string `protobuf:"bytes,1,rep,name=UID,proto3" json:"UID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTweetsByUserRequest) Reset()         { *m = GetTweetsByUserRequest{} }
func (m *GetTweetsByUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetTweetsByUserRequest) ProtoMessage()    {}
func (*GetTweetsByUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *GetTweetsByUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTweetsByUserRequest.Unmarshal(m, b)
}
func (m *GetTweetsByUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTweetsByUserRequest.Marshal(b, m, deterministic)
}
func (m *GetTweetsByUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTweetsByUserRequest.Merge(m, src)
}
func (m *GetTweetsByUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetTweetsByUserRequest.Size(m)
}
func (m *GetTweetsByUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTweetsByUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTweetsByUserRequest proto.InternalMessageInfo

func (m *GetTweetsByUserRequest) GetUID() []string {
	if m != nil {
		return m.UID
	}
	return nil
}

type GetTweetsByUserResponse struct {
	Tweets               []*Tweet `protobuf:"bytes,1,rep,name=Tweets,proto3" json:"Tweets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTweetsByUserResponse) Reset()         { *m = GetTweetsByUserResponse{} }
func (m *GetTweetsByUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetTweetsByUserResponse) ProtoMessage()    {}
func (*GetTweetsByUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *GetTweetsByUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTweetsByUserResponse.Unmarshal(m, b)
}
func (m *GetTweetsByUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTweetsByUserResponse.Marshal(b, m, deterministic)
}
func (m *GetTweetsByUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTweetsByUserResponse.Merge(m, src)
}
func (m *GetTweetsByUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetTweetsByUserResponse.Size(m)
}
func (m *GetTweetsByUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTweetsByUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTweetsByUserResponse proto.InternalMessageInfo

func (m *GetTweetsByUserResponse) GetTweets() []*Tweet {
	if m != nil {
		return m.Tweets
	}
	return nil
}

func init() {
	proto.RegisterType((*NewTweetRequest)(nil), "NewTweetRequest")
	proto.RegisterType((*NewTweetResponse)(nil), "NewTweetResponse")
	proto.RegisterType((*GetTweetRequest)(nil), "GetTweetRequest")
	proto.RegisterType((*GetTweetResponse)(nil), "GetTweetResponse")
	proto.RegisterType((*GetTweetsByUserRequest)(nil), "GetTweetsByUserRequest")
	proto.RegisterType((*GetTweetsByUserResponse)(nil), "GetTweetsByUserResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0x16, 0x6b, 0x3b, 0xa2, 0xd9, 0xce, 0xc1, 0x86, 0xe0, 0x41, 0x56, 0x44, 0xf1, 0x30,
	0x42, 0x3d, 0x79, 0xad, 0x01, 0xc9, 0x45, 0xa4, 0x34, 0x2f, 0x60, 0x3b, 0x07, 0x2f, 0xd9, 0x9a,
	0x5d, 0x2d, 0xbe, 0x96, 0x4f, 0x28, 0xd9, 0x1f, 0x57, 0xd3, 0xf4, 0xb6, 0x33, 0xf3, 0xcd, 0x7c,
	0x3f, 0x0b, 0x27, 0x9a, 0x9b, 0xcf, 0xb7, 0x15, 0xd3, 0xa6, 0x51, 0x46, 0xe5, 0xa7, 0x2b, 0x55,
	0x1b, 0xae, 0xcd, 0xda, 0xd5, 0xb2, 0x84, 0xf4, 0x99, 0xb7, 0xcb, 0x2d, 0xb3, 0x59, 0xf0, 0xfb,
	0x07, 0x6b, 0x83, 0x19, 0x1c, 0x3d, 0x3a, 0x50, 0x96, 0x5c, 0x24, 0x37, 0xe3, 0x45, 0x28, 0xf1,
	0x1c, 0xc6, 0x2f, 0x4a, 0x1b, 0x6e, 0xaa, 0xb2, 0xc8, 0x0e, 0xec, 0x2c, 0x36, 0x24, 0x82, 0x88,
	0xa7, 0xf4, 0x46, 0xd5, 0x9a, 0xe5, 0x25, 0xa4, 0x4f, 0x6c, 0xfe, 0x9d, 0x17, 0x30, 0x58, 0x96,
	0x85, 0x3f, 0xdd, 0x3e, 0xe5, 0x03, 0x88, 0x08, 0x72, 0x8b, 0x78, 0x05, 0x87, 0xb6, 0x61, 0x71,
	0xc7, 0xb3, 0x94, 0x7e, 0x75, 0x3b, 0x9c, 0x9b, 0xca, 0x5b, 0x38, 0x0b, 0xab, 0x7a, 0xfe, 0x55,
	0x69, 0x6e, 0xfe, 0xd0, 0x54, 0x96, 0x66, 0xd0, 0xd2, 0xb4, 0xfa, 0xe6, 0x30, 0xdd, 0xc1, 0x7a,
	0xb6, 0x6b, 0x18, 0xba, 0xbe, 0xc5, 0xf7, 0xd0, 0xf9, 0xf1, 0xec, 0x3b, 0x81, 0x91, 0x4f, 0x63,
	0x8d, 0x77, 0x30, 0x0a, 0x86, 0x51, 0x50, 0x27, 0xc6, 0x7c, 0x42, 0xdd, 0x34, 0xda, 0x85, 0xa0,
	0x00, 0x05, 0x75, 0x82, 0xc9, 0x27, 0xb4, 0x93, 0x42, 0x11, 0xe3, 0xf3, 0x92, 0x71, 0x4a, 0xfd,
	0x86, 0xf3, 0x8c, 0xf6, 0xb8, 0x7b, 0x1d, 0xda, 0xaf, 0xbe, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xa0, 0xba, 0x5a, 0x5f, 0x0b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContentdClient is the client API for Contentd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentdClient interface {
	NewTweet(ctx context.Context, in *NewTweetRequest, opts ...grpc.CallOption) (*NewTweetResponse, error)
	GetTweet(ctx context.Context, in *GetTweetRequest, opts ...grpc.CallOption) (*GetTweetResponse, error)
	GetTweetsByUser(ctx context.Context, in *GetTweetsByUserRequest, opts ...grpc.CallOption) (*GetTweetsByUserResponse, error)
}

type contentdClient struct {
	cc *grpc.ClientConn
}

func NewContentdClient(cc *grpc.ClientConn) ContentdClient {
	return &contentdClient{cc}
}

func (c *contentdClient) NewTweet(ctx context.Context, in *NewTweetRequest, opts ...grpc.CallOption) (*NewTweetResponse, error) {
	out := new(NewTweetResponse)
	err := c.cc.Invoke(ctx, "/Contentd/NewTweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentdClient) GetTweet(ctx context.Context, in *GetTweetRequest, opts ...grpc.CallOption) (*GetTweetResponse, error) {
	out := new(GetTweetResponse)
	err := c.cc.Invoke(ctx, "/Contentd/GetTweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentdClient) GetTweetsByUser(ctx context.Context, in *GetTweetsByUserRequest, opts ...grpc.CallOption) (*GetTweetsByUserResponse, error) {
	out := new(GetTweetsByUserResponse)
	err := c.cc.Invoke(ctx, "/Contentd/GetTweetsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentdServer is the server API for Contentd service.
type ContentdServer interface {
	NewTweet(context.Context, *NewTweetRequest) (*NewTweetResponse, error)
	GetTweet(context.Context, *GetTweetRequest) (*GetTweetResponse, error)
	GetTweetsByUser(context.Context, *GetTweetsByUserRequest) (*GetTweetsByUserResponse, error)
}

func RegisterContentdServer(s *grpc.Server, srv ContentdServer) {
	s.RegisterService(&_Contentd_serviceDesc, srv)
}

func _Contentd_NewTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentdServer).NewTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contentd/NewTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentdServer).NewTweet(ctx, req.(*NewTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contentd_GetTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentdServer).GetTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contentd/GetTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentdServer).GetTweet(ctx, req.(*GetTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contentd_GetTweetsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTweetsByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentdServer).GetTweetsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contentd/GetTweetsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentdServer).GetTweetsByUser(ctx, req.(*GetTweetsByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contentd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Contentd",
	HandlerType: (*ContentdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewTweet",
			Handler:    _Contentd_NewTweet_Handler,
		},
		{
			MethodName: "GetTweet",
			Handler:    _Contentd_GetTweet_Handler,
		},
		{
			MethodName: "GetTweetsByUser",
			Handler:    _Contentd_GetTweetsByUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
