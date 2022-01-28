package types

type Tick int

func (t Tick) Iterate() {
	t += 1
}
