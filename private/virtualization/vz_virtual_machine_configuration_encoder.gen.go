// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachineConfigurationEncoder] class.
var (
	_VZVirtualMachineConfigurationEncoderClass     VZVirtualMachineConfigurationEncoderClass
	_VZVirtualMachineConfigurationEncoderClassOnce sync.Once
)

func getVZVirtualMachineConfigurationEncoderClass() VZVirtualMachineConfigurationEncoderClass {
	_VZVirtualMachineConfigurationEncoderClassOnce.Do(func() {
		_VZVirtualMachineConfigurationEncoderClass = VZVirtualMachineConfigurationEncoderClass{class: objc.GetClass("_VZVirtualMachineConfigurationEncoder")}
	})
	return _VZVirtualMachineConfigurationEncoderClass
}

// GetVZVirtualMachineConfigurationEncoderClass returns the class object for _VZVirtualMachineConfigurationEncoder.
func GetVZVirtualMachineConfigurationEncoderClass() VZVirtualMachineConfigurationEncoderClass {
	return getVZVirtualMachineConfigurationEncoderClass()
}

type VZVirtualMachineConfigurationEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtualMachineConfigurationEncoderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineConfigurationEncoderClass) Alloc() VZVirtualMachineConfigurationEncoder {
	rv := objc.Send[VZVirtualMachineConfigurationEncoder](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtualMachineConfigurationEncoder.DataWithConfigurationFormatError]
//   - [VZVirtualMachineConfigurationEncoder.InitWithBaseURL]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationEncoder
type VZVirtualMachineConfigurationEncoder struct {
	objectivec.Object
}

// VZVirtualMachineConfigurationEncoderFromID constructs a [VZVirtualMachineConfigurationEncoder] from an objc.ID.
func VZVirtualMachineConfigurationEncoderFromID(id objc.ID) VZVirtualMachineConfigurationEncoder {
	return VZVirtualMachineConfigurationEncoder{objectivec.Object{ID: id}}
}

// Ensure VZVirtualMachineConfigurationEncoder implements IVZVirtualMachineConfigurationEncoder.
var _ IVZVirtualMachineConfigurationEncoder = VZVirtualMachineConfigurationEncoder{}

// An interface definition for the [VZVirtualMachineConfigurationEncoder] class.
//
// # Methods
//
//   - [IVZVirtualMachineConfigurationEncoder.DataWithConfigurationFormatError]
//   - [IVZVirtualMachineConfigurationEncoder.InitWithBaseURL]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationEncoder
type IVZVirtualMachineConfigurationEncoder interface {
	objectivec.IObject

	// Topic: Methods

	DataWithConfigurationFormatError(configuration objectivec.IObject, format uint64) (objectivec.IObject, error)
	InitWithBaseURL(url foundation.INSURL) VZVirtualMachineConfigurationEncoder
}

// Init initializes the instance.
func (v VZVirtualMachineConfigurationEncoder) Init() VZVirtualMachineConfigurationEncoder {
	rv := objc.Send[VZVirtualMachineConfigurationEncoder](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineConfigurationEncoder) Autorelease() VZVirtualMachineConfigurationEncoder {
	rv := objc.Send[VZVirtualMachineConfigurationEncoder](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineConfigurationEncoder creates a new VZVirtualMachineConfigurationEncoder instance.
func NewVZVirtualMachineConfigurationEncoder() VZVirtualMachineConfigurationEncoder {
	class := getVZVirtualMachineConfigurationEncoderClass()
	rv := objc.Send[VZVirtualMachineConfigurationEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationEncoder/initWithBaseURL:
func NewVZVirtualMachineConfigurationEncoderWithBaseURL(url foundation.INSURL) VZVirtualMachineConfigurationEncoder {
	instance := getVZVirtualMachineConfigurationEncoderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBaseURL:"), url)
	return VZVirtualMachineConfigurationEncoderFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationEncoder/dataWithConfiguration:format:error:
func (v VZVirtualMachineConfigurationEncoder) DataWithConfigurationFormatError(configuration objectivec.IObject, format uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("dataWithConfiguration:format:error:"), configuration, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationEncoder/initWithBaseURL:
func (v VZVirtualMachineConfigurationEncoder) InitWithBaseURL(url foundation.INSURL) VZVirtualMachineConfigurationEncoder {
	rv := objc.Send[VZVirtualMachineConfigurationEncoder](v.ID, objc.Sel("initWithBaseURL:"), url)
	return rv
}
