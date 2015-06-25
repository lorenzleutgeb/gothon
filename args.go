package main

// Holds arguments to an invocation of a function.
// as Python support positional arguments and keyword
// arguments, this encapsulates both.
type args struct {
	Positional []Object
	Keyword    map[string]Object
}

func (a args) IsEmpty() bool {
	return len(a.Positional)+len(a.Keyword) == 0
}

func newArgs() *args {
	return &args{
		make([]Object, 3),
		make(map[string]Object),
	}
}

// Callable is every function that takes Python arguments.
type Callable func(args *args) Object
