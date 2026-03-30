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

// The class instance for the [VZCPUExitContext] class.
var (
	_VZCPUExitContextClass     VZCPUExitContextClass
	_VZCPUExitContextClassOnce sync.Once
)

func getVZCPUExitContextClass() VZCPUExitContextClass {
	_VZCPUExitContextClassOnce.Do(func() {
		_VZCPUExitContextClass = VZCPUExitContextClass{class: objc.GetClass("_VZCPUExitContext")}
	})
	return _VZCPUExitContextClass
}

// GetVZCPUExitContextClass returns the class object for _VZCPUExitContext.
func GetVZCPUExitContextClass() VZCPUExitContextClass {
	return getVZCPUExitContextClass()
}

type VZCPUExitContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCPUExitContextClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCPUExitContextClass) Alloc() VZCPUExitContext {
	rv := objc.Send[VZCPUExitContext](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCPUExitContext.CPUExit]
//   - [VZCPUExitContext.CPUIndex]
//   - [VZCPUExitContext.GetPhysicalAddressForVirtualAddressError]
//   - [VZCPUExitContext.GetRegisterValueError]
//   - [VZCPUExitContext.GetSIMDRegisterValueError]
//   - [VZCPUExitContext.GetSystemRegisterValueError]
//   - [VZCPUExitContext.GuestMemoryAtPhysicalAddressLengthError]
//   - [VZCPUExitContext.SetRegisterValueError]
//   - [VZCPUExitContext.SetSIMDRegisterValueError]
//   - [VZCPUExitContext.SetSystemRegisterValueError]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext
type VZCPUExitContext struct {
	objectivec.Object
}

// VZCPUExitContextFromID constructs a [VZCPUExitContext] from an objc.ID.
func VZCPUExitContextFromID(id objc.ID) VZCPUExitContext {
	return VZCPUExitContext{objectivec.Object{ID: id}}
}

// Ensure VZCPUExitContext implements IVZCPUExitContext.
var _ IVZCPUExitContext = VZCPUExitContext{}

// An interface definition for the [VZCPUExitContext] class.
//
// # Methods
//
//   - [IVZCPUExitContext.CPUExit]
//   - [IVZCPUExitContext.CPUIndex]
//   - [IVZCPUExitContext.GetPhysicalAddressForVirtualAddressError]
//   - [IVZCPUExitContext.GetRegisterValueError]
//   - [IVZCPUExitContext.GetSIMDRegisterValueError]
//   - [IVZCPUExitContext.GetSystemRegisterValueError]
//   - [IVZCPUExitContext.GuestMemoryAtPhysicalAddressLengthError]
//   - [IVZCPUExitContext.SetRegisterValueError]
//   - [IVZCPUExitContext.SetSIMDRegisterValueError]
//   - [IVZCPUExitContext.SetSystemRegisterValueError]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext
type IVZCPUExitContext interface {
	objectivec.IObject

	// Topic: Methods

	CPUExit() objectivec.IObject
	CPUIndex() uint64
	GetPhysicalAddressForVirtualAddressError(address2 uint64) (uint64, error)
	GetRegisterValueError(register uint32) (uint64, error)
	GetSIMDRegisterValueError(sIMDRegister uint32, value []objectivec.IObject) (bool, error)
	GetSystemRegisterValueError(register uint16) (uint64, error)
	GuestMemoryAtPhysicalAddressLengthError(address uint64, length uint64) (objectivec.IObject, error)
	SetRegisterValueError(register uint32, value uint64) (bool, error)
	SetSIMDRegisterValueError(sIMDRegister uint32, value objectivec.IObject) (bool, error)
	SetSystemRegisterValueError(register uint16, value uint64) (bool, error)
}

// Init initializes the instance.
func (v VZCPUExitContext) Init() VZCPUExitContext {
	rv := objc.Send[VZCPUExitContext](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCPUExitContext) Autorelease() VZCPUExitContext {
	rv := objc.Send[VZCPUExitContext](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCPUExitContext creates a new VZCPUExitContext instance.
func NewVZCPUExitContext() VZCPUExitContext {
	class := getVZCPUExitContextClass()
	rv := objc.Send[VZCPUExitContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext/getPhysicalAddress:forVirtualAddress:error:
func (v VZCPUExitContext) GetPhysicalAddressForVirtualAddressError(address2 uint64) (uint64, error) {
	var address uint64
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("getPhysicalAddress:forVirtualAddress:error:"), unsafe.Pointer(&address), address2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0, errors.New("getPhysicalAddress:forVirtualAddress:error: returned NO with nil NSError")
	}
	return address, nil
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext/getRegister:value:error:
func (v VZCPUExitContext) GetRegisterValueError(register uint32) (uint64, error) {
	var value uint64
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("getRegister:value:error:"), register, unsafe.Pointer(&value), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0, errors.New("getRegister:value:error: returned NO with nil NSError")
	}
	return value, nil
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext/getSIMDRegister:value:error:
func (v VZCPUExitContext) GetSIMDRegisterValueError(sIMDRegister uint32, value []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("getSIMDRegister:value:error:"), sIMDRegister, objectivec.IObjectSliceToNSArray(value), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getSIMDRegister:value:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext/getSystemRegister:value:error:
func (v VZCPUExitContext) GetSystemRegisterValueError(register uint16) (uint64, error) {
	var value uint64
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("getSystemRegister:value:error:"), register, unsafe.Pointer(&value), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0, errors.New("getSystemRegister:value:error: returned NO with nil NSError")
	}
	return value, nil
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext/guestMemoryAtPhysicalAddress:length:error:
func (v VZCPUExitContext) GuestMemoryAtPhysicalAddressLengthError(address uint64, length uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("guestMemoryAtPhysicalAddress:length:error:"), address, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext/setRegister:value:error:
func (v VZCPUExitContext) SetRegisterValueError(register uint32, value uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("setRegister:value:error:"), register, value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setRegister:value:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext/setSIMDRegister:value:error:
func (v VZCPUExitContext) SetSIMDRegisterValueError(sIMDRegister uint32, value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("setSIMDRegister:value:error:"), sIMDRegister, value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setSIMDRegister:value:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext/setSystemRegister:value:error:
func (v VZCPUExitContext) SetSystemRegisterValueError(register uint16, value uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("setSystemRegister:value:error:"), register, value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setSystemRegister:value:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext/CPUExit
func (v VZCPUExitContext) CPUExit() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("CPUExit"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCPUExitContext/CPUIndex
func (v VZCPUExitContext) CPUIndex() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("CPUIndex"))
	return rv
}
