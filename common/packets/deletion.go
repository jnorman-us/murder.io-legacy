package packets

import (
	"github.com/josephnormandev/murder/common/types"
	"time"
)

type Deletion struct {
	types.ID
	Offset time.Duration
}
