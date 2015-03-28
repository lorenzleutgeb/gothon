package gothon

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Frame struct {
	stack  *Stack
	blocks []Block
}

type Block struct {
	start, end int
}

func HasArg(opcode byte) bool {
	return opcode > 89
}

func (frame *Frame) Run(code Code) Object {
	fmt.Printf("\x1b[32;1m%s\x1b[0m\n", code.Name.string)

	frame.blocks = make([]Block, 0)

	pc := 0
	var op, first, second byte

	for pc < len(code.Instructions) {
		op = code.Instructions[pc]

		fmt.Printf("\x1b[0;32m%4d %-17s", pc, opcode[op])

		if HasArg(op) {
			first, second = code.Instructions[pc+1], code.Instructions[pc+2]
			fmt.Printf(" %3d %3d", first, second)
			pc += 2
		} else {
			first, second = 0, 0
			fmt.Printf("        ")
		}

		fmt.Printf(" Ϟ %s\x1b[0m\n", frame.stack)

		fmt.Printf("%d\n", pc%3)
		//fmt.Printf("%d/%d - %x\n", pc, (len(code.Instructions)), md5.Sum(code.Instructions))

		pc++

		switch op {
		case POP_TOP:
			frame.stack.Pop()
		case LOAD_CONST:
			frame.stack.Push((*code.Consts)[first])
		case STORE_NAME:
			(*code.Names)[first] = frame.stack.Pop()
		case LOAD_NAME:
			frame.stack.Push((*code.Names)[first])
		case LOAD_FAST:
			frame.stack.Push((*code.Varnames)[first])
		case MAKE_FUNCTION:
			if fqn, ok := frame.stack.Pop().(*String); ok {
				if code, ok := frame.stack.Pop().(*Code); ok {
					result := &Function{
						Name: fqn,
						Code: code,
					}
					frame.stack.Push(result)
				}
			}
		case CALL_FUNCTION:
			if second > 0 {
				panic("Keyword parameters are not implemented.")
			}

			args := make(Tuple, int(first))

			for j := int(first) - 1; j > -1; j-- {
				args[j] = frame.stack.Pop()
			}

			o := frame.stack.Pop()

			if function, ok := o.(*Function); ok {
				function.Code.Varnames = &args
				frame.stack.Push(frame.Run(*function.Code))
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
				} else if s.string == "set" {
					// TODO(flowlo): Implement iterating over args
					frame.stack.Push(&Set{})
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
					frame.stack.Push(&String{text})
				} else if s.string == code.Name.string {
					code.Varnames = &args
					frame.stack.Push(frame.Run(code))
				} else {
					panic(fmt.Sprintf("Unknown function: %s", s))
				}
			}

		case BINARY_MULTIPLY: // TODO(flowlo): implement this for floats
			if a, ok := frame.stack.Pop().(*Int); ok {
				if b, ok := frame.stack.Pop().(*Int); ok {
					frame.stack.Push(&Int{a.int32 * b.int32})
				}
			}
		case BINARY_ADD: // TODO(flowlo): implement this for floats
			if a, ok := frame.stack.Pop().(*Int); ok {
				if b, ok := frame.stack.Pop().(*Int); ok {
					frame.stack.Push(&Int{a.int32 + b.int32})
				}
			}
		case RETURN_VALUE:
			return frame.stack.Pop()

		case LOAD_GLOBAL:
			frame.stack.Push((*code.Names)[first])

		case COMPARE_OP:
			rightx := frame.stack.Pop()
			leftx := frame.stack.Pop()
			if right, ok := rightx.(*Int); ok {
				if left, ok := leftx.(*Int); ok {
					switch first {
					case OP_LT:
						if left.int32 < right.int32 {
							frame.stack.Push(&True{})
						} else {
							frame.stack.Push(&False{})
						}
					case OP_LEQ:
						if left.int32 <= right.int32 {
							frame.stack.Push(&True{})
						} else {
							frame.stack.Push(&False{})
						}

					case OP_EQ:
						if left.int32 == right.int32 {
							frame.stack.Push(&True{})
						} else {
							frame.stack.Push(&False{})
						}
					case OP_GT:
						if left.int32 > right.int32 {
							frame.stack.Push(&True{})
						} else {
							frame.stack.Push(&False{})
						}
					case OP_GE:
						if left.int32 >= right.int32 {
							frame.stack.Push(&True{})
						} else {
							frame.stack.Push(&False{})
						}
					default:
						panic("Comparison operator not implemented.")
					}
				}
			} else if right, ok := rightx.(*String); ok {
				if left, ok := leftx.(*String); ok {
					if first == OP_EQ || first == OP_IS {
						if left.string == right.string {
							frame.stack.Push(&True{})
						} else {
							frame.stack.Push(&False{})
						}
					} else if first == OP_ISNT {
						if left.string != right.string {
							frame.stack.Push(&True{})
						} else {
							frame.stack.Push(&False{})
						}
					} else {
						panic(fmt.Sprintf("Unimplemented operator: %d", int(first)))
					}
				}
			} else {
				panic(fmt.Sprintf("Cannot compare %T and %T", rightx, leftx))
			}

		case POP_JUMP_IF_FALSE:
			o := frame.stack.Pop()

			if _, ok := o.(*False); ok {
				pc = int(first) + int(second) * 256
			}

		case BINARY_SUBTRACT: // TODO(flowlo): Implement this for floats
			if right, ok := frame.stack.Pop().(*Int); ok {
				if left, ok := frame.stack.Pop().(*Int); ok {
					frame.stack.Push(&Int{left.int32 - right.int32})
				}
			}

		case NOP:
		case ROT_TWO:
			a := frame.stack.Pop()
			b := frame.stack.Pop()
			frame.stack.Push(a)
			frame.stack.Push(b)
		case ROT_THREE:
			a := frame.stack.Pop()
			b := frame.stack.Pop()
			c := frame.stack.Pop()
			frame.stack.Push(b)
			frame.stack.Push(a)
			frame.stack.Push(c)
		//case DUP_TOP:
		//case DUP_TOP_TWO:
		case UNARY_POSITIVE: // TODO(flowlo): Implement for floats
			o := frame.stack.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = +a.int32
			}
		case UNARY_NEGATIVE: // TODO(flowlo): Implement for floats
			o := frame.stack.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = -a.int32
			}
		case UNARY_NOT: // TODO(flowlo): Implement for floats
			o := frame.stack.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = a.int32 // TODO
			}
		case UNARY_INVERT: // TODO(flowlo): Implement for floats
			o := frame.stack.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = ^a.int32
			}
		case IMPORT_NAME:
			name := (*code.Names)[first]
			fromlist := frame.stack.Pop()
			level := frame.stack.Pop()
			fmt.Printf("import %s with %s and %s\n", name, fromlist, level)

			if s, ok := name.(*String); ok {
				if s.string == "sys" {
					frame.stack.Push(NewSys())
				}
			}
		case LOAD_ATTR:
			name := (*code.Names)[first]
			o := frame.stack.Pop()

			ao, ok := o.(*Code)
			if !ok {
				panic(fmt.Sprintf("TypeError: %v of type %T has no attributes!", o, o))
			}
			a, err := ao.GetAttribute(name, nil)
			if err != nil {
				panic(err.Error())
			}
			frame.stack.Push(a)
			//frame.stack.Push(&String{"NOT IMPLEMENTED"})
		case JUMP_ABSOLUTE:
			pc = int(first)
		case SETUP_LOOP:
			block := &Block{pc, pc + int(first)}
			frame.blocks = append(frame.blocks, *block)
		case POP_BLOCK:
			frame.blocks = frame.blocks[:len(frame.blocks)-1]
		default:
			fmt.Sprintf("\x1b[31;1mSkipped\x1b[0m unknown opcode: %d\n", op)
			panic(fmt.Sprintf("Unknown opcode: %d", op))
		}
	}
	return &Null{}
}

func (frame *Frame) Execute(module Module) {
	frame.stack = new(Stack)
	frame.Run(*module.Code)
}
