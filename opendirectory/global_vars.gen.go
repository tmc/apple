// Code generated from Apple documentation. DO NOT EDIT.

package opendirectory

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// See: https://developer.apple.com/documentation/OpenDirectory/ODFrameworkErrorDomain
	ODFrameworkErrorDomain string
	// ODSessionProxyAddress is the address to connect to via proxy. The value is of type [NSString].
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/ODSessionProxyAddress
	ODSessionProxyAddress string
	// ODSessionProxyPassword is the password to connect with via proxy. The value is of type [NSString].
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/ODSessionProxyPassword
	ODSessionProxyPassword string
	// ODSessionProxyPort is the port to connect to via proxy. The value is of type [NSNumber].
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/ODSessionProxyPort
	ODSessionProxyPort string
	// ODSessionProxyUsername is the username to connect with via proxy. The value is of type [NSString].
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/ODSessionProxyUsername
	ODSessionProxyUsername string
	// See: https://developer.apple.com/documentation/OpenDirectory/ODTrustTypeAnonymous
	ODTrustTypeAnonymous string
	// See: https://developer.apple.com/documentation/OpenDirectory/ODTrustTypeJoined
	ODTrustTypeJoined string
	// See: https://developer.apple.com/documentation/OpenDirectory/ODTrustTypeUsingCredentials
	ODTrustTypeUsingCredentials string
)

var (
	// KODAttributeTypeAccessControlEntry is the attribute type used to store directory access control directives.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAccessControlEntry
	KODAttributeTypeAccessControlEntry unsafe.Pointer
	// KODAttributeTypeAddressLine1 is the attribute type used to store the first line of a user’s address data.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAddressLine1
	KODAttributeTypeAddressLine1 unsafe.Pointer
	// KODAttributeTypeAddressLine2 is the attribute type used to store the second line of a user’s address data.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAddressLine2
	KODAttributeTypeAddressLine2 unsafe.Pointer
	// KODAttributeTypeAddressLine3 is the attribute type used to store the third line of a user’s address data.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAddressLine3
	KODAttributeTypeAddressLine3 unsafe.Pointer
	// KODAttributeTypeAdminLimits is the attribute type of an XML property list that indicates what an admin user can edit. Found in records of type `kODRecordTypeUsers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAdminLimits
	KODAttributeTypeAdminLimits unsafe.Pointer
	// KODAttributeTypeAdvertisedServices is the attribute type used to specify (Bonjour) advertised services.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAdvertisedServices
	KODAttributeTypeAdvertisedServices unsafe.Pointer
	// KODAttributeTypeAlias is the attribute type used to specify an alias that contains a pointer to another record, node, or attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAlias
	KODAttributeTypeAlias unsafe.Pointer
	// KODAttributeTypeAllAttributes is the attribute type used in requesting all attribute types in a search.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAllAttributes
	KODAttributeTypeAllAttributes unsafe.Pointer
	// KODAttributeTypeAllTypes is the attribute type used to indicate all attribute types for a given record type the config node.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAllTypes
	KODAttributeTypeAllTypes unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAltSecurityIdentities
	KODAttributeTypeAltSecurityIdentities unsafe.Pointer
	// KODAttributeTypeAreaCode is the attribute type used to store a user’s area code.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAreaCode
	KODAttributeTypeAreaCode unsafe.Pointer
	// KODAttributeTypeAttrListRefCount is the attribute type used to specify the total count of attribute list references for a process.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAttrListRefCount
	KODAttributeTypeAttrListRefCount unsafe.Pointer
	// KODAttributeTypeAttrListRefs is the attribute type used to store the attribute list references for a process.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAttrListRefs
	KODAttributeTypeAttrListRefs unsafe.Pointer
	// KODAttributeTypeAttrListValueRefCount is the attribute type used to specify the total count of attribute list value references for a process.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAttrListValueRefCount
	KODAttributeTypeAttrListValueRefCount unsafe.Pointer
	// KODAttributeTypeAttrListValueRefs is the attribute type used to store the attribute list value references for a process.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAttrListValueRefs
	KODAttributeTypeAttrListValueRefs unsafe.Pointer
	// KODAttributeTypeAuthCredential is the attribute type used to store an authentication credential used to authenticate to a directory.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAuthCredential
	KODAttributeTypeAuthCredential unsafe.Pointer
	// KODAttributeTypeAuthMethod is the attribute type used to specify a record’s authentication method.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAuthMethod
	KODAttributeTypeAuthMethod unsafe.Pointer
	// KODAttributeTypeAuthenticationAuthority is the attribute type used to specify the mechanism used to verify or set a user’s password. Typically found in records of type `kODRecordTypeUsers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAuthenticationAuthority
	KODAttributeTypeAuthenticationAuthority unsafe.Pointer
	// KODAttributeTypeAuthenticationHint is the attribute type of an authentication hint attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAuthenticationHint
	KODAttributeTypeAuthenticationHint unsafe.Pointer
	// KODAttributeTypeAuthorityRevocationList is the attribute type of the authority revocation list attribute, which defines certificate authority certificates that are no longer trusted. Typically found in records of type `kODRecordTypeCertificateAuthorities`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAuthorityRevocationList
	KODAttributeTypeAuthorityRevocationList unsafe.Pointer
	// KODAttributeTypeAutomaticSearchPath is the attribute type used to specify the automatic search path used by a search node.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAutomaticSearchPath
	KODAttributeTypeAutomaticSearchPath unsafe.Pointer
	// KODAttributeTypeAutomountInformation is the attribute type used to store automount information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeAutomountInformation
	KODAttributeTypeAutomountInformation unsafe.Pointer
	// KODAttributeTypeBirthday is the attribute type of a birthday attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeBirthday
	KODAttributeTypeBirthday unsafe.Pointer
	// KODAttributeTypeBootParams is the attribute type used to store boot parameters. Typically found in host or machine records.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeBootParams
	KODAttributeTypeBootParams unsafe.Pointer
	// KODAttributeTypeBuildVersion is the attribute type used to specify the build version.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeBuildVersion
	KODAttributeTypeBuildVersion unsafe.Pointer
	// KODAttributeTypeBuilding is the attribute type used to store a user’s building information. Typically found in records of type `kODRecordTypeUsers` or `kODRecordTypePeople`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeBuilding
	KODAttributeTypeBuilding unsafe.Pointer
	// KODAttributeTypeCACertificate is the attribute type of a certificate authority certificate attribute, which contains the binary of the certificate. Typically found in records of type `kODRecordTypeCertificateAuthority`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeCACertificate
	KODAttributeTypeCACertificate unsafe.Pointer
	// KODAttributeTypeCapacity is the attribute type of a capacity attribute, which indicates the capacity of a resource.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeCapacity
	KODAttributeTypeCapacity unsafe.Pointer
	// KODAttributeTypeCertificateRevocationList is the attribute type of the certificate revocation list attribute, which defines certificates that are no longer trusted. Typically found in records of type `kODRecordTypeCertificateAuthorities`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeCertificateRevocationList
	KODAttributeTypeCertificateRevocationList unsafe.Pointer
	// KODAttributeTypeCity is the attribute type used to store a user’s city information. Typically found in records of type `kODRecordTypeUsers` or `kODRecordTypePeople`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeCity
	KODAttributeTypeCity unsafe.Pointer
	// KODAttributeTypeComment is the attribute type of an unformatted comment attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeComment
	KODAttributeTypeComment unsafe.Pointer
	// KODAttributeTypeCompany is the attribute type used to store a user’s company information. Typically found in records of type `kODRecordTypeUsers` or `kODRecordTypePeople`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeCompany
	KODAttributeTypeCompany unsafe.Pointer
	// KODAttributeTypeComputers is the attribute type used to store a list of computers.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeComputers
	KODAttributeTypeComputers unsafe.Pointer
	// KODAttributeTypeConfigAvailable is the attribute type for the config available flag.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeConfigAvailable
	KODAttributeTypeConfigAvailable unsafe.Pointer
	// KODAttributeTypeConfigFile is the attribute type used to specify the config file name.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeConfigFile
	KODAttributeTypeConfigFile unsafe.Pointer
	// KODAttributeTypeContactGUID is the attribute type of the contact GUID attribute. Typically found in records of type `kODRecordTypeGroups`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeContactGUID
	KODAttributeTypeContactGUID unsafe.Pointer
	// KODAttributeTypeContactPerson is the attribute type of the contact person attribute, which indicates the contact person of a machine.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeContactPerson
	KODAttributeTypeContactPerson unsafe.Pointer
	// KODAttributeTypeCopyTimestamp is the attribute type used to store a timestamp used in local account caching.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeCopyTimestamp
	KODAttributeTypeCopyTimestamp unsafe.Pointer
	// KODAttributeTypeCoreFWVersion is the attribute type used to specify the version of the core framework.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeCoreFWVersion
	KODAttributeTypeCoreFWVersion unsafe.Pointer
	// KODAttributeTypeCountry is the attribute type used to store a user’s country or region information. Typically found in records of type `kODRecordTypeUsers` or `kODRecordTypePeople`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeCountry
	KODAttributeTypeCountry unsafe.Pointer
	// KODAttributeTypeCreationTimestamp is the attribute type of the creation timestamp attribute, which indicates the time the record was created.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeCreationTimestamp
	KODAttributeTypeCreationTimestamp unsafe.Pointer
	// KODAttributeTypeCrossCertificatePair is the attribute type of the cross certificate attribute, which contains the binary of two certificates that verify one another. Typically found in records of type `kODRecordTypeCertificateAuthorities`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeCrossCertificatePair
	KODAttributeTypeCrossCertificatePair unsafe.Pointer
	// KODAttributeTypeCustomSearchPath is the attribute type used to specify an admin-configured search path used by a search node.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeCustomSearchPath
	KODAttributeTypeCustomSearchPath unsafe.Pointer
	// KODAttributeTypeDNSDomain is the attribute type of the DNS domain attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeDNSDomain
	KODAttributeTypeDNSDomain unsafe.Pointer
	// KODAttributeTypeDNSName is the attribute type used to specify a DNS resolver nameserver.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeDNSName
	KODAttributeTypeDNSName unsafe.Pointer
	// KODAttributeTypeDNSNameServer is the attribute type of the DNS name server attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeDNSNameServer
	KODAttributeTypeDNSNameServer unsafe.Pointer
	// KODAttributeTypeDataStamp is the attribute type of the data stamp attribute, which is used for checksum and accessing metadata.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeDataStamp
	KODAttributeTypeDataStamp unsafe.Pointer
	// KODAttributeTypeDateRecordCreated is the attribute type used to store a record’s creation date.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeDateRecordCreated
	KODAttributeTypeDateRecordCreated unsafe.Pointer
	// KODAttributeTypeDepartment is the attribute type used to store a user’s department information. Typically found in records of type `kODRecordTypeUsers` or `kODRecordTypePeople`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeDepartment
	KODAttributeTypeDepartment unsafe.Pointer
	// KODAttributeTypeDirRefCount is the attribute type used to specify the total count of directory references for a process.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeDirRefCount
	KODAttributeTypeDirRefCount unsafe.Pointer
	// KODAttributeTypeDirRefs is the attribute type used to store the directory references for a process.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeDirRefs
	KODAttributeTypeDirRefs unsafe.Pointer
	// KODAttributeTypeEMailAddress is the attribute type used to store a user’s email address. Typically found in records of type `kODRecordTypeUsers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeEMailAddress
	KODAttributeTypeEMailAddress unsafe.Pointer
	// KODAttributeTypeEMailContacts is the attribute type used to store a user’s custom email information. Typically found in records of type `kODRecordTypeUsers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeEMailContacts
	KODAttributeTypeEMailContacts unsafe.Pointer
	// KODAttributeTypeENetAddress is the attribute type of the ethernet address attribute, which specifies a record’s ethernet address (MAC address).
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeENetAddress
	KODAttributeTypeENetAddress unsafe.Pointer
	// KODAttributeTypeExpire is the attribute type of the expiration attribute, which indicates the expiration date or time of a record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeExpire
	KODAttributeTypeExpire unsafe.Pointer
	// KODAttributeTypeFWVersion is the attribute type used to specify the version of the framework.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeFWVersion
	KODAttributeTypeFWVersion unsafe.Pointer
	// KODAttributeTypeFaxNumber is the attribute type used to store a user’s fax number. Typically found in records of type `kODRecordTypeUsers` or `kODRecordTypePeople`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeFaxNumber
	KODAttributeTypeFaxNumber unsafe.Pointer
	// KODAttributeTypeFirstName is the attribute type of the first name attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeFirstName
	KODAttributeTypeFirstName unsafe.Pointer
	// KODAttributeTypeFullName is the attribute type of the full name attribute, which indicates the full name of the record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeFullName
	KODAttributeTypeFullName unsafe.Pointer
	// KODAttributeTypeFunctionalState is the attribute type used to specify the functional state of the plug-in.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeFunctionalState
	KODAttributeTypeFunctionalState unsafe.Pointer
	// KODAttributeTypeGUID is the attribute type of the GUID attribute, which indicates a record’s 128-bit GUID.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeGUID
	KODAttributeTypeGUID unsafe.Pointer
	// KODAttributeTypeGroup is the attribute type used to store a list of groups.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeGroup
	KODAttributeTypeGroup unsafe.Pointer
	// KODAttributeTypeGroupMembers is the attribute type used to specify the GUID values of members of a group that are not groups.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeGroupMembers
	KODAttributeTypeGroupMembers unsafe.Pointer
	// KODAttributeTypeGroupMembership is the attribute type used to specify a list of users that belong to a given group.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeGroupMembership
	KODAttributeTypeGroupMembership unsafe.Pointer
	// KODAttributeTypeGroupServices is the attribute type used to specify an XML property list that defines a group’s services. Typically found in records of type `kODRecordTypeGroups`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeGroupServices
	KODAttributeTypeGroupServices unsafe.Pointer
	// KODAttributeTypeHTML is the attribute type used to specify an HTML location.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeHTML
	KODAttributeTypeHTML unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeHardwareUUID
	KODAttributeTypeHardwareUUID unsafe.Pointer
	// KODAttributeTypeHomeDirectory is the attribute type used to specify the allowed usage of a user’s home directory in bytes. Typically found in records of type `kODRecordTypeUsers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeHomeDirectory
	KODAttributeTypeHomeDirectory unsafe.Pointer
	// KODAttributeTypeHomeDirectoryQuota is the attribute type of the home directory quota attribute, which is listed in bytes. Typically found in records of type `kODRecordTypeUsers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeHomeDirectoryQuota
	KODAttributeTypeHomeDirectoryQuota unsafe.Pointer
	// KODAttributeTypeHomeDirectorySoftQuota is the attribute type of the home directory soft quota attribute, which is listed in bytes. This specifies a size limit at which users are notified that they are reaching their hard quota. Typically found in records of type `kODRecordTypeUsers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeHomeDirectorySoftQuota
	KODAttributeTypeHomeDirectorySoftQuota unsafe.Pointer
	// KODAttributeTypeHomeLocOwner is the attribute type of the workgroup shared home directory attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeHomeLocOwner
	KODAttributeTypeHomeLocOwner unsafe.Pointer
	// KODAttributeTypeHomePhoneNumber is the attribute type used to store a user’s home phone number. Typically found in records of type `kODRecordTypeUsers` or `kODRecordTypePeople`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeHomePhoneNumber
	KODAttributeTypeHomePhoneNumber unsafe.Pointer
	// KODAttributeTypeIMHandle is the attribute type used to store a user’s instant messaging handles. Typically found in records of type `kODRecordTypeUsers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeIMHandle
	KODAttributeTypeIMHandle unsafe.Pointer
	// KODAttributeTypeIPAddress is the attribute type used to specify an IP address.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeIPAddress
	KODAttributeTypeIPAddress unsafe.Pointer
	// KODAttributeTypeIPAddressAndENetAddress is the attribute type used to specify a pairing of an IPv4 or IPv6 address with an ethernet address. Typically found in records of type `kODRecordTypeComputers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeIPAddressAndENetAddress
	KODAttributeTypeIPAddressAndENetAddress unsafe.Pointer
	// KODAttributeTypeIPv6Address is the attribute type used to specify an IPv6 address. Typically found in records of type `kODRecordTypeComputers` and `kODRecordTypeHosts`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeIPv6Address
	KODAttributeTypeIPv6Address unsafe.Pointer
	// KODAttributeTypeInternetAlias is the attribute type of the internet alias attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeInternetAlias
	KODAttributeTypeInternetAlias unsafe.Pointer
	// KODAttributeTypeJPEGPhoto is the attribute type used to store binary picture data in JPEG format. Typically found in records of type `kODRecordTypeUsers`, `kODRecordTypePeople`, and `kODRecordTypeGroups`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeJPEGPhoto
	KODAttributeTypeJPEGPhoto unsafe.Pointer
	// KODAttributeTypeJobTitle is the attribute type used to store a user’s job title.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeJobTitle
	KODAttributeTypeJobTitle unsafe.Pointer
	// KODAttributeTypeKDCAuthKey is the attribute type used to store a KDC primary key.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeKDCAuthKey
	KODAttributeTypeKDCAuthKey unsafe.Pointer
	// KODAttributeTypeKDCConfigData is the attribute type of the KDC configuration file attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeKDCConfigData
	KODAttributeTypeKDCConfigData unsafe.Pointer
	// KODAttributeTypeKerberosRealm is an attribute type used to support Kerberos SMB server services.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeKerberosRealm
	KODAttributeTypeKerberosRealm unsafe.Pointer
	// KODAttributeTypeKerberosServices is the attribute type of the Kerberos services attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeKerberosServices
	KODAttributeTypeKerberosServices unsafe.Pointer
	// KODAttributeTypeKeywords is the attribute type used to specify keywords for search capability.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeKeywords
	KODAttributeTypeKeywords unsafe.Pointer
	// KODAttributeTypeLDAPReadReplicas is the attribute type used to specify a list of LDAP server URLs that can be used to read directory data.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeLDAPReadReplicas
	KODAttributeTypeLDAPReadReplicas unsafe.Pointer
	// KODAttributeTypeLDAPSearchBaseSuffix is the attribute type of the LDAP server search base suffix attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeLDAPSearchBaseSuffix
	KODAttributeTypeLDAPSearchBaseSuffix unsafe.Pointer
	// KODAttributeTypeLDAPWriteReplicas is the attribute type used to specify a list of LDAP server URLs that can be used to write directory data.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeLDAPWriteReplicas
	KODAttributeTypeLDAPWriteReplicas unsafe.Pointer
	// KODAttributeTypeLastName is the attribute type of the last name attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeLastName
	KODAttributeTypeLastName unsafe.Pointer
	// KODAttributeTypeLocalOnlySearchPath is the attribute type used to specify the local-only search path used by a search node.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeLocalOnlySearchPath
	KODAttributeTypeLocalOnlySearchPath unsafe.Pointer
	// KODAttributeTypeLocaleRelay is the attribute type used to specify a relay server for a locale.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeLocaleRelay
	KODAttributeTypeLocaleRelay unsafe.Pointer
	// KODAttributeTypeLocaleSubnets is the attribute type used to specify the subnets for a locale.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeLocaleSubnets
	KODAttributeTypeLocaleSubnets unsafe.Pointer
	// KODAttributeTypeLocation is the attribute type of the location attribute, which indicates the domain names a service is available from. Typically found in service record types, such as `kODRecordTypeAFPServer`, `kODRecordTypeLDAPServer`, and `kODRecordTypeWebServer`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeLocation
	KODAttributeTypeLocation unsafe.Pointer
	// KODAttributeTypeMCXFlags is the attribute type of the MCX flags attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMCXFlags
	KODAttributeTypeMCXFlags unsafe.Pointer
	// KODAttributeTypeMCXSettings is the attribute type of the MCX settings attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMCXSettings
	KODAttributeTypeMCXSettings unsafe.Pointer
	// KODAttributeTypeMIME is the attribute type used to store data of a fully qualified MIME type.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMIME
	KODAttributeTypeMIME unsafe.Pointer
	// KODAttributeTypeMailAttribute is the attribute type of the mail attribute, which contains mail account configuration data.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMailAttribute
	KODAttributeTypeMailAttribute unsafe.Pointer
	// KODAttributeTypeMapCoordinates is the attribute type used to store coordinates of a user’s location.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMapCoordinates
	KODAttributeTypeMapCoordinates unsafe.Pointer
	// KODAttributeTypeMapGUID is the attribute type of the map GUID attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMapGUID
	KODAttributeTypeMapGUID unsafe.Pointer
	// KODAttributeTypeMapURI is the attribute type used to specify the URI of a user’s location.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMapURI
	KODAttributeTypeMapURI unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMetaAmbiguousName
	KODAttributeTypeMetaAmbiguousName unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMetaAugmentedAttributes
	KODAttributeTypeMetaAugmentedAttributes unsafe.Pointer
	// KODAttributeTypeMetaAutomountMap is the attribute type used to query for records of type `kODRecordTypeAutomount` that are associated with a specific record of type `kODRecordTypeAutomountMap`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMetaAutomountMap
	KODAttributeTypeMetaAutomountMap unsafe.Pointer
	// KODAttributeTypeMetaNodeLocation is the attribute type used to retrieve the registered node name with the directory node plug-in.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMetaNodeLocation
	KODAttributeTypeMetaNodeLocation unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMetaRecordName
	KODAttributeTypeMetaRecordName unsafe.Pointer
	// KODAttributeTypeMiddleName is the attribute type of the middle name attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMiddleName
	KODAttributeTypeMiddleName unsafe.Pointer
	// KODAttributeTypeMobileNumber is the attribute type used to store a user’s mobile phone information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeMobileNumber
	KODAttributeTypeMobileNumber unsafe.Pointer
	// KODAttributeTypeModificationTimestamp is the attribute type of the modification timestamp attribute, which indicates the time the record was last modified.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeModificationTimestamp
	KODAttributeTypeModificationTimestamp unsafe.Pointer
	// KODAttributeTypeNFSHomeDirectory is the attribute type of the NFS home directory attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNFSHomeDirectory
	KODAttributeTypeNFSHomeDirectory unsafe.Pointer
	// KODAttributeTypeNTDomainComputerAccount is an attribute type used to support Kerberos SMB server services.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNTDomainComputerAccount
	KODAttributeTypeNTDomainComputerAccount unsafe.Pointer
	// KODAttributeTypeNamePrefix is the attribute type used to store a user’s title prefix.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNamePrefix
	KODAttributeTypeNamePrefix unsafe.Pointer
	// KODAttributeTypeNameSuffix is the attribute type used to specify a user’s title suffix.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNameSuffix
	KODAttributeTypeNameSuffix unsafe.Pointer
	// KODAttributeTypeNativeOnly is the attribute type used in requesting only native attribute types in a search.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNativeOnly
	KODAttributeTypeNativeOnly unsafe.Pointer
	// KODAttributeTypeNestedGroups is the attribute type used to specify a list of nested group GUID values in a group attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNestedGroups
	KODAttributeTypeNestedGroups unsafe.Pointer
	// KODAttributeTypeNetGroupTriplet is the attribute type used to specify a node’s list of subnodes.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNetGroupTriplet
	KODAttributeTypeNetGroupTriplet unsafe.Pointer
	// KODAttributeTypeNetGroups is the attribute type used to specify a list of net groups that a user or host record is a member of.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNetGroups
	KODAttributeTypeNetGroups unsafe.Pointer
	// KODAttributeTypeNetworkInterfaces is the attribute type used to specify network interfaces.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNetworkInterfaces
	KODAttributeTypeNetworkInterfaces unsafe.Pointer
	// KODAttributeTypeNetworkNumber is the attribute type used to specify a network number. Typically found in records of type `kODRecordTypeNetworks`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNetworkNumber
	KODAttributeTypeNetworkNumber unsafe.Pointer
	// KODAttributeTypeNickName is the attribute type used to store a user’s nickname.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNickName
	KODAttributeTypeNickName unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNodeOptions
	KODAttributeTypeNodeOptions unsafe.Pointer
	// KODAttributeTypeNodePath is the attribute type used in neighborhood records to specify the node to search while looking up aliases in the neighborhood.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNodePath
	KODAttributeTypeNodePath unsafe.Pointer
	// KODAttributeTypeNodeRefCount is the attribute type used to specify the total count of node references for a process.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNodeRefCount
	KODAttributeTypeNodeRefCount unsafe.Pointer
	// KODAttributeTypeNodeRefs is the attribute type used to store the node references for a process.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNodeRefs
	KODAttributeTypeNodeRefs unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNodeSASLRealm
	KODAttributeTypeNodeSASLRealm unsafe.Pointer
	// KODAttributeTypeNote is the attribute type of the last name attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNote
	KODAttributeTypeNote unsafe.Pointer
	// KODAttributeTypeNumTableList is the attribute type used to summarize reference table entries as attribute values from the configure node.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeNumTableList
	KODAttributeTypeNumTableList unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeOperatingSystem
	KODAttributeTypeOperatingSystem unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeOperatingSystemVersion
	KODAttributeTypeOperatingSystemVersion unsafe.Pointer
	// KODAttributeTypeOrganizationInfo is the attribute type used to store a user’s organization information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeOrganizationInfo
	KODAttributeTypeOrganizationInfo unsafe.Pointer
	// KODAttributeTypeOrganizationName is the attribute type used to store a user’s organization name.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeOrganizationName
	KODAttributeTypeOrganizationName unsafe.Pointer
	// KODAttributeTypeOriginalHomeDirectory is the attribute type used to store a home directory URL used in local account caching.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeOriginalHomeDirectory
	KODAttributeTypeOriginalHomeDirectory unsafe.Pointer
	// KODAttributeTypeOriginalNFSHomeDirectory is the attribute type used to store an NFS home directory URL used in local account caching.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeOriginalNFSHomeDirectory
	KODAttributeTypeOriginalNFSHomeDirectory unsafe.Pointer
	// KODAttributeTypeOriginalNodeName is the attribute type used to store a node name used in local account caching.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeOriginalNodeName
	KODAttributeTypeOriginalNodeName unsafe.Pointer
	// KODAttributeTypeOwner is the attribute type of the owner attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeOwner
	KODAttributeTypeOwner unsafe.Pointer
	// KODAttributeTypeOwnerGUID is the attribute type of the owner GUID attribute. Typically found in records of type `kODRecordTypeGroups`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeOwnerGUID
	KODAttributeTypeOwnerGUID unsafe.Pointer
	// KODAttributeTypePGPPublicKey is the attribute type used to specify a Pretty Good Privacy public key.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePGPPublicKey
	KODAttributeTypePGPPublicKey unsafe.Pointer
	// KODAttributeTypePIDValue is the attribute type used to specify the PID value.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePIDValue
	KODAttributeTypePIDValue unsafe.Pointer
	// KODAttributeTypePagerNumber is the attribute type used to store a user’s pager number.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePagerNumber
	KODAttributeTypePagerNumber unsafe.Pointer
	// KODAttributeTypeParentLocales is the attribute type for specifying locales of the parent.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeParentLocales
	KODAttributeTypeParentLocales unsafe.Pointer
	// KODAttributeTypePassword is the attribute type of the password attribute.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePassword
	KODAttributeTypePassword unsafe.Pointer
	// KODAttributeTypePasswordPlus is the attribute type of the attribute that holds marker data to indicate possible authentication redirection.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePasswordPlus
	KODAttributeTypePasswordPlus unsafe.Pointer
	// KODAttributeTypePasswordPolicyOptions is the attribute type of the password policy options attribute. Typically found in records of type `kODRecordTypePresetUsers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePasswordPolicyOptions
	KODAttributeTypePasswordPolicyOptions unsafe.Pointer
	// KODAttributeTypePasswordServerList is the attribute type of the password server list attribute, which contains the password server’s replication information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePasswordServerList
	KODAttributeTypePasswordServerList unsafe.Pointer
	// KODAttributeTypePasswordServerLocation is the attribute type of the password server location attribute, which specifies the IP address or domain name of the password server associated with a given directory node.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePasswordServerLocation
	KODAttributeTypePasswordServerLocation unsafe.Pointer
	// KODAttributeTypePhoneContacts is the attribute type used to store a user’s custom phone information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePhoneContacts
	KODAttributeTypePhoneContacts unsafe.Pointer
	// KODAttributeTypePhoneNumber is the attribute type used to store a user’s phone number.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePhoneNumber
	KODAttributeTypePhoneNumber unsafe.Pointer
	// KODAttributeTypePicture is the attribute type of the picture attribute, which specifies the path of the picture for each user displayed in the login window.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePicture
	KODAttributeTypePicture unsafe.Pointer
	// KODAttributeTypePlugInInfo is the attribute type used to specify information about the plug-in that is serving a particular directory node.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePlugInInfo
	KODAttributeTypePlugInInfo unsafe.Pointer
	// KODAttributeTypePluginIndex is the attribute type used to specify the plug-in index.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePluginIndex
	KODAttributeTypePluginIndex unsafe.Pointer
	// KODAttributeTypePort is the attribute type of the port attribute, which indicates the port number a service is available on.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePort
	KODAttributeTypePort unsafe.Pointer
	// KODAttributeTypePostalAddress is the attribute type used to store a user’s postal address.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePostalAddress
	KODAttributeTypePostalAddress unsafe.Pointer
	// KODAttributeTypePostalAddressContacts is the attribute type used to store a user’s custom postal address information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePostalAddressContacts
	KODAttributeTypePostalAddressContacts unsafe.Pointer
	// KODAttributeTypePostalCode is the attribute type used to store a user’s postal code.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePostalCode
	KODAttributeTypePostalCode unsafe.Pointer
	// KODAttributeTypePresetUserIsAdmin is the attribute type used to indicate whether users created from a given preset are administrators. Typically found in records of type `kODRecordTypePresetUsers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePresetUserIsAdmin
	KODAttributeTypePresetUserIsAdmin unsafe.Pointer
	// KODAttributeTypePrimaryComputerGUID is the attribute type used to define the primary computer of a computer group. Typically found in records of type kODRecordTypeComputerGroups.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrimaryComputerGUID
	KODAttributeTypePrimaryComputerGUID unsafe.Pointer
	// KODAttributeTypePrimaryComputerList is the attribute type of the primary computer list attribute, which indicates the computer list a given computer record is associated with.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrimaryComputerList
	KODAttributeTypePrimaryComputerList unsafe.Pointer
	// KODAttributeTypePrimaryGroupID is the attribute type used to define a user’s primary group.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrimaryGroupID
	KODAttributeTypePrimaryGroupID unsafe.Pointer
	// KODAttributeTypePrimaryLocale is the attribute type for specifying the primary locale.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrimaryLocale
	KODAttributeTypePrimaryLocale unsafe.Pointer
	// KODAttributeTypePrimaryNTDomain is an attribute type used to support Kerberos SMB server services.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrimaryNTDomain
	KODAttributeTypePrimaryNTDomain unsafe.Pointer
	// KODAttributeTypePrintServiceInfoText is the attribute type used to define a printer’s service info in plaintext.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrintServiceInfoText
	KODAttributeTypePrintServiceInfoText unsafe.Pointer
	// KODAttributeTypePrintServiceInfoXML is the attribute type used to define a printer’s service info in XML.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrintServiceInfoXML
	KODAttributeTypePrintServiceInfoXML unsafe.Pointer
	// KODAttributeTypePrintServiceUserData is the attribute type used to define a printer’s quota configuration and statistics in XML. Typically found in records of type `kODRecordTypeUsers` or `kODRecordTypePrintServiceUser`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrintServiceUserData
	KODAttributeTypePrintServiceUserData unsafe.Pointer
	// KODAttributeTypePrinter1284DeviceID is the attribute type used to define a printer’s IEEE 1284 device ID.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrinter1284DeviceID
	KODAttributeTypePrinter1284DeviceID unsafe.Pointer
	// KODAttributeTypePrinterLPRHost is the attribute type used to define a printer’s LPR host.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrinterLPRHost
	KODAttributeTypePrinterLPRHost unsafe.Pointer
	// KODAttributeTypePrinterLPRQueue is the attribute type used to define a printer’s LPR queue.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrinterLPRQueue
	KODAttributeTypePrinterLPRQueue unsafe.Pointer
	// KODAttributeTypePrinterMakeAndModel is the attribute type used to define a printer’s make and model. Based on the IPP Printing Specification RFC and IETF IPP-LDAP Printer Record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrinterMakeAndModel
	KODAttributeTypePrinterMakeAndModel unsafe.Pointer
	// KODAttributeTypePrinterType is the attribute type used to define a printer’s type.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrinterType
	KODAttributeTypePrinterType unsafe.Pointer
	// KODAttributeTypePrinterURI is the attribute type used to define a printer’s URI. Based on the IPP Printing Specification RFC and IETF IPP-LDAP Printer Record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrinterURI
	KODAttributeTypePrinterURI unsafe.Pointer
	// KODAttributeTypePrinterXRISupported is the attribute type used to define additional URIs supported by a printer. Based on the IPP Printing Specification RFC and IETF IPP-LDAP Printer Record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePrinterXRISupported
	KODAttributeTypePrinterXRISupported unsafe.Pointer
	// KODAttributeTypeProcessName is the attribute type used to specify the process name.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeProcessName
	KODAttributeTypeProcessName unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeProfiles
	KODAttributeTypeProfiles unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeProfilesTimestamp
	KODAttributeTypeProfilesTimestamp unsafe.Pointer
	// KODAttributeTypeProtocolNumber is the attribute type used to specify a protocol number. Typically found in records of type `kODRecordTypeProtocols`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeProtocolNumber
	KODAttributeTypeProtocolNumber unsafe.Pointer
	// KODAttributeTypeProtocols is the attribute type used to specify a list of protocols.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeProtocols
	KODAttributeTypeProtocols unsafe.Pointer
	// KODAttributeTypePwdAgingPolicy is the attribute type used to store a record’s password aging policy.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypePwdAgingPolicy
	KODAttributeTypePwdAgingPolicy unsafe.Pointer
	// KODAttributeTypeRPCNumber is the attribute type used to specify an RPC number.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeRPCNumber
	KODAttributeTypeRPCNumber unsafe.Pointer
	// KODAttributeTypeReadOnlyNode is the attribute type used to specify a read-only node.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeReadOnlyNode
	KODAttributeTypeReadOnlyNode unsafe.Pointer
	// KODAttributeTypeRealUserID is the attribute type of the real user ID attribute, which is used by Managed Client.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeRealUserID
	KODAttributeTypeRealUserID unsafe.Pointer
	// KODAttributeTypeRecRefCount is the attribute type used to specify the total count of record references for a process.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeRecRefCount
	KODAttributeTypeRecRefCount unsafe.Pointer
	// KODAttributeTypeRecRefs is the attribute type used to store the record references for a process.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeRecRefs
	KODAttributeTypeRecRefs unsafe.Pointer
	// KODAttributeTypeRecordName is the attribute type used to specify a list of names for a record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeRecordName
	KODAttributeTypeRecordName unsafe.Pointer
	// KODAttributeTypeRecordType is the attribute type used to specify the type of a record or a directory node.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeRecordType
	KODAttributeTypeRecordType unsafe.Pointer
	// KODAttributeTypeRelationships is the attribute type used to specify a user’s relationships.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeRelationships
	KODAttributeTypeRelationships unsafe.Pointer
	// KODAttributeTypeRelativeDNPrefix is the attribute type used to map the first native LDAP attribute type needed for building a relative distinguished name.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeRelativeDNPrefix
	KODAttributeTypeRelativeDNPrefix unsafe.Pointer
	// KODAttributeTypeResourceInfo is the attribute type used to specify resource record information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeResourceInfo
	KODAttributeTypeResourceInfo unsafe.Pointer
	// KODAttributeTypeResourceType is the attribute type used to specify the type of a resource record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeResourceType
	KODAttributeTypeResourceType unsafe.Pointer
	// KODAttributeTypeSMBAcctFlags is the attribute type used as an account control flag.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBAcctFlags
	KODAttributeTypeSMBAcctFlags unsafe.Pointer
	// KODAttributeTypeSMBGroupRID is the attribute type used to define PDC SMB interaction with DirectoryService.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBGroupRID
	KODAttributeTypeSMBGroupRID unsafe.Pointer
	// KODAttributeTypeSMBHome is the attribute type used to define the UNC address of a Windows home directory mount point.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBHome
	KODAttributeTypeSMBHome unsafe.Pointer
	// KODAttributeTypeSMBHomeDrive is the attribute type used to define the drive letter of a Windows home directory mount point.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBHomeDrive
	KODAttributeTypeSMBHomeDrive unsafe.Pointer
	// KODAttributeTypeSMBKickoffTime is the attribute type used to define kickoff time in SMB interaction.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBKickoffTime
	KODAttributeTypeSMBKickoffTime unsafe.Pointer
	// KODAttributeTypeSMBLogoffTime is the attribute type used to define logoff time in SMB interaction.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBLogoffTime
	KODAttributeTypeSMBLogoffTime unsafe.Pointer
	// KODAttributeTypeSMBLogonTime is the attribute type used to define logon time in SMB interaction.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBLogonTime
	KODAttributeTypeSMBLogonTime unsafe.Pointer
	// KODAttributeTypeSMBPWDLastSet is an attribute type used in SMB interaction.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBPWDLastSet
	KODAttributeTypeSMBPWDLastSet unsafe.Pointer
	// KODAttributeTypeSMBPrimaryGroupSID is the attribute type used to define an SMB primary group’s security ID, which is stored as a string of up to 64 bytes.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBPrimaryGroupSID
	KODAttributeTypeSMBPrimaryGroupSID unsafe.Pointer
	// KODAttributeTypeSMBProfilePath is the attribute type used to define desktop management information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBProfilePath
	KODAttributeTypeSMBProfilePath unsafe.Pointer
	// KODAttributeTypeSMBRID is an attribute type used in SMB interaction.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBRID
	KODAttributeTypeSMBRID unsafe.Pointer
	// KODAttributeTypeSMBSID is the attribute type used to define an SMB security ID, which is stored as a string of up to 64 bytes. Typically found in records of type `kODRecordTypeUsers`, `kODRecordTypeGroups`, and `kODRecordTypeComputers`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBSID
	KODAttributeTypeSMBSID unsafe.Pointer
	// KODAttributeTypeSMBScriptPath is the attribute type used to define an SMB login script path.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBScriptPath
	KODAttributeTypeSMBScriptPath unsafe.Pointer
	// KODAttributeTypeSMBUserWorkstations is the attribute type used to define a list of workstations a user can log in from.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSMBUserWorkstations
	KODAttributeTypeSMBUserWorkstations unsafe.Pointer
	// KODAttributeTypeSchema is the attribute type used to specify a record’s list of attribute types.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSchema
	KODAttributeTypeSchema unsafe.Pointer
	// KODAttributeTypeSearchPath is the attribute type used to specify the search path used by a search node.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSearchPath
	KODAttributeTypeSearchPath unsafe.Pointer
	// KODAttributeTypeSearchPolicy is the attribute type used to specify the search policy used by a search node.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSearchPolicy
	KODAttributeTypeSearchPolicy unsafe.Pointer
	// KODAttributeTypeServiceType is the attribute type used to define an SMB login script path.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeServiceType
	KODAttributeTypeServiceType unsafe.Pointer
	// KODAttributeTypeServicesLocator is the attribute type used to specify the URI of a record’s calendar.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeServicesLocator
	KODAttributeTypeServicesLocator unsafe.Pointer
	// KODAttributeTypeSetupAdvertising is the attribute type used to define the raw service type of a service. For instance, a service record of type `kODRecordTypeWebServer` could have a service of type `http` or `https`.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSetupAdvertising
	KODAttributeTypeSetupAdvertising unsafe.Pointer
	// KODAttributeTypeSetupAutoRegister is an attribute type used for automatic population in Setup Assistant.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSetupAutoRegister
	KODAttributeTypeSetupAutoRegister unsafe.Pointer
	// KODAttributeTypeSetupLocation is an attribute type used for automatic population in Setup Assistant.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSetupLocation
	KODAttributeTypeSetupLocation unsafe.Pointer
	// KODAttributeTypeSetupOccupation is an attribute type used for automatic population in Setup Assistant.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSetupOccupation
	KODAttributeTypeSetupOccupation unsafe.Pointer
	// KODAttributeTypeStandardOnly is the attribute type used in requesting only standard attribute types in a search.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeStandardOnly
	KODAttributeTypeStandardOnly unsafe.Pointer
	// KODAttributeTypeState is the attribute type used to specify a user’s state or province.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeState
	KODAttributeTypeState unsafe.Pointer
	// KODAttributeTypeStreet is the attribute type used to specify a user’s street address.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeStreet
	KODAttributeTypeStreet unsafe.Pointer
	// KODAttributeTypeSubNodes is the attribute type used to specify a node’s list of subnodes.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeSubNodes
	KODAttributeTypeSubNodes unsafe.Pointer
	// KODAttributeTypeTimePackage is the attribute type used to group a record’s creation, modification, and backup timestamps.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeTimePackage
	KODAttributeTypeTimePackage unsafe.Pointer
	// KODAttributeTypeTimeToLive is the attribute type used to specify how long to cache a record’s attribute values. Specified in seconds.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeTimeToLive
	KODAttributeTypeTimeToLive unsafe.Pointer
	// KODAttributeTypeTotalRefCount is the attribute type used to specify the total count of references for a process.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeTotalRefCount
	KODAttributeTypeTotalRefCount unsafe.Pointer
	// KODAttributeTypeTotalSize is an attribute type used for checksum and accessing metadata.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeTotalSize
	KODAttributeTypeTotalSize unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeTrustInformation
	KODAttributeTypeTrustInformation unsafe.Pointer
	// KODAttributeTypeURL is the attribute type used to specify a list of URLs.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeURL
	KODAttributeTypeURL unsafe.Pointer
	// KODAttributeTypeUniqueID is the attribute type used to define a user’s unique 32-bit ID in the legacy manner.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeUniqueID
	KODAttributeTypeUniqueID unsafe.Pointer
	// KODAttributeTypeUserCertificate is the attribute type used to store the binary of a user’s certificate. Typically found in user records.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeUserCertificate
	KODAttributeTypeUserCertificate unsafe.Pointer
	// KODAttributeTypeUserPKCS12Data is the attribute type used to store binary data in PKCS #12 format, including keys and certificates. Typically found in user records.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeUserPKCS12Data
	KODAttributeTypeUserPKCS12Data unsafe.Pointer
	// KODAttributeTypeUserSMIMECertificate is the attribute type used to store the binary of a user’s SMIME certificate. Typically found in user records.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeUserSMIMECertificate
	KODAttributeTypeUserSMIMECertificate unsafe.Pointer
	// KODAttributeTypeUserShell is the attribute type used to specify a user’s shell setting.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeUserShell
	KODAttributeTypeUserShell unsafe.Pointer
	// KODAttributeTypeVFSDumpFreq is an attribute type used to support mount records.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeVFSDumpFreq
	KODAttributeTypeVFSDumpFreq unsafe.Pointer
	// KODAttributeTypeVFSLinkDir is an attribute type used to support mount records.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeVFSLinkDir
	KODAttributeTypeVFSLinkDir unsafe.Pointer
	// KODAttributeTypeVFSOpts is an attribute type used to support mount records.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeVFSOpts
	KODAttributeTypeVFSOpts unsafe.Pointer
	// KODAttributeTypeVFSPassNo is an attribute type used to support mount records.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeVFSPassNo
	KODAttributeTypeVFSPassNo unsafe.Pointer
	// KODAttributeTypeVFSType is an attribute type used to support mount records.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeVFSType
	KODAttributeTypeVFSType unsafe.Pointer
	// KODAttributeTypeVersion is the attribute type used to specify the version.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeVersion
	KODAttributeTypeVersion unsafe.Pointer
	// KODAttributeTypeWeblogURI is the attribute type used to specify the URI of a user’s weblog.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeWeblogURI
	KODAttributeTypeWeblogURI unsafe.Pointer
	// KODAttributeTypeXMLPlist is the attribute type used to specify an XML property list.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAttributeTypeXMLPlist
	KODAttributeTypeXMLPlist unsafe.Pointer
	// KODRecordTypeAFPServer is the record type of an AFP server record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeAFPServer
	KODRecordTypeAFPServer unsafe.Pointer
	// KODRecordTypeAliases is the record type of an alias record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeAliases
	KODRecordTypeAliases unsafe.Pointer
	// KODRecordTypeAttributeTypes is a record that represents each possible attribute type.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeAttributeTypes
	KODRecordTypeAttributeTypes unsafe.Pointer
	// KODRecordTypeAugments is the record type of augmented record data.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeAugments
	KODRecordTypeAugments unsafe.Pointer
	// KODRecordTypeAutoServerSetup is the record type used to discover automated server setup information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeAutoServerSetup
	KODRecordTypeAutoServerSetup unsafe.Pointer
	// KODRecordTypeAutomount is the record type of automount record data.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeAutomount
	KODRecordTypeAutomount unsafe.Pointer
	// KODRecordTypeAutomountMap is the record type of automount map record data.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeAutomountMap
	KODRecordTypeAutomountMap unsafe.Pointer
	// KODRecordTypeBootp is the record type of a record in the local node for storing bootp information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeBootp
	KODRecordTypeBootp unsafe.Pointer
	// KODRecordTypeCertificateAuthorities is the record type of a record that contains certificate authority information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeCertificateAuthorities
	KODRecordTypeCertificateAuthorities unsafe.Pointer
	// KODRecordTypeComputerGroups is the record type of a computer group record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeComputerGroups
	KODRecordTypeComputerGroups unsafe.Pointer
	// KODRecordTypeComputerLists is the record type of a computer list record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeComputerLists
	KODRecordTypeComputerLists unsafe.Pointer
	// KODRecordTypeComputers is the record type of a computer record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeComputers
	KODRecordTypeComputers unsafe.Pointer
	// KODRecordTypeConfiguration is the record type of a configuration record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeConfiguration
	KODRecordTypeConfiguration unsafe.Pointer
	// KODRecordTypeEthernets is the record type of a record in a node storing ethernet information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeEthernets
	KODRecordTypeEthernets unsafe.Pointer
	// KODRecordTypeFTPServer is the record type of an FTP server record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeFTPServer
	KODRecordTypeFTPServer unsafe.Pointer
	// KODRecordTypeFileMakerServers is the record type of a FileMaker server record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeFileMakerServers
	KODRecordTypeFileMakerServers unsafe.Pointer
	// KODRecordTypeGroups is the record type of a group record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeGroups
	KODRecordTypeGroups unsafe.Pointer
	// KODRecordTypeHostServices is the record type of a record in the local node for storing host services information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeHostServices
	KODRecordTypeHostServices unsafe.Pointer
	// KODRecordTypeHosts is the record type of a host record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeHosts
	KODRecordTypeHosts unsafe.Pointer
	// KODRecordTypeLDAPServer is the record type of an LDAP server record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeLDAPServer
	KODRecordTypeLDAPServer unsafe.Pointer
	// KODRecordTypeLocations is the record type of a location record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeLocations
	KODRecordTypeLocations unsafe.Pointer
	// KODRecordTypeMounts is the record type of a mount record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeMounts
	KODRecordTypeMounts unsafe.Pointer
	// KODRecordTypeNFS is the record type of an NFS record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeNFS
	KODRecordTypeNFS unsafe.Pointer
	// KODRecordTypeNetDomains is the record type of a record in the local node for storing net domain information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeNetDomains
	KODRecordTypeNetDomains unsafe.Pointer
	// KODRecordTypeNetGroups is the record type of a record in the local node for storing net group information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeNetGroups
	KODRecordTypeNetGroups unsafe.Pointer
	// KODRecordTypeNetworks is the record type of a network record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeNetworks
	KODRecordTypeNetworks unsafe.Pointer
	// KODRecordTypePeople is the record type of a “people” record, used for storing contact information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypePeople
	KODRecordTypePeople unsafe.Pointer
	// KODRecordTypePresetComputerGroups is the record type used for presets in computer group record creation.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypePresetComputerGroups
	KODRecordTypePresetComputerGroups unsafe.Pointer
	// KODRecordTypePresetComputerLists is the record type used for presets in computer list record creation.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypePresetComputerLists
	KODRecordTypePresetComputerLists unsafe.Pointer
	// KODRecordTypePresetComputers is the record type used for presets in computer record creation.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypePresetComputers
	KODRecordTypePresetComputers unsafe.Pointer
	// KODRecordTypePresetGroups is the record type used for presets in group record creation.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypePresetGroups
	KODRecordTypePresetGroups unsafe.Pointer
	// KODRecordTypePresetUsers is the record type used for presets in user record creation.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypePresetUsers
	KODRecordTypePresetUsers unsafe.Pointer
	// KODRecordTypePrintService is the record type of a print service record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypePrintService
	KODRecordTypePrintService unsafe.Pointer
	// KODRecordTypePrintServiceUser is the record type of a record in the local node for storing a user’s quota usage information.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypePrintServiceUser
	KODRecordTypePrintServiceUser unsafe.Pointer
	// KODRecordTypePrinters is the record type of a printer record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypePrinters
	KODRecordTypePrinters unsafe.Pointer
	// KODRecordTypeProtocols is the record type of a protocol record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeProtocols
	KODRecordTypeProtocols unsafe.Pointer
	// KODRecordTypeQTSServer is the record type of a QuickTime streaming server record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeQTSServer
	KODRecordTypeQTSServer unsafe.Pointer
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeQueryInformation
	KODRecordTypeQueryInformation unsafe.Pointer
	// KODRecordTypeRPC is the record type of a remote procedure call record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeRPC
	KODRecordTypeRPC unsafe.Pointer
	// KODRecordTypeRecordTypes is a record that represents each possible record type.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeRecordTypes
	KODRecordTypeRecordTypes unsafe.Pointer
	// KODRecordTypeResources is the record type of a resource used in group services.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeResources
	KODRecordTypeResources unsafe.Pointer
	// KODRecordTypeSMBServer is the record type of an SMB server record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeSMBServer
	KODRecordTypeSMBServer unsafe.Pointer
	// KODRecordTypeServer is the record type of a generic server record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeServer
	KODRecordTypeServer unsafe.Pointer
	// KODRecordTypeServices is the record type of a directory based service.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeServices
	KODRecordTypeServices unsafe.Pointer
	// KODRecordTypeSharePoints is the record type of a SharePoint record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeSharePoints
	KODRecordTypeSharePoints unsafe.Pointer
	// KODRecordTypeUsers is the record type of a user record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeUsers
	KODRecordTypeUsers unsafe.Pointer
	// KODRecordTypeWebServer is the record type of a web server record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODRecordTypeWebServer
	KODRecordTypeWebServer unsafe.Pointer
)

var (
	// KODAuthenticationType2WayRandom is the authentication type used to specify two way random authentication.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationType2WayRandom
	KODAuthenticationType2WayRandom ODAuthenticationType
	// KODAuthenticationType2WayRandomChangePasswd is the authentication type used to change a user’s password using two way random authentication.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationType2WayRandomChangePasswd
	KODAuthenticationType2WayRandomChangePasswd ODAuthenticationType
	// KODAuthenticationTypeAPOP is the authentication type used to specify APOP authentication.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeAPOP
	KODAuthenticationTypeAPOP ODAuthenticationType
	// KODAuthenticationTypeCRAM_MD5 is the authentication type used to specify CRAM MD5 authentication.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeCRAM_MD5
	KODAuthenticationTypeCRAM_MD5 ODAuthenticationType
	// KODAuthenticationTypeChangePasswd is the authentication type used to change a user’s password using CRAM MD5 authentication.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeChangePasswd
	KODAuthenticationTypeChangePasswd ODAuthenticationType
	// KODAuthenticationTypeClearText is the authentication type used to specify cleartext authentication.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeClearText
	KODAuthenticationTypeClearText ODAuthenticationType
	// KODAuthenticationTypeCrypt is the authentication type used to specify crypt authentication, which uses a crypt password stored in a user’s record if available.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeCrypt
	KODAuthenticationTypeCrypt ODAuthenticationType
	// KODAuthenticationTypeDIGEST_MD5 is the authentication type used to specify digest MD5 authentication.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeDIGEST_MD5
	KODAuthenticationTypeDIGEST_MD5 ODAuthenticationType
	// KODAuthenticationTypeDeleteUser is the authentication type used to specify that a user on an Apple password server be deleted.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeDeleteUser
	KODAuthenticationTypeDeleteUser ODAuthenticationType
	// KODAuthenticationTypeGetEffectivePolicy is the authentication type used to access the policies applied to a user.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeGetEffectivePolicy
	KODAuthenticationTypeGetEffectivePolicy ODAuthenticationType
	// KODAuthenticationTypeGetGlobalPolicy is the authentication type used to access the global authentication policy.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeGetGlobalPolicy
	KODAuthenticationTypeGetGlobalPolicy ODAuthenticationType
	// KODAuthenticationTypeGetKerberosPrincipal is the authentication type used to access the name of the Kerberos principal.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeGetKerberosPrincipal
	KODAuthenticationTypeGetKerberosPrincipal ODAuthenticationType
	// KODAuthenticationTypeGetPolicy is the authentication type used to specify that the plug-in should determine the authentication method to use.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeGetPolicy
	KODAuthenticationTypeGetPolicy ODAuthenticationType
	// KODAuthenticationTypeGetUserData is the authentication type used to access user data on an Apple password server.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeGetUserData
	KODAuthenticationTypeGetUserData ODAuthenticationType
	// KODAuthenticationTypeGetUserName is the authentication type used to access a username on an Apple password server.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeGetUserName
	KODAuthenticationTypeGetUserName ODAuthenticationType
	// KODAuthenticationTypeKerberosTickets is the authentication type used to provide write access to LDAP with an existing Kerberos ticket.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeKerberosTickets
	KODAuthenticationTypeKerberosTickets ODAuthenticationType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeMPPEPrimaryKeys
	KODAuthenticationTypeMPPEPrimaryKeys ODAuthenticationType
	// KODAuthenticationTypeMSCHAP2 is the authentication type used to specify MS-CHAPv2 encryption.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeMSCHAP2
	KODAuthenticationTypeMSCHAP2 ODAuthenticationType
	// KODAuthenticationTypeNTLMv2 is the authentication type used to verify an NTLMv2 challenge and response.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeNTLMv2
	KODAuthenticationTypeNTLMv2 ODAuthenticationType
	// KODAuthenticationTypeNTLMv2WithSessionKey is the authentication type used to verify an NTLMv2 challenge and response and retrieve session keys in a single call.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeNTLMv2WithSessionKey
	KODAuthenticationTypeNTLMv2WithSessionKey ODAuthenticationType
	// KODAuthenticationTypeNewUser is the authentication type used to create a new user on an Apple password server.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeNewUser
	KODAuthenticationTypeNewUser ODAuthenticationType
	// KODAuthenticationTypeNewUserWithPolicy is the authentication type used to create a new user with specified policy settings on an Apple password server.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeNewUserWithPolicy
	KODAuthenticationTypeNewUserWithPolicy ODAuthenticationType
	// KODAuthenticationTypeNodeNativeClearTextOK is the authentication type used to specify that the plug-in should determine the authentication method to use. It also specifies that cleartext is an acceptable authentication method.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeNodeNativeClearTextOK
	KODAuthenticationTypeNodeNativeClearTextOK ODAuthenticationType
	// KODAuthenticationTypeNodeNativeNoClearText is the authentication type used to specify that the plug-in should determine the authentication method to use. It also specifies that cleartext is not an acceptable authentication method.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeNodeNativeNoClearText
	KODAuthenticationTypeNodeNativeNoClearText ODAuthenticationType
	// KODAuthenticationTypeReadSecureHash is the authentication type used to access the SHA1 or seeded SHA1 hash for a local user.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeReadSecureHash
	KODAuthenticationTypeReadSecureHash ODAuthenticationType
	// KODAuthenticationTypeSMBNTv2UserSessionKey is the authentication type used to generate an NTLMv2 user session key.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSMBNTv2UserSessionKey
	KODAuthenticationTypeSMBNTv2UserSessionKey ODAuthenticationType
	// KODAuthenticationTypeSMBWorkstationCredentialSessionKey is the authentication type used to generate an SMB workstation credential session key.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSMBWorkstationCredentialSessionKey
	KODAuthenticationTypeSMBWorkstationCredentialSessionKey ODAuthenticationType
	// KODAuthenticationTypeSMB_LM_Key is the authentication type used to specify SMB LAN manager authentication.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSMB_LM_Key
	KODAuthenticationTypeSMB_LM_Key ODAuthenticationType
	// KODAuthenticationTypeSMB_NT_Key is the authentication type used to specify SMB NT authentication.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSMB_NT_Key
	KODAuthenticationTypeSMB_NT_Key ODAuthenticationType
	// KODAuthenticationTypeSMB_NT_UserSessionKey is the authentication type used by Samba to access session keys on an Apple password server.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSMB_NT_UserSessionKey
	KODAuthenticationTypeSMB_NT_UserSessionKey ODAuthenticationType
	// KODAuthenticationTypeSMB_NT_WithUserSessionKey is the authentication type used by Samba to authenticate and access session keys on an Apple password server.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSMB_NT_WithUserSessionKey
	KODAuthenticationTypeSMB_NT_WithUserSessionKey ODAuthenticationType
	// KODAuthenticationTypeSetGlobalPolicy is the authentication type used to set the global authentication policy.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSetGlobalPolicy
	KODAuthenticationTypeSetGlobalPolicy ODAuthenticationType
	// KODAuthenticationTypeSetLMHash is the authentication type used to set the LAN manager hash for an account.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSetLMHash
	KODAuthenticationTypeSetLMHash ODAuthenticationType
	// KODAuthenticationTypeSetNTHash is the authentication type used to set the NT hash for a user.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSetNTHash
	KODAuthenticationTypeSetNTHash ODAuthenticationType
	// KODAuthenticationTypeSetPassword is the authentication type used to set a password.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSetPassword
	KODAuthenticationTypeSetPassword ODAuthenticationType
	// KODAuthenticationTypeSetPasswordAsCurrent is the authentication type used to set a password using the current credentials.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSetPasswordAsCurrent
	KODAuthenticationTypeSetPasswordAsCurrent ODAuthenticationType
	// KODAuthenticationTypeSetPolicy is the authentication type used to specify that the plug-in should determine the authentication method to use.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSetPolicy
	KODAuthenticationTypeSetPolicy ODAuthenticationType
	// KODAuthenticationTypeSetPolicyAsCurrent is the authentication type used to set the authentication policy using the current credentials.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSetPolicyAsCurrent
	KODAuthenticationTypeSetPolicyAsCurrent ODAuthenticationType
	// KODAuthenticationTypeSetUserData is the authentication type used to set user data on an Apple password server.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSetUserData
	KODAuthenticationTypeSetUserData ODAuthenticationType
	// KODAuthenticationTypeSetUserName is the authentication type used to set a username on an Apple password server.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSetUserName
	KODAuthenticationTypeSetUserName ODAuthenticationType
	// KODAuthenticationTypeSetWorkstationPassword is an authentication type used to support PDC SMB interaction with Directory Services.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeSetWorkstationPassword
	KODAuthenticationTypeSetWorkstationPassword ODAuthenticationType
	// KODAuthenticationTypeWithAuthorizationRef is the authentication type used to allow root access to local directories with valid authorization.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeWithAuthorizationRef
	KODAuthenticationTypeWithAuthorizationRef ODAuthenticationType
	// KODAuthenticationTypeWriteSecureHash is the authentication type used to enable a root process to write the secure hash of a user record.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODAuthenticationTypeWriteSecureHash
	KODAuthenticationTypeWriteSecureHash ODAuthenticationType
)

var (
	// See: https://developer.apple.com/documentation/OpenDirectory/kODBackOffSeconds
	KODBackOffSeconds ODErrorUserInfoKeyType
)

var (
	// KODErrorDomainFramework is the error domain used for errors from the Open Directory framework.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODErrorDomainFramework
	KODErrorDomainFramework string
	// See: https://developer.apple.com/documentation/OpenDirectory/kODModuleConfigOptionConnectionIdleDisconnect
	KODModuleConfigOptionConnectionIdleDisconnect string
	// See: https://developer.apple.com/documentation/OpenDirectory/kODModuleConfigOptionConnectionSetupTimeout
	KODModuleConfigOptionConnectionSetupTimeout string
	// See: https://developer.apple.com/documentation/OpenDirectory/kODModuleConfigOptionManInTheMiddle
	KODModuleConfigOptionManInTheMiddle string
	// See: https://developer.apple.com/documentation/OpenDirectory/kODModuleConfigOptionPacketEncryption
	KODModuleConfigOptionPacketEncryption string
	// See: https://developer.apple.com/documentation/OpenDirectory/kODModuleConfigOptionPacketSigning
	KODModuleConfigOptionPacketSigning string
	// See: https://developer.apple.com/documentation/OpenDirectory/kODModuleConfigOptionQueryTimeout
	KODModuleConfigOptionQueryTimeout string
	// See: https://developer.apple.com/documentation/OpenDirectory/kODNodeOptionsQuerySkippedSubnode
	KODNodeOptionsQuerySkippedSubnode string
	// KODSessionProxyAddress is the address to connect to via proxy. The value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODSessionProxyAddress
	KODSessionProxyAddress string
	// KODSessionProxyPassword is the password to connect with via proxy. The value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODSessionProxyPassword
	KODSessionProxyPassword string
	// KODSessionProxyPort is the port to connect to via proxy. The value is of type [CFNumber].
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODSessionProxyPort
	KODSessionProxyPort string
	// KODSessionProxyUsername is the username to connect with via proxy. The value is of type [CFString].
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODSessionProxyUsername
	KODSessionProxyUsername string
)

var (
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeCreationTime
	KODPolicyAttributeCreationTime ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeCurrentDate
	KODPolicyAttributeCurrentDate ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeCurrentDayOfWeek
	KODPolicyAttributeCurrentDayOfWeek ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeCurrentTime
	KODPolicyAttributeCurrentTime ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeCurrentTimeOfDay
	KODPolicyAttributeCurrentTimeOfDay ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeDaysUntilExpiration
	KODPolicyAttributeDaysUntilExpiration ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeEnableAtTimeOfDay
	KODPolicyAttributeEnableAtTimeOfDay ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeEnableOnDate
	KODPolicyAttributeEnableOnDate ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeEnableOnDayOfWeek
	KODPolicyAttributeEnableOnDayOfWeek ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeExpiresAtTimeOfDay
	KODPolicyAttributeExpiresAtTimeOfDay ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeExpiresEveryNDays
	KODPolicyAttributeExpiresEveryNDays ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeExpiresOnDate
	KODPolicyAttributeExpiresOnDate ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeExpiresOnDayOfWeek
	KODPolicyAttributeExpiresOnDayOfWeek ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeFailedAuthentications
	KODPolicyAttributeFailedAuthentications ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeLastAuthenticationTime
	KODPolicyAttributeLastAuthenticationTime ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeLastFailedAuthenticationTime
	KODPolicyAttributeLastFailedAuthenticationTime ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeLastPasswordChangeTime
	KODPolicyAttributeLastPasswordChangeTime ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeMaximumFailedAuthentications
	KODPolicyAttributeMaximumFailedAuthentications ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeNewPasswordRequiredTime
	KODPolicyAttributeNewPasswordRequiredTime ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributePassword
	KODPolicyAttributePassword ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributePasswordHashes
	KODPolicyAttributePasswordHashes ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributePasswordHistory
	KODPolicyAttributePasswordHistory ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributePasswordHistoryDepth
	KODPolicyAttributePasswordHistoryDepth ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeRecordName
	KODPolicyAttributeRecordName ODPolicyAttributeType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyAttributeRecordType
	KODPolicyAttributeRecordType ODPolicyAttributeType
)

var (
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyCategoryAuthentication
	KODPolicyCategoryAuthentication ODPolicyCategoryType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyCategoryPasswordChange
	KODPolicyCategoryPasswordChange ODPolicyCategoryType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyCategoryPasswordContent
	KODPolicyCategoryPasswordContent ODPolicyCategoryType
)

var (
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyKeyContent
	KODPolicyKeyContent ODPolicyKeyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyKeyContentDescription
	KODPolicyKeyContentDescription ODPolicyKeyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyKeyEvaluationDetails
	KODPolicyKeyEvaluationDetails ODPolicyKeyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyKeyIdentifier
	KODPolicyKeyIdentifier ODPolicyKeyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyKeyParameters
	KODPolicyKeyParameters ODPolicyKeyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyKeyPolicySatisfied
	KODPolicyKeyPolicySatisfied ODPolicyKeyType
)

var (
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypeAccountExpiresOnDate
	KODPolicyTypeAccountExpiresOnDate ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypeAccountMaximumFailedLogins
	KODPolicyTypeAccountMaximumFailedLogins ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypeAccountMaximumMinutesOfNonUse
	KODPolicyTypeAccountMaximumMinutesOfNonUse ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypeAccountMaximumMinutesUntilDisabled
	KODPolicyTypeAccountMaximumMinutesUntilDisabled ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypeAccountMinutesUntilFailedLoginReset
	KODPolicyTypeAccountMinutesUntilFailedLoginReset ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypePasswordCannotBeAccountName
	KODPolicyTypePasswordCannotBeAccountName ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypePasswordChangeRequired
	KODPolicyTypePasswordChangeRequired ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypePasswordHistory
	KODPolicyTypePasswordHistory ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypePasswordMaximumAgeInMinutes
	KODPolicyTypePasswordMaximumAgeInMinutes ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypePasswordMaximumNumberOfCharacters
	KODPolicyTypePasswordMaximumNumberOfCharacters ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypePasswordMinimumNumberOfCharacters
	KODPolicyTypePasswordMinimumNumberOfCharacters ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypePasswordRequiresAlpha
	KODPolicyTypePasswordRequiresAlpha ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypePasswordRequiresMixedCase
	KODPolicyTypePasswordRequiresMixedCase ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypePasswordRequiresNumeric
	KODPolicyTypePasswordRequiresNumeric ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypePasswordRequiresSymbol
	KODPolicyTypePasswordRequiresSymbol ODPolicyType
	// See: https://developer.apple.com/documentation/OpenDirectory/kODPolicyTypePasswordSelfModification
	KODPolicyTypePasswordSelfModification ODPolicyType
)

var (
	// KODSessionDefault is the default session. Used if there is no need to create a specific reference.
	//
	// See: https://developer.apple.com/documentation/OpenDirectory/kODSessionDefault
	KODSessionDefault ODSessionRef
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ODFrameworkErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ODFrameworkErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ODSessionProxyAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ODSessionProxyAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ODSessionProxyPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ODSessionProxyPassword = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ODSessionProxyPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ODSessionProxyPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ODSessionProxyUsername"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ODSessionProxyUsername = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ODTrustTypeAnonymous"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ODTrustTypeAnonymous = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ODTrustTypeJoined"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ODTrustTypeJoined = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "ODTrustTypeUsingCredentials"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ODTrustTypeUsingCredentials = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAccessControlEntry"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAccessControlEntry = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAddressLine1"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAddressLine1 = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAddressLine2"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAddressLine2 = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAddressLine3"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAddressLine3 = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAdminLimits"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAdminLimits = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAdvertisedServices"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAdvertisedServices = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAlias"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAlias = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAllAttributes"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAllAttributes = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAllTypes"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAllTypes = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAltSecurityIdentities"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAltSecurityIdentities = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAreaCode"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAreaCode = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAttrListRefCount"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAttrListRefCount = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAttrListRefs"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAttrListRefs = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAttrListValueRefCount"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAttrListValueRefCount = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAttrListValueRefs"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAttrListValueRefs = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAuthCredential"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAuthCredential = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAuthMethod"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAuthMethod = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAuthenticationAuthority"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAuthenticationAuthority = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAuthenticationHint"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAuthenticationHint = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAuthorityRevocationList"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAuthorityRevocationList = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAutomaticSearchPath"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAutomaticSearchPath = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeAutomountInformation"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeAutomountInformation = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeBirthday"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeBirthday = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeBootParams"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeBootParams = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeBuildVersion"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeBuildVersion = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeBuilding"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeBuilding = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeCACertificate"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeCACertificate = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeCapacity"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeCapacity = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeCertificateRevocationList"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeCertificateRevocationList = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeCity"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeCity = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeComment"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeComment = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeCompany"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeCompany = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeComputers"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeComputers = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeConfigAvailable"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeConfigAvailable = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeConfigFile"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeConfigFile = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeContactGUID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeContactGUID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeContactPerson"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeContactPerson = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeCopyTimestamp"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeCopyTimestamp = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeCoreFWVersion"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeCoreFWVersion = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeCountry"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeCountry = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeCreationTimestamp"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeCreationTimestamp = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeCrossCertificatePair"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeCrossCertificatePair = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeCustomSearchPath"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeCustomSearchPath = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeDNSDomain"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeDNSDomain = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeDNSName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeDNSName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeDNSNameServer"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeDNSNameServer = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeDataStamp"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeDataStamp = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeDateRecordCreated"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeDateRecordCreated = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeDepartment"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeDepartment = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeDirRefCount"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeDirRefCount = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeDirRefs"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeDirRefs = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeEMailAddress"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeEMailAddress = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeEMailContacts"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeEMailContacts = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeENetAddress"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeENetAddress = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeExpire"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeExpire = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeFWVersion"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeFWVersion = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeFaxNumber"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeFaxNumber = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeFirstName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeFirstName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeFullName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeFullName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeFunctionalState"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeFunctionalState = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeGUID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeGUID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeGroup"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeGroup = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeGroupMembers"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeGroupMembers = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeGroupMembership"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeGroupMembership = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeGroupServices"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeGroupServices = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeHTML"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeHTML = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeHardwareUUID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeHardwareUUID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeHomeDirectory"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeHomeDirectory = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeHomeDirectoryQuota"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeHomeDirectoryQuota = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeHomeDirectorySoftQuota"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeHomeDirectorySoftQuota = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeHomeLocOwner"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeHomeLocOwner = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeHomePhoneNumber"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeHomePhoneNumber = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeIMHandle"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeIMHandle = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeIPAddress"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeIPAddress = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeIPAddressAndENetAddress"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeIPAddressAndENetAddress = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeIPv6Address"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeIPv6Address = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeInternetAlias"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeInternetAlias = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeJPEGPhoto"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeJPEGPhoto = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeJobTitle"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeJobTitle = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeKDCAuthKey"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeKDCAuthKey = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeKDCConfigData"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeKDCConfigData = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeKerberosRealm"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeKerberosRealm = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeKerberosServices"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeKerberosServices = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeKeywords"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeKeywords = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeLDAPReadReplicas"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeLDAPReadReplicas = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeLDAPSearchBaseSuffix"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeLDAPSearchBaseSuffix = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeLDAPWriteReplicas"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeLDAPWriteReplicas = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeLastName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeLastName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeLocalOnlySearchPath"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeLocalOnlySearchPath = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeLocaleRelay"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeLocaleRelay = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeLocaleSubnets"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeLocaleSubnets = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeLocation"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeLocation = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMCXFlags"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMCXFlags = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMCXSettings"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMCXSettings = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMIME"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMIME = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMailAttribute"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMailAttribute = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMapCoordinates"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMapCoordinates = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMapGUID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMapGUID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMapURI"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMapURI = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMetaAmbiguousName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMetaAmbiguousName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMetaAugmentedAttributes"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMetaAugmentedAttributes = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMetaAutomountMap"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMetaAutomountMap = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMetaNodeLocation"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMetaNodeLocation = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMetaRecordName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMetaRecordName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMiddleName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMiddleName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeMobileNumber"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeMobileNumber = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeModificationTimestamp"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeModificationTimestamp = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNFSHomeDirectory"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNFSHomeDirectory = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNTDomainComputerAccount"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNTDomainComputerAccount = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNamePrefix"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNamePrefix = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNameSuffix"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNameSuffix = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNativeOnly"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNativeOnly = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNestedGroups"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNestedGroups = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNetGroupTriplet"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNetGroupTriplet = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNetGroups"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNetGroups = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNetworkInterfaces"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNetworkInterfaces = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNetworkNumber"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNetworkNumber = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNickName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNickName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNodeOptions"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNodeOptions = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNodePath"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNodePath = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNodeRefCount"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNodeRefCount = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNodeRefs"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNodeRefs = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNodeSASLRealm"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNodeSASLRealm = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNote"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNote = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeNumTableList"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeNumTableList = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeOperatingSystem"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeOperatingSystem = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeOperatingSystemVersion"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeOperatingSystemVersion = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeOrganizationInfo"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeOrganizationInfo = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeOrganizationName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeOrganizationName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeOriginalHomeDirectory"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeOriginalHomeDirectory = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeOriginalNFSHomeDirectory"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeOriginalNFSHomeDirectory = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeOriginalNodeName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeOriginalNodeName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeOwner"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeOwner = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeOwnerGUID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeOwnerGUID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePGPPublicKey"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePGPPublicKey = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePIDValue"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePIDValue = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePagerNumber"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePagerNumber = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeParentLocales"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeParentLocales = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePassword"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePassword = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePasswordPlus"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePasswordPlus = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePasswordPolicyOptions"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePasswordPolicyOptions = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePasswordServerList"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePasswordServerList = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePasswordServerLocation"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePasswordServerLocation = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePhoneContacts"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePhoneContacts = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePhoneNumber"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePhoneNumber = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePicture"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePicture = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePlugInInfo"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePlugInInfo = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePluginIndex"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePluginIndex = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePort"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePort = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePostalAddress"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePostalAddress = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePostalAddressContacts"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePostalAddressContacts = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePostalCode"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePostalCode = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePresetUserIsAdmin"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePresetUserIsAdmin = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrimaryComputerGUID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrimaryComputerGUID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrimaryComputerList"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrimaryComputerList = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrimaryGroupID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrimaryGroupID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrimaryLocale"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrimaryLocale = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrimaryNTDomain"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrimaryNTDomain = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrintServiceInfoText"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrintServiceInfoText = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrintServiceInfoXML"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrintServiceInfoXML = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrintServiceUserData"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrintServiceUserData = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrinter1284DeviceID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrinter1284DeviceID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrinterLPRHost"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrinterLPRHost = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrinterLPRQueue"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrinterLPRQueue = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrinterMakeAndModel"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrinterMakeAndModel = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrinterType"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrinterType = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrinterURI"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrinterURI = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePrinterXRISupported"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePrinterXRISupported = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeProcessName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeProcessName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeProfiles"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeProfiles = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeProfilesTimestamp"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeProfilesTimestamp = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeProtocolNumber"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeProtocolNumber = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeProtocols"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeProtocols = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypePwdAgingPolicy"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypePwdAgingPolicy = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeRPCNumber"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeRPCNumber = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeReadOnlyNode"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeReadOnlyNode = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeRealUserID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeRealUserID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeRecRefCount"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeRecRefCount = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeRecRefs"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeRecRefs = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeRecordName"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeRecordName = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeRecordType"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeRecordType = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeRelationships"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeRelationships = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeRelativeDNPrefix"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeRelativeDNPrefix = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeResourceInfo"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeResourceInfo = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeResourceType"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeResourceType = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBAcctFlags"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBAcctFlags = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBGroupRID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBGroupRID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBHome"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBHome = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBHomeDrive"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBHomeDrive = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBKickoffTime"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBKickoffTime = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBLogoffTime"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBLogoffTime = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBLogonTime"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBLogonTime = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBPWDLastSet"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBPWDLastSet = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBPrimaryGroupSID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBPrimaryGroupSID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBProfilePath"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBProfilePath = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBRID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBRID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBSID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBSID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBScriptPath"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBScriptPath = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSMBUserWorkstations"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSMBUserWorkstations = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSchema"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSchema = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSearchPath"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSearchPath = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSearchPolicy"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSearchPolicy = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeServiceType"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeServiceType = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeServicesLocator"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeServicesLocator = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSetupAdvertising"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSetupAdvertising = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSetupAutoRegister"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSetupAutoRegister = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSetupLocation"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSetupLocation = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSetupOccupation"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSetupOccupation = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeStandardOnly"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeStandardOnly = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeState"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeState = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeStreet"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeStreet = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeSubNodes"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeSubNodes = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeTimePackage"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeTimePackage = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeTimeToLive"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeTimeToLive = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeTotalRefCount"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeTotalRefCount = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeTotalSize"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeTotalSize = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeTrustInformation"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeTrustInformation = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeURL"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeURL = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeUniqueID"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeUniqueID = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeUserCertificate"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeUserCertificate = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeUserPKCS12Data"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeUserPKCS12Data = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeUserSMIMECertificate"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeUserSMIMECertificate = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeUserShell"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeUserShell = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeVFSDumpFreq"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeVFSDumpFreq = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeVFSLinkDir"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeVFSLinkDir = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeVFSOpts"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeVFSOpts = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeVFSPassNo"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeVFSPassNo = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeVFSType"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeVFSType = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeVersion"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeVersion = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeWeblogURI"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeWeblogURI = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAttributeTypeXMLPlist"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODAttributeTypeXMLPlist = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationType2WayRandom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationType2WayRandom = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationType2WayRandomChangePasswd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationType2WayRandomChangePasswd = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeAPOP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeAPOP = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeCRAM_MD5"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeCRAM_MD5 = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeChangePasswd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeChangePasswd = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeClearText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeClearText = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeCrypt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeCrypt = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeDIGEST_MD5"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeDIGEST_MD5 = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeDeleteUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeDeleteUser = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeGetEffectivePolicy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeGetEffectivePolicy = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeGetGlobalPolicy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeGetGlobalPolicy = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeGetKerberosPrincipal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeGetKerberosPrincipal = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeGetPolicy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeGetPolicy = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeGetUserData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeGetUserData = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeGetUserName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeGetUserName = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeKerberosTickets"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeKerberosTickets = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeMPPEPrimaryKeys"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeMPPEPrimaryKeys = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeMSCHAP2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeMSCHAP2 = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeNTLMv2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeNTLMv2 = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeNTLMv2WithSessionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeNTLMv2WithSessionKey = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeNewUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeNewUser = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeNewUserWithPolicy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeNewUserWithPolicy = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeNodeNativeClearTextOK"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeNodeNativeClearTextOK = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeNodeNativeNoClearText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeNodeNativeNoClearText = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeReadSecureHash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeReadSecureHash = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSMBNTv2UserSessionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSMBNTv2UserSessionKey = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSMBWorkstationCredentialSessionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSMBWorkstationCredentialSessionKey = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSMB_LM_Key"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSMB_LM_Key = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSMB_NT_Key"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSMB_NT_Key = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSMB_NT_UserSessionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSMB_NT_UserSessionKey = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSMB_NT_WithUserSessionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSMB_NT_WithUserSessionKey = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSetGlobalPolicy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSetGlobalPolicy = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSetLMHash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSetLMHash = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSetNTHash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSetNTHash = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSetPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSetPassword = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSetPasswordAsCurrent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSetPasswordAsCurrent = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSetPolicy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSetPolicy = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSetPolicyAsCurrent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSetPolicyAsCurrent = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSetUserData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSetUserData = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSetUserName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSetUserName = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeSetWorkstationPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeSetWorkstationPassword = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeWithAuthorizationRef"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeWithAuthorizationRef = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODAuthenticationTypeWriteSecureHash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODAuthenticationTypeWriteSecureHash = ODAuthenticationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODBackOffSeconds"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODBackOffSeconds = ODErrorUserInfoKeyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODErrorDomainFramework"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODErrorDomainFramework = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODModuleConfigOptionConnectionIdleDisconnect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODModuleConfigOptionConnectionIdleDisconnect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODModuleConfigOptionConnectionSetupTimeout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODModuleConfigOptionConnectionSetupTimeout = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODModuleConfigOptionManInTheMiddle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODModuleConfigOptionManInTheMiddle = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODModuleConfigOptionPacketEncryption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODModuleConfigOptionPacketEncryption = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODModuleConfigOptionPacketSigning"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODModuleConfigOptionPacketSigning = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODModuleConfigOptionQueryTimeout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODModuleConfigOptionQueryTimeout = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODNodeOptionsQuerySkippedSubnode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODNodeOptionsQuerySkippedSubnode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeCreationTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeCreationTime = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeCurrentDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeCurrentDate = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeCurrentDayOfWeek"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeCurrentDayOfWeek = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeCurrentTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeCurrentTime = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeCurrentTimeOfDay"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeCurrentTimeOfDay = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeDaysUntilExpiration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeDaysUntilExpiration = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeEnableAtTimeOfDay"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeEnableAtTimeOfDay = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeEnableOnDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeEnableOnDate = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeEnableOnDayOfWeek"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeEnableOnDayOfWeek = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeExpiresAtTimeOfDay"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeExpiresAtTimeOfDay = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeExpiresEveryNDays"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeExpiresEveryNDays = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeExpiresOnDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeExpiresOnDate = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeExpiresOnDayOfWeek"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeExpiresOnDayOfWeek = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeFailedAuthentications"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeFailedAuthentications = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeLastAuthenticationTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeLastAuthenticationTime = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeLastFailedAuthenticationTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeLastFailedAuthenticationTime = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeLastPasswordChangeTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeLastPasswordChangeTime = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeMaximumFailedAuthentications"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeMaximumFailedAuthentications = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeNewPasswordRequiredTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeNewPasswordRequiredTime = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributePassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributePassword = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributePasswordHashes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributePasswordHashes = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributePasswordHistory"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributePasswordHistory = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributePasswordHistoryDepth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributePasswordHistoryDepth = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeRecordName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeRecordName = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyAttributeRecordType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyAttributeRecordType = ODPolicyAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyCategoryAuthentication"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyCategoryAuthentication = ODPolicyCategoryType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyCategoryPasswordChange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyCategoryPasswordChange = ODPolicyCategoryType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyCategoryPasswordContent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyCategoryPasswordContent = ODPolicyCategoryType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyKeyContent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyKeyContent = ODPolicyKeyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyKeyContentDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyKeyContentDescription = ODPolicyKeyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyKeyEvaluationDetails"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyKeyEvaluationDetails = ODPolicyKeyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyKeyIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyKeyIdentifier = ODPolicyKeyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyKeyParameters"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyKeyParameters = ODPolicyKeyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyKeyPolicySatisfied"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyKeyPolicySatisfied = ODPolicyKeyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypeAccountExpiresOnDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypeAccountExpiresOnDate = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypeAccountMaximumFailedLogins"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypeAccountMaximumFailedLogins = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypeAccountMaximumMinutesOfNonUse"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypeAccountMaximumMinutesOfNonUse = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypeAccountMaximumMinutesUntilDisabled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypeAccountMaximumMinutesUntilDisabled = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypeAccountMinutesUntilFailedLoginReset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypeAccountMinutesUntilFailedLoginReset = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypePasswordCannotBeAccountName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypePasswordCannotBeAccountName = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypePasswordChangeRequired"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypePasswordChangeRequired = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypePasswordHistory"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypePasswordHistory = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypePasswordMaximumAgeInMinutes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypePasswordMaximumAgeInMinutes = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypePasswordMaximumNumberOfCharacters"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypePasswordMaximumNumberOfCharacters = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypePasswordMinimumNumberOfCharacters"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypePasswordMinimumNumberOfCharacters = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypePasswordRequiresAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypePasswordRequiresAlpha = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypePasswordRequiresMixedCase"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypePasswordRequiresMixedCase = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypePasswordRequiresNumeric"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypePasswordRequiresNumeric = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypePasswordRequiresSymbol"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypePasswordRequiresSymbol = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODPolicyTypePasswordSelfModification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODPolicyTypePasswordSelfModification = ODPolicyType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeAFPServer"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeAFPServer = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeAliases"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeAliases = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeAttributeTypes"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeAttributeTypes = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeAugments"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeAugments = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeAutoServerSetup"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeAutoServerSetup = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeAutomount"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeAutomount = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeAutomountMap"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeAutomountMap = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeBootp"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeBootp = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeCertificateAuthorities"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeCertificateAuthorities = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeComputerGroups"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeComputerGroups = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeComputerLists"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeComputerLists = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeComputers"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeComputers = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeConfiguration"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeConfiguration = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeEthernets"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeEthernets = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeFTPServer"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeFTPServer = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeFileMakerServers"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeFileMakerServers = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeGroups"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeGroups = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeHostServices"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeHostServices = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeHosts"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeHosts = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeLDAPServer"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeLDAPServer = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeLocations"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeLocations = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeMounts"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeMounts = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeNFS"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeNFS = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeNetDomains"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeNetDomains = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeNetGroups"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeNetGroups = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeNetworks"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeNetworks = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypePeople"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypePeople = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypePresetComputerGroups"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypePresetComputerGroups = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypePresetComputerLists"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypePresetComputerLists = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypePresetComputers"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypePresetComputers = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypePresetGroups"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypePresetGroups = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypePresetUsers"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypePresetUsers = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypePrintService"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypePrintService = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypePrintServiceUser"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypePrintServiceUser = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypePrinters"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypePrinters = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeProtocols"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeProtocols = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeQTSServer"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeQTSServer = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeQueryInformation"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeQueryInformation = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeRPC"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeRPC = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeRecordTypes"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeRecordTypes = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeResources"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeResources = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeSMBServer"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeSMBServer = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeServer"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeServer = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeServices"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeServices = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeSharePoints"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeSharePoints = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeUsers"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeUsers = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODRecordTypeWebServer"); err == nil && ptr != 0 {
		// Opaque/struct symbol (e.g. CSSM_GUID): store the symbol address
		// itself. The symbol is not a pointer, so dereferencing its first
		// word as one is both wrong and trips -race checkptr alignment for
		// sub-pointer-aligned structs.
		KODRecordTypeWebServer = unsafe.Pointer(ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODSessionDefault"); err == nil && ptr != 0 {
		KODSessionDefault = objc.ValueAt[ODSessionRef](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODSessionProxyAddress"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODSessionProxyAddress = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODSessionProxyPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODSessionProxyPassword = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODSessionProxyPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODSessionProxyPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kODSessionProxyUsername"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KODSessionProxyUsername = objc.GoString(cstr)
			}
		}
	}

}
