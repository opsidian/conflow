// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

type Context struct {
	resolver Resolver
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) WithResolver(resolver Resolver) *Context {
	c.resolver = resolver
	return c
}

func (c *Context) ResolveSchema(uri string) (Schema, error) {
	if c != nil && c.resolver != nil {
		s, err := c.resolver.ResolveSchema(uri)
		if s != nil || err != nil {
			return s, err
		}
	}

	return Get(uri)
}
