package ext4

import "encoding/binary"

// pointersPerBlock is the number of LE u32 block pointers in one indirect
// block (4096 / 4 = 1024).
func pointersPerBlock(blockSize int) int { return blockSize / 4 }

// indirectBlock is an allocated indirect block: its own block number plus the
// pointer slice to be written into it.
type indirectBlock struct {
	block uint32
	data  []uint32
}

// blockMap holds the result of mapping a sequence of data block numbers onto
// the classic ext4 indirect layout: the 15-entry i_block array and any
// indirect blocks that must be written out.
type blockMap struct {
	iBlock    [15]uint32
	indirects []indirectBlock // includes single/double/triple at every level
}

// alloc allocates one block number.
type alloc func() (uint32, error)

// buildBlockMap maps data[i] (the file's i-th data block) into the classic
// 12 direct + single + double + triple indirect structure, allocating indirect
// blocks via the supplied allocator. data holds the already-allocated data
// block numbers in file order.
func buildBlockMap(data []uint32, ptrs int, alloc alloc) (blockMap, error) {
	var m blockMap
	n := len(data)
	i := 0

	// 12 direct pointers.
	for ; i < 12 && i < n; i++ {
		m.iBlock[i] = data[i]
	}
	if i >= n {
		return m, nil
	}

	// Single indirect: ptrs blocks.
	if rem := n - i; rem > 0 {
		take := rem
		if take > ptrs {
			take = ptrs
		}
		single, err := alloc()
		if err != nil {
			return m, err
		}
		m.iBlock[12] = single
		ind := indirectBlock{block: single, data: append([]uint32(nil), data[i:i+take]...)}
		m.indirects = append(m.indirects, ind)
		i += take
	}
	if i >= n {
		return m, nil
	}

	// Double indirect: ptrs single-indirect blocks, each mapping ptrs data.
	if rem := n - i; rem > 0 {
		cap2 := ptrs * ptrs
		take := rem
		if take > cap2 {
			take = cap2
		}
		double, err := alloc()
		if err != nil {
			return m, err
		}
		m.iBlock[13] = double
		if err := m.fillIndirect(2, double, data[i:i+take], ptrs, alloc); err != nil {
			return m, err
		}
		i += take
	}
	if i >= n {
		return m, nil
	}

	// Triple indirect: the remainder.
	rem := n - i
	triple, err := alloc()
	if err != nil {
		return m, err
	}
	m.iBlock[14] = triple
	if err := m.fillIndirect(3, triple, data[i:i+rem], ptrs, alloc); err != nil {
		return m, err
	}
	return m, nil
}

// fillIndirect populates an indirect block of the given level (1=single,
// 2=double, 3=triple) that addresses the slice data, allocating any child
// indirect blocks. The block at parent has already been allocated.
func (m *blockMap) fillIndirect(level int, parent uint32, data []uint32, ptrs int, alloc alloc) error {
	if level == 1 {
		m.indirects = append(m.indirects, indirectBlock{
			block: parent,
			data:  append([]uint32(nil), data...),
		})
		return nil
	}
	// Each child at the next-lower level addresses childCap data blocks.
	childCap := 1
	for k := 1; k < level; k++ {
		childCap *= ptrs
	}
	var ptrSlice []uint32
	for off := 0; off < len(data); off += childCap {
		end := off + childCap
		if end > len(data) {
			end = len(data)
		}
		child, err := alloc()
		if err != nil {
			return err
		}
		ptrSlice = append(ptrSlice, child)
		if err := m.fillIndirect(level-1, child, data[off:end], ptrs, alloc); err != nil {
			return err
		}
	}
	m.indirects = append(m.indirects, indirectBlock{block: parent, data: ptrSlice})
	return nil
}

// countIndirectBlocks returns how many indirect blocks are required to map n
// data blocks (used for sizing before allocation).
func countIndirectBlocks(n, ptrs int) int {
	if n <= 12 {
		return 0
	}
	count := 0
	rem := n - 12

	// Single indirect.
	count++ // the single-indirect block itself
	if rem <= ptrs {
		return count
	}
	rem -= ptrs

	// Double indirect.
	cap2 := ptrs * ptrs
	d := rem
	if d > cap2 {
		d = cap2
	}
	count++                   // the double-indirect block
	count += ceilDiv(d, ptrs) // its single-indirect children
	if rem <= cap2 {
		return count
	}
	rem -= cap2

	// Triple indirect (remainder).
	count++ // the triple-indirect block
	doubles := ceilDiv(rem, cap2)
	count += doubles // its double-indirect children
	count += ceilDiv(rem, ptrs)
	return count
}

func ceilDiv(a, b int) int { return (a + b - 1) / b }

// encodeIndirect serializes an indirect block's pointer slice into a full
// block, zero-padding the unused tail.
func encodeIndirect(ind indirectBlock, blockSize int) []byte {
	b := make([]byte, blockSize)
	le := binary.LittleEndian
	for i, p := range ind.data {
		le.PutUint32(b[i*4:], p)
	}
	return b
}
