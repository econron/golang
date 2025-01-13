//
//  QueryService用のParam型とResult型を定義したprotoファイル
//

// ライセンスヘッダ:バージョン3を利用

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.29.2
// source: proto/query.proto

// パッケージの宣言

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CategoryQuery_List_FullMethodName = "/proto.CategoryQuery/List"
	CategoryQuery_ById_FullMethodName = "/proto.CategoryQuery/ById"
)

// CategoryQueryClient is the client API for CategoryQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// *****************************************
// 商品カテゴリと商品の問合せサービス型の定義
// *****************************************
//
//	商品カテゴリ問合せサービス型
type CategoryQueryClient interface {
	// すべてのカテゴリを取得して返す
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoriesResult, error)
	// 指定されたIDのカテゴリを取得して返す
	ById(ctx context.Context, in *CategoryParam, opts ...grpc.CallOption) (*CategoryResult, error)
}

type categoryQueryClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryQueryClient(cc grpc.ClientConnInterface) CategoryQueryClient {
	return &categoryQueryClient{cc}
}

func (c *categoryQueryClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoriesResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoriesResult)
	err := c.cc.Invoke(ctx, CategoryQuery_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryQueryClient) ById(ctx context.Context, in *CategoryParam, opts ...grpc.CallOption) (*CategoryResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryResult)
	err := c.cc.Invoke(ctx, CategoryQuery_ById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryQueryServer is the server API for CategoryQuery service.
// All implementations must embed UnimplementedCategoryQueryServer
// for forward compatibility
//
// *****************************************
// 商品カテゴリと商品の問合せサービス型の定義
// *****************************************
//
//	商品カテゴリ問合せサービス型
type CategoryQueryServer interface {
	// すべてのカテゴリを取得して返す
	List(context.Context, *emptypb.Empty) (*CategoriesResult, error)
	// 指定されたIDのカテゴリを取得して返す
	ById(context.Context, *CategoryParam) (*CategoryResult, error)
	mustEmbedUnimplementedCategoryQueryServer()
}

// UnimplementedCategoryQueryServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryQueryServer struct {
}

func (UnimplementedCategoryQueryServer) List(context.Context, *emptypb.Empty) (*CategoriesResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCategoryQueryServer) ById(context.Context, *CategoryParam) (*CategoryResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ById not implemented")
}
func (UnimplementedCategoryQueryServer) mustEmbedUnimplementedCategoryQueryServer() {}

// UnsafeCategoryQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryQueryServer will
// result in compilation errors.
type UnsafeCategoryQueryServer interface {
	mustEmbedUnimplementedCategoryQueryServer()
}

func RegisterCategoryQueryServer(s grpc.ServiceRegistrar, srv CategoryQueryServer) {
	s.RegisterService(&CategoryQuery_ServiceDesc, srv)
}

func _CategoryQuery_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryQueryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryQuery_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryQueryServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryQuery_ById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryQueryServer).ById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryQuery_ById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryQueryServer).ById(ctx, req.(*CategoryParam))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryQuery_ServiceDesc is the grpc.ServiceDesc for CategoryQuery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryQuery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CategoryQuery",
	HandlerType: (*CategoryQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CategoryQuery_List_Handler,
		},
		{
			MethodName: "ById",
			Handler:    _CategoryQuery_ById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/query.proto",
}

const (
	ProductQuery_ListStream_FullMethodName = "/proto.ProductQuery/ListStream"
	ProductQuery_List_FullMethodName       = "/proto.ProductQuery/List"
	ProductQuery_ById_FullMethodName       = "/proto.ProductQuery/ById"
	ProductQuery_ByKeyword_FullMethodName  = "/proto.ProductQuery/ByKeyword"
)

// ProductQueryClient is the client API for ProductQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
//	商品問合せサービス型
type ProductQueryClient interface {
	// すべての商品を取得して返す(Server streaming RPC)
	ListStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ProductQuery_ListStreamClient, error)
	// すべての商品を取得して返す
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ProductsResult, error)
	// 指定されたIDの商品を取得して返す
	ById(ctx context.Context, in *ProductParam, opts ...grpc.CallOption) (*ProductResult, error)
	// 指定されたキーワードの商品を取得して返す
	ByKeyword(ctx context.Context, in *ProductParam, opts ...grpc.CallOption) (*ProductsResult, error)
}

type productQueryClient struct {
	cc grpc.ClientConnInterface
}

func NewProductQueryClient(cc grpc.ClientConnInterface) ProductQueryClient {
	return &productQueryClient{cc}
}

func (c *productQueryClient) ListStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ProductQuery_ListStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProductQuery_ServiceDesc.Streams[0], ProductQuery_ListStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &productQueryListStreamClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductQuery_ListStreamClient interface {
	Recv() (*Product, error)
	grpc.ClientStream
}

type productQueryListStreamClient struct {
	grpc.ClientStream
}

func (x *productQueryListStreamClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productQueryClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ProductsResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductsResult)
	err := c.cc.Invoke(ctx, ProductQuery_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productQueryClient) ById(ctx context.Context, in *ProductParam, opts ...grpc.CallOption) (*ProductResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductResult)
	err := c.cc.Invoke(ctx, ProductQuery_ById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productQueryClient) ByKeyword(ctx context.Context, in *ProductParam, opts ...grpc.CallOption) (*ProductsResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductsResult)
	err := c.cc.Invoke(ctx, ProductQuery_ByKeyword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductQueryServer is the server API for ProductQuery service.
// All implementations must embed UnimplementedProductQueryServer
// for forward compatibility
//
//	商品問合せサービス型
type ProductQueryServer interface {
	// すべての商品を取得して返す(Server streaming RPC)
	ListStream(*emptypb.Empty, ProductQuery_ListStreamServer) error
	// すべての商品を取得して返す
	List(context.Context, *emptypb.Empty) (*ProductsResult, error)
	// 指定されたIDの商品を取得して返す
	ById(context.Context, *ProductParam) (*ProductResult, error)
	// 指定されたキーワードの商品を取得して返す
	ByKeyword(context.Context, *ProductParam) (*ProductsResult, error)
	mustEmbedUnimplementedProductQueryServer()
}

// UnimplementedProductQueryServer must be embedded to have forward compatible implementations.
type UnimplementedProductQueryServer struct {
}

func (UnimplementedProductQueryServer) ListStream(*emptypb.Empty, ProductQuery_ListStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListStream not implemented")
}
func (UnimplementedProductQueryServer) List(context.Context, *emptypb.Empty) (*ProductsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProductQueryServer) ById(context.Context, *ProductParam) (*ProductResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ById not implemented")
}
func (UnimplementedProductQueryServer) ByKeyword(context.Context, *ProductParam) (*ProductsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByKeyword not implemented")
}
func (UnimplementedProductQueryServer) mustEmbedUnimplementedProductQueryServer() {}

// UnsafeProductQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductQueryServer will
// result in compilation errors.
type UnsafeProductQueryServer interface {
	mustEmbedUnimplementedProductQueryServer()
}

func RegisterProductQueryServer(s grpc.ServiceRegistrar, srv ProductQueryServer) {
	s.RegisterService(&ProductQuery_ServiceDesc, srv)
}

func _ProductQuery_ListStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductQueryServer).ListStream(m, &productQueryListStreamServer{ServerStream: stream})
}

type ProductQuery_ListStreamServer interface {
	Send(*Product) error
	grpc.ServerStream
}

type productQueryListStreamServer struct {
	grpc.ServerStream
}

func (x *productQueryListStreamServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductQuery_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductQueryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductQuery_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductQueryServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductQuery_ById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductQueryServer).ById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductQuery_ById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductQueryServer).ById(ctx, req.(*ProductParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductQuery_ByKeyword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductQueryServer).ByKeyword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductQuery_ByKeyword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductQueryServer).ByKeyword(ctx, req.(*ProductParam))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductQuery_ServiceDesc is the grpc.ServiceDesc for ProductQuery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductQuery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductQuery",
	HandlerType: (*ProductQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ProductQuery_List_Handler,
		},
		{
			MethodName: "ById",
			Handler:    _ProductQuery_ById_Handler,
		},
		{
			MethodName: "ByKeyword",
			Handler:    _ProductQuery_ByKeyword_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListStream",
			Handler:       _ProductQuery_ListStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/query.proto",
}
