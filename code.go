package gothon

import (
	"encoding/binary"
	"log"
)

type Code struct {
	Argcount int32			// is the number of positional arguments (including arguments with default values)
	Kwonlyargcount int32
	Nlocals int32			// is the number of local variables used by the function (including arguments)
	Stacksize int32		// is the required stack size (including local variables)
	Flags int32				// is an integer encoding a number of flags for the interpreter:
									// bit 0x04 is set if the function uses the *arguments syntax to accept an arbitrary number of positional arguments
									// bit 0x08 is set if the function uses the **keywords syntax to accept arbitrary keyword arguments
									// bit 0x20 is set if the function is a generator
									// Future feature declarations (from __future__ import division) also use bits in co_flags to indicate whether a code object was compiled with a particular feature enabled: bit 0x2000 is set if the function was compiled with future division enabled; bits 0x10 and 0x1000 were used in earlier versions of Python.
	
	Instructions []byte				// is a string representing the sequence of bytecode instructions
	
	Consts *Tuple		// is a tuple containing the literals used by the bytecode. If a code object represents a function, the first item in consts is the documentation string of the function, or None if undefined.
	Names *Tuple			// is a tuple containing the names used by the bytecode
	Varnames *Tuple		// is a tuple containing the names of the local variables (starting with the argument names)
	Freevars *Tuple		// is a tuple containing the names of free variables
	Cellvars *Tuple		// is a tuple containing the names of local variables that are referenced by nested functions
	
	Filename *String		// is the filename from which the code was compiled
	Name *String				// gives the function name
	
	Firstlineno int32		// is the first line number of the function
	Lnotab *String			// is a string encoding the mapping from bytecode offsets to line numbers (for details see the source code of the interpreter)
}

func (code *Code) Read(reader *Reader) {
	binary.Read(reader, binary.LittleEndian, &code.Argcount)
	binary.Read(reader, binary.LittleEndian, &code.Kwonlyargcount)
	binary.Read(reader, binary.LittleEndian, &code.Nlocals)
	binary.Read(reader, binary.LittleEndian, &code.Stacksize)
	binary.Read(reader, binary.LittleEndian, &code.Flags)
	
	c, _ := reader.ReadByte()
	
	if c != 's' {
		log.Fatal("Expected code block but got anything else")
	}
	
	var size int32
	binary.Read(reader, binary.LittleEndian, &size)
	code.Instructions = make([]byte, size)
	reader.Read(code.Instructions)
	
	code.Consts = reader.ReadTuple()
	code.Names = reader.ReadTuple()
	code.Varnames = reader.ReadTuple()
	code.Freevars = reader.ReadTuple()
	code.Cellvars = reader.ReadTuple()
	code.Filename = reader.ReadString()
	code.Name = reader.ReadString()
	binary.Read(reader, binary.LittleEndian, &code.Firstlineno)
	code.Lnotab = reader.ReadString()
}
