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

// Api Endpoints for MeetingService service

func NewMeetingServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MeetingService service

type MeetingService interface {
	//会议室预订
	MeetingRegister(ctx context.Context, in *ReqMeetingRegister, opts ...client.CallOption) (*ResMeetingRegister, error)
}

type meetingService struct {
	c    client.Client
	name string
}

func NewMeetingService(name string, c client.Client) MeetingService {
	return &meetingService{
		c:    c,
		name: name,
	}
}

func (c *meetingService) MeetingRegister(ctx context.Context, in *ReqMeetingRegister, opts ...client.CallOption) (*ResMeetingRegister, error) {
	req := c.c.NewRequest(c.name, "MeetingService.MeetingRegister", in)
	out := new(ResMeetingRegister)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MeetingService service

type MeetingServiceHandler interface {
	//会议室预订
	MeetingRegister(context.Context, *ReqMeetingRegister, *ResMeetingRegister) error
}

func RegisterMeetingServiceHandler(s server.Server, hdlr MeetingServiceHandler, opts ...server.HandlerOption) error {
	type meetingService interface {
		MeetingRegister(ctx context.Context, in *ReqMeetingRegister, out *ResMeetingRegister) error
	}
	type MeetingService struct {
		meetingService
	}
	h := &meetingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&MeetingService{h}, opts...))
}

type meetingServiceHandler struct {
	MeetingServiceHandler
}

func (h *meetingServiceHandler) MeetingRegister(ctx context.Context, in *ReqMeetingRegister, out *ResMeetingRegister) error {
	return h.MeetingServiceHandler.MeetingRegister(ctx, in, out)
}
