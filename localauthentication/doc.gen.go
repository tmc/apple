// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

// Package localauthentication provides Go bindings for the LocalAuthentication framework.
//
// Authenticate users biometrically or with a passphrase they already know.
//
// Many users rely on biometric authentication like Face ID, Touch ID, or Optic ID to enable secure, effortless access to their devices. As a fallback option, and for devices without biometry, a passcode or password serves a similar purpose. Use the LocalAuthentication framework to leverage these mechanisms in your app and extend authentication procedures your app already implements.
//
// # Essentials
//
//   - Logging a User into Your App with Face ID or Touch ID: Supplement your own authentication scheme with biometric authentication, making it easy for users to access sensitive parts of your app.
//   - Accessing Keychain Items with Face ID or Touch ID: Protect a keychain item with biometric authentication.
//
// # Authentication and access
//
//   - LARight: A grouped set of requirements that gate access to a resource or operation.
//   - LARight.State: The possible states for a right during authorization.
//   - LAContext: A mechanism for evaluating authentication policies and access controls. ([LAPolicy], [LABiometryType], [LAAccessControlOperation], [LACredentialType])
//
// # Persistence
//
//   - LARightStore: A container for data protected by a right.
//   - LAPersistedRight: A right that gates access to a key and a secret.
//   - LASecret: Data that’s protected by a persisted right.
//
// # Key pairs
//
//   - LAPublicKey: The public portion of an asymmetric key pair.
//   - LAPrivateKey: The private portion of an asymmetric key pair.
//
// # Requirements
//
//   - LAAuthenticationRequirement: A set of requirements that protect a right.
//   - LABiometryFallbackRequirement: A set of requirements to fall back on if biometrics aren’t present.
//
// # Authentication views
//
//   - LocalAuthenticationView: A SwiftUI view that displays an authentication interface.
//
// # Errors
//
//   - LAError: Errors issued by the LocalAuthentication framework.
//   - LAError.Code: Errors issued by the LocalAuthentication framework.
//   - LAErrorDomain: The error domain that the framework uses when issuing errors.
//
// # Classes
//
//   - LADomainState
//   - LADomainStateBiometry
//   - LADomainStateCompanion
//   - LAEnvironment
//
// # Variables
//
//   - kLAAccessControlOperationCreateItem
//   - kLAAccessControlOperationCreateKey
//   - kLAAccessControlOperationUseItem
//   - kLAAccessControlOperationUseKeyDecrypt
//   - kLAAccessControlOperationUseKeyKeyExchange
//   - kLAAccessControlOperationUseKeySign
//   - kLACompanionTypeMac
//   - kLACompanionTypeNone
//   - kLACompanionTypeVision
//   - kLACompanionTypeWatch
//   - kLAErrorCompanionNotAvailable
//   - kLAPolicyDeviceOwnerAuthenticationWithBiometricsOrCompanion
//   - kLAPolicyDeviceOwnerAuthenticationWithCompanion
//
// # Enumerations
//
//   - LACompanionType
//
// # Key Types
//
//   - [LAContext] - A mechanism for evaluating authentication policies and access controls.
//   - [LARight] - A grouped set of requirements that gate access to a resource or operation.
//   - [LAPrivateKey] - The private portion of an asymmetric key pair.
//   - [LARightStore] - A container for data protected by a right.
//   - [LAEnvironmentMechanismBiometry]
//   - [LAPersistedRight] - A right that gates access to a key and a secret.
//   - [LAPublicKey] - The public portion of an asymmetric key pair.
//   - [LAAuthenticationRequirement] - A set of requirements that protect a right.
//   - [LAEnvironment]
//   - [LAEnvironmentState]
//
// [LocalAuthentication Documentation]: https://developer.apple.com/documentation/LocalAuthentication
package localauthentication

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the LocalAuthentication library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/LocalAuthentication.framework/LocalAuthentication",
	"/usr/lib/libLocalAuthentication.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: LocalAuthentication: failed to load framework from any known path\n")
}
