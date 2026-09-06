package jaccl

import "fmt"

const ringLargeMessageBytes = 64 << 10

// reduceRing applies the accumulation order used by standalone JACCL's ring
// reduce-scatter. Each incoming block has already been reduced upstream, so
// the order is right-associated and observable for non-associative sums and
// NaNs.
func reduceRing(dst, gathered []byte, size, blockSize int, dtype DType, op ReduceOp, wires int) error {
	width, err := dtype.Size()
	if err != nil {
		return err
	}
	if size < 2 || wires < 1 || blockSize < 0 || blockSize%width != 0 || len(dst) != blockSize || len(gathered) != size*blockSize || !op.valid() {
		return fmt.Errorf("ring reduce shape: %w", ErrProtocol)
	}
	count := blockSize / width
	if count == 0 {
		return nil
	}
	chunkSize, split, reverse := ringReductionPlan(count, size, wires, blockSize)
	for element := 0; element < count; element++ {
		chunk, offset := element/chunkSize, element%chunkSize
		start, step := (chunk+1)%size, 1
		if reverse || offset < split {
			start, step = (chunk+size-1)%size, -1
		}
		byteOffset := element * width
		var accumulated [8]byte
		last := (start + step*(size-1) + size) % size
		copy(accumulated[:width], gathered[last*blockSize+byteOffset:last*blockSize+byteOffset+width])
		for rank := size - 2; rank >= 0; rank-- {
			peer := (start + step*rank + size) % size
			value := dst[byteOffset : byteOffset+width]
			copy(value, gathered[peer*blockSize+byteOffset:peer*blockSize+byteOffset+width])
			if err := reduceValue(value, accumulated[:width], dtype, op); err != nil {
				return err
			}
			copy(accumulated[:width], value)
		}
	}
	return nil
}

// ringReductionPlan mirrors RingGroup::all_reduce and RingImpl::all_reduce.
// A one-direction path reduces every chunk counter-clockwise. The two-way
// path reduces its first per-chunk region counter-clockwise and its remainder
// clockwise. split is unused when reverse is true.
func ringReductionPlan(count, size, wires, bytes int) (chunkSize, split int, reverse bool) {
	chunkSize = (count + size - 1) / size
	if count < size*2*wires {
		return chunkSize, chunkSize, true
	}
	usedWires := 1
	if bytes > ringLargeMessageBytes {
		usedWires = wires
	}
	perWire := (chunkSize + 2*usedWires - 1) / (2 * usedWires)
	return chunkSize, usedWires * perWire, false
}

// meshUsesRingReduction mirrors MeshGroup::all_reduce. Large mesh reductions
// use its internal one-wire RingImpl rather than MeshImpl's rank-major path.
func meshUsesRingReduction(size, blockSize int, dtype DType) bool {
	if size <= 2 {
		return false
	}
	width, err := dtype.Size()
	if err != nil || blockSize < 0 || blockSize%width != 0 {
		return false
	}
	count := blockSize / width
	if dtype == BFloat16 && count > 256*1024 {
		return true
	}
	return count >= (8*1024*1024)/width
}
