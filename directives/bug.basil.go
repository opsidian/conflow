// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/opsidian/conflow/basil/schema"
	"github.com/opsidian/conflow/conflow"
)

// BugInterpreter is the basil interpreter for the Bug block
type BugInterpreter struct {
	s schema.Schema
}

func (i BugInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Metadata: schema.Metadata{
				Annotations: map[string]string{"eval_stage": "ignore"},
			},
			Name: "Bug",
			Properties: map[string]schema.Schema{
				"description": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"value": "true"},
					},
				},
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
			},
			Required: []string{"description"},
		}
	}
	return i.s
}

// Create creates a new Bug block
func (i BugInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &Bug{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BugInterpreter) ValueParamName() conflow.ID {
	return "description"
}

// ParseContext returns with the parse context for the block
func (i BugInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Bug
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BugInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "description":
		return b.(*Bug).description
	case "id":
		return b.(*Bug).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in Bug", name))
	}
}

func (i BugInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Bug)
	switch name {
	case "description":
		b.description = value.(string)
	}
	return nil
}

func (i BugInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
