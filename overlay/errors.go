// Copyright 2022-2025 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package overlay

import (
	"errors"

	highoverlay "github.com/pb33f/libopenapi/datamodel/high/overlay"
)

// Warning represents a non-fatal issue encountered during overlay application.
type Warning struct {
	Action  *highoverlay.Action
	Target  string
	Message string
}

func (w *Warning) String() string { _ = "STUB: not implemented"; return "" }

// OverlayError represents an error that occurred during an overlay application.
type OverlayError struct {
	Action *highoverlay.Action
	Cause  error
}

func (e *OverlayError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *OverlayError) Unwrap() error {
	_ = "STUB: not implemented"

	// Sentinel errors for overlay operations.
	return nil
}

var (
	// Parsing errors
	ErrInvalidOverlay      = errors.New("invalid overlay document")
	ErrMissingOverlayField = errors.New("missing required 'overlay' field")
	ErrMissingInfo         = errors.New("missing required 'info' field")
	ErrMissingActions      = errors.New("missing required 'actions' field")
	ErrEmptyActions        = errors.New("actions array must contain at least one action")

	// JSONPath errors
	ErrInvalidJSONPath = errors.New("invalid JSONPath expression")
	ErrPrimitiveTarget = errors.New("JSONPath target resolved to primitive/null; must be object or array")

	// Application errors
	ErrNoTargetDocument = errors.New("no target document provided")

	// Copy action errors
	ErrCopySourceNotFound = errors.New("copy source JSONPath matched zero nodes")
	ErrCopySourceMultiple = errors.New("copy source JSONPath must match exactly one node")
	ErrCopyTypeMismatch   = errors.New("copy source and target must be the same type")
)
