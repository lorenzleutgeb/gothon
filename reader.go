package main

import (
	"bufio"
	"container/list"
	"encoding/binary"
	"fmt"
	"log"
)

type Reader struct {
	bufio.Reader
	module *Module
}

func (r *Reader) ReadObject() Object {
	c, err := r.ReadByte()

	if err != nil {
		log.Fatal(err)
	}

	if *debug {
		fmt.Printf("\x1b[34;1m%c\x1b[0m", c&0x7f)
	}

	if c == TYPE_OBREF {
		var index int32
		binary.Read(r, binary.LittleEndian, &index)
		if *debug {
			fmt.Printf("%d", index)
		}
		return *r.module.References[index-1]
	}

	o := r.createObject(c & 0x7f)

	// TODO(flowlo): Prevent references to null by checking here
	// before adding a reference
	if c&0x80 != 0 {
		r.module.References = append(r.module.References, &o)
	}

	return o
}

func (r *Reader) createObject(c byte) Object {
	var result Object

	switch c {
	case TYPE_NULL:
		result = &Null{}
	case TYPE_NONE:
		result = &None{}
	case TYPE_FALSE:
		result = &False{}
	case TYPE_TRUE:
		result = &True{}
	case TYPE_INT:
		result = r.readInt()
	case TYPE_STRING, TYPE_INTERNED, TYPE_STRINGREF, TYPE_ASCII, TYPE_ASCII_INTERNED, TYPE_SHORT_ASCII, TYPE_SHORT_ASCII_INTERNED, TYPE_UNICODE:
		result = r.readString(c)
	case TYPE_TUPLE, TYPE_SMALL_TUPLE:
		result = r.readTuple(c)
	case TYPE_LIST:
		result = r.readList()
	case TYPE_DICT:
		result = r.readDict()
	case TYPE_CODE:
		result = r.readCode()
	case TYPE_SET, TYPE_FROZENSET:
		result = r.readSet(c)
	default:
		panic("bad type specifier")
	}
	return result
}
func (r *Reader) readTuple(t byte) Tuple {
	size := 0
	if t == TYPE_TUPLE {
		var raw int32
		binary.Read(r, binary.LittleEndian, &raw)
		size = int(raw)
	} else if t == TYPE_SMALL_TUPLE {
		var raw int8
		binary.Read(r, binary.LittleEndian, &raw)
		size = int(raw)
	} else {
		panic("unkown tuple type")
	}

	tuple := make(Tuple, size)

	for i := 0; i < int(size); i++ {
		tuple[i] = r.ReadObject()
	}
	return tuple
}

func (r *Reader) readList() List {
	var size int32
	binary.Read(r, binary.LittleEndian, &size)

	result := *list.New()

	for result.Len() < int(size) {
		tmp := r.ReadObject()
		result.PushBack(tmp)
	}

	return List{List: result}
}

func (r *Reader) readDict() Dictionary {
	dict := make(map[Object]Object)

	for {
		key := r.ReadObject()
		if _, ok := key.(*Null); ok {
			break
		}
		dict[key] = r.ReadObject()
	}

	return dict
}

func (r *Reader) readString(t byte) String {
	var size int32

	if t == TYPE_STRINGREF {
		binary.Read(r, binary.LittleEndian, &size)
		return *r.module.Interns[int(size)]
	} else if t == TYPE_SHORT_ASCII || t == TYPE_SHORT_ASCII_INTERNED {
		// TODO(flowlo): Handle error
		tmp, _ := r.ReadByte()
		size = int32(tmp)
	} else if t == TYPE_STRING || t == TYPE_INTERNED || t == TYPE_UNICODE {
		binary.Read(r, binary.LittleEndian, &size)
	} else {
		panic("unknown string type")
	}
	var result = make([]byte, size)
	r.Read(result)

	s := String{string: string(result)}

	if t == TYPE_INTERNED || t == TYPE_ASCII_INTERNED || t == TYPE_SHORT_ASCII_INTERNED {
		r.module.Interns = append(r.module.Interns, &s)
	}

	return s
}

func (r *Reader) readCode() (code Code) {
	binary.Read(r, binary.LittleEndian, &code.PosArgCnt)
	binary.Read(r, binary.LittleEndian, &code.KwArgCnt)
	binary.Read(r, binary.LittleEndian, &code.Nlocals)
	binary.Read(r, binary.LittleEndian, &code.Stacksize)
	binary.Read(r, binary.LittleEndian, &code.Flags)

	c, _ := r.ReadByte()

	if c != TYPE_STRING {
		panic(fmt.Sprintf("Expected code block but got anything else (\"%s\")", string(c)))
	}

	var size int32
	binary.Read(r, binary.LittleEndian, &size)

	code.Instructions = make([]byte, size)
	r.Read(code.Instructions)

	code.Consts = r.ReadObject().(Tuple)
	code.Names = r.ReadObject().(Tuple)
	code.Varnames = r.ReadObject().(Tuple)
	code.Freevars = r.ReadObject().(Tuple)
	code.Cellvars = r.ReadObject().(Tuple)
	code.Filename = r.ReadObject().(String)
	code.Name = r.ReadObject().(String)
	binary.Read(r, binary.LittleEndian, &code.Firstlineno)
	code.Lnotab = r.ReadObject().(String)
	return
}

func (r *Reader) readSet(t byte) (set Set) {
	if t == TYPE_FROZENSET {
		panic("frozen sets are not implemented")
	} else if t != TYPE_SET {
		panic("unknown set type")
	}

	var size int32
	binary.Read(r, binary.LittleEndian, &size)

	set = make(map[Object]struct{}, size)

	for i := 0; i < int(size); i++ {
		item := r.ReadObject()
		set[item] = *new(struct{})
	}
	return
}

func (r *Reader) readInt() (i Int) {
	binary.Read(r, binary.LittleEndian, &i.int32)
	return
}

func NewReader(br bufio.Reader) Reader {
	return Reader{
		br,
		&Module{},
	}
}
