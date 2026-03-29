// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMetalDeviceObserver] class.
var (
	_MLMetalDeviceObserverClass     MLMetalDeviceObserverClass
	_MLMetalDeviceObserverClassOnce sync.Once
)

func getMLMetalDeviceObserverClass() MLMetalDeviceObserverClass {
	_MLMetalDeviceObserverClassOnce.Do(func() {
		_MLMetalDeviceObserverClass = MLMetalDeviceObserverClass{class: objc.GetClass("MLMetalDeviceObserver")}
	})
	return _MLMetalDeviceObserverClass
}

// GetMLMetalDeviceObserverClass returns the class object for MLMetalDeviceObserver.
func GetMLMetalDeviceObserverClass() MLMetalDeviceObserverClass {
	return getMLMetalDeviceObserverClass()
}

type MLMetalDeviceObserverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMetalDeviceObserverClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMetalDeviceObserverClass) Alloc() MLMetalDeviceObserver {
	rv := objc.Send[MLMetalDeviceObserver](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLMetalDeviceObserver.CopyAllMTLDevicesWithHandlerDeviceObserver]
//   - [MLMetalDeviceObserver.StartObservingWithBlockDeviceObserver]
//   - [MLMetalDeviceObserver.StopObserving]
// See: https://developer.apple.com/documentation/CoreML/MLMetalDeviceObserver
type MLMetalDeviceObserver struct {
	objectivec.Object
}

// MLMetalDeviceObserverFromID constructs a [MLMetalDeviceObserver] from an objc.ID.
func MLMetalDeviceObserverFromID(id objc.ID) MLMetalDeviceObserver {
	return MLMetalDeviceObserver{objectivec.Object{ID: id}}
}
// Ensure MLMetalDeviceObserver implements IMLMetalDeviceObserver.
var _ IMLMetalDeviceObserver = MLMetalDeviceObserver{}

// An interface definition for the [MLMetalDeviceObserver] class.
//
// # Methods
//
//   - [IMLMetalDeviceObserver.CopyAllMTLDevicesWithHandlerDeviceObserver]
//   - [IMLMetalDeviceObserver.StartObservingWithBlockDeviceObserver]
//   - [IMLMetalDeviceObserver.StopObserving]
//
// See: https://developer.apple.com/documentation/CoreML/MLMetalDeviceObserver
type IMLMetalDeviceObserver interface {
	objectivec.IObject

	// Topic: Methods

	CopyAllMTLDevicesWithHandlerDeviceObserver(handler VoidHandler, observer []objectivec.IObject) objectivec.IObject
	StartObservingWithBlockDeviceObserver(block VoidHandler, observer []objectivec.IObject) objectivec.IObject
	StopObserving(observing objectivec.IObject)
}

// Init initializes the instance.
func (m MLMetalDeviceObserver) Init() MLMetalDeviceObserver {
	rv := objc.Send[MLMetalDeviceObserver](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMetalDeviceObserver) Autorelease() MLMetalDeviceObserver {
	rv := objc.Send[MLMetalDeviceObserver](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMetalDeviceObserver creates a new MLMetalDeviceObserver instance.
func NewMLMetalDeviceObserver() MLMetalDeviceObserver {
	class := getMLMetalDeviceObserverClass()
	rv := objc.Send[MLMetalDeviceObserver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMetalDeviceObserver/copyAllMTLDevicesWithHandler:deviceObserver:
func (m MLMetalDeviceObserver) CopyAllMTLDevicesWithHandlerDeviceObserver(handler VoidHandler, observer []objectivec.IObject) objectivec.IObject {
_block0, _ := NewVoidBlock(handler)
	rv := objc.Send[objc.ID](m.ID, objc.Sel("copyAllMTLDevicesWithHandler:deviceObserver:"), _block0, observer)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMetalDeviceObserver/startObservingWithBlock:deviceObserver:
func (m MLMetalDeviceObserver) StartObservingWithBlockDeviceObserver(block VoidHandler, observer []objectivec.IObject) objectivec.IObject {
_block0, _ := NewVoidBlock(block)
	rv := objc.Send[objc.ID](m.ID, objc.Sel("startObservingWithBlock:deviceObserver:"), _block0, observer)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLMetalDeviceObserver/stopObserving:
func (m MLMetalDeviceObserver) StopObserving(observing objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("stopObserving:"), observing)
}

