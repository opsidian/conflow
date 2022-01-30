// Code generated by Conflow. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/conflowio/conflow/src/conflow"
	"github.com/conflowio/conflow/src/schema"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{"block.conflow.io/type": "directive"},
			ID:          "github.com/conflowio/conflow/src/schema/directives.UniqueItems",
		},
		Name: "UniqueItems",
		Parameters: map[string]schema.Schema{
			"id": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{"block.conflow.io/id": "true"},
					ReadOnly:    true,
				},
				Format: "conflow.ID",
			},
		},
	})
}

// UniqueItemsInterpreter is the Conflow interpreter for the UniqueItems block
type UniqueItemsInterpreter struct {
}

func (i UniqueItemsInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/schema/directives.UniqueItems")
	return s
}

// Create creates a new UniqueItems block
func (i UniqueItemsInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &UniqueItems{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i UniqueItemsInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i UniqueItemsInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *UniqueItems
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i UniqueItemsInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*UniqueItems).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in UniqueItems", name))
	}
}

func (i UniqueItemsInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}

func (i UniqueItemsInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
