package ext4

import (
	"fmt"
	"os"
)

// write serializes the whole filesystem into the image.
func (l *layout) write(im *image) error {
	if err := l.writeSuperAndGDT(im); err != nil {
		return err
	}
	if err := l.writeBitmaps(im); err != nil {
		return err
	}
	if err := l.writeInodesAndData(im); err != nil {
		return err
	}
	return nil
}

// buildSuperblock returns the superblock with final counts.
func (l *layout) buildSuperblock() *superblock {
	return &superblock{
		inodesCount:     l.inodesCount,
		blocksCount:     l.blocksCount,
		freeBlocksCount: uint32(l.blocks.countClear()),
		freeInodesCount: uint32(l.inodes.countClear()),
		firstDataBlock:  0, // 4K blocks: superblock lives in block 0
		blockSize:       uint32(l.blockSize),
		blocksPerGroup:  uint32(l.blocksPerGroup),
		inodesPerGroup:  uint32(l.inodesPerGroup),
		inodeSize:       inodeSize,
		uuid:            l.uuid,
		groupCount:      uint32(l.groupCount),
		mkfsTime:        l.mkfsTime,
	}
}

// buildGDT returns the encoded group descriptor table (all descriptors).
func (l *layout) buildGDT() []byte {
	gdt := make([]byte, l.gdtBlocks*l.blockSize)
	for g := 0; g < l.groupCount; g++ {
		gm := l.group(g)
		gd := &groupDesc{
			blockBitmap:     gm.blockBitmap,
			inodeBitmap:     gm.inodeBitmap,
			inodeTable:      gm.inodeTable,
			freeBlocksCount: uint16(l.groupFreeBlocks(g)),
			freeInodesCount: uint16(l.groupFreeInodes(g)),
			usedDirsCount:   uint16(l.usedDir[g]),
		}
		copy(gdt[g*groupDescSize:], gd.encode())
	}
	return gdt
}

// writeSuperAndGDT writes the primary superblock and GDT plus a full backup
// copy at the start of every other group (sparse_super OFF).
func (l *layout) writeSuperAndGDT(im *image) error {
	sb := l.buildSuperblock()
	gdt := l.buildGDT()

	for g := 0; g < l.groupCount; g++ {
		gm := l.group(g)
		enc := sb.encode()
		setBlockGroupNr(enc, g)
		if g == 0 {
			// Primary SB at byte offset 1024 (block 0 holds boot+SB).
			if err := im.writeAt(1024, enc); err != nil {
				return err
			}
		} else {
			// Backup SB occupies the start of the group's first block.
			if err := im.writeBlock(gm.superBlock, enc); err != nil {
				return err
			}
		}
		// GDT (primary or backup) at gdtStart.
		for i := 0; i < l.gdtBlocks; i++ {
			start := i * l.blockSize
			end := start + l.blockSize
			if end > len(gdt) {
				end = len(gdt)
			}
			if err := im.writeBlock(gm.gdtStart+uint32(i), gdt[start:end]); err != nil {
				return err
			}
		}
	}
	return nil
}

// setBlockGroupNr writes s_block_group_nr (offset 0x5A) into an encoded SB.
func setBlockGroupNr(sb []byte, g int) {
	sb[0x5A] = byte(g)
	sb[0x5B] = byte(g >> 8)
}

// writeBitmaps writes per-group block and inode bitmaps, sliced from the global
// bitmaps with trailing padding bits set.
func (l *layout) writeBitmaps(im *image) error {
	bs := l.blockSize
	for g := 0; g < l.groupCount; g++ {
		gm := l.group(g)

		// Block bitmap: blocksPerGroup bits.
		bb := make([]byte, bs)
		base := g * l.blocksPerGroup
		for i := 0; i < l.blocksPerGroup; i++ {
			if l.blocks.get(base + i) {
				bb[i/8] |= 1 << uint(i%8)
			}
		}
		// Trailing bits beyond blocksPerGroup within the block: set as used.
		for i := l.blocksPerGroup; i < bs*8; i++ {
			bb[i/8] |= 1 << uint(i%8)
		}
		if err := im.writeBlock(gm.blockBitmap, bb); err != nil {
			return err
		}

		// Inode bitmap: inodesPerGroup bits.
		ib := make([]byte, bs)
		ibase := g * l.inodesPerGroup
		for i := 0; i < l.inodesPerGroup; i++ {
			if l.inodes.get(ibase + i) {
				ib[i/8] |= 1 << uint(i%8)
			}
		}
		for i := l.inodesPerGroup; i < bs*8; i++ {
			ib[i/8] |= 1 << uint(i%8)
		}
		if err := im.writeBlock(gm.inodeBitmap, ib); err != nil {
			return err
		}
	}
	return nil
}

// writeInodesAndData writes every inode into its group's inode table and the
// associated data and indirect blocks.
func (l *layout) writeInodesAndData(im *image) error {
	for _, it := range l.items {
		in := l.buildInode(it)
		if err := l.writeInode(im, it.ino, in); err != nil {
			return err
		}
		if err := l.writeItemData(im, it); err != nil {
			return err
		}
	}
	return nil
}

// writeInode writes a single inode into the correct group inode table slot.
func (l *layout) writeInode(im *image, ino uint32, in *inode) error {
	g := int(ino-1) / l.inodesPerGroup
	idx := int(ino-1) % l.inodesPerGroup
	gm := l.group(g)
	off := int64(gm.inodeTable)*int64(l.blockSize) + int64(idx)*inodeSize
	return im.writeAt(off, in.encode())
}

// buildInode constructs the on-disk inode for an item.
func (l *layout) buildInode(it *item) *inode {
	in := &inode{
		mode:  it.mode,
		uid:   it.uid,
		gid:   it.gid,
		size:  it.size,
		links: it.links,
		atime: it.atime,
		mtime: it.mtime,
		ctime: it.ctime,
	}
	if it.isDir {
		// Directory size is the byte length of its data blocks.
		in.size = int64(len(it.dirBlocks) * l.blockSize)
	}
	if it.fastLink != nil {
		in.inlined = it.fastLink
		in.size = int64(len(it.fastLink))
		in.blocks = 0
		return in
	}
	in.bmapInto(it.bmap)
	// i_blocks counts every 512-byte sector of data + indirect blocks.
	total := len(it.dataBlocks) + len(it.bmap.indirects)
	in.blocks = uint32(total * (l.blockSize / 512))
	return in
}

// bmapInto copies a block map into the inode's i_block array.
func (in *inode) bmapInto(m blockMap) {
	in.block = m.iBlock
}

// writeItemData writes an item's data blocks and indirect blocks.
func (l *layout) writeItemData(im *image, it *item) error {
	if it.fastLink != nil {
		return nil
	}
	// Data blocks.
	if it.isDir {
		for i, blk := range it.dataBlocks {
			if err := im.writeBlock(blk, it.dirBlocks[i]); err != nil {
				return err
			}
		}
	} else if it.isSymlink {
		// Slow symlink: target in one or more data blocks.
		target := []byte(it.n.link)
		if err := l.writeBytes(im, it.dataBlocks, target); err != nil {
			return err
		}
	} else {
		if err := l.writeFile(im, it); err != nil {
			return err
		}
	}
	// Indirect blocks.
	for _, ind := range it.bmap.indirects {
		if err := im.writeBlock(ind.block, encodeIndirect(ind, l.blockSize)); err != nil {
			return err
		}
	}
	return nil
}

// writeBytes writes a byte slice across the given data blocks.
func (l *layout) writeBytes(im *image, blocks []uint32, data []byte) error {
	bs := l.blockSize
	for i, blk := range blocks {
		start := i * bs
		end := start + bs
		if end > len(data) {
			end = len(data)
		}
		if err := im.writeBlock(blk, data[start:end]); err != nil {
			return err
		}
	}
	return nil
}

// writeFile streams a regular file's content into its data blocks.
func (l *layout) writeFile(im *image, it *item) error {
	if it.size == 0 {
		return nil
	}
	f, err := os.Open(it.n.path)
	if err != nil {
		return fmt.Errorf("open %s: %w", it.n.path, err)
	}
	defer f.Close()

	bs := l.blockSize
	buf := make([]byte, bs)
	for _, blk := range it.dataBlocks {
		n, err := readFull(f, buf)
		if n > 0 {
			if werr := im.writeBlock(blk, buf[:n]); werr != nil {
				return werr
			}
		}
		if err != nil {
			break
		}
	}
	return nil
}

// readFull reads up to len(buf) bytes, returning the count and a non-nil error
// only at EOF after a short read.
func readFull(f *os.File, buf []byte) (int, error) {
	total := 0
	for total < len(buf) {
		n, err := f.Read(buf[total:])
		total += n
		if err != nil {
			return total, err
		}
	}
	return total, nil
}

// groupFreeBlocks counts clear bits in group g's block bitmap range.
func (l *layout) groupFreeBlocks(g int) int {
	base := g * l.blocksPerGroup
	free := 0
	for i := 0; i < l.blocksPerGroup; i++ {
		if !l.blocks.get(base + i) {
			free++
		}
	}
	return free
}

// groupFreeInodes counts clear bits in group g's inode bitmap range.
func (l *layout) groupFreeInodes(g int) int {
	base := g * l.inodesPerGroup
	free := 0
	for i := 0; i < l.inodesPerGroup; i++ {
		if !l.inodes.get(base + i) {
			free++
		}
	}
	return free
}
