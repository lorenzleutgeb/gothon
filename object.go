package gothon

type Object interface {
	Read(*Reader, byte)
	//	MarshalJSON() (byte, error)
}
