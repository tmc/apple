// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCache] class.
var (
	_NSCacheClass     NSCacheClass
	_NSCacheClassOnce sync.Once
)

func getNSCacheClass() NSCacheClass {
	_NSCacheClassOnce.Do(func() {
		_NSCacheClass = NSCacheClass{class: objc.GetClass("NSCache")}
	})
	return _NSCacheClass
}

// GetNSCacheClass returns the class object for NSCache.
func GetNSCacheClass() NSCacheClass {
	return getNSCacheClass()
}

type NSCacheClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCacheClass) Alloc() NSCache {
	rv := objc.Send[NSCache](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A mutable collection you use to temporarily store transient key-value pairs
// that are subject to eviction when resources are low.
//
// # Overview
// 
// Cache objects differ from other mutable collections in a few ways:
// 
// - The [NSCache] class incorporates various auto-eviction policies, which
// ensure that a cache doesn’t use too much of the system’s memory. If
// memory is needed by other applications, these policies remove some items
// from the cache, minimizing its memory footprint. - You can add, remove, and
// query items in the cache from different threads without having to lock the
// cache yourself. - Unlike an [NSMutableDictionary] object, a cache does not
// copy the key objects that are put into it.
// 
// You typically use [NSCache] objects to temporarily store objects with
// transient data that are expensive to create. Reusing these objects can
// provide performance benefits, because their values do not have to be
// recalculated. However, the objects are not critical to the application and
// can be discarded if memory is tight. If discarded, their values will have
// to be recomputed again when needed.
// 
// Objects that have subcomponents that can be discarded when not being used
// can adopt the [NSDiscardableContent] protocol to improve cache eviction
// behavior. By default, [NSDiscardableContent] objects in a cache are
// automatically removed if their content is discarded, although this
// automatic removal policy can be changed. If an [NSDiscardableContent]
// object is put into the cache, the cache calls [DiscardContentIfPossible] on
// it upon its removal.
//
// # Managing the Name
//
//   - [NSCache.Name]: The name of the cache.
//   - [NSCache.SetName]
//
// # Managing Cache Size
//
//   - [NSCache.CountLimit]: The maximum number of objects the cache should hold.
//   - [NSCache.SetCountLimit]
//   - [NSCache.TotalCostLimit]: The maximum total cost that the cache can hold before it starts evicting objects.
//   - [NSCache.SetTotalCostLimit]
//
// # Managing Discardable Content
//
//   - [NSCache.EvictsObjectsWithDiscardedContent]: Whether the cache will automatically evict discardable-content objects whose content has been discarded.
//   - [NSCache.SetEvictsObjectsWithDiscardedContent]
//
// # Managing the Delegate
//
//   - [NSCache.Delegate]: The cache’s delegate.
//   - [NSCache.SetDelegate]
//
// # Getting a Cached Value
//
//   - [NSCache.ObjectForKey]: Returns the value associated with a given key.
//
// # Adding and Removing Cached Values
//
//   - [NSCache.SetObjectForKey]: Sets the value of the specified key in the cache.
//   - [NSCache.SetObjectForKeyCost]: Sets the value of the specified key in the cache, and associates the key-value pair with the specified cost.
//   - [NSCache.RemoveObjectForKey]: Removes the value of the specified key in the cache.
//   - [NSCache.RemoveAllObjects]: Empties the cache.
//
// See: https://developer.apple.com/documentation/Foundation/NSCache
type NSCache struct {
	objectivec.Object
}

// NSCacheFromID constructs a [NSCache] from an objc.ID.
//
// A mutable collection you use to temporarily store transient key-value pairs
// that are subject to eviction when resources are low.
func NSCacheFromID(id objc.ID) NSCache {
	return NSCache{objectivec.Object{ID: id}}
}
// NOTE: NSCache adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCache] class.
//
// # Managing the Name
//
//   - [INSCache.Name]: The name of the cache.
//   - [INSCache.SetName]
//
// # Managing Cache Size
//
//   - [INSCache.CountLimit]: The maximum number of objects the cache should hold.
//   - [INSCache.SetCountLimit]
//   - [INSCache.TotalCostLimit]: The maximum total cost that the cache can hold before it starts evicting objects.
//   - [INSCache.SetTotalCostLimit]
//
// # Managing Discardable Content
//
//   - [INSCache.EvictsObjectsWithDiscardedContent]: Whether the cache will automatically evict discardable-content objects whose content has been discarded.
//   - [INSCache.SetEvictsObjectsWithDiscardedContent]
//
// # Managing the Delegate
//
//   - [INSCache.Delegate]: The cache’s delegate.
//   - [INSCache.SetDelegate]
//
// # Getting a Cached Value
//
//   - [INSCache.ObjectForKey]: Returns the value associated with a given key.
//
// # Adding and Removing Cached Values
//
//   - [INSCache.SetObjectForKey]: Sets the value of the specified key in the cache.
//   - [INSCache.SetObjectForKeyCost]: Sets the value of the specified key in the cache, and associates the key-value pair with the specified cost.
//   - [INSCache.RemoveObjectForKey]: Removes the value of the specified key in the cache.
//   - [INSCache.RemoveAllObjects]: Empties the cache.
//
// See: https://developer.apple.com/documentation/Foundation/NSCache
type INSCache interface {
	objectivec.IObject

	// Topic: Managing the Name

	// The name of the cache.
	Name() string
	SetName(value string)

	// Topic: Managing Cache Size

	// The maximum number of objects the cache should hold.
	CountLimit() uint
	SetCountLimit(value uint)
	// The maximum total cost that the cache can hold before it starts evicting objects.
	TotalCostLimit() uint
	SetTotalCostLimit(value uint)

	// Topic: Managing Discardable Content

	// Whether the cache will automatically evict discardable-content objects whose content has been discarded.
	EvictsObjectsWithDiscardedContent() bool
	SetEvictsObjectsWithDiscardedContent(value bool)

	// Topic: Managing the Delegate

	// The cache’s delegate.
	Delegate() NSCacheDelegate
	SetDelegate(value NSCacheDelegate)

	// Topic: Getting a Cached Value

	// Returns the value associated with a given key.
	ObjectForKey(key objectivec.IObject) objectivec.IObject

	// Topic: Adding and Removing Cached Values

	// Sets the value of the specified key in the cache.
	SetObjectForKey(obj objectivec.IObject, key objectivec.IObject)
	// Sets the value of the specified key in the cache, and associates the key-value pair with the specified cost.
	SetObjectForKeyCost(obj objectivec.IObject, key objectivec.IObject, g uint)
	// Removes the value of the specified key in the cache.
	RemoveObjectForKey(key objectivec.IObject)
	// Empties the cache.
	RemoveAllObjects()
}

// Init initializes the instance.
func (c NSCache) Init() NSCache {
	rv := objc.Send[NSCache](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCache) Autorelease() NSCache {
	rv := objc.Send[NSCache](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCache creates a new NSCache instance.
func NewNSCache() NSCache {
	class := getNSCacheClass()
	rv := objc.Send[NSCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the value associated with a given key.
//
// key: An object identifying the value.
//
// # Return Value
// 
// The value associated with `key`, or `nil` if no value is associated with
// `key`.
//
// See: https://developer.apple.com/documentation/Foundation/NSCache/object(forKey:)
func (c NSCache) ObjectForKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("objectForKey:"), key)
	return objectivec.Object{ID: rv}
}

// Sets the value of the specified key in the cache.
//
// obj: The object to be stored in the cache.
//
// key: The key with which to associate the value.
//
// # Discussion
// 
// Unlike an [NSMutableDictionary] object, a cache does not copy the key
// objects that are put into it.
//
// See: https://developer.apple.com/documentation/Foundation/NSCache/setObject(_:forKey:)
func (c NSCache) SetObjectForKey(obj objectivec.IObject, key objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("setObject:forKey:"), obj, key)
}

// Sets the value of the specified key in the cache, and associates the
// key-value pair with the specified cost.
//
// obj: The object to store in the cache.
//
// key: The key with which to associate the value.
//
// g: The cost with which to associate the key-value pair.
//
// # Discussion
// 
// The `cost` value is used to compute a sum encompassing the costs of all the
// objects in the cache. When memory is limited or when the total cost of the
// cache eclipses the maximum allowed total cost, the cache could begin an
// eviction process to remove some of its elements. However, this eviction
// process is not in a guaranteed order. As a consequence, if you try to
// manipulate the cost values to achieve some specific behavior, the
// consequences could be detrimental to your program. Typically, the obvious
// cost is the size of the value in bytes. If that information is not readily
// available, you should not go through the trouble of trying to compute it,
// as doing so will drive up the cost of using the cache. Pass in `0` for the
// cost value if you otherwise have nothing useful to pass, or simply use the
// `` method, which does not require a `cost` value to be passed in.
// 
// Unlike an [NSMutableDictionary] object, a cache does not copy the key
// objects that are put into it.
//
// See: https://developer.apple.com/documentation/Foundation/NSCache/setObject(_:forKey:cost:)
func (c NSCache) SetObjectForKeyCost(obj objectivec.IObject, key objectivec.IObject, g uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("setObject:forKey:cost:"), obj, key, g)
}

// Removes the value of the specified key in the cache.
//
// key: The key identifying the value to be removed.
//
// See: https://developer.apple.com/documentation/Foundation/NSCache/removeObject(forKey:)
func (c NSCache) RemoveObjectForKey(key objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeObjectForKey:"), key)
}

// Empties the cache.
//
// See: https://developer.apple.com/documentation/Foundation/NSCache/removeAllObjects()
func (c NSCache) RemoveAllObjects() {
	objc.Send[objc.ID](c.ID, objc.Sel("removeAllObjects"))
}

// The name of the cache.
//
// # Discussion
// 
// The default value is an empty string (””).
//
// See: https://developer.apple.com/documentation/Foundation/NSCache/name
func (c NSCache) Name() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
func (c NSCache) SetName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setName:"), objc.String(value))
}

// The maximum number of objects the cache should hold.
//
// # Discussion
// 
// If `0`, there is no count limit. The default value is `0`.
// 
// This is not a strict limit—if the cache goes over the limit, an object in
// the cache could be evicted instantly, later, or possibly never, depending
// on the implementation details of the cache.
//
// See: https://developer.apple.com/documentation/Foundation/NSCache/countLimit
func (c NSCache) CountLimit() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("countLimit"))
	return rv
}
func (c NSCache) SetCountLimit(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setCountLimit:"), value)
}

// The maximum total cost that the cache can hold before it starts evicting
// objects.
//
// # Discussion
// 
// If `0`, there is no total cost limit. The default value is `0`.
// 
// When you add an object to the cache, you may pass in a specified cost for
// the object, such as the size in bytes of the object. If adding this object
// to the cache causes the cache’s total cost to rise above
// `totalCostLimit`, the cache may automatically evict objects until its total
// cost falls below `totalCostLimit`. The order in which the cache evicts
// objects is not guaranteed.
// 
// This is not a strict limit, and if the cache goes over the limit, an object
// in the cache could be evicted instantly, at a later point in time, or
// possibly never, all depending on the implementation details of the cache.
//
// See: https://developer.apple.com/documentation/Foundation/NSCache/totalCostLimit
func (c NSCache) TotalCostLimit() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("totalCostLimit"))
	return rv
}
func (c NSCache) SetTotalCostLimit(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setTotalCostLimit:"), value)
}

// Whether the cache will automatically evict discardable-content objects
// whose content has been discarded.
//
// # Discussion
// 
// If [true], the cache will evict a discardable-content object after its
// content is discarded. If [false], it will not. The default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCache/evictsObjectsWithDiscardedContent
func (c NSCache) EvictsObjectsWithDiscardedContent() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("evictsObjectsWithDiscardedContent"))
	return rv
}
func (c NSCache) SetEvictsObjectsWithDiscardedContent(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEvictsObjectsWithDiscardedContent:"), value)
}

// The cache’s delegate.
//
// # Discussion
// 
// The delegate must adopt the [NSCacheDelegate] protocol.
//
// See: https://developer.apple.com/documentation/Foundation/NSCache/delegate
func (c NSCache) Delegate() NSCacheDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return NSCacheDelegateObjectFromID(rv)
}
func (c NSCache) SetDelegate(value NSCacheDelegate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelegate:"), value)
}

