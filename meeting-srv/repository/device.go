package repository

import (
	"meeting-srv/database"
	"meeting-srv/model"
)

type DeviceRepository struct {
}

//创建设备
func (d *DeviceRepository) CreateDevice(device *model.Device) error {
	return database.Db.Create(device).Error
}

//更新设备
func (d *DeviceRepository) UpdateDevice(device *model.Device) error {
	return database.Db.Save(device).Error
}

//删除设备
func (d *DeviceRepository) DeleteDevice(device *model.Device) error {
	//删除设备与会议室关联
	err := database.Db.Where("device_id = ?", device.Id).Delete(model.RoomDevice{}).Error
	if err != nil {
		return err
	}
	//删除设备
	return database.Db.Delete(device).Error
}

//获取设备列表
func (d *DeviceRepository) GetDevices(name string, pageSize int64, page int64, order string, sortBy string) (devices []*model.Device, count int, err error) {
	//查询数据
	query := database.Db.Model(devices)
	if name != "" {
		query = query.Where("name like ?", "%"+name+"%")
	}
	err = query.Order(order + " " + sortBy).
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&devices).Error
	if err != nil {
		return
	}

	//计数
	err = query.Count(&count).Error
	return
}

//获取设备详情
func (d *DeviceRepository) GetDevice(id int64) (device model.Device, err error) {
	err = database.Db.Where("id = ?", id).Find(&device).Error
	return
}

//根据ID批量获取设备
func (d *DeviceRepository) GetDevicesByIds(ids []string) (devices []model.Device, err error) {
	err = database.Db.Find(&devices, ids).Error
	return
}

//根据指定自动获取设备
func (d *DeviceRepository) GetByField(field string, val string) (device model.Device, err error) {
	err = database.Db.Where(field+" = ?", val).Find(&device).Error
	return
}

//验证设备名称是否重复
func (d *DeviceRepository) ValidateNameRepeat(name string, id int64) (device model.Device, err error) {
	query := database.Db.Where("name = ?", name)
	if id != 0 {
		query = query.Where("id <> ?", id)
	}
	err = query.Find(&device).Error
	return
}
