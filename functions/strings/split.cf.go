// Code generated by Conflow. DO NOT EDIT.

package strings

import (
	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
	"github.com/conflowio/parsley/parsley"
)

// SplitInterpreter is the conflow interpreter for the Split function
type SplitInterpreter struct {
	s schema.Schema
}

func (i SplitInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Function{
			Metadata: schema.Metadata{
				Description: "It slices s into all substrings separated by sep and returns a slice of\nthe substrings between those separators.\n\nIf s does not contain sep and sep is not empty, Split returns a\nslice of length 1 whose only element is s.\n\nIf sep is empty, Split splits after each UTF-8 sequence. If both s\nand sep are empty, Split returns an empty slice.",
			},
			Parameters: schema.Parameters{
				schema.NamedSchema{
					Name:   "s",
					Schema: &schema.String{},
				},
				schema.NamedSchema{
					Name:   "sep",
					Schema: &schema.String{},
				},
			},
			Result: &schema.Array{
				Items: &schema.Untyped{},
			},
		}
	}
	return i.s
}

// Eval returns with the result of the function
func (i SplitInterpreter) Eval(ctx interface{}, node conflow.FunctionNode) (interface{}, parsley.Error) {
	parameters := i.Schema().(*schema.Function).GetParameters()
	arguments := node.ArgumentNodes()

	arg0, evalErr := parsley.EvaluateNode(ctx, arguments[0])
	if evalErr != nil {
		return nil, evalErr
	}
	if err := parameters[0].Schema.ValidateValue(arg0); err != nil {
		return nil, parsley.NewError(arguments[0].Pos(), err)
	}
	var val0 = arg0.(string)

	arg1, evalErr := parsley.EvaluateNode(ctx, arguments[1])
	if evalErr != nil {
		return nil, evalErr
	}
	if err := parameters[1].Schema.ValidateValue(arg1); err != nil {
		return nil, parsley.NewError(arguments[1].Pos(), err)
	}
	var val1 = arg1.(string)

	return Split(val0, val1), nil
}
