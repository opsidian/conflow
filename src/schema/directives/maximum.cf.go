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
			ID: "github.com/conflowio/conflow/src/schema/directives.Maximum",
		},
		Properties: map[string]schema.Schema{
			"id": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
			"value": &schema.Untyped{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.Value: "true",
					},
				},
			},
		},
		Required: []string{"value"},
	})
}

// MaximumInterpreter is the Conflow interpreter for the Maximum block
type MaximumInterpreter struct {
}

func (i MaximumInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/schema/directives.Maximum")
	return s
}

// Create creates a new Maximum block
func (i MaximumInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &Maximum{}
	b.id = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i MaximumInterpreter) ValueParamName() conflow.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i MaximumInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Maximum
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i MaximumInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*Maximum).id
	case "value":
		return b.(*Maximum).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in Maximum", name))
	}
}

func (i MaximumInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Maximum)
	switch name {
	case "value":
		b.value = value
	}
	return nil
}

func (i MaximumInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
