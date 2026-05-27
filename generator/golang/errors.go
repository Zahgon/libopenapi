// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

import (
	"errors"
)

var (
	ErrNilSchema          = errors.New("nil schema")
	ErrNilType            = errors.New("nil type")
	ErrUnsupportedType    = errors.New("unsupported type")
	ErrUnsupportedMapKey  = errors.New("unsupported map key")
	ErrInvalidPackageName = errors.New("invalid package name")
)

func wrapPath(err error, path string) error { _ = "STUB: not implemented"; return nil }
