package gothon

import (
	"bufio"
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
		result := &Ellipsis{} */
	case 'i':
		result = &Int{}
/*	case 'f'
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
	case 'u':
		result = &String{}
/*	case '?'
		result := &Unknown{} */
	case '<':
		result = &Set{}
/*	case '>'
		result := &FrozenSet{} */
	default:
		log.Fatalf("Read unknown type specifier byte \"%c\"", c)
	}
	result.Read(reader)
	return result
}

func (reader *Reader) ReadTuple() *Tuple {
	tmp := reader.ReadObject()
	if value, ok := tmp.(*Tuple) ; ok {
		return value
	}
	return nil
}

func (reader *Reader) ReadString() *String {
	tmp := reader.ReadObject()
	if value, ok := tmp.(*String) ; ok {
		return value
	}
	return nil
}

func (reader *Reader) ReadCode() *Code {
	tmp := reader.ReadObject()
	if value, ok := tmp.(*Code) ; ok {
		return value
	}
	return nil
}
