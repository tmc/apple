// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[VZCustomVirtioDeviceProvider](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCustomVirtioDeviceProvider._connectionIdentifier]
//   - [VZCustomVirtioDeviceProvider._init]
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
type IVZCustomVirtioDeviceProvider interface {
	objectivec.IObject

	// Topic: Methods

	_connectionIdentifier() unsafe.Pointer
	_init() objectivec.IObject
}

// Init initializes the instance.
func (v VZCustomVirtioDeviceProvider) Init() VZCustomVirtioDeviceProvider {
	rv := objc.SendIfResponds[VZCustomVirtioDeviceProvider](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomVirtioDeviceProvider) Autorelease() VZCustomVirtioDeviceProvider {
	rv := objc.SendIfResponds[VZCustomVirtioDeviceProvider](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomVirtioDeviceProvider creates a new VZCustomVirtioDeviceProvider instance.
func NewVZCustomVirtioDeviceProvider() VZCustomVirtioDeviceProvider {
	class := getVZCustomVirtioDeviceProviderClass()
	rv := objc.SendIfResponds[VZCustomVirtioDeviceProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZCustomVirtioDeviceProvider) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (v VZCustomVirtioDeviceProvider) _connectionIdentifier() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("_connectionIdentifier"))
	return rv
}

// CanConnectionIdentifier reports whether the receiver responds to the private selector _connectionIdentifier.
func (v VZCustomVirtioDeviceProvider) CanConnectionIdentifier() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_connectionIdentifier"))
}

// ConnectionIdentifier is an exported wrapper for the private property _connectionIdentifier.
func (v VZCustomVirtioDeviceProvider) ConnectionIdentifier() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_connectionIdentifier")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_connectionIdentifier"}
	}
	return v._connectionIdentifier(), nil
}
