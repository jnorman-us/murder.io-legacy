package timestamp

import (
	"time"
)

type Timestamp struct {
	Tick
	startTime time.Time
}

func NewTimestamp() *Timestamp {
	return &Timestamp{
		startTime: time.Now(),
	}
}

func (t *Timestamp) GetOffset() time.Duration {
	return time.Now().Sub(t.startTime)
}

func (t *Timestamp) GetOffsetBytes() byte {
	var duration = t.GetOffset()
	duration /= time.Millisecond
	return byte(duration)
}

func (t *Timestamp) TimeTick() {
	t.Tick++
	t.startTime = time.Now()
}
