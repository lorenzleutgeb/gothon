package gothon

import (
	"encoding/binary"
	"encoding/json"
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

func (this *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.string)
}

func (this *String) String() string {
	return "\"" + this.string + "\""
}
