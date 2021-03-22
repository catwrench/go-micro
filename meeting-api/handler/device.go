package handler

import (
	"context"
	"fmt"
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

type DeviceController struct {
	BaseController
}

//查询设备列表
func (dc *DeviceController) GetDevices(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.GetDevices{}
	if qErr := c.ShouldBind(&params); qErr != nil {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(qErr.Error()))
		return
	}

	//字段验证
	errs := validator.Builder.Struct(params)
	if vErr := validator.TransError(errs); vErr != "" {
		log.Error("GetDevices：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}
	fmt.Println("请求 params:", params)

	//远程调用 DeviceService.GetDevices
	res, err := serviceclient.DeviceServiceClient.GetDevices(context.TODO(), &proto.ReqGetDevices{
		Page:     params.Page,
		PageSize: params.PageSize,
		Order:    params.Order,
		SortBy:   params.SortBy,
		Name:     params.Name,
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
	}

	//响应
	if res.Code == lib.Err.GetInt64("error.ok") {
		c.JSON(http.StatusOK, lib.Success().WithData(data))
	} else {
		c.JSON(http.StatusOK, lib.NewResponse(res.Code).WithMsg(res.Message).WithData(data))
	}
}

//查询设备详情
func (dc *DeviceController) GetDevice(c *gin.Context) {
	//请求参数处理
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg("id不能为空"))
		return
	}

	//远程调用 DeviceService.GetDevice
	res, err := serviceclient.DeviceServiceClient.GetDevice(context.TODO(), &proto.ReqGetDevice{
		Id: id,
	})
	if err != nil {
		log.Info("未知错误：", err)
		c.JSON(http.StatusOK, lib.NewResponse(1000))
		return
	}
	fmt.Println("DeviceService.GetDevice res:", res)

	//需要获取 proto.any格式中的data,需要进行解码
	//解码any格式
	data, paErr := utils.UnmarshalProtoAny(res.Data)
	if paErr != nil {
		log.Error("UnmarshalProtoAny err:", data)
		c.JSON(http.StatusOK, lib.NewResponse(1000))
	}

	//响应
	if res.Code == lib.Err.GetInt64("error.ok") {
		c.JSON(http.StatusOK, lib.Success().WithData(data))
	} else {
		c.JSON(http.StatusOK, lib.NewResponse(res.Code).WithMsg(res.Message).WithData(data))
	}
}

//新增设备
func (dc *DeviceController) CreateDevice(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.CreateDevice{}
	if qErr := c.ShouldBind(&params); qErr != nil {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(qErr.Error()))
		return
	}

	//字段验证
	errs := validator.Builder.Struct(params)
	if vErr := validator.TransError(errs); vErr != "" {
		log.Error("CreateDevice：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}
	fmt.Println("请求 params:", params)

	//远程调用 DeviceService.CreateDevice
	res, err := serviceclient.DeviceServiceClient.CreateDevice(context.TODO(), &proto.ReqCreateDevice{
		Name:     params.Name,
		Sn:       params.Sn,
		ImageUrl: params.ImageUrl,
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
	}

	//响应
	if res.Code == lib.Err.GetInt64("error.ok") {
		c.JSON(http.StatusOK, lib.Success().WithData(data))
	} else {
		c.JSON(http.StatusOK, lib.NewResponse(res.Code))
	}
	return
}

//编辑设备
func (dc *DeviceController) UpdateDevice(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.UpdateDevice{}
	qErr := c.ShouldBind(&params)
	if qErr != nil {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(qErr.Error()))
		return
	}
	params.Id, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	//字段验证
	errs := validator.Builder.Struct(params)
	vErr := validator.TransError(errs)
	if vErr != "" {
		log.Error("UpdateDevice：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}
	fmt.Println("请求 params:", params)

	//远程调用 DeviceService.UpdateDevice
	res, err := serviceclient.DeviceServiceClient.UpdateDevice(context.TODO(), &proto.ReqUpdateDevice{
		Id:       params.Id,
		Name:     params.Name,
		Sn:       params.Sn,
		ImageUrl: params.ImageUrl,
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

//启用/禁用设备
func (dc *DeviceController) DelDevice(c *gin.Context) {
	//获取并绑定请求参数
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	//远程调用 DeviceService.DeleteDevice
	res, err := serviceclient.DeviceServiceClient.DeleteDevice(context.TODO(), &proto.ReqDeleteDevice{
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
