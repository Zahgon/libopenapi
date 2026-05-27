// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	low "github.com/pb33f/libopenapi/datamodel/low/v2"
	"github.com/pb33f/libopenapi/orderedmap"
)

// ParameterDefinitions is a high-level representation of a Swagger / OpenAPI 2 Parameters Definitions object
// that is backed by a low-level one.
//
// ParameterDefinitions holds parameters to be reused across operations. Parameter definitions can be
// referenced to the ones defined here. It does not define global operation parameters
//   - https://swagger.io/specification/v2/#parametersDefinitionsObject
type ParameterDefinitions struct {
	Definitions *orderedmap.Map[string, *Parameter]
	low         *low.ParameterDefinitions
}

// NewParametersDefinitions creates a new instance of a high-level ParameterDefinitions, from a low-level one.
// Every parameter is extracted asynchronously due to the potential depth
func NewParametersDefinitions(parametersDefinitions *low.ParameterDefinitions) *ParameterDefinitions {
	_ = "STUB: not implemented"
	return nil
}

// GoLow returns the low-level ParameterDefinitions instance that backs the low-level one.
func (p *ParameterDefinitions) GoLow() *low.ParameterDefinitions {
	_ = "STUB: not implemented"
	return nil
}
