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

// The class instance for the [VZMacGraphicsDisplay] class.
var (
	_VZMacGraphicsDisplayClass     VZMacGraphicsDisplayClass
	_VZMacGraphicsDisplayClassOnce sync.Once
)

func getVZMacGraphicsDisplayClass() VZMacGraphicsDisplayClass {
	_VZMacGraphicsDisplayClassOnce.Do(func() {
		_VZMacGraphicsDisplayClass = VZMacGraphicsDisplayClass{class: objc.GetClass("VZMacGraphicsDisplay")}
	})
	return _VZMacGraphicsDisplayClass
}

// GetVZMacGraphicsDisplayClass returns the class object for VZMacGraphicsDisplay.
func GetVZMacGraphicsDisplayClass() VZMacGraphicsDisplayClass {
	return getVZMacGraphicsDisplayClass()
}

type VZMacGraphicsDisplayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacGraphicsDisplayClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacGraphicsDisplayClass) Alloc() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacGraphicsDisplay._connectionType]
//   - [VZMacGraphicsDisplay._displayIdentifier]
//   - [VZMacGraphicsDisplay._displayMode]
//   - [VZMacGraphicsDisplay.ReconfigureWithConfigurationError]
//   - [VZMacGraphicsDisplay.InitWithConfigurationError]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay
type VZMacGraphicsDisplay struct {
	VZGraphicsDisplay
}

// VZMacGraphicsDisplayFromID constructs a [VZMacGraphicsDisplay] from an objc.ID.
func VZMacGraphicsDisplayFromID(id objc.ID) VZMacGraphicsDisplay {
	return VZMacGraphicsDisplay{VZGraphicsDisplay: VZGraphicsDisplayFromID(id)}
}

// Ensure VZMacGraphicsDisplay implements IVZMacGraphicsDisplay.
var _ IVZMacGraphicsDisplay = VZMacGraphicsDisplay{}

// An interface definition for the [VZMacGraphicsDisplay] class.
//
// # Methods
//
//   - [IVZMacGraphicsDisplay._connectionType]
//   - [IVZMacGraphicsDisplay._displayIdentifier]
//   - [IVZMacGraphicsDisplay._displayMode]
//   - [IVZMacGraphicsDisplay.ReconfigureWithConfigurationError]
//   - [IVZMacGraphicsDisplay.InitWithConfigurationError]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay
type IVZMacGraphicsDisplay interface {
	IVZGraphicsDisplay

	// Topic: Methods

	_connectionType() int64
	_displayIdentifier() objectivec.IObject
	_displayMode() int64
	ReconfigureWithConfigurationError(configuration objectivec.IObject) (bool, error)
	InitWithConfigurationError(configuration objectivec.IObject) (VZMacGraphicsDisplay, error)
}

// Init initializes the instance.
func (m VZMacGraphicsDisplay) Init() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacGraphicsDisplay) Autorelease() VZMacGraphicsDisplay {
	rv := objc.Send[VZMacGraphicsDisplay](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDisplay creates a new VZMacGraphicsDisplay instance.
func NewVZMacGraphicsDisplay() VZMacGraphicsDisplay {
	class := getVZMacGraphicsDisplayClass()
	rv := objc.Send[VZMacGraphicsDisplay](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay/initWithConfiguration:error:
func NewMacGraphicsDisplayWithConfigurationError(configuration objectivec.IObject) (VZMacGraphicsDisplay, error) {
	var errorPtr objc.ID
	instance := getVZMacGraphicsDisplayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZMacGraphicsDisplay{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZMacGraphicsDisplayFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:
func NewMacGraphicsDisplayWithVirtualMachineGraphicsDeviceIndexFramebufferIndexUuid(machine objectivec.IObject, index uint64, index2 uint64, uuid objectivec.IObject) VZMacGraphicsDisplay {
	instance := getVZMacGraphicsDisplayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:"), machine, index, index2, uuid)
	return VZMacGraphicsDisplayFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay/_connectionType
func (m VZMacGraphicsDisplay) _connectionType() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("_connectionType"))
	return rv
}

// ConnectionType is an exported wrapper for the private method _connectionType.
func (m VZMacGraphicsDisplay) ConnectionType() int64 {
	return m._connectionType()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay/_displayIdentifier
func (m VZMacGraphicsDisplay) _displayIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_displayIdentifier"))
	return objectivec.Object{ID: rv}
}

// DisplayIdentifier is an exported wrapper for the private method _displayIdentifier.
func (m VZMacGraphicsDisplay) DisplayIdentifier() objectivec.IObject {
	return m._displayIdentifier()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay/_displayMode
func (m VZMacGraphicsDisplay) _displayMode() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("_displayMode"))
	return rv
}

// DisplayMode is an exported wrapper for the private method _displayMode.
func (m VZMacGraphicsDisplay) DisplayMode() int64 {
	return m._displayMode()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay/reconfigureWithConfiguration:error:
func (m VZMacGraphicsDisplay) ReconfigureWithConfigurationError(configuration objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("reconfigureWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("reconfigureWithConfiguration:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplay/initWithConfiguration:error:
func (m VZMacGraphicsDisplay) InitWithConfigurationError(configuration objectivec.IObject) (VZMacGraphicsDisplay, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZMacGraphicsDisplay{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZMacGraphicsDisplayFromID(rv), nil

}
