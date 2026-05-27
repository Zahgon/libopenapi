// Copyright 2023-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"go.yaml.in/yaml/v4"
)

func (index *SpecIndex) extractReferenceAt(
	node, parent *yaml.Node,
	keyIndex int,
	seenPath []string,
	scope *SchemaIdScope,
	poly bool,
	pName string,
) *Reference {
	_ = "STUB: not implemented"
	return nil
}

func (index *SpecIndex) registerSchemaIDAt(node *yaml.Node, keyIndex int, seenPath []string, parentBaseUri string) {
	_ = "STUB: not implemented"
	return
}

func (index *SpecIndex) resolveReferenceTarget(value string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// Reference contains a fragment (e.g. "file.yaml#/components/schemas/Foo" or "#/definitions/Bar").

// Fragment-only local ref — prefix with the spec's absolute path.

// Absolute HTTP URL with fragment — use as-is.

// Absolute local file path with fragment — use as-is.

// Relative path with a configured BaseURL — resolve against the base URL.

// Relative local file path — resolve against the spec's directory.

// No fragment, absolute HTTP URL — use as-is.

// No fragment, not a bare anchor — whole-file reference or relative path.

// Spec root is remote — resolve the relative path against the remote URL.

// Relative local file path — resolve against BaseURL if configured, else the spec's directory.

func (index *SpecIndex) recordReferenceByLine(value string, line int) {
	_ = "STUB: not implemented"
	return
}

func extractSiblingRefProperties(node *yaml.Node) (map[string]*yaml.Node, []*yaml.Node) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (index *SpecIndex) storeReferenceWithSiblings(
	node, parent *yaml.Node,
	keyIndex int,
	ref *Reference,
	isExtensionPath bool,
	path, value string,
) {
	_ = "STUB: not implemented"
	return
}
