// Copyright (c) 2018 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package block

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
	"github.com/opsidian/parsley/parsley"
)

// Node tokens
const (
	TokenBlock     = "BLOCK"
	TokenBlockBody = "BLOCK_BODY"
	TokenParameter = "PARAMETER"
)

var _ basil.BlockNode = &Node{}

// Node is a block node
type Node struct {
	typeNode     *basil.IDNode
	idNode       *basil.IDNode
	children     []basil.Node
	readerPos    parsley.Pos
	interpreter  basil.BlockInterpreter
	dependencies basil.Dependencies
	evalStage    basil.EvalStage
}

// ID returns with the id of the block
func (n *Node) ID() basil.ID {
	return n.idNode.ID()
}

// TypeNode returns with the type node
func (n *Node) BlockType() basil.ID {
	return n.typeNode.ID()
}

// Token returns with the node's token
func (n *Node) Token() string {
	return "BLOCK"
}

// Type returns with the node's type
func (n *Node) Type() string {
	return string(n.typeNode.ID())
}

// EvalStage returns with the evaluation stage
func (n *Node) EvalStage() basil.EvalStage {
	return n.evalStage
}

// Dependencies returns the blocks/parameters this block depends on
func (n *Node) Dependencies() basil.Dependencies {
	return n.dependencies
}

// Interpreter returns with the interpreter
func (n *Node) Interpreter() basil.BlockInterpreter {
	return n.interpreter
}

// Provides returns with the all the defined blocked node ids inside this block
func (n *Node) Provides() []basil.ID {
	var providedBlockIDs []basil.ID
	for _, c := range n.children {
		if b, ok := c.(basil.BlockNode); ok {
			providedBlockIDs = append(providedBlockIDs, b.ID())
			providedBlockIDs = append(providedBlockIDs, b.Provides()...)
		}
	}
	return providedBlockIDs
}

// StaticCheck runs static analysis on the node
func (n *Node) StaticCheck(ctx interface{}) parsley.Error {
	parseCtx := ctx.(*basil.ParseContext)

	if n.interpreter.HasForeignID() {
		if _, exists := parseCtx.BlockNode(n.ID()); !exists {
			return parsley.NewErrorf(n.idNode.Pos(), "%q is referencing a non-existing block", n.ID())
		}
	}

	params := n.interpreter.Params()
	requiredParams := make(map[basil.ID]bool, len(params))
	for name, param := range params {
		if param.IsRequired {
			requiredParams[name] = false
		}
	}

	for _, child := range n.Children() {
		switch c := child.(type) {
		case basil.BlockNode:
		case basil.ParameterNode:
			param, exists := params[c.Name()]

			switch {
			case exists && c.IsDeclaration():
				return parsley.NewErrorf(c.Pos(), "%q parameter already exists. Use \"=\" to set the parameter value or use a different name", c.Name())
			case !exists && !c.IsDeclaration():
				return parsley.NewErrorf(c.Pos(), "%q parameter does not exist", c.Name())
			case param.IsOutput:
				return parsley.NewErrorf(c.Pos(), "%q is an output parameter and can not be set", c.Name())
			}

			if err := variable.CheckNodeType(c, param.Type); err != nil {
				return err
			}

			if param.IsRequired {
				requiredParams[c.Name()] = true
			}
		}
	}

	for paramName, isSet := range requiredParams {
		if !isSet {
			return parsley.NewError(n.Pos(), fmt.Errorf("%q parameter is required", paramName))
		}
	}

	return nil
}

// Value creates a new block
func (n *Node) Value(userCtx interface{}) (interface{}, parsley.Error) {
	container := NewContainer(userCtx.(basil.EvalContext), n, nil)
	container.Run()

	return container.Value()
}

// Pos returns with the node's position
func (n *Node) Pos() parsley.Pos {
	return n.typeNode.Pos()
}

// ReaderPos returns with the reader's position
func (n *Node) ReaderPos() parsley.Pos {
	return n.readerPos
}

// SetReaderPos amends the reader position using the given function
func (n *Node) SetReaderPos(f func(parsley.Pos) parsley.Pos) {
	n.readerPos = f(n.readerPos)
}

// Children returns with the parameter and child block nodes
func (n *Node) Children() []basil.Node {
	return n.children
}

// ParamType returns with the given parameter's type if it exists, otherwise it returns false
func (n *Node) ParamType(name basil.ID) (string, bool) {
	for _, child := range n.children {
		if paramNode, ok := child.(basil.ParameterNode); ok {
			if paramNode.Name() == name {
				return paramNode.Type(), true
			}
		}
	}

	if param, ok := n.interpreter.Params()[name]; ok {
		return param.Type, true
	}

	return "", false
}

// Walk runs the given function on all child nodes
func (n *Node) Walk(f func(n parsley.Node) bool) bool {
	for _, node := range n.children {
		if parsley.Walk(node, f) {
			return true
		}
	}

	return false
}

func (n *Node) String() string {
	return fmt.Sprintf("%s{%s, %s, %s, %d..%d}", n.Token(), n.typeNode, n.idNode, n.children, n.Pos(), n.ReaderPos())
}
