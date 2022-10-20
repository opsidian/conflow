// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package template

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/conflowio/conflow/src/util"
)

type HeaderParams struct {
	Package string
	Imports map[string]string
}

const headerTemplate = `
{{ $root := . -}}

// Code generated by Conflow. DO NOT EDIT.

package {{ .Package }}

{{ if importsNotEmpty -}}
import (
	{{ range sortedImportKeys -}}
	{{ if index $root.Imports . -}}
	{{ if ne (index $root.Imports .) (last .) }}{{ index $root.Imports . }} {{ end }}{{ printf "%q" . }}
	{{ end -}}
	{{ end -}}
)
{{ end -}}
`

func GenerateHeader(params HeaderParams) (*bytes.Buffer, error) {
	headerTmpl := template.New("header")
	headerTmpl.Funcs(map[string]interface{}{
		"importsNotEmpty": func() bool {
			for _, name := range params.Imports {
				if name != "" {
					return true
				}
			}
			return false
		},
		"sortedImportKeys": func() []string {
			return util.SortedMapKeys(params.Imports)
		},
		"last": func(path string) string {
			parts := strings.Split(path, "/")
			return parts[len(parts)-1]
		},
	})
	if _, parseErr := headerTmpl.Parse(headerTemplate); parseErr != nil {
		return nil, parseErr
	}

	res := &bytes.Buffer{}
	err := headerTmpl.Execute(res, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}
