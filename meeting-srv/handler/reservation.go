package handler

import (
	"context"
	"fmt"
	"github.com/golang-module/carbon"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
	"meeting-srv/lib"
	"meeting-srv/model"
	"meeting-srv/repository"
	"meeting-srv/serviceclient"
	proto "meeting-srv/submodules/common/protob"
	"meeting-srv/submodules/common/utils"
	"strconv"
)

type ReservationService struct {
	repo repository.ReservationRepository
}

//查询会议列表
func (r ReservationService) GetReservations(ctx context.Context, req *proto.ReqGetReservations, res *proto.Response) error {
	//数据查询
	dbRes, count, getErr := r.repo.GetReservations(
		req.RoomId,
		req.StartDate,
		req.EndDate,
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
	lib.Success().WithData(dbRes).ToProtoRes(res)
	return nil
}

//查询会议详情
func (r ReservationService) GetReservation(ctx context.Context, req *proto.ReqGetReservation, res *proto.Response) error {
	dbRes, getErr := r.repo.GetReservation(req.Id)
	if getErr != nil {
		if getErr == gorm.ErrRecordNotFound {
			lib.NewResponse(20111).ToProtoRes(res)
			return nil
		}
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}

	lib.Success().WithData(dbRes).ToProtoRes(res)
	return nil
}

//新增会议
func (r ReservationService) CreateReservation(ctx context.Context, req *proto.ReqCreateReservation, res *proto.Response) error {
	//会议时间重复验证
	errCode := r.validateTimeRepeat(req.RoomId, req.Date, req.StartTime, req.EndTime, 0)
	if errCode != lib.Err.GetInt64("error.ok") {
		lib.NewResponse(errCode).ToProtoRes(res)
		return nil
	}

	//获取会议参与人(将发起人也作为参与人一起查询)
	uRes, err := serviceclient.UserServiceClient.BaseInfoEmployeeListNoPage(context.TODO(), &proto.ReqBaseInfoEmployeeListNoPage{
		Token:    req.Token,
		Ids:      req.ParticipantsIds + "," + strconv.FormatInt(req.InitiatorId, 10),
		RealName: "",
		Mobile:   "",
		Email:    "",
		CId:      strconv.FormatInt(req.CId, 10),
	})
	if err != nil {
		log.Info("未知错误：", err)
		lib.NewResponse(30100).ToProtoRes(res)
		return nil
	}

	//需要获取 proto.any格式中的data,需要进行解码
	//解码any格式
	type userRes struct {
		SuId     int64  `json:"su_id"`
		Email    string `json:"email"`
		Mobile   string `json:"mobile"`
		RealName string `json:"real_name"`
		Sex      int64  `json:"sex"`
	}
	var data []userRes
	paErr := utils.UnmarshalProtoAnyTo(uRes.Data, &data)
	if paErr != nil {
		log.Error("UnmarshalProtoAny err:", data)
		lib.NewResponse(1000)
		return nil
	}
	fmt.Println("BaseInfoEmployeeListNoPage data:", data)

	//转换用户中心结果为 model.Participant
	var participants []model.Participant
	for _, v := range data {
		participants = append(participants, model.Participant{
			UserId:   v.SuId,
			UserName: v.RealName,
		})
	}

	create := &model.Reservation{
		RoomId:        req.RoomId,
		Title:         req.Title,
		Content:       req.Content,
		InitiatorId:   req.InitiatorId,
		InitiatorName: "",
		Date:          carbon.ToDateString{carbon.ParseByFormat(req.Date, "Y-m-d")},
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		Participants:  participants,
	}
	//验证发起人
	for _, v := range participants {
		if v.UserId == req.InitiatorId {
			create.InitiatorName = v.UserName
			break
		}
	}
	if create.InitiatorName == "" {
		lib.NewResponse(20114).ToProtoRes(res)
		return nil
	}

	//创建会议及参与人
	createErr := r.repo.CreateReservation(create)
	if createErr != nil {
		log.Error("repo.CreateReservation error:", createErr)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	lib.Success().ToProtoRes(res)
	return nil
}

//会议时间重复验证
func (r ReservationService) validateTimeRepeat(roomId int64, date string, startTime string, endTime string, id int64) (code int64) {
	reservation, err := r.repo.ValidateTimeRepeat(roomId, date, startTime, endTime, id)
	if err != nil {
		log.Error("ValidateTimeRepeat error:", err)
		if err != gorm.ErrRecordNotFound {
			code = 1001
			return
		}
	}
	if reservation.Id != 0 {
		code = 20110
		return
	}
	code = lib.Err.GetInt64("error.ok")
	return
}
