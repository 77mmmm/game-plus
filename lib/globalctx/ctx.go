package globalctx

import (
	"game-pro/model"
	"game-pro/model/repository"
)

type GlobalCtx struct {
	LeaderBoard model.LeaderBoardRepository
}

func GetGlobalCtx() *GlobalCtx {
	return &GlobalCtx{
		LeaderBoard: repository.NewLeaderBoard(),
	}
}
