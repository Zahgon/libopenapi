// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	low "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// buildLowSecurityScheme builds a low-level SecurityScheme from a resolved YAML node.
func buildLowSecurityScheme(node *yaml.Node, idx *index.SpecIndex) (*low.SecurityScheme, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SecurityScheme represents a high-level OpenAPI 3+ SecurityScheme object that is backed by a low-level one.
//
// Defines a security scheme that can be used by the operations.
//
// Supported schemes are HTTP authentication, an API key (either as a header, a cookie parameter or as a query parameter),
// mutual TLS (use of a client certificate), OAuth2's common flows (implicit, password, client credentials and
// authorization code) as defined in RFC6749 (https://www.rfc-editor.org/rfc/rfc6749), and OpenID Connect Discovery.
// Please note that as of 2020, the implicit  flow is about to be deprecated by OAuth 2.0 Security Best Current Practice.
// Recommended for most use case is Authorization Code Grant flow with PKCE.
//   - https://spec.openapis.org/oas/v3.1.0#security-scheme-object
type SecurityScheme struct {
	Reference         string                              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Type              string                              `json:"type,omitempty" yaml:"type,omitempty"`
	Description       string                              `json:"description,omitempty" yaml:"description,omitempty"`
	Name              string                              `json:"name,omitempty" yaml:"name,omitempty"`
	In                string                              `json:"in,omitempty" yaml:"in,omitempty"`
	Scheme            string                              `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	BearerFormat      string                              `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows             *OAuthFlows                         `json:"flows,omitempty" yaml:"flows,omitempty"`
	OpenIdConnectUrl  string                              `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
	OAuth2MetadataUrl string                              `json:"oauth2MetadataUrl,omitempty" yaml:"oauth2MetadataUrl,omitempty"` // OpenAPI 3.2+ OAuth2 metadata URL
	Deprecated        bool                                `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`               // OpenAPI 3.2+ deprecated flag
	Extensions        *orderedmap.Map[string, *yaml.Node] `json:"-" yaml:"-"`
	low               *low.SecurityScheme
}

// NewSecurityScheme creates a new high-level SecurityScheme from a low-level one.
func NewSecurityScheme(ss *low.SecurityScheme) *SecurityScheme {
	_ = "STUB: not implemented"
	return nil
}

// GoLow returns the low-level SecurityScheme that was used to create the high-level one.
func (s *SecurityScheme) GoLow() *low.SecurityScheme {
	_ = "STUB: not implemented"

	// GoLowUntyped will return the low-level SecurityScheme instance that was used to create the high-level one, with no type
	return nil
}

func (s *SecurityScheme) GoLowUntyped() any {
	_ = "STUB: not implemented"

	// IsReference returns true if this SecurityScheme is a reference to another SecurityScheme definition.
	return *new(any)
}

func (s *SecurityScheme) IsReference() bool { _ = "STUB: not implemented"; return false }

// GetReference returns the reference string if this is a reference SecurityScheme.
func (s *SecurityScheme) GetReference() string {
	_ = "STUB: not implemented"

	// Render will return a YAML representation of the SecurityScheme object as a byte slice.
	return ""
}

func (s *SecurityScheme) Render() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalYAML will create a ready to render YAML representation of the SecurityScheme object.
		nil
}

func (s *SecurityScheme) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	// Handle reference-only security scheme
	return nil, nil
}

// MarshalYAMLInline will create a ready to render YAML representation of the SecurityScheme object,
// with all references resolved inline.
func (s *SecurityScheme) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	// reference-only objects render as $ref nodes
	return nil, nil
}

// resolve external reference if present

// buildLowSecurityScheme never returns an error, so we can ignore it

// MarshalYAMLInlineWithContext will create a ready to render YAML representation of the SecurityScheme object,
// resolving any references inline where possible. Uses the provided context for cycle detection.
// The ctx parameter should be *base.InlineRenderContext but is typed as any to satisfy the
// high.RenderableInlineWithContext interface without import cycles.
func (s *SecurityScheme) MarshalYAMLInlineWithContext(ctx any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolve external reference if present

// buildLowSecurityScheme never returns an error, so we can ignore it

// CreateSecuritySchemeRef creates a SecurityScheme that renders as a $ref to another security scheme definition.
// This is useful when building OpenAPI specs programmatically and you want to reference
// a security scheme defined in components/securitySchemes rather than inlining the full definition.
//
// Example:
//
//	ss := v3.CreateSecuritySchemeRef("#/components/securitySchemes/BearerAuth")
//
// Renders as:
//
//	$ref: '#/components/securitySchemes/BearerAuth'
func CreateSecuritySchemeRef(ref string) *SecurityScheme { _ = "STUB: not implemented"; return nil }
