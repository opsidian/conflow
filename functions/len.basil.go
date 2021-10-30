// Code generated by Basil. DO NOT EDIT.

package functions

import (
	"github.com/opsidian/conflow/basil"
	"github.com/opsidian/conflow/basil/schema"
	"github.com/opsidian/parsley/parsley"
)

// LenInterpreter is the basil interpreter for the Len function
type LenInterpreter struct {
	s schema.Schema
}

func (i LenInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Function{
			Metadata: schema.Metadata{
				Description: "It returns with the length of the variable\nFor strings it means the count of UTF-8 characters\nFor arrays and maps it means the number of items/entries",
			},
			Parameters: schema.Parameters{
				schema.NamedSchema{
					Name: "value",
					Schema: &schema.Untyped{
						Types: []string{"string", "array", "map"},
					},
				},
			},
			Result: &schema.Integer{},
		}
	}
	return i.s
}

// Eval returns with the result of the function
func (i LenInterpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
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

	return Len(val0), nil
}
