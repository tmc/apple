// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLProtocol] class.
var (
	_URLProtocolClass     URLProtocolClass
	_URLProtocolClassOnce sync.Once
)

func getURLProtocolClass() URLProtocolClass {
	_URLProtocolClassOnce.Do(func() {
		_URLProtocolClass = URLProtocolClass{class: objc.GetClass("NSURLProtocol")}
	})
	return _URLProtocolClass
}

// GetURLProtocolClass returns the class object for NSURLProtocol.
func GetURLProtocolClass() URLProtocolClass {
	return getURLProtocolClass()
}

type URLProtocolClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLProtocolClass) Alloc() URLProtocol {
	rv := objc.Send[URLProtocol](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that handles the loading of protocol-specific URL data.
//
// # Overview
// 
// Don’t instantiate a [NSURLProtocol] subclass directly. Instead, create
// subclasses for any custom protocols or URL schemes that your app supports.
// When a download starts, the system creates the appropriate protocol object
// to handle the corresponding URL request. You define your protocol class and
// call the [RegisterClass] class method during your app’s launch time so
// that the system is aware of your protocol.
// 
// To support the customization of protocol-specific requests, create
// extensions to the [URLRequest] class to provide any custom API that you
// need. You can store and retrieve protocol-specific request data by using
// [NSURLProtocol]’s class methods [PropertyForKeyInRequest] and
// [SetPropertyForKeyInRequest].
// 
// Create a [NSURLResponse] for each request your subclass processes
// successfully. You may want to create a custom [NSURLResponse] class to
// provide protocol specific information.
// 
// # Subclassing notes
// 
// When overriding methods of this class, be aware that methods that take a
// `task` parameter are preferred by the system to those that do not.
// Therefore, you should override the task-based methods when subclassing, as
// follows:
// 
// Swift:
// 
// - Initialization — Override [InitWithTaskCachedResponseClient] instead of
// or in addition to [InitWithRequestCachedResponseClient]. Also override the
// task-based [CanInitWithTask] instead of or in addition to the request-based
// [CanInitWithRequest].
// 
// Objective-C:
// 
// - Initialization — Override [CanInitWithTask] and
// [InitWithTaskCachedResponseClient] instead of or in addition to
// [CanInitWithRequest] and [InitWithRequestCachedResponseClient].
//
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
//
// # Creating protocol objects
//
//   - [URLProtocol.InitWithRequestCachedResponseClient]: Creates a URL protocol instance to handle the request.
//   - [URLProtocol.InitWithTaskCachedResponseClient]: Creates a URL protocol instance to handle the task.
//
// # Starting and stopping downloads
//
//   - [URLProtocol.StartLoading]: Starts protocol-specific loading of the request.
//   - [URLProtocol.StopLoading]: Stops protocol-specific loading of the request.
//
// # Getting protocol attributes
//
//   - [URLProtocol.CachedResponse]: The protocol’s cached response.
//   - [URLProtocol.Client]: The object the protocol uses to communicate with the URL loading system.
//   - [URLProtocol.Request]: The protocol’s request.
//   - [URLProtocol.Task]: The protocol’s task.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol
type URLProtocol struct {
	objectivec.Object
}

// URLProtocolFromID constructs a [URLProtocol] from an objc.ID.
//
// An abstract class that handles the loading of protocol-specific URL data.
func URLProtocolFromID(id objc.ID) URLProtocol {
	return NSURLProtocol{objectivec.Object{ID: id}}
}

// NSURLProtocolFromID is an alias for [URLProtocolFromID] for cross-framework compatibility.
func NSURLProtocolFromID(id objc.ID) URLProtocol { return URLProtocolFromID(id) }
// NOTE: URLProtocol adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLProtocol] class.
//
// # Creating protocol objects
//
//   - [IURLProtocol.InitWithRequestCachedResponseClient]: Creates a URL protocol instance to handle the request.
//   - [IURLProtocol.InitWithTaskCachedResponseClient]: Creates a URL protocol instance to handle the task.
//
// # Starting and stopping downloads
//
//   - [IURLProtocol.StartLoading]: Starts protocol-specific loading of the request.
//   - [IURLProtocol.StopLoading]: Stops protocol-specific loading of the request.
//
// # Getting protocol attributes
//
//   - [IURLProtocol.CachedResponse]: The protocol’s cached response.
//   - [IURLProtocol.Client]: The object the protocol uses to communicate with the URL loading system.
//   - [IURLProtocol.Request]: The protocol’s request.
//   - [IURLProtocol.Task]: The protocol’s task.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol
type IURLProtocol interface {
	objectivec.IObject

	// Topic: Creating protocol objects

	// Creates a URL protocol instance to handle the request.
	InitWithRequestCachedResponseClient(request INSURLRequest, cachedResponse INSCachedURLResponse, client NSURLProtocolClient) URLProtocol
	// Creates a URL protocol instance to handle the task.
	InitWithTaskCachedResponseClient(task INSURLSessionTask, cachedResponse INSCachedURLResponse, client NSURLProtocolClient) URLProtocol

	// Topic: Starting and stopping downloads

	// Starts protocol-specific loading of the request.
	StartLoading()
	// Stops protocol-specific loading of the request.
	StopLoading()

	// Topic: Getting protocol attributes

	// The protocol’s cached response.
	CachedResponse() INSCachedURLResponse
	// The object the protocol uses to communicate with the URL loading system.
	Client() NSURLProtocolClient
	// The protocol’s request.
	Request() INSURLRequest
	// The protocol’s task.
	Task() INSURLSessionTask

	// An array of extra protocol subclasses that handle requests in a session.
	ProtocolClasses() objc.Class
	SetProtocolClasses(value objc.Class)
}

// Init initializes the instance.
func (u URLProtocol) Init() URLProtocol {
	rv := objc.Send[URLProtocol](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLProtocol) Autorelease() URLProtocol {
	rv := objc.Send[URLProtocol](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLProtocol creates a new URLProtocol instance.
func NewURLProtocol() URLProtocol {
	class := getURLProtocolClass()
	rv := objc.Send[URLProtocol](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a URL protocol instance to handle the request.
//
// request: The URL request for the URL protocol object. This request is retained.
//
// cachedResponse: A cached response for the request; it may be `nil` if there is no existing
// cached response for the request.
//
// client: An object that provides an implementation of the [NSURLProtocolClient]
// protocol that this instance uses to communicate with the URL Loading
// System. This client object is retained.
//
// # Return Value
// 
// The initialized protocol object.
//
// # Discussion
// 
// Subclasses should override this method to do any custom initialization.
// Don’t call this method explicitly. When you register your custom protocol
// class, the system will initialize instances of your protocol as needed.
// 
// This is the designated initializer for [NSURLProtocol].
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/init(request:cachedResponse:client:)
func NewURLProtocolWithRequestCachedResponseClient(request INSURLRequest, cachedResponse INSCachedURLResponse, client NSURLProtocolClient) URLProtocol {
	instance := getURLProtocolClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRequest:cachedResponse:client:"), request, cachedResponse, client)
	return URLProtocolFromID(rv)
}

// Creates a URL protocol instance to handle the task.
//
// task: A task containing a URL request to be performed by the protocol.
//
// cachedResponse: A cached response for the task; may be `nil` if there is no existing cached
// response for the task.
//
// client: An object that provides an implementation of the [NSURLProtocolClient]
// protocol that this instance uses to communicate with the URL loading
// system. This client object is retained.
//
// # Return Value
// 
// The initialized protocol object.
//
// # Discussion
// 
// Subclasses should override this method to do any custom initialization.
// Don’t call this method explicitly. When you register your custom protocol
// class, the system will initialize instances of your protocol as needed.
// 
// This initializer calls through to [InitWithRequestCachedResponseClient].
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/init(task:cachedResponse:client:)
func NewURLProtocolWithTaskCachedResponseClient(task INSURLSessionTask, cachedResponse INSCachedURLResponse, client NSURLProtocolClient) URLProtocol {
	instance := getURLProtocolClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTask:cachedResponse:client:"), task, cachedResponse, client)
	return URLProtocolFromID(rv)
}

// Creates a URL protocol instance to handle the request.
//
// request: The URL request for the URL protocol object. This request is retained.
//
// cachedResponse: A cached response for the request; it may be `nil` if there is no existing
// cached response for the request.
//
// client: An object that provides an implementation of the [NSURLProtocolClient]
// protocol that this instance uses to communicate with the URL Loading
// System. This client object is retained.
//
// # Return Value
// 
// The initialized protocol object.
//
// # Discussion
// 
// Subclasses should override this method to do any custom initialization.
// Don’t call this method explicitly. When you register your custom protocol
// class, the system will initialize instances of your protocol as needed.
// 
// This is the designated initializer for [NSURLProtocol].
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/init(request:cachedResponse:client:)
func (u URLProtocol) InitWithRequestCachedResponseClient(request INSURLRequest, cachedResponse INSCachedURLResponse, client NSURLProtocolClient) URLProtocol {
	rv := objc.Send[URLProtocol](u.ID, objc.Sel("initWithRequest:cachedResponse:client:"), request, cachedResponse, client)
	return rv
}
// Creates a URL protocol instance to handle the task.
//
// task: A task containing a URL request to be performed by the protocol.
//
// cachedResponse: A cached response for the task; may be `nil` if there is no existing cached
// response for the task.
//
// client: An object that provides an implementation of the [NSURLProtocolClient]
// protocol that this instance uses to communicate with the URL loading
// system. This client object is retained.
//
// # Return Value
// 
// The initialized protocol object.
//
// # Discussion
// 
// Subclasses should override this method to do any custom initialization.
// Don’t call this method explicitly. When you register your custom protocol
// class, the system will initialize instances of your protocol as needed.
// 
// This initializer calls through to [InitWithRequestCachedResponseClient].
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/init(task:cachedResponse:client:)
func (u URLProtocol) InitWithTaskCachedResponseClient(task INSURLSessionTask, cachedResponse INSCachedURLResponse, client NSURLProtocolClient) URLProtocol {
	rv := objc.Send[URLProtocol](u.ID, objc.Sel("initWithTask:cachedResponse:client:"), task, cachedResponse, client)
	return rv
}
// Starts protocol-specific loading of the request.
//
// # Discussion
// 
// When this method is called, the subclass implementation should start
// loading the request, providing feedback to the URL loading system via the
// [NSURLProtocolClient] protocol.
// 
// Subclasses must implement this method.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/startLoading()
func (u URLProtocol) StartLoading() {
	objc.Send[objc.ID](u.ID, objc.Sel("startLoading"))
}
// Stops protocol-specific loading of the request.
//
// # Discussion
// 
// When this method is called, the subclass implementation should stop loading
// a request. This could be in response to a cancel operation, so protocol
// implementations must be able to handle this call while a load is in
// progress. When your protocol receives a call to this method, it should also
// stop sending notifications to the client.
// 
// Subclasses must implement this method.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/stopLoading()
func (u URLProtocol) StopLoading() {
	objc.Send[objc.ID](u.ID, objc.Sel("stopLoading"))
}

// Attempts to register a subclass of [NSURLProtocol], making it visible to
// the URL loading system.
//
// protocolClass: The subclass to register.
//
// # Return Value
// 
// [true] if the registration is successful, [false] otherwise. The only
// failure condition is if `protocolClass` is not a subclass of
// [NSURLProtocol].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Register any custom [NSURLProtocol] subclasses prior to making URL
// requests. When the URL loading system begins to load a request, it tries to
// initialize each registered protocol class with the specified request. The
// first [NSURLProtocol] subclass to return [true] when sent a
// [CanInitWithRequest] message is used to load the request. There is no
// guarantee that all registered protocol classes will be consulted.
// 
// Classes are consulted in the reverse order of their registration. A similar
// design governs the process to create the canonical form of a request with
// [CanonicalRequestForRequest].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/registerClass(_:)
func (_URLProtocolClass URLProtocolClass) RegisterClass(protocolClass objc.Class) bool {
	rv := objc.Send[bool](objc.ID(_URLProtocolClass.class), objc.Sel("registerClass:"), protocolClass)
	return rv
}
// Unregisters the specified subclass of [NSURLProtocol].
//
// protocolClass: The subclass of [NSURLProtocol] to unregister.
//
// # Discussion
// 
// After this method is invoked, `protocolClass` is no longer consulted by the
// URL loading system.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/unregisterClass(_:)
func (_URLProtocolClass URLProtocolClass) UnregisterClass(protocolClass objc.Class) {
	objc.Send[objc.ID](objc.ID(_URLProtocolClass.class), objc.Sel("unregisterClass:"), protocolClass)
}
// Determines whether the protocol subclass can handle the specified request.
//
// request: The request to be handled.
//
// # Return Value
// 
// [true] if the protocol subclass can handle `request`, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// A subclass should inspect `request` and determine whether or not the
// implementation can perform a load with that request.
// 
// This is an abstract method and subclasses must provide an implementation.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/canInit(with:)-76brg
func (_URLProtocolClass URLProtocolClass) CanInitWithRequest(request INSURLRequest) bool {
	rv := objc.Send[bool](objc.ID(_URLProtocolClass.class), objc.Sel("canInitWithRequest:"), request)
	return rv
}
// Determines whether the protocol subclass can handle the specified task.
//
// task: A URL session task containing the request to be handled.
//
// # Discussion
// 
// A subclass should inspect the task’s request and determine whether or not
// the implementation can perform a load with that task.
// 
// This is an abstract method and subclasses must provide an implementation.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/canInit(with:)-18gbo
func (_URLProtocolClass URLProtocolClass) CanInitWithTask(task INSURLSessionTask) bool {
	rv := objc.Send[bool](objc.ID(_URLProtocolClass.class), objc.Sel("canInitWithTask:"), task)
	return rv
}
// Fetches the property associated with the specified key in the specified
// request.
//
// key: The key of the desired property.
//
// request: The request whose properties are to be queried.
//
// # Return Value
// 
// The property associated with `key`, or `nil` if no property has been stored
// for `key`.
//
// # Discussion
// 
// Use this method to access protocol-specific information associated with
// [URLRequest] objects.
//
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/property(forKey:in:)
func (_URLProtocolClass URLProtocolClass) PropertyForKeyInRequest(key string, request INSURLRequest) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_URLProtocolClass.class), objc.Sel("propertyForKey:inRequest:"), objc.String(key), request)
	return objectivec.Object{ID: rv}
}
// Sets the property associated with the specified key in the specified
// request.
//
// value: The value to set for the specified property.
//
// key: The key for the specified property.
//
// request: The request for which to create the property.
//
// # Discussion
// 
// Use this method to provide an interface for protocol implementors to
// customize protocol-specific information associated with [URLRequest]
// objects.
//
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/setProperty(_:forKey:in:)
func (_URLProtocolClass URLProtocolClass) SetPropertyForKeyInRequest(value objectivec.IObject, key string, request INSMutableURLRequest) {
	objc.Send[objc.ID](objc.ID(_URLProtocolClass.class), objc.Sel("setProperty:forKey:inRequest:"), value, objc.String(key), request)
}
// Removes the property associated with the specified key in the specified
// request.
//
// key: The key whose value should be removed.
//
// request: The request from which to remove the property value.
//
// # Discussion
// 
// This method is used to provide an interface for protocol implementors to
// customize protocol-specific information associated with [URLRequest]
// objects, or [NSMutableURLRequest] objects in Objective-C.
//
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/removeProperty(forKey:in:)
func (_URLProtocolClass URLProtocolClass) RemovePropertyForKeyInRequest(key string, request INSMutableURLRequest) {
	objc.Send[objc.ID](objc.ID(_URLProtocolClass.class), objc.Sel("removePropertyForKey:inRequest:"), objc.String(key), request)
}
// Returns a canonical version of the specified request.
//
// request: The request whose canonical version is desired.
//
// # Return Value
// 
// The canonical form of `request`.
//
// # Discussion
// 
// It is up to each concrete protocol implementation to define what
// “canonical” means. A protocol should guarantee that the same input
// request always yields the same canonical form.
// 
// Special consideration should be given when implementing this method,
// because the canonical form of a request is used to lookup objects in the
// URL cache, a process which performs equality checks between [URLRequest]
// instances.
// 
// This is an abstract method and subclasses must provide an implementation.
//
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/canonicalRequest(for:)
func (_URLProtocolClass URLProtocolClass) CanonicalRequestForRequest(request INSURLRequest) NSURLRequest {
	rv := objc.Send[objc.ID](objc.ID(_URLProtocolClass.class), objc.Sel("canonicalRequestForRequest:"), request)
	return NSURLRequestFromID(rv)
}
// A Boolean value indicating whether two requests are equivalent for cache
// purposes.
//
// a: The request to compare with `bRequest`.
//
// b: The request to compare with `aRequest`.
//
// # Return Value
// 
// [true] if `aRequest` and `bRequest` are equivalent for cache purposes,
// [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Requests are considered equivalent for cache purposes if and only if they
// would be handled by the same protocol and that protocol declares them
// equivalent after performing implementation-specific checks.
// 
// The [NSURLProtocol] implementation of this method compares the URLs of the
// requests to determine if the requests should be considered equivalent.
// Subclasses can override this method to provide protocol-specific
// comparisons.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/requestIsCacheEquivalent(_:to:)
func (_URLProtocolClass URLProtocolClass) RequestIsCacheEquivalentToRequest(a INSURLRequest, b INSURLRequest) bool {
	rv := objc.Send[bool](objc.ID(_URLProtocolClass.class), objc.Sel("requestIsCacheEquivalent:toRequest:"), a, b)
	return rv
}

// The protocol’s cached response.
//
// # Discussion
// 
// If not overridden in a subclass, this method returns the cached response
// stored at initialization time.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/cachedResponse
func (u URLProtocol) CachedResponse() INSCachedURLResponse {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("cachedResponse"))
	return NSCachedURLResponseFromID(objc.ID(rv))
}
// The object the protocol uses to communicate with the URL loading system.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/client
func (u URLProtocol) Client() NSURLProtocolClient {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("client"))
	return NSURLProtocolClientObjectFromID(rv)
}
// The protocol’s request.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/request
func (u URLProtocol) Request() INSURLRequest {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("request"))
	return NSURLRequestFromID(objc.ID(rv))
}
// The protocol’s task.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocol/task
func (u URLProtocol) Task() INSURLSessionTask {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("task"))
	return NSURLSessionTaskFromID(objc.ID(rv))
}
// An array of extra protocol subclasses that handle requests in a session.
//
// See: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/protocolclasses
func (u URLProtocol) ProtocolClasses() objc.Class {
	rv := objc.Send[objc.Class](u.ID, objc.Sel("protocolClasses"))
	return rv
}
func (u URLProtocol) SetProtocolClasses(value objc.Class) {
	objc.Send[struct{}](u.ID, objc.Sel("setProtocolClasses:"), value)
}

