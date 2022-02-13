package types

type Change struct {
	Flag bool
}

func (c *Change) Set() {
	c.Flag = true
}

func (c *Change) Changed() bool {
	return c.Flag
}

func (c *Change) Reset() {
	c.Flag = false
}
