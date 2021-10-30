// Code generated by Conflow. DO NOT EDIT.

package math

import (
	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
	"github.com/conflowio/parsley/parsley"
)

// CeilInterpreter is the conflow interpreter for the Ceil function
type CeilInterpreter struct {
	s schema.Schema
}

func (i CeilInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Function{
			Metadata: schema.Metadata{
				Description: "It returns the least integer value greater than or equal to x.",
			},
			Parameters: schema.Parameters{
				schema.NamedSchema{
					Name: "number",
					Schema: &schema.Untyped{
						Types: []string{"integer", "number"},
					},
				},
			},
			Result: &schema.Integer{},
		}
	}
	return i.s
}

// Eval returns with the result of the function
func (i CeilInterpreter) Eval(ctx interface{}, node conflow.FunctionNode) (interface{}, parsley.Error) {
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

	return Ceil(val0), nil
}
