// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRTrialTaskResult] class.
var (
	_MLRTrialTaskResultClass     MLRTrialTaskResultClass
	_MLRTrialTaskResultClassOnce sync.Once
)

func getMLRTrialTaskResultClass() MLRTrialTaskResultClass {
	_MLRTrialTaskResultClassOnce.Do(func() {
		_MLRTrialTaskResultClass = MLRTrialTaskResultClass{class: objc.GetClass("MLRTrialTaskResult")}
	})
	return _MLRTrialTaskResultClass
}

// GetMLRTrialTaskResultClass returns the class object for MLRTrialTaskResult.
func GetMLRTrialTaskResultClass() MLRTrialTaskResultClass {
	return getMLRTrialTaskResultClass()
}

type MLRTrialTaskResultClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRTrialTaskResultClass) Alloc() MLRTrialTaskResult {
	rv := objc.Send[MLRTrialTaskResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRTrialTaskResult.JSONResult]
//   - [MLRTrialTaskResult.SubmitForTaskError]
//   - [MLRTrialTaskResult.InitWithJSONResult]
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTaskResult
type MLRTrialTaskResult struct {
	objectivec.Object
}

// MLRTrialTaskResultFromID constructs a [MLRTrialTaskResult] from an objc.ID.
func MLRTrialTaskResultFromID(id objc.ID) MLRTrialTaskResult {
	return MLRTrialTaskResult{objectivec.Object{ID: id}}
}
// Ensure MLRTrialTaskResult implements IMLRTrialTaskResult.
var _ IMLRTrialTaskResult = MLRTrialTaskResult{}

// An interface definition for the [MLRTrialTaskResult] class.
//
// # Methods
//
//   - [IMLRTrialTaskResult.JSONResult]
//   - [IMLRTrialTaskResult.SubmitForTaskError]
//   - [IMLRTrialTaskResult.InitWithJSONResult]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTaskResult
type IMLRTrialTaskResult interface {
	objectivec.IObject

	// Topic: Methods

	JSONResult() foundation.INSDictionary
	SubmitForTaskError(task objectivec.IObject) (bool, error)
	InitWithJSONResult(jSONResult objectivec.IObject) MLRTrialTaskResult
}

// Init initializes the instance.
func (r MLRTrialTaskResult) Init() MLRTrialTaskResult {
	rv := objc.Send[MLRTrialTaskResult](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRTrialTaskResult) Autorelease() MLRTrialTaskResult {
	rv := objc.Send[MLRTrialTaskResult](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRTrialTaskResult creates a new MLRTrialTaskResult instance.
func NewMLRTrialTaskResult() MLRTrialTaskResult {
	class := getMLRTrialTaskResultClass()
	rv := objc.Send[MLRTrialTaskResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTaskResult/initWithJSONResult:
func NewRTrialTaskResultWithJSONResult(jSONResult objectivec.IObject) MLRTrialTaskResult {
	instance := getMLRTrialTaskResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithJSONResult:"), jSONResult)
	return MLRTrialTaskResultFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTaskResult/submitForTask:error:
func (r MLRTrialTaskResult) SubmitForTaskError(task objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](r.ID, objc.Sel("submitForTask:error:"), task, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("submitForTask:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTaskResult/initWithJSONResult:
func (r MLRTrialTaskResult) InitWithJSONResult(jSONResult objectivec.IObject) MLRTrialTaskResult {
	rv := objc.Send[MLRTrialTaskResult](r.ID, objc.Sel("initWithJSONResult:"), jSONResult)
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialTaskResult/JSONResult
func (r MLRTrialTaskResult) JSONResult() foundation.INSDictionary {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("JSONResult"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

