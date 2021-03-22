package repository

import (
	"notice-srv/database"
	"notice-srv/model"
)

type NoticeRepository struct {
}

//创建消息
func (d *NoticeRepository) CreateNotice(notice *model.Notice) error {
	return database.Db.Create(&notice).Error
}

//更新消息
func (d *NoticeRepository) UpdateNotice(notice *model.Notice) error {
	return database.Db.Save(&notice).Error
}

//获取消息列表
func (d *NoticeRepository) GetNotices(name string, pageSize int64, page int64, order string, sortBy string) (notices []*model.Notice, count int, err error) {
	//查询数据
	query := database.Db.Model(notices)
	if name != "" {
		query = query.Where("name like ?", "%"+name+"%")
	}
	err = query.Order(order + " " + sortBy).
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&notices).Error
	if err != nil {
		return
	}

	//计数
	err = query.Count(&count).Error
	return
}

//获取消息详情
func (d *NoticeRepository) GetNotice(id int64) (notice model.Notice, err error) {
	err = database.Db.Where("id = ?", id).Find(&notice).Error
	return
}

//根据ID批量获取消息
func (d *NoticeRepository) GetNoticesByIds(ids []string) (notices []model.Notice, err error) {
	err = database.Db.Find(&notices, ids).Error
	return
}

//根据指定字段获取消息
func (d *NoticeRepository) GetByField(field string, val string) (notice model.Notice, err error) {
	err = database.Db.Where(field+" = ?", val).Find(&notice).Error
	return
}
