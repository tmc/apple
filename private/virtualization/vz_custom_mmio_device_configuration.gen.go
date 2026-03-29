// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[VZCustomMMIODeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZCustomMMIODeviceConfiguration.MMIORegions]
//   - [VZCustomMMIODeviceConfiguration.SetMMIORegions]
//   - [VZCustomMMIODeviceConfiguration.AdditionalProperties]
//   - [VZCustomMMIODeviceConfiguration.SetAdditionalProperties]
//   - [VZCustomMMIODeviceConfiguration.AdditionalXPCProperties]
//   - [VZCustomMMIODeviceConfiguration.SetAdditionalXPCProperties]
//   - [VZCustomMMIODeviceConfiguration.EncodeWithEncoder]
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
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration
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
//   - [IVZCustomMMIODeviceConfiguration.EncodeWithEncoder]
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
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration
type IVZCustomMMIODeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	MMIORegions() foundation.INSArray
	SetMMIORegions(value foundation.INSArray)
	AdditionalProperties() foundation.INSDictionary
	SetAdditionalProperties(value foundation.INSDictionary)
	AdditionalXPCProperties() objectivec.Object
	SetAdditionalXPCProperties(value objectivec.Object)
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	Irqs() foundation.INSArray
	SetIrqs(value foundation.INSArray)
	Provider() *VZCustomMMIODeviceProvider
	SetProvider(value *VZCustomMMIODeviceProvider)
	SupportsSaveRestore() bool
	SetSupportsSaveRestore(value bool)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZCustomMMIODeviceConfiguration) Init() VZCustomMMIODeviceConfiguration {
	rv := objc.Send[VZCustomMMIODeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomMMIODeviceConfiguration) Autorelease() VZCustomMMIODeviceConfiguration {
	rv := objc.Send[VZCustomMMIODeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomMMIODeviceConfiguration creates a new VZCustomMMIODeviceConfiguration instance.
func NewVZCustomMMIODeviceConfiguration() VZCustomMMIODeviceConfiguration {
	class := getVZCustomMMIODeviceConfigurationClass()
	rv := objc.Send[VZCustomMMIODeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration/encodeWithEncoder:
func (v VZCustomMMIODeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration/MMIORegions
func (v VZCustomMMIODeviceConfiguration) MMIORegions() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("MMIORegions"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZCustomMMIODeviceConfiguration) SetMMIORegions(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("setMMIORegions:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration/additionalProperties
func (v VZCustomMMIODeviceConfiguration) AdditionalProperties() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("additionalProperties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (v VZCustomMMIODeviceConfiguration) SetAdditionalProperties(value foundation.INSDictionary) {
	objc.Send[struct{}](v.ID, objc.Sel("setAdditionalProperties:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration/additionalXPCProperties
func (v VZCustomMMIODeviceConfiguration) AdditionalXPCProperties() objectivec.Object {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("additionalXPCProperties"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (v VZCustomMMIODeviceConfiguration) SetAdditionalXPCProperties(value objectivec.Object) {
	objc.Send[struct{}](v.ID, objc.Sel("setAdditionalXPCProperties:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration/debugDescription
func (v VZCustomMMIODeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration/description
func (v VZCustomMMIODeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration/hash
func (v VZCustomMMIODeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration/irqs
func (v VZCustomMMIODeviceConfiguration) Irqs() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("irqs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZCustomMMIODeviceConfiguration) SetIrqs(value foundation.INSArray) {
	objc.Send[struct{}](v.ID, objc.Sel("setIrqs:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration/provider
func (v VZCustomMMIODeviceConfiguration) Provider() *VZCustomMMIODeviceProvider {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("provider"))
	if rv == 0 {
		return nil
	}
	val := VZCustomMMIODeviceProviderFromID(objc.ID(rv))
	return &val
}
func (v VZCustomMMIODeviceConfiguration) SetProvider(value *VZCustomMMIODeviceProvider) {
	if value == nil {
		objc.Send[struct{}](v.ID, objc.Sel("setProvider:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](v.ID, objc.Sel("setProvider:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration/superclass
func (v VZCustomMMIODeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomMMIODeviceConfiguration/supportsSaveRestore
func (v VZCustomMMIODeviceConfiguration) SupportsSaveRestore() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("supportsSaveRestore"))
	return rv
}
func (v VZCustomMMIODeviceConfiguration) SetSupportsSaveRestore(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSupportsSaveRestore:"), value)
}

