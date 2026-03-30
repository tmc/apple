// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSequence] class.
var (
	_MLSequenceClass     MLSequenceClass
	_MLSequenceClassOnce sync.Once
)

func getMLSequenceClass() MLSequenceClass {
	_MLSequenceClassOnce.Do(func() {
		_MLSequenceClass = MLSequenceClass{class: objc.GetClass("MLSequence")}
	})
	return _MLSequenceClass
}

// GetMLSequenceClass returns the class object for MLSequence.
func GetMLSequenceClass() MLSequenceClass {
	return getMLSequenceClass()
}

type MLSequenceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSequenceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSequenceClass) Alloc() MLSequence {
	rv := objc.Send[MLSequence](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A machine learning collection type that stores a series of strings or
// integers.
//
// # Overview
//
// A sequence stores a series of integers or strings of any length as the
// underlying type of an [MLFeatureValue]. Some classifier models —
// typically natural language models, such as an [NLTagger] — produce an
// [MLSequence] feature value from their output features.
//
// # Identifying the sequence’s element type
//
//   - [MLSequence.Type]: The underlying type of the sequence’s elements.
//
// # Retrieving the Sequence’s Values
//
//   - [MLSequence.StringValues]: An array of strings in the sequence.
//   - [MLSequence.Int64Values]: An array of 64-bit integers in the sequence.
//
// See: https://developer.apple.com/documentation/CoreML/MLSequence
//
// [NLTagger]: https://developer.apple.com/documentation/NaturalLanguage/NLTagger
type MLSequence struct {
	objectivec.Object
}

// MLSequenceFromID constructs a [MLSequence] from an objc.ID.
//
// A machine learning collection type that stores a series of strings or
// integers.
func MLSequenceFromID(id objc.ID) MLSequence {
	return MLSequence{objectivec.Object{ID: id}}
}

// NOTE: MLSequence adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLSequence] class.
//
// # Identifying the sequence’s element type
//
//   - [IMLSequence.Type]: The underlying type of the sequence’s elements.
//
// # Retrieving the Sequence’s Values
//
//   - [IMLSequence.StringValues]: An array of strings in the sequence.
//   - [IMLSequence.Int64Values]: An array of 64-bit integers in the sequence.
//
// See: https://developer.apple.com/documentation/CoreML/MLSequence
type IMLSequence interface {
	objectivec.IObject

	// Topic: Identifying the sequence’s element type

	// The underlying type of the sequence’s elements.
	Type() MLFeatureType

	// Topic: Retrieving the Sequence’s Values

	// An array of strings in the sequence.
	StringValues() []string
	// An array of 64-bit integers in the sequence.
	Int64Values() []foundation.NSNumber

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s MLSequence) Init() MLSequence {
	rv := objc.Send[MLSequence](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSequence) Autorelease() MLSequence {
	rv := objc.Send[MLSequence](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSequence creates a new MLSequence instance.
func NewMLSequence() MLSequence {
	class := getMLSequenceClass()
	rv := objc.Send[MLSequence](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an empty sequence of strings or integers.
//
// type: An [MLFeatureType] instance that determines the sequence’s element type,
// which must be either [MLFeatureTypeString] or [MLFeatureTypeInt64].
//
// See: https://developer.apple.com/documentation/CoreML/MLSequence/init(empty:)
//
// [MLFeatureType]: https://developer.apple.com/documentation/CoreML/MLFeatureType
func NewSequenceEmptySequenceWithType(type_ MLFeatureType) MLSequence {
	rv := objc.Send[objc.ID](objc.ID(getMLSequenceClass().class), objc.Sel("emptySequenceWithType:"), type_)
	return MLSequenceFromID(rv)
}

// Creates a sequence of integers from an array of numbers.
//
// int64Values: An array of integer values represented as [NSNumber] instances.
//
// See: https://developer.apple.com/documentation/CoreML/MLSequence/init(int64s:)
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func NewSequenceWithInt64Array(int64Values []foundation.NSNumber) MLSequence {
	rv := objc.Send[objc.ID](objc.ID(getMLSequenceClass().class), objc.Sel("sequenceWithInt64Array:"), objectivec.IObjectSliceToNSArray(int64Values))
	return MLSequenceFromID(rv)
}

// Creates a sequence of strings from a string array.
//
// stringValues: The array of strings for the sequence.
//
// See: https://developer.apple.com/documentation/CoreML/MLSequence/init(strings:)
func NewSequenceWithStringArray(stringValues []string) MLSequence {
	rv := objc.Send[objc.ID](objc.ID(getMLSequenceClass().class), objc.Sel("sequenceWithStringArray:"), objectivec.StringSliceToNSArray(stringValues))
	return MLSequenceFromID(rv)
}

func (s MLSequence) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The underlying type of the sequence’s elements.
//
// # Discussion
//
// The sequence’s underlying element type can only be either
// [MLFeatureTypeString] or [MLFeatureTypeInt64]. Use this value to determine
// whether to access [StringValues] or [Int64Values] at runtime.
//
// See: https://developer.apple.com/documentation/CoreML/MLSequence/type
func (s MLSequence) Type() MLFeatureType {
	rv := objc.Send[MLFeatureType](s.ID, objc.Sel("type"))
	return MLFeatureType(rv)
}

// An array of strings in the sequence.
//
// # Discussion
//
// Only use this property when the sequence’s [Type] is
// [MLFeatureTypeString].
//
// See: https://developer.apple.com/documentation/CoreML/MLSequence/stringValues
func (s MLSequence) StringValues() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("stringValues"))
	return objc.ConvertSliceToStrings(rv)
}

// An array of 64-bit integers in the sequence.
//
// # Discussion
//
// Only use this property when the sequence’s [Type] is
// [MLFeatureTypeInt64].
//
// See: https://developer.apple.com/documentation/CoreML/MLSequence/int64Values
func (s MLSequence) Int64Values() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("int64Values"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
