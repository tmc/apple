// Code generated from Apple documentation. DO NOT EDIT.

package endpointsecurity

import (
	"unsafe"
	"github.com/tmc/apple/objectivec"
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
type Es_handler_block_t = func(objectivec.IObject, *Es_message_t)

// See: https://developer.apple.com/documentation/EndpointSecurity/es_sha256_t
type Es_sha256_t = unsafe.Pointer

// Es_statfs_t is this typedef is no longer used, but exists for API backwards compatibility.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_statfs_t
type Es_statfs_t = unsafe.Pointer

