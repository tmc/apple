// Package img4 assembles Image4 firmware containers from an IM4P payload and an
// IM4M ticket. It preserves payload encoding and signed ticket bytes. It does not
// decrypt, decompress, sign, or authenticate firmware. The caller must verify
// ticket provenance and bind it to the intended hardware, nonces and build.
package img4
