// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package directives

import (
	"github.com/opsidian/conflow/conflow"
	"github.com/opsidian/conflow/conflow/schema"
)

// @block {
//   eval_stage = "parse"
// }
type Number struct {
	schema.Number
}

func (n *Number) ApplyToParameterConfig(config *conflow.ParameterConfig) {
	config.Schema = &n.Number
}
