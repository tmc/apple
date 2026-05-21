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
func (v VZMacHardwareModel) Init() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacHardwareModel) Autorelease() VZMacHardwareModel {
	rv := objc.Send[VZMacHardwareModel](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacHardwareModel creates a new VZMacHardwareModel instance.
func NewVZMacHardwareModel() VZMacHardwareModel {
	class := getVZMacHardwareModelClass()
	rv := objc.Send[VZMacHardwareModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_VZMacHardwareModelClass VZMacHardwareModelClass) _defaultBoardIDForPlatformVersion(version uint32) uint32 {
	rv := objc.Send[uint32](objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_defaultBoardIDForPlatformVersion:"), version)
	return rv
}

// DefaultBoardIDForPlatformVersion is an exported wrapper for the private method _defaultBoardIDForPlatformVersion.
func (_VZMacHardwareModelClass VZMacHardwareModelClass) DefaultBoardIDForPlatformVersion(version uint32) (uint32, error) {
	if !objc.RespondsToSelector(objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_defaultBoardIDForPlatformVersion:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_defaultBoardIDForPlatformVersion:"}
		return 0, err
	}
	return _VZMacHardwareModelClass._defaultBoardIDForPlatformVersion(version), nil
}

// CanDefaultBoardIDForPlatformVersion reports whether the receiver responds to the private selector _defaultBoardIDForPlatformVersion:.
func (_VZMacHardwareModelClass VZMacHardwareModelClass) CanDefaultBoardIDForPlatformVersion() bool {
	return objc.RespondsToSelector(objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_defaultBoardIDForPlatformVersion:"))
}
func (_VZMacHardwareModelClass VZMacHardwareModelClass) _defaultHardwareModel() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_defaultHardwareModel"))
	return objectivec.Object{ID: rv}
}

// DefaultHardwareModel is an exported wrapper for the private method _defaultHardwareModel.
func (_VZMacHardwareModelClass VZMacHardwareModelClass) DefaultHardwareModel() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_defaultHardwareModel")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_defaultHardwareModel"}
		return nil, err
	}
	return _VZMacHardwareModelClass._defaultHardwareModel(), nil
}

// CanDefaultHardwareModel reports whether the receiver responds to the private selector _defaultHardwareModel.
func (_VZMacHardwareModelClass VZMacHardwareModelClass) CanDefaultHardwareModel() bool {
	return objc.RespondsToSelector(objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_defaultHardwareModel"))
}
func (_VZMacHardwareModelClass VZMacHardwareModelClass) _hardwareModelWithDescriptor(descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_hardwareModelWithDescriptor:"), descriptor)
	return objectivec.Object{ID: rv}
}

// HardwareModelWithDescriptor is an exported wrapper for the private method _hardwareModelWithDescriptor.
func (_VZMacHardwareModelClass VZMacHardwareModelClass) HardwareModelWithDescriptor(descriptor objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_hardwareModelWithDescriptor:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_hardwareModelWithDescriptor:"}
		return nil, err
	}
	return _VZMacHardwareModelClass._hardwareModelWithDescriptor(descriptor), nil
}

// CanHardwareModelWithDescriptor reports whether the receiver responds to the private selector _hardwareModelWithDescriptor:.
func (_VZMacHardwareModelClass VZMacHardwareModelClass) CanHardwareModelWithDescriptor() bool {
	return objc.RespondsToSelector(objc.ID(_VZMacHardwareModelClass.class), objc.Sel("_hardwareModelWithDescriptor:"))
}

func (v VZMacHardwareModel) _boardID() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("_boardID"))
	return rv
}

// CanBoardID reports whether the receiver responds to the private selector _boardID.
func (v VZMacHardwareModel) CanBoardID() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_boardID"))
}

// BoardID is an exported wrapper for the private property _boardID.
func (v VZMacHardwareModel) BoardID() (uint32, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_boardID")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_boardID"}
	}
	return v._boardID(), nil
}
func (v VZMacHardwareModel) _isa() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("_isa"))
	return rv
}

// CanIsa reports whether the receiver responds to the private selector _isa.
func (v VZMacHardwareModel) CanIsa() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_isa"))
}

// Isa is an exported wrapper for the private property _isa.
func (v VZMacHardwareModel) Isa() (int64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_isa")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_isa"}
	}
	return v._isa(), nil
}
func (v VZMacHardwareModel) _variantID() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("_variantID"))
	return rv
}

// CanVariantID reports whether the receiver responds to the private selector _variantID.
func (v VZMacHardwareModel) CanVariantID() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_variantID"))
}

// VariantID is an exported wrapper for the private property _variantID.
func (v VZMacHardwareModel) VariantID() (uint32, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_variantID")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_variantID"}
	}
	return v._variantID(), nil
}
func (v VZMacHardwareModel) _variantName() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_variantName"))
	return foundation.NSStringFromID(rv).String()
}

// CanVariantName reports whether the receiver responds to the private selector _variantName.
func (v VZMacHardwareModel) CanVariantName() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_variantName"))
}

// VariantName is an exported wrapper for the private property _variantName.
func (v VZMacHardwareModel) VariantName() (string, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_variantName")) {
		return "", &objc.UnrecognizedSelectorError{Selector: "_variantName"}
	}
	return v._variantName(), nil
}
func (v VZMacHardwareModel) Supported() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supported"))
	return rv
}
