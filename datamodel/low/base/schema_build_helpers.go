// Copyright 2022-2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"context"

	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

type resolvedSchemaBuildInput struct {
	ctx         context.Context
	idx         *index.SpecIndex
	valueNode   *yaml.Node
	scopeNode   *yaml.Node
	refNode     *yaml.Node
	transformed *yaml.Node
	refLocation string
}

func buildPropertyMap(ctx context.Context, parent *Schema, root *yaml.Node, idx *index.SpecIndex, label string) (*low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[*SchemaProxy]]], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildDependentRequiredMap builds an ordered map of string arrays for the dependentRequired property
func buildDependentRequiredMap(root *yaml.Node, label string) (*low.NodeReference[*orderedmap.Map[low.KeyReference[string], low.ValueReference[[]string]]], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extract extensions from schema
func (s *Schema) extractExtensions(root *yaml.Node) { _ = "STUB: not implemented"; return }

// buildSchemaProxy builds out a SchemaProxy for a single node.
func buildSchemaProxy(ctx context.Context, idx *index.SpecIndex, kn, vn, scopeNode, rf, transformed *yaml.Node, refLocation string) low.ValueReference[*SchemaProxy] {
	_ = "STUB: not implemented"
	return nil
}

// buildSchema builds out a child schema for parent schema. Expected to be a singular schema object.
func buildSchema(ctx context.Context, labelNode, valueNode *yaml.Node, idx *index.SpecIndex) (low.ValueReference[*SchemaProxy], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildSchemaList builds out child schemas for a parent schema. Expected to be an array of schema objects.
func buildSchemaList(ctx context.Context, labelNode, valueNode *yaml.Node, idx *index.SpecIndex) ([]low.ValueReference[*SchemaProxy], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func assignBuiltSchema(ctx context.Context, labelNode, valueNode *yaml.Node, idx *index.SpecIndex, dst *low.ValueReference[*SchemaProxy]) error {
	_ = "STUB: not implemented"
	return nil
}

func assignBuiltSchemaList(ctx context.Context, labelNode, valueNode *yaml.Node, idx *index.SpecIndex, dst *[]low.ValueReference[*SchemaProxy]) error {
	_ = "STUB: not implemented"
	return nil
}

func resolveSchemaBuildInput(ctx context.Context, valueNode *yaml.Node, idx *index.SpecIndex, errFormat string) (resolvedSchemaBuildInput, error) {
	_ = "STUB: not implemented"
	return *new(resolvedSchemaBuildInput), nil
}

func schemaReferenceBuildError(errFormat string, valueNode *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}
