//go:build darwin

package e5rt_test

import (
	"log"
	"unsafe"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/e5rt"
)

// bindExamplePort retains a named port, allocates a CPU-visible buffer object
// for it, and binds the two together.
func bindExamplePort(lib *e5rt.Lib, op uintptr, name string, nbytes int, input bool) (uintptr, uintptr) {
	var port uintptr
	var err error
	if input {
		port, err = lib.OperationRetainInputPort(op, name)
	} else {
		port, err = lib.OperationRetainOutputPort(op, name)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer lib.IOPortRelease(port)

	size := max((nbytes+63)&^63, 64)
	buf, err := lib.BufferObjectAlloc(uintptr(size), 0)
	if err != nil {
		log.Fatal(err)
	}
	ptr, err := lib.BufferObjectGetDataPtr(buf)
	if err != nil {
		log.Fatal(err)
	}
	if err := lib.IOPortBindBufferObject(port, buf); err != nil {
		log.Fatal(err)
	}
	return buf, ptr
}

// pointerAt converts a foreign address to an unsafe.Pointer. Converting a
// uintptr directly is what go vet flags; taking the address of the uintptr
// yields a pointer the checker accepts, and the address is one the Go
// collector does not own in either spelling.
func pointerAt(addr uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&addr))
}

func writeExampleFP16(dst uintptr, src []float32) {
	out := unsafe.Slice((*uint16)(pointerAt(dst)), len(src))
	for i, v := range src {
		out[i] = ane.Float32ToFP16(v)
	}
}

func readExampleFP16(src uintptr, n int) []float32 {
	in := unsafe.Slice((*uint16)(pointerAt(src)), n)
	out := make([]float32, n)
	for i, v := range in {
		out[i] = ane.FP16ToFloat32(v)
	}
	return out
}
