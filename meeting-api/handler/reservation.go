package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"github.com/micro/go-micro/v2/util/log"
	"meeting-api/lib"
	"meeting-api/serviceclient"
	proto "meeting-api/submodules/common/protob"
	"meeting-api/submodules/common/utils"
	"meeting-api/validator"
	"net/http"
	"strconv"
)

type ReservationController struct {
	BaseController
}

func (rc *ReservationController) GetReservations(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.GetReservations{}
	if qErr := c.ShouldBind(&params); qErr != nil {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(qErr.Error()))
		return
	}
	//字段验证
	errs := validator.Builder.Struct(params)
	if vErr := validator.TransError(errs); vErr != "" {
		log.Error("GetReservations：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}
	//远程调用 ReservationService.GetReservations
	res, err := serviceclient.ReservationServiceClient.GetReservations(context.TODO(), &proto.ReqGetReservations{
		RoomId:    params.RoomId,
		StartDate: params.StartDate,
		EndDate:   params.EndDate,
		Page:      params.Page,
		PageSize:  params.PageSize,
		Order:     params.Order,
		SortBy:    params.SortBy,
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

func (rc *ReservationController) GetReservation(c *gin.Context) {
	//请求参数处理
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg("id不能为空"))
		return
	}

	//远程调用 ReservationService.GetReservation
	res, err := serviceclient.ReservationServiceClient.GetReservation(context.TODO(), &proto.ReqGetReservation{
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

func (rc *ReservationController) CreateReservation(c *gin.Context) {
	//获取并绑定请求参数
	params := validator.CreateReservation{}
	if qErr := c.ShouldBind(&params); qErr != nil {
		log.Info("qErr:", qErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(qErr.Error()))
		return
	}

	//毕竟开始结束时间
	if carbon.Parse("2006-01-02 " + params.StartTime).Gte(carbon.Parse("2006-01-02 " + params.EndTime)) {
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg("会议开始时间不能大于等于结束时间"))
		return
	}

	//字段验证
	errs := validator.Builder.Struct(params)
	if vErr := validator.TransError(errs); vErr != "" {
		log.Error("CreateReservation：", vErr)
		c.JSON(http.StatusOK, lib.NewResponse(1001).WithMsg(vErr))
		return
	}

	//远程调用 ReservationService.CreateReservation
	res, err := serviceclient.ReservationServiceClient.CreateReservation(context.TODO(), &proto.ReqCreateReservation{
		RoomId:          params.RoomId,
		Title:           params.Title,
		Content:         params.Content,
		InitiatorId:     params.InitiatorId,
		ParticipantsIds: params.ParticipantIds,
		Date:            carbon.Parse(params.Date).ToDateString(),
		StartTime:       carbon.Parse("2006-01-02 " + params.StartTime).ToTimeString(),
		EndTime:         carbon.Parse("2006-01-02 " + params.EndTime).ToTimeString(),
		Token:           c.GetHeader("Authorization"),
		CId:             params.CId,
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
