package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"meeting-api/lib"
	"meeting-api/serviceclient"
	proto "meeting-api/submodules/common/protob"
	"meeting-api/submodules/common/utils"
	"meeting-api/validator"
	"net/http"
	"strconv"
)

type RoomController struct {
	BaseController
}

func (rc *RoomController) GetRooms(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.GetRooms{}
	if qErr := c.ShouldBind(&params); qErr != nil {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(qErr.Error()))
		return
	}
	//字段验证
	errs := validator.Builder.Struct(params)
	if vErr := validator.TransError(errs); vErr != "" {
		log.Error("GetRooms：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}
	//远程调用 RoomService.GetRooms
	res, err := serviceclient.RoomServiceClient.GetRooms(context.TODO(), &proto.ReqGetRooms{
		SpaceId:  params.SpaceId,
		Page:     params.Page,
		PageSize: params.PageSize,
		Order:    params.Order,
		SortBy:   params.SortBy,
	})
	if err != nil {
		log.Info("未知错误：", err)
		c.JSON(http.StatusOK, lib.NewResponse(1000))
		return
	}

	//需要获取 proto.any格式中的data,需要进行解码
	//解码any格式
	data, paErr := utils.UnmarshalProtoAny(res.Data)
	if paErr != nil {
		log.Error("UnmarshalProtoAny err:", data)
		c.JSON(http.StatusOK, lib.NewResponse(1000))
		return
	}

	//响应
	if res.Code == lib.Err.GetInt64("error.ok") {
		c.JSON(http.StatusOK, lib.Success().WithData(data))
	} else {
		c.JSON(http.StatusOK, lib.NewResponse(res.Code).WithMsg(res.Message).WithData(data))
	}
	return
}

func (rc *RoomController) GetRoom(c *gin.Context) {
	//请求参数处理
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg("id不能为空"))
		return
	}

	//远程调用 RoomService.GetRoom
	res, err := serviceclient.RoomServiceClient.GetRoom(context.TODO(), &proto.ReqGetRoom{
		Id: id,
	})
	if err != nil {
		log.Info("未知错误：", err)
		c.JSON(http.StatusOK, lib.NewResponse(1000))
		return
	}

	//需要获取 proto.any格式中的data,需要进行解码
	//解码any格式
	data, paErr := utils.UnmarshalProtoAny(res.Data)
	if paErr != nil {
		log.Error("UnmarshalProtoAny err:", data)
		c.JSON(http.StatusOK, lib.NewResponse(1000))
		return
	}

	//响应
	if res.Code == lib.Err.GetInt64("error.ok") {
		c.JSON(http.StatusOK, lib.Success().WithData(data))
	} else {
		c.JSON(http.StatusOK, lib.NewResponse(res.Code))
	}
	return
}

func (rc *RoomController) CreateRoom(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.CreateRoom{}
	if qErr := c.ShouldBind(&params); qErr != nil {
		log.Info("qErr:", qErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(qErr.Error()))
		return
	}

	//字段验证
	errs := validator.Builder.Struct(params)
	if vErr := validator.TransError(errs); vErr != "" {
		log.Error("CreateRoom：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}

	//远程调用 RoomService.CreateRoom
	res, err := serviceclient.RoomServiceClient.CreateRoom(context.TODO(), &proto.ReqCreateRoom{
		SpaceId:     params.SpaceId,
		Name:        params.Name,
		OneStatus:   &proto.ReqCreateRoom_Status{Status: params.Status},
		ImageUrl:    params.ImageUrl,
		CapacityMin: params.CapacityMin,
		CapacityMax: params.CapacityMax,
		DeviceIds:   params.DeviceIds,
	})
	if err != nil {
		log.Info("未知错误：", err)
		c.JSON(http.StatusOK, lib.NewResponse(1000))
		return
	}

	//响应
	if res.Code == lib.Err.GetInt64("error.ok") {
		c.JSON(http.StatusOK, lib.Success())
	} else {
		c.JSON(http.StatusOK, lib.NewResponse(res.Code))
	}
	return
}

func (rc RoomController) UpdateRoom(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.UpdateRoom{}
	if qErr := c.ShouldBind(&params); qErr != nil {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(qErr.Error()))
		return
	}
	params.Id, _ = strconv.ParseInt(c.Param("id"), 10, 64)

	//字段验证
	errs := validator.Builder.Struct(params)
	vErr := validator.TransError(errs)
	if vErr != "" {
		log.Error("UpdateRoom：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}

	//远程调用 RoomService.UpdateRoom
	res, err := serviceclient.RoomServiceClient.UpdateRoom(context.TODO(), &proto.ReqUpdateRoom{
		Id:          params.Id,
		SpaceId:     params.SpaceId,
		Name:        params.Name,
		OneStatus:   &proto.ReqUpdateRoom_Status{Status: params.Status},
		ImageUrl:    params.ImageUrl,
		CapacityMin: params.CapacityMin,
		CapacityMax: params.CapacityMax,
		DeviceIds:   params.DeviceIds,
	})
	if err != nil {
		log.Info("未知错误：", err)
		c.JSON(http.StatusOK, lib.NewResponse(1000))
		return
	}

	//响应
	if res.Code == lib.Err.GetInt64("error.ok") {
		c.JSON(http.StatusOK, lib.Success())
	} else {
		c.JSON(http.StatusOK, lib.NewResponse(res.Code))
	}
	return
}

func (rc RoomController) DeleteRoom(c *gin.Context) {
	//获取并绑定请求参数
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	//远程调用 RoomService.DeleteDevice
	res, err := serviceclient.RoomServiceClient.DeleteRoom(context.TODO(), &proto.ReqDeleteRoom{
		Id: id,
	})
	if err != nil {
		log.Info("未知错误：", err)
		c.JSON(http.StatusOK, lib.NewResponse(1000))
		return
	}

	//响应
	if res.Code == lib.Err.GetInt64("error.ok") {
		c.JSON(http.StatusOK, lib.Success())
	} else {
		c.JSON(http.StatusOK, lib.NewResponse(res.Code))
	}
	return
}
