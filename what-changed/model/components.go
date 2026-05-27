// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/orderedmap"
)

// ComponentsChanges represents changes made to both OpenAPI and Swagger documents. This model is based on OpenAPI 3
// components, however it's also used to contain Swagger definitions changes. Swagger for some reason decided to not
// contain definitions inside a single parent like Components, and instead scattered them across the root of the
// Swagger document, giving everything a `Definitions` postfix. This design attempts to unify those models into
// a single entity that contains all changes.
//
// Schemas are treated differently from every other component / definition in this library. Schemas can be highly
// recursive, and are not resolved by the model, every ref is recorded, but it's not looked at essentially. This means
// that when what-changed performs a check, everything that is *not* a schema is checked *inline*, Those references are
// resolved in place and a change is recorded in place. Schemas however are *not* resolved. which means no change
// will be recorded in place for any object referencing it.
//
// That is why there is a separate SchemaChanges object in ComponentsChanges. Schemas are checked at the source, and
// not inline when referenced. A schema change will only be found once, however a change to ANY other definition or
// component, will be found inline (and will duplicate for every use).
//
// The other oddity here is SecuritySchemes. For some reason OpenAPI does not use a $ref for these entities, it
// uses a name lookup, which means there are no direct links between any model and a security scheme reference.
// So like Schemas, SecuritySchemes are treated differently and handled individually.
//
// An important note: Everything EXCEPT Schemas and SecuritySchemes is ONLY checked for additions or removals.
// modifications are not checked, these checks occur in-place by implementing objects as they are autp-resolved
// when the model is built.
type ComponentsChanges struct {
	*PropertyChanges
	SchemaChanges         map[string]*SchemaChanges         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	SecuritySchemeChanges map[string]*SecuritySchemeChanges `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	ExtensionChanges      *ExtensionChanges                 `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// CompareComponents will compare OpenAPI components for any changes. Accepts Swagger Definition objects
// like ParameterDefinitions or Definitions etc.
func CompareComponents(l, r any) *ComponentsChanges { _ = "STUB: not implemented"; return nil }

// Swagger Parameters

// Swagger Responses

// Swagger Schemas

// Swagger Security Definitions

// OpenAPI Components

//if low.AreEqual(lComponents, rComponents) {
//	return nil
//}

// run as fast as we can, thread all the things.

type componentComparison struct {
	prop   string
	result any
}

// run a generic comparison in a thread which in turn splits checks into further threads.
func runComparison[T any, R any](l, r *orderedmap.Map[low.KeyReference[string], low.ValueReference[T]],
	changes *[]*Change, label string, compareFunc func(l, r T) R, doneChan chan componentComparison,
) {
	_ = "STUB: not implemented"
	// for schemas
	return
}

// GetAllChanges returns a slice of all changes made between Callback objects
func (c *ComponentsChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns total changes for all Components and Definitions
func (c *ComponentsChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns all breaking changes found for all Components and Definitions
func (c *ComponentsChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }
