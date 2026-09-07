// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacHardwareModelDescriptor] class.
var (
	_VZMacHardwareModelDescriptorClass     VZMacHardwareModelDescriptorClass
	_VZMacHardwareModelDescriptorClassOnce sync.Once
)

func getVZMacHardwareModelDescriptorClass() VZMacHardwareModelDescriptorClass {
	_VZMacHardwareModelDescriptorClassOnce.Do(func() {
		_VZMacHardwareModelDescriptorClass = VZMacHardwareModelDescriptorClass{class: objc.GetClass("_VZMacHardwareModelDescriptor")}
	})
	return _VZMacHardwareModelDescriptorClass
}

// GetVZMacHardwareModelDescriptorClass returns the class object for _VZMacHardwareModelDescriptor.
func GetVZMacHardwareModelDescriptorClass() VZMacHardwareModelDescriptorClass {
	return getVZMacHardwareModelDescriptorClass()
}

type VZMacHardwareModelDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacHardwareModelDescriptorClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacHardwareModelDescriptorClass) Alloc() VZMacHardwareModelDescriptor {
	rv := objc.SendIfResponds[VZMacHardwareModelDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacHardwareModelDescriptor.SetBoardID]
//   - [VZMacHardwareModelDescriptor.SetISA]
//   - [VZMacHardwareModelDescriptor.SetInitialGuestMacOSVersion]
//   - [VZMacHardwareModelDescriptor.SetMinimumSupportedHostOSVersion]
//   - [VZMacHardwareModelDescriptor.SetPlatformVersion]
//   - [VZMacHardwareModelDescriptor.SetVariantIDVariantName]
type VZMacHardwareModelDescriptor struct {
	objectivec.Object
}

// VZMacHardwareModelDescriptorFromID constructs a [VZMacHardwareModelDescriptor] from an objc.ID.
func VZMacHardwareModelDescriptorFromID(id objc.ID) VZMacHardwareModelDescriptor {
	return VZMacHardwareModelDescriptor{objectivec.Object{ID: id}}
}

// Ensure VZMacHardwareModelDescriptor implements IVZMacHardwareModelDescriptor.
var _ IVZMacHardwareModelDescriptor = VZMacHardwareModelDescriptor{}

// An interface definition for the [VZMacHardwareModelDescriptor] class.
//
// # Methods
//
//   - [IVZMacHardwareModelDescriptor.SetBoardID]
//   - [IVZMacHardwareModelDescriptor.SetISA]
//   - [IVZMacHardwareModelDescriptor.SetInitialGuestMacOSVersion]
//   - [IVZMacHardwareModelDescriptor.SetMinimumSupportedHostOSVersion]
//   - [IVZMacHardwareModelDescriptor.SetPlatformVersion]
//   - [IVZMacHardwareModelDescriptor.SetVariantIDVariantName]
type IVZMacHardwareModelDescriptor interface {
	objectivec.IObject

	// Topic: Methods

	SetBoardID(id uint32)
	SetISA(isa int64)
	SetInitialGuestMacOSVersion(oSVersion unsafe.Pointer)
	SetMinimumSupportedHostOSVersion(oSVersion unsafe.Pointer)
	SetPlatformVersion(version uint32)
	SetVariantIDVariantName(id uint32, name objectivec.IObject)
}

// Init initializes the instance.
func (v VZMacHardwareModelDescriptor) Init() VZMacHardwareModelDescriptor {
	rv := objc.SendIfResponds[VZMacHardwareModelDescriptor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacHardwareModelDescriptor) Autorelease() VZMacHardwareModelDescriptor {
	rv := objc.SendIfResponds[VZMacHardwareModelDescriptor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacHardwareModelDescriptor creates a new VZMacHardwareModelDescriptor instance.
func NewVZMacHardwareModelDescriptor() VZMacHardwareModelDescriptor {
	class := getVZMacHardwareModelDescriptorClass()
	rv := objc.SendIfResponds[VZMacHardwareModelDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZMacHardwareModelDescriptor) SetBoardID(id uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("setBoardID:"), id)
}
func (v VZMacHardwareModelDescriptor) SetISA(isa int64) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("setISA:"), isa)
}
func (v VZMacHardwareModelDescriptor) SetInitialGuestMacOSVersion(oSVersion unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("setInitialGuestMacOSVersion:"), oSVersion)
}
func (v VZMacHardwareModelDescriptor) SetMinimumSupportedHostOSVersion(oSVersion unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("setMinimumSupportedHostOSVersion:"), oSVersion)
}
func (v VZMacHardwareModelDescriptor) SetPlatformVersion(version uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("setPlatformVersion:"), version)
}
func (v VZMacHardwareModelDescriptor) SetVariantIDVariantName(id uint32, name objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("setVariantID:variantName:"), id, name)
}
