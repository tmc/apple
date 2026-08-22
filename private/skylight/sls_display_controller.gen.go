// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSDisplayController] class.
var (
	_SLSDisplayControllerClass     SLSDisplayControllerClass
	_SLSDisplayControllerClassOnce sync.Once
)

func getSLSDisplayControllerClass() SLSDisplayControllerClass {
	_SLSDisplayControllerClassOnce.Do(func() {
		_SLSDisplayControllerClass = SLSDisplayControllerClass{class: objc.GetClass("SLSDisplayController")}
	})
	return _SLSDisplayControllerClass
}

// GetSLSDisplayControllerClass returns the class object for SLSDisplayController.
func GetSLSDisplayControllerClass() SLSDisplayControllerClass {
	return getSLSDisplayControllerClass()
}

type SLSDisplayControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSDisplayControllerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSDisplayControllerClass) Alloc() SLSDisplayController {
	rv := objc.SendIfResponds[SLSDisplayController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSDisplayController.AbortContrastEnhancerRampError]
//   - [SLSDisplayController.AbortWhitePointRampError]
//   - [SLSDisplayController.BrightnessAvailable]
//   - [SLSDisplayController.SetBrightnessAvailable]
//   - [SLSDisplayController.BrightnessCapabilities]
//   - [SLSDisplayController.SetBrightnessCapabilities]
//   - [SLSDisplayController.CommitBrightness]
//   - [SLSDisplayController.CommitBrightnessTimeouts]
//   - [SLSDisplayController.ContainerId]
//   - [SLSDisplayController.SetContainerId]
//   - [SLSDisplayController.DisplayAttachmentSeed]
//   - [SLSDisplayController.SetDisplayAttachmentSeed]
//   - [SLSDisplayController.DisplayId]
//   - [SLSDisplayController.SetDisplayId]
//   - [SLSDisplayController.DisplayInfo]
//   - [SLSDisplayController.DisplayType]
//   - [SLSDisplayController.SetDisplayType]
//   - [SLSDisplayController.GetLinearBrightnessError]
//   - [SLSDisplayController.GetNitsError]
//   - [SLSDisplayController.IsOnline]
//   - [SLSDisplayController.MaximumLuminance]
//   - [SLSDisplayController.SetMaximumLuminance]
//   - [SLSDisplayController.NativeWhitePoint]
//   - [SLSDisplayController.SetNativeWhitePoint]
//   - [SLSDisplayController.NotificationQueue]
//   - [SLSDisplayController.PostNotificationPayload]
//   - [SLSDisplayController.ProductId]
//   - [SLSDisplayController.SetProductId]
//   - [SLSDisplayController.RegisterForNotificationsWithBlock]
//   - [SLSDisplayController.SerialNumber]
//   - [SLSDisplayController.SetSerialNumber]
//   - [SLSDisplayController.SetAmbient]
//   - [SLSDisplayController.SetApplyPolicy]
//   - [SLSDisplayController.SetBrightnessLimit]
//   - [SLSDisplayController.SetContrastEnhancer]
//   - [SLSDisplayController.SetContrastEnhancerRampDurationError]
//   - [SLSDisplayController.SetContrastPreservation]
//   - [SLSDisplayController.SetDimMessagingTimeout]
//   - [SLSDisplayController.SetFilteredAmbient]
//   - [SLSDisplayController.SetHeadroom]
//   - [SLSDisplayController.SetHighAmbientAdaptation]
//   - [SLSDisplayController.SetIndicatorBrightness]
//   - [SLSDisplayController.SetIndicatorBrightnessLimit]
//   - [SLSDisplayController.SetLinearBrightnessError]
//   - [SLSDisplayController.SetLowAmbientAdaptation]
//   - [SLSDisplayController.SetNotificationQueue]
//   - [SLSDisplayController.SetPotentialHeadroom]
//   - [SLSDisplayController.SetReferenceHeadroom]
//   - [SLSDisplayController.SetSDRBrightness]
//   - [SLSDisplayController.SetShieldingTimeout]
//   - [SLSDisplayController.SetSleepMessagingTimeout]
//   - [SLSDisplayController.SetWhitePointRampDurationError]
//   - [SLSDisplayController.UnregisterNotificationBlocks]
//   - [SLSDisplayController.Uuid]
//   - [SLSDisplayController.SetUuid]
//   - [SLSDisplayController.VendorId]
//   - [SLSDisplayController.SetVendorId]
//   - [SLSDisplayController.WhitePointAvailable]
//   - [SLSDisplayController.SetWhitePointAvailable]
//   - [SLSDisplayController.WhitePointD50XYZ]
//   - [SLSDisplayController.SetWhitePointD50XYZ]
//   - [SLSDisplayController.InitWithDisplayIdCapabilitiesClient]
//   - [SLSDisplayController.DebugDescription]
//   - [SLSDisplayController.Description]
//   - [SLSDisplayController.Hash]
//   - [SLSDisplayController.Online]
//   - [SLSDisplayController.SetOnline]
//   - [SLSDisplayController.Superclass]
type SLSDisplayController struct {
	objectivec.Object
}

// SLSDisplayControllerFromID constructs a [SLSDisplayController] from an objc.ID.
func SLSDisplayControllerFromID(id objc.ID) SLSDisplayController {
	return SLSDisplayController{objectivec.Object{ID: id}}
}

// Ensure SLSDisplayController implements ISLSDisplayController.
var _ ISLSDisplayController = SLSDisplayController{}

// An interface definition for the [SLSDisplayController] class.
//
// # Methods
//
//   - [ISLSDisplayController.AbortContrastEnhancerRampError]
//   - [ISLSDisplayController.AbortWhitePointRampError]
//   - [ISLSDisplayController.BrightnessAvailable]
//   - [ISLSDisplayController.SetBrightnessAvailable]
//   - [ISLSDisplayController.BrightnessCapabilities]
//   - [ISLSDisplayController.SetBrightnessCapabilities]
//   - [ISLSDisplayController.CommitBrightness]
//   - [ISLSDisplayController.CommitBrightnessTimeouts]
//   - [ISLSDisplayController.ContainerId]
//   - [ISLSDisplayController.SetContainerId]
//   - [ISLSDisplayController.DisplayAttachmentSeed]
//   - [ISLSDisplayController.SetDisplayAttachmentSeed]
//   - [ISLSDisplayController.DisplayId]
//   - [ISLSDisplayController.SetDisplayId]
//   - [ISLSDisplayController.DisplayInfo]
//   - [ISLSDisplayController.DisplayType]
//   - [ISLSDisplayController.SetDisplayType]
//   - [ISLSDisplayController.GetLinearBrightnessError]
//   - [ISLSDisplayController.GetNitsError]
//   - [ISLSDisplayController.IsOnline]
//   - [ISLSDisplayController.MaximumLuminance]
//   - [ISLSDisplayController.SetMaximumLuminance]
//   - [ISLSDisplayController.NativeWhitePoint]
//   - [ISLSDisplayController.SetNativeWhitePoint]
//   - [ISLSDisplayController.NotificationQueue]
//   - [ISLSDisplayController.PostNotificationPayload]
//   - [ISLSDisplayController.ProductId]
//   - [ISLSDisplayController.SetProductId]
//   - [ISLSDisplayController.RegisterForNotificationsWithBlock]
//   - [ISLSDisplayController.SerialNumber]
//   - [ISLSDisplayController.SetSerialNumber]
//   - [ISLSDisplayController.SetAmbient]
//   - [ISLSDisplayController.SetApplyPolicy]
//   - [ISLSDisplayController.SetBrightnessLimit]
//   - [ISLSDisplayController.SetContrastEnhancer]
//   - [ISLSDisplayController.SetContrastEnhancerRampDurationError]
//   - [ISLSDisplayController.SetContrastPreservation]
//   - [ISLSDisplayController.SetDimMessagingTimeout]
//   - [ISLSDisplayController.SetFilteredAmbient]
//   - [ISLSDisplayController.SetHeadroom]
//   - [ISLSDisplayController.SetHighAmbientAdaptation]
//   - [ISLSDisplayController.SetIndicatorBrightness]
//   - [ISLSDisplayController.SetIndicatorBrightnessLimit]
//   - [ISLSDisplayController.SetLinearBrightnessError]
//   - [ISLSDisplayController.SetLowAmbientAdaptation]
//   - [ISLSDisplayController.SetNotificationQueue]
//   - [ISLSDisplayController.SetPotentialHeadroom]
//   - [ISLSDisplayController.SetReferenceHeadroom]
//   - [ISLSDisplayController.SetSDRBrightness]
//   - [ISLSDisplayController.SetShieldingTimeout]
//   - [ISLSDisplayController.SetSleepMessagingTimeout]
//   - [ISLSDisplayController.SetWhitePointRampDurationError]
//   - [ISLSDisplayController.UnregisterNotificationBlocks]
//   - [ISLSDisplayController.Uuid]
//   - [ISLSDisplayController.SetUuid]
//   - [ISLSDisplayController.VendorId]
//   - [ISLSDisplayController.SetVendorId]
//   - [ISLSDisplayController.WhitePointAvailable]
//   - [ISLSDisplayController.SetWhitePointAvailable]
//   - [ISLSDisplayController.WhitePointD50XYZ]
//   - [ISLSDisplayController.SetWhitePointD50XYZ]
//   - [ISLSDisplayController.InitWithDisplayIdCapabilitiesClient]
//   - [ISLSDisplayController.DebugDescription]
//   - [ISLSDisplayController.Description]
//   - [ISLSDisplayController.Hash]
//   - [ISLSDisplayController.Online]
//   - [ISLSDisplayController.SetOnline]
//   - [ISLSDisplayController.Superclass]
type ISLSDisplayController interface {
	objectivec.IObject

	// Topic: Methods

	AbortContrastEnhancerRampError() (float32, error)
	AbortWhitePointRampError(ramp unsafe.Pointer) (bool, error)
	BrightnessAvailable() bool
	SetBrightnessAvailable(value bool)
	BrightnessCapabilities() foundation.INSDictionary
	SetBrightnessCapabilities(value foundation.INSDictionary)
	CommitBrightness(brightness []objectivec.IObject) bool
	CommitBrightnessTimeouts(timeouts []objectivec.IObject) bool
	ContainerId() foundation.NSUUID
	SetContainerId(value foundation.NSUUID)
	DisplayAttachmentSeed() uint64
	SetDisplayAttachmentSeed(value uint64)
	DisplayId() int
	SetDisplayId(value int)
	DisplayInfo() objectivec.IObject
	DisplayType() uint32
	SetDisplayType(value uint32)
	GetLinearBrightnessError() (float32, error)
	GetNitsError() (float32, error)
	IsOnline() bool
	MaximumLuminance() float32
	SetMaximumLuminance(value float32)
	NativeWhitePoint() unsafe.Pointer
	SetNativeWhitePoint(value unsafe.Pointer)
	NotificationQueue() objectivec.IObject
	PostNotificationPayload(notification objectivec.IObject, payload objectivec.IObject)
	ProductId() uint64
	SetProductId(value uint64)
	RegisterForNotificationsWithBlock(notifications objectivec.IObject, block VoidHandler)
	SerialNumber() uint64
	SetSerialNumber(value uint64)
	SetAmbient(ambient float32)
	SetApplyPolicy()
	SetBrightnessLimit(limit float32)
	SetContrastEnhancer(enhancer float32)
	SetContrastEnhancerRampDurationError(enhancer float32, duration float64) (bool, error)
	SetContrastPreservation(preservation float32)
	SetDimMessagingTimeout(timeout float64)
	SetFilteredAmbient(ambient float32)
	SetHeadroom(headroom float32)
	SetHighAmbientAdaptation(adaptation float32)
	SetIndicatorBrightness(brightness float32)
	SetIndicatorBrightnessLimit(limit float32)
	SetLinearBrightnessError(brightness float32) (bool, error)
	SetLowAmbientAdaptation(adaptation float32)
	SetNotificationQueue(queue objectivec.IObject)
	SetPotentialHeadroom(headroom float32)
	SetReferenceHeadroom(headroom float32)
	SetSDRBrightness(sDRBrightness float32)
	SetShieldingTimeout(timeout float64)
	SetSleepMessagingTimeout(timeout float64)
	SetWhitePointRampDurationError(point unsafe.Pointer, duration float64) (bool, error)
	UnregisterNotificationBlocks()
	Uuid() foundation.NSUUID
	SetUuid(value foundation.NSUUID)
	VendorId() uint64
	SetVendorId(value uint64)
	WhitePointAvailable() bool
	SetWhitePointAvailable(value bool)
	WhitePointD50XYZ() bool
	SetWhitePointD50XYZ(value bool)
	InitWithDisplayIdCapabilitiesClient(id int, capabilities objectivec.IObject, client objectivec.IObject) SLSDisplayController
	DebugDescription() string
	Description() string
	Hash() uint64
	Online() bool
	SetOnline(value bool)
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (s SLSDisplayController) Init() SLSDisplayController {
	rv := objc.SendIfResponds[SLSDisplayController](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSDisplayController) Autorelease() SLSDisplayController {
	rv := objc.SendIfResponds[SLSDisplayController](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSDisplayController creates a new SLSDisplayController instance.
func NewSLSDisplayController() SLSDisplayController {
	class := getSLSDisplayControllerClass()
	rv := objc.SendIfResponds[SLSDisplayController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSDisplayControllerWithDisplayIdCapabilitiesClient(id int, capabilities objectivec.IObject, client objectivec.IObject) SLSDisplayController {
	instance := getSLSDisplayControllerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDisplayId:capabilities:client:"), id, capabilities, client)
	return SLSDisplayControllerFromID(rv)
}

func (s SLSDisplayController) AbortContrastEnhancerRampError() (float32, error) {
	var ramp float32
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("abortContrastEnhancerRamp:error:"), unsafe.Pointer(&ramp), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("abortContrastEnhancerRamp:error: returned NO with nil NSError")
	}
	return ramp, nil
}
func (s SLSDisplayController) AbortWhitePointRampError(ramp unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("abortWhitePointRamp:error:"), ramp, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("abortWhitePointRamp:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (s SLSDisplayController) CommitBrightness(brightness []objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("commitBrightness:"), objectivec.IObjectSliceToNSArray(brightness))
	return rv
}
func (s SLSDisplayController) CommitBrightnessTimeouts(timeouts []objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("commitBrightnessTimeouts:"), objectivec.IObjectSliceToNSArray(timeouts))
	return rv
}
func (s SLSDisplayController) DisplayInfo() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("displayInfo"))
	return objectivec.Object{ID: rv}
}
func (s SLSDisplayController) GetLinearBrightnessError() (float32, error) {
	var brightness float32
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("getLinearBrightness:error:"), unsafe.Pointer(&brightness), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("getLinearBrightness:error: returned NO with nil NSError")
	}
	return brightness, nil
}
func (s SLSDisplayController) GetNitsError() (float32, error) {
	var nits float32
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("getNits:error:"), unsafe.Pointer(&nits), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("getNits:error: returned NO with nil NSError")
	}
	return nits, nil
}
func (s SLSDisplayController) IsOnline() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isOnline"))
	return rv
}
func (s SLSDisplayController) NotificationQueue() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("notificationQueue"))
	return objectivec.Object{ID: rv}
}
func (s SLSDisplayController) PostNotificationPayload(notification objectivec.IObject, payload objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("postNotification:payload:"), notification, payload)
}
func (s SLSDisplayController) RegisterForNotificationsWithBlock(notifications objectivec.IObject, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("registerForNotifications:withBlock:"), notifications, _block1)
}
func (s SLSDisplayController) SetAmbient(ambient float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setAmbient:"), ambient)
}
func (s SLSDisplayController) SetApplyPolicy() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setApplyPolicy"))
}
func (s SLSDisplayController) SetBrightnessLimit(limit float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setBrightnessLimit:"), limit)
}
func (s SLSDisplayController) SetContrastEnhancer(enhancer float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setContrastEnhancer:"), enhancer)
}
func (s SLSDisplayController) SetContrastEnhancerRampDurationError(enhancer float32, duration float64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("setContrastEnhancer:rampDuration:error:"), enhancer, duration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setContrastEnhancer:rampDuration:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (s SLSDisplayController) SetContrastPreservation(preservation float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setContrastPreservation:"), preservation)
}
func (s SLSDisplayController) SetDimMessagingTimeout(timeout float64) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setDimMessagingTimeout:"), timeout)
}
func (s SLSDisplayController) SetFilteredAmbient(ambient float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setFilteredAmbient:"), ambient)
}
func (s SLSDisplayController) SetHeadroom(headroom float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setHeadroom:"), headroom)
}
func (s SLSDisplayController) SetHighAmbientAdaptation(adaptation float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setHighAmbientAdaptation:"), adaptation)
}
func (s SLSDisplayController) SetIndicatorBrightness(brightness float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setIndicatorBrightness:"), brightness)
}
func (s SLSDisplayController) SetIndicatorBrightnessLimit(limit float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setIndicatorBrightnessLimit:"), limit)
}
func (s SLSDisplayController) SetLinearBrightnessError(brightness float32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("setLinearBrightness:error:"), brightness, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setLinearBrightness:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (s SLSDisplayController) SetLowAmbientAdaptation(adaptation float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setLowAmbientAdaptation:"), adaptation)
}
func (s SLSDisplayController) SetNotificationQueue(queue objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setNotificationQueue:"), queue)
}
func (s SLSDisplayController) SetPotentialHeadroom(headroom float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setPotentialHeadroom:"), headroom)
}
func (s SLSDisplayController) SetReferenceHeadroom(headroom float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setReferenceHeadroom:"), headroom)
}
func (s SLSDisplayController) SetSDRBrightness(sDRBrightness float32) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setSDRBrightness:"), sDRBrightness)
}
func (s SLSDisplayController) SetShieldingTimeout(timeout float64) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setShieldingTimeout:"), timeout)
}
func (s SLSDisplayController) SetSleepMessagingTimeout(timeout float64) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setSleepMessagingTimeout:"), timeout)
}
func (s SLSDisplayController) SetWhitePointRampDurationError(point unsafe.Pointer, duration float64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("setWhitePoint:rampDuration:error:"), point, duration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setWhitePoint:rampDuration:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (s SLSDisplayController) UnregisterNotificationBlocks() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("unregisterNotificationBlocks"))
}
func (s SLSDisplayController) InitWithDisplayIdCapabilitiesClient(id int, capabilities objectivec.IObject, client objectivec.IObject) SLSDisplayController {
	rv := objc.SendIfResponds[SLSDisplayController](s.ID, objc.Sel("initWithDisplayId:capabilities:client:"), id, capabilities, client)
	return rv
}

func (s SLSDisplayController) BrightnessAvailable() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("brightnessAvailable"))
	return rv
}
func (s SLSDisplayController) SetBrightnessAvailable(value bool) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setBrightnessAvailable:"), value)
}
func (s SLSDisplayController) BrightnessCapabilities() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("brightnessCapabilities"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s SLSDisplayController) SetBrightnessCapabilities(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setBrightnessCapabilities:"), value)
}
func (s SLSDisplayController) ContainerId() foundation.NSUUID {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("containerId"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (s SLSDisplayController) SetContainerId(value foundation.NSUUID) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setContainerId:"), value)
}
func (s SLSDisplayController) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSDisplayController) Description() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSDisplayController) DisplayAttachmentSeed() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("displayAttachmentSeed"))
	return rv
}
func (s SLSDisplayController) SetDisplayAttachmentSeed(value uint64) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setDisplayAttachmentSeed:"), value)
}
func (s SLSDisplayController) DisplayId() int {
	rv := objc.SendIfResponds[int](s.ID, objc.Sel("displayId"))
	return rv
}
func (s SLSDisplayController) SetDisplayId(value int) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setDisplayId:"), value)
}
func (s SLSDisplayController) DisplayType() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("displayType"))
	return rv
}
func (s SLSDisplayController) SetDisplayType(value uint32) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setDisplayType:"), value)
}
func (s SLSDisplayController) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("hash"))
	return rv
}
func (s SLSDisplayController) MaximumLuminance() float32 {
	rv := objc.SendIfResponds[float32](s.ID, objc.Sel("maximumLuminance"))
	return rv
}
func (s SLSDisplayController) SetMaximumLuminance(value float32) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setMaximumLuminance:"), value)
}
func (s SLSDisplayController) NativeWhitePoint() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("nativeWhitePoint"))
	return rv
}
func (s SLSDisplayController) SetNativeWhitePoint(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setNativeWhitePoint:"), value)
}
func (s SLSDisplayController) Online() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("online"))
	return rv
}
func (s SLSDisplayController) SetOnline(value bool) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setOnline:"), value)
}
func (s SLSDisplayController) ProductId() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("productId"))
	return rv
}
func (s SLSDisplayController) SetProductId(value uint64) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setProductId:"), value)
}
func (s SLSDisplayController) SerialNumber() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("serialNumber"))
	return rv
}
func (s SLSDisplayController) SetSerialNumber(value uint64) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setSerialNumber:"), value)
}
func (s SLSDisplayController) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](s.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (s SLSDisplayController) Uuid() foundation.NSUUID {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (s SLSDisplayController) SetUuid(value foundation.NSUUID) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setUuid:"), value)
}
func (s SLSDisplayController) VendorId() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("vendorId"))
	return rv
}
func (s SLSDisplayController) SetVendorId(value uint64) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setVendorId:"), value)
}
func (s SLSDisplayController) WhitePointAvailable() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("whitePointAvailable"))
	return rv
}
func (s SLSDisplayController) SetWhitePointAvailable(value bool) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setWhitePointAvailable:"), value)
}
func (s SLSDisplayController) WhitePointD50XYZ() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("whitePointD50XYZ"))
	return rv
}
func (s SLSDisplayController) SetWhitePointD50XYZ(value bool) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setWhitePointD50XYZ:"), value)
}

// RegisterForNotificationsWithBlockSync is a synchronous wrapper around [SLSDisplayController.RegisterForNotificationsWithBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSDisplayController) RegisterForNotificationsWithBlockSync(ctx context.Context, notifications objectivec.IObject) error {
	done := make(chan struct{}, 1)
	s.RegisterForNotificationsWithBlock(notifications, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
