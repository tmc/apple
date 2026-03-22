// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// SCRecordingOutputDelegate protocol.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputDelegate
type SCRecordingOutputDelegate interface {
	objectivec.IObject
}

// SCRecordingOutputDelegateObject wraps an existing Objective-C object that conforms to the SCRecordingOutputDelegate protocol.
type SCRecordingOutputDelegateObject struct {
	objectivec.Object
}
func (o SCRecordingOutputDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// SCRecordingOutputDelegateObjectFromID constructs a [SCRecordingOutputDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SCRecordingOutputDelegateObjectFromID(id objc.ID) SCRecordingOutputDelegateObject {
	return SCRecordingOutputDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputDelegate/recordingOutput(_:didFailWithError:)
func (o SCRecordingOutputDelegateObject) RecordingOutputDidFailWithError(recordingOutput ISCRecordingOutput, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("recordingOutput:didFailWithError:"), recordingOutput, error_)
	}
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputDelegate/recordingOutputDidFinishRecording(_:)
func (o SCRecordingOutputDelegateObject) RecordingOutputDidFinishRecording(recordingOutput ISCRecordingOutput) {
	objc.Send[struct{}](o.ID, objc.Sel("recordingOutputDidFinishRecording:"), recordingOutput)
	}
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCRecordingOutputDelegate/recordingOutputDidStartRecording(_:)
func (o SCRecordingOutputDelegateObject) RecordingOutputDidStartRecording(recordingOutput ISCRecordingOutput) {
	objc.Send[struct{}](o.ID, objc.Sel("recordingOutputDidStartRecording:"), recordingOutput)
	}

// SCRecordingOutputDelegateConfig holds optional typed callbacks for [SCRecordingOutputDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/screencapturekit/screcordingoutputdelegate
type SCRecordingOutputDelegateConfig struct {

	// Instance Methods
	RecordingOutputDidFailWithError func(recordingOutput SCRecordingOutput, error_ foundation.NSError)
	RecordingOutputDidFinishRecording func(recordingOutput SCRecordingOutput)
	RecordingOutputDidStartRecording func(recordingOutput SCRecordingOutput)
}

// NewSCRecordingOutputDelegate creates an Objective-C object implementing the [SCRecordingOutputDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [SCRecordingOutputDelegateObject] satisfies the [SCRecordingOutputDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/screencapturekit/screcordingoutputdelegate
func NewSCRecordingOutputDelegate(config SCRecordingOutputDelegateConfig) SCRecordingOutputDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoSCRecordingOutputDelegate_%d", n)

	var methods []objc.MethodDef

	if config.RecordingOutputDidFailWithError != nil {
		fn := config.RecordingOutputDidFailWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("recordingOutput:didFailWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, recordingOutputID objc.ID, error_ID objc.ID) {
				recordingOutput := SCRecordingOutputFromID(recordingOutputID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(recordingOutput, error_)
			},
		})
	}

	if config.RecordingOutputDidFinishRecording != nil {
		fn := config.RecordingOutputDidFinishRecording
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("recordingOutputDidFinishRecording:"),
			Fn: func(self objc.ID, _cmd objc.SEL, recordingOutputID objc.ID) {
				recordingOutput := SCRecordingOutputFromID(recordingOutputID)
				fn(recordingOutput)
			},
		})
	}

	if config.RecordingOutputDidStartRecording != nil {
		fn := config.RecordingOutputDidStartRecording
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("recordingOutputDidStartRecording:"),
			Fn: func(self objc.ID, _cmd objc.SEL, recordingOutputID objc.ID) {
				recordingOutput := SCRecordingOutputFromID(recordingOutputID)
				fn(recordingOutput)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("SCRecordingOutputDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewSCRecordingOutputDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return SCRecordingOutputDelegateObjectFromID(instance)
}

