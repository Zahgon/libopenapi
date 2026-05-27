// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package arazzo

import (
	"context"
	"time"

	"github.com/pb33f/libopenapi/arazzo/expression"
	high "github.com/pb33f/libopenapi/datamodel/high/arazzo"
	"github.com/pb33f/libopenapi/orderedmap"
)

// actionTypeRequest groups the parameters for processActionTypeResult,
// normalizing both success and failure actions into a common structure.
type actionTypeRequest struct {
	actionType     string
	workflowId     string
	stepId         string
	retryAfterSec  float64
	retryLimit     int64
	currentRetries int
}

type stepActionResult struct {
	endWorkflow   bool
	retryCurrent  bool
	retryAfter    time.Duration
	jumpToStepIdx int
}

func (e *Engine) processSuccessActions(
	ctx context.Context,
	step *high.Step,
	wf *high.Workflow,
	exprCtx *expression.Context,
	state *executionState,
	stepIndexByID map[string]int,
) (*stepActionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) processFailureActions(
	ctx context.Context,
	step *high.Step,
	wf *high.Workflow,
	exprCtx *expression.Context,
	state *executionState,
	stepIndexByID map[string]int,
	currentRetries int,
) (*stepActionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) processActionTypeResult(
	ctx context.Context,
	req *actionTypeRequest,
	exprCtx *expression.Context,
	state *executionState,
	stepIndexByID map[string]int,
) (*stepActionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) selectSuccessAction(stepActions, workflowActions []*high.SuccessAction, exprCtx *expression.Context) (*high.SuccessAction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) selectFailureAction(stepActions, workflowActions []*high.FailureAction, exprCtx *expression.Context) (*high.FailureAction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) findMatchingSuccessAction(actions []*high.SuccessAction, exprCtx *expression.Context) (*high.SuccessAction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) findMatchingFailureAction(actions []*high.FailureAction, exprCtx *expression.Context) (*high.FailureAction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// findMatchingAction iterates actions, resolves component references, evaluates criteria,
// and returns the first action whose criteria all pass.
func findMatchingAction[T any](
	actions []T,
	resolve func(T) (T, error),
	getCriteria func(T) []*high.Criterion,
	evalCriteria func([]*high.Criterion, *expression.Context) (bool, error),
	exprCtx *expression.Context,
) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (e *Engine) resolveSuccessAction(action *high.SuccessAction) (*high.SuccessAction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) resolveFailureAction(action *high.FailureAction) (*high.FailureAction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// lookupComponent resolves a $components reference against an ordered map.
func lookupComponent[T any](ref, prefix string, componentMap *orderedmap.Map[string, T]) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// evaluateActionCriteria evaluates all criteria for an action, using per-engine caches.
func (e *Engine) evaluateActionCriteria(criteria []*high.Criterion, exprCtx *expression.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func workflowFailureError(workflowID string, wfResult *WorkflowResult) error {
	_ = "STUB: not implemented"
	return nil
}
