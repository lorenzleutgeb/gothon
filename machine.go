package gothon

import (
	"log"
)

type Machine struct {
	stack *Stack
}

func HasArg(opcode byte) bool {
	return opcode > 89
}

func (machine *Machine) Run(code Code, pc int) {
	log.Printf("(%d) %d", len(code.Instructions), code.Instructions)

	pc = 0
	var op, first, second byte
	
	for pc < len(code.Instructions) {
		op = code.Instructions[pc]
		pc++
		
		if HasArg(op) {
			first, second = code.Instructions[pc], code.Instructions[pc + 1]
			pc += 2
		} else {
			first, second = 0, 0
		}
		
		switch op {
		case POP_TOP:
			machine.stack.Pop()
		case LOAD_CONST:
			machine.stack.Push((*code.Consts)[first])
		case STORE_NAME:
			(*code.Consts)[first] = machine.stack.Pop()
		case LOAD_NAME:
			machine.stack.Push((*code.Consts)[first])
		case MAKE_FUNCTION:
			if fqn, ok := machine.stack.Pop().(*String) ; ok {
				if code, ok := machine.stack.Pop().(*Code) ; ok {
					machine.stack.Push(&Function{fqn, code})
				}
			}
		case CALL_FUNCTION:
			// first: The low byte of argc indicates the number of positional parameters,
			// second: the high byte the number of keyword parameters.
			// On the stack, the opcode finds the keyword parameters first.
			// For each keyword argument, the value is on top of the key.
			// Below the keyword parameters, the positional parameters are on the stack, with the right-most parameter on top.
			// Below the parameters, the function object to call is on the stack. Pops all function arguments, and the function itself off the stack, and pushes the return value.
			
			// TODO Pop keyword parameters
			
			for j := 0; j < int(first); j++ {
				machine.stack.Pop()
			}
			
			if function, ok := machine.stack.Pop().(*Function) ; ok {
				machine.Run(*function.Code, 0)
			}
			
		case BINARY_MULTIPLY: // TODO implement this for floats
			if a, ok := machine.stack.Pop().(*Int) ; ok {
				if b, ok := machine.stack.Pop().(*Int) ; ok {
					machine.stack.Push(&Int{a.int32 * b.int32})
				}
			}
		default:
			log.Printf("%-5d %-4d %-4d %-15s", op, first, second, mnemonic[op])
		}
		log.Printf("Stack after executing %s: %s", mnemonic[op], machine.stack)
	}
}

func (machine *Machine) Execute(module Module) {
	machine.stack = new(Stack)
	machine.Run(*module.Code, 0)
}
