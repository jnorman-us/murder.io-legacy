package innocent

import (
	"github.com/josephnormandev/murder/common/types"
)

func (i *Innocent) HandleInputStateChange(s types.Input) {
	i.input = s
}
