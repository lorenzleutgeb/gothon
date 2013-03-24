package gothon

type Dictionary map[Object]Object

func (this *Dictionary) Read(reader *Reader) {
	*this = make(map[Object]Object)

	for {
		key := reader.ReadObject()
		if _, ok := key.(*Null) ; ok { break }
		value := reader.ReadObject()
		(*this)[key] = value
	}
}
