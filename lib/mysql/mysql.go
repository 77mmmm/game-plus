package mysql

import (
	"fmt"
	"game-pro/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/game-pro?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		return nil, fmt.Errorf("mysql init err: %s", err)
	}
	log.Println("mysql init success")
	err = db.DB().Ping()
	if err != nil {
		return nil, fmt.Errorf("mysql ping err: %s", err)
	}
	db.AutoMigrate(&model.LeaderBoard{})
	db.SingularTable(true)
	return db, nil
}
