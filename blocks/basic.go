// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package blocks

import "github.com/opsidian/conflow/basil"

// @block
type Basic struct {
	// @id
	id basil.ID
}

func (b *Basic) ID() basil.ID {
	return b.id
}
