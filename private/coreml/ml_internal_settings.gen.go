// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLInternalSettings] class.
var (
	_MLInternalSettingsClass     MLInternalSettingsClass
	_MLInternalSettingsClassOnce sync.Once
)

func getMLInternalSettingsClass() MLInternalSettingsClass {
	_MLInternalSettingsClassOnce.Do(func() {
		_MLInternalSettingsClass = MLInternalSettingsClass{class: objc.GetClass("MLInternalSettings")}
	})
	return _MLInternalSettingsClass
}

// GetMLInternalSettingsClass returns the class object for MLInternalSettings.
func GetMLInternalSettingsClass() MLInternalSettingsClass {
	return getMLInternalSettingsClass()
}

type MLInternalSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLInternalSettingsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLInternalSettingsClass) Alloc() MLInternalSettings {
	rv := objc.Send[MLInternalSettings](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLInternalSettings.EncodeWithCoder]
//   - [MLInternalSettings.IsNeuralNetworkGPUPathForbidden]
//   - [MLInternalSettings.RestrictNeuralNetworksFromUsingANE]
//   - [MLInternalSettings.SetRestrictNeuralNetworksFromUsingANE]
//   - [MLInternalSettings.RestrictNeuralNetworksToUseCPUOnly]
//   - [MLInternalSettings.SetRestrictNeuralNetworksToUseCPUOnly]
//   - [MLInternalSettings.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings
type MLInternalSettings struct {
	objectivec.Object
}

// MLInternalSettingsFromID constructs a [MLInternalSettings] from an objc.ID.
func MLInternalSettingsFromID(id objc.ID) MLInternalSettings {
	return MLInternalSettings{objectivec.Object{ID: id}}
}

// Ensure MLInternalSettings implements IMLInternalSettings.
var _ IMLInternalSettings = MLInternalSettings{}

// An interface definition for the [MLInternalSettings] class.
//
// # Methods
//
//   - [IMLInternalSettings.EncodeWithCoder]
//   - [IMLInternalSettings.IsNeuralNetworkGPUPathForbidden]
//   - [IMLInternalSettings.RestrictNeuralNetworksFromUsingANE]
//   - [IMLInternalSettings.SetRestrictNeuralNetworksFromUsingANE]
//   - [IMLInternalSettings.RestrictNeuralNetworksToUseCPUOnly]
//   - [IMLInternalSettings.SetRestrictNeuralNetworksToUseCPUOnly]
//   - [IMLInternalSettings.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings
type IMLInternalSettings interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	IsNeuralNetworkGPUPathForbidden() bool
	RestrictNeuralNetworksFromUsingANE() bool
	SetRestrictNeuralNetworksFromUsingANE(value bool)
	RestrictNeuralNetworksToUseCPUOnly() bool
	SetRestrictNeuralNetworksToUseCPUOnly(value bool)
	InitWithCoder(coder foundation.INSCoder) MLInternalSettings
}

// Init initializes the instance.
func (i MLInternalSettings) Init() MLInternalSettings {
	rv := objc.Send[MLInternalSettings](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MLInternalSettings) Autorelease() MLInternalSettings {
	rv := objc.Send[MLInternalSettings](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLInternalSettings creates a new MLInternalSettings instance.
func NewMLInternalSettings() MLInternalSettings {
	class := getMLInternalSettingsClass()
	rv := objc.Send[MLInternalSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings/initWithCoder:
func NewInternalSettingsWithCoder(coder objectivec.IObject) MLInternalSettings {
	instance := getMLInternalSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLInternalSettingsFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings/encodeWithCoder:
func (i MLInternalSettings) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings/initWithCoder:
func (i MLInternalSettings) InitWithCoder(coder foundation.INSCoder) MLInternalSettings {
	rv := objc.Send[MLInternalSettings](i.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings/deviceHasANE
func (_MLInternalSettingsClass MLInternalSettingsClass) DeviceHasANE() bool {
	rv := objc.Send[bool](objc.ID(_MLInternalSettingsClass.class), objc.Sel("deviceHasANE"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings/globalSettings
func (_MLInternalSettingsClass MLInternalSettingsClass) GlobalSettings() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLInternalSettingsClass.class), objc.Sel("globalSettings"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings/globalSettingsFromSettings:
func (_MLInternalSettingsClass MLInternalSettingsClass) GlobalSettingsFromSettings(settings objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLInternalSettingsClass.class), objc.Sel("globalSettingsFromSettings:"), settings)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings/supportsSecureCoding
func (_MLInternalSettingsClass MLInternalSettingsClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLInternalSettingsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings/isNeuralNetworkGPUPathForbidden
func (i MLInternalSettings) IsNeuralNetworkGPUPathForbidden() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isNeuralNetworkGPUPathForbidden"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings/restrictNeuralNetworksFromUsingANE
func (i MLInternalSettings) RestrictNeuralNetworksFromUsingANE() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("restrictNeuralNetworksFromUsingANE"))
	return rv
}
func (i MLInternalSettings) SetRestrictNeuralNetworksFromUsingANE(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setRestrictNeuralNetworksFromUsingANE:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLInternalSettings/restrictNeuralNetworksToUseCPUOnly
func (i MLInternalSettings) RestrictNeuralNetworksToUseCPUOnly() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("restrictNeuralNetworksToUseCPUOnly"))
	return rv
}
func (i MLInternalSettings) SetRestrictNeuralNetworksToUseCPUOnly(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setRestrictNeuralNetworksToUseCPUOnly:"), value)
}
