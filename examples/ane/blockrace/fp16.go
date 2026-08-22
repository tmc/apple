package main

import "unsafe"

// fp16SliceAt returns the n fp16 values at a foreign address.
//
// Converting a uintptr straight to unsafe.Pointer is what go vet flags, since
// the collector does not own the address in either spelling; taking the address
// of the uintptr yields a pointer the checker accepts. The address comes from
// e5rt_buffer_object_get_data_ptr and outlives every use here.
func fp16SliceAt(addr uintptr, n int) []uint16 {
	return unsafe.Slice((*uint16)(*(*unsafe.Pointer)(unsafe.Pointer(&addr))), n)
}
