// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomMMIODeviceProvider] class.
var (
	_VZCustomMMIODeviceProviderClass     VZCustomMMIODeviceProviderClass
	_VZCustomMMIODeviceProviderClassOnce sync.Once
)

func getVZCustomMMIODeviceProviderClass() VZCustomMMIODeviceProviderClass {
	_VZCustomMMIODeviceProviderClassOnce.Do(func() {
		_VZCustomMMIODeviceProviderClass = VZCustomMMIODeviceProviderClass{class: objc.GetClass("_VZCustomMMIODeviceProvider")}
	})
	return _VZCustomMMIODeviceProviderClass
}

// GetVZCustomMMIODeviceProviderClass returns the class object for _VZCustomMMIODeviceProvider.
func GetVZCustomMMIODeviceProviderClass() VZCustomMMIODeviceProviderClass {
	return getVZCustomMMIODeviceProviderClass()
}

type VZCustomMMIODeviceProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomMMIODeviceProviderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomMMIODeviceProviderClass) Alloc() VZCustomMMIODeviceProvider {
	rv := objc.Send[VZCustomMMIODeviceProvider](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCustomMMIODeviceProvider._connectionIdentifier]
//   - [VZCustomMMIODeviceProvider._init]
//   - [VZCustomMMIODeviceProvider.EncodeWithEncoder]
//   - [VZCustomMMIODeviceProvider.DebugDescription]
//   - [VZCustomMMIODeviceProvider.Description]
//   - [VZCustomMMIODeviceProvider.Hash]
//   - [VZCustomMMIODeviceProvider.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceProvider
type VZCustomMMIODeviceProvider struct {
	objectivec.Object
}

// VZCustomMMIODeviceProviderFromID constructs a [VZCustomMMIODeviceProvider] from an objc.ID.
func VZCustomMMIODeviceProviderFromID(id objc.ID) VZCustomMMIODeviceProvider {
	return VZCustomMMIODeviceProvider{objectivec.Object{ID: id}}
}

// Ensure VZCustomMMIODeviceProvider implements IVZCustomMMIODeviceProvider.
var _ IVZCustomMMIODeviceProvider = VZCustomMMIODeviceProvider{}

// An interface definition for the [VZCustomMMIODeviceProvider] class.
//
// # Methods
//
//   - [IVZCustomMMIODeviceProvider._connectionIdentifier]
//   - [IVZCustomMMIODeviceProvider._init]
//   - [IVZCustomMMIODeviceProvider.EncodeWithEncoder]
//   - [IVZCustomMMIODeviceProvider.DebugDescription]
//   - [IVZCustomMMIODeviceProvider.Description]
//   - [IVZCustomMMIODeviceProvider.Hash]
//   - [IVZCustomMMIODeviceProvider.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceProvider
type IVZCustomMMIODeviceProvider interface {
	objectivec.IObject

	// Topic: Methods

	_connectionIdentifier() objectivec.IObject
	_init() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZCustomMMIODeviceProvider) Init() VZCustomMMIODeviceProvider {
	rv := objc.Send[VZCustomMMIODeviceProvider](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomMMIODeviceProvider) Autorelease() VZCustomMMIODeviceProvider {
	rv := objc.Send[VZCustomMMIODeviceProvider](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomMMIODeviceProvider creates a new VZCustomMMIODeviceProvider instance.
func NewVZCustomMMIODeviceProvider() VZCustomMMIODeviceProvider {
	class := getVZCustomMMIODeviceProviderClass()
	rv := objc.Send[VZCustomMMIODeviceProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceProvider/_init
func (v VZCustomMMIODeviceProvider) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceProvider/encodeWithEncoder:
func (v VZCustomMMIODeviceProvider) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceProvider/_connectionIdentifier
func (v VZCustomMMIODeviceProvider) _connectionIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_connectionIdentifier"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceProvider/debugDescription
func (v VZCustomMMIODeviceProvider) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceProvider/description
func (v VZCustomMMIODeviceProvider) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceProvider/hash
func (v VZCustomMMIODeviceProvider) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceProvider/superclass
func (v VZCustomMMIODeviceProvider) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
