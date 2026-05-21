// Code generated from Apple documentation for EndpointSecurity. DO NOT EDIT.

// Package endpointsecurity provides Go bindings for the EndpointSecurity framework.
//
// Develop system extensions that enhance user security.
//
// Endpoint Security is a C API for monitoring system events for potentially
// malicious activity. You can write your client in any language that supports
// native calls. Your client registers with Endpoint Security to authorize
// pending events, or receive notifications of events that already occurred.
// These events include process executions, mounting file systems, forking
// processes, and raising signals.
//
// # Event Monitoring
//
//   - Client: An opaque type that maintains Endpoint Security client state, and functions related to this type. ([Es_handler_block_t], [Es_new_client_result_t], [Es_event_type_t], [Es_auth_result_t], [Es_respond_result_t])
//   - Message: A type used by Endpoint Security to notify your client when a monitored action occurs. ([Es_message_t], [Es_result_t], [Es_string_token_t], [Es_token_t])
//   - [Event Types]: Types used by messages to deliver details specific to different kinds of Endpoint Security events. ([Es_file_t], [Es_event_access_t], [Es_event_clone_t], [Es_event_copyfile_t], [Es_event_create_t])
//   - [Monitoring System Events with Endpoint Security]: Receive notifications and authorization requests for sensitive operations by creating an Endpoint Security client for your app.
//
// # Entitlements
//
//   - [com.apple.developer.endpoint-security.client]: The entitlement required to monitor system events for potentially malicious activity.
//
// # Type Aliases
//
//   - [Es_statfs_t]: This typedef is no longer used, but exists for API backwards compatibility.
//
// # Enumerations
//
//   - [Es_cs_validation_category_t]: es_cs_validation_category
//   - [Es_tcc_authorization_reason_t]: ess_tcc_authorization_reason_t
//   - [Es_tcc_authorization_right_t]: ess_tcc_authorization_right_t
//   - [Es_tcc_event_type_t]
//   - [Es_tcc_identity_type_t]: es_tcc_identity_type_t//
//
// [Event Types]: https://developer.apple.com/documentation/endpointsecurity/event-types
// [Monitoring System Events with Endpoint Security]: https://developer.apple.com/documentation/endpointsecurity/monitoring-system-events-with-endpoint-security
// [com.apple.developer.endpoint-security.client]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.developer.endpoint-security.client
package endpointsecurity

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the EndpointSecurity library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/EndpointSecurity.framework/EndpointSecurity",
	"/usr/lib/libEndpointSecurity.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: EndpointSecurity: failed to load framework from any known path\n")
}
