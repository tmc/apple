// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIRenderTask] class.
var (
	_CIRenderTaskClass     CIRenderTaskClass
	_CIRenderTaskClassOnce sync.Once
)

func getCIRenderTaskClass() CIRenderTaskClass {
	_CIRenderTaskClassOnce.Do(func() {
		_CIRenderTaskClass = CIRenderTaskClass{class: objc.GetClass("CIRenderTask")}
	})
	return _CIRenderTaskClass
}

// GetCIRenderTaskClass returns the class object for CIRenderTask.
func GetCIRenderTaskClass() CIRenderTaskClass {
	return getCIRenderTaskClass()
}

type CIRenderTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIRenderTaskClass) Alloc() CIRenderTask {
	rv := objc.Send[CIRenderTask](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A single render task.
//
// # Overview
// 
// A single render task issued in conjunction with [CIRenderDestination].
// 
// A [CIRenderTask] object appears in Xcode Quick Look as a graph.
//
// # Instance Methods
//
//   - [CIRenderTask.WaitUntilCompletedAndReturnError]: Waits until the [CIRenderTask](<doc://com.apple.coreimage/documentation/CoreImage/CIRenderTask>) finishes and returns.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderTask
type CIRenderTask struct {
	objectivec.Object
}

// CIRenderTaskFromID constructs a [CIRenderTask] from an objc.ID.
//
// A single render task.
func CIRenderTaskFromID(id objc.ID) CIRenderTask {
	return CIRenderTask{objectivec.Object{ID: id}}
}
// NOTE: CIRenderTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIRenderTask] class.
//
// # Instance Methods
//
//   - [ICIRenderTask.WaitUntilCompletedAndReturnError]: Waits until the [CIRenderTask](<doc://com.apple.coreimage/documentation/CoreImage/CIRenderTask>) finishes and returns.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderTask
type ICIRenderTask interface {
	objectivec.IObject

	// Topic: Instance Methods

	// Waits until the [CIRenderTask](<doc://com.apple.coreimage/documentation/CoreImage/CIRenderTask>) finishes and returns.
	WaitUntilCompletedAndReturnError() (ICIRenderInfo, error)
}

// Init initializes the instance.
func (r CIRenderTask) Init() CIRenderTask {
	rv := objc.Send[CIRenderTask](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r CIRenderTask) Autorelease() CIRenderTask {
	rv := objc.Send[CIRenderTask](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIRenderTask creates a new CIRenderTask instance.
func NewCIRenderTask() CIRenderTask {
	class := getCIRenderTaskClass()
	rv := objc.Send[CIRenderTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Waits until the [CIRenderTask] finishes and returns.
//
// # Discussion
// 
// Synchronously blocks execution until the [CIRenderTask] either completes or
// fails (with error). Calling this method after
// [StartTaskToRenderToDestinationError] or
// [StartTaskToRenderFromRectToDestinationAtPointError] makes the render task
// behave synchronously, as if the CPU and GPU were operating as a single
// unit.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRenderTask/waitUntilCompleted()
func (r CIRenderTask) WaitUntilCompletedAndReturnError() (ICIRenderInfo, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](r.ID, objc.Sel("waitUntilCompletedAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return CIRenderInfo{}, foundation.NSErrorFrom(errorPtr)
	}
	return CIRenderInfoFromID(rv), nil

}

