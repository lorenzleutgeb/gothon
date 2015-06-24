package main

type Null struct {
}

func (null *Null) Read(reader *Reader, t byte) {
	return
}
