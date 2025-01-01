// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: travel.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TravelAgencyService_BookTravel_FullMethodName    = "/TravelAgencyService/BookTravel"
	TravelAgencyService_CancelBooking_FullMethodName = "/TravelAgencyService/CancelBooking"
	TravelAgencyService_CancelTravel_FullMethodName  = "/TravelAgencyService/CancelTravel"
	TravelAgencyService_ApproveTravel_FullMethodName = "/TravelAgencyService/ApproveTravel"
	TravelAgencyService_FinishTravel_FullMethodName  = "/TravelAgencyService/FinishTravel"
)

// TravelAgencyServiceClient is the client API for TravelAgencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TravelAgencyServiceClient interface {
	BookTravel(ctx context.Context, in *BookTravelRequest, opts ...grpc.CallOption) (*BookTravelResponse, error)
	CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*CancelBookingResponse, error)
	CancelTravel(ctx context.Context, in *CancelTravelRequest, opts ...grpc.CallOption) (*CancelTravelResponse, error)
	ApproveTravel(ctx context.Context, in *ApproveTravelRequest, opts ...grpc.CallOption) (*ApproveTravelResponse, error)
	FinishTravel(ctx context.Context, in *FinishTravelRequest, opts ...grpc.CallOption) (*FinishTravelResponse, error)
}

type travelAgencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTravelAgencyServiceClient(cc grpc.ClientConnInterface) TravelAgencyServiceClient {
	return &travelAgencyServiceClient{cc}
}

func (c *travelAgencyServiceClient) BookTravel(ctx context.Context, in *BookTravelRequest, opts ...grpc.CallOption) (*BookTravelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookTravelResponse)
	err := c.cc.Invoke(ctx, TravelAgencyService_BookTravel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelAgencyServiceClient) CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*CancelBookingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelBookingResponse)
	err := c.cc.Invoke(ctx, TravelAgencyService_CancelBooking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelAgencyServiceClient) CancelTravel(ctx context.Context, in *CancelTravelRequest, opts ...grpc.CallOption) (*CancelTravelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelTravelResponse)
	err := c.cc.Invoke(ctx, TravelAgencyService_CancelTravel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelAgencyServiceClient) ApproveTravel(ctx context.Context, in *ApproveTravelRequest, opts ...grpc.CallOption) (*ApproveTravelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApproveTravelResponse)
	err := c.cc.Invoke(ctx, TravelAgencyService_ApproveTravel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelAgencyServiceClient) FinishTravel(ctx context.Context, in *FinishTravelRequest, opts ...grpc.CallOption) (*FinishTravelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FinishTravelResponse)
	err := c.cc.Invoke(ctx, TravelAgencyService_FinishTravel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TravelAgencyServiceServer is the server API for TravelAgencyService service.
// All implementations must embed UnimplementedTravelAgencyServiceServer
// for forward compatibility.
type TravelAgencyServiceServer interface {
	BookTravel(context.Context, *BookTravelRequest) (*BookTravelResponse, error)
	CancelBooking(context.Context, *CancelBookingRequest) (*CancelBookingResponse, error)
	CancelTravel(context.Context, *CancelTravelRequest) (*CancelTravelResponse, error)
	ApproveTravel(context.Context, *ApproveTravelRequest) (*ApproveTravelResponse, error)
	FinishTravel(context.Context, *FinishTravelRequest) (*FinishTravelResponse, error)
	mustEmbedUnimplementedTravelAgencyServiceServer()
}

// UnimplementedTravelAgencyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTravelAgencyServiceServer struct{}

func (UnimplementedTravelAgencyServiceServer) BookTravel(context.Context, *BookTravelRequest) (*BookTravelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookTravel not implemented")
}
func (UnimplementedTravelAgencyServiceServer) CancelBooking(context.Context, *CancelBookingRequest) (*CancelBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedTravelAgencyServiceServer) CancelTravel(context.Context, *CancelTravelRequest) (*CancelTravelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTravel not implemented")
}
func (UnimplementedTravelAgencyServiceServer) ApproveTravel(context.Context, *ApproveTravelRequest) (*ApproveTravelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveTravel not implemented")
}
func (UnimplementedTravelAgencyServiceServer) FinishTravel(context.Context, *FinishTravelRequest) (*FinishTravelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishTravel not implemented")
}
func (UnimplementedTravelAgencyServiceServer) mustEmbedUnimplementedTravelAgencyServiceServer() {}
func (UnimplementedTravelAgencyServiceServer) testEmbeddedByValue()                             {}

// UnsafeTravelAgencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TravelAgencyServiceServer will
// result in compilation errors.
type UnsafeTravelAgencyServiceServer interface {
	mustEmbedUnimplementedTravelAgencyServiceServer()
}

func RegisterTravelAgencyServiceServer(s grpc.ServiceRegistrar, srv TravelAgencyServiceServer) {
	// If the following call pancis, it indicates UnimplementedTravelAgencyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TravelAgencyService_ServiceDesc, srv)
}

func _TravelAgencyService_BookTravel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookTravelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelAgencyServiceServer).BookTravel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TravelAgencyService_BookTravel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelAgencyServiceServer).BookTravel(ctx, req.(*BookTravelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TravelAgencyService_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelAgencyServiceServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TravelAgencyService_CancelBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelAgencyServiceServer).CancelBooking(ctx, req.(*CancelBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TravelAgencyService_CancelTravel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTravelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelAgencyServiceServer).CancelTravel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TravelAgencyService_CancelTravel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelAgencyServiceServer).CancelTravel(ctx, req.(*CancelTravelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TravelAgencyService_ApproveTravel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveTravelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelAgencyServiceServer).ApproveTravel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TravelAgencyService_ApproveTravel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelAgencyServiceServer).ApproveTravel(ctx, req.(*ApproveTravelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TravelAgencyService_FinishTravel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishTravelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelAgencyServiceServer).FinishTravel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TravelAgencyService_FinishTravel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelAgencyServiceServer).FinishTravel(ctx, req.(*FinishTravelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TravelAgencyService_ServiceDesc is the grpc.ServiceDesc for TravelAgencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TravelAgencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TravelAgencyService",
	HandlerType: (*TravelAgencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BookTravel",
			Handler:    _TravelAgencyService_BookTravel_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _TravelAgencyService_CancelBooking_Handler,
		},
		{
			MethodName: "CancelTravel",
			Handler:    _TravelAgencyService_CancelTravel_Handler,
		},
		{
			MethodName: "ApproveTravel",
			Handler:    _TravelAgencyService_ApproveTravel_Handler,
		},
		{
			MethodName: "FinishTravel",
			Handler:    _TravelAgencyService_FinishTravel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "travel.proto",
}