package gothon

const (
	POP_TOP                 = 1
	ROT_TWO                 = 2
	ROT_THREE               = 3
	DUP_TOP                 = 4
	DUP_TOP_TWO             = 5
	NOP                     = 9
	UNARY_POSITIVE          = 10
	UNARY_NEGATIVE          = 11
	UNARY_NOT               = 12
	UNARY_INVERT            = 15
	BINARY_MATRIX_MULTIPLY  = 16
	INPLACE_MATRIX_MULTIPLY = 17
	BINARY_POWER            = 19
	BINARY_MULTIPLY         = 20
	BINARY_MODULO           = 22
	BINARY_ADD              = 23
	BINARY_SUBTRACT         = 24
	BINARY_SUBSCR           = 25
	BINARY_FLOOR_DIVIDE     = 26
	BINARY_TRUE_DIVIDE      = 27
	INPLACE_FLOOR_DIVIDE    = 28
	INPLACE_TRUE_DIVIDE     = 29
	STORE_MAP               = 54
	INPLACE_ADD             = 55
	INPLACE_SUBTRACT        = 56
	INPLACE_MULTIPLY        = 57
	INPLACE_MODULO          = 59
	STORE_SUBSCR            = 60
	DELETE_SUBSCR           = 61
	BINARY_LSHIFT           = 62
	BINARY_RSHIFT           = 63
	BINARY_AND              = 64
	BINARY_XOR              = 65
	BINARY_OR               = 66
	INPLACE_POWER           = 67
	GET_ITER                = 68
	PRINT_EXPR              = 70
	LOAD_BUILD_CLASS        = 71
	YIELD_FROM              = 72
	INPLACE_LSHIFT          = 75
	INPLACE_RSHIFT          = 76
	INPLACE_AND             = 77
	INPLACE_XOR             = 78
	INPLACE_OR              = 79
	BREAK_LOOP              = 80
	WITH_CLEANUP            = 81
	RETURN_VALUE            = 83
	IMPORT_STAR             = 84
	YIELD_VALUE             = 86
	POP_BLOCK               = 87
	END_FINALLY             = 88
	POP_EXCEPT              = 89
	HAVE_ARGUMENT           = 90
	STORE_NAME              = 90
	DELETE_NAME             = 91
	UNPACK_SEQUENCE         = 92
	FOR_ITER                = 93
	UNPACK_EX               = 94
	STORE_ATTR              = 95
	DELETE_ATTR             = 96
	STORE_GLOBAL            = 97
	DELETE_GLOBAL           = 98
	LOAD_CONST              = 100
	LOAD_NAME               = 101
	BUILD_TUPLE             = 102
	BUILD_LIST              = 103
	BUILD_SET               = 104
	BUILD_MAP               = 105
	LOAD_ATTR               = 106
	COMPARE_OP              = 107
	IMPORT_NAME             = 108
	IMPORT_FROM             = 109
	JUMP_FORWARD            = 110
	JUMP_IF_FALSE_OR_POP    = 111
	JUMP_IF_TRUE_OR_POP     = 112
	JUMP_ABSOLUTE           = 113
	POP_JUMP_IF_FALSE       = 114
	POP_JUMP_IF_TRUE        = 115
	LOAD_GLOBAL             = 116
	CONTINUE_LOOP           = 119
	SETUP_LOOP              = 120
	SETUP_EXCEPT            = 121
	SETUP_FINALLY           = 122
	LOAD_FAST               = 124
	STORE_FAST              = 125
	DELETE_FAST             = 126
	RAISE_VARARGS           = 130
	CALL_FUNCTION           = 131
	MAKE_FUNCTION           = 132
	BUILD_SLICE             = 133
	MAKE_CLOSURE            = 134
	LOAD_CLOSURE            = 135
	LOAD_DEREF              = 136
	STORE_DEREF             = 137
	DELETE_DEREF            = 138
	CALL_FUNCTION_VAR       = 140
	CALL_FUNCTION_KW        = 141
	CALL_FUNCTION_VAR_KW    = 142
	SETUP_WITH              = 143
	EXTENDED_ARG            = 144
	LIST_APPEND             = 145
	SET_ADD                 = 146
	MAP_ADD                 = 147
	LOAD_CLASSDEREF         = 148
)

var opcode = map[byte]string{
	POP_TOP:                 "POP_TOP",
	ROT_TWO:                 "ROT_TWO",
	ROT_THREE:               "ROT_THREE",
	DUP_TOP:                 "DUP_TOP",
	DUP_TOP_TWO:             "DUP_TOP_TWO",
	NOP:                     "NOP",
	UNARY_POSITIVE:          "UNARY_POSITIVE",
	UNARY_NEGATIVE:          "UNARY_NEGATIVE",
	UNARY_NOT:               "UNARY_NOT",
	UNARY_INVERT:            "UNARY_INVERT",
	BINARY_MATRIX_MULTIPLY:  "BINARY_MATRIX_MULTIPLY",
	INPLACE_MATRIX_MULTIPLY: "INPLACE_MATRIX_MULTIPLY",
	BINARY_POWER:            "BINARY_POWER",
	BINARY_MULTIPLY:         "BINARY_MULTIPLY",
	BINARY_MODULO:           "BINARY_MODULO",
	BINARY_ADD:              "BINARY_ADD",
	BINARY_SUBTRACT:         "BINARY_SUBTRACT",
	BINARY_SUBSCR:           "BINARY_SUBSCR",
	BINARY_FLOOR_DIVIDE:     "BINARY_FLOOR_DIVIDE",
	BINARY_TRUE_DIVIDE:      "BINARY_TRUE_DIVIDE",
	INPLACE_FLOOR_DIVIDE:    "INPLACE_FLOOR_DIVIDE",
	INPLACE_TRUE_DIVIDE:     "INPLACE_TRUE_DIVIDE",
	STORE_MAP:               "STORE_MAP",
	INPLACE_ADD:             "INPLACE_ADD",
	INPLACE_SUBTRACT:        "INPLACE_SUBTRACT",
	INPLACE_MULTIPLY:        "INPLACE_MULTIPLY",
	INPLACE_MODULO:          "INPLACE_MODULO",
	STORE_SUBSCR:            "STORE_SUBSCR",
	DELETE_SUBSCR:           "DELETE_SUBSCR",
	BINARY_LSHIFT:           "BINARY_LSHIFT",
	BINARY_RSHIFT:           "BINARY_RSHIFT",
	BINARY_AND:              "BINARY_AND",
	BINARY_XOR:              "BINARY_XOR",
	BINARY_OR:               "BINARY_OR",
	INPLACE_POWER:           "INPLACE_POWER",
	GET_ITER:                "GET_ITER",
	PRINT_EXPR:              "PRINT_EXPR",
	LOAD_BUILD_CLASS:        "LOAD_BUILD_CLASS",
	YIELD_FROM:              "YIELD_FROM",
	INPLACE_LSHIFT:          "INPLACE_LSHIFT",
	INPLACE_RSHIFT:          "INPLACE_RSHIFT",
	INPLACE_AND:             "INPLACE_AND",
	INPLACE_XOR:             "INPLACE_XOR",
	INPLACE_OR:              "INPLACE_OR",
	BREAK_LOOP:              "BREAK_LOOP",
	WITH_CLEANUP:            "WITH_CLEANUP",
	RETURN_VALUE:            "RETURN_VALUE",
	IMPORT_STAR:             "IMPORT_STAR",
	YIELD_VALUE:             "YIELD_VALUE",
	POP_BLOCK:               "POP_BLOCK",
	END_FINALLY:             "END_FINALLY",
	POP_EXCEPT:              "POP_EXCEPT",
	STORE_NAME:              "STORE_NAME",
	DELETE_NAME:             "DELETE_NAME",
	UNPACK_SEQUENCE:         "UNPACK_SEQUENCE",
	FOR_ITER:                "FOR_ITER",
	UNPACK_EX:               "UNPACK_EX",
	STORE_ATTR:              "STORE_ATTR",
	DELETE_ATTR:             "DELETE_ATTR",
	STORE_GLOBAL:            "STORE_GLOBAL",
	DELETE_GLOBAL:           "DELETE_GLOBAL",
	LOAD_CONST:              "LOAD_CONST",
	LOAD_NAME:               "LOAD_NAME",
	BUILD_TUPLE:             "BUILD_TUPLE",
	BUILD_LIST:              "BUILD_LIST",
	BUILD_SET:               "BUILD_SET",
	BUILD_MAP:               "BUILD_MAP",
	LOAD_ATTR:               "LOAD_ATTR",
	COMPARE_OP:              "COMPARE_OP",
	IMPORT_NAME:             "IMPORT_NAME",
	IMPORT_FROM:             "IMPORT_FROM",
	JUMP_FORWARD:            "JUMP_FORWARD",
	JUMP_IF_FALSE_OR_POP:    "JUMP_IF_FALSE_OR_POP",
	JUMP_IF_TRUE_OR_POP:     "JUMP_IF_TRUE_OR_POP",
	JUMP_ABSOLUTE:           "JUMP_ABSOLUTE",
	POP_JUMP_IF_FALSE:       "POP_JUMP_IF_FALSE",
	POP_JUMP_IF_TRUE:        "POP_JUMP_IF_TRUE",
	LOAD_GLOBAL:             "LOAD_GLOBAL",
	CONTINUE_LOOP:           "CONTINUE_LOOP",
	SETUP_LOOP:              "SETUP_LOOP",
	SETUP_EXCEPT:            "SETUP_EXCEPT",
	SETUP_FINALLY:           "SETUP_FINALLY",
	LOAD_FAST:               "LOAD_FAST",
	STORE_FAST:              "STORE_FAST",
	DELETE_FAST:             "DELETE_FAST",
	RAISE_VARARGS:           "RAISE_VARARGS",
	CALL_FUNCTION:           "CALL_FUNCTION",
	MAKE_FUNCTION:           "MAKE_FUNCTION",
	BUILD_SLICE:             "BUILD_SLICE",
	MAKE_CLOSURE:            "MAKE_CLOSURE",
	LOAD_CLOSURE:            "LOAD_CLOSURE",
	LOAD_DEREF:              "LOAD_DEREF",
	STORE_DEREF:             "STORE_DEREF",
	DELETE_DEREF:            "DELETE_DEREF",
	CALL_FUNCTION_VAR:       "CALL_FUNCTION_VAR",
	CALL_FUNCTION_KW:        "CALL_FUNCTION_KW",
	CALL_FUNCTION_VAR_KW:    "CALL_FUNCTION_VAR_KW",
	SETUP_WITH:              "SETUP_WITH",
	EXTENDED_ARG:            "EXTENDED_ARG",
	LIST_APPEND:             "LIST_APPEND",
	SET_ADD:                 "SET_ADD",
	MAP_ADD:                 "MAP_ADD",
	LOAD_CLASSDEREF:         "LOAD_CLASSDEREF",
}
