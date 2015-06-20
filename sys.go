package gothon

import (
	"container/list"
	"os"
	"runtime"
)

// NewSys constructs the sys module for Gothon.
// Refer to https://docs.python.org/3.4/library/sys.html
func NewSys() *Code {
	sys := &Code{}
	sys.Name = &String{"sys"}
	sys.Filename = &String{"builtin"}
	sys.AddAttribute("copyright", &String{"Copyright (c) 2015 Lorenz Leutgeb."})
	sys.AddAttribute("builtin_module_names", &Tuple{&String{"sys"}})
	sys.AddAttribute("platform", &String{runtime.GOOS})
	sys.AddAttribute("exit", &String{"Nope, you are not exiting."})
	argList := list.New()
	for _, arg := range os.Args {
		argList.PushBack(arg)
	}
	sys.AddAttribute("argv", &List{*argList})
	return sys
}
