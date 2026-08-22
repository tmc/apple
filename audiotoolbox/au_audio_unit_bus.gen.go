// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AUAudioUnitBus] class.
var (
	_AUAudioUnitBusClass     AUAudioUnitBusClass
	_AUAudioUnitBusClassOnce sync.Once
)

func getAUAudioUnitBusClass() AUAudioUnitBusClass {
	_AUAudioUnitBusClassOnce.Do(func() {
		_AUAudioUnitBusClass = AUAudioUnitBusClass{class: objc.GetClass("AUAudioUnitBus")}
	})
	return _AUAudioUnitBusClass
}

// GetAUAudioUnitBusClass returns the class object for AUAudioUnitBus.
func GetAUAudioUnitBusClass() AUAudioUnitBusClass {
	return getAUAudioUnitBusClass()
}

type AUAudioUnitBusClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AUAudioUnitBusClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AUAudioUnitBusClass) Alloc() AUAudioUnitBus {
	rv := objc.Send[AUAudioUnitBus](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A class that defines an input or output connection point on an audio unit.
//
// # Bus Methods and Properties
//
//   - [AUAudioUnitBus.SetFormatError]: Sets the bus’s audio format.
//   - [AUAudioUnitBus.Format]: The audio format and channel layout of audio being transferred on the bus.
//   - [AUAudioUnitBus.IsEnabled]: Determines whether the bus is active.
//   - [AUAudioUnitBus.SetEnabled]
//   - [AUAudioUnitBus.Name]: A name for the bus.
//   - [AUAudioUnitBus.SetName]
//   - [AUAudioUnitBus.Index]: The index of this bus in its containing array.
//   - [AUAudioUnitBus.BusType]: The bus type.
//   - [AUAudioUnitBus.OwnerAudioUnit]: The audio unit that owns the bus.
//   - [AUAudioUnitBus.SupportedChannelLayoutTags]: An array of audio channel layout tags.
//   - [AUAudioUnitBus.ContextPresentationLatency]: Information about latency in the audio unit’s processing context.
//   - [AUAudioUnitBus.SetContextPresentationLatency]
//   - [AUAudioUnitBus.ShouldAllocateBuffer]
//   - [AUAudioUnitBus.SetShouldAllocateBuffer]
//
// # Audio Unit Implementations
//
//   - [AUAudioUnitBus.InitWithFormatError]: Initializes a bus object with a specific format.
//   - [AUAudioUnitBus.SupportedChannelCounts]: An array of numbers indicating the supported number of channels for this bus.
//   - [AUAudioUnitBus.SetSupportedChannelCounts]
//   - [AUAudioUnitBus.MaximumChannelCount]: The maximum number of channels supported for this bus.
//   - [AUAudioUnitBus.SetMaximumChannelCount]
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus
type AUAudioUnitBus struct {
	objectivec.Object
}

// AUAudioUnitBusFromID constructs a [AUAudioUnitBus] from an objc.ID.
//
// A class that defines an input or output connection point on an audio unit.
func AUAudioUnitBusFromID(id objc.ID) AUAudioUnitBus {
	return AUAudioUnitBus{objectivec.Object{ID: id}}
}

// NOTE: AUAudioUnitBus adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AUAudioUnitBus] class.
//
// # Bus Methods and Properties
//
//   - [IAUAudioUnitBus.SetFormatError]: Sets the bus’s audio format.
//   - [IAUAudioUnitBus.Format]: The audio format and channel layout of audio being transferred on the bus.
//   - [IAUAudioUnitBus.IsEnabled]: Determines whether the bus is active.
//   - [IAUAudioUnitBus.SetEnabled]
//   - [IAUAudioUnitBus.Name]: A name for the bus.
//   - [IAUAudioUnitBus.SetName]
//   - [IAUAudioUnitBus.Index]: The index of this bus in its containing array.
//   - [IAUAudioUnitBus.BusType]: The bus type.
//   - [IAUAudioUnitBus.OwnerAudioUnit]: The audio unit that owns the bus.
//   - [IAUAudioUnitBus.SupportedChannelLayoutTags]: An array of audio channel layout tags.
//   - [IAUAudioUnitBus.ContextPresentationLatency]: Information about latency in the audio unit’s processing context.
//   - [IAUAudioUnitBus.SetContextPresentationLatency]
//   - [IAUAudioUnitBus.ShouldAllocateBuffer]
//   - [IAUAudioUnitBus.SetShouldAllocateBuffer]
//
// # Audio Unit Implementations
//
//   - [IAUAudioUnitBus.InitWithFormatError]: Initializes a bus object with a specific format.
//   - [IAUAudioUnitBus.SupportedChannelCounts]: An array of numbers indicating the supported number of channels for this bus.
//   - [IAUAudioUnitBus.SetSupportedChannelCounts]
//   - [IAUAudioUnitBus.MaximumChannelCount]: The maximum number of channels supported for this bus.
//   - [IAUAudioUnitBus.SetMaximumChannelCount]
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus
type IAUAudioUnitBus interface {
	objectivec.IObject

	// Topic: Bus Methods and Properties

	// Sets the bus’s audio format.
	SetFormatError(format objectivec.IObject) (bool, error)
	// The audio format and channel layout of audio being transferred on the bus.
	Format() objectivec.IObject
	// Determines whether the bus is active.
	IsEnabled() bool
	SetEnabled(value bool)
	// A name for the bus.
	Name() string
	SetName(value string)
	// The index of this bus in its containing array.
	Index() uint
	// The bus type.
	BusType() AUAudioUnitBusType
	// The audio unit that owns the bus.
	OwnerAudioUnit() IAUAudioUnit
	// An array of audio channel layout tags.
	SupportedChannelLayoutTags() []foundation.NSNumber
	// Information about latency in the audio unit’s processing context.
	ContextPresentationLatency() foundation.NSTimeInterval
	SetContextPresentationLatency(value foundation.NSTimeInterval)
	ShouldAllocateBuffer() bool
	SetShouldAllocateBuffer(value bool)

	// Topic: Audio Unit Implementations

	// Initializes a bus object with a specific format.
	InitWithFormatError(format objectivec.IObject) (AUAudioUnitBus, error)
	// An array of numbers indicating the supported number of channels for this bus.
	SupportedChannelCounts() []foundation.NSNumber
	SetSupportedChannelCounts(value []foundation.NSNumber)
	// The maximum number of channels supported for this bus.
	MaximumChannelCount() AUAudioChannelCount
	SetMaximumChannelCount(value AUAudioChannelCount)
}

// Init initializes the instance.
func (a AUAudioUnitBus) Init() AUAudioUnitBus {
	rv := objc.Send[AUAudioUnitBus](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AUAudioUnitBus) Autorelease() AUAudioUnitBus {
	rv := objc.Send[AUAudioUnitBus](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAUAudioUnitBus creates a new AUAudioUnitBus instance.
func NewAUAudioUnitBus() AUAudioUnitBus {
	class := getAUAudioUnitBusClass()
	rv := objc.Send[AUAudioUnitBus](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a bus object with a specific format.
//
// format: The initial audio format.
//
// # Return Value
//
// A newly-initialized bus object, or `nil` if the operation failed.
//
// # Discussion
//
// Audio units can generally be expected to support the [AVAudioFormat]
// standard format (deinterleaved 32-bit float), at any sample rate.
//
// Channel counts can be more complex. See the
// [AUAudioUnit.ChannelCapabilities] reference for a more complete discussion.
//
// Initialization will fail and return an error if the specified format is
// unsupported for the bus.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/init(format:)
//
// [AVAudioFormat]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat
func NewAudioUnitBusWithFormatError(format objectivec.IObject) (AUAudioUnitBus, error) {
	var errorPtr objc.ID
	instance := getAUAudioUnitBusClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:error:"), format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AUAudioUnitBus{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return AUAudioUnitBus{}, objc.ErrInitFailed
	}
	return AUAudioUnitBusFromID(rv), nil
}

// Sets the bus’s audio format.
//
// format: The desired audio format.
//
// format is a [*avfaudio.AVAudioFormat].
//
// # Discussion
//
// - false if the operation failed.
//
// # Discussion
//
// Audio units can generally be expected to support the [AVAudioFormat]
// standard format (deinterleaved 32-bit float), at any sample rate.
//
// Channel counts can be more complex. See the
// [AUAudioUnit.ChannelCapabilities] reference for a more complete discussion.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/setFormat(_:)
//
// [AVAudioFormat]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat
func (a AUAudioUnitBus) SetFormatError(format objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("setFormat:error:"), format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setFormat:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Initializes a bus object with a specific format.
//
// format: The initial audio format.
//
// format is a [*avfaudio.AVAudioFormat].
//
// # Return Value
//
// A newly-initialized bus object, or `nil` if the operation failed.
//
// # Discussion
//
// Audio units can generally be expected to support the [AVAudioFormat]
// standard format (deinterleaved 32-bit float), at any sample rate.
//
// Channel counts can be more complex. See the
// [AUAudioUnit.ChannelCapabilities] reference for a more complete discussion.
//
// Initialization will fail and return an error if the specified format is
// unsupported for the bus.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/init(format:)
//
// [AVAudioFormat]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat
func (a AUAudioUnitBus) InitWithFormatError(format objectivec.IObject) (AUAudioUnitBus, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithFormat:error:"), format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AUAudioUnitBus{}, foundation.NSErrorFrom(errorPtr)
	}
	return AUAudioUnitBusFromID(rv), nil

}

// The audio format and channel layout of audio being transferred on the bus.
//
// # Discussion
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_StreamFormat` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/format
func (a AUAudioUnitBus) Format() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("format"))
	return objectivec.Object{ID: rv}
}

// Determines whether the bus is active.
//
// # Discussion
//
// Hosts must enable input busses before using them. This allows an audio unit
// to be prepared to render a large number of inputs, but avoid the work of
// preparing to pull inputs which are not in use.
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_MakeConnection` and
// `kAudioUnitProperty_SetRenderCallback` APIs.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/isEnabled
func (a AUAudioUnitBus) IsEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isEnabled"))
	return rv
}
func (a AUAudioUnitBus) SetEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setEnabled:"), value)
}

// A name for the bus.
//
// # Discussion
//
// The bus name can be set by the host.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/name
func (a AUAudioUnitBus) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (a AUAudioUnitBus) SetName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setName:"), objc.String(value))
}

// The index of this bus in its containing array.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/index
func (a AUAudioUnitBus) Index() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("index"))
	return rv
}

// The bus type.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/busType
func (a AUAudioUnitBus) BusType() AUAudioUnitBusType {
	rv := objc.Send[AUAudioUnitBusType](a.ID, objc.Sel("busType"))
	return AUAudioUnitBusType(rv)
}

// The audio unit that owns the bus.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/ownerAudioUnit
func (a AUAudioUnitBus) OwnerAudioUnit() IAUAudioUnit {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("ownerAudioUnit"))
	return AUAudioUnitFromID(objc.ID(rv))
}

// An array of audio channel layout tags.
//
// # Discussion
//
// The array contains [NSNumber] objects representing [AudioChannelLayoutTag]
// values.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/supportedChannelLayoutTags
//
// [AudioChannelLayoutTag]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelLayoutTag
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (a AUAudioUnitBus) SupportedChannelLayoutTags() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("supportedChannelLayoutTags"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Information about latency in the audio unit’s processing context.
//
// # Discussion
//
// A host may set this property to describe the presentation latency, in
// seconds, of its input and/or output audio data. A value of `0` means either
// no latency or unknown latency.
//
// A host should set this property on each active bus, since, for example, the
// audio routing path to each of multiple output busses may differ. The
// meaning of this property’s value differs between input and output busses,
// as described below:
//
// - For input busses, this value describes how long ago the audio arriving on
// this bus was acquired.
//
// For example, when reading from a file to the first audio unit in a chain,
// the input presentation latency is zero. For audio input from a device, this
// initial input latency is the presentation latency of the device itself
// (i.e. the device’s offset and latency). A second chained audio unit’s
// input presentation latency is the input presentation latency of the first
// unit, plus the processing latency of the first unit.
//
// - For output busses, this value describes how long it will be before the
// output audio of an audio unit is presented.
//
// For example, when writing to a file, the output presentation latency of the
// last audio unit in a chain is zero. When the audio from that audio unit is
// to be played to a device, then that initial presentation latency will be
// the presentation latency of the device itself (i.e. the I/O buffer size)
// plus the device’s safety offset and latency. A previously chained audio
// unit’s output presentation latency is the last unit’s presentation
// latency plus its processing latency.
//
// For a given audio unit anywhere within a mixing graph, the input and output
// presentation latencies describe to that unit how long from the moment of
// generation it has taken for its input to arrive, and how long it will take
// for its output to be presented.
//
// This version 3 property is bridged to the version 2
// `kAudioUnitProperty_PresentationLatency` API.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/contextPresentationLatency
func (a AUAudioUnitBus) ContextPresentationLatency() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](a.ID, objc.Sel("contextPresentationLatency"))
	return foundation.NSTimeInterval(rv)
}
func (a AUAudioUnitBus) SetContextPresentationLatency(value foundation.NSTimeInterval) {
	objc.Send[struct{}](a.ID, objc.Sel("setContextPresentationLatency:"), value)
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/shouldAllocateBuffer
func (a AUAudioUnitBus) ShouldAllocateBuffer() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("shouldAllocateBuffer"))
	return rv
}
func (a AUAudioUnitBus) SetShouldAllocateBuffer(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setShouldAllocateBuffer:"), value)
}

// An array of numbers indicating the supported number of channels for this
// bus.
//
// # Discussion
//
// If the value of this property is `nil`, then any number less than or equal
// to the value of [AUAudioUnitBus.MaximumChannelCount] is supported. If
// setting a new value on this property makes the current bus format
// unsupported, then the value of [AUAudioUnitBus.Format] is set to `nil`.
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/supportedChannelCounts
func (a AUAudioUnitBus) SupportedChannelCounts() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("supportedChannelCounts"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (a AUAudioUnitBus) SetSupportedChannelCounts(value []foundation.NSNumber) {
	objc.Send[struct{}](a.ID, objc.Sel("setSupportedChannelCounts:"), objectivec.IObjectSliceToNSArray(value))
}

// The maximum number of channels supported for this bus.
//
// # Discussion
//
// If the value of [AUAudioUnitBus.SupportedChannelCounts] is set, then this
// value is derived from it. If setting a new value on this property makes the
// current bus format unsupported, then the value of [AUAudioUnitBus.Format]
// is set to `nil`.
//
// The default value is `UINT_MAX`.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/maximumChannelCount
func (a AUAudioUnitBus) MaximumChannelCount() AUAudioChannelCount {
	rv := objc.Send[AUAudioChannelCount](a.ID, objc.Sel("maximumChannelCount"))
	return AUAudioChannelCount(rv)
}
func (a AUAudioUnitBus) SetMaximumChannelCount(value AUAudioChannelCount) {
	objc.Send[struct{}](a.ID, objc.Sel("setMaximumChannelCount:"), value)
}
