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

// Items is a low-level representation of a Swagger / OpenAPI 2 Items object.
//
// Items is a limited subset of JSON-Schema's items object. It is used by parameter definitions that are not
// located in "body". Items, is actually identical to a Header, except it does not have description.
//   - https://swagger.io/specification/v2/#itemsObject
type Items struct {
	Type             low.NodeReference[string]
	Format           low.NodeReference[string]
	CollectionFormat low.NodeReference[string]
	Items            low.NodeReference[*Items]
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
func (i *Items) FindExtension(ext string) *low.ValueReference[*yaml.Node] {
	_ = "STUB: not implemented"
	return nil
}

// GetExtensions returns all Items extensions and satisfies the low.HasExtensions interface.
func (i *Items) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent Hash of the Items object
}

func (itm *Items) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Build will build out items and default value.
func (i *Items) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// IsHeader compliance methods

func (i *Items) GetType() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetFormat() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetItems() *low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetCollectionFormat() *low.NodeReference[string] {
	_ = "STUB: not implemented"
	return nil
}

func (i *Items) GetDescription() *low.NodeReference[string] {
	_ = "STUB: not implemented"
	// not implemented, but required to align with header contract
	return nil
}

func (i *Items) GetDefault() *low.NodeReference[*yaml.Node] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetMaximum() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetExclusiveMaximum() *low.NodeReference[bool] {
	_ = "STUB: not implemented"
	return nil
}

func (i *Items) GetMinimum() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetExclusiveMinimum() *low.NodeReference[bool] {
	_ = "STUB: not implemented"
	return nil
}

func (i *Items) GetMaxLength() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetMinLength() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetPattern() *low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetMaxItems() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetMinItems() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetUniqueItems() *low.NodeReference[bool] { _ = "STUB: not implemented"; return nil }

func (i *Items) GetEnum() *low.NodeReference[[]low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil
}

func (i *Items) GetMultipleOf() *low.NodeReference[int] { _ = "STUB: not implemented"; return nil }
