package types

type Change struct {
	changed bool
}

func (c *Change) Set() {
	c.changed = true
}

func (c *Change) Changed() bool {
	return c.changed
}

func (c *Change) Reset() {
	c.changed = false
}
