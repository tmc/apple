//go:build darwin

package ane

// StateHandle manages KV cache state for stateful MIL models.
//
// It wraps a Kernel compiled from MIL programs containing read_state
// and coreml_update_state ops, tracking the current sequence position
// for incremental decode.
type StateHandle struct {
	kernel   *Kernel
	maxSeq   int
	position int
}

// NewStateHandle creates a StateHandle wrapping a stateful Kernel.
// maxSeq is the maximum sequence length the KV cache supports.
func NewStateHandle(k *Kernel, maxSeq int) *StateHandle {
	return &StateHandle{
		kernel: k,
		maxSeq: maxSeq,
	}
}

// Kernel returns the underlying Kernel.
func (s *StateHandle) Kernel() *Kernel { return s.kernel }

// Position returns the current sequence position in the KV cache.
func (s *StateHandle) Position() int { return s.position }

// MaxSeq returns the maximum sequence length.
func (s *StateHandle) MaxSeq() int { return s.maxSeq }

// Reset zeros the sequence position. The actual state buffer clearing
// happens on the next evaluation cycle through the ANE framework.
func (s *StateHandle) Reset() {
	s.position = 0
}

// Advance increments the sequence position by n tokens.
func (s *StateHandle) Advance(n int) {
	s.position += n
	if s.position > s.maxSeq {
		s.position = s.maxSeq
	}
}

// Remaining returns how many tokens can still be appended.
func (s *StateHandle) Remaining() int {
	return s.maxSeq - s.position
}

// Close releases the underlying kernel.
func (s *StateHandle) Close() error {
	return s.kernel.Close()
}
