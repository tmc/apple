//go:build darwin

package vmnet_test

import (
	"fmt"

	"github.com/tmc/apple/vmnet"
)

func ExampleOperatingModes() {
	fmt.Println(vmnet.VMNET_HOST_MODE)
	fmt.Println(vmnet.VMNET_SHARED_MODE)
	fmt.Println(vmnet.VMNET_BRIDGED_MODE)

	// Output:
	// VMNET_HOST_MODE
	// VMNET_SHARED_MODE
	// VMNET_BRIDGED_MODE
}

func ExampleVmnetReturn() {
	fmt.Println(vmnet.VMNET_SUCCESS)
	fmt.Println(vmnet.VMNET_FAILURE)

	// Output:
	// VMNET_SUCCESS
	// VMNET_FAILURE
}
