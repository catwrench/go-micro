package model

type Space struct {
	Id        int64      `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	Name      string     `gorm:"column:name;type:varchar(30);comment:'地点名称';unique_index" json:"name"`
	Lng       float64    `gorm:"column:lng;type:decimal(13,6);comment:'经度'" json:"lng"`
	Lat       float64    `gorm:"column:lat;type:decimal(13,6);comment:'纬度'" json:"lat"`
	Status    string     `gorm:"column:status;type:tinyint(4);comment:'启用状态：0禁用、1启用'" json:"status"`
	CreatedAt TimeNormal `gorm:"column:created_at;type:datetime(0)" json:"created_at"`
}

//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为spaces（结构体+s）
func (Space) TableName() string {
	return "mt_space"
}
