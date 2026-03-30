// Code generated from Apple documentation. DO NOT EDIT.

package networkextension

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/security"
)

// DataErrorHandler handles A block that will be executed by the system on an internal system thread when some data is read from the flow.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NEAppProxyTCPFlow.ReadDataWithCompletionHandler]
//   - [NWTCPConnection.ReadLengthCompletionHandler]
//   - [NWTCPConnection.ReadMinimumLengthMaximumLengthCompletionHandler]
type DataErrorHandler = func(*foundation.NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NEAppProxyTCPFlow.ReadDataWithCompletionHandler]
//   - [NWTCPConnection.ReadLengthCompletionHandler]
//   - [NWTCPConnection.ReadMinimumLengthMaximumLengthCompletionHandler]
func NewDataErrorBlock(handler DataErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// DataHandler handles A block to be executed by the Tunnel Provider when it is finished handling the message.
//
// Used by:
//   - [NETunnelProvider.HandleAppMessageCompletionHandler]
//   - [NETunnelProviderSession.SendProviderMessageReturnErrorResponseHandler]
type DataHandler = func(*foundation.NSData)

// NewDataBlock wraps a Go [DataHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NETunnelProvider.HandleAppMessageCompletionHandler]
//   - [NETunnelProviderSession.SendProviderMessageReturnErrorResponseHandler]
func NewDataBlock(handler DataHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles Called when the open operation is complete.
//   - error: A `nil` value indicates the flow opened successfully. A non-`nil` value indicates the flow could not be opened. See [NEAppProxyFlowError](<doc://com.apple.networkextension/documentation/NetworkExtension/NEAppProxyFlowError-swift.struct>) for a list of expected error codes.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NEAppProxyFlow.OpenWithLocalEndpointCompletionHandler]
//   - [NEAppProxyFlow.OpenWithLocalFlowEndpointCompletionHandler]
//   - [NEAppProxyProvider.StartProxyWithOptionsCompletionHandler]
//   - [NEAppProxyTCPFlow.WriteDataWithCompletionHandler]
//   - [NEAppProxyUDPFlow.WriteDatagramsSentByEndpointsCompletionHandler]
//   - [NEAppProxyUDPFlow.WriteDatagramsSentByFlowEndpointsCompletionHandler]
//   - [NEDNSProxyManager.LoadFromPreferencesWithCompletionHandler]
//   - [NEDNSProxyManager.RemoveFromPreferencesWithCompletionHandler]
//   - [NEDNSProxyManager.SaveToPreferencesWithCompletionHandler]
//   - [NEDNSProxyProvider.StartProxyWithOptionsCompletionHandler]
//   - [NEDNSSettingsManager.LoadFromPreferencesWithCompletionHandler]
//   - [NEDNSSettingsManager.RemoveFromPreferencesWithCompletionHandler]
//   - [NEDNSSettingsManager.SaveToPreferencesWithCompletionHandler]
//   - [NEFilterDataProvider.ApplySettingsCompletionHandler]
//   - [NEFilterManager.LoadFromPreferencesWithCompletionHandler]
//   - [NEFilterManager.RemoveFromPreferencesWithCompletionHandler]
//   - [NEFilterManager.SaveToPreferencesWithCompletionHandler]
//   - [NEFilterProvider.StartFilterWithCompletionHandler]
//   - [NEPacketTunnelProvider.StartTunnelWithOptionsCompletionHandler]
//   - [NERelayManager.GetLastClientErrorsCompletionHandler]
//   - [NERelayManager.LoadFromPreferencesWithCompletionHandler]
//   - [NERelayManager.RemoveFromPreferencesWithCompletionHandler]
//   - [NERelayManager.SaveToPreferencesWithCompletionHandler]
//   - [NETunnelProvider.SetTunnelNetworkSettingsCompletionHandler]
//   - [NEVPNConnection.FetchLastDisconnectErrorWithCompletionHandler]
//   - [NEVPNManager.LoadFromPreferencesWithCompletionHandler]
//   - [NEVPNManager.RemoveFromPreferencesWithCompletionHandler]
//   - [NEVPNManager.SaveToPreferencesWithCompletionHandler]
//   - [NWTCPConnection.WriteCompletionHandler]
//   - [NWUDPSession.WriteDatagramCompletionHandler]
//   - [NWUDPSession.WriteMultipleDatagramsCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NEAppProxyFlow.OpenWithLocalEndpointCompletionHandler]
//   - [NEAppProxyFlow.OpenWithLocalFlowEndpointCompletionHandler]
//   - [NEAppProxyProvider.StartProxyWithOptionsCompletionHandler]
//   - [NEAppProxyTCPFlow.WriteDataWithCompletionHandler]
//   - [NEAppProxyUDPFlow.WriteDatagramsSentByEndpointsCompletionHandler]
//   - [NEAppProxyUDPFlow.WriteDatagramsSentByFlowEndpointsCompletionHandler]
//   - [NEDNSProxyManager.LoadFromPreferencesWithCompletionHandler]
//   - [NEDNSProxyManager.RemoveFromPreferencesWithCompletionHandler]
//   - [NEDNSProxyManager.SaveToPreferencesWithCompletionHandler]
//   - [NEDNSProxyProvider.StartProxyWithOptionsCompletionHandler]
//   - [NEDNSSettingsManager.LoadFromPreferencesWithCompletionHandler]
//   - [NEDNSSettingsManager.RemoveFromPreferencesWithCompletionHandler]
//   - [NEDNSSettingsManager.SaveToPreferencesWithCompletionHandler]
//   - [NEFilterDataProvider.ApplySettingsCompletionHandler]
//   - [NEFilterManager.LoadFromPreferencesWithCompletionHandler]
//   - [NEFilterManager.RemoveFromPreferencesWithCompletionHandler]
//   - [NEFilterManager.SaveToPreferencesWithCompletionHandler]
//   - [NEFilterProvider.StartFilterWithCompletionHandler]
//   - [NEPacketTunnelProvider.StartTunnelWithOptionsCompletionHandler]
//   - [NERelayManager.GetLastClientErrorsCompletionHandler]
//   - [NERelayManager.LoadFromPreferencesWithCompletionHandler]
//   - [NERelayManager.RemoveFromPreferencesWithCompletionHandler]
//   - [NERelayManager.SaveToPreferencesWithCompletionHandler]
//   - [NETunnelProvider.SetTunnelNetworkSettingsCompletionHandler]
//   - [NEVPNConnection.FetchLastDisconnectErrorWithCompletionHandler]
//   - [NEVPNManager.LoadFromPreferencesWithCompletionHandler]
//   - [NEVPNManager.RemoveFromPreferencesWithCompletionHandler]
//   - [NEVPNManager.SaveToPreferencesWithCompletionHandler]
//   - [NWTCPConnection.WriteCompletionHandler]
//   - [NWUDPSession.WriteDatagramCompletionHandler]
//   - [NWUDPSession.WriteMultipleDatagramsCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NEAppProxyProviderManagerArrayErrorHandler handles A block that takes an NSArray of NEAppProxyProviderManager objects, and an NSError object.
//
// Used by:
//   - [NEAppProxyProviderManager.LoadAllFromPreferencesWithCompletionHandler]
type NEAppProxyProviderManagerArrayErrorHandler = func(*[]NEAppProxyProviderManager, error)

// NEPacketArrayHandler is the signature for a completion handler block.
//
// Used by:
//   - [NEPacketTunnelFlow.ReadPacketObjectsWithCompletionHandler]
type NEPacketArrayHandler = func(*[]NEPacket)

// NERelayManagerArrayErrorHandler handles A block that receives an array of NERelayManager objects.
//
// Used by:
//   - [NERelayManager.LoadAllManagersFromPreferencesWithCompletionHandler]
type NERelayManagerArrayErrorHandler = func(*[]NERelayManager, error)

// NETransparentProxyManagerArrayErrorHandler handles A Swift closure or an ObjectiveC block that receives as parameters an array of transparent proxy manager instances loaded from disk and an error.
//
// Used by:
//   - [NETransparentProxyManager.LoadAllFromPreferencesWithCompletionHandler]
type NETransparentProxyManagerArrayErrorHandler = func(*[]NETransparentProxyManager, error)

// NETunnelProviderManagerArrayErrorHandler handles A block that takes an NSArray of [NETunnelProviderManager] objects, and an NSError object.
//
// Used by:
//   - [NETunnelProviderManager.LoadAllFromPreferencesWithCompletionHandler]
type NETunnelProviderManagerArrayErrorHandler = func(*[]NETunnelProviderManager, error)

// NEURLFilterVerdictHandler handles A block that the system calls when it completes validation.
//
// Used by:
//   - [NEURLFilter.VerdictForURLCompletionHandler]
type NEURLFilterVerdictHandler = func(NEURLFilterVerdict)

// NewNEURLFilterVerdictBlock wraps a Go [NEURLFilterVerdictHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NEURLFilter.VerdictForURLCompletionHandler]
func NewNEURLFilterVerdictBlock(handler NEURLFilterVerdictHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal NEURLFilterVerdict) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSDataArrayErrorHandler handles A handler called when datagrams have been read, or when an error has occurred.
//
// Used by:
//   - [NWUDPSession.SetReadHandlerMaxDatagrams]
type NSDataArrayErrorHandler = func(*[]foundation.NSData, error)

// SecTrustRefHandler handles The completion handler for passing the SecTrust object to the connection.
//
// Used by:
//   - [NWTCPConnectionAuthenticationDelegate.EvaluateTrustForConnectionPeerCertificateChainCompletionHandler]
type SecTrustRefHandler = func(security.SecTrustRef)

// NewSecTrustRefBlock wraps a Go [SecTrustRefHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NWTCPConnectionAuthenticationDelegate.EvaluateTrustForConnectionPeerCertificateChainCompletionHandler]
func NewSecTrustRefBlock(handler SecTrustRefHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal security.SecTrustRef) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles A block that must be executed when the proxy is fully stopped.
//
// Used by:
//   - [NEAppProxyProvider.StopProxyWithReasonCompletionHandler]
//   - [NEAppProxyUDPFlow.ReadDatagramsAndFlowEndpointsWithCompletionHandler]
//   - [NEAppProxyUDPFlow.ReadDatagramsWithCompletionHandler]
//   - [NEDNSProxyProvider.StopProxyWithReasonCompletionHandler]
//   - [NEFilterProvider.StopFilterWithReasonCompletionHandler]
//   - [NEPacketTunnelFlow.ReadPacketsWithCompletionHandler]
//   - [NEPacketTunnelProvider.StopTunnelWithReasonCompletionHandler]
//   - [NEProvider.SleepWithCompletionHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NEAppProxyProvider.StopProxyWithReasonCompletionHandler]
//   - [NEAppProxyUDPFlow.ReadDatagramsAndFlowEndpointsWithCompletionHandler]
//   - [NEAppProxyUDPFlow.ReadDatagramsWithCompletionHandler]
//   - [NEDNSProxyProvider.StopProxyWithReasonCompletionHandler]
//   - [NEFilterProvider.StopFilterWithReasonCompletionHandler]
//   - [NEPacketTunnelFlow.ReadPacketsWithCompletionHandler]
//   - [NEPacketTunnelProvider.StopTunnelWithReasonCompletionHandler]
//   - [NEProvider.SleepWithCompletionHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// idArrayHandler handles The completion handler for passing an identity and certificate chain to the connection.
//
// Used by:
//   - [NWTCPConnectionAuthenticationDelegate.ProvideIdentityForConnectionCompletionHandler]
type idArrayHandler = func(security.SecIdentityRef)

// NewidArrayBlock wraps a Go [idArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NWTCPConnectionAuthenticationDelegate.ProvideIdentityForConnectionCompletionHandler]
func NewidArrayBlock(handler idArrayHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal security.SecIdentityRef) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}
