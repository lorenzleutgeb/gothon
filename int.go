package gothon

import (
	"encoding/binary"
	"fmt"
)

type Int struct {
	int32
}

func (this *Int) Read(reader *Reader, t byte) {
	binary.Read(reader, binary.LittleEndian, &this.int32)
}

func (this *Int) String() string {
	return fmt.Sprintf("%d", this.int32)
}
