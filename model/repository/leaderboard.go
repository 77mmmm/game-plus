package repository

import (
	"game-pro/model"
	"game-pro/model/cache"
	"game-pro/model/dao"
)

var leaderboard_key = "leader_board"

type LeaderBoard struct {
	client *cache.RedisRepository
	db     *dao.MysqlRepository
}

func NewLeaderBoard() *LeaderBoard {
	return &LeaderBoard{
		db: dao.NewMysqlRepository(),
		//client: cache.NewRedisRepository(),
	}
}

func (l *LeaderBoard) AddUser(username int64, userscore int64) error {
	err := l.db.AddUser(username, userscore)
	if err != nil {
		return err
	}
	return nil
}

func (l *LeaderBoard) UpdateUser(username int64, score int64) error {
	err := l.db.UpdateUserScore(username, score)
	if err != nil {
		return err
	}
	return nil
}

func (l *LeaderBoard) GetUserScore(username int64) (int64, error) {
	score, err := l.db.GetUserScore(username)
	if err != nil {
		return -1, err
	}
	return score, nil
}

func (l *LeaderBoard) GetUserRank(username int64) (int64, error) {
	rank, err := l.db.GetUserRank(username)
	if err != nil {
		return -1, err
	}
	return rank, nil
}

func (l *LeaderBoard) GetBoard(start int, end int) ([]model.LeaderBoard, error) {
	leaderboard, err := l.db.GetLimitLeaderBoard(start, end)
	if err != nil {
		return nil, err
	}
	return leaderboard, nil
}

func (l *LeaderBoard) RemoveUser(username int64) error {
	err := l.db.RemoveUser(username)
	if err != nil {
		return err
	}
	return nil
}
