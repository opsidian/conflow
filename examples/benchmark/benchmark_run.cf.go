// Code generated by Conflow. DO NOT EDIT.

package main

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
				annotations.Type: "configuration",
			},
			ID: "github.com/conflowio/conflow/examples/benchmark.BenchmarkRun",
		},
		Properties: map[string]schema.Schema{
			"cnt": &schema.Integer{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.EvalStage: "close",
					},
					ReadOnly: true,
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
	})
}

// NewBenchmarkRunWithDefaults creates a new BenchmarkRun instance with default values
func NewBenchmarkRunWithDefaults() *BenchmarkRun {
	b := &BenchmarkRun{}
	return b
}

// BenchmarkRunInterpreter is the Conflow interpreter for the BenchmarkRun block
type BenchmarkRunInterpreter struct {
}

func (i BenchmarkRunInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/examples/benchmark.BenchmarkRun")
	return s
}

// Create creates a new BenchmarkRun block
func (i BenchmarkRunInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := NewBenchmarkRunWithDefaults()
	b.id = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BenchmarkRunInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BenchmarkRunInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *BenchmarkRun
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BenchmarkRunInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "cnt":
		return b.(*BenchmarkRun).cnt
	case "id":
		return b.(*BenchmarkRun).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in BenchmarkRun", name))
	}
}

func (i BenchmarkRunInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}

func (i BenchmarkRunInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
