// Code generated by Conflow. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/conflowio/conflow/pkg/conflow"
	"github.com/conflowio/conflow/pkg/conflow/annotations"
	"github.com/conflowio/conflow/pkg/schema"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{
				annotations.Type: "directive",
			},
			Description: "It is the directive for marking functions as conflow functions",
			ID:          "github.com/conflowio/conflow/pkg/schema/directives.Function",
		},
		ParameterNames: map[string]string{"Path": "path"},
		Properties: map[string]schema.Schema{
			"Path": &schema.String{},
			"id": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
		},
	})
}

// FunctionInterpreter is the Conflow interpreter for the Function block
type FunctionInterpreter struct {
}

func (i FunctionInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/pkg/schema/directives.Function")
	return s
}

// Create creates a new Function block
func (i FunctionInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &Function{}
	b.id = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i FunctionInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i FunctionInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Function
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i FunctionInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "path":
		return b.(*Function).Path
	case "id":
		return b.(*Function).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in Function", name))
	}
}

func (i FunctionInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Function)
	switch name {
	case "path":
		b.Path = schema.Value[string](value)
	}
	return nil
}

func (i FunctionInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
