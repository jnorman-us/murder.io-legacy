package action

type Action byte

type action struct {
	Add    Action
	Update Action
	Delete Action
}

var Actions = action{
	Add:    0x01,
	Update: 0x02,
	Delete: 0x03,
}
