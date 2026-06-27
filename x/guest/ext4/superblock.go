package ext4

import (
	"encoding/binary"
	"time"
)

const (
	superMagic = 0xEF53

	featureIncompatFiletype  = 0x0002
	featureRoCompatLargeFile = 0x0002

	errorsRO   = 1
	stateOK    = 1
	revDynamic = 1
	firstIno   = 11
)

// superblock holds the parameters needed to serialize the 1024-byte ext4
// superblock. Counts are filled in after layout.
type superblock struct {
	inodesCount     uint32
	blocksCount     uint32
	freeBlocksCount uint32
	freeInodesCount uint32
	firstDataBlock  uint32
	blockSize       uint32
	blocksPerGroup  uint32
	inodesPerGroup  uint32
	inodeSize       uint16
	uuid            [16]byte
	groupCount      uint32
	mkfsTime        time.Time
}

// encode serializes the superblock into a 1024-byte buffer.
func (s *superblock) encode() []byte {
	b := make([]byte, 1024)
	le := binary.LittleEndian

	le.PutUint32(b[0x00:], s.inodesCount)
	le.PutUint32(b[0x04:], s.blocksCount)
	le.PutUint32(b[0x08:], 0) // s_r_blocks_count (reserved blocks)
	le.PutUint32(b[0x0C:], s.freeBlocksCount)
	le.PutUint32(b[0x10:], s.freeInodesCount)
	le.PutUint32(b[0x14:], s.firstDataBlock)
	le.PutUint32(b[0x18:], log2(s.blockSize)-10) // s_log_block_size (2 for 4096)
	le.PutUint32(b[0x1C:], log2(s.blockSize)-10) // s_log_cluster_size
	le.PutUint32(b[0x20:], s.blocksPerGroup)
	le.PutUint32(b[0x24:], s.blocksPerGroup) // s_clusters_per_group
	le.PutUint32(b[0x28:], s.inodesPerGroup)
	le.PutUint32(b[0x2C:], uint32(s.mkfsTime.Unix())) // s_mtime
	le.PutUint32(b[0x30:], uint32(s.mkfsTime.Unix())) // s_wtime
	le.PutUint16(b[0x34:], 0)                         // s_mnt_count
	le.PutUint16(b[0x36:], 0xFFFF)                    // s_max_mnt_count (-1)
	le.PutUint16(b[0x38:], superMagic)
	le.PutUint16(b[0x3A:], stateOK)
	le.PutUint16(b[0x3C:], errorsRO)
	le.PutUint16(b[0x3E:], 0)                         // s_minor_rev_level
	le.PutUint32(b[0x40:], uint32(s.mkfsTime.Unix())) // s_lastcheck
	le.PutUint32(b[0x44:], 0)                         // s_checkinterval
	le.PutUint32(b[0x48:], 0)                         // s_creator_os (Linux)
	le.PutUint32(b[0x4C:], revDynamic)
	le.PutUint16(b[0x50:], 0) // s_def_resuid
	le.PutUint16(b[0x52:], 0) // s_def_resgid

	// EXT4_DYNAMIC_REV fields.
	le.PutUint32(b[0x54:], firstIno)
	le.PutUint16(b[0x58:], s.inodeSize)
	le.PutUint16(b[0x5A:], 0) // s_block_group_nr (primary copy)
	le.PutUint32(b[0x5C:], 0) // s_feature_compat
	le.PutUint32(b[0x60:], featureIncompatFiletype)
	le.PutUint32(b[0x64:], featureRoCompatLargeFile)
	copy(b[0x68:0x78], s.uuid[:]) // s_uuid
	// s_volume_name 0x78 (16 bytes) left empty.

	le.PutUint16(b[0xFE:], 0)  // s_desc_size only meaningful with 64bit; 0 => 32
	le.PutUint32(b[0x100:], 0) // s_default_mount_opts
	le.PutUint32(b[0x104:], 0) // s_first_meta_bg

	// s_checksum_type 0x175 = 0, s_checksum 0x3FC = 0 (left zero).
	return b
}

// log2 returns the base-2 log of a power-of-two value.
func log2(v uint32) uint32 {
	n := uint32(0)
	for v > 1 {
		v >>= 1
		n++
	}
	return n
}
