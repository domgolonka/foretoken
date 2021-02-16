// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: email.proto

package proto

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetEmailListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emails []string `protobuf:"bytes,1,rep,name=emails,proto3" json:"emails,omitempty"`
}

func (x *GetEmailListResponse) Reset() {
	*x = GetEmailListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailListResponse) ProtoMessage() {}

func (x *GetEmailListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailListResponse.ProtoReflect.Descriptor instead.
func (*GetEmailListResponse) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{0}
}

func (x *GetEmailListResponse) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

var File_email_proto protoreflect.FileDescriptor

var file_email_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x32, 0xd7, 0x01, 0x0a, 0x0c, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x70, 0x61, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_email_proto_rawDescOnce sync.Once
	file_email_proto_rawDescData = file_email_proto_rawDesc
)

func file_email_proto_rawDescGZIP() []byte {
	file_email_proto_rawDescOnce.Do(func() {
		file_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_proto_rawDescData)
	})
	return file_email_proto_rawDescData
}

var file_email_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_email_proto_goTypes = []interface{}{
	(*GetEmailListResponse)(nil), // 0: GetEmailListResponse
	(*empty.Empty)(nil),          // 1: google.protobuf.Empty
}
var file_email_proto_depIdxs = []int32{
	1, // 0: EmailService.GetDisposableList:input_type -> google.protobuf.Empty
	1, // 1: EmailService.GetGenericList:input_type -> google.protobuf.Empty
	1, // 2: EmailService.GetSpamList:input_type -> google.protobuf.Empty
	0, // 3: EmailService.GetDisposableList:output_type -> GetEmailListResponse
	0, // 4: EmailService.GetGenericList:output_type -> GetEmailListResponse
	0, // 5: EmailService.GetSpamList:output_type -> GetEmailListResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_email_proto_init() }
func file_email_proto_init() {
	if File_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_email_proto_goTypes,
		DependencyIndexes: file_email_proto_depIdxs,
		MessageInfos:      file_email_proto_msgTypes,
	}.Build()
	File_email_proto = out.File
	file_email_proto_rawDesc = nil
	file_email_proto_goTypes = nil
	file_email_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailServiceClient interface {
	GetDisposableList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEmailListResponse, error)
	GetGenericList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEmailListResponse, error)
	GetSpamList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEmailListResponse, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) GetDisposableList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEmailListResponse, error) {
	out := new(GetEmailListResponse)
	err := c.cc.Invoke(ctx, "/EmailService/GetDisposableList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) GetGenericList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEmailListResponse, error) {
	out := new(GetEmailListResponse)
	err := c.cc.Invoke(ctx, "/EmailService/GetGenericList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) GetSpamList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEmailListResponse, error) {
	out := new(GetEmailListResponse)
	err := c.cc.Invoke(ctx, "/EmailService/GetSpamList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
type EmailServiceServer interface {
	GetDisposableList(context.Context, *empty.Empty) (*GetEmailListResponse, error)
	GetGenericList(context.Context, *empty.Empty) (*GetEmailListResponse, error)
	GetSpamList(context.Context, *empty.Empty) (*GetEmailListResponse, error)
}

// UnimplementedEmailServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (*UnimplementedEmailServiceServer) GetDisposableList(context.Context, *empty.Empty) (*GetEmailListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDisposableList not implemented")
}
func (*UnimplementedEmailServiceServer) GetGenericList(context.Context, *empty.Empty) (*GetEmailListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenericList not implemented")
}
func (*UnimplementedEmailServiceServer) GetSpamList(context.Context, *empty.Empty) (*GetEmailListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpamList not implemented")
}

func RegisterEmailServiceServer(s *grpc.Server, srv EmailServiceServer) {
	s.RegisterService(&_EmailService_serviceDesc, srv)
}

func _EmailService_GetDisposableList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetDisposableList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmailService/GetDisposableList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetDisposableList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_GetGenericList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetGenericList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmailService/GetGenericList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetGenericList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_GetSpamList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetSpamList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmailService/GetSpamList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetSpamList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDisposableList",
			Handler:    _EmailService_GetDisposableList_Handler,
		},
		{
			MethodName: "GetGenericList",
			Handler:    _EmailService_GetGenericList_Handler,
		},
		{
			MethodName: "GetSpamList",
			Handler:    _EmailService_GetSpamList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}
