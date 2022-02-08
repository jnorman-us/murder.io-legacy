package main

import (
	"github.com/gorilla/mux"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/server/game"
	"github.com/josephnormandev/murder/server/ws"
	"log"
	"net/http"
)

var wsServer *ws.Server
var soleGame *game.ServerGame

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

func main() {
	soleGame = game.NewServerGame(0)
	wsServer = ws.NewServer()

	var lobby = soleGame.GetLobby()
	wsServer.AddLobby(lobby)
	go lobby.Send()

	soleGame.SetPlayers(names)

	var staticFiles = http.FileServer(http.Dir("./server/static"))

	var router = mux.NewRouter().StrictSlash(true)
	router.Handle("/ws/{id}", wsServer)
	router.PathPrefix("/").Handler(staticFiles)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
