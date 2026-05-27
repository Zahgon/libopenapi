package json

import (
	"go.yaml.in/yaml/v4"
)

// YAMLNodeToJSON converts yaml/json stored in a yaml.Node to json ordered matching the original yaml/json
func YAMLNodeToJSON(node *yaml.Node, indentation string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleYAMLNode(node *yaml.Node) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func handleMappingNode(node *yaml.Node) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// unreachable code in test case, but kept for safety

func handleSequenceNode(node *yaml.Node) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// unreachable code in test case, but kept for safety

func handleScalarNode(node *yaml.Node) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
