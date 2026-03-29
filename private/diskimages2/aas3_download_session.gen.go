// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AAS3DownloadSession] class.
var (
	_AAS3DownloadSessionClass     AAS3DownloadSessionClass
	_AAS3DownloadSessionClassOnce sync.Once
)

func getAAS3DownloadSessionClass() AAS3DownloadSessionClass {
	_AAS3DownloadSessionClassOnce.Do(func() {
		_AAS3DownloadSessionClass = AAS3DownloadSessionClass{class: objc.GetClass("AAS3DownloadSession")}
	})
	return _AAS3DownloadSessionClass
}

// GetAAS3DownloadSessionClass returns the class object for AAS3DownloadSession.
func GetAAS3DownloadSessionClass() AAS3DownloadSessionClass {
	return getAAS3DownloadSessionClass()
}

type AAS3DownloadSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AAS3DownloadSessionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AAS3DownloadSessionClass) Alloc() AAS3DownloadSession {
	rv := objc.Send[AAS3DownloadSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AAS3DownloadSession.AddBytesDownloaded]
//   - [AAS3DownloadSession.AddRequest]
//   - [AAS3DownloadSession.BytesDownloaded]
//   - [AAS3DownloadSession.SetBytesDownloaded]
//   - [AAS3DownloadSession.Cache]
//   - [AAS3DownloadSession.SetCache]
//   - [AAS3DownloadSession.CacheDocument]
//   - [AAS3DownloadSession.CacheLock]
//   - [AAS3DownloadSession.Cancelled]
//   - [AAS3DownloadSession.SetCancelled]
//   - [AAS3DownloadSession.EnqueueRequestWithSizeAtOffsetDestinationBufferDestinationStreamCompletionSemaphore]
//   - [AAS3DownloadSession.InvalidateAndCancel]
//   - [AAS3DownloadSession.IsCancelled]
//   - [AAS3DownloadSession.MaxAttempts]
//   - [AAS3DownloadSession.MaxRequests]
//   - [AAS3DownloadSession.PauseInterval]
//   - [AAS3DownloadSession.ReadToAsyncByteStreamSizeAtOffset]
//   - [AAS3DownloadSession.ReadToBufferSizeAtOffset]
//   - [AAS3DownloadSession.RemoveRequest]
//   - [AAS3DownloadSession.Requests]
//   - [AAS3DownloadSession.RequestsLock]
//   - [AAS3DownloadSession.RequestsSem]
//   - [AAS3DownloadSession.StreamBase]
//   - [AAS3DownloadSession.SyncRequests]
//   - [AAS3DownloadSession.Url]
//   - [AAS3DownloadSession.UrlSession]
//   - [AAS3DownloadSession.InitWithURLStreamBaseMaxAttemptsPauseIntervalMaxRequestsInFlight]
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession
type AAS3DownloadSession struct {
	objectivec.Object
}

// AAS3DownloadSessionFromID constructs a [AAS3DownloadSession] from an objc.ID.
func AAS3DownloadSessionFromID(id objc.ID) AAS3DownloadSession {
	return AAS3DownloadSession{objectivec.Object{ID: id}}
}
// Ensure AAS3DownloadSession implements IAAS3DownloadSession.
var _ IAAS3DownloadSession = AAS3DownloadSession{}

// An interface definition for the [AAS3DownloadSession] class.
//
// # Methods
//
//   - [IAAS3DownloadSession.AddBytesDownloaded]
//   - [IAAS3DownloadSession.AddRequest]
//   - [IAAS3DownloadSession.BytesDownloaded]
//   - [IAAS3DownloadSession.SetBytesDownloaded]
//   - [IAAS3DownloadSession.Cache]
//   - [IAAS3DownloadSession.SetCache]
//   - [IAAS3DownloadSession.CacheDocument]
//   - [IAAS3DownloadSession.CacheLock]
//   - [IAAS3DownloadSession.Cancelled]
//   - [IAAS3DownloadSession.SetCancelled]
//   - [IAAS3DownloadSession.EnqueueRequestWithSizeAtOffsetDestinationBufferDestinationStreamCompletionSemaphore]
//   - [IAAS3DownloadSession.InvalidateAndCancel]
//   - [IAAS3DownloadSession.IsCancelled]
//   - [IAAS3DownloadSession.MaxAttempts]
//   - [IAAS3DownloadSession.MaxRequests]
//   - [IAAS3DownloadSession.PauseInterval]
//   - [IAAS3DownloadSession.ReadToAsyncByteStreamSizeAtOffset]
//   - [IAAS3DownloadSession.ReadToBufferSizeAtOffset]
//   - [IAAS3DownloadSession.RemoveRequest]
//   - [IAAS3DownloadSession.Requests]
//   - [IAAS3DownloadSession.RequestsLock]
//   - [IAAS3DownloadSession.RequestsSem]
//   - [IAAS3DownloadSession.StreamBase]
//   - [IAAS3DownloadSession.SyncRequests]
//   - [IAAS3DownloadSession.Url]
//   - [IAAS3DownloadSession.UrlSession]
//   - [IAAS3DownloadSession.InitWithURLStreamBaseMaxAttemptsPauseIntervalMaxRequestsInFlight]
//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession
type IAAS3DownloadSession interface {
	objectivec.IObject

	// Topic: Methods

	AddBytesDownloaded(downloaded uint64)
	AddRequest(request objectivec.IObject) int
	BytesDownloaded() objectivec.IObject
	SetBytesDownloaded(value objectivec.IObject)
	Cache() foundation.INSData
	SetCache(value foundation.INSData)
	CacheDocument(document objectivec.IObject)
	CacheLock() foundation.NSLock
	Cancelled() objectivec.IObject
	SetCancelled(value objectivec.IObject)
	EnqueueRequestWithSizeAtOffsetDestinationBufferDestinationStreamCompletionSemaphore(size uint64, offset int64, buffer string, stream unsafe.Pointer, semaphore objectivec.IObject) objectivec.IObject
	InvalidateAndCancel()
	IsCancelled() int
	MaxAttempts() uint32
	MaxRequests() uint32
	PauseInterval() float32
	ReadToAsyncByteStreamSizeAtOffset(stream unsafe.Pointer, size uint64, offset int64) int
	ReadToBufferSizeAtOffset(buffer unsafe.Pointer, size uint64, offset int64) int64
	RemoveRequest(request objectivec.IObject)
	Requests() foundation.INSSet
	RequestsLock() foundation.NSLock
	RequestsSem() objectivec.Object
	StreamBase() objectivec.IObject
	SyncRequests() int
	Url() foundation.INSURL
	UrlSession() *foundation.NSURLSession
	InitWithURLStreamBaseMaxAttemptsPauseIntervalMaxRequestsInFlight(url foundation.INSURL, base objectivec.IObject, attempts uint32, interval float32, flight uint32) AAS3DownloadSession
}

// Init initializes the instance.
func (a AAS3DownloadSession) Init() AAS3DownloadSession {
	rv := objc.Send[AAS3DownloadSession](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AAS3DownloadSession) Autorelease() AAS3DownloadSession {
	rv := objc.Send[AAS3DownloadSession](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAAS3DownloadSession creates a new AAS3DownloadSession instance.
func NewAAS3DownloadSession() AAS3DownloadSession {
	class := getAAS3DownloadSessionClass()
	rv := objc.Send[AAS3DownloadSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/initWithURL:streamBase:maxAttempts:pauseInterval:maxRequestsInFlight:
func NewAAS3DownloadSessionWithURLStreamBaseMaxAttemptsPauseIntervalMaxRequestsInFlight(url foundation.INSURL, base objectivec.IObject, attempts uint32, interval float32, flight uint32) AAS3DownloadSession {
	instance := getAAS3DownloadSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:streamBase:maxAttempts:pauseInterval:maxRequestsInFlight:"), url, base, attempts, interval, flight)
	return AAS3DownloadSessionFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/addBytesDownloaded:
func (a AAS3DownloadSession) AddBytesDownloaded(downloaded uint64) {
	objc.Send[objc.ID](a.ID, objc.Sel("addBytesDownloaded:"), downloaded)
}
//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/addRequest:
func (a AAS3DownloadSession) AddRequest(request objectivec.IObject) int {
	rv := objc.Send[int](a.ID, objc.Sel("addRequest:"), request)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/cacheDocument:
func (a AAS3DownloadSession) CacheDocument(document objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("cacheDocument:"), document)
}
//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/enqueueRequestWithSize:atOffset:destinationBuffer:destinationStream:completionSemaphore:
func (a AAS3DownloadSession) EnqueueRequestWithSizeAtOffsetDestinationBufferDestinationStreamCompletionSemaphore(size uint64, offset int64, buffer string, stream unsafe.Pointer, semaphore objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("enqueueRequestWithSize:atOffset:destinationBuffer:destinationStream:completionSemaphore:"), size, offset, unsafe.Pointer(unsafe.StringData(buffer + "\x00")), stream, semaphore)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/invalidateAndCancel
func (a AAS3DownloadSession) InvalidateAndCancel() {
	objc.Send[objc.ID](a.ID, objc.Sel("invalidateAndCancel"))
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/isCancelled
func (a AAS3DownloadSession) IsCancelled() int {
	rv := objc.Send[int](a.ID, objc.Sel("isCancelled"))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/readToAsyncByteStream:size:atOffset:
func (a AAS3DownloadSession) ReadToAsyncByteStreamSizeAtOffset(stream unsafe.Pointer, size uint64, offset int64) int {
	rv := objc.Send[int](a.ID, objc.Sel("readToAsyncByteStream:size:atOffset:"), stream, size, offset)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/readToBuffer:size:atOffset:
func (a AAS3DownloadSession) ReadToBufferSizeAtOffset(buffer unsafe.Pointer, size uint64, offset int64) int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("readToBuffer:size:atOffset:"), buffer, size, offset)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/removeRequest:
func (a AAS3DownloadSession) RemoveRequest(request objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeRequest:"), request)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/syncRequests
func (a AAS3DownloadSession) SyncRequests() int {
	rv := objc.Send[int](a.ID, objc.Sel("syncRequests"))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/initWithURL:streamBase:maxAttempts:pauseInterval:maxRequestsInFlight:
func (a AAS3DownloadSession) InitWithURLStreamBaseMaxAttemptsPauseIntervalMaxRequestsInFlight(url foundation.INSURL, base objectivec.IObject, attempts uint32, interval float32, flight uint32) AAS3DownloadSession {
	rv := objc.Send[AAS3DownloadSession](a.ID, objc.Sel("initWithURL:streamBase:maxAttempts:pauseInterval:maxRequestsInFlight:"), url, base, attempts, interval, flight)
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/completeRequest:data:response:error:
func (_AAS3DownloadSessionClass AAS3DownloadSessionClass) CompleteRequestDataResponseError(request objectivec.IObject, data objectivec.IObject, response objectivec.IObject, error_ objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_AAS3DownloadSessionClass.class), objc.Sel("completeRequest:data:response:error:"), request, data, response, error_)
}

// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/bytesDownloaded
func (a AAS3DownloadSession) BytesDownloaded() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("bytesDownloaded"))
	return objectivec.Object{ID: rv}
}
func (a AAS3DownloadSession) SetBytesDownloaded(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setBytesDownloaded:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/cache
func (a AAS3DownloadSession) Cache() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cache"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (a AAS3DownloadSession) SetCache(value foundation.INSData) {
	objc.Send[struct{}](a.ID, objc.Sel("setCache:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/cacheLock
func (a AAS3DownloadSession) CacheLock() foundation.NSLock {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cacheLock"))
	return foundation.NSLockFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/cancelled
func (a AAS3DownloadSession) Cancelled() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cancelled"))
	return objectivec.Object{ID: rv}
}
func (a AAS3DownloadSession) SetCancelled(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setCancelled:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/maxAttempts
func (a AAS3DownloadSession) MaxAttempts() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("maxAttempts"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/maxRequests
func (a AAS3DownloadSession) MaxRequests() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("maxRequests"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/pauseInterval
func (a AAS3DownloadSession) PauseInterval() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pauseInterval"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/requests
func (a AAS3DownloadSession) Requests() foundation.INSSet {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("requests"))
	return foundation.NSSetFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/requestsLock
func (a AAS3DownloadSession) RequestsLock() foundation.NSLock {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("requestsLock"))
	return foundation.NSLockFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/requestsSem
func (a AAS3DownloadSession) RequestsSem() objectivec.Object {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("requestsSem"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/streamBase
func (a AAS3DownloadSession) StreamBase() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("streamBase"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/url
func (a AAS3DownloadSession) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/AAS3DownloadSession/urlSession
func (a AAS3DownloadSession) UrlSession() *foundation.NSURLSession {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("urlSession"))
	if rv == 0 {
		return nil
	}
	val := foundation.NSURLSessionFromID(objc.ID(rv))
	return &val
}

