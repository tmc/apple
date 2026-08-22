// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/Security/AuthorizationContextFlags
type AuthorizationContextFlags uint32

const (
	// KAuthorizationContextFlagExtractable: It is possible for the authorization client to use the AuthorizationCopyInfo(_:_:_:) function to obtain the value.
	KAuthorizationContextFlagExtractable AuthorizationContextFlags = 1
	// KAuthorizationContextFlagSticky: This data persists through an interrupted or failed evaluation.
	KAuthorizationContextFlagSticky AuthorizationContextFlags = 4
	// KAuthorizationContextFlagVolatile: The value is not saved for the authorization client.
	KAuthorizationContextFlagVolatile AuthorizationContextFlags = 2
)

func (e AuthorizationContextFlags) String() string {
	switch e {
	case KAuthorizationContextFlagExtractable:
		return "KAuthorizationContextFlagExtractable"
	case KAuthorizationContextFlagSticky:
		return "KAuthorizationContextFlagSticky"
	case KAuthorizationContextFlagVolatile:
		return "KAuthorizationContextFlagVolatile"
	default:
		return fmt.Sprintf("AuthorizationContextFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/AuthorizationFlags
type AuthorizationFlags uint32

const (
	// KAuthorizationFlagDefaults: An empty flag set that you use as a placeholder when you don’t want any of the other flags.
	KAuthorizationFlagDefaults AuthorizationFlags = 0
	// KAuthorizationFlagDestroyRights: A flag that instructs the Security Server to revoke authorization.
	KAuthorizationFlagDestroyRights AuthorizationFlags = 8
	// KAuthorizationFlagExtendRights: A flag that permits the Security Server to attempt to grant the rights requested.
	KAuthorizationFlagExtendRights AuthorizationFlags = 2
	// KAuthorizationFlagInteractionAllowed: A flag that permits user interaction as needed.
	KAuthorizationFlagInteractionAllowed AuthorizationFlags = 1
	// KAuthorizationFlagNoData: Private flag.
	KAuthorizationFlagNoData AuthorizationFlags = 1048576
	// KAuthorizationFlagPartialRights: A flag that permits the Security Server to grant rights on an individual basis.
	KAuthorizationFlagPartialRights AuthorizationFlags = 4
	// KAuthorizationFlagPreAuthorize: A flag that instructs the Security Server to preauthorize the rights requested.
	KAuthorizationFlagPreAuthorize     AuthorizationFlags = 16
	KAuthorizationFlagSkipInternalAuth AuthorizationFlags = 512
)

func (e AuthorizationFlags) String() string {
	switch e {
	case KAuthorizationFlagDefaults:
		return "KAuthorizationFlagDefaults"
	case KAuthorizationFlagDestroyRights:
		return "KAuthorizationFlagDestroyRights"
	case KAuthorizationFlagExtendRights:
		return "KAuthorizationFlagExtendRights"
	case KAuthorizationFlagInteractionAllowed:
		return "KAuthorizationFlagInteractionAllowed"
	case KAuthorizationFlagNoData:
		return "KAuthorizationFlagNoData"
	case KAuthorizationFlagPartialRights:
		return "KAuthorizationFlagPartialRights"
	case KAuthorizationFlagPreAuthorize:
		return "KAuthorizationFlagPreAuthorize"
	case KAuthorizationFlagSkipInternalAuth:
		return "KAuthorizationFlagSkipInternalAuth"
	default:
		return fmt.Sprintf("AuthorizationFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/AuthorizationResult
type AuthorizationResult uint32

const (
	// KAuthorizationResultAllow: The authorization operation succeeded and authorization should be granted.
	KAuthorizationResultAllow AuthorizationResult = 0
	// KAuthorizationResultDeny: The authorization operation succeeded and authorization should be denied.
	KAuthorizationResultDeny AuthorizationResult = 1
	// KAuthorizationResultUndefined: The authorization operation failed and should not be retried for this session.
	KAuthorizationResultUndefined AuthorizationResult = 2
	// KAuthorizationResultUserCanceled: The user has requested that the authorization evaluation be terminated.
	KAuthorizationResultUserCanceled AuthorizationResult = 3
)

func (e AuthorizationResult) String() string {
	switch e {
	case KAuthorizationResultAllow:
		return "KAuthorizationResultAllow"
	case KAuthorizationResultDeny:
		return "KAuthorizationResultDeny"
	case KAuthorizationResultUndefined:
		return "KAuthorizationResultUndefined"
	case KAuthorizationResultUserCanceled:
		return "KAuthorizationResultUserCanceled"
	default:
		return fmt.Sprintf("AuthorizationResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/CMSCertificateChainMode
type CMSCertificateChainMode uint32

const (
	// KCMSCertificateChain: Include the signer certificate chain up to but not including the root certificate.
	KCMSCertificateChain CMSCertificateChainMode = 2
	// KCMSCertificateChainWithRoot: Include the entire signer certificate chain, including the root certificate.
	KCMSCertificateChainWithRoot       CMSCertificateChainMode = 3
	KCMSCertificateChainWithRootOrFail CMSCertificateChainMode = 4
	// KCMSCertificateNone: Don’t include any certificates.
	KCMSCertificateNone CMSCertificateChainMode = 0
	// KCMSCertificateSignerOnly: Only include signer certificates.
	KCMSCertificateSignerOnly CMSCertificateChainMode = 1
)

func (e CMSCertificateChainMode) String() string {
	switch e {
	case KCMSCertificateChain:
		return "KCMSCertificateChain"
	case KCMSCertificateChainWithRoot:
		return "KCMSCertificateChainWithRoot"
	case KCMSCertificateChainWithRootOrFail:
		return "KCMSCertificateChainWithRootOrFail"
	case KCMSCertificateNone:
		return "KCMSCertificateNone"
	case KCMSCertificateSignerOnly:
		return "KCMSCertificateSignerOnly"
	default:
		return fmt.Sprintf("CMSCertificateChainMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/CMSSignedAttributes
type CMSSignedAttributes uint32

const (
	// KCMSAttrAppleCodesigningHashAgility: Include Apple codesigning hash agility.
	KCMSAttrAppleCodesigningHashAgility CMSSignedAttributes = 0x10
	// KCMSAttrAppleCodesigningHashAgilityV2: Include Apple codesigning hash agility, version 2.
	KCMSAttrAppleCodesigningHashAgilityV2 CMSSignedAttributes = 0x20
	// KCMSAttrAppleExpirationTime: Include the expiration time.
	KCMSAttrAppleExpirationTime CMSSignedAttributes = 0x40
	// KCMSAttrNone: No attributes.
	KCMSAttrNone CMSSignedAttributes = 0
	// KCMSAttrSigningTime: Include the signing time.
	KCMSAttrSigningTime CMSSignedAttributes = 0x8
	// KCMSAttrSmimeCapabilities: Identify signature, encryption, and digest algorithms supported by the encoder.
	KCMSAttrSmimeCapabilities CMSSignedAttributes = 0x1
	// KCMSAttrSmimeEncryptionKeyPrefs: Indicate that the signing certificate included with the message is the preferred one for S/MIME encryption.
	KCMSAttrSmimeEncryptionKeyPrefs CMSSignedAttributes = 0x2
	// KCMSAttrSmimeMSEncryptionKeyPrefs: Indicate that the signing certificate included with the message is the preferred one for S/MIME encryption, but using an attribute object identifier (OID) preferred by Microsoft.
	KCMSAttrSmimeMSEncryptionKeyPrefs CMSSignedAttributes = 0x4
)

func (e CMSSignedAttributes) String() string {
	switch e {
	case KCMSAttrAppleCodesigningHashAgility:
		return "KCMSAttrAppleCodesigningHashAgility"
	case KCMSAttrAppleCodesigningHashAgilityV2:
		return "KCMSAttrAppleCodesigningHashAgilityV2"
	case KCMSAttrAppleExpirationTime:
		return "KCMSAttrAppleExpirationTime"
	case KCMSAttrNone:
		return "KCMSAttrNone"
	case KCMSAttrSigningTime:
		return "KCMSAttrSigningTime"
	case KCMSAttrSmimeCapabilities:
		return "KCMSAttrSmimeCapabilities"
	case KCMSAttrSmimeEncryptionKeyPrefs:
		return "KCMSAttrSmimeEncryptionKeyPrefs"
	case KCMSAttrSmimeMSEncryptionKeyPrefs:
		return "KCMSAttrSmimeMSEncryptionKeyPrefs"
	default:
		return fmt.Sprintf("CMSSignedAttributes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/CMSSignerStatus
type CMSSignerStatus uint32

const (
	// KCMSSignerInvalidCert: The message was signed but the signer’s certificate could not be verified.
	KCMSSignerInvalidCert CMSSignerStatus = 4
	// KCMSSignerInvalidIndex: The specified value for the signer index (`signerIndex` parameter) is greater than the number of signers of the message minus one (`signerIndex > (numSigners – 1)`).
	KCMSSignerInvalidIndex CMSSignerStatus = 5
	// KCMSSignerInvalidSignature: The message was signed but the signature is invalid.
	KCMSSignerInvalidSignature CMSSignerStatus = 3
	// KCMSSignerNeedsDetachedContent: The message was signed but has detached content.
	KCMSSignerNeedsDetachedContent CMSSignerStatus = 2
	// KCMSSignerUnsigned: The message was not signed.
	KCMSSignerUnsigned CMSSignerStatus = 0
	// KCMSSignerValid: The message was signed and both the signature and the signer certificate have been verified.
	KCMSSignerValid CMSSignerStatus = 1
)

func (e CMSSignerStatus) String() string {
	switch e {
	case KCMSSignerInvalidCert:
		return "KCMSSignerInvalidCert"
	case KCMSSignerInvalidIndex:
		return "KCMSSignerInvalidIndex"
	case KCMSSignerInvalidSignature:
		return "KCMSSignerInvalidSignature"
	case KCMSSignerNeedsDetachedContent:
		return "KCMSSignerNeedsDetachedContent"
	case KCMSSignerUnsigned:
		return "KCMSSignerUnsigned"
	case KCMSSignerValid:
		return "KCMSSignerValid"
	default:
		return fmt.Sprintf("CMSSignerStatus(%d)", e)
	}
}

type CssmAcBaseAcError int32

const (
	CSSMERR_AC_INVALID_BASE_ACLS          CssmAcBaseAcError = -2147405567
	CSSMERR_AC_INVALID_ENCODING           CssmAcBaseAcError = -2147405565
	CSSMERR_AC_INVALID_REQUESTOR          CssmAcBaseAcError = -2147405563
	CSSMERR_AC_INVALID_REQUEST_DESCRIPTOR CssmAcBaseAcError = -2147405562
	CSSMERR_AC_INVALID_TUPLE_CREDENTIALS  CssmAcBaseAcError = -2147405566
	CSSMERR_AC_INVALID_VALIDITY_PERIOD    CssmAcBaseAcError = -2147405564
	CSSM_AC_BASE_AC_ERROR                 CssmAcBaseAcError = -2147405568
)

func (e CssmAcBaseAcError) String() string {
	switch e {
	case CSSMERR_AC_INVALID_BASE_ACLS:
		return "CSSMERR_AC_INVALID_BASE_ACLS"
	case CSSMERR_AC_INVALID_ENCODING:
		return "CSSMERR_AC_INVALID_ENCODING"
	case CSSMERR_AC_INVALID_REQUESTOR:
		return "CSSMERR_AC_INVALID_REQUESTOR"
	case CSSMERR_AC_INVALID_REQUEST_DESCRIPTOR:
		return "CSSMERR_AC_INVALID_REQUEST_DESCRIPTOR"
	case CSSMERR_AC_INVALID_TUPLE_CREDENTIALS:
		return "CSSMERR_AC_INVALID_TUPLE_CREDENTIALS"
	case CSSMERR_AC_INVALID_VALIDITY_PERIOD:
		return "CSSMERR_AC_INVALID_VALIDITY_PERIOD"
	case CSSM_AC_BASE_AC_ERROR:
		return "CSSM_AC_BASE_AC_ERROR"
	default:
		return fmt.Sprintf("CssmAcBaseAcError(%d)", e)
	}
}

type CssmAclAuthorizationChangeAcl uint32

const (
	CSSM_ACL_AUTHORIZATION_CHANGE_ACL   CssmAclAuthorizationChangeAcl = 65536
	CSSM_ACL_AUTHORIZATION_CHANGE_OWNER CssmAclAuthorizationChangeAcl = 65537
	CSSM_ACL_AUTHORIZATION_INTEGRITY    CssmAclAuthorizationChangeAcl = 65539
	CSSM_ACL_AUTHORIZATION_PARTITION_ID CssmAclAuthorizationChangeAcl = 65538
	CSSM_ACL_AUTHORIZATION_PREAUTH_BASE CssmAclAuthorizationChangeAcl = 16842752
	CSSM_ACL_AUTHORIZATION_PREAUTH_END  CssmAclAuthorizationChangeAcl = 16908288
)

func (e CssmAclAuthorizationChangeAcl) String() string {
	switch e {
	case CSSM_ACL_AUTHORIZATION_CHANGE_ACL:
		return "CSSM_ACL_AUTHORIZATION_CHANGE_ACL"
	case CSSM_ACL_AUTHORIZATION_CHANGE_OWNER:
		return "CSSM_ACL_AUTHORIZATION_CHANGE_OWNER"
	case CSSM_ACL_AUTHORIZATION_INTEGRITY:
		return "CSSM_ACL_AUTHORIZATION_INTEGRITY"
	case CSSM_ACL_AUTHORIZATION_PARTITION_ID:
		return "CSSM_ACL_AUTHORIZATION_PARTITION_ID"
	case CSSM_ACL_AUTHORIZATION_PREAUTH_BASE:
		return "CSSM_ACL_AUTHORIZATION_PREAUTH_BASE"
	case CSSM_ACL_AUTHORIZATION_PREAUTH_END:
		return "CSSM_ACL_AUTHORIZATION_PREAUTH_END"
	default:
		return fmt.Sprintf("CssmAclAuthorizationChangeAcl(%d)", e)
	}
}

type CssmAclAuthorizationTagVendorDefinedStart uint32

const (
	CSSM_ACL_AUTHORIZATION_ANY                      CssmAclAuthorizationTagVendorDefinedStart = 1
	CSSM_ACL_AUTHORIZATION_DBS_CREATE               CssmAclAuthorizationTagVendorDefinedStart = 22
	CSSM_ACL_AUTHORIZATION_DBS_DELETE               CssmAclAuthorizationTagVendorDefinedStart = 23
	CSSM_ACL_AUTHORIZATION_DB_DELETE                CssmAclAuthorizationTagVendorDefinedStart = 17
	CSSM_ACL_AUTHORIZATION_DB_INSERT                CssmAclAuthorizationTagVendorDefinedStart = 19
	CSSM_ACL_AUTHORIZATION_DB_MODIFY                CssmAclAuthorizationTagVendorDefinedStart = 20
	CSSM_ACL_AUTHORIZATION_DB_READ                  CssmAclAuthorizationTagVendorDefinedStart = 21
	CSSM_ACL_AUTHORIZATION_DECRYPT                  CssmAclAuthorizationTagVendorDefinedStart = 24
	CSSM_ACL_AUTHORIZATION_DELETE                   CssmAclAuthorizationTagVendorDefinedStart = 25
	CSSM_ACL_AUTHORIZATION_DERIVE                   CssmAclAuthorizationTagVendorDefinedStart = 28
	CSSM_ACL_AUTHORIZATION_ENCRYPT                  CssmAclAuthorizationTagVendorDefinedStart = 35
	CSSM_ACL_AUTHORIZATION_EXPORT_CLEAR             CssmAclAuthorizationTagVendorDefinedStart = 37
	CSSM_ACL_AUTHORIZATION_EXPORT_WRAPPED           CssmAclAuthorizationTagVendorDefinedStart = 38
	CSSM_ACL_AUTHORIZATION_GENKEY                   CssmAclAuthorizationTagVendorDefinedStart = 41
	CSSM_ACL_AUTHORIZATION_IMPORT_CLEAR             CssmAclAuthorizationTagVendorDefinedStart = 47
	CSSM_ACL_AUTHORIZATION_IMPORT_WRAPPED           CssmAclAuthorizationTagVendorDefinedStart = 48
	CSSM_ACL_AUTHORIZATION_LOGIN                    CssmAclAuthorizationTagVendorDefinedStart = 57
	CSSM_ACL_AUTHORIZATION_MAC                      CssmAclAuthorizationTagVendorDefinedStart = 59
	CSSM_ACL_AUTHORIZATION_SIGN                     CssmAclAuthorizationTagVendorDefinedStart = 115
	CSSM_ACL_AUTHORIZATION_TAG_VENDOR_DEFINED_START CssmAclAuthorizationTagVendorDefinedStart = 0x10000
)

func (e CssmAclAuthorizationTagVendorDefinedStart) String() string {
	switch e {
	case CSSM_ACL_AUTHORIZATION_ANY:
		return "CSSM_ACL_AUTHORIZATION_ANY"
	case CSSM_ACL_AUTHORIZATION_DBS_CREATE:
		return "CSSM_ACL_AUTHORIZATION_DBS_CREATE"
	case CSSM_ACL_AUTHORIZATION_DBS_DELETE:
		return "CSSM_ACL_AUTHORIZATION_DBS_DELETE"
	case CSSM_ACL_AUTHORIZATION_DB_DELETE:
		return "CSSM_ACL_AUTHORIZATION_DB_DELETE"
	case CSSM_ACL_AUTHORIZATION_DB_INSERT:
		return "CSSM_ACL_AUTHORIZATION_DB_INSERT"
	case CSSM_ACL_AUTHORIZATION_DB_MODIFY:
		return "CSSM_ACL_AUTHORIZATION_DB_MODIFY"
	case CSSM_ACL_AUTHORIZATION_DB_READ:
		return "CSSM_ACL_AUTHORIZATION_DB_READ"
	case CSSM_ACL_AUTHORIZATION_DECRYPT:
		return "CSSM_ACL_AUTHORIZATION_DECRYPT"
	case CSSM_ACL_AUTHORIZATION_DELETE:
		return "CSSM_ACL_AUTHORIZATION_DELETE"
	case CSSM_ACL_AUTHORIZATION_DERIVE:
		return "CSSM_ACL_AUTHORIZATION_DERIVE"
	case CSSM_ACL_AUTHORIZATION_ENCRYPT:
		return "CSSM_ACL_AUTHORIZATION_ENCRYPT"
	case CSSM_ACL_AUTHORIZATION_EXPORT_CLEAR:
		return "CSSM_ACL_AUTHORIZATION_EXPORT_CLEAR"
	case CSSM_ACL_AUTHORIZATION_EXPORT_WRAPPED:
		return "CSSM_ACL_AUTHORIZATION_EXPORT_WRAPPED"
	case CSSM_ACL_AUTHORIZATION_GENKEY:
		return "CSSM_ACL_AUTHORIZATION_GENKEY"
	case CSSM_ACL_AUTHORIZATION_IMPORT_CLEAR:
		return "CSSM_ACL_AUTHORIZATION_IMPORT_CLEAR"
	case CSSM_ACL_AUTHORIZATION_IMPORT_WRAPPED:
		return "CSSM_ACL_AUTHORIZATION_IMPORT_WRAPPED"
	case CSSM_ACL_AUTHORIZATION_LOGIN:
		return "CSSM_ACL_AUTHORIZATION_LOGIN"
	case CSSM_ACL_AUTHORIZATION_MAC:
		return "CSSM_ACL_AUTHORIZATION_MAC"
	case CSSM_ACL_AUTHORIZATION_SIGN:
		return "CSSM_ACL_AUTHORIZATION_SIGN"
	case CSSM_ACL_AUTHORIZATION_TAG_VENDOR_DEFINED_START:
		return "CSSM_ACL_AUTHORIZATION_TAG_VENDOR_DEFINED_START"
	default:
		return fmt.Sprintf("CssmAclAuthorizationTagVendorDefinedStart(%d)", e)
	}
}

type CssmAclCodeSignature uint32

const (
	CSSM_ACL_CODE_SIGNATURE_INVALID CssmAclCodeSignature = 0
	CSSM_ACL_CODE_SIGNATURE_OSX     CssmAclCodeSignature = 1
)

func (e CssmAclCodeSignature) String() string {
	switch e {
	case CSSM_ACL_CODE_SIGNATURE_INVALID:
		return "CSSM_ACL_CODE_SIGNATURE_INVALID"
	case CSSM_ACL_CODE_SIGNATURE_OSX:
		return "CSSM_ACL_CODE_SIGNATURE_OSX"
	default:
		return fmt.Sprintf("CssmAclCodeSignature(%d)", e)
	}
}

type CssmAclEditMode uint32

const (
	CSSM_ACL_EDIT_MODE_ADD     CssmAclEditMode = 1
	CSSM_ACL_EDIT_MODE_DELETE  CssmAclEditMode = 2
	CSSM_ACL_EDIT_MODE_REPLACE CssmAclEditMode = 3
)

func (e CssmAclEditMode) String() string {
	switch e {
	case CSSM_ACL_EDIT_MODE_ADD:
		return "CSSM_ACL_EDIT_MODE_ADD"
	case CSSM_ACL_EDIT_MODE_DELETE:
		return "CSSM_ACL_EDIT_MODE_DELETE"
	case CSSM_ACL_EDIT_MODE_REPLACE:
		return "CSSM_ACL_EDIT_MODE_REPLACE"
	default:
		return fmt.Sprintf("CssmAclEditMode(%d)", e)
	}
}

type CssmAclKeychainPrompt uint32

const (
	CSSM_ACL_KEYCHAIN_PROMPT_INVALID            CssmAclKeychainPrompt = 0x40
	CSSM_ACL_KEYCHAIN_PROMPT_INVALID_ACT        CssmAclKeychainPrompt = 0x80
	CSSM_ACL_KEYCHAIN_PROMPT_REQUIRE_PASSPHRASE CssmAclKeychainPrompt = 0x1
	CSSM_ACL_KEYCHAIN_PROMPT_UNSIGNED           CssmAclKeychainPrompt = 0x10
	CSSM_ACL_KEYCHAIN_PROMPT_UNSIGNED_ACT       CssmAclKeychainPrompt = 0x20
)

func (e CssmAclKeychainPrompt) String() string {
	switch e {
	case CSSM_ACL_KEYCHAIN_PROMPT_INVALID:
		return "CSSM_ACL_KEYCHAIN_PROMPT_INVALID"
	case CSSM_ACL_KEYCHAIN_PROMPT_INVALID_ACT:
		return "CSSM_ACL_KEYCHAIN_PROMPT_INVALID_ACT"
	case CSSM_ACL_KEYCHAIN_PROMPT_REQUIRE_PASSPHRASE:
		return "CSSM_ACL_KEYCHAIN_PROMPT_REQUIRE_PASSPHRASE"
	case CSSM_ACL_KEYCHAIN_PROMPT_UNSIGNED:
		return "CSSM_ACL_KEYCHAIN_PROMPT_UNSIGNED"
	case CSSM_ACL_KEYCHAIN_PROMPT_UNSIGNED_ACT:
		return "CSSM_ACL_KEYCHAIN_PROMPT_UNSIGNED_ACT"
	default:
		return fmt.Sprintf("CssmAclKeychainPrompt(%d)", e)
	}
}

type CssmAclKeychainPromptCurrent uint32

const (
	CSSM_ACL_KEYCHAIN_PROMPT_CURRENT_VERSION CssmAclKeychainPromptCurrent = 0x101
)

func (e CssmAclKeychainPromptCurrent) String() string {
	switch e {
	case CSSM_ACL_KEYCHAIN_PROMPT_CURRENT_VERSION:
		return "CSSM_ACL_KEYCHAIN_PROMPT_CURRENT_VERSION"
	default:
		return fmt.Sprintf("CssmAclKeychainPromptCurrent(%d)", e)
	}
}

type CssmAclMatch uint32

const (
	CSSM_ACL_MATCH_BITS       CssmAclMatch = 3
	CSSM_ACL_MATCH_GID        CssmAclMatch = 0x2
	CSSM_ACL_MATCH_HONOR_ROOT CssmAclMatch = 0x100
	CSSM_ACL_MATCH_UID        CssmAclMatch = 0x1
)

func (e CssmAclMatch) String() string {
	switch e {
	case CSSM_ACL_MATCH_BITS:
		return "CSSM_ACL_MATCH_BITS"
	case CSSM_ACL_MATCH_GID:
		return "CSSM_ACL_MATCH_GID"
	case CSSM_ACL_MATCH_HONOR_ROOT:
		return "CSSM_ACL_MATCH_HONOR_ROOT"
	case CSSM_ACL_MATCH_UID:
		return "CSSM_ACL_MATCH_UID"
	default:
		return fmt.Sprintf("CssmAclMatch(%d)", e)
	}
}

type CssmAclPreauthTracking uint32

const (
	CSSM_ACL_PREAUTH_TRACKING_AUTHORIZED CssmAclPreauthTracking = 0x80000000
	CSSM_ACL_PREAUTH_TRACKING_BLOCKED    CssmAclPreauthTracking = 0
	CSSM_ACL_PREAUTH_TRACKING_COUNT_MASK CssmAclPreauthTracking = 0xff
	CSSM_ACL_PREAUTH_TRACKING_UNKNOWN    CssmAclPreauthTracking = 0x40000000
)

func (e CssmAclPreauthTracking) String() string {
	switch e {
	case CSSM_ACL_PREAUTH_TRACKING_AUTHORIZED:
		return "CSSM_ACL_PREAUTH_TRACKING_AUTHORIZED"
	case CSSM_ACL_PREAUTH_TRACKING_BLOCKED:
		return "CSSM_ACL_PREAUTH_TRACKING_BLOCKED"
	case CSSM_ACL_PREAUTH_TRACKING_COUNT_MASK:
		return "CSSM_ACL_PREAUTH_TRACKING_COUNT_MASK"
	case CSSM_ACL_PREAUTH_TRACKING_UNKNOWN:
		return "CSSM_ACL_PREAUTH_TRACKING_UNKNOWN"
	default:
		return fmt.Sprintf("CssmAclPreauthTracking(%d)", e)
	}
}

type CssmAclProcessSelectorCurrent uint32

const (
	CSSM_ACL_PROCESS_SELECTOR_CURRENT_VERSION CssmAclProcessSelectorCurrent = 0x101
)

func (e CssmAclProcessSelectorCurrent) String() string {
	switch e {
	case CSSM_ACL_PROCESS_SELECTOR_CURRENT_VERSION:
		return "CSSM_ACL_PROCESS_SELECTOR_CURRENT_VERSION"
	default:
		return fmt.Sprintf("CssmAclProcessSelectorCurrent(%d)", e)
	}
}

type CssmAclSubjectTypeAny uint32

const (
	CSSM_ACL_SUBJECT_TYPE_ANY                 CssmAclSubjectTypeAny = 1
	CSSM_ACL_SUBJECT_TYPE_BIOMETRIC           CssmAclSubjectTypeAny = 8
	CSSM_ACL_SUBJECT_TYPE_EXT_PAM_NAME        CssmAclSubjectTypeAny = 78
	CSSM_ACL_SUBJECT_TYPE_HASHED_SUBJECT      CssmAclSubjectTypeAny = 44
	CSSM_ACL_SUBJECT_TYPE_LOGIN_NAME          CssmAclSubjectTypeAny = 58
	CSSM_ACL_SUBJECT_TYPE_PASSWORD            CssmAclSubjectTypeAny = 79
	CSSM_ACL_SUBJECT_TYPE_PROMPTED_BIOMETRIC  CssmAclSubjectTypeAny = 83
	CSSM_ACL_SUBJECT_TYPE_PROMPTED_PASSWORD   CssmAclSubjectTypeAny = 84
	CSSM_ACL_SUBJECT_TYPE_PROTECTED_BIOMETRIC CssmAclSubjectTypeAny = 86
	CSSM_ACL_SUBJECT_TYPE_PROTECTED_PASSWORD  CssmAclSubjectTypeAny = 87
	CSSM_ACL_SUBJECT_TYPE_PUBLIC_KEY          CssmAclSubjectTypeAny = 89
	CSSM_ACL_SUBJECT_TYPE_THRESHOLD           CssmAclSubjectTypeAny = 123
)

func (e CssmAclSubjectTypeAny) String() string {
	switch e {
	case CSSM_ACL_SUBJECT_TYPE_ANY:
		return "CSSM_ACL_SUBJECT_TYPE_ANY"
	case CSSM_ACL_SUBJECT_TYPE_BIOMETRIC:
		return "CSSM_ACL_SUBJECT_TYPE_BIOMETRIC"
	case CSSM_ACL_SUBJECT_TYPE_EXT_PAM_NAME:
		return "CSSM_ACL_SUBJECT_TYPE_EXT_PAM_NAME"
	case CSSM_ACL_SUBJECT_TYPE_HASHED_SUBJECT:
		return "CSSM_ACL_SUBJECT_TYPE_HASHED_SUBJECT"
	case CSSM_ACL_SUBJECT_TYPE_LOGIN_NAME:
		return "CSSM_ACL_SUBJECT_TYPE_LOGIN_NAME"
	case CSSM_ACL_SUBJECT_TYPE_PASSWORD:
		return "CSSM_ACL_SUBJECT_TYPE_PASSWORD"
	case CSSM_ACL_SUBJECT_TYPE_PROMPTED_BIOMETRIC:
		return "CSSM_ACL_SUBJECT_TYPE_PROMPTED_BIOMETRIC"
	case CSSM_ACL_SUBJECT_TYPE_PROMPTED_PASSWORD:
		return "CSSM_ACL_SUBJECT_TYPE_PROMPTED_PASSWORD"
	case CSSM_ACL_SUBJECT_TYPE_PROTECTED_BIOMETRIC:
		return "CSSM_ACL_SUBJECT_TYPE_PROTECTED_BIOMETRIC"
	case CSSM_ACL_SUBJECT_TYPE_PROTECTED_PASSWORD:
		return "CSSM_ACL_SUBJECT_TYPE_PROTECTED_PASSWORD"
	case CSSM_ACL_SUBJECT_TYPE_PUBLIC_KEY:
		return "CSSM_ACL_SUBJECT_TYPE_PUBLIC_KEY"
	case CSSM_ACL_SUBJECT_TYPE_THRESHOLD:
		return "CSSM_ACL_SUBJECT_TYPE_THRESHOLD"
	default:
		return fmt.Sprintf("CssmAclSubjectTypeAny(%d)", e)
	}
}

type CssmAclSubjectTypeKeychainPrompt uint32

const (
	CSSM_ACL_SUBJECT_TYPE_ASYMMETRIC_KEY  CssmAclSubjectTypeKeychainPrompt = 65547
	CSSM_ACL_SUBJECT_TYPE_CODE_SIGNATURE  CssmAclSubjectTypeKeychainPrompt = 116
	CSSM_ACL_SUBJECT_TYPE_COMMENT         CssmAclSubjectTypeKeychainPrompt = 12
	CSSM_ACL_SUBJECT_TYPE_KEYCHAIN_PROMPT CssmAclSubjectTypeKeychainPrompt = 65536
	CSSM_ACL_SUBJECT_TYPE_PARTITION       CssmAclSubjectTypeKeychainPrompt = 65548
	CSSM_ACL_SUBJECT_TYPE_PREAUTH         CssmAclSubjectTypeKeychainPrompt = 65545
	CSSM_ACL_SUBJECT_TYPE_PREAUTH_SOURCE  CssmAclSubjectTypeKeychainPrompt = 65546
	CSSM_ACL_SUBJECT_TYPE_PROCESS         CssmAclSubjectTypeKeychainPrompt = 65539
	CSSM_ACL_SUBJECT_TYPE_SYMMETRIC_KEY   CssmAclSubjectTypeKeychainPrompt = 65541
)

func (e CssmAclSubjectTypeKeychainPrompt) String() string {
	switch e {
	case CSSM_ACL_SUBJECT_TYPE_ASYMMETRIC_KEY:
		return "CSSM_ACL_SUBJECT_TYPE_ASYMMETRIC_KEY"
	case CSSM_ACL_SUBJECT_TYPE_CODE_SIGNATURE:
		return "CSSM_ACL_SUBJECT_TYPE_CODE_SIGNATURE"
	case CSSM_ACL_SUBJECT_TYPE_COMMENT:
		return "CSSM_ACL_SUBJECT_TYPE_COMMENT"
	case CSSM_ACL_SUBJECT_TYPE_KEYCHAIN_PROMPT:
		return "CSSM_ACL_SUBJECT_TYPE_KEYCHAIN_PROMPT"
	case CSSM_ACL_SUBJECT_TYPE_PARTITION:
		return "CSSM_ACL_SUBJECT_TYPE_PARTITION"
	case CSSM_ACL_SUBJECT_TYPE_PREAUTH:
		return "CSSM_ACL_SUBJECT_TYPE_PREAUTH"
	case CSSM_ACL_SUBJECT_TYPE_PREAUTH_SOURCE:
		return "CSSM_ACL_SUBJECT_TYPE_PREAUTH_SOURCE"
	case CSSM_ACL_SUBJECT_TYPE_PROCESS:
		return "CSSM_ACL_SUBJECT_TYPE_PROCESS"
	case CSSM_ACL_SUBJECT_TYPE_SYMMETRIC_KEY:
		return "CSSM_ACL_SUBJECT_TYPE_SYMMETRIC_KEY"
	default:
		return fmt.Sprintf("CssmAclSubjectTypeKeychainPrompt(%d)", e)
	}
}

type CssmAddr uint32

const (
	CSSM_ADDR_CUSTOM   CssmAddr = 1
	CSSM_ADDR_NAME     CssmAddr = 4
	CSSM_ADDR_NONE     CssmAddr = 0
	CSSM_ADDR_SOCKADDR CssmAddr = 3
	CSSM_ADDR_URL      CssmAddr = 2
)

func (e CssmAddr) String() string {
	switch e {
	case CSSM_ADDR_CUSTOM:
		return "CSSM_ADDR_CUSTOM"
	case CSSM_ADDR_NAME:
		return "CSSM_ADDR_NAME"
	case CSSM_ADDR_NONE:
		return "CSSM_ADDR_NONE"
	case CSSM_ADDR_SOCKADDR:
		return "CSSM_ADDR_SOCKADDR"
	case CSSM_ADDR_URL:
		return "CSSM_ADDR_URL"
	default:
		return fmt.Sprintf("CssmAddr(%d)", e)
	}
}

type CssmAlgclass uint32

const (
	CSSM_ALGCLASS_ASYMMETRIC CssmAlgclass = 8
	CSSM_ALGCLASS_CUSTOM     CssmAlgclass = 1
	CSSM_ALGCLASS_DERIVEKEY  CssmAlgclass = 10
	CSSM_ALGCLASS_DIGEST     CssmAlgclass = 4
	CSSM_ALGCLASS_KEYGEN     CssmAlgclass = 9
	CSSM_ALGCLASS_MAC        CssmAlgclass = 7
	CSSM_ALGCLASS_NONE       CssmAlgclass = 0
	CSSM_ALGCLASS_RANDOMGEN  CssmAlgclass = 5
	CSSM_ALGCLASS_SIGNATURE  CssmAlgclass = 2
	CSSM_ALGCLASS_SYMMETRIC  CssmAlgclass = 3
	CSSM_ALGCLASS_UNIQUEGEN  CssmAlgclass = 6
)

func (e CssmAlgclass) String() string {
	switch e {
	case CSSM_ALGCLASS_ASYMMETRIC:
		return "CSSM_ALGCLASS_ASYMMETRIC"
	case CSSM_ALGCLASS_CUSTOM:
		return "CSSM_ALGCLASS_CUSTOM"
	case CSSM_ALGCLASS_DERIVEKEY:
		return "CSSM_ALGCLASS_DERIVEKEY"
	case CSSM_ALGCLASS_DIGEST:
		return "CSSM_ALGCLASS_DIGEST"
	case CSSM_ALGCLASS_KEYGEN:
		return "CSSM_ALGCLASS_KEYGEN"
	case CSSM_ALGCLASS_MAC:
		return "CSSM_ALGCLASS_MAC"
	case CSSM_ALGCLASS_NONE:
		return "CSSM_ALGCLASS_NONE"
	case CSSM_ALGCLASS_RANDOMGEN:
		return "CSSM_ALGCLASS_RANDOMGEN"
	case CSSM_ALGCLASS_SIGNATURE:
		return "CSSM_ALGCLASS_SIGNATURE"
	case CSSM_ALGCLASS_SYMMETRIC:
		return "CSSM_ALGCLASS_SYMMETRIC"
	case CSSM_ALGCLASS_UNIQUEGEN:
		return "CSSM_ALGCLASS_UNIQUEGEN"
	default:
		return fmt.Sprintf("CssmAlgclass(%d)", e)
	}
}

type CssmAlgidAppleYarrow uint32

const (
	CSSM_ALGID_AES               CssmAlgidAppleYarrow = 2147483649
	CSSM_ALGID_APPLE_YARROW      CssmAlgidAppleYarrow = 2147483648
	CSSM_ALGID_ASC               CssmAlgidAppleYarrow = 2147483655
	CSSM_ALGID_ECDH_X963_KDF     CssmAlgidAppleYarrow = 2147483677
	CSSM_ALGID_ECDSA_SPECIFIED   CssmAlgidAppleYarrow = 2147483676
	CSSM_ALGID_ENTROPY_DEFAULT   CssmAlgidAppleYarrow = 2147483665
	CSSM_ALGID_FEE               CssmAlgidAppleYarrow = 2147483650
	CSSM_ALGID_FEED              CssmAlgidAppleYarrow = 2147483653
	CSSM_ALGID_FEEDEXP           CssmAlgidAppleYarrow = 2147483654
	CSSM_ALGID_FEE_MD5           CssmAlgidAppleYarrow = 2147483651
	CSSM_ALGID_FEE_SHA1          CssmAlgidAppleYarrow = 2147483652
	CSSM_ALGID_KEYCHAIN_KEY      CssmAlgidAppleYarrow = 2147483657
	CSSM_ALGID_OPENSSH1          CssmAlgidAppleYarrow = 2147483671
	CSSM_ALGID_PBE_OPENSSL_MD5   CssmAlgidAppleYarrow = 2147483661
	CSSM_ALGID_PKCS12_PBE_ENCR   CssmAlgidAppleYarrow = 2147483658
	CSSM_ALGID_PKCS12_PBE_MAC    CssmAlgidAppleYarrow = 2147483659
	CSSM_ALGID_SECURE_PASSPHRASE CssmAlgidAppleYarrow = 2147483660
	CSSM_ALGID_SHA1HMAC_LEGACY   CssmAlgidAppleYarrow = 2147483656
	CSSM_ALGID_SHA224            CssmAlgidAppleYarrow = 2147483666
	CSSM_ALGID_SHA224WithECDSA   CssmAlgidAppleYarrow = 2147483672
	CSSM_ALGID_SHA224WithRSA     CssmAlgidAppleYarrow = 2147483667
	CSSM_ALGID_SHA256            CssmAlgidAppleYarrow = 2147483662
	CSSM_ALGID_SHA256WithECDSA   CssmAlgidAppleYarrow = 2147483673
	CSSM_ALGID_SHA256WithRSA     CssmAlgidAppleYarrow = 2147483668
	CSSM_ALGID_SHA384            CssmAlgidAppleYarrow = 2147483663
	CSSM_ALGID_SHA384WithECDSA   CssmAlgidAppleYarrow = 2147483674
	CSSM_ALGID_SHA384WithRSA     CssmAlgidAppleYarrow = 2147483669
	CSSM_ALGID_SHA512            CssmAlgidAppleYarrow = 2147483664
	CSSM_ALGID_SHA512WithECDSA   CssmAlgidAppleYarrow = 2147483675
	CSSM_ALGID_SHA512WithRSA     CssmAlgidAppleYarrow = 2147483670
	CSSM_ALGID__FIRST_UNUSED     CssmAlgidAppleYarrow = 2147483678
)

func (e CssmAlgidAppleYarrow) String() string {
	switch e {
	case CSSM_ALGID_AES:
		return "CSSM_ALGID_AES"
	case CSSM_ALGID_APPLE_YARROW:
		return "CSSM_ALGID_APPLE_YARROW"
	case CSSM_ALGID_ASC:
		return "CSSM_ALGID_ASC"
	case CSSM_ALGID_ECDH_X963_KDF:
		return "CSSM_ALGID_ECDH_X963_KDF"
	case CSSM_ALGID_ECDSA_SPECIFIED:
		return "CSSM_ALGID_ECDSA_SPECIFIED"
	case CSSM_ALGID_ENTROPY_DEFAULT:
		return "CSSM_ALGID_ENTROPY_DEFAULT"
	case CSSM_ALGID_FEE:
		return "CSSM_ALGID_FEE"
	case CSSM_ALGID_FEED:
		return "CSSM_ALGID_FEED"
	case CSSM_ALGID_FEEDEXP:
		return "CSSM_ALGID_FEEDEXP"
	case CSSM_ALGID_FEE_MD5:
		return "CSSM_ALGID_FEE_MD5"
	case CSSM_ALGID_FEE_SHA1:
		return "CSSM_ALGID_FEE_SHA1"
	case CSSM_ALGID_KEYCHAIN_KEY:
		return "CSSM_ALGID_KEYCHAIN_KEY"
	case CSSM_ALGID_OPENSSH1:
		return "CSSM_ALGID_OPENSSH1"
	case CSSM_ALGID_PBE_OPENSSL_MD5:
		return "CSSM_ALGID_PBE_OPENSSL_MD5"
	case CSSM_ALGID_PKCS12_PBE_ENCR:
		return "CSSM_ALGID_PKCS12_PBE_ENCR"
	case CSSM_ALGID_PKCS12_PBE_MAC:
		return "CSSM_ALGID_PKCS12_PBE_MAC"
	case CSSM_ALGID_SECURE_PASSPHRASE:
		return "CSSM_ALGID_SECURE_PASSPHRASE"
	case CSSM_ALGID_SHA1HMAC_LEGACY:
		return "CSSM_ALGID_SHA1HMAC_LEGACY"
	case CSSM_ALGID_SHA224:
		return "CSSM_ALGID_SHA224"
	case CSSM_ALGID_SHA224WithECDSA:
		return "CSSM_ALGID_SHA224WithECDSA"
	case CSSM_ALGID_SHA224WithRSA:
		return "CSSM_ALGID_SHA224WithRSA"
	case CSSM_ALGID_SHA256:
		return "CSSM_ALGID_SHA256"
	case CSSM_ALGID_SHA256WithECDSA:
		return "CSSM_ALGID_SHA256WithECDSA"
	case CSSM_ALGID_SHA256WithRSA:
		return "CSSM_ALGID_SHA256WithRSA"
	case CSSM_ALGID_SHA384:
		return "CSSM_ALGID_SHA384"
	case CSSM_ALGID_SHA384WithECDSA:
		return "CSSM_ALGID_SHA384WithECDSA"
	case CSSM_ALGID_SHA384WithRSA:
		return "CSSM_ALGID_SHA384WithRSA"
	case CSSM_ALGID_SHA512:
		return "CSSM_ALGID_SHA512"
	case CSSM_ALGID_SHA512WithECDSA:
		return "CSSM_ALGID_SHA512WithECDSA"
	case CSSM_ALGID_SHA512WithRSA:
		return "CSSM_ALGID_SHA512WithRSA"
	case CSSM_ALGID__FIRST_UNUSED:
		return "CSSM_ALGID__FIRST_UNUSED"
	default:
		return fmt.Sprintf("CssmAlgidAppleYarrow(%d)", e)
	}
}

type CssmAlgidNone uint32

const (
	CSSM_ALGID_3DES                CssmAlgidNone = 77
	CSSM_ALGID_3DES_1KEY           CssmAlgidNone = 20
	CSSM_ALGID_3DES_1KEY_EEE       CssmAlgidNone = 19
	CSSM_ALGID_3DES_2KEY           CssmAlgidNone = 18
	CSSM_ALGID_3DES_2KEY_EDE       CssmAlgidNone = 18
	CSSM_ALGID_3DES_2KEY_EEE       CssmAlgidNone = 21
	CSSM_ALGID_3DES_3KEY           CssmAlgidNone = 17
	CSSM_ALGID_3DES_3KEY_EDE       CssmAlgidNone = 17
	CSSM_ALGID_3DES_3KEY_EEE       CssmAlgidNone = 20
	CSSM_ALGID_BATON               CssmAlgidNone = 72
	CSSM_ALGID_BLOWFISH            CssmAlgidNone = 28
	CSSM_ALGID_CAST                CssmAlgidNone = 27
	CSSM_ALGID_CAST3               CssmAlgidNone = 53
	CSSM_ALGID_CAST5               CssmAlgidNone = 54
	CSSM_ALGID_CDMF                CssmAlgidNone = 52
	CSSM_ALGID_CRAB                CssmAlgidNone = 41
	CSSM_ALGID_CUSTOM              CssmAlgidNone = 1
	CSSM_ALGID_ConcatBaseAndData   CssmAlgidNone = 58
	CSSM_ALGID_ConcatBaseAndKey    CssmAlgidNone = 56
	CSSM_ALGID_ConcatDataAndBase   CssmAlgidNone = 59
	CSSM_ALGID_ConcatKeyAndBase    CssmAlgidNone = 57
	CSSM_ALGID_DES                 CssmAlgidNone = 14
	CSSM_ALGID_DESRandom           CssmAlgidNone = 50
	CSSM_ALGID_DESX                CssmAlgidNone = 15
	CSSM_ALGID_DH                  CssmAlgidNone = 2
	CSSM_ALGID_DSA                 CssmAlgidNone = 43
	CSSM_ALGID_DSA_BSAFE           CssmAlgidNone = 83
	CSSM_ALGID_ECAES               CssmAlgidNone = 90
	CSSM_ALGID_ECC                 CssmAlgidNone = 93
	CSSM_ALGID_ECDH                CssmAlgidNone = 84
	CSSM_ALGID_ECDSA               CssmAlgidNone = 73
	CSSM_ALGID_ECES                CssmAlgidNone = 89
	CSSM_ALGID_ECMQV               CssmAlgidNone = 85
	CSSM_ALGID_ECNRA               CssmAlgidNone = 87
	CSSM_ALGID_ElGamal             CssmAlgidNone = 46
	CSSM_ALGID_ExtractFromKey      CssmAlgidNone = 61
	CSSM_ALGID_FASTHASH            CssmAlgidNone = 76
	CSSM_ALGID_FEAL                CssmAlgidNone = 32
	CSSM_ALGID_FIPS186Random       CssmAlgidNone = 92
	CSSM_ALGID_FortezzaTimestamp   CssmAlgidNone = 80
	CSSM_ALGID_GOST                CssmAlgidNone = 39
	CSSM_ALGID_GenericSecret       CssmAlgidNone = 55
	CSSM_ALGID_HAVAL               CssmAlgidNone = 10
	CSSM_ALGID_HAVAL3              CssmAlgidNone = 98
	CSSM_ALGID_HAVAL4              CssmAlgidNone = 99
	CSSM_ALGID_HAVAL5              CssmAlgidNone = 100
	CSSM_ALGID_IBCHASH             CssmAlgidNone = 12
	CSSM_ALGID_IDEA                CssmAlgidNone = 22
	CSSM_ALGID_IntelPlatformRandom CssmAlgidNone = 96
	CSSM_ALGID_JUNIPER             CssmAlgidNone = 75
	CSSM_ALGID_KEA                 CssmAlgidNone = 4
	CSSM_ALGID_KHAFRE              CssmAlgidNone = 37
	CSSM_ALGID_KHUFU               CssmAlgidNone = 36
	CSSM_ALGID_LAST                CssmAlgidNone = 2147483647
	CSSM_ALGID_LOKI                CssmAlgidNone = 35
	CSSM_ALGID_LUCIFER             CssmAlgidNone = 30
	CSSM_ALGID_MADRYGA             CssmAlgidNone = 31
	CSSM_ALGID_MAYFLY              CssmAlgidNone = 74
	CSSM_ALGID_MD2                 CssmAlgidNone = 5
	CSSM_ALGID_MD2Random           CssmAlgidNone = 47
	CSSM_ALGID_MD2WithRSA          CssmAlgidNone = 45
	CSSM_ALGID_MD4                 CssmAlgidNone = 6
	CSSM_ALGID_MD5                 CssmAlgidNone = 7
	CSSM_ALGID_MD5HMAC             CssmAlgidNone = 102
	CSSM_ALGID_MD5Random           CssmAlgidNone = 48
	CSSM_ALGID_MD5WithRSA          CssmAlgidNone = 44
	CSSM_ALGID_MMB                 CssmAlgidNone = 38
	CSSM_ALGID_MQV                 CssmAlgidNone = 94
	CSSM_ALGID_NHASH               CssmAlgidNone = 9
	CSSM_ALGID_NONE                CssmAlgidNone = 0
	CSSM_ALGID_NRA                 CssmAlgidNone = 95
	CSSM_ALGID_PH                  CssmAlgidNone = 3
	CSSM_ALGID_PKCS12_SHA1_PBE     CssmAlgidNone = 86
	CSSM_ALGID_PKCS5_PBKDF1_MD2    CssmAlgidNone = 68
	CSSM_ALGID_PKCS5_PBKDF1_MD5    CssmAlgidNone = 67
	CSSM_ALGID_PKCS5_PBKDF1_SHA1   CssmAlgidNone = 69
	CSSM_ALGID_PKCS5_PBKDF2        CssmAlgidNone = 103
	CSSM_ALGID_RC2                 CssmAlgidNone = 23
	CSSM_ALGID_RC4                 CssmAlgidNone = 25
	CSSM_ALGID_RC5                 CssmAlgidNone = 24
	CSSM_ALGID_RDES                CssmAlgidNone = 16
	CSSM_ALGID_REDOC               CssmAlgidNone = 33
	CSSM_ALGID_REDOC3              CssmAlgidNone = 34
	CSSM_ALGID_RIPEMAC             CssmAlgidNone = 13
	CSSM_ALGID_RIPEMD              CssmAlgidNone = 11
	CSSM_ALGID_RSA                 CssmAlgidNone = 42
	CSSM_ALGID_RUNNING_COUNTER     CssmAlgidNone = 104
	CSSM_ALGID_SAFER               CssmAlgidNone = 40
	CSSM_ALGID_SEAL                CssmAlgidNone = 26
	CSSM_ALGID_SHA1                CssmAlgidNone = 8
	CSSM_ALGID_SHA1HMAC            CssmAlgidNone = 91
	CSSM_ALGID_SHA1WithDSA         CssmAlgidNone = 81
	CSSM_ALGID_SHA1WithECDSA       CssmAlgidNone = 82
	CSSM_ALGID_SHA1WithECNRA       CssmAlgidNone = 88
	CSSM_ALGID_SHA1WithRSA         CssmAlgidNone = 51
	CSSM_ALGID_SHARandom           CssmAlgidNone = 49
	CSSM_ALGID_SKIPJACK            CssmAlgidNone = 29
	CSSM_ALGID_SSL3KeyAndMacDerive CssmAlgidNone = 64
	CSSM_ALGID_SSL3MD5             CssmAlgidNone = 78
	CSSM_ALGID_SSL3MD5_MAC         CssmAlgidNone = 65
	CSSM_ALGID_SSL3PrePrimaryGen   CssmAlgidNone = 62
	CSSM_ALGID_SSL3PrimaryDerive   CssmAlgidNone = 63
	CSSM_ALGID_SSL3SHA1            CssmAlgidNone = 79
	CSSM_ALGID_SSL3SHA1_MAC        CssmAlgidNone = 66
	CSSM_ALGID_TIGER               CssmAlgidNone = 101
	CSSM_ALGID_UTC                 CssmAlgidNone = 97
	CSSM_ALGID_VENDOR_DEFINED      CssmAlgidNone = 2147483648
	CSSM_ALGID_WrapLynks           CssmAlgidNone = 70
	CSSM_ALGID_WrapSET_OAEP        CssmAlgidNone = 71
	CSSM_ALGID_XORBaseAndData      CssmAlgidNone = 60
	// Deprecated: use CSSM_ALGID_SSL3PrimaryDerive.
	CSSM_ALGID_SSL3MasterDerive CssmAlgidNone = 63
	// Deprecated: use CSSM_ALGID_SSL3PrePrimaryGen.
	CSSM_ALGID_SSL3PreMasterGen CssmAlgidNone = 62
)

func (e CssmAlgidNone) String() string {
	switch e {
	case CSSM_ALGID_3DES:
		return "CSSM_ALGID_3DES"
	case CSSM_ALGID_3DES_1KEY:
		return "CSSM_ALGID_3DES_1KEY"
	case CSSM_ALGID_3DES_1KEY_EEE:
		return "CSSM_ALGID_3DES_1KEY_EEE"
	case CSSM_ALGID_3DES_2KEY:
		return "CSSM_ALGID_3DES_2KEY"
	case CSSM_ALGID_3DES_2KEY_EEE:
		return "CSSM_ALGID_3DES_2KEY_EEE"
	case CSSM_ALGID_3DES_3KEY:
		return "CSSM_ALGID_3DES_3KEY"
	case CSSM_ALGID_BATON:
		return "CSSM_ALGID_BATON"
	case CSSM_ALGID_BLOWFISH:
		return "CSSM_ALGID_BLOWFISH"
	case CSSM_ALGID_CAST:
		return "CSSM_ALGID_CAST"
	case CSSM_ALGID_CAST3:
		return "CSSM_ALGID_CAST3"
	case CSSM_ALGID_CAST5:
		return "CSSM_ALGID_CAST5"
	case CSSM_ALGID_CDMF:
		return "CSSM_ALGID_CDMF"
	case CSSM_ALGID_CRAB:
		return "CSSM_ALGID_CRAB"
	case CSSM_ALGID_CUSTOM:
		return "CSSM_ALGID_CUSTOM"
	case CSSM_ALGID_ConcatBaseAndData:
		return "CSSM_ALGID_ConcatBaseAndData"
	case CSSM_ALGID_ConcatBaseAndKey:
		return "CSSM_ALGID_ConcatBaseAndKey"
	case CSSM_ALGID_ConcatDataAndBase:
		return "CSSM_ALGID_ConcatDataAndBase"
	case CSSM_ALGID_ConcatKeyAndBase:
		return "CSSM_ALGID_ConcatKeyAndBase"
	case CSSM_ALGID_DES:
		return "CSSM_ALGID_DES"
	case CSSM_ALGID_DESRandom:
		return "CSSM_ALGID_DESRandom"
	case CSSM_ALGID_DESX:
		return "CSSM_ALGID_DESX"
	case CSSM_ALGID_DH:
		return "CSSM_ALGID_DH"
	case CSSM_ALGID_DSA:
		return "CSSM_ALGID_DSA"
	case CSSM_ALGID_DSA_BSAFE:
		return "CSSM_ALGID_DSA_BSAFE"
	case CSSM_ALGID_ECAES:
		return "CSSM_ALGID_ECAES"
	case CSSM_ALGID_ECC:
		return "CSSM_ALGID_ECC"
	case CSSM_ALGID_ECDH:
		return "CSSM_ALGID_ECDH"
	case CSSM_ALGID_ECDSA:
		return "CSSM_ALGID_ECDSA"
	case CSSM_ALGID_ECES:
		return "CSSM_ALGID_ECES"
	case CSSM_ALGID_ECMQV:
		return "CSSM_ALGID_ECMQV"
	case CSSM_ALGID_ECNRA:
		return "CSSM_ALGID_ECNRA"
	case CSSM_ALGID_ElGamal:
		return "CSSM_ALGID_ElGamal"
	case CSSM_ALGID_ExtractFromKey:
		return "CSSM_ALGID_ExtractFromKey"
	case CSSM_ALGID_FASTHASH:
		return "CSSM_ALGID_FASTHASH"
	case CSSM_ALGID_FEAL:
		return "CSSM_ALGID_FEAL"
	case CSSM_ALGID_FIPS186Random:
		return "CSSM_ALGID_FIPS186Random"
	case CSSM_ALGID_FortezzaTimestamp:
		return "CSSM_ALGID_FortezzaTimestamp"
	case CSSM_ALGID_GOST:
		return "CSSM_ALGID_GOST"
	case CSSM_ALGID_GenericSecret:
		return "CSSM_ALGID_GenericSecret"
	case CSSM_ALGID_HAVAL:
		return "CSSM_ALGID_HAVAL"
	case CSSM_ALGID_HAVAL3:
		return "CSSM_ALGID_HAVAL3"
	case CSSM_ALGID_HAVAL4:
		return "CSSM_ALGID_HAVAL4"
	case CSSM_ALGID_HAVAL5:
		return "CSSM_ALGID_HAVAL5"
	case CSSM_ALGID_IBCHASH:
		return "CSSM_ALGID_IBCHASH"
	case CSSM_ALGID_IDEA:
		return "CSSM_ALGID_IDEA"
	case CSSM_ALGID_IntelPlatformRandom:
		return "CSSM_ALGID_IntelPlatformRandom"
	case CSSM_ALGID_JUNIPER:
		return "CSSM_ALGID_JUNIPER"
	case CSSM_ALGID_KEA:
		return "CSSM_ALGID_KEA"
	case CSSM_ALGID_KHAFRE:
		return "CSSM_ALGID_KHAFRE"
	case CSSM_ALGID_KHUFU:
		return "CSSM_ALGID_KHUFU"
	case CSSM_ALGID_LAST:
		return "CSSM_ALGID_LAST"
	case CSSM_ALGID_LOKI:
		return "CSSM_ALGID_LOKI"
	case CSSM_ALGID_LUCIFER:
		return "CSSM_ALGID_LUCIFER"
	case CSSM_ALGID_MADRYGA:
		return "CSSM_ALGID_MADRYGA"
	case CSSM_ALGID_MAYFLY:
		return "CSSM_ALGID_MAYFLY"
	case CSSM_ALGID_MD2:
		return "CSSM_ALGID_MD2"
	case CSSM_ALGID_MD2Random:
		return "CSSM_ALGID_MD2Random"
	case CSSM_ALGID_MD2WithRSA:
		return "CSSM_ALGID_MD2WithRSA"
	case CSSM_ALGID_MD4:
		return "CSSM_ALGID_MD4"
	case CSSM_ALGID_MD5:
		return "CSSM_ALGID_MD5"
	case CSSM_ALGID_MD5HMAC:
		return "CSSM_ALGID_MD5HMAC"
	case CSSM_ALGID_MD5Random:
		return "CSSM_ALGID_MD5Random"
	case CSSM_ALGID_MD5WithRSA:
		return "CSSM_ALGID_MD5WithRSA"
	case CSSM_ALGID_MMB:
		return "CSSM_ALGID_MMB"
	case CSSM_ALGID_MQV:
		return "CSSM_ALGID_MQV"
	case CSSM_ALGID_NHASH:
		return "CSSM_ALGID_NHASH"
	case CSSM_ALGID_NONE:
		return "CSSM_ALGID_NONE"
	case CSSM_ALGID_NRA:
		return "CSSM_ALGID_NRA"
	case CSSM_ALGID_PH:
		return "CSSM_ALGID_PH"
	case CSSM_ALGID_PKCS12_SHA1_PBE:
		return "CSSM_ALGID_PKCS12_SHA1_PBE"
	case CSSM_ALGID_PKCS5_PBKDF1_MD2:
		return "CSSM_ALGID_PKCS5_PBKDF1_MD2"
	case CSSM_ALGID_PKCS5_PBKDF1_MD5:
		return "CSSM_ALGID_PKCS5_PBKDF1_MD5"
	case CSSM_ALGID_PKCS5_PBKDF1_SHA1:
		return "CSSM_ALGID_PKCS5_PBKDF1_SHA1"
	case CSSM_ALGID_PKCS5_PBKDF2:
		return "CSSM_ALGID_PKCS5_PBKDF2"
	case CSSM_ALGID_RC2:
		return "CSSM_ALGID_RC2"
	case CSSM_ALGID_RC4:
		return "CSSM_ALGID_RC4"
	case CSSM_ALGID_RC5:
		return "CSSM_ALGID_RC5"
	case CSSM_ALGID_RDES:
		return "CSSM_ALGID_RDES"
	case CSSM_ALGID_REDOC:
		return "CSSM_ALGID_REDOC"
	case CSSM_ALGID_REDOC3:
		return "CSSM_ALGID_REDOC3"
	case CSSM_ALGID_RIPEMAC:
		return "CSSM_ALGID_RIPEMAC"
	case CSSM_ALGID_RIPEMD:
		return "CSSM_ALGID_RIPEMD"
	case CSSM_ALGID_RSA:
		return "CSSM_ALGID_RSA"
	case CSSM_ALGID_RUNNING_COUNTER:
		return "CSSM_ALGID_RUNNING_COUNTER"
	case CSSM_ALGID_SAFER:
		return "CSSM_ALGID_SAFER"
	case CSSM_ALGID_SEAL:
		return "CSSM_ALGID_SEAL"
	case CSSM_ALGID_SHA1:
		return "CSSM_ALGID_SHA1"
	case CSSM_ALGID_SHA1HMAC:
		return "CSSM_ALGID_SHA1HMAC"
	case CSSM_ALGID_SHA1WithDSA:
		return "CSSM_ALGID_SHA1WithDSA"
	case CSSM_ALGID_SHA1WithECDSA:
		return "CSSM_ALGID_SHA1WithECDSA"
	case CSSM_ALGID_SHA1WithECNRA:
		return "CSSM_ALGID_SHA1WithECNRA"
	case CSSM_ALGID_SHA1WithRSA:
		return "CSSM_ALGID_SHA1WithRSA"
	case CSSM_ALGID_SHARandom:
		return "CSSM_ALGID_SHARandom"
	case CSSM_ALGID_SKIPJACK:
		return "CSSM_ALGID_SKIPJACK"
	case CSSM_ALGID_SSL3KeyAndMacDerive:
		return "CSSM_ALGID_SSL3KeyAndMacDerive"
	case CSSM_ALGID_SSL3MD5:
		return "CSSM_ALGID_SSL3MD5"
	case CSSM_ALGID_SSL3MD5_MAC:
		return "CSSM_ALGID_SSL3MD5_MAC"
	case CSSM_ALGID_SSL3PrePrimaryGen:
		return "CSSM_ALGID_SSL3PrePrimaryGen"
	case CSSM_ALGID_SSL3PrimaryDerive:
		return "CSSM_ALGID_SSL3PrimaryDerive"
	case CSSM_ALGID_SSL3SHA1:
		return "CSSM_ALGID_SSL3SHA1"
	case CSSM_ALGID_SSL3SHA1_MAC:
		return "CSSM_ALGID_SSL3SHA1_MAC"
	case CSSM_ALGID_TIGER:
		return "CSSM_ALGID_TIGER"
	case CSSM_ALGID_UTC:
		return "CSSM_ALGID_UTC"
	case CSSM_ALGID_VENDOR_DEFINED:
		return "CSSM_ALGID_VENDOR_DEFINED"
	case CSSM_ALGID_WrapLynks:
		return "CSSM_ALGID_WrapLynks"
	case CSSM_ALGID_WrapSET_OAEP:
		return "CSSM_ALGID_WrapSET_OAEP"
	case CSSM_ALGID_XORBaseAndData:
		return "CSSM_ALGID_XORBaseAndData"
	default:
		return fmt.Sprintf("CssmAlgidNone(%d)", e)
	}
}

type CssmAlgmode uint32

const (
	CSSM_ALGMODE_BC             CssmAlgmode = 14
	CSSM_ALGMODE_CBC            CssmAlgmode = 4
	CSSM_ALGMODE_CBC128         CssmAlgmode = 36
	CSSM_ALGMODE_CBC64          CssmAlgmode = 25
	CSSM_ALGMODE_CBCC           CssmAlgmode = 16
	CSSM_ALGMODE_CBCPD          CssmAlgmode = 20
	CSSM_ALGMODE_CBCPadIV8      CssmAlgmode = 6
	CSSM_ALGMODE_CBC_IV8        CssmAlgmode = 5
	CSSM_ALGMODE_CFB            CssmAlgmode = 7
	CSSM_ALGMODE_CFB16          CssmAlgmode = 29
	CSSM_ALGMODE_CFB32          CssmAlgmode = 28
	CSSM_ALGMODE_CFB8           CssmAlgmode = 30
	CSSM_ALGMODE_CFBPadIV8      CssmAlgmode = 9
	CSSM_ALGMODE_CFB_IV8        CssmAlgmode = 8
	CSSM_ALGMODE_COUNTER        CssmAlgmode = 13
	CSSM_ALGMODE_CUSTOM         CssmAlgmode = 1
	CSSM_ALGMODE_ECB            CssmAlgmode = 2
	CSSM_ALGMODE_ECB128         CssmAlgmode = 34
	CSSM_ALGMODE_ECB64          CssmAlgmode = 24
	CSSM_ALGMODE_ECB96          CssmAlgmode = 35
	CSSM_ALGMODE_ECBPad         CssmAlgmode = 3
	CSSM_ALGMODE_ISO_9796       CssmAlgmode = 41
	CSSM_ALGMODE_LAST           CssmAlgmode = 2147483647
	CSSM_ALGMODE_NONE           CssmAlgmode = 0
	CSSM_ALGMODE_OAEP_HASH      CssmAlgmode = 37
	CSSM_ALGMODE_OFB            CssmAlgmode = 10
	CSSM_ALGMODE_OFB64          CssmAlgmode = 26
	CSSM_ALGMODE_OFBNLF         CssmAlgmode = 17
	CSSM_ALGMODE_OFBPadIV8      CssmAlgmode = 12
	CSSM_ALGMODE_OFB_IV8        CssmAlgmode = 11
	CSSM_ALGMODE_PBC            CssmAlgmode = 18
	CSSM_ALGMODE_PCBC           CssmAlgmode = 15
	CSSM_ALGMODE_PFB            CssmAlgmode = 19
	CSSM_ALGMODE_PKCS1_EME_OAEP CssmAlgmode = 39
	CSSM_ALGMODE_PKCS1_EME_V15  CssmAlgmode = 38
	CSSM_ALGMODE_PKCS1_EMSA_V15 CssmAlgmode = 40
	CSSM_ALGMODE_PRIVATE_KEY    CssmAlgmode = 22
	CSSM_ALGMODE_PRIVATE_WRAP   CssmAlgmode = 32
	CSSM_ALGMODE_PUBLIC_KEY     CssmAlgmode = 21
	CSSM_ALGMODE_RELAYX         CssmAlgmode = 33
	CSSM_ALGMODE_SHUFFLE        CssmAlgmode = 23
	CSSM_ALGMODE_VENDOR_DEFINED CssmAlgmode = 2147483648
	CSSM_ALGMODE_WRAP           CssmAlgmode = 31
	CSSM_ALGMODE_X9_31          CssmAlgmode = 42
)

func (e CssmAlgmode) String() string {
	switch e {
	case CSSM_ALGMODE_BC:
		return "CSSM_ALGMODE_BC"
	case CSSM_ALGMODE_CBC:
		return "CSSM_ALGMODE_CBC"
	case CSSM_ALGMODE_CBC128:
		return "CSSM_ALGMODE_CBC128"
	case CSSM_ALGMODE_CBC64:
		return "CSSM_ALGMODE_CBC64"
	case CSSM_ALGMODE_CBCC:
		return "CSSM_ALGMODE_CBCC"
	case CSSM_ALGMODE_CBCPD:
		return "CSSM_ALGMODE_CBCPD"
	case CSSM_ALGMODE_CBCPadIV8:
		return "CSSM_ALGMODE_CBCPadIV8"
	case CSSM_ALGMODE_CBC_IV8:
		return "CSSM_ALGMODE_CBC_IV8"
	case CSSM_ALGMODE_CFB:
		return "CSSM_ALGMODE_CFB"
	case CSSM_ALGMODE_CFB16:
		return "CSSM_ALGMODE_CFB16"
	case CSSM_ALGMODE_CFB32:
		return "CSSM_ALGMODE_CFB32"
	case CSSM_ALGMODE_CFB8:
		return "CSSM_ALGMODE_CFB8"
	case CSSM_ALGMODE_CFBPadIV8:
		return "CSSM_ALGMODE_CFBPadIV8"
	case CSSM_ALGMODE_CFB_IV8:
		return "CSSM_ALGMODE_CFB_IV8"
	case CSSM_ALGMODE_COUNTER:
		return "CSSM_ALGMODE_COUNTER"
	case CSSM_ALGMODE_CUSTOM:
		return "CSSM_ALGMODE_CUSTOM"
	case CSSM_ALGMODE_ECB:
		return "CSSM_ALGMODE_ECB"
	case CSSM_ALGMODE_ECB128:
		return "CSSM_ALGMODE_ECB128"
	case CSSM_ALGMODE_ECB64:
		return "CSSM_ALGMODE_ECB64"
	case CSSM_ALGMODE_ECB96:
		return "CSSM_ALGMODE_ECB96"
	case CSSM_ALGMODE_ECBPad:
		return "CSSM_ALGMODE_ECBPad"
	case CSSM_ALGMODE_ISO_9796:
		return "CSSM_ALGMODE_ISO_9796"
	case CSSM_ALGMODE_LAST:
		return "CSSM_ALGMODE_LAST"
	case CSSM_ALGMODE_NONE:
		return "CSSM_ALGMODE_NONE"
	case CSSM_ALGMODE_OAEP_HASH:
		return "CSSM_ALGMODE_OAEP_HASH"
	case CSSM_ALGMODE_OFB:
		return "CSSM_ALGMODE_OFB"
	case CSSM_ALGMODE_OFB64:
		return "CSSM_ALGMODE_OFB64"
	case CSSM_ALGMODE_OFBNLF:
		return "CSSM_ALGMODE_OFBNLF"
	case CSSM_ALGMODE_OFBPadIV8:
		return "CSSM_ALGMODE_OFBPadIV8"
	case CSSM_ALGMODE_OFB_IV8:
		return "CSSM_ALGMODE_OFB_IV8"
	case CSSM_ALGMODE_PBC:
		return "CSSM_ALGMODE_PBC"
	case CSSM_ALGMODE_PCBC:
		return "CSSM_ALGMODE_PCBC"
	case CSSM_ALGMODE_PFB:
		return "CSSM_ALGMODE_PFB"
	case CSSM_ALGMODE_PKCS1_EME_OAEP:
		return "CSSM_ALGMODE_PKCS1_EME_OAEP"
	case CSSM_ALGMODE_PKCS1_EME_V15:
		return "CSSM_ALGMODE_PKCS1_EME_V15"
	case CSSM_ALGMODE_PKCS1_EMSA_V15:
		return "CSSM_ALGMODE_PKCS1_EMSA_V15"
	case CSSM_ALGMODE_PRIVATE_KEY:
		return "CSSM_ALGMODE_PRIVATE_KEY"
	case CSSM_ALGMODE_PRIVATE_WRAP:
		return "CSSM_ALGMODE_PRIVATE_WRAP"
	case CSSM_ALGMODE_PUBLIC_KEY:
		return "CSSM_ALGMODE_PUBLIC_KEY"
	case CSSM_ALGMODE_RELAYX:
		return "CSSM_ALGMODE_RELAYX"
	case CSSM_ALGMODE_SHUFFLE:
		return "CSSM_ALGMODE_SHUFFLE"
	case CSSM_ALGMODE_VENDOR_DEFINED:
		return "CSSM_ALGMODE_VENDOR_DEFINED"
	case CSSM_ALGMODE_WRAP:
		return "CSSM_ALGMODE_WRAP"
	case CSSM_ALGMODE_X9_31:
		return "CSSM_ALGMODE_X9_31"
	default:
		return fmt.Sprintf("CssmAlgmode(%d)", e)
	}
}

type CssmApple uint32

const (
	CSSM_APPLECSPDL_DB_CHANGE_PASSWORD CssmApple = 5
	CSSM_APPLECSPDL_DB_GET_HANDLE      CssmApple = 6
	CSSM_APPLECSPDL_DB_GET_SETTINGS    CssmApple = 2
	CSSM_APPLECSPDL_DB_IS_LOCKED       CssmApple = 4
	CSSM_APPLECSPDL_DB_LOCK            CssmApple = 0
	CSSM_APPLECSPDL_DB_SET_SETTINGS    CssmApple = 3
	CSSM_APPLECSPDL_DB_UNLOCK          CssmApple = 1
	CSSM_APPLECSP_KEYDIGEST            CssmApple = 0x100
	CSSM_APPLECSP_PUBKEY               CssmApple = 0x101
	CSSM_APPLESCPDL_CSP_GET_KEYHANDLE  CssmApple = 7
	CSSM_APPLE_PRIVATE_CSPDL_CODE_10   CssmApple = 10
	CSSM_APPLE_PRIVATE_CSPDL_CODE_11   CssmApple = 11
	CSSM_APPLE_PRIVATE_CSPDL_CODE_12   CssmApple = 12
	CSSM_APPLE_PRIVATE_CSPDL_CODE_13   CssmApple = 13
	CSSM_APPLE_PRIVATE_CSPDL_CODE_14   CssmApple = 14
	CSSM_APPLE_PRIVATE_CSPDL_CODE_15   CssmApple = 15
	CSSM_APPLE_PRIVATE_CSPDL_CODE_16   CssmApple = 16
	CSSM_APPLE_PRIVATE_CSPDL_CODE_17   CssmApple = 17
	CSSM_APPLE_PRIVATE_CSPDL_CODE_18   CssmApple = 18
	CSSM_APPLE_PRIVATE_CSPDL_CODE_19   CssmApple = 19
	CSSM_APPLE_PRIVATE_CSPDL_CODE_20   CssmApple = 20
	CSSM_APPLE_PRIVATE_CSPDL_CODE_21   CssmApple = 21
	CSSM_APPLE_PRIVATE_CSPDL_CODE_22   CssmApple = 22
	CSSM_APPLE_PRIVATE_CSPDL_CODE_23   CssmApple = 23
	CSSM_APPLE_PRIVATE_CSPDL_CODE_24   CssmApple = 24
	CSSM_APPLE_PRIVATE_CSPDL_CODE_25   CssmApple = 25
	CSSM_APPLE_PRIVATE_CSPDL_CODE_26   CssmApple = 26
	CSSM_APPLE_PRIVATE_CSPDL_CODE_27   CssmApple = 27
	CSSM_APPLE_PRIVATE_CSPDL_CODE_28   CssmApple = 28
	CSSM_APPLE_PRIVATE_CSPDL_CODE_8    CssmApple = 8
	CSSM_APPLE_PRIVATE_CSPDL_CODE_9    CssmApple = 9
)

func (e CssmApple) String() string {
	switch e {
	case CSSM_APPLECSPDL_DB_CHANGE_PASSWORD:
		return "CSSM_APPLECSPDL_DB_CHANGE_PASSWORD"
	case CSSM_APPLECSPDL_DB_GET_HANDLE:
		return "CSSM_APPLECSPDL_DB_GET_HANDLE"
	case CSSM_APPLECSPDL_DB_GET_SETTINGS:
		return "CSSM_APPLECSPDL_DB_GET_SETTINGS"
	case CSSM_APPLECSPDL_DB_IS_LOCKED:
		return "CSSM_APPLECSPDL_DB_IS_LOCKED"
	case CSSM_APPLECSPDL_DB_LOCK:
		return "CSSM_APPLECSPDL_DB_LOCK"
	case CSSM_APPLECSPDL_DB_SET_SETTINGS:
		return "CSSM_APPLECSPDL_DB_SET_SETTINGS"
	case CSSM_APPLECSPDL_DB_UNLOCK:
		return "CSSM_APPLECSPDL_DB_UNLOCK"
	case CSSM_APPLECSP_KEYDIGEST:
		return "CSSM_APPLECSP_KEYDIGEST"
	case CSSM_APPLECSP_PUBKEY:
		return "CSSM_APPLECSP_PUBKEY"
	case CSSM_APPLESCPDL_CSP_GET_KEYHANDLE:
		return "CSSM_APPLESCPDL_CSP_GET_KEYHANDLE"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_10:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_10"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_11:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_11"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_12:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_12"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_13:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_13"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_14:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_14"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_15:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_15"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_16:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_16"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_17:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_17"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_18:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_18"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_19:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_19"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_20:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_20"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_21:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_21"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_22:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_22"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_23:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_23"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_24:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_24"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_25:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_25"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_26:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_26"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_27:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_27"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_28:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_28"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_8:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_8"
	case CSSM_APPLE_PRIVATE_CSPDL_CODE_9:
		return "CSSM_APPLE_PRIVATE_CSPDL_CODE_9"
	default:
		return fmt.Sprintf("CssmApple(%d)", e)
	}
}

type CssmAppleUnlockType uint32

const (
	CSSM_APPLE_UNLOCK_TYPE_KEYBAG          CssmAppleUnlockType = 3
	CSSM_APPLE_UNLOCK_TYPE_KEY_DIRECT      CssmAppleUnlockType = 1
	CSSM_APPLE_UNLOCK_TYPE_WRAPPED_PRIVATE CssmAppleUnlockType = 2
)

func (e CssmAppleUnlockType) String() string {
	switch e {
	case CSSM_APPLE_UNLOCK_TYPE_KEYBAG:
		return "CSSM_APPLE_UNLOCK_TYPE_KEYBAG"
	case CSSM_APPLE_UNLOCK_TYPE_KEY_DIRECT:
		return "CSSM_APPLE_UNLOCK_TYPE_KEY_DIRECT"
	case CSSM_APPLE_UNLOCK_TYPE_WRAPPED_PRIVATE:
		return "CSSM_APPLE_UNLOCK_TYPE_WRAPPED_PRIVATE"
	default:
		return fmt.Sprintf("CssmAppleUnlockType(%d)", e)
	}
}

type CssmAppledlOpenParameters uint32

const (
	CSSM_APPLEDL_OPEN_PARAMETERS_VERSION CssmAppledlOpenParameters = 1
)

func (e CssmAppledlOpenParameters) String() string {
	switch e {
	case CSSM_APPLEDL_OPEN_PARAMETERS_VERSION:
		return "CSSM_APPLEDL_OPEN_PARAMETERS_VERSION"
	default:
		return fmt.Sprintf("CssmAppledlOpenParameters(%d)", e)
	}
}

type CssmApplefiledl uint32

const (
	CSSM_APPLEFILEDL_COMMIT            CssmApplefiledl = 1
	CSSM_APPLEFILEDL_DELETE_FILE       CssmApplefiledl = 6
	CSSM_APPLEFILEDL_MAKE_BACKUP       CssmApplefiledl = 4
	CSSM_APPLEFILEDL_MAKE_COPY         CssmApplefiledl = 5
	CSSM_APPLEFILEDL_ROLLBACK          CssmApplefiledl = 2
	CSSM_APPLEFILEDL_TAKE_FILE_LOCK    CssmApplefiledl = 3
	CSSM_APPLEFILEDL_TOGGLE_AUTOCOMMIT CssmApplefiledl = 0
)

func (e CssmApplefiledl) String() string {
	switch e {
	case CSSM_APPLEFILEDL_COMMIT:
		return "CSSM_APPLEFILEDL_COMMIT"
	case CSSM_APPLEFILEDL_DELETE_FILE:
		return "CSSM_APPLEFILEDL_DELETE_FILE"
	case CSSM_APPLEFILEDL_MAKE_BACKUP:
		return "CSSM_APPLEFILEDL_MAKE_BACKUP"
	case CSSM_APPLEFILEDL_MAKE_COPY:
		return "CSSM_APPLEFILEDL_MAKE_COPY"
	case CSSM_APPLEFILEDL_ROLLBACK:
		return "CSSM_APPLEFILEDL_ROLLBACK"
	case CSSM_APPLEFILEDL_TAKE_FILE_LOCK:
		return "CSSM_APPLEFILEDL_TAKE_FILE_LOCK"
	case CSSM_APPLEFILEDL_TOGGLE_AUTOCOMMIT:
		return "CSSM_APPLEFILEDL_TOGGLE_AUTOCOMMIT"
	default:
		return fmt.Sprintf("CssmApplefiledl(%d)", e)
	}
}

type CssmApplex509cl uint32

const (
	CSSM_APPLEX509CL_OBTAIN_CSR CssmApplex509cl = 0
	CSSM_APPLEX509CL_VERIFY_CSR CssmApplex509cl = 1
)

func (e CssmApplex509cl) String() string {
	switch e {
	case CSSM_APPLEX509CL_OBTAIN_CSR:
		return "CSSM_APPLEX509CL_OBTAIN_CSR"
	case CSSM_APPLEX509CL_VERIFY_CSR:
		return "CSSM_APPLEX509CL_VERIFY_CSR"
	default:
		return fmt.Sprintf("CssmApplex509cl(%d)", e)
	}
}

type CssmAscOptimize uint32

const (
	CSSM_ASC_OPTIMIZE_ASCII     CssmAscOptimize = 5
	CSSM_ASC_OPTIMIZE_DEFAULT   CssmAscOptimize = 0
	CSSM_ASC_OPTIMIZE_SECURITY  CssmAscOptimize = 2
	CSSM_ASC_OPTIMIZE_SIZE      CssmAscOptimize = 1
	CSSM_ASC_OPTIMIZE_TIME      CssmAscOptimize = 3
	CSSM_ASC_OPTIMIZE_TIME_SIZE CssmAscOptimize = 4
)

func (e CssmAscOptimize) String() string {
	switch e {
	case CSSM_ASC_OPTIMIZE_ASCII:
		return "CSSM_ASC_OPTIMIZE_ASCII"
	case CSSM_ASC_OPTIMIZE_DEFAULT:
		return "CSSM_ASC_OPTIMIZE_DEFAULT"
	case CSSM_ASC_OPTIMIZE_SECURITY:
		return "CSSM_ASC_OPTIMIZE_SECURITY"
	case CSSM_ASC_OPTIMIZE_SIZE:
		return "CSSM_ASC_OPTIMIZE_SIZE"
	case CSSM_ASC_OPTIMIZE_TIME:
		return "CSSM_ASC_OPTIMIZE_TIME"
	case CSSM_ASC_OPTIMIZE_TIME_SIZE:
		return "CSSM_ASC_OPTIMIZE_TIME_SIZE"
	default:
		return fmt.Sprintf("CssmAscOptimize(%d)", e)
	}
}

type CssmAttachRead uint32

const (
	CSSM_ATTACH_READ_ONLY CssmAttachRead = 0x1
)

func (e CssmAttachRead) String() string {
	switch e {
	case CSSM_ATTACH_READ_ONLY:
		return "CSSM_ATTACH_READ_ONLY"
	default:
		return fmt.Sprintf("CssmAttachRead(%d)", e)
	}
}

type CssmAttributeDataNone uint32

const (
	CSSM_ATTRIBUTE_DATA_ACCESS_CREDENTIALS CssmAttributeDataNone = 0x80000000
	CSSM_ATTRIBUTE_DATA_CRYPTO_DATA        CssmAttributeDataNone = 0x30000000
	CSSM_ATTRIBUTE_DATA_CSSM_DATA          CssmAttributeDataNone = 0x20000000
	CSSM_ATTRIBUTE_DATA_DATE               CssmAttributeDataNone = 0x60000000
	CSSM_ATTRIBUTE_DATA_DL_DB_HANDLE       CssmAttributeDataNone = 0x2000000
	CSSM_ATTRIBUTE_DATA_KEY                CssmAttributeDataNone = 0x40000000
	CSSM_ATTRIBUTE_DATA_KR_PROFILE         CssmAttributeDataNone = 0x3000000
	CSSM_ATTRIBUTE_DATA_NONE               CssmAttributeDataNone = 0
	CSSM_ATTRIBUTE_DATA_RANGE              CssmAttributeDataNone = 0x70000000
	CSSM_ATTRIBUTE_DATA_STRING             CssmAttributeDataNone = 0x50000000
	CSSM_ATTRIBUTE_DATA_UINT32             CssmAttributeDataNone = 0x10000000
	CSSM_ATTRIBUTE_DATA_VERSION            CssmAttributeDataNone = 0x1000000
	CSSM_ATTRIBUTE_TYPE_MASK               CssmAttributeDataNone = 0xff000000
)

func (e CssmAttributeDataNone) String() string {
	switch e {
	case CSSM_ATTRIBUTE_DATA_ACCESS_CREDENTIALS:
		return "CSSM_ATTRIBUTE_DATA_ACCESS_CREDENTIALS"
	case CSSM_ATTRIBUTE_DATA_CRYPTO_DATA:
		return "CSSM_ATTRIBUTE_DATA_CRYPTO_DATA"
	case CSSM_ATTRIBUTE_DATA_CSSM_DATA:
		return "CSSM_ATTRIBUTE_DATA_CSSM_DATA"
	case CSSM_ATTRIBUTE_DATA_DATE:
		return "CSSM_ATTRIBUTE_DATA_DATE"
	case CSSM_ATTRIBUTE_DATA_DL_DB_HANDLE:
		return "CSSM_ATTRIBUTE_DATA_DL_DB_HANDLE"
	case CSSM_ATTRIBUTE_DATA_KEY:
		return "CSSM_ATTRIBUTE_DATA_KEY"
	case CSSM_ATTRIBUTE_DATA_KR_PROFILE:
		return "CSSM_ATTRIBUTE_DATA_KR_PROFILE"
	case CSSM_ATTRIBUTE_DATA_NONE:
		return "CSSM_ATTRIBUTE_DATA_NONE"
	case CSSM_ATTRIBUTE_DATA_RANGE:
		return "CSSM_ATTRIBUTE_DATA_RANGE"
	case CSSM_ATTRIBUTE_DATA_STRING:
		return "CSSM_ATTRIBUTE_DATA_STRING"
	case CSSM_ATTRIBUTE_DATA_UINT32:
		return "CSSM_ATTRIBUTE_DATA_UINT32"
	case CSSM_ATTRIBUTE_DATA_VERSION:
		return "CSSM_ATTRIBUTE_DATA_VERSION"
	case CSSM_ATTRIBUTE_TYPE_MASK:
		return "CSSM_ATTRIBUTE_TYPE_MASK"
	default:
		return fmt.Sprintf("CssmAttributeDataNone(%d)", e)
	}
}

type CssmAttributeNone uint32

const (
	CSSM_ATTRIBUTE_ACCESS_CREDENTIALS   CssmAttributeNone = 2147483684
	CSSM_ATTRIBUTE_ALG_ID               CssmAttributeNone = 268435485
	CSSM_ATTRIBUTE_ALG_PARAMS           CssmAttributeNone = 536870928
	CSSM_ATTRIBUTE_BASE                 CssmAttributeNone = 536870939
	CSSM_ATTRIBUTE_BLOCK_SIZE           CssmAttributeNone = 268435468
	CSSM_ATTRIBUTE_CSP_HANDLE           CssmAttributeNone = 268435490
	CSSM_ATTRIBUTE_CUSTOM               CssmAttributeNone = 536870913
	CSSM_ATTRIBUTE_DESCRIPTION          CssmAttributeNone = 1342177282
	CSSM_ATTRIBUTE_DL_DB_HANDLE         CssmAttributeNone = 33554467
	CSSM_ATTRIBUTE_EFFECTIVE_BITS       CssmAttributeNone = 268435476
	CSSM_ATTRIBUTE_END_DATE             CssmAttributeNone = 1610612758
	CSSM_ATTRIBUTE_INIT_VECTOR          CssmAttributeNone = 536870916
	CSSM_ATTRIBUTE_ITERATION_COUNT      CssmAttributeNone = 268435486
	CSSM_ATTRIBUTE_IV_SIZE              CssmAttributeNone = 268435471
	CSSM_ATTRIBUTE_KEY                  CssmAttributeNone = 1073741827
	CSSM_ATTRIBUTE_KEYATTR              CssmAttributeNone = 268435480
	CSSM_ATTRIBUTE_KEYUSAGE             CssmAttributeNone = 268435479
	CSSM_ATTRIBUTE_KEY_LENGTH           CssmAttributeNone = 268435466
	CSSM_ATTRIBUTE_KEY_LENGTH_RANGE     CssmAttributeNone = 1879048203
	CSSM_ATTRIBUTE_KEY_TYPE             CssmAttributeNone = 268435474
	CSSM_ATTRIBUTE_KRPROFILE_LOCAL      CssmAttributeNone = 50331680
	CSSM_ATTRIBUTE_KRPROFILE_REMOTE     CssmAttributeNone = 50331681
	CSSM_ATTRIBUTE_LABEL                CssmAttributeNone = 536870929
	CSSM_ATTRIBUTE_MODE                 CssmAttributeNone = 268435475
	CSSM_ATTRIBUTE_NONE                 CssmAttributeNone = 0
	CSSM_ATTRIBUTE_OUTPUT_SIZE          CssmAttributeNone = 268435469
	CSSM_ATTRIBUTE_PADDING              CssmAttributeNone = 268435462
	CSSM_ATTRIBUTE_PASSPHRASE           CssmAttributeNone = 805306377
	CSSM_ATTRIBUTE_PRIME                CssmAttributeNone = 536870938
	CSSM_ATTRIBUTE_PRIVATE_KEY_FORMAT   CssmAttributeNone = 268435494
	CSSM_ATTRIBUTE_PUBLIC_KEY_FORMAT    CssmAttributeNone = 268435493
	CSSM_ATTRIBUTE_RANDOM               CssmAttributeNone = 536870919
	CSSM_ATTRIBUTE_ROUNDS               CssmAttributeNone = 268435470
	CSSM_ATTRIBUTE_ROUNDS_RANGE         CssmAttributeNone = 1879048223
	CSSM_ATTRIBUTE_SALT                 CssmAttributeNone = 536870917
	CSSM_ATTRIBUTE_SEED                 CssmAttributeNone = 805306376
	CSSM_ATTRIBUTE_START_DATE           CssmAttributeNone = 1610612757
	CSSM_ATTRIBUTE_SUBPRIME             CssmAttributeNone = 536870940
	CSSM_ATTRIBUTE_SYMMETRIC_KEY_FORMAT CssmAttributeNone = 268435495
	CSSM_ATTRIBUTE_VERSION              CssmAttributeNone = 16777241
	CSSM_ATTRIBUTE_WRAPPED_KEY_FORMAT   CssmAttributeNone = 268435496
)

func (e CssmAttributeNone) String() string {
	switch e {
	case CSSM_ATTRIBUTE_ACCESS_CREDENTIALS:
		return "CSSM_ATTRIBUTE_ACCESS_CREDENTIALS"
	case CSSM_ATTRIBUTE_ALG_ID:
		return "CSSM_ATTRIBUTE_ALG_ID"
	case CSSM_ATTRIBUTE_ALG_PARAMS:
		return "CSSM_ATTRIBUTE_ALG_PARAMS"
	case CSSM_ATTRIBUTE_BASE:
		return "CSSM_ATTRIBUTE_BASE"
	case CSSM_ATTRIBUTE_BLOCK_SIZE:
		return "CSSM_ATTRIBUTE_BLOCK_SIZE"
	case CSSM_ATTRIBUTE_CSP_HANDLE:
		return "CSSM_ATTRIBUTE_CSP_HANDLE"
	case CSSM_ATTRIBUTE_CUSTOM:
		return "CSSM_ATTRIBUTE_CUSTOM"
	case CSSM_ATTRIBUTE_DESCRIPTION:
		return "CSSM_ATTRIBUTE_DESCRIPTION"
	case CSSM_ATTRIBUTE_DL_DB_HANDLE:
		return "CSSM_ATTRIBUTE_DL_DB_HANDLE"
	case CSSM_ATTRIBUTE_EFFECTIVE_BITS:
		return "CSSM_ATTRIBUTE_EFFECTIVE_BITS"
	case CSSM_ATTRIBUTE_END_DATE:
		return "CSSM_ATTRIBUTE_END_DATE"
	case CSSM_ATTRIBUTE_INIT_VECTOR:
		return "CSSM_ATTRIBUTE_INIT_VECTOR"
	case CSSM_ATTRIBUTE_ITERATION_COUNT:
		return "CSSM_ATTRIBUTE_ITERATION_COUNT"
	case CSSM_ATTRIBUTE_IV_SIZE:
		return "CSSM_ATTRIBUTE_IV_SIZE"
	case CSSM_ATTRIBUTE_KEY:
		return "CSSM_ATTRIBUTE_KEY"
	case CSSM_ATTRIBUTE_KEYATTR:
		return "CSSM_ATTRIBUTE_KEYATTR"
	case CSSM_ATTRIBUTE_KEYUSAGE:
		return "CSSM_ATTRIBUTE_KEYUSAGE"
	case CSSM_ATTRIBUTE_KEY_LENGTH:
		return "CSSM_ATTRIBUTE_KEY_LENGTH"
	case CSSM_ATTRIBUTE_KEY_LENGTH_RANGE:
		return "CSSM_ATTRIBUTE_KEY_LENGTH_RANGE"
	case CSSM_ATTRIBUTE_KEY_TYPE:
		return "CSSM_ATTRIBUTE_KEY_TYPE"
	case CSSM_ATTRIBUTE_KRPROFILE_LOCAL:
		return "CSSM_ATTRIBUTE_KRPROFILE_LOCAL"
	case CSSM_ATTRIBUTE_KRPROFILE_REMOTE:
		return "CSSM_ATTRIBUTE_KRPROFILE_REMOTE"
	case CSSM_ATTRIBUTE_LABEL:
		return "CSSM_ATTRIBUTE_LABEL"
	case CSSM_ATTRIBUTE_MODE:
		return "CSSM_ATTRIBUTE_MODE"
	case CSSM_ATTRIBUTE_NONE:
		return "CSSM_ATTRIBUTE_NONE"
	case CSSM_ATTRIBUTE_OUTPUT_SIZE:
		return "CSSM_ATTRIBUTE_OUTPUT_SIZE"
	case CSSM_ATTRIBUTE_PADDING:
		return "CSSM_ATTRIBUTE_PADDING"
	case CSSM_ATTRIBUTE_PASSPHRASE:
		return "CSSM_ATTRIBUTE_PASSPHRASE"
	case CSSM_ATTRIBUTE_PRIME:
		return "CSSM_ATTRIBUTE_PRIME"
	case CSSM_ATTRIBUTE_PRIVATE_KEY_FORMAT:
		return "CSSM_ATTRIBUTE_PRIVATE_KEY_FORMAT"
	case CSSM_ATTRIBUTE_PUBLIC_KEY_FORMAT:
		return "CSSM_ATTRIBUTE_PUBLIC_KEY_FORMAT"
	case CSSM_ATTRIBUTE_RANDOM:
		return "CSSM_ATTRIBUTE_RANDOM"
	case CSSM_ATTRIBUTE_ROUNDS:
		return "CSSM_ATTRIBUTE_ROUNDS"
	case CSSM_ATTRIBUTE_ROUNDS_RANGE:
		return "CSSM_ATTRIBUTE_ROUNDS_RANGE"
	case CSSM_ATTRIBUTE_SALT:
		return "CSSM_ATTRIBUTE_SALT"
	case CSSM_ATTRIBUTE_SEED:
		return "CSSM_ATTRIBUTE_SEED"
	case CSSM_ATTRIBUTE_START_DATE:
		return "CSSM_ATTRIBUTE_START_DATE"
	case CSSM_ATTRIBUTE_SUBPRIME:
		return "CSSM_ATTRIBUTE_SUBPRIME"
	case CSSM_ATTRIBUTE_SYMMETRIC_KEY_FORMAT:
		return "CSSM_ATTRIBUTE_SYMMETRIC_KEY_FORMAT"
	case CSSM_ATTRIBUTE_VERSION:
		return "CSSM_ATTRIBUTE_VERSION"
	case CSSM_ATTRIBUTE_WRAPPED_KEY_FORMAT:
		return "CSSM_ATTRIBUTE_WRAPPED_KEY_FORMAT"
	default:
		return fmt.Sprintf("CssmAttributeNone(%d)", e)
	}
}

type CssmAttributePublicKey uint32

const (
	CSSM_ATTRIBUTE_ALERT_TITLE       CssmAttributePublicKey = 545259527
	CSSM_ATTRIBUTE_ASC_OPTIMIZATION  CssmAttributePublicKey = 276824067
	CSSM_ATTRIBUTE_FEE_CURVE_TYPE    CssmAttributePublicKey = 276824066
	CSSM_ATTRIBUTE_FEE_PRIME_TYPE    CssmAttributePublicKey = 276824065
	CSSM_ATTRIBUTE_PARAM_KEY         CssmAttributePublicKey = 1082130437
	CSSM_ATTRIBUTE_PROMPT            CssmAttributePublicKey = 545259526
	CSSM_ATTRIBUTE_PUBLIC_KEY        CssmAttributePublicKey = 1082130432
	CSSM_ATTRIBUTE_RSA_BLINDING      CssmAttributePublicKey = 276824068
	CSSM_ATTRIBUTE_VERIFY_PASSPHRASE CssmAttributePublicKey = 276824072
)

func (e CssmAttributePublicKey) String() string {
	switch e {
	case CSSM_ATTRIBUTE_ALERT_TITLE:
		return "CSSM_ATTRIBUTE_ALERT_TITLE"
	case CSSM_ATTRIBUTE_ASC_OPTIMIZATION:
		return "CSSM_ATTRIBUTE_ASC_OPTIMIZATION"
	case CSSM_ATTRIBUTE_FEE_CURVE_TYPE:
		return "CSSM_ATTRIBUTE_FEE_CURVE_TYPE"
	case CSSM_ATTRIBUTE_FEE_PRIME_TYPE:
		return "CSSM_ATTRIBUTE_FEE_PRIME_TYPE"
	case CSSM_ATTRIBUTE_PARAM_KEY:
		return "CSSM_ATTRIBUTE_PARAM_KEY"
	case CSSM_ATTRIBUTE_PROMPT:
		return "CSSM_ATTRIBUTE_PROMPT"
	case CSSM_ATTRIBUTE_PUBLIC_KEY:
		return "CSSM_ATTRIBUTE_PUBLIC_KEY"
	case CSSM_ATTRIBUTE_RSA_BLINDING:
		return "CSSM_ATTRIBUTE_RSA_BLINDING"
	case CSSM_ATTRIBUTE_VERIFY_PASSPHRASE:
		return "CSSM_ATTRIBUTE_VERIFY_PASSPHRASE"
	default:
		return fmt.Sprintf("CssmAttributePublicKey(%d)", e)
	}
}

type CssmAttributeVendor uint32

const (
	CSSM_ATTRIBUTE_VENDOR_DEFINED CssmAttributeVendor = 0x800000
)

func (e CssmAttributeVendor) String() string {
	switch e {
	case CSSM_ATTRIBUTE_VENDOR_DEFINED:
		return "CSSM_ATTRIBUTE_VENDOR_DEFINED"
	default:
		return fmt.Sprintf("CssmAttributeVendor(%d)", e)
	}
}

type CssmBaseError int32

const (
	CSSM_AC_BASE_ERROR           CssmBaseError = -2147405824
	CSSM_AC_PRIVATE_ERROR        CssmBaseError = -2147404800
	CSSM_BASE_ERROR              CssmBaseError = -0x7fff0000
	CSSM_CL_BASE_ERROR           CssmBaseError = -2147411968
	CSSM_CL_PRIVATE_ERROR        CssmBaseError = -2147410944
	CSSM_CSP_BASE_ERROR          CssmBaseError = -2147416064
	CSSM_CSP_PRIVATE_ERROR       CssmBaseError = -2147415040
	CSSM_CSSM_BASE_ERROR         CssmBaseError = -2147418112
	CSSM_CSSM_PRIVATE_ERROR      CssmBaseError = -2147417088
	CSSM_DL_BASE_ERROR           CssmBaseError = -2147414016
	CSSM_DL_PRIVATE_ERROR        CssmBaseError = -2147412992
	CSSM_ERRORCODE_COMMON_EXTENT CssmBaseError = 0x100
	CSSM_ERRORCODE_CUSTOM_OFFSET CssmBaseError = 0x400
	CSSM_ERRORCODE_MODULE_EXTENT CssmBaseError = 0x800
	CSSM_KR_BASE_ERROR           CssmBaseError = -2147407872
	CSSM_KR_PRIVATE_ERROR        CssmBaseError = -2147406848
	CSSM_TP_BASE_ERROR           CssmBaseError = -2147409920
	CSSM_TP_PRIVATE_ERROR        CssmBaseError = -2147408896
)

func (e CssmBaseError) String() string {
	switch e {
	case CSSM_AC_BASE_ERROR:
		return "CSSM_AC_BASE_ERROR"
	case CSSM_AC_PRIVATE_ERROR:
		return "CSSM_AC_PRIVATE_ERROR"
	case CSSM_BASE_ERROR:
		return "CSSM_BASE_ERROR"
	case CSSM_CL_BASE_ERROR:
		return "CSSM_CL_BASE_ERROR"
	case CSSM_CL_PRIVATE_ERROR:
		return "CSSM_CL_PRIVATE_ERROR"
	case CSSM_CSP_BASE_ERROR:
		return "CSSM_CSP_BASE_ERROR"
	case CSSM_CSP_PRIVATE_ERROR:
		return "CSSM_CSP_PRIVATE_ERROR"
	case CSSM_CSSM_PRIVATE_ERROR:
		return "CSSM_CSSM_PRIVATE_ERROR"
	case CSSM_DL_BASE_ERROR:
		return "CSSM_DL_BASE_ERROR"
	case CSSM_DL_PRIVATE_ERROR:
		return "CSSM_DL_PRIVATE_ERROR"
	case CSSM_ERRORCODE_COMMON_EXTENT:
		return "CSSM_ERRORCODE_COMMON_EXTENT"
	case CSSM_ERRORCODE_CUSTOM_OFFSET:
		return "CSSM_ERRORCODE_CUSTOM_OFFSET"
	case CSSM_ERRORCODE_MODULE_EXTENT:
		return "CSSM_ERRORCODE_MODULE_EXTENT"
	case CSSM_KR_BASE_ERROR:
		return "CSSM_KR_BASE_ERROR"
	case CSSM_KR_PRIVATE_ERROR:
		return "CSSM_KR_PRIVATE_ERROR"
	case CSSM_TP_BASE_ERROR:
		return "CSSM_TP_BASE_ERROR"
	case CSSM_TP_PRIVATE_ERROR:
		return "CSSM_TP_PRIVATE_ERROR"
	default:
		return fmt.Sprintf("CssmBaseError(%d)", e)
	}
}

type CssmCLBaseCLError int32

const (
	CSSMERR_CL_INVALID_BUNDLE_INFO    CssmCLBaseCLError = -2147411708
	CSSMERR_CL_INVALID_BUNDLE_POINTER CssmCLBaseCLError = -2147411711
	CSSMERR_CL_INVALID_CACHE_HANDLE   CssmCLBaseCLError = -2147411710
	CSSMERR_CL_INVALID_CRL_INDEX      CssmCLBaseCLError = -2147411707
	CSSMERR_CL_INVALID_RESULTS_HANDLE CssmCLBaseCLError = -2147411709
	CSSMERR_CL_INVALID_SCOPE          CssmCLBaseCLError = -2147411706
	CSSMERR_CL_NO_FIELD_VALUES        CssmCLBaseCLError = -2147411705
	CSSMERR_CL_SCOPE_NOT_SUPPORTED    CssmCLBaseCLError = -2147411704
	CSSM_CL_BASE_CL_ERROR             CssmCLBaseCLError = -2147411712
)

func (e CssmCLBaseCLError) String() string {
	switch e {
	case CSSMERR_CL_INVALID_BUNDLE_INFO:
		return "CSSMERR_CL_INVALID_BUNDLE_INFO"
	case CSSMERR_CL_INVALID_BUNDLE_POINTER:
		return "CSSMERR_CL_INVALID_BUNDLE_POINTER"
	case CSSMERR_CL_INVALID_CACHE_HANDLE:
		return "CSSMERR_CL_INVALID_CACHE_HANDLE"
	case CSSMERR_CL_INVALID_CRL_INDEX:
		return "CSSMERR_CL_INVALID_CRL_INDEX"
	case CSSMERR_CL_INVALID_RESULTS_HANDLE:
		return "CSSMERR_CL_INVALID_RESULTS_HANDLE"
	case CSSMERR_CL_INVALID_SCOPE:
		return "CSSMERR_CL_INVALID_SCOPE"
	case CSSMERR_CL_NO_FIELD_VALUES:
		return "CSSMERR_CL_NO_FIELD_VALUES"
	case CSSMERR_CL_SCOPE_NOT_SUPPORTED:
		return "CSSMERR_CL_SCOPE_NOT_SUPPORTED"
	case CSSM_CL_BASE_CL_ERROR:
		return "CSSM_CL_BASE_CL_ERROR"
	default:
		return fmt.Sprintf("CssmCLBaseCLError(%d)", e)
	}
}

type CssmCLTemplate uint32

const (
	CSSM_CL_TEMPLATE_INTERMEDIATE_CERT CssmCLTemplate = 1
	CSSM_CL_TEMPLATE_PKIX_CERTTEMPLATE CssmCLTemplate = 2
)

func (e CssmCLTemplate) String() string {
	switch e {
	case CSSM_CL_TEMPLATE_INTERMEDIATE_CERT:
		return "CSSM_CL_TEMPLATE_INTERMEDIATE_CERT"
	case CSSM_CL_TEMPLATE_PKIX_CERTTEMPLATE:
		return "CSSM_CL_TEMPLATE_PKIX_CERTTEMPLATE"
	default:
		return fmt.Sprintf("CssmCLTemplate(%d)", e)
	}
}

type CssmCertBundleEncoding uint32

const (
	CSSM_CERT_BUNDLE_ENCODING_BER     CssmCertBundleEncoding = 0x2
	CSSM_CERT_BUNDLE_ENCODING_CUSTOM  CssmCertBundleEncoding = 0x1
	CSSM_CERT_BUNDLE_ENCODING_DER     CssmCertBundleEncoding = 0x3
	CSSM_CERT_BUNDLE_ENCODING_PGP     CssmCertBundleEncoding = 0x5
	CSSM_CERT_BUNDLE_ENCODING_SEXPR   CssmCertBundleEncoding = 0x4
	CSSM_CERT_BUNDLE_ENCODING_UNKNOWN CssmCertBundleEncoding = 0
)

func (e CssmCertBundleEncoding) String() string {
	switch e {
	case CSSM_CERT_BUNDLE_ENCODING_BER:
		return "CSSM_CERT_BUNDLE_ENCODING_BER"
	case CSSM_CERT_BUNDLE_ENCODING_CUSTOM:
		return "CSSM_CERT_BUNDLE_ENCODING_CUSTOM"
	case CSSM_CERT_BUNDLE_ENCODING_DER:
		return "CSSM_CERT_BUNDLE_ENCODING_DER"
	case CSSM_CERT_BUNDLE_ENCODING_PGP:
		return "CSSM_CERT_BUNDLE_ENCODING_PGP"
	case CSSM_CERT_BUNDLE_ENCODING_SEXPR:
		return "CSSM_CERT_BUNDLE_ENCODING_SEXPR"
	case CSSM_CERT_BUNDLE_ENCODING_UNKNOWN:
		return "CSSM_CERT_BUNDLE_ENCODING_UNKNOWN"
	default:
		return fmt.Sprintf("CssmCertBundleEncoding(%d)", e)
	}
}

type CssmCertBundleUnknown uint32

const (
	CSSM_CERT_BUNDLE_CUSTOM                      CssmCertBundleUnknown = 0x1
	CSSM_CERT_BUNDLE_LAST                        CssmCertBundleUnknown = 0x7fff
	CSSM_CERT_BUNDLE_PFX                         CssmCertBundleUnknown = 0x5
	CSSM_CERT_BUNDLE_PGP_KEYRING                 CssmCertBundleUnknown = 0x7
	CSSM_CERT_BUNDLE_PKCS12                      CssmCertBundleUnknown = 0x4
	CSSM_CERT_BUNDLE_PKCS7_SIGNED_DATA           CssmCertBundleUnknown = 0x2
	CSSM_CERT_BUNDLE_PKCS7_SIGNED_ENVELOPED_DATA CssmCertBundleUnknown = 0x3
	CSSM_CERT_BUNDLE_SPKI_SEQUENCE               CssmCertBundleUnknown = 0x6
	CSSM_CERT_BUNDLE_UNKNOWN                     CssmCertBundleUnknown = 0
	CSSM_CL_CUSTOM_CERT_BUNDLE_TYPE              CssmCertBundleUnknown = 0x8000
)

func (e CssmCertBundleUnknown) String() string {
	switch e {
	case CSSM_CERT_BUNDLE_CUSTOM:
		return "CSSM_CERT_BUNDLE_CUSTOM"
	case CSSM_CERT_BUNDLE_LAST:
		return "CSSM_CERT_BUNDLE_LAST"
	case CSSM_CERT_BUNDLE_PFX:
		return "CSSM_CERT_BUNDLE_PFX"
	case CSSM_CERT_BUNDLE_PGP_KEYRING:
		return "CSSM_CERT_BUNDLE_PGP_KEYRING"
	case CSSM_CERT_BUNDLE_PKCS12:
		return "CSSM_CERT_BUNDLE_PKCS12"
	case CSSM_CERT_BUNDLE_PKCS7_SIGNED_DATA:
		return "CSSM_CERT_BUNDLE_PKCS7_SIGNED_DATA"
	case CSSM_CERT_BUNDLE_PKCS7_SIGNED_ENVELOPED_DATA:
		return "CSSM_CERT_BUNDLE_PKCS7_SIGNED_ENVELOPED_DATA"
	case CSSM_CERT_BUNDLE_SPKI_SEQUENCE:
		return "CSSM_CERT_BUNDLE_SPKI_SEQUENCE"
	case CSSM_CERT_BUNDLE_UNKNOWN:
		return "CSSM_CERT_BUNDLE_UNKNOWN"
	case CSSM_CL_CUSTOM_CERT_BUNDLE_TYPE:
		return "CSSM_CL_CUSTOM_CERT_BUNDLE_TYPE"
	default:
		return fmt.Sprintf("CssmCertBundleUnknown(%d)", e)
	}
}

type CssmCertEncodingUnknown uint32

const (
	CSSM_CERT_ENCODING_BER       CssmCertEncodingUnknown = 0x2
	CSSM_CERT_ENCODING_CUSTOM    CssmCertEncodingUnknown = 0x1
	CSSM_CERT_ENCODING_DER       CssmCertEncodingUnknown = 0x3
	CSSM_CERT_ENCODING_LAST      CssmCertEncodingUnknown = 0x7fff
	CSSM_CERT_ENCODING_MULTIPLE  CssmCertEncodingUnknown = 0x7ffe
	CSSM_CERT_ENCODING_NDR       CssmCertEncodingUnknown = 0x4
	CSSM_CERT_ENCODING_PGP       CssmCertEncodingUnknown = 0x6
	CSSM_CERT_ENCODING_SEXPR     CssmCertEncodingUnknown = 0x5
	CSSM_CERT_ENCODING_UNKNOWN   CssmCertEncodingUnknown = 0
	CSSM_CL_CUSTOM_CERT_ENCODING CssmCertEncodingUnknown = 0x8000
)

func (e CssmCertEncodingUnknown) String() string {
	switch e {
	case CSSM_CERT_ENCODING_BER:
		return "CSSM_CERT_ENCODING_BER"
	case CSSM_CERT_ENCODING_CUSTOM:
		return "CSSM_CERT_ENCODING_CUSTOM"
	case CSSM_CERT_ENCODING_DER:
		return "CSSM_CERT_ENCODING_DER"
	case CSSM_CERT_ENCODING_LAST:
		return "CSSM_CERT_ENCODING_LAST"
	case CSSM_CERT_ENCODING_MULTIPLE:
		return "CSSM_CERT_ENCODING_MULTIPLE"
	case CSSM_CERT_ENCODING_NDR:
		return "CSSM_CERT_ENCODING_NDR"
	case CSSM_CERT_ENCODING_PGP:
		return "CSSM_CERT_ENCODING_PGP"
	case CSSM_CERT_ENCODING_SEXPR:
		return "CSSM_CERT_ENCODING_SEXPR"
	case CSSM_CERT_ENCODING_UNKNOWN:
		return "CSSM_CERT_ENCODING_UNKNOWN"
	case CSSM_CL_CUSTOM_CERT_ENCODING:
		return "CSSM_CL_CUSTOM_CERT_ENCODING"
	default:
		return fmt.Sprintf("CssmCertEncodingUnknown(%d)", e)
	}
}

type CssmCertParseFormatNone uint32

const (
	CSSM_CERT_PARSE_FORMAT_COMPLEX   CssmCertParseFormatNone = 0x3
	CSSM_CERT_PARSE_FORMAT_CUSTOM    CssmCertParseFormatNone = 0x1
	CSSM_CERT_PARSE_FORMAT_LAST      CssmCertParseFormatNone = 0x7fff
	CSSM_CERT_PARSE_FORMAT_MULTIPLE  CssmCertParseFormatNone = 0x7ffe
	CSSM_CERT_PARSE_FORMAT_NONE      CssmCertParseFormatNone = 0
	CSSM_CERT_PARSE_FORMAT_OID_NAMED CssmCertParseFormatNone = 0x4
	CSSM_CERT_PARSE_FORMAT_SEXPR     CssmCertParseFormatNone = 0x2
	CSSM_CERT_PARSE_FORMAT_TUPLE     CssmCertParseFormatNone = 0x5
	CSSM_CL_CUSTOM_CERT_PARSE_FORMAT CssmCertParseFormatNone = 0x8000
)

func (e CssmCertParseFormatNone) String() string {
	switch e {
	case CSSM_CERT_PARSE_FORMAT_COMPLEX:
		return "CSSM_CERT_PARSE_FORMAT_COMPLEX"
	case CSSM_CERT_PARSE_FORMAT_CUSTOM:
		return "CSSM_CERT_PARSE_FORMAT_CUSTOM"
	case CSSM_CERT_PARSE_FORMAT_LAST:
		return "CSSM_CERT_PARSE_FORMAT_LAST"
	case CSSM_CERT_PARSE_FORMAT_MULTIPLE:
		return "CSSM_CERT_PARSE_FORMAT_MULTIPLE"
	case CSSM_CERT_PARSE_FORMAT_NONE:
		return "CSSM_CERT_PARSE_FORMAT_NONE"
	case CSSM_CERT_PARSE_FORMAT_OID_NAMED:
		return "CSSM_CERT_PARSE_FORMAT_OID_NAMED"
	case CSSM_CERT_PARSE_FORMAT_SEXPR:
		return "CSSM_CERT_PARSE_FORMAT_SEXPR"
	case CSSM_CERT_PARSE_FORMAT_TUPLE:
		return "CSSM_CERT_PARSE_FORMAT_TUPLE"
	case CSSM_CL_CUSTOM_CERT_PARSE_FORMAT:
		return "CSSM_CL_CUSTOM_CERT_PARSE_FORMAT"
	default:
		return fmt.Sprintf("CssmCertParseFormatNone(%d)", e)
	}
}

type CssmCertStatus uint32

const (
	CSSM_CERT_STATUS_EXPIRED                      CssmCertStatus = 0x1
	CSSM_CERT_STATUS_IS_FROM_NET                  CssmCertStatus = 0x20
	CSSM_CERT_STATUS_IS_IN_ANCHORS                CssmCertStatus = 0x8
	CSSM_CERT_STATUS_IS_IN_INPUT_CERTS            CssmCertStatus = 0x4
	CSSM_CERT_STATUS_IS_ROOT                      CssmCertStatus = 0x10
	CSSM_CERT_STATUS_NOT_VALID_YET                CssmCertStatus = 0x2
	CSSM_CERT_STATUS_TRUST_SETTINGS_DENY          CssmCertStatus = 0x400
	CSSM_CERT_STATUS_TRUST_SETTINGS_FOUND_ADMIN   CssmCertStatus = 0x80
	CSSM_CERT_STATUS_TRUST_SETTINGS_FOUND_SYSTEM  CssmCertStatus = 0x100
	CSSM_CERT_STATUS_TRUST_SETTINGS_FOUND_USER    CssmCertStatus = 0x40
	CSSM_CERT_STATUS_TRUST_SETTINGS_IGNORED_ERROR CssmCertStatus = 0x800
	CSSM_CERT_STATUS_TRUST_SETTINGS_TRUST         CssmCertStatus = 0x200
)

func (e CssmCertStatus) String() string {
	switch e {
	case CSSM_CERT_STATUS_EXPIRED:
		return "CSSM_CERT_STATUS_EXPIRED"
	case CSSM_CERT_STATUS_IS_FROM_NET:
		return "CSSM_CERT_STATUS_IS_FROM_NET"
	case CSSM_CERT_STATUS_IS_IN_ANCHORS:
		return "CSSM_CERT_STATUS_IS_IN_ANCHORS"
	case CSSM_CERT_STATUS_IS_IN_INPUT_CERTS:
		return "CSSM_CERT_STATUS_IS_IN_INPUT_CERTS"
	case CSSM_CERT_STATUS_IS_ROOT:
		return "CSSM_CERT_STATUS_IS_ROOT"
	case CSSM_CERT_STATUS_NOT_VALID_YET:
		return "CSSM_CERT_STATUS_NOT_VALID_YET"
	case CSSM_CERT_STATUS_TRUST_SETTINGS_DENY:
		return "CSSM_CERT_STATUS_TRUST_SETTINGS_DENY"
	case CSSM_CERT_STATUS_TRUST_SETTINGS_FOUND_ADMIN:
		return "CSSM_CERT_STATUS_TRUST_SETTINGS_FOUND_ADMIN"
	case CSSM_CERT_STATUS_TRUST_SETTINGS_FOUND_SYSTEM:
		return "CSSM_CERT_STATUS_TRUST_SETTINGS_FOUND_SYSTEM"
	case CSSM_CERT_STATUS_TRUST_SETTINGS_FOUND_USER:
		return "CSSM_CERT_STATUS_TRUST_SETTINGS_FOUND_USER"
	case CSSM_CERT_STATUS_TRUST_SETTINGS_IGNORED_ERROR:
		return "CSSM_CERT_STATUS_TRUST_SETTINGS_IGNORED_ERROR"
	case CSSM_CERT_STATUS_TRUST_SETTINGS_TRUST:
		return "CSSM_CERT_STATUS_TRUST_SETTINGS_TRUST"
	default:
		return fmt.Sprintf("CssmCertStatus(%d)", e)
	}
}

type CssmCertUnknown uint32

const (
	CSSM_CERT_ACL_ENTRY       CssmCertUnknown = 0xc
	CSSM_CERT_Intel           CssmCertUnknown = 0x8
	CSSM_CERT_LAST            CssmCertUnknown = 0x7fff
	CSSM_CERT_MULTIPLE        CssmCertUnknown = 0x7ffe
	CSSM_CERT_PGP             CssmCertUnknown = 0x4
	CSSM_CERT_SDSIv1          CssmCertUnknown = 0x6
	CSSM_CERT_SPKI            CssmCertUnknown = 0x5
	CSSM_CERT_TUPLE           CssmCertUnknown = 0xb
	CSSM_CERT_UNKNOWN         CssmCertUnknown = 0
	CSSM_CERT_X9_ATTRIBUTE    CssmCertUnknown = 0xa
	CSSM_CERT_X_509_ATTRIBUTE CssmCertUnknown = 0x9
	CSSM_CERT_X_509v1         CssmCertUnknown = 0x1
	CSSM_CERT_X_509v2         CssmCertUnknown = 0x2
	CSSM_CERT_X_509v3         CssmCertUnknown = 0x3
	CSSM_CL_CUSTOM_CERT_TYPE  CssmCertUnknown = 0x8000
)

func (e CssmCertUnknown) String() string {
	switch e {
	case CSSM_CERT_ACL_ENTRY:
		return "CSSM_CERT_ACL_ENTRY"
	case CSSM_CERT_Intel:
		return "CSSM_CERT_Intel"
	case CSSM_CERT_LAST:
		return "CSSM_CERT_LAST"
	case CSSM_CERT_MULTIPLE:
		return "CSSM_CERT_MULTIPLE"
	case CSSM_CERT_PGP:
		return "CSSM_CERT_PGP"
	case CSSM_CERT_SDSIv1:
		return "CSSM_CERT_SDSIv1"
	case CSSM_CERT_SPKI:
		return "CSSM_CERT_SPKI"
	case CSSM_CERT_TUPLE:
		return "CSSM_CERT_TUPLE"
	case CSSM_CERT_UNKNOWN:
		return "CSSM_CERT_UNKNOWN"
	case CSSM_CERT_X9_ATTRIBUTE:
		return "CSSM_CERT_X9_ATTRIBUTE"
	case CSSM_CERT_X_509_ATTRIBUTE:
		return "CSSM_CERT_X_509_ATTRIBUTE"
	case CSSM_CERT_X_509v1:
		return "CSSM_CERT_X_509v1"
	case CSSM_CERT_X_509v2:
		return "CSSM_CERT_X_509v2"
	case CSSM_CERT_X_509v3:
		return "CSSM_CERT_X_509v3"
	case CSSM_CL_CUSTOM_CERT_TYPE:
		return "CSSM_CL_CUSTOM_CERT_TYPE"
	default:
		return fmt.Sprintf("CssmCertUnknown(%d)", e)
	}
}

type CssmCertgroup uint32

const (
	CSSM_CERTGROUP_CERT_PAIR    CssmCertgroup = 0x3
	CSSM_CERTGROUP_DATA         CssmCertgroup = 0
	CSSM_CERTGROUP_ENCODED_CERT CssmCertgroup = 0x1
	CSSM_CERTGROUP_PARSED_CERT  CssmCertgroup = 0x2
)

func (e CssmCertgroup) String() string {
	switch e {
	case CSSM_CERTGROUP_CERT_PAIR:
		return "CSSM_CERTGROUP_CERT_PAIR"
	case CSSM_CERTGROUP_DATA:
		return "CSSM_CERTGROUP_DATA"
	case CSSM_CERTGROUP_ENCODED_CERT:
		return "CSSM_CERTGROUP_ENCODED_CERT"
	case CSSM_CERTGROUP_PARSED_CERT:
		return "CSSM_CERTGROUP_PARSED_CERT"
	default:
		return fmt.Sprintf("CssmCertgroup(%d)", e)
	}
}

type CssmContextEvent uint32

const (
	CSSM_CONTEXT_EVENT_CREATE CssmContextEvent = 1
	CSSM_CONTEXT_EVENT_DELETE CssmContextEvent = 2
	CSSM_CONTEXT_EVENT_UPDATE CssmContextEvent = 3
)

func (e CssmContextEvent) String() string {
	switch e {
	case CSSM_CONTEXT_EVENT_CREATE:
		return "CSSM_CONTEXT_EVENT_CREATE"
	case CSSM_CONTEXT_EVENT_DELETE:
		return "CSSM_CONTEXT_EVENT_DELETE"
	case CSSM_CONTEXT_EVENT_UPDATE:
		return "CSSM_CONTEXT_EVENT_UPDATE"
	default:
		return fmt.Sprintf("CssmContextEvent(%d)", e)
	}
}

type CssmCrlEncoding uint32

const (
	CSSM_CRL_ENCODING_BER      CssmCrlEncoding = 0x2
	CSSM_CRL_ENCODING_BLOOM    CssmCrlEncoding = 0x4
	CSSM_CRL_ENCODING_CUSTOM   CssmCrlEncoding = 0x1
	CSSM_CRL_ENCODING_DER      CssmCrlEncoding = 0x3
	CSSM_CRL_ENCODING_MULTIPLE CssmCrlEncoding = 0x7ffe
	CSSM_CRL_ENCODING_SEXPR    CssmCrlEncoding = 0x5
	CSSM_CRL_ENCODING_UNKNOWN  CssmCrlEncoding = 0
)

func (e CssmCrlEncoding) String() string {
	switch e {
	case CSSM_CRL_ENCODING_BER:
		return "CSSM_CRL_ENCODING_BER"
	case CSSM_CRL_ENCODING_BLOOM:
		return "CSSM_CRL_ENCODING_BLOOM"
	case CSSM_CRL_ENCODING_CUSTOM:
		return "CSSM_CRL_ENCODING_CUSTOM"
	case CSSM_CRL_ENCODING_DER:
		return "CSSM_CRL_ENCODING_DER"
	case CSSM_CRL_ENCODING_MULTIPLE:
		return "CSSM_CRL_ENCODING_MULTIPLE"
	case CSSM_CRL_ENCODING_SEXPR:
		return "CSSM_CRL_ENCODING_SEXPR"
	case CSSM_CRL_ENCODING_UNKNOWN:
		return "CSSM_CRL_ENCODING_UNKNOWN"
	default:
		return fmt.Sprintf("CssmCrlEncoding(%d)", e)
	}
}

type CssmCrlParseFormatNone uint32

const (
	CSSM_CL_CUSTOM_CRL_PARSE_FORMAT CssmCrlParseFormatNone = 0x8000
	CSSM_CRL_PARSE_FORMAT_COMPLEX   CssmCrlParseFormatNone = 0x3
	CSSM_CRL_PARSE_FORMAT_CUSTOM    CssmCrlParseFormatNone = 0x1
	CSSM_CRL_PARSE_FORMAT_LAST      CssmCrlParseFormatNone = 0x7fff
	CSSM_CRL_PARSE_FORMAT_MULTIPLE  CssmCrlParseFormatNone = 0x7ffe
	CSSM_CRL_PARSE_FORMAT_NONE      CssmCrlParseFormatNone = 0
	CSSM_CRL_PARSE_FORMAT_OID_NAMED CssmCrlParseFormatNone = 0x4
	CSSM_CRL_PARSE_FORMAT_SEXPR     CssmCrlParseFormatNone = 0x2
	CSSM_CRL_PARSE_FORMAT_TUPLE     CssmCrlParseFormatNone = 0x5
)

func (e CssmCrlParseFormatNone) String() string {
	switch e {
	case CSSM_CL_CUSTOM_CRL_PARSE_FORMAT:
		return "CSSM_CL_CUSTOM_CRL_PARSE_FORMAT"
	case CSSM_CRL_PARSE_FORMAT_COMPLEX:
		return "CSSM_CRL_PARSE_FORMAT_COMPLEX"
	case CSSM_CRL_PARSE_FORMAT_CUSTOM:
		return "CSSM_CRL_PARSE_FORMAT_CUSTOM"
	case CSSM_CRL_PARSE_FORMAT_LAST:
		return "CSSM_CRL_PARSE_FORMAT_LAST"
	case CSSM_CRL_PARSE_FORMAT_MULTIPLE:
		return "CSSM_CRL_PARSE_FORMAT_MULTIPLE"
	case CSSM_CRL_PARSE_FORMAT_NONE:
		return "CSSM_CRL_PARSE_FORMAT_NONE"
	case CSSM_CRL_PARSE_FORMAT_OID_NAMED:
		return "CSSM_CRL_PARSE_FORMAT_OID_NAMED"
	case CSSM_CRL_PARSE_FORMAT_SEXPR:
		return "CSSM_CRL_PARSE_FORMAT_SEXPR"
	case CSSM_CRL_PARSE_FORMAT_TUPLE:
		return "CSSM_CRL_PARSE_FORMAT_TUPLE"
	default:
		return fmt.Sprintf("CssmCrlParseFormatNone(%d)", e)
	}
}

type CssmCrlType uint32

const (
	CSSM_CRL_TYPE_MULTIPLE CssmCrlType = 0x7ffe
	CSSM_CRL_TYPE_SPKI     CssmCrlType = 0x3
	CSSM_CRL_TYPE_UNKNOWN  CssmCrlType = 0
	CSSM_CRL_TYPE_X_509v1  CssmCrlType = 0x1
	CSSM_CRL_TYPE_X_509v2  CssmCrlType = 0x2
)

func (e CssmCrlType) String() string {
	switch e {
	case CSSM_CRL_TYPE_MULTIPLE:
		return "CSSM_CRL_TYPE_MULTIPLE"
	case CSSM_CRL_TYPE_SPKI:
		return "CSSM_CRL_TYPE_SPKI"
	case CSSM_CRL_TYPE_UNKNOWN:
		return "CSSM_CRL_TYPE_UNKNOWN"
	case CSSM_CRL_TYPE_X_509v1:
		return "CSSM_CRL_TYPE_X_509v1"
	case CSSM_CRL_TYPE_X_509v2:
		return "CSSM_CRL_TYPE_X_509v2"
	default:
		return fmt.Sprintf("CssmCrlType(%d)", e)
	}
}

type CssmCrlgroup uint32

const (
	CSSM_CRLGROUP_CRL_PAIR    CssmCrlgroup = 0x3
	CSSM_CRLGROUP_DATA        CssmCrlgroup = 0
	CSSM_CRLGROUP_ENCODED_CRL CssmCrlgroup = 0x1
	CSSM_CRLGROUP_PARSED_CRL  CssmCrlgroup = 0x2
)

func (e CssmCrlgroup) String() string {
	switch e {
	case CSSM_CRLGROUP_CRL_PAIR:
		return "CSSM_CRLGROUP_CRL_PAIR"
	case CSSM_CRLGROUP_DATA:
		return "CSSM_CRLGROUP_DATA"
	case CSSM_CRLGROUP_ENCODED_CRL:
		return "CSSM_CRLGROUP_ENCODED_CRL"
	case CSSM_CRLGROUP_PARSED_CRL:
		return "CSSM_CRLGROUP_PARSED_CRL"
	default:
		return fmt.Sprintf("CssmCrlgroup(%d)", e)
	}
}

type CssmCspBaseCspError int32

const (
	CSSMERR_CSP_ALGID_MISMATCH                    CssmCspBaseCspError = -2147415789
	CSSMERR_CSP_ALREADY_LOGGED_IN                 CssmCspBaseCspError = -2147415726
	CSSMERR_CSP_ATTACH_HANDLE_BUSY                CssmCspBaseCspError = -2147415802
	CSSMERR_CSP_BLOCK_SIZE_MISMATCH               CssmCspBaseCspError = -2147415731
	CSSMERR_CSP_CRYPTO_DATA_CALLBACK_FAILED       CssmCspBaseCspError = -2147415722
	CSSMERR_CSP_DEVICE_ERROR                      CssmCspBaseCspError = -2147415804
	CSSMERR_CSP_DEVICE_MEMORY_ERROR               CssmCspBaseCspError = -2147415803
	CSSMERR_CSP_DEVICE_VERIFY_FAILED              CssmCspBaseCspError = -2147415728
	CSSMERR_CSP_INPUT_LENGTH_ERROR                CssmCspBaseCspError = -2147415807
	CSSMERR_CSP_INVALID_ALGORITHM                 CssmCspBaseCspError = -2147415759
	CSSMERR_CSP_INVALID_ATTR_ACCESS_CREDENTIALS   CssmCspBaseCspError = -2147415678
	CSSMERR_CSP_INVALID_ATTR_ALG_PARAMS           CssmCspBaseCspError = -2147415704
	CSSMERR_CSP_INVALID_ATTR_BASE                 CssmCspBaseCspError = -2147415686
	CSSMERR_CSP_INVALID_ATTR_BLOCK_SIZE           CssmCspBaseCspError = -2147415738
	CSSMERR_CSP_INVALID_ATTR_DL_DB_HANDLE         CssmCspBaseCspError = -2147415680
	CSSMERR_CSP_INVALID_ATTR_EFFECTIVE_BITS       CssmCspBaseCspError = -2147415696
	CSSMERR_CSP_INVALID_ATTR_END_DATE             CssmCspBaseCspError = -2147415692
	CSSMERR_CSP_INVALID_ATTR_INIT_VECTOR          CssmCspBaseCspError = -2147415752
	CSSMERR_CSP_INVALID_ATTR_ITERATION_COUNT      CssmCspBaseCspError = -2147415682
	CSSMERR_CSP_INVALID_ATTR_KEY                  CssmCspBaseCspError = -2147415754
	CSSMERR_CSP_INVALID_ATTR_KEY_LENGTH           CssmCspBaseCspError = -2147415740
	CSSMERR_CSP_INVALID_ATTR_KEY_TYPE             CssmCspBaseCspError = -2147415700
	CSSMERR_CSP_INVALID_ATTR_LABEL                CssmCspBaseCspError = -2147415702
	CSSMERR_CSP_INVALID_ATTR_MODE                 CssmCspBaseCspError = -2147415698
	CSSMERR_CSP_INVALID_ATTR_OUTPUT_SIZE          CssmCspBaseCspError = -2147415708
	CSSMERR_CSP_INVALID_ATTR_PADDING              CssmCspBaseCspError = -2147415748
	CSSMERR_CSP_INVALID_ATTR_PASSPHRASE           CssmCspBaseCspError = -2147415742
	CSSMERR_CSP_INVALID_ATTR_PRIME                CssmCspBaseCspError = -2147415688
	CSSMERR_CSP_INVALID_ATTR_PRIVATE_KEY_FORMAT   CssmCspBaseCspError = -2147415674
	CSSMERR_CSP_INVALID_ATTR_PUBLIC_KEY_FORMAT    CssmCspBaseCspError = -2147415676
	CSSMERR_CSP_INVALID_ATTR_RANDOM               CssmCspBaseCspError = -2147415746
	CSSMERR_CSP_INVALID_ATTR_ROUNDS               CssmCspBaseCspError = -2147415706
	CSSMERR_CSP_INVALID_ATTR_SALT                 CssmCspBaseCspError = -2147415750
	CSSMERR_CSP_INVALID_ATTR_SEED                 CssmCspBaseCspError = -2147415744
	CSSMERR_CSP_INVALID_ATTR_START_DATE           CssmCspBaseCspError = -2147415694
	CSSMERR_CSP_INVALID_ATTR_SUBPRIME             CssmCspBaseCspError = -2147415684
	CSSMERR_CSP_INVALID_ATTR_SYMMETRIC_KEY_FORMAT CssmCspBaseCspError = -2147415672
	CSSMERR_CSP_INVALID_ATTR_VERSION              CssmCspBaseCspError = -2147415690
	CSSMERR_CSP_INVALID_ATTR_WRAPPED_KEY_FORMAT   CssmCspBaseCspError = -2147415670
	CSSMERR_CSP_INVALID_CONTEXT                   CssmCspBaseCspError = -2147415760
	CSSMERR_CSP_INVALID_DATA_COUNT                CssmCspBaseCspError = -2147415768
	CSSMERR_CSP_INVALID_DIGEST_ALGORITHM          CssmCspBaseCspError = -2147415723
	CSSMERR_CSP_INVALID_INPUT_VECTOR              CssmCspBaseCspError = -2147415766
	CSSMERR_CSP_INVALID_KEY                       CssmCspBaseCspError = -2147415792
	CSSMERR_CSP_INVALID_KEYATTR_MASK              CssmCspBaseCspError = -2147415780
	CSSMERR_CSP_INVALID_KEYUSAGE_MASK             CssmCspBaseCspError = -2147415782
	CSSMERR_CSP_INVALID_KEY_CLASS                 CssmCspBaseCspError = -2147415790
	CSSMERR_CSP_INVALID_KEY_FORMAT                CssmCspBaseCspError = -2147415776
	CSSMERR_CSP_INVALID_KEY_LABEL                 CssmCspBaseCspError = -2147415778
	CSSMERR_CSP_INVALID_KEY_POINTER               CssmCspBaseCspError = -2147415783
	CSSMERR_CSP_INVALID_KEY_REFERENCE             CssmCspBaseCspError = -2147415791
	CSSMERR_CSP_INVALID_LOGIN_NAME                CssmCspBaseCspError = -2147415727
	CSSMERR_CSP_INVALID_OUTPUT_VECTOR             CssmCspBaseCspError = -2147415765
	CSSMERR_CSP_INVALID_SIGNATURE                 CssmCspBaseCspError = -2147415733
	CSSMERR_CSP_KEY_BLOB_TYPE_INCORRECT           CssmCspBaseCspError = -2147415787
	CSSMERR_CSP_KEY_HEADER_INCONSISTENT           CssmCspBaseCspError = -2147415786
	CSSMERR_CSP_KEY_LABEL_ALREADY_EXISTS          CssmCspBaseCspError = -2147415724
	CSSMERR_CSP_KEY_USAGE_INCORRECT               CssmCspBaseCspError = -2147415788
	CSSMERR_CSP_MISSING_ATTR_ACCESS_CREDENTIALS   CssmCspBaseCspError = -2147415677
	CSSMERR_CSP_MISSING_ATTR_ALG_PARAMS           CssmCspBaseCspError = -2147415703
	CSSMERR_CSP_MISSING_ATTR_BASE                 CssmCspBaseCspError = -2147415685
	CSSMERR_CSP_MISSING_ATTR_BLOCK_SIZE           CssmCspBaseCspError = -2147415737
	CSSMERR_CSP_MISSING_ATTR_DL_DB_HANDLE         CssmCspBaseCspError = -2147415679
	CSSMERR_CSP_MISSING_ATTR_EFFECTIVE_BITS       CssmCspBaseCspError = -2147415695
	CSSMERR_CSP_MISSING_ATTR_END_DATE             CssmCspBaseCspError = -2147415691
	CSSMERR_CSP_MISSING_ATTR_INIT_VECTOR          CssmCspBaseCspError = -2147415751
	CSSMERR_CSP_MISSING_ATTR_ITERATION_COUNT      CssmCspBaseCspError = -2147415681
	CSSMERR_CSP_MISSING_ATTR_KEY                  CssmCspBaseCspError = -2147415753
	CSSMERR_CSP_MISSING_ATTR_KEY_LENGTH           CssmCspBaseCspError = -2147415739
	CSSMERR_CSP_MISSING_ATTR_KEY_TYPE             CssmCspBaseCspError = -2147415699
	CSSMERR_CSP_MISSING_ATTR_LABEL                CssmCspBaseCspError = -2147415701
	CSSMERR_CSP_MISSING_ATTR_MODE                 CssmCspBaseCspError = -2147415697
	CSSMERR_CSP_MISSING_ATTR_OUTPUT_SIZE          CssmCspBaseCspError = -2147415707
	CSSMERR_CSP_MISSING_ATTR_PADDING              CssmCspBaseCspError = -2147415747
	CSSMERR_CSP_MISSING_ATTR_PASSPHRASE           CssmCspBaseCspError = -2147415741
	CSSMERR_CSP_MISSING_ATTR_PRIME                CssmCspBaseCspError = -2147415687
	CSSMERR_CSP_MISSING_ATTR_PRIVATE_KEY_FORMAT   CssmCspBaseCspError = -2147415673
	CSSMERR_CSP_MISSING_ATTR_PUBLIC_KEY_FORMAT    CssmCspBaseCspError = -2147415675
	CSSMERR_CSP_MISSING_ATTR_RANDOM               CssmCspBaseCspError = -2147415745
	CSSMERR_CSP_MISSING_ATTR_ROUNDS               CssmCspBaseCspError = -2147415705
	CSSMERR_CSP_MISSING_ATTR_SALT                 CssmCspBaseCspError = -2147415749
	CSSMERR_CSP_MISSING_ATTR_SEED                 CssmCspBaseCspError = -2147415743
	CSSMERR_CSP_MISSING_ATTR_START_DATE           CssmCspBaseCspError = -2147415693
	CSSMERR_CSP_MISSING_ATTR_SUBPRIME             CssmCspBaseCspError = -2147415683
	CSSMERR_CSP_MISSING_ATTR_SYMMETRIC_KEY_FORMAT CssmCspBaseCspError = -2147415671
	CSSMERR_CSP_MISSING_ATTR_VERSION              CssmCspBaseCspError = -2147415689
	CSSMERR_CSP_MISSING_ATTR_WRAPPED_KEY_FORMAT   CssmCspBaseCspError = -2147415669
	CSSMERR_CSP_NOT_LOGGED_IN                     CssmCspBaseCspError = -2147415801
	CSSMERR_CSP_OUTPUT_LENGTH_ERROR               CssmCspBaseCspError = -2147415806
	CSSMERR_CSP_PRIVATE_KEY_ALREADY_EXISTS        CssmCspBaseCspError = -2147415725
	CSSMERR_CSP_PRIVATE_KEY_NOT_FOUND             CssmCspBaseCspError = -2147415730
	CSSMERR_CSP_PRIVILEGE_NOT_SUPPORTED           CssmCspBaseCspError = -2147415805
	CSSMERR_CSP_PUBLIC_KEY_INCONSISTENT           CssmCspBaseCspError = -2147415729
	CSSMERR_CSP_QUERY_SIZE_UNKNOWN                CssmCspBaseCspError = -2147415732
	CSSMERR_CSP_STAGED_OPERATION_IN_PROGRESS      CssmCspBaseCspError = -2147415736
	CSSMERR_CSP_STAGED_OPERATION_NOT_STARTED      CssmCspBaseCspError = -2147415735
	CSSMERR_CSP_UNSUPPORTED_KEYATTR_MASK          CssmCspBaseCspError = -2147415779
	CSSMERR_CSP_UNSUPPORTED_KEYUSAGE_MASK         CssmCspBaseCspError = -2147415781
	CSSMERR_CSP_UNSUPPORTED_KEY_FORMAT            CssmCspBaseCspError = -2147415785
	CSSMERR_CSP_UNSUPPORTED_KEY_LABEL             CssmCspBaseCspError = -2147415777
	CSSMERR_CSP_UNSUPPORTED_KEY_SIZE              CssmCspBaseCspError = -2147415784
	CSSMERR_CSP_VECTOR_OF_BUFS_UNSUPPORTED        CssmCspBaseCspError = -2147415767
	CSSMERR_CSP_VERIFY_FAILED                     CssmCspBaseCspError = -2147415734
	CSSM_CSP_BASE_CSP_ERROR                       CssmCspBaseCspError = -2147415808
)

func (e CssmCspBaseCspError) String() string {
	switch e {
	case CSSMERR_CSP_ALGID_MISMATCH:
		return "CSSMERR_CSP_ALGID_MISMATCH"
	case CSSMERR_CSP_ALREADY_LOGGED_IN:
		return "CSSMERR_CSP_ALREADY_LOGGED_IN"
	case CSSMERR_CSP_ATTACH_HANDLE_BUSY:
		return "CSSMERR_CSP_ATTACH_HANDLE_BUSY"
	case CSSMERR_CSP_BLOCK_SIZE_MISMATCH:
		return "CSSMERR_CSP_BLOCK_SIZE_MISMATCH"
	case CSSMERR_CSP_CRYPTO_DATA_CALLBACK_FAILED:
		return "CSSMERR_CSP_CRYPTO_DATA_CALLBACK_FAILED"
	case CSSMERR_CSP_DEVICE_ERROR:
		return "CSSMERR_CSP_DEVICE_ERROR"
	case CSSMERR_CSP_DEVICE_MEMORY_ERROR:
		return "CSSMERR_CSP_DEVICE_MEMORY_ERROR"
	case CSSMERR_CSP_DEVICE_VERIFY_FAILED:
		return "CSSMERR_CSP_DEVICE_VERIFY_FAILED"
	case CSSMERR_CSP_INPUT_LENGTH_ERROR:
		return "CSSMERR_CSP_INPUT_LENGTH_ERROR"
	case CSSMERR_CSP_INVALID_ALGORITHM:
		return "CSSMERR_CSP_INVALID_ALGORITHM"
	case CSSMERR_CSP_INVALID_ATTR_ACCESS_CREDENTIALS:
		return "CSSMERR_CSP_INVALID_ATTR_ACCESS_CREDENTIALS"
	case CSSMERR_CSP_INVALID_ATTR_ALG_PARAMS:
		return "CSSMERR_CSP_INVALID_ATTR_ALG_PARAMS"
	case CSSMERR_CSP_INVALID_ATTR_BASE:
		return "CSSMERR_CSP_INVALID_ATTR_BASE"
	case CSSMERR_CSP_INVALID_ATTR_BLOCK_SIZE:
		return "CSSMERR_CSP_INVALID_ATTR_BLOCK_SIZE"
	case CSSMERR_CSP_INVALID_ATTR_DL_DB_HANDLE:
		return "CSSMERR_CSP_INVALID_ATTR_DL_DB_HANDLE"
	case CSSMERR_CSP_INVALID_ATTR_EFFECTIVE_BITS:
		return "CSSMERR_CSP_INVALID_ATTR_EFFECTIVE_BITS"
	case CSSMERR_CSP_INVALID_ATTR_END_DATE:
		return "CSSMERR_CSP_INVALID_ATTR_END_DATE"
	case CSSMERR_CSP_INVALID_ATTR_INIT_VECTOR:
		return "CSSMERR_CSP_INVALID_ATTR_INIT_VECTOR"
	case CSSMERR_CSP_INVALID_ATTR_ITERATION_COUNT:
		return "CSSMERR_CSP_INVALID_ATTR_ITERATION_COUNT"
	case CSSMERR_CSP_INVALID_ATTR_KEY:
		return "CSSMERR_CSP_INVALID_ATTR_KEY"
	case CSSMERR_CSP_INVALID_ATTR_KEY_LENGTH:
		return "CSSMERR_CSP_INVALID_ATTR_KEY_LENGTH"
	case CSSMERR_CSP_INVALID_ATTR_KEY_TYPE:
		return "CSSMERR_CSP_INVALID_ATTR_KEY_TYPE"
	case CSSMERR_CSP_INVALID_ATTR_LABEL:
		return "CSSMERR_CSP_INVALID_ATTR_LABEL"
	case CSSMERR_CSP_INVALID_ATTR_MODE:
		return "CSSMERR_CSP_INVALID_ATTR_MODE"
	case CSSMERR_CSP_INVALID_ATTR_OUTPUT_SIZE:
		return "CSSMERR_CSP_INVALID_ATTR_OUTPUT_SIZE"
	case CSSMERR_CSP_INVALID_ATTR_PADDING:
		return "CSSMERR_CSP_INVALID_ATTR_PADDING"
	case CSSMERR_CSP_INVALID_ATTR_PASSPHRASE:
		return "CSSMERR_CSP_INVALID_ATTR_PASSPHRASE"
	case CSSMERR_CSP_INVALID_ATTR_PRIME:
		return "CSSMERR_CSP_INVALID_ATTR_PRIME"
	case CSSMERR_CSP_INVALID_ATTR_PRIVATE_KEY_FORMAT:
		return "CSSMERR_CSP_INVALID_ATTR_PRIVATE_KEY_FORMAT"
	case CSSMERR_CSP_INVALID_ATTR_PUBLIC_KEY_FORMAT:
		return "CSSMERR_CSP_INVALID_ATTR_PUBLIC_KEY_FORMAT"
	case CSSMERR_CSP_INVALID_ATTR_RANDOM:
		return "CSSMERR_CSP_INVALID_ATTR_RANDOM"
	case CSSMERR_CSP_INVALID_ATTR_ROUNDS:
		return "CSSMERR_CSP_INVALID_ATTR_ROUNDS"
	case CSSMERR_CSP_INVALID_ATTR_SALT:
		return "CSSMERR_CSP_INVALID_ATTR_SALT"
	case CSSMERR_CSP_INVALID_ATTR_SEED:
		return "CSSMERR_CSP_INVALID_ATTR_SEED"
	case CSSMERR_CSP_INVALID_ATTR_START_DATE:
		return "CSSMERR_CSP_INVALID_ATTR_START_DATE"
	case CSSMERR_CSP_INVALID_ATTR_SUBPRIME:
		return "CSSMERR_CSP_INVALID_ATTR_SUBPRIME"
	case CSSMERR_CSP_INVALID_ATTR_SYMMETRIC_KEY_FORMAT:
		return "CSSMERR_CSP_INVALID_ATTR_SYMMETRIC_KEY_FORMAT"
	case CSSMERR_CSP_INVALID_ATTR_VERSION:
		return "CSSMERR_CSP_INVALID_ATTR_VERSION"
	case CSSMERR_CSP_INVALID_ATTR_WRAPPED_KEY_FORMAT:
		return "CSSMERR_CSP_INVALID_ATTR_WRAPPED_KEY_FORMAT"
	case CSSMERR_CSP_INVALID_CONTEXT:
		return "CSSMERR_CSP_INVALID_CONTEXT"
	case CSSMERR_CSP_INVALID_DATA_COUNT:
		return "CSSMERR_CSP_INVALID_DATA_COUNT"
	case CSSMERR_CSP_INVALID_DIGEST_ALGORITHM:
		return "CSSMERR_CSP_INVALID_DIGEST_ALGORITHM"
	case CSSMERR_CSP_INVALID_INPUT_VECTOR:
		return "CSSMERR_CSP_INVALID_INPUT_VECTOR"
	case CSSMERR_CSP_INVALID_KEY:
		return "CSSMERR_CSP_INVALID_KEY"
	case CSSMERR_CSP_INVALID_KEYATTR_MASK:
		return "CSSMERR_CSP_INVALID_KEYATTR_MASK"
	case CSSMERR_CSP_INVALID_KEYUSAGE_MASK:
		return "CSSMERR_CSP_INVALID_KEYUSAGE_MASK"
	case CSSMERR_CSP_INVALID_KEY_CLASS:
		return "CSSMERR_CSP_INVALID_KEY_CLASS"
	case CSSMERR_CSP_INVALID_KEY_FORMAT:
		return "CSSMERR_CSP_INVALID_KEY_FORMAT"
	case CSSMERR_CSP_INVALID_KEY_LABEL:
		return "CSSMERR_CSP_INVALID_KEY_LABEL"
	case CSSMERR_CSP_INVALID_KEY_POINTER:
		return "CSSMERR_CSP_INVALID_KEY_POINTER"
	case CSSMERR_CSP_INVALID_KEY_REFERENCE:
		return "CSSMERR_CSP_INVALID_KEY_REFERENCE"
	case CSSMERR_CSP_INVALID_LOGIN_NAME:
		return "CSSMERR_CSP_INVALID_LOGIN_NAME"
	case CSSMERR_CSP_INVALID_OUTPUT_VECTOR:
		return "CSSMERR_CSP_INVALID_OUTPUT_VECTOR"
	case CSSMERR_CSP_INVALID_SIGNATURE:
		return "CSSMERR_CSP_INVALID_SIGNATURE"
	case CSSMERR_CSP_KEY_BLOB_TYPE_INCORRECT:
		return "CSSMERR_CSP_KEY_BLOB_TYPE_INCORRECT"
	case CSSMERR_CSP_KEY_HEADER_INCONSISTENT:
		return "CSSMERR_CSP_KEY_HEADER_INCONSISTENT"
	case CSSMERR_CSP_KEY_LABEL_ALREADY_EXISTS:
		return "CSSMERR_CSP_KEY_LABEL_ALREADY_EXISTS"
	case CSSMERR_CSP_KEY_USAGE_INCORRECT:
		return "CSSMERR_CSP_KEY_USAGE_INCORRECT"
	case CSSMERR_CSP_MISSING_ATTR_ACCESS_CREDENTIALS:
		return "CSSMERR_CSP_MISSING_ATTR_ACCESS_CREDENTIALS"
	case CSSMERR_CSP_MISSING_ATTR_ALG_PARAMS:
		return "CSSMERR_CSP_MISSING_ATTR_ALG_PARAMS"
	case CSSMERR_CSP_MISSING_ATTR_BASE:
		return "CSSMERR_CSP_MISSING_ATTR_BASE"
	case CSSMERR_CSP_MISSING_ATTR_BLOCK_SIZE:
		return "CSSMERR_CSP_MISSING_ATTR_BLOCK_SIZE"
	case CSSMERR_CSP_MISSING_ATTR_DL_DB_HANDLE:
		return "CSSMERR_CSP_MISSING_ATTR_DL_DB_HANDLE"
	case CSSMERR_CSP_MISSING_ATTR_EFFECTIVE_BITS:
		return "CSSMERR_CSP_MISSING_ATTR_EFFECTIVE_BITS"
	case CSSMERR_CSP_MISSING_ATTR_END_DATE:
		return "CSSMERR_CSP_MISSING_ATTR_END_DATE"
	case CSSMERR_CSP_MISSING_ATTR_INIT_VECTOR:
		return "CSSMERR_CSP_MISSING_ATTR_INIT_VECTOR"
	case CSSMERR_CSP_MISSING_ATTR_ITERATION_COUNT:
		return "CSSMERR_CSP_MISSING_ATTR_ITERATION_COUNT"
	case CSSMERR_CSP_MISSING_ATTR_KEY:
		return "CSSMERR_CSP_MISSING_ATTR_KEY"
	case CSSMERR_CSP_MISSING_ATTR_KEY_LENGTH:
		return "CSSMERR_CSP_MISSING_ATTR_KEY_LENGTH"
	case CSSMERR_CSP_MISSING_ATTR_KEY_TYPE:
		return "CSSMERR_CSP_MISSING_ATTR_KEY_TYPE"
	case CSSMERR_CSP_MISSING_ATTR_LABEL:
		return "CSSMERR_CSP_MISSING_ATTR_LABEL"
	case CSSMERR_CSP_MISSING_ATTR_MODE:
		return "CSSMERR_CSP_MISSING_ATTR_MODE"
	case CSSMERR_CSP_MISSING_ATTR_OUTPUT_SIZE:
		return "CSSMERR_CSP_MISSING_ATTR_OUTPUT_SIZE"
	case CSSMERR_CSP_MISSING_ATTR_PADDING:
		return "CSSMERR_CSP_MISSING_ATTR_PADDING"
	case CSSMERR_CSP_MISSING_ATTR_PASSPHRASE:
		return "CSSMERR_CSP_MISSING_ATTR_PASSPHRASE"
	case CSSMERR_CSP_MISSING_ATTR_PRIME:
		return "CSSMERR_CSP_MISSING_ATTR_PRIME"
	case CSSMERR_CSP_MISSING_ATTR_PRIVATE_KEY_FORMAT:
		return "CSSMERR_CSP_MISSING_ATTR_PRIVATE_KEY_FORMAT"
	case CSSMERR_CSP_MISSING_ATTR_PUBLIC_KEY_FORMAT:
		return "CSSMERR_CSP_MISSING_ATTR_PUBLIC_KEY_FORMAT"
	case CSSMERR_CSP_MISSING_ATTR_RANDOM:
		return "CSSMERR_CSP_MISSING_ATTR_RANDOM"
	case CSSMERR_CSP_MISSING_ATTR_ROUNDS:
		return "CSSMERR_CSP_MISSING_ATTR_ROUNDS"
	case CSSMERR_CSP_MISSING_ATTR_SALT:
		return "CSSMERR_CSP_MISSING_ATTR_SALT"
	case CSSMERR_CSP_MISSING_ATTR_SEED:
		return "CSSMERR_CSP_MISSING_ATTR_SEED"
	case CSSMERR_CSP_MISSING_ATTR_START_DATE:
		return "CSSMERR_CSP_MISSING_ATTR_START_DATE"
	case CSSMERR_CSP_MISSING_ATTR_SUBPRIME:
		return "CSSMERR_CSP_MISSING_ATTR_SUBPRIME"
	case CSSMERR_CSP_MISSING_ATTR_SYMMETRIC_KEY_FORMAT:
		return "CSSMERR_CSP_MISSING_ATTR_SYMMETRIC_KEY_FORMAT"
	case CSSMERR_CSP_MISSING_ATTR_VERSION:
		return "CSSMERR_CSP_MISSING_ATTR_VERSION"
	case CSSMERR_CSP_MISSING_ATTR_WRAPPED_KEY_FORMAT:
		return "CSSMERR_CSP_MISSING_ATTR_WRAPPED_KEY_FORMAT"
	case CSSMERR_CSP_NOT_LOGGED_IN:
		return "CSSMERR_CSP_NOT_LOGGED_IN"
	case CSSMERR_CSP_OUTPUT_LENGTH_ERROR:
		return "CSSMERR_CSP_OUTPUT_LENGTH_ERROR"
	case CSSMERR_CSP_PRIVATE_KEY_ALREADY_EXISTS:
		return "CSSMERR_CSP_PRIVATE_KEY_ALREADY_EXISTS"
	case CSSMERR_CSP_PRIVATE_KEY_NOT_FOUND:
		return "CSSMERR_CSP_PRIVATE_KEY_NOT_FOUND"
	case CSSMERR_CSP_PRIVILEGE_NOT_SUPPORTED:
		return "CSSMERR_CSP_PRIVILEGE_NOT_SUPPORTED"
	case CSSMERR_CSP_PUBLIC_KEY_INCONSISTENT:
		return "CSSMERR_CSP_PUBLIC_KEY_INCONSISTENT"
	case CSSMERR_CSP_QUERY_SIZE_UNKNOWN:
		return "CSSMERR_CSP_QUERY_SIZE_UNKNOWN"
	case CSSMERR_CSP_STAGED_OPERATION_IN_PROGRESS:
		return "CSSMERR_CSP_STAGED_OPERATION_IN_PROGRESS"
	case CSSMERR_CSP_STAGED_OPERATION_NOT_STARTED:
		return "CSSMERR_CSP_STAGED_OPERATION_NOT_STARTED"
	case CSSMERR_CSP_UNSUPPORTED_KEYATTR_MASK:
		return "CSSMERR_CSP_UNSUPPORTED_KEYATTR_MASK"
	case CSSMERR_CSP_UNSUPPORTED_KEYUSAGE_MASK:
		return "CSSMERR_CSP_UNSUPPORTED_KEYUSAGE_MASK"
	case CSSMERR_CSP_UNSUPPORTED_KEY_FORMAT:
		return "CSSMERR_CSP_UNSUPPORTED_KEY_FORMAT"
	case CSSMERR_CSP_UNSUPPORTED_KEY_LABEL:
		return "CSSMERR_CSP_UNSUPPORTED_KEY_LABEL"
	case CSSMERR_CSP_UNSUPPORTED_KEY_SIZE:
		return "CSSMERR_CSP_UNSUPPORTED_KEY_SIZE"
	case CSSMERR_CSP_VECTOR_OF_BUFS_UNSUPPORTED:
		return "CSSMERR_CSP_VECTOR_OF_BUFS_UNSUPPORTED"
	case CSSMERR_CSP_VERIFY_FAILED:
		return "CSSMERR_CSP_VERIFY_FAILED"
	case CSSM_CSP_BASE_CSP_ERROR:
		return "CSSM_CSP_BASE_CSP_ERROR"
	default:
		return fmt.Sprintf("CssmCspBaseCspError(%d)", e)
	}
}

type CssmCspRdr uint32

const (
	CSSM_CSP_RDR_EXISTS       CssmCspRdr = 0x2
	CSSM_CSP_RDR_HW           CssmCspRdr = 0x4
	CSSM_CSP_RDR_TOKENPRESENT CssmCspRdr = 0x1
)

func (e CssmCspRdr) String() string {
	switch e {
	case CSSM_CSP_RDR_EXISTS:
		return "CSSM_CSP_RDR_EXISTS"
	case CSSM_CSP_RDR_HW:
		return "CSSM_CSP_RDR_HW"
	case CSSM_CSP_RDR_TOKENPRESENT:
		return "CSSM_CSP_RDR_TOKENPRESENT"
	default:
		return fmt.Sprintf("CssmCspRdr(%d)", e)
	}
}

type CssmCspSoftware uint32

const (
	CSSM_CSP_HARDWARE CssmCspSoftware = 2
	CSSM_CSP_HYBRID   CssmCspSoftware = 3
	CSSM_CSP_SOFTWARE CssmCspSoftware = 1
)

func (e CssmCspSoftware) String() string {
	switch e {
	case CSSM_CSP_HARDWARE:
		return "CSSM_CSP_HARDWARE"
	case CSSM_CSP_HYBRID:
		return "CSSM_CSP_HYBRID"
	case CSSM_CSP_SOFTWARE:
		return "CSSM_CSP_SOFTWARE"
	default:
		return fmt.Sprintf("CssmCspSoftware(%d)", e)
	}
}

type CssmCspTok uint32

const (
	CSSM_CSP_TOK_CLOCK_EXISTS CssmCspTok = 0x40
	CSSM_CSP_TOK_RNG          CssmCspTok = 0x1
)

func (e CssmCspTok) String() string {
	switch e {
	case CSSM_CSP_TOK_CLOCK_EXISTS:
		return "CSSM_CSP_TOK_CLOCK_EXISTS"
	case CSSM_CSP_TOK_RNG:
		return "CSSM_CSP_TOK_RNG"
	default:
		return fmt.Sprintf("CssmCspTok(%d)", e)
	}
}

type CssmCspTokWriteProtected uint32

const (
	CSSM_CSP_STORES_CERTIFICATES      CssmCspTokWriteProtected = 0x8000000
	CSSM_CSP_STORES_GENERIC           CssmCspTokWriteProtected = 0x10000000
	CSSM_CSP_STORES_PRIVATE_KEYS      CssmCspTokWriteProtected = 0x1000000
	CSSM_CSP_STORES_PUBLIC_KEYS       CssmCspTokWriteProtected = 0x2000000
	CSSM_CSP_STORES_SESSION_KEYS      CssmCspTokWriteProtected = 0x4000000
	CSSM_CSP_TOK_LOGIN_REQUIRED       CssmCspTokWriteProtected = 0x4
	CSSM_CSP_TOK_PRIVATE_KEY_PASSWORD CssmCspTokWriteProtected = 0x400000
	CSSM_CSP_TOK_PROT_AUTHENTICATION  CssmCspTokWriteProtected = 0x100
	CSSM_CSP_TOK_SESSION_KEY_PASSWORD CssmCspTokWriteProtected = 0x200000
	CSSM_CSP_TOK_USER_PIN_EXPIRED     CssmCspTokWriteProtected = 0x100000
	CSSM_CSP_TOK_USER_PIN_INITIALIZED CssmCspTokWriteProtected = 0x8
	CSSM_CSP_TOK_WRITE_PROTECTED      CssmCspTokWriteProtected = 0x2
)

func (e CssmCspTokWriteProtected) String() string {
	switch e {
	case CSSM_CSP_STORES_CERTIFICATES:
		return "CSSM_CSP_STORES_CERTIFICATES"
	case CSSM_CSP_STORES_GENERIC:
		return "CSSM_CSP_STORES_GENERIC"
	case CSSM_CSP_STORES_PRIVATE_KEYS:
		return "CSSM_CSP_STORES_PRIVATE_KEYS"
	case CSSM_CSP_STORES_PUBLIC_KEYS:
		return "CSSM_CSP_STORES_PUBLIC_KEYS"
	case CSSM_CSP_STORES_SESSION_KEYS:
		return "CSSM_CSP_STORES_SESSION_KEYS"
	case CSSM_CSP_TOK_LOGIN_REQUIRED:
		return "CSSM_CSP_TOK_LOGIN_REQUIRED"
	case CSSM_CSP_TOK_PRIVATE_KEY_PASSWORD:
		return "CSSM_CSP_TOK_PRIVATE_KEY_PASSWORD"
	case CSSM_CSP_TOK_PROT_AUTHENTICATION:
		return "CSSM_CSP_TOK_PROT_AUTHENTICATION"
	case CSSM_CSP_TOK_SESSION_KEY_PASSWORD:
		return "CSSM_CSP_TOK_SESSION_KEY_PASSWORD"
	case CSSM_CSP_TOK_USER_PIN_EXPIRED:
		return "CSSM_CSP_TOK_USER_PIN_EXPIRED"
	case CSSM_CSP_TOK_USER_PIN_INITIALIZED:
		return "CSSM_CSP_TOK_USER_PIN_INITIALIZED"
	case CSSM_CSP_TOK_WRITE_PROTECTED:
		return "CSSM_CSP_TOK_WRITE_PROTECTED"
	default:
		return fmt.Sprintf("CssmCspTokWriteProtected(%d)", e)
	}
}

type CssmCssmBaseCssmError int32

const (
	CSSMERR_CSSM_ADDIN_AUTHENTICATE_FAILED             CssmCssmBaseCssmError = -2147417828
	CSSMERR_CSSM_ADDIN_LOAD_FAILED                     CssmCssmBaseCssmError = -2147417834
	CSSMERR_CSSM_ADDIN_UNLOAD_FAILED                   CssmCssmBaseCssmError = -2147417832
	CSSMERR_CSSM_ATTRIBUTE_NOT_IN_CONTEXT              CssmCssmBaseCssmError = -2147417822
	CSSMERR_CSSM_BUFFER_TOO_SMALL                      CssmCssmBaseCssmError = -2147417824
	CSSMERR_CSSM_EMM_AUTHENTICATE_FAILED               CssmCssmBaseCssmError = -2147417829
	CSSMERR_CSSM_EMM_LOAD_FAILED                       CssmCssmBaseCssmError = -2147417836
	CSSMERR_CSSM_EMM_UNLOAD_FAILED                     CssmCssmBaseCssmError = -2147417835
	CSSMERR_CSSM_EVENT_NOTIFICATION_CALLBACK_NOT_FOUND CssmCssmBaseCssmError = -2147417819
	CSSMERR_CSSM_INVALID_ADDIN_FUNCTION_TABLE          CssmCssmBaseCssmError = -2147417830
	CSSMERR_CSSM_INVALID_ATTRIBUTE                     CssmCssmBaseCssmError = -2147417823
	CSSMERR_CSSM_INVALID_KEY_HIERARCHY                 CssmCssmBaseCssmError = -2147417833
	CSSMERR_CSSM_INVALID_PVC                           CssmCssmBaseCssmError = -2147417837
	CSSMERR_CSSM_INVALID_SERVICE_MASK                  CssmCssmBaseCssmError = -2147417827
	CSSMERR_CSSM_INVALID_SUBSERVICEID                  CssmCssmBaseCssmError = -2147417825
	CSSMERR_CSSM_LIB_REF_NOT_FOUND                     CssmCssmBaseCssmError = -2147417831
	CSSMERR_CSSM_MODULE_MANAGER_INITIALIZE_FAIL        CssmCssmBaseCssmError = -2147417821
	CSSMERR_CSSM_MODULE_MANAGER_NOT_FOUND              CssmCssmBaseCssmError = -2147417820
	CSSMERR_CSSM_MODULE_NOT_LOADED                     CssmCssmBaseCssmError = -2147417826
	CSSMERR_CSSM_PVC_ALREADY_CONFIGURED                CssmCssmBaseCssmError = -2147417838
	CSSMERR_CSSM_SCOPE_NOT_SUPPORTED                   CssmCssmBaseCssmError = -2147417839
	CSSM_CSSM_BASE_CSSM_ERROR                          CssmCssmBaseCssmError = -2147417840
)

func (e CssmCssmBaseCssmError) String() string {
	switch e {
	case CSSMERR_CSSM_ADDIN_AUTHENTICATE_FAILED:
		return "CSSMERR_CSSM_ADDIN_AUTHENTICATE_FAILED"
	case CSSMERR_CSSM_ADDIN_LOAD_FAILED:
		return "CSSMERR_CSSM_ADDIN_LOAD_FAILED"
	case CSSMERR_CSSM_ADDIN_UNLOAD_FAILED:
		return "CSSMERR_CSSM_ADDIN_UNLOAD_FAILED"
	case CSSMERR_CSSM_ATTRIBUTE_NOT_IN_CONTEXT:
		return "CSSMERR_CSSM_ATTRIBUTE_NOT_IN_CONTEXT"
	case CSSMERR_CSSM_BUFFER_TOO_SMALL:
		return "CSSMERR_CSSM_BUFFER_TOO_SMALL"
	case CSSMERR_CSSM_EMM_AUTHENTICATE_FAILED:
		return "CSSMERR_CSSM_EMM_AUTHENTICATE_FAILED"
	case CSSMERR_CSSM_EMM_LOAD_FAILED:
		return "CSSMERR_CSSM_EMM_LOAD_FAILED"
	case CSSMERR_CSSM_EMM_UNLOAD_FAILED:
		return "CSSMERR_CSSM_EMM_UNLOAD_FAILED"
	case CSSMERR_CSSM_EVENT_NOTIFICATION_CALLBACK_NOT_FOUND:
		return "CSSMERR_CSSM_EVENT_NOTIFICATION_CALLBACK_NOT_FOUND"
	case CSSMERR_CSSM_INVALID_ADDIN_FUNCTION_TABLE:
		return "CSSMERR_CSSM_INVALID_ADDIN_FUNCTION_TABLE"
	case CSSMERR_CSSM_INVALID_ATTRIBUTE:
		return "CSSMERR_CSSM_INVALID_ATTRIBUTE"
	case CSSMERR_CSSM_INVALID_KEY_HIERARCHY:
		return "CSSMERR_CSSM_INVALID_KEY_HIERARCHY"
	case CSSMERR_CSSM_INVALID_PVC:
		return "CSSMERR_CSSM_INVALID_PVC"
	case CSSMERR_CSSM_INVALID_SERVICE_MASK:
		return "CSSMERR_CSSM_INVALID_SERVICE_MASK"
	case CSSMERR_CSSM_INVALID_SUBSERVICEID:
		return "CSSMERR_CSSM_INVALID_SUBSERVICEID"
	case CSSMERR_CSSM_LIB_REF_NOT_FOUND:
		return "CSSMERR_CSSM_LIB_REF_NOT_FOUND"
	case CSSMERR_CSSM_MODULE_MANAGER_INITIALIZE_FAIL:
		return "CSSMERR_CSSM_MODULE_MANAGER_INITIALIZE_FAIL"
	case CSSMERR_CSSM_MODULE_MANAGER_NOT_FOUND:
		return "CSSMERR_CSSM_MODULE_MANAGER_NOT_FOUND"
	case CSSMERR_CSSM_MODULE_NOT_LOADED:
		return "CSSMERR_CSSM_MODULE_NOT_LOADED"
	case CSSMERR_CSSM_PVC_ALREADY_CONFIGURED:
		return "CSSMERR_CSSM_PVC_ALREADY_CONFIGURED"
	case CSSMERR_CSSM_SCOPE_NOT_SUPPORTED:
		return "CSSMERR_CSSM_SCOPE_NOT_SUPPORTED"
	case CSSM_CSSM_BASE_CSSM_ERROR:
		return "CSSM_CSSM_BASE_CSSM_ERROR"
	default:
		return fmt.Sprintf("CssmCssmBaseCssmError(%d)", e)
	}
}

type CssmCustomCommonErrorExtent uint32

const (
	CSSM_CUSTOM_COMMON_ERROR_EXTENT                 CssmCustomCommonErrorExtent = 0xe0
	CSSM_ERRCODE_DEVICE_FAILED                      CssmCustomCommonErrorExtent = 0xe5
	CSSM_ERRCODE_DEVICE_RESET                       CssmCustomCommonErrorExtent = 0xe4
	CSSM_ERRCODE_INSUFFICIENT_CLIENT_IDENTIFICATION CssmCustomCommonErrorExtent = 0xe3
	CSSM_ERRCODE_IN_DARK_WAKE                       CssmCustomCommonErrorExtent = 0xe6
	CSSM_ERRCODE_NO_USER_INTERACTION                CssmCustomCommonErrorExtent = 0xe0
	CSSM_ERRCODE_SERVICE_NOT_AVAILABLE              CssmCustomCommonErrorExtent = 0xe2
	CSSM_ERRCODE_USER_CANCELED                      CssmCustomCommonErrorExtent = 0xe1
)

func (e CssmCustomCommonErrorExtent) String() string {
	switch e {
	case CSSM_CUSTOM_COMMON_ERROR_EXTENT:
		return "CSSM_CUSTOM_COMMON_ERROR_EXTENT"
	case CSSM_ERRCODE_DEVICE_FAILED:
		return "CSSM_ERRCODE_DEVICE_FAILED"
	case CSSM_ERRCODE_DEVICE_RESET:
		return "CSSM_ERRCODE_DEVICE_RESET"
	case CSSM_ERRCODE_INSUFFICIENT_CLIENT_IDENTIFICATION:
		return "CSSM_ERRCODE_INSUFFICIENT_CLIENT_IDENTIFICATION"
	case CSSM_ERRCODE_IN_DARK_WAKE:
		return "CSSM_ERRCODE_IN_DARK_WAKE"
	case CSSM_ERRCODE_SERVICE_NOT_AVAILABLE:
		return "CSSM_ERRCODE_SERVICE_NOT_AVAILABLE"
	case CSSM_ERRCODE_USER_CANCELED:
		return "CSSM_ERRCODE_USER_CANCELED"
	default:
		return fmt.Sprintf("CssmCustomCommonErrorExtent(%d)", e)
	}
}

type CssmD uint32

const (
	CSSM_DB_RECORDTYPE_APP_DEFINED_END   CssmD = 0xffffffff
	CSSM_DB_RECORDTYPE_APP_DEFINED_START CssmD = 0x80000000
	CSSM_DB_RECORDTYPE_OPEN_GROUP_END    CssmD = 18
	CSSM_DB_RECORDTYPE_OPEN_GROUP_START  CssmD = 0xa
	CSSM_DB_RECORDTYPE_SCHEMA_END        CssmD = 4
	CSSM_DB_RECORDTYPE_SCHEMA_START      CssmD = 0
	CSSM_DL_DB_RECORD_ALL_KEYS           CssmD = 18
	CSSM_DL_DB_RECORD_ANY                CssmD = 10
	CSSM_DL_DB_RECORD_CERT               CssmD = 11
	CSSM_DL_DB_RECORD_CRL                CssmD = 12
	CSSM_DL_DB_RECORD_GENERIC            CssmD = 14
	CSSM_DL_DB_RECORD_POLICY             CssmD = 13
	CSSM_DL_DB_RECORD_PRIVATE_KEY        CssmD = 16
	CSSM_DL_DB_RECORD_PUBLIC_KEY         CssmD = 15
	CSSM_DL_DB_RECORD_SYMMETRIC_KEY      CssmD = 17
	CSSM_DL_DB_SCHEMA_ATTRIBUTES         CssmD = 2
	CSSM_DL_DB_SCHEMA_INDEXES            CssmD = 1
	CSSM_DL_DB_SCHEMA_INFO               CssmD = 0
	CSSM_DL_DB_SCHEMA_PARSING_MODULE     CssmD = 3
)

func (e CssmD) String() string {
	switch e {
	case CSSM_DB_RECORDTYPE_APP_DEFINED_END:
		return "CSSM_DB_RECORDTYPE_APP_DEFINED_END"
	case CSSM_DB_RECORDTYPE_APP_DEFINED_START:
		return "CSSM_DB_RECORDTYPE_APP_DEFINED_START"
	case CSSM_DB_RECORDTYPE_OPEN_GROUP_END:
		return "CSSM_DB_RECORDTYPE_OPEN_GROUP_END"
	case CSSM_DB_RECORDTYPE_OPEN_GROUP_START:
		return "CSSM_DB_RECORDTYPE_OPEN_GROUP_START"
	case CSSM_DB_RECORDTYPE_SCHEMA_END:
		return "CSSM_DB_RECORDTYPE_SCHEMA_END"
	case CSSM_DB_RECORDTYPE_SCHEMA_START:
		return "CSSM_DB_RECORDTYPE_SCHEMA_START"
	case CSSM_DL_DB_RECORD_CERT:
		return "CSSM_DL_DB_RECORD_CERT"
	case CSSM_DL_DB_RECORD_CRL:
		return "CSSM_DL_DB_RECORD_CRL"
	case CSSM_DL_DB_RECORD_GENERIC:
		return "CSSM_DL_DB_RECORD_GENERIC"
	case CSSM_DL_DB_RECORD_POLICY:
		return "CSSM_DL_DB_RECORD_POLICY"
	case CSSM_DL_DB_RECORD_PRIVATE_KEY:
		return "CSSM_DL_DB_RECORD_PRIVATE_KEY"
	case CSSM_DL_DB_RECORD_PUBLIC_KEY:
		return "CSSM_DL_DB_RECORD_PUBLIC_KEY"
	case CSSM_DL_DB_RECORD_SYMMETRIC_KEY:
		return "CSSM_DL_DB_RECORD_SYMMETRIC_KEY"
	case CSSM_DL_DB_SCHEMA_ATTRIBUTES:
		return "CSSM_DL_DB_SCHEMA_ATTRIBUTES"
	case CSSM_DL_DB_SCHEMA_INDEXES:
		return "CSSM_DL_DB_SCHEMA_INDEXES"
	case CSSM_DL_DB_SCHEMA_PARSING_MODULE:
		return "CSSM_DL_DB_SCHEMA_PARSING_MODULE"
	default:
		return fmt.Sprintf("CssmD(%d)", e)
	}
}

type CssmDbAccessRead uint32

const (
	CSSM_DB_ACCESS_PRIVILEGED CssmDbAccessRead = 0x4
	CSSM_DB_ACCESS_READ       CssmDbAccessRead = 0x1
	CSSM_DB_ACCESS_WRITE      CssmDbAccessRead = 0x2
)

func (e CssmDbAccessRead) String() string {
	switch e {
	case CSSM_DB_ACCESS_PRIVILEGED:
		return "CSSM_DB_ACCESS_PRIVILEGED"
	case CSSM_DB_ACCESS_READ:
		return "CSSM_DB_ACCESS_READ"
	case CSSM_DB_ACCESS_WRITE:
		return "CSSM_DB_ACCESS_WRITE"
	default:
		return fmt.Sprintf("CssmDbAccessRead(%d)", e)
	}
}

type CssmDbAccessReset uint32

const (
	CSSM_DB_ACCESS_RESET CssmDbAccessReset = 0x10000
)

func (e CssmDbAccessReset) String() string {
	switch e {
	case CSSM_DB_ACCESS_RESET:
		return "CSSM_DB_ACCESS_RESET"
	default:
		return fmt.Sprintf("CssmDbAccessReset(%d)", e)
	}
}

type CssmDbAttributeFormat uint32

const (
	CSSM_DB_ATTRIBUTE_FORMAT_BIG_NUM      CssmDbAttributeFormat = 3
	CSSM_DB_ATTRIBUTE_FORMAT_BLOB         CssmDbAttributeFormat = 6
	CSSM_DB_ATTRIBUTE_FORMAT_COMPLEX      CssmDbAttributeFormat = 8
	CSSM_DB_ATTRIBUTE_FORMAT_MULTI_UINT32 CssmDbAttributeFormat = 7
	CSSM_DB_ATTRIBUTE_FORMAT_REAL         CssmDbAttributeFormat = 4
	CSSM_DB_ATTRIBUTE_FORMAT_SINT32       CssmDbAttributeFormat = 1
	CSSM_DB_ATTRIBUTE_FORMAT_STRING       CssmDbAttributeFormat = 0
	CSSM_DB_ATTRIBUTE_FORMAT_TIME_DATE    CssmDbAttributeFormat = 5
	CSSM_DB_ATTRIBUTE_FORMAT_UINT32       CssmDbAttributeFormat = 2
)

func (e CssmDbAttributeFormat) String() string {
	switch e {
	case CSSM_DB_ATTRIBUTE_FORMAT_BIG_NUM:
		return "CSSM_DB_ATTRIBUTE_FORMAT_BIG_NUM"
	case CSSM_DB_ATTRIBUTE_FORMAT_BLOB:
		return "CSSM_DB_ATTRIBUTE_FORMAT_BLOB"
	case CSSM_DB_ATTRIBUTE_FORMAT_COMPLEX:
		return "CSSM_DB_ATTRIBUTE_FORMAT_COMPLEX"
	case CSSM_DB_ATTRIBUTE_FORMAT_MULTI_UINT32:
		return "CSSM_DB_ATTRIBUTE_FORMAT_MULTI_UINT32"
	case CSSM_DB_ATTRIBUTE_FORMAT_REAL:
		return "CSSM_DB_ATTRIBUTE_FORMAT_REAL"
	case CSSM_DB_ATTRIBUTE_FORMAT_SINT32:
		return "CSSM_DB_ATTRIBUTE_FORMAT_SINT32"
	case CSSM_DB_ATTRIBUTE_FORMAT_STRING:
		return "CSSM_DB_ATTRIBUTE_FORMAT_STRING"
	case CSSM_DB_ATTRIBUTE_FORMAT_TIME_DATE:
		return "CSSM_DB_ATTRIBUTE_FORMAT_TIME_DATE"
	case CSSM_DB_ATTRIBUTE_FORMAT_UINT32:
		return "CSSM_DB_ATTRIBUTE_FORMAT_UINT32"
	default:
		return fmt.Sprintf("CssmDbAttributeFormat(%d)", e)
	}
}

type CssmDbAttributeNameAs uint32

const (
	CSSM_DB_ATTRIBUTE_NAME_AS_INTEGER CssmDbAttributeNameAs = 2
	CSSM_DB_ATTRIBUTE_NAME_AS_OID     CssmDbAttributeNameAs = 1
	CSSM_DB_ATTRIBUTE_NAME_AS_STRING  CssmDbAttributeNameAs = 0
)

func (e CssmDbAttributeNameAs) String() string {
	switch e {
	case CSSM_DB_ATTRIBUTE_NAME_AS_INTEGER:
		return "CSSM_DB_ATTRIBUTE_NAME_AS_INTEGER"
	case CSSM_DB_ATTRIBUTE_NAME_AS_OID:
		return "CSSM_DB_ATTRIBUTE_NAME_AS_OID"
	case CSSM_DB_ATTRIBUTE_NAME_AS_STRING:
		return "CSSM_DB_ATTRIBUTE_NAME_AS_STRING"
	default:
		return fmt.Sprintf("CssmDbAttributeNameAs(%d)", e)
	}
}

type CssmDbCertUse uint32

const (
	CSSM_DB_CERT_USE_OWNER   CssmDbCertUse = 0x4
	CSSM_DB_CERT_USE_PRIVACY CssmDbCertUse = 0x20
	CSSM_DB_CERT_USE_REVOKED CssmDbCertUse = 0x8
	CSSM_DB_CERT_USE_SIGNING CssmDbCertUse = 0x10
	CSSM_DB_CERT_USE_SYSTEM  CssmDbCertUse = 0x2
	CSSM_DB_CERT_USE_TRUSTED CssmDbCertUse = 0x1
)

func (e CssmDbCertUse) String() string {
	switch e {
	case CSSM_DB_CERT_USE_OWNER:
		return "CSSM_DB_CERT_USE_OWNER"
	case CSSM_DB_CERT_USE_PRIVACY:
		return "CSSM_DB_CERT_USE_PRIVACY"
	case CSSM_DB_CERT_USE_REVOKED:
		return "CSSM_DB_CERT_USE_REVOKED"
	case CSSM_DB_CERT_USE_SIGNING:
		return "CSSM_DB_CERT_USE_SIGNING"
	case CSSM_DB_CERT_USE_SYSTEM:
		return "CSSM_DB_CERT_USE_SYSTEM"
	case CSSM_DB_CERT_USE_TRUSTED:
		return "CSSM_DB_CERT_USE_TRUSTED"
	default:
		return fmt.Sprintf("CssmDbCertUse(%d)", e)
	}
}

type CssmDbDatastores uint32

const (
	CSSM_DB_DATASTORES_UNKNOWN CssmDbDatastores = 0xffffffff
)

func (e CssmDbDatastores) String() string {
	switch e {
	case CSSM_DB_DATASTORES_UNKNOWN:
		return "CSSM_DB_DATASTORES_UNKNOWN"
	default:
		return fmt.Sprintf("CssmDbDatastores(%d)", e)
	}
}

type CssmDbEqual uint32

const (
	CSSM_DB_CONTAINS                   CssmDbEqual = 4
	CSSM_DB_CONTAINS_FINAL_SUBSTRING   CssmDbEqual = 6
	CSSM_DB_CONTAINS_INITIAL_SUBSTRING CssmDbEqual = 5
	CSSM_DB_EQUAL                      CssmDbEqual = 0
	CSSM_DB_GREATER_THAN               CssmDbEqual = 3
	CSSM_DB_LESS_THAN                  CssmDbEqual = 2
	CSSM_DB_NOT_EQUAL                  CssmDbEqual = 1
)

func (e CssmDbEqual) String() string {
	switch e {
	case CSSM_DB_CONTAINS:
		return "CSSM_DB_CONTAINS"
	case CSSM_DB_CONTAINS_FINAL_SUBSTRING:
		return "CSSM_DB_CONTAINS_FINAL_SUBSTRING"
	case CSSM_DB_CONTAINS_INITIAL_SUBSTRING:
		return "CSSM_DB_CONTAINS_INITIAL_SUBSTRING"
	case CSSM_DB_EQUAL:
		return "CSSM_DB_EQUAL"
	case CSSM_DB_GREATER_THAN:
		return "CSSM_DB_GREATER_THAN"
	case CSSM_DB_LESS_THAN:
		return "CSSM_DB_LESS_THAN"
	case CSSM_DB_NOT_EQUAL:
		return "CSSM_DB_NOT_EQUAL"
	default:
		return fmt.Sprintf("CssmDbEqual(%d)", e)
	}
}

type CssmDbIndex uint32

const (
	CSSM_DB_INDEX_NONUNIQUE CssmDbIndex = 1
	CSSM_DB_INDEX_UNIQUE    CssmDbIndex = 0
)

func (e CssmDbIndex) String() string {
	switch e {
	case CSSM_DB_INDEX_NONUNIQUE:
		return "CSSM_DB_INDEX_NONUNIQUE"
	case CSSM_DB_INDEX_UNIQUE:
		return "CSSM_DB_INDEX_UNIQUE"
	default:
		return fmt.Sprintf("CssmDbIndex(%d)", e)
	}
}

type CssmDbIndexOn uint32

const (
	CSSM_DB_INDEX_ON_ATTRIBUTE CssmDbIndexOn = 1
	CSSM_DB_INDEX_ON_RECORD    CssmDbIndexOn = 2
	CSSM_DB_INDEX_ON_UNKNOWN   CssmDbIndexOn = 0
)

func (e CssmDbIndexOn) String() string {
	switch e {
	case CSSM_DB_INDEX_ON_ATTRIBUTE:
		return "CSSM_DB_INDEX_ON_ATTRIBUTE"
	case CSSM_DB_INDEX_ON_RECORD:
		return "CSSM_DB_INDEX_ON_RECORD"
	case CSSM_DB_INDEX_ON_UNKNOWN:
		return "CSSM_DB_INDEX_ON_UNKNOWN"
	default:
		return fmt.Sprintf("CssmDbIndexOn(%d)", e)
	}
}

type CssmDbModifyAttribute uint32

const (
	CSSM_DB_MODIFY_ATTRIBUTE_ADD     CssmDbModifyAttribute = 1
	CSSM_DB_MODIFY_ATTRIBUTE_DELETE  CssmDbModifyAttribute = 2
	CSSM_DB_MODIFY_ATTRIBUTE_NONE    CssmDbModifyAttribute = 0
	CSSM_DB_MODIFY_ATTRIBUTE_REPLACE CssmDbModifyAttribute = 3
)

func (e CssmDbModifyAttribute) String() string {
	switch e {
	case CSSM_DB_MODIFY_ATTRIBUTE_ADD:
		return "CSSM_DB_MODIFY_ATTRIBUTE_ADD"
	case CSSM_DB_MODIFY_ATTRIBUTE_DELETE:
		return "CSSM_DB_MODIFY_ATTRIBUTE_DELETE"
	case CSSM_DB_MODIFY_ATTRIBUTE_NONE:
		return "CSSM_DB_MODIFY_ATTRIBUTE_NONE"
	case CSSM_DB_MODIFY_ATTRIBUTE_REPLACE:
		return "CSSM_DB_MODIFY_ATTRIBUTE_REPLACE"
	default:
		return fmt.Sprintf("CssmDbModifyAttribute(%d)", e)
	}
}

type CssmDbNone uint32

const (
	CSSM_DB_AND  CssmDbNone = 1
	CSSM_DB_NONE CssmDbNone = 0
	CSSM_DB_OR   CssmDbNone = 2
)

func (e CssmDbNone) String() string {
	switch e {
	case CSSM_DB_AND:
		return "CSSM_DB_AND"
	case CSSM_DB_NONE:
		return "CSSM_DB_NONE"
	case CSSM_DB_OR:
		return "CSSM_DB_OR"
	default:
		return fmt.Sprintf("CssmDbNone(%d)", e)
	}
}

type CssmDbTransactionalMode uint32

const (
	CSSM_DB_FILESYSTEMSCAN_MODE CssmDbTransactionalMode = 1
	CSSM_DB_TRANSACTIONAL_MODE  CssmDbTransactionalMode = 0
)

func (e CssmDbTransactionalMode) String() string {
	switch e {
	case CSSM_DB_FILESYSTEMSCAN_MODE:
		return "CSSM_DB_FILESYSTEMSCAN_MODE"
	case CSSM_DB_TRANSACTIONAL_MODE:
		return "CSSM_DB_TRANSACTIONAL_MODE"
	default:
		return fmt.Sprintf("CssmDbTransactionalMode(%d)", e)
	}
}

type CssmDl uint32

const (
	CSSM_DL_CUSTOM    CssmDl = 1
	CSSM_DL_FFS       CssmDl = 5
	CSSM_DL_LDAP      CssmDl = 2
	CSSM_DL_MEMORY    CssmDl = 6
	CSSM_DL_ODBC      CssmDl = 3
	CSSM_DL_PKCS11    CssmDl = 4
	CSSM_DL_REMOTEDIR CssmDl = 7
	CSSM_DL_UNKNOWN   CssmDl = 0
)

func (e CssmDl) String() string {
	switch e {
	case CSSM_DL_CUSTOM:
		return "CSSM_DL_CUSTOM"
	case CSSM_DL_FFS:
		return "CSSM_DL_FFS"
	case CSSM_DL_LDAP:
		return "CSSM_DL_LDAP"
	case CSSM_DL_MEMORY:
		return "CSSM_DL_MEMORY"
	case CSSM_DL_ODBC:
		return "CSSM_DL_ODBC"
	case CSSM_DL_PKCS11:
		return "CSSM_DL_PKCS11"
	case CSSM_DL_REMOTEDIR:
		return "CSSM_DL_REMOTEDIR"
	case CSSM_DL_UNKNOWN:
		return "CSSM_DL_UNKNOWN"
	default:
		return fmt.Sprintf("CssmDl(%d)", e)
	}
}

type CssmDlBaseDlError int32

const (
	CSSMERR_DL_DATABASE_CORRUPT                CssmDlBaseDlError = -2147413759
	CSSMERR_DL_DATASTORE_ALREADY_EXISTS        CssmDlBaseDlError = -2147413736
	CSSMERR_DL_DATASTORE_DOESNOT_EXIST         CssmDlBaseDlError = -2147413737
	CSSMERR_DL_DATASTORE_IS_OPEN               CssmDlBaseDlError = -2147413734
	CSSMERR_DL_DB_LOCKED                       CssmDlBaseDlError = -2147413735
	CSSMERR_DL_ENDOFDATA                       CssmDlBaseDlError = -2147413715
	CSSMERR_DL_FIELD_SPECIFIED_MULTIPLE        CssmDlBaseDlError = -2147413742
	CSSMERR_DL_INCOMPATIBLE_FIELD_FORMAT       CssmDlBaseDlError = -2147413741
	CSSMERR_DL_INVALID_ACCESS_REQUEST          CssmDlBaseDlError = -2147413724
	CSSMERR_DL_INVALID_DB_LOCATION             CssmDlBaseDlError = -2147413725
	CSSMERR_DL_INVALID_DB_NAME                 CssmDlBaseDlError = -2147413738
	CSSMERR_DL_INVALID_FIELD_NAME              CssmDlBaseDlError = -2147413750
	CSSMERR_DL_INVALID_INDEX_INFO              CssmDlBaseDlError = -2147413723
	CSSMERR_DL_INVALID_MODIFY_MODE             CssmDlBaseDlError = -2147413718
	CSSMERR_DL_INVALID_NEW_OWNER               CssmDlBaseDlError = -2147413721
	CSSMERR_DL_INVALID_OPEN_PARAMETERS         CssmDlBaseDlError = -2147413717
	CSSMERR_DL_INVALID_PARSING_MODULE          CssmDlBaseDlError = -2147413740
	CSSMERR_DL_INVALID_QUERY                   CssmDlBaseDlError = -2147413714
	CSSMERR_DL_INVALID_RECORDTYPE              CssmDlBaseDlError = -2147413751
	CSSMERR_DL_INVALID_RECORD_INDEX            CssmDlBaseDlError = -2147413752
	CSSMERR_DL_INVALID_RECORD_UID              CssmDlBaseDlError = -2147413720
	CSSMERR_DL_INVALID_RESULTS_HANDLE          CssmDlBaseDlError = -2147413726
	CSSMERR_DL_INVALID_SELECTION_TAG           CssmDlBaseDlError = -2147413722
	CSSMERR_DL_INVALID_UNIQUE_INDEX_DATA       CssmDlBaseDlError = -2147413719
	CSSMERR_DL_INVALID_VALUE                   CssmDlBaseDlError = -2147413713
	CSSMERR_DL_MISSING_VALUE                   CssmDlBaseDlError = -2147413732
	CSSMERR_DL_MULTIPLE_VALUES_UNSUPPORTED     CssmDlBaseDlError = -2147413712
	CSSMERR_DL_RECORD_MODIFIED                 CssmDlBaseDlError = -2147413716
	CSSMERR_DL_RECORD_NOT_FOUND                CssmDlBaseDlError = -2147413733
	CSSMERR_DL_STALE_UNIQUE_RECORD             CssmDlBaseDlError = -2147413711
	CSSMERR_DL_UNSUPPORTED_FIELD_FORMAT        CssmDlBaseDlError = -2147413749
	CSSMERR_DL_UNSUPPORTED_INDEX_INFO          CssmDlBaseDlError = -2147413748
	CSSMERR_DL_UNSUPPORTED_LOCALITY            CssmDlBaseDlError = -2147413747
	CSSMERR_DL_UNSUPPORTED_NUM_ATTRIBUTES      CssmDlBaseDlError = -2147413746
	CSSMERR_DL_UNSUPPORTED_NUM_INDEXES         CssmDlBaseDlError = -2147413745
	CSSMERR_DL_UNSUPPORTED_NUM_RECORDTYPES     CssmDlBaseDlError = -2147413744
	CSSMERR_DL_UNSUPPORTED_NUM_SELECTION_PREDS CssmDlBaseDlError = -2147413729
	CSSMERR_DL_UNSUPPORTED_OPERATOR            CssmDlBaseDlError = -2147413727
	CSSMERR_DL_UNSUPPORTED_QUERY               CssmDlBaseDlError = -2147413731
	CSSMERR_DL_UNSUPPORTED_QUERY_LIMITS        CssmDlBaseDlError = -2147413730
	CSSMERR_DL_UNSUPPORTED_RECORDTYPE          CssmDlBaseDlError = -2147413743
	CSSM_DL_BASE_DL_ERROR                      CssmDlBaseDlError = -2147413760
)

func (e CssmDlBaseDlError) String() string {
	switch e {
	case CSSMERR_DL_DATABASE_CORRUPT:
		return "CSSMERR_DL_DATABASE_CORRUPT"
	case CSSMERR_DL_DATASTORE_ALREADY_EXISTS:
		return "CSSMERR_DL_DATASTORE_ALREADY_EXISTS"
	case CSSMERR_DL_DATASTORE_DOESNOT_EXIST:
		return "CSSMERR_DL_DATASTORE_DOESNOT_EXIST"
	case CSSMERR_DL_DATASTORE_IS_OPEN:
		return "CSSMERR_DL_DATASTORE_IS_OPEN"
	case CSSMERR_DL_DB_LOCKED:
		return "CSSMERR_DL_DB_LOCKED"
	case CSSMERR_DL_ENDOFDATA:
		return "CSSMERR_DL_ENDOFDATA"
	case CSSMERR_DL_FIELD_SPECIFIED_MULTIPLE:
		return "CSSMERR_DL_FIELD_SPECIFIED_MULTIPLE"
	case CSSMERR_DL_INCOMPATIBLE_FIELD_FORMAT:
		return "CSSMERR_DL_INCOMPATIBLE_FIELD_FORMAT"
	case CSSMERR_DL_INVALID_ACCESS_REQUEST:
		return "CSSMERR_DL_INVALID_ACCESS_REQUEST"
	case CSSMERR_DL_INVALID_DB_LOCATION:
		return "CSSMERR_DL_INVALID_DB_LOCATION"
	case CSSMERR_DL_INVALID_DB_NAME:
		return "CSSMERR_DL_INVALID_DB_NAME"
	case CSSMERR_DL_INVALID_FIELD_NAME:
		return "CSSMERR_DL_INVALID_FIELD_NAME"
	case CSSMERR_DL_INVALID_INDEX_INFO:
		return "CSSMERR_DL_INVALID_INDEX_INFO"
	case CSSMERR_DL_INVALID_MODIFY_MODE:
		return "CSSMERR_DL_INVALID_MODIFY_MODE"
	case CSSMERR_DL_INVALID_NEW_OWNER:
		return "CSSMERR_DL_INVALID_NEW_OWNER"
	case CSSMERR_DL_INVALID_OPEN_PARAMETERS:
		return "CSSMERR_DL_INVALID_OPEN_PARAMETERS"
	case CSSMERR_DL_INVALID_PARSING_MODULE:
		return "CSSMERR_DL_INVALID_PARSING_MODULE"
	case CSSMERR_DL_INVALID_QUERY:
		return "CSSMERR_DL_INVALID_QUERY"
	case CSSMERR_DL_INVALID_RECORDTYPE:
		return "CSSMERR_DL_INVALID_RECORDTYPE"
	case CSSMERR_DL_INVALID_RECORD_INDEX:
		return "CSSMERR_DL_INVALID_RECORD_INDEX"
	case CSSMERR_DL_INVALID_RECORD_UID:
		return "CSSMERR_DL_INVALID_RECORD_UID"
	case CSSMERR_DL_INVALID_RESULTS_HANDLE:
		return "CSSMERR_DL_INVALID_RESULTS_HANDLE"
	case CSSMERR_DL_INVALID_SELECTION_TAG:
		return "CSSMERR_DL_INVALID_SELECTION_TAG"
	case CSSMERR_DL_INVALID_UNIQUE_INDEX_DATA:
		return "CSSMERR_DL_INVALID_UNIQUE_INDEX_DATA"
	case CSSMERR_DL_INVALID_VALUE:
		return "CSSMERR_DL_INVALID_VALUE"
	case CSSMERR_DL_MISSING_VALUE:
		return "CSSMERR_DL_MISSING_VALUE"
	case CSSMERR_DL_MULTIPLE_VALUES_UNSUPPORTED:
		return "CSSMERR_DL_MULTIPLE_VALUES_UNSUPPORTED"
	case CSSMERR_DL_RECORD_MODIFIED:
		return "CSSMERR_DL_RECORD_MODIFIED"
	case CSSMERR_DL_RECORD_NOT_FOUND:
		return "CSSMERR_DL_RECORD_NOT_FOUND"
	case CSSMERR_DL_STALE_UNIQUE_RECORD:
		return "CSSMERR_DL_STALE_UNIQUE_RECORD"
	case CSSMERR_DL_UNSUPPORTED_FIELD_FORMAT:
		return "CSSMERR_DL_UNSUPPORTED_FIELD_FORMAT"
	case CSSMERR_DL_UNSUPPORTED_INDEX_INFO:
		return "CSSMERR_DL_UNSUPPORTED_INDEX_INFO"
	case CSSMERR_DL_UNSUPPORTED_LOCALITY:
		return "CSSMERR_DL_UNSUPPORTED_LOCALITY"
	case CSSMERR_DL_UNSUPPORTED_NUM_ATTRIBUTES:
		return "CSSMERR_DL_UNSUPPORTED_NUM_ATTRIBUTES"
	case CSSMERR_DL_UNSUPPORTED_NUM_INDEXES:
		return "CSSMERR_DL_UNSUPPORTED_NUM_INDEXES"
	case CSSMERR_DL_UNSUPPORTED_NUM_RECORDTYPES:
		return "CSSMERR_DL_UNSUPPORTED_NUM_RECORDTYPES"
	case CSSMERR_DL_UNSUPPORTED_NUM_SELECTION_PREDS:
		return "CSSMERR_DL_UNSUPPORTED_NUM_SELECTION_PREDS"
	case CSSMERR_DL_UNSUPPORTED_OPERATOR:
		return "CSSMERR_DL_UNSUPPORTED_OPERATOR"
	case CSSMERR_DL_UNSUPPORTED_QUERY:
		return "CSSMERR_DL_UNSUPPORTED_QUERY"
	case CSSMERR_DL_UNSUPPORTED_QUERY_LIMITS:
		return "CSSMERR_DL_UNSUPPORTED_QUERY_LIMITS"
	case CSSMERR_DL_UNSUPPORTED_RECORDTYPE:
		return "CSSMERR_DL_UNSUPPORTED_RECORDTYPE"
	case CSSM_DL_BASE_DL_ERROR:
		return "CSSM_DL_BASE_DL_ERROR"
	default:
		return fmt.Sprintf("CssmDlBaseDlError(%d)", e)
	}
}

type CssmDlDbRecord uint32

const (
	CSSM_DL_DB_RECORD_APPLESHARE_PASSWORD CssmDlDbRecord = 2147483650
	CSSM_DL_DB_RECORD_EXTENDED_ATTRIBUTE  CssmDlDbRecord = 2147487748
	CSSM_DL_DB_RECORD_GENERIC_PASSWORD    CssmDlDbRecord = 2147483648
	CSSM_DL_DB_RECORD_INTERNET_PASSWORD   CssmDlDbRecord = 2147483649
	CSSM_DL_DB_RECORD_METADATA            CssmDlDbRecord = 2147516416
	CSSM_DL_DB_RECORD_UNLOCK_REFERRAL     CssmDlDbRecord = 2147487747
	CSSM_DL_DB_RECORD_USER_TRUST          CssmDlDbRecord = 2147487745
	CSSM_DL_DB_RECORD_X509_CERTIFICATE    CssmDlDbRecord = 2147487744
	CSSM_DL_DB_RECORD_X509_CRL            CssmDlDbRecord = 2147487746
)

func (e CssmDlDbRecord) String() string {
	switch e {
	case CSSM_DL_DB_RECORD_APPLESHARE_PASSWORD:
		return "CSSM_DL_DB_RECORD_APPLESHARE_PASSWORD"
	case CSSM_DL_DB_RECORD_EXTENDED_ATTRIBUTE:
		return "CSSM_DL_DB_RECORD_EXTENDED_ATTRIBUTE"
	case CSSM_DL_DB_RECORD_GENERIC_PASSWORD:
		return "CSSM_DL_DB_RECORD_GENERIC_PASSWORD"
	case CSSM_DL_DB_RECORD_INTERNET_PASSWORD:
		return "CSSM_DL_DB_RECORD_INTERNET_PASSWORD"
	case CSSM_DL_DB_RECORD_METADATA:
		return "CSSM_DL_DB_RECORD_METADATA"
	case CSSM_DL_DB_RECORD_UNLOCK_REFERRAL:
		return "CSSM_DL_DB_RECORD_UNLOCK_REFERRAL"
	case CSSM_DL_DB_RECORD_USER_TRUST:
		return "CSSM_DL_DB_RECORD_USER_TRUST"
	case CSSM_DL_DB_RECORD_X509_CERTIFICATE:
		return "CSSM_DL_DB_RECORD_X509_CERTIFICATE"
	case CSSM_DL_DB_RECORD_X509_CRL:
		return "CSSM_DL_DB_RECORD_X509_CRL"
	default:
		return fmt.Sprintf("CssmDlDbRecord(%d)", e)
	}
}

type CssmElapsedTime int32

const (
	CSSM_ELAPSED_TIME_COMPLETE CssmElapsedTime = -2
	CSSM_ELAPSED_TIME_UNKNOWN  CssmElapsedTime = -1
)

func (e CssmElapsedTime) String() string {
	switch e {
	case CSSM_ELAPSED_TIME_COMPLETE:
		return "CSSM_ELAPSED_TIME_COMPLETE"
	case CSSM_ELAPSED_TIME_UNKNOWN:
		return "CSSM_ELAPSED_TIME_UNKNOWN"
	default:
		return fmt.Sprintf("CssmElapsedTime(%d)", e)
	}
}

type CssmErrcodeInternalError uint32

const (
	CSSM_ERRCODE_FUNCTION_FAILED               CssmErrcodeInternalError = 0xa
	CSSM_ERRCODE_FUNCTION_NOT_IMPLEMENTED      CssmErrcodeInternalError = 0x7
	CSSM_ERRCODE_INTERNAL_ERROR                CssmErrcodeInternalError = 0x1
	CSSM_ERRCODE_INVALID_GUID                  CssmErrcodeInternalError = 0xc
	CSSM_ERRCODE_INVALID_INPUT_POINTER         CssmErrcodeInternalError = 0x5
	CSSM_ERRCODE_INVALID_OUTPUT_POINTER        CssmErrcodeInternalError = 0x6
	CSSM_ERRCODE_INVALID_POINTER               CssmErrcodeInternalError = 0x4
	CSSM_ERRCODE_MDS_ERROR                     CssmErrcodeInternalError = 0x3
	CSSM_ERRCODE_MEMORY_ERROR                  CssmErrcodeInternalError = 0x2
	CSSM_ERRCODE_MODULE_MANIFEST_VERIFY_FAILED CssmErrcodeInternalError = 0xb
	CSSM_ERRCODE_OS_ACCESS_DENIED              CssmErrcodeInternalError = 0x9
	CSSM_ERRCODE_SELF_CHECK_FAILED             CssmErrcodeInternalError = 0x8
)

func (e CssmErrcodeInternalError) String() string {
	switch e {
	case CSSM_ERRCODE_FUNCTION_FAILED:
		return "CSSM_ERRCODE_FUNCTION_FAILED"
	case CSSM_ERRCODE_FUNCTION_NOT_IMPLEMENTED:
		return "CSSM_ERRCODE_FUNCTION_NOT_IMPLEMENTED"
	case CSSM_ERRCODE_INTERNAL_ERROR:
		return "CSSM_ERRCODE_INTERNAL_ERROR"
	case CSSM_ERRCODE_INVALID_GUID:
		return "CSSM_ERRCODE_INVALID_GUID"
	case CSSM_ERRCODE_INVALID_INPUT_POINTER:
		return "CSSM_ERRCODE_INVALID_INPUT_POINTER"
	case CSSM_ERRCODE_INVALID_OUTPUT_POINTER:
		return "CSSM_ERRCODE_INVALID_OUTPUT_POINTER"
	case CSSM_ERRCODE_INVALID_POINTER:
		return "CSSM_ERRCODE_INVALID_POINTER"
	case CSSM_ERRCODE_MDS_ERROR:
		return "CSSM_ERRCODE_MDS_ERROR"
	case CSSM_ERRCODE_MEMORY_ERROR:
		return "CSSM_ERRCODE_MEMORY_ERROR"
	case CSSM_ERRCODE_MODULE_MANIFEST_VERIFY_FAILED:
		return "CSSM_ERRCODE_MODULE_MANIFEST_VERIFY_FAILED"
	case CSSM_ERRCODE_OS_ACCESS_DENIED:
		return "CSSM_ERRCODE_OS_ACCESS_DENIED"
	case CSSM_ERRCODE_SELF_CHECK_FAILED:
		return "CSSM_ERRCODE_SELF_CHECK_FAILED"
	default:
		return fmt.Sprintf("CssmErrcodeInternalError(%d)", e)
	}
}

type CssmErrcodeInvalidContextHandle uint32

const (
	CSSM_ERRCODE_CRL_ALREADY_SIGNED        CssmErrcodeInvalidContextHandle = 0x47
	CSSM_ERRCODE_INCOMPATIBLE_VERSION      CssmErrcodeInvalidContextHandle = 0x41
	CSSM_ERRCODE_INVALID_AC_HANDLE         CssmErrcodeInvalidContextHandle = 0x55
	CSSM_ERRCODE_INVALID_CERTGROUP_POINTER CssmErrcodeInvalidContextHandle = 0x42
	CSSM_ERRCODE_INVALID_CERT_POINTER      CssmErrcodeInvalidContextHandle = 0x43
	CSSM_ERRCODE_INVALID_CL_HANDLE         CssmErrcodeInvalidContextHandle = 0x52
	CSSM_ERRCODE_INVALID_CONTEXT_HANDLE    CssmErrcodeInvalidContextHandle = 0x40
	CSSM_ERRCODE_INVALID_CRL_POINTER       CssmErrcodeInvalidContextHandle = 0x44
	CSSM_ERRCODE_INVALID_CRYPTO_DATA       CssmErrcodeInvalidContextHandle = 0x58
	CSSM_ERRCODE_INVALID_CSP_HANDLE        CssmErrcodeInvalidContextHandle = 0x50
	CSSM_ERRCODE_INVALID_DATA              CssmErrcodeInvalidContextHandle = 0x46
	CSSM_ERRCODE_INVALID_DB_HANDLE         CssmErrcodeInvalidContextHandle = 0x4a
	CSSM_ERRCODE_INVALID_DB_LIST           CssmErrcodeInvalidContextHandle = 0x4c
	CSSM_ERRCODE_INVALID_DB_LIST_POINTER   CssmErrcodeInvalidContextHandle = 0x4d
	CSSM_ERRCODE_INVALID_DL_HANDLE         CssmErrcodeInvalidContextHandle = 0x51
	CSSM_ERRCODE_INVALID_FIELD_POINTER     CssmErrcodeInvalidContextHandle = 0x45
	CSSM_ERRCODE_INVALID_KR_HANDLE         CssmErrcodeInvalidContextHandle = 0x54
	CSSM_ERRCODE_INVALID_NETWORK_ADDR      CssmErrcodeInvalidContextHandle = 0x57
	CSSM_ERRCODE_INVALID_NUMBER_OF_FIELDS  CssmErrcodeInvalidContextHandle = 0x48
	CSSM_ERRCODE_INVALID_PASSTHROUGH_ID    CssmErrcodeInvalidContextHandle = 0x56
	CSSM_ERRCODE_INVALID_TP_HANDLE         CssmErrcodeInvalidContextHandle = 0x53
	CSSM_ERRCODE_PRIVILEGE_NOT_GRANTED     CssmErrcodeInvalidContextHandle = 0x4b
	CSSM_ERRCODE_UNKNOWN_FORMAT            CssmErrcodeInvalidContextHandle = 0x4e
	CSSM_ERRCODE_UNKNOWN_TAG               CssmErrcodeInvalidContextHandle = 0x4f
	CSSM_ERRCODE_VERIFICATION_FAILURE      CssmErrcodeInvalidContextHandle = 0x49
)

func (e CssmErrcodeInvalidContextHandle) String() string {
	switch e {
	case CSSM_ERRCODE_CRL_ALREADY_SIGNED:
		return "CSSM_ERRCODE_CRL_ALREADY_SIGNED"
	case CSSM_ERRCODE_INCOMPATIBLE_VERSION:
		return "CSSM_ERRCODE_INCOMPATIBLE_VERSION"
	case CSSM_ERRCODE_INVALID_AC_HANDLE:
		return "CSSM_ERRCODE_INVALID_AC_HANDLE"
	case CSSM_ERRCODE_INVALID_CERTGROUP_POINTER:
		return "CSSM_ERRCODE_INVALID_CERTGROUP_POINTER"
	case CSSM_ERRCODE_INVALID_CERT_POINTER:
		return "CSSM_ERRCODE_INVALID_CERT_POINTER"
	case CSSM_ERRCODE_INVALID_CL_HANDLE:
		return "CSSM_ERRCODE_INVALID_CL_HANDLE"
	case CSSM_ERRCODE_INVALID_CONTEXT_HANDLE:
		return "CSSM_ERRCODE_INVALID_CONTEXT_HANDLE"
	case CSSM_ERRCODE_INVALID_CRL_POINTER:
		return "CSSM_ERRCODE_INVALID_CRL_POINTER"
	case CSSM_ERRCODE_INVALID_CRYPTO_DATA:
		return "CSSM_ERRCODE_INVALID_CRYPTO_DATA"
	case CSSM_ERRCODE_INVALID_CSP_HANDLE:
		return "CSSM_ERRCODE_INVALID_CSP_HANDLE"
	case CSSM_ERRCODE_INVALID_DATA:
		return "CSSM_ERRCODE_INVALID_DATA"
	case CSSM_ERRCODE_INVALID_DB_HANDLE:
		return "CSSM_ERRCODE_INVALID_DB_HANDLE"
	case CSSM_ERRCODE_INVALID_DB_LIST:
		return "CSSM_ERRCODE_INVALID_DB_LIST"
	case CSSM_ERRCODE_INVALID_DB_LIST_POINTER:
		return "CSSM_ERRCODE_INVALID_DB_LIST_POINTER"
	case CSSM_ERRCODE_INVALID_DL_HANDLE:
		return "CSSM_ERRCODE_INVALID_DL_HANDLE"
	case CSSM_ERRCODE_INVALID_FIELD_POINTER:
		return "CSSM_ERRCODE_INVALID_FIELD_POINTER"
	case CSSM_ERRCODE_INVALID_KR_HANDLE:
		return "CSSM_ERRCODE_INVALID_KR_HANDLE"
	case CSSM_ERRCODE_INVALID_NETWORK_ADDR:
		return "CSSM_ERRCODE_INVALID_NETWORK_ADDR"
	case CSSM_ERRCODE_INVALID_NUMBER_OF_FIELDS:
		return "CSSM_ERRCODE_INVALID_NUMBER_OF_FIELDS"
	case CSSM_ERRCODE_INVALID_PASSTHROUGH_ID:
		return "CSSM_ERRCODE_INVALID_PASSTHROUGH_ID"
	case CSSM_ERRCODE_INVALID_TP_HANDLE:
		return "CSSM_ERRCODE_INVALID_TP_HANDLE"
	case CSSM_ERRCODE_PRIVILEGE_NOT_GRANTED:
		return "CSSM_ERRCODE_PRIVILEGE_NOT_GRANTED"
	case CSSM_ERRCODE_UNKNOWN_FORMAT:
		return "CSSM_ERRCODE_UNKNOWN_FORMAT"
	case CSSM_ERRCODE_UNKNOWN_TAG:
		return "CSSM_ERRCODE_UNKNOWN_TAG"
	case CSSM_ERRCODE_VERIFICATION_FAILURE:
		return "CSSM_ERRCODE_VERIFICATION_FAILURE"
	default:
		return fmt.Sprintf("CssmErrcodeInvalidContextHandle(%d)", e)
	}
}

type CssmErrcodeOperationAuthDenied uint32

const (
	CSSM_ERRCODE_ACL_ADD_FAILED                 CssmErrcodeOperationAuthDenied = 0x36
	CSSM_ERRCODE_ACL_BASE_CERTS_NOT_SUPPORTED   CssmErrcodeOperationAuthDenied = 0x27
	CSSM_ERRCODE_ACL_CHALLENGE_CALLBACK_FAILED  CssmErrcodeOperationAuthDenied = 0x2d
	CSSM_ERRCODE_ACL_CHANGE_FAILED              CssmErrcodeOperationAuthDenied = 0x31
	CSSM_ERRCODE_ACL_DELETE_FAILED              CssmErrcodeOperationAuthDenied = 0x34
	CSSM_ERRCODE_ACL_ENTRY_TAG_NOT_FOUND        CssmErrcodeOperationAuthDenied = 0x2f
	CSSM_ERRCODE_ACL_REPLACE_FAILED             CssmErrcodeOperationAuthDenied = 0x35
	CSSM_ERRCODE_ACL_SUBJECT_TYPE_NOT_SUPPORTED CssmErrcodeOperationAuthDenied = 0x2b
	CSSM_ERRCODE_INVALID_ACCESS_CREDENTIALS     CssmErrcodeOperationAuthDenied = 0x25
	CSSM_ERRCODE_INVALID_ACL_BASE_CERTS         CssmErrcodeOperationAuthDenied = 0x26
	CSSM_ERRCODE_INVALID_ACL_CHALLENGE_CALLBACK CssmErrcodeOperationAuthDenied = 0x2c
	CSSM_ERRCODE_INVALID_ACL_EDIT_MODE          CssmErrcodeOperationAuthDenied = 0x30
	CSSM_ERRCODE_INVALID_ACL_ENTRY_TAG          CssmErrcodeOperationAuthDenied = 0x2e
	CSSM_ERRCODE_INVALID_ACL_SUBJECT_VALUE      CssmErrcodeOperationAuthDenied = 0x2a
	CSSM_ERRCODE_INVALID_NEW_ACL_ENTRY          CssmErrcodeOperationAuthDenied = 0x32
	CSSM_ERRCODE_INVALID_NEW_ACL_OWNER          CssmErrcodeOperationAuthDenied = 0x33
	CSSM_ERRCODE_INVALID_SAMPLE_VALUE           CssmErrcodeOperationAuthDenied = 0x28
	CSSM_ERRCODE_OBJECT_ACL_NOT_SUPPORTED       CssmErrcodeOperationAuthDenied = 0x23
	CSSM_ERRCODE_OBJECT_ACL_REQUIRED            CssmErrcodeOperationAuthDenied = 0x24
	CSSM_ERRCODE_OBJECT_MANIP_AUTH_DENIED       CssmErrcodeOperationAuthDenied = 0x22
	CSSM_ERRCODE_OBJECT_USE_AUTH_DENIED         CssmErrcodeOperationAuthDenied = 0x21
	CSSM_ERRCODE_OPERATION_AUTH_DENIED          CssmErrcodeOperationAuthDenied = 0x20
	CSSM_ERRCODE_SAMPLE_VALUE_NOT_SUPPORTED     CssmErrcodeOperationAuthDenied = 0x29
)

func (e CssmErrcodeOperationAuthDenied) String() string {
	switch e {
	case CSSM_ERRCODE_ACL_ADD_FAILED:
		return "CSSM_ERRCODE_ACL_ADD_FAILED"
	case CSSM_ERRCODE_ACL_BASE_CERTS_NOT_SUPPORTED:
		return "CSSM_ERRCODE_ACL_BASE_CERTS_NOT_SUPPORTED"
	case CSSM_ERRCODE_ACL_CHALLENGE_CALLBACK_FAILED:
		return "CSSM_ERRCODE_ACL_CHALLENGE_CALLBACK_FAILED"
	case CSSM_ERRCODE_ACL_CHANGE_FAILED:
		return "CSSM_ERRCODE_ACL_CHANGE_FAILED"
	case CSSM_ERRCODE_ACL_DELETE_FAILED:
		return "CSSM_ERRCODE_ACL_DELETE_FAILED"
	case CSSM_ERRCODE_ACL_ENTRY_TAG_NOT_FOUND:
		return "CSSM_ERRCODE_ACL_ENTRY_TAG_NOT_FOUND"
	case CSSM_ERRCODE_ACL_REPLACE_FAILED:
		return "CSSM_ERRCODE_ACL_REPLACE_FAILED"
	case CSSM_ERRCODE_ACL_SUBJECT_TYPE_NOT_SUPPORTED:
		return "CSSM_ERRCODE_ACL_SUBJECT_TYPE_NOT_SUPPORTED"
	case CSSM_ERRCODE_INVALID_ACCESS_CREDENTIALS:
		return "CSSM_ERRCODE_INVALID_ACCESS_CREDENTIALS"
	case CSSM_ERRCODE_INVALID_ACL_BASE_CERTS:
		return "CSSM_ERRCODE_INVALID_ACL_BASE_CERTS"
	case CSSM_ERRCODE_INVALID_ACL_CHALLENGE_CALLBACK:
		return "CSSM_ERRCODE_INVALID_ACL_CHALLENGE_CALLBACK"
	case CSSM_ERRCODE_INVALID_ACL_EDIT_MODE:
		return "CSSM_ERRCODE_INVALID_ACL_EDIT_MODE"
	case CSSM_ERRCODE_INVALID_ACL_ENTRY_TAG:
		return "CSSM_ERRCODE_INVALID_ACL_ENTRY_TAG"
	case CSSM_ERRCODE_INVALID_ACL_SUBJECT_VALUE:
		return "CSSM_ERRCODE_INVALID_ACL_SUBJECT_VALUE"
	case CSSM_ERRCODE_INVALID_NEW_ACL_ENTRY:
		return "CSSM_ERRCODE_INVALID_NEW_ACL_ENTRY"
	case CSSM_ERRCODE_INVALID_NEW_ACL_OWNER:
		return "CSSM_ERRCODE_INVALID_NEW_ACL_OWNER"
	case CSSM_ERRCODE_INVALID_SAMPLE_VALUE:
		return "CSSM_ERRCODE_INVALID_SAMPLE_VALUE"
	case CSSM_ERRCODE_OBJECT_ACL_NOT_SUPPORTED:
		return "CSSM_ERRCODE_OBJECT_ACL_NOT_SUPPORTED"
	case CSSM_ERRCODE_OBJECT_ACL_REQUIRED:
		return "CSSM_ERRCODE_OBJECT_ACL_REQUIRED"
	case CSSM_ERRCODE_OBJECT_MANIP_AUTH_DENIED:
		return "CSSM_ERRCODE_OBJECT_MANIP_AUTH_DENIED"
	case CSSM_ERRCODE_OBJECT_USE_AUTH_DENIED:
		return "CSSM_ERRCODE_OBJECT_USE_AUTH_DENIED"
	case CSSM_ERRCODE_OPERATION_AUTH_DENIED:
		return "CSSM_ERRCODE_OPERATION_AUTH_DENIED"
	case CSSM_ERRCODE_SAMPLE_VALUE_NOT_SUPPORTED:
		return "CSSM_ERRCODE_SAMPLE_VALUE_NOT_SUPPORTED"
	default:
		return fmt.Sprintf("CssmErrcodeOperationAuthDenied(%d)", e)
	}
}

type CssmEstimatedTime int32

const (
	CSSM_ESTIMATED_TIME_UNKNOWN CssmEstimatedTime = -1
)

func (e CssmEstimatedTime) String() string {
	switch e {
	case CSSM_ESTIMATED_TIME_UNKNOWN:
		return "CSSM_ESTIMATED_TIME_UNKNOWN"
	default:
		return fmt.Sprintf("CssmEstimatedTime(%d)", e)
	}
}

type CssmEvidenceForm uint32

const (
	CSSM_EVIDENCE_FORM_CERT          CssmEvidenceForm = 0x1
	CSSM_EVIDENCE_FORM_CERT_ID       CssmEvidenceForm = 0x3
	CSSM_EVIDENCE_FORM_CRL           CssmEvidenceForm = 0x2
	CSSM_EVIDENCE_FORM_CRL_ID        CssmEvidenceForm = 0x4
	CSSM_EVIDENCE_FORM_CRL_NEXTTIME  CssmEvidenceForm = 0x7
	CSSM_EVIDENCE_FORM_CRL_THISTIME  CssmEvidenceForm = 0x6
	CSSM_EVIDENCE_FORM_POLICYINFO    CssmEvidenceForm = 0x8
	CSSM_EVIDENCE_FORM_TUPLEGROUP    CssmEvidenceForm = 0x9
	CSSM_EVIDENCE_FORM_UNSPECIFIC    CssmEvidenceForm = 0
	CSSM_EVIDENCE_FORM_VERIFIER_TIME CssmEvidenceForm = 0x5
)

func (e CssmEvidenceForm) String() string {
	switch e {
	case CSSM_EVIDENCE_FORM_CERT:
		return "CSSM_EVIDENCE_FORM_CERT"
	case CSSM_EVIDENCE_FORM_CERT_ID:
		return "CSSM_EVIDENCE_FORM_CERT_ID"
	case CSSM_EVIDENCE_FORM_CRL:
		return "CSSM_EVIDENCE_FORM_CRL"
	case CSSM_EVIDENCE_FORM_CRL_ID:
		return "CSSM_EVIDENCE_FORM_CRL_ID"
	case CSSM_EVIDENCE_FORM_CRL_NEXTTIME:
		return "CSSM_EVIDENCE_FORM_CRL_NEXTTIME"
	case CSSM_EVIDENCE_FORM_CRL_THISTIME:
		return "CSSM_EVIDENCE_FORM_CRL_THISTIME"
	case CSSM_EVIDENCE_FORM_POLICYINFO:
		return "CSSM_EVIDENCE_FORM_POLICYINFO"
	case CSSM_EVIDENCE_FORM_TUPLEGROUP:
		return "CSSM_EVIDENCE_FORM_TUPLEGROUP"
	case CSSM_EVIDENCE_FORM_UNSPECIFIC:
		return "CSSM_EVIDENCE_FORM_UNSPECIFIC"
	case CSSM_EVIDENCE_FORM_VERIFIER_TIME:
		return "CSSM_EVIDENCE_FORM_VERIFIER_TIME"
	default:
		return fmt.Sprintf("CssmEvidenceForm(%d)", e)
	}
}

type CssmEvidenceFormApple uint32

const (
	CSSM_EVIDENCE_FORM_APPLE_CERTGROUP CssmEvidenceFormApple = 2147483649
	CSSM_EVIDENCE_FORM_APPLE_CERT_INFO CssmEvidenceFormApple = 2147483650
	CSSM_EVIDENCE_FORM_APPLE_HEADER    CssmEvidenceFormApple = 2147483648
)

func (e CssmEvidenceFormApple) String() string {
	switch e {
	case CSSM_EVIDENCE_FORM_APPLE_CERTGROUP:
		return "CSSM_EVIDENCE_FORM_APPLE_CERTGROUP"
	case CSSM_EVIDENCE_FORM_APPLE_CERT_INFO:
		return "CSSM_EVIDENCE_FORM_APPLE_CERT_INFO"
	case CSSM_EVIDENCE_FORM_APPLE_HEADER:
		return "CSSM_EVIDENCE_FORM_APPLE_HEADER"
	default:
		return fmt.Sprintf("CssmEvidenceFormApple(%d)", e)
	}
}

type CssmFalse uint32

const (
	CSSM_FALSE CssmFalse = 0
	CSSM_TRUE  CssmFalse = 1
)

func (e CssmFalse) String() string {
	switch e {
	case CSSM_FALSE:
		return "CSSM_FALSE"
	case CSSM_TRUE:
		return "CSSM_TRUE"
	default:
		return fmt.Sprintf("CssmFalse(%d)", e)
	}
}

type CssmFeeCurveType uint32

const (
	CSSM_FEE_CURVE_TYPE_ANSI_X9_62  CssmFeeCurveType = 3
	CSSM_FEE_CURVE_TYPE_DEFAULT     CssmFeeCurveType = 0
	CSSM_FEE_CURVE_TYPE_MONTGOMERY  CssmFeeCurveType = 1
	CSSM_FEE_CURVE_TYPE_WEIERSTRASS CssmFeeCurveType = 2
)

func (e CssmFeeCurveType) String() string {
	switch e {
	case CSSM_FEE_CURVE_TYPE_ANSI_X9_62:
		return "CSSM_FEE_CURVE_TYPE_ANSI_X9_62"
	case CSSM_FEE_CURVE_TYPE_DEFAULT:
		return "CSSM_FEE_CURVE_TYPE_DEFAULT"
	case CSSM_FEE_CURVE_TYPE_MONTGOMERY:
		return "CSSM_FEE_CURVE_TYPE_MONTGOMERY"
	case CSSM_FEE_CURVE_TYPE_WEIERSTRASS:
		return "CSSM_FEE_CURVE_TYPE_WEIERSTRASS"
	default:
		return fmt.Sprintf("CssmFeeCurveType(%d)", e)
	}
}

type CssmFeePrimeType uint32

const (
	CSSM_FEE_PRIME_TYPE_DEFAULT  CssmFeePrimeType = 0
	CSSM_FEE_PRIME_TYPE_FEE      CssmFeePrimeType = 2
	CSSM_FEE_PRIME_TYPE_GENERAL  CssmFeePrimeType = 3
	CSSM_FEE_PRIME_TYPE_MERSENNE CssmFeePrimeType = 1
)

func (e CssmFeePrimeType) String() string {
	switch e {
	case CSSM_FEE_PRIME_TYPE_DEFAULT:
		return "CSSM_FEE_PRIME_TYPE_DEFAULT"
	case CSSM_FEE_PRIME_TYPE_FEE:
		return "CSSM_FEE_PRIME_TYPE_FEE"
	case CSSM_FEE_PRIME_TYPE_GENERAL:
		return "CSSM_FEE_PRIME_TYPE_GENERAL"
	case CSSM_FEE_PRIME_TYPE_MERSENNE:
		return "CSSM_FEE_PRIME_TYPE_MERSENNE"
	default:
		return fmt.Sprintf("CssmFeePrimeType(%d)", e)
	}
}

type CssmFieldvalueComplexData uint32

const (
	CSSM_FIELDVALUE_COMPLEX_DATA_TYPE CssmFieldvalueComplexData = 0xffffffff
)

func (e CssmFieldvalueComplexData) String() string {
	switch e {
	case CSSM_FIELDVALUE_COMPLEX_DATA_TYPE:
		return "CSSM_FIELDVALUE_COMPLEX_DATA_TYPE"
	default:
		return fmt.Sprintf("CssmFieldvalueComplexData(%d)", e)
	}
}

type CssmHint uint32

const (
	CSSM_HINT_ADDRESS_APP CssmHint = 1
	CSSM_HINT_ADDRESS_SP  CssmHint = 2
	CSSM_HINT_NONE        CssmHint = 0
)

func (e CssmHint) String() string {
	switch e {
	case CSSM_HINT_ADDRESS_APP:
		return "CSSM_HINT_ADDRESS_APP"
	case CSSM_HINT_ADDRESS_SP:
		return "CSSM_HINT_ADDRESS_SP"
	case CSSM_HINT_NONE:
		return "CSSM_HINT_NONE"
	default:
		return fmt.Sprintf("CssmHint(%d)", e)
	}
}

type CssmInvalid uint32

const (
	CSSM_INVALID_HANDLE CssmInvalid = 0
)

func (e CssmInvalid) String() string {
	switch e {
	case CSSM_INVALID_HANDLE:
		return "CSSM_INVALID_HANDLE"
	default:
		return fmt.Sprintf("CssmInvalid(%d)", e)
	}
}

type CssmKeyHierarchy uint32

const (
	CSSM_KEY_HIERARCHY_EXPORT CssmKeyHierarchy = 2
	CSSM_KEY_HIERARCHY_INTEG  CssmKeyHierarchy = 1
	CSSM_KEY_HIERARCHY_NONE   CssmKeyHierarchy = 0
)

func (e CssmKeyHierarchy) String() string {
	switch e {
	case CSSM_KEY_HIERARCHY_EXPORT:
		return "CSSM_KEY_HIERARCHY_EXPORT"
	case CSSM_KEY_HIERARCHY_INTEG:
		return "CSSM_KEY_HIERARCHY_INTEG"
	case CSSM_KEY_HIERARCHY_NONE:
		return "CSSM_KEY_HIERARCHY_NONE"
	default:
		return fmt.Sprintf("CssmKeyHierarchy(%d)", e)
	}
}

type CssmKeyattr uint32

const (
	CSSM_KEYATTR_ALWAYS_SENSITIVE  CssmKeyattr = 0x10
	CSSM_KEYATTR_EXTRACTABLE       CssmKeyattr = 0x20
	CSSM_KEYATTR_MODIFIABLE        CssmKeyattr = 0x4
	CSSM_KEYATTR_NEVER_EXTRACTABLE CssmKeyattr = 0x40
	CSSM_KEYATTR_PERMANENT         CssmKeyattr = 0x1
	CSSM_KEYATTR_PRIVATE           CssmKeyattr = 0x2
	CSSM_KEYATTR_RETURN_DATA       CssmKeyattr = 0x10000000
	CSSM_KEYATTR_RETURN_DEFAULT    CssmKeyattr = 0
	CSSM_KEYATTR_RETURN_NONE       CssmKeyattr = 0x40000000
	CSSM_KEYATTR_RETURN_REF        CssmKeyattr = 0x20000000
	CSSM_KEYATTR_SENSITIVE         CssmKeyattr = 0x8
)

func (e CssmKeyattr) String() string {
	switch e {
	case CSSM_KEYATTR_ALWAYS_SENSITIVE:
		return "CSSM_KEYATTR_ALWAYS_SENSITIVE"
	case CSSM_KEYATTR_EXTRACTABLE:
		return "CSSM_KEYATTR_EXTRACTABLE"
	case CSSM_KEYATTR_MODIFIABLE:
		return "CSSM_KEYATTR_MODIFIABLE"
	case CSSM_KEYATTR_NEVER_EXTRACTABLE:
		return "CSSM_KEYATTR_NEVER_EXTRACTABLE"
	case CSSM_KEYATTR_PERMANENT:
		return "CSSM_KEYATTR_PERMANENT"
	case CSSM_KEYATTR_PRIVATE:
		return "CSSM_KEYATTR_PRIVATE"
	case CSSM_KEYATTR_RETURN_DATA:
		return "CSSM_KEYATTR_RETURN_DATA"
	case CSSM_KEYATTR_RETURN_DEFAULT:
		return "CSSM_KEYATTR_RETURN_DEFAULT"
	case CSSM_KEYATTR_RETURN_NONE:
		return "CSSM_KEYATTR_RETURN_NONE"
	case CSSM_KEYATTR_RETURN_REF:
		return "CSSM_KEYATTR_RETURN_REF"
	case CSSM_KEYATTR_SENSITIVE:
		return "CSSM_KEYATTR_SENSITIVE"
	default:
		return fmt.Sprintf("CssmKeyattr(%d)", e)
	}
}

type CssmKeyattrP uint32

const (
	CSSM_KEYATTR_PARTIAL            CssmKeyattrP = 0x10000
	CSSM_KEYATTR_PUBLIC_KEY_ENCRYPT CssmKeyattrP = 0x20000
)

func (e CssmKeyattrP) String() string {
	switch e {
	case CSSM_KEYATTR_PARTIAL:
		return "CSSM_KEYATTR_PARTIAL"
	case CSSM_KEYATTR_PUBLIC_KEY_ENCRYPT:
		return "CSSM_KEYATTR_PUBLIC_KEY_ENCRYPT"
	default:
		return fmt.Sprintf("CssmKeyattrP(%d)", e)
	}
}

type CssmKeyblob uint32

const (
	CSSM_KEYBLOB_OTHER     CssmKeyblob = 0xffffffff
	CSSM_KEYBLOB_RAW       CssmKeyblob = 0
	CSSM_KEYBLOB_REFERENCE CssmKeyblob = 2
	CSSM_KEYBLOB_WRAPPED   CssmKeyblob = 3
)

func (e CssmKeyblob) String() string {
	switch e {
	case CSSM_KEYBLOB_OTHER:
		return "CSSM_KEYBLOB_OTHER"
	case CSSM_KEYBLOB_RAW:
		return "CSSM_KEYBLOB_RAW"
	case CSSM_KEYBLOB_REFERENCE:
		return "CSSM_KEYBLOB_REFERENCE"
	case CSSM_KEYBLOB_WRAPPED:
		return "CSSM_KEYBLOB_WRAPPED"
	default:
		return fmt.Sprintf("CssmKeyblob(%d)", e)
	}
}

type CssmKeyblobRawFormatNone uint32

const (
	CSSM_KEYBLOB_RAW_FORMAT_BSAFE        CssmKeyblobRawFormatNone = 6
	CSSM_KEYBLOB_RAW_FORMAT_CCA          CssmKeyblobRawFormatNone = 9
	CSSM_KEYBLOB_RAW_FORMAT_FIPS186      CssmKeyblobRawFormatNone = 5
	CSSM_KEYBLOB_RAW_FORMAT_MSCAPI       CssmKeyblobRawFormatNone = 3
	CSSM_KEYBLOB_RAW_FORMAT_NONE         CssmKeyblobRawFormatNone = 0
	CSSM_KEYBLOB_RAW_FORMAT_OCTET_STRING CssmKeyblobRawFormatNone = 12
	CSSM_KEYBLOB_RAW_FORMAT_OTHER        CssmKeyblobRawFormatNone = 0xffffffff
	CSSM_KEYBLOB_RAW_FORMAT_PGP          CssmKeyblobRawFormatNone = 4
	CSSM_KEYBLOB_RAW_FORMAT_PKCS1        CssmKeyblobRawFormatNone = 1
	CSSM_KEYBLOB_RAW_FORMAT_PKCS3        CssmKeyblobRawFormatNone = 2
	CSSM_KEYBLOB_RAW_FORMAT_PKCS8        CssmKeyblobRawFormatNone = 10
	CSSM_KEYBLOB_RAW_FORMAT_SPKI         CssmKeyblobRawFormatNone = 11
)

func (e CssmKeyblobRawFormatNone) String() string {
	switch e {
	case CSSM_KEYBLOB_RAW_FORMAT_BSAFE:
		return "CSSM_KEYBLOB_RAW_FORMAT_BSAFE"
	case CSSM_KEYBLOB_RAW_FORMAT_CCA:
		return "CSSM_KEYBLOB_RAW_FORMAT_CCA"
	case CSSM_KEYBLOB_RAW_FORMAT_FIPS186:
		return "CSSM_KEYBLOB_RAW_FORMAT_FIPS186"
	case CSSM_KEYBLOB_RAW_FORMAT_MSCAPI:
		return "CSSM_KEYBLOB_RAW_FORMAT_MSCAPI"
	case CSSM_KEYBLOB_RAW_FORMAT_NONE:
		return "CSSM_KEYBLOB_RAW_FORMAT_NONE"
	case CSSM_KEYBLOB_RAW_FORMAT_OCTET_STRING:
		return "CSSM_KEYBLOB_RAW_FORMAT_OCTET_STRING"
	case CSSM_KEYBLOB_RAW_FORMAT_OTHER:
		return "CSSM_KEYBLOB_RAW_FORMAT_OTHER"
	case CSSM_KEYBLOB_RAW_FORMAT_PGP:
		return "CSSM_KEYBLOB_RAW_FORMAT_PGP"
	case CSSM_KEYBLOB_RAW_FORMAT_PKCS1:
		return "CSSM_KEYBLOB_RAW_FORMAT_PKCS1"
	case CSSM_KEYBLOB_RAW_FORMAT_PKCS3:
		return "CSSM_KEYBLOB_RAW_FORMAT_PKCS3"
	case CSSM_KEYBLOB_RAW_FORMAT_PKCS8:
		return "CSSM_KEYBLOB_RAW_FORMAT_PKCS8"
	case CSSM_KEYBLOB_RAW_FORMAT_SPKI:
		return "CSSM_KEYBLOB_RAW_FORMAT_SPKI"
	default:
		return fmt.Sprintf("CssmKeyblobRawFormatNone(%d)", e)
	}
}

type CssmKeyblobRawFormatVendor uint32

const (
	CSSM_KEYBLOB_RAW_FORMAT_VENDOR_DEFINED CssmKeyblobRawFormatVendor = 0x80000000
)

func (e CssmKeyblobRawFormatVendor) String() string {
	switch e {
	case CSSM_KEYBLOB_RAW_FORMAT_VENDOR_DEFINED:
		return "CSSM_KEYBLOB_RAW_FORMAT_VENDOR_DEFINED"
	default:
		return fmt.Sprintf("CssmKeyblobRawFormatVendor(%d)", e)
	}
}

type CssmKeyblobRawFormatX509 uint32

const (
	CSSM_KEYBLOB_RAW_FORMAT_OPENSSH  CssmKeyblobRawFormatX509 = 2147483649
	CSSM_KEYBLOB_RAW_FORMAT_OPENSSH2 CssmKeyblobRawFormatX509 = 2147483651
	CSSM_KEYBLOB_RAW_FORMAT_OPENSSL  CssmKeyblobRawFormatX509 = 2147483650
	CSSM_KEYBLOB_RAW_FORMAT_X509     CssmKeyblobRawFormatX509 = 2147483648
)

func (e CssmKeyblobRawFormatX509) String() string {
	switch e {
	case CSSM_KEYBLOB_RAW_FORMAT_OPENSSH:
		return "CSSM_KEYBLOB_RAW_FORMAT_OPENSSH"
	case CSSM_KEYBLOB_RAW_FORMAT_OPENSSH2:
		return "CSSM_KEYBLOB_RAW_FORMAT_OPENSSH2"
	case CSSM_KEYBLOB_RAW_FORMAT_OPENSSL:
		return "CSSM_KEYBLOB_RAW_FORMAT_OPENSSL"
	case CSSM_KEYBLOB_RAW_FORMAT_X509:
		return "CSSM_KEYBLOB_RAW_FORMAT_X509"
	default:
		return fmt.Sprintf("CssmKeyblobRawFormatX509(%d)", e)
	}
}

type CssmKeyblobRefFormat uint32

const (
	CSSM_KEYBLOB_REF_FORMAT_INTEGER CssmKeyblobRefFormat = 0
	CSSM_KEYBLOB_REF_FORMAT_OTHER   CssmKeyblobRefFormat = 0xffffffff
	CSSM_KEYBLOB_REF_FORMAT_SPKI    CssmKeyblobRefFormat = 2
	CSSM_KEYBLOB_REF_FORMAT_STRING  CssmKeyblobRefFormat = 1
)

func (e CssmKeyblobRefFormat) String() string {
	switch e {
	case CSSM_KEYBLOB_REF_FORMAT_INTEGER:
		return "CSSM_KEYBLOB_REF_FORMAT_INTEGER"
	case CSSM_KEYBLOB_REF_FORMAT_OTHER:
		return "CSSM_KEYBLOB_REF_FORMAT_OTHER"
	case CSSM_KEYBLOB_REF_FORMAT_SPKI:
		return "CSSM_KEYBLOB_REF_FORMAT_SPKI"
	case CSSM_KEYBLOB_REF_FORMAT_STRING:
		return "CSSM_KEYBLOB_REF_FORMAT_STRING"
	default:
		return fmt.Sprintf("CssmKeyblobRefFormat(%d)", e)
	}
}

type CssmKeyblobWrappedFormatAppleCustom uint32

const (
	CSSM_KEYBLOB_WRAPPED_FORMAT_APPLE_CUSTOM CssmKeyblobWrappedFormatAppleCustom = 100
	CSSM_KEYBLOB_WRAPPED_FORMAT_OPENSSH1     CssmKeyblobWrappedFormatAppleCustom = 102
	CSSM_KEYBLOB_WRAPPED_FORMAT_OPENSSL      CssmKeyblobWrappedFormatAppleCustom = 101
)

func (e CssmKeyblobWrappedFormatAppleCustom) String() string {
	switch e {
	case CSSM_KEYBLOB_WRAPPED_FORMAT_APPLE_CUSTOM:
		return "CSSM_KEYBLOB_WRAPPED_FORMAT_APPLE_CUSTOM"
	case CSSM_KEYBLOB_WRAPPED_FORMAT_OPENSSH1:
		return "CSSM_KEYBLOB_WRAPPED_FORMAT_OPENSSH1"
	case CSSM_KEYBLOB_WRAPPED_FORMAT_OPENSSL:
		return "CSSM_KEYBLOB_WRAPPED_FORMAT_OPENSSL"
	default:
		return fmt.Sprintf("CssmKeyblobWrappedFormatAppleCustom(%d)", e)
	}
}

type CssmKeyblobWrappedFormatNone uint32

const (
	CSSM_KEYBLOB_WRAPPED_FORMAT_MSCAPI CssmKeyblobWrappedFormatNone = 3
	CSSM_KEYBLOB_WRAPPED_FORMAT_NONE   CssmKeyblobWrappedFormatNone = 0
	CSSM_KEYBLOB_WRAPPED_FORMAT_OTHER  CssmKeyblobWrappedFormatNone = 0xffffffff
	CSSM_KEYBLOB_WRAPPED_FORMAT_PKCS7  CssmKeyblobWrappedFormatNone = 2
	CSSM_KEYBLOB_WRAPPED_FORMAT_PKCS8  CssmKeyblobWrappedFormatNone = 1
)

func (e CssmKeyblobWrappedFormatNone) String() string {
	switch e {
	case CSSM_KEYBLOB_WRAPPED_FORMAT_MSCAPI:
		return "CSSM_KEYBLOB_WRAPPED_FORMAT_MSCAPI"
	case CSSM_KEYBLOB_WRAPPED_FORMAT_NONE:
		return "CSSM_KEYBLOB_WRAPPED_FORMAT_NONE"
	case CSSM_KEYBLOB_WRAPPED_FORMAT_OTHER:
		return "CSSM_KEYBLOB_WRAPPED_FORMAT_OTHER"
	case CSSM_KEYBLOB_WRAPPED_FORMAT_PKCS7:
		return "CSSM_KEYBLOB_WRAPPED_FORMAT_PKCS7"
	case CSSM_KEYBLOB_WRAPPED_FORMAT_PKCS8:
		return "CSSM_KEYBLOB_WRAPPED_FORMAT_PKCS8"
	default:
		return fmt.Sprintf("CssmKeyblobWrappedFormatNone(%d)", e)
	}
}

type CssmKeyclass uint32

const (
	CSSM_KEYCLASS_OTHER       CssmKeyclass = 0xffffffff
	CSSM_KEYCLASS_PRIVATE_KEY CssmKeyclass = 1
	CSSM_KEYCLASS_PUBLIC_KEY  CssmKeyclass = 0
	CSSM_KEYCLASS_SECRET_PART CssmKeyclass = 3
	CSSM_KEYCLASS_SESSION_KEY CssmKeyclass = 2
)

func (e CssmKeyclass) String() string {
	switch e {
	case CSSM_KEYCLASS_OTHER:
		return "CSSM_KEYCLASS_OTHER"
	case CSSM_KEYCLASS_PRIVATE_KEY:
		return "CSSM_KEYCLASS_PRIVATE_KEY"
	case CSSM_KEYCLASS_PUBLIC_KEY:
		return "CSSM_KEYCLASS_PUBLIC_KEY"
	case CSSM_KEYCLASS_SECRET_PART:
		return "CSSM_KEYCLASS_SECRET_PART"
	case CSSM_KEYCLASS_SESSION_KEY:
		return "CSSM_KEYCLASS_SESSION_KEY"
	default:
		return fmt.Sprintf("CssmKeyclass(%d)", e)
	}
}

type CssmKeyheader uint32

const (
	CSSM_KEYHEADER_VERSION CssmKeyheader = 2
)

func (e CssmKeyheader) String() string {
	switch e {
	case CSSM_KEYHEADER_VERSION:
		return "CSSM_KEYHEADER_VERSION"
	default:
		return fmt.Sprintf("CssmKeyheader(%d)", e)
	}
}

type CssmKeyuse uint32

const (
	CSSM_KEYUSE_ANY            CssmKeyuse = 0x80000000
	CSSM_KEYUSE_DECRYPT        CssmKeyuse = 0x2
	CSSM_KEYUSE_DERIVE         CssmKeyuse = 0x100
	CSSM_KEYUSE_ENCRYPT        CssmKeyuse = 0x1
	CSSM_KEYUSE_SIGN           CssmKeyuse = 0x4
	CSSM_KEYUSE_SIGN_RECOVER   CssmKeyuse = 0x10
	CSSM_KEYUSE_UNWRAP         CssmKeyuse = 0x80
	CSSM_KEYUSE_VERIFY         CssmKeyuse = 0x8
	CSSM_KEYUSE_VERIFY_RECOVER CssmKeyuse = 0x20
	CSSM_KEYUSE_WRAP           CssmKeyuse = 0x40
)

func (e CssmKeyuse) String() string {
	switch e {
	case CSSM_KEYUSE_ANY:
		return "CSSM_KEYUSE_ANY"
	case CSSM_KEYUSE_DECRYPT:
		return "CSSM_KEYUSE_DECRYPT"
	case CSSM_KEYUSE_DERIVE:
		return "CSSM_KEYUSE_DERIVE"
	case CSSM_KEYUSE_ENCRYPT:
		return "CSSM_KEYUSE_ENCRYPT"
	case CSSM_KEYUSE_SIGN:
		return "CSSM_KEYUSE_SIGN"
	case CSSM_KEYUSE_SIGN_RECOVER:
		return "CSSM_KEYUSE_SIGN_RECOVER"
	case CSSM_KEYUSE_UNWRAP:
		return "CSSM_KEYUSE_UNWRAP"
	case CSSM_KEYUSE_VERIFY:
		return "CSSM_KEYUSE_VERIFY"
	case CSSM_KEYUSE_VERIFY_RECOVER:
		return "CSSM_KEYUSE_VERIFY_RECOVER"
	case CSSM_KEYUSE_WRAP:
		return "CSSM_KEYUSE_WRAP"
	default:
		return fmt.Sprintf("CssmKeyuse(%d)", e)
	}
}

type CssmListElement uint32

const (
	CSSM_LIST_ELEMENT_DATUM   CssmListElement = 0
	CSSM_LIST_ELEMENT_SUBLIST CssmListElement = 0x1
	CSSM_LIST_ELEMENT_WORDID  CssmListElement = 0x2
)

func (e CssmListElement) String() string {
	switch e {
	case CSSM_LIST_ELEMENT_DATUM:
		return "CSSM_LIST_ELEMENT_DATUM"
	case CSSM_LIST_ELEMENT_SUBLIST:
		return "CSSM_LIST_ELEMENT_SUBLIST"
	case CSSM_LIST_ELEMENT_WORDID:
		return "CSSM_LIST_ELEMENT_WORDID"
	default:
		return fmt.Sprintf("CssmListElement(%d)", e)
	}
}

type CssmListType uint32

const (
	CSSM_LIST_TYPE_CUSTOM  CssmListType = 1
	CSSM_LIST_TYPE_SEXPR   CssmListType = 2
	CSSM_LIST_TYPE_UNKNOWN CssmListType = 0
)

func (e CssmListType) String() string {
	switch e {
	case CSSM_LIST_TYPE_CUSTOM:
		return "CSSM_LIST_TYPE_CUSTOM"
	case CSSM_LIST_TYPE_SEXPR:
		return "CSSM_LIST_TYPE_SEXPR"
	case CSSM_LIST_TYPE_UNKNOWN:
		return "CSSM_LIST_TYPE_UNKNOWN"
	default:
		return fmt.Sprintf("CssmListType(%d)", e)
	}
}

type CssmMds int32

const (
	CSSM_MDS_BASE_ERROR    CssmMds = -2147414016
	CSSM_MDS_PRIVATE_ERROR CssmMds = -2147412992
)

func (e CssmMds) String() string {
	switch e {
	case CSSM_MDS_BASE_ERROR:
		return "CSSM_MDS_BASE_ERROR"
	case CSSM_MDS_PRIVATE_ERROR:
		return "CSSM_MDS_PRIVATE_ERROR"
	default:
		return fmt.Sprintf("CssmMds(%d)", e)
	}
}

type CssmModuleString uint32

const (
	CSSM_MODULE_STRING_SIZE CssmModuleString = 64
)

func (e CssmModuleString) String() string {
	switch e {
	case CSSM_MODULE_STRING_SIZE:
		return "CSSM_MODULE_STRING_SIZE"
	default:
		return fmt.Sprintf("CssmModuleString(%d)", e)
	}
}

type CssmNetProto uint32

const (
	CSSM_NET_PROTO_CMP         CssmNetProto = 10
	CSSM_NET_PROTO_CMPS        CssmNetProto = 11
	CSSM_NET_PROTO_CUSTOM      CssmNetProto = 1
	CSSM_NET_PROTO_FTP         CssmNetProto = 7
	CSSM_NET_PROTO_FTPS        CssmNetProto = 8
	CSSM_NET_PROTO_LDAP        CssmNetProto = 3
	CSSM_NET_PROTO_LDAPNS      CssmNetProto = 5
	CSSM_NET_PROTO_LDAPS       CssmNetProto = 4
	CSSM_NET_PROTO_NONE        CssmNetProto = 0
	CSSM_NET_PROTO_OCSP        CssmNetProto = 9
	CSSM_NET_PROTO_UNSPECIFIED CssmNetProto = 2
	CSSM_NET_PROTO_X500DAP     CssmNetProto = 6
)

func (e CssmNetProto) String() string {
	switch e {
	case CSSM_NET_PROTO_CMP:
		return "CSSM_NET_PROTO_CMP"
	case CSSM_NET_PROTO_CMPS:
		return "CSSM_NET_PROTO_CMPS"
	case CSSM_NET_PROTO_CUSTOM:
		return "CSSM_NET_PROTO_CUSTOM"
	case CSSM_NET_PROTO_FTP:
		return "CSSM_NET_PROTO_FTP"
	case CSSM_NET_PROTO_FTPS:
		return "CSSM_NET_PROTO_FTPS"
	case CSSM_NET_PROTO_LDAP:
		return "CSSM_NET_PROTO_LDAP"
	case CSSM_NET_PROTO_LDAPNS:
		return "CSSM_NET_PROTO_LDAPNS"
	case CSSM_NET_PROTO_LDAPS:
		return "CSSM_NET_PROTO_LDAPS"
	case CSSM_NET_PROTO_NONE:
		return "CSSM_NET_PROTO_NONE"
	case CSSM_NET_PROTO_OCSP:
		return "CSSM_NET_PROTO_OCSP"
	case CSSM_NET_PROTO_UNSPECIFIED:
		return "CSSM_NET_PROTO_UNSPECIFIED"
	case CSSM_NET_PROTO_X500DAP:
		return "CSSM_NET_PROTO_X500DAP"
	default:
		return fmt.Sprintf("CssmNetProto(%d)", e)
	}
}

type CssmNotify uint32

const (
	CSSM_NOTIFY_FAULT  CssmNotify = 3
	CSSM_NOTIFY_INSERT CssmNotify = 1
	CSSM_NOTIFY_REMOVE CssmNotify = 2
)

func (e CssmNotify) String() string {
	switch e {
	case CSSM_NOTIFY_FAULT:
		return "CSSM_NOTIFY_FAULT"
	case CSSM_NOTIFY_INSERT:
		return "CSSM_NOTIFY_INSERT"
	case CSSM_NOTIFY_REMOVE:
		return "CSSM_NOTIFY_REMOVE"
	default:
		return fmt.Sprintf("CssmNotify(%d)", e)
	}
}

type CssmOk uint32

const (
	CSSM_OK CssmOk = 0
)

func (e CssmOk) String() string {
	switch e {
	case CSSM_OK:
		return "CSSM_OK"
	default:
		return fmt.Sprintf("CssmOk(%d)", e)
	}
}

type CssmPadding uint32

const (
	CSSM_PADDING_ALTERNATE      CssmPadding = 4
	CSSM_PADDING_CIPHERSTEALING CssmPadding = 8
	CSSM_PADDING_CUSTOM         CssmPadding = 1
	CSSM_PADDING_FF             CssmPadding = 5
	CSSM_PADDING_NONE           CssmPadding = 0
	CSSM_PADDING_ONE            CssmPadding = 3
	CSSM_PADDING_PKCS1          CssmPadding = 10
	CSSM_PADDING_PKCS5          CssmPadding = 6
	CSSM_PADDING_PKCS7          CssmPadding = 7
	CSSM_PADDING_RANDOM         CssmPadding = 9
	CSSM_PADDING_SIGRAW         CssmPadding = 11
	CSSM_PADDING_VENDOR_DEFINED CssmPadding = 2147483648
	CSSM_PADDING_ZERO           CssmPadding = 2
)

func (e CssmPadding) String() string {
	switch e {
	case CSSM_PADDING_ALTERNATE:
		return "CSSM_PADDING_ALTERNATE"
	case CSSM_PADDING_CIPHERSTEALING:
		return "CSSM_PADDING_CIPHERSTEALING"
	case CSSM_PADDING_CUSTOM:
		return "CSSM_PADDING_CUSTOM"
	case CSSM_PADDING_FF:
		return "CSSM_PADDING_FF"
	case CSSM_PADDING_NONE:
		return "CSSM_PADDING_NONE"
	case CSSM_PADDING_ONE:
		return "CSSM_PADDING_ONE"
	case CSSM_PADDING_PKCS1:
		return "CSSM_PADDING_PKCS1"
	case CSSM_PADDING_PKCS5:
		return "CSSM_PADDING_PKCS5"
	case CSSM_PADDING_PKCS7:
		return "CSSM_PADDING_PKCS7"
	case CSSM_PADDING_RANDOM:
		return "CSSM_PADDING_RANDOM"
	case CSSM_PADDING_SIGRAW:
		return "CSSM_PADDING_SIGRAW"
	case CSSM_PADDING_VENDOR_DEFINED:
		return "CSSM_PADDING_VENDOR_DEFINED"
	case CSSM_PADDING_ZERO:
		return "CSSM_PADDING_ZERO"
	default:
		return fmt.Sprintf("CssmPadding(%d)", e)
	}
}

type CssmPaddingApple uint32

const (
	CSSM_PADDING_APPLE_SSLv2 CssmPaddingApple = 2147483648
)

func (e CssmPaddingApple) String() string {
	switch e {
	case CSSM_PADDING_APPLE_SSLv2:
		return "CSSM_PADDING_APPLE_SSLv2"
	default:
		return fmt.Sprintf("CssmPaddingApple(%d)", e)
	}
}

type CssmPkcs5Pbkdf2PrfHmac uint32

const (
	CSSM_PKCS5_PBKDF2_PRF_HMAC_SHA1 CssmPkcs5Pbkdf2PrfHmac = 0
)

func (e CssmPkcs5Pbkdf2PrfHmac) String() string {
	switch e {
	case CSSM_PKCS5_PBKDF2_PRF_HMAC_SHA1:
		return "CSSM_PKCS5_PBKDF2_PRF_HMAC_SHA1"
	default:
		return fmt.Sprintf("CssmPkcs5Pbkdf2PrfHmac(%d)", e)
	}
}

type CssmPkcsOaep uint32

const (
	CSSM_PKCS_OAEP_MGF1_MD5  CssmPkcsOaep = 2
	CSSM_PKCS_OAEP_MGF1_SHA1 CssmPkcsOaep = 1
	CSSM_PKCS_OAEP_MGF_NONE  CssmPkcsOaep = 0
)

func (e CssmPkcsOaep) String() string {
	switch e {
	case CSSM_PKCS_OAEP_MGF1_MD5:
		return "CSSM_PKCS_OAEP_MGF1_MD5"
	case CSSM_PKCS_OAEP_MGF1_SHA1:
		return "CSSM_PKCS_OAEP_MGF1_SHA1"
	case CSSM_PKCS_OAEP_MGF_NONE:
		return "CSSM_PKCS_OAEP_MGF_NONE"
	default:
		return fmt.Sprintf("CssmPkcsOaep(%d)", e)
	}
}

type CssmPkcsOaepPsource uint32

const (
	CSSM_PKCS_OAEP_PSOURCE_NONE       CssmPkcsOaepPsource = 0
	CSSM_PKCS_OAEP_PSOURCE_Pspecified CssmPkcsOaepPsource = 1
)

func (e CssmPkcsOaepPsource) String() string {
	switch e {
	case CSSM_PKCS_OAEP_PSOURCE_NONE:
		return "CSSM_PKCS_OAEP_PSOURCE_NONE"
	case CSSM_PKCS_OAEP_PSOURCE_Pspecified:
		return "CSSM_PKCS_OAEP_PSOURCE_Pspecified"
	default:
		return fmt.Sprintf("CssmPkcsOaepPsource(%d)", e)
	}
}

type CssmPrivilegeScope uint32

const (
	CSSM_PRIVILEGE_SCOPE_NONE    CssmPrivilegeScope = 0
	CSSM_PRIVILEGE_SCOPE_PROCESS CssmPrivilegeScope = 1
	CSSM_PRIVILEGE_SCOPE_THREAD  CssmPrivilegeScope = 2
)

func (e CssmPrivilegeScope) String() string {
	switch e {
	case CSSM_PRIVILEGE_SCOPE_NONE:
		return "CSSM_PRIVILEGE_SCOPE_NONE"
	case CSSM_PRIVILEGE_SCOPE_PROCESS:
		return "CSSM_PRIVILEGE_SCOPE_PROCESS"
	case CSSM_PRIVILEGE_SCOPE_THREAD:
		return "CSSM_PRIVILEGE_SCOPE_THREAD"
	default:
		return fmt.Sprintf("CssmPrivilegeScope(%d)", e)
	}
}

type CssmPvc uint32

const (
	CSSM_PVC_APP  CssmPvc = 1
	CSSM_PVC_NONE CssmPvc = 0
	CSSM_PVC_SP   CssmPvc = 2
)

func (e CssmPvc) String() string {
	switch e {
	case CSSM_PVC_APP:
		return "CSSM_PVC_APP"
	case CSSM_PVC_NONE:
		return "CSSM_PVC_NONE"
	case CSSM_PVC_SP:
		return "CSSM_PVC_SP"
	default:
		return fmt.Sprintf("CssmPvc(%d)", e)
	}
}

type CssmQueryReturn uint32

const (
	CSSM_QUERY_RETURN_DATA CssmQueryReturn = 0x1
)

func (e CssmQueryReturn) String() string {
	switch e {
	case CSSM_QUERY_RETURN_DATA:
		return "CSSM_QUERY_RETURN_DATA"
	default:
		return fmt.Sprintf("CssmQueryReturn(%d)", e)
	}
}

type CssmQuerySizelimit uint32

const (
	CSSM_QUERY_SIZELIMIT_NONE CssmQuerySizelimit = 0
)

func (e CssmQuerySizelimit) String() string {
	switch e {
	case CSSM_QUERY_SIZELIMIT_NONE:
		return "CSSM_QUERY_SIZELIMIT_NONE"
	default:
		return fmt.Sprintf("CssmQuerySizelimit(%d)", e)
	}
}

type CssmQueryTimelimit uint32

const (
	CSSM_QUERY_TIMELIMIT_NONE CssmQueryTimelimit = 0
)

func (e CssmQueryTimelimit) String() string {
	switch e {
	case CSSM_QUERY_TIMELIMIT_NONE:
		return "CSSM_QUERY_TIMELIMIT_NONE"
	default:
		return fmt.Sprintf("CssmQueryTimelimit(%d)", e)
	}
}

type CssmSampleTypeKeychainPrompt uint32

const (
	CSSM_SAMPLE_TYPE_ASYMMETRIC_KEY       CssmSampleTypeKeychainPrompt = 65547
	CSSM_SAMPLE_TYPE_COMMENT              CssmSampleTypeKeychainPrompt = 12
	CSSM_SAMPLE_TYPE_KEYBAG_KEY           CssmSampleTypeKeychainPrompt = 65549
	CSSM_SAMPLE_TYPE_KEYCHAIN_CHANGE_LOCK CssmSampleTypeKeychainPrompt = 65538
	CSSM_SAMPLE_TYPE_KEYCHAIN_LOCK        CssmSampleTypeKeychainPrompt = 65537
	CSSM_SAMPLE_TYPE_KEYCHAIN_PROMPT      CssmSampleTypeKeychainPrompt = 65536
	CSSM_SAMPLE_TYPE_PREAUTH              CssmSampleTypeKeychainPrompt = 65545
	CSSM_SAMPLE_TYPE_PROCESS              CssmSampleTypeKeychainPrompt = 65539
	CSSM_SAMPLE_TYPE_RETRY_ID             CssmSampleTypeKeychainPrompt = 85
	CSSM_SAMPLE_TYPE_SYMMETRIC_KEY        CssmSampleTypeKeychainPrompt = 65541
)

func (e CssmSampleTypeKeychainPrompt) String() string {
	switch e {
	case CSSM_SAMPLE_TYPE_ASYMMETRIC_KEY:
		return "CSSM_SAMPLE_TYPE_ASYMMETRIC_KEY"
	case CSSM_SAMPLE_TYPE_COMMENT:
		return "CSSM_SAMPLE_TYPE_COMMENT"
	case CSSM_SAMPLE_TYPE_KEYBAG_KEY:
		return "CSSM_SAMPLE_TYPE_KEYBAG_KEY"
	case CSSM_SAMPLE_TYPE_KEYCHAIN_CHANGE_LOCK:
		return "CSSM_SAMPLE_TYPE_KEYCHAIN_CHANGE_LOCK"
	case CSSM_SAMPLE_TYPE_KEYCHAIN_LOCK:
		return "CSSM_SAMPLE_TYPE_KEYCHAIN_LOCK"
	case CSSM_SAMPLE_TYPE_KEYCHAIN_PROMPT:
		return "CSSM_SAMPLE_TYPE_KEYCHAIN_PROMPT"
	case CSSM_SAMPLE_TYPE_PREAUTH:
		return "CSSM_SAMPLE_TYPE_PREAUTH"
	case CSSM_SAMPLE_TYPE_PROCESS:
		return "CSSM_SAMPLE_TYPE_PROCESS"
	case CSSM_SAMPLE_TYPE_RETRY_ID:
		return "CSSM_SAMPLE_TYPE_RETRY_ID"
	case CSSM_SAMPLE_TYPE_SYMMETRIC_KEY:
		return "CSSM_SAMPLE_TYPE_SYMMETRIC_KEY"
	default:
		return fmt.Sprintf("CssmSampleTypeKeychainPrompt(%d)", e)
	}
}

type CssmSampleTypePassword uint32

const (
	CSSM_SAMPLE_TYPE_BIOMETRIC           CssmSampleTypePassword = 8
	CSSM_SAMPLE_TYPE_HASHED_PASSWORD     CssmSampleTypePassword = 43
	CSSM_SAMPLE_TYPE_PASSWORD            CssmSampleTypePassword = 79
	CSSM_SAMPLE_TYPE_PROMPTED_BIOMETRIC  CssmSampleTypePassword = 83
	CSSM_SAMPLE_TYPE_PROMPTED_PASSWORD   CssmSampleTypePassword = 84
	CSSM_SAMPLE_TYPE_PROTECTED_BIOMETRIC CssmSampleTypePassword = 86
	CSSM_SAMPLE_TYPE_PROTECTED_PASSWORD  CssmSampleTypePassword = 87
	CSSM_SAMPLE_TYPE_SIGNED_NONCE        CssmSampleTypePassword = 117
	CSSM_SAMPLE_TYPE_SIGNED_SECRET       CssmSampleTypePassword = 118
	CSSM_SAMPLE_TYPE_THRESHOLD           CssmSampleTypePassword = 123
)

func (e CssmSampleTypePassword) String() string {
	switch e {
	case CSSM_SAMPLE_TYPE_BIOMETRIC:
		return "CSSM_SAMPLE_TYPE_BIOMETRIC"
	case CSSM_SAMPLE_TYPE_HASHED_PASSWORD:
		return "CSSM_SAMPLE_TYPE_HASHED_PASSWORD"
	case CSSM_SAMPLE_TYPE_PASSWORD:
		return "CSSM_SAMPLE_TYPE_PASSWORD"
	case CSSM_SAMPLE_TYPE_PROMPTED_BIOMETRIC:
		return "CSSM_SAMPLE_TYPE_PROMPTED_BIOMETRIC"
	case CSSM_SAMPLE_TYPE_PROMPTED_PASSWORD:
		return "CSSM_SAMPLE_TYPE_PROMPTED_PASSWORD"
	case CSSM_SAMPLE_TYPE_PROTECTED_BIOMETRIC:
		return "CSSM_SAMPLE_TYPE_PROTECTED_BIOMETRIC"
	case CSSM_SAMPLE_TYPE_PROTECTED_PASSWORD:
		return "CSSM_SAMPLE_TYPE_PROTECTED_PASSWORD"
	case CSSM_SAMPLE_TYPE_SIGNED_NONCE:
		return "CSSM_SAMPLE_TYPE_SIGNED_NONCE"
	case CSSM_SAMPLE_TYPE_SIGNED_SECRET:
		return "CSSM_SAMPLE_TYPE_SIGNED_SECRET"
	case CSSM_SAMPLE_TYPE_THRESHOLD:
		return "CSSM_SAMPLE_TYPE_THRESHOLD"
	default:
		return fmt.Sprintf("CssmSampleTypePassword(%d)", e)
	}
}

type CssmService uint32

const (
	CSSM_SERVICE_AC   CssmService = 0x20
	CSSM_SERVICE_CL   CssmService = 0x8
	CSSM_SERVICE_CSP  CssmService = 0x2
	CSSM_SERVICE_CSSM CssmService = 0x1
	CSSM_SERVICE_DL   CssmService = 0x4
	CSSM_SERVICE_KR   CssmService = 0x40
	CSSM_SERVICE_TP   CssmService = 0x10
)

func (e CssmService) String() string {
	switch e {
	case CSSM_SERVICE_AC:
		return "CSSM_SERVICE_AC"
	case CSSM_SERVICE_CL:
		return "CSSM_SERVICE_CL"
	case CSSM_SERVICE_CSP:
		return "CSSM_SERVICE_CSP"
	case CSSM_SERVICE_CSSM:
		return "CSSM_SERVICE_CSSM"
	case CSSM_SERVICE_DL:
		return "CSSM_SERVICE_DL"
	case CSSM_SERVICE_KR:
		return "CSSM_SERVICE_KR"
	case CSSM_SERVICE_TP:
		return "CSSM_SERVICE_TP"
	default:
		return fmt.Sprintf("CssmService(%d)", e)
	}
}

type CssmTp uint32

const (
	CSSM_TP_CERT_DIR_UPDATE   CssmTp = 0x8
	CSSM_TP_CERT_NOTIFY_RENEW CssmTp = 0x4
	CSSM_TP_CERT_PUBLISH      CssmTp = 0x2
	CSSM_TP_CRL_DISTRIBUTE    CssmTp = 0x10
	CSSM_TP_KEY_ARCHIVE       CssmTp = 0x1
)

func (e CssmTp) String() string {
	switch e {
	case CSSM_TP_CERT_DIR_UPDATE:
		return "CSSM_TP_CERT_DIR_UPDATE"
	case CSSM_TP_CERT_NOTIFY_RENEW:
		return "CSSM_TP_CERT_NOTIFY_RENEW"
	case CSSM_TP_CERT_PUBLISH:
		return "CSSM_TP_CERT_PUBLISH"
	case CSSM_TP_CRL_DISTRIBUTE:
		return "CSSM_TP_CRL_DISTRIBUTE"
	case CSSM_TP_KEY_ARCHIVE:
		return "CSSM_TP_KEY_ARCHIVE"
	default:
		return fmt.Sprintf("CssmTp(%d)", e)
	}
}

type CssmTpActionAllowExpired uint32

const (
	CSSM_TP_ACTION_ALLOW_EXPIRED        CssmTpActionAllowExpired = 0x1
	CSSM_TP_ACTION_ALLOW_EXPIRED_ROOT   CssmTpActionAllowExpired = 0x8
	CSSM_TP_ACTION_FETCH_CERT_FROM_NET  CssmTpActionAllowExpired = 0x4
	CSSM_TP_ACTION_IMPLICIT_ANCHORS     CssmTpActionAllowExpired = 0x40
	CSSM_TP_ACTION_LEAF_IS_CA           CssmTpActionAllowExpired = 0x2
	CSSM_TP_ACTION_REQUIRE_REV_PER_CERT CssmTpActionAllowExpired = 0x10
	CSSM_TP_ACTION_TRUST_SETTINGS       CssmTpActionAllowExpired = 0x20
)

func (e CssmTpActionAllowExpired) String() string {
	switch e {
	case CSSM_TP_ACTION_ALLOW_EXPIRED:
		return "CSSM_TP_ACTION_ALLOW_EXPIRED"
	case CSSM_TP_ACTION_ALLOW_EXPIRED_ROOT:
		return "CSSM_TP_ACTION_ALLOW_EXPIRED_ROOT"
	case CSSM_TP_ACTION_FETCH_CERT_FROM_NET:
		return "CSSM_TP_ACTION_FETCH_CERT_FROM_NET"
	case CSSM_TP_ACTION_IMPLICIT_ANCHORS:
		return "CSSM_TP_ACTION_IMPLICIT_ANCHORS"
	case CSSM_TP_ACTION_LEAF_IS_CA:
		return "CSSM_TP_ACTION_LEAF_IS_CA"
	case CSSM_TP_ACTION_REQUIRE_REV_PER_CERT:
		return "CSSM_TP_ACTION_REQUIRE_REV_PER_CERT"
	case CSSM_TP_ACTION_TRUST_SETTINGS:
		return "CSSM_TP_ACTION_TRUST_SETTINGS"
	default:
		return fmt.Sprintf("CssmTpActionAllowExpired(%d)", e)
	}
}

type CssmTpActionDefault uint32

const (
	CSSM_TP_ACTION_DEFAULT CssmTpActionDefault = 0
)

func (e CssmTpActionDefault) String() string {
	switch e {
	case CSSM_TP_ACTION_DEFAULT:
		return "CSSM_TP_ACTION_DEFAULT"
	default:
		return fmt.Sprintf("CssmTpActionDefault(%d)", e)
	}
}

type CssmTpActionRequireCrlPerCert uint32

const (
	CSSM_TP_ACTION_CRL_SUFFICIENT         CssmTpActionRequireCrlPerCert = 0x4
	CSSM_TP_ACTION_FETCH_CRL_FROM_NET     CssmTpActionRequireCrlPerCert = 0x2
	CSSM_TP_ACTION_REQUIRE_CRL_IF_PRESENT CssmTpActionRequireCrlPerCert = 0x8
	CSSM_TP_ACTION_REQUIRE_CRL_PER_CERT   CssmTpActionRequireCrlPerCert = 0x1
)

func (e CssmTpActionRequireCrlPerCert) String() string {
	switch e {
	case CSSM_TP_ACTION_CRL_SUFFICIENT:
		return "CSSM_TP_ACTION_CRL_SUFFICIENT"
	case CSSM_TP_ACTION_FETCH_CRL_FROM_NET:
		return "CSSM_TP_ACTION_FETCH_CRL_FROM_NET"
	case CSSM_TP_ACTION_REQUIRE_CRL_IF_PRESENT:
		return "CSSM_TP_ACTION_REQUIRE_CRL_IF_PRESENT"
	case CSSM_TP_ACTION_REQUIRE_CRL_PER_CERT:
		return "CSSM_TP_ACTION_REQUIRE_CRL_PER_CERT"
	default:
		return fmt.Sprintf("CssmTpActionRequireCrlPerCert(%d)", e)
	}
}

type CssmTpAuthorityRequestC uint32

const (
	CSSM_TP_AUTHORITY_REQUEST_CERTISSUE      CssmTpAuthorityRequestC = 0x1
	CSSM_TP_AUTHORITY_REQUEST_CERTNOTARIZE   CssmTpAuthorityRequestC = 0x6
	CSSM_TP_AUTHORITY_REQUEST_CERTRESUME     CssmTpAuthorityRequestC = 0x4
	CSSM_TP_AUTHORITY_REQUEST_CERTREVOKE     CssmTpAuthorityRequestC = 0x2
	CSSM_TP_AUTHORITY_REQUEST_CERTSUSPEND    CssmTpAuthorityRequestC = 0x3
	CSSM_TP_AUTHORITY_REQUEST_CERTUSERECOVER CssmTpAuthorityRequestC = 0x7
	CSSM_TP_AUTHORITY_REQUEST_CERTVERIFY     CssmTpAuthorityRequestC = 0x5
	CSSM_TP_AUTHORITY_REQUEST_CRLISSUE       CssmTpAuthorityRequestC = 0x100
)

func (e CssmTpAuthorityRequestC) String() string {
	switch e {
	case CSSM_TP_AUTHORITY_REQUEST_CERTISSUE:
		return "CSSM_TP_AUTHORITY_REQUEST_CERTISSUE"
	case CSSM_TP_AUTHORITY_REQUEST_CERTNOTARIZE:
		return "CSSM_TP_AUTHORITY_REQUEST_CERTNOTARIZE"
	case CSSM_TP_AUTHORITY_REQUEST_CERTRESUME:
		return "CSSM_TP_AUTHORITY_REQUEST_CERTRESUME"
	case CSSM_TP_AUTHORITY_REQUEST_CERTREVOKE:
		return "CSSM_TP_AUTHORITY_REQUEST_CERTREVOKE"
	case CSSM_TP_AUTHORITY_REQUEST_CERTSUSPEND:
		return "CSSM_TP_AUTHORITY_REQUEST_CERTSUSPEND"
	case CSSM_TP_AUTHORITY_REQUEST_CERTUSERECOVER:
		return "CSSM_TP_AUTHORITY_REQUEST_CERTUSERECOVER"
	case CSSM_TP_AUTHORITY_REQUEST_CERTVERIFY:
		return "CSSM_TP_AUTHORITY_REQUEST_CERTVERIFY"
	case CSSM_TP_AUTHORITY_REQUEST_CRLISSUE:
		return "CSSM_TP_AUTHORITY_REQUEST_CRLISSUE"
	default:
		return fmt.Sprintf("CssmTpAuthorityRequestC(%d)", e)
	}
}

type CssmTpBaseTpError int32

const (
	CSSMERR_TP_AUTHENTICATION_FAILED              CssmTpBaseTpError = -2147409657
	CSSMERR_TP_CERTGROUP_INCOMPLETE               CssmTpBaseTpError = -2147409656
	CSSMERR_TP_CERTIFICATE_CANT_OPERATE           CssmTpBaseTpError = -2147409655
	CSSMERR_TP_CERT_EXPIRED                       CssmTpBaseTpError = -2147409654
	CSSMERR_TP_CERT_NOT_VALID_YET                 CssmTpBaseTpError = -2147409653
	CSSMERR_TP_CERT_REVOKED                       CssmTpBaseTpError = -2147409652
	CSSMERR_TP_CERT_SUSPENDED                     CssmTpBaseTpError = -2147409651
	CSSMERR_TP_INSUFFICIENT_CREDENTIALS           CssmTpBaseTpError = -2147409650
	CSSMERR_TP_INVALID_ACTION                     CssmTpBaseTpError = -2147409649
	CSSMERR_TP_INVALID_ACTION_DATA                CssmTpBaseTpError = -2147409648
	CSSMERR_TP_INVALID_ANCHOR_CERT                CssmTpBaseTpError = -2147409646
	CSSMERR_TP_INVALID_AUTHORITY                  CssmTpBaseTpError = -2147409645
	CSSMERR_TP_INVALID_CALLBACK                   CssmTpBaseTpError = -2147409625
	CSSMERR_TP_INVALID_CALLERAUTH_CONTEXT_POINTER CssmTpBaseTpError = -2147409663
	CSSMERR_TP_INVALID_CERTGROUP                  CssmTpBaseTpError = -2147409660
	CSSMERR_TP_INVALID_CERTIFICATE                CssmTpBaseTpError = -2147409643
	CSSMERR_TP_INVALID_CERT_AUTHORITY             CssmTpBaseTpError = -2147409642
	CSSMERR_TP_INVALID_CRL                        CssmTpBaseTpError = -2147409638
	CSSMERR_TP_INVALID_CRLGROUP                   CssmTpBaseTpError = -2147409659
	CSSMERR_TP_INVALID_CRLGROUP_POINTER           CssmTpBaseTpError = -2147409658
	CSSMERR_TP_INVALID_CRL_AUTHORITY              CssmTpBaseTpError = -2147409641
	CSSMERR_TP_INVALID_CRL_ENCODING               CssmTpBaseTpError = -2147409640
	CSSMERR_TP_INVALID_CRL_TYPE                   CssmTpBaseTpError = -2147409639
	CSSMERR_TP_INVALID_FORM_TYPE                  CssmTpBaseTpError = -2147409637
	CSSMERR_TP_INVALID_ID                         CssmTpBaseTpError = -2147409636
	CSSMERR_TP_INVALID_IDENTIFIER                 CssmTpBaseTpError = -2147409635
	CSSMERR_TP_INVALID_IDENTIFIER_POINTER         CssmTpBaseTpError = -2147409662
	CSSMERR_TP_INVALID_INDEX                      CssmTpBaseTpError = -2147409634
	CSSMERR_TP_INVALID_KEYCACHE_HANDLE            CssmTpBaseTpError = -2147409661
	CSSMERR_TP_INVALID_NAME                       CssmTpBaseTpError = -2147409633
	CSSMERR_TP_INVALID_POLICY_IDENTIFIERS         CssmTpBaseTpError = -2147409632
	CSSMERR_TP_INVALID_REASON                     CssmTpBaseTpError = -2147409630
	CSSMERR_TP_INVALID_REQUEST_INPUTS             CssmTpBaseTpError = -2147409629
	CSSMERR_TP_INVALID_RESPONSE_VECTOR            CssmTpBaseTpError = -2147409628
	CSSMERR_TP_INVALID_SIGNATURE                  CssmTpBaseTpError = -2147409627
	CSSMERR_TP_INVALID_STOP_ON_POLICY             CssmTpBaseTpError = -2147409626
	CSSMERR_TP_INVALID_TIMESTRING                 CssmTpBaseTpError = -2147409631
	CSSMERR_TP_INVALID_TUPLE                      CssmTpBaseTpError = -2147409624
	CSSMERR_TP_INVALID_TUPLEGROUP                 CssmTpBaseTpError = -2147409614
	CSSMERR_TP_INVALID_TUPLEGROUP_POINTER         CssmTpBaseTpError = -2147409615
	CSSMERR_TP_NOT_SIGNER                         CssmTpBaseTpError = -2147409623
	CSSMERR_TP_NOT_TRUSTED                        CssmTpBaseTpError = -2147409622
	CSSMERR_TP_NO_DEFAULT_AUTHORITY               CssmTpBaseTpError = -2147409621
	CSSMERR_TP_REJECTED_FORM                      CssmTpBaseTpError = -2147409620
	CSSMERR_TP_REQUEST_LOST                       CssmTpBaseTpError = -2147409619
	CSSMERR_TP_REQUEST_REJECTED                   CssmTpBaseTpError = -2147409618
	CSSMERR_TP_UNSUPPORTED_ADDR_TYPE              CssmTpBaseTpError = -2147409617
	CSSMERR_TP_UNSUPPORTED_SERVICE                CssmTpBaseTpError = -2147409616
	CSSMERR_TP_VERIFY_ACTION_FAILED               CssmTpBaseTpError = -2147409644
	CSSM_TP_BASE_TP_ERROR                         CssmTpBaseTpError = -2147409664
)

func (e CssmTpBaseTpError) String() string {
	switch e {
	case CSSMERR_TP_AUTHENTICATION_FAILED:
		return "CSSMERR_TP_AUTHENTICATION_FAILED"
	case CSSMERR_TP_CERTGROUP_INCOMPLETE:
		return "CSSMERR_TP_CERTGROUP_INCOMPLETE"
	case CSSMERR_TP_CERTIFICATE_CANT_OPERATE:
		return "CSSMERR_TP_CERTIFICATE_CANT_OPERATE"
	case CSSMERR_TP_CERT_EXPIRED:
		return "CSSMERR_TP_CERT_EXPIRED"
	case CSSMERR_TP_CERT_NOT_VALID_YET:
		return "CSSMERR_TP_CERT_NOT_VALID_YET"
	case CSSMERR_TP_CERT_REVOKED:
		return "CSSMERR_TP_CERT_REVOKED"
	case CSSMERR_TP_CERT_SUSPENDED:
		return "CSSMERR_TP_CERT_SUSPENDED"
	case CSSMERR_TP_INSUFFICIENT_CREDENTIALS:
		return "CSSMERR_TP_INSUFFICIENT_CREDENTIALS"
	case CSSMERR_TP_INVALID_ACTION:
		return "CSSMERR_TP_INVALID_ACTION"
	case CSSMERR_TP_INVALID_ACTION_DATA:
		return "CSSMERR_TP_INVALID_ACTION_DATA"
	case CSSMERR_TP_INVALID_ANCHOR_CERT:
		return "CSSMERR_TP_INVALID_ANCHOR_CERT"
	case CSSMERR_TP_INVALID_AUTHORITY:
		return "CSSMERR_TP_INVALID_AUTHORITY"
	case CSSMERR_TP_INVALID_CALLBACK:
		return "CSSMERR_TP_INVALID_CALLBACK"
	case CSSMERR_TP_INVALID_CALLERAUTH_CONTEXT_POINTER:
		return "CSSMERR_TP_INVALID_CALLERAUTH_CONTEXT_POINTER"
	case CSSMERR_TP_INVALID_CERTGROUP:
		return "CSSMERR_TP_INVALID_CERTGROUP"
	case CSSMERR_TP_INVALID_CERTIFICATE:
		return "CSSMERR_TP_INVALID_CERTIFICATE"
	case CSSMERR_TP_INVALID_CERT_AUTHORITY:
		return "CSSMERR_TP_INVALID_CERT_AUTHORITY"
	case CSSMERR_TP_INVALID_CRL:
		return "CSSMERR_TP_INVALID_CRL"
	case CSSMERR_TP_INVALID_CRLGROUP:
		return "CSSMERR_TP_INVALID_CRLGROUP"
	case CSSMERR_TP_INVALID_CRLGROUP_POINTER:
		return "CSSMERR_TP_INVALID_CRLGROUP_POINTER"
	case CSSMERR_TP_INVALID_CRL_AUTHORITY:
		return "CSSMERR_TP_INVALID_CRL_AUTHORITY"
	case CSSMERR_TP_INVALID_CRL_ENCODING:
		return "CSSMERR_TP_INVALID_CRL_ENCODING"
	case CSSMERR_TP_INVALID_CRL_TYPE:
		return "CSSMERR_TP_INVALID_CRL_TYPE"
	case CSSMERR_TP_INVALID_FORM_TYPE:
		return "CSSMERR_TP_INVALID_FORM_TYPE"
	case CSSMERR_TP_INVALID_ID:
		return "CSSMERR_TP_INVALID_ID"
	case CSSMERR_TP_INVALID_IDENTIFIER:
		return "CSSMERR_TP_INVALID_IDENTIFIER"
	case CSSMERR_TP_INVALID_IDENTIFIER_POINTER:
		return "CSSMERR_TP_INVALID_IDENTIFIER_POINTER"
	case CSSMERR_TP_INVALID_INDEX:
		return "CSSMERR_TP_INVALID_INDEX"
	case CSSMERR_TP_INVALID_KEYCACHE_HANDLE:
		return "CSSMERR_TP_INVALID_KEYCACHE_HANDLE"
	case CSSMERR_TP_INVALID_NAME:
		return "CSSMERR_TP_INVALID_NAME"
	case CSSMERR_TP_INVALID_POLICY_IDENTIFIERS:
		return "CSSMERR_TP_INVALID_POLICY_IDENTIFIERS"
	case CSSMERR_TP_INVALID_REASON:
		return "CSSMERR_TP_INVALID_REASON"
	case CSSMERR_TP_INVALID_REQUEST_INPUTS:
		return "CSSMERR_TP_INVALID_REQUEST_INPUTS"
	case CSSMERR_TP_INVALID_RESPONSE_VECTOR:
		return "CSSMERR_TP_INVALID_RESPONSE_VECTOR"
	case CSSMERR_TP_INVALID_SIGNATURE:
		return "CSSMERR_TP_INVALID_SIGNATURE"
	case CSSMERR_TP_INVALID_STOP_ON_POLICY:
		return "CSSMERR_TP_INVALID_STOP_ON_POLICY"
	case CSSMERR_TP_INVALID_TIMESTRING:
		return "CSSMERR_TP_INVALID_TIMESTRING"
	case CSSMERR_TP_INVALID_TUPLE:
		return "CSSMERR_TP_INVALID_TUPLE"
	case CSSMERR_TP_INVALID_TUPLEGROUP:
		return "CSSMERR_TP_INVALID_TUPLEGROUP"
	case CSSMERR_TP_INVALID_TUPLEGROUP_POINTER:
		return "CSSMERR_TP_INVALID_TUPLEGROUP_POINTER"
	case CSSMERR_TP_NOT_SIGNER:
		return "CSSMERR_TP_NOT_SIGNER"
	case CSSMERR_TP_NOT_TRUSTED:
		return "CSSMERR_TP_NOT_TRUSTED"
	case CSSMERR_TP_NO_DEFAULT_AUTHORITY:
		return "CSSMERR_TP_NO_DEFAULT_AUTHORITY"
	case CSSMERR_TP_REJECTED_FORM:
		return "CSSMERR_TP_REJECTED_FORM"
	case CSSMERR_TP_REQUEST_LOST:
		return "CSSMERR_TP_REQUEST_LOST"
	case CSSMERR_TP_REQUEST_REJECTED:
		return "CSSMERR_TP_REQUEST_REJECTED"
	case CSSMERR_TP_UNSUPPORTED_ADDR_TYPE:
		return "CSSMERR_TP_UNSUPPORTED_ADDR_TYPE"
	case CSSMERR_TP_UNSUPPORTED_SERVICE:
		return "CSSMERR_TP_UNSUPPORTED_SERVICE"
	case CSSMERR_TP_VERIFY_ACTION_FAILED:
		return "CSSMERR_TP_VERIFY_ACTION_FAILED"
	case CSSM_TP_BASE_TP_ERROR:
		return "CSSM_TP_BASE_TP_ERROR"
	default:
		return fmt.Sprintf("CssmTpBaseTpError(%d)", e)
	}
}

type CssmTpCertchangeNone uint32

const (
	CSSM_TP_CERTCHANGE_HOLD    CssmTpCertchangeNone = 0x2
	CSSM_TP_CERTCHANGE_NONE    CssmTpCertchangeNone = 0
	CSSM_TP_CERTCHANGE_RELEASE CssmTpCertchangeNone = 0x3
	CSSM_TP_CERTCHANGE_REVOKE  CssmTpCertchangeNone = 0x1
)

func (e CssmTpCertchangeNone) String() string {
	switch e {
	case CSSM_TP_CERTCHANGE_HOLD:
		return "CSSM_TP_CERTCHANGE_HOLD"
	case CSSM_TP_CERTCHANGE_NONE:
		return "CSSM_TP_CERTCHANGE_NONE"
	case CSSM_TP_CERTCHANGE_RELEASE:
		return "CSSM_TP_CERTCHANGE_RELEASE"
	case CSSM_TP_CERTCHANGE_REVOKE:
		return "CSSM_TP_CERTCHANGE_REVOKE"
	default:
		return fmt.Sprintf("CssmTpCertchangeNone(%d)", e)
	}
}

type CssmTpCertchangeReason uint32

const (
	CSSM_TP_CERTCHANGE_REASON_AFFILIATIONCHANGE   CssmTpCertchangeReason = 0x4
	CSSM_TP_CERTCHANGE_REASON_CACOMPROMISE        CssmTpCertchangeReason = 0x2
	CSSM_TP_CERTCHANGE_REASON_CEASEOPERATION      CssmTpCertchangeReason = 0x3
	CSSM_TP_CERTCHANGE_REASON_HOLDRELEASE         CssmTpCertchangeReason = 0x7
	CSSM_TP_CERTCHANGE_REASON_KEYCOMPROMISE       CssmTpCertchangeReason = 0x1
	CSSM_TP_CERTCHANGE_REASON_SUPERCEDED          CssmTpCertchangeReason = 0x5
	CSSM_TP_CERTCHANGE_REASON_SUSPECTEDCOMPROMISE CssmTpCertchangeReason = 0x6
	CSSM_TP_CERTCHANGE_REASON_UNKNOWN             CssmTpCertchangeReason = 0
)

func (e CssmTpCertchangeReason) String() string {
	switch e {
	case CSSM_TP_CERTCHANGE_REASON_AFFILIATIONCHANGE:
		return "CSSM_TP_CERTCHANGE_REASON_AFFILIATIONCHANGE"
	case CSSM_TP_CERTCHANGE_REASON_CACOMPROMISE:
		return "CSSM_TP_CERTCHANGE_REASON_CACOMPROMISE"
	case CSSM_TP_CERTCHANGE_REASON_CEASEOPERATION:
		return "CSSM_TP_CERTCHANGE_REASON_CEASEOPERATION"
	case CSSM_TP_CERTCHANGE_REASON_HOLDRELEASE:
		return "CSSM_TP_CERTCHANGE_REASON_HOLDRELEASE"
	case CSSM_TP_CERTCHANGE_REASON_KEYCOMPROMISE:
		return "CSSM_TP_CERTCHANGE_REASON_KEYCOMPROMISE"
	case CSSM_TP_CERTCHANGE_REASON_SUPERCEDED:
		return "CSSM_TP_CERTCHANGE_REASON_SUPERCEDED"
	case CSSM_TP_CERTCHANGE_REASON_SUSPECTEDCOMPROMISE:
		return "CSSM_TP_CERTCHANGE_REASON_SUSPECTEDCOMPROMISE"
	case CSSM_TP_CERTCHANGE_REASON_UNKNOWN:
		return "CSSM_TP_CERTCHANGE_REASON_UNKNOWN"
	default:
		return fmt.Sprintf("CssmTpCertchangeReason(%d)", e)
	}
}

type CssmTpCertchangeStatusUnknown uint32

const (
	CSSM_TP_CERTCHANGE_NOT_AUTHORIZED CssmTpCertchangeStatusUnknown = 0x5
	CSSM_TP_CERTCHANGE_OK             CssmTpCertchangeStatusUnknown = 0x1
	CSSM_TP_CERTCHANGE_OKWITHNEWTIME  CssmTpCertchangeStatusUnknown = 0x2
	CSSM_TP_CERTCHANGE_REJECTED       CssmTpCertchangeStatusUnknown = 0x4
	CSSM_TP_CERTCHANGE_STATUS_UNKNOWN CssmTpCertchangeStatusUnknown = 0
	CSSM_TP_CERTCHANGE_WRONGCA        CssmTpCertchangeStatusUnknown = 0x3
)

func (e CssmTpCertchangeStatusUnknown) String() string {
	switch e {
	case CSSM_TP_CERTCHANGE_NOT_AUTHORIZED:
		return "CSSM_TP_CERTCHANGE_NOT_AUTHORIZED"
	case CSSM_TP_CERTCHANGE_OK:
		return "CSSM_TP_CERTCHANGE_OK"
	case CSSM_TP_CERTCHANGE_OKWITHNEWTIME:
		return "CSSM_TP_CERTCHANGE_OKWITHNEWTIME"
	case CSSM_TP_CERTCHANGE_REJECTED:
		return "CSSM_TP_CERTCHANGE_REJECTED"
	case CSSM_TP_CERTCHANGE_STATUS_UNKNOWN:
		return "CSSM_TP_CERTCHANGE_STATUS_UNKNOWN"
	case CSSM_TP_CERTCHANGE_WRONGCA:
		return "CSSM_TP_CERTCHANGE_WRONGCA"
	default:
		return fmt.Sprintf("CssmTpCertchangeStatusUnknown(%d)", e)
	}
}

type CssmTpCertissue uint32

const (
	CSSM_TP_CERTISSUE_NOT_AUTHORIZED    CssmTpCertissue = 0x5
	CSSM_TP_CERTISSUE_OK                CssmTpCertissue = 0x1
	CSSM_TP_CERTISSUE_OKWITHCERTMODS    CssmTpCertissue = 0x2
	CSSM_TP_CERTISSUE_OKWITHSERVICEMODS CssmTpCertissue = 0x3
	CSSM_TP_CERTISSUE_REJECTED          CssmTpCertissue = 0x4
	CSSM_TP_CERTISSUE_STATUS_UNKNOWN    CssmTpCertissue = 0
	CSSM_TP_CERTISSUE_WILL_BE_REVOKED   CssmTpCertissue = 0x6
)

func (e CssmTpCertissue) String() string {
	switch e {
	case CSSM_TP_CERTISSUE_NOT_AUTHORIZED:
		return "CSSM_TP_CERTISSUE_NOT_AUTHORIZED"
	case CSSM_TP_CERTISSUE_OK:
		return "CSSM_TP_CERTISSUE_OK"
	case CSSM_TP_CERTISSUE_OKWITHCERTMODS:
		return "CSSM_TP_CERTISSUE_OKWITHCERTMODS"
	case CSSM_TP_CERTISSUE_OKWITHSERVICEMODS:
		return "CSSM_TP_CERTISSUE_OKWITHSERVICEMODS"
	case CSSM_TP_CERTISSUE_REJECTED:
		return "CSSM_TP_CERTISSUE_REJECTED"
	case CSSM_TP_CERTISSUE_STATUS_UNKNOWN:
		return "CSSM_TP_CERTISSUE_STATUS_UNKNOWN"
	case CSSM_TP_CERTISSUE_WILL_BE_REVOKED:
		return "CSSM_TP_CERTISSUE_WILL_BE_REVOKED"
	default:
		return fmt.Sprintf("CssmTpCertissue(%d)", e)
	}
}

type CssmTpCertnotarize uint32

const (
	CSSM_TP_CERTNOTARIZE_NOT_AUTHORIZED    CssmTpCertnotarize = 0x5
	CSSM_TP_CERTNOTARIZE_OK                CssmTpCertnotarize = 0x1
	CSSM_TP_CERTNOTARIZE_OKWITHOUTFIELDS   CssmTpCertnotarize = 0x2
	CSSM_TP_CERTNOTARIZE_OKWITHSERVICEMODS CssmTpCertnotarize = 0x3
	CSSM_TP_CERTNOTARIZE_REJECTED          CssmTpCertnotarize = 0x4
	CSSM_TP_CERTNOTARIZE_STATUS_UNKNOWN    CssmTpCertnotarize = 0
)

func (e CssmTpCertnotarize) String() string {
	switch e {
	case CSSM_TP_CERTNOTARIZE_NOT_AUTHORIZED:
		return "CSSM_TP_CERTNOTARIZE_NOT_AUTHORIZED"
	case CSSM_TP_CERTNOTARIZE_OK:
		return "CSSM_TP_CERTNOTARIZE_OK"
	case CSSM_TP_CERTNOTARIZE_OKWITHOUTFIELDS:
		return "CSSM_TP_CERTNOTARIZE_OKWITHOUTFIELDS"
	case CSSM_TP_CERTNOTARIZE_OKWITHSERVICEMODS:
		return "CSSM_TP_CERTNOTARIZE_OKWITHSERVICEMODS"
	case CSSM_TP_CERTNOTARIZE_REJECTED:
		return "CSSM_TP_CERTNOTARIZE_REJECTED"
	case CSSM_TP_CERTNOTARIZE_STATUS_UNKNOWN:
		return "CSSM_TP_CERTNOTARIZE_STATUS_UNKNOWN"
	default:
		return fmt.Sprintf("CssmTpCertnotarize(%d)", e)
	}
}

type CssmTpCertreclaim uint32

const (
	CSSM_TP_CERTRECLAIM_NOMATCH        CssmTpCertreclaim = 0x2
	CSSM_TP_CERTRECLAIM_NOT_AUTHORIZED CssmTpCertreclaim = 0x4
	CSSM_TP_CERTRECLAIM_OK             CssmTpCertreclaim = 0x1
	CSSM_TP_CERTRECLAIM_REJECTED       CssmTpCertreclaim = 0x3
	CSSM_TP_CERTRECLAIM_STATUS_UNKNOWN CssmTpCertreclaim = 0
)

func (e CssmTpCertreclaim) String() string {
	switch e {
	case CSSM_TP_CERTRECLAIM_NOMATCH:
		return "CSSM_TP_CERTRECLAIM_NOMATCH"
	case CSSM_TP_CERTRECLAIM_NOT_AUTHORIZED:
		return "CSSM_TP_CERTRECLAIM_NOT_AUTHORIZED"
	case CSSM_TP_CERTRECLAIM_OK:
		return "CSSM_TP_CERTRECLAIM_OK"
	case CSSM_TP_CERTRECLAIM_REJECTED:
		return "CSSM_TP_CERTRECLAIM_REJECTED"
	case CSSM_TP_CERTRECLAIM_STATUS_UNKNOWN:
		return "CSSM_TP_CERTRECLAIM_STATUS_UNKNOWN"
	default:
		return fmt.Sprintf("CssmTpCertreclaim(%d)", e)
	}
}

type CssmTpCertverify uint32

const (
	CSSM_TP_CERTVERIFY_EXPIRED                   CssmTpCertverify = 0x5
	CSSM_TP_CERTVERIFY_INVALID                   CssmTpCertverify = 0x2
	CSSM_TP_CERTVERIFY_INVALID_AUTHORITY         CssmTpCertverify = 0x7
	CSSM_TP_CERTVERIFY_INVALID_BASIC_CONSTRAINTS CssmTpCertverify = 0xd
	CSSM_TP_CERTVERIFY_INVALID_CERTGROUP         CssmTpCertverify = 0xa
	CSSM_TP_CERTVERIFY_INVALID_CERT_VALUE        CssmTpCertverify = 0x9
	CSSM_TP_CERTVERIFY_INVALID_CRL_DIST_PT       CssmTpCertverify = 0xe
	CSSM_TP_CERTVERIFY_INVALID_NAME_TREE         CssmTpCertverify = 0xf
	CSSM_TP_CERTVERIFY_INVALID_POLICY            CssmTpCertverify = 0xb
	CSSM_TP_CERTVERIFY_INVALID_POLICY_IDS        CssmTpCertverify = 0xc
	CSSM_TP_CERTVERIFY_INVALID_SIGNATURE         CssmTpCertverify = 0x8
	CSSM_TP_CERTVERIFY_NOT_VALID_YET             CssmTpCertverify = 0x6
	CSSM_TP_CERTVERIFY_REVOKED                   CssmTpCertverify = 0x3
	CSSM_TP_CERTVERIFY_SUSPENDED                 CssmTpCertverify = 0x4
	CSSM_TP_CERTVERIFY_UNKNOWN                   CssmTpCertverify = 0
	CSSM_TP_CERTVERIFY_UNKNOWN_CRITICAL_EXT      CssmTpCertverify = 0x10
	CSSM_TP_CERTVERIFY_VALID                     CssmTpCertverify = 0x1
)

func (e CssmTpCertverify) String() string {
	switch e {
	case CSSM_TP_CERTVERIFY_EXPIRED:
		return "CSSM_TP_CERTVERIFY_EXPIRED"
	case CSSM_TP_CERTVERIFY_INVALID:
		return "CSSM_TP_CERTVERIFY_INVALID"
	case CSSM_TP_CERTVERIFY_INVALID_AUTHORITY:
		return "CSSM_TP_CERTVERIFY_INVALID_AUTHORITY"
	case CSSM_TP_CERTVERIFY_INVALID_BASIC_CONSTRAINTS:
		return "CSSM_TP_CERTVERIFY_INVALID_BASIC_CONSTRAINTS"
	case CSSM_TP_CERTVERIFY_INVALID_CERTGROUP:
		return "CSSM_TP_CERTVERIFY_INVALID_CERTGROUP"
	case CSSM_TP_CERTVERIFY_INVALID_CERT_VALUE:
		return "CSSM_TP_CERTVERIFY_INVALID_CERT_VALUE"
	case CSSM_TP_CERTVERIFY_INVALID_CRL_DIST_PT:
		return "CSSM_TP_CERTVERIFY_INVALID_CRL_DIST_PT"
	case CSSM_TP_CERTVERIFY_INVALID_NAME_TREE:
		return "CSSM_TP_CERTVERIFY_INVALID_NAME_TREE"
	case CSSM_TP_CERTVERIFY_INVALID_POLICY:
		return "CSSM_TP_CERTVERIFY_INVALID_POLICY"
	case CSSM_TP_CERTVERIFY_INVALID_POLICY_IDS:
		return "CSSM_TP_CERTVERIFY_INVALID_POLICY_IDS"
	case CSSM_TP_CERTVERIFY_INVALID_SIGNATURE:
		return "CSSM_TP_CERTVERIFY_INVALID_SIGNATURE"
	case CSSM_TP_CERTVERIFY_NOT_VALID_YET:
		return "CSSM_TP_CERTVERIFY_NOT_VALID_YET"
	case CSSM_TP_CERTVERIFY_REVOKED:
		return "CSSM_TP_CERTVERIFY_REVOKED"
	case CSSM_TP_CERTVERIFY_SUSPENDED:
		return "CSSM_TP_CERTVERIFY_SUSPENDED"
	case CSSM_TP_CERTVERIFY_UNKNOWN:
		return "CSSM_TP_CERTVERIFY_UNKNOWN"
	case CSSM_TP_CERTVERIFY_UNKNOWN_CRITICAL_EXT:
		return "CSSM_TP_CERTVERIFY_UNKNOWN_CRITICAL_EXT"
	case CSSM_TP_CERTVERIFY_VALID:
		return "CSSM_TP_CERTVERIFY_VALID"
	default:
		return fmt.Sprintf("CssmTpCertverify(%d)", e)
	}
}

type CssmTpConfirm uint32

const (
	CSSM_TP_CONFIRM_ACCEPT         CssmTpConfirm = 0x1
	CSSM_TP_CONFIRM_REJECT         CssmTpConfirm = 0x2
	CSSM_TP_CONFIRM_STATUS_UNKNOWN CssmTpConfirm = 0
)

func (e CssmTpConfirm) String() string {
	switch e {
	case CSSM_TP_CONFIRM_ACCEPT:
		return "CSSM_TP_CONFIRM_ACCEPT"
	case CSSM_TP_CONFIRM_REJECT:
		return "CSSM_TP_CONFIRM_REJECT"
	case CSSM_TP_CONFIRM_STATUS_UNKNOWN:
		return "CSSM_TP_CONFIRM_STATUS_UNKNOWN"
	default:
		return fmt.Sprintf("CssmTpConfirm(%d)", e)
	}
}

type CssmTpCrlissue uint32

const (
	CSSM_TP_CRLISSUE_INVALID_DOMAIN     CssmTpCrlissue = 0x3
	CSSM_TP_CRLISSUE_NOT_AUTHORIZED     CssmTpCrlissue = 0x6
	CSSM_TP_CRLISSUE_NOT_CURRENT        CssmTpCrlissue = 0x2
	CSSM_TP_CRLISSUE_OK                 CssmTpCrlissue = 0x1
	CSSM_TP_CRLISSUE_REJECTED           CssmTpCrlissue = 0x5
	CSSM_TP_CRLISSUE_STATUS_UNKNOWN     CssmTpCrlissue = 0
	CSSM_TP_CRLISSUE_UNKNOWN_IDENTIFIER CssmTpCrlissue = 0x4
)

func (e CssmTpCrlissue) String() string {
	switch e {
	case CSSM_TP_CRLISSUE_INVALID_DOMAIN:
		return "CSSM_TP_CRLISSUE_INVALID_DOMAIN"
	case CSSM_TP_CRLISSUE_NOT_AUTHORIZED:
		return "CSSM_TP_CRLISSUE_NOT_AUTHORIZED"
	case CSSM_TP_CRLISSUE_NOT_CURRENT:
		return "CSSM_TP_CRLISSUE_NOT_CURRENT"
	case CSSM_TP_CRLISSUE_OK:
		return "CSSM_TP_CRLISSUE_OK"
	case CSSM_TP_CRLISSUE_REJECTED:
		return "CSSM_TP_CRLISSUE_REJECTED"
	case CSSM_TP_CRLISSUE_STATUS_UNKNOWN:
		return "CSSM_TP_CRLISSUE_STATUS_UNKNOWN"
	case CSSM_TP_CRLISSUE_UNKNOWN_IDENTIFIER:
		return "CSSM_TP_CRLISSUE_UNKNOWN_IDENTIFIER"
	default:
		return fmt.Sprintf("CssmTpCrlissue(%d)", e)
	}
}

type CssmTpFormType uint32

const (
	CSSM_TP_FORM_TYPE_GENERIC      CssmTpFormType = 0
	CSSM_TP_FORM_TYPE_REGISTRATION CssmTpFormType = 0x1
)

func (e CssmTpFormType) String() string {
	switch e {
	case CSSM_TP_FORM_TYPE_GENERIC:
		return "CSSM_TP_FORM_TYPE_GENERIC"
	case CSSM_TP_FORM_TYPE_REGISTRATION:
		return "CSSM_TP_FORM_TYPE_REGISTRATION"
	default:
		return fmt.Sprintf("CssmTpFormType(%d)", e)
	}
}

type CssmTpStopOn uint32

const (
	CSSM_TP_STOP_ON_FIRST_FAIL CssmTpStopOn = 3
	CSSM_TP_STOP_ON_FIRST_PASS CssmTpStopOn = 2
	CSSM_TP_STOP_ON_NONE       CssmTpStopOn = 1
	CSSM_TP_STOP_ON_POLICY     CssmTpStopOn = 0
)

func (e CssmTpStopOn) String() string {
	switch e {
	case CSSM_TP_STOP_ON_FIRST_FAIL:
		return "CSSM_TP_STOP_ON_FIRST_FAIL"
	case CSSM_TP_STOP_ON_FIRST_PASS:
		return "CSSM_TP_STOP_ON_FIRST_PASS"
	case CSSM_TP_STOP_ON_NONE:
		return "CSSM_TP_STOP_ON_NONE"
	case CSSM_TP_STOP_ON_POLICY:
		return "CSSM_TP_STOP_ON_POLICY"
	default:
		return fmt.Sprintf("CssmTpStopOn(%d)", e)
	}
}

type CssmUsee uint32

const (
	CSSM_USEE_AUTHENTICATION CssmUsee = 6
	CSSM_USEE_DOMESTIC       CssmUsee = 1
	CSSM_USEE_FINANCIAL      CssmUsee = 2
	CSSM_USEE_INSURANCE      CssmUsee = 9
	CSSM_USEE_KEYEXCH        CssmUsee = 7
	CSSM_USEE_KRENT          CssmUsee = 4
	CSSM_USEE_KRLE           CssmUsee = 3
	CSSM_USEE_LAST           CssmUsee = 0xff
	CSSM_USEE_MEDICAL        CssmUsee = 8
	CSSM_USEE_NONE           CssmUsee = 0
	CSSM_USEE_SSL            CssmUsee = 5
	CSSM_USEE_WEAK           CssmUsee = 10
)

func (e CssmUsee) String() string {
	switch e {
	case CSSM_USEE_AUTHENTICATION:
		return "CSSM_USEE_AUTHENTICATION"
	case CSSM_USEE_DOMESTIC:
		return "CSSM_USEE_DOMESTIC"
	case CSSM_USEE_FINANCIAL:
		return "CSSM_USEE_FINANCIAL"
	case CSSM_USEE_INSURANCE:
		return "CSSM_USEE_INSURANCE"
	case CSSM_USEE_KEYEXCH:
		return "CSSM_USEE_KEYEXCH"
	case CSSM_USEE_KRENT:
		return "CSSM_USEE_KRENT"
	case CSSM_USEE_KRLE:
		return "CSSM_USEE_KRLE"
	case CSSM_USEE_LAST:
		return "CSSM_USEE_LAST"
	case CSSM_USEE_MEDICAL:
		return "CSSM_USEE_MEDICAL"
	case CSSM_USEE_NONE:
		return "CSSM_USEE_NONE"
	case CSSM_USEE_SSL:
		return "CSSM_USEE_SSL"
	case CSSM_USEE_WEAK:
		return "CSSM_USEE_WEAK"
	default:
		return fmt.Sprintf("CssmUsee(%d)", e)
	}
}

type CssmValueNot int32

const (
	CSSM_VALUE_NOT_AVAILABLE CssmValueNot = -1
)

func (e CssmValueNot) String() string {
	switch e {
	case CSSM_VALUE_NOT_AVAILABLE:
		return "CSSM_VALUE_NOT_AVAILABLE"
	default:
		return fmt.Sprintf("CssmValueNot(%d)", e)
	}
}

type CssmWordidKeychainPrompt uint32

const (
	CSSM_WORDID_ASYMMETRIC_KEY       CssmWordidKeychainPrompt = 65547
	CSSM_WORDID_KEY                  CssmWordidKeychainPrompt = 65543
	CSSM_WORDID_KEYBAG_KEY           CssmWordidKeychainPrompt = 65549
	CSSM_WORDID_KEYCHAIN_CHANGE_LOCK CssmWordidKeychainPrompt = 65538
	CSSM_WORDID_KEYCHAIN_LOCK        CssmWordidKeychainPrompt = 65537
	CSSM_WORDID_KEYCHAIN_PROMPT      CssmWordidKeychainPrompt = 65536
	CSSM_WORDID_PARTITION            CssmWordidKeychainPrompt = 65548
	CSSM_WORDID_PIN                  CssmWordidKeychainPrompt = 65544
	CSSM_WORDID_PREAUTH              CssmWordidKeychainPrompt = 65545
	CSSM_WORDID_PREAUTH_SOURCE       CssmWordidKeychainPrompt = 65546
	CSSM_WORDID_PROCESS              CssmWordidKeychainPrompt = 65539
	CSSM_WORDID_SYMMETRIC_KEY        CssmWordidKeychainPrompt = 65541
	CSSM_WORDID_SYSTEM               CssmWordidKeychainPrompt = 65542
	CSSM_WORDID__FIRST_UNUSED        CssmWordidKeychainPrompt = 65550
	CSSM_WORDID__RESERVED_1          CssmWordidKeychainPrompt = 65540
)

func (e CssmWordidKeychainPrompt) String() string {
	switch e {
	case CSSM_WORDID_ASYMMETRIC_KEY:
		return "CSSM_WORDID_ASYMMETRIC_KEY"
	case CSSM_WORDID_KEY:
		return "CSSM_WORDID_KEY"
	case CSSM_WORDID_KEYBAG_KEY:
		return "CSSM_WORDID_KEYBAG_KEY"
	case CSSM_WORDID_KEYCHAIN_CHANGE_LOCK:
		return "CSSM_WORDID_KEYCHAIN_CHANGE_LOCK"
	case CSSM_WORDID_KEYCHAIN_LOCK:
		return "CSSM_WORDID_KEYCHAIN_LOCK"
	case CSSM_WORDID_KEYCHAIN_PROMPT:
		return "CSSM_WORDID_KEYCHAIN_PROMPT"
	case CSSM_WORDID_PARTITION:
		return "CSSM_WORDID_PARTITION"
	case CSSM_WORDID_PIN:
		return "CSSM_WORDID_PIN"
	case CSSM_WORDID_PREAUTH:
		return "CSSM_WORDID_PREAUTH"
	case CSSM_WORDID_PREAUTH_SOURCE:
		return "CSSM_WORDID_PREAUTH_SOURCE"
	case CSSM_WORDID_PROCESS:
		return "CSSM_WORDID_PROCESS"
	case CSSM_WORDID_SYMMETRIC_KEY:
		return "CSSM_WORDID_SYMMETRIC_KEY"
	case CSSM_WORDID_SYSTEM:
		return "CSSM_WORDID_SYSTEM"
	case CSSM_WORDID__FIRST_UNUSED:
		return "CSSM_WORDID__FIRST_UNUSED"
	case CSSM_WORDID__RESERVED_1:
		return "CSSM_WORDID__RESERVED_1"
	default:
		return fmt.Sprintf("CssmWordidKeychainPrompt(%d)", e)
	}
}

type CssmWordidUnk int32

const (
	CSSM_WORDID_A                    CssmWordidUnk = 2
	CSSM_WORDID_ACL                  CssmWordidUnk = 3
	CSSM_WORDID_ALPHA                CssmWordidUnk = 4
	CSSM_WORDID_B                    CssmWordidUnk = 5
	CSSM_WORDID_BER                  CssmWordidUnk = 6
	CSSM_WORDID_BINARY               CssmWordidUnk = 7
	CSSM_WORDID_BIOMETRIC            CssmWordidUnk = 8
	CSSM_WORDID_C                    CssmWordidUnk = 9
	CSSM_WORDID_CANCELED             CssmWordidUnk = 10
	CSSM_WORDID_CERT                 CssmWordidUnk = 11
	CSSM_WORDID_COMMENT              CssmWordidUnk = 12
	CSSM_WORDID_CRL                  CssmWordidUnk = 13
	CSSM_WORDID_CUSTOM               CssmWordidUnk = 14
	CSSM_WORDID_D                    CssmWordidUnk = 15
	CSSM_WORDID_DATE                 CssmWordidUnk = 16
	CSSM_WORDID_DBS_CREATE           CssmWordidUnk = 22
	CSSM_WORDID_DBS_DELETE           CssmWordidUnk = 23
	CSSM_WORDID_DB_DELETE            CssmWordidUnk = 17
	CSSM_WORDID_DB_EXEC_STORED_QUERY CssmWordidUnk = 18
	CSSM_WORDID_DB_INSERT            CssmWordidUnk = 19
	CSSM_WORDID_DB_MODIFY            CssmWordidUnk = 20
	CSSM_WORDID_DB_READ              CssmWordidUnk = 21
	CSSM_WORDID_DECRYPT              CssmWordidUnk = 24
	CSSM_WORDID_DELETE               CssmWordidUnk = 25
	CSSM_WORDID_DELTA_CRL            CssmWordidUnk = 26
	CSSM_WORDID_DER                  CssmWordidUnk = 27
	CSSM_WORDID_DERIVE               CssmWordidUnk = 28
	CSSM_WORDID_DISPLAY              CssmWordidUnk = 29
	CSSM_WORDID_DO                   CssmWordidUnk = 30
	CSSM_WORDID_DSA                  CssmWordidUnk = 31
	CSSM_WORDID_DSA_SHA1             CssmWordidUnk = 32
	CSSM_WORDID_E                    CssmWordidUnk = 33
	CSSM_WORDID_ELGAMAL              CssmWordidUnk = 34
	CSSM_WORDID_ENCRYPT              CssmWordidUnk = 35
	CSSM_WORDID_ENTRY                CssmWordidUnk = 36
	CSSM_WORDID_EXPORT_CLEAR         CssmWordidUnk = 37
	CSSM_WORDID_EXPORT_WRAPPED       CssmWordidUnk = 38
	CSSM_WORDID_G                    CssmWordidUnk = 39
	CSSM_WORDID_GE                   CssmWordidUnk = 40
	CSSM_WORDID_GENKEY               CssmWordidUnk = 41
	CSSM_WORDID_HASH                 CssmWordidUnk = 42
	CSSM_WORDID_HASHED_PASSWORD      CssmWordidUnk = 43
	CSSM_WORDID_HASHED_SUBJECT       CssmWordidUnk = 44
	CSSM_WORDID_HAVAL                CssmWordidUnk = 45
	CSSM_WORDID_IBCHASH              CssmWordidUnk = 46
	CSSM_WORDID_IMPORT_CLEAR         CssmWordidUnk = 47
	CSSM_WORDID_IMPORT_WRAPPED       CssmWordidUnk = 48
	CSSM_WORDID_INTEL                CssmWordidUnk = 49
	CSSM_WORDID_ISSUER               CssmWordidUnk = 50
	CSSM_WORDID_ISSUER_INFO          CssmWordidUnk = 51
	CSSM_WORDID_KEA                  CssmWordidUnk = 53
	CSSM_WORDID_KEYHOLDER            CssmWordidUnk = 54
	CSSM_WORDID_K_OF_N               CssmWordidUnk = 52
	CSSM_WORDID_L                    CssmWordidUnk = 55
	CSSM_WORDID_LE                   CssmWordidUnk = 56
	CSSM_WORDID_LOGIN                CssmWordidUnk = 57
	CSSM_WORDID_LOGIN_NAME           CssmWordidUnk = 58
	CSSM_WORDID_MAC                  CssmWordidUnk = 59
	CSSM_WORDID_MD2                  CssmWordidUnk = 60
	CSSM_WORDID_MD2WITHRSA           CssmWordidUnk = 61
	CSSM_WORDID_MD4                  CssmWordidUnk = 62
	CSSM_WORDID_MD5                  CssmWordidUnk = 63
	CSSM_WORDID_MD5WITHRSA           CssmWordidUnk = 64
	CSSM_WORDID_N                    CssmWordidUnk = 65
	CSSM_WORDID_NAME                 CssmWordidUnk = 66
	CSSM_WORDID_NDR                  CssmWordidUnk = 67
	CSSM_WORDID_NHASH                CssmWordidUnk = 68
	CSSM_WORDID_NOT_AFTER            CssmWordidUnk = 69
	CSSM_WORDID_NOT_BEFORE           CssmWordidUnk = 70
	CSSM_WORDID_NULL                 CssmWordidUnk = 71
	CSSM_WORDID_NUMERIC              CssmWordidUnk = 72
	CSSM_WORDID_OBJECT_HASH          CssmWordidUnk = 73
	CSSM_WORDID_ONE_TIME             CssmWordidUnk = 74
	CSSM_WORDID_ONLINE               CssmWordidUnk = 75
	CSSM_WORDID_OWNER                CssmWordidUnk = 76
	CSSM_WORDID_P                    CssmWordidUnk = 77
	CSSM_WORDID_PAM_NAME             CssmWordidUnk = 78
	CSSM_WORDID_PASSWORD             CssmWordidUnk = 79
	CSSM_WORDID_PGP                  CssmWordidUnk = 80
	CSSM_WORDID_PREFIX               CssmWordidUnk = 81
	CSSM_WORDID_PRIVATE_KEY          CssmWordidUnk = 82
	CSSM_WORDID_PROMPTED_BIOMETRIC   CssmWordidUnk = 83
	CSSM_WORDID_PROMPTED_PASSWORD    CssmWordidUnk = 84
	CSSM_WORDID_PROPAGATE            CssmWordidUnk = 85
	CSSM_WORDID_PROTECTED_BIOMETRIC  CssmWordidUnk = 86
	CSSM_WORDID_PROTECTED_PASSWORD   CssmWordidUnk = 87
	CSSM_WORDID_PROTECTED_PIN        CssmWordidUnk = 88
	CSSM_WORDID_PUBLIC_KEY           CssmWordidUnk = 89
	CSSM_WORDID_PUBLIC_KEY_FROM_CERT CssmWordidUnk = 90
	CSSM_WORDID_Q                    CssmWordidUnk = 91
	CSSM_WORDID_RANGE                CssmWordidUnk = 92
	CSSM_WORDID_REVAL                CssmWordidUnk = 93
	CSSM_WORDID_RIPEMAC              CssmWordidUnk = 94
	CSSM_WORDID_RIPEMD               CssmWordidUnk = 95
	CSSM_WORDID_RIPEMD160            CssmWordidUnk = 96
	CSSM_WORDID_RSA                  CssmWordidUnk = 97
	CSSM_WORDID_RSA_ISO9796          CssmWordidUnk = 98
	CSSM_WORDID_RSA_PKCS             CssmWordidUnk = 99
	CSSM_WORDID_RSA_PKCS1            CssmWordidUnk = 102
	CSSM_WORDID_RSA_PKCS1_MD5        CssmWordidUnk = 103
	CSSM_WORDID_RSA_PKCS1_SHA1       CssmWordidUnk = 104
	CSSM_WORDID_RSA_PKCS1_SIG        CssmWordidUnk = 105
	CSSM_WORDID_RSA_PKCS_MD5         CssmWordidUnk = 100
	CSSM_WORDID_RSA_PKCS_SHA1        CssmWordidUnk = 101
	CSSM_WORDID_RSA_RAW              CssmWordidUnk = 106
	CSSM_WORDID_SDSIV1               CssmWordidUnk = 107
	CSSM_WORDID_SEQUENCE             CssmWordidUnk = 108
	CSSM_WORDID_SET                  CssmWordidUnk = 109
	CSSM_WORDID_SEXPR                CssmWordidUnk = 110
	CSSM_WORDID_SHA1                 CssmWordidUnk = 111
	CSSM_WORDID_SHA1WITHDSA          CssmWordidUnk = 112
	CSSM_WORDID_SHA1WITHECDSA        CssmWordidUnk = 113
	CSSM_WORDID_SHA1WITHRSA          CssmWordidUnk = 114
	CSSM_WORDID_SIGN                 CssmWordidUnk = 115
	CSSM_WORDID_SIGNATURE            CssmWordidUnk = 116
	CSSM_WORDID_SIGNED_NONCE         CssmWordidUnk = 117
	CSSM_WORDID_SIGNED_SECRET        CssmWordidUnk = 118
	CSSM_WORDID_SPKI                 CssmWordidUnk = 119
	CSSM_WORDID_SUBJECT              CssmWordidUnk = 120
	CSSM_WORDID_SUBJECT_INFO         CssmWordidUnk = 121
	CSSM_WORDID_TAG                  CssmWordidUnk = 122
	CSSM_WORDID_THRESHOLD            CssmWordidUnk = 123
	CSSM_WORDID_TIME                 CssmWordidUnk = 124
	CSSM_WORDID_URI                  CssmWordidUnk = 125
	CSSM_WORDID_VENDOR_END           CssmWordidUnk = 0x7fff0000
	CSSM_WORDID_VENDOR_START         CssmWordidUnk = 0x10000
	CSSM_WORDID_VERSION              CssmWordidUnk = 126
	CSSM_WORDID_X509V1               CssmWordidUnk = 128
	CSSM_WORDID_X509V2               CssmWordidUnk = 129
	CSSM_WORDID_X509V3               CssmWordidUnk = 130
	CSSM_WORDID_X509_ATTRIBUTE       CssmWordidUnk = 127
	CSSM_WORDID_X9_ATTRIBUTE         CssmWordidUnk = 131
	CSSM_WORDID__NLU_                CssmWordidUnk = 0
	CSSM_WORDID__STAR_               CssmWordidUnk = 1
	CSSM_WORDID__UNK_                CssmWordidUnk = -1
)

func (e CssmWordidUnk) String() string {
	switch e {
	case CSSM_WORDID_A:
		return "CSSM_WORDID_A"
	case CSSM_WORDID_ACL:
		return "CSSM_WORDID_ACL"
	case CSSM_WORDID_ALPHA:
		return "CSSM_WORDID_ALPHA"
	case CSSM_WORDID_B:
		return "CSSM_WORDID_B"
	case CSSM_WORDID_BER:
		return "CSSM_WORDID_BER"
	case CSSM_WORDID_BINARY:
		return "CSSM_WORDID_BINARY"
	case CSSM_WORDID_BIOMETRIC:
		return "CSSM_WORDID_BIOMETRIC"
	case CSSM_WORDID_C:
		return "CSSM_WORDID_C"
	case CSSM_WORDID_CANCELED:
		return "CSSM_WORDID_CANCELED"
	case CSSM_WORDID_CERT:
		return "CSSM_WORDID_CERT"
	case CSSM_WORDID_COMMENT:
		return "CSSM_WORDID_COMMENT"
	case CSSM_WORDID_CRL:
		return "CSSM_WORDID_CRL"
	case CSSM_WORDID_CUSTOM:
		return "CSSM_WORDID_CUSTOM"
	case CSSM_WORDID_D:
		return "CSSM_WORDID_D"
	case CSSM_WORDID_DATE:
		return "CSSM_WORDID_DATE"
	case CSSM_WORDID_DBS_CREATE:
		return "CSSM_WORDID_DBS_CREATE"
	case CSSM_WORDID_DBS_DELETE:
		return "CSSM_WORDID_DBS_DELETE"
	case CSSM_WORDID_DB_DELETE:
		return "CSSM_WORDID_DB_DELETE"
	case CSSM_WORDID_DB_EXEC_STORED_QUERY:
		return "CSSM_WORDID_DB_EXEC_STORED_QUERY"
	case CSSM_WORDID_DB_INSERT:
		return "CSSM_WORDID_DB_INSERT"
	case CSSM_WORDID_DB_MODIFY:
		return "CSSM_WORDID_DB_MODIFY"
	case CSSM_WORDID_DB_READ:
		return "CSSM_WORDID_DB_READ"
	case CSSM_WORDID_DECRYPT:
		return "CSSM_WORDID_DECRYPT"
	case CSSM_WORDID_DELETE:
		return "CSSM_WORDID_DELETE"
	case CSSM_WORDID_DELTA_CRL:
		return "CSSM_WORDID_DELTA_CRL"
	case CSSM_WORDID_DER:
		return "CSSM_WORDID_DER"
	case CSSM_WORDID_DERIVE:
		return "CSSM_WORDID_DERIVE"
	case CSSM_WORDID_DISPLAY:
		return "CSSM_WORDID_DISPLAY"
	case CSSM_WORDID_DO:
		return "CSSM_WORDID_DO"
	case CSSM_WORDID_DSA:
		return "CSSM_WORDID_DSA"
	case CSSM_WORDID_DSA_SHA1:
		return "CSSM_WORDID_DSA_SHA1"
	case CSSM_WORDID_E:
		return "CSSM_WORDID_E"
	case CSSM_WORDID_ELGAMAL:
		return "CSSM_WORDID_ELGAMAL"
	case CSSM_WORDID_ENCRYPT:
		return "CSSM_WORDID_ENCRYPT"
	case CSSM_WORDID_ENTRY:
		return "CSSM_WORDID_ENTRY"
	case CSSM_WORDID_EXPORT_CLEAR:
		return "CSSM_WORDID_EXPORT_CLEAR"
	case CSSM_WORDID_EXPORT_WRAPPED:
		return "CSSM_WORDID_EXPORT_WRAPPED"
	case CSSM_WORDID_G:
		return "CSSM_WORDID_G"
	case CSSM_WORDID_GE:
		return "CSSM_WORDID_GE"
	case CSSM_WORDID_GENKEY:
		return "CSSM_WORDID_GENKEY"
	case CSSM_WORDID_HASH:
		return "CSSM_WORDID_HASH"
	case CSSM_WORDID_HASHED_PASSWORD:
		return "CSSM_WORDID_HASHED_PASSWORD"
	case CSSM_WORDID_HASHED_SUBJECT:
		return "CSSM_WORDID_HASHED_SUBJECT"
	case CSSM_WORDID_HAVAL:
		return "CSSM_WORDID_HAVAL"
	case CSSM_WORDID_IBCHASH:
		return "CSSM_WORDID_IBCHASH"
	case CSSM_WORDID_IMPORT_CLEAR:
		return "CSSM_WORDID_IMPORT_CLEAR"
	case CSSM_WORDID_IMPORT_WRAPPED:
		return "CSSM_WORDID_IMPORT_WRAPPED"
	case CSSM_WORDID_INTEL:
		return "CSSM_WORDID_INTEL"
	case CSSM_WORDID_ISSUER:
		return "CSSM_WORDID_ISSUER"
	case CSSM_WORDID_ISSUER_INFO:
		return "CSSM_WORDID_ISSUER_INFO"
	case CSSM_WORDID_KEA:
		return "CSSM_WORDID_KEA"
	case CSSM_WORDID_KEYHOLDER:
		return "CSSM_WORDID_KEYHOLDER"
	case CSSM_WORDID_K_OF_N:
		return "CSSM_WORDID_K_OF_N"
	case CSSM_WORDID_L:
		return "CSSM_WORDID_L"
	case CSSM_WORDID_LE:
		return "CSSM_WORDID_LE"
	case CSSM_WORDID_LOGIN:
		return "CSSM_WORDID_LOGIN"
	case CSSM_WORDID_LOGIN_NAME:
		return "CSSM_WORDID_LOGIN_NAME"
	case CSSM_WORDID_MAC:
		return "CSSM_WORDID_MAC"
	case CSSM_WORDID_MD2:
		return "CSSM_WORDID_MD2"
	case CSSM_WORDID_MD2WITHRSA:
		return "CSSM_WORDID_MD2WITHRSA"
	case CSSM_WORDID_MD4:
		return "CSSM_WORDID_MD4"
	case CSSM_WORDID_MD5:
		return "CSSM_WORDID_MD5"
	case CSSM_WORDID_MD5WITHRSA:
		return "CSSM_WORDID_MD5WITHRSA"
	case CSSM_WORDID_N:
		return "CSSM_WORDID_N"
	case CSSM_WORDID_NAME:
		return "CSSM_WORDID_NAME"
	case CSSM_WORDID_NDR:
		return "CSSM_WORDID_NDR"
	case CSSM_WORDID_NHASH:
		return "CSSM_WORDID_NHASH"
	case CSSM_WORDID_NOT_AFTER:
		return "CSSM_WORDID_NOT_AFTER"
	case CSSM_WORDID_NOT_BEFORE:
		return "CSSM_WORDID_NOT_BEFORE"
	case CSSM_WORDID_NULL:
		return "CSSM_WORDID_NULL"
	case CSSM_WORDID_NUMERIC:
		return "CSSM_WORDID_NUMERIC"
	case CSSM_WORDID_OBJECT_HASH:
		return "CSSM_WORDID_OBJECT_HASH"
	case CSSM_WORDID_ONE_TIME:
		return "CSSM_WORDID_ONE_TIME"
	case CSSM_WORDID_ONLINE:
		return "CSSM_WORDID_ONLINE"
	case CSSM_WORDID_OWNER:
		return "CSSM_WORDID_OWNER"
	case CSSM_WORDID_P:
		return "CSSM_WORDID_P"
	case CSSM_WORDID_PAM_NAME:
		return "CSSM_WORDID_PAM_NAME"
	case CSSM_WORDID_PASSWORD:
		return "CSSM_WORDID_PASSWORD"
	case CSSM_WORDID_PGP:
		return "CSSM_WORDID_PGP"
	case CSSM_WORDID_PREFIX:
		return "CSSM_WORDID_PREFIX"
	case CSSM_WORDID_PRIVATE_KEY:
		return "CSSM_WORDID_PRIVATE_KEY"
	case CSSM_WORDID_PROMPTED_BIOMETRIC:
		return "CSSM_WORDID_PROMPTED_BIOMETRIC"
	case CSSM_WORDID_PROMPTED_PASSWORD:
		return "CSSM_WORDID_PROMPTED_PASSWORD"
	case CSSM_WORDID_PROPAGATE:
		return "CSSM_WORDID_PROPAGATE"
	case CSSM_WORDID_PROTECTED_BIOMETRIC:
		return "CSSM_WORDID_PROTECTED_BIOMETRIC"
	case CSSM_WORDID_PROTECTED_PASSWORD:
		return "CSSM_WORDID_PROTECTED_PASSWORD"
	case CSSM_WORDID_PROTECTED_PIN:
		return "CSSM_WORDID_PROTECTED_PIN"
	case CSSM_WORDID_PUBLIC_KEY:
		return "CSSM_WORDID_PUBLIC_KEY"
	case CSSM_WORDID_PUBLIC_KEY_FROM_CERT:
		return "CSSM_WORDID_PUBLIC_KEY_FROM_CERT"
	case CSSM_WORDID_Q:
		return "CSSM_WORDID_Q"
	case CSSM_WORDID_RANGE:
		return "CSSM_WORDID_RANGE"
	case CSSM_WORDID_REVAL:
		return "CSSM_WORDID_REVAL"
	case CSSM_WORDID_RIPEMAC:
		return "CSSM_WORDID_RIPEMAC"
	case CSSM_WORDID_RIPEMD:
		return "CSSM_WORDID_RIPEMD"
	case CSSM_WORDID_RIPEMD160:
		return "CSSM_WORDID_RIPEMD160"
	case CSSM_WORDID_RSA:
		return "CSSM_WORDID_RSA"
	case CSSM_WORDID_RSA_ISO9796:
		return "CSSM_WORDID_RSA_ISO9796"
	case CSSM_WORDID_RSA_PKCS:
		return "CSSM_WORDID_RSA_PKCS"
	case CSSM_WORDID_RSA_PKCS1:
		return "CSSM_WORDID_RSA_PKCS1"
	case CSSM_WORDID_RSA_PKCS1_MD5:
		return "CSSM_WORDID_RSA_PKCS1_MD5"
	case CSSM_WORDID_RSA_PKCS1_SHA1:
		return "CSSM_WORDID_RSA_PKCS1_SHA1"
	case CSSM_WORDID_RSA_PKCS1_SIG:
		return "CSSM_WORDID_RSA_PKCS1_SIG"
	case CSSM_WORDID_RSA_PKCS_MD5:
		return "CSSM_WORDID_RSA_PKCS_MD5"
	case CSSM_WORDID_RSA_PKCS_SHA1:
		return "CSSM_WORDID_RSA_PKCS_SHA1"
	case CSSM_WORDID_RSA_RAW:
		return "CSSM_WORDID_RSA_RAW"
	case CSSM_WORDID_SDSIV1:
		return "CSSM_WORDID_SDSIV1"
	case CSSM_WORDID_SEQUENCE:
		return "CSSM_WORDID_SEQUENCE"
	case CSSM_WORDID_SET:
		return "CSSM_WORDID_SET"
	case CSSM_WORDID_SEXPR:
		return "CSSM_WORDID_SEXPR"
	case CSSM_WORDID_SHA1:
		return "CSSM_WORDID_SHA1"
	case CSSM_WORDID_SHA1WITHDSA:
		return "CSSM_WORDID_SHA1WITHDSA"
	case CSSM_WORDID_SHA1WITHECDSA:
		return "CSSM_WORDID_SHA1WITHECDSA"
	case CSSM_WORDID_SHA1WITHRSA:
		return "CSSM_WORDID_SHA1WITHRSA"
	case CSSM_WORDID_SIGN:
		return "CSSM_WORDID_SIGN"
	case CSSM_WORDID_SIGNATURE:
		return "CSSM_WORDID_SIGNATURE"
	case CSSM_WORDID_SIGNED_NONCE:
		return "CSSM_WORDID_SIGNED_NONCE"
	case CSSM_WORDID_SIGNED_SECRET:
		return "CSSM_WORDID_SIGNED_SECRET"
	case CSSM_WORDID_SPKI:
		return "CSSM_WORDID_SPKI"
	case CSSM_WORDID_SUBJECT:
		return "CSSM_WORDID_SUBJECT"
	case CSSM_WORDID_SUBJECT_INFO:
		return "CSSM_WORDID_SUBJECT_INFO"
	case CSSM_WORDID_TAG:
		return "CSSM_WORDID_TAG"
	case CSSM_WORDID_THRESHOLD:
		return "CSSM_WORDID_THRESHOLD"
	case CSSM_WORDID_TIME:
		return "CSSM_WORDID_TIME"
	case CSSM_WORDID_URI:
		return "CSSM_WORDID_URI"
	case CSSM_WORDID_VENDOR_END:
		return "CSSM_WORDID_VENDOR_END"
	case CSSM_WORDID_VENDOR_START:
		return "CSSM_WORDID_VENDOR_START"
	case CSSM_WORDID_VERSION:
		return "CSSM_WORDID_VERSION"
	case CSSM_WORDID_X509V1:
		return "CSSM_WORDID_X509V1"
	case CSSM_WORDID_X509V2:
		return "CSSM_WORDID_X509V2"
	case CSSM_WORDID_X509V3:
		return "CSSM_WORDID_X509V3"
	case CSSM_WORDID_X509_ATTRIBUTE:
		return "CSSM_WORDID_X509_ATTRIBUTE"
	case CSSM_WORDID_X9_ATTRIBUTE:
		return "CSSM_WORDID_X9_ATTRIBUTE"
	case CSSM_WORDID__NLU_:
		return "CSSM_WORDID__NLU_"
	case CSSM_WORDID__STAR_:
		return "CSSM_WORDID__STAR_"
	case CSSM_WORDID__UNK_:
		return "CSSM_WORDID__UNK_"
	default:
		return fmt.Sprintf("CssmWordidUnk(%d)", e)
	}
}

type Cssmerr int32

const (
	CSSMERR_AC_DEVICE_FAILED                        Cssmerr = -2147405595
	CSSMERR_AC_DEVICE_RESET                         Cssmerr = -2147405596
	CSSMERR_AC_INSUFFICIENT_CLIENT_IDENTIFICATION   Cssmerr = -2147405597
	CSSMERR_AC_IN_DARK_WAKE                         Cssmerr = -2147405594
	CSSMERR_AC_NO_USER_INTERACTION                  Cssmerr = -2147405600
	CSSMERR_AC_SERVICE_NOT_AVAILABLE                Cssmerr = -2147405598
	CSSMERR_AC_USER_CANCELED                        Cssmerr = -2147405599
	CSSMERR_CL_DEVICE_FAILED                        Cssmerr = -2147411739
	CSSMERR_CL_DEVICE_RESET                         Cssmerr = -2147411740
	CSSMERR_CL_INSUFFICIENT_CLIENT_IDENTIFICATION   Cssmerr = -2147411741
	CSSMERR_CL_IN_DARK_WAKE                         Cssmerr = -2147411738
	CSSMERR_CL_NO_USER_INTERACTION                  Cssmerr = -2147411744
	CSSMERR_CL_SERVICE_NOT_AVAILABLE                Cssmerr = -2147411742
	CSSMERR_CL_USER_CANCELED                        Cssmerr = -2147411743
	CSSMERR_CSP_DEVICE_FAILED                       Cssmerr = -2147415835
	CSSMERR_CSP_DEVICE_RESET                        Cssmerr = -2147415836
	CSSMERR_CSP_INSUFFICIENT_CLIENT_IDENTIFICATION  Cssmerr = -2147415837
	CSSMERR_CSP_IN_DARK_WAKE                        Cssmerr = -2147415834
	CSSMERR_CSP_NO_USER_INTERACTION                 Cssmerr = -2147415840
	CSSMERR_CSP_SERVICE_NOT_AVAILABLE               Cssmerr = -2147415838
	CSSMERR_CSP_USER_CANCELED                       Cssmerr = -2147415839
	CSSMERR_CSSM_DEVICE_FAILED                      Cssmerr = -2147417883
	CSSMERR_CSSM_DEVICE_RESET                       Cssmerr = -2147417884
	CSSMERR_CSSM_INSUFFICIENT_CLIENT_IDENTIFICATION Cssmerr = -2147417885
	CSSMERR_CSSM_IN_DARK_WAKE                       Cssmerr = -2147417882
	CSSMERR_CSSM_NO_USER_INTERACTION                Cssmerr = -2147417888
	CSSMERR_CSSM_SERVICE_NOT_AVAILABLE              Cssmerr = -2147417886
	CSSMERR_CSSM_USER_CANCELED                      Cssmerr = -2147417887
	CSSMERR_DL_DEVICE_FAILED                        Cssmerr = -2147413787
	CSSMERR_DL_DEVICE_RESET                         Cssmerr = -2147413788
	CSSMERR_DL_INSUFFICIENT_CLIENT_IDENTIFICATION   Cssmerr = -2147413789
	CSSMERR_DL_IN_DARK_WAKE                         Cssmerr = -2147413786
	CSSMERR_DL_NO_USER_INTERACTION                  Cssmerr = -2147413792
	CSSMERR_DL_SERVICE_NOT_AVAILABLE                Cssmerr = -2147413790
	CSSMERR_DL_USER_CANCELED                        Cssmerr = -2147413791
	CSSMERR_TP_DEVICE_FAILED                        Cssmerr = -2147409691
	CSSMERR_TP_DEVICE_RESET                         Cssmerr = -2147409692
	CSSMERR_TP_INSUFFICIENT_CLIENT_IDENTIFICATION   Cssmerr = -2147409693
	CSSMERR_TP_IN_DARK_WAKE                         Cssmerr = -2147409690
	CSSMERR_TP_NO_USER_INTERACTION                  Cssmerr = -2147409696
	CSSMERR_TP_SERVICE_NOT_AVAILABLE                Cssmerr = -2147409694
	CSSMERR_TP_USER_CANCELED                        Cssmerr = -2147409695
)

func (e Cssmerr) String() string {
	switch e {
	case CSSMERR_AC_DEVICE_FAILED:
		return "CSSMERR_AC_DEVICE_FAILED"
	case CSSMERR_AC_DEVICE_RESET:
		return "CSSMERR_AC_DEVICE_RESET"
	case CSSMERR_AC_INSUFFICIENT_CLIENT_IDENTIFICATION:
		return "CSSMERR_AC_INSUFFICIENT_CLIENT_IDENTIFICATION"
	case CSSMERR_AC_IN_DARK_WAKE:
		return "CSSMERR_AC_IN_DARK_WAKE"
	case CSSMERR_AC_NO_USER_INTERACTION:
		return "CSSMERR_AC_NO_USER_INTERACTION"
	case CSSMERR_AC_SERVICE_NOT_AVAILABLE:
		return "CSSMERR_AC_SERVICE_NOT_AVAILABLE"
	case CSSMERR_AC_USER_CANCELED:
		return "CSSMERR_AC_USER_CANCELED"
	case CSSMERR_CL_DEVICE_FAILED:
		return "CSSMERR_CL_DEVICE_FAILED"
	case CSSMERR_CL_DEVICE_RESET:
		return "CSSMERR_CL_DEVICE_RESET"
	case CSSMERR_CL_INSUFFICIENT_CLIENT_IDENTIFICATION:
		return "CSSMERR_CL_INSUFFICIENT_CLIENT_IDENTIFICATION"
	case CSSMERR_CL_IN_DARK_WAKE:
		return "CSSMERR_CL_IN_DARK_WAKE"
	case CSSMERR_CL_NO_USER_INTERACTION:
		return "CSSMERR_CL_NO_USER_INTERACTION"
	case CSSMERR_CL_SERVICE_NOT_AVAILABLE:
		return "CSSMERR_CL_SERVICE_NOT_AVAILABLE"
	case CSSMERR_CL_USER_CANCELED:
		return "CSSMERR_CL_USER_CANCELED"
	case CSSMERR_CSP_DEVICE_FAILED:
		return "CSSMERR_CSP_DEVICE_FAILED"
	case CSSMERR_CSP_DEVICE_RESET:
		return "CSSMERR_CSP_DEVICE_RESET"
	case CSSMERR_CSP_INSUFFICIENT_CLIENT_IDENTIFICATION:
		return "CSSMERR_CSP_INSUFFICIENT_CLIENT_IDENTIFICATION"
	case CSSMERR_CSP_IN_DARK_WAKE:
		return "CSSMERR_CSP_IN_DARK_WAKE"
	case CSSMERR_CSP_NO_USER_INTERACTION:
		return "CSSMERR_CSP_NO_USER_INTERACTION"
	case CSSMERR_CSP_SERVICE_NOT_AVAILABLE:
		return "CSSMERR_CSP_SERVICE_NOT_AVAILABLE"
	case CSSMERR_CSP_USER_CANCELED:
		return "CSSMERR_CSP_USER_CANCELED"
	case CSSMERR_CSSM_DEVICE_FAILED:
		return "CSSMERR_CSSM_DEVICE_FAILED"
	case CSSMERR_CSSM_DEVICE_RESET:
		return "CSSMERR_CSSM_DEVICE_RESET"
	case CSSMERR_CSSM_INSUFFICIENT_CLIENT_IDENTIFICATION:
		return "CSSMERR_CSSM_INSUFFICIENT_CLIENT_IDENTIFICATION"
	case CSSMERR_CSSM_IN_DARK_WAKE:
		return "CSSMERR_CSSM_IN_DARK_WAKE"
	case CSSMERR_CSSM_NO_USER_INTERACTION:
		return "CSSMERR_CSSM_NO_USER_INTERACTION"
	case CSSMERR_CSSM_SERVICE_NOT_AVAILABLE:
		return "CSSMERR_CSSM_SERVICE_NOT_AVAILABLE"
	case CSSMERR_CSSM_USER_CANCELED:
		return "CSSMERR_CSSM_USER_CANCELED"
	case CSSMERR_DL_DEVICE_FAILED:
		return "CSSMERR_DL_DEVICE_FAILED"
	case CSSMERR_DL_DEVICE_RESET:
		return "CSSMERR_DL_DEVICE_RESET"
	case CSSMERR_DL_INSUFFICIENT_CLIENT_IDENTIFICATION:
		return "CSSMERR_DL_INSUFFICIENT_CLIENT_IDENTIFICATION"
	case CSSMERR_DL_IN_DARK_WAKE:
		return "CSSMERR_DL_IN_DARK_WAKE"
	case CSSMERR_DL_NO_USER_INTERACTION:
		return "CSSMERR_DL_NO_USER_INTERACTION"
	case CSSMERR_DL_SERVICE_NOT_AVAILABLE:
		return "CSSMERR_DL_SERVICE_NOT_AVAILABLE"
	case CSSMERR_DL_USER_CANCELED:
		return "CSSMERR_DL_USER_CANCELED"
	case CSSMERR_TP_DEVICE_FAILED:
		return "CSSMERR_TP_DEVICE_FAILED"
	case CSSMERR_TP_DEVICE_RESET:
		return "CSSMERR_TP_DEVICE_RESET"
	case CSSMERR_TP_INSUFFICIENT_CLIENT_IDENTIFICATION:
		return "CSSMERR_TP_INSUFFICIENT_CLIENT_IDENTIFICATION"
	case CSSMERR_TP_IN_DARK_WAKE:
		return "CSSMERR_TP_IN_DARK_WAKE"
	case CSSMERR_TP_NO_USER_INTERACTION:
		return "CSSMERR_TP_NO_USER_INTERACTION"
	case CSSMERR_TP_SERVICE_NOT_AVAILABLE:
		return "CSSMERR_TP_SERVICE_NOT_AVAILABLE"
	case CSSMERR_TP_USER_CANCELED:
		return "CSSMERR_TP_USER_CANCELED"
	default:
		return fmt.Sprintf("Cssmerr(%d)", e)
	}
}

type CssmerrAc int32

const (
	CSSMERR_AC_FUNCTION_FAILED          CssmerrAc = -2147405814
	CSSMERR_AC_FUNCTION_NOT_IMPLEMENTED CssmerrAc = -2147405817
	CSSMERR_AC_INTERNAL_ERROR           CssmerrAc = -2147405823
	CSSMERR_AC_INVALID_CL_HANDLE        CssmerrAc = -2147405742
	CSSMERR_AC_INVALID_CONTEXT_HANDLE   CssmerrAc = -2147405760
	CSSMERR_AC_INVALID_DATA             CssmerrAc = -2147405754
	CSSMERR_AC_INVALID_DB_HANDLE        CssmerrAc = -2147405750
	CSSMERR_AC_INVALID_DB_LIST          CssmerrAc = -2147405748
	CSSMERR_AC_INVALID_DB_LIST_POINTER  CssmerrAc = -2147405747
	CSSMERR_AC_INVALID_DL_HANDLE        CssmerrAc = -2147405743
	CSSMERR_AC_INVALID_INPUT_POINTER    CssmerrAc = -2147405819
	CSSMERR_AC_INVALID_OUTPUT_POINTER   CssmerrAc = -2147405818
	CSSMERR_AC_INVALID_PASSTHROUGH_ID   CssmerrAc = -2147405738
	CSSMERR_AC_INVALID_POINTER          CssmerrAc = -2147405820
	CSSMERR_AC_INVALID_TP_HANDLE        CssmerrAc = -2147405741
	CSSMERR_AC_MDS_ERROR                CssmerrAc = -2147405821
	CSSMERR_AC_MEMORY_ERROR             CssmerrAc = -2147405822
	CSSMERR_AC_OS_ACCESS_DENIED         CssmerrAc = -2147405815
	CSSMERR_AC_SELF_CHECK_FAILED        CssmerrAc = -2147405816
)

func (e CssmerrAc) String() string {
	switch e {
	case CSSMERR_AC_FUNCTION_FAILED:
		return "CSSMERR_AC_FUNCTION_FAILED"
	case CSSMERR_AC_FUNCTION_NOT_IMPLEMENTED:
		return "CSSMERR_AC_FUNCTION_NOT_IMPLEMENTED"
	case CSSMERR_AC_INTERNAL_ERROR:
		return "CSSMERR_AC_INTERNAL_ERROR"
	case CSSMERR_AC_INVALID_CL_HANDLE:
		return "CSSMERR_AC_INVALID_CL_HANDLE"
	case CSSMERR_AC_INVALID_CONTEXT_HANDLE:
		return "CSSMERR_AC_INVALID_CONTEXT_HANDLE"
	case CSSMERR_AC_INVALID_DATA:
		return "CSSMERR_AC_INVALID_DATA"
	case CSSMERR_AC_INVALID_DB_HANDLE:
		return "CSSMERR_AC_INVALID_DB_HANDLE"
	case CSSMERR_AC_INVALID_DB_LIST:
		return "CSSMERR_AC_INVALID_DB_LIST"
	case CSSMERR_AC_INVALID_DB_LIST_POINTER:
		return "CSSMERR_AC_INVALID_DB_LIST_POINTER"
	case CSSMERR_AC_INVALID_DL_HANDLE:
		return "CSSMERR_AC_INVALID_DL_HANDLE"
	case CSSMERR_AC_INVALID_INPUT_POINTER:
		return "CSSMERR_AC_INVALID_INPUT_POINTER"
	case CSSMERR_AC_INVALID_OUTPUT_POINTER:
		return "CSSMERR_AC_INVALID_OUTPUT_POINTER"
	case CSSMERR_AC_INVALID_PASSTHROUGH_ID:
		return "CSSMERR_AC_INVALID_PASSTHROUGH_ID"
	case CSSMERR_AC_INVALID_POINTER:
		return "CSSMERR_AC_INVALID_POINTER"
	case CSSMERR_AC_INVALID_TP_HANDLE:
		return "CSSMERR_AC_INVALID_TP_HANDLE"
	case CSSMERR_AC_MDS_ERROR:
		return "CSSMERR_AC_MDS_ERROR"
	case CSSMERR_AC_MEMORY_ERROR:
		return "CSSMERR_AC_MEMORY_ERROR"
	case CSSMERR_AC_OS_ACCESS_DENIED:
		return "CSSMERR_AC_OS_ACCESS_DENIED"
	case CSSMERR_AC_SELF_CHECK_FAILED:
		return "CSSMERR_AC_SELF_CHECK_FAILED"
	default:
		return fmt.Sprintf("CssmerrAc(%d)", e)
	}
}

type CssmerrAppleDotmac int32

const (
	CSSMERR_APPLE_DOTMAC_CSR_VERIFY_FAIL          CssmerrAppleDotmac = -2147408785
	CSSMERR_APPLE_DOTMAC_FAILED_CONSISTENCY_CHECK CssmerrAppleDotmac = -2147408784
	CSSMERR_APPLE_DOTMAC_NO_REQ_PENDING           CssmerrAppleDotmac = -2147408786
	CSSMERR_APPLE_DOTMAC_REQ_IS_PENDING           CssmerrAppleDotmac = -2147408787
	CSSMERR_APPLE_DOTMAC_REQ_QUEUED               CssmerrAppleDotmac = -2147408796
	CSSMERR_APPLE_DOTMAC_REQ_REDIRECT             CssmerrAppleDotmac = -2147408795
	CSSMERR_APPLE_DOTMAC_REQ_SERVER_ALREADY_EXIST CssmerrAppleDotmac = -2147408789
	CSSMERR_APPLE_DOTMAC_REQ_SERVER_AUTH          CssmerrAppleDotmac = -2147408792
	CSSMERR_APPLE_DOTMAC_REQ_SERVER_ERR           CssmerrAppleDotmac = -2147408794
	CSSMERR_APPLE_DOTMAC_REQ_SERVER_NOT_AVAIL     CssmerrAppleDotmac = -2147408790
	CSSMERR_APPLE_DOTMAC_REQ_SERVER_PARAM         CssmerrAppleDotmac = -2147408793
	CSSMERR_APPLE_DOTMAC_REQ_SERVER_SERVICE_ERROR CssmerrAppleDotmac = -2147408788
	CSSMERR_APPLE_DOTMAC_REQ_SERVER_UNIMPL        CssmerrAppleDotmac = -2147408791
)

func (e CssmerrAppleDotmac) String() string {
	switch e {
	case CSSMERR_APPLE_DOTMAC_CSR_VERIFY_FAIL:
		return "CSSMERR_APPLE_DOTMAC_CSR_VERIFY_FAIL"
	case CSSMERR_APPLE_DOTMAC_FAILED_CONSISTENCY_CHECK:
		return "CSSMERR_APPLE_DOTMAC_FAILED_CONSISTENCY_CHECK"
	case CSSMERR_APPLE_DOTMAC_NO_REQ_PENDING:
		return "CSSMERR_APPLE_DOTMAC_NO_REQ_PENDING"
	case CSSMERR_APPLE_DOTMAC_REQ_IS_PENDING:
		return "CSSMERR_APPLE_DOTMAC_REQ_IS_PENDING"
	case CSSMERR_APPLE_DOTMAC_REQ_QUEUED:
		return "CSSMERR_APPLE_DOTMAC_REQ_QUEUED"
	case CSSMERR_APPLE_DOTMAC_REQ_REDIRECT:
		return "CSSMERR_APPLE_DOTMAC_REQ_REDIRECT"
	case CSSMERR_APPLE_DOTMAC_REQ_SERVER_ALREADY_EXIST:
		return "CSSMERR_APPLE_DOTMAC_REQ_SERVER_ALREADY_EXIST"
	case CSSMERR_APPLE_DOTMAC_REQ_SERVER_AUTH:
		return "CSSMERR_APPLE_DOTMAC_REQ_SERVER_AUTH"
	case CSSMERR_APPLE_DOTMAC_REQ_SERVER_ERR:
		return "CSSMERR_APPLE_DOTMAC_REQ_SERVER_ERR"
	case CSSMERR_APPLE_DOTMAC_REQ_SERVER_NOT_AVAIL:
		return "CSSMERR_APPLE_DOTMAC_REQ_SERVER_NOT_AVAIL"
	case CSSMERR_APPLE_DOTMAC_REQ_SERVER_PARAM:
		return "CSSMERR_APPLE_DOTMAC_REQ_SERVER_PARAM"
	case CSSMERR_APPLE_DOTMAC_REQ_SERVER_SERVICE_ERROR:
		return "CSSMERR_APPLE_DOTMAC_REQ_SERVER_SERVICE_ERROR"
	case CSSMERR_APPLE_DOTMAC_REQ_SERVER_UNIMPL:
		return "CSSMERR_APPLE_DOTMAC_REQ_SERVER_UNIMPL"
	default:
		return fmt.Sprintf("CssmerrAppleDotmac(%d)", e)
	}
}

type CssmerrAppledl int32

const (
	CSSMERR_APPLEDL_DISK_FULL                  CssmerrAppledl = -2147412991
	CSSMERR_APPLEDL_FILE_TOO_BIG               CssmerrAppledl = -2147412989
	CSSMERR_APPLEDL_INCOMPATIBLE_DATABASE_BLOB CssmerrAppledl = -2147412986
	CSSMERR_APPLEDL_INCOMPATIBLE_KEY_BLOB      CssmerrAppledl = -2147412985
	CSSMERR_APPLEDL_INVALID_DATABASE_BLOB      CssmerrAppledl = -2147412988
	CSSMERR_APPLEDL_INVALID_KEY_BLOB           CssmerrAppledl = -2147412987
	CSSMERR_APPLEDL_INVALID_OPEN_PARAMETERS    CssmerrAppledl = -2147412992
	CSSMERR_APPLEDL_QUOTA_EXCEEDED             CssmerrAppledl = -2147412990
)

func (e CssmerrAppledl) String() string {
	switch e {
	case CSSMERR_APPLEDL_DISK_FULL:
		return "CSSMERR_APPLEDL_DISK_FULL"
	case CSSMERR_APPLEDL_FILE_TOO_BIG:
		return "CSSMERR_APPLEDL_FILE_TOO_BIG"
	case CSSMERR_APPLEDL_INCOMPATIBLE_DATABASE_BLOB:
		return "CSSMERR_APPLEDL_INCOMPATIBLE_DATABASE_BLOB"
	case CSSMERR_APPLEDL_INCOMPATIBLE_KEY_BLOB:
		return "CSSMERR_APPLEDL_INCOMPATIBLE_KEY_BLOB"
	case CSSMERR_APPLEDL_INVALID_DATABASE_BLOB:
		return "CSSMERR_APPLEDL_INVALID_DATABASE_BLOB"
	case CSSMERR_APPLEDL_INVALID_KEY_BLOB:
		return "CSSMERR_APPLEDL_INVALID_KEY_BLOB"
	case CSSMERR_APPLEDL_INVALID_OPEN_PARAMETERS:
		return "CSSMERR_APPLEDL_INVALID_OPEN_PARAMETERS"
	case CSSMERR_APPLEDL_QUOTA_EXCEEDED:
		return "CSSMERR_APPLEDL_QUOTA_EXCEEDED"
	default:
		return fmt.Sprintf("CssmerrAppledl(%d)", e)
	}
}

type CssmerrAppletp int32

const (
	CSSMERR_APPLETP_BAD_CERT_FROM_ISSUER         CssmerrAppletp = -2147408873
	CSSMERR_APPLETP_CA_PIN_MISMATCH              CssmerrAppletp = -2147408836
	CSSMERR_APPLETP_CERT_NOT_FOUND_FROM_ISSUER   CssmerrAppletp = -2147408874
	CSSMERR_APPLETP_CODE_SIGN_DEVELOPMENT        CssmerrAppletp = -2147408845
	CSSMERR_APPLETP_CRL_BAD_URI                  CssmerrAppletp = -2147408881
	CSSMERR_APPLETP_CRL_EXPIRED                  CssmerrAppletp = -2147408885
	CSSMERR_APPLETP_CRL_INVALID_ANCHOR_CERT      CssmerrAppletp = -2147408877
	CSSMERR_APPLETP_CRL_NOT_FOUND                CssmerrAppletp = -2147408883
	CSSMERR_APPLETP_CRL_NOT_TRUSTED              CssmerrAppletp = -2147408878
	CSSMERR_APPLETP_CRL_NOT_VALID_YET            CssmerrAppletp = -2147408884
	CSSMERR_APPLETP_CRL_POLICY_FAIL              CssmerrAppletp = -2147408876
	CSSMERR_APPLETP_CRL_SERVER_DOWN              CssmerrAppletp = -2147408882
	CSSMERR_APPLETP_CS_BAD_CERT_CHAIN_LENGTH     CssmerrAppletp = -2147408849
	CSSMERR_APPLETP_CS_BAD_PATH_LENGTH           CssmerrAppletp = -2147408847
	CSSMERR_APPLETP_CS_NO_BASIC_CONSTRAINTS      CssmerrAppletp = -2147408848
	CSSMERR_APPLETP_CS_NO_EXTENDED_KEY_USAGE     CssmerrAppletp = -2147408846
	CSSMERR_APPLETP_EXT_KEYUSAGE_NOT_CRITICAL    CssmerrAppletp = -2147408838
	CSSMERR_APPLETP_HOSTNAME_MISMATCH            CssmerrAppletp = -2147408896
	CSSMERR_APPLETP_IDENTIFIER_MISSING           CssmerrAppletp = -2147408837
	CSSMERR_APPLETP_IDP_FAIL                     CssmerrAppletp = -2147408875
	CSSMERR_APPLETP_INCOMPLETE_REVOCATION_CHECK  CssmerrAppletp = -2147408861
	CSSMERR_APPLETP_INVALID_AUTHORITY_ID         CssmerrAppletp = -2147408892
	CSSMERR_APPLETP_INVALID_CA                   CssmerrAppletp = -2147408893
	CSSMERR_APPLETP_INVALID_EMPTY_SUBJECT        CssmerrAppletp = -2147408841
	CSSMERR_APPLETP_INVALID_EXTENDED_KEY_USAGE   CssmerrAppletp = -2147408889
	CSSMERR_APPLETP_INVALID_ID_LINKAGE           CssmerrAppletp = -2147408888
	CSSMERR_APPLETP_INVALID_KEY_USAGE            CssmerrAppletp = -2147408890
	CSSMERR_APPLETP_INVALID_ROOT                 CssmerrAppletp = -2147408886
	CSSMERR_APPLETP_INVALID_SUBJECT_ID           CssmerrAppletp = -2147408891
	CSSMERR_APPLETP_LEAF_PIN_MISMATCH            CssmerrAppletp = -2147408835
	CSSMERR_APPLETP_MISSING_REQUIRED_EXTENSION   CssmerrAppletp = -2147408839
	CSSMERR_APPLETP_NETWORK_FAILURE              CssmerrAppletp = -2147408860
	CSSMERR_APPLETP_NO_BASIC_CONSTRAINTS         CssmerrAppletp = -2147408894
	CSSMERR_APPLETP_OCSP_BAD_REQUEST             CssmerrAppletp = -2147408864
	CSSMERR_APPLETP_OCSP_BAD_RESPONSE            CssmerrAppletp = -2147408865
	CSSMERR_APPLETP_OCSP_INVALID_ANCHOR_CERT     CssmerrAppletp = -2147408858
	CSSMERR_APPLETP_OCSP_NONCE_MISMATCH          CssmerrAppletp = -2147408850
	CSSMERR_APPLETP_OCSP_NOT_TRUSTED             CssmerrAppletp = -2147408859
	CSSMERR_APPLETP_OCSP_NO_SIGNER               CssmerrAppletp = -2147408856
	CSSMERR_APPLETP_OCSP_RESP_INTERNAL_ERR       CssmerrAppletp = -2147408854
	CSSMERR_APPLETP_OCSP_RESP_MALFORMED_REQ      CssmerrAppletp = -2147408855
	CSSMERR_APPLETP_OCSP_RESP_SIG_REQUIRED       CssmerrAppletp = -2147408852
	CSSMERR_APPLETP_OCSP_RESP_TRY_LATER          CssmerrAppletp = -2147408853
	CSSMERR_APPLETP_OCSP_RESP_UNAUTHORIZED       CssmerrAppletp = -2147408851
	CSSMERR_APPLETP_OCSP_SIG_ERROR               CssmerrAppletp = -2147408857
	CSSMERR_APPLETP_OCSP_STATUS_UNRECOGNIZED     CssmerrAppletp = -2147408862
	CSSMERR_APPLETP_OCSP_UNAVAILABLE             CssmerrAppletp = -2147408863
	CSSMERR_APPLETP_PATH_LEN_CONSTRAINT          CssmerrAppletp = -2147408887
	CSSMERR_APPLETP_RS_BAD_CERT_CHAIN_LENGTH     CssmerrAppletp = -2147408844
	CSSMERR_APPLETP_RS_BAD_EXTENDED_KEY_USAGE    CssmerrAppletp = -2147408843
	CSSMERR_APPLETP_SMIME_BAD_EXT_KEY_USE        CssmerrAppletp = -2147408871
	CSSMERR_APPLETP_SMIME_BAD_KEY_USE            CssmerrAppletp = -2147408870
	CSSMERR_APPLETP_SMIME_EMAIL_ADDRS_NOT_FOUND  CssmerrAppletp = -2147408872
	CSSMERR_APPLETP_SMIME_KEYUSAGE_NOT_CRITICAL  CssmerrAppletp = -2147408869
	CSSMERR_APPLETP_SMIME_NO_EMAIL_ADDRS         CssmerrAppletp = -2147408868
	CSSMERR_APPLETP_SMIME_SUBJ_ALT_NAME_NOT_CRIT CssmerrAppletp = -2147408867
	CSSMERR_APPLETP_SSL_BAD_EXT_KEY_USE          CssmerrAppletp = -2147408866
	CSSMERR_APPLETP_TRUST_SETTING_DENY           CssmerrAppletp = -2147408842
	CSSMERR_APPLETP_UNKNOWN_CERT_EXTEN           CssmerrAppletp = -2147408880
	CSSMERR_APPLETP_UNKNOWN_CRITICAL_EXTEN       CssmerrAppletp = -2147408895
	CSSMERR_APPLETP_UNKNOWN_CRL_EXTEN            CssmerrAppletp = -2147408879
	CSSMERR_APPLETP_UNKNOWN_QUAL_CERT_STATEMENT  CssmerrAppletp = -2147408840
)

func (e CssmerrAppletp) String() string {
	switch e {
	case CSSMERR_APPLETP_BAD_CERT_FROM_ISSUER:
		return "CSSMERR_APPLETP_BAD_CERT_FROM_ISSUER"
	case CSSMERR_APPLETP_CA_PIN_MISMATCH:
		return "CSSMERR_APPLETP_CA_PIN_MISMATCH"
	case CSSMERR_APPLETP_CERT_NOT_FOUND_FROM_ISSUER:
		return "CSSMERR_APPLETP_CERT_NOT_FOUND_FROM_ISSUER"
	case CSSMERR_APPLETP_CODE_SIGN_DEVELOPMENT:
		return "CSSMERR_APPLETP_CODE_SIGN_DEVELOPMENT"
	case CSSMERR_APPLETP_CRL_BAD_URI:
		return "CSSMERR_APPLETP_CRL_BAD_URI"
	case CSSMERR_APPLETP_CRL_EXPIRED:
		return "CSSMERR_APPLETP_CRL_EXPIRED"
	case CSSMERR_APPLETP_CRL_INVALID_ANCHOR_CERT:
		return "CSSMERR_APPLETP_CRL_INVALID_ANCHOR_CERT"
	case CSSMERR_APPLETP_CRL_NOT_FOUND:
		return "CSSMERR_APPLETP_CRL_NOT_FOUND"
	case CSSMERR_APPLETP_CRL_NOT_TRUSTED:
		return "CSSMERR_APPLETP_CRL_NOT_TRUSTED"
	case CSSMERR_APPLETP_CRL_NOT_VALID_YET:
		return "CSSMERR_APPLETP_CRL_NOT_VALID_YET"
	case CSSMERR_APPLETP_CRL_POLICY_FAIL:
		return "CSSMERR_APPLETP_CRL_POLICY_FAIL"
	case CSSMERR_APPLETP_CRL_SERVER_DOWN:
		return "CSSMERR_APPLETP_CRL_SERVER_DOWN"
	case CSSMERR_APPLETP_CS_BAD_CERT_CHAIN_LENGTH:
		return "CSSMERR_APPLETP_CS_BAD_CERT_CHAIN_LENGTH"
	case CSSMERR_APPLETP_CS_BAD_PATH_LENGTH:
		return "CSSMERR_APPLETP_CS_BAD_PATH_LENGTH"
	case CSSMERR_APPLETP_CS_NO_BASIC_CONSTRAINTS:
		return "CSSMERR_APPLETP_CS_NO_BASIC_CONSTRAINTS"
	case CSSMERR_APPLETP_CS_NO_EXTENDED_KEY_USAGE:
		return "CSSMERR_APPLETP_CS_NO_EXTENDED_KEY_USAGE"
	case CSSMERR_APPLETP_EXT_KEYUSAGE_NOT_CRITICAL:
		return "CSSMERR_APPLETP_EXT_KEYUSAGE_NOT_CRITICAL"
	case CSSMERR_APPLETP_HOSTNAME_MISMATCH:
		return "CSSMERR_APPLETP_HOSTNAME_MISMATCH"
	case CSSMERR_APPLETP_IDENTIFIER_MISSING:
		return "CSSMERR_APPLETP_IDENTIFIER_MISSING"
	case CSSMERR_APPLETP_IDP_FAIL:
		return "CSSMERR_APPLETP_IDP_FAIL"
	case CSSMERR_APPLETP_INCOMPLETE_REVOCATION_CHECK:
		return "CSSMERR_APPLETP_INCOMPLETE_REVOCATION_CHECK"
	case CSSMERR_APPLETP_INVALID_AUTHORITY_ID:
		return "CSSMERR_APPLETP_INVALID_AUTHORITY_ID"
	case CSSMERR_APPLETP_INVALID_CA:
		return "CSSMERR_APPLETP_INVALID_CA"
	case CSSMERR_APPLETP_INVALID_EMPTY_SUBJECT:
		return "CSSMERR_APPLETP_INVALID_EMPTY_SUBJECT"
	case CSSMERR_APPLETP_INVALID_EXTENDED_KEY_USAGE:
		return "CSSMERR_APPLETP_INVALID_EXTENDED_KEY_USAGE"
	case CSSMERR_APPLETP_INVALID_ID_LINKAGE:
		return "CSSMERR_APPLETP_INVALID_ID_LINKAGE"
	case CSSMERR_APPLETP_INVALID_KEY_USAGE:
		return "CSSMERR_APPLETP_INVALID_KEY_USAGE"
	case CSSMERR_APPLETP_INVALID_ROOT:
		return "CSSMERR_APPLETP_INVALID_ROOT"
	case CSSMERR_APPLETP_INVALID_SUBJECT_ID:
		return "CSSMERR_APPLETP_INVALID_SUBJECT_ID"
	case CSSMERR_APPLETP_LEAF_PIN_MISMATCH:
		return "CSSMERR_APPLETP_LEAF_PIN_MISMATCH"
	case CSSMERR_APPLETP_MISSING_REQUIRED_EXTENSION:
		return "CSSMERR_APPLETP_MISSING_REQUIRED_EXTENSION"
	case CSSMERR_APPLETP_NETWORK_FAILURE:
		return "CSSMERR_APPLETP_NETWORK_FAILURE"
	case CSSMERR_APPLETP_NO_BASIC_CONSTRAINTS:
		return "CSSMERR_APPLETP_NO_BASIC_CONSTRAINTS"
	case CSSMERR_APPLETP_OCSP_BAD_REQUEST:
		return "CSSMERR_APPLETP_OCSP_BAD_REQUEST"
	case CSSMERR_APPLETP_OCSP_BAD_RESPONSE:
		return "CSSMERR_APPLETP_OCSP_BAD_RESPONSE"
	case CSSMERR_APPLETP_OCSP_INVALID_ANCHOR_CERT:
		return "CSSMERR_APPLETP_OCSP_INVALID_ANCHOR_CERT"
	case CSSMERR_APPLETP_OCSP_NONCE_MISMATCH:
		return "CSSMERR_APPLETP_OCSP_NONCE_MISMATCH"
	case CSSMERR_APPLETP_OCSP_NOT_TRUSTED:
		return "CSSMERR_APPLETP_OCSP_NOT_TRUSTED"
	case CSSMERR_APPLETP_OCSP_NO_SIGNER:
		return "CSSMERR_APPLETP_OCSP_NO_SIGNER"
	case CSSMERR_APPLETP_OCSP_RESP_INTERNAL_ERR:
		return "CSSMERR_APPLETP_OCSP_RESP_INTERNAL_ERR"
	case CSSMERR_APPLETP_OCSP_RESP_MALFORMED_REQ:
		return "CSSMERR_APPLETP_OCSP_RESP_MALFORMED_REQ"
	case CSSMERR_APPLETP_OCSP_RESP_SIG_REQUIRED:
		return "CSSMERR_APPLETP_OCSP_RESP_SIG_REQUIRED"
	case CSSMERR_APPLETP_OCSP_RESP_TRY_LATER:
		return "CSSMERR_APPLETP_OCSP_RESP_TRY_LATER"
	case CSSMERR_APPLETP_OCSP_RESP_UNAUTHORIZED:
		return "CSSMERR_APPLETP_OCSP_RESP_UNAUTHORIZED"
	case CSSMERR_APPLETP_OCSP_SIG_ERROR:
		return "CSSMERR_APPLETP_OCSP_SIG_ERROR"
	case CSSMERR_APPLETP_OCSP_STATUS_UNRECOGNIZED:
		return "CSSMERR_APPLETP_OCSP_STATUS_UNRECOGNIZED"
	case CSSMERR_APPLETP_OCSP_UNAVAILABLE:
		return "CSSMERR_APPLETP_OCSP_UNAVAILABLE"
	case CSSMERR_APPLETP_PATH_LEN_CONSTRAINT:
		return "CSSMERR_APPLETP_PATH_LEN_CONSTRAINT"
	case CSSMERR_APPLETP_RS_BAD_CERT_CHAIN_LENGTH:
		return "CSSMERR_APPLETP_RS_BAD_CERT_CHAIN_LENGTH"
	case CSSMERR_APPLETP_RS_BAD_EXTENDED_KEY_USAGE:
		return "CSSMERR_APPLETP_RS_BAD_EXTENDED_KEY_USAGE"
	case CSSMERR_APPLETP_SMIME_BAD_EXT_KEY_USE:
		return "CSSMERR_APPLETP_SMIME_BAD_EXT_KEY_USE"
	case CSSMERR_APPLETP_SMIME_BAD_KEY_USE:
		return "CSSMERR_APPLETP_SMIME_BAD_KEY_USE"
	case CSSMERR_APPLETP_SMIME_EMAIL_ADDRS_NOT_FOUND:
		return "CSSMERR_APPLETP_SMIME_EMAIL_ADDRS_NOT_FOUND"
	case CSSMERR_APPLETP_SMIME_KEYUSAGE_NOT_CRITICAL:
		return "CSSMERR_APPLETP_SMIME_KEYUSAGE_NOT_CRITICAL"
	case CSSMERR_APPLETP_SMIME_NO_EMAIL_ADDRS:
		return "CSSMERR_APPLETP_SMIME_NO_EMAIL_ADDRS"
	case CSSMERR_APPLETP_SMIME_SUBJ_ALT_NAME_NOT_CRIT:
		return "CSSMERR_APPLETP_SMIME_SUBJ_ALT_NAME_NOT_CRIT"
	case CSSMERR_APPLETP_SSL_BAD_EXT_KEY_USE:
		return "CSSMERR_APPLETP_SSL_BAD_EXT_KEY_USE"
	case CSSMERR_APPLETP_TRUST_SETTING_DENY:
		return "CSSMERR_APPLETP_TRUST_SETTING_DENY"
	case CSSMERR_APPLETP_UNKNOWN_CERT_EXTEN:
		return "CSSMERR_APPLETP_UNKNOWN_CERT_EXTEN"
	case CSSMERR_APPLETP_UNKNOWN_CRITICAL_EXTEN:
		return "CSSMERR_APPLETP_UNKNOWN_CRITICAL_EXTEN"
	case CSSMERR_APPLETP_UNKNOWN_CRL_EXTEN:
		return "CSSMERR_APPLETP_UNKNOWN_CRL_EXTEN"
	case CSSMERR_APPLETP_UNKNOWN_QUAL_CERT_STATEMENT:
		return "CSSMERR_APPLETP_UNKNOWN_QUAL_CERT_STATEMENT"
	default:
		return fmt.Sprintf("CssmerrAppletp(%d)", e)
	}
}

type CssmerrCL int32

const (
	CSSMERR_CL_CRL_ALREADY_SIGNED        CssmerrCL = -2147411897
	CSSMERR_CL_FUNCTION_FAILED           CssmerrCL = -2147411958
	CSSMERR_CL_FUNCTION_NOT_IMPLEMENTED  CssmerrCL = -2147411961
	CSSMERR_CL_INTERNAL_ERROR            CssmerrCL = -2147411967
	CSSMERR_CL_INVALID_CERTGROUP_POINTER CssmerrCL = -2147411902
	CSSMERR_CL_INVALID_CERT_POINTER      CssmerrCL = -2147411901
	CSSMERR_CL_INVALID_CONTEXT_HANDLE    CssmerrCL = -2147411904
	CSSMERR_CL_INVALID_CRL_POINTER       CssmerrCL = -2147411900
	CSSMERR_CL_INVALID_DATA              CssmerrCL = -2147411898
	CSSMERR_CL_INVALID_FIELD_POINTER     CssmerrCL = -2147411899
	CSSMERR_CL_INVALID_INPUT_POINTER     CssmerrCL = -2147411963
	CSSMERR_CL_INVALID_NUMBER_OF_FIELDS  CssmerrCL = -2147411896
	CSSMERR_CL_INVALID_OUTPUT_POINTER    CssmerrCL = -2147411962
	CSSMERR_CL_INVALID_PASSTHROUGH_ID    CssmerrCL = -2147411882
	CSSMERR_CL_INVALID_POINTER           CssmerrCL = -2147411964
	CSSMERR_CL_MDS_ERROR                 CssmerrCL = -2147411965
	CSSMERR_CL_MEMORY_ERROR              CssmerrCL = -2147411966
	CSSMERR_CL_OS_ACCESS_DENIED          CssmerrCL = -2147411959
	CSSMERR_CL_SELF_CHECK_FAILED         CssmerrCL = -2147411960
	CSSMERR_CL_UNKNOWN_FORMAT            CssmerrCL = -2147411890
	CSSMERR_CL_UNKNOWN_TAG               CssmerrCL = -2147411889
	CSSMERR_CL_VERIFICATION_FAILURE      CssmerrCL = -2147411895
)

func (e CssmerrCL) String() string {
	switch e {
	case CSSMERR_CL_CRL_ALREADY_SIGNED:
		return "CSSMERR_CL_CRL_ALREADY_SIGNED"
	case CSSMERR_CL_FUNCTION_FAILED:
		return "CSSMERR_CL_FUNCTION_FAILED"
	case CSSMERR_CL_FUNCTION_NOT_IMPLEMENTED:
		return "CSSMERR_CL_FUNCTION_NOT_IMPLEMENTED"
	case CSSMERR_CL_INTERNAL_ERROR:
		return "CSSMERR_CL_INTERNAL_ERROR"
	case CSSMERR_CL_INVALID_CERTGROUP_POINTER:
		return "CSSMERR_CL_INVALID_CERTGROUP_POINTER"
	case CSSMERR_CL_INVALID_CERT_POINTER:
		return "CSSMERR_CL_INVALID_CERT_POINTER"
	case CSSMERR_CL_INVALID_CONTEXT_HANDLE:
		return "CSSMERR_CL_INVALID_CONTEXT_HANDLE"
	case CSSMERR_CL_INVALID_CRL_POINTER:
		return "CSSMERR_CL_INVALID_CRL_POINTER"
	case CSSMERR_CL_INVALID_DATA:
		return "CSSMERR_CL_INVALID_DATA"
	case CSSMERR_CL_INVALID_FIELD_POINTER:
		return "CSSMERR_CL_INVALID_FIELD_POINTER"
	case CSSMERR_CL_INVALID_INPUT_POINTER:
		return "CSSMERR_CL_INVALID_INPUT_POINTER"
	case CSSMERR_CL_INVALID_NUMBER_OF_FIELDS:
		return "CSSMERR_CL_INVALID_NUMBER_OF_FIELDS"
	case CSSMERR_CL_INVALID_OUTPUT_POINTER:
		return "CSSMERR_CL_INVALID_OUTPUT_POINTER"
	case CSSMERR_CL_INVALID_PASSTHROUGH_ID:
		return "CSSMERR_CL_INVALID_PASSTHROUGH_ID"
	case CSSMERR_CL_INVALID_POINTER:
		return "CSSMERR_CL_INVALID_POINTER"
	case CSSMERR_CL_MDS_ERROR:
		return "CSSMERR_CL_MDS_ERROR"
	case CSSMERR_CL_MEMORY_ERROR:
		return "CSSMERR_CL_MEMORY_ERROR"
	case CSSMERR_CL_OS_ACCESS_DENIED:
		return "CSSMERR_CL_OS_ACCESS_DENIED"
	case CSSMERR_CL_SELF_CHECK_FAILED:
		return "CSSMERR_CL_SELF_CHECK_FAILED"
	case CSSMERR_CL_UNKNOWN_FORMAT:
		return "CSSMERR_CL_UNKNOWN_FORMAT"
	case CSSMERR_CL_UNKNOWN_TAG:
		return "CSSMERR_CL_UNKNOWN_TAG"
	case CSSMERR_CL_VERIFICATION_FAILURE:
		return "CSSMERR_CL_VERIFICATION_FAILURE"
	default:
		return fmt.Sprintf("CssmerrCL(%d)", e)
	}
}

type CssmerrCspAppleAddApplicationAclSubject int32

const (
	CSSMERR_CSPDL_APPLE_DL_CONVERSION_ERROR       CssmerrCspAppleAddApplicationAclSubject = -2147415035
	CSSMERR_CSP_APPLE_ADD_APPLICATION_ACL_SUBJECT CssmerrCspAppleAddApplicationAclSubject = -2147415040
	CSSMERR_CSP_APPLE_INVALID_KEY_END_DATE        CssmerrCspAppleAddApplicationAclSubject = -2147415036
	CSSMERR_CSP_APPLE_INVALID_KEY_START_DATE      CssmerrCspAppleAddApplicationAclSubject = -2147415037
	CSSMERR_CSP_APPLE_PUBLIC_KEY_INCOMPLETE       CssmerrCspAppleAddApplicationAclSubject = -2147415039
	CSSMERR_CSP_APPLE_SIGNATURE_MISMATCH          CssmerrCspAppleAddApplicationAclSubject = -2147415038
	CSSMERR_CSP_APPLE_SSLv2_ROLLBACK              CssmerrCspAppleAddApplicationAclSubject = -2147415034
)

func (e CssmerrCspAppleAddApplicationAclSubject) String() string {
	switch e {
	case CSSMERR_CSPDL_APPLE_DL_CONVERSION_ERROR:
		return "CSSMERR_CSPDL_APPLE_DL_CONVERSION_ERROR"
	case CSSMERR_CSP_APPLE_ADD_APPLICATION_ACL_SUBJECT:
		return "CSSMERR_CSP_APPLE_ADD_APPLICATION_ACL_SUBJECT"
	case CSSMERR_CSP_APPLE_INVALID_KEY_END_DATE:
		return "CSSMERR_CSP_APPLE_INVALID_KEY_END_DATE"
	case CSSMERR_CSP_APPLE_INVALID_KEY_START_DATE:
		return "CSSMERR_CSP_APPLE_INVALID_KEY_START_DATE"
	case CSSMERR_CSP_APPLE_PUBLIC_KEY_INCOMPLETE:
		return "CSSMERR_CSP_APPLE_PUBLIC_KEY_INCOMPLETE"
	case CSSMERR_CSP_APPLE_SIGNATURE_MISMATCH:
		return "CSSMERR_CSP_APPLE_SIGNATURE_MISMATCH"
	case CSSMERR_CSP_APPLE_SSLv2_ROLLBACK:
		return "CSSMERR_CSP_APPLE_SSLv2_ROLLBACK"
	default:
		return fmt.Sprintf("CssmerrCspAppleAddApplicationAclSubject(%d)", e)
	}
}

type CssmerrCspInternalError int32

const (
	CSSMERR_CSP_FUNCTION_FAILED          CssmerrCspInternalError = -2147416054
	CSSMERR_CSP_FUNCTION_NOT_IMPLEMENTED CssmerrCspInternalError = -2147416057
	CSSMERR_CSP_INTERNAL_ERROR           CssmerrCspInternalError = -2147416063
	CSSMERR_CSP_INVALID_INPUT_POINTER    CssmerrCspInternalError = -2147416059
	CSSMERR_CSP_INVALID_OUTPUT_POINTER   CssmerrCspInternalError = -2147416058
	CSSMERR_CSP_INVALID_POINTER          CssmerrCspInternalError = -2147416060
	CSSMERR_CSP_MDS_ERROR                CssmerrCspInternalError = -2147416061
	CSSMERR_CSP_MEMORY_ERROR             CssmerrCspInternalError = -2147416062
	CSSMERR_CSP_OS_ACCESS_DENIED         CssmerrCspInternalError = -2147416055
	CSSMERR_CSP_SELF_CHECK_FAILED        CssmerrCspInternalError = -2147416056
)

func (e CssmerrCspInternalError) String() string {
	switch e {
	case CSSMERR_CSP_FUNCTION_FAILED:
		return "CSSMERR_CSP_FUNCTION_FAILED"
	case CSSMERR_CSP_FUNCTION_NOT_IMPLEMENTED:
		return "CSSMERR_CSP_FUNCTION_NOT_IMPLEMENTED"
	case CSSMERR_CSP_INTERNAL_ERROR:
		return "CSSMERR_CSP_INTERNAL_ERROR"
	case CSSMERR_CSP_INVALID_INPUT_POINTER:
		return "CSSMERR_CSP_INVALID_INPUT_POINTER"
	case CSSMERR_CSP_INVALID_OUTPUT_POINTER:
		return "CSSMERR_CSP_INVALID_OUTPUT_POINTER"
	case CSSMERR_CSP_INVALID_POINTER:
		return "CSSMERR_CSP_INVALID_POINTER"
	case CSSMERR_CSP_MDS_ERROR:
		return "CSSMERR_CSP_MDS_ERROR"
	case CSSMERR_CSP_MEMORY_ERROR:
		return "CSSMERR_CSP_MEMORY_ERROR"
	case CSSMERR_CSP_OS_ACCESS_DENIED:
		return "CSSMERR_CSP_OS_ACCESS_DENIED"
	case CSSMERR_CSP_SELF_CHECK_FAILED:
		return "CSSMERR_CSP_SELF_CHECK_FAILED"
	default:
		return fmt.Sprintf("CssmerrCspInternalError(%d)", e)
	}
}

type CssmerrCspInvalidContextHandle int32

const (
	CSSMERR_CSP_INVALID_CONTEXT_HANDLE CssmerrCspInvalidContextHandle = -2147416000
	CSSMERR_CSP_INVALID_CRYPTO_DATA    CssmerrCspInvalidContextHandle = -2147415976
	CSSMERR_CSP_INVALID_DATA           CssmerrCspInvalidContextHandle = -2147415994
	CSSMERR_CSP_INVALID_PASSTHROUGH_ID CssmerrCspInvalidContextHandle = -2147415978
	CSSMERR_CSP_PRIVILEGE_NOT_GRANTED  CssmerrCspInvalidContextHandle = -2147415989
)

func (e CssmerrCspInvalidContextHandle) String() string {
	switch e {
	case CSSMERR_CSP_INVALID_CONTEXT_HANDLE:
		return "CSSMERR_CSP_INVALID_CONTEXT_HANDLE"
	case CSSMERR_CSP_INVALID_CRYPTO_DATA:
		return "CSSMERR_CSP_INVALID_CRYPTO_DATA"
	case CSSMERR_CSP_INVALID_DATA:
		return "CSSMERR_CSP_INVALID_DATA"
	case CSSMERR_CSP_INVALID_PASSTHROUGH_ID:
		return "CSSMERR_CSP_INVALID_PASSTHROUGH_ID"
	case CSSMERR_CSP_PRIVILEGE_NOT_GRANTED:
		return "CSSMERR_CSP_PRIVILEGE_NOT_GRANTED"
	default:
		return fmt.Sprintf("CssmerrCspInvalidContextHandle(%d)", e)
	}
}

type CssmerrCspOperationAuthDenied int32

const (
	CSSMERR_CSP_ACL_ADD_FAILED                 CssmerrCspOperationAuthDenied = -2147416010
	CSSMERR_CSP_ACL_BASE_CERTS_NOT_SUPPORTED   CssmerrCspOperationAuthDenied = -2147416025
	CSSMERR_CSP_ACL_CHALLENGE_CALLBACK_FAILED  CssmerrCspOperationAuthDenied = -2147416019
	CSSMERR_CSP_ACL_CHANGE_FAILED              CssmerrCspOperationAuthDenied = -2147416015
	CSSMERR_CSP_ACL_DELETE_FAILED              CssmerrCspOperationAuthDenied = -2147416012
	CSSMERR_CSP_ACL_ENTRY_TAG_NOT_FOUND        CssmerrCspOperationAuthDenied = -2147416017
	CSSMERR_CSP_ACL_REPLACE_FAILED             CssmerrCspOperationAuthDenied = -2147416011
	CSSMERR_CSP_ACL_SUBJECT_TYPE_NOT_SUPPORTED CssmerrCspOperationAuthDenied = -2147416021
	CSSMERR_CSP_INVALID_ACCESS_CREDENTIALS     CssmerrCspOperationAuthDenied = -2147416027
	CSSMERR_CSP_INVALID_ACL_BASE_CERTS         CssmerrCspOperationAuthDenied = -2147416026
	CSSMERR_CSP_INVALID_ACL_CHALLENGE_CALLBACK CssmerrCspOperationAuthDenied = -2147416020
	CSSMERR_CSP_INVALID_ACL_EDIT_MODE          CssmerrCspOperationAuthDenied = -2147416016
	CSSMERR_CSP_INVALID_ACL_ENTRY_TAG          CssmerrCspOperationAuthDenied = -2147416018
	CSSMERR_CSP_INVALID_ACL_SUBJECT_VALUE      CssmerrCspOperationAuthDenied = -2147416022
	CSSMERR_CSP_INVALID_NEW_ACL_ENTRY          CssmerrCspOperationAuthDenied = -2147416014
	CSSMERR_CSP_INVALID_NEW_ACL_OWNER          CssmerrCspOperationAuthDenied = -2147416013
	CSSMERR_CSP_INVALID_SAMPLE_VALUE           CssmerrCspOperationAuthDenied = -2147416024
	CSSMERR_CSP_OBJECT_ACL_NOT_SUPPORTED       CssmerrCspOperationAuthDenied = -2147416029
	CSSMERR_CSP_OBJECT_ACL_REQUIRED            CssmerrCspOperationAuthDenied = -2147416028
	CSSMERR_CSP_OBJECT_MANIP_AUTH_DENIED       CssmerrCspOperationAuthDenied = -2147416030
	CSSMERR_CSP_OBJECT_USE_AUTH_DENIED         CssmerrCspOperationAuthDenied = -2147416031
	CSSMERR_CSP_OPERATION_AUTH_DENIED          CssmerrCspOperationAuthDenied = -2147416032
	CSSMERR_CSP_SAMPLE_VALUE_NOT_SUPPORTED     CssmerrCspOperationAuthDenied = -2147416023
)

func (e CssmerrCspOperationAuthDenied) String() string {
	switch e {
	case CSSMERR_CSP_ACL_ADD_FAILED:
		return "CSSMERR_CSP_ACL_ADD_FAILED"
	case CSSMERR_CSP_ACL_BASE_CERTS_NOT_SUPPORTED:
		return "CSSMERR_CSP_ACL_BASE_CERTS_NOT_SUPPORTED"
	case CSSMERR_CSP_ACL_CHALLENGE_CALLBACK_FAILED:
		return "CSSMERR_CSP_ACL_CHALLENGE_CALLBACK_FAILED"
	case CSSMERR_CSP_ACL_CHANGE_FAILED:
		return "CSSMERR_CSP_ACL_CHANGE_FAILED"
	case CSSMERR_CSP_ACL_DELETE_FAILED:
		return "CSSMERR_CSP_ACL_DELETE_FAILED"
	case CSSMERR_CSP_ACL_ENTRY_TAG_NOT_FOUND:
		return "CSSMERR_CSP_ACL_ENTRY_TAG_NOT_FOUND"
	case CSSMERR_CSP_ACL_REPLACE_FAILED:
		return "CSSMERR_CSP_ACL_REPLACE_FAILED"
	case CSSMERR_CSP_ACL_SUBJECT_TYPE_NOT_SUPPORTED:
		return "CSSMERR_CSP_ACL_SUBJECT_TYPE_NOT_SUPPORTED"
	case CSSMERR_CSP_INVALID_ACCESS_CREDENTIALS:
		return "CSSMERR_CSP_INVALID_ACCESS_CREDENTIALS"
	case CSSMERR_CSP_INVALID_ACL_BASE_CERTS:
		return "CSSMERR_CSP_INVALID_ACL_BASE_CERTS"
	case CSSMERR_CSP_INVALID_ACL_CHALLENGE_CALLBACK:
		return "CSSMERR_CSP_INVALID_ACL_CHALLENGE_CALLBACK"
	case CSSMERR_CSP_INVALID_ACL_EDIT_MODE:
		return "CSSMERR_CSP_INVALID_ACL_EDIT_MODE"
	case CSSMERR_CSP_INVALID_ACL_ENTRY_TAG:
		return "CSSMERR_CSP_INVALID_ACL_ENTRY_TAG"
	case CSSMERR_CSP_INVALID_ACL_SUBJECT_VALUE:
		return "CSSMERR_CSP_INVALID_ACL_SUBJECT_VALUE"
	case CSSMERR_CSP_INVALID_NEW_ACL_ENTRY:
		return "CSSMERR_CSP_INVALID_NEW_ACL_ENTRY"
	case CSSMERR_CSP_INVALID_NEW_ACL_OWNER:
		return "CSSMERR_CSP_INVALID_NEW_ACL_OWNER"
	case CSSMERR_CSP_INVALID_SAMPLE_VALUE:
		return "CSSMERR_CSP_INVALID_SAMPLE_VALUE"
	case CSSMERR_CSP_OBJECT_ACL_NOT_SUPPORTED:
		return "CSSMERR_CSP_OBJECT_ACL_NOT_SUPPORTED"
	case CSSMERR_CSP_OBJECT_ACL_REQUIRED:
		return "CSSMERR_CSP_OBJECT_ACL_REQUIRED"
	case CSSMERR_CSP_OBJECT_MANIP_AUTH_DENIED:
		return "CSSMERR_CSP_OBJECT_MANIP_AUTH_DENIED"
	case CSSMERR_CSP_OBJECT_USE_AUTH_DENIED:
		return "CSSMERR_CSP_OBJECT_USE_AUTH_DENIED"
	case CSSMERR_CSP_OPERATION_AUTH_DENIED:
		return "CSSMERR_CSP_OPERATION_AUTH_DENIED"
	case CSSMERR_CSP_SAMPLE_VALUE_NOT_SUPPORTED:
		return "CSSMERR_CSP_SAMPLE_VALUE_NOT_SUPPORTED"
	default:
		return fmt.Sprintf("CssmerrCspOperationAuthDenied(%d)", e)
	}
}

type CssmerrCssmInternalError int32

const (
	CSSMERR_CSSM_FUNCTION_FAILED               CssmerrCssmInternalError = -2147418102
	CSSMERR_CSSM_FUNCTION_NOT_IMPLEMENTED      CssmerrCssmInternalError = -2147418105
	CSSMERR_CSSM_INTERNAL_ERROR                CssmerrCssmInternalError = -2147418111
	CSSMERR_CSSM_INVALID_GUID                  CssmerrCssmInternalError = -2147418100
	CSSMERR_CSSM_INVALID_INPUT_POINTER         CssmerrCssmInternalError = -2147418107
	CSSMERR_CSSM_INVALID_OUTPUT_POINTER        CssmerrCssmInternalError = -2147418106
	CSSMERR_CSSM_INVALID_POINTER               CssmerrCssmInternalError = -2147418108
	CSSMERR_CSSM_MDS_ERROR                     CssmerrCssmInternalError = -2147418109
	CSSMERR_CSSM_MEMORY_ERROR                  CssmerrCssmInternalError = -2147418110
	CSSMERR_CSSM_MODULE_MANIFEST_VERIFY_FAILED CssmerrCssmInternalError = -2147418101
	CSSMERR_CSSM_OS_ACCESS_DENIED              CssmerrCssmInternalError = -2147418103
	CSSMERR_CSSM_SELF_CHECK_FAILED             CssmerrCssmInternalError = -2147418104
)

func (e CssmerrCssmInternalError) String() string {
	switch e {
	case CSSMERR_CSSM_FUNCTION_FAILED:
		return "CSSMERR_CSSM_FUNCTION_FAILED"
	case CSSMERR_CSSM_FUNCTION_NOT_IMPLEMENTED:
		return "CSSMERR_CSSM_FUNCTION_NOT_IMPLEMENTED"
	case CSSMERR_CSSM_INTERNAL_ERROR:
		return "CSSMERR_CSSM_INTERNAL_ERROR"
	case CSSMERR_CSSM_INVALID_GUID:
		return "CSSMERR_CSSM_INVALID_GUID"
	case CSSMERR_CSSM_INVALID_INPUT_POINTER:
		return "CSSMERR_CSSM_INVALID_INPUT_POINTER"
	case CSSMERR_CSSM_INVALID_OUTPUT_POINTER:
		return "CSSMERR_CSSM_INVALID_OUTPUT_POINTER"
	case CSSMERR_CSSM_INVALID_POINTER:
		return "CSSMERR_CSSM_INVALID_POINTER"
	case CSSMERR_CSSM_MDS_ERROR:
		return "CSSMERR_CSSM_MDS_ERROR"
	case CSSMERR_CSSM_MEMORY_ERROR:
		return "CSSMERR_CSSM_MEMORY_ERROR"
	case CSSMERR_CSSM_MODULE_MANIFEST_VERIFY_FAILED:
		return "CSSMERR_CSSM_MODULE_MANIFEST_VERIFY_FAILED"
	case CSSMERR_CSSM_OS_ACCESS_DENIED:
		return "CSSMERR_CSSM_OS_ACCESS_DENIED"
	case CSSMERR_CSSM_SELF_CHECK_FAILED:
		return "CSSMERR_CSSM_SELF_CHECK_FAILED"
	default:
		return fmt.Sprintf("CssmerrCssmInternalError(%d)", e)
	}
}

type CssmerrCssmInvalidAddinHandle int32

const (
	CSSMERR_CSSM_FUNCTION_INTEGRITY_FAIL CssmerrCssmInvalidAddinHandle = -2147417851
	CSSMERR_CSSM_INVALID_ADDIN_HANDLE    CssmerrCssmInvalidAddinHandle = -2147417855
	CSSMERR_CSSM_INVALID_HANDLE_USAGE    CssmerrCssmInvalidAddinHandle = -2147417853
	CSSMERR_CSSM_NOT_INITIALIZED         CssmerrCssmInvalidAddinHandle = -2147417854
	CSSMERR_CSSM_PVC_REFERENT_NOT_FOUND  CssmerrCssmInvalidAddinHandle = -2147417852
)

func (e CssmerrCssmInvalidAddinHandle) String() string {
	switch e {
	case CSSMERR_CSSM_FUNCTION_INTEGRITY_FAIL:
		return "CSSMERR_CSSM_FUNCTION_INTEGRITY_FAIL"
	case CSSMERR_CSSM_INVALID_ADDIN_HANDLE:
		return "CSSMERR_CSSM_INVALID_ADDIN_HANDLE"
	case CSSMERR_CSSM_INVALID_HANDLE_USAGE:
		return "CSSMERR_CSSM_INVALID_HANDLE_USAGE"
	case CSSMERR_CSSM_NOT_INITIALIZED:
		return "CSSMERR_CSSM_NOT_INITIALIZED"
	case CSSMERR_CSSM_PVC_REFERENT_NOT_FOUND:
		return "CSSMERR_CSSM_PVC_REFERENT_NOT_FOUND"
	default:
		return fmt.Sprintf("CssmerrCssmInvalidAddinHandle(%d)", e)
	}
}

type CssmerrCssmInvalidContextHandle int32

const (
	CSSMERR_CSSM_INCOMPATIBLE_VERSION   CssmerrCssmInvalidContextHandle = -2147418047
	CSSMERR_CSSM_INVALID_CONTEXT_HANDLE CssmerrCssmInvalidContextHandle = -2147418048
	CSSMERR_CSSM_PRIVILEGE_NOT_GRANTED  CssmerrCssmInvalidContextHandle = -2147418037
)

func (e CssmerrCssmInvalidContextHandle) String() string {
	switch e {
	case CSSMERR_CSSM_INCOMPATIBLE_VERSION:
		return "CSSMERR_CSSM_INCOMPATIBLE_VERSION"
	case CSSMERR_CSSM_INVALID_CONTEXT_HANDLE:
		return "CSSMERR_CSSM_INVALID_CONTEXT_HANDLE"
	case CSSMERR_CSSM_PRIVILEGE_NOT_GRANTED:
		return "CSSMERR_CSSM_PRIVILEGE_NOT_GRANTED"
	default:
		return fmt.Sprintf("CssmerrCssmInvalidContextHandle(%d)", e)
	}
}

type CssmerrDlInternalError int32

const (
	CSSMERR_DL_FUNCTION_FAILED          CssmerrDlInternalError = -2147414006
	CSSMERR_DL_FUNCTION_NOT_IMPLEMENTED CssmerrDlInternalError = -2147414009
	CSSMERR_DL_INTERNAL_ERROR           CssmerrDlInternalError = -2147414015
	CSSMERR_DL_INVALID_CL_HANDLE        CssmerrDlInternalError = -2147413934
	CSSMERR_DL_INVALID_CSP_HANDLE       CssmerrDlInternalError = -2147413936
	CSSMERR_DL_INVALID_DB_LIST_POINTER  CssmerrDlInternalError = -2147413939
	CSSMERR_DL_INVALID_DL_HANDLE        CssmerrDlInternalError = -2147413935
	CSSMERR_DL_INVALID_INPUT_POINTER    CssmerrDlInternalError = -2147414011
	CSSMERR_DL_INVALID_OUTPUT_POINTER   CssmerrDlInternalError = -2147414010
	CSSMERR_DL_INVALID_POINTER          CssmerrDlInternalError = -2147414012
	CSSMERR_DL_MDS_ERROR                CssmerrDlInternalError = -2147414013
	CSSMERR_DL_MEMORY_ERROR             CssmerrDlInternalError = -2147414014
	CSSMERR_DL_OS_ACCESS_DENIED         CssmerrDlInternalError = -2147414007
	CSSMERR_DL_SELF_CHECK_FAILED        CssmerrDlInternalError = -2147414008
)

func (e CssmerrDlInternalError) String() string {
	switch e {
	case CSSMERR_DL_FUNCTION_FAILED:
		return "CSSMERR_DL_FUNCTION_FAILED"
	case CSSMERR_DL_FUNCTION_NOT_IMPLEMENTED:
		return "CSSMERR_DL_FUNCTION_NOT_IMPLEMENTED"
	case CSSMERR_DL_INTERNAL_ERROR:
		return "CSSMERR_DL_INTERNAL_ERROR"
	case CSSMERR_DL_INVALID_CL_HANDLE:
		return "CSSMERR_DL_INVALID_CL_HANDLE"
	case CSSMERR_DL_INVALID_CSP_HANDLE:
		return "CSSMERR_DL_INVALID_CSP_HANDLE"
	case CSSMERR_DL_INVALID_DB_LIST_POINTER:
		return "CSSMERR_DL_INVALID_DB_LIST_POINTER"
	case CSSMERR_DL_INVALID_DL_HANDLE:
		return "CSSMERR_DL_INVALID_DL_HANDLE"
	case CSSMERR_DL_INVALID_INPUT_POINTER:
		return "CSSMERR_DL_INVALID_INPUT_POINTER"
	case CSSMERR_DL_INVALID_OUTPUT_POINTER:
		return "CSSMERR_DL_INVALID_OUTPUT_POINTER"
	case CSSMERR_DL_INVALID_POINTER:
		return "CSSMERR_DL_INVALID_POINTER"
	case CSSMERR_DL_MDS_ERROR:
		return "CSSMERR_DL_MDS_ERROR"
	case CSSMERR_DL_MEMORY_ERROR:
		return "CSSMERR_DL_MEMORY_ERROR"
	case CSSMERR_DL_OS_ACCESS_DENIED:
		return "CSSMERR_DL_OS_ACCESS_DENIED"
	case CSSMERR_DL_SELF_CHECK_FAILED:
		return "CSSMERR_DL_SELF_CHECK_FAILED"
	default:
		return fmt.Sprintf("CssmerrDlInternalError(%d)", e)
	}
}

type CssmerrDlInvalid int32

const (
	CSSMERR_DL_INVALID_DB_HANDLE      CssmerrDlInvalid = -2147413942
	CSSMERR_DL_INVALID_NETWORK_ADDR   CssmerrDlInvalid = -2147413929
	CSSMERR_DL_INVALID_PASSTHROUGH_ID CssmerrDlInvalid = -2147413930
)

func (e CssmerrDlInvalid) String() string {
	switch e {
	case CSSMERR_DL_INVALID_DB_HANDLE:
		return "CSSMERR_DL_INVALID_DB_HANDLE"
	case CSSMERR_DL_INVALID_NETWORK_ADDR:
		return "CSSMERR_DL_INVALID_NETWORK_ADDR"
	case CSSMERR_DL_INVALID_PASSTHROUGH_ID:
		return "CSSMERR_DL_INVALID_PASSTHROUGH_ID"
	default:
		return fmt.Sprintf("CssmerrDlInvalid(%d)", e)
	}
}

type CssmerrDlOperationAuthDenied int32

const (
	CSSMERR_DL_ACL_ADD_FAILED                 CssmerrDlOperationAuthDenied = -2147413962
	CSSMERR_DL_ACL_BASE_CERTS_NOT_SUPPORTED   CssmerrDlOperationAuthDenied = -2147413977
	CSSMERR_DL_ACL_CHALLENGE_CALLBACK_FAILED  CssmerrDlOperationAuthDenied = -2147413971
	CSSMERR_DL_ACL_CHANGE_FAILED              CssmerrDlOperationAuthDenied = -2147413967
	CSSMERR_DL_ACL_DELETE_FAILED              CssmerrDlOperationAuthDenied = -2147413964
	CSSMERR_DL_ACL_ENTRY_TAG_NOT_FOUND        CssmerrDlOperationAuthDenied = -2147413969
	CSSMERR_DL_ACL_REPLACE_FAILED             CssmerrDlOperationAuthDenied = -2147413963
	CSSMERR_DL_ACL_SUBJECT_TYPE_NOT_SUPPORTED CssmerrDlOperationAuthDenied = -2147413973
	CSSMERR_DL_INVALID_ACCESS_CREDENTIALS     CssmerrDlOperationAuthDenied = -2147413979
	CSSMERR_DL_INVALID_ACL_BASE_CERTS         CssmerrDlOperationAuthDenied = -2147413978
	CSSMERR_DL_INVALID_ACL_CHALLENGE_CALLBACK CssmerrDlOperationAuthDenied = -2147413972
	CSSMERR_DL_INVALID_ACL_EDIT_MODE          CssmerrDlOperationAuthDenied = -2147413968
	CSSMERR_DL_INVALID_ACL_ENTRY_TAG          CssmerrDlOperationAuthDenied = -2147413970
	CSSMERR_DL_INVALID_ACL_SUBJECT_VALUE      CssmerrDlOperationAuthDenied = -2147413974
	CSSMERR_DL_INVALID_NEW_ACL_ENTRY          CssmerrDlOperationAuthDenied = -2147413966
	CSSMERR_DL_INVALID_NEW_ACL_OWNER          CssmerrDlOperationAuthDenied = -2147413965
	CSSMERR_DL_INVALID_SAMPLE_VALUE           CssmerrDlOperationAuthDenied = -2147413976
	CSSMERR_DL_OBJECT_ACL_NOT_SUPPORTED       CssmerrDlOperationAuthDenied = -2147413981
	CSSMERR_DL_OBJECT_ACL_REQUIRED            CssmerrDlOperationAuthDenied = -2147413980
	CSSMERR_DL_OBJECT_MANIP_AUTH_DENIED       CssmerrDlOperationAuthDenied = -2147413982
	CSSMERR_DL_OBJECT_USE_AUTH_DENIED         CssmerrDlOperationAuthDenied = -2147413983
	CSSMERR_DL_OPERATION_AUTH_DENIED          CssmerrDlOperationAuthDenied = -2147413984
	CSSMERR_DL_SAMPLE_VALUE_NOT_SUPPORTED     CssmerrDlOperationAuthDenied = -2147413975
)

func (e CssmerrDlOperationAuthDenied) String() string {
	switch e {
	case CSSMERR_DL_ACL_ADD_FAILED:
		return "CSSMERR_DL_ACL_ADD_FAILED"
	case CSSMERR_DL_ACL_BASE_CERTS_NOT_SUPPORTED:
		return "CSSMERR_DL_ACL_BASE_CERTS_NOT_SUPPORTED"
	case CSSMERR_DL_ACL_CHALLENGE_CALLBACK_FAILED:
		return "CSSMERR_DL_ACL_CHALLENGE_CALLBACK_FAILED"
	case CSSMERR_DL_ACL_CHANGE_FAILED:
		return "CSSMERR_DL_ACL_CHANGE_FAILED"
	case CSSMERR_DL_ACL_DELETE_FAILED:
		return "CSSMERR_DL_ACL_DELETE_FAILED"
	case CSSMERR_DL_ACL_ENTRY_TAG_NOT_FOUND:
		return "CSSMERR_DL_ACL_ENTRY_TAG_NOT_FOUND"
	case CSSMERR_DL_ACL_REPLACE_FAILED:
		return "CSSMERR_DL_ACL_REPLACE_FAILED"
	case CSSMERR_DL_ACL_SUBJECT_TYPE_NOT_SUPPORTED:
		return "CSSMERR_DL_ACL_SUBJECT_TYPE_NOT_SUPPORTED"
	case CSSMERR_DL_INVALID_ACCESS_CREDENTIALS:
		return "CSSMERR_DL_INVALID_ACCESS_CREDENTIALS"
	case CSSMERR_DL_INVALID_ACL_BASE_CERTS:
		return "CSSMERR_DL_INVALID_ACL_BASE_CERTS"
	case CSSMERR_DL_INVALID_ACL_CHALLENGE_CALLBACK:
		return "CSSMERR_DL_INVALID_ACL_CHALLENGE_CALLBACK"
	case CSSMERR_DL_INVALID_ACL_EDIT_MODE:
		return "CSSMERR_DL_INVALID_ACL_EDIT_MODE"
	case CSSMERR_DL_INVALID_ACL_ENTRY_TAG:
		return "CSSMERR_DL_INVALID_ACL_ENTRY_TAG"
	case CSSMERR_DL_INVALID_ACL_SUBJECT_VALUE:
		return "CSSMERR_DL_INVALID_ACL_SUBJECT_VALUE"
	case CSSMERR_DL_INVALID_NEW_ACL_ENTRY:
		return "CSSMERR_DL_INVALID_NEW_ACL_ENTRY"
	case CSSMERR_DL_INVALID_NEW_ACL_OWNER:
		return "CSSMERR_DL_INVALID_NEW_ACL_OWNER"
	case CSSMERR_DL_INVALID_SAMPLE_VALUE:
		return "CSSMERR_DL_INVALID_SAMPLE_VALUE"
	case CSSMERR_DL_OBJECT_ACL_NOT_SUPPORTED:
		return "CSSMERR_DL_OBJECT_ACL_NOT_SUPPORTED"
	case CSSMERR_DL_OBJECT_ACL_REQUIRED:
		return "CSSMERR_DL_OBJECT_ACL_REQUIRED"
	case CSSMERR_DL_OBJECT_MANIP_AUTH_DENIED:
		return "CSSMERR_DL_OBJECT_MANIP_AUTH_DENIED"
	case CSSMERR_DL_OBJECT_USE_AUTH_DENIED:
		return "CSSMERR_DL_OBJECT_USE_AUTH_DENIED"
	case CSSMERR_DL_OPERATION_AUTH_DENIED:
		return "CSSMERR_DL_OPERATION_AUTH_DENIED"
	case CSSMERR_DL_SAMPLE_VALUE_NOT_SUPPORTED:
		return "CSSMERR_DL_SAMPLE_VALUE_NOT_SUPPORTED"
	default:
		return fmt.Sprintf("CssmerrDlOperationAuthDenied(%d)", e)
	}
}

type CssmerrTp int32

const (
	CSSMERR_TP_CRL_ALREADY_SIGNED        CssmerrTp = -2147409849
	CSSMERR_TP_FUNCTION_FAILED           CssmerrTp = -2147409910
	CSSMERR_TP_FUNCTION_NOT_IMPLEMENTED  CssmerrTp = -2147409913
	CSSMERR_TP_INTERNAL_ERROR            CssmerrTp = -2147409919
	CSSMERR_TP_INVALID_CERTGROUP_POINTER CssmerrTp = -2147409854
	CSSMERR_TP_INVALID_CERT_POINTER      CssmerrTp = -2147409853
	CSSMERR_TP_INVALID_CL_HANDLE         CssmerrTp = -2147409838
	CSSMERR_TP_INVALID_CONTEXT_HANDLE    CssmerrTp = -2147409856
	CSSMERR_TP_INVALID_CRL_POINTER       CssmerrTp = -2147409852
	CSSMERR_TP_INVALID_CSP_HANDLE        CssmerrTp = -2147409840
	CSSMERR_TP_INVALID_DATA              CssmerrTp = -2147409850
	CSSMERR_TP_INVALID_DB_HANDLE         CssmerrTp = -2147409846
	CSSMERR_TP_INVALID_DB_LIST           CssmerrTp = -2147409844
	CSSMERR_TP_INVALID_DB_LIST_POINTER   CssmerrTp = -2147409843
	CSSMERR_TP_INVALID_DL_HANDLE         CssmerrTp = -2147409839
	CSSMERR_TP_INVALID_FIELD_POINTER     CssmerrTp = -2147409851
	CSSMERR_TP_INVALID_INPUT_POINTER     CssmerrTp = -2147409915
	CSSMERR_TP_INVALID_NETWORK_ADDR      CssmerrTp = -2147409833
	CSSMERR_TP_INVALID_NUMBER_OF_FIELDS  CssmerrTp = -2147409848
	CSSMERR_TP_INVALID_OUTPUT_POINTER    CssmerrTp = -2147409914
	CSSMERR_TP_INVALID_PASSTHROUGH_ID    CssmerrTp = -2147409834
	CSSMERR_TP_INVALID_POINTER           CssmerrTp = -2147409916
	CSSMERR_TP_MDS_ERROR                 CssmerrTp = -2147409917
	CSSMERR_TP_MEMORY_ERROR              CssmerrTp = -2147409918
	CSSMERR_TP_OS_ACCESS_DENIED          CssmerrTp = -2147409911
	CSSMERR_TP_SELF_CHECK_FAILED         CssmerrTp = -2147409912
	CSSMERR_TP_UNKNOWN_FORMAT            CssmerrTp = -2147409842
	CSSMERR_TP_UNKNOWN_TAG               CssmerrTp = -2147409841
	CSSMERR_TP_VERIFICATION_FAILURE      CssmerrTp = -2147409847
)

func (e CssmerrTp) String() string {
	switch e {
	case CSSMERR_TP_CRL_ALREADY_SIGNED:
		return "CSSMERR_TP_CRL_ALREADY_SIGNED"
	case CSSMERR_TP_FUNCTION_FAILED:
		return "CSSMERR_TP_FUNCTION_FAILED"
	case CSSMERR_TP_FUNCTION_NOT_IMPLEMENTED:
		return "CSSMERR_TP_FUNCTION_NOT_IMPLEMENTED"
	case CSSMERR_TP_INTERNAL_ERROR:
		return "CSSMERR_TP_INTERNAL_ERROR"
	case CSSMERR_TP_INVALID_CERTGROUP_POINTER:
		return "CSSMERR_TP_INVALID_CERTGROUP_POINTER"
	case CSSMERR_TP_INVALID_CERT_POINTER:
		return "CSSMERR_TP_INVALID_CERT_POINTER"
	case CSSMERR_TP_INVALID_CL_HANDLE:
		return "CSSMERR_TP_INVALID_CL_HANDLE"
	case CSSMERR_TP_INVALID_CONTEXT_HANDLE:
		return "CSSMERR_TP_INVALID_CONTEXT_HANDLE"
	case CSSMERR_TP_INVALID_CRL_POINTER:
		return "CSSMERR_TP_INVALID_CRL_POINTER"
	case CSSMERR_TP_INVALID_CSP_HANDLE:
		return "CSSMERR_TP_INVALID_CSP_HANDLE"
	case CSSMERR_TP_INVALID_DATA:
		return "CSSMERR_TP_INVALID_DATA"
	case CSSMERR_TP_INVALID_DB_HANDLE:
		return "CSSMERR_TP_INVALID_DB_HANDLE"
	case CSSMERR_TP_INVALID_DB_LIST:
		return "CSSMERR_TP_INVALID_DB_LIST"
	case CSSMERR_TP_INVALID_DB_LIST_POINTER:
		return "CSSMERR_TP_INVALID_DB_LIST_POINTER"
	case CSSMERR_TP_INVALID_DL_HANDLE:
		return "CSSMERR_TP_INVALID_DL_HANDLE"
	case CSSMERR_TP_INVALID_FIELD_POINTER:
		return "CSSMERR_TP_INVALID_FIELD_POINTER"
	case CSSMERR_TP_INVALID_INPUT_POINTER:
		return "CSSMERR_TP_INVALID_INPUT_POINTER"
	case CSSMERR_TP_INVALID_NETWORK_ADDR:
		return "CSSMERR_TP_INVALID_NETWORK_ADDR"
	case CSSMERR_TP_INVALID_NUMBER_OF_FIELDS:
		return "CSSMERR_TP_INVALID_NUMBER_OF_FIELDS"
	case CSSMERR_TP_INVALID_OUTPUT_POINTER:
		return "CSSMERR_TP_INVALID_OUTPUT_POINTER"
	case CSSMERR_TP_INVALID_PASSTHROUGH_ID:
		return "CSSMERR_TP_INVALID_PASSTHROUGH_ID"
	case CSSMERR_TP_INVALID_POINTER:
		return "CSSMERR_TP_INVALID_POINTER"
	case CSSMERR_TP_MDS_ERROR:
		return "CSSMERR_TP_MDS_ERROR"
	case CSSMERR_TP_MEMORY_ERROR:
		return "CSSMERR_TP_MEMORY_ERROR"
	case CSSMERR_TP_OS_ACCESS_DENIED:
		return "CSSMERR_TP_OS_ACCESS_DENIED"
	case CSSMERR_TP_SELF_CHECK_FAILED:
		return "CSSMERR_TP_SELF_CHECK_FAILED"
	case CSSMERR_TP_UNKNOWN_FORMAT:
		return "CSSMERR_TP_UNKNOWN_FORMAT"
	case CSSMERR_TP_UNKNOWN_TAG:
		return "CSSMERR_TP_UNKNOWN_TAG"
	case CSSMERR_TP_VERIFICATION_FAILURE:
		return "CSSMERR_TP_VERIFICATION_FAILURE"
	default:
		return fmt.Sprintf("CssmerrTp(%d)", e)
	}
}

type ErrAuthorization int32

const (
	// ErrAuthorizationBadAddress: The requested socket address is invalid.
	ErrAuthorizationBadAddress ErrAuthorization = -60033
	// ErrAuthorizationCanceled: The user canceled the operation.
	ErrAuthorizationCanceled ErrAuthorization = -60006
	// ErrAuthorizationDenied: The Security Server denied authorization for one or more requested rights.
	ErrAuthorizationDenied ErrAuthorization = -60005
	// ErrAuthorizationExternalizeNotAllowed: The Security Server denied externalization of the authorization reference.
	ErrAuthorizationExternalizeNotAllowed ErrAuthorization = -60009
	// ErrAuthorizationInteractionNotAllowed: The Security Server denied authorization because no user interaction is allowed.
	ErrAuthorizationInteractionNotAllowed ErrAuthorization = -60007
	// ErrAuthorizationInternal: An unrecognized internal error occurred.
	ErrAuthorizationInternal ErrAuthorization = -60008
	// ErrAuthorizationInternalizeNotAllowed: The Security Server denied internalization of the authorization reference.
	ErrAuthorizationInternalizeNotAllowed ErrAuthorization = -60010
	// ErrAuthorizationInvalidFlags: The flags parameter is invalid.
	ErrAuthorizationInvalidFlags ErrAuthorization = -60011
	// ErrAuthorizationInvalidPointer: The authorizedRights parameter is invalid.
	ErrAuthorizationInvalidPointer ErrAuthorization = -60004
	// ErrAuthorizationInvalidRef: The authorization parameter is invalid.
	ErrAuthorizationInvalidRef ErrAuthorization = -60002
	// ErrAuthorizationInvalidSet: The set parameter is invalid.
	ErrAuthorizationInvalidSet ErrAuthorization = -60001
	// ErrAuthorizationInvalidTag: The tag parameter is invalid.
	ErrAuthorizationInvalidTag ErrAuthorization = -60003
	// ErrAuthorizationSuccess: The operation completed successfully.
	ErrAuthorizationSuccess ErrAuthorization = 0
	// ErrAuthorizationToolEnvironmentError: The attempt to execute the tool failed to return a success or an error code.
	ErrAuthorizationToolEnvironmentError ErrAuthorization = -60032
	// ErrAuthorizationToolExecuteFailure: The tool failed to execute.
	ErrAuthorizationToolExecuteFailure ErrAuthorization = -60031
)

func (e ErrAuthorization) String() string {
	switch e {
	case ErrAuthorizationBadAddress:
		return "ErrAuthorizationBadAddress"
	case ErrAuthorizationCanceled:
		return "ErrAuthorizationCanceled"
	case ErrAuthorizationDenied:
		return "ErrAuthorizationDenied"
	case ErrAuthorizationExternalizeNotAllowed:
		return "ErrAuthorizationExternalizeNotAllowed"
	case ErrAuthorizationInteractionNotAllowed:
		return "ErrAuthorizationInteractionNotAllowed"
	case ErrAuthorizationInternal:
		return "ErrAuthorizationInternal"
	case ErrAuthorizationInternalizeNotAllowed:
		return "ErrAuthorizationInternalizeNotAllowed"
	case ErrAuthorizationInvalidFlags:
		return "ErrAuthorizationInvalidFlags"
	case ErrAuthorizationInvalidPointer:
		return "ErrAuthorizationInvalidPointer"
	case ErrAuthorizationInvalidRef:
		return "ErrAuthorizationInvalidRef"
	case ErrAuthorizationInvalidSet:
		return "ErrAuthorizationInvalidSet"
	case ErrAuthorizationInvalidTag:
		return "ErrAuthorizationInvalidTag"
	case ErrAuthorizationSuccess:
		return "ErrAuthorizationSuccess"
	case ErrAuthorizationToolEnvironmentError:
		return "ErrAuthorizationToolEnvironmentError"
	case ErrAuthorizationToolExecuteFailure:
		return "ErrAuthorizationToolExecuteFailure"
	default:
		return fmt.Sprintf("ErrAuthorization(%d)", e)
	}
}

type ErrSSL int32

const (
	// ErrSSLATSCertificateHashAlgorithmViolation: The peer certificate hash algorithm isn’t App Transport Security compliant.
	ErrSSLATSCertificateHashAlgorithmViolation ErrSSL = -9885
	// ErrSSLATSCertificateTrustViolation: The peer certificate wasn’t issued by a trusted peer.
	ErrSSLATSCertificateTrustViolation ErrSSL = -9886
	// ErrSSLATSCiphersuiteViolation: The selected ciphersuite isn’t App Transport Security compliant.
	ErrSSLATSCiphersuiteViolation ErrSSL = -9882
	// ErrSSLATSLeafCertificateHashAlgorithmViolation: The peer leaf certificate hash algorithm isn’t App Transport Security compliant.
	ErrSSLATSLeafCertificateHashAlgorithmViolation ErrSSL = -9884
	// ErrSSLATSMinimumKeySizeViolation: The peer key size isn’t App Transport Security compliant.
	ErrSSLATSMinimumKeySizeViolation ErrSSL = -9883
	// ErrSSLATSMinimumVersionViolation: The minimum protocol version isn’t App Transport Security compliant.
	ErrSSLATSMinimumVersionViolation ErrSSL = -9881
	// ErrSSLATSViolation: An App Transport Security violation occurred.
	ErrSSLATSViolation ErrSSL = -9880
	// ErrSSLBadCert: Bad certificate format.
	ErrSSLBadCert ErrSSL = -9808
	// ErrSSLBadCertificateStatusResponse: Bad OCSP response.
	ErrSSLBadCertificateStatusResponse ErrSSL = -9862
	// ErrSSLBadCipherSuite: A bad SSL cipher suite was encountered.
	ErrSSLBadCipherSuite ErrSSL = -9818
	// ErrSSLBadConfiguration: A configuration error occurred.
	ErrSSLBadConfiguration ErrSSL = -9848
	// ErrSSLBadRecordMac: A record with a bad message authentication code (MAC) was encountered.
	ErrSSLBadRecordMac ErrSSL = -9846
	// ErrSSLBufferOverflow: An insufficient buffer was provided.
	ErrSSLBufferOverflow ErrSSL = -9817
	// ErrSSLCertExpired: The certificate chain had an expired certificate.
	ErrSSLCertExpired ErrSSL = -9814
	// ErrSSLCertNotYetValid: The certificate chain had a certificatethat is not yet valid.
	ErrSSLCertNotYetValid ErrSSL = -9815
	// ErrSSLCertificateRequired: Certificate required.
	ErrSSLCertificateRequired ErrSSL = -9863
	// ErrSSLClientCertRequested: The server has requested a client certificate.
	ErrSSLClientCertRequested ErrSSL = -9842
	// ErrSSLClientHelloReceived: A non-fatal result for providing a server name indication.
	ErrSSLClientHelloReceived ErrSSL = -9851
	// ErrSSLClosedAbort: The connection closed due to an error.
	ErrSSLClosedAbort ErrSSL = -9806
	// ErrSSLClosedGraceful: The connection closed gracefully.
	ErrSSLClosedGraceful ErrSSL = -9805
	// ErrSSLClosedNoNotify: The server closed the session with no notification.
	ErrSSLClosedNoNotify ErrSSL = -9816
	// ErrSSLConfigurationFailed: TLS configuration failed.
	ErrSSLConfigurationFailed ErrSSL = -9854
	// ErrSSLConnectionRefused: The peer dropped the connection before responding.
	ErrSSLConnectionRefused ErrSSL = -9844
	// ErrSSLCrypto: An underlying cryptographic error was encountered.
	ErrSSLCrypto ErrSSL = -9809
	// ErrSSLDecodeError: Decode failed.
	ErrSSLDecodeError ErrSSL = -9859
	// ErrSSLDecompressFail: Decompression failed.
	ErrSSLDecompressFail ErrSSL = -9857
	// ErrSSLDecryptionFail: Decryption failed.
	ErrSSLDecryptionFail    ErrSSL = -9845
	ErrSSLEarlyDataRejected ErrSSL = -9890
	// ErrSSLFatalAlert: A fatal alert was encountered.
	ErrSSLFatalAlert ErrSSL = -9802
	// ErrSSLHandshakeFail: Handshake failed.
	ErrSSLHandshakeFail ErrSSL = -9858
	// ErrSSLHostNameMismatch: The host name you connected with does not match any of the host names allowed by the certificate.
	ErrSSLHostNameMismatch ErrSSL = -9843
	// ErrSSLIllegalParam: An illegal parameter was encountered.
	ErrSSLIllegalParam ErrSSL = -9830
	// ErrSSLInappropriateFallback: Inappropriate fallback.
	ErrSSLInappropriateFallback ErrSSL = -9860
	// ErrSSLInternal: Internal error.
	ErrSSLInternal ErrSSL = -9810
	// ErrSSLMissingExtension: Missing extension.
	ErrSSLMissingExtension ErrSSL = -9861
	// ErrSSLModuleAttach: Module attach failure.
	ErrSSLModuleAttach ErrSSL = -9811
	// ErrSSLNegotiation: The cipher suite negotiation failed.
	ErrSSLNegotiation ErrSSL = -9801
	// ErrSSLNetworkTimeout: Network timeout triggered.
	ErrSSLNetworkTimeout ErrSSL = -9853
	// ErrSSLNoRootCert: No root certificate for the certificate chain.
	ErrSSLNoRootCert ErrSSL = -9813
	// ErrSSLPeerAccessDenied: Access was denied.
	ErrSSLPeerAccessDenied ErrSSL = -9832
	// ErrSSLPeerAuthCompleted: A non-fatal result indicating the peer certificate is valid, or was ignored if verification is disabled.
	ErrSSLPeerAuthCompleted ErrSSL = -9841
	// ErrSSLPeerBadCert: A bad certificate was encountered.
	ErrSSLPeerBadCert ErrSSL = -9825
	// ErrSSLPeerBadRecordMac: A record with a bad message authentication code (MAC) was encountered.
	ErrSSLPeerBadRecordMac ErrSSL = -9820
	// ErrSSLPeerCertExpired: The certificate expired.
	ErrSSLPeerCertExpired ErrSSL = -9828
	// ErrSSLPeerCertRevoked: The certificate was revoked.
	ErrSSLPeerCertRevoked ErrSSL = -9827
	// ErrSSLPeerCertUnknown: The certificate is unknown.
	ErrSSLPeerCertUnknown ErrSSL = -9829
	// ErrSSLPeerDecodeError: A decoding error occurred.
	ErrSSLPeerDecodeError ErrSSL = -9833
	// ErrSSLPeerDecompressFail: Decompression failed.
	ErrSSLPeerDecompressFail ErrSSL = -9823
	// ErrSSLPeerDecryptError: A decryption error occurred.
	ErrSSLPeerDecryptError ErrSSL = -9834
	// ErrSSLPeerDecryptionFail: Decryption failed.
	ErrSSLPeerDecryptionFail ErrSSL = -9821
	// ErrSSLPeerExportRestriction: An export restriction occurred.
	ErrSSLPeerExportRestriction ErrSSL = -9835
	// ErrSSLPeerHandshakeFail: The handshake failed.
	ErrSSLPeerHandshakeFail ErrSSL = -9824
	// ErrSSLPeerInsufficientSecurity: There is insufficient security for this operation.
	ErrSSLPeerInsufficientSecurity ErrSSL = -9837
	// ErrSSLPeerInternalError: An internal error occurred.
	ErrSSLPeerInternalError ErrSSL = -9838
	// ErrSSLPeerNoRenegotiation: No renegotiation is allowed.
	ErrSSLPeerNoRenegotiation ErrSSL = -9840
	// ErrSSLPeerProtocolVersion: A bad protocol version was encountered.
	ErrSSLPeerProtocolVersion ErrSSL = -9836
	// ErrSSLPeerRecordOverflow: A record overflow occurred.
	ErrSSLPeerRecordOverflow ErrSSL = -9822
	// ErrSSLPeerUnexpectedMsg: An unexpected message was received.
	ErrSSLPeerUnexpectedMsg ErrSSL = -9819
	// ErrSSLPeerUnknownCA: An unknown certificate authority was encountered.
	ErrSSLPeerUnknownCA ErrSSL = -9831
	// ErrSSLPeerUnsupportedCert: An unsupported certificate format was encountered.
	ErrSSLPeerUnsupportedCert ErrSSL = -9826
	// ErrSSLPeerUserCancelled: The user canceled the operation.
	ErrSSLPeerUserCancelled ErrSSL = -9839
	// ErrSSLProtocol: SSL protocol error.
	ErrSSLProtocol ErrSSL = -9800
	// ErrSSLRecordOverflow: A record overflow occurred.
	ErrSSLRecordOverflow ErrSSL = -9847
	// ErrSSLSessionNotFound: An attempt to restore an unknown session failed.
	ErrSSLSessionNotFound ErrSSL = -9804
	// ErrSSLTransportReset: Transport (socket) shutdown, for example, TCP RST or FIN.
	ErrSSLTransportReset ErrSSL = -9852
	// ErrSSLUnexpectedMessage: Peer rejected unexpected message.
	ErrSSLUnexpectedMessage ErrSSL = -9856
	ErrSSLUnexpectedRecord  ErrSSL = -9849
	// ErrSSLUnknownPSKIdentity: Unknown PSK identity.
	ErrSSLUnknownPSKIdentity ErrSSL = -9864
	// ErrSSLUnknownRootCert: Certificate chain is valid, but root is nottrusted.
	ErrSSLUnknownRootCert ErrSSL = -9812
	// ErrSSLUnrecognizedName: Unknown or unrecognized name.
	ErrSSLUnrecognizedName ErrSSL = -9865
	// ErrSSLUnsupportedExtension: Unsupported TLS extension.
	ErrSSLUnsupportedExtension ErrSSL = -9855
	// ErrSSLWeakPeerEphemeralDHKey: Indicates a weak ephemeral dh key.
	ErrSSLWeakPeerEphemeralDHKey ErrSSL = -9850
	// ErrSSLWouldBlock: Function is blocked; waiting for I/O.
	ErrSSLWouldBlock ErrSSL = -9803
	// ErrSSLXCertChainInvalid: Invalid certificate chain.
	ErrSSLXCertChainInvalid ErrSSL = -9807
)

func (e ErrSSL) String() string {
	switch e {
	case ErrSSLATSCertificateHashAlgorithmViolation:
		return "ErrSSLATSCertificateHashAlgorithmViolation"
	case ErrSSLATSCertificateTrustViolation:
		return "ErrSSLATSCertificateTrustViolation"
	case ErrSSLATSCiphersuiteViolation:
		return "ErrSSLATSCiphersuiteViolation"
	case ErrSSLATSLeafCertificateHashAlgorithmViolation:
		return "ErrSSLATSLeafCertificateHashAlgorithmViolation"
	case ErrSSLATSMinimumKeySizeViolation:
		return "ErrSSLATSMinimumKeySizeViolation"
	case ErrSSLATSMinimumVersionViolation:
		return "ErrSSLATSMinimumVersionViolation"
	case ErrSSLATSViolation:
		return "ErrSSLATSViolation"
	case ErrSSLBadCert:
		return "ErrSSLBadCert"
	case ErrSSLBadCertificateStatusResponse:
		return "ErrSSLBadCertificateStatusResponse"
	case ErrSSLBadCipherSuite:
		return "ErrSSLBadCipherSuite"
	case ErrSSLBadConfiguration:
		return "ErrSSLBadConfiguration"
	case ErrSSLBadRecordMac:
		return "ErrSSLBadRecordMac"
	case ErrSSLBufferOverflow:
		return "ErrSSLBufferOverflow"
	case ErrSSLCertExpired:
		return "ErrSSLCertExpired"
	case ErrSSLCertNotYetValid:
		return "ErrSSLCertNotYetValid"
	case ErrSSLCertificateRequired:
		return "ErrSSLCertificateRequired"
	case ErrSSLClientCertRequested:
		return "ErrSSLClientCertRequested"
	case ErrSSLClientHelloReceived:
		return "ErrSSLClientHelloReceived"
	case ErrSSLClosedAbort:
		return "ErrSSLClosedAbort"
	case ErrSSLClosedGraceful:
		return "ErrSSLClosedGraceful"
	case ErrSSLClosedNoNotify:
		return "ErrSSLClosedNoNotify"
	case ErrSSLConfigurationFailed:
		return "ErrSSLConfigurationFailed"
	case ErrSSLConnectionRefused:
		return "ErrSSLConnectionRefused"
	case ErrSSLCrypto:
		return "ErrSSLCrypto"
	case ErrSSLDecodeError:
		return "ErrSSLDecodeError"
	case ErrSSLDecompressFail:
		return "ErrSSLDecompressFail"
	case ErrSSLDecryptionFail:
		return "ErrSSLDecryptionFail"
	case ErrSSLEarlyDataRejected:
		return "ErrSSLEarlyDataRejected"
	case ErrSSLFatalAlert:
		return "ErrSSLFatalAlert"
	case ErrSSLHandshakeFail:
		return "ErrSSLHandshakeFail"
	case ErrSSLHostNameMismatch:
		return "ErrSSLHostNameMismatch"
	case ErrSSLIllegalParam:
		return "ErrSSLIllegalParam"
	case ErrSSLInappropriateFallback:
		return "ErrSSLInappropriateFallback"
	case ErrSSLInternal:
		return "ErrSSLInternal"
	case ErrSSLMissingExtension:
		return "ErrSSLMissingExtension"
	case ErrSSLModuleAttach:
		return "ErrSSLModuleAttach"
	case ErrSSLNegotiation:
		return "ErrSSLNegotiation"
	case ErrSSLNetworkTimeout:
		return "ErrSSLNetworkTimeout"
	case ErrSSLNoRootCert:
		return "ErrSSLNoRootCert"
	case ErrSSLPeerAccessDenied:
		return "ErrSSLPeerAccessDenied"
	case ErrSSLPeerAuthCompleted:
		return "ErrSSLPeerAuthCompleted"
	case ErrSSLPeerBadCert:
		return "ErrSSLPeerBadCert"
	case ErrSSLPeerBadRecordMac:
		return "ErrSSLPeerBadRecordMac"
	case ErrSSLPeerCertExpired:
		return "ErrSSLPeerCertExpired"
	case ErrSSLPeerCertRevoked:
		return "ErrSSLPeerCertRevoked"
	case ErrSSLPeerCertUnknown:
		return "ErrSSLPeerCertUnknown"
	case ErrSSLPeerDecodeError:
		return "ErrSSLPeerDecodeError"
	case ErrSSLPeerDecompressFail:
		return "ErrSSLPeerDecompressFail"
	case ErrSSLPeerDecryptError:
		return "ErrSSLPeerDecryptError"
	case ErrSSLPeerDecryptionFail:
		return "ErrSSLPeerDecryptionFail"
	case ErrSSLPeerExportRestriction:
		return "ErrSSLPeerExportRestriction"
	case ErrSSLPeerHandshakeFail:
		return "ErrSSLPeerHandshakeFail"
	case ErrSSLPeerInsufficientSecurity:
		return "ErrSSLPeerInsufficientSecurity"
	case ErrSSLPeerInternalError:
		return "ErrSSLPeerInternalError"
	case ErrSSLPeerNoRenegotiation:
		return "ErrSSLPeerNoRenegotiation"
	case ErrSSLPeerProtocolVersion:
		return "ErrSSLPeerProtocolVersion"
	case ErrSSLPeerRecordOverflow:
		return "ErrSSLPeerRecordOverflow"
	case ErrSSLPeerUnexpectedMsg:
		return "ErrSSLPeerUnexpectedMsg"
	case ErrSSLPeerUnknownCA:
		return "ErrSSLPeerUnknownCA"
	case ErrSSLPeerUnsupportedCert:
		return "ErrSSLPeerUnsupportedCert"
	case ErrSSLPeerUserCancelled:
		return "ErrSSLPeerUserCancelled"
	case ErrSSLProtocol:
		return "ErrSSLProtocol"
	case ErrSSLRecordOverflow:
		return "ErrSSLRecordOverflow"
	case ErrSSLSessionNotFound:
		return "ErrSSLSessionNotFound"
	case ErrSSLTransportReset:
		return "ErrSSLTransportReset"
	case ErrSSLUnexpectedMessage:
		return "ErrSSLUnexpectedMessage"
	case ErrSSLUnexpectedRecord:
		return "ErrSSLUnexpectedRecord"
	case ErrSSLUnknownPSKIdentity:
		return "ErrSSLUnknownPSKIdentity"
	case ErrSSLUnknownRootCert:
		return "ErrSSLUnknownRootCert"
	case ErrSSLUnrecognizedName:
		return "ErrSSLUnrecognizedName"
	case ErrSSLUnsupportedExtension:
		return "ErrSSLUnsupportedExtension"
	case ErrSSLWeakPeerEphemeralDHKey:
		return "ErrSSLWeakPeerEphemeralDHKey"
	case ErrSSLWouldBlock:
		return "ErrSSLWouldBlock"
	case ErrSSLXCertChainInvalid:
		return "ErrSSLXCertChainInvalid"
	default:
		return fmt.Sprintf("ErrSSL(%d)", e)
	}
}

type ErrSecCSUnimplemented int32

const (
	// ErrSecCSAmbiguousBundleFormat: The bundle could be an app or a framework.
	ErrSecCSAmbiguousBundleFormat ErrSecCSUnimplemented = -67011
	// ErrSecCSBadBundleFormat: The bundle format is unrecognized, invalid, or unsuitable.
	ErrSecCSBadBundleFormat ErrSecCSUnimplemented = -67028
	// ErrSecCSBadCallbackValue: The monitor callback returned invalid value.
	ErrSecCSBadCallbackValue ErrSecCSUnimplemented = -67020
	// ErrSecCSBadDictionaryFormat: A required information property list (Info.plist) file or resource is malformed.
	ErrSecCSBadDictionaryFormat ErrSecCSUnimplemented = -67058
	// ErrSecCSBadDiskImageFormat: The disk image format unrecognized, invalid, or unsuitable.
	ErrSecCSBadDiskImageFormat ErrSecCSUnimplemented = -67001
	// ErrSecCSBadFrameworkVersion: The embedded framework contains a modified or invalid version.
	ErrSecCSBadFrameworkVersion ErrSecCSUnimplemented = -67009
	// ErrSecCSBadLVArch: The library validation flag cannot be used with an i386 binary.
	ErrSecCSBadLVArch ErrSecCSUnimplemented = -67017
	// ErrSecCSBadMainExecutable: The main executable failed strict validation.
	ErrSecCSBadMainExecutable ErrSecCSUnimplemented = -67010
	// ErrSecCSBadNestedCode: The nested code is modified or invalid.
	ErrSecCSBadNestedCode ErrSecCSUnimplemented = -67021
	// ErrSecCSBadObjectFormat: The object file format invalid or unsuitable.
	ErrSecCSBadObjectFormat ErrSecCSUnimplemented = -67049
	// ErrSecCSBadResource: A sealed resource is missing or invalid.
	ErrSecCSBadResource ErrSecCSUnimplemented = -67054
	// ErrSecCSBadTeamIdentifier: A Team Identifier is wrong or inappropriate.
	ErrSecCSBadTeamIdentifier     ErrSecCSUnimplemented = -66997
	ErrSecCSCMSConstructionFailed ErrSecCSUnimplemented = -66991
	// ErrSecCSCMSTooLarge: The signature is too large to embed.
	ErrSecCSCMSTooLarge ErrSecCSUnimplemented = -67036
	// ErrSecCSCancelled: The operation was terminated by explicit cancellation.
	ErrSecCSCancelled ErrSecCSUnimplemented = -67006
	// ErrSecCSDBAccess: Cannot access signature database.
	ErrSecCSDBAccess ErrSecCSUnimplemented = -67032
	// ErrSecCSDBDenied: Access to signature database denied.
	ErrSecCSDBDenied ErrSecCSUnimplemented = -67033
	// ErrSecCSDSStoreSymlink: A `.DS_Store` file can’t be a symlink.
	ErrSecCSDSStoreSymlink ErrSecCSUnimplemented = -67012
	// ErrSecCSDbCorrupt: A system database or file is corrupt.
	ErrSecCSDbCorrupt ErrSecCSUnimplemented = -67024
	// ErrSecCSFileHardQuarantined: File open or execution not allowed.
	ErrSecCSFileHardQuarantined ErrSecCSUnimplemented = -67026
	// ErrSecCSGuestInvalid: The identity of guest code has been invalidated.
	ErrSecCSGuestInvalid ErrSecCSUnimplemented = -67063
	// ErrSecCSHelperFailed: The codesign_allocate helper tool can’t be found or used.
	ErrSecCSHelperFailed ErrSecCSUnimplemented = -67019
	// ErrSecCSHostProtocolContradiction: Host protocol violation: contradictory hosting modes.
	ErrSecCSHostProtocolContradiction ErrSecCSUnimplemented = -67043
	// ErrSecCSHostProtocolDedicationError: Host protocol violation: operation not allowed with or for a dedicated guest.
	ErrSecCSHostProtocolDedicationError ErrSecCSUnimplemented = -67042
	// ErrSecCSHostProtocolInvalidAttribute: Code signing host returned invalid or inconsistent attributes for guest code.
	ErrSecCSHostProtocolInvalidAttribute ErrSecCSUnimplemented = -67031
	// ErrSecCSHostProtocolInvalidHash: Host protocol violation: invalid hash of guest code.
	ErrSecCSHostProtocolInvalidHash ErrSecCSUnimplemented = -67035
	// ErrSecCSHostProtocolNotProxy: Host protocol violation: proxy hosting not engaged.
	ErrSecCSHostProtocolNotProxy ErrSecCSUnimplemented = -67041
	// ErrSecCSHostProtocolRelativePath: Host protocol violation: absolute guest path required.
	ErrSecCSHostProtocolRelativePath ErrSecCSUnimplemented = -67044
	// ErrSecCSHostProtocolStateError: Host protocol violation: invalid guest state change request.
	ErrSecCSHostProtocolStateError ErrSecCSUnimplemented = -67040
	// ErrSecCSHostProtocolUnrelated: Host protocol violation: the specified code is not a guest of the specified code signing host.
	ErrSecCSHostProtocolUnrelated ErrSecCSUnimplemented = -67039
	// ErrSecCSHostReject: Code rejected its host.
	ErrSecCSHostReject ErrSecCSUnimplemented = -67047
	// ErrSecCSInfoPlistFailed: The Info.plist file or the signature has been modified.
	ErrSecCSInfoPlistFailed ErrSecCSUnimplemented = -67030
	// ErrSecCSInternalError: Internal error in Code Signing Services subsystem.
	ErrSecCSInternalError ErrSecCSUnimplemented = -67048
	// ErrSecCSInvalidAssociatedFileData: Resource fork, Finder information, or similar detritus not allowed.
	ErrSecCSInvalidAssociatedFileData ErrSecCSUnimplemented = -66999
	// ErrSecCSInvalidAttributeValues: An attribute value associated with a key is out of range or is the wrong type.
	ErrSecCSInvalidAttributeValues ErrSecCSUnimplemented = -67066
	// ErrSecCSInvalidEntitlements: Encountered an invalid entitlement plist.
	ErrSecCSInvalidEntitlements ErrSecCSUnimplemented = -66994
	// ErrSecCSInvalidFlags: Invalid or inappropriate API flags specified.
	ErrSecCSInvalidFlags ErrSecCSUnimplemented = -67070
	// ErrSecCSInvalidObjectRef: Invalid API object reference.
	ErrSecCSInvalidObjectRef ErrSecCSUnimplemented = -67071
	// ErrSecCSInvalidPlatform: Invalid platform identifier or platform mismatch.
	ErrSecCSInvalidPlatform ErrSecCSUnimplemented = -67005
	// ErrSecCSInvalidRuntimeVersion: An invalid runtime version was explicity set.
	ErrSecCSInvalidRuntimeVersion ErrSecCSUnimplemented = -66993
	// ErrSecCSInvalidSymlink: Invalid destination for symbolic link in bundle.
	ErrSecCSInvalidSymlink ErrSecCSUnimplemented = -67003
	// ErrSecCSInvalidTeamIdentifier: A Team Identifier string is invalid.
	ErrSecCSInvalidTeamIdentifier ErrSecCSUnimplemented = -66998
	// ErrSecCSMultipleGuests: Code signing host has more than one block of guest code with this attribute value.
	ErrSecCSMultipleGuests ErrSecCSUnimplemented = -67064
	// ErrSecCSNoMainExecutable: The code has no main executable file.
	ErrSecCSNoMainExecutable ErrSecCSUnimplemented = -67029
	// ErrSecCSNoMatches: No matches were found for a search or update operation.
	ErrSecCSNoMatches ErrSecCSUnimplemented = -67027
	// ErrSecCSNoSuchCode: Code signing host has no guest code with the requested attributes.
	ErrSecCSNoSuchCode ErrSecCSUnimplemented = -67065
	// ErrSecCSNotAHost: This code is not a code signing host.
	ErrSecCSNotAHost ErrSecCSUnimplemented = -67046
	// ErrSecCSNotAppLike: The code is valid but does not seem to be an app.
	ErrSecCSNotAppLike ErrSecCSUnimplemented = -67002
	// ErrSecCSNotSupported: Operation not supported for this type of code.
	ErrSecCSNotSupported ErrSecCSUnimplemented = -67037
	// ErrSecCSObjectRequired: A required pointer argument was null.
	ErrSecCSObjectRequired ErrSecCSUnimplemented = -67069
	// ErrSecCSOutdated: The presented data is out of date.
	ErrSecCSOutdated ErrSecCSUnimplemented = -67025
	// ErrSecCSRegularFile: The main executable or Info.plist must be a regular file (and not, for example, a symbolic link).
	ErrSecCSRegularFile        ErrSecCSUnimplemented = -67015
	ErrSecCSRemoteSignerFailed ErrSecCSUnimplemented = -66990
	// ErrSecCSReqFailed: The code failed to satisfy one of the code requirements.
	ErrSecCSReqFailed ErrSecCSUnimplemented = -67050
	// ErrSecCSReqInvalid: Invalid or corrupted code requirements.
	ErrSecCSReqInvalid ErrSecCSUnimplemented = -67052
	// ErrSecCSReqUnsupported: Unsupported type or version of code requirements.
	ErrSecCSReqUnsupported ErrSecCSUnimplemented = -67051
	// ErrSecCSResourceDirectoryFailed: A directory or its signature has been modified and is therefore invalid.
	ErrSecCSResourceDirectoryFailed ErrSecCSUnimplemented = -67023
	// ErrSecCSResourceNotSupported: Found an unsupported resource.
	ErrSecCSResourceNotSupported ErrSecCSUnimplemented = -67016
	// ErrSecCSResourceRulesInvalid: Invalid resource selection rule or rules.
	ErrSecCSResourceRulesInvalid ErrSecCSUnimplemented = -67053
	// ErrSecCSResourcesInvalid: The sealed resource directory is invalid.
	ErrSecCSResourcesInvalid ErrSecCSUnimplemented = -67055
	// ErrSecCSResourcesNotFound: Cannot find sealed resources in code.
	ErrSecCSResourcesNotFound ErrSecCSUnimplemented = -67056
	// ErrSecCSResourcesNotSealed: Resources are not sealed by the signature.
	ErrSecCSResourcesNotSealed ErrSecCSUnimplemented = -67057
	// ErrSecCSRevokedNotarization: Notarization indicates this code has been revoked.
	ErrSecCSRevokedNotarization ErrSecCSUnimplemented = -66992
	// ErrSecCSSigDBAccess: Can’t access signature database.
	ErrSecCSSigDBAccess ErrSecCSUnimplemented = -67032
	// ErrSecCSSigDBDenied: Access to signature database denied.
	ErrSecCSSigDBDenied ErrSecCSUnimplemented = -67033
	// ErrSecCSSignatureFailed: Code or signature modified.
	ErrSecCSSignatureFailed ErrSecCSUnimplemented = -67061
	// ErrSecCSSignatureInvalid: Invalid format for signature.
	ErrSecCSSignatureInvalid ErrSecCSUnimplemented = -67045
	// ErrSecCSSignatureNotVerifiable: Signature cannot be read.
	ErrSecCSSignatureNotVerifiable ErrSecCSUnimplemented = -67060
	// ErrSecCSSignatureUnsupported: Unsupported type or version of signature.
	ErrSecCSSignatureUnsupported ErrSecCSUnimplemented = -67059
	// ErrSecCSSignatureUntrusted: The signature is valid but signer isn’t trusted.
	ErrSecCSSignatureUntrusted ErrSecCSUnimplemented = -66996
	// ErrSecCSStaticCodeChanged: The code on disk has been modified after the code started running.
	ErrSecCSStaticCodeChanged ErrSecCSUnimplemented = -67034
	// ErrSecCSStaticCodeNotFound: Cannot find code object on disk.
	ErrSecCSStaticCodeNotFound ErrSecCSUnimplemented = -67068
	// ErrSecCSTooBig: The code is too big for current signing format.
	ErrSecCSTooBig ErrSecCSUnimplemented = -67004
	// ErrSecCSUnimplementedValue: Unimplemented code signing feature.
	ErrSecCSUnimplementedValue ErrSecCSUnimplemented = -67072
	// ErrSecCSUnsealedAppRoot: Unsealed contents present in the bundle root.
	ErrSecCSUnsealedAppRoot ErrSecCSUnimplemented = -67014
	// ErrSecCSUnsealedFrameworkRoot: Unsealed contents present in the root directory of an embedded framework.
	ErrSecCSUnsealedFrameworkRoot ErrSecCSUnimplemented = -67008
	// ErrSecCSUnsigned: Code object is not signed.
	ErrSecCSUnsigned ErrSecCSUnimplemented = -67062
	// ErrSecCSUnsignedNestedCode: Nested code is unsigned.
	ErrSecCSUnsignedNestedCode ErrSecCSUnimplemented = -67022
	// ErrSecCSUnsupportedDigestAlgorithm: The signature digest algorithm(s) specified are not supported.
	ErrSecCSUnsupportedDigestAlgorithm ErrSecCSUnimplemented = -67000
	// ErrSecCSUnsupportedGuestAttributes: Cannot locate guest code using this attribute set.
	ErrSecCSUnsupportedGuestAttributes ErrSecCSUnimplemented = -67067
	ErrSecCSVetoed                     ErrSecCSUnimplemented = -67018
	// ErrSecCSWeakResourceEnvelope: The resource envelope is obsolete (version 1 signature).
	ErrSecCSWeakResourceEnvelope ErrSecCSUnimplemented = -67007
	// ErrSecCSWeakResourceRules: The resource envelope is obsolete (custom omit rules).
	ErrSecCSWeakResourceRules ErrSecCSUnimplemented = -67013
	// ErrSecMultipleExecSegments: The image contains multiple executable segments.
	ErrSecMultipleExecSegments ErrSecCSUnimplemented = -66995
)

func (e ErrSecCSUnimplemented) String() string {
	switch e {
	case ErrSecCSAmbiguousBundleFormat:
		return "ErrSecCSAmbiguousBundleFormat"
	case ErrSecCSBadBundleFormat:
		return "ErrSecCSBadBundleFormat"
	case ErrSecCSBadCallbackValue:
		return "ErrSecCSBadCallbackValue"
	case ErrSecCSBadDictionaryFormat:
		return "ErrSecCSBadDictionaryFormat"
	case ErrSecCSBadDiskImageFormat:
		return "ErrSecCSBadDiskImageFormat"
	case ErrSecCSBadFrameworkVersion:
		return "ErrSecCSBadFrameworkVersion"
	case ErrSecCSBadLVArch:
		return "ErrSecCSBadLVArch"
	case ErrSecCSBadMainExecutable:
		return "ErrSecCSBadMainExecutable"
	case ErrSecCSBadNestedCode:
		return "ErrSecCSBadNestedCode"
	case ErrSecCSBadObjectFormat:
		return "ErrSecCSBadObjectFormat"
	case ErrSecCSBadResource:
		return "ErrSecCSBadResource"
	case ErrSecCSBadTeamIdentifier:
		return "ErrSecCSBadTeamIdentifier"
	case ErrSecCSCMSConstructionFailed:
		return "ErrSecCSCMSConstructionFailed"
	case ErrSecCSCMSTooLarge:
		return "ErrSecCSCMSTooLarge"
	case ErrSecCSCancelled:
		return "ErrSecCSCancelled"
	case ErrSecCSDBAccess:
		return "ErrSecCSDBAccess"
	case ErrSecCSDBDenied:
		return "ErrSecCSDBDenied"
	case ErrSecCSDSStoreSymlink:
		return "ErrSecCSDSStoreSymlink"
	case ErrSecCSDbCorrupt:
		return "ErrSecCSDbCorrupt"
	case ErrSecCSFileHardQuarantined:
		return "ErrSecCSFileHardQuarantined"
	case ErrSecCSGuestInvalid:
		return "ErrSecCSGuestInvalid"
	case ErrSecCSHelperFailed:
		return "ErrSecCSHelperFailed"
	case ErrSecCSHostProtocolContradiction:
		return "ErrSecCSHostProtocolContradiction"
	case ErrSecCSHostProtocolDedicationError:
		return "ErrSecCSHostProtocolDedicationError"
	case ErrSecCSHostProtocolInvalidAttribute:
		return "ErrSecCSHostProtocolInvalidAttribute"
	case ErrSecCSHostProtocolInvalidHash:
		return "ErrSecCSHostProtocolInvalidHash"
	case ErrSecCSHostProtocolNotProxy:
		return "ErrSecCSHostProtocolNotProxy"
	case ErrSecCSHostProtocolRelativePath:
		return "ErrSecCSHostProtocolRelativePath"
	case ErrSecCSHostProtocolStateError:
		return "ErrSecCSHostProtocolStateError"
	case ErrSecCSHostProtocolUnrelated:
		return "ErrSecCSHostProtocolUnrelated"
	case ErrSecCSHostReject:
		return "ErrSecCSHostReject"
	case ErrSecCSInfoPlistFailed:
		return "ErrSecCSInfoPlistFailed"
	case ErrSecCSInternalError:
		return "ErrSecCSInternalError"
	case ErrSecCSInvalidAssociatedFileData:
		return "ErrSecCSInvalidAssociatedFileData"
	case ErrSecCSInvalidAttributeValues:
		return "ErrSecCSInvalidAttributeValues"
	case ErrSecCSInvalidEntitlements:
		return "ErrSecCSInvalidEntitlements"
	case ErrSecCSInvalidFlags:
		return "ErrSecCSInvalidFlags"
	case ErrSecCSInvalidObjectRef:
		return "ErrSecCSInvalidObjectRef"
	case ErrSecCSInvalidPlatform:
		return "ErrSecCSInvalidPlatform"
	case ErrSecCSInvalidRuntimeVersion:
		return "ErrSecCSInvalidRuntimeVersion"
	case ErrSecCSInvalidSymlink:
		return "ErrSecCSInvalidSymlink"
	case ErrSecCSInvalidTeamIdentifier:
		return "ErrSecCSInvalidTeamIdentifier"
	case ErrSecCSMultipleGuests:
		return "ErrSecCSMultipleGuests"
	case ErrSecCSNoMainExecutable:
		return "ErrSecCSNoMainExecutable"
	case ErrSecCSNoMatches:
		return "ErrSecCSNoMatches"
	case ErrSecCSNoSuchCode:
		return "ErrSecCSNoSuchCode"
	case ErrSecCSNotAHost:
		return "ErrSecCSNotAHost"
	case ErrSecCSNotAppLike:
		return "ErrSecCSNotAppLike"
	case ErrSecCSNotSupported:
		return "ErrSecCSNotSupported"
	case ErrSecCSObjectRequired:
		return "ErrSecCSObjectRequired"
	case ErrSecCSOutdated:
		return "ErrSecCSOutdated"
	case ErrSecCSRegularFile:
		return "ErrSecCSRegularFile"
	case ErrSecCSRemoteSignerFailed:
		return "ErrSecCSRemoteSignerFailed"
	case ErrSecCSReqFailed:
		return "ErrSecCSReqFailed"
	case ErrSecCSReqInvalid:
		return "ErrSecCSReqInvalid"
	case ErrSecCSReqUnsupported:
		return "ErrSecCSReqUnsupported"
	case ErrSecCSResourceDirectoryFailed:
		return "ErrSecCSResourceDirectoryFailed"
	case ErrSecCSResourceNotSupported:
		return "ErrSecCSResourceNotSupported"
	case ErrSecCSResourceRulesInvalid:
		return "ErrSecCSResourceRulesInvalid"
	case ErrSecCSResourcesInvalid:
		return "ErrSecCSResourcesInvalid"
	case ErrSecCSResourcesNotFound:
		return "ErrSecCSResourcesNotFound"
	case ErrSecCSResourcesNotSealed:
		return "ErrSecCSResourcesNotSealed"
	case ErrSecCSRevokedNotarization:
		return "ErrSecCSRevokedNotarization"
	case ErrSecCSSignatureFailed:
		return "ErrSecCSSignatureFailed"
	case ErrSecCSSignatureInvalid:
		return "ErrSecCSSignatureInvalid"
	case ErrSecCSSignatureNotVerifiable:
		return "ErrSecCSSignatureNotVerifiable"
	case ErrSecCSSignatureUnsupported:
		return "ErrSecCSSignatureUnsupported"
	case ErrSecCSSignatureUntrusted:
		return "ErrSecCSSignatureUntrusted"
	case ErrSecCSStaticCodeChanged:
		return "ErrSecCSStaticCodeChanged"
	case ErrSecCSStaticCodeNotFound:
		return "ErrSecCSStaticCodeNotFound"
	case ErrSecCSTooBig:
		return "ErrSecCSTooBig"
	case ErrSecCSUnimplementedValue:
		return "ErrSecCSUnimplementedValue"
	case ErrSecCSUnsealedAppRoot:
		return "ErrSecCSUnsealedAppRoot"
	case ErrSecCSUnsealedFrameworkRoot:
		return "ErrSecCSUnsealedFrameworkRoot"
	case ErrSecCSUnsigned:
		return "ErrSecCSUnsigned"
	case ErrSecCSUnsignedNestedCode:
		return "ErrSecCSUnsignedNestedCode"
	case ErrSecCSUnsupportedDigestAlgorithm:
		return "ErrSecCSUnsupportedDigestAlgorithm"
	case ErrSecCSUnsupportedGuestAttributes:
		return "ErrSecCSUnsupportedGuestAttributes"
	case ErrSecCSVetoed:
		return "ErrSecCSVetoed"
	case ErrSecCSWeakResourceEnvelope:
		return "ErrSecCSWeakResourceEnvelope"
	case ErrSecCSWeakResourceRules:
		return "ErrSecCSWeakResourceRules"
	case ErrSecMultipleExecSegments:
		return "ErrSecMultipleExecSegments"
	default:
		return fmt.Sprintf("ErrSecCSUnimplemented(%d)", e)
	}
}

type ErrSecSuccess int32

const (
	// ErrSecACLAddFailed: An ACL add operation failed.
	ErrSecACLAddFailed ErrSecSuccess = -67698
	// ErrSecACLChangeFailed: An ACL change operation failed.
	ErrSecACLChangeFailed ErrSecSuccess = -67699
	// ErrSecACLDeleteFailed: An ACL delete operation failed.
	ErrSecACLDeleteFailed ErrSecSuccess = -67696
	// ErrSecACLNotSimple: The access control list is not in standard simple form.
	ErrSecACLNotSimple ErrSecSuccess = -25240
	// ErrSecACLReplaceFailed: An ACL replace operation failed.
	ErrSecACLReplaceFailed ErrSecSuccess = -67697
	// ErrSecAddinLoadFailed: The add-in load operation failed.
	ErrSecAddinLoadFailed ErrSecSuccess = -67711
	// ErrSecAddinUnloadFailed: The add-in unload operation failed.
	ErrSecAddinUnloadFailed ErrSecSuccess = -67714
	// ErrSecAlgorithmMismatch: An algorithm mismatch occurred.
	ErrSecAlgorithmMismatch ErrSecSuccess = -67730
	// ErrSecAllocate: Failed to allocate memory.
	ErrSecAllocate ErrSecSuccess = -108
	// ErrSecAlreadyLoggedIn: The user is already logged in.
	ErrSecAlreadyLoggedIn ErrSecSuccess = -67814
	// ErrSecAppleAddAppACLSubject: Adding an application ACL subject failed.
	ErrSecAppleAddAppACLSubject ErrSecSuccess = -67589
	// ErrSecAppleInvalidKeyEndDate: The specified key has an invalid end date.
	ErrSecAppleInvalidKeyEndDate ErrSecSuccess = -67593
	// ErrSecAppleInvalidKeyStartDate: The specified key has an invalid start date.
	ErrSecAppleInvalidKeyStartDate ErrSecSuccess = -67592
	// ErrSecApplePublicKeyIncomplete: The public key is incomplete.
	ErrSecApplePublicKeyIncomplete ErrSecSuccess = -67590
	// ErrSecAppleSSLv2Rollback: A SSLv2 rollback error has occurred.
	ErrSecAppleSSLv2Rollback ErrSecSuccess = -67595
	// ErrSecAppleSignatureMismatch: A signature mismatch has occurred.
	ErrSecAppleSignatureMismatch ErrSecSuccess = -67591
	// ErrSecAttachHandleBusy: The CSP handle was busy.
	ErrSecAttachHandleBusy ErrSecSuccess = -67728
	// ErrSecAttributeNotInContext: An attribute was not in the context.
	ErrSecAttributeNotInContext ErrSecSuccess = -67720
	// ErrSecAuthFailed: Authorization and/or authentication failed.
	ErrSecAuthFailed ErrSecSuccess = -25293
	// ErrSecBadReq: Bad parameter or invalid state for operation.
	ErrSecBadReq ErrSecSuccess = -909
	// ErrSecBlockSizeMismatch: A block size mismatch occurred.
	ErrSecBlockSizeMismatch ErrSecSuccess = -67810
	// ErrSecBufferTooSmall: The buffer is too small.
	ErrSecBufferTooSmall ErrSecSuccess = -25301
	// ErrSecCRLAlreadySigned: The certificate revocation list is already signed.
	ErrSecCRLAlreadySigned ErrSecSuccess = -67684
	// ErrSecCRLBadURI: The certificate revocation list has a bad uniform resource identifier.
	ErrSecCRLBadURI ErrSecSuccess = -67617
	// ErrSecCRLExpired: The certificate revocation list has expired.
	ErrSecCRLExpired ErrSecSuccess = -67613
	// ErrSecCRLNotFound: The certificate revocation list was not found.
	ErrSecCRLNotFound ErrSecSuccess = -67615
	// ErrSecCRLNotTrusted: The certificate revocation list is not trusted.
	ErrSecCRLNotTrusted ErrSecSuccess = -67620
	// ErrSecCRLNotValidYet: The certificate revocation list is not yet valid.
	ErrSecCRLNotValidYet ErrSecSuccess = -67614
	// ErrSecCRLPolicyFailed: The certificate revocation list policy failed.
	ErrSecCRLPolicyFailed ErrSecSuccess = -67621
	// ErrSecCRLServerDown: The certificate revocation list server is down.
	ErrSecCRLServerDown ErrSecSuccess = -67616
	// ErrSecCallbackFailed: A callback failed.
	ErrSecCallbackFailed ErrSecSuccess = -67695
	// ErrSecCertificateCannotOperate: The certificate cannot operate.
	ErrSecCertificateCannotOperate      ErrSecSuccess = -67817
	ErrSecCertificateDuplicateExtension ErrSecSuccess = -67903
	// ErrSecCertificateExpired: An expired certificate was detected.
	ErrSecCertificateExpired ErrSecSuccess = -67818
	ErrSecCertificateIsCA    ErrSecSuccess = -67902
	// ErrSecCertificateNameNotAllowed: The requested name isn’t allowed for this certificate.
	ErrSecCertificateNameNotAllowed ErrSecSuccess = -67900
	// ErrSecCertificateNotValidYet: The certificate is not yet valid.
	ErrSecCertificateNotValidYet ErrSecSuccess = -67819
	// ErrSecCertificatePolicyNotAllowed: The requested policy isn’t allowed for this certificate.
	ErrSecCertificatePolicyNotAllowed ErrSecSuccess = -67899
	// ErrSecCertificateRevoked: The certificate was revoked.
	ErrSecCertificateRevoked ErrSecSuccess = -67820
	// ErrSecCertificateSuspended: The certificate was suspended.
	ErrSecCertificateSuspended ErrSecSuccess = -67821
	// ErrSecCertificateValidityPeriodTooLong: The validity period in the certificate exceeds the maximum allowed period.
	ErrSecCertificateValidityPeriodTooLong ErrSecSuccess = -67901
	// ErrSecCodeSigningBadCertChainLength: Code signing encountered an incorrect certificate chain length.
	ErrSecCodeSigningBadCertChainLength ErrSecSuccess = -67647
	// ErrSecCodeSigningBadPathLengthConstraint: Code signing encountered an incorrect path length constraint.
	ErrSecCodeSigningBadPathLengthConstraint ErrSecSuccess = -67649
	// ErrSecCodeSigningDevelopment: Code signing indicated use of a development-only certificate.
	ErrSecCodeSigningDevelopment ErrSecSuccess = -67651
	// ErrSecCodeSigningNoBasicConstraints: Code signing found no basic constraints.
	ErrSecCodeSigningNoBasicConstraints ErrSecSuccess = -67648
	// ErrSecCodeSigningNoExtendedKeyUsage: Code signing found no extended key usage.
	ErrSecCodeSigningNoExtendedKeyUsage ErrSecSuccess = -67650
	// ErrSecConversionError: A conversion error has occurred.
	ErrSecConversionError ErrSecSuccess = -67594
	// ErrSecCoreFoundationUnknown: An unknown Core Foundation error occurred.
	ErrSecCoreFoundationUnknown ErrSecSuccess = -4960
	// ErrSecCreateChainFailed: The attempt to create a certificate chain failed.
	ErrSecCreateChainFailed ErrSecSuccess = -25318
	// ErrSecDataNotAvailable: The data is not available.
	ErrSecDataNotAvailable ErrSecSuccess = -25316
	// ErrSecDataNotModifiable: The data is not modifiable.
	ErrSecDataNotModifiable ErrSecSuccess = -25317
	// ErrSecDataTooLarge: The data is too large for the particular data type.
	ErrSecDataTooLarge ErrSecSuccess = -25302
	// ErrSecDatabaseLocked: The database is locked.
	ErrSecDatabaseLocked ErrSecSuccess = -67869
	// ErrSecDatastoreIsOpen: The data store is open.
	ErrSecDatastoreIsOpen ErrSecSuccess = -67870
	// ErrSecDecode: Unable to decode the provided data.
	ErrSecDecode ErrSecSuccess = -26275
	// ErrSecDeviceError: A device error was encountered.
	ErrSecDeviceError ErrSecSuccess = -67727
	// ErrSecDeviceFailed: A device failure has occurred.
	ErrSecDeviceFailed ErrSecSuccess = -67588
	// ErrSecDeviceReset: A device reset has occurred.
	ErrSecDeviceReset ErrSecSuccess = -67587
	// ErrSecDeviceVerifyFailed: A device verification failure has occurred.
	ErrSecDeviceVerifyFailed ErrSecSuccess = -67812
	// ErrSecDiskFull: The disk is full.
	ErrSecDiskFull ErrSecSuccess = -34
	// ErrSecDuplicateCallback: More than one callback of the same name exists.
	ErrSecDuplicateCallback ErrSecSuccess = -25297
	// ErrSecDuplicateItem: The item already exists.
	ErrSecDuplicateItem ErrSecSuccess = -25299
	// ErrSecDuplicateKeychain: A keychain with the same name already exists.
	ErrSecDuplicateKeychain ErrSecSuccess = -25296
	// ErrSecEMMLoadFailed: The elective module manager load failed.
	ErrSecEMMLoadFailed ErrSecSuccess = -67709
	// ErrSecEMMUnloadFailed: The elective module manager unload has failed.
	ErrSecEMMUnloadFailed ErrSecSuccess = -67710
	// ErrSecEndOfData: An end-of-data was detected.
	ErrSecEndOfData ErrSecSuccess = -67634
	// ErrSecEventNotificationCallbackNotFound: An event notification callback was not found.
	ErrSecEventNotificationCallbackNotFound ErrSecSuccess = -67723
	// ErrSecExtendedKeyUsageNotCritical: The extended key usage extension was not marked critical.
	ErrSecExtendedKeyUsageNotCritical ErrSecSuccess = -67881
	// ErrSecFieldSpecifiedMultiple: Too many fields were specified.
	ErrSecFieldSpecifiedMultiple ErrSecSuccess = -67866
	// ErrSecFileTooBig: The file is too big.
	ErrSecFileTooBig ErrSecSuccess = -67597
	// ErrSecFunctionFailed: A function has failed.
	ErrSecFunctionFailed ErrSecSuccess = -67677
	// ErrSecFunctionIntegrityFail: A function address is not within the verified module.
	ErrSecFunctionIntegrityFail ErrSecSuccess = -67670
	// ErrSecHostNameMismatch: A host name mismatch has occurred.
	ErrSecHostNameMismatch ErrSecSuccess = -67602
	// ErrSecIDPFailure: The issuing distribution point is not valid.
	ErrSecIDPFailure ErrSecSuccess = -67622
	// ErrSecIO: I/O error.
	ErrSecIO ErrSecSuccess = -36
	// ErrSecInDarkWake: The user interface cannot be displayed because the system is in a dark wake state.
	ErrSecInDarkWake ErrSecSuccess = -25320
	// ErrSecIncompatibleDatabaseBlob: The specified database has an incompatible blob.
	ErrSecIncompatibleDatabaseBlob ErrSecSuccess = -67600
	// ErrSecIncompatibleFieldFormat: The field format is incompatible.
	ErrSecIncompatibleFieldFormat ErrSecSuccess = -67867
	// ErrSecIncompatibleKeyBlob: The specified database has an incompatible key blob.
	ErrSecIncompatibleKeyBlob ErrSecSuccess = -67601
	// ErrSecIncompatibleVersion: The version is incompatible.
	ErrSecIncompatibleVersion ErrSecSuccess = -67704
	// ErrSecIncompleteCertRevocationCheck: An incomplete certificate revocation check occurred.
	ErrSecIncompleteCertRevocationCheck ErrSecSuccess = -67635
	// ErrSecInputLengthError: An input length error occurred.
	ErrSecInputLengthError ErrSecSuccess = -67724
	// ErrSecInsufficientClientID: The client ID is incorrect.
	ErrSecInsufficientClientID ErrSecSuccess = -67586
	// ErrSecInsufficientCredentials: Insufficient credentials were detected.
	ErrSecInsufficientCredentials ErrSecSuccess = -67822
	// ErrSecInteractionNotAllowed: Interaction with the Security Server is not allowed.
	ErrSecInteractionNotAllowed ErrSecSuccess = -25308
	// ErrSecInteractionRequired: User interaction is required.
	ErrSecInteractionRequired ErrSecSuccess = -25315
	// ErrSecInternalComponent: An internal component experienced an error.
	ErrSecInternalComponent ErrSecSuccess = -2070
	// ErrSecInternalError: An internal error occurred.
	ErrSecInternalError ErrSecSuccess = -67671
	// ErrSecInvalidACL: An invalid access control list was detected.
	ErrSecInvalidACL ErrSecSuccess = -67702
	// ErrSecInvalidAccessCredentials: Invalid access credentials were detected.
	ErrSecInvalidAccessCredentials ErrSecSuccess = -67700
	// ErrSecInvalidAccessRequest: The access request is invalid.
	ErrSecInvalidAccessRequest ErrSecSuccess = -67876
	// ErrSecInvalidAction: The action is invalid.
	ErrSecInvalidAction ErrSecSuccess = -67823
	// ErrSecInvalidAddinFunctionTable: An invalid add-in function table was detected.
	ErrSecInvalidAddinFunctionTable ErrSecSuccess = -67716
	// ErrSecInvalidAlgorithm: An invalid algorithm was detected.
	ErrSecInvalidAlgorithm ErrSecSuccess = -67747
	// ErrSecInvalidAlgorithmParms: An algorithm parameters attribute is not valid.
	ErrSecInvalidAlgorithmParms ErrSecSuccess = -67770
	// ErrSecInvalidAttributeAccessCredentials: An access credentials attribute is not valid.
	ErrSecInvalidAttributeAccessCredentials ErrSecSuccess = -67796
	// ErrSecInvalidAttributeBase: A base attribute is not valid.
	ErrSecInvalidAttributeBase ErrSecSuccess = -67788
	// ErrSecInvalidAttributeBlockSize: A block size attribute is not valid.
	ErrSecInvalidAttributeBlockSize ErrSecSuccess = -67764
	// ErrSecInvalidAttributeDLDBHandle: A database handle attribute is not valid.
	ErrSecInvalidAttributeDLDBHandle ErrSecSuccess = -67794
	// ErrSecInvalidAttributeEffectiveBits: An effective bits attribute is not valid.
	ErrSecInvalidAttributeEffectiveBits ErrSecSuccess = -67778
	// ErrSecInvalidAttributeEndDate: An end date attribute is not valid.
	ErrSecInvalidAttributeEndDate ErrSecSuccess = -67782
	// ErrSecInvalidAttributeInitVector: An init vector attribute is not valid.
	ErrSecInvalidAttributeInitVector ErrSecSuccess = -67750
	// ErrSecInvalidAttributeIterationCount: An iteration count attribute is not valid.
	ErrSecInvalidAttributeIterationCount ErrSecSuccess = -67792
	// ErrSecInvalidAttributeKey: A key attribute is not valid.
	ErrSecInvalidAttributeKey ErrSecSuccess = -67748
	// ErrSecInvalidAttributeKeyLength: A key length attribute is not valid.
	ErrSecInvalidAttributeKeyLength ErrSecSuccess = -67762
	// ErrSecInvalidAttributeKeyType: A key type attribute is not valid.
	ErrSecInvalidAttributeKeyType ErrSecSuccess = -67774
	// ErrSecInvalidAttributeLabel: A label attribute is not valid.
	ErrSecInvalidAttributeLabel ErrSecSuccess = -67772
	// ErrSecInvalidAttributeMode: A mode attribute is not valid.
	ErrSecInvalidAttributeMode ErrSecSuccess = -67776
	// ErrSecInvalidAttributeOutputSize: An output size attribute is not valid.
	ErrSecInvalidAttributeOutputSize ErrSecSuccess = -67766
	// ErrSecInvalidAttributePadding: A padding attribute is not valid.
	ErrSecInvalidAttributePadding ErrSecSuccess = -67754
	// ErrSecInvalidAttributePassphrase: A passphrase attribute is not valid.
	ErrSecInvalidAttributePassphrase ErrSecSuccess = -67760
	// ErrSecInvalidAttributePrime: A prime attribute is not valid.
	ErrSecInvalidAttributePrime ErrSecSuccess = -67786
	// ErrSecInvalidAttributePrivateKeyFormat: A private key format attribute is not valid.
	ErrSecInvalidAttributePrivateKeyFormat ErrSecSuccess = -67800
	// ErrSecInvalidAttributePublicKeyFormat: A public key format attribute is not valid.
	ErrSecInvalidAttributePublicKeyFormat ErrSecSuccess = -67798
	// ErrSecInvalidAttributeRandom: A random number attribute is not valid.
	ErrSecInvalidAttributeRandom ErrSecSuccess = -67756
	// ErrSecInvalidAttributeRounds: The number of rounds attribute is not valid.
	ErrSecInvalidAttributeRounds ErrSecSuccess = -67768
	// ErrSecInvalidAttributeSalt: A salt attribute is not valid.
	ErrSecInvalidAttributeSalt ErrSecSuccess = -67752
	// ErrSecInvalidAttributeSeed: A seed attribute is not valid.
	ErrSecInvalidAttributeSeed ErrSecSuccess = -67758
	// ErrSecInvalidAttributeStartDate: A start date attribute is not valid.
	ErrSecInvalidAttributeStartDate ErrSecSuccess = -67780
	// ErrSecInvalidAttributeSubprime: A subprime attribute is not valid.
	ErrSecInvalidAttributeSubprime ErrSecSuccess = -67790
	// ErrSecInvalidAttributeSymmetricKeyFormat: A symmetric key format attribute is not valid.
	ErrSecInvalidAttributeSymmetricKeyFormat ErrSecSuccess = -67802
	// ErrSecInvalidAttributeVersion: A version attribute is not valid.
	ErrSecInvalidAttributeVersion ErrSecSuccess = -67784
	// ErrSecInvalidAttributeWrappedKeyFormat: A wrapped key format attribute is not valid.
	ErrSecInvalidAttributeWrappedKeyFormat ErrSecSuccess = -67804
	// ErrSecInvalidAuthority: The authority is not valid.
	ErrSecInvalidAuthority ErrSecSuccess = -67824
	// ErrSecInvalidAuthorityKeyID: The authority key ID is not valid.
	ErrSecInvalidAuthorityKeyID ErrSecSuccess = -67606
	// ErrSecInvalidBaseACLs: The base access control lists are not valid.
	ErrSecInvalidBaseACLs ErrSecSuccess = -67851
	// ErrSecInvalidBundleInfo: The bundle information is not valid.
	ErrSecInvalidBundleInfo ErrSecSuccess = -67857
	// ErrSecInvalidCRL: The certificate revocation list is not valid.
	ErrSecInvalidCRL          ErrSecSuccess = -67830
	ErrSecInvalidCRLAuthority ErrSecSuccess = -67827
	// ErrSecInvalidCRLEncoding: The certificate revocation list encoding is not valid.
	ErrSecInvalidCRLEncoding ErrSecSuccess = -67828
	// ErrSecInvalidCRLGroup: An invalid certificate revocation list group was detected.
	ErrSecInvalidCRLGroup ErrSecSuccess = -67816
	// ErrSecInvalidCRLIndex: The certificate revocation list index is not valid.
	ErrSecInvalidCRLIndex ErrSecSuccess = -67858
	// ErrSecInvalidCRLType: The certificate revocation list type is not valid.
	ErrSecInvalidCRLType ErrSecSuccess = -67829
	// ErrSecInvalidCallback: The callback is not valid.
	ErrSecInvalidCallback ErrSecSuccess = -25298
	// ErrSecInvalidCertAuthority: The certificate authority is not valid.
	ErrSecInvalidCertAuthority ErrSecSuccess = -67826
	// ErrSecInvalidCertificateGroup: An invalid certificate group was detected.
	ErrSecInvalidCertificateGroup ErrSecSuccess = -67691
	// ErrSecInvalidCertificateRef: An invalid certificate reference was detected.
	ErrSecInvalidCertificateRef ErrSecSuccess = -67690
	// ErrSecInvalidContext: An invalid context was detected.
	ErrSecInvalidContext ErrSecSuccess = -67746
	// ErrSecInvalidDBList: An invalid DB list was detected.
	ErrSecInvalidDBList ErrSecSuccess = -67681
	// ErrSecInvalidDBLocation: The database location is not valid.
	ErrSecInvalidDBLocation ErrSecSuccess = -67875
	// ErrSecInvalidData: Invalid data was detected.
	ErrSecInvalidData ErrSecSuccess = -67673
	// ErrSecInvalidDatabaseBlob: The specified database has an invalid blob.
	ErrSecInvalidDatabaseBlob ErrSecSuccess = -67598
	// ErrSecInvalidDigestAlgorithm: An invalid digest algorithm was detected.
	ErrSecInvalidDigestAlgorithm ErrSecSuccess = -67815
	// ErrSecInvalidEncoding: The encoding is not valid.
	ErrSecInvalidEncoding ErrSecSuccess = -67853
	// ErrSecInvalidExtendedKeyUsage: The extended key usage is not valid.
	ErrSecInvalidExtendedKeyUsage ErrSecSuccess = -67609
	// ErrSecInvalidFormType: The form type is not valid.
	ErrSecInvalidFormType ErrSecSuccess = -67831
	// ErrSecInvalidGUID: An invalid GUID was detected.
	ErrSecInvalidGUID ErrSecSuccess = -67679
	// ErrSecInvalidHandle: An invalid handle was encountered.
	ErrSecInvalidHandle ErrSecSuccess = -67680
	// ErrSecInvalidHandleUsage: The common security services manager handle does not match with the service type.
	ErrSecInvalidHandleUsage ErrSecSuccess = -67668
	// ErrSecInvalidID: The ID is not valid.
	ErrSecInvalidID ErrSecSuccess = -67832
	// ErrSecInvalidIDLinkage: The ID linkage is not valid.
	ErrSecInvalidIDLinkage ErrSecSuccess = -67610
	// ErrSecInvalidIdentifier: The identifier is not valid.
	ErrSecInvalidIdentifier ErrSecSuccess = -67833
	// ErrSecInvalidIndex: The index is not valid.
	ErrSecInvalidIndex ErrSecSuccess = -67834
	// ErrSecInvalidIndexInfo: The index information is not valid.
	ErrSecInvalidIndexInfo ErrSecSuccess = -67877
	// ErrSecInvalidInputVector: The input vector is not valid.
	ErrSecInvalidInputVector ErrSecSuccess = -67744
	// ErrSecInvalidItemRef: The item reference is invalid.
	ErrSecInvalidItemRef ErrSecSuccess = -25304
	// ErrSecInvalidKeyAttributeMask: The key attribute mask is not valid.
	ErrSecInvalidKeyAttributeMask ErrSecSuccess = -67738
	// ErrSecInvalidKeyBlob: The specified database has an invalid key blob.
	ErrSecInvalidKeyBlob ErrSecSuccess = -67599
	// ErrSecInvalidKeyFormat: The key format is not valid.
	ErrSecInvalidKeyFormat ErrSecSuccess = -67742
	// ErrSecInvalidKeyHierarchy: An invalid key hierarchy was detected.
	ErrSecInvalidKeyHierarchy ErrSecSuccess = -67713
	// ErrSecInvalidKeyLabel: The key label is not valid.
	ErrSecInvalidKeyLabel ErrSecSuccess = -67740
	// ErrSecInvalidKeyRef: An invalid key was encountered.
	ErrSecInvalidKeyRef ErrSecSuccess = -67712
	// ErrSecInvalidKeyUsageForPolicy: The key usage is not valid for the specified policy.
	ErrSecInvalidKeyUsageForPolicy ErrSecSuccess = -67608
	// ErrSecInvalidKeyUsageMask: The key usage mask is not valid.
	ErrSecInvalidKeyUsageMask ErrSecSuccess = -67736
	// ErrSecInvalidKeychain: The keychain is not valid.
	ErrSecInvalidKeychain ErrSecSuccess = -25295
	// ErrSecInvalidLoginName: An invalid login name was detected.
	ErrSecInvalidLoginName ErrSecSuccess = -67813
	// ErrSecInvalidModifyMode: The modify mode is not valid.
	ErrSecInvalidModifyMode ErrSecSuccess = -67879
	// ErrSecInvalidName: An invalid name was detected.
	ErrSecInvalidName ErrSecSuccess = -67689
	// ErrSecInvalidNetworkAddress: An invalid network address was detected.
	ErrSecInvalidNetworkAddress ErrSecSuccess = -67683
	// ErrSecInvalidNewOwner: The new owner is not valid.
	ErrSecInvalidNewOwner ErrSecSuccess = -67878
	// ErrSecInvalidNumberOfFields: An invalid number of fields were detected.
	ErrSecInvalidNumberOfFields ErrSecSuccess = -67685
	// ErrSecInvalidOutputVector: The output vector is not valid.
	ErrSecInvalidOutputVector ErrSecSuccess = -67745
	// ErrSecInvalidOwnerEdit: An invalid attempt to change the owner of an item.
	ErrSecInvalidOwnerEdit ErrSecSuccess = -25244
	// ErrSecInvalidPVC: An invalid pointer validation checking policy was detected.
	ErrSecInvalidPVC ErrSecSuccess = -67708
	// ErrSecInvalidParsingModule: The parsing module is not valid.
	ErrSecInvalidParsingModule ErrSecSuccess = -67868
	// ErrSecInvalidPassthroughID: An invalid passthrough ID was detected.
	ErrSecInvalidPassthroughID ErrSecSuccess = -67682
	// ErrSecInvalidPasswordRef: The password reference is invalid.
	ErrSecInvalidPasswordRef ErrSecSuccess = -25261
	// ErrSecInvalidPointer: An invalid pointer was detected.
	ErrSecInvalidPointer ErrSecSuccess = -67675
	// ErrSecInvalidPolicyIdentifiers: The policy identifiers are not valid.
	ErrSecInvalidPolicyIdentifiers ErrSecSuccess = -67835
	// ErrSecInvalidPrefsDomain: The preference domain specified is invalid.
	ErrSecInvalidPrefsDomain ErrSecSuccess = -25319
	// ErrSecInvalidQuery: The specified query is not valid.
	ErrSecInvalidQuery ErrSecSuccess = -67693
	// ErrSecInvalidReason: The trust policy reason is not valid.
	ErrSecInvalidReason ErrSecSuccess = -67837
	// ErrSecInvalidRecord: An invalid record was detected.
	ErrSecInvalidRecord ErrSecSuccess = -67701
	// ErrSecInvalidRequestInputs: The request inputs are not valid.
	ErrSecInvalidRequestInputs ErrSecSuccess = -67838
	// ErrSecInvalidRequestor: The requestor is not valid.
	ErrSecInvalidRequestor ErrSecSuccess = -67855
	// ErrSecInvalidResponseVector: The response vector is not valid.
	ErrSecInvalidResponseVector ErrSecSuccess = -67839
	// ErrSecInvalidRoot: The root or anchor certificate is not valid.
	ErrSecInvalidRoot ErrSecSuccess = -67612
	// ErrSecInvalidSampleValue: An invalid sample value was detected.
	ErrSecInvalidSampleValue ErrSecSuccess = -67703
	// ErrSecInvalidScope: An invalid scope was detected.
	ErrSecInvalidScope ErrSecSuccess = -67706
	// ErrSecInvalidSearchRef: The search reference is invalid.
	ErrSecInvalidSearchRef ErrSecSuccess = -25305
	// ErrSecInvalidServiceMask: An invalid service mask was detected.
	ErrSecInvalidServiceMask ErrSecSuccess = -67717
	// ErrSecInvalidSignature: An invalid signature was detected.
	ErrSecInvalidSignature ErrSecSuccess = -67688
	// ErrSecInvalidStopOnPolicy: The stop-on policy is not valid.
	ErrSecInvalidStopOnPolicy ErrSecSuccess = -67840
	// ErrSecInvalidSubServiceID: An invalid sub-service ID was detected.
	ErrSecInvalidSubServiceID ErrSecSuccess = -67719
	// ErrSecInvalidSubjectKeyID: The subject key ID is not valid.
	ErrSecInvalidSubjectKeyID ErrSecSuccess = -67607
	// ErrSecInvalidSubjectName: An invalid certificate subject name was detected.
	ErrSecInvalidSubjectName ErrSecSuccess = -67655
	// ErrSecInvalidTimeString: The time specified is not valid.
	ErrSecInvalidTimeString ErrSecSuccess = -67836
	// ErrSecInvalidTrustSetting: The trust setting is invalid.
	ErrSecInvalidTrustSetting ErrSecSuccess = -25242
	// ErrSecInvalidTrustSettings: The trust settings record is corrupted.
	ErrSecInvalidTrustSettings ErrSecSuccess = -25262
	// ErrSecInvalidTuple: The tuple is not valid.
	ErrSecInvalidTuple ErrSecSuccess = -67841
	// ErrSecInvalidTupleCredentials: The tuple credentials are not valid.
	ErrSecInvalidTupleCredentials ErrSecSuccess = -67852
	// ErrSecInvalidTupleGroup: The tuple group is not valid.
	ErrSecInvalidTupleGroup ErrSecSuccess = -67850
	// ErrSecInvalidValidityPeriod: The validity period is not valid.
	ErrSecInvalidValidityPeriod ErrSecSuccess = -67854
	// ErrSecInvalidValue: An invalid value was detected.
	ErrSecInvalidValue ErrSecSuccess = -67694
	// ErrSecItemNotFound: The item cannot be found.
	ErrSecItemNotFound ErrSecSuccess = -25300
	// ErrSecKeyBlobTypeIncorrect: The key blob type is incorrect.
	ErrSecKeyBlobTypeIncorrect ErrSecSuccess = -67732
	// ErrSecKeyHeaderInconsistent: The key header is inconsistent.
	ErrSecKeyHeaderInconsistent ErrSecSuccess = -67733
	// ErrSecKeyIsSensitive: The key must be wrapped to be exported.
	ErrSecKeyIsSensitive ErrSecSuccess = -25258
	// ErrSecKeySizeNotAllowed: The key size is not allowed.
	ErrSecKeySizeNotAllowed ErrSecSuccess = -25311
	// ErrSecKeyUsageIncorrect: The key usage is incorrect.
	ErrSecKeyUsageIncorrect ErrSecSuccess = -67731
	// ErrSecLibraryReferenceNotFound: A library reference was not found.
	ErrSecLibraryReferenceNotFound ErrSecSuccess = -67715
	// ErrSecMDSError: A module directory service error occurred.
	ErrSecMDSError ErrSecSuccess = -67674
	// ErrSecMemoryError: A memory error occurred.
	ErrSecMemoryError ErrSecSuccess = -67672
	// ErrSecMissingAlgorithmParms: An algorithm parameters attribute is missing.
	ErrSecMissingAlgorithmParms ErrSecSuccess = -67771
	// ErrSecMissingAttributeAccessCredentials: An access credentials attribute is missing.
	ErrSecMissingAttributeAccessCredentials ErrSecSuccess = -67797
	// ErrSecMissingAttributeBase: A base attribute is missing.
	ErrSecMissingAttributeBase ErrSecSuccess = -67789
	// ErrSecMissingAttributeBlockSize: A block size attribute is missing.
	ErrSecMissingAttributeBlockSize ErrSecSuccess = -67765
	// ErrSecMissingAttributeDLDBHandle: A database handle attribute is missing.
	ErrSecMissingAttributeDLDBHandle ErrSecSuccess = -67795
	// ErrSecMissingAttributeEffectiveBits: An effective bits attribute is missing.
	ErrSecMissingAttributeEffectiveBits ErrSecSuccess = -67779
	// ErrSecMissingAttributeEndDate: An end date attribute is missing.
	ErrSecMissingAttributeEndDate ErrSecSuccess = -67783
	// ErrSecMissingAttributeInitVector: An init vector attribute is missing.
	ErrSecMissingAttributeInitVector ErrSecSuccess = -67751
	// ErrSecMissingAttributeIterationCount: An iteration count attribute is missing.
	ErrSecMissingAttributeIterationCount ErrSecSuccess = -67793
	// ErrSecMissingAttributeKey: A key attribute is missing.
	ErrSecMissingAttributeKey ErrSecSuccess = -67749
	// ErrSecMissingAttributeKeyLength: A key length attribute is missing.
	ErrSecMissingAttributeKeyLength ErrSecSuccess = -67763
	// ErrSecMissingAttributeKeyType: A key type attribute is missing.
	ErrSecMissingAttributeKeyType ErrSecSuccess = -67775
	// ErrSecMissingAttributeLabel: A label attribute is missing.
	ErrSecMissingAttributeLabel ErrSecSuccess = -67773
	// ErrSecMissingAttributeMode: A mode attribute is missing.
	ErrSecMissingAttributeMode ErrSecSuccess = -67777
	// ErrSecMissingAttributeOutputSize: An output size attribute is missing.
	ErrSecMissingAttributeOutputSize ErrSecSuccess = -67767
	// ErrSecMissingAttributePadding: A padding attribute is missing.
	ErrSecMissingAttributePadding ErrSecSuccess = -67755
	// ErrSecMissingAttributePassphrase: A passphrase attribute is missing.
	ErrSecMissingAttributePassphrase ErrSecSuccess = -67761
	// ErrSecMissingAttributePrime: A prime attribute is missing.
	ErrSecMissingAttributePrime ErrSecSuccess = -67787
	// ErrSecMissingAttributePrivateKeyFormat: A private key format attribute is missing.
	ErrSecMissingAttributePrivateKeyFormat ErrSecSuccess = -67801
	// ErrSecMissingAttributePublicKeyFormat: A public key format attribute is missing.
	ErrSecMissingAttributePublicKeyFormat ErrSecSuccess = -67799
	// ErrSecMissingAttributeRandom: A random number attribute is missing.
	ErrSecMissingAttributeRandom ErrSecSuccess = -67757
	// ErrSecMissingAttributeRounds: The number of rounds attribute is missing.
	ErrSecMissingAttributeRounds ErrSecSuccess = -67769
	// ErrSecMissingAttributeSalt: A salt attribute is missing.
	ErrSecMissingAttributeSalt ErrSecSuccess = -67753
	// ErrSecMissingAttributeSeed: A seed attribute is missing.
	ErrSecMissingAttributeSeed ErrSecSuccess = -67759
	// ErrSecMissingAttributeStartDate: A start date attribute is missing.
	ErrSecMissingAttributeStartDate ErrSecSuccess = -67781
	// ErrSecMissingAttributeSubprime: A subprime attribute is missing.
	ErrSecMissingAttributeSubprime ErrSecSuccess = -67791
	// ErrSecMissingAttributeSymmetricKeyFormat: A symmetric key format attribute is missing.
	ErrSecMissingAttributeSymmetricKeyFormat ErrSecSuccess = -67803
	// ErrSecMissingAttributeVersion: A version attribute is missing.
	ErrSecMissingAttributeVersion ErrSecSuccess = -67785
	// ErrSecMissingAttributeWrappedKeyFormat: A wrapped key format attribute is missing.
	ErrSecMissingAttributeWrappedKeyFormat ErrSecSuccess = -67805
	// ErrSecMissingEntitlement: A required entitlement is missing.
	ErrSecMissingEntitlement            ErrSecSuccess = -34018
	ErrSecMissingQualifiedCertStatement ErrSecSuccess = -67904
	// ErrSecMissingRequiredExtension: A required certificate extension is missing.
	ErrSecMissingRequiredExtension ErrSecSuccess = -67880
	// ErrSecMissingValue: A missing value was detected.
	ErrSecMissingValue ErrSecSuccess = -67871
	// ErrSecMobileMeCSRVerifyFailure: A MobileMe certificate signing request verification failure occurred.
	ErrSecMobileMeCSRVerifyFailure ErrSecSuccess = -67665
	// ErrSecMobileMeFailedConsistencyCheck: MobileMe found a failed consistency check.
	ErrSecMobileMeFailedConsistencyCheck ErrSecSuccess = -67666
	// ErrSecMobileMeNoRequestPending: MobileMe has no request pending.
	ErrSecMobileMeNoRequestPending ErrSecSuccess = -67664
	// ErrSecMobileMeRequestAlreadyPending: A MobileMe request is already pending.
	ErrSecMobileMeRequestAlreadyPending ErrSecSuccess = -67663
	// ErrSecMobileMeRequestQueued: The MobileMe request will be sent during the next connection.
	ErrSecMobileMeRequestQueued ErrSecSuccess = -67657
	// ErrSecMobileMeRequestRedirected: The MobileMe request was redirected.
	ErrSecMobileMeRequestRedirected ErrSecSuccess = -67658
	// ErrSecMobileMeServerAlreadyExists: The MobileMe server reported that the item already exists.
	ErrSecMobileMeServerAlreadyExists ErrSecSuccess = -67661
	// ErrSecMobileMeServerError: A MobileMe server error occurred.
	ErrSecMobileMeServerError ErrSecSuccess = -67659
	// ErrSecMobileMeServerNotAvailable: The MobileMe server is not available.
	ErrSecMobileMeServerNotAvailable ErrSecSuccess = -67660
	// ErrSecMobileMeServerServiceErr: A MobileMe service error occurred.
	ErrSecMobileMeServerServiceErr ErrSecSuccess = -67662
	// ErrSecModuleManagerInitializeFailed: A module failed to initialize.
	ErrSecModuleManagerInitializeFailed ErrSecSuccess = -67721
	// ErrSecModuleManagerNotFound: A module was not found.
	ErrSecModuleManagerNotFound ErrSecSuccess = -67722
	// ErrSecModuleManifestVerifyFailed: A module manifest verification failure occurred.
	ErrSecModuleManifestVerifyFailed ErrSecSuccess = -67678
	// ErrSecModuleNotLoaded: A module was not loaded.
	ErrSecModuleNotLoaded ErrSecSuccess = -67718
	// ErrSecMultiplePrivKeys: An attempt was made to import multiple private keys.
	ErrSecMultiplePrivKeys ErrSecSuccess = -25259
	// ErrSecMultipleValuesUnsupported: Multiple values are not supported.
	ErrSecMultipleValuesUnsupported ErrSecSuccess = -67842
	// ErrSecNetworkFailure: A network failure occurred.
	ErrSecNetworkFailure ErrSecSuccess = -67636
	// ErrSecNoAccessForItem: The specified item has no access control.
	ErrSecNoAccessForItem ErrSecSuccess = -25243
	// ErrSecNoBasicConstraints: No basic constraints were found.
	ErrSecNoBasicConstraints ErrSecSuccess = -67604
	// ErrSecNoBasicConstraintsCA: No basic CA constraints were found.
	ErrSecNoBasicConstraintsCA ErrSecSuccess = -67605
	// ErrSecNoCertificateModule: There is no certificate module available.
	ErrSecNoCertificateModule ErrSecSuccess = -25313
	// ErrSecNoDefaultAuthority: No default authority was detected.
	ErrSecNoDefaultAuthority ErrSecSuccess = -67844
	// ErrSecNoDefaultKeychain: A default keychain does not exist.
	ErrSecNoDefaultKeychain ErrSecSuccess = -25307
	// ErrSecNoFieldValues: No field values were detected.
	ErrSecNoFieldValues ErrSecSuccess = -67859
	// ErrSecNoPolicyModule: There is no policy module available.
	ErrSecNoPolicyModule ErrSecSuccess = -25314
	// ErrSecNoStorageModule: There is no storage module available.
	ErrSecNoStorageModule ErrSecSuccess = -25312
	// ErrSecNoSuchAttr: The attribute does not exist.
	ErrSecNoSuchAttr ErrSecSuccess = -25303
	// ErrSecNoSuchClass: The keychain item class does not exist.
	ErrSecNoSuchClass ErrSecSuccess = -25306
	// ErrSecNoSuchKeychain: The keychain does not exist.
	ErrSecNoSuchKeychain ErrSecSuccess = -25294
	// ErrSecNoTrustSettings: No trust settings were found.
	ErrSecNoTrustSettings ErrSecSuccess = -25263
	// ErrSecNotAvailable: No trust results are available.
	ErrSecNotAvailable ErrSecSuccess = -25291
	// ErrSecNotInitialized: A function was called without initializing the common security services manager.
	ErrSecNotInitialized ErrSecSuccess = -67667
	// ErrSecNotLoggedIn: You are not logged in.
	ErrSecNotLoggedIn ErrSecSuccess = -67729
	// ErrSecNotSigner: The certificate is not signed by its proposed parent.
	ErrSecNotSigner ErrSecSuccess = -26267
	// ErrSecNotTrusted: The trust policy is not trusted.
	ErrSecNotTrusted ErrSecSuccess = -67843
	// ErrSecOCSPBadRequest: The online certificate status protocol (OCSP) request is incorrect or cannot be parsed.
	ErrSecOCSPBadRequest ErrSecSuccess = -67631
	// ErrSecOCSPBadResponse: The online certificate status protocol (OCSP) response is incorrect or cannot be parsed.
	ErrSecOCSPBadResponse ErrSecSuccess = -67630
	// ErrSecOCSPNoSigner: The online certificate status protocol (OCSP) response has no signer.
	ErrSecOCSPNoSigner ErrSecSuccess = -67640
	// ErrSecOCSPNotTrustedToAnchor: The online certificate status protocol (OCSP) response is not trusted to a root or anchor certificate.
	ErrSecOCSPNotTrustedToAnchor ErrSecSuccess = -67637
	// ErrSecOCSPResponderInternalError: The online certificate status protocol (OCSP) responder detected an internal error.
	ErrSecOCSPResponderInternalError ErrSecSuccess = -67642
	// ErrSecOCSPResponderMalformedReq: The online certificate status protocol (OCSP) responder detected a malformed request.
	ErrSecOCSPResponderMalformedReq ErrSecSuccess = -67641
	// ErrSecOCSPResponderSignatureRequired: The online certificate status protocol (OCSP) responder requires a signature.
	ErrSecOCSPResponderSignatureRequired ErrSecSuccess = -67644
	// ErrSecOCSPResponderTryLater: The online certificate status protocol (OCSP) responder is busy, try again later.
	ErrSecOCSPResponderTryLater ErrSecSuccess = -67643
	// ErrSecOCSPResponderUnauthorized: The online certificate status protocol (OCSP) responder rejects the request as unauthorized.
	ErrSecOCSPResponderUnauthorized ErrSecSuccess = -67645
	// ErrSecOCSPResponseNonceMismatch: The online certificate status protocol (OCSP) response nonce does not match the request.
	ErrSecOCSPResponseNonceMismatch ErrSecSuccess = -67646
	// ErrSecOCSPSignatureError: The online certificate status protocol (OCSP) response has an invalid signature.
	ErrSecOCSPSignatureError ErrSecSuccess = -67639
	// ErrSecOCSPStatusUnrecognized: The online certificate status protocol (OCSP) server does not recognize this certificate.
	ErrSecOCSPStatusUnrecognized ErrSecSuccess = -67633
	// ErrSecOCSPUnavailable: The online certificate status protocol (OCSP) service is unavailable.
	ErrSecOCSPUnavailable ErrSecSuccess = -67632
	// ErrSecOpWr: The file is already open with write permission.
	ErrSecOpWr ErrSecSuccess = -49
	// ErrSecOutputLengthError: An output length error was detected.
	ErrSecOutputLengthError ErrSecSuccess = -67725
	// ErrSecPVCAlreadyConfigured: The PVC is already configured.
	ErrSecPVCAlreadyConfigured ErrSecSuccess = -67707
	// ErrSecPVCReferentNotFound: A reference to the calling module was not found in the list of authorized callers.
	ErrSecPVCReferentNotFound ErrSecSuccess = -67669
	// ErrSecParam: One or more parameters passed to the function are not valid.
	ErrSecParam ErrSecSuccess = -50
	// ErrSecPassphraseRequired: A password is required for import or export.
	ErrSecPassphraseRequired ErrSecSuccess = -25260
	// ErrSecPathLengthConstraintExceeded: The path length constraint was exceeded.
	ErrSecPathLengthConstraintExceeded ErrSecSuccess = -67611
	// ErrSecPkcs12VerifyFailure: MAC verification failed during PKCS12 Import.
	ErrSecPkcs12VerifyFailure ErrSecSuccess = -25264
	// ErrSecPolicyNotFound: The specified policy cannot be found.
	ErrSecPolicyNotFound ErrSecSuccess = -25241
	// ErrSecPrivilegeNotGranted: The privilege is not granted.
	ErrSecPrivilegeNotGranted ErrSecSuccess = -67705
	// ErrSecPrivilegeNotSupported: The privilege is not supported.
	ErrSecPrivilegeNotSupported ErrSecSuccess = -67726
	// ErrSecPublicKeyInconsistent: The public key is inconsistent.
	ErrSecPublicKeyInconsistent ErrSecSuccess = -67811
	// ErrSecQuerySizeUnknown: The query size is unknown.
	ErrSecQuerySizeUnknown ErrSecSuccess = -67809
	// ErrSecQuotaExceeded: The quota was exceeded.
	ErrSecQuotaExceeded ErrSecSuccess = -67596
	// ErrSecReadOnly: Read-only error.
	ErrSecReadOnly ErrSecSuccess = -25292
	// ErrSecReadOnlyAttr: The attribute is read-only.
	ErrSecReadOnlyAttr ErrSecSuccess = -25309
	// ErrSecRecordModified: The record is modified.
	ErrSecRecordModified ErrSecSuccess = -67638
	// ErrSecRejectedForm: The trust policy has a rejected form.
	ErrSecRejectedForm ErrSecSuccess = -67845
	// ErrSecRequestDescriptor: The request descriptor is not valid.
	ErrSecRequestDescriptor ErrSecSuccess = -67856
	// ErrSecRequestLost: The request is lost.
	ErrSecRequestLost ErrSecSuccess = -67846
	// ErrSecRequestRejected: The request is rejected.
	ErrSecRequestRejected ErrSecSuccess = -67847
	// ErrSecResourceSignBadCertChainLength: Resource signing detects an incorrect certificate chain length.
	ErrSecResourceSignBadCertChainLength ErrSecSuccess = -67652
	// ErrSecResourceSignBadExtKeyUsage: Resource signing detects an error in the extended key usage.
	ErrSecResourceSignBadExtKeyUsage ErrSecSuccess = -67653
	ErrSecRestrictedAPI              ErrSecSuccess = -34020
	// ErrSecSMIMEBadExtendedKeyUsage: The appropriate extended key usage for SMIME is not found.
	ErrSecSMIMEBadExtendedKeyUsage ErrSecSuccess = -67624
	// ErrSecSMIMEBadKeyUsage: The key usage is not compatible with SMIME.
	ErrSecSMIMEBadKeyUsage ErrSecSuccess = -67625
	// ErrSecSMIMEEmailAddressesNotFound: An email address mismatch was detected.
	ErrSecSMIMEEmailAddressesNotFound ErrSecSuccess = -67623
	// ErrSecSMIMEKeyUsageNotCritical: The key usage extension is not marked as critical.
	ErrSecSMIMEKeyUsageNotCritical ErrSecSuccess = -67626
	// ErrSecSMIMENoEmailAddress: No email address is found in the certificate.
	ErrSecSMIMENoEmailAddress ErrSecSuccess = -67627
	// ErrSecSMIMESubjAltNameNotCritical: The subject alternative name extension is not marked as critical.
	ErrSecSMIMESubjAltNameNotCritical ErrSecSuccess = -67628
	// ErrSecSSLBadExtendedKeyUsage: The appropriate extended key usage for SSL is not found.
	ErrSecSSLBadExtendedKeyUsage ErrSecSuccess = -67629
	// ErrSecSelfCheckFailed: Self-check failed.
	ErrSecSelfCheckFailed ErrSecSuccess = -67676
	// ErrSecServiceNotAvailable: Self-check failed.
	ErrSecServiceNotAvailable ErrSecSuccess = -67585
	// ErrSecSigningTimeMissing: A signing time is missing.
	ErrSecSigningTimeMissing ErrSecSuccess = -67894
	// ErrSecStagedOperationInProgress: A staged operation is in progress.
	ErrSecStagedOperationInProgress ErrSecSuccess = -67806
	// ErrSecStagedOperationNotStarted: A staged operation was not started.
	ErrSecStagedOperationNotStarted ErrSecSuccess = -67807
	// ErrSecSuccessValue: No error.
	ErrSecSuccessValue ErrSecSuccess = 0
	// ErrSecTagNotFound: The specified tag is not found.
	ErrSecTagNotFound ErrSecSuccess = -67692
	// ErrSecTimestampAddInfoNotAvailable: The additional information requested is not available.
	ErrSecTimestampAddInfoNotAvailable ErrSecSuccess = -67892
	// ErrSecTimestampBadAlg: Found an unrecognized or unsupported algorithm identifier (AI) in timestamp.
	ErrSecTimestampBadAlg ErrSecSuccess = -67886
	// ErrSecTimestampBadDataFormat: The timestamp data submitted has the wrong format.
	ErrSecTimestampBadDataFormat ErrSecSuccess = -67888
	// ErrSecTimestampBadRequest: The timestamp transaction is not permitted or supported.
	ErrSecTimestampBadRequest ErrSecSuccess = -67887
	// ErrSecTimestampInvalid: The timestamp is not valid.
	ErrSecTimestampInvalid ErrSecSuccess = -67883
	// ErrSecTimestampMissing: A timestamp is expected but is not found.
	ErrSecTimestampMissing ErrSecSuccess = -67882
	// ErrSecTimestampNotTrusted: The timestamp is not trusted.
	ErrSecTimestampNotTrusted ErrSecSuccess = -67884
	// ErrSecTimestampRejection: A timestamp transaction is rejected.
	ErrSecTimestampRejection ErrSecSuccess = -67895
	// ErrSecTimestampRevocationNotification: A timestamp authority revocation notification is issued.
	ErrSecTimestampRevocationNotification ErrSecSuccess = -67898
	// ErrSecTimestampRevocationWarning: A timestamp authority revocation warning is issued.
	ErrSecTimestampRevocationWarning   ErrSecSuccess = -67897
	ErrSecTimestampServiceNotAvailable ErrSecSuccess = -67885
	// ErrSecTimestampSystemFailure: The timestamp request cannot be handled due to a system failure.
	ErrSecTimestampSystemFailure ErrSecSuccess = -67893
	// ErrSecTimestampTimeNotAvailable: The time source for the timestamp authority is not available.
	ErrSecTimestampTimeNotAvailable ErrSecSuccess = -67889
	// ErrSecTimestampUnacceptedExtension: The requested extension is not supported by the timestamp authority.
	ErrSecTimestampUnacceptedExtension ErrSecSuccess = -67891
	// ErrSecTimestampUnacceptedPolicy: The requested policy is not supported by the timestamp authority.
	ErrSecTimestampUnacceptedPolicy ErrSecSuccess = -67890
	// ErrSecTimestampWaiting: A timestamp transaction is waiting.
	ErrSecTimestampWaiting ErrSecSuccess = -67896
	// ErrSecTrustNotAvailable: No trust results are available.
	ErrSecTrustNotAvailable ErrSecSuccess = -25245
	// ErrSecTrustSettingDeny: The trust setting for this policy is set to Deny.
	ErrSecTrustSettingDeny ErrSecSuccess = -67654
	// ErrSecUnimplemented: A function or operation is not implemented.
	ErrSecUnimplemented ErrSecSuccess = -4
	// ErrSecUnknownCRLExtension: An unknown certificate revocation list extension was detected.
	ErrSecUnknownCRLExtension ErrSecSuccess = -67619
	// ErrSecUnknownCertExtension: An unknown certificate extension was detected.
	ErrSecUnknownCertExtension ErrSecSuccess = -67618
	// ErrSecUnknownCriticalExtensionFlag: There is an unknown critical extension flag.
	ErrSecUnknownCriticalExtensionFlag ErrSecSuccess = -67603
	// ErrSecUnknownFormat: The item you are trying to import has an unknown format.
	ErrSecUnknownFormat ErrSecSuccess = -25257
	// ErrSecUnknownQualifiedCertStatement: An unknown qualified certificate statement was detected.
	ErrSecUnknownQualifiedCertStatement ErrSecSuccess = -67656
	// ErrSecUnknownTag: An unknown tag was detected.
	ErrSecUnknownTag ErrSecSuccess = -67687
	// ErrSecUnsupportedAddressType: The address type is not supported.
	ErrSecUnsupportedAddressType ErrSecSuccess = -67848
	// ErrSecUnsupportedFieldFormat: The field format is not supported.
	ErrSecUnsupportedFieldFormat ErrSecSuccess = -67860
	// ErrSecUnsupportedFormat: The specified import or export format is not supported.
	ErrSecUnsupportedFormat ErrSecSuccess = -25256
	// ErrSecUnsupportedIndexInfo: The index information is not supported.
	ErrSecUnsupportedIndexInfo ErrSecSuccess = -67861
	// ErrSecUnsupportedKeyAttributeMask: The key attribute mask is not supported.
	ErrSecUnsupportedKeyAttributeMask ErrSecSuccess = -67739
	// ErrSecUnsupportedKeyFormat: The key header format is not supported.
	ErrSecUnsupportedKeyFormat ErrSecSuccess = -67734
	// ErrSecUnsupportedKeyLabel: The key label is not supported.
	ErrSecUnsupportedKeyLabel ErrSecSuccess = -67741
	// ErrSecUnsupportedKeySize: The key size is not supported.
	ErrSecUnsupportedKeySize ErrSecSuccess = -67735
	// ErrSecUnsupportedKeyUsageMask: The key usage mask is not supported.
	ErrSecUnsupportedKeyUsageMask ErrSecSuccess = -67737
	// ErrSecUnsupportedLocality: The locality is not supported.
	ErrSecUnsupportedLocality ErrSecSuccess = -67862
	// ErrSecUnsupportedNumAttributes: The number of attributes is not supported.
	ErrSecUnsupportedNumAttributes ErrSecSuccess = -67863
	// ErrSecUnsupportedNumIndexes: The number of indexes is not supported.
	ErrSecUnsupportedNumIndexes ErrSecSuccess = -67864
	// ErrSecUnsupportedNumRecordTypes: The number of record types is not supported.
	ErrSecUnsupportedNumRecordTypes ErrSecSuccess = -67865
	// ErrSecUnsupportedNumSelectionPreds: The number of selection predicates is not supported.
	ErrSecUnsupportedNumSelectionPreds ErrSecSuccess = -67873
	// ErrSecUnsupportedOperator: The operator is not supported.
	ErrSecUnsupportedOperator ErrSecSuccess = -67874
	// ErrSecUnsupportedQueryLimits: The query limits are not supported.
	ErrSecUnsupportedQueryLimits ErrSecSuccess = -67872
	// ErrSecUnsupportedService: The service is not supported.
	ErrSecUnsupportedService ErrSecSuccess = -67849
	// ErrSecUnsupportedVectorOfBuffers: The vector of buffers is not supported.
	ErrSecUnsupportedVectorOfBuffers ErrSecSuccess = -67743
	// ErrSecUserCanceled: User canceled the operation.
	ErrSecUserCanceled ErrSecSuccess = -128
	// ErrSecVerificationFailure: A verification failure occurred.
	ErrSecVerificationFailure ErrSecSuccess = -67686
	// ErrSecVerifyActionFailed: A verify action failed.
	ErrSecVerifyActionFailed ErrSecSuccess = -67825
	// ErrSecVerifyFailed: A cryptographic verification failure occurred.
	ErrSecVerifyFailed ErrSecSuccess = -67808
	// ErrSecWrPerm: Write permissions error.
	ErrSecWrPerm ErrSecSuccess = -61
	// ErrSecWrongSecVersion: The version is incorrect.
	ErrSecWrongSecVersion ErrSecSuccess = -25310
	// Deprecated.
	ErrSecDskFull ErrSecSuccess = -34
	// Deprecated: use ErrSecInvalidCRLAuthority.
	ErrSecInvaldCRLAuthority ErrSecSuccess = -67827
	// Deprecated: use ErrSecInvalidTupleCredentials.
	ErrSecInvalidTupleCredendtials ErrSecSuccess = -67852
)

func (e ErrSecSuccess) String() string {
	switch e {
	case ErrSecACLAddFailed:
		return "ErrSecACLAddFailed"
	case ErrSecACLChangeFailed:
		return "ErrSecACLChangeFailed"
	case ErrSecACLDeleteFailed:
		return "ErrSecACLDeleteFailed"
	case ErrSecACLNotSimple:
		return "ErrSecACLNotSimple"
	case ErrSecACLReplaceFailed:
		return "ErrSecACLReplaceFailed"
	case ErrSecAddinLoadFailed:
		return "ErrSecAddinLoadFailed"
	case ErrSecAddinUnloadFailed:
		return "ErrSecAddinUnloadFailed"
	case ErrSecAlgorithmMismatch:
		return "ErrSecAlgorithmMismatch"
	case ErrSecAllocate:
		return "ErrSecAllocate"
	case ErrSecAlreadyLoggedIn:
		return "ErrSecAlreadyLoggedIn"
	case ErrSecAppleAddAppACLSubject:
		return "ErrSecAppleAddAppACLSubject"
	case ErrSecAppleInvalidKeyEndDate:
		return "ErrSecAppleInvalidKeyEndDate"
	case ErrSecAppleInvalidKeyStartDate:
		return "ErrSecAppleInvalidKeyStartDate"
	case ErrSecApplePublicKeyIncomplete:
		return "ErrSecApplePublicKeyIncomplete"
	case ErrSecAppleSSLv2Rollback:
		return "ErrSecAppleSSLv2Rollback"
	case ErrSecAppleSignatureMismatch:
		return "ErrSecAppleSignatureMismatch"
	case ErrSecAttachHandleBusy:
		return "ErrSecAttachHandleBusy"
	case ErrSecAttributeNotInContext:
		return "ErrSecAttributeNotInContext"
	case ErrSecAuthFailed:
		return "ErrSecAuthFailed"
	case ErrSecBadReq:
		return "ErrSecBadReq"
	case ErrSecBlockSizeMismatch:
		return "ErrSecBlockSizeMismatch"
	case ErrSecBufferTooSmall:
		return "ErrSecBufferTooSmall"
	case ErrSecCRLAlreadySigned:
		return "ErrSecCRLAlreadySigned"
	case ErrSecCRLBadURI:
		return "ErrSecCRLBadURI"
	case ErrSecCRLExpired:
		return "ErrSecCRLExpired"
	case ErrSecCRLNotFound:
		return "ErrSecCRLNotFound"
	case ErrSecCRLNotTrusted:
		return "ErrSecCRLNotTrusted"
	case ErrSecCRLNotValidYet:
		return "ErrSecCRLNotValidYet"
	case ErrSecCRLPolicyFailed:
		return "ErrSecCRLPolicyFailed"
	case ErrSecCRLServerDown:
		return "ErrSecCRLServerDown"
	case ErrSecCallbackFailed:
		return "ErrSecCallbackFailed"
	case ErrSecCertificateCannotOperate:
		return "ErrSecCertificateCannotOperate"
	case ErrSecCertificateDuplicateExtension:
		return "ErrSecCertificateDuplicateExtension"
	case ErrSecCertificateExpired:
		return "ErrSecCertificateExpired"
	case ErrSecCertificateIsCA:
		return "ErrSecCertificateIsCA"
	case ErrSecCertificateNameNotAllowed:
		return "ErrSecCertificateNameNotAllowed"
	case ErrSecCertificateNotValidYet:
		return "ErrSecCertificateNotValidYet"
	case ErrSecCertificatePolicyNotAllowed:
		return "ErrSecCertificatePolicyNotAllowed"
	case ErrSecCertificateRevoked:
		return "ErrSecCertificateRevoked"
	case ErrSecCertificateSuspended:
		return "ErrSecCertificateSuspended"
	case ErrSecCertificateValidityPeriodTooLong:
		return "ErrSecCertificateValidityPeriodTooLong"
	case ErrSecCodeSigningBadCertChainLength:
		return "ErrSecCodeSigningBadCertChainLength"
	case ErrSecCodeSigningBadPathLengthConstraint:
		return "ErrSecCodeSigningBadPathLengthConstraint"
	case ErrSecCodeSigningDevelopment:
		return "ErrSecCodeSigningDevelopment"
	case ErrSecCodeSigningNoBasicConstraints:
		return "ErrSecCodeSigningNoBasicConstraints"
	case ErrSecCodeSigningNoExtendedKeyUsage:
		return "ErrSecCodeSigningNoExtendedKeyUsage"
	case ErrSecConversionError:
		return "ErrSecConversionError"
	case ErrSecCoreFoundationUnknown:
		return "ErrSecCoreFoundationUnknown"
	case ErrSecCreateChainFailed:
		return "ErrSecCreateChainFailed"
	case ErrSecDataNotAvailable:
		return "ErrSecDataNotAvailable"
	case ErrSecDataNotModifiable:
		return "ErrSecDataNotModifiable"
	case ErrSecDataTooLarge:
		return "ErrSecDataTooLarge"
	case ErrSecDatabaseLocked:
		return "ErrSecDatabaseLocked"
	case ErrSecDatastoreIsOpen:
		return "ErrSecDatastoreIsOpen"
	case ErrSecDecode:
		return "ErrSecDecode"
	case ErrSecDeviceError:
		return "ErrSecDeviceError"
	case ErrSecDeviceFailed:
		return "ErrSecDeviceFailed"
	case ErrSecDeviceReset:
		return "ErrSecDeviceReset"
	case ErrSecDeviceVerifyFailed:
		return "ErrSecDeviceVerifyFailed"
	case ErrSecDiskFull:
		return "ErrSecDiskFull"
	case ErrSecDuplicateCallback:
		return "ErrSecDuplicateCallback"
	case ErrSecDuplicateItem:
		return "ErrSecDuplicateItem"
	case ErrSecDuplicateKeychain:
		return "ErrSecDuplicateKeychain"
	case ErrSecEMMLoadFailed:
		return "ErrSecEMMLoadFailed"
	case ErrSecEMMUnloadFailed:
		return "ErrSecEMMUnloadFailed"
	case ErrSecEndOfData:
		return "ErrSecEndOfData"
	case ErrSecEventNotificationCallbackNotFound:
		return "ErrSecEventNotificationCallbackNotFound"
	case ErrSecExtendedKeyUsageNotCritical:
		return "ErrSecExtendedKeyUsageNotCritical"
	case ErrSecFieldSpecifiedMultiple:
		return "ErrSecFieldSpecifiedMultiple"
	case ErrSecFileTooBig:
		return "ErrSecFileTooBig"
	case ErrSecFunctionFailed:
		return "ErrSecFunctionFailed"
	case ErrSecFunctionIntegrityFail:
		return "ErrSecFunctionIntegrityFail"
	case ErrSecHostNameMismatch:
		return "ErrSecHostNameMismatch"
	case ErrSecIDPFailure:
		return "ErrSecIDPFailure"
	case ErrSecIO:
		return "ErrSecIO"
	case ErrSecInDarkWake:
		return "ErrSecInDarkWake"
	case ErrSecIncompatibleDatabaseBlob:
		return "ErrSecIncompatibleDatabaseBlob"
	case ErrSecIncompatibleFieldFormat:
		return "ErrSecIncompatibleFieldFormat"
	case ErrSecIncompatibleKeyBlob:
		return "ErrSecIncompatibleKeyBlob"
	case ErrSecIncompatibleVersion:
		return "ErrSecIncompatibleVersion"
	case ErrSecIncompleteCertRevocationCheck:
		return "ErrSecIncompleteCertRevocationCheck"
	case ErrSecInputLengthError:
		return "ErrSecInputLengthError"
	case ErrSecInsufficientClientID:
		return "ErrSecInsufficientClientID"
	case ErrSecInsufficientCredentials:
		return "ErrSecInsufficientCredentials"
	case ErrSecInteractionNotAllowed:
		return "ErrSecInteractionNotAllowed"
	case ErrSecInteractionRequired:
		return "ErrSecInteractionRequired"
	case ErrSecInternalComponent:
		return "ErrSecInternalComponent"
	case ErrSecInternalError:
		return "ErrSecInternalError"
	case ErrSecInvalidACL:
		return "ErrSecInvalidACL"
	case ErrSecInvalidAccessCredentials:
		return "ErrSecInvalidAccessCredentials"
	case ErrSecInvalidAccessRequest:
		return "ErrSecInvalidAccessRequest"
	case ErrSecInvalidAction:
		return "ErrSecInvalidAction"
	case ErrSecInvalidAddinFunctionTable:
		return "ErrSecInvalidAddinFunctionTable"
	case ErrSecInvalidAlgorithm:
		return "ErrSecInvalidAlgorithm"
	case ErrSecInvalidAlgorithmParms:
		return "ErrSecInvalidAlgorithmParms"
	case ErrSecInvalidAttributeAccessCredentials:
		return "ErrSecInvalidAttributeAccessCredentials"
	case ErrSecInvalidAttributeBase:
		return "ErrSecInvalidAttributeBase"
	case ErrSecInvalidAttributeBlockSize:
		return "ErrSecInvalidAttributeBlockSize"
	case ErrSecInvalidAttributeDLDBHandle:
		return "ErrSecInvalidAttributeDLDBHandle"
	case ErrSecInvalidAttributeEffectiveBits:
		return "ErrSecInvalidAttributeEffectiveBits"
	case ErrSecInvalidAttributeEndDate:
		return "ErrSecInvalidAttributeEndDate"
	case ErrSecInvalidAttributeInitVector:
		return "ErrSecInvalidAttributeInitVector"
	case ErrSecInvalidAttributeIterationCount:
		return "ErrSecInvalidAttributeIterationCount"
	case ErrSecInvalidAttributeKey:
		return "ErrSecInvalidAttributeKey"
	case ErrSecInvalidAttributeKeyLength:
		return "ErrSecInvalidAttributeKeyLength"
	case ErrSecInvalidAttributeKeyType:
		return "ErrSecInvalidAttributeKeyType"
	case ErrSecInvalidAttributeLabel:
		return "ErrSecInvalidAttributeLabel"
	case ErrSecInvalidAttributeMode:
		return "ErrSecInvalidAttributeMode"
	case ErrSecInvalidAttributeOutputSize:
		return "ErrSecInvalidAttributeOutputSize"
	case ErrSecInvalidAttributePadding:
		return "ErrSecInvalidAttributePadding"
	case ErrSecInvalidAttributePassphrase:
		return "ErrSecInvalidAttributePassphrase"
	case ErrSecInvalidAttributePrime:
		return "ErrSecInvalidAttributePrime"
	case ErrSecInvalidAttributePrivateKeyFormat:
		return "ErrSecInvalidAttributePrivateKeyFormat"
	case ErrSecInvalidAttributePublicKeyFormat:
		return "ErrSecInvalidAttributePublicKeyFormat"
	case ErrSecInvalidAttributeRandom:
		return "ErrSecInvalidAttributeRandom"
	case ErrSecInvalidAttributeRounds:
		return "ErrSecInvalidAttributeRounds"
	case ErrSecInvalidAttributeSalt:
		return "ErrSecInvalidAttributeSalt"
	case ErrSecInvalidAttributeSeed:
		return "ErrSecInvalidAttributeSeed"
	case ErrSecInvalidAttributeStartDate:
		return "ErrSecInvalidAttributeStartDate"
	case ErrSecInvalidAttributeSubprime:
		return "ErrSecInvalidAttributeSubprime"
	case ErrSecInvalidAttributeSymmetricKeyFormat:
		return "ErrSecInvalidAttributeSymmetricKeyFormat"
	case ErrSecInvalidAttributeVersion:
		return "ErrSecInvalidAttributeVersion"
	case ErrSecInvalidAttributeWrappedKeyFormat:
		return "ErrSecInvalidAttributeWrappedKeyFormat"
	case ErrSecInvalidAuthority:
		return "ErrSecInvalidAuthority"
	case ErrSecInvalidAuthorityKeyID:
		return "ErrSecInvalidAuthorityKeyID"
	case ErrSecInvalidBaseACLs:
		return "ErrSecInvalidBaseACLs"
	case ErrSecInvalidBundleInfo:
		return "ErrSecInvalidBundleInfo"
	case ErrSecInvalidCRL:
		return "ErrSecInvalidCRL"
	case ErrSecInvalidCRLAuthority:
		return "ErrSecInvalidCRLAuthority"
	case ErrSecInvalidCRLEncoding:
		return "ErrSecInvalidCRLEncoding"
	case ErrSecInvalidCRLGroup:
		return "ErrSecInvalidCRLGroup"
	case ErrSecInvalidCRLIndex:
		return "ErrSecInvalidCRLIndex"
	case ErrSecInvalidCRLType:
		return "ErrSecInvalidCRLType"
	case ErrSecInvalidCallback:
		return "ErrSecInvalidCallback"
	case ErrSecInvalidCertAuthority:
		return "ErrSecInvalidCertAuthority"
	case ErrSecInvalidCertificateGroup:
		return "ErrSecInvalidCertificateGroup"
	case ErrSecInvalidCertificateRef:
		return "ErrSecInvalidCertificateRef"
	case ErrSecInvalidContext:
		return "ErrSecInvalidContext"
	case ErrSecInvalidDBList:
		return "ErrSecInvalidDBList"
	case ErrSecInvalidDBLocation:
		return "ErrSecInvalidDBLocation"
	case ErrSecInvalidData:
		return "ErrSecInvalidData"
	case ErrSecInvalidDatabaseBlob:
		return "ErrSecInvalidDatabaseBlob"
	case ErrSecInvalidDigestAlgorithm:
		return "ErrSecInvalidDigestAlgorithm"
	case ErrSecInvalidEncoding:
		return "ErrSecInvalidEncoding"
	case ErrSecInvalidExtendedKeyUsage:
		return "ErrSecInvalidExtendedKeyUsage"
	case ErrSecInvalidFormType:
		return "ErrSecInvalidFormType"
	case ErrSecInvalidGUID:
		return "ErrSecInvalidGUID"
	case ErrSecInvalidHandle:
		return "ErrSecInvalidHandle"
	case ErrSecInvalidHandleUsage:
		return "ErrSecInvalidHandleUsage"
	case ErrSecInvalidID:
		return "ErrSecInvalidID"
	case ErrSecInvalidIDLinkage:
		return "ErrSecInvalidIDLinkage"
	case ErrSecInvalidIdentifier:
		return "ErrSecInvalidIdentifier"
	case ErrSecInvalidIndex:
		return "ErrSecInvalidIndex"
	case ErrSecInvalidIndexInfo:
		return "ErrSecInvalidIndexInfo"
	case ErrSecInvalidInputVector:
		return "ErrSecInvalidInputVector"
	case ErrSecInvalidItemRef:
		return "ErrSecInvalidItemRef"
	case ErrSecInvalidKeyAttributeMask:
		return "ErrSecInvalidKeyAttributeMask"
	case ErrSecInvalidKeyBlob:
		return "ErrSecInvalidKeyBlob"
	case ErrSecInvalidKeyFormat:
		return "ErrSecInvalidKeyFormat"
	case ErrSecInvalidKeyHierarchy:
		return "ErrSecInvalidKeyHierarchy"
	case ErrSecInvalidKeyLabel:
		return "ErrSecInvalidKeyLabel"
	case ErrSecInvalidKeyRef:
		return "ErrSecInvalidKeyRef"
	case ErrSecInvalidKeyUsageForPolicy:
		return "ErrSecInvalidKeyUsageForPolicy"
	case ErrSecInvalidKeyUsageMask:
		return "ErrSecInvalidKeyUsageMask"
	case ErrSecInvalidKeychain:
		return "ErrSecInvalidKeychain"
	case ErrSecInvalidLoginName:
		return "ErrSecInvalidLoginName"
	case ErrSecInvalidModifyMode:
		return "ErrSecInvalidModifyMode"
	case ErrSecInvalidName:
		return "ErrSecInvalidName"
	case ErrSecInvalidNetworkAddress:
		return "ErrSecInvalidNetworkAddress"
	case ErrSecInvalidNewOwner:
		return "ErrSecInvalidNewOwner"
	case ErrSecInvalidNumberOfFields:
		return "ErrSecInvalidNumberOfFields"
	case ErrSecInvalidOutputVector:
		return "ErrSecInvalidOutputVector"
	case ErrSecInvalidOwnerEdit:
		return "ErrSecInvalidOwnerEdit"
	case ErrSecInvalidPVC:
		return "ErrSecInvalidPVC"
	case ErrSecInvalidParsingModule:
		return "ErrSecInvalidParsingModule"
	case ErrSecInvalidPassthroughID:
		return "ErrSecInvalidPassthroughID"
	case ErrSecInvalidPasswordRef:
		return "ErrSecInvalidPasswordRef"
	case ErrSecInvalidPointer:
		return "ErrSecInvalidPointer"
	case ErrSecInvalidPolicyIdentifiers:
		return "ErrSecInvalidPolicyIdentifiers"
	case ErrSecInvalidPrefsDomain:
		return "ErrSecInvalidPrefsDomain"
	case ErrSecInvalidQuery:
		return "ErrSecInvalidQuery"
	case ErrSecInvalidReason:
		return "ErrSecInvalidReason"
	case ErrSecInvalidRecord:
		return "ErrSecInvalidRecord"
	case ErrSecInvalidRequestInputs:
		return "ErrSecInvalidRequestInputs"
	case ErrSecInvalidRequestor:
		return "ErrSecInvalidRequestor"
	case ErrSecInvalidResponseVector:
		return "ErrSecInvalidResponseVector"
	case ErrSecInvalidRoot:
		return "ErrSecInvalidRoot"
	case ErrSecInvalidSampleValue:
		return "ErrSecInvalidSampleValue"
	case ErrSecInvalidScope:
		return "ErrSecInvalidScope"
	case ErrSecInvalidSearchRef:
		return "ErrSecInvalidSearchRef"
	case ErrSecInvalidServiceMask:
		return "ErrSecInvalidServiceMask"
	case ErrSecInvalidSignature:
		return "ErrSecInvalidSignature"
	case ErrSecInvalidStopOnPolicy:
		return "ErrSecInvalidStopOnPolicy"
	case ErrSecInvalidSubServiceID:
		return "ErrSecInvalidSubServiceID"
	case ErrSecInvalidSubjectKeyID:
		return "ErrSecInvalidSubjectKeyID"
	case ErrSecInvalidSubjectName:
		return "ErrSecInvalidSubjectName"
	case ErrSecInvalidTimeString:
		return "ErrSecInvalidTimeString"
	case ErrSecInvalidTrustSetting:
		return "ErrSecInvalidTrustSetting"
	case ErrSecInvalidTrustSettings:
		return "ErrSecInvalidTrustSettings"
	case ErrSecInvalidTuple:
		return "ErrSecInvalidTuple"
	case ErrSecInvalidTupleCredentials:
		return "ErrSecInvalidTupleCredentials"
	case ErrSecInvalidTupleGroup:
		return "ErrSecInvalidTupleGroup"
	case ErrSecInvalidValidityPeriod:
		return "ErrSecInvalidValidityPeriod"
	case ErrSecInvalidValue:
		return "ErrSecInvalidValue"
	case ErrSecItemNotFound:
		return "ErrSecItemNotFound"
	case ErrSecKeyBlobTypeIncorrect:
		return "ErrSecKeyBlobTypeIncorrect"
	case ErrSecKeyHeaderInconsistent:
		return "ErrSecKeyHeaderInconsistent"
	case ErrSecKeyIsSensitive:
		return "ErrSecKeyIsSensitive"
	case ErrSecKeySizeNotAllowed:
		return "ErrSecKeySizeNotAllowed"
	case ErrSecKeyUsageIncorrect:
		return "ErrSecKeyUsageIncorrect"
	case ErrSecLibraryReferenceNotFound:
		return "ErrSecLibraryReferenceNotFound"
	case ErrSecMDSError:
		return "ErrSecMDSError"
	case ErrSecMemoryError:
		return "ErrSecMemoryError"
	case ErrSecMissingAlgorithmParms:
		return "ErrSecMissingAlgorithmParms"
	case ErrSecMissingAttributeAccessCredentials:
		return "ErrSecMissingAttributeAccessCredentials"
	case ErrSecMissingAttributeBase:
		return "ErrSecMissingAttributeBase"
	case ErrSecMissingAttributeBlockSize:
		return "ErrSecMissingAttributeBlockSize"
	case ErrSecMissingAttributeDLDBHandle:
		return "ErrSecMissingAttributeDLDBHandle"
	case ErrSecMissingAttributeEffectiveBits:
		return "ErrSecMissingAttributeEffectiveBits"
	case ErrSecMissingAttributeEndDate:
		return "ErrSecMissingAttributeEndDate"
	case ErrSecMissingAttributeInitVector:
		return "ErrSecMissingAttributeInitVector"
	case ErrSecMissingAttributeIterationCount:
		return "ErrSecMissingAttributeIterationCount"
	case ErrSecMissingAttributeKey:
		return "ErrSecMissingAttributeKey"
	case ErrSecMissingAttributeKeyLength:
		return "ErrSecMissingAttributeKeyLength"
	case ErrSecMissingAttributeKeyType:
		return "ErrSecMissingAttributeKeyType"
	case ErrSecMissingAttributeLabel:
		return "ErrSecMissingAttributeLabel"
	case ErrSecMissingAttributeMode:
		return "ErrSecMissingAttributeMode"
	case ErrSecMissingAttributeOutputSize:
		return "ErrSecMissingAttributeOutputSize"
	case ErrSecMissingAttributePadding:
		return "ErrSecMissingAttributePadding"
	case ErrSecMissingAttributePassphrase:
		return "ErrSecMissingAttributePassphrase"
	case ErrSecMissingAttributePrime:
		return "ErrSecMissingAttributePrime"
	case ErrSecMissingAttributePrivateKeyFormat:
		return "ErrSecMissingAttributePrivateKeyFormat"
	case ErrSecMissingAttributePublicKeyFormat:
		return "ErrSecMissingAttributePublicKeyFormat"
	case ErrSecMissingAttributeRandom:
		return "ErrSecMissingAttributeRandom"
	case ErrSecMissingAttributeRounds:
		return "ErrSecMissingAttributeRounds"
	case ErrSecMissingAttributeSalt:
		return "ErrSecMissingAttributeSalt"
	case ErrSecMissingAttributeSeed:
		return "ErrSecMissingAttributeSeed"
	case ErrSecMissingAttributeStartDate:
		return "ErrSecMissingAttributeStartDate"
	case ErrSecMissingAttributeSubprime:
		return "ErrSecMissingAttributeSubprime"
	case ErrSecMissingAttributeSymmetricKeyFormat:
		return "ErrSecMissingAttributeSymmetricKeyFormat"
	case ErrSecMissingAttributeVersion:
		return "ErrSecMissingAttributeVersion"
	case ErrSecMissingAttributeWrappedKeyFormat:
		return "ErrSecMissingAttributeWrappedKeyFormat"
	case ErrSecMissingEntitlement:
		return "ErrSecMissingEntitlement"
	case ErrSecMissingQualifiedCertStatement:
		return "ErrSecMissingQualifiedCertStatement"
	case ErrSecMissingRequiredExtension:
		return "ErrSecMissingRequiredExtension"
	case ErrSecMissingValue:
		return "ErrSecMissingValue"
	case ErrSecMobileMeCSRVerifyFailure:
		return "ErrSecMobileMeCSRVerifyFailure"
	case ErrSecMobileMeFailedConsistencyCheck:
		return "ErrSecMobileMeFailedConsistencyCheck"
	case ErrSecMobileMeNoRequestPending:
		return "ErrSecMobileMeNoRequestPending"
	case ErrSecMobileMeRequestAlreadyPending:
		return "ErrSecMobileMeRequestAlreadyPending"
	case ErrSecMobileMeRequestQueued:
		return "ErrSecMobileMeRequestQueued"
	case ErrSecMobileMeRequestRedirected:
		return "ErrSecMobileMeRequestRedirected"
	case ErrSecMobileMeServerAlreadyExists:
		return "ErrSecMobileMeServerAlreadyExists"
	case ErrSecMobileMeServerError:
		return "ErrSecMobileMeServerError"
	case ErrSecMobileMeServerNotAvailable:
		return "ErrSecMobileMeServerNotAvailable"
	case ErrSecMobileMeServerServiceErr:
		return "ErrSecMobileMeServerServiceErr"
	case ErrSecModuleManagerInitializeFailed:
		return "ErrSecModuleManagerInitializeFailed"
	case ErrSecModuleManagerNotFound:
		return "ErrSecModuleManagerNotFound"
	case ErrSecModuleManifestVerifyFailed:
		return "ErrSecModuleManifestVerifyFailed"
	case ErrSecModuleNotLoaded:
		return "ErrSecModuleNotLoaded"
	case ErrSecMultiplePrivKeys:
		return "ErrSecMultiplePrivKeys"
	case ErrSecMultipleValuesUnsupported:
		return "ErrSecMultipleValuesUnsupported"
	case ErrSecNetworkFailure:
		return "ErrSecNetworkFailure"
	case ErrSecNoAccessForItem:
		return "ErrSecNoAccessForItem"
	case ErrSecNoBasicConstraints:
		return "ErrSecNoBasicConstraints"
	case ErrSecNoBasicConstraintsCA:
		return "ErrSecNoBasicConstraintsCA"
	case ErrSecNoCertificateModule:
		return "ErrSecNoCertificateModule"
	case ErrSecNoDefaultAuthority:
		return "ErrSecNoDefaultAuthority"
	case ErrSecNoDefaultKeychain:
		return "ErrSecNoDefaultKeychain"
	case ErrSecNoFieldValues:
		return "ErrSecNoFieldValues"
	case ErrSecNoPolicyModule:
		return "ErrSecNoPolicyModule"
	case ErrSecNoStorageModule:
		return "ErrSecNoStorageModule"
	case ErrSecNoSuchAttr:
		return "ErrSecNoSuchAttr"
	case ErrSecNoSuchClass:
		return "ErrSecNoSuchClass"
	case ErrSecNoSuchKeychain:
		return "ErrSecNoSuchKeychain"
	case ErrSecNoTrustSettings:
		return "ErrSecNoTrustSettings"
	case ErrSecNotAvailable:
		return "ErrSecNotAvailable"
	case ErrSecNotInitialized:
		return "ErrSecNotInitialized"
	case ErrSecNotLoggedIn:
		return "ErrSecNotLoggedIn"
	case ErrSecNotSigner:
		return "ErrSecNotSigner"
	case ErrSecNotTrusted:
		return "ErrSecNotTrusted"
	case ErrSecOCSPBadRequest:
		return "ErrSecOCSPBadRequest"
	case ErrSecOCSPBadResponse:
		return "ErrSecOCSPBadResponse"
	case ErrSecOCSPNoSigner:
		return "ErrSecOCSPNoSigner"
	case ErrSecOCSPNotTrustedToAnchor:
		return "ErrSecOCSPNotTrustedToAnchor"
	case ErrSecOCSPResponderInternalError:
		return "ErrSecOCSPResponderInternalError"
	case ErrSecOCSPResponderMalformedReq:
		return "ErrSecOCSPResponderMalformedReq"
	case ErrSecOCSPResponderSignatureRequired:
		return "ErrSecOCSPResponderSignatureRequired"
	case ErrSecOCSPResponderTryLater:
		return "ErrSecOCSPResponderTryLater"
	case ErrSecOCSPResponderUnauthorized:
		return "ErrSecOCSPResponderUnauthorized"
	case ErrSecOCSPResponseNonceMismatch:
		return "ErrSecOCSPResponseNonceMismatch"
	case ErrSecOCSPSignatureError:
		return "ErrSecOCSPSignatureError"
	case ErrSecOCSPStatusUnrecognized:
		return "ErrSecOCSPStatusUnrecognized"
	case ErrSecOCSPUnavailable:
		return "ErrSecOCSPUnavailable"
	case ErrSecOpWr:
		return "ErrSecOpWr"
	case ErrSecOutputLengthError:
		return "ErrSecOutputLengthError"
	case ErrSecPVCAlreadyConfigured:
		return "ErrSecPVCAlreadyConfigured"
	case ErrSecPVCReferentNotFound:
		return "ErrSecPVCReferentNotFound"
	case ErrSecParam:
		return "ErrSecParam"
	case ErrSecPassphraseRequired:
		return "ErrSecPassphraseRequired"
	case ErrSecPathLengthConstraintExceeded:
		return "ErrSecPathLengthConstraintExceeded"
	case ErrSecPkcs12VerifyFailure:
		return "ErrSecPkcs12VerifyFailure"
	case ErrSecPolicyNotFound:
		return "ErrSecPolicyNotFound"
	case ErrSecPrivilegeNotGranted:
		return "ErrSecPrivilegeNotGranted"
	case ErrSecPrivilegeNotSupported:
		return "ErrSecPrivilegeNotSupported"
	case ErrSecPublicKeyInconsistent:
		return "ErrSecPublicKeyInconsistent"
	case ErrSecQuerySizeUnknown:
		return "ErrSecQuerySizeUnknown"
	case ErrSecQuotaExceeded:
		return "ErrSecQuotaExceeded"
	case ErrSecReadOnly:
		return "ErrSecReadOnly"
	case ErrSecReadOnlyAttr:
		return "ErrSecReadOnlyAttr"
	case ErrSecRecordModified:
		return "ErrSecRecordModified"
	case ErrSecRejectedForm:
		return "ErrSecRejectedForm"
	case ErrSecRequestDescriptor:
		return "ErrSecRequestDescriptor"
	case ErrSecRequestLost:
		return "ErrSecRequestLost"
	case ErrSecRequestRejected:
		return "ErrSecRequestRejected"
	case ErrSecResourceSignBadCertChainLength:
		return "ErrSecResourceSignBadCertChainLength"
	case ErrSecResourceSignBadExtKeyUsage:
		return "ErrSecResourceSignBadExtKeyUsage"
	case ErrSecRestrictedAPI:
		return "ErrSecRestrictedAPI"
	case ErrSecSMIMEBadExtendedKeyUsage:
		return "ErrSecSMIMEBadExtendedKeyUsage"
	case ErrSecSMIMEBadKeyUsage:
		return "ErrSecSMIMEBadKeyUsage"
	case ErrSecSMIMEEmailAddressesNotFound:
		return "ErrSecSMIMEEmailAddressesNotFound"
	case ErrSecSMIMEKeyUsageNotCritical:
		return "ErrSecSMIMEKeyUsageNotCritical"
	case ErrSecSMIMENoEmailAddress:
		return "ErrSecSMIMENoEmailAddress"
	case ErrSecSMIMESubjAltNameNotCritical:
		return "ErrSecSMIMESubjAltNameNotCritical"
	case ErrSecSSLBadExtendedKeyUsage:
		return "ErrSecSSLBadExtendedKeyUsage"
	case ErrSecSelfCheckFailed:
		return "ErrSecSelfCheckFailed"
	case ErrSecServiceNotAvailable:
		return "ErrSecServiceNotAvailable"
	case ErrSecSigningTimeMissing:
		return "ErrSecSigningTimeMissing"
	case ErrSecStagedOperationInProgress:
		return "ErrSecStagedOperationInProgress"
	case ErrSecStagedOperationNotStarted:
		return "ErrSecStagedOperationNotStarted"
	case ErrSecSuccessValue:
		return "ErrSecSuccessValue"
	case ErrSecTagNotFound:
		return "ErrSecTagNotFound"
	case ErrSecTimestampAddInfoNotAvailable:
		return "ErrSecTimestampAddInfoNotAvailable"
	case ErrSecTimestampBadAlg:
		return "ErrSecTimestampBadAlg"
	case ErrSecTimestampBadDataFormat:
		return "ErrSecTimestampBadDataFormat"
	case ErrSecTimestampBadRequest:
		return "ErrSecTimestampBadRequest"
	case ErrSecTimestampInvalid:
		return "ErrSecTimestampInvalid"
	case ErrSecTimestampMissing:
		return "ErrSecTimestampMissing"
	case ErrSecTimestampNotTrusted:
		return "ErrSecTimestampNotTrusted"
	case ErrSecTimestampRejection:
		return "ErrSecTimestampRejection"
	case ErrSecTimestampRevocationNotification:
		return "ErrSecTimestampRevocationNotification"
	case ErrSecTimestampRevocationWarning:
		return "ErrSecTimestampRevocationWarning"
	case ErrSecTimestampServiceNotAvailable:
		return "ErrSecTimestampServiceNotAvailable"
	case ErrSecTimestampSystemFailure:
		return "ErrSecTimestampSystemFailure"
	case ErrSecTimestampTimeNotAvailable:
		return "ErrSecTimestampTimeNotAvailable"
	case ErrSecTimestampUnacceptedExtension:
		return "ErrSecTimestampUnacceptedExtension"
	case ErrSecTimestampUnacceptedPolicy:
		return "ErrSecTimestampUnacceptedPolicy"
	case ErrSecTimestampWaiting:
		return "ErrSecTimestampWaiting"
	case ErrSecTrustNotAvailable:
		return "ErrSecTrustNotAvailable"
	case ErrSecTrustSettingDeny:
		return "ErrSecTrustSettingDeny"
	case ErrSecUnimplemented:
		return "ErrSecUnimplemented"
	case ErrSecUnknownCRLExtension:
		return "ErrSecUnknownCRLExtension"
	case ErrSecUnknownCertExtension:
		return "ErrSecUnknownCertExtension"
	case ErrSecUnknownCriticalExtensionFlag:
		return "ErrSecUnknownCriticalExtensionFlag"
	case ErrSecUnknownFormat:
		return "ErrSecUnknownFormat"
	case ErrSecUnknownQualifiedCertStatement:
		return "ErrSecUnknownQualifiedCertStatement"
	case ErrSecUnknownTag:
		return "ErrSecUnknownTag"
	case ErrSecUnsupportedAddressType:
		return "ErrSecUnsupportedAddressType"
	case ErrSecUnsupportedFieldFormat:
		return "ErrSecUnsupportedFieldFormat"
	case ErrSecUnsupportedFormat:
		return "ErrSecUnsupportedFormat"
	case ErrSecUnsupportedIndexInfo:
		return "ErrSecUnsupportedIndexInfo"
	case ErrSecUnsupportedKeyAttributeMask:
		return "ErrSecUnsupportedKeyAttributeMask"
	case ErrSecUnsupportedKeyFormat:
		return "ErrSecUnsupportedKeyFormat"
	case ErrSecUnsupportedKeyLabel:
		return "ErrSecUnsupportedKeyLabel"
	case ErrSecUnsupportedKeySize:
		return "ErrSecUnsupportedKeySize"
	case ErrSecUnsupportedKeyUsageMask:
		return "ErrSecUnsupportedKeyUsageMask"
	case ErrSecUnsupportedLocality:
		return "ErrSecUnsupportedLocality"
	case ErrSecUnsupportedNumAttributes:
		return "ErrSecUnsupportedNumAttributes"
	case ErrSecUnsupportedNumIndexes:
		return "ErrSecUnsupportedNumIndexes"
	case ErrSecUnsupportedNumRecordTypes:
		return "ErrSecUnsupportedNumRecordTypes"
	case ErrSecUnsupportedNumSelectionPreds:
		return "ErrSecUnsupportedNumSelectionPreds"
	case ErrSecUnsupportedOperator:
		return "ErrSecUnsupportedOperator"
	case ErrSecUnsupportedQueryLimits:
		return "ErrSecUnsupportedQueryLimits"
	case ErrSecUnsupportedService:
		return "ErrSecUnsupportedService"
	case ErrSecUnsupportedVectorOfBuffers:
		return "ErrSecUnsupportedVectorOfBuffers"
	case ErrSecUserCanceled:
		return "ErrSecUserCanceled"
	case ErrSecVerificationFailure:
		return "ErrSecVerificationFailure"
	case ErrSecVerifyActionFailed:
		return "ErrSecVerifyActionFailed"
	case ErrSecVerifyFailed:
		return "ErrSecVerifyFailed"
	case ErrSecWrPerm:
		return "ErrSecWrPerm"
	case ErrSecWrongSecVersion:
		return "ErrSecWrongSecVersion"
	default:
		return fmt.Sprintf("ErrSecSuccess(%d)", e)
	}
}

type ErrSecureDownloadInvalid int32

const (
	ErrSecureDownloadInvalidDownload ErrSecureDownloadInvalid = -20053
	ErrSecureDownloadInvalidTicket   ErrSecureDownloadInvalid = -20052
)

func (e ErrSecureDownloadInvalid) String() string {
	switch e {
	case ErrSecureDownloadInvalidDownload:
		return "ErrSecureDownloadInvalidDownload"
	case ErrSecureDownloadInvalidTicket:
		return "ErrSecureDownloadInvalidTicket"
	default:
		return fmt.Sprintf("ErrSecureDownloadInvalid(%d)", e)
	}
}

type ErrSession int32

const (
	// ErrSessionAuthorizationDenied: Authorization denied.
	ErrSessionAuthorizationDenied ErrSession = -60502
	// ErrSessionInternal: An unrecognized internal error occurred.
	ErrSessionInternal ErrSession = -60008
	// ErrSessionInvalidAttributes: Detected an invalid set of request attribute bits.
	ErrSessionInvalidAttributes ErrSession = -60501
	// ErrSessionInvalidFlags: Encountered invalid flags or options.
	ErrSessionInvalidFlags ErrSession = -60011
	// ErrSessionInvalidId: Detected an invalid session ID.
	ErrSessionInvalidId ErrSession = -60500
	// ErrSessionSuccess: The operation completed successfully.
	ErrSessionSuccess ErrSession = 0
	// ErrSessionValueNotSet: The requested session attribute has not been set.
	ErrSessionValueNotSet ErrSession = -60503
)

func (e ErrSession) String() string {
	switch e {
	case ErrSessionAuthorizationDenied:
		return "ErrSessionAuthorizationDenied"
	case ErrSessionInternal:
		return "ErrSessionInternal"
	case ErrSessionInvalidAttributes:
		return "ErrSessionInvalidAttributes"
	case ErrSessionInvalidFlags:
		return "ErrSessionInvalidFlags"
	case ErrSessionInvalidId:
		return "ErrSessionInvalidId"
	case ErrSessionSuccess:
		return "ErrSessionSuccess"
	case ErrSessionValueNotSet:
		return "ErrSessionValueNotSet"
	default:
		return fmt.Sprintf("ErrSession(%d)", e)
	}
}

type KAuthorizationCallbacks uint32

const (
	KAuthorizationCallbacksVersion KAuthorizationCallbacks = 4
)

func (e KAuthorizationCallbacks) String() string {
	switch e {
	case KAuthorizationCallbacksVersion:
		return "KAuthorizationCallbacksVersion"
	default:
		return fmt.Sprintf("KAuthorizationCallbacks(%d)", e)
	}
}

type KAuthorizationFlagCanNotPre uint32

const (
	// KAuthorizationFlagCanNotPreAuthorize: Indicates the Security Server could not preauthorizethe right.
	KAuthorizationFlagCanNotPreAuthorize KAuthorizationFlagCanNotPre = 1
)

func (e KAuthorizationFlagCanNotPre) String() string {
	switch e {
	case KAuthorizationFlagCanNotPreAuthorize:
		return "KAuthorizationFlagCanNotPreAuthorize"
	default:
		return fmt.Sprintf("KAuthorizationFlagCanNotPre(%d)", e)
	}
}

type KAuthorizationPluginInterface uint32

const (
	KAuthorizationPluginInterfaceVersion KAuthorizationPluginInterface = 0
)

func (e KAuthorizationPluginInterface) String() string {
	switch e {
	case KAuthorizationPluginInterfaceVersion:
		return "KAuthorizationPluginInterfaceVersion"
	default:
		return fmt.Sprintf("KAuthorizationPluginInterface(%d)", e)
	}
}

type KSecCSCheckAllArchitectures uint32

const (
	KSecCSAllowNetworkAccess KSecCSCheckAllArchitectures = 65536
	// KSecCSBasicValidateOnly: Do not validate either the main executable or the bundle resources, if any.
	KSecCSBasicValidateOnly KSecCSCheckAllArchitectures = 6
	// KSecCSCheckAllArchitecturesValue: For multi-architecture (universal) Mach-O programs, validate all architectures included.
	KSecCSCheckAllArchitecturesValue   KSecCSCheckAllArchitectures = 1
	KSecCSCheckGatekeeperArchitectures KSecCSCheckAllArchitectures = 65
	// KSecCSCheckNestedCode: For code in bundle form, locate and recursively check embedded code.
	KSecCSCheckNestedCode KSecCSCheckAllArchitectures = 8
	// KSecCSDoNotValidateExecutable: Do not validate the contents of the main executable.
	KSecCSDoNotValidateExecutable KSecCSCheckAllArchitectures = 2
	// KSecCSDoNotValidateResources: Do not validate the presence and contents of all bundle resources (if any).
	KSecCSDoNotValidateResources   KSecCSCheckAllArchitectures = 4
	KSecCSFastExecutableValidation KSecCSCheckAllArchitectures = 131072
	KSecCSFullReport               KSecCSCheckAllArchitectures = 32
	KSecCSRestrictSidebandData     KSecCSCheckAllArchitectures = 512
	KSecCSRestrictSymlinks         KSecCSCheckAllArchitectures = 128
	KSecCSRestrictToAppLike        KSecCSCheckAllArchitectures = 256
	KSecCSSingleThreaded           KSecCSCheckAllArchitectures = 4096
	// KSecCSStrictValidate: Perform additional checks to ensure the validity of code in bundle form.
	KSecCSStrictValidate         KSecCSCheckAllArchitectures = 16
	KSecCSUseSoftwareSigningCert KSecCSCheckAllArchitectures = 1024
	KSecCSValidatePEH            KSecCSCheckAllArchitectures = 2048
)

func (e KSecCSCheckAllArchitectures) String() string {
	switch e {
	case KSecCSAllowNetworkAccess:
		return "KSecCSAllowNetworkAccess"
	case KSecCSBasicValidateOnly:
		return "KSecCSBasicValidateOnly"
	case KSecCSCheckAllArchitecturesValue:
		return "KSecCSCheckAllArchitecturesValue"
	case KSecCSCheckGatekeeperArchitectures:
		return "KSecCSCheckGatekeeperArchitectures"
	case KSecCSCheckNestedCode:
		return "KSecCSCheckNestedCode"
	case KSecCSDoNotValidateExecutable:
		return "KSecCSDoNotValidateExecutable"
	case KSecCSDoNotValidateResources:
		return "KSecCSDoNotValidateResources"
	case KSecCSFastExecutableValidation:
		return "KSecCSFastExecutableValidation"
	case KSecCSFullReport:
		return "KSecCSFullReport"
	case KSecCSRestrictSidebandData:
		return "KSecCSRestrictSidebandData"
	case KSecCSRestrictSymlinks:
		return "KSecCSRestrictSymlinks"
	case KSecCSRestrictToAppLike:
		return "KSecCSRestrictToAppLike"
	case KSecCSSingleThreaded:
		return "KSecCSSingleThreaded"
	case KSecCSStrictValidate:
		return "KSecCSStrictValidate"
	case KSecCSUseSoftwareSigningCert:
		return "KSecCSUseSoftwareSigningCert"
	case KSecCSValidatePEH:
		return "KSecCSValidatePEH"
	default:
		return fmt.Sprintf("KSecCSCheckAllArchitectures(%d)", e)
	}
}

type KSecCSDedicatedHost uint32

const (
	// KSecCSDedicatedHostValue: Declares dedicated hosting for the given host.
	KSecCSDedicatedHostValue KSecCSDedicatedHost = 1
	// KSecCSGenerateGuestHash: Ask the host to generate the unique binary identifier (kSecCodeInfoUnique) from the copy on disk at the path given.
	KSecCSGenerateGuestHash KSecCSDedicatedHost = 2
)

func (e KSecCSDedicatedHost) String() string {
	switch e {
	case KSecCSDedicatedHostValue:
		return "KSecCSDedicatedHostValue"
	case KSecCSGenerateGuestHash:
		return "KSecCSGenerateGuestHash"
	default:
		return fmt.Sprintf("KSecCSDedicatedHost(%d)", e)
	}
}

type KSecCSInternalInformation uint32

const (
	KSecCSCalculateCMSDigest KSecCSInternalInformation = 64
	// KSecCSContentInformation: More information about the file system contents making up the signed code on disk.
	KSecCSContentInformation KSecCSInternalInformation = 16
	// KSecCSDynamicInformation: Dynamic validity information about running code.
	KSecCSDynamicInformation KSecCSInternalInformation = 8
	// KSecCSInternalInformationValue: Internal code signing information.
	KSecCSInternalInformationValue KSecCSInternalInformation = 1
	// KSecCSRequirementInformation: Code requirements—including the designated requirement—embedded in the code.
	KSecCSRequirementInformation KSecCSInternalInformation = 4
	// KSecCSSigningInformation: Cryptographic signing information.
	KSecCSSigningInformation KSecCSInternalInformation = 2
	// KSecCSSkipResourceDirectory: Suppress validating the resource directory.
	KSecCSSkipResourceDirectory KSecCSInternalInformation = 32
)

func (e KSecCSInternalInformation) String() string {
	switch e {
	case KSecCSCalculateCMSDigest:
		return "KSecCSCalculateCMSDigest"
	case KSecCSContentInformation:
		return "KSecCSContentInformation"
	case KSecCSDynamicInformation:
		return "KSecCSDynamicInformation"
	case KSecCSInternalInformationValue:
		return "KSecCSInternalInformationValue"
	case KSecCSRequirementInformation:
		return "KSecCSRequirementInformation"
	case KSecCSSigningInformation:
		return "KSecCSSigningInformation"
	case KSecCSSkipResourceDirectory:
		return "KSecCSSkipResourceDirectory"
	default:
		return fmt.Sprintf("KSecCSInternalInformation(%d)", e)
	}
}

type KSecCSUseAll uint32

const (
	// KSecCSUseAllArchitectures: Flag for requesting all architectures.
	KSecCSUseAllArchitectures KSecCSUseAll = 1
)

func (e KSecCSUseAll) String() string {
	switch e {
	case KSecCSUseAllArchitectures:
		return "KSecCSUseAllArchitectures"
	default:
		return fmt.Sprintf("KSecCSUseAll(%d)", e)
	}
}

type KSecKeyKeyClass int32

const (
	// KSecKeyAlias: Type blob; currently unused.
	KSecKeyAlias KSecKeyKeyClass = 2
	// KSecKeyAlwaysSensitive: Type uint32; value is nonzero.
	KSecKeyAlwaysSensitive KSecKeyKeyClass = 15
	// KSecKeyApplicationTag: Type blob; currently unused.
	KSecKeyApplicationTag KSecKeyKeyClass = 7
	// KSecKeyDecrypt: Type uint32; value is nonzero.
	KSecKeyDecrypt KSecKeyKeyClass = 19
	// KSecKeyEffectiveKeySize: Type uint32; value is the effective number of bits in this key.
	KSecKeyEffectiveKeySize KSecKeyKeyClass = 11
	// KSecKeyEncrypt: Type uint32; value is nonzero.
	KSecKeyEncrypt KSecKeyKeyClass = 18
	// KSecKeyEndDate: Type `CSSM_DATE`.
	KSecKeyEndDate KSecKeyKeyClass = 13
	// KSecKeyExtractable: Type uint32; value is nonzero.
	KSecKeyExtractable KSecKeyKeyClass = 16
	// KSecKeyKeyClassValue: Type uint32 (`CSSM_KEYCLASS`); value is one of `CSSM_KEYCLASS_PUBLIC_KEY`, `CSSM_KEYCLASS_PRIVATE_KEY` or `CSSM_KEYCLASS_SESSION_KEY`.
	KSecKeyKeyClassValue KSecKeyKeyClass = 0
	// KSecKeyKeyCreator: Type data.
	KSecKeyKeyCreator KSecKeyKeyClass = 8
	// KSecKeyKeySizeInBits: Type uint32; value is the number of bits in this key.
	KSecKeyKeySizeInBits KSecKeyKeyClass = 10
	// KSecKeyKeyType: Type uint32; value is a CSSM algorithm (`CSSM_ALGORITHMS`) representing the algorithm associated with this key.
	KSecKeyKeyType KSecKeyKeyClass = 9
	// KSecKeyLabel: # Discussion
	KSecKeyLabel KSecKeyKeyClass = 6
	// KSecKeyModifiable: Type uint32; value is nonzero.
	KSecKeyModifiable KSecKeyKeyClass = 5
	// KSecKeyNeverExtractable: Type uint32; value is nonzero.
	KSecKeyNeverExtractable KSecKeyKeyClass = 17
	// KSecKeyPermanent: Type uint32; value is nonzero.
	KSecKeyPermanent KSecKeyKeyClass = 3
	// KSecKeyPrintName: Type blob; human readable name of the key.
	KSecKeyPrintName KSecKeyKeyClass = 1
	// KSecKeyPrivate: Type uint32; value is nonzero.
	KSecKeyPrivate KSecKeyKeyClass = 4
	// KSecKeySensitive: Type uint32; value is nonzero.
	KSecKeySensitive KSecKeyKeyClass = 14
	// KSecKeySign: Type uint32, value is nonzero.
	KSecKeySign KSecKeyKeyClass = 21
	// KSecKeySignRecover: Type uint32.
	KSecKeySignRecover KSecKeyKeyClass = 23
	// KSecKeyUnwrap: Type uint32; value is nonzero.
	KSecKeyUnwrap KSecKeyKeyClass = 26
	// KSecKeyVerify: Type uint32, value is nonzero.
	KSecKeyVerify KSecKeyKeyClass = 22
	// KSecKeyVerifyRecover: Type uint32.
	KSecKeyVerifyRecover KSecKeyKeyClass = 24
	// KSecKeyWrap: Type uint32; value is nonzero.
	KSecKeyWrap KSecKeyKeyClass = 25
	// Deprecated.
	KSecKeyDerive KSecKeyKeyClass = 20
	// Deprecated.
	KSecKeyStartDate KSecKeyKeyClass = 12
)

func (e KSecKeyKeyClass) String() string {
	switch e {
	case KSecKeyAlias:
		return "KSecKeyAlias"
	case KSecKeyAlwaysSensitive:
		return "KSecKeyAlwaysSensitive"
	case KSecKeyApplicationTag:
		return "KSecKeyApplicationTag"
	case KSecKeyDecrypt:
		return "KSecKeyDecrypt"
	case KSecKeyEffectiveKeySize:
		return "KSecKeyEffectiveKeySize"
	case KSecKeyEncrypt:
		return "KSecKeyEncrypt"
	case KSecKeyEndDate:
		return "KSecKeyEndDate"
	case KSecKeyExtractable:
		return "KSecKeyExtractable"
	case KSecKeyKeyClassValue:
		return "KSecKeyKeyClassValue"
	case KSecKeyKeyCreator:
		return "KSecKeyKeyCreator"
	case KSecKeyKeySizeInBits:
		return "KSecKeyKeySizeInBits"
	case KSecKeyKeyType:
		return "KSecKeyKeyType"
	case KSecKeyLabel:
		return "KSecKeyLabel"
	case KSecKeyModifiable:
		return "KSecKeyModifiable"
	case KSecKeyNeverExtractable:
		return "KSecKeyNeverExtractable"
	case KSecKeyPermanent:
		return "KSecKeyPermanent"
	case KSecKeyPrintName:
		return "KSecKeyPrintName"
	case KSecKeyPrivate:
		return "KSecKeyPrivate"
	case KSecKeySensitive:
		return "KSecKeySensitive"
	case KSecKeySign:
		return "KSecKeySign"
	case KSecKeySignRecover:
		return "KSecKeySignRecover"
	case KSecKeyUnwrap:
		return "KSecKeyUnwrap"
	case KSecKeyVerify:
		return "KSecKeyVerify"
	case KSecKeyVerifyRecover:
		return "KSecKeyVerifyRecover"
	case KSecKeyWrap:
		return "KSecKeyWrap"
	case KSecKeyDerive:
		return "KSecKeyDerive"
	case KSecKeyStartDate:
		return "KSecKeyStartDate"
	default:
		return fmt.Sprintf("KSecKeyKeyClass(%d)", e)
	}
}

type KSecNo uint32

const (
	// KSecNoGuest: Not a valid [SecGuestRef] object.
	KSecNoGuest KSecNo = 0
)

func (e KSecNo) String() string {
	switch e {
	case KSecNoGuest:
		return "KSecNoGuest"
	default:
		return fmt.Sprintf("KSecNo(%d)", e)
	}
}

type KSecRevocation uint

const (
	// KSecRevocationCRLMethod: Perform revocation checking using the CRL (Certification Revocation List) method.
	KSecRevocationCRLMethod KSecRevocation = 2
	// KSecRevocationNetworkAccessDisabled: Consult only locally cached replies; do not use network access.
	KSecRevocationNetworkAccessDisabled KSecRevocation = 16
	// KSecRevocationOCSPMethod: Perform revocation     checking using OCSP (Online Certificate Status Protocol).
	KSecRevocationOCSPMethod KSecRevocation = 1
	// KSecRevocationPreferCRL: Prefer CRL revocation checking over OCSP; by default, OCSP is preferred.
	KSecRevocationPreferCRL KSecRevocation = 4
	// KSecRevocationRequirePositiveResponse: Require a positive response to pass the policy.
	KSecRevocationRequirePositiveResponse KSecRevocation = 8
	// KSecRevocationUseAnyAvailableMethod: Perform either OCSP or CRL checking.
	KSecRevocationUseAnyAvailableMethod KSecRevocation = 3
)

func (e KSecRevocation) String() string {
	switch e {
	case KSecRevocationCRLMethod:
		return "KSecRevocationCRLMethod"
	case KSecRevocationNetworkAccessDisabled:
		return "KSecRevocationNetworkAccessDisabled"
	case KSecRevocationOCSPMethod:
		return "KSecRevocationOCSPMethod"
	case KSecRevocationPreferCRL:
		return "KSecRevocationPreferCRL"
	case KSecRevocationRequirePositiveResponse:
		return "KSecRevocationRequirePositiveResponse"
	case KSecRevocationUseAnyAvailableMethod:
		return "KSecRevocationUseAnyAvailableMethod"
	default:
		return fmt.Sprintf("KSecRevocation(%d)", e)
	}
}

type KSecSubjectItemAttr uint32

const (
	// KSecCertEncodingItemAttr: Certificate encoding.
	KSecCertEncodingItemAttr KSecSubjectItemAttr = 'c'<<24 | 'e'<<16 | 'n'<<8 | 'c' // 'cenc'
	// KSecCertTypeItemAttr: Certificate type.
	KSecCertTypeItemAttr KSecSubjectItemAttr = 'c'<<24 | 't'<<16 | 'y'<<8 | 'p' // 'ctyp'
	// KSecIssuerItemAttr: DER-encoded issuer distinguished name.
	KSecIssuerItemAttr KSecSubjectItemAttr = 'i'<<24 | 's'<<16 | 's'<<8 | 'u' // 'issu'
	// KSecPublicKeyHashItemAttr: Public key hash.
	KSecPublicKeyHashItemAttr KSecSubjectItemAttr = 'h'<<24 | 'p'<<16 | 'k'<<8 | 'y' // 'hpky'
	// KSecSerialNumberItemAttr: DER-encoded certificate serial number (without the tag and length).
	KSecSerialNumberItemAttr KSecSubjectItemAttr = 's'<<24 | 'n'<<16 | 'b'<<8 | 'r' // 'snbr'
	// KSecSubjectItemAttrValue: DER-encoded subject distinguished name.
	KSecSubjectItemAttrValue KSecSubjectItemAttr = 's'<<24 | 'u'<<16 | 'b'<<8 | 'j' // 'subj'
	// KSecSubjectKeyIdentifierItemAttr: Subject key identifier.
	KSecSubjectKeyIdentifierItemAttr KSecSubjectItemAttr = 's'<<24 | 'k'<<16 | 'i'<<8 | 'd' // 'skid'
)

func (e KSecSubjectItemAttr) String() string {
	switch e {
	case KSecCertEncodingItemAttr:
		return "KSecCertEncodingItemAttr"
	case KSecCertTypeItemAttr:
		return "KSecCertTypeItemAttr"
	case KSecIssuerItemAttr:
		return "KSecIssuerItemAttr"
	case KSecPublicKeyHashItemAttr:
		return "KSecPublicKeyHashItemAttr"
	case KSecSerialNumberItemAttr:
		return "KSecSerialNumberItemAttr"
	case KSecSubjectItemAttrValue:
		return "KSecSubjectItemAttrValue"
	case KSecSubjectKeyIdentifierItemAttr:
		return "KSecSubjectKeyIdentifierItemAttr"
	default:
		return fmt.Sprintf("KSecSubjectItemAttr(%d)", e)
	}
}

type KSecTransform int

const (
	// Deprecated.
	KSecTransformErrorAbortInProgress KSecTransform = 19
	// Deprecated.
	KSecTransformErrorAborted KSecTransform = 20
	// Deprecated.
	KSecTransformErrorAttributeNotFound KSecTransform = 1
	// Deprecated.
	KSecTransformErrorInvalidAlgorithm KSecTransform = 6
	// Deprecated.
	KSecTransformErrorInvalidConnection KSecTransform = 15
	// Deprecated.
	KSecTransformErrorInvalidInput KSecTransform = 10
	// Deprecated.
	KSecTransformErrorInvalidInputDictionary KSecTransform = 5
	// Deprecated.
	KSecTransformErrorInvalidLength KSecTransform = 7
	// Deprecated.
	KSecTransformErrorInvalidOperation KSecTransform = 2
	// Deprecated.
	KSecTransformErrorInvalidType KSecTransform = 8
	// Deprecated.
	KSecTransformErrorMissingParameter KSecTransform = 14
	// Deprecated.
	KSecTransformErrorMoreThanOneOutput KSecTransform = 4
	// Deprecated.
	KSecTransformErrorNameAlreadyRegistered KSecTransform = 11
	// Deprecated.
	KSecTransformErrorNotInitializedCorrectly KSecTransform = 3
	// Deprecated.
	KSecTransformErrorUnsupportedAttribute KSecTransform = 12
	// Deprecated.
	KSecTransformInvalidArgument KSecTransform = 21
	// Deprecated.
	KSecTransformInvalidOverride KSecTransform = 17
	// Deprecated.
	KSecTransformOperationNotSupportedOnGroup KSecTransform = 13
	// Deprecated.
	KSecTransformTransformIsExecuting KSecTransform = 16
	// Deprecated.
	KSecTransformTransformIsNotRegistered KSecTransform = 18
)

func (e KSecTransform) String() string {
	switch e {
	case KSecTransformErrorAbortInProgress:
		return "KSecTransformErrorAbortInProgress"
	case KSecTransformErrorAborted:
		return "KSecTransformErrorAborted"
	case KSecTransformErrorAttributeNotFound:
		return "KSecTransformErrorAttributeNotFound"
	case KSecTransformErrorInvalidAlgorithm:
		return "KSecTransformErrorInvalidAlgorithm"
	case KSecTransformErrorInvalidConnection:
		return "KSecTransformErrorInvalidConnection"
	case KSecTransformErrorInvalidInput:
		return "KSecTransformErrorInvalidInput"
	case KSecTransformErrorInvalidInputDictionary:
		return "KSecTransformErrorInvalidInputDictionary"
	case KSecTransformErrorInvalidLength:
		return "KSecTransformErrorInvalidLength"
	case KSecTransformErrorInvalidOperation:
		return "KSecTransformErrorInvalidOperation"
	case KSecTransformErrorInvalidType:
		return "KSecTransformErrorInvalidType"
	case KSecTransformErrorMissingParameter:
		return "KSecTransformErrorMissingParameter"
	case KSecTransformErrorMoreThanOneOutput:
		return "KSecTransformErrorMoreThanOneOutput"
	case KSecTransformErrorNameAlreadyRegistered:
		return "KSecTransformErrorNameAlreadyRegistered"
	case KSecTransformErrorNotInitializedCorrectly:
		return "KSecTransformErrorNotInitializedCorrectly"
	case KSecTransformErrorUnsupportedAttribute:
		return "KSecTransformErrorUnsupportedAttribute"
	case KSecTransformInvalidArgument:
		return "KSecTransformInvalidArgument"
	case KSecTransformInvalidOverride:
		return "KSecTransformInvalidOverride"
	case KSecTransformOperationNotSupportedOnGroup:
		return "KSecTransformOperationNotSupportedOnGroup"
	case KSecTransformTransformIsExecuting:
		return "KSecTransformTransformIsExecuting"
	case KSecTransformTransformIsNotRegistered:
		return "KSecTransformTransformIsNotRegistered"
	default:
		return fmt.Sprintf("KSecTransform(%d)", e)
	}
}

type KSecUnlockStateStatus uint32

const (
	// KSecReadPermStatus: Indicates the keychain is readable.
	KSecReadPermStatus KSecUnlockStateStatus = 2
	// KSecUnlockStateStatusValue: Indicates the keychain is unlocked.
	KSecUnlockStateStatusValue KSecUnlockStateStatus = 1
	// KSecWritePermStatus: Indicates the keychain is writable.
	KSecWritePermStatus KSecUnlockStateStatus = 4
)

func (e KSecUnlockStateStatus) String() string {
	switch e {
	case KSecReadPermStatus:
		return "KSecReadPermStatus"
	case KSecUnlockStateStatusValue:
		return "KSecUnlockStateStatusValue"
	case KSecWritePermStatus:
		return "KSecWritePermStatus"
	default:
		return fmt.Sprintf("KSecUnlockStateStatus(%d)", e)
	}
}

type KSecUseOnlyUID uint32

const (
	// KSecHonorRoot: The access control list should treat the root user as a typical user for ownership purposes.
	KSecHonorRoot KSecUseOnlyUID = 0x100
	// KSecMatchBits: The access control list should be owned by users whose ID matches the specified user ID or who are members of a group whose ID matches the specified group ID parameter.
	KSecMatchBits KSecUseOnlyUID = 3
	// KSecUseOnlyGID: The access control list should be owned by users that are members of a group matching the specified group ID parameter.
	KSecUseOnlyGID KSecUseOnlyUID = 2
	// KSecUseOnlyUIDValue: The access control list should be owned by the user matching the specified user ID parameter.
	KSecUseOnlyUIDValue KSecUseOnlyUID = 1
)

func (e KSecUseOnlyUID) String() string {
	switch e {
	case KSecHonorRoot:
		return "KSecHonorRoot"
	case KSecMatchBits:
		return "KSecMatchBits"
	case KSecUseOnlyGID:
		return "KSecUseOnlyGID"
	case KSecUseOnlyUIDValue:
		return "KSecUseOnlyUIDValue"
	default:
		return fmt.Sprintf("KSecUseOnlyUID(%d)", e)
	}
}

type NoSecuritySession uint32

const (
	// CallerSecuritySession: A value that is a placeholder for the caller’s session.
	CallerSecuritySession NoSecuritySession = 4294967295
	// NoSecuritySessionValue: Not a valid session.
	NoSecuritySessionValue NoSecuritySession = 0
)

func (e NoSecuritySession) String() string {
	switch e {
	case CallerSecuritySession:
		return "CallerSecuritySession"
	case NoSecuritySessionValue:
		return "NoSecuritySessionValue"
	default:
		return fmt.Sprintf("NoSecuritySession(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SSLAuthenticate
type SSLAuthenticate int32

const (
	// KAlwaysAuthenticate: Indicates that client-side authentication is required.
	KAlwaysAuthenticate SSLAuthenticate = 1
	// KNeverAuthenticate: Indicates that client-side authentication is not required.
	KNeverAuthenticate SSLAuthenticate = 0
	// KTryAuthenticate: Indicates that client-side authentication should be attempted.
	KTryAuthenticate SSLAuthenticate = 2
)

func (e SSLAuthenticate) String() string {
	switch e {
	case KAlwaysAuthenticate:
		return "KAlwaysAuthenticate"
	case KNeverAuthenticate:
		return "KNeverAuthenticate"
	case KTryAuthenticate:
		return "KTryAuthenticate"
	default:
		return fmt.Sprintf("SSLAuthenticate(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SSLCiphersuiteGroup
type SSLCiphersuiteGroup int32

const (
	KSSLCiphersuiteGroupATS              SSLCiphersuiteGroup = 3
	KSSLCiphersuiteGroupATSCompatibility SSLCiphersuiteGroup = 4
	KSSLCiphersuiteGroupATSFCP_v2_1      SSLCiphersuiteGroup = 5
	KSSLCiphersuiteGroupCompatibility    SSLCiphersuiteGroup = 1
	KSSLCiphersuiteGroupDefault          SSLCiphersuiteGroup = 0
	KSSLCiphersuiteGroupLegacy           SSLCiphersuiteGroup = 2
)

func (e SSLCiphersuiteGroup) String() string {
	switch e {
	case KSSLCiphersuiteGroupATS:
		return "KSSLCiphersuiteGroupATS"
	case KSSLCiphersuiteGroupATSCompatibility:
		return "KSSLCiphersuiteGroupATSCompatibility"
	case KSSLCiphersuiteGroupATSFCP_v2_1:
		return "KSSLCiphersuiteGroupATSFCP_v2_1"
	case KSSLCiphersuiteGroupCompatibility:
		return "KSSLCiphersuiteGroupCompatibility"
	case KSSLCiphersuiteGroupDefault:
		return "KSSLCiphersuiteGroupDefault"
	case KSSLCiphersuiteGroupLegacy:
		return "KSSLCiphersuiteGroupLegacy"
	default:
		return fmt.Sprintf("SSLCiphersuiteGroup(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SSLClientCertificateState
type SSLClientCertificateState int32

const (
	// Deprecated.
	KSSLClientCertNone SSLClientCertificateState = 0
	// Deprecated.
	KSSLClientCertRejected SSLClientCertificateState = 3
	// Deprecated.
	KSSLClientCertRequested SSLClientCertificateState = 1
	// Deprecated.
	KSSLClientCertSent SSLClientCertificateState = 2
)

func (e SSLClientCertificateState) String() string {
	switch e {
	case KSSLClientCertNone:
		return "KSSLClientCertNone"
	case KSSLClientCertRejected:
		return "KSSLClientCertRejected"
	case KSSLClientCertRequested:
		return "KSSLClientCertRequested"
	case KSSLClientCertSent:
		return "KSSLClientCertSent"
	default:
		return fmt.Sprintf("SSLClientCertificateState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SSLConnectionType
type SSLConnectionType int32

const (
	// Deprecated.
	KSSLDatagramType SSLConnectionType = 1
	// Deprecated.
	KSSLStreamType SSLConnectionType = 0
)

func (e SSLConnectionType) String() string {
	switch e {
	case KSSLDatagramType:
		return "KSSLDatagramType"
	case KSSLStreamType:
		return "KSSLStreamType"
	default:
		return fmt.Sprintf("SSLConnectionType(%d)", e)
	}
}

type SSLNullWithNullNull uint16

const (
	// SSL_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA: Session key size conforms to pre-1998 US export restrictions.
	SSL_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA SSLNullWithNullNull = 0x11
	SSL_DHE_DSS_WITH_3DES_EDE_CBC_SHA     SSLNullWithNullNull = 0x13
	SSL_DHE_DSS_WITH_DES_CBC_SHA          SSLNullWithNullNull = 0x12
	// SSL_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA: Session key size conforms to pre-1998 US export restrictions.
	SSL_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA SSLNullWithNullNull = 0x14
	SSL_DHE_RSA_WITH_3DES_EDE_CBC_SHA     SSLNullWithNullNull = 0x16
	SSL_DHE_RSA_WITH_DES_CBC_SHA          SSLNullWithNullNull = 0x15
	SSL_DH_DSS_EXPORT_WITH_DES40_CBC_SHA  SSLNullWithNullNull = 0xb
	SSL_DH_DSS_WITH_3DES_EDE_CBC_SHA      SSLNullWithNullNull = 0xd
	SSL_DH_DSS_WITH_DES_CBC_SHA           SSLNullWithNullNull = 0xc
	// SSL_DH_RSA_EXPORT_WITH_DES40_CBC_SHA: Session key size conforms to pre-1998 US export restrictions.
	SSL_DH_RSA_EXPORT_WITH_DES40_CBC_SHA SSLNullWithNullNull = 0xe
	SSL_DH_RSA_WITH_3DES_EDE_CBC_SHA     SSLNullWithNullNull = 0x10
	SSL_DH_RSA_WITH_DES_CBC_SHA          SSLNullWithNullNull = 0xf
	// SSL_DH_anon_EXPORT_WITH_DES40_CBC_SHA: Session key size conforms to pre-1998 US export restrictions.
	SSL_DH_anon_EXPORT_WITH_DES40_CBC_SHA SSLNullWithNullNull = 0x19
	// SSL_DH_anon_EXPORT_WITH_RC4_40_MD5: Session key size conforms to pre-1998 US export restrictions.
	SSL_DH_anon_EXPORT_WITH_RC4_40_MD5     SSLNullWithNullNull = 0x17
	SSL_DH_anon_WITH_3DES_EDE_CBC_SHA      SSLNullWithNullNull = 0x1b
	SSL_DH_anon_WITH_DES_CBC_SHA           SSLNullWithNullNull = 0x1a
	SSL_DH_anon_WITH_RC4_128_MD5           SSLNullWithNullNull = 0x18
	SSL_FORTEZZA_DMS_WITH_FORTEZZA_CBC_SHA SSLNullWithNullNull = 0x1d
	SSL_FORTEZZA_DMS_WITH_NULL_SHA         SSLNullWithNullNull = 0x1c
	SSL_NO_SUCH_CIPHERSUITE                SSLNullWithNullNull = 0xffff
	SSL_NULL_WITH_NULL_NULL                SSLNullWithNullNull = 0
	// SSL_RSA_EXPORT_WITH_DES40_CBC_SHA: Session key size conforms to pre-1998 US export restrictions.
	SSL_RSA_EXPORT_WITH_DES40_CBC_SHA SSLNullWithNullNull = 0x8
	// SSL_RSA_EXPORT_WITH_RC2_CBC_40_MD5: Session key size conforms to pre-1998 US export restrictions.
	SSL_RSA_EXPORT_WITH_RC2_CBC_40_MD5 SSLNullWithNullNull = 0x6
	// SSL_RSA_EXPORT_WITH_RC4_40_MD5: Session key size conforms to pre-1998 US export restrictions.
	SSL_RSA_EXPORT_WITH_RC4_40_MD5 SSLNullWithNullNull = 0x3
	// SSL_RSA_WITH_3DES_EDE_CBC_MD5: This value can be specified for SSL 2 but not SSL 3.
	SSL_RSA_WITH_3DES_EDE_CBC_MD5 SSLNullWithNullNull = 0xff83
	SSL_RSA_WITH_3DES_EDE_CBC_SHA SSLNullWithNullNull = 0xa
	// SSL_RSA_WITH_DES_CBC_MD5: This value can be specified for SSL 2 but not SSL 3.
	SSL_RSA_WITH_DES_CBC_MD5 SSLNullWithNullNull = 0xff82
	SSL_RSA_WITH_DES_CBC_SHA SSLNullWithNullNull = 0x9
	// SSL_RSA_WITH_IDEA_CBC_MD5: This value can be specified for SSL 2 but not SSL 3.
	SSL_RSA_WITH_IDEA_CBC_MD5 SSLNullWithNullNull = 0xff81
	SSL_RSA_WITH_IDEA_CBC_SHA SSLNullWithNullNull = 0x7
	SSL_RSA_WITH_NULL_MD5     SSLNullWithNullNull = 0x1
	SSL_RSA_WITH_NULL_SHA     SSLNullWithNullNull = 0x2
	// SSL_RSA_WITH_RC2_CBC_MD5: This value can be specified for SSL 2 but not SSL 3.
	SSL_RSA_WITH_RC2_CBC_MD5                      SSLNullWithNullNull = 0xff80
	SSL_RSA_WITH_RC4_128_MD5                      SSLNullWithNullNull = 0x4
	SSL_RSA_WITH_RC4_128_SHA                      SSLNullWithNullNull = 0x5
	TLS_AES_128_CCM_8_SHA256                      SSLNullWithNullNull = 0x1305
	TLS_AES_128_CCM_SHA256                        SSLNullWithNullNull = 0x1304
	TLS_AES_128_GCM_SHA256                        SSLNullWithNullNull = 0x1301
	TLS_AES_256_GCM_SHA384                        SSLNullWithNullNull = 0x1302
	TLS_CHACHA20_POLY1305_SHA256                  SSLNullWithNullNull = 0x1303
	TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA             SSLNullWithNullNull = 0x13
	TLS_DHE_DSS_WITH_AES_128_CBC_SHA              SSLNullWithNullNull = 0x32
	TLS_DHE_DSS_WITH_AES_128_CBC_SHA256           SSLNullWithNullNull = 0x40
	TLS_DHE_DSS_WITH_AES_128_GCM_SHA256           SSLNullWithNullNull = 0xa2
	TLS_DHE_DSS_WITH_AES_256_CBC_SHA              SSLNullWithNullNull = 0x38
	TLS_DHE_DSS_WITH_AES_256_CBC_SHA256           SSLNullWithNullNull = 0x6a
	TLS_DHE_DSS_WITH_AES_256_GCM_SHA384           SSLNullWithNullNull = 0xa3
	TLS_DHE_PSK_WITH_3DES_EDE_CBC_SHA             SSLNullWithNullNull = 0x8f
	TLS_DHE_PSK_WITH_AES_128_CBC_SHA              SSLNullWithNullNull = 0x90
	TLS_DHE_PSK_WITH_AES_128_CBC_SHA256           SSLNullWithNullNull = 0xb2
	TLS_DHE_PSK_WITH_AES_128_GCM_SHA256           SSLNullWithNullNull = 0xaa
	TLS_DHE_PSK_WITH_AES_256_CBC_SHA              SSLNullWithNullNull = 0x91
	TLS_DHE_PSK_WITH_AES_256_CBC_SHA384           SSLNullWithNullNull = 0xb3
	TLS_DHE_PSK_WITH_AES_256_GCM_SHA384           SSLNullWithNullNull = 0xab
	TLS_DHE_PSK_WITH_NULL_SHA                     SSLNullWithNullNull = 0x2d
	TLS_DHE_PSK_WITH_NULL_SHA256                  SSLNullWithNullNull = 0xb4
	TLS_DHE_PSK_WITH_NULL_SHA384                  SSLNullWithNullNull = 0xb5
	TLS_DHE_PSK_WITH_RC4_128_SHA                  SSLNullWithNullNull = 0x8e
	TLS_DHE_RSA_WITH_3DES_EDE_CBC_SHA             SSLNullWithNullNull = 0x16
	TLS_DHE_RSA_WITH_AES_128_CBC_SHA              SSLNullWithNullNull = 0x33
	TLS_DHE_RSA_WITH_AES_128_CBC_SHA256           SSLNullWithNullNull = 0x67
	TLS_DHE_RSA_WITH_AES_128_GCM_SHA256           SSLNullWithNullNull = 0x9e
	TLS_DHE_RSA_WITH_AES_256_CBC_SHA              SSLNullWithNullNull = 0x39
	TLS_DHE_RSA_WITH_AES_256_CBC_SHA256           SSLNullWithNullNull = 0x6b
	TLS_DHE_RSA_WITH_AES_256_GCM_SHA384           SSLNullWithNullNull = 0x9f
	TLS_DH_DSS_WITH_3DES_EDE_CBC_SHA              SSLNullWithNullNull = 0xd
	TLS_DH_DSS_WITH_AES_128_CBC_SHA               SSLNullWithNullNull = 0x30
	TLS_DH_DSS_WITH_AES_128_CBC_SHA256            SSLNullWithNullNull = 0x3e
	TLS_DH_DSS_WITH_AES_128_GCM_SHA256            SSLNullWithNullNull = 0xa4
	TLS_DH_DSS_WITH_AES_256_CBC_SHA               SSLNullWithNullNull = 0x36
	TLS_DH_DSS_WITH_AES_256_CBC_SHA256            SSLNullWithNullNull = 0x68
	TLS_DH_DSS_WITH_AES_256_GCM_SHA384            SSLNullWithNullNull = 0xa5
	TLS_DH_RSA_WITH_3DES_EDE_CBC_SHA              SSLNullWithNullNull = 0x10
	TLS_DH_RSA_WITH_AES_128_CBC_SHA               SSLNullWithNullNull = 0x31
	TLS_DH_RSA_WITH_AES_128_CBC_SHA256            SSLNullWithNullNull = 0x3f
	TLS_DH_RSA_WITH_AES_128_GCM_SHA256            SSLNullWithNullNull = 0xa0
	TLS_DH_RSA_WITH_AES_256_CBC_SHA               SSLNullWithNullNull = 0x37
	TLS_DH_RSA_WITH_AES_256_CBC_SHA256            SSLNullWithNullNull = 0x69
	TLS_DH_RSA_WITH_AES_256_GCM_SHA384            SSLNullWithNullNull = 0xa1
	TLS_DH_anon_WITH_3DES_EDE_CBC_SHA             SSLNullWithNullNull = 0x1b
	TLS_DH_anon_WITH_AES_128_CBC_SHA              SSLNullWithNullNull = 0x34
	TLS_DH_anon_WITH_AES_128_CBC_SHA256           SSLNullWithNullNull = 0x6c
	TLS_DH_anon_WITH_AES_128_GCM_SHA256           SSLNullWithNullNull = 0xa6
	TLS_DH_anon_WITH_AES_256_CBC_SHA              SSLNullWithNullNull = 0x3a
	TLS_DH_anon_WITH_AES_256_CBC_SHA256           SSLNullWithNullNull = 0x6d
	TLS_DH_anon_WITH_AES_256_GCM_SHA384           SSLNullWithNullNull = 0xa7
	TLS_DH_anon_WITH_RC4_128_MD5                  SSLNullWithNullNull = 0x18
	TLS_ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA         SSLNullWithNullNull = 0xc008
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA          SSLNullWithNullNull = 0xc009
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256       SSLNullWithNullNull = 0xc023
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256       SSLNullWithNullNull = 0xc02b
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA          SSLNullWithNullNull = 0xc00a
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384       SSLNullWithNullNull = 0xc024
	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384       SSLNullWithNullNull = 0xc02c
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 SSLNullWithNullNull = 0xcca9
	TLS_ECDHE_ECDSA_WITH_NULL_SHA                 SSLNullWithNullNull = 0xc006
	TLS_ECDHE_ECDSA_WITH_RC4_128_SHA              SSLNullWithNullNull = 0xc007
	TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA            SSLNullWithNullNull = 0xc035
	TLS_ECDHE_PSK_WITH_AES_256_CBC_SHA            SSLNullWithNullNull = 0xc036
	TLS_ECDHE_PSK_WITH_CHACHA20_POLY1305_SHA256   SSLNullWithNullNull = 0xccac
	TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA           SSLNullWithNullNull = 0xc012
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA            SSLNullWithNullNull = 0xc013
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256         SSLNullWithNullNull = 0xc027
	TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256         SSLNullWithNullNull = 0xc02f
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA            SSLNullWithNullNull = 0xc014
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384         SSLNullWithNullNull = 0xc028
	TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384         SSLNullWithNullNull = 0xc030
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256   SSLNullWithNullNull = 0xcca8
	TLS_ECDHE_RSA_WITH_NULL_SHA                   SSLNullWithNullNull = 0xc010
	TLS_ECDHE_RSA_WITH_RC4_128_SHA                SSLNullWithNullNull = 0xc011
	TLS_ECDH_ECDSA_WITH_3DES_EDE_CBC_SHA          SSLNullWithNullNull = 0xc003
	TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA           SSLNullWithNullNull = 0xc004
	TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA256        SSLNullWithNullNull = 0xc025
	TLS_ECDH_ECDSA_WITH_AES_128_GCM_SHA256        SSLNullWithNullNull = 0xc02d
	TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA           SSLNullWithNullNull = 0xc005
	TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA384        SSLNullWithNullNull = 0xc026
	TLS_ECDH_ECDSA_WITH_AES_256_GCM_SHA384        SSLNullWithNullNull = 0xc02e
	TLS_ECDH_ECDSA_WITH_NULL_SHA                  SSLNullWithNullNull = 0xc001
	TLS_ECDH_ECDSA_WITH_RC4_128_SHA               SSLNullWithNullNull = 0xc002
	TLS_ECDH_RSA_WITH_3DES_EDE_CBC_SHA            SSLNullWithNullNull = 0xc00d
	TLS_ECDH_RSA_WITH_AES_128_CBC_SHA             SSLNullWithNullNull = 0xc00e
	TLS_ECDH_RSA_WITH_AES_128_CBC_SHA256          SSLNullWithNullNull = 0xc029
	TLS_ECDH_RSA_WITH_AES_128_GCM_SHA256          SSLNullWithNullNull = 0xc031
	TLS_ECDH_RSA_WITH_AES_256_CBC_SHA             SSLNullWithNullNull = 0xc00f
	TLS_ECDH_RSA_WITH_AES_256_CBC_SHA384          SSLNullWithNullNull = 0xc02a
	TLS_ECDH_RSA_WITH_AES_256_GCM_SHA384          SSLNullWithNullNull = 0xc032
	TLS_ECDH_RSA_WITH_NULL_SHA                    SSLNullWithNullNull = 0xc00b
	TLS_ECDH_RSA_WITH_RC4_128_SHA                 SSLNullWithNullNull = 0xc00c
	TLS_ECDH_anon_WITH_3DES_EDE_CBC_SHA           SSLNullWithNullNull = 0xc017
	TLS_ECDH_anon_WITH_AES_128_CBC_SHA            SSLNullWithNullNull = 0xc018
	TLS_ECDH_anon_WITH_AES_256_CBC_SHA            SSLNullWithNullNull = 0xc019
	TLS_ECDH_anon_WITH_NULL_SHA                   SSLNullWithNullNull = 0xc015
	TLS_ECDH_anon_WITH_RC4_128_SHA                SSLNullWithNullNull = 0xc016
	TLS_EMPTY_RENEGOTIATION_INFO_SCSV             SSLNullWithNullNull = 0xff
	TLS_NULL_WITH_NULL_NULL                       SSLNullWithNullNull = 0
	TLS_PSK_WITH_3DES_EDE_CBC_SHA                 SSLNullWithNullNull = 0x8b
	TLS_PSK_WITH_AES_128_CBC_SHA                  SSLNullWithNullNull = 0x8c
	TLS_PSK_WITH_AES_128_CBC_SHA256               SSLNullWithNullNull = 0xae
	TLS_PSK_WITH_AES_128_GCM_SHA256               SSLNullWithNullNull = 0xa8
	TLS_PSK_WITH_AES_256_CBC_SHA                  SSLNullWithNullNull = 0x8d
	TLS_PSK_WITH_AES_256_CBC_SHA384               SSLNullWithNullNull = 0xaf
	TLS_PSK_WITH_AES_256_GCM_SHA384               SSLNullWithNullNull = 0xa9
	TLS_PSK_WITH_CHACHA20_POLY1305_SHA256         SSLNullWithNullNull = 0xccab
	TLS_PSK_WITH_NULL_SHA                         SSLNullWithNullNull = 0x2c
	TLS_PSK_WITH_NULL_SHA256                      SSLNullWithNullNull = 0xb0
	TLS_PSK_WITH_NULL_SHA384                      SSLNullWithNullNull = 0xb1
	TLS_PSK_WITH_RC4_128_SHA                      SSLNullWithNullNull = 0x8a
	TLS_RSA_PSK_WITH_3DES_EDE_CBC_SHA             SSLNullWithNullNull = 0x93
	TLS_RSA_PSK_WITH_AES_128_CBC_SHA              SSLNullWithNullNull = 0x94
	TLS_RSA_PSK_WITH_AES_128_CBC_SHA256           SSLNullWithNullNull = 0xb6
	TLS_RSA_PSK_WITH_AES_128_GCM_SHA256           SSLNullWithNullNull = 0xac
	TLS_RSA_PSK_WITH_AES_256_CBC_SHA              SSLNullWithNullNull = 0x95
	TLS_RSA_PSK_WITH_AES_256_CBC_SHA384           SSLNullWithNullNull = 0xb7
	TLS_RSA_PSK_WITH_AES_256_GCM_SHA384           SSLNullWithNullNull = 0xad
	TLS_RSA_PSK_WITH_NULL_SHA                     SSLNullWithNullNull = 0x2e
	TLS_RSA_PSK_WITH_NULL_SHA256                  SSLNullWithNullNull = 0xb8
	TLS_RSA_PSK_WITH_NULL_SHA384                  SSLNullWithNullNull = 0xb9
	TLS_RSA_PSK_WITH_RC4_128_SHA                  SSLNullWithNullNull = 0x92
	TLS_RSA_WITH_3DES_EDE_CBC_SHA                 SSLNullWithNullNull = 0xa
	TLS_RSA_WITH_AES_128_CBC_SHA                  SSLNullWithNullNull = 0x2f
	TLS_RSA_WITH_AES_128_CBC_SHA256               SSLNullWithNullNull = 0x3c
	TLS_RSA_WITH_AES_128_GCM_SHA256               SSLNullWithNullNull = 0x9c
	TLS_RSA_WITH_AES_256_CBC_SHA                  SSLNullWithNullNull = 0x35
	TLS_RSA_WITH_AES_256_CBC_SHA256               SSLNullWithNullNull = 0x3d
	TLS_RSA_WITH_AES_256_GCM_SHA384               SSLNullWithNullNull = 0x9d
	TLS_RSA_WITH_NULL_MD5                         SSLNullWithNullNull = 0x1
	TLS_RSA_WITH_NULL_SHA                         SSLNullWithNullNull = 0x2
	TLS_RSA_WITH_NULL_SHA256                      SSLNullWithNullNull = 0x3b
	TLS_RSA_WITH_RC4_128_MD5                      SSLNullWithNullNull = 0x4
	TLS_RSA_WITH_RC4_128_SHA                      SSLNullWithNullNull = 0x5
)

func (e SSLNullWithNullNull) String() string {
	switch e {
	case SSL_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA:
		return "SSL_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA"
	case SSL_DHE_DSS_WITH_3DES_EDE_CBC_SHA:
		return "SSL_DHE_DSS_WITH_3DES_EDE_CBC_SHA"
	case SSL_DHE_DSS_WITH_DES_CBC_SHA:
		return "SSL_DHE_DSS_WITH_DES_CBC_SHA"
	case SSL_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA:
		return "SSL_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA"
	case SSL_DHE_RSA_WITH_3DES_EDE_CBC_SHA:
		return "SSL_DHE_RSA_WITH_3DES_EDE_CBC_SHA"
	case SSL_DHE_RSA_WITH_DES_CBC_SHA:
		return "SSL_DHE_RSA_WITH_DES_CBC_SHA"
	case SSL_DH_DSS_EXPORT_WITH_DES40_CBC_SHA:
		return "SSL_DH_DSS_EXPORT_WITH_DES40_CBC_SHA"
	case SSL_DH_DSS_WITH_3DES_EDE_CBC_SHA:
		return "SSL_DH_DSS_WITH_3DES_EDE_CBC_SHA"
	case SSL_DH_DSS_WITH_DES_CBC_SHA:
		return "SSL_DH_DSS_WITH_DES_CBC_SHA"
	case SSL_DH_RSA_EXPORT_WITH_DES40_CBC_SHA:
		return "SSL_DH_RSA_EXPORT_WITH_DES40_CBC_SHA"
	case SSL_DH_RSA_WITH_3DES_EDE_CBC_SHA:
		return "SSL_DH_RSA_WITH_3DES_EDE_CBC_SHA"
	case SSL_DH_RSA_WITH_DES_CBC_SHA:
		return "SSL_DH_RSA_WITH_DES_CBC_SHA"
	case SSL_DH_anon_EXPORT_WITH_DES40_CBC_SHA:
		return "SSL_DH_anon_EXPORT_WITH_DES40_CBC_SHA"
	case SSL_DH_anon_EXPORT_WITH_RC4_40_MD5:
		return "SSL_DH_anon_EXPORT_WITH_RC4_40_MD5"
	case SSL_DH_anon_WITH_3DES_EDE_CBC_SHA:
		return "SSL_DH_anon_WITH_3DES_EDE_CBC_SHA"
	case SSL_DH_anon_WITH_DES_CBC_SHA:
		return "SSL_DH_anon_WITH_DES_CBC_SHA"
	case SSL_DH_anon_WITH_RC4_128_MD5:
		return "SSL_DH_anon_WITH_RC4_128_MD5"
	case SSL_FORTEZZA_DMS_WITH_FORTEZZA_CBC_SHA:
		return "SSL_FORTEZZA_DMS_WITH_FORTEZZA_CBC_SHA"
	case SSL_FORTEZZA_DMS_WITH_NULL_SHA:
		return "SSL_FORTEZZA_DMS_WITH_NULL_SHA"
	case SSL_NO_SUCH_CIPHERSUITE:
		return "SSL_NO_SUCH_CIPHERSUITE"
	case SSL_NULL_WITH_NULL_NULL:
		return "SSL_NULL_WITH_NULL_NULL"
	case SSL_RSA_EXPORT_WITH_DES40_CBC_SHA:
		return "SSL_RSA_EXPORT_WITH_DES40_CBC_SHA"
	case SSL_RSA_EXPORT_WITH_RC2_CBC_40_MD5:
		return "SSL_RSA_EXPORT_WITH_RC2_CBC_40_MD5"
	case SSL_RSA_EXPORT_WITH_RC4_40_MD5:
		return "SSL_RSA_EXPORT_WITH_RC4_40_MD5"
	case SSL_RSA_WITH_3DES_EDE_CBC_MD5:
		return "SSL_RSA_WITH_3DES_EDE_CBC_MD5"
	case SSL_RSA_WITH_3DES_EDE_CBC_SHA:
		return "SSL_RSA_WITH_3DES_EDE_CBC_SHA"
	case SSL_RSA_WITH_DES_CBC_MD5:
		return "SSL_RSA_WITH_DES_CBC_MD5"
	case SSL_RSA_WITH_DES_CBC_SHA:
		return "SSL_RSA_WITH_DES_CBC_SHA"
	case SSL_RSA_WITH_IDEA_CBC_MD5:
		return "SSL_RSA_WITH_IDEA_CBC_MD5"
	case SSL_RSA_WITH_IDEA_CBC_SHA:
		return "SSL_RSA_WITH_IDEA_CBC_SHA"
	case SSL_RSA_WITH_NULL_MD5:
		return "SSL_RSA_WITH_NULL_MD5"
	case SSL_RSA_WITH_NULL_SHA:
		return "SSL_RSA_WITH_NULL_SHA"
	case SSL_RSA_WITH_RC2_CBC_MD5:
		return "SSL_RSA_WITH_RC2_CBC_MD5"
	case SSL_RSA_WITH_RC4_128_MD5:
		return "SSL_RSA_WITH_RC4_128_MD5"
	case SSL_RSA_WITH_RC4_128_SHA:
		return "SSL_RSA_WITH_RC4_128_SHA"
	case TLS_AES_128_CCM_8_SHA256:
		return "TLS_AES_128_CCM_8_SHA256"
	case TLS_AES_128_CCM_SHA256:
		return "TLS_AES_128_CCM_SHA256"
	case TLS_AES_128_GCM_SHA256:
		return "TLS_AES_128_GCM_SHA256"
	case TLS_AES_256_GCM_SHA384:
		return "TLS_AES_256_GCM_SHA384"
	case TLS_CHACHA20_POLY1305_SHA256:
		return "TLS_CHACHA20_POLY1305_SHA256"
	case TLS_DHE_DSS_WITH_AES_128_CBC_SHA:
		return "TLS_DHE_DSS_WITH_AES_128_CBC_SHA"
	case TLS_DHE_DSS_WITH_AES_128_CBC_SHA256:
		return "TLS_DHE_DSS_WITH_AES_128_CBC_SHA256"
	case TLS_DHE_DSS_WITH_AES_128_GCM_SHA256:
		return "TLS_DHE_DSS_WITH_AES_128_GCM_SHA256"
	case TLS_DHE_DSS_WITH_AES_256_CBC_SHA:
		return "TLS_DHE_DSS_WITH_AES_256_CBC_SHA"
	case TLS_DHE_DSS_WITH_AES_256_CBC_SHA256:
		return "TLS_DHE_DSS_WITH_AES_256_CBC_SHA256"
	case TLS_DHE_DSS_WITH_AES_256_GCM_SHA384:
		return "TLS_DHE_DSS_WITH_AES_256_GCM_SHA384"
	case TLS_DHE_PSK_WITH_3DES_EDE_CBC_SHA:
		return "TLS_DHE_PSK_WITH_3DES_EDE_CBC_SHA"
	case TLS_DHE_PSK_WITH_AES_128_CBC_SHA:
		return "TLS_DHE_PSK_WITH_AES_128_CBC_SHA"
	case TLS_DHE_PSK_WITH_AES_128_CBC_SHA256:
		return "TLS_DHE_PSK_WITH_AES_128_CBC_SHA256"
	case TLS_DHE_PSK_WITH_AES_128_GCM_SHA256:
		return "TLS_DHE_PSK_WITH_AES_128_GCM_SHA256"
	case TLS_DHE_PSK_WITH_AES_256_CBC_SHA:
		return "TLS_DHE_PSK_WITH_AES_256_CBC_SHA"
	case TLS_DHE_PSK_WITH_AES_256_CBC_SHA384:
		return "TLS_DHE_PSK_WITH_AES_256_CBC_SHA384"
	case TLS_DHE_PSK_WITH_AES_256_GCM_SHA384:
		return "TLS_DHE_PSK_WITH_AES_256_GCM_SHA384"
	case TLS_DHE_PSK_WITH_NULL_SHA:
		return "TLS_DHE_PSK_WITH_NULL_SHA"
	case TLS_DHE_PSK_WITH_NULL_SHA256:
		return "TLS_DHE_PSK_WITH_NULL_SHA256"
	case TLS_DHE_PSK_WITH_NULL_SHA384:
		return "TLS_DHE_PSK_WITH_NULL_SHA384"
	case TLS_DHE_PSK_WITH_RC4_128_SHA:
		return "TLS_DHE_PSK_WITH_RC4_128_SHA"
	case TLS_DHE_RSA_WITH_AES_128_CBC_SHA:
		return "TLS_DHE_RSA_WITH_AES_128_CBC_SHA"
	case TLS_DHE_RSA_WITH_AES_128_CBC_SHA256:
		return "TLS_DHE_RSA_WITH_AES_128_CBC_SHA256"
	case TLS_DHE_RSA_WITH_AES_128_GCM_SHA256:
		return "TLS_DHE_RSA_WITH_AES_128_GCM_SHA256"
	case TLS_DHE_RSA_WITH_AES_256_CBC_SHA:
		return "TLS_DHE_RSA_WITH_AES_256_CBC_SHA"
	case TLS_DHE_RSA_WITH_AES_256_CBC_SHA256:
		return "TLS_DHE_RSA_WITH_AES_256_CBC_SHA256"
	case TLS_DHE_RSA_WITH_AES_256_GCM_SHA384:
		return "TLS_DHE_RSA_WITH_AES_256_GCM_SHA384"
	case TLS_DH_DSS_WITH_AES_128_CBC_SHA:
		return "TLS_DH_DSS_WITH_AES_128_CBC_SHA"
	case TLS_DH_DSS_WITH_AES_128_CBC_SHA256:
		return "TLS_DH_DSS_WITH_AES_128_CBC_SHA256"
	case TLS_DH_DSS_WITH_AES_128_GCM_SHA256:
		return "TLS_DH_DSS_WITH_AES_128_GCM_SHA256"
	case TLS_DH_DSS_WITH_AES_256_CBC_SHA:
		return "TLS_DH_DSS_WITH_AES_256_CBC_SHA"
	case TLS_DH_DSS_WITH_AES_256_CBC_SHA256:
		return "TLS_DH_DSS_WITH_AES_256_CBC_SHA256"
	case TLS_DH_DSS_WITH_AES_256_GCM_SHA384:
		return "TLS_DH_DSS_WITH_AES_256_GCM_SHA384"
	case TLS_DH_RSA_WITH_AES_128_CBC_SHA:
		return "TLS_DH_RSA_WITH_AES_128_CBC_SHA"
	case TLS_DH_RSA_WITH_AES_128_CBC_SHA256:
		return "TLS_DH_RSA_WITH_AES_128_CBC_SHA256"
	case TLS_DH_RSA_WITH_AES_128_GCM_SHA256:
		return "TLS_DH_RSA_WITH_AES_128_GCM_SHA256"
	case TLS_DH_RSA_WITH_AES_256_CBC_SHA:
		return "TLS_DH_RSA_WITH_AES_256_CBC_SHA"
	case TLS_DH_RSA_WITH_AES_256_CBC_SHA256:
		return "TLS_DH_RSA_WITH_AES_256_CBC_SHA256"
	case TLS_DH_RSA_WITH_AES_256_GCM_SHA384:
		return "TLS_DH_RSA_WITH_AES_256_GCM_SHA384"
	case TLS_DH_anon_WITH_AES_128_CBC_SHA:
		return "TLS_DH_anon_WITH_AES_128_CBC_SHA"
	case TLS_DH_anon_WITH_AES_128_CBC_SHA256:
		return "TLS_DH_anon_WITH_AES_128_CBC_SHA256"
	case TLS_DH_anon_WITH_AES_128_GCM_SHA256:
		return "TLS_DH_anon_WITH_AES_128_GCM_SHA256"
	case TLS_DH_anon_WITH_AES_256_CBC_SHA:
		return "TLS_DH_anon_WITH_AES_256_CBC_SHA"
	case TLS_DH_anon_WITH_AES_256_CBC_SHA256:
		return "TLS_DH_anon_WITH_AES_256_CBC_SHA256"
	case TLS_DH_anon_WITH_AES_256_GCM_SHA384:
		return "TLS_DH_anon_WITH_AES_256_GCM_SHA384"
	case TLS_ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA:
		return "TLS_ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA"
	case TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:
		return "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA"
	case TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256:
		return "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256"
	case TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:
		return "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
	case TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:
		return "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA"
	case TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384:
		return "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384"
	case TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:
		return "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
	case TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256:
		return "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256"
	case TLS_ECDHE_ECDSA_WITH_NULL_SHA:
		return "TLS_ECDHE_ECDSA_WITH_NULL_SHA"
	case TLS_ECDHE_ECDSA_WITH_RC4_128_SHA:
		return "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA"
	case TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA:
		return "TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA"
	case TLS_ECDHE_PSK_WITH_AES_256_CBC_SHA:
		return "TLS_ECDHE_PSK_WITH_AES_256_CBC_SHA"
	case TLS_ECDHE_PSK_WITH_CHACHA20_POLY1305_SHA256:
		return "TLS_ECDHE_PSK_WITH_CHACHA20_POLY1305_SHA256"
	case TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA:
		return "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA"
	case TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:
		return "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA"
	case TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256:
		return "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256"
	case TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:
		return "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"
	case TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:
		return "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA"
	case TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384:
		return "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384"
	case TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:
		return "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
	case TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256:
		return "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256"
	case TLS_ECDHE_RSA_WITH_NULL_SHA:
		return "TLS_ECDHE_RSA_WITH_NULL_SHA"
	case TLS_ECDHE_RSA_WITH_RC4_128_SHA:
		return "TLS_ECDHE_RSA_WITH_RC4_128_SHA"
	case TLS_ECDH_ECDSA_WITH_3DES_EDE_CBC_SHA:
		return "TLS_ECDH_ECDSA_WITH_3DES_EDE_CBC_SHA"
	case TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA:
		return "TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA"
	case TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA256:
		return "TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA256"
	case TLS_ECDH_ECDSA_WITH_AES_128_GCM_SHA256:
		return "TLS_ECDH_ECDSA_WITH_AES_128_GCM_SHA256"
	case TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA:
		return "TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA"
	case TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA384:
		return "TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA384"
	case TLS_ECDH_ECDSA_WITH_AES_256_GCM_SHA384:
		return "TLS_ECDH_ECDSA_WITH_AES_256_GCM_SHA384"
	case TLS_ECDH_ECDSA_WITH_NULL_SHA:
		return "TLS_ECDH_ECDSA_WITH_NULL_SHA"
	case TLS_ECDH_ECDSA_WITH_RC4_128_SHA:
		return "TLS_ECDH_ECDSA_WITH_RC4_128_SHA"
	case TLS_ECDH_RSA_WITH_3DES_EDE_CBC_SHA:
		return "TLS_ECDH_RSA_WITH_3DES_EDE_CBC_SHA"
	case TLS_ECDH_RSA_WITH_AES_128_CBC_SHA:
		return "TLS_ECDH_RSA_WITH_AES_128_CBC_SHA"
	case TLS_ECDH_RSA_WITH_AES_128_CBC_SHA256:
		return "TLS_ECDH_RSA_WITH_AES_128_CBC_SHA256"
	case TLS_ECDH_RSA_WITH_AES_128_GCM_SHA256:
		return "TLS_ECDH_RSA_WITH_AES_128_GCM_SHA256"
	case TLS_ECDH_RSA_WITH_AES_256_CBC_SHA:
		return "TLS_ECDH_RSA_WITH_AES_256_CBC_SHA"
	case TLS_ECDH_RSA_WITH_AES_256_CBC_SHA384:
		return "TLS_ECDH_RSA_WITH_AES_256_CBC_SHA384"
	case TLS_ECDH_RSA_WITH_AES_256_GCM_SHA384:
		return "TLS_ECDH_RSA_WITH_AES_256_GCM_SHA384"
	case TLS_ECDH_RSA_WITH_NULL_SHA:
		return "TLS_ECDH_RSA_WITH_NULL_SHA"
	case TLS_ECDH_RSA_WITH_RC4_128_SHA:
		return "TLS_ECDH_RSA_WITH_RC4_128_SHA"
	case TLS_ECDH_anon_WITH_3DES_EDE_CBC_SHA:
		return "TLS_ECDH_anon_WITH_3DES_EDE_CBC_SHA"
	case TLS_ECDH_anon_WITH_AES_128_CBC_SHA:
		return "TLS_ECDH_anon_WITH_AES_128_CBC_SHA"
	case TLS_ECDH_anon_WITH_AES_256_CBC_SHA:
		return "TLS_ECDH_anon_WITH_AES_256_CBC_SHA"
	case TLS_ECDH_anon_WITH_NULL_SHA:
		return "TLS_ECDH_anon_WITH_NULL_SHA"
	case TLS_ECDH_anon_WITH_RC4_128_SHA:
		return "TLS_ECDH_anon_WITH_RC4_128_SHA"
	case TLS_EMPTY_RENEGOTIATION_INFO_SCSV:
		return "TLS_EMPTY_RENEGOTIATION_INFO_SCSV"
	case TLS_PSK_WITH_3DES_EDE_CBC_SHA:
		return "TLS_PSK_WITH_3DES_EDE_CBC_SHA"
	case TLS_PSK_WITH_AES_128_CBC_SHA:
		return "TLS_PSK_WITH_AES_128_CBC_SHA"
	case TLS_PSK_WITH_AES_128_CBC_SHA256:
		return "TLS_PSK_WITH_AES_128_CBC_SHA256"
	case TLS_PSK_WITH_AES_128_GCM_SHA256:
		return "TLS_PSK_WITH_AES_128_GCM_SHA256"
	case TLS_PSK_WITH_AES_256_CBC_SHA:
		return "TLS_PSK_WITH_AES_256_CBC_SHA"
	case TLS_PSK_WITH_AES_256_CBC_SHA384:
		return "TLS_PSK_WITH_AES_256_CBC_SHA384"
	case TLS_PSK_WITH_AES_256_GCM_SHA384:
		return "TLS_PSK_WITH_AES_256_GCM_SHA384"
	case TLS_PSK_WITH_CHACHA20_POLY1305_SHA256:
		return "TLS_PSK_WITH_CHACHA20_POLY1305_SHA256"
	case TLS_PSK_WITH_NULL_SHA:
		return "TLS_PSK_WITH_NULL_SHA"
	case TLS_PSK_WITH_NULL_SHA256:
		return "TLS_PSK_WITH_NULL_SHA256"
	case TLS_PSK_WITH_NULL_SHA384:
		return "TLS_PSK_WITH_NULL_SHA384"
	case TLS_PSK_WITH_RC4_128_SHA:
		return "TLS_PSK_WITH_RC4_128_SHA"
	case TLS_RSA_PSK_WITH_3DES_EDE_CBC_SHA:
		return "TLS_RSA_PSK_WITH_3DES_EDE_CBC_SHA"
	case TLS_RSA_PSK_WITH_AES_128_CBC_SHA:
		return "TLS_RSA_PSK_WITH_AES_128_CBC_SHA"
	case TLS_RSA_PSK_WITH_AES_128_CBC_SHA256:
		return "TLS_RSA_PSK_WITH_AES_128_CBC_SHA256"
	case TLS_RSA_PSK_WITH_AES_128_GCM_SHA256:
		return "TLS_RSA_PSK_WITH_AES_128_GCM_SHA256"
	case TLS_RSA_PSK_WITH_AES_256_CBC_SHA:
		return "TLS_RSA_PSK_WITH_AES_256_CBC_SHA"
	case TLS_RSA_PSK_WITH_AES_256_CBC_SHA384:
		return "TLS_RSA_PSK_WITH_AES_256_CBC_SHA384"
	case TLS_RSA_PSK_WITH_AES_256_GCM_SHA384:
		return "TLS_RSA_PSK_WITH_AES_256_GCM_SHA384"
	case TLS_RSA_PSK_WITH_NULL_SHA:
		return "TLS_RSA_PSK_WITH_NULL_SHA"
	case TLS_RSA_PSK_WITH_NULL_SHA256:
		return "TLS_RSA_PSK_WITH_NULL_SHA256"
	case TLS_RSA_PSK_WITH_NULL_SHA384:
		return "TLS_RSA_PSK_WITH_NULL_SHA384"
	case TLS_RSA_PSK_WITH_RC4_128_SHA:
		return "TLS_RSA_PSK_WITH_RC4_128_SHA"
	case TLS_RSA_WITH_AES_128_CBC_SHA:
		return "TLS_RSA_WITH_AES_128_CBC_SHA"
	case TLS_RSA_WITH_AES_128_CBC_SHA256:
		return "TLS_RSA_WITH_AES_128_CBC_SHA256"
	case TLS_RSA_WITH_AES_128_GCM_SHA256:
		return "TLS_RSA_WITH_AES_128_GCM_SHA256"
	case TLS_RSA_WITH_AES_256_CBC_SHA:
		return "TLS_RSA_WITH_AES_256_CBC_SHA"
	case TLS_RSA_WITH_AES_256_CBC_SHA256:
		return "TLS_RSA_WITH_AES_256_CBC_SHA256"
	case TLS_RSA_WITH_AES_256_GCM_SHA384:
		return "TLS_RSA_WITH_AES_256_GCM_SHA384"
	case TLS_RSA_WITH_NULL_SHA256:
		return "TLS_RSA_WITH_NULL_SHA256"
	default:
		return fmt.Sprintf("SSLNullWithNullNull(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SSLProtocol
type SSLProtocol int32

const (
	// Deprecated.
	KDTLSProtocol1 SSLProtocol = 9
	// Deprecated.
	KDTLSProtocol12 SSLProtocol = 11
	// Deprecated.
	KSSLProtocol2 SSLProtocol = 1
	// Deprecated.
	KSSLProtocol3 SSLProtocol = 2
	// Deprecated.
	KSSLProtocol3Only SSLProtocol = 3
	// Deprecated.
	KSSLProtocolAll SSLProtocol = 6
	// Deprecated.
	KSSLProtocolUnknown SSLProtocol = 0
	// Deprecated.
	KTLSProtocol1 SSLProtocol = 4
	// Deprecated.
	KTLSProtocol11 SSLProtocol = 7
	// Deprecated.
	KTLSProtocol12 SSLProtocol = 8
	// Deprecated.
	KTLSProtocol13 SSLProtocol = 10
	// Deprecated.
	KTLSProtocol1Only SSLProtocol = 5
	// Deprecated.
	KTLSProtocolMaxSupported SSLProtocol = 999
)

func (e SSLProtocol) String() string {
	switch e {
	case KDTLSProtocol1:
		return "KDTLSProtocol1"
	case KDTLSProtocol12:
		return "KDTLSProtocol12"
	case KSSLProtocol2:
		return "KSSLProtocol2"
	case KSSLProtocol3:
		return "KSSLProtocol3"
	case KSSLProtocol3Only:
		return "KSSLProtocol3Only"
	case KSSLProtocolAll:
		return "KSSLProtocolAll"
	case KSSLProtocolUnknown:
		return "KSSLProtocolUnknown"
	case KTLSProtocol1:
		return "KTLSProtocol1"
	case KTLSProtocol11:
		return "KTLSProtocol11"
	case KTLSProtocol12:
		return "KTLSProtocol12"
	case KTLSProtocol13:
		return "KTLSProtocol13"
	case KTLSProtocol1Only:
		return "KTLSProtocol1Only"
	case KTLSProtocolMaxSupported:
		return "KTLSProtocolMaxSupported"
	default:
		return fmt.Sprintf("SSLProtocol(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SSLProtocolSide
type SSLProtocolSide int32

const (
	// Deprecated.
	KSSLClientSide SSLProtocolSide = 1
	// Deprecated.
	KSSLServerSide SSLProtocolSide = 0
)

func (e SSLProtocolSide) String() string {
	switch e {
	case KSSLClientSide:
		return "KSSLClientSide"
	case KSSLServerSide:
		return "KSSLServerSide"
	default:
		return fmt.Sprintf("SSLProtocolSide(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SSLSessionOption
type SSLSessionOption int32

const (
	// Deprecated.
	KSSLSessionOptionAllowRenegotiation SSLSessionOption = 8
	// Deprecated.
	KSSLSessionOptionAllowServerIdentityChange SSLSessionOption = 5
	// Deprecated.
	KSSLSessionOptionBreakOnCertRequested SSLSessionOption = 1
	// Deprecated.
	KSSLSessionOptionBreakOnClientAuth SSLSessionOption = 2
	// Deprecated.
	KSSLSessionOptionBreakOnClientHello SSLSessionOption = 7
	// Deprecated.
	KSSLSessionOptionBreakOnServerAuth SSLSessionOption = 0
	// Deprecated.
	KSSLSessionOptionEnableSessionTickets SSLSessionOption = 9
	// Deprecated.
	KSSLSessionOptionFallback SSLSessionOption = 6
	// Deprecated.
	KSSLSessionOptionFalseStart SSLSessionOption = 3
	// Deprecated.
	KSSLSessionOptionSendOneByteRecord SSLSessionOption = 4
)

func (e SSLSessionOption) String() string {
	switch e {
	case KSSLSessionOptionAllowRenegotiation:
		return "KSSLSessionOptionAllowRenegotiation"
	case KSSLSessionOptionAllowServerIdentityChange:
		return "KSSLSessionOptionAllowServerIdentityChange"
	case KSSLSessionOptionBreakOnCertRequested:
		return "KSSLSessionOptionBreakOnCertRequested"
	case KSSLSessionOptionBreakOnClientAuth:
		return "KSSLSessionOptionBreakOnClientAuth"
	case KSSLSessionOptionBreakOnClientHello:
		return "KSSLSessionOptionBreakOnClientHello"
	case KSSLSessionOptionBreakOnServerAuth:
		return "KSSLSessionOptionBreakOnServerAuth"
	case KSSLSessionOptionEnableSessionTickets:
		return "KSSLSessionOptionEnableSessionTickets"
	case KSSLSessionOptionFallback:
		return "KSSLSessionOptionFallback"
	case KSSLSessionOptionFalseStart:
		return "KSSLSessionOptionFalseStart"
	case KSSLSessionOptionSendOneByteRecord:
		return "KSSLSessionOptionSendOneByteRecord"
	default:
		return fmt.Sprintf("SSLSessionOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SSLSessionState
type SSLSessionState int32

const (
	// Deprecated.
	KSSLAborted SSLSessionState = 4
	// Deprecated.
	KSSLClosed SSLSessionState = 3
	// Deprecated.
	KSSLConnected SSLSessionState = 2
	// Deprecated.
	KSSLHandshake SSLSessionState = 1
	// Deprecated.
	KSSLIdle SSLSessionState = 0
)

func (e SSLSessionState) String() string {
	switch e {
	case KSSLAborted:
		return "KSSLAborted"
	case KSSLClosed:
		return "KSSLClosed"
	case KSSLConnected:
		return "KSSLConnected"
	case KSSLHandshake:
		return "KSSLHandshake"
	case KSSLIdle:
		return "KSSLIdle"
	default:
		return fmt.Sprintf("SSLSessionState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags
type SecAccessControlCreateFlags uint

const (
	// KSecAccessControlAnd: Indicates that all constraints must be satisfied.
	KSecAccessControlAnd SecAccessControlCreateFlags = 32768
	// KSecAccessControlApplicationPassword: Option to use an application-provided password for data encryption key generation.
	KSecAccessControlApplicationPassword SecAccessControlCreateFlags = 2147483648
	// KSecAccessControlBiometryAny: Constraint to access an item with Touch ID for any enrolled fingers, or Face ID.
	KSecAccessControlBiometryAny SecAccessControlCreateFlags = 2
	// KSecAccessControlBiometryCurrentSet: Constraint to access an item with Touch ID for currently enrolled fingers, or from Face ID with the currently enrolled user.
	KSecAccessControlBiometryCurrentSet SecAccessControlCreateFlags = 8
	KSecAccessControlCompanion          SecAccessControlCreateFlags = 32
	// KSecAccessControlDevicePasscode: Constraint to access an item with a passcode.
	KSecAccessControlDevicePasscode SecAccessControlCreateFlags = 16
	// KSecAccessControlOr: Indicates that at least one constraint must be satisfied.
	KSecAccessControlOr SecAccessControlCreateFlags = 16384
	// KSecAccessControlPrivateKeyUsage: Enable a private key to be used in signing a block of data or verifying a signed block.
	KSecAccessControlPrivateKeyUsage SecAccessControlCreateFlags = 1073741824
	// KSecAccessControlUserPresence: Constraint to access an item with either biometry or passcode.
	KSecAccessControlUserPresence SecAccessControlCreateFlags = 1
	// Deprecated: use KSecAccessControlBiometryAny.
	KSecAccessControlTouchIDAny SecAccessControlCreateFlags = 2
	// Deprecated: use KSecAccessControlBiometryCurrentSet.
	KSecAccessControlTouchIDCurrentSet SecAccessControlCreateFlags = 8
	// Deprecated: use KSecAccessControlCompanion.
	KSecAccessControlWatch SecAccessControlCreateFlags = 32
)

func (e SecAccessControlCreateFlags) String() string {
	switch e {
	case KSecAccessControlAnd:
		return "KSecAccessControlAnd"
	case KSecAccessControlApplicationPassword:
		return "KSecAccessControlApplicationPassword"
	case KSecAccessControlBiometryAny:
		return "KSecAccessControlBiometryAny"
	case KSecAccessControlBiometryCurrentSet:
		return "KSecAccessControlBiometryCurrentSet"
	case KSecAccessControlCompanion:
		return "KSecAccessControlCompanion"
	case KSecAccessControlDevicePasscode:
		return "KSecAccessControlDevicePasscode"
	case KSecAccessControlOr:
		return "KSecAccessControlOr"
	case KSecAccessControlPrivateKeyUsage:
		return "KSecAccessControlPrivateKeyUsage"
	case KSecAccessControlUserPresence:
		return "KSecAccessControlUserPresence"
	default:
		return fmt.Sprintf("SecAccessControlCreateFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecAuthenticationType
type SecAuthenticationType uint32

const (
	// KSecAuthenticationTypeAny: Specifies that any authentication type is acceptable.
	KSecAuthenticationTypeAny SecAuthenticationType = 0
	// KSecAuthenticationTypeDPA: Specifies Distributed Password authentication.
	KSecAuthenticationTypeDPA SecAuthenticationType = 'a'<<24 | 'a'<<16 | 'p'<<8 | 'd' // 'aapd'
	// KSecAuthenticationTypeDefault: Specifies the default authentication type.
	KSecAuthenticationTypeDefault SecAuthenticationType = 't'<<24 | 'l'<<16 | 'f'<<8 | 'd' // 'tlfd'
	// KSecAuthenticationTypeHTMLForm: Specifies HTML form based authentication.
	KSecAuthenticationTypeHTMLForm SecAuthenticationType = 'm'<<24 | 'r'<<16 | 'o'<<8 | 'f' // 'mrof'
	// KSecAuthenticationTypeHTTPBasic: Specifies HTTP Basic authentication.
	KSecAuthenticationTypeHTTPBasic SecAuthenticationType = 'p'<<24 | 't'<<16 | 't'<<8 | 'h' // 'ptth'
	// KSecAuthenticationTypeHTTPDigest: Specifies HTTP Digest Access authentication.
	KSecAuthenticationTypeHTTPDigest SecAuthenticationType = 'd'<<24 | 't'<<16 | 't'<<8 | 'h' // 'dtth'
	// KSecAuthenticationTypeMSN: Specifies Microsoft Network default authentication.
	KSecAuthenticationTypeMSN SecAuthenticationType = 'a'<<24 | 'n'<<16 | 's'<<8 | 'm' // 'ansm'
	// KSecAuthenticationTypeNTLM: Specifies Windows NT LAN Manager authentication.
	KSecAuthenticationTypeNTLM SecAuthenticationType = 'm'<<24 | 'l'<<16 | 't'<<8 | 'n' // 'mltn'
	// KSecAuthenticationTypeRPA: Specifies Remote Password authentication.
	KSecAuthenticationTypeRPA SecAuthenticationType = 'a'<<24 | 'a'<<16 | 'p'<<8 | 'r' // 'aapr'
)

func (e SecAuthenticationType) String() string {
	switch e {
	case KSecAuthenticationTypeAny:
		return "KSecAuthenticationTypeAny"
	case KSecAuthenticationTypeDPA:
		return "KSecAuthenticationTypeDPA"
	case KSecAuthenticationTypeDefault:
		return "KSecAuthenticationTypeDefault"
	case KSecAuthenticationTypeHTMLForm:
		return "KSecAuthenticationTypeHTMLForm"
	case KSecAuthenticationTypeHTTPBasic:
		return "KSecAuthenticationTypeHTTPBasic"
	case KSecAuthenticationTypeHTTPDigest:
		return "KSecAuthenticationTypeHTTPDigest"
	case KSecAuthenticationTypeMSN:
		return "KSecAuthenticationTypeMSN"
	case KSecAuthenticationTypeNTLM:
		return "KSecAuthenticationTypeNTLM"
	case KSecAuthenticationTypeRPA:
		return "KSecAuthenticationTypeRPA"
	default:
		return fmt.Sprintf("SecAuthenticationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecCSDigestAlgorithm
type SecCSDigestAlgorithm uint32

const (
	KSecCodeSignatureHashSHA1            SecCSDigestAlgorithm = 1
	KSecCodeSignatureHashSHA256          SecCSDigestAlgorithm = 2
	KSecCodeSignatureHashSHA256Truncated SecCSDigestAlgorithm = 3
	KSecCodeSignatureHashSHA384          SecCSDigestAlgorithm = 4
	KSecCodeSignatureHashSHA512          SecCSDigestAlgorithm = 5
	KSecCodeSignatureNoHash              SecCSDigestAlgorithm = 0
)

func (e SecCSDigestAlgorithm) String() string {
	switch e {
	case KSecCodeSignatureHashSHA1:
		return "KSecCodeSignatureHashSHA1"
	case KSecCodeSignatureHashSHA256:
		return "KSecCodeSignatureHashSHA256"
	case KSecCodeSignatureHashSHA256Truncated:
		return "KSecCodeSignatureHashSHA256Truncated"
	case KSecCodeSignatureHashSHA384:
		return "KSecCodeSignatureHashSHA384"
	case KSecCodeSignatureHashSHA512:
		return "KSecCodeSignatureHashSHA512"
	case KSecCodeSignatureNoHash:
		return "KSecCodeSignatureNoHash"
	default:
		return fmt.Sprintf("SecCSDigestAlgorithm(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecCSFlags
type SecCSFlags uint32

const (
	KSecCSApplyEmbeddedPolicy SecCSFlags = 33554432
	KSecCSCheckTrustedAnchors SecCSFlags = 134217728
	// KSecCSConsiderExpiration: Consider expired certificates invalid.
	KSecCSConsiderExpiration SecCSFlags = 2147483648
	// KSecCSDefaultFlags: No flags (use the default behavior).
	KSecCSDefaultFlags SecCSFlags = 0
	// KSecCSEnforceRevocationChecks: # Discussion
	KSecCSEnforceRevocationChecks       SecCSFlags = 1073741824
	KSecCSMatchGuestRequirementInKernel SecCSFlags = 8388608
	KSecCSNoNetworkAccess               SecCSFlags = 536870912
	KSecCSQuickCheck                    SecCSFlags = 67108864
	KSecCSReportProgress                SecCSFlags = 268435456
	KSecCSStripDisallowedXattrs         SecCSFlags = 16777216
)

func (e SecCSFlags) String() string {
	switch e {
	case KSecCSApplyEmbeddedPolicy:
		return "KSecCSApplyEmbeddedPolicy"
	case KSecCSCheckTrustedAnchors:
		return "KSecCSCheckTrustedAnchors"
	case KSecCSConsiderExpiration:
		return "KSecCSConsiderExpiration"
	case KSecCSDefaultFlags:
		return "KSecCSDefaultFlags"
	case KSecCSEnforceRevocationChecks:
		return "KSecCSEnforceRevocationChecks"
	case KSecCSMatchGuestRequirementInKernel:
		return "KSecCSMatchGuestRequirementInKernel"
	case KSecCSNoNetworkAccess:
		return "KSecCSNoNetworkAccess"
	case KSecCSQuickCheck:
		return "KSecCSQuickCheck"
	case KSecCSReportProgress:
		return "KSecCSReportProgress"
	case KSecCSStripDisallowedXattrs:
		return "KSecCSStripDisallowedXattrs"
	default:
		return fmt.Sprintf("SecCSFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags
type SecCodeSignatureFlags uint32

const (
	// KSecCodeSignatureAdhoc: Must be used without a signing identity.
	KSecCodeSignatureAdhoc SecCodeSignatureFlags = 0x2
	// KSecCodeSignatureEnforcement: Enforce code signing.
	KSecCodeSignatureEnforcement SecCodeSignatureFlags = 0x1000
	// KSecCodeSignatureForceExpiration: Always set the considerExpiration flag when validating the code.
	KSecCodeSignatureForceExpiration SecCodeSignatureFlags = 0x400
	// KSecCodeSignatureForceHard: Always set the hard status flag on launch.
	KSecCodeSignatureForceHard SecCodeSignatureFlags = 0x100
	// KSecCodeSignatureForceKill: Always set the termination status flag on launch.
	KSecCodeSignatureForceKill SecCodeSignatureFlags = 0x200
	// KSecCodeSignatureHost: May host guest code.
	KSecCodeSignatureHost SecCodeSignatureFlags = 0x1
	// KSecCodeSignatureLibraryValidation: Require library validation.
	KSecCodeSignatureLibraryValidation SecCodeSignatureFlags = 0x2000
	KSecCodeSignatureLinkerSigned      SecCodeSignatureFlags = 0x20000
	// KSecCodeSignatureRestrict: Restrict dyld loading.
	KSecCodeSignatureRestrict SecCodeSignatureFlags = 0x800
	// KSecCodeSignatureRuntime: Apply runtime hardening policies as required by the hardened runtime version.
	KSecCodeSignatureRuntime SecCodeSignatureFlags = 0x10000
)

func (e SecCodeSignatureFlags) String() string {
	switch e {
	case KSecCodeSignatureAdhoc:
		return "KSecCodeSignatureAdhoc"
	case KSecCodeSignatureEnforcement:
		return "KSecCodeSignatureEnforcement"
	case KSecCodeSignatureForceExpiration:
		return "KSecCodeSignatureForceExpiration"
	case KSecCodeSignatureForceHard:
		return "KSecCodeSignatureForceHard"
	case KSecCodeSignatureForceKill:
		return "KSecCodeSignatureForceKill"
	case KSecCodeSignatureHost:
		return "KSecCodeSignatureHost"
	case KSecCodeSignatureLibraryValidation:
		return "KSecCodeSignatureLibraryValidation"
	case KSecCodeSignatureLinkerSigned:
		return "KSecCodeSignatureLinkerSigned"
	case KSecCodeSignatureRestrict:
		return "KSecCodeSignatureRestrict"
	case KSecCodeSignatureRuntime:
		return "KSecCodeSignatureRuntime"
	default:
		return fmt.Sprintf("SecCodeSignatureFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecCodeStatus
type SecCodeStatus uint32

const (
	// KSecCodeStatusDebugged: The code has been debugged by another process that was allowed to do so.
	KSecCodeStatusDebugged SecCodeStatus = 0x10000000
	// KSecCodeStatusHard: The code prefers to be denied access to resources if gaining access would invalidate it.
	KSecCodeStatusHard SecCodeStatus = 0x100
	// KSecCodeStatusKill: The code wants to be terminated if it ever loses its validity.
	KSecCodeStatusKill SecCodeStatus = 0x200
	// KSecCodeStatusPlatform: The code ships with the operating system and is signed by Apple.
	KSecCodeStatusPlatform SecCodeStatus = 0x4000000
	// KSecCodeStatusValid: The code is dynamically valid.
	KSecCodeStatusValid SecCodeStatus = 0x1
)

func (e SecCodeStatus) String() string {
	switch e {
	case KSecCodeStatusDebugged:
		return "KSecCodeStatusDebugged"
	case KSecCodeStatusHard:
		return "KSecCodeStatusHard"
	case KSecCodeStatusKill:
		return "KSecCodeStatusKill"
	case KSecCodeStatusPlatform:
		return "KSecCodeStatusPlatform"
	case KSecCodeStatusValid:
		return "KSecCodeStatusValid"
	default:
		return fmt.Sprintf("SecCodeStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecCredentialType
type SecCredentialType uint32

const (
	// Deprecated.
	KSecCredentialTypeDefault SecCredentialType = 0
	// Deprecated.
	KSecCredentialTypeNoUI SecCredentialType = 2
	// Deprecated.
	KSecCredentialTypeWithUI SecCredentialType = 1
)

func (e SecCredentialType) String() string {
	switch e {
	case KSecCredentialTypeDefault:
		return "KSecCredentialTypeDefault"
	case KSecCredentialTypeNoUI:
		return "KSecCredentialTypeNoUI"
	case KSecCredentialTypeWithUI:
		return "KSecCredentialTypeWithUI"
	default:
		return fmt.Sprintf("SecCredentialType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecExternalFormat
type SecExternalFormat uint32

const (
	// KSecFormatBSAFE: Format for asymmetric keys.
	KSecFormatBSAFE SecExternalFormat = 3
	// KSecFormatNetscapeCertSequence: Set of certificates in the Netscape Certificate Sequence format.
	KSecFormatNetscapeCertSequence SecExternalFormat = 13
	// KSecFormatOpenSSL: Format for asymmetric (public/private) keys.
	KSecFormatOpenSSL SecExternalFormat = 1
	// KSecFormatPEMSequence: Sequence of certificates and keys with PEM armor.
	KSecFormatPEMSequence SecExternalFormat = 10
	// KSecFormatPKCS12: Set of certificates and private keys.
	KSecFormatPKCS12 SecExternalFormat = 12
	// KSecFormatPKCS7: Sequence of certificates, no PEM armor.
	KSecFormatPKCS7 SecExternalFormat = 11
	// KSecFormatRawKey: Format for symmetric keys.
	KSecFormatRawKey SecExternalFormat = 4
	// KSecFormatSSH: OpenSSH 1 format for asymmetric (public/private) keys.
	KSecFormatSSH SecExternalFormat = 2
	// KSecFormatSSHv2: OpenSSH 2 format for public keys.
	KSecFormatSSHv2 SecExternalFormat = 14
	// KSecFormatUnknown: # Discussion
	KSecFormatUnknown SecExternalFormat = 0
	// KSecFormatWrappedLSH: Not supported.
	KSecFormatWrappedLSH SecExternalFormat = 8
	// KSecFormatWrappedOpenSSL: Format for wrapped symmetric and private keys.
	KSecFormatWrappedOpenSSL SecExternalFormat = 6
	// KSecFormatWrappedPKCS8: Format for wrapped symmetric and private keys.
	KSecFormatWrappedPKCS8 SecExternalFormat = 5
	// KSecFormatWrappedSSH: OpenSSH 1 format for wrapped symmetric and private keys.
	KSecFormatWrappedSSH SecExternalFormat = 7
	// KSecFormatX509Cert: Format for certificates.
	KSecFormatX509Cert SecExternalFormat = 9
)

func (e SecExternalFormat) String() string {
	switch e {
	case KSecFormatBSAFE:
		return "KSecFormatBSAFE"
	case KSecFormatNetscapeCertSequence:
		return "KSecFormatNetscapeCertSequence"
	case KSecFormatOpenSSL:
		return "KSecFormatOpenSSL"
	case KSecFormatPEMSequence:
		return "KSecFormatPEMSequence"
	case KSecFormatPKCS12:
		return "KSecFormatPKCS12"
	case KSecFormatPKCS7:
		return "KSecFormatPKCS7"
	case KSecFormatRawKey:
		return "KSecFormatRawKey"
	case KSecFormatSSH:
		return "KSecFormatSSH"
	case KSecFormatSSHv2:
		return "KSecFormatSSHv2"
	case KSecFormatUnknown:
		return "KSecFormatUnknown"
	case KSecFormatWrappedLSH:
		return "KSecFormatWrappedLSH"
	case KSecFormatWrappedOpenSSL:
		return "KSecFormatWrappedOpenSSL"
	case KSecFormatWrappedPKCS8:
		return "KSecFormatWrappedPKCS8"
	case KSecFormatWrappedSSH:
		return "KSecFormatWrappedSSH"
	case KSecFormatX509Cert:
		return "KSecFormatX509Cert"
	default:
		return fmt.Sprintf("SecExternalFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecExternalItemType
type SecExternalItemType uint32

const (
	// KSecItemTypeAggregate: Indicates a set of certificates or certificates and private keys.
	KSecItemTypeAggregate SecExternalItemType = 5
	// KSecItemTypeCertificate: Indicates a certificate.
	KSecItemTypeCertificate SecExternalItemType = 4
	// KSecItemTypePrivateKey: Indicates a private key.
	KSecItemTypePrivateKey SecExternalItemType = 1
	// KSecItemTypePublicKey: Indicates a public key.
	KSecItemTypePublicKey SecExternalItemType = 2
	// KSecItemTypeSessionKey: Indicates a session key.
	KSecItemTypeSessionKey SecExternalItemType = 3
	// KSecItemTypeUnknown: Indicates that the caller does not know the type of information being imported or exported.
	KSecItemTypeUnknown SecExternalItemType = 0
)

func (e SecExternalItemType) String() string {
	switch e {
	case KSecItemTypeAggregate:
		return "KSecItemTypeAggregate"
	case KSecItemTypeCertificate:
		return "KSecItemTypeCertificate"
	case KSecItemTypePrivateKey:
		return "KSecItemTypePrivateKey"
	case KSecItemTypePublicKey:
		return "KSecItemTypePublicKey"
	case KSecItemTypeSessionKey:
		return "KSecItemTypeSessionKey"
	case KSecItemTypeUnknown:
		return "KSecItemTypeUnknown"
	default:
		return fmt.Sprintf("SecExternalItemType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecItemAttr
type SecItemAttr uint32

const (
	// KSecAccountItemAttr: Identifies the account attribute.
	KSecAccountItemAttr SecItemAttr = 'a'<<24 | 'c'<<16 | 'c'<<8 | 't' // 'acct'
	// KSecAddressItemAttr: Identifies the address attribute.
	KSecAddressItemAttr SecItemAttr = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'r' // 'addr'
	// KSecAlias: Indicates an alias.
	KSecAlias SecItemAttr = 'a'<<24 | 'l'<<16 | 'i'<<8 | 's' // 'alis'
	// KSecAuthenticationTypeItemAttr: Identifies the authentication type attribute.
	KSecAuthenticationTypeItemAttr SecItemAttr = 'a'<<24 | 't'<<16 | 'y'<<8 | 'p' // 'atyp'
	// KSecCertificateEncoding: Indicates a `CSSM_CERT_ENCODING` type.
	KSecCertificateEncoding SecItemAttr = 'c'<<24 | 'e'<<16 | 'n'<<8 | 'c' // 'cenc'
	// KSecCertificateType: Indicates a `CSSM_CERT_TYPE` type.
	KSecCertificateType SecItemAttr = 'c'<<24 | 't'<<16 | 'y'<<8 | 'p' // 'ctyp'
	// KSecCommentItemAttr: Identifies the comment attribute.
	KSecCommentItemAttr SecItemAttr = 'i'<<24 | 'c'<<16 | 'm'<<8 | 't' // 'icmt'
	// KSecCreationDateItemAttr: Identifies the creation date attribute.
	KSecCreationDateItemAttr SecItemAttr = 'c'<<24 | 'd'<<16 | 'a'<<8 | 't' // 'cdat'
	// KSecCreatorItemAttr: Identifies the creator attribute.
	KSecCreatorItemAttr SecItemAttr = 'c'<<24 | 'r'<<16 | 't'<<8 | 'r' // 'crtr'
	// KSecCrlEncoding: Indicates a `CSSM_CRL_ENCODING` type.
	KSecCrlEncoding SecItemAttr = 'c'<<24 | 'r'<<16 | 'n'<<8 | 'c' // 'crnc'
	// KSecCrlType: Indicates a `CSSM_CRL_TYPE` type.
	KSecCrlType SecItemAttr = 'c'<<24 | 'r'<<16 | 't'<<8 | 'p' // 'crtp'
	// KSecCustomIconItemAttr: Identifies the custom icon attribute.
	KSecCustomIconItemAttr SecItemAttr = 'c'<<24 | 'u'<<16 | 's'<<8 | 'i' // 'cusi'
	// KSecDescriptionItemAttr: Identifies the description attribute.
	KSecDescriptionItemAttr SecItemAttr = 'd'<<24 | 'e'<<16 | 's'<<8 | 'c' // 'desc'
	// KSecGenericItemAttr: Identifies the generic attribute.
	KSecGenericItemAttr SecItemAttr = 'g'<<24 | 'e'<<16 | 'n'<<8 | 'a' // 'gena'
	// KSecInvisibleItemAttr: Identifies the invisible attribute.
	KSecInvisibleItemAttr SecItemAttr = 'i'<<24 | 'n'<<16 | 'v'<<8 | 'i' // 'invi'
	// KSecLabelItemAttr: Identifies the label attribute.
	KSecLabelItemAttr SecItemAttr = 'l'<<24 | 'a'<<16 | 'b'<<8 | 'l' // 'labl'
	// KSecModDateItemAttr: Identifies the modification date attribute.
	KSecModDateItemAttr SecItemAttr = 'm'<<24 | 'd'<<16 | 'a'<<8 | 't' // 'mdat'
	// KSecNegativeItemAttr: Identifies the negative attribute.
	KSecNegativeItemAttr SecItemAttr = 'n'<<24 | 'e'<<16 | 'g'<<8 | 'a' // 'nega'
	// KSecPathItemAttr: Identifies the path attribute.
	KSecPathItemAttr SecItemAttr = 'p'<<24 | 'a'<<16 | 't'<<8 | 'h' // 'path'
	// KSecPortItemAttr: Identifies the port attribute.
	KSecPortItemAttr SecItemAttr = 'p'<<24 | 'o'<<16 | 'r'<<8 | 't' // 'port'
	// KSecProtocolItemAttr: Identifies the protocol attribute.
	KSecProtocolItemAttr SecItemAttr = 'p'<<24 | 't'<<16 | 'c'<<8 | 'l' // 'ptcl'
	// KSecScriptCodeItemAttr: Identifies the script code attribute.
	KSecScriptCodeItemAttr SecItemAttr = 's'<<24 | 'c'<<16 | 'r'<<8 | 'p' // 'scrp'
	// KSecSecurityDomainItemAttr: Identifies the security domain attribute.
	KSecSecurityDomainItemAttr SecItemAttr = 's'<<24 | 'd'<<16 | 'm'<<8 | 'n' // 'sdmn'
	// KSecServerItemAttr: Identifies the server attribute.
	KSecServerItemAttr SecItemAttr = 's'<<24 | 'r'<<16 | 'v'<<8 | 'r' // 'srvr'
	// KSecServiceItemAttr: Identifies the service attribute.
	KSecServiceItemAttr SecItemAttr = 's'<<24 | 'v'<<16 | 'c'<<8 | 'e' // 'svce'
	// KSecSignatureItemAttr: Identifies the server signature attribute.
	KSecSignatureItemAttr SecItemAttr = 's'<<24 | 's'<<16 | 'i'<<8 | 'g' // 'ssig'
	// KSecTypeItemAttr: Identifies the type attribute.
	KSecTypeItemAttr SecItemAttr = 't'<<24 | 'y'<<16 | 'p'<<8 | 'e' // 'type'
	// KSecVolumeItemAttr: Identifies the volume attribute.
	KSecVolumeItemAttr SecItemAttr = 'v'<<24 | 'l'<<16 | 'm'<<8 | 'e' // 'vlme'
)

func (e SecItemAttr) String() string {
	switch e {
	case KSecAccountItemAttr:
		return "KSecAccountItemAttr"
	case KSecAddressItemAttr:
		return "KSecAddressItemAttr"
	case KSecAlias:
		return "KSecAlias"
	case KSecAuthenticationTypeItemAttr:
		return "KSecAuthenticationTypeItemAttr"
	case KSecCertificateEncoding:
		return "KSecCertificateEncoding"
	case KSecCertificateType:
		return "KSecCertificateType"
	case KSecCommentItemAttr:
		return "KSecCommentItemAttr"
	case KSecCreationDateItemAttr:
		return "KSecCreationDateItemAttr"
	case KSecCreatorItemAttr:
		return "KSecCreatorItemAttr"
	case KSecCrlEncoding:
		return "KSecCrlEncoding"
	case KSecCrlType:
		return "KSecCrlType"
	case KSecCustomIconItemAttr:
		return "KSecCustomIconItemAttr"
	case KSecDescriptionItemAttr:
		return "KSecDescriptionItemAttr"
	case KSecGenericItemAttr:
		return "KSecGenericItemAttr"
	case KSecInvisibleItemAttr:
		return "KSecInvisibleItemAttr"
	case KSecLabelItemAttr:
		return "KSecLabelItemAttr"
	case KSecModDateItemAttr:
		return "KSecModDateItemAttr"
	case KSecNegativeItemAttr:
		return "KSecNegativeItemAttr"
	case KSecPathItemAttr:
		return "KSecPathItemAttr"
	case KSecPortItemAttr:
		return "KSecPortItemAttr"
	case KSecProtocolItemAttr:
		return "KSecProtocolItemAttr"
	case KSecScriptCodeItemAttr:
		return "KSecScriptCodeItemAttr"
	case KSecSecurityDomainItemAttr:
		return "KSecSecurityDomainItemAttr"
	case KSecServerItemAttr:
		return "KSecServerItemAttr"
	case KSecServiceItemAttr:
		return "KSecServiceItemAttr"
	case KSecSignatureItemAttr:
		return "KSecSignatureItemAttr"
	case KSecTypeItemAttr:
		return "KSecTypeItemAttr"
	case KSecVolumeItemAttr:
		return "KSecVolumeItemAttr"
	default:
		return fmt.Sprintf("SecItemAttr(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecItemClass
type SecItemClass uint32

const (
	// KSecCertificateItemClass: Indicates that the item is an X509 certificate.
	KSecCertificateItemClass SecItemClass = 0x80001000
	// KSecGenericPasswordItemClass: Indicates that the item is a generic password.
	KSecGenericPasswordItemClass SecItemClass = 'g'<<24 | 'e'<<16 | 'n'<<8 | 'p' // 'genp'
	// KSecInternetPasswordItemClass: Indicates that the item is an Internet password.
	KSecInternetPasswordItemClass SecItemClass = 'i'<<24 | 'n'<<16 | 'e'<<8 | 't' // 'inet'
	// KSecPrivateKeyItemClass: Indicates that the item is a private key of a public-private pair.
	KSecPrivateKeyItemClass SecItemClass = 0x10
	// KSecPublicKeyItemClass: Indicates that the item is a public key of a public-private pair.
	KSecPublicKeyItemClass SecItemClass = 0xf
	// KSecSymmetricKeyItemClass: Indicates that the item is a private key used for symmetric-key encryption.
	KSecSymmetricKeyItemClass SecItemClass = 0x11
	// Deprecated.
	KSecAppleSharePasswordItemClass SecItemClass = 'a'<<24 | 's'<<16 | 'h'<<8 | 'p' // 'ashp'
)

func (e SecItemClass) String() string {
	switch e {
	case KSecCertificateItemClass:
		return "KSecCertificateItemClass"
	case KSecGenericPasswordItemClass:
		return "KSecGenericPasswordItemClass"
	case KSecInternetPasswordItemClass:
		return "KSecInternetPasswordItemClass"
	case KSecPrivateKeyItemClass:
		return "KSecPrivateKeyItemClass"
	case KSecPublicKeyItemClass:
		return "KSecPublicKeyItemClass"
	case KSecSymmetricKeyItemClass:
		return "KSecSymmetricKeyItemClass"
	case KSecAppleSharePasswordItemClass:
		return "KSecAppleSharePasswordItemClass"
	default:
		return fmt.Sprintf("SecItemClass(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecItemImportExportFlags
type SecItemImportExportFlags uint32

const (
	// KSecItemPemArmour: A flag that indicates the exported data should have PEM armor.
	KSecItemPemArmour SecItemImportExportFlags = 0x1
)

func (e SecItemImportExportFlags) String() string {
	switch e {
	case KSecItemPemArmour:
		return "KSecItemPemArmour"
	default:
		return fmt.Sprintf("SecItemImportExportFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecKeyImportExportFlags
type SecKeyImportExportFlags uint32

const (
	// KSecKeyImportOnlyOne: A flag that you set to prevent importing more than one private key.
	KSecKeyImportOnlyOne SecKeyImportExportFlags = 0x1
	// KSecKeyNoAccessControl: A flag that indicates imported private keys have no access object attached to them.
	KSecKeyNoAccessControl SecKeyImportExportFlags = 0x4
	// KSecKeySecurePassphrase: A flag that indicates the user should be prompted for a passphrase on import or export.
	KSecKeySecurePassphrase SecKeyImportExportFlags = 0x2
)

func (e SecKeyImportExportFlags) String() string {
	switch e {
	case KSecKeyImportOnlyOne:
		return "KSecKeyImportOnlyOne"
	case KSecKeyNoAccessControl:
		return "KSecKeyNoAccessControl"
	case KSecKeySecurePassphrase:
		return "KSecKeySecurePassphrase"
	default:
		return fmt.Sprintf("SecKeyImportExportFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecKeyOperationType
type SecKeyOperationType int

const (
	KSecKeyOperationTypeDecrypt     SecKeyOperationType = 3
	KSecKeyOperationTypeEncrypt     SecKeyOperationType = 2
	KSecKeyOperationTypeKeyExchange SecKeyOperationType = 4
	KSecKeyOperationTypeSign        SecKeyOperationType = 0
	KSecKeyOperationTypeVerify      SecKeyOperationType = 1
)

func (e SecKeyOperationType) String() string {
	switch e {
	case KSecKeyOperationTypeDecrypt:
		return "KSecKeyOperationTypeDecrypt"
	case KSecKeyOperationTypeEncrypt:
		return "KSecKeyOperationTypeEncrypt"
	case KSecKeyOperationTypeKeyExchange:
		return "KSecKeyOperationTypeKeyExchange"
	case KSecKeyOperationTypeSign:
		return "KSecKeyOperationTypeSign"
	case KSecKeyOperationTypeVerify:
		return "KSecKeyOperationTypeVerify"
	default:
		return fmt.Sprintf("SecKeyOperationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecKeySizes
type SecKeySizes uint32

const (
	// Deprecated.
	KSec3DES192 SecKeySizes = 192
	// Deprecated.
	KSecAES128 SecKeySizes = 128
	// Deprecated.
	KSecAES192 SecKeySizes = 192
	// Deprecated.
	KSecAES256 SecKeySizes = 256
	// Deprecated.
	KSecDefaultKeySize SecKeySizes = 0
	// Deprecated.
	KSecRSAMax SecKeySizes = 4096
	// Deprecated.
	KSecRSAMin SecKeySizes = 1024
	// Deprecated.
	KSecp192r1 SecKeySizes = 192
	// Deprecated.
	KSecp256r1 SecKeySizes = 256
	// Deprecated.
	KSecp384r1 SecKeySizes = 384
	// Deprecated.
	KSecp521r1 SecKeySizes = 521
)

func (e SecKeySizes) String() string {
	switch e {
	case KSec3DES192:
		return "KSec3DES192"
	case KSecAES128:
		return "KSecAES128"
	case KSecAES256:
		return "KSecAES256"
	case KSecDefaultKeySize:
		return "KSecDefaultKeySize"
	case KSecRSAMax:
		return "KSecRSAMax"
	case KSecRSAMin:
		return "KSecRSAMin"
	case KSecp384r1:
		return "KSecp384r1"
	case KSecp521r1:
		return "KSecp521r1"
	default:
		return fmt.Sprintf("SecKeySizes(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecKeyUsage
type SecKeyUsage uint32

const (
	// KSecKeyUsageAll: All flags set.
	KSecKeyUsageAll SecKeyUsage = 0x7fffffff
	// KSecKeyUsageCRLSign: The [CRLSign] bit is set in KeyUsage extension.
	KSecKeyUsageCRLSign SecKeyUsage = 64
	// KSecKeyUsageContentCommitment: The [ContentCommitment] bit is set in KeyUsage extension.
	KSecKeyUsageContentCommitment SecKeyUsage = 2
	// KSecKeyUsageCritical: The KeyUsage extension is marked critical.
	KSecKeyUsageCritical SecKeyUsage = 2147483648
	// KSecKeyUsageDataEncipherment: The [DataEncipherment] bit is set in KeyUsage extension.
	KSecKeyUsageDataEncipherment SecKeyUsage = 8
	// KSecKeyUsageDecipherOnly: The [DecipherOnly] bit is set in KeyUsage extension.
	KSecKeyUsageDecipherOnly SecKeyUsage = 256
	// KSecKeyUsageDigitalSignature: The [DigitalSignature] bit is set in KeyUsage extension.
	KSecKeyUsageDigitalSignature SecKeyUsage = 1
	// KSecKeyUsageEncipherOnly: The [EncipherOnly] bit is set in KeyUsage extension.
	KSecKeyUsageEncipherOnly SecKeyUsage = 128
	// KSecKeyUsageKeyAgreement: The [KeyAgreement] bit is set in KeyUsage extension.
	KSecKeyUsageKeyAgreement SecKeyUsage = 16
	// KSecKeyUsageKeyCertSign: The [KeyCertSign] bit is set in KeyUsage extension.
	KSecKeyUsageKeyCertSign SecKeyUsage = 32
	// KSecKeyUsageKeyEncipherment: The [KeyEncipherment] bit is set in KeyUsage extension.
	KSecKeyUsageKeyEncipherment SecKeyUsage = 4
	// KSecKeyUsageNonRepudiation: The [NonRepudiation] bit is set in KeyUsage extension.
	KSecKeyUsageNonRepudiation SecKeyUsage = 2
	KSecKeyUsageUnspecified    SecKeyUsage = 0
)

func (e SecKeyUsage) String() string {
	switch e {
	case KSecKeyUsageAll:
		return "KSecKeyUsageAll"
	case KSecKeyUsageCRLSign:
		return "KSecKeyUsageCRLSign"
	case KSecKeyUsageContentCommitment:
		return "KSecKeyUsageContentCommitment"
	case KSecKeyUsageCritical:
		return "KSecKeyUsageCritical"
	case KSecKeyUsageDataEncipherment:
		return "KSecKeyUsageDataEncipherment"
	case KSecKeyUsageDecipherOnly:
		return "KSecKeyUsageDecipherOnly"
	case KSecKeyUsageDigitalSignature:
		return "KSecKeyUsageDigitalSignature"
	case KSecKeyUsageEncipherOnly:
		return "KSecKeyUsageEncipherOnly"
	case KSecKeyUsageKeyAgreement:
		return "KSecKeyUsageKeyAgreement"
	case KSecKeyUsageKeyCertSign:
		return "KSecKeyUsageKeyCertSign"
	case KSecKeyUsageKeyEncipherment:
		return "KSecKeyUsageKeyEncipherment"
	case KSecKeyUsageUnspecified:
		return "KSecKeyUsageUnspecified"
	default:
		return fmt.Sprintf("SecKeyUsage(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecKeychainEvent
type SecKeychainEvent uint32

const (
	// KSecAddEvent: Indicates an item was added to a keychain.
	KSecAddEvent SecKeychainEvent = 3
	// KSecDefaultChangedEvent: Indicates that a different keychain was specified as the default.
	KSecDefaultChangedEvent SecKeychainEvent = 9
	// KSecDeleteEvent: Indicates an item was deleted from a keychain.
	KSecDeleteEvent SecKeychainEvent = 4
	// KSecKeychainListChangedEvent: Indicates the list of keychains has changed.
	KSecKeychainListChangedEvent SecKeychainEvent = 11
	// KSecLockEvent: Indicates a keychain was locked.
	KSecLockEvent SecKeychainEvent = 1
	// KSecPasswordChangedEvent: Indicates the keychain password was changed.
	KSecPasswordChangedEvent SecKeychainEvent = 6
	// KSecTrustSettingsChangedEvent: Indicates trust settings have changed.
	KSecTrustSettingsChangedEvent SecKeychainEvent = 12
	// KSecUnlockEvent: Indicates a keychain was successfully unlocked.
	KSecUnlockEvent SecKeychainEvent = 2
	// KSecUpdateEvent: Indicates a keychain item was updated.
	KSecUpdateEvent SecKeychainEvent = 5
	// Deprecated.
	KSecDataAccessEvent SecKeychainEvent = 10
)

func (e SecKeychainEvent) String() string {
	switch e {
	case KSecAddEvent:
		return "KSecAddEvent"
	case KSecDefaultChangedEvent:
		return "KSecDefaultChangedEvent"
	case KSecDeleteEvent:
		return "KSecDeleteEvent"
	case KSecKeychainListChangedEvent:
		return "KSecKeychainListChangedEvent"
	case KSecLockEvent:
		return "KSecLockEvent"
	case KSecPasswordChangedEvent:
		return "KSecPasswordChangedEvent"
	case KSecTrustSettingsChangedEvent:
		return "KSecTrustSettingsChangedEvent"
	case KSecUnlockEvent:
		return "KSecUnlockEvent"
	case KSecUpdateEvent:
		return "KSecUpdateEvent"
	case KSecDataAccessEvent:
		return "KSecDataAccessEvent"
	default:
		return fmt.Sprintf("SecKeychainEvent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecKeychainEventMask
type SecKeychainEventMask uint32

const (
	// KSecAddEventMask: If the bit specified by this mask is set, your callback function is invoked when an item is added to a keychain.
	KSecAddEventMask SecKeychainEventMask = 8
	// KSecDefaultChangedEventMask: If the bit specified by this mask is set, your callback function is invoked when a different keychain is specified as the default.
	KSecDefaultChangedEventMask SecKeychainEventMask = 512
	// KSecDeleteEventMask: If the bit specified by this mask is set, your callback function is invoked when an item is deleted from a keychain.
	KSecDeleteEventMask SecKeychainEventMask = 16
	// KSecEveryEventMask: If all the bits are set, your callback function is invoked whenever any event occurs.
	KSecEveryEventMask SecKeychainEventMask = 0xffffffff
	// KSecKeychainListChangedMask: If the bit specified by this mask is set, your callback function is invoked when a keychain list is changed.
	KSecKeychainListChangedMask SecKeychainEventMask = 2048
	// KSecLockEventMask: If the bit specified by this mask is set, your callback function is invoked when a keychain is locked.
	KSecLockEventMask SecKeychainEventMask = 2
	// KSecPasswordChangedEventMask: If the bit specified by this mask is set, your callback function is invoked when the keychain password is changed.
	KSecPasswordChangedEventMask SecKeychainEventMask = 64
	// KSecTrustSettingsChangedEventMask: If the bit specified by this mask is set, your callback function is invoked when there is a change in certificate trust settings.
	KSecTrustSettingsChangedEventMask SecKeychainEventMask = 4096
	// KSecUnlockEventMask: If the bit specified by this mask is set, your callback function is invoked when a keychain is unlocked.
	KSecUnlockEventMask SecKeychainEventMask = 4
	// KSecUpdateEventMask: If the bit specified by this mask is set, your callback function is invoked when a keychain item is updated.
	KSecUpdateEventMask SecKeychainEventMask = 32
	// Deprecated.
	KSecDataAccessEventMask SecKeychainEventMask = 1024
)

func (e SecKeychainEventMask) String() string {
	switch e {
	case KSecAddEventMask:
		return "KSecAddEventMask"
	case KSecDefaultChangedEventMask:
		return "KSecDefaultChangedEventMask"
	case KSecDeleteEventMask:
		return "KSecDeleteEventMask"
	case KSecEveryEventMask:
		return "KSecEveryEventMask"
	case KSecKeychainListChangedMask:
		return "KSecKeychainListChangedMask"
	case KSecLockEventMask:
		return "KSecLockEventMask"
	case KSecPasswordChangedEventMask:
		return "KSecPasswordChangedEventMask"
	case KSecTrustSettingsChangedEventMask:
		return "KSecTrustSettingsChangedEventMask"
	case KSecUnlockEventMask:
		return "KSecUnlockEventMask"
	case KSecUpdateEventMask:
		return "KSecUpdateEventMask"
	case KSecDataAccessEventMask:
		return "KSecDataAccessEventMask"
	default:
		return fmt.Sprintf("SecKeychainEventMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecKeychainPromptSelector
type SecKeychainPromptSelector uint16

const (
	// KSecKeychainPromptInvalid: Indicates that a passphrase should be required when an application with an invalid signature attempts to use the keychain, overriding the system default.
	KSecKeychainPromptInvalid SecKeychainPromptSelector = 0x40
	// KSecKeychainPromptInvalidAct: Indicates that a passphrase should be required when an application with an invalid signature attempts to use the keychain.
	KSecKeychainPromptInvalidAct SecKeychainPromptSelector = 0x80
	// KSecKeychainPromptRequirePassphase: Indicates that a passphrase should be required for every access.
	KSecKeychainPromptRequirePassphase SecKeychainPromptSelector = 0x1
	// KSecKeychainPromptUnsigned: Indicates that a passphrase should be required when an unsigned application attempts to use the keychain, overriding the system default.
	KSecKeychainPromptUnsigned SecKeychainPromptSelector = 0x10
	// KSecKeychainPromptUnsignedAct: Indicates that a passphrase should be required when an unsigned application attempts to use the keychain.
	KSecKeychainPromptUnsignedAct SecKeychainPromptSelector = 0x20
)

func (e SecKeychainPromptSelector) String() string {
	switch e {
	case KSecKeychainPromptInvalid:
		return "KSecKeychainPromptInvalid"
	case KSecKeychainPromptInvalidAct:
		return "KSecKeychainPromptInvalidAct"
	case KSecKeychainPromptRequirePassphase:
		return "KSecKeychainPromptRequirePassphase"
	case KSecKeychainPromptUnsigned:
		return "KSecKeychainPromptUnsigned"
	case KSecKeychainPromptUnsignedAct:
		return "KSecKeychainPromptUnsignedAct"
	default:
		return fmt.Sprintf("SecKeychainPromptSelector(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecPadding
type SecPadding uint32

const (
	// Deprecated.
	KSecPaddingNone SecPadding = 0
	// Deprecated.
	KSecPaddingOAEP SecPadding = 2
	// Deprecated.
	KSecPaddingPKCS1 SecPadding = 1
	// Deprecated.
	KSecPaddingPKCS1MD2 SecPadding = 0x8000
	// Deprecated.
	KSecPaddingPKCS1MD5 SecPadding = 0x8001
	// Deprecated.
	KSecPaddingPKCS1SHA1 SecPadding = 0x8002
	// Deprecated.
	KSecPaddingPKCS1SHA224 SecPadding = 0x8003
	// Deprecated.
	KSecPaddingPKCS1SHA256 SecPadding = 0x8004
	// Deprecated.
	KSecPaddingPKCS1SHA384 SecPadding = 0x8005
	// Deprecated.
	KSecPaddingPKCS1SHA512 SecPadding = 0x8006
	// Deprecated.
	KSecPaddingSigRaw SecPadding = 0x4000
)

func (e SecPadding) String() string {
	switch e {
	case KSecPaddingNone:
		return "KSecPaddingNone"
	case KSecPaddingOAEP:
		return "KSecPaddingOAEP"
	case KSecPaddingPKCS1:
		return "KSecPaddingPKCS1"
	case KSecPaddingPKCS1MD2:
		return "KSecPaddingPKCS1MD2"
	case KSecPaddingPKCS1MD5:
		return "KSecPaddingPKCS1MD5"
	case KSecPaddingPKCS1SHA1:
		return "KSecPaddingPKCS1SHA1"
	case KSecPaddingPKCS1SHA224:
		return "KSecPaddingPKCS1SHA224"
	case KSecPaddingPKCS1SHA256:
		return "KSecPaddingPKCS1SHA256"
	case KSecPaddingPKCS1SHA384:
		return "KSecPaddingPKCS1SHA384"
	case KSecPaddingPKCS1SHA512:
		return "KSecPaddingPKCS1SHA512"
	case KSecPaddingSigRaw:
		return "KSecPaddingSigRaw"
	default:
		return fmt.Sprintf("SecPadding(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecPreferencesDomain
type SecPreferencesDomain int32

const (
	// KSecPreferencesDomainCommon: Indicates the preferences are common to everyone.
	KSecPreferencesDomainCommon SecPreferencesDomain = 2
	// KSecPreferencesDomainDynamic: Indicates a dynamic search list (typically provided by removable keychains such as smart cards).
	KSecPreferencesDomainDynamic SecPreferencesDomain = 3
	// KSecPreferencesDomainSystem: Indicates the system or daemon preference domain preferences.
	KSecPreferencesDomainSystem SecPreferencesDomain = 1
	// KSecPreferencesDomainUser: Indicates the user preference domain preferences.
	KSecPreferencesDomainUser SecPreferencesDomain = 0
)

func (e SecPreferencesDomain) String() string {
	switch e {
	case KSecPreferencesDomainCommon:
		return "KSecPreferencesDomainCommon"
	case KSecPreferencesDomainDynamic:
		return "KSecPreferencesDomainDynamic"
	case KSecPreferencesDomainSystem:
		return "KSecPreferencesDomainSystem"
	case KSecPreferencesDomainUser:
		return "KSecPreferencesDomainUser"
	default:
		return fmt.Sprintf("SecPreferencesDomain(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecProtocolType
type SecProtocolType uint32

const (
	// KSecProtocolTypeAFP: Indicates AFP over TCP.
	KSecProtocolTypeAFP SecProtocolType = 'a'<<24 | 'f'<<16 | 'p'<<8 | ' ' // 'afp '
	// KSecProtocolTypeAny: Indicates that any protocol is acceptable.
	KSecProtocolTypeAny SecProtocolType = 0
	// KSecProtocolTypeAppleTalk: Indicates AFP over AppleTalk.
	KSecProtocolTypeAppleTalk SecProtocolType = 'a'<<24 | 't'<<16 | 'l'<<8 | 'k' // 'atlk'
	// KSecProtocolTypeCIFS: Indicates CIFS.
	KSecProtocolTypeCIFS SecProtocolType = 'c'<<24 | 'i'<<16 | 'f'<<8 | 's' // 'cifs'
	// KSecProtocolTypeCVSpserver: Indicates CVS pserver.
	KSecProtocolTypeCVSpserver SecProtocolType = 'c'<<24 | 'v'<<16 | 's'<<8 | 'p' // 'cvsp'
	// KSecProtocolTypeDAAP: Indicates DAAP.
	KSecProtocolTypeDAAP SecProtocolType = 'd'<<24 | 'a'<<16 | 'a'<<8 | 'p' // 'daap'
	// KSecProtocolTypeEPPC: Indicates Remote Apple Events.
	KSecProtocolTypeEPPC SecProtocolType = 'e'<<24 | 'p'<<16 | 'p'<<8 | 'c' // 'eppc'
	// KSecProtocolTypeFTP: Indicates FTP.
	KSecProtocolTypeFTP SecProtocolType = 'f'<<24 | 't'<<16 | 'p'<<8 | ' ' // 'ftp '
	// KSecProtocolTypeFTPAccount: Indicates a client side FTP account.
	KSecProtocolTypeFTPAccount SecProtocolType = 'f'<<24 | 't'<<16 | 'p'<<8 | 'a' // 'ftpa'
	// KSecProtocolTypeFTPProxy: Indicates FTP proxy.
	KSecProtocolTypeFTPProxy SecProtocolType = 'f'<<24 | 't'<<16 | 'p'<<8 | 'x' // 'ftpx'
	// KSecProtocolTypeFTPS: Indicates FTP over TLS/SSL.
	KSecProtocolTypeFTPS SecProtocolType = 'f'<<24 | 't'<<16 | 'p'<<8 | 's' // 'ftps'
	// KSecProtocolTypeHTTP: Indicates HTTP.
	KSecProtocolTypeHTTP SecProtocolType = 'h'<<24 | 't'<<16 | 't'<<8 | 'p' // 'http'
	// KSecProtocolTypeHTTPProxy: Indicates HTTP proxy.
	KSecProtocolTypeHTTPProxy SecProtocolType = 'h'<<24 | 't'<<16 | 'p'<<8 | 'x' // 'htpx'
	// KSecProtocolTypeHTTPS: Indicates HTTP over TLS/SSL.
	KSecProtocolTypeHTTPS SecProtocolType = 'h'<<24 | 't'<<16 | 'p'<<8 | 's' // 'htps'
	// KSecProtocolTypeHTTPSProxy: Indicates HTTPS proxy.
	KSecProtocolTypeHTTPSProxy SecProtocolType = 'h'<<24 | 't'<<16 | 's'<<8 | 'x' // 'htsx'
	// KSecProtocolTypeIMAP: Indicates IMAP.
	KSecProtocolTypeIMAP SecProtocolType = 'i'<<24 | 'm'<<16 | 'a'<<8 | 'p' // 'imap'
	// KSecProtocolTypeIMAPS: Indicates IMAP4 over TLS/SSL.
	KSecProtocolTypeIMAPS SecProtocolType = 'i'<<24 | 'm'<<16 | 'p'<<8 | 's' // 'imps'
	// KSecProtocolTypeIPP: Indicates IPP.
	KSecProtocolTypeIPP SecProtocolType = 'i'<<24 | 'p'<<16 | 'p'<<8 | ' ' // 'ipp '
	// KSecProtocolTypeIRC: Indicates IRC.
	KSecProtocolTypeIRC SecProtocolType = 'i'<<24 | 'r'<<16 | 'c'<<8 | ' ' // 'irc '
	// KSecProtocolTypeIRCS: Indicates IRC over TLS/SSL.
	KSecProtocolTypeIRCS SecProtocolType = 'i'<<24 | 'r'<<16 | 'c'<<8 | 's' // 'ircs'
	// KSecProtocolTypeLDAP: Indicates LDAP.
	KSecProtocolTypeLDAP SecProtocolType = 'l'<<24 | 'd'<<16 | 'a'<<8 | 'p' // 'ldap'
	// KSecProtocolTypeLDAPS: Indicates LDAP over TLS/SSL.
	KSecProtocolTypeLDAPS SecProtocolType = 'l'<<24 | 'd'<<16 | 'p'<<8 | 's' // 'ldps'
	// KSecProtocolTypeNNTP: Indicates NNTP.
	KSecProtocolTypeNNTP SecProtocolType = 'n'<<24 | 'n'<<16 | 't'<<8 | 'p' // 'nntp'
	// KSecProtocolTypeNNTPS: Indicates NNTP over TLS/SSL.
	KSecProtocolTypeNNTPS SecProtocolType = 'n'<<24 | 't'<<16 | 'p'<<8 | 's' // 'ntps'
	// KSecProtocolTypePOP3: Indicates POP3.
	KSecProtocolTypePOP3 SecProtocolType = 'p'<<24 | 'o'<<16 | 'p'<<8 | '3' // 'pop3'
	// KSecProtocolTypePOP3S: Indicates POP3 over TLS/SSL.
	KSecProtocolTypePOP3S SecProtocolType = 'p'<<24 | 'o'<<16 | 'p'<<8 | 's' // 'pops'
	// KSecProtocolTypeRTSP: Indicates RTSP.
	KSecProtocolTypeRTSP SecProtocolType = 'r'<<24 | 't'<<16 | 's'<<8 | 'p' // 'rtsp'
	// KSecProtocolTypeRTSPProxy: Indicates RTSP proxy.
	KSecProtocolTypeRTSPProxy SecProtocolType = 'r'<<24 | 't'<<16 | 's'<<8 | 'x' // 'rtsx'
	// KSecProtocolTypeSMB: Indicates SMB.
	KSecProtocolTypeSMB SecProtocolType = 's'<<24 | 'm'<<16 | 'b'<<8 | ' ' // 'smb '
	// KSecProtocolTypeSMTP: Indicates SMTP.
	KSecProtocolTypeSMTP SecProtocolType = 's'<<24 | 'm'<<16 | 't'<<8 | 'p' // 'smtp'
	// KSecProtocolTypeSOCKS: Indicates SOCKS.
	KSecProtocolTypeSOCKS SecProtocolType = 's'<<24 | 'o'<<16 | 'x'<<8 | ' ' // 'sox '
	// KSecProtocolTypeSSH: Indicates SSH.
	KSecProtocolTypeSSH SecProtocolType = 's'<<24 | 's'<<16 | 'h'<<8 | ' ' // 'ssh '
	// KSecProtocolTypeSVN: Indicates Subversion.
	KSecProtocolTypeSVN SecProtocolType = 's'<<24 | 'v'<<16 | 'n'<<8 | ' ' // 'svn '
	// KSecProtocolTypeTelnet: Indicates Telnet.
	KSecProtocolTypeTelnet SecProtocolType = 't'<<24 | 'e'<<16 | 'l'<<8 | 'n' // 'teln'
	// KSecProtocolTypeTelnetS: Indicates Telnet over TLS/SSL.
	KSecProtocolTypeTelnetS SecProtocolType = 't'<<24 | 'e'<<16 | 'l'<<8 | 's' // 'tels'
)

func (e SecProtocolType) String() string {
	switch e {
	case KSecProtocolTypeAFP:
		return "KSecProtocolTypeAFP"
	case KSecProtocolTypeAny:
		return "KSecProtocolTypeAny"
	case KSecProtocolTypeAppleTalk:
		return "KSecProtocolTypeAppleTalk"
	case KSecProtocolTypeCIFS:
		return "KSecProtocolTypeCIFS"
	case KSecProtocolTypeCVSpserver:
		return "KSecProtocolTypeCVSpserver"
	case KSecProtocolTypeDAAP:
		return "KSecProtocolTypeDAAP"
	case KSecProtocolTypeEPPC:
		return "KSecProtocolTypeEPPC"
	case KSecProtocolTypeFTP:
		return "KSecProtocolTypeFTP"
	case KSecProtocolTypeFTPAccount:
		return "KSecProtocolTypeFTPAccount"
	case KSecProtocolTypeFTPProxy:
		return "KSecProtocolTypeFTPProxy"
	case KSecProtocolTypeFTPS:
		return "KSecProtocolTypeFTPS"
	case KSecProtocolTypeHTTP:
		return "KSecProtocolTypeHTTP"
	case KSecProtocolTypeHTTPProxy:
		return "KSecProtocolTypeHTTPProxy"
	case KSecProtocolTypeHTTPS:
		return "KSecProtocolTypeHTTPS"
	case KSecProtocolTypeHTTPSProxy:
		return "KSecProtocolTypeHTTPSProxy"
	case KSecProtocolTypeIMAP:
		return "KSecProtocolTypeIMAP"
	case KSecProtocolTypeIMAPS:
		return "KSecProtocolTypeIMAPS"
	case KSecProtocolTypeIPP:
		return "KSecProtocolTypeIPP"
	case KSecProtocolTypeIRC:
		return "KSecProtocolTypeIRC"
	case KSecProtocolTypeIRCS:
		return "KSecProtocolTypeIRCS"
	case KSecProtocolTypeLDAP:
		return "KSecProtocolTypeLDAP"
	case KSecProtocolTypeLDAPS:
		return "KSecProtocolTypeLDAPS"
	case KSecProtocolTypeNNTP:
		return "KSecProtocolTypeNNTP"
	case KSecProtocolTypeNNTPS:
		return "KSecProtocolTypeNNTPS"
	case KSecProtocolTypePOP3:
		return "KSecProtocolTypePOP3"
	case KSecProtocolTypePOP3S:
		return "KSecProtocolTypePOP3S"
	case KSecProtocolTypeRTSP:
		return "KSecProtocolTypeRTSP"
	case KSecProtocolTypeRTSPProxy:
		return "KSecProtocolTypeRTSPProxy"
	case KSecProtocolTypeSMB:
		return "KSecProtocolTypeSMB"
	case KSecProtocolTypeSMTP:
		return "KSecProtocolTypeSMTP"
	case KSecProtocolTypeSOCKS:
		return "KSecProtocolTypeSOCKS"
	case KSecProtocolTypeSSH:
		return "KSecProtocolTypeSSH"
	case KSecProtocolTypeSVN:
		return "KSecProtocolTypeSVN"
	case KSecProtocolTypeTelnet:
		return "KSecProtocolTypeTelnet"
	case KSecProtocolTypeTelnetS:
		return "KSecProtocolTypeTelnetS"
	default:
		return fmt.Sprintf("SecProtocolType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecRequirementType
type SecRequirementType uint32

const (
	// KSecDesignatedRequirementType: A designated requirement.
	KSecDesignatedRequirementType SecRequirementType = 3
	// KSecGuestRequirementType: What guests this code may run.
	KSecGuestRequirementType SecRequirementType = 2
	// KSecHostRequirementType: What hosts may run this code.
	KSecHostRequirementType SecRequirementType = 1
	// KSecInvalidRequirementType: Invalid type of requirement.
	KSecInvalidRequirementType SecRequirementType = 6
	// KSecLibraryRequirementType: What libraries this code may link against.
	KSecLibraryRequirementType SecRequirementType = 4
	// KSecPluginRequirementType: What plug-ins this code may load.
	KSecPluginRequirementType SecRequirementType = 5
	// KSecRequirementTypeCount: The number of valid requirement types.
	KSecRequirementTypeCount SecRequirementType = 6
)

func (e SecRequirementType) String() string {
	switch e {
	case KSecDesignatedRequirementType:
		return "KSecDesignatedRequirementType"
	case KSecGuestRequirementType:
		return "KSecGuestRequirementType"
	case KSecHostRequirementType:
		return "KSecHostRequirementType"
	case KSecInvalidRequirementType:
		return "KSecInvalidRequirementType"
	case KSecLibraryRequirementType:
		return "KSecLibraryRequirementType"
	case KSecPluginRequirementType:
		return "KSecPluginRequirementType"
	default:
		return fmt.Sprintf("SecRequirementType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType
type SecTransformMetaAttributeType int

const (
	// Deprecated.
	KSecTransformMetaAttributeCanCycle SecTransformMetaAttributeType = 7
	// Deprecated.
	KSecTransformMetaAttributeDeferred SecTransformMetaAttributeType = 5
	// Deprecated.
	KSecTransformMetaAttributeExternalize SecTransformMetaAttributeType = 8
	// Deprecated.
	KSecTransformMetaAttributeHasInboundConnection SecTransformMetaAttributeType = 10
	// Deprecated.
	KSecTransformMetaAttributeHasOutboundConnections SecTransformMetaAttributeType = 9
	// Deprecated.
	KSecTransformMetaAttributeName SecTransformMetaAttributeType = 1
	// Deprecated.
	KSecTransformMetaAttributeRef SecTransformMetaAttributeType = 2
	// Deprecated.
	KSecTransformMetaAttributeRequired SecTransformMetaAttributeType = 3
	// Deprecated.
	KSecTransformMetaAttributeRequiresOutboundConnection SecTransformMetaAttributeType = 4
	// Deprecated.
	KSecTransformMetaAttributeStream SecTransformMetaAttributeType = 6
	// Deprecated.
	KSecTransformMetaAttributeValue SecTransformMetaAttributeType = 0
)

func (e SecTransformMetaAttributeType) String() string {
	switch e {
	case KSecTransformMetaAttributeCanCycle:
		return "KSecTransformMetaAttributeCanCycle"
	case KSecTransformMetaAttributeDeferred:
		return "KSecTransformMetaAttributeDeferred"
	case KSecTransformMetaAttributeExternalize:
		return "KSecTransformMetaAttributeExternalize"
	case KSecTransformMetaAttributeHasInboundConnection:
		return "KSecTransformMetaAttributeHasInboundConnection"
	case KSecTransformMetaAttributeHasOutboundConnections:
		return "KSecTransformMetaAttributeHasOutboundConnections"
	case KSecTransformMetaAttributeName:
		return "KSecTransformMetaAttributeName"
	case KSecTransformMetaAttributeRef:
		return "KSecTransformMetaAttributeRef"
	case KSecTransformMetaAttributeRequired:
		return "KSecTransformMetaAttributeRequired"
	case KSecTransformMetaAttributeRequiresOutboundConnection:
		return "KSecTransformMetaAttributeRequiresOutboundConnection"
	case KSecTransformMetaAttributeStream:
		return "KSecTransformMetaAttributeStream"
	case KSecTransformMetaAttributeValue:
		return "KSecTransformMetaAttributeValue"
	default:
		return fmt.Sprintf("SecTransformMetaAttributeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecTrustOptionFlags
type SecTrustOptionFlags uint32

const (
	// KSecTrustOptionAllowExpired: Allow expired certificates (except for the root certificate).
	KSecTrustOptionAllowExpired SecTrustOptionFlags = 0x1
	// KSecTrustOptionAllowExpiredRoot: Allow expired root certificates.
	KSecTrustOptionAllowExpiredRoot SecTrustOptionFlags = 0x8
	// KSecTrustOptionFetchIssuerFromNet: Allow network downloads of CA certificates.
	KSecTrustOptionFetchIssuerFromNet SecTrustOptionFlags = 0x4
	// KSecTrustOptionImplicitAnchors: Treat properly self-signed certificates as anchors implicitly.
	KSecTrustOptionImplicitAnchors SecTrustOptionFlags = 0x40
	// KSecTrustOptionLeafIsCA: Allow CA certificates as leaf certificates.
	KSecTrustOptionLeafIsCA SecTrustOptionFlags = 0x2
	// KSecTrustOptionRequireRevPerCert: Require a positive revocation check for each certificate.
	KSecTrustOptionRequireRevPerCert SecTrustOptionFlags = 0x10
	// KSecTrustOptionUseTrustSettings: Use TrustSettings instead of anchors.
	KSecTrustOptionUseTrustSettings SecTrustOptionFlags = 0x20
)

func (e SecTrustOptionFlags) String() string {
	switch e {
	case KSecTrustOptionAllowExpired:
		return "KSecTrustOptionAllowExpired"
	case KSecTrustOptionAllowExpiredRoot:
		return "KSecTrustOptionAllowExpiredRoot"
	case KSecTrustOptionFetchIssuerFromNet:
		return "KSecTrustOptionFetchIssuerFromNet"
	case KSecTrustOptionImplicitAnchors:
		return "KSecTrustOptionImplicitAnchors"
	case KSecTrustOptionLeafIsCA:
		return "KSecTrustOptionLeafIsCA"
	case KSecTrustOptionRequireRevPerCert:
		return "KSecTrustOptionRequireRevPerCert"
	case KSecTrustOptionUseTrustSettings:
		return "KSecTrustOptionUseTrustSettings"
	default:
		return fmt.Sprintf("SecTrustOptionFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecTrustResultType
type SecTrustResultType uint32

const (
	// KSecTrustResultDeny: The user specified that the certificate should not be trusted.
	KSecTrustResultDeny SecTrustResultType = 3
	// KSecTrustResultFatalTrustFailure: Trust is denied and no simple fix is available.
	KSecTrustResultFatalTrustFailure SecTrustResultType = 6
	// KSecTrustResultInvalid: An indication of an invalid setting or result.
	KSecTrustResultInvalid SecTrustResultType = 0
	// KSecTrustResultOtherError: A value that indicates a failure other than trust evaluation.
	KSecTrustResultOtherError SecTrustResultType = 7
	// KSecTrustResultProceed: The user granted permission to trust the certificate for the purposes designated in the specified policies.
	KSecTrustResultProceed SecTrustResultType = 1
	// KSecTrustResultRecoverableTrustFailure: Trust is denied, but recovery may be possible.
	KSecTrustResultRecoverableTrustFailure SecTrustResultType = 5
	// KSecTrustResultUnspecified: The user did not specify a trust setting.
	KSecTrustResultUnspecified SecTrustResultType = 4
	// Deprecated.
	KSecTrustResultConfirm SecTrustResultType = 2
)

func (e SecTrustResultType) String() string {
	switch e {
	case KSecTrustResultDeny:
		return "KSecTrustResultDeny"
	case KSecTrustResultFatalTrustFailure:
		return "KSecTrustResultFatalTrustFailure"
	case KSecTrustResultInvalid:
		return "KSecTrustResultInvalid"
	case KSecTrustResultOtherError:
		return "KSecTrustResultOtherError"
	case KSecTrustResultProceed:
		return "KSecTrustResultProceed"
	case KSecTrustResultRecoverableTrustFailure:
		return "KSecTrustResultRecoverableTrustFailure"
	case KSecTrustResultUnspecified:
		return "KSecTrustResultUnspecified"
	case KSecTrustResultConfirm:
		return "KSecTrustResultConfirm"
	default:
		return fmt.Sprintf("SecTrustResultType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecTrustSettingsDomain
type SecTrustSettingsDomain uint32

const (
	// KSecTrustSettingsDomainAdmin: Locally administered, system-wide trust settings.
	KSecTrustSettingsDomainAdmin SecTrustSettingsDomain = 1
	// KSecTrustSettingsDomainSystem: System trust settings.
	KSecTrustSettingsDomainSystem SecTrustSettingsDomain = 2
	// KSecTrustSettingsDomainUser: Per-user trust settings.
	KSecTrustSettingsDomainUser SecTrustSettingsDomain = 0
)

func (e SecTrustSettingsDomain) String() string {
	switch e {
	case KSecTrustSettingsDomainAdmin:
		return "KSecTrustSettingsDomainAdmin"
	case KSecTrustSettingsDomainSystem:
		return "KSecTrustSettingsDomainSystem"
	case KSecTrustSettingsDomainUser:
		return "KSecTrustSettingsDomainUser"
	default:
		return fmt.Sprintf("SecTrustSettingsDomain(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecTrustSettingsKeyUsage
type SecTrustSettingsKeyUsage uint32

const (
	// KSecTrustSettingsKeyUseAny: The key can be used for any purpose.
	KSecTrustSettingsKeyUseAny SecTrustSettingsKeyUsage = 0xffffffff
	// KSecTrustSettingsKeyUseEnDecryptData: The key can be used to encrypt or decrypt data.
	KSecTrustSettingsKeyUseEnDecryptData SecTrustSettingsKeyUsage = 0x2
	// KSecTrustSettingsKeyUseEnDecryptKey: The key can be used to encrypt or decrypt (wrap or unwrap) a key.
	KSecTrustSettingsKeyUseEnDecryptKey SecTrustSettingsKeyUsage = 0x4
	// KSecTrustSettingsKeyUseKeyExchange: The key is a private key that has been shared using a key exchange protocol, such as Diffie-Hellman key exchange.
	KSecTrustSettingsKeyUseKeyExchange SecTrustSettingsKeyUsage = 0x20
	// KSecTrustSettingsKeyUseSignCert: The key can be used to sign a certificate or verify a signature.
	KSecTrustSettingsKeyUseSignCert SecTrustSettingsKeyUsage = 0x8
	// KSecTrustSettingsKeyUseSignRevocation: The key can be used to sign an OCSP (online certificate status protocol) message or CRL (certificate verification list), or to verify a signature.
	KSecTrustSettingsKeyUseSignRevocation SecTrustSettingsKeyUsage = 0x10
	// KSecTrustSettingsKeyUseSignature: The key can be used to sign data or verify a signature.
	KSecTrustSettingsKeyUseSignature SecTrustSettingsKeyUsage = 0x1
)

func (e SecTrustSettingsKeyUsage) String() string {
	switch e {
	case KSecTrustSettingsKeyUseAny:
		return "KSecTrustSettingsKeyUseAny"
	case KSecTrustSettingsKeyUseEnDecryptData:
		return "KSecTrustSettingsKeyUseEnDecryptData"
	case KSecTrustSettingsKeyUseEnDecryptKey:
		return "KSecTrustSettingsKeyUseEnDecryptKey"
	case KSecTrustSettingsKeyUseKeyExchange:
		return "KSecTrustSettingsKeyUseKeyExchange"
	case KSecTrustSettingsKeyUseSignCert:
		return "KSecTrustSettingsKeyUseSignCert"
	case KSecTrustSettingsKeyUseSignRevocation:
		return "KSecTrustSettingsKeyUseSignRevocation"
	case KSecTrustSettingsKeyUseSignature:
		return "KSecTrustSettingsKeyUseSignature"
	default:
		return fmt.Sprintf("SecTrustSettingsKeyUsage(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecTrustSettingsResult
type SecTrustSettingsResult uint32

const (
	// KSecTrustSettingsResultDeny: This certificate is explicitly distrusted.
	KSecTrustSettingsResultDeny SecTrustSettingsResult = 3
	// KSecTrustSettingsResultInvalid: Never valid in a trust settings array or in an API call.
	KSecTrustSettingsResultInvalid SecTrustSettingsResult = 0
	// KSecTrustSettingsResultTrustAsRoot: This non-root certificate is explicitly trusted as if it were a trusted root.
	KSecTrustSettingsResultTrustAsRoot SecTrustSettingsResult = 2
	// KSecTrustSettingsResultTrustRoot: This root certificate is explicitly trusted.
	KSecTrustSettingsResultTrustRoot SecTrustSettingsResult = 1
	// KSecTrustSettingsResultUnspecified: This certificate is neither trusted nor distrusted.
	KSecTrustSettingsResultUnspecified SecTrustSettingsResult = 4
)

func (e SecTrustSettingsResult) String() string {
	switch e {
	case KSecTrustSettingsResultDeny:
		return "KSecTrustSettingsResultDeny"
	case KSecTrustSettingsResultInvalid:
		return "KSecTrustSettingsResultInvalid"
	case KSecTrustSettingsResultTrustAsRoot:
		return "KSecTrustSettingsResultTrustAsRoot"
	case KSecTrustSettingsResultTrustRoot:
		return "KSecTrustSettingsResultTrustRoot"
	case KSecTrustSettingsResultUnspecified:
		return "KSecTrustSettingsResultUnspecified"
	default:
		return fmt.Sprintf("SecTrustSettingsResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SessionAttributeBits
type SessionAttributeBits uint32

const (
	// SessionHasGraphicAccess: A bit that indicates a graphic subsystem is available.
	SessionHasGraphicAccess SessionAttributeBits = 0x10
	// SessionHasTTY: A bit that indicates `/dev/tty` is available.
	SessionHasTTY SessionAttributeBits = 0x20
	// SessionIsRemote: A bit that indicates the session was initiated over the network.
	SessionIsRemote SessionAttributeBits = 0x1000
	// SessionIsRoot: A bit that indicates the session is the root session.
	SessionIsRoot SessionAttributeBits = 0x1
)

func (e SessionAttributeBits) String() string {
	switch e {
	case SessionHasGraphicAccess:
		return "SessionHasGraphicAccess"
	case SessionHasTTY:
		return "SessionHasTTY"
	case SessionIsRemote:
		return "SessionIsRemote"
	case SessionIsRoot:
		return "SessionIsRoot"
	default:
		return fmt.Sprintf("SessionAttributeBits(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SessionCreationFlags
type SessionCreationFlags uint32

const (
	// SessionKeepCurrentBootstrap: The caller has allocated sub-bootstrap.
	SessionKeepCurrentBootstrap SessionCreationFlags = 0x8000
)

func (e SessionCreationFlags) String() string {
	switch e {
	case SessionKeepCurrentBootstrap:
		return "SessionKeepCurrentBootstrap"
	default:
		return fmt.Sprintf("SessionCreationFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/SecureDownloadTrustCallbackResult
type SecureDownloadTrustCallbackResult uint32

const (
	KSecureDownloadDoNotEvaluateSigner SecureDownloadTrustCallbackResult = 0
	KSecureDownloadEvaluateSigner      SecureDownloadTrustCallbackResult = 1
	KSecureDownloadFailEvaluation      SecureDownloadTrustCallbackResult = 2
)

func (e SecureDownloadTrustCallbackResult) String() string {
	switch e {
	case KSecureDownloadDoNotEvaluateSigner:
		return "KSecureDownloadDoNotEvaluateSigner"
	case KSecureDownloadEvaluateSigner:
		return "KSecureDownloadEvaluateSigner"
	case KSecureDownloadFailEvaluation:
		return "KSecureDownloadFailEvaluation"
	default:
		return fmt.Sprintf("SecureDownloadTrustCallbackResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/CE_CrlDistributionPointNameType
type CE_CrlDistributionPointNameType uint32

const (
	// Deprecated.
	CE_CDNT_FullName CE_CrlDistributionPointNameType = 0
	// Deprecated.
	CE_CDNT_NameRelativeToCrlIssuer CE_CrlDistributionPointNameType = 1
)

func (e CE_CrlDistributionPointNameType) String() string {
	switch e {
	case CE_CDNT_FullName:
		return "CE_CDNT_FullName"
	case CE_CDNT_NameRelativeToCrlIssuer:
		return "CE_CDNT_NameRelativeToCrlIssuer"
	default:
		return fmt.Sprintf("CE_CrlDistributionPointNameType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/CE_DataType-c.enum
type CE_DataType uint32

const (
	DT_AuthorityInfoAccess      CE_DataType = 14
	DT_AuthorityKeyID           CE_DataType = 0
	DT_BasicConstraints         CE_DataType = 6
	DT_CertPolicies             CE_DataType = 7
	DT_CrlDistributionPoints    CE_DataType = 12
	DT_CrlNumber                CE_DataType = 9
	DT_CrlReason                CE_DataType = 11
	DT_DeltaCrl                 CE_DataType = 10
	DT_ExtendedKeyUsage         CE_DataType = 5
	DT_InhibitAnyPolicy         CE_DataType = 20
	DT_IssuerAltName            CE_DataType = 4
	DT_IssuingDistributionPoint CE_DataType = 13
	DT_KeyUsage                 CE_DataType = 2
	DT_NameConstraints          CE_DataType = 17
	DT_NetscapeCertType         CE_DataType = 8
	DT_Other                    CE_DataType = 15
	DT_PolicyConstraints        CE_DataType = 19
	DT_PolicyMappings           CE_DataType = 18
	DT_QC_Statements            CE_DataType = 16
	DT_SubjectAltName           CE_DataType = 3
	DT_SubjectKeyID             CE_DataType = 1
)

func (e CE_DataType) String() string {
	switch e {
	case DT_AuthorityInfoAccess:
		return "DT_AuthorityInfoAccess"
	case DT_AuthorityKeyID:
		return "DT_AuthorityKeyID"
	case DT_BasicConstraints:
		return "DT_BasicConstraints"
	case DT_CertPolicies:
		return "DT_CertPolicies"
	case DT_CrlDistributionPoints:
		return "DT_CrlDistributionPoints"
	case DT_CrlNumber:
		return "DT_CrlNumber"
	case DT_CrlReason:
		return "DT_CrlReason"
	case DT_DeltaCrl:
		return "DT_DeltaCrl"
	case DT_ExtendedKeyUsage:
		return "DT_ExtendedKeyUsage"
	case DT_InhibitAnyPolicy:
		return "DT_InhibitAnyPolicy"
	case DT_IssuerAltName:
		return "DT_IssuerAltName"
	case DT_IssuingDistributionPoint:
		return "DT_IssuingDistributionPoint"
	case DT_KeyUsage:
		return "DT_KeyUsage"
	case DT_NameConstraints:
		return "DT_NameConstraints"
	case DT_NetscapeCertType:
		return "DT_NetscapeCertType"
	case DT_Other:
		return "DT_Other"
	case DT_PolicyConstraints:
		return "DT_PolicyConstraints"
	case DT_PolicyMappings:
		return "DT_PolicyMappings"
	case DT_QC_Statements:
		return "DT_QC_Statements"
	case DT_SubjectAltName:
		return "DT_SubjectAltName"
	case DT_SubjectKeyID:
		return "DT_SubjectKeyID"
	default:
		return fmt.Sprintf("CE_DataType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/CE_GeneralNameType-c.enum
type CE_GeneralNameType int

const (
	GNT_DNSName       CE_GeneralNameType = 2
	GNT_DirectoryName CE_GeneralNameType = 4
	GNT_EdiPartyName  CE_GeneralNameType = 5
	GNT_IPAddress     CE_GeneralNameType = 7
	GNT_OtherName     CE_GeneralNameType = 0
	GNT_RFC822Name    CE_GeneralNameType = 1
	GNT_RegisteredID  CE_GeneralNameType = 8
	GNT_URI           CE_GeneralNameType = 6
	GNT_X400Address   CE_GeneralNameType = 3
)

func (e CE_GeneralNameType) String() string {
	switch e {
	case GNT_DNSName:
		return "GNT_DNSName"
	case GNT_DirectoryName:
		return "GNT_DirectoryName"
	case GNT_EdiPartyName:
		return "GNT_EdiPartyName"
	case GNT_IPAddress:
		return "GNT_IPAddress"
	case GNT_OtherName:
		return "GNT_OtherName"
	case GNT_RFC822Name:
		return "GNT_RFC822Name"
	case GNT_RegisteredID:
		return "GNT_RegisteredID"
	case GNT_URI:
		return "GNT_URI"
	case GNT_X400Address:
		return "GNT_X400Address"
	default:
		return fmt.Sprintf("CE_GeneralNameType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/cssm_appledl_open_parameters_mask
type Cssm_appledl_open_parameters_mask uint32

const (
	KCSSM_APPLEDL_MASK_MODE Cssm_appledl_open_parameters_mask = 1
)

func (e Cssm_appledl_open_parameters_mask) String() string {
	switch e {
	case KCSSM_APPLEDL_MASK_MODE:
		return "KCSSM_APPLEDL_MASK_MODE"
	default:
		return fmt.Sprintf("Cssm_appledl_open_parameters_mask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/extension_data_format
type Extension_data_format uint32

const (
	CSSM_X509_DATAFORMAT_ENCODED Extension_data_format = 0
	CSSM_X509_DATAFORMAT_PAIR    Extension_data_format = 2
	CSSM_X509_DATAFORMAT_PARSED  Extension_data_format = 1
)

func (e Extension_data_format) String() string {
	switch e {
	case CSSM_X509_DATAFORMAT_ENCODED:
		return "CSSM_X509_DATAFORMAT_ENCODED"
	case CSSM_X509_DATAFORMAT_PAIR:
		return "CSSM_X509_DATAFORMAT_PAIR"
	case CSSM_X509_DATAFORMAT_PARSED:
		return "CSSM_X509_DATAFORMAT_PARSED"
	default:
		return fmt.Sprintf("Extension_data_format(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/tls_ciphersuite_group_t
type Tls_ciphersuite_group_t uint16

const (
	Tls_ciphersuite_group_ats               Tls_ciphersuite_group_t = 3
	Tls_ciphersuite_group_ats_compatibility Tls_ciphersuite_group_t = 4
	Tls_ciphersuite_group_ats_fcp_v2_1      Tls_ciphersuite_group_t = 5
	Tls_ciphersuite_group_compatibility     Tls_ciphersuite_group_t = 1
	Tls_ciphersuite_group_default           Tls_ciphersuite_group_t = 0
	Tls_ciphersuite_group_legacy            Tls_ciphersuite_group_t = 2
)

func (e Tls_ciphersuite_group_t) String() string {
	switch e {
	case Tls_ciphersuite_group_ats:
		return "Tls_ciphersuite_group_ats"
	case Tls_ciphersuite_group_ats_compatibility:
		return "Tls_ciphersuite_group_ats_compatibility"
	case Tls_ciphersuite_group_ats_fcp_v2_1:
		return "Tls_ciphersuite_group_ats_fcp_v2_1"
	case Tls_ciphersuite_group_compatibility:
		return "Tls_ciphersuite_group_compatibility"
	case Tls_ciphersuite_group_default:
		return "Tls_ciphersuite_group_default"
	case Tls_ciphersuite_group_legacy:
		return "Tls_ciphersuite_group_legacy"
	default:
		return fmt.Sprintf("Tls_ciphersuite_group_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/tls_ciphersuite_t
type Tls_ciphersuite_t uint16

const (
	Tls_ciphersuite_AES_128_GCM_SHA256                        Tls_ciphersuite_t = 0x1301
	Tls_ciphersuite_AES_256_GCM_SHA384                        Tls_ciphersuite_t = 0x1302
	Tls_ciphersuite_CHACHA20_POLY1305_SHA256                  Tls_ciphersuite_t = 0x1303
	Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_CBC_SHA          Tls_ciphersuite_t = 0xc009
	Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256       Tls_ciphersuite_t = 0xc023
	Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256       Tls_ciphersuite_t = 0xc02b
	Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_CBC_SHA          Tls_ciphersuite_t = 0xc00a
	Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384       Tls_ciphersuite_t = 0xc024
	Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384       Tls_ciphersuite_t = 0xc02c
	Tls_ciphersuite_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 Tls_ciphersuite_t = 0xcca9
	Tls_ciphersuite_ECDHE_RSA_WITH_AES_128_CBC_SHA            Tls_ciphersuite_t = 0xc013
	Tls_ciphersuite_ECDHE_RSA_WITH_AES_128_CBC_SHA256         Tls_ciphersuite_t = 0xc027
	Tls_ciphersuite_ECDHE_RSA_WITH_AES_128_GCM_SHA256         Tls_ciphersuite_t = 0xc02f
	Tls_ciphersuite_ECDHE_RSA_WITH_AES_256_CBC_SHA            Tls_ciphersuite_t = 0xc014
	Tls_ciphersuite_ECDHE_RSA_WITH_AES_256_CBC_SHA384         Tls_ciphersuite_t = 0xc028
	Tls_ciphersuite_ECDHE_RSA_WITH_AES_256_GCM_SHA384         Tls_ciphersuite_t = 0xc030
	Tls_ciphersuite_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256   Tls_ciphersuite_t = 0xcca8
	Tls_ciphersuite_RSA_WITH_AES_128_CBC_SHA                  Tls_ciphersuite_t = 0x2f
	Tls_ciphersuite_RSA_WITH_AES_128_CBC_SHA256               Tls_ciphersuite_t = 0x3c
	Tls_ciphersuite_RSA_WITH_AES_128_GCM_SHA256               Tls_ciphersuite_t = 0x9c
	Tls_ciphersuite_RSA_WITH_AES_256_CBC_SHA                  Tls_ciphersuite_t = 0x35
	Tls_ciphersuite_RSA_WITH_AES_256_CBC_SHA256               Tls_ciphersuite_t = 0x3d
	Tls_ciphersuite_RSA_WITH_AES_256_GCM_SHA384               Tls_ciphersuite_t = 0x9d
	// Deprecated.
	Tls_ciphersuite_ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA Tls_ciphersuite_t = 0xc008
	// Deprecated.
	Tls_ciphersuite_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA Tls_ciphersuite_t = 0xc012
	// Deprecated.
	Tls_ciphersuite_RSA_WITH_3DES_EDE_CBC_SHA Tls_ciphersuite_t = 0xa
)

func (e Tls_ciphersuite_t) String() string {
	switch e {
	case Tls_ciphersuite_AES_128_GCM_SHA256:
		return "Tls_ciphersuite_AES_128_GCM_SHA256"
	case Tls_ciphersuite_AES_256_GCM_SHA384:
		return "Tls_ciphersuite_AES_256_GCM_SHA384"
	case Tls_ciphersuite_CHACHA20_POLY1305_SHA256:
		return "Tls_ciphersuite_CHACHA20_POLY1305_SHA256"
	case Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:
		return "Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_CBC_SHA"
	case Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256:
		return "Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256"
	case Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:
		return "Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
	case Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:
		return "Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_CBC_SHA"
	case Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384:
		return "Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384"
	case Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:
		return "Tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
	case Tls_ciphersuite_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256:
		return "Tls_ciphersuite_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256"
	case Tls_ciphersuite_ECDHE_RSA_WITH_AES_128_CBC_SHA:
		return "Tls_ciphersuite_ECDHE_RSA_WITH_AES_128_CBC_SHA"
	case Tls_ciphersuite_ECDHE_RSA_WITH_AES_128_CBC_SHA256:
		return "Tls_ciphersuite_ECDHE_RSA_WITH_AES_128_CBC_SHA256"
	case Tls_ciphersuite_ECDHE_RSA_WITH_AES_128_GCM_SHA256:
		return "Tls_ciphersuite_ECDHE_RSA_WITH_AES_128_GCM_SHA256"
	case Tls_ciphersuite_ECDHE_RSA_WITH_AES_256_CBC_SHA:
		return "Tls_ciphersuite_ECDHE_RSA_WITH_AES_256_CBC_SHA"
	case Tls_ciphersuite_ECDHE_RSA_WITH_AES_256_CBC_SHA384:
		return "Tls_ciphersuite_ECDHE_RSA_WITH_AES_256_CBC_SHA384"
	case Tls_ciphersuite_ECDHE_RSA_WITH_AES_256_GCM_SHA384:
		return "Tls_ciphersuite_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
	case Tls_ciphersuite_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256:
		return "Tls_ciphersuite_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256"
	case Tls_ciphersuite_RSA_WITH_AES_128_CBC_SHA:
		return "Tls_ciphersuite_RSA_WITH_AES_128_CBC_SHA"
	case Tls_ciphersuite_RSA_WITH_AES_128_CBC_SHA256:
		return "Tls_ciphersuite_RSA_WITH_AES_128_CBC_SHA256"
	case Tls_ciphersuite_RSA_WITH_AES_128_GCM_SHA256:
		return "Tls_ciphersuite_RSA_WITH_AES_128_GCM_SHA256"
	case Tls_ciphersuite_RSA_WITH_AES_256_CBC_SHA:
		return "Tls_ciphersuite_RSA_WITH_AES_256_CBC_SHA"
	case Tls_ciphersuite_RSA_WITH_AES_256_CBC_SHA256:
		return "Tls_ciphersuite_RSA_WITH_AES_256_CBC_SHA256"
	case Tls_ciphersuite_RSA_WITH_AES_256_GCM_SHA384:
		return "Tls_ciphersuite_RSA_WITH_AES_256_GCM_SHA384"
	case Tls_ciphersuite_ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA:
		return "Tls_ciphersuite_ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA"
	case Tls_ciphersuite_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA:
		return "Tls_ciphersuite_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA"
	case Tls_ciphersuite_RSA_WITH_3DES_EDE_CBC_SHA:
		return "Tls_ciphersuite_RSA_WITH_3DES_EDE_CBC_SHA"
	default:
		return fmt.Sprintf("Tls_ciphersuite_t(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Security/tls_protocol_version_t
type Tls_protocol_version_t uint16

const (
	// Tls_protocol_version_DTLSv12: The DTLS 1.2 protocol.
	Tls_protocol_version_DTLSv12 Tls_protocol_version_t = 0xfefd
	// Tls_protocol_version_TLSv12: The TLS 1.2 protocol.
	Tls_protocol_version_TLSv12 Tls_protocol_version_t = 0x303
	// Tls_protocol_version_TLSv13: The TLS 1.3 protocol.
	Tls_protocol_version_TLSv13 Tls_protocol_version_t = 0x304
	// Deprecated.
	Tls_protocol_version_DTLSv10 Tls_protocol_version_t = 0xfeff
	// Deprecated.
	Tls_protocol_version_TLSv10 Tls_protocol_version_t = 0x301
	// Deprecated.
	Tls_protocol_version_TLSv11 Tls_protocol_version_t = 0x302
)

func (e Tls_protocol_version_t) String() string {
	switch e {
	case Tls_protocol_version_DTLSv12:
		return "Tls_protocol_version_DTLSv12"
	case Tls_protocol_version_TLSv12:
		return "Tls_protocol_version_TLSv12"
	case Tls_protocol_version_TLSv13:
		return "Tls_protocol_version_TLSv13"
	case Tls_protocol_version_DTLSv10:
		return "Tls_protocol_version_DTLSv10"
	case Tls_protocol_version_TLSv10:
		return "Tls_protocol_version_TLSv10"
	case Tls_protocol_version_TLSv11:
		return "Tls_protocol_version_TLSv11"
	default:
		return fmt.Sprintf("Tls_protocol_version_t(%d)", e)
	}
}

// CssmAppledlOpenParametersMask is a Go-name alias for Cssm_appledl_open_parameters_mask.
type CssmAppledlOpenParametersMask = Cssm_appledl_open_parameters_mask

// ExtensionDataFormat is a Go-name alias for Extension_data_format.
type ExtensionDataFormat = Extension_data_format

// TLSCiphersuiteGroup is a Go-name alias for Tls_ciphersuite_group_t.
type TLSCiphersuiteGroup = Tls_ciphersuite_group_t

// TLSCiphersuite is a Go-name alias for Tls_ciphersuite_t.
type TLSCiphersuite = Tls_ciphersuite_t

// TLSProtocolVersion is a Go-name alias for Tls_protocol_version_t.
type TLSProtocolVersion = Tls_protocol_version_t
