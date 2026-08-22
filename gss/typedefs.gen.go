// Code generated from Apple documentation. DO NOT EDIT.

package gss

import (
	"unsafe"
)

// OM_uint32 is a 32-bit unsigned integer.
//
// See: https://developer.apple.com/documentation/GSS/OM_uint32
type OM_uint32 = uint32

// OM_uint64 is a 64-bit unsigned integer.
//
// See: https://developer.apple.com/documentation/GSS/OM_uint64
type OM_uint64 = uint64

// Gss_OID is a pointer to the OID descriptor that exchanges object identifiers with many GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_OID
type Gss_OID = uintptr

// Gss_OID_desc is the OID descriptor that exchanges object identifiers with many GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_OID_desc
type Gss_OID_desc = Gss_OID_desc_struct

// Gss_OID_set is a pointer to a descriptor that manages an array of OID descriptors.
//
// See: https://developer.apple.com/documentation/GSS/gss_OID_set
type Gss_OID_set = uintptr

// Gss_OID_set_desc is the descriptor that manages an array of OID descriptors.
//
// See: https://developer.apple.com/documentation/GSS/gss_OID_set_desc
type Gss_OID_set_desc = Gss_OID_set_desc_struct

// Gss_auth_identity_t is a pointer to an opaque object used to manage authentication identities.
//
// See: https://developer.apple.com/documentation/GSS/gss_auth_identity_t
type Gss_auth_identity_t = uintptr

// Gss_buffer_desc is the buffer descriptor that you use to exchange octet streams with many GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_buffer_desc
type Gss_buffer_desc = Gss_buffer_desc_struct

// Gss_buffer_set_desc is the descriptor that you use to manage an array of buffer descriptors.
//
// See: https://developer.apple.com/documentation/GSS/gss_buffer_set_desc
type Gss_buffer_set_desc = Gss_buffer_set_desc_struct

// Gss_buffer_set_t is a pointer to the descriptor that you use to manage an array of buffer descriptors.
//
// See: https://developer.apple.com/documentation/GSS/gss_buffer_set_t
type Gss_buffer_set_t = uintptr

// Gss_buffer_t is a pointer to a buffer descriptor that you use to exchange octet streams with many GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_buffer_t
type Gss_buffer_t = uintptr

// Gss_channel_bindings_t is a pointer to a channel bindings descriptor that specifies the communications channel used to carry a context.
//
// See: https://developer.apple.com/documentation/GSS/gss_channel_bindings_t
type Gss_channel_bindings_t = uintptr

// Gss_const_OID is a pointer to an immutable OID descriptor exchanges object identifiers with many GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_const_OID
type Gss_const_OID = unsafe.Pointer

// Gss_const_OID_set is a pointer to an immutable descriptor manages an array of OID descriptors.
//
// See: https://developer.apple.com/documentation/GSS/gss_const_OID_set
type Gss_const_OID_set = unsafe.Pointer

// Gss_const_buffer_t is a pointer to an immutable buffer descriptor that you use to exchange octet streams with many GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_const_buffer_t
type Gss_const_buffer_t = unsafe.Pointer

// Gss_const_channel_bindings_t is a pointer to an immutable channel bindings descriptor that you use to specify the communications channel used to carry a context.
//
// See: https://developer.apple.com/documentation/GSS/gss_const_channel_bindings_t
type Gss_const_channel_bindings_t = uintptr

// Gss_const_cred_id_t is a pointer to an immutable opaque type that you use to exchange a credential object with GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_const_cred_id_t
type Gss_const_cred_id_t = uintptr

// Gss_const_ctx_id_t is a pointer to an immutable opaque type that you use to communicate context pointers with GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_const_ctx_id_t
type Gss_const_ctx_id_t = unsafe.Pointer

// Gss_const_name_t is a pointer to an immutable version of the opaque descriptor used to exchange name objects with GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_const_name_t
type Gss_const_name_t = uintptr

// Gss_cred_id_t is a pointer to an opaque type that you use to exchange a credential object with GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_cred_id_t
type Gss_cred_id_t = uintptr

// Gss_cred_usage_t is a credential usage value.
//
// See: https://developer.apple.com/documentation/GSS/gss_cred_usage_t
type Gss_cred_usage_t = int32

// Gss_ctx_id_t is a pointer to an opaque type that you use to communicate context pointers with GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_ctx_id_t
type Gss_ctx_id_t = uintptr

// Gss_iov_buffer_desc is the structure for a vectored I/O buffer and its defined type.
//
// See: https://developer.apple.com/documentation/GSS/gss_iov_buffer_desc
type Gss_iov_buffer_desc = Gss_iov_buffer_desc_struct

// Gss_iov_buffer_t is the structure for a vectored I/O buffer and its defined type.
//
// See: https://developer.apple.com/documentation/GSS/gss_iov_buffer_t
type Gss_iov_buffer_t = uintptr

// Gss_krb5_cfx_keydata_t is the structure of a Kerberos context and acceptor-asserted key.
//
// See: https://developer.apple.com/documentation/GSS/gss_krb5_cfx_keydata_t
type Gss_krb5_cfx_keydata_t = Gss_krb5_cfx_keydata

// Gss_krb5_lucid_context_v1_t is the structure of a Kerberos context.
//
// See: https://developer.apple.com/documentation/GSS/gss_krb5_lucid_context_v1_t
type Gss_krb5_lucid_context_v1_t = Gss_krb5_lucid_context_v1

// Gss_krb5_lucid_context_version_t is the structure for determining the returned Kerberos lucid context structure version.
//
// See: https://developer.apple.com/documentation/GSS/gss_krb5_lucid_context_version_t
type Gss_krb5_lucid_context_version_t = Gss_krb5_lucid_context_version

// Gss_krb5_lucid_key_t is the structure for a Kerberos encryption key.
//
// See: https://developer.apple.com/documentation/GSS/gss_krb5_lucid_key_t
type Gss_krb5_lucid_key_t = Gss_krb5_lucid_key

// Gss_krb5_rfc1964_keydata_t is the structure for an RFC 1964-compliant Kerberos encryption key.
//
// See: https://developer.apple.com/documentation/GSS/gss_krb5_rfc1964_keydata_t
type Gss_krb5_rfc1964_keydata_t = Gss_krb5_rfc1964_keydata

// Gss_name_t is a pointer to an opaque type that you use to communicate name objects with GSS-API functions.
//
// See: https://developer.apple.com/documentation/GSS/gss_name_t
type Gss_name_t = uintptr

// Gss_qop_t is a quality of protection setting.
//
// See: https://developer.apple.com/documentation/GSS/gss_qop_t
type Gss_qop_t = uint32

// Gss_status_id_t is a pointer to a status result.
//
// See: https://developer.apple.com/documentation/GSS/gss_status_id_t
type Gss_status_id_t = *OM_uint32

// Gss_uint32 is a 32-bit unsigned integer.
//
// See: https://developer.apple.com/documentation/GSS/gss_uint32
type Gss_uint32 = uint32

// OmUint32 is a Go-name alias for OM_uint32.
type OmUint32 = OM_uint32

// OmUint64 is a Go-name alias for OM_uint64.
type OmUint64 = OM_uint64

// GssOid is a Go-name alias for Gss_OID.
type GssOid = Gss_OID

// GssOidDesc is a Go-name alias for Gss_OID_desc.
type GssOidDesc = Gss_OID_desc

// GssOidSet is a Go-name alias for Gss_OID_set.
type GssOidSet = Gss_OID_set

// GssOidSetDesc is a Go-name alias for Gss_OID_set_desc.
type GssOidSetDesc = Gss_OID_set_desc

// GssAuthIdentity is a Go-name alias for Gss_auth_identity_t.
type GssAuthIdentity = Gss_auth_identity_t

// GssBufferDesc is a Go-name alias for Gss_buffer_desc.
type GssBufferDesc = Gss_buffer_desc

// GssBufferSetDesc is a Go-name alias for Gss_buffer_set_desc.
type GssBufferSetDesc = Gss_buffer_set_desc

// GssBufferSet is a Go-name alias for Gss_buffer_set_t.
type GssBufferSet = Gss_buffer_set_t

// GssBuffer is a Go-name alias for Gss_buffer_t.
type GssBuffer = Gss_buffer_t

// GssChannelBindings is a Go-name alias for Gss_channel_bindings_t.
type GssChannelBindings = Gss_channel_bindings_t

// GssConstOid is a Go-name alias for Gss_const_OID.
type GssConstOid = Gss_const_OID

// GssConstOidSet is a Go-name alias for Gss_const_OID_set.
type GssConstOidSet = Gss_const_OID_set

// GssConstBuffer is a Go-name alias for Gss_const_buffer_t.
type GssConstBuffer = Gss_const_buffer_t

// GssConstChannelBindings is a Go-name alias for Gss_const_channel_bindings_t.
type GssConstChannelBindings = Gss_const_channel_bindings_t

// GssConstCredID is a Go-name alias for Gss_const_cred_id_t.
type GssConstCredID = Gss_const_cred_id_t

// GssConstCtxID is a Go-name alias for Gss_const_ctx_id_t.
type GssConstCtxID = Gss_const_ctx_id_t

// GssConstName is a Go-name alias for Gss_const_name_t.
type GssConstName = Gss_const_name_t

// GssCredID is a Go-name alias for Gss_cred_id_t.
type GssCredID = Gss_cred_id_t

// GssCredUsage is a Go-name alias for Gss_cred_usage_t.
type GssCredUsage = Gss_cred_usage_t

// GssCtxID is a Go-name alias for Gss_ctx_id_t.
type GssCtxID = Gss_ctx_id_t

// GssIovBufferDesc is a Go-name alias for Gss_iov_buffer_desc.
type GssIovBufferDesc = Gss_iov_buffer_desc

// GssIovBuffer is a Go-name alias for Gss_iov_buffer_t.
type GssIovBuffer = Gss_iov_buffer_t

// GssKrb5CfxKeydata is a Go-name alias for Gss_krb5_cfx_keydata_t.
type GssKrb5CfxKeydata = Gss_krb5_cfx_keydata_t

// GssKrb5LucidContextV1 is a Go-name alias for Gss_krb5_lucid_context_v1_t.
type GssKrb5LucidContextV1 = Gss_krb5_lucid_context_v1_t

// GssKrb5LucidContextVersion is a Go-name alias for Gss_krb5_lucid_context_version_t.
type GssKrb5LucidContextVersion = Gss_krb5_lucid_context_version_t

// GssKrb5LucidKey is a Go-name alias for Gss_krb5_lucid_key_t.
type GssKrb5LucidKey = Gss_krb5_lucid_key_t

// GssKrb5Rfc1964Keydata is a Go-name alias for Gss_krb5_rfc1964_keydata_t.
type GssKrb5Rfc1964Keydata = Gss_krb5_rfc1964_keydata_t

// GssName is a Go-name alias for Gss_name_t.
type GssName = Gss_name_t

// GssQop is a Go-name alias for Gss_qop_t.
type GssQop = Gss_qop_t

// GssStatusID is a Go-name alias for Gss_status_id_t.
type GssStatusID = Gss_status_id_t

// GssUint32 is a Go-name alias for Gss_uint32.
type GssUint32 = Gss_uint32
