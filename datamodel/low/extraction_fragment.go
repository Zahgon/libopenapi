// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package low

import (
	"go.yaml.in/yaml/v4"
)

// navigateReferenceFragment navigates a local JSON Pointer fragment within a YAML node tree.
// Supported fragment formats are "#/path/to/node", "/path/to/node", and "#/" for the root.
func navigateReferenceFragment(root *yaml.Node, fragment string) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

func lookupFragmentMapValue(node *yaml.Node, key string) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

func lookupFragmentSequenceValue(node *yaml.Node, segment string) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}
