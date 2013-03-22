package gothon

import ("container/list")

type Machine struct {
   stack list.List
}

func HasArg(opcode byte) bool {
	return opcode > 89
}
