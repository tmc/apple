// Code generated from Apple documentation for GSS. DO NOT EDIT.

package gss

import (
	"unsafe"
)

// C struct types

// Gss_OID_desc_struct - The structure for an OID descriptor that exchanges object identifiers with many GSS-API functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_OID_desc_struct
type Gss_OID_desc_struct struct {
	Length   OM_uint32      // The number of octets in the object identifier.
	Elements unsafe.Pointer // A pointer to the octets that make up the object identifier.

}

// Gss_OID_set_desc_struct - The structure for an OID set descriptor that manages an array of OID descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_OID_set_desc_struct
type Gss_OID_set_desc_struct struct {
	Count    uintptr        // The number of OID descriptors in the array.
	Elements unsafe.Pointer // A pointer to the array of OID descriptors.

}

// Gss_buffer_desc_struct - The structure for a buffer descriptor that you use to exchange octet streams with many GSS-API functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_buffer_desc_struct
type Gss_buffer_desc_struct struct {
	Length uintptr        // The number of bytes held in the buffer.
	Value  unsafe.Pointer // A pointer to the bytes in the buffer.

}

// Gss_buffer_set_desc_struct - The structure for a buffer set descriptor that you use to manage an array of buffer descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_buffer_set_desc_struct
type Gss_buffer_set_desc_struct struct {
	Count    uintptr          // The number of buffer descriptors in the array.
	Elements *Gss_buffer_desc // A pointer to an array of buffer descriptors.

}

// Gss_channel_bindings_struct - The structure defining a channel bindings descriptor that specifies the communications channel used to carry a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_channel_bindings_struct
type Gss_channel_bindings_struct struct {
	Initiator_addrtype OM_uint32       // The type of address contained in the [initiator_address](<https://developer.apple.com/documentation/GSS/gss_channel_bindings_struct/initiator_address>) field.
	Initiator_address  Gss_buffer_desc // The network address of the acceptor, in the form specified by [initiator_addrtype](<https://developer.apple.com/documentation/GSS/gss_channel_bindings_struct/initiator_addrtype>).
	Acceptor_addrtype  OM_uint32       // The type of address contained in the [acceptor_address](<https://developer.apple.com/documentation/GSS/gss_channel_bindings_struct/acceptor_address>) field.
	Acceptor_address   Gss_buffer_desc // The network address of the acceptor, in the form specified by [acceptor_addrtype](<https://developer.apple.com/documentation/GSS/gss_channel_bindings_struct/acceptor_addrtype>).
	Application_data   Gss_buffer_desc // Application specific data for use in communicating a channel binding.

}

// Gss_iov_buffer_desc_struct - The structure for a vectored I/O buffer and its defined type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_iov_buffer_desc_struct
type Gss_iov_buffer_desc_struct struct {
	Type   OM_uint32       // The buffer type.
	Buffer Gss_buffer_desc // The buffer length and contents.

}

// Gss_krb5_cfx_keydata
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_cfx_keydata
type Gss_krb5_cfx_keydata struct {
	Have_acceptor_subkey OM_uint32            // The flag that indicates if the Kerberos session acceptor subkey is available.
	Ctx_key              Gss_krb5_lucid_key_t // The Kerberos session context key.
	Acceptor_subkey      Gss_krb5_lucid_key_t // The Kerberos session acceptor subkey.

}

// Gss_krb5_lucid_context_v1
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_lucid_context_v1
type Gss_krb5_lucid_context_v1 struct {
	Version    OM_uint32                  // The structure version number.
	Initiate   OM_uint32                  // The flag indicating if the role is initiator.
	Endtime    OM_uint32                  // The expiration time of the context.
	Send_seq   OM_uint64                  // The send sequence number.
	Recv_seq   OM_uint64                  // The receive sequence number.
	Protocol   OM_uint32                  // The protocol to use.
	Rfc1964_kd Gss_krb5_rfc1964_keydata_t // The RFC-1964 key data.
	Cfx_kd     Gss_krb5_cfx_keydata_t     // The key data structure for Kerberos 5.

}

// Gss_krb5_lucid_context_version
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_lucid_context_version
type Gss_krb5_lucid_context_version struct {
	Version OM_uint32 // The structure version number.

}

// Gss_krb5_lucid_key
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_lucid_key
type Gss_krb5_lucid_key struct {
	Type   OM_uint32      // The key encryption type.
	Length OM_uint32      // The length of the key data.
	Data   unsafe.Pointer // The actual key data.

}

// Gss_krb5_rfc1964_keydata
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_rfc1964_keydata
type Gss_krb5_rfc1964_keydata struct {
	Sign_alg OM_uint32            // The signing algorithm.
	Seal_alg OM_uint32            // The seal/encrypt algorithm.
	Ctx_key  Gss_krb5_lucid_key_t // The context key (Kerberos session key or subkey).

}
