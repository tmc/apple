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
	rv := objc.SendIfResponds[VZGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZGraphicsDeviceConfiguration._graphicsDevice]
//   - [VZGraphicsDeviceConfiguration._init]
//   - [VZGraphicsDeviceConfiguration._initWithConfiguration]
//   - [VZGraphicsDeviceConfiguration.MakeGraphicsDeviceForVirtualMachineGraphicsDeviceIndex]
//   - [VZGraphicsDeviceConfiguration.ValidateWithError]
//   - [VZGraphicsDeviceConfiguration.DebugDescription]
//   - [VZGraphicsDeviceConfiguration.Description]
//   - [VZGraphicsDeviceConfiguration.Hash]
//   - [VZGraphicsDeviceConfiguration.Superclass]
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
//   - [IVZGraphicsDeviceConfiguration._graphicsDevice]
//   - [IVZGraphicsDeviceConfiguration._init]
//   - [IVZGraphicsDeviceConfiguration._initWithConfiguration]
//   - [IVZGraphicsDeviceConfiguration.MakeGraphicsDeviceForVirtualMachineGraphicsDeviceIndex]
//   - [IVZGraphicsDeviceConfiguration.ValidateWithError]
//   - [IVZGraphicsDeviceConfiguration.DebugDescription]
//   - [IVZGraphicsDeviceConfiguration.Description]
//   - [IVZGraphicsDeviceConfiguration.Hash]
//   - [IVZGraphicsDeviceConfiguration.Superclass]
type IVZGraphicsDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_graphicsDevice() unsafe.Pointer
	_init() objectivec.IObject
	_initWithConfiguration(configuration unsafe.Pointer) objectivec.IObject
	MakeGraphicsDeviceForVirtualMachineGraphicsDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	ValidateWithError() (bool, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZGraphicsDeviceConfiguration) Init() VZGraphicsDeviceConfiguration {
	rv := objc.SendIfResponds[VZGraphicsDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZGraphicsDeviceConfiguration) Autorelease() VZGraphicsDeviceConfiguration {
	rv := objc.SendIfResponds[VZGraphicsDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDeviceConfiguration creates a new VZGraphicsDeviceConfiguration instance.
func NewVZGraphicsDeviceConfiguration() VZGraphicsDeviceConfiguration {
	class := getVZGraphicsDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZGraphicsDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZGraphicsDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZGraphicsDeviceConfiguration) _initWithConfiguration(configuration unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_initWithConfiguration:"), configuration)
	return objectivec.Object{ID: rv}
}

// InitWithConfiguration is an exported wrapper for the private method _initWithConfiguration.
func (v VZGraphicsDeviceConfiguration) InitWithConfiguration(configuration unsafe.Pointer) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initWithConfiguration:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithConfiguration:"}
		return nil, err
	}
	return v._initWithConfiguration(configuration), nil
}

// CanInitWithConfiguration reports whether the receiver responds to the private selector _initWithConfiguration:.
func (v VZGraphicsDeviceConfiguration) CanInitWithConfiguration() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initWithConfiguration:"))
}
func (v VZGraphicsDeviceConfiguration) MakeGraphicsDeviceForVirtualMachineGraphicsDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("makeGraphicsDeviceForVirtualMachine:graphicsDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}
func (v VZGraphicsDeviceConfiguration) ValidateWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("validateWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateWithError: returned NO with nil NSError")
	}
	return rv, nil

}

func (v VZGraphicsDeviceConfiguration) _graphicsDevice() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("_graphicsDevice"))
	return rv
}

// CanGraphicsDevice reports whether the receiver responds to the private selector _graphicsDevice.
func (v VZGraphicsDeviceConfiguration) CanGraphicsDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_graphicsDevice"))
}

// GraphicsDevice is an exported wrapper for the private property _graphicsDevice.
func (v VZGraphicsDeviceConfiguration) GraphicsDevice() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_graphicsDevice")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_graphicsDevice"}
	}
	return v._graphicsDevice(), nil
}
func (v VZGraphicsDeviceConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZGraphicsDeviceConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZGraphicsDeviceConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZGraphicsDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
