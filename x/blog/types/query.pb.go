// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// this line is used by starport scaffolding # 3
type QueryGetPostRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetPostRequest) Reset()         { *m = QueryGetPostRequest{} }
func (m *QueryGetPostRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetPostRequest) ProtoMessage()    {}
func (*QueryGetPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e22cd6d352f3384, []int{0}
}
func (m *QueryGetPostRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetPostRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetPostRequest.Merge(m, src)
}
func (m *QueryGetPostRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetPostRequest proto.InternalMessageInfo

func (m *QueryGetPostRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetPostResponse struct {
	Post *Post `protobuf:"bytes,1,opt,name=Post,proto3" json:"Post,omitempty"`
}

func (m *QueryGetPostResponse) Reset()         { *m = QueryGetPostResponse{} }
func (m *QueryGetPostResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetPostResponse) ProtoMessage()    {}
func (*QueryGetPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e22cd6d352f3384, []int{1}
}
func (m *QueryGetPostResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetPostResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetPostResponse.Merge(m, src)
}
func (m *QueryGetPostResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetPostResponse proto.InternalMessageInfo

func (m *QueryGetPostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type QueryAllPostRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPostRequest) Reset()         { *m = QueryAllPostRequest{} }
func (m *QueryAllPostRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllPostRequest) ProtoMessage()    {}
func (*QueryAllPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e22cd6d352f3384, []int{2}
}
func (m *QueryAllPostRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPostRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPostRequest.Merge(m, src)
}
func (m *QueryAllPostRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPostRequest proto.InternalMessageInfo

func (m *QueryAllPostRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllPostResponse struct {
	Post       []*Post             `protobuf:"bytes,1,rep,name=Post,proto3" json:"Post,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPostResponse) Reset()         { *m = QueryAllPostResponse{} }
func (m *QueryAllPostResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllPostResponse) ProtoMessage()    {}
func (*QueryAllPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e22cd6d352f3384, []int{3}
}
func (m *QueryAllPostResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPostResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPostResponse.Merge(m, src)
}
func (m *QueryAllPostResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPostResponse proto.InternalMessageInfo

func (m *QueryAllPostResponse) GetPost() []*Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *QueryAllPostResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetPostRequest)(nil), "user.planet.blog.QueryGetPostRequest")
	proto.RegisterType((*QueryGetPostResponse)(nil), "user.planet.blog.QueryGetPostResponse")
	proto.RegisterType((*QueryAllPostRequest)(nil), "user.planet.blog.QueryAllPostRequest")
	proto.RegisterType((*QueryAllPostResponse)(nil), "user.planet.blog.QueryAllPostResponse")
}

func init() { proto.RegisterFile("blog/query.proto", fileDescriptor_6e22cd6d352f3384) }

var fileDescriptor_6e22cd6d352f3384 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4f, 0xf2, 0x30,
	0x1c, 0xc7, 0xd9, 0x1e, 0x9e, 0xe7, 0x49, 0x6a, 0xa2, 0xa4, 0x1a, 0x42, 0x50, 0x17, 0x33, 0x02,
	0x1a, 0x0e, 0x6d, 0xc0, 0xb3, 0x07, 0x38, 0xc8, 0x15, 0x39, 0x9a, 0x78, 0xe8, 0xa0, 0x99, 0x4b,
	0xca, 0x5a, 0x68, 0x67, 0x40, 0xe3, 0xc5, 0xab, 0x17, 0x13, 0xdf, 0x94, 0x47, 0x12, 0x2f, 0x1e,
	0x0d, 0xf8, 0x36, 0x4c, 0xcc, 0xda, 0x05, 0x19, 0x7f, 0xc2, 0x71, 0xdb, 0xf7, 0xdb, 0xcf, 0xe7,
	0xf7, 0xeb, 0x40, 0xce, 0x63, 0xdc, 0xc7, 0x83, 0x88, 0x0e, 0xc7, 0x48, 0x0c, 0xb9, 0xe2, 0x30,
	0x17, 0x49, 0x3a, 0x44, 0x82, 0x91, 0x90, 0x2a, 0x14, 0x7f, 0x2d, 0x1e, 0xf9, 0x9c, 0xfb, 0x8c,
	0x62, 0x22, 0x02, 0x4c, 0xc2, 0x90, 0x2b, 0xa2, 0x02, 0x1e, 0x4a, 0x93, 0x2f, 0x56, 0xbb, 0x5c,
	0xf6, 0xb9, 0xc4, 0x1e, 0x91, 0xd4, 0x1c, 0x84, 0xef, 0x6a, 0x1e, 0x55, 0xa4, 0x86, 0x05, 0xf1,
	0x83, 0x50, 0x87, 0x93, 0xec, 0x9e, 0xa6, 0x09, 0x2e, 0x95, 0x79, 0xe1, 0x96, 0xc1, 0xfe, 0x55,
	0x5c, 0x69, 0x51, 0xd5, 0xe6, 0x52, 0x75, 0xe8, 0x20, 0xa2, 0x52, 0xc1, 0x5d, 0x60, 0x07, 0xbd,
	0x82, 0x75, 0x62, 0x9d, 0x65, 0x3b, 0x76, 0xd0, 0x73, 0x9b, 0xe0, 0x20, 0x1d, 0x93, 0x82, 0x87,
	0x92, 0xc2, 0x2a, 0xc8, 0xc6, 0xcf, 0x3a, 0xb9, 0x53, 0xcf, 0xa3, 0x65, 0x75, 0xa4, 0xd3, 0x3a,
	0xe3, 0xde, 0x24, 0xa8, 0x06, 0x63, 0x8b, 0xa8, 0x4b, 0x00, 0x7e, 0x35, 0x93, 0x83, 0x2a, 0xc8,
	0xcc, 0x84, 0xe2, 0x99, 0x90, 0x59, 0x4e, 0x32, 0x13, 0x6a, 0x13, 0x9f, 0x26, 0xdd, 0xce, 0x42,
	0xd3, 0x7d, 0xb6, 0x12, 0xc7, 0xf9, 0xf9, 0x2b, 0x8e, 0x7f, 0xb6, 0x39, 0xc2, 0x56, 0x4a, 0xc6,
	0xd6, 0x32, 0xa7, 0x5b, 0x65, 0x0c, 0x68, 0xd1, 0xa6, 0xfe, 0x6d, 0x81, 0xbf, 0xda, 0x06, 0xde,
	0x1b, 0x3c, 0x2c, 0xaf, 0x82, 0xd7, 0x6c, 0xbe, 0x58, 0xd9, 0x16, 0x33, 0x30, 0xb7, 0xf4, 0xf4,
	0xfe, 0xf5, 0x6a, 0x1f, 0xc3, 0x43, 0x1c, 0xe7, 0xb1, 0xc9, 0xe3, 0xf9, 0xf5, 0xe2, 0x87, 0xa0,
	0xf7, 0x08, 0x47, 0xe0, 0x7f, 0x5c, 0x6a, 0x30, 0xb6, 0x11, 0x9f, 0xbe, 0x8d, 0x8d, 0xf8, 0xa5,
	0xa5, 0xba, 0x8e, 0xc6, 0x17, 0x60, 0x7e, 0x3d, 0xbe, 0x79, 0xf1, 0x36, 0x75, 0xac, 0xc9, 0xd4,
	0xb1, 0x3e, 0xa7, 0x8e, 0xf5, 0x32, 0x73, 0x32, 0x93, 0x99, 0x93, 0xf9, 0x98, 0x39, 0x99, 0xeb,
	0x92, 0x1f, 0xa8, 0xdb, 0xc8, 0x43, 0x5d, 0xde, 0x4f, 0x75, 0x47, 0xa6, 0xad, 0xc6, 0x82, 0x4a,
	0xef, 0x9f, 0xfe, 0x3b, 0xcf, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xa2, 0x55, 0x0f, 0x1e,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a post by id.
	Post(ctx context.Context, in *QueryGetPostRequest, opts ...grpc.CallOption) (*QueryGetPostResponse, error)
	// Queries a list of post items.
	PostAll(ctx context.Context, in *QueryAllPostRequest, opts ...grpc.CallOption) (*QueryAllPostResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Post(ctx context.Context, in *QueryGetPostRequest, opts ...grpc.CallOption) (*QueryGetPostResponse, error) {
	out := new(QueryGetPostResponse)
	err := c.cc.Invoke(ctx, "/user.planet.blog.Query/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PostAll(ctx context.Context, in *QueryAllPostRequest, opts ...grpc.CallOption) (*QueryAllPostResponse, error) {
	out := new(QueryAllPostResponse)
	err := c.cc.Invoke(ctx, "/user.planet.blog.Query/PostAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a post by id.
	Post(context.Context, *QueryGetPostRequest) (*QueryGetPostResponse, error)
	// Queries a list of post items.
	PostAll(context.Context, *QueryAllPostRequest) (*QueryAllPostResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Post(ctx context.Context, req *QueryGetPostRequest) (*QueryGetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (*UnimplementedQueryServer) PostAll(ctx context.Context, req *QueryAllPostRequest) (*QueryAllPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.planet.blog.Query/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Post(ctx, req.(*QueryGetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PostAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PostAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.planet.blog.Query/PostAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PostAll(ctx, req.(*QueryAllPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.planet.blog.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _Query_Post_Handler,
		},
		{
			MethodName: "PostAll",
			Handler:    _Query_PostAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog/query.proto",
}

func (m *QueryGetPostRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetPostRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetPostRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetPostResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetPostResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetPostResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Post != nil {
		{
			size, err := m.Post.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllPostRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPostRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPostRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllPostResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPostResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPostResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Post) > 0 {
		for iNdEx := len(m.Post) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Post[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetPostRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetPostResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Post != nil {
		l = m.Post.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllPostRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllPostResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Post) > 0 {
		for _, e := range m.Post {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetPostRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetPostRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetPostRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetPostResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetPostResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetPostResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Post", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Post == nil {
				m.Post = &Post{}
			}
			if err := m.Post.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllPostRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllPostRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPostRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllPostResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllPostResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPostResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Post", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Post = append(m.Post, &Post{})
			if err := m.Post[len(m.Post)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
