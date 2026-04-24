// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSGUIClient] class.
var (
	_SLSGUIClientClass     SLSGUIClientClass
	_SLSGUIClientClassOnce sync.Once
)

func getSLSGUIClientClass() SLSGUIClientClass {
	_SLSGUIClientClassOnce.Do(func() {
		_SLSGUIClientClass = SLSGUIClientClass{class: objc.GetClass("SLSGUIClient")}
	})
	return _SLSGUIClientClass
}

// GetSLSGUIClientClass returns the class object for SLSGUIClient.
func GetSLSGUIClientClass() SLSGUIClientClass {
	return getSLSGUIClientClass()
}

type SLSGUIClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSGUIClientClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSGUIClientClass) Alloc() SLSGUIClient {
	rv := objc.Send[SLSGUIClient](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSGUIClient.RequestDisplaysIdleError]
//   - [SLSGUIClient.Service]
//   - [SLSGUIClient.SetService]
//   - [SLSGUIClient.ValidateIdleRequest]
//   - [SLSGUIClient.InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSGUIClient
type SLSGUIClient struct {
	SLSDisplayControlClient
}

// SLSGUIClientFromID constructs a [SLSGUIClient] from an objc.ID.
func SLSGUIClientFromID(id objc.ID) SLSGUIClient {
	return SLSGUIClient{SLSDisplayControlClient: SLSDisplayControlClientFromID(id)}
}

// Ensure SLSGUIClient implements ISLSGUIClient.
var _ ISLSGUIClient = SLSGUIClient{}

// An interface definition for the [SLSGUIClient] class.
//
// # Methods
//
//   - [ISLSGUIClient.RequestDisplaysIdleError]
//   - [ISLSGUIClient.Service]
//   - [ISLSGUIClient.SetService]
//   - [ISLSGUIClient.ValidateIdleRequest]
//   - [ISLSGUIClient.InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSGUIClient
type ISLSGUIClient interface {
	ISLSDisplayControlClient

	// Topic: Methods

	RequestDisplaysIdleError(idle objectivec.IObject) (uint64, error)
	Service() ISLSXPCService
	SetService(value ISLSXPCService)
	ValidateIdleRequest(request objectivec.IObject) int
	InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock(gUIClient objectivec.IObject, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) SLSGUIClient
}

// Init initializes the instance.
func (s SLSGUIClient) Init() SLSGUIClient {
	rv := objc.Send[SLSGUIClient](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSGUIClient) Autorelease() SLSGUIClient {
	rv := objc.Send[SLSGUIClient](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSGUIClient creates a new SLSGUIClient instance.
func NewSLSGUIClient() SLSGUIClient {
	class := getSLSGUIClientClass()
	rv := objc.Send[SLSGUIClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSGUIClient/requestDisplaysIdle:error:
func (s SLSGUIClient) RequestDisplaysIdleError(idle objectivec.IObject) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](s.ID, objc.Sel("requestDisplaysIdle:error:"), idle, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLSGUIClient/validateIdleRequest:
func (s SLSGUIClient) ValidateIdleRequest(request objectivec.IObject) int {
	rv := objc.Send[int](s.ID, objc.Sel("validateIdleRequest:"), request)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSGUIClient/initGUIClient:error:notifyQueue:notificationType:notificationBlock:
func (s SLSGUIClient) InitGUIClientErrorNotifyQueueNotificationTypeNotificationBlock(gUIClient objectivec.IObject, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) SLSGUIClient {
	_block4, _ := NewVoidBlock(block)
	rv := objc.Send[SLSGUIClient](s.ID, objc.Sel("initGUIClient:error:notifyQueue:notificationType:notificationBlock:"), gUIClient, error_, queue, type_, _block4)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSGUIClient/configGUIClient:error:notifyQueue:notificationType:notificationBlock:
func (_SLSGUIClientClass SLSGUIClientClass) ConfigGUIClientErrorNotifyQueueNotificationTypeNotificationBlock(gUIClient objectivec.IObject, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) objectivec.IObject {
	_block4, _ := NewVoidBlock(block)
	rv := objc.Send[objc.ID](objc.ID(_SLSGUIClientClass.class), objc.Sel("configGUIClient:error:notifyQueue:notificationType:notificationBlock:"), gUIClient, error_, queue, type_, _block4)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSGUIClient/service
func (s SLSGUIClient) Service() ISLSXPCService {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("service"))
	return SLSXPCServiceFromID(objc.ID(rv))
}
func (s SLSGUIClient) SetService(value ISLSXPCService) {
	objc.Send[struct{}](s.ID, objc.Sel("setService:"), value)
}
