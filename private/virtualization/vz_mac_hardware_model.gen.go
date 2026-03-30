// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacHardwareModel] class.
var (
	_VZMacHardwareModelClass     VZMacHardwareModelClass
	_VZMacHardwareModelClassOnce sync.Once
)

func getVZMacHardwareModelClass() VZMacHardwareModelClass {
	_VZMacHardwareModelClassOnce.Do(func() {
		_VZMacHardwareModelClass = VZMacHardwareModelClass{class: objc.GetClass("VZMacHardwareModel")}
	})
	return _VZMacHardwareModelClass
}

// GetVZMacHardwareModelClass returns the class object for VZMacHardwareModel.
func GetVZMacHardwareModelClass() VZMacHardwareModelClass {
	return getVZMacHardwareModelClass()
}

type VZMacHardwareModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacHardwareModelClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacHardwareModelClass) Alloc() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacHardwareModel._boardID]
//   - [VZMacHardwareModel._isa]
//   - [VZMacHardwareModel._variantID]
//   - [VZMacHardwareModel._variantName]
//   - [VZMacHardwareModel.Supported]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel
type VZMacHardwareModel struct {
	objectivec.Object
}

// VZMacHardwareModelFromID constructs a [VZMacHardwareModel] from an objc.ID.
func VZMacHardwareModelFromID(id objc.ID) VZMacHardwareModel {
	return VZMacHardwareModel{objectivec.Object{ID: id}}
}

// Ensure VZMacHardwareModel implements IVZMacHardwareModel.
var _ IVZMacHardwareModel = VZMacHardwareModel{}

// An interface definition for the [VZMacHardwareModel] class.
//
// # Methods
//
//   - [IVZMacHardwareModel._boardID]
//   - [IVZMacHardwareModel._isa]
//   - [IVZMacHardwareModel._variantID]
//   - [IVZMacHardwareModel._variantName]
//   - [IVZMacHardwareModel.Supported]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel
type IVZMacHardwareModel interface {
	objectivec.IObject

	// Topic: Methods

	_boardID() uint32
	_isa() int64
	_variantID() uint32
	_variantName() string
	Supported() bool
}

// Init initializes the instance.
func (m VZMacHardwareModel) Init() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacHardwareModel) Autorelease() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacHardwareModel creates a new VZMacHardwareModel instance.
func NewVZMacHardwareModel() VZMacHardwareModel {
	class := getVZMacHardwareModelClass()
	rv := objc.Send[VZMacHardwareModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/_defaultBoardIDForPlatformVersion:
func (_VZMacHardwareModelClass VZMacHardwareModelClass) _defaultBoardIDForPlatformVersion(version uint32) uint32 {
	rv := objc.Send[uint32](objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_defaultBoardIDForPlatformVersion:"), version)
	return rv
}

// DefaultBoardIDForPlatformVersion is an exported wrapper for the private method _defaultBoardIDForPlatformVersion.
func (_VZMacHardwareModelClass VZMacHardwareModelClass) DefaultBoardIDForPlatformVersion(version uint32) uint32 {
	return _VZMacHardwareModelClass._defaultBoardIDForPlatformVersion(version)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/_defaultHardwareModel
func (_VZMacHardwareModelClass VZMacHardwareModelClass) _defaultHardwareModel() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_defaultHardwareModel"))
	return objectivec.Object{ID: rv}
}

// DefaultHardwareModel is an exported wrapper for the private method _defaultHardwareModel.
func (_VZMacHardwareModelClass VZMacHardwareModelClass) DefaultHardwareModel() objectivec.IObject {
	return _VZMacHardwareModelClass._defaultHardwareModel()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/_hardwareModelWithDescriptor:
func (_VZMacHardwareModelClass VZMacHardwareModelClass) _hardwareModelWithDescriptor(descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_hardwareModelWithDescriptor:"), descriptor)
	return objectivec.Object{ID: rv}
}

// HardwareModelWithDescriptor is an exported wrapper for the private method _hardwareModelWithDescriptor.
func (_VZMacHardwareModelClass VZMacHardwareModelClass) HardwareModelWithDescriptor(descriptor objectivec.IObject) objectivec.IObject {
	return _VZMacHardwareModelClass._hardwareModelWithDescriptor(descriptor)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/_boardID
func (m VZMacHardwareModel) _boardID() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("_boardID"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/_isa
func (m VZMacHardwareModel) _isa() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("_isa"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/_variantID
func (m VZMacHardwareModel) _variantID() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("_variantID"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/_variantName
func (m VZMacHardwareModel) _variantName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_variantName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacHardwareModel/supported
func (m VZMacHardwareModel) Supported() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supported"))
	return rv
}
