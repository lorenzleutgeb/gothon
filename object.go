package gothon

type Object interface {
	Read(*Reader)
//	MarshalJSON() (byte, error)
}
