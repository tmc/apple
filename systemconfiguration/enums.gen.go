// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

package systemconfiguration

import (
	"fmt"
)

type KSCBondStatus int

const (
	// KSCBondStatusLinkInvalid: The link state is not valid (such as down, half-duplex, or wrong speed).
	KSCBondStatusLinkInvalid KSCBondStatus = 1
	// KSCBondStatusNoPartner: The port on the switch to which the device is connected doesn’t seem to have 802.3ad Link Aggregation enabled.
	KSCBondStatusNoPartner KSCBondStatus = 2
	// KSCBondStatusNotInActiveGroup: Communication with a partner is occurring, but the link aggregation group is different from the one that is active.
	KSCBondStatusNotInActiveGroup KSCBondStatus = 3
	// KSCBondStatusOK: The status is valid (for example, enabled, active, running, and so on).
	KSCBondStatusOK KSCBondStatus = 0
	// KSCBondStatusUnknown: Nonspecific failure.
	KSCBondStatusUnknown KSCBondStatus = 999
)

func (e KSCBondStatus) String() string {
	switch e {
	case KSCBondStatusLinkInvalid:
		return "KSCBondStatusLinkInvalid"
	case KSCBondStatusNoPartner:
		return "KSCBondStatusNoPartner"
	case KSCBondStatusNotInActiveGroup:
		return "KSCBondStatusNotInActiveGroup"
	case KSCBondStatusOK:
		return "KSCBondStatusOK"
	case KSCBondStatusUnknown:
		return "KSCBondStatusUnknown"
	default:
		return fmt.Sprintf("KSCBondStatus(%d)", e)
	}
}

type KSCNetworkFlags uint

const (
	// KSCNetworkFlagsConnectionAutomatic: The specified node name or address can be reached using the current network configuration, but a connection must first be established.
	KSCNetworkFlagsConnectionAutomatic KSCNetworkFlags = 8
	// KSCNetworkFlagsConnectionRequired: The specified node name or address can be reached using the current network configuration, but a connection must first be established.
	KSCNetworkFlagsConnectionRequired KSCNetworkFlags = 4
	// KSCNetworkFlagsInterventionRequired: The specified node name or address can be reached using the current network configuration, but a connection must first be established.
	KSCNetworkFlagsInterventionRequired KSCNetworkFlags = 16
	// KSCNetworkFlagsIsDirect: Network traffic to the specified node name or address does not go through a gateway, but is routed directly to one of the interfaces in the system.
	KSCNetworkFlagsIsDirect KSCNetworkFlags = 131072
	// KSCNetworkFlagsIsLocalAddress: The specified node name or address is one associated with a network interface on the current system.
	KSCNetworkFlagsIsLocalAddress KSCNetworkFlags = 65536
	// KSCNetworkFlagsReachable: The specified node name or address can be reached using the current network configuration.
	KSCNetworkFlagsReachable KSCNetworkFlags = 2
	// KSCNetworkFlagsTransientConnection: The specified node name or address can be reached via a transient connection, such as PPP.
	KSCNetworkFlagsTransientConnection KSCNetworkFlags = 1
)

func (e KSCNetworkFlags) String() string {
	switch e {
	case KSCNetworkFlagsConnectionAutomatic:
		return "KSCNetworkFlagsConnectionAutomatic"
	case KSCNetworkFlagsConnectionRequired:
		return "KSCNetworkFlagsConnectionRequired"
	case KSCNetworkFlagsInterventionRequired:
		return "KSCNetworkFlagsInterventionRequired"
	case KSCNetworkFlagsIsDirect:
		return "KSCNetworkFlagsIsDirect"
	case KSCNetworkFlagsIsLocalAddress:
		return "KSCNetworkFlagsIsLocalAddress"
	case KSCNetworkFlagsReachable:
		return "KSCNetworkFlagsReachable"
	case KSCNetworkFlagsTransientConnection:
		return "KSCNetworkFlagsTransientConnection"
	default:
		return fmt.Sprintf("KSCNetworkFlags(%d)", e)
	}
}

type KSCStatus int

const (
	// KSCStatusAccessError: Permission is denied; you must be root to obtain a lock.
	KSCStatusAccessError KSCStatus = 1003
	// KSCStatusConnectionIgnore: Network connection information is not available at this time.
	KSCStatusConnectionIgnore KSCStatus = 4003
	// KSCStatusConnectionNoService: Network service for the connection is not available.
	KSCStatusConnectionNoService KSCStatus = 4002
	// KSCStatusFailed: A nonspecific failure occurred.
	KSCStatusFailed KSCStatus = 1001
	// KSCStatusInvalidArgument: An invalid argument was specified.
	KSCStatusInvalidArgument KSCStatus = 1002
	// KSCStatusKeyExists: # Discussion
	KSCStatusKeyExists KSCStatus = 1005
	// KSCStatusLocked: A lock is already held.
	KSCStatusLocked KSCStatus = 1006
	// KSCStatusMaxLink: The maximum link count is exceeded.
	KSCStatusMaxLink KSCStatus = 3006
	// KSCStatusNeedLock: A lock is required for this operation.
	KSCStatusNeedLock KSCStatus = 1007
	// KSCStatusNoConfigFile: The configuration file cannot be found.
	KSCStatusNoConfigFile KSCStatus = 3003
	// KSCStatusNoKey: No such key.
	KSCStatusNoKey KSCStatus = 1004
	// KSCStatusNoLink: No such link exists.
	KSCStatusNoLink KSCStatus = 3004
	// KSCStatusNoPrefsSession: The preferences session is not active.
	KSCStatusNoPrefsSession KSCStatus = 3001
	// KSCStatusNoStoreServer: The configuration daemon is not available or no longer available.
	KSCStatusNoStoreServer KSCStatus = 2002
	// KSCStatusNoStoreSession: The configuration daemon session is not active.
	KSCStatusNoStoreSession KSCStatus = 2001
	// KSCStatusNotifierActive: Notifier is currently active.
	KSCStatusNotifierActive KSCStatus = 2003
	// KSCStatusOK: The call was successful.
	KSCStatusOK KSCStatus = 0
	// KSCStatusPrefsBusy: A preferences update is currently in progress.
	KSCStatusPrefsBusy KSCStatus = 3002
	// KSCStatusReachabilityUnknown: Network reachability cannot be determined.
	KSCStatusReachabilityUnknown KSCStatus = 4001
	// KSCStatusStale: A write was attempted on a stale version of the object.
	KSCStatusStale KSCStatus = 3005
)

func (e KSCStatus) String() string {
	switch e {
	case KSCStatusAccessError:
		return "KSCStatusAccessError"
	case KSCStatusConnectionIgnore:
		return "KSCStatusConnectionIgnore"
	case KSCStatusConnectionNoService:
		return "KSCStatusConnectionNoService"
	case KSCStatusFailed:
		return "KSCStatusFailed"
	case KSCStatusInvalidArgument:
		return "KSCStatusInvalidArgument"
	case KSCStatusKeyExists:
		return "KSCStatusKeyExists"
	case KSCStatusLocked:
		return "KSCStatusLocked"
	case KSCStatusMaxLink:
		return "KSCStatusMaxLink"
	case KSCStatusNeedLock:
		return "KSCStatusNeedLock"
	case KSCStatusNoConfigFile:
		return "KSCStatusNoConfigFile"
	case KSCStatusNoKey:
		return "KSCStatusNoKey"
	case KSCStatusNoLink:
		return "KSCStatusNoLink"
	case KSCStatusNoPrefsSession:
		return "KSCStatusNoPrefsSession"
	case KSCStatusNoStoreServer:
		return "KSCStatusNoStoreServer"
	case KSCStatusNoStoreSession:
		return "KSCStatusNoStoreSession"
	case KSCStatusNotifierActive:
		return "KSCStatusNotifierActive"
	case KSCStatusOK:
		return "KSCStatusOK"
	case KSCStatusPrefsBusy:
		return "KSCStatusPrefsBusy"
	case KSCStatusReachabilityUnknown:
		return "KSCStatusReachabilityUnknown"
	case KSCStatusStale:
		return "KSCStatusStale"
	default:
		return fmt.Sprintf("KSCStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus
type SCNetworkConnectionPPPStatus int32

const (
	// KSCNetworkConnectionPPPAuthenticating: PPP is authenticating to the server (PAP, CHAP, MS-CHAP, or EAP protocols).
	KSCNetworkConnectionPPPAuthenticating SCNetworkConnectionPPPStatus = 5
	// KSCNetworkConnectionPPPConnected: PPP is now fully connected for at least one networking layer.
	KSCNetworkConnectionPPPConnected SCNetworkConnectionPPPStatus = 8
	// KSCNetworkConnectionPPPConnectingLink: PPP is connecting the lower connection layer (for example, the modem is dialing out).
	KSCNetworkConnectionPPPConnectingLink SCNetworkConnectionPPPStatus = 2
	// KSCNetworkConnectionPPPDialOnTraffic: PPP is waiting for networking traffic to automatically establish the connection.
	KSCNetworkConnectionPPPDialOnTraffic SCNetworkConnectionPPPStatus = 3
	// KSCNetworkConnectionPPPDisconnected: PPP is disconnected.
	KSCNetworkConnectionPPPDisconnected SCNetworkConnectionPPPStatus = 0
	// KSCNetworkConnectionPPPDisconnectingLink: PPP is disconnecting the lower level (for example, the modem is hanging up).
	KSCNetworkConnectionPPPDisconnectingLink SCNetworkConnectionPPPStatus = 10
	// KSCNetworkConnectionPPPHoldingLinkOff: PPP is disconnected and maintaining the link temporarily off.
	KSCNetworkConnectionPPPHoldingLinkOff SCNetworkConnectionPPPStatus = 11
	// KSCNetworkConnectionPPPInitializing: PPP is initializing.
	KSCNetworkConnectionPPPInitializing SCNetworkConnectionPPPStatus = 1
	// KSCNetworkConnectionPPPNegotiatingLink: The PPP lower layer is connected and PPP is negotiating the link layer (LCP protocol).
	KSCNetworkConnectionPPPNegotiatingLink SCNetworkConnectionPPPStatus = 4
	// KSCNetworkConnectionPPPNegotiatingNetwork: PPP is now authenticated and negotiating the networking layer (IPCP or IPv6CP protocols).
	KSCNetworkConnectionPPPNegotiatingNetwork SCNetworkConnectionPPPStatus = 7
	// KSCNetworkConnectionPPPSuspended: PPP is suspended as a result of the suspend command (for example, when a V.92 Modem is On Hold).
	KSCNetworkConnectionPPPSuspended SCNetworkConnectionPPPStatus = 12
	// KSCNetworkConnectionPPPTerminating: PPP networking and link protocols are terminating.
	KSCNetworkConnectionPPPTerminating SCNetworkConnectionPPPStatus = 9
	// KSCNetworkConnectionPPPWaitingForCallBack: PPP is waiting for the server to call back.
	KSCNetworkConnectionPPPWaitingForCallBack SCNetworkConnectionPPPStatus = 6
	// KSCNetworkConnectionPPPWaitingForRedial: PPP has found a busy server and is waiting for redial.
	KSCNetworkConnectionPPPWaitingForRedial SCNetworkConnectionPPPStatus = 13
)

func (e SCNetworkConnectionPPPStatus) String() string {
	switch e {
	case KSCNetworkConnectionPPPAuthenticating:
		return "KSCNetworkConnectionPPPAuthenticating"
	case KSCNetworkConnectionPPPConnected:
		return "KSCNetworkConnectionPPPConnected"
	case KSCNetworkConnectionPPPConnectingLink:
		return "KSCNetworkConnectionPPPConnectingLink"
	case KSCNetworkConnectionPPPDialOnTraffic:
		return "KSCNetworkConnectionPPPDialOnTraffic"
	case KSCNetworkConnectionPPPDisconnected:
		return "KSCNetworkConnectionPPPDisconnected"
	case KSCNetworkConnectionPPPDisconnectingLink:
		return "KSCNetworkConnectionPPPDisconnectingLink"
	case KSCNetworkConnectionPPPHoldingLinkOff:
		return "KSCNetworkConnectionPPPHoldingLinkOff"
	case KSCNetworkConnectionPPPInitializing:
		return "KSCNetworkConnectionPPPInitializing"
	case KSCNetworkConnectionPPPNegotiatingLink:
		return "KSCNetworkConnectionPPPNegotiatingLink"
	case KSCNetworkConnectionPPPNegotiatingNetwork:
		return "KSCNetworkConnectionPPPNegotiatingNetwork"
	case KSCNetworkConnectionPPPSuspended:
		return "KSCNetworkConnectionPPPSuspended"
	case KSCNetworkConnectionPPPTerminating:
		return "KSCNetworkConnectionPPPTerminating"
	case KSCNetworkConnectionPPPWaitingForCallBack:
		return "KSCNetworkConnectionPPPWaitingForCallBack"
	case KSCNetworkConnectionPPPWaitingForRedial:
		return "KSCNetworkConnectionPPPWaitingForRedial"
	default:
		return fmt.Sprintf("SCNetworkConnectionPPPStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStatus
type SCNetworkConnectionStatus int32

const (
	// KSCNetworkConnectionConnected: The network connection is connected.
	KSCNetworkConnectionConnected SCNetworkConnectionStatus = 2
	// KSCNetworkConnectionConnecting: The network connection is connecting.
	KSCNetworkConnectionConnecting SCNetworkConnectionStatus = 1
	// KSCNetworkConnectionDisconnected: The network connection is disconnected.
	KSCNetworkConnectionDisconnected SCNetworkConnectionStatus = 0
	// KSCNetworkConnectionDisconnecting: The network connection is disconnecting.
	KSCNetworkConnectionDisconnecting SCNetworkConnectionStatus = 3
	// KSCNetworkConnectionInvalid: The network connection refers to an invalid service.
	KSCNetworkConnectionInvalid SCNetworkConnectionStatus = -1
)

func (e SCNetworkConnectionStatus) String() string {
	switch e {
	case KSCNetworkConnectionConnected:
		return "KSCNetworkConnectionConnected"
	case KSCNetworkConnectionConnecting:
		return "KSCNetworkConnectionConnecting"
	case KSCNetworkConnectionDisconnected:
		return "KSCNetworkConnectionDisconnected"
	case KSCNetworkConnectionDisconnecting:
		return "KSCNetworkConnectionDisconnecting"
	case KSCNetworkConnectionInvalid:
		return "KSCNetworkConnectionInvalid"
	default:
		return fmt.Sprintf("SCNetworkConnectionStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags
type SCNetworkReachabilityFlags uint32

const (
	// KSCNetworkReachabilityFlagsInterventionRequired: The specified node name or address can be reached using the current network configuration, but a connection must first be established.
	KSCNetworkReachabilityFlagsInterventionRequired SCNetworkReachabilityFlags = 16
	// KSCNetworkReachabilityFlagsIsDirect: Network traffic to the specified node name or address will not go through a gateway, but is routed directly to one of the interfaces in the system.
	KSCNetworkReachabilityFlagsIsDirect SCNetworkReachabilityFlags = 131072
	// KSCNetworkReachabilityFlagsIsLocalAddress: The specified node name or address is one that is associated with a network interface on the current system.
	KSCNetworkReachabilityFlagsIsLocalAddress SCNetworkReachabilityFlags = 65536
	// KSCNetworkReachabilityFlagsIsWWAN: The specified node name or address can be reached via a cellular connection, such as EDGE or GPRS.
	KSCNetworkReachabilityFlagsIsWWAN SCNetworkReachabilityFlags = 131073
	// KSCNetworkReachabilityFlagsReachable: The specified node name or address can be reached using the current network configuration.
	KSCNetworkReachabilityFlagsReachable SCNetworkReachabilityFlags = 2
	// KSCNetworkReachabilityFlagsTransientConnection: The specified node name or address can be reached via a transient connection, such as PPP.
	KSCNetworkReachabilityFlagsTransientConnection SCNetworkReachabilityFlags = 1
)

func (e SCNetworkReachabilityFlags) String() string {
	switch e {
	case KSCNetworkReachabilityFlagsInterventionRequired:
		return "KSCNetworkReachabilityFlagsInterventionRequired"
	case KSCNetworkReachabilityFlagsIsDirect:
		return "KSCNetworkReachabilityFlagsIsDirect"
	case KSCNetworkReachabilityFlagsIsLocalAddress:
		return "KSCNetworkReachabilityFlagsIsLocalAddress"
	case KSCNetworkReachabilityFlagsIsWWAN:
		return "KSCNetworkReachabilityFlagsIsWWAN"
	case KSCNetworkReachabilityFlagsReachable:
		return "KSCNetworkReachabilityFlagsReachable"
	case KSCNetworkReachabilityFlagsTransientConnection:
		return "KSCNetworkReachabilityFlagsTransientConnection"
	default:
		return fmt.Sprintf("SCNetworkReachabilityFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesNotification
type SCPreferencesNotification uint32

const ()
