package repository

import (
	"meeting-srv/database"
	"meeting-srv/model"
)

type SpaceRepository struct {
}

//创建地点
func (s *SpaceRepository) CreateSpace(space *model.Space) error {
	return database.Db.Create(space).Error
}

//更新地点
func (s *SpaceRepository) UpdateSpace(space *model.Space) error {
	return database.Db.Save(space).Error
}

//更新地点状态
func (s *SpaceRepository) UpdateSpaceStatus(space *model.Space) error {
	data := map[string]interface{}{
		"status": space.Status,
	}
	return database.Db.Model(&model.Space{}).Update(data).Error
}

//删除地点
func (s *SpaceRepository) DelSpace(space *model.Space) error {
	return database.Db.Delete(space).Error
}

//查询地点详情
func (s *SpaceRepository) GetSpace(id int64) (space model.Space, err error) {
	err = database.Db.Where("id = ?", id).Find(&space).Error
	return
}

//查询地点详情通过指定字段
func (s *SpaceRepository) GetByField(field string, val string) (space model.Space, err error) {
	err = database.Db.Where(field+" = ?", val).Find(&space).Error
	return
}

//获取地点列表（分页）
func (s *SpaceRepository) GetSpaces(status string, pageSize int64, page int64, order string, sortBy string) (spaces []*model.Space, count int, err error) {
	//查询数据
	query := database.Db.Model(spaces)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err = query.Order(order + " " + sortBy).
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&spaces).Error
	if err != nil {
		return
	}

	//计数
	err = query.Count(&count).Error
	return
}

//获取所有地点
func (s *SpaceRepository) GetAllSpaces(order string, sortBy string) (spaces []*model.Space, err error) {
	//查询数据
	err = database.Db.Model(&model.Space{}).Order(order + " " + sortBy).Find(&spaces).Error
	if err != nil {
		return
	}
	return
}

//地点名称重复验证
func (s *SpaceRepository) ValidateNameRepeat(name string, id int64) (space model.Space, err error) {
	query := database.Db.Where("name = ?", name)
	if id != 0 {
		query = query.Where("id <> ?", id)
	}
	err = query.Find(&space).Error
	return
}
