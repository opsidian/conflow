// Code generated by Conflow. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/conflowio/conflow/pkg/conflow"
	"github.com/conflowio/conflow/pkg/conflow/annotations"
	"github.com/conflowio/conflow/pkg/schema"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{
				annotations.EvalStage: "ignore",
				annotations.Type:      "directive",
			},
			ID: "github.com/conflowio/conflow/pkg/directives.Deprecated",
		},
		Properties: map[string]schema.Schema{
			"description": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.Value: "true",
					},
				},
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
		Required: []string{"description"},
	})
}

// NewDeprecatedWithDefaults creates a new Deprecated instance with default values
func NewDeprecatedWithDefaults() *Deprecated {
	b := &Deprecated{}
	return b
}

// DeprecatedInterpreter is the Conflow interpreter for the Deprecated block
type DeprecatedInterpreter struct {
}

func (i DeprecatedInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/pkg/directives.Deprecated")
	return s
}

// Create creates a new Deprecated block
func (i DeprecatedInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := NewDeprecatedWithDefaults()
	b.id = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i DeprecatedInterpreter) ValueParamName() conflow.ID {
	return "description"
}

// ParseContext returns with the parse context for the block
func (i DeprecatedInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Deprecated
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i DeprecatedInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "description":
		return b.(*Deprecated).description
	case "id":
		return b.(*Deprecated).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in Deprecated", name))
	}
}

func (i DeprecatedInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Deprecated)
	switch name {
	case "description":
		b.description = schema.Value[string](value)
	}
	return nil
}

func (i DeprecatedInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
