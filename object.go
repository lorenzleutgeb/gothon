package main

import (
	"fmt"
)

type Object interface {
	Read(*Reader, byte)
	//	MarshalJSON() (byte, error)
}

type AttributedObject struct {
	attr map[string]Object
}

// Constructs an AttributedObject without any
// attributes set
func NewAttributedObject() *AttributedObject {
	ao := new(AttributedObject)
	ao.attr = make(map[string]Object)
	return ao
}

// Retrieves an attribute
func (ao *AttributedObject) GetAttribute(name, fallback Object) (Object, error) {
	key, ok := name.(*String)
	if !ok {
		panic("Can only retrieve attributes by string.")
	}
	value, found := ao.attr[key.string]
	if found {
		return value, nil
	}
	if fallback != nil {
		return fallback, nil
	}
	return nil, fmt.Errorf("AttributeError: Unknown object has no attribute '%s'", key.string)
}

func (ao *AttributedObject) AddAttribute(key string, value Object) {
	if ao.attr == nil {
		ao.attr = make(map[string]Object)
	}
	ao.attr[key] = value
}

func (ao AttributedObject) Read(r *Reader, t byte) {
	panic("AttributeObject.Read is a dummy!")
}
