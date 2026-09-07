package main

import (
	"io"
	"runtime"
	"strings"
	"testing"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/kernel"
)

// selfMem is this process's own address space as an io.ReaderAt, so the
// out-of-process NSString decoder can be exercised against real CFString
// objects with known contents. It reads through mach_vm_read_overwrite on our
// own task — the same primitive the tool uses against amfid — which both
// avoids reconstructing a pointer from an integer address and returns an error
// on an unmapped address instead of faulting.
type selfMem struct{}

func (selfMem) ReadAt(b []byte, off int64) (int, error) {
	if off <= 0 {
		return 0, io.EOF
	}
	got, err := readTaskMem(kernel.Mach_task_self(), uint64(off), len(b))
	if err != nil {
		return 0, err
	}
	copy(b, got)
	return len(got), nil
}

// makeNSString builds an NSString in this process and returns its pointer.
func makeNSString(t *testing.T, s string) uint64 {
	t.Helper()
	if _, err := purego.Dlopen("/System/Library/Frameworks/Foundation.framework/Foundation", purego.RTLD_NOW|purego.RTLD_GLOBAL); err != nil {
		t.Skipf("dlopen Foundation: %v", err)
	}
	var getClass func(string) uintptr
	var selReg func(string) uintptr
	var msgSend func(uintptr, uintptr, uintptr) uintptr
	purego.RegisterLibFunc(&getClass, purego.RTLD_DEFAULT, "objc_getClass")
	purego.RegisterLibFunc(&selReg, purego.RTLD_DEFAULT, "sel_registerName")
	purego.RegisterLibFunc(&msgSend, purego.RTLD_DEFAULT, "objc_msgSend")
	cls := getClass("NSString")
	sel := selReg("stringWithUTF8String:")
	cs := append([]byte(s), 0)
	p := msgSend(cls, sel, uintptr(unsafe.Pointer(&cs[0])))
	if p == 0 {
		t.Fatalf("stringWithUTF8String:(%q) returned nil", s)
	}
	// Keep the source bytes alive until the call has copied them.
	runtime.KeepAlive(cs)
	return uint64(p)
}

func TestReadSigningIdentifier(t *testing.T) {
	// Realistic reverse-DNS identifiers of varying length exercise both the
	// short-length-byte inline form and the length-word inline form; all are
	// immutable ASCII, the shape a real signing identifier takes.
	cases := []string{
		"com.apple.covevm",
		"com.apple.virtualization",
		"a.b.c.d.e.f.g.h",
		strings.Repeat("com.apple.x.", 40) + "end",
	}
	for _, want := range cases {
		ptr := makeNSString(t, want)
		if ptr&(1<<63) != 0 {
			// A tagged pointer is deliberately unsupported; the decoder must
			// reject it rather than guess.
			if _, err := readSigningIdentifier(selfMem{}, ptr); err == nil {
				t.Errorf("readSigningIdentifier(tagged %q) = nil error, want rejection", want)
			}
			continue
		}
		got, err := readSigningIdentifier(selfMem{}, ptr)
		if err != nil {
			t.Errorf("readSigningIdentifier(%q): %v", want, err)
			continue
		}
		if got != want {
			t.Errorf("readSigningIdentifier(%q) = %q", want, got)
		}
	}
}

func TestReadSigningIdentifierFailsClosed(t *testing.T) {
	if _, err := readSigningIdentifier(selfMem{}, 0); err == nil {
		t.Error("null pointer: want error, got nil")
	}
	if _, err := readSigningIdentifier(selfMem{}, 1<<63|0x1000); err == nil {
		t.Error("tagged pointer: want error, got nil")
	}
}
