package model

type Notice struct {
	Id         int64      `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	Content    string     `gorm:"column:content;type:text(0);comment:'消息内容'" json:"content"`
	Channel    string     `gorm:"column:channel;type:varchar(20);comment:'消息渠道:微信通知、短信、邮件、站内信'" json:"channel"`
	Action     string     `gorm:"column:action;type:varchar(30);comment:'消息动作：如点赞、预订、收藏'" json:"action"`
	TargetId   int64      `gorm:"column:target_id;type:int(11);comment:'目标id：如文章id、订单id'" json:"target_id"`
	TargetType string     `gorm:"column:target_type;type:varchar(30);comment:'目标类型:如文章、订单'" json:"target_type"`
	CreatedAt  TimeNormal `gorm:"column:created_at;type:datetime(0)" json:"created_at"`
}

func (Notice) TableName() string {
	return "nt_notice"
}
