package gothon

import (
	"encoding/binary"
)

type Tuple []Object

func (this *Tuple) Read(reader *Reader) {
	var size int32
	binary.Read(reader, binary.LittleEndian, &size)
	
	*this = make([]Object, size)

	for i := 0; i < int(size); i++ {
		(*this)[i] = reader.ReadObject()
	}
}
