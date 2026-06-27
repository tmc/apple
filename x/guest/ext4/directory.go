package ext4

import "encoding/binary"

// Directory entry file types (ext2_dir_entry_2.file_type), enabled by the
// FILETYPE incompat feature.
const (
	ftReg     = 1
	ftDir     = 2
	ftSymlink = 7
)

// dirEntry is a linear directory entry to be serialized as ext2_dir_entry_2.
type dirEntry struct {
	ino      uint32
	name     string
	fileType uint8
}

// entryLen returns the minimum on-disk record length for a name: the 8-byte
// header plus the name rounded up to a 4-byte boundary.
func entryLen(name string) int {
	return (8 + len(name) + 3) &^ 3
}

// packDirBlocks lays out entries into 4096-byte directory blocks. Each block
// begins fresh, no entry straddles a block boundary, and the final entry in
// every block has its rec_len extended to the end of the block.
func packDirBlocks(entries []dirEntry, blockSize int) [][]byte {
	var blocks [][]byte
	var cur []dirEntry
	used := 0

	flush := func() {
		if len(cur) == 0 {
			return
		}
		blocks = append(blocks, encodeDirBlock(cur, blockSize))
		cur = nil
		used = 0
	}

	for _, e := range entries {
		need := entryLen(e.name)
		if used+need > blockSize {
			flush()
		}
		cur = append(cur, e)
		used += need
	}
	flush()
	return blocks
}

// encodeDirBlock serializes one directory block; the last entry's rec_len is
// stretched to fill the block.
func encodeDirBlock(entries []dirEntry, blockSize int) []byte {
	b := make([]byte, blockSize)
	le := binary.LittleEndian
	off := 0
	for i, e := range entries {
		recLen := entryLen(e.name)
		if i == len(entries)-1 {
			recLen = blockSize - off
		}
		le.PutUint32(b[off:], e.ino)
		le.PutUint16(b[off+4:], uint16(recLen))
		b[off+6] = uint8(len(e.name))
		b[off+7] = e.fileType
		copy(b[off+8:], e.name)
		off += recLen
	}
	return b
}
