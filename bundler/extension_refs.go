// Copyright 2023-2024 Princess Beef Heavy Industries, LLC / Dave Shanley
// https://pb33f.io
// SPDX-License-Identifier: MIT

package bundler

import (
	"context"

	"go.yaml.in/yaml/v4"

	"github.com/pb33f/libopenapi/index"
)

// resolveExtensionRefs resolves $ref pointers within extension fields (x-*).
// Extensions are stored as raw *yaml.Node and don't go through the MarshalYAMLInline()
// chain during rendering, so refs inside extensions need to be resolved separately.
//
// NOTE: This mutates the model's yaml.Node objects in-place.
// This follows the pattern used in compose() for composed bundling.
func resolveExtensionRefs(rolodex *index.Rolodex) { _ = "STUB: not implemented"; return }

// Process root index

// Process all external indexes

func resolveExtensionRefsFromIndex(idx *index.SpecIndex, rolodex *index.Rolodex) {
	_ = "STUB: not implemented"
	return
}

// Skip invalid refs and circular refs (already detected by indexer)

// Resolve the reference

func resolveExtensionRefContent(ctx context.Context, ref *index.Reference, _ *index.Rolodex) *yaml.Node {
	_ = "STUB: not implemented"
	// Use FindComponent which handles all reference types including:
	// - #/components/... refs (local component lookups)
	// - File refs (via lookupRolodex internally)
	// - Both YAML and raw text files
	return nil
}

// Deep copy to avoid mutating original component

// deepCopyNode creates a deep copy of a yaml.Node tree.
func deepCopyNode(node *yaml.Node) *yaml.Node { _ = "STUB: not implemented"; return nil }

// unwrap document nodes

// create copy

// deep copy children

func replaceRefNodeWithContent(refNode, content *yaml.Node) { _ = "STUB: not implemented"; return }

// replace refNode in-place with resolved content

// rewriteExtensionRefsForComposedBundle rebases $ref values found under x-* extension
// keys from their original source file location to the bundled root document.
func rewriteExtensionRefsForComposedBundle(rolodex *index.Rolodex) {
	_ = "STUB: not implemented"
	return
}

func rewriteExtensionRefsForComposedIndex(sourceIdx, rootIdx *index.SpecIndex) {
	_ = "STUB: not implemented"
	return
}

func walkAndRewriteComposedExtensionRefs(node *yaml.Node, sourceIdx, rootIdx *index.SpecIndex, inExtension bool) {
	_ = "STUB: not implemented"
	return
}

func rebaseExtensionRefForComposed(refValue string, sourceIdx, rootIdx *index.SpecIndex) string {
	_ = "STUB: not implemented"
	return ""
}

func splitRefPathAndFragment(refValue string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func isExternalRefURI(refPath string) bool { _ = "STUB: not implemented"; return false }

func specDir(idx *index.SpecIndex) string { _ = "STUB: not implemented"; return "" }
