package gothon

import (
	"log"
)

type Machine struct {
	stack *Stack
	pc int
}

func HasArg(opcode byte) bool {
	return opcode > 89
}

func (machine *Machine) Execute(module Module) {
	log.Printf("(%d) %d", len(module.Code.Instructions), module.Code.Instructions)

	machine.pc, machine.stack = 0, new(Stack)
	var op, first, second byte
	
	for machine.pc < len(module.Code.Instructions) {
		op = module.Code.Instructions[machine.pc]
		machine.pc++
		
		if HasArg(op) {
			first, second = module.Code.Instructions[machine.pc], module.Code.Instructions[machine.pc + 1]
			machine.pc += 2
		} else {
			first, second = 0, 0
		}
		
		switch op {
		case LOAD_CONST:
			machine.stack.Push((*module.Code.Consts)[first])
		case STORE_NAME:
			(*module.Code.Consts)[first] = machine.stack.Pop()
//		case MAKE_FUNCTION:
//			machine.stack.Push(&Function{machine.stack.Pop(), machine.stack.Pop()})
		default:
			log.Printf("%-5d %-4d %-4d %-15s", op, first, second, mnemonic[op])
		}
	}
}
