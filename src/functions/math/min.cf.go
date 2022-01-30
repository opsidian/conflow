// Code generated by Conflow. DO NOT EDIT.

package math

import (
	"github.com/conflowio/conflow/src/schema"
)

func init() {
	schema.Register(&schema.Function{
		Metadata: schema.Metadata{
			Description: "It returns with the lowest value",
			ID:          "github.com/conflowio/conflow/src/functions/math.Min",
		},
		AdditionalParameters: &schema.NamedSchema{
			Name: "rest",
			Schema: &schema.Untyped{
				Types: []string{"integer", "number"},
			},
		},
		Parameters: schema.Parameters{
			schema.NamedSchema{
				Name: "min",
				Schema: &schema.Untyped{
					Types: []string{"integer", "number"},
				},
			},
		},
		Result: &schema.Number{},
	})
}

// MinInterpreter is the Conflow interpreter for the Min function
type MinInterpreter struct {
}

func (i MinInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/functions/math.Min")
	return s
}

// Eval returns with the result of the function
func (i MinInterpreter) Eval(ctx interface{}, args []interface{}) (interface{}, error) {
	var val0 = args[0]
	var variadicArgs []interface{}
	for p := 1; p < len(args); p++ {
		var val = args[p]
		variadicArgs = append(variadicArgs, val)
	}
	return Min(val0, variadicArgs...), nil
}
