// Copyright 2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package index

import (
	"context"

	"go.yaml.in/yaml/v4"
)

func isSchemaContainingNode(v string) bool { _ = "STUB: not implemented"; return false }

func isMapOfSchemaContainingNode(v string) bool { _ = "STUB: not implemented"; return false }

func isArrayOfSchemaContainingNode(v string) bool { _ = "STUB: not implemented"; return false }

// underOpenAPIExamplePath reports whether seenPath is under an OpenAPI example or examples
// keyword (sample data, not schema). A segment named "example" or "examples" that is preceded
// by "properties" or "patternProperties" is a schema property name, not an OpenAPI keyword.
func underOpenAPIExamplePath(seenPath []string) bool { _ = "STUB: not implemented"; return false }

func isOpenAPIExampleKeywordSegment(seenPath []string, idx int) bool {
	_ = "STUB: not implemented"
	return false
}

// underOpenAPIExamplePayloadPath reports whether seenPath points to raw example payload content.
// A bare `example` path is not payload by itself because libopenapi still supports a direct
// `$ref` wrapper there for bundling. Once traversal moves below `example`, or into an Example
// Object's `value`/`dataValue`, the path is payload and nested `$ref` keys should be ignored.
func underOpenAPIExamplePayloadPath(seenPath []string) bool {
	_ = "STUB: not implemented"
	return false
}

// isDirectOpenAPIExampleValuePath reports whether seenPath points at the value of an OpenAPI
// `example` field itself. This is used to allow a top-level `$ref` wrapper while still skipping
// traversal into arbitrary example payload objects.
func isDirectOpenAPIExampleValuePath(seenPath []string) bool {
	_ = "STUB: not implemented"
	return false
}

// ExtractRefs will return a deduplicated slice of references for every unique ref found in the document.
// The total number of refs, will generally be much higher, you can extract those from GetRawReferenceCount()
func (index *SpecIndex) ExtractRefs(ctx context.Context, node, parent *yaml.Node, seenPath []string, level int, poly bool, pName string) []*Reference {
	_ = "STUB: not implemented"
	return nil
}
