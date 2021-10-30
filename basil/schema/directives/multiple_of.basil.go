// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/opsidian/conflow/basil"
	"github.com/opsidian/conflow/basil/schema"
)

// MultipleOfInterpreter is the basil interpreter for the MultipleOf block
type MultipleOfInterpreter struct {
	s schema.Schema
}

func (i MultipleOfInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "MultipleOf",
			Properties: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
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

// Create creates a new MultipleOf block
func (i MultipleOfInterpreter) CreateBlock(id basil.ID, blockCtx *basil.BlockContext) basil.Block {
	return &MultipleOf{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i MultipleOfInterpreter) ValueParamName() basil.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i MultipleOfInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *MultipleOf
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i MultipleOfInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*MultipleOf).id
	case "value":
		return b.(*MultipleOf).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in MultipleOf", name))
	}
}

func (i MultipleOfInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*MultipleOf)
	switch name {
	case "value":
		b.value = value
	}
	return nil
}

func (i MultipleOfInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
