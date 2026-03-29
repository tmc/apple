
// Code generated from Apple documentation for EndpointSecurity. DO NOT EDIT.

// Package endpointsecurity provides Go bindings for the EndpointSecurity framework.
//
// Develop system extensions that enhance user security.
//
// Endpoint Security is a C API for monitoring system events for potentially malicious activity. You can write your client in any language that supports native calls. Your client registers with Endpoint Security to authorize pending events, or receive notifications of events that already occurred. These events include process executions, mounting file systems, forking processes, and raising signals.
//
// # Event Monitoring
//
//   - Client: An opaque type that maintains Endpoint Security client state, and functions related to this type. ([Es_handler_block_t], [Es_muted_processes_t], [Es_muted_paths_t])
//   - Message: A type used by Endpoint Security to notify your client when a monitored action occurs. ([Es_message_t], [Es_result_t], [Es_string_token_t], [Es_token_t])
//   - Event Types: Types used by messages to deliver details specific to different kinds of Endpoint Security events. ([Es_file_t], [Es_event_access_t], [Es_event_clone_t], [Es_event_copyfile_t], [Es_event_create_t])
//
// # Entitlements
//
//   - com.apple.developer.endpoint-security.client: The entitlement required to monitor system events for potentially malicious activity.
//
// # Variables
//
//   - ES_CS_VALIDATION_CATEGORY_APP_STORE
//   - ES_CS_VALIDATION_CATEGORY_DEVELOPER_ID
//   - ES_CS_VALIDATION_CATEGORY_DEVELOPMENT
//   - ES_CS_VALIDATION_CATEGORY_ENTERPRISE
//   - ES_CS_VALIDATION_CATEGORY_INVALID
//   - ES_CS_VALIDATION_CATEGORY_LOCAL_SIGNING
//   - ES_CS_VALIDATION_CATEGORY_NONE
//   - ES_CS_VALIDATION_CATEGORY_OOPJIT
//   - ES_CS_VALIDATION_CATEGORY_PLATFORM
//   - ES_CS_VALIDATION_CATEGORY_ROSETTA
//   - ES_CS_VALIDATION_CATEGORY_TESTFLIGHT
//   - ES_EVENT_TYPE_NOTIFY_TCC_MODIFY
//   - ES_TCC_AUTHORIZATION_REASON_APP_TYPE_POLICY: A system process changed the authorization right
//   - ES_TCC_AUTHORIZATION_REASON_ENTITLED: A system process changed the authorization right
//   - ES_TCC_AUTHORIZATION_REASON_ERROR
//   - ES_TCC_AUTHORIZATION_REASON_MDM_POLICY: A system process changed the authorization right
//   - ES_TCC_AUTHORIZATION_REASON_MISSING_USAGE_STRING: A system process changed the authorization right
//   - ES_TCC_AUTHORIZATION_REASON_NONE
//   - ES_TCC_AUTHORIZATION_REASON_PREFLIGHT_UNKNOWN: A system process changed the authorization right
//   - ES_TCC_AUTHORIZATION_REASON_PROMPT_CANCEL: A system process changed the authorization right
//   - ES_TCC_AUTHORIZATION_REASON_PROMPT_TIMEOUT: A system process changed the authorization right
//   - ES_TCC_AUTHORIZATION_REASON_SERVICE_OVERRIDE_POLICY: A system process changed the authorization right
//   - ES_TCC_AUTHORIZATION_REASON_SERVICE_POLICY: A system process changed the authorization right
//   - ES_TCC_AUTHORIZATION_REASON_SYSTEM_SET: User changed the authorization right via Preferences
//   - ES_TCC_AUTHORIZATION_REASON_USER_CONSENT
//   - ES_TCC_AUTHORIZATION_REASON_USER_SET: User answered a prompt
//   - ES_TCC_AUTHORIZATION_RIGHT_ADD_MODIFY_ADDED
//   - ES_TCC_AUTHORIZATION_RIGHT_ALLOWED
//   - ES_TCC_AUTHORIZATION_RIGHT_DENIED
//   - ES_TCC_AUTHORIZATION_RIGHT_LEARN_MORE
//   - ES_TCC_AUTHORIZATION_RIGHT_LIMITED
//   - ES_TCC_AUTHORIZATION_RIGHT_SESSION_PID
//   - ES_TCC_AUTHORIZATION_RIGHT_UNKNOWN
//   - ES_TCC_EVENT_TYPE_CREATE
//   - ES_TCC_EVENT_TYPE_DELETE
//   - ES_TCC_EVENT_TYPE_MODIFY
//   - ES_TCC_EVENT_TYPE_UNKNOWN
//   - ES_TCC_IDENTITY_TYPE_BUNDLE_ID
//   - ES_TCC_IDENTITY_TYPE_EXECUTABLE_PATH
//   - ES_TCC_IDENTITY_TYPE_FILE_PROVIDER_DOMAIN_ID
//   - ES_TCC_IDENTITY_TYPE_POLICY_ID
//
// # Type Aliases
//
//   - es_statfs_t: This typedef is no longer used, but exists for API backwards compatibility.
//
// [EndpointSecurity Documentation]: https://developer.apple.com/documentation/EndpointSecurity
package endpointsecurity

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/EndpointSecurity.framework/EndpointSecurity"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: EndpointSecurity: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

