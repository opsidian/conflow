// Code generated by Basil. DO NOT EDIT.
package main

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type TickerInterpreter struct{}

// Create creates a new Ticker block
func (i TickerInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Ticker{
		id:   id,
		tick: make(chan basil.BlockMessage, 0),
	}
}

// Params returns with the list of valid parameters
func (i TickerInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"interval": {
			Type:       "time.Duration",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: true,
			IsOutput:   false,
		},
		"count": {
			Type:       "int64",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"ticks": {
			Type:       "int64",
			EvalStage:  basil.EvalStages["close"],
			IsRequired: false,
			IsOutput:   true,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i TickerInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return map[basil.ID]basil.BlockDescriptor{
		"tick": {
			EvalStage:  basil.EvalStages["close"],
			IsRequired: true,
			IsOutput:   true,
			IsMany:     false,
		},
	}
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i TickerInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i TickerInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i TickerInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Ticker
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i TickerInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Ticker).id
	case "interval":
		return b.(*Ticker).interval
	case "count":
		return b.(*Ticker).count
	case "ticks":
		return b.(*Ticker).ticks
	default:
		panic(fmt.Errorf("unexpected parameter %q in Ticker", name))
	}
}

func (i TickerInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*Ticker)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	case "interval":
		b.interval, err = variable.TimeDurationValue(value)
	case "count":
		b.count, err = variable.IntegerValue(value)
	}
	return err
}

func (i TickerInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i TickerInterpreter) ProcessChannels(blockContainer basil.BlockContainer) {
	b := blockContainer.Block().(*Ticker)
	go func() {
		for blockMessage := range b.tick {
			blockContainer.PublishBlock("tick", blockMessage)
		}
	}()
}

func (i TickerInterpreter) CloseChannels(blockContainer basil.BlockContainer) {
	b := blockContainer.Block().(*Ticker)
	close(b.tick)
}
