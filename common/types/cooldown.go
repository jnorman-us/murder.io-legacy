package types

type CoolDown struct {
	time, remaining int
}

func NewCoolDown(time int) CoolDown {
	return CoolDown{
		time:      time,
		remaining: 0,
	}
}

func (c *CoolDown) Ready() bool {
	return c.remaining == 0
}

func (c *CoolDown) Tick() {
	if c.remaining > 0 {
		c.remaining--
	}
}

func (c *CoolDown) Reset() {
	c.remaining = c.time
}
