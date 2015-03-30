package gothon

import (
	"encoding/binary"
	"fmt"
)

type Int struct {
	int32
}

func (i *Int) Read(reader *Reader, t byte) {
	binary.Read(reader, binary.LittleEndian, &i.int32)
}

func (i *Int) String() string {
	return fmt.Sprintf("%d", i.int32)
}
