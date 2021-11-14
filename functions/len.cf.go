// Code generated by Conflow. DO NOT EDIT.

package functions

import (
	"github.com/conflowio/conflow/conflow/schema"
)

// LenInterpreter is the conflow interpreter for the Len function
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
func (i LenInterpreter) Eval(ctx interface{}, args []interface{}) (interface{}, error) {
	var val0 = args[0]
	return Len(val0), nil
}
