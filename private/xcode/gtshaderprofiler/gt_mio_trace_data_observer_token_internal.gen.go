// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioTraceDataObserverTokenInternal] class.
var (
	_GTMioTraceDataObserverTokenInternalClass     GTMioTraceDataObserverTokenInternalClass
	_GTMioTraceDataObserverTokenInternalClassOnce sync.Once
)

func getGTMioTraceDataObserverTokenInternalClass() GTMioTraceDataObserverTokenInternalClass {
	_GTMioTraceDataObserverTokenInternalClassOnce.Do(func() {
		_GTMioTraceDataObserverTokenInternalClass = GTMioTraceDataObserverTokenInternalClass{class: objc.GetClass("GTMioTraceDataObserverTokenInternal")}
	})
	return _GTMioTraceDataObserverTokenInternalClass
}

// GetGTMioTraceDataObserverTokenInternalClass returns the class object for GTMioTraceDataObserverTokenInternal.
func GetGTMioTraceDataObserverTokenInternalClass() GTMioTraceDataObserverTokenInternalClass {
	return getGTMioTraceDataObserverTokenInternalClass()
}

type GTMioTraceDataObserverTokenInternalClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceDataObserverTokenInternalClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceDataObserverTokenInternalClass) Alloc() GTMioTraceDataObserverTokenInternal {
	rv := objc.Send[GTMioTraceDataObserverTokenInternal](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioTraceDataObserverTokenInternal.Cancel]
//   - [GTMioTraceDataObserverTokenInternal.InitWithTraceDataIdentifier]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataObserverTokenInternal
type GTMioTraceDataObserverTokenInternal struct {
	objectivec.Object
}

// GTMioTraceDataObserverTokenInternalFromID constructs a [GTMioTraceDataObserverTokenInternal] from an objc.ID.
func GTMioTraceDataObserverTokenInternalFromID(id objc.ID) GTMioTraceDataObserverTokenInternal {
	return GTMioTraceDataObserverTokenInternal{objectivec.Object{ID: id}}
}
// Ensure GTMioTraceDataObserverTokenInternal implements IGTMioTraceDataObserverTokenInternal.
var _ IGTMioTraceDataObserverTokenInternal = GTMioTraceDataObserverTokenInternal{}

// An interface definition for the [GTMioTraceDataObserverTokenInternal] class.
//
// # Methods
//
//   - [IGTMioTraceDataObserverTokenInternal.Cancel]
//   - [IGTMioTraceDataObserverTokenInternal.InitWithTraceDataIdentifier]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataObserverTokenInternal
type IGTMioTraceDataObserverTokenInternal interface {
	objectivec.IObject

	// Topic: Methods

	Cancel()
	InitWithTraceDataIdentifier(data objectivec.IObject, identifier uint64) GTMioTraceDataObserverTokenInternal
}

// Init initializes the instance.
func (g GTMioTraceDataObserverTokenInternal) Init() GTMioTraceDataObserverTokenInternal {
	rv := objc.Send[GTMioTraceDataObserverTokenInternal](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceDataObserverTokenInternal) Autorelease() GTMioTraceDataObserverTokenInternal {
	rv := objc.Send[GTMioTraceDataObserverTokenInternal](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceDataObserverTokenInternal creates a new GTMioTraceDataObserverTokenInternal instance.
func NewGTMioTraceDataObserverTokenInternal() GTMioTraceDataObserverTokenInternal {
	class := getGTMioTraceDataObserverTokenInternalClass()
	rv := objc.Send[GTMioTraceDataObserverTokenInternal](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataObserverTokenInternal/initWithTraceData:identifier:
func NewGTMioTraceDataObserverTokenInternalWithTraceDataIdentifier(data objectivec.IObject, identifier uint64) GTMioTraceDataObserverTokenInternal {
	instance := getGTMioTraceDataObserverTokenInternalClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraceData:identifier:"), data, identifier)
	return GTMioTraceDataObserverTokenInternalFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataObserverTokenInternal/cancel
func (g GTMioTraceDataObserverTokenInternal) Cancel() {
	objc.Send[objc.ID](g.ID, objc.Sel("cancel"))
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataObserverTokenInternal/initWithTraceData:identifier:
func (g GTMioTraceDataObserverTokenInternal) InitWithTraceDataIdentifier(data objectivec.IObject, identifier uint64) GTMioTraceDataObserverTokenInternal {
	rv := objc.Send[GTMioTraceDataObserverTokenInternal](g.ID, objc.Sel("initWithTraceData:identifier:"), data, identifier)
	return rv
}

