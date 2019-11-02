// Code generated by Basil. DO NOT EDIT.
package blocks

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type PrintInterpreter struct{}

// Create creates a new Print block
func (i PrintInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Print{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i PrintInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"value": {
			Type:       "interface{}",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: true,
			IsOutput:   false,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i PrintInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i PrintInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i PrintInterpreter) ValueParamName() basil.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i PrintInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Print
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i PrintInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Print).id
	case "value":
		return b.(*Print).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in Print", name))
	}
}

func (i PrintInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*Print)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	case "value":
		b.value, err = variable.AnyValue(value)
	}
	return err
}

func (i PrintInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i PrintInterpreter) EvalStage() basil.EvalStage {
	var nilBlock *Print
	if b, ok := basil.Block(nilBlock).(basil.EvalStageAware); ok {
		return b.EvalStage()
	}

	return basil.EvalStageUndefined
}
