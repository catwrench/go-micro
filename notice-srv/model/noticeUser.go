package model

import "github.com/golang-module/carbon"

type NoticeUser struct {
	Id        int64                   `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	NotifyId  int64                   `gorm:"column:notify_id;type:int(11);comment:'消息id'" json:"notify_id"`
	UserId    int64                   `gorm:"column:user_id;type:int(11);comment:'接收人id'" json:"user_id"`
	UserType  string                  `gorm:"column:user_type;type:varchar(30);comment:'接收人类型：前台用户、后台用户'" json:"user_type"`
	ReadTime  carbon.ToDateTimeString `gorm:"column:read_time;type:datetime(0);comment:'阅读时间'" json:"read_time"`
	CreatedAt TimeNormal              `gorm:"column:created_at;type:datetime(0)" json:"created_at"`
}

func (NoticeUser) TableName() string {
	return "nt_notice_user"
}
