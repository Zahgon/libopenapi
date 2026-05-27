// Copyright 2023-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"go.yaml.in/yaml/v4"
)

func (index *SpecIndex) collectInlineSchemaDefinition(parent, node *yaml.Node, seenPath []string, keyIndex int) {
	_ = "STUB: not implemented"
	return
}

func (index *SpecIndex) collectMapSchemaDefinitions(parent, node *yaml.Node, seenPath []string, keyIndex int) {
	_ = "STUB: not implemented"
	return
}

func (index *SpecIndex) collectArraySchemaDefinitions(parent, node *yaml.Node, seenPath []string, keyIndex int) {
	_ = "STUB: not implemented"
	return
}

func (index *SpecIndex) appendInlineSchemaDefinition(ref *Reference) {
	_ = "STUB: not implemented"
	return
}

func inlineSchemaIsObjectOrArray(node *yaml.Node) bool { _ = "STUB: not implemented"; return false }
