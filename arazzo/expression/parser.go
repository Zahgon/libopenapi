// Copyright 2022-2026 Princess Beef Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package expression

// tcharTable is a 128-byte lookup table for RFC 7230 token characters.
// tchar = "!" / "#" / "$" / "%" / "&" / "'" / "*" / "+" / "-" / "." /
//
//	"^" / "_" / "`" / "|" / "~" / DIGIT / ALPHA
var tcharTable [128]bool

func init() {
	for c := 'a'; c <= 'z'; c++ {
		tcharTable[c] = true
	}
	for c := 'A'; c <= 'Z'; c++ {
		tcharTable[c] = true
	}
	for c := '0'; c <= '9'; c++ {
		tcharTable[c] = true
	}
	for _, c := range "!#$%&'*+-.^_`|~" {
		tcharTable[c] = true
	}
}

func isTchar(c byte) bool { _ = "STUB: not implemented"; return false }

// Parse parses a single Arazzo runtime expression. Returns a value type to avoid heap allocation.
func Parse(input string) (Expression, error) {
	_ = "STUB: not implemented"
	return *new(Expression), nil
}

// Fast prefix dispatch on second character

// $url

// $method

// $statusCode, $steps., $sourceDescriptions.

// $request., $response.

// $inputs.

// $outputs.

// $workflows.

// $components.

// parseSource parses $request.{source} or $response.{source} expressions.
func parseSource(input, prefix string, headerType, queryType, pathType, bodyType ExpressionType) (Expression, error) {
	_ = "STUB: not implemented"
	return *new(Expression), nil
}

// Validate tchar for header names

// parseNamedExpression parses expressions like $steps.{name}[.tail], $workflows.{name}[.tail], etc.
func parseNamedExpression(input, prefix string, exprType ExpressionType) (Expression, error) {
	_ = "STUB: not implemented"
	return *new(Expression), nil
}

// Find the first dot to split name from tail

// parseComponents parses $components.{name} and $components.parameters.{name} expressions.
func parseComponents(input string) (Expression, error) {
	_ = "STUB: not implemented"
	return *new(Expression), nil
}

// Special case: $components.parameters.{name}

// General: $components.{name}[.tail]

// ParseEmbedded parses a string that may contain embedded runtime expressions in {$...} blocks.
// Returns alternating literal and expression tokens.
func ParseEmbedded(input string) ([]Token, error) { _ = "STUB: not implemented"; return nil, nil }

// Find the next embedded expression start.

// No more expressions, rest is literal

// Add literal before the brace

// Find closing brace

// Extract and parse the expression (without the surrounding braces).

// Validate checks whether a string is a valid runtime expression without allocating a full AST.
func Validate(input string) error { _ = "STUB: not implemented"; return nil }
