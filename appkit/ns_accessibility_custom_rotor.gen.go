// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAccessibilityCustomRotor] class.
var (
	_NSAccessibilityCustomRotorClass     NSAccessibilityCustomRotorClass
	_NSAccessibilityCustomRotorClassOnce sync.Once
)

func getNSAccessibilityCustomRotorClass() NSAccessibilityCustomRotorClass {
	_NSAccessibilityCustomRotorClassOnce.Do(func() {
		_NSAccessibilityCustomRotorClass = NSAccessibilityCustomRotorClass{class: objc.GetClass("NSAccessibilityCustomRotor")}
	})
	return _NSAccessibilityCustomRotorClass
}

// GetNSAccessibilityCustomRotorClass returns the class object for NSAccessibilityCustomRotor.
func GetNSAccessibilityCustomRotorClass() NSAccessibilityCustomRotorClass {
	return getNSAccessibilityCustomRotorClass()
}

type NSAccessibilityCustomRotorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAccessibilityCustomRotorClass) Alloc() NSAccessibilityCustomRotor {
	rv := objc.Send[NSAccessibilityCustomRotor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A context-sensitive function that helps VoiceOver users find the next
// instance of a related accessibility element.
//
// # Overview
// 
// Assistive apps, like VoiceOver, provide interfaces to quickly search apps
// for content of a specific type. For example, in a web browser, a user can
// quickly explore a list of navigational links or buttons using VoiceOver’s
// content menus.
// 
// [NSAccessibilityCustomRotor] provides a way for apps to vend their own
// content menus. For example, Pages can create a custom rotor that allows
// assistive apps to search the Pages document for all headings.
//
// # Creating a Rotor
//
//   - [NSAccessibilityCustomRotor.InitWithLabelItemSearchDelegate]: Creates a custom rotor with the specified label and item search delegate.
//   - [NSAccessibilityCustomRotor.InitWithRotorTypeItemSearchDelegate]: Creates a custom rotor with the specified rotor type and item search delegate.
//
// # Navigating to the Next Item
//
//   - [NSAccessibilityCustomRotor.ItemSearchDelegate]: The delegate for finding the next item result.
//   - [NSAccessibilityCustomRotor.SetItemSearchDelegate]
//
// # Loading the Item
//
//   - [NSAccessibilityCustomRotor.ItemLoadingDelegate]: The delegate for loading item results that don’t have a backing UI element at loading time.
//   - [NSAccessibilityCustomRotor.SetItemLoadingDelegate]
//
// # Getting the Rotor Type
//
//   - [NSAccessibilityCustomRotor.Type]: The type of content that the rotor represents.
//   - [NSAccessibilityCustomRotor.SetType]
//
// # Identifying the Rotor
//
//   - [NSAccessibilityCustomRotor.Label]: The localized label that assistive apps use to describe the custom rotor.
//   - [NSAccessibilityCustomRotor.SetLabel]
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor
type NSAccessibilityCustomRotor struct {
	objectivec.Object
}

// NSAccessibilityCustomRotorFromID constructs a [NSAccessibilityCustomRotor] from an objc.ID.
//
// A context-sensitive function that helps VoiceOver users find the next
// instance of a related accessibility element.
func NSAccessibilityCustomRotorFromID(id objc.ID) NSAccessibilityCustomRotor {
	return NSAccessibilityCustomRotor{objectivec.Object{ID: id}}
}
// NOTE: NSAccessibilityCustomRotor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAccessibilityCustomRotor] class.
//
// # Creating a Rotor
//
//   - [INSAccessibilityCustomRotor.InitWithLabelItemSearchDelegate]: Creates a custom rotor with the specified label and item search delegate.
//   - [INSAccessibilityCustomRotor.InitWithRotorTypeItemSearchDelegate]: Creates a custom rotor with the specified rotor type and item search delegate.
//
// # Navigating to the Next Item
//
//   - [INSAccessibilityCustomRotor.ItemSearchDelegate]: The delegate for finding the next item result.
//   - [INSAccessibilityCustomRotor.SetItemSearchDelegate]
//
// # Loading the Item
//
//   - [INSAccessibilityCustomRotor.ItemLoadingDelegate]: The delegate for loading item results that don’t have a backing UI element at loading time.
//   - [INSAccessibilityCustomRotor.SetItemLoadingDelegate]
//
// # Getting the Rotor Type
//
//   - [INSAccessibilityCustomRotor.Type]: The type of content that the rotor represents.
//   - [INSAccessibilityCustomRotor.SetType]
//
// # Identifying the Rotor
//
//   - [INSAccessibilityCustomRotor.Label]: The localized label that assistive apps use to describe the custom rotor.
//   - [INSAccessibilityCustomRotor.SetLabel]
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor
type INSAccessibilityCustomRotor interface {
	objectivec.IObject

	// Topic: Creating a Rotor

	// Creates a custom rotor with the specified label and item search delegate.
	InitWithLabelItemSearchDelegate(label string, itemSearchDelegate NSAccessibilityCustomRotorItemSearchDelegate) NSAccessibilityCustomRotor
	// Creates a custom rotor with the specified rotor type and item search delegate.
	InitWithRotorTypeItemSearchDelegate(rotorType NSAccessibilityCustomRotorType, itemSearchDelegate NSAccessibilityCustomRotorItemSearchDelegate) NSAccessibilityCustomRotor

	// Topic: Navigating to the Next Item

	// The delegate for finding the next item result.
	ItemSearchDelegate() NSAccessibilityCustomRotorItemSearchDelegate
	SetItemSearchDelegate(value NSAccessibilityCustomRotorItemSearchDelegate)

	// Topic: Loading the Item

	// The delegate for loading item results that don’t have a backing UI element at loading time.
	ItemLoadingDelegate() NSAccessibilityElementLoading
	SetItemLoadingDelegate(value NSAccessibilityElementLoading)

	// Topic: Getting the Rotor Type

	// The type of content that the rotor represents.
	Type() NSAccessibilityCustomRotorType
	SetType(value NSAccessibilityCustomRotorType)

	// Topic: Identifying the Rotor

	// The localized label that assistive apps use to describe the custom rotor.
	Label() string
	SetLabel(value string)
}

// Init initializes the instance.
func (a NSAccessibilityCustomRotor) Init() NSAccessibilityCustomRotor {
	rv := objc.Send[NSAccessibilityCustomRotor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAccessibilityCustomRotor) Autorelease() NSAccessibilityCustomRotor {
	rv := objc.Send[NSAccessibilityCustomRotor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAccessibilityCustomRotor creates a new NSAccessibilityCustomRotor instance.
func NewNSAccessibilityCustomRotor() NSAccessibilityCustomRotor {
	class := getNSAccessibilityCustomRotorClass()
	rv := objc.Send[NSAccessibilityCustomRotor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a custom rotor with the specified label and item search delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/init(label:itemSearchDelegate:)
func NewAccessibilityCustomRotorWithLabelItemSearchDelegate(label string, itemSearchDelegate NSAccessibilityCustomRotorItemSearchDelegate) NSAccessibilityCustomRotor {
	instance := getNSAccessibilityCustomRotorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLabel:itemSearchDelegate:"), objc.String(label), itemSearchDelegate)
	return NSAccessibilityCustomRotorFromID(rv)
}

// Creates a custom rotor with the specified rotor type and item search
// delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/init(rotorType:itemSearchDelegate:)
func NewAccessibilityCustomRotorWithRotorTypeItemSearchDelegate(rotorType NSAccessibilityCustomRotorType, itemSearchDelegate NSAccessibilityCustomRotorItemSearchDelegate) NSAccessibilityCustomRotor {
	instance := getNSAccessibilityCustomRotorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRotorType:itemSearchDelegate:"), rotorType, itemSearchDelegate)
	return NSAccessibilityCustomRotorFromID(rv)
}

// Creates a custom rotor with the specified label and item search delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/init(label:itemSearchDelegate:)
func (a NSAccessibilityCustomRotor) InitWithLabelItemSearchDelegate(label string, itemSearchDelegate NSAccessibilityCustomRotorItemSearchDelegate) NSAccessibilityCustomRotor {
	rv := objc.Send[NSAccessibilityCustomRotor](a.ID, objc.Sel("initWithLabel:itemSearchDelegate:"), objc.String(label), itemSearchDelegate)
	return rv
}

// Creates a custom rotor with the specified rotor type and item search
// delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/init(rotorType:itemSearchDelegate:)
func (a NSAccessibilityCustomRotor) InitWithRotorTypeItemSearchDelegate(rotorType NSAccessibilityCustomRotorType, itemSearchDelegate NSAccessibilityCustomRotorItemSearchDelegate) NSAccessibilityCustomRotor {
	rv := objc.Send[NSAccessibilityCustomRotor](a.ID, objc.Sel("initWithRotorType:itemSearchDelegate:"), rotorType, itemSearchDelegate)
	return rv
}

// The delegate for finding the next item result.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/itemSearchDelegate
func (a NSAccessibilityCustomRotor) ItemSearchDelegate() NSAccessibilityCustomRotorItemSearchDelegate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("itemSearchDelegate"))
	return NSAccessibilityCustomRotorItemSearchDelegateObjectFromID(rv)
}
func (a NSAccessibilityCustomRotor) SetItemSearchDelegate(value NSAccessibilityCustomRotorItemSearchDelegate) {
	objc.Send[struct{}](a.ID, objc.Sel("setItemSearchDelegate:"), value)
}

// The delegate for loading item results that don’t have a backing UI
// element at loading time.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/itemLoadingDelegate
func (a NSAccessibilityCustomRotor) ItemLoadingDelegate() NSAccessibilityElementLoading {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("itemLoadingDelegate"))
	return NSAccessibilityElementLoadingObjectFromID(rv)
}
func (a NSAccessibilityCustomRotor) SetItemLoadingDelegate(value NSAccessibilityElementLoading) {
	objc.Send[struct{}](a.ID, objc.Sel("setItemLoadingDelegate:"), value)
}

// The type of content that the rotor represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/type
func (a NSAccessibilityCustomRotor) Type() NSAccessibilityCustomRotorType {
	rv := objc.Send[NSAccessibilityCustomRotorType](a.ID, objc.Sel("type"))
	return NSAccessibilityCustomRotorType(rv)
}
func (a NSAccessibilityCustomRotor) SetType(value NSAccessibilityCustomRotorType) {
	objc.Send[struct{}](a.ID, objc.Sel("setType:"), value)
}

// The localized label that assistive apps use to describe the custom rotor.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/label
func (a NSAccessibilityCustomRotor) Label() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (a NSAccessibilityCustomRotor) SetLabel(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setLabel:"), objc.String(value))
}

