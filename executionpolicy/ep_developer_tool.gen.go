// Code generated from Apple documentation for ExecutionPolicy. DO NOT EDIT.

package executionpolicy

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EPDeveloperTool] class.
var (
	_EPDeveloperToolClass     EPDeveloperToolClass
	_EPDeveloperToolClassOnce sync.Once
)

func getEPDeveloperToolClass() EPDeveloperToolClass {
	_EPDeveloperToolClassOnce.Do(func() {
		_EPDeveloperToolClass = EPDeveloperToolClass{class: objc.GetClass("EPDeveloperTool")}
	})
	return _EPDeveloperToolClass
}

// GetEPDeveloperToolClass returns the class object for EPDeveloperTool.
func GetEPDeveloperToolClass() EPDeveloperToolClass {
	return getEPDeveloperToolClass()
}

type EPDeveloperToolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EPDeveloperToolClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EPDeveloperToolClass) Alloc() EPDeveloperTool {
	rv := objc.Send[EPDeveloperTool](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [EPDeveloperTool.AuthorizationStatus]
//
// # Instance Methods
//
//   - [EPDeveloperTool.RequestDeveloperToolAccessWithCompletionHandler]
//
// See: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperTool
type EPDeveloperTool struct {
	objectivec.Object
}

// EPDeveloperToolFromID constructs a [EPDeveloperTool] from an objc.ID.
func EPDeveloperToolFromID(id objc.ID) EPDeveloperTool {
	return EPDeveloperTool{objectivec.Object{ID: id}}
}

// NOTE: EPDeveloperTool adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [EPDeveloperTool] class.
//
// # Instance Properties
//
//   - [IEPDeveloperTool.AuthorizationStatus]
//
// # Instance Methods
//
//   - [IEPDeveloperTool.RequestDeveloperToolAccessWithCompletionHandler]
//
// See: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperTool
type IEPDeveloperTool interface {
	objectivec.IObject

	// Topic: Instance Properties

	AuthorizationStatus() EPDeveloperToolStatus

	// Topic: Instance Methods

	RequestDeveloperToolAccessWithCompletionHandler(handler BoolHandler)
}

// Init initializes the instance.
func (e EPDeveloperTool) Init() EPDeveloperTool {
	rv := objc.Send[EPDeveloperTool](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EPDeveloperTool) Autorelease() EPDeveloperTool {
	rv := objc.Send[EPDeveloperTool](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEPDeveloperTool creates a new EPDeveloperTool instance.
func NewEPDeveloperTool() EPDeveloperTool {
	class := getEPDeveloperToolClass()
	rv := objc.Send[EPDeveloperTool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// # Discussion
//
// See: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperTool/requestAccess(completionHandler:)
func (e EPDeveloperTool) RequestDeveloperToolAccessWithCompletionHandler(handler BoolHandler) {
	_block0, _ := NewBoolBlock(handler)
	objc.Send[objc.ID](e.ID, objc.Sel("requestDeveloperToolAccessWithCompletionHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperTool/authorizationStatus
func (e EPDeveloperTool) AuthorizationStatus() EPDeveloperToolStatus {
	rv := objc.Send[EPDeveloperToolStatus](e.ID, objc.Sel("authorizationStatus"))
	return EPDeveloperToolStatus(rv)
}

// RequestDeveloperToolAccess is a synchronous wrapper around [EPDeveloperTool.RequestDeveloperToolAccessWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (e EPDeveloperTool) RequestDeveloperToolAccess(ctx context.Context) (bool, error) {
	done := make(chan bool, 1)
	e.RequestDeveloperToolAccessWithCompletionHandler(func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}
