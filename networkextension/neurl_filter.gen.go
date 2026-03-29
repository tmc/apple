// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEURLFilter] class.
var (
	_NEURLFilterClass     NEURLFilterClass
	_NEURLFilterClassOnce sync.Once
)

func getNEURLFilterClass() NEURLFilterClass {
	_NEURLFilterClassOnce.Do(func() {
		_NEURLFilterClass = NEURLFilterClass{class: objc.GetClass("NEURLFilter")}
	})
	return _NEURLFilterClass
}

// GetNEURLFilterClass returns the class object for NEURLFilter.
func GetNEURLFilterClass() NEURLFilterClass {
	return getNEURLFilterClass()
}

type NEURLFilterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEURLFilterClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEURLFilterClass) Alloc() NEURLFilter {
	rv := objc.Send[NEURLFilter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A class used to voluntarily validate URLs for apps that don’t use WebKit
// or the URL session API.
//
// # Overview
// 
// When using networking frameworks other than WebKit or Foundation’s
// [URLSession], use the [NEURLFilter] API to evaluate URLs before potentially
// connecting to a restricted or malicious site. Call the class method
// [verdict(for:)] to check a URL and honor the “allow” or “deny”
// verdict. Don’t connect to any URL that receives a “deny” verdict.
//
// [URLSession]: https://developer.apple.com/documentation/Foundation/URLSession
// [verdict(for:)]: https://developer.apple.com/documentation/NetworkExtension/NEURLFilter/verdict(for:)
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEURLFilter
type NEURLFilter struct {
	objectivec.Object
}

// NEURLFilterFromID constructs a [NEURLFilter] from an objc.ID.
//
// A class used to voluntarily validate URLs for apps that don’t use WebKit
// or the URL session API.
func NEURLFilterFromID(id objc.ID) NEURLFilter {
	return NEURLFilter{objectivec.Object{ID: id}}
}
// NOTE: NEURLFilter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEURLFilter] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEURLFilter
type INEURLFilter interface {
	objectivec.IObject
}

// Init initializes the instance.
func (u NEURLFilter) Init() NEURLFilter {
	rv := objc.Send[NEURLFilter](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NEURLFilter) Autorelease() NEURLFilter {
	rv := objc.Send[NEURLFilter](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEURLFilter creates a new NEURLFilter instance.
func NewNEURLFilter() NEURLFilter {
	class := getNEURLFilterClass()
	rv := objc.Send[NEURLFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

