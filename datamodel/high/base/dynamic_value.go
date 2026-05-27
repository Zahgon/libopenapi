// Copyright 2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

// DynamicValue is used to hold multiple possible types for a schema property. There are two values, a left
// value (A) and a right value (B). The A and B values represent different types that a property can have,
// not necessarily different OpenAPI versions.
//
// For example:
//   - additionalProperties: A = SchemaProxy (when it's a schema), B = bool (when it's a boolean)
//   - items: A = SchemaProxy (when it's a schema), B = bool (when it's a boolean in 3.1)
//   - type: A = string (single type), B = []string (multiple types in 3.1)
//   - exclusiveMinimum: A = bool (in 3.0), B = float64 (in 3.1)
//
// The N value indicates which value is set (0 = A, 1 == B), preventing the need to check both values.
type DynamicValue[A any, B any] struct {
	N         int // 0 == A, 1 == B
	A         A
	B         B
	inline    bool
	renderCtx any // Context for inline rendering (typed as any to avoid import cycles)
}

// IsA will return true if the 'A' or left value is set.
func (d *DynamicValue[A, B]) IsA() bool {
	_ = "STUB: not implemented"

	// IsB will return true if the 'B' or right value is set.
	return false
}

func (d *DynamicValue[A, B]) IsB() bool { _ = "STUB: not implemented"; return false }

func (d *DynamicValue[A, B]) Render() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *DynamicValue[A, B]) RenderInline() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalYAML will create a ready to render YAML representation of the DynamicValue object.
func (d *DynamicValue[A, B]) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	// this is a custom renderer, we can't use the NodeBuilder out of the gate.
	return nil, nil
}

// prefer context-aware method when context is available

// fall back to context-less method

// MarshalYAMLInline will create a ready to render YAML representation of the DynamicValue object. The
// references will be inlined instead of kept as references.
func (d *DynamicValue[A, B]) MarshalYAMLInline() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalYAMLInlineWithContext will create a ready to render YAML representation of the DynamicValue object.
// The references will be inlined and the provided context will be passed through to nested schemas.
// The ctx parameter should be *InlineRenderContext but is typed as any to avoid import cycles.
func (d *DynamicValue[A, B]) MarshalYAMLInlineWithContext(ctx any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
