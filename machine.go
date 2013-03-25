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
	log.Printf("%s (%d) %d", code.Name.string, len(code.Instructions), code.Instructions)

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
		case LOAD_FAST:
			machine.stack.Push((*code.Varnames)[first])
		case MAKE_FUNCTION:
			if fqn, ok := machine.stack.Pop().(*String) ; ok {
				if code, ok := machine.stack.Pop().(*Code) ; ok {
					result := &Function{fqn, code}
					
/*					if first > 0 { // TODO fix this. horribly ugly!
						args := make(Tuple, int(first))
			
						for j := int(first) - 1; j > -1; j-- {
							args[j] = machine.stack.Pop()
						}
						
						result.Code.Varnames = &args
					}
*/		
					machine.stack.Push(result)
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
			
			args := make(Tuple, int(first))
			
			for j := int(first) - 1; j > -1; j-- {
				args[j] = machine.stack.Pop()
			}
			
			if function, ok := machine.stack.Pop().(*Function) ; ok {
				function.Code.Varnames = &args
				machine.Run(*function.Code, 0)
			}
			
		case BINARY_MULTIPLY: // TODO implement this for floats
			if a, ok := machine.stack.Pop().(*Int) ; ok {
				if b, ok := machine.stack.Pop().(*Int) ; ok {
					machine.stack.Push(&Int{a.int32 * b.int32})
				}
			}
		case BINARY_ADD: // TODO implement this for floats
			if a, ok := machine.stack.Pop().(*Int) ; ok {
				if b, ok := machine.stack.Pop().(*Int) ; ok {
					machine.stack.Push(&Int{a.int32 + b.int32})
				}
			}
		}
		if op != CALL_FUNCTION {
			log.Printf("%-15s (%-3d) %-4d %-4d : %s", mnemonic[op], op, first, second, machine.stack)
		}
	}
}

func (machine *Machine) Execute(module Module) {
	machine.stack = new(Stack)
	machine.Run(*module.Code, 0)
}
