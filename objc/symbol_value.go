package objc

import "unsafe"

// IDValueAt loads an Objective-C object ID stored at a symbol address.
//
// Dynamic libraries commonly export Objective-C object constants as pointers to
// storage containing the real object ID. Dlsym returns the storage address, so
// callers must load the pointer-sized value stored there.
func IDValueAt(addr uintptr) ID {
	if addr == 0 {
		return 0
	}
	return *(*ID)(unsafe.Pointer(addr))
}
