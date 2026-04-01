// Code generated from Apple documentation. DO NOT EDIT.

package endpointsecurity

import (
	"unsafe"
)

// See: https://developer.apple.com/documentation/EndpointSecurity/es_cdhash_t
type Es_cdhash_t = unsafe.Pointer

// Es_client_t is an opaque type that stores the Endpoint Security client state.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_client_t
type Es_client_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/EndpointSecurity/es_graphical_session_id_t
type Es_graphical_session_id_t = uint32

// Es_handler_block_t is a block that handles a message received from Endpoint Security.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_handler_block_t
type Es_handler_block_t = func(*Es_client_t, *Es_message_t)

// See: https://developer.apple.com/documentation/EndpointSecurity/es_sha256_t
type Es_sha256_t = unsafe.Pointer

// Es_statfs_t is this typedef is no longer used, but exists for API backwards compatibility.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_statfs_t
type Es_statfs_t = unsafe.Pointer

// Es_auto_unlock_type_t aliases the generated auto-unlock enum.

type Es_auto_unlock_type_t = EsAutoUnlock

// Es_mute_inverted_return_t aliases the generated mute state enum.

type Es_mute_inverted_return_t = EsMute

// Es_openssh_login_result_type_t aliases the generated OpenSSH login result enum.

type Es_openssh_login_result_type_t = EsOpenssh

// Es_set_or_clear_t aliases the generated set-or-clear enum.

type Es_set_or_clear_t = Es

// Es_xpc_domain_type_t aliases the generated XPC domain enum.

type Es_xpc_domain_type_t = EsXPCDomainType
