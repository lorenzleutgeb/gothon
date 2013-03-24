package pyc

import (
	"encoding/binary"
)

type String struct {
	string
}

func (this *String) Read(reader *Reader) {
	var size int32
	binary.Read(reader, binary.LittleEndian, &size)
	var result = make([]byte, size)
	reader.Read(result)
	this.string = string(result)
}
