// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAccessibilityCustomRotorItemResult] class.
var (
	_NSAccessibilityCustomRotorItemResultClass     NSAccessibilityCustomRotorItemResultClass
	_NSAccessibilityCustomRotorItemResultClassOnce sync.Once
)

func getNSAccessibilityCustomRotorItemResultClass() NSAccessibilityCustomRotorItemResultClass {
	_NSAccessibilityCustomRotorItemResultClassOnce.Do(func() {
		_NSAccessibilityCustomRotorItemResultClass = NSAccessibilityCustomRotorItemResultClass{class: objc.GetClass("NSAccessibilityCustomRotorItemResult")}
	})
	return _NSAccessibilityCustomRotorItemResultClass
}

// GetNSAccessibilityCustomRotorItemResultClass returns the class object for NSAccessibilityCustomRotorItemResult.
func GetNSAccessibilityCustomRotorItemResultClass() NSAccessibilityCustomRotorItemResultClass {
	return getNSAccessibilityCustomRotorItemResultClass()
}

type NSAccessibilityCustomRotorItemResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSAccessibilityCustomRotorItemResultClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAccessibilityCustomRotorItemResultClass) Alloc() NSAccessibilityCustomRotorItemResult {
	rv := objc.Send[NSAccessibilityCustomRotorItemResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A target accessibility element that a custom rotor references.
//
// # Creating an Item Result
//
//   - [NSAccessibilityCustomRotorItemResult.InitWithTargetElement]: Creates an item result with the specified target element.
//   - [NSAccessibilityCustomRotorItemResult.InitWithItemLoadingTokenCustomLabel]: Creates an item result with the specified item load token and custom label.
//
// # Identifying an Item Result
//
//   - [NSAccessibilityCustomRotorItemResult.TargetElement]: A target element that references an element to message for accessibility properties.
//   - [NSAccessibilityCustomRotorItemResult.ItemLoadingToken]: A token to determine which item to return.
//   - [NSAccessibilityCustomRotorItemResult.TargetRange]: A range that specifies the area of interest for text-based elements.
//   - [NSAccessibilityCustomRotorItemResult.SetTargetRange]
//   - [NSAccessibilityCustomRotorItemResult.CustomLabel]: A localized label to use instead of the default item label to describe the item result.
//   - [NSAccessibilityCustomRotorItemResult.SetCustomLabel]
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult
type NSAccessibilityCustomRotorItemResult struct {
	objectivec.Object
}

// NSAccessibilityCustomRotorItemResultFromID constructs a [NSAccessibilityCustomRotorItemResult] from an objc.ID.
//
// A target accessibility element that a custom rotor references.
func NSAccessibilityCustomRotorItemResultFromID(id objc.ID) NSAccessibilityCustomRotorItemResult {
	return NSAccessibilityCustomRotorItemResult{objectivec.Object{ID: id}}
}

// NOTE: NSAccessibilityCustomRotorItemResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAccessibilityCustomRotorItemResult] class.
//
// # Creating an Item Result
//
//   - [INSAccessibilityCustomRotorItemResult.InitWithTargetElement]: Creates an item result with the specified target element.
//   - [INSAccessibilityCustomRotorItemResult.InitWithItemLoadingTokenCustomLabel]: Creates an item result with the specified item load token and custom label.
//
// # Identifying an Item Result
//
//   - [INSAccessibilityCustomRotorItemResult.TargetElement]: A target element that references an element to message for accessibility properties.
//   - [INSAccessibilityCustomRotorItemResult.ItemLoadingToken]: A token to determine which item to return.
//   - [INSAccessibilityCustomRotorItemResult.TargetRange]: A range that specifies the area of interest for text-based elements.
//   - [INSAccessibilityCustomRotorItemResult.SetTargetRange]
//   - [INSAccessibilityCustomRotorItemResult.CustomLabel]: A localized label to use instead of the default item label to describe the item result.
//   - [INSAccessibilityCustomRotorItemResult.SetCustomLabel]
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult
type INSAccessibilityCustomRotorItemResult interface {
	objectivec.IObject

	// Topic: Creating an Item Result

	// Creates an item result with the specified target element.
	InitWithTargetElement(targetElement NSAccessibilityElement) NSAccessibilityCustomRotorItemResult
	// Creates an item result with the specified item load token and custom label.
	InitWithItemLoadingTokenCustomLabel(itemLoadingToken objectivec.IObject, customLabel string) NSAccessibilityCustomRotorItemResult

	// Topic: Identifying an Item Result

	// A target element that references an element to message for accessibility properties.
	TargetElement() NSAccessibilityElement
	// A token to determine which item to return.
	ItemLoadingToken() NSAccessibilityLoadingToken
	// A range that specifies the area of interest for text-based elements.
	TargetRange() foundation.NSRange
	SetTargetRange(value foundation.NSRange)
	// A localized label to use instead of the default item label to describe the item result.
	CustomLabel() string
	SetCustomLabel(value string)

	// The current item that determines where the search starts.
	CurrentItem() INSAccessibilityCustomRotorItemResult
	SetCurrentItem(value INSAccessibilityCustomRotorItemResult)
}

// Init initializes the instance.
func (a NSAccessibilityCustomRotorItemResult) Init() NSAccessibilityCustomRotorItemResult {
	rv := objc.Send[NSAccessibilityCustomRotorItemResult](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAccessibilityCustomRotorItemResult) Autorelease() NSAccessibilityCustomRotorItemResult {
	rv := objc.Send[NSAccessibilityCustomRotorItemResult](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAccessibilityCustomRotorItemResult creates a new NSAccessibilityCustomRotorItemResult instance.
func NewNSAccessibilityCustomRotorItemResult() NSAccessibilityCustomRotorItemResult {
	class := getNSAccessibilityCustomRotorItemResultClass()
	rv := objc.Send[NSAccessibilityCustomRotorItemResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an item result with the specified item load token and custom label.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/init(itemLoadingToken:customLabel:)
func NewAccessibilityCustomRotorItemResultWithItemLoadingTokenCustomLabel(itemLoadingToken objectivec.IObject, customLabel string) NSAccessibilityCustomRotorItemResult {
	instance := getNSAccessibilityCustomRotorItemResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItemLoadingToken:customLabel:"), itemLoadingToken, objc.String(customLabel))
	return NSAccessibilityCustomRotorItemResultFromID(rv)
}

// Creates an item result with the specified target element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/init(targetElement:)
func NewAccessibilityCustomRotorItemResultWithTargetElement(targetElement NSAccessibilityElement) NSAccessibilityCustomRotorItemResult {
	instance := getNSAccessibilityCustomRotorItemResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTargetElement:"), targetElement)
	return NSAccessibilityCustomRotorItemResultFromID(rv)
}

// Creates an item result with the specified target element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/init(targetElement:)
func (a NSAccessibilityCustomRotorItemResult) InitWithTargetElement(targetElement NSAccessibilityElement) NSAccessibilityCustomRotorItemResult {
	rv := objc.Send[NSAccessibilityCustomRotorItemResult](a.ID, objc.Sel("initWithTargetElement:"), targetElement)
	return rv
}

// Creates an item result with the specified item load token and custom label.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/init(itemLoadingToken:customLabel:)
func (a NSAccessibilityCustomRotorItemResult) InitWithItemLoadingTokenCustomLabel(itemLoadingToken objectivec.IObject, customLabel string) NSAccessibilityCustomRotorItemResult {
	rv := objc.Send[NSAccessibilityCustomRotorItemResult](a.ID, objc.Sel("initWithItemLoadingToken:customLabel:"), itemLoadingToken, objc.String(customLabel))
	return rv
}

// A target element that references an element to message for accessibility
// properties.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/targetElement
func (a NSAccessibilityCustomRotorItemResult) TargetElement() NSAccessibilityElement {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("targetElement"))
	return NSAccessibilityElementFromID(objc.ID(rv))
}

// A token to determine which item to return.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/itemLoadingToken
func (a NSAccessibilityCustomRotorItemResult) ItemLoadingToken() NSAccessibilityLoadingToken {
	rv := objc.Send[NSAccessibilityLoadingToken](a.ID, objc.Sel("itemLoadingToken"))
	return NSAccessibilityLoadingToken(rv)
}

// A range that specifies the area of interest for text-based elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/targetRange
func (a NSAccessibilityCustomRotorItemResult) TargetRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](a.ID, objc.Sel("targetRange"))
	return foundation.NSRange(rv)
}
func (a NSAccessibilityCustomRotorItemResult) SetTargetRange(value foundation.NSRange) {
	objc.Send[struct{}](a.ID, objc.Sel("setTargetRange:"), value)
}

// A localized label to use instead of the default item label to describe the
// item result.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/customLabel
func (a NSAccessibilityCustomRotorItemResult) CustomLabel() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("customLabel"))
	return foundation.NSStringFromID(rv).String()
}
func (a NSAccessibilityCustomRotorItemResult) SetCustomLabel(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setCustomLabel:"), objc.String(value))
}

// The current item that determines where the search starts.
//
// See: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/searchparameters/currentitem
func (a NSAccessibilityCustomRotorItemResult) CurrentItem() INSAccessibilityCustomRotorItemResult {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("currentItem"))
	return NSAccessibilityCustomRotorItemResultFromID(objc.ID(rv))
}
func (a NSAccessibilityCustomRotorItemResult) SetCurrentItem(value INSAccessibilityCustomRotorItemResult) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurrentItem:"), value)
}
