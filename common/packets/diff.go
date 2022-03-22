package packets

type FloatDiff struct {
	field  byte
	offset byte
	value  float32
}

type IntDiff struct {
	field  byte
	offset byte
	value  int32
}

type StringDiff struct {
	field  byte
	offset byte
	value  string
}

type Addition struct {
	data   *Data
	offset byte
}

type Deletion struct {
	data   *Data
	offset byte
}
