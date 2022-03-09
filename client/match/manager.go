package match

import (
	"context"
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

func (m *Manager) Connect(background context.Context, hostname string, port int, username types.UserID) error {
	m.wsClient = ws.NewClient(m.packets, hostname, port, username)
	err := m.wsClient.Connect(background)
	fmt.Println("Stopping connection!")
	return err
}

func (m *Manager) UpdateTick(background context.Context) error {
	for range time.Tick(updateTime) {
		select {
		case <-background.Done():
			fmt.Println("Stopping update tick")
			return background.Err()
		default:
			m.engine.UpdatePhysics(updateTime)
		}
	}
	return nil
}

func (m *Manager) SteadyTick(background context.Context) error {
	for range time.Tick(steadyTime) {
		select {
		case <-background.Done():
			fmt.Println("Stopping steady tick")
			return background.Err()
		default:
			err := m.packets.SteadyTick()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
