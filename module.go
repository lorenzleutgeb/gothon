package gothon

import (
	"log"
	"time"
	"container/list"
	"encoding/binary"
)

type Module struct {
	Code *Code
	Interns list.List
	mtime time.Time
	Size uint32
	Version uint16
}

func (module *Module) Read(reader *Reader) {
	binary.Read(reader, binary.LittleEndian, &module.Version)
	var check uint16
	binary.Read(reader, binary.LittleEndian, &check)

	if check != 0xa0d {
		log.Fatalf("Second two bytes of magic number are incorrect (0x0a0d expected, 0x%x found)", check)
	}

	var mtime int32
	binary.Read(reader, binary.LittleEndian, &mtime)
	module.mtime = time.Unix(int64(mtime), 0)
	binary.Read(reader, binary.LittleEndian, &module.Size)
	module.Code = reader.ReadCode()
}
