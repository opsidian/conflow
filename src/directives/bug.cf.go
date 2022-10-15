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
				annotations.EvalStage: "ignore",
				annotations.Type:      "directive",
			},
			ID: "github.com/conflowio/conflow/src/directives.Bug",
		},
		Name: "Bug",
		Parameters: map[string]schema.Schema{
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

// BugInterpreter is the Conflow interpreter for the Bug block
type BugInterpreter struct {
}

func (i BugInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/directives.Bug")
	return s
}

// Create creates a new Bug block
func (i BugInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &Bug{}
	b.id = id
	return b
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

func (i BugInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
