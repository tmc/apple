// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[VZVirtualMachineConfigurationDecoder](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtualMachineConfigurationDecoder.ConfigurationFromDataFormatError]
//   - [VZVirtualMachineConfigurationDecoder.Delegate]
//   - [VZVirtualMachineConfigurationDecoder.SetDelegate]
//   - [VZVirtualMachineConfigurationDecoder.InitWithBaseURL]
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
type IVZVirtualMachineConfigurationDecoder interface {
	objectivec.IObject

	// Topic: Methods

	ConfigurationFromDataFormatError(data objectivec.IObject, format *uint64) (objectivec.IObject, error)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	InitWithBaseURL(url foundation.NSURL) VZVirtualMachineConfigurationDecoder
}

// Init initializes the instance.
func (v VZVirtualMachineConfigurationDecoder) Init() VZVirtualMachineConfigurationDecoder {
	rv := objc.SendIfResponds[VZVirtualMachineConfigurationDecoder](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineConfigurationDecoder) Autorelease() VZVirtualMachineConfigurationDecoder {
	rv := objc.SendIfResponds[VZVirtualMachineConfigurationDecoder](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineConfigurationDecoder creates a new VZVirtualMachineConfigurationDecoder instance.
func NewVZVirtualMachineConfigurationDecoder() VZVirtualMachineConfigurationDecoder {
	class := getVZVirtualMachineConfigurationDecoderClass()
	rv := objc.SendIfResponds[VZVirtualMachineConfigurationDecoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZVirtualMachineConfigurationDecoderWithBaseURL(url foundation.NSURL) VZVirtualMachineConfigurationDecoder {
	instance := getVZVirtualMachineConfigurationDecoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBaseURL:"), url)
	return VZVirtualMachineConfigurationDecoderFromID(rv)
}

func (v VZVirtualMachineConfigurationDecoder) ConfigurationFromDataFormatError(data objectivec.IObject, format *uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("configurationFromData:format:error:"), data, unsafe.Pointer(format), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (v VZVirtualMachineConfigurationDecoder) InitWithBaseURL(url foundation.NSURL) VZVirtualMachineConfigurationDecoder {
	rv := objc.SendIfResponds[VZVirtualMachineConfigurationDecoder](v.ID, objc.Sel("initWithBaseURL:"), url)
	return rv
}

func (v VZVirtualMachineConfigurationDecoder) Delegate() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("delegate"))
	return rv
}
func (v VZVirtualMachineConfigurationDecoder) SetDelegate(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}
