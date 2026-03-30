// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomVirtioDeviceProvider] class.
var (
	_VZCustomVirtioDeviceProviderClass     VZCustomVirtioDeviceProviderClass
	_VZCustomVirtioDeviceProviderClassOnce sync.Once
)

func getVZCustomVirtioDeviceProviderClass() VZCustomVirtioDeviceProviderClass {
	_VZCustomVirtioDeviceProviderClassOnce.Do(func() {
		_VZCustomVirtioDeviceProviderClass = VZCustomVirtioDeviceProviderClass{class: objc.GetClass("_VZCustomVirtioDeviceProvider")}
	})
	return _VZCustomVirtioDeviceProviderClass
}

// GetVZCustomVirtioDeviceProviderClass returns the class object for _VZCustomVirtioDeviceProvider.
func GetVZCustomVirtioDeviceProviderClass() VZCustomVirtioDeviceProviderClass {
	return getVZCustomVirtioDeviceProviderClass()
}

type VZCustomVirtioDeviceProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomVirtioDeviceProviderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomVirtioDeviceProviderClass) Alloc() VZCustomVirtioDeviceProvider {
	rv := objc.Send[VZCustomVirtioDeviceProvider](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCustomVirtioDeviceProvider._connectionIdentifier]
//   - [VZCustomVirtioDeviceProvider._init]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceProvider
type VZCustomVirtioDeviceProvider struct {
	objectivec.Object
}

// VZCustomVirtioDeviceProviderFromID constructs a [VZCustomVirtioDeviceProvider] from an objc.ID.
func VZCustomVirtioDeviceProviderFromID(id objc.ID) VZCustomVirtioDeviceProvider {
	return VZCustomVirtioDeviceProvider{objectivec.Object{ID: id}}
}

// Ensure VZCustomVirtioDeviceProvider implements IVZCustomVirtioDeviceProvider.
var _ IVZCustomVirtioDeviceProvider = VZCustomVirtioDeviceProvider{}

// An interface definition for the [VZCustomVirtioDeviceProvider] class.
//
// # Methods
//
//   - [IVZCustomVirtioDeviceProvider._connectionIdentifier]
//   - [IVZCustomVirtioDeviceProvider._init]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceProvider
type IVZCustomVirtioDeviceProvider interface {
	objectivec.IObject

	// Topic: Methods

	_connectionIdentifier() objectivec.IObject
	_init() objectivec.IObject
}

// Init initializes the instance.
func (v VZCustomVirtioDeviceProvider) Init() VZCustomVirtioDeviceProvider {
	rv := objc.Send[VZCustomVirtioDeviceProvider](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomVirtioDeviceProvider) Autorelease() VZCustomVirtioDeviceProvider {
	rv := objc.Send[VZCustomVirtioDeviceProvider](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomVirtioDeviceProvider creates a new VZCustomVirtioDeviceProvider instance.
func NewVZCustomVirtioDeviceProvider() VZCustomVirtioDeviceProvider {
	class := getVZCustomVirtioDeviceProviderClass()
	rv := objc.Send[VZCustomVirtioDeviceProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceProvider/_init
func (v VZCustomVirtioDeviceProvider) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDeviceProvider/_connectionIdentifier
func (v VZCustomVirtioDeviceProvider) _connectionIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_connectionIdentifier"))
	return objectivec.Object{ID: rv}
}
