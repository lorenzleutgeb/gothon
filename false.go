package gothon

type False struct {
	bool
}

func (this *False) Read(reader *Reader, t byte) {
	this.bool = false
}
