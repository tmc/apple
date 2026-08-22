// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"fmt"
)

type KODExpirationTime int32

const (
	KODExpirationTimeExpired      KODExpirationTime = 0
	KODExpirationTimeNeverExpires KODExpirationTime = -1
)

func (e KODExpirationTime) String() string {
	switch e {
	case KODExpirationTimeExpired:
		return "KODExpirationTimeExpired"
	case KODExpirationTimeNeverExpires:
		return "KODExpirationTimeNeverExpires"
	default:
		return fmt.Sprintf("KODExpirationTime(%d)", e)
	}
}

type KODMatch uint32

const (
	KODMatchAny         KODMatch = 0x1
	KODMatchBeginsWith  KODMatch = 0x2002
	KODMatchContains    KODMatch = 0x2004
	KODMatchEndsWith    KODMatch = 0x2003
	KODMatchEqualTo     KODMatch = 0x2001
	KODMatchGreaterThan KODMatch = 0x2006
	// Deprecated.
	KODMatchInsensitiveBeginsWith KODMatch = 0x2102
	// Deprecated.
	KODMatchInsensitiveContains KODMatch = 0x2104
)

func (e KODMatch) String() string {
	switch e {
	case KODMatchAny:
		return "KODMatchAny"
	case KODMatchBeginsWith:
		return "KODMatchBeginsWith"
	case KODMatchContains:
		return "KODMatchContains"
	case KODMatchEndsWith:
		return "KODMatchEndsWith"
	case KODMatchEqualTo:
		return "KODMatchEqualTo"
	case KODMatchGreaterThan:
		return "KODMatchGreaterThan"
	case KODMatchInsensitiveBeginsWith:
		return "KODMatchInsensitiveBeginsWith"
	case KODMatchInsensitiveContains:
		return "KODMatchInsensitiveContains"
	default:
		return fmt.Sprintf("KODMatch(%d)", e)
	}
}

type KODNodeType uint32

const (
	// KODNodeTypeAuthentication: A node used for authentication or record lookups.
	KODNodeTypeAuthentication KODNodeType = 0x2201
	// KODNodeTypeConfigure: A node that specifically refers to the Directory Services configuration.
	KODNodeTypeConfigure KODNodeType = 0x2202
	// KODNodeTypeContacts: A node used for applications that handle contact data.
	KODNodeTypeContacts KODNodeType = 0x2204
	// KODNodeTypeLocalNodes: A node that specifically looks at the local directory.
	KODNodeTypeLocalNodes KODNodeType = 0x2200
	// KODNodeTypeNetwork: A node used for looking up network resource type data.
	KODNodeTypeNetwork KODNodeType = 0x2205
)

func (e KODNodeType) String() string {
	switch e {
	case KODNodeTypeAuthentication:
		return "KODNodeTypeAuthentication"
	case KODNodeTypeConfigure:
		return "KODNodeTypeConfigure"
	case KODNodeTypeContacts:
		return "KODNodeTypeContacts"
	case KODNodeTypeLocalNodes:
		return "KODNodeTypeLocalNodes"
	case KODNodeTypeNetwork:
		return "KODNodeTypeNetwork"
	default:
		return fmt.Sprintf("KODNodeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODFrameworkErrors
type ODFrameworkErrors uint32

const (
	// KODErrorCredentialsAccountDisabled: The account is disabled.
	KODErrorCredentialsAccountDisabled ODFrameworkErrors = 5301
	// KODErrorCredentialsAccountExpired: The account is expired.
	KODErrorCredentialsAccountExpired ODFrameworkErrors = 5302
	// KODErrorCredentialsAccountInactive: The account is inactive.
	KODErrorCredentialsAccountInactive ODFrameworkErrors = 5303
	KODErrorCredentialsAccountLocked   ODFrameworkErrors = 5305
	// KODErrorCredentialsAccountNotFound: The authentication server could not find the provided account.
	KODErrorCredentialsAccountNotFound          ODFrameworkErrors = 5300
	KODErrorCredentialsAccountTemporarilyLocked ODFrameworkErrors = 5304
	// KODErrorCredentialsContactMaster: The authentication server contacted is not the primary server, and the requested operation requires the primary server.
	KODErrorCredentialsContactMaster  ODFrameworkErrors = 5204
	KODErrorCredentialsContactPrimary ODFrameworkErrors = 5204
	// KODErrorCredentialsInvalid: The provided credentials are invalid with the current node.
	KODErrorCredentialsInvalid ODFrameworkErrors = 5000
	// KODErrorCredentialsInvalidComputer: The account is not permitted to log into this computer.
	KODErrorCredentialsInvalidComputer ODFrameworkErrors = 5501
	// KODErrorCredentialsInvalidLogonHours: The logon attempt was not within set logon hours.
	KODErrorCredentialsInvalidLogonHours ODFrameworkErrors = 5500
	// KODErrorCredentialsMethodNotSupported: The extended authentication method is not supported.
	KODErrorCredentialsMethodNotSupported ODFrameworkErrors = 5100
	// KODErrorCredentialsNotAuthorized: The operation, such as changing a password, is not permitted with current privileges.
	KODErrorCredentialsNotAuthorized ODFrameworkErrors = 5101
	// KODErrorCredentialsOperationFailed: The requested operation failed.
	KODErrorCredentialsOperationFailed ODFrameworkErrors = 5103
	// KODErrorCredentialsParameterError: An invalid parameter was provided.
	KODErrorCredentialsParameterError ODFrameworkErrors = 5102
	// KODErrorCredentialsPasswordChangeRequired: The password must be changed.
	KODErrorCredentialsPasswordChangeRequired ODFrameworkErrors = 5401
	// KODErrorCredentialsPasswordChangeTooSoon: The password was changed too recently to be changed again.
	KODErrorCredentialsPasswordChangeTooSoon ODFrameworkErrors = 5407
	// KODErrorCredentialsPasswordExpired: The password has expired and must be changed.
	KODErrorCredentialsPasswordExpired ODFrameworkErrors = 5400
	// KODErrorCredentialsPasswordNeedsDigit: The provided password needs at least one digit.
	KODErrorCredentialsPasswordNeedsDigit ODFrameworkErrors = 5406
	// KODErrorCredentialsPasswordNeedsLetter: The provided password needs at least one letter.
	KODErrorCredentialsPasswordNeedsLetter ODFrameworkErrors = 5405
	// KODErrorCredentialsPasswordQualityFailed: The provided password did not meet minimum quality requirements.
	KODErrorCredentialsPasswordQualityFailed ODFrameworkErrors = 5402
	// KODErrorCredentialsPasswordTooLong: The provided password is too long.
	KODErrorCredentialsPasswordTooLong ODFrameworkErrors = 5404
	// KODErrorCredentialsPasswordTooShort: The provided password is too short.
	KODErrorCredentialsPasswordTooShort ODFrameworkErrors = 5403
	// KODErrorCredentialsPasswordUnrecoverable: The password could not be recovered from the authentication database.
	KODErrorCredentialsPasswordUnrecoverable ODFrameworkErrors = 5408
	// KODErrorCredentialsServerCommunicationError: The authentication server encountered a communication error.
	KODErrorCredentialsServerCommunicationError ODFrameworkErrors = 5205
	// KODErrorCredentialsServerError: The authentication server encountered an error.
	KODErrorCredentialsServerError ODFrameworkErrors = 5202
	// KODErrorCredentialsServerNotFound: The authentication server could not be found.
	KODErrorCredentialsServerNotFound ODFrameworkErrors = 5201
	// KODErrorCredentialsServerTimeout: The authentication server timed out.
	KODErrorCredentialsServerTimeout ODFrameworkErrors = 5203
	// KODErrorCredentialsServerUnreachable: The authentication server could not be reached.
	KODErrorCredentialsServerUnreachable ODFrameworkErrors = 5200
	// KODErrorDaemonError: The daemon has encountered an undefined error.
	KODErrorDaemonError ODFrameworkErrors = 10002
	// KODErrorNodeConnectionFailed: The node connection failed.
	KODErrorNodeConnectionFailed ODFrameworkErrors = 2100
	KODErrorNodeDisabled         ODFrameworkErrors = 2002
	// KODErrorNodeUnknownHost: The host provided is invalid.
	KODErrorNodeUnknownHost ODFrameworkErrors = 2200
	// KODErrorNodeUnknownName: The node name provided does not exist and cannot be opened.
	KODErrorNodeUnknownName ODFrameworkErrors = 2000
	// KODErrorNodeUnknownType: The node type provided is not a known value.
	KODErrorNodeUnknownType ODFrameworkErrors = 2001
	// KODErrorPluginError: A plug-in has encountered an undefined error.
	KODErrorPluginError ODFrameworkErrors = 10001
	// KODErrorPluginOperationNotSupported: The plug-in does not support the requested operation.
	KODErrorPluginOperationNotSupported ODFrameworkErrors = 10000
	KODErrorPluginOperationTimeout      ODFrameworkErrors = 10003
	KODErrorPolicyOutOfRange            ODFrameworkErrors = 6001
	KODErrorPolicyUnsupported           ODFrameworkErrors = 6000
	// KODErrorQueryInvalidMatchType: An invalid match type was provided in the query.
	KODErrorQueryInvalidMatchType ODFrameworkErrors = 3100
	// KODErrorQuerySynchronize: A query synchronization has been initiated.
	KODErrorQuerySynchronize ODFrameworkErrors = 3000
	// KODErrorQueryTimeout: The query timed out.
	KODErrorQueryTimeout ODFrameworkErrors = 3102
	// KODErrorQueryUnsupportedMatchType: An unsupported match type was provided in the query.
	KODErrorQueryUnsupportedMatchType ODFrameworkErrors = 3101
	// KODErrorRecordAlreadyExists: The record create failed because the record already exists.
	KODErrorRecordAlreadyExists ODFrameworkErrors = 4102
	// KODErrorRecordAttributeNotFound: The requested attribute could not be found in the record.
	KODErrorRecordAttributeNotFound ODFrameworkErrors = 4201
	// KODErrorRecordAttributeUnknownType: The attribute type is unknown.
	KODErrorRecordAttributeUnknownType ODFrameworkErrors = 4200
	// KODErrorRecordAttributeValueNotFound: The requested attribute value could not be found in the record.
	KODErrorRecordAttributeValueNotFound ODFrameworkErrors = 4203
	// KODErrorRecordAttributeValueSchemaError: The attribute value does not meet schema requirements.
	KODErrorRecordAttributeValueSchemaError ODFrameworkErrors = 4202
	KODErrorRecordInvalidType               ODFrameworkErrors = 4101
	KODErrorRecordNoLongerExists            ODFrameworkErrors = 4104
	// KODErrorRecordParameterError: An invalid parameter was provided.
	KODErrorRecordParameterError ODFrameworkErrors = 4100
	// KODErrorRecordPermissionError: The changes were denied due to insufficient permissions.
	KODErrorRecordPermissionError ODFrameworkErrors = 4001
	// KODErrorRecordReadOnlyNode: The record cannot be modified.
	KODErrorRecordReadOnlyNode ODFrameworkErrors = 4000
	// KODErrorRecordTypeDisabled: The record type is disabled by policy for a plug-in.
	KODErrorRecordTypeDisabled ODFrameworkErrors = 4103
	// KODErrorSessionDaemonNotRunning: The daemon is not running.
	KODErrorSessionDaemonNotRunning ODFrameworkErrors = 1002
	// KODErrorSessionDaemonRefused: The daemon refused the session.
	KODErrorSessionDaemonRefused ODFrameworkErrors = 1003
	// KODErrorSessionLocalOnlyDaemonInUse: A normal request was issued when the local-only daemon was in use.
	KODErrorSessionLocalOnlyDaemonInUse ODFrameworkErrors = 1000
	// KODErrorSessionNormalDaemonInUse: A local-only request was issued when the normal daemon was in use.
	KODErrorSessionNormalDaemonInUse ODFrameworkErrors = 1001
	// KODErrorSessionProxyCommunicationError: There was a communication error with the remote daemon.
	KODErrorSessionProxyCommunicationError ODFrameworkErrors = 1100
	// KODErrorSessionProxyIPUnreachable: The proxy did not respond.
	KODErrorSessionProxyIPUnreachable ODFrameworkErrors = 1102
	// KODErrorSessionProxyUnknownHost: The proxy could not be resolved.
	KODErrorSessionProxyUnknownHost ODFrameworkErrors = 1103
	// KODErrorSessionProxyVersionMismatch: Versions mismatch between the remote daemon and the local framework.
	KODErrorSessionProxyVersionMismatch ODFrameworkErrors = 1101
	KODErrorSuccess                     ODFrameworkErrors = 0
)

func (e ODFrameworkErrors) String() string {
	switch e {
	case KODErrorCredentialsAccountDisabled:
		return "KODErrorCredentialsAccountDisabled"
	case KODErrorCredentialsAccountExpired:
		return "KODErrorCredentialsAccountExpired"
	case KODErrorCredentialsAccountInactive:
		return "KODErrorCredentialsAccountInactive"
	case KODErrorCredentialsAccountLocked:
		return "KODErrorCredentialsAccountLocked"
	case KODErrorCredentialsAccountNotFound:
		return "KODErrorCredentialsAccountNotFound"
	case KODErrorCredentialsAccountTemporarilyLocked:
		return "KODErrorCredentialsAccountTemporarilyLocked"
	case KODErrorCredentialsContactMaster:
		return "KODErrorCredentialsContactMaster"
	case KODErrorCredentialsInvalid:
		return "KODErrorCredentialsInvalid"
	case KODErrorCredentialsInvalidComputer:
		return "KODErrorCredentialsInvalidComputer"
	case KODErrorCredentialsInvalidLogonHours:
		return "KODErrorCredentialsInvalidLogonHours"
	case KODErrorCredentialsMethodNotSupported:
		return "KODErrorCredentialsMethodNotSupported"
	case KODErrorCredentialsNotAuthorized:
		return "KODErrorCredentialsNotAuthorized"
	case KODErrorCredentialsOperationFailed:
		return "KODErrorCredentialsOperationFailed"
	case KODErrorCredentialsParameterError:
		return "KODErrorCredentialsParameterError"
	case KODErrorCredentialsPasswordChangeRequired:
		return "KODErrorCredentialsPasswordChangeRequired"
	case KODErrorCredentialsPasswordChangeTooSoon:
		return "KODErrorCredentialsPasswordChangeTooSoon"
	case KODErrorCredentialsPasswordExpired:
		return "KODErrorCredentialsPasswordExpired"
	case KODErrorCredentialsPasswordNeedsDigit:
		return "KODErrorCredentialsPasswordNeedsDigit"
	case KODErrorCredentialsPasswordNeedsLetter:
		return "KODErrorCredentialsPasswordNeedsLetter"
	case KODErrorCredentialsPasswordQualityFailed:
		return "KODErrorCredentialsPasswordQualityFailed"
	case KODErrorCredentialsPasswordTooLong:
		return "KODErrorCredentialsPasswordTooLong"
	case KODErrorCredentialsPasswordTooShort:
		return "KODErrorCredentialsPasswordTooShort"
	case KODErrorCredentialsPasswordUnrecoverable:
		return "KODErrorCredentialsPasswordUnrecoverable"
	case KODErrorCredentialsServerCommunicationError:
		return "KODErrorCredentialsServerCommunicationError"
	case KODErrorCredentialsServerError:
		return "KODErrorCredentialsServerError"
	case KODErrorCredentialsServerNotFound:
		return "KODErrorCredentialsServerNotFound"
	case KODErrorCredentialsServerTimeout:
		return "KODErrorCredentialsServerTimeout"
	case KODErrorCredentialsServerUnreachable:
		return "KODErrorCredentialsServerUnreachable"
	case KODErrorDaemonError:
		return "KODErrorDaemonError"
	case KODErrorNodeConnectionFailed:
		return "KODErrorNodeConnectionFailed"
	case KODErrorNodeDisabled:
		return "KODErrorNodeDisabled"
	case KODErrorNodeUnknownHost:
		return "KODErrorNodeUnknownHost"
	case KODErrorNodeUnknownName:
		return "KODErrorNodeUnknownName"
	case KODErrorNodeUnknownType:
		return "KODErrorNodeUnknownType"
	case KODErrorPluginError:
		return "KODErrorPluginError"
	case KODErrorPluginOperationNotSupported:
		return "KODErrorPluginOperationNotSupported"
	case KODErrorPluginOperationTimeout:
		return "KODErrorPluginOperationTimeout"
	case KODErrorPolicyOutOfRange:
		return "KODErrorPolicyOutOfRange"
	case KODErrorPolicyUnsupported:
		return "KODErrorPolicyUnsupported"
	case KODErrorQueryInvalidMatchType:
		return "KODErrorQueryInvalidMatchType"
	case KODErrorQuerySynchronize:
		return "KODErrorQuerySynchronize"
	case KODErrorQueryTimeout:
		return "KODErrorQueryTimeout"
	case KODErrorQueryUnsupportedMatchType:
		return "KODErrorQueryUnsupportedMatchType"
	case KODErrorRecordAlreadyExists:
		return "KODErrorRecordAlreadyExists"
	case KODErrorRecordAttributeNotFound:
		return "KODErrorRecordAttributeNotFound"
	case KODErrorRecordAttributeUnknownType:
		return "KODErrorRecordAttributeUnknownType"
	case KODErrorRecordAttributeValueNotFound:
		return "KODErrorRecordAttributeValueNotFound"
	case KODErrorRecordAttributeValueSchemaError:
		return "KODErrorRecordAttributeValueSchemaError"
	case KODErrorRecordInvalidType:
		return "KODErrorRecordInvalidType"
	case KODErrorRecordNoLongerExists:
		return "KODErrorRecordNoLongerExists"
	case KODErrorRecordParameterError:
		return "KODErrorRecordParameterError"
	case KODErrorRecordPermissionError:
		return "KODErrorRecordPermissionError"
	case KODErrorRecordReadOnlyNode:
		return "KODErrorRecordReadOnlyNode"
	case KODErrorRecordTypeDisabled:
		return "KODErrorRecordTypeDisabled"
	case KODErrorSessionDaemonNotRunning:
		return "KODErrorSessionDaemonNotRunning"
	case KODErrorSessionDaemonRefused:
		return "KODErrorSessionDaemonRefused"
	case KODErrorSessionLocalOnlyDaemonInUse:
		return "KODErrorSessionLocalOnlyDaemonInUse"
	case KODErrorSessionNormalDaemonInUse:
		return "KODErrorSessionNormalDaemonInUse"
	case KODErrorSessionProxyCommunicationError:
		return "KODErrorSessionProxyCommunicationError"
	case KODErrorSessionProxyIPUnreachable:
		return "KODErrorSessionProxyIPUnreachable"
	case KODErrorSessionProxyUnknownHost:
		return "KODErrorSessionProxyUnknownHost"
	case KODErrorSessionProxyVersionMismatch:
		return "KODErrorSessionProxyVersionMismatch"
	case KODErrorSuccess:
		return "KODErrorSuccess"
	default:
		return fmt.Sprintf("ODFrameworkErrors(%d)", e)
	}
}

type ODPacketEncryption uint32

const (
	ODPacketEncryptionAllow    ODPacketEncryption = 1
	ODPacketEncryptionDisabled ODPacketEncryption = 0
	ODPacketEncryptionRequired ODPacketEncryption = 2
	ODPacketEncryptionSSL      ODPacketEncryption = 3
)

func (e ODPacketEncryption) String() string {
	switch e {
	case ODPacketEncryptionAllow:
		return "ODPacketEncryptionAllow"
	case ODPacketEncryptionDisabled:
		return "ODPacketEncryptionDisabled"
	case ODPacketEncryptionRequired:
		return "ODPacketEncryptionRequired"
	case ODPacketEncryptionSSL:
		return "ODPacketEncryptionSSL"
	default:
		return fmt.Sprintf("ODPacketEncryption(%d)", e)
	}
}

type ODPacketSigning uint32

const (
	ODPacketSigningAllow    ODPacketSigning = 1
	ODPacketSigningDisabled ODPacketSigning = 0
	ODPacketSigningRequired ODPacketSigning = 2
)

func (e ODPacketSigning) String() string {
	switch e {
	case ODPacketSigningAllow:
		return "ODPacketSigningAllow"
	case ODPacketSigningDisabled:
		return "ODPacketSigningDisabled"
	case ODPacketSigningRequired:
		return "ODPacketSigningRequired"
	default:
		return fmt.Sprintf("ODPacketSigning(%d)", e)
	}
}
