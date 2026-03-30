// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct/Code
type NEAppProxyFlowError int

const (
	// NEAppProxyFlowErrorAborted: The flow was aborted.
	NEAppProxyFlowErrorAborted NEAppProxyFlowError = 5
	// NEAppProxyFlowErrorDatagramTooLarge: A caller attempted to write a datagram that was larger than the socket’s receive window.
	NEAppProxyFlowErrorDatagramTooLarge NEAppProxyFlowError = 9
	// NEAppProxyFlowErrorHostUnreachable: An attempt to reach the remote endpoint of the flow failed.
	NEAppProxyFlowErrorHostUnreachable NEAppProxyFlowError = 3
	// NEAppProxyFlowErrorInternal: An internal error occurred while handling the flow.
	NEAppProxyFlowErrorInternal NEAppProxyFlowError = 8
	// NEAppProxyFlowErrorInvalidArgument: A proxy flow method received an invalid argument.
	NEAppProxyFlowErrorInvalidArgument NEAppProxyFlowError = 4
	// NEAppProxyFlowErrorNotConnected: The flow is not fully opened.
	NEAppProxyFlowErrorNotConnected NEAppProxyFlowError = 1
	// NEAppProxyFlowErrorPeerReset: The remote peer closed the flow.
	NEAppProxyFlowErrorPeerReset NEAppProxyFlowError = 2
	// NEAppProxyFlowErrorReadAlreadyPending: A read operation on the flow is already pending.
	NEAppProxyFlowErrorReadAlreadyPending NEAppProxyFlowError = 10
	// NEAppProxyFlowErrorRefused: Connecting the flow to its remote endpoint failed.
	NEAppProxyFlowErrorRefused NEAppProxyFlowError = 6
	// NEAppProxyFlowErrorTimedOut: The flow timed out.
	NEAppProxyFlowErrorTimedOut NEAppProxyFlowError = 7
)

func (e NEAppProxyFlowError) String() string {
	switch e {
	case NEAppProxyFlowErrorAborted:
		return "NEAppProxyFlowErrorAborted"
	case NEAppProxyFlowErrorDatagramTooLarge:
		return "NEAppProxyFlowErrorDatagramTooLarge"
	case NEAppProxyFlowErrorHostUnreachable:
		return "NEAppProxyFlowErrorHostUnreachable"
	case NEAppProxyFlowErrorInternal:
		return "NEAppProxyFlowErrorInternal"
	case NEAppProxyFlowErrorInvalidArgument:
		return "NEAppProxyFlowErrorInvalidArgument"
	case NEAppProxyFlowErrorNotConnected:
		return "NEAppProxyFlowErrorNotConnected"
	case NEAppProxyFlowErrorPeerReset:
		return "NEAppProxyFlowErrorPeerReset"
	case NEAppProxyFlowErrorReadAlreadyPending:
		return "NEAppProxyFlowErrorReadAlreadyPending"
	case NEAppProxyFlowErrorRefused:
		return "NEAppProxyFlowErrorRefused"
	case NEAppProxyFlowErrorTimedOut:
		return "NEAppProxyFlowErrorTimedOut"
	default:
		return fmt.Sprintf("NEAppProxyFlowError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManagerError-swift.struct/Code
type NEAppPushManagerError int

const (
	// NEAppPushManagerErrorConfigurationInvalid: An error code that indicates the app push configuration is invalid.
	NEAppPushManagerErrorConfigurationInvalid NEAppPushManagerError = 1
	// NEAppPushManagerErrorConfigurationNotLoaded: An error code that indicates the manager hasn’t loaded the app push configuration.
	NEAppPushManagerErrorConfigurationNotLoaded NEAppPushManagerError = 2
	// NEAppPushManagerErrorInactiveSession: An error code that indicates an invalid attempt to perform an operation on an inactive session.
	NEAppPushManagerErrorInactiveSession NEAppPushManagerError = 4
	// NEAppPushManagerErrorInternalError: An error code that indicates an internal error in the local push connectivity framework.
	NEAppPushManagerErrorInternalError NEAppPushManagerError = 3
)

func (e NEAppPushManagerError) String() string {
	switch e {
	case NEAppPushManagerErrorConfigurationInvalid:
		return "NEAppPushManagerErrorConfigurationInvalid"
	case NEAppPushManagerErrorConfigurationNotLoaded:
		return "NEAppPushManagerErrorConfigurationNotLoaded"
	case NEAppPushManagerErrorInactiveSession:
		return "NEAppPushManagerErrorInactiveSession"
	case NEAppPushManagerErrorInternalError:
		return "NEAppPushManagerErrorInternalError"
	default:
		return fmt.Sprintf("NEAppPushManagerError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProtocol
type NEDNSProtocol int

const (
	// NEDNSProtocolCleartext: The DNS server uses cleartext UDP or TCP over port 53.
	NEDNSProtocolCleartext NEDNSProtocol = 1
	// NEDNSProtocolHTTPS: The DNS server uses DNS-over-HTTPS.
	NEDNSProtocolHTTPS NEDNSProtocol = 3
	// NEDNSProtocolTLS: The DNS server uses DNS-over-TLS.
	NEDNSProtocolTLS NEDNSProtocol = 2
)

func (e NEDNSProtocol) String() string {
	switch e {
	case NEDNSProtocolCleartext:
		return "NEDNSProtocolCleartext"
	case NEDNSProtocolHTTPS:
		return "NEDNSProtocolHTTPS"
	case NEDNSProtocolTLS:
		return "NEDNSProtocolTLS"
	default:
		return fmt.Sprintf("NEDNSProtocol(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManagerError
type NEDNSProxyManagerError int

const (
	// NEDNSProxyManagerErrorConfigurationCannotBeRemoved: Unremovable DNS proxy configuration.
	NEDNSProxyManagerErrorConfigurationCannotBeRemoved NEDNSProxyManagerError = 4
	// NEDNSProxyManagerErrorConfigurationDisabled: Disabled DNS proxy configuration.
	NEDNSProxyManagerErrorConfigurationDisabled NEDNSProxyManagerError = 2
	// NEDNSProxyManagerErrorConfigurationInvalid: Invalid DNS proxy configuration that cannot be stored.
	NEDNSProxyManagerErrorConfigurationInvalid NEDNSProxyManagerError = 1
	// NEDNSProxyManagerErrorConfigurationStale: Outdated DNS proxy configuration that needs to be loaded.
	NEDNSProxyManagerErrorConfigurationStale NEDNSProxyManagerError = 3
)

func (e NEDNSProxyManagerError) String() string {
	switch e {
	case NEDNSProxyManagerErrorConfigurationCannotBeRemoved:
		return "NEDNSProxyManagerErrorConfigurationCannotBeRemoved"
	case NEDNSProxyManagerErrorConfigurationDisabled:
		return "NEDNSProxyManagerErrorConfigurationDisabled"
	case NEDNSProxyManagerErrorConfigurationInvalid:
		return "NEDNSProxyManagerErrorConfigurationInvalid"
	case NEDNSProxyManagerErrorConfigurationStale:
		return "NEDNSProxyManagerErrorConfigurationStale"
	default:
		return fmt.Sprintf("NEDNSProxyManagerError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManagerError
type NEDNSSettingsManagerError int

const (
	// NEDNSSettingsManagerErrorConfigurationCannotBeRemoved: An error code that indicates removing the DNS settings manager failed.
	NEDNSSettingsManagerErrorConfigurationCannotBeRemoved NEDNSSettingsManagerError = 4
	// NEDNSSettingsManagerErrorConfigurationDisabled: An error code that indicates the DNS settings manager isn’t enabled.
	NEDNSSettingsManagerErrorConfigurationDisabled NEDNSSettingsManagerError = 2
	// NEDNSSettingsManagerErrorConfigurationInvalid: An error code that indicates the DNS settings manager is invalid.
	NEDNSSettingsManagerErrorConfigurationInvalid NEDNSSettingsManagerError = 1
	// NEDNSSettingsManagerErrorConfigurationStale: An error code that indicates the DNS settings manager isn’t loaded.
	NEDNSSettingsManagerErrorConfigurationStale NEDNSSettingsManagerError = 3
)

func (e NEDNSSettingsManagerError) String() string {
	switch e {
	case NEDNSSettingsManagerErrorConfigurationCannotBeRemoved:
		return "NEDNSSettingsManagerErrorConfigurationCannotBeRemoved"
	case NEDNSSettingsManagerErrorConfigurationDisabled:
		return "NEDNSSettingsManagerErrorConfigurationDisabled"
	case NEDNSSettingsManagerErrorConfigurationInvalid:
		return "NEDNSSettingsManagerErrorConfigurationInvalid"
	case NEDNSSettingsManagerErrorConfigurationStale:
		return "NEDNSSettingsManagerErrorConfigurationStale"
	default:
		return fmt.Sprintf("NEDNSSettingsManagerError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRuleAction
type NEEvaluateConnectionRuleAction int

const (
	// NEEvaluateConnectionRuleActionConnectIfNeeded: Start the VPN if connections to the matching hostname cannot be resolved.
	NEEvaluateConnectionRuleActionConnectIfNeeded NEEvaluateConnectionRuleAction = 1
	// NEEvaluateConnectionRuleActionNeverConnect: Do not start the VPN.
	NEEvaluateConnectionRuleActionNeverConnect NEEvaluateConnectionRuleAction = 2
)

func (e NEEvaluateConnectionRuleAction) String() string {
	switch e {
	case NEEvaluateConnectionRuleActionConnectIfNeeded:
		return "NEEvaluateConnectionRuleActionConnectIfNeeded"
	case NEEvaluateConnectionRuleActionNeverConnect:
		return "NEEvaluateConnectionRuleActionNeverConnect"
	default:
		return fmt.Sprintf("NEEvaluateConnectionRuleAction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterAction
type NEFilterAction int

const (
	// NEFilterActionAllow: Allow the flow.
	NEFilterActionAllow NEFilterAction = 1
	// NEFilterActionDrop: Drop the flow.
	NEFilterActionDrop NEFilterAction = 2
	// NEFilterActionFilterData: Filter data on the flow.
	NEFilterActionFilterData NEFilterAction = 4
	// NEFilterActionInvalid: Invalid action used to represent an error.
	NEFilterActionInvalid NEFilterAction = 0
	// NEFilterActionRemediate: Remediate the flow.
	NEFilterActionRemediate NEFilterAction = 3
)

func (e NEFilterAction) String() string {
	switch e {
	case NEFilterActionAllow:
		return "NEFilterActionAllow"
	case NEFilterActionDrop:
		return "NEFilterActionDrop"
	case NEFilterActionFilterData:
		return "NEFilterActionFilterData"
	case NEFilterActionInvalid:
		return "NEFilterActionInvalid"
	case NEFilterActionRemediate:
		return "NEFilterActionRemediate"
	default:
		return fmt.Sprintf("NEFilterAction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataAttribute
type NEFilterDataAttribute int

const (
	// NEFilterDataAttributeHasIPHeader: An attribute that indicates the data includes an IP header.
	NEFilterDataAttributeHasIPHeader NEFilterDataAttribute = 1
)

func (e NEFilterDataAttribute) String() string {
	switch e {
	case NEFilterDataAttributeHasIPHeader:
		return "NEFilterDataAttributeHasIPHeader"
	default:
		return fmt.Sprintf("NEFilterDataAttribute(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManagerError
type NEFilterManagerError int

const (
	// NEFilterManagerErrorConfigurationCannotBeRemoved: An error code that indicates removing the configuration isn’t allowed.
	NEFilterManagerErrorConfigurationCannotBeRemoved NEFilterManagerError = 4
	// NEFilterManagerErrorConfigurationDisabled: An error code that indicates the filter configuration isn’t enabled.
	NEFilterManagerErrorConfigurationDisabled NEFilterManagerError = 2
	// NEFilterManagerErrorConfigurationInternalError: An error code that indicates an internal configuration error occurred.
	NEFilterManagerErrorConfigurationInternalError NEFilterManagerError = 6
	// NEFilterManagerErrorConfigurationInvalid: An error code that indicates the filter configuration is invalid.
	NEFilterManagerErrorConfigurationInvalid NEFilterManagerError = 1
	// NEFilterManagerErrorConfigurationPermissionDenied: An error code that indicates the configuration lacks permission.
	NEFilterManagerErrorConfigurationPermissionDenied NEFilterManagerError = 5
	// NEFilterManagerErrorConfigurationStale: An error code that indicates another process modfied the filter configuration since the last time the app loaded the configuration.
	NEFilterManagerErrorConfigurationStale NEFilterManagerError = 3
)

func (e NEFilterManagerError) String() string {
	switch e {
	case NEFilterManagerErrorConfigurationCannotBeRemoved:
		return "NEFilterManagerErrorConfigurationCannotBeRemoved"
	case NEFilterManagerErrorConfigurationDisabled:
		return "NEFilterManagerErrorConfigurationDisabled"
	case NEFilterManagerErrorConfigurationInternalError:
		return "NEFilterManagerErrorConfigurationInternalError"
	case NEFilterManagerErrorConfigurationInvalid:
		return "NEFilterManagerErrorConfigurationInvalid"
	case NEFilterManagerErrorConfigurationPermissionDenied:
		return "NEFilterManagerErrorConfigurationPermissionDenied"
	case NEFilterManagerErrorConfigurationStale:
		return "NEFilterManagerErrorConfigurationStale"
	default:
		return fmt.Sprintf("NEFilterManagerError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterManager/Grade-swift.enum
type NEFilterManagerGrade int

const (
	// NEFilterManagerGradeFirewall: A grade for filters that act as firewalls, blocking some network traffic.
	NEFilterManagerGradeFirewall NEFilterManagerGrade = 1
	// NEFilterManagerGradeInspector: A grade for filters that act as inspectors of network traffic.
	NEFilterManagerGradeInspector NEFilterManagerGrade = 2
)

func (e NEFilterManagerGrade) String() string {
	switch e {
	case NEFilterManagerGradeFirewall:
		return "NEFilterManagerGradeFirewall"
	case NEFilterManagerGradeInspector:
		return "NEFilterManagerGradeInspector"
	default:
		return fmt.Sprintf("NEFilterManagerGrade(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/Verdict
type NEFilterPacketProviderVerdict int

const (
	// NEFilterPacketProviderVerdictAllow: A verdict to allow a packet.
	NEFilterPacketProviderVerdictAllow NEFilterPacketProviderVerdict = 0
	// NEFilterPacketProviderVerdictDelay: A verdict to delay a packet until a future verdict.
	NEFilterPacketProviderVerdictDelay NEFilterPacketProviderVerdict = 2
	// NEFilterPacketProviderVerdictDrop: A verdict to drop a packet.
	NEFilterPacketProviderVerdictDrop NEFilterPacketProviderVerdict = 1
)

func (e NEFilterPacketProviderVerdict) String() string {
	switch e {
	case NEFilterPacketProviderVerdictAllow:
		return "NEFilterPacketProviderVerdictAllow"
	case NEFilterPacketProviderVerdictDelay:
		return "NEFilterPacketProviderVerdictDelay"
	case NEFilterPacketProviderVerdictDrop:
		return "NEFilterPacketProviderVerdictDrop"
	default:
		return fmt.Sprintf("NEFilterPacketProviderVerdict(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Event-swift.enum
type NEFilterReportEvent int

const (
	// NEFilterReportEventDataDecision: A type of event indicating the report is about a pass/block decision made after analyzing some amount of a flow’s data.
	NEFilterReportEventDataDecision NEFilterReportEvent = 2
	// NEFilterReportEventFlowClosed: A type of event indicating the report is for a flow’s closing.
	NEFilterReportEventFlowClosed NEFilterReportEvent = 3
	// NEFilterReportEventNewFlow: A type of event indicating the report is for a new flow.
	NEFilterReportEventNewFlow NEFilterReportEvent = 1
	// NEFilterReportEventStatistics: A type of event indicating the report is for the latest statistics of the flow.
	NEFilterReportEventStatistics NEFilterReportEvent = 4
)

func (e NEFilterReportEvent) String() string {
	switch e {
	case NEFilterReportEventDataDecision:
		return "NEFilterReportEventDataDecision"
	case NEFilterReportEventFlowClosed:
		return "NEFilterReportEventFlowClosed"
	case NEFilterReportEventNewFlow:
		return "NEFilterReportEventNewFlow"
	case NEFilterReportEventStatistics:
		return "NEFilterReportEventStatistics"
	default:
		return fmt.Sprintf("NEFilterReportEvent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Frequency
type NEFilterReportFrequency int

const (
	// NEFilterReportFrequencyHigh: A low frequency of reports, about once every half-second.
	NEFilterReportFrequencyHigh NEFilterReportFrequency = 3
	// NEFilterReportFrequencyLow: A low frequency of reports, about once every five seconds.
	NEFilterReportFrequencyLow NEFilterReportFrequency = 1
	// NEFilterReportFrequencyMedium: A low frequency of reports, about once every second.
	NEFilterReportFrequencyMedium NEFilterReportFrequency = 2
	// NEFilterReportFrequencyNone: A frequency value that indicates no report delivery.
	NEFilterReportFrequencyNone NEFilterReportFrequency = 0
)

func (e NEFilterReportFrequency) String() string {
	switch e {
	case NEFilterReportFrequencyHigh:
		return "NEFilterReportFrequencyHigh"
	case NEFilterReportFrequencyLow:
		return "NEFilterReportFrequencyLow"
	case NEFilterReportFrequencyMedium:
		return "NEFilterReportFrequencyMedium"
	case NEFilterReportFrequencyNone:
		return "NEFilterReportFrequencyNone"
	default:
		return fmt.Sprintf("NEFilterReportFrequency(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/TLSVersion
type NEHotspotConfigurationEAPTLSVersion int

const (
	// NEHotspotConfigurationEAPTLSVersion_1_0: Network EAPTLS version 1.0.
	NEHotspotConfigurationEAPTLSVersion_1_0 NEHotspotConfigurationEAPTLSVersion = 0
	// NEHotspotConfigurationEAPTLSVersion_1_1: Network EAPTLS version 1.1.
	NEHotspotConfigurationEAPTLSVersion_1_1 NEHotspotConfigurationEAPTLSVersion = 1
	// NEHotspotConfigurationEAPTLSVersion_1_2: Network EAPTLS version 1.2.
	NEHotspotConfigurationEAPTLSVersion_1_2 NEHotspotConfigurationEAPTLSVersion = 2
)

func (e NEHotspotConfigurationEAPTLSVersion) String() string {
	switch e {
	case NEHotspotConfigurationEAPTLSVersion_1_0:
		return "NEHotspotConfigurationEAPTLSVersion_1_0"
	case NEHotspotConfigurationEAPTLSVersion_1_1:
		return "NEHotspotConfigurationEAPTLSVersion_1_1"
	case NEHotspotConfigurationEAPTLSVersion_1_2:
		return "NEHotspotConfigurationEAPTLSVersion_1_2"
	default:
		return fmt.Sprintf("NEHotspotConfigurationEAPTLSVersion(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/EAPType
type NEHotspotConfigurationEAPType int

const (
	// NEHotspotConfigurationEAPTypeEAPFAST: Network EAP type is [EAPFAST].
	NEHotspotConfigurationEAPTypeEAPFAST NEHotspotConfigurationEAPType = 43
	// NEHotspotConfigurationEAPTypeEAPPEAP: Network EAP type is [EAPPEAP].
	NEHotspotConfigurationEAPTypeEAPPEAP NEHotspotConfigurationEAPType = 25
	// NEHotspotConfigurationEAPTypeEAPTLS: Network EAP type is [EAPTLS].
	NEHotspotConfigurationEAPTypeEAPTLS NEHotspotConfigurationEAPType = 13
	// NEHotspotConfigurationEAPTypeEAPTTLS: Network EAP type is [EAPTTLS].
	NEHotspotConfigurationEAPTypeEAPTTLS NEHotspotConfigurationEAPType = 21
)

func (e NEHotspotConfigurationEAPType) String() string {
	switch e {
	case NEHotspotConfigurationEAPTypeEAPFAST:
		return "NEHotspotConfigurationEAPTypeEAPFAST"
	case NEHotspotConfigurationEAPTypeEAPPEAP:
		return "NEHotspotConfigurationEAPTypeEAPPEAP"
	case NEHotspotConfigurationEAPTypeEAPTLS:
		return "NEHotspotConfigurationEAPTypeEAPTLS"
	case NEHotspotConfigurationEAPTypeEAPTTLS:
		return "NEHotspotConfigurationEAPTypeEAPTTLS"
	default:
		return fmt.Sprintf("NEHotspotConfigurationEAPType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationError
type NEHotspotConfigurationError int

const (
	// NEHotspotConfigurationErrorAlreadyAssociated: The configuration is already associated with the hotspot.
	NEHotspotConfigurationErrorAlreadyAssociated NEHotspotConfigurationError = 13
	// NEHotspotConfigurationErrorApplicationIsNotInForeground: The application is not running in the foreground.
	NEHotspotConfigurationErrorApplicationIsNotInForeground NEHotspotConfigurationError = 14
	// NEHotspotConfigurationErrorInternal: Internal error, otherwise undefined.
	NEHotspotConfigurationErrorInternal NEHotspotConfigurationError = 8
	// NEHotspotConfigurationErrorInvalid: The configuration is not valid.
	NEHotspotConfigurationErrorInvalid NEHotspotConfigurationError = 0
	// NEHotspotConfigurationErrorInvalidEAPSettings: EAP settings are not valid.
	NEHotspotConfigurationErrorInvalidEAPSettings NEHotspotConfigurationError = 4
	// NEHotspotConfigurationErrorInvalidHS20DomainName: The HS 2.0 domain name is not valid.
	NEHotspotConfigurationErrorInvalidHS20DomainName NEHotspotConfigurationError = 6
	// NEHotspotConfigurationErrorInvalidHS20Settings: The HS 2.0 settings are not valid.
	NEHotspotConfigurationErrorInvalidHS20Settings NEHotspotConfigurationError = 5
	// NEHotspotConfigurationErrorInvalidSSID: The SSID value is not valid.
	NEHotspotConfigurationErrorInvalidSSID NEHotspotConfigurationError = 1
	// NEHotspotConfigurationErrorInvalidSSIDPrefix: The SSID prefix used to create the hotspot configuration is invalid.
	NEHotspotConfigurationErrorInvalidSSIDPrefix NEHotspotConfigurationError = 15
	// NEHotspotConfigurationErrorInvalidWEPPassphrase: The WEP passphrase is not valid.
	NEHotspotConfigurationErrorInvalidWEPPassphrase NEHotspotConfigurationError = 3
	// NEHotspotConfigurationErrorInvalidWPAPassphrase: The WPA passphrase is not valid.
	NEHotspotConfigurationErrorInvalidWPAPassphrase NEHotspotConfigurationError = 2
	// NEHotspotConfigurationErrorJoinOnceNotSupported: The join-once option isn’t support for EAP configuration.
	NEHotspotConfigurationErrorJoinOnceNotSupported NEHotspotConfigurationError = 12
	// NEHotspotConfigurationErrorPending: The network configuration action has not completed.
	NEHotspotConfigurationErrorPending NEHotspotConfigurationError = 9
	// NEHotspotConfigurationErrorSystemConfiguration: The system configuration is not valid.
	NEHotspotConfigurationErrorSystemConfiguration NEHotspotConfigurationError = 10
	NEHotspotConfigurationErrorSystemDenied        NEHotspotConfigurationError = 17
	// NEHotspotConfigurationErrorUnknown: An unknown error has occurred.
	NEHotspotConfigurationErrorUnknown NEHotspotConfigurationError = 11
	// NEHotspotConfigurationErrorUserDenied: The user has refused the network configuration.
	NEHotspotConfigurationErrorUserDenied       NEHotspotConfigurationError = 7
	NEHotspotConfigurationErrorUserUnauthorized NEHotspotConfigurationError = 16
)

func (e NEHotspotConfigurationError) String() string {
	switch e {
	case NEHotspotConfigurationErrorAlreadyAssociated:
		return "NEHotspotConfigurationErrorAlreadyAssociated"
	case NEHotspotConfigurationErrorApplicationIsNotInForeground:
		return "NEHotspotConfigurationErrorApplicationIsNotInForeground"
	case NEHotspotConfigurationErrorInternal:
		return "NEHotspotConfigurationErrorInternal"
	case NEHotspotConfigurationErrorInvalid:
		return "NEHotspotConfigurationErrorInvalid"
	case NEHotspotConfigurationErrorInvalidEAPSettings:
		return "NEHotspotConfigurationErrorInvalidEAPSettings"
	case NEHotspotConfigurationErrorInvalidHS20DomainName:
		return "NEHotspotConfigurationErrorInvalidHS20DomainName"
	case NEHotspotConfigurationErrorInvalidHS20Settings:
		return "NEHotspotConfigurationErrorInvalidHS20Settings"
	case NEHotspotConfigurationErrorInvalidSSID:
		return "NEHotspotConfigurationErrorInvalidSSID"
	case NEHotspotConfigurationErrorInvalidSSIDPrefix:
		return "NEHotspotConfigurationErrorInvalidSSIDPrefix"
	case NEHotspotConfigurationErrorInvalidWEPPassphrase:
		return "NEHotspotConfigurationErrorInvalidWEPPassphrase"
	case NEHotspotConfigurationErrorInvalidWPAPassphrase:
		return "NEHotspotConfigurationErrorInvalidWPAPassphrase"
	case NEHotspotConfigurationErrorJoinOnceNotSupported:
		return "NEHotspotConfigurationErrorJoinOnceNotSupported"
	case NEHotspotConfigurationErrorPending:
		return "NEHotspotConfigurationErrorPending"
	case NEHotspotConfigurationErrorSystemConfiguration:
		return "NEHotspotConfigurationErrorSystemConfiguration"
	case NEHotspotConfigurationErrorSystemDenied:
		return "NEHotspotConfigurationErrorSystemDenied"
	case NEHotspotConfigurationErrorUnknown:
		return "NEHotspotConfigurationErrorUnknown"
	case NEHotspotConfigurationErrorUserDenied:
		return "NEHotspotConfigurationErrorUserDenied"
	case NEHotspotConfigurationErrorUserUnauthorized:
		return "NEHotspotConfigurationErrorUserUnauthorized"
	default:
		return fmt.Sprintf("NEHotspotConfigurationError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEHotspotEAPSettings/TTLSInnerAuthenticationType-swift.enum
type NEHotspotConfigurationTTLSInnerAuthenticationType int

const (
	// NEHotspotConfigurationEAPTTLSInnerAuthenticationCHAP: Network EAPTTLS inner authentication type is CHAP.
	NEHotspotConfigurationEAPTTLSInnerAuthenticationCHAP NEHotspotConfigurationTTLSInnerAuthenticationType = 1
	// NEHotspotConfigurationEAPTTLSInnerAuthenticationEAP: Network EAPTTLS inner authentication type is EAP.
	NEHotspotConfigurationEAPTTLSInnerAuthenticationEAP NEHotspotConfigurationTTLSInnerAuthenticationType = 4
	// NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAP: Network EAPTTLS inner authentication type is MSCHAP.
	NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAP NEHotspotConfigurationTTLSInnerAuthenticationType = 2
	// NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAPv2: Network EAPTTLS inner authentication type is MSCHAP, version 2.
	NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAPv2 NEHotspotConfigurationTTLSInnerAuthenticationType = 3
	// NEHotspotConfigurationEAPTTLSInnerAuthenticationPAP: Network EAPTTLS inner authentication type is PAP.
	NEHotspotConfigurationEAPTTLSInnerAuthenticationPAP NEHotspotConfigurationTTLSInnerAuthenticationType = 0
)

func (e NEHotspotConfigurationTTLSInnerAuthenticationType) String() string {
	switch e {
	case NEHotspotConfigurationEAPTTLSInnerAuthenticationCHAP:
		return "NEHotspotConfigurationEAPTTLSInnerAuthenticationCHAP"
	case NEHotspotConfigurationEAPTTLSInnerAuthenticationEAP:
		return "NEHotspotConfigurationEAPTTLSInnerAuthenticationEAP"
	case NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAP:
		return "NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAP"
	case NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAPv2:
		return "NEHotspotConfigurationEAPTTLSInnerAuthenticationMSCHAPv2"
	case NEHotspotConfigurationEAPTTLSInnerAuthenticationPAP:
		return "NEHotspotConfigurationEAPTTLSInnerAuthenticationPAP"
	default:
		return fmt.Sprintf("NEHotspotConfigurationTTLSInnerAuthenticationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommandType
type NEHotspotHelperCommandType int

const (
	// KNEHotspotHelperCommandTypeAuthenticate: Authenticate to the network.
	KNEHotspotHelperCommandTypeAuthenticate NEHotspotHelperCommandType = 3
	// KNEHotspotHelperCommandTypeEvaluate: Evaluate the network.
	KNEHotspotHelperCommandTypeEvaluate NEHotspotHelperCommandType = 2
	// KNEHotspotHelperCommandTypeFilterScanList: Filter the Wi-Fi scan list.
	KNEHotspotHelperCommandTypeFilterScanList NEHotspotHelperCommandType = 1
	// KNEHotspotHelperCommandTypeLogoff: Logoff the network.
	KNEHotspotHelperCommandTypeLogoff NEHotspotHelperCommandType = 6
	// KNEHotspotHelperCommandTypeMaintain: Maintain the connection to the network.
	KNEHotspotHelperCommandTypeMaintain NEHotspotHelperCommandType = 5
	// KNEHotspotHelperCommandTypeNone: Placeholder for the null command.
	KNEHotspotHelperCommandTypeNone NEHotspotHelperCommandType = 0
	// KNEHotspotHelperCommandTypePresentUI: Present user interface.
	KNEHotspotHelperCommandTypePresentUI NEHotspotHelperCommandType = 4
)

func (e NEHotspotHelperCommandType) String() string {
	switch e {
	case KNEHotspotHelperCommandTypeAuthenticate:
		return "KNEHotspotHelperCommandTypeAuthenticate"
	case KNEHotspotHelperCommandTypeEvaluate:
		return "KNEHotspotHelperCommandTypeEvaluate"
	case KNEHotspotHelperCommandTypeFilterScanList:
		return "KNEHotspotHelperCommandTypeFilterScanList"
	case KNEHotspotHelperCommandTypeLogoff:
		return "KNEHotspotHelperCommandTypeLogoff"
	case KNEHotspotHelperCommandTypeMaintain:
		return "KNEHotspotHelperCommandTypeMaintain"
	case KNEHotspotHelperCommandTypeNone:
		return "KNEHotspotHelperCommandTypeNone"
	case KNEHotspotHelperCommandTypePresentUI:
		return "KNEHotspotHelperCommandTypePresentUI"
	default:
		return fmt.Sprintf("NEHotspotHelperCommandType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperConfidence
type NEHotspotHelperConfidence int

const (
	// KNEHotspotHelperConfidenceHigh: The helper has high confidence in being able to handle the network.
	KNEHotspotHelperConfidenceHigh NEHotspotHelperConfidence = 2
	// KNEHotspotHelperConfidenceLow: The helper has some confidence in being able to handle the network.
	KNEHotspotHelperConfidenceLow NEHotspotHelperConfidence = 1
	// KNEHotspotHelperConfidenceNone: The helper is unable to handle the network.
	KNEHotspotHelperConfidenceNone NEHotspotHelperConfidence = 0
)

func (e NEHotspotHelperConfidence) String() string {
	switch e {
	case KNEHotspotHelperConfidenceHigh:
		return "KNEHotspotHelperConfidenceHigh"
	case KNEHotspotHelperConfidenceLow:
		return "KNEHotspotHelperConfidenceLow"
	case KNEHotspotHelperConfidenceNone:
		return "KNEHotspotHelperConfidenceNone"
	default:
		return fmt.Sprintf("NEHotspotHelperConfidence(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperResult
type NEHotspotHelperResult int

const (
	// KNEHotspotHelperResultAuthenticationRequired: The network requires authentication again.
	KNEHotspotHelperResultAuthenticationRequired NEHotspotHelperResult = 4
	// KNEHotspotHelperResultCommandNotRecognized: The helper did not recognize the command type.
	KNEHotspotHelperResultCommandNotRecognized NEHotspotHelperResult = 3
	// KNEHotspotHelperResultFailure: The command failed to be handled.
	KNEHotspotHelperResultFailure NEHotspotHelperResult = 1
	// KNEHotspotHelperResultSuccess: The command was handled successfully.
	KNEHotspotHelperResultSuccess NEHotspotHelperResult = 0
	// KNEHotspotHelperResultTemporaryFailure: The Hotspot Helper app determined that it is temporarily unable to perform the authentication.
	KNEHotspotHelperResultTemporaryFailure NEHotspotHelperResult = 6
	// KNEHotspotHelperResultUIRequired: The operation requires user interaction.
	KNEHotspotHelperResultUIRequired NEHotspotHelperResult = 2
	// KNEHotspotHelperResultUnsupportedNetwork: After attempting to authenticate, the Hotspot Helper app determined that it can’t perform the authentication.
	KNEHotspotHelperResultUnsupportedNetwork NEHotspotHelperResult = 5
)

func (e NEHotspotHelperResult) String() string {
	switch e {
	case KNEHotspotHelperResultAuthenticationRequired:
		return "KNEHotspotHelperResultAuthenticationRequired"
	case KNEHotspotHelperResultCommandNotRecognized:
		return "KNEHotspotHelperResultCommandNotRecognized"
	case KNEHotspotHelperResultFailure:
		return "KNEHotspotHelperResultFailure"
	case KNEHotspotHelperResultSuccess:
		return "KNEHotspotHelperResultSuccess"
	case KNEHotspotHelperResultTemporaryFailure:
		return "KNEHotspotHelperResultTemporaryFailure"
	case KNEHotspotHelperResultUIRequired:
		return "KNEHotspotHelperResultUIRequired"
	case KNEHotspotHelperResultUnsupportedNetwork:
		return "KNEHotspotHelperResultUnsupportedNetwork"
	default:
		return fmt.Sprintf("NEHotspotHelperResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEHotspotNetworkSecurityType
type NEHotspotNetworkSecurityType int

const (
	// NEHotspotNetworkSecurityTypeEnterprise: A security type to represent use of Wi-Fi protected access (WPA), WPA2, and WPA3 standards using enterprise-level seciurity.
	NEHotspotNetworkSecurityTypeEnterprise NEHotspotNetworkSecurityType = 3
	// NEHotspotNetworkSecurityTypeOpen: A security type to represent an open network with no security protocol.
	NEHotspotNetworkSecurityTypeOpen NEHotspotNetworkSecurityType = 0
	// NEHotspotNetworkSecurityTypePersonal: A security type to represent use of Wi-Fi protected access (WPA), WPA2, and WPA3 standards using a pre-shared secret.
	NEHotspotNetworkSecurityTypePersonal NEHotspotNetworkSecurityType = 2
	// NEHotspotNetworkSecurityTypeUnknown: A value that represents an unknown security type.
	NEHotspotNetworkSecurityTypeUnknown NEHotspotNetworkSecurityType = 4
	// NEHotspotNetworkSecurityTypeWEP: A security type to represent use of Wired Equivalent Privacy (WEP).
	NEHotspotNetworkSecurityTypeWEP NEHotspotNetworkSecurityType = 1
)

func (e NEHotspotNetworkSecurityType) String() string {
	switch e {
	case NEHotspotNetworkSecurityTypeEnterprise:
		return "NEHotspotNetworkSecurityTypeEnterprise"
	case NEHotspotNetworkSecurityTypeOpen:
		return "NEHotspotNetworkSecurityTypeOpen"
	case NEHotspotNetworkSecurityTypePersonal:
		return "NEHotspotNetworkSecurityTypePersonal"
	case NEHotspotNetworkSecurityTypeUnknown:
		return "NEHotspotNetworkSecurityTypeUnknown"
	case NEHotspotNetworkSecurityTypeWEP:
		return "NEHotspotNetworkSecurityTypeWEP"
	default:
		return fmt.Sprintf("NEHotspotNetworkSecurityType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/Protocol
type NENetworkRuleProtocol int

const (
	// NENetworkRuleProtocolAny: A rule protocol to match TCP and UDP traffic.
	NENetworkRuleProtocolAny NENetworkRuleProtocol = 0
	// NENetworkRuleProtocolTCP: A rule protocol to match TCP traffic.
	NENetworkRuleProtocolTCP NENetworkRuleProtocol = 1
	// NENetworkRuleProtocolUDP: A rule protocol to match UDP traffic.
	NENetworkRuleProtocolUDP NENetworkRuleProtocol = 2
)

func (e NENetworkRuleProtocol) String() string {
	switch e {
	case NENetworkRuleProtocolAny:
		return "NENetworkRuleProtocolAny"
	case NENetworkRuleProtocolTCP:
		return "NENetworkRuleProtocolTCP"
	case NENetworkRuleProtocolUDP:
		return "NENetworkRuleProtocolUDP"
	default:
		return fmt.Sprintf("NENetworkRuleProtocol(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleAction
type NEOnDemandRuleAction int

const (
	// NEOnDemandRuleActionConnect: Start the VPN connection for every connection attempt.
	NEOnDemandRuleActionConnect NEOnDemandRuleAction = 1
	// NEOnDemandRuleActionDisconnect: Do not start the VPN connection, and disconnect the VPN connection if it is not currently disconnected.
	NEOnDemandRuleActionDisconnect NEOnDemandRuleAction = 2
	// NEOnDemandRuleActionEvaluateConnection: Start the VPN after evaluating the destination host being accessed against the rule’s parameters.
	NEOnDemandRuleActionEvaluateConnection NEOnDemandRuleAction = 3
	// NEOnDemandRuleActionIgnore: Do not start the VPN connection, but do not disconnect it if it is currently connected.
	NEOnDemandRuleActionIgnore NEOnDemandRuleAction = 4
)

func (e NEOnDemandRuleAction) String() string {
	switch e {
	case NEOnDemandRuleActionConnect:
		return "NEOnDemandRuleActionConnect"
	case NEOnDemandRuleActionDisconnect:
		return "NEOnDemandRuleActionDisconnect"
	case NEOnDemandRuleActionEvaluateConnection:
		return "NEOnDemandRuleActionEvaluateConnection"
	case NEOnDemandRuleActionIgnore:
		return "NEOnDemandRuleActionIgnore"
	default:
		return fmt.Sprintf("NEOnDemandRuleAction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleInterfaceType
type NEOnDemandRuleInterfaceType int

const (
	// NEOnDemandRuleInterfaceTypeAny: Match any interface type
	NEOnDemandRuleInterfaceTypeAny NEOnDemandRuleInterfaceType = 0
	// NEOnDemandRuleInterfaceTypeCellular: Match cellular data interfaces
	NEOnDemandRuleInterfaceTypeCellular NEOnDemandRuleInterfaceType = 3
	// NEOnDemandRuleInterfaceTypeEthernet: Match wired ethernet interfaces
	NEOnDemandRuleInterfaceTypeEthernet NEOnDemandRuleInterfaceType = 1
	// NEOnDemandRuleInterfaceTypeWiFi: Match Wi-Fi interfaces
	NEOnDemandRuleInterfaceTypeWiFi NEOnDemandRuleInterfaceType = 2
)

func (e NEOnDemandRuleInterfaceType) String() string {
	switch e {
	case NEOnDemandRuleInterfaceTypeAny:
		return "NEOnDemandRuleInterfaceTypeAny"
	case NEOnDemandRuleInterfaceTypeCellular:
		return "NEOnDemandRuleInterfaceTypeCellular"
	case NEOnDemandRuleInterfaceTypeEthernet:
		return "NEOnDemandRuleInterfaceTypeEthernet"
	case NEOnDemandRuleInterfaceTypeWiFi:
		return "NEOnDemandRuleInterfaceTypeWiFi"
	default:
		return fmt.Sprintf("NEOnDemandRuleInterfaceType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEProviderStopReason
type NEProviderStopReason int

const (
	NEProviderStopReasonAppUpdate NEProviderStopReason = 16
	// NEProviderStopReasonAuthenticationCanceled: The authentication process was canceled.
	NEProviderStopReasonAuthenticationCanceled NEProviderStopReason = 6
	// NEProviderStopReasonConfigurationDisabled: The configuration was disabled.
	NEProviderStopReasonConfigurationDisabled NEProviderStopReason = 9
	// NEProviderStopReasonConfigurationFailed: The configuration is invalid.
	NEProviderStopReasonConfigurationFailed NEProviderStopReason = 7
	// NEProviderStopReasonConfigurationRemoved: The configuration was removed.
	NEProviderStopReasonConfigurationRemoved NEProviderStopReason = 10
	// NEProviderStopReasonConnectionFailed: The connection failed.
	NEProviderStopReasonConnectionFailed NEProviderStopReason = 14
	// NEProviderStopReasonIdleTimeout: The session timed out.
	NEProviderStopReasonIdleTimeout NEProviderStopReason = 8
	// NEProviderStopReasonInternalError: The provider encountered an internal error.
	NEProviderStopReasonInternalError NEProviderStopReason = 17
	// NEProviderStopReasonNoNetworkAvailable: No network connectivity is currently available.
	NEProviderStopReasonNoNetworkAvailable NEProviderStopReason = 3
	// NEProviderStopReasonNone: No specific reason.
	NEProviderStopReasonNone NEProviderStopReason = 0
	// NEProviderStopReasonProviderDisabled: The provider was disabled.
	NEProviderStopReasonProviderDisabled NEProviderStopReason = 5
	// NEProviderStopReasonProviderFailed: The provider failed to function correctly.
	NEProviderStopReasonProviderFailed NEProviderStopReason = 2
	// NEProviderStopReasonSleep: A stop reason indicating the configuration enabled disconnect on sleep and the device went to sleep.
	NEProviderStopReasonSleep NEProviderStopReason = 15
	// NEProviderStopReasonSuperceded: The configuration was superceded by a higher-priority configuration.
	NEProviderStopReasonSuperceded NEProviderStopReason = 11
	// NEProviderStopReasonUnrecoverableNetworkChange: The device’s network connectivity changed.
	NEProviderStopReasonUnrecoverableNetworkChange NEProviderStopReason = 4
	// NEProviderStopReasonUserInitiated: The user stopped the provider extension.
	NEProviderStopReasonUserInitiated NEProviderStopReason = 1
	// NEProviderStopReasonUserLogout: The user logged out.
	NEProviderStopReasonUserLogout NEProviderStopReason = 12
	// NEProviderStopReasonUserSwitch: The current console user changed.
	NEProviderStopReasonUserSwitch NEProviderStopReason = 13
)

func (e NEProviderStopReason) String() string {
	switch e {
	case NEProviderStopReasonAppUpdate:
		return "NEProviderStopReasonAppUpdate"
	case NEProviderStopReasonAuthenticationCanceled:
		return "NEProviderStopReasonAuthenticationCanceled"
	case NEProviderStopReasonConfigurationDisabled:
		return "NEProviderStopReasonConfigurationDisabled"
	case NEProviderStopReasonConfigurationFailed:
		return "NEProviderStopReasonConfigurationFailed"
	case NEProviderStopReasonConfigurationRemoved:
		return "NEProviderStopReasonConfigurationRemoved"
	case NEProviderStopReasonConnectionFailed:
		return "NEProviderStopReasonConnectionFailed"
	case NEProviderStopReasonIdleTimeout:
		return "NEProviderStopReasonIdleTimeout"
	case NEProviderStopReasonInternalError:
		return "NEProviderStopReasonInternalError"
	case NEProviderStopReasonNoNetworkAvailable:
		return "NEProviderStopReasonNoNetworkAvailable"
	case NEProviderStopReasonNone:
		return "NEProviderStopReasonNone"
	case NEProviderStopReasonProviderDisabled:
		return "NEProviderStopReasonProviderDisabled"
	case NEProviderStopReasonProviderFailed:
		return "NEProviderStopReasonProviderFailed"
	case NEProviderStopReasonSleep:
		return "NEProviderStopReasonSleep"
	case NEProviderStopReasonSuperceded:
		return "NEProviderStopReasonSuperceded"
	case NEProviderStopReasonUnrecoverableNetworkChange:
		return "NEProviderStopReasonUnrecoverableNetworkChange"
	case NEProviderStopReasonUserInitiated:
		return "NEProviderStopReasonUserInitiated"
	case NEProviderStopReasonUserLogout:
		return "NEProviderStopReasonUserLogout"
	case NEProviderStopReasonUserSwitch:
		return "NEProviderStopReasonUserSwitch"
	default:
		return fmt.Sprintf("NEProviderStopReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerClientError
type NERelayManagerClientError int

const (
	NERelayManagerClientErrorCertificateExpired       NERelayManagerClientError = 7
	NERelayManagerClientErrorCertificateInvalid       NERelayManagerClientError = 6
	NERelayManagerClientErrorCertificateMissing       NERelayManagerClientError = 5
	NERelayManagerClientErrorDNSFailed                NERelayManagerClientError = 2
	NERelayManagerClientErrorNone                     NERelayManagerClientError = 1
	NERelayManagerClientErrorOther                    NERelayManagerClientError = 10
	NERelayManagerClientErrorServerCertificateExpired NERelayManagerClientError = 9
	NERelayManagerClientErrorServerCertificateInvalid NERelayManagerClientError = 8
	NERelayManagerClientErrorServerDisconnected       NERelayManagerClientError = 4
	NERelayManagerClientErrorServerUnreachable        NERelayManagerClientError = 3
)

func (e NERelayManagerClientError) String() string {
	switch e {
	case NERelayManagerClientErrorCertificateExpired:
		return "NERelayManagerClientErrorCertificateExpired"
	case NERelayManagerClientErrorCertificateInvalid:
		return "NERelayManagerClientErrorCertificateInvalid"
	case NERelayManagerClientErrorCertificateMissing:
		return "NERelayManagerClientErrorCertificateMissing"
	case NERelayManagerClientErrorDNSFailed:
		return "NERelayManagerClientErrorDNSFailed"
	case NERelayManagerClientErrorNone:
		return "NERelayManagerClientErrorNone"
	case NERelayManagerClientErrorOther:
		return "NERelayManagerClientErrorOther"
	case NERelayManagerClientErrorServerCertificateExpired:
		return "NERelayManagerClientErrorServerCertificateExpired"
	case NERelayManagerClientErrorServerCertificateInvalid:
		return "NERelayManagerClientErrorServerCertificateInvalid"
	case NERelayManagerClientErrorServerDisconnected:
		return "NERelayManagerClientErrorServerDisconnected"
	case NERelayManagerClientErrorServerUnreachable:
		return "NERelayManagerClientErrorServerUnreachable"
	default:
		return fmt.Sprintf("NERelayManagerClientError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NERelayManagerError
type NERelayManagerError int

const (
	// NERelayManagerErrorConfigurationCannotBeRemoved: An error code that indicates removing the relay manager failed.
	NERelayManagerErrorConfigurationCannotBeRemoved NERelayManagerError = 4
	// NERelayManagerErrorConfigurationDisabled: An error code that indicates the relay manager isn’t enabled.
	NERelayManagerErrorConfigurationDisabled NERelayManagerError = 2
	// NERelayManagerErrorConfigurationInvalid: An error code that indicates the relay manager is invalid.
	NERelayManagerErrorConfigurationInvalid NERelayManagerError = 1
	// NERelayManagerErrorConfigurationStale: An error code that indicates the relay manager isn’t loaded.
	NERelayManagerErrorConfigurationStale NERelayManagerError = 3
)

func (e NERelayManagerError) String() string {
	switch e {
	case NERelayManagerErrorConfigurationCannotBeRemoved:
		return "NERelayManagerErrorConfigurationCannotBeRemoved"
	case NERelayManagerErrorConfigurationDisabled:
		return "NERelayManagerErrorConfigurationDisabled"
	case NERelayManagerErrorConfigurationInvalid:
		return "NERelayManagerErrorConfigurationInvalid"
	case NERelayManagerErrorConfigurationStale:
		return "NERelayManagerErrorConfigurationStale"
	default:
		return fmt.Sprintf("NERelayManagerError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NETrafficDirection
type NETrafficDirection int

const (
	// NETrafficDirectionAny: A direction that matches either inbound or outbound traffic.
	NETrafficDirectionAny NETrafficDirection = 0
	// NETrafficDirectionInbound: The inbound traffic direction.
	NETrafficDirectionInbound NETrafficDirection = 1
	// NETrafficDirectionOutbound: The outbound traffic direction.
	NETrafficDirectionOutbound NETrafficDirection = 2
)

func (e NETrafficDirection) String() string {
	switch e {
	case NETrafficDirectionAny:
		return "NETrafficDirectionAny"
	case NETrafficDirectionInbound:
		return "NETrafficDirectionInbound"
	case NETrafficDirectionOutbound:
		return "NETrafficDirectionOutbound"
	default:
		return fmt.Sprintf("NETrafficDirection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderError-swift.struct/Code
type NETunnelProviderError int

const (
	// NETunnelProviderErrorNetworkSettingsCanceled: The request to set or clear the tunnel network settings was canceled.
	NETunnelProviderErrorNetworkSettingsCanceled NETunnelProviderError = 2
	// NETunnelProviderErrorNetworkSettingsFailed: The request to set or clear the tunnel network settings failed.
	NETunnelProviderErrorNetworkSettingsFailed NETunnelProviderError = 3
	// NETunnelProviderErrorNetworkSettingsInvalid: The provided tunnel network settings are invalid.
	NETunnelProviderErrorNetworkSettingsInvalid NETunnelProviderError = 1
)

func (e NETunnelProviderError) String() string {
	switch e {
	case NETunnelProviderErrorNetworkSettingsCanceled:
		return "NETunnelProviderErrorNetworkSettingsCanceled"
	case NETunnelProviderErrorNetworkSettingsFailed:
		return "NETunnelProviderErrorNetworkSettingsFailed"
	case NETunnelProviderErrorNetworkSettingsInvalid:
		return "NETunnelProviderErrorNetworkSettingsInvalid"
	default:
		return fmt.Sprintf("NETunnelProviderError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderRoutingMethod
type NETunnelProviderRoutingMethod int

const (
	// NETunnelProviderRoutingMethodDestinationIP: Route network traffic to the tunnel based on destination IP.
	NETunnelProviderRoutingMethodDestinationIP NETunnelProviderRoutingMethod = 1
	// NETunnelProviderRoutingMethodNetworkRule: A routing method that routes traffic based on network rule objects specified by the provider.
	NETunnelProviderRoutingMethodNetworkRule NETunnelProviderRoutingMethod = 3
	// NETunnelProviderRoutingMethodSourceApplication: Route network traffic to the tunnel based on source application.
	NETunnelProviderRoutingMethodSourceApplication NETunnelProviderRoutingMethod = 2
)

func (e NETunnelProviderRoutingMethod) String() string {
	switch e {
	case NETunnelProviderRoutingMethodDestinationIP:
		return "NETunnelProviderRoutingMethodDestinationIP"
	case NETunnelProviderRoutingMethodNetworkRule:
		return "NETunnelProviderRoutingMethodNetworkRule"
	case NETunnelProviderRoutingMethodSourceApplication:
		return "NETunnelProviderRoutingMethodSourceApplication"
	default:
		return fmt.Sprintf("NETunnelProviderRoutingMethod(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEURLFilter/Verdict
type NEURLFilterVerdict int

const (
	// NEURLFilterVerdictAllow: A verdict that indicates that accessing the URL is allowed.
	NEURLFilterVerdictAllow NEURLFilterVerdict = 2
	// NEURLFilterVerdictDeny: A verdict that indicates that accessing the URL is denied.
	NEURLFilterVerdictDeny NEURLFilterVerdict = 3
	// NEURLFilterVerdictUnknown: A verdict that indicates URL validation failed.
	NEURLFilterVerdictUnknown NEURLFilterVerdict = 1
)

func (e NEURLFilterVerdict) String() string {
	switch e {
	case NEURLFilterVerdictAllow:
		return "NEURLFilterVerdictAllow"
	case NEURLFilterVerdictDeny:
		return "NEURLFilterVerdictDeny"
	case NEURLFilterVerdictUnknown:
		return "NEURLFilterVerdictUnknown"
	default:
		return fmt.Sprintf("NEURLFilterVerdict(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionError
type NEVPNConnectionError int

const (
	// NEVPNConnectionErrorAuthenticationFailed: An error code that indicates the VPN connection failed because the VPN server rejected the user credentials.
	NEVPNConnectionErrorAuthenticationFailed NEVPNConnectionError = 8
	// NEVPNConnectionErrorClientCertificateExpired: An error code that indicates the client certfiicate’s validity period has passed.
	NEVPNConnectionErrorClientCertificateExpired NEVPNConnectionError = 11
	// NEVPNConnectionErrorClientCertificateInvalid: An error code that indicates the client certfiicate is invalid.
	NEVPNConnectionErrorClientCertificateInvalid NEVPNConnectionError = 9
	// NEVPNConnectionErrorClientCertificateNotYetValid: An error code that indicates the client certfiicate won’t be valid until some time in the future.
	NEVPNConnectionErrorClientCertificateNotYetValid NEVPNConnectionError = 10
	// NEVPNConnectionErrorConfigurationFailed: An error code that indicates the VPN connection failed because the configuration is invalid.
	NEVPNConnectionErrorConfigurationFailed NEVPNConnectionError = 4
	// NEVPNConnectionErrorConfigurationNotFound: An error code that indicates the VPN connection failed because the system couldn’t find a configuration.
	NEVPNConnectionErrorConfigurationNotFound NEVPNConnectionError = 13
	// NEVPNConnectionErrorNegotiationFailed: An error code that indicates the VPN connection failed because the negotiation failed.
	NEVPNConnectionErrorNegotiationFailed NEVPNConnectionError = 15
	// NEVPNConnectionErrorNoNetworkAvailable: An error code that indicates the VPN connection failed because the system isn’t connected to a network.
	NEVPNConnectionErrorNoNetworkAvailable NEVPNConnectionError = 2
	// NEVPNConnectionErrorOverslept: An error code that indicates the system slept for an extended period of time, causing the VPN connection to terminate.
	NEVPNConnectionErrorOverslept NEVPNConnectionError = 1
	// NEVPNConnectionErrorPluginDisabled: An error code that indicates the VPN plugin isn’t available or needs an update.
	NEVPNConnectionErrorPluginDisabled NEVPNConnectionError = 14
	// NEVPNConnectionErrorPluginFailed: An error code that indicates the VPN plugin failed unexpectedly.
	NEVPNConnectionErrorPluginFailed NEVPNConnectionError = 12
	// NEVPNConnectionErrorServerAddressResolutionFailed: An error code that indicates the VPN connection failed because the system couldn’t determine the VPN server address.
	NEVPNConnectionErrorServerAddressResolutionFailed NEVPNConnectionError = 5
	// NEVPNConnectionErrorServerCertificateExpired: An error code that indicates the server certfiicate’s validity period has passed.
	NEVPNConnectionErrorServerCertificateExpired NEVPNConnectionError = 19
	// NEVPNConnectionErrorServerCertificateInvalid: An error code that indicates the server certfiicate is invalid.
	NEVPNConnectionErrorServerCertificateInvalid NEVPNConnectionError = 17
	// NEVPNConnectionErrorServerCertificateNotYetValid: An error code that indicates the server certfiicate won’t be valid until some time in the future.
	NEVPNConnectionErrorServerCertificateNotYetValid NEVPNConnectionError = 18
	// NEVPNConnectionErrorServerDead: An error code that indicates the VPN connection failed because the VPN server has stopped responding.
	NEVPNConnectionErrorServerDead NEVPNConnectionError = 7
	// NEVPNConnectionErrorServerDisconnected: An error code that indicates the VPN connection failed because the VPN server terminated the connection.
	NEVPNConnectionErrorServerDisconnected NEVPNConnectionError = 16
	// NEVPNConnectionErrorServerNotResponding: An error code that indicates the VPN connection failed because the VPN server isn’t responding.
	NEVPNConnectionErrorServerNotResponding NEVPNConnectionError = 6
	// NEVPNConnectionErrorUnrecoverableNetworkChange: An error code that indicates network conditions changed such that the VPN connection needed to terminate.
	NEVPNConnectionErrorUnrecoverableNetworkChange NEVPNConnectionError = 3
)

func (e NEVPNConnectionError) String() string {
	switch e {
	case NEVPNConnectionErrorAuthenticationFailed:
		return "NEVPNConnectionErrorAuthenticationFailed"
	case NEVPNConnectionErrorClientCertificateExpired:
		return "NEVPNConnectionErrorClientCertificateExpired"
	case NEVPNConnectionErrorClientCertificateInvalid:
		return "NEVPNConnectionErrorClientCertificateInvalid"
	case NEVPNConnectionErrorClientCertificateNotYetValid:
		return "NEVPNConnectionErrorClientCertificateNotYetValid"
	case NEVPNConnectionErrorConfigurationFailed:
		return "NEVPNConnectionErrorConfigurationFailed"
	case NEVPNConnectionErrorConfigurationNotFound:
		return "NEVPNConnectionErrorConfigurationNotFound"
	case NEVPNConnectionErrorNegotiationFailed:
		return "NEVPNConnectionErrorNegotiationFailed"
	case NEVPNConnectionErrorNoNetworkAvailable:
		return "NEVPNConnectionErrorNoNetworkAvailable"
	case NEVPNConnectionErrorOverslept:
		return "NEVPNConnectionErrorOverslept"
	case NEVPNConnectionErrorPluginDisabled:
		return "NEVPNConnectionErrorPluginDisabled"
	case NEVPNConnectionErrorPluginFailed:
		return "NEVPNConnectionErrorPluginFailed"
	case NEVPNConnectionErrorServerAddressResolutionFailed:
		return "NEVPNConnectionErrorServerAddressResolutionFailed"
	case NEVPNConnectionErrorServerCertificateExpired:
		return "NEVPNConnectionErrorServerCertificateExpired"
	case NEVPNConnectionErrorServerCertificateInvalid:
		return "NEVPNConnectionErrorServerCertificateInvalid"
	case NEVPNConnectionErrorServerCertificateNotYetValid:
		return "NEVPNConnectionErrorServerCertificateNotYetValid"
	case NEVPNConnectionErrorServerDead:
		return "NEVPNConnectionErrorServerDead"
	case NEVPNConnectionErrorServerDisconnected:
		return "NEVPNConnectionErrorServerDisconnected"
	case NEVPNConnectionErrorServerNotResponding:
		return "NEVPNConnectionErrorServerNotResponding"
	case NEVPNConnectionErrorUnrecoverableNetworkChange:
		return "NEVPNConnectionErrorUnrecoverableNetworkChange"
	default:
		return fmt.Sprintf("NEVPNConnectionError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNError-swift.struct/Code
type NEVPNError int

const (
	// NEVPNErrorConfigurationDisabled: An error code indicating the VPN configuration associated with the VPN manager isn’t enabled.
	NEVPNErrorConfigurationDisabled NEVPNError = 2
	// NEVPNErrorConfigurationInvalid: An error code indicating the VPN configuration associated with the VPN manager object is invalid.
	NEVPNErrorConfigurationInvalid NEVPNError = 1
	// NEVPNErrorConfigurationReadWriteFailed: An error code that indicates an error occurred while reading or writing the Network Extension preferences.
	NEVPNErrorConfigurationReadWriteFailed NEVPNError = 5
	// NEVPNErrorConfigurationStale: An error code that indicates another process modfied the VPN configuration since the last time the app loaded the configuration.
	NEVPNErrorConfigurationStale NEVPNError = 4
	// NEVPNErrorConfigurationUnknown: An error code that indicates that unspecified error occurred.
	NEVPNErrorConfigurationUnknown NEVPNError = 6
	// NEVPNErrorConnectionFailed: The connection to the VPN server failed.
	NEVPNErrorConnectionFailed NEVPNError = 3
)

func (e NEVPNError) String() string {
	switch e {
	case NEVPNErrorConfigurationDisabled:
		return "NEVPNErrorConfigurationDisabled"
	case NEVPNErrorConfigurationInvalid:
		return "NEVPNErrorConfigurationInvalid"
	case NEVPNErrorConfigurationReadWriteFailed:
		return "NEVPNErrorConfigurationReadWriteFailed"
	case NEVPNErrorConfigurationStale:
		return "NEVPNErrorConfigurationStale"
	case NEVPNErrorConfigurationUnknown:
		return "NEVPNErrorConfigurationUnknown"
	case NEVPNErrorConnectionFailed:
		return "NEVPNErrorConnectionFailed"
	default:
		return fmt.Sprintf("NEVPNError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEAuthenticationMethod
type NEVPNIKEAuthenticationMethod int

const (
	// NEVPNIKEAuthenticationMethodCertificate: Use a certificate and private key as the authentication credential.
	NEVPNIKEAuthenticationMethodCertificate NEVPNIKEAuthenticationMethod = 1
	// NEVPNIKEAuthenticationMethodNone: Do not authenticate with the IPSec server.
	NEVPNIKEAuthenticationMethodNone NEVPNIKEAuthenticationMethod = 0
	// NEVPNIKEAuthenticationMethodSharedSecret: Use a shared secret as the authentication credential.
	NEVPNIKEAuthenticationMethodSharedSecret NEVPNIKEAuthenticationMethod = 2
)

func (e NEVPNIKEAuthenticationMethod) String() string {
	switch e {
	case NEVPNIKEAuthenticationMethodCertificate:
		return "NEVPNIKEAuthenticationMethodCertificate"
	case NEVPNIKEAuthenticationMethodNone:
		return "NEVPNIKEAuthenticationMethodNone"
	case NEVPNIKEAuthenticationMethodSharedSecret:
		return "NEVPNIKEAuthenticationMethodSharedSecret"
	default:
		return fmt.Sprintf("NEVPNIKEAuthenticationMethod(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2CertificateType
type NEVPNIKEv2CertificateType int

const (
	// NEVPNIKEv2CertificateTypeECDSA256: The ECDSA with p-256 curve certificate type.
	NEVPNIKEv2CertificateTypeECDSA256 NEVPNIKEv2CertificateType = 2
	// NEVPNIKEv2CertificateTypeECDSA384: The ECDSA with p-384 curve certificate type.
	NEVPNIKEv2CertificateTypeECDSA384 NEVPNIKEv2CertificateType = 3
	// NEVPNIKEv2CertificateTypeECDSA521: The ECDSA with p-521 curve certificate type.
	NEVPNIKEv2CertificateTypeECDSA521 NEVPNIKEv2CertificateType = 4
	// NEVPNIKEv2CertificateTypeEd25519: The Edwards 25519 curve certificate type.
	NEVPNIKEv2CertificateTypeEd25519 NEVPNIKEv2CertificateType = 5
	// NEVPNIKEv2CertificateTypeRSA: The RSA certificate type.
	NEVPNIKEv2CertificateTypeRSA NEVPNIKEv2CertificateType = 1
	// NEVPNIKEv2CertificateTypeRSAPSS: The RSA-PSS certificate type.
	NEVPNIKEv2CertificateTypeRSAPSS NEVPNIKEv2CertificateType = 6
)

func (e NEVPNIKEv2CertificateType) String() string {
	switch e {
	case NEVPNIKEv2CertificateTypeECDSA256:
		return "NEVPNIKEv2CertificateTypeECDSA256"
	case NEVPNIKEv2CertificateTypeECDSA384:
		return "NEVPNIKEv2CertificateTypeECDSA384"
	case NEVPNIKEv2CertificateTypeECDSA521:
		return "NEVPNIKEv2CertificateTypeECDSA521"
	case NEVPNIKEv2CertificateTypeEd25519:
		return "NEVPNIKEv2CertificateTypeEd25519"
	case NEVPNIKEv2CertificateTypeRSA:
		return "NEVPNIKEv2CertificateTypeRSA"
	case NEVPNIKEv2CertificateTypeRSAPSS:
		return "NEVPNIKEv2CertificateTypeRSAPSS"
	default:
		return fmt.Sprintf("NEVPNIKEv2CertificateType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DeadPeerDetectionRate
type NEVPNIKEv2DeadPeerDetectionRate int

const (
	// NEVPNIKEv2DeadPeerDetectionRateHigh: Run dead peer detection once every 1 minute.
	NEVPNIKEv2DeadPeerDetectionRateHigh NEVPNIKEv2DeadPeerDetectionRate = 3
	// NEVPNIKEv2DeadPeerDetectionRateLow: Run dead peer detection once every 30 minutes.
	NEVPNIKEv2DeadPeerDetectionRateLow NEVPNIKEv2DeadPeerDetectionRate = 1
	// NEVPNIKEv2DeadPeerDetectionRateMedium: Run dead peer detection once every 10 minutes.
	NEVPNIKEv2DeadPeerDetectionRateMedium NEVPNIKEv2DeadPeerDetectionRate = 2
	// NEVPNIKEv2DeadPeerDetectionRateNone: Do not perform dead peer detection.
	NEVPNIKEv2DeadPeerDetectionRateNone NEVPNIKEv2DeadPeerDetectionRate = 0
)

func (e NEVPNIKEv2DeadPeerDetectionRate) String() string {
	switch e {
	case NEVPNIKEv2DeadPeerDetectionRateHigh:
		return "NEVPNIKEv2DeadPeerDetectionRateHigh"
	case NEVPNIKEv2DeadPeerDetectionRateLow:
		return "NEVPNIKEv2DeadPeerDetectionRateLow"
	case NEVPNIKEv2DeadPeerDetectionRateMedium:
		return "NEVPNIKEv2DeadPeerDetectionRateMedium"
	case NEVPNIKEv2DeadPeerDetectionRateNone:
		return "NEVPNIKEv2DeadPeerDetectionRateNone"
	default:
		return fmt.Sprintf("NEVPNIKEv2DeadPeerDetectionRate(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup
type NEVPNIKEv2DiffieHellmanGroup int

const (
	// NEVPNIKEv2DiffieHellmanGroup1: Diffie Hellman group 1 (768-bit modular exponential [MODP]).
	NEVPNIKEv2DiffieHellmanGroup1 NEVPNIKEv2DiffieHellmanGroup = 1
	// NEVPNIKEv2DiffieHellmanGroup14: Diffie Hellman group 14 (2048-bit modular exponential [MODP]).
	NEVPNIKEv2DiffieHellmanGroup14 NEVPNIKEv2DiffieHellmanGroup = 14
	// NEVPNIKEv2DiffieHellmanGroup15: Diffie Hellman group 15 (3072-bit modular exponential [MODP]).
	NEVPNIKEv2DiffieHellmanGroup15 NEVPNIKEv2DiffieHellmanGroup = 15
	// NEVPNIKEv2DiffieHellmanGroup16: Diffie Hellman group 16 (4096-bit modular exponential [MODP]).
	NEVPNIKEv2DiffieHellmanGroup16 NEVPNIKEv2DiffieHellmanGroup = 16
	// NEVPNIKEv2DiffieHellmanGroup17: Diffie Hellman group 17 (6144-bit modular exponential [MODP]).
	NEVPNIKEv2DiffieHellmanGroup17 NEVPNIKEv2DiffieHellmanGroup = 17
	// NEVPNIKEv2DiffieHellmanGroup18: Diffie Hellman group 18 (8192-bit modular exponential [MODP]).
	NEVPNIKEv2DiffieHellmanGroup18 NEVPNIKEv2DiffieHellmanGroup = 18
	// NEVPNIKEv2DiffieHellmanGroup19: Diffie Hellman group 19 (256-bit random elliptic curve group over GF[P] [ECP]).
	NEVPNIKEv2DiffieHellmanGroup19 NEVPNIKEv2DiffieHellmanGroup = 19
	// NEVPNIKEv2DiffieHellmanGroup2: Diffie Hellman group 2 (1024-bit modular exponential [MODP]).
	NEVPNIKEv2DiffieHellmanGroup2 NEVPNIKEv2DiffieHellmanGroup = 2
	// NEVPNIKEv2DiffieHellmanGroup20: Diffie Hellman group 20 (384-bit random elliptic curve group over GF[P] [ECP]).
	NEVPNIKEv2DiffieHellmanGroup20 NEVPNIKEv2DiffieHellmanGroup = 20
	// NEVPNIKEv2DiffieHellmanGroup21: Diffie Hellman group 21 (521-bit random elliptic curve group over GF[P] [ECP]).
	NEVPNIKEv2DiffieHellmanGroup21 NEVPNIKEv2DiffieHellmanGroup = 21
	// NEVPNIKEv2DiffieHellmanGroup31: Diffie Hellman group 31 (Curve 25519).
	NEVPNIKEv2DiffieHellmanGroup31 NEVPNIKEv2DiffieHellmanGroup = 31
	// NEVPNIKEv2DiffieHellmanGroup32: Diffie Hellman group 32 (Curve 448).
	NEVPNIKEv2DiffieHellmanGroup32 NEVPNIKEv2DiffieHellmanGroup = 32
	// NEVPNIKEv2DiffieHellmanGroup5: Diffie Hellman group 5 (1536-bit modular exponential [MODP]).
	NEVPNIKEv2DiffieHellmanGroup5 NEVPNIKEv2DiffieHellmanGroup = 5
	// NEVPNIKEv2DiffieHellmanGroupInvalid: A value indicating the group is not a valid Diffie-Hellman group.
	NEVPNIKEv2DiffieHellmanGroupInvalid NEVPNIKEv2DiffieHellmanGroup = 0
)

func (e NEVPNIKEv2DiffieHellmanGroup) String() string {
	switch e {
	case NEVPNIKEv2DiffieHellmanGroup1:
		return "NEVPNIKEv2DiffieHellmanGroup1"
	case NEVPNIKEv2DiffieHellmanGroup14:
		return "NEVPNIKEv2DiffieHellmanGroup14"
	case NEVPNIKEv2DiffieHellmanGroup15:
		return "NEVPNIKEv2DiffieHellmanGroup15"
	case NEVPNIKEv2DiffieHellmanGroup16:
		return "NEVPNIKEv2DiffieHellmanGroup16"
	case NEVPNIKEv2DiffieHellmanGroup17:
		return "NEVPNIKEv2DiffieHellmanGroup17"
	case NEVPNIKEv2DiffieHellmanGroup18:
		return "NEVPNIKEv2DiffieHellmanGroup18"
	case NEVPNIKEv2DiffieHellmanGroup19:
		return "NEVPNIKEv2DiffieHellmanGroup19"
	case NEVPNIKEv2DiffieHellmanGroup2:
		return "NEVPNIKEv2DiffieHellmanGroup2"
	case NEVPNIKEv2DiffieHellmanGroup20:
		return "NEVPNIKEv2DiffieHellmanGroup20"
	case NEVPNIKEv2DiffieHellmanGroup21:
		return "NEVPNIKEv2DiffieHellmanGroup21"
	case NEVPNIKEv2DiffieHellmanGroup31:
		return "NEVPNIKEv2DiffieHellmanGroup31"
	case NEVPNIKEv2DiffieHellmanGroup32:
		return "NEVPNIKEv2DiffieHellmanGroup32"
	case NEVPNIKEv2DiffieHellmanGroup5:
		return "NEVPNIKEv2DiffieHellmanGroup5"
	case NEVPNIKEv2DiffieHellmanGroupInvalid:
		return "NEVPNIKEv2DiffieHellmanGroupInvalid"
	default:
		return fmt.Sprintf("NEVPNIKEv2DiffieHellmanGroup(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2EncryptionAlgorithm
type NEVPNIKEv2EncryptionAlgorithm int

const (
	// NEVPNIKEv2EncryptionAlgorithm3DES: Triple Data Encryption Algorithm (aka 3DES)
	NEVPNIKEv2EncryptionAlgorithm3DES NEVPNIKEv2EncryptionAlgorithm = 2
	// NEVPNIKEv2EncryptionAlgorithmAES256: Advanced Encryption Standard 256 bit (AES256).
	NEVPNIKEv2EncryptionAlgorithmAES256 NEVPNIKEv2EncryptionAlgorithm = 4
	// NEVPNIKEv2EncryptionAlgorithmAES256GCM: Advanced Encryption Standard 256-bit Galois/Counter Mode (AES256GCM).
	NEVPNIKEv2EncryptionAlgorithmAES256GCM NEVPNIKEv2EncryptionAlgorithm = 6
	// NEVPNIKEv2EncryptionAlgorithmChaCha20Poly1305: ChaCha20 and Poly1305 (ChaCha20Poly1305).
	NEVPNIKEv2EncryptionAlgorithmChaCha20Poly1305 NEVPNIKEv2EncryptionAlgorithm = 7
	// NEVPNIKEv2EncryptionAlgorithmDES: Data Encryption Standard (DES)
	NEVPNIKEv2EncryptionAlgorithmDES NEVPNIKEv2EncryptionAlgorithm = 1
	// Deprecated.
	NEVPNIKEv2EncryptionAlgorithmAES128 NEVPNIKEv2EncryptionAlgorithm = 3
	// Deprecated.
	NEVPNIKEv2EncryptionAlgorithmAES128GCM NEVPNIKEv2EncryptionAlgorithm = 5
)

func (e NEVPNIKEv2EncryptionAlgorithm) String() string {
	switch e {
	case NEVPNIKEv2EncryptionAlgorithm3DES:
		return "NEVPNIKEv2EncryptionAlgorithm3DES"
	case NEVPNIKEv2EncryptionAlgorithmAES256:
		return "NEVPNIKEv2EncryptionAlgorithmAES256"
	case NEVPNIKEv2EncryptionAlgorithmAES256GCM:
		return "NEVPNIKEv2EncryptionAlgorithmAES256GCM"
	case NEVPNIKEv2EncryptionAlgorithmChaCha20Poly1305:
		return "NEVPNIKEv2EncryptionAlgorithmChaCha20Poly1305"
	case NEVPNIKEv2EncryptionAlgorithmDES:
		return "NEVPNIKEv2EncryptionAlgorithmDES"
	case NEVPNIKEv2EncryptionAlgorithmAES128:
		return "NEVPNIKEv2EncryptionAlgorithmAES128"
	case NEVPNIKEv2EncryptionAlgorithmAES128GCM:
		return "NEVPNIKEv2EncryptionAlgorithmAES128GCM"
	default:
		return fmt.Sprintf("NEVPNIKEv2EncryptionAlgorithm(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2IntegrityAlgorithm
type NEVPNIKEv2IntegrityAlgorithm int

const (
	// NEVPNIKEv2IntegrityAlgorithmSHA160: SHA-1 160-bit.
	NEVPNIKEv2IntegrityAlgorithmSHA160 NEVPNIKEv2IntegrityAlgorithm = 2
	// NEVPNIKEv2IntegrityAlgorithmSHA256: SHA-2 256-bit.
	NEVPNIKEv2IntegrityAlgorithmSHA256 NEVPNIKEv2IntegrityAlgorithm = 3
	// NEVPNIKEv2IntegrityAlgorithmSHA384: SHA-2 384-bit.
	NEVPNIKEv2IntegrityAlgorithmSHA384 NEVPNIKEv2IntegrityAlgorithm = 4
	// NEVPNIKEv2IntegrityAlgorithmSHA512: SHA-2 512-bit.
	NEVPNIKEv2IntegrityAlgorithmSHA512 NEVPNIKEv2IntegrityAlgorithm = 5
	// NEVPNIKEv2IntegrityAlgorithmSHA96: SHA-1 96-bit.
	NEVPNIKEv2IntegrityAlgorithmSHA96 NEVPNIKEv2IntegrityAlgorithm = 1
)

func (e NEVPNIKEv2IntegrityAlgorithm) String() string {
	switch e {
	case NEVPNIKEv2IntegrityAlgorithmSHA160:
		return "NEVPNIKEv2IntegrityAlgorithmSHA160"
	case NEVPNIKEv2IntegrityAlgorithmSHA256:
		return "NEVPNIKEv2IntegrityAlgorithmSHA256"
	case NEVPNIKEv2IntegrityAlgorithmSHA384:
		return "NEVPNIKEv2IntegrityAlgorithmSHA384"
	case NEVPNIKEv2IntegrityAlgorithmSHA512:
		return "NEVPNIKEv2IntegrityAlgorithmSHA512"
	case NEVPNIKEv2IntegrityAlgorithmSHA96:
		return "NEVPNIKEv2IntegrityAlgorithmSHA96"
	default:
		return fmt.Sprintf("NEVPNIKEv2IntegrityAlgorithm(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PostQuantumKeyExchangeMethod
type NEVPNIKEv2PostQuantumKeyExchangeMethod int

const (
	// NEVPNIKEv2PostQuantumKeyExchangeMethod36: Instructs the server to use the ML-KEM-768 key exchange method.
	NEVPNIKEv2PostQuantumKeyExchangeMethod36 NEVPNIKEv2PostQuantumKeyExchangeMethod = 36
	// NEVPNIKEv2PostQuantumKeyExchangeMethod37: Instructs the server to use the ML-KEM-1024 key exchange method.
	NEVPNIKEv2PostQuantumKeyExchangeMethod37 NEVPNIKEv2PostQuantumKeyExchangeMethod = 37
	// NEVPNIKEv2PostQuantumKeyExchangeMethodNone: Instructs the server not to use a quantum-secure key exchange method.
	NEVPNIKEv2PostQuantumKeyExchangeMethodNone NEVPNIKEv2PostQuantumKeyExchangeMethod = 0
)

func (e NEVPNIKEv2PostQuantumKeyExchangeMethod) String() string {
	switch e {
	case NEVPNIKEv2PostQuantumKeyExchangeMethod36:
		return "NEVPNIKEv2PostQuantumKeyExchangeMethod36"
	case NEVPNIKEv2PostQuantumKeyExchangeMethod37:
		return "NEVPNIKEv2PostQuantumKeyExchangeMethod37"
	case NEVPNIKEv2PostQuantumKeyExchangeMethodNone:
		return "NEVPNIKEv2PostQuantumKeyExchangeMethodNone"
	default:
		return fmt.Sprintf("NEVPNIKEv2PostQuantumKeyExchangeMethod(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2TLSVersion
type NEVPNIKEv2TLSVersion int

const (
	// NEVPNIKEv2TLSVersion1_0: A value to use TLS version 1.0.
	NEVPNIKEv2TLSVersion1_0 NEVPNIKEv2TLSVersion = 1
	// NEVPNIKEv2TLSVersion1_1: A value to use TLS version 1.1.
	NEVPNIKEv2TLSVersion1_1 NEVPNIKEv2TLSVersion = 2
	// NEVPNIKEv2TLSVersion1_2: A value to use TLS version 1.2.
	NEVPNIKEv2TLSVersion1_2 NEVPNIKEv2TLSVersion = 3
	// NEVPNIKEv2TLSVersionDefault: A value to use the default TLS configuration.
	NEVPNIKEv2TLSVersionDefault NEVPNIKEv2TLSVersion = 0
)

func (e NEVPNIKEv2TLSVersion) String() string {
	switch e {
	case NEVPNIKEv2TLSVersion1_0:
		return "NEVPNIKEv2TLSVersion1_0"
	case NEVPNIKEv2TLSVersion1_1:
		return "NEVPNIKEv2TLSVersion1_1"
	case NEVPNIKEv2TLSVersion1_2:
		return "NEVPNIKEv2TLSVersion1_2"
	case NEVPNIKEv2TLSVersionDefault:
		return "NEVPNIKEv2TLSVersionDefault"
	default:
		return fmt.Sprintf("NEVPNIKEv2TLSVersion(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatus
type NEVPNStatus int

const (
	// NEVPNStatusConnected: The VPN is connected.
	NEVPNStatusConnected NEVPNStatus = 3
	// NEVPNStatusConnecting: The VPN is in the process of connecting.
	NEVPNStatusConnecting NEVPNStatus = 2
	// NEVPNStatusDisconnected: The VPN is disconnected.
	NEVPNStatusDisconnected NEVPNStatus = 1
	// NEVPNStatusDisconnecting: The VPN is in the process of disconnecting.
	NEVPNStatusDisconnecting NEVPNStatus = 5
	// NEVPNStatusInvalid: The associated VPN configuration doesn’t exist in the Network Extension preferences or isn’t enabled.
	NEVPNStatusInvalid NEVPNStatus = 0
	// NEVPNStatusReasserting: The VPN is in the process of reconnecting.
	NEVPNStatusReasserting NEVPNStatus = 4
)

func (e NEVPNStatus) String() string {
	switch e {
	case NEVPNStatusConnected:
		return "NEVPNStatusConnected"
	case NEVPNStatusConnecting:
		return "NEVPNStatusConnecting"
	case NEVPNStatusDisconnected:
		return "NEVPNStatusDisconnected"
	case NEVPNStatusDisconnecting:
		return "NEVPNStatusDisconnecting"
	case NEVPNStatusInvalid:
		return "NEVPNStatusInvalid"
	case NEVPNStatusReasserting:
		return "NEVPNStatusReasserting"
	default:
		return fmt.Sprintf("NEVPNStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NWPathStatus
type NWPathStatus int

const (
	// Deprecated.
	NWPathStatusInvalid NWPathStatus = 0
	// Deprecated.
	NWPathStatusSatisfiable NWPathStatus = 3
	// Deprecated.
	NWPathStatusSatisfied NWPathStatus = 1
	// Deprecated.
	NWPathStatusUnsatisfied NWPathStatus = 2
)

func (e NWPathStatus) String() string {
	switch e {
	case NWPathStatusInvalid:
		return "NWPathStatusInvalid"
	case NWPathStatusSatisfiable:
		return "NWPathStatusSatisfiable"
	case NWPathStatusSatisfied:
		return "NWPathStatusSatisfied"
	case NWPathStatusUnsatisfied:
		return "NWPathStatusUnsatisfied"
	default:
		return fmt.Sprintf("NWPathStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnectionState
type NWTCPConnectionState int

const (
	// Deprecated.
	NWTCPConnectionStateCancelled NWTCPConnectionState = 5
	// Deprecated.
	NWTCPConnectionStateConnected NWTCPConnectionState = 3
	// Deprecated.
	NWTCPConnectionStateConnecting NWTCPConnectionState = 1
	// Deprecated.
	NWTCPConnectionStateDisconnected NWTCPConnectionState = 4
	// Deprecated.
	NWTCPConnectionStateInvalid NWTCPConnectionState = 0
	// Deprecated.
	NWTCPConnectionStateWaiting NWTCPConnectionState = 2
)

func (e NWTCPConnectionState) String() string {
	switch e {
	case NWTCPConnectionStateCancelled:
		return "NWTCPConnectionStateCancelled"
	case NWTCPConnectionStateConnected:
		return "NWTCPConnectionStateConnected"
	case NWTCPConnectionStateConnecting:
		return "NWTCPConnectionStateConnecting"
	case NWTCPConnectionStateDisconnected:
		return "NWTCPConnectionStateDisconnected"
	case NWTCPConnectionStateInvalid:
		return "NWTCPConnectionStateInvalid"
	case NWTCPConnectionStateWaiting:
		return "NWTCPConnectionStateWaiting"
	default:
		return fmt.Sprintf("NWTCPConnectionState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NetworkExtension/NWUDPSessionState
type NWUDPSessionState int

const (
	// Deprecated.
	NWUDPSessionStateCancelled NWUDPSessionState = 5
	// Deprecated.
	NWUDPSessionStateFailed NWUDPSessionState = 4
	// Deprecated.
	NWUDPSessionStateInvalid NWUDPSessionState = 0
	// Deprecated.
	NWUDPSessionStatePreparing NWUDPSessionState = 2
	// Deprecated.
	NWUDPSessionStateReady NWUDPSessionState = 3
	// Deprecated.
	NWUDPSessionStateWaiting NWUDPSessionState = 1
)

func (e NWUDPSessionState) String() string {
	switch e {
	case NWUDPSessionStateCancelled:
		return "NWUDPSessionStateCancelled"
	case NWUDPSessionStateFailed:
		return "NWUDPSessionStateFailed"
	case NWUDPSessionStateInvalid:
		return "NWUDPSessionStateInvalid"
	case NWUDPSessionStatePreparing:
		return "NWUDPSessionStatePreparing"
	case NWUDPSessionStateReady:
		return "NWUDPSessionStateReady"
	case NWUDPSessionStateWaiting:
		return "NWUDPSessionStateWaiting"
	default:
		return fmt.Sprintf("NWUDPSessionState(%d)", e)
	}
}
