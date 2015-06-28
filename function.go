package main

// Function is an implementation of some Python function.
// A Function can be external (some CPython bytecode) or
// internal (implemented in Go inside gothon)
type Function struct {
	AttributedObject

	// For internal Functions
	Name     string
	Callable Callable

	// For external Functions
	Code *Code
}

// IsInternal tells whether this Function is implemented
// in Go (part of the interpreter itself)
func (f *Function) IsInternal() bool {
	return f.Code == nil
}

// Call allocates a new stack frame as execution context
// for this Function, executes it and returns the result
// of the invocation.
func (f *Function) Call(args *args) Object {
	if f.IsInternal() {
		return f.Callable(args)
	}

	frame := NewFrame(f.Code)

	if args != nil {
		for i, value := range args.Positional {
			name := f.Code.Varnames[i].(String)
			frame.names[name.string] = value
		}
		if len(args.Keyword) > 0 {
			panic("Keyword arguments are not implemented")
		}
	}

	return frame.Execute()
}

func (f Function) String() string {
	if f.IsInternal() {
		return "<internal function \"" + f.Name + "\">"
	}
	return "<external function \"" + f.Code.String() + "\">"
}

// NewInternalFunction allocates a new Function that is implemented in
// Go.
func NewInternalFunction(name string, callable Callable) Function {
	return Function{
		Name:     name,
		Callable: callable,
	}
}

// NewExternalFunction allocates a new Function that encapsulates
// CPython bytecode.
func NewExternalFunction(name string, code *Code) Function {
	return Function{
		Name: name,
		Code: code,
	}
}
