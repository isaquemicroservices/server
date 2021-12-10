// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: protos/product/product.proto

package product

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

// Product model of one product
type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

// ListProducts model of one list of products
type ListProducts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ListProducts) Reset() {
	*x = ListProducts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProducts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProducts) ProtoMessage() {}

func (x *ListProducts) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProducts.ProtoReflect.Descriptor instead.
func (*ListProducts) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *ListProducts) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

// Params data model from params
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{2}
}

func (x *Params) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Params) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Void model null
type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_protos_product_product_proto_rawDescGZIP(), []int{3}
}

var File_protos_product_product_proto protoreflect.FileDescriptor

var file_protos_product_product_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x34, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x2c, 0x0a, 0x06, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69,
	0x64, 0x32, 0x61, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1e, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0d, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x00, 0x12, 0x1b, 0x0a,
	0x04, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x07, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x08,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x18, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x05, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_product_product_proto_rawDescOnce sync.Once
	file_protos_product_product_proto_rawDescData = file_protos_product_product_proto_rawDesc
)

func file_protos_product_product_proto_rawDescGZIP() []byte {
	file_protos_product_product_proto_rawDescOnce.Do(func() {
		file_protos_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_product_product_proto_rawDescData)
	})
	return file_protos_product_product_proto_rawDescData
}

var file_protos_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_product_product_proto_goTypes = []interface{}{
	(*Product)(nil),      // 0: Product
	(*ListProducts)(nil), // 1: ListProducts
	(*Params)(nil),       // 2: Params
	(*Void)(nil),         // 3: Void
}
var file_protos_product_product_proto_depIdxs = []int32{
	0, // 0: ListProducts.products:type_name -> Product
	3, // 1: Products.List:input_type -> Void
	2, // 2: Products.Show:input_type -> Params
	0, // 3: Products.Add:input_type -> Product
	1, // 4: Products.List:output_type -> ListProducts
	0, // 5: Products.Show:output_type -> Product
	3, // 6: Products.Add:output_type -> Void
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_product_product_proto_init() }
func file_protos_product_product_proto_init() {
	if File_protos_product_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_product_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_protos_product_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProducts); i {
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
		file_protos_product_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
		file_protos_product_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_protos_product_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_product_product_proto_goTypes,
		DependencyIndexes: file_protos_product_product_proto_depIdxs,
		MessageInfos:      file_protos_product_product_proto_msgTypes,
	}.Build()
	File_protos_product_product_proto = out.File
	file_protos_product_product_proto_rawDesc = nil
	file_protos_product_product_proto_goTypes = nil
	file_protos_product_product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductsClient is the client API for Products service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductsClient interface {
	// List function to show one list of products
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ListProducts, error)
	// Show model to show details of one product
	Show(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Product, error)
	// Add function to create product
	Add(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Void, error)
}

type productsClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsClient(cc grpc.ClientConnInterface) ProductsClient {
	return &productsClient{cc}
}

func (c *productsClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ListProducts, error) {
	out := new(ListProducts)
	err := c.cc.Invoke(ctx, "/Products/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) Show(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/Products/Show", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) Add(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Products/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServer is the server API for Products service.
type ProductsServer interface {
	// List function to show one list of products
	List(context.Context, *Void) (*ListProducts, error)
	// Show model to show details of one product
	Show(context.Context, *Params) (*Product, error)
	// Add function to create product
	Add(context.Context, *Product) (*Void, error)
}

// UnimplementedProductsServer can be embedded to have forward compatible implementations.
type UnimplementedProductsServer struct {
}

func (*UnimplementedProductsServer) List(context.Context, *Void) (*ListProducts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedProductsServer) Show(context.Context, *Params) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (*UnimplementedProductsServer) Add(context.Context, *Product) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterProductsServer(s *grpc.Server, srv ProductsServer) {
	s.RegisterService(&_Products_serviceDesc, srv)
}

func _Products_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Products/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Products/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).Show(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Products/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).Add(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

var _Products_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Products",
	HandlerType: (*ProductsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Products_List_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _Products_Show_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Products_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/product/product.proto",
}
