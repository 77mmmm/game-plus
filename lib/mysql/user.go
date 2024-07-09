package mysql

import (
	"fmt"
	"game-pro/model"
	"github.com/jinzhu/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	db, err := InitDb()
	if err != nil {
		panic(err)
	}
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetLeaderBoard(start, end int) ([]model.LeaderBoard, error) {
	var user []model.LeaderBoard
	err := r.db.Table("leaderboard").Select("uuid").Order("score desc,scoretime asc").Limit(end - start).Offset(start).Scan(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("leaderboard fail to get the data,err:%s", err)

	}
	return user, nil
}

func (r *UserRepository) GetUserScore(username int64) (int64, error) {
	var user model.LeaderBoard
	var score int64
	err := r.db.Table("leaderboard").Where("uuid = ?", username).First(&user).Error
	if err != nil {
		return -1, fmt.Errorf("user not found,err:%s", err)
	}
	score = user.Score
	return score, nil

}

func (r *UserRepository) GetUserRank(username int64) (int64, error) {
	var user model.LeaderBoard
	err := r.db.Table("leaderboard").Select("score").Where("uuid = ?", username).First(&user).Error
	if err != nil {
		return -1, fmt.Errorf("userrank not found,err:%s", err)
	}
	score := user.Score
	var count int64
	err = r.db.Table("leaderboard").Where("score>?", score).Order("scoretime asc").Count(&count).Error
	return count + 1, nil
}

func (r *UserRepository) AddUser(username int64, userscore int64) error {
	tx := r.db.Begin()
	var list []model.LeaderBoard
	err := tx.Model(&model.LeaderBoard{}).Set("gorm:query_option", "LOCK IN SHARE MODE").Where("uuid = ?", username).Limit(1).Find(&list).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("query cannot add,err:%s", err)
	}
	if len(list) != 0 {
		tx.Rollback()
		return fmt.Errorf("user exist,err:%s", err)
	}
	err = tx.Model(&model.LeaderBoard{}).Create(model.LeaderBoard{
		Uuid:      username,
		Score:     userscore,
		Scoretime: time.Now(),
	}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("user cannot add,err:%s", err)
	}
	tx.Commit()
	return nil
}

func (r *UserRepository) UpdateUser(username int64, userscore int64) error {
	tx := r.db.Begin()
	var list []model.LeaderBoard
	err := tx.Model(&model.LeaderBoard{}).Where("uuid = ?", username).Limit(1).Find(&list).Error
	if err != nil {
		tx.Rollback()
	}
	if len(list) == 0 {
		tx.Rollback()
		return fmt.Errorf("user not found,err:%s", err)
	}
	user := list[0]
	user.Score = userscore
	user.Scoretime = time.Now()
	err1 := tx.Model(&model.LeaderBoard{}).Save(user).Error
	if err1 != nil {
		tx.Rollback()
		return fmt.Errorf("user cannot update,err:%s", err)
	}
	tx.Commit()
	return nil

}

func (r *UserRepository) DeleteUser(username int64) error {
	tx := r.db.Begin()
	var list []model.LeaderBoard
	err := tx.Where("uuid=?", username).Delete(&list).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("user cannot delete,err:%s", err)
	}
	tx.Commit()
	return nil

}
