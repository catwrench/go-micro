package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //初始化数据库驱动
	"github.com/micro/go-micro/v2/util/log"
	"meeting-srv/lib"
	"meeting-srv/model"
)

var Db *gorm.DB

func init() {
	log.Info("开始初始化数据库连接")
	var err error
	Db, err = gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			lib.Config.Db.Username,
			lib.Config.Db.Password,
			lib.Config.Db.Addr,
			lib.Config.Db.Port,
			lib.Config.Db.Name,
		))
	if err != nil {
		log.Fatal("连接数据库失败：", err)
	}
	//最大空闲连接
	Db.DB().SetMaxIdleConns(10)
	//最大开启连接
	Db.DB().SetMaxOpenConns(50)

	//初始化表
	fmt.Println("初始化数据库表：model")
	Db.AutoMigrate(
		&model.Device{},
		&model.Participant{},
		&model.Participant{},
		&model.Reservation{},
		&model.Room{},
		&model.RoomDevice{},
		&model.Space{},
	)
	fmt.Println("初始化数据库表结束")
}

func GetDB() *gorm.DB {
	return Db
}
