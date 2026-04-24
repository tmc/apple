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

// The class instance for the [SLSDisplayControlClient] class.
var (
	_SLSDisplayControlClientClass     SLSDisplayControlClientClass
	_SLSDisplayControlClientClassOnce sync.Once
)

func getSLSDisplayControlClientClass() SLSDisplayControlClientClass {
	_SLSDisplayControlClientClassOnce.Do(func() {
		_SLSDisplayControlClientClass = SLSDisplayControlClientClass{class: objc.GetClass("SLSDisplayControlClient")}
	})
	return _SLSDisplayControlClientClass
}

// GetSLSDisplayControlClientClass returns the class object for SLSDisplayControlClient.
func GetSLSDisplayControlClientClass() SLSDisplayControlClientClass {
	return getSLSDisplayControlClientClass()
}

type SLSDisplayControlClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSDisplayControlClientClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSDisplayControlClientClass) Alloc() SLSDisplayControlClient {
	rv := objc.Send[SLSDisplayControlClient](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSDisplayControlClient.AddNotificationTypeTo]
//   - [SLSDisplayControlClient.AddPayloadOfTypeTo]
//   - [SLSDisplayControlClient.AddUUIDTo]
//   - [SLSDisplayControlClient.CfStringToCStringPtr]
//   - [SLSDisplayControlClient.Configured]
//   - [SLSDisplayControlClient.SetConfigured]
//   - [SLSDisplayControlClient.CreateNSErrorWithCGError]
//   - [SLSDisplayControlClient.DecodeNotificationNotifyTypeUuidPayloadTypePayload]
//   - [SLSDisplayControlClient.Enabled]
//   - [SLSDisplayControlClient.SetEnabled]
//   - [SLSDisplayControlClient.EncodeCommandWithUUIDPayloadTypePayload]
//   - [SLSDisplayControlClient.IsTypeOfClassAClassError]
//   - [SLSDisplayControlClient.PayloadTypeToCFType]
//   - [SLSDisplayControlClient.RegisterClientPortNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClient.RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClient.RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [SLSDisplayControlClient.Semaphore]
//   - [SLSDisplayControlClient.SetSemaphore]
//   - [SLSDisplayControlClient.SemaphoreSignal]
//   - [SLSDisplayControlClient.SemaphoreWait]
//   - [SLSDisplayControlClient.SetNotification]
//   - [SLSDisplayControlClient.TerminateConnection]
//   - [SLSDisplayControlClient.XpcEventToNotification]
//   - [SLSDisplayControlClient.DebugDescription]
//   - [SLSDisplayControlClient.Description]
//   - [SLSDisplayControlClient.Hash]
//   - [SLSDisplayControlClient.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient
type SLSDisplayControlClient struct {
	objectivec.Object
}

// SLSDisplayControlClientFromID constructs a [SLSDisplayControlClient] from an objc.ID.
func SLSDisplayControlClientFromID(id objc.ID) SLSDisplayControlClient {
	return SLSDisplayControlClient{objectivec.Object{ID: id}}
}

// Ensure SLSDisplayControlClient implements ISLSDisplayControlClient.
var _ ISLSDisplayControlClient = SLSDisplayControlClient{}

// An interface definition for the [SLSDisplayControlClient] class.
//
// # Methods
//
//   - [ISLSDisplayControlClient.AddNotificationTypeTo]
//   - [ISLSDisplayControlClient.AddPayloadOfTypeTo]
//   - [ISLSDisplayControlClient.AddUUIDTo]
//   - [ISLSDisplayControlClient.CfStringToCStringPtr]
//   - [ISLSDisplayControlClient.Configured]
//   - [ISLSDisplayControlClient.SetConfigured]
//   - [ISLSDisplayControlClient.CreateNSErrorWithCGError]
//   - [ISLSDisplayControlClient.DecodeNotificationNotifyTypeUuidPayloadTypePayload]
//   - [ISLSDisplayControlClient.Enabled]
//   - [ISLSDisplayControlClient.SetEnabled]
//   - [ISLSDisplayControlClient.EncodeCommandWithUUIDPayloadTypePayload]
//   - [ISLSDisplayControlClient.IsTypeOfClassAClassError]
//   - [ISLSDisplayControlClient.PayloadTypeToCFType]
//   - [ISLSDisplayControlClient.RegisterClientPortNotifyQueueNotificationTypeNotificationBlock]
//   - [ISLSDisplayControlClient.RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [ISLSDisplayControlClient.RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock]
//   - [ISLSDisplayControlClient.Semaphore]
//   - [ISLSDisplayControlClient.SetSemaphore]
//   - [ISLSDisplayControlClient.SemaphoreSignal]
//   - [ISLSDisplayControlClient.SemaphoreWait]
//   - [ISLSDisplayControlClient.SetNotification]
//   - [ISLSDisplayControlClient.TerminateConnection]
//   - [ISLSDisplayControlClient.XpcEventToNotification]
//   - [ISLSDisplayControlClient.DebugDescription]
//   - [ISLSDisplayControlClient.Description]
//   - [ISLSDisplayControlClient.Hash]
//   - [ISLSDisplayControlClient.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient
type ISLSDisplayControlClient interface {
	objectivec.IObject

	// Topic: Methods

	AddNotificationTypeTo(type_ uint64, to objectivec.IObject)
	AddPayloadOfTypeTo(payload objectivec.IObject, type_ uint64, to objectivec.IObject)
	AddUUIDTo(uuid uint64, to objectivec.IObject)
	CfStringToCStringPtr(ptr objectivec.IObject) string
	Configured() bool
	SetConfigured(value bool)
	CreateNSErrorWithCGError(nSError []objectivec.IObject, cGError int64)
	DecodeNotificationNotifyTypeUuidPayloadTypePayload(notification objectivec.IObject, type_ unsafe.Pointer, uuid unsafe.Pointer, type_2 unsafe.Pointer, payload []objectivec.IObject)
	Enabled() bool
	SetEnabled(value bool)
	EncodeCommandWithUUIDPayloadTypePayload(command uint64, uuid unsafe.Pointer, type_ uint64, payload objectivec.IObject) objectivec.IObject
	IsTypeOfClassAClassError(class objectivec.IObject, class2 objc.Class) (int, bool)
	PayloadTypeToCFType(cFType uint64) uint64
	RegisterClientPortNotifyQueueNotificationTypeNotificationBlock(client []objectivec.IObject, port uint32, queue objectivec.IObject, type_ uint64, block VoidHandler)
	RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock(client objectivec.IObject, autoreconnect bool, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) objectivec.IObject
	RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock(gUIClient objectivec.IObject, port uint32, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) objectivec.IObject
	Semaphore() objectivec.Object
	SetSemaphore(value objectivec.Object)
	SemaphoreSignal()
	SemaphoreWait(wait byte) int
	SetNotification(notification VoidHandler)
	TerminateConnection()
	XpcEventToNotification(notification objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s SLSDisplayControlClient) Init() SLSDisplayControlClient {
	rv := objc.Send[SLSDisplayControlClient](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSDisplayControlClient) Autorelease() SLSDisplayControlClient {
	rv := objc.Send[SLSDisplayControlClient](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSDisplayControlClient creates a new SLSDisplayControlClient instance.
func NewSLSDisplayControlClient() SLSDisplayControlClient {
	class := getSLSDisplayControlClientClass()
	rv := objc.Send[SLSDisplayControlClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/addNotificationType:to:
func (s SLSDisplayControlClient) AddNotificationTypeTo(type_ uint64, to objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("addNotificationType:to:"), type_, to)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/addPayload:ofType:to:
func (s SLSDisplayControlClient) AddPayloadOfTypeTo(payload objectivec.IObject, type_ uint64, to objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("addPayload:ofType:to:"), payload, type_, to)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/addUUID:to:
func (s SLSDisplayControlClient) AddUUIDTo(uuid uint64, to objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("addUUID:to:"), uuid, to)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/cfStringToCStringPtr:
func (s SLSDisplayControlClient) CfStringToCStringPtr(ptr objectivec.IObject) string {
	rv := objc.Send[*byte](s.ID, objc.Sel("cfStringToCStringPtr:"), ptr)
	return objc.GoString(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/createNSError:withCGError:
func (s SLSDisplayControlClient) CreateNSErrorWithCGError(nSError []objectivec.IObject, cGError int64) {
	objc.Send[objc.ID](s.ID, objc.Sel("createNSError:withCGError:"), objectivec.IObjectSliceToNSArray(nSError), cGError)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/decodeNotification:notifyType:uuid:payloadType:payload:
func (s SLSDisplayControlClient) DecodeNotificationNotifyTypeUuidPayloadTypePayload(notification objectivec.IObject, type_ unsafe.Pointer, uuid unsafe.Pointer, type_2 unsafe.Pointer, payload []objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("decodeNotification:notifyType:uuid:payloadType:payload:"), notification, type_, uuid, type_2, objectivec.IObjectSliceToNSArray(payload))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/encodeCommand:withUUID:payloadType:payload:
func (s SLSDisplayControlClient) EncodeCommandWithUUIDPayloadTypePayload(command uint64, uuid unsafe.Pointer, type_ uint64, payload objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("encodeCommand:withUUID:payloadType:payload:"), command, uuid, type_, payload)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/isTypeOfClass:aClass:error:
func (s SLSDisplayControlClient) IsTypeOfClassAClassError(class objectivec.IObject, class2 objc.Class) (int, bool) {
	var error_ int
	rv := objc.Send[bool](s.ID, objc.Sel("isTypeOfClass:aClass:error:"), class, class2, unsafe.Pointer(&error_))
	return error_, rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/payloadTypeToCFType:
func (s SLSDisplayControlClient) PayloadTypeToCFType(cFType uint64) uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("payloadTypeToCFType:"), cFType)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/registerClient:port:notifyQueue:notificationType:notificationBlock:
func (s SLSDisplayControlClient) RegisterClientPortNotifyQueueNotificationTypeNotificationBlock(client []objectivec.IObject, port uint32, queue objectivec.IObject, type_ uint64, block VoidHandler) {
	_block4, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("registerClient:port:notifyQueue:notificationType:notificationBlock:"), client, port, queue, type_, _block4)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/registerDaemonClient:withAutoreconnect:error:notifyQueue:notificationType:notificationBlock:
func (s SLSDisplayControlClient) RegisterDaemonClientWithAutoreconnectErrorNotifyQueueNotificationTypeNotificationBlock(client objectivec.IObject, autoreconnect bool, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) objectivec.IObject {
	_block5, _ := NewVoidBlock(block)
	rv := objc.Send[objc.ID](s.ID, objc.Sel("registerDaemonClient:withAutoreconnect:error:notifyQueue:notificationType:notificationBlock:"), client, autoreconnect, error_, queue, type_, _block5)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/registerGUIClient:connectionPort:error:notifyQueue:notificationType:notificationBlock:
func (s SLSDisplayControlClient) RegisterGUIClientConnectionPortErrorNotifyQueueNotificationTypeNotificationBlock(gUIClient objectivec.IObject, port uint32, error_ []objectivec.IObject, queue objectivec.IObject, type_ uint64, block VoidHandler) objectivec.IObject {
	_block5, _ := NewVoidBlock(block)
	rv := objc.Send[objc.ID](s.ID, objc.Sel("registerGUIClient:connectionPort:error:notifyQueue:notificationType:notificationBlock:"), gUIClient, port, error_, queue, type_, _block5)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/semaphoreSignal
func (s SLSDisplayControlClient) SemaphoreSignal() {
	objc.Send[objc.ID](s.ID, objc.Sel("semaphoreSignal"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/semaphoreWait:
func (s SLSDisplayControlClient) SemaphoreWait(wait byte) int {
	rv := objc.Send[int](s.ID, objc.Sel("semaphoreWait:"), wait)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/setNotification:
func (s SLSDisplayControlClient) SetNotification(notification VoidHandler) {
	_block0, _ := NewVoidBlock(notification)
	objc.Send[objc.ID](s.ID, objc.Sel("setNotification:"), _block0)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/terminateConnection
func (s SLSDisplayControlClient) TerminateConnection() {
	objc.Send[objc.ID](s.ID, objc.Sel("terminateConnection"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/xpcEventToNotification:
func (s SLSDisplayControlClient) XpcEventToNotification(notification objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("xpcEventToNotification:"), notification)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/configured
func (s SLSDisplayControlClient) Configured() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("configured"))
	return rv
}
func (s SLSDisplayControlClient) SetConfigured(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setConfigured:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/debugDescription
func (s SLSDisplayControlClient) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/description
func (s SLSDisplayControlClient) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/enabled
func (s SLSDisplayControlClient) Enabled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("enabled"))
	return rv
}
func (s SLSDisplayControlClient) SetEnabled(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setEnabled:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/hash
func (s SLSDisplayControlClient) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/semaphore
func (s SLSDisplayControlClient) Semaphore() objectivec.Object {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("semaphore"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s SLSDisplayControlClient) SetSemaphore(value objectivec.Object) {
	objc.Send[struct{}](s.ID, objc.Sel("setSemaphore:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayControlClient/superclass
func (s SLSDisplayControlClient) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}

// SetNotificationSync is a synchronous wrapper around [SLSDisplayControlClient.SetNotification].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSDisplayControlClient) SetNotificationSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.SetNotification(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
