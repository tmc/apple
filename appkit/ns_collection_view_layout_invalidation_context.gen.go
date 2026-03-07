// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionViewLayoutInvalidationContext] class.
var (
	_NSCollectionViewLayoutInvalidationContextClass     NSCollectionViewLayoutInvalidationContextClass
	_NSCollectionViewLayoutInvalidationContextClassOnce sync.Once
)

func getNSCollectionViewLayoutInvalidationContextClass() NSCollectionViewLayoutInvalidationContextClass {
	_NSCollectionViewLayoutInvalidationContextClassOnce.Do(func() {
		_NSCollectionViewLayoutInvalidationContextClass = NSCollectionViewLayoutInvalidationContextClass{class: objc.GetClass("NSCollectionViewLayoutInvalidationContext")}
	})
	return _NSCollectionViewLayoutInvalidationContextClass
}

// GetNSCollectionViewLayoutInvalidationContextClass returns the class object for NSCollectionViewLayoutInvalidationContext.
func GetNSCollectionViewLayoutInvalidationContextClass() NSCollectionViewLayoutInvalidationContextClass {
	return getNSCollectionViewLayoutInvalidationContextClass()
}

type NSCollectionViewLayoutInvalidationContextClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewLayoutInvalidationContextClass) Alloc() NSCollectionViewLayoutInvalidationContext {
	rv := objc.Send[NSCollectionViewLayoutInvalidationContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that identifies the portions of your layout that need to be
// updated.
//
// # Overview
// 
// Invalidation contexts are a way to improve the efficiency of layout
// operations and must be supported explicitly by the layout object. Instead
// of invalidating the entire layout, you can create an invalidation layout
// object that specifies only the portions of the layout that changed. You
// then pass that invalidation context to the [InvalidateLayoutWithContext]
// method of the layout object.
// 
// Typically, you ask the layout object to create an invalidation context for
// you. The [NSCollectionViewLayout] class defines methods for creating a
// supported invalidation context. If you define a custom layout, you can
// define additional methods for creating invalidation contexts with custom
// information. Layout objects may also create invalidation contexts in
// response to specific changes. For example, layout objects automatically
// create invalidation contexts when you change the collection view’s data
// source, when you insert or delete items, and when you reload the collection
// view’s data.
// 
// # Subclassing Notes
// 
// If you define a custom layout object, you can also subclass
// [NSCollectionViewLayoutInvalidationContext] and add properties that are
// specific to your layout object. Creating a custom invalidation context is
// not required and should only be done when your layout object has additional
// ways to optimize the layout process.
// 
// Fore more information about how to support custom invalidation contexts in
// your layout objects, see [NSCollectionViewLayout].
//
// # Invalidating the Collection View Data
//
//   - [NSCollectionViewLayoutInvalidationContext.InvalidateEverything]: A Boolean that indicates whether all layout data should be marked as invalid.
//   - [NSCollectionViewLayoutInvalidationContext.InvalidateDataSourceCounts]: A Boolean that indicates whether the layout object should ask for new section and item counts.
//
// # Invalidating the Content Area
//
//   - [NSCollectionViewLayoutInvalidationContext.ContentOffsetAdjustment]: The delta value to add to the collection view’s content offset.
//   - [NSCollectionViewLayoutInvalidationContext.SetContentOffsetAdjustment]
//   - [NSCollectionViewLayoutInvalidationContext.ContentSizeAdjustment]: The delta value to add to the collection view’s content size.
//   - [NSCollectionViewLayoutInvalidationContext.SetContentSizeAdjustment]
//
// # Invalidating Specific Items
//
//   - [NSCollectionViewLayoutInvalidationContext.InvalidateItemsAtIndexPaths]: Marks the specified items as invalid so that their layout information can be updated.
//   - [NSCollectionViewLayoutInvalidationContext.InvalidateSupplementaryElementsOfKindAtIndexPaths]: Marks the specified supplementary views as invalid so that their layout information can be updated.
//   - [NSCollectionViewLayoutInvalidationContext.InvalidateDecorationElementsOfKindAtIndexPaths]: Marks the specified decoration views as invalid so that their layout information can be updated.
//   - [NSCollectionViewLayoutInvalidationContext.InvalidatedItemIndexPaths]: The set of items whose layout attributes are invalid.
//   - [NSCollectionViewLayoutInvalidationContext.InvalidatedSupplementaryIndexPaths]: A dictionary containing the supplementary views whose layout attributes are invalid.
//   - [NSCollectionViewLayoutInvalidationContext.InvalidatedDecorationIndexPaths]: A dictionary containing the decoration views whose layout attributes are invalid.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext
type NSCollectionViewLayoutInvalidationContext struct {
	objectivec.Object
}

// NSCollectionViewLayoutInvalidationContextFromID constructs a [NSCollectionViewLayoutInvalidationContext] from an objc.ID.
//
// An object that identifies the portions of your layout that need to be
// updated.
func NSCollectionViewLayoutInvalidationContextFromID(id objc.ID) NSCollectionViewLayoutInvalidationContext {
	return NSCollectionViewLayoutInvalidationContext{objectivec.Object{id}}
}
// NOTE: NSCollectionViewLayoutInvalidationContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCollectionViewLayoutInvalidationContext] class.
//
// # Invalidating the Collection View Data
//
//   - [INSCollectionViewLayoutInvalidationContext.InvalidateEverything]: A Boolean that indicates whether all layout data should be marked as invalid.
//   - [INSCollectionViewLayoutInvalidationContext.InvalidateDataSourceCounts]: A Boolean that indicates whether the layout object should ask for new section and item counts.
//
// # Invalidating the Content Area
//
//   - [INSCollectionViewLayoutInvalidationContext.ContentOffsetAdjustment]: The delta value to add to the collection view’s content offset.
//   - [INSCollectionViewLayoutInvalidationContext.SetContentOffsetAdjustment]
//   - [INSCollectionViewLayoutInvalidationContext.ContentSizeAdjustment]: The delta value to add to the collection view’s content size.
//   - [INSCollectionViewLayoutInvalidationContext.SetContentSizeAdjustment]
//
// # Invalidating Specific Items
//
//   - [INSCollectionViewLayoutInvalidationContext.InvalidateItemsAtIndexPaths]: Marks the specified items as invalid so that their layout information can be updated.
//   - [INSCollectionViewLayoutInvalidationContext.InvalidateSupplementaryElementsOfKindAtIndexPaths]: Marks the specified supplementary views as invalid so that their layout information can be updated.
//   - [INSCollectionViewLayoutInvalidationContext.InvalidateDecorationElementsOfKindAtIndexPaths]: Marks the specified decoration views as invalid so that their layout information can be updated.
//   - [INSCollectionViewLayoutInvalidationContext.InvalidatedItemIndexPaths]: The set of items whose layout attributes are invalid.
//   - [INSCollectionViewLayoutInvalidationContext.InvalidatedSupplementaryIndexPaths]: A dictionary containing the supplementary views whose layout attributes are invalid.
//   - [INSCollectionViewLayoutInvalidationContext.InvalidatedDecorationIndexPaths]: A dictionary containing the decoration views whose layout attributes are invalid.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext
type INSCollectionViewLayoutInvalidationContext interface {
	objectivec.IObject

	// Topic: Invalidating the Collection View Data

	// A Boolean that indicates whether all layout data should be marked as invalid.
	InvalidateEverything() bool
	// A Boolean that indicates whether the layout object should ask for new section and item counts.
	InvalidateDataSourceCounts() bool

	// Topic: Invalidating the Content Area

	// The delta value to add to the collection view’s content offset.
	ContentOffsetAdjustment() corefoundation.CGPoint
	SetContentOffsetAdjustment(value corefoundation.CGPoint)
	// The delta value to add to the collection view’s content size.
	ContentSizeAdjustment() corefoundation.CGSize
	SetContentSizeAdjustment(value corefoundation.CGSize)

	// Topic: Invalidating Specific Items

	// Marks the specified items as invalid so that their layout information can be updated.
	InvalidateItemsAtIndexPaths(indexPaths foundation.INSSet)
	// Marks the specified supplementary views as invalid so that their layout information can be updated.
	InvalidateSupplementaryElementsOfKindAtIndexPaths(elementKind NSCollectionViewSupplementaryElementKind, indexPaths foundation.INSSet)
	// Marks the specified decoration views as invalid so that their layout information can be updated.
	InvalidateDecorationElementsOfKindAtIndexPaths(elementKind NSCollectionViewDecorationElementKind, indexPaths foundation.INSSet)
	// The set of items whose layout attributes are invalid.
	InvalidatedItemIndexPaths() foundation.INSSet
	// A dictionary containing the supplementary views whose layout attributes are invalid.
	InvalidatedSupplementaryIndexPaths() foundation.INSDictionary
	// A dictionary containing the decoration views whose layout attributes are invalid.
	InvalidatedDecorationIndexPaths() foundation.INSDictionary
}





// Init initializes the instance.
func (c NSCollectionViewLayoutInvalidationContext) Init() NSCollectionViewLayoutInvalidationContext {
	rv := objc.Send[NSCollectionViewLayoutInvalidationContext](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewLayoutInvalidationContext) Autorelease() NSCollectionViewLayoutInvalidationContext {
	rv := objc.Send[NSCollectionViewLayoutInvalidationContext](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewLayoutInvalidationContext creates a new NSCollectionViewLayoutInvalidationContext instance.
func NewNSCollectionViewLayoutInvalidationContext() NSCollectionViewLayoutInvalidationContext {
	class := getNSCollectionViewLayoutInvalidationContextClass()
	rv := objc.Send[NSCollectionViewLayoutInvalidationContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Marks the specified items as invalid so that their layout information can
// be updated.
//
// indexPaths: A set of [NSIndexPath] objects. Each index path represents an item whose
// layout needs to be recomputed.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// Call this method when you want the layout object to recompute attributes
// for a specific set of items. The items you provide are added to the
// [InvalidatedItemIndexPaths] property. You can call this method more than
// once to create a merged set of items.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidateItems(at:)
func (c NSCollectionViewLayoutInvalidationContext) InvalidateItemsAtIndexPaths(indexPaths foundation.INSSet) {
	objc.Send[objc.ID](c.ID, objc.Sel("invalidateItemsAtIndexPaths:"), indexPaths)
}

// Marks the specified supplementary views as invalid so that their layout
// information can be updated.
//
// elementKind: A string that identifies the type of the supplementary views. This
// parameter must not be `nil` or an empty string.
//
// indexPaths: A set of [NSIndexPath] objects. Each index path contains the section in
// which the supplementary view appears.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// Call this method when you want the layout object to recompute attributes
// for one or more supplementary views. All of the views must be of the type
// specified by the `elementKind` parameter. The method adds the views you
// specify to the [InvalidatedSupplementaryIndexPaths] property. You can call
// this method more than once for the specified `elementKind` value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidateSupplementaryElements(ofKind:at:)
func (c NSCollectionViewLayoutInvalidationContext) InvalidateSupplementaryElementsOfKindAtIndexPaths(elementKind NSCollectionViewSupplementaryElementKind, indexPaths foundation.INSSet) {
	objc.Send[objc.ID](c.ID, objc.Sel("invalidateSupplementaryElementsOfKind:atIndexPaths:"), objc.String(string(elementKind)), indexPaths)
}

// Marks the specified decoration views as invalid so that their layout
// information can be updated.
//
// elementKind: A string that identifies the type of the decoration views. This parameter
// must not be `nil` or an empty string.
//
// indexPaths: A set of [NSIndexPath] objects. Each index path contains the section in
// which the decoration view appears.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// Call this method when you want the layout object to recompute attributes
// for one or more decoration views. All of the views must be of the type
// specified by the `elementKind` parameter. The method adds the views you
// specify to the [InvalidatedDecorationIndexPaths] property. You can call
// this method more than once for the specified `elementKind` value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidateDecorationElements(ofKind:at:)
func (c NSCollectionViewLayoutInvalidationContext) InvalidateDecorationElementsOfKindAtIndexPaths(elementKind NSCollectionViewDecorationElementKind, indexPaths foundation.INSSet) {
	objc.Send[objc.ID](c.ID, objc.Sel("invalidateDecorationElementsOfKind:atIndexPaths:"), objc.String(string(elementKind)), indexPaths)
}












// A Boolean that indicates whether all layout data should be marked as
// invalid.
//
// # Discussion
// 
// The collection view sets this property in response to specific layout
// invalidation scenarios. For example, the collection view sets the property
// to [true] when you change the current layout object, change the data source
// of the collection view, or call the [ReloadData] method and subsequently
// request a layout invalidation context.
// 
// When this property is set to [true], the layout object must throw away all
// previous layout information and recompute it.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidateEverything
func (c NSCollectionViewLayoutInvalidationContext) InvalidateEverything() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("invalidateEverything"))
	return rv
}



// A Boolean that indicates whether the layout object should ask for new
// section and item counts.
//
// # Discussion
// 
// The collection view sets this property in response to specific layout
// invalidation scenarios. For example, the collection view sets the property
// to [true] when you insert or delete items or call the collection view’s
// [ReloadData] method.
// 
// When this property is set to [true], the layout object must query the data
// source for the new number of sections and items. IT should also update its
// layout based on the updated number of sections and items.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidateDataSourceCounts
func (c NSCollectionViewLayoutInvalidationContext) InvalidateDataSourceCounts() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("invalidateDataSourceCounts"))
	return rv
}



// The delta value to add to the collection view’s content offset.
//
// # Discussion
// 
// The content offset adjustment shifts the position of content inside the
// collection view by the specified amount. You use this value to make tweaks
// based on how you want to present your content. For example, you might use
// it to ensure that the first line of items is always lined up at the same
// position in the collection view’s visible rectangle. When making
// adjustments, you can specify both positive and negative values.
// 
// The default value of this property is [NSZeroSize].
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/contentOffsetAdjustment
func (c NSCollectionViewLayoutInvalidationContext) ContentOffsetAdjustment() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("contentOffsetAdjustment"))
	return corefoundation.CGPoint(rv)
}
func (c NSCollectionViewLayoutInvalidationContext) SetContentOffsetAdjustment(value corefoundation.CGPoint) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentOffsetAdjustment:"), value)
}



// The delta value to add to the collection view’s content size.
//
// # Discussion
// 
// Use this property to update the size of the collection view’s content
// area, as computed by the associated layout object. The default value of
// this property is [NSZeroSize]. Changing the value causes the collection
// view to add the specified height and width values to its content size.
// Thus, positive values grow the content area and negative values shrink it.
// You might add space around the content area to provide a visual buffer for
// your collection view content.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/contentSizeAdjustment
func (c NSCollectionViewLayoutInvalidationContext) ContentSizeAdjustment() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("contentSizeAdjustment"))
	return corefoundation.CGSize(rv)
}
func (c NSCollectionViewLayoutInvalidationContext) SetContentSizeAdjustment(value corefoundation.CGSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentSizeAdjustment:"), value)
}



// The set of items whose layout attributes are invalid.
//
// # Discussion
// 
// The set contains zero or more [NSIndexPath] objects, each of which
// identifies an invalid item.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidatedItemIndexPaths
func (c NSCollectionViewLayoutInvalidationContext) InvalidatedItemIndexPaths() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("invalidatedItemIndexPaths"))
	return foundation.NSSetFromID(objc.ID(rv))
}



// A dictionary containing the supplementary views whose layout attributes are
// invalid.
//
// # Discussion
// 
// The keys in this dictionary are the element kind strings of the
// supplementary views. The value for each key is an [NSSet] object containing
// one or more [NSIndexPath] objects, each of which identifies the section
// containing the supplementary view.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
// [NSSet]: https://developer.apple.com/documentation/Foundation/NSSet
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidatedSupplementaryIndexPaths
func (c NSCollectionViewLayoutInvalidationContext) InvalidatedSupplementaryIndexPaths() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("invalidatedSupplementaryIndexPaths"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}



// A dictionary containing the decoration views whose layout attributes are
// invalid.
//
// # Discussion
// 
// The keys in this dictionary are the element kind strings of the decoration
// views. The value for each key is an [NSSet] object containing one or more
// [NSIndexPath] objects, each of which identifies the section containing the
// decoration view.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
// [NSSet]: https://developer.apple.com/documentation/Foundation/NSSet
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutInvalidationContext/invalidatedDecorationIndexPaths
func (c NSCollectionViewLayoutInvalidationContext) InvalidatedDecorationIndexPaths() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("invalidatedDecorationIndexPaths"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
























