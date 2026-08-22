// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
)

// C struct types

// AuthorizationCallbacks - The interface implemented by the Security Server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCallbacks
type AuthorizationCallbacks struct {
	Version               uint32                                                                                       // The engine callback version.
	SetResult             func(AuthorizationEngineRef, AuthorizationResult) int32                                      // Returns the result of an authorization operation.
	RequestInterrupt      func(AuthorizationEngineRef) int32                                                           // Requests the authorization engine to interrupt the currently active authorization mechanism.
	DidDeactivate         func(AuthorizationEngineRef) int32                                                           // Reports the successful deactivation of an authorization mechanism.
	GetContextValue       func(AuthorizationEngineRef, AuthorizationString, *AuthorizationContextFlags, uintptr) int32 // Reads a value collected during authorization.
	SetContextValue       func(AuthorizationEngineRef, AuthorizationString, uint, uintptr) int32                       // Stores data collected during authorization as a key-value pair.
	GetHintValue          func(AuthorizationEngineRef, AuthorizationString, uintptr) int32                             // Reads a value stored by the plug-in authorization mechanism.
	SetHintValue          func(AuthorizationEngineRef, AuthorizationString, uintptr) int32                             // Stores data needed during authorization as a key-value pair.
	GetArguments          func(AuthorizationEngineRef, uintptr) int32                                                  // Reads the arguments for this authorization mechanism from the authorization policy database.
	GetSessionId          func(AuthorizationEngineRef, unsafe.Pointer) int32                                           // Reads the session ID.
	GetImmutableHintValue func(AuthorizationEngineRef, AuthorizationString, uintptr) int32                             // Reads an immutable value stored by the plug-in authorization mechanism.
	GetLAContext          func(AuthorizationEngineRef, unsafe.Pointer) int32                                           // Constructs a local authentication context.
	GetTokenIdentities    func(AuthorizationEngineRef, unsafe.Pointer, uintptr) int32                                  // Returns an array of identities available on tokens.
	GetTKTokenWatcher     func(AuthorizationEngineRef, unsafe.Pointer) int32                                           // Constructs a token watcher.
	RemoveHintValue       func(AuthorizationEngineRef, AuthorizationString) int32                                      // Removes a value stored by the plug-in authorization mechanism.
	RemoveContextValue    func(AuthorizationEngineRef, AuthorizationString) int32                                      // Removes a value collected during authorization.

}

// AuthorizationExternalForm - The external representation of an authorization reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationExternalForm
type AuthorizationExternalForm struct {
	Bytes [32]int8 // An array of characters representing the external form of an authorization reference.

}

// AuthorizationItem - A structure containing information about an authorization right or the authorization environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationItem
type AuthorizationItem struct {
	Name        AuthorizationString // The required name of the authorization right or environment data.
	ValueLength uintptr             // The number of bytes in the value field.
	Value       unsafe.Pointer      // A pointer to information pertaining to the name field.
	Flags       uint32              // Reserved option bits.

}

// AuthorizationItemSet - A structure containing a set of authorization items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationItemSet
type AuthorizationItemSet struct {
	Count uint32             // The number of elements in the `items` array.
	Items *AuthorizationItem // A pointer to an array of authorization items.

}

// AuthorizationPluginInterface - The interface that must be implemented by your plug-in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationPluginInterface
type AuthorizationPluginInterface struct {
	Version             uint32                                                                                  // The plug-in interface version.
	PluginDestroy       func(unsafe.Pointer) int32                                                              // Notifies the plug-in that it is about to be unloaded.
	MechanismCreate     func(unsafe.Pointer, AuthorizationEngineRef, AuthorizationString, unsafe.Pointer) int32 // Creates an authorization mechanism.
	MechanismInvoke     func(unsafe.Pointer) int32                                                              // Invokes an authorization mechanism to perform an authorization operation.
	MechanismDeactivate func(unsafe.Pointer) int32                                                              // Deactivates an authorization mechanism.
	MechanismDestroy    func(unsafe.Pointer) int32                                                              // Destroys an authorization mechanism.

}

// AuthorizationValue - A structure used to pass data between the authorization engine and the plug-in mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationValue
type AuthorizationValue struct {
	Length uintptr
	Data   unsafe.Pointer
}

// AuthorizationValueVector - A structure used to pass arguments from the authorization policy database to the authorization mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationValueVector
type AuthorizationValueVector struct {
	Count  uint32
	Values *AuthorizationValue
}

// CE_AccessDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_AccessDescription
type CE_AccessDescription struct {
	AccessMethod   [2]uint64
	AccessLocation CE_GeneralName
}

// CE_AuthorityInfoAccess
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_AuthorityInfoAccess
type CE_AuthorityInfoAccess struct {
	NumAccessDescriptions uint32
	AccessDescriptions    *CE_AccessDescription
}

// CE_AuthorityKeyID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_AuthorityKeyID
type CE_AuthorityKeyID struct {
	KeyIdentifierPresent CSSM_BOOL
	KeyIdentifier        SecAsn1Item
	GeneralNamesPresent  CSSM_BOOL
	GeneralNames         *CE_GeneralNames
	SerialNumberPresent  CSSM_BOOL
	SerialNumber         SecAsn1Item
}

// CE_BasicConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_BasicConstraints
type CE_BasicConstraints struct {
	CA                       CSSM_BOOL
	PathLenConstraintPresent CSSM_BOOL
	PathLenConstraint        uint32
}

// CE_CRLDistPointsSyntax
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CRLDistPointsSyntax
type CE_CRLDistPointsSyntax struct {
	NumDistPoints uint32
	DistPoints    *CE_CRLDistributionPoint
}

// CE_CRLDistributionPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CRLDistributionPoint
type CE_CRLDistributionPoint struct {
	DistPointName  *CE_DistributionPointName
	ReasonsPresent CSSM_BOOL
	Reasons        uint8
	CrlIssuer      *CE_GeneralNames
}

// CE_CertPolicies
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CertPolicies
type CE_CertPolicies struct {
	NumPolicies uint32
	Policies    *CE_PolicyInformation
}

// CE_Data is a C union type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_Data
type CE_Data [8]uint64

// AuthorityKeyID returns the union interpreted as *CE_AuthorityKeyID.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) AuthorityKeyID() *CE_AuthorityKeyID {
	return (*CE_AuthorityKeyID)(unsafe.Pointer(u))
}

// SubjectKeyID returns the union interpreted as *CE_SubjectKeyID.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) SubjectKeyID() *CE_SubjectKeyID {
	return (*CE_SubjectKeyID)(unsafe.Pointer(u))
}

// KeyUsage returns the union interpreted as *CE_KeyUsage.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) KeyUsage() *CE_KeyUsage {
	return (*CE_KeyUsage)(unsafe.Pointer(u))
}

// SubjectAltName returns the union interpreted as *CE_GeneralNames.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) SubjectAltName() *CE_GeneralNames {
	return (*CE_GeneralNames)(unsafe.Pointer(u))
}

// IssuerAltName returns the union interpreted as *CE_GeneralNames.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) IssuerAltName() *CE_GeneralNames {
	return (*CE_GeneralNames)(unsafe.Pointer(u))
}

// ExtendedKeyUsage returns the union interpreted as *CE_ExtendedKeyUsage.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) ExtendedKeyUsage() *CE_ExtendedKeyUsage {
	return (*CE_ExtendedKeyUsage)(unsafe.Pointer(u))
}

// BasicConstraints returns the union interpreted as *CE_BasicConstraints.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) BasicConstraints() *CE_BasicConstraints {
	return (*CE_BasicConstraints)(unsafe.Pointer(u))
}

// CertPolicies returns the union interpreted as *CE_CertPolicies.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) CertPolicies() *CE_CertPolicies {
	return (*CE_CertPolicies)(unsafe.Pointer(u))
}

// NetscapeCertType returns the union interpreted as *CE_NetscapeCertType.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) NetscapeCertType() *CE_NetscapeCertType {
	return (*CE_NetscapeCertType)(unsafe.Pointer(u))
}

// CrlNumber returns the union interpreted as *CE_CrlNumber.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) CrlNumber() *CE_CrlNumber {
	return (*CE_CrlNumber)(unsafe.Pointer(u))
}

// DeltaCrl returns the union interpreted as *CE_DeltaCrl.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) DeltaCrl() *CE_DeltaCrl {
	return (*CE_DeltaCrl)(unsafe.Pointer(u))
}

// CrlReason returns the union interpreted as *CE_CrlReason.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) CrlReason() *CE_CrlReason {
	return (*CE_CrlReason)(unsafe.Pointer(u))
}

// CrlDistPoints returns the union interpreted as *CE_CRLDistPointsSyntax.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) CrlDistPoints() *CE_CRLDistPointsSyntax {
	return (*CE_CRLDistPointsSyntax)(unsafe.Pointer(u))
}

// IssuingDistPoint returns the union interpreted as *CE_IssuingDistributionPoint.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) IssuingDistPoint() *CE_IssuingDistributionPoint {
	return (*CE_IssuingDistributionPoint)(unsafe.Pointer(u))
}

// AuthorityInfoAccess returns the union interpreted as *CE_AuthorityInfoAccess.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) AuthorityInfoAccess() *CE_AuthorityInfoAccess {
	return (*CE_AuthorityInfoAccess)(unsafe.Pointer(u))
}

// QualifiedCertStatements returns the union interpreted as *CE_QC_Statements.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) QualifiedCertStatements() *CE_QC_Statements {
	return (*CE_QC_Statements)(unsafe.Pointer(u))
}

// NameConstraints returns the union interpreted as *CE_NameConstraints.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) NameConstraints() *CE_NameConstraints {
	return (*CE_NameConstraints)(unsafe.Pointer(u))
}

// PolicyMappings returns the union interpreted as *CE_PolicyMappings.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) PolicyMappings() *CE_PolicyMappings {
	return (*CE_PolicyMappings)(unsafe.Pointer(u))
}

// PolicyConstraints returns the union interpreted as *CE_PolicyConstraints.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) PolicyConstraints() *CE_PolicyConstraints {
	return (*CE_PolicyConstraints)(unsafe.Pointer(u))
}

// InhibitAnyPolicy returns the union interpreted as *CE_InhibitAnyPolicy.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) InhibitAnyPolicy() *CE_InhibitAnyPolicy {
	return (*CE_InhibitAnyPolicy)(unsafe.Pointer(u))
}

// RawData returns the union interpreted as *SecAsn1Item.
// The returned pointer aliases the receiver's memory.
func (u *CE_Data) RawData() *SecAsn1Item {
	return (*SecAsn1Item)(unsafe.Pointer(u))
}

// CE_DataAndType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_DataAndType
type CE_DataAndType struct {
	Type      CE_DataType
	Extension CE_Data
	Critical  CSSM_BOOL
}

// CE_DistributionPointName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_DistributionPointName
type CE_DistributionPointName struct {
	NameType CE_CrlDistributionPointNameType
	Dpn      [1]uint64
}

// CE_ExtendedKeyUsage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_ExtendedKeyUsage-c.struct
type CE_ExtendedKeyUsage struct {
	NumPurposes uint32
	Purposes    unsafe.Pointer
}

// CE_GeneralName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralName
type CE_GeneralName struct {
	NameType   CE_GeneralNameType
	BerEncoded CSSM_BOOL
	Name       SecAsn1Item
}

// CE_GeneralNames
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralNames
type CE_GeneralNames struct {
	NumNames    uint32
	GeneralName *CE_GeneralName
}

// CE_GeneralSubtree
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralSubtree
type CE_GeneralSubtree struct {
	Base           *CE_GeneralNames
	Minimum        uint32
	MaximumPresent CSSM_BOOL
	Maximum        uint32
}

// CE_GeneralSubtrees
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralSubtrees
type CE_GeneralSubtrees struct {
	NumSubtrees uint32
	Subtrees    *CE_GeneralSubtree
}

// CE_IssuingDistributionPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_IssuingDistributionPoint
type CE_IssuingDistributionPoint struct {
	DistPointName          *CE_DistributionPointName
	OnlyUserCertsPresent   CSSM_BOOL
	OnlyUserCerts          CSSM_BOOL
	OnlyCACertsPresent     CSSM_BOOL
	OnlyCACerts            CSSM_BOOL
	OnlySomeReasonsPresent CSSM_BOOL
	OnlySomeReasons        uint8
	IndirectCrlPresent     CSSM_BOOL
	IndirectCrl            CSSM_BOOL
}

// CE_NameConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_NameConstraints
type CE_NameConstraints struct {
	Permitted *CE_GeneralSubtrees
	Excluded  *CE_GeneralSubtrees
}

// CE_OtherName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_OtherName
type CE_OtherName struct {
	TypeId [2]uint64
	Value  SecAsn1Item
}

// CE_PolicyConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyConstraints
type CE_PolicyConstraints struct {
	RequireExplicitPolicyPresent CSSM_BOOL
	RequireExplicitPolicy        uint32
	InhibitPolicyMappingPresent  CSSM_BOOL
	InhibitPolicyMapping         uint32
}

// CE_PolicyInformation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyInformation
type CE_PolicyInformation struct {
	CertPolicyId        [2]uint64
	NumPolicyQualifiers uint32
	PolicyQualifiers    *CE_PolicyQualifierInfo
}

// CE_PolicyMapping
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyMapping
type CE_PolicyMapping struct {
	IssuerDomainPolicy  [2]uint64
	SubjectDomainPolicy [2]uint64
}

// CE_PolicyMappings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyMappings
type CE_PolicyMappings struct {
	NumPolicyMappings uint32
	PolicyMappings    *CE_PolicyMapping
}

// CE_PolicyQualifierInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyQualifierInfo
type CE_PolicyQualifierInfo struct {
	PolicyQualifierId [2]uint64
	Qualifier         SecAsn1Item
}

// CE_QC_Statement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_QC_Statement
type CE_QC_Statement struct {
	StatementId   [2]uint64
	SemanticsInfo *CE_SemanticsInformation
	OtherInfo     unsafe.Pointer
}

// CE_QC_Statements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_QC_Statements
type CE_QC_Statements struct {
	NumQCStatements uint32
	QcStatements    *CE_QC_Statement
}

// CE_SemanticsInformation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_SemanticsInformation
type CE_SemanticsInformation struct {
	SemanticsIdentifier         unsafe.Pointer
	NameRegistrationAuthorities unsafe.Pointer
}

// CSSM_APPLE_CL_CSR_REQUEST
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_CL_CSR_REQUEST
type CSSM_APPLE_CL_CSR_REQUEST struct {
	SubjectNameX509   unsafe.Pointer
	SignatureAlg      CSSM_ALGORITHMS
	SignatureOid      [2]uint64
	CspHand           CSSM_CSP_HANDLE
	SubjectPublicKey  unsafe.Pointer
	SubjectPrivateKey unsafe.Pointer
	ChallengeString   *byte
}

// CSSM_APPLE_TP_ACTION_DATA
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_ACTION_DATA
type CSSM_APPLE_TP_ACTION_DATA struct {
	Version     uint32
	ActionFlags CSSM_APPLE_TP_ACTION_FLAGS
}

// CSSM_APPLE_TP_CERT_REQUEST
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_CERT_REQUEST
type CSSM_APPLE_TP_CERT_REQUEST struct {
	CspHand          CSSM_CSP_HANDLE
	ClHand           CSSM_CL_HANDLE
	SerialNumber     uint32
	NumSubjectNames  uint32
	SubjectNames     *CSSM_APPLE_TP_NAME_OID
	NumIssuerNames   uint32
	IssuerNames      *CSSM_APPLE_TP_NAME_OID
	IssuerNameX509   unsafe.Pointer
	CertPublicKey    unsafe.Pointer
	IssuerPrivateKey unsafe.Pointer
	SignatureAlg     CSSM_ALGORITHMS
	SignatureOid     [2]uint64
	NotBefore        uint32
	NotAfter         uint32
	NumExtensions    uint32
	Extensions       *CE_DataAndType
	ChallengeString  *byte
}

// CSSM_APPLE_TP_CRL_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_CRL_OPTIONS
type CSSM_APPLE_TP_CRL_OPTIONS struct {
	Version  uint32
	CrlFlags CSSM_APPLE_TP_CRL_OPT_FLAGS
	CrlStore unsafe.Pointer
}

// CSSM_APPLE_TP_NAME_OID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_NAME_OID
type CSSM_APPLE_TP_NAME_OID struct {
	String *byte
	Oid    unsafe.Pointer
}

// CSSM_APPLE_TP_SMIME_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_SMIME_OPTIONS
type CSSM_APPLE_TP_SMIME_OPTIONS struct {
	Version        uint32
	IntendedUsage  CE_KeyUsage
	SenderEmailLen uint32
	SenderEmail    *byte
}

// CSSM_APPLE_TP_SSL_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_SSL_OPTIONS
type CSSM_APPLE_TP_SSL_OPTIONS struct {
	Version       uint32
	ServerNameLen uint32
	ServerName    *byte
	Flags         uint32
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
	StatusBits     CSSM_TP_APPLE_CERT_STATUS
	NumStatusCodes uint32
	StatusCodes    unsafe.Pointer
	Index          uint32
	DlDbHandle     [2]uint64
	UniqueRecord   unsafe.Pointer
	CrlReason      int32
}

// CSSM_TUPLE
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TUPLE
type CSSM_TUPLE struct {
	Issuer           [3]uint64
	Subject          [3]uint64
	Delegate         CSSM_BOOL
	AuthorizationTag [3]uint64
	ValidityPeriod   [3]uint64
}

// SecAsn1AlgId - A structure identifying an ASN.1 algorithm by its OID, and its corresponding parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1AlgId
type SecAsn1AlgId struct {
	Algorithm  [2]uint64
	Parameters SecAsn1Item
}

// SecAsn1PubKeyInfo - A structure containing a public key and its associated algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1PubKeyInfo
type SecAsn1PubKeyInfo struct {
	Algorithm        SecAsn1AlgId
	SubjectPublicKey SecAsn1Item
}

// SecAsn1Template_struct - A structure that defines one element of a BER or DER encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1Template_struct
type SecAsn1Template_struct struct {
	Kind   uint32
	Offset uint32
	Sub    unsafe.Pointer
	Size   uint32
}

// SecItemImportExportKeyParameters - The import/export parameter structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemImportExportKeyParameters
type SecItemImportExportKeyParameters struct {
	Version       uint32                     // The version of this structure.
	Flags         SecKeyImportExportFlags    // The bitwise [OR] of zero or more key import/export flags.
	Passphrase    corefoundation.CFTypeRef   // The password to use during key import or export.
	AlertTitle    corefoundation.CFStringRef // The title to display in the secure passphrase alert panel.
	AlertPrompt   corefoundation.CFStringRef // The prompt to display in the secure passphrase alert panel.
	AccessRef     SecAccessRef               // Specifies the initial access controls of imported private keys.
	KeyUsage      corefoundation.CFArrayRef  // An array containing usage attributes applied to a key on import.
	KeyAttributes corefoundation.CFArrayRef  // An array containing zero or more key attributes for an imported key.

}

// SecKeyImportExportParameters - The legacy import/export parameter structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyImportExportParameters
type SecKeyImportExportParameters struct {
	Version       uint32                     // The version of this structure.
	Flags         SecKeyImportExportFlags    // The bitwise [OR] of zero or more key import/export flags.
	Passphrase    corefoundation.CFTypeRef   // The password to use during key import or export.
	AlertTitle    corefoundation.CFStringRef // The title to display in the secure passphrase alert panel.
	AlertPrompt   corefoundation.CFStringRef // The prompt to display in the secure passphrase alert panel.
	AccessRef     SecAccessRef               // Specifies the initial access controls of imported private keys.
	KeyUsage      CSSM_KEYUSE                // A word of bits constituting the low-level use flags for imported keys.
	KeyAttributes CSSM_KEYATTR_FLAGS         // A word of bits constituting the low-level attribute flags for imported keys.

}

// SecKeychainAttribute - A structure that holds a single keychain attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttribute
type SecKeychainAttribute struct {
	Tag    SecKeychainAttrType // A 4-byte attribute tag.
	Length uint32              // The length of the buffer pointed to by data.
	Data   unsafe.Pointer      // A pointer to the attribute data.

}

// SecKeychainAttributeInfo - A structure that represents an attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttributeInfo
type SecKeychainAttributeInfo struct {
	Count  uint32  // The number of tag-format pairs in the respective arrays.
	Tag    *uint32 // A pointer to the first attribute tag in the array.
	Format *uint32 // A pointer to the first attribute format in the array.

}

// SecKeychainAttributeList - A list of keychain attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttributeList
type SecKeychainAttributeList struct {
	Count uint32                // The number of keychain attributes in the array.
	Attr  *SecKeychainAttribute // A pointer to the first keychain attribute in the array.

}

// SecKeychainCallbackInfo - Information about a keychain event that keychain services deliver to your app via a callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainCallbackInfo
type SecKeychainCallbackInfo struct {
	Version  uint32             // The version of this structure.
	Item     SecKeychainItemRef // A reference to the keychain item in which the event occurred. If the event did not involve an item, this field is not valid.
	Keychain SecKeychainRef     // A reference to the keychain in which the event occurred. If the event did not involve a keychain, this field is not valid.
	Pid      int32              // The ID of the process that generated this event.

}

// SecKeychainSettings - A structure that contains information about keychain settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSettings
type SecKeychainSettings struct {
	Version         uint32 // The keychain version.
	LockOnSleep     bool   // A Boolean value indicating whether the keychain locks when the system sleeps.
	UseLockInterval bool   // A Boolean value indicating whether the keychain automatically locks after a certain period of time.
	LockInterval    uint32 // The number of seconds to wait before the keychain locks.

}

// Cssm_access_credentials
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_access_credentials-c.struct
type Cssm_access_credentials struct {
	EntryTag  CSSM_STRING
	BaseCerts [7]uint64
	Samples   [2]uint64
	Callback  unsafe.Pointer
	CallerCtx unsafe.Pointer
}

// Cssm_acl_edit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_edit-c.struct
type Cssm_acl_edit struct {
	EditMode       CSSM_ACL_EDIT_MODE
	OldEntryHandle CSSM_ACL_HANDLE
	NewEntry       unsafe.Pointer
}

// Cssm_acl_entry_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_entry_info-c.struct
type Cssm_acl_entry_info struct {
	EntryPublicInfo [19]uint64
	EntryHandle     CSSM_ACL_HANDLE
}

// Cssm_acl_entry_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_entry_input-c.struct
type Cssm_acl_entry_input struct {
	Prototype     [19]uint64
	Callback      unsafe.Pointer
	CallerContext unsafe.Pointer
}

// Cssm_acl_entry_prototype
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_entry_prototype-c.struct
type Cssm_acl_entry_prototype struct {
	TypedSubject  [3]uint64
	Delegate      CSSM_BOOL
	Authorization [2]uint64
	TimeRange     [4]uint64
	EntryTag      CSSM_STRING
}

// Cssm_acl_keychain_prompt_selector
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_keychain_prompt_selector-swift.struct
type Cssm_acl_keychain_prompt_selector struct {
	Version uint16
	Flags   uint16
}

// Cssm_acl_owner_prototype
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_owner_prototype-c.struct
type Cssm_acl_owner_prototype struct {
	TypedSubject [3]uint64
	Delegate     CSSM_BOOL
}

// Cssm_acl_process_subject_selector
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_process_subject_selector-swift.struct
type Cssm_acl_process_subject_selector struct {
	Version uint16
	Mask    uint16
	Uid     uint32
	Gid     uint32
}

// Cssm_acl_validity_period
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_validity_period-c.struct
type Cssm_acl_validity_period struct {
	StartDate SecAsn1Item
	EndDate   SecAsn1Item
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
	Length     uint32
	Version    uint32
	AutoCommit CSSM_BOOL
	Mask       uint32
	Mode       uint16
}

// Cssm_authorizationgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_authorizationgroup-swift.struct
type Cssm_authorizationgroup struct {
	NumberOfAuthTags uint32
	AuthTags         unsafe.Pointer
}

// Cssm_base_certs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_base_certs-c.struct
type Cssm_base_certs struct {
	TPHandle CSSM_TP_HANDLE
	CLHandle CSSM_CL_HANDLE
	Certs    [5]uint64
}

// Cssm_cert_bundle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_cert_bundle-c.struct
type Cssm_cert_bundle struct {
	BundleHeader [2]uint32
	Bundle       SecAsn1Item
}

// Cssm_cert_bundle_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_cert_bundle_header-c.struct
type Cssm_cert_bundle_header struct {
	BundleType     CSSM_CERT_BUNDLE_TYPE
	BundleEncoding CSSM_CERT_BUNDLE_ENCODING
}

// Cssm_cert_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_cert_pair-c.struct
type Cssm_cert_pair struct {
	EncodedCert [3]uint64
	ParsedCert  [2]uint64
}

// Cssm_certgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_certgroup-c.struct
type Cssm_certgroup struct {
	CertEncoding    CSSM_CERT_ENCODING
	CertGroupType   CSSM_CERTGROUP_TYPE
	CertType        CSSM_CERT_TYPE
	GroupList       unsafe.Pointer
	NumCerts        uint32
	Reserved        unsafe.Pointer
	PairCertList    unsafe.Pointer
	CertList        unsafe.Pointer
	EncodedCertList unsafe.Pointer
	ParsedCertList  unsafe.Pointer
}

// Cssm_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_context-c.struct
type Cssm_context struct {
	ContextType          CSSM_CONTEXT_TYPE
	AlgorithmType        CSSM_ALGORITHMS
	NumberOfAttributes   uint32
	ContextAttributes    unsafe.Pointer
	CSPHandle            CSSM_CSP_HANDLE
	Privileged           CSSM_BOOL
	EncryptionProhibited uint32
	WorkFactor           uint32
	Reserved             uint32
}

// Cssm_context_attribute
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_context_attribute-c.struct
type Cssm_context_attribute struct {
	AttributeType   CSSM_ATTRIBUTE_TYPE
	AttributeLength uint32
	Attribute       unsafe.Pointer
}

// Cssm_crl_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_crl_pair-c.struct
type Cssm_crl_pair struct {
	EncodedCrl [3]uint64
	ParsedCrl  [2]uint64
}

// Cssm_crlgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_crlgroup-c.struct
type Cssm_crlgroup struct {
	CrlEncoding    CSSM_CRL_ENCODING
	CrlGroupType   CSSM_CRLGROUP_TYPE
	CrlType        CSSM_CRL_TYPE
	GroupCrlList   unsafe.Pointer
	NumberOfCrls   uint32
	CrlList        unsafe.Pointer
	EncodedCrlList unsafe.Pointer
	PairCrlList    unsafe.Pointer
	ParsedCrlList  unsafe.Pointer
}

// Cssm_crypto_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_crypto_data-c.struct
type Cssm_crypto_data struct {
	Param     SecAsn1Item
	Callback  unsafe.Pointer
	CallerCtx unsafe.Pointer
}

// Cssm_csp_operational_statistics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_csp_operational_statistics-swift.struct
type Cssm_csp_operational_statistics struct {
	UserAuthenticated         CSSM_BOOL
	DeviceFlags               CSSM_CSP_FLAGS
	TokenMaxSessionCount      uint32
	TokenOpenedSessionCount   uint32
	TokenMaxRWSessionCount    uint32
	TokenOpenedRWSessionCount uint32
	TokenTotalPublicMem       uint32
	TokenFreePublicMem        uint32
	TokenTotalPrivateMem      uint32
	TokenFreePrivateMem       uint32
}

// Cssm_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_data-swift.struct
type Cssm_data struct {
	Length uintptr
	Data   *uint8
}

// Cssm_date
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_date-swift.struct
type Cssm_date struct {
	Year  [4]uint8
	Month [2]uint8
	Day   [2]uint8
}

// Cssm_db_attribute_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_attribute_data-c.struct
type Cssm_db_attribute_data struct {
	Info           [4]uint64
	NumberOfValues uint32
	Value          unsafe.Pointer
}

// Cssm_db_attribute_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_attribute_info-c.struct
type Cssm_db_attribute_info struct {
	AttributeNameFormat CSSM_DB_ATTRIBUTE_NAME_FORMAT
	Label               unsafe.Pointer
	AttributeFormat     CSSM_DB_ATTRIBUTE_FORMAT
}

// Cssm_db_index_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_index_info-c.struct
type Cssm_db_index_info struct {
	IndexType           CSSM_DB_INDEX_TYPE
	IndexedDataLocation CSSM_DB_INDEXED_DATA_LOCATION
	Info                [4]uint64
}

// Cssm_db_parsing_module_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_parsing_module_info-c.struct
type Cssm_db_parsing_module_info struct {
	RecordType          CSSM_DB_RECORDTYPE
	ModuleSubserviceUid [8]uint32
}

// Cssm_db_record_attribute_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_record_attribute_data-c.struct
type Cssm_db_record_attribute_data struct {
	DataRecordType      CSSM_DB_RECORDTYPE
	SemanticInformation uint32
	NumberOfAttributes  uint32
	AttributeData       unsafe.Pointer
}

// Cssm_db_record_attribute_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_record_attribute_info-c.struct
type Cssm_db_record_attribute_info struct {
	DataRecordType     CSSM_DB_RECORDTYPE
	NumberOfAttributes uint32
	AttributeInfo      unsafe.Pointer
}

// Cssm_db_record_index_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_record_index_info-c.struct
type Cssm_db_record_index_info struct {
	DataRecordType  CSSM_DB_RECORDTYPE
	NumberOfIndexes uint32
	IndexInfo       unsafe.Pointer
}

// Cssm_db_schema_attribute_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_schema_attribute_info-c.struct
type Cssm_db_schema_attribute_info struct {
	AttributeId     uint32
	AttributeName   *byte
	AttributeNameID [2]uint64
	DataType        CSSM_DB_ATTRIBUTE_FORMAT
}

// Cssm_db_schema_index_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_schema_index_info-swift.struct
type Cssm_db_schema_index_info struct {
	AttributeId         uint32
	IndexId             uint32
	IndexType           CSSM_DB_INDEX_TYPE
	IndexedDataLocation CSSM_DB_INDEXED_DATA_LOCATION
}

// Cssm_db_unique_record
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_unique_record-c.struct
type Cssm_db_unique_record struct {
	RecordLocator    [5]uint64
	RecordIdentifier SecAsn1Item
}

// Cssm_dbinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_dbinfo-c.struct
type Cssm_dbinfo struct {
	NumberOfRecordTypes   uint32
	DefaultParsingModules unsafe.Pointer
	RecordAttributeNames  unsafe.Pointer
	RecordIndexes         unsafe.Pointer
	IsLocal               CSSM_BOOL
	AccessPath            *byte
	Reserved              unsafe.Pointer
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
	NumHandles uint32
	DLDBHandle unsafe.Pointer
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
	CertType     CSSM_CERT_TYPE
	CertEncoding CSSM_CERT_ENCODING
	CertBlob     SecAsn1Item
}

// Cssm_encoded_crl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_encoded_crl-c.struct
type Cssm_encoded_crl struct {
	CrlType     CSSM_CRL_TYPE
	CrlEncoding CSSM_CRL_ENCODING
	CrlBlob     SecAsn1Item
}

// Cssm_evidence
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_evidence-c.struct
type Cssm_evidence struct {
	EvidenceForm CSSM_EVIDENCE_FORM
	Evidence     unsafe.Pointer
}

// Cssm_field
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_field-c.struct
type Cssm_field struct {
	FieldOid   [2]uint64
	FieldValue SecAsn1Item
}

// Cssm_fieldgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_fieldgroup-c.struct
type Cssm_fieldgroup struct {
	NumberOfFields int32
	Fields         unsafe.Pointer
}

// Cssm_func_name_addr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_func_name_addr-swift.struct
type Cssm_func_name_addr struct {
	Name    CSSM_STRING
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
	Data4 [8]uint8
}

// Cssm_kea_derive_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kea_derive_params-c.struct
type Cssm_kea_derive_params struct {
	Rb SecAsn1Item
	Yb SecAsn1Item
}

// Cssm_key
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_key-c.struct
type Cssm_key struct {
	KeyHeader [19]uint32
	KeyData   SecAsn1Item
}

// Cssm_key_size
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_key_size-swift.struct
type Cssm_key_size struct {
	LogicalKeySizeInBits   uint32
	EffectiveKeySizeInBits uint32
}

// Cssm_keyheader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_keyheader-c.struct
type Cssm_keyheader struct {
	HeaderVersion        CSSM_HEADERVERSION
	CspId                [4]uint32
	BlobType             CSSM_KEYBLOB_TYPE
	Format               CSSM_KEYBLOB_FORMAT
	AlgorithmId          CSSM_ALGORITHMS
	KeyClass             CSSM_KEYCLASS
	LogicalKeySizeInBits uint32
	KeyAttr              CSSM_KEYATTR_FLAGS
	KeyUsage             CSSM_KEYUSE
	StartDate            [8]byte
	EndDate              [8]byte
	WrapAlgorithmId      CSSM_ALGORITHMS
	WrapMode             CSSM_ENCRYPT_MODE
	Reserved             uint32
}

// Cssm_kr_name
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_name-swift.struct
type Cssm_kr_name struct {
	Type   uint8
	Length uint8
	Name   *byte
}

// Cssm_kr_policy_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_policy_info-c.struct
type Cssm_kr_policy_info struct {
	KrbNotAllowed   CSSM_BOOL
	NumberOfEntries uint32
	PolicyEntry     unsafe.Pointer
}

// Cssm_kr_policy_list_item
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_policy_list_item-c.struct
type Cssm_kr_policy_list_item struct {
	Next         unsafe.Pointer
	AlgorithmId  CSSM_ALGORITHMS
	Mode         CSSM_ENCRYPT_MODE
	MaxKeyLength uint32
	MaxRounds    uint32
	WorkFactor   uint8
	PolicyFlags  CSSM_KR_POLICY_FLAGS
	AlgClass     CSSM_CONTEXT_TYPE
}

// Cssm_kr_profile
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_profile-c.struct
type Cssm_kr_profile struct {
	UserName                 [2]uint64
	UserCertificate          unsafe.Pointer
	KRSCertChain             unsafe.Pointer
	LE_KRANum                uint8
	LE_KRACertChainList      unsafe.Pointer
	ENT_KRANum               uint8
	ENT_KRACertChainList     unsafe.Pointer
	INDIV_KRANum             uint8
	INDIV_KRACertChainList   unsafe.Pointer
	INDIV_AuthenticationInfo unsafe.Pointer
	KRSPFlags                uint32
	KRSPExtensions           unsafe.Pointer
}

// Cssm_kr_wrappedproductinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_wrappedproductinfo
type Cssm_kr_wrappedproductinfo struct {
	StandardVersion     [2]uint32
	StandardDescription CSSM_STRING
	ProductVersion      [2]uint32
	ProductDescription  CSSM_STRING
	ProductVendor       CSSM_STRING
	ProductFlags        uint32
}

// Cssm_krsubservice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_krsubservice-c.struct
type Cssm_krsubservice struct {
	SubServiceId   uint32
	Description    *byte
	WrappedProduct [56]uint32
}

// Cssm_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_list-swift.struct
type Cssm_list struct {
	ListType CSSM_LIST_TYPE
	Head     CSSM_LIST_ELEMENT_PTR
	Tail     CSSM_LIST_ELEMENT_PTR
}

// Cssm_list_element
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_list_element-c.struct
type Cssm_list_element struct {
	Element     unsafe.Pointer
	ElementType CSSM_LIST_ELEMENT_TYPE
	NextElement *Cssm_list_element
	WordID      CSSM_WORDID_TYPE
	Sublist     [3]uint64
	Word        SecAsn1Item
}

// Cssm_manager_event_notification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_manager_event_notification-c.struct
type Cssm_manager_event_notification struct {
	DestinationModuleManagerType CSSM_SERVICE_MASK
	SourceModuleManagerType      CSSM_SERVICE_MASK
	Event                        CSSM_MANAGER_EVENT_TYPES
	EventId                      uint32
	EventData                    SecAsn1Item
}

// Cssm_manager_registration_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_manager_registration_info-c.struct
type Cssm_manager_registration_info struct {
	Initialize              func(uint32, uint32) int32
	Terminate               func() int32
	RegisterDispatchTable   func(uintptr) int32
	DeregisterDispatchTable func() int32
	EventNotifyManager      func(uintptr) int32
	RefreshFunctionTable    func(uintptr, uint32) int32
}

// Cssm_memory_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_memory_funcs-swift.struct
type Cssm_memory_funcs struct {
	Malloc_func  CSSM_MALLOC
	Free_func    CSSM_FREE
	Realloc_func CSSM_REALLOC
	Calloc_func  CSSM_CALLOC
	AllocRef     unsafe.Pointer
}

// Cssm_module_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_module_funcs-c.struct
type Cssm_module_funcs struct {
	ServiceType          CSSM_SERVICE_TYPE
	NumberOfServiceFuncs uint32
	ServiceFuncs         unsafe.Pointer
}

// Cssm_name_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_name_list-swift.struct
type Cssm_name_list struct {
	NumStrings uint32
	String     *byte
}

// Cssm_net_address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_net_address-c.struct
type Cssm_net_address struct {
	AddressType CSSM_NET_ADDRESS_TYPE
	Address     SecAsn1Item
}

// Cssm_parsed_cert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_parsed_cert-swift.struct
type Cssm_parsed_cert struct {
	CertType         CSSM_CERT_TYPE
	ParsedCertFormat CSSM_CERT_PARSE_FORMAT
	ParsedCert       unsafe.Pointer
}

// Cssm_parsed_crl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_parsed_crl-swift.struct
type Cssm_parsed_crl struct {
	CrlType         CSSM_CRL_TYPE
	ParsedCrlFormat CSSM_CRL_PARSE_FORMAT
	ParsedCrl       unsafe.Pointer
}

// Cssm_pkcs1_oaep_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_pkcs1_oaep_params-c.struct
type Cssm_pkcs1_oaep_params struct {
	HashAlgorithm uint32
	HashParams    SecAsn1Item
	MGF           CSSM_PKCS_OAEP_MGF
	MGFParams     SecAsn1Item
	PSource       CSSM_PKCS_OAEP_PSOURCE
	PSourceParams SecAsn1Item
}

// Cssm_pkcs5_pbkdf1_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_pkcs5_pbkdf1_params-c.struct
type Cssm_pkcs5_pbkdf1_params struct {
	Passphrase SecAsn1Item
	InitVector SecAsn1Item
}

// Cssm_pkcs5_pbkdf2_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_pkcs5_pbkdf2_params-c.struct
type Cssm_pkcs5_pbkdf2_params struct {
	Passphrase           SecAsn1Item
	PseudoRandomFunction CSSM_PKCS5_PBKDF2_PRF
}

// Cssm_query
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_query-c.struct
type Cssm_query struct {
	RecordType             CSSM_DB_RECORDTYPE
	Conjunctive            CSSM_DB_CONJUNCTIVE
	NumSelectionPredicates uint32
	SelectionPredicate     unsafe.Pointer
	QueryLimits            [2]uint32
	QueryFlags             CSSM_QUERY_FLAGS
}

// Cssm_query_limits
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_query_limits-c.struct
type Cssm_query_limits struct {
	TimeLimit uint32
	SizeLimit uint32
}

// Cssm_query_size_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_query_size_data-swift.struct
type Cssm_query_size_data struct {
	SizeInputBlock  uint32
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
	AccessCred      unsafe.Pointer
	InitialAclEntry [21]uint64
}

// Cssm_sample
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_sample-c.struct
type Cssm_sample struct {
	TypedSample [3]uint64
	Verifier    unsafe.Pointer
}

// Cssm_samplegroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_samplegroup-c.struct
type Cssm_samplegroup struct {
	NumberOfSamples uint32
	Samples         unsafe.Pointer
}

// Cssm_selection_predicate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_selection_predicate-c.struct
type Cssm_selection_predicate struct {
	DbOperator CSSM_DB_OPERATOR
	Attribute  [6]uint64
}

// Cssm_spi_ac_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_ac_funcs-c.struct
type Cssm_spi_ac_funcs struct {
	AuthCompute func(int, uintptr, uintptr, uint32, uintptr, uintptr, uintptr, uintptr) int32
	PassThrough func(int, int, int, uint64, uintptr, uint32, unsafe.Pointer, unsafe.Pointer) int32
}

// Cssm_spi_cl_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_cl_funcs-c.struct
type Cssm_spi_cl_funcs struct {
	CertCreateTemplate           func(int, uint32, uintptr, uintptr) int32
	CertGetAllTemplateFields     func(int, uintptr, *uint32, uintptr) int32
	CertSign                     func(int, uint64, uintptr, uintptr, uint32, uintptr) int32
	CertVerify                   func(int, uint64, uintptr, uintptr, uintptr, uint32) int32
	CertVerifyWithKey            func(int, uint64, uintptr) int32
	CertGetFirstFieldValue       func(int, uintptr, uintptr, *int, *uint32, uintptr) int32
	CertGetNextFieldValue        func(int, int, uintptr) int32
	CertAbortQuery               func(int, int) int32
	CertGetKeyInfo               func(int, uintptr, uintptr) int32
	CertGetAllFields             func(int, uintptr, *uint32, uintptr) int32
	FreeFields                   func(int, uint32, uintptr) int32
	FreeFieldValue               func(int, uintptr, uintptr) int32
	CertCache                    func(int, uintptr, *int) int32
	CertGetFirstCachedFieldValue func(int, int, uintptr, *int, *uint32, uintptr) int32
	CertGetNextCachedFieldValue  func(int, int, uintptr) int32
	CertAbortCache               func(int, int) int32
	CertGroupToSignedBundle      func(int, uint64, uintptr, uintptr, uintptr) int32
	CertGroupFromVerifiedBundle  func(int, uint64, uintptr, uintptr, uintptr) int32
	CertDescribeFormat           func(int, *uint32, uintptr) int32
	CrlCreateTemplate            func(int, uint32, uintptr, uintptr) int32
	CrlSetFields                 func(int, uint32, uintptr, uintptr, uintptr) int32
	CrlAddCert                   func(int, uint64, uintptr, uint32, uintptr, uintptr, uintptr) int32
	CrlRemoveCert                func(int, uintptr, uintptr, uintptr) int32
	CrlSign                      func(int, uint64, uintptr, uintptr, uint32, uintptr) int32
	CrlVerify                    func(int, uint64, uintptr, uintptr, uintptr, uint32) int32
	CrlVerifyWithKey             func(int, uint64, uintptr) int32
	IsCertInCrl                  func(int, uintptr, uintptr, *int32) int32
	CrlGetFirstFieldValue        func(int, uintptr, uintptr, *int, *uint32, uintptr) int32
	CrlGetNextFieldValue         func(int, int, uintptr) int32
	CrlAbortQuery                func(int, int) int32
	CrlGetAllFields              func(int, uintptr, *uint32, uintptr) int32
	CrlCache                     func(int, uintptr, *int) int32
	IsCertInCachedCrl            func(int, uintptr, int, *int32, uintptr) int32
	CrlGetFirstCachedFieldValue  func(int, int, uintptr, uintptr, *int, *uint32, uintptr) int32
	CrlGetNextCachedFieldValue   func(int, int, uintptr) int32
	CrlGetAllCachedRecordFields  func(int, int, uintptr, *uint32, uintptr) int32
	CrlAbortCache                func(int, int) int32
	CrlDescribeFormat            func(int, *uint32, uintptr) int32
	PassThrough                  func(int, uint64, uint32, unsafe.Pointer, unsafe.Pointer) int32
}

// Cssm_spi_csp_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_csp_funcs-c.struct
type Cssm_spi_csp_funcs struct {
	ChangeKeyAcl                  func(int, uintptr, uintptr, uintptr) int32
	ChangeKeyOwner                func(int, uintptr, uintptr, uintptr) int32
	ChangeLoginAcl                func(int, uintptr, uintptr) int32
	ChangeLoginOwner              func(int, uintptr, uintptr) int32
	DecryptData                   func(int, uint64, uintptr, uintptr, uint32, uintptr, uint32, *uint, uintptr, uint64) int32
	DecryptDataFinal              func(int, uint64, uintptr) int32
	DecryptDataInit               func(int, uint64, uintptr, uint64) int32
	DecryptDataUpdate             func(int, uint64, uintptr, uint32, uintptr, uint32, *uint) int32
	DeriveKey                     func(int, uint64, uintptr, uintptr, uint32, uint32, uintptr, uintptr, uintptr) int32
	DigestData                    func(int, uint64, uintptr, uintptr, uint32, uintptr) int32
	DigestDataClone               func(int, uint64, uint64) int32
	DigestDataFinal               func(int, uint64, uintptr) int32
	DigestDataInit                func(int, uint64, uintptr) int32
	DigestDataUpdate              func(int, uint64, uintptr, uint32) int32
	EncryptData                   func(int, uint64, uintptr, uintptr, uint32, uintptr, uint32, *uint, uintptr, uint64) int32
	EncryptDataFinal              func(int, uint64, uintptr) int32
	EncryptDataInit               func(int, uint64, uintptr, uint64) int32
	EncryptDataUpdate             func(int, uint64, uintptr, uint32, uintptr, uint32, *uint) int32
	EventNotify                   func(int, uint32, uint64, uintptr) int32
	FreeKey                       func(int, uintptr, uintptr, int32) int32
	GenerateAlgorithmParams       func(int, uint64, uintptr, uint32, uintptr, *uint32, uintptr) int32
	GenerateKey                   func(int, uint64, uintptr, uint32, uint32, uintptr, uintptr, uintptr, uint64) int32
	GenerateKeyPair               func(int, uint64, uintptr, uint32, uint32, uintptr, uintptr, uint32, uint32, uintptr, uintptr, uintptr, uint64) int32
	GenerateMac                   func(int, uint64, uintptr, uintptr, uint32, uintptr) int32
	GenerateMacFinal              func(int, uint64, uintptr) int32
	GenerateMacInit               func(int, uint64, uintptr) int32
	GenerateMacUpdate             func(int, uint64, uintptr, uint32) int32
	GenerateRandom                func(int, uint64, uintptr, uintptr) int32
	GetKeyOwner                   func(int, uintptr, uintptr) int32
	GetLoginOwner                 func(int, uintptr) int32
	GetOperationalStatistics      func(int, uintptr) int32
	GetTimeValue                  func(int, uint32, uintptr) int32
	Login                         func(int, uintptr, uintptr, unsafe.Pointer) int32
	Logout                        func(int) int32
	ObtainPrivateKeyFromPublicKey func(int, uintptr, uintptr) int32
	PassThrough                   func(int, uint64, uintptr, uint32, unsafe.Pointer, unsafe.Pointer) int32
	QueryKeySizeInBits            func(int, uint64, uintptr, uintptr, uintptr) int32
	QuerySize                     func(int, uint64, uintptr, int32, uint32, uintptr) int32
	RetrieveCounter               func(int, uintptr) int32
	RetrieveUniqueId              func(int, uintptr) int32
	SignData                      func(int, uint64, uintptr, uintptr, uint32, uint32, uintptr) int32
	SignDataFinal                 func(int, uint64, uintptr) int32
	SignDataInit                  func(int, uint64, uintptr) int32
	SignDataUpdate                func(int, uint64, uintptr, uint32) int32
	UnwrapKey                     func(int, uint64, uintptr, uintptr, uintptr, uint32, uint32, uintptr, uintptr, uintptr, uintptr, uint64) int32
	VerifyData                    func(int, uint64, uintptr, uintptr, uint32, uint32, uintptr) int32
	VerifyDataFinal               func(int, uint64, uintptr) int32
	VerifyDataInit                func(int, uint64, uintptr) int32
	VerifyDataUpdate              func(int, uint64, uintptr, uint32) int32
	VerifyDevice                  func(int, uintptr) int32
	VerifyMac                     func(int, uint64, uintptr, uintptr, uint32, uintptr) int32
	VerifyMacFinal                func(int, uint64, uintptr) int32
	VerifyMacInit                 func(int, uint64, uintptr) int32
	VerifyMacUpdate               func(int, uint64, uintptr, uint32) int32
	WrapKey                       func(int, uint64, uintptr, uintptr, uintptr, uintptr, uintptr, uint64) int32
}

// Cssm_spi_dl_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_dl_funcs-c.struct
type Cssm_spi_dl_funcs struct {
	Authenticate              func(Cssm_dl_db_handle, uint32, uintptr) int32
	ChangeDbAcl               func(Cssm_dl_db_handle, uintptr, uintptr) int32
	ChangeDbOwner             func(Cssm_dl_db_handle, uintptr, uintptr) int32
	CreateRelation            func(Cssm_dl_db_handle, uint32, AuthorizationString, uint32, uintptr, uint32, uintptr) int32
	DataAbortQuery            func(Cssm_dl_db_handle, int) int32
	DataDelete                func(Cssm_dl_db_handle, uintptr) int32
	DataGetFirst              func(Cssm_dl_db_handle, uintptr, *int, uintptr, uintptr, uintptr) int32
	DataGetFromUniqueRecordId func(Cssm_dl_db_handle, uintptr, uintptr, uintptr) int32
	DataGetNext               func(Cssm_dl_db_handle, int, uintptr, uintptr, uintptr) int32
	DataInsert                func(Cssm_dl_db_handle, uint32, uintptr, uintptr, uintptr) int32
	DataModify                func(Cssm_dl_db_handle, uint32, uintptr, uintptr, uintptr, uint32) int32
	DbClose                   func(Cssm_dl_db_handle) int32
	DbCreate                  func(int, AuthorizationString, uintptr, uintptr, uint32, uintptr, unsafe.Pointer, *int) int32
	DbDelete                  func(int, AuthorizationString, uintptr, uintptr) int32
	DbOpen                    func(int, AuthorizationString, uintptr, uint32, uintptr, unsafe.Pointer, *int) int32
	DestroyRelation           func(Cssm_dl_db_handle, uint32) int32
	FreeNameList              func(int, uintptr) int32
	FreeUniqueRecord          func(Cssm_dl_db_handle, uintptr) int32
	GetDbNameFromHandle       func(Cssm_dl_db_handle, *byte) int32
	GetDbNames                func(int, uintptr) int32
	GetDbOwner                func(Cssm_dl_db_handle, uintptr) int32
	PassThrough               func(Cssm_dl_db_handle, uint32, unsafe.Pointer, unsafe.Pointer) int32
}

// Cssm_spi_kr_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_kr_funcs-c.struct
type Cssm_spi_kr_funcs struct {
	RegistrationRequest    func(uint32, uint64, uintptr, uintptr, uintptr, uint32, *int32, *int) int32
	RegistrationRetrieve   func(uint32, int, *int32, uintptr) int32
	GenerateRecoveryFields func(uint32, uint64, uintptr, uint64, uintptr, uintptr, uint32, uintptr) int32
	ProcessRecoveryFields  func(uint32, uint64, uintptr, uint64, uintptr, uintptr, uint32, uintptr) int32
	RecoveryRequest        func(uint32, uint64, uintptr, uintptr, uintptr, *int32, *int) int32
	RecoveryRetrieve       func(uint32, int, *int32, *int, *uint32) int32
	GetRecoveredObject     func(uint32, int, uint32, int, uintptr, uint32, uintptr, uintptr) int32
	RecoveryRequestAbort   func(uint32, int) int32
	PassThrough            func(uint32, uint64, uintptr, uint64, uintptr, uint32, unsafe.Pointer, unsafe.Pointer) int32
}

// Cssm_spi_tp_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_tp_funcs-c.struct
type Cssm_spi_tp_funcs struct {
	SubmitCredRequest         func(int, uintptr, uint32, uintptr, uintptr, *int32, uintptr) int32
	RetrieveCredResult        func(int, uintptr, uintptr, *int32, *int32, uintptr) int32
	ConfirmCredResult         func(int, uintptr, uintptr, uintptr, uintptr) int32
	ReceiveConfirmation       func(int, uintptr, uintptr, *int32) int32
	CertReclaimKey            func(int, uintptr, uint32, uint64, int, uintptr) int32
	CertReclaimAbort          func(int, uint64) int32
	FormRequest               func(int, uintptr, uint32, uintptr) int32
	FormSubmit                func(int, uint32, uintptr, uintptr, uintptr, uintptr) int32
	CertGroupVerify           func(int, int, int, uintptr, uintptr, uintptr) int32
	CertCreateTemplate        func(int, int, uint32, uintptr, uintptr) int32
	CertGetAllTemplateFields  func(int, int, uintptr, *uint32, uintptr) int32
	CertSign                  func(int, int, uint64, uintptr, uintptr, uintptr, uintptr, uintptr) int32
	CrlVerify                 func(int, int, int, uintptr, uintptr, uintptr, uintptr) int32
	CrlCreateTemplate         func(int, int, uint32, uintptr, uintptr) int32
	CertRevoke                func(int, int, int, uintptr, uintptr, uintptr, uintptr, uintptr, uint32, uintptr) int32
	CertRemoveFromCrlTemplate func(int, int, int, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) int32
	CrlSign                   func(int, int, uint64, uintptr, uintptr, uintptr, uintptr, uintptr) int32
	ApplyCrlToDb              func(int, int, int, uintptr, uintptr, uintptr, uintptr) int32
	CertGroupConstruct        func(int, int, int, uintptr, unsafe.Pointer, uintptr, uintptr) int32
	CertGroupPrune            func(int, int, uintptr, uintptr, uintptr) int32
	CertGroupToTupleGroup     func(int, int, uintptr, uintptr) int32
	TupleGroupToCertGroup     func(int, int, uintptr, uintptr) int32
	PassThrough               func(int, int, uint64, uintptr, uint32, unsafe.Pointer, unsafe.Pointer) int32
}

// Cssm_state_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_state_funcs-c.struct
type Cssm_state_funcs struct {
	Cssm_GetAttachFunctions        func(int, uint32, unsafe.Pointer, uintptr, *int32) int32
	Cssm_ReleaseAttachFunctions    func(int) int32
	Cssm_GetAppMemoryFunctions     func(int, uintptr) int32
	Cssm_IsFuncCallValid           func(int, CSSM_PROC_ADDR, CSSM_PROC_ADDR, uint64, *uint64, uint32, *int32) int32
	Cssm_DeregisterManagerServices func(uintptr) int32
	Cssm_DeliverModuleManagerEvent func(uintptr) int32
}

// Cssm_subservice_uid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_subservice_uid-c.struct
type Cssm_subservice_uid struct {
	Guid           [4]uint32
	Version        [2]uint32
	SubserviceId   uint32
	SubserviceType CSSM_SERVICE_TYPE
}

// Cssm_tp_authority_id
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_authority_id-c.struct
type Cssm_tp_authority_id struct {
	AuthorityCert     unsafe.Pointer
	AuthorityLocation unsafe.Pointer
}

// Cssm_tp_callerauth_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_callerauth_context-c.struct
type Cssm_tp_callerauth_context struct {
	Policy                   [3]uint64
	VerifyTime               CSSM_TIMESTRING
	VerificationAbortOn      CSSM_TP_STOP_ON
	CallbackWithVerifiedCert unsafe.Pointer
	NumberOfAnchorCerts      uint32
	AnchorCerts              unsafe.Pointer
	DBList                   unsafe.Pointer
	CallerCredentials        unsafe.Pointer
}

// Cssm_tp_certchange_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certchange_input-c.struct
type Cssm_tp_certchange_input struct {
	Action            CSSM_TP_CERTCHANGE_ACTION
	Reason            CSSM_TP_CERTCHANGE_REASON
	CLHandle          CSSM_CL_HANDLE
	Cert              unsafe.Pointer
	ChangeInfo        unsafe.Pointer
	StartTime         CSSM_TIMESTRING
	CallerCredentials unsafe.Pointer
}

// Cssm_tp_certchange_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certchange_output-c.struct
type Cssm_tp_certchange_output struct {
	ActionStatus CSSM_TP_CERTCHANGE_STATUS
	RevokeInfo   [4]uint64
}

// Cssm_tp_certissue_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certissue_input-c.struct
type Cssm_tp_certissue_input struct {
	CSPSubserviceUid        [8]uint32
	CLHandle                CSSM_CL_HANDLE
	NumberOfTemplateFields  uint32
	SubjectCertFields       unsafe.Pointer
	MoreServiceRequests     CSSM_TP_SERVICES
	NumberOfServiceControls uint32
	ServiceControls         unsafe.Pointer
	UserCredentials         unsafe.Pointer
}

// Cssm_tp_certissue_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certissue_output-c.struct
type Cssm_tp_certissue_output struct {
	IssueStatus              CSSM_TP_CERTISSUE_STATUS
	CertGroup                unsafe.Pointer
	PerformedServiceRequests CSSM_TP_SERVICES
}

// Cssm_tp_certnotarize_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certnotarize_input-c.struct
type Cssm_tp_certnotarize_input struct {
	CLHandle                CSSM_CL_HANDLE
	NumberOfFields          uint32
	MoreFields              unsafe.Pointer
	SignScope               unsafe.Pointer
	ScopeSize               uint32
	MoreServiceRequests     CSSM_TP_SERVICES
	NumberOfServiceControls uint32
	ServiceControls         unsafe.Pointer
	UserCredentials         unsafe.Pointer
}

// Cssm_tp_certnotarize_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certnotarize_output-c.struct
type Cssm_tp_certnotarize_output struct {
	NotarizeStatus           CSSM_TP_CERTNOTARIZE_STATUS
	NotarizedCertGroup       unsafe.Pointer
	PerformedServiceRequests CSSM_TP_SERVICES
}

// Cssm_tp_certreclaim_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certreclaim_input-c.struct
type Cssm_tp_certreclaim_input struct {
	CLHandle                CSSM_CL_HANDLE
	NumberOfSelectionFields uint32
	SelectionFields         unsafe.Pointer
	UserCredentials         unsafe.Pointer
}

// Cssm_tp_certreclaim_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certreclaim_output-c.struct
type Cssm_tp_certreclaim_output struct {
	ReclaimStatus      CSSM_TP_CERTRECLAIM_STATUS
	ReclaimedCertGroup unsafe.Pointer
	KeyCacheHandle     CSSM_LONG_HANDLE
}

// Cssm_tp_certverify_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certverify_input-c.struct
type Cssm_tp_certverify_input struct {
	CLHandle      CSSM_CL_HANDLE
	Cert          unsafe.Pointer
	VerifyContext unsafe.Pointer
}

// Cssm_tp_certverify_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certverify_output-c.struct
type Cssm_tp_certverify_output struct {
	VerifyStatus     CSSM_TP_CERTVERIFY_STATUS
	NumberOfEvidence uint32
	Evidence         unsafe.Pointer
}

// Cssm_tp_confirm_response
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_confirm_response-c.struct
type Cssm_tp_confirm_response struct {
	NumberOfResponses uint32
	Responses         CSSM_TP_CONFIRM_STATUS_PTR
}

// Cssm_tp_crlissue_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_crlissue_input-c.struct
type Cssm_tp_crlissue_input struct {
	CLHandle          CSSM_CL_HANDLE
	CrlIdentifier     uint32
	CrlThisTime       CSSM_TIMESTRING
	PolicyIdentifier  unsafe.Pointer
	CallerCredentials unsafe.Pointer
}

// Cssm_tp_crlissue_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_crlissue_output-c.struct
type Cssm_tp_crlissue_output struct {
	IssueStatus CSSM_TP_CRLISSUE_STATUS
	Crl         unsafe.Pointer
	CrlNextTime CSSM_TIMESTRING
}

// Cssm_tp_policyinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_policyinfo-c.struct
type Cssm_tp_policyinfo struct {
	NumberOfPolicyIds uint32
	PolicyIds         unsafe.Pointer
	PolicyControl     unsafe.Pointer
}

// Cssm_tp_request_set
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_request_set-c.struct
type Cssm_tp_request_set struct {
	NumberOfRequests uint32
	Requests         unsafe.Pointer
}

// Cssm_tp_result_set
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_result_set-swift.struct
type Cssm_tp_result_set struct {
	NumberOfResults uint32
	Results         unsafe.Pointer
}

// Cssm_tp_verify_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_verify_context-c.struct
type Cssm_tp_verify_context struct {
	Action     CSSM_TP_ACTION
	ActionData SecAsn1Item
	Crls       [4]uint64
	Cred       unsafe.Pointer
}

// Cssm_tp_verify_context_result
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_verify_context_result-c.struct
type Cssm_tp_verify_context_result struct {
	NumberOfEvidences uint32
	Evidence          unsafe.Pointer
}

// Cssm_tuplegroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tuplegroup-c.struct
type Cssm_tuplegroup struct {
	NumberOfTuples uint32
	Tuples         unsafe.Pointer
}

// Cssm_upcalls
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_upcalls-c.struct
type Cssm_upcalls struct {
	Malloc_func        unsafe.Pointer
	Free_func          unsafe.Pointer
	Realloc_func       unsafe.Pointer
	Calloc_func        unsafe.Pointer
	CcToHandle_func    func(uint64, *int) int32
	GetModuleInfo_func func(int, uintptr, uintptr, *uint32, *uint32, *uint32, *uint32, uintptr, uintptr, uint32) int32
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
	ExtnId   [2]uint64
	Critical CSSM_BOOL
	Format   CSSM_X509EXT_DATA_FORMAT
	Value    unsafe.Pointer
	BERvalue SecAsn1Item
}

// Cssm_x509_extensionTagAndValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_extensionTagAndValue
type Cssm_x509_extensionTagAndValue struct {
	Type  CSSM_BER_TAG
	Value SecAsn1Item
}

// Cssm_x509_extensions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_extensions-c.struct
type Cssm_x509_extensions struct {
	NumberOfExtensions uint32
	Extensions         unsafe.Pointer
}

// Cssm_x509_name
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_name-c.struct
type Cssm_x509_name struct {
	NumberOfRDNs              uint32
	RelativeDistinguishedName unsafe.Pointer
}

// Cssm_x509_rdn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_rdn-c.struct
type Cssm_x509_rdn struct {
	NumberOfPairs         uint32
	AttributeTypeAndValue unsafe.Pointer
}

// Cssm_x509_revoked_cert_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_revoked_cert_entry-c.struct
type Cssm_x509_revoked_cert_entry struct {
	CertificateSerialNumber SecAsn1Item
	RevocationDate          [3]uint64
	Extensions              [2]uint64
}

// Cssm_x509_revoked_cert_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_revoked_cert_list-c.struct
type Cssm_x509_revoked_cert_list struct {
	NumberOfRevokedCertEntries uint32
	RevokedCertEntry           unsafe.Pointer
}

// Cssm_x509_signature
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_signature-c.struct
type Cssm_x509_signature struct {
	AlgorithmIdentifier SecAsn1AlgId
	Encrypted           SecAsn1Item
}

// Cssm_x509_signed_certificate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_signed_certificate-c.struct
type Cssm_x509_signed_certificate struct {
	Certificate [30]uint64
	Signature   [6]uint64
}

// Cssm_x509_signed_crl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_signed_crl-c.struct
type Cssm_x509_signed_crl struct {
	TbsCertList [17]uint64
	Signature   [6]uint64
}

// Cssm_x509_tbs_certificate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_tbs_certificate-c.struct
type Cssm_x509_tbs_certificate struct {
	Version                 SecAsn1Item
	SerialNumber            SecAsn1Item
	Signature               SecAsn1AlgId
	Issuer                  [2]uint64
	Validity                [6]uint64
	Subject                 [2]uint64
	SubjectPublicKeyInfo    SecAsn1PubKeyInfo
	IssuerUniqueIdentifier  SecAsn1Item
	SubjectUniqueIdentifier SecAsn1Item
	Extensions              [2]uint64
}

// Cssm_x509_tbs_certlist
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_tbs_certlist-c.struct
type Cssm_x509_tbs_certlist struct {
	Version             SecAsn1Item
	Signature           SecAsn1AlgId
	Issuer              [2]uint64
	ThisUpdate          [3]uint64
	NextUpdate          [3]uint64
	RevokedCertificates unsafe.Pointer
	Extensions          [2]uint64
}

// Cssm_x509_time
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_time-c.struct
type Cssm_x509_time struct {
	TimeType CSSM_BER_TAG
	Time     SecAsn1Item
}

// Cssm_x509_type_value_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_type_value_pair-c.struct
type Cssm_x509_type_value_pair struct {
	Type      [2]uint64
	ValueType CSSM_BER_TAG
	Value     SecAsn1Item
}

// Cssm_x509ext_basicConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_basicConstraints-c.struct
type Cssm_x509ext_basicConstraints struct {
	CA                       CSSM_BOOL
	PathLenConstraintPresent CSSM_X509_OPTION
	PathLenConstraint        uint32
}

// Cssm_x509ext_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_pair-c.struct
type Cssm_x509ext_pair struct {
	TagAndValue [3]uint64
	ParsedValue unsafe.Pointer
}

// Cssm_x509ext_policyInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_policyInfo-c.struct
type Cssm_x509ext_policyInfo struct {
	PolicyIdentifier [2]uint64
	PolicyQualifiers [2]uint64
}

// Cssm_x509ext_policyQualifierInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_policyQualifierInfo-c.struct
type Cssm_x509ext_policyQualifierInfo struct {
	PolicyQualifierId [2]uint64
	Value             SecAsn1Item
}

// Cssm_x509ext_policyQualifiers
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_policyQualifiers-c.struct
type Cssm_x509ext_policyQualifiers struct {
	NumberOfPolicyQualifiers uint32
	PolicyQualifier          unsafe.Pointer
}

// Mds_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/mds_funcs-c.struct
type Mds_funcs struct {
	DbOpen                    func(int, AuthorizationString, uintptr, uint32, uintptr, unsafe.Pointer, *int) int32
	DbClose                   func(Cssm_dl_db_handle) int32
	GetDbNames                func(int, uintptr) int32
	GetDbNameFromHandle       func(Cssm_dl_db_handle, *byte) int32
	FreeNameList              func(int, uintptr) int32
	DataInsert                func(Cssm_dl_db_handle, uint32, uintptr, uintptr, uintptr) int32
	DataDelete                func(Cssm_dl_db_handle, uintptr) int32
	DataModify                func(Cssm_dl_db_handle, uint32, uintptr, uintptr, uintptr, uint32) int32
	DataGetFirst              func(Cssm_dl_db_handle, uintptr, *int, uintptr, uintptr, uintptr) int32
	DataGetNext               func(Cssm_dl_db_handle, int, uintptr, uintptr, uintptr) int32
	DataAbortQuery            func(Cssm_dl_db_handle, int) int32
	DataGetFromUniqueRecordId func(Cssm_dl_db_handle, uintptr, uintptr, uintptr) int32
	FreeUniqueRecord          func(Cssm_dl_db_handle, uintptr) int32
	CreateRelation            func(Cssm_dl_db_handle, uint32, AuthorizationString, uint32, uintptr, uint32, uintptr) int32
	DestroyRelation           func(Cssm_dl_db_handle, uint32) int32
}

// X509_validity
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/x509_validity
type X509_validity struct {
	NotBefore [3]uint64
	NotAfter  [3]uint64
}
