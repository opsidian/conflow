// Code generated by Basil. DO NOT EDIT.
package main

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type BenchmarkInterpreter struct{}

// Create creates a new Benchmark block
func (i BenchmarkInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Benchmark{
		id:  id,
		run: make(chan basil.BlockMessage, 0),
	}
}

// Params returns with the list of valid parameters
func (i BenchmarkInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"duration": {
			Type:       "time.Duration",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: true,
			IsOutput:   false,
		},
		"elapsed": {
			Type:       "time.Duration",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"counter": {
			Type:       "int64",
			EvalStage:  basil.EvalStages["close"],
			IsRequired: false,
			IsOutput:   true,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i BenchmarkInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return map[basil.ID]basil.BlockDescriptor{
		"run": {
			EvalStage:  basil.EvalStages["close"],
			IsRequired: true,
			IsOutput:   true,
			IsMany:     false,
		},
	}
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i BenchmarkInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i BenchmarkInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BenchmarkInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Benchmark
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BenchmarkInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Benchmark).id
	case "duration":
		return b.(*Benchmark).duration
	case "elapsed":
		return b.(*Benchmark).elapsed
	case "counter":
		return b.(*Benchmark).counter
	default:
		panic(fmt.Errorf("unexpected parameter %q in Benchmark", name))
	}
}

func (i BenchmarkInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*Benchmark)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	case "duration":
		b.duration, err = variable.TimeDurationValue(value)
	case "elapsed":
		b.elapsed, err = variable.TimeDurationValue(value)
	}
	return err
}

func (i BenchmarkInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i BenchmarkInterpreter) ProcessChannels(blockContainer basil.BlockContainer) {
	b := blockContainer.Block().(*Benchmark)
	go func() {
		for blockMessage := range b.run {
			blockContainer.PublishBlock("run", blockMessage)
		}
	}()
}

func (i BenchmarkInterpreter) CloseChannels(blockContainer basil.BlockContainer) {
	b := blockContainer.Block().(*Benchmark)
	close(b.run)
}
