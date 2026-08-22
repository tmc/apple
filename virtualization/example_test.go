//go:build darwin

package virtualization_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/virtualization"
)

func ExampleVZDiskImageCachingMode() {
	fmt.Println(virtualization.VZDiskImageCachingModeAutomatic)
	fmt.Println(virtualization.VZDiskImageCachingModeUncached)
	fmt.Println(virtualization.VZDiskImageCachingModeCached)

	// Output:
	// VZDiskImageCachingModeAutomatic
	// VZDiskImageCachingModeUncached
	// VZDiskImageCachingModeCached
}

func ExampleVZVirtualMachineConfigurationClass() {
	cls := virtualization.GetVZVirtualMachineConfigurationClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)

	// Output:
	// class name: VZVirtualMachineConfiguration
}
