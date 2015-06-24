package main

type None struct {
}

func (none *None) Read(reader *Reader, t byte) {
	return
}

func (none *None) String() string {
	return "None"
}
