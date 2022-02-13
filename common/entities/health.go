package entities

type Health struct {
	MaxHealth int
	Health    int
}

func (h *Health) SetHealth(health int) {
	h.MaxHealth = health
	h.Health = health
}

func (h *Health) Damage(damage int) {
	h.Health -= damage
	if h.Health < 0 {
		h.Health = 0
	}
}

func (h *Health) Dead() bool {
	return h.Health == 0
}

func (h *Health) Heal(health int) {
	h.Health += health
}

func (h *Health) GetMaxHealth() int {
	return h.MaxHealth
}
