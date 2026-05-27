package datamodel

import (
	"errors"

	"github.com/pb33f/libopenapi/orderedmap"
)

type (
	ActionFunc[T any]                   func(T) error
	TranslateFunc[IN any, OUT any]      func(IN) (OUT, error)
	TranslateSliceFunc[IN any, OUT any] func(int, IN) (OUT, error)
	TranslateMapFunc[IN any, OUT any]   func(IN) (OUT, error)
	ResultFunc[V any]                   func(V) error
)

type continueError struct {
	error
}

var Continue = &continueError{error: errors.New("Continue")}

type indexedResult[OUT any] struct {
	idx    int
	cont   bool
	output OUT
	err    error
}

type pipelineResult[OUT any] struct {
	seq    int
	cont   bool
	output OUT
	err    error
}

// TranslateSliceParallel iterates a slice in parallel and calls translate()
// asynchronously.
// translate() may return `datamodel.Continue` to continue iteration.
// translate() or result() may return `io.EOF` to break iteration.
// Results are provided sequentially to result() in stable order from slice.
func TranslateSliceParallel[IN any, OUT any](in []IN, translate TranslateSliceFunc[IN, OUT], result ActionFunc[OUT]) error {
	_ = "STUB: not implemented"
	return nil
}

// Buffered to len(in) so workers never block on send.

// Bounded worker pool.

// Enqueue work, then close doneChan after all workers finish.

// Deliver results in stable order using a pending map.

// Flush contiguous completed results starting from nextIdx.

// Check errors first, even when result callback is nil.

// TranslateMapParallel iterates a `*orderedmap.Map` in parallel and calls translate()
// asynchronously.
// translate() or result() may return `io.EOF` to break iteration.
// Safely handles nil pointer.
// Results are provided sequentially to result() in stable order from `*orderedmap.Map`.
func TranslateMapParallel[K comparable, V any, RV any](m *orderedmap.Map[K, V], translate TranslateFunc[orderedmap.Pair[K, V], RV], result ResultFunc[RV]) error {
	_ = "STUB: not implemented"
	return nil
}

// Snapshot pairs for indexed access.

// Bounded worker pool.

// Enqueue work, then close doneChan after all workers finish.

// Deliver results in stable order.

type pipelineWork[IN any] struct {
	seq   int
	input IN
}

// TranslatePipeline processes input sequentially through predicate(), sends to
// translate() in parallel, then outputs in stable order.
// translate() may return `datamodel.Continue` to continue iteration.
// Caller must close `in` channel to indicate EOF.
// TranslatePipeline closes `out` channel to indicate EOF.
func TranslatePipeline[IN any, OUT any](in <-chan IN, out chan<- OUT, translate TranslateFunc[IN, OUT]) error {
	_ = "STUB: not implemented"
	return nil
}

// Launch worker pool.

// Send the error result so the collector can detect it

// Iterate input, assign sequence numbers, send to workers.

// Close resultChan after all workers and the enqueue goroutine finish.

// Collect results in stable order, send to output channel.

// Error already stored in reterr by the worker
