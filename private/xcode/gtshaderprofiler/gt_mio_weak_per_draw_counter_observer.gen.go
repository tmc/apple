// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioWeakPerDrawCounterObserver] class.
var (
	_GTMioWeakPerDrawCounterObserverClass     GTMioWeakPerDrawCounterObserverClass
	_GTMioWeakPerDrawCounterObserverClassOnce sync.Once
)

func getGTMioWeakPerDrawCounterObserverClass() GTMioWeakPerDrawCounterObserverClass {
	_GTMioWeakPerDrawCounterObserverClassOnce.Do(func() {
		_GTMioWeakPerDrawCounterObserverClass = GTMioWeakPerDrawCounterObserverClass{class: objc.GetClass("GTMioWeakPerDrawCounterObserver")}
	})
	return _GTMioWeakPerDrawCounterObserverClass
}

// GetGTMioWeakPerDrawCounterObserverClass returns the class object for GTMioWeakPerDrawCounterObserver.
func GetGTMioWeakPerDrawCounterObserverClass() GTMioWeakPerDrawCounterObserverClass {
	return getGTMioWeakPerDrawCounterObserverClass()
}

type GTMioWeakPerDrawCounterObserverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioWeakPerDrawCounterObserverClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioWeakPerDrawCounterObserverClass) Alloc() GTMioWeakPerDrawCounterObserver {
	rv := objc.Send[GTMioWeakPerDrawCounterObserver](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioWeakPerDrawCounterObserver.Observer]
//   - [GTMioWeakPerDrawCounterObserver.SetObserver]
//   - [GTMioWeakPerDrawCounterObserver.InitWithObserver]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioWeakPerDrawCounterObserver
type GTMioWeakPerDrawCounterObserver struct {
	objectivec.Object
}

// GTMioWeakPerDrawCounterObserverFromID constructs a [GTMioWeakPerDrawCounterObserver] from an objc.ID.
func GTMioWeakPerDrawCounterObserverFromID(id objc.ID) GTMioWeakPerDrawCounterObserver {
	return GTMioWeakPerDrawCounterObserver{objectivec.Object{ID: id}}
}

// Ensure GTMioWeakPerDrawCounterObserver implements IGTMioWeakPerDrawCounterObserver.
var _ IGTMioWeakPerDrawCounterObserver = GTMioWeakPerDrawCounterObserver{}

// An interface definition for the [GTMioWeakPerDrawCounterObserver] class.
//
// # Methods
//
//   - [IGTMioWeakPerDrawCounterObserver.Observer]
//   - [IGTMioWeakPerDrawCounterObserver.SetObserver]
//   - [IGTMioWeakPerDrawCounterObserver.InitWithObserver]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioWeakPerDrawCounterObserver
type IGTMioWeakPerDrawCounterObserver interface {
	objectivec.IObject

	// Topic: Methods

	Observer() objectivec.IObject
	SetObserver(value objectivec.IObject)
	InitWithObserver(observer objectivec.IObject) GTMioWeakPerDrawCounterObserver
}

// Init initializes the instance.
func (g GTMioWeakPerDrawCounterObserver) Init() GTMioWeakPerDrawCounterObserver {
	rv := objc.Send[GTMioWeakPerDrawCounterObserver](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioWeakPerDrawCounterObserver) Autorelease() GTMioWeakPerDrawCounterObserver {
	rv := objc.Send[GTMioWeakPerDrawCounterObserver](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioWeakPerDrawCounterObserver creates a new GTMioWeakPerDrawCounterObserver instance.
func NewGTMioWeakPerDrawCounterObserver() GTMioWeakPerDrawCounterObserver {
	class := getGTMioWeakPerDrawCounterObserverClass()
	rv := objc.Send[GTMioWeakPerDrawCounterObserver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioWeakPerDrawCounterObserver/initWithObserver:
func NewGTMioWeakPerDrawCounterObserverWithObserver(observer objectivec.IObject) GTMioWeakPerDrawCounterObserver {
	instance := getGTMioWeakPerDrawCounterObserverClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObserver:"), observer)
	return GTMioWeakPerDrawCounterObserverFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioWeakPerDrawCounterObserver/initWithObserver:
func (g GTMioWeakPerDrawCounterObserver) InitWithObserver(observer objectivec.IObject) GTMioWeakPerDrawCounterObserver {
	rv := objc.Send[GTMioWeakPerDrawCounterObserver](g.ID, objc.Sel("initWithObserver:"), observer)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioWeakPerDrawCounterObserver/observer
func (g GTMioWeakPerDrawCounterObserver) Observer() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("observer"))
	return objectivec.Object{ID: rv}
}
func (g GTMioWeakPerDrawCounterObserver) SetObserver(value objectivec.IObject) {
	objc.Send[struct{}](g.ID, objc.Sel("setObserver:"), value)
}
