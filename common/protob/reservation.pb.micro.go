// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/meeting/reservation.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ReservationService service

func NewReservationServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ReservationService service

type ReservationService interface {
	//会议预订列表
	GetReservations(ctx context.Context, in *ReqGetReservations, opts ...client.CallOption) (*Response, error)
	//会议详情
	GetReservation(ctx context.Context, in *ReqGetReservation, opts ...client.CallOption) (*Response, error)
	//预订会议室
	CreateReservation(ctx context.Context, in *ReqCreateReservation, opts ...client.CallOption) (*Response, error)
}

type reservationService struct {
	c    client.Client
	name string
}

func NewReservationService(name string, c client.Client) ReservationService {
	return &reservationService{
		c:    c,
		name: name,
	}
}

func (c *reservationService) GetReservations(ctx context.Context, in *ReqGetReservations, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ReservationService.GetReservations", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) GetReservation(ctx context.Context, in *ReqGetReservation, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ReservationService.GetReservation", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) CreateReservation(ctx context.Context, in *ReqCreateReservation, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ReservationService.CreateReservation", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReservationService service

type ReservationServiceHandler interface {
	//会议预订列表
	GetReservations(context.Context, *ReqGetReservations, *Response) error
	//会议详情
	GetReservation(context.Context, *ReqGetReservation, *Response) error
	//预订会议室
	CreateReservation(context.Context, *ReqCreateReservation, *Response) error
}

func RegisterReservationServiceHandler(s server.Server, hdlr ReservationServiceHandler, opts ...server.HandlerOption) error {
	type reservationService interface {
		GetReservations(ctx context.Context, in *ReqGetReservations, out *Response) error
		GetReservation(ctx context.Context, in *ReqGetReservation, out *Response) error
		CreateReservation(ctx context.Context, in *ReqCreateReservation, out *Response) error
	}
	type ReservationService struct {
		reservationService
	}
	h := &reservationServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ReservationService{h}, opts...))
}

type reservationServiceHandler struct {
	ReservationServiceHandler
}

func (h *reservationServiceHandler) GetReservations(ctx context.Context, in *ReqGetReservations, out *Response) error {
	return h.ReservationServiceHandler.GetReservations(ctx, in, out)
}

func (h *reservationServiceHandler) GetReservation(ctx context.Context, in *ReqGetReservation, out *Response) error {
	return h.ReservationServiceHandler.GetReservation(ctx, in, out)
}

func (h *reservationServiceHandler) CreateReservation(ctx context.Context, in *ReqCreateReservation, out *Response) error {
	return h.ReservationServiceHandler.CreateReservation(ctx, in, out)
}