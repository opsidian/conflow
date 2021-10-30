// Code generated by Conflow. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
)

// DefaultInterpreter is the conflow interpreter for the Default block
type DefaultInterpreter struct {
	s schema.Schema
}

func (i DefaultInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "Default",
			Properties: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "conflow.ID",
				},
				"value": &schema.Untyped{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"value": "true"},
					},
				},
			},
			Required: []string{"value"},
		}
	}
	return i.s
}

// Create creates a new Default block
func (i DefaultInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &Default{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i DefaultInterpreter) ValueParamName() conflow.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i DefaultInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Default
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i DefaultInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*Default).id
	case "value":
		return b.(*Default).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in Default", name))
	}
}

func (i DefaultInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Default)
	switch name {
	case "value":
		b.value = value
	}
	return nil
}

func (i DefaultInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
