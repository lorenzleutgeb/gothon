package main

import (
	"bufio"
	"encoding/binary"
	"time"
)

type Module struct {
	AttributedObject
	Code       *Code
	Interns    []*String
	mtime      time.Time
	Size       uint32
	Version    uint16
	References []*Object
}

func NewModule(br *bufio.Reader) (module *Module) {
	module = new(Module)

	module.References = make([]*Object, 0)

	r := Reader{
		*br,
		module,
	}

	binary.Read(&r, binary.LittleEndian, &module.Version)

	if module.Version != 0x0cee {
		panic("invalid version")
	}

	var check uint16
	binary.Read(&r, binary.LittleEndian, &check)

	if check != 0x0a0d {
		panic("invalid magic")
	}

	var rawMtime int32
	binary.Read(&r, binary.LittleEndian, &rawMtime)
	module.mtime = time.Unix(int64(rawMtime), 0)
	binary.Read(&r, binary.LittleEndian, &module.Size)

	code := r.ReadObject().(Code)
	module.Code = &code

	return module
}
