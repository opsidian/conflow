// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/opsidian/conflow/basil/schema"
	"github.com/opsidian/conflow/conflow"
)

// TriggersInterpreter is the basil interpreter for the Triggers block
type TriggersInterpreter struct {
	s schema.Schema
}

func (i TriggersInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Metadata: schema.Metadata{
				Annotations: map[string]string{"eval_stage": "resolve"},
			},
			Name: "Triggers",
			Properties: map[string]schema.Schema{
				"block_ids": &schema.Array{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"value": "true"},
					},
					Items: &schema.Untyped{},
				},
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
			},
			PropertyNames: map[string]string{"block_ids": "blockIDs"},
			Required:      []string{"block_ids"},
		}
	}
	return i.s
}

// Create creates a new Triggers block
func (i TriggersInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &Triggers{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i TriggersInterpreter) ValueParamName() conflow.ID {
	return "block_ids"
}

// ParseContext returns with the parse context for the block
func (i TriggersInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Triggers
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i TriggersInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "block_ids":
		return b.(*Triggers).blockIDs
	case "id":
		return b.(*Triggers).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in Triggers", name))
	}
}

func (i TriggersInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Triggers)
	switch name {
	case "block_ids":
		b.blockIDs = value.([]interface{})
	}
	return nil
}

func (i TriggersInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
