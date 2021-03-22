package handler

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
	"meeting-srv/lib"
	"meeting-srv/model"
	"meeting-srv/repository"
	proto "meeting-srv/submodules/common/protob"
)

type SpaceService struct {
	repo     repository.SpaceRepository
	repoRoom repository.RoomRepository
}

//获取所有地点
func (s *SpaceService) GetAllSpaces(ctx context.Context, req *proto.ReqGetAllSpaces, res *proto.Response) error {
	//数据查询
	dbRes, getErr := s.repo.GetAllSpaces(
		req.Order,
		req.SortBy,
	)
	if getErr != nil {
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}

	lib.Success().WithData(dbRes).ToProtoRes(res)
	return nil
}

//查询地点列表
func (s *SpaceService) GetSpaces(ctx context.Context, req *proto.ReqGetSpaces, res *proto.Response) error {
	//数据查询
	dbRes, count, getErr := s.repo.GetSpaces(
		req.GetStatus(),
		req.PageSize,
		req.Page,
		req.Order,
		req.SortBy,
	)
	if getErr != nil {
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}

	data := make(map[string]interface{})
	data["total"] = count
	data["data"] = dbRes
	lib.Success().WithData(data).ToProtoRes(res)
	return nil
}

//查询地点详情
func (s *SpaceService) GetSpace(ctx context.Context, req *proto.ReqGetSpace, res *proto.Response) error {
	dbRes, getErr := s.repo.GetSpace(req.Id)
	if getErr != nil {
		if getErr == gorm.ErrRecordNotFound {
			lib.NewResponse(20103).ToProtoRes(res)
			return nil
		}
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}

	lib.Success().WithData(dbRes).ToProtoRes(res)
	return nil
}

//新增地点
func (s *SpaceService) CreateSpace(ctx context.Context, req *proto.ReqCreateSpace, res *proto.Response) error {
	//验证名称是否重复
	errCode := s.validateNameRepeat(req.Name, 0)
	if errCode != lib.Err.GetInt64("error.ok") {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}

	//创建地点
	createErr := s.repo.CreateSpace(&model.Space{
		Name:   req.Name,
		Lng:    req.Lng,
		Lat:    req.Lat,
		Status: req.GetStatus(),
	})
	if createErr != nil {
		log.Error("repo.CreateSpace error:", createErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	lib.Success().ToProtoRes(res)
	return nil
}

//编辑地点
func (s *SpaceService) UpdateSpace(ctx context.Context, req *proto.ReqUpdateSpace, res *proto.Response) error {
	//验证地点
	errCode := s.validateSpace(req.Id)
	ok := lib.Err.GetInt64("error.ok")
	if errCode != ok {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}

	//验证地点名称是否重复
	errCode = s.validateNameRepeat(req.Name, req.Id)
	if errCode != ok {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}

	//更新地点
	upErr := s.repo.UpdateSpace(&model.Space{
		Id:     req.Id,
		Name:   req.Name,
		Lng:    req.Lng,
		Lat:    req.Lat,
		Status: req.GetStatus(),
	})
	if upErr != nil {
		log.Error("repo.UpdateSpace error:", upErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	lib.Success().ToProtoRes(res)
	return nil
}

//删除地点
func (s *SpaceService) DelSpace(ctx context.Context, req *proto.ReqDelSpace, res *proto.Response) error {
	//验证地点
	errCode := s.validateSpace(req.Id)
	if errCode != lib.Err.GetInt64("error.ok") {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}

	//检查是否存在会议室绑定了地点
	rooms, sErr := s.repoRoom.GetRoomsBySpaceId(req.Id)
	if sErr != nil {
		log.Error("repoRoom.GetRoomsBySpaceId error:", sErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	if cap(rooms) != 0 {
		lib.NewResponse(20112).ToProtoRes(res)
		return nil
	}
	//删除地点
	delErr := s.repo.DelSpace(&model.Space{
		Id: req.Id,
	})
	if delErr != nil {
		log.Error("repo.DelSpace error:", delErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	lib.Success().ToProtoRes(res)
	return nil
}

//启用/禁用地点
func (s *SpaceService) UpdateSpaceStatus(ctx context.Context, req *proto.ReqUpdateSpaceStatus, res *proto.Response) error {
	//验证地点
	errCode := s.validateSpace(req.Id)
	if errCode != lib.Err.GetInt64("error.ok") {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}
	upErr := s.repo.UpdateSpaceStatus(&model.Space{
		Id:     req.Id,
		Status: req.GetStatus(),
	})
	if upErr != nil {
		log.Error("repo.UpdateSpaceStatus error:", upErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	lib.Success().ToProtoRes(res)
	return nil
}

//验证validateSpace是否存在
func (s *SpaceService) validateSpace(id int64) (code int64) {
	_, err := s.repo.GetSpace(id)
	if err == gorm.ErrRecordNotFound {
		code = 20103
		return
	}
	if err != nil {
		log.Error("validateSpace error:", err)
		code = 1001
		return
	}
	code = lib.Err.GetInt64("error.ok")
	return
}

//验证地点名称是否重复
func (s *SpaceService) validateNameRepeat(name string, id int64) (code int64) {
	device, err := s.repo.ValidateNameRepeat(name, id)
	if err != nil {
		log.Error("ValidateNameRepeat error:", err)
		if err != gorm.ErrRecordNotFound {
			code = 1001
			return
		}
	}
	if device.Id != 0 {
		code = 20102
		return
	}
	code = lib.Err.GetInt64("error.ok")
	return
}
