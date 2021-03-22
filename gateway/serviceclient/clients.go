package serviceclient

import (
	proto "gateway/submodules/common/protob"
	"github.com/micro/go-micro/v2/registry"
)

var (
	UserServiceClient proto.UserService
	MeetingApiNode    *registry.Node
)
