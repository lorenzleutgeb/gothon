package gothon

type False struct {
	bool
}

func (f *False) Read(reader *Reader, t byte) {
	f.bool = false
}
