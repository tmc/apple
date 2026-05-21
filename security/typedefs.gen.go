// Code generated from Apple documentation. DO NOT EDIT.

package security

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
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
type AuthorizationMechanismId = kernel.Pointer

// AuthorizationMechanismRef is a handle passed by the plug-in to the authorization engine when creating an instance of a mechanism.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationMechanismRef
type AuthorizationMechanismRef = kernel.Pointer

// AuthorizationPluginId is an unused identifier for a plug-in.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationPluginId
type AuthorizationPluginId = kernel.Pointer

// AuthorizationPluginRef is a handle passed by the plug-in to the authorization engine when the plug-in is initiated.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationPluginRef
type AuthorizationPluginRef = kernel.Pointer

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
type AuthorizationSessionId = kernel.Pointer

// AuthorizationString is a zero-terminated string in UTF-8 encoding.
//
// See: https://developer.apple.com/documentation/Security/AuthorizationString
type AuthorizationString = *byte

// See: https://developer.apple.com/documentation/Security/CE_CrlNumber
type CE_CrlNumber = uint32

// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CE_CrlReason
type CE_CrlReason = uint32

// See: https://developer.apple.com/documentation/Security/CE_DeltaCrl
type CE_DeltaCrl = uint32

// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CE_InhibitAnyPolicy
type CE_InhibitAnyPolicy = uint32

// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CE_KeyUsage
type CE_KeyUsage = uint16

// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CE_NetscapeCertType
type CE_NetscapeCertType = uint16

// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/Security/CE_SubjectKeyID
type CE_SubjectKeyID = string

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
type CSSM_CALLOC = func(uint, uint, kernel.Pointer) kernel.Pointer

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
type CSSM_DL_CUSTOM_ATTRIBUTES = kernel.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_DL_FFS_ATTRIBUTES
type CSSM_DL_FFS_ATTRIBUTES = kernel.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_DL_HANDLE
type CSSM_DL_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_DL_LDAP_ATTRIBUTES
type CSSM_DL_LDAP_ATTRIBUTES = kernel.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_DL_ODBC_ATTRIBUTES
type CSSM_DL_ODBC_ATTRIBUTES = kernel.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_DL_PKCS11_ATTRIBUTE
type CSSM_DL_PKCS11_ATTRIBUTE = uintptr

// See: https://developer.apple.com/documentation/Security/CSSM_DL_PKCS11_ATTRIBUTE_PTR
type CSSM_DL_PKCS11_ATTRIBUTE_PTR = uintptr

// See: https://developer.apple.com/documentation/Security/CSSM_ENCRYPT_MODE
type CSSM_ENCRYPT_MODE = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_EVIDENCE_FORM
type CSSM_EVIDENCE_FORM = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_FREE
type CSSM_FREE = func(kernel.Pointer, kernel.Pointer)

// See: https://developer.apple.com/documentation/Security/CSSM_HANDLE
type CSSM_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_HANDLE_PTR
type CSSM_HANDLE_PTR = kernel.Pointer

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
type CSSM_MALLOC = func(uint, kernel.Pointer) kernel.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_MANAGER_EVENT_TYPES
type CSSM_MANAGER_EVENT_TYPES = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_MODULE_EVENT
type CSSM_MODULE_EVENT = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_MODULE_EVENT_PTR
type CSSM_MODULE_EVENT_PTR = *uint32

// See: https://developer.apple.com/documentation/Security/CSSM_MODULE_HANDLE
type CSSM_MODULE_HANDLE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_MODULE_HANDLE_PTR
type CSSM_MODULE_HANDLE_PTR = kernel.Pointer

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
type CSSM_PROC_ADDR_PTR = kernel.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_PVC_MODE
type CSSM_PVC_MODE = unsafe.Pointer

// See: https://developer.apple.com/documentation/Security/CSSM_QUERY_FLAGS
type CSSM_QUERY_FLAGS = uint32

// See: https://developer.apple.com/documentation/Security/CSSM_REALLOC
type CSSM_REALLOC = func(kernel.Pointer, uint, kernel.Pointer) kernel.Pointer

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
type CSSM_STRING = kernel.Pointer

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
type OpaqueSecAccessRef = kernel.Pointer

// See: https://developer.apple.com/documentation/Security/OpaqueSecCertificateRef
type OpaqueSecCertificateRef = kernel.Pointer

// See: https://developer.apple.com/documentation/Security/OpaqueSecIdentityRef
type OpaqueSecIdentityRef = kernel.Pointer

// See: https://developer.apple.com/documentation/Security/OpaqueSecKeyRef
type OpaqueSecKeyRef = string

// SSLCipherSuite is a type for storing cipher suite values.
//
// See: https://developer.apple.com/documentation/Security/SSLCipherSuite
type SSLCipherSuite = uint16

// SSLConnectionRef is a pointer to an opaque I/O connection object.
//
// See: https://developer.apple.com/documentation/Security/SSLConnectionRef
type SSLConnectionRef = kernel.Pointer

// SSLContextRef is an opaque type that represents an SSL session context object.
//
// See: https://developer.apple.com/documentation/Security/SSLContext
type SSLContextRef uintptr

// SSLReadFunc is a pointer to a customized read function that secure transport calls to read data from the connection.
//
// See: https://developer.apple.com/documentation/Security/SSLReadFunc
type SSLReadFunc = func(kernel.Pointer, kernel.Pointer, uint) int

// SSLWriteFunc is a pointer to a customized write function that secure transport calls to write data to the connection.
//
// See: https://developer.apple.com/documentation/Security/SSLWriteFunc
type SSLWriteFunc = func(kernel.Pointer, kernel.Pointer, uint) int

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

// SecAsn1Item is a structure holding DER encoded data.
//
// Deprecated: Deprecated since macOS 12.0. SecAsn1 is not supported
//
// See: https://developer.apple.com/documentation/Security/SecAsn1Item
type SecAsn1Item = Cssm_data

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
type SecGuestRef = kernel.U_int32_t

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
type SecKeyGeneratePairBlock = func(string, string, kernel.Pointer)

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
type SecMessageBlock = func(kernel.Pointer, kernel.Pointer, uint32)

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
type SecTransformDataBlock = func(kernel.Pointer) kernel.Pointer

// SecTransformImplementationRef is an opaque pointer to a block that implements an instance of a transform.
//
// See: https://developer.apple.com/documentation/Security/SecTransformImplementationRef
type SecTransformImplementationRef uintptr

// SecTransformInstanceBlock is a block that you return from a transform creation function.
//
// See: https://developer.apple.com/documentation/Security/SecTransformInstanceBlock
type SecTransformInstanceBlock = func() kernel.Pointer

// SecTrustCallback is a block called with the results of an asynchronous trust evaluation.
//
// See: https://developer.apple.com/documentation/Security/SecTrustCallback
type SecTrustCallback = func(kernel.Pointer, SecTrustResultType)

// SecTrustRef is an object used to evaluate trust.
//
// See: https://developer.apple.com/documentation/Security/SecTrust
type SecTrustRef uintptr

// SecTrustWithErrorCallback is a block called with the results of an asynchronous trust evaluation.
//
// See: https://developer.apple.com/documentation/Security/SecTrustWithErrorCallback
type SecTrustWithErrorCallback = func(kernel.Pointer, bool, kernel.Pointer)

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

// Sec_certificate_tFromID constructs a [Sec_certificate_t] from an objc.ID.
func Sec_certificate_tFromID(id objc.ID) Sec_certificate_t {
	return Sec_certificate_t{ID: id}
}

// See: https://developer.apple.com/documentation/Security/sec_identity_t
type Sec_identity_t = objectivec.Object

// Sec_identity_tFromID constructs a [Sec_identity_t] from an objc.ID.
func Sec_identity_tFromID(id objc.ID) Sec_identity_t {
	return Sec_identity_t{ID: id}
}

// Sec_object_t is a `sec_object` is a generic, ARC-able type wrapper for common CoreFoundation Security types.
//
// See: https://developer.apple.com/documentation/Security/sec_object_t
type Sec_object_t = objectivec.Object

// Sec_object_tFromID constructs a [Sec_object_t] from an objc.ID.
func Sec_object_tFromID(id objc.ID) Sec_object_t {
	return Sec_object_t{ID: id}
}

// See: https://developer.apple.com/documentation/Security/sec_protocol_challenge_complete_t
type Sec_protocol_challenge_complete_t = func(objectivec.Object)

// See: https://developer.apple.com/documentation/Security/sec_protocol_challenge_t
type Sec_protocol_challenge_t = func(objectivec.Object, func(*objectivec.Object))

// See: https://developer.apple.com/documentation/Security/sec_protocol_key_update_complete_t
type Sec_protocol_key_update_complete_t = func()

// See: https://developer.apple.com/documentation/Security/sec_protocol_key_update_t
type Sec_protocol_key_update_t = func(objectivec.Object, func())

// Sec_protocol_metadata_t is a `sec_protocol_metadata` instance conatins read-only properties of a connected and configured security protocol. Clients use this object to read information about a protocol instance. Properties include, for example, the negotiated TLS version, ciphersuite, and peer certificates.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_metadata_t
type Sec_protocol_metadata_t = objectivec.Object

// Sec_protocol_metadata_tFromID constructs a [Sec_protocol_metadata_t] from an objc.ID.
func Sec_protocol_metadata_tFromID(id objc.ID) Sec_protocol_metadata_t {
	return Sec_protocol_metadata_t{ID: id}
}

// Sec_protocol_options_t is a `sec_protocol_options` instance is a container of options for security protocol instances, such as TLS. Protocol options are used to configure security protocols in the network stack. For example, clients may set the maximum and minimum allowed TLS versions through protocol options.
//
// See: https://developer.apple.com/documentation/Security/sec_protocol_options_t
type Sec_protocol_options_t = objectivec.Object

// Sec_protocol_options_tFromID constructs a [Sec_protocol_options_t] from an objc.ID.
func Sec_protocol_options_tFromID(id objc.ID) Sec_protocol_options_t {
	return Sec_protocol_options_t{ID: id}
}

// See: https://developer.apple.com/documentation/Security/sec_protocol_pre_shared_key_selection_complete_t
type Sec_protocol_pre_shared_key_selection_complete_t = func(objectivec.Object)

// See: https://developer.apple.com/documentation/Security/sec_protocol_pre_shared_key_selection_t
type Sec_protocol_pre_shared_key_selection_t = func(objectivec.Object, objectivec.Object, func(*objectivec.Object))

// See: https://developer.apple.com/documentation/Security/sec_protocol_verify_complete_t
type Sec_protocol_verify_complete_t = func(bool)

// See: https://developer.apple.com/documentation/Security/sec_protocol_verify_t
type Sec_protocol_verify_t = func(objectivec.Object, objectivec.Object, func(bool))

// Sec_trust_t is these are os_object compatible and ARC-able wrappers around existing CoreFoundation Security types, including: SecTrustRef, SecIdentityRef, and SecCertificateRef. They allow clients to use these types in os_object-type APIs and data structures. The underlying CoreFoundation types may be extracted and used by clients as needed.
//
// See: https://developer.apple.com/documentation/Security/sec_trust_t
type Sec_trust_t = objectivec.Object

// Sec_trust_tFromID constructs a [Sec_trust_t] from an objc.ID.
func Sec_trust_tFromID(id objc.ID) Sec_trust_t {
	return Sec_trust_t{ID: id}
}

// See: https://developer.apple.com/documentation/Security/sint16
type Sint16 = int16

// See: https://developer.apple.com/documentation/Security/sint32
type Sint32 = int32

// See: https://developer.apple.com/documentation/Security/sint64
type Sint64 = int64

// See: https://developer.apple.com/documentation/Security/sint8
type Sint8 = int8

// CeCrlNumber is a Go-name alias for CE_CrlNumber.
type CeCrlNumber = CE_CrlNumber

// CeCrlReason is a Go-name alias for CE_CrlReason.
type CeCrlReason = CE_CrlReason

// CeDeltaCrl is a Go-name alias for CE_DeltaCrl.
type CeDeltaCrl = CE_DeltaCrl

// CeInhibitAnyPolicy is a Go-name alias for CE_InhibitAnyPolicy.
type CeInhibitAnyPolicy = CE_InhibitAnyPolicy

// CeKeyUsage is a Go-name alias for CE_KeyUsage.
type CeKeyUsage = CE_KeyUsage

// CeNetscapeCertType is a Go-name alias for CE_NetscapeCertType.
type CeNetscapeCertType = CE_NetscapeCertType

// CeSubjectKeyID is a Go-name alias for CE_SubjectKeyID.
type CeSubjectKeyID = CE_SubjectKeyID

// CssmAclAuthorizationTag is a Go-name alias for CSSM_ACL_AUTHORIZATION_TAG.
type CssmAclAuthorizationTag = CSSM_ACL_AUTHORIZATION_TAG

// CssmAclHandle is a Go-name alias for CSSM_ACL_HANDLE.
type CssmAclHandle = CSSM_ACL_HANDLE

// CssmAclKeychainPromptSelector is a Go-name alias for CSSM_ACL_KEYCHAIN_PROMPT_SELECTOR.
type CssmAclKeychainPromptSelector = CSSM_ACL_KEYCHAIN_PROMPT_SELECTOR

// CssmAclPreauthTrackingState is a Go-name alias for CSSM_ACL_PREAUTH_TRACKING_STATE.
type CssmAclPreauthTrackingState = CSSM_ACL_PREAUTH_TRACKING_STATE

// CssmAclProcessSubjectSelector is a Go-name alias for CSSM_ACL_PROCESS_SUBJECT_SELECTOR.
type CssmAclProcessSubjectSelector = CSSM_ACL_PROCESS_SUBJECT_SELECTOR

// CssmAcHandle is a Go-name alias for CSSM_AC_HANDLE.
type CssmAcHandle = CSSM_AC_HANDLE

// CssmAlgorithms is a Go-name alias for CSSM_ALGORITHMS.
type CssmAlgorithms = CSSM_ALGORITHMS

// CssmApplecspdlDbChangePasswordParameters is a Go-name alias for CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS.
type CssmApplecspdlDbChangePasswordParameters = CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS

// CssmApplecspdlDbChangePasswordParametersPtr is a Go-name alias for CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS_PTR.
type CssmApplecspdlDbChangePasswordParametersPtr = CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS_PTR

// CssmApplecspdlDbIsLockedParameters is a Go-name alias for CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS.
type CssmApplecspdlDbIsLockedParameters = CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS

// CssmApplecspdlDbIsLockedParametersPtr is a Go-name alias for CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS_PTR.
type CssmApplecspdlDbIsLockedParametersPtr = CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS_PTR

// CssmApplecspdlDbSettingsParameters is a Go-name alias for CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS.
type CssmApplecspdlDbSettingsParameters = CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS

// CssmApplecspdlDbSettingsParametersPtr is a Go-name alias for CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS_PTR.
type CssmApplecspdlDbSettingsParametersPtr = CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS_PTR

// CssmAppledlOpenParametersPtr is a Go-name alias for CSSM_APPLEDL_OPEN_PARAMETERS_PTR.
type CssmAppledlOpenParametersPtr = CSSM_APPLEDL_OPEN_PARAMETERS_PTR

// CssmAppleTpActionFlags is a Go-name alias for CSSM_APPLE_TP_ACTION_FLAGS.
type CssmAppleTpActionFlags = CSSM_APPLE_TP_ACTION_FLAGS

// CssmAppleTpCrlOptFlags is a Go-name alias for CSSM_APPLE_TP_CRL_OPT_FLAGS.
type CssmAppleTpCrlOptFlags = CSSM_APPLE_TP_CRL_OPT_FLAGS

// CssmAttachFlags is a Go-name alias for CSSM_ATTACH_FLAGS.
type CssmAttachFlags = CSSM_ATTACH_FLAGS

// CssmAttributeType is a Go-name alias for CSSM_ATTRIBUTE_TYPE.
type CssmAttributeType = CSSM_ATTRIBUTE_TYPE

// CssmBerTag is a Go-name alias for CSSM_BER_TAG.
type CssmBerTag = CSSM_BER_TAG

// CssmBitmask is a Go-name alias for CSSM_BITMASK.
type CssmBitmask = CSSM_BITMASK

// CssmBool is a Go-name alias for CSSM_BOOL.
type CssmBool = CSSM_BOOL

// CssmCalloc is a Go-name alias for CSSM_CALLOC.
type CssmCalloc = CSSM_CALLOC

// CssmCcHandle is a Go-name alias for CSSM_CC_HANDLE.
type CssmCcHandle = CSSM_CC_HANDLE

// CssmCertgroupType is a Go-name alias for CSSM_CERTGROUP_TYPE.
type CssmCertgroupType = CSSM_CERTGROUP_TYPE

// CssmCertgroupTypePtr is a Go-name alias for CSSM_CERTGROUP_TYPE_PTR.
type CssmCertgroupTypePtr = CSSM_CERTGROUP_TYPE_PTR

// CssmCertBundleType is a Go-name alias for CSSM_CERT_BUNDLE_TYPE.
type CssmCertBundleType = CSSM_CERT_BUNDLE_TYPE

// CssmCertEncoding is a Go-name alias for CSSM_CERT_ENCODING.
type CssmCertEncoding = CSSM_CERT_ENCODING

// CssmCertEncodingPtr is a Go-name alias for CSSM_CERT_ENCODING_PTR.
type CssmCertEncodingPtr = CSSM_CERT_ENCODING_PTR

// CssmCertParseFormat is a Go-name alias for CSSM_CERT_PARSE_FORMAT.
type CssmCertParseFormat = CSSM_CERT_PARSE_FORMAT

// CssmCertParseFormatPtr is a Go-name alias for CSSM_CERT_PARSE_FORMAT_PTR.
type CssmCertParseFormatPtr = CSSM_CERT_PARSE_FORMAT_PTR

// CssmCertType is a Go-name alias for CSSM_CERT_TYPE.
type CssmCertType = CSSM_CERT_TYPE

// CssmCertTypePtr is a Go-name alias for CSSM_CERT_TYPE_PTR.
type CssmCertTypePtr = CSSM_CERT_TYPE_PTR

// CssmClHandle is a Go-name alias for CSSM_CL_HANDLE.
type CssmClHandle = CSSM_CL_HANDLE

// CssmClTemplateType is a Go-name alias for CSSM_CL_TEMPLATE_TYPE.
type CssmClTemplateType = CSSM_CL_TEMPLATE_TYPE

// CssmContextType is a Go-name alias for CSSM_CONTEXT_TYPE.
type CssmContextType = CSSM_CONTEXT_TYPE

// CssmCrlgroupType is a Go-name alias for CSSM_CRLGROUP_TYPE.
type CssmCrlgroupType = CSSM_CRLGROUP_TYPE

// CssmCrlgroupTypePtr is a Go-name alias for CSSM_CRLGROUP_TYPE_PTR.
type CssmCrlgroupTypePtr = CSSM_CRLGROUP_TYPE_PTR

// CssmCrlEncodingPtr is a Go-name alias for CSSM_CRL_ENCODING_PTR.
type CssmCrlEncodingPtr = CSSM_CRL_ENCODING_PTR

// CssmCrlParseFormat is a Go-name alias for CSSM_CRL_PARSE_FORMAT.
type CssmCrlParseFormat = CSSM_CRL_PARSE_FORMAT

// CssmCrlParseFormatPtr is a Go-name alias for CSSM_CRL_PARSE_FORMAT_PTR.
type CssmCrlParseFormatPtr = CSSM_CRL_PARSE_FORMAT_PTR

// CssmCrlTypePtr is a Go-name alias for CSSM_CRL_TYPE_PTR.
type CssmCrlTypePtr = CSSM_CRL_TYPE_PTR

// CssmCsptype is a Go-name alias for CSSM_CSPTYPE.
type CssmCsptype = CSSM_CSPTYPE

// CssmCspFlags is a Go-name alias for CSSM_CSP_FLAGS.
type CssmCspFlags = CSSM_CSP_FLAGS

// CssmCspHandle is a Go-name alias for CSSM_CSP_HANDLE.
type CssmCspHandle = CSSM_CSP_HANDLE

// CssmCspReaderFlags is a Go-name alias for CSSM_CSP_READER_FLAGS.
type CssmCspReaderFlags = CSSM_CSP_READER_FLAGS

// CssmDbAccessType is a Go-name alias for CSSM_DB_ACCESS_TYPE.
type CssmDbAccessType = CSSM_DB_ACCESS_TYPE

// CssmDbAccessTypePtr is a Go-name alias for CSSM_DB_ACCESS_TYPE_PTR.
type CssmDbAccessTypePtr = CSSM_DB_ACCESS_TYPE_PTR

// CssmDbAttributeFormatPtr is a Go-name alias for CSSM_DB_ATTRIBUTE_FORMAT_PTR.
type CssmDbAttributeFormatPtr = CSSM_DB_ATTRIBUTE_FORMAT_PTR

// CssmDbAttributeNameFormat is a Go-name alias for CSSM_DB_ATTRIBUTE_NAME_FORMAT.
type CssmDbAttributeNameFormat = CSSM_DB_ATTRIBUTE_NAME_FORMAT

// CssmDbAttributeNameFormatPtr is a Go-name alias for CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR.
type CssmDbAttributeNameFormatPtr = CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR

// CssmDbConjunctive is a Go-name alias for CSSM_DB_CONJUNCTIVE.
type CssmDbConjunctive = CSSM_DB_CONJUNCTIVE

// CssmDbConjunctivePtr is a Go-name alias for CSSM_DB_CONJUNCTIVE_PTR.
type CssmDbConjunctivePtr = CSSM_DB_CONJUNCTIVE_PTR

// CssmDbHandle is a Go-name alias for CSSM_DB_HANDLE.
type CssmDbHandle = CSSM_DB_HANDLE

// CssmDbIndexedDataLocation is a Go-name alias for CSSM_DB_INDEXED_DATA_LOCATION.
type CssmDbIndexedDataLocation = CSSM_DB_INDEXED_DATA_LOCATION

// CssmDbIndexType is a Go-name alias for CSSM_DB_INDEX_TYPE.
type CssmDbIndexType = CSSM_DB_INDEX_TYPE

// CssmDbModifyMode is a Go-name alias for CSSM_DB_MODIFY_MODE.
type CssmDbModifyMode = CSSM_DB_MODIFY_MODE

// CssmDbOperator is a Go-name alias for CSSM_DB_OPERATOR.
type CssmDbOperator = CSSM_DB_OPERATOR

// CssmDbOperatorPtr is a Go-name alias for CSSM_DB_OPERATOR_PTR.
type CssmDbOperatorPtr = CSSM_DB_OPERATOR_PTR

// CssmDbRecordtype is a Go-name alias for CSSM_DB_RECORDTYPE.
type CssmDbRecordtype = CSSM_DB_RECORDTYPE

// CssmDbRetrievalModes is a Go-name alias for CSSM_DB_RETRIEVAL_MODES.
type CssmDbRetrievalModes = CSSM_DB_RETRIEVAL_MODES

// CssmDltype is a Go-name alias for CSSM_DLTYPE.
type CssmDltype = CSSM_DLTYPE

// CssmDltypePtr is a Go-name alias for CSSM_DLTYPE_PTR.
type CssmDltypePtr = CSSM_DLTYPE_PTR

// CssmDlCustomAttributes is a Go-name alias for CSSM_DL_CUSTOM_ATTRIBUTES.
type CssmDlCustomAttributes = CSSM_DL_CUSTOM_ATTRIBUTES

// CssmDlFfsAttributes is a Go-name alias for CSSM_DL_FFS_ATTRIBUTES.
type CssmDlFfsAttributes = CSSM_DL_FFS_ATTRIBUTES

// CssmDlHandle is a Go-name alias for CSSM_DL_HANDLE.
type CssmDlHandle = CSSM_DL_HANDLE

// CssmDlLdapAttributes is a Go-name alias for CSSM_DL_LDAP_ATTRIBUTES.
type CssmDlLdapAttributes = CSSM_DL_LDAP_ATTRIBUTES

// CssmDlOdbcAttributes is a Go-name alias for CSSM_DL_ODBC_ATTRIBUTES.
type CssmDlOdbcAttributes = CSSM_DL_ODBC_ATTRIBUTES

// CssmDlPkcs11Attribute is a Go-name alias for CSSM_DL_PKCS11_ATTRIBUTE.
type CssmDlPkcs11Attribute = CSSM_DL_PKCS11_ATTRIBUTE

// CssmDlPkcs11AttributePtr is a Go-name alias for CSSM_DL_PKCS11_ATTRIBUTE_PTR.
type CssmDlPkcs11AttributePtr = CSSM_DL_PKCS11_ATTRIBUTE_PTR

// CssmEncryptMode is a Go-name alias for CSSM_ENCRYPT_MODE.
type CssmEncryptMode = CSSM_ENCRYPT_MODE

// CssmFree is a Go-name alias for CSSM_FREE.
type CssmFree = CSSM_FREE

// CssmHandle is a Go-name alias for CSSM_HANDLE.
type CssmHandle = CSSM_HANDLE

// CssmHandlePtr is a Go-name alias for CSSM_HANDLE_PTR.
type CssmHandlePtr = CSSM_HANDLE_PTR

// CssmHeaderversion is a Go-name alias for CSSM_HEADERVERSION.
type CssmHeaderversion = CSSM_HEADERVERSION

// CssmIntptr is a Go-name alias for CSSM_INTPTR.
type CssmIntptr = CSSM_INTPTR

// CssmKeyattrFlags is a Go-name alias for CSSM_KEYATTR_FLAGS.
type CssmKeyattrFlags = CSSM_KEYATTR_FLAGS

// CssmKeyblobFormat is a Go-name alias for CSSM_KEYBLOB_FORMAT.
type CssmKeyblobFormat = CSSM_KEYBLOB_FORMAT

// CssmKeyblobType is a Go-name alias for CSSM_KEYBLOB_TYPE.
type CssmKeyblobType = CSSM_KEYBLOB_TYPE

// CssmKeyType is a Go-name alias for CSSM_KEY_TYPE.
type CssmKeyType = CSSM_KEY_TYPE

// CssmKrspHandle is a Go-name alias for CSSM_KRSP_HANDLE.
type CssmKrspHandle = CSSM_KRSP_HANDLE

// CssmKrPolicyFlags is a Go-name alias for CSSM_KR_POLICY_FLAGS.
type CssmKrPolicyFlags = CSSM_KR_POLICY_FLAGS

// CssmKrPolicyType is a Go-name alias for CSSM_KR_POLICY_TYPE.
type CssmKrPolicyType = CSSM_KR_POLICY_TYPE

// CssmListElementPtr is a Go-name alias for CSSM_LIST_ELEMENT_PTR.
type CssmListElementPtr = CSSM_LIST_ELEMENT_PTR

// CssmListElementType is a Go-name alias for CSSM_LIST_ELEMENT_TYPE.
type CssmListElementType = CSSM_LIST_ELEMENT_TYPE

// CssmListElementTypePtr is a Go-name alias for CSSM_LIST_ELEMENT_TYPE_PTR.
type CssmListElementTypePtr = CSSM_LIST_ELEMENT_TYPE_PTR

// CssmListTypePtr is a Go-name alias for CSSM_LIST_TYPE_PTR.
type CssmListTypePtr = CSSM_LIST_TYPE_PTR

// CssmLongHandle is a Go-name alias for CSSM_LONG_HANDLE.
type CssmLongHandle = CSSM_LONG_HANDLE

// CssmLongHandlePtr is a Go-name alias for CSSM_LONG_HANDLE_PTR.
type CssmLongHandlePtr = CSSM_LONG_HANDLE_PTR

// CssmMalloc is a Go-name alias for CSSM_MALLOC.
type CssmMalloc = CSSM_MALLOC

// CssmManagerEventTypes is a Go-name alias for CSSM_MANAGER_EVENT_TYPES.
type CssmManagerEventTypes = CSSM_MANAGER_EVENT_TYPES

// CssmModuleEvent is a Go-name alias for CSSM_MODULE_EVENT.
type CssmModuleEvent = CSSM_MODULE_EVENT

// CssmModuleEventPtr is a Go-name alias for CSSM_MODULE_EVENT_PTR.
type CssmModuleEventPtr = CSSM_MODULE_EVENT_PTR

// CssmModuleHandle is a Go-name alias for CSSM_MODULE_HANDLE.
type CssmModuleHandle = CSSM_MODULE_HANDLE

// CssmModuleHandlePtr is a Go-name alias for CSSM_MODULE_HANDLE_PTR.
type CssmModuleHandlePtr = CSSM_MODULE_HANDLE_PTR

// CssmNetAddressType is a Go-name alias for CSSM_NET_ADDRESS_TYPE.
type CssmNetAddressType = CSSM_NET_ADDRESS_TYPE

// CssmNetProtocol is a Go-name alias for CSSM_NET_PROTOCOL.
type CssmNetProtocol = CSSM_NET_PROTOCOL

// CssmPkcs5Pbkdf2Prf is a Go-name alias for CSSM_PKCS5_PBKDF2_PRF.
type CssmPkcs5Pbkdf2Prf = CSSM_PKCS5_PBKDF2_PRF

// CssmPkcsOaepMgf is a Go-name alias for CSSM_PKCS_OAEP_MGF.
type CssmPkcsOaepMgf = CSSM_PKCS_OAEP_MGF

// CssmPrivilege is a Go-name alias for CSSM_PRIVILEGE.
type CssmPrivilege = CSSM_PRIVILEGE

// CssmProcAddr is a Go-name alias for CSSM_PROC_ADDR.
type CssmProcAddr = CSSM_PROC_ADDR

// CssmProcAddrPtr is a Go-name alias for CSSM_PROC_ADDR_PTR.
type CssmProcAddrPtr = CSSM_PROC_ADDR_PTR

// CssmPvcMode is a Go-name alias for CSSM_PVC_MODE.
type CssmPvcMode = CSSM_PVC_MODE

// CssmQueryFlags is a Go-name alias for CSSM_QUERY_FLAGS.
type CssmQueryFlags = CSSM_QUERY_FLAGS

// CssmRealloc is a Go-name alias for CSSM_REALLOC.
type CssmRealloc = CSSM_REALLOC

// CssmReturn is a Go-name alias for CSSM_RETURN.
type CssmReturn = CSSM_RETURN

// CssmScFlags is a Go-name alias for CSSM_SC_FLAGS.
type CssmScFlags = CSSM_SC_FLAGS

// CssmServiceMask is a Go-name alias for CSSM_SERVICE_MASK.
type CssmServiceMask = CSSM_SERVICE_MASK

// CssmServiceType is a Go-name alias for CSSM_SERVICE_TYPE.
type CssmServiceType = CSSM_SERVICE_TYPE

// CssmSize is a Go-name alias for CSSM_SIZE.
type CssmSize = CSSM_SIZE

// CssmString is a Go-name alias for CSSM_STRING.
type CssmString = CSSM_STRING

// CssmTimestring is a Go-name alias for CSSM_TIMESTRING.
type CssmTimestring = CSSM_TIMESTRING

// CssmTpAppleCertStatus is a Go-name alias for CSSM_TP_APPLE_CERT_STATUS.
type CssmTpAppleCertStatus = CSSM_TP_APPLE_CERT_STATUS

// CssmTpAuthorityRequestType is a Go-name alias for CSSM_TP_AUTHORITY_REQUEST_TYPE.
type CssmTpAuthorityRequestType = CSSM_TP_AUTHORITY_REQUEST_TYPE

// CssmTpAuthorityRequestTypePtr is a Go-name alias for CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR.
type CssmTpAuthorityRequestTypePtr = CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR

// CssmTpCertchangeAction is a Go-name alias for CSSM_TP_CERTCHANGE_ACTION.
type CssmTpCertchangeAction = CSSM_TP_CERTCHANGE_ACTION

// CssmTpCertchangeStatus is a Go-name alias for CSSM_TP_CERTCHANGE_STATUS.
type CssmTpCertchangeStatus = CSSM_TP_CERTCHANGE_STATUS

// CssmTpCertissueStatus is a Go-name alias for CSSM_TP_CERTISSUE_STATUS.
type CssmTpCertissueStatus = CSSM_TP_CERTISSUE_STATUS

// CssmTpCertnotarizeStatus is a Go-name alias for CSSM_TP_CERTNOTARIZE_STATUS.
type CssmTpCertnotarizeStatus = CSSM_TP_CERTNOTARIZE_STATUS

// CssmTpCertreclaimStatus is a Go-name alias for CSSM_TP_CERTRECLAIM_STATUS.
type CssmTpCertreclaimStatus = CSSM_TP_CERTRECLAIM_STATUS

// CssmTpCertverifyStatus is a Go-name alias for CSSM_TP_CERTVERIFY_STATUS.
type CssmTpCertverifyStatus = CSSM_TP_CERTVERIFY_STATUS

// CssmTpConfirmStatus is a Go-name alias for CSSM_TP_CONFIRM_STATUS.
type CssmTpConfirmStatus = CSSM_TP_CONFIRM_STATUS

// CssmTpConfirmStatusPtr is a Go-name alias for CSSM_TP_CONFIRM_STATUS_PTR.
type CssmTpConfirmStatusPtr = CSSM_TP_CONFIRM_STATUS_PTR

// CssmTpCrlissueStatus is a Go-name alias for CSSM_TP_CRLISSUE_STATUS.
type CssmTpCrlissueStatus = CSSM_TP_CRLISSUE_STATUS

// CssmTpHandle is a Go-name alias for CSSM_TP_HANDLE.
type CssmTpHandle = CSSM_TP_HANDLE

// CssmTpServices is a Go-name alias for CSSM_TP_SERVICES.
type CssmTpServices = CSSM_TP_SERVICES

// CssmUseeTag is a Go-name alias for CSSM_USEE_TAG.
type CssmUseeTag = CSSM_USEE_TAG

// CssmWordidType is a Go-name alias for CSSM_WORDID_TYPE.
type CssmWordidType = CSSM_WORDID_TYPE

// CssmX509extDataFormat is a Go-name alias for CSSM_X509EXT_DATA_FORMAT.
type CssmX509extDataFormat = CSSM_X509EXT_DATA_FORMAT

// CssmX509Option is a Go-name alias for CSSM_X509_OPTION.
type CssmX509Option = CSSM_X509_OPTION

// MdsHandle is a Go-name alias for MDS_HANDLE.
type MdsHandle = MDS_HANDLE

// SecCertificate is a Go-name alias for Sec_certificate_t.
type SecCertificate = Sec_certificate_t

// SecIdentity is a Go-name alias for Sec_identity_t.
type SecIdentity = Sec_identity_t

// SecObject is a Go-name alias for Sec_object_t.
type SecObject = Sec_object_t

// SecProtocolChallengeComplete is a Go-name alias for Sec_protocol_challenge_complete_t.
type SecProtocolChallengeComplete = Sec_protocol_challenge_complete_t

// SecProtocolChallenge is a Go-name alias for Sec_protocol_challenge_t.
type SecProtocolChallenge = Sec_protocol_challenge_t

// SecProtocolKeyUpdateComplete is a Go-name alias for Sec_protocol_key_update_complete_t.
type SecProtocolKeyUpdateComplete = Sec_protocol_key_update_complete_t

// SecProtocolKeyUpdate is a Go-name alias for Sec_protocol_key_update_t.
type SecProtocolKeyUpdate = Sec_protocol_key_update_t

// SecProtocolMetadata is a Go-name alias for Sec_protocol_metadata_t.
type SecProtocolMetadata = Sec_protocol_metadata_t

// SecProtocolOptions is a Go-name alias for Sec_protocol_options_t.
type SecProtocolOptions = Sec_protocol_options_t

// SecProtocolPreSharedKeySelectionComplete is a Go-name alias for Sec_protocol_pre_shared_key_selection_complete_t.
type SecProtocolPreSharedKeySelectionComplete = Sec_protocol_pre_shared_key_selection_complete_t

// SecProtocolPreSharedKeySelection is a Go-name alias for Sec_protocol_pre_shared_key_selection_t.
type SecProtocolPreSharedKeySelection = Sec_protocol_pre_shared_key_selection_t

// SecProtocolVerifyComplete is a Go-name alias for Sec_protocol_verify_complete_t.
type SecProtocolVerifyComplete = Sec_protocol_verify_complete_t

// SecProtocolVerify is a Go-name alias for Sec_protocol_verify_t.
type SecProtocolVerify = Sec_protocol_verify_t

// SecTrust is a Go-name alias for Sec_trust_t.
type SecTrust = Sec_trust_t
