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
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration
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
func (m VZMacGraphicsDisplayConfiguration) Init() VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacGraphicsDisplayConfiguration) Autorelease() VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDisplayConfiguration creates a new VZMacGraphicsDisplayConfiguration instance.
func NewVZMacGraphicsDisplayConfiguration() VZMacGraphicsDisplayConfiguration {
	class := getVZMacGraphicsDisplayConfigurationClass()
	rv := objc.Send[VZMacGraphicsDisplayConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/_setConnectionType:
func (m VZMacGraphicsDisplayConfiguration) _setConnectionType(type_ int64) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setConnectionType:"), type_)
}

// SetConnectionType is an exported wrapper for the private method _setConnectionType.
func (m VZMacGraphicsDisplayConfiguration) SetConnectionType(type_ int64) {
	m._setConnectionType(type_)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/_setDisplayIdentifier:
func (m VZMacGraphicsDisplayConfiguration) _setDisplayIdentifier(identifier objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setDisplayIdentifier:"), identifier)
}

// SetDisplayIdentifier is an exported wrapper for the private method _setDisplayIdentifier.
func (m VZMacGraphicsDisplayConfiguration) SetDisplayIdentifier(identifier objectivec.IObject) {
	m._setDisplayIdentifier(identifier)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/_setDisplayMode:
func (m VZMacGraphicsDisplayConfiguration) _setDisplayMode(mode int64) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setDisplayMode:"), mode)
}

// SetDisplayMode is an exported wrapper for the private method _setDisplayMode.
func (m VZMacGraphicsDisplayConfiguration) SetDisplayMode(mode int64) {
	m._setDisplayMode(mode)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/_connectionType
func (m VZMacGraphicsDisplayConfiguration) _connectionType() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("_connectionType"))
	return rv
}
func (m VZMacGraphicsDisplayConfiguration) Set_connectionType(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("set_connectionType:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/_displayIdentifier
func (m VZMacGraphicsDisplayConfiguration) _displayIdentifier() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_displayIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (m VZMacGraphicsDisplayConfiguration) Set_displayIdentifier(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("set_displayIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDisplayConfiguration/_displayMode
func (m VZMacGraphicsDisplayConfiguration) _displayMode() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("_displayMode"))
	return rv
}
func (m VZMacGraphicsDisplayConfiguration) Set_displayMode(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("set_displayMode:"), value)
}
