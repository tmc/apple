// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomMMIODeviceConfiguration] class.
var (
	_VZCustomMMIODeviceConfigurationClass     VZCustomMMIODeviceConfigurationClass
	_VZCustomMMIODeviceConfigurationClassOnce sync.Once
)

func getVZCustomMMIODeviceConfigurationClass() VZCustomMMIODeviceConfigurationClass {
	_VZCustomMMIODeviceConfigurationClassOnce.Do(func() {
		_VZCustomMMIODeviceConfigurationClass = VZCustomMMIODeviceConfigurationClass{class: objc.GetClass("_VZCustomMMIODeviceConfiguration")}
	})
	return _VZCustomMMIODeviceConfigurationClass
}

// GetVZCustomMMIODeviceConfigurationClass returns the class object for _VZCustomMMIODeviceConfiguration.
func GetVZCustomMMIODeviceConfigurationClass() VZCustomMMIODeviceConfigurationClass {
	return getVZCustomMMIODeviceConfigurationClass()
}

type VZCustomMMIODeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomMMIODeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomMMIODeviceConfigurationClass) Alloc() VZCustomMMIODeviceConfiguration {
	rv := objc.SendIfResponds[VZCustomMMIODeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCustomMMIODeviceConfiguration.MMIORegions]
//   - [VZCustomMMIODeviceConfiguration.SetMMIORegions]
//   - [VZCustomMMIODeviceConfiguration.AdditionalProperties]
//   - [VZCustomMMIODeviceConfiguration.SetAdditionalProperties]
//   - [VZCustomMMIODeviceConfiguration.AdditionalXPCProperties]
//   - [VZCustomMMIODeviceConfiguration.SetAdditionalXPCProperties]
//   - [VZCustomMMIODeviceConfiguration.Irqs]
//   - [VZCustomMMIODeviceConfiguration.SetIrqs]
//   - [VZCustomMMIODeviceConfiguration.Provider]
//   - [VZCustomMMIODeviceConfiguration.SetProvider]
//   - [VZCustomMMIODeviceConfiguration.SupportsSaveRestore]
//   - [VZCustomMMIODeviceConfiguration.SetSupportsSaveRestore]
//   - [VZCustomMMIODeviceConfiguration.DebugDescription]
//   - [VZCustomMMIODeviceConfiguration.Description]
//   - [VZCustomMMIODeviceConfiguration.Hash]
//   - [VZCustomMMIODeviceConfiguration.Superclass]
type VZCustomMMIODeviceConfiguration struct {
	objectivec.Object
}

// VZCustomMMIODeviceConfigurationFromID constructs a [VZCustomMMIODeviceConfiguration] from an objc.ID.
func VZCustomMMIODeviceConfigurationFromID(id objc.ID) VZCustomMMIODeviceConfiguration {
	return VZCustomMMIODeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZCustomMMIODeviceConfiguration implements IVZCustomMMIODeviceConfiguration.
var _ IVZCustomMMIODeviceConfiguration = VZCustomMMIODeviceConfiguration{}

// An interface definition for the [VZCustomMMIODeviceConfiguration] class.
//
// # Methods
//
//   - [IVZCustomMMIODeviceConfiguration.MMIORegions]
//   - [IVZCustomMMIODeviceConfiguration.SetMMIORegions]
//   - [IVZCustomMMIODeviceConfiguration.AdditionalProperties]
//   - [IVZCustomMMIODeviceConfiguration.SetAdditionalProperties]
//   - [IVZCustomMMIODeviceConfiguration.AdditionalXPCProperties]
//   - [IVZCustomMMIODeviceConfiguration.SetAdditionalXPCProperties]
//   - [IVZCustomMMIODeviceConfiguration.Irqs]
//   - [IVZCustomMMIODeviceConfiguration.SetIrqs]
//   - [IVZCustomMMIODeviceConfiguration.Provider]
//   - [IVZCustomMMIODeviceConfiguration.SetProvider]
//   - [IVZCustomMMIODeviceConfiguration.SupportsSaveRestore]
//   - [IVZCustomMMIODeviceConfiguration.SetSupportsSaveRestore]
//   - [IVZCustomMMIODeviceConfiguration.DebugDescription]
//   - [IVZCustomMMIODeviceConfiguration.Description]
//   - [IVZCustomMMIODeviceConfiguration.Hash]
//   - [IVZCustomMMIODeviceConfiguration.Superclass]
type IVZCustomMMIODeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	MMIORegions() foundation.INSArray
	SetMMIORegions(value foundation.INSArray)
	AdditionalProperties() foundation.INSDictionary
	SetAdditionalProperties(value foundation.INSDictionary)
	AdditionalXPCProperties() objectivec.Object
	SetAdditionalXPCProperties(value objectivec.Object)
	Irqs() foundation.INSArray
	SetIrqs(value foundation.INSArray)
	Provider() IVZCustomMMIODeviceProvider
	SetProvider(value IVZCustomMMIODeviceProvider)
	SupportsSaveRestore() bool
	SetSupportsSaveRestore(value bool)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZCustomMMIODeviceConfiguration) Init() VZCustomMMIODeviceConfiguration {
	rv := objc.SendIfResponds[VZCustomMMIODeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomMMIODeviceConfiguration) Autorelease() VZCustomMMIODeviceConfiguration {
	rv := objc.SendIfResponds[VZCustomMMIODeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomMMIODeviceConfiguration creates a new VZCustomMMIODeviceConfiguration instance.
func NewVZCustomMMIODeviceConfiguration() VZCustomMMIODeviceConfiguration {
	class := getVZCustomMMIODeviceConfigurationClass()
	rv := objc.SendIfResponds[VZCustomMMIODeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZCustomMMIODeviceConfiguration) MMIORegions() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("MMIORegions"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZCustomMMIODeviceConfiguration) SetMMIORegions(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setMMIORegions:"), value)
}
func (v VZCustomMMIODeviceConfiguration) AdditionalProperties() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("additionalProperties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (v VZCustomMMIODeviceConfiguration) SetAdditionalProperties(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setAdditionalProperties:"), value)
}
func (v VZCustomMMIODeviceConfiguration) AdditionalXPCProperties() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("additionalXPCProperties"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (v VZCustomMMIODeviceConfiguration) SetAdditionalXPCProperties(value objectivec.Object) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setAdditionalXPCProperties:"), value)
}
func (v VZCustomMMIODeviceConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZCustomMMIODeviceConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZCustomMMIODeviceConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZCustomMMIODeviceConfiguration) Irqs() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("irqs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZCustomMMIODeviceConfiguration) SetIrqs(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setIrqs:"), value)
}
func (v VZCustomMMIODeviceConfiguration) Provider() IVZCustomMMIODeviceProvider {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("provider"))
	return VZCustomMMIODeviceProviderFromID(objc.ID(rv))
}
func (v VZCustomMMIODeviceConfiguration) SetProvider(value IVZCustomMMIODeviceProvider) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setProvider:"), value)
}
func (v VZCustomMMIODeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (v VZCustomMMIODeviceConfiguration) SupportsSaveRestore() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("supportsSaveRestore"))
	return rv
}
func (v VZCustomMMIODeviceConfiguration) SetSupportsSaveRestore(value bool) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setSupportsSaveRestore:"), value)
}
