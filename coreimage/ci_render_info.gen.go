// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIRenderInfo] class.
var (
	_CIRenderInfoClass     CIRenderInfoClass
	_CIRenderInfoClassOnce sync.Once
)

func getCIRenderInfoClass() CIRenderInfoClass {
	_CIRenderInfoClassOnce.Do(func() {
		_CIRenderInfoClass = CIRenderInfoClass{class: objc.GetClass("CIRenderInfo")}
	})
	return _CIRenderInfoClass
}

// GetCIRenderInfoClass returns the class object for CIRenderInfo.
func GetCIRenderInfoClass() CIRenderInfoClass {
	return getCIRenderInfoClass()
}

type CIRenderInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIRenderInfoClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIRenderInfoClass) Alloc() CIRenderInfo {
	rv := objc.Send[CIRenderInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An encapsulation of a render task’s timing, passes, and pixels processed.
//
// # Overview
//
// A [CIRenderInfo] object allows Xcode Quick Look to visualize the render
// graph with detailed timing information.
//
// # Instance Properties
//
//   - [CIRenderInfo.KernelExecutionTime]: The amount of time a render spent executing kernels.
//   - [CIRenderInfo.PassCount]: The number of passes the render took.
//   - [CIRenderInfo.PixelsProcessed]: The number of pixels the render produced executing kernels.
//   - [CIRenderInfo.KernelCompileTime]
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderInfo
type CIRenderInfo struct {
	objectivec.Object
}

// CIRenderInfoFromID constructs a [CIRenderInfo] from an objc.ID.
//
// An encapsulation of a render task’s timing, passes, and pixels processed.
func CIRenderInfoFromID(id objc.ID) CIRenderInfo {
	return CIRenderInfo{objectivec.Object{ID: id}}
}

// NOTE: CIRenderInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIRenderInfo] class.
//
// # Instance Properties
//
//   - [ICIRenderInfo.KernelExecutionTime]: The amount of time a render spent executing kernels.
//   - [ICIRenderInfo.PassCount]: The number of passes the render took.
//   - [ICIRenderInfo.PixelsProcessed]: The number of pixels the render produced executing kernels.
//   - [ICIRenderInfo.KernelCompileTime]
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderInfo
type ICIRenderInfo interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The amount of time a render spent executing kernels.
	KernelExecutionTime() float64
	// The number of passes the render took.
	PassCount() int
	// The number of pixels the render produced executing kernels.
	PixelsProcessed() int
	KernelCompileTime() float64
}

// Init initializes the instance.
func (r CIRenderInfo) Init() CIRenderInfo {
	rv := objc.Send[CIRenderInfo](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r CIRenderInfo) Autorelease() CIRenderInfo {
	rv := objc.Send[CIRenderInfo](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIRenderInfo creates a new CIRenderInfo instance.
func NewCIRenderInfo() CIRenderInfo {
	class := getCIRenderInfoClass()
	rv := objc.Send[CIRenderInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The amount of time a render spent executing kernels.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderInfo/kernelExecutionTime
func (r CIRenderInfo) KernelExecutionTime() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("kernelExecutionTime"))
	return rv
}

// The number of passes the render took.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderInfo/passCount
func (r CIRenderInfo) PassCount() int {
	rv := objc.Send[int](r.ID, objc.Sel("passCount"))
	return rv
}

// The number of pixels the render produced executing kernels.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderInfo/pixelsProcessed
func (r CIRenderInfo) PixelsProcessed() int {
	rv := objc.Send[int](r.ID, objc.Sel("pixelsProcessed"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIRenderInfo/kernelCompileTime
func (r CIRenderInfo) KernelCompileTime() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("kernelCompileTime"))
	return rv
}
