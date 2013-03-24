package pyc

import (
	"bufio"
	"reflect"
	"log"
)

type Reader struct {
	bufio.Reader
}

func (reader *Reader) ReadObject() Object {
	c, err := reader.ReadByte()
	
	if err != nil {
		log.Fatal(err)
	}
	
	var result Object
	
	switch c {
	case '0':
		result = &Null{}
	case 'N':
		result = &None{}
	case 'F':
		result = &False{}
	case 'T':
		result = &True{}
/*	case 'S':
		result := &StopIter{}
	case '.'
		result := &Ellipsis{}
	case 'i'
		result := &Int{}
	case 'f'
		result := &Float{}
	case 'g'
		result := &BinaryFloat{}
	case 'x'
		result := &Complex{}
	case 'y'
		result := &BinaryComplex{}
	case 'l'
		result := Long */
	case 's':
		result = &String{}
	case '(':
		result = &Tuple{}
	case '[':
		result = &List{}
	case '{':
		result = &Dictionary{}
	case 'c':
		result = &Code{}
/*	case 'u'
		result := &Unicode{}
	case '?'
		result := &Unknown{} */
	case '<':
		result = &Set{}
/*	case '>'
		result := &FrozenSet{} */
	default:
		log.Fatalf("Read unknown type specifier byte \"%x\"", c)
	}
	
	return result
}

func (reader *Reader) ReadExpected(object Object) {
	tmp := reader.ReadObject()

	if reflect.TypeOf(tmp) == reflect.TypeOf(object) {
		object.Read(reader)
		return
	}

	log.Fatalf("Failed type assertion. gothon.pyc.Reader.ReadExpected read %T  but expected %T", tmp, object)
}
