// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETImagePreprocessor] class.
var (
	_ETImagePreprocessorClass     ETImagePreprocessorClass
	_ETImagePreprocessorClassOnce sync.Once
)

func getETImagePreprocessorClass() ETImagePreprocessorClass {
	_ETImagePreprocessorClassOnce.Do(func() {
		_ETImagePreprocessorClass = ETImagePreprocessorClass{class: objc.GetClass("ETImagePreprocessor")}
	})
	return _ETImagePreprocessorClass
}

// GetETImagePreprocessorClass returns the class object for ETImagePreprocessor.
func GetETImagePreprocessorClass() ETImagePreprocessorClass {
	return getETImagePreprocessorClass()
}

type ETImagePreprocessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETImagePreprocessorClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETImagePreprocessorClass) Alloc() ETImagePreprocessor {
	rv := objc.Send[ETImagePreprocessor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETImagePreprocessor.LoadSrcBufferWithCGImage]
//   - [ETImagePreprocessor.Preprocess]
//   - [ETImagePreprocessor.TensorWithCGImage]
//   - [ETImagePreprocessor.TensorWithPath]
//   - [ETImagePreprocessor.InitWithImagePreprocessParams]
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessor
type ETImagePreprocessor struct {
	objectivec.Object
}

// ETImagePreprocessorFromID constructs a [ETImagePreprocessor] from an objc.ID.
func ETImagePreprocessorFromID(id objc.ID) ETImagePreprocessor {
	return ETImagePreprocessor{objectivec.Object{ID: id}}
}
// Ensure ETImagePreprocessor implements IETImagePreprocessor.
var _ IETImagePreprocessor = ETImagePreprocessor{}

// An interface definition for the [ETImagePreprocessor] class.
//
// # Methods
//
//   - [IETImagePreprocessor.LoadSrcBufferWithCGImage]
//   - [IETImagePreprocessor.Preprocess]
//   - [IETImagePreprocessor.TensorWithCGImage]
//   - [IETImagePreprocessor.TensorWithPath]
//   - [IETImagePreprocessor.InitWithImagePreprocessParams]
//
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessor
type IETImagePreprocessor interface {
	objectivec.IObject

	// Topic: Methods

	LoadSrcBufferWithCGImage(cGImage *coregraphics.CGImageRef)
	Preprocess()
	TensorWithCGImage(cGImage *coregraphics.CGImageRef) objectivec.IObject
	TensorWithPath(path objectivec.IObject) objectivec.IObject
	InitWithImagePreprocessParams(params objectivec.IObject) ETImagePreprocessor
}

// Init initializes the instance.
func (e ETImagePreprocessor) Init() ETImagePreprocessor {
	rv := objc.Send[ETImagePreprocessor](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETImagePreprocessor) Autorelease() ETImagePreprocessor {
	rv := objc.Send[ETImagePreprocessor](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETImagePreprocessor creates a new ETImagePreprocessor instance.
func NewETImagePreprocessor() ETImagePreprocessor {
	class := getETImagePreprocessorClass()
	rv := objc.Send[ETImagePreprocessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessor/initWithImagePreprocessParams:
func NewETImagePreprocessorWithImagePreprocessParams(params objectivec.IObject) ETImagePreprocessor {
	instance := getETImagePreprocessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImagePreprocessParams:"), params)
	return ETImagePreprocessorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessor/loadSrcBufferWithCGImage:
func (e ETImagePreprocessor) LoadSrcBufferWithCGImage(cGImage *coregraphics.CGImageRef) {
	objc.Send[objc.ID](e.ID, objc.Sel("loadSrcBufferWithCGImage:"), cGImage)
}
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessor/preprocess
func (e ETImagePreprocessor) Preprocess() {
	objc.Send[objc.ID](e.ID, objc.Sel("preprocess"))
}
//
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessor/tensorWithCGImage:
func (e ETImagePreprocessor) TensorWithCGImage(cGImage *coregraphics.CGImageRef) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("tensorWithCGImage:"), cGImage)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessor/tensorWithPath:
func (e ETImagePreprocessor) TensorWithPath(path objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("tensorWithPath:"), path)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessor/initWithImagePreprocessParams:
func (e ETImagePreprocessor) InitWithImagePreprocessParams(params objectivec.IObject) ETImagePreprocessor {
	rv := objc.Send[ETImagePreprocessor](e.ID, objc.Sel("initWithImagePreprocessParams:"), params)
	return rv
}

