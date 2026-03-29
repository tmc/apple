// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFontCollection] class.
var (
	_NSFontCollectionClass     NSFontCollectionClass
	_NSFontCollectionClassOnce sync.Once
)

func getNSFontCollectionClass() NSFontCollectionClass {
	_NSFontCollectionClassOnce.Do(func() {
		_NSFontCollectionClass = NSFontCollectionClass{class: objc.GetClass("NSFontCollection")}
	})
	return _NSFontCollectionClass
}

// GetNSFontCollectionClass returns the class object for NSFontCollection.
func GetNSFontCollectionClass() NSFontCollectionClass {
	return getNSFontCollectionClass()
}

type NSFontCollectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFontCollectionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFontCollectionClass) Alloc() NSFontCollection {
	rv := objc.Send[NSFontCollection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A font collection, which is a group of font descriptors taken together as a
// single object.
//
// # Overview
// 
// You can publicize the font collection as a named collection and it is
// presented through the System user interface such as the font panel and Font
// Book. The queries can be modified using the [NSMutableFontCollection]
// subclass.
//
// # Getting the Font Descriptors
//
//   - [NSFontCollection.MatchingDescriptors]: An array of font descriptors matching the logical descriptors.
//   - [NSFontCollection.MatchingDescriptorsForFamily]: Returns an array of font descriptors matching the logical descriptors for the given font family.
//   - [NSFontCollection.MatchingDescriptorsForFamilyOptions]: Returns an array of font descriptors matching the logical descriptors for the given font family and options.
//   - [NSFontCollection.MatchingDescriptorsWithOptions]: Returns an array of font descriptors matching the logical descriptors with the given options.
//   - [NSFontCollection.QueryDescriptors]: An array of font descriptors whose matching results produce the collection’s matching descriptors.
//   - [NSFontCollection.ExclusionDescriptors]: A list of query font descriptors whose matching results are excluded from the list of matching descriptors.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection
type NSFontCollection struct {
	objectivec.Object
}

// NSFontCollectionFromID constructs a [NSFontCollection] from an objc.ID.
//
// A font collection, which is a group of font descriptors taken together as a
// single object.
func NSFontCollectionFromID(id objc.ID) NSFontCollection {
	return NSFontCollection{objectivec.Object{ID: id}}
}
// NOTE: NSFontCollection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFontCollection] class.
//
// # Getting the Font Descriptors
//
//   - [INSFontCollection.MatchingDescriptors]: An array of font descriptors matching the logical descriptors.
//   - [INSFontCollection.MatchingDescriptorsForFamily]: Returns an array of font descriptors matching the logical descriptors for the given font family.
//   - [INSFontCollection.MatchingDescriptorsForFamilyOptions]: Returns an array of font descriptors matching the logical descriptors for the given font family and options.
//   - [INSFontCollection.MatchingDescriptorsWithOptions]: Returns an array of font descriptors matching the logical descriptors with the given options.
//   - [INSFontCollection.QueryDescriptors]: An array of font descriptors whose matching results produce the collection’s matching descriptors.
//   - [INSFontCollection.ExclusionDescriptors]: A list of query font descriptors whose matching results are excluded from the list of matching descriptors.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection
type INSFontCollection interface {
	objectivec.IObject

	// Topic: Getting the Font Descriptors

	// An array of font descriptors matching the logical descriptors.
	MatchingDescriptors() []NSFontDescriptor
	// Returns an array of font descriptors matching the logical descriptors for the given font family.
	MatchingDescriptorsForFamily(family string) []NSFontDescriptor
	// Returns an array of font descriptors matching the logical descriptors for the given font family and options.
	MatchingDescriptorsForFamilyOptions(family string, options foundation.INSDictionary) []NSFontDescriptor
	// Returns an array of font descriptors matching the logical descriptors with the given options.
	MatchingDescriptorsWithOptions(options foundation.INSDictionary) []NSFontDescriptor
	// An array of font descriptors whose matching results produce the collection’s matching descriptors.
	QueryDescriptors() []NSFontDescriptor
	// A list of query font descriptors whose matching results are excluded from the list of matching descriptors.
	ExclusionDescriptors() []NSFontDescriptor

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NSFontCollection) Init() NSFontCollection {
	rv := objc.Send[NSFontCollection](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFontCollection) Autorelease() NSFontCollection {
	rv := objc.Send[NSFontCollection](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFontCollection creates a new NSFontCollection instance.
func NewNSFontCollection() NSFontCollection {
	class := getNSFontCollectionClass()
	rv := objc.Send[NSFontCollection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a font collection matching the given descriptors.
//
// queryDescriptors: The descriptors used to match the returned collection.
//
// # Return Value
// 
// The font collection matching the given descriptors.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(descriptors:)
func NewFontCollectionWithDescriptors(queryDescriptors []NSFontDescriptor) NSFontCollection {
	rv := objc.Send[objc.ID](objc.ID(getNSFontCollectionClass().class), objc.Sel("fontCollectionWithDescriptors:"), objectivec.IObjectSliceToNSArray(queryDescriptors))
	return NSFontCollectionFromID(rv)
}

// Returns a collection of fonts matching the given locale.
//
// locale: The locale to match.
//
// # Return Value
// 
// A collection of fonts matching the given locale.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(locale:)
func NewFontCollectionWithLocale(locale foundation.NSLocale) NSFontCollection {
	rv := objc.Send[objc.ID](objc.ID(getNSFontCollectionClass().class), objc.Sel("fontCollectionWithLocale:"), locale)
	return NSFontCollectionFromID(rv)
}

// Creates a named font collection object.
//
// name: The name of the collection.
//
// # Return Value
// 
// The named font collection.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(name:)
func NewFontCollectionWithName(name NSFontCollectionName) NSFontCollection {
	rv := objc.Send[objc.ID](objc.ID(getNSFontCollectionClass().class), objc.Sel("fontCollectionWithName:"), objc.String(string(name)))
	return NSFontCollectionFromID(rv)
}

// Creates a font collection with the specified name and font visibility.
//
// name: The name of the collection.
//
// visibility: The visibility of the collection.
//
// # Return Value
// 
// The font collection with the specified name and visibility.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/init(name:visibility:)
func NewFontCollectionWithNameVisibility(name NSFontCollectionName, visibility NSFontCollectionVisibility) NSFontCollection {
	rv := objc.Send[objc.ID](objc.ID(getNSFontCollectionClass().class), objc.Sel("fontCollectionWithName:visibility:"), objc.String(string(name)), visibility)
	return NSFontCollectionFromID(rv)
}

// Returns an array of font descriptors matching the logical descriptors for
// the given font family.
//
// family: The font family whose descriptors are matched.
//
// # Return Value
// 
// The [MatchingDescriptors] for the given family.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/matchingDescriptors(forFamily:)
func (f NSFontCollection) MatchingDescriptorsForFamily(family string) []NSFontDescriptor {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("matchingDescriptorsForFamily:"), objc.String(family))
	return objc.ConvertSlice(rv, func(id objc.ID) NSFontDescriptor {
		return NSFontDescriptorFromID(id)
	})
}
// Returns an array of font descriptors matching the logical descriptors for
// the given font family and options.
//
// family: The font family whose descriptors are matched.
//
// options: A dictionary containing any combination of the `Matching Descriptors
// Options` keys or `nil`.
//
// # Return Value
// 
// The [MatchingDescriptors] for the given family and options.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/matchingDescriptors(forFamily:options:)
func (f NSFontCollection) MatchingDescriptorsForFamilyOptions(family string, options foundation.INSDictionary) []NSFontDescriptor {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("matchingDescriptorsForFamily:options:"), objc.String(family), options)
	return objc.ConvertSlice(rv, func(id objc.ID) NSFontDescriptor {
		return NSFontDescriptorFromID(id)
	})
}
// Returns an array of font descriptors matching the logical descriptors with
// the given options.
//
// options: A dictionary containing any combination of the `Matching Descriptors
// Options` keys or `nil`.
//
// # Return Value
// 
// The [MatchingDescriptors] for the given options.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/matchingDescriptors(options:)
func (f NSFontCollection) MatchingDescriptorsWithOptions(options foundation.INSDictionary) []NSFontDescriptor {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("matchingDescriptorsWithOptions:"), options)
	return objc.ConvertSlice(rv, func(id objc.ID) NSFontDescriptor {
		return NSFontDescriptorFromID(id)
	})
}
func (f NSFontCollection) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Renames the font collection with the specified name and visibility to the
// second name specified.
//
// oldName: The collection to rename.
//
// visibility: The visibility of the collection to rename.
//
// newName: The new name to give to the collection.
//
// # Discussion
// 
// Named collections are shown by user interfaces such as the Font panel. When
// you change the collection, you must show it again to see the changes
// reflected on disk or in the Font panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/rename(fromName:visibility:toName:)
func (_NSFontCollectionClass NSFontCollectionClass) RenameFontCollectionWithNameVisibilityToNameError(oldName NSFontCollectionName, visibility NSFontCollectionVisibility, newName NSFontCollectionName) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_NSFontCollectionClass.class), objc.Sel("renameFontCollectionWithName:visibility:toName:error:"), objc.String(string(oldName)), visibility, objc.String(string(newName)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("renameFontCollectionWithName:visibility:toName:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Make the given font collection visible by giving it a name.
//
// collection: The font collection to make visible.
//
// name: The name to associate with the collection.
//
// visibility: The visibility of the collection to show.
//
// # Discussion
// 
// Named collections are shown by user interfaces such as the Font panel. When
// you change the collection, you must show it again to see the changes
// reflected on disk or in the Font panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/show(_:withName:visibility:)
func (_NSFontCollectionClass NSFontCollectionClass) ShowFontCollectionWithNameVisibilityError(collection INSFontCollection, name NSFontCollectionName, visibility NSFontCollectionVisibility) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_NSFontCollectionClass.class), objc.Sel("showFontCollection:withName:visibility:error:"), collection, objc.String(string(name)), visibility, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("showFontCollection:withName:visibility:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Remove from view the named font collection with the specified visibility.
//
// name: The name of the collection.
//
// visibility: The visibility of the collection.
//
// # Discussion
// 
// For a persistent font collection, this method deletes the named font
// collection from disk.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/hide(withName:visibility:)
func (_NSFontCollectionClass NSFontCollectionClass) HideFontCollectionWithNameVisibilityError(name NSFontCollectionName, visibility NSFontCollectionVisibility) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_NSFontCollectionClass.class), objc.Sel("hideFontCollectionWithName:visibility:error:"), objc.String(string(name)), visibility, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("hideFontCollectionWithName:visibility:error: returned NO with nil NSError")
	}
	return rv, nil

}

// An array of font descriptors matching the logical descriptors.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/matchingDescriptors
func (f NSFontCollection) MatchingDescriptors() []NSFontDescriptor {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("matchingDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSFontDescriptor {
		return NSFontDescriptorFromID(id)
	})
}
// An array of font descriptors whose matching results produce the
// collection’s matching descriptors.
//
// # Discussion
// 
// The font descriptors matching [ExclusionDescriptors] are removed from
// [MatchingDescriptors]
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/queryDescriptors
func (f NSFontCollection) QueryDescriptors() []NSFontDescriptor {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("queryDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSFontDescriptor {
		return NSFontDescriptorFromID(id)
	})
}
// A list of query font descriptors whose matching results are excluded from
// the list of matching descriptors.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/exclusionDescriptors
func (f NSFontCollection) ExclusionDescriptors() []NSFontDescriptor {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("exclusionDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSFontDescriptor {
		return NSFontDescriptorFromID(id)
	})
}

// The font collection that matches all registered fonts.
//
// # Return Value
// 
// The collection of all fonts available to the current application.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/withAllAvailableDescriptors
func (_NSFontCollectionClass NSFontCollectionClass) FontCollectionWithAllAvailableDescriptors() NSFontCollection {
	rv := objc.Send[objc.ID](objc.ID(_NSFontCollectionClass.class), objc.Sel("fontCollectionWithAllAvailableDescriptors"))
	return NSFontCollectionFromID(objc.ID(rv))
}
// Returns all named collections visible to this process.
//
// # Return Value
// 
// [NSString] objects containing the names of all the named collections.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/allFontCollectionNames
func (_NSFontCollectionClass NSFontCollectionClass) AllFontCollectionNames() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSFontCollectionClass.class), objc.Sel("allFontCollectionNames"))
	return objc.ConvertSliceToStrings(rv)
}

