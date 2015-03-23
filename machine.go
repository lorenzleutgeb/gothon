package gothon

import (
	"fmt"
)

type Machine struct {
	stack *Stack
}

func HasArg(opcode byte) bool {
	return opcode > 89
}

func (machine *Machine) Run(code Code) Object {
	fmt.Printf("\x1b[32;1m%s\x1b[0m\n", code.Name.string)

	pc := 0
	var op, first, second byte

	for pc < len(code.Instructions) {
		op = code.Instructions[pc]
		if HasArg(op) {
			first, second = code.Instructions[pc+1], code.Instructions[pc+2]
			pc += 2
		} else {
			first, second = 0, 0
		}

		fmt.Printf("%-4d %-15s (%-3d) %-4d %-4d : %s\n", pc, opcode[op], op, first, second, machine.stack)

		pc += 1

		switch op {
		case POP_TOP:
			machine.stack.Pop()
		case LOAD_CONST:
			machine.stack.Push((*code.Consts)[first])
		case STORE_NAME:
			(*code.Names)[first] = machine.stack.Pop()
		case LOAD_NAME:
			machine.stack.Push((*code.Names)[first])
		case LOAD_FAST:
			machine.stack.Push((*code.Varnames)[first])
		case MAKE_FUNCTION:
			if fqn, ok := machine.stack.Pop().(*String); ok {
				if code, ok := machine.stack.Pop().(*Code); ok {
					result := &Function{
						Name: fqn,
						Code: code,
					}
					machine.stack.Push(result)
				}
			}
		case CALL_FUNCTION:
			if second > 0 {
				panic("Keyword parameters are not implemented.")
			}

			args := make(Tuple, int(first))

			for j := int(first) - 1; j > -1; j-- {
				args[j] = machine.stack.Pop()
			}

			o := machine.stack.Pop()

			if function, ok := o.(*Function); ok {
				function.Code.Varnames = &args
				machine.stack.Push(machine.Run(*function.Code))
			} else if s, ok := o.(*String); ok {
				if s.string == "print" { // print is built-in
					for _, arg := range args {
						fmt.Printf("\x1b[31;1m%+v\x1b[0m", arg)
					}
					fmt.Println()
				} else if s.string == code.Name.string {
					code.Varnames = &args
					machine.stack.Push(machine.Run(code))
				} else {
					panic(fmt.Sprintf("Unknown function: %s", s))
				}
			}

		case BINARY_MULTIPLY: // TODO(flowlo): implement this for floats
			if a, ok := machine.stack.Pop().(*Int); ok {
				if b, ok := machine.stack.Pop().(*Int); ok {
					machine.stack.Push(&Int{a.int32 * b.int32})
				}
			}
		case BINARY_ADD: // TODO(flowlo): implement this for floats
			if a, ok := machine.stack.Pop().(*Int); ok {
				if b, ok := machine.stack.Pop().(*Int); ok {
					machine.stack.Push(&Int{a.int32 + b.int32})
				}
			}
		case RETURN_VALUE:
			return machine.stack.Pop()

		case LOAD_GLOBAL:
			machine.stack.Push((*code.Names)[first])

		case COMPARE_OP:
			if right, ok := machine.stack.Pop().(*Int); ok {
				if left, ok := machine.stack.Pop().(*Int); ok {
					switch first {
					case OP_LT:
						if left.int32 < right.int32 {
							machine.stack.Push(&True{})
						} else {
							machine.stack.Push(&False{})
						}
					case OP_LEQ:
						if left.int32 <= right.int32 {
							machine.stack.Push(&True{})
						} else {
							machine.stack.Push(&False{})
						}

					case OP_EQ:
						if left.int32 == right.int32 {
							machine.stack.Push(&True{})
						} else {
							machine.stack.Push(&False{})
						}
					case OP_GT:
						if left.int32 > right.int32 {
							machine.stack.Push(&True{})
						} else {
							machine.stack.Push(&False{})
						}
					case OP_GE:
						if left.int32 >= right.int32 {
							machine.stack.Push(&True{})
						} else {
							machine.stack.Push(&False{})
						}
					default:
						panic("Comparison operator not implemented.")
					}
				}
			}

		case POP_JUMP_IF_FALSE:
			o := machine.stack.Pop()

			if _, ok := o.(*False); ok {
				pc = int(first)
			}

		case BINARY_SUBTRACT: // TODO(flowlo): Implement this for floats
			if right, ok := machine.stack.Pop().(*Int); ok {
				if left, ok := machine.stack.Pop().(*Int); ok {
					machine.stack.Push(&Int{left.int32 - right.int32})
				}
			}

		case NOP:
		case ROT_TWO:
			a := machine.stack.Pop()
			b := machine.stack.Pop()
			machine.stack.Push(a)
			machine.stack.Push(b)
		case ROT_THREE:
			a := machine.stack.Pop()
			b := machine.stack.Pop()
			c := machine.stack.Pop()
			machine.stack.Push(b)
			machine.stack.Push(a)
			machine.stack.Push(c)
		//case DUP_TOP:
		//case DUP_TOP_TWO:
		case UNARY_POSITIVE: // TODO(flowlo): Implement for floats
			o := machine.stack.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = +a.int32
			}
		case UNARY_NEGATIVE: // TODO(flowlo): Implement for floats
			o := machine.stack.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = -a.int32
			}
		case UNARY_NOT: // TODO(flowlo): Implement for floats
			o := machine.stack.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = a.int32 // TODO
			}
		case UNARY_INVERT: // TODO(flowlo): Implement for floats
			o := machine.stack.Pop()
			if a, ok := o.(*Int); ok {
				a.int32 = ^a.int32
			}
		case IMPORT_NAME:
			name := (*code.Names)[first]
			fromlist := machine.stack.Pop()
			level := machine.stack.Pop()
			fmt.Printf("import %s with %s and %s\n", name, fromlist, level)

			if s, ok := name.(*String); ok {
				if s.string == "sys" {
					machine.stack.Push(NewSys())
				}
			}
		case LOAD_ATTR:
			name := (*code.Names)[first]
			o := machine.stack.Pop()

			ao, ok := o.(*Code)
			if !ok {
				panic(fmt.Sprintf("TypeError: %v of type %T has no attributes!", o, o))
			}
			a, err := ao.GetAttribute(name, nil)
			if err != nil {
				panic(err.Error())
			}
			machine.stack.Push(a)
			//machine.stack.Push(&String{"NOT IMPLEMENTED"})
		default:
			fmt.Println("\x1b[31;1mSKIPPED\x1b[0m")
		}
	}
	return &Null{}
}

func (machine *Machine) Execute(module Module) {
	machine.stack = new(Stack)
	machine.Run(*module.Code)
}
