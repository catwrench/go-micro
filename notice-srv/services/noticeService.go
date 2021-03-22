package services

import (
	"notice-srv/lib"
	"notice-srv/model"
	"notice-srv/repository"
)

type NoticeBase interface {
	Create(*model.Notice, []*model.NoticeUser, *BindMeetingParticipants) *lib.Response
	List()
	Detail()
	Delete()
	Read()
}

func NewNoticeBase() NoticeBase {
	return &noticeBase{}
}

type noticeBase struct {
	repo           repository.NoticeRepository
	noticeUserRepo repository.NoticeUserRepository
}

func (n noticeBase) Create(notice *model.Notice, noticeUsers []*model.NoticeUser, participants *BindMeetingParticipants) *lib.Response {
	if notice.Action != "meeting-participant" {
		return lib.NewResponse(1000).WithMsg("消息动作错误")
	}
	if notice.Channel != "wechat" {
		return lib.NewResponse(1000).WithMsg("发送渠道错误")
	}
	//记录消息通知
	err := n.repo.CreateNotice(notice)
	if err != nil {
		return lib.NewResponse(1000).WithMsg("创建消息失败")
	}

	//记录要通知的人
	for _, noticeUser := range noticeUsers {
		noticeUser.NotifyId = notice.Id
	}
	err = n.noticeUserRepo.CreateMany(noticeUsers)
	if err != nil {
		return lib.NewResponse(1000).WithMsg("创建消息通知的人")
	}

	//调起渠道发送记录

	return lib.Success()
}

func (n noticeBase) List() {
	panic("implement me")
}

func (n noticeBase) Detail() {
	panic("implement me")
}

func (n noticeBase) Delete() {
	panic("implement me")
}

func (n noticeBase) Read() {
	panic("implement me")
}
