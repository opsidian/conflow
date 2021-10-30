// Code generated by Basil. DO NOT EDIT.

package test

import (
	"github.com/opsidian/conflow/basil"
	"github.com/opsidian/conflow/basil/schema"
	"github.com/opsidian/parsley/parsley"
)

// TestFunc0Interpreter is the basil interpreter for the testFunc0 function
type TestFunc0Interpreter struct {
	s schema.Schema
}

func (i TestFunc0Interpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Function{
			Result: &schema.String{},
		}
	}
	return i.s
}

// Eval returns with the result of the function
func (i TestFunc0Interpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	return testFunc0(), nil
}
