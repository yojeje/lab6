// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: protobuf.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IngenierosClient is the client API for Ingenieros service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IngenierosClient interface {
	EnviarBroker(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*Direccion, error)
	EnviarServidor(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*Reloj, error)
}

type ingenierosClient struct {
	cc grpc.ClientConnInterface
}

func NewIngenierosClient(cc grpc.ClientConnInterface) IngenierosClient {
	return &ingenierosClient{cc}
}

func (c *ingenierosClient) EnviarBroker(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*Direccion, error) {
	out := new(Direccion)
	err := c.cc.Invoke(ctx, "/grpc.Ingenieros/EnviarBroker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingenierosClient) EnviarServidor(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*Reloj, error) {
	out := new(Reloj)
	err := c.cc.Invoke(ctx, "/grpc.Ingenieros/EnviarServidor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngenierosServer is the server API for Ingenieros service.
// All implementations must embed UnimplementedIngenierosServer
// for forward compatibility
type IngenierosServer interface {
	EnviarBroker(context.Context, *Comando) (*Direccion, error)
	EnviarServidor(context.Context, *Comando) (*Reloj, error)
	mustEmbedUnimplementedIngenierosServer()
}

// UnimplementedIngenierosServer must be embedded to have forward compatible implementations.
type UnimplementedIngenierosServer struct {
}

func (UnimplementedIngenierosServer) EnviarBroker(context.Context, *Comando) (*Direccion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarBroker not implemented")
}
func (UnimplementedIngenierosServer) EnviarServidor(context.Context, *Comando) (*Reloj, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarServidor not implemented")
}
func (UnimplementedIngenierosServer) mustEmbedUnimplementedIngenierosServer() {}

// UnsafeIngenierosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IngenierosServer will
// result in compilation errors.
type UnsafeIngenierosServer interface {
	mustEmbedUnimplementedIngenierosServer()
}

func RegisterIngenierosServer(s grpc.ServiceRegistrar, srv IngenierosServer) {
	s.RegisterService(&Ingenieros_ServiceDesc, srv)
}

func _Ingenieros_EnviarBroker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngenierosServer).EnviarBroker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Ingenieros/EnviarBroker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngenierosServer).EnviarBroker(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ingenieros_EnviarServidor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngenierosServer).EnviarServidor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Ingenieros/EnviarServidor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngenierosServer).EnviarServidor(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

// Ingenieros_ServiceDesc is the grpc.ServiceDesc for Ingenieros service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ingenieros_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Ingenieros",
	HandlerType: (*IngenierosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarBroker",
			Handler:    _Ingenieros_EnviarBroker_Handler,
		},
		{
			MethodName: "EnviarServidor",
			Handler:    _Ingenieros_EnviarServidor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf.proto",
}

// KaisClient is the client API for Kais service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KaisClient interface {
	GetEnemigosBroker(ctx context.Context, in *Informacion, opts ...grpc.CallOption) (*Direccion, error)
	GetEnemigosServidor(ctx context.Context, in *Direccion, opts ...grpc.CallOption) (*Enemigos, error)
}

type kaisClient struct {
	cc grpc.ClientConnInterface
}

func NewKaisClient(cc grpc.ClientConnInterface) KaisClient {
	return &kaisClient{cc}
}

func (c *kaisClient) GetEnemigosBroker(ctx context.Context, in *Informacion, opts ...grpc.CallOption) (*Direccion, error) {
	out := new(Direccion)
	err := c.cc.Invoke(ctx, "/grpc.Kais/GetEnemigosBroker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaisClient) GetEnemigosServidor(ctx context.Context, in *Direccion, opts ...grpc.CallOption) (*Enemigos, error) {
	out := new(Enemigos)
	err := c.cc.Invoke(ctx, "/grpc.Kais/GetEnemigosServidor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KaisServer is the server API for Kais service.
// All implementations must embed UnimplementedKaisServer
// for forward compatibility
type KaisServer interface {
	GetEnemigosBroker(context.Context, *Informacion) (*Direccion, error)
	GetEnemigosServidor(context.Context, *Direccion) (*Enemigos, error)
	mustEmbedUnimplementedKaisServer()
}

// UnimplementedKaisServer must be embedded to have forward compatible implementations.
type UnimplementedKaisServer struct {
}

func (UnimplementedKaisServer) GetEnemigosBroker(context.Context, *Informacion) (*Direccion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnemigosBroker not implemented")
}
func (UnimplementedKaisServer) GetEnemigosServidor(context.Context, *Direccion) (*Enemigos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnemigosServidor not implemented")
}
func (UnimplementedKaisServer) mustEmbedUnimplementedKaisServer() {}

// UnsafeKaisServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KaisServer will
// result in compilation errors.
type UnsafeKaisServer interface {
	mustEmbedUnimplementedKaisServer()
}

func RegisterKaisServer(s grpc.ServiceRegistrar, srv KaisServer) {
	s.RegisterService(&Kais_ServiceDesc, srv)
}

func _Kais_GetEnemigosBroker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Informacion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaisServer).GetEnemigosBroker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Kais/GetEnemigosBroker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaisServer).GetEnemigosBroker(ctx, req.(*Informacion))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kais_GetEnemigosServidor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Direccion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaisServer).GetEnemigosServidor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Kais/GetEnemigosServidor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaisServer).GetEnemigosServidor(ctx, req.(*Direccion))
	}
	return interceptor(ctx, in, info, handler)
}

// Kais_ServiceDesc is the grpc.ServiceDesc for Kais service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kais_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Kais",
	HandlerType: (*KaisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEnemigosBroker",
			Handler:    _Kais_GetEnemigosBroker_Handler,
		},
		{
			MethodName: "GetEnemigosServidor",
			Handler:    _Kais_GetEnemigosServidor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf.proto",
}

// ConflictosClient is the client API for Conflictos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConflictosClient interface {
	ResolverConsistencia(ctx context.Context, in *ConsistenciaRequest, opts ...grpc.CallOption) (*Direccion, error)
}

type conflictosClient struct {
	cc grpc.ClientConnInterface
}

func NewConflictosClient(cc grpc.ClientConnInterface) ConflictosClient {
	return &conflictosClient{cc}
}

func (c *conflictosClient) ResolverConsistencia(ctx context.Context, in *ConsistenciaRequest, opts ...grpc.CallOption) (*Direccion, error) {
	out := new(Direccion)
	err := c.cc.Invoke(ctx, "/grpc.Conflictos/ResolverConsistencia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConflictosServer is the server API for Conflictos service.
// All implementations must embed UnimplementedConflictosServer
// for forward compatibility
type ConflictosServer interface {
	ResolverConsistencia(context.Context, *ConsistenciaRequest) (*Direccion, error)
	mustEmbedUnimplementedConflictosServer()
}

// UnimplementedConflictosServer must be embedded to have forward compatible implementations.
type UnimplementedConflictosServer struct {
}

func (UnimplementedConflictosServer) ResolverConsistencia(context.Context, *ConsistenciaRequest) (*Direccion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolverConsistencia not implemented")
}
func (UnimplementedConflictosServer) mustEmbedUnimplementedConflictosServer() {}

// UnsafeConflictosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConflictosServer will
// result in compilation errors.
type UnsafeConflictosServer interface {
	mustEmbedUnimplementedConflictosServer()
}

func RegisterConflictosServer(s grpc.ServiceRegistrar, srv ConflictosServer) {
	s.RegisterService(&Conflictos_ServiceDesc, srv)
}

func _Conflictos_ResolverConsistencia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsistenciaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConflictosServer).ResolverConsistencia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Conflictos/ResolverConsistencia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConflictosServer).ResolverConsistencia(ctx, req.(*ConsistenciaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Conflictos_ServiceDesc is the grpc.ServiceDesc for Conflictos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conflictos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Conflictos",
	HandlerType: (*ConflictosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResolverConsistencia",
			Handler:    _Conflictos_ResolverConsistencia_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf.proto",
}
