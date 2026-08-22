// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

// Package cryptotokenkit provides Go bindings for the CryptoTokenKit framework.
//
// Access security tokens and the cryptographic assets they store.
//
// You use the CryptoTokenKit framework to easily access cryptographic tokens.
// Tokens are physical devices built in to the system, located on attached
// hardware (like a smart card), or accessible through a network connection.
// Tokens store cryptographic objects like keys and certificates. They also
// may perform operations—for example, encryption or digital signature
// verification—using these objects. You use the framework to work with a
// token’s assets as if they were part of your system, even though they
// remain secured by the token.
//
// # Smart Cards
//
//   - [Using Cryptographic Assets Stored on a Smart Card]: Access certificates, keys, and identities stored on a smart card as if they were part of the keychain.
//   - [TKSmartCardSlotManager]: An interface to all available smart card reader slots.
//   - [TKSmartCardSlot]: A single smart card reader slot in the system. ([TKSmartCardATR])
//   - [TKSmartCard]: A representation of a smart card. ([TKSmartCardProtocol], [TKSmartCardPINFormat], [TKSmartCardUserInteraction], [TKSmartCardUserInteractionForPINOperation], [TKSmartCardUserInteractionForSecurePINChange])
//
// # Smart Card App Extensions
//
//   - [Authenticating Users with a Cryptographic Token]: Grant access to user accounts and the keychain by creating a smart card app extension.
//   - [Configuring Smart Card Authentication]: Set preferences for smart card authentication operations, including those on managed devices.
//   - [TKSmartCardTokenDriver]: The driver that acts as an entry point for smart card app extensions. ([TKSmartCardTokenDriverDelegate])
//   - [TKSmartCardToken]: A representation of a smart card based cryptographic token. ([TKSmartCard], [TKTLVRecord], [TKBERTLVRecord], [TKCompactTLVRecord], [TKSimpleTLVRecord])
//   - [TKSmartCardTokenSession]: A token session that is based on a smart card token.
//
// # Tokens
//
//   - [TKTokenWatcher]: An object that tracks the tokens available in the system.
//   - [TKTokenDriver]: A base class for building token drivers. ([TKTokenDriverDelegate])
//   - [TKToken]: A representation of a hardware-based cryptographic token. ([TKTokenDelegate], [TKTokenKeychainContents], [TKTokenKeychainItem], [TKTokenKeychainCertificate], [TKTokenKeychainKey])
//   - [TKTokenSession]: A token session that manages the authentication state of a token. ([TKTokenSessionDelegate])
//
// # Errors
//
//   - [TKErrorDomain]: The domain for all CryptoTokenKit framework errors.
//   - [TKErrorCode]: Error codes from CryptoTokenKit.
//
// # Classes
//
//   - [TKTokenWatcherTokenInfo]//
//
// # Key Types
//
//   - [TKSmartCard] - A representation of a smart card.
//   - [TKTokenKeychainKey] - A token’s key as stored in the keychain.
//   - [TKSmartCardPINFormat] - The formatting properties for a PIN, such as character encoding and length constraints.
//   - [TKSmartCardATR] - A parsed ATR (Answer To Reset) message from a Smart Card.
//   - [TKSmartCardSlot] - A single smart card reader slot in the system.
//   - [TKSmartCardSlotManager] - An interface to all available smart card reader slots.
//   - [TKSmartCardUserInteraction] - The base class for encapsulating user interaction with a Smart Card reader.
//   - [TKSmartCardUserInteractionForPINOperation] - A representation of user interaction for secure PIN operations on a Smart Card reader.
//   - [TKTLVRecord] - The base class encapsulating a Tag-Length-Value record.
//   - [TKToken] - A representation of a hardware-based cryptographic token.
//
// [Authenticating Users with a Cryptographic Token]: https://developer.apple.com/documentation/cryptotokenkit/authenticating-users-with-a-cryptographic-token
// [Configuring Smart Card Authentication]: https://developer.apple.com/documentation/cryptotokenkit/configuring-smart-card-authentication
// [Using Cryptographic Assets Stored on a Smart Card]: https://developer.apple.com/documentation/cryptotokenkit/using-cryptographic-assets-stored-on-a-smart-card
package cryptotokenkit

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CryptoTokenKit library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CryptoTokenKit.framework/CryptoTokenKit",
	"/usr/lib/libCryptoTokenKit.dylib",
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
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: CryptoTokenKit: failed to load framework from any known path\n")
	}
}
