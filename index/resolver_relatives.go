// Copyright 2022-2026 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package index

import (
	"go.yaml.in/yaml/v4"
)

func (resolver *Resolver) extractRelatives(
	ref *Reference,
	node, parent *yaml.Node,
	foundRelatives map[string]bool,
	journey []*Reference,
	resolve bool,
	depth int,
	schemaIDBase string,
) []*Reference {
	_ = "STUB: not implemented"
	return nil
}

func (resolver *Resolver) extractRelativesWithState(
	ref *Reference,
	node, parent *yaml.Node,
	state relativeWalkState,
) []*Reference {
	_ = "STUB: not implemented"
	// Guard against stack overflow from deeply nested or circular specs.
	// Journey tracks the reference chain (100 max); depth tracks recursive calls (500 max).
	return nil
}

func (resolver *Resolver) handleRelativeDepthLimit(ref *Reference, state relativeWalkState) []*Reference {
	_ = "STUB: not implemented"
	return nil
}

func (resolver *Resolver) extractNestedRelatives(
	ref *Reference,
	parent, node *yaml.Node,
	state relativeWalkState,
) []*Reference {
	_ = "STUB: not implemented"
	return nil
}

func (resolver *Resolver) extractRelativeReference(
	ref *Reference,
	node, parent, keyNode *yaml.Node,
	keyIndex int,
	state relativeWalkState,
) (*Reference, bool, bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

func parentArraySchemaType(parent *yaml.Node) string { _ = "STUB: not implemented"; return "" }

func shouldExtractPolymorphicRelatives(parent, keyNode *yaml.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func isInsidePropertiesNode(parent *yaml.Node) bool { _ = "STUB: not implemented"; return false }

func (resolver *Resolver) buildRelativeLookupDefinitions(ref *Reference, value, currentBase string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// Reference contains a fragment (e.g. "other.yaml#/components/schemas/Foo").

// Has a file/URL prefix before the fragment.

// Absolute HTTP URL — use as-is.

// Relative file path, but the parent ref is remote — resolve against the URL.

// Relative file path with a local parent — resolve against the parent's directory.

// Fragment-only ref (e.g. "#/definitions/Bar") — resolve against the parent's base location.

// No fragment, absolute HTTP URL — use as-is.

// No fragment, relative file path — resolve against the parent's base location.
