// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacGraphicsDisplayConfiguration] class.
var (
	_VZMacGraphicsDisplayConfigurationClass     VZMacGraphicsDisplayConfigurationClass
	_VZMacGraphicsDisplayConfigurationClassOnce sync.Once
)

func getVZMacGraphicsDisplayConfigurationClass() VZMacGraphicsDisplayConfigurationClass {
	_VZMacGraphicsDisplayConfigurationClassOnce.Do(func() {
		_VZMacGraphicsDisplayConfigurationClass = VZMacGraphicsDisplayConfigurationClass{class: objc.GetClass("VZMacGraphicsDisplayConfiguration")}
	})
	return _VZMacGraphicsDisplayConfigurationClass
}

// GetVZMacGraphicsDisplayConfigurationClass returns the class object for VZMacGraphicsDisplayConfiguration.
func GetVZMacGraphicsDisplayConfigurationClass() VZMacGraphicsDisplayConfigurationClass {
	return getVZMacGraphicsDisplayConfigurationClass()
}

type VZMacGraphicsDisplayConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacGraphicsDisplayConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacGraphicsDisplayConfigurationClass) Alloc() VZMacGraphicsDisplayConfiguration {
	rv := objc.SendIfResponds[VZMacGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacGraphicsDisplayConfiguration._connectionType]
//   - [VZMacGraphicsDisplayConfiguration.Set_connectionType]
//   - [VZMacGraphicsDisplayConfiguration._displayIdentifier]
//   - [VZMacGraphicsDisplayConfiguration.Set_displayIdentifier]
//   - [VZMacGraphicsDisplayConfiguration._displayMode]
//   - [VZMacGraphicsDisplayConfiguration.Set_displayMode]
//   - [VZMacGraphicsDisplayConfiguration._setConnectionType]
//   - [VZMacGraphicsDisplayConfiguration._setDisplayIdentifier]
//   - [VZMacGraphicsDisplayConfiguration._setDisplayMode]
type VZMacGraphicsDisplayConfiguration struct {
	VZGraphicsDisplayConfiguration
}

// VZMacGraphicsDisplayConfigurationFromID constructs a [VZMacGraphicsDisplayConfiguration] from an objc.ID.
func VZMacGraphicsDisplayConfigurationFromID(id objc.ID) VZMacGraphicsDisplayConfiguration {
	return VZMacGraphicsDisplayConfiguration{VZGraphicsDisplayConfiguration: VZGraphicsDisplayConfigurationFromID(id)}
}

// Ensure VZMacGraphicsDisplayConfiguration implements IVZMacGraphicsDisplayConfiguration.
var _ IVZMacGraphicsDisplayConfiguration = VZMacGraphicsDisplayConfiguration{}

// An interface definition for the [VZMacGraphicsDisplayConfiguration] class.
//
// # Methods
//
//   - [IVZMacGraphicsDisplayConfiguration._connectionType]
//   - [IVZMacGraphicsDisplayConfiguration.Set_connectionType]
//   - [IVZMacGraphicsDisplayConfiguration._displayIdentifier]
//   - [IVZMacGraphicsDisplayConfiguration.Set_displayIdentifier]
//   - [IVZMacGraphicsDisplayConfiguration._displayMode]
//   - [IVZMacGraphicsDisplayConfiguration.Set_displayMode]
//   - [IVZMacGraphicsDisplayConfiguration._setConnectionType]
//   - [IVZMacGraphicsDisplayConfiguration._setDisplayIdentifier]
//   - [IVZMacGraphicsDisplayConfiguration._setDisplayMode]
type IVZMacGraphicsDisplayConfiguration interface {
	IVZGraphicsDisplayConfiguration

	// Topic: Methods

	_connectionType() int64
	Set_connectionType(value int64)
	_displayIdentifier() string
	Set_displayIdentifier(value string)
	_displayMode() int64
	Set_displayMode(value int64)
	_setConnectionType(type_ int64)
	_setDisplayIdentifier(identifier objectivec.IObject)
	_setDisplayMode(mode int64)
}

// Init initializes the instance.
func (v VZMacGraphicsDisplayConfiguration) Init() VZMacGraphicsDisplayConfiguration {
	rv := objc.SendIfResponds[VZMacGraphicsDisplayConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacGraphicsDisplayConfiguration) Autorelease() VZMacGraphicsDisplayConfiguration {
	rv := objc.SendIfResponds[VZMacGraphicsDisplayConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDisplayConfiguration creates a new VZMacGraphicsDisplayConfiguration instance.
func NewVZMacGraphicsDisplayConfiguration() VZMacGraphicsDisplayConfiguration {
	class := getVZMacGraphicsDisplayConfigurationClass()
	rv := objc.SendIfResponds[VZMacGraphicsDisplayConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZMacGraphicsDisplayConfiguration) _setConnectionType(type_ int64) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setConnectionType:"), type_)
}

// SetConnectionType is an exported wrapper for the private method _setConnectionType.
func (v VZMacGraphicsDisplayConfiguration) SetConnectionType(type_ int64) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setConnectionType:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setConnectionType:"}
		return err
	}
	v._setConnectionType(type_)
	return nil
}

// CanSetConnectionType reports whether the receiver responds to the private selector _setConnectionType:.
func (v VZMacGraphicsDisplayConfiguration) CanSetConnectionType() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setConnectionType:"))
}
func (v VZMacGraphicsDisplayConfiguration) _setDisplayIdentifier(identifier objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setDisplayIdentifier:"), identifier)
}

// SetDisplayIdentifier is an exported wrapper for the private method _setDisplayIdentifier.
func (v VZMacGraphicsDisplayConfiguration) SetDisplayIdentifier(identifier objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setDisplayIdentifier:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setDisplayIdentifier:"}
		return err
	}
	v._setDisplayIdentifier(identifier)
	return nil
}

// CanSetDisplayIdentifier reports whether the receiver responds to the private selector _setDisplayIdentifier:.
func (v VZMacGraphicsDisplayConfiguration) CanSetDisplayIdentifier() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setDisplayIdentifier:"))
}
func (v VZMacGraphicsDisplayConfiguration) _setDisplayMode(mode int64) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setDisplayMode:"), mode)
}

// SetDisplayMode is an exported wrapper for the private method _setDisplayMode.
func (v VZMacGraphicsDisplayConfiguration) SetDisplayMode(mode int64) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setDisplayMode:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setDisplayMode:"}
		return err
	}
	v._setDisplayMode(mode)
	return nil
}

// CanSetDisplayMode reports whether the receiver responds to the private selector _setDisplayMode:.
func (v VZMacGraphicsDisplayConfiguration) CanSetDisplayMode() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setDisplayMode:"))
}

func (v VZMacGraphicsDisplayConfiguration) _connectionType() int64 {
	rv := objc.SendIfResponds[int64](v.ID, objc.Sel("_connectionType"))
	return rv
}

// CanConnectionType reports whether the receiver responds to the private selector _connectionType.
func (v VZMacGraphicsDisplayConfiguration) CanConnectionType() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_connectionType"))
}

// ConnectionType is an exported wrapper for the private property _connectionType.
func (v VZMacGraphicsDisplayConfiguration) ConnectionType() (int64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_connectionType")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_connectionType"}
	}
	return v._connectionType(), nil
}
func (v VZMacGraphicsDisplayConfiguration) Set_connectionType(value int64) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_connectionType:"), value)
}
func (v VZMacGraphicsDisplayConfiguration) _displayIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_displayIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// CanDisplayIdentifier reports whether the receiver responds to the private selector _displayIdentifier.
func (v VZMacGraphicsDisplayConfiguration) CanDisplayIdentifier() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_displayIdentifier"))
}

// DisplayIdentifier is an exported wrapper for the private property _displayIdentifier.
func (v VZMacGraphicsDisplayConfiguration) DisplayIdentifier() (string, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_displayIdentifier")) {
		return "", &objc.UnrecognizedSelectorError{Selector: "_displayIdentifier"}
	}
	return v._displayIdentifier(), nil
}
func (v VZMacGraphicsDisplayConfiguration) Set_displayIdentifier(value string) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_displayIdentifier:"), objc.String(value))
}
func (v VZMacGraphicsDisplayConfiguration) _displayMode() int64 {
	rv := objc.SendIfResponds[int64](v.ID, objc.Sel("_displayMode"))
	return rv
}

// CanDisplayMode reports whether the receiver responds to the private selector _displayMode.
func (v VZMacGraphicsDisplayConfiguration) CanDisplayMode() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_displayMode"))
}

// DisplayMode is an exported wrapper for the private property _displayMode.
func (v VZMacGraphicsDisplayConfiguration) DisplayMode() (int64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_displayMode")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_displayMode"}
	}
	return v._displayMode(), nil
}
func (v VZMacGraphicsDisplayConfiguration) Set_displayMode(value int64) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_displayMode:"), value)
}
