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

type DeviceService struct {
	repo repository.DeviceRepository
}

//查询设备列表
func (d *DeviceService) GetDevices(ctx context.Context, req *proto.ReqGetDevices, res *proto.Response) error {
	//数据查询
	dbRes, count, getErr := d.repo.GetDevices(
		req.Name,
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

//查询设备详情
func (d *DeviceService) GetDevice(ctx context.Context, req *proto.ReqGetDevice, res *proto.Response) error {
	dbRes, getErr := d.repo.GetDevice(req.Id)
	if getErr != nil {
		if getErr == gorm.ErrRecordNotFound {
			lib.NewResponse(20106).ToProtoRes(res)
			return nil
		}
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}

	lib.Success().WithData(dbRes).ToProtoRes(res)
	return nil
}

//新增设备
func (d *DeviceService) CreateDevice(ctx context.Context, req *proto.ReqCreateDevice, res *proto.Response) error {
	//验证设备名称是否重复
	errCode := d.validateDeviceNameRepeat(req.Name, 0)
	if errCode != lib.Err.GetInt64("error.ok") {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}

	//创建设备
	device := &model.Device{
		Name:     req.Name,
		Sn:       req.Sn,
		ImageUrl: req.ImageUrl,
	}
	createErr := d.repo.CreateDevice(device)
	if createErr != nil {
		log.Error("repo.CreateDevice error:", createErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	log.Info("device:", device)
	lib.Success().WithData(device).ToProtoRes(res)
	return nil
}

//编辑设备
func (d *DeviceService) UpdateDevice(ctx context.Context, req *proto.ReqUpdateDevice, res *proto.Response) error {
	//验证设备
	errCode := d.validateDevice(req.Id)
	ok := lib.Err.GetInt64("error.ok")
	if errCode != ok {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}
	//验证设备名称是否重复
	errCode = d.validateDeviceNameRepeat(req.Name, req.Id)
	if errCode != ok {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}

	//更新设备
	upErr := d.repo.UpdateDevice(&model.Device{
		Id:       req.Id,
		Name:     req.Name,
		Sn:       req.Sn,
		ImageUrl: req.ImageUrl,
	})
	if upErr != nil {
		log.Error("repo.UpdateDevice error:", upErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	lib.Success().ToProtoRes(res)
	return nil
}

//删除设备
func (d *DeviceService) DeleteDevice(ctx context.Context, req *proto.ReqDeleteDevice, res *proto.Response) error {
	errCode := d.validateDevice(req.Id)
	if errCode != lib.Err.GetInt64("error.ok") {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}

	delErr := d.repo.DeleteDevice(&model.Device{
		Id: req.Id,
	})
	if delErr != nil {
		log.Error("repo.DelDevice error:", delErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	lib.Success().ToProtoRes(res)
	return nil
}

//验证设备是否存在
func (d *DeviceService) validateDevice(id int64) (code int64) {
	_, err := d.repo.GetDevice(id)
	if err == gorm.ErrRecordNotFound {
		code = 20106
		return
	}
	if err != nil {
		log.Error("validateDevice error:", err)
		code = 1001
	}
	code = lib.Err.GetInt64("error.ok")
	return
}

//验证设备名称是否重复
func (d *DeviceService) validateDeviceNameRepeat(name string, id int64) (code int64) {
	device, err := d.repo.ValidateNameRepeat(name, id)
	if err != nil {
		log.Error("ValidateNameRepeat error:", err)
		if err != gorm.ErrRecordNotFound {
			code = 1001
			return
		}
	}
	if device.Id != 0 {
		code = 20105
		return
	}
	code = lib.Err.GetInt64("error.ok")
	return
}
