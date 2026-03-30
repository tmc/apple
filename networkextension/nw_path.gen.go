// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWPath] class.
var (
	_NWPathClass     NWPathClass
	_NWPathClassOnce sync.Once
)

func getNWPathClass() NWPathClass {
	_NWPathClassOnce.Do(func() {
		_NWPathClass = NWPathClass{class: objc.GetClass("NWPath")}
	})
	return _NWPathClass
}

// GetNWPathClass returns the class object for NWPath.
func GetNWPathClass() NWPathClass {
	return getNWPathClass()
}

type NWPathClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWPathClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWPathClass) Alloc() NWPath {
	rv := objc.Send[NWPath](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The path made by a network connection, including information about its
// viability.
//
// # Overview
//
// For example, if the path status is [NWPathStatusSatisfied], then a
// connection attempt will be made.
//
// When attached to a specific connection, a path takes all of the connection
// parameters into account. For example, if the route for a connection changes
// or is removed, the path will reflect that change. Note that every path is
// evaluated within the context of the process it is running in, and may be
// different across processes.
//
// [NWPath] is a static object, and properties of the path will never change.
// To monitor changing network status, use Key-Value Observing (KVO) to watch
// a path property on another object. For information about KVO, see
// [Key-Value Observing Programming Guide].
//
// # Getting network path properties
//
//   - [NWPath.Status]: The evaluated status of the network path.
//   - [NWPath.Expensive]: A Boolean that indicates whether or not the path uses an expensive interface.
//   - [NWPath.Constrained]: A Boolean that indicates whether or not the path uses a constrained interface, such as when using low-data mode.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWPath
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
type NWPath struct {
	objectivec.Object
}

// NWPathFromID constructs a [NWPath] from an objc.ID.
//
// The path made by a network connection, including information about its
// viability.
func NWPathFromID(id objc.ID) NWPath {
	return NWPath{objectivec.Object{ID: id}}
}

// NOTE: NWPath adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NWPath] class.
//
// # Getting network path properties
//
//   - [INWPath.Status]: The evaluated status of the network path.
//   - [INWPath.Expensive]: A Boolean that indicates whether or not the path uses an expensive interface.
//   - [INWPath.Constrained]: A Boolean that indicates whether or not the path uses a constrained interface, such as when using low-data mode.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWPath
type INWPath interface {
	objectivec.IObject

	// Topic: Getting network path properties

	// The evaluated status of the network path.
	Status() NWPathStatus
	// A Boolean that indicates whether or not the path uses an expensive interface.
	Expensive() bool
	// A Boolean that indicates whether or not the path uses a constrained interface, such as when using low-data mode.
	Constrained() bool
}

// Init initializes the instance.
func (n NWPath) Init() NWPath {
	rv := objc.Send[NWPath](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWPath) Autorelease() NWPath {
	rv := objc.Send[NWPath](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWPath creates a new NWPath instance.
func NewNWPath() NWPath {
	class := getNWPathClass()
	rv := objc.Send[NWPath](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The evaluated status of the network path.
//
// # Discussion
//
// The status of a path indicates whether or not the process is able to make
// connection attempts to any, or a specific, network endpoint. A satisfied
// status does not guarantee that a connection will be successful, but it does
// ensure that there is some interface over which an attempt can be made.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWPath/status
func (n NWPath) Status() NWPathStatus {
	rv := objc.Send[NWPathStatus](n.ID, objc.Sel("status"))
	return NWPathStatus(rv)
}

// A Boolean that indicates whether or not the path uses an expensive
// interface.
//
// # Discussion
//
// Returns YES is the path uses an interface that is considered expensive,
// such as when using a cellular data plan.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWPath/isExpensive
func (n NWPath) Expensive() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isExpensive"))
	return rv
}

// A Boolean that indicates whether or not the path uses a constrained
// interface, such as when using low-data mode.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWPath/isConstrained
func (n NWPath) Constrained() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isConstrained"))
	return rv
}
