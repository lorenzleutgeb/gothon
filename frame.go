package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Frame is a code object in execution context.
// It has a data stack and a block stack.
type Frame struct {
	stack  []Object
	blocks []block
	code   *Code
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
	fmt.Printf("\x1b[32;1m%s\x1b[0m\n", f.code.Name.string)

	f.blocks = make([]block, 0)

	pc := 0
	var op, first, second byte

	for pc < len(f.code.Instructions) {
		op = f.code.Instructions[pc]

		fmt.Printf("\x1b[0;32m%4d %-17s", pc, opcode[op])

		if hasArg(op) {
			first, second = f.code.Instructions[pc+1], f.code.Instructions[pc+2]
			fmt.Printf(" %3d %3d", first, second)
			pc += 2
		} else {
			first, second = 0, 0
			fmt.Printf("        ")
		}

		fmt.Printf(" Ϟ %s\x1b[0m\n", f.stack)

		fmt.Printf("%d\n", pc%3)
		//fmt.Printf("%d/%d - %x\n", pc, (len(code.Instructions)), md5.Sum(code.Instructions))

		pc++

		switch op {
		case POP_TOP:
			f.Pop()
		case LOAD_CONST:
			f.Push((*f.code.Consts)[first])
		case STORE_NAME:
			(*f.code.Names)[first] = f.Pop()
		case LOAD_NAME:
			f.Push((*f.code.Names)[first])
		case LOAD_FAST:
			f.Push((*f.code.Varnames)[first])
		case MAKE_FUNCTION:
			if fqn, ok := f.Pop().(*String); ok {
				if code, ok := f.Pop().(*Code); ok {
					result := &Function{
						Name: fqn,
						Code: code,
					}
					f.Push(result)
				}
			}
		case CALL_FUNCTION:
			if second > 0 {
				panic("Keyword parameters are not implemented.")
			}

			args := make(Tuple, int(first))

			for j := int(first) - 1; j > -1; j-- {
				args[j] = f.Pop()
			}

			o := f.Pop()

			if function, ok := o.(*Function); ok {
				function.Code.Varnames = &args

				invoc := NewFrame(function.Code)

				f.Push(invoc.Execute())
			} else if s, ok := o.(*String); ok {
				if s.string == "print" { // print is built-in
					for _, arg := range args {
						if str, ok := arg.(*String); ok {
							fmt.Printf("%s", str.string)
						} else {
							fmt.Printf("\x1b[31;1mPrinting something that's not a string:\x1b[0m %+v", arg)
						}
					}
					fmt.Println()
					f.Push(&None{})
				} else if s.string == "set" {
					// TODO(flowlo): Implement iterating over args
					f.Push(&Set{})
				} else if s.string == "input" {
					if len(args) > 0 {
						if str, ok := args[0].(*String); ok {
							fmt.Printf("%s", str.string)
						} else {
							fmt.Printf("\x1b[31;1mPrinting something that's not a string:\x1b[0m %+v", args[0])
						}
					}
					reader := bufio.NewReader(os.Stdin)
					text, _ := reader.ReadString('\n')
					text = strings.TrimRight(text, "\r\n")
					f.Push(&String{text})
				} else if s.string == f.code.Name.string { // Recursive call
					invoc := NewFrame(f.code)
					invoc.code.Varnames = &args
					f.Push(invoc.Execute())
				} else {
					panic(fmt.Sprintf("Unknown function: %s", s))
				}
			}

		case BINARY_MULTIPLY: // TODO(flowlo): implement this for floats
			if a, ok := f.Pop().(*Int); ok {
				if b, ok := f.Pop().(*Int); ok {
					f.Push(&Int{a.int32 * b.int32})
				}
			}
		case BINARY_ADD: // TODO(flowlo): implement this for floats
			if a, ok := f.Pop().(*Int); ok {
				if b, ok := f.Pop().(*Int); ok {
					f.Push(&Int{a.int32 + b.int32})
				}
			}
		case RETURN_VALUE:
			return f.Pop()

		case LOAD_GLOBAL:
			f.Push((*f.code.Names)[first])

		case COMPARE_OP:
			rightx := f.Pop()
			leftx := f.Pop()
			if right, ok := rightx.(*Int); ok {
				if left, ok := leftx.(*Int); ok {
					switch first {
					case OP_LT:
						if left.int32 < right.int32 {
							f.Push(&True{})
						} else {
							f.Push(&False{})
						}
					case OP_LEQ:
						if left.int32 <= right.int32 {
							f.Push(&True{})
						} else {
							f.Push(&False{})
						}

					case OP_EQ:
						if left.int32 == right.int32 {
							f.Push(&True{})
						} else {
							f.Push(&False{})
						}
					case OP_GT:
						if left.int32 > right.int32 {
							f.Push(&True{})
						} else {
							f.Push(&False{})
						}
					case OP_GE:
						if left.int32 >= right.int32 {
							f.Push(&True{})
						} else {
							f.Push(&False{})
						}
					default:
						panic("Comparison operator not implemented.")
					}
				}
			} else if right, ok := rightx.(*String); ok {
				if left, ok := leftx.(*String); ok {
					if first == OP_EQ || first == OP_IS {
						if left.string == right.string {
							f.Push(&True{})
						} else {
							f.Push(&False{})
						}
					} else if first == OP_ISNT {
						if left.string != right.string {
							f.Push(&True{})
						} else {
							f.Push(&False{})
						}
					} else {
						panic(fmt.Sprintf("Unimplemented operator: %d", int(first)))
					}
				}
			} else {
				panic(fmt.Sprintf("Cannot compare %T and %T", rightx, leftx))
			}

		case POP_JUMP_IF_FALSE:
			o := f.Pop()

			if _, ok := o.(*False); ok {
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
		//case DUP_TOP:
		//case DUP_TOP_TWO:
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
			name := (*f.code.Names)[first]
			fromlist := f.Pop()
			level := f.Pop()
			fmt.Printf("import %s with %s and %s\n", name, fromlist, level)

			if s, ok := name.(*String); ok {
				if s.string == "sys" {
					f.Push(NewSys())
				}
			}
		case LOAD_ATTR:
			name := (*f.code.Names)[first]
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
			//f.Push(&String{"NOT IMPLEMENTED"})
		case JUMP_ABSOLUTE:
			pc = int(first)
		case SETUP_LOOP:
			block := &block{pc, pc + int(first)}
			f.blocks = append(f.blocks, *block)
		case POP_BLOCK:
			f.blocks = f.blocks[:len(f.blocks)-1]
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
	f.stack = make([]Object, code.Stacksize)
	// TODO(flowlo): What's a good capacity here?
	f.blocks = make([]block, 1)
	f.code = code
	return f
}
