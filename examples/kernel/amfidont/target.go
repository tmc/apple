package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	"github.com/tmc/apple/kernel"
)

// amfidPath is the daemon this tool relaxes verdicts inside.
const amfidPath = "/usr/libexec/amfid"

// procAllPIDs is PROC_ALL_PIDS from <sys/proc_info.h>: list every pid.
const procAllPIDs = 1

// findAmfid returns the pid of the running amfid, or an error if it is not
// found. It resolves proc_listpids and proc_pidpath up front, before attach.
func findAmfid() (int, error) {
	var listPIDs func(uint32, uint32, unsafe.Pointer, int32) int32
	var pidPath func(int32, unsafe.Pointer, uint32) int32
	if err := registerFunc(&listPIDs, "proc_listpids"); err != nil {
		return 0, err
	}
	if err := registerFunc(&pidPath, "proc_pidpath"); err != nil {
		return 0, err
	}

	pids := make([]int32, 8192)
	n := listPIDs(procAllPIDs, 0, unsafe.Pointer(&pids[0]), int32(len(pids)*4))
	if n <= 0 {
		return 0, fmt.Errorf("proc_listpids returned %d", n)
	}
	count := int(n) / 4
	buf := make([]byte, 4096)
	for _, pid := range pids[:count] {
		if pid == 0 {
			continue
		}
		m := pidPath(pid, unsafe.Pointer(&buf[0]), uint32(len(buf)))
		if m <= 0 {
			continue
		}
		if string(buf[:m]) == amfidPath {
			return int(pid), nil
		}
	}
	return 0, fmt.Errorf("%s is not running", amfidPath)
}

// translateIMP returns the address of localIMP in a task whose shared-cache
// slide is targetSlide, given that localIMP was resolved in this process under
// ourSlide. The slides must be equal for the result to be a valid breakpoint
// address: the validator IMPs live in the shared cache, and a differing slide
// would put the target's copy elsewhere. Rather than breakpoint a computed
// address inside a system daemon on an unverified slide, a mismatch is an
// error.
func translateIMP(localIMP, ourSlide, targetSlide uint64) (uint64, error) {
	if ourSlide != targetSlide {
		return 0, fmt.Errorf("shared-cache slide mismatch: ours=%#x target=%#x; refusing to breakpoint an unverified address", ourSlide, targetSlide)
	}
	return localIMP - ourSlide + targetSlide, nil
}

// taskDyldInfoFlavor is TASK_DYLD_INFO and taskDyldInfoCount is
// TASK_DYLD_INFO_COUNT from <mach/task_info.h>: task_dyld_info_data_t is
// {all_image_info_addr uint64, all_image_info_size uint64,
// all_image_info_format int32}, five 32-bit words. ABI-pinned, undocumented,
// so written here rather than generated.
const (
	taskDyldInfoFlavor = 17
	taskDyldInfoCount  = 5
)

// dyldSharedCacheSlideOffset is the byte offset of sharedCacheSlide within
// struct dyld_all_image_infos (64-bit) from <mach-o/dyld_images.h>. The layout
// preceding it has been stable across releases; ABI-pinned, undocumented.
const dyldSharedCacheSlideOffset = 152

// selfSharedCacheSlide reads this process's dyld shared-cache slide.
func selfSharedCacheSlide() (uint64, error) {
	return taskSharedCacheSlideByPort(kernel.Mach_task_self())
}

// taskSharedCacheSlide reads pid's dyld shared-cache slide via its task port.
func taskSharedCacheSlide(pid int) (uint64, error) {
	var taskForPID func(uint32, int32, *uint32) int32
	if err := registerFunc(&taskForPID, "task_for_pid"); err != nil {
		return 0, err
	}
	var task uint32
	if kr := taskForPID(kernel.Mach_task_self(), int32(pid), &task); kr != 0 {
		return 0, fmt.Errorf("task_for_pid(%d): kr=%d", pid, kr)
	}
	return taskSharedCacheSlideByPort(task)
}

// taskSharedCacheSlideByPort reads the shared-cache slide from the
// dyld_all_image_infos of the task named by port.
func taskSharedCacheSlideByPort(task uint32) (uint64, error) {
	addr, err := dyldAllImageInfoAddr(task)
	if err != nil {
		return 0, err
	}
	b, err := readTaskMem(task, addr+dyldSharedCacheSlideOffset, 8)
	if err != nil {
		return 0, fmt.Errorf("read sharedCacheSlide: %w", err)
	}
	return binary.LittleEndian.Uint64(b), nil
}

// dyldAllImageInfoAddr returns the all_image_info_addr field of task's
// TASK_DYLD_INFO.
func dyldAllImageInfoAddr(task uint32) (uint64, error) {
	var info [taskDyldInfoCount]int32
	cnt := uint32(taskDyldInfoCount)
	if kr := kernel.Task_info(task, taskDyldInfoFlavor, &info[0], &cnt); kr != 0 {
		return 0, fmt.Errorf("task_info(TASK_DYLD_INFO): kr=%d", kr)
	}
	return uint64(uint32(info[0])) | uint64(uint32(info[1]))<<32, nil
}

// readTaskMem reads n bytes at addr from the task named by port into a fresh
// buffer, using mach_vm_read_overwrite so the copy lands in storage we own.
func readTaskMem(task uint32, addr uint64, n int) ([]byte, error) {
	buf := make([]byte, n)
	var outsize uint64
	kr := kernel.Mach_vm_read_overwrite(task, addr, uint64(n), uint64(uintptr(unsafe.Pointer(&buf[0]))), &outsize)
	if kr != 0 {
		return nil, fmt.Errorf("mach_vm_read_overwrite at %#x: kr=%d", addr, kr)
	}
	return buf[:outsize], nil
}
