// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEIOSurfaceOutputSets] class.
var (
	_ANEIOSurfaceOutputSetsClass     ANEIOSurfaceOutputSetsClass
	_ANEIOSurfaceOutputSetsClassOnce sync.Once
)

func getANEIOSurfaceOutputSetsClass() ANEIOSurfaceOutputSetsClass {
	_ANEIOSurfaceOutputSetsClassOnce.Do(func() {
		_ANEIOSurfaceOutputSetsClass = ANEIOSurfaceOutputSetsClass{class: objc.GetClass("_ANEIOSurfaceOutputSets")}
	})
	return _ANEIOSurfaceOutputSetsClass
}

// GetANEIOSurfaceOutputSetsClass returns the class object for _ANEIOSurfaceOutputSets.
func GetANEIOSurfaceOutputSetsClass() ANEIOSurfaceOutputSetsClass {
	return getANEIOSurfaceOutputSetsClass()
}

type ANEIOSurfaceOutputSetsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEIOSurfaceOutputSetsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEIOSurfaceOutputSetsClass) Alloc() ANEIOSurfaceOutputSets {
	rv := objc.Send[ANEIOSurfaceOutputSets](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEIOSurfaceOutputSets.EncodeWithCoder]
//   - [ANEIOSurfaceOutputSets.OutputBuffer]
//   - [ANEIOSurfaceOutputSets.StatsSurRef]
//   - [ANEIOSurfaceOutputSets.InitWithCoder]
//   - [ANEIOSurfaceOutputSets.InitWithstatsSurRefOutputBuffer]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceOutputSets
type ANEIOSurfaceOutputSets struct {
	objectivec.Object
}

// ANEIOSurfaceOutputSetsFromID constructs a [ANEIOSurfaceOutputSets] from an objc.ID.
func ANEIOSurfaceOutputSetsFromID(id objc.ID) ANEIOSurfaceOutputSets {
	return ANEIOSurfaceOutputSets{objectivec.Object{ID: id}}
}
// Ensure ANEIOSurfaceOutputSets implements IANEIOSurfaceOutputSets.
var _ IANEIOSurfaceOutputSets = ANEIOSurfaceOutputSets{}

// An interface definition for the [ANEIOSurfaceOutputSets] class.
//
// # Methods
//
//   - [IANEIOSurfaceOutputSets.EncodeWithCoder]
//   - [IANEIOSurfaceOutputSets.OutputBuffer]
//   - [IANEIOSurfaceOutputSets.StatsSurRef]
//   - [IANEIOSurfaceOutputSets.InitWithCoder]
//   - [IANEIOSurfaceOutputSets.InitWithstatsSurRefOutputBuffer]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceOutputSets
type IANEIOSurfaceOutputSets interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	OutputBuffer() foundation.INSArray
	StatsSurRef() coregraphics.IOSurfaceRef
	InitWithCoder(coder foundation.INSCoder) ANEIOSurfaceOutputSets
	InitWithstatsSurRefOutputBuffer(ref coregraphics.IOSurfaceRef, buffer objectivec.IObject) ANEIOSurfaceOutputSets
}

// Init initializes the instance.
func (a ANEIOSurfaceOutputSets) Init() ANEIOSurfaceOutputSets {
	rv := objc.Send[ANEIOSurfaceOutputSets](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEIOSurfaceOutputSets) Autorelease() ANEIOSurfaceOutputSets {
	rv := objc.Send[ANEIOSurfaceOutputSets](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEIOSurfaceOutputSets creates a new ANEIOSurfaceOutputSets instance.
func NewANEIOSurfaceOutputSets() ANEIOSurfaceOutputSets {
	class := getANEIOSurfaceOutputSetsClass()
	rv := objc.Send[ANEIOSurfaceOutputSets](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceOutputSets/initWithCoder:
func NewANEIOSurfaceOutputSetsWithCoder(coder objectivec.IObject) ANEIOSurfaceOutputSets {
	instance := getANEIOSurfaceOutputSetsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEIOSurfaceOutputSetsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceOutputSets/initWithstatsSurRef:outputBuffer:
func NewANEIOSurfaceOutputSetsWithstatsSurRefOutputBuffer(ref coregraphics.IOSurfaceRef, buffer objectivec.IObject) ANEIOSurfaceOutputSets {
	instance := getANEIOSurfaceOutputSetsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithstatsSurRef:outputBuffer:"), ref, buffer)
	return ANEIOSurfaceOutputSetsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceOutputSets/encodeWithCoder:
func (a ANEIOSurfaceOutputSets) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceOutputSets/initWithCoder:
func (a ANEIOSurfaceOutputSets) InitWithCoder(coder foundation.INSCoder) ANEIOSurfaceOutputSets {
	rv := objc.Send[ANEIOSurfaceOutputSets](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceOutputSets/initWithstatsSurRef:outputBuffer:
func (a ANEIOSurfaceOutputSets) InitWithstatsSurRefOutputBuffer(ref coregraphics.IOSurfaceRef, buffer objectivec.IObject) ANEIOSurfaceOutputSets {
	rv := objc.Send[ANEIOSurfaceOutputSets](a.ID, objc.Sel("initWithstatsSurRef:outputBuffer:"), ref, buffer)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceOutputSets/objectWithstatsSurRef:outputBuffer:
func (_ANEIOSurfaceOutputSetsClass ANEIOSurfaceOutputSetsClass) ObjectWithstatsSurRefOutputBuffer(ref coregraphics.IOSurfaceRef, buffer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEIOSurfaceOutputSetsClass.class), objc.Sel("objectWithstatsSurRef:outputBuffer:"), ref, buffer)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceOutputSets/supportsSecureCoding
func (_ANEIOSurfaceOutputSetsClass ANEIOSurfaceOutputSetsClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_ANEIOSurfaceOutputSetsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceOutputSets/outputBuffer
func (a ANEIOSurfaceOutputSets) OutputBuffer() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputBuffer"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceOutputSets/statsSurRef
func (a ANEIOSurfaceOutputSets) StatsSurRef() coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](a.ID, objc.Sel("statsSurRef"))
	return coregraphics.IOSurfaceRef(rv)
}

