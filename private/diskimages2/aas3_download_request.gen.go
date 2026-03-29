// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AAS3DownloadRequest] class.
var (
	_AAS3DownloadRequestClass     AAS3DownloadRequestClass
	_AAS3DownloadRequestClassOnce sync.Once
)

func getAAS3DownloadRequestClass() AAS3DownloadRequestClass {
	_AAS3DownloadRequestClassOnce.Do(func() {
		_AAS3DownloadRequestClass = AAS3DownloadRequestClass{class: objc.GetClass("AAS3DownloadRequest")}
	})
	return _AAS3DownloadRequestClass
}

// GetAAS3DownloadRequestClass returns the class object for AAS3DownloadRequest.
func GetAAS3DownloadRequestClass() AAS3DownloadRequestClass {
	return getAAS3DownloadRequestClass()
}

type AAS3DownloadRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AAS3DownloadRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AAS3DownloadRequestClass) Alloc() AAS3DownloadRequest {
	rv := objc.Send[AAS3DownloadRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AAS3DownloadRequest.Buf]
//   - [AAS3DownloadRequest.SetBuf]
//   - [AAS3DownloadRequest.CreateAndResumeTask]
//   - [AAS3DownloadRequest.DownloadSession]
//   - [AAS3DownloadRequest.SetDownloadSession]
//   - [AAS3DownloadRequest.Nbyte]
//   - [AAS3DownloadRequest.SetNbyte]
//   - [AAS3DownloadRequest.Offset]
//   - [AAS3DownloadRequest.SetOffset]
//   - [AAS3DownloadRequest.PauseInterval]
//   - [AAS3DownloadRequest.SetPauseInterval]
//   - [AAS3DownloadRequest.RemainingAttempts]
//   - [AAS3DownloadRequest.SetRemainingAttempts]
//   - [AAS3DownloadRequest.Sem]
//   - [AAS3DownloadRequest.SetSem]
//   - [AAS3DownloadRequest.Status]
//   - [AAS3DownloadRequest.SetStatus]
//   - [AAS3DownloadRequest.Stream]
//   - [AAS3DownloadRequest.SetStream]
//   - [AAS3DownloadRequest.UrlRequest]
//   - [AAS3DownloadRequest.SetUrlRequest]
//   - [AAS3DownloadRequest.InitWithSessionSizeAtOffsetDestinationBufferDestinationStreamCompletionSemaphore]
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest
type AAS3DownloadRequest struct {
	objectivec.Object
}

// AAS3DownloadRequestFromID constructs a [AAS3DownloadRequest] from an objc.ID.
func AAS3DownloadRequestFromID(id objc.ID) AAS3DownloadRequest {
	return AAS3DownloadRequest{objectivec.Object{ID: id}}
}
// Ensure AAS3DownloadRequest implements IAAS3DownloadRequest.
var _ IAAS3DownloadRequest = AAS3DownloadRequest{}

// An interface definition for the [AAS3DownloadRequest] class.
//
// # Methods
//
//   - [IAAS3DownloadRequest.Buf]
//   - [IAAS3DownloadRequest.SetBuf]
//   - [IAAS3DownloadRequest.CreateAndResumeTask]
//   - [IAAS3DownloadRequest.DownloadSession]
//   - [IAAS3DownloadRequest.SetDownloadSession]
//   - [IAAS3DownloadRequest.Nbyte]
//   - [IAAS3DownloadRequest.SetNbyte]
//   - [IAAS3DownloadRequest.Offset]
//   - [IAAS3DownloadRequest.SetOffset]
//   - [IAAS3DownloadRequest.PauseInterval]
//   - [IAAS3DownloadRequest.SetPauseInterval]
//   - [IAAS3DownloadRequest.RemainingAttempts]
//   - [IAAS3DownloadRequest.SetRemainingAttempts]
//   - [IAAS3DownloadRequest.Sem]
//   - [IAAS3DownloadRequest.SetSem]
//   - [IAAS3DownloadRequest.Status]
//   - [IAAS3DownloadRequest.SetStatus]
//   - [IAAS3DownloadRequest.Stream]
//   - [IAAS3DownloadRequest.SetStream]
//   - [IAAS3DownloadRequest.UrlRequest]
//   - [IAAS3DownloadRequest.SetUrlRequest]
//   - [IAAS3DownloadRequest.InitWithSessionSizeAtOffsetDestinationBufferDestinationStreamCompletionSemaphore]
//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest
type IAAS3DownloadRequest interface {
	objectivec.IObject

	// Topic: Methods

	Buf() string
	SetBuf(value string)
	CreateAndResumeTask() int
	DownloadSession() IAAS3DownloadSession
	SetDownloadSession(value IAAS3DownloadSession)
	Nbyte() uint64
	SetNbyte(value uint64)
	Offset() int64
	SetOffset(value int64)
	PauseInterval() float32
	SetPauseInterval(value float32)
	RemainingAttempts() uint32
	SetRemainingAttempts(value uint32)
	Sem() objectivec.Object
	SetSem(value objectivec.Object)
	Status() int
	SetStatus(value int)
	Stream() unsafe.Pointer
	SetStream(value unsafe.Pointer)
	UrlRequest() foundation.NSMutableURLRequest
	SetUrlRequest(value foundation.NSMutableURLRequest)
	InitWithSessionSizeAtOffsetDestinationBufferDestinationStreamCompletionSemaphore(session objectivec.IObject, size uint64, offset int64, buffer string, stream unsafe.Pointer, semaphore objectivec.IObject) AAS3DownloadRequest
}

// Init initializes the instance.
func (a AAS3DownloadRequest) Init() AAS3DownloadRequest {
	rv := objc.Send[AAS3DownloadRequest](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AAS3DownloadRequest) Autorelease() AAS3DownloadRequest {
	rv := objc.Send[AAS3DownloadRequest](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAAS3DownloadRequest creates a new AAS3DownloadRequest instance.
func NewAAS3DownloadRequest() AAS3DownloadRequest {
	class := getAAS3DownloadRequestClass()
	rv := objc.Send[AAS3DownloadRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/initWithSession:size:atOffset:destinationBuffer:destinationStream:completionSemaphore:
func NewAAS3DownloadRequestWithSessionSizeAtOffsetDestinationBufferDestinationStreamCompletionSemaphore(session objectivec.IObject, size uint64, offset int64, buffer string, stream unsafe.Pointer, semaphore objectivec.IObject) AAS3DownloadRequest {
	instance := getAAS3DownloadRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:size:atOffset:destinationBuffer:destinationStream:completionSemaphore:"), session, size, offset, unsafe.Pointer(unsafe.StringData(buffer + "\x00")), stream, semaphore)
	return AAS3DownloadRequestFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/createAndResumeTask
func (a AAS3DownloadRequest) CreateAndResumeTask() int {
	rv := objc.Send[int](a.ID, objc.Sel("createAndResumeTask"))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/initWithSession:size:atOffset:destinationBuffer:destinationStream:completionSemaphore:
func (a AAS3DownloadRequest) InitWithSessionSizeAtOffsetDestinationBufferDestinationStreamCompletionSemaphore(session objectivec.IObject, size uint64, offset int64, buffer string, stream unsafe.Pointer, semaphore objectivec.IObject) AAS3DownloadRequest {
	rv := objc.Send[AAS3DownloadRequest](a.ID, objc.Sel("initWithSession:size:atOffset:destinationBuffer:destinationStream:completionSemaphore:"), session, size, offset, unsafe.Pointer(unsafe.StringData(buffer + "\x00")), stream, semaphore)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/buf
func (a AAS3DownloadRequest) Buf() string {
	rv := objc.Send[*byte](a.ID, objc.Sel("buf"))
	return objc.GoString(rv)
}
func (a AAS3DownloadRequest) SetBuf(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setBuf:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/downloadSession
func (a AAS3DownloadRequest) DownloadSession() IAAS3DownloadSession {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("downloadSession"))
	return AAS3DownloadSessionFromID(objc.ID(rv))
}
func (a AAS3DownloadRequest) SetDownloadSession(value IAAS3DownloadSession) {
	objc.Send[struct{}](a.ID, objc.Sel("setDownloadSession:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/nbyte
func (a AAS3DownloadRequest) Nbyte() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("nbyte"))
	return rv
}
func (a AAS3DownloadRequest) SetNbyte(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setNbyte:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/offset
func (a AAS3DownloadRequest) Offset() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("offset"))
	return rv
}
func (a AAS3DownloadRequest) SetOffset(value int64) {
	objc.Send[struct{}](a.ID, objc.Sel("setOffset:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/pauseInterval
func (a AAS3DownloadRequest) PauseInterval() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pauseInterval"))
	return rv
}
func (a AAS3DownloadRequest) SetPauseInterval(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPauseInterval:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/remainingAttempts
func (a AAS3DownloadRequest) RemainingAttempts() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("remainingAttempts"))
	return rv
}
func (a AAS3DownloadRequest) SetRemainingAttempts(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setRemainingAttempts:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/sem
func (a AAS3DownloadRequest) Sem() objectivec.Object {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sem"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (a AAS3DownloadRequest) SetSem(value objectivec.Object) {
	objc.Send[struct{}](a.ID, objc.Sel("setSem:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/status
func (a AAS3DownloadRequest) Status() int {
	rv := objc.Send[int](a.ID, objc.Sel("status"))
	return rv
}
func (a AAS3DownloadRequest) SetStatus(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setStatus:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/stream
func (a AAS3DownloadRequest) Stream() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("stream"))
	return rv
}
func (a AAS3DownloadRequest) SetStream(value unsafe.Pointer) {
	objc.Send[struct{}](a.ID, objc.Sel("setStream:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadRequest/urlRequest
func (a AAS3DownloadRequest) UrlRequest() foundation.NSMutableURLRequest {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("urlRequest"))
	return foundation.NSMutableURLRequestFromID(objc.ID(rv))
}
func (a AAS3DownloadRequest) SetUrlRequest(value foundation.NSMutableURLRequest) {
	objc.Send[struct{}](a.ID, objc.Sel("setUrlRequest:"), value)
}

