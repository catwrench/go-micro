package wechat

import (
	"fmt"
	"github.com/gowechat/example/config"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

// 公众号操作样例
type OfficialAccount struct {
	Wc              *wechat.Wechat
	OfficialAccount *officialaccount.OfficialAccount
}

//创建微信公众号实例
func NewWechatOfficialAccount(wc *wechat.Wechat) *OfficialAccount {
	//加载配置
	globalCfg := config.GetConfig()
	//创建公众号配置实例
	offCfg := &offConfig.Config{
		AppID:          globalCfg.AppID,
		AppSecret:      globalCfg.AppSecret,
		Token:          globalCfg.Token,
		EncodingAESKey: globalCfg.EncodingAESKey,
	}
	fmt.Println("offCfg:", offCfg)
	//获取公众号实例
	officialAccount := wc.GetOfficialAccount(offCfg)
	return &OfficialAccount{
		Wc:              wc,
		OfficialAccount: officialAccount,
	}
}

//获取AccessToken
func (wa *OfficialAccount) GetAccessToken() (string, error) {
	ak, err := wa.OfficialAccount.GetAccessToken()
	if err != nil {
		log.Error("GetAccessToken error:", err)
		return "", err
	}
	return ak, nil
}

//发送模板消息
func (wa *OfficialAccount) SendTemplate(template *message.TemplateMessage) (msgId int64, err error) {
	return wa.OfficialAccount.GetTemplate().Send(template)
}
