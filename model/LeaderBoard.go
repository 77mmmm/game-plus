package model

import "time"

type LeaderBoard struct {
	Uuid            int64
	Score           int64
	ScoreCreateTime time.Time
	ScoreUpdateTime time.Time
}
