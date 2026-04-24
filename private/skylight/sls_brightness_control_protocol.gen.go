// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"errors"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSBrightnessControl protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl
type SLSBrightnessControl interface {
	objectivec.IObject

	// AbortContrastEnhancerRampError protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/abortContrastEnhancerRamp:error:
	AbortContrastEnhancerRampError() (float32, error)

	// AbortWhitePointRampError protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/abortWhitePointRamp:error:
	AbortWhitePointRampError(ramp objectivec.IObject) (bool, error)

	// BrightnessAvailable protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/brightnessAvailable
	BrightnessAvailable() bool

	// DisplayId protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/displayId
	DisplayId() int

	// DisplayType protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/displayType
	DisplayType() uint32

	// GetLinearBrightnessError protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/getLinearBrightness:error:
	GetLinearBrightnessError() (float32, error)

	// GetNitsError protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/getNits:error:
	GetNitsError() (float32, error)

	// IsOnline protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/isOnline
	IsOnline() bool

	// MaximumLuminance protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/maximumLuminance
	MaximumLuminance() float32

	// NativeWhitePoint protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/nativeWhitePoint
	NativeWhitePoint() objectivec.IObject

	// ProductId protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/productId
	ProductId() uint64

	// SerialNumber protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/serialNumber
	SerialNumber() uint64

	// SetContrastEnhancerRampDurationError protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/setContrastEnhancer:rampDuration:error:
	SetContrastEnhancerRampDurationError(enhancer float32, duration float64) (bool, error)

	// SetLinearBrightnessError protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/setLinearBrightness:error:
	SetLinearBrightnessError(brightness float32) (bool, error)

	// SetWhitePointRampDurationError protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/setWhitePoint:rampDuration:error:
	SetWhitePointRampDurationError(point objectivec.IObject, duration float64) (bool, error)

	// UnregisterNotificationBlocks protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/unregisterNotificationBlocks
	UnregisterNotificationBlocks()

	// VendorId protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/vendorId
	VendorId() uint64

	// WhitePointAvailable protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/whitePointAvailable
	WhitePointAvailable() bool

	// WhitePointD50XYZ protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/whitePointD50XYZ
	WhitePointD50XYZ() bool
}

// SLSBrightnessControlObject wraps an existing Objective-C object that conforms to the SLSBrightnessControl protocol.
type SLSBrightnessControlObject struct {
	objectivec.Object
}

func (o SLSBrightnessControlObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLSBrightnessControlObjectFromID constructs a [SLSBrightnessControlObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLSBrightnessControlObjectFromID(id objc.ID) SLSBrightnessControlObject {
	return SLSBrightnessControlObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/abortContrastEnhancerRamp:error:
func (o SLSBrightnessControlObject) AbortContrastEnhancerRampError() (float32, error) {
	var ramp float32
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("abortContrastEnhancerRamp:error:"), unsafe.Pointer(&ramp), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("abortContrastEnhancerRamp:error: returned NO with nil NSError")
	}
	return ramp, nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/abortWhitePointRamp:error:
func (o SLSBrightnessControlObject) AbortWhitePointRampError(ramp objectivec.IObject) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("abortWhitePointRamp:error:"), ramp)
	if err != nil {
		return false, err
	}
	return rv, nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/brightnessAvailable
func (o SLSBrightnessControlObject) BrightnessAvailable() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("brightnessAvailable"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/brightnessCapabilities
func (o SLSBrightnessControlObject) BrightnessCapabilities() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("brightnessCapabilities"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/containerId
func (o SLSBrightnessControlObject) ContainerId() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("containerId"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/displayId
func (o SLSBrightnessControlObject) DisplayId() int {
	rv := objc.Send[int](o.ID, objc.Sel("displayId"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/displayType
func (o SLSBrightnessControlObject) DisplayType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("displayType"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/getLinearBrightness:error:
func (o SLSBrightnessControlObject) GetLinearBrightnessError() (float32, error) {
	var brightness float32
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("getLinearBrightness:error:"), unsafe.Pointer(&brightness), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("getLinearBrightness:error: returned NO with nil NSError")
	}
	return brightness, nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/getNits:error:
func (o SLSBrightnessControlObject) GetNitsError() (float32, error) {
	var nits float32
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("getNits:error:"), unsafe.Pointer(&nits), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("getNits:error: returned NO with nil NSError")
	}
	return nits, nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/isOnline
func (o SLSBrightnessControlObject) IsOnline() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isOnline"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/maximumLuminance
func (o SLSBrightnessControlObject) MaximumLuminance() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("maximumLuminance"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/nativeWhitePoint
func (o SLSBrightnessControlObject) NativeWhitePoint() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("nativeWhitePoint"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/productId
func (o SLSBrightnessControlObject) ProductId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("productId"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/serialNumber
func (o SLSBrightnessControlObject) SerialNumber() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("serialNumber"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/setContrastEnhancer:rampDuration:error:
func (o SLSBrightnessControlObject) SetContrastEnhancerRampDurationError(enhancer float32, duration float64) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("setContrastEnhancer:rampDuration:error:"), enhancer, duration)
	if err != nil {
		return false, err
	}
	return rv, nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/setLinearBrightness:error:
func (o SLSBrightnessControlObject) SetLinearBrightnessError(brightness float32) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("setLinearBrightness:error:"), brightness)
	if err != nil {
		return false, err
	}
	return rv, nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/setNotificationQueue:
func (o SLSBrightnessControlObject) SetNotificationQueue(queue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setNotificationQueue:"), queue)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/setWhitePoint:rampDuration:error:
func (o SLSBrightnessControlObject) SetWhitePointRampDurationError(point objectivec.IObject, duration float64) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("setWhitePoint:rampDuration:error:"), point, duration)
	if err != nil {
		return false, err
	}
	return rv, nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/unregisterNotificationBlocks
func (o SLSBrightnessControlObject) UnregisterNotificationBlocks() {
	objc.Send[struct{}](o.ID, objc.Sel("unregisterNotificationBlocks"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/uuid
func (o SLSBrightnessControlObject) Uuid() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("uuid"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/vendorId
func (o SLSBrightnessControlObject) VendorId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("vendorId"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/whitePointAvailable
func (o SLSBrightnessControlObject) WhitePointAvailable() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("whitePointAvailable"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControl/whitePointD50XYZ
func (o SLSBrightnessControlObject) WhitePointD50XYZ() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("whitePointD50XYZ"))
	return rv
}
