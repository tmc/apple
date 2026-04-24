// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBrightnessControlClient] class.
var (
	_SLSBrightnessControlClientClass     SLSBrightnessControlClientClass
	_SLSBrightnessControlClientClassOnce sync.Once
)

func getSLSBrightnessControlClientClass() SLSBrightnessControlClientClass {
	_SLSBrightnessControlClientClassOnce.Do(func() {
		_SLSBrightnessControlClientClass = SLSBrightnessControlClientClass{class: objc.GetClass("SLSBrightnessControlClient")}
	})
	return _SLSBrightnessControlClientClass
}

// GetSLSBrightnessControlClientClass returns the class object for SLSBrightnessControlClient.
func GetSLSBrightnessControlClientClass() SLSBrightnessControlClientClass {
	return getSLSBrightnessControlClientClass()
}

type SLSBrightnessControlClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBrightnessControlClientClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBrightnessControlClientClass) Alloc() SLSBrightnessControlClient {
	rv := objc.Send[SLSBrightnessControlClient](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBrightnessControlClient.BrightnessControls]
//   - [SLSBrightnessControlClient.CommitBrightnessPolicy]
//   - [SLSBrightnessControlClient.ControllerWithId]
//   - [SLSBrightnessControlClient.Displays]
//   - [SLSBrightnessControlClient.SetDisplays]
//   - [SLSBrightnessControlClient.GetFloatWithKeyFromReply]
//   - [SLSBrightnessControlClient.GetWhitePointMatrixWithKeyFromReply]
//   - [SLSBrightnessControlClient.HandleServerMessage]
//   - [SLSBrightnessControlClient.RequestAbortRampCommandDisplayError]
//   - [SLSBrightnessControlClient.RequestBrightnessPolicyError]
//   - [SLSBrightnessControlClient.RequestBrightnessTimeoutsError]
//   - [SLSBrightnessControlClient.RequestBulkBrightnessChangeError]
//   - [SLSBrightnessControlClient.RequestGetValueCommandDisplayError]
//   - [SLSBrightnessControlClient.RequestSetContrastEnhancerDurationDisplayError]
//   - [SLSBrightnessControlClient.RequestSetWhitePointDurationDisplayError]
//   - [SLSBrightnessControlClient.SendRequestCommandError]
//   - [SLSBrightnessControlClient.SendSynchronousRequestCommandError]
//   - [SLSBrightnessControlClient.Service]
//   - [SLSBrightnessControlClient.SetService]
//   - [SLSBrightnessControlClient.SetDimMessagingPolicy]
//   - [SLSBrightnessControlClient.SetNotifyBlock]
//   - [SLSBrightnessControlClient.SetShieldingPolicy]
//   - [SLSBrightnessControlClient.SetSleepMessagingPolicy]
//   - [SLSBrightnessControlClient.InitBrightnessControlClientNotifyQueueNotificationBlock]
//   - [SLSBrightnessControlClient.InitBrightnessControlClientNotifyQueueNotificationTypeNotificationBlock]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient
type SLSBrightnessControlClient struct {
	SLSDisplayControlClient
}

// SLSBrightnessControlClientFromID constructs a [SLSBrightnessControlClient] from an objc.ID.
func SLSBrightnessControlClientFromID(id objc.ID) SLSBrightnessControlClient {
	return SLSBrightnessControlClient{SLSDisplayControlClient: SLSDisplayControlClientFromID(id)}
}

// Ensure SLSBrightnessControlClient implements ISLSBrightnessControlClient.
var _ ISLSBrightnessControlClient = SLSBrightnessControlClient{}

// An interface definition for the [SLSBrightnessControlClient] class.
//
// # Methods
//
//   - [ISLSBrightnessControlClient.BrightnessControls]
//   - [ISLSBrightnessControlClient.CommitBrightnessPolicy]
//   - [ISLSBrightnessControlClient.ControllerWithId]
//   - [ISLSBrightnessControlClient.Displays]
//   - [ISLSBrightnessControlClient.SetDisplays]
//   - [ISLSBrightnessControlClient.GetFloatWithKeyFromReply]
//   - [ISLSBrightnessControlClient.GetWhitePointMatrixWithKeyFromReply]
//   - [ISLSBrightnessControlClient.HandleServerMessage]
//   - [ISLSBrightnessControlClient.RequestAbortRampCommandDisplayError]
//   - [ISLSBrightnessControlClient.RequestBrightnessPolicyError]
//   - [ISLSBrightnessControlClient.RequestBrightnessTimeoutsError]
//   - [ISLSBrightnessControlClient.RequestBulkBrightnessChangeError]
//   - [ISLSBrightnessControlClient.RequestGetValueCommandDisplayError]
//   - [ISLSBrightnessControlClient.RequestSetContrastEnhancerDurationDisplayError]
//   - [ISLSBrightnessControlClient.RequestSetWhitePointDurationDisplayError]
//   - [ISLSBrightnessControlClient.SendRequestCommandError]
//   - [ISLSBrightnessControlClient.SendSynchronousRequestCommandError]
//   - [ISLSBrightnessControlClient.Service]
//   - [ISLSBrightnessControlClient.SetService]
//   - [ISLSBrightnessControlClient.SetDimMessagingPolicy]
//   - [ISLSBrightnessControlClient.SetNotifyBlock]
//   - [ISLSBrightnessControlClient.SetShieldingPolicy]
//   - [ISLSBrightnessControlClient.SetSleepMessagingPolicy]
//   - [ISLSBrightnessControlClient.InitBrightnessControlClientNotifyQueueNotificationBlock]
//   - [ISLSBrightnessControlClient.InitBrightnessControlClientNotifyQueueNotificationTypeNotificationBlock]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient
type ISLSBrightnessControlClient interface {
	ISLSDisplayControlClient

	// Topic: Methods

	BrightnessControls(controls []objectivec.IObject) objectivec.IObject
	CommitBrightnessPolicy(policy []objectivec.IObject) bool
	ControllerWithId(id uint32) objectivec.IObject
	Displays() foundation.INSArray
	SetDisplays(value foundation.INSArray)
	GetFloatWithKeyFromReply(key string, reply objectivec.IObject) (float32, bool)
	GetWhitePointMatrixWithKeyFromReply(matrix objectivec.IObject, key string, reply objectivec.IObject) bool
	HandleServerMessage(message objectivec.IObject)
	RequestAbortRampCommandDisplayError(ramp unsafe.Pointer, command uint64, display int) (int, error)
	RequestBrightnessPolicyError(policy objectivec.IObject) (uint64, error)
	RequestBrightnessTimeoutsError(timeouts objectivec.IObject) (uint64, error)
	RequestBulkBrightnessChangeError(change objectivec.IObject) (uint64, error)
	RequestGetValueCommandDisplayError(value unsafe.Pointer, command uint64, display int) (int, error)
	RequestSetContrastEnhancerDurationDisplayError(enhancer float32, duration float64, display int) (uint64, error)
	RequestSetWhitePointDurationDisplayError(point objectivec.IObject, duration float64, display int) (uint64, error)
	SendRequestCommandError(request objectivec.IObject, command uint64) (uint64, error)
	SendSynchronousRequestCommandError(request objectivec.IObject, command uint64) (objectivec.IObject, error)
	Service() ISLSXPCService
	SetService(value ISLSXPCService)
	SetDimMessagingPolicy(policy byte)
	SetNotifyBlock(block VoidHandler)
	SetShieldingPolicy(policy byte)
	SetSleepMessagingPolicy(policy byte)
	InitBrightnessControlClientNotifyQueueNotificationBlock(client []objectivec.IObject, queue objectivec.IObject, block VoidHandler) SLSBrightnessControlClient
	InitBrightnessControlClientNotifyQueueNotificationTypeNotificationBlock(client []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) SLSBrightnessControlClient
}

// Init initializes the instance.
func (s SLSBrightnessControlClient) Init() SLSBrightnessControlClient {
	rv := objc.Send[SLSBrightnessControlClient](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBrightnessControlClient) Autorelease() SLSBrightnessControlClient {
	rv := objc.Send[SLSBrightnessControlClient](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBrightnessControlClient creates a new SLSBrightnessControlClient instance.
func NewSLSBrightnessControlClient() SLSBrightnessControlClient {
	class := getSLSBrightnessControlClientClass()
	rv := objc.Send[SLSBrightnessControlClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/brightnessControls:
func (s SLSBrightnessControlClient) BrightnessControls(controls []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("brightnessControls:"), objectivec.IObjectSliceToNSArray(controls))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/commitBrightnessPolicy:
func (s SLSBrightnessControlClient) CommitBrightnessPolicy(policy []objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("commitBrightnessPolicy:"), objectivec.IObjectSliceToNSArray(policy))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/controllerWithId:
func (s SLSBrightnessControlClient) ControllerWithId(id uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("controllerWithId:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/getFloat:withKey:fromReply:
func (s SLSBrightnessControlClient) GetFloatWithKeyFromReply(key string, reply objectivec.IObject) (float32, bool) {
	var float float32
	rv := objc.Send[bool](s.ID, objc.Sel("getFloat:withKey:fromReply:"), unsafe.Pointer(&float), key, reply)
	return float, rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/getWhitePointMatrix:withKey:fromReply:
func (s SLSBrightnessControlClient) GetWhitePointMatrixWithKeyFromReply(matrix objectivec.IObject, key string, reply objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("getWhitePointMatrix:withKey:fromReply:"), matrix, unsafe.Pointer(unsafe.StringData(key+"\x00")), reply)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/handleServerMessage:
func (s SLSBrightnessControlClient) HandleServerMessage(message objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("handleServerMessage:"), message)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/requestAbortRamp:command:display:error:
func (s SLSBrightnessControlClient) RequestAbortRampCommandDisplayError(ramp unsafe.Pointer, command uint64, display int) (int, error) {
	var errorPtr objc.ID
	rv := objc.Send[int](s.ID, objc.Sel("requestAbortRamp:command:display:error:"), ramp, command, display, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/requestBrightnessPolicy:error:
func (s SLSBrightnessControlClient) RequestBrightnessPolicyError(policy objectivec.IObject) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](s.ID, objc.Sel("requestBrightnessPolicy:error:"), policy, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/requestBrightnessTimeouts:error:
func (s SLSBrightnessControlClient) RequestBrightnessTimeoutsError(timeouts objectivec.IObject) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](s.ID, objc.Sel("requestBrightnessTimeouts:error:"), timeouts, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/requestBulkBrightnessChange:error:
func (s SLSBrightnessControlClient) RequestBulkBrightnessChangeError(change objectivec.IObject) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](s.ID, objc.Sel("requestBulkBrightnessChange:error:"), change, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/requestGetValue:command:display:error:
func (s SLSBrightnessControlClient) RequestGetValueCommandDisplayError(value unsafe.Pointer, command uint64, display int) (int, error) {
	var errorPtr objc.ID
	rv := objc.Send[int](s.ID, objc.Sel("requestGetValue:command:display:error:"), value, command, display, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/requestSetContrastEnhancer:duration:display:error:
func (s SLSBrightnessControlClient) RequestSetContrastEnhancerDurationDisplayError(enhancer float32, duration float64, display int) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](s.ID, objc.Sel("requestSetContrastEnhancer:duration:display:error:"), enhancer, duration, display, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/requestSetWhitePoint:duration:display:error:
func (s SLSBrightnessControlClient) RequestSetWhitePointDurationDisplayError(point objectivec.IObject, duration float64, display int) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](s.ID, objc.Sel("requestSetWhitePoint:duration:display:error:"), point, duration, display, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/sendRequest:command:error:
func (s SLSBrightnessControlClient) SendRequestCommandError(request objectivec.IObject, command uint64) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](s.ID, objc.Sel("sendRequest:command:error:"), request, command, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/sendSynchronousRequest:command:error:
func (s SLSBrightnessControlClient) SendSynchronousRequestCommandError(request objectivec.IObject, command uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("sendSynchronousRequest:command:error:"), request, command, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/setDimMessagingPolicy:
func (s SLSBrightnessControlClient) SetDimMessagingPolicy(policy byte) {
	objc.Send[objc.ID](s.ID, objc.Sel("setDimMessagingPolicy:"), policy)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/setNotifyBlock:
func (s SLSBrightnessControlClient) SetNotifyBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("setNotifyBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/setShieldingPolicy:
func (s SLSBrightnessControlClient) SetShieldingPolicy(policy byte) {
	objc.Send[objc.ID](s.ID, objc.Sel("setShieldingPolicy:"), policy)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/setSleepMessagingPolicy:
func (s SLSBrightnessControlClient) SetSleepMessagingPolicy(policy byte) {
	objc.Send[objc.ID](s.ID, objc.Sel("setSleepMessagingPolicy:"), policy)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/initBrightnessControlClient:notifyQueue:notificationBlock:
func (s SLSBrightnessControlClient) InitBrightnessControlClientNotifyQueueNotificationBlock(client []objectivec.IObject, queue objectivec.IObject, block VoidHandler) SLSBrightnessControlClient {
	_block2, _ := NewVoidBlock(block)
	rv := objc.Send[SLSBrightnessControlClient](s.ID, objc.Sel("initBrightnessControlClient:notifyQueue:notificationBlock:"), client, queue, _block2)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/initBrightnessControlClient:notifyQueue:notificationType:notificationBlock:
func (s SLSBrightnessControlClient) InitBrightnessControlClientNotifyQueueNotificationTypeNotificationBlock(client []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) SLSBrightnessControlClient {
	_block3, _ := NewVoidBlock(block)
	rv := objc.Send[SLSBrightnessControlClient](s.ID, objc.Sel("initBrightnessControlClient:notifyQueue:notificationType:notificationBlock:"), client, queue, type_, _block3)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/displays
func (s SLSBrightnessControlClient) Displays() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displays"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s SLSBrightnessControlClient) SetDisplays(value foundation.INSArray) {
	objc.Send[struct{}](s.ID, objc.Sel("setDisplays:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBrightnessControlClient/service
func (s SLSBrightnessControlClient) Service() ISLSXPCService {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("service"))
	return SLSXPCServiceFromID(objc.ID(rv))
}
func (s SLSBrightnessControlClient) SetService(value ISLSXPCService) {
	objc.Send[struct{}](s.ID, objc.Sel("setService:"), value)
}

// SetNotifyBlockSync is a synchronous wrapper around [SLSBrightnessControlClient.SetNotifyBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSBrightnessControlClient) SetNotifyBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SetNotifyBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
