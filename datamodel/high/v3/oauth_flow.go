// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	lowv3 "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// OAuthFlow represents a high-level OpenAPI 3+ OAuthFlow object that is backed by a low-level one.
//   - https://spec.openapis.org/oas/v3.1.0#oauth-flow-object
type OAuthFlow struct {
	AuthorizationUrl string                              `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenUrl         string                              `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshUrl       string                              `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           *orderedmap.Map[string, string]     `json:"scopes,renderZero" yaml:"scopes,renderZero"`
	Extensions       *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low              *lowv3.OAuthFlow
}

// NewOAuthFlow creates a new high-level OAuthFlow instance from a low-level one.
func NewOAuthFlow(flow *lowv3.OAuthFlow) *OAuthFlow { _ = "STUB: not implemented"; return nil }

// GoLow returns the low-level OAuthFlow instance used to create the high-level one.
func (o *OAuthFlow) GoLow() *lowv3.OAuthFlow {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level Discriminator instance that was used to create the high-level one, with no type
	return nil
}

func (o *OAuthFlow) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the OAuthFlow object as a byte slice.
	return *new(any)
}

func (o *OAuthFlow) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the OAuthFlow object.
		nil
}

func (o *OAuthFlow) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
