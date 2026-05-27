// Ordered map container
// Works like the Golang `map` built-in, but preserves order that key/value
// pairs were added when iterating.

package orderedmap

import (
	"context"
	"iter"
	"reflect"

	wk8orderedmap "github.com/pb33f/ordered-map/v2"
)

// Pair represents a key/value pair in an ordered map returned for iteration.
type Pair[K comparable, V any] interface {
	Key() K
	KeyPtr() *K
	Value() V
	ValuePtr() *V
	Next() Pair[K, V]
}

// Map represents an ordered map where the key must be a comparable type, the ordering is based on insertion order.
type Map[K comparable, V any] struct {
	*wk8orderedmap.OrderedMap[K, V]
}

type wrapPair[K comparable, V any] struct {
	*wk8orderedmap.Pair[K, V]
}

// New creates an ordered map generic object.
func New[K comparable, V any]() *Map[K, V] { _ = "STUB: not implemented"; return nil }

// GetKeyType returns the reflection type of the key.
func (o *Map[K, V]) GetKeyType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

// GetValueType returns the reflection type of the value.
func (o *Map[K, V]) GetValueType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

// GetOrZero will return the value for the key if it exists, otherwise it will return the zero value for the value type.
func (o *Map[K, V]) GetOrZero(k K) V { _ = "STUB: not implemented"; return *new(V) }

// First returns the first pair in the map useful for iteration.
func (o *Map[K, V]) First() Pair[K, V] { _ = "STUB: not implemented"; return nil }

// FromOldest returns an iterator that yields the oldest key-value pair in the map.
func (o *Map[K, V]) FromOldest() iter.Seq2[K, V] { _ = "STUB: not implemented"; return nil }

// FromNewest returns an iterator that yields the newest key-value pair in the map.
func (o *Map[K, V]) FromNewest() iter.Seq2[K, V] { _ = "STUB: not implemented"; return nil }

// FromNewest returns an iterator that yields the newest key-value pair in the map.
func (o *Map[K, V]) KeysFromOldest() iter.Seq[K] { _ = "STUB: not implemented"; return nil }

// KeysFromNewest returns an iterator that yields the newest key in the map.
func (o *Map[K, V]) KeysFromNewest() iter.Seq[K] { _ = "STUB: not implemented"; return nil }

// ValuesFromOldest returns an iterator that yields the oldest value in the map.
func (o *Map[K, V]) ValuesFromOldest() iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// ValuesFromNewest returns an iterator that yields the newest value in the map.
func (o *Map[K, V]) ValuesFromNewest() iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// From creates a new ordered map from an iterator.
func From[K comparable, V any](iter iter.Seq2[K, V]) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// NewPair instantiates a `Pair` object for use with `FromPairs()`.
func NewPair[K comparable, V any](key K, value V) Pair[K, V] { _ = "STUB: not implemented"; return nil }

// FromPairs creates an `OrderedMap` from an array of pairs.
// Use `NewPair()` to generate input parameters.
func FromPairs[K comparable, V any](pairs ...Pair[K, V]) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// IsZero is required to support `omitempty` tag for YAML/JSON marshaling.
func (o *Map[K, V]) IsZero() bool {
	_ = "STUB: not implemented"

	// Next returns the next pair in the map when iterating.
	return false
}

func (p *wrapPair[K, V]) Next() Pair[K, V] { _ = "STUB: not implemented"; return nil }

// Key returns the key of the pair.
func (p *wrapPair[K, V]) Key() K {
	_ = "STUB: not implemented"

	// KeyPtr returns a pointer to the key of the pair.
	return *new(K)
}

func (p *wrapPair[K, V]) KeyPtr() *K {
	_ = "STUB: not implemented"

	// Value returns the value of the pair.
	return nil
}

func (p *wrapPair[K, V]) Value() V {
	_ = "STUB: not implemented"
	return *

	// ValuePtr returns a pointer to the value of the pair.
	new(V)
}

func (p *wrapPair[K, V]) ValuePtr() *V { _ = "STUB: not implemented"; return nil }

// Len returns the length of a container implementing a `Len()` method.
// Safely returns zero on nil pointer.
func Len[K comparable, V any](m *Map[K, V]) int { _ = "STUB: not implemented"; return 0 }

// Iterate the map in order.
// Safely handles nil pointer.
// Be sure to iterate to end or cancel the context when done to release
// resources.
func Iterate[K comparable, V any](ctx context.Context, m *Map[K, V]) <-chan Pair[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// ToOrderedMap converts a `map` to `OrderedMap`.
func ToOrderedMap[K comparable, V any](m map[K]V) *Map[K, V] { _ = "STUB: not implemented"; return nil }

// First returns map's first pair for iteration.
// Safely handles nil pointer.
func First[K comparable, V any](m *Map[K, V]) Pair[K, V] { _ = "STUB: not implemented"; return nil }

// Cast converts `any` to `Map`.
func Cast[K comparable, V any](v any) *Map[K, V] { _ = "STUB: not implemented"; return nil }

// SortAlpha sorts the map by keys in alphabetical order.
func SortAlpha[K comparable, V any](m *Map[K, V]) *Map[K, V] { _ = "STUB: not implemented"; return nil }
