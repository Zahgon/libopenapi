// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Operation represents a low-level Swagger / OpenAPI 2 Operation object.
//
// It describes a single API operation on a path.
//   - https://swagger.io/specification/v2/#operationObject
type Operation struct {
	Tags         low.NodeReference[[]low.ValueReference[string]]
	Summary      low.NodeReference[string]
	Description  low.NodeReference[string]
	ExternalDocs low.NodeReference[*base.ExternalDoc]
	OperationId  low.NodeReference[string]
	Consumes     low.NodeReference[[]low.ValueReference[string]]
	Produces     low.NodeReference[[]low.ValueReference[string]]
	Parameters   low.NodeReference[[]low.ValueReference[*Parameter]]
	Responses    low.NodeReference[*Responses]
	Schemes      low.NodeReference[[]low.ValueReference[string]]
	Deprecated   low.NodeReference[bool]
	Security     low.NodeReference[[]low.ValueReference[*base.SecurityRequirement]]
	Extensions   *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
}

// Build will extract external docs, extensions, parameters, responses and security requirements.
func (o *Operation) Build(ctx context.Context, _, root *yaml.Node, idx *index.SpecIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// extract externalDocs

// extract parameters

// extract responses

// extract security

// Hash will return a consistent Hash of the Operation object
func (o *Operation) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// methods to satisfy swagger operations interface

func (o *Operation) GetTags() low.NodeReference[[]low.ValueReference[string]] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operation) GetSummary() low.NodeReference[string] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetDescription() low.NodeReference[string] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operation) GetExternalDocs() low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetOperationId() low.NodeReference[string] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operation) GetDeprecated() low.NodeReference[bool] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operation) GetResponses() low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetParameters() low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetSecurity() low.NodeReference[any] { _ = "STUB: not implemented"; return nil }

func (o *Operation) GetSchemes() low.NodeReference[[]low.ValueReference[string]] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operation) GetProduces() low.NodeReference[[]low.ValueReference[string]] {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operation) GetConsumes() low.NodeReference[[]low.ValueReference[string]] {
	_ = "STUB: not implemented"
	return nil
}
