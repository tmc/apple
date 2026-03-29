// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSBindingSelectionMarker] class.
var (
	_NSBindingSelectionMarkerClass     NSBindingSelectionMarkerClass
	_NSBindingSelectionMarkerClassOnce sync.Once
)

func getNSBindingSelectionMarkerClass() NSBindingSelectionMarkerClass {
	_NSBindingSelectionMarkerClassOnce.Do(func() {
		_NSBindingSelectionMarkerClass = NSBindingSelectionMarkerClass{class: objc.GetClass("NSBindingSelectionMarker")}
	})
	return _NSBindingSelectionMarkerClass
}

// GetNSBindingSelectionMarkerClass returns the class object for NSBindingSelectionMarker.
func GetNSBindingSelectionMarkerClass() NSBindingSelectionMarkerClass {
	return getNSBindingSelectionMarkerClass()
}

type NSBindingSelectionMarkerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBindingSelectionMarkerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBindingSelectionMarkerClass) Alloc() NSBindingSelectionMarker {
	rv := objc.Send[NSBindingSelectionMarker](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker
type NSBindingSelectionMarker struct {
	objectivec.Object
}

// NSBindingSelectionMarkerFromID constructs a [NSBindingSelectionMarker] from an objc.ID.
func NSBindingSelectionMarkerFromID(id objc.ID) NSBindingSelectionMarker {
	return NSBindingSelectionMarker{objectivec.Object{ID: id}}
}
// NOTE: NSBindingSelectionMarker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBindingSelectionMarker] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker
type INSBindingSelectionMarker interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b NSBindingSelectionMarker) Init() NSBindingSelectionMarker {
	rv := objc.Send[NSBindingSelectionMarker](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBindingSelectionMarker) Autorelease() NSBindingSelectionMarker {
	rv := objc.Send[NSBindingSelectionMarker](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBindingSelectionMarker creates a new NSBindingSelectionMarker instance.
func NewNSBindingSelectionMarker() NSBindingSelectionMarker {
	class := getNSBindingSelectionMarkerClass()
	rv := objc.Send[NSBindingSelectionMarker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/defaultPlaceholder(for:on:withBinding:)
func (_NSBindingSelectionMarkerClass NSBindingSelectionMarkerClass) DefaultPlaceholderForMarkerOnClassWithBinding(marker INSBindingSelectionMarker, objectClass objc.Class, binding string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSBindingSelectionMarkerClass.class), objc.Sel("defaultPlaceholderForMarker:onClass:withBinding:"), marker, objectClass, objc.String(binding))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/setDefaultPlaceholder(_:for:on:withBinding:)
func (_NSBindingSelectionMarkerClass NSBindingSelectionMarkerClass) SetDefaultPlaceholderForMarkerOnClassWithBinding(placeholder objectivec.IObject, marker INSBindingSelectionMarker, objectClass objc.Class, binding string) {
	objc.Send[objc.ID](objc.ID(_NSBindingSelectionMarkerClass.class), objc.Sel("setDefaultPlaceholder:forMarker:onClass:withBinding:"), placeholder, marker, objectClass, objc.String(binding))
}

// See: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/multipleValues
func (_NSBindingSelectionMarkerClass NSBindingSelectionMarkerClass) MultipleValuesSelectionMarker() NSBindingSelectionMarker {
	rv := objc.Send[objc.ID](objc.ID(_NSBindingSelectionMarkerClass.class), objc.Sel("multipleValuesSelectionMarker"))
	return NSBindingSelectionMarkerFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/noSelection
func (_NSBindingSelectionMarkerClass NSBindingSelectionMarkerClass) NoSelectionMarker() NSBindingSelectionMarker {
	rv := objc.Send[objc.ID](objc.ID(_NSBindingSelectionMarkerClass.class), objc.Sel("noSelectionMarker"))
	return NSBindingSelectionMarkerFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/notApplicable
func (_NSBindingSelectionMarkerClass NSBindingSelectionMarkerClass) NotApplicableSelectionMarker() NSBindingSelectionMarker {
	rv := objc.Send[objc.ID](objc.ID(_NSBindingSelectionMarkerClass.class), objc.Sel("notApplicableSelectionMarker"))
	return NSBindingSelectionMarkerFromID(objc.ID(rv))
}

