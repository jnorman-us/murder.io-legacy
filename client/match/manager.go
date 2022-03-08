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
	"time"
)

const steadyTime = time.Millisecond * 1000 / 20
const updateTime = time.Millisecond * 1000 / 60

type Manager struct {
	world.World
	game.Game

	Username types.UserID

	engine   *engine.Manager
	packets  *ws.Manager
	drawer   *drawer.Drawer
	inputs   *input.Manager
	wsClient *ws.Client
}

func NewManager() *Manager {
	var manager = &Manager{
		Game: *game.NewGame(),
	}
	var spawner = world.Spawner(manager)
	manager.World = *world.NewWorld(&spawner)

	var gEngine = engine.NewManager()
	var gDrawer = drawer.NewDrawer()
	var packets = ws.NewManager()
	var inputs = input.NewManager()

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

func (m *Manager) Connect(hostname string, port int, username types.UserID) {
	m.wsClient = ws.NewClient(m.packets, hostname, port, username)
	go (func() {
		err := m.wsClient.Connect()
		if err != nil {
			fmt.Printf("Error with WS! %v\n", err)
		}
	})()
}

func (m *Manager) Update() {
	for range time.Tick(updateTime) {
		m.engine.UpdatePhysics(updateTime)
	}
}

func (m *Manager) SteadyTick() {
	for range time.Tick(steadyTime) {
		err := m.packets.SteadyTick()
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
