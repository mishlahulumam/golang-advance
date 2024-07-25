package shortlink

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion8

const (
	ShortlinkService_CreateShortlink_FullMethodName = "/proto.shortlink_service.v1.ShortlinkService/CreateShortlink"
	ShortlinkService_GetLongUrl_FullMethodName      = "/proto.shortlink_service.v1.ShortlinkService/GetLongUrl"
)

type ShortlinkServiceClient interface {
	CreateShortlink(ctx context.Context, in *CreateShortlinkRequest, opts ...grpc.CallOption) (*MutationResponse, error)
	GetLongUrl(ctx context.Context, in *GetLongUrlRequest, opts ...grpc.CallOption) (*GetLongUrlResponse, error)
}

type shortlinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShortlinkServiceClient(cc grpc.ClientConnInterface) ShortlinkServiceClient {
	return &shortlinkServiceClient{cc}
}

func (c *shortlinkServiceClient) CreateShortlink(ctx context.Context, in *CreateShortlinkRequest, opts ...grpc.CallOption) (*MutationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, ShortlinkService_CreateShortlink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortlinkServiceClient) GetLongUrl(ctx context.Context, in *GetLongUrlRequest, opts ...grpc.CallOption) (*GetLongUrlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLongUrlResponse)
	err := c.cc.Invoke(ctx, ShortlinkService_GetLongUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type ShortlinkServiceServer interface {
	CreateShortlink(context.Context, *CreateShortlinkRequest) (*MutationResponse, error)
	GetLongUrl(context.Context, *GetLongUrlRequest) (*GetLongUrlResponse, error)
	mustEmbedUnimplementedShortlinkServiceServer()
}

type UnimplementedShortlinkServiceServer struct {
}

func (UnimplementedShortlinkServiceServer) CreateShortlink(context.Context, *CreateShortlinkRequest) (*MutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShortlink not implemented")
}
func (UnimplementedShortlinkServiceServer) GetLongUrl(context.Context, *GetLongUrlRequest) (*GetLongUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLongUrl not implemented")
}
func (UnimplementedShortlinkServiceServer) mustEmbedUnimplementedShortlinkServiceServer() {}

type UnsafeShortlinkServiceServer interface {
	mustEmbedUnimplementedShortlinkServiceServer()
}

func RegisterShortlinkServiceServer(s grpc.ServiceRegistrar, srv ShortlinkServiceServer) {
	s.RegisterService(&ShortlinkService_ServiceDesc, srv)
}

func _ShortlinkService_CreateShortlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShortlinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortlinkServiceServer).CreateShortlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortlinkService_CreateShortlink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortlinkServiceServer).CreateShortlink(ctx, req.(*CreateShortlinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortlinkService_GetLongUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLongUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortlinkServiceServer).GetLongUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortlinkService_GetLongUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortlinkServiceServer).GetLongUrl(ctx, req.(*GetLongUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var ShortlinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.shortlink_service.v1.ShortlinkService",
	HandlerType: (*ShortlinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShortlink",
			Handler:    _ShortlinkService_CreateShortlink_Handler,
		},
		{
			MethodName: "GetLongUrl",
			Handler:    _ShortlinkService_GetLongUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/shortlink/v1/shortlink.proto",
}
