// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSequnceAsFeatureValueArray] class.
var (
	_MLSequnceAsFeatureValueArrayClass     MLSequnceAsFeatureValueArrayClass
	_MLSequnceAsFeatureValueArrayClassOnce sync.Once
)

func getMLSequnceAsFeatureValueArrayClass() MLSequnceAsFeatureValueArrayClass {
	_MLSequnceAsFeatureValueArrayClassOnce.Do(func() {
		_MLSequnceAsFeatureValueArrayClass = MLSequnceAsFeatureValueArrayClass{class: objc.GetClass("MLSequnceAsFeatureValueArray")}
	})
	return _MLSequnceAsFeatureValueArrayClass
}

// GetMLSequnceAsFeatureValueArrayClass returns the class object for MLSequnceAsFeatureValueArray.
func GetMLSequnceAsFeatureValueArrayClass() MLSequnceAsFeatureValueArrayClass {
	return getMLSequnceAsFeatureValueArrayClass()
}

type MLSequnceAsFeatureValueArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSequnceAsFeatureValueArrayClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSequnceAsFeatureValueArrayClass) Alloc() MLSequnceAsFeatureValueArray {
	rv := objc.Send[MLSequnceAsFeatureValueArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSequnceAsFeatureValueArray.Sequence]
//   - [MLSequnceAsFeatureValueArray.InitWrappingSequence]
type MLSequnceAsFeatureValueArray struct {
	foundation.NSArray
}

// MLSequnceAsFeatureValueArrayFromID constructs a [MLSequnceAsFeatureValueArray] from an objc.ID.
func MLSequnceAsFeatureValueArrayFromID(id objc.ID) MLSequnceAsFeatureValueArray {
	return MLSequnceAsFeatureValueArray{NSArray: foundation.NSArrayFromID(id)}
}

// Ensure MLSequnceAsFeatureValueArray implements IMLSequnceAsFeatureValueArray.
var _ IMLSequnceAsFeatureValueArray = MLSequnceAsFeatureValueArray{}

// An interface definition for the [MLSequnceAsFeatureValueArray] class.
//
// # Methods
//
//   - [IMLSequnceAsFeatureValueArray.Sequence]
//   - [IMLSequnceAsFeatureValueArray.InitWrappingSequence]
type IMLSequnceAsFeatureValueArray interface {
	foundation.INSArray

	// Topic: Methods

	Sequence() IMLSequence
	InitWrappingSequence(sequence objectivec.IObject) MLSequnceAsFeatureValueArray
}

// Init initializes the instance.
func (m MLSequnceAsFeatureValueArray) Init() MLSequnceAsFeatureValueArray {
	rv := objc.Send[MLSequnceAsFeatureValueArray](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSequnceAsFeatureValueArray) Autorelease() MLSequnceAsFeatureValueArray {
	rv := objc.Send[MLSequnceAsFeatureValueArray](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSequnceAsFeatureValueArray creates a new MLSequnceAsFeatureValueArray instance.
func NewMLSequnceAsFeatureValueArray() MLSequnceAsFeatureValueArray {
	class := getMLSequnceAsFeatureValueArrayClass()
	rv := objc.Send[MLSequnceAsFeatureValueArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSequnceAsFeatureValueArrayWrappingSequence(sequence objectivec.IObject) MLSequnceAsFeatureValueArray {
	instance := getMLSequnceAsFeatureValueArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWrappingSequence:"), sequence)
	return MLSequnceAsFeatureValueArrayFromID(rv)
}

func (m MLSequnceAsFeatureValueArray) InitWrappingSequence(sequence objectivec.IObject) MLSequnceAsFeatureValueArray {
	rv := objc.Send[MLSequnceAsFeatureValueArray](m.ID, objc.Sel("initWrappingSequence:"), sequence)
	return rv
}

func (m MLSequnceAsFeatureValueArray) Sequence() IMLSequence {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sequence"))
	return MLSequenceFromID(objc.ID(rv))
}
