// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	lowv3 "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// buildLowLink builds a low-level Link from a resolved YAML node.
func buildLowLink(node *yaml.Node, idx *index.SpecIndex) (*lowv3.Link, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Link represents a high-level OpenAPI 3+ Link object that is backed by a low-level one.
//
// The Link object represents a possible design-time link for a response. The presence of a link does not guarantee the
// caller's ability to successfully invoke it, rather it provides a known relationship and traversal mechanism between
// responses and other operations.
//
// Unlike dynamic links (i.e. links provided in the response payload), the OAS linking mechanism does not require
// link information in the runtime response.
//
// For computing links, and providing instructions to execute them, a runtime expression is used for accessing values
// in an operation and using them as parameters while invoking the linked operation.
//   - https://spec.openapis.org/oas/v3.1.0#link-object
type Link struct {
	Reference    string                              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	OperationRef string                              `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	OperationId  string                              `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   *orderedmap.Map[string, string]     `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  string                              `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Description  string                              `json:"description,omitempty" yaml:"description,omitempty"`
	Server       *Server                             `json:"server,omitempty" yaml:"server,omitempty"`
	Extensions   *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low          *lowv3.Link
}

// NewLink will create a new high-level Link instance from a low-level one.
func NewLink(link *lowv3.Link) *Link { _ = "STUB: not implemented"; return nil }

// GoLow will return the low-level Link instance used to create the high-level one.
func (l *Link) GoLow() *lowv3.Link {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Link instance that was used to create the high-level one, with no type
	return nil
}

func (l *Link) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// IsReference returns true if this Link is a reference to another Link definition.
	return *new(any)
}

func (l *Link) IsReference() bool { _ = "STUB: not implemented"; return false }

// GetReference returns the reference string if this is a reference Link.
func (l *Link) GetReference() string {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the Link object as a byte slice.
	return ""
}

func (l *Link) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the Link object.
		nil
}

func (l *Link) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	// Handle reference-only link
	return nil, nil
}

// MarshalYAMLInline will create a ready to render YAML representation of the Link object,
// with all references resolved inline.
func (l *Link) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	// reference-only objects render as $ref nodes
	return nil, nil
}

// resolve external reference if present

// buildLowLink never returns an error, so we can ignore it

// MarshalYAMLInlineWithContext will create a ready to render YAML representation of the Link object,
// resolving any references inline where possible. Uses the provided context for cycle detection.
// The ctx parameter should be *base.InlineRenderContext but is typed as any to satisfy the
// high.RenderableInlineWithContext interface without import cycles.
func (l *Link) MarshalYAMLInlineWithContext(ctx any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolve external reference if present

// buildLowLink never returns an error, so we can ignore it

// CreateLinkRef creates a Link that renders as a $ref to another link definition.
// This is useful when building OpenAPI specs programmatically, and you want to reference
// a link defined in components/links rather than inlining the full definition.
//
// Example:
//
//	link := v3.CreateLinkRef("#/components/links/GetUserByUserId")
//
// Renders as:
//
//	$ref: '#/components/links/GetUserByUserId'
func CreateLinkRef(ref string) *Link { _ = "STUB: not implemented"; return nil }
