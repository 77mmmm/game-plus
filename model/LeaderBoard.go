package model

import "time"

type LeaderBoard struct {
	Uuid      int64
	Score     int64
	Scoretime time.Time
}
