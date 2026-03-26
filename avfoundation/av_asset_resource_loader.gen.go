// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetResourceLoader] class.
var (
	_AVAssetResourceLoaderClass     AVAssetResourceLoaderClass
	_AVAssetResourceLoaderClassOnce sync.Once
)

func getAVAssetResourceLoaderClass() AVAssetResourceLoaderClass {
	_AVAssetResourceLoaderClassOnce.Do(func() {
		_AVAssetResourceLoaderClass = AVAssetResourceLoaderClass{class: objc.GetClass("AVAssetResourceLoader")}
	})
	return _AVAssetResourceLoaderClass
}

// GetAVAssetResourceLoaderClass returns the class object for AVAssetResourceLoader.
func GetAVAssetResourceLoaderClass() AVAssetResourceLoaderClass {
	return getAVAssetResourceLoaderClass()
}

type AVAssetResourceLoaderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetResourceLoaderClass) Alloc() AVAssetResourceLoader {
	rv := objc.Send[AVAssetResourceLoader](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that mediates resource requests from a URL asset.
//
// # Overview
// 
// You do not create resource loader objects yourself. Instead, you retrieve a
// resource loader from the [AVAssetResourceLoader.ResourceLoader] property of an [AVURLAsset]
// object and use it to assign your custom delegate object.
// 
// The delegate you associate with this object must adopt the
// [AVAssetResourceLoaderDelegate] protocol. For more information, see
// [AVAssetResourceLoaderDelegate].
//
// # Accessing the delegate
//
//   - [AVAssetResourceLoader.SetDelegateQueue]: Sets the delegate and dispatch queue to use with the resource loader.
//   - [AVAssetResourceLoader.Delegate]: The delegate object to use when handling resource requests.
//   - [AVAssetResourceLoader.DelegateQueue]: The dispatch queue to use when handling resource requests.
//
// # Loading content keys
//
//   - [AVAssetResourceLoader.PreloadsEligibleContentKeys]: A Boolean value that indicates whether content keys will be loaded as quickly as possible.
//   - [AVAssetResourceLoader.SetPreloadsEligibleContentKeys]
//
// # Supporting Common Media Client Data
//
//   - [AVAssetResourceLoader.SendsCommonMediaClientDataAsHTTPHeaders]: A Boolean value that indicates whether to enable attaching Common Media Client Data as HTTP request headers.
//   - [AVAssetResourceLoader.SetSendsCommonMediaClientDataAsHTTPHeaders]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader
type AVAssetResourceLoader struct {
	objectivec.Object
}

// AVAssetResourceLoaderFromID constructs a [AVAssetResourceLoader] from an objc.ID.
//
// An object that mediates resource requests from a URL asset.
func AVAssetResourceLoaderFromID(id objc.ID) AVAssetResourceLoader {
	return AVAssetResourceLoader{objectivec.Object{ID: id}}
}
// NOTE: AVAssetResourceLoader adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetResourceLoader] class.
//
// # Accessing the delegate
//
//   - [IAVAssetResourceLoader.SetDelegateQueue]: Sets the delegate and dispatch queue to use with the resource loader.
//   - [IAVAssetResourceLoader.Delegate]: The delegate object to use when handling resource requests.
//   - [IAVAssetResourceLoader.DelegateQueue]: The dispatch queue to use when handling resource requests.
//
// # Loading content keys
//
//   - [IAVAssetResourceLoader.PreloadsEligibleContentKeys]: A Boolean value that indicates whether content keys will be loaded as quickly as possible.
//   - [IAVAssetResourceLoader.SetPreloadsEligibleContentKeys]
//
// # Supporting Common Media Client Data
//
//   - [IAVAssetResourceLoader.SendsCommonMediaClientDataAsHTTPHeaders]: A Boolean value that indicates whether to enable attaching Common Media Client Data as HTTP request headers.
//   - [IAVAssetResourceLoader.SetSendsCommonMediaClientDataAsHTTPHeaders]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader
type IAVAssetResourceLoader interface {
	objectivec.IObject

	// Topic: Accessing the delegate

	// Sets the delegate and dispatch queue to use with the resource loader.
	SetDelegateQueue(delegate AVAssetResourceLoaderDelegate, delegateQueue dispatch.Queue)
	// The delegate object to use when handling resource requests.
	Delegate() AVAssetResourceLoaderDelegate
	// The dispatch queue to use when handling resource requests.
	DelegateQueue() dispatch.Queue

	// Topic: Loading content keys

	// A Boolean value that indicates whether content keys will be loaded as quickly as possible.
	PreloadsEligibleContentKeys() bool
	SetPreloadsEligibleContentKeys(value bool)

	// Topic: Supporting Common Media Client Data

	// A Boolean value that indicates whether to enable attaching Common Media Client Data as HTTP request headers.
	SendsCommonMediaClientDataAsHTTPHeaders() bool
	SetSendsCommonMediaClientDataAsHTTPHeaders(value bool)

	// The resource loader for the asset.
	ResourceLoader() IAVAssetResourceLoader
	SetResourceLoader(value IAVAssetResourceLoader)
}

// Init initializes the instance.
func (a AVAssetResourceLoader) Init() AVAssetResourceLoader {
	rv := objc.Send[AVAssetResourceLoader](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetResourceLoader) Autorelease() AVAssetResourceLoader {
	rv := objc.Send[AVAssetResourceLoader](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetResourceLoader creates a new AVAssetResourceLoader instance.
func NewAVAssetResourceLoader() AVAssetResourceLoader {
	class := getAVAssetResourceLoaderClass()
	rv := objc.Send[AVAssetResourceLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets the delegate and dispatch queue to use with the resource loader.
//
// delegate: The delegate object to query when handling resource requests. You may
// specify `nil` if you want to clear the delegate object. The resource loader
// does not store a strong reference to the delegate object.
//
// delegateQueue: The dispatch queue on which to execute resource requests. If the `delegate`
// parameter is not `nil`, this parameter must also not be `nil` and must
// contain a valid dispatch queue. However, if `delegate` is `nil`, this
// parameter may also be `nil`.
// 
// The resource loader maintains a strong reference to the dispatch queue you
// specify.
//
// # Discussion
// 
// You use this method to specify the object to use when handling resource
// requests and the dispatch queue on which to process those requests.
// Resource requests are processed synchronously on the dispatch queue you
// provide.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/setDelegate(_:queue:)
func (a AVAssetResourceLoader) SetDelegateQueue(delegate AVAssetResourceLoaderDelegate, delegateQueue dispatch.Queue) {
	objc.Send[objc.ID](a.ID, objc.Sel("setDelegate:queue:"), delegate, uintptr(delegateQueue.Handle()))
}

// The delegate object to use when handling resource requests.
//
// # Discussion
// 
// The delegate object is responsible for indicating whether or not it is able
// to handle a resource request. And for those requests it does handle, the
// delegate object must initiate the loading of the requested resource.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/delegate
func (a AVAssetResourceLoader) Delegate() AVAssetResourceLoaderDelegate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("delegate"))
	return AVAssetResourceLoaderDelegateObjectFromID(rv)
}
// The dispatch queue to use when handling resource requests.
//
// # Discussion
// 
// Resource requests are processed synchronously on the specified dispatch
// queue.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/delegateQueue
func (a AVAssetResourceLoader) DelegateQueue() dispatch.Queue {
	rv := objc.Send[uintptr](a.ID, objc.Sel("delegateQueue"))
	return dispatch.QueueFromHandle(rv)
}
// A Boolean value that indicates whether content keys will be loaded as
// quickly as possible.
//
// # Discussion
// 
// Set this property to `true` to load eligible keys. This may result in
// network activity. All work done as a result of setting this property to
// `true` is performed asynchronously.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/preloadsEligibleContentKeys
func (a AVAssetResourceLoader) PreloadsEligibleContentKeys() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("preloadsEligibleContentKeys"))
	return rv
}
func (a AVAssetResourceLoader) SetPreloadsEligibleContentKeys(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setPreloadsEligibleContentKeys:"), value)
}
// A Boolean value that indicates whether to enable attaching Common Media
// Client Data as HTTP request headers.
//
// # Discussion
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/sendsCommonMediaClientDataAsHTTPHeaders
func (a AVAssetResourceLoader) SendsCommonMediaClientDataAsHTTPHeaders() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("sendsCommonMediaClientDataAsHTTPHeaders"))
	return rv
}
func (a AVAssetResourceLoader) SetSendsCommonMediaClientDataAsHTTPHeaders(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setSendsCommonMediaClientDataAsHTTPHeaders:"), value)
}
// The resource loader for the asset.
//
// See: https://developer.apple.com/documentation/avfoundation/avurlasset/resourceloader
func (a AVAssetResourceLoader) ResourceLoader() IAVAssetResourceLoader {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("resourceLoader"))
	return AVAssetResourceLoaderFromID(objc.ID(rv))
}
func (a AVAssetResourceLoader) SetResourceLoader(value IAVAssetResourceLoader) {
	objc.Send[struct{}](a.ID, objc.Sel("setResourceLoader:"), value)
}

