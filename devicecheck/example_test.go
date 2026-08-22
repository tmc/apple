//go:build darwin

package devicecheck_test

import (
	"fmt"

	"github.com/tmc/apple/devicecheck"
)

func ExampleDCError() {
	err := devicecheck.DCErrorFeatureUnsupported
	fmt.Println(err)
	// Output:
	// DCErrorFeatureUnsupported
}

func ExampleDCDevice_IsSupported() {
	device := devicecheck.GetDCDeviceClass().CurrentDevice()
	fmt.Printf("DCDevice supported: %t\n", device.IsSupported())
	// Output:
	// DCDevice supported: true
}

func ExampleDCAppAttestService_IsSupported() {
	service := devicecheck.GetDCAppAttestServiceClass().SharedService()
	fmt.Printf("DCAppAttestService supported: %t\n", service.IsSupported())
	// Output:
	// DCAppAttestService supported: false
}
