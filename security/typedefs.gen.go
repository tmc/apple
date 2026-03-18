// Code generated from Apple documentation. DO NOT EDIT.

package security

import (
	"unsafe"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

type AuthorizationAsyncCallback = func(int, *AuthorizationItemSet)

type AuthorizationEngineRef uintptr

type AuthorizationEnvironment = AuthorizationItemSet

type AuthorizationMechanismId = unsafe.Pointer

type AuthorizationMechanismRef uintptr

type AuthorizationPluginId = unsafe.Pointer

type AuthorizationPluginRef uintptr

type AuthorizationRef uintptr

type AuthorizationRights = AuthorizationItemSet

type AuthorizationSessionId = unsafe.Pointer

type AuthorizationString = *byte

type CE_CrlNumber = uint32

type CE_DeltaCrl = uint32

type CE_ExtendedKeyUsage = unsafe.Pointer

type CMSDecoderRef uintptr

type CMSEncoderRef uintptr

type CSSM_ACL_AUTHORIZATION_TAG = int32

type CSSM_ACL_EDIT_MODE = uint32

type CSSM_ACL_HANDLE = unsafe.Pointer

type CSSM_ACL_KEYCHAIN_PROMPT_SELECTOR = Cssm_acl_keychain_prompt_selector

type CSSM_ACL_PREAUTH_TRACKING_STATE = uint32

type CSSM_ACL_PROCESS_SUBJECT_SELECTOR = Cssm_acl_process_subject_selector

type CSSM_ACL_SUBJECT_TYPE = int32

type CSSM_AC_HANDLE = unsafe.Pointer

type CSSM_ALGORITHMS = uint32

type CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS = Cssm_applecspdl_db_change_password_parameters

type CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS_PTR = uintptr

type CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS = Cssm_applecspdl_db_is_locked_parameters

type CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS_PTR = uintptr

type CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS = Cssm_applecspdl_db_settings_parameters

type CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS_PTR = uintptr

type CSSM_APPLEDL_OPEN_PARAMETERS = Cssm_appledl_open_parameters

type CSSM_APPLEDL_OPEN_PARAMETERS_PTR = uintptr

type CSSM_APPLE_TP_ACTION_FLAGS = uint32

type CSSM_APPLE_TP_CRL_OPT_FLAGS = uint32

type CSSM_ATTACH_FLAGS = uint32

type CSSM_ATTRIBUTE_TYPE = uint32

type CSSM_BER_TAG = uint8

type CSSM_BITMASK = uint32

type CSSM_BOOL = int32

type CSSM_CALLOC = func(uint, uint, unsafe.Pointer) unsafe.Pointer

type CSSM_CC_HANDLE = unsafe.Pointer

type CSSM_CERTGROUP_TYPE = uint32

type CSSM_CERTGROUP_TYPE_PTR = *uint32

type CSSM_CERT_BUNDLE_ENCODING = uint32

type CSSM_CERT_BUNDLE_TYPE = uint32

type CSSM_CERT_ENCODING = uint32

type CSSM_CERT_ENCODING_PTR = *uint32

type CSSM_CERT_PARSE_FORMAT = uint32

type CSSM_CERT_PARSE_FORMAT_PTR = *uint32

type CSSM_CERT_TYPE = uint32

type CSSM_CERT_TYPE_PTR = *uint32

type CSSM_CL_HANDLE = unsafe.Pointer

type CSSM_CL_TEMPLATE_TYPE = uint32

type CSSM_CONTEXT_EVENT = uint32

type CSSM_CONTEXT_TYPE = uint32

type CSSM_CRLGROUP_TYPE = uint32

type CSSM_CRLGROUP_TYPE_PTR = *uint32

type CSSM_CRL_ENCODING = uint32

type CSSM_CRL_ENCODING_PTR = *uint32

type CSSM_CRL_PARSE_FORMAT = uint32

type CSSM_CRL_PARSE_FORMAT_PTR = *uint32

type CSSM_CRL_TYPE = uint32

type CSSM_CRL_TYPE_PTR = *uint32

type CSSM_CSPTYPE = uint32

type CSSM_CSP_FLAGS = uint32

type CSSM_CSP_HANDLE = unsafe.Pointer

type CSSM_CSP_READER_FLAGS = uint32

type CSSM_DB_ACCESS_TYPE = uint32

type CSSM_DB_ACCESS_TYPE_PTR = *uint32

type CSSM_DB_ATTRIBUTE_FORMAT = uint32

type CSSM_DB_ATTRIBUTE_FORMAT_PTR = *uint32

type CSSM_DB_ATTRIBUTE_NAME_FORMAT = uint32

type CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR = *uint32

type CSSM_DB_CONJUNCTIVE = uint32

type CSSM_DB_CONJUNCTIVE_PTR = *uint32

type CSSM_DB_HANDLE = unsafe.Pointer

type CSSM_DB_INDEXED_DATA_LOCATION = uint32

type CSSM_DB_INDEX_TYPE = uint32

type CSSM_DB_MODIFY_MODE = uint32

type CSSM_DB_OPERATOR = uint32

type CSSM_DB_OPERATOR_PTR = *uint32

type CSSM_DB_RECORDTYPE = uint32

type CSSM_DB_RETRIEVAL_MODES = uint32

type CSSM_DLTYPE = uint32

type CSSM_DLTYPE_PTR = *uint32

type CSSM_DL_CUSTOM_ATTRIBUTES = unsafe.Pointer

type CSSM_DL_FFS_ATTRIBUTES = unsafe.Pointer

type CSSM_DL_HANDLE = unsafe.Pointer

type CSSM_DL_LDAP_ATTRIBUTES = unsafe.Pointer

type CSSM_DL_ODBC_ATTRIBUTES = unsafe.Pointer

type CSSM_DL_PKCS11_ATTRIBUTE = uintptr

type CSSM_DL_PKCS11_ATTRIBUTE_PTR = uintptr

type CSSM_ENCRYPT_MODE = uint32

type CSSM_EVIDENCE_FORM = uint32

type CSSM_FREE = func(unsafe.Pointer, unsafe.Pointer)

type CSSM_HANDLE = unsafe.Pointer

type CSSM_HANDLE_PTR = unsafe.Pointer

type CSSM_HEADERVERSION = uint32

type CSSM_INTPTR = int

type CSSM_KEYATTR_FLAGS = uint32

type CSSM_KEYBLOB_FORMAT = uint32

type CSSM_KEYBLOB_TYPE = uint32

type CSSM_KEYCLASS = uint32

type CSSM_KEYUSE = uint32

type CSSM_KEY_HIERARCHY = unsafe.Pointer

type CSSM_KEY_TYPE = unsafe.Pointer

type CSSM_KRSP_HANDLE = uint32

type CSSM_KR_POLICY_FLAGS = uint32

type CSSM_KR_POLICY_TYPE = uint32

type CSSM_LIST_ELEMENT_PTR = uintptr

type CSSM_LIST_ELEMENT_TYPE = uint32

type CSSM_LIST_ELEMENT_TYPE_PTR = *uint32

type CSSM_LIST_TYPE = uint32

type CSSM_LIST_TYPE_PTR = *uint32

type CSSM_LONG_HANDLE = uint64

type CSSM_LONG_HANDLE_PTR = *uint64

type CSSM_MALLOC = func(uint, unsafe.Pointer) unsafe.Pointer

type CSSM_MANAGER_EVENT_TYPES = uint32

type CSSM_MODULE_EVENT = uint32

type CSSM_MODULE_EVENT_PTR = *uint32

type CSSM_MODULE_HANDLE = unsafe.Pointer

type CSSM_MODULE_HANDLE_PTR = unsafe.Pointer

type CSSM_NET_ADDRESS_TYPE = uint32

type CSSM_NET_PROTOCOL = uint32

type CSSM_PADDING = uint32

type CSSM_PKCS5_PBKDF2_PRF = uint32

type CSSM_PKCS_OAEP_MGF = uint32

type CSSM_PKCS_OAEP_PSOURCE = uint32

type CSSM_PRIVILEGE = uint64

type CSSM_PRIVILEGE_SCOPE = uint32

type CSSM_PROC_ADDR = func()

type CSSM_PROC_ADDR_PTR = unsafe.Pointer

type CSSM_PVC_MODE = unsafe.Pointer

type CSSM_QUERY_FLAGS = uint32

type CSSM_REALLOC = func(unsafe.Pointer, uint, unsafe.Pointer) unsafe.Pointer

type CSSM_RETURN = int32

type CSSM_SAMPLE_TYPE = unsafe.Pointer

type CSSM_SC_FLAGS = uint32

type CSSM_SERVICE_MASK = uint32

type CSSM_SERVICE_TYPE = unsafe.Pointer

type CSSM_SIZE = uintptr

type CSSM_STRING = unsafe.Pointer

type CSSM_TIMESTRING = *byte

type CSSM_TP_ACTION = uint32

type CSSM_TP_APPLE_CERT_STATUS = uint32

type CSSM_TP_AUTHORITY_REQUEST_TYPE = uint32

type CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR = *uint32

type CSSM_TP_CERTCHANGE_ACTION = uint32

type CSSM_TP_CERTCHANGE_REASON = uint32

type CSSM_TP_CERTCHANGE_STATUS = uint32

type CSSM_TP_CERTISSUE_STATUS = uint32

type CSSM_TP_CERTNOTARIZE_STATUS = uint32

type CSSM_TP_CERTRECLAIM_STATUS = uint32

type CSSM_TP_CERTVERIFY_STATUS = uint32

type CSSM_TP_CONFIRM_STATUS = uint32

type CSSM_TP_CONFIRM_STATUS_PTR = *uint32

type CSSM_TP_CRLISSUE_STATUS = uint32

type CSSM_TP_FORM_TYPE = uint32

type CSSM_TP_HANDLE = unsafe.Pointer

type CSSM_TP_SERVICES = uint32

type CSSM_TP_STOP_ON = uint32

type CSSM_USEE_TAG = unsafe.Pointer

type CSSM_WORDID_TYPE = int32

type CSSM_X509EXT_DATA_FORMAT = Extension_data_format

type CSSM_X509_OPTION = unsafe.Pointer

type MDS_HANDLE = unsafe.Pointer

type OpaqueSecAccessRef uintptr

type OpaqueSecCertificateRef uintptr

type OpaqueSecIdentityRef uintptr

type OpaqueSecKeyRef = string

type SSLCipherSuite = uint16

type SSLConnectionRef uintptr

type SSLContextRef uintptr

type SSLReadFunc = func(unsafe.Pointer, unsafe.Pointer, *uint) int

type SSLWriteFunc = func(unsafe.Pointer, unsafe.Pointer, *uint) int

type SecACLRef uintptr

type SecAccessControlRef uintptr

type SecAccessOwnerType = uint32

type SecAccessRef uintptr

type SecCertificateRef uintptr

type SecCodeRef uintptr

type SecGuestRef = objectivec.IObject

type SecIdentityRef uintptr

type SecIdentitySearchRef uintptr

type SecKeyAlgorithm = corefoundation.CFStringRef

type SecKeyGeneratePairBlock = func(string, string, objectivec.IObject)

type SecKeyKeyExchangeParameter = corefoundation.CFStringRef

type SecKeyRef uintptr

type SecKeychainAttrType = uint32

type SecKeychainAttributePtr = *SecKeychainAttribute

type SecKeychainItemRef uintptr

type SecKeychainRef uintptr

type SecKeychainSearchRef uintptr

type SecKeychainStatus = uint32

type SecMessageBlock = func(unsafe.Pointer, objectivec.IObject, uint32)

type SecPasswordRef uintptr

type SecPolicyRef uintptr

type SecPolicySearchRef uintptr

type SecPublicKeyHash = uint8

type SecRandomRef uintptr

type SecRequirementRef uintptr

type SecStaticCodeRef uintptr

type SecTaskRef uintptr

type SecTransformDataBlock = func(unsafe.Pointer) unsafe.Pointer

type SecTransformImplementationRef uintptr

type SecTransformInstanceBlock = func() objectivec.IObject

type SecTrustCallback = func(objectivec.IObject, SecTrustResultType)

type SecTrustRef uintptr

type SecTrustWithErrorCallback = func(objectivec.IObject, bool, objectivec.IObject)

type SecTrustedApplicationRef uintptr

type SecureDownloadRef uintptr

type SecuritySessionId = uint32

type Sec_certificate_t = objectivec.Object

type Sec_identity_t = objectivec.Object

type Sec_object_t = objectivec.Object

type Sec_protocol_challenge_complete_t = func(objectivec.Object)

type Sec_protocol_challenge_t = func(*objectivec.Object)

type Sec_protocol_key_update_complete_t = func()

type Sec_protocol_key_update_t = func(*objectivec.Object, func())

type Sec_protocol_metadata_t = objectivec.Object

type Sec_protocol_options_t = objectivec.Object

type Sec_protocol_pre_shared_key_selection_complete_t = func(objectivec.Object)

type Sec_protocol_pre_shared_key_selection_t = func(*objectivec.Object)

type Sec_protocol_verify_complete_t = func(bool)

type Sec_protocol_verify_t = func(*objectivec.Object, func(unsafe.Pointer))

type Sec_trust_t = objectivec.Object

type Sint16 = int16

type Sint32 = int32

type Sint64 = int64

type Sint8 = int8

