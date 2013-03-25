package gothon

type Function struct {
	Name *String
	Code *Code
}

func (function *Function) Read(reader *Reader) {
}

func (function *Function) String() string {
	return "[func: " + function.Name.string + "]"
}
