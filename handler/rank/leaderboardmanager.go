package rank

import (
	"fmt"
	"game-pro/lib/globalctx"
	"github.com/valyala/fasthttp"
)

func UpdateUser(ctx *fasthttp.RequestCtx) {
	gctx := globalctx.GetGlobalCtx()
	values := ctx.QueryArgs()
	username, err := values.GetUint("username")
	if err != nil {
		ctx.Error(fmt.Sprintf("error get username: %s", err), fasthttp.StatusBadRequest)
	}
	score, err1 := values.GetUint("score")
	if err1 != nil {
		ctx.Error(fmt.Sprintf("error get score: %s", err), fasthttp.StatusBadRequest)
	}
	err = gctx.LeaderBoard.UpdateUser(int64(username), int64(score))
	if err != nil {
		ctx.Error(fmt.Sprintf("error update leaderboard: %s", err), fasthttp.StatusInternalServerError)
	}
	fmt.Fprintf(ctx, "%s:update success", username)
}

func AddUserIntoLeaderBoard(ctx *fasthttp.RequestCtx) {
	gctx := globalctx.GetGlobalCtx()
	values := ctx.QueryArgs()
	username, err := values.GetUint("username")
	if err != nil {
		ctx.Error(fmt.Sprintf("error add username: %s", err), fasthttp.StatusBadRequest)
	}
	score, err1 := values.GetUint("score")
	if err1 != nil {
		ctx.Error(fmt.Sprintf("error add score: %s", err), fasthttp.StatusBadRequest)
	}
	err = gctx.LeaderBoard.AddUser(int64(username), int64(score))
	if err != nil {
		ctx.Error(fmt.Sprintf("error add user:%s", err), fasthttp.StatusBadRequest)
	}
	fmt.Fprintf(ctx, "%s:add success", username)

}

func RemoveUserFromLeaderBoard(ctx *fasthttp.RequestCtx) {
	gctx := globalctx.GetGlobalCtx()
	values := ctx.QueryArgs()
	username, err := values.GetUint("username")
	if err != nil {
		ctx.Error(fmt.Sprintf("error get username: %s", err), fasthttp.StatusBadRequest)
	}
	err = gctx.LeaderBoard.RemoveUser(int64(username))
	if err != nil {
		ctx.Error(fmt.Sprintf("error remove user:%s", err), fasthttp.StatusBadRequest)
	}
	fmt.Fprintf(ctx, "%s:remove success", username)
}
