package repository

import (
	"meeting-srv/database"
	"meeting-srv/model"
)

type RoomRepository struct {
}

//创建会议室
func (s *RoomRepository) CreateRoom(room *model.Room) error {
	return database.Db.Create(room).Error
}

//更新会议室
func (s *RoomRepository) UpdateRoom(room *model.Room) error {
	//先替换关联
	err := database.Db.Model(&room).Association("Devices").Replace(room.Devices).Error
	if err != nil {
		return err
	}
	//再更新实体，（这个时候room对象的关联已经被删除了）
	return database.Db.Save(&room).Error
}

//更新会议室状态
func (s *RoomRepository) UpdateRoomStatus(room *model.Room) error {
	data := map[string]interface{}{
		"status": room.Status,
	}
	return database.Db.Model(&model.Room{}).Update(data).Error
}

//删除会议室
func (s *RoomRepository) DelRoom(room *model.Room) error {
	return database.Db.Delete(room).Error
}

//查询会议室详情
func (s *RoomRepository) GetRoom(id int64) (room model.Room, err error) {
	err = database.Db.Where("id = ?", id).Preload("Devices").Find(&room).Error
	return
}

//查询会议室详情通过指定字段
func (s *RoomRepository) GetByField(field string, val string) (room model.Room, err error) {
	err = database.Db.Where(field+" = ?", val).Find(&room).Error
	return
}

//获取会议室列表（分页）
func (s *RoomRepository) GetRooms(spaceId int64, status string, pageSize int64, page int64, order string, sortBy string) (rooms []*model.Room, count int, err error) {
	//查询数据
	query := database.Db.Model(rooms)
	if spaceId != 0 {
		query = query.Where("space_id = ?", spaceId)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err = query.Preload("Devices").
		Order(order + " " + sortBy).
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&rooms).Error
	if err != nil {
		return
	}

	//计数
	countQuery := query
	if spaceId != 0 {
		countQuery = countQuery.Where("space_id = ?", spaceId)
	}
	if status != "" {
		countQuery = countQuery.Where("status = ?", status)
	}
	err = countQuery.Count(&count).Error
	return
}

//会议室名称重复验证
func (s *RoomRepository) ValidateNameRepeat(name string, spaceId int64, id int64) (room model.Room, err error) {
	query := database.Db.Where("name = ? and space_id = ?", name, spaceId)
	if id != 0 {
		query = query.Where("id <> ?", id)
	}
	err = query.Find(&room).Error
	return
}

//根据地点获取会议室列表
func (s *RoomRepository) GetRoomsBySpaceId(spaceId int64) (rooms []*model.Room, err error) {
	//查询数据
	err = database.Db.Model(rooms).Where("space_id = ?", spaceId).Find(&rooms).Error
	return
}
