// Code generated from Apple documentation for DeviceCheck. DO NOT EDIT.

package devicecheck

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct/Code
type DCError int

const (
	// DCErrorFeatureUnsupported: DeviceCheck is unavailable on this device.
	DCErrorFeatureUnsupported DCError = 1
	// DCErrorInvalidInput: An error code that indicates when your app provides data that isn’t formatted correctly.
	DCErrorInvalidInput DCError = 2
	// DCErrorInvalidKey: An error caused by a failed attempt to use the App Attest key.
	DCErrorInvalidKey DCError = 3
	// DCErrorServerUnavailable: An error that indicates a failed attempt to contact the App Attest service during an attestation.
	DCErrorServerUnavailable DCError = 4
	// DCErrorUnknownSystemFailure: A failure has occurred, such as the failure to generate a token.
	DCErrorUnknownSystemFailure DCError = 0
)

func (e DCError) String() string {
	switch e {
	case DCErrorFeatureUnsupported:
		return "DCErrorFeatureUnsupported"
	case DCErrorInvalidInput:
		return "DCErrorInvalidInput"
	case DCErrorInvalidKey:
		return "DCErrorInvalidKey"
	case DCErrorServerUnavailable:
		return "DCErrorServerUnavailable"
	case DCErrorUnknownSystemFailure:
		return "DCErrorUnknownSystemFailure"
	default:
		return fmt.Sprintf("DCError(%d)", e)
	}
}
