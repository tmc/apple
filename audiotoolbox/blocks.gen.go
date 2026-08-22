// Code generated from Apple documentation. DO NOT EDIT.

package audiotoolbox

import (
	"unsafe"

	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/coremidi"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// AUEventListenerBlock handles completion with primitive and object results.

// NewAUEventListenerBlock wraps a Go [AUEventListenerBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUEventListenerBlock(handler AUEventListenerBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 *AudioUnitEvent, extra1 uint64, extra2 float32) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUImplementorValueObserver handles A block called to notify the audio unit implementation of changes to a parameter value.

// NewAUImplementorValueObserverBlock wraps a Go [AUImplementorValueObserver] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUImplementorValueObserverBlock(handler AUImplementorValueObserver) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive AUParameter, extra0 float32) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUImplementorValueProvider handles A block called to fetch a parameter’s current value from the audio unit implementation.

// NewAUImplementorValueProviderBlock wraps a Go [AUImplementorValueProvider] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUImplementorValueProviderBlock(handler AUImplementorValueProvider) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal AUParameter) float32 {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUInputHandler handles A block to notify the host of an I/O unit that an input is available.

// NewAUInputHandlerBlock wraps a Go [AUInputHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUInputHandlerBlock(handler AUInputHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *AudioUnitRenderActionFlags, extra0 *coreaudiotypes.AudioTimeStamp, extra1 uint32, extra2 int) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUMIDICIProfileChangedBlock handles completion with primitive and object results.

// NewAUMIDICIProfileChangedBlock wraps a Go [AUMIDICIProfileChangedBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUMIDICIProfileChangedBlock(handler AUMIDICIProfileChangedBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint32, extra0 uint32, extra1 *coremidi.MIDICIProfile, extra2 bool) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUParameterAutomationObserver handles completion with primitive and object results.

// NewAUParameterAutomationObserverBlock wraps a Go [AUParameterAutomationObserver] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUParameterAutomationObserverBlock(handler AUParameterAutomationObserver) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int, extra0 *AUParameterAutomationEvent) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUParameterFloat32Handler is the signature for a completion handler block.
type AUParameterFloat32Handler = func(*AUParameter, float32)

// NewAUParameterFloat32Block wraps a Go [AUParameterFloat32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUParameterFloat32Block(handler AUParameterFloat32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 float32) {
		var result *AUParameter
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := AUParameterFromID(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUParameterListenerBlock handles completion with primitive and object results.

// NewAUParameterListenerBlock wraps a Go [AUParameterListenerBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUParameterListenerBlock(handler AUParameterListenerBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 *AudioUnitParameter, extra1 float32) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUParameterObserver handles A block called after the value of a parameter changes.

// NewAUParameterObserverBlock wraps a Go [AUParameterObserver] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUParameterObserverBlock(handler AUParameterObserver) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint64, extra0 float32) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUParameterRecordingObserver handles A block called to record parameter changes as automation events.

// NewAUParameterRecordingObserverBlock wraps a Go [AUParameterRecordingObserver] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUParameterRecordingObserverBlock(handler AUParameterRecordingObserver) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int, extra0 *AURecordedParameterEvent) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// AURenderContextObserver handles A custom block that tells the audio unit which thread context to use for the next render cycle.

// NewAURenderContextObserverBlock wraps a Go [AURenderContextObserver] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAURenderContextObserverBlock(handler AURenderContextObserver) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *AudioUnitRenderContext) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AURenderObserver handles A block called when an audio unit renders audio.

// NewAURenderObserverBlock wraps a Go [AURenderObserver] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAURenderObserverBlock(handler AURenderObserver) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint, extra0 *coreaudiotypes.AudioTimeStamp, extra1 uint32, extra2 int) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUScheduleMIDIEventBlock handles A block to schedule MIDI events.

// NewAUScheduleMIDIEventBlock wraps a Go [AUScheduleMIDIEventBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUScheduleMIDIEventBlock(handler AUScheduleMIDIEventBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int, extra0 uint32, extra1 int, extra2 *uint8) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUScheduleParameterBlock handles A block to schedule parameter changes.

// NewAUScheduleParameterBlock wraps a Go [AUScheduleParameterBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUScheduleParameterBlock(handler AUScheduleParameterBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int, extra0 uint32, extra1 uint64, extra2 float32) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// AUViewControllerBaseHandler is the signature for a completion handler block.
//
// Used by:
//   - [AUAudioUnit.RequestViewControllerWithCompletionHandler]
type AUViewControllerBaseHandler = func(*uintptr)

// AUVoiceIOMutedSpeechActivityEventListener handles A block that the system calls to indicate speech activity while the user has the microphone muted.

// NewAUVoiceIOMutedSpeechActivityEventListenerBlock wraps a Go [AUVoiceIOMutedSpeechActivityEventListener] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAUVoiceIOMutedSpeechActivityEventListenerBlock(handler AUVoiceIOMutedSpeechActivityEventListener) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal AUVoiceIOSpeechActivityEvent) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AudioComponentInstanceErrorHandler handles The block called when instantiation has completed.
//   - audioUnit: An initialized audio unit if the operation succeeded, or `nil` if it failed.
//   - error: An error if the operation failed, or `nil` if it succeeded.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AUAudioUnit.InstantiateWithComponentDescriptionOptionsCompletionHandler]
type AudioComponentInstanceErrorHandler = func(AudioComponentInstance, error)

// NewAudioComponentInstanceErrorBlock wraps a Go [AudioComponentInstanceErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AUAudioUnit.InstantiateWithComponentDescriptionOptionsCompletionHandler]
func NewAudioComponentInstanceErrorBlock(handler AudioComponentInstanceErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result AudioComponentInstance = AudioComponentInstance(resultID)
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// AudioQueueInputCallbackBlock handles completion with primitive and object results.

// NewAudioQueueInputCallbackBlock wraps a Go [AudioQueueInputCallbackBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAudioQueueInputCallbackBlock(handler AudioQueueInputCallbackBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *uintptr, extra0 *AudioQueueBuffer, extra1 *coreaudiotypes.AudioTimeStamp, extra2 uint32, extra3 *coreaudiotypes.AudioStreamPacketDescription) {
		handler(primitive, extra0, extra1, extra2, extra3)
	})
	return objc.ID(block), func() { block.Release() }
}

// AudioQueueOutputCallbackBlock handles completion with primitive and object results.

// NewAudioQueueOutputCallbackBlock wraps a Go [AudioQueueOutputCallbackBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAudioQueueOutputCallbackBlock(handler AudioQueueOutputCallbackBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *uintptr, extra0 *AudioQueueBuffer) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// AudioUnitRemoteControlEventListener handles completion with a primitive value.

// NewAudioUnitRemoteControlEventListenerBlock wraps a Go [AudioUnitRemoteControlEventListener] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAudioUnitRemoteControlEventListenerBlock(handler AudioUnitRemoteControlEventListener) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal AudioUnitRemoteControlEvent) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// AudioUnitRenderActionFlagsAudioTimeStampUint32Int64Handler handles completion with primitive and object results.
type AudioUnitRenderActionFlagsAudioTimeStampUint32Int64Handler = func(*AudioUnitRenderActionFlags, *coreaudiotypes.AudioTimeStamp, uint32, int64)

// NewAudioUnitRenderActionFlagsAudioTimeStampUint32Int64Block wraps a Go [AudioUnitRenderActionFlagsAudioTimeStampUint32Int64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAudioUnitRenderActionFlagsAudioTimeStampUint32Int64Block(handler AudioUnitRenderActionFlagsAudioTimeStampUint32Int64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *AudioUnitRenderActionFlags, extra0 *coreaudiotypes.AudioTimeStamp, extra1 uint32, extra2 int64) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolAUHostTransportStateFlagsHandler is the signature for a completion handler block.
type BoolAUHostTransportStateFlagsHandler = func([]float64, []float64, []float64) bool

// BoolFloat64Handler is the signature for a completion handler block.
type BoolFloat64Handler = func([]float64, *int, []float64, *int, []float64) bool

// CallHostBlock handles a primitive value and returns a primitive value.

// NewCallHostBlock wraps a Go [CallHostBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCallHostBlock(handler CallHostBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID) foundation.INSDictionary {
		var primitiveVal foundation.INSDictionary
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitiveVal = foundation.NSDictionaryFromID(primitiveID)
		}
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// Float32AUParameterHandler is the signature for a completion handler block.
type Float32AUParameterHandler = func(*AUParameter) float32

// NewFloat32AUParameterBlock wraps a Go [Float32AUParameterHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewFloat32AUParameterBlock(handler Float32AUParameterHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) float32 {
		var result *AUParameter
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := AUParameterFromID(resultID)
			result = &v
		}
		return handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// Float32AUParameterStringHandler is the signature for a completion handler block.
type Float32AUParameterStringHandler = func(*AUParameter, string) float32

// NewFloat32AUParameterStringBlock wraps a Go [Float32AUParameterStringHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewFloat32AUParameterStringBlock(handler Float32AUParameterStringHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) float32 {
		var result *AUParameter
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := AUParameterFromID(resultID)
			result = &v
		}
		var extra0 string = objc.IDToString(extra0ID)
		return handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// INSDictionaryDictionaryHandler is the signature for a completion handler block.
type INSDictionaryDictionaryHandler = func(*foundation.INSDictionary) foundation.INSDictionary

// Int64AUParameterAutomationEventHandler handles completion with primitive and object results.
//
// Used by:
//   - [AUParameterNode.TokenByAddingParameterAutomationObserver]
type Int64AUParameterAutomationEventHandler = func(int64, *AUParameterAutomationEvent)

// NewInt64AUParameterAutomationEventBlock wraps a Go [Int64AUParameterAutomationEventHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AUParameterNode.TokenByAddingParameterAutomationObserver]
func NewInt64AUParameterAutomationEventBlock(handler Int64AUParameterAutomationEventHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int64, extra0 *AUParameterAutomationEvent) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// Int64AURecordedParameterEventHandler handles A block called to record parameter changes.
//
// Used by:
//   - [AUParameterNode.TokenByAddingParameterRecordingObserver]
type Int64AURecordedParameterEventHandler = func(int64, *AURecordedParameterEvent)

// NewInt64AURecordedParameterEventBlock wraps a Go [Int64AURecordedParameterEventHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AUParameterNode.TokenByAddingParameterRecordingObserver]
func NewInt64AURecordedParameterEventBlock(handler Int64AURecordedParameterEventHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int64, extra0 *AURecordedParameterEvent) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// Int64Uint32Uint64Float32Handler handles completion with primitive and object results.
type Int64Uint32Uint64Float32Handler = func(int64, uint32, uint64, float32)

// NewInt64Uint32Uint64Float32Block wraps a Go [Int64Uint32Uint64Float32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewInt64Uint32Uint64Float32Block(handler Int64Uint32Uint64Float32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int64, extra0 uint32, extra1 uint64, extra2 float32) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// IntAudioTimeStampUint32Int64Handler handles The block to call.
//
// Used by:
//   - [AUAudioUnit.TokenByAddingRenderObserver]
type IntAudioTimeStampUint32Int64Handler = func(int, *coreaudiotypes.AudioTimeStamp, uint32, int64)

// NewIntAudioTimeStampUint32Int64Block wraps a Go [IntAudioTimeStampUint32Int64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AUAudioUnit.TokenByAddingRenderObserver]
func NewIntAudioTimeStampUint32Int64Block(handler IntAudioTimeStampUint32Int64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int, extra0 *coreaudiotypes.AudioTimeStamp, extra1 uint32, extra2 int64) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// IntAudioUnitRenderActionFlagsHandler handles completion with primitive and object results.
type IntAudioUnitRenderActionFlagsHandler = func(*AudioUnitRenderActionFlags, *coreaudiotypes.AudioTimeStamp, uint32, int64, *coreaudiotypes.AudioBufferList, unsafe.Pointer) int

// NewIntAudioUnitRenderActionFlagsBlock wraps a Go [IntAudioUnitRenderActionFlagsHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIntAudioUnitRenderActionFlagsBlock(handler IntAudioUnitRenderActionFlagsHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *AudioUnitRenderActionFlags, extra0 *coreaudiotypes.AudioTimeStamp, extra1 uint32, extra2 int64, extra3 *coreaudiotypes.AudioBufferList, extra4 unsafe.Pointer) int {
		return handler(primitive, extra0, extra1, extra2, extra3, extra4)
	})
	return objc.ID(block), func() { block.Release() }
}

// IntInt64Handler handles a primitive value and returns a primitive value.
type IntInt64Handler = func(int64) int

// NewIntInt64Block wraps a Go [IntInt64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewIntInt64Block(handler IntInt64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal int64) int {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// StringAUParameterNodeInt64Handler is the signature for a completion handler block.
type StringAUParameterNodeInt64Handler = func(*AUParameterNode, int64) string

// NewStringAUParameterNodeInt64Block wraps a Go [StringAUParameterNodeInt64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewStringAUParameterNodeInt64Block(handler StringAUParameterNodeInt64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 int64) string {
		var result *AUParameterNode
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := AUParameterNodeFromID(resultID)
			result = &v
		}
		return handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// StringVoidHandler is the signature for a completion handler block.
type StringVoidHandler = func() string

// NewStringVoidBlock wraps a Go [StringVoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewStringVoidBlock(handler StringVoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) string {
		return handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// Uint32Uint32MIDICIProfileBoolHandler handles completion with primitive and object results.
type Uint32Uint32MIDICIProfileBoolHandler = func(uint32, uint32, *coremidi.MIDICIProfile, bool)

// NewUint32Uint32MIDICIProfileBoolBlock wraps a Go [Uint32Uint32MIDICIProfileBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewUint32Uint32MIDICIProfileBoolBlock(handler Uint32Uint32MIDICIProfileBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint32, extra0 uint32, extra1ID objc.ID, extra2 bool) {
		var extra1 *coremidi.MIDICIProfile
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			v := coremidi.MIDICIProfileFromID(extra1ID)
			extra1 = &v
		}
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// Uint64Float32Handler handles A block called after the value of a parameter has changed.
//
// Used by:
//   - [AUParameterNode.TokenByAddingParameterObserver]
type Uint64Float32Handler = func(uint64, float32)

// NewUint64Float32Block wraps a Go [Uint64Float32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AUParameterNode.TokenByAddingParameterObserver]
func NewUint64Float32Block(handler Uint64Float32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive uint64, extra0 float32) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// conststructAudioUnitRenderContextHandler handles completion with a primitive value.
type conststructAudioUnitRenderContextHandler = func(*AudioUnitRenderContext)

// NewconststructAudioUnitRenderContextBlock wraps a Go [conststructAudioUnitRenderContextHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewconststructAudioUnitRenderContextBlock(handler conststructAudioUnitRenderContextHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *AudioUnitRenderContext) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}
