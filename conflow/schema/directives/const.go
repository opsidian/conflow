// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package directives

import (
	"fmt"

	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
)

// @block "directive"
type Const struct {
	// @id
	id conflow.ID
	// @value
	// @required
	value interface{}
}

func (c *Const) ID() conflow.ID {
	return c.id
}

func (c *Const) ApplyToSchema(s schema.Schema) error {
	value, err := s.ValidateValue(c.value)
	if err != nil {
		return fmt.Errorf("const value is invalid: %w", err)
	}

	switch st := s.(type) {
	case *schema.Array:
		st.Const = value.([]interface{})
	case *schema.Boolean:
		st.Const = schema.BooleanPtr(value.(bool))
	case *schema.Integer:
		st.Const = schema.IntegerPtr(value.(int64))
	case *schema.Map:
		st.Const = value.(map[string]interface{})
	case *schema.Number:
		st.Const = schema.NumberPtr(value.(float64))
	case *schema.String:
		st.Const = schema.StringPtr(value.(string))
	case *schema.Object:
		st.Const = schema.ObjectPtr(value.(map[string]interface{}))
	default:
		return fmt.Errorf("const directive can not be applied to %T", s)
	}

	return nil
}
