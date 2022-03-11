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

const steadyTime = time.Millisecond * 1000 / 5

type Manager struct {
	world.World
	game.Game

	time     *types.Time
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
	var time = types.NewTime()
	var m = &Manager{
		time: time,
		Game: *game.NewGame(),
	}
	var spawner = world.Spawner(m)
	m.World = *world.NewWorld(&spawner)
	var additions = world.NewAdditions(&m.World, time)
	var deletions = world.NewDeletions(&m.World, time)
	m.SetAdditions(additions)
	m.SetDeletions(deletions)

	var gEngine = engine.NewManager()
	var gDrawer = drawer.NewDrawer()
	var packets = ws.NewManager()
	var inputs = input.NewManager()

	var wsSpawner = ws.Spawner(m)
	var inputsSystem = ws.System(inputs)
	var gameListener = ws.Listener(m)
	var additionsListener = ws.Listener(additions)
	var deletionsListener = ws.FutureListener(deletions)
	var futurePositionListener = ws.FutureListener(gEngine)
	packets.SetSpawner(&wsSpawner)
	packets.AddSystem(&inputsSystem)
	packets.AddListener(&gameListener)
	packets.AddListener(&additionsListener)
	packets.AddFutureListener(&deletionsListener)
	packets.AddFutureListener(&futurePositionListener)

	m.drawer = gDrawer
	m.engine = gEngine
	m.packets = packets
	m.inputs = inputs

	return m
}

func (m *Manager) ExposeFunctions(doc js.Value, group *errgroup.Group, ctx context.Context) {
	doc.Set("connectToServer", js.FuncOf(func(this js.Value, values []js.Value) interface{} {
		var hostname = values[0].String()
		var port = values[1].Int()
		var username = types.UserID(values[2].String())

		group.Go(func() error {
			return m.Connect(ctx, hostname, port, username)
		})
		return nil
	}))
	doc.Set("setInputs", js.FuncOf(m.inputs.SetInputs))
	doc.Set("drawUpdate", js.FuncOf(m.drawer.DrawUpdate))
	doc.Set("centerUpdate", js.FuncOf(m.drawer.CenterUpdate))
	doc.Set("engineUpdate", js.FuncOf(m.Update))
}
