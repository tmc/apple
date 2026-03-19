// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The interface used by [URLProtocol](<doc://com.apple.foundation/documentation/Foundation/URLProtocol>) subclasses to communicate with the URL Loading System.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient
type NSURLProtocolClient interface {
	objectivec.IObject

	// Tells the client that the protocol implementation has created a response object for the request.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:didReceive:cacheStoragePolicy:)
	URLProtocolDidReceiveResponseCacheStoragePolicy(protocol_ INSURLProtocol, response INSURLResponse, policy NSURLCacheStoragePolicy)

	// Tells the client that the protocol implementation has been redirected.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:wasRedirectedTo:redirectResponse:)
	URLProtocolWasRedirectedToRequestRedirectResponse(protocol_ INSURLProtocol, request INSURLRequest, redirectResponse INSURLResponse)

	// Tells the client that a cached response is valid.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:cachedResponseIsValid:)
	URLProtocolCachedResponseIsValid(protocol_ INSURLProtocol, cachedResponse INSCachedURLResponse)

	// Tells the client that an authentication challenge has been canceled.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:didCancel:)
	URLProtocolDidCancelAuthenticationChallenge(protocol_ INSURLProtocol, challenge INSURLAuthenticationChallenge)

	// Tells the client that the URL Loading System received an authentication challenge.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:didReceive:)
	URLProtocolDidReceiveAuthenticationChallenge(protocol_ INSURLProtocol, challenge INSURLAuthenticationChallenge)

	// Tells the client that the load request failed due to an error.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:didFailWithError:)
	URLProtocolDidFailWithError(protocol_ INSURLProtocol, error_ INSError)

	// Tells the client that the protocol implementation has loaded some data.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:didLoad:)
	URLProtocolDidLoadData(protocol_ INSURLProtocol, data INSData)

	// Tells the client that the protocol implementation has finished loading.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocolDidFinishLoading(_:)
	URLProtocolDidFinishLoading(protocol_ INSURLProtocol)
}

// NSURLProtocolClientObject wraps an existing Objective-C object that conforms to the NSURLProtocolClient protocol.
type NSURLProtocolClientObject struct {
	objectivec.Object
}
func (o NSURLProtocolClientObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSURLProtocolClientObjectFromID constructs a [NSURLProtocolClientObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSURLProtocolClientObjectFromID(id objc.ID) NSURLProtocolClientObject {
	return NSURLProtocolClientObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the client that the protocol implementation has created a response
// object for the request.
//
// protocol: The URL protocol object sending the message.
//
// response: The newly available response object.
//
// policy: The cache storage policy for the response.
//
// # Discussion
// 
// The implementation should use the provided cache storage policy to
// determine whether to store the response in a cache.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:didReceive:cacheStoragePolicy:)
func (o NSURLProtocolClientObject) URLProtocolDidReceiveResponseCacheStoragePolicy(protocol_ INSURLProtocol, response INSURLResponse, policy NSURLCacheStoragePolicy) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLProtocol:didReceiveResponse:cacheStoragePolicy:"), protocol_, response, policy)
	}
// Tells the client that the protocol implementation has been redirected.
//
// protocol: The URL protocol object sending the message.
//
// request: The new request that the original request was redirected to.
//
// redirectResponse: The response from the original request that caused the redirect.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:wasRedirectedTo:redirectResponse:)
func (o NSURLProtocolClientObject) URLProtocolWasRedirectedToRequestRedirectResponse(protocol_ INSURLProtocol, request INSURLRequest, redirectResponse INSURLResponse) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLProtocol:wasRedirectedToRequest:redirectResponse:"), protocol_, request, redirectResponse)
	}
// Tells the client that a cached response is valid.
//
// protocol: The URL protocol object sending the message.
//
// cachedResponse: The cached response whose validity is being communicated.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:cachedResponseIsValid:)
func (o NSURLProtocolClientObject) URLProtocolCachedResponseIsValid(protocol_ INSURLProtocol, cachedResponse INSCachedURLResponse) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLProtocol:cachedResponseIsValid:"), protocol_, cachedResponse)
	}
// Tells the client that an authentication challenge has been canceled.
//
// protocol: The URL protocol object sending the message.
//
// challenge: The authentication challenge that was canceled.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:didCancel:)
func (o NSURLProtocolClientObject) URLProtocolDidCancelAuthenticationChallenge(protocol_ INSURLProtocol, challenge INSURLAuthenticationChallenge) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLProtocol:didCancelAuthenticationChallenge:"), protocol_, challenge)
	}
// Tells the client that the URL Loading System received an authentication
// challenge.
//
// protocol: The URL protocol object sending the message.
//
// challenge: The authentication challenge that has been received.
//
// # Discussion
// 
// The protocol client guarantees that it will answer the request on the same
// thread that called this method. The client may add a default credential to
// the challenge it issues to the connection delegate, if `protocol` did not
// provide one.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:didReceive:)
func (o NSURLProtocolClientObject) URLProtocolDidReceiveAuthenticationChallenge(protocol_ INSURLProtocol, challenge INSURLAuthenticationChallenge) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLProtocol:didReceiveAuthenticationChallenge:"), protocol_, challenge)
	}
// Tells the client that the load request failed due to an error.
//
// protocol: The URL protocol object sending the message.
//
// error: The error that caused the failure of the load request.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:didFailWithError:)
func (o NSURLProtocolClientObject) URLProtocolDidFailWithError(protocol_ INSURLProtocol, error_ INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLProtocol:didFailWithError:"), protocol_, error_)
	}
// Tells the client that the protocol implementation has loaded some data.
//
// protocol: The URL protocol object sending the message.
//
// data: The data being made available.
//
// # Discussion
// 
// The data object must contain only new data loaded since the previous
// invocation of this method.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocol(_:didLoad:)
func (o NSURLProtocolClientObject) URLProtocolDidLoadData(protocol_ INSURLProtocol, data INSData) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLProtocol:didLoadData:"), protocol_, data)
	}
// Tells the client that the protocol implementation has finished loading.
//
// protocol: The URL protocol object sending the message.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtocolClient/urlProtocolDidFinishLoading(_:)
func (o NSURLProtocolClientObject) URLProtocolDidFinishLoading(protocol_ INSURLProtocol) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLProtocolDidFinishLoading:"), protocol_)
	}

