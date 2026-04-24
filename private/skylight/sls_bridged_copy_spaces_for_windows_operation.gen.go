// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedCopySpacesForWindowsOperation] class.
var (
	_SLSBridgedCopySpacesForWindowsOperationClass     SLSBridgedCopySpacesForWindowsOperationClass
	_SLSBridgedCopySpacesForWindowsOperationClassOnce sync.Once
)

func getSLSBridgedCopySpacesForWindowsOperationClass() SLSBridgedCopySpacesForWindowsOperationClass {
	_SLSBridgedCopySpacesForWindowsOperationClassOnce.Do(func() {
		_SLSBridgedCopySpacesForWindowsOperationClass = SLSBridgedCopySpacesForWindowsOperationClass{class: objc.GetClass("SLSBridgedCopySpacesForWindowsOperation")}
	})
	return _SLSBridgedCopySpacesForWindowsOperationClass
}

// GetSLSBridgedCopySpacesForWindowsOperationClass returns the class object for SLSBridgedCopySpacesForWindowsOperation.
func GetSLSBridgedCopySpacesForWindowsOperationClass() SLSBridgedCopySpacesForWindowsOperationClass {
	return getSLSBridgedCopySpacesForWindowsOperationClass()
}

type SLSBridgedCopySpacesForWindowsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedCopySpacesForWindowsOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedCopySpacesForWindowsOperationClass) Alloc() SLSBridgedCopySpacesForWindowsOperation {
	rv := objc.Send[SLSBridgedCopySpacesForWindowsOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopySpacesForWindowsOperation.MakeResultWithNumbers]
//   - [SLSBridgedCopySpacesForWindowsOperation.Options]
//   - [SLSBridgedCopySpacesForWindowsOperation.Windows]
//   - [SLSBridgedCopySpacesForWindowsOperation.InitWithOptionsWindows]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesForWindowsOperation
type SLSBridgedCopySpacesForWindowsOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedCopySpacesForWindowsOperationFromID constructs a [SLSBridgedCopySpacesForWindowsOperation] from an objc.ID.
func SLSBridgedCopySpacesForWindowsOperationFromID(id objc.ID) SLSBridgedCopySpacesForWindowsOperation {
	return SLSBridgedCopySpacesForWindowsOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedCopySpacesForWindowsOperation implements ISLSBridgedCopySpacesForWindowsOperation.
var _ ISLSBridgedCopySpacesForWindowsOperation = SLSBridgedCopySpacesForWindowsOperation{}

// An interface definition for the [SLSBridgedCopySpacesForWindowsOperation] class.
//
// # Methods
//
//   - [ISLSBridgedCopySpacesForWindowsOperation.MakeResultWithNumbers]
//   - [ISLSBridgedCopySpacesForWindowsOperation.Options]
//   - [ISLSBridgedCopySpacesForWindowsOperation.Windows]
//   - [ISLSBridgedCopySpacesForWindowsOperation.InitWithOptionsWindows]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesForWindowsOperation
type ISLSBridgedCopySpacesForWindowsOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject
	Options() uint32
	Windows() foundation.INSArray
	InitWithOptionsWindows(options uint32, windows objectivec.IObject) SLSBridgedCopySpacesForWindowsOperation
}

// Init initializes the instance.
func (s SLSBridgedCopySpacesForWindowsOperation) Init() SLSBridgedCopySpacesForWindowsOperation {
	rv := objc.Send[SLSBridgedCopySpacesForWindowsOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopySpacesForWindowsOperation) Autorelease() SLSBridgedCopySpacesForWindowsOperation {
	rv := objc.Send[SLSBridgedCopySpacesForWindowsOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopySpacesForWindowsOperation creates a new SLSBridgedCopySpacesForWindowsOperation instance.
func NewSLSBridgedCopySpacesForWindowsOperation() SLSBridgedCopySpacesForWindowsOperation {
	class := getSLSBridgedCopySpacesForWindowsOperationClass()
	rv := objc.Send[SLSBridgedCopySpacesForWindowsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesForWindowsOperation/initWithCoder:
func NewSLSBridgedCopySpacesForWindowsOperationWithCoder(coder objectivec.IObject) SLSBridgedCopySpacesForWindowsOperation {
	instance := getSLSBridgedCopySpacesForWindowsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopySpacesForWindowsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesForWindowsOperation/initWithOptions:windows:
func NewSLSBridgedCopySpacesForWindowsOperationWithOptionsWindows(options uint32, windows objectivec.IObject) SLSBridgedCopySpacesForWindowsOperation {
	instance := getSLSBridgedCopySpacesForWindowsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:windows:"), options, windows)
	return SLSBridgedCopySpacesForWindowsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesForWindowsOperation/makeResultWithNumbers:
func (s SLSBridgedCopySpacesForWindowsOperation) MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithNumbers:"), numbers)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesForWindowsOperation/initWithOptions:windows:
func (s SLSBridgedCopySpacesForWindowsOperation) InitWithOptionsWindows(options uint32, windows objectivec.IObject) SLSBridgedCopySpacesForWindowsOperation {
	rv := objc.Send[SLSBridgedCopySpacesForWindowsOperation](s.ID, objc.Sel("initWithOptions:windows:"), options, windows)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesForWindowsOperation/options
func (s SLSBridgedCopySpacesForWindowsOperation) Options() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("options"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesForWindowsOperation/windows
func (s SLSBridgedCopySpacesForWindowsOperation) Windows() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("windows"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
