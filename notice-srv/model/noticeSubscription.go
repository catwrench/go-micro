package model

type NoticeSubscription struct {
	Id        int64       `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	UserId    int64       `gorm:"column:user_id;type:int(11);comment:'接收人id'" json:"user_id"`
	UserType  string      `gorm:"column:user_type;type:varchar(30);comment:'接收人类型：前台用户、后台用户'" json:"user_type"`
	Config    interface{} `gorm:"column:config;type:json(0);comment:'配置类型：string配置类型:bool是否启用'" json:"config"`
	CreatedAt TimeNormal  `gorm:"column:created_at;type:datetime(0)" json:"created_at"`
}

func (NoticeSubscription) TableName() string {
	return "nt_notice_subscription"
}
