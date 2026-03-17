// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEBuffer] class.
var (
	_ANEBufferClass     ANEBufferClass
	_ANEBufferClassOnce sync.Once
)

func getANEBufferClass() ANEBufferClass {
	_ANEBufferClassOnce.Do(func() {
		_ANEBufferClass = ANEBufferClass{class: objc.GetClass("_ANEBuffer")}
	})
	return _ANEBufferClass
}

// GetANEBufferClass returns the class object for _ANEBuffer.
func GetANEBufferClass() ANEBufferClass {
	return getANEBufferClass()
}

type ANEBufferClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEBufferClass) Alloc() ANEBuffer {
	rv := objc.Send[ANEBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEBuffer.EncodeWithCoder]
//   - [ANEBuffer.IoSurfaceObject]
//   - [ANEBuffer.Source]
//   - [ANEBuffer.SymbolIndex]
//   - [ANEBuffer.InitWithCoder]
//   - [ANEBuffer.InitWithIOSurfaceObjectSymbolIndexSource]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer
type ANEBuffer struct {
	objectivec.Object
}

// ANEBufferFromID constructs a [ANEBuffer] from an objc.ID.
func ANEBufferFromID(id objc.ID) ANEBuffer {
	return ANEBuffer{objectivec.Object{ID: id}}
}
// Ensure ANEBuffer implements IANEBuffer.
var _ IANEBuffer = ANEBuffer{}

// An interface definition for the [ANEBuffer] class.
//
// # Methods
//
//   - [IANEBuffer.EncodeWithCoder]
//   - [IANEBuffer.IoSurfaceObject]
//   - [IANEBuffer.Source]
//   - [IANEBuffer.SymbolIndex]
//   - [IANEBuffer.InitWithCoder]
//   - [IANEBuffer.InitWithIOSurfaceObjectSymbolIndexSource]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer
type IANEBuffer interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	IoSurfaceObject() *ANEIOSurfaceObject
	Source() int64
	SymbolIndex() foundation.NSNumber
	InitWithCoder(coder foundation.INSCoder) ANEBuffer
	InitWithIOSurfaceObjectSymbolIndexSource(object objectivec.IObject, index objectivec.IObject, source int64) ANEBuffer
}

// Init initializes the instance.
func (a ANEBuffer) Init() ANEBuffer {
	rv := objc.Send[ANEBuffer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEBuffer) Autorelease() ANEBuffer {
	rv := objc.Send[ANEBuffer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEBuffer creates a new ANEBuffer instance.
func NewANEBuffer() ANEBuffer {
	class := getANEBufferClass()
	rv := objc.Send[ANEBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer/initWithCoder:
func NewANEBufferWithCoder(coder objectivec.IObject) ANEBuffer {
	instance := getANEBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEBufferFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer/initWithIOSurfaceObject:symbolIndex:source:
func NewANEBufferWithIOSurfaceObjectSymbolIndexSource(object objectivec.IObject, index objectivec.IObject, source int64) ANEBuffer {
	instance := getANEBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOSurfaceObject:symbolIndex:source:"), object, index, source)
	return ANEBufferFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer/encodeWithCoder:
func (a ANEBuffer) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer/initWithCoder:
func (a ANEBuffer) InitWithCoder(coder foundation.INSCoder) ANEBuffer {
	rv := objc.Send[ANEBuffer](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer/initWithIOSurfaceObject:symbolIndex:source:
func (a ANEBuffer) InitWithIOSurfaceObjectSymbolIndexSource(object objectivec.IObject, index objectivec.IObject, source int64) ANEBuffer {
	rv := objc.Send[ANEBuffer](a.ID, objc.Sel("initWithIOSurfaceObject:symbolIndex:source:"), object, index, source)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer/bufferWithIOSurfaceObject:symbolIndex:source:
func (_ANEBufferClass ANEBufferClass) BufferWithIOSurfaceObjectSymbolIndexSource(object objectivec.IObject, index objectivec.IObject, source int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEBufferClass.class), objc.Sel("bufferWithIOSurfaceObject:symbolIndex:source:"), object, index, source)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer/supportsSecureCoding
func (_ANEBufferClass ANEBufferClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_ANEBufferClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer/ioSurfaceObject
func (a ANEBuffer) IoSurfaceObject() *ANEIOSurfaceObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("ioSurfaceObject"))
	if rv == 0 {
		return nil
	}
	val := ANEIOSurfaceObjectFromID(objc.ID(rv))
	return &val
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer/source
func (a ANEBuffer) Source() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("source"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEBuffer/symbolIndex
func (a ANEBuffer) SymbolIndex() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("symbolIndex"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

