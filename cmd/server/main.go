package main

import (
	"fmt"
	"game-pro/handler/rank"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {

	r := fasthttprouter.New()
	leaderboardclient(r)
	leaderboardmanager(r)
	fmt.Println("Server listening on port 8081...")
	err := fasthttp.ListenAndServe(":8081", r.Handler)
	if err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}

}

func leaderboardclient(r *fasthttprouter.Router) {
	r.GET("/leaderboard/rank", rank.GetUserRank)
	r.GET("/leaderboard/userscore", rank.GetUserScore)
	r.GET("/leaderboard/board", rank.GetPageRank)
}
func leaderboardmanager(r *fasthttprouter.Router) {
	r.POST("/leaderboard/manager/add", rank.AddUserIntoLeaderBoard)
	r.POST("/leaderboard/manager/update", rank.UpdateUser)
}
