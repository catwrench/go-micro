package model

type Device struct {
	Id        int64      `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	Name      string     `gorm:"column:name;type:varchar(30);comment:'设备名称'" json:"name"`
	Sn        string     `gorm:"column:sn;type:varchar(30);comment:'设备编号'" json:"sn"`
	ImageUrl  string     `gorm:"column:image_url;type:varchar(200);comment:'设备图'" json:"image_url"`
	CreatedAt TimeNormal `gorm:"column:created_at;type:datetime(0)" json:"created_at"`
}

func (Device) TableName() string {
	return "mt_device"
}
