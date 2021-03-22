package model

type RoomDevice struct {
	DeviceId  int64      `gorm:"column:device_id;type:int(11);comment:'设备id'" json:"device_id"`
	RoomId    int64      `gorm:"column:room_id;type:int(11);comment:'会议室id'" json:"room_id"`
	CreatedAt TimeNormal `gorm:"column:created_at;type:datetime(0)" json:"created_at"`
}

func (RoomDevice) TableName() string {
	return "mt_room_device"
}
