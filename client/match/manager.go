package match

import (
	"context"
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/engine"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/client/worldin"
	"github.com/josephnormandev/murder/client/ws"
	"github.com/josephnormandev/murder/common/game"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
	"golang.org/x/sync/errgroup"
	"syscall/js"
)

type Manager struct {
	world.World
	game.Game

	Username types.UserID

	engine   *engine.Manager
	packets  *ws.Manager
	drawer   *drawer.Drawer
	inputs   *input.Manager
	worldIn  *worldin.Manager
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

	var worldOutput = worldin.Output(&m.World)
	var worldIn = worldin.NewManager(packets, &worldOutput)

	m.drawer = gDrawer
	m.engine = gEngine
	m.packets = packets
	m.inputs = inputs
	m.worldIn = worldIn

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
