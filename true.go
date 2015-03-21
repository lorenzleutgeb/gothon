package gothon

type True struct {
	bool
}

func (this *True) Read(reader *Reader, t byte) {
	this.bool = true
	return
}
