// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSDisplayPowerControlClient] class.
var (
	_SLSDisplayPowerControlClientClass     SLSDisplayPowerControlClientClass
	_SLSDisplayPowerControlClientClassOnce sync.Once
)

func getSLSDisplayPowerControlClientClass() SLSDisplayPowerControlClientClass {
	_SLSDisplayPowerControlClientClassOnce.Do(func() {
		_SLSDisplayPowerControlClientClass = SLSDisplayPowerControlClientClass{class: objc.GetClass("SLSDisplayPowerControlClient")}
	})
	return _SLSDisplayPowerControlClientClass
}

// GetSLSDisplayPowerControlClientClass returns the class object for SLSDisplayPowerControlClient.
func GetSLSDisplayPowerControlClientClass() SLSDisplayPowerControlClientClass {
	return getSLSDisplayPowerControlClientClass()
}

type SLSDisplayPowerControlClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSDisplayPowerControlClientClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSDisplayPowerControlClientClass) Alloc() SLSDisplayPowerControlClient {
	rv := objc.Send[SLSDisplayPowerControlClient](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSDisplayPowerControlClient.RequestStateChangeError]
//   - [SLSDisplayPowerControlClient.SendStateChangeRequestUuid]
//   - [SLSDisplayPowerControlClient.Service]
//   - [SLSDisplayPowerControlClient.SetService]
//   - [SLSDisplayPowerControlClient.InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayPowerControlClient.InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayPowerControlClient
type SLSDisplayPowerControlClient struct {
	SLSDisplayControlClient
}

// SLSDisplayPowerControlClientFromID constructs a [SLSDisplayPowerControlClient] from an objc.ID.
func SLSDisplayPowerControlClientFromID(id objc.ID) SLSDisplayPowerControlClient {
	return SLSDisplayPowerControlClient{SLSDisplayControlClient: SLSDisplayControlClientFromID(id)}
}

// Ensure SLSDisplayPowerControlClient implements ISLSDisplayPowerControlClient.
var _ ISLSDisplayPowerControlClient = SLSDisplayPowerControlClient{}

// An interface definition for the [SLSDisplayPowerControlClient] class.
//
// # Methods
//
//   - [ISLSDisplayPowerControlClient.RequestStateChangeError]
//   - [ISLSDisplayPowerControlClient.SendStateChangeRequestUuid]
//   - [ISLSDisplayPowerControlClient.Service]
//   - [ISLSDisplayPowerControlClient.SetService]
//   - [ISLSDisplayPowerControlClient.InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//   - [ISLSDisplayPowerControlClient.InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayPowerControlClient
type ISLSDisplayPowerControlClient interface {
	ISLSDisplayControlClient

	// Topic: Methods

	RequestStateChangeError(change objectivec.IObject) (uint64, error)
	SendStateChangeRequestUuid(request objectivec.IObject, uuid unsafe.Pointer) int
	Service() ISLSXPCService
	SetService(value ISLSXPCService)
	InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock(client []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) SLSDisplayPowerControlClient
	InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock(client []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) SLSDisplayPowerControlClient
}

// Init initializes the instance.
func (s SLSDisplayPowerControlClient) Init() SLSDisplayPowerControlClient {
	rv := objc.Send[SLSDisplayPowerControlClient](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSDisplayPowerControlClient) Autorelease() SLSDisplayPowerControlClient {
	rv := objc.Send[SLSDisplayPowerControlClient](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSDisplayPowerControlClient creates a new SLSDisplayPowerControlClient instance.
func NewSLSDisplayPowerControlClient() SLSDisplayPowerControlClient {
	class := getSLSDisplayPowerControlClientClass()
	rv := objc.Send[SLSDisplayPowerControlClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayPowerControlClient/requestStateChange:error:
func (s SLSDisplayPowerControlClient) RequestStateChangeError(change objectivec.IObject) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](s.ID, objc.Sel("requestStateChange:error:"), change, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayPowerControlClient/sendStateChangeRequest:uuid:
func (s SLSDisplayPowerControlClient) SendStateChangeRequestUuid(request objectivec.IObject, uuid unsafe.Pointer) int {
	rv := objc.Send[int](s.ID, objc.Sel("sendStateChangeRequest:uuid:"), request, uuid)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayPowerControlClient/initAsyncPowerControlClient:notifyQueue:notificationType:notificationBlock:
func (s SLSDisplayPowerControlClient) InitAsyncPowerControlClientNotifyQueueNotificationTypeNotificationBlock(client []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) SLSDisplayPowerControlClient {
	_block3, _ := NewVoidBlock(block)
	rv := objc.Send[SLSDisplayPowerControlClient](s.ID, objc.Sel("initAsyncPowerControlClient:notifyQueue:notificationType:notificationBlock:"), client, queue, type_, _block3)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayPowerControlClient/initPowerControlClient:notifyQueue:notificationType:notificationBlock:
func (s SLSDisplayPowerControlClient) InitPowerControlClientNotifyQueueNotificationTypeNotificationBlock(client []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) SLSDisplayPowerControlClient {
	_block3, _ := NewVoidBlock(block)
	rv := objc.Send[SLSDisplayPowerControlClient](s.ID, objc.Sel("initPowerControlClient:notifyQueue:notificationType:notificationBlock:"), client, queue, type_, _block3)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayPowerControlClient/service
func (s SLSDisplayPowerControlClient) Service() ISLSXPCService {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("service"))
	return SLSXPCServiceFromID(objc.ID(rv))
}
func (s SLSDisplayPowerControlClient) SetService(value ISLSXPCService) {
	objc.Send[struct{}](s.ID, objc.Sel("setService:"), value)
}
