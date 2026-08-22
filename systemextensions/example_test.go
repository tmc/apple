//go:build darwin

package systemextensions_test

import (
	"fmt"

	"github.com/tmc/apple/systemextensions"
)

func ExampleOSSystemExtensionErrorCode() {
	errCode := systemextensions.OSSystemExtensionErrorUnknown
	fmt.Println(errCode)

	authErr := systemextensions.OSSystemExtensionErrorAuthorizationRequired
	fmt.Println(authErr)

	// Output:
	// OSSystemExtensionErrorUnknown
	// OSSystemExtensionErrorAuthorizationRequired
}

func ExampleOSSystemExtensionRequestResult() {
	completed := systemextensions.OSSystemExtensionRequestCompleted
	fmt.Println(completed)

	reboot := systemextensions.OSSystemExtensionRequestWillCompleteAfterReboot
	fmt.Println(reboot)

	// Output:
	// OSSystemExtensionRequestCompleted
	// OSSystemExtensionRequestWillCompleteAfterReboot
}

func ExampleGetOSSystemExtensionManagerClass() {
	class := systemextensions.GetOSSystemExtensionManagerClass()
	if class.Class() != 0 {
		fmt.Println("OSSystemExtensionManager class loaded")
	}

	// Output:
	// OSSystemExtensionManager class loaded
}
