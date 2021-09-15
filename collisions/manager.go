package collisions

type Manager struct {
	ActorActorCollidables map[int32]*ActorActorCollidable
	ActorWallCollidables  map[int32]*ActorWallCollidable
	WallActorCollidables  map[int32]*WallActorCollidable
}

func NewManager() *Manager {
	return &Manager{
		ActorActorCollidables: map[int32]*ActorActorCollidable{},
		ActorWallCollidables:  map[int32]*ActorWallCollidable{},
		WallActorCollidables:  map[int32]*WallActorCollidable{},
	}
}

func (m *Manager) AddActorActor(id int32, a *ActorActorCollidable) {
	m.ActorActorCollidables[id] = a
}

func (m *Manager) RemoveActorActor(id int32) {
	delete(m.ActorActorCollidables, id)
}

func (m *Manager) AddActorWall(id int32, a *ActorWallCollidable) {
	m.ActorWallCollidables[id] = a
}

func (m *Manager) RemoveActorWall(id int32) {
	delete(m.ActorWallCollidables, id)
}

func (m *Manager) AddWallActor(id int32, w *WallActorCollidable) {
	m.WallActorCollidables[id] = w
}

func (m *Manager) RemoveWallActor(id int32) {
	delete(m.WallActorCollidables, id)
}
