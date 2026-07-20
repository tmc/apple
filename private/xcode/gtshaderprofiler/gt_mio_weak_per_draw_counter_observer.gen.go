// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

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
type IGTMioWeakPerDrawCounterObserver interface {
	objectivec.IObject

	// Topic: Methods

	Observer() unsafe.Pointer
	SetObserver(value unsafe.Pointer)
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

func NewGTMioWeakPerDrawCounterObserverWithObserver(observer objectivec.IObject) GTMioWeakPerDrawCounterObserver {
	instance := getGTMioWeakPerDrawCounterObserverClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObserver:"), observer)
	return GTMioWeakPerDrawCounterObserverFromID(rv)
}

func (g GTMioWeakPerDrawCounterObserver) InitWithObserver(observer objectivec.IObject) GTMioWeakPerDrawCounterObserver {
	rv := objc.Send[GTMioWeakPerDrawCounterObserver](g.ID, objc.Sel("initWithObserver:"), observer)
	return rv
}

func (g GTMioWeakPerDrawCounterObserver) Observer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("observer"))
	return rv
}
func (g GTMioWeakPerDrawCounterObserver) SetObserver(value unsafe.Pointer) {
	objc.Send[struct{}](g.ID, objc.Sel("setObserver:"), value)
}
