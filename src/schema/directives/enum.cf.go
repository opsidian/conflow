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
				annotations.Type: "directive",
			},
			ID: "github.com/conflowio/conflow/src/schema/directives.Enum",
		},
		Parameters: map[string]schema.Schema{
			"id": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
			"values": &schema.Array{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.Value: "true",
					},
				},
				Items: &schema.Untyped{},
			},
		},
		Required: []string{"values"},
	})
}

// EnumInterpreter is the Conflow interpreter for the Enum block
type EnumInterpreter struct {
}

func (i EnumInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/schema/directives.Enum")
	return s
}

// Create creates a new Enum block
func (i EnumInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &Enum{}
	b.id = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i EnumInterpreter) ValueParamName() conflow.ID {
	return "values"
}

// ParseContext returns with the parse context for the block
func (i EnumInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Enum
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i EnumInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*Enum).id
	case "values":
		return b.(*Enum).values
	default:
		panic(fmt.Errorf("unexpected parameter %q in Enum", name))
	}
}

func (i EnumInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Enum)
	switch name {
	case "values":
		b.values = value.([]interface{})
	}
	return nil
}

func (i EnumInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
