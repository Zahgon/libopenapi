// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	"context"
	"time"

	"github.com/pb33f/libopenapi/arazzo/expression"
	high "github.com/pb33f/libopenapi/datamodel/high/arazzo"
	"go.yaml.in/yaml/v4"
)

func (e *Engine) executeStep(ctx context.Context, step *high.Step, wf *high.Workflow, exprCtx *expression.Context, state *executionState) *StepResult {
	_ = "STUB: not implemented"
	// retained for future per-workflow step configuration
	return nil
}

func (e *Engine) evaluateStepSuccessCriteria(step *high.Step, exprCtx *expression.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) buildExecutionRequest(step *high.Step, exprCtx *expression.Context) (*ExecutionRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) resolveStepSource(step *high.Step) *ResolvedSource {
	_ = "STUB: not implemented"
	return nil
}

// Deterministic fallback: use the first source from the document's ordered list.

func (e *Engine) resolveParameter(param *high.Parameter) (*high.Parameter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) resolveYAMLNodeValue(node *yaml.Node, exprCtx *expression.Context) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (e *Engine) resolveExpressionValues(value any, exprCtx *expression.Context) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (e *Engine) applyPayloadReplacements(payload any, replacements []*high.PayloadReplacement, exprCtx *expression.Context, stepId string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func setJSONPointerValue(root map[string]any, pointer string, value any) error {
	_ = "STUB: not implemented"
	return nil
}

func valueNeedsResolution(v any) bool { _ = "STUB: not implemented"; return false }

func sliceNeedsResolution(items []any) bool { _ = "STUB: not implemented"; return false }

func mapAnyNeedsResolution(items map[any]any) bool { _ = "STUB: not implemented"; return false }

func mapNeedsResolution(items map[string]any) bool { _ = "STUB: not implemented"; return false }

func (e *Engine) evaluateStringValue(input string, exprCtx *expression.Context) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (e *Engine) populateStepOutputs(step *high.Step, result *StepResult, exprCtx *expression.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) populateWorkflowOutputs(wf *high.Workflow, result *WorkflowResult, exprCtx *expression.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func firstHeaderValues(headers map[string][]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func sleepWithContext(ctx context.Context, d time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
