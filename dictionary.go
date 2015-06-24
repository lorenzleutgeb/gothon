package main

type Dictionary map[Object]Object

func (dict *Dictionary) Read(reader *Reader, t byte) {
	*dict = make(map[Object]Object)

	for {
		key := reader.ReadObject()
		if _, ok := key.(*Null); ok {
			break
		}
		value := reader.ReadObject()
		(*dict)[key] = value
	}
}
