package main

import (
	"container/list"
	"fmt"
)

type Null struct{}

func (n Null) String() string {
	return "null"
}

type None struct{}

func (none None) String() string {
	return "None"
}

type False struct {
	bool
}

func (f False) String() string {
	return "False"
}

type True struct {
	bool
}

func (t True) String() string {
	return "True"
}

type String struct {
	string
}

func (s String) String() string {
	return "\"" + s.string + "\""
}

type Tuple []Object

func (t Tuple) String() string {
	result := "<tuple"
	for _, v := range t {
		result = result + " " + v.String()
	}
	return result + ">"
}

type Dictionary map[Object]Object

func (d Dictionary) String() string {
	return "<dict>"
}

type Set map[Object]struct{}

func (s Set) String() string {
	return "<set>"
}

type Int struct {
	int32
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.int32)
}

type List struct {
	list.List
}

func (l List) String() string {
	return "<list>"
}
