// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/compute/v1/disk_type_service.proto

package compute

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetDiskTypeRequest struct {
	// ID of the disk type to return information about.
	// To get the disk type ID use a [DiskTypeService.List] request.
	DiskTypeId           string   `protobuf:"bytes,1,opt,name=disk_type_id,json=diskTypeId,proto3" json:"disk_type_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDiskTypeRequest) Reset()         { *m = GetDiskTypeRequest{} }
func (m *GetDiskTypeRequest) String() string { return proto.CompactTextString(m) }
func (*GetDiskTypeRequest) ProtoMessage()    {}
func (*GetDiskTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0433b2f525f69d83, []int{0}
}

func (m *GetDiskTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiskTypeRequest.Unmarshal(m, b)
}
func (m *GetDiskTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiskTypeRequest.Marshal(b, m, deterministic)
}
func (m *GetDiskTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiskTypeRequest.Merge(m, src)
}
func (m *GetDiskTypeRequest) XXX_Size() int {
	return xxx_messageInfo_GetDiskTypeRequest.Size(m)
}
func (m *GetDiskTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiskTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiskTypeRequest proto.InternalMessageInfo

func (m *GetDiskTypeRequest) GetDiskTypeId() string {
	if m != nil {
		return m.DiskTypeId
	}
	return ""
}

type ListDiskTypesRequest struct {
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListDiskTypesResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [ListDiskTypesResponse.next_page_token] returned by a previous list request.
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDiskTypesRequest) Reset()         { *m = ListDiskTypesRequest{} }
func (m *ListDiskTypesRequest) String() string { return proto.CompactTextString(m) }
func (*ListDiskTypesRequest) ProtoMessage()    {}
func (*ListDiskTypesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0433b2f525f69d83, []int{1}
}

func (m *ListDiskTypesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDiskTypesRequest.Unmarshal(m, b)
}
func (m *ListDiskTypesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDiskTypesRequest.Marshal(b, m, deterministic)
}
func (m *ListDiskTypesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDiskTypesRequest.Merge(m, src)
}
func (m *ListDiskTypesRequest) XXX_Size() int {
	return xxx_messageInfo_ListDiskTypesRequest.Size(m)
}
func (m *ListDiskTypesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDiskTypesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDiskTypesRequest proto.InternalMessageInfo

func (m *ListDiskTypesRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListDiskTypesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListDiskTypesResponse struct {
	// List of disk types.
	DiskTypes []*DiskType `protobuf:"bytes,1,rep,name=disk_types,json=diskTypes,proto3" json:"disk_types,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListDiskTypesRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListDiskTypesRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDiskTypesResponse) Reset()         { *m = ListDiskTypesResponse{} }
func (m *ListDiskTypesResponse) String() string { return proto.CompactTextString(m) }
func (*ListDiskTypesResponse) ProtoMessage()    {}
func (*ListDiskTypesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0433b2f525f69d83, []int{2}
}

func (m *ListDiskTypesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDiskTypesResponse.Unmarshal(m, b)
}
func (m *ListDiskTypesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDiskTypesResponse.Marshal(b, m, deterministic)
}
func (m *ListDiskTypesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDiskTypesResponse.Merge(m, src)
}
func (m *ListDiskTypesResponse) XXX_Size() int {
	return xxx_messageInfo_ListDiskTypesResponse.Size(m)
}
func (m *ListDiskTypesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDiskTypesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDiskTypesResponse proto.InternalMessageInfo

func (m *ListDiskTypesResponse) GetDiskTypes() []*DiskType {
	if m != nil {
		return m.DiskTypes
	}
	return nil
}

func (m *ListDiskTypesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDiskTypeRequest)(nil), "yandex.cloud.compute.v1.GetDiskTypeRequest")
	proto.RegisterType((*ListDiskTypesRequest)(nil), "yandex.cloud.compute.v1.ListDiskTypesRequest")
	proto.RegisterType((*ListDiskTypesResponse)(nil), "yandex.cloud.compute.v1.ListDiskTypesResponse")
}

func init() {
	proto.RegisterFile("yandex/cloud/compute/v1/disk_type_service.proto", fileDescriptor_0433b2f525f69d83)
}

var fileDescriptor_0433b2f525f69d83 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0x26, 0xed, 0xba, 0x98, 0xa7, 0xb2, 0x30, 0xb8, 0x6c, 0x89, 0x2e, 0xac, 0x41, 0xba, 0x05,
	0x6d, 0xa6, 0x59, 0x8f, 0xae, 0x20, 0x55, 0x58, 0x04, 0x0f, 0x92, 0xdd, 0x93, 0x97, 0x90, 0x36,
	0x8f, 0x38, 0xb4, 0xce, 0x8c, 0x3b, 0x93, 0xd0, 0x56, 0x3c, 0xf8, 0xe3, 0xe4, 0xd5, 0xbb, 0xff,
	0x4e, 0xbd, 0xfb, 0x2f, 0x78, 0xf0, 0x6f, 0xf0, 0x24, 0x99, 0x24, 0xb5, 0xad, 0x0d, 0xdd, 0xdb,
	0x30, 0xef, 0xfb, 0xbe, 0xf7, 0xcd, 0xfb, 0xe6, 0x01, 0x9d, 0x46, 0x3c, 0xc6, 0x09, 0x1d, 0x8e,
	0x45, 0x1a, 0xd3, 0xa1, 0x78, 0x2b, 0x53, 0x8d, 0x34, 0xf3, 0x69, 0xcc, 0xd4, 0x28, 0xd4, 0x53,
	0x89, 0xa1, 0xc2, 0xcb, 0x8c, 0x0d, 0xd1, 0x93, 0x97, 0x42, 0x0b, 0x72, 0x50, 0x10, 0x3c, 0x43,
	0xf0, 0x4a, 0x82, 0x97, 0xf9, 0xce, 0xdd, 0x44, 0x88, 0x64, 0x8c, 0x34, 0x92, 0x8c, 0x46, 0x9c,
	0x0b, 0x1d, 0x69, 0x26, 0xb8, 0x2a, 0x68, 0xce, 0xf1, 0xd6, 0x3e, 0x25, 0xf0, 0x70, 0x05, 0x98,
	0x45, 0x63, 0x16, 0x1b, 0xa1, 0xa2, 0xec, 0x9e, 0x02, 0x39, 0x43, 0xfd, 0x9c, 0xa9, 0xd1, 0xc5,
	0x54, 0x62, 0x80, 0xef, 0x52, 0x54, 0x9a, 0xb4, 0xe1, 0xe6, 0x3f, 0xbf, 0x2c, 0x6e, 0x59, 0x47,
	0x56, 0xc7, 0xee, 0xef, 0xfc, 0x9e, 0xfb, 0x56, 0x00, 0x71, 0x09, 0x7e, 0x11, 0xbb, 0x0c, 0x6e,
	0xbf, 0x64, 0x6a, 0x41, 0x57, 0x15, 0xff, 0x18, 0x6c, 0x19, 0x25, 0x18, 0x2a, 0x36, 0x43, 0x43,
	0x6e, 0xf6, 0xe1, 0xcf, 0xdc, 0xdf, 0x3d, 0x7d, 0xe2, 0xf7, 0x7a, 0xbd, 0xe0, 0x7a, 0x5e, 0x3c,
	0x67, 0x33, 0x24, 0x1d, 0x00, 0x03, 0xd4, 0x62, 0x84, 0xbc, 0xd5, 0x30, 0x6d, 0xec, 0xaf, 0x3f,
	0xfc, 0x6b, 0x06, 0x19, 0x18, 0x95, 0x8b, 0xbc, 0xe6, 0x7e, 0xb4, 0x60, 0x7f, 0xad, 0x97, 0x92,
	0x82, 0x2b, 0x24, 0x4f, 0x01, 0x16, 0x66, 0x55, 0xcb, 0x3a, 0x6a, 0x76, 0x6e, 0x9c, 0xdc, 0xf3,
	0x6a, 0xc6, 0xea, 0x2d, 0x9e, 0x6a, 0x57, 0xef, 0x50, 0xa4, 0x0d, 0x7b, 0x1c, 0x27, 0x3a, 0x5c,
	0xb7, 0x12, 0xdc, 0xca, 0xaf, 0x5f, 0x55, 0x1e, 0x4e, 0xbe, 0x37, 0x60, 0xaf, 0xe2, 0x9f, 0x17,
	0x29, 0x92, 0xcf, 0x16, 0x34, 0xcf, 0x50, 0x93, 0x07, 0xb5, 0x1d, 0xff, 0x9f, 0xaf, 0xb3, 0xdd,
	0x9e, 0xfb, 0xf0, 0xd3, 0xcf, 0x5f, 0xdf, 0x1a, 0x6d, 0x72, 0x7f, 0x3d, 0x5c, 0x63, 0x99, 0xbe,
	0x5f, 0xce, 0xe7, 0x03, 0xf9, 0x62, 0xc1, 0x4e, 0x3e, 0x1d, 0xd2, 0xad, 0x55, 0xde, 0x14, 0x94,
	0xe3, 0x5d, 0x15, 0x5e, 0xcc, 0xda, 0x3d, 0x34, 0xae, 0x0e, 0xc8, 0xfe, 0x46, 0x57, 0xfd, 0x01,
	0xdc, 0x59, 0xd1, 0x8b, 0x24, 0x5b, 0xd2, 0x7c, 0xfd, 0x2c, 0x61, 0xfa, 0x4d, 0x3a, 0xc8, 0xaf,
	0xca, 0x3d, 0xe9, 0x16, 0xdf, 0x32, 0x11, 0xdd, 0x04, 0xb9, 0xf9, 0x91, 0x75, 0x0b, 0xf4, 0xb8,
	0x3c, 0x0e, 0x76, 0x0d, 0xec, 0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xcb, 0xd9, 0xc1,
	0x6a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiskTypeServiceClient is the client API for DiskTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiskTypeServiceClient interface {
	// Returns the information about specified disk type.
	//
	// To get the list of available disk types, make a [List] request.
	Get(ctx context.Context, in *GetDiskTypeRequest, opts ...grpc.CallOption) (*DiskType, error)
	// Retrieves the list of disk types for the specified folder.
	List(ctx context.Context, in *ListDiskTypesRequest, opts ...grpc.CallOption) (*ListDiskTypesResponse, error)
}

type diskTypeServiceClient struct {
	cc *grpc.ClientConn
}

func NewDiskTypeServiceClient(cc *grpc.ClientConn) DiskTypeServiceClient {
	return &diskTypeServiceClient{cc}
}

func (c *diskTypeServiceClient) Get(ctx context.Context, in *GetDiskTypeRequest, opts ...grpc.CallOption) (*DiskType, error) {
	out := new(DiskType)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.DiskTypeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskTypeServiceClient) List(ctx context.Context, in *ListDiskTypesRequest, opts ...grpc.CallOption) (*ListDiskTypesResponse, error) {
	out := new(ListDiskTypesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.DiskTypeService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiskTypeServiceServer is the server API for DiskTypeService service.
type DiskTypeServiceServer interface {
	// Returns the information about specified disk type.
	//
	// To get the list of available disk types, make a [List] request.
	Get(context.Context, *GetDiskTypeRequest) (*DiskType, error)
	// Retrieves the list of disk types for the specified folder.
	List(context.Context, *ListDiskTypesRequest) (*ListDiskTypesResponse, error)
}

// UnimplementedDiskTypeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDiskTypeServiceServer struct {
}

func (*UnimplementedDiskTypeServiceServer) Get(ctx context.Context, req *GetDiskTypeRequest) (*DiskType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDiskTypeServiceServer) List(ctx context.Context, req *ListDiskTypesRequest) (*ListDiskTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterDiskTypeServiceServer(s *grpc.Server, srv DiskTypeServiceServer) {
	s.RegisterService(&_DiskTypeService_serviceDesc, srv)
}

func _DiskTypeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiskTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskTypeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.DiskTypeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskTypeServiceServer).Get(ctx, req.(*GetDiskTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskTypeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiskTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskTypeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.DiskTypeService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskTypeServiceServer).List(ctx, req.(*ListDiskTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiskTypeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.compute.v1.DiskTypeService",
	HandlerType: (*DiskTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DiskTypeService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DiskTypeService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/compute/v1/disk_type_service.proto",
}
