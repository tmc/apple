// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureTimecodeGenerator] class.
var (
	_AVCaptureTimecodeGeneratorClass     AVCaptureTimecodeGeneratorClass
	_AVCaptureTimecodeGeneratorClassOnce sync.Once
)

func getAVCaptureTimecodeGeneratorClass() AVCaptureTimecodeGeneratorClass {
	_AVCaptureTimecodeGeneratorClassOnce.Do(func() {
		_AVCaptureTimecodeGeneratorClass = AVCaptureTimecodeGeneratorClass{class: objc.GetClass("AVCaptureTimecodeGenerator")}
	})
	return _AVCaptureTimecodeGeneratorClass
}

// GetAVCaptureTimecodeGeneratorClass returns the class object for AVCaptureTimecodeGenerator.
func GetAVCaptureTimecodeGeneratorClass() AVCaptureTimecodeGeneratorClass {
	return getAVCaptureTimecodeGeneratorClass()
}

type AVCaptureTimecodeGeneratorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureTimecodeGeneratorClass) Alloc() AVCaptureTimecodeGenerator {
	rv := objc.Send[AVCaptureTimecodeGenerator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Generates and synchronizes timecode data from various sources for precise
// video and audio synchronization.
//
// # Overview
// 
// The [AVCaptureTimecodeGenerator] class supports multiple timecode sources,
// including frame counting, system clock synchronization, and MIDI timecode
// input (MTC). Suitable for playback, recording, or other time-sensitive
// operations where precise timecode metadata is required.
// 
// Use the [AVCaptureTimecodeGenerator.StartSynchronizationWithTimecodeSource] method to set up the
// desired timecode source.
//
// # Generating timecode
//
//   - [AVCaptureTimecodeGenerator.GenerateInitialTimecode]: Generates an initial timecode intended to be the first in a sequence.
//
// # Managing sources
//
//   - [AVCaptureTimecodeGenerator.CurrentSource]: The active timecode source used by [AVCaptureTimecodeGenerator](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureTimecodeGenerator>) to maintain clock synchronization for accurate timecode generation.
//   - [AVCaptureTimecodeGenerator.AvailableSources]: An array of available timecode synchronization sources that can be used by the timecode generator.
//   - [AVCaptureTimecodeGenerator.StartSynchronizationWithTimecodeSource]: Synchronizes the generator with the specified timecode source.
//
// # Configuring the generator
//
//   - [AVCaptureTimecodeGenerator.SynchronizationTimeout]: The maximum time interval allowed for source synchronization attempts before timing out.
//   - [AVCaptureTimecodeGenerator.SetSynchronizationTimeout]
//   - [AVCaptureTimecodeGenerator.TimecodeAlignmentOffset]: The time offset, in seconds, applied to the generated timecode.
//   - [AVCaptureTimecodeGenerator.SetTimecodeAlignmentOffset]
//   - [AVCaptureTimecodeGenerator.TimecodeFrameDuration]: The frame duration that the generator will use to generate timecodes.
//   - [AVCaptureTimecodeGenerator.SetTimecodeFrameDuration]
//   - [AVCaptureTimecodeGenerator.SetDelegateQueue]: Assigns a delegate to receive real-time timecode updates and specifies a queue for callbacks.
//
// # Handling delegate callbacks
//
//   - [AVCaptureTimecodeGenerator.Delegate]: The delegate that receives timecode updates from the timecode generator.
//   - [AVCaptureTimecodeGenerator.DelegateCallbackQueue]: The dispatch queue on which delegate callbacks are invoked.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator
type AVCaptureTimecodeGenerator struct {
	objectivec.Object
}

// AVCaptureTimecodeGeneratorFromID constructs a [AVCaptureTimecodeGenerator] from an objc.ID.
//
// Generates and synchronizes timecode data from various sources for precise
// video and audio synchronization.
func AVCaptureTimecodeGeneratorFromID(id objc.ID) AVCaptureTimecodeGenerator {
	return AVCaptureTimecodeGenerator{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureTimecodeGenerator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureTimecodeGenerator] class.
//
// # Generating timecode
//
//   - [IAVCaptureTimecodeGenerator.GenerateInitialTimecode]: Generates an initial timecode intended to be the first in a sequence.
//
// # Managing sources
//
//   - [IAVCaptureTimecodeGenerator.CurrentSource]: The active timecode source used by [AVCaptureTimecodeGenerator](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureTimecodeGenerator>) to maintain clock synchronization for accurate timecode generation.
//   - [IAVCaptureTimecodeGenerator.AvailableSources]: An array of available timecode synchronization sources that can be used by the timecode generator.
//   - [IAVCaptureTimecodeGenerator.StartSynchronizationWithTimecodeSource]: Synchronizes the generator with the specified timecode source.
//
// # Configuring the generator
//
//   - [IAVCaptureTimecodeGenerator.SynchronizationTimeout]: The maximum time interval allowed for source synchronization attempts before timing out.
//   - [IAVCaptureTimecodeGenerator.SetSynchronizationTimeout]
//   - [IAVCaptureTimecodeGenerator.TimecodeAlignmentOffset]: The time offset, in seconds, applied to the generated timecode.
//   - [IAVCaptureTimecodeGenerator.SetTimecodeAlignmentOffset]
//   - [IAVCaptureTimecodeGenerator.TimecodeFrameDuration]: The frame duration that the generator will use to generate timecodes.
//   - [IAVCaptureTimecodeGenerator.SetTimecodeFrameDuration]
//   - [IAVCaptureTimecodeGenerator.SetDelegateQueue]: Assigns a delegate to receive real-time timecode updates and specifies a queue for callbacks.
//
// # Handling delegate callbacks
//
//   - [IAVCaptureTimecodeGenerator.Delegate]: The delegate that receives timecode updates from the timecode generator.
//   - [IAVCaptureTimecodeGenerator.DelegateCallbackQueue]: The dispatch queue on which delegate callbacks are invoked.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator
type IAVCaptureTimecodeGenerator interface {
	objectivec.IObject

	// Topic: Generating timecode

	// Generates an initial timecode intended to be the first in a sequence.
	GenerateInitialTimecode() AVCaptureTimecode

	// Topic: Managing sources

	// The active timecode source used by [AVCaptureTimecodeGenerator](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureTimecodeGenerator>) to maintain clock synchronization for accurate timecode generation.
	CurrentSource() IAVCaptureTimecodeSource
	// An array of available timecode synchronization sources that can be used by the timecode generator.
	AvailableSources() []AVCaptureTimecodeSource
	// Synchronizes the generator with the specified timecode source.
	StartSynchronizationWithTimecodeSource(source IAVCaptureTimecodeSource)

	// Topic: Configuring the generator

	// The maximum time interval allowed for source synchronization attempts before timing out.
	SynchronizationTimeout() float64
	SetSynchronizationTimeout(value float64)
	// The time offset, in seconds, applied to the generated timecode.
	TimecodeAlignmentOffset() float64
	SetTimecodeAlignmentOffset(value float64)
	// The frame duration that the generator will use to generate timecodes.
	TimecodeFrameDuration() objectivec.IObject
	SetTimecodeFrameDuration(value objectivec.IObject)
	// Assigns a delegate to receive real-time timecode updates and specifies a queue for callbacks.
	SetDelegateQueue(delegate AVCaptureTimecodeGeneratorDelegate, callbackQueue dispatch.Queue)

	// Topic: Handling delegate callbacks

	// The delegate that receives timecode updates from the timecode generator.
	Delegate() AVCaptureTimecodeGeneratorDelegate
	// The dispatch queue on which delegate callbacks are invoked.
	DelegateCallbackQueue() dispatch.Queue
}

// Init initializes the instance.
func (c AVCaptureTimecodeGenerator) Init() AVCaptureTimecodeGenerator {
	rv := objc.Send[AVCaptureTimecodeGenerator](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureTimecodeGenerator) Autorelease() AVCaptureTimecodeGenerator {
	rv := objc.Send[AVCaptureTimecodeGenerator](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureTimecodeGenerator creates a new AVCaptureTimecodeGenerator instance.
func NewAVCaptureTimecodeGenerator() AVCaptureTimecodeGenerator {
	class := getAVCaptureTimecodeGeneratorClass()
	rv := objc.Send[AVCaptureTimecodeGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Generates an initial timecode intended to be the first in a sequence.
//
// # Return Value
// 
// A populated [AVCaptureTimecode] structure.
//
// [AVCaptureTimecode]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/generateInitialTimecode()
func (c AVCaptureTimecodeGenerator) GenerateInitialTimecode() AVCaptureTimecode {
	rv := objc.Send[AVCaptureTimecode](c.ID, objc.Sel("generateInitialTimecode"))
	return AVCaptureTimecode(rv)
}
// Synchronizes the generator with the specified timecode source.
//
// source: The timecode source for synchronization.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/startSynchronization(source:)
func (c AVCaptureTimecodeGenerator) StartSynchronizationWithTimecodeSource(source IAVCaptureTimecodeSource) {
	objc.Send[objc.ID](c.ID, objc.Sel("startSynchronizationWithTimecodeSource:"), source)
}
// Assigns a delegate to receive real-time timecode updates and specifies a
// queue for callbacks.
//
// delegate: An object conforming to the [AVCaptureTimecodeGeneratorDelegate] protocol.
//
// callbackQueue: The dispatch queue on which the delegate methods are invoked. The
// `callbackQueue` parameter may not be `nil`, except when setting the
// [AVCaptureTimecodeGeneratorDelegate] to `nil`, otherwise [SetDelegateQueue]
// throws an [NSInvalidArgumentException].
//
// # Discussion
// 
// Use this method to configure a delegate that handles timecode updates. The
// specified `queue` ensures thread-safe invocation of delegate methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/setDelegate(_:queue:)
func (c AVCaptureTimecodeGenerator) SetDelegateQueue(delegate AVCaptureTimecodeGeneratorDelegate, callbackQueue dispatch.Queue) {
	objc.Send[objc.ID](c.ID, objc.Sel("setDelegate:queue:"), delegate, uintptr(callbackQueue.Handle()))
}

// The active timecode source used by [AVCaptureTimecodeGenerator] to maintain
// clock synchronization for accurate timecode generation.
//
// # Discussion
// 
// Indicates the active timecode source, as defined in the
// [AVCaptureTimecodeSynchronizationSourceType] enum. If an
// [AVCaptureTimecodeGenerator] becomes disconnected from its source, it
// continues generating timecodes using historical data from its ring buffer.
// This approach allows the generator to maintain synchronization during brief
// disruptions, as is common in cinema workflows where timecode signals may
// experience discontinuities.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/currentSource
func (c AVCaptureTimecodeGenerator) CurrentSource() IAVCaptureTimecodeSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("currentSource"))
	return AVCaptureTimecodeSourceFromID(objc.ID(rv))
}
// An array of available timecode synchronization sources that can be used by
// the timecode generator.
//
// # Return Value
// 
// A read-only array of [AVCaptureTimecodeSource] objects representing the
// available timecode synchronization sources.
// 
// # Discussion
// 
// This property provides a list of [AVCaptureTimecodeSource] objects
// representing the available timecode sources with which the generator can
// synchronize. The sources may include built-in options such as the frame
// counter and real-time clock, as well as dynamically detected sources such
// as connected MIDI or HID devices.
// 
// This array is key-value observable, allowing you to monitor changes in
// real-time. For example, when a new MIDI device is connected, the array is
// updated to include the corresponding timecode source.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/availableSources
func (c AVCaptureTimecodeGenerator) AvailableSources() []AVCaptureTimecodeSource {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableSources"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureTimecodeSource {
		return AVCaptureTimecodeSourceFromID(id)
	})
}
// The maximum time interval allowed for source synchronization attempts
// before timing out.
//
// # Discussion
// 
// This property specifies the duration, in seconds, that the
// [AVCaptureTimecodeGenerator] will attempt to synchronize with a timecode
// source before timing out if synchronization cannot be achieved. If this
// threshold is exceeded, the synchronization status updates to reflect a
// timeout, and your
// [TimecodeGeneratorTransitionedToSynchronizationStatusForSource] delegate
// method fires, informing you of the event. The default value is 15 seconds.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/synchronizationTimeout
func (c AVCaptureTimecodeGenerator) SynchronizationTimeout() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("synchronizationTimeout"))
	return rv
}
func (c AVCaptureTimecodeGenerator) SetSynchronizationTimeout(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setSynchronizationTimeout:"), value)
}
// The time offset, in seconds, applied to the generated timecode.
//
// # Discussion
// 
// This offset allows fine-tuning of time alignment for synchronization with
// external sources or to accommodate any intentional delay. The default value
// is 0 seconds.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/timecodeAlignmentOffset
func (c AVCaptureTimecodeGenerator) TimecodeAlignmentOffset() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("timecodeAlignmentOffset"))
	return rv
}
func (c AVCaptureTimecodeGenerator) SetTimecodeAlignmentOffset(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setTimecodeAlignmentOffset:"), value)
}
// The frame duration that the generator will use to generate timecodes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/timecodeFrameDuration
func (c AVCaptureTimecodeGenerator) TimecodeFrameDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("timecodeFrameDuration"))
	return objectivec.Object{ID: rv}
}
func (c AVCaptureTimecodeGenerator) SetTimecodeFrameDuration(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setTimecodeFrameDuration:"), value)
}
// The delegate that receives timecode updates from the timecode generator.
//
// # Discussion
// 
// You can use your [Delegate] to receive real-time timecode updates.
// Implement the `` method in your delegate to handle updates.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/delegate
func (c AVCaptureTimecodeGenerator) Delegate() AVCaptureTimecodeGeneratorDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return AVCaptureTimecodeGeneratorDelegateObjectFromID(rv)
}
// The dispatch queue on which delegate callbacks are invoked.
//
// # Discussion
// 
// Provides the queue set in [SetDelegateQueue]. If no delegate is assigned,
// this property is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/delegateCallbackQueue
func (c AVCaptureTimecodeGenerator) DelegateCallbackQueue() dispatch.Queue {
	rv := objc.Send[uintptr](c.ID, objc.Sel("delegateCallbackQueue"))
	return dispatch.QueueFromHandle(rv)
}

// A frame counter timecode source that operates independently of any internal
// or external synchronization.
//
// # Discussion
// 
// This class property represents a standalone timecode source that advances
// based purely on frame count, independent of any real-time or external
// synchronization. It is ideal for scenarios where a simple, self-contained
// timing reference is sufficient, without requiring alignment to system
// clocks or external devices.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/frameCountSource
func (_AVCaptureTimecodeGeneratorClass AVCaptureTimecodeGeneratorClass) FrameCountSource() AVCaptureTimecodeSource {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureTimecodeGeneratorClass.class), objc.Sel("frameCountSource"))
	return AVCaptureTimecodeSourceFromID(objc.ID(rv))
}
// A predefined timecode source synchronized to the real-time system clock.
//
// # Discussion
// 
// This class property provides a default timecode source based on the
// real-time system clock, requiring no external device. It is ideal for live
// events or scenarios where alignment with the current time of day is
// necessary.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/realTimeClockSource
func (_AVCaptureTimecodeGeneratorClass AVCaptureTimecodeGeneratorClass) RealTimeClockSource() AVCaptureTimecodeSource {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureTimecodeGeneratorClass.class), objc.Sel("realTimeClockSource"))
	return AVCaptureTimecodeSourceFromID(objc.ID(rv))
}

