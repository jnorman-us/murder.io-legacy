package types

type Change struct {
	Flag bool
}

func (c *Change) MarkDirty() {
	c.Flag = true
}

func (c *Change) Dirty() bool {
	return c.Flag
}

func (c *Change) CleanDirt() {
	c.Flag = false
}
