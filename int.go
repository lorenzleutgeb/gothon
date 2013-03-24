package gothon

import (
	"encoding/binary"
)

type Int struct {
	int32
}

func (this *Int) Read(reader *Reader) {
	binary.Read(reader, binary.LittleEndian, &this.int32)
}
