// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomMMIODeviceDelegateProvider] class.
var (
	_VZCustomMMIODeviceDelegateProviderClass     VZCustomMMIODeviceDelegateProviderClass
	_VZCustomMMIODeviceDelegateProviderClassOnce sync.Once
)

func getVZCustomMMIODeviceDelegateProviderClass() VZCustomMMIODeviceDelegateProviderClass {
	_VZCustomMMIODeviceDelegateProviderClassOnce.Do(func() {
		_VZCustomMMIODeviceDelegateProviderClass = VZCustomMMIODeviceDelegateProviderClass{class: objc.GetClass("_VZCustomMMIODeviceDelegateProvider")}
	})
	return _VZCustomMMIODeviceDelegateProviderClass
}

// GetVZCustomMMIODeviceDelegateProviderClass returns the class object for _VZCustomMMIODeviceDelegateProvider.
func GetVZCustomMMIODeviceDelegateProviderClass() VZCustomMMIODeviceDelegateProviderClass {
	return getVZCustomMMIODeviceDelegateProviderClass()
}

type VZCustomMMIODeviceDelegateProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomMMIODeviceDelegateProviderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomMMIODeviceDelegateProviderClass) Alloc() VZCustomMMIODeviceDelegateProvider {
	rv := objc.Send[VZCustomMMIODeviceDelegateProvider](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZCustomMMIODeviceDelegateProvider.Delegate]
//   - [VZCustomMMIODeviceDelegateProvider.DeviceQueue]
//   - [VZCustomMMIODeviceDelegateProvider.InitWithDeviceQueueDelegate]
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceDelegateProvider
type VZCustomMMIODeviceDelegateProvider struct {
	VZCustomMMIODeviceProvider
}

// VZCustomMMIODeviceDelegateProviderFromID constructs a [VZCustomMMIODeviceDelegateProvider] from an objc.ID.
func VZCustomMMIODeviceDelegateProviderFromID(id objc.ID) VZCustomMMIODeviceDelegateProvider {
	return VZCustomMMIODeviceDelegateProvider{VZCustomMMIODeviceProvider: VZCustomMMIODeviceProviderFromID(id)}
}
// Ensure VZCustomMMIODeviceDelegateProvider implements IVZCustomMMIODeviceDelegateProvider.
var _ IVZCustomMMIODeviceDelegateProvider = VZCustomMMIODeviceDelegateProvider{}

// An interface definition for the [VZCustomMMIODeviceDelegateProvider] class.
//
// # Methods
//
//   - [IVZCustomMMIODeviceDelegateProvider.Delegate]
//   - [IVZCustomMMIODeviceDelegateProvider.DeviceQueue]
//   - [IVZCustomMMIODeviceDelegateProvider.InitWithDeviceQueueDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceDelegateProvider
type IVZCustomMMIODeviceDelegateProvider interface {
	IVZCustomMMIODeviceProvider

	// Topic: Methods

	Delegate() objectivec.IObject
	DeviceQueue() objectivec.Object
	InitWithDeviceQueueDelegate(queue objectivec.IObject, delegate objectivec.IObject) VZCustomMMIODeviceDelegateProvider
}

// Init initializes the instance.
func (v VZCustomMMIODeviceDelegateProvider) Init() VZCustomMMIODeviceDelegateProvider {
	rv := objc.Send[VZCustomMMIODeviceDelegateProvider](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomMMIODeviceDelegateProvider) Autorelease() VZCustomMMIODeviceDelegateProvider {
	rv := objc.Send[VZCustomMMIODeviceDelegateProvider](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomMMIODeviceDelegateProvider creates a new VZCustomMMIODeviceDelegateProvider instance.
func NewVZCustomMMIODeviceDelegateProvider() VZCustomMMIODeviceDelegateProvider {
	class := getVZCustomMMIODeviceDelegateProviderClass()
	rv := objc.Send[VZCustomMMIODeviceDelegateProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceDelegateProvider/initWithDeviceQueue:delegate:
func NewVZCustomMMIODeviceDelegateProviderWithDeviceQueueDelegate(queue objectivec.IObject, delegate objectivec.IObject) VZCustomMMIODeviceDelegateProvider {
	instance := getVZCustomMMIODeviceDelegateProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDeviceQueue:delegate:"), queue, delegate)
	return VZCustomMMIODeviceDelegateProviderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceDelegateProvider/initWithDeviceQueue:delegate:
func (v VZCustomMMIODeviceDelegateProvider) InitWithDeviceQueueDelegate(queue objectivec.IObject, delegate objectivec.IObject) VZCustomMMIODeviceDelegateProvider {
	rv := objc.Send[VZCustomMMIODeviceDelegateProvider](v.ID, objc.Sel("initWithDeviceQueue:delegate:"), queue, delegate)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceDelegateProvider/delegate
func (v VZCustomMMIODeviceDelegateProvider) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceDelegateProvider/deviceQueue
func (v VZCustomMMIODeviceDelegateProvider) DeviceQueue() objectivec.Object {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("deviceQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

