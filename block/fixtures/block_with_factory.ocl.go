// Code generated by ocl generate. DO NOT EDIT.
package fixtures

import (
	"fmt"

	"github.com/opsidian/ocl/ocl"
	"github.com/opsidian/ocl/util"
	"github.com/opsidian/parsley/parsley"
)

// NewBlockWithFactoryFactory creates a new BlockWithFactory block factory
func NewBlockWithFactoryFactory(
	typeNode parsley.Node,
	idNode parsley.Node,
	paramNodes map[string]parsley.Node,
	blockNodes []parsley.Node,
) (ocl.BlockFactory, parsley.Error) {
	return &BlockWithFactoryFactory{
		typeNode:   typeNode,
		idNode:     idNode,
		paramNodes: paramNodes,
		blockNodes: blockNodes,
	}, nil
}

// BlockWithFactoryFactory will create and evaluate a BlockWithFactory block
type BlockWithFactoryFactory struct {
	typeNode    parsley.Node
	idNode      parsley.Node
	paramNodes  map[string]parsley.Node
	blockNodes  []parsley.Node
	shortFormat bool
}

// CreateBlock creates a new BlockWithFactory block
func (f *BlockWithFactoryFactory) CreateBlock(parentCtx interface{}) (ocl.Block, interface{}, parsley.Error) {
	var err parsley.Error

	block := &BlockWithFactory{}

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
			case *BlockSimpleFactory:
				block.BlockFactories = append(block.BlockFactories, b)
			default:
				panic(fmt.Sprintf("block type %T is not supported in BlockWithFactory, please open a bug ticket", childBlockFactory))
			}

		}
	}

	return block, ctx, nil
}

// EvalBlock evaluates all fields belonging to the given stage on a BlockWithFactory block
func (f *BlockWithFactoryFactory) EvalBlock(ctx interface{}, stage string, res ocl.Block) parsley.Error {
	var err parsley.Error

	if preInterpreter, ok := res.(ocl.BlockPreInterpreter); ok {
		if err = preInterpreter.PreEval(ctx, stage); err != nil {
			return err
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
func (f *BlockWithFactoryFactory) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (f *BlockWithFactoryFactory) HasShortFormat() bool {
	return false
}
