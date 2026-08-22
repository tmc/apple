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
	rv := objc.SendIfResponds[VZMacMachineIdentifier](objc.ID(vc.class), objc.Sel("alloc"))
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
type IVZMacMachineIdentifier interface {
	objectivec.IObject

	// Topic: Methods

	_ECID() uint64
	_ECIDChecksDisabled() bool
	_serialNumber() IVZMacSerialNumber
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZMacMachineIdentifier) Init() VZMacMachineIdentifier {
	rv := objc.SendIfResponds[VZMacMachineIdentifier](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacMachineIdentifier) Autorelease() VZMacMachineIdentifier {
	rv := objc.SendIfResponds[VZMacMachineIdentifier](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacMachineIdentifier creates a new VZMacMachineIdentifier instance.
func NewVZMacMachineIdentifier() VZMacMachineIdentifier {
	class := getVZMacMachineIdentifierClass()
	rv := objc.SendIfResponds[VZMacMachineIdentifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) _machineIdentifierForVirtualMachineClone() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineClone"))
	return objectivec.Object{ID: rv}
}

// MachineIdentifierForVirtualMachineClone is an exported wrapper for the private method _machineIdentifierForVirtualMachineClone.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) MachineIdentifierForVirtualMachineClone() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineClone")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_machineIdentifierForVirtualMachineClone"}
		return nil, err
	}
	return _VZMacMachineIdentifierClass._machineIdentifierForVirtualMachineClone(), nil
}

// CanMachineIdentifierForVirtualMachineClone reports whether the receiver responds to the private selector _machineIdentifierForVirtualMachineClone.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) CanMachineIdentifierForVirtualMachineClone() bool {
	return objc.RespondsToSelector(objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineClone"))
}
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) _machineIdentifierForVirtualMachineCloneWithECIDSerialNumber(ecid uint64, number objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineCloneWithECID:serialNumber:"), ecid, number)
	return objectivec.Object{ID: rv}
}

// MachineIdentifierForVirtualMachineCloneWithECIDSerialNumber is an exported wrapper for the private method _machineIdentifierForVirtualMachineCloneWithECIDSerialNumber.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) MachineIdentifierForVirtualMachineCloneWithECIDSerialNumber(ecid uint64, number objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineCloneWithECID:serialNumber:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_machineIdentifierForVirtualMachineCloneWithECID:serialNumber:"}
		return nil, err
	}
	return _VZMacMachineIdentifierClass._machineIdentifierForVirtualMachineCloneWithECIDSerialNumber(ecid, number), nil
}

// CanMachineIdentifierForVirtualMachineCloneWithECIDSerialNumber reports whether the receiver responds to the private selector _machineIdentifierForVirtualMachineCloneWithECID:serialNumber:.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) CanMachineIdentifierForVirtualMachineCloneWithECIDSerialNumber() bool {
	return objc.RespondsToSelector(objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineCloneWithECID:serialNumber:"))
}
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) _machineIdentifierForVirtualMachineCloneWithSerialNumber(number objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineCloneWithSerialNumber:"), number)
	return objectivec.Object{ID: rv}
}

// MachineIdentifierForVirtualMachineCloneWithSerialNumber is an exported wrapper for the private method _machineIdentifierForVirtualMachineCloneWithSerialNumber.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) MachineIdentifierForVirtualMachineCloneWithSerialNumber(number objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineCloneWithSerialNumber:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_machineIdentifierForVirtualMachineCloneWithSerialNumber:"}
		return nil, err
	}
	return _VZMacMachineIdentifierClass._machineIdentifierForVirtualMachineCloneWithSerialNumber(number), nil
}

// CanMachineIdentifierForVirtualMachineCloneWithSerialNumber reports whether the receiver responds to the private selector _machineIdentifierForVirtualMachineCloneWithSerialNumber:.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) CanMachineIdentifierForVirtualMachineCloneWithSerialNumber() bool {
	return objc.RespondsToSelector(objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierForVirtualMachineCloneWithSerialNumber:"))
}
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) _machineIdentifierWithECIDSerialNumber(ecid uint64, number objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierWithECID:serialNumber:"), ecid, number)
	return objectivec.Object{ID: rv}
}

// MachineIdentifierWithECIDSerialNumber is an exported wrapper for the private method _machineIdentifierWithECIDSerialNumber.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) MachineIdentifierWithECIDSerialNumber(ecid uint64, number objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierWithECID:serialNumber:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_machineIdentifierWithECID:serialNumber:"}
		return nil, err
	}
	return _VZMacMachineIdentifierClass._machineIdentifierWithECIDSerialNumber(ecid, number), nil
}

// CanMachineIdentifierWithECIDSerialNumber reports whether the receiver responds to the private selector _machineIdentifierWithECID:serialNumber:.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) CanMachineIdentifierWithECIDSerialNumber() bool {
	return objc.RespondsToSelector(objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierWithECID:serialNumber:"))
}
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) _machineIdentifierWithSerialNumber(number objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierWithSerialNumber:"), number)
	return objectivec.Object{ID: rv}
}

// MachineIdentifierWithSerialNumber is an exported wrapper for the private method _machineIdentifierWithSerialNumber.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) MachineIdentifierWithSerialNumber(number objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierWithSerialNumber:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_machineIdentifierWithSerialNumber:"}
		return nil, err
	}
	return _VZMacMachineIdentifierClass._machineIdentifierWithSerialNumber(number), nil
}

// CanMachineIdentifierWithSerialNumber reports whether the receiver responds to the private selector _machineIdentifierWithSerialNumber:.
func (_VZMacMachineIdentifierClass VZMacMachineIdentifierClass) CanMachineIdentifierWithSerialNumber() bool {
	return objc.RespondsToSelector(objc.ID(_VZMacMachineIdentifierClass.class), objc.Sel("_machineIdentifierWithSerialNumber:"))
}

func (v VZMacMachineIdentifier) _ECID() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("_ECID"))
	return rv
}

// CanECID reports whether the receiver responds to the private selector _ECID.
func (v VZMacMachineIdentifier) CanECID() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_ECID"))
}

// ECID is an exported wrapper for the private property _ECID.
func (v VZMacMachineIdentifier) ECID() (uint64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_ECID")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_ECID"}
	}
	return v._ECID(), nil
}
func (v VZMacMachineIdentifier) _ECIDChecksDisabled() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_ECIDChecksDisabled"))
	return rv
}

// CanECIDChecksDisabled reports whether the receiver responds to the private selector _ECIDChecksDisabled.
func (v VZMacMachineIdentifier) CanECIDChecksDisabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_ECIDChecksDisabled"))
}

// ECIDChecksDisabled is an exported wrapper for the private property _ECIDChecksDisabled.
func (v VZMacMachineIdentifier) ECIDChecksDisabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_ECIDChecksDisabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_ECIDChecksDisabled"}
	}
	return v._ECIDChecksDisabled(), nil
}
func (v VZMacMachineIdentifier) _serialNumber() IVZMacSerialNumber {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_serialNumber"))
	return VZMacSerialNumberFromID(objc.ID(rv))
}

// CanSerialNumber reports whether the receiver responds to the private selector _serialNumber.
func (v VZMacMachineIdentifier) CanSerialNumber() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_serialNumber"))
}

// SerialNumber is an exported wrapper for the private property _serialNumber.
func (v VZMacMachineIdentifier) SerialNumber() (IVZMacSerialNumber, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_serialNumber")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_serialNumber"}
	}
	return v._serialNumber(), nil
}
func (v VZMacMachineIdentifier) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMacMachineIdentifier) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMacMachineIdentifier) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZMacMachineIdentifier) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
