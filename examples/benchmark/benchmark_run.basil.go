// Code generated by Basil. DO NOT EDIT.
package main

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type BenchmarkRunInterpreter struct{}

// Create creates a new BenchmarkRun block
func (i BenchmarkRunInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &BenchmarkRun{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i BenchmarkRunInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"cnt": {
			Type:       "int64",
			EvalStage:  basil.EvalStages["close"],
			IsRequired: false,
			IsOutput:   true,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i BenchmarkRunInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i BenchmarkRunInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i BenchmarkRunInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BenchmarkRunInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *BenchmarkRun
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BenchmarkRunInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*BenchmarkRun).id
	case "cnt":
		return b.(*BenchmarkRun).cnt
	default:
		panic(fmt.Errorf("unexpected parameter %q in BenchmarkRun", name))
	}
}

func (i BenchmarkRunInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*BenchmarkRun)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	}
	return err
}

func (i BenchmarkRunInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i BenchmarkRunInterpreter) ProcessChannels(blockContainer basil.BlockContainer) {
}

func (i BenchmarkRunInterpreter) CloseChannels(blockContainer basil.BlockContainer) {
}
