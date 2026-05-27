// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package expression

import (
	"go.yaml.in/yaml/v4"
)

// Context holds runtime values for expression evaluation.
type Context struct {
	URL             string
	Method          string
	StatusCode      int
	RequestHeaders  map[string]string
	RequestQuery    map[string]string
	RequestPath     map[string]string
	RequestBody     *yaml.Node
	ResponseHeaders map[string]string
	ResponseBody    *yaml.Node
	Inputs          map[string]any
	Outputs         map[string]any
	Steps           map[string]*StepContext
	Workflows       map[string]*WorkflowContext
	SourceDescs     map[string]*SourceDescContext
	Components      *ComponentsContext
}

// StepContext holds inputs and outputs for a specific step.
type StepContext struct {
	Inputs  map[string]any
	Outputs map[string]any
}

// WorkflowContext holds inputs and outputs for a specific workflow.
type WorkflowContext struct {
	Inputs  map[string]any
	Outputs map[string]any
}

// SourceDescContext holds resolved source description data.
type SourceDescContext struct {
	URL string
}

// ComponentsContext holds resolved component data.
type ComponentsContext struct {
	Parameters     map[string]any
	SuccessActions map[string]any
	FailureActions map[string]any
	Inputs         map[string]any
}

// Evaluate resolves a parsed Expression against a Context.
func Evaluate(expr Expression, ctx *Context) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// EvaluateString parses and evaluates a runtime expression string in one call.
func EvaluateString(input string, ctx *Context) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func resolveSteps(expr Expression, ctx *Context) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func splitTail(tail string) (segment, rest string) { _ = "STUB: not implemented"; return "", "" }

func resolveStepTail(tail string, sc *StepContext, stepName string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func resolveWorkflows(expr Expression, ctx *Context) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func resolveSourceDescriptions(expr Expression, ctx *Context) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func resolveComponents(expr Expression, ctx *Context) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// resolveJSONPointer navigates a yaml.Node tree using a JSON Pointer (RFC 6901).
// The pointer should start with "/" (the leading "#" has already been stripped).
func resolveJSONPointer(node *yaml.Node, pointer string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Unwrap document nodes

// Find next segment boundary

// Unescape JSON Pointer: ~1 -> /, ~0 -> ~

// UnescapeJSONPointer applies RFC 6901 unescaping: ~1 -> /, ~0 -> ~
func UnescapeJSONPointer(s string) string { _ = "STUB: not implemented"; return "" }

// yamlNodeToValue converts a yaml.Node to a Go native value.
func yamlNodeToValue(node *yaml.Node) any { _ = "STUB: not implemented"; return *new(any) }

// resolveDeepValue traverses into a resolved component value using a dot-separated path.
func resolveDeepValue(v any, path, componentType, componentName string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
