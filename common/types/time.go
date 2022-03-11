package types

import "time"

type Time struct {
	Tick
	startTime time.Time
}

func NewTime() *Time {
	return &Time{
		startTime: time.Now(),
	}
}

func (t *Time) Reset() {
	t.startTime = time.Now()
}

func (t *Time) GetOffset() time.Duration {
	return time.Now().Sub(t.startTime)
}
