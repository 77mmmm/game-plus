package dao

import (
	"fmt"
	"game-pro/lib/mysql"
	"game-pro/model"
)

type MysqlRepository struct {
	db *mysql.UserRepository
}

func NewMysqlRepository() *MysqlRepository {
	return &MysqlRepository{
		db: mysql.NewUserRepository(),
	}
}

func (m *MysqlRepository) GetUserRank(username int64) (int64, error) {
	rank, err := m.db.GetUserRank(username)
	if err != nil {
		return -1, fmt.Errorf("Query Ploblem: %v", err)
	}
	return rank, nil
}
func (m *MysqlRepository) GetUserScore(username int64) (int64, error) {
	score, err := m.db.GetUserScore(username)
	if err != nil {
		return -1, fmt.Errorf("Query Ploblem: %v", err)
	}
	return score, nil
}

func (m *MysqlRepository) UpdateUserScore(username int64, score int64) error {
	err := m.db.UpdateUser(username, score)
	if err != nil {
		return fmt.Errorf("Query Ploblem: %v", err)
	}
	return nil

}

func (m *MysqlRepository) GetLimitLeaderBoard(start, end int) ([]model.LeaderBoard, error) {
	leaderboard, err := m.db.GetLeaderBoard(start, end)
	if err != nil {
		return nil, fmt.Errorf("Query Ploblem: %v", err)
	}
	return leaderboard, nil

}

func (m *MysqlRepository) AddUser(username int64, score int64) error {
	err := m.db.AddUser(username, score)
	if err != nil {
		return fmt.Errorf("Add Ploblem: %v", err)
	}
	return nil
}

func (m *MysqlRepository) RemoveUser(username int64) error {
	err := m.db.DeleteUser(username)
	if err != nil {
		return fmt.Errorf("Remove Ploblem: %v", err)
	}
	return nil
}
