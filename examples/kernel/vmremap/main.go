//go:build darwin

// vmremap demonstrates a shared mapping within the current Mach task.
package main

import (
	"fmt"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/x/mach"
)

func main() {
	const size = 16384
	var sourceAddress uint64
	if kr := kernel.Mach_vm_allocate(kernel.Mach_task_self(), &sourceAddress, size, 1); kr != 0 {
		fmt.Println("vmremap unavailable: allocate:", kr)
		return
	}
	defer func() { _ = kernel.Mach_vm_deallocate(kernel.Mach_task_self(), sourceAddress, size) }()
	backing := mach.Bytes(uintptr(sourceAddress), size)
	backing[0] = 1
	addr, err := mach.Remap(uintptr(sourceAddress), size, mach.RemapOptions{})
	if err != nil {
		fmt.Println("vmremap unavailable:", err)
		return
	}
	defer func() { _ = mach.Unmap(addr, size) }()

	mapped := mach.Bytes(addr, size)
	mapped[0] = 2
	fmt.Printf("shared mapping: source=%d mapping=%d\n", backing[0], mapped[0])
}
