// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: comunicacion.proto

package comunicacion

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Llamada struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Llamada) Reset() {
	*x = Llamada{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comunicacion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Llamada) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Llamada) ProtoMessage() {}

func (x *Llamada) ProtoReflect() protoreflect.Message {
	mi := &file_comunicacion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Llamada.ProtoReflect.Descriptor instead.
func (*Llamada) Descriptor() ([]byte, []int) {
	return file_comunicacion_proto_rawDescGZIP(), []int{0}
}

func (x *Llamada) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_comunicacion_proto protoreflect.FileDescriptor

var file_comunicacion_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x63, 0x69,
	0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x07, 0x6c, 0x6c, 0x61, 0x6d, 0x61, 0x64, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x32, 0x45, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x6e, 0x64, 0x6f,
	0x12, 0x36, 0x0a, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x75, 0x6e,
	0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x2e, 0x6c, 0x6c, 0x61, 0x6d, 0x61, 0x64, 0x61, 0x1a,
	0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x2e, 0x6c,
	0x6c, 0x61, 0x6d, 0x61, 0x64, 0x61, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comunicacion_proto_rawDescOnce sync.Once
	file_comunicacion_proto_rawDescData = file_comunicacion_proto_rawDesc
)

func file_comunicacion_proto_rawDescGZIP() []byte {
	file_comunicacion_proto_rawDescOnce.Do(func() {
		file_comunicacion_proto_rawDescData = protoimpl.X.CompressGZIP(file_comunicacion_proto_rawDescData)
	})
	return file_comunicacion_proto_rawDescData
}

var file_comunicacion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_comunicacion_proto_goTypes = []interface{}{
	(*Llamada)(nil), // 0: comunicacion.llamada
}
var file_comunicacion_proto_depIdxs = []int32{
	0, // 0: comunicacion.Comunicando.call:input_type -> comunicacion.llamada
	0, // 1: comunicacion.Comunicando.call:output_type -> comunicacion.llamada
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comunicacion_proto_init() }
func file_comunicacion_proto_init() {
	if File_comunicacion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comunicacion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Llamada); i {
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
			RawDescriptor: file_comunicacion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comunicacion_proto_goTypes,
		DependencyIndexes: file_comunicacion_proto_depIdxs,
		MessageInfos:      file_comunicacion_proto_msgTypes,
	}.Build()
	File_comunicacion_proto = out.File
	file_comunicacion_proto_rawDesc = nil
	file_comunicacion_proto_goTypes = nil
	file_comunicacion_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ComunicandoClient is the client API for Comunicando service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComunicandoClient interface {
	Call(ctx context.Context, in *Llamada, opts ...grpc.CallOption) (*Llamada, error)
}

type comunicandoClient struct {
	cc grpc.ClientConnInterface
}

func NewComunicandoClient(cc grpc.ClientConnInterface) ComunicandoClient {
	return &comunicandoClient{cc}
}

func (c *comunicandoClient) Call(ctx context.Context, in *Llamada, opts ...grpc.CallOption) (*Llamada, error) {
	out := new(Llamada)
	err := c.cc.Invoke(ctx, "/comunicacion.Comunicando/call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComunicandoServer is the server API for Comunicando service.
type ComunicandoServer interface {
	Call(context.Context, *Llamada) (*Llamada, error)
}

// UnimplementedComunicandoServer can be embedded to have forward compatible implementations.
type UnimplementedComunicandoServer struct {
}

func (*UnimplementedComunicandoServer) Call(context.Context, *Llamada) (*Llamada, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterComunicandoServer(s *grpc.Server, srv ComunicandoServer) {
	s.RegisterService(&_Comunicando_serviceDesc, srv)
}

func _Comunicando_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Llamada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunicandoServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunicacion.Comunicando/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunicandoServer).Call(ctx, req.(*Llamada))
	}
	return interceptor(ctx, in, info, handler)
}

var _Comunicando_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comunicacion.Comunicando",
	HandlerType: (*ComunicandoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "call",
			Handler:    _Comunicando_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comunicacion.proto",
}
