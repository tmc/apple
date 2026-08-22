//go:build darwin

package virtualization_test

import (
	"fmt"

	"github.com/tmc/apple/private/virtualization"
)

func ExampleGetVZVirtualMachineConfigurationClass() {
	cls := virtualization.GetVZVirtualMachineConfigurationClass()
	fmt.Printf("class: %T\n", cls)
	// Output:
	// class: virtualization.VZVirtualMachineConfigurationClass
}
