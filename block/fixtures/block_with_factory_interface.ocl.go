// Code generated by ocl generate. DO NOT EDIT.
package fixtures

import (
	"fmt"
	"strings"

	"github.com/opsidian/ocl/ocl"
	"github.com/opsidian/ocl/util"
	"github.com/opsidian/parsley/parsley"
)

// NewBlockWithFactoryInterfaceFactory creates a new BlockWithFactoryInterface block factory
func NewBlockWithFactoryInterfaceFactory(
	typeNode parsley.Node,
	idNode parsley.Node,
	paramNodes map[string]parsley.Node,
	blockNodes []parsley.Node,
) (ocl.BlockFactory, parsley.Error) {
	return &BlockWithFactoryInterfaceFactory{
		typeNode:   typeNode,
		idNode:     idNode,
		paramNodes: paramNodes,
		blockNodes: blockNodes,
	}, nil
}

// BlockWithFactoryInterfaceFactory will create and evaluate a BlockWithFactoryInterface block
type BlockWithFactoryInterfaceFactory struct {
	typeNode    parsley.Node
	idNode      parsley.Node
	paramNodes  map[string]parsley.Node
	blockNodes  []parsley.Node
	shortFormat bool
}

// CreateBlock creates a new BlockWithFactoryInterface block
func (f *BlockWithFactoryInterfaceFactory) CreateBlock(parentCtx interface{}) (ocl.Block, interface{}, parsley.Error) {
	var err parsley.Error

	block := &BlockWithFactoryInterface{}

	if block.IDField, err = util.NodeStringValue(f.idNode, parentCtx); err != nil {
		return nil, nil, err
	}

	ctx := block.Context(parentCtx)

	if len(f.blockNodes) > 0 {
		var childBlockFactory interface{}
		for _, childBlock := range f.blockNodes {
			if childBlockFactory, err = childBlock.Value(ctx); err != nil {
				return nil, nil, err
			}
			switch b := childBlockFactory.(type) {
			case BlockFactoryInterface:
				block.BlockFactories = append(block.BlockFactories, b)
			default:
				panic(fmt.Sprintf("block type %T is not supported in BlockWithFactoryInterface, please open a bug ticket", childBlockFactory))
			}

		}
	}

	return block, ctx, nil
}

// EvalBlock evaluates all fields belonging to the given stage on a BlockWithFactoryInterface block
func (f *BlockWithFactoryInterfaceFactory) EvalBlock(ctx interface{}, stage string, res ocl.Block) parsley.Error {
	var err parsley.Error

	if preInterpreter, ok := res.(ocl.BlockPreInterpreter); ok {
		if err = preInterpreter.PreEval(ctx, stage); err != nil {
			return err
		}
	}

	block, ok := res.(*BlockWithFactoryInterface)
	if !ok {
		panic("result must be a type of *BlockWithFactoryInterface")
	}

	validParamNames := map[string]struct{}{
		"block_factories": struct{}{},
	}

	for paramName, paramNode := range f.paramNodes {
		if !strings.HasPrefix(paramName, "_") {
			if _, valid := validParamNames[paramName]; !valid {
				return parsley.NewError(paramNode.Pos(), fmt.Errorf("%q parameter does not exist", paramName))
			}
		}
	}

	if !f.shortFormat {
		switch stage {
		case "default":
		default:
			panic(fmt.Sprintf("unknown stage: %s", stage))
		}

		switch stage {
		case "default":
			var childBlock ocl.Block
			var childBlockCtx interface{}
			for _, childBlockFactory := range block.BlockFactories {
				if childBlock, childBlockCtx, err = childBlockFactory.CreateBlock(ctx); err != nil {
					return err
				}

				if err = childBlockFactory.EvalBlock(childBlockCtx, stage, childBlock); err != nil {
					return err
				}

				panic(fmt.Sprintf("block type %T is not supported in BlockFactories, please open a bug ticket", childBlock))
			}
		default:
			panic(fmt.Sprintf("unknown stage: %s", stage))
		}
	}

	if postInterpreter, ok := res.(ocl.BlockPostInterpreter); ok {
		if err = postInterpreter.PostEval(ctx, stage); err != nil {
			return err
		}
	}

	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (f *BlockWithFactoryInterfaceFactory) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (f *BlockWithFactoryInterfaceFactory) HasShortFormat() bool {
	return false
}
