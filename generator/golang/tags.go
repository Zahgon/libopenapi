// Copyright 2026 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package golang

import (
	"reflect"
)

type fieldTag struct {
	name          string
	skip          bool
	omitempty     bool
	stringEncoded bool
	hasName       bool
	openapi       openAPIMetadata
}

func parseJSONTag(field reflect.StructField) fieldTag {
	_ = "STUB: not implemented"
	return *new(fieldTag)
}

func tagLiteral(name string, required bool, jsonTags, yamlTags, omitEmpty bool, openapiTag string) string {
	_ = "STUB: not implemented"
	return ""
}

func escapeStructTagValue(value string) string { _ = "STUB: not implemented"; return "" }
