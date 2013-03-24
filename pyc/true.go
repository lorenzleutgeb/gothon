package pyc

type True struct {
	bool
}

func (this *True) Read(reader *Reader) {
	this.bool = true
	return
}
