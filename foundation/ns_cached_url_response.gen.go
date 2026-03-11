// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CachedURLResponse] class.
var (
	_CachedURLResponseClass     CachedURLResponseClass
	_CachedURLResponseClassOnce sync.Once
)

func getCachedURLResponseClass() CachedURLResponseClass {
	_CachedURLResponseClassOnce.Do(func() {
		_CachedURLResponseClass = CachedURLResponseClass{class: objc.GetClass("NSCachedURLResponse")}
	})
	return _CachedURLResponseClass
}

// GetCachedURLResponseClass returns the class object for NSCachedURLResponse.
func GetCachedURLResponseClass() CachedURLResponseClass {
	return getCachedURLResponseClass()
}

type CachedURLResponseClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CachedURLResponseClass) Alloc() CachedURLResponse {
	rv := objc.Send[CachedURLResponse](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}







// A cached response to a URL request.
//
// # Overview
// 
// A [NSCachedURLResponse] object provides the server’s response metadata in
// the form of a [NSURLResponse] object, along with an [NSData] object
// containing the actual cached content data. Its storage policy determines
// whether the response should be cached on disk, in memory, or not at all.
// 
// Cached responses also contain a user info dictionary where you can store
// app-specific information about the cached item.
// 
// The [NSURLCache] class stores and retrieves instances of
// [NSCachedURLResponse].
//
// [NSData]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/OldStylePlists/OldStylePLists.html#//apple_ref/doc/uid/20001012-47169
//
// # Creating a cached URL response
//
//   - [CachedURLResponse.InitWithResponseData]: Creates a cached URL response instance.
//   - [CachedURLResponse.InitWithResponseDataUserInfoStoragePolicy]: Creates a cached URL response with a given server response, data, user-info dictionary, and storage policy.
//
// # Getting cached URL response properties
//
//   - [CachedURLResponse.Data]: The cached response’s data.
//   - [CachedURLResponse.Response]: The URL response object associated with the instance.
//   - [CachedURLResponse.StoragePolicy]: The cached response’s storage policy.
//   - [CachedURLResponse.UserInfo]: The cached response’s user info dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/CachedURLResponse
type CachedURLResponse struct {
	objectivec.Object
}

// CachedURLResponseFromID constructs a [CachedURLResponse] from an objc.ID.
//
// A cached response to a URL request.
func CachedURLResponseFromID(id objc.ID) CachedURLResponse {
	return NSCachedURLResponse{objectivec.Object{id}}
}

// NSCachedURLResponseFromID is an alias for [CachedURLResponseFromID] for cross-framework compatibility.
func NSCachedURLResponseFromID(id objc.ID) CachedURLResponse { return CachedURLResponseFromID(id) }
// NOTE: CachedURLResponse adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [CachedURLResponse] class.
//
// # Creating a cached URL response
//
//   - [ICachedURLResponse.InitWithResponseData]: Creates a cached URL response instance.
//   - [ICachedURLResponse.InitWithResponseDataUserInfoStoragePolicy]: Creates a cached URL response with a given server response, data, user-info dictionary, and storage policy.
//
// # Getting cached URL response properties
//
//   - [ICachedURLResponse.Data]: The cached response’s data.
//   - [ICachedURLResponse.Response]: The URL response object associated with the instance.
//   - [ICachedURLResponse.StoragePolicy]: The cached response’s storage policy.
//   - [ICachedURLResponse.UserInfo]: The cached response’s user info dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/CachedURLResponse
type ICachedURLResponse interface {
	objectivec.IObject
	NSCopying

	// Topic: Creating a cached URL response

	// Creates a cached URL response instance.
	InitWithResponseData(response INSURLResponse, data INSData) CachedURLResponse
	// Creates a cached URL response with a given server response, data, user-info dictionary, and storage policy.
	InitWithResponseDataUserInfoStoragePolicy(response INSURLResponse, data INSData, userInfo INSDictionary, storagePolicy NSURLCacheStoragePolicy) CachedURLResponse

	// Topic: Getting cached URL response properties

	// The cached response’s data.
	Data() INSData
	// The URL response object associated with the instance.
	Response() INSURLResponse
	// The cached response’s storage policy.
	StoragePolicy() NSURLCacheStoragePolicy
	// The cached response’s user info dictionary.
	UserInfo() INSDictionary

	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
	InitWithCoder(coder INSCoder) CachedURLResponse
}





// Init initializes the instance.
func (c CachedURLResponse) Init() CachedURLResponse {
	rv := objc.Send[CachedURLResponse](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CachedURLResponse) Autorelease() CachedURLResponse {
	rv := objc.Send[CachedURLResponse](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCachedURLResponse creates a new CachedURLResponse instance.
func NewCachedURLResponse() CachedURLResponse {
	class := getCachedURLResponseClass()
	rv := objc.Send[CachedURLResponse](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewCachedURLResponseWithCoder(coder INSCoder) CachedURLResponse {
	instance := getCachedURLResponseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CachedURLResponseFromID(rv)
}


// Creates a cached URL response instance.
//
// response: The response to cache.
//
// data: The data to cache.
//
// # Return Value
// 
// A cached URL response object, containing the response and data.
//
// # Discussion
// 
// The cache storage policy is set to the default, [URLCacheStorageAllowed],
// and the user info dictionary is set to `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/CachedURLResponse/init(response:data:)
func NewCachedURLResponseWithResponseData(response INSURLResponse, data INSData) CachedURLResponse {
	instance := getCachedURLResponseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResponse:data:"), response, data)
	return CachedURLResponseFromID(rv)
}


// Creates a cached URL response with a given server response, data, user-info
// dictionary, and storage policy.
//
// response: The response to cache.
//
// data: The data to cache.
//
// userInfo: An optional dictionary of user information. May be `nil`.
//
// storagePolicy: The storage policy for the cached response.
//
// # Return Value
// 
// A cached URL response object, containing the response and data.
//
// See: https://developer.apple.com/documentation/Foundation/CachedURLResponse/init(response:data:userInfo:storagePolicy:)
func NewCachedURLResponseWithResponseDataUserInfoStoragePolicy(response INSURLResponse, data INSData, userInfo INSDictionary, storagePolicy NSURLCacheStoragePolicy) CachedURLResponse {
	instance := getCachedURLResponseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResponse:data:userInfo:storagePolicy:"), response, data, userInfo, storagePolicy)
	return CachedURLResponseFromID(rv)
}







// Creates a cached URL response instance.
//
// response: The response to cache.
//
// data: The data to cache.
//
// # Return Value
// 
// A cached URL response object, containing the response and data.
//
// # Discussion
// 
// The cache storage policy is set to the default, [URLCacheStorageAllowed],
// and the user info dictionary is set to `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/CachedURLResponse/init(response:data:)
func (c CachedURLResponse) InitWithResponseData(response INSURLResponse, data INSData) CachedURLResponse {
	rv := objc.Send[CachedURLResponse](c.ID, objc.Sel("initWithResponse:data:"), response, data)
	return rv
}

// Creates a cached URL response with a given server response, data, user-info
// dictionary, and storage policy.
//
// response: The response to cache.
//
// data: The data to cache.
//
// userInfo: An optional dictionary of user information. May be `nil`.
//
// storagePolicy: The storage policy for the cached response.
//
// # Return Value
// 
// A cached URL response object, containing the response and data.
//
// See: https://developer.apple.com/documentation/Foundation/CachedURLResponse/init(response:data:userInfo:storagePolicy:)
func (c CachedURLResponse) InitWithResponseDataUserInfoStoragePolicy(response INSURLResponse, data INSData, userInfo INSDictionary, storagePolicy NSURLCacheStoragePolicy) CachedURLResponse {
	rv := objc.Send[CachedURLResponse](c.ID, objc.Sel("initWithResponse:data:userInfo:storagePolicy:"), response, data, userInfo, storagePolicy)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (c CachedURLResponse) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (c CachedURLResponse) InitWithCoder(coder INSCoder) CachedURLResponse {
	rv := objc.Send[CachedURLResponse](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}












// The cached response’s data.
//
// See: https://developer.apple.com/documentation/Foundation/CachedURLResponse/data
func (c CachedURLResponse) Data() INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("data"))
	return NSDataFromID(objc.ID(rv))
}



// The URL response object associated with the instance.
//
// See: https://developer.apple.com/documentation/Foundation/CachedURLResponse/response
func (c CachedURLResponse) Response() INSURLResponse {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("response"))
	return NSURLResponseFromID(objc.ID(rv))
}



// The cached response’s storage policy.
//
// See: https://developer.apple.com/documentation/Foundation/CachedURLResponse/storagePolicy
func (c CachedURLResponse) StoragePolicy() NSURLCacheStoragePolicy {
	rv := objc.Send[NSURLCacheStoragePolicy](c.ID, objc.Sel("storagePolicy"))
	return NSURLCacheStoragePolicy(rv)
}



// The cached response’s user info dictionary.
//
// # Discussion
// 
// This value is `nil` if there is no user info dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/CachedURLResponse/userInfo
func (c CachedURLResponse) UserInfo() INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userInfo"))
	return NSDictionaryFromID(objc.ID(rv))
}














			// Protocol methods for NSCopying
			

















