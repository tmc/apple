// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationInt32Result] class.
var (
	_SLSBridgedWindowManagementOperationInt32ResultClass     SLSBridgedWindowManagementOperationInt32ResultClass
	_SLSBridgedWindowManagementOperationInt32ResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationInt32ResultClass() SLSBridgedWindowManagementOperationInt32ResultClass {
	_SLSBridgedWindowManagementOperationInt32ResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationInt32ResultClass = SLSBridgedWindowManagementOperationInt32ResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationInt32Result")}
	})
	return _SLSBridgedWindowManagementOperationInt32ResultClass
}

// GetSLSBridgedWindowManagementOperationInt32ResultClass returns the class object for SLSBridgedWindowManagementOperationInt32Result.
func GetSLSBridgedWindowManagementOperationInt32ResultClass() SLSBridgedWindowManagementOperationInt32ResultClass {
	return getSLSBridgedWindowManagementOperationInt32ResultClass()
}

type SLSBridgedWindowManagementOperationInt32ResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationInt32ResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationInt32ResultClass) Alloc() SLSBridgedWindowManagementOperationInt32Result {
	rv := objc.Send[SLSBridgedWindowManagementOperationInt32Result](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationInt32Result.Int32Value]
//   - [SLSBridgedWindowManagementOperationInt32Result.InitWithInt32Value]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationInt32Result
type SLSBridgedWindowManagementOperationInt32Result struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationInt32ResultFromID constructs a [SLSBridgedWindowManagementOperationInt32Result] from an objc.ID.
func SLSBridgedWindowManagementOperationInt32ResultFromID(id objc.ID) SLSBridgedWindowManagementOperationInt32Result {
	return SLSBridgedWindowManagementOperationInt32Result{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationInt32Result implements ISLSBridgedWindowManagementOperationInt32Result.
var _ ISLSBridgedWindowManagementOperationInt32Result = SLSBridgedWindowManagementOperationInt32Result{}

// An interface definition for the [SLSBridgedWindowManagementOperationInt32Result] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationInt32Result.Int32Value]
//   - [ISLSBridgedWindowManagementOperationInt32Result.InitWithInt32Value]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationInt32Result
type ISLSBridgedWindowManagementOperationInt32Result interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	Int32Value() int
	InitWithInt32Value(int32Value int) SLSBridgedWindowManagementOperationInt32Result
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationInt32Result) Init() SLSBridgedWindowManagementOperationInt32Result {
	rv := objc.Send[SLSBridgedWindowManagementOperationInt32Result](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationInt32Result) Autorelease() SLSBridgedWindowManagementOperationInt32Result {
	rv := objc.Send[SLSBridgedWindowManagementOperationInt32Result](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationInt32Result creates a new SLSBridgedWindowManagementOperationInt32Result instance.
func NewSLSBridgedWindowManagementOperationInt32Result() SLSBridgedWindowManagementOperationInt32Result {
	class := getSLSBridgedWindowManagementOperationInt32ResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationInt32Result](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationInt32Result/initWithCoder:
func NewSLSBridgedWindowManagementOperationInt32ResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationInt32Result {
	instance := getSLSBridgedWindowManagementOperationInt32ResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationInt32ResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationInt32Result/initWithInt32Value:
func NewSLSBridgedWindowManagementOperationInt32ResultWithInt32Value(int32Value int) SLSBridgedWindowManagementOperationInt32Result {
	instance := getSLSBridgedWindowManagementOperationInt32ResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInt32Value:"), int32Value)
	return SLSBridgedWindowManagementOperationInt32ResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationInt32Result/initWithInt32Value:
func (s SLSBridgedWindowManagementOperationInt32Result) InitWithInt32Value(int32Value int) SLSBridgedWindowManagementOperationInt32Result {
	rv := objc.Send[SLSBridgedWindowManagementOperationInt32Result](s.ID, objc.Sel("initWithInt32Value:"), int32Value)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationInt32Result/int32Value
func (s SLSBridgedWindowManagementOperationInt32Result) Int32Value() int {
	rv := objc.Send[int](s.ID, objc.Sel("int32Value"))
	return rv
}
