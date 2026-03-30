// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioSocketDeviceObserver] class.
var (
	_VZVirtioSocketDeviceObserverClass     VZVirtioSocketDeviceObserverClass
	_VZVirtioSocketDeviceObserverClassOnce sync.Once
)

func getVZVirtioSocketDeviceObserverClass() VZVirtioSocketDeviceObserverClass {
	_VZVirtioSocketDeviceObserverClassOnce.Do(func() {
		_VZVirtioSocketDeviceObserverClass = VZVirtioSocketDeviceObserverClass{class: objc.GetClass("_VZVirtioSocketDeviceObserver")}
	})
	return _VZVirtioSocketDeviceObserverClass
}

// GetVZVirtioSocketDeviceObserverClass returns the class object for _VZVirtioSocketDeviceObserver.
func GetVZVirtioSocketDeviceObserverClass() VZVirtioSocketDeviceObserverClass {
	return getVZVirtioSocketDeviceObserverClass()
}

type VZVirtioSocketDeviceObserverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSocketDeviceObserverClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSocketDeviceObserverClass) Alloc() VZVirtioSocketDeviceObserver {
	rv := objc.Send[VZVirtioSocketDeviceObserver](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtioSocketDeviceObserver._initWithConnectionQueueDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioSocketDeviceObserver
type VZVirtioSocketDeviceObserver struct {
	objectivec.Object
}

// VZVirtioSocketDeviceObserverFromID constructs a [VZVirtioSocketDeviceObserver] from an objc.ID.
func VZVirtioSocketDeviceObserverFromID(id objc.ID) VZVirtioSocketDeviceObserver {
	return VZVirtioSocketDeviceObserver{objectivec.Object{ID: id}}
}

// Ensure VZVirtioSocketDeviceObserver implements IVZVirtioSocketDeviceObserver.
var _ IVZVirtioSocketDeviceObserver = VZVirtioSocketDeviceObserver{}

// An interface definition for the [VZVirtioSocketDeviceObserver] class.
//
// # Methods
//
//   - [IVZVirtioSocketDeviceObserver._initWithConnectionQueueDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioSocketDeviceObserver
type IVZVirtioSocketDeviceObserver interface {
	objectivec.IObject

	// Topic: Methods

	_initWithConnectionQueueDelegate(connection objectivec.IObject, queue objectivec.IObject, delegate objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (v VZVirtioSocketDeviceObserver) Init() VZVirtioSocketDeviceObserver {
	rv := objc.Send[VZVirtioSocketDeviceObserver](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSocketDeviceObserver) Autorelease() VZVirtioSocketDeviceObserver {
	rv := objc.Send[VZVirtioSocketDeviceObserver](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketDeviceObserver creates a new VZVirtioSocketDeviceObserver instance.
func NewVZVirtioSocketDeviceObserver() VZVirtioSocketDeviceObserver {
	class := getVZVirtioSocketDeviceObserverClass()
	rv := objc.Send[VZVirtioSocketDeviceObserver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioSocketDeviceObserver/_initWithConnection:queue:delegate:
func (v VZVirtioSocketDeviceObserver) _initWithConnectionQueueDelegate(connection objectivec.IObject, queue objectivec.IObject, delegate objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_initWithConnection:queue:delegate:"), connection, queue, delegate)
	return objectivec.Object{ID: rv}
}

// InitWithConnectionQueueDelegate is an exported wrapper for the private method _initWithConnectionQueueDelegate.
func (v VZVirtioSocketDeviceObserver) InitWithConnectionQueueDelegate(connection objectivec.IObject, queue objectivec.IObject, delegate objectivec.IObject) objectivec.IObject {
	return v._initWithConnectionQueueDelegate(connection, queue, delegate)
}
