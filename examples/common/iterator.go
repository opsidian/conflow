// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package common

import (
	"context"

	"github.com/opsidian/conflow/basil"
	"github.com/opsidian/conflow/basil/block"
)

// @block
type Iterator struct {
	// @id
	id basil.ID
	// @required
	count int64
	// @generated
	it *It
	// @dependency
	blockPublisher basil.BlockPublisher
}

func (i *Iterator) ID() basil.ID {
	return i.id
}

func (it *Iterator) Run(ctx context.Context) (basil.Result, error) {
	for i := int64(0); i < it.count; i++ {
		_, err := it.blockPublisher.PublishBlock(&It{
			id:    it.it.id,
			value: i,
		}, nil)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (it *Iterator) ParseContextOverride() basil.ParseContextOverride {
	return basil.ParseContextOverride{
		BlockTransformerRegistry: block.InterpreterRegistry{
			"it": ItInterpreter{},
		},
	}
}

// @block
type It struct {
	// @id
	id basil.ID
	// @read_only
	value int64
}

func (i *It) ID() basil.ID {
	return i.id
}
