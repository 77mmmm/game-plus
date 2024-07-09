package model

type LeaderBoardRepository interface {
	AddUser(username int64, userscore int64) error
	GetUserScore(username int64) (int64, error)
	GetUserRank(username int64) (int64, error)
	GetBoard(start, end int) ([]LeaderBoard, error)
	UpdateUser(username int64, score int64) error
	RemoveUser(username int64) error
}
