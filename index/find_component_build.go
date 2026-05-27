// Copyright 2023-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"go.yaml.in/yaml/v4"
)

func cloneFoundComponentReference(index *SpecIndex, found *Reference, componentID, absoluteFilePath string) *Reference {
	_ = "STUB: not implemented"
	return nil
}

// buildResolvedComponentReference constructs a fully resolved Reference for a component.
// source is the original unresolved ref (may be nil for fresh lookups); componentID is the
// JSON Pointer fragment (e.g. "#/components/schemas/Pet"); absoluteFilePath is the file
// or URL where the component lives.
func buildResolvedComponentReference(
	index *SpecIndex,
	source *Reference,
	componentID, absoluteFilePath, name, path string,
	node *yaml.Node,
) *Reference {
	_ = "STUB: not implemented"
	return nil
}

func lookupComponentParentNode(index *SpecIndex, componentID, fullDef string) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

func cloneRequiredRefProperties(source map[string][]string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func cloneSiblingProperties(source map[string]*yaml.Node) map[string]*yaml.Node {
	_ = "STUB: not implemented"
	return nil
}
