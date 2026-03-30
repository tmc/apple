// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacMachineIdentifier] class.
var (
	_VZMacMachineIdentifierClass     VZMacMachineIdentifierClass
	_VZMacMachineIdentifierClassOnce sync.Once
)

func getVZMacMachineIdentifierClass() VZMacMachineIdentifierClass {
	_VZMacMachineIdentifierClassOnce.Do(func() {
		_VZMacMachineIdentifierClass = VZMacMachineIdentifierClass{class: objc.GetClass("VZMacMachineIdentifier")}
	})
	return _VZMacMachineIdentifierClass
}

// GetVZMacMachineIdentifierClass returns the class object for VZMacMachineIdentifier.
func GetVZMacMachineIdentifierClass() VZMacMachineIdentifierClass {
	return getVZMacMachineIdentifierClass()
}

type VZMacMachineIdentifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacMachineIdentifierClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacMachineIdentifierClass) Alloc() VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacMachineIdentifier._ECID]
//   - [VZMacMachineIdentifier._ECIDChecksDisabled]
//   - [VZMacMachineIdentifier._serialNumber]
//   - [VZMacMachineIdentifier.DebugDescription]
//   - [VZMacMachineIdentifier.Description]
//   - [VZMacMachineIdentifier.Hash]
//   - [VZMacMachineIdentifier.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier
type VZMacMachineIdentifier struct {
	objectivec.Object
}

// VZMacMachineIdentifierFromID constructs a [VZMacMachineIdentifier] from an objc.ID.
func VZMacMachineIdentifierFromID(id objc.ID) VZMacMachineIdentifier {
	return VZMacMachineIdentifier{objectivec.Object{ID: id}}
}

// Ensure VZMacMachineIdentifier implements IVZMacMachineIdentifier.
var _ IVZMacMachineIdentifier = VZMacMachineIdentifier{}

// An interface definition for the [VZMacMachineIdentifier] class.
//
// # Methods
//
//   - [IVZMacMachineIdentifier._ECID]
//   - [IVZMacMachineIdentifier._ECIDChecksDisabled]
//   - [IVZMacMachineIdentifier._serialNumber]
//   - [IVZMacMachineIdentifier.DebugDescription]
//   - [IVZMacMachineIdentifier.Description]
//   - [IVZMacMachineIdentifier.Hash]
//   - [IVZMacMachineIdentifier.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier
type IVZMacMachineIdentifier interface {
	objectivec.IObject

	// Topic: Methods

	_ECID() uint64
	_ECIDChecksDisabled() bool
	_serialNumber() *VZMacSerialNumber
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (m VZMacMachineIdentifier) Init() VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacMachineIdentifier) Autorelease() VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacMachineIdentifier creates a new VZMacMachineIdentifier instance.
func NewVZMacMachineIdentifier() VZMacMachineIdentifier {
	class := getVZMacMachineIdentifierClass()
	rv := objc.Send[VZMacMachineIdentifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/_machineIdentifierForVirtualMachineClone
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) _machineIdentifierForVirtualMachineClone() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineClone"))
	return objectivec.Object{ID: rv}
}

// MachineIdentifierForVirtualMachineClone is an exported wrapper for the private method _machineIdentifierForVirtualMachineClone.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) MachineIdentifierForVirtualMachineClone() objectivec.IObject {
	return _VZMacMachineIdentifierClass._machineIdentifierForVirtualMachineClone()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/_machineIdentifierForVirtualMachineCloneWithECID:serialNumber:
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) _machineIdentifierForVirtualMachineCloneWithECIDSerialNumber(ecid uint64, number objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineCloneWithECID:serialNumber:"), ecid, number)
	return objectivec.Object{ID: rv}
}

// MachineIdentifierForVirtualMachineCloneWithECIDSerialNumber is an exported wrapper for the private method _machineIdentifierForVirtualMachineCloneWithECIDSerialNumber.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) MachineIdentifierForVirtualMachineCloneWithECIDSerialNumber(ecid uint64, number objectivec.IObject) objectivec.IObject {
	return _VZMacMachineIdentifierClass._machineIdentifierForVirtualMachineCloneWithECIDSerialNumber(ecid, number)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/_machineIdentifierForVirtualMachineCloneWithSerialNumber:
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) _machineIdentifierForVirtualMachineCloneWithSerialNumber(number objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineCloneWithSerialNumber:"), number)
	return objectivec.Object{ID: rv}
}

// MachineIdentifierForVirtualMachineCloneWithSerialNumber is an exported wrapper for the private method _machineIdentifierForVirtualMachineCloneWithSerialNumber.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) MachineIdentifierForVirtualMachineCloneWithSerialNumber(number objectivec.IObject) objectivec.IObject {
	return _VZMacMachineIdentifierClass._machineIdentifierForVirtualMachineCloneWithSerialNumber(number)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/_machineIdentifierWithECID:serialNumber:
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) _machineIdentifierWithECIDSerialNumber(ecid uint64, number objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierWithECID:serialNumber:"), ecid, number)
	return objectivec.Object{ID: rv}
}

// MachineIdentifierWithECIDSerialNumber is an exported wrapper for the private method _machineIdentifierWithECIDSerialNumber.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) MachineIdentifierWithECIDSerialNumber(ecid uint64, number objectivec.IObject) objectivec.IObject {
	return _VZMacMachineIdentifierClass._machineIdentifierWithECIDSerialNumber(ecid, number)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/_machineIdentifierWithSerialNumber:
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) _machineIdentifierWithSerialNumber(number objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierWithSerialNumber:"), number)
	return objectivec.Object{ID: rv}
}

// MachineIdentifierWithSerialNumber is an exported wrapper for the private method _machineIdentifierWithSerialNumber.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) MachineIdentifierWithSerialNumber(number objectivec.IObject) objectivec.IObject {
	return _VZMacMachineIdentifierClass._machineIdentifierWithSerialNumber(number)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/_ECID
func (m VZMacMachineIdentifier) _ECID() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("_ECID"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/_ECIDChecksDisabled
func (m VZMacMachineIdentifier) _ECIDChecksDisabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_ECIDChecksDisabled"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/_serialNumber
func (m VZMacMachineIdentifier) _serialNumber() *VZMacSerialNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_serialNumber"))
	if rv == 0 {
		return nil
	}
	val := VZMacSerialNumberFromID(objc.ID(rv))
	return &val
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/debugDescription
func (m VZMacMachineIdentifier) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/description
func (m VZMacMachineIdentifier) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/hash
func (m VZMacMachineIdentifier) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/superclass
func (m VZMacMachineIdentifier) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}
