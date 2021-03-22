package model

import (
	"github.com/golang-module/carbon"
)

type Participant struct {
	Id            int64                   `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	ReservationId int64                   `gorm:"column:reservation_id;type:int(11);comment:'会议预订id'" json:"reservation_id"`
	UserId        int64                   `gorm:"column:user_id;type:int(11);comment:'参与人id'" json:"user_id"`
	UserName      string                  `gorm:"column:username;type:varchar(30);comment:'参与人名称'" json:"user_name"`
	SignAt        carbon.ToDateTimeString `gorm:"column:sign_at;type:datetime(0);comment:'签到时间'" json:"sign_at"`
	CreatedAt     TimeNormal              `gorm:"column:created_at;type:datetime(0)" json:"created_at"`
}

func (Participant) TableName() string {
	return "mt_participants"
}
