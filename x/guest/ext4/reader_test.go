package ext4

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

// reader is a minimal read-only ext4 parser used by the tests to validate the
// images Build produces. It understands only the classic feature set this
// package emits.
type reader struct {
	f              *os.File
	blockSize      int
	inodeSize      int
	inodesPerGroup int
	blocksPerGroup int
	groupCount     int
	inodeTables    []uint32 // per-group inode table start block
}

// rinode is a parsed inode.
type rinode struct {
	mode    uint16
	uid     uint32
	gid     uint32
	size    int64
	links   uint16
	blocks  uint32
	block   [15]uint32
	inline  []byte // first 60 i_block bytes, for symlink target extraction
	atime   time.Time
	mtime   time.Time
	atimeNS uint32
	mtimeNS uint32
}

// open parses the superblock and group descriptors.
func openReader(path string) (*reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	r := &reader{f: f}
	sb := make([]byte, 1024)
	if _, err := f.ReadAt(sb, 1024); err != nil {
		f.Close()
		return nil, fmt.Errorf("read superblock: %w", err)
	}
	le := binary.LittleEndian
	if mag := le.Uint16(sb[0x38:]); mag != superMagic {
		f.Close()
		return nil, fmt.Errorf("bad magic %#x", mag)
	}
	r.blockSize = 1024 << le.Uint32(sb[0x18:])
	r.inodeSize = int(le.Uint16(sb[0x58:]))
	r.inodesPerGroup = int(le.Uint32(sb[0x28:]))
	r.blocksPerGroup = int(le.Uint32(sb[0x20:]))
	inodesCount := le.Uint32(sb[0x00:])
	r.groupCount = ceilDiv(int(inodesCount), r.inodesPerGroup)

	// Read GDT from block 1.
	gdtBytes := make([]byte, r.groupCount*groupDescSize)
	if _, err := f.ReadAt(gdtBytes, int64(r.blockSize)); err != nil {
		f.Close()
		return nil, fmt.Errorf("read gdt: %w", err)
	}
	r.inodeTables = make([]uint32, r.groupCount)
	for g := 0; g < r.groupCount; g++ {
		r.inodeTables[g] = le.Uint32(gdtBytes[g*groupDescSize+0x08:])
	}
	return r, nil
}

func (r *reader) close() error { return r.f.Close() }

// superRaw returns the raw 1024-byte superblock.
func (r *reader) superRaw() ([]byte, error) {
	sb := make([]byte, 1024)
	_, err := r.f.ReadAt(sb, 1024)
	return sb, err
}

// readInode reads and parses inode ino.
func (r *reader) readInode(ino uint32) (*rinode, error) {
	g := int(ino-1) / r.inodesPerGroup
	idx := int(ino-1) % r.inodesPerGroup
	off := int64(r.inodeTables[g])*int64(r.blockSize) + int64(idx)*int64(r.inodeSize)
	buf := make([]byte, r.inodeSize)
	if _, err := r.f.ReadAt(buf, off); err != nil {
		return nil, fmt.Errorf("read inode %d: %w", ino, err)
	}
	le := binary.LittleEndian
	in := &rinode{
		mode:   le.Uint16(buf[0x00:]),
		links:  le.Uint16(buf[0x1A:]),
		blocks: le.Uint32(buf[0x1C:]),
	}
	in.uid = uint32(le.Uint16(buf[0x02:])) | uint32(le.Uint16(buf[0x78:]))<<16
	in.gid = uint32(le.Uint16(buf[0x18:])) | uint32(le.Uint16(buf[0x7A:]))<<16
	sizeLo := uint64(le.Uint32(buf[0x04:]))
	sizeHi := uint64(le.Uint32(buf[0x6C:]))
	in.size = int64(sizeHi<<32 | sizeLo)
	in.atimeNS = le.Uint32(buf[0x8C:]) >> 2
	in.mtimeNS = le.Uint32(buf[0x88:]) >> 2
	in.atime = time.Unix(int64(le.Uint32(buf[0x08:])), int64(in.atimeNS))
	in.mtime = time.Unix(int64(le.Uint32(buf[0x10:])), int64(in.mtimeNS))
	in.inline = append([]byte(nil), buf[0x28:0x28+60]...)
	for i := 0; i < 15; i++ {
		in.block[i] = le.Uint32(buf[0x28+i*4:])
	}
	return in, nil
}

// readBlock reads block n.
func (r *reader) readBlock(n uint32) ([]byte, error) {
	buf := make([]byte, r.blockSize)
	_, err := r.f.ReadAt(buf, int64(n)*int64(r.blockSize))
	return buf, err
}

// dataBlocks returns the data block numbers for an inode in file order,
// following the classic indirect map up to count blocks.
func (r *reader) dataBlocks(in *rinode, count int) ([]uint32, error) {
	ptrs := pointersPerBlock(r.blockSize)
	var out []uint32

	for i := 0; i < 12 && len(out) < count; i++ {
		out = append(out, in.block[i])
	}
	if len(out) >= count {
		return out[:count], nil
	}
	// single
	if err := r.followIndirect(in.block[12], 1, ptrs, count, &out); err != nil {
		return nil, err
	}
	if len(out) >= count {
		return out[:count], nil
	}
	if err := r.followIndirect(in.block[13], 2, ptrs, count, &out); err != nil {
		return nil, err
	}
	if len(out) >= count {
		return out[:count], nil
	}
	if err := r.followIndirect(in.block[14], 3, ptrs, count, &out); err != nil {
		return nil, err
	}
	if len(out) > count {
		out = out[:count]
	}
	return out, nil
}

func (r *reader) followIndirect(blk uint32, level, ptrs, count int, out *[]uint32) error {
	if blk == 0 || len(*out) >= count {
		return nil
	}
	buf, err := r.readBlock(blk)
	if err != nil {
		return err
	}
	le := binary.LittleEndian
	for i := 0; i < ptrs && len(*out) < count; i++ {
		p := le.Uint32(buf[i*4:])
		if p == 0 {
			continue
		}
		if level == 1 {
			*out = append(*out, p)
		} else {
			if err := r.followIndirect(p, level-1, ptrs, count, out); err != nil {
				return err
			}
		}
	}
	return nil
}

// readContent reads size bytes of a file's data following its block map.
func (r *reader) readContent(in *rinode) ([]byte, error) {
	if in.size == 0 {
		return nil, nil
	}
	count := ceilDiv(int(in.size), r.blockSize)
	blocks, err := r.dataBlocks(in, count)
	if err != nil {
		return nil, err
	}
	out := make([]byte, 0, in.size)
	for _, b := range blocks {
		buf, err := r.readBlock(b)
		if err != nil {
			return nil, err
		}
		out = append(out, buf...)
	}
	return out[:in.size], nil
}

// symlinkTarget returns the target of a symlink inode.
func (r *reader) symlinkTarget(in *rinode) (string, error) {
	if in.size <= 60 {
		return string(in.inline[:in.size]), nil
	}
	content, err := r.readContent(in)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// rdirEntry is a parsed directory entry.
type rdirEntry struct {
	ino      uint32
	name     string
	fileType uint8
}

// readDir parses the linear directory entries of a directory inode.
func (r *reader) readDir(in *rinode) ([]rdirEntry, error) {
	count := ceilDiv(int(in.size), r.blockSize)
	blocks, err := r.dataBlocks(in, count)
	if err != nil {
		return nil, err
	}
	le := binary.LittleEndian
	var entries []rdirEntry
	for _, b := range blocks {
		buf, err := r.readBlock(b)
		if err != nil {
			return nil, err
		}
		off := 0
		for off+8 <= len(buf) {
			ino := le.Uint32(buf[off:])
			recLen := int(le.Uint16(buf[off+4:]))
			nameLen := int(buf[off+6])
			ft := buf[off+7]
			if recLen < 8 {
				break
			}
			if ino != 0 {
				name := string(buf[off+8 : off+8+nameLen])
				entries = append(entries, rdirEntry{ino: ino, name: name, fileType: ft})
			}
			off += recLen
		}
	}
	return entries, nil
}
