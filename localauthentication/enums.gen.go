// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/LocalAuthentication/LAAccessControlOperation
type LAAccessControlOperation int

const (
	// LAAccessControlOperationCreateItem: Specifies that access control is used for item creation.
	LAAccessControlOperationCreateItem LAAccessControlOperation = 0
	// LAAccessControlOperationCreateKey: Specifies that access control is used for key creation.
	LAAccessControlOperationCreateKey LAAccessControlOperation = 2
	// LAAccessControlOperationUseItem: Specifies that access control is used for accessing an existing item.
	LAAccessControlOperationUseItem LAAccessControlOperation = 1
	// LAAccessControlOperationUseKeyDecrypt: Specifies that access control is used for data decryption using existing key.
	LAAccessControlOperationUseKeyDecrypt LAAccessControlOperation = 4
	// LAAccessControlOperationUseKeyKeyExchange: Specifies that access control is used for key exchange.
	LAAccessControlOperationUseKeyKeyExchange LAAccessControlOperation = 5
	// LAAccessControlOperationUseKeySign: Specifies that access control is used for accessing an existing key.
	LAAccessControlOperationUseKeySign LAAccessControlOperation = 3
)

func (e LAAccessControlOperation) String() string {
	switch e {
	case LAAccessControlOperationCreateItem:
		return "LAAccessControlOperationCreateItem"
	case LAAccessControlOperationCreateKey:
		return "LAAccessControlOperationCreateKey"
	case LAAccessControlOperationUseItem:
		return "LAAccessControlOperationUseItem"
	case LAAccessControlOperationUseKeyDecrypt:
		return "LAAccessControlOperationUseKeyDecrypt"
	case LAAccessControlOperationUseKeyKeyExchange:
		return "LAAccessControlOperationUseKeyKeyExchange"
	case LAAccessControlOperationUseKeySign:
		return "LAAccessControlOperationUseKeySign"
	default:
		return fmt.Sprintf("LAAccessControlOperation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/LocalAuthentication/LABiometryType
type LABiometryType int

const (
	// LABiometryTypeFaceID: The device supports Face ID.
	LABiometryTypeFaceID LABiometryType = 2
	// LABiometryTypeNone: No biometry type is supported.
	LABiometryTypeNone LABiometryType = 0
	// LABiometryTypeOpticID: The device supports Optic ID.
	LABiometryTypeOpticID LABiometryType = 4
	// LABiometryTypeTouchID: The device supports Touch ID.
	LABiometryTypeTouchID LABiometryType = 1
	// Deprecated.
	LABiometryNone LABiometryType = 0
)

func (e LABiometryType) String() string {
	switch e {
	case LABiometryTypeFaceID:
		return "LABiometryTypeFaceID"
	case LABiometryTypeNone:
		return "LABiometryTypeNone"
	case LABiometryTypeOpticID:
		return "LABiometryTypeOpticID"
	case LABiometryTypeTouchID:
		return "LABiometryTypeTouchID"
	default:
		return fmt.Sprintf("LABiometryType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/LocalAuthentication/LACompanionType
type LACompanionType int

const (
	// LACompanionTypeMac: Paired Mac
	LACompanionTypeMac LACompanionType = 2
	// LACompanionTypeVision: Paired Vision Pro
	LACompanionTypeVision LACompanionType = 4
	// LACompanionTypeWatch: Paired Apple Watch
	LACompanionTypeWatch LACompanionType = 1
)

func (e LACompanionType) String() string {
	switch e {
	case LACompanionTypeMac:
		return "LACompanionTypeMac"
	case LACompanionTypeVision:
		return "LACompanionTypeVision"
	case LACompanionTypeWatch:
		return "LACompanionTypeWatch"
	default:
		return fmt.Sprintf("LACompanionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/LocalAuthentication/LACredentialType
type LACredentialType int

const (
	// LACredentialTypeApplicationPassword: Specifies that a password is provided by the application.
	LACredentialTypeApplicationPassword LACredentialType = 0
	LACredentialTypeSmartCardPIN        LACredentialType = -3
)

func (e LACredentialType) String() string {
	switch e {
	case LACredentialTypeApplicationPassword:
		return "LACredentialTypeApplicationPassword"
	case LACredentialTypeSmartCardPIN:
		return "LACredentialTypeSmartCardPIN"
	default:
		return fmt.Sprintf("LACredentialType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code
type LAError int

const (
	// LAErrorAppCancel: The app canceled authentication.
	LAErrorAppCancel LAError = -9
	// LAErrorAuthenticationFailed: The user failed to provide valid credentials.
	LAErrorAuthenticationFailed LAError = -1
	// LAErrorBiometryDisconnected: The device supports biometry only using a removable accessory, but the paired accessory isn’t connected.
	LAErrorBiometryDisconnected LAError = -13
	// LAErrorBiometryLockout: Biometry is locked because there were too many failed attempts.
	LAErrorBiometryLockout LAError = -8
	// LAErrorBiometryNotAvailable: Biometry is not available on the device.
	LAErrorBiometryNotAvailable LAError = -6
	// LAErrorBiometryNotEnrolled: The user has no enrolled biometric identities.
	LAErrorBiometryNotEnrolled LAError = -7
	// LAErrorBiometryNotPaired: The device supports biometry only using a removable accessory, but no accessory is paired.
	LAErrorBiometryNotPaired LAError = -12
	// LAErrorCompanionNotAvailable: Authentication could not start because there was no paired companion device nearby.
	LAErrorCompanionNotAvailable LAError = -11
	// LAErrorInvalidContext: The context was previously invalidated.
	LAErrorInvalidContext    LAError = -10
	LAErrorInvalidDimensions LAError = -12
	// LAErrorNotInteractive: Displaying the required authentication user interface is forbidden.
	LAErrorNotInteractive LAError = -1004
	// LAErrorPasscodeNotSet: A passcode isn’t set on the device.
	LAErrorPasscodeNotSet LAError = -5
	// LAErrorSystemCancel: The system canceled authentication.
	LAErrorSystemCancel LAError = -4
	// LAErrorUserCancel: The user tapped the cancel button in the authentication dialog.
	LAErrorUserCancel LAError = -2
	// LAErrorUserFallback: The user tapped the fallback button in the authentication dialog, but no fallback is available for the authentication policy.
	LAErrorUserFallback LAError = -3
	// Deprecated.
	LAErrorTouchIDLockout LAError = -6
	// Deprecated.
	LAErrorTouchIDNotAvailable LAError = -6
	// Deprecated.
	LAErrorTouchIDNotEnrolled LAError = -7
	// Deprecated.
	LAErrorWatchNotAvailable LAError = -11
)

func (e LAError) String() string {
	switch e {
	case LAErrorAppCancel:
		return "LAErrorAppCancel"
	case LAErrorAuthenticationFailed:
		return "LAErrorAuthenticationFailed"
	case LAErrorBiometryDisconnected:
		return "LAErrorBiometryDisconnected"
	case LAErrorBiometryLockout:
		return "LAErrorBiometryLockout"
	case LAErrorBiometryNotAvailable:
		return "LAErrorBiometryNotAvailable"
	case LAErrorBiometryNotEnrolled:
		return "LAErrorBiometryNotEnrolled"
	case LAErrorBiometryNotPaired:
		return "LAErrorBiometryNotPaired"
	case LAErrorCompanionNotAvailable:
		return "LAErrorCompanionNotAvailable"
	case LAErrorInvalidContext:
		return "LAErrorInvalidContext"
	case LAErrorNotInteractive:
		return "LAErrorNotInteractive"
	case LAErrorPasscodeNotSet:
		return "LAErrorPasscodeNotSet"
	case LAErrorSystemCancel:
		return "LAErrorSystemCancel"
	case LAErrorUserCancel:
		return "LAErrorUserCancel"
	case LAErrorUserFallback:
		return "LAErrorUserFallback"
	default:
		return fmt.Sprintf("LAError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy
type LAPolicy int

const (
	// LAPolicyDeviceOwnerAuthentication: User authentication with biometry, Apple Watch, or the device passcode.
	LAPolicyDeviceOwnerAuthentication LAPolicy = 2
	// LAPolicyDeviceOwnerAuthenticationWithBiometrics: User authentication with biometry.
	LAPolicyDeviceOwnerAuthenticationWithBiometrics LAPolicy = 1
	// LAPolicyDeviceOwnerAuthenticationWithBiometricsOrCompanion: Device owner will be authenticated by biometry or a companion device e.g.
	LAPolicyDeviceOwnerAuthenticationWithBiometricsOrCompanion LAPolicy = 4
	// LAPolicyDeviceOwnerAuthenticationWithCompanion: Device owner will be authenticated by a companion device e.g.
	LAPolicyDeviceOwnerAuthenticationWithCompanion LAPolicy = 3
	// LAPolicyDeviceOwnerAuthenticationWithWristDetection: User authentication with wrist detection on watchOS.
	LAPolicyDeviceOwnerAuthenticationWithWristDetection LAPolicy = 5
	// Deprecated.
	LAPolicyDeviceOwnerAuthenticationWithBiometricsOrWatch LAPolicy = 4
	// Deprecated.
	LAPolicyDeviceOwnerAuthenticationWithWatch LAPolicy = 3
)

func (e LAPolicy) String() string {
	switch e {
	case LAPolicyDeviceOwnerAuthentication:
		return "LAPolicyDeviceOwnerAuthentication"
	case LAPolicyDeviceOwnerAuthenticationWithBiometrics:
		return "LAPolicyDeviceOwnerAuthenticationWithBiometrics"
	case LAPolicyDeviceOwnerAuthenticationWithBiometricsOrCompanion:
		return "LAPolicyDeviceOwnerAuthenticationWithBiometricsOrCompanion"
	case LAPolicyDeviceOwnerAuthenticationWithCompanion:
		return "LAPolicyDeviceOwnerAuthenticationWithCompanion"
	case LAPolicyDeviceOwnerAuthenticationWithWristDetection:
		return "LAPolicyDeviceOwnerAuthenticationWithWristDetection"
	default:
		return fmt.Sprintf("LAPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/LocalAuthentication/LARight/State-swift.enum
type LARightState int

const (
	// LARightStateAuthorized: The authorization completed successfully.
	LARightStateAuthorized LARightState = 2
	// LARightStateAuthorizing: The authorization is in progress but not completed.
	LARightStateAuthorizing LARightState = 1
	// LARightStateNotAuthorized: The authorization failed.
	LARightStateNotAuthorized LARightState = 3
	// LARightStateUnknown: The authorization is in an unknown state.
	LARightStateUnknown LARightState = 0
)

func (e LARightState) String() string {
	switch e {
	case LARightStateAuthorized:
		return "LARightStateAuthorized"
	case LARightStateAuthorizing:
		return "LARightStateAuthorizing"
	case LARightStateNotAuthorized:
		return "LARightStateNotAuthorized"
	case LARightStateUnknown:
		return "LARightStateUnknown"
	default:
		return fmt.Sprintf("LARightState(%d)", e)
	}
}
