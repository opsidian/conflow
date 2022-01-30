// Code generated by Conflow. DO NOT EDIT.

package strings

import (
	"github.com/conflowio/conflow/src/schema"
)

func init() {
	schema.Register(&schema.Function{
		Metadata: schema.Metadata{
			Description: "It returns a copy of the string s with all Unicode letters mapped to their lower case.",
			ID:          "github.com/conflowio/conflow/src/functions/strings.Lower",
		},
		Parameters: schema.Parameters{
			schema.NamedSchema{
				Name:   "s",
				Schema: &schema.String{},
			},
		},
		Result: &schema.String{},
	})
}

// LowerInterpreter is the Conflow interpreter for the Lower function
type LowerInterpreter struct {
}

func (i LowerInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/functions/strings.Lower")
	return s
}

// Eval returns with the result of the function
func (i LowerInterpreter) Eval(ctx interface{}, args []interface{}) (interface{}, error) {
	var val0 = args[0].(string)
	return Lower(val0), nil
}
