package pyc

import (
	"encoding/binary"
)

type Code struct {
	argcount int32			// is the number of positional arguments (including arguments with default values)
	kwonlyargcount int32
	nlocals int32			// is the number of local variables used by the function (including arguments)
	stacksize int32		// is the required stack size (including local variables)
	flags int32				// is an integer encoding a number of flags for the interpreter:
									// bit 0x04 is set if the function uses the *arguments syntax to accept an arbitrary number of positional arguments
									// bit 0x08 is set if the function uses the **keywords syntax to accept arbitrary keyword arguments
									// bit 0x20 is set if the function is a generator
									// Future feature declarations (from __future__ import division) also use bits in co_flags to indicate whether a code object was compiled with a particular feature enabled: bit 0x2000 is set if the function was compiled with future division enabled; bits 0x10 and 0x1000 were used in earlier versions of Python.
	
	code String				// is a string representing the sequence of bytecode instructions
	
	consts Tuple		// is a tuple containing the literals used by the bytecode. If a code object represents a function, the first item in consts is the documentation string of the function, or None if undefined.
	names Tuple			// is a tuple containing the names used by the bytecode
	varnames Tuple		// is a tuple containing the names of the local variables (starting with the argument names)
	freevars Tuple		// is a tuple containing the names of free variables
	cellvars Tuple		// is a tuple containing the names of local variables that are referenced by nested functions
	
	filename String		// is the filename from which the code was compiled
	name String				// gives the function name
	
	firstlineno int32		// is the first line number of the function
	lnotab String			// is a string encoding the mapping from bytecode offsets to line numbers (for details see the source code of the interpreter)
}

func (code *Code) Read(reader *Reader) {
	binary.Read(reader, binary.LittleEndian, &code.argcount)
	binary.Read(reader, binary.LittleEndian, &code.kwonlyargcount)
	binary.Read(reader, binary.LittleEndian, &code.nlocals)
	binary.Read(reader, binary.LittleEndian, &code.stacksize)
	binary.Read(reader, binary.LittleEndian, &code.flags)
	reader.ReadExpected(&code.code)
	reader.ReadExpected(&code.consts)
	reader.ReadExpected(&code.names)
	reader.ReadExpected(&code.varnames)
	reader.ReadExpected(&code.freevars)
	reader.ReadExpected(&code.cellvars)
	reader.ReadExpected(&code.filename)
	reader.ReadExpected(&code.name)
	binary.Read(reader, binary.LittleEndian, &code.firstlineno)
	reader.ReadExpected(&code.lnotab)
}
