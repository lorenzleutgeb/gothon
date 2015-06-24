package main

import (
	"encoding/binary"
)

type Set map[Object]struct{}

func (set *Set) Read(reader *Reader, t byte) {
	var size int32
	binary.Read(reader, binary.LittleEndian, &size)

	*set = make(map[Object]struct{}, size)

	for i := 0; i < int(size); i++ {
		item := reader.ReadObject()
		(*set)[item] = *new(struct{})
	}
}
