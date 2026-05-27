package v3

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/orderedmap"
	"go.yaml.in/yaml/v4"
)

// ServerVariable represents a low-level OpenAPI 3+ ServerVariable object.
//
// ServerVariable is an object representing a Server Variable for server URL template substitution.
// - https://spec.openapis.org/oas/v3.1.0#server-variable-object
//
// This is the only struct that is not Buildable, it's not used by anything other than a Server instance,
// and it has nothing to build that requires it to be buildable.
type ServerVariable struct {
	Enum        []low.NodeReference[string]
	Default     low.NodeReference[string]
	Description low.NodeReference[string]
	Extensions  *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]]
	KeyNode     *yaml.Node
	RootNode    *yaml.Node
	*low.Reference
	low.NodeMap
}

// GetRootNode returns the root yaml node of the ServerVariable object.
func (s *ServerVariable) GetRootNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetKeyNode returns the key yaml node of the ServerVariable object.
	return nil
}

func (s *ServerVariable) GetKeyNode() *yaml.Node {
	_ = "STUB: not implemented"

	// GetExtensions returns all extensions and satisfies the low.HasExtensions interface.
	return nil
}

func (s *ServerVariable) GetExtensions() *orderedmap.Map[low.KeyReference[string], low.ValueReference[*yaml.Node]] {
	_ = "STUB: not implemented"
	return nil

	// Hash will return a consistent Hash of the ServerVariable object
}

func (s *ServerVariable) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

// Pre-allocate and sort enum values
