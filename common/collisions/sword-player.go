package collisions

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
)

type SwordPlayer interface {
	GetCollider() *collider.Collider
	GetWielder() int
	GetWielderUsername() string
}

type PlayerSword interface {
	CheckCollision(*collider.Collider) bool
	GetID() int
	SlainBy(int, string)
}

func (m *Manager) AddSwordPlayer(id int, s *SwordPlayer) {
	m.SwordPlayers[id] = s
}

func (m *Manager) RemoveSwordPlayer(id int) {
	delete(m.SwordPlayers, id)
}

func (m *Manager) AddPlayerSword(id int, p *PlayerSword) {
	m.PlayerSwords[id] = p
}

func (m *Manager) RemovePlayerSword(id int) {
	delete(m.PlayerSwords, id)
}

func (m *Manager) resolveSwordPlayers() {
	for _, p := range m.PlayerSwords {
		var player = *p
		for _, s := range m.SwordPlayers {
			var sword = *s
			var swordCollider = sword.GetCollider()
			if player.CheckCollision(swordCollider) {
				if player.GetID() != sword.GetWielder() {
					player.SlainBy(sword.GetWielder(), sword.GetWielderUsername())

				}
			}
		}
	}
}
