// Code generated from Apple documentation for Security. DO NOT EDIT.

package security
import (
	"unsafe"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// C struct types
// AuthorizationCallbacks - The interface implemented by the Security Server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCallbacks
type AuthorizationCallbacks struct {
	SetResult func(uintptr, AuthorizationResult) int // Returns the result of an authorization operation.
	RequestInterrupt func(uintptr) int // Requests the authorization engine to interrupt the currently active authorization mechanism.
	DidDeactivate func(uintptr) int // Reports the successful deactivation of an authorization mechanism.
	GetContextValue func(uintptr, *byte, *AuthorizationContextFlags, uintptr) int // Reads a value collected during authorization.
	SetContextValue func(uintptr, *byte, uint, uintptr) int // Stores data collected during authorization as a key-value pair.
	RemoveContextValue func(uintptr, *byte) int // Removes a value collected during authorization.
	GetHintValue func(uintptr, *byte, uintptr) int // Reads a value stored by the plug-in authorization mechanism.
	GetImmutableHintValue func(uintptr, *byte, uintptr) int // Reads an immutable value stored by the plug-in authorization mechanism.
	SetHintValue func(uintptr, *byte, uintptr) int // Stores data needed during authorization as a key-value pair.
	RemoveHintValue func(uintptr, *byte) int // Removes a value stored by the plug-in authorization mechanism.
	GetArguments func(uintptr, uintptr) int // Reads the arguments for this authorization mechanism from the authorization policy database.
	GetSessionId func(uintptr, unsafe.Pointer) int // Reads the session ID.
	GetLAContext func(uintptr, unsafe.Pointer) int // Constructs a local authentication context.
	GetTokenIdentities func(uintptr, unsafe.Pointer, uintptr) int // Returns an array of identities available on tokens.
	GetTKTokenWatcher func(uintptr, unsafe.Pointer) int // Constructs a token watcher.
	Version uint32 // The engine callback version.

}

// AuthorizationExternalForm - The external representation of an authorization reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationExternalForm
type AuthorizationExternalForm struct {
	Bytes int8 // An array of characters representing the external form of an authorization reference.

}

// AuthorizationItem - A structure containing information about an authorization right or the authorization environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationItem
type AuthorizationItem struct {
	Name AuthorizationString // The required name of the authorization right or environment data.
	ValueLength uintptr // The number of bytes in the value field.
	Value unsafe.Pointer // A pointer to information pertaining to the name field.
	Flags uint32 // Reserved option bits.

}

// AuthorizationItemSet - A structure containing a set of authorization items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationItemSet
type AuthorizationItemSet struct {
	Count uint32 // The number of elements in the `items` array.
	Items *AuthorizationItem // A pointer to an array of authorization items.

}

// AuthorizationPluginInterface - The interface that must be implemented by your plug-in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationPluginInterface
type AuthorizationPluginInterface struct {
	MechanismCreate func(unsafe.Pointer, uintptr, *byte, unsafe.Pointer) int // Creates an authorization mechanism.
	MechanismDeactivate func(unsafe.Pointer) int // Deactivates an authorization mechanism.
	MechanismDestroy func(unsafe.Pointer) int // Destroys an authorization mechanism.
	MechanismInvoke func(unsafe.Pointer) int // Invokes an authorization mechanism to perform an authorization operation.
	PluginDestroy func(unsafe.Pointer) int // Notifies the plug-in that it is about to be unloaded.
	Version uint32 // The plug-in interface version.

}

// AuthorizationValue - A structure used to pass data between the authorization engine and the plug-in mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationValue
type AuthorizationValue struct {
	Data unsafe.Pointer
	Length uintptr

}

// AuthorizationValueVector - A structure used to pass arguments from the authorization policy database to the authorization mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationValueVector
type AuthorizationValueVector struct {
	Count uint32
	Values *AuthorizationValue

}

// CSSM_APPLE_CL_CSR_REQUEST
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_CL_CSR_REQUEST
type CSSM_APPLE_CL_CSR_REQUEST struct {
	SubjectNameX509 unsafe.Pointer
	SignatureAlg CSSM_ALGORITHMS
	SignatureOid unsafe.Pointer
	CspHand CSSM_CSP_HANDLE
	ChallengeString *byte
	SubjectPrivateKey unsafe.Pointer
	SubjectPublicKey unsafe.Pointer

}

// CSSM_APPLE_TP_ACTION_DATA
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_ACTION_DATA
type CSSM_APPLE_TP_ACTION_DATA struct {
	Version uint32
	ActionFlags CSSM_APPLE_TP_ACTION_FLAGS

}

// CSSM_APPLE_TP_CERT_REQUEST
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_CERT_REQUEST
type CSSM_APPLE_TP_CERT_REQUEST struct {
	CspHand CSSM_CSP_HANDLE
	ClHand CSSM_CL_HANDLE
	SerialNumber uint32
	NumSubjectNames uint32
	NumIssuerNames uint32
	IssuerNameX509 unsafe.Pointer
	SignatureAlg CSSM_ALGORITHMS
	SignatureOid unsafe.Pointer
	NotBefore uint32
	NotAfter uint32
	NumExtensions uint32
	CertPublicKey unsafe.Pointer
	ChallengeString *byte
	Extensions unsafe.Pointer
	IssuerNames *CSSM_APPLE_TP_NAME_OID
	IssuerPrivateKey unsafe.Pointer
	SubjectNames *CSSM_APPLE_TP_NAME_OID

}

// CSSM_APPLE_TP_CRL_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_CRL_OPTIONS
type CSSM_APPLE_TP_CRL_OPTIONS struct {
	Version uint32
	CrlFlags CSSM_APPLE_TP_CRL_OPT_FLAGS
	CrlStore unsafe.Pointer

}

// CSSM_APPLE_TP_NAME_OID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_NAME_OID
type CSSM_APPLE_TP_NAME_OID struct {
	Oid unsafe.Pointer
	String *byte

}

// CSSM_APPLE_TP_SMIME_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_SMIME_OPTIONS
type CSSM_APPLE_TP_SMIME_OPTIONS struct {
	Version uint32
	IntendedUsage unsafe.Pointer
	SenderEmailLen uint32
	SenderEmail *byte

}

// CSSM_APPLE_TP_SSL_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_SSL_OPTIONS
type CSSM_APPLE_TP_SSL_OPTIONS struct {
	Version uint32
	ServerNameLen uint32
	Flags uint32
	ServerName *byte

}

// CSSM_TP_APPLE_EVIDENCE_HEADER
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_APPLE_EVIDENCE_HEADER
type CSSM_TP_APPLE_EVIDENCE_HEADER struct {
	Version uint32

}

// CSSM_TP_APPLE_EVIDENCE_INFO
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_APPLE_EVIDENCE_INFO
type CSSM_TP_APPLE_EVIDENCE_INFO struct {
	CrlReason int32
	DlDbHandle unsafe.Pointer
	Index uint32
	NumStatusCodes uint32
	StatusBits CSSM_TP_APPLE_CERT_STATUS
	StatusCodes unsafe.Pointer
	UniqueRecord unsafe.Pointer

}

// CSSM_TUPLE
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TUPLE
type CSSM_TUPLE struct {
	AuthorizationTag unsafe.Pointer
	Delegate CSSM_BOOL
	Issuer unsafe.Pointer
	Subject unsafe.Pointer
	ValidityPeriod unsafe.Pointer

}

// SecAsn1AlgId - A structure identifying an ASN.1 algorithm by its OID, and its corresponding parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1AlgId
type SecAsn1AlgId struct {
	Algorithm unsafe.Pointer
	Parameters unsafe.Pointer

}

// SecAsn1PubKeyInfo - A structure containing a public key and its associated algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1PubKeyInfo
type SecAsn1PubKeyInfo struct {
	SubjectPublicKey unsafe.Pointer
	Algorithm SecAsn1AlgId

}

// SecAsn1Template_struct - A structure that defines one element of a BER or DER encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1Template_struct
type SecAsn1Template_struct struct {
	Kind uint32
	Offset uint32
	Size uint32
	Sub unsafe.Pointer

}

// SecCEBasicConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCEBasicConstraints
type SecCEBasicConstraints struct {
	Critical bool
	IsCA bool
	PathLenConstraint uint32
	PathLenConstraintPresent bool
	Present bool

}

// SecCEPolicyConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCEPolicyConstraints
type SecCEPolicyConstraints struct {
	Critical bool
	InhibitPolicyMapping uint32
	InhibitPolicyMappingPresent bool
	Present bool
	RequireExplicitPolicy uint32
	RequireExplicitPolicyPresent bool

}

// SecItemImportExportKeyParameters - The import/export parameter structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemImportExportKeyParameters
type SecItemImportExportKeyParameters struct {
	AccessRef SecAccessRef // Specifies the initial access controls of imported private keys.
	AlertPrompt corefoundation.CFStringRef // The prompt to display in the secure passphrase alert panel.
	AlertTitle corefoundation.CFStringRef // The title to display in the secure passphrase alert panel.
	Flags SecKeyImportExportFlags // The bitwise [OR] of zero or more key import/export flags.
	KeyAttributes corefoundation.CFArrayRef // An array containing zero or more key attributes for an imported key.
	KeyUsage corefoundation.CFArrayRef // An array containing usage attributes applied to a key on import.
	Passphrase corefoundation.CFTypeRef // The password to use during key import or export.
	Version uint32 // The version of this structure.

}

// SecKeyImportExportParameters - The legacy import/export parameter structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyImportExportParameters
type SecKeyImportExportParameters struct {
	AccessRef SecAccessRef // Specifies the initial access controls of imported private keys.
	AlertPrompt corefoundation.CFStringRef // The prompt to display in the secure passphrase alert panel.
	AlertTitle corefoundation.CFStringRef // The title to display in the secure passphrase alert panel.
	Flags SecKeyImportExportFlags // The bitwise [OR] of zero or more key import/export flags.
	KeyAttributes CSSM_KEYATTR_FLAGS // A word of bits constituting the low-level attribute flags for imported keys.
	KeyUsage CSSM_KEYUSE // A word of bits constituting the low-level use flags for imported keys.
	Passphrase corefoundation.CFTypeRef // The password to use during key import or export.
	Version uint32 // The version of this structure.

}

// SecKeychainAttribute - A structure that holds a single keychain attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttribute
type SecKeychainAttribute struct {
	Data unsafe.Pointer // A pointer to the attribute data.
	Length uint32 // The length of the buffer pointed to by data.
	Tag SecKeychainAttrType // A 4-byte attribute tag.

}

// SecKeychainAttributeInfo - A structure that represents an attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttributeInfo
type SecKeychainAttributeInfo struct {
	Count uint32 // The number of tag-format pairs in the respective arrays.
	Format *uint32 // A pointer to the first attribute format in the array.
	Tag *uint32 // A pointer to the first attribute tag in the array.

}

// SecKeychainAttributeList - A list of keychain attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttributeList
type SecKeychainAttributeList struct {
	Attr *SecKeychainAttribute // A pointer to the first keychain attribute in the array.
	Count uint32 // The number of keychain attributes in the array.

}

// SecKeychainCallbackInfo - Information about a keychain event that keychain services deliver to your app via a callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainCallbackInfo
type SecKeychainCallbackInfo struct {
	Item SecKeychainItemRef // A reference to the keychain item in which the event occurred. If the event did not involve an item, this field is not valid.
	Keychain SecKeychainRef // A reference to the keychain in which the event occurred. If the event did not involve a keychain, this field is not valid.
	Pid int32 // The ID of the process that generated this event.
	Version uint32 // The version of this structure.

}

// SecKeychainSettings - A structure that contains information about keychain settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSettings
type SecKeychainSettings struct {
	Version uint32 // The keychain version.
	LockOnSleep bool // A Boolean value indicating whether the keychain locks when the system sleeps.
	UseLockInterval bool // A Boolean value indicating whether the keychain automatically locks after a certain period of time.
	LockInterval uint32 // The number of seconds to wait before the keychain locks.

}

// _CE_AccessDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_AccessDescription
type _CE_AccessDescription struct {
	AccessLocation string
	AccessMethod unsafe.Pointer

}

// _CE_AuthorityInfoAccess
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_AuthorityInfoAccess
type _CE_AuthorityInfoAccess struct {
	AccessDescriptions unsafe.Pointer
	NumAccessDescriptions uint32

}

// _CE_AuthorityKeyID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_AuthorityKeyID
type _CE_AuthorityKeyID struct {
	GeneralNames unsafe.Pointer
	GeneralNamesPresent CSSM_BOOL
	KeyIdentifier unsafe.Pointer
	KeyIdentifierPresent CSSM_BOOL
	SerialNumber unsafe.Pointer
	SerialNumberPresent CSSM_BOOL

}

// _CE_BasicConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_BasicConstraints
type _CE_BasicConstraints struct {
	CA CSSM_BOOL
	PathLenConstraint uint32
	PathLenConstraintPresent CSSM_BOOL

}

// _CE_CRLDistPointsSyntax
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CRLDistPointsSyntax
type _CE_CRLDistPointsSyntax struct {
	DistPoints unsafe.Pointer
	NumDistPoints uint32

}

// _CE_CRLDistributionPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CRLDistributionPoint
type _CE_CRLDistributionPoint struct {
	CrlIssuer unsafe.Pointer
	DistPointName unsafe.Pointer
	Reasons unsafe.Pointer
	ReasonsPresent CSSM_BOOL

}

// _CE_CertPolicies
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CertPolicies
type _CE_CertPolicies struct {
	NumPolicies uint32
	Policies unsafe.Pointer

}

// _CE_DataAndType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_DataAndType
type _CE_DataAndType struct {
	Critical CSSM_BOOL
	Extension unsafe.Pointer
	Type CE_DataType

}

// _CE_DistributionPointName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_DistributionPointName
type _CE_DistributionPointName struct {
	Dpn unsafe.Pointer
	NameType CE_CrlDistributionPointNameType

}

// _CE_ExtendedKeyUsage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_ExtendedKeyUsage-c.struct
type _CE_ExtendedKeyUsage struct {
	NumPurposes uint32
	Purposes unsafe.Pointer

}

// _CE_GeneralName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralName
type _CE_GeneralName struct {
	BerEncoded CSSM_BOOL
	Name unsafe.Pointer
	NameType int

}

// _CE_GeneralNames
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralNames
type _CE_GeneralNames struct {
	GeneralName unsafe.Pointer
	NumNames uint32

}

// _CE_GeneralSubtree
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralSubtree
type _CE_GeneralSubtree struct {
	Base unsafe.Pointer
	Maximum uint32
	MaximumPresent CSSM_BOOL
	Minimum uint32

}

// _CE_GeneralSubtrees
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralSubtrees
type _CE_GeneralSubtrees struct {
	NumSubtrees uint32
	Subtrees unsafe.Pointer

}

// _CE_IssuingDistributionPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_IssuingDistributionPoint
type _CE_IssuingDistributionPoint struct {
	DistPointName unsafe.Pointer
	IndirectCrl CSSM_BOOL
	IndirectCrlPresent CSSM_BOOL
	OnlyCACerts CSSM_BOOL
	OnlyCACertsPresent CSSM_BOOL
	OnlySomeReasons unsafe.Pointer
	OnlySomeReasonsPresent CSSM_BOOL
	OnlyUserCerts CSSM_BOOL
	OnlyUserCertsPresent CSSM_BOOL

}

// _CE_NameConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_NameConstraints
type _CE_NameConstraints struct {
	Excluded unsafe.Pointer
	Permitted unsafe.Pointer

}

// _CE_OtherName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_OtherName
type _CE_OtherName struct {
	TypeId unsafe.Pointer
	Value unsafe.Pointer

}

// _CE_PolicyConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyConstraints
type _CE_PolicyConstraints struct {
	InhibitPolicyMapping uint32
	InhibitPolicyMappingPresent CSSM_BOOL
	RequireExplicitPolicy uint32
	RequireExplicitPolicyPresent CSSM_BOOL

}

// _CE_PolicyInformation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyInformation
type _CE_PolicyInformation struct {
	CertPolicyId unsafe.Pointer
	NumPolicyQualifiers uint32
	PolicyQualifiers unsafe.Pointer

}

// _CE_PolicyMapping
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyMapping
type _CE_PolicyMapping struct {
	IssuerDomainPolicy unsafe.Pointer
	SubjectDomainPolicy unsafe.Pointer

}

// _CE_PolicyMappings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyMappings
type _CE_PolicyMappings struct {
	NumPolicyMappings uint32
	PolicyMappings unsafe.Pointer

}

// _CE_PolicyQualifierInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyQualifierInfo
type _CE_PolicyQualifierInfo struct {
	PolicyQualifierId unsafe.Pointer
	Qualifier unsafe.Pointer

}

// _CE_QC_Statement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_QC_Statement
type _CE_QC_Statement struct {
	OtherInfo unsafe.Pointer
	SemanticsInfo unsafe.Pointer
	StatementId unsafe.Pointer

}

// _CE_QC_Statements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_QC_Statements
type _CE_QC_Statements struct {
	NumQCStatements uint32
	QcStatements unsafe.Pointer

}

// _CE_SemanticsInformation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_SemanticsInformation
type _CE_SemanticsInformation struct {
	NameRegistrationAuthorities unsafe.Pointer
	SemanticsIdentifier unsafe.Pointer

}

// Cssm_access_credentials
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_access_credentials-c.struct
type Cssm_access_credentials struct {
	BaseCerts unsafe.Pointer
	Callback unsafe.Pointer
	CallerCtx unsafe.Pointer
	EntryTag CSSM_STRING
	Samples unsafe.Pointer

}

// Cssm_acl_edit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_edit-c.struct
type Cssm_acl_edit struct {
	EditMode CSSM_ACL_EDIT_MODE
	NewEntry unsafe.Pointer
	OldEntryHandle CSSM_ACL_HANDLE

}

// Cssm_acl_entry_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_entry_info-c.struct
type Cssm_acl_entry_info struct {
	EntryHandle CSSM_ACL_HANDLE
	EntryPublicInfo unsafe.Pointer

}

// Cssm_acl_entry_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_entry_input-c.struct
type Cssm_acl_entry_input struct {
	Callback unsafe.Pointer
	CallerContext unsafe.Pointer
	Prototype unsafe.Pointer

}

// Cssm_acl_entry_prototype
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_entry_prototype-c.struct
type Cssm_acl_entry_prototype struct {
	Authorization unsafe.Pointer
	Delegate CSSM_BOOL
	EntryTag CSSM_STRING
	TimeRange unsafe.Pointer
	TypedSubject unsafe.Pointer

}

// Cssm_acl_keychain_prompt_selector
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_keychain_prompt_selector-swift.struct
type Cssm_acl_keychain_prompt_selector struct {
	Version uint16
	Flags uint16

}

// Cssm_acl_owner_prototype
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_owner_prototype-c.struct
type Cssm_acl_owner_prototype struct {
	Delegate CSSM_BOOL
	TypedSubject unsafe.Pointer

}

// Cssm_acl_process_subject_selector
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_process_subject_selector-swift.struct
type Cssm_acl_process_subject_selector struct {
	Version uint16
	Mask uint16
	Uid uint32
	Gid uint32

}

// Cssm_acl_validity_period
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_validity_period-c.struct
type Cssm_acl_validity_period struct {
	EndDate unsafe.Pointer
	StartDate unsafe.Pointer

}

// Cssm_applecspdl_db_change_password_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_applecspdl_db_change_password_parameters-swift.struct
type Cssm_applecspdl_db_change_password_parameters struct {
	AccessCredentials unsafe.Pointer

}

// Cssm_applecspdl_db_is_locked_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_applecspdl_db_is_locked_parameters-swift.struct
type Cssm_applecspdl_db_is_locked_parameters struct {
	IsLocked uint8

}

// Cssm_applecspdl_db_settings_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_applecspdl_db_settings_parameters-swift.struct
type Cssm_applecspdl_db_settings_parameters struct {
	IdleTimeout uint32
	LockOnSleep uint8

}

// Cssm_appledl_open_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_appledl_open_parameters-swift.struct
type Cssm_appledl_open_parameters struct {
	Length uint32
	Version uint32
	AutoCommit CSSM_BOOL
	Mask uint32
	Mode uint16

}

// Cssm_authorizationgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_authorizationgroup-swift.struct
type Cssm_authorizationgroup struct {
	NumberOfAuthTags uint32
	AuthTags unsafe.Pointer

}

// Cssm_base_certs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_base_certs-c.struct
type Cssm_base_certs struct {
	CLHandle CSSM_CL_HANDLE
	Certs unsafe.Pointer
	TPHandle CSSM_TP_HANDLE

}

// Cssm_cert_bundle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_cert_bundle-c.struct
type Cssm_cert_bundle struct {
	Bundle unsafe.Pointer
	BundleHeader unsafe.Pointer

}

// Cssm_cert_bundle_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_cert_bundle_header-c.struct
type Cssm_cert_bundle_header struct {
	BundleEncoding CSSM_CERT_BUNDLE_ENCODING
	BundleType CSSM_CERT_BUNDLE_TYPE

}

// Cssm_cert_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_cert_pair-c.struct
type Cssm_cert_pair struct {
	EncodedCert unsafe.Pointer
	ParsedCert unsafe.Pointer

}

// Cssm_certgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_certgroup-c.struct
type Cssm_certgroup struct {
	CertEncoding CSSM_CERT_ENCODING
	CertGroupType CSSM_CERTGROUP_TYPE
	CertType CSSM_CERT_TYPE
	GroupList unsafe.Pointer
	NumCerts uint32
	Reserved unsafe.Pointer

}

// Cssm_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_context-c.struct
type Cssm_context struct {
	AlgorithmType CSSM_ALGORITHMS
	CSPHandle CSSM_CSP_HANDLE
	ContextAttributes unsafe.Pointer
	ContextType CSSM_CONTEXT_TYPE
	EncryptionProhibited uint32
	NumberOfAttributes uint32
	Privileged CSSM_BOOL
	Reserved uint32
	WorkFactor uint32

}

// Cssm_context_attribute
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_context_attribute-c.struct
type Cssm_context_attribute struct {
	Attribute unsafe.Pointer
	AttributeLength uint32
	AttributeType CSSM_ATTRIBUTE_TYPE

}

// Cssm_crl_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_crl_pair-c.struct
type Cssm_crl_pair struct {
	EncodedCrl unsafe.Pointer
	ParsedCrl unsafe.Pointer

}

// Cssm_crlgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_crlgroup-c.struct
type Cssm_crlgroup struct {
	CrlEncoding CSSM_CRL_ENCODING
	CrlGroupType CSSM_CRLGROUP_TYPE
	CrlType CSSM_CRL_TYPE
	GroupCrlList unsafe.Pointer
	NumberOfCrls uint32

}

// Cssm_crypto_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_crypto_data-c.struct
type Cssm_crypto_data struct {
	Callback unsafe.Pointer
	CallerCtx unsafe.Pointer
	Param unsafe.Pointer

}

// Cssm_csp_operational_statistics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_csp_operational_statistics-swift.struct
type Cssm_csp_operational_statistics struct {
	UserAuthenticated CSSM_BOOL
	DeviceFlags CSSM_CSP_FLAGS
	TokenMaxSessionCount uint32
	TokenOpenedSessionCount uint32
	TokenMaxRWSessionCount uint32
	TokenOpenedRWSessionCount uint32
	TokenTotalPublicMem uint32
	TokenFreePublicMem uint32
	TokenTotalPrivateMem uint32
	TokenFreePrivateMem uint32

}

// Cssm_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_data-swift.struct
type Cssm_data struct {
	Length uintptr
	Data *uint8

}

// Cssm_date
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_date-swift.struct
type Cssm_date struct {
	Day uint8
	Month uint8
	Year uint8

}

// Cssm_db_attribute_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_attribute_data-c.struct
type Cssm_db_attribute_data struct {
	Info unsafe.Pointer
	NumberOfValues uint32
	Value unsafe.Pointer

}

// Cssm_db_attribute_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_attribute_info-c.struct
type Cssm_db_attribute_info struct {
	AttributeFormat CSSM_DB_ATTRIBUTE_FORMAT
	AttributeNameFormat CSSM_DB_ATTRIBUTE_NAME_FORMAT
	Label unsafe.Pointer

}

// Cssm_db_index_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_index_info-c.struct
type Cssm_db_index_info struct {
	IndexType CSSM_DB_INDEX_TYPE
	IndexedDataLocation CSSM_DB_INDEXED_DATA_LOCATION
	Info unsafe.Pointer

}

// Cssm_db_parsing_module_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_parsing_module_info-c.struct
type Cssm_db_parsing_module_info struct {
	ModuleSubserviceUid unsafe.Pointer
	RecordType CSSM_DB_RECORDTYPE

}

// Cssm_db_record_attribute_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_record_attribute_data-c.struct
type Cssm_db_record_attribute_data struct {
	AttributeData unsafe.Pointer
	DataRecordType CSSM_DB_RECORDTYPE
	NumberOfAttributes uint32
	SemanticInformation uint32

}

// Cssm_db_record_attribute_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_record_attribute_info-c.struct
type Cssm_db_record_attribute_info struct {
	AttributeInfo unsafe.Pointer
	DataRecordType CSSM_DB_RECORDTYPE
	NumberOfAttributes uint32

}

// Cssm_db_record_index_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_record_index_info-c.struct
type Cssm_db_record_index_info struct {
	DataRecordType CSSM_DB_RECORDTYPE
	IndexInfo unsafe.Pointer
	NumberOfIndexes uint32

}

// Cssm_db_schema_attribute_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_schema_attribute_info-c.struct
type Cssm_db_schema_attribute_info struct {
	AttributeId uint32
	AttributeName *byte
	AttributeNameID unsafe.Pointer
	DataType CSSM_DB_ATTRIBUTE_FORMAT

}

// Cssm_db_schema_index_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_schema_index_info-swift.struct
type Cssm_db_schema_index_info struct {
	AttributeId uint32
	IndexId uint32
	IndexType CSSM_DB_INDEX_TYPE
	IndexedDataLocation CSSM_DB_INDEXED_DATA_LOCATION

}

// Cssm_db_unique_record
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_unique_record-c.struct
type Cssm_db_unique_record struct {
	RecordIdentifier unsafe.Pointer
	RecordLocator unsafe.Pointer

}

// Cssm_dbinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_dbinfo-c.struct
type Cssm_dbinfo struct {
	AccessPath *byte
	DefaultParsingModules unsafe.Pointer
	IsLocal CSSM_BOOL
	NumberOfRecordTypes uint32
	RecordAttributeNames unsafe.Pointer
	RecordIndexes unsafe.Pointer
	Reserved unsafe.Pointer

}

// Cssm_dl_db_handle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_dl_db_handle-swift.struct
type Cssm_dl_db_handle struct {
	DLHandle CSSM_DL_HANDLE
	DBHandle CSSM_DB_HANDLE

}

// Cssm_dl_db_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_dl_db_list-c.struct
type Cssm_dl_db_list struct {
	DLDBHandle unsafe.Pointer
	NumHandles uint32

}

// Cssm_dl_pkcs11_attributes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_dl_pkcs11_attributes
type Cssm_dl_pkcs11_attributes struct {
	DeviceAccessFlags uint32

}

// Cssm_encoded_cert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_encoded_cert-c.struct
type Cssm_encoded_cert struct {
	CertBlob unsafe.Pointer
	CertEncoding CSSM_CERT_ENCODING
	CertType CSSM_CERT_TYPE

}

// Cssm_encoded_crl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_encoded_crl-c.struct
type Cssm_encoded_crl struct {
	CrlBlob unsafe.Pointer
	CrlEncoding CSSM_CRL_ENCODING
	CrlType CSSM_CRL_TYPE

}

// Cssm_evidence
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_evidence-c.struct
type Cssm_evidence struct {
	Evidence unsafe.Pointer
	EvidenceForm CSSM_EVIDENCE_FORM

}

// Cssm_field
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_field-c.struct
type Cssm_field struct {
	FieldOid unsafe.Pointer
	FieldValue unsafe.Pointer

}

// Cssm_fieldgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_fieldgroup-c.struct
type Cssm_fieldgroup struct {
	Fields unsafe.Pointer
	NumberOfFields int

}

// Cssm_func_name_addr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_func_name_addr-swift.struct
type Cssm_func_name_addr struct {
	Name CSSM_STRING
	Address CSSM_PROC_ADDR

}

// Cssm_guid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_guid-swift.struct
type Cssm_guid struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 uint8

}

// Cssm_kea_derive_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kea_derive_params-c.struct
type Cssm_kea_derive_params struct {
	Rb unsafe.Pointer
	Yb unsafe.Pointer

}

// Cssm_key
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_key-c.struct
type Cssm_key struct {
	KeyData unsafe.Pointer
	KeyHeader unsafe.Pointer

}

// Cssm_key_size
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_key_size-swift.struct
type Cssm_key_size struct {
	LogicalKeySizeInBits uint32
	EffectiveKeySizeInBits uint32

}

// Cssm_keyheader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_keyheader-c.struct
type Cssm_keyheader struct {
	AlgorithmId CSSM_ALGORITHMS
	BlobType CSSM_KEYBLOB_TYPE
	CspId unsafe.Pointer
	EndDate unsafe.Pointer
	Format CSSM_KEYBLOB_FORMAT
	HeaderVersion CSSM_HEADERVERSION
	KeyAttr CSSM_KEYATTR_FLAGS
	KeyClass CSSM_KEYCLASS
	KeyUsage CSSM_KEYUSE
	LogicalKeySizeInBits uint32
	Reserved uint32
	StartDate unsafe.Pointer
	WrapAlgorithmId CSSM_ALGORITHMS
	WrapMode CSSM_ENCRYPT_MODE

}

// Cssm_kr_name
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_name-swift.struct
type Cssm_kr_name struct {
	Type uint8
	Length uint8
	Name *byte

}

// Cssm_kr_policy_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_policy_info-c.struct
type Cssm_kr_policy_info struct {
	KrbNotAllowed CSSM_BOOL
	NumberOfEntries uint32
	PolicyEntry unsafe.Pointer

}

// Cssm_kr_policy_list_item
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_policy_list_item-c.struct
type Cssm_kr_policy_list_item struct {
	AlgClass CSSM_CONTEXT_TYPE
	AlgorithmId CSSM_ALGORITHMS
	MaxKeyLength uint32
	MaxRounds uint32
	Mode CSSM_ENCRYPT_MODE
	PolicyFlags CSSM_KR_POLICY_FLAGS
	WorkFactor uint8
	Next objectivec.IObject

}

// Cssm_kr_profile
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_profile-c.struct
type Cssm_kr_profile struct {
	ENT_KRACertChainList unsafe.Pointer
	ENT_KRANum uint8
	INDIV_AuthenticationInfo unsafe.Pointer
	INDIV_KRACertChainList unsafe.Pointer
	INDIV_KRANum uint8
	KRSCertChain unsafe.Pointer
	KRSPExtensions unsafe.Pointer
	KRSPFlags uint32
	LE_KRACertChainList unsafe.Pointer
	LE_KRANum uint8
	UserCertificate unsafe.Pointer
	UserName unsafe.Pointer

}

// Cssm_kr_wrappedproductinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_wrappedproductinfo
type Cssm_kr_wrappedproductinfo struct {
	ProductDescription CSSM_STRING
	ProductFlags uint32
	ProductVendor CSSM_STRING
	ProductVersion unsafe.Pointer
	StandardDescription CSSM_STRING
	StandardVersion unsafe.Pointer

}

// Cssm_krsubservice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_krsubservice-c.struct
type Cssm_krsubservice struct {
	Description *byte
	SubServiceId uint32
	WrappedProduct unsafe.Pointer

}

// Cssm_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_list-swift.struct
type Cssm_list struct {
	ListType CSSM_LIST_TYPE
	Head CSSM_LIST_ELEMENT_PTR
	Tail CSSM_LIST_ELEMENT_PTR

}

// Cssm_list_element
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_list_element-c.struct
type Cssm_list_element struct {
	Element unsafe.Pointer
	ElementType CSSM_LIST_ELEMENT_TYPE
	NextElement *Cssm_list_element
	WordID CSSM_WORDID_TYPE

}

// Cssm_manager_event_notification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_manager_event_notification-c.struct
type Cssm_manager_event_notification struct {
	DestinationModuleManagerType CSSM_SERVICE_MASK
	Event CSSM_MANAGER_EVENT_TYPES
	EventData unsafe.Pointer
	EventId uint32
	SourceModuleManagerType CSSM_SERVICE_MASK

}

// Cssm_manager_registration_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_manager_registration_info-c.struct
type Cssm_manager_registration_info struct {
	DeregisterDispatchTable func() int
	EventNotifyManager func(uintptr) int
	Initialize func(uint, uint) int
	RefreshFunctionTable func(uintptr, uint) int
	RegisterDispatchTable func(uintptr) int
	Terminate func() int

}

// Cssm_memory_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_memory_funcs-swift.struct
type Cssm_memory_funcs struct {
	Malloc_func CSSM_MALLOC
	Free_func CSSM_FREE
	Realloc_func CSSM_REALLOC
	Calloc_func CSSM_CALLOC
	AllocRef unsafe.Pointer

}

// Cssm_module_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_module_funcs-c.struct
type Cssm_module_funcs struct {
	NumberOfServiceFuncs uint32
	ServiceFuncs CSSM_PROC_ADDR
	ServiceType CSSM_SERVICE_TYPE

}

// Cssm_name_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_name_list-swift.struct
type Cssm_name_list struct {
	NumStrings uint32
	String *byte

}

// Cssm_net_address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_net_address-c.struct
type Cssm_net_address struct {
	Address unsafe.Pointer
	AddressType CSSM_NET_ADDRESS_TYPE

}

// Cssm_parsed_cert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_parsed_cert-swift.struct
type Cssm_parsed_cert struct {
	CertType CSSM_CERT_TYPE
	ParsedCertFormat CSSM_CERT_PARSE_FORMAT
	ParsedCert unsafe.Pointer

}

// Cssm_parsed_crl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_parsed_crl-swift.struct
type Cssm_parsed_crl struct {
	CrlType CSSM_CRL_TYPE
	ParsedCrlFormat CSSM_CRL_PARSE_FORMAT
	ParsedCrl unsafe.Pointer

}

// Cssm_pkcs1_oaep_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_pkcs1_oaep_params-c.struct
type Cssm_pkcs1_oaep_params struct {
	HashAlgorithm uint32
	HashParams unsafe.Pointer
	MGF CSSM_PKCS_OAEP_MGF
	MGFParams unsafe.Pointer
	PSource CSSM_PKCS_OAEP_PSOURCE
	PSourceParams unsafe.Pointer

}

// Cssm_pkcs5_pbkdf1_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_pkcs5_pbkdf1_params-c.struct
type Cssm_pkcs5_pbkdf1_params struct {
	InitVector unsafe.Pointer
	Passphrase unsafe.Pointer

}

// Cssm_pkcs5_pbkdf2_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_pkcs5_pbkdf2_params-c.struct
type Cssm_pkcs5_pbkdf2_params struct {
	Passphrase unsafe.Pointer
	PseudoRandomFunction CSSM_PKCS5_PBKDF2_PRF

}

// Cssm_query
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_query-c.struct
type Cssm_query struct {
	Conjunctive CSSM_DB_CONJUNCTIVE
	NumSelectionPredicates uint32
	QueryFlags CSSM_QUERY_FLAGS
	QueryLimits unsafe.Pointer
	RecordType CSSM_DB_RECORDTYPE
	SelectionPredicate unsafe.Pointer

}

// Cssm_query_limits
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_query_limits-c.struct
type Cssm_query_limits struct {
	SizeLimit uint32
	TimeLimit uint32

}

// Cssm_query_size_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_query_size_data-swift.struct
type Cssm_query_size_data struct {
	SizeInputBlock uint32
	SizeOutputBlock uint32

}

// Cssm_range
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_range-swift.struct
type Cssm_range struct {
	Min uint32
	Max uint32

}

// Cssm_resource_control_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_resource_control_context-c.struct
type Cssm_resource_control_context struct {
	AccessCred unsafe.Pointer
	InitialAclEntry unsafe.Pointer

}

// Cssm_sample
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_sample-c.struct
type Cssm_sample struct {
	TypedSample unsafe.Pointer
	Verifier unsafe.Pointer

}

// Cssm_samplegroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_samplegroup-c.struct
type Cssm_samplegroup struct {
	NumberOfSamples uint32
	Samples unsafe.Pointer

}

// Cssm_selection_predicate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_selection_predicate-c.struct
type Cssm_selection_predicate struct {
	Attribute unsafe.Pointer
	DbOperator CSSM_DB_OPERATOR

}

// Cssm_spi_ac_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_ac_funcs-c.struct
type Cssm_spi_ac_funcs struct {
	AuthCompute func(int, uintptr, uintptr, uint, uintptr, uintptr, uintptr, uintptr) int
	PassThrough func(int, int, int, uint64, uintptr, uint, unsafe.Pointer, unsafe.Pointer) int

}

// Cssm_spi_cl_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_cl_funcs-c.struct
type Cssm_spi_cl_funcs struct {
	CertAbortCache func(int, int) int
	CertAbortQuery func(int, int) int
	CertCache func(int, uintptr, *int) int
	CertCreateTemplate func(int, uint, uintptr, uintptr) int
	CertDescribeFormat func(int, *uint, uintptr) int
	CertGetAllFields func(int, uintptr, *uint, uintptr) int
	CertGetAllTemplateFields func(int, uintptr, *uint, uintptr) int
	CertGetFirstCachedFieldValue func(int, int, uintptr, *int, *uint, uintptr) int
	CertGetFirstFieldValue func(int, uintptr, uintptr, *int, *uint, uintptr) int
	CertGetKeyInfo func(int, uintptr, uintptr) int
	CertGetNextCachedFieldValue func(int, int, uintptr) int
	CertGetNextFieldValue func(int, int, uintptr) int
	CertGroupFromVerifiedBundle func(int, uint64, uintptr, uintptr, uintptr) int
	CertGroupToSignedBundle func(int, uint64, uintptr, uintptr, uintptr) int
	CertSign func(int, uint64, uintptr, uintptr, uint, uintptr) int
	CertVerify func(int, uint64, uintptr, uintptr, uintptr, uint) int
	CertVerifyWithKey func(int, uint64, uintptr) int
	CrlAbortCache func(int, int) int
	CrlAbortQuery func(int, int) int
	CrlAddCert func(int, uint64, uintptr, uint, uintptr, uintptr, uintptr) int
	CrlCache func(int, uintptr, *int) int
	CrlCreateTemplate func(int, uint, uintptr, uintptr) int
	CrlDescribeFormat func(int, *uint, uintptr) int
	CrlGetAllCachedRecordFields func(int, int, uintptr, *uint, uintptr) int
	CrlGetAllFields func(int, uintptr, *uint, uintptr) int
	CrlGetFirstCachedFieldValue func(int, int, uintptr, uintptr, *int, *uint, uintptr) int
	CrlGetFirstFieldValue func(int, uintptr, uintptr, *int, *uint, uintptr) int
	CrlGetNextCachedFieldValue func(int, int, uintptr) int
	CrlGetNextFieldValue func(int, int, uintptr) int
	CrlRemoveCert func(int, uintptr, uintptr, uintptr) int
	CrlSetFields func(int, uint, uintptr, uintptr, uintptr) int
	CrlSign func(int, uint64, uintptr, uintptr, uint, uintptr) int
	CrlVerify func(int, uint64, uintptr, uintptr, uintptr, uint) int
	CrlVerifyWithKey func(int, uint64, uintptr) int
	FreeFieldValue func(int, uintptr, uintptr) int
	FreeFields func(int, uint, uintptr) int
	IsCertInCachedCrl func(int, uintptr, int, []int, uintptr) int
	IsCertInCrl func(int, uintptr, uintptr, []int) int
	PassThrough func(int, uint64, uint, unsafe.Pointer, unsafe.Pointer) int

}

// Cssm_spi_csp_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_csp_funcs-c.struct
type Cssm_spi_csp_funcs struct {
	ChangeKeyAcl func(int, uintptr, uintptr, uintptr) int
	ChangeKeyOwner func(int, uintptr, uintptr, uintptr) int
	ChangeLoginAcl func(int, uintptr, uintptr) int
	ChangeLoginOwner func(int, uintptr, uintptr) int
	DecryptData func(int, uint64, uintptr, uintptr, uint, uintptr, uint, *uint, uintptr, uint64) int
	DecryptDataFinal func(int, uint64, uintptr) int
	DecryptDataInit func(int, uint64, uintptr, uint64) int
	DecryptDataUpdate func(int, uint64, uintptr, uint, uintptr, uint, *uint) int
	DeriveKey func(int, uint64, uintptr, uintptr, uint, uint, uintptr, uintptr, uintptr) int
	DigestData func(int, uint64, uintptr, uintptr, uint, uintptr) int
	DigestDataClone func(int, uint64, uint64) int
	DigestDataFinal func(int, uint64, uintptr) int
	DigestDataInit func(int, uint64, uintptr) int
	DigestDataUpdate func(int, uint64, uintptr, uint) int
	EncryptData func(int, uint64, uintptr, uintptr, uint, uintptr, uint, *uint, uintptr, uint64) int
	EncryptDataFinal func(int, uint64, uintptr) int
	EncryptDataInit func(int, uint64, uintptr, uint64) int
	EncryptDataUpdate func(int, uint64, uintptr, uint, uintptr, uint, *uint) int
	EventNotify func(int, uint, uint64, uintptr) int
	FreeKey func(int, uintptr, uintptr, int) int
	GenerateAlgorithmParams func(int, uint64, uintptr, uint, uintptr, *uint, uintptr) int
	GenerateKey func(int, uint64, uintptr, uint, uint, uintptr, uintptr, uintptr, uint64) int
	GenerateKeyPair func(int, uint64, uintptr, uint, uint, uintptr, uintptr, uint, uint, uintptr, uintptr, uintptr, uint64) int
	GenerateMac func(int, uint64, uintptr, uintptr, uint, uintptr) int
	GenerateMacFinal func(int, uint64, uintptr) int
	GenerateMacInit func(int, uint64, uintptr) int
	GenerateMacUpdate func(int, uint64, uintptr, uint) int
	GenerateRandom func(int, uint64, uintptr, uintptr) int
	GetKeyAcl func(int, uintptr, [68]objectivec.IObject, *uint, uintptr) int
	GetKeyOwner func(int, uintptr, uintptr) int
	GetLoginAcl func(int, [68]objectivec.IObject, *uint, uintptr) int
	GetLoginOwner func(int, uintptr) int
	GetOperationalStatistics func(int, uintptr) int
	GetTimeValue func(int, uint, uintptr) int
	Login func(int, uintptr, uintptr, unsafe.Pointer) int
	Logout func(int) int
	ObtainPrivateKeyFromPublicKey func(int, uintptr, uintptr) int
	PassThrough func(int, uint64, uintptr, uint, unsafe.Pointer, unsafe.Pointer) int
	QueryKeySizeInBits func(int, uint64, uintptr, uintptr, uintptr) int
	QuerySize func(int, uint64, uintptr, int, uint, uintptr) int
	RetrieveCounter func(int, uintptr) int
	RetrieveUniqueId func(int, uintptr) int
	SignData func(int, uint64, uintptr, uintptr, uint, uint, uintptr) int
	SignDataFinal func(int, uint64, uintptr) int
	SignDataInit func(int, uint64, uintptr) int
	SignDataUpdate func(int, uint64, uintptr, uint) int
	UnwrapKey func(int, uint64, uintptr, uintptr, uintptr, uint, uint, uintptr, uintptr, uintptr, uintptr, uint64) int
	VerifyData func(int, uint64, uintptr, uintptr, uint, uint, uintptr) int
	VerifyDataFinal func(int, uint64, uintptr) int
	VerifyDataInit func(int, uint64, uintptr) int
	VerifyDataUpdate func(int, uint64, uintptr, uint) int
	VerifyDevice func(int, uintptr) int
	VerifyMac func(int, uint64, uintptr, uintptr, uint, uintptr) int
	VerifyMacFinal func(int, uint64, uintptr) int
	VerifyMacInit func(int, uint64, uintptr) int
	VerifyMacUpdate func(int, uint64, uintptr, uint) int
	WrapKey func(int, uint64, uintptr, uintptr, uintptr, uintptr, uintptr, uint64) int

}

// Cssm_spi_dl_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_dl_funcs-c.struct
type Cssm_spi_dl_funcs struct {
	Authenticate func(Cssm_dl_db_handle, uint, uintptr) int
	ChangeDbAcl func(Cssm_dl_db_handle, uintptr, uintptr) int
	ChangeDbOwner func(Cssm_dl_db_handle, uintptr, uintptr) int
	CreateRelation func(Cssm_dl_db_handle, uint, *byte, uint, uintptr, uint, uintptr) int
	DataAbortQuery func(Cssm_dl_db_handle, int) int
	DataDelete func(Cssm_dl_db_handle, uintptr) int
	DataGetFirst func(Cssm_dl_db_handle, uintptr, *int, uintptr, uintptr, uintptr) int
	DataGetFromUniqueRecordId func(Cssm_dl_db_handle, uintptr, uintptr, uintptr) int
	DataGetNext func(Cssm_dl_db_handle, int, uintptr, uintptr, uintptr) int
	DataInsert func(Cssm_dl_db_handle, uint, uintptr, uintptr, uintptr) int
	DataModify func(Cssm_dl_db_handle, uint, uintptr, uintptr, uintptr, uint) int
	DbClose func(Cssm_dl_db_handle) int
	DbCreate func(int, *byte, uintptr, uintptr, uint, uintptr, unsafe.Pointer, *int) int
	DbDelete func(int, *byte, uintptr, uintptr) int
	DbOpen func(int, *byte, uintptr, uint, uintptr, unsafe.Pointer, *int) int
	DestroyRelation func(Cssm_dl_db_handle, uint) int
	FreeNameList func(int, uintptr) int
	FreeUniqueRecord func(Cssm_dl_db_handle, uintptr) int
	GetDbAcl func(Cssm_dl_db_handle, [68]objectivec.IObject, *uint, uintptr) int
	GetDbNameFromHandle func(Cssm_dl_db_handle, *byte) int
	GetDbNames func(int, uintptr) int
	GetDbOwner func(Cssm_dl_db_handle, uintptr) int
	PassThrough func(Cssm_dl_db_handle, uint, unsafe.Pointer, unsafe.Pointer) int

}

// Cssm_spi_kr_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_kr_funcs-c.struct
type Cssm_spi_kr_funcs struct {
	GenerateRecoveryFields func(uint, uint64, uintptr, uint64, uintptr, uintptr, uint, uintptr) int
	GetRecoveredObject func(uint, int, uint, int, uintptr, uint, uintptr, uintptr) int
	PassThrough func(uint, uint64, uintptr, uint64, uintptr, uint, unsafe.Pointer, unsafe.Pointer) int
	ProcessRecoveryFields func(uint, uint64, uintptr, uint64, uintptr, uintptr, uint, uintptr) int
	RecoveryRequest func(uint, uint64, uintptr, uintptr, uintptr, []int, *int) int
	RecoveryRequestAbort func(uint, int) int
	RecoveryRetrieve func(uint, int, []int, *int, *uint) int
	RegistrationRequest func(uint, uint64, uintptr, uintptr, uintptr, uint, []int, *int) int
	RegistrationRetrieve func(uint, int, []int, uintptr) int

}

// Cssm_spi_tp_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_tp_funcs-c.struct
type Cssm_spi_tp_funcs struct {
	ApplyCrlToDb func(int, int, int, uintptr, uintptr, uintptr, uintptr) int
	CertCreateTemplate func(int, int, uint, uintptr, uintptr) int
	CertGetAllTemplateFields func(int, int, uintptr, *uint, uintptr) int
	CertGroupConstruct func(int, int, int, uintptr, unsafe.Pointer, uintptr, uintptr) int
	CertGroupPrune func(int, int, uintptr, uintptr, uintptr) int
	CertGroupToTupleGroup func(int, int, uintptr, uintptr) int
	CertGroupVerify func(int, int, int, uintptr, uintptr, uintptr) int
	CertReclaimAbort func(int, uint64) int
	CertReclaimKey func(int, uintptr, uint, uint64, int, uintptr) int
	CertRemoveFromCrlTemplate func(int, int, int, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) int
	CertRevoke func(int, int, int, uintptr, uintptr, uintptr, uintptr, uintptr, uint, uintptr) int
	CertSign func(int, int, uint64, uintptr, uintptr, uintptr, uintptr, uintptr) int
	ConfirmCredResult func(int, uintptr, uintptr, uintptr, uintptr) int
	CrlCreateTemplate func(int, int, uint, uintptr, uintptr) int
	CrlSign func(int, int, uint64, uintptr, uintptr, uintptr, uintptr, uintptr) int
	CrlVerify func(int, int, int, uintptr, uintptr, uintptr, uintptr) int
	FormRequest func(int, uintptr, uint, uintptr) int
	FormSubmit func(int, uint, uintptr, uintptr, uintptr, uintptr) int
	PassThrough func(int, int, uint64, uintptr, uint, unsafe.Pointer, unsafe.Pointer) int
	ReceiveConfirmation func(int, uintptr, uintptr, []int) int
	RetrieveCredResult func(int, uintptr, uintptr, []int, []int, uintptr) int
	SubmitCredRequest func(int, uintptr, uint, uintptr, uintptr, []int, uintptr) int
	TupleGroupToCertGroup func(int, int, uintptr, uintptr) int

}

// Cssm_state_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_state_funcs-c.struct
type Cssm_state_funcs struct {
	Cssm_DeliverModuleManagerEvent func(uintptr) int
	Cssm_DeregisterManagerServices func(uintptr) int
	Cssm_GetAppMemoryFunctions func(int, uintptr) int
	Cssm_GetAttachFunctions func(int, uint, unsafe.Pointer, uintptr, []int) int
	Cssm_IsFuncCallValid func(int, func(), func(), uint64, *uint64, uint, []int) int
	Cssm_ReleaseAttachFunctions func(int) int

}

// Cssm_subservice_uid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_subservice_uid-c.struct
type Cssm_subservice_uid struct {
	Guid unsafe.Pointer
	SubserviceId uint32
	SubserviceType CSSM_SERVICE_TYPE
	Version unsafe.Pointer

}

// Cssm_tp_authority_id
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_authority_id-c.struct
type Cssm_tp_authority_id struct {
	AuthorityCert unsafe.Pointer
	AuthorityLocation unsafe.Pointer

}

// Cssm_tp_callerauth_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_callerauth_context-c.struct
type Cssm_tp_callerauth_context struct {
	AnchorCerts unsafe.Pointer
	CallbackWithVerifiedCert unsafe.Pointer
	CallerCredentials unsafe.Pointer
	DBList unsafe.Pointer
	NumberOfAnchorCerts uint32
	Policy unsafe.Pointer
	VerificationAbortOn CSSM_TP_STOP_ON
	VerifyTime CSSM_TIMESTRING

}

// Cssm_tp_certchange_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certchange_input-c.struct
type Cssm_tp_certchange_input struct {
	Action CSSM_TP_CERTCHANGE_ACTION
	CLHandle CSSM_CL_HANDLE
	CallerCredentials unsafe.Pointer
	Cert unsafe.Pointer
	ChangeInfo unsafe.Pointer
	Reason CSSM_TP_CERTCHANGE_REASON
	StartTime CSSM_TIMESTRING

}

// Cssm_tp_certchange_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certchange_output-c.struct
type Cssm_tp_certchange_output struct {
	ActionStatus CSSM_TP_CERTCHANGE_STATUS
	RevokeInfo unsafe.Pointer

}

// Cssm_tp_certissue_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certissue_input-c.struct
type Cssm_tp_certissue_input struct {
	CLHandle CSSM_CL_HANDLE
	CSPSubserviceUid unsafe.Pointer
	MoreServiceRequests CSSM_TP_SERVICES
	NumberOfServiceControls uint32
	NumberOfTemplateFields uint32
	ServiceControls unsafe.Pointer
	SubjectCertFields unsafe.Pointer
	UserCredentials unsafe.Pointer

}

// Cssm_tp_certissue_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certissue_output-c.struct
type Cssm_tp_certissue_output struct {
	CertGroup unsafe.Pointer
	IssueStatus CSSM_TP_CERTISSUE_STATUS
	PerformedServiceRequests CSSM_TP_SERVICES

}

// Cssm_tp_certnotarize_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certnotarize_input-c.struct
type Cssm_tp_certnotarize_input struct {
	CLHandle CSSM_CL_HANDLE
	MoreFields unsafe.Pointer
	MoreServiceRequests CSSM_TP_SERVICES
	NumberOfFields uint32
	NumberOfServiceControls uint32
	ScopeSize uint32
	ServiceControls unsafe.Pointer
	SignScope unsafe.Pointer
	UserCredentials unsafe.Pointer

}

// Cssm_tp_certnotarize_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certnotarize_output-c.struct
type Cssm_tp_certnotarize_output struct {
	NotarizeStatus CSSM_TP_CERTNOTARIZE_STATUS
	NotarizedCertGroup unsafe.Pointer
	PerformedServiceRequests CSSM_TP_SERVICES

}

// Cssm_tp_certreclaim_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certreclaim_input-c.struct
type Cssm_tp_certreclaim_input struct {
	CLHandle CSSM_CL_HANDLE
	NumberOfSelectionFields uint32
	SelectionFields unsafe.Pointer
	UserCredentials unsafe.Pointer

}

// Cssm_tp_certreclaim_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certreclaim_output-c.struct
type Cssm_tp_certreclaim_output struct {
	KeyCacheHandle CSSM_LONG_HANDLE
	ReclaimStatus CSSM_TP_CERTRECLAIM_STATUS
	ReclaimedCertGroup unsafe.Pointer

}

// Cssm_tp_certverify_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certverify_input-c.struct
type Cssm_tp_certverify_input struct {
	CLHandle CSSM_CL_HANDLE
	Cert unsafe.Pointer
	VerifyContext unsafe.Pointer

}

// Cssm_tp_certverify_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certverify_output-c.struct
type Cssm_tp_certverify_output struct {
	Evidence unsafe.Pointer
	NumberOfEvidence uint32
	VerifyStatus CSSM_TP_CERTVERIFY_STATUS

}

// Cssm_tp_confirm_response
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_confirm_response-c.struct
type Cssm_tp_confirm_response struct {
	NumberOfResponses uint32
	Responses CSSM_TP_CONFIRM_STATUS_PTR

}

// Cssm_tp_crlissue_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_crlissue_input-c.struct
type Cssm_tp_crlissue_input struct {
	CLHandle CSSM_CL_HANDLE
	CallerCredentials unsafe.Pointer
	CrlIdentifier uint32
	CrlThisTime CSSM_TIMESTRING
	PolicyIdentifier unsafe.Pointer

}

// Cssm_tp_crlissue_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_crlissue_output-c.struct
type Cssm_tp_crlissue_output struct {
	Crl unsafe.Pointer
	CrlNextTime CSSM_TIMESTRING
	IssueStatus CSSM_TP_CRLISSUE_STATUS

}

// Cssm_tp_policyinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_policyinfo-c.struct
type Cssm_tp_policyinfo struct {
	NumberOfPolicyIds uint32
	PolicyControl unsafe.Pointer
	PolicyIds unsafe.Pointer

}

// Cssm_tp_request_set
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_request_set-c.struct
type Cssm_tp_request_set struct {
	NumberOfRequests uint32
	Requests unsafe.Pointer

}

// Cssm_tp_result_set
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_result_set-swift.struct
type Cssm_tp_result_set struct {
	NumberOfResults uint32
	Results unsafe.Pointer

}

// Cssm_tp_verify_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_verify_context-c.struct
type Cssm_tp_verify_context struct {
	Action CSSM_TP_ACTION
	ActionData unsafe.Pointer
	Cred unsafe.Pointer
	Crls unsafe.Pointer

}

// Cssm_tp_verify_context_result
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_verify_context_result-c.struct
type Cssm_tp_verify_context_result struct {
	Evidence unsafe.Pointer
	NumberOfEvidences uint32

}

// Cssm_tuplegroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tuplegroup-c.struct
type Cssm_tuplegroup struct {
	NumberOfTuples uint32
	Tuples unsafe.Pointer

}

// Cssm_upcalls
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_upcalls-c.struct
type Cssm_upcalls struct {
	CcToHandle_func func(uint64, *int) int
	GetModuleInfo_func func(int, uintptr, uintptr, *uint, *uint, *uint, *uint, uintptr, uintptr, uint) int
	Calloc_func unsafe.Pointer
	Free_func unsafe.Pointer
	Malloc_func unsafe.Pointer
	Realloc_func unsafe.Pointer

}

// Cssm_version
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_version-swift.struct
type Cssm_version struct {
	Major uint32
	Minor uint32

}

// Cssm_x509_extension
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_extension-c.struct
type Cssm_x509_extension struct {
	BERvalue unsafe.Pointer
	Critical CSSM_BOOL
	ExtnId unsafe.Pointer
	Format CSSM_X509EXT_DATA_FORMAT
	Value unsafe.Pointer

}

// Cssm_x509_extensionTagAndValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_extensionTagAndValue
type Cssm_x509_extensionTagAndValue struct {
	Type CSSM_BER_TAG
	Value unsafe.Pointer

}

// Cssm_x509_extensions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_extensions-c.struct
type Cssm_x509_extensions struct {
	Extensions unsafe.Pointer
	NumberOfExtensions uint32

}

// Cssm_x509_name
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_name-c.struct
type Cssm_x509_name struct {
	RelativeDistinguishedName unsafe.Pointer
	NumberOfRDNs uint32

}

// Cssm_x509_rdn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_rdn-c.struct
type Cssm_x509_rdn struct {
	AttributeTypeAndValue unsafe.Pointer
	NumberOfPairs uint32

}

// Cssm_x509_revoked_cert_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_revoked_cert_entry-c.struct
type Cssm_x509_revoked_cert_entry struct {
	CertificateSerialNumber unsafe.Pointer
	Extensions unsafe.Pointer
	RevocationDate unsafe.Pointer

}

// Cssm_x509_revoked_cert_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_revoked_cert_list-c.struct
type Cssm_x509_revoked_cert_list struct {
	NumberOfRevokedCertEntries uint32
	RevokedCertEntry unsafe.Pointer

}

// Cssm_x509_signature
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_signature-c.struct
type Cssm_x509_signature struct {
	AlgorithmIdentifier SecAsn1AlgId
	Encrypted unsafe.Pointer

}

// Cssm_x509_signed_certificate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_signed_certificate-c.struct
type Cssm_x509_signed_certificate struct {
	Certificate unsafe.Pointer
	Signature unsafe.Pointer

}

// Cssm_x509_signed_crl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_signed_crl-c.struct
type Cssm_x509_signed_crl struct {
	Signature unsafe.Pointer
	TbsCertList unsafe.Pointer

}

// Cssm_x509_tbs_certificate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_tbs_certificate-c.struct
type Cssm_x509_tbs_certificate struct {
	Extensions unsafe.Pointer
	Issuer unsafe.Pointer
	IssuerUniqueIdentifier unsafe.Pointer
	SerialNumber unsafe.Pointer
	Signature SecAsn1AlgId
	Subject unsafe.Pointer
	SubjectPublicKeyInfo SecAsn1PubKeyInfo
	SubjectUniqueIdentifier unsafe.Pointer
	Validity unsafe.Pointer
	Version unsafe.Pointer

}

// Cssm_x509_tbs_certlist
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_tbs_certlist-c.struct
type Cssm_x509_tbs_certlist struct {
	Extensions unsafe.Pointer
	Issuer unsafe.Pointer
	NextUpdate unsafe.Pointer
	RevokedCertificates unsafe.Pointer
	Signature SecAsn1AlgId
	ThisUpdate unsafe.Pointer
	Version unsafe.Pointer

}

// Cssm_x509_time
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_time-c.struct
type Cssm_x509_time struct {
	Time unsafe.Pointer
	TimeType CSSM_BER_TAG

}

// Cssm_x509_type_value_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_type_value_pair-c.struct
type Cssm_x509_type_value_pair struct {
	Type unsafe.Pointer
	Value unsafe.Pointer
	ValueType CSSM_BER_TAG

}

// Cssm_x509ext_basicConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_basicConstraints-c.struct
type Cssm_x509ext_basicConstraints struct {
	CA CSSM_BOOL
	PathLenConstraint uint32
	PathLenConstraintPresent CSSM_X509_OPTION

}

// Cssm_x509ext_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_pair-c.struct
type Cssm_x509ext_pair struct {
	ParsedValue unsafe.Pointer
	TagAndValue unsafe.Pointer

}

// Cssm_x509ext_policyInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_policyInfo-c.struct
type Cssm_x509ext_policyInfo struct {
	PolicyIdentifier unsafe.Pointer
	PolicyQualifiers unsafe.Pointer

}

// Cssm_x509ext_policyQualifierInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_policyQualifierInfo-c.struct
type Cssm_x509ext_policyQualifierInfo struct {
	PolicyQualifierId unsafe.Pointer
	Value unsafe.Pointer

}

// Cssm_x509ext_policyQualifiers
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_policyQualifiers-c.struct
type Cssm_x509ext_policyQualifiers struct {
	NumberOfPolicyQualifiers uint32
	PolicyQualifier unsafe.Pointer

}

// Mds_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/mds_funcs-c.struct
type Mds_funcs struct {
	CreateRelation func(Cssm_dl_db_handle, uint, *byte, uint, uintptr, uint, uintptr) int
	DataAbortQuery func(Cssm_dl_db_handle, int) int
	DataDelete func(Cssm_dl_db_handle, uintptr) int
	DataGetFirst func(Cssm_dl_db_handle, uintptr, *int, uintptr, uintptr, uintptr) int
	DataGetFromUniqueRecordId func(Cssm_dl_db_handle, uintptr, uintptr, uintptr) int
	DataGetNext func(Cssm_dl_db_handle, int, uintptr, uintptr, uintptr) int
	DataInsert func(Cssm_dl_db_handle, uint, uintptr, uintptr, uintptr) int
	DataModify func(Cssm_dl_db_handle, uint, uintptr, uintptr, uintptr, uint) int
	DbClose func(Cssm_dl_db_handle) int
	DbOpen func(int, *byte, uintptr, uint, uintptr, unsafe.Pointer, *int) int
	DestroyRelation func(Cssm_dl_db_handle, uint) int
	FreeNameList func(int, uintptr) int
	FreeUniqueRecord func(Cssm_dl_db_handle, uintptr) int
	GetDbNameFromHandle func(Cssm_dl_db_handle, *byte) int
	GetDbNames func(int, uintptr) int

}

// X509_validity
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/x509_validity
type X509_validity struct {
	NotAfter unsafe.Pointer
	NotBefore unsafe.Pointer

}

