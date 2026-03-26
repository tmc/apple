// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A protocol for receiving real-time timecode updates and error notifications from a timecode generator.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGeneratorDelegate
type AVCaptureTimecodeGeneratorDelegate interface {
	objectivec.IObject

	// Notifies the delegate when new, unaligned timecodes are parsed from the specified source.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGeneratorDelegate/timecodeGenerator(_:didReceiveUpdate:from:)
	TimecodeGeneratorDidReceiveUpdateFromSource(generator IAVCaptureTimecodeGenerator, timecode AVCaptureTimecode, source IAVCaptureTimecodeSource)

	// Notifies the delegate when the list of available timecode synchronization sources is updated.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGeneratorDelegate/timecodeGenerator(_:didUpdateAvailableSources:)
	TimecodeGeneratorDidUpdateAvailableSources(generator IAVCaptureTimecodeGenerator, availableSources []AVCaptureTimecodeSource)

	// Notifies the delegate when the synchronization status of a timecode source changes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGeneratorDelegate/timecodeGenerator(_:transitionedTo:for:)
	TimecodeGeneratorTransitionedToSynchronizationStatusForSource(generator IAVCaptureTimecodeGenerator, synchronizationStatus AVCaptureTimecodeGeneratorSynchronizationStatus, source IAVCaptureTimecodeSource)
}

// AVCaptureTimecodeGeneratorDelegateObject wraps an existing Objective-C object that conforms to the AVCaptureTimecodeGeneratorDelegate protocol.
type AVCaptureTimecodeGeneratorDelegateObject struct {
	objectivec.Object
}
func (o AVCaptureTimecodeGeneratorDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVCaptureTimecodeGeneratorDelegateObjectFromID constructs a [AVCaptureTimecodeGeneratorDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVCaptureTimecodeGeneratorDelegateObjectFromID(id objc.ID) AVCaptureTimecodeGeneratorDelegateObject {
	return AVCaptureTimecodeGeneratorDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Notifies the delegate when new, unaligned timecodes are parsed from the
// specified source.
//
// generator: The timecode generator providing the update.
//
// timecode: The updated timecode data.
//
// source: The source from which the timecode was received.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGeneratorDelegate/timecodeGenerator(_:didReceiveUpdate:from:)
func (o AVCaptureTimecodeGeneratorDelegateObject) TimecodeGeneratorDidReceiveUpdateFromSource(generator IAVCaptureTimecodeGenerator, timecode AVCaptureTimecode, source IAVCaptureTimecodeSource) {
	objc.Send[struct{}](o.ID, objc.Sel("timecodeGenerator:didReceiveUpdate:fromSource:"), generator, timecode, source)
	}
// Notifies the delegate when the list of available timecode synchronization
// sources is updated.
//
// generator: The [AVCaptureTimecodeGenerator] instance providing the source list update.
//
// availableSources: An array of [AVCaptureTimecodeSource] objects representing the available
// timecode synchronization sources.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGeneratorDelegate/timecodeGenerator(_:didUpdateAvailableSources:)
func (o AVCaptureTimecodeGeneratorDelegateObject) TimecodeGeneratorDidUpdateAvailableSources(generator IAVCaptureTimecodeGenerator, availableSources []AVCaptureTimecodeSource) {
	objc.Send[struct{}](o.ID, objc.Sel("timecodeGenerator:didUpdateAvailableSources:"), generator, objectivec.IObjectSliceToNSArray(availableSources))
	}
// Notifies the delegate when the synchronization status of a timecode source
// changes.
//
// generator: The [AVCaptureTimecodeGenerator] instance providing the status update.
//
// synchronizationStatus: The updated synchronization state.
//
// source: The internal or external source to which the generator synchronizes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGeneratorDelegate/timecodeGenerator(_:transitionedTo:for:)
func (o AVCaptureTimecodeGeneratorDelegateObject) TimecodeGeneratorTransitionedToSynchronizationStatusForSource(generator IAVCaptureTimecodeGenerator, synchronizationStatus AVCaptureTimecodeGeneratorSynchronizationStatus, source IAVCaptureTimecodeSource) {
	objc.Send[struct{}](o.ID, objc.Sel("timecodeGenerator:transitionedToSynchronizationStatus:forSource:"), generator, synchronizationStatus, source)
	}

// AVCaptureTimecodeGeneratorDelegateConfig holds optional typed callbacks for [AVCaptureTimecodeGeneratorDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegeneratordelegate
type AVCaptureTimecodeGeneratorDelegateConfig struct {

	// Other Methods
	// TimecodeGeneratorDidReceiveUpdateFromSource — Notifies the delegate when new, unaligned timecodes are parsed from the specified source.
	TimecodeGeneratorDidReceiveUpdateFromSource func(generator AVCaptureTimecodeGenerator, timecode AVCaptureTimecode, source AVCaptureTimecodeSource)
	// TimecodeGeneratorTransitionedToSynchronizationStatusForSource — Notifies the delegate when the synchronization status of a timecode source changes.
	TimecodeGeneratorTransitionedToSynchronizationStatusForSource func(generator AVCaptureTimecodeGenerator, synchronizationStatus AVCaptureTimecodeGeneratorSynchronizationStatus, source AVCaptureTimecodeSource)
}

// NewAVCaptureTimecodeGeneratorDelegate creates an Objective-C object implementing the [AVCaptureTimecodeGeneratorDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVCaptureTimecodeGeneratorDelegateObject] satisfies the [AVCaptureTimecodeGeneratorDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegeneratordelegate
func NewAVCaptureTimecodeGeneratorDelegate(config AVCaptureTimecodeGeneratorDelegateConfig) AVCaptureTimecodeGeneratorDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVCaptureTimecodeGeneratorDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TimecodeGeneratorDidReceiveUpdateFromSource != nil {
		fn := config.TimecodeGeneratorDidReceiveUpdateFromSource
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("timecodeGenerator:didReceiveUpdate:fromSource:"),
			Fn: func(self objc.ID, _cmd objc.SEL, generatorID objc.ID, timecode AVCaptureTimecode, sourceID objc.ID) {
				generator := AVCaptureTimecodeGeneratorFromID(generatorID)
				source := AVCaptureTimecodeSourceFromID(sourceID)
				fn(generator, timecode, source)
			},
		})
	}

	if config.TimecodeGeneratorTransitionedToSynchronizationStatusForSource != nil {
		fn := config.TimecodeGeneratorTransitionedToSynchronizationStatusForSource
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("timecodeGenerator:transitionedToSynchronizationStatus:forSource:"),
			Fn: func(self objc.ID, _cmd objc.SEL, generatorID objc.ID, synchronizationStatus AVCaptureTimecodeGeneratorSynchronizationStatus, sourceID objc.ID) {
				generator := AVCaptureTimecodeGeneratorFromID(generatorID)
				source := AVCaptureTimecodeSourceFromID(sourceID)
				fn(generator, synchronizationStatus, source)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVCaptureTimecodeGeneratorDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVCaptureTimecodeGeneratorDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVCaptureTimecodeGeneratorDelegateObjectFromID(instance)
}

