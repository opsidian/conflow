// Code generated by Basil. DO NOT EDIT.
package directives

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type InputInterpreter struct{}

// Create creates a new Input block
func (i InputInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Input{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i InputInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"type": {
			Type:       "string",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: true,
			IsOutput:   false,
		},
		"required": {
			Type:       "bool",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i InputInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i InputInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i InputInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i InputInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Input
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i InputInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Input).id
	case "type":
		return b.(*Input).inputType
	case "required":
		return b.(*Input).required
	default:
		panic(fmt.Errorf("unexpected parameter %q in Input", name))
	}
}

func (i InputInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*Input)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	case "type":
		b.inputType, err = variable.StringValue(value)
	case "required":
		b.required, err = variable.BoolValue(value)
	}
	return err
}

func (i InputInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i InputInterpreter) EvalStage() basil.EvalStage {
	var nilBlock *Input
	if b, ok := basil.Block(nilBlock).(basil.EvalStageAware); ok {
		return b.EvalStage()
	}

	return basil.EvalStageUndefined
}
