package gothon

import (
	"encoding/binary"
	"container/list"
)

type List struct {
	list.List
}

func (this *List) Read(reader *Reader) {
	var size int32
	binary.Read(reader, binary.LittleEndian, &size)
	
	this.List = *list.New()

	for this.Len() < int(size) {
		tmp := reader.ReadObject()
		this.PushBack(tmp)
	}
}
