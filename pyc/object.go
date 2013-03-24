package pyc

type Object interface {
	Read(*Reader)
}
