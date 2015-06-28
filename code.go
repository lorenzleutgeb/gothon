package main

import (
	"fmt"
	"path"
)

// Code is a code object unmarshaled from CPython bytecode,
// it holds executable instructions and additional metadata.
type Code struct {
	AttributedObject

	PosArgCnt, // number of positional arguments (including arguments with default values)
	KwArgCnt, // number of keyword arguments
	Nlocals, // number of local variables used by the function (including arguments)
	Stacksize, // required stack size (including local variables)
	Flags uint32

	Instructions []byte

	Consts, // contains the literals used by the bytecode. If a code object represents a function, the first item in consts is the documentation string of the function, or None if undefined.
	Names, // is a tuple containing the names used by the bytecode
	Varnames, // is a tuple containing the names of the local variables (starting with the argument names)
	Freevars, // is a tuple containing the names of free variables
	Cellvars Tuple // is a tuple containing the names of local variables that are referenced by nested functions

	Filename String // is the filename from which the code was compiled
	Name     String // gives the function name

	Firstlineno int32  // is the first line number of the function
	Lnotab      String // is a string encoding the mapping from bytecode offsets to line numbers (for details see the source code of the interpreter)
}

func (code Code) String() string {
	return fmt.Sprintf("<code \"%s\" \".../%s:%d\">", code.Name.string, path.Base(code.Filename.string), code.Firstlineno)
}
