// Code generated by Conflow. DO NOT EDIT.

package blocks

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
				annotations.Type: "generator",
			},
			ID: "github.com/conflowio/conflow/src/blocks.LineScanner",
		},
		Name: "LineScanner",
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
			"input": &schema.Untyped{},
			"line": &schema.Reference{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.EvalStage: "init",
						annotations.Generated: "true",
					},
				},
				Nullable: true,
				Ref:      "github.com/conflowio/conflow/src/blocks.Line",
			},
		},
		Required: []string{"input", "line"},
	})
}

// LineScannerInterpreter is the Conflow interpreter for the LineScanner block
type LineScannerInterpreter struct {
}

func (i LineScannerInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/blocks.LineScanner")
	return s
}

// Create creates a new LineScanner block
func (i LineScannerInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &LineScanner{}
	b.id = id
	b.blockPublisher = blockCtx.BlockPublisher()
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i LineScannerInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i LineScannerInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *LineScanner
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i LineScannerInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*LineScanner).id
	case "input":
		return b.(*LineScanner).input
	default:
		panic(fmt.Errorf("unexpected parameter %q in LineScanner", name))
	}
}

func (i LineScannerInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*LineScanner)
	switch name {
	case "input":
		b.input = value
	}
	return nil
}

func (i LineScannerInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	b := block.(*LineScanner)
	switch name {
	case "line":
		b.line = value.(*Line)
	}
	return nil
}
