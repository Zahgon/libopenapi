// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	low "github.com/pb33f/libopenapi/datamodel/low/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Arazzo represents a high-level Arazzo document.
// https://spec.openapis.org/arazzo/v1.0.1
type Arazzo struct {
	Arazzo             string                              `json:"arazzo,omitempty" yaml:"arazzo,omitempty"`
	Info               *Info                               `json:"info,omitempty" yaml:"info,omitempty"`
	SourceDescriptions []*SourceDescription                `json:"sourceDescriptions,omitempty" yaml:"sourceDescriptions,omitempty"`
	Workflows          []*Workflow                         `json:"workflows,omitempty" yaml:"workflows,omitempty"`
	Components         *Components                         `json:"components,omitempty" yaml:"components,omitempty"`
	Extensions         *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	openAPISourceDocs  []*v3.Document
	low                *low.Arazzo
}

// NewArazzo creates a new high-level Arazzo instance from a low-level one.
func NewArazzo(a *low.Arazzo) *Arazzo { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level Arazzo instance used to create the high-level one.
func (a *Arazzo) GoLow() *low.Arazzo {
	_ = "STUB: not implemented"

	// GoLowUntyped returns the low-level Arazzo instance with no type.
	return nil
}

func (a *Arazzo) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// AddOpenAPISourceDocument attaches one or more OpenAPI source documents to this Arazzo model.
	// Attached documents are runtime metadata and are not rendered or serialized.
	return *new(any)
}

func (a *Arazzo) AddOpenAPISourceDocument(docs ...*v3.Document) { _ = "STUB: not implemented"; return }

// GetOpenAPISourceDocuments returns attached OpenAPI source documents.
func (a *Arazzo) GetOpenAPISourceDocuments() []*v3.Document { _ = "STUB: not implemented"; return nil }

// Render returns a YAML representation of the Arazzo object as a byte slice.
func (a *Arazzo) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML creates a ready to render YAML representation of the Arazzo object.
		nil
}

func (a *Arazzo) MarshalYAML() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
