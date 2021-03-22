package wechat

import (
	"fmt"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"notice-srv/lib"
)

// 小程序
type MiniProgram struct {
	Wc          *wechat.Wechat
	MiniProgram *miniprogram.MiniProgram
}

//创建小程序实例
func NewWechatMiniProgram(wc *wechat.Wechat) *MiniProgram {
	//创建小程序配置实例
	miniCfg := &miniConfig.Config{
		AppID:     lib.Config.MiniProgram.AppID,
		AppSecret: lib.Config.MiniProgram.AppSecret,
	}
	fmt.Println("miniCfg:", miniCfg)
	//获取小程序实例
	miniProgram := wc.GetMiniProgram(miniCfg)
	return &MiniProgram{
		Wc:          wc,
		MiniProgram: miniProgram,
	}
}
