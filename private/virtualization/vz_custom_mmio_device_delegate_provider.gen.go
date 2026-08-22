// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[VZCustomMMIODeviceDelegateProvider](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCustomMMIODeviceDelegateProvider.Delegate]
//   - [VZCustomMMIODeviceDelegateProvider.DeviceQueue]
//   - [VZCustomMMIODeviceDelegateProvider.InitWithDeviceQueueDelegate]
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
type IVZCustomMMIODeviceDelegateProvider interface {
	IVZCustomMMIODeviceProvider

	// Topic: Methods

	Delegate() unsafe.Pointer
	DeviceQueue() objectivec.Object
	InitWithDeviceQueueDelegate(queue objectivec.IObject, delegate objectivec.IObject) VZCustomMMIODeviceDelegateProvider
}

// Init initializes the instance.
func (v VZCustomMMIODeviceDelegateProvider) Init() VZCustomMMIODeviceDelegateProvider {
	rv := objc.SendIfResponds[VZCustomMMIODeviceDelegateProvider](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomMMIODeviceDelegateProvider) Autorelease() VZCustomMMIODeviceDelegateProvider {
	rv := objc.SendIfResponds[VZCustomMMIODeviceDelegateProvider](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomMMIODeviceDelegateProvider creates a new VZCustomMMIODeviceDelegateProvider instance.
func NewVZCustomMMIODeviceDelegateProvider() VZCustomMMIODeviceDelegateProvider {
	class := getVZCustomMMIODeviceDelegateProviderClass()
	rv := objc.SendIfResponds[VZCustomMMIODeviceDelegateProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZCustomMMIODeviceDelegateProviderWithDeviceQueueDelegate(queue objectivec.IObject, delegate objectivec.IObject) VZCustomMMIODeviceDelegateProvider {
	instance := getVZCustomMMIODeviceDelegateProviderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDeviceQueue:delegate:"), queue, delegate)
	return VZCustomMMIODeviceDelegateProviderFromID(rv)
}

func (v VZCustomMMIODeviceDelegateProvider) InitWithDeviceQueueDelegate(queue objectivec.IObject, delegate objectivec.IObject) VZCustomMMIODeviceDelegateProvider {
	rv := objc.SendIfResponds[VZCustomMMIODeviceDelegateProvider](v.ID, objc.Sel("initWithDeviceQueue:delegate:"), queue, delegate)
	return rv
}

func (v VZCustomMMIODeviceDelegateProvider) Delegate() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("delegate"))
	return rv
}
func (v VZCustomMMIODeviceDelegateProvider) DeviceQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("deviceQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
