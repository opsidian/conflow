// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/opsidian/conflow/basil"
	"github.com/opsidian/conflow/basil/schema"
)

// ConstInterpreter is the basil interpreter for the Const block
type ConstInterpreter struct {
	s schema.Schema
}

func (i ConstInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "Const",
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

// Create creates a new Const block
func (i ConstInterpreter) CreateBlock(id basil.ID, blockCtx *basil.BlockContext) basil.Block {
	return &Const{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i ConstInterpreter) ValueParamName() basil.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i ConstInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Const
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i ConstInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Const).id
	case "value":
		return b.(*Const).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in Const", name))
	}
}

func (i ConstInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Const)
	switch name {
	case "value":
		b.value = value
	}
	return nil
}

func (i ConstInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
