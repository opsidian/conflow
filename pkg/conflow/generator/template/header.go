// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package template

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"
)

type HeaderParams struct {
	Package       string
	Imports       map[string]string
	LocalPrefixes []string
	Init          []string
}

const headerTemplate = `
{{ $root := . -}}

// Code generated by Conflow. DO NOT EDIT.

package {{ .Package }}

{{ if importsNotEmpty -}}
import (
	{{ $lastGroup := -1 -}}
	{{ range sortedImportKeys -}}
	{{ $currentGroup := importGroup . -}}
	{{ if and (ne $lastGroup -1) (ne $lastGroup $currentGroup) }}
	{{ end -}}
	{{ if index $root.Imports . -}}
	{{ if ne (index $root.Imports .) (last .) }}{{ index $root.Imports . }} {{ end }}{{ printf "%q" . }}
	{{ end -}}
	{{ $lastGroup = $currentGroup -}}
	{{ end -}}
)
{{ end -}}

{{ if .Init -}}
func init() {
{{ range $init := .Init -}}
{{ indent . }}
{{ end -}}
}
{{ end -}}
`

func GenerateHeader(params HeaderParams) ([]byte, error) {
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
			return sortedImportKeys(params.Imports, params.LocalPrefixes)
		},
		"last": func(path string) string {
			parts := strings.Split(path, "/")
			return parts[len(parts)-1]
		},
		"indent": func(s string) string {
			return fmt.Sprintf("\t%s", strings.ReplaceAll(s, "\n", "\n\t"))
		},
		"importGroup": func(path string) int {
			return importGroup(path, params.LocalPrefixes)
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

	return res.Bytes(), nil
}

func sortedImportKeys(imports map[string]string, localPrefixes []string) []string {
	paths := make([]string, 0, len(imports))
	for path := range imports {
		paths = append(paths, path)
	}
	sort.Slice(paths, func(i, j int) bool {
		ipath := paths[i]
		jpath := paths[j]
		igroup := importGroup(ipath, localPrefixes)
		jgroup := importGroup(jpath, localPrefixes)
		if igroup != jgroup {
			return igroup < jgroup
		}

		return ipath < jpath
	})
	return paths
}

func importGroup(importPath string, localPrefixes []string) int {
	for _, p := range localPrefixes {
		if strings.HasPrefix(importPath, p) || strings.TrimSuffix(p, "/") == importPath {
			return 2
		}
	}

	firstComponent := strings.Split(importPath, "/")[0]
	if strings.Contains(firstComponent, ".") {
		return 1
	}

	return 0
}
