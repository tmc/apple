// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The delegate of an [NSCache](<doc://com.apple.foundation/documentation/Foundation/NSCache>) object implements this protocol to perform specialized actions when an object is about to be evicted or removed from the cache.
//
// See: https://developer.apple.com/documentation/Foundation/NSCacheDelegate
type NSCacheDelegate interface {
	objectivec.IObject
}

// NSCacheDelegateObject wraps an existing Objective-C object that conforms to the NSCacheDelegate protocol.
type NSCacheDelegateObject struct {
	objectivec.Object
}
func (o NSCacheDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSCacheDelegateObjectFromID constructs a [NSCacheDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCacheDelegateObjectFromID(id objc.ID) NSCacheDelegateObject {
	return NSCacheDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Called when an object is about to be evicted or removed from the cache.
//
// cache: The cache with which the object of interest is associated.
//
// obj: The object of interest in the cache.
//
// # Discussion
// 
// It is not possible to modify `cache` from within the implementation of this
// delegate method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCacheDelegate/cache(_:willEvictObject:)
func (o NSCacheDelegateObject) CacheWillEvictObject(cache INSCache, obj objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("cache:willEvictObject:"), cache, obj)
	}

