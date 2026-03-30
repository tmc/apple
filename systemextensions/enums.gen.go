// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code
type OSSystemExtensionErrorCode int

const (
	// OSSystemExtensionErrorAuthorizationRequired: An error code that indicates the system was unable to obtain the proper authorization.
	OSSystemExtensionErrorAuthorizationRequired OSSystemExtensionErrorCode = 13
	// OSSystemExtensionErrorCodeSignatureInvalid: An error code that indicates the extension’s signature is invalid.
	OSSystemExtensionErrorCodeSignatureInvalid OSSystemExtensionErrorCode = 8
	// OSSystemExtensionErrorDuplicateExtensionIdentifer: An error code that indicates the extension identifier duplicates an existing identifier.
	OSSystemExtensionErrorDuplicateExtensionIdentifer OSSystemExtensionErrorCode = 6
	// OSSystemExtensionErrorExtensionMissingIdentifier: An error code that indicates the extension identifier is missing.
	OSSystemExtensionErrorExtensionMissingIdentifier OSSystemExtensionErrorCode = 5
	// OSSystemExtensionErrorExtensionNotFound: An error code that indicates the manager can’t find the system extension.
	OSSystemExtensionErrorExtensionNotFound OSSystemExtensionErrorCode = 4
	// OSSystemExtensionErrorForbiddenBySystemPolicy: An error code that indicates the system policy prohibits activating the system extension.
	OSSystemExtensionErrorForbiddenBySystemPolicy OSSystemExtensionErrorCode = 10
	// OSSystemExtensionErrorMissingEntitlement: An error code that indicates the system extension lacks a required entitlement.
	OSSystemExtensionErrorMissingEntitlement OSSystemExtensionErrorCode = 2
	// OSSystemExtensionErrorRequestCanceled: An error code that indicates the system extension manager request was canceled.
	OSSystemExtensionErrorRequestCanceled OSSystemExtensionErrorCode = 11
	// OSSystemExtensionErrorRequestSuperseded: An error code that indicates the system extension request failed because the system already has a pending request for the same identifier.
	OSSystemExtensionErrorRequestSuperseded OSSystemExtensionErrorCode = 12
	// OSSystemExtensionErrorUnknown: An error code that indicates an unknown error occurred.
	OSSystemExtensionErrorUnknown OSSystemExtensionErrorCode = 1
	// OSSystemExtensionErrorUnknownExtensionCategory: An error code that indicates the extension manager can’t recognize the extension’s category identifier.
	OSSystemExtensionErrorUnknownExtensionCategory OSSystemExtensionErrorCode = 7
	// OSSystemExtensionErrorUnsupportedParentBundleLocation: An error code that indicates the extension’s parent app isn’t in a valid location for activation.
	OSSystemExtensionErrorUnsupportedParentBundleLocation OSSystemExtensionErrorCode = 3
	// OSSystemExtensionErrorValidationFailed: An error code that indicates the manager can’t validate the extension.
	OSSystemExtensionErrorValidationFailed OSSystemExtensionErrorCode = 9
)

func (e OSSystemExtensionErrorCode) String() string {
	switch e {
	case OSSystemExtensionErrorAuthorizationRequired:
		return "OSSystemExtensionErrorAuthorizationRequired"
	case OSSystemExtensionErrorCodeSignatureInvalid:
		return "OSSystemExtensionErrorCodeSignatureInvalid"
	case OSSystemExtensionErrorDuplicateExtensionIdentifer:
		return "OSSystemExtensionErrorDuplicateExtensionIdentifer"
	case OSSystemExtensionErrorExtensionMissingIdentifier:
		return "OSSystemExtensionErrorExtensionMissingIdentifier"
	case OSSystemExtensionErrorExtensionNotFound:
		return "OSSystemExtensionErrorExtensionNotFound"
	case OSSystemExtensionErrorForbiddenBySystemPolicy:
		return "OSSystemExtensionErrorForbiddenBySystemPolicy"
	case OSSystemExtensionErrorMissingEntitlement:
		return "OSSystemExtensionErrorMissingEntitlement"
	case OSSystemExtensionErrorRequestCanceled:
		return "OSSystemExtensionErrorRequestCanceled"
	case OSSystemExtensionErrorRequestSuperseded:
		return "OSSystemExtensionErrorRequestSuperseded"
	case OSSystemExtensionErrorUnknown:
		return "OSSystemExtensionErrorUnknown"
	case OSSystemExtensionErrorUnknownExtensionCategory:
		return "OSSystemExtensionErrorUnknownExtensionCategory"
	case OSSystemExtensionErrorUnsupportedParentBundleLocation:
		return "OSSystemExtensionErrorUnsupportedParentBundleLocation"
	case OSSystemExtensionErrorValidationFailed:
		return "OSSystemExtensionErrorValidationFailed"
	default:
		return fmt.Sprintf("OSSystemExtensionErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/ReplacementAction
type OSSystemExtensionReplacementAction int

const (
	// OSSystemExtensionReplacementActionCancel: An action that tells the manager to cancel replacement of a system extension.
	OSSystemExtensionReplacementActionCancel OSSystemExtensionReplacementAction = 0
	// OSSystemExtensionReplacementActionReplace: An action that tells the manager to replace an existing system extension.
	OSSystemExtensionReplacementActionReplace OSSystemExtensionReplacementAction = 1
)

func (e OSSystemExtensionReplacementAction) String() string {
	switch e {
	case OSSystemExtensionReplacementActionCancel:
		return "OSSystemExtensionReplacementActionCancel"
	case OSSystemExtensionReplacementActionReplace:
		return "OSSystemExtensionReplacementActionReplace"
	default:
		return fmt.Sprintf("OSSystemExtensionReplacementAction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/Result
type OSSystemExtensionRequestResult int

const (
	// OSSystemExtensionRequestCompleted: The request completed successfully.
	OSSystemExtensionRequestCompleted OSSystemExtensionRequestResult = 0
	// OSSystemExtensionRequestWillCompleteAfterReboot: The request requires a restart to complete successfully.
	OSSystemExtensionRequestWillCompleteAfterReboot OSSystemExtensionRequestResult = 1
)

func (e OSSystemExtensionRequestResult) String() string {
	switch e {
	case OSSystemExtensionRequestCompleted:
		return "OSSystemExtensionRequestCompleted"
	case OSSystemExtensionRequestWillCompleteAfterReboot:
		return "OSSystemExtensionRequestWillCompleteAfterReboot"
	default:
		return fmt.Sprintf("OSSystemExtensionRequestResult(%d)", e)
	}
}
