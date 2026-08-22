//go:build darwin

package kernel_test

import (
	"fmt"

	"github.com/tmc/apple/kernel"
)

func ExampleMachMsgHdr() {
	hdr := kernel.MachMsgHdr{
		Bits:       0,
		Size:       24,
		RemotePort: 0,
		LocalPort:  0,
		ID:         100,
	}
	fmt.Printf("Size: %d, ID: %d, Bits: %d\n", hdr.Size, hdr.ID, hdr.Bits)

	// Output:
	// Size: 24, ID: 100, Bits: 0
}
