// Code generated by Basil. DO NOT EDIT.
package directives

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type BugInterpreter struct{}

// Create creates a new Bug block
func (i BugInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Bug{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i BugInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"description": {
			Type:       "string",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i BugInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i BugInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i BugInterpreter) ValueParamName() basil.ID {
	return "description"
}

// ParseContext returns with the parse context for the block
func (i BugInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Bug
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BugInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Bug).id
	case "description":
		return b.(*Bug).description
	default:
		panic(fmt.Errorf("unexpected parameter %q in Bug", name))
	}
}

func (i BugInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*Bug)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	case "description":
		b.description, err = variable.StringValue(value)
	}
	return err
}

func (i BugInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
