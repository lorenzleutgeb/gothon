package gothon

type True struct {
	bool
}

func (this *True) Read(reader *Reader) {
	this.bool = true
	return
}
