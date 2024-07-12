package mysql

import (
	"fmt"
	"log"

	"game-pro/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	dsn := "root:52Tiananmen.@tcp(z1.juhong.live:3306)/game-pro?charset=utf8&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("mysql init err: %s", err)
	}
	log.Println("mysql init success")
	_ = db.AutoMigrate(&model.LeaderBoard{})
	return db, nil
}
