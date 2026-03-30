// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

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
	rv := objc.Send[VZMacHardwareModelDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacHardwareModelDescriptor
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
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacHardwareModelDescriptor
type IVZMacHardwareModelDescriptor interface {
	objectivec.IObject

	// Topic: Methods

	SetBoardID(id uint32)
	SetISA(isa int64)
	SetInitialGuestMacOSVersion(oSVersion objectivec.IObject)
	SetMinimumSupportedHostOSVersion(oSVersion objectivec.IObject)
	SetPlatformVersion(version uint32)
	SetVariantIDVariantName(id uint32, name objectivec.IObject)
}

// Init initializes the instance.
func (v VZMacHardwareModelDescriptor) Init() VZMacHardwareModelDescriptor {
	rv := objc.Send[VZMacHardwareModelDescriptor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacHardwareModelDescriptor) Autorelease() VZMacHardwareModelDescriptor {
	rv := objc.Send[VZMacHardwareModelDescriptor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacHardwareModelDescriptor creates a new VZMacHardwareModelDescriptor instance.
func NewVZMacHardwareModelDescriptor() VZMacHardwareModelDescriptor {
	class := getVZMacHardwareModelDescriptorClass()
	rv := objc.Send[VZMacHardwareModelDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacHardwareModelDescriptor/setBoardID:
func (v VZMacHardwareModelDescriptor) SetBoardID(id uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("setBoardID:"), id)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacHardwareModelDescriptor/setISA:
func (v VZMacHardwareModelDescriptor) SetISA(isa int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("setISA:"), isa)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacHardwareModelDescriptor/setInitialGuestMacOSVersion:
func (v VZMacHardwareModelDescriptor) SetInitialGuestMacOSVersion(oSVersion objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("setInitialGuestMacOSVersion:"), oSVersion)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacHardwareModelDescriptor/setMinimumSupportedHostOSVersion:
func (v VZMacHardwareModelDescriptor) SetMinimumSupportedHostOSVersion(oSVersion objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("setMinimumSupportedHostOSVersion:"), oSVersion)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacHardwareModelDescriptor/setPlatformVersion:
func (v VZMacHardwareModelDescriptor) SetPlatformVersion(version uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("setPlatformVersion:"), version)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacHardwareModelDescriptor/setVariantID:variantName:
func (v VZMacHardwareModelDescriptor) SetVariantIDVariantName(id uint32, name objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("setVariantID:variantName:"), id, name)
}
