package machdebug

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/kernel"
)

// pageSize is the target's page size, which is this process's page size:
// both are Apple silicon user tasks.
var pageSize = uint64(os.Getpagesize())

// ReadAt reads from the target's address space at off, implementing
// [io.ReaderAt]. It is what lets debug/macho, encoding/binary and
// io.SectionReader work against a live process with no new API.
//
// A short read reports [io.EOF], the [io.ReaderAt] contract for an
// unreadable tail; the common cause is an unmapped page, not an end of file.
func (p *Process) ReadAt(b []byte, off int64) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	if off < 0 {
		return 0, fmt.Errorf("machdebug: read at %#x: negative address", off)
	}
	if err := p.alive(); err != nil {
		return 0, err
	}

	// mach_vm_read_overwrite stops at the first unmapped page rather than
	// returning a partial count, so read page by page and report how far we
	// got. Callers that hand us one small buffer take a single round trip.
	var n int
	for n < len(b) {
		want := len(b) - n
		if end := pageEnd(uint64(off) + uint64(n)); uint64(want) > end-(uint64(off)+uint64(n)) {
			want = int(end - (uint64(off) + uint64(n)))
		}
		var pin runtime.Pinner
		pin.Pin(&b[n])
		var got kernel.Mach_vm_size_t
		kr := kernel.Mach_vm_read_overwrite(
			kernel.Vm_map_read_t(p.task),
			kernel.Mach_vm_address_t(uint64(off)+uint64(n)),
			kernel.Mach_vm_size_t(want),
			kernel.Mach_vm_address_t(uintptr(unsafe.Pointer(&b[n]))),
			&got)
		pin.Unpin()
		if kr != kernSuccess {
			if n == 0 {
				return 0, fmt.Errorf("machdebug: read %d bytes at %#x: %w", len(b), off, kernErr(kr))
			}
			return n, io.EOF
		}
		n += int(got)
		if got == 0 {
			return n, io.EOF
		}
	}
	return n, nil
}

// WriteAt writes to the target's address space at off, implementing
// [io.WriterAt]. Writes to read-only pages, including shared-cache text, go
// through copy-on-write so the change is private to the target task.
func (p *Process) WriteAt(b []byte, off int64) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	if off < 0 {
		return 0, fmt.Errorf("machdebug: write at %#x: negative address", off)
	}
	if err := p.alive(); err != nil {
		return 0, err
	}
	if kr := p.rawWrite(uint64(off), b); kr == kernSuccess {
		return len(b), nil
	} else if kr != kernProtectionFailure {
		return 0, fmt.Errorf("machdebug: write %d bytes at %#x: %w", len(b), off, kernErr(kr))
	}
	if err := p.writeProtected(uint64(off), b); err != nil {
		return 0, err
	}
	return len(b), nil
}

// rawWrite is a bare mach_vm_write, returning the kern_return_t.
func (p *Process) rawWrite(addr uint64, b []byte) kernel.Kern_return_t {
	var pin runtime.Pinner
	pin.Pin(&b[0])
	kr := kernel.Mach_vm_write(p.task,
		kernel.Mach_vm_address_t(addr),
		kernel.Vm_offset_t(uintptr(unsafe.Pointer(&b[0]))),
		kernel.Mach_msg_type_number_t(len(b)))
	pin.Unpin()
	return kr
}

// writeProtected writes b at addr over pages that are not currently
// writable: it raises protection with VM_PROT_COPY, writes, and puts the
// original protection back.
//
// VM_PROT_COPY is the whole trick and is not optional. Text is mapped
// read-only and shared with every other process that maps the same file (or
// the dyld shared cache). VM_PROT_COPY asks the kernel for a private
// copy-on-write copy of the range first, so the patch lands only in this
// task; without it the protect either fails outright with
// KERN_PROTECTION_FAILURE or, on a range the kernel is willing to make
// writable in place, corrupts every process sharing the page.
//
// The task is suspended for the duration. Raising a text page to read-write
// takes execute away from it, and any other thread that fetches an
// instruction from that page in the window between the two protect calls
// takes an unrecoverable EXC_BAD_ACCESS — which is a crash in the target
// caused entirely by the debugger, and a flaky one, since it depends on where
// the other threads happen to be. Suspending is the only way to close the
// window.
func (p *Process) writeProtected(addr uint64, b []byte) error {
	start := addr &^ (pageSize - 1)
	end := pageEnd(addr + uint64(len(b)) - 1)
	size := end - start

	old, err := p.protection(start)
	if err != nil {
		return err
	}
	if kr := kernel.Task_suspend(kernel.Task_read_t(p.task)); kr != kernSuccess {
		return fmt.Errorf("machdebug: task_suspend: %w", kernErr(kr))
	}
	defer kernel.Task_resume(kernel.Task_read_t(p.task))
	kr := kernel.Mach_vm_protect(p.task, kernel.Mach_vm_address_t(start), kernel.Mach_vm_size_t(size),
		0, kernel.Vm_prot_t(vmProtRead|vmProtWrite|vmProtCopy))
	if kr != kernSuccess {
		return fmt.Errorf("machdebug: protect %#x+%#x rw|copy: %w", start, size, kernErr(kr))
	}
	werr := p.rawWrite(addr, b)
	if werr == kernSuccess {
		p.flushICache(start, size)
	}
	// Put the original protection back whether or not the write worked.
	if kr := kernel.Mach_vm_protect(p.task, kernel.Mach_vm_address_t(start), kernel.Mach_vm_size_t(size),
		0, kernel.Vm_prot_t(old)); kr != kernSuccess && werr == kernSuccess {
		return fmt.Errorf("machdebug: restore protection %#x+%#x to %#x: %w", start, size, old, kernErr(kr))
	}
	if werr != kernSuccess {
		return fmt.Errorf("machdebug: write %d bytes at %#x: %w", len(b), addr, kernErr(werr))
	}
	return nil
}

// flushICache invalidates the target's instruction caches over a range.
//
// Patching text is a data write, and the cores are free to keep executing
// what their instruction caches already hold. Without this the target can
// re-take a breakpoint we already removed, whereupon we decline an exception
// nobody set and the trap kills it.
func (p *Process) flushICache(addr, size uint64) {
	val := kernel.Vm_machine_attribute_val_t(mattrValICacheFlush)
	kernel.Mach_vm_machine_attribute(p.task, kernel.Mach_vm_address_t(addr),
		kernel.Mach_vm_size_t(size), kernel.Vm_machine_attribute_t(mattrCache), &val)
}

// protection reports the current protection of the region containing addr.
func (p *Process) protection(addr uint64) (int32, error) {
	// vm_region_basic_info_data_64_t: protection is the first word.
	var info [vmRegionBasicInfo64Count]int32
	a := kernel.Mach_vm_address_t(addr)
	var size kernel.Mach_vm_size_t
	cnt := kernel.Mach_msg_type_number_t(vmRegionBasicInfo64Count)
	var objName uint32
	kr := kernel.Mach_vm_region(kernel.Vm_map_read_t(p.task), &a, &size,
		kernel.Vm_region_flavor_t(vmRegionBasicInfo64),
		kernel.Vm_region_info_t(&info[0]), &cnt, &objName)
	if kr != kernSuccess {
		return 0, fmt.Errorf("machdebug: region at %#x: %w", addr, kernErr(kr))
	}
	// mach_vm_region hands back a send right to the region's object name;
	// dropping it silently leaks a port every time we patch.
	if objName != 0 {
		kernel.Mach_port_deallocate(kernel.Mach_task_self(), kernel.Mach_port_name_t(objName))
	}
	return info[0], nil
}

// pageEnd returns the first address after addr's page.
func pageEnd(addr uint64) uint64 { return (addr &^ (pageSize - 1)) + pageSize }
