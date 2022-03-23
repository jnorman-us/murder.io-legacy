package packets

type FloatDiff struct {
	Field  byte
	Offset byte
	Value  float32
}

type IntDiff struct {
	Field  byte
	Offset byte
	Value  int32
}

type StringDiff struct {
	Field  byte
	Offset byte
	Value  string
}

type Addition struct {
	Data   *Data
	Offset byte
}

type Deletion struct {
	Data   *Data
	Offset byte
}
