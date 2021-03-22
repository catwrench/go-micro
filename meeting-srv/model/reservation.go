package model

import (
	"github.com/golang-module/carbon"
)

type Reservation struct {
	Id            int64               `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	RoomId        int64               `gorm:"column:room_id;type:int(11);comment:'会议室id'" json:"room_id"`
	Title         string              `gorm:"column:title;type:varchar(30);comment:'会议主题'" json:"title"`
	Content       string              `gorm:"column:content;type:varchar(500);comment:'会议内容'" json:"content"`
	InitiatorId   int64               `gorm:"column:initiator_id;type:int(11);comment:'发起人id'" json:"initiator_id"`
	InitiatorName string              `gorm:"column:initiator_name;type:varchar(30);comment:'发起人'" json:"initiator_name"`
	Date          carbon.ToDateString `gorm:"column:date;type:date;comment:'会议日期'" json:"date"`
	StartTime     string              `gorm:"column:start_time;type:time;comment:'会议开始时间'" json:"start_time"`
	EndTime       string              `gorm:"column:end_time;type:time;comment:'会议结束时间'" json:"end_time"`
	CreatedAt     TimeNormal          `gorm:"column:created_at;type:datetime(0)" json:"created_at"`
	//ForeignKey定义外键，AssociationForeignKey定义关联表的id
	Participants []Participant `gorm:"ForeignKey:ReservationId;AssociationForeignKey:Id" json:"participants"`
}

func (Reservation) TableName() string {
	return "mt_reservation"
}
