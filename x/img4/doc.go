// Package img4 assembles Image4 firmware containers and reads IM4M properties.
// Assembly preserves payload encoding and signed ticket bytes. The package does
// not decrypt, decompress, sign, or authenticate firmware. The caller must verify
// ticket provenance and bind it to the intended hardware, nonces and build.
package img4
