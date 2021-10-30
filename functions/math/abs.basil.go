// Code generated by Basil. DO NOT EDIT.

package math

import (
	"github.com/opsidian/conflow/basil"
	"github.com/opsidian/conflow/basil/function"
	"github.com/opsidian/conflow/basil/schema"
	"github.com/opsidian/parsley/parsley"
)

// AbsInterpreter is the basil interpreter for the Abs function
type AbsInterpreter struct {
	s schema.Schema
}

func (i AbsInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Function{
			Metadata: schema.Metadata{
				Description: "It returns the absolute value of the given number",
			},
			Parameters: schema.Parameters{
				schema.NamedSchema{
					Name: "value",
					Schema: &schema.Untyped{
						Types: []string{"integer", "number"},
					},
				},
			},
			Result:         &schema.Untyped{},
			ResultTypeFrom: "value",
		}
	}
	return i.s
}

// Eval returns with the result of the function
func (i AbsInterpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	parameters := i.Schema().(*schema.Function).GetParameters()
	arguments := node.ArgumentNodes()

	arg0, evalErr := parsley.EvaluateNode(ctx, arguments[0])
	if evalErr != nil {
		return nil, evalErr
	}
	if err := parameters[0].Schema.ValidateValue(arg0); err != nil {
		return nil, parsley.NewError(arguments[0].Pos(), err)
	}
	var val0 = arg0

	res, resErr := Abs(val0)
	if resErr != nil {
		if funcErr, ok := resErr.(*function.Error); ok {
			return nil, parsley.NewError(arguments[funcErr.ArgIndex].Pos(), funcErr.Err)
		}
		return nil, parsley.NewError(node.Pos(), resErr)
	}

	return res, nil
}
