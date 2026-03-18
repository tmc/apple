// Code generated from Apple documentation. DO NOT EDIT.

package security

import (
	"unsafe"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// AuthorizationAsyncCallback is a block used as a callback for the asynchronous version of copying authorization rights.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationAsyncCallback
type AuthorizationAsyncCallback = func(int, *AuthorizationItemSet)

// AuthorizationEngineRef is handle passed from the authorization engine to an instance of a mechanism in a plug-in.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationEngineRef
type AuthorizationEngineRef uintptr

// AuthorizationEnvironment is an authorization item set designated to hold environment information relevant to authorization decisions.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationEnvironment
type AuthorizationEnvironment = AuthorizationItemSet

// AuthorizationMechanismId is the mechanism ID specified in the authorization policy database is passed to the plug-in to create the appropriate mechanism.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationMechanismId
type AuthorizationMechanismId = unsafe.Pointer

// AuthorizationMechanismRef is a handle passed by the plug-in to the authorization engine when creating an instance of a mechanism.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationMechanismRef
type AuthorizationMechanismRef uintptr

// AuthorizationPluginId is an unused identifier for a plug-in.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationPluginId
type AuthorizationPluginId = unsafe.Pointer

// AuthorizationPluginRef is a handle passed by the plug-in to the authorization engine when the plug-in is initiated.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationPluginRef
type AuthorizationPluginRef uintptr

// AuthorizationRef is a pointer to an opaque authorization reference structure.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationRef
type AuthorizationRef uintptr

// AuthorizationRights is an authorization item set designated to represent a set of rights.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationRights
type AuthorizationRights = AuthorizationItemSet

// AuthorizationSessionId is a unique value for an authorization session, provided by the authorization engine.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationSessionId
type AuthorizationSessionId = unsafe.Pointer

// AuthorizationString is a zero-terminated string in UTF-8 encoding.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationString
type AuthorizationString = *byte

// See: https://developer.apple.com/documentation/Security/CE_CrlNumber
type CE_CrlNumber = uint32

// See: https://developer.apple.com/documentation/Security/CE_DeltaCrl
type CE_DeltaCrl = uint32

// See: https://developer.apple.com/documentation/Security/CE_ExtendedKeyUsage-swift.typealias
type CE_ExtendedKeyUsage = unsafe.Pointer

// CMSDecoderRef is an opaque reference to a CMS decoder object.
//
// See: https://developer.apple.com/documentation/Security/CMSDecoder
type CMSDecoderRef uintptr

// CMSEncoderRef is opaque reference to a CMS encoder object.
//
// See: https://developer.apple.com/documentation/Security/CMSEncoder
type CMSEncoderRef uintptr

// See: https://developer.apple.com/documentation/Security/CSSM_ACL_AUTHORIZATION_TAG
type CSSM_ACL_AUTHORIZATION_TAG = int32

// See: https://developer.apple.com/documentation/Security/CSSM_ACL_EDIT_MODE
type CSSM_ACL_EDIT_MODE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_ACL_HANDLE
type CSSM_ACL_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_ACL_KEYCHAIN_PROMPT_SELECTOR-swift.typealias
type CSSM_ACL_KEYCHAIN_PROMPT_SELECTOR = Cssm_acl_keychain_prompt_selector

// See: https://developer.apple.com/documentation/Security/CSSM_ACL_PREAUTH_TRACKING_STATE
type CSSM_ACL_PREAUTH_TRACKING_STATE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_ACL_PROCESS_SUBJECT_SELECTOR-swift.typealias
type CSSM_ACL_PROCESS_SUBJECT_SELECTOR = Cssm_acl_process_subject_selector

// See: https://developer.apple.com/documentation/Security/CSSM_ACL_SUBJECT_TYPE
type CSSM_ACL_SUBJECT_TYPE = int32

// See: https://developer.apple.com/documentation/Security/CSSM_AC_HANDLE
type CSSM_AC_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_ALGORITHMS
type CSSM_ALGORITHMS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS-swift.typealias
type CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS = Cssm_applecspdl_db_change_password_parameters

// See: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS_PTR
type CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS_PTR = uintptr

// See: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS-swift.typealias
type CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS = Cssm_applecspdl_db_is_locked_parameters

// See: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS_PTR
type CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS_PTR = uintptr

// See: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS-swift.typealias
type CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS = Cssm_applecspdl_db_settings_parameters

// See: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS_PTR
type CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS_PTR = uintptr

// See: https://developer.apple.com/documentation/Security/CSSM_APPLEDL_OPEN_PARAMETERS-swift.typealias
type CSSM_APPLEDL_OPEN_PARAMETERS = Cssm_appledl_open_parameters

// See: https://developer.apple.com/documentation/Security/CSSM_APPLEDL_OPEN_PARAMETERS_PTR
type CSSM_APPLEDL_OPEN_PARAMETERS_PTR = uintptr

// See: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_ACTION_FLAGS
type CSSM_APPLE_TP_ACTION_FLAGS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_CRL_OPT_FLAGS
type CSSM_APPLE_TP_CRL_OPT_FLAGS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_ATTACH_FLAGS
type CSSM_ATTACH_FLAGS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_ATTRIBUTE_TYPE
type CSSM_ATTRIBUTE_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_BER_TAG
type CSSM_BER_TAG = uint8

// See: https://developer.apple.com/documentation/Security/CSSM_BITMASK
type CSSM_BITMASK = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_BOOL
type CSSM_BOOL = int32

// See: https://developer.apple.com/documentation/Security/CSSM_CALLOC
type CSSM_CALLOC = func(uint, uint, unsafe.Pointer) unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_CC_HANDLE
type CSSM_CC_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_CERTGROUP_TYPE
type CSSM_CERTGROUP_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CERTGROUP_TYPE_PTR
type CSSM_CERTGROUP_TYPE_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CERT_BUNDLE_ENCODING
type CSSM_CERT_BUNDLE_ENCODING = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CERT_BUNDLE_TYPE
type CSSM_CERT_BUNDLE_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CERT_ENCODING
type CSSM_CERT_ENCODING = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CERT_ENCODING_PTR
type CSSM_CERT_ENCODING_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CERT_PARSE_FORMAT
type CSSM_CERT_PARSE_FORMAT = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CERT_PARSE_FORMAT_PTR
type CSSM_CERT_PARSE_FORMAT_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CERT_TYPE
type CSSM_CERT_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CERT_TYPE_PTR
type CSSM_CERT_TYPE_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CL_HANDLE
type CSSM_CL_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_CL_TEMPLATE_TYPE
type CSSM_CL_TEMPLATE_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CONTEXT_EVENT
type CSSM_CONTEXT_EVENT = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CONTEXT_TYPE
type CSSM_CONTEXT_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CRLGROUP_TYPE
type CSSM_CRLGROUP_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CRLGROUP_TYPE_PTR
type CSSM_CRLGROUP_TYPE_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CRL_ENCODING
type CSSM_CRL_ENCODING = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CRL_ENCODING_PTR
type CSSM_CRL_ENCODING_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CRL_PARSE_FORMAT
type CSSM_CRL_PARSE_FORMAT = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CRL_PARSE_FORMAT_PTR
type CSSM_CRL_PARSE_FORMAT_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CRL_TYPE
type CSSM_CRL_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CRL_TYPE_PTR
type CSSM_CRL_TYPE_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CSPTYPE
type CSSM_CSPTYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CSP_FLAGS
type CSSM_CSP_FLAGS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_CSP_HANDLE
type CSSM_CSP_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_CSP_READER_FLAGS
type CSSM_CSP_READER_FLAGS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_ACCESS_TYPE
type CSSM_DB_ACCESS_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_ACCESS_TYPE_PTR
type CSSM_DB_ACCESS_TYPE_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_FORMAT
type CSSM_DB_ATTRIBUTE_FORMAT = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_FORMAT_PTR
type CSSM_DB_ATTRIBUTE_FORMAT_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_NAME_FORMAT
type CSSM_DB_ATTRIBUTE_NAME_FORMAT = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR
type CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_CONJUNCTIVE
type CSSM_DB_CONJUNCTIVE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_CONJUNCTIVE_PTR
type CSSM_DB_CONJUNCTIVE_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_HANDLE
type CSSM_DB_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_DB_INDEXED_DATA_LOCATION
type CSSM_DB_INDEXED_DATA_LOCATION = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_INDEX_TYPE
type CSSM_DB_INDEX_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_MODIFY_MODE
type CSSM_DB_MODIFY_MODE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_OPERATOR
type CSSM_DB_OPERATOR = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_OPERATOR_PTR
type CSSM_DB_OPERATOR_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_RECORDTYPE
type CSSM_DB_RECORDTYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DB_RETRIEVAL_MODES
type CSSM_DB_RETRIEVAL_MODES = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DLTYPE
type CSSM_DLTYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DLTYPE_PTR
type CSSM_DLTYPE_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_DL_CUSTOM_ATTRIBUTES
type CSSM_DL_CUSTOM_ATTRIBUTES = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_DL_FFS_ATTRIBUTES
type CSSM_DL_FFS_ATTRIBUTES = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_DL_HANDLE
type CSSM_DL_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_DL_LDAP_ATTRIBUTES
type CSSM_DL_LDAP_ATTRIBUTES = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_DL_ODBC_ATTRIBUTES
type CSSM_DL_ODBC_ATTRIBUTES = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_DL_PKCS11_ATTRIBUTE
type CSSM_DL_PKCS11_ATTRIBUTE = uintptr

// See: https://developer.apple.com/documentation/Security/CSSM_DL_PKCS11_ATTRIBUTE_PTR
type CSSM_DL_PKCS11_ATTRIBUTE_PTR = uintptr

// See: https://developer.apple.com/documentation/Security/CSSM_ENCRYPT_MODE
type CSSM_ENCRYPT_MODE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_EVIDENCE_FORM
type CSSM_EVIDENCE_FORM = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_FREE
type CSSM_FREE = func(unsafe.Pointer, unsafe.Pointer)

// See: https://developer.apple.com/documentation/Security/CSSM_HANDLE
type CSSM_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_HANDLE_PTR
type CSSM_HANDLE_PTR = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_HEADERVERSION
type CSSM_HEADERVERSION = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_INTPTR
type CSSM_INTPTR = int

// See: https://developer.apple.com/documentation/Security/CSSM_KEYATTR_FLAGS
type CSSM_KEYATTR_FLAGS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_KEYBLOB_FORMAT
type CSSM_KEYBLOB_FORMAT = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_KEYBLOB_TYPE
type CSSM_KEYBLOB_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_KEYCLASS
type CSSM_KEYCLASS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_KEYUSE
type CSSM_KEYUSE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_KEY_HIERARCHY
type CSSM_KEY_HIERARCHY = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_KEY_TYPE
type CSSM_KEY_TYPE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_KRSP_HANDLE
type CSSM_KRSP_HANDLE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_KR_POLICY_FLAGS
type CSSM_KR_POLICY_FLAGS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_KR_POLICY_TYPE
type CSSM_KR_POLICY_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_LIST_ELEMENT_PTR
type CSSM_LIST_ELEMENT_PTR = uintptr

// See: https://developer.apple.com/documentation/Security/CSSM_LIST_ELEMENT_TYPE
type CSSM_LIST_ELEMENT_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_LIST_ELEMENT_TYPE_PTR
type CSSM_LIST_ELEMENT_TYPE_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_LIST_TYPE
type CSSM_LIST_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_LIST_TYPE_PTR
type CSSM_LIST_TYPE_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_LONG_HANDLE
type CSSM_LONG_HANDLE = uint64

// See: https://developer.apple.com/documentation/Security/CSSM_LONG_HANDLE_PTR
type CSSM_LONG_HANDLE_PTR = *uint64

// See: https://developer.apple.com/documentation/Security/CSSM_MALLOC
type CSSM_MALLOC = func(uint, unsafe.Pointer) unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_MANAGER_EVENT_TYPES
type CSSM_MANAGER_EVENT_TYPES = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_MODULE_EVENT
type CSSM_MODULE_EVENT = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_MODULE_EVENT_PTR
type CSSM_MODULE_EVENT_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_MODULE_HANDLE
type CSSM_MODULE_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_MODULE_HANDLE_PTR
type CSSM_MODULE_HANDLE_PTR = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_NET_ADDRESS_TYPE
type CSSM_NET_ADDRESS_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_NET_PROTOCOL
type CSSM_NET_PROTOCOL = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_PADDING
type CSSM_PADDING = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_PKCS5_PBKDF2_PRF
type CSSM_PKCS5_PBKDF2_PRF = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_PKCS_OAEP_MGF
type CSSM_PKCS_OAEP_MGF = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_PKCS_OAEP_PSOURCE
type CSSM_PKCS_OAEP_PSOURCE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_PRIVILEGE
type CSSM_PRIVILEGE = uint64

// See: https://developer.apple.com/documentation/Security/CSSM_PRIVILEGE_SCOPE
type CSSM_PRIVILEGE_SCOPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_PROC_ADDR
type CSSM_PROC_ADDR = func()

// See: https://developer.apple.com/documentation/Security/CSSM_PROC_ADDR_PTR
type CSSM_PROC_ADDR_PTR = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_PVC_MODE
type CSSM_PVC_MODE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_QUERY_FLAGS
type CSSM_QUERY_FLAGS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_REALLOC
type CSSM_REALLOC = func(unsafe.Pointer, uint, unsafe.Pointer) unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_RETURN
type CSSM_RETURN = int32

// See: https://developer.apple.com/documentation/Security/CSSM_SAMPLE_TYPE
type CSSM_SAMPLE_TYPE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_SC_FLAGS
type CSSM_SC_FLAGS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_SERVICE_MASK
type CSSM_SERVICE_MASK = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_SERVICE_TYPE
type CSSM_SERVICE_TYPE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_SIZE
type CSSM_SIZE = uintptr

// See: https://developer.apple.com/documentation/Security/CSSM_STRING
type CSSM_STRING = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_TIMESTRING
type CSSM_TIMESTRING = *byte

// See: https://developer.apple.com/documentation/Security/CSSM_TP_ACTION
type CSSM_TP_ACTION = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_APPLE_CERT_STATUS
type CSSM_TP_APPLE_CERT_STATUS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_AUTHORITY_REQUEST_TYPE
type CSSM_TP_AUTHORITY_REQUEST_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR
type CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_CERTCHANGE_ACTION
type CSSM_TP_CERTCHANGE_ACTION = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_CERTCHANGE_REASON
type CSSM_TP_CERTCHANGE_REASON = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_CERTCHANGE_STATUS
type CSSM_TP_CERTCHANGE_STATUS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_CERTISSUE_STATUS
type CSSM_TP_CERTISSUE_STATUS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_CERTNOTARIZE_STATUS
type CSSM_TP_CERTNOTARIZE_STATUS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_CERTRECLAIM_STATUS
type CSSM_TP_CERTRECLAIM_STATUS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_CERTVERIFY_STATUS
type CSSM_TP_CERTVERIFY_STATUS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_CONFIRM_STATUS
type CSSM_TP_CONFIRM_STATUS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_CONFIRM_STATUS_PTR
type CSSM_TP_CONFIRM_STATUS_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_CRLISSUE_STATUS
type CSSM_TP_CRLISSUE_STATUS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_FORM_TYPE
type CSSM_TP_FORM_TYPE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_HANDLE
type CSSM_TP_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_TP_SERVICES
type CSSM_TP_SERVICES = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_TP_STOP_ON
type CSSM_TP_STOP_ON = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_USEE_TAG
type CSSM_USEE_TAG = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_WORDID_TYPE
type CSSM_WORDID_TYPE = int32

// See: https://developer.apple.com/documentation/Security/CSSM_X509EXT_DATA_FORMAT
type CSSM_X509EXT_DATA_FORMAT = Extension_data_format

// See: https://developer.apple.com/documentation/Security/CSSM_X509_OPTION
type CSSM_X509_OPTION = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/MDS_HANDLE
type MDS_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/OpaqueSecAccessRef
type OpaqueSecAccessRef uintptr

// See: https://developer.apple.com/documentation/Security/OpaqueSecCertificateRef
type OpaqueSecCertificateRef uintptr

// See: https://developer.apple.com/documentation/Security/OpaqueSecIdentityRef
type OpaqueSecIdentityRef uintptr

// See: https://developer.apple.com/documentation/Security/OpaqueSecKeyRef
type OpaqueSecKeyRef = string

// SSLCipherSuite is a type for storing cipher suite values.
//
// See: https://developer.apple.com/documentation/Security/SSLCipherSuite
type SSLCipherSuite = uint16

// SSLConnectionRef is a pointer to an opaque I/O connection object.
//
// See: https://developer.apple.com/documentation/Security/SSLConnectionRef
type SSLConnectionRef uintptr

// SSLContextRef is an opaque type that represents an SSL session context object.
//
// See: https://developer.apple.com/documentation/Security/SSLContext
type SSLContextRef uintptr

// SSLReadFunc is a pointer to a customized read function that secure transport calls to read data from the connection.
//
// See: https://developer.apple.com/documentation/Security/SSLReadFunc
type SSLReadFunc = func(unsafe.Pointer, unsafe.Pointer, *uint) int

// SSLWriteFunc is a pointer to a customized write function that secure transport calls to write data to the connection.
//
// See: https://developer.apple.com/documentation/Security/SSLWriteFunc
type SSLWriteFunc = func(unsafe.Pointer, unsafe.Pointer, *uint) int

// SecACLRef is an opaque type that represents information about an ACL entry.
//
// See: https://developer.apple.com/documentation/Security/SecACL
type SecACLRef uintptr

// SecAccessControlRef is an opaque type that contains information about how a keychain item may be used.
//
// See: https://developer.apple.com/documentation/Security/SecAccessControl
type SecAccessControlRef uintptr

// SecAccessOwnerType is a type for flags that enable you to configure ACL ownership.
//
// See: https://developer.apple.com/documentation/Security/SecAccessOwnerType
type SecAccessOwnerType = uint32

// SecAccessRef is an opaque type that identifies a keychain item’s access information.
//
// See: https://developer.apple.com/documentation/Security/SecAccess
type SecAccessRef uintptr

// SecCertificateRef is an abstract Core Foundation-type object representing an X.509 certificate.
//
// See: https://developer.apple.com/documentation/Security/SecCertificate
type SecCertificateRef uintptr

// SecCodeRef is a code object representing signed code running on the system.
//
// See: https://developer.apple.com/documentation/Security/SecCode
type SecCodeRef uintptr

// SecGuestRef is a reference to a guest object, which identifies a particular block of guest code in the context of its code signing host.
//
// See: https://developer.apple.com/documentation/Security/SecGuestRef
type SecGuestRef = objectivec.IObject

// SecIdentityRef is an abstract Core Foundation-type object representing an identity.
//
// See: https://developer.apple.com/documentation/Security/SecIdentity
type SecIdentityRef uintptr

// SecIdentitySearchRef is contains information about an identity search.
//
// See: https://developer.apple.com/documentation/Security/SecIdentitySearch
type SecIdentitySearchRef uintptr

// SecKeyAlgorithm is the algorithms that cryptographic keys enable.
//
// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm
type SecKeyAlgorithm = corefoundation.CFStringRef

// SecKeyGeneratePairBlock is a block called with the results of a call to [SecKeyGeneratePairAsync(_:_:_:)].
//
// See: https://developer.apple.com/documentation/Security/SecKeyGeneratePairBlock
type SecKeyGeneratePairBlock = func(string, string, objectivec.IObject)

// SecKeyKeyExchangeParameter is the dictionary keys used to specify Diffie-Hellman key exchange parameters.
//
// See: https://developer.apple.com/documentation/Security/SecKeyKeyExchangeParameter
type SecKeyKeyExchangeParameter = corefoundation.CFStringRef

// SecKeyRef is an object that represents a cryptographic key.
//
// See: https://developer.apple.com/documentation/Security/SecKey
type SecKeyRef uintptr

// SecKeychainAttrType is the keychain attribute type.
//
// See: https://developer.apple.com/documentation/Security/SecKeychainAttrType
type SecKeychainAttrType = uint32

// SecKeychainAttributePtr is a pointer to a keychain attribute structure.
//
// See: https://developer.apple.com/documentation/Security/SecKeychainAttributePtr
type SecKeychainAttributePtr = *SecKeychainAttribute

// SecKeychainItemRef is an opaque type that represents a keychain item.
//
// See: https://developer.apple.com/documentation/Security/SecKeychainItem
type SecKeychainItemRef uintptr

// SecKeychainRef is an opaque type that represents a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecKeychain
type SecKeychainRef uintptr

// SecKeychainSearchRef is an opaque type that contains information about a keychain search.
//
// See: https://developer.apple.com/documentation/Security/SecKeychainSearch
type SecKeychainSearchRef uintptr

// SecKeychainStatus is a value that defines the current status of a keychain.
//
// See: https://developer.apple.com/documentation/Security/SecKeychainStatus
type SecKeychainStatus = uint32

// SecMessageBlock is a block that delivers messages during asynchronous operations.
//
// See: https://developer.apple.com/documentation/Security/SecMessageBlock
type SecMessageBlock = func(unsafe.Pointer, objectivec.IObject, uint32)

// SecPasswordRef is contains information about a password.
//
// See: https://developer.apple.com/documentation/Security/SecPassword
type SecPasswordRef uintptr

// SecPolicyRef is an object that represents a trust policy.
//
// See: https://developer.apple.com/documentation/Security/SecPolicy
type SecPolicyRef uintptr

// SecPolicySearchRef is an object that contains information about a policy search.
//
// See: https://developer.apple.com/documentation/Security/SecPolicySearch
type SecPolicySearchRef uintptr

// SecPublicKeyHash is a container for a 20-byte public key hash.
//
// See: https://developer.apple.com/documentation/Security/SecPublicKeyHash
type SecPublicKeyHash = uint8

// SecRandomRef is an abstract Core Foundation-type object containing information about a random number generator.
//
// See: https://developer.apple.com/documentation/Security/SecRandomRef
type SecRandomRef uintptr

// SecRequirementRef is a code requirement object.
//
// See: https://developer.apple.com/documentation/Security/SecRequirement
type SecRequirementRef uintptr

// SecStaticCodeRef is a static code object representing signed code on disk.
//
// See: https://developer.apple.com/documentation/Security/SecStaticCode
type SecStaticCodeRef uintptr

// SecTaskRef is the Core Foundation type representing a task.
//
// See: https://developer.apple.com/documentation/Security/SecTask
type SecTaskRef uintptr

// SecTransformDataBlock is a block used to override the default data handling for a transform.
//
// See: https://developer.apple.com/documentation/Security/SecTransformDataBlock
type SecTransformDataBlock = func(unsafe.Pointer) unsafe.Pointer

// SecTransformImplementationRef is an opaque pointer to a block that implements an instance of a transform.
//
// See: https://developer.apple.com/documentation/Security/SecTransformImplementationRef
type SecTransformImplementationRef uintptr

// SecTransformInstanceBlock is a block that you return from a transform creation function.
//
// See: https://developer.apple.com/documentation/Security/SecTransformInstanceBlock
type SecTransformInstanceBlock = func() objectivec.IObject

// SecTrustCallback is a block called with the results of an asynchronous trust evaluation.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCallback
type SecTrustCallback = func(objectivec.IObject, SecTrustResultType)

// SecTrustRef is an object used to evaluate trust.
//
// See: https://developer.apple.com/documentation/Security/SecTrust
type SecTrustRef uintptr

// SecTrustWithErrorCallback is a block called with the results of an asynchronous trust evaluation.
//
// See: https://developer.apple.com/documentation/Security/SecTrustWithErrorCallback
type SecTrustWithErrorCallback = func(objectivec.IObject, bool, objectivec.IObject)

// SecTrustedApplicationRef is an opaque type that contains information about a trusted app.
//
// See: https://developer.apple.com/documentation/Security/SecTrustedApplication
type SecTrustedApplicationRef uintptr

// SecureDownloadRef is an opaque type representing a secure download object.
//
// See: https://developer.apple.com/documentation/Security/SecureDownloadRef
type SecureDownloadRef uintptr

// SecuritySessionId is a type that contains an authorization session identifier.
//
// See: https://developer.apple.com/documentation/Security/SecuritySessionId
type SecuritySessionId = uint32

// See: https://developer.apple.com/documentation/Security/sec_certificate_t
type Sec_certificate_t = objectivec.Object

// See: https://developer.apple.com/documentation/Security/sec_identity_t
type Sec_identity_t = objectivec.Object

// Sec_object_t is a `sec_object` is a generic, ARC-able type wrapper for common CoreFoundation Security types.
//
// See: https://developer.apple.com/documentation/Security/sec_object_t
type Sec_object_t = objectivec.Object

// See: https://developer.apple.com/documentation/Security/sec_protocol_challenge_complete_t
type Sec_protocol_challenge_complete_t = func(objectivec.Object)

// See: https://developer.apple.com/documentation/Security/sec_protocol_challenge_t
type Sec_protocol_challenge_t = func(*objectivec.Object)

// See: https://developer.apple.com/documentation/Security/sec_protocol_key_update_complete_t
type Sec_protocol_key_update_complete_t = func()

// See: https://developer.apple.com/documentation/Security/sec_protocol_key_update_t
type Sec_protocol_key_update_t = func(*objectivec.Object, func())

// Sec_protocol_metadata_t is a `sec_protocol_metadata` instance conatins read-only properties of a connected and configured security protocol. Clients use this object to read information about a protocol instance. Properties include, for example, the negotiated TLS version, ciphersuite, and peer certificates.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_t
type Sec_protocol_metadata_t = objectivec.Object

// Sec_protocol_options_t is a `sec_protocol_options` instance is a container of options for security protocol instances, such as TLS. Protocol options are used to configure security protocols in the network stack. For example, clients may set the maximum and minimum allowed TLS versions through protocol options.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_t
type Sec_protocol_options_t = objectivec.Object

// See: https://developer.apple.com/documentation/Security/sec_protocol_pre_shared_key_selection_complete_t
type Sec_protocol_pre_shared_key_selection_complete_t = func(objectivec.Object)

// See: https://developer.apple.com/documentation/Security/sec_protocol_pre_shared_key_selection_t
type Sec_protocol_pre_shared_key_selection_t = func(*objectivec.Object)

// See: https://developer.apple.com/documentation/Security/sec_protocol_verify_complete_t
type Sec_protocol_verify_complete_t = func(bool)

// See: https://developer.apple.com/documentation/Security/sec_protocol_verify_t
type Sec_protocol_verify_t = func(*objectivec.Object, func(unsafe.Pointer))

// Sec_trust_t is these are os_object compatible and ARC-able wrappers around existing CoreFoundation Security types, including: SecTrustRef, SecIdentityRef, and SecCertificateRef. They allow clients to use these types in os_object-type APIs and data structures. The underlying CoreFoundation types may be extracted and used by clients as needed.
//
// See: https://developer.apple.com/documentation/Security/sec_trust_t
type Sec_trust_t = objectivec.Object

// See: https://developer.apple.com/documentation/Security/sint16
type Sint16 = int16

// See: https://developer.apple.com/documentation/Security/sint32
type Sint32 = int32

// See: https://developer.apple.com/documentation/Security/sint64
type Sint64 = int64

// See: https://developer.apple.com/documentation/Security/sint8
type Sint8 = int8

