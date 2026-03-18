// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLCache] class.
var (
	_URLCacheClass     URLCacheClass
	_URLCacheClassOnce sync.Once
)

func getURLCacheClass() URLCacheClass {
	_URLCacheClassOnce.Do(func() {
		_URLCacheClass = URLCacheClass{class: objc.GetClass("NSURLCache")}
	})
	return _URLCacheClass
}

// GetURLCacheClass returns the class object for NSURLCache.
func GetURLCacheClass() URLCacheClass {
	return getURLCacheClass()
}

type URLCacheClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLCacheClass) Alloc() URLCache {
	rv := objc.Send[URLCache](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// An object that maps URL requests to cached response objects.
//
// # Overview
// 
// The [NSURLCache] class implements the caching of responses to URL load
// requests, by mapping [NSURLRequest] objects to [NSCachedURLResponse]
// objects. It provides a composite in-memory and on-disk cache, and lets you
// manipulate the sizes of both the in-memory and on-disk portions. You can
// also control the path where cache data is persistently stored.
// 
// # Thread safety
// 
// In iOS 8 and later, and macOS 10.10 and later, [NSURLCache] is thread safe.
// 
// Although [NSURLCache] instance methods can safely be called from multiple
// execution contexts at the same time, be aware that methods like
// [CachedResponseForRequest] and [StoreCachedResponseForRequest] have an
// unavoidable race condition when attempting to read or write responses for
// the same request.
// 
// Subclasses of [NSURLCache] must implement overridden methods in such a
// thread-safe manner.
// 
// # Subclassing notes
// 
// The [NSURLCache] class is meant to be used as-is, but you can subclass it
// when you have specific needs. For example, you might want to screen which
// responses are cached, or reimplement the storage mechanism for security or
// other reasons.
// 
// When overriding methods of this class, be aware that methods that take a
// `task` parameter are preferred by the system to those that do not.
// Therefore, you should override the task-based methods when subclassing, as
// follows:
// 
// - Storing responses in the cache — Override the task-based
// [StoreCachedResponseForDataTask], instead of or in addition to the
// request-based [StoreCachedResponseForRequest]. - Getting responses from the
// cache — Override [GetCachedResponseForDataTaskCompletionHandler], instead
// of or in addition to [CachedResponseForRequest]. - Removing cached
// responses — Override the task-based [RemoveCachedResponseForDataTask],
// instead of or in addition to the request-based
// [RemoveCachedResponseForRequest].
//
// # Getting and storing cached objects
//
//   - [URLCache.CachedResponseForRequest]: Returns the cached URL response in the cache for the specified URL request.
//   - [URLCache.StoreCachedResponseForRequest]: Stores a cached URL response for a specified request.
//   - [URLCache.GetCachedResponseForDataTaskCompletionHandler]: Gets the cached URL response for a data task, passing it to the provided completion handler.
//   - [URLCache.StoreCachedResponseForDataTask]: Stores a cached URL response for a specified data task.
//
// # Removing cached objects
//
//   - [URLCache.RemoveCachedResponseForRequest]: Removes the cached URL response for a specified URL request.
//   - [URLCache.RemoveCachedResponseForDataTask]: Removes the cached URL response for a specified data task.
//   - [URLCache.RemoveCachedResponsesSinceDate]: Clears the given cache of any cached responses since the provided date.
//   - [URLCache.RemoveAllCachedResponses]: Clears the receiver’s cache, removing all stored cached URL responses.
//
// # Getting and setting on-disk cache properties
//
//   - [URLCache.CurrentDiskUsage]: The current size of the on-disk cache, in bytes.
//   - [URLCache.DiskCapacity]: The capacity of the on-disk cache, in bytes.
//   - [URLCache.SetDiskCapacity]
//
// # Getting and setting in-memory cache properties
//
//   - [URLCache.CurrentMemoryUsage]: The current size of the in-memory cache, in bytes.
//   - [URLCache.MemoryCapacity]: The capacity of the in-memory cache, in bytes.
//   - [URLCache.SetMemoryCapacity]
//
// See: https://developer.apple.com/documentation/Foundation/URLCache
type URLCache struct {
	objectivec.Object
}

// URLCacheFromID constructs a [URLCache] from an objc.ID.
//
// An object that maps URL requests to cached response objects.
func URLCacheFromID(id objc.ID) URLCache {
	return NSURLCache{objectivec.Object{ID: id}}
}

// NSURLCacheFromID is an alias for [URLCacheFromID] for cross-framework compatibility.
func NSURLCacheFromID(id objc.ID) URLCache { return URLCacheFromID(id) }
// NOTE: URLCache adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLCache] class.
//
// # Getting and storing cached objects
//
//   - [IURLCache.CachedResponseForRequest]: Returns the cached URL response in the cache for the specified URL request.
//   - [IURLCache.StoreCachedResponseForRequest]: Stores a cached URL response for a specified request.
//   - [IURLCache.GetCachedResponseForDataTaskCompletionHandler]: Gets the cached URL response for a data task, passing it to the provided completion handler.
//   - [IURLCache.StoreCachedResponseForDataTask]: Stores a cached URL response for a specified data task.
//
// # Removing cached objects
//
//   - [IURLCache.RemoveCachedResponseForRequest]: Removes the cached URL response for a specified URL request.
//   - [IURLCache.RemoveCachedResponseForDataTask]: Removes the cached URL response for a specified data task.
//   - [IURLCache.RemoveCachedResponsesSinceDate]: Clears the given cache of any cached responses since the provided date.
//   - [IURLCache.RemoveAllCachedResponses]: Clears the receiver’s cache, removing all stored cached URL responses.
//
// # Getting and setting on-disk cache properties
//
//   - [IURLCache.CurrentDiskUsage]: The current size of the on-disk cache, in bytes.
//   - [IURLCache.DiskCapacity]: The capacity of the on-disk cache, in bytes.
//   - [IURLCache.SetDiskCapacity]
//
// # Getting and setting in-memory cache properties
//
//   - [IURLCache.CurrentMemoryUsage]: The current size of the in-memory cache, in bytes.
//   - [IURLCache.MemoryCapacity]: The capacity of the in-memory cache, in bytes.
//   - [IURLCache.SetMemoryCapacity]
//
// See: https://developer.apple.com/documentation/Foundation/URLCache
type IURLCache interface {
	objectivec.IObject

	// Topic: Getting and storing cached objects

	// Returns the cached URL response in the cache for the specified URL request.
	CachedResponseForRequest(request INSURLRequest) INSCachedURLResponse
	// Stores a cached URL response for a specified request.
	StoreCachedResponseForRequest(cachedResponse INSCachedURLResponse, request INSURLRequest)
	// Gets the cached URL response for a data task, passing it to the provided completion handler.
	GetCachedResponseForDataTaskCompletionHandler(dataTask INSURLSessionDataTask, completionHandler CachedURLResponseHandler)
	// Stores a cached URL response for a specified data task.
	StoreCachedResponseForDataTask(cachedResponse INSCachedURLResponse, dataTask INSURLSessionDataTask)

	// Topic: Removing cached objects

	// Removes the cached URL response for a specified URL request.
	RemoveCachedResponseForRequest(request INSURLRequest)
	// Removes the cached URL response for a specified data task.
	RemoveCachedResponseForDataTask(dataTask INSURLSessionDataTask)
	// Clears the given cache of any cached responses since the provided date.
	RemoveCachedResponsesSinceDate(date INSDate)
	// Clears the receiver’s cache, removing all stored cached URL responses.
	RemoveAllCachedResponses()

	// Topic: Getting and setting on-disk cache properties

	// The current size of the on-disk cache, in bytes.
	CurrentDiskUsage() uint
	// The capacity of the on-disk cache, in bytes.
	DiskCapacity() uint
	SetDiskCapacity(value uint)

	// Topic: Getting and setting in-memory cache properties

	// The current size of the in-memory cache, in bytes.
	CurrentMemoryUsage() uint
	// The capacity of the in-memory cache, in bytes.
	MemoryCapacity() uint
	SetMemoryCapacity(value uint)

	// Creates a URL cache object with the specified memory and disk capacities, in the specified directory.
	InitWithMemoryCapacityDiskCapacityDirectoryURL(memoryCapacity uint, diskCapacity uint, directoryURL INSURL) URLCache
}

// Init initializes the instance.
func (u URLCache) Init() URLCache {
	rv := objc.Send[URLCache](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLCache) Autorelease() URLCache {
	rv := objc.Send[URLCache](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLCache creates a new URLCache instance.
func NewURLCache() URLCache {
	class := getURLCacheClass()
	rv := objc.Send[URLCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a URL cache object with the specified memory and disk capacities,
// in the specified directory.
//
// memoryCapacity: The memory capacity of the cache, in bytes.
//
// diskCapacity: The disk capacity of the cache, in bytes.
//
// directoryURL: The path to an on-disk directory at which to store the on-disk cache. If
// `directory` is `nil`, the cache uses a default directory.
//
// # Discussion
// 
// A disk cache measured in the tens of megabytes should be acceptable in most
// cases.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLCache/initWithMemoryCapacity:diskCapacity:directoryURL:
func NewURLCacheWithMemoryCapacityDiskCapacityDirectoryURL(memoryCapacity uint, diskCapacity uint, directoryURL INSURL) URLCache {
	instance := getURLCacheClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMemoryCapacity:diskCapacity:directoryURL:"), memoryCapacity, diskCapacity, directoryURL)
	return URLCacheFromID(rv)
}

// Creates a URL cache object with the specified values.
//
// memoryCapacity: The memory capacity of the cache, in bytes.
//
// diskCapacity: The disk capacity of the cache, in bytes.
//
// path: In macOS, `path` is the location at which to store the on-disk cache.
// 
// In iOS, `path` is the name of a subdirectory of the application’s default
// cache directory in which to store the on-disk cache (the subdirectory is
// created if it does not exist).
//
// # Return Value
// 
// The initialized cache object.
//
// # Discussion
// 
// The returned cache instance is backed by disk, so you have more leeway when
// choosing the capacity for this kind of cache. A disk cache measured in the
// tens of megabytes should be acceptable in most cases.
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/init(memoryCapacity:diskCapacity:diskPath:)
func NewURLCacheWithMemoryCapacityDiskCapacityDiskPath(memoryCapacity uint, diskCapacity uint, path string) URLCache {
	instance := getURLCacheClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMemoryCapacity:diskCapacity:diskPath:"), memoryCapacity, diskCapacity, objc.String(path))
	return URLCacheFromID(rv)
}

// Returns the cached URL response in the cache for the specified URL request.
//
// request: The URL request whose cached response is desired.
//
// # Return Value
// 
// The cached URL response for `request`, or `nil` if no response has been
// cached.
//
// # Discussion
// 
// If you override this method, you should also override
// [GetCachedResponseForDataTaskCompletionHandler].
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/cachedResponse(for:)
func (u URLCache) CachedResponseForRequest(request INSURLRequest) INSCachedURLResponse {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("cachedResponseForRequest:"), request)
	return NSCachedURLResponseFromID(rv)
}

// Stores a cached URL response for a specified request.
//
// cachedResponse: The cached URL response to store.
//
// request: The request for which the cached URL response is being stored.
//
// # Discussion
// 
// If you override this method, you should also override
// [StoreCachedResponseForDataTask].
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/storeCachedResponse(_:for:)-7p7bl
func (u URLCache) StoreCachedResponseForRequest(cachedResponse INSCachedURLResponse, request INSURLRequest) {
	objc.Send[objc.ID](u.ID, objc.Sel("storeCachedResponse:forRequest:"), cachedResponse, request)
}

// Gets the cached URL response for a data task, passing it to the provided
// completion handler.
//
// dataTask: The data task whose cached URL response is desired.
//
// completionHandler: A completion handler that receives the cached URL response for the data
// task’s request, or `nil` if no response is found in the cache.
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/getCachedResponse(for:completionHandler:)
func (u URLCache) GetCachedResponseForDataTaskCompletionHandler(dataTask INSURLSessionDataTask, completionHandler CachedURLResponseHandler) {
_block1, _cleanup1 := NewCachedURLResponseBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](u.ID, objc.Sel("getCachedResponseForDataTask:completionHandler:"), dataTask, _block1)
}

// Stores a cached URL response for a specified data task.
//
// cachedResponse: The cached URL response to store for this data task.
//
// dataTask: The data task whose response is to be cached.
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/storeCachedResponse(_:for:)-8uq91
func (u URLCache) StoreCachedResponseForDataTask(cachedResponse INSCachedURLResponse, dataTask INSURLSessionDataTask) {
	objc.Send[objc.ID](u.ID, objc.Sel("storeCachedResponse:forDataTask:"), cachedResponse, dataTask)
}

// Removes the cached URL response for a specified URL request.
//
// request: The URL request whose cached URL response should be removed. If there is no
// corresponding cached URL response, no action is taken.
//
// # Discussion
// 
// If you override this method, you should also override
// [RemoveCachedResponseForDataTask].
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/removeCachedResponse(for:)-1dh89
func (u URLCache) RemoveCachedResponseForRequest(request INSURLRequest) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeCachedResponseForRequest:"), request)
}

// Removes the cached URL response for a specified data task.
//
// dataTask: A task whose URL request’s corresponding cached URL response should be
// removed. If there is no corresponding cached URL response, no action is
// taken.
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/removeCachedResponse(for:)-1zwp6
func (u URLCache) RemoveCachedResponseForDataTask(dataTask INSURLSessionDataTask) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeCachedResponseForDataTask:"), dataTask)
}

// Clears the given cache of any cached responses since the provided date.
//
// date: The earliest date of responses that should remain in the cache. Any
// responses with dates later than this parameter should be removed.
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/removeCachedResponses(since:)
func (u URLCache) RemoveCachedResponsesSinceDate(date INSDate) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeCachedResponsesSinceDate:"), date)
}

// Clears the receiver’s cache, removing all stored cached URL responses.
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/removeAllCachedResponses()
func (u URLCache) RemoveAllCachedResponses() {
	objc.Send[objc.ID](u.ID, objc.Sel("removeAllCachedResponses"))
}

// Creates a URL cache object with the specified memory and disk capacities,
// in the specified directory.
//
// memoryCapacity: The memory capacity of the cache, in bytes.
//
// diskCapacity: The disk capacity of the cache, in bytes.
//
// directoryURL: The path to an on-disk directory at which to store the on-disk cache. If
// `directory` is `nil`, the cache uses a default directory.
//
// # Discussion
// 
// A disk cache measured in the tens of megabytes should be acceptable in most
// cases.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLCache/initWithMemoryCapacity:diskCapacity:directoryURL:
func (u URLCache) InitWithMemoryCapacityDiskCapacityDirectoryURL(memoryCapacity uint, diskCapacity uint, directoryURL INSURL) URLCache {
	rv := objc.Send[URLCache](u.ID, objc.Sel("initWithMemoryCapacity:diskCapacity:directoryURL:"), memoryCapacity, diskCapacity, directoryURL)
	return rv
}

// The current size of the on-disk cache, in bytes.
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/currentDiskUsage
func (u URLCache) CurrentDiskUsage() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("currentDiskUsage"))
	return rv
}

// The capacity of the on-disk cache, in bytes.
//
// # Discussion
// 
// When set, the on-disk cache will truncate its contents to the given size,
// if necessary.
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/diskCapacity
func (u URLCache) DiskCapacity() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("diskCapacity"))
	return rv
}
func (u URLCache) SetDiskCapacity(value uint) {
	objc.Send[struct{}](u.ID, objc.Sel("setDiskCapacity:"), value)
}

// The current size of the in-memory cache, in bytes.
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/currentMemoryUsage
func (u URLCache) CurrentMemoryUsage() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("currentMemoryUsage"))
	return rv
}

// The capacity of the in-memory cache, in bytes.
//
// # Discussion
// 
// At the time this property is set, the in-memory cache will truncate its
// contents to the size given, if necessary.
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/memoryCapacity
func (u URLCache) MemoryCapacity() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("memoryCapacity"))
	return rv
}
func (u URLCache) SetMemoryCapacity(value uint) {
	objc.Send[struct{}](u.ID, objc.Sel("setMemoryCapacity:"), value)
}

// The shared URL cache instance.
//
// # Discussion
// 
// If your app doesn’t have special caching requirements or constraints, the
// default shared cache instance should be acceptable. Alternatively, you can
// create a custom [NSURLCache] object and set it as the shared cache instance
// (use `+[NSURLCache setSharedURLCache]` in Objective-C). You should do so
// before making any calls to this method.
//
// See: https://developer.apple.com/documentation/Foundation/URLCache/shared
func (_URLCacheClass URLCacheClass) SharedURLCache() URLCache {
	rv := objc.Send[objc.ID](objc.ID(_URLCacheClass.class), objc.Sel("sharedURLCache"))
	return NSURLCacheFromID(objc.ID(rv))
}
func (_URLCacheClass URLCacheClass) SetSharedURLCache(value URLCache) {
	objc.Send[struct{}](objc.ID(_URLCacheClass.class), objc.Sel("setSharedURLCache:"), value)
}

// GetCachedResponseForDataTask is a synchronous wrapper around [URLCache.GetCachedResponseForDataTaskCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u URLCache) GetCachedResponseForDataTask(ctx context.Context, dataTask INSURLSessionDataTask) (*NSCachedURLResponse, error) {
	done := make(chan *NSCachedURLResponse, 1)
	u.GetCachedResponseForDataTaskCompletionHandler(dataTask, func(val *NSCachedURLResponse) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

