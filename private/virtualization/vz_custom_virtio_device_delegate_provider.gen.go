// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomVirtioDeviceDelegateProvider] class.
var (
	_VZCustomVirtioDeviceDelegateProviderClass     VZCustomVirtioDeviceDelegateProviderClass
	_VZCustomVirtioDeviceDelegateProviderClassOnce sync.Once
)

func getVZCustomVirtioDeviceDelegateProviderClass() VZCustomVirtioDeviceDelegateProviderClass {
	_VZCustomVirtioDeviceDelegateProviderClassOnce.Do(func() {
		_VZCustomVirtioDeviceDelegateProviderClass = VZCustomVirtioDeviceDelegateProviderClass{class: objc.GetClass("_VZCustomVirtioDeviceDelegateProvider")}
	})
	return _VZCustomVirtioDeviceDelegateProviderClass
}

// GetVZCustomVirtioDeviceDelegateProviderClass returns the class object for _VZCustomVirtioDeviceDelegateProvider.
func GetVZCustomVirtioDeviceDelegateProviderClass() VZCustomVirtioDeviceDelegateProviderClass {
	return getVZCustomVirtioDeviceDelegateProviderClass()
}

type VZCustomVirtioDeviceDelegateProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomVirtioDeviceDelegateProviderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomVirtioDeviceDelegateProviderClass) Alloc() VZCustomVirtioDeviceDelegateProvider {
	rv := objc.Send[VZCustomVirtioDeviceDelegateProvider](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZCustomVirtioDeviceDelegateProvider.Delegate]
//   - [VZCustomVirtioDeviceDelegateProvider.DeviceQueue]
//   - [VZCustomVirtioDeviceDelegateProvider.InitWithDeviceQueueDelegate]
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceDelegateProvider
type VZCustomVirtioDeviceDelegateProvider struct {
	VZCustomVirtioDeviceProvider
}

// VZCustomVirtioDeviceDelegateProviderFromID constructs a [VZCustomVirtioDeviceDelegateProvider] from an objc.ID.
func VZCustomVirtioDeviceDelegateProviderFromID(id objc.ID) VZCustomVirtioDeviceDelegateProvider {
	return VZCustomVirtioDeviceDelegateProvider{VZCustomVirtioDeviceProvider: VZCustomVirtioDeviceProviderFromID(id)}
}
// Ensure VZCustomVirtioDeviceDelegateProvider implements IVZCustomVirtioDeviceDelegateProvider.
var _ IVZCustomVirtioDeviceDelegateProvider = VZCustomVirtioDeviceDelegateProvider{}

// An interface definition for the [VZCustomVirtioDeviceDelegateProvider] class.
//
// # Methods
//
//   - [IVZCustomVirtioDeviceDelegateProvider.Delegate]
//   - [IVZCustomVirtioDeviceDelegateProvider.DeviceQueue]
//   - [IVZCustomVirtioDeviceDelegateProvider.InitWithDeviceQueueDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceDelegateProvider
type IVZCustomVirtioDeviceDelegateProvider interface {
	IVZCustomVirtioDeviceProvider

	// Topic: Methods

	Delegate() objectivec.IObject
	DeviceQueue() objectivec.Object
	InitWithDeviceQueueDelegate(queue objectivec.IObject, delegate objectivec.IObject) VZCustomVirtioDeviceDelegateProvider
}

// Init initializes the instance.
func (v VZCustomVirtioDeviceDelegateProvider) Init() VZCustomVirtioDeviceDelegateProvider {
	rv := objc.Send[VZCustomVirtioDeviceDelegateProvider](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomVirtioDeviceDelegateProvider) Autorelease() VZCustomVirtioDeviceDelegateProvider {
	rv := objc.Send[VZCustomVirtioDeviceDelegateProvider](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomVirtioDeviceDelegateProvider creates a new VZCustomVirtioDeviceDelegateProvider instance.
func NewVZCustomVirtioDeviceDelegateProvider() VZCustomVirtioDeviceDelegateProvider {
	class := getVZCustomVirtioDeviceDelegateProviderClass()
	rv := objc.Send[VZCustomVirtioDeviceDelegateProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceDelegateProvider/initWithDeviceQueue:delegate:
func NewVZCustomVirtioDeviceDelegateProviderWithDeviceQueueDelegate(queue objectivec.IObject, delegate objectivec.IObject) VZCustomVirtioDeviceDelegateProvider {
	instance := getVZCustomVirtioDeviceDelegateProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDeviceQueue:delegate:"), queue, delegate)
	return VZCustomVirtioDeviceDelegateProviderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceDelegateProvider/initWithDeviceQueue:delegate:
func (v VZCustomVirtioDeviceDelegateProvider) InitWithDeviceQueueDelegate(queue objectivec.IObject, delegate objectivec.IObject) VZCustomVirtioDeviceDelegateProvider {
	rv := objc.Send[VZCustomVirtioDeviceDelegateProvider](v.ID, objc.Sel("initWithDeviceQueue:delegate:"), queue, delegate)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceDelegateProvider/delegate
func (v VZCustomVirtioDeviceDelegateProvider) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceDelegateProvider/deviceQueue
func (v VZCustomVirtioDeviceDelegateProvider) DeviceQueue() objectivec.Object {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("deviceQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

