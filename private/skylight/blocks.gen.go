// Code generated from Apple documentation. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// DictionaryHandler is the signature for a completion handler block.
//
// Used by:
//   - [SLSBrightnessControl.RegisterForNotificationsWithBlock]
type DictionaryHandler = func(*foundation.INSDictionary)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [CPXRemoteViewEventPendingConnection.InitWithConnectionHandler]
//   - [CPXRemoteViewEventPendingConnection.SetHandler]
type ErrorHandler = func()

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CPXRemoteViewEventPendingConnection.InitWithConnectionHandler]
//   - [CPXRemoteViewEventPendingConnection.SetHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// NumberErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [CPXRemoteViewEventProtocolServerCallsClient.SendEventToHostFullDispatchReply]
type NumberErrorHandler = func(*foundation.NSNumber, error)

// NewNumberErrorBlock wraps a Go [NumberErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CPXRemoteViewEventProtocolServerCallsClient.SendEventToHostFullDispatchReply]
func NewNumberErrorBlock(handler NumberErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSNumber
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSNumberFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// SLDataTimelineProcessHandler is the signature for a completion handler block.
//
// Used by:
//   - [SLDataTimelineSessionProcessCollection.ProcessesApplyBlock]
type SLDataTimelineProcessHandler = func(*objectivec.Object)

// NewSLDataTimelineProcessBlock wraps a Go [SLDataTimelineProcessHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SLDataTimelineSessionProcessCollection.ProcessesApplyBlock]
func NewSLDataTimelineProcessBlock(handler SLDataTimelineProcessHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *objectivec.Object
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := objectivec.ObjectFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// SLDataTimelineServerSnapshotHandler is the signature for a completion handler block.
//
// Used by:
//   - [SLDataTimelineSnapshotCollection.SnapshotsApplyBlock]
type SLDataTimelineServerSnapshotHandler = func(*objectivec.Object)

// NewSLDataTimelineServerSnapshotBlock wraps a Go [SLDataTimelineServerSnapshotHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SLDataTimelineSnapshotCollection.SnapshotsApplyBlock]
func NewSLDataTimelineServerSnapshotBlock(handler SLDataTimelineServerSnapshotHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *objectivec.Object
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := objectivec.ObjectFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// SLDataTimelineSessionHandler is the signature for a completion handler block.
//
// Used by:
//   - [SLDataTimelineServerSnapshot.SessionsApplyBlock]
type SLDataTimelineSessionHandler = func(*objectivec.Object)

// NewSLDataTimelineSessionBlock wraps a Go [SLDataTimelineSessionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [SLDataTimelineServerSnapshot.SessionsApplyBlock]
func NewSLDataTimelineSessionBlock(handler SLDataTimelineSessionHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *objectivec.Object
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := objectivec.ObjectFromID(resultID)
			result = &v
		}
		handler(result)
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
//   - [SLSDisplayControlClientProtocol.RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClientProtocol.RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayController.RegisterForNotificationsWithBlock]
//   - [SLSDisplayManager.RegisterPowerStateNotificationRegistrationIDSendInitialStateQueueRefconNotificationOptionNotificationBlockNotificationPayloadBlock]
//   - [SLSDisplayPowerControlClient.InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClient.InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClientProtocol.InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClientProtocol.InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSEventAuthenticationMessage.ValidateWithOptionsAndResultBlock]
//   - [SLSFullScreenPidReporter.ReportFullScreenStatusWithFilterAndHandler]
//   - [SLSFullScreenPidReporter.SetDisconnectHandler]
//   - [SLSGUIClient.ConfigGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClient.InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClientProtocol.ConfigGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClientProtocol.InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
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
//   - [SLSXPCServiceProtocol.InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock]
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
//   - [SLSDisplayControlClientProtocol.RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClientProtocol.RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayController.RegisterForNotificationsWithBlock]
//   - [SLSDisplayManager.RegisterPowerStateNotificationRegistrationIDSendInitialStateQueueRefconNotificationOptionNotificationBlockNotificationPayloadBlock]
//   - [SLSDisplayPowerControlClient.InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClient.InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClientProtocol.InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClientProtocol.InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSEventAuthenticationMessage.ValidateWithOptionsAndResultBlock]
//   - [SLSFullScreenPidReporter.ReportFullScreenStatusWithFilterAndHandler]
//   - [SLSFullScreenPidReporter.SetDisconnectHandler]
//   - [SLSGUIClient.ConfigGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClient.InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClientProtocol.ConfigGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSGUIClientProtocol.InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
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
//   - [SLSXPCServiceProtocol.InitConnectionWithNameNotificationQueueWithAutoreconnectErrorHandlerNotificationBlock]
//   - [SLScreenTelemetryConnection.ConnectionWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock]
//   - [SLScreenTelemetryConnection.InitWithZoneWidthZoneHeightZoneRowsZoneColumnsSamplingIntervalQueueAndUpdateBlock]
//   - [SLSharingSessionManager.SetDelegateBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
