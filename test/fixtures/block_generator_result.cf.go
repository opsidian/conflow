// Code generated by Conflow. DO NOT EDIT.

package fixtures

import (
	"fmt"
	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
)

// BlockGeneratorResultInterpreter is the conflow interpreter for the BlockGeneratorResult block
type BlockGeneratorResultInterpreter struct {
	s schema.Schema
}

func (i BlockGeneratorResultInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "BlockGeneratorResult",
			Properties: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "conflow.ID",
				},
				"value": &schema.Untyped{},
			},
		}
	}
	return i.s
}

// Create creates a new BlockGeneratorResult block
func (i BlockGeneratorResultInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &BlockGeneratorResult{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BlockGeneratorResultInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockGeneratorResultInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *BlockGeneratorResult
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockGeneratorResultInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*BlockGeneratorResult).id
	case "value":
		return b.(*BlockGeneratorResult).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockGeneratorResult", name))
	}
}

func (i BlockGeneratorResultInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*BlockGeneratorResult)
	switch name {
	case "value":
		b.value = value
	}
	return nil
}

func (i BlockGeneratorResultInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
