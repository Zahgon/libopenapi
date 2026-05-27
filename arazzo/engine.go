// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	"context"

	"github.com/pb33f/libopenapi/arazzo/expression"
	high "github.com/pb33f/libopenapi/datamodel/high/arazzo"
)

const maxWorkflowDepth = 32
const maxStepTransitions = 1024

// Executor defines the interface for executing API calls.
type Executor interface {
	Execute(ctx context.Context, req *ExecutionRequest) (*ExecutionResponse, error)
}

// ExecutionRequest represents a request to execute an API operation.
type ExecutionRequest struct {
	Source        *ResolvedSource
	OperationID   string
	OperationPath string
	Method        string
	Parameters    map[string]any
	RequestBody   any
	ContentType   string
}

// ExecutionResponse represents the response from an API operation execution.
type ExecutionResponse struct {
	StatusCode int
	Headers    map[string][]string
	Body       any
	URL        string // Actual request URL (populated by Executor)
	Method     string // HTTP method used (populated by Executor)
}

// EngineConfig configures engine behavior.
type EngineConfig struct {
	RetainResponseBodies bool // If false, nil out response bodies after extracting outputs
}

// Engine orchestrates the execution of Arazzo workflows.
// An Engine is NOT safe for concurrent use from multiple goroutines.
type Engine struct {
	document         *high.Arazzo
	executor         Executor
	sources          map[string]*ResolvedSource
	defaultSource    *ResolvedSource // cached for single-source fast path
	sourceOrder      []string        // deterministic source ordering from document
	workflows        map[string]*high.Workflow
	config           *EngineConfig
	exprCache        map[string]expression.Expression
	criterionCaches  *criterionCaches
	cachedComponents *expression.ComponentsContext // immutable component maps, built once
}

// NewEngine creates a new Engine for executing Arazzo workflows.
func NewEngine(doc *high.Arazzo, executor Executor, sources []*ResolvedSource) *Engine {
	_ = "STUB: not implemented"
	return nil
}

// Cache a default source for the single-source fast path to avoid map iteration per step.

// Build deterministic source ordering from the document's ordered SourceDescriptions list.

// NewEngineWithConfig creates a new Engine with custom configuration.
func NewEngineWithConfig(doc *high.Arazzo, executor Executor, sources []*ResolvedSource, config *EngineConfig) *Engine {
	_ = "STUB: not implemented"
	return nil
}

// ClearCaches resets all per-engine caches (expressions, regex, JSONPath).
func (e *Engine) ClearCaches() { _ = "STUB: not implemented"; return }

// RunWorkflow executes a single workflow by its ID.
func (e *Engine) RunWorkflow(ctx context.Context, workflowId string, inputs map[string]any) (*WorkflowResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunAll executes all workflows in dependency order.
func (e *Engine) RunAll(ctx context.Context, inputs map[string]map[string]any) (*RunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Topological sort on dependsOn

type executionState struct {
	workflowResults  map[string]*WorkflowResult
	workflowContexts map[string]*expression.WorkflowContext
	activeWorkflows  map[string]struct{}
	depth            int
}

func (e *Engine) runWorkflow(ctx context.Context, workflowId string, inputs map[string]any, state *executionState) (*WorkflowResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Error is non-fatal: unresolvable component input expressions fall back to raw YAML nodes.

func (e *Engine) topologicalSort() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func dependencyExecutionError(wf *high.Workflow, workflowResults map[string]*WorkflowResult) error {
	_ = "STUB: not implemented"
	return nil
}

func workflowExecutionFailureResult(workflowID string, inputs map[string]any, execErr error) *WorkflowResult {
	_ = "STUB: not implemented"
	return nil
}

func stepFailureOrDefault(stepID string, stepErr error) error {
	_ = "STUB: not implemented"
	return nil
}

// parseExpression parses and caches an expression.
func (e *Engine) parseExpression(input string) (expression.Expression, error) {
	_ = "STUB: not implemented"
	return *new(expression.Expression), nil
}

// buildCachedComponents builds the immutable portion of the components context once.
// Parameters, SuccessActions, and FailureActions are read-only and shared across workflow runs.
// Inputs are resolved per-run because they may contain runtime expressions.
func (e *Engine) buildCachedComponents() *expression.ComponentsContext {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) newExpressionContext(inputs map[string]any, state *executionState) (*expression.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func copyWorkflowContexts(src map[string]*expression.WorkflowContext) map[string]*expression.WorkflowContext {
	_ = "STUB: not implemented"
	return nil
}
