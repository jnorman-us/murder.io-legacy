package match

import (
	"context"
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/engine"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/client/ws"
	"github.com/josephnormandev/murder/common/game"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
	"golang.org/x/sync/errgroup"
	"syscall/js"
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

	RunContext context.Context
	RunGroup   *errgroup.Group
}

func NewManager() *Manager {
	var m = &Manager{
		Game: *game.NewGame(),
	}
	var spawner = world.Spawner(m)
	m.World = *world.NewWorld(&spawner)

	var gEngine = engine.NewManager()
	var gDrawer = drawer.NewDrawer()
	var packets = ws.NewManager()
	var inputs = input.NewManager()

	var wsSpawner = ws.Spawner(m)
	var inputsSystem = ws.System(inputs)
	var gameListener = ws.Listener(m)
	var deletionsListener = ws.Listener(m.Deletions())
	var futurePositionListener = ws.FutureListener(gEngine)
	packets.SetSpawner(&wsSpawner)
	packets.AddSystem(&inputsSystem)
	packets.AddListener(&gameListener)
	packets.AddListener(&deletionsListener)
	packets.AddFutureListener(&futurePositionListener)

	m.drawer = gDrawer
	m.engine = gEngine
	m.packets = packets
	m.inputs = inputs

	m.RunGroup, m.RunContext = errgroup.WithContext(context.Background())

	return m
}

func (m *Manager) ExposeFunctions(document js.Value) {
	document.Set("connectToServer", js.FuncOf(m.Connect))
	document.Set("setInputs", js.FuncOf(m.inputs.SetInputs))
	document.Set("getDrawData", js.FuncOf(m.drawer.GetDrawData))
}
