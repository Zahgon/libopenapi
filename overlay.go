// Copyright 2022-2025 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package libopenapi

import (
	"github.com/pb33f/libopenapi/datamodel"
	highoverlay "github.com/pb33f/libopenapi/datamodel/high/overlay"
	"github.com/pb33f/libopenapi/overlay"
)

// OverlayResult contains the result of applying an overlay to a target document.
type OverlayResult struct {
	// Bytes raw YAML bytes of the modified document after the overlay has been applied.
	Bytes []byte

	// OverlayDocument is the modified document, ready to have a model built from it.
	// The document is created using the same configuration as the input document
	// (for ApplyOverlay and ApplyOverlayFromBytes), or with a default configuration
	// (for ApplyOverlayToSpecBytes and ApplyOverlayFromBytesToSpecBytes).
	OverlayDocument Document

	// Warnings that occurred during overlay application.
	Warnings []*overlay.Warning
}

// NewOverlayDocument creates a new overlay document from the provided bytes.
// The overlay document can then be applied to a target OpenAPI document using ApplyOverlay.
func NewOverlayDocument(overlayBytes []byte) (*highoverlay.Overlay, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApplyOverlay applies the overlay to the target document and returns the modified document.
// This is the primary entry point for an overlay application when working with Document objects.
//
// The returned OverlayDocument uses the same configuration as the input document.
func ApplyOverlay(document Document, ov *highoverlay.Overlay) (*OverlayResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApplyOverlayFromBytes applies an overlay (provided as bytes) to the target document.
// This is a convenience function when you have a Document but the overlay as raw bytes.
//
// The returned OverlayDocument uses the same configuration as the input document.
func ApplyOverlayFromBytes(document Document, overlayBytes []byte) (*OverlayResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApplyOverlayToSpecBytes applies the overlay to the target document bytes.
// Use this when you have raw spec bytes and a parsed Overlay object.
//
// The returned OverlayDocument uses a default document configuration.
func ApplyOverlayToSpecBytes(docBytes []byte, ov *highoverlay.Overlay) (*OverlayResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApplyOverlayFromBytesToSpecBytes applies an overlay to target document bytes,
// where both the overlay and target document are provided as raw bytes.
// This is the most convenient function when you don't need to configure either document.
//
// The returned OverlayDocument uses a default document configuration.
func ApplyOverlayFromBytesToSpecBytes(docBytes, overlayBytes []byte) (*OverlayResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// applyOverlayToBytesWithConfig is the internal function that applies the overlay to bytes
// and creates a Document with the specified configuration (nil for default).
func applyOverlayToBytesWithConfig(targetBytes []byte, ov *highoverlay.Overlay, config *datamodel.DocumentConfiguration) (*OverlayResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
