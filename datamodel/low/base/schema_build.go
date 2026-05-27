// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"context"

	"github.com/pb33f/libopenapi/index"
	"go.yaml.in/yaml/v4"
)

// Build will perform a number of operations.
// Extraction of the following happens in this method:
//   - Extensions
//   - Type
//   - ExclusiveMinimum and ExclusiveMaximum
//   - Examples
//   - AdditionalProperties
//   - Discriminator
//   - ExternalDocs
//   - XML
//   - Properties
//   - AllOf, OneOf, AnyOf
//   - Not
//   - Items
//   - PrefixItems
//   - If
//   - Else
//   - Then
//   - DependentSchemas
//   - PatternProperties
//   - PropertyNames
//   - UnevaluatedItems
//   - UnevaluatedProperties
//   - Anchor
func (s *Schema) Build(ctx context.Context, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}
