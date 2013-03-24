package pyc

import (
	"encoding/binary"
)

type Set map[Object]struct{}

func (this *Set) Read(reader *Reader) {
	var size int32
	binary.Read(reader, binary.LittleEndian, &size)
	
	*this = make(map[Object]struct{}, size)

	for i := 0 ; i < int(size) ; i++ {
		item := reader.ReadObject()
		(*this)[item] = *new(struct{})
	}
}
