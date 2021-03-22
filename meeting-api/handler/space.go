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

type SpaceController struct {
	BaseController
}

//查询所有地点列表
func (sc *SpaceController) GetAllSpaces(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.GetAllSpaces{}
	qErr := c.ShouldBind(&params)
	if qErr != nil {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(qErr.Error()))
		return
	}
	//字段验证
	errs := validator.Builder.Struct(params)
	vErr := validator.TransError(errs)
	if vErr != "" {
		log.Error("GetAllSpaces：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}

	//添加链路追踪子span
	ctx, ok := lib.ContextWithSpan(c)
	if ok == false {
		log.Error("get spanContext err")
	}
	//远程调用 SpaceService.GetAllSpaces
	res, err := serviceclient.SpaceServiceClient.GetAllSpaces(ctx, &proto.ReqGetAllSpaces{
		Order:  params.Order,
		SortBy: params.SortBy,
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

//查询地点列表
func (sc *SpaceController) GetSpaces(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.GetSpaces{}
	qErr := c.ShouldBind(&params)
	if qErr != nil {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(qErr.Error()))
		return
	}
	//字段验证
	errs := validator.Builder.Struct(params)
	vErr := validator.TransError(errs)
	if vErr != "" {
		log.Error("GetSpaces：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}
	fmt.Println("请求 params:", params)

	//远程调用 SpaceService.GetSpaces
	res, err := serviceclient.SpaceServiceClient.GetSpaces(context.TODO(), &proto.ReqGetSpaces{
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
	fmt.Println("SpaceServiceClient.GetSpaces res:", res)

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

//查询地点详情
func (sc *SpaceController) GetSpace(c *gin.Context) {
	//请求参数处理
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg("id不能为空"))
		return
	}

	//远程调用 SpaceService.GetSpaces
	res, err := serviceclient.SpaceServiceClient.GetSpace(context.TODO(), &proto.ReqGetSpace{
		Id: id,
	})
	if err != nil {
		log.Info("未知错误：", err)
		c.JSON(http.StatusOK, lib.NewResponse(1000))
		return
	}
	fmt.Println("SpaceServiceClient.GetSpace res:", res)

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

//新增地点
func (sc *SpaceController) CreateSpace(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.CreateSpace{}
	qErr := c.ShouldBind(&params)
	if qErr != nil {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(qErr.Error()))
		return
	}
	//字段验证
	errs := validator.Builder.Struct(params)
	vErr := validator.TransError(errs)
	if vErr != "" {
		log.Error("CreateSpace：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}
	fmt.Println("请求 params:", params)

	//远程调用 SpaceService.CreateSpace
	res, err := serviceclient.SpaceServiceClient.CreateSpace(context.TODO(), &proto.ReqCreateSpace{
		Name:      params.Name,
		Lng:       params.Lng,
		Lat:       params.Lat,
		OneStatus: &proto.ReqCreateSpace_Status{Status: params.Status},
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

//编辑地点
func (sc *SpaceController) UpdateSpace(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.UpdateSpace{}
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
		log.Error("UpdateSpace：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}
	fmt.Println("请求 params:", params)

	//远程调用 SpaceService.UpdateSpace
	res, err := serviceclient.SpaceServiceClient.UpdateSpace(context.TODO(), &proto.ReqUpdateSpace{
		Id:        params.Id,
		Name:      params.Name,
		Lng:       params.Lng,
		Lat:       params.Lat,
		OneStatus: &proto.ReqUpdateSpace_Status{Status: params.Status},
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

//启用/禁用地点
func (sc *SpaceController) DelSpace(c *gin.Context) {
	//获取并绑定请求参数
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	//远程调用 SpaceService.DelSpace
	res, err := serviceclient.SpaceServiceClient.DelSpace(context.TODO(), &proto.ReqDelSpace{
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

//启用/禁用地点
func (sc *SpaceController) UpdateSpaceStatus(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.UpdateSpaceStatus{}
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
		log.Error("UpdateSpaceStatus：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}
	fmt.Println("请求 params:", params)

	//远程调用 SpaceService.UpdateSpaceStatus
	res, err := serviceclient.SpaceServiceClient.UpdateSpaceStatus(context.TODO(), &proto.ReqUpdateSpaceStatus{
		Id:        params.Id,
		OneStatus: &proto.ReqUpdateSpaceStatus_Status{Status: params.Status},
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
