package shorturl

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion8

const (
	UrlService_GetLong_FullMethodName  = "/proto.user_service.v1.UrlService/GetLong"
	UrlService_ShortUrl_FullMethodName = "/proto.user_service.v1.UrlService/ShortUrl"
	UrlService_Redirect_FullMethodName = "/proto.user_service.v1.UrlService/Redirect"
)

type UrlServiceClient interface {
	GetLong(ctx context.Context, in *GetLongReq, opts ...grpc.CallOption) (*GetLongRes, error)
	ShortUrl(ctx context.Context, in *ShortUrlReq, opts ...grpc.CallOption) (*ShortUrlRes, error)
	Redirect(ctx context.Context, in *RedirectReq, opts ...grpc.CallOption) (*RedirectRes, error)
}

type urlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlServiceClient(cc grpc.ClientConnInterface) UrlServiceClient {
	return &urlServiceClient{cc}
}

func (c *urlServiceClient) GetLong(ctx context.Context, in *GetLongReq, opts ...grpc.CallOption) (*GetLongRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLongRes)
	err := c.cc.Invoke(ctx, UrlService_GetLong_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) ShortUrl(ctx context.Context, in *ShortUrlReq, opts ...grpc.CallOption) (*ShortUrlRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShortUrlRes)
	err := c.cc.Invoke(ctx, UrlService_ShortUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlServiceClient) Redirect(ctx context.Context, in *RedirectReq, opts ...grpc.CallOption) (*RedirectRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RedirectRes)
	err := c.cc.Invoke(ctx, UrlService_Redirect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type UrlServiceServer interface {
	GetLong(context.Context, *GetLongReq) (*GetLongRes, error)
	ShortUrl(context.Context, *ShortUrlReq) (*ShortUrlRes, error)
	Redirect(context.Context, *RedirectReq) (*RedirectRes, error)
	mustEmbedUnimplementedUrlServiceServer()
}

type UnimplementedUrlServiceServer struct {
}

func (UnimplementedUrlServiceServer) GetLong(context.Context, *GetLongReq) (*GetLongRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLong not implemented")
}
func (UnimplementedUrlServiceServer) ShortUrl(context.Context, *ShortUrlReq) (*ShortUrlRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortUrl not implemented")
}
func (UnimplementedUrlServiceServer) Redirect(context.Context, *RedirectReq) (*RedirectRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redirect not implemented")
}
func (UnimplementedUrlServiceServer) mustEmbedUnimplementedUrlServiceServer() {}

type UnsafeUrlServiceServer interface {
	mustEmbedUnimplementedUrlServiceServer()
}

func RegisterUrlServiceServer(s grpc.ServiceRegistrar, srv UrlServiceServer) {
	s.RegisterService(&UrlService_ServiceDesc, srv)
}

func _UrlService_GetLong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLongReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).GetLong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlService_GetLong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).GetLong(ctx, req.(*GetLongReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_ShortUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortUrlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).ShortUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlService_ShortUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).ShortUrl(ctx, req.(*ShortUrlReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlService_Redirect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedirectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlServiceServer).Redirect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlService_Redirect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlServiceServer).Redirect(ctx, req.(*RedirectReq))
	}
	return interceptor(ctx, in, info, handler)
}

var UrlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.user_service.v1.UrlService",
	HandlerType: (*UrlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLong",
			Handler:    _UrlService_GetLong_Handler,
		},
		{
			MethodName: "ShortUrl",
			Handler:    _UrlService_ShortUrl_Handler,
		},
		{
			MethodName: "Redirect",
			Handler:    _UrlService_Redirect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/shorturl_service/v1/shorturl.proto",
}
