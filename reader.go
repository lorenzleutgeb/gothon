package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
)

type Reader struct {
	bufio.Reader
	Module Module
}

func (reader *Reader) ReadObject() Object {
	c, err := reader.ReadByte()

	if err != nil {
		log.Fatal(err)
	}

	var result Object

	fmt.Printf("\x1b[34;1m%c\x1b[0m", c&0x7f)

	if c == 'r' {
		var index int32
		binary.Read(reader, binary.LittleEndian, &index)
		fmt.Printf("%d", index)
		return reader.Module.GetReference(int(index))
	}

	switch c & 0x7f {
	case TYPE_NULL:
		result = &Null{}
	case TYPE_NONE:
		result = &None{}
	case TYPE_FALSE:
		result = &False{}
	case TYPE_TRUE:
		result = &True{}
		/*	case 'S':
				result := &StopIter{}
			case '.'
				result := &Ellipsis{} */
	case TYPE_INT:
		result = &Int{}
		/*	case 'I':
				result = &Int64{}
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
	case TYPE_SHORT_ASCII_INTERNED:
		fallthrough
	case TYPE_SHORT_ASCII:
		fallthrough
	case TYPE_STRING:
		result = &String{}
		//	case 't': result = &Interned{}
		//	case 'R': result = &StringRef{}
	case TYPE_TUPLE:
		result = &Tuple{}
	case TYPE_SMALL_TUPLE:
		result = &Tuple{}
	case TYPE_LIST:
		result = &List{}
	case TYPE_DICT:
		result = &Dictionary{}
	case TYPE_CODE:
		result = &Code{}
	case TYPE_UNICODE:
		result = &String{}
		/*	case '?'
			result := &Unknown{} */
	case TYPE_SET:
		result = &Set{}
		/*	case '>'
			result := &FrozenSet{} */
	default:
		panic(fmt.Sprintf("Reached bad type specifier '%c'", c))
	}

	// TODO(flowlo): Prevent references to null by checking here
	// before adding a reference
	if c&0x80 != 0 {
		fmt.Print("+")
		reader.Module.AddReference(result)
		//fmt.Printf("After adding reference %p: %+v\n", result, reader.Module)
	}

	// c needs masking here, references are already stored!
	result.Read(reader, c&0x7f)

	return result
}

func (reader *Reader) ReadTuple() *Tuple {
	tmp := reader.ReadObject()
	if value, ok := tmp.(*Tuple); ok {
		return value
	}
	return nil
}

func (reader *Reader) ReadString() *String {
	tmp := reader.ReadObject()
	if value, ok := tmp.(*String); ok {
		return value
	}
	return nil
}

func (reader *Reader) ReadCode() *Code {
	tmp := reader.ReadObject()
	if value, ok := tmp.(*Code); ok {
		return value
	}
	return nil
}
