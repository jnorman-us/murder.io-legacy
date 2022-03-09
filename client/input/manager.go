package input

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/types"
	"syscall/js"
)

type Manager struct {
	inputs  types.Input
	current map[int]bool
}

func NewManager() *Manager {
	var input = &Manager{
		inputs: types.Input{},
	}
	return input
}

func (m *Manager) SetInputs(this js.Value, values []js.Value) interface{} {
	var jsObject = values[0]
	var forwardValue = jsObject.Get("forward")
	var backwardValue = jsObject.Get("backward")
	var leftValue = jsObject.Get("left")
	var rightValue = jsObject.Get("right")
	var attackClickValue = jsObject.Get("attackClick")
	var rangedClickValue = jsObject.Get("rangedClick")
	var specialValue = jsObject.Get("special")
	var directionValue = jsObject.Get("direction")

	if forwardValue.IsUndefined() || forwardValue.Type() != js.TypeBoolean ||
		backwardValue.IsUndefined() || backwardValue.Type() != js.TypeBoolean ||
		leftValue.IsUndefined() || leftValue.Type() != js.TypeBoolean ||
		rightValue.IsUndefined() || rightValue.Type() != js.TypeBoolean ||
		attackClickValue.IsUndefined() || attackClickValue.Type() != js.TypeBoolean ||
		rangedClickValue.IsUndefined() || rangedClickValue.Type() != js.TypeBoolean ||
		specialValue.IsUndefined() || specialValue.Type() != js.TypeBoolean ||
		directionValue.IsUndefined() || directionValue.Type() != js.TypeNumber {
		return js.Error{
			Value: js.ValueOf("Incorrect input parameters"),
		}
	}

	var inputs = types.Input{
		Forward:     forwardValue.Bool(),
		Backward:    backwardValue.Bool(),
		Left:        leftValue.Bool(),
		Right:       rightValue.Bool(),
		AttackClick: attackClickValue.Bool(),
		RangedClick: rangedClickValue.Bool(),
		Special:     specialValue.Bool(),
		Direction:   directionValue.Float(),
	}

	m.inputs = inputs
	return nil
}

func (m *Manager) GetChannel() byte {
	return 0x02
}

func (m *Manager) GetData(e *gob.Encoder) error {
	err := e.Encode(m.inputs)
	return err
}
