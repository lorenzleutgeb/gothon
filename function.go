package main

type Function struct {
	AttributedObject

	// For internal Functions
	Name     string
	Callable Callable

	// For external Functions
	Code *Code
}

func (f *Function) IsInternal() bool {
	return f.Code == nil
}

func (f *Function) Call(args *args) Object {
	if f.IsInternal() {
		return f.Callable(args)
	}

	frame := NewFrame(f.Code)

	if args != nil {
		for i, value := range args.Positional {
			name := f.Code.Varnames[i]
			frame.names[name.String()] = value
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

func NewInternalFunction(name string, callable Callable) Function {
	return Function{
		Name:     name,
		Callable: callable,
	}
}

func NewExternalFunction(name string, code *Code) Function {
	return Function{
		Name: name,
		Code: code,
	}
}
