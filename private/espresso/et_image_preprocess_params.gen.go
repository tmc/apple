// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETImagePreprocessParams] class.
var (
	_ETImagePreprocessParamsClass     ETImagePreprocessParamsClass
	_ETImagePreprocessParamsClassOnce sync.Once
)

func getETImagePreprocessParamsClass() ETImagePreprocessParamsClass {
	_ETImagePreprocessParamsClassOnce.Do(func() {
		_ETImagePreprocessParamsClass = ETImagePreprocessParamsClass{class: objc.GetClass("ETImagePreprocessParams")}
	})
	return _ETImagePreprocessParamsClass
}

// GetETImagePreprocessParamsClass returns the class object for ETImagePreprocessParams.
func GetETImagePreprocessParamsClass() ETImagePreprocessParamsClass {
	return getETImagePreprocessParamsClass()
}

type ETImagePreprocessParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETImagePreprocessParamsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETImagePreprocessParamsClass) Alloc() ETImagePreprocessParams {
	rv := objc.Send[ETImagePreprocessParams](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETImagePreprocessParams.Bias_b]
//   - [ETImagePreprocessParams.Bias_g]
//   - [ETImagePreprocessParams.Bias_r]
//   - [ETImagePreprocessParams.Channels]
//   - [ETImagePreprocessParams.Height]
//   - [ETImagePreprocessParams.Network_wants_bgr]
//   - [ETImagePreprocessParams.Scale]
//   - [ETImagePreprocessParams.Width]
//   - [ETImagePreprocessParams.InitWithHeightWidthNumChannelsScaleBiasRBiasGBiasBNetworkWantBGR]
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams
type ETImagePreprocessParams struct {
	objectivec.Object
}

// ETImagePreprocessParamsFromID constructs a [ETImagePreprocessParams] from an objc.ID.
func ETImagePreprocessParamsFromID(id objc.ID) ETImagePreprocessParams {
	return ETImagePreprocessParams{objectivec.Object{ID: id}}
}
// Ensure ETImagePreprocessParams implements IETImagePreprocessParams.
var _ IETImagePreprocessParams = ETImagePreprocessParams{}

// An interface definition for the [ETImagePreprocessParams] class.
//
// # Methods
//
//   - [IETImagePreprocessParams.Bias_b]
//   - [IETImagePreprocessParams.Bias_g]
//   - [IETImagePreprocessParams.Bias_r]
//   - [IETImagePreprocessParams.Channels]
//   - [IETImagePreprocessParams.Height]
//   - [IETImagePreprocessParams.Network_wants_bgr]
//   - [IETImagePreprocessParams.Scale]
//   - [IETImagePreprocessParams.Width]
//   - [IETImagePreprocessParams.InitWithHeightWidthNumChannelsScaleBiasRBiasGBiasBNetworkWantBGR]
//
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams
type IETImagePreprocessParams interface {
	objectivec.IObject

	// Topic: Methods

	Bias_b() float32
	Bias_g() float32
	Bias_r() float32
	Channels() uint64
	Height() uint64
	Network_wants_bgr() bool
	Scale() float32
	Width() uint64
	InitWithHeightWidthNumChannelsScaleBiasRBiasGBiasBNetworkWantBGR(height uint64, width uint64, channels uint64, scale float32, r float32, g float32, b float32, bgr bool) ETImagePreprocessParams
}

// Init initializes the instance.
func (e ETImagePreprocessParams) Init() ETImagePreprocessParams {
	rv := objc.Send[ETImagePreprocessParams](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETImagePreprocessParams) Autorelease() ETImagePreprocessParams {
	rv := objc.Send[ETImagePreprocessParams](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETImagePreprocessParams creates a new ETImagePreprocessParams instance.
func NewETImagePreprocessParams() ETImagePreprocessParams {
	class := getETImagePreprocessParamsClass()
	rv := objc.Send[ETImagePreprocessParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams/initWithHeight:Width:NumChannels:Scale:BiasR:BiasG:BiasB:NetworkWantBGR:
func NewETImagePreprocessParamsWithHeightWidthNumChannelsScaleBiasRBiasGBiasBNetworkWantBGR(height uint64, width uint64, channels uint64, scale float32, r float32, g float32, b float32, bgr bool) ETImagePreprocessParams {
	instance := getETImagePreprocessParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHeight:Width:NumChannels:Scale:BiasR:BiasG:BiasB:NetworkWantBGR:"), height, width, channels, scale, r, g, b, bgr)
	return ETImagePreprocessParamsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams/initWithHeight:Width:NumChannels:Scale:BiasR:BiasG:BiasB:NetworkWantBGR:
func (e ETImagePreprocessParams) InitWithHeightWidthNumChannelsScaleBiasRBiasGBiasBNetworkWantBGR(height uint64, width uint64, channels uint64, scale float32, r float32, g float32, b float32, bgr bool) ETImagePreprocessParams {
	rv := objc.Send[ETImagePreprocessParams](e.ID, objc.Sel("initWithHeight:Width:NumChannels:Scale:BiasR:BiasG:BiasB:NetworkWantBGR:"), height, width, channels, scale, r, g, b, bgr)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams/bias_b
func (e ETImagePreprocessParams) Bias_b() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("bias_b"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams/bias_g
func (e ETImagePreprocessParams) Bias_g() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("bias_g"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams/bias_r
func (e ETImagePreprocessParams) Bias_r() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("bias_r"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams/channels
func (e ETImagePreprocessParams) Channels() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("channels"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams/height
func (e ETImagePreprocessParams) Height() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("height"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams/network_wants_bgr
func (e ETImagePreprocessParams) Network_wants_bgr() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("network_wants_bgr"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams/scale
func (e ETImagePreprocessParams) Scale() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("scale"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/ETImagePreprocessParams/width
func (e ETImagePreprocessParams) Width() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("width"))
	return rv
}

