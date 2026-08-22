package signpost

import (
	"bytes"
	"debug/macho"
	"os"
	"sync"
	"unsafe"
)

// The unified logging system records signpost names by reference: logd stores
// an offset into the emitting image's __TEXT,__oslogstring section, and the
// log tools resolve it from the binary on disk at decode time. A name passed
// from the Go heap keeps interval pairing intact but decodes as
// "<missing name>". (Verified empirically 2026-08-11: identical emits with a
// heap pointer, a __TEXT,__rodata pointer, and a __TEXT,__oslogstring pointer
// decode as missing, missing, and the name, respectively.)
//
// Go source cannot place strings in __oslogstring, but a linked-in object
// file (.syso) can. namePool indexes every NUL-terminated string found in the
// running binary's __oslogstring section; emit consults it and passes the
// in-image pointer when the name is present, so those names decode. Names not
// in the pool fall back to the heap copy and keep today's behavior.
//
// To give a binary decodable names, assemble a string pool into it:
//
//	// names.s (assemble with: clang -c names.s -o names.syso)
//	.section __TEXT,__oslogstring,cstring_literals
//	.asciz "GET /work"
//	.asciz "render"
var (
	poolOnce sync.Once
	pool     map[string]*byte
)

func namePool() map[string]*byte {
	poolOnce.Do(func() {
		exe, err := os.Executable()
		if err != nil {
			return
		}
		f, err := macho.Open(exe)
		if err != nil {
			return
		}
		defer f.Close()
		sect := f.Section("__oslogstring")
		if sect == nil {
			return
		}
		data, err := sect.Data()
		if err != nil {
			return
		}
		var textVM uint64
		for _, l := range f.Loads {
			if s, ok := l.(*macho.Segment); ok && s.Name == "__TEXT" {
				textVM = s.Addr
			}
		}
		hdr := dsoHandle()
		if hdr == nil || textVM == 0 {
			return
		}
		// The section's runtime address is the header pointer plus the
		// section's offset from __TEXT's start; deriving it with unsafe.Add
		// keeps the arithmetic anchored on a live pointer.
		base := unsafe.Add(hdr, int(sect.Addr)-int(textVM))

		pool = make(map[string]*byte)
		off := 0
		for off < len(data) {
			end := bytes.IndexByte(data[off:], 0)
			if end < 0 {
				break
			}
			name := string(data[off : off+end])
			if _, ok := pool[name]; !ok {
				pool[name] = (*byte)(unsafe.Add(base, off))
			}
			off += end + 1
		}
	})
	return pool
}
