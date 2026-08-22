// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAtomicStoreCacheNode] class.
var (
	_NSAtomicStoreCacheNodeClass     NSAtomicStoreCacheNodeClass
	_NSAtomicStoreCacheNodeClassOnce sync.Once
)

func getNSAtomicStoreCacheNodeClass() NSAtomicStoreCacheNodeClass {
	_NSAtomicStoreCacheNodeClassOnce.Do(func() {
		_NSAtomicStoreCacheNodeClass = NSAtomicStoreCacheNodeClass{class: objc.GetClass("NSAtomicStoreCacheNode")}
	})
	return _NSAtomicStoreCacheNodeClass
}

// GetNSAtomicStoreCacheNodeClass returns the class object for NSAtomicStoreCacheNode.
func GetNSAtomicStoreCacheNodeClass() NSAtomicStoreCacheNodeClass {
	return getNSAtomicStoreCacheNodeClass()
}

type NSAtomicStoreCacheNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSAtomicStoreCacheNodeClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAtomicStoreCacheNodeClass) Alloc() NSAtomicStoreCacheNode {
	rv := objc.Send[NSAtomicStoreCacheNode](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A concrete class that you use to represent basic nodes in a Core Data
// atomic store.
//
// # Overview
//
// A node represents a single record in a persistent store.
//
// You can subclass [NSAtomicStoreCacheNode] to provide custom behavior.
//
// # Initializing a Cache Node
//
//   - [NSAtomicStoreCacheNode.InitWithObjectID]: Returns a cache node for the given managed object ID.
//
// # Managing Node Data
//
//   - [NSAtomicStoreCacheNode.ObjectID]: The managed object ID of the node.
//   - [NSAtomicStoreCacheNode.PropertyCache]: The property cache dictionary of the node.
//   - [NSAtomicStoreCacheNode.SetPropertyCache]
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStoreCacheNode
type NSAtomicStoreCacheNode struct {
	objectivec.Object
}

// NSAtomicStoreCacheNodeFromID constructs a [NSAtomicStoreCacheNode] from an objc.ID.
//
// A concrete class that you use to represent basic nodes in a Core Data
// atomic store.
func NSAtomicStoreCacheNodeFromID(id objc.ID) NSAtomicStoreCacheNode {
	return NSAtomicStoreCacheNode{objectivec.Object{ID: id}}
}

// NOTE: NSAtomicStoreCacheNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAtomicStoreCacheNode] class.
//
// # Initializing a Cache Node
//
//   - [INSAtomicStoreCacheNode.InitWithObjectID]: Returns a cache node for the given managed object ID.
//
// # Managing Node Data
//
//   - [INSAtomicStoreCacheNode.ObjectID]: The managed object ID of the node.
//   - [INSAtomicStoreCacheNode.PropertyCache]: The property cache dictionary of the node.
//   - [INSAtomicStoreCacheNode.SetPropertyCache]
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStoreCacheNode
type INSAtomicStoreCacheNode interface {
	objectivec.IObject

	// Topic: Initializing a Cache Node

	// Returns a cache node for the given managed object ID.
	InitWithObjectID(moid INSManagedObjectID) NSAtomicStoreCacheNode

	// Topic: Managing Node Data

	// The managed object ID of the node.
	ObjectID() INSManagedObjectID
	// The property cache dictionary of the node.
	PropertyCache() foundation.INSDictionary
	SetPropertyCache(value foundation.INSDictionary)
}

// Init initializes the instance.
func (a NSAtomicStoreCacheNode) Init() NSAtomicStoreCacheNode {
	rv := objc.Send[NSAtomicStoreCacheNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAtomicStoreCacheNode) Autorelease() NSAtomicStoreCacheNode {
	rv := objc.Send[NSAtomicStoreCacheNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAtomicStoreCacheNode creates a new NSAtomicStoreCacheNode instance.
func NewNSAtomicStoreCacheNode() NSAtomicStoreCacheNode {
	class := getNSAtomicStoreCacheNodeClass()
	rv := objc.Send[NSAtomicStoreCacheNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a cache node for the given managed object ID.
//
// moid: A managed object ID.
//
// # Return Value
//
// A cache node for the given managed object ID, or `nil` if the node could
// not be initialized.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStoreCacheNode/init(objectID:)
func NewAtomicStoreCacheNodeWithObjectID(moid INSManagedObjectID) NSAtomicStoreCacheNode {
	instance := getNSAtomicStoreCacheNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObjectID:"), moid)
	return NSAtomicStoreCacheNodeFromID(rv)
}

// Returns a cache node for the given managed object ID.
//
// moid: A managed object ID.
//
// # Return Value
//
// A cache node for the given managed object ID, or `nil` if the node could
// not be initialized.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStoreCacheNode/init(objectID:)
func (a NSAtomicStoreCacheNode) InitWithObjectID(moid INSManagedObjectID) NSAtomicStoreCacheNode {
	rv := objc.Send[NSAtomicStoreCacheNode](a.ID, objc.Sel("initWithObjectID:"), moid)
	return rv
}

// The managed object ID of the node.
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStoreCacheNode/objectID
func (a NSAtomicStoreCacheNode) ObjectID() INSManagedObjectID {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("objectID"))
	return NSManagedObjectIDFromID(objc.ID(rv))
}

// The property cache dictionary of the node.
//
// # Discussion
//
// This dictionary is used by [NSAtomicStoreCacheNode.ValueForKey] and
// [NSAtomicStoreCacheNode.SetValueForKey] for property values. This property
// is `nil` unless it has been explicitly set or non-`nil` values have been
// set for keys using [NSAtomicStoreCacheNode.SetValueForKey].
//
// See: https://developer.apple.com/documentation/CoreData/NSAtomicStoreCacheNode/propertyCache
func (a NSAtomicStoreCacheNode) PropertyCache() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("propertyCache"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a NSAtomicStoreCacheNode) SetPropertyCache(value foundation.INSDictionary) {
	objc.Send[struct{}](a.ID, objc.Sel("setPropertyCache:"), value)
}
