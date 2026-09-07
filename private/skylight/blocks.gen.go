// Code generated from Apple documentation. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// IObjectINSDictionaryHandler handles completion with primitive and object results.
//
// Used by:
//   - [SLSBrightnessControl.RegisterForNotificationsWithBlock]
type IObjectINSDictionaryHandler = func(objectivec.IObject, foundation.INSDictionary)

// NewIObjectINSDictionaryBlock wraps a Go [IObjectINSDictionaryHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SLSBrightnessControl.RegisterForNotificationsWithBlock]
func NewIObjectINSDictionaryBlock(handler IObjectINSDictionaryHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0ID objc.ID) {
		var primitive objectivec.IObject
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			obj := objectivec.ObjectFromID(primitiveID)
			primitive = &obj
		}
		var extra0 foundation.INSDictionary
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			extra0 = foundation.NSDictionaryFromID(extra0ID)
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NumberErrorHandler handles completion with primitive result and optional error.
//
// Used by:
//   - [CPXRemoteViewEventProtocolServerCallsClient.SendEventToHostFullDispatchReply]
type NumberErrorHandler = func(foundation.NSNumber, error)

// NewNumberErrorBlock wraps a Go [NumberErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CPXRemoteViewEventProtocolServerCallsClient.SendEventToHostFullDispatchReply]
func NewNumberErrorBlock(handler NumberErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal foundation.NSNumber, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// SLDataTimelineProcessHandler is the signature for a completion handler block.
//
// Used by:
//   - [SLDataTimelineSessionProcessCollection.ProcessesApplyBlock]
type SLDataTimelineProcessHandler = func(*unsafe.Pointer)

// SLDataTimelineServerSnapshotHandler is the signature for a completion handler block.
//
// Used by:
//   - [SLDataTimelineSnapshotCollection.SnapshotsApplyBlock]
type SLDataTimelineServerSnapshotHandler = func(*unsafe.Pointer)

// SLDataTimelineSessionHandler is the signature for a completion handler block.
//
// Used by:
//   - [SLDataTimelineServerSnapshot.SessionsApplyBlock]
type SLDataTimelineSessionHandler = func(*unsafe.Pointer)

// UnsafePointerHandler handles completion with a primitive value.
//
// Used by:
//   - [SLSDisplayControlClientProtocol.RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClientProtocol.RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClientProtocol.InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClientProtocol.InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClientProtocol.ConfigGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClientProtocol.InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSXPCServiceProtocol.InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock]
type UnsafePointerHandler = func(unsafe.Pointer)

// NewUnsafePointerBlock wraps a Go [UnsafePointerHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SLSDisplayControlClientProtocol.RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClientProtocol.RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClientProtocol.InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClientProtocol.InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClientProtocol.ConfigGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClientProtocol.InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSXPCServiceProtocol.InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock]
func NewUnsafePointerBlock(handler UnsafePointerHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [CPXEventDeferringPolicy.Build]
//   - [CPXRemoteViewEventManager.PassEventUpstreamToHostFullDispatchReply]
//   - [CPXRemoteViewEventManager.SendEventToHostPidFullDispatchReply]
//   - [CPXSymbolicHotKeyRegistrar.RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc]
//   - [CPXSymbolicHotKeyRegistry.RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc]
//   - [ECTestOnlyEventAuthenticationMessage.ValidateWithOptionsAndResultBlock]
//   - [PKGCoreUITransaction.UpdateLayerKeyRendererWork]
//   - [PKGCoreUITransaction._scheduleRendererWorkMainThreadWork]
//   - [PKGCoreUIWork.SetMainThreadWork]
//   - [PKGCoreUIWork.SetRendererWork]
//   - [SLContentStream.CreateScreenshotPropertiesQueueHandlerError]
//   - [SLContentStream.InitWithFilterPropertiesQueueHandlerError]
//   - [SLContentStream.InitWithFilterPropertiesQueueHandler]
//   - [SLContentStream.SetHandler]
//   - [SLDataTimelineConfig.ConfigWithNameAndUpdateBlock]
//   - [SLDataTimelineConfig.CreateCancellableMachRecvSourceWithQueueCancelActionError]
//   - [SLDataTimelineConfig.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler]
//   - [SLDataTimelineConfig.EstablishConnectionWithResultBlock]
//   - [SLDataTimelineConfig.InitWithNameAndUpdateBlock]
//   - [SLDataTimelineServerSnapshotEntry.SessionsApplyBlock]
//   - [SLDataTimelineSessionSnapshotEntry.ProcessesApplyBlock]
//   - [SLDisplayPresetDeviceManager.StartWithBlockOnQueue]
//   - [SLSBrightnessControlClient.InitBrightnessControlClientNotifyQueueNotificationBlock]
//   - [SLSBrightnessControlClient.InitBrightnessControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSBrightnessControlClient.SetNotifyBlock]
//   - [SLSDisplayControlClient.RegisterClientPortNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClient.RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClient.RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClient.SetNotification]
//   - [SLSDisplayController.RegisterForNotificationsWithBlock]
//   - [SLSDisplayManager.RegisterPowerStateNotificationRegistrationIDSendInitialStateQueueRefconNotificationOptionNotificationBlockNotificationPayloadBlock]
//   - [SLSDisplayPowerControlClient.InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClient.InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSEventAuthenticationMessage.ValidateWithOptionsAndResultBlock]
//   - [SLSFullScreenPidReporter.ReportFullScreenStatusWithFilterAndHandler]
//   - [SLSFullScreenPidReporter.SetDisconnectHandler]
//   - [SLSGUIClient.ConfigGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClient.InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSRemoteViewEventClient.ActivateWithHandlerInvalidationHandler]
//   - [SLSRemoteViewEventClient.SendEventToHostFullDispatchReply]
//   - [SLSRemoteViewEventClient.ServicePassEventUpstreamToHostFullDispatchReply]
//   - [SLSScreenshot.CreateScreenshotPropertiesQueueHandlerError]
//   - [SLSScreenshot.SetHandler]
//   - [SLSSpaceWindowManager._performBatchingCallouts]
//   - [SLSWindowManagementBridgeDelegate.PerformWindowManagementBridgeTransactionUsingBlock]
//   - [SLSWindowManagementFallbackBridge.PerformWindowManagementBridgeTransactionUsingBlock]
//   - [SLSXPCService.CreateCancellableMachRecvSourceWithQueueErrorCancelAction]
//   - [SLSXPCService.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler]
//   - [SLSXPCService.InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock]
//   - [SLSXPCService.InitWithConnectionErrorHandlerNotificationBlock]
//   - [SLSXPCService.SetClientErrorBlock]
//   - [SLSXPCService.SetClientNotificationBlock]
//   - [SLSXPCService.SetErrorBlock]
//   - [SLSXPCService.SetNotificationBlock]
//   - [SLScreenTelemetryConnection.ConnectionWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock]
//   - [SLScreenTelemetryConnection.InitWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock]
//   - [SLSharingSessionManager.SetDelegateBlock]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CPXEventDeferringPolicy.Build]
//   - [CPXRemoteViewEventManager.PassEventUpstreamToHostFullDispatchReply]
//   - [CPXRemoteViewEventManager.SendEventToHostPidFullDispatchReply]
//   - [CPXSymbolicHotKeyRegistrar.RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc]
//   - [CPXSymbolicHotKeyRegistry.RegisterSymbolicHotKeyConnectionHotKeyIDSymbolicHotKeyOptionCallbackFunc]
//   - [ECTestOnlyEventAuthenticationMessage.ValidateWithOptionsAndResultBlock]
//   - [PKGCoreUITransaction.UpdateLayerKeyRendererWork]
//   - [PKGCoreUITransaction._scheduleRendererWorkMainThreadWork]
//   - [PKGCoreUIWork.SetMainThreadWork]
//   - [PKGCoreUIWork.SetRendererWork]
//   - [SLContentStream.CreateScreenshotPropertiesQueueHandlerError]
//   - [SLContentStream.InitWithFilterPropertiesQueueHandlerError]
//   - [SLContentStream.InitWithFilterPropertiesQueueHandler]
//   - [SLContentStream.SetHandler]
//   - [SLDataTimelineConfig.ConfigWithNameAndUpdateBlock]
//   - [SLDataTimelineConfig.CreateCancellableMachRecvSourceWithQueueCancelActionError]
//   - [SLDataTimelineConfig.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler]
//   - [SLDataTimelineConfig.EstablishConnectionWithResultBlock]
//   - [SLDataTimelineConfig.InitWithNameAndUpdateBlock]
//   - [SLDataTimelineServerSnapshotEntry.SessionsApplyBlock]
//   - [SLDataTimelineSessionSnapshotEntry.ProcessesApplyBlock]
//   - [SLDisplayPresetDeviceManager.StartWithBlockOnQueue]
//   - [SLSBrightnessControlClient.InitBrightnessControlClientNotifyQueueNotificationBlock]
//   - [SLSBrightnessControlClient.InitBrightnessControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSBrightnessControlClient.SetNotifyBlock]
//   - [SLSDisplayControlClient.RegisterClientPortNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClient.RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClient.RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClient.SetNotification]
//   - [SLSDisplayController.RegisterForNotificationsWithBlock]
//   - [SLSDisplayManager.RegisterPowerStateNotificationRegistrationIDSendInitialStateQueueRefconNotificationOptionNotificationBlockNotificationPayloadBlock]
//   - [SLSDisplayPowerControlClient.InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClient.InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSEventAuthenticationMessage.ValidateWithOptionsAndResultBlock]
//   - [SLSFullScreenPidReporter.ReportFullScreenStatusWithFilterAndHandler]
//   - [SLSFullScreenPidReporter.SetDisconnectHandler]
//   - [SLSGUIClient.ConfigGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClient.InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSRemoteViewEventClient.ActivateWithHandlerInvalidationHandler]
//   - [SLSRemoteViewEventClient.SendEventToHostFullDispatchReply]
//   - [SLSRemoteViewEventClient.ServicePassEventUpstreamToHostFullDispatchReply]
//   - [SLSScreenshot.CreateScreenshotPropertiesQueueHandlerError]
//   - [SLSScreenshot.SetHandler]
//   - [SLSSpaceWindowManager._performBatchingCallouts]
//   - [SLSWindowManagementBridgeDelegate.PerformWindowManagementBridgeTransactionUsingBlock]
//   - [SLSWindowManagementFallbackBridge.PerformWindowManagementBridgeTransactionUsingBlock]
//   - [SLSXPCService.CreateCancellableMachRecvSourceWithQueueErrorCancelAction]
//   - [SLSXPCService.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler]
//   - [SLSXPCService.InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock]
//   - [SLSXPCService.InitWithConnectionErrorHandlerNotificationBlock]
//   - [SLSXPCService.SetClientErrorBlock]
//   - [SLSXPCService.SetClientNotificationBlock]
//   - [SLSXPCService.SetErrorBlock]
//   - [SLSXPCService.SetNotificationBlock]
//   - [SLScreenTelemetryConnection.ConnectionWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock]
//   - [SLScreenTelemetryConnection.InitWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock]
//   - [SLSharingSessionManager.SetDelegateBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
