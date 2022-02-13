package main

import (
	"github.com/gorilla/mux"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/server/match"
	"github.com/josephnormandev/murder/server/ws"
	"log"
	"net/http"
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

	var lobby = soleGame.GetPackets()
	wsServer.AddLobby(lobby)

	go soleGame.Tick()
	go lobby.Send()

	soleGame.SetPlayers(names)

	for _, name := range names {
		var drifter = drifter.NewDrifter()
		drifter.UserID = name
		drifter.SetPosition(types.NewRandomVector(0, 0, 400, 400))
		soleGame.AddDrifter(drifter)
	}

	for _, position := range polePositions {
		var newPole = pole.NewPole()
		newPole.SetPosition(position)
		soleGame.AddPole(newPole)
	}

	var staticFiles = http.FileServer(http.Dir("./server/static"))

	var router = mux.NewRouter().StrictSlash(true)
	router.Handle("/ws/{id}", wsServer)
	router.PathPrefix("/").Handler(staticFiles)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
