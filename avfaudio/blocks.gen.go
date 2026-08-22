// Code generated from Apple documentation. DO NOT EDIT.

package avfaudio

import (
	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal AVAudioApplicationMicrophoneInjectionPermission) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioBufferHandler handles The system calls this closure with the generated audio buffer.
//
// Used by:
//   - [AVSpeechSynthesizer.WriteUtteranceToBufferCallbackToMarkerCallback]
//   - [AVSpeechSynthesizer.WriteUtteranceToBufferCallback]
type AVAudioBufferHandler = func(*AVAudioBuffer)

// NewAVAudioBufferBlock wraps a Go [AVAudioBufferHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVSpeechSynthesizer.WriteUtteranceToBufferCallbackToMarkerCallback]
//   - [AVSpeechSynthesizer.WriteUtteranceToBufferCallback]
func NewAVAudioBufferBlock(handler AVAudioBufferHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *AVAudioBuffer
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := AVAudioBufferFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioBufferUint32Handler handles A block the framework calls to get input data.
//
// Used by:
//   - [AVAudioConverter.ConvertToBufferErrorWithInputFromBlock]
type AVAudioBufferUint32Handler = func(uint32, *AVAudioConverterInputStatus) AVAudioBuffer

// NewAVAudioBufferUint32Block wraps a Go [AVAudioBufferUint32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioConverter.ConvertToBufferErrorWithInputFromBlock]
func NewAVAudioBufferUint32Block(handler AVAudioBufferUint32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint32, extra0 *AVAudioConverterInputStatus) objc.ID {
		return handler(primitive, extra0).ID
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioEngineManualRenderingStatusUint32Handler handles completion with primitive and object results.
type AVAudioEngineManualRenderingStatusUint32Handler = func(uint32, *coreaudiotypes.AudioBufferList, *int32) AVAudioEngineManualRenderingStatus

// NewAVAudioEngineManualRenderingStatusUint32Block wraps a Go [AVAudioEngineManualRenderingStatusUint32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAVAudioEngineManualRenderingStatusUint32Block(handler AVAudioEngineManualRenderingStatusUint32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint32, extra0 *coreaudiotypes.AudioBufferList, extra1 *int32) AVAudioEngineManualRenderingStatus {
		return handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioIONodeInputBlock handles The type that represents a block to render operation calls to get input data when in manual rendering mode.

// NewAVAudioIONodeInputBlock wraps a Go [AVAudioIONodeInputBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAVAudioIONodeInputBlock(handler AVAudioIONodeInputBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal uint32) *coreaudiotypes.AudioBufferList {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioNodeCompletionHandler handles A general callback handler for an audio node.

// NewAVAudioNodeCompletionHandlerBlock wraps a Go [AVAudioNodeCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAVAudioNodeCompletionHandlerBlock(handler AVAudioNodeCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioPCMBufferAVAudioTimeHandler handles A block the framework calls with audio buffers.
//
// Used by:
//   - [AVAudioNode.InstallTapOnBusBufferSizeFormatBlock]
type AVAudioPCMBufferAVAudioTimeHandler = func(*AVAudioPCMBuffer, *AVAudioTime)

// NewAVAudioPCMBufferAVAudioTimeBlock wraps a Go [AVAudioPCMBufferAVAudioTimeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioNode.InstallTapOnBusBufferSizeFormatBlock]
func NewAVAudioPCMBufferAVAudioTimeBlock(handler AVAudioPCMBufferAVAudioTimeHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *AVAudioPCMBuffer
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := AVAudioPCMBufferFromID(resultID)
			result = &v
		}
		var extra0 *AVAudioTime
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := AVAudioTimeFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioPlayerNodeCompletionCallbackTypeHandler handles The handler the system calls after the player schedules the file for playback on the render thread, or the player stops.
//
// Used by:
//   - [AVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleBufferCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleFileAtTimeCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionCallbackTypeCompletionHandler]
type AVAudioPlayerNodeCompletionCallbackTypeHandler = func(int)

// NewAVAudioPlayerNodeCompletionCallbackTypeBlock wraps a Go [AVAudioPlayerNodeCompletionCallbackTypeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleBufferCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleFileAtTimeCompletionCallbackTypeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionCallbackTypeCompletionHandler]
func NewAVAudioPlayerNodeCompletionCallbackTypeBlock(handler AVAudioPlayerNodeCompletionCallbackTypeHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal int) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioPlayerNodeCompletionHandler handles The callback handler for buffer or file completion.

// NewAVAudioPlayerNodeCompletionHandlerBlock wraps a Go [AVAudioPlayerNodeCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAVAudioPlayerNodeCompletionHandlerBlock(handler AVAudioPlayerNodeCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal AVAudioPlayerNodeCompletionCallbackType) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAudioSequencerUserCallback handles A callback the sequencer calls asynchronously during playback when it encounters a user event.

// NewAVAudioSequencerUserCallbackBlock wraps a Go [AVAudioSequencerUserCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAVAudioSequencerUserCallbackBlock(handler AVAudioSequencerUserCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive AVMusicTrack, extra0 foundation.NSData, extra1 float64) {
		handler(primitive, extra0, extra1)
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
	if handler == nil {
		return 0, func() {}
	}
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal AVAudioVoiceProcessingSpeechActivityEvent) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVMIDIPlayerCompletionHandler handles A callback the system invokes when MIDI playback completes.

// NewAVMIDIPlayerCompletionHandlerBlock wraps a Go [AVMIDIPlayerCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAVMIDIPlayerCompletionHandlerBlock(handler AVMIDIPlayerCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// AVMusicEventEnumerationBlock handles A type you use to enumerate and remove music events, if necessary.

// AVMusicEventFloat64Int8Handler handles The block to call for each event.
//
// Used by:
//   - [AVMusicTrack.EnumerateEventsInRangeUsingBlock]
type AVMusicEventFloat64Int8Handler = func(*AVMusicEvent, []float64, *int8)

// AVMusicTrackDataFloat64Handler handles The user callback that the system calls.
//
// Used by:
//   - [AVAudioSequencer.SetUserCallback]
type AVMusicTrackDataFloat64Handler = func(*AVMusicTrack, *foundation.NSData, float64)

// NewAVMusicTrackDataFloat64Block wraps a Go [AVMusicTrackDataFloat64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioSequencer.SetUserCallback]
func NewAVMusicTrackDataFloat64Block(handler AVMusicTrackDataFloat64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, extra1 float64) {
		var result *AVMusicTrack
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := AVMusicTrackFromID(resultID)
			result = &v
		}
		var extra0 *foundation.NSData
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := foundation.NSDataFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVSpeechSynthesisMarkerArrayAVSpeechSynthesisProviderRequestHandler is the signature for a completion handler block.
type AVSpeechSynthesisMarkerArrayAVSpeechSynthesisProviderRequestHandler = func(*[]AVSpeechSynthesisMarker, *AVSpeechSynthesisProviderRequest)

// NewAVSpeechSynthesisMarkerArrayAVSpeechSynthesisProviderRequestBlock wraps a Go [AVSpeechSynthesisMarkerArrayAVSpeechSynthesisProviderRequestHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAVSpeechSynthesisMarkerArrayAVSpeechSynthesisProviderRequestBlock(handler AVSpeechSynthesisMarkerArrayAVSpeechSynthesisProviderRequestHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *[]AVSpeechSynthesisMarker
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]AVSpeechSynthesisMarker, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = AVSpeechSynthesisMarkerFromID(item.GetID())
			}
			result = &res
		}
		var extra0 *AVSpeechSynthesisProviderRequest
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := AVSpeechSynthesisProviderRequestFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVSpeechSynthesisMarkerArrayHandler handles A callback that the system invokes with marker information.
//
// Used by:
//   - [AVSpeechSynthesizer.WriteUtteranceToBufferCallbackToMarkerCallback]
type AVSpeechSynthesisMarkerArrayHandler = func(*[]AVSpeechSynthesisMarker)

// NewAVSpeechSynthesisMarkerArrayBlock wraps a Go [AVSpeechSynthesisMarkerArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVSpeechSynthesizer.WriteUtteranceToBufferCallbackToMarkerCallback]
func NewAVSpeechSynthesisMarkerArrayBlock(handler AVSpeechSynthesisMarkerArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]AVSpeechSynthesisMarker
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]AVSpeechSynthesisMarker, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = AVSpeechSynthesisMarkerFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal AVSpeechSynthesisPersonalVoiceAuthorizationStatus) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVSpeechSynthesisProviderOutputBlock handles A type that represents the method for sending marker information to the host.

// AVSpeechSynthesizerBufferCallback handles A type that defines a callback that receives a buffer of generated speech.

// NewAVSpeechSynthesizerBufferCallbackBlock wraps a Go [AVSpeechSynthesizerBufferCallback] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAVSpeechSynthesizerBufferCallbackBlock(handler AVSpeechSynthesizerBufferCallback) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal AVAudioBuffer) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVSpeechSynthesizerMarkerCallback handles A type that defines a callback that receives speech markers.

// AudioBufferListUint32Handler handles The block the engine calls on the input node to get the audio to send to the output when operating in the manual rendering mode.
//
// Used by:
//   - [AVAudioInputNode.SetManualRenderingInputPCMFormatInputBlock]
type AudioBufferListUint32Handler = func(uint32) *coreaudiotypes.AudioBufferList

// NewAudioBufferListUint32Block wraps a Go [AudioBufferListUint32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioInputNode.SetManualRenderingInputPCMFormatInputBlock]
func NewAudioBufferListUint32Block(handler AudioBufferListUint32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal uint32) *coreaudiotypes.AudioBufferList {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolAVAudioUnitComponentInt8Handler handles The block to apply to the audio unit components.
//   - comp: A block to test.
//   - stop: A reference to a Boolean value. To stop further processing of the search, the block sets the value to [true](<https://developer.apple.com/documentation/Swift/true>). The stop argument is an out-only argument. Only set this Boolean to [true](<https://developer.apple.com/documentation/Swift/true>) within the block.
//
// Used by:
//   - [AVAudioUnitComponentManager.ComponentsPassingTest]
type BoolAVAudioUnitComponentInt8Handler = func(*AVAudioUnitComponent, *int8) bool

// NewBoolAVAudioUnitComponentInt8Block wraps a Go [BoolAVAudioUnitComponentInt8Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioUnitComponentManager.ComponentsPassingTest]
func NewBoolAVAudioUnitComponentInt8Block(handler BoolAVAudioUnitComponentInt8Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 *int8) bool {
		var result *AVAudioUnitComponent
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := AVAudioUnitComponentFromID(resultID)
			result = &v
		}
		return handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolBoolHandler handles A callback that the system invokes when the input mute state changes.
//
// Used by:
//   - [AVAudioApplication.SetInputMuteStateChangeHandlerError]
type BoolBoolHandler = func(bool) bool

// NewBoolBoolBlock wraps a Go [BoolBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioApplication.SetInputMuteStateChangeHandlerError]
func NewBoolBoolBlock(handler BoolBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolErrorHandler handles A completion handler the system calls asynchronously when the system completes audio routing arbitration.
//   - defaultDeviceChanged: A Boolean value that indicates whether the system switched the AirPods to the macOS device.
//   - error: An error object that indicates why the request failed, or [nil](<https://developer.apple.com/documentation/ObjectiveC/nil-227m0>) if the request succeeded.
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolHandler handles A Boolean value that indicates whether the user grants the app permission to record audio.
//
// Used by:
//   - [AVAudioApplication.RequestRecordPermissionWithCompletionHandler]
type BoolHandler = func(bool)

// NewBoolBlock wraps a Go [BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioApplication.RequestRecordPermissionWithCompletionHandler]
func NewBoolBlock(handler BoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles If not [NULL], the source node’s event list block calls this on the real-time thread.
//
// Used by:
//   - [AVAudioEngine.ConnectMIDIToFormatEventListBlock]
//   - [AVAudioEngine.ConnectMIDIToNodesFormatEventListBlock]
type ErrorHandler = func()

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioEngine.ConnectMIDIToFormatEventListBlock]
//   - [AVAudioEngine.ConnectMIDIToNodesFormatEventListBlock]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// IntAudioTimeStampHandler handles The block that receives audio data from the input.
//
// Used by:
//   - [AVAudioSinkNode.InitWithReceiverBlock]
type IntAudioTimeStampHandler = func(*coreaudiotypes.AudioTimeStamp, uint32, *coreaudiotypes.AudioBufferList) int

// NewIntAudioTimeStampBlock wraps a Go [IntAudioTimeStampHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioSinkNode.InitWithReceiverBlock]
func NewIntAudioTimeStampBlock(handler IntAudioTimeStampHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *coreaudiotypes.AudioTimeStamp, extra0 uint32, extra1 *coreaudiotypes.AudioBufferList) int {
		return handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// IntInt8Handler handles The block to supply audio data to the output.
//
// Used by:
//   - [AVAudioSourceNode.InitWithFormatRenderBlock]
//   - [AVAudioSourceNode.InitWithRenderBlock]
type IntInt8Handler = func(*int8, *coreaudiotypes.AudioTimeStamp, uint32, *coreaudiotypes.AudioBufferList) int

// NewIntInt8Block wraps a Go [IntInt8Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioSourceNode.InitWithFormatRenderBlock]
//   - [AVAudioSourceNode.InitWithRenderBlock]
func NewIntInt8Block(handler IntInt8Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *int8, extra0 *coreaudiotypes.AudioTimeStamp, extra1 uint32, extra2 *coreaudiotypes.AudioBufferList) int {
		return handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles The handler the system calls after the player schedules the file for playback on the render thread, or the player stops.
//
// Used by:
//   - [AVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleBufferCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleFileAtTimeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler]
//   - [AVMIDIPlayer.Play]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleBufferCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleFileAtTimeCompletionHandler]
//   - [AVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler]
//   - [AVMIDIPlayer.Play]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// constAudioBufferListHandler handles The closure the method invokes when the resulting PCM buffer object deallocates.
//
// Used by:
//   - [AVAudioPCMBuffer.InitWithPCMFormatBufferListNoCopyDeallocator]
type constAudioBufferListHandler = func(*coreaudiotypes.AudioBufferList)

// NewconstAudioBufferListBlock wraps a Go [constAudioBufferListHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAudioPCMBuffer.InitWithPCMFormatBufferListNoCopyDeallocator]
func NewconstAudioBufferListBlock(handler constAudioBufferListHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *coreaudiotypes.AudioBufferList) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}
