package match

import (
	"fmt"
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/engine"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/client/ws"
	"github.com/josephnormandev/murder/common/game"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
)

type Manager struct {
	world.World
	game.Game

	Username types.UserID

	engine  *engine.Manager
	packets *ws.Manager
	drawer  *drawer.Drawer
	inputs  *input.Manager
}

func NewManager() *Manager {
	var manager = &Manager{
		World: *world.NewWorld(),
		Game:  *game.NewGame(),
	}
	var spawner = world.Spawner(manager)
	manager.SetSpawner(&spawner)

	var gEngine = engine.NewManager()
	var gDrawer = drawer.NewDrawer()
	var packets = ws.NewManager()
	var inputs = input.NewManager()

	var sizeable = input.Sizeable(gDrawer)
	inputs.SetSizeable(&sizeable)

	var wsSpawner = ws.Spawner(manager)
	var inputsSystem = ws.System(inputs)
	var gameListener = ws.Listener(manager)
	var deletionsListener = ws.Listener(manager.Deletions())
	var futurePositionListener = ws.FutureListener(gEngine)
	packets.SetSpawner(&wsSpawner)
	packets.AddSystem(&inputsSystem)
	packets.AddListener(&gameListener)
	packets.AddListener(&deletionsListener)
	packets.AddFutureListener(&futurePositionListener)

	manager.drawer = gDrawer
	manager.engine = gEngine
	manager.packets = packets
	manager.inputs = inputs

	return manager
}

func (m *Manager) GetPackets() *ws.Manager {
	return m.packets
}

func (m *Manager) Start() {
	m.drawer.Start(m.engine.UpdatePhysics)
}

func (m *Manager) SteadyTick() {
	err := m.packets.SteadyTick()
	if err != nil {
		fmt.Println(err)
	}
}
