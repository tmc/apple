// Code generated from Apple documentation for GSS. DO NOT EDIT.

// Package gss provides Go bindings for the GSS framework.
//
// Conduct secure, authenticated network transactions.
//
// The open source Generic Security Service Application Programming Interface
// (GSS-API) defines a standardized interface through which the operating
// system vends secure data transport operations. The GSS framework provides
// an implementation of the interface and the underlying libraries.
//
// # Memory and Context
//
//   - [Allocating and Releasing Objects]: Manage memory and object lifetimes.
//   - [Function Status]: Evaluate return values that most GSS-API functions use to indicate the outcome of an operation. ([OM_uint32], [OM_uint64])
//   - [Buffer Management]: Allocate and deallocate buffers with structures that hold a variety of data. ([Gss_iov_buffer_desc_struct])
//   - [Context Services]: Use context services to manage secure operations between endpoints. ([Gss_channel_bindings_struct])
//
// # Credentials
//
//   - [Credential Management]: Securely establish connections between endpoints.
//   - [Security Mechanisms]: Provide a security mechanism for your implementation.
//
// # Names and Object Identifiers
//
//   - [Name Handling]: Manage names for GSS-API principals such as a person, a machine, or an application.
//   - [Object Identifiers]: Store security mechanisms, QOPs (Quality of Protection values), and name types.
//
// # Messages
//
//   - [Token Management]: Establish secure communication with tokens.
//   - [Message Protection]: Provide cryptographic protection to secure message integrity.
//   - [Kerberos Implementation]: Establish secure connections using the Kerberos implementation of GSS-API.
//
// # Structure and macros
//
//   - [Structures and macros] ([Gss_krb5_cfx_keydata], [Gss_krb5_lucid_context_v1], [Gss_krb5_lucid_context_version], [Gss_krb5_lucid_key], [Gss_krb5_rfc1964_keydata])//
//
// [Allocating and Releasing Objects]: https://developer.apple.com/documentation/gss/allocating-and-releasing-objects
// [Buffer Management]: https://developer.apple.com/documentation/gss/buffer-management
// [Context Services]: https://developer.apple.com/documentation/gss/context-services
// [Credential Management]: https://developer.apple.com/documentation/gss/credential-management
// [Function Status]: https://developer.apple.com/documentation/gss/function-status
// [Kerberos Implementation]: https://developer.apple.com/documentation/gss/kerberos-implementation
// [Message Protection]: https://developer.apple.com/documentation/gss/message-protection
// [Name Handling]: https://developer.apple.com/documentation/gss/name-handling
// [Object Identifiers]: https://developer.apple.com/documentation/gss/object-identifiers
// [Security Mechanisms]: https://developer.apple.com/documentation/gss/security-mechanisms
// [Structures and macros]: https://developer.apple.com/documentation/gss/structures-and-macros
// [Token Management]: https://developer.apple.com/documentation/gss/token-management
package gss

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the GSS library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/GSS.framework/GSS",
	"/usr/lib/libGSS.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: GSS: failed to load framework from any known path\n")
	}
}
