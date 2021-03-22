package model

type Room struct {
	Id          int64      `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	SpaceId     int64      `gorm:"column:space_id;type:int(11);comment:'所属地点ID'" json:"space_id"`
	Name        string     `gorm:"column:name;type:varchar(30);comment:'会议室名称'" json:"name"`
	Status      string     `gorm:"column:status;type:tinyint(4);comment:'启用状态：0禁用、1启用'" json:"status"`
	ImageUrl    string     `gorm:"column:image_url;type:varchar(200);comment:'会议室图片'" json:"image_url"`
	CapacityMin int64      `gorm:"column:capacity_min;type:int(11);comment:'建议使用人数（最小）'" json:"capacity_min"`
	CapacityMax int64      `gorm:"column:capacity_max;type:int(11);comment:'建议使用人数（最大）'" json:"capacity_max"`
	CreatedAt   TimeNormal `gorm:"column:created_at;type:datetime(0)" json:"created_at"`
	Devices     []Device   `gorm:"many2many:mt_room_device;association_autoupdate:false;association_autocreate:false" json:"devices"`
}

func (Room) TableName() string {
	return "mt_room"
}
