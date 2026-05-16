// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/coreservices"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCPhysicalInputElementCollection] class.
var (
	_GCPhysicalInputElementCollectionClass     GCPhysicalInputElementCollectionClass
	_GCPhysicalInputElementCollectionClassOnce sync.Once
)

func getGCPhysicalInputElementCollectionClass() GCPhysicalInputElementCollectionClass {
	_GCPhysicalInputElementCollectionClassOnce.Do(func() {
		_GCPhysicalInputElementCollectionClass = GCPhysicalInputElementCollectionClass{class: objc.GetClass("GCPhysicalInputElementCollection")}
	})
	return _GCPhysicalInputElementCollectionClass
}

// GetGCPhysicalInputElementCollectionClass returns the class object for GCPhysicalInputElementCollection.
func GetGCPhysicalInputElementCollectionClass() GCPhysicalInputElementCollectionClass {
	return getGCPhysicalInputElementCollectionClass()
}

type GCPhysicalInputElementCollectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCPhysicalInputElementCollectionClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCPhysicalInputElementCollectionClass) Alloc() GCPhysicalInputElementCollection {
	rv := objc.Send[GCPhysicalInputElementCollection](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A collection of physical input elements.
//
// # Getting elements in the collection
//
//   - [GCPhysicalInputElementCollection.ElementEnumerator]: Returns an enumerator to iterate the elements in the collection.
//   - [GCPhysicalInputElementCollection.Count]: The number of elements in the collection.
//
// # Accessing elements by key and alias
//
//   - [GCPhysicalInputElementCollection.ObjectForKeyedSubscript]: Returns the element in the collection for the specified key.
//   - [GCPhysicalInputElementCollection.ElementForAlias]: Returns the element in the collection that uses the specified alias.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class
type GCPhysicalInputElementCollection struct {
	objectivec.Object
}

// GCPhysicalInputElementCollectionFromID constructs a [GCPhysicalInputElementCollection] from an objc.ID.
//
// A collection of physical input elements.
func GCPhysicalInputElementCollectionFromID(id objc.ID) GCPhysicalInputElementCollection {
	return GCPhysicalInputElementCollection{objectivec.Object{ID: id}}
}

// NOTE: GCPhysicalInputElementCollection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCPhysicalInputElementCollection] class.
//
// # Getting elements in the collection
//
//   - [IGCPhysicalInputElementCollection.ElementEnumerator]: Returns an enumerator to iterate the elements in the collection.
//   - [IGCPhysicalInputElementCollection.Count]: The number of elements in the collection.
//
// # Accessing elements by key and alias
//
//   - [IGCPhysicalInputElementCollection.ObjectForKeyedSubscript]: Returns the element in the collection for the specified key.
//   - [IGCPhysicalInputElementCollection.ElementForAlias]: Returns the element in the collection that uses the specified alias.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class
type IGCPhysicalInputElementCollection interface {
	objectivec.IObject

	// Topic: Getting elements in the collection

	// Returns an enumerator to iterate the elements in the collection.
	ElementEnumerator() foundation.NSEnumerator
	// The number of elements in the collection.
	Count() uint

	// Topic: Accessing elements by key and alias

	// Returns the element in the collection for the specified key.
	ObjectForKeyedSubscript(key coreservices.Key) objectivec.IObject
	// Returns the element in the collection that uses the specified alias.
	ElementForAlias(alias coreservices.Key) objectivec.IObject
}

// Init initializes the instance.
func (g GCPhysicalInputElementCollection) Init() GCPhysicalInputElementCollection {
	rv := objc.Send[GCPhysicalInputElementCollection](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCPhysicalInputElementCollection) Autorelease() GCPhysicalInputElementCollection {
	rv := objc.Send[GCPhysicalInputElementCollection](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCPhysicalInputElementCollection creates a new GCPhysicalInputElementCollection instance.
func NewGCPhysicalInputElementCollection() GCPhysicalInputElementCollection {
	class := getGCPhysicalInputElementCollectionClass()
	rv := objc.Send[GCPhysicalInputElementCollection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an enumerator to iterate the elements in the collection.
//
// # Return Value
//
// An enumerator for the collection.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class/elementEnumerator
func (g GCPhysicalInputElementCollection) ElementEnumerator() foundation.NSEnumerator {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("elementEnumerator"))
	return foundation.NSEnumeratorFromID(rv)
}

// Returns the element in the collection for the specified key.
//
// key: The key that identifies the element.
//
// # Return Value
//
// An element in the collection.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class/objectForKeyedSubscript:
func (g GCPhysicalInputElementCollection) ObjectForKeyedSubscript(key coreservices.Key) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return objectivec.Object{ID: rv}
}

// Returns the element in the collection that uses the specified alias.
//
// alias: An alias for the element.
//
// # Return Value
//
// An element in the collection.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class/elementForAlias:
func (g GCPhysicalInputElementCollection) ElementForAlias(alias coreservices.Key) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("elementForAlias:"), alias)
	return objectivec.Object{ID: rv}
}

// The number of elements in the collection.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class/count
func (g GCPhysicalInputElementCollection) Count() uint {
	rv := objc.Send[uint](g.ID, objc.Sel("count"))
	return rv
}
