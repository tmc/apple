// Code generated from Apple documentation. DO NOT EDIT.

package security

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/Security/gGuidAppleCSP
	gGuidAppleCSP unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/gGuidAppleCSPDL
	gGuidAppleCSPDL unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/gGuidAppleDotMacDL
	gGuidAppleDotMacDL unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/gGuidAppleDotMacTP
	gGuidAppleDotMacTP unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/gGuidAppleFileDL
	gGuidAppleFileDL unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/gGuidAppleLDAPDL
	gGuidAppleLDAPDL unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/gGuidAppleSdCSPDL
	gGuidAppleSdCSPDL unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/gGuidAppleX509CL
	gGuidAppleX509CL unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/gGuidAppleX509TP
	gGuidAppleX509TP unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/gGuidCssm
	gGuidCssm unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidAdCAIssuer
	oidAdCAIssuer unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidAdOCSP
	oidAdOCSP unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidAnsip384r1
	oidAnsip384r1 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidAnsip521r1
	oidAnsip521r1 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidAnyExtendedKeyUsage
	oidAnyExtendedKeyUsage unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidAnyPolicy
	oidAnyPolicy unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidAuthorityInfoAccess
	oidAuthorityInfoAccess unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidAuthorityKeyIdentifier
	oidAuthorityKeyIdentifier unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidBasicConstraints
	oidBasicConstraints unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidCertificatePolicies
	oidCertificatePolicies unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidCommonName
	oidCommonName unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidCountryName
	oidCountryName unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidCrlDistributionPoints
	oidCrlDistributionPoints unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidDescription
	oidDescription unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidEcPrime192v1
	oidEcPrime192v1 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidEcPrime256v1
	oidEcPrime256v1 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidEcPubKey
	oidEcPubKey unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidEmailAddress
	oidEmailAddress unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidEntrustVersInfo
	oidEntrustVersInfo unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidExtendedKeyUsage
	oidExtendedKeyUsage unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidExtendedKeyUsageClientAuth
	oidExtendedKeyUsageClientAuth unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidExtendedKeyUsageCodeSigning
	oidExtendedKeyUsageCodeSigning unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidExtendedKeyUsageEmailProtection
	oidExtendedKeyUsageEmailProtection unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidExtendedKeyUsageIPSec
	oidExtendedKeyUsageIPSec unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidExtendedKeyUsageMicrosoftSGC
	oidExtendedKeyUsageMicrosoftSGC unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidExtendedKeyUsageNetscapeSGC
	oidExtendedKeyUsageNetscapeSGC unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidExtendedKeyUsageOCSPSigning
	oidExtendedKeyUsageOCSPSigning unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidExtendedKeyUsageServerAuth
	oidExtendedKeyUsageServerAuth unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidExtendedKeyUsageTimeStamping
	oidExtendedKeyUsageTimeStamping unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidFee
	oidFee unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidFriendlyName
	oidFriendlyName unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidGoogleEmbeddedSignedCertificateTimestamp
	oidGoogleEmbeddedSignedCertificateTimestamp unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidGoogleOCSPSignedCertificateTimestamp
	oidGoogleOCSPSignedCertificateTimestamp unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidInhibitAnyPolicy
	oidInhibitAnyPolicy unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidIssuerAltName
	oidIssuerAltName unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidKeyUsage
	oidKeyUsage unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidLocalKeyId
	oidLocalKeyId unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidLocalityName
	oidLocalityName unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidMSNTPrincipalName
	oidMSNTPrincipalName unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidMd2
	oidMd2 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidMd2Rsa
	oidMd2Rsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidMd4
	oidMd4 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidMd4Rsa
	oidMd4Rsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidMd5
	oidMd5 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidMd5Fee
	oidMd5Fee unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidMd5Rsa
	oidMd5Rsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidNameConstraints
	oidNameConstraints unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidNetscapeCertType
	oidNetscapeCertType unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidOrganizationName
	oidOrganizationName unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidOrganizationalUnitName
	oidOrganizationalUnitName unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidPolicyConstraints
	oidPolicyConstraints unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidPolicyMappings
	oidPolicyMappings unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidPrivateKeyUsagePeriod
	oidPrivateKeyUsagePeriod unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidQtCps
	oidQtCps unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidQtUNotice
	oidQtUNotice unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidRsa
	oidRsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha1
	oidSha1 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha1Dsa
	oidSha1Dsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha1DsaCommonOIW
	oidSha1DsaCommonOIW unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha1DsaOIW
	oidSha1DsaOIW unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha1Ecdsa
	oidSha1Ecdsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha1Fee
	oidSha1Fee unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha1Rsa
	oidSha1Rsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha1RsaOIW
	oidSha1RsaOIW unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha224
	oidSha224 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha224Ecdsa
	oidSha224Ecdsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha224Rsa
	oidSha224Rsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha256
	oidSha256 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha256Ecdsa
	oidSha256Ecdsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha256Rsa
	oidSha256Rsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha384
	oidSha384 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha384Ecdsa
	oidSha384Ecdsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha384Rsa
	oidSha384Rsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha512
	oidSha512 unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha512Ecdsa
	oidSha512Ecdsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSha512Rsa
	oidSha512Rsa unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidStateOrProvinceName
	oidStateOrProvinceName unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSubjectAltName
	oidSubjectAltName unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSubjectInfoAccess
	oidSubjectInfoAccess unsafe.Pointer
	// See: https://developer.apple.com/documentation/Security/oidSubjectKeyIdentifier
	oidSubjectKeyIdentifier unsafe.Pointer
)

var (
	// See: https://developer.apple.com/documentation/Security/kCMSEncoderDigestAlgorithmSHA1
	KCMSEncoderDigestAlgorithmSHA1 string
	// See: https://developer.apple.com/documentation/Security/kCMSEncoderDigestAlgorithmSHA256
	KCMSEncoderDigestAlgorithmSHA256 string
	// KSecACLAuthorizationAny is no restrictions. This ACL entry applies to all operations available to the caller.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationAny
	KSecACLAuthorizationAny string
	// KSecACLAuthorizationChangeACL is change an access control list entry.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationChangeACL
	KSecACLAuthorizationChangeACL string
	// KSecACLAuthorizationChangeOwner is for internal system use only. Use the `CSSM_ACL_AUTHORIZATION_CHANGE_ACL` tag for changes to owner ACL entries.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationChangeOwner
	KSecACLAuthorizationChangeOwner string
	// KSecACLAuthorizationDecrypt is decrypt data.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationDecrypt
	KSecACLAuthorizationDecrypt string
	// KSecACLAuthorizationDelete is delete this item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationDelete
	KSecACLAuthorizationDelete string
	// KSecACLAuthorizationDerive is derive a new key from another key.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationDerive
	KSecACLAuthorizationDerive string
	// KSecACLAuthorizationEncrypt is encrypt data.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationEncrypt
	KSecACLAuthorizationEncrypt string
	// KSecACLAuthorizationExportClear is export an unencrypted key.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationExportClear
	KSecACLAuthorizationExportClear string
	// KSecACLAuthorizationExportWrapped is export a wrapped (that is, encrypted) key. This tag is checked on the key being exported; in addition, the `CSSM_ACL_AUTHORIZATION_ENCRYPT` tag is checked for any key used in the wrapping operation.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationExportWrapped
	KSecACLAuthorizationExportWrapped string
	// KSecACLAuthorizationGenKey is generate a key.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationGenKey
	KSecACLAuthorizationGenKey string
	// KSecACLAuthorizationImportClear is import an unencrypted key.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationImportClear
	KSecACLAuthorizationImportClear string
	// KSecACLAuthorizationImportWrapped is import an encrypted key. This tag is checked on the key being imported; in addition, the `CSSM_ACL_AUTHORIZATION_DECRYPT` tag is checked for any key used in the unwrapping operation.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationImportWrapped
	KSecACLAuthorizationImportWrapped string
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationIntegrity
	KSecACLAuthorizationIntegrity string
	// KSecACLAuthorizationKeychainCreate is create a new keychain.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationKeychainCreate
	KSecACLAuthorizationKeychainCreate string
	// KSecACLAuthorizationKeychainDelete is delete a keychain.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationKeychainDelete
	KSecACLAuthorizationKeychainDelete string
	// KSecACLAuthorizationKeychainItemDelete is delete an item from a keychain.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationKeychainItemDelete
	KSecACLAuthorizationKeychainItemDelete string
	// KSecACLAuthorizationKeychainItemInsert is insert an item into a keychain.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationKeychainItemInsert
	KSecACLAuthorizationKeychainItemInsert string
	// KSecACLAuthorizationKeychainItemModify is modify an item in a keychain.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationKeychainItemModify
	KSecACLAuthorizationKeychainItemModify string
	// KSecACLAuthorizationKeychainItemRead is read an item from a keychain.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationKeychainItemRead
	KSecACLAuthorizationKeychainItemRead string
	// KSecACLAuthorizationLogin is use for a CSP (smart card) login.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationLogin
	KSecACLAuthorizationLogin string
	// KSecACLAuthorizationMAC is create or verify a message authentication code.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationMAC
	KSecACLAuthorizationMAC string
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationPartitionID
	KSecACLAuthorizationPartitionID string
	// KSecACLAuthorizationSign is digitally sign data.
	//
	// See: https://developer.apple.com/documentation/Security/kSecACLAuthorizationSign
	KSecACLAuthorizationSign string
	// KSecAttrAccess is a key with a value that indicates access control list settings for the item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAccess
	KSecAttrAccess string
	// KSecAttrAccessControl is a key with a value that’s an access control instance indicating access control settings for the item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAccessControl
	KSecAttrAccessControl string
	// KSecAttrAccessGroup is a key with a value that’s a string indicating the access group the item is in.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAccessGroup
	KSecAttrAccessGroup string
	// KSecAttrAccessGroupToken is the access group containing items provided by external tokens.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAccessGroupToken
	KSecAttrAccessGroupToken string
	// KSecAttrAccessible is a key with a value that indicates when the keychain item is accessible.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAccessible
	KSecAttrAccessible string
	// KSecAttrAccessibleAfterFirstUnlock is the data in the keychain item cannot be accessed after a restart until the device has been unlocked once by the user.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAccessibleAfterFirstUnlock
	KSecAttrAccessibleAfterFirstUnlock string
	// KSecAttrAccessibleAfterFirstUnlockThisDeviceOnly is the data in the keychain item cannot be accessed after a restart until the device has been unlocked once by the user.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAccessibleAfterFirstUnlockThisDeviceOnly
	KSecAttrAccessibleAfterFirstUnlockThisDeviceOnly string
	// KSecAttrAccessibleWhenPasscodeSetThisDeviceOnly is the data in the keychain can only be accessed when the device is unlocked. Only available if a passcode is set on the device.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAccessibleWhenPasscodeSetThisDeviceOnly
	KSecAttrAccessibleWhenPasscodeSetThisDeviceOnly string
	// KSecAttrAccessibleWhenUnlocked is the data in the keychain item can be accessed only while the device is unlocked by the user.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAccessibleWhenUnlocked
	KSecAttrAccessibleWhenUnlocked string
	// KSecAttrAccessibleWhenUnlockedThisDeviceOnly is the data in the keychain item can be accessed only while the device is unlocked by the user.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAccessibleWhenUnlockedThisDeviceOnly
	KSecAttrAccessibleWhenUnlockedThisDeviceOnly string
	// KSecAttrAccount is a key whose value is a string indicating the item’s account name.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAccount
	KSecAttrAccount string
	// KSecAttrApplicationLabel is a key whose value indicates the item’s application label.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrApplicationLabel
	KSecAttrApplicationLabel string
	// KSecAttrApplicationTag is a key whose value indicates the item’s private tag.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrApplicationTag
	KSecAttrApplicationTag string
	// KSecAttrAuthenticationType is a key whose value indicates the item’s authentication scheme.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAuthenticationType
	KSecAttrAuthenticationType string
	// KSecAttrAuthenticationTypeDPA is distributed Password authentication.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAuthenticationTypeDPA
	KSecAttrAuthenticationTypeDPA string
	// KSecAttrAuthenticationTypeDefault is the default authentication type.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAuthenticationTypeDefault
	KSecAttrAuthenticationTypeDefault string
	// KSecAttrAuthenticationTypeHTMLForm is hTML form based authentication.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAuthenticationTypeHTMLForm
	KSecAttrAuthenticationTypeHTMLForm string
	// KSecAttrAuthenticationTypeHTTPBasic is hTTP Basic authentication.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAuthenticationTypeHTTPBasic
	KSecAttrAuthenticationTypeHTTPBasic string
	// KSecAttrAuthenticationTypeHTTPDigest is hTTP Digest Access authentication.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAuthenticationTypeHTTPDigest
	KSecAttrAuthenticationTypeHTTPDigest string
	// KSecAttrAuthenticationTypeMSN is microsoft Network default authentication.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAuthenticationTypeMSN
	KSecAttrAuthenticationTypeMSN string
	// KSecAttrAuthenticationTypeNTLM is windows NT LAN Manager authentication.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAuthenticationTypeNTLM
	KSecAttrAuthenticationTypeNTLM string
	// KSecAttrAuthenticationTypeRPA is remote Password authentication.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrAuthenticationTypeRPA
	KSecAttrAuthenticationTypeRPA string
	// KSecAttrCanDecrypt is a key whose value is a Boolean that indicates whether the cryptographic key can be used for decryption.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrCanDecrypt
	KSecAttrCanDecrypt string
	// KSecAttrCanDerive is a key whose value is a Boolean that indicates whether the cryptographic key can be used for derivation.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrCanDerive
	KSecAttrCanDerive string
	// KSecAttrCanEncrypt is a key whose value is a Boolean that indicates whether the cryptographic key can be used for encryption.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrCanEncrypt
	KSecAttrCanEncrypt string
	// KSecAttrCanSign is a key whose value is a Boolean that indicates whether the cryptographic key can be used for digital signing.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrCanSign
	KSecAttrCanSign string
	// KSecAttrCanUnwrap is a key whose value is a Boolean that indicates whether the cryptographic key can be used for unwrapping.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrCanUnwrap
	KSecAttrCanUnwrap string
	// KSecAttrCanVerify is a key whose value is a Boolean that indicates whether the cryptographic key can be used for signature verification.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrCanVerify
	KSecAttrCanVerify string
	// KSecAttrCanWrap is a key whose value is a Boolean that indicates whether the cryptographic key can be used for wrapping.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrCanWrap
	KSecAttrCanWrap string
	// KSecAttrCertificateEncoding is a key whose value indicates the item’s certificate encoding.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrCertificateEncoding
	KSecAttrCertificateEncoding string
	// KSecAttrCertificateType is a key whose value indicates the item’s certificate type.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrCertificateType
	KSecAttrCertificateType string
	// KSecAttrComment is a key with a value that’s a string indicating a comment associated with the item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrComment
	KSecAttrComment string
	// KSecAttrCreationDate is a key with a value that indicates the item’s creation date.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrCreationDate
	KSecAttrCreationDate string
	// KSecAttrCreator is a key with a value that indicates the item’s creator.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrCreator
	KSecAttrCreator string
	// KSecAttrDescription is a key with a value that’s a string indicating the item’s description.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrDescription
	KSecAttrDescription string
	// KSecAttrEffectiveKeySize is a key whose value indicates the effective number of bits in a cryptographic key.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrEffectiveKeySize
	KSecAttrEffectiveKeySize string
	// KSecAttrGeneric is a key whose value indicates the item’s user-defined attributes.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrGeneric
	KSecAttrGeneric string
	// KSecAttrIsExtractable is a key whose value indicates the item’s extractability.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrIsExtractable
	KSecAttrIsExtractable string
	// KSecAttrIsInvisible is a key with a value that’s a Boolean indicating the item’s visibility.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrIsInvisible
	KSecAttrIsInvisible string
	// KSecAttrIsNegative is a key with a value that’s a Boolean indicating whether the item has a valid password.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrIsNegative
	KSecAttrIsNegative string
	// KSecAttrIsPermanent is a key whose value indicates the item’s permanence.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrIsPermanent
	KSecAttrIsPermanent string
	// KSecAttrIsSensitive is a key whose value indicates the item’s sensitivity.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrIsSensitive
	KSecAttrIsSensitive string
	// KSecAttrIssuer is a key whose value indicates the item’s issuer.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrIssuer
	KSecAttrIssuer string
	// KSecAttrKeyClass is a key whose value indicates the item’s cryptographic key class.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyClass
	KSecAttrKeyClass string
	// KSecAttrKeyClassPrivate is a private key of a public-private pair.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyClassPrivate
	KSecAttrKeyClassPrivate string
	// KSecAttrKeyClassPublic is a public key of a public-private pair.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyClassPublic
	KSecAttrKeyClassPublic string
	// KSecAttrKeyClassSymmetric is a private key used for symmetric-key encryption and decryption.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyClassSymmetric
	KSecAttrKeyClassSymmetric string
	// KSecAttrKeySizeInBits is a key whose value indicates the number of bits in a cryptographic key.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeySizeInBits
	KSecAttrKeySizeInBits string
	// KSecAttrKeyType is a key whose value indicates the item’s algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyType
	KSecAttrKeyType string
	// KSecAttrKeyType3DES is 3DES algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyType3DES
	KSecAttrKeyType3DES string
	// KSecAttrKeyTypeAES is aES algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyTypeAES
	KSecAttrKeyTypeAES string
	// KSecAttrKeyTypeCAST is cAST algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyTypeCAST
	KSecAttrKeyTypeCAST string
	// KSecAttrKeyTypeDES is dES algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyTypeDES
	KSecAttrKeyTypeDES string
	// KSecAttrKeyTypeDSA is dSA algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyTypeDSA
	KSecAttrKeyTypeDSA string
	// KSecAttrKeyTypeECSECPrimeRandom is elliptic curve algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyTypeECSECPrimeRandom
	KSecAttrKeyTypeECSECPrimeRandom string
	// KSecAttrKeyTypeRC2 is rC2 algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyTypeRC2
	KSecAttrKeyTypeRC2 string
	// KSecAttrKeyTypeRC4 is rC4 algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyTypeRC4
	KSecAttrKeyTypeRC4 string
	// KSecAttrKeyTypeRSA is rSA algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrKeyTypeRSA
	KSecAttrKeyTypeRSA string
	// KSecAttrLabel is a key with a value that’s a string indicating the item’s label.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrLabel
	KSecAttrLabel string
	// KSecAttrModificationDate is a key with a value that indicates the item’s most recent modification date.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrModificationDate
	KSecAttrModificationDate string
	// KSecAttrPRF is a key whose value indicates the item’s pseudorandom function.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrPRF
	KSecAttrPRF string
	// KSecAttrPRFHmacAlgSHA1 is use the SHA1 algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrPRFHmacAlgSHA1
	KSecAttrPRFHmacAlgSHA1 string
	// KSecAttrPRFHmacAlgSHA224 is use the SHA224 algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrPRFHmacAlgSHA224
	KSecAttrPRFHmacAlgSHA224 string
	// KSecAttrPRFHmacAlgSHA256 is use the SHA256 algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrPRFHmacAlgSHA256
	KSecAttrPRFHmacAlgSHA256 string
	// KSecAttrPRFHmacAlgSHA384 is use the SHA384 algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrPRFHmacAlgSHA384
	KSecAttrPRFHmacAlgSHA384 string
	// KSecAttrPRFHmacAlgSHA512 is use the SHA512 algorithm.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrPRFHmacAlgSHA512
	KSecAttrPRFHmacAlgSHA512 string
	// KSecAttrPath is a key whose value is a string indicating the item’s path attribute.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrPath
	KSecAttrPath string
	// See: https://developer.apple.com/documentation/Security/kSecAttrPersistantReference
	KSecAttrPersistantReference string
	// See: https://developer.apple.com/documentation/Security/kSecAttrPersistentReference
	KSecAttrPersistentReference string
	// KSecAttrPort is a key whose value indicates the item’s port.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrPort
	KSecAttrPort string
	// KSecAttrProtocol is a key whose value indicates the item’s protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocol
	KSecAttrProtocol string
	// KSecAttrProtocolAFP is aFP over TCP.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolAFP
	KSecAttrProtocolAFP string
	// KSecAttrProtocolAppleTalk is aFP over AppleTalk.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolAppleTalk
	KSecAttrProtocolAppleTalk string
	// KSecAttrProtocolDAAP is dAAP protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolDAAP
	KSecAttrProtocolDAAP string
	// KSecAttrProtocolEPPC is remote Apple Events.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolEPPC
	KSecAttrProtocolEPPC string
	// KSecAttrProtocolFTP is fTP protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolFTP
	KSecAttrProtocolFTP string
	// KSecAttrProtocolFTPAccount is a client side FTP account.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolFTPAccount
	KSecAttrProtocolFTPAccount string
	// KSecAttrProtocolFTPProxy is fTP proxy.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolFTPProxy
	KSecAttrProtocolFTPProxy string
	// KSecAttrProtocolFTPS is fTP over TLS/SSL.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolFTPS
	KSecAttrProtocolFTPS string
	// KSecAttrProtocolHTTP is hTTP protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolHTTP
	KSecAttrProtocolHTTP string
	// KSecAttrProtocolHTTPProxy is hTTP proxy.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolHTTPProxy
	KSecAttrProtocolHTTPProxy string
	// KSecAttrProtocolHTTPS is hTTP over TLS/SSL.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolHTTPS
	KSecAttrProtocolHTTPS string
	// KSecAttrProtocolHTTPSProxy is hTTPS proxy.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolHTTPSProxy
	KSecAttrProtocolHTTPSProxy string
	// KSecAttrProtocolIMAP is iMAP protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolIMAP
	KSecAttrProtocolIMAP string
	// KSecAttrProtocolIMAPS is iMAP over TLS/SSL.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolIMAPS
	KSecAttrProtocolIMAPS string
	// KSecAttrProtocolIPP is iPP protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolIPP
	KSecAttrProtocolIPP string
	// KSecAttrProtocolIRC is iRC protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolIRC
	KSecAttrProtocolIRC string
	// KSecAttrProtocolIRCS is iRC over TLS/SSL.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolIRCS
	KSecAttrProtocolIRCS string
	// KSecAttrProtocolLDAP is lDAP protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolLDAP
	KSecAttrProtocolLDAP string
	// KSecAttrProtocolLDAPS is lDAP over TLS/SSL.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolLDAPS
	KSecAttrProtocolLDAPS string
	// KSecAttrProtocolNNTP is nNTP protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolNNTP
	KSecAttrProtocolNNTP string
	// KSecAttrProtocolNNTPS is nNTP over TLS/SSL.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolNNTPS
	KSecAttrProtocolNNTPS string
	// KSecAttrProtocolPOP3 is pOP3 protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolPOP3
	KSecAttrProtocolPOP3 string
	// KSecAttrProtocolPOP3S is pOP3 over TLS/SSL.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolPOP3S
	KSecAttrProtocolPOP3S string
	// KSecAttrProtocolRTSP is rTSP protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolRTSP
	KSecAttrProtocolRTSP string
	// KSecAttrProtocolRTSPProxy is rTSP proxy.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolRTSPProxy
	KSecAttrProtocolRTSPProxy string
	// KSecAttrProtocolSMB is sMB protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolSMB
	KSecAttrProtocolSMB string
	// KSecAttrProtocolSMTP is sMTP protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolSMTP
	KSecAttrProtocolSMTP string
	// KSecAttrProtocolSOCKS is sOCKS protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolSOCKS
	KSecAttrProtocolSOCKS string
	// KSecAttrProtocolSSH is sSH protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolSSH
	KSecAttrProtocolSSH string
	// KSecAttrProtocolTelnet is telnet protocol.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolTelnet
	KSecAttrProtocolTelnet string
	// KSecAttrProtocolTelnetS is telnet over TLS/SSL.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrProtocolTelnetS
	KSecAttrProtocolTelnetS string
	// KSecAttrPublicKeyHash is a key whose value indicates the item’s public key hash.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrPublicKeyHash
	KSecAttrPublicKeyHash string
	// KSecAttrRounds is a key whose value indicates the number of rounds to run the pseudorandom function.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrRounds
	KSecAttrRounds string
	// KSecAttrSalt is a key whose value indicates the salt to use for this item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrSalt
	KSecAttrSalt string
	// KSecAttrSecurityDomain is a key whose value is a string indicating the item’s security domain.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrSecurityDomain
	KSecAttrSecurityDomain string
	// KSecAttrSerialNumber is a key whose value indicates the item’s serial number.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrSerialNumber
	KSecAttrSerialNumber string
	// KSecAttrServer is a key whose value is a string indicating the item’s server.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrServer
	KSecAttrServer string
	// KSecAttrService is a key whose value is a string indicating the item’s service.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrService
	KSecAttrService string
	// KSecAttrSubject is a key whose value indicates the item’s subject name.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrSubject
	KSecAttrSubject string
	// KSecAttrSubjectKeyID is a key whose value indicates the item’s subject key ID.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrSubjectKeyID
	KSecAttrSubjectKeyID string
	// KSecAttrSyncViewHint is a key with a value that’s a string that provides a sync view hint.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrSyncViewHint
	KSecAttrSyncViewHint string
	// KSecAttrSynchronizable is a key with a value that’s a string indicating whether the item synchronizes through iCloud.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrSynchronizable
	KSecAttrSynchronizable string
	// KSecAttrSynchronizableAny is specifies that both synchronizable and non-synchronizable results should be returned from a query.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrSynchronizableAny
	KSecAttrSynchronizableAny string
	// KSecAttrTokenID is a key whose value indicates that a cryptographic key is in an external store.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrTokenID
	KSecAttrTokenID string
	// KSecAttrTokenIDSecureEnclave is specifies an item should be stored in the device’s Secure Enclave.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrTokenIDSecureEnclave
	KSecAttrTokenIDSecureEnclave string
	// KSecAttrType is a key with a value that indicates the item’s type.
	//
	// See: https://developer.apple.com/documentation/Security/kSecAttrType
	KSecAttrType string
	// KSecCFErrorArchitecture is a key whose value is a string containing the name of the architecture that is causing the problem.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorArchitecture
	KSecCFErrorArchitecture string
	// KSecCFErrorGuestAttributes is a key whose value is a Core Foundation object containing an attribute that is unrecognized or that contains a value of the wrong type.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorGuestAttributes
	KSecCFErrorGuestAttributes string
	// KSecCFErrorInfoPlist is a key whose value is a Core Foundation object identifying the invalid component or key in the dictionary.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorInfoPlist
	KSecCFErrorInfoPlist string
	// KSecCFErrorPath is a key whose value is a URL indicating the subcomponent containing the error.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorPath
	KSecCFErrorPath string
	// KSecCFErrorPattern is a key whose value is a string containing a regular expression that’s part of a resource specification that did not parse correctly.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorPattern
	KSecCFErrorPattern string
	// KSecCFErrorRequirementSyntax is a key whose value is a string containing a compilation error generated when parsing a requirement.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorRequirementSyntax
	KSecCFErrorRequirementSyntax string
	// KSecCFErrorResourceAdded is a key whose value is a URL pointing to the resource on disk that is not included in the signed resources for the code.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorResourceAdded
	KSecCFErrorResourceAdded string
	// KSecCFErrorResourceAltered is a key whose value is a URL pointing to the resource on disk that has been altered.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorResourceAltered
	KSecCFErrorResourceAltered string
	// KSecCFErrorResourceMissing is a key whose value is a URL indicating the location of the missing resource as it is specified in the [CodeResources] file.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorResourceMissing
	KSecCFErrorResourceMissing string
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorResourceRecursive
	KSecCFErrorResourceRecursive string
	// KSecCFErrorResourceSeal is a key whose value is a Core Foundation object containing the part of the resource seal that had a problem.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorResourceSeal
	KSecCFErrorResourceSeal string
	// KSecCFErrorResourceSideband is a key whose value is a URL representing a sealed resource with invalid sideband data (resource fork, etc.).
	//
	// See: https://developer.apple.com/documentation/Security/kSecCFErrorResourceSideband
	KSecCFErrorResourceSideband string
	// KSecClass is a dictionary key whose value is the item’s class.
	//
	// See: https://developer.apple.com/documentation/Security/kSecClass
	KSecClass string
	// KSecClassCertificate is the value that indicates a certificate item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecClassCertificate
	KSecClassCertificate string
	// KSecClassGenericPassword is the value that indicates a generic password item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecClassGenericPassword
	KSecClassGenericPassword string
	// KSecClassIdentity is the value that indicates an identity item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecClassIdentity
	KSecClassIdentity string
	// KSecClassInternetPassword is the value that indicates an Internet password item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecClassInternetPassword
	KSecClassInternetPassword string
	// KSecClassKey is the value that indicates a cryptographic key item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecClassKey
	KSecClassKey string
	// KSecCodeAttributeArchitecture is a key whose value is a string that indicates an architecture, such as `i386` or `x86_64`.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeAttributeArchitecture
	KSecCodeAttributeArchitecture string
	// KSecCodeAttributeBundleVersion is a key whose value indicates the bundle version.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeAttributeBundleVersion
	KSecCodeAttributeBundleVersion string
	// KSecCodeAttributeSubarchitecture is a key whose value is a string indicating a specific processor type, such as `i686` or `core2`.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeAttributeSubarchitecture
	KSecCodeAttributeSubarchitecture string
	// KSecCodeAttributeUniversalFileOffset is a key whose value indicates the offset of a Mach-O specific slice of a universal Mach-O file.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeAttributeUniversalFileOffset
	KSecCodeAttributeUniversalFileOffset string
	// KSecCodeInfoCMS is a key whose value is the CMS cryptographic object that secures the code signature.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoCMS
	KSecCodeInfoCMS string
	// KSecCodeInfoCdHashes is a key whose value is an array containing the unique binary identifier for every digest algorithm supported in the signature.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoCdHashes
	KSecCodeInfoCdHashes string
	// KSecCodeInfoCertificates is a key whose value is an array of certificates representing the certificate chain of the signing certificate as seen by the system.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoCertificates
	KSecCodeInfoCertificates string
	// KSecCodeInfoChangedFiles is a key whose value is a list of all files in the code that may have been modified by the process of signing it.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoChangedFiles
	KSecCodeInfoChangedFiles string
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoDefaultDesignatedLightweightCodeRequirement
	KSecCodeInfoDefaultDesignatedLightweightCodeRequirement string
	// KSecCodeInfoDesignatedRequirement is a keys whose value is the designated requirement of the code.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoDesignatedRequirement
	KSecCodeInfoDesignatedRequirement string
	// KSecCodeInfoDigestAlgorithm is a key whose value is a number indicating the cryptographic hash function.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoDigestAlgorithm
	KSecCodeInfoDigestAlgorithm string
	// KSecCodeInfoDigestAlgorithms is a key whose value is a list of the kinds of cryptographic hash functions available within the signature.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoDigestAlgorithms
	KSecCodeInfoDigestAlgorithms string
	// KSecCodeInfoEntitlements is a key whose value represents the embedded entitlement blob of the code, if any.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoEntitlements
	KSecCodeInfoEntitlements string
	// KSecCodeInfoEntitlementsDict is a key whose value is a dictionary of embedded entitlements.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoEntitlementsDict
	KSecCodeInfoEntitlementsDict string
	// KSecCodeInfoFlags is a key whose value indicates the static (on-disk) state of the object.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoFlags
	KSecCodeInfoFlags string
	// KSecCodeInfoFormat is a key whose value is a string representing the type and format of the code in a form suitable for display to a knowledgeable user.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoFormat
	KSecCodeInfoFormat string
	// KSecCodeInfoIdentifier is a key whose value is the signing identifier sealed into the signature.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoIdentifier
	KSecCodeInfoIdentifier string
	// KSecCodeInfoImplicitDesignatedRequirement is a key whose value is the designated requirement (DR) that the system generated—or would have generated—for the code in the absence of an explicitly-declared DR.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoImplicitDesignatedRequirement
	KSecCodeInfoImplicitDesignatedRequirement string
	// KSecCodeInfoMainExecutable is a key whose value is a URL locating the main executable file of the code.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoMainExecutable
	KSecCodeInfoMainExecutable string
	// KSecCodeInfoPList is a key whose value is an information dictionary containing the contents of the secured `Info.Plist()` file as seen by Code Signing Services.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoPList
	KSecCodeInfoPList string
	// KSecCodeInfoPlatformIdentifier is a key whose value identifies the operating system release with which the code is associated, if any.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoPlatformIdentifier
	KSecCodeInfoPlatformIdentifier string
	// KSecCodeInfoRequirementData is a key whose value is the internal requirements of the code as a binary blob.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoRequirementData
	KSecCodeInfoRequirementData string
	// KSecCodeInfoRequirements is a key whose value is the internal requirements of the code as a text string in canonical syntax.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoRequirements
	KSecCodeInfoRequirements string
	// KSecCodeInfoRuntimeVersion is a key whose value represents the runtime version.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoRuntimeVersion
	KSecCodeInfoRuntimeVersion string
	// KSecCodeInfoSource is the source of the code signature used for the code object in a format suitable for display.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoSource
	KSecCodeInfoSource string
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoStapledNotarizationTicket
	KSecCodeInfoStapledNotarizationTicket string
	// KSecCodeInfoStatus is a key whose value is the set of code status flags for the running code.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoStatus
	KSecCodeInfoStatus string
	// KSecCodeInfoTeamIdentifier is a key whose value is the team identifier.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoTeamIdentifier
	KSecCodeInfoTeamIdentifier string
	// KSecCodeInfoTime is a key whose value is the signing date embedded in the code signature.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoTime
	KSecCodeInfoTime string
	// KSecCodeInfoTimestamp is a key whose value indicates the actual signing date.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoTimestamp
	KSecCodeInfoTimestamp string
	// KSecCodeInfoTrust is a key whose value is the trust object the system uses to evaluate the validity of the code’s signature.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoTrust
	KSecCodeInfoTrust string
	// KSecCodeInfoUnique is a key whose value is a binary number that uniquely identifies static code.
	//
	// See: https://developer.apple.com/documentation/Security/kSecCodeInfoUnique
	KSecCodeInfoUnique string
	// KSecDecodeTypeAttribute is the encoding used by a decode transform.
	//
	// See: https://developer.apple.com/documentation/Security/kSecDecodeTypeAttribute
	KSecDecodeTypeAttribute string
	// KSecGuestAttributeArchitecture is a key whose value is a number representing the CPU type under which the guest code is designed to run.
	//
	// See: https://developer.apple.com/documentation/Security/kSecGuestAttributeArchitecture
	KSecGuestAttributeArchitecture string
	// See: https://developer.apple.com/documentation/Security/kSecGuestAttributeAudit
	KSecGuestAttributeAudit string
	// KSecGuestAttributeCanonical is a key whose value is the guest code object for that guest.
	//
	// See: https://developer.apple.com/documentation/Security/kSecGuestAttributeCanonical
	KSecGuestAttributeCanonical string
	// See: https://developer.apple.com/documentation/Security/kSecGuestAttributeDynamicCode
	KSecGuestAttributeDynamicCode string
	// See: https://developer.apple.com/documentation/Security/kSecGuestAttributeDynamicCodeInfoPlist
	KSecGuestAttributeDynamicCodeInfoPlist string
	// KSecGuestAttributeHash is a key whose value is a data object containing the SHA-1 hash of the code directory.
	//
	// See: https://developer.apple.com/documentation/Security/kSecGuestAttributeHash
	KSecGuestAttributeHash string
	// KSecGuestAttributeMachPort is not implemented.
	//
	// See: https://developer.apple.com/documentation/Security/kSecGuestAttributeMachPort
	KSecGuestAttributeMachPort string
	// KSecGuestAttributePid is a key whose value is an integer of type `pid_t` representing a process ID (PID), usually of the kernel’s guest.
	//
	// See: https://developer.apple.com/documentation/Security/kSecGuestAttributePid
	KSecGuestAttributePid string
	// KSecGuestAttributeSubarchitecture is a key whose value is a number representing the CPU subtype under which the guest code is designed to run.
	//
	// See: https://developer.apple.com/documentation/Security/kSecGuestAttributeSubarchitecture
	KSecGuestAttributeSubarchitecture string
	// KSecIdentityDomainDefault is the system-wide default identity.
	//
	// See: https://developer.apple.com/documentation/Security/kSecIdentityDomainDefault
	KSecIdentityDomainDefault string
	// KSecIdentityDomainKerberosKDC is kerberos Key Distribution Center (KDC) identity.
	//
	// See: https://developer.apple.com/documentation/Security/kSecIdentityDomainKerberosKDC
	KSecIdentityDomainKerberosKDC string
	// KSecImportExportAccess is an initial access control list represented by a [SecAccess] object.
	//
	// See: https://developer.apple.com/documentation/Security/kSecImportExportAccess
	KSecImportExportAccess string
	// KSecImportExportKeychain is a keychain represented by a SecKeychainRef to be used as the target when importing or exporting.
	//
	// See: https://developer.apple.com/documentation/Security/kSecImportExportKeychain
	KSecImportExportKeychain string
	// KSecImportExportPassphrase is a passphrase (represented by a [CFStringRef] object) to be used when exporting to or importing from PKCS#12 format.
	//
	// See: https://developer.apple.com/documentation/Security/kSecImportExportPassphrase
	KSecImportExportPassphrase string
	// KSecImportItemCertChain is certificate list.
	//
	// See: https://developer.apple.com/documentation/Security/kSecImportItemCertChain
	KSecImportItemCertChain string
	// KSecImportItemIdentity is identity object.
	//
	// See: https://developer.apple.com/documentation/Security/kSecImportItemIdentity
	KSecImportItemIdentity string
	// KSecImportItemKeyID is key ID.
	//
	// See: https://developer.apple.com/documentation/Security/kSecImportItemKeyID
	KSecImportItemKeyID string
	// KSecImportItemLabel is item label.
	//
	// See: https://developer.apple.com/documentation/Security/kSecImportItemLabel
	KSecImportItemLabel string
	// KSecImportItemTrust is trust management object.
	//
	// See: https://developer.apple.com/documentation/Security/kSecImportItemTrust
	KSecImportItemTrust string
	// See: https://developer.apple.com/documentation/Security/kSecImportToMemoryOnly
	KSecImportToMemoryOnly string
	// KSecInputIsDigest is the input is a digest of the original data.
	//
	// See: https://developer.apple.com/documentation/Security/kSecInputIsDigest
	KSecInputIsDigest string
	// KSecInputIsPlainText is the input is plain text.
	//
	// See: https://developer.apple.com/documentation/Security/kSecInputIsPlainText
	KSecInputIsPlainText string
	// KSecKeyAttributeName is the cryptographic key associated with a transform.
	//
	// See: https://developer.apple.com/documentation/Security/kSecKeyAttributeName
	KSecKeyAttributeName string
	// KSecMatchCaseInsensitive is a key whose value is a Boolean indicating whether case-insensitive matching is performed.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchCaseInsensitive
	KSecMatchCaseInsensitive string
	// KSecMatchDiacriticInsensitive is a key whose value is a Boolean indicating whether diacritic-insensitive matching is performed.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchDiacriticInsensitive
	KSecMatchDiacriticInsensitive string
	// KSecMatchEmailAddressIfPresent is a key whose value is a string to match against a certificate or identity’s email address.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchEmailAddressIfPresent
	KSecMatchEmailAddressIfPresent string
	// See: https://developer.apple.com/documentation/Security/kSecMatchHostOrSubdomainOfHost
	KSecMatchHostOrSubdomainOfHost string
	// KSecMatchIssuers is a key whose value is a string to match against a certificate or identity’s issuers.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchIssuers
	KSecMatchIssuers string
	// KSecMatchItemList is a key whose value indicates a list of items to search.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchItemList
	KSecMatchItemList string
	// KSecMatchLimit is a key whose value indicates the match limit.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchLimit
	KSecMatchLimit string
	// KSecMatchLimitAll is a value that corresponds to matching an unlimited number of items.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchLimitAll
	KSecMatchLimitAll string
	// KSecMatchLimitOne is a value that corresponds to matching exactly one item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchLimitOne
	KSecMatchLimitOne string
	// KSecMatchPolicy is a key whose value indicates a policy with which a matching certificate or identity must verify.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchPolicy
	KSecMatchPolicy string
	// KSecMatchSearchList is a key whose value indicates a list of items to search.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchSearchList
	KSecMatchSearchList string
	// KSecMatchSubjectContains is a key whose value is a string to look for in a certificate or identity’s subject.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchSubjectContains
	KSecMatchSubjectContains string
	// KSecMatchSubjectEndsWith is a key whose value is a string to match against the end of a certificate or identity’s subject.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchSubjectEndsWith
	KSecMatchSubjectEndsWith string
	// KSecMatchSubjectStartsWith is a key whose value is a string to match against the beginning of a certificate or identity’s subject.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchSubjectStartsWith
	KSecMatchSubjectStartsWith string
	// KSecMatchSubjectWholeString is a key whose value is a string to exactly match a certificate or identity’s subject.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchSubjectWholeString
	KSecMatchSubjectWholeString string
	// KSecMatchTrustedOnly is a key whose value is a Boolean indicating whether untrusted certificates should be returned.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchTrustedOnly
	KSecMatchTrustedOnly string
	// KSecMatchValidOnDate is a key whose value indicates the validity date.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchValidOnDate
	KSecMatchValidOnDate string
	// KSecMatchWidthInsensitive is a key whose value is a Boolean indicating whether width-insensitive matching is performed.
	//
	// See: https://developer.apple.com/documentation/Security/kSecMatchWidthInsensitive
	KSecMatchWidthInsensitive string
	// See: https://developer.apple.com/documentation/Security/kSecOIDADC_CERT_POLICY
	KSecOIDADC_CERT_POLICY string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_CERT_POLICY
	KSecOIDAPPLE_CERT_POLICY string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EKU_CODE_SIGNING
	KSecOIDAPPLE_EKU_CODE_SIGNING string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EKU_CODE_SIGNING_DEV
	KSecOIDAPPLE_EKU_CODE_SIGNING_DEV string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EKU_ICHAT_ENCRYPTION
	KSecOIDAPPLE_EKU_ICHAT_ENCRYPTION string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EKU_ICHAT_SIGNING
	KSecOIDAPPLE_EKU_ICHAT_SIGNING string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EKU_RESOURCE_SIGNING
	KSecOIDAPPLE_EKU_RESOURCE_SIGNING string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EKU_SYSTEM_IDENTITY
	KSecOIDAPPLE_EKU_SYSTEM_IDENTITY string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EXTENSION
	KSecOIDAPPLE_EXTENSION string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EXTENSION_AAI_INTERMEDIATE
	KSecOIDAPPLE_EXTENSION_AAI_INTERMEDIATE string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EXTENSION_ADC_APPLE_SIGNING
	KSecOIDAPPLE_EXTENSION_ADC_APPLE_SIGNING string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EXTENSION_ADC_DEV_SIGNING
	KSecOIDAPPLE_EXTENSION_ADC_DEV_SIGNING string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EXTENSION_APPLEID_INTERMEDIATE
	KSecOIDAPPLE_EXTENSION_APPLEID_INTERMEDIATE string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EXTENSION_APPLE_SIGNING
	KSecOIDAPPLE_EXTENSION_APPLE_SIGNING string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EXTENSION_CODE_SIGNING
	KSecOIDAPPLE_EXTENSION_CODE_SIGNING string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EXTENSION_INTERMEDIATE_MARKER
	KSecOIDAPPLE_EXTENSION_INTERMEDIATE_MARKER string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EXTENSION_ITMS_INTERMEDIATE
	KSecOIDAPPLE_EXTENSION_ITMS_INTERMEDIATE string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAPPLE_EXTENSION_WWDR_INTERMEDIATE
	KSecOIDAPPLE_EXTENSION_WWDR_INTERMEDIATE string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAuthorityInfoAccess
	KSecOIDAuthorityInfoAccess string
	// See: https://developer.apple.com/documentation/Security/kSecOIDAuthorityKeyIdentifier
	KSecOIDAuthorityKeyIdentifier string
	// See: https://developer.apple.com/documentation/Security/kSecOIDBasicConstraints
	KSecOIDBasicConstraints string
	// See: https://developer.apple.com/documentation/Security/kSecOIDBiometricInfo
	KSecOIDBiometricInfo string
	// See: https://developer.apple.com/documentation/Security/kSecOIDCSSMKeyStruct
	KSecOIDCSSMKeyStruct string
	// See: https://developer.apple.com/documentation/Security/kSecOIDCertIssuer
	KSecOIDCertIssuer string
	// See: https://developer.apple.com/documentation/Security/kSecOIDCertificatePolicies
	KSecOIDCertificatePolicies string
	// See: https://developer.apple.com/documentation/Security/kSecOIDClientAuth
	KSecOIDClientAuth string
	// See: https://developer.apple.com/documentation/Security/kSecOIDCollectiveStateProvinceName
	KSecOIDCollectiveStateProvinceName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDCollectiveStreetAddress
	KSecOIDCollectiveStreetAddress string
	// See: https://developer.apple.com/documentation/Security/kSecOIDCommonName
	KSecOIDCommonName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDCountryName
	KSecOIDCountryName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDCrlDistributionPoints
	KSecOIDCrlDistributionPoints string
	// See: https://developer.apple.com/documentation/Security/kSecOIDCrlNumber
	KSecOIDCrlNumber string
	// See: https://developer.apple.com/documentation/Security/kSecOIDCrlReason
	KSecOIDCrlReason string
	// See: https://developer.apple.com/documentation/Security/kSecOIDDOTMAC_CERT_EMAIL_ENCRYPT
	KSecOIDDOTMAC_CERT_EMAIL_ENCRYPT string
	// See: https://developer.apple.com/documentation/Security/kSecOIDDOTMAC_CERT_EMAIL_SIGN
	KSecOIDDOTMAC_CERT_EMAIL_SIGN string
	// See: https://developer.apple.com/documentation/Security/kSecOIDDOTMAC_CERT_EXTENSION
	KSecOIDDOTMAC_CERT_EXTENSION string
	// See: https://developer.apple.com/documentation/Security/kSecOIDDOTMAC_CERT_IDENTITY
	KSecOIDDOTMAC_CERT_IDENTITY string
	// See: https://developer.apple.com/documentation/Security/kSecOIDDOTMAC_CERT_POLICY
	KSecOIDDOTMAC_CERT_POLICY string
	// See: https://developer.apple.com/documentation/Security/kSecOIDDeltaCrlIndicator
	KSecOIDDeltaCrlIndicator string
	// See: https://developer.apple.com/documentation/Security/kSecOIDDescription
	KSecOIDDescription string
	// See: https://developer.apple.com/documentation/Security/kSecOIDEKU_IPSec
	KSecOIDEKU_IPSec string
	// See: https://developer.apple.com/documentation/Security/kSecOIDEmailAddress
	KSecOIDEmailAddress string
	// See: https://developer.apple.com/documentation/Security/kSecOIDEmailProtection
	KSecOIDEmailProtection string
	// See: https://developer.apple.com/documentation/Security/kSecOIDExtendedKeyUsage
	KSecOIDExtendedKeyUsage string
	// See: https://developer.apple.com/documentation/Security/kSecOIDExtendedKeyUsageAny
	KSecOIDExtendedKeyUsageAny string
	// See: https://developer.apple.com/documentation/Security/kSecOIDExtendedUseCodeSigning
	KSecOIDExtendedUseCodeSigning string
	// See: https://developer.apple.com/documentation/Security/kSecOIDGivenName
	KSecOIDGivenName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDHoldInstructionCode
	KSecOIDHoldInstructionCode string
	// See: https://developer.apple.com/documentation/Security/kSecOIDInvalidityDate
	KSecOIDInvalidityDate string
	// See: https://developer.apple.com/documentation/Security/kSecOIDIssuerAltName
	KSecOIDIssuerAltName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDIssuingDistributionPoint
	KSecOIDIssuingDistributionPoint string
	// See: https://developer.apple.com/documentation/Security/kSecOIDIssuingDistributionPoints
	KSecOIDIssuingDistributionPoints string
	// See: https://developer.apple.com/documentation/Security/kSecOIDKERBv5_PKINIT_KP_CLIENT_AUTH
	KSecOIDKERBv5_PKINIT_KP_CLIENT_AUTH string
	// See: https://developer.apple.com/documentation/Security/kSecOIDKERBv5_PKINIT_KP_KDC
	KSecOIDKERBv5_PKINIT_KP_KDC string
	// See: https://developer.apple.com/documentation/Security/kSecOIDKeyUsage
	KSecOIDKeyUsage string
	// See: https://developer.apple.com/documentation/Security/kSecOIDLocalityName
	KSecOIDLocalityName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDMS_NTPrincipalName
	KSecOIDMS_NTPrincipalName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDMicrosoftSGC
	KSecOIDMicrosoftSGC string
	// See: https://developer.apple.com/documentation/Security/kSecOIDNameConstraints
	KSecOIDNameConstraints string
	// See: https://developer.apple.com/documentation/Security/kSecOIDNetscapeCertSequence
	KSecOIDNetscapeCertSequence string
	// See: https://developer.apple.com/documentation/Security/kSecOIDNetscapeCertType
	KSecOIDNetscapeCertType string
	// See: https://developer.apple.com/documentation/Security/kSecOIDNetscapeSGC
	KSecOIDNetscapeSGC string
	// See: https://developer.apple.com/documentation/Security/kSecOIDOCSPSigning
	KSecOIDOCSPSigning string
	// See: https://developer.apple.com/documentation/Security/kSecOIDOrganizationName
	KSecOIDOrganizationName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDOrganizationalUnitName
	KSecOIDOrganizationalUnitName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDPolicyConstraints
	KSecOIDPolicyConstraints string
	// See: https://developer.apple.com/documentation/Security/kSecOIDPolicyMappings
	KSecOIDPolicyMappings string
	// See: https://developer.apple.com/documentation/Security/kSecOIDPrivateKeyUsagePeriod
	KSecOIDPrivateKeyUsagePeriod string
	// See: https://developer.apple.com/documentation/Security/kSecOIDQC_Statements
	KSecOIDQC_Statements string
	// See: https://developer.apple.com/documentation/Security/kSecOIDSRVName
	KSecOIDSRVName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDSerialNumber
	KSecOIDSerialNumber string
	// See: https://developer.apple.com/documentation/Security/kSecOIDServerAuth
	KSecOIDServerAuth string
	// See: https://developer.apple.com/documentation/Security/kSecOIDStateProvinceName
	KSecOIDStateProvinceName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDStreetAddress
	KSecOIDStreetAddress string
	// See: https://developer.apple.com/documentation/Security/kSecOIDSubjectAltName
	KSecOIDSubjectAltName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDSubjectDirectoryAttributes
	KSecOIDSubjectDirectoryAttributes string
	// See: https://developer.apple.com/documentation/Security/kSecOIDSubjectEmailAddress
	KSecOIDSubjectEmailAddress string
	// See: https://developer.apple.com/documentation/Security/kSecOIDSubjectInfoAccess
	KSecOIDSubjectInfoAccess string
	// See: https://developer.apple.com/documentation/Security/kSecOIDSubjectKeyIdentifier
	KSecOIDSubjectKeyIdentifier string
	// See: https://developer.apple.com/documentation/Security/kSecOIDSubjectPicture
	KSecOIDSubjectPicture string
	// See: https://developer.apple.com/documentation/Security/kSecOIDSubjectSignatureBitmap
	KSecOIDSubjectSignatureBitmap string
	// See: https://developer.apple.com/documentation/Security/kSecOIDSurname
	KSecOIDSurname string
	// See: https://developer.apple.com/documentation/Security/kSecOIDTimeStamping
	KSecOIDTimeStamping string
	// See: https://developer.apple.com/documentation/Security/kSecOIDTitle
	KSecOIDTitle string
	// See: https://developer.apple.com/documentation/Security/kSecOIDUseExemptions
	KSecOIDUseExemptions string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1CertificateIssuerUniqueId
	KSecOIDX509V1CertificateIssuerUniqueId string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1CertificateSubjectUniqueId
	KSecOIDX509V1CertificateSubjectUniqueId string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1IssuerName
	KSecOIDX509V1IssuerName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1IssuerNameCStruct
	KSecOIDX509V1IssuerNameCStruct string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1IssuerNameLDAP
	KSecOIDX509V1IssuerNameLDAP string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1IssuerNameStd
	KSecOIDX509V1IssuerNameStd string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SerialNumber
	KSecOIDX509V1SerialNumber string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1Signature
	KSecOIDX509V1Signature string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SignatureAlgorithm
	KSecOIDX509V1SignatureAlgorithm string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SignatureAlgorithmParameters
	KSecOIDX509V1SignatureAlgorithmParameters string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SignatureAlgorithmTBS
	KSecOIDX509V1SignatureAlgorithmTBS string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SignatureCStruct
	KSecOIDX509V1SignatureCStruct string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SignatureStruct
	KSecOIDX509V1SignatureStruct string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SubjectName
	KSecOIDX509V1SubjectName string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SubjectNameCStruct
	KSecOIDX509V1SubjectNameCStruct string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SubjectNameLDAP
	KSecOIDX509V1SubjectNameLDAP string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SubjectNameStd
	KSecOIDX509V1SubjectNameStd string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SubjectPublicKey
	KSecOIDX509V1SubjectPublicKey string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SubjectPublicKeyAlgorithm
	KSecOIDX509V1SubjectPublicKeyAlgorithm string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SubjectPublicKeyAlgorithmParameters
	KSecOIDX509V1SubjectPublicKeyAlgorithmParameters string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1SubjectPublicKeyCStruct
	KSecOIDX509V1SubjectPublicKeyCStruct string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1ValidityNotAfter
	KSecOIDX509V1ValidityNotAfter string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1ValidityNotBefore
	KSecOIDX509V1ValidityNotBefore string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V1Version
	KSecOIDX509V1Version string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3Certificate
	KSecOIDX509V3Certificate string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3CertificateCStruct
	KSecOIDX509V3CertificateCStruct string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3CertificateExtensionCStruct
	KSecOIDX509V3CertificateExtensionCStruct string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3CertificateExtensionCritical
	KSecOIDX509V3CertificateExtensionCritical string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3CertificateExtensionId
	KSecOIDX509V3CertificateExtensionId string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3CertificateExtensionStruct
	KSecOIDX509V3CertificateExtensionStruct string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3CertificateExtensionType
	KSecOIDX509V3CertificateExtensionType string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3CertificateExtensionValue
	KSecOIDX509V3CertificateExtensionValue string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3CertificateExtensionsCStruct
	KSecOIDX509V3CertificateExtensionsCStruct string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3CertificateExtensionsStruct
	KSecOIDX509V3CertificateExtensionsStruct string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3CertificateNumberOfExtensions
	KSecOIDX509V3CertificateNumberOfExtensions string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3SignedCertificate
	KSecOIDX509V3SignedCertificate string
	// See: https://developer.apple.com/documentation/Security/kSecOIDX509V3SignedCertificateCStruct
	KSecOIDX509V3SignedCertificateCStruct string
	// KSecPolicyAppleCodeSigning is policy for use in evaluating Apple code signing certificates.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleCodeSigning
	KSecPolicyAppleCodeSigning string
	// KSecPolicyAppleEAP is extensible Authentication Protocol. Functionally identical to SSL policy. A separate OID is provided to facilitate per-policy, per-certificate trust settings using the [SecTrust] mechanism.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleEAP
	KSecPolicyAppleEAP string
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleEAPClient
	KSecPolicyAppleEAPClient string
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleEAPServer
	KSecPolicyAppleEAPServer string
	// KSecPolicyAppleIDValidation is policy for use in evaluating Apple ID certificates.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleIDValidation
	KSecPolicyAppleIDValidation string
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleIPSecClient
	KSecPolicyAppleIPSecClient string
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleIPSecServer
	KSecPolicyAppleIPSecServer string
	// KSecPolicyAppleIPsec is policy for use in IPsec communication. Functionally identical to SSL policy. A separate OID is provided to facilitate per-policy, per-certificate trust settings using the [SecTrust] mechanism.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleIPsec
	KSecPolicyAppleIPsec string
	// KSecPolicyApplePKINITClient is kerberos Pkinit client certificate validation.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyApplePKINITClient
	KSecPolicyApplePKINITClient string
	// KSecPolicyApplePKINITServer is kerberos Pkinit server certificate validation.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyApplePKINITServer
	KSecPolicyApplePKINITServer string
	// See: https://developer.apple.com/documentation/Security/kSecPolicyApplePassbookSigning
	KSecPolicyApplePassbookSigning string
	// See: https://developer.apple.com/documentation/Security/kSecPolicyApplePayIssuerEncryption
	KSecPolicyApplePayIssuerEncryption string
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleRevocation
	KSecPolicyAppleRevocation string
	// KSecPolicyAppleSMIME is basic X509 plus email address verification and [KeyUsage] enforcement per RFC 2632.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleSMIME
	KSecPolicyAppleSMIME string
	// KSecPolicyAppleSSL is basic X509 plus host name verification per RFC 2818.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleSSL
	KSecPolicyAppleSSL string
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleSSLClient
	KSecPolicyAppleSSLClient string
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleSSLServer
	KSecPolicyAppleSSLServer string
	// KSecPolicyAppleTimeStamping is policy that causes evaluation of the validity of the time stamp on a signature. This can be used to allow verification that a certificate was valid at the time that something was signed with that certificate even if the certificate is no longer valid.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleTimeStamping
	KSecPolicyAppleTimeStamping string
	// KSecPolicyAppleX509Basic is basic X509-style certificate evaluation.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyAppleX509Basic
	KSecPolicyAppleX509Basic string
	// KSecPolicyClient is if true, indicates this policy should be evaluated against the client certificate. If false, the policy is evaluated against the certificate for the server. Default is false.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyClient
	KSecPolicyClient string
	// KSecPolicyKU_CRLSign is if true, the certificate’s key usage must allow it to be used for signing certificate revocation lists (CRLs).
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyKU_CRLSign
	KSecPolicyKU_CRLSign string
	// KSecPolicyKU_DataEncipherment is if true, the certificate’s key usage must allow it to be used for data encryption.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyKU_DataEncipherment
	KSecPolicyKU_DataEncipherment string
	// KSecPolicyKU_DecipherOnly is if true, the certificate’s key usage must allow it to be used for decryption.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyKU_DecipherOnly
	KSecPolicyKU_DecipherOnly string
	// KSecPolicyKU_DigitalSignature is if true, the certificate’s key usage must allow it to be used for signing.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyKU_DigitalSignature
	KSecPolicyKU_DigitalSignature string
	// KSecPolicyKU_EncipherOnly is if true, the certificate’s key usage must allow it to be used for encryption.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyKU_EncipherOnly
	KSecPolicyKU_EncipherOnly string
	// KSecPolicyKU_KeyAgreement is if true, the certificate’s key usage must allow it to be used for key agreement.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyKU_KeyAgreement
	KSecPolicyKU_KeyAgreement string
	// KSecPolicyKU_KeyCertSign is if true, the certificate’s key usage must allow it to be used for signing certificates.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyKU_KeyCertSign
	KSecPolicyKU_KeyCertSign string
	// KSecPolicyKU_KeyEncipherment is if true, the certificate’s key usage must allow it to be used for key encryption.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyKU_KeyEncipherment
	KSecPolicyKU_KeyEncipherment string
	// KSecPolicyKU_NonRepudiation is if true, the certificate’s key usage must allow it to be used for non-repudiation.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyKU_NonRepudiation
	KSecPolicyKU_NonRepudiation string
	// KSecPolicyMacAppStoreReceipt is policy for use in evaluating Mac App Store receipts.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyMacAppStoreReceipt
	KSecPolicyMacAppStoreReceipt string
	// KSecPolicyName is a name ([CFStringRef]) that the certificate must match to satisfy this policy. For SSL/TLS, this specifies the server name which must match the common name of the certificate. For S/MIME, this specifies the RFC 822 email address.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyName
	KSecPolicyName string
	// KSecPolicyOid is the object identifier that defines the policy type ([CFStringRef]). All policies have a value for this key.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPolicyOid
	KSecPolicyOid string
	// See: https://developer.apple.com/documentation/Security/kSecPolicyRevocationFlags
	KSecPolicyRevocationFlags string
	// See: https://developer.apple.com/documentation/Security/kSecPolicyTeamIdentifier
	KSecPolicyTeamIdentifier string
	// KSecPrivateKeyAttrs is a key whose value is a dictionary of cryptographic key attributes specific to a private key.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPrivateKeyAttrs
	KSecPrivateKeyAttrs string
	// KSecPropertyKeyLabel is a key whose value is the label for a certificate property.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyKeyLabel
	KSecPropertyKeyLabel string
	// KSecPropertyKeyLocalizedLabel is a key whose value is the localized label for a certificate property.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyKeyLocalizedLabel
	KSecPropertyKeyLocalizedLabel string
	// KSecPropertyKeyType is a key whose value indicates the type of certificate property.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyKeyType
	KSecPropertyKeyType string
	// KSecPropertyKeyValue is a key whose value is the value for a certificate property.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyKeyValue
	KSecPropertyKeyValue string
	// KSecPropertyTypeArray is specifies a key whose value is an array.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyTypeArray
	KSecPropertyTypeArray string
	// KSecPropertyTypeData is a key whose value is a data object.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyTypeData
	KSecPropertyTypeData string
	// KSecPropertyTypeDate is specifies a key whose value is a string containing a date (or a string listing the bytes of an invalid date).
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyTypeDate
	KSecPropertyTypeDate string
	// KSecPropertyTypeError is specifies a key whose value is a string containing the reason for a trust evaluation failure.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyTypeError
	KSecPropertyTypeError string
	// KSecPropertyTypeNumber is specifies a key whose value is a number.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyTypeNumber
	KSecPropertyTypeNumber string
	// KSecPropertyTypeSection is a key whose value is a string describing the name of a field in the certificate (`CFSTR("Subject Name")`, for example).
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyTypeSection
	KSecPropertyTypeSection string
	// KSecPropertyTypeString is a key whose value is a string.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyTypeString
	KSecPropertyTypeString string
	// KSecPropertyTypeSuccess is a key whose value is a string describing a trust evaluation success.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyTypeSuccess
	KSecPropertyTypeSuccess string
	// KSecPropertyTypeTitle is specifies a key whose value is a string containing the title (display name) of the certificate.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyTypeTitle
	KSecPropertyTypeTitle string
	// KSecPropertyTypeURL is specifies a key whose value is a URL.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyTypeURL
	KSecPropertyTypeURL string
	// KSecPropertyTypeWarning is a key whose value is a string describing a trust evaluation warning.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPropertyTypeWarning
	KSecPropertyTypeWarning string
	// KSecPublicKeyAttrs is a key whose value is a dictionary of cryptographic key attributes specific to a public key.
	//
	// See: https://developer.apple.com/documentation/Security/kSecPublicKeyAttrs
	KSecPublicKeyAttrs string
	// KSecReturnAttributes is a key whose value is a Boolean indicating whether or not to return item attributes.
	//
	// See: https://developer.apple.com/documentation/Security/kSecReturnAttributes
	KSecReturnAttributes string
	// KSecReturnData is a key whose value is a Boolean that indicates whether or not to return item data.
	//
	// See: https://developer.apple.com/documentation/Security/kSecReturnData
	KSecReturnData string
	// KSecReturnPersistentRef is a key whose value is a Boolean indicating whether or not to return a persistent reference to an item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecReturnPersistentRef
	KSecReturnPersistentRef string
	// KSecReturnRef is a key whose value is a Boolean indicating whether or not to return a reference to an item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecReturnRef
	KSecReturnRef string
	// KSecSharedPassword is a dictionary key whose value is the shared password.
	//
	// See: https://developer.apple.com/documentation/Security/kSecSharedPassword
	KSecSharedPassword string
	// KSecSignatureAttributeName is the cryptographic signature associated with a transform.
	//
	// See: https://developer.apple.com/documentation/Security/kSecSignatureAttributeName
	KSecSignatureAttributeName string
	// KSecTransformAbortOriginatorKey is the key in an error’s `userInfo` dictionary whose value indicates the transform that caused the chain to abort.
	//
	// See: https://developer.apple.com/documentation/Security/kSecTransformAbortOriginatorKey
	KSecTransformAbortOriginatorKey string
	// KSecTransformErrorDomain is the domain of any error object created by a transform on failure.
	//
	// See: https://developer.apple.com/documentation/Security/kSecTransformErrorDomain
	KSecTransformErrorDomain string
	// KSecTransformPreviousErrorKey is the key in an error’s `userInfo` dictionary whose value specifies the previous error when multiple errors occur during transform evaluation.
	//
	// See: https://developer.apple.com/documentation/Security/kSecTransformPreviousErrorKey
	KSecTransformPreviousErrorKey string
	// KSecTrustCertificateTransparency is a key whose value is a Boolean used to indicate Certificate Transparency.
	//
	// See: https://developer.apple.com/documentation/Security/kSecTrustCertificateTransparency
	KSecTrustCertificateTransparency string
	// KSecTrustEvaluationDate is a key whose value indicates the time that the trust evaluation took place.
	//
	// See: https://developer.apple.com/documentation/Security/kSecTrustEvaluationDate
	KSecTrustEvaluationDate string
	// KSecTrustExtendedValidation is a key whose value is a Boolean used to indicate Extended Validation.
	//
	// See: https://developer.apple.com/documentation/Security/kSecTrustExtendedValidation
	KSecTrustExtendedValidation string
	// KSecTrustOrganizationName is a key whose value is the organization name field of the subject of the leaf certificate.
	//
	// See: https://developer.apple.com/documentation/Security/kSecTrustOrganizationName
	KSecTrustOrganizationName string
	// See: https://developer.apple.com/documentation/Security/kSecTrustQCStatements
	KSecTrustQCStatements string
	// See: https://developer.apple.com/documentation/Security/kSecTrustQWACValidation
	KSecTrustQWACValidation string
	// KSecTrustResultValue is a key whose value represents the trust evaluation result.
	//
	// See: https://developer.apple.com/documentation/Security/kSecTrustResultValue
	KSecTrustResultValue string
	// KSecTrustRevocationChecked is a key whose value indicates the outcome of revocation checking during trust evaluation.
	//
	// See: https://developer.apple.com/documentation/Security/kSecTrustRevocationChecked
	KSecTrustRevocationChecked string
	// KSecTrustRevocationValidUntilDate is a key whose value indicates the earliest date at which revocation information becomes stale.
	//
	// See: https://developer.apple.com/documentation/Security/kSecTrustRevocationValidUntilDate
	KSecTrustRevocationValidUntilDate string
	// KSecUseAuthenticationContext is a key whose value indicates a local authentication context to use.
	//
	// See: https://developer.apple.com/documentation/Security/kSecUseAuthenticationContext
	KSecUseAuthenticationContext string
	// KSecUseAuthenticationUI is a key whose value indicates whether the user is prompted for authentication.
	//
	// See: https://developer.apple.com/documentation/Security/kSecUseAuthenticationUI
	KSecUseAuthenticationUI string
	// KSecUseAuthenticationUISkip is a value that indicates items requiring user authentication should be skipped.
	//
	// See: https://developer.apple.com/documentation/Security/kSecUseAuthenticationUISkip
	KSecUseAuthenticationUISkip string
	// KSecUseDataProtectionKeychain is a key whose value indicates whether to treat macOS keychain items like iOS keychain items.
	//
	// See: https://developer.apple.com/documentation/Security/kSecUseDataProtectionKeychain
	KSecUseDataProtectionKeychain string
	// KSecUseKeychain is a key whose value is a keychain to operate on.
	//
	// See: https://developer.apple.com/documentation/Security/kSecUseKeychain
	KSecUseKeychain string
	// KSecValueData is a key whose value is the item’s data.
	//
	// See: https://developer.apple.com/documentation/Security/kSecValueData
	KSecValueData string
	// KSecValuePersistentRef is a key whose value is a persistent reference to the item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecValuePersistentRef
	KSecValuePersistentRef string
	// KSecValueRef is a key whose value is a reference to the item.
	//
	// See: https://developer.apple.com/documentation/Security/kSecValueRef
	KSecValueRef string
)

var (
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeCofactor
	KSecKeyAlgorithmECDHKeyExchangeCofactor SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeCofactorX963SHA1
	KSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA1 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeCofactorX963SHA224
	KSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA224 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeCofactorX963SHA256
	KSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA256 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeCofactorX963SHA384
	KSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA384 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeCofactorX963SHA512
	KSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA512 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeStandard
	KSecKeyAlgorithmECDHKeyExchangeStandard SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeStandardX963SHA1
	KSecKeyAlgorithmECDHKeyExchangeStandardX963SHA1 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeStandardX963SHA224
	KSecKeyAlgorithmECDHKeyExchangeStandardX963SHA224 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeStandardX963SHA256
	KSecKeyAlgorithmECDHKeyExchangeStandardX963SHA256 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeStandardX963SHA384
	KSecKeyAlgorithmECDHKeyExchangeStandardX963SHA384 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdhKeyExchangeStandardX963SHA512
	KSecKeyAlgorithmECDHKeyExchangeStandardX963SHA512 SecKeyAlgorithm
	// KSecKeyAlgorithmECDSASignatureDigestRFC4754 is an algorithm for generating message digest signatures.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestRFC4754
	KSecKeyAlgorithmECDSASignatureDigestRFC4754 SecKeyAlgorithm
	// KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA1 is an algorithm for generating signatures of SHA1 message digests.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestRFC4754SHA1
	KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA1 SecKeyAlgorithm
	// KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA224 is an algorithm for generating signatures of SHA224 message digests.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestRFC4754SHA224
	KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA224 SecKeyAlgorithm
	// KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA256 is an algorithm for generating signatures of SHA256 message digests.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestRFC4754SHA256
	KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA256 SecKeyAlgorithm
	// KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA384 is an algorithm for generating signatures of SHA384 message digests.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestRFC4754SHA384
	KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA384 SecKeyAlgorithm
	// KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA512 is an algorithm for generating signatures of SHA512 message digests.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestRFC4754SHA512
	KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA512 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestX962
	KSecKeyAlgorithmECDSASignatureDigestX962 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestX962SHA1
	KSecKeyAlgorithmECDSASignatureDigestX962SHA1 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestX962SHA224
	KSecKeyAlgorithmECDSASignatureDigestX962SHA224 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestX962SHA256
	KSecKeyAlgorithmECDSASignatureDigestX962SHA256 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestX962SHA384
	KSecKeyAlgorithmECDSASignatureDigestX962SHA384 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureDigestX962SHA512
	KSecKeyAlgorithmECDSASignatureDigestX962SHA512 SecKeyAlgorithm
	// KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA1 is an algorithm for generating message signatures by calculating and signing the SHA1 message digest.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureMessageRFC4754SHA1
	KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA1 SecKeyAlgorithm
	// KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA224 is an algorithm for generating message signatures by calculating and signing the SHA224 message digest.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureMessageRFC4754SHA224
	KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA224 SecKeyAlgorithm
	// KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA256 is an algorithm for generating message signatures by calculating and signing the SHA256 message digest.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureMessageRFC4754SHA256
	KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA256 SecKeyAlgorithm
	// KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA384 is an algorithm for generating message signatures by calculating and signing the SHA384 message digest.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureMessageRFC4754SHA384
	KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA384 SecKeyAlgorithm
	// KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA512 is an algorithm for generating message signatures by calculating and signing the SHA512 message digest.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureMessageRFC4754SHA512
	KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA512 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureMessageX962SHA1
	KSecKeyAlgorithmECDSASignatureMessageX962SHA1 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureMessageX962SHA224
	KSecKeyAlgorithmECDSASignatureMessageX962SHA224 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureMessageX962SHA256
	KSecKeyAlgorithmECDSASignatureMessageX962SHA256 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureMessageX962SHA384
	KSecKeyAlgorithmECDSASignatureMessageX962SHA384 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureMessageX962SHA512
	KSecKeyAlgorithmECDSASignatureMessageX962SHA512 SecKeyAlgorithm
	//
	// Deprecated: Deprecated since macOS 14.0.
	//
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/ecdsaSignatureRFC4754
	KSecKeyAlgorithmECDSASignatureRFC4754 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionCofactorVariableIVX963SHA224AESGCM
	KSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA224AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionCofactorVariableIVX963SHA256AESGCM
	KSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA256AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionCofactorVariableIVX963SHA384AESGCM
	KSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA384AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionCofactorVariableIVX963SHA512AESGCM
	KSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA512AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionCofactorX963SHA1AESGCM
	KSecKeyAlgorithmECIESEncryptionCofactorX963SHA1AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionCofactorX963SHA224AESGCM
	KSecKeyAlgorithmECIESEncryptionCofactorX963SHA224AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionCofactorX963SHA256AESGCM
	KSecKeyAlgorithmECIESEncryptionCofactorX963SHA256AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionCofactorX963SHA384AESGCM
	KSecKeyAlgorithmECIESEncryptionCofactorX963SHA384AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionCofactorX963SHA512AESGCM
	KSecKeyAlgorithmECIESEncryptionCofactorX963SHA512AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionStandardVariableIVX963SHA224AESGCM
	KSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA224AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionStandardVariableIVX963SHA256AESGCM
	KSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA256AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionStandardVariableIVX963SHA384AESGCM
	KSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA384AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionStandardVariableIVX963SHA512AESGCM
	KSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA512AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionStandardX963SHA1AESGCM
	KSecKeyAlgorithmECIESEncryptionStandardX963SHA1AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionStandardX963SHA224AESGCM
	KSecKeyAlgorithmECIESEncryptionStandardX963SHA224AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionStandardX963SHA256AESGCM
	KSecKeyAlgorithmECIESEncryptionStandardX963SHA256AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionStandardX963SHA384AESGCM
	KSecKeyAlgorithmECIESEncryptionStandardX963SHA384AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/eciesEncryptionStandardX963SHA512AESGCM
	KSecKeyAlgorithmECIESEncryptionStandardX963SHA512AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionOAEPSHA1
	KSecKeyAlgorithmRSAEncryptionOAEPSHA1 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionOAEPSHA1AESGCM
	KSecKeyAlgorithmRSAEncryptionOAEPSHA1AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionOAEPSHA224
	KSecKeyAlgorithmRSAEncryptionOAEPSHA224 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionOAEPSHA224AESGCM
	KSecKeyAlgorithmRSAEncryptionOAEPSHA224AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionOAEPSHA256
	KSecKeyAlgorithmRSAEncryptionOAEPSHA256 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionOAEPSHA256AESGCM
	KSecKeyAlgorithmRSAEncryptionOAEPSHA256AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionOAEPSHA384
	KSecKeyAlgorithmRSAEncryptionOAEPSHA384 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionOAEPSHA384AESGCM
	KSecKeyAlgorithmRSAEncryptionOAEPSHA384AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionOAEPSHA512
	KSecKeyAlgorithmRSAEncryptionOAEPSHA512 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionOAEPSHA512AESGCM
	KSecKeyAlgorithmRSAEncryptionOAEPSHA512AESGCM SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionPKCS1
	KSecKeyAlgorithmRSAEncryptionPKCS1 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaEncryptionRaw
	KSecKeyAlgorithmRSAEncryptionRaw SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureDigestPKCS1v15Raw
	KSecKeyAlgorithmRSASignatureDigestPKCS1v15Raw SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureDigestPKCS1v15SHA1
	KSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA1 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureDigestPKCS1v15SHA224
	KSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA224 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureDigestPKCS1v15SHA256
	KSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA256 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureDigestPKCS1v15SHA384
	KSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA384 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureDigestPKCS1v15SHA512
	KSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA512 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureDigestPSSSHA1
	KSecKeyAlgorithmRSASignatureDigestPSSSHA1 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureDigestPSSSHA224
	KSecKeyAlgorithmRSASignatureDigestPSSSHA224 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureDigestPSSSHA256
	KSecKeyAlgorithmRSASignatureDigestPSSSHA256 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureDigestPSSSHA384
	KSecKeyAlgorithmRSASignatureDigestPSSSHA384 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureDigestPSSSHA512
	KSecKeyAlgorithmRSASignatureDigestPSSSHA512 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureMessagePKCS1v15SHA1
	KSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA1 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureMessagePKCS1v15SHA224
	KSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA224 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureMessagePKCS1v15SHA256
	KSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA256 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureMessagePKCS1v15SHA384
	KSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA384 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureMessagePKCS1v15SHA512
	KSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA512 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureMessagePSSSHA1
	KSecKeyAlgorithmRSASignatureMessagePSSSHA1 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureMessagePSSSHA224
	KSecKeyAlgorithmRSASignatureMessagePSSSHA224 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureMessagePSSSHA256
	KSecKeyAlgorithmRSASignatureMessagePSSSHA256 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureMessagePSSSHA384
	KSecKeyAlgorithmRSASignatureMessagePSSSHA384 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureMessagePSSSHA512
	KSecKeyAlgorithmRSASignatureMessagePSSSHA512 SecKeyAlgorithm
	// See: https://developer.apple.com/documentation/Security/SecKeyAlgorithm/rsaSignatureRaw
	KSecKeyAlgorithmRSASignatureRaw SecKeyAlgorithm
)

var (
	// See: https://developer.apple.com/documentation/Security/SecKeyKeyExchangeParameter/requestedSize
	KSecKeyKeyExchangeParameterRequestedSize SecKeyKeyExchangeParameter
	// See: https://developer.apple.com/documentation/Security/SecKeyKeyExchangeParameter/sharedInfo
	KSecKeyKeyExchangeParameterSharedInfo SecKeyKeyExchangeParameter
)

var (
	// KSecRandomDefault is an alias for the default random number generator.
	//
	// See: https://developer.apple.com/documentation/Security/kSecRandomDefault
	KSecRandomDefault SecRandomRef
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "gGuidAppleCSP"); err == nil && ptr != 0 {
		gGuidAppleCSP = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "gGuidAppleCSPDL"); err == nil && ptr != 0 {
		gGuidAppleCSPDL = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "gGuidAppleDotMacDL"); err == nil && ptr != 0 {
		gGuidAppleDotMacDL = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "gGuidAppleDotMacTP"); err == nil && ptr != 0 {
		gGuidAppleDotMacTP = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "gGuidAppleFileDL"); err == nil && ptr != 0 {
		gGuidAppleFileDL = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "gGuidAppleLDAPDL"); err == nil && ptr != 0 {
		gGuidAppleLDAPDL = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "gGuidAppleSdCSPDL"); err == nil && ptr != 0 {
		gGuidAppleSdCSPDL = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "gGuidAppleX509CL"); err == nil && ptr != 0 {
		gGuidAppleX509CL = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "gGuidAppleX509TP"); err == nil && ptr != 0 {
		gGuidAppleX509TP = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "gGuidCssm"); err == nil && ptr != 0 {
		gGuidCssm = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSEncoderDigestAlgorithmSHA1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSEncoderDigestAlgorithmSHA1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCMSEncoderDigestAlgorithmSHA256"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCMSEncoderDigestAlgorithmSHA256 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationAny"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationAny = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationChangeACL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationChangeACL = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationChangeOwner"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationChangeOwner = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationDecrypt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationDecrypt = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationDelete"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationDelete = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationDerive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationDerive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationEncrypt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationEncrypt = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationExportClear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationExportClear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationExportWrapped"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationExportWrapped = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationGenKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationGenKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationImportClear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationImportClear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationImportWrapped"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationImportWrapped = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationIntegrity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationIntegrity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationKeychainCreate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationKeychainCreate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationKeychainDelete"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationKeychainDelete = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationKeychainItemDelete"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationKeychainItemDelete = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationKeychainItemInsert"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationKeychainItemInsert = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationKeychainItemModify"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationKeychainItemModify = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationKeychainItemRead"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationKeychainItemRead = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationLogin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationLogin = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationMAC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationMAC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationPartitionID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationPartitionID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecACLAuthorizationSign"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecACLAuthorizationSign = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAccess"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAccess = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAccessControl"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAccessControl = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAccessGroup"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAccessGroup = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAccessGroupToken"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAccessGroupToken = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAccessible"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAccessible = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAccessibleAfterFirstUnlock"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAccessibleAfterFirstUnlock = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAccessibleAfterFirstUnlockThisDeviceOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAccessibleAfterFirstUnlockThisDeviceOnly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAccessibleWhenPasscodeSetThisDeviceOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAccessibleWhenPasscodeSetThisDeviceOnly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAccessibleWhenUnlocked"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAccessibleWhenUnlocked = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAccessibleWhenUnlockedThisDeviceOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAccessibleWhenUnlockedThisDeviceOnly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAccount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAccount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrApplicationLabel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrApplicationLabel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrApplicationTag"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrApplicationTag = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAuthenticationType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAuthenticationType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAuthenticationTypeDPA"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAuthenticationTypeDPA = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAuthenticationTypeDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAuthenticationTypeDefault = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAuthenticationTypeHTMLForm"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAuthenticationTypeHTMLForm = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAuthenticationTypeHTTPBasic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAuthenticationTypeHTTPBasic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAuthenticationTypeHTTPDigest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAuthenticationTypeHTTPDigest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAuthenticationTypeMSN"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAuthenticationTypeMSN = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAuthenticationTypeNTLM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAuthenticationTypeNTLM = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrAuthenticationTypeRPA"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrAuthenticationTypeRPA = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrCanDecrypt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrCanDecrypt = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrCanDerive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrCanDerive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrCanEncrypt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrCanEncrypt = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrCanSign"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrCanSign = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrCanUnwrap"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrCanUnwrap = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrCanVerify"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrCanVerify = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrCanWrap"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrCanWrap = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrCertificateEncoding"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrCertificateEncoding = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrCertificateType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrCertificateType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrComment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrComment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrCreationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrCreationDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrCreator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrCreator = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrEffectiveKeySize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrEffectiveKeySize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrGeneric"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrGeneric = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrIsExtractable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrIsExtractable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrIsInvisible"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrIsInvisible = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrIsNegative"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrIsNegative = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrIsPermanent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrIsPermanent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrIsSensitive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrIsSensitive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrIssuer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrIssuer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyClassPrivate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyClassPrivate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyClassPublic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyClassPublic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyClassSymmetric"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyClassSymmetric = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeySizeInBits"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeySizeInBits = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyType3DES"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyType3DES = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyTypeAES"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyTypeAES = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyTypeCAST"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyTypeCAST = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyTypeDES"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyTypeDES = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyTypeDSA"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyTypeDSA = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyTypeECSECPrimeRandom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyTypeECSECPrimeRandom = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyTypeRC2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyTypeRC2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyTypeRC4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyTypeRC4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrKeyTypeRSA"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrKeyTypeRSA = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrLabel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrLabel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrModificationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrModificationDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrPRF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrPRF = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrPRFHmacAlgSHA1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrPRFHmacAlgSHA1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrPRFHmacAlgSHA224"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrPRFHmacAlgSHA224 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrPRFHmacAlgSHA256"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrPRFHmacAlgSHA256 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrPRFHmacAlgSHA384"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrPRFHmacAlgSHA384 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrPRFHmacAlgSHA512"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrPRFHmacAlgSHA512 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrPath"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrPath = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrPersistantReference"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrPersistantReference = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrPersistentReference"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrPersistentReference = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocol"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocol = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolAFP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolAFP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolAppleTalk"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolAppleTalk = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolDAAP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolDAAP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolEPPC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolEPPC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolFTP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolFTP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolFTPAccount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolFTPAccount = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolFTPProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolFTPProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolFTPS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolFTPS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolHTTP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolHTTP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolHTTPProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolHTTPProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolHTTPS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolHTTPS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolHTTPSProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolHTTPSProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolIMAP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolIMAP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolIMAPS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolIMAPS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolIPP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolIPP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolIRC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolIRC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolIRCS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolIRCS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolLDAP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolLDAP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolLDAPS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolLDAPS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolNNTP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolNNTP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolNNTPS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolNNTPS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolPOP3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolPOP3 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolPOP3S"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolPOP3S = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolRTSP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolRTSP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolRTSPProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolRTSPProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolSMB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolSMB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolSMTP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolSMTP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolSOCKS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolSOCKS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolSSH"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolSSH = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolTelnet"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolTelnet = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrProtocolTelnetS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrProtocolTelnetS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrPublicKeyHash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrPublicKeyHash = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrRounds"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrRounds = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrSalt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrSalt = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrSecurityDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrSecurityDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrSerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrSerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrServer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrServer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrService"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrService = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrSubject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrSubject = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrSubjectKeyID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrSubjectKeyID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrSyncViewHint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrSyncViewHint = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrSynchronizable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrSynchronizable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrSynchronizableAny"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrSynchronizableAny = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrTokenID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrTokenID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrTokenIDSecureEnclave"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrTokenIDSecureEnclave = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecAttrType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecAttrType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorArchitecture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorArchitecture = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorGuestAttributes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorGuestAttributes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorInfoPlist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorInfoPlist = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorPath"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorPath = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorPattern"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorPattern = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorRequirementSyntax"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorRequirementSyntax = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorResourceAdded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorResourceAdded = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorResourceAltered"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorResourceAltered = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorResourceMissing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorResourceMissing = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorResourceRecursive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorResourceRecursive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorResourceSeal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorResourceSeal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCFErrorResourceSideband"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCFErrorResourceSideband = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecClass = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecClassCertificate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecClassCertificate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecClassGenericPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecClassGenericPassword = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecClassIdentity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecClassIdentity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecClassInternetPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecClassInternetPassword = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecClassKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecClassKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeAttributeArchitecture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeAttributeArchitecture = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeAttributeBundleVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeAttributeBundleVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeAttributeSubarchitecture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeAttributeSubarchitecture = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeAttributeUniversalFileOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeAttributeUniversalFileOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoCMS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoCMS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoCdHashes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoCdHashes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoCertificates"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoCertificates = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoChangedFiles"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoChangedFiles = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoDefaultDesignatedLightweightCodeRequirement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoDefaultDesignatedLightweightCodeRequirement = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoDesignatedRequirement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoDesignatedRequirement = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoDigestAlgorithm"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoDigestAlgorithm = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoDigestAlgorithms"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoDigestAlgorithms = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoEntitlements"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoEntitlements = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoEntitlementsDict"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoEntitlementsDict = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoFlags"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoFlags = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoFormat = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoImplicitDesignatedRequirement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoImplicitDesignatedRequirement = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoMainExecutable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoMainExecutable = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoPList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoPList = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoPlatformIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoPlatformIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoRequirementData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoRequirementData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoRequirements"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoRequirements = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoRuntimeVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoRuntimeVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoSource = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoStapledNotarizationTicket"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoStapledNotarizationTicket = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoStatus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoStatus = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoTeamIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoTeamIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoTime = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoTimestamp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoTimestamp = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoTrust"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoTrust = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecCodeInfoUnique"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecCodeInfoUnique = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecDecodeTypeAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecDecodeTypeAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecGuestAttributeArchitecture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecGuestAttributeArchitecture = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecGuestAttributeAudit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecGuestAttributeAudit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecGuestAttributeCanonical"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecGuestAttributeCanonical = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecGuestAttributeDynamicCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecGuestAttributeDynamicCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecGuestAttributeDynamicCodeInfoPlist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecGuestAttributeDynamicCodeInfoPlist = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecGuestAttributeHash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecGuestAttributeHash = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecGuestAttributeMachPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecGuestAttributeMachPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecGuestAttributePid"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecGuestAttributePid = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecGuestAttributeSubarchitecture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecGuestAttributeSubarchitecture = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecIdentityDomainDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecIdentityDomainDefault = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecIdentityDomainKerberosKDC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecIdentityDomainKerberosKDC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecImportExportAccess"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecImportExportAccess = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecImportExportKeychain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecImportExportKeychain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecImportExportPassphrase"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecImportExportPassphrase = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecImportItemCertChain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecImportItemCertChain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecImportItemIdentity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecImportItemIdentity = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecImportItemKeyID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecImportItemKeyID = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecImportItemLabel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecImportItemLabel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecImportItemTrust"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecImportItemTrust = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecImportToMemoryOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecImportToMemoryOnly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecInputIsDigest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecInputIsDigest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecInputIsPlainText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecInputIsPlainText = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeCofactor"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeCofactor = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA224"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA224 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA256"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA256 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA384"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA384 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA512"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeCofactorX963SHA512 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeStandard"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeStandard = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeStandardX963SHA1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeStandardX963SHA1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeStandardX963SHA224"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeStandardX963SHA224 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeStandardX963SHA256"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeStandardX963SHA256 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeStandardX963SHA384"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeStandardX963SHA384 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDHKeyExchangeStandardX963SHA512"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDHKeyExchangeStandardX963SHA512 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestRFC4754"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestRFC4754 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestRFC4754SHA1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestRFC4754SHA224"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA224 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestRFC4754SHA256"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA256 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestRFC4754SHA384"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA384 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestRFC4754SHA512"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestRFC4754SHA512 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestX962"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestX962 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestX962SHA1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestX962SHA1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestX962SHA224"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestX962SHA224 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestX962SHA256"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestX962SHA256 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestX962SHA384"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestX962SHA384 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureDigestX962SHA512"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureDigestX962SHA512 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureMessageRFC4754SHA1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureMessageRFC4754SHA224"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA224 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureMessageRFC4754SHA256"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA256 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureMessageRFC4754SHA384"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA384 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureMessageRFC4754SHA512"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureMessageRFC4754SHA512 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureMessageX962SHA1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureMessageX962SHA1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureMessageX962SHA224"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureMessageX962SHA224 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureMessageX962SHA256"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureMessageX962SHA256 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureMessageX962SHA384"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureMessageX962SHA384 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureMessageX962SHA512"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureMessageX962SHA512 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECDSASignatureRFC4754"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECDSASignatureRFC4754 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA224AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA224AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA256AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA256AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA384AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA384AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA512AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionCofactorVariableIVX963SHA512AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionCofactorX963SHA1AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionCofactorX963SHA1AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionCofactorX963SHA224AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionCofactorX963SHA224AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionCofactorX963SHA256AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionCofactorX963SHA256AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionCofactorX963SHA384AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionCofactorX963SHA384AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionCofactorX963SHA512AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionCofactorX963SHA512AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA224AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA224AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA256AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA256AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA384AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA384AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA512AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionStandardVariableIVX963SHA512AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionStandardX963SHA1AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionStandardX963SHA1AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionStandardX963SHA224AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionStandardX963SHA224AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionStandardX963SHA256AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionStandardX963SHA256AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionStandardX963SHA384AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionStandardX963SHA384AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmECIESEncryptionStandardX963SHA512AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmECIESEncryptionStandardX963SHA512AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionOAEPSHA1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionOAEPSHA1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionOAEPSHA1AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionOAEPSHA1AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionOAEPSHA224"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionOAEPSHA224 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionOAEPSHA224AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionOAEPSHA224AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionOAEPSHA256"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionOAEPSHA256 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionOAEPSHA256AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionOAEPSHA256AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionOAEPSHA384"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionOAEPSHA384 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionOAEPSHA384AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionOAEPSHA384AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionOAEPSHA512"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionOAEPSHA512 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionOAEPSHA512AESGCM"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionOAEPSHA512AESGCM = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionPKCS1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionPKCS1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSAEncryptionRaw"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSAEncryptionRaw = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureDigestPKCS1v15Raw"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureDigestPKCS1v15Raw = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA224"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA224 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA256"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA256 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA384"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA384 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA512"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureDigestPKCS1v15SHA512 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureDigestPSSSHA1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureDigestPSSSHA1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureDigestPSSSHA224"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureDigestPSSSHA224 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureDigestPSSSHA256"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureDigestPSSSHA256 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureDigestPSSSHA384"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureDigestPSSSHA384 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureDigestPSSSHA512"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureDigestPSSSHA512 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA224"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA224 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA256"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA256 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA384"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA384 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA512"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureMessagePKCS1v15SHA512 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureMessagePSSSHA1"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureMessagePSSSHA1 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureMessagePSSSHA224"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureMessagePSSSHA224 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureMessagePSSSHA256"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureMessagePSSSHA256 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureMessagePSSSHA384"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureMessagePSSSHA384 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureMessagePSSSHA512"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureMessagePSSSHA512 = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAlgorithmRSASignatureRaw"); err == nil && ptr != 0 {
		KSecKeyAlgorithmRSASignatureRaw = *(*SecKeyAlgorithm)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecKeyAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyKeyExchangeParameterRequestedSize"); err == nil && ptr != 0 {
		KSecKeyKeyExchangeParameterRequestedSize = *(*SecKeyKeyExchangeParameter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecKeyKeyExchangeParameterSharedInfo"); err == nil && ptr != 0 {
		KSecKeyKeyExchangeParameterSharedInfo = *(*SecKeyKeyExchangeParameter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchCaseInsensitive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchCaseInsensitive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchDiacriticInsensitive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchDiacriticInsensitive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchEmailAddressIfPresent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchEmailAddressIfPresent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchHostOrSubdomainOfHost"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchHostOrSubdomainOfHost = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchIssuers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchIssuers = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchItemList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchItemList = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchLimit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchLimit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchLimitAll"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchLimitAll = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchLimitOne"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchLimitOne = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchPolicy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchPolicy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchSearchList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchSearchList = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchSubjectContains"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchSubjectContains = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchSubjectEndsWith"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchSubjectEndsWith = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchSubjectStartsWith"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchSubjectStartsWith = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchSubjectWholeString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchSubjectWholeString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchTrustedOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchTrustedOnly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchValidOnDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchValidOnDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecMatchWidthInsensitive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecMatchWidthInsensitive = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDADC_CERT_POLICY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDADC_CERT_POLICY = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_CERT_POLICY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_CERT_POLICY = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EKU_CODE_SIGNING"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EKU_CODE_SIGNING = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EKU_CODE_SIGNING_DEV"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EKU_CODE_SIGNING_DEV = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EKU_ICHAT_ENCRYPTION"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EKU_ICHAT_ENCRYPTION = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EKU_ICHAT_SIGNING"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EKU_ICHAT_SIGNING = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EKU_RESOURCE_SIGNING"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EKU_RESOURCE_SIGNING = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EKU_SYSTEM_IDENTITY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EKU_SYSTEM_IDENTITY = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EXTENSION"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EXTENSION = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EXTENSION_AAI_INTERMEDIATE"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EXTENSION_AAI_INTERMEDIATE = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EXTENSION_ADC_APPLE_SIGNING"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EXTENSION_ADC_APPLE_SIGNING = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EXTENSION_ADC_DEV_SIGNING"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EXTENSION_ADC_DEV_SIGNING = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EXTENSION_APPLEID_INTERMEDIATE"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EXTENSION_APPLEID_INTERMEDIATE = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EXTENSION_APPLE_SIGNING"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EXTENSION_APPLE_SIGNING = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EXTENSION_CODE_SIGNING"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EXTENSION_CODE_SIGNING = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EXTENSION_INTERMEDIATE_MARKER"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EXTENSION_INTERMEDIATE_MARKER = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EXTENSION_ITMS_INTERMEDIATE"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EXTENSION_ITMS_INTERMEDIATE = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAPPLE_EXTENSION_WWDR_INTERMEDIATE"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAPPLE_EXTENSION_WWDR_INTERMEDIATE = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAuthorityInfoAccess"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAuthorityInfoAccess = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDAuthorityKeyIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDAuthorityKeyIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDBasicConstraints"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDBasicConstraints = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDBiometricInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDBiometricInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDCSSMKeyStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDCSSMKeyStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDCertIssuer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDCertIssuer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDCertificatePolicies"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDCertificatePolicies = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDClientAuth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDClientAuth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDCollectiveStateProvinceName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDCollectiveStateProvinceName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDCollectiveStreetAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDCollectiveStreetAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDCommonName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDCommonName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDCountryName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDCountryName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDCrlDistributionPoints"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDCrlDistributionPoints = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDCrlNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDCrlNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDCrlReason"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDCrlReason = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDDOTMAC_CERT_EMAIL_ENCRYPT"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDDOTMAC_CERT_EMAIL_ENCRYPT = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDDOTMAC_CERT_EMAIL_SIGN"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDDOTMAC_CERT_EMAIL_SIGN = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDDOTMAC_CERT_EXTENSION"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDDOTMAC_CERT_EXTENSION = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDDOTMAC_CERT_IDENTITY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDDOTMAC_CERT_IDENTITY = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDDOTMAC_CERT_POLICY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDDOTMAC_CERT_POLICY = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDDeltaCrlIndicator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDDeltaCrlIndicator = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDEKU_IPSec"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDEKU_IPSec = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDEmailAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDEmailAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDEmailProtection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDEmailProtection = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDExtendedKeyUsage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDExtendedKeyUsage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDExtendedKeyUsageAny"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDExtendedKeyUsageAny = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDExtendedUseCodeSigning"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDExtendedUseCodeSigning = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDGivenName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDGivenName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDHoldInstructionCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDHoldInstructionCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDInvalidityDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDInvalidityDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDIssuerAltName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDIssuerAltName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDIssuingDistributionPoint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDIssuingDistributionPoint = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDIssuingDistributionPoints"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDIssuingDistributionPoints = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDKERBv5_PKINIT_KP_CLIENT_AUTH"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDKERBv5_PKINIT_KP_CLIENT_AUTH = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDKERBv5_PKINIT_KP_KDC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDKERBv5_PKINIT_KP_KDC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDKeyUsage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDKeyUsage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDLocalityName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDLocalityName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDMS_NTPrincipalName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDMS_NTPrincipalName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDMicrosoftSGC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDMicrosoftSGC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDNameConstraints"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDNameConstraints = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDNetscapeCertSequence"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDNetscapeCertSequence = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDNetscapeCertType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDNetscapeCertType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDNetscapeSGC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDNetscapeSGC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDOCSPSigning"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDOCSPSigning = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDOrganizationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDOrganizationName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDOrganizationalUnitName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDOrganizationalUnitName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDPolicyConstraints"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDPolicyConstraints = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDPolicyMappings"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDPolicyMappings = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDPrivateKeyUsagePeriod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDPrivateKeyUsagePeriod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDQC_Statements"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDQC_Statements = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDSRVName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDSRVName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDSerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDSerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDServerAuth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDServerAuth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDStateProvinceName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDStateProvinceName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDStreetAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDStreetAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDSubjectAltName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDSubjectAltName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDSubjectDirectoryAttributes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDSubjectDirectoryAttributes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDSubjectEmailAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDSubjectEmailAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDSubjectInfoAccess"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDSubjectInfoAccess = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDSubjectKeyIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDSubjectKeyIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDSubjectPicture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDSubjectPicture = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDSubjectSignatureBitmap"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDSubjectSignatureBitmap = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDSurname"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDSurname = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDTimeStamping"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDTimeStamping = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDTitle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDUseExemptions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDUseExemptions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1CertificateIssuerUniqueId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1CertificateIssuerUniqueId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1CertificateSubjectUniqueId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1CertificateSubjectUniqueId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1IssuerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1IssuerName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1IssuerNameCStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1IssuerNameCStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1IssuerNameLDAP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1IssuerNameLDAP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1IssuerNameStd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1IssuerNameStd = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SerialNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SerialNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1Signature"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1Signature = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SignatureAlgorithm"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SignatureAlgorithm = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SignatureAlgorithmParameters"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SignatureAlgorithmParameters = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SignatureAlgorithmTBS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SignatureAlgorithmTBS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SignatureCStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SignatureCStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SignatureStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SignatureStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SubjectName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SubjectName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SubjectNameCStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SubjectNameCStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SubjectNameLDAP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SubjectNameLDAP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SubjectNameStd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SubjectNameStd = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SubjectPublicKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SubjectPublicKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SubjectPublicKeyAlgorithm"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SubjectPublicKeyAlgorithm = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SubjectPublicKeyAlgorithmParameters"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SubjectPublicKeyAlgorithmParameters = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1SubjectPublicKeyCStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1SubjectPublicKeyCStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1ValidityNotAfter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1ValidityNotAfter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1ValidityNotBefore"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1ValidityNotBefore = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V1Version"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V1Version = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3Certificate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3Certificate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3CertificateCStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3CertificateCStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3CertificateExtensionCStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3CertificateExtensionCStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3CertificateExtensionCritical"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3CertificateExtensionCritical = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3CertificateExtensionId"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3CertificateExtensionId = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3CertificateExtensionStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3CertificateExtensionStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3CertificateExtensionType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3CertificateExtensionType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3CertificateExtensionValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3CertificateExtensionValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3CertificateExtensionsCStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3CertificateExtensionsCStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3CertificateExtensionsStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3CertificateExtensionsStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3CertificateNumberOfExtensions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3CertificateNumberOfExtensions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3SignedCertificate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3SignedCertificate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecOIDX509V3SignedCertificateCStruct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecOIDX509V3SignedCertificateCStruct = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleCodeSigning"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleCodeSigning = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleEAP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleEAP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleEAPClient"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleEAPClient = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleEAPServer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleEAPServer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleIDValidation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleIDValidation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleIPSecClient"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleIPSecClient = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleIPSecServer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleIPSecServer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleIPsec"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleIPsec = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyApplePKINITClient"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyApplePKINITClient = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyApplePKINITServer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyApplePKINITServer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyApplePassbookSigning"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyApplePassbookSigning = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyApplePayIssuerEncryption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyApplePayIssuerEncryption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleRevocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleRevocation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleSMIME"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleSMIME = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleSSL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleSSL = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleSSLClient"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleSSLClient = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleSSLServer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleSSLServer = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleTimeStamping"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleTimeStamping = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyAppleX509Basic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyAppleX509Basic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyClient"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyClient = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyKU_CRLSign"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyKU_CRLSign = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyKU_DataEncipherment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyKU_DataEncipherment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyKU_DecipherOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyKU_DecipherOnly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyKU_DigitalSignature"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyKU_DigitalSignature = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyKU_EncipherOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyKU_EncipherOnly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyKU_KeyAgreement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyKU_KeyAgreement = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyKU_KeyCertSign"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyKU_KeyCertSign = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyKU_KeyEncipherment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyKU_KeyEncipherment = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyKU_NonRepudiation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyKU_NonRepudiation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyMacAppStoreReceipt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyMacAppStoreReceipt = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyOid"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyOid = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyRevocationFlags"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyRevocationFlags = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPolicyTeamIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPolicyTeamIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPrivateKeyAttrs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPrivateKeyAttrs = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyKeyLabel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyKeyLabel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyKeyLocalizedLabel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyKeyLocalizedLabel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyKeyType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyKeyType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyKeyValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyKeyValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyTypeArray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyTypeArray = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyTypeData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyTypeData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyTypeDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyTypeDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyTypeError"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyTypeError = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyTypeNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyTypeNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyTypeSection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyTypeSection = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyTypeString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyTypeString = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyTypeSuccess"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyTypeSuccess = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyTypeTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyTypeTitle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyTypeURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyTypeURL = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPropertyTypeWarning"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPropertyTypeWarning = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecPublicKeyAttrs"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecPublicKeyAttrs = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecRandomDefault"); err == nil && ptr != 0 {
		KSecRandomDefault = *(*SecRandomRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecReturnAttributes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecReturnAttributes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecReturnData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecReturnData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecReturnPersistentRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecReturnPersistentRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecReturnRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecReturnRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecSharedPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecSharedPassword = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecSignatureAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecSignatureAttributeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTransformAbortOriginatorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTransformAbortOriginatorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTransformErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTransformErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTransformPreviousErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTransformPreviousErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTrustCertificateTransparency"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTrustCertificateTransparency = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTrustEvaluationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTrustEvaluationDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTrustExtendedValidation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTrustExtendedValidation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTrustOrganizationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTrustOrganizationName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTrustQCStatements"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTrustQCStatements = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTrustQWACValidation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTrustQWACValidation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTrustResultValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTrustResultValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTrustRevocationChecked"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTrustRevocationChecked = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecTrustRevocationValidUntilDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecTrustRevocationValidUntilDate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecUseAuthenticationContext"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecUseAuthenticationContext = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecUseAuthenticationUI"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecUseAuthenticationUI = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecUseAuthenticationUISkip"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecUseAuthenticationUISkip = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecUseDataProtectionKeychain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecUseDataProtectionKeychain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecUseKeychain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecUseKeychain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecValueData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecValueData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecValuePersistentRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecValuePersistentRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kSecValueRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KSecValueRef = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidAdCAIssuer"); err == nil && ptr != 0 {
		oidAdCAIssuer = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidAdOCSP"); err == nil && ptr != 0 {
		oidAdOCSP = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidAnsip384r1"); err == nil && ptr != 0 {
		oidAnsip384r1 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidAnsip521r1"); err == nil && ptr != 0 {
		oidAnsip521r1 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidAnyExtendedKeyUsage"); err == nil && ptr != 0 {
		oidAnyExtendedKeyUsage = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidAnyPolicy"); err == nil && ptr != 0 {
		oidAnyPolicy = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidAuthorityInfoAccess"); err == nil && ptr != 0 {
		oidAuthorityInfoAccess = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidAuthorityKeyIdentifier"); err == nil && ptr != 0 {
		oidAuthorityKeyIdentifier = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidBasicConstraints"); err == nil && ptr != 0 {
		oidBasicConstraints = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidCertificatePolicies"); err == nil && ptr != 0 {
		oidCertificatePolicies = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidCommonName"); err == nil && ptr != 0 {
		oidCommonName = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidCountryName"); err == nil && ptr != 0 {
		oidCountryName = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidCrlDistributionPoints"); err == nil && ptr != 0 {
		oidCrlDistributionPoints = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidDescription"); err == nil && ptr != 0 {
		oidDescription = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidEcPrime192v1"); err == nil && ptr != 0 {
		oidEcPrime192v1 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidEcPrime256v1"); err == nil && ptr != 0 {
		oidEcPrime256v1 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidEcPubKey"); err == nil && ptr != 0 {
		oidEcPubKey = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidEmailAddress"); err == nil && ptr != 0 {
		oidEmailAddress = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidEntrustVersInfo"); err == nil && ptr != 0 {
		oidEntrustVersInfo = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidExtendedKeyUsage"); err == nil && ptr != 0 {
		oidExtendedKeyUsage = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidExtendedKeyUsageClientAuth"); err == nil && ptr != 0 {
		oidExtendedKeyUsageClientAuth = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidExtendedKeyUsageCodeSigning"); err == nil && ptr != 0 {
		oidExtendedKeyUsageCodeSigning = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidExtendedKeyUsageEmailProtection"); err == nil && ptr != 0 {
		oidExtendedKeyUsageEmailProtection = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidExtendedKeyUsageIPSec"); err == nil && ptr != 0 {
		oidExtendedKeyUsageIPSec = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidExtendedKeyUsageMicrosoftSGC"); err == nil && ptr != 0 {
		oidExtendedKeyUsageMicrosoftSGC = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidExtendedKeyUsageNetscapeSGC"); err == nil && ptr != 0 {
		oidExtendedKeyUsageNetscapeSGC = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidExtendedKeyUsageOCSPSigning"); err == nil && ptr != 0 {
		oidExtendedKeyUsageOCSPSigning = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidExtendedKeyUsageServerAuth"); err == nil && ptr != 0 {
		oidExtendedKeyUsageServerAuth = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidExtendedKeyUsageTimeStamping"); err == nil && ptr != 0 {
		oidExtendedKeyUsageTimeStamping = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidFee"); err == nil && ptr != 0 {
		oidFee = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidFriendlyName"); err == nil && ptr != 0 {
		oidFriendlyName = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidGoogleEmbeddedSignedCertificateTimestamp"); err == nil && ptr != 0 {
		oidGoogleEmbeddedSignedCertificateTimestamp = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidGoogleOCSPSignedCertificateTimestamp"); err == nil && ptr != 0 {
		oidGoogleOCSPSignedCertificateTimestamp = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidInhibitAnyPolicy"); err == nil && ptr != 0 {
		oidInhibitAnyPolicy = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidIssuerAltName"); err == nil && ptr != 0 {
		oidIssuerAltName = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidKeyUsage"); err == nil && ptr != 0 {
		oidKeyUsage = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidLocalKeyId"); err == nil && ptr != 0 {
		oidLocalKeyId = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidLocalityName"); err == nil && ptr != 0 {
		oidLocalityName = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidMSNTPrincipalName"); err == nil && ptr != 0 {
		oidMSNTPrincipalName = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidMd2"); err == nil && ptr != 0 {
		oidMd2 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidMd2Rsa"); err == nil && ptr != 0 {
		oidMd2Rsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidMd4"); err == nil && ptr != 0 {
		oidMd4 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidMd4Rsa"); err == nil && ptr != 0 {
		oidMd4Rsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidMd5"); err == nil && ptr != 0 {
		oidMd5 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidMd5Fee"); err == nil && ptr != 0 {
		oidMd5Fee = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidMd5Rsa"); err == nil && ptr != 0 {
		oidMd5Rsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidNameConstraints"); err == nil && ptr != 0 {
		oidNameConstraints = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidNetscapeCertType"); err == nil && ptr != 0 {
		oidNetscapeCertType = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidOrganizationName"); err == nil && ptr != 0 {
		oidOrganizationName = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidOrganizationalUnitName"); err == nil && ptr != 0 {
		oidOrganizationalUnitName = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidPolicyConstraints"); err == nil && ptr != 0 {
		oidPolicyConstraints = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidPolicyMappings"); err == nil && ptr != 0 {
		oidPolicyMappings = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidPrivateKeyUsagePeriod"); err == nil && ptr != 0 {
		oidPrivateKeyUsagePeriod = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidQtCps"); err == nil && ptr != 0 {
		oidQtCps = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidQtUNotice"); err == nil && ptr != 0 {
		oidQtUNotice = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidRsa"); err == nil && ptr != 0 {
		oidRsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha1"); err == nil && ptr != 0 {
		oidSha1 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha1Dsa"); err == nil && ptr != 0 {
		oidSha1Dsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha1DsaCommonOIW"); err == nil && ptr != 0 {
		oidSha1DsaCommonOIW = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha1DsaOIW"); err == nil && ptr != 0 {
		oidSha1DsaOIW = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha1Ecdsa"); err == nil && ptr != 0 {
		oidSha1Ecdsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha1Fee"); err == nil && ptr != 0 {
		oidSha1Fee = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha1Rsa"); err == nil && ptr != 0 {
		oidSha1Rsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha1RsaOIW"); err == nil && ptr != 0 {
		oidSha1RsaOIW = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha224"); err == nil && ptr != 0 {
		oidSha224 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha224Ecdsa"); err == nil && ptr != 0 {
		oidSha224Ecdsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha224Rsa"); err == nil && ptr != 0 {
		oidSha224Rsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha256"); err == nil && ptr != 0 {
		oidSha256 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha256Ecdsa"); err == nil && ptr != 0 {
		oidSha256Ecdsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha256Rsa"); err == nil && ptr != 0 {
		oidSha256Rsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha384"); err == nil && ptr != 0 {
		oidSha384 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha384Ecdsa"); err == nil && ptr != 0 {
		oidSha384Ecdsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha384Rsa"); err == nil && ptr != 0 {
		oidSha384Rsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha512"); err == nil && ptr != 0 {
		oidSha512 = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha512Ecdsa"); err == nil && ptr != 0 {
		oidSha512Ecdsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSha512Rsa"); err == nil && ptr != 0 {
		oidSha512Rsa = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidStateOrProvinceName"); err == nil && ptr != 0 {
		oidStateOrProvinceName = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSubjectAltName"); err == nil && ptr != 0 {
		oidSubjectAltName = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSubjectInfoAccess"); err == nil && ptr != 0 {
		oidSubjectInfoAccess = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "oidSubjectKeyIdentifier"); err == nil && ptr != 0 {
		oidSubjectKeyIdentifier = *(*unsafe.Pointer)(unsafe.Pointer(ptr))
	}

}
