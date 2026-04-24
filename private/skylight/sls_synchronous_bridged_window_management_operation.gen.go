// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSSynchronousBridgedWindowManagementOperation] class.
var (
	_SLSSynchronousBridgedWindowManagementOperationClass     SLSSynchronousBridgedWindowManagementOperationClass
	_SLSSynchronousBridgedWindowManagementOperationClassOnce sync.Once
)

func getSLSSynchronousBridgedWindowManagementOperationClass() SLSSynchronousBridgedWindowManagementOperationClass {
	_SLSSynchronousBridgedWindowManagementOperationClassOnce.Do(func() {
		_SLSSynchronousBridgedWindowManagementOperationClass = SLSSynchronousBridgedWindowManagementOperationClass{class: objc.GetClass("SLSSynchronousBridgedWindowManagementOperation")}
	})
	return _SLSSynchronousBridgedWindowManagementOperationClass
}

// GetSLSSynchronousBridgedWindowManagementOperationClass returns the class object for SLSSynchronousBridgedWindowManagementOperation.
func GetSLSSynchronousBridgedWindowManagementOperationClass() SLSSynchronousBridgedWindowManagementOperationClass {
	return getSLSSynchronousBridgedWindowManagementOperationClass()
}

type SLSSynchronousBridgedWindowManagementOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSSynchronousBridgedWindowManagementOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSSynchronousBridgedWindowManagementOperationClass) Alloc() SLSSynchronousBridgedWindowManagementOperation {
	rv := objc.Send[SLSSynchronousBridgedWindowManagementOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSSynchronousBridgedWindowManagementOperation._init]
//   - [SLSSynchronousBridgedWindowManagementOperation.EncodeWithCoder]
//   - [SLSSynchronousBridgedWindowManagementOperation.InvokeFallback]
//   - [SLSSynchronousBridgedWindowManagementOperation.PerformWithWMBridgeDelegate]
//   - [SLSSynchronousBridgedWindowManagementOperation.InitWithCoder]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSynchronousBridgedWindowManagementOperation
type SLSSynchronousBridgedWindowManagementOperation struct {
	objectivec.Object
}

// SLSSynchronousBridgedWindowManagementOperationFromID constructs a [SLSSynchronousBridgedWindowManagementOperation] from an objc.ID.
func SLSSynchronousBridgedWindowManagementOperationFromID(id objc.ID) SLSSynchronousBridgedWindowManagementOperation {
	return SLSSynchronousBridgedWindowManagementOperation{objectivec.Object{ID: id}}
}

// Ensure SLSSynchronousBridgedWindowManagementOperation implements ISLSSynchronousBridgedWindowManagementOperation.
var _ ISLSSynchronousBridgedWindowManagementOperation = SLSSynchronousBridgedWindowManagementOperation{}

// An interface definition for the [SLSSynchronousBridgedWindowManagementOperation] class.
//
// # Methods
//
//   - [ISLSSynchronousBridgedWindowManagementOperation._init]
//   - [ISLSSynchronousBridgedWindowManagementOperation.EncodeWithCoder]
//   - [ISLSSynchronousBridgedWindowManagementOperation.InvokeFallback]
//   - [ISLSSynchronousBridgedWindowManagementOperation.PerformWithWMBridgeDelegate]
//   - [ISLSSynchronousBridgedWindowManagementOperation.InitWithCoder]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSynchronousBridgedWindowManagementOperation
type ISLSSynchronousBridgedWindowManagementOperation interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	InvokeFallback() objectivec.IObject
	PerformWithWMBridgeDelegate() objectivec.IObject
	InitWithCoder(coder foundation.INSCoder) SLSSynchronousBridgedWindowManagementOperation
}

// Init initializes the instance.
func (s SLSSynchronousBridgedWindowManagementOperation) Init() SLSSynchronousBridgedWindowManagementOperation {
	rv := objc.Send[SLSSynchronousBridgedWindowManagementOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSSynchronousBridgedWindowManagementOperation) Autorelease() SLSSynchronousBridgedWindowManagementOperation {
	rv := objc.Send[SLSSynchronousBridgedWindowManagementOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSSynchronousBridgedWindowManagementOperation creates a new SLSSynchronousBridgedWindowManagementOperation instance.
func NewSLSSynchronousBridgedWindowManagementOperation() SLSSynchronousBridgedWindowManagementOperation {
	class := getSLSSynchronousBridgedWindowManagementOperationClass()
	rv := objc.Send[SLSSynchronousBridgedWindowManagementOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSynchronousBridgedWindowManagementOperation/initWithCoder:
func NewSLSSynchronousBridgedWindowManagementOperationWithCoder(coder objectivec.IObject) SLSSynchronousBridgedWindowManagementOperation {
	instance := getSLSSynchronousBridgedWindowManagementOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSSynchronousBridgedWindowManagementOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSynchronousBridgedWindowManagementOperation/_init
func (s SLSSynchronousBridgedWindowManagementOperation) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSynchronousBridgedWindowManagementOperation/encodeWithCoder:
func (s SLSSynchronousBridgedWindowManagementOperation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSynchronousBridgedWindowManagementOperation/invokeFallback
func (s SLSSynchronousBridgedWindowManagementOperation) InvokeFallback() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("invokeFallback"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSynchronousBridgedWindowManagementOperation/performWithWMBridgeDelegate
func (s SLSSynchronousBridgedWindowManagementOperation) PerformWithWMBridgeDelegate() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("performWithWMBridgeDelegate"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSynchronousBridgedWindowManagementOperation/initWithCoder:
func (s SLSSynchronousBridgedWindowManagementOperation) InitWithCoder(coder foundation.INSCoder) SLSSynchronousBridgedWindowManagementOperation {
	rv := objc.Send[SLSSynchronousBridgedWindowManagementOperation](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSynchronousBridgedWindowManagementOperation/supportsSecureCoding
func (_SLSSynchronousBridgedWindowManagementOperationClass SLSSynchronousBridgedWindowManagementOperationClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_SLSSynchronousBridgedWindowManagementOperationClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
