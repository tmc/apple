// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZLinearFramebufferGraphicsDisplay] class.
var (
	_VZLinearFramebufferGraphicsDisplayClass     VZLinearFramebufferGraphicsDisplayClass
	_VZLinearFramebufferGraphicsDisplayClassOnce sync.Once
)

func getVZLinearFramebufferGraphicsDisplayClass() VZLinearFramebufferGraphicsDisplayClass {
	_VZLinearFramebufferGraphicsDisplayClassOnce.Do(func() {
		_VZLinearFramebufferGraphicsDisplayClass = VZLinearFramebufferGraphicsDisplayClass{class: objc.GetClass("_VZLinearFramebufferGraphicsDisplay")}
	})
	return _VZLinearFramebufferGraphicsDisplayClass
}

// GetVZLinearFramebufferGraphicsDisplayClass returns the class object for _VZLinearFramebufferGraphicsDisplay.
func GetVZLinearFramebufferGraphicsDisplayClass() VZLinearFramebufferGraphicsDisplayClass {
	return getVZLinearFramebufferGraphicsDisplayClass()
}

type VZLinearFramebufferGraphicsDisplayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZLinearFramebufferGraphicsDisplayClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinearFramebufferGraphicsDisplayClass) Alloc() VZLinearFramebufferGraphicsDisplay {
	rv := objc.Send[VZLinearFramebufferGraphicsDisplay](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZLinearFramebufferGraphicsDisplay._displayConfiguration]
//   - [VZLinearFramebufferGraphicsDisplay.ReconfigureWithConfigurationError]
//   - [VZLinearFramebufferGraphicsDisplay.ReconfigureWithSizeInPixelsError]
//   - [VZLinearFramebufferGraphicsDisplay.SizeInPixels]
// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDisplay
type VZLinearFramebufferGraphicsDisplay struct {
	VZGraphicsDisplay
}

// VZLinearFramebufferGraphicsDisplayFromID constructs a [VZLinearFramebufferGraphicsDisplay] from an objc.ID.
func VZLinearFramebufferGraphicsDisplayFromID(id objc.ID) VZLinearFramebufferGraphicsDisplay {
	return VZLinearFramebufferGraphicsDisplay{VZGraphicsDisplay: VZGraphicsDisplayFromID(id)}
}
// Ensure VZLinearFramebufferGraphicsDisplay implements IVZLinearFramebufferGraphicsDisplay.
var _ IVZLinearFramebufferGraphicsDisplay = VZLinearFramebufferGraphicsDisplay{}

// An interface definition for the [VZLinearFramebufferGraphicsDisplay] class.
//
// # Methods
//
//   - [IVZLinearFramebufferGraphicsDisplay._displayConfiguration]
//   - [IVZLinearFramebufferGraphicsDisplay.ReconfigureWithConfigurationError]
//   - [IVZLinearFramebufferGraphicsDisplay.ReconfigureWithSizeInPixelsError]
//   - [IVZLinearFramebufferGraphicsDisplay.SizeInPixels]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDisplay
type IVZLinearFramebufferGraphicsDisplay interface {
	IVZGraphicsDisplay

	// Topic: Methods

	_displayConfiguration() objectivec.IObject
	ReconfigureWithConfigurationError(configuration objectivec.IObject) (bool, error)
	ReconfigureWithSizeInPixelsError(pixels corefoundation.CGSize) (bool, error)
	SizeInPixels() corefoundation.CGSize
}

// Init initializes the instance.
func (v VZLinearFramebufferGraphicsDisplay) Init() VZLinearFramebufferGraphicsDisplay {
	rv := objc.Send[VZLinearFramebufferGraphicsDisplay](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZLinearFramebufferGraphicsDisplay) Autorelease() VZLinearFramebufferGraphicsDisplay {
	rv := objc.Send[VZLinearFramebufferGraphicsDisplay](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinearFramebufferGraphicsDisplay creates a new VZLinearFramebufferGraphicsDisplay instance.
func NewVZLinearFramebufferGraphicsDisplay() VZLinearFramebufferGraphicsDisplay {
	class := getVZLinearFramebufferGraphicsDisplayClass()
	rv := objc.Send[VZLinearFramebufferGraphicsDisplay](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:
func NewVZLinearFramebufferGraphicsDisplayWithVirtualMachineGraphicsDeviceIndexFramebufferIndexUuid(machine objectivec.IObject, index uint64, index2 uint64, uuid objectivec.IObject) VZLinearFramebufferGraphicsDisplay {
	instance := getVZLinearFramebufferGraphicsDisplayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:"), machine, index, index2, uuid)
	return VZLinearFramebufferGraphicsDisplayFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDisplay/_displayConfiguration
func (v VZLinearFramebufferGraphicsDisplay) _displayConfiguration() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_displayConfiguration"))
	return objectivec.Object{ID: rv}
}

// DisplayConfiguration is an exported wrapper for the private method _displayConfiguration.
func (v VZLinearFramebufferGraphicsDisplay) DisplayConfiguration() objectivec.IObject {
	return v._displayConfiguration()
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDisplay/reconfigureWithConfiguration:error:
func (v VZLinearFramebufferGraphicsDisplay) ReconfigureWithConfigurationError(configuration objectivec.IObject) (bool, error) {
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
//
// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDisplay/reconfigureWithSizeInPixels:error:
func (v VZLinearFramebufferGraphicsDisplay) ReconfigureWithSizeInPixelsError(pixels corefoundation.CGSize) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("reconfigureWithSizeInPixels:error:"), pixels, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("reconfigureWithSizeInPixels:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDisplay/sizeInPixels
func (v VZLinearFramebufferGraphicsDisplay) SizeInPixels() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("sizeInPixels"))
	return corefoundation.CGSize(rv)
}

