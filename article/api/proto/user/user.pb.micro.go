// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserSrv service

type UserSrvService interface {
	// 登录
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// 获取用户
	User(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (UserSrv_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (UserSrv_PingPongService, error)
}

type userSrvService struct {
	c    client.Client
	name string
}

func NewUserSrvService(name string, c client.Client) UserSrvService {
	return &userSrvService{
		c:    c,
		name: name,
	}
}

func (c *userSrvService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "UserSrv.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) User(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "UserSrv.User", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (UserSrv_StreamService, error) {
	req := c.c.NewRequest(c.name, "UserSrv.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &userSrvServiceStream{stream}, nil
}

type UserSrv_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type userSrvServiceStream struct {
	stream client.Stream
}

func (x *userSrvServiceStream) Close() error {
	return x.stream.Close()
}

func (x *userSrvServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *userSrvServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userSrvServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userSrvServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userSrvService) PingPong(ctx context.Context, opts ...client.CallOption) (UserSrv_PingPongService, error) {
	req := c.c.NewRequest(c.name, "UserSrv.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &userSrvServicePingPong{stream}, nil
}

type UserSrv_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type userSrvServicePingPong struct {
	stream client.Stream
}

func (x *userSrvServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *userSrvServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *userSrvServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userSrvServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userSrvServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *userSrvServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for UserSrv service

type UserSrvHandler interface {
	// 登录
	Login(context.Context, *LoginRequest, *LoginResponse) error
	// 获取用户
	User(context.Context, *UserRequest, *UserResponse) error
	Stream(context.Context, *StreamingRequest, UserSrv_StreamStream) error
	PingPong(context.Context, UserSrv_PingPongStream) error
}

func RegisterUserSrvHandler(s server.Server, hdlr UserSrvHandler, opts ...server.HandlerOption) error {
	type userSrv interface {
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		User(ctx context.Context, in *UserRequest, out *UserResponse) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type UserSrv struct {
		userSrv
	}
	h := &userSrvHandler{hdlr}
	return s.Handle(s.NewHandler(&UserSrv{h}, opts...))
}

type userSrvHandler struct {
	UserSrvHandler
}

func (h *userSrvHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserSrvHandler.Login(ctx, in, out)
}

func (h *userSrvHandler) User(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserSrvHandler.User(ctx, in, out)
}

func (h *userSrvHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.UserSrvHandler.Stream(ctx, m, &userSrvStreamStream{stream})
}

type UserSrv_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type userSrvStreamStream struct {
	stream server.Stream
}

func (x *userSrvStreamStream) Close() error {
	return x.stream.Close()
}

func (x *userSrvStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *userSrvStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userSrvStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userSrvStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *userSrvHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.UserSrvHandler.PingPong(ctx, &userSrvPingPongStream{stream})
}

type UserSrv_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type userSrvPingPongStream struct {
	stream server.Stream
}

func (x *userSrvPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *userSrvPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *userSrvPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userSrvPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userSrvPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *userSrvPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}