// Code generated by Conflow. DO NOT EDIT.

package math

import (
	"github.com/conflowio/conflow/src/schema"
)

func init() {
	schema.Register(&schema.Function{
		Metadata: schema.Metadata{
			Description: "It returns the least integer value greater than or equal to x.",
			ID:          "github.com/conflowio/conflow/src/functions/math.Ceil",
		},
		Parameters: schema.Parameters{
			schema.NamedSchema{
				Name: "number",
				Schema: &schema.Any{
					Types: []string{"integer", "number"},
				},
			},
		},
		Result: &schema.Integer{},
	})
}

// CeilInterpreter is the Conflow interpreter for the Ceil function
type CeilInterpreter struct {
}

func (i CeilInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/functions/math.Ceil")
	return s
}

// Eval returns with the result of the function
func (i CeilInterpreter) Eval(ctx interface{}, args []interface{}) (interface{}, error) {
	var val0 = args[0]
	return Ceil(val0), nil
}
