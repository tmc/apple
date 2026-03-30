// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMutableFontCollection] class.
var (
	_NSMutableFontCollectionClass     NSMutableFontCollectionClass
	_NSMutableFontCollectionClassOnce sync.Once
)

func getNSMutableFontCollectionClass() NSMutableFontCollectionClass {
	_NSMutableFontCollectionClassOnce.Do(func() {
		_NSMutableFontCollectionClass = NSMutableFontCollectionClass{class: objc.GetClass("NSMutableFontCollection")}
	})
	return _NSMutableFontCollectionClass
}

// GetNSMutableFontCollectionClass returns the class object for NSMutableFontCollection.
func GetNSMutableFontCollectionClass() NSMutableFontCollectionClass {
	return getNSMutableFontCollectionClass()
}

type NSMutableFontCollectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMutableFontCollectionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableFontCollectionClass) Alloc() NSMutableFontCollection {
	rv := objc.Send[NSMutableFontCollection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A mutable collection of font descriptors taken together as a single object.
//
// # Overview
//
// You can use this class to modify the search queries for the font
// descriptors used by the parent [NSFontCollection] class.
//
// # Getting the Font Descriptors
//
//   - [NSMutableFontCollection.AddQueryForDescriptors]: Edits the query and exclusion arrays by adding the specified font descriptors.
//   - [NSMutableFontCollection.RemoveQueryForDescriptors]: Edits the query and exclusion arrays by removing the specified font descriptors.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection
type NSMutableFontCollection struct {
	NSFontCollection
}

// NSMutableFontCollectionFromID constructs a [NSMutableFontCollection] from an objc.ID.
//
// A mutable collection of font descriptors taken together as a single object.
func NSMutableFontCollectionFromID(id objc.ID) NSMutableFontCollection {
	return NSMutableFontCollection{NSFontCollection: NSFontCollectionFromID(id)}
}

// NOTE: NSMutableFontCollection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMutableFontCollection] class.
//
// # Getting the Font Descriptors
//
//   - [INSMutableFontCollection.AddQueryForDescriptors]: Edits the query and exclusion arrays by adding the specified font descriptors.
//   - [INSMutableFontCollection.RemoveQueryForDescriptors]: Edits the query and exclusion arrays by removing the specified font descriptors.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection
type INSMutableFontCollection interface {
	INSFontCollection

	// Topic: Getting the Font Descriptors

	// Edits the query and exclusion arrays by adding the specified font descriptors.
	AddQueryForDescriptors(descriptors []NSFontDescriptor)
	// Edits the query and exclusion arrays by removing the specified font descriptors.
	RemoveQueryForDescriptors(descriptors []NSFontDescriptor)
}

// Init initializes the instance.
func (m NSMutableFontCollection) Init() NSMutableFontCollection {
	rv := objc.Send[NSMutableFontCollection](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableFontCollection) Autorelease() NSMutableFontCollection {
	rv := objc.Send[NSMutableFontCollection](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableFontCollection creates a new NSMutableFontCollection instance.
func NewNSMutableFontCollection() NSMutableFontCollection {
	class := getNSMutableFontCollectionClass()
	rv := objc.Send[NSMutableFontCollection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a mutable font collection containing the fonts that match the
// specified font descriptors.
//
// queryDescriptors: One or more query descriptors describing the fonts to include in the
// collection.
//
// # Return Value
//
// A mutable font collection object.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(descriptors:)
func NewMutableFontCollectionWithDescriptors(queryDescriptors []NSFontDescriptor) NSMutableFontCollection {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableFontCollectionClass().class), objc.Sel("fontCollectionWithDescriptors:"), objectivec.IObjectSliceToNSArray(queryDescriptors))
	return NSMutableFontCollectionFromID(rv)
}

// Creates a mutable font collection containing fonts suitable for the
// specified locale.
//
// locale: The locale associated with the fonts you want.
//
// # Return Value
//
// A mutable collection of fonts matching the specified locale.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(locale:)
func NewMutableFontCollectionWithLocale(locale foundation.NSLocale) NSMutableFontCollection {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableFontCollectionClass().class), objc.Sel("fontCollectionWithLocale:"), locale)
	return NSMutableFontCollectionFromID(rv)
}

// Creates a mutable named font collection object.
//
// name: The name to apply to the font collection.
//
// # Return Value
//
// A mutable font collection object.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(name:)
func NewMutableFontCollectionWithName(name NSFontCollectionName) NSMutableFontCollection {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableFontCollectionClass().class), objc.Sel("fontCollectionWithName:"), objc.String(string(name)))
	return NSMutableFontCollectionFromID(rv)
}

// Creates a mutable font collection with the specified name and font
// visibility.
//
// name: The name to apply to the font collection.
//
// visibility: The visibility of the fonts in the collection.
//
// # Return Value
//
// A mutable font collection object.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/init(name:visibility:)
func NewMutableFontCollectionWithNameVisibility(name NSFontCollectionName, visibility NSFontCollectionVisibility) NSMutableFontCollection {
	rv := objc.Send[objc.ID](objc.ID(getNSMutableFontCollectionClass().class), objc.Sel("fontCollectionWithName:visibility:"), objc.String(string(name)), visibility)
	return NSMutableFontCollectionFromID(rv)
}

// Edits the query and exclusion arrays by adding the specified font
// descriptors.
//
// descriptors: The font descriptors to add to the query.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/addQuery(for:)
func (m NSMutableFontCollection) AddQueryForDescriptors(descriptors []NSFontDescriptor) {
	objc.Send[objc.ID](m.ID, objc.Sel("addQueryForDescriptors:"), objectivec.IObjectSliceToNSArray(descriptors))
}

// Edits the query and exclusion arrays by removing the specified font
// descriptors.
//
// descriptors: The font descriptors to add to the query.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableFontCollection/removeQuery(for:)
func (m NSMutableFontCollection) RemoveQueryForDescriptors(descriptors []NSFontDescriptor) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeQueryForDescriptors:"), objectivec.IObjectSliceToNSArray(descriptors))
}
