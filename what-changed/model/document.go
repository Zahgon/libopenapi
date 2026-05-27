// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

// Package model
//
// What-changed models are unified across OpenAPI and Swagger. Everything is kept flat for simplicity, so please
// excuse the size of the package. There is a lot of data to crunch!
//
// Every model in here is either universal (works across both versions of OpenAPI) or is bound to a specific version
// of OpenAPI. There is only a single model however - so version specific objects are marked accordingly.
package model

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
)

// DocumentChanges represents all the changes made to an OpenAPI document.
type DocumentChanges struct {
	*PropertyChanges
	InfoChanges                *InfoChanges                  `json:"info,omitempty" yaml:"info,omitempty"`
	PathsChanges               *PathsChanges                 `json:"paths,omitempty" yaml:"paths,omitempty"`
	TagChanges                 []*TagChanges                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocChanges         *ExternalDocChanges           `json:"externalDoc,omitempty" yaml:"externalDoc,omitempty"`
	WebhookChanges             map[string]*PathItemChanges   `json:"webhooks,omitempty" yaml:"webhooks,omitempty"`
	ServerChanges              []*ServerChanges              `json:"servers,omitempty" yaml:"servers,omitempty"`
	SecurityRequirementChanges []*SecurityRequirementChanges `json:"securityRequirements,omitempty" yaml:"securityRequirements,omitempty"`
	ComponentsChanges          *ComponentsChanges            `json:"components,omitempty" yaml:"components,omitempty"`
	ExtensionChanges           *ExtensionChanges             `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// TotalChanges returns a total count of all changes made in the Document
func (d *DocumentChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// GetAllChanges returns a slice of all changes made between Document objects
func (d *DocumentChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalBreakingChanges returns a total count of all breaking changes made in the Document
func (d *DocumentChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareDocuments will compare any two OpenAPI documents (either Swagger or OpenAPI) and return a pointer to
// DocumentChanges that outlines everything that was found to have changed.
func CompareDocuments(l, r any) *DocumentChanges { _ = "STUB: not implemented"; return nil }

// reset schema hashmap

// clear hash cache to ensure clean state for comparison

// version

// host

// base path

// schemes

// consumes

// produces

// tags

// paths

// external docs

// info

// security

// components / definitions
// swagger (damn you) decided to put all this stuff at the document root, rather than cleanly
// placing it under a parent, like they did with OpenAPI. This means picking through each definition
// creating a new set of changes and then morphing them into a single changes object.

// version

// schema dialect

// $self field (3.2+)

// tags

// paths

// external docs

// info

// security

// compare components.

// compare servers

// compare webhooks

// extensions

func compareDocumentExternalDocs(l, r low.HasExternalDocs, dc *DocumentChanges, changes *[]*Change) {
	_ = "STUB: not implemented"
	// external docs
	return
}

func compareDocumentInfo(l, r *low.NodeReference[*base.Info], dc *DocumentChanges, changes *[]*Change) {
	_ = "STUB: not implemented"
	// info
	return
}
