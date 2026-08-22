//go:build darwin

package mach

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/tmc/apple/kernel"
)

// RemapOptions controls a mapping made by Remap.
type RemapOptions struct {
	// Copy requests copy-on-write mappings when true. When false, the source
	// and destination share physical pages.
	Copy bool
	// Protection is the initial and maximum VM protection. Zero means read,
	// write, and execute, matching the small demonstration use case.
	Protection kernel.VmProt
}

// Remap maps size bytes from src into the current task and returns the mapped
// address. The returned address is outside the Go heap. Call Unmap exactly
// once when the byte slice is no longer in use.
func Remap(src uintptr, size uintptr, opts RemapOptions) (uintptr, error) {
	if src == 0 {
		return 0, fmt.Errorf("mach: remap nil source")
	}
	if size == 0 {
		return 0, fmt.Errorf("mach: remap zero size")
	}
	page := uintptr(os.Getpagesize())
	if src%page != 0 || size%page != 0 {
		return 0, fmt.Errorf("mach: remap address and size must be page aligned")
	}
	prot := opts.Protection
	if prot == 0 {
		prot = 7
	}
	copyPages := kernel.Boolean_t(0)
	if opts.Copy {
		copyPages = 1
	}
	var dst uint64
	kr := kernel.Mach_vm_remap(kernel.Mach_task_self(), &dst, uint64(size), 0, 1, kernel.Mach_task_self(), uint64(src), copyPages, &prot, &prot, 2)
	if err := kernError("mach_vm_remap", kr); err != nil {
		return 0, err
	}
	if dst == 0 {
		return 0, fmt.Errorf("mach: remap returned nil address")
	}
	return uintptr(dst), nil
}

// Bytes returns a byte view of a mapping returned by Remap. The caller must
// keep the mapping alive until all uses of the returned slice have finished.
func Bytes(addr uintptr, size uintptr) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(addr)), size)
}

// Unmap releases a mapping returned by Remap.
func Unmap(addr uintptr, size uintptr) error {
	if addr == 0 || size == 0 {
		return fmt.Errorf("mach: unmap invalid mapping")
	}
	return kernError("mach_vm_deallocate", kernel.Mach_vm_deallocate(kernel.Mach_task_self(), uint64(addr), uint64(size)))
}
