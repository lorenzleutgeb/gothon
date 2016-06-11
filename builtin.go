package main

import (
	"fmt"
	"strings"
)

var builtin map[string]Object

func init() {
	builtin = map[string]Object{
		"print": NewInternalFunction("print", func(args *args) Object {
			if len(args.Keyword) > 0 {
				panic("print() is not fancy")
			}

			s := make([]string, len(args.Positional))

			for i, v := range args.Positional {
				s[i] = v.String()
			}

			fmt.Println(strings.Join(s, " "))
			return None{}
		}),
		"set": NewInternalFunction("set", func(args *args) Object {
			if !args.IsEmpty() {
				panic("set() with iterable not implemented")
			}
			return Set{}
		}),
		"abs": NewInternalFunction("abs", func(args *args) Object {
			if len(args.Keyword) > 0 || len(args.Positional) != 1 {
				return None{}
			}
			num := args.Positional[0]
			switch t := num.(type) {
			case Int:
				if t.int32 < 0 {
					return Int{-t.int32}
				}
			default:
				panic("cannot take abs() of unknown type")
			}
			return num
		}),
		"__import__": NewInternalFunction("__import__", func(args *args) Object {
			panic("__import__() is not yet implemented")
		}),
		"len": NewInternalFunction("len", func(args *args) Object {
			res := len(args.Keyword) + len(args.Positional)
			result := Int{int32(res)}
			return result
		}),
		"sum": NewInternalFunction("sum", func(args *args) Object {
			if len(args.Keyword) > 0 || len(args.Positional) < 1 {
				panic("all need to be numbers to use sum()")
			}
			sum := Int{0}
			for _,v := range args.Positional {
				sum.int32 += int32(v)
			}

			return sum
		}),
	}
}
