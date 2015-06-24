package main

import (
	"container/list"
	"encoding/binary"
)

type List struct {
	list.List
}

func (l *List) Read(reader *Reader, t byte) {
	var size int32
	binary.Read(reader, binary.LittleEndian, &size)

	l.List = *list.New()

	for l.Len() < int(size) {
		tmp := reader.ReadObject()
		l.PushBack(tmp)
	}
}
