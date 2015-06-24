package main

import (
	"encoding/binary"
	"encoding/json"
)

type String struct {
	string
}

func (s *String) Read(reader *Reader, t byte) {
	var size int32
	// short ASCII strings
	if t == 'Z' || t == 'z' {
		// TODO(flowlo): Handle error
		tmp, _ := reader.ReadByte()
		size = int32(tmp)
	} else {
		binary.Read(reader, binary.LittleEndian, &size)
	}
	var result = make([]byte, size)
	reader.Read(result)
	s.string = string(result)
	//fmt.Printf("Just read string of length %d\n", size)
}

func (s *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.string)
}

func (s *String) String() string {
	return "\"" + s.string + "\""
}
