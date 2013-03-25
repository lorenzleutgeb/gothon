package gothon

type None struct {
}

func (none *None) Read(reader *Reader) {
	return
}

func (none *None) String() string {
	return "None"
}
