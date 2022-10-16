// Code generated by Conflow. DO NOT EDIT.

package math

import (
	"github.com/conflowio/conflow/src/schema"
)

func init() {
	schema.Register(&schema.Function{
		Metadata: schema.Metadata{
			Description: "It returns the absolute value of the given number",
			ID:          "github.com/conflowio/conflow/src/functions/math.Abs",
		},
		Parameters: schema.Parameters{
			schema.NamedSchema{
				Name: "value",
				Schema: &schema.Any{
					Types: []string{"integer", "number"},
				},
			},
		},
		Result:         &schema.Any{},
		ResultTypeFrom: "value",
	})
}

// AbsInterpreter is the Conflow interpreter for the Abs function
type AbsInterpreter struct {
}

func (i AbsInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/functions/math.Abs")
	return s
}

// Eval returns with the result of the function
func (i AbsInterpreter) Eval(ctx interface{}, args []interface{}) (interface{}, error) {
	var val0 = args[0]
	return Abs(val0)
}
