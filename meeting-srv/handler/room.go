package handler

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
	"meeting-srv/lib"
	"meeting-srv/model"
	"meeting-srv/repository"
	proto "meeting-srv/submodules/common/protob"
	"strings"
)

type RoomService struct {
	repo            repository.RoomRepository
	repoDevice      repository.DeviceRepository
	repoReservation repository.ReservationRepository
}

//查询会议室列表
func (r RoomService) GetRooms(ctx context.Context, req *proto.ReqGetRooms, res *proto.Response) error {
	//数据查询
	dbRes, count, getErr := r.repo.GetRooms(
		req.SpaceId,
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

//查询会议室详情
func (r RoomService) GetRoom(ctx context.Context, req *proto.ReqGetRoom, res *proto.Response) error {
	dbRes, getErr := r.repo.GetRoom(req.Id)
	if getErr != nil {
		if getErr == gorm.ErrRecordNotFound {
			lib.NewResponse(20109).ToProtoRes(res)
			return nil
		}
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}

	lib.Success().WithData(dbRes).ToProtoRes(res)
	return nil
}

//新增会议室
func (r RoomService) CreateRoom(ctx context.Context, req *proto.ReqCreateRoom, res *proto.Response) error {
	//验证会议室名称是否重复
	errCode := r.validateNameRepeat(req.Name, req.SpaceId, 0)
	if errCode != lib.Err.GetInt64("error.ok") {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}

	//批量获取设备
	deviceIds := strings.Split(req.DeviceIds, ",")
	devices, gErr := r.repoDevice.GetDevicesByIds(deviceIds)
	if gErr != nil {
		log.Error("GetDevicesByIds error:", gErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	//创建会议室
	createErr := r.repo.CreateRoom(&model.Room{
		SpaceId:     req.SpaceId,
		Name:        req.Name,
		Status:      req.GetStatus(),
		ImageUrl:    req.ImageUrl,
		CapacityMin: req.CapacityMin,
		CapacityMax: req.CapacityMax,
		Devices:     devices,
	})
	if createErr != nil {
		log.Error("repo.CreateRoom error:", createErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	lib.Success().ToProtoRes(res)
	return nil
}

//编辑会议室
func (r RoomService) UpdateRoom(ctx context.Context, req *proto.ReqUpdateRoom, res *proto.Response) error {
	//验证会议室
	errCode := r.validateRoom(req.Id)
	ok := lib.Err.GetInt64("error.ok")
	if errCode != ok {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}
	//验证会议室名称是否重复
	errCode = r.validateNameRepeat(req.Name, req.SpaceId, req.Id)
	if errCode != ok {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}

	//批量获取设备
	deviceIds := strings.Split(req.DeviceIds, ",")
	devices, gErr := r.repoDevice.GetDevicesByIds(deviceIds)
	if gErr != nil {
		log.Error("GetDevicesByIds error:", gErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}

	//更新会议室
	upErr := r.repo.UpdateRoom(&model.Room{
		Id:          req.Id,
		SpaceId:     req.SpaceId,
		Name:        req.Name,
		Status:      req.GetStatus(),
		ImageUrl:    req.ImageUrl,
		CapacityMin: req.CapacityMin,
		CapacityMax: req.CapacityMax,
		Devices:     devices,
	})
	if upErr != nil {
		log.Error("repo.UpdateRoom error:", upErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	lib.Success().ToProtoRes(res)
	return nil
}

//删除会议室
func (r RoomService) DeleteRoom(ctx context.Context, req *proto.ReqDeleteRoom, res *proto.Response) error {
	//验证会议室
	errCode := r.validateRoom(req.Id)
	if errCode != lib.Err.GetInt64("error.ok") {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}

	//验证会议室是否已预约
	count, cErr := r.repoReservation.GetCountOfRoomReservation(req.Id)
	if cErr != nil {
		log.Error("repoReservation.GetCountOfRoomReservation error:", cErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	if count != 0 {
		lib.NewResponse(20113).ToProtoRes(res)
		return nil
	}

	//删除会议室
	delErr := r.repo.DelRoom(&model.Room{
		Id: req.Id,
	})
	if delErr != nil {
		log.Error("repo.DelRoom error:", delErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	lib.Success().ToProtoRes(res)
	return nil
}

//验证会议室是否存在
func (r *RoomService) validateRoom(id int64) (code int64) {
	_, err := r.repo.GetRoom(id)
	if err == gorm.ErrRecordNotFound {
		code = 20109
		return
	}
	if err != nil {
		log.Error("validateRoom error:", err)
		code = 1001
		return
	}
	code = lib.Err.GetInt64("error.ok")
	return
}

//验证会议室名称是否重复
func (r RoomService) validateNameRepeat(name string, spaceId int64, id int64) (code int64) {
	room, err := r.repo.ValidateNameRepeat(name, spaceId, id)
	if err != nil {
		log.Error("ValidateNameRepeat error:", err)
		if err != gorm.ErrRecordNotFound {
			code = 1001
			return
		}
	}
	if room.Id != 0 {
		code = 20108
		return
	}
	code = lib.Err.GetInt64("error.ok")
	return
}
