package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

// maxIdentifierLen bounds the bytes read for a signing identifier, so a bogus
// length in a mis-decoded object cannot drive an enormous read.
const maxIdentifierLen = 4096

// CFString info-byte flags, read off macOS 27 by dumping __NSCFString objects
// created from C strings. The info byte sits at offset ptrSize (8) inside the
// object. ABI-pinned to the CoreFoundation implementation, undocumented.
const (
	cfStrMutable       = 1 << 0 // mutable storage
	cfStrHasLengthByte = 1 << 2 // inline data is prefixed by a one-byte length
	cfStrUnicode       = 1 << 4 // contents are UTF-16, not 8-bit
	cfStrStorageMask   = 3 << 5 // 0 means inline contents, non-zero means an external pointer
)

// readSigningIdentifier reads the NSString* at ptr out of the address space r
// and returns its bytes as a Go string. It decodes only the shape a signing
// identifier actually takes — an immutable, non-unicode CFString, inline or
// external — and fails closed on anything else (tagged pointers, mutable or
// unicode storage). Reverse-DNS signing identifiers are ASCII and longer than
// the tagged-pointer threshold, so this covers the real case; refusing the
// rest is the safe direction, because a wrong identity read must never lead to
// a patch.
//
// r is the target's memory (machdebug.Process implements io.ReaderAt); in
// tests it is this process's own memory, so the decoder is exercised against
// real CFString objects with known contents.
func readSigningIdentifier(r io.ReaderAt, ptr uint64) (string, error) {
	if ptr == 0 {
		return "", fmt.Errorf("signing identifier pointer is null")
	}
	if ptr&(1<<63) != 0 {
		return "", fmt.Errorf("signing identifier is a tagged-pointer string (%#x); unsupported", ptr)
	}

	// Object header: isa (8), cfinfo word (8), then either an inline length
	// byte and data, or a length word, or an external pointer and length word.
	var hdr [32]byte
	if _, err := r.ReadAt(hdr[:], int64(ptr)); err != nil {
		return "", fmt.Errorf("read NSString header at %#x: %w", ptr, err)
	}
	info := hdr[8]
	if info&cfStrMutable != 0 {
		return "", fmt.Errorf("signing identifier is a mutable string; unsupported")
	}
	if info&cfStrUnicode != 0 {
		return "", fmt.Errorf("signing identifier is a unicode string; unsupported")
	}

	if info&cfStrStorageMask == 0 {
		// Inline contents.
		if info&cfStrHasLengthByte != 0 {
			n := int(hdr[16])
			return readAt(r, ptr+17, n)
		}
		n := int(binary.LittleEndian.Uint64(hdr[16:24]))
		return readAt(r, ptr+24, n)
	}

	// External contents: {isa, cfinfo, char* contents, length}.
	dataPtr := binary.LittleEndian.Uint64(hdr[16:24])
	n := int(binary.LittleEndian.Uint64(hdr[24:32]))
	if dataPtr == 0 {
		return "", fmt.Errorf("external string has null contents pointer")
	}
	return readAt(r, dataPtr, n)
}

// readAt reads exactly n bytes at addr from r, bounding n first.
func readAt(r io.ReaderAt, addr uint64, n int) (string, error) {
	if n < 0 || n > maxIdentifierLen {
		return "", fmt.Errorf("implausible string length %d", n)
	}
	b := make([]byte, n)
	if _, err := r.ReadAt(b, int64(addr)); err != nil {
		return "", fmt.Errorf("read string contents at %#x: %w", addr, err)
	}
	return string(b), nil
}
