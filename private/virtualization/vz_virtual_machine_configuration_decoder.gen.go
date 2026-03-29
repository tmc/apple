// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachineConfigurationDecoder] class.
var (
	_VZVirtualMachineConfigurationDecoderClass     VZVirtualMachineConfigurationDecoderClass
	_VZVirtualMachineConfigurationDecoderClassOnce sync.Once
)

func getVZVirtualMachineConfigurationDecoderClass() VZVirtualMachineConfigurationDecoderClass {
	_VZVirtualMachineConfigurationDecoderClassOnce.Do(func() {
		_VZVirtualMachineConfigurationDecoderClass = VZVirtualMachineConfigurationDecoderClass{class: objc.GetClass("_VZVirtualMachineConfigurationDecoder")}
	})
	return _VZVirtualMachineConfigurationDecoderClass
}

// GetVZVirtualMachineConfigurationDecoderClass returns the class object for _VZVirtualMachineConfigurationDecoder.
func GetVZVirtualMachineConfigurationDecoderClass() VZVirtualMachineConfigurationDecoderClass {
	return getVZVirtualMachineConfigurationDecoderClass()
}

type VZVirtualMachineConfigurationDecoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtualMachineConfigurationDecoderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineConfigurationDecoderClass) Alloc() VZVirtualMachineConfigurationDecoder {
	rv := objc.Send[VZVirtualMachineConfigurationDecoder](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZVirtualMachineConfigurationDecoder.ConfigurationFromDataFormatError]
//   - [VZVirtualMachineConfigurationDecoder.Delegate]
//   - [VZVirtualMachineConfigurationDecoder.SetDelegate]
//   - [VZVirtualMachineConfigurationDecoder.InitWithBaseURL]
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationDecoder
type VZVirtualMachineConfigurationDecoder struct {
	objectivec.Object
}

// VZVirtualMachineConfigurationDecoderFromID constructs a [VZVirtualMachineConfigurationDecoder] from an objc.ID.
func VZVirtualMachineConfigurationDecoderFromID(id objc.ID) VZVirtualMachineConfigurationDecoder {
	return VZVirtualMachineConfigurationDecoder{objectivec.Object{ID: id}}
}
// Ensure VZVirtualMachineConfigurationDecoder implements IVZVirtualMachineConfigurationDecoder.
var _ IVZVirtualMachineConfigurationDecoder = VZVirtualMachineConfigurationDecoder{}

// An interface definition for the [VZVirtualMachineConfigurationDecoder] class.
//
// # Methods
//
//   - [IVZVirtualMachineConfigurationDecoder.ConfigurationFromDataFormatError]
//   - [IVZVirtualMachineConfigurationDecoder.Delegate]
//   - [IVZVirtualMachineConfigurationDecoder.SetDelegate]
//   - [IVZVirtualMachineConfigurationDecoder.InitWithBaseURL]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationDecoder
type IVZVirtualMachineConfigurationDecoder interface {
	objectivec.IObject

	// Topic: Methods

	ConfigurationFromDataFormatError(data objectivec.IObject, format unsafe.Pointer) (objectivec.IObject, error)
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	InitWithBaseURL(url foundation.INSURL) VZVirtualMachineConfigurationDecoder
}

// Init initializes the instance.
func (v VZVirtualMachineConfigurationDecoder) Init() VZVirtualMachineConfigurationDecoder {
	rv := objc.Send[VZVirtualMachineConfigurationDecoder](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineConfigurationDecoder) Autorelease() VZVirtualMachineConfigurationDecoder {
	rv := objc.Send[VZVirtualMachineConfigurationDecoder](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineConfigurationDecoder creates a new VZVirtualMachineConfigurationDecoder instance.
func NewVZVirtualMachineConfigurationDecoder() VZVirtualMachineConfigurationDecoder {
	class := getVZVirtualMachineConfigurationDecoderClass()
	rv := objc.Send[VZVirtualMachineConfigurationDecoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationDecoder/initWithBaseURL:
func NewVZVirtualMachineConfigurationDecoderWithBaseURL(url foundation.INSURL) VZVirtualMachineConfigurationDecoder {
	instance := getVZVirtualMachineConfigurationDecoderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBaseURL:"), url)
	return VZVirtualMachineConfigurationDecoderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationDecoder/configurationFromData:format:error:
func (v VZVirtualMachineConfigurationDecoder) ConfigurationFromDataFormatError(data objectivec.IObject, format unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("configurationFromData:format:error:"), data, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationDecoder/initWithBaseURL:
func (v VZVirtualMachineConfigurationDecoder) InitWithBaseURL(url foundation.INSURL) VZVirtualMachineConfigurationDecoder {
	rv := objc.Send[VZVirtualMachineConfigurationDecoder](v.ID, objc.Sel("initWithBaseURL:"), url)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationDecoder/delegate
func (v VZVirtualMachineConfigurationDecoder) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (v VZVirtualMachineConfigurationDecoder) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}

