// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAccessibilityCustomRotorSearchParameters] class.
var (
	_NSAccessibilityCustomRotorSearchParametersClass     NSAccessibilityCustomRotorSearchParametersClass
	_NSAccessibilityCustomRotorSearchParametersClassOnce sync.Once
)

func getNSAccessibilityCustomRotorSearchParametersClass() NSAccessibilityCustomRotorSearchParametersClass {
	_NSAccessibilityCustomRotorSearchParametersClassOnce.Do(func() {
		_NSAccessibilityCustomRotorSearchParametersClass = NSAccessibilityCustomRotorSearchParametersClass{class: objc.GetClass("NSAccessibilityCustomRotorSearchParameters")}
	})
	return _NSAccessibilityCustomRotorSearchParametersClass
}

// GetNSAccessibilityCustomRotorSearchParametersClass returns the class object for NSAccessibilityCustomRotorSearchParameters.
func GetNSAccessibilityCustomRotorSearchParametersClass() NSAccessibilityCustomRotorSearchParametersClass {
	return getNSAccessibilityCustomRotorSearchParametersClass()
}

type NSAccessibilityCustomRotorSearchParametersClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAccessibilityCustomRotorSearchParametersClass) Alloc() NSAccessibilityCustomRotorSearchParameters {
	rv := objc.Send[NSAccessibilityCustomRotorSearchParameters](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Search parameters for a custom rotor.
//
// # Overview
// 
// Use these parameters to determine the next matching
// [NSAccessibilityCustomRotorItemResult].
//
// # Managing the Current Item
//
//   - [NSAccessibilityCustomRotorSearchParameters.CurrentItem]: The current item that determines where the search starts.
//   - [NSAccessibilityCustomRotorSearchParameters.SetCurrentItem]
//
// # Specifying the Filter String
//
//   - [NSAccessibilityCustomRotorSearchParameters.FilterString]: A string of text to filter the results against.
//   - [NSAccessibilityCustomRotorSearchParameters.SetFilterString]
//
// # Specifying Search Direction
//
//   - [NSAccessibilityCustomRotorSearchParameters.SearchDirection]: The direction to search for an item result.
//   - [NSAccessibilityCustomRotorSearchParameters.SetSearchDirection]
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters
type NSAccessibilityCustomRotorSearchParameters struct {
	objectivec.Object
}

// NSAccessibilityCustomRotorSearchParametersFromID constructs a [NSAccessibilityCustomRotorSearchParameters] from an objc.ID.
//
// Search parameters for a custom rotor.
func NSAccessibilityCustomRotorSearchParametersFromID(id objc.ID) NSAccessibilityCustomRotorSearchParameters {
	return NSAccessibilityCustomRotorSearchParameters{objectivec.Object{ID: id}}
}
// NOTE: NSAccessibilityCustomRotorSearchParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAccessibilityCustomRotorSearchParameters] class.
//
// # Managing the Current Item
//
//   - [INSAccessibilityCustomRotorSearchParameters.CurrentItem]: The current item that determines where the search starts.
//   - [INSAccessibilityCustomRotorSearchParameters.SetCurrentItem]
//
// # Specifying the Filter String
//
//   - [INSAccessibilityCustomRotorSearchParameters.FilterString]: A string of text to filter the results against.
//   - [INSAccessibilityCustomRotorSearchParameters.SetFilterString]
//
// # Specifying Search Direction
//
//   - [INSAccessibilityCustomRotorSearchParameters.SearchDirection]: The direction to search for an item result.
//   - [INSAccessibilityCustomRotorSearchParameters.SetSearchDirection]
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters
type INSAccessibilityCustomRotorSearchParameters interface {
	objectivec.IObject

	// Topic: Managing the Current Item

	// The current item that determines where the search starts.
	CurrentItem() INSAccessibilityCustomRotorItemResult
	SetCurrentItem(value INSAccessibilityCustomRotorItemResult)

	// Topic: Specifying the Filter String

	// A string of text to filter the results against.
	FilterString() string
	SetFilterString(value string)

	// Topic: Specifying Search Direction

	// The direction to search for an item result.
	SearchDirection() NSAccessibilityCustomRotorSearchDirection
	SetSearchDirection(value NSAccessibilityCustomRotorSearchDirection)
}

// Init initializes the instance.
func (a NSAccessibilityCustomRotorSearchParameters) Init() NSAccessibilityCustomRotorSearchParameters {
	rv := objc.Send[NSAccessibilityCustomRotorSearchParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAccessibilityCustomRotorSearchParameters) Autorelease() NSAccessibilityCustomRotorSearchParameters {
	rv := objc.Send[NSAccessibilityCustomRotorSearchParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAccessibilityCustomRotorSearchParameters creates a new NSAccessibilityCustomRotorSearchParameters instance.
func NewNSAccessibilityCustomRotorSearchParameters() NSAccessibilityCustomRotorSearchParameters {
	class := getNSAccessibilityCustomRotorSearchParametersClass()
	rv := objc.Send[NSAccessibilityCustomRotorSearchParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The current item that determines where the search starts.
//
// # Discussion
// 
// If this value is `nil`, [SearchDirection] determines the current item. A
// search direction of [NSAccessibilityCustomRotor.SearchDirection.next]
// begins the search from the first item, and a search direction of
// [NSAccessibilityCustomRotor.SearchDirection.previous] begins the search
// from the last item.
//
// [NSAccessibilityCustomRotor.SearchDirection.next]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchDirection/next
// [NSAccessibilityCustomRotor.SearchDirection.previous]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchDirection/previous
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/currentItem
func (a NSAccessibilityCustomRotorSearchParameters) CurrentItem() INSAccessibilityCustomRotorItemResult {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("currentItem"))
	return NSAccessibilityCustomRotorItemResultFromID(objc.ID(rv))
}
func (a NSAccessibilityCustomRotorSearchParameters) SetCurrentItem(value INSAccessibilityCustomRotorItemResult) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurrentItem:"), value)
}
// A string of text to filter the results against.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/filterString
func (a NSAccessibilityCustomRotorSearchParameters) FilterString() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("filterString"))
	return foundation.NSStringFromID(rv).String()
}
func (a NSAccessibilityCustomRotorSearchParameters) SetFilterString(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setFilterString:"), objc.String(value))
}
// The direction to search for an item result.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/searchDirection
func (a NSAccessibilityCustomRotorSearchParameters) SearchDirection() NSAccessibilityCustomRotorSearchDirection {
	rv := objc.Send[NSAccessibilityCustomRotorSearchDirection](a.ID, objc.Sel("searchDirection"))
	return NSAccessibilityCustomRotorSearchDirection(rv)
}
func (a NSAccessibilityCustomRotorSearchParameters) SetSearchDirection(value NSAccessibilityCustomRotorSearchDirection) {
	objc.Send[struct{}](a.ID, objc.Sel("setSearchDirection:"), value)
}

