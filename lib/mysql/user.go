package mysql

import (
	"fmt"
	"game-pro/model"
	"gorm.io/gorm"
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

	err := r.db.Select("uuid,score,score_update_time").Table("leaderboard").Where("score>?", 400000).Order("score desc,score_update_time asc").Limit(end - start).Offset(start).Find(&user).Error
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
	var count1 int64
	var count2 int64
	subQuery1 := r.db.Table("leaderboard").Where("uuid=?", username)
	err := r.db.Table("leaderboard as l,(?) as t", subQuery1).Where("l.score>t.score").Order("l.score desc").Count(&count1).Error
	if err != nil {
		return -1, fmt.Errorf("userrank not found,err:%s", err)
	}
	err = r.db.Table("leaderboard as l,(?) as t", subQuery1).Where("l.score=t.score and l.score_update_time<t.score_update_time").Count(&count2).Error
	return count1 + count2 + 1, nil
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
		Uuid:  username,
		Score: userscore,
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
	err := tx.Model(&model.LeaderBoard{}).Set("gorm:query_option", "LOCK IN SHARE MODE").Where("uuid = ?", username).Limit(1).Find(&list).Error
	if err != nil {
		tx.Rollback()
	}
	if len(list) == 0 {
		tx.Rollback()
		return fmt.Errorf("user not found,err:%s", err)
	}
	user := list[0]
	user.Score = userscore
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
	err := tx.Where("uuid=?", username).Set("gorm:query_option", "LOCK IN SHARE MODE").Delete(&list).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("user cannot delete,err:%s", err)
	}
	tx.Commit()
	return nil

}
