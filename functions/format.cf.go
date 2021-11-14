// Code generated by Conflow. DO NOT EDIT.

package functions

import (
	"github.com/conflowio/conflow/conflow/schema"
)

// FormatInterpreter is the conflow interpreter for the Format function
type FormatInterpreter struct {
	s schema.Schema
}

func (i FormatInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Function{
			Metadata: schema.Metadata{
				Description: "It formats according to a format specifier and returns the resulting string.",
			},
			AdditionalParameters: &schema.NamedSchema{
				Name:   "values",
				Schema: &schema.Untyped{},
			},
			Parameters: schema.Parameters{
				schema.NamedSchema{
					Name:   "format",
					Schema: &schema.String{},
				},
			},
			Result: &schema.String{},
		}
	}
	return i.s
}

// Eval returns with the result of the function
func (i FormatInterpreter) Eval(ctx interface{}, args []interface{}) (interface{}, error) {
	var val0 = args[0].(string)
	var variadicArgs []interface{}
	for p := 1; p < len(args); p++ {
		var val = args[p]
		variadicArgs = append(variadicArgs, val)
	}
	return Format(val0, variadicArgs...), nil
}
