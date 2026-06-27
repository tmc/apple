package ext4

import (
	"encoding/binary"
	"time"
)

// File-type bits for i_mode (octal, matching POSIX S_IF*).
const (
	sIFLNK  = 0xA000
	sIFREG  = 0x8000
	sIFDIR  = 0x4000
	sIFMASK = 0xF000
)

// inodeSize is the on-disk inode size for this filesystem.
const inodeSize = 256

// inode holds the fields needed to serialize a 256-byte ext4_inode.
//
// Field offsets follow the canonical ext4 on-disk layout (struct ext4_inode):
//
//	i_mode          0x00 u16
//	i_uid           0x02 u16   (low 16 bits)
//	i_size_lo       0x04 u32
//	i_atime         0x08 u32
//	i_ctime         0x0C u32
//	i_mtime         0x10 u32
//	i_dtime         0x14 u32
//	i_gid           0x18 u16   (low 16 bits)
//	i_links_count   0x1A u16
//	i_blocks_lo     0x1C u32   (count of 512-byte sectors, incl. indirect)
//	i_flags         0x20 u32
//	i_block[15]     0x28 60 bytes (direct[12] + single + double + triple)
//	i_generation    0x64 u32
//	i_file_acl_lo   0x68 u32
//	i_size_high     0x6C u32   (a.k.a. i_dir_acl for regular files)
//	i_obso_faddr    0x70 u32
//	osd2.l_i_blocks_high  0x74 u16
//	osd2.l_i_file_acl_high 0x76 u16
//	osd2.l_i_uid_high     0x78 u16
//	osd2.l_i_gid_high     0x7A u16
//	osd2.l_i_checksum_lo  0x7C u16
//	osd2.l_i_reserved     0x7E u16
//	i_extra_isize   0x80 u16   (=32)
//	i_checksum_hi   0x82 u16
//	i_ctime_extra   0x84 u32   (nsec<<2 | epoch; epoch=0)
//	i_mtime_extra   0x88 u32
//	i_atime_extra   0x8C u32
//	i_crtime        0x90 u32
//	i_crtime_extra  0x94 u32
type inode struct {
	mode    uint16
	uid     uint32
	gid     uint32
	size    int64
	links   uint16
	blocks  uint32 // 512-byte sector count, including indirect blocks
	flags   uint32
	block   [15]uint32 // direct[0..11], single[12], double[13], triple[14]
	inlined []byte     // fast-symlink target stored in the i_block area
	atime   time.Time
	mtime   time.Time
	ctime   time.Time
}

// encodeTimeExtra returns the 32-bit i_*time_extra value: the nanosecond
// component shifted left by 2 with the two low epoch bits cleared (times after
// the 2038 wrap are not represented).
func encodeTimeExtra(t time.Time) uint32 {
	return uint32(t.Nanosecond()) << 2
}

// encode serializes the inode into a 256-byte buffer.
func (in *inode) encode() []byte {
	b := make([]byte, inodeSize)
	le := binary.LittleEndian

	le.PutUint16(b[0x00:], in.mode)
	le.PutUint16(b[0x02:], uint16(in.uid))
	le.PutUint32(b[0x04:], uint32(in.size))
	le.PutUint32(b[0x08:], uint32(in.atime.Unix()))
	le.PutUint32(b[0x0C:], uint32(in.ctime.Unix()))
	le.PutUint32(b[0x10:], uint32(in.mtime.Unix()))
	// i_dtime 0x14 left zero
	le.PutUint16(b[0x18:], uint16(in.gid))
	le.PutUint16(b[0x1A:], in.links)
	le.PutUint32(b[0x1C:], in.blocks)
	le.PutUint32(b[0x20:], in.flags)

	if in.inlined != nil {
		copy(b[0x28:0x28+60], in.inlined)
	} else {
		for i := 0; i < 15; i++ {
			le.PutUint32(b[0x28+i*4:], in.block[i])
		}
	}

	// i_generation (0x64) and i_file_acl_lo (0x68) are left zero: a freshly
	// built filesystem has no inode-reuse history and no extended attributes.
	// Both are valid zero values that mke2fs also emits for a clean image.

	// i_size_high (regular files use the upper 32 bits via LARGE_FILE).
	le.PutUint32(b[0x6C:], uint32(uint64(in.size)>>32))

	// osd2 high halves of uid/gid.
	le.PutUint16(b[0x78:], uint16(in.uid>>16))
	le.PutUint16(b[0x7A:], uint16(in.gid>>16))

	le.PutUint16(b[0x80:], 32) // i_extra_isize
	le.PutUint32(b[0x84:], encodeTimeExtra(in.ctime))
	le.PutUint32(b[0x88:], encodeTimeExtra(in.mtime))
	le.PutUint32(b[0x8C:], encodeTimeExtra(in.atime))
	// i_crtime / i_crtime_extra: birth time of a freshly materialized inode is
	// its change time. Written because i_extra_isize=32 advertises this region.
	le.PutUint32(b[0x90:], uint32(in.ctime.Unix()))
	le.PutUint32(b[0x94:], encodeTimeExtra(in.ctime))

	return b
}
