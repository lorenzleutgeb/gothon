package gothon
 
type Stack struct {
	top *Element
	size int
}
 
type Element struct {
	value Object
	next *Element
}

func (s *Stack) Len() int {
	return s.size
}
 
func (s *Stack) Push(value Object) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value Object) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}
