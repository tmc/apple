// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEIOSurfaceObject] class.
var (
	_ANEIOSurfaceObjectClass     ANEIOSurfaceObjectClass
	_ANEIOSurfaceObjectClassOnce sync.Once
)

func getANEIOSurfaceObjectClass() ANEIOSurfaceObjectClass {
	_ANEIOSurfaceObjectClassOnce.Do(func() {
		_ANEIOSurfaceObjectClass = ANEIOSurfaceObjectClass{class: objc.GetClass("_ANEIOSurfaceObject")}
	})
	return _ANEIOSurfaceObjectClass
}

// GetANEIOSurfaceObjectClass returns the class object for _ANEIOSurfaceObject.
func GetANEIOSurfaceObjectClass() ANEIOSurfaceObjectClass {
	return getANEIOSurfaceObjectClass()
}

type ANEIOSurfaceObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEIOSurfaceObjectClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEIOSurfaceObjectClass) Alloc() ANEIOSurfaceObject {
	rv := objc.SendIfResponds[ANEIOSurfaceObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEIOSurfaceObject.EncodeWithCoder]
//   - [ANEIOSurfaceObject.IoSurface]
//   - [ANEIOSurfaceObject.StartOffset]
//   - [ANEIOSurfaceObject.InitWithCoder]
//   - [ANEIOSurfaceObject.InitWithIOSurfaceStartOffsetShouldRetain]
type ANEIOSurfaceObject struct {
	objectivec.Object
}

// ANEIOSurfaceObjectFromID constructs a [ANEIOSurfaceObject] from an objc.ID.
func ANEIOSurfaceObjectFromID(id objc.ID) ANEIOSurfaceObject {
	return ANEIOSurfaceObject{objectivec.Object{ID: id}}
}

// Ensure ANEIOSurfaceObject implements IANEIOSurfaceObject.
var _ IANEIOSurfaceObject = ANEIOSurfaceObject{}

// An interface definition for the [ANEIOSurfaceObject] class.
//
// # Methods
//
//   - [IANEIOSurfaceObject.EncodeWithCoder]
//   - [IANEIOSurfaceObject.IoSurface]
//   - [IANEIOSurfaceObject.StartOffset]
//   - [IANEIOSurfaceObject.InitWithCoder]
//   - [IANEIOSurfaceObject.InitWithIOSurfaceStartOffsetShouldRetain]
type IANEIOSurfaceObject interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	IoSurface() iosurface.IOSurfaceRef
	StartOffset() foundation.NSNumber
	InitWithCoder(coder foundation.INSCoder) ANEIOSurfaceObject
	InitWithIOSurfaceStartOffsetShouldRetain(iOSurface iosurface.IOSurfaceRef, offset objectivec.IObject, retain bool) ANEIOSurfaceObject
}

// Init initializes the instance.
func (a ANEIOSurfaceObject) Init() ANEIOSurfaceObject {
	rv := objc.SendIfResponds[ANEIOSurfaceObject](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEIOSurfaceObject) Autorelease() ANEIOSurfaceObject {
	rv := objc.SendIfResponds[ANEIOSurfaceObject](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEIOSurfaceObject creates a new ANEIOSurfaceObject instance.
func NewANEIOSurfaceObject() ANEIOSurfaceObject {
	class := getANEIOSurfaceObjectClass()
	rv := objc.SendIfResponds[ANEIOSurfaceObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEIOSurfaceObjectWithCoder(coder objectivec.IObject) ANEIOSurfaceObject {
	instance := getANEIOSurfaceObjectClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEIOSurfaceObjectFromID(rv)
}

func NewANEIOSurfaceObjectWithIOSurfaceStartOffsetShouldRetain(iOSurface iosurface.IOSurfaceRef, offset objectivec.IObject, retain bool) ANEIOSurfaceObject {
	instance := getANEIOSurfaceObjectClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithIOSurface:startOffset:shouldRetain:"), iOSurface, offset, retain)
	return ANEIOSurfaceObjectFromID(rv)
}

func (a ANEIOSurfaceObject) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (a ANEIOSurfaceObject) InitWithCoder(coder foundation.INSCoder) ANEIOSurfaceObject {
	rv := objc.SendIfResponds[ANEIOSurfaceObject](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a ANEIOSurfaceObject) InitWithIOSurfaceStartOffsetShouldRetain(iOSurface iosurface.IOSurfaceRef, offset objectivec.IObject, retain bool) ANEIOSurfaceObject {
	rv := objc.SendIfResponds[ANEIOSurfaceObject](a.ID, objc.Sel("initWithIOSurface:startOffset:shouldRetain:"), iOSurface, offset, retain)
	return rv
}

func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) CreateIOSurfaceWithWidthPixel_sizeHeight(width int, pixel_size int, height int) iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("createIOSurfaceWithWidth:pixel_size:height:"), width, pixel_size, height)
	return iosurface.IOSurfaceRef(rv)
}
func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) CreateIOSurfaceWithWidthPixel_sizeHeightBytesPerElement(width int, pixel_size int, height int, element int) iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("createIOSurfaceWithWidth:pixel_size:height:bytesPerElement:"), width, pixel_size, height, element)
	return iosurface.IOSurfaceRef(rv)
}
func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) ObjectWithIOSurface(iOSurface iosurface.IOSurfaceRef) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("objectWithIOSurface:"), iOSurface)
	return objectivec.Object{ID: rv}
}
func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) ObjectWithIOSurfaceStartOffset(iOSurface iosurface.IOSurfaceRef, offset objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("objectWithIOSurface:startOffset:"), iOSurface, offset)
	return objectivec.Object{ID: rv}
}
func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) ObjectWithIOSurfaceNoRetainStartOffset(retain iosurface.IOSurfaceRef, offset objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("objectWithIOSurfaceNoRetain:startOffset:"), retain, offset)
	return objectivec.Object{ID: rv}
}
func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (a ANEIOSurfaceObject) IoSurface() iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](a.ID, objc.Sel("ioSurface"))
	return iosurface.IOSurfaceRef(rv)
}
func (a ANEIOSurfaceObject) StartOffset() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("startOffset"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
