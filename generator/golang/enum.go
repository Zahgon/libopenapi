// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

import (
	"go.yaml.in/yaml/v4"
)

type enumShape struct {
	goType        string
	constants     bool
	mixed         bool
	nullable      bool
	nonNullValues int
}

func enumShapeFor(nodes []*yaml.Node) enumShape { _ = "STUB: not implemented"; return *new(enumShape) }

func enumFamily(node *yaml.Node) string { _ = "STUB: not implemented"; return "" }

func enumHasNull(nodes []*yaml.Node) bool { _ = "STUB: not implemented"; return false }

func enumLiteral(node *yaml.Node, goType string) string { _ = "STUB: not implemented"; return "" }

func nodeIsNull(node *yaml.Node) bool { _ = "STUB: not implemented"; return false }
