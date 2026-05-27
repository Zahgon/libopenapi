// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	low "github.com/pb33f/libopenapi/datamodel/low/v2"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// SecurityScheme is a high-level representation of a Swagger / OpenAPI 2 SecurityScheme object
// backed by a low-level one.
//
// SecurityScheme allows the definition of a security scheme that can be used by the operations. Supported schemes are
// basic authentication, an API key (either as a header or as a query parameter) and OAuth2's common flows
// (implicit, password, application and access code)
//   - https://swagger.io/specification/v2/#securityDefinitionsObject
type SecurityScheme struct {
	Type             string
	Description      string
	Name             string
	In               string
	Flow             string
	AuthorizationUrl string
	TokenUrl         string
	Scopes           *Scopes
	Extensions       *orderedmap.Map[string, *yaml.Node]
	low              *low.SecurityScheme
}

// NewSecurityScheme creates a new instance of SecurityScheme from a low-level one.
func NewSecurityScheme(securityScheme *low.SecurityScheme) *SecurityScheme {
	_ = "STUB: not implemented"
	return nil
}

// GoLow returns the low-level SecurityScheme that was used to create the high-level one.
func (s *SecurityScheme) GoLow() *low.SecurityScheme { _ = "STUB: not implemented"; return nil }
