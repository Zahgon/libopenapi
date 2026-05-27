package index

import (
	"go.yaml.in/yaml/v4"
)

// ResolveRefsInNode resolves local $ref values in a YAML node using the provided
// index. If a mapping contains sibling keys alongside $ref, sibling keys are
// preserved and merged into the resolved mapping (sibling values take precedence).
func ResolveRefsInNode(node *yaml.Node, idx *SpecIndex) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

func resolveRefsInNode(node *yaml.Node, idx *SpecIndex, seen map[string]struct{}) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// resolveRefsInMappingNode handles $ref resolution for a single mapping node, including sibling merging.
func resolveRefsInMappingNode(node *yaml.Node, idx *SpecIndex, seen map[string]struct{}) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// This helper is intentionally local-only; keep external refs intact.

// Fallback: keep original mapping (with $ref) but still resolve sibling values.

// hasNonRefSiblings returns true if the mapping node contains keys other than "$ref".
func hasNonRefSiblings(node *yaml.Node) bool { _ = "STUB: not implemented"; return false }

// findRefInMappingNode extracts the "$ref" value from a mapping node, if present.
func findRefInMappingNode(node *yaml.Node) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// extractResolvedSiblingPairs collects all non-$ref key-value pairs from a mapping, resolving their values.
func extractResolvedSiblingPairs(node *yaml.Node, idx *SpecIndex, seen map[string]struct{}) []*yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// cloneMappingNodeWithResolvedChildren shallow-clones a mapping node, recursively resolving each child value.
func cloneMappingNodeWithResolvedChildren(node *yaml.Node, idx *SpecIndex, seen map[string]struct{}) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// mergeResolvedMappingWithSiblings combines a resolved mapping with sibling key-value pairs; siblings win on conflict.
func mergeResolvedMappingWithSiblings(resolved *yaml.Node, siblings []*yaml.Node) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}
