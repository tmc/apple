//go:build darwin

package localauthentication_test

import (
	"fmt"

	"github.com/tmc/apple/localauthentication"
)

func ExampleLAContext() {
	ctx := localauthentication.NewLAContext()
	ctx.SetLocalizedReason("Authenticate to access secure data")
	fmt.Println(ctx.LocalizedReason())
	ctx.Invalidate()

	// Output:
	// Authenticate to access secure data
}

func ExampleLAPolicy() {
	policy := localauthentication.LAPolicyDeviceOwnerAuthenticationWithBiometrics
	fmt.Println(policy)

	// Output:
	// LAPolicyDeviceOwnerAuthenticationWithBiometrics
}

func ExampleLABiometryType() {
	t := localauthentication.LABiometryTypeTouchID
	fmt.Println(t)

	// Output:
	// LABiometryTypeTouchID
}

func ExampleLAError() {
	err := localauthentication.LAErrorAuthenticationFailed
	fmt.Println(err)

	// Output:
	// LAErrorAuthenticationFailed
}
