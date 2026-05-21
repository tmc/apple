// Code generated from Apple documentation. DO NOT EDIT.

package networkextension

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
	if handler == nil {
		return 0, func() {}
	}
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
	if handler == nil {
		return 0, func() {}
	}
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}

// NEAppProxyProviderManagerArrayErrorHandler handles A block that takes an NSArray of NEAppProxyProviderManager objects, and an NSError object.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NEAppProxyProviderManager.LoadAllFromPreferencesWithCompletionHandler]
type NEAppProxyProviderManagerArrayErrorHandler = func(*[]NEAppProxyProviderManager, error)

// NewNEAppProxyProviderManagerArrayErrorBlock wraps a Go [NEAppProxyProviderManagerArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NEAppProxyProviderManager.LoadAllFromPreferencesWithCompletionHandler]
func NewNEAppProxyProviderManagerArrayErrorBlock(handler NEAppProxyProviderManagerArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]NEAppProxyProviderManager
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NEAppProxyProviderManager, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NEAppProxyProviderManagerFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NEPacketArrayHandler is the signature for a completion handler block.
//
// Used by:
//   - [NEPacketTunnelFlow.ReadPacketObjectsWithCompletionHandler]
type NEPacketArrayHandler = func(*[]NEPacket)

// NewNEPacketArrayBlock wraps a Go [NEPacketArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NEPacketTunnelFlow.ReadPacketObjectsWithCompletionHandler]
func NewNEPacketArrayBlock(handler NEPacketArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]NEPacket
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NEPacket, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NEPacketFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// NERelayManagerArrayErrorHandler handles A block that receives an array of NERelayManager objects.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NERelayManager.LoadAllManagersFromPreferencesWithCompletionHandler]
type NERelayManagerArrayErrorHandler = func(*[]NERelayManager, error)

// NewNERelayManagerArrayErrorBlock wraps a Go [NERelayManagerArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NERelayManager.LoadAllManagersFromPreferencesWithCompletionHandler]
func NewNERelayManagerArrayErrorBlock(handler NERelayManagerArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]NERelayManager
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NERelayManager, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NERelayManagerFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NETransparentProxyManagerArrayErrorHandler handles A Swift closure or an ObjectiveC block that receives as parameters an array of transparent proxy manager instances loaded from disk and an error.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NETransparentProxyManager.LoadAllFromPreferencesWithCompletionHandler]
type NETransparentProxyManagerArrayErrorHandler = func(*[]NETransparentProxyManager, error)

// NewNETransparentProxyManagerArrayErrorBlock wraps a Go [NETransparentProxyManagerArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NETransparentProxyManager.LoadAllFromPreferencesWithCompletionHandler]
func NewNETransparentProxyManagerArrayErrorBlock(handler NETransparentProxyManagerArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]NETransparentProxyManager
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NETransparentProxyManager, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NETransparentProxyManagerFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NETunnelProviderManagerArrayErrorHandler handles A block that takes an NSArray of [NETunnelProviderManager] objects, and an NSError object.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NETunnelProviderManager.LoadAllFromPreferencesWithCompletionHandler]
type NETunnelProviderManagerArrayErrorHandler = func(*[]NETunnelProviderManager, error)

// NewNETunnelProviderManagerArrayErrorBlock wraps a Go [NETunnelProviderManagerArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NETunnelProviderManager.LoadAllFromPreferencesWithCompletionHandler]
func NewNETunnelProviderManagerArrayErrorBlock(handler NETunnelProviderManagerArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]NETunnelProviderManager
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NETunnelProviderManager, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NETunnelProviderManagerFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NEURLFilterVerdict) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSDataArrayErrorHandler handles A handler called when datagrams have been read, or when an error has occurred.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NWUDPSession.SetReadHandlerMaxDatagrams]
type NSDataArrayErrorHandler = func(*[]foundation.NSData, error)

// NewNSDataArrayErrorBlock wraps a Go [NSDataArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NWUDPSession.SetReadHandlerMaxDatagrams]
func NewNSDataArrayErrorBlock(handler NSDataArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]foundation.NSData, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = foundation.NSDataFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NSDataArrayNSNumberArrayHandler handles A Swift closure or an ObjectiveC block that runs when some packets are read from the TUN interface.
//
// Used by:
//   - [NEPacketTunnelFlow.ReadPacketsWithCompletionHandler]
type NSDataArrayNSNumberArrayHandler = func(*[]foundation.NSData, *[]foundation.NSNumber)

// NewNSDataArrayNSNumberArrayBlock wraps a Go [NSDataArrayNSNumberArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NEPacketTunnelFlow.ReadPacketsWithCompletionHandler]
func NewNSDataArrayNSNumberArrayBlock(handler NSDataArrayNSNumberArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *[]foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]foundation.NSData, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = foundation.NSDataFromID(item.GetID())
			}
			result = &res
		}
		var extra0 *[]foundation.NSNumber
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(extra0ID)
			count := obj.Count()
			res := make([]foundation.NSNumber, count)
			for j := uint(0); j < count; j++ {
				item := obj.ObjectAtIndex(j)
				res[j] = foundation.NSNumberFromID(item.GetID())
			}
			extra0 = &res
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSDataArrayNWEndpointArrayErrorHandler handles A block that will be executed by the system on an internal system thread when datagrams have been read from the flow.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NEAppProxyUDPFlow.ReadDatagramsWithCompletionHandler]
type NSDataArrayNWEndpointArrayErrorHandler = func(*[]foundation.NSData, *[]NWEndpoint, error)

// NewNSDataArrayNWEndpointArrayErrorBlock wraps a Go [NSDataArrayNWEndpointArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NEAppProxyUDPFlow.ReadDatagramsWithCompletionHandler]
func NewNSDataArrayNWEndpointArrayErrorBlock(handler NSDataArrayNWEndpointArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *[]foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]foundation.NSData, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = foundation.NSDataFromID(item.GetID())
			}
			result = &res
		}
		var extra0 *[]NWEndpoint
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(extra0ID)
			count := obj.Count()
			res := make([]NWEndpoint, count)
			for j := uint(0); j < count; j++ {
				item := obj.ObjectAtIndex(j)
				res[j] = NWEndpointFromID(item.GetID())
			}
			extra0 = &res
		}
		handler(result, extra0, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NSDataArrayObjectArrayErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [NEAppProxyUDPFlow.ReadDatagramsAndFlowEndpointsWithCompletionHandler]
type NSDataArrayObjectArrayErrorHandler = func(*[]foundation.NSData, *[]objectivec.Object, error)

// NewNSDataArrayObjectArrayErrorBlock wraps a Go [NSDataArrayObjectArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NEAppProxyUDPFlow.ReadDatagramsAndFlowEndpointsWithCompletionHandler]
func NewNSDataArrayObjectArrayErrorBlock(handler NSDataArrayObjectArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result *[]foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]foundation.NSData, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = foundation.NSDataFromID(item.GetID())
			}
			result = &res
		}
		var extra0 *[]objectivec.Object
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(extra0ID)
			count := obj.Count()
			res := make([]objectivec.Object, count)
			for j := uint(0); j < count; j++ {
				item := obj.ObjectAtIndex(j)
				res[j] = objectivec.ObjectFromID(item.GetID())
			}
			extra0 = &res
		}
		handler(result, extra0, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// SecIdentityRefArrayHandler handles The completion handler for passing an identity and certificate chain to the connection.
//
// Used by:
//   - [NWTCPConnectionAuthenticationDelegate.ProvideIdentityForConnectionCompletionHandler]
type SecIdentityRefArrayHandler = func(security.SecIdentity, *foundation.NSArray)

// NewSecIdentityRefArrayBlock wraps a Go [SecIdentityRefArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NWTCPConnectionAuthenticationDelegate.ProvideIdentityForConnectionCompletionHandler]
func NewSecIdentityRefArrayBlock(handler SecIdentityRefArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0ID objc.ID) {
		primitive := security.SecIdentity{ID: primitiveID}
		var extra0 *foundation.NSArray
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := foundation.NSArrayFromID(extra0ID)
			extra0 = &v
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// SecTrustRefHandler handles The completion handler for passing the SecTrust object to the connection.
//
// Used by:
//   - [NWTCPConnectionAuthenticationDelegate.EvaluateTrustForConnectionPeerCertificateChainCompletionHandler]
type SecTrustRefHandler = func(security.SecTrust)

// NewSecTrustRefBlock wraps a Go [SecTrustRefHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NWTCPConnectionAuthenticationDelegate.EvaluateTrustForConnectionPeerCertificateChainCompletionHandler]
func NewSecTrustRefBlock(handler SecTrustRefHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal objc.ID) {
		handler(security.SecTrust{ID: primitiveVal})
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles A block that must be executed when the proxy is fully stopped.
//
// Used by:
//   - [NEAppProxyProvider.StopProxyWithReasonCompletionHandler]
//   - [NEDNSProxyProvider.StopProxyWithReasonCompletionHandler]
//   - [NEFilterProvider.StopFilterWithReasonCompletionHandler]
//   - [NEPacketTunnelProvider.StopTunnelWithReasonCompletionHandler]
//   - [NEProvider.SleepWithCompletionHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NEAppProxyProvider.StopProxyWithReasonCompletionHandler]
//   - [NEDNSProxyProvider.StopProxyWithReasonCompletionHandler]
//   - [NEFilterProvider.StopFilterWithReasonCompletionHandler]
//   - [NEPacketTunnelProvider.StopTunnelWithReasonCompletionHandler]
//   - [NEProvider.SleepWithCompletionHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
