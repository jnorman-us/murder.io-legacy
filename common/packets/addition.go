package packets

import (
	"github.com/josephnormandev/murder/common/types"
	"time"
)

// Addition is a struct for holding data about the spawning of an
// object and is only used for the bare minimum spawn details
// since further information will come later
type Addition struct {
	types.ID
	Offset time.Duration
	Class  byte
}
