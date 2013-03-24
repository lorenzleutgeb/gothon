package pyc

import (
	"log"
	"time"
	"encoding/binary"
	"container/list"
)

type Module struct {
	code Code
	interns list.List
	mtime time.Time
	size uint32
	version uint16
}

func (module *Module) Read(reader *Reader) {
	binary.Read(reader, binary.LittleEndian, &module.version)
	var check uint16
	binary.Read(reader, binary.LittleEndian, &check)

	if check != 0xa0d {
		log.Fatalf("Second two bytes of magic number are incorrect (0x0a0d expected, 0x%x found)", check)
	}

	var mtime int32
	binary.Read(reader, binary.LittleEndian, &mtime)
	module.mtime = time.Unix(int64(mtime), 0)
	binary.Read(reader, binary.LittleEndian, &module.size)
	reader.ReadExpected(&module.code)
}
