//go:build darwin

package hypervisor_test

import (
	"fmt"

	"github.com/tmc/apple/hypervisor"
)

func ExampleHVReturn() {
	fmt.Println(hypervisor.HVSuccess)
	fmt.Println(hypervisor.HVError)
	fmt.Println(hypervisor.HVBadArgument)

	// Output:
	// HVSuccess
	// HVError
	// HVBadArgument
}
