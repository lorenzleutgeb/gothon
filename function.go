package main

type Function struct {
	AttributedObject
	Name *String
	Code *Code
}

func (function *Function) Read(reader *Reader, t byte) {
}

func (function *Function) String() string {
	return "<funct \"" + function.Name.string + "\">"
}
