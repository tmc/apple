// Package ext4 builds a mountable ext4 filesystem image from a directory tree.
//
// Build writes an ext4 image file populated from an unpacked rootfs directory
// without mounting, loopback devices, or privileges, so it runs on a macOS host
// in an unprivileged build step. The Linux guest kernel later mounts the image
// read-write at /dev/vda; the on-disk format is therefore produced byte-for-byte
// to match what the kernel expects.
//
// The emitted filesystem is a deliberately classic ext4:
//
//   - 4096-byte blocks, 256-byte inodes, revision 1 (dynamic).
//   - Classic indirect block maps (no extents).
//   - feature_incompat = FILETYPE only.
//   - feature_ro_compat = LARGE_FILE only.
//   - No journal, no flex_bg, no sparse_super, no metadata_csum, no htree.
//
// These choices keep the writer small and fully deterministic while remaining
// mountable read-write by any modern Linux kernel, which treats all of the
// omitted features as optional.
//
// Non-goals:
//
//   - No journaling (has_journal) — the image is rebuilt per container.
//   - No extents — classic 12 direct + single/double/triple indirect maps only.
//   - No metadata checksums — s_checksum_type is 0 and metadata_csum is off, so
//     no CRC32c computation is required anywhere.
//   - No sparse_super — every block group carries a full superblock + group
//     descriptor table backup.
//   - No htree directory indexing, no inline data, no extended attributes.
//   - No resize_inode, no 64bit, no meta_bg, no huge_file.
//   - Device nodes, FIFOs, and sockets in the source tree are skipped.
//
// The package depends only on the Go standard library.
package ext4
