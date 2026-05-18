// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGenericPlatformConfiguration] class.
var (
	_VZGenericPlatformConfigurationClass     VZGenericPlatformConfigurationClass
	_VZGenericPlatformConfigurationClassOnce sync.Once
)

func getVZGenericPlatformConfigurationClass() VZGenericPlatformConfigurationClass {
	_VZGenericPlatformConfigurationClassOnce.Do(func() {
		_VZGenericPlatformConfigurationClass = VZGenericPlatformConfigurationClass{class: objc.GetClass("VZGenericPlatformConfiguration")}
	})
	return _VZGenericPlatformConfigurationClass
}

// GetVZGenericPlatformConfigurationClass returns the class object for VZGenericPlatformConfiguration.
func GetVZGenericPlatformConfigurationClass() VZGenericPlatformConfigurationClass {
	return getVZGenericPlatformConfigurationClass()
}

type VZGenericPlatformConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGenericPlatformConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGenericPlatformConfigurationClass) Alloc() VZGenericPlatformConfiguration {
	rv := objc.Send[VZGenericPlatformConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZGenericPlatformConfiguration._fineGrainTrapsEmulationEnabled]
//   - [VZGenericPlatformConfiguration.Set_fineGrainTrapsEmulationEnabled]
//   - [VZGenericPlatformConfiguration._guestType]
//   - [VZGenericPlatformConfiguration.Set_guestType]
//   - [VZGenericPlatformConfiguration._performanceMonitoringUnitEmulationEnabled]
//   - [VZGenericPlatformConfiguration.Set_performanceMonitoringUnitEmulationEnabled]
//   - [VZGenericPlatformConfiguration._setFineGrainedTrapsEmulationEnabled]
//   - [VZGenericPlatformConfiguration._setGuestType]
//   - [VZGenericPlatformConfiguration._setPerformanceMonitoringUnitEmulationEnabled]
//   - [VZGenericPlatformConfiguration.NestedVirtualizationEnabled]
//   - [VZGenericPlatformConfiguration.SetNestedVirtualizationEnabled]
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration
type VZGenericPlatformConfiguration struct {
	VZPlatformConfiguration
}

// VZGenericPlatformConfigurationFromID constructs a [VZGenericPlatformConfiguration] from an objc.ID.
func VZGenericPlatformConfigurationFromID(id objc.ID) VZGenericPlatformConfiguration {
	return VZGenericPlatformConfiguration{VZPlatformConfiguration: VZPlatformConfigurationFromID(id)}
}

// Ensure VZGenericPlatformConfiguration implements IVZGenericPlatformConfiguration.
var _ IVZGenericPlatformConfiguration = VZGenericPlatformConfiguration{}

// An interface definition for the [VZGenericPlatformConfiguration] class.
//
// # Methods
//
//   - [IVZGenericPlatformConfiguration._fineGrainTrapsEmulationEnabled]
//   - [IVZGenericPlatformConfiguration.Set_fineGrainTrapsEmulationEnabled]
//   - [IVZGenericPlatformConfiguration._guestType]
//   - [IVZGenericPlatformConfiguration.Set_guestType]
//   - [IVZGenericPlatformConfiguration._performanceMonitoringUnitEmulationEnabled]
//   - [IVZGenericPlatformConfiguration.Set_performanceMonitoringUnitEmulationEnabled]
//   - [IVZGenericPlatformConfiguration._setFineGrainedTrapsEmulationEnabled]
//   - [IVZGenericPlatformConfiguration._setGuestType]
//   - [IVZGenericPlatformConfiguration._setPerformanceMonitoringUnitEmulationEnabled]
//   - [IVZGenericPlatformConfiguration.NestedVirtualizationEnabled]
//   - [IVZGenericPlatformConfiguration.SetNestedVirtualizationEnabled]
//
// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration
type IVZGenericPlatformConfiguration interface {
	IVZPlatformConfiguration

	// Topic: Methods

	_fineGrainTrapsEmulationEnabled() bool
	Set_fineGrainTrapsEmulationEnabled(value bool)
	_guestType() string
	Set_guestType(value string)
	_performanceMonitoringUnitEmulationEnabled() bool
	Set_performanceMonitoringUnitEmulationEnabled(value bool)
	_setFineGrainedTrapsEmulationEnabled(enabled bool)
	_setGuestType(type_ objectivec.IObject)
	_setPerformanceMonitoringUnitEmulationEnabled(enabled bool)
	NestedVirtualizationEnabled() bool
	SetNestedVirtualizationEnabled(value bool)
}

// Init initializes the instance.
func (v VZGenericPlatformConfiguration) Init() VZGenericPlatformConfiguration {
	rv := objc.Send[VZGenericPlatformConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZGenericPlatformConfiguration) Autorelease() VZGenericPlatformConfiguration {
	rv := objc.Send[VZGenericPlatformConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGenericPlatformConfiguration creates a new VZGenericPlatformConfiguration instance.
func NewVZGenericPlatformConfiguration() VZGenericPlatformConfiguration {
	class := getVZGenericPlatformConfigurationClass()
	rv := objc.Send[VZGenericPlatformConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/_setFineGrainedTrapsEmulationEnabled:
func (v VZGenericPlatformConfiguration) _setFineGrainedTrapsEmulationEnabled(enabled bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setFineGrainedTrapsEmulationEnabled:"), enabled)
}

// SetFineGrainedTrapsEmulationEnabled is an exported wrapper for the private method _setFineGrainedTrapsEmulationEnabled.
func (v VZGenericPlatformConfiguration) SetFineGrainedTrapsEmulationEnabled(enabled bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setFineGrainedTrapsEmulationEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setFineGrainedTrapsEmulationEnabled:"}
		return err
	}
	v._setFineGrainedTrapsEmulationEnabled(enabled)
	return nil
}

// CanSetFineGrainedTrapsEmulationEnabled reports whether the receiver responds to the private selector _setFineGrainedTrapsEmulationEnabled:.
func (v VZGenericPlatformConfiguration) CanSetFineGrainedTrapsEmulationEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setFineGrainedTrapsEmulationEnabled:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/_setGuestType:
func (v VZGenericPlatformConfiguration) _setGuestType(type_ objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setGuestType:"), type_)
}

// SetGuestType is an exported wrapper for the private method _setGuestType.
func (v VZGenericPlatformConfiguration) SetGuestType(type_ objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setGuestType:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setGuestType:"}
		return err
	}
	v._setGuestType(type_)
	return nil
}

// CanSetGuestType reports whether the receiver responds to the private selector _setGuestType:.
func (v VZGenericPlatformConfiguration) CanSetGuestType() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setGuestType:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/_setPerformanceMonitoringUnitEmulationEnabled:
func (v VZGenericPlatformConfiguration) _setPerformanceMonitoringUnitEmulationEnabled(enabled bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setPerformanceMonitoringUnitEmulationEnabled:"), enabled)
}

// SetPerformanceMonitoringUnitEmulationEnabled is an exported wrapper for the private method _setPerformanceMonitoringUnitEmulationEnabled.
func (v VZGenericPlatformConfiguration) SetPerformanceMonitoringUnitEmulationEnabled(enabled bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setPerformanceMonitoringUnitEmulationEnabled:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setPerformanceMonitoringUnitEmulationEnabled:"}
		return err
	}
	v._setPerformanceMonitoringUnitEmulationEnabled(enabled)
	return nil
}

// CanSetPerformanceMonitoringUnitEmulationEnabled reports whether the receiver responds to the private selector _setPerformanceMonitoringUnitEmulationEnabled:.
func (v VZGenericPlatformConfiguration) CanSetPerformanceMonitoringUnitEmulationEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setPerformanceMonitoringUnitEmulationEnabled:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/_fineGrainTrapsEmulationEnabled
func (v VZGenericPlatformConfiguration) _fineGrainTrapsEmulationEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_fineGrainTrapsEmulationEnabled"))
	return rv
}

// CanFineGrainTrapsEmulationEnabled reports whether the receiver responds to the private selector _fineGrainTrapsEmulationEnabled.
func (v VZGenericPlatformConfiguration) CanFineGrainTrapsEmulationEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_fineGrainTrapsEmulationEnabled"))
}

// FineGrainTrapsEmulationEnabled is an exported wrapper for the private property _fineGrainTrapsEmulationEnabled.
func (v VZGenericPlatformConfiguration) FineGrainTrapsEmulationEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_fineGrainTrapsEmulationEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_fineGrainTrapsEmulationEnabled"}
	}
	return v._fineGrainTrapsEmulationEnabled(), nil
}
func (v VZGenericPlatformConfiguration) Set_fineGrainTrapsEmulationEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_fineGrainTrapsEmulationEnabled:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/_guestType
func (v VZGenericPlatformConfiguration) _guestType() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_guestType"))
	return foundation.NSStringFromID(rv).String()
}

// CanGuestType reports whether the receiver responds to the private selector _guestType.
func (v VZGenericPlatformConfiguration) CanGuestType() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_guestType"))
}

// GuestType is an exported wrapper for the private property _guestType.
func (v VZGenericPlatformConfiguration) GuestType() (string, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_guestType")) {
		return "", &objc.UnrecognizedSelectorError{Selector: "_guestType"}
	}
	return v._guestType(), nil
}
func (v VZGenericPlatformConfiguration) Set_guestType(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("set_guestType:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/_performanceMonitoringUnitEmulationEnabled
func (v VZGenericPlatformConfiguration) _performanceMonitoringUnitEmulationEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_performanceMonitoringUnitEmulationEnabled"))
	return rv
}

// CanPerformanceMonitoringUnitEmulationEnabled reports whether the receiver responds to the private selector _performanceMonitoringUnitEmulationEnabled.
func (v VZGenericPlatformConfiguration) CanPerformanceMonitoringUnitEmulationEnabled() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_performanceMonitoringUnitEmulationEnabled"))
}

// PerformanceMonitoringUnitEmulationEnabled is an exported wrapper for the private property _performanceMonitoringUnitEmulationEnabled.
func (v VZGenericPlatformConfiguration) PerformanceMonitoringUnitEmulationEnabled() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_performanceMonitoringUnitEmulationEnabled")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_performanceMonitoringUnitEmulationEnabled"}
	}
	return v._performanceMonitoringUnitEmulationEnabled(), nil
}
func (v VZGenericPlatformConfiguration) Set_performanceMonitoringUnitEmulationEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("set_performanceMonitoringUnitEmulationEnabled:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZGenericPlatformConfiguration/nestedVirtualizationEnabled
func (v VZGenericPlatformConfiguration) NestedVirtualizationEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("nestedVirtualizationEnabled"))
	return rv
}
func (v VZGenericPlatformConfiguration) SetNestedVirtualizationEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setNestedVirtualizationEnabled:"), value)
}
