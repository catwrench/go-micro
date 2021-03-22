package handler

import (
	"context"
	"notice-srv/lib"
	"notice-srv/model"
	"notice-srv/repository"
	"notice-srv/services"
	proto "notice-srv/submodules/common/protob"
	"notice-srv/submodules/common/utils"
)

type NoticeService struct {
	repo      repository.NoticeRepository
	noticeSrv services.NoticeBase
}

//获取消息列表
func (this *NoticeService) GetNotices(ctx context.Context, req *proto.ReqGetNotices, res *proto.Response) error {
	//数据查询
	dbRes, count, getErr := this.repo.GetNotices(
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

//创建消息
func (this *NoticeService) CreateNotice(
	c context.Context,
	req *proto.ReqNoticeCreate,
	res *proto.Response,
) error {
	/**
	发送一条消息的流程：
		外部服务：
		eg:发送会议室预订通知
		1、选择微信渠道，获取发送基本信息，（包含open_id,app_id,user_name）
		2、根据微信模板id组装发送内容
		3、rpc调用消息服务
		消息服务：
		4、验证请求参数
		5、生成消息记录
		6、（可扩展是否加载配置）
		7、根据发送渠道，执行发送（每个渠道需要实现发送接口）
		6、记录发送结果数据，并返回

		结构：
		消息对象模型 struct
			参数验证 params
			验证器 validator
			接收对象 user
			消息实体 notice
			消息子实体 noticeUser
			订阅配置 noticeSubscription
			日志记录 log
	*/

	//加载模板模板数据，将请求参数中的变量加载进模板，完成参数绑定
	template := &services.BindMeetingParticipants{}
	err := utils.UnmarshalProtoAnyTo(req.Template, template)
	if err != nil {
		lib.NewResponse(1000).WithMsg("解析失败")
		return nil
	}

	//组装模型
	notice := &model.Notice{}
	var noticeUsers []*model.NoticeUser
	noticeUsers = append(noticeUsers, &model.NoticeUser{
		UserId:   req.UserId,
		UserType: req.UserType,
	})

	//创建消息推送
	resp := this.noticeSrv.Create(notice, noticeUsers, template)
	resp.ToProtoRes(res)
	return nil

	////组装请求参数对象
	//repoSrv := &services.RepoNoticeCreate{}
	//createReq := &services.Request{
	//	UserId:   req.UserId,
	//	UserName: req.UserName,
	//	OpenId:   req.OpenId,
	//	Content:  req.Content,
	//	Template: req.Template,
	//}
	////初始化通知并开始创建
	//errCreate := repoSrv.Init(createReq).
	//	RegisterBehaviour(
	//		&services.ParamsBehaviour{},
	//		&services.NoticeBehaviour{},
	//		&services.NoticeUserBehaviour{},
	//		&services.SendBehaviour{},
	//).
	//	Run()
	//if errCreate != nil {
	//	log.Error("创建消息提醒 error:", errCreate)
	//	lib.NewResponse(1000).WithMsg("创建消息通知失败，请稍后重试")
	//}
	//lib.Success()
}
