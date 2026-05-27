package orderedmap

import (
	"github.com/pb33f/libopenapi/datamodel/high/nodes"
	"go.yaml.in/yaml/v4"
)

type marshaler interface {
	MarshalYAML() (interface{}, error)
}

type NodeBuilder interface {
	AddYAMLNode(parent *yaml.Node, entry *nodes.NodeEntry) *yaml.Node
}

type MapToYamlNoder interface {
	ToYamlNode(n NodeBuilder, l any) *yaml.Node
}

type hasValueNode interface {
	GetValueNode() *yaml.Node
}

type hasValueUntyped interface {
	GetValueUntyped() any
}

type findValueUntyped interface {
	FindValueUntyped(k string) any
}

// ToYamlNode converts the ordered map to a yaml node ready for marshalling.
func (o *Map[K, V]) ToYamlNode(n NodeBuilder, l any) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

// TODO marshal inline?

func findKeyNode(key string, m *yaml.Node) *yaml.Node { _ = "STUB: not implemented"; return nil }

// FindValueUntyped finds a value in the ordered map by key if the stored value for that key implements GetValueUntyped otherwise just returns the value.
func (o *Map[K, V]) FindValueUntyped(key string) any { _ = "STUB: not implemented"; return *new(any) }
