// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/opsidian/conflow/basil"
	"github.com/opsidian/conflow/basil/schema"
)

// InputInterpreter is the basil interpreter for the Input block
type InputInterpreter struct {
	s schema.Schema
}

func (i InputInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Metadata: schema.Metadata{
				Annotations: map[string]string{"eval_stage": "parse"},
			},
			Name: "Input",
			Properties: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
				"required": &schema.Boolean{},
			},
		}
	}
	return i.s
}

// Create creates a new Input block
func (i InputInterpreter) CreateBlock(id basil.ID, blockCtx *basil.BlockContext) basil.Block {
	return &Input{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i InputInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i InputInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Input
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i InputInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Input).id
	case "required":
		return b.(*Input).required
	default:
		panic(fmt.Errorf("unexpected parameter %q in Input", name))
	}
}

func (i InputInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Input)
	switch name {
	case "required":
		b.required = value.(bool)
	}
	return nil
}

func (i InputInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
