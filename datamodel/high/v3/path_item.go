// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	lowV3 "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// buildLowPathItem builds a low-level PathItem from a resolved YAML node.
func buildLowPathItem(node *yaml.Node, idx *index.SpecIndex) (*lowV3.PathItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	get = iota
	put
	post
	del
	options
	head
	patch
	trace
	query
)

// PathItem represents a high-level OpenAPI 3+ PathItem object backed by a low-level one.
//
// Describes the operations available on a single path. A Path Item MAY be empty, due to ACL constraints.
// The path itself is still exposed to the documentation viewer but they will not know which operations and parameters
// are available.
//   - https://spec.openapis.org/oas/v3.1.0#path-item-object
type PathItem struct {
	Reference            string                              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Description          string                              `json:"description,omitempty" yaml:"description,omitempty"`
	Summary              string                              `json:"summary,omitempty" yaml:"summary,omitempty"`
	Get                  *Operation                          `json:"get,omitempty" yaml:"get,omitempty"`
	Put                  *Operation                          `json:"put,omitempty" yaml:"put,omitempty"`
	Post                 *Operation                          `json:"post,omitempty" yaml:"post,omitempty"`
	Delete               *Operation                          `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options              *Operation                          `json:"options,omitempty" yaml:"options,omitempty"`
	Head                 *Operation                          `json:"head,omitempty" yaml:"head,omitempty"`
	Patch                *Operation                          `json:"patch,omitempty" yaml:"patch,omitempty"`
	Trace                *Operation                          `json:"trace,omitempty" yaml:"trace,omitempty"`
	Query                *Operation                          `json:"query,omitempty" yaml:"query,omitempty"`
	AdditionalOperations *orderedmap.Map[string, *Operation] `json:"additionalOperations,omitempty" yaml:"additionalOperations,omitempty"` // OpenAPI 3.2+ additional operations
	Servers              []*Server                           `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters           []*Parameter                        `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Extensions           *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low                  *lowV3.PathItem
}

// NewPathItem creates a new high-level PathItem instance from a low-level one.
func NewPathItem(pathItem *lowV3.PathItem) *PathItem { _ = "STUB: not implemented"; return nil }

// build operation async

// build out operations async.

// build additional operations if present

// GoLow returns the low level instance of PathItem, used to build the high-level one.
func (p *PathItem) GoLow() *lowV3.PathItem {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level PathItem instance that was used to create the high-level one, with no type
	return nil
}

func (p *PathItem) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// IsReference returns true if this PathItem is a reference to another PathItem definition.
	return *new(any)
}

func (p *PathItem) IsReference() bool { _ = "STUB: not implemented"; return false }

// GetReference returns the reference string if this is a reference PathItem.
func (p *PathItem) GetReference() string { _ = "STUB: not implemented"; return "" }

func (p *PathItem) GetOperations() *orderedmap.Map[string, *Operation] {
	_ = "STUB: not implemented"
	return nil
}

// TODO: this is a bit of a hack, but it works for now. We might just want to actually pull the data out of the document as a map and split it into the individual operations

// add additional operations if present - get line numbers from low-level KeyNodes

// find the corresponding high-level operation

// Render will return a YAML representation of the PathItem object as a byte slice.
func (p *PathItem) Render() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *PathItem) RenderInline() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalYAML will create a ready to render YAML representation of the PathItem object.
func (p *PathItem) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	// Handle reference-only path item
	return nil, nil
}

// MarshalYAMLInline will create a ready to render YAML representation of the PathItem object,
// with all references resolved inline.
func (p *PathItem) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	// reference-only objects render as $ref nodes
	return nil, nil
}

// resolve external reference if present

// MarshalYAMLInlineWithContext will create a ready to render YAML representation of the PathItem object,
// resolving any references inline where possible. Uses the provided context for cycle detection.
// The ctx parameter should be *base.InlineRenderContext but is typed as any to satisfy the
// high.RenderableInlineWithContext interface without import cycles.
func (p *PathItem) MarshalYAMLInlineWithContext(ctx any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolve external reference if present

// CreatePathItemRef creates a PathItem that renders as a $ref to another path item definition.
// This is useful when building OpenAPI specs programmatically and you want to reference
// a path item defined in components/pathItems rather than inlining the full definition.
//
// Example:
//
//	pi := v3.CreatePathItemRef("#/components/pathItems/CommonPathItem")
//
// Renders as:
//
//	$ref: '#/components/pathItems/CommonPathItem'
func CreatePathItemRef(ref string) *PathItem { _ = "STUB: not implemented"; return nil }
