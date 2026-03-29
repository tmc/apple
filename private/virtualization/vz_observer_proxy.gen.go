// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZObserverProxy] class.
var (
	_VZObserverProxyClass     VZObserverProxyClass
	_VZObserverProxyClassOnce sync.Once
)

func getVZObserverProxyClass() VZObserverProxyClass {
	_VZObserverProxyClassOnce.Do(func() {
		_VZObserverProxyClass = VZObserverProxyClass{class: objc.GetClass("VZObserverProxy")}
	})
	return _VZObserverProxyClass
}

// GetVZObserverProxyClass returns the class object for VZObserverProxy.
func GetVZObserverProxyClass() VZObserverProxyClass {
	return getVZObserverProxyClass()
}

type VZObserverProxyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZObserverProxyClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZObserverProxyClass) Alloc() VZObserverProxy {
	rv := objc.Send[VZObserverProxy](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZObserverProxy
type VZObserverProxy struct {
	objectivec.Object
}

// VZObserverProxyFromID constructs a [VZObserverProxy] from an objc.ID.
func VZObserverProxyFromID(id objc.ID) VZObserverProxy {
	return VZObserverProxy{objectivec.Object{ID: id}}
}
// Ensure VZObserverProxy implements IVZObserverProxy.
var _ IVZObserverProxy = VZObserverProxy{}

// An interface definition for the [VZObserverProxy] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZObserverProxy
type IVZObserverProxy interface {
	objectivec.IObject
}

// Init initializes the instance.
func (o VZObserverProxy) Init() VZObserverProxy {
	rv := objc.Send[VZObserverProxy](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o VZObserverProxy) Autorelease() VZObserverProxy {
	rv := objc.Send[VZObserverProxy](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZObserverProxy creates a new VZObserverProxy instance.
func NewVZObserverProxy() VZObserverProxy {
	class := getVZObserverProxyClass()
	rv := objc.Send[VZObserverProxy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

