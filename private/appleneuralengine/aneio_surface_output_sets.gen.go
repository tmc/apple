// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[ANEIOSurfaceOutputSets](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEIOSurfaceOutputSets.EncodeWithCoder]
//   - [ANEIOSurfaceOutputSets.OutputBuffer]
//   - [ANEIOSurfaceOutputSets.StatsSurRef]
//   - [ANEIOSurfaceOutputSets.InitWithCoder]
//   - [ANEIOSurfaceOutputSets.InitWithstatsSurRefOutputBuffer]
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
type IANEIOSurfaceOutputSets interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	OutputBuffer() foundation.INSArray
	StatsSurRef() iosurface.IOSurfaceRef
	InitWithCoder(coder foundation.INSCoder) ANEIOSurfaceOutputSets
	InitWithstatsSurRefOutputBuffer(ref iosurface.IOSurfaceRef, buffer objectivec.IObject) ANEIOSurfaceOutputSets
}

// Init initializes the instance.
func (a ANEIOSurfaceOutputSets) Init() ANEIOSurfaceOutputSets {
	rv := objc.SendIfResponds[ANEIOSurfaceOutputSets](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEIOSurfaceOutputSets) Autorelease() ANEIOSurfaceOutputSets {
	rv := objc.SendIfResponds[ANEIOSurfaceOutputSets](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEIOSurfaceOutputSets creates a new ANEIOSurfaceOutputSets instance.
func NewANEIOSurfaceOutputSets() ANEIOSurfaceOutputSets {
	class := getANEIOSurfaceOutputSetsClass()
	rv := objc.SendIfResponds[ANEIOSurfaceOutputSets](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEIOSurfaceOutputSetsWithCoder(coder objectivec.IObject) ANEIOSurfaceOutputSets {
	instance := getANEIOSurfaceOutputSetsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEIOSurfaceOutputSetsFromID(rv)
}

func NewANEIOSurfaceOutputSetsWithstatsSurRefOutputBuffer(ref iosurface.IOSurfaceRef, buffer objectivec.IObject) ANEIOSurfaceOutputSets {
	instance := getANEIOSurfaceOutputSetsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithstatsSurRef:outputBuffer:"), ref, buffer)
	return ANEIOSurfaceOutputSetsFromID(rv)
}

func (a ANEIOSurfaceOutputSets) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (a ANEIOSurfaceOutputSets) InitWithCoder(coder foundation.INSCoder) ANEIOSurfaceOutputSets {
	rv := objc.SendIfResponds[ANEIOSurfaceOutputSets](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a ANEIOSurfaceOutputSets) InitWithstatsSurRefOutputBuffer(ref iosurface.IOSurfaceRef, buffer objectivec.IObject) ANEIOSurfaceOutputSets {
	rv := objc.SendIfResponds[ANEIOSurfaceOutputSets](a.ID, objc.Sel("initWithstatsSurRef:outputBuffer:"), ref, buffer)
	return rv
}

func (_ANEIOSurfaceOutputSetsClass ANEIOSurfaceOutputSetsClass) ObjectWithstatsSurRefOutputBuffer(ref iosurface.IOSurfaceRef, buffer objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEIOSurfaceOutputSetsClass.class), objc.Sel("objectWithstatsSurRef:outputBuffer:"), ref, buffer)
	return objectivec.Object{ID: rv}
}
func (_ANEIOSurfaceOutputSetsClass ANEIOSurfaceOutputSetsClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_ANEIOSurfaceOutputSetsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (a ANEIOSurfaceOutputSets) OutputBuffer() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("outputBuffer"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANEIOSurfaceOutputSets) StatsSurRef() iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](a.ID, objc.Sel("statsSurRef"))
	return iosurface.IOSurfaceRef(rv)
}
