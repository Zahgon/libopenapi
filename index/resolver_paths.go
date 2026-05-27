// Copyright 2022-2026 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package index

import (
	"go.yaml.in/yaml/v4"
)

func (resolver *Resolver) buildDefPath(ref *Reference, l string) string {
	_ = "STUB: not implemented"
	return ""
}

// Reference contains a fragment (e.g. "file.yaml#/components/schemas/Foo" or "#/definitions/Bar").

// Has a file/URL portion before the fragment.

// Relative file path — resolve against the parent ref's location.

// Parent is a remote URL: resolve relative path against the URL's directory.

// Parent is a local file path: resolve relative to the parent's directory.

// Absolute HTTP URL with a fragment — use as-is.

// HTTP URL with no fragment content — use just the URL part.

// Fragment-only ref (e.g. "#/components/schemas/Foo") with a remote parent.

// Fragment-only ref with a fragment-only parent — keep as local fragment.

// Fragment-only ref with a local file parent — prepend the file portion.

// No fragment, absolute HTTP URL — use as-is.

// No fragment, relative path with a remote parent — resolve against the URL.

// No fragment, local relative path — resolve against the parent's directory.

func (resolver *Resolver) resolveLocalRefPath(base, ref string) string {
	_ = "STUB: not implemented"
	return ""
}

func (resolver *Resolver) buildDefPathWithSchemaBase(ref *Reference, l string, schemaIDBase string) string {
	_ = "STUB: not implemented"
	return ""
}

func (resolver *Resolver) resolveSchemaIdBase(parentBase string, node *yaml.Node) string {
	_ = "STUB: not implemented"
	return ""
}

// ResolvePendingNodes applies deferred node content replacements that were collected during resolution.
func (resolver *Resolver) ResolvePendingNodes() { _ = "STUB: not implemented"; return }
