// Code generated from Apple documentation. DO NOT EDIT.

package endpointsecurity

import (
	"unsafe"

	"github.com/tmc/apple/kernel"
)

// See: https://developer.apple.com/documentation/EndpointSecurity/es_cdhash_t
type EsCdhash = [20]uint8

// EsClient is an opaque type that stores the Endpoint Security client state.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_client_t
// EsClient is an unresolved C aggregate typedef.
type EsClient unsafe.Pointer

// See: https://developer.apple.com/documentation/EndpointSecurity/es_graphical_session_id_t
type EsGraphicalSessionID = uint32

// EsHandlerBlock is a block that handles a message received from Endpoint Security.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_handler_block_t
type EsHandlerBlock = func(*Es_client_t, *Es_message_t)

// See: https://developer.apple.com/documentation/EndpointSecurity/es_sha256_t
type EsSha256 = [32]uint8

// EsStatfs is this typedef is no longer used, but exists for API backwards compatibility.
//
// See: https://developer.apple.com/documentation/EndpointSecurity/es_statfs_t
type EsStatfs = kernel.Pointer

// Es_cdhash_t is a C-name alias for EsCdhash.
type Es_cdhash_t = EsCdhash

// Es_client_t is a C-name alias for EsClient.
type Es_client_t = EsClient

// Es_graphical_session_id_t is a C-name alias for EsGraphicalSessionID.
type Es_graphical_session_id_t = EsGraphicalSessionID

// Es_handler_block_t is a C-name alias for EsHandlerBlock.
type Es_handler_block_t = EsHandlerBlock

// Es_sha256_t is a C-name alias for EsSha256.
type Es_sha256_t = EsSha256

// Es_statfs_t is a C-name alias for EsStatfs.
type Es_statfs_t = EsStatfs
