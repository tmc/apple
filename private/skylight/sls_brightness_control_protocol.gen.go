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
type SLSBrightnessControl interface {
	objectivec.IObject

	// AbortContrastEnhancerRampError protocol.
	AbortContrastEnhancerRampError() (float32, error)

	// AbortWhitePointRampError protocol.
	AbortWhitePointRampError(ramp unsafe.Pointer) (bool, error)

	// BrightnessAvailable protocol.
	BrightnessAvailable() bool

	// BrightnessCapabilities protocol.
	BrightnessCapabilities() objectivec.IObject

	// ContainerId protocol.
	ContainerId() objectivec.IObject

	// DisplayId protocol.
	DisplayId() int

	// DisplayType protocol.
	DisplayType() uint32

	// GetLinearBrightnessError protocol.
	GetLinearBrightnessError() (float32, error)

	// GetNitsError protocol.
	GetNitsError() (float32, error)

	// IsOnline protocol.
	IsOnline() bool

	// MaximumLuminance protocol.
	MaximumLuminance() float32

	// NativeWhitePoint protocol.
	NativeWhitePoint() unsafe.Pointer

	// ProductId protocol.
	ProductId() uint64

	// RegisterForNotificationsWithBlock protocol.
	RegisterForNotificationsWithBlock(notifications objectivec.IObject, block ObjectHandler)

	// SerialNumber protocol.
	SerialNumber() uint64

	// SetContrastEnhancerRampDurationError protocol.
	SetContrastEnhancerRampDurationError(enhancer float32, duration float64) (bool, error)

	// SetLinearBrightnessError protocol.
	SetLinearBrightnessError(brightness float32) (bool, error)

	// SetNotificationQueue protocol.
	SetNotificationQueue(queue objectivec.IObject)

	// SetWhitePointRampDurationError protocol.
	SetWhitePointRampDurationError(point unsafe.Pointer, duration float64) (bool, error)

	// UnregisterNotificationBlocks protocol.
	UnregisterNotificationBlocks()

	// Uuid protocol.
	Uuid() objectivec.IObject

	// VendorId protocol.
	VendorId() uint64

	// WhitePointAvailable protocol.
	WhitePointAvailable() bool

	// WhitePointD50XYZ protocol.
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
func (o SLSBrightnessControlObject) AbortWhitePointRampError(ramp unsafe.Pointer) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("abortWhitePointRamp:error:"), ramp)
	if err != nil {
		return false, err
	}
	return rv, nil
}
func (o SLSBrightnessControlObject) BrightnessAvailable() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("brightnessAvailable"))
	return rv
}
func (o SLSBrightnessControlObject) BrightnessCapabilities() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("brightnessCapabilities"))
	return objectivec.Object{ID: rv}
}
func (o SLSBrightnessControlObject) ContainerId() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("containerId"))
	return objectivec.Object{ID: rv}
}
func (o SLSBrightnessControlObject) DisplayId() int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("displayId"))
	return rv
}
func (o SLSBrightnessControlObject) DisplayType() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("displayType"))
	return rv
}
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
func (o SLSBrightnessControlObject) IsOnline() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isOnline"))
	return rv
}
func (o SLSBrightnessControlObject) MaximumLuminance() float32 {
	rv := objc.SendIfResponds[float32](o.ID, objc.Sel("maximumLuminance"))
	return rv
}
func (o SLSBrightnessControlObject) NativeWhitePoint() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("nativeWhitePoint"))
	return rv
}
func (o SLSBrightnessControlObject) ProductId() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("productId"))
	return rv
}
func (o SLSBrightnessControlObject) RegisterForNotificationsWithBlock(notifications objectivec.IObject, block ObjectHandler) {
	_block1, _cleanup1 := NewObjectBlock(block)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("registerForNotifications:withBlock:"), notifications, objc.ID(_block1))
}
func (o SLSBrightnessControlObject) SerialNumber() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("serialNumber"))
	return rv
}
func (o SLSBrightnessControlObject) SetContrastEnhancerRampDurationError(enhancer float32, duration float64) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("setContrastEnhancer:rampDuration:error:"), enhancer, duration)
	if err != nil {
		return false, err
	}
	return rv, nil
}
func (o SLSBrightnessControlObject) SetLinearBrightnessError(brightness float32) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("setLinearBrightness:error:"), brightness)
	if err != nil {
		return false, err
	}
	return rv, nil
}
func (o SLSBrightnessControlObject) SetNotificationQueue(queue objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setNotificationQueue:"), queue)
}
func (o SLSBrightnessControlObject) SetWhitePointRampDurationError(point unsafe.Pointer, duration float64) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("setWhitePoint:rampDuration:error:"), point, duration)
	if err != nil {
		return false, err
	}
	return rv, nil
}
func (o SLSBrightnessControlObject) UnregisterNotificationBlocks() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("unregisterNotificationBlocks"))
}
func (o SLSBrightnessControlObject) Uuid() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("uuid"))
	return objectivec.Object{ID: rv}
}
func (o SLSBrightnessControlObject) VendorId() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("vendorId"))
	return rv
}
func (o SLSBrightnessControlObject) WhitePointAvailable() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("whitePointAvailable"))
	return rv
}
func (o SLSBrightnessControlObject) WhitePointD50XYZ() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("whitePointD50XYZ"))
	return rv
}
