package input

type Input struct {
	Forward     bool // w
	Backward    bool // s
	Left        bool // a
	Right       bool // d
	AttackClick bool // left-click
	RangedClick bool // right-click
	Special     bool // space bar
}

func (s Input) Equals(o Input) bool {
	return s.Forward == o.Forward &&
		s.Backward == o.Backward &&
		s.Left == o.Left &&
		s.Right == o.Right &&
		s.AttackClick == o.AttackClick &&
		s.RangedClick == o.RangedClick &&
		s.Special == o.Special
}