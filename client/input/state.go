package input

type Input struct {
	forward     bool // w
	backward    bool // s
	left        bool // a
	right       bool // d
	attackClick bool // left-click
	rangedClick bool // right-click
	special     bool // space bar
}

func (s Input) Equals(o Input) bool {
	return s.forward == o.forward &&
		s.backward == o.backward &&
		s.left == o.left &&
		s.right == o.right &&
		s.attackClick == o.attackClick &&
		s.rangedClick == o.rangedClick &&
		s.special == o.special
}
