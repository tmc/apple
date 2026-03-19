// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRTaskResult] class.
var (
	_MLRTaskResultClass     MLRTaskResultClass
	_MLRTaskResultClassOnce sync.Once
)

func getMLRTaskResultClass() MLRTaskResultClass {
	_MLRTaskResultClassOnce.Do(func() {
		_MLRTaskResultClass = MLRTaskResultClass{class: objc.GetClass("MLRTaskResult")}
	})
	return _MLRTaskResultClass
}

// GetMLRTaskResultClass returns the class object for MLRTaskResult.
func GetMLRTaskResultClass() MLRTaskResultClass {
	return getMLRTaskResultClass()
}

type MLRTaskResultClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRTaskResultClass) Alloc() MLRTaskResult {
	rv := objc.Send[MLRTaskResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRTaskResult.JSONResult]
//   - [MLRTaskResult.EncodeWithCoder]
//   - [MLRTaskResult.Vector]
//   - [MLRTaskResult.InitWithCoder]
//   - [MLRTaskResult.InitWithJSONResultUnprivatizedVector]
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskResult
type MLRTaskResult struct {
	objectivec.Object
}

// MLRTaskResultFromID constructs a [MLRTaskResult] from an objc.ID.
func MLRTaskResultFromID(id objc.ID) MLRTaskResult {
	return MLRTaskResult{objectivec.Object{ID: id}}
}
// Ensure MLRTaskResult implements IMLRTaskResult.
var _ IMLRTaskResult = MLRTaskResult{}

// An interface definition for the [MLRTaskResult] class.
//
// # Methods
//
//   - [IMLRTaskResult.JSONResult]
//   - [IMLRTaskResult.EncodeWithCoder]
//   - [IMLRTaskResult.Vector]
//   - [IMLRTaskResult.InitWithCoder]
//   - [IMLRTaskResult.InitWithJSONResultUnprivatizedVector]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskResult
type IMLRTaskResult interface {
	objectivec.IObject

	// Topic: Methods

	JSONResult() foundation.INSDictionary
	EncodeWithCoder(coder foundation.INSCoder)
	Vector() foundation.INSData
	InitWithCoder(coder foundation.INSCoder) MLRTaskResult
	InitWithJSONResultUnprivatizedVector(jSONResult objectivec.IObject, vector objectivec.IObject) MLRTaskResult
}

// Init initializes the instance.
func (r MLRTaskResult) Init() MLRTaskResult {
	rv := objc.Send[MLRTaskResult](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRTaskResult) Autorelease() MLRTaskResult {
	rv := objc.Send[MLRTaskResult](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRTaskResult creates a new MLRTaskResult instance.
func NewMLRTaskResult() MLRTaskResult {
	class := getMLRTaskResultClass()
	rv := objc.Send[MLRTaskResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskResult/initWithCoder:
func NewRTaskResultWithCoder(coder objectivec.IObject) MLRTaskResult {
	instance := getMLRTaskResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLRTaskResultFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskResult/initWithJSONResult:unprivatizedVector:
func NewRTaskResultWithJSONResultUnprivatizedVector(jSONResult objectivec.IObject, vector objectivec.IObject) MLRTaskResult {
	instance := getMLRTaskResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithJSONResult:unprivatizedVector:"), jSONResult, vector)
	return MLRTaskResultFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskResult/encodeWithCoder:
func (r MLRTaskResult) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskResult/initWithCoder:
func (r MLRTaskResult) InitWithCoder(coder foundation.INSCoder) MLRTaskResult {
	rv := objc.Send[MLRTaskResult](r.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskResult/initWithJSONResult:unprivatizedVector:
func (r MLRTaskResult) InitWithJSONResultUnprivatizedVector(jSONResult objectivec.IObject, vector objectivec.IObject) MLRTaskResult {
	rv := objc.Send[MLRTaskResult](r.ID, objc.Sel("initWithJSONResult:unprivatizedVector:"), jSONResult, vector)
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskResult/supportsSecureCoding
func (_MLRTaskResultClass MLRTaskResultClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLRTaskResultClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskResult/JSONResult
func (r MLRTaskResult) JSONResult() foundation.INSDictionary {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("JSONResult"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskResult/vector
func (r MLRTaskResult) Vector() foundation.INSData {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("vector"))
	return foundation.NSDataFromID(objc.ID(rv))
}

