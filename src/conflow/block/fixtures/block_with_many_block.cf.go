// Code generated by Conflow. DO NOT EDIT.

package fixtures

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
				annotations.Type: "configuration",
			},
			ID: "github.com/conflowio/conflow/src/conflow/block/fixtures.BlockWithManyBlock",
		},
		ParameterNames: map[string]string{"BlockSimple": "block_simple", "IDField": "id_field"},
		Properties: map[string]schema.Schema{
			"BlockSimple": &schema.Array{
				Items: &schema.Reference{
					Nullable: true,
					Ref:      "github.com/conflowio/conflow/src/conflow/block/fixtures.BlockSimple",
				},
			},
			"IDField": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
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
	b := &BlockWithManyBlock{}
	b.IDField = id
	return b
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

func (i BlockWithManyBlockInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	b := block.(*BlockWithManyBlock)
	switch name {
	case "block_simple":
		b.BlockSimple = append(b.BlockSimple, value.(*BlockSimple))
	}
	return nil
}
