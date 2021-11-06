// Code generated by Conflow. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
)

// MinLengthInterpreter is the conflow interpreter for the MinLength block
type MinLengthInterpreter struct {
	s schema.Schema
}

func (i MinLengthInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "MinLength",
			Properties: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "conflow.ID",
				},
				"value": &schema.Integer{
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

// Create creates a new MinLength block
func (i MinLengthInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &MinLength{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i MinLengthInterpreter) ValueParamName() conflow.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i MinLengthInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *MinLength
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i MinLengthInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*MinLength).id
	case "value":
		return b.(*MinLength).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in MinLength", name))
	}
}

func (i MinLengthInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*MinLength)
	switch name {
	case "value":
		b.value = value.(int64)
	}
	return nil
}

func (i MinLengthInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
