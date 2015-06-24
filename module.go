package main

import (
	"container/list"
	"encoding/binary"
	"log"
	"time"
)

type Module struct {
	AttributedObject
	Code       *Code
	Interns    list.List
	mtime      time.Time
	Size       uint32
	Version    uint16
	References []*Object
}

func (module *Module) Read(reader *Reader, t byte) {
	module.References = make([]*Object, 0)

	binary.Read(reader, binary.LittleEndian, &module.Version)
	var check uint16
	binary.Read(reader, binary.LittleEndian, &check)

	if check != 0xa0d {
		log.Fatalf("Second two bytes of magic number are incorrect (0x0a0d expected, 0x%x found)", check)
	}

	var mtime int32
	binary.Read(reader, binary.LittleEndian, &mtime)
	module.mtime = time.Unix(int64(mtime), 0)
	//log.Printf("Last modified: %v", module.mtime)
	binary.Read(reader, binary.LittleEndian, &module.Size)
	module.Code = reader.ReadCode()
}

func (module *Module) AddReference(o Object) {
	module.References = append(module.References, &o)
}

func (module *Module) GetReference(index int) Object {
	return *module.References[index]
}
