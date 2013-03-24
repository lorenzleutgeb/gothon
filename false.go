package gothon

type False struct {
	bool
}

func (this *False) Read(reader *Reader) {
	this.bool = false
}
