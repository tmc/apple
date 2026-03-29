// Code generated from Apple documentation. DO NOT EDIT.

package avfaudio

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [AVAudioApplication.RequestRecordPermissionWithCompletionHandler]
//   - [AVAudioRoutingArbiter.BeginArbitrationWithAudioSessionCategoryModeOptionsCompletionHandler]
//   - [AVAudioRoutingArbiter.BeginArbitrationWithBTSessionCategoryModeFlagsCompletionHandler]
//   - [AVAudioUnitComponent.ValidateWithResultsInCompletionHandler]
//   - [AVSpeechSynthesisVoice.SpeechVoicesIncludingSuperCompactWithCompletionHandler]
//   - [AVSpeechSynthesisVoice.SpeechVoicesWithCompletionHandler]
//   - [AVSpeechSynthesisVoice._speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler]
//   - [AVSpeechSynthesisVoice._speechVoicesIncludingSiriWithCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioApplication.RequestRecordPermissionWithCompletionHandler]
//   - [AVAudioRoutingArbiter.BeginArbitrationWithAudioSessionCategoryModeOptionsCompletionHandler]
//   - [AVAudioRoutingArbiter.BeginArbitrationWithBTSessionCategoryModeFlagsCompletionHandler]
//   - [AVAudioUnitComponent.ValidateWithResultsInCompletionHandler]
//   - [AVSpeechSynthesisVoice.SpeechVoicesIncludingSuperCompactWithCompletionHandler]
//   - [AVSpeechSynthesisVoice.SpeechVoicesWithCompletionHandler]
//   - [AVSpeechSynthesisVoice._speechVoicesIncludingSiriAndSuperCompactWithCompletionHandler]
//   - [AVSpeechSynthesisVoice._speechVoicesIncludingSiriWithCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [AVAudioApplication.PrivateCallInputMuteHandlerBlockInputMutedIsTopDownMuteContext]
//   - [AVAudioApplication.PrivateSetInputMuteStateChangeHandlerError]
//   - [AVAudioEngine.ConnectMIDIToFormatBlock]
//   - [AVAudioEngine.ConnectMIDIToNodesFormatBlock]
//   - [AVAudioPlayerNode.CallLegacyCompletionHandlerForTypeLegacyHandler]
//   - [AVAudioSinkNode.PullInputBlockFromReceiverBlock]
//   - [AVAudioSourceNode.PullInputBlockFromRenderBlock]
//   - [AVSpeechUtterance.SetAudioBufferCallback]
//   - [AVSpeechUtterance.SetMarkerCallback]
//   - [AVVCAudioDeviceManager.IterateOverProcessObjectList]
//   - [AVVCMetricsManager.MeasureElapseTimeForMetricBlock]
//   - [AVVCSessionFactory.SessionForContextClientTypeCompletion]
//   - [AVVCSessionFactory.SessionForContextCompletion]
//   - [AVVCSessionFactory.SessionManagerForContextClientTypeCompletion]
//   - [AVVCSessionFactory.SetSessionWasCreatedBlock]
//   - [AVVCSessionFactory.SetSessionWillBeDestroyedBlock]
//   - [AVVoiceController.ConfigureAlertBehaviorForStreamCompletion]
//   - [AVVoiceController.DeactivateAudioSessionForStreamWithOptionsCompletion]
//   - [AVVoiceController.EnableTriangleModeForStreamEnableWithCompletion]
//   - [AVVoiceController.GetDeviceLatenciesForStreamWithCompletion]
//   - [AVVoiceController.GetInputChannelInfoForStreamCompletion]
//   - [AVVoiceController.GetPlaybackRouteForStreamWithCompletion]
//   - [AVVoiceController.PlayAlertWithOverrideCompletion]
//   - [AVVoiceController.PrepareRecordForStreamCompletion]
//   - [AVVoiceController.RemoveStreamCompletion]
//   - [AVVoiceController.SetContextCompletion]
//   - [AVVoiceController.SetRecordStatusChangeBlock]
//   - [AVVoiceController.StartKeepAliveQueueForStreamCompletion]
//   - [AVVoiceController.StartRecordForStreamCompletion]
//   - [AVVoiceController.StartRecordWithSettingsCompletionAlertCompletionAudioCallback]
//   - [AVVoiceController.StopKeepAliveQueueForStreamCompletion]
//   - [AVVoiceController.StopRecordForStreamCompletion]
//   - [AVVoiceTriggerClient.EnableBargeInModeCompletionBlock]
//   - [AVVoiceTriggerClient.EnableListeningOnPortsCompletionBlock]
//   - [AVVoiceTriggerClient.EnableSpeakerStateListeningCompletionBlock]
//   - [AVVoiceTriggerClient.EnableVoiceTriggerListeningCompletionBlock]
//   - [AVVoiceTriggerClient.GetInputChannelInfoCompletion]
//   - [AVVoiceTriggerClient.ListeningEnabledCompletionBlock]
//   - [AVVoiceTriggerClient.PortStateActiveCompletionBlock]
//   - [AVVoiceTriggerClient.SetAggressiveECModeCompletionBlock]
//   - [AVVoiceTriggerClient.SetAvvcServerCrashedBlock]
//   - [AVVoiceTriggerClient.SetAvvcServerResetBlock]
//   - [AVVoiceTriggerClient.SetListeningPropertyCompletionBlock]
//   - [AVVoiceTriggerClient.SetPortStateChangedBlock]
//   - [AVVoiceTriggerClient.SetServerCrashedBlock]
//   - [AVVoiceTriggerClient.SetServerResetBlock]
//   - [AVVoiceTriggerClient.SetSiriClientRecordStateChangedBlock]
//   - [AVVoiceTriggerClient.SetSpeakerMuteStateChangedBlock]
//   - [AVVoiceTriggerClient.SetSpeakerStateChangedBlock]
//   - [AVVoiceTriggerClient.SetVoiceTriggerBlock]
//   - [AVVoiceTriggerClient.SiriClientsRecordingCompletionBlock]
//   - [AVVoiceTriggerClient.SpeakerStateActiveCompletionBlock]
//   - [AVVoiceTriggerClient.SpeakerStateMutedCompletionBlock]
//   - [AVVoiceTriggerClient.UpdateVoiceTriggerConfigurationCompletionBlock]
//   - [AVVoiceTriggerClient.VoiceTriggerPastDataFramesAvailableCompletion]
//   - [AVVoiceTriggerClientPortManager.InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationBlock]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioApplication.PrivateCallInputMuteHandlerBlockInputMutedIsTopDownMuteContext]
//   - [AVAudioApplication.PrivateSetInputMuteStateChangeHandlerError]
//   - [AVAudioEngine.ConnectMIDIToFormatBlock]
//   - [AVAudioEngine.ConnectMIDIToNodesFormatBlock]
//   - [AVAudioPlayerNode.CallLegacyCompletionHandlerForTypeLegacyHandler]
//   - [AVAudioSinkNode.PullInputBlockFromReceiverBlock]
//   - [AVAudioSourceNode.PullInputBlockFromRenderBlock]
//   - [AVSpeechUtterance.SetAudioBufferCallback]
//   - [AVSpeechUtterance.SetMarkerCallback]
//   - [AVVCAudioDeviceManager.IterateOverProcessObjectList]
//   - [AVVCMetricsManager.MeasureElapseTimeForMetricBlock]
//   - [AVVCSessionFactory.SessionForContextClientTypeCompletion]
//   - [AVVCSessionFactory.SessionForContextCompletion]
//   - [AVVCSessionFactory.SessionManagerForContextClientTypeCompletion]
//   - [AVVCSessionFactory.SetSessionWasCreatedBlock]
//   - [AVVCSessionFactory.SetSessionWillBeDestroyedBlock]
//   - [AVVoiceController.ConfigureAlertBehaviorForStreamCompletion]
//   - [AVVoiceController.DeactivateAudioSessionForStreamWithOptionsCompletion]
//   - [AVVoiceController.EnableTriangleModeForStreamEnableWithCompletion]
//   - [AVVoiceController.GetDeviceLatenciesForStreamWithCompletion]
//   - [AVVoiceController.GetInputChannelInfoForStreamCompletion]
//   - [AVVoiceController.GetPlaybackRouteForStreamWithCompletion]
//   - [AVVoiceController.PlayAlertWithOverrideCompletion]
//   - [AVVoiceController.PrepareRecordForStreamCompletion]
//   - [AVVoiceController.RemoveStreamCompletion]
//   - [AVVoiceController.SetContextCompletion]
//   - [AVVoiceController.SetRecordStatusChangeBlock]
//   - [AVVoiceController.StartKeepAliveQueueForStreamCompletion]
//   - [AVVoiceController.StartRecordForStreamCompletion]
//   - [AVVoiceController.StartRecordWithSettingsCompletionAlertCompletionAudioCallback]
//   - [AVVoiceController.StopKeepAliveQueueForStreamCompletion]
//   - [AVVoiceController.StopRecordForStreamCompletion]
//   - [AVVoiceTriggerClient.EnableBargeInModeCompletionBlock]
//   - [AVVoiceTriggerClient.EnableListeningOnPortsCompletionBlock]
//   - [AVVoiceTriggerClient.EnableSpeakerStateListeningCompletionBlock]
//   - [AVVoiceTriggerClient.EnableVoiceTriggerListeningCompletionBlock]
//   - [AVVoiceTriggerClient.GetInputChannelInfoCompletion]
//   - [AVVoiceTriggerClient.ListeningEnabledCompletionBlock]
//   - [AVVoiceTriggerClient.PortStateActiveCompletionBlock]
//   - [AVVoiceTriggerClient.SetAggressiveECModeCompletionBlock]
//   - [AVVoiceTriggerClient.SetAvvcServerCrashedBlock]
//   - [AVVoiceTriggerClient.SetAvvcServerResetBlock]
//   - [AVVoiceTriggerClient.SetListeningPropertyCompletionBlock]
//   - [AVVoiceTriggerClient.SetPortStateChangedBlock]
//   - [AVVoiceTriggerClient.SetServerCrashedBlock]
//   - [AVVoiceTriggerClient.SetServerResetBlock]
//   - [AVVoiceTriggerClient.SetSiriClientRecordStateChangedBlock]
//   - [AVVoiceTriggerClient.SetSpeakerMuteStateChangedBlock]
//   - [AVVoiceTriggerClient.SetSpeakerStateChangedBlock]
//   - [AVVoiceTriggerClient.SetVoiceTriggerBlock]
//   - [AVVoiceTriggerClient.SiriClientsRecordingCompletionBlock]
//   - [AVVoiceTriggerClient.SpeakerStateActiveCompletionBlock]
//   - [AVVoiceTriggerClient.SpeakerStateMutedCompletionBlock]
//   - [AVVoiceTriggerClient.UpdateVoiceTriggerConfigurationCompletionBlock]
//   - [AVVoiceTriggerClient.VoiceTriggerPastDataFramesAvailableCompletion]
//   - [AVVoiceTriggerClientPortManager.InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

