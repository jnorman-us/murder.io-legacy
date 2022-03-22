package packets

import (
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/types/action"
)

type Packet struct {
	ID         types.ID
	Channel    types.Channel // decoder channel
	Action     action.Action
	Offset     byte
	FloatDiff  []FloatDiff
	IntDiff    []IntDiff
	StringDiff []StringDiff
}
