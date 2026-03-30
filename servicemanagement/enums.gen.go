// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

package servicemanagement

import (
	"fmt"
)

type KSMError int

const (
	// KSMErrorAlreadyRegistered: The application is already registered.
	KSMErrorAlreadyRegistered KSMError = 12
	// KSMErrorAuthorizationFailure: The authorization requested failed.
	KSMErrorAuthorizationFailure KSMError = 4
	// KSMErrorInternalFailure: An internal failure has occurred.
	KSMErrorInternalFailure KSMError = 2
	// KSMErrorInvalidPlist: The app’s property list is invalid.
	KSMErrorInvalidPlist KSMError = 10
	// KSMErrorInvalidSignature: The app’s code signature doesn’t meet the requirements to perform the operation.
	KSMErrorInvalidSignature KSMError = 3
	KSMErrorJobMustBeEnabled KSMError = 9
	// KSMErrorJobNotFound: The system can’t find the specified job.
	KSMErrorJobNotFound      KSMError = 6
	KSMErrorJobPlistNotFound KSMError = 8
	// KSMErrorLaunchDeniedByUser: The user denied the app’s launch request.
	KSMErrorLaunchDeniedByUser KSMError = 11
	// KSMErrorServiceUnavailable: The service necessary to perform this operation is unavailable or is no longer accepting requests.
	KSMErrorServiceUnavailable KSMError = 7
	// KSMErrorToolNotValid: The specified path doesn’t exist or the helper tool at the specified path isn’t valid.
	KSMErrorToolNotValid KSMError = 5
)

func (e KSMError) String() string {
	switch e {
	case KSMErrorAlreadyRegistered:
		return "KSMErrorAlreadyRegistered"
	case KSMErrorAuthorizationFailure:
		return "KSMErrorAuthorizationFailure"
	case KSMErrorInternalFailure:
		return "KSMErrorInternalFailure"
	case KSMErrorInvalidPlist:
		return "KSMErrorInvalidPlist"
	case KSMErrorInvalidSignature:
		return "KSMErrorInvalidSignature"
	case KSMErrorJobMustBeEnabled:
		return "KSMErrorJobMustBeEnabled"
	case KSMErrorJobNotFound:
		return "KSMErrorJobNotFound"
	case KSMErrorJobPlistNotFound:
		return "KSMErrorJobPlistNotFound"
	case KSMErrorLaunchDeniedByUser:
		return "KSMErrorLaunchDeniedByUser"
	case KSMErrorServiceUnavailable:
		return "KSMErrorServiceUnavailable"
	case KSMErrorToolNotValid:
		return "KSMErrorToolNotValid"
	default:
		return fmt.Sprintf("KSMError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum
type SMAppServiceStatus int

const (
	// SMAppServiceStatusEnabled: The service has been successfully registered and is eligible to run.
	SMAppServiceStatusEnabled SMAppServiceStatus = 1
	// SMAppServiceStatusNotFound: An error occurred and the framework couldn’t find this service.
	SMAppServiceStatusNotFound SMAppServiceStatus = 3
	// SMAppServiceStatusNotRegistered: The service hasn’t registered with the Service Management framework, or the service attempted to reregister after it was already registered.
	SMAppServiceStatusNotRegistered SMAppServiceStatus = 0
	// SMAppServiceStatusRequiresApproval: The service has been successfully registered, but the user needs to take action in System Preferences.
	SMAppServiceStatusRequiresApproval SMAppServiceStatus = 2
)

func (e SMAppServiceStatus) String() string {
	switch e {
	case SMAppServiceStatusEnabled:
		return "SMAppServiceStatusEnabled"
	case SMAppServiceStatusNotFound:
		return "SMAppServiceStatusNotFound"
	case SMAppServiceStatusNotRegistered:
		return "SMAppServiceStatusNotRegistered"
	case SMAppServiceStatusRequiresApproval:
		return "SMAppServiceStatusRequiresApproval"
	default:
		return fmt.Sprintf("SMAppServiceStatus(%d)", e)
	}
}
