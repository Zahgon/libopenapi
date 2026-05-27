// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	"errors"
)

// Document errors
var (
	ErrInvalidArazzo             = errors.New("invalid arazzo document")
	ErrMissingArazzoField        = errors.New("missing required 'arazzo' field")
	ErrMissingInfo               = errors.New("missing required 'info' field")
	ErrMissingSourceDescriptions = errors.New("missing required 'sourceDescriptions' field")
	ErrEmptySourceDescriptions   = errors.New("sourceDescriptions must have at least one entry")
	ErrMissingWorkflows          = errors.New("missing required 'workflows' field")
	ErrEmptyWorkflows            = errors.New("workflows must have at least one entry")
)

// Workflow errors
var (
	ErrMissingWorkflowId   = errors.New("missing required 'workflowId'")
	ErrMissingSteps        = errors.New("missing required 'steps'")
	ErrEmptySteps          = errors.New("steps must have at least one entry")
	ErrDuplicateWorkflowId = errors.New("duplicate workflowId")
)

// Step errors
var (
	ErrMissingStepId         = errors.New("missing required 'stepId'")
	ErrDuplicateStepId       = errors.New("duplicate stepId within workflow")
	ErrStepMutualExclusion   = errors.New("step must have exactly one of operationId, operationPath, or workflowId")
	ErrExecutorNotConfigured = errors.New("executor is not configured")
)

// Parameter errors
var (
	ErrMissingParameterName  = errors.New("missing required 'name'")
	ErrMissingParameterIn    = errors.New("missing required 'in' for operation parameter")
	ErrInvalidParameterIn    = errors.New("'in' must be path, query, header, or cookie")
	ErrMissingParameterValue = errors.New("missing required 'value'")
)

// Action errors
var (
	ErrMissingActionName     = errors.New("missing required 'name'")
	ErrMissingActionType     = errors.New("missing required 'type'")
	ErrInvalidSuccessType    = errors.New("success action type must be 'end' or 'goto'")
	ErrInvalidFailureType    = errors.New("failure action type must be 'end', 'retry', or 'goto'")
	ErrActionMutualExclusion = errors.New("action cannot have both workflowId and stepId")
	ErrGotoRequiresTarget    = errors.New("goto action requires workflowId or stepId")
	ErrStepIdNotInWorkflow   = errors.New("stepId must reference a step in the current workflow")
)

// Criterion errors
var (
	ErrMissingCondition = errors.New("missing required 'condition'")
)

// Expression errors
var (
	ErrInvalidExpression       = errors.New("invalid runtime expression")
	ErrUnknownExpressionPrefix = errors.New("unknown expression prefix")
)

// Reference errors
var (
	ErrUnresolvedWorkflowRef  = errors.New("workflowId references unknown workflow")
	ErrUnresolvedSourceDesc   = errors.New("sourceDescription reference not found")
	ErrUnresolvedOperationRef = errors.New("operation reference not found")
	ErrOperationSourceMapping = errors.New("operation source mapping failed")
	ErrUnresolvedComponent    = errors.New("component reference not found")
	ErrCircularDependency     = errors.New("circular workflow dependency detected")
)

// Source description errors
var (
	ErrSourceDescLoadFailed = errors.New("failed to load source description")
)

// ValidationError represents a structured validation error with source location.
type ValidationError struct {
	Path   string // e.g. "workflows[0].steps[2].parameters[1]"
	Line   int
	Column int
	Cause  error
}

func (e *ValidationError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ValidationError) Unwrap() error {
	_ = "STUB: not implemented"

	// StepFailureError represents a step execution failure with structured context.
	return nil
}

type StepFailureError struct {
	StepId         string
	CriterionIndex int // -1 if not criterion-related
	Message        string
	Cause          error
}

func (e *StepFailureError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *StepFailureError) Unwrap() error {
	_ = "STUB: not implemented"

	// Warning represents a non-fatal validation issue.
	return nil
}

type Warning struct {
	Path    string
	Line    int
	Column  int
	Message string
}

func (w *Warning) String() string { _ = "STUB: not implemented"; return "" }

// ValidationResult holds all validation errors and warnings.
type ValidationResult struct {
	Errors   []*ValidationError
	Warnings []*Warning
}

// HasErrors returns true if there are any validation errors.
func (r *ValidationResult) HasErrors() bool { _ = "STUB: not implemented"; return false }

// HasWarnings returns true if there are any validation warnings.
func (r *ValidationResult) HasWarnings() bool { _ = "STUB: not implemented"; return false }

// Error implements the error interface, returning all errors as a combined string.
func (r *ValidationResult) Error() string { _ = "STUB: not implemented"; return "" }

// Unwrap returns the individual validation errors for use with errors.Is/As (Go 1.20+).
func (r *ValidationResult) Unwrap() []error { _ = "STUB: not implemented"; return nil }
