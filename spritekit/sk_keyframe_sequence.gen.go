// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKKeyframeSequence] class.
var (
	_SKKeyframeSequenceClass     SKKeyframeSequenceClass
	_SKKeyframeSequenceClassOnce sync.Once
)

func getSKKeyframeSequenceClass() SKKeyframeSequenceClass {
	_SKKeyframeSequenceClassOnce.Do(func() {
		_SKKeyframeSequenceClass = SKKeyframeSequenceClass{class: objc.GetClass("SKKeyframeSequence")}
	})
	return _SKKeyframeSequenceClass
}

// GetSKKeyframeSequenceClass returns the class object for SKKeyframeSequence.
func GetSKKeyframeSequenceClass() SKKeyframeSequenceClass {
	return getSKKeyframeSequenceClass()
}

type SKKeyframeSequenceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKKeyframeSequenceClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKKeyframeSequenceClass) Alloc() SKKeyframeSequence {
	rv := objc.Send[SKKeyframeSequence](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object that performs interpolation between values specified at different
// times (keyframes).
//
// # Overview
//
// The primary use for an [SKKeyframeSequence] object is to animate properties
// on particles emitted by an [SKEmitterNode] object, but it can also be used
// for your general interpolation needs across a discrete set of inputs.
//
// When a keyframe sequence is used with an emitter node, particles determine
// their values by sampling the keyframe sequence. The sequence replaces the
// normal simulation performed by the emitter node.
//
// # First Steps
//
//   - [SKKeyframeSequence.InitWithKeyframeValuesTimes]: Initializes a keyframe sequence with an initial set of values and times.
//   - [SKKeyframeSequence.InitWithCapacity]: Initializes a new keyframe sequence.
//   - [SKKeyframeSequence.InitWithCoder]
//
// # Sequence Building
//
//   - [SKKeyframeSequence.AddKeyframeValueTime]: Adds a keyframe to the sequence.
//   - [SKKeyframeSequence.RemoveKeyframeAtIndex]: Removes a keyframe from the sequence.
//   - [SKKeyframeSequence.RemoveLastKeyframe]: Removes the last value in the sequence.
//   - [SKKeyframeSequence.SetKeyframeTimeForIndex]: Changes the time for a specific keyframe.
//   - [SKKeyframeSequence.SetKeyframeValueForIndex]: Changes the value for a specific keyframe.
//   - [SKKeyframeSequence.SetKeyframeValueTimeForIndex]: Replaces a keyframe in the sequence with a new keyframe.
//
// # Sequence Running
//
//   - [SKKeyframeSequence.SampleAtTime]: Calculates the sample at a particular time.
//
// # Sequence Information
//
//   - [SKKeyframeSequence.Count]: The number of keyframes in the sequence.
//   - [SKKeyframeSequence.GetKeyframeTimeForIndex]: Gets the time for a keyframe in the sequence.
//   - [SKKeyframeSequence.GetKeyframeValueForIndex]: Gets the value for a keyframe in the sequence.
//
// # Interpolation Modifiers
//
//   - [SKKeyframeSequence.InterpolationMode]: The mode used to determine how values for times between the keyframes are calculated.
//   - [SKKeyframeSequence.SetInterpolationMode]
//   - [SKKeyframeSequence.RepeatMode]: The mode used to determine how the keyframe sequence repeats.
//   - [SKKeyframeSequence.SetRepeatMode]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence
type SKKeyframeSequence struct {
	objectivec.Object
}

// SKKeyframeSequenceFromID constructs a [SKKeyframeSequence] from an objc.ID.
//
// An object that performs interpolation between values specified at different
// times (keyframes).
func SKKeyframeSequenceFromID(id objc.ID) SKKeyframeSequence {
	return SKKeyframeSequence{objectivec.Object{ID: id}}
}

// NOTE: SKKeyframeSequence adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKKeyframeSequence] class.
//
// # First Steps
//
//   - [ISKKeyframeSequence.InitWithKeyframeValuesTimes]: Initializes a keyframe sequence with an initial set of values and times.
//   - [ISKKeyframeSequence.InitWithCapacity]: Initializes a new keyframe sequence.
//   - [ISKKeyframeSequence.InitWithCoder]
//
// # Sequence Building
//
//   - [ISKKeyframeSequence.AddKeyframeValueTime]: Adds a keyframe to the sequence.
//   - [ISKKeyframeSequence.RemoveKeyframeAtIndex]: Removes a keyframe from the sequence.
//   - [ISKKeyframeSequence.RemoveLastKeyframe]: Removes the last value in the sequence.
//   - [ISKKeyframeSequence.SetKeyframeTimeForIndex]: Changes the time for a specific keyframe.
//   - [ISKKeyframeSequence.SetKeyframeValueForIndex]: Changes the value for a specific keyframe.
//   - [ISKKeyframeSequence.SetKeyframeValueTimeForIndex]: Replaces a keyframe in the sequence with a new keyframe.
//
// # Sequence Running
//
//   - [ISKKeyframeSequence.SampleAtTime]: Calculates the sample at a particular time.
//
// # Sequence Information
//
//   - [ISKKeyframeSequence.Count]: The number of keyframes in the sequence.
//   - [ISKKeyframeSequence.GetKeyframeTimeForIndex]: Gets the time for a keyframe in the sequence.
//   - [ISKKeyframeSequence.GetKeyframeValueForIndex]: Gets the value for a keyframe in the sequence.
//
// # Interpolation Modifiers
//
//   - [ISKKeyframeSequence.InterpolationMode]: The mode used to determine how values for times between the keyframes are calculated.
//   - [ISKKeyframeSequence.SetInterpolationMode]
//   - [ISKKeyframeSequence.RepeatMode]: The mode used to determine how the keyframe sequence repeats.
//   - [ISKKeyframeSequence.SetRepeatMode]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence
type ISKKeyframeSequence interface {
	objectivec.IObject

	// Topic: First Steps

	// Initializes a keyframe sequence with an initial set of values and times.
	InitWithKeyframeValuesTimes(values foundation.INSArray, times []foundation.NSNumber) SKKeyframeSequence
	// Initializes a new keyframe sequence.
	InitWithCapacity(numItems uint) SKKeyframeSequence
	InitWithCoder(aDecoder foundation.INSCoder) SKKeyframeSequence

	// Topic: Sequence Building

	// Adds a keyframe to the sequence.
	AddKeyframeValueTime(value objectivec.IObject, time float64)
	// Removes a keyframe from the sequence.
	RemoveKeyframeAtIndex(index uint)
	// Removes the last value in the sequence.
	RemoveLastKeyframe()
	// Changes the time for a specific keyframe.
	SetKeyframeTimeForIndex(time float64, index uint)
	// Changes the value for a specific keyframe.
	SetKeyframeValueForIndex(value objectivec.IObject, index uint)
	// Replaces a keyframe in the sequence with a new keyframe.
	SetKeyframeValueTimeForIndex(value objectivec.IObject, time float64, index uint)

	// Topic: Sequence Running

	// Calculates the sample at a particular time.
	SampleAtTime(time float64) objectivec.IObject

	// Topic: Sequence Information

	// The number of keyframes in the sequence.
	Count() uint
	// Gets the time for a keyframe in the sequence.
	GetKeyframeTimeForIndex(index uint) float64
	// Gets the value for a keyframe in the sequence.
	GetKeyframeValueForIndex(index uint) objectivec.IObject

	// Topic: Interpolation Modifiers

	// The mode used to determine how values for times between the keyframes are calculated.
	InterpolationMode() SKInterpolationMode
	SetInterpolationMode(value SKInterpolationMode)
	// The mode used to determine how the keyframe sequence repeats.
	RepeatMode() SKRepeatMode
	SetRepeatMode(value SKRepeatMode)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (k SKKeyframeSequence) Init() SKKeyframeSequence {
	rv := objc.Send[SKKeyframeSequence](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k SKKeyframeSequence) Autorelease() SKKeyframeSequence {
	rv := objc.Send[SKKeyframeSequence](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKKeyframeSequence creates a new SKKeyframeSequence instance.
func NewSKKeyframeSequence() SKKeyframeSequence {
	class := getSKKeyframeSequenceClass()
	rv := objc.Send[SKKeyframeSequence](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new keyframe sequence.
//
// numItems: The initial capacity of the new sequence.
//
// # Return Value
//
// A newly initialized empty sequence.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/init(capacity:)
func NewKeyframeSequenceWithCapacity(numItems uint) SKKeyframeSequence {
	instance := getSKKeyframeSequenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	return SKKeyframeSequenceFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/init(coder:)
func NewKeyframeSequenceWithCoder(aDecoder foundation.INSCoder) SKKeyframeSequence {
	instance := getSKKeyframeSequenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKKeyframeSequenceFromID(rv)
}

// Initializes a keyframe sequence with an initial set of values and times.
//
// values: An array of value objects that define the keyframe values for the sequence.
//
// times: An array of [NSNumber] objects containing floating-point values that
// specify the time values for the keyframes.
//
// # Return Value
//
// A newly initialized sequence.
//
// # Discussion
//
// The two arrays must have an identical number of elements. The keyframes in
// the new sequence are stored in the same order as the input arrays.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/init(keyframeValues:times:)
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func NewKeyframeSequenceWithKeyframeValuesTimes(values foundation.INSArray, times []foundation.NSNumber) SKKeyframeSequence {
	instance := getSKKeyframeSequenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyframeValues:times:"), values, objectivec.IObjectSliceToNSArray(times))
	return SKKeyframeSequenceFromID(rv)
}

// Initializes a keyframe sequence with an initial set of values and times.
//
// values: An array of value objects that define the keyframe values for the sequence.
//
// times: An array of [NSNumber] objects containing floating-point values that
// specify the time values for the keyframes.
//
// # Return Value
//
// A newly initialized sequence.
//
// # Discussion
//
// The two arrays must have an identical number of elements. The keyframes in
// the new sequence are stored in the same order as the input arrays.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/init(keyframeValues:times:)
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (k SKKeyframeSequence) InitWithKeyframeValuesTimes(values foundation.INSArray, times []foundation.NSNumber) SKKeyframeSequence {
	rv := objc.Send[SKKeyframeSequence](k.ID, objc.Sel("initWithKeyframeValues:times:"), values, objectivec.IObjectSliceToNSArray(times))
	return rv
}

// Initializes a new keyframe sequence.
//
// numItems: The initial capacity of the new sequence.
//
// # Return Value
//
// A newly initialized empty sequence.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/init(capacity:)
func (k SKKeyframeSequence) InitWithCapacity(numItems uint) SKKeyframeSequence {
	rv := objc.Send[SKKeyframeSequence](k.ID, objc.Sel("initWithCapacity:"), numItems)
	return rv
}

// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/init(coder:)
func (k SKKeyframeSequence) InitWithCoder(aDecoder foundation.INSCoder) SKKeyframeSequence {
	rv := objc.Send[SKKeyframeSequence](k.ID, objc.Sel("initWithCoder:"), aDecoder)
	return rv
}

// Adds a keyframe to the sequence.
//
// value: An object that defines the value to add. It must have the same class as
// other value objects stored in the sequence.
//
// time: The corresponding time.
//
// # Discussion
//
// The new keyframe is appended to the end of the array.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/addKeyframeValue(_:time:)
func (k SKKeyframeSequence) AddKeyframeValueTime(value objectivec.IObject, time float64) {
	objc.Send[objc.ID](k.ID, objc.Sel("addKeyframeValue:time:"), value, time)
}

// Removes a keyframe from the sequence.
//
// index: The index of the keyframe value to remove.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/removeKeyframe(at:)
func (k SKKeyframeSequence) RemoveKeyframeAtIndex(index uint) {
	objc.Send[objc.ID](k.ID, objc.Sel("removeKeyframeAtIndex:"), index)
}

// Removes the last value in the sequence.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/removeLastKeyframe()
func (k SKKeyframeSequence) RemoveLastKeyframe() {
	objc.Send[objc.ID](k.ID, objc.Sel("removeLastKeyframe"))
}

// Changes the time for a specific keyframe.
//
// time: The new time value for the keyframe.
//
// index: The index of the keyframe to change.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/setKeyframeTime(_:for:)
func (k SKKeyframeSequence) SetKeyframeTimeForIndex(time float64, index uint) {
	objc.Send[objc.ID](k.ID, objc.Sel("setKeyframeTime:forIndex:"), time, index)
}

// Changes the value for a specific keyframe.
//
// value: The new value for the keyframe.
//
// index: The index of the keyframe to change.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/setKeyframeValue(_:for:)
func (k SKKeyframeSequence) SetKeyframeValueForIndex(value objectivec.IObject, index uint) {
	objc.Send[objc.ID](k.ID, objc.Sel("setKeyframeValue:forIndex:"), value, index)
}

// Replaces a keyframe in the sequence with a new keyframe.
//
// value: The new value for the keyframe.
//
// time: The new time for the keyframe.
//
// index: The index of the keyframe to change.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/setKeyframeValue(_:time:for:)
func (k SKKeyframeSequence) SetKeyframeValueTimeForIndex(value objectivec.IObject, time float64, index uint) {
	objc.Send[objc.ID](k.ID, objc.Sel("setKeyframeValue:time:forIndex:"), value, time, index)
}

// Calculates the sample at a particular time.
//
// time: The time value to sample.
//
// # Return Value
//
// An object that contains the interpolated sample. The class of this object
// matches the class of the values stored in the keyframe sequence — either
// an [NSNumber] or an [SKColor].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/sample(atTime:)
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (k SKKeyframeSequence) SampleAtTime(time float64) objectivec.IObject {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("sampleAtTime:"), time)
	return objectivec.Object{ID: rv}
}

// The number of keyframes in the sequence.
//
// # Return Value
//
// The number of keyframes in the sequence.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/count()
func (k SKKeyframeSequence) Count() uint {
	rv := objc.Send[uint](k.ID, objc.Sel("count"))
	return rv
}

// Gets the time for a keyframe in the sequence.
//
// index: The index of the keyframe.
//
// # Return Value
//
// The time value for the keyframe.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/getKeyframeTime(for:)
func (k SKKeyframeSequence) GetKeyframeTimeForIndex(index uint) float64 {
	rv := objc.Send[float64](k.ID, objc.Sel("getKeyframeTimeForIndex:"), index)
	return rv
}

// Gets the value for a keyframe in the sequence.
//
// index: The index of the keyframe.
//
// # Return Value
//
// The value object for the keyframe.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/getKeyframeValue(for:)
func (k SKKeyframeSequence) GetKeyframeValueForIndex(index uint) objectivec.IObject {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("getKeyframeValueForIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (k SKKeyframeSequence) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](k.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The mode used to determine how values for times between the keyframes are
// calculated.
//
// # Discussion
//
// The possible values are defined in [SKInterpolationMode]. The default value
// is [SKInterpolationMode.linear].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/interpolationMode
//
// [SKInterpolationMode.linear]: https://developer.apple.com/documentation/SpriteKit/SKInterpolationMode/linear
// [SKInterpolationMode]: https://developer.apple.com/documentation/SpriteKit/SKInterpolationMode
func (k SKKeyframeSequence) InterpolationMode() SKInterpolationMode {
	rv := objc.Send[SKInterpolationMode](k.ID, objc.Sel("interpolationMode"))
	return SKInterpolationMode(rv)
}
func (k SKKeyframeSequence) SetInterpolationMode(value SKInterpolationMode) {
	objc.Send[struct{}](k.ID, objc.Sel("setInterpolationMode:"), value)
}

// The mode used to determine how the keyframe sequence repeats.
//
// # Discussion
//
// The possible values are defined in [SKRepeatMode]. The default value is
// [SKRepeatMode.clamp].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKKeyframeSequence/repeatMode
//
// [SKRepeatMode.clamp]: https://developer.apple.com/documentation/SpriteKit/SKRepeatMode/clamp
// [SKRepeatMode]: https://developer.apple.com/documentation/SpriteKit/SKRepeatMode
func (k SKKeyframeSequence) RepeatMode() SKRepeatMode {
	rv := objc.Send[SKRepeatMode](k.ID, objc.Sel("repeatMode"))
	return SKRepeatMode(rv)
}
func (k SKKeyframeSequence) SetRepeatMode(value SKRepeatMode) {
	objc.Send[struct{}](k.ID, objc.Sel("setRepeatMode:"), value)
}
