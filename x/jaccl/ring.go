package jaccl

import "fmt"

// ringGatherBlocks circulates rank-sized blocks clockwise. stage transfers one
// byte stripe from this rank to its right neighbor and receives the matching
// stripe from its left neighbor. It is separate from the native transport so
// its rank-major reconstruction can be tested without RDMA hardware.
func ringGatherBlocks(rank, size int, dst, src []byte, wires int, stage func(step, wire int, send, recv []byte) error) error {
	if size < 1 || rank < 0 || rank >= size || wires < 1 || len(dst) != size*len(src) || stage == nil {
		return fmt.Errorf("ring gather shape: %w", ErrProtocol)
	}
	if size == 1 || len(src) == 0 {
		copy(dst, src)
		return nil
	}
	copy(dst[rank*len(src):], src)
	forward := src
	for step := 0; step < size-1; step++ {
		block := (rank - step - 1 + size) % size
		received := dst[block*len(src) : (block+1)*len(src)]
		for wire := 0; wire < wires; wire++ {
			start, end := ringWireRange(len(src), wire, wires)
			if start == end {
				continue
			}
			if err := stage(step, wire, forward[start:end], received[start:end]); err != nil {
				return fmt.Errorf("ring step %d wire %d: %w", step, wire, err)
			}
		}
		forward = received
	}
	return nil
}

func ringWireRange(length, wire, wires int) (int, int) {
	start := length * wire / wires
	end := length * (wire + 1) / wires
	return start, end
}
