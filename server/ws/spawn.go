package ws

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/types"
)

type Spawn interface {
	GetID() types.ID
	GetClass() types.Channel
	Dirty() bool
	CleanDirt()
	GetStartData() data.Data
	GetData() data.Data
}

func (l *Lobby) AddSpawn(id types.ID, s *Spawn) {
	l.Lock()
	defer l.Unlock()

	l.spawns[id] = s
	l.additions[id] = s
	l.addData[id] = (*s).GetStartData()
	l.addTimes[id] = l.time.GetOffset()
}

func (l *Lobby) RemoveSpawn(id types.ID) {
	l.Lock()
	defer l.Unlock()

	var s = l.spawns[id]
	l.deletions[id] = s
	l.deleteData[id] = (*s).GetData()
	l.deleteTimes[id] = l.time.GetOffset()

	delete(l.spawns, id)
}
