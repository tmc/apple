package ext4

import "encoding/binary"

// groupDescSize is the on-disk size of a non-64bit block group descriptor.
const groupDescSize = 32

// groupDesc is one block group descriptor (32-byte, non-64bit form).
type groupDesc struct {
	blockBitmap     uint32
	inodeBitmap     uint32
	inodeTable      uint32
	freeBlocksCount uint16
	freeInodesCount uint16
	usedDirsCount   uint16
}

// encode serializes the descriptor into 32 bytes.
func (g *groupDesc) encode() []byte {
	b := make([]byte, groupDescSize)
	le := binary.LittleEndian
	le.PutUint32(b[0x00:], g.blockBitmap)
	le.PutUint32(b[0x04:], g.inodeBitmap)
	le.PutUint32(b[0x08:], g.inodeTable)
	le.PutUint16(b[0x0C:], g.freeBlocksCount)
	le.PutUint16(b[0x0E:], g.freeInodesCount)
	le.PutUint16(b[0x10:], g.usedDirsCount)
	// 0x12 bg_flags, 0x14 bg_exclude_bitmap_lo, 0x18.. checksums: all zero.
	return b
}
