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
func (m MLInternalSettings) Init() MLInternalSettings {
	rv := objc.Send[MLInternalSettings](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLInternalSettings) Autorelease() MLInternalSettings {
	rv := objc.Send[MLInternalSettings](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLInternalSettings creates a new MLInternalSettings instance.
func NewMLInternalSettings() MLInternalSettings {
	class := getMLInternalSettingsClass()
	rv := objc.Send[MLInternalSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewInternalSettingsWithCoder(coder objectivec.IObject) MLInternalSettings {
	instance := getMLInternalSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLInternalSettingsFromID(rv)
}

func (m MLInternalSettings) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (m MLInternalSettings) InitWithCoder(coder foundation.INSCoder) MLInternalSettings {
	rv := objc.Send[MLInternalSettings](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_MLInternalSettingsClass MLInternalSettingsClass) DeviceHasANE() bool {
	rv := objc.Send[bool](objc.ID(_MLInternalSettingsClass.class), objc.Sel("deviceHasANE"))
	return rv
}
func (_MLInternalSettingsClass MLInternalSettingsClass) GlobalSettings() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLInternalSettingsClass.class), objc.Sel("globalSettings"))
	return objectivec.Object{ID: rv}
}
func (_MLInternalSettingsClass MLInternalSettingsClass) GlobalSettingsFromSettings(settings objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLInternalSettingsClass.class), objc.Sel("globalSettingsFromSettings:"), settings)
	return objectivec.Object{ID: rv}
}
func (_MLInternalSettingsClass MLInternalSettingsClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLInternalSettingsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLInternalSettings) IsNeuralNetworkGPUPathForbidden() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isNeuralNetworkGPUPathForbidden"))
	return rv
}
func (m MLInternalSettings) RestrictNeuralNetworksFromUsingANE() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("restrictNeuralNetworksFromUsingANE"))
	return rv
}
func (m MLInternalSettings) SetRestrictNeuralNetworksFromUsingANE(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setRestrictNeuralNetworksFromUsingANE:"), value)
}
func (m MLInternalSettings) RestrictNeuralNetworksToUseCPUOnly() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("restrictNeuralNetworksToUseCPUOnly"))
	return rv
}
func (m MLInternalSettings) SetRestrictNeuralNetworksToUseCPUOnly(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setRestrictNeuralNetworksToUseCPUOnly:"), value)
}
