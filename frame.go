package main

import (
	"fmt"
	"os"
)

// Frame is a code object in execution context.
// It has a data stack and a block stack.
type Frame struct {
	stack  []Object
	blocks []block
	code   *Code
	names  map[string]Object
}

type block struct {
	start, end int
}

func hasArg(opcode byte) bool {
	return opcode > 89
}

// Push pushes an item to the stack associated
// with the frame.
func (f *Frame) Push(item Object) {
	f.stack = append(f.stack, item)
}

// Pop pops an item to the stack associated
// with the frame.
func (f *Frame) Pop() Object {
	item := f.stack[len(f.stack)-1]
	f.stack = f.stack[:len(f.stack)-1]
	return item
}

// Peek gives you what's on top of the stack,
// but does not remove the item.
func (f *Frame) Peek() Object {
	return f.stack[len(f.stack)-1]
}

// Execute brings the frame to execution. The attached
// code will be executed starting at the first instruction
// and interpreted.
func (f *Frame) Execute() Object {
	if *debug {
		fmt.Println()
	}

	f.blocks = make([]block, 0)

	pc := 0
	var op, first, second byte

	for pc < len(f.code.Instructions) {
		op = f.code.Instructions[pc]

		if *debug {
			fmt.Printf("\x1b[0;32m%4d %-17s", pc, opcode[op])
		}

		if hasArg(op) {
			first, second = f.code.Instructions[pc+1], f.code.Instructions[pc+2]
			if *debug {
				fmt.Printf(" %3d %3d", first, second)
			}
			pc += 2
		} else {
			first, second = 0, 0
			if *debug {
				fmt.Printf("        ")
			}
		}

		if *debug {
			fmt.Printf(" %s\x1b[0m\n", f.stack)
		}

		pc++

		switch op {
		case POP_TOP:
			f.Pop()
		case LOAD_CONST:
			f.Push(f.code.Consts[first])
		case STORE_NAME:
			f.code.Names[first] = f.Pop()
		case LOAD_NAME:
			f.Push(f.code.Names[first])
		case LOAD_FAST:
			name := f.code.Varnames[int(first)].(String).string
			f.Push(f.names[name])
		case MAKE_FUNCTION:
			name := f.Pop().(String)
			code := f.Pop().(Code)
			f.Push(NewExternalFunction(name.string, &code))
		case CALL_FUNCTION:
			if second > 0 {
				panic("Keyword parameters are not implemented.")
			}

			pos := make([]Object, int(first))

			for j := int(first) - 1; j > -1; j-- {
				pos[j] = f.Pop()
			}

			args := args{Positional: pos}

			o := f.Pop()

			if function, ok := o.(Function); ok {
				f.Push(function.Call(&args))
			} else if fname, ok := o.(String); ok {
				value, isBuiltin := builtin[fname.string]

				if isBuiltin {
					f.Push(value)
				} else {
					panic("can only resolve named functions to builtins")
				}

				function := value.(Function)
				f.Push(function.Call(&args))
			} else {
				fmt.Fprintf(os.Stderr, "%+v", o)
				panic("unknown function call")
			}
		case BINARY_MULTIPLY: // TODO(flowlo): implement this for floats
			if a, ok := f.Pop().(Int); ok {
				if b, ok := f.Pop().(Int); ok {
					f.Push(Int{a.int32 * b.int32})
				}
			}
		case BINARY_ADD: // TODO(flowlo): implement this for floats
			if a, ok := f.Pop().(Int); ok {
				if b, ok := f.Pop().(Int); ok {
					f.Push(Int{a.int32 + b.int32})
				}
			}
		case RETURN_VALUE:
			return f.Pop()
		case LOAD_GLOBAL:
			name := f.code.Names[first].(String).string

			value, isBuiltin := builtin[name]

			if isBuiltin {
				f.Push(value)
			} else {
				panic("lookup of globals other than builtins not implemented")
			}
		case COMPARE_OP:
			rightx := f.Pop()
			leftx := f.Pop()
			if right, ok := rightx.(Int); ok {
				if left, ok := leftx.(Int); ok {
					switch first {
					case OP_LT:
						if left.int32 < right.int32 {
							f.Push(True{})
						} else {
							f.Push(False{})
						}
					case OP_LEQ:
						if left.int32 <= right.int32 {
							f.Push(True{})
						} else {
							f.Push(False{})
						}

					case OP_EQ:
						if left.int32 == right.int32 {
							f.Push(True{})
						} else {
							f.Push(False{})
						}
					case OP_GT:
						if left.int32 > right.int32 {
							f.Push(True{})
						} else {
							f.Push(False{})
						}
					case OP_GE:
						if left.int32 >= right.int32 {
							f.Push(True{})
						} else {
							f.Push(False{})
						}
					default:
						panic("comparison operator not implemented.")
					}
				}
			} else if right, ok := rightx.(String); ok {
				if left, ok := leftx.(String); ok {
					if first == OP_EQ || first == OP_IS {
						if left.string == right.string {
							f.Push(True{})
						} else {
							f.Push(False{})
						}
					} else if first == OP_ISNT {
						if left.string != right.string {
							f.Push(True{})
						} else {
							f.Push(False{})
						}
					} else {
						panic(fmt.Sprintf("unimplemented operator: %d", int(first)))
					}
				}
			} else {
				panic(fmt.Sprintf("cannot compare %T and %T", rightx, leftx))
			}

		case POP_JUMP_IF_FALSE:
			if _, ok := f.Pop().(False); ok {
				pc = int(first) + int(second)*256
			}
		case BINARY_SUBTRACT: // TODO(flowlo): Implement this for floats
			if right, ok := f.Pop().(*Int); ok {
				if left, ok := f.Pop().(*Int); ok {
					f.Push(&Int{left.int32 - right.int32})
				}
			}
		case NOP:
		case ROT_TWO:
			a := f.Pop()
			b := f.Pop()
			f.Push(a)
			f.Push(b)
		case ROT_THREE:
			a := f.Pop()
			b := f.Pop()
			c := f.Pop()
			f.Push(b)
			f.Push(a)
			f.Push(c)
		case UNARY_POSITIVE: // TODO(flowlo): Implement for floats
			o := f.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = +a.int32
			}
		case UNARY_NEGATIVE: // TODO(flowlo): Implement for floats
			o := f.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = -a.int32
			}
		case UNARY_NOT: // TODO(flowlo): Implement for floats
			o := f.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = ^a.int32
			}
		case UNARY_INVERT: // TODO(flowlo): Implement for floats
			o := f.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = ^a.int32
			}
		case IMPORT_NAME:
			name := f.code.Names[first]
			fromlist := f.Pop()
			level := f.Pop()
			fmt.Printf("import %s with %s and %s\n", name, fromlist, level)

			imp := builtin["__import__"].(Function)
			arg := &args{
				Positional: []Object{name},
			}
			f.Push(imp.Call(arg))
		case LOAD_ATTR:
			name := f.code.Names[first]
			o := f.Pop()

			ao, ok := o.(*Code)
			if !ok {
				panic(fmt.Sprintf("TypeError: %v of type %T has no attributes!", o, o))
			}
			a, err := ao.GetAttribute(name, nil)
			if err != nil {
				panic(err.Error())
			}
			f.Push(a)
		case JUMP_ABSOLUTE:
			pc = int(first)
		case SETUP_LOOP:
			block := &block{pc, pc + int(first)}
			f.blocks = append(f.blocks, *block)
		case POP_BLOCK:
			f.blocks = f.blocks[:len(f.blocks)-1]
		case STORE_FAST:
			name := f.code.Varnames[int(first)].(String).string
			f.names[name] = f.Pop()
		default:
			fmt.Sprintf("\x1b[31;1mSkipped\x1b[0m unknown opcode: %d\n", op)
			panic(fmt.Sprintf("Unknown opcode: %d", op))
		}
	}
	return &Null{}
}

// NewFrame constructs a new Frame to execute the passed code. This
// allocates memory for the execution stack and sets everything
// up.
func NewFrame(code *Code) *Frame {
	f := new(Frame)
	f.stack = make([]Object, 0)
	// TODO(flowlo): What's a good capacity here?
	f.blocks = make([]block, 0)
	f.code = code
	f.names = make(map[string]Object, 0)
	return f
}
