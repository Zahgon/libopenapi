// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Header Represents a low-level Swagger / OpenAPI 2 Header object.
//
// A Header is essentially identical to a Parameter, except it does not contain 'name' or 'in' properties.
//   - https://swagger.io/specification/v2/#headerObject
type Header struct {
	Type             low.NodeReference[string]
	Format           low.NodeReference[string]
	Description      low.NodeReference[string]
	Items            low.NodeReference[*Items]
	CollectionFormat low.NodeReference[string]
	Default          low.NodeReference[*yaml.Node]
	Maximum          low.NodeReference[int]
	ExclusiveMaximum low.NodeReference[bool]
	Minimum          low.NodeReference[int]
	ExclusiveMinimum low.NodeReference[bool]
	MaxLength        low.NodeReference[int]
	MinLength        low.NodeReference[int]
	Pattern          low.NodeReference[string]
	MaxItems         low.NodeReference[int]
	MinItems         low.NodeReference[int]
	UniqueItems      low.NodeReference[bool]
	Enum             low.NodeReference[[]low.ValueReference[*yaml.Node]]
	MultipleOf       low.NodeReference[int]
	Extensions       *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
}

// FindExtension will attempt to locate an extension value using a name lookup.
func (h *Header) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all Header extensions and satisfies the low.HasExtensions interface.
func (h *Header) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Build will build out items, extensions and default value from the supplied node.
}

func (h *Header) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Hash will return a consistent Hash of the Header object
func (hdr *Header) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Getter methods to satisfy SwaggerHeader interface.

func (h *Header) GetType() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetDescription() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetFormat() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetItems() *low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetCollectionFormat() *low.NodeReference[string] {
	_ = "STUB: not implemented"
	return nil
}

func (h *Header) GetDefault() *low.NodeReference[*yaml.Node] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetMaximum() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetExclusiveMaximum() *low.NodeReference[bool] {
	_ = "STUB: not implemented"
	return nil
}

func (h *Header) GetMinimum() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetExclusiveMinimum() *low.NodeReference[bool] {
	_ = "STUB: not implemented"
	return nil
}

func (h *Header) GetMaxLength() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetMinLength() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetPattern() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetMaxItems() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetMinItems() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetUniqueItems() *low.NodeReference[bool] { _ = "STUB: not implemented"; return nil }

func (h *Header) GetEnum() *low.NodeReference[[]low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil
}

func (h *Header) GetMultipleOf() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }
