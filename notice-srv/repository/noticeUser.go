package repository

import (
	"notice-srv/database"
	"notice-srv/model"
)

type NoticeUserRepository struct {
}

//创建消息接收人
func (d *NoticeUserRepository) Create(model *model.NoticeUser) error {
	return database.Db.Create(&model).Error
}

//批量创建消息接收人
func (d *NoticeUserRepository) CreateMany(models []*model.NoticeUser) (err error) {
	for _, m := range models {
		err = database.Db.Create(&m).Error
		if err != nil {
			return
		}
	}
	return
}

//更新消息接收人
func (d *NoticeUserRepository) Update(model *model.NoticeUser) error {
	return database.Db.Save(&model).Error
}

//获取消息接收人列表
func (d *NoticeUserRepository) GetMany(name string, pageSize int64, page int64, order string, sortBy string) (models []*model.NoticeUser, count int, err error) {
	//查询数据
	query := database.Db.Model(models)
	if name != "" {
		query = query.Where("name like ?", "%"+name+"%")
	}
	err = query.Order(order + " " + sortBy).
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&models).Error
	if err != nil {
		return
	}

	//计数
	err = query.Count(&count).Error
	return
}

//获取消息接收人详情
func (d *NoticeUserRepository) GetOne(id int64) (model model.NoticeUser, err error) {
	err = database.Db.Where("id = ?", id).Find(&model).Error
	return
}

//根据消息id批量获取被提醒人
func (d *NoticeUserRepository) GetManyByNoticeId(noticeId int64) (model []*model.NoticeUser, err error) {
	err = database.Db.Where("notice_id = ?", noticeId).Find(&model).Error
	return
}

//根据消息id获取被提醒人
func (d *NoticeUserRepository) GetOneByNoticeId(noticeId int64) (model model.NoticeUser, err error) {
	err = database.Db.Where("noticeId = ?", noticeId).Find(&model).Error
	return
}
