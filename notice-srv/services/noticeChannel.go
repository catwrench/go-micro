package services

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	libWechat "notice-srv/lib/wechat"
	"notice-srv/model"
	"notice-srv/repository"
)

//消息发送渠道接口
type NoticeChannel interface {
	Send(notice *model.Notice) (errorCode int64)
}

func NewWechatChannel() NoticeChannel {
	return &wechatChannel{}
}

//--------------------------------------
//微信模板消息渠道实现
type wechatChannel struct {
	repo repository.NoticeUserRepository
}

func (wc *wechatChannel) Send(notice *model.Notice) (errorCode int64) {
	//初始化微信公众号实例
	offAccount := libWechat.NewWechatOfficialAccount(libWechat.InitWechat())

	//获取accessToken
	ak, err := offAccount.GetAccessToken()
	if err != nil {
		log.Info("获取微信accessToken error:", ak, err)
		return 20201
	}

	//获取待发送的人
	noticeUser, err := wc.repo.GetOneByNoticeId(notice.Id)
	if err != nil {
		log.Info("GetManyByNoticeId error:", noticeUser, err)
		return 20201
	}

	//发送消息
	var msgId int64
	msgId, err = offAccount.SendTemplate(&message.TemplateMessage{
		//ToUser:     noticeUser.UserId,
		TemplateID: "h1-e7A4a5gR0gzn1VF42RnpKtFEm5oZ4S2Nx2f6BM9s",
		URL:        "",
		Data:       nil,
		MiniProgram: struct {
			AppID    string `json:"appid"`
			PagePath string `json:"pagepath"`
		}{
			AppID:    "",
			PagePath: "",
		},
	})
	if err != nil {
		log.Info("offAccount.SendTemplate error:", msgId, err)
		return 20201
	}

	return
}
