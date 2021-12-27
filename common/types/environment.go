package types

type Environment bool

func (e Environment) IsClient() bool {
	return e == true
}

func (e Environment) IsServer() bool {
	return e == false
}

func ServerEnvironment() Environment {
	return false
}

func ClientEnvironment() Environment {
	return true
}
