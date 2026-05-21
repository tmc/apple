// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSAsynchronousBridgedWindowManagementOperation] class.
var (
	_SLSAsynchronousBridgedWindowManagementOperationClass     SLSAsynchronousBridgedWindowManagementOperationClass
	_SLSAsynchronousBridgedWindowManagementOperationClassOnce sync.Once
)

func getSLSAsynchronousBridgedWindowManagementOperationClass() SLSAsynchronousBridgedWindowManagementOperationClass {
	_SLSAsynchronousBridgedWindowManagementOperationClassOnce.Do(func() {
		_SLSAsynchronousBridgedWindowManagementOperationClass = SLSAsynchronousBridgedWindowManagementOperationClass{class: objc.GetClass("SLSAsynchronousBridgedWindowManagementOperation")}
	})
	return _SLSAsynchronousBridgedWindowManagementOperationClass
}

// GetSLSAsynchronousBridgedWindowManagementOperationClass returns the class object for SLSAsynchronousBridgedWindowManagementOperation.
func GetSLSAsynchronousBridgedWindowManagementOperationClass() SLSAsynchronousBridgedWindowManagementOperationClass {
	return getSLSAsynchronousBridgedWindowManagementOperationClass()
}

type SLSAsynchronousBridgedWindowManagementOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSAsynchronousBridgedWindowManagementOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSAsynchronousBridgedWindowManagementOperationClass) Alloc() SLSAsynchronousBridgedWindowManagementOperation {
	rv := objc.Send[SLSAsynchronousBridgedWindowManagementOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSAsynchronousBridgedWindowManagementOperation._init]
//   - [SLSAsynchronousBridgedWindowManagementOperation.EncodeWithCoder]
//   - [SLSAsynchronousBridgedWindowManagementOperation.InvokeFallback]
//   - [SLSAsynchronousBridgedWindowManagementOperation.PerformWithWMBridgeDelegate]
//   - [SLSAsynchronousBridgedWindowManagementOperation.InitWithCoder]
type SLSAsynchronousBridgedWindowManagementOperation struct {
	objectivec.Object
}

// SLSAsynchronousBridgedWindowManagementOperationFromID constructs a [SLSAsynchronousBridgedWindowManagementOperation] from an objc.ID.
func SLSAsynchronousBridgedWindowManagementOperationFromID(id objc.ID) SLSAsynchronousBridgedWindowManagementOperation {
	return SLSAsynchronousBridgedWindowManagementOperation{objectivec.Object{ID: id}}
}

// Ensure SLSAsynchronousBridgedWindowManagementOperation implements ISLSAsynchronousBridgedWindowManagementOperation.
var _ ISLSAsynchronousBridgedWindowManagementOperation = SLSAsynchronousBridgedWindowManagementOperation{}

// An interface definition for the [SLSAsynchronousBridgedWindowManagementOperation] class.
//
// # Methods
//
//   - [ISLSAsynchronousBridgedWindowManagementOperation._init]
//   - [ISLSAsynchronousBridgedWindowManagementOperation.EncodeWithCoder]
//   - [ISLSAsynchronousBridgedWindowManagementOperation.InvokeFallback]
//   - [ISLSAsynchronousBridgedWindowManagementOperation.PerformWithWMBridgeDelegate]
//   - [ISLSAsynchronousBridgedWindowManagementOperation.InitWithCoder]
type ISLSAsynchronousBridgedWindowManagementOperation interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	InvokeFallback()
	PerformWithWMBridgeDelegate()
	InitWithCoder(coder foundation.INSCoder) SLSAsynchronousBridgedWindowManagementOperation
}

// Init initializes the instance.
func (s SLSAsynchronousBridgedWindowManagementOperation) Init() SLSAsynchronousBridgedWindowManagementOperation {
	rv := objc.Send[SLSAsynchronousBridgedWindowManagementOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSAsynchronousBridgedWindowManagementOperation) Autorelease() SLSAsynchronousBridgedWindowManagementOperation {
	rv := objc.Send[SLSAsynchronousBridgedWindowManagementOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSAsynchronousBridgedWindowManagementOperation creates a new SLSAsynchronousBridgedWindowManagementOperation instance.
func NewSLSAsynchronousBridgedWindowManagementOperation() SLSAsynchronousBridgedWindowManagementOperation {
	class := getSLSAsynchronousBridgedWindowManagementOperationClass()
	rv := objc.Send[SLSAsynchronousBridgedWindowManagementOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSAsynchronousBridgedWindowManagementOperationWithCoder(coder objectivec.IObject) SLSAsynchronousBridgedWindowManagementOperation {
	instance := getSLSAsynchronousBridgedWindowManagementOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSAsynchronousBridgedWindowManagementOperationFromID(rv)
}

func (s SLSAsynchronousBridgedWindowManagementOperation) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (s SLSAsynchronousBridgedWindowManagementOperation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (s SLSAsynchronousBridgedWindowManagementOperation) InvokeFallback() {
	objc.Send[objc.ID](s.ID, objc.Sel("invokeFallback"))
}
func (s SLSAsynchronousBridgedWindowManagementOperation) PerformWithWMBridgeDelegate() {
	objc.Send[objc.ID](s.ID, objc.Sel("performWithWMBridgeDelegate"))
}
func (s SLSAsynchronousBridgedWindowManagementOperation) InitWithCoder(coder foundation.INSCoder) SLSAsynchronousBridgedWindowManagementOperation {
	rv := objc.Send[SLSAsynchronousBridgedWindowManagementOperation](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_SLSAsynchronousBridgedWindowManagementOperationClass SLSAsynchronousBridgedWindowManagementOperationClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_SLSAsynchronousBridgedWindowManagementOperationClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
