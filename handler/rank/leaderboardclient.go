package rank

import (
	"encoding/json"
	"fmt"
	"game-pro/lib/globalctx"
	"github.com/valyala/fasthttp"
	"log"
)

func GetUserRank(ctx *fasthttp.RequestCtx) {

	gctx := globalctx.GetGlobalCtx()

	values := ctx.QueryArgs()

	username, err := values.GetUint("username")
	if err != nil {
		ctx.Error(fmt.Sprintf("error get username: %s", err), fasthttp.StatusBadRequest)
		return
	}
	rank, err1 := gctx.LeaderBoard.GetUserRank(int64(username))
	if err1 != nil {
		ctx.Error(fmt.Sprintf("error get rank: %s", err1), fasthttp.StatusBadRequest)
		return
	}
	rankJSON, err2 := json.Marshal(rank)
	if err2 != nil {
		ctx.Error(fmt.Sprintf("the trans has done: %s", err2), fasthttp.StatusBadRequest)
		return
	}
	ctx.SetBody(rankJSON)
}

func GetUserScore(ctx *fasthttp.RequestCtx) {
	log.Println("连接成功")
	gctx := globalctx.GetGlobalCtx()
	values := ctx.QueryArgs()
	username, err := values.GetUint("username")
	if err != nil {
		ctx.Error(fmt.Sprintf("error get username: %s", err), fasthttp.StatusBadRequest)
		return
	}
	score, err1 := gctx.LeaderBoard.GetUserScore(int64(username))
	if err1 != nil {
		ctx.Error(fmt.Sprintf("error get score: %s", err1), fasthttp.StatusBadRequest)
	}
	scoreJSON, err2 := json.Marshal(score)
	if err2 != nil {
		ctx.Error(fmt.Sprintf("the trans has done: %s", err2), fasthttp.StatusBadRequest)
	}
	ctx.SetBody(scoreJSON)
}
func GetPageRank(ctx *fasthttp.RequestCtx) {
	gctx := globalctx.GetGlobalCtx()
	values := ctx.QueryArgs()
	page, err := values.GetUint("page")
	pagesize, err := values.GetUint("pagesize")
	if err != nil {
		ctx.Error(fmt.Sprintf("error get page: %s", err), fasthttp.StatusBadRequest)
		return
	}
	if page <= 1 {
		ctx.Error("invalid page argument", fasthttp.StatusBadRequest)
	}
	board, err1 := gctx.LeaderBoard.GetBoard(page*pagesize-pagesize, page*pagesize)
	if err1 != nil {
		ctx.Error(fmt.Sprintf("error get board: %s", err1), fasthttp.StatusBadRequest)
		return
	}
	p := page*pagesize - pagesize
	for _, v := range board {
		boardJSON, err2 := json.Marshal(v)
		if err2 != nil {
			ctx.Error(fmt.Sprintf("the trans has done: %s", err2), fasthttp.StatusBadRequest)
		}
		fmt.Fprintf(ctx, "rank:%d,mes:%s\n", p+1, string(boardJSON))
		p++

	}
}
