// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
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

// Alloc allocates memory for a new instance of the class.
func (ac ANEIOSurfaceObjectClass) Alloc() ANEIOSurfaceObject {
	rv := objc.Send[ANEIOSurfaceObject](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [ANEIOSurfaceObject.EncodeWithCoder]
//   - [ANEIOSurfaceObject.IoSurface]
//   - [ANEIOSurfaceObject.MetalBufferWithDeviceMultiBufferFrame]
//   - [ANEIOSurfaceObject.StartOffset]
//   - [ANEIOSurfaceObject.InitWithCoder]
//   - [ANEIOSurfaceObject.InitWithIOSurfaceStartOffsetShouldRetain]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject
type ANEIOSurfaceObject struct {
	objectivec.Object
}

// ANEIOSurfaceObjectFromID constructs a [ANEIOSurfaceObject] from an objc.ID.
func ANEIOSurfaceObjectFromID(id objc.ID) ANEIOSurfaceObject {
	return ANEIOSurfaceObject{objectivec.Object{id}}
}
// Ensure ANEIOSurfaceObject implements IANEIOSurfaceObject.
var _ IANEIOSurfaceObject = ANEIOSurfaceObject{}





// An interface definition for the [ANEIOSurfaceObject] class.
//
// # Methods
//
//   - [IANEIOSurfaceObject.EncodeWithCoder]
//   - [IANEIOSurfaceObject.IoSurface]
//   - [IANEIOSurfaceObject.MetalBufferWithDeviceMultiBufferFrame]
//   - [IANEIOSurfaceObject.StartOffset]
//   - [IANEIOSurfaceObject.InitWithCoder]
//   - [IANEIOSurfaceObject.InitWithIOSurfaceStartOffsetShouldRetain]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject
type IANEIOSurfaceObject interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder objectivec.IObject)
	IoSurface() coregraphics.IOSurfaceRef
	MetalBufferWithDeviceMultiBufferFrame(device objectivec.IObject, frame uint64) metal.MTLBuffer
	StartOffset() foundation.NSNumber
	InitWithCoder(coder objectivec.IObject) ANEIOSurfaceObject
	InitWithIOSurfaceStartOffsetShouldRetain(iOSurface coregraphics.IOSurfaceRef, offset objectivec.IObject, retain bool) ANEIOSurfaceObject
}





// Init initializes the instance.
func (a ANEIOSurfaceObject) Init() ANEIOSurfaceObject {
	rv := objc.Send[ANEIOSurfaceObject](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEIOSurfaceObject) Autorelease() ANEIOSurfaceObject {
	rv := objc.Send[ANEIOSurfaceObject](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEIOSurfaceObject creates a new ANEIOSurfaceObject instance.
func NewANEIOSurfaceObject() ANEIOSurfaceObject {
	class := getANEIOSurfaceObjectClass()
	rv := objc.Send[ANEIOSurfaceObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/initWithCoder:
func NewANEIOSurfaceObjectWithCoder(coder objectivec.IObject) ANEIOSurfaceObject {
	instance := getANEIOSurfaceObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEIOSurfaceObjectFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/initWithIOSurface:startOffset:shouldRetain:
func NewANEIOSurfaceObjectWithIOSurfaceStartOffsetShouldRetain(iOSurface coregraphics.IOSurfaceRef, offset objectivec.IObject, retain bool) ANEIOSurfaceObject {
	instance := getANEIOSurfaceObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOSurface:startOffset:shouldRetain:"), iOSurface, offset, retain)
	return ANEIOSurfaceObjectFromID(rv)
}







//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/encodeWithCoder:
func (a ANEIOSurfaceObject) EncodeWithCoder(coder objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/metalBufferWithDevice:multiBufferFrame:
func (a ANEIOSurfaceObject) MetalBufferWithDeviceMultiBufferFrame(device objectivec.IObject, frame uint64) metal.MTLBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("metalBufferWithDevice:multiBufferFrame:"), device, frame)
	return metal.MTLBufferObjectFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/initWithCoder:
func (a ANEIOSurfaceObject) InitWithCoder(coder objectivec.IObject) ANEIOSurfaceObject {
	rv := objc.Send[ANEIOSurfaceObject](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/initWithIOSurface:startOffset:shouldRetain:
func (a ANEIOSurfaceObject) InitWithIOSurfaceStartOffsetShouldRetain(iOSurface coregraphics.IOSurfaceRef, offset objectivec.IObject, retain bool) ANEIOSurfaceObject {
	rv := objc.Send[ANEIOSurfaceObject](a.ID, objc.Sel("initWithIOSurface:startOffset:shouldRetain:"), iOSurface, offset, retain)
	return rv
}





//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/createIOSurfaceWithWidth:pixel_size:height:
func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) CreateIOSurfaceWithWidthPixel_sizeHeight(width int, pixel_size int, height int) coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("createIOSurfaceWithWidth:pixel_size:height:"), width, pixel_size, height)
	return coregraphics.IOSurfaceRef(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/createIOSurfaceWithWidth:pixel_size:height:bytesPerElement:
func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) CreateIOSurfaceWithWidthPixel_sizeHeightBytesPerElement(width int, pixel_size int, height int, element int) coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("createIOSurfaceWithWidth:pixel_size:height:bytesPerElement:"), width, pixel_size, height, element)
	return coregraphics.IOSurfaceRef(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/objectWithIOSurface:
func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) ObjectWithIOSurface(iOSurface coregraphics.IOSurfaceRef) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("objectWithIOSurface:"), iOSurface)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/objectWithIOSurface:startOffset:
func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) ObjectWithIOSurfaceStartOffset(iOSurface coregraphics.IOSurfaceRef, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("objectWithIOSurface:startOffset:"), iOSurface, offset)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/objectWithIOSurfaceNoRetain:startOffset:
func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) ObjectWithIOSurfaceNoRetainStartOffset(retain coregraphics.IOSurfaceRef, offset objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("objectWithIOSurfaceNoRetain:startOffset:"), retain, offset)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/supportsSecureCoding
func (_ANEIOSurfaceObjectClass ANEIOSurfaceObjectClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_ANEIOSurfaceObjectClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}








// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/ioSurface
func (a ANEIOSurfaceObject) IoSurface() coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](a.ID, objc.Sel("ioSurface"))
	return coregraphics.IOSurfaceRef(rv)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEIOSurfaceObject/startOffset
func (a ANEIOSurfaceObject) StartOffset() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("startOffset"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

















