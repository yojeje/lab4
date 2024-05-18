// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: Director.proto

package Director

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

// MontoClient is the client API for Monto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MontoClient interface {
	ConsultarMonto(ctx context.Context, in *MontoRequest, opts ...grpc.CallOption) (*Vacio, error)
}

type montoClient struct {
	cc grpc.ClientConnInterface
}

func NewMontoClient(cc grpc.ClientConnInterface) MontoClient {
	return &montoClient{cc}
}

func (c *montoClient) ConsultarMonto(ctx context.Context, in *MontoRequest, opts ...grpc.CallOption) (*Vacio, error) {
	out := new(Vacio)
	err := c.cc.Invoke(ctx, "/grpc.monto/ConsultarMonto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MontoServer is the server API for Monto service.
// All implementations must embed UnimplementedMontoServer
// for forward compatibility
type MontoServer interface {
	ConsultarMonto(context.Context, *MontoRequest) (*Vacio, error)
	mustEmbedUnimplementedMontoServer()
}

// UnimplementedMontoServer must be embedded to have forward compatible implementations.
type UnimplementedMontoServer struct {
}

func (UnimplementedMontoServer) ConsultarMonto(context.Context, *MontoRequest) (*Vacio, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultarMonto not implemented")
}
func (UnimplementedMontoServer) mustEmbedUnimplementedMontoServer() {}

// UnsafeMontoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MontoServer will
// result in compilation errors.
type UnsafeMontoServer interface {
	mustEmbedUnimplementedMontoServer()
}

func RegisterMontoServer(s grpc.ServiceRegistrar, srv MontoServer) {
	s.RegisterService(&Monto_ServiceDesc, srv)
}

func _Monto_ConsultarMonto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MontoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MontoServer).ConsultarMonto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.monto/ConsultarMonto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MontoServer).ConsultarMonto(ctx, req.(*MontoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Monto_ServiceDesc is the grpc.ServiceDesc for Monto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Monto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.monto",
	HandlerType: (*MontoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConsultarMonto",
			Handler:    _Monto_ConsultarMonto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Director.proto",
}

// MercenarioClient is the client API for Mercenario service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MercenarioClient interface {
	IdentificarMercenario(ctx context.Context, in *Vacio, opts ...grpc.CallOption) (*IDRequest, error)
}

type mercenarioClient struct {
	cc grpc.ClientConnInterface
}

func NewMercenarioClient(cc grpc.ClientConnInterface) MercenarioClient {
	return &mercenarioClient{cc}
}

func (c *mercenarioClient) IdentificarMercenario(ctx context.Context, in *Vacio, opts ...grpc.CallOption) (*IDRequest, error) {
	out := new(IDRequest)
	err := c.cc.Invoke(ctx, "/grpc.Mercenario/IdentificarMercenario", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MercenarioServer is the server API for Mercenario service.
// All implementations must embed UnimplementedMercenarioServer
// for forward compatibility
type MercenarioServer interface {
	IdentificarMercenario(context.Context, *Vacio) (*IDRequest, error)
	mustEmbedUnimplementedMercenarioServer()
}

// UnimplementedMercenarioServer must be embedded to have forward compatible implementations.
type UnimplementedMercenarioServer struct {
}

func (UnimplementedMercenarioServer) IdentificarMercenario(context.Context, *Vacio) (*IDRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IdentificarMercenario not implemented")
}
func (UnimplementedMercenarioServer) mustEmbedUnimplementedMercenarioServer() {}

// UnsafeMercenarioServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MercenarioServer will
// result in compilation errors.
type UnsafeMercenarioServer interface {
	mustEmbedUnimplementedMercenarioServer()
}

func RegisterMercenarioServer(s grpc.ServiceRegistrar, srv MercenarioServer) {
	s.RegisterService(&Mercenario_ServiceDesc, srv)
}

func _Mercenario_IdentificarMercenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vacio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercenarioServer).IdentificarMercenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Mercenario/IdentificarMercenario",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercenarioServer).IdentificarMercenario(ctx, req.(*Vacio))
	}
	return interceptor(ctx, in, info, handler)
}

// Mercenario_ServiceDesc is the grpc.ServiceDesc for Mercenario service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mercenario_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Mercenario",
	HandlerType: (*MercenarioServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IdentificarMercenario",
			Handler:    _Mercenario_IdentificarMercenario_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Director.proto",
}

// ConectadosClient is the client API for Conectados service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConectadosClient interface {
	UsuariosConectados(ctx context.Context, in *Vacio, opts ...grpc.CallOption) (*Conectados, error)
}

type conectadosClient struct {
	cc grpc.ClientConnInterface
}

func NewConectadosClient(cc grpc.ClientConnInterface) ConectadosClient {
	return &conectadosClient{cc}
}

func (c *conectadosClient) UsuariosConectados(ctx context.Context, in *Vacio, opts ...grpc.CallOption) (*Conectados, error) {
	out := new(Conectados)
	err := c.cc.Invoke(ctx, "/grpc.Conectados/UsuariosConectados", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConectadosServer is the server API for Conectados service.
// All implementations must embed UnimplementedConectadosServer
// for forward compatibility
type ConectadosServer interface {
	UsuariosConectados(context.Context, *Vacio) (*Conectados, error)
	mustEmbedUnimplementedConectadosServer()
}

// UnimplementedConectadosServer must be embedded to have forward compatible implementations.
type UnimplementedConectadosServer struct {
}

func (UnimplementedConectadosServer) UsuariosConectados(context.Context, *Vacio) (*Conectados, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsuariosConectados not implemented")
}
func (UnimplementedConectadosServer) mustEmbedUnimplementedConectadosServer() {}

// UnsafeConectadosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConectadosServer will
// result in compilation errors.
type UnsafeConectadosServer interface {
	mustEmbedUnimplementedConectadosServer()
}

func RegisterConectadosServer(s grpc.ServiceRegistrar, srv ConectadosServer) {
	s.RegisterService(&Conectados_ServiceDesc, srv)
}

func _Conectados_UsuariosConectados_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vacio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConectadosServer).UsuariosConectados(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Conectados/UsuariosConectados",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConectadosServer).UsuariosConectados(ctx, req.(*Vacio))
	}
	return interceptor(ctx, in, info, handler)
}

// Conectados_ServiceDesc is the grpc.ServiceDesc for Conectados service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conectados_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Conectados",
	HandlerType: (*ConectadosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UsuariosConectados",
			Handler:    _Conectados_UsuariosConectados_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Director.proto",
}

// Piso1Client is the client API for Piso1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Piso1Client interface {
	ProbabilidadArma(ctx context.Context, in *Arma, opts ...grpc.CallOption) (*Supervivencia, error)
}

type piso1Client struct {
	cc grpc.ClientConnInterface
}

func NewPiso1Client(cc grpc.ClientConnInterface) Piso1Client {
	return &piso1Client{cc}
}

func (c *piso1Client) ProbabilidadArma(ctx context.Context, in *Arma, opts ...grpc.CallOption) (*Supervivencia, error) {
	out := new(Supervivencia)
	err := c.cc.Invoke(ctx, "/grpc.Piso1/ProbabilidadArma", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Piso1Server is the server API for Piso1 service.
// All implementations must embed UnimplementedPiso1Server
// for forward compatibility
type Piso1Server interface {
	ProbabilidadArma(context.Context, *Arma) (*Supervivencia, error)
	mustEmbedUnimplementedPiso1Server()
}

// UnimplementedPiso1Server must be embedded to have forward compatible implementations.
type UnimplementedPiso1Server struct {
}

func (UnimplementedPiso1Server) ProbabilidadArma(context.Context, *Arma) (*Supervivencia, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProbabilidadArma not implemented")
}
func (UnimplementedPiso1Server) mustEmbedUnimplementedPiso1Server() {}

// UnsafePiso1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Piso1Server will
// result in compilation errors.
type UnsafePiso1Server interface {
	mustEmbedUnimplementedPiso1Server()
}

func RegisterPiso1Server(s grpc.ServiceRegistrar, srv Piso1Server) {
	s.RegisterService(&Piso1_ServiceDesc, srv)
}

func _Piso1_ProbabilidadArma_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Arma)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Piso1Server).ProbabilidadArma(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Piso1/ProbabilidadArma",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Piso1Server).ProbabilidadArma(ctx, req.(*Arma))
	}
	return interceptor(ctx, in, info, handler)
}

// Piso1_ServiceDesc is the grpc.ServiceDesc for Piso1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Piso1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Piso1",
	HandlerType: (*Piso1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProbabilidadArma",
			Handler:    _Piso1_ProbabilidadArma_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Director.proto",
}

// Piso2Client is the client API for Piso2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Piso2Client interface {
	ProbabilidadCaminoCorrecto(ctx context.Context, in *CaminoElegido, opts ...grpc.CallOption) (*CaminoCorrecto, error)
}

type piso2Client struct {
	cc grpc.ClientConnInterface
}

func NewPiso2Client(cc grpc.ClientConnInterface) Piso2Client {
	return &piso2Client{cc}
}

func (c *piso2Client) ProbabilidadCaminoCorrecto(ctx context.Context, in *CaminoElegido, opts ...grpc.CallOption) (*CaminoCorrecto, error) {
	out := new(CaminoCorrecto)
	err := c.cc.Invoke(ctx, "/grpc.Piso2/ProbabilidadCaminoCorrecto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Piso2Server is the server API for Piso2 service.
// All implementations must embed UnimplementedPiso2Server
// for forward compatibility
type Piso2Server interface {
	ProbabilidadCaminoCorrecto(context.Context, *CaminoElegido) (*CaminoCorrecto, error)
	mustEmbedUnimplementedPiso2Server()
}

// UnimplementedPiso2Server must be embedded to have forward compatible implementations.
type UnimplementedPiso2Server struct {
}

func (UnimplementedPiso2Server) ProbabilidadCaminoCorrecto(context.Context, *CaminoElegido) (*CaminoCorrecto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProbabilidadCaminoCorrecto not implemented")
}
func (UnimplementedPiso2Server) mustEmbedUnimplementedPiso2Server() {}

// UnsafePiso2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Piso2Server will
// result in compilation errors.
type UnsafePiso2Server interface {
	mustEmbedUnimplementedPiso2Server()
}

func RegisterPiso2Server(s grpc.ServiceRegistrar, srv Piso2Server) {
	s.RegisterService(&Piso2_ServiceDesc, srv)
}

func _Piso2_ProbabilidadCaminoCorrecto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaminoElegido)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Piso2Server).ProbabilidadCaminoCorrecto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Piso2/ProbabilidadCaminoCorrecto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Piso2Server).ProbabilidadCaminoCorrecto(ctx, req.(*CaminoElegido))
	}
	return interceptor(ctx, in, info, handler)
}

// Piso2_ServiceDesc is the grpc.ServiceDesc for Piso2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Piso2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Piso2",
	HandlerType: (*Piso2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProbabilidadCaminoCorrecto",
			Handler:    _Piso2_ProbabilidadCaminoCorrecto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Director.proto",
}

// Piso3Client is the client API for Piso3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Piso3Client interface {
	ProbabiliadEnfrentamiento(ctx context.Context, in *Vacio, opts ...grpc.CallOption) (*Enfrentamiento, error)
}

type piso3Client struct {
	cc grpc.ClientConnInterface
}

func NewPiso3Client(cc grpc.ClientConnInterface) Piso3Client {
	return &piso3Client{cc}
}

func (c *piso3Client) ProbabiliadEnfrentamiento(ctx context.Context, in *Vacio, opts ...grpc.CallOption) (*Enfrentamiento, error) {
	out := new(Enfrentamiento)
	err := c.cc.Invoke(ctx, "/grpc.Piso3/ProbabiliadEnfrentamiento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Piso3Server is the server API for Piso3 service.
// All implementations must embed UnimplementedPiso3Server
// for forward compatibility
type Piso3Server interface {
	ProbabiliadEnfrentamiento(context.Context, *Vacio) (*Enfrentamiento, error)
	mustEmbedUnimplementedPiso3Server()
}

// UnimplementedPiso3Server must be embedded to have forward compatible implementations.
type UnimplementedPiso3Server struct {
}

func (UnimplementedPiso3Server) ProbabiliadEnfrentamiento(context.Context, *Vacio) (*Enfrentamiento, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProbabiliadEnfrentamiento not implemented")
}
func (UnimplementedPiso3Server) mustEmbedUnimplementedPiso3Server() {}

// UnsafePiso3Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Piso3Server will
// result in compilation errors.
type UnsafePiso3Server interface {
	mustEmbedUnimplementedPiso3Server()
}

func RegisterPiso3Server(s grpc.ServiceRegistrar, srv Piso3Server) {
	s.RegisterService(&Piso3_ServiceDesc, srv)
}

func _Piso3_ProbabiliadEnfrentamiento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vacio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Piso3Server).ProbabiliadEnfrentamiento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Piso3/ProbabiliadEnfrentamiento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Piso3Server).ProbabiliadEnfrentamiento(ctx, req.(*Vacio))
	}
	return interceptor(ctx, in, info, handler)
}

// Piso3_ServiceDesc is the grpc.ServiceDesc for Piso3 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Piso3_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Piso3",
	HandlerType: (*Piso3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProbabiliadEnfrentamiento",
			Handler:    _Piso3_ProbabiliadEnfrentamiento_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Director.proto",
}
