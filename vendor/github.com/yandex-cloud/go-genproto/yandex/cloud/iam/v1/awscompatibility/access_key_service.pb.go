// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/iam/v1/awscompatibility/access_key_service.proto

package awscompatibility

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type GetAccessKeyRequest struct {
	// ID of the AccessKey resource to return.
	// To get the access key ID, use a [AccessKeyService.List] request.
	AccessKeyId          string   `protobuf:"bytes,1,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccessKeyRequest) Reset()         { *m = GetAccessKeyRequest{} }
func (m *GetAccessKeyRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccessKeyRequest) ProtoMessage()    {}
func (*GetAccessKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e7adea57f560ab, []int{0}
}

func (m *GetAccessKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccessKeyRequest.Unmarshal(m, b)
}
func (m *GetAccessKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccessKeyRequest.Marshal(b, m, deterministic)
}
func (m *GetAccessKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccessKeyRequest.Merge(m, src)
}
func (m *GetAccessKeyRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccessKeyRequest.Size(m)
}
func (m *GetAccessKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccessKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccessKeyRequest proto.InternalMessageInfo

func (m *GetAccessKeyRequest) GetAccessKeyId() string {
	if m != nil {
		return m.AccessKeyId
	}
	return ""
}

type ListAccessKeysRequest struct {
	// ID of the service account to list access keys for.
	// To get the service account ID, use a [yandex.cloud.iam.v1.ServiceAccountService.List] request.
	// If not specified, it defaults to the subject that made the request.
	ServiceAccountId string `protobuf:"bytes,1,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListAccessKeysResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token]
	// to the [ListAccessKeysResponse.next_page_token]
	// returned by a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccessKeysRequest) Reset()         { *m = ListAccessKeysRequest{} }
func (m *ListAccessKeysRequest) String() string { return proto.CompactTextString(m) }
func (*ListAccessKeysRequest) ProtoMessage()    {}
func (*ListAccessKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e7adea57f560ab, []int{1}
}

func (m *ListAccessKeysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccessKeysRequest.Unmarshal(m, b)
}
func (m *ListAccessKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccessKeysRequest.Marshal(b, m, deterministic)
}
func (m *ListAccessKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccessKeysRequest.Merge(m, src)
}
func (m *ListAccessKeysRequest) XXX_Size() int {
	return xxx_messageInfo_ListAccessKeysRequest.Size(m)
}
func (m *ListAccessKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccessKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccessKeysRequest proto.InternalMessageInfo

func (m *ListAccessKeysRequest) GetServiceAccountId() string {
	if m != nil {
		return m.ServiceAccountId
	}
	return ""
}

func (m *ListAccessKeysRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListAccessKeysRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListAccessKeysResponse struct {
	// List of access keys.
	AccessKeys []*AccessKey `protobuf:"bytes,1,rep,name=access_keys,json=accessKeys,proto3" json:"access_keys,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListAccessKeysRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListAccessKeysRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccessKeysResponse) Reset()         { *m = ListAccessKeysResponse{} }
func (m *ListAccessKeysResponse) String() string { return proto.CompactTextString(m) }
func (*ListAccessKeysResponse) ProtoMessage()    {}
func (*ListAccessKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e7adea57f560ab, []int{2}
}

func (m *ListAccessKeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccessKeysResponse.Unmarshal(m, b)
}
func (m *ListAccessKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccessKeysResponse.Marshal(b, m, deterministic)
}
func (m *ListAccessKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccessKeysResponse.Merge(m, src)
}
func (m *ListAccessKeysResponse) XXX_Size() int {
	return xxx_messageInfo_ListAccessKeysResponse.Size(m)
}
func (m *ListAccessKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccessKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccessKeysResponse proto.InternalMessageInfo

func (m *ListAccessKeysResponse) GetAccessKeys() []*AccessKey {
	if m != nil {
		return m.AccessKeys
	}
	return nil
}

func (m *ListAccessKeysResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type CreateAccessKeyRequest struct {
	// ID of the service account to create an access key for.
	// To get the service account ID, use a [yandex.cloud.iam.v1.ServiceAccountService.List] request.
	// If not specified, it defaults to the subject that made the request.
	ServiceAccountId string `protobuf:"bytes,1,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Description of the access key.
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccessKeyRequest) Reset()         { *m = CreateAccessKeyRequest{} }
func (m *CreateAccessKeyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccessKeyRequest) ProtoMessage()    {}
func (*CreateAccessKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e7adea57f560ab, []int{3}
}

func (m *CreateAccessKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccessKeyRequest.Unmarshal(m, b)
}
func (m *CreateAccessKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccessKeyRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccessKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccessKeyRequest.Merge(m, src)
}
func (m *CreateAccessKeyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccessKeyRequest.Size(m)
}
func (m *CreateAccessKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccessKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccessKeyRequest proto.InternalMessageInfo

func (m *CreateAccessKeyRequest) GetServiceAccountId() string {
	if m != nil {
		return m.ServiceAccountId
	}
	return ""
}

func (m *CreateAccessKeyRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateAccessKeyResponse struct {
	// AccessKey resource.
	AccessKey *AccessKey `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// Secret access key.
	// The key is AWS compatible.
	Secret               string   `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccessKeyResponse) Reset()         { *m = CreateAccessKeyResponse{} }
func (m *CreateAccessKeyResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAccessKeyResponse) ProtoMessage()    {}
func (*CreateAccessKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e7adea57f560ab, []int{4}
}

func (m *CreateAccessKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccessKeyResponse.Unmarshal(m, b)
}
func (m *CreateAccessKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccessKeyResponse.Marshal(b, m, deterministic)
}
func (m *CreateAccessKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccessKeyResponse.Merge(m, src)
}
func (m *CreateAccessKeyResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAccessKeyResponse.Size(m)
}
func (m *CreateAccessKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccessKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccessKeyResponse proto.InternalMessageInfo

func (m *CreateAccessKeyResponse) GetAccessKey() *AccessKey {
	if m != nil {
		return m.AccessKey
	}
	return nil
}

func (m *CreateAccessKeyResponse) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type DeleteAccessKeyRequest struct {
	// ID of the access key to delete.
	// To get the access key ID, use a [AccessKeyService.List] request.
	AccessKeyId          string   `protobuf:"bytes,1,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAccessKeyRequest) Reset()         { *m = DeleteAccessKeyRequest{} }
func (m *DeleteAccessKeyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAccessKeyRequest) ProtoMessage()    {}
func (*DeleteAccessKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82e7adea57f560ab, []int{5}
}

func (m *DeleteAccessKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAccessKeyRequest.Unmarshal(m, b)
}
func (m *DeleteAccessKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAccessKeyRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAccessKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAccessKeyRequest.Merge(m, src)
}
func (m *DeleteAccessKeyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAccessKeyRequest.Size(m)
}
func (m *DeleteAccessKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAccessKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAccessKeyRequest proto.InternalMessageInfo

func (m *DeleteAccessKeyRequest) GetAccessKeyId() string {
	if m != nil {
		return m.AccessKeyId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAccessKeyRequest)(nil), "yandex.cloud.iam.v1.awscompatibility.GetAccessKeyRequest")
	proto.RegisterType((*ListAccessKeysRequest)(nil), "yandex.cloud.iam.v1.awscompatibility.ListAccessKeysRequest")
	proto.RegisterType((*ListAccessKeysResponse)(nil), "yandex.cloud.iam.v1.awscompatibility.ListAccessKeysResponse")
	proto.RegisterType((*CreateAccessKeyRequest)(nil), "yandex.cloud.iam.v1.awscompatibility.CreateAccessKeyRequest")
	proto.RegisterType((*CreateAccessKeyResponse)(nil), "yandex.cloud.iam.v1.awscompatibility.CreateAccessKeyResponse")
	proto.RegisterType((*DeleteAccessKeyRequest)(nil), "yandex.cloud.iam.v1.awscompatibility.DeleteAccessKeyRequest")
}

func init() {
	proto.RegisterFile("yandex/cloud/iam/v1/awscompatibility/access_key_service.proto", fileDescriptor_82e7adea57f560ab)
}

var fileDescriptor_82e7adea57f560ab = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xc1, 0x4f, 0x13, 0x4f,
	0x14, 0xc7, 0x33, 0x14, 0x1a, 0xfa, 0xfa, 0x23, 0x3f, 0x32, 0xc6, 0xda, 0x54, 0x4d, 0xc8, 0x86,
	0x60, 0x45, 0xd9, 0x69, 0x51, 0x48, 0x94, 0xf6, 0x00, 0x6a, 0x08, 0x6a, 0x0c, 0x29, 0x5e, 0xf4,
	0xd2, 0x4c, 0x77, 0x9f, 0x75, 0x42, 0xbb, 0xb3, 0x76, 0xa6, 0x85, 0x62, 0x38, 0x68, 0xe2, 0x85,
	0xa3, 0x5e, 0x3d, 0xf9, 0x0f, 0x70, 0xf3, 0x5f, 0x80, 0x93, 0x17, 0xff, 0x05, 0x0f, 0xfe, 0x0d,
	0x9e, 0xcc, 0xce, 0x6e, 0x5b, 0x0a, 0x35, 0x59, 0x7a, 0xdc, 0x79, 0xf3, 0x7d, 0xef, 0x3b, 0x9f,
	0xf7, 0x66, 0x16, 0xca, 0x5d, 0xee, 0xb9, 0xb8, 0xcf, 0x9c, 0x86, 0x6c, 0xbb, 0x4c, 0xf0, 0x26,
	0xeb, 0x14, 0x19, 0xdf, 0x53, 0x8e, 0x6c, 0xfa, 0x5c, 0x8b, 0x9a, 0x68, 0x08, 0xdd, 0x65, 0xdc,
	0x71, 0x50, 0xa9, 0xea, 0x2e, 0x76, 0xab, 0x0a, 0x5b, 0x1d, 0xe1, 0xa0, 0xed, 0xb7, 0xa4, 0x96,
	0x74, 0x3e, 0x94, 0xdb, 0x46, 0x6e, 0x0b, 0xde, 0xb4, 0x3b, 0x45, 0xfb, 0xbc, 0x3c, 0x77, 0xa3,
	0x2e, 0x65, 0xbd, 0x81, 0x8c, 0xfb, 0x82, 0x71, 0xcf, 0x93, 0x9a, 0x6b, 0x21, 0x3d, 0x15, 0xe6,
	0xc8, 0x5d, 0x8f, 0xa2, 0xe6, 0xab, 0xd6, 0x7e, 0xc3, 0xb0, 0xe9, 0xeb, 0x6e, 0x14, 0x5c, 0xb9,
	0xa4, 0xbf, 0x48, 0x76, 0x73, 0x48, 0xd6, 0xe1, 0x0d, 0xe1, 0x9a, 0x9a, 0x61, 0xd8, 0xda, 0x84,
	0x2b, 0x9b, 0xa8, 0xd7, 0x8d, 0xea, 0x19, 0x76, 0x2b, 0xf8, 0xae, 0x8d, 0x4a, 0xd3, 0x02, 0xcc,
	0x9c, 0x39, 0xa9, 0x70, 0xb3, 0x64, 0x8e, 0xe4, 0x53, 0x1b, 0xff, 0xfd, 0x3e, 0x29, 0x92, 0xa3,
	0xd3, 0xe2, 0x64, 0xa9, 0xbc, 0x52, 0xa8, 0xa4, 0x79, 0x4f, 0xb6, 0xe5, 0x5a, 0xdf, 0x08, 0x5c,
	0x7d, 0x2e, 0xd4, 0x20, 0x95, 0xea, 0xe5, 0x5a, 0x05, 0x1a, 0xa1, 0xaa, 0x72, 0xc7, 0x91, 0x6d,
	0x4f, 0x0f, 0x12, 0x4e, 0xf7, 0x93, 0xcd, 0x46, 0x7b, 0xd6, 0xc3, 0x2d, 0x5b, 0x2e, 0xbd, 0x05,
	0x29, 0x9f, 0xd7, 0xb1, 0xaa, 0xc4, 0x01, 0x66, 0x27, 0xe6, 0x48, 0x3e, 0xb1, 0x01, 0x7f, 0x4e,
	0x8a, 0xc9, 0x52, 0xb9, 0x58, 0x28, 0x14, 0x2a, 0xd3, 0x41, 0x70, 0x47, 0x1c, 0x20, 0xcd, 0x03,
	0x98, 0x8d, 0x5a, 0xee, 0xa2, 0x97, 0x4d, 0x98, 0xc4, 0xa9, 0xa3, 0xd3, 0xe2, 0x94, 0xd9, 0x59,
	0x31, 0x59, 0x5e, 0x06, 0x31, 0xeb, 0x33, 0x81, 0xcc, 0x79, 0x93, 0xca, 0x97, 0x9e, 0x42, 0xba,
	0x0d, 0xe9, 0xc1, 0x89, 0x55, 0x96, 0xcc, 0x25, 0xf2, 0xe9, 0x65, 0x66, 0xc7, 0xe9, 0xaa, 0x3d,
	0xc0, 0x07, 0x7d, 0x24, 0x8a, 0x2e, 0xc0, 0xff, 0x1e, 0xee, 0xeb, 0xea, 0x19, 0x6f, 0xc1, 0x29,
	0x52, 0x95, 0x99, 0x60, 0x79, 0xbb, 0x6f, 0xea, 0x10, 0x32, 0x8f, 0x5a, 0xc8, 0x35, 0x5e, 0xe8,
	0xc2, 0xb8, 0xe4, 0xee, 0x40, 0xda, 0x45, 0xe5, 0xb4, 0x84, 0x1f, 0x74, 0x3a, 0xac, 0xda, 0x23,
	0xb2, 0xbc, 0xb2, 0x5a, 0x39, 0x1b, 0xb5, 0x3e, 0x10, 0xb8, 0x76, 0xa1, 0x7e, 0x04, 0xe5, 0x05,
	0xc0, 0x00, 0x8a, 0x29, 0x3c, 0x06, 0x93, 0x54, 0x9f, 0x09, 0xcd, 0x40, 0x52, 0xa1, 0xd3, 0x42,
	0x1d, 0x91, 0x88, 0xbe, 0xac, 0xa7, 0x90, 0x79, 0x8c, 0x0d, 0x1c, 0x81, 0xe0, 0xd2, 0x83, 0xb8,
	0xfc, 0x63, 0x0a, 0x66, 0xfb, 0x69, 0x76, 0x42, 0x34, 0xf4, 0x98, 0xc0, 0x64, 0xd0, 0x78, 0xba,
	0x16, 0xcf, 0xfd, 0xc8, 0x49, 0xce, 0x95, 0xc6, 0x13, 0x87, 0x30, 0xad, 0xbb, 0x1f, 0x7f, 0xfe,
	0xfa, 0x32, 0xb1, 0x40, 0xe7, 0xcd, 0xe5, 0xe5, 0x7b, 0x6a, 0x69, 0xf8, 0xea, 0x06, 0xd7, 0x79,
	0x30, 0x3d, 0xc7, 0x04, 0x12, 0x9b, 0xa8, 0xe9, 0x83, 0x78, 0x35, 0x47, 0x5c, 0xe2, 0xdc, 0x65,
	0x3b, 0x65, 0x95, 0x8c, 0xc3, 0x55, 0x7a, 0x3f, 0x8e, 0x43, 0xf6, 0x7e, 0xa8, 0x31, 0x87, 0xf4,
	0x3b, 0x81, 0x64, 0x38, 0x48, 0x34, 0x26, 0xa8, 0xd1, 0x63, 0x9f, 0x2b, 0x8f, 0xa9, 0x8e, 0x38,
	0x33, 0x73, 0x8a, 0xdb, 0x56, 0x2c, 0xce, 0x0f, 0xc9, 0x22, 0xfd, 0x4a, 0x20, 0x19, 0x8e, 0x5f,
	0x5c, 0xe3, 0xa3, 0x87, 0x35, 0x97, 0xb1, 0xc3, 0x07, 0xdc, 0xee, 0x3d, 0xe0, 0xf6, 0x93, 0xe0,
	0x01, 0xef, 0x71, 0x5d, 0x1c, 0x8b, 0xeb, 0xc6, 0x27, 0x02, 0xf9, 0x21, 0x57, 0xdc, 0x17, 0xff,
	0x72, 0xf6, 0xfa, 0x55, 0x5d, 0xe8, 0xb7, 0xed, 0x9a, 0xed, 0xc8, 0x26, 0x0b, 0x45, 0x4b, 0xe1,
	0xcb, 0x5f, 0x97, 0x4b, 0x75, 0xf4, 0x8c, 0x31, 0x16, 0xe7, 0x4f, 0xb2, 0x76, 0x7e, 0xa1, 0x96,
	0x34, 0xe2, 0x7b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x39, 0x5f, 0xe5, 0x29, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccessKeyServiceClient is the client API for AccessKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessKeyServiceClient interface {
	// Retrieves the list of access keys for the specified service account.
	List(ctx context.Context, in *ListAccessKeysRequest, opts ...grpc.CallOption) (*ListAccessKeysResponse, error)
	// Returns the specified access key.
	//
	// To get the list of available access keys, make a [List] request.
	Get(ctx context.Context, in *GetAccessKeyRequest, opts ...grpc.CallOption) (*AccessKey, error)
	// Creates an access key for the specified service account.
	Create(ctx context.Context, in *CreateAccessKeyRequest, opts ...grpc.CallOption) (*CreateAccessKeyResponse, error)
	// Deletes the specified access key.
	Delete(ctx context.Context, in *DeleteAccessKeyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type accessKeyServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccessKeyServiceClient(cc *grpc.ClientConn) AccessKeyServiceClient {
	return &accessKeyServiceClient{cc}
}

func (c *accessKeyServiceClient) List(ctx context.Context, in *ListAccessKeysRequest, opts ...grpc.CallOption) (*ListAccessKeysResponse, error) {
	out := new(ListAccessKeysResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.awscompatibility.AccessKeyService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyServiceClient) Get(ctx context.Context, in *GetAccessKeyRequest, opts ...grpc.CallOption) (*AccessKey, error) {
	out := new(AccessKey)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.awscompatibility.AccessKeyService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyServiceClient) Create(ctx context.Context, in *CreateAccessKeyRequest, opts ...grpc.CallOption) (*CreateAccessKeyResponse, error) {
	out := new(CreateAccessKeyResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.awscompatibility.AccessKeyService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyServiceClient) Delete(ctx context.Context, in *DeleteAccessKeyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.awscompatibility.AccessKeyService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessKeyServiceServer is the server API for AccessKeyService service.
type AccessKeyServiceServer interface {
	// Retrieves the list of access keys for the specified service account.
	List(context.Context, *ListAccessKeysRequest) (*ListAccessKeysResponse, error)
	// Returns the specified access key.
	//
	// To get the list of available access keys, make a [List] request.
	Get(context.Context, *GetAccessKeyRequest) (*AccessKey, error)
	// Creates an access key for the specified service account.
	Create(context.Context, *CreateAccessKeyRequest) (*CreateAccessKeyResponse, error)
	// Deletes the specified access key.
	Delete(context.Context, *DeleteAccessKeyRequest) (*empty.Empty, error)
}

// UnimplementedAccessKeyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccessKeyServiceServer struct {
}

func (*UnimplementedAccessKeyServiceServer) List(ctx context.Context, req *ListAccessKeysRequest) (*ListAccessKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedAccessKeyServiceServer) Get(ctx context.Context, req *GetAccessKeyRequest) (*AccessKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAccessKeyServiceServer) Create(ctx context.Context, req *CreateAccessKeyRequest) (*CreateAccessKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAccessKeyServiceServer) Delete(ctx context.Context, req *DeleteAccessKeyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAccessKeyServiceServer(s *grpc.Server, srv AccessKeyServiceServer) {
	s.RegisterService(&_AccessKeyService_serviceDesc, srv)
}

func _AccessKeyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccessKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessKeyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.awscompatibility.AccessKeyService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessKeyServiceServer).List(ctx, req.(*ListAccessKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessKeyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessKeyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.awscompatibility.AccessKeyService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessKeyServiceServer).Get(ctx, req.(*GetAccessKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessKeyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccessKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessKeyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.awscompatibility.AccessKeyService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessKeyServiceServer).Create(ctx, req.(*CreateAccessKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessKeyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccessKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessKeyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.awscompatibility.AccessKeyService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessKeyServiceServer).Delete(ctx, req.(*DeleteAccessKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessKeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.iam.v1.awscompatibility.AccessKeyService",
	HandlerType: (*AccessKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AccessKeyService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AccessKeyService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccessKeyService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccessKeyService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/iam/v1/awscompatibility/access_key_service.proto",
}
