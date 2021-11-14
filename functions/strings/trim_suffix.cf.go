// Code generated by Conflow. DO NOT EDIT.

package strings

import (
	"github.com/conflowio/conflow/conflow/schema"
)

// TrimSuffixInterpreter is the conflow interpreter for the TrimSuffix function
type TrimSuffixInterpreter struct {
	s schema.Schema
}

func (i TrimSuffixInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Function{
			Metadata: schema.Metadata{
				Description: "It returns s without the provided trailing suffix string.\nIf s doesn't end with suffix, s is returned unchanged.",
			},
			Parameters: schema.Parameters{
				schema.NamedSchema{
					Name:   "s",
					Schema: &schema.String{},
				},
				schema.NamedSchema{
					Name:   "suffix",
					Schema: &schema.String{},
				},
			},
			Result: &schema.String{},
		}
	}
	return i.s
}

// Eval returns with the result of the function
func (i TrimSuffixInterpreter) Eval(ctx interface{}, args []interface{}) (interface{}, error) {
	var val0 = args[0].(string)
	var val1 = args[1].(string)
	return TrimSuffix(val0, val1), nil
}
