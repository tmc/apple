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

// The class instance for the [VZVirtioGraphicsScanout] class.
var (
	_VZVirtioGraphicsScanoutClass     VZVirtioGraphicsScanoutClass
	_VZVirtioGraphicsScanoutClassOnce sync.Once
)

func getVZVirtioGraphicsScanoutClass() VZVirtioGraphicsScanoutClass {
	_VZVirtioGraphicsScanoutClassOnce.Do(func() {
		_VZVirtioGraphicsScanoutClass = VZVirtioGraphicsScanoutClass{class: objc.GetClass("VZVirtioGraphicsScanout")}
	})
	return _VZVirtioGraphicsScanoutClass
}

// GetVZVirtioGraphicsScanoutClass returns the class object for VZVirtioGraphicsScanout.
func GetVZVirtioGraphicsScanoutClass() VZVirtioGraphicsScanoutClass {
	return getVZVirtioGraphicsScanoutClass()
}

type VZVirtioGraphicsScanoutClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioGraphicsScanoutClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioGraphicsScanoutClass) Alloc() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtioGraphicsScanout.ReconfigureWithConfigurationError]
//   - [VZVirtioGraphicsScanout.InitWithConfigurationError]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanout
type VZVirtioGraphicsScanout struct {
	VZGraphicsDisplay
}

// VZVirtioGraphicsScanoutFromID constructs a [VZVirtioGraphicsScanout] from an objc.ID.
func VZVirtioGraphicsScanoutFromID(id objc.ID) VZVirtioGraphicsScanout {
	return VZVirtioGraphicsScanout{VZGraphicsDisplay: VZGraphicsDisplayFromID(id)}
}

// Ensure VZVirtioGraphicsScanout implements IVZVirtioGraphicsScanout.
var _ IVZVirtioGraphicsScanout = VZVirtioGraphicsScanout{}

// An interface definition for the [VZVirtioGraphicsScanout] class.
//
// # Methods
//
//   - [IVZVirtioGraphicsScanout.ReconfigureWithConfigurationError]
//   - [IVZVirtioGraphicsScanout.InitWithConfigurationError]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanout
type IVZVirtioGraphicsScanout interface {
	IVZGraphicsDisplay

	// Topic: Methods

	ReconfigureWithConfigurationError(configuration objectivec.IObject) (bool, error)
	InitWithConfigurationError(configuration objectivec.IObject) (VZVirtioGraphicsScanout, error)
}

// Init initializes the instance.
func (v VZVirtioGraphicsScanout) Init() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioGraphicsScanout) Autorelease() VZVirtioGraphicsScanout {
	rv := objc.Send[VZVirtioGraphicsScanout](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsScanout creates a new VZVirtioGraphicsScanout instance.
func NewVZVirtioGraphicsScanout() VZVirtioGraphicsScanout {
	class := getVZVirtioGraphicsScanoutClass()
	rv := objc.Send[VZVirtioGraphicsScanout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanout/initWithConfiguration:error:
func NewVirtioGraphicsScanoutWithConfigurationError(configuration objectivec.IObject) (VZVirtioGraphicsScanout, error) {
	var errorPtr objc.ID
	instance := getVZVirtioGraphicsScanoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZVirtioGraphicsScanout{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZVirtioGraphicsScanoutFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:
func NewVirtioGraphicsScanoutWithVirtualMachineGraphicsDeviceIndexFramebufferIndexUuid(machine objectivec.IObject, index uint64, index2 uint64, uuid objectivec.IObject) VZVirtioGraphicsScanout {
	instance := getVZVirtioGraphicsScanoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:"), machine, index, index2, uuid)
	return VZVirtioGraphicsScanoutFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanout/reconfigureWithConfiguration:error:
func (v VZVirtioGraphicsScanout) ReconfigureWithConfigurationError(configuration objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("reconfigureWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("reconfigureWithConfiguration:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanout/initWithConfiguration:error:
func (v VZVirtioGraphicsScanout) InitWithConfigurationError(configuration objectivec.IObject) (VZVirtioGraphicsScanout, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZVirtioGraphicsScanout{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZVirtioGraphicsScanoutFromID(rv), nil

}
