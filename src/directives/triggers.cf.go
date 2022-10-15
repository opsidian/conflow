// Code generated by Conflow. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/conflowio/conflow/src/conflow"
	"github.com/conflowio/conflow/src/conflow/annotations"
	"github.com/conflowio/conflow/src/schema"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{
				annotations.EvalStage: "resolve",
				annotations.Type:      "directive",
			},
			ID: "github.com/conflowio/conflow/src/directives.Triggers",
		},
		JSONPropertyNames: map[string]string{"block_ids": "blockIDs"},
		Name:              "Triggers",
		Parameters: map[string]schema.Schema{
			"block_ids": &schema.Array{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.Value: "true",
					},
				},
				Items: &schema.Untyped{},
			},
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
		Required: []string{"block_ids"},
	})
}

// TriggersInterpreter is the Conflow interpreter for the Triggers block
type TriggersInterpreter struct {
}

func (i TriggersInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/directives.Triggers")
	return s
}

// Create creates a new Triggers block
func (i TriggersInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &Triggers{}
	b.id = id
	return b
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

func (i TriggersInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
