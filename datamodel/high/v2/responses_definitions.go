// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	low "github.com/pb33f/libopenapi/datamodel/low/v2"
	"github.com/pb33f/libopenapi/orderedmap"
)

// ResponsesDefinitions is a high-level representation of a Swagger / OpenAPI 2 Responses Definitions object.
// that is backed by a low-level one.
//
// ResponsesDefinitions is an object to hold responses to be reused across operations. Response definitions can be
// referenced to the ones defined here. It does not define global operation responses
//   - https://swagger.io/specification/v2/#responsesDefinitionsObject
type ResponsesDefinitions struct {
	Definitions *orderedmap.Map[string, *Response]
	low         *low.ResponsesDefinitions
}

// NewResponsesDefinitions will create a new high-level instance of ResponsesDefinitions from a low-level one.
func NewResponsesDefinitions(responsesDefinitions *low.ResponsesDefinitions) *ResponsesDefinitions {
	_ = "STUB: not implemented"
	return nil
}

// GoLow returns the low-level ResponsesDefinitions used to create the high-level one.
func (r *ResponsesDefinitions) GoLow() *low.ResponsesDefinitions {
	_ = "STUB: not implemented"
	return nil
}
