// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETLossConfig] class.
var (
	_ETLossConfigClass     ETLossConfigClass
	_ETLossConfigClassOnce sync.Once
)

func getETLossConfigClass() ETLossConfigClass {
	_ETLossConfigClassOnce.Do(func() {
		_ETLossConfigClass = ETLossConfigClass{class: objc.GetClass("ETLossConfig")}
	})
	return _ETLossConfigClass
}

// GetETLossConfigClass returns the class object for ETLossConfig.
func GetETLossConfigClass() ETLossConfigClass {
	return getETLossConfigClass()
}

type ETLossConfigClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETLossConfigClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETLossConfigClass) Alloc() ETLossConfig {
	rv := objc.Send[ETLossConfig](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETLossConfig.Custom_network_path]
//   - [ETLossConfig.SetCustom_network_path]
//   - [ETLossConfig.Label_name]
//   - [ETLossConfig.SetLabel_name]
//   - [ETLossConfig.Label_shape]
//   - [ETLossConfig.SetLabel_shape]
//   - [ETLossConfig.Loss_name]
//   - [ETLossConfig.SetLoss_name]
//   - [ETLossConfig.Mode]
//   - [ETLossConfig.SetMode]
//   - [ETLossConfig.Output_name]
//   - [ETLossConfig.SetOutput_name]
//
// See: https://developer.apple.com/documentation/Espresso/ETLossConfig
type ETLossConfig struct {
	objectivec.Object
}

// ETLossConfigFromID constructs a [ETLossConfig] from an objc.ID.
func ETLossConfigFromID(id objc.ID) ETLossConfig {
	return ETLossConfig{objectivec.Object{ID: id}}
}

// Ensure ETLossConfig implements IETLossConfig.
var _ IETLossConfig = ETLossConfig{}

// An interface definition for the [ETLossConfig] class.
//
// # Methods
//
//   - [IETLossConfig.Custom_network_path]
//   - [IETLossConfig.SetCustom_network_path]
//   - [IETLossConfig.Label_name]
//   - [IETLossConfig.SetLabel_name]
//   - [IETLossConfig.Label_shape]
//   - [IETLossConfig.SetLabel_shape]
//   - [IETLossConfig.Loss_name]
//   - [IETLossConfig.SetLoss_name]
//   - [IETLossConfig.Mode]
//   - [IETLossConfig.SetMode]
//   - [IETLossConfig.Output_name]
//   - [IETLossConfig.SetOutput_name]
//
// See: https://developer.apple.com/documentation/Espresso/ETLossConfig
type IETLossConfig interface {
	objectivec.IObject

	// Topic: Methods

	Custom_network_path() string
	SetCustom_network_path(value string)
	Label_name() string
	SetLabel_name(value string)
	Label_shape() foundation.INSArray
	SetLabel_shape(value foundation.INSArray)
	Loss_name() string
	SetLoss_name(value string)
	Mode() uint64
	SetMode(value uint64)
	Output_name() string
	SetOutput_name(value string)
}

// Init initializes the instance.
func (e ETLossConfig) Init() ETLossConfig {
	rv := objc.Send[ETLossConfig](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETLossConfig) Autorelease() ETLossConfig {
	rv := objc.Send[ETLossConfig](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETLossConfig creates a new ETLossConfig instance.
func NewETLossConfig() ETLossConfig {
	class := getETLossConfigClass()
	rv := objc.Send[ETLossConfig](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETLossConfig/L2Loss
func (_ETLossConfigClass ETLossConfigClass) L2Loss() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ETLossConfigClass.class), objc.Sel("L2Loss"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETLossConfig/softmaxCrossEntropyLoss
func (_ETLossConfigClass ETLossConfigClass) SoftmaxCrossEntropyLoss() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ETLossConfigClass.class), objc.Sel("softmaxCrossEntropyLoss"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETLossConfig/custom_network_path
func (e ETLossConfig) Custom_network_path() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("custom_network_path"))
	return foundation.NSStringFromID(rv).String()
}
func (e ETLossConfig) SetCustom_network_path(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setCustom_network_path:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Espresso/ETLossConfig/label_name
func (e ETLossConfig) Label_name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("label_name"))
	return foundation.NSStringFromID(rv).String()
}
func (e ETLossConfig) SetLabel_name(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setLabel_name:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Espresso/ETLossConfig/label_shape
func (e ETLossConfig) Label_shape() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("label_shape"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETLossConfig) SetLabel_shape(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setLabel_shape:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETLossConfig/loss_name
func (e ETLossConfig) Loss_name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("loss_name"))
	return foundation.NSStringFromID(rv).String()
}
func (e ETLossConfig) SetLoss_name(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setLoss_name:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Espresso/ETLossConfig/mode
func (e ETLossConfig) Mode() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("mode"))
	return rv
}
func (e ETLossConfig) SetMode(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setMode:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETLossConfig/output_name
func (e ETLossConfig) Output_name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("output_name"))
	return foundation.NSStringFromID(rv).String()
}
func (e ETLossConfig) SetOutput_name(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setOutput_name:"), objc.String(value))
}
