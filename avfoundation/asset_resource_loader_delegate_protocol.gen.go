// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods you can implement to handle resource-loading requests coming from a URL asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoaderDelegate
type AVAssetResourceLoaderDelegate interface {
	objectivec.IObject
}

// AVAssetResourceLoaderDelegateObject wraps an existing Objective-C object that conforms to the AVAssetResourceLoaderDelegate protocol.
type AVAssetResourceLoaderDelegateObject struct {
	objectivec.Object
}

func (o AVAssetResourceLoaderDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAssetResourceLoaderDelegateObjectFromID constructs a [AVAssetResourceLoaderDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAssetResourceLoaderDelegateObjectFromID(id objc.ID) AVAssetResourceLoaderDelegateObject {
	return AVAssetResourceLoaderDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate if it wants to load the requested resource.
//
// resourceLoader: The resource loader object that is making the request.
//
// loadingRequest: The loading request object that contains information about the requested
// resource.
//
// # Return Value
//
// true if your delegate can load the resource specified by the
// `loadingRequest` parameter or false if it cannot.
//
// # Discussion
//
// The resource loader object calls this method when assistance is required of
// your code to load the specified resource. For example, the resource loader
// might call this method to load decryption keys that have been specified
// using a custom URL scheme.
//
// Returning true from this method, implies only that the receiver will load,
// or at least attempt to load, the resource. In some implementations, the
// actual work of loading the resource might be initiated on another thread,
// running asynchronously to the resource loading delegate; whether the work
// begins immediately or merely soon is an implementation detail of the client
// application.
//
// You can load the resource synchronously or asynchronously. In both cases,
// you must indicate success or failure of the operation by calling the
// [finishLoading(with:data:redirect:)] or [FinishLoadingWithError] method of
// the request object when you finish. If you load the resource
// asynchronously, you must also store a strong reference to the object in the
// `loadingRequest` parameter before returning from this method.
//
// If you return false from this method, the resource loader treats the
// loading of the resource as having failed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoaderDelegate/resourceLoader(_:shouldWaitForLoadingOfRequestedResource:)
//
// [finishLoading(with:data:redirect:)]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequest/finishLoading(with:data:redirect:)
func (o AVAssetResourceLoaderDelegateObject) ResourceLoaderShouldWaitForLoadingOfRequestedResource(resourceLoader IAVAssetResourceLoader, loadingRequest IAVAssetResourceLoadingRequest) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("resourceLoader:shouldWaitForLoadingOfRequestedResource:"), resourceLoader, loadingRequest)
	return rv
}

// Tells the delegate when assistance is required of the application to renew
// a resource.
//
// resourceLoader: The resource loader.
//
// renewalRequest: An instance of [AVAssetResourceRenewalRequest] that provides information
// about the requested resource.
//
// # Return Value
//
// true if the delegate can renew the resource; otherwise false.
//
// # Discussion
//
// Delegates receive this message when assistance is required to renew a
// resource previously loaded by
// [ResourceLoaderShouldWaitForLoadingOfRequestedResource]. For example, this
// method is invoked to for decryption keys that require renewal, as indicated
// in a response to a prior invocation of
// [ResourceLoaderShouldWaitForLoadingOfRequestedResource].
//
// If the result is true, the resource loader expects invocation, either
// subsequently or immediately, of either the [AVAssetResourceRenewalRequest]
// method `finishLoading` or “. If you intend to finish loading the resource
// after your handling of this message returns, you must retain the
// `renewalRequest` until after loading is finished.
//
// If the result is false, the resource loader treats the loading of the
// resource as having failed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoaderDelegate/resourceLoader(_:shouldWaitForRenewalOfRequestedResource:)
func (o AVAssetResourceLoaderDelegateObject) ResourceLoaderShouldWaitForRenewalOfRequestedResource(resourceLoader IAVAssetResourceLoader, renewalRequest IAVAssetResourceRenewalRequest) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("resourceLoader:shouldWaitForRenewalOfRequestedResource:"), resourceLoader, renewalRequest)
	return rv
}

// Informs the delegate that a prior loading request has been cancelled.
//
// resourceLoader: The resource loader.
//
// loadingRequest: The loading request that has been cancelled.
//
// # Discussion
//
// Previously issued loading requests can be cancelled when data from the
// resource is no longer required or when a loading request is superseded by
// new requests for data from the same resource.
//
// For example, if to complete a seek operation it becomes necessary to load a
// range of bytes that’s different from a range previously requested, the
// prior request may be cancelled while the delegate is still handling it.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoaderDelegate/resourceLoader(_:didCancel:)-3nl51
func (o AVAssetResourceLoaderDelegateObject) ResourceLoaderDidCancelLoadingRequest(resourceLoader IAVAssetResourceLoader, loadingRequest IAVAssetResourceLoadingRequest) {
	objc.Send[struct{}](o.ID, objc.Sel("resourceLoader:didCancelLoadingRequest:"), resourceLoader, loadingRequest)
}

// Tells the delegate that assistance is required of the application to
// respond to an authentication challenge.
//
// resourceLoader: The resource loader.
//
// authenticationChallenge: The authentication challenge.
//
// # Return Value
//
// true if the resource loader should wait for a response to the
// authentication challenge; otherwise false.
//
// # Discussion
//
// Delegates receive this message when assistance is required of the
// application to respond to an authentication challenge.
//
// Return true if you expect a response either subsequently or immediately to
// the authenticationChallenger object’s sender.
//
// If you intend to respond to the authentication challenge after your
// handling of “ returns, you must retain the authenticationChallenge until
// after your response has been made.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoaderDelegate/resourceLoader(_:shouldWaitForResponseTo:)
func (o AVAssetResourceLoaderDelegateObject) ResourceLoaderShouldWaitForResponseToAuthenticationChallenge(resourceLoader IAVAssetResourceLoader, authenticationChallenge foundation.NSURLAuthenticationChallenge) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("resourceLoader:shouldWaitForResponseToAuthenticationChallenge:"), resourceLoader, authenticationChallenge)
	return rv
}

// Informs the delegate that a prior authentication challenge has been
// cancelled.
//
// resourceLoader: The resource loader.
//
// authenticationChallenge: The authentication challenge that has been cancelled.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoaderDelegate/resourceLoader(_:didCancel:)-1wqin
func (o AVAssetResourceLoaderDelegateObject) ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader IAVAssetResourceLoader, authenticationChallenge foundation.NSURLAuthenticationChallenge) {
	objc.Send[struct{}](o.ID, objc.Sel("resourceLoader:didCancelAuthenticationChallenge:"), resourceLoader, authenticationChallenge)
}

// AVAssetResourceLoaderDelegateConfig holds optional typed callbacks for [AVAssetResourceLoaderDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate
type AVAssetResourceLoaderDelegateConfig struct {

	// Processing resource requests
	// ResourceLoaderShouldWaitForLoadingOfRequestedResource — Asks the delegate if it wants to load the requested resource.
	ResourceLoaderShouldWaitForLoadingOfRequestedResource func(resourceLoader AVAssetResourceLoader, loadingRequest AVAssetResourceLoadingRequest) bool
	// ResourceLoaderShouldWaitForRenewalOfRequestedResource — Tells the delegate when assistance is required of the application to renew a resource.
	ResourceLoaderShouldWaitForRenewalOfRequestedResource func(resourceLoader AVAssetResourceLoader, renewalRequest AVAssetResourceRenewalRequest) bool

	// Other Methods
	// ResourceLoaderDidCancelLoadingRequest — Informs the delegate that a prior loading request has been cancelled.
	ResourceLoaderDidCancelLoadingRequest func(resourceLoader AVAssetResourceLoader, loadingRequest AVAssetResourceLoadingRequest)
	// ResourceLoaderShouldWaitForResponseToAuthenticationChallenge — Tells the delegate that assistance is required of the application to respond to an authentication challenge.
	ResourceLoaderShouldWaitForResponseToAuthenticationChallenge func(resourceLoader AVAssetResourceLoader, authenticationChallenge foundation.NSURLAuthenticationChallenge) bool
	// ResourceLoaderDidCancelAuthenticationChallenge — Informs the delegate that a prior authentication challenge has been cancelled.
	ResourceLoaderDidCancelAuthenticationChallenge func(resourceLoader AVAssetResourceLoader, authenticationChallenge foundation.NSURLAuthenticationChallenge)
}

// NewAVAssetResourceLoaderDelegate creates an Objective-C object implementing the [AVAssetResourceLoaderDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVAssetResourceLoaderDelegateObject] satisfies the [AVAssetResourceLoaderDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate
func NewAVAssetResourceLoaderDelegate(config AVAssetResourceLoaderDelegateConfig) AVAssetResourceLoaderDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVAssetResourceLoaderDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ResourceLoaderShouldWaitForLoadingOfRequestedResource != nil {
		fn := config.ResourceLoaderShouldWaitForLoadingOfRequestedResource
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("resourceLoader:shouldWaitForLoadingOfRequestedResource:"),
			Fn: func(self objc.ID, _cmd objc.SEL, resourceLoaderID objc.ID, loadingRequestID objc.ID) bool {
				resourceLoader := AVAssetResourceLoaderFromID(resourceLoaderID)
				loadingRequest := AVAssetResourceLoadingRequestFromID(loadingRequestID)
				return fn(resourceLoader, loadingRequest)
			},
		})
	}

	if config.ResourceLoaderShouldWaitForRenewalOfRequestedResource != nil {
		fn := config.ResourceLoaderShouldWaitForRenewalOfRequestedResource
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("resourceLoader:shouldWaitForRenewalOfRequestedResource:"),
			Fn: func(self objc.ID, _cmd objc.SEL, resourceLoaderID objc.ID, renewalRequestID objc.ID) bool {
				resourceLoader := AVAssetResourceLoaderFromID(resourceLoaderID)
				renewalRequest := AVAssetResourceRenewalRequestFromID(renewalRequestID)
				return fn(resourceLoader, renewalRequest)
			},
		})
	}

	if config.ResourceLoaderDidCancelLoadingRequest != nil {
		fn := config.ResourceLoaderDidCancelLoadingRequest
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("resourceLoader:didCancelLoadingRequest:"),
			Fn: func(self objc.ID, _cmd objc.SEL, resourceLoaderID objc.ID, loadingRequestID objc.ID) {
				resourceLoader := AVAssetResourceLoaderFromID(resourceLoaderID)
				loadingRequest := AVAssetResourceLoadingRequestFromID(loadingRequestID)
				fn(resourceLoader, loadingRequest)
			},
		})
	}

	if config.ResourceLoaderShouldWaitForResponseToAuthenticationChallenge != nil {
		fn := config.ResourceLoaderShouldWaitForResponseToAuthenticationChallenge
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("resourceLoader:shouldWaitForResponseToAuthenticationChallenge:"),
			Fn: func(self objc.ID, _cmd objc.SEL, resourceLoaderID objc.ID, authenticationChallengeID objc.ID) bool {
				resourceLoader := AVAssetResourceLoaderFromID(resourceLoaderID)
				authenticationChallenge := foundation.NSURLAuthenticationChallengeFromID(authenticationChallengeID)
				return fn(resourceLoader, authenticationChallenge)
			},
		})
	}

	if config.ResourceLoaderDidCancelAuthenticationChallenge != nil {
		fn := config.ResourceLoaderDidCancelAuthenticationChallenge
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("resourceLoader:didCancelAuthenticationChallenge:"),
			Fn: func(self objc.ID, _cmd objc.SEL, resourceLoaderID objc.ID, authenticationChallengeID objc.ID) {
				resourceLoader := AVAssetResourceLoaderFromID(resourceLoaderID)
				authenticationChallenge := foundation.NSURLAuthenticationChallengeFromID(authenticationChallengeID)
				fn(resourceLoader, authenticationChallenge)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVAssetResourceLoaderDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVAssetResourceLoaderDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVAssetResourceLoaderDelegateObjectFromID(instance)
}
