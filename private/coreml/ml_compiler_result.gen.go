// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCompilerResult] class.
var (
	_MLCompilerResultClass     MLCompilerResultClass
	_MLCompilerResultClassOnce sync.Once
)

func getMLCompilerResultClass() MLCompilerResultClass {
	_MLCompilerResultClassOnce.Do(func() {
		_MLCompilerResultClass = MLCompilerResultClass{class: objc.GetClass("MLCompilerResult")}
	})
	return _MLCompilerResultClass
}

// GetMLCompilerResultClass returns the class object for MLCompilerResult.
func GetMLCompilerResultClass() MLCompilerResultClass {
	return getMLCompilerResultClass()
}

type MLCompilerResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCompilerResultClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCompilerResultClass) Alloc() MLCompilerResult {
	rv := objc.Send[MLCompilerResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLCompilerResult.OutputFiles]
//   - [MLCompilerResult.SetOutputFiles]
// See: https://developer.apple.com/documentation/CoreML/MLCompilerResult
type MLCompilerResult struct {
	objectivec.Object
}

// MLCompilerResultFromID constructs a [MLCompilerResult] from an objc.ID.
func MLCompilerResultFromID(id objc.ID) MLCompilerResult {
	return MLCompilerResult{objectivec.Object{ID: id}}
}
// Ensure MLCompilerResult implements IMLCompilerResult.
var _ IMLCompilerResult = MLCompilerResult{}

// An interface definition for the [MLCompilerResult] class.
//
// # Methods
//
//   - [IMLCompilerResult.OutputFiles]
//   - [IMLCompilerResult.SetOutputFiles]
//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerResult
type IMLCompilerResult interface {
	objectivec.IObject

	// Topic: Methods

	OutputFiles() foundation.INSArray
	SetOutputFiles(value foundation.INSArray)
}

// Init initializes the instance.
func (c MLCompilerResult) Init() MLCompilerResult {
	rv := objc.Send[MLCompilerResult](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCompilerResult) Autorelease() MLCompilerResult {
	rv := objc.Send[MLCompilerResult](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCompilerResult creates a new MLCompilerResult instance.
func NewMLCompilerResult() MLCompilerResult {
	class := getMLCompilerResultClass()
	rv := objc.Send[MLCompilerResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerResult/resultWithArchive:
func (_MLCompilerResultClass MLCompilerResultClass) ResultWithArchive(archive unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerResultClass.class), objc.Sel("resultWithArchive:"), archive)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerResult/resultWithOutputFiles:
func (_MLCompilerResultClass MLCompilerResultClass) ResultWithOutputFiles(files objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerResultClass.class), objc.Sel("resultWithOutputFiles:"), files)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerResult/outputFiles
func (c MLCompilerResult) OutputFiles() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputFiles"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (c MLCompilerResult) SetOutputFiles(value foundation.INSArray) {
	objc.Send[struct{}](c.ID, objc.Sel("setOutputFiles:"), value)
}

