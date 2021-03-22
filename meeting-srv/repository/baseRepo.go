package repository

import (
	"github.com/jinzhu/gorm"
	"meeting-srv/database"
)

type BaseRepo struct {
	model gorm.Model
}

//创建
func (r *BaseRepo) CreateModel(model interface{}) error {
	return database.Db.Create(model).Error
}

//删除
func (r *BaseRepo) DelModel(model interface{}) error {
	return database.Db.Delete(model).Error
}

//查询详情
func (r *BaseRepo) GetModel(model interface{}, id int64) (interface{}, error) {
	err := database.Db.Where("id = ?", id).Find(&model).Error
	return model, err
}

//获取列表（分页）
func (r *BaseRepo) GetModels(model interface{}, pageSize int64, page int64, order string, sortBy string) (models []interface{}, count int, err error) {
	//查询数据
	err = database.Db.Model(&model).
		Order(order + " " + sortBy).
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&models).Error
	if err != nil {
		return
	}

	//计数
	err = database.Db.Model(&model).Count(&count).Error
	return
}
