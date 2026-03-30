// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGraphicsDeviceConfiguration] class.
var (
	_VZGraphicsDeviceConfigurationClass     VZGraphicsDeviceConfigurationClass
	_VZGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZGraphicsDeviceConfigurationClass() VZGraphicsDeviceConfigurationClass {
	_VZGraphicsDeviceConfigurationClassOnce.Do(func() {
		_VZGraphicsDeviceConfigurationClass = VZGraphicsDeviceConfigurationClass{class: objc.GetClass("VZGraphicsDeviceConfiguration")}
	})
	return _VZGraphicsDeviceConfigurationClass
}

// GetVZGraphicsDeviceConfigurationClass returns the class object for VZGraphicsDeviceConfiguration.
func GetVZGraphicsDeviceConfigurationClass() VZGraphicsDeviceConfigurationClass {
	return getVZGraphicsDeviceConfigurationClass()
}

type VZGraphicsDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGraphicsDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGraphicsDeviceConfigurationClass) Alloc() VZGraphicsDeviceConfiguration {
	rv := objc.Send[VZGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZGraphicsDeviceConfiguration._init]
//   - [VZGraphicsDeviceConfiguration._initWithConfiguration]
//   - [VZGraphicsDeviceConfiguration.MakeGraphicsDeviceForVirtualMachineGraphicsDeviceIndex]
//   - [VZGraphicsDeviceConfiguration.ValidateWithError]
//   - [VZGraphicsDeviceConfiguration._graphicsDevice]
//   - [VZGraphicsDeviceConfiguration.DebugDescription]
//   - [VZGraphicsDeviceConfiguration.Description]
//   - [VZGraphicsDeviceConfiguration.Hash]
//   - [VZGraphicsDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration
type VZGraphicsDeviceConfiguration struct {
	objectivec.Object
}

// VZGraphicsDeviceConfigurationFromID constructs a [VZGraphicsDeviceConfiguration] from an objc.ID.
func VZGraphicsDeviceConfigurationFromID(id objc.ID) VZGraphicsDeviceConfiguration {
	return VZGraphicsDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZGraphicsDeviceConfiguration implements IVZGraphicsDeviceConfiguration.
var _ IVZGraphicsDeviceConfiguration = VZGraphicsDeviceConfiguration{}

// An interface definition for the [VZGraphicsDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZGraphicsDeviceConfiguration._init]
//   - [IVZGraphicsDeviceConfiguration._initWithConfiguration]
//   - [IVZGraphicsDeviceConfiguration.MakeGraphicsDeviceForVirtualMachineGraphicsDeviceIndex]
//   - [IVZGraphicsDeviceConfiguration.ValidateWithError]
//   - [IVZGraphicsDeviceConfiguration._graphicsDevice]
//   - [IVZGraphicsDeviceConfiguration.DebugDescription]
//   - [IVZGraphicsDeviceConfiguration.Description]
//   - [IVZGraphicsDeviceConfiguration.Hash]
//   - [IVZGraphicsDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration
type IVZGraphicsDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_initWithConfiguration(configuration unsafe.Pointer) objectivec.IObject
	MakeGraphicsDeviceForVirtualMachineGraphicsDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	ValidateWithError() (bool, error)
	_graphicsDevice() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g VZGraphicsDeviceConfiguration) Init() VZGraphicsDeviceConfiguration {
	rv := objc.Send[VZGraphicsDeviceConfiguration](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VZGraphicsDeviceConfiguration) Autorelease() VZGraphicsDeviceConfiguration {
	rv := objc.Send[VZGraphicsDeviceConfiguration](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDeviceConfiguration creates a new VZGraphicsDeviceConfiguration instance.
func NewVZGraphicsDeviceConfiguration() VZGraphicsDeviceConfiguration {
	class := getVZGraphicsDeviceConfigurationClass()
	rv := objc.Send[VZGraphicsDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration/_init
func (g VZGraphicsDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration/_initWithConfiguration:
func (g VZGraphicsDeviceConfiguration) _initWithConfiguration(configuration unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_initWithConfiguration:"), configuration)
	return objectivec.Object{ID: rv}
}

// InitWithConfiguration is an exported wrapper for the private method _initWithConfiguration.
func (g VZGraphicsDeviceConfiguration) InitWithConfiguration(configuration unsafe.Pointer) objectivec.IObject {
	return g._initWithConfiguration(configuration)
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration/makeGraphicsDeviceForVirtualMachine:graphicsDeviceIndex:
func (g VZGraphicsDeviceConfiguration) MakeGraphicsDeviceForVirtualMachineGraphicsDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("makeGraphicsDeviceForVirtualMachine:graphicsDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration/validateWithError:
func (g VZGraphicsDeviceConfiguration) ValidateWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](g.ID, objc.Sel("validateWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration/_graphicsDevice
func (g VZGraphicsDeviceConfiguration) _graphicsDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_graphicsDevice"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration/debugDescription
func (g VZGraphicsDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration/description
func (g VZGraphicsDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration/hash
func (g VZGraphicsDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration/superclass
func (g VZGraphicsDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
