// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
)

// OAuthFlowsChanges represents changes found between two OpenAPI OAuthFlows objects.
type OAuthFlowsChanges struct {
	*PropertyChanges
	ImplicitChanges          *OAuthFlowChanges `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	PasswordChanges          *OAuthFlowChanges `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentialsChanges *OAuthFlowChanges `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCodeChanges *OAuthFlowChanges `json:"authCode,omitempty" yaml:"authCode,omitempty"`
	DeviceChanges            *OAuthFlowChanges `json:"device,omitempty" yaml:"device,omitempty"` // OpenAPI 3.2+ device flow
	ExtensionChanges         *ExtensionChanges `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between OAuthFlows objects
func (o *OAuthFlowsChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the number of changes made between two OAuthFlows instances.
func (o *OAuthFlowsChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the number of breaking changes made between two OAuthFlows objects.
func (o *OAuthFlowsChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareOAuthFlows compares a left and right OAuthFlows object. If changes are found a pointer to *OAuthFlowsChanges
// is returned, otherwise nil is returned.
func CompareOAuthFlows(l, r *v3.OAuthFlows) *OAuthFlowsChanges {
	_ = "STUB: not implemented"
	return nil
}

// client credentials

// implicit

// password

// auth code

// device flow (OpenAPI 3.2+)

// OAuthFlowChanges represents an OpenAPI OAuthFlow object.
type OAuthFlowChanges struct {
	*PropertyChanges
	ExtensionChanges *ExtensionChanges `json:"extensions,omitempty" yaml:"extensions,omitempty"`
}

// GetAllChanges returns a slice of all changes made between OAuthFlow objects
func (o *OAuthFlowChanges) GetAllChanges() []*Change { _ = "STUB: not implemented"; return nil }

// TotalChanges returns the total number of changes made between two OAuthFlow objects
func (o *OAuthFlowChanges) TotalChanges() int { _ = "STUB: not implemented"; return 0 }

// TotalBreakingChanges returns the total number of breaking changes made between two OAuthFlow objects
func (o *OAuthFlowChanges) TotalBreakingChanges() int { _ = "STUB: not implemented"; return 0 }

// CompareOAuthFlow checks a left and a right OAuthFlow object for changes. If found, returns a pointer to
// an OAuthFlowChanges instance, or nil if nothing is found.
func CompareOAuthFlow(l, r *v3.OAuthFlow) *OAuthFlowChanges { _ = "STUB: not implemented"; return nil }
