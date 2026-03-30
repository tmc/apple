// Code generated from Apple documentation. DO NOT EDIT.

package avfaudio

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AVAudioApplicationMicrophoneInjectionPermissionHandler handles completion with a primitive value.
//
// Used by:
//   - [AVAudioApplication.RequestMicrophoneInjectionPermissionWithCompletionHandler]
type AVAudioApplicationMicrophoneInjectionPermissionHandler = func(AVAudioApplicationMicrophoneInjectionPermission)

// NewAVAudioApplicationMicrophoneInjectionPermissionBlock wraps a Go [AVAudioApplicationMicrophoneInjectionPermissionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioApplication.RequestMicrophoneInjectionPermissionWithCompletionHandler]
func NewAVAudioApplicationMicrophoneInjectionPermissionBlock(handler AVAudioApplicationMicrophoneInjectionPermissionHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal AVAudioApplicationMicrophoneInjectionPermission) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioUnitComponentHandler handles The block to apply to the audio unit components.
//   - comp: A block to test.
//   - stop: A reference to a Boolean value. To stop further processing of the search, the block sets the value to [true](<doc://com.apple.documentation/documentation/Swift/true>). The stop argument is an out-only argument. Only set this Boolean to [true](<doc://com.apple.documentation/documentation/Swift/true>) within the block.
//
// Used by:
//   - [AVAudioUnitComponentManager.ComponentsPassingTest]
type AVAudioUnitComponentHandler = func(*AVAudioUnitComponent)

// NewAVAudioUnitComponentBlock wraps a Go [AVAudioUnitComponentHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioUnitComponentManager.ComponentsPassingTest]
func NewAVAudioUnitComponentBlock(handler AVAudioUnitComponentHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *AVAudioUnitComponent
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := AVAudioUnitComponentFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioUnitErrorHandler handles A handler the framework calls in an arbitrary thread context when creation completes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVAudioUnit.InstantiateWithComponentDescriptionOptionsCompletionHandler]
type AVAudioUnitErrorHandler = func(*AVAudioUnit, error)

// NewAVAudioUnitErrorBlock wraps a Go [AVAudioUnitErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioUnit.InstantiateWithComponentDescriptionOptionsCompletionHandler]
func NewAVAudioUnitErrorBlock(handler AVAudioUnitErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVAudioUnit
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := AVAudioUnitFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioVoiceProcessingSpeechActivityEventHandler handles completion with a primitive value.
//
// Used by:
//   - [AVAudioInputNode.SetMutedSpeechActivityEventListener]
type AVAudioVoiceProcessingSpeechActivityEventHandler = func(AVAudioVoiceProcessingSpeechActivityEvent)

// NewAVAudioVoiceProcessingSpeechActivityEventBlock wraps a Go [AVAudioVoiceProcessingSpeechActivityEventHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioInputNode.SetMutedSpeechActivityEventListener]
func NewAVAudioVoiceProcessingSpeechActivityEventBlock(handler AVAudioVoiceProcessingSpeechActivityEventHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal AVAudioVoiceProcessingSpeechActivityEvent) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVSpeechSynthesisPersonalVoiceAuthorizationStatusHandler handles A completion handler that the system calls after the user responds to a request to authorize use of personal voices, which receives the authorization status as an argument.
//
// Used by:
//   - [AVSpeechSynthesizer.RequestPersonalVoiceAuthorizationWithCompletionHandler]
type AVSpeechSynthesisPersonalVoiceAuthorizationStatusHandler = func(AVSpeechSynthesisPersonalVoiceAuthorizationStatus)

// NewAVSpeechSynthesisPersonalVoiceAuthorizationStatusBlock wraps a Go [AVSpeechSynthesisPersonalVoiceAuthorizationStatusHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVSpeechSynthesizer.RequestPersonalVoiceAuthorizationWithCompletionHandler]
func NewAVSpeechSynthesisPersonalVoiceAuthorizationStatusBlock(handler AVSpeechSynthesisPersonalVoiceAuthorizationStatusHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal AVSpeechSynthesisPersonalVoiceAuthorizationStatus) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolErrorHandler handles A completion handler the system calls asynchronously when the system completes audio routing arbitration.
//   - defaultDeviceChanged: A Boolean value that indicates whether the system switched the AirPods to the macOS device.
//   - error: An error object that indicates why the request failed, or [nil](<doc://com.apple.documentation/documentation/ObjectiveC/nil-227m0>) if the request succeeded.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVAudioRoutingArbiter.BeginArbitrationWithCategoryCompletionHandler]
type BoolErrorHandler = func(bool, error)

// NewBoolErrorBlock wraps a Go [BoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioRoutingArbiter.BeginArbitrationWithCategoryCompletionHandler]
func NewBoolErrorBlock(handler BoolErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolHandler handles A Boolean value that indicates whether the user grants the app permission to record audio.
//
// Used by:
//   - [AVAudioApplication.RequestRecordPermissionWithCompletionHandler]
//   - [AVAudioApplication.SetInputMuteStateChangeHandlerError]
type BoolHandler = func(bool)

// NewBoolBlock wraps a Go [BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioApplication.RequestRecordPermissionWithCompletionHandler]
//   - [AVAudioApplication.SetInputMuteStateChangeHandlerError]
func NewBoolBlock(handler BoolHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles The handler the system calls after the player schedules the file for playback on the render thread, or the player stops.
//
// Used by:
//   - [AVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleBufferCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleBufferCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleFileAtTimeCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleFileAtTimeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler]
//   - [AVMIDIPlayer.Play]
type ErrorHandler = func()

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleBufferCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleBufferCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleFileAtTimeCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleFileAtTimeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler]
//   - [AVMIDIPlayer.Play]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// constAudioBufferListHandler handles The closure the method invokes when the resulting PCM buffer object deallocates.
//
// Used by:
//   - [AVAudioPCMBuffer.InitWithPCMFormatBufferListNoCopyDeallocator]
type constAudioBufferListHandler = func(*objectivec.Object)

// NewconstAudioBufferListBlock wraps a Go [constAudioBufferListHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioPCMBuffer.InitWithPCMFormatBufferListNoCopyDeallocator]
func NewconstAudioBufferListBlock(handler constAudioBufferListHandler) (objc.ID, func()) {
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
