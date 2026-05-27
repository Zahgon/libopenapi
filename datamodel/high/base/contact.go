// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	low "github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// Contact represents a high-level representation of the Contact definitions found at
//
//	v2 - https://swagger.io/specification/v2/#contactObject
//	v3 - https://spec.openapis.org/oas/v3.1.0#contact-object
type Contact struct {
	Name       string                              `json:"name,omitempty" yaml:"name,omitempty"`
	URL        string                              `json:"url,omitempty" yaml:"url,omitempty"`
	Email      string                              `json:"email,omitempty" yaml:"email,omitempty"`
	Extensions *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low        *low.Contact                        `json:"-" yaml:"-"` // low-level representation
}

// NewContact will create a new Contact instance using a low-level Contact
func NewContact(contact *low.Contact) *Contact { _ = "STUB: not implemented"; return nil }

// GoLow returns the low level Contact object used to create the high-level one.
func (c *Contact) GoLow() *low.Contact {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Contact instance that was used to create the high-level one, with no type
	return nil
}

func (c *Contact) GoLowUntyped() any { _ = "STUB: not implemented"; return *new(any) }

func (c *Contact) Render() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Contact) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
