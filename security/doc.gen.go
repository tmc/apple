// Code generated from Apple documentation for Security. DO NOT EDIT.

// Package security provides Go bindings for the Security framework.
//
// Secure the data your app manages, and control access to your app.
//
// Use the Security framework to protect information, establish trust, and control access to software. Broadly, security services support these goals:
//
// # Essentials
//
//   - Security updates: Learn about important changes to Security.
//
// # Authorization and authentication
//
//   - Password AutoFill: Streamline your app’s login and onboarding procedures.
//   - Shared Web Credentials: Share credentials between iOS apps and their website counterparts.
//   - Authorization Services: Access restricted areas of the operating system, and control access to particular features of your macOS app. ([AuthorizationFlags], [AuthorizationRef], [AuthorizationItem], [AuthorizationItemSet], [AuthorizationRights])
//   - Authorization Plug-ins: Extend the authorization services API by creating plug-ins that can participate in authorization decisions.
//   - Sessions: Manage login, authorization, and security sessions in macOS. ([SessionCreationFlags], [SessionAttributeBits], [SecuritySessionId])
//   - One-time codes: Streamline entry of authentication and recovery codes.
//
// # Secure data
//
//   - Keychain services: Securely store small chunks of data on behalf of the user.
//   - Preventing Insecure Network Connections: Enforce secure network links in your app by relying on App Transport Security.
//
// # Secure code
//
//   - Code Signing Services: Examine and validate signed code running on the system. ([SecCSFlags], [SecCodeSignatureFlags], [SecCSDigestAlgorithm], [SecRequirementType], [SecCodeStatus])
//   - Notarizing macOS software before distribution: Give users even more confidence in your macOS software by submitting it to Apple for notarization.
//   - Preparing your app to work with pointer authentication: Test your app against the arm64e architecture to ensure that it works seamlessly with enhanced security features.
//   - App Sandbox: Restrict access to system resources and user data in macOS apps to contain damage if an app becomes compromised.
//   - Hardened Runtime: Manage security protections and resource access for your macOS apps.
//   - Disabling and Enabling System Integrity Protection: Disable system protections only temporarily during development to test drivers, kernel extensions, and other low-level code.
//   - Using the latest code signature format: Update legacy app code signatures so your app runs on current OS releases.
//   - Updating Mac Software: Implement Mac software updates without causing code-signing crashes.
//   - TN3125: Inside Code Signing: Provisioning Profiles: Learn how provisioning profiles enable third-party code to run on Apple platforms.
//
// # Launch environment constraints
//
//   - Applying launch environment and library constraints: Limit the libraries your process loads, and the situations where it runs.
//   - Defining launch environment and library constraints: Restrict your app’s components to their expected contexts.
//   - Constraining a tool’s launch environment: Improve the security of your macOS app by limiting the ways its components can run.
//
// # Cryptography
//
//   - Complying with Encryption Export Regulations: Declare the use of encryption in your app to streamline the app submission process.
//   - Certificate, Key, and Trust Services: Establish trust using certificates and cryptographic keys.
//   - Cryptographic Message Syntax Services: Cryptographically sign and encrypt S/MIME messages. ([CMSSignedAttributes], [CMSCertificateChainMode], [CMSSignerStatus])
//   - Randomization Services: Generate cryptographically secure random numbers. ([SecRandomRef])
//   - Security Transforms: Perform cryptographic functions like encoding, encryption, signing, and signature verification. ([SecTransformInstanceBlock], [SecTransformImplementationRef], [SecTransformMetaAttributeType], [SecTransformDataBlock], [SecMessageBlock])
//   - ASN.1: Encode and decode Distinguished Encoding Rules (DER) and Basic Encoding Rules (BER) data streams. ([SecAsn1Item], [SecAsn1Template_struct], [SecAsn1Template_struct], [SecAsn1AlgId], [SecAsn1PubKeyInfo])
//
// # Result codes
//
//   - Security Framework Result Codes: Evaluate result codes common to many Security framework functions.
//
// # Legacy interfaces
//
//   - Common Security Services Manager: A set of open source modules underpinning the legacy implementation of the Security framework. ([Cssm_applecspdl_db_change_password_parameters], [Cssm_applecspdl_db_is_locked_parameters], [Cssm_applecspdl_db_settings_parameters], [Cssm_appledl_open_parameters], [Cssm_authorizationgroup])
//   - Secure Transport: Secure network communication using standardized transport layer security mechanisms. ([SSLProtocolSide], [SSLConnectionType], [SSLSessionOption], [SSLReadFunc], [SSLWriteFunc])
//   - Secure Download: Implement Apple’s Secure Download System in macOS.
//   - Security legacy reference: Learn about legacy APIs. ([MDS_HANDLE], [OS_sec_certificate], [OS_sec_identity], [OS_sec_object], [OS_sec_protocol_metadata])
//
// # Variables
//
//   - CSSM_APPLE_PRIVATE_CSPDL_CODE_28
//   - TLS_ECDHE_PSK_WITH_CHACHA20_POLY1305_SHA256
//   - errSecMissingQualifiedCertStatement
//   - kSecPolicyAppleEAPClient
//   - kSecPolicyAppleEAPServer
//   - kSecPolicyAppleIPSecClient
//   - kSecPolicyAppleIPSecServer
//   - kSecPolicyAppleSSLClient
//   - kSecPolicyAppleSSLServer
//   - kSecTrustQCStatements
//   - kSecTrustQWACValidation
//
// # Functions
//
//   - SecIdentityCreate(_:_:_:)
//   - sec_protocol_metadata_copy_negotiated_protocol(_:)
//   - sec_protocol_metadata_copy_server_name(_:)
//
// # Type Aliases
//
//   - CE_DataType
//   - CE_ExtendedKeyUsage
//   - CE_GeneralNameType
//
// [Security Documentation]: https://developer.apple.com/documentation/Security
package security

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Security library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Security.framework/Security",
	"/usr/lib/libSecurity.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: Security: failed to load framework from any known path\n")
}
