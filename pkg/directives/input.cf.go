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
				annotations.EvalStage: "parse",
				annotations.Type:      "directive",
			},
			ID: "github.com/conflowio/conflow/pkg/directives.Input",
		},
		ParameterNames: map[string]string{"schema": "type"},
		Properties: map[string]schema.Schema{
			"id": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
			"required": &schema.Boolean{},
			"schema": &schema.Reference{
				Ref: "github.com/conflowio/conflow/pkg/schema.Schema",
			},
		},
		Required: []string{"type"},
	})
}

// InputInterpreter is the Conflow interpreter for the Input block
type InputInterpreter struct {
}

func (i InputInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/pkg/directives.Input")
	return s
}

// Create creates a new Input block
func (i InputInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &Input{}
	b.id = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i InputInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i InputInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Input
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i InputInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*Input).id
	case "required":
		return b.(*Input).required
	default:
		panic(fmt.Errorf("unexpected parameter %q in Input", name))
	}
}

func (i InputInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Input)
	switch name {
	case "required":
		b.required = schema.Value[bool](value)
	}
	return nil
}

func (i InputInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	b := block.(*Input)
	switch name {
	case "type":
		b.schema = value.(schema.Schema)
	}
	return nil
}
