// Code generated from Apple documentation for ExecutionPolicy. DO NOT EDIT.

package executionpolicy

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EPExecutionPolicy] class.
var (
	_EPExecutionPolicyClass     EPExecutionPolicyClass
	_EPExecutionPolicyClassOnce sync.Once
)

func getEPExecutionPolicyClass() EPExecutionPolicyClass {
	_EPExecutionPolicyClassOnce.Do(func() {
		_EPExecutionPolicyClass = EPExecutionPolicyClass{class: objc.GetClass("EPExecutionPolicy")}
	})
	return _EPExecutionPolicyClass
}

// GetEPExecutionPolicyClass returns the class object for EPExecutionPolicy.
func GetEPExecutionPolicyClass() EPExecutionPolicyClass {
	return getEPExecutionPolicyClass()
}

type EPExecutionPolicyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EPExecutionPolicyClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EPExecutionPolicyClass) Alloc() EPExecutionPolicy {
	rv := objc.Send[EPExecutionPolicy](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Instance Methods
//
//   - [EPExecutionPolicy.AddPolicyExceptionForURLError]
//
// See: https://developer.apple.com/documentation/ExecutionPolicy/EPExecutionPolicy
type EPExecutionPolicy struct {
	objectivec.Object
}

// EPExecutionPolicyFromID constructs a [EPExecutionPolicy] from an objc.ID.
func EPExecutionPolicyFromID(id objc.ID) EPExecutionPolicy {
	return EPExecutionPolicy{objectivec.Object{ID: id}}
}

// NOTE: EPExecutionPolicy adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [EPExecutionPolicy] class.
//
// # Instance Methods
//
//   - [IEPExecutionPolicy.AddPolicyExceptionForURLError]
//
// See: https://developer.apple.com/documentation/ExecutionPolicy/EPExecutionPolicy
type IEPExecutionPolicy interface {
	objectivec.IObject

	// Topic: Instance Methods

	AddPolicyExceptionForURLError(url foundation.NSURL) (bool, error)
}

// Init initializes the instance.
func (e EPExecutionPolicy) Init() EPExecutionPolicy {
	rv := objc.Send[EPExecutionPolicy](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EPExecutionPolicy) Autorelease() EPExecutionPolicy {
	rv := objc.Send[EPExecutionPolicy](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEPExecutionPolicy creates a new EPExecutionPolicy instance.
func NewEPExecutionPolicy() EPExecutionPolicy {
	class := getEPExecutionPolicyClass()
	rv := objc.Send[EPExecutionPolicy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ExecutionPolicy/EPExecutionPolicy/addException(for:)
func (e EPExecutionPolicy) AddPolicyExceptionForURLError(url foundation.NSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("addPolicyExceptionForURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addPolicyExceptionForURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
