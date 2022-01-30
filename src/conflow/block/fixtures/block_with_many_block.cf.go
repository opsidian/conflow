// Code generated by Conflow. DO NOT EDIT.

package fixtures

import (
	"fmt"
	"github.com/conflowio/conflow/src/conflow"
	"github.com/conflowio/conflow/src/schema"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{"block.conflow.io/type": "configuration"},
			ID:          "github.com/conflowio/conflow/src/conflow/block/fixtures.BlockWithManyBlock",
		},
		JSONPropertyNames: map[string]string{"block_simple": "BlockSimple", "id_field": "IDField"},
		Name:              "BlockWithManyBlock",
		Parameters: map[string]schema.Schema{
			"block_simple": &schema.Array{
				Items: &schema.Reference{
					Nullable: true,
					Ref:      "github.com/conflowio/conflow/src/conflow/block/fixtures.BlockSimple",
				},
			},
			"id_field": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{"block.conflow.io/id": "true"},
					ReadOnly:    true,
				},
				Format: "conflow.ID",
			},
		},
	})
}

// BlockWithManyBlockInterpreter is the Conflow interpreter for the BlockWithManyBlock block
type BlockWithManyBlockInterpreter struct {
}

func (i BlockWithManyBlockInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/conflow/block/fixtures.BlockWithManyBlock")
	return s
}

// Create creates a new BlockWithManyBlock block
func (i BlockWithManyBlockInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &BlockWithManyBlock{
		IDField: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BlockWithManyBlockInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockWithManyBlockInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *BlockWithManyBlock
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockWithManyBlockInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id_field":
		return b.(*BlockWithManyBlock).IDField
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockWithManyBlock", name))
	}
}

func (i BlockWithManyBlockInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}

func (i BlockWithManyBlockInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*BlockWithManyBlock)
	switch name {
	case "block_simple":
		b.BlockSimple = append(b.BlockSimple, value.(*BlockSimple))
	}
	return nil
}
