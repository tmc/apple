package ext4

import (
	"fmt"
	"sort"
	"time"
)

// item is a single ext4 inode to be written, along with the directory entries
// it contains (for directories) and the data blocks it owns.
type item struct {
	ino     uint32
	n       *node // source node (nil for the synthetic lost+found)
	entries []dirEntry
	mode    uint16
	uid     uint32
	gid     uint32
	size    int64
	links   uint16
	atime   time.Time
	mtime   time.Time
	ctime   time.Time

	isDir     bool
	isSymlink bool
	fastLink  []byte // inline symlink target, when <= 60 bytes

	dataBlocks []uint32 // file/dir data blocks, in order
	dirBlocks  [][]byte // encoded directory blocks (dirs only)
	bmap       blockMap
}

// layout is the computed on-disk arrangement.
type layout struct {
	blockSize      int
	blocksPerGroup int
	inodesPerGroup int
	groupCount     int
	blocksCount    uint32
	inodesCount    uint32

	gdtBlocks      int // blocks occupied by the group descriptor table
	inodeTblBlocks int // blocks per group for the inode table

	items   []*item          // ordered by inode number (index = ino-1 for assigned)
	byIno   map[uint32]*item // assigned inodes only
	blocks  *bitmap          // global block bitmap (all groups, flat)
	inodes  *bitmap          // global inode bitmap (all groups, flat)
	usedDir []int            // used-dir count per group

	mkfsTime time.Time
	uuid     [16]byte
}

// newLayout assigns inodes and blocks for the tree rooted at root.
func newLayout(root *node, blockSize, blocksPerGroup int) (*layout, error) {
	l := &layout{
		blockSize:      blockSize,
		blocksPerGroup: blocksPerGroup,
		byIno:          map[uint32]*item{},
	}

	// Phase 1: assign inode numbers and build directory entry lists.
	if err := l.assignInodes(root); err != nil {
		return nil, err
	}

	// Phase 2: fix filesystem geometry from inode and rough data-block counts.
	if err := l.computeGeometry(); err != nil {
		return nil, err
	}

	// Phase 3: allocate data + indirect blocks against the bitmaps.
	if err := l.allocate(); err != nil {
		return nil, err
	}

	return l, nil
}

// assignInodes walks the tree assigning inode numbers (root=2, lost+found=11,
// then 12..) and constructs the directory entry lists.
func (l *layout) assignInodes(root *node) error {
	next := uint32(firstIno + 1) // 12

	// Reserve lost+found at inode 11.
	lostFound := &item{
		ino:   firstIno,
		mode:  sIFDIR | 0o700,
		isDir: true,
		links: 2,
		atime: root.atime,
		mtime: root.mtime,
		ctime: root.ctime,
	}

	root.ino = 2
	rootItem := l.newDirItem(root, 2)

	// Root contains "." -> 2, ".." -> 2, "lost+found" -> 11, then children.
	rootItem.entries = []dirEntry{
		{ino: 2, name: ".", fileType: ftDir},
		{ino: 2, name: "..", fileType: ftDir},
		{ino: firstIno, name: "lost+found", fileType: ftDir},
	}
	lostFound.entries = []dirEntry{
		{ino: firstIno, name: ".", fileType: ftDir},
		{ino: 2, name: "..", fileType: ftDir},
	}

	l.add(rootItem)
	l.add(lostFound)

	// link count for root: 2 (self + ".") plus one per immediate subdir (their
	// "..") plus lost+found.
	rootLinks := 2 + 1 // base + lost+found
	for _, c := range root.children {
		if c.kind == kindDir {
			rootLinks++
		}
	}
	rootItem.links = uint16(rootLinks)

	// Depth-first assignment over the children, in sorted order.
	var assign func(parent *node, parentItem *item) error
	assign = func(parent *node, parentItem *item) error {
		for _, c := range parent.children {
			var ino uint32
			switch {
			case c.kind == kindReg && c.shared != nil:
				if c.shared.ino == 0 {
					c.shared.ino = next
					next++
					it := l.newRegItem(c, c.shared.ino)
					it.links = uint16(c.shared.links)
					l.add(it)
				}
				ino = c.shared.ino
			default:
				ino = next
				next++
				c.ino = ino
				switch c.kind {
				case kindDir:
					it := l.newDirItem(c, ino)
					links := 2 // self + "."
					for _, gc := range c.children {
						if gc.kind == kindDir {
							links++
						}
					}
					it.links = uint16(links)
					it.entries = []dirEntry{
						{ino: ino, name: ".", fileType: ftDir},
						{ino: parentItem.ino, name: "..", fileType: ftDir},
					}
					l.add(it)
				case kindReg:
					it := l.newRegItem(c, ino)
					it.links = 1
					l.add(it)
				case kindSymlink:
					it := l.newSymlinkItem(c, ino)
					it.links = 1
					l.add(it)
				}
			}

			ft := fileTypeOf(c)
			parentItem.entries = append(parentItem.entries, dirEntry{
				ino: ino, name: c.name, fileType: ft,
			})

			if c.kind == kindDir && c.shared == nil {
				if err := assign(c, l.byIno[ino]); err != nil {
					return err
				}
			}
		}
		return nil
	}
	if err := assign(root, rootItem); err != nil {
		return err
	}

	l.inodesCount = next - 1
	return nil
}

func fileTypeOf(n *node) uint8 {
	switch n.kind {
	case kindDir:
		return ftDir
	case kindSymlink:
		return ftSymlink
	default:
		return ftReg
	}
}

func (l *layout) add(it *item) {
	l.items = append(l.items, it)
	l.byIno[it.ino] = it
}

func (l *layout) newDirItem(n *node, ino uint32) *item {
	return &item{
		ino: ino, n: n, isDir: true,
		mode: n.mode, uid: n.uid, gid: n.gid,
		atime: n.atime, mtime: n.mtime, ctime: n.ctime,
	}
}

func (l *layout) newRegItem(n *node, ino uint32) *item {
	return &item{
		ino: ino, n: n,
		mode: n.mode, uid: n.uid, gid: n.gid, size: n.size,
		atime: n.atime, mtime: n.mtime, ctime: n.ctime,
	}
}

func (l *layout) newSymlinkItem(n *node, ino uint32) *item {
	it := &item{
		ino: ino, n: n, isSymlink: true,
		mode: n.mode, uid: n.uid, gid: n.gid, size: n.size,
		atime: n.atime, mtime: n.mtime, ctime: n.ctime,
	}
	if len(n.link) <= 60 {
		it.fastLink = []byte(n.link)
	}
	return it
}

// computeGeometry fixes inodes-per-group, group count, and metadata block
// layout, then sizes the filesystem to hold all data with headroom.
func (l *layout) computeGeometry() error {
	bs := l.blockSize
	ptrs := pointersPerBlock(bs)

	// Estimate data blocks required (file + directory + indirect blocks).
	var dataBlocks int
	for _, it := range l.items {
		nb := l.itemDataBlocks(it)
		dataBlocks += nb + countIndirectBlocks(nb, ptrs)
	}

	// Inode count: reserved (firstIno) + assigned beyond. We already counted
	// assigned in inodesCount; ensure at least firstIno reserved exist.
	inodes := int(l.inodesCount)
	if inodes < firstIno {
		inodes = firstIno
	}
	// A little inode headroom.
	inodes += inodes/10 + 16

	// Iterate group count: metadata depends on group count, which depends on
	// total blocks, which depends on metadata. Converge upward.
	groups := 1
	for {
		inodesPerGroup := ceilDiv(inodes, groups)
		// inodes_per_group must be a multiple of 8 (bitmap byte alignment) and
		// fit the inode bitmap (one block).
		inodesPerGroup = (inodesPerGroup + 7) &^ 7
		if inodesPerGroup > bs*8 {
			inodesPerGroup = bs * 8
		}
		inodeTblBlocks := ceilDiv(inodesPerGroup*inodeSize, bs)
		gdtBlocks := ceilDiv(groups*groupDescSize, bs)

		// Per-group metadata = SB(1) + GDT + block bitmap(1) + inode bitmap(1)
		// + inode table. Group 0's SB shares block 0 with the boot record.
		perGroupMeta := 1 + gdtBlocks + 1 + 1 + inodeTblBlocks
		dataPerGroup := l.blocksPerGroup - perGroupMeta
		if dataPerGroup <= 0 {
			return fmt.Errorf("blocks_per_group %d too small for metadata", l.blocksPerGroup)
		}

		// total inodes available with this geometry
		totalInodes := inodesPerGroup * groups

		neededDataGroups := ceilDiv(dataBlocks, dataPerGroup)
		neededInodeGroups := ceilDiv(inodes, inodesPerGroup)
		need := neededDataGroups
		if neededInodeGroups > need {
			need = neededInodeGroups
		}
		if need < 1 {
			need = 1
		}
		if need > groups {
			groups = need
			continue
		}

		// Converged. Add a bit of free-block headroom to the final group.
		l.groupCount = groups
		l.inodesPerGroup = inodesPerGroup
		l.inodeTblBlocks = inodeTblBlocks
		l.gdtBlocks = gdtBlocks
		l.inodesCount = uint32(totalInodes)
		l.blocksCount = uint32(groups * l.blocksPerGroup)

		// The last group may extend past needed data; that is fine, the extra
		// blocks are simply free. But the total must not exceed addressable
		// range; for our sizes it never does.
		_ = totalInodes
		break
	}

	l.usedDir = make([]int, l.groupCount)
	l.blocks = newBitmap(int(l.blocksCount))
	l.inodes = newBitmap(l.groupCount * l.inodesPerGroup)
	return nil
}

// itemDataBlocks returns the number of (non-indirect) data blocks an item's
// content occupies.
func (l *layout) itemDataBlocks(it *item) int {
	bs := l.blockSize
	switch {
	case it.isSymlink:
		if it.fastLink != nil {
			return 0
		}
		return ceilDiv(int(it.size), bs)
	case it.isDir:
		// Directory blocks are determined by packing entries; estimate by
		// total entry length (the exact value is computed at write time, but
		// for sizing the conservative ceil is correct because packing never
		// uses more blocks than ceil(totalLen/bs)+1).
		total := 0
		for _, e := range it.entries {
			total += entryLen(e.name)
		}
		nb := ceilDiv(total, bs)
		if nb == 0 {
			nb = 1
		}
		// Packing may push an entry to a new block; allow one extra.
		return nb + 1
	default:
		return ceilDiv(int(it.size), bs)
	}
}

// groupMetaBlocks returns, for group g, the first block of its data region and
// the block numbers of its block bitmap, inode bitmap, and inode table start.
type groupMeta struct {
	firstBlock  uint32 // first block of the group
	hasSuper    bool
	superBlock  uint32 // block holding SB+GDT backup (sparse_super OFF: always)
	gdtStart    uint32
	blockBitmap uint32
	inodeBitmap uint32
	inodeTable  uint32
	dataStart   uint32
	dataEnd     uint32 // exclusive
}

func (l *layout) group(g int) groupMeta {
	bs := l.blockSize
	_ = bs
	first := uint32(g * l.blocksPerGroup)
	gm := groupMeta{firstBlock: first, hasSuper: true}

	cur := first
	if g == 0 {
		// Block 0 holds boot record + primary superblock; GDT follows at block 1.
		gm.superBlock = 0
		cur = 1
		gm.gdtStart = cur
		cur += uint32(l.gdtBlocks)
	} else {
		// Backup SB at first block, GDT after.
		gm.superBlock = first
		cur = first + 1
		gm.gdtStart = cur
		cur += uint32(l.gdtBlocks)
	}
	gm.blockBitmap = cur
	cur++
	gm.inodeBitmap = cur
	cur++
	gm.inodeTable = cur
	cur += uint32(l.inodeTblBlocks)
	gm.dataStart = cur
	gm.dataEnd = first + uint32(l.blocksPerGroup)
	return gm
}

// allocate reserves metadata blocks in the bitmap, then allocates data and
// indirect blocks for every item via deterministic first-fit.
func (l *layout) allocate() error {
	// Reserve metadata blocks for every group.
	for g := 0; g < l.groupCount; g++ {
		gm := l.group(g)
		for blk := gm.firstBlock; blk < gm.dataStart; blk++ {
			l.blocks.set(int(blk))
		}
	}
	// Mark inode bitmap padding (inodes beyond count within last group bitmap
	// handled below) and reserved inodes 1..firstIno.
	for ino := 1; ino <= firstIno; ino++ {
		l.inodes.set(ino - 1)
	}
	// Mark every assigned inode as used.
	for ino := uint32(1); ino <= l.inodesCount; ino++ {
		if ino <= firstIno {
			continue // already reserved
		}
		if _, ok := l.byIno[ino]; ok {
			l.inodes.set(int(ino) - 1)
		}
	}
	// inode 2 (root) is within 1..firstIno reserved range, already set.

	// Deterministic block allocator: first-fit forward scan.
	nextFree := l.firstDataBlock()
	allocOne := func() (uint32, error) {
		for {
			if nextFree >= int(l.blocksCount) {
				return 0, fmt.Errorf("out of blocks during allocation")
			}
			if !l.blocks.get(nextFree) {
				l.blocks.set(nextFree)
				b := uint32(nextFree)
				nextFree++
				return b, nil
			}
			nextFree++
		}
	}

	ptrs := pointersPerBlock(l.blockSize)

	// Allocate per item in inode order for reproducibility.
	order := make([]*item, len(l.items))
	copy(order, l.items)
	sort.Slice(order, func(i, j int) bool { return order[i].ino < order[j].ino })

	for _, it := range order {
		nb := l.contentBlockCount(it)
		if nb == 0 {
			continue
		}
		data := make([]uint32, nb)
		for i := range data {
			blk, err := allocOne()
			if err != nil {
				return err
			}
			data[i] = blk
		}
		it.dataBlocks = data
		bmap, err := buildBlockMap(data, ptrs, allocOne)
		if err != nil {
			return err
		}
		it.bmap = bmap

		if it.isDir {
			grp := int(it.ino-1) / l.inodesPerGroup
			if grp >= 0 && grp < len(l.usedDir) {
				l.usedDir[grp]++
			}
		}
	}
	return nil
}

// firstDataBlock returns the lowest data block (group 0's data start).
func (l *layout) firstDataBlock() int {
	return int(l.group(0).dataStart)
}

// contentBlockCount returns the exact non-indirect block count for an item,
// using real directory packing where applicable.
func (l *layout) contentBlockCount(it *item) int {
	switch {
	case it.isSymlink:
		if it.fastLink != nil {
			return 0
		}
		return ceilDiv(int(it.size), l.blockSize)
	case it.isDir:
		blocks := packDirBlocks(it.entries, l.blockSize)
		it.dirBlocks = blocks
		return len(blocks)
	default:
		return ceilDiv(int(it.size), l.blockSize)
	}
}
