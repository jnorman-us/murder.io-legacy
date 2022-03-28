package main

import (
	"github.com/gorilla/mux"
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/server/match"
	"github.com/josephnormandev/murder/server/ws"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var wsServer *ws.Server
var soleGame *match.Match

var names = []types.UserID{
	"Jellotinous",
	"Wine_Craft",
	"Xiehang",
	"TheStorminNorman",
	"ShadowDragon",
	"Society Member",
	"Envii",
	"Jinseng",
	"Laerir",
	"JoeyD",
	"Beta Tester",
}

var polePositions = []types.Vector{
	types.NewVector(0, 0),
	types.NewVector(500, 0),
	types.NewVector(0, 500),
	types.NewVector(500, 500),
}

func main() {
	soleGame = match.NewMatch(0)
	wsServer = ws.NewServer()

	var lobby = soleGame.GetLobby()
	wsServer.AddLobby(lobby)

	soleGame.SetPlayers(names)

	rand.Seed(time.Now().UnixNano())
	for _, name := range names {
		var d = dimetrodon.NewDimetrodon()
		d.UserID = name
		d.SetPosition(types.NewRandomVector(0, 0, 400, 400))
		soleGame.AddDimetrodon(d)
		if rand.Intn(6) == 1 {
			d.Input.AttackClick = true
		}
		if rand.Intn(2) == 1 {
			d.Input.Left = true
		} else {
			d.Input.Right = true
		}
	}

	for _, position := range polePositions {
		var newPole = pole.NewPole()
		newPole.SetPosition(position)
		soleGame.AddPole(newPole)
	}

	go soleGame.Tick()
	go soleGame.Send()

	var staticFiles = http.FileServer(http.Dir("./server/static"))

	var router = mux.NewRouter().StrictSlash(true)
	router.Handle("/ws/{id}", wsServer)
	router.PathPrefix("/").Handler(staticFiles)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
