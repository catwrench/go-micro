package repository

import (
	"github.com/golang-module/carbon"
	"meeting-srv/database"
	"meeting-srv/model"
)

type ReservationRepository struct {
	base BaseRepo
}

//创建会议预约
func (r *ReservationRepository) CreateReservation(reservation *model.Reservation) error {
	return database.Db.Create(&reservation).Error
}

//删除会议
func (r *ReservationRepository) DelReservation(reservation *model.Reservation) error {
	return database.Db.Delete(reservation).Error
}

//查询会议详情
func (r *ReservationRepository) GetReservation(id int64) (reservation model.Reservation, err error) {
	err = database.Db.Where("id = ?", id).Preload("Participants").Find(&reservation).Error
	return
}

//获取会议列表（分页）
func (r *ReservationRepository) GetReservations(roomId int64, startDate string, endDate string, pageSize int64, page int64, order string, sortBy string) (reservations []*model.Reservation, count int, err error) {
	//查询数据
	query := database.Db.Model(&model.Reservation{}).Where("room_id = ?", roomId)
	if startDate != "" {
		query = query.Where("date >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("date <= ?", endDate)
	}
	err = query.Preload("Participants").Order(order + " " + sortBy).
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&reservations).Error
	if err != nil {
		return
	}

	//计数
	err = query.Count(&count).Error
	return
}

//会议时间重复验证
func (r *ReservationRepository) ValidateTimeRepeat(roomId int64, date string, startTime string, endTime string, id int64) (reservation model.Reservation, err error) {
	/**
	验证思路：共4种情况：1、全包 2、全内 3、左交叉 4、右交叉
	*/
	query := database.Db.Where("room_id = ? and date = ?", roomId, date).
		Where("(start_time <= ? and end_time >= ?)"+
			" or (start_time >= ? and end_time <= ?)"+
			" or (start_time < ? and end_time > ? and end_time < ?)"+
			" or (start_time > ? and start_time < ? and end_time > ?)",
			startTime, endTime,
			startTime, endTime,
			startTime, startTime, endTime,
			startTime, endTime, endTime,
		)

	if id != 0 {
		query = query.Where("id <> ?", id)
	}
	err = query.Find(&reservation).Error
	return
}

//获取会议室预订数量（当前时间之后的）
func (r *ReservationRepository) GetCountOfRoomReservation(roomId int64) (count int, err error) {
	date := carbon.Now().ToDateString()
	time := carbon.Now().ToTimeString()
	err = database.Db.Model(model.Reservation{}).
		Where("room_id = ?", roomId).
		Where("(date > ? ) or (date = ? and start_time > ?)", date, date, time).
		Count(&count).Error
	return
}
