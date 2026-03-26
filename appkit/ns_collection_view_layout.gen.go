// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionViewLayout] class.
var (
	_NSCollectionViewLayoutClass     NSCollectionViewLayoutClass
	_NSCollectionViewLayoutClassOnce sync.Once
)

func getNSCollectionViewLayoutClass() NSCollectionViewLayoutClass {
	_NSCollectionViewLayoutClassOnce.Do(func() {
		_NSCollectionViewLayoutClass = NSCollectionViewLayoutClass{class: objc.GetClass("NSCollectionViewLayout")}
	})
	return _NSCollectionViewLayoutClass
}

// GetNSCollectionViewLayoutClass returns the class object for NSCollectionViewLayout.
func GetNSCollectionViewLayoutClass() NSCollectionViewLayoutClass {
	return getNSCollectionViewLayoutClass()
}

type NSCollectionViewLayoutClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewLayoutClass) Alloc() NSCollectionViewLayout {
	rv := objc.Send[NSCollectionViewLayout](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class that you subclass and use to generate layout
// information for a collection view.
//
// # Overview
// 
// The job of a layout object is to perform the calculations needed to
// determine the placement and appearance of items, supplementary views, and
// other content in the collection view. The layout object does not apply the
// layout attributes it generates to the views in your interface. Instead, it
// passes those layout attributes to the collection view, which then creates
// the needed views and applies the layout attributes to them.
// 
// You do not create instances of this class directly. Instead, you create
// instances of one of its subclasses and associate that object with your
// collection view either programmatically (using the [CollectionViewLayout]
// property) or at design time in Interface Builder. Changing the layout
// object of a collection view forces an immediate update of the layout
// information.
// 
// Collection views support many different types of elements, most of which
// are visual and all of which require layout attributes:
// 
// - are the main elements managed by the layout. Each item represents a
// single piece of data in the collection view. A collection view can have a
// single group of items or it can divide the items into multiple sections. -
// are optional views associated with a specific section. The layout object
// defines the placement and use of supplementary views. For example, grid and
// flow layouts use supplementary views to implement headers and footers for
// each section. Supplementary views cannot be selected by the user. - are
// visual adornments used to implement themes or to present visual content
// that is unrelated to the data being managed by the collection view.
// Decoration views are optional and the layout object defines their use and
// placement. - supply a drop target for dragged content. Gaps do not have a
// direct visual representation, but they do have layout attributes, which the
// collection view uses for hit testing. The layout object provides attributes
// for inter-item gaps only when asked to do so.
// 
// Each concrete layout object defines a specific organization for the
// contained elements and provides the appropriate layout attributes. The
// placement and appearance of items is determined entirely by the layout
// object. The [NSCollectionViewFlowLayout] and [NSCollectionViewGridLayout]
// subclasses define variants of a grid-based layout, but you can create
// custom layouts that arrange elements in different ways. For example, you
// might define a layout class that arranges items in a circle or define a
// class that groups items into stacks that resemble a pile of photos on a
// table.
// 
// # Subclassing Notes
// 
// Layout objects provide information about the position and visual state of
// elements in the collection view. When defining an [NSCollectionViewLayout]
// subclass, first have an idea for how items in your layout will appear
// onscreen. Layouts should provide a well-defined organization for their
// elements and should not morph between different organizations. If you want
// your collection view to change how it organizes elements, define separate
// layout objects for each organization and switch between them at runtime.
// 
// Remember that a layout object does not directly create the views for which
// it provides layout information. Creation of the views is handled by the
// collection view, which creates them with the help of its data source
// object. The layout object only assists the collection view by providing
// information about the position and appearance of the views that the
// collection view creates.
// 
// Your [NSCollectionViewLayout] subclass should be mostly self contained and
// should not rely on the collection view. Layout objects can retrieve the
// number of sections and items from the collection view, but you should not
// design layout objects that require knowledge of the data being displayed.
// The separation of the layout information from the displayed data offers
// flexibility and makes it easier to reuse layout objects.
// 
// # Methods to Override
// 
// When defining a custom layout class, always override the following methods
// and properties:
// 
// - [NSCollectionViewLayout.PrepareLayout] - [NSCollectionViewLayout.CollectionViewContentSize] -
// [NSCollectionViewLayout.LayoutAttributesForElementsInRect] - [NSCollectionViewLayout.LayoutAttributesForItemAtIndexPath]
// - [NSCollectionViewLayout.LayoutAttributesForSupplementaryViewOfKindAtIndexPath] (if your layout
// supports supplementary views) -
// [NSCollectionViewLayout.LayoutAttributesForDecorationViewOfKindAtIndexPath] (if your layout
// supports decoration views) - [NSCollectionViewLayout.ShouldInvalidateLayoutForBoundsChange]
// 
// These methods provide the fundamental layout information that the
// collection view needs to configure the items and other views that it
// displays. Layout objects need to implement only the methods associated with
// the types of elements they support. So if your layout object does not
// include decoration views, you do not need to implement the
// [NSCollectionViewLayout.LayoutAttributesForDecorationViewOfKindAtIndexPath] method or any other
// methods relating to decoration views.
// 
// In addition to the preceding methods, it is recommended that you implement
// several other methods in your custom layout objects. The insertion or
// deletion of items usually involves animations to move those items into
// position. Providing initial or final layout attributes makes your layout
// more engaging by improving the animations associated with the insertion or
// deletion of items. Other methods provide additional support for
// layout-related behaviors.
// 
// - [NSCollectionViewLayout.InitialLayoutAttributesForAppearingItemAtIndexPath] -
// [NSCollectionViewLayout.InitialLayoutAttributesForAppearingSupplementaryElementOfKindAtIndexPath]
// - [NSCollectionViewLayout.InitialLayoutAttributesForAppearingDecorationElementOfKindAtIndexPath] -
// [NSCollectionViewLayout.FinalLayoutAttributesForDisappearingItemAtIndexPath] -
// [NSCollectionViewLayout.FinalLayoutAttributesForDisappearingSupplementaryElementOfKindAtIndexPath]
// - [NSCollectionViewLayout.FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath]
// - [NSCollectionViewLayout.LayoutAttributesForDropTargetAtPoint] (if your layout supports dropping
// content between items - [NSCollectionViewLayout.LayoutAttributesForInterItemGapBeforeIndexPath]
// (if your layout supports dropping content in gaps between elements)
// 
// # Understanding the Layout Process
// 
// Although the collection view drives the layout process, layout objects do
// all the work of calculating the layout information used during that
// process. The layout object acts as a semi-passive utility object, computing
// the location of elements and returning layout attributes only when asked to
// do so. Layout occurs when the collection view is first displayed and
// reoccurs whenever all or part of the layout is marked as invalid.
// 
// During the layout process, the collection view calls several methods of
// your layout object to gather information. In particular, it calls three
// very important methods, whose implementations drive the core layout
// behavior.
// 
// - Use the [NSCollectionViewLayout.PrepareLayout] method to perform your initial layout
// calculations. These calculations provide the basis for everything the
// layout object does later. - Use the [NSCollectionViewLayout.CollectionViewContentSize] method to
// return the smallest rectangle that completely encloses all of the elements
// in the collection view. Use the calculations from your [NSCollectionViewLayout.PrepareLayout]
// method to specify this rectangle. - Use the
// [NSCollectionViewLayout.LayoutAttributesForElementsInRect] method to return the layout attributes
// for all elements in the specified rectangle. The collection view typically
// requests only the subset of visible elements, but may include elements that
// are just offscreen.
// 
// The [NSCollectionViewLayout.PrepareLayout] method is your chance to perform the main calculations
// associated with the layout process. Use this method to generate an initial
// list of layout attributes for your content. For example, use this method to
// calculate the frame rectangles of all elements in the collection view.
// Performing all of these calculations up front and caching the resulting
// data is often simpler than trying to compute attributes for individual
// items later.
// 
// In addition to the [NSCollectionViewLayout.LayoutAttributesForElementsInRect] method, the
// collection view may call other methods to retrieve layout attributes for
// specific items. By performing your calculations in advance, your
// implementations of those methods should be able to return cached
// information without having to recompute that information first. The only
// time your layout object needs to recompute its layout information is when
// your app invalidates the layout. For example, you might invalidate the
// layout when the user inserts or deletes items.
// 
// # Adding Decoration Views to Your Layout
// 
// Layouts can use decoration views to create a specific visual appearance for
// the collection view’s content. Decoration views are fully configured
// visual elements that the layout object provides. When defining your own
// custom layout objects, you might use decoration views to provide a dynamic
// background or add visual styling in and around the items of the collection
// view. The layout object defines what decoration views are needed and
// registers an appropriate class or nib file for each view using the
// [NSCollectionViewLayout.RegisterClassForDecorationViewOfKind] or
// [NSCollectionViewLayout.RegisterNibForDecorationViewOfKind] method.
// 
// To create a decoration view, return an appropriate layout attributes object
// from your layout object’s [NSCollectionViewLayout.LayoutAttributesForElementsInRect] method.
// When the collection view receives attributes for a decoration view, it
// creates that view using the class or nib file that your layout object
// registered. Because the collection view creates them, decoration views must
// be fully configured at registration time. You cannot add content to a
// decoration view or change its configuration later.
// 
// # Optimizing Layout Performance Using Invalidation Contexts
// 
// When designing custom layouts, you can improve performance by updating only
// the portions of your layout that actually changed. The [NSCollectionViewLayout.InvalidateLayout]
// method forces the collection view to throw away all of its layout
// information and recompute it, which is inefficient if most of the layout
// has not changed. A better way to update your layout is using the
// [NSCollectionViewLayout.InvalidateLayoutWithContext] method, which uses the provided context
// object to invalidate only the parts of the layout that changed.
// 
// Support for invalidation contexts must be built into the implementation of
// your layout object. At a minimum, you must override the
// [NSCollectionViewLayout.InvalidateLayoutWithContext] method and use it to mark the parts of your
// layout that changed. You might do this by setting flags or throwing away
// cached layout attributes for the changed elements. Your
// [NSCollectionViewLayout.InvalidateLayoutWithContext] method should also call `super` so that the
// collection view can initiate the layout update process at a future time.
// 
// If your layout object supports more fine-grained invalidation than the
// [NSCollectionViewLayoutInvalidationContext] class provides, you can
// subclass and add your invalidation information there. If you define a
// custom context class, override the [NSCollectionViewLayout.InvalidationContextClass] property in
// your layout object so that the collection view knows which class to
// instantiate. Similarly, other parts of your app should create instances of
// your custom context class and use them to invalidate the layout.
//
// # Getting the Collection View
//
//   - [NSCollectionViewLayout.CollectionView]: The collection view object currently using this layout.
//
// # Providing Layout Information
//
//   - [NSCollectionViewLayout.PrepareLayout]: Prepares the layout object to begin laying out content.
//   - [NSCollectionViewLayout.CollectionViewContentSize]: The width and height of the collection view’s contents.
//   - [NSCollectionViewLayout.LayoutAttributesForElementsInRect]: Returns the layout attribute objects for all items and views in the specified rectangle.
//   - [NSCollectionViewLayout.LayoutAttributesForItemAtIndexPath]: Returns the layout attributes for the item at the specified index path.
//   - [NSCollectionViewLayout.LayoutAttributesForSupplementaryViewOfKindAtIndexPath]: Returns the layout attributes of the supplementary view at the specified location in your layout.
//   - [NSCollectionViewLayout.LayoutAttributesForDecorationViewOfKindAtIndexPath]: Returns the layout attributes of the decoration view at the specified location in your layout.
//   - [NSCollectionViewLayout.LayoutAttributesForDropTargetAtPoint]: Returns layout attributes for the drop target at the specified point.
//   - [NSCollectionViewLayout.LayoutAttributesForInterItemGapBeforeIndexPath]: Returns layout attributes for the inter-item gap at the specified location in your layout.
//   - [NSCollectionViewLayout.TargetContentOffsetForProposedContentOffset]: Returns the offset value to use after an animated layout update or change.
//   - [NSCollectionViewLayout.TargetContentOffsetForProposedContentOffsetWithScrollingVelocity]: Returns the offset value to use for the collection view’s content at the end of scrolling.
//
// # Responding to Collection View Updates
//
//   - [NSCollectionViewLayout.PrepareForCollectionViewUpdates]: Performs needed tasks before items are inserted, deleted, or moved within the collection view.
//   - [NSCollectionViewLayout.FinalizeCollectionViewUpdates]: Performs needed steps after items are inserted, deleted, or moved within a collection view.
//   - [NSCollectionViewLayout.IndexPathsToInsertForSupplementaryViewOfKind]: Returns the index paths for any supplementary views that the layout object wants to add to the collection view.
//   - [NSCollectionViewLayout.IndexPathsToInsertForDecorationViewOfKind]: Returns the index paths for any decoration views that the layout object wants to add to the collection view.
//   - [NSCollectionViewLayout.InitialLayoutAttributesForAppearingItemAtIndexPath]: Returns the starting layout information for an item being inserted into the collection view.
//   - [NSCollectionViewLayout.InitialLayoutAttributesForAppearingSupplementaryElementOfKindAtIndexPath]: Returns the starting layout information for a supplementary view being added to the collection view.
//   - [NSCollectionViewLayout.InitialLayoutAttributesForAppearingDecorationElementOfKindAtIndexPath]: Returns the starting layout information for a decoration view being added to the collection view.
//   - [NSCollectionViewLayout.IndexPathsToDeleteForSupplementaryViewOfKind]: Returns the index paths for any supplementary views that the layout object wants to remove from the collection view.
//   - [NSCollectionViewLayout.IndexPathsToDeleteForDecorationViewOfKind]: Returns index paths for any decoration views that the layout object wants to remove from the collection view.
//   - [NSCollectionViewLayout.FinalLayoutAttributesForDisappearingItemAtIndexPath]: Returns the ending layout information for an item being removed from the collection view.
//   - [NSCollectionViewLayout.FinalLayoutAttributesForDisappearingSupplementaryElementOfKindAtIndexPath]: Returns the ending layout information for a supplementary view being removed from the collection view.
//   - [NSCollectionViewLayout.FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath]: Returns the ending layout information for a decoration view being removed from the collection view.
//
// # Invalidating the Layout
//
//   - [NSCollectionViewLayout.InvalidateLayout]: Invalidates all layout information and triggers a layout update.
//   - [NSCollectionViewLayout.InvalidateLayoutWithContext]: Invalidates specific parts of the layout using the specified context object.
//   - [NSCollectionViewLayout.ShouldInvalidateLayoutForBoundsChange]: Returns a Boolean indicating whether a bounds change triggers a layout update.
//   - [NSCollectionViewLayout.ShouldInvalidateLayoutForPreferredLayoutAttributesWithOriginalAttributes]: Returns a Boolean indicating whether changes to a cell’s layout attributes trigger a larger layout update.
//   - [NSCollectionViewLayout.InvalidationContextForBoundsChange]: Returns an invalidation context object that defines the portions of the layout that need to be updated.
//   - [NSCollectionViewLayout.InvalidationContextForPreferredLayoutAttributesWithOriginalAttributes]: Returns an invalidation context object that defines the portions of the layout that need to be updated.
//
// # Coordinating Animated Changes
//
//   - [NSCollectionViewLayout.PrepareForAnimatedBoundsChange]: Prepares the layout object for animated changes to the collection view’s bounds or for the insertion or deletion of items.
//   - [NSCollectionViewLayout.FinalizeAnimatedBoundsChange]: Cleans up after any animated changes to the collection view’s bounds or after the insertion or deletion of items.
//
// # Registering Decoration Views
//
//   - [NSCollectionViewLayout.RegisterClassForDecorationViewOfKind]: Registers a class to use when creating the layout’s decoration views.
//   - [NSCollectionViewLayout.RegisterNibForDecorationViewOfKind]: Registers a nib file to use when creating the layout’s decoration views.
//
// # Transitioning Between Layouts
//
//   - [NSCollectionViewLayout.PrepareForTransitionFromLayout]: Prepares the layout object to be installed in the collection view.
//   - [NSCollectionViewLayout.PrepareForTransitionToLayout]: Prepares the layout object to be uninstalled from the collection view.
//   - [NSCollectionViewLayout.FinalizeLayoutTransition]: Performs any final steps related to a layout transition before the transition animations actually occur.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout
type NSCollectionViewLayout struct {
	objectivec.Object
}

// NSCollectionViewLayoutFromID constructs a [NSCollectionViewLayout] from an objc.ID.
//
// An abstract base class that you subclass and use to generate layout
// information for a collection view.
func NSCollectionViewLayoutFromID(id objc.ID) NSCollectionViewLayout {
	return NSCollectionViewLayout{objectivec.Object{ID: id}}
}
// NOTE: NSCollectionViewLayout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionViewLayout] class.
//
// # Getting the Collection View
//
//   - [INSCollectionViewLayout.CollectionView]: The collection view object currently using this layout.
//
// # Providing Layout Information
//
//   - [INSCollectionViewLayout.PrepareLayout]: Prepares the layout object to begin laying out content.
//   - [INSCollectionViewLayout.CollectionViewContentSize]: The width and height of the collection view’s contents.
//   - [INSCollectionViewLayout.LayoutAttributesForElementsInRect]: Returns the layout attribute objects for all items and views in the specified rectangle.
//   - [INSCollectionViewLayout.LayoutAttributesForItemAtIndexPath]: Returns the layout attributes for the item at the specified index path.
//   - [INSCollectionViewLayout.LayoutAttributesForSupplementaryViewOfKindAtIndexPath]: Returns the layout attributes of the supplementary view at the specified location in your layout.
//   - [INSCollectionViewLayout.LayoutAttributesForDecorationViewOfKindAtIndexPath]: Returns the layout attributes of the decoration view at the specified location in your layout.
//   - [INSCollectionViewLayout.LayoutAttributesForDropTargetAtPoint]: Returns layout attributes for the drop target at the specified point.
//   - [INSCollectionViewLayout.LayoutAttributesForInterItemGapBeforeIndexPath]: Returns layout attributes for the inter-item gap at the specified location in your layout.
//   - [INSCollectionViewLayout.TargetContentOffsetForProposedContentOffset]: Returns the offset value to use after an animated layout update or change.
//   - [INSCollectionViewLayout.TargetContentOffsetForProposedContentOffsetWithScrollingVelocity]: Returns the offset value to use for the collection view’s content at the end of scrolling.
//
// # Responding to Collection View Updates
//
//   - [INSCollectionViewLayout.PrepareForCollectionViewUpdates]: Performs needed tasks before items are inserted, deleted, or moved within the collection view.
//   - [INSCollectionViewLayout.FinalizeCollectionViewUpdates]: Performs needed steps after items are inserted, deleted, or moved within a collection view.
//   - [INSCollectionViewLayout.IndexPathsToInsertForSupplementaryViewOfKind]: Returns the index paths for any supplementary views that the layout object wants to add to the collection view.
//   - [INSCollectionViewLayout.IndexPathsToInsertForDecorationViewOfKind]: Returns the index paths for any decoration views that the layout object wants to add to the collection view.
//   - [INSCollectionViewLayout.InitialLayoutAttributesForAppearingItemAtIndexPath]: Returns the starting layout information for an item being inserted into the collection view.
//   - [INSCollectionViewLayout.InitialLayoutAttributesForAppearingSupplementaryElementOfKindAtIndexPath]: Returns the starting layout information for a supplementary view being added to the collection view.
//   - [INSCollectionViewLayout.InitialLayoutAttributesForAppearingDecorationElementOfKindAtIndexPath]: Returns the starting layout information for a decoration view being added to the collection view.
//   - [INSCollectionViewLayout.IndexPathsToDeleteForSupplementaryViewOfKind]: Returns the index paths for any supplementary views that the layout object wants to remove from the collection view.
//   - [INSCollectionViewLayout.IndexPathsToDeleteForDecorationViewOfKind]: Returns index paths for any decoration views that the layout object wants to remove from the collection view.
//   - [INSCollectionViewLayout.FinalLayoutAttributesForDisappearingItemAtIndexPath]: Returns the ending layout information for an item being removed from the collection view.
//   - [INSCollectionViewLayout.FinalLayoutAttributesForDisappearingSupplementaryElementOfKindAtIndexPath]: Returns the ending layout information for a supplementary view being removed from the collection view.
//   - [INSCollectionViewLayout.FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath]: Returns the ending layout information for a decoration view being removed from the collection view.
//
// # Invalidating the Layout
//
//   - [INSCollectionViewLayout.InvalidateLayout]: Invalidates all layout information and triggers a layout update.
//   - [INSCollectionViewLayout.InvalidateLayoutWithContext]: Invalidates specific parts of the layout using the specified context object.
//   - [INSCollectionViewLayout.ShouldInvalidateLayoutForBoundsChange]: Returns a Boolean indicating whether a bounds change triggers a layout update.
//   - [INSCollectionViewLayout.ShouldInvalidateLayoutForPreferredLayoutAttributesWithOriginalAttributes]: Returns a Boolean indicating whether changes to a cell’s layout attributes trigger a larger layout update.
//   - [INSCollectionViewLayout.InvalidationContextForBoundsChange]: Returns an invalidation context object that defines the portions of the layout that need to be updated.
//   - [INSCollectionViewLayout.InvalidationContextForPreferredLayoutAttributesWithOriginalAttributes]: Returns an invalidation context object that defines the portions of the layout that need to be updated.
//
// # Coordinating Animated Changes
//
//   - [INSCollectionViewLayout.PrepareForAnimatedBoundsChange]: Prepares the layout object for animated changes to the collection view’s bounds or for the insertion or deletion of items.
//   - [INSCollectionViewLayout.FinalizeAnimatedBoundsChange]: Cleans up after any animated changes to the collection view’s bounds or after the insertion or deletion of items.
//
// # Registering Decoration Views
//
//   - [INSCollectionViewLayout.RegisterClassForDecorationViewOfKind]: Registers a class to use when creating the layout’s decoration views.
//   - [INSCollectionViewLayout.RegisterNibForDecorationViewOfKind]: Registers a nib file to use when creating the layout’s decoration views.
//
// # Transitioning Between Layouts
//
//   - [INSCollectionViewLayout.PrepareForTransitionFromLayout]: Prepares the layout object to be installed in the collection view.
//   - [INSCollectionViewLayout.PrepareForTransitionToLayout]: Prepares the layout object to be uninstalled from the collection view.
//   - [INSCollectionViewLayout.FinalizeLayoutTransition]: Performs any final steps related to a layout transition before the transition animations actually occur.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout
type INSCollectionViewLayout interface {
	objectivec.IObject

	// Topic: Getting the Collection View

	// The collection view object currently using this layout.
	CollectionView() INSCollectionView

	// Topic: Providing Layout Information

	// Prepares the layout object to begin laying out content.
	PrepareLayout()
	// The width and height of the collection view’s contents.
	CollectionViewContentSize() corefoundation.CGSize
	// Returns the layout attribute objects for all items and views in the specified rectangle.
	LayoutAttributesForElementsInRect(rect corefoundation.CGRect) []NSCollectionViewLayoutAttributes
	// Returns the layout attributes for the item at the specified index path.
	LayoutAttributesForItemAtIndexPath(indexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes
	// Returns the layout attributes of the supplementary view at the specified location in your layout.
	LayoutAttributesForSupplementaryViewOfKindAtIndexPath(elementKind NSCollectionViewSupplementaryElementKind, indexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes
	// Returns the layout attributes of the decoration view at the specified location in your layout.
	LayoutAttributesForDecorationViewOfKindAtIndexPath(elementKind NSCollectionViewDecorationElementKind, indexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes
	// Returns layout attributes for the drop target at the specified point.
	LayoutAttributesForDropTargetAtPoint(pointInCollectionView corefoundation.CGPoint) INSCollectionViewLayoutAttributes
	// Returns layout attributes for the inter-item gap at the specified location in your layout.
	LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes
	// Returns the offset value to use after an animated layout update or change.
	TargetContentOffsetForProposedContentOffset(proposedContentOffset corefoundation.CGPoint) corefoundation.CGPoint
	// Returns the offset value to use for the collection view’s content at the end of scrolling.
	TargetContentOffsetForProposedContentOffsetWithScrollingVelocity(proposedContentOffset corefoundation.CGPoint, velocity corefoundation.CGPoint) corefoundation.CGPoint

	// Topic: Responding to Collection View Updates

	// Performs needed tasks before items are inserted, deleted, or moved within the collection view.
	PrepareForCollectionViewUpdates(updateItems []NSCollectionViewUpdateItem)
	// Performs needed steps after items are inserted, deleted, or moved within a collection view.
	FinalizeCollectionViewUpdates()
	// Returns the index paths for any supplementary views that the layout object wants to add to the collection view.
	IndexPathsToInsertForSupplementaryViewOfKind(elementKind NSCollectionViewSupplementaryElementKind) foundation.INSSet
	// Returns the index paths for any decoration views that the layout object wants to add to the collection view.
	IndexPathsToInsertForDecorationViewOfKind(elementKind NSCollectionViewDecorationElementKind) foundation.INSSet
	// Returns the starting layout information for an item being inserted into the collection view.
	InitialLayoutAttributesForAppearingItemAtIndexPath(itemIndexPath foundation.INSIndexPath) NSCollectionViewLayout
	// Returns the starting layout information for a supplementary view being added to the collection view.
	InitialLayoutAttributesForAppearingSupplementaryElementOfKindAtIndexPath(elementKind NSCollectionViewSupplementaryElementKind, elementIndexPath foundation.INSIndexPath) NSCollectionViewLayout
	// Returns the starting layout information for a decoration view being added to the collection view.
	InitialLayoutAttributesForAppearingDecorationElementOfKindAtIndexPath(elementKind NSCollectionViewDecorationElementKind, decorationIndexPath foundation.INSIndexPath) NSCollectionViewLayout
	// Returns the index paths for any supplementary views that the layout object wants to remove from the collection view.
	IndexPathsToDeleteForSupplementaryViewOfKind(elementKind NSCollectionViewSupplementaryElementKind) foundation.INSSet
	// Returns index paths for any decoration views that the layout object wants to remove from the collection view.
	IndexPathsToDeleteForDecorationViewOfKind(elementKind NSCollectionViewDecorationElementKind) foundation.INSSet
	// Returns the ending layout information for an item being removed from the collection view.
	FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes
	// Returns the ending layout information for a supplementary view being removed from the collection view.
	FinalLayoutAttributesForDisappearingSupplementaryElementOfKindAtIndexPath(elementKind NSCollectionViewSupplementaryElementKind, elementIndexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes
	// Returns the ending layout information for a decoration view being removed from the collection view.
	FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath(elementKind NSCollectionViewDecorationElementKind, decorationIndexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes

	// Topic: Invalidating the Layout

	// Invalidates all layout information and triggers a layout update.
	InvalidateLayout()
	// Invalidates specific parts of the layout using the specified context object.
	InvalidateLayoutWithContext(context INSCollectionViewLayoutInvalidationContext)
	// Returns a Boolean indicating whether a bounds change triggers a layout update.
	ShouldInvalidateLayoutForBoundsChange(newBounds corefoundation.CGRect) bool
	// Returns a Boolean indicating whether changes to a cell’s layout attributes trigger a larger layout update.
	ShouldInvalidateLayoutForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes INSCollectionViewLayoutAttributes, originalAttributes INSCollectionViewLayoutAttributes) bool
	// Returns an invalidation context object that defines the portions of the layout that need to be updated.
	InvalidationContextForBoundsChange(newBounds corefoundation.CGRect) INSCollectionViewLayoutInvalidationContext
	// Returns an invalidation context object that defines the portions of the layout that need to be updated.
	InvalidationContextForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes INSCollectionViewLayoutAttributes, originalAttributes INSCollectionViewLayoutAttributes) INSCollectionViewLayoutInvalidationContext

	// Topic: Coordinating Animated Changes

	// Prepares the layout object for animated changes to the collection view’s bounds or for the insertion or deletion of items.
	PrepareForAnimatedBoundsChange(oldBounds corefoundation.CGRect)
	// Cleans up after any animated changes to the collection view’s bounds or after the insertion or deletion of items.
	FinalizeAnimatedBoundsChange()

	// Topic: Registering Decoration Views

	// Registers a class to use when creating the layout’s decoration views.
	RegisterClassForDecorationViewOfKind(viewClass objc.Class, elementKind NSCollectionViewDecorationElementKind)
	// Registers a nib file to use when creating the layout’s decoration views.
	RegisterNibForDecorationViewOfKind(nib INSNib, elementKind NSCollectionViewDecorationElementKind)

	// Topic: Transitioning Between Layouts

	// Prepares the layout object to be installed in the collection view.
	PrepareForTransitionFromLayout(oldLayout INSCollectionViewLayout)
	// Prepares the layout object to be uninstalled from the collection view.
	PrepareForTransitionToLayout(newLayout INSCollectionViewLayout)
	// Performs any final steps related to a layout transition before the transition animations actually occur.
	FinalizeLayoutTransition()

	// The layout object used to organize the collection view’s content.
	CollectionViewLayout() INSCollectionViewLayout
	SetCollectionViewLayout(value INSCollectionViewLayout)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c NSCollectionViewLayout) Init() NSCollectionViewLayout {
	rv := objc.Send[NSCollectionViewLayout](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewLayout) Autorelease() NSCollectionViewLayout {
	rv := objc.Send[NSCollectionViewLayout](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewLayout creates a new NSCollectionViewLayout instance.
func NewNSCollectionViewLayout() NSCollectionViewLayout {
	class := getNSCollectionViewLayoutClass()
	rv := objc.Send[NSCollectionViewLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Prepares the layout object to begin laying out content.
//
// # Discussion
// 
// The default implementation of this method does nothing. During the layout
// cycle, the collection view calls this method first to give you a chance to
// prepare any data needed during the layout operation. When defining a custom
// layout, you can override this method and use it to set up data structures
// or perform any initial computations needed to perform the layout later.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/prepare()
func (c NSCollectionViewLayout) PrepareLayout() {
	objc.Send[objc.ID](c.ID, objc.Sel("prepareLayout"))
}
// Returns the layout attribute objects for all items and views in the
// specified rectangle.
//
// rect: The rectangle (specified in the collection view’s coordinate system)
// containing the target views.
//
// # Return Value
// 
// An array of layout attribute objects containing the layout information for
// the enclosed items and views.
//
// # Discussion
// 
// Subclasses must override this method. In your implementation, return layout
// attributes for all items, supplementary views, and decoration views that
// intersect the specified rectangle.
// 
// For each element, always return a layout object of the appropriate type
// (item, supplementary, or decoration). The collection view differentiates
// between attributes for items, supplementary views, and decoration views and
// uses the differences to decide how to create and manage the corresponding
// views. Use the [LayoutAttributesForItemAtIndexPath],
// [LayoutAttributesForSupplementaryViewOfKindAtIndexPath], and
// [LayoutAttributesForDecorationViewOfKindAtIndexPath] methods to create new
// layout attribute objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForElements(in:)
func (c NSCollectionViewLayout) LayoutAttributesForElementsInRect(rect corefoundation.CGRect) []NSCollectionViewLayoutAttributes {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("layoutAttributesForElementsInRect:"), rect)
	return objc.ConvertSlice(rv, func(id objc.ID) NSCollectionViewLayoutAttributes {
		return NSCollectionViewLayoutAttributesFromID(id)
	})
}
// Returns the layout attributes for the item at the specified index path.
//
// indexPath: The index path of the item whose attributes are requested.
//
// # Return Value
// 
// A layout attributes object containing the layout information to apply to
// the item.
//
// # Discussion
// 
// Subclasses must override this method. In your implementation, create an
// instance of the appropriate layout attributes class and fill the resulting
// object with the layout information for the specified item. The default
// implementation of this method does nothing and returns `nil`.
// 
// You can also call this method from other layout-related methods when you
// want to retrieve layout information for items. Call this method only for
// items. Do not call it to retrieve layout attributes for supplementary views
// or decoration views.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForItem(at:)
func (c NSCollectionViewLayout) LayoutAttributesForItemAtIndexPath(indexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("layoutAttributesForItemAtIndexPath:"), indexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}
// Returns the layout attributes of the supplementary view at the specified
// location in your layout.
//
// elementKind: A string that identifies the type of the supplementary view. Use this
// string to differentiate the supplementary views in a given section.
//
// indexPath: The index path of the supplementary view. Use this parameter to determine
// which section contains the supplementary view.
//
// # Return Value
// 
// A layout attributes object containing the layout information to apply to
// the supplementary view.
//
// # Discussion
// 
// If your layout includes supplementary views, you must override this method.
// In your implementation, create an instance of the appropriate layout
// attributes class and fill the resulting object with the layout information
// for the corresponding supplementary view. You define the supported
// supplementary views by assigning each one a string that identifies its
// kind. Use the `elementKind` and `indexPath` properties to identify the
// specific supplementary view whose attributes were requested.
// 
// You can call this method from other layout-related methods when you want to
// retrieve layout information for supplementary views. Call this method only
// for supplementary views. Do not call it to retrieve layout attributes for
// items or decoration views.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForSupplementaryView(ofKind:at:)
func (c NSCollectionViewLayout) LayoutAttributesForSupplementaryViewOfKindAtIndexPath(elementKind NSCollectionViewSupplementaryElementKind, indexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("layoutAttributesForSupplementaryViewOfKind:atIndexPath:"), objc.String(string(elementKind)), indexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}
// Returns the layout attributes of the decoration view at the specified
// location in your layout.
//
// elementKind: A string that identifies the type of the decoration view. Use this string
// to differentiate the decoration views in a given section.
//
// indexPath: The index path of the decoration view. Use this parameter to determine
// which section contains the decoration view.
//
// # Return Value
// 
// A layout attributes object containing the layout information to apply to
// the decoration view.
//
// # Discussion
// 
// If your layout includes decoration views, you must override this method. In
// your implementation, create an instance of the appropriate layout
// attributes class and fill the resulting object with the layout information
// for the corresponding decoration view. You define the supported decoration
// views by assigning each one a string that identifies its kind. Use the
// `elementKind` and `indexPath` properties to identify the specific
// decoration view whose attributes were requested.
// 
// You can call this method from other layout-related methods when you want to
// retrieve layout information for decoration views. Call this method only for
// decoration views. Do not call it to retrieve layout attributes for items or
// supplementary views.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForDecorationView(ofKind:at:)
func (c NSCollectionViewLayout) LayoutAttributesForDecorationViewOfKindAtIndexPath(elementKind NSCollectionViewDecorationElementKind, indexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("layoutAttributesForDecorationViewOfKind:atIndexPath:"), objc.String(string(elementKind)), indexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}
// Returns layout attributes for the drop target at the specified point.
//
// pointInCollectionView: A point in the collection view’s coordinate system. Use this point to
// determine whether the drop would occur on an item or between items.
//
// # Return Value
// 
// A layout attributes object that identifies the element in which the drop
// would occur.
//
// # Discussion
// 
// The default implementation of this method tests the specified point to see
// if it lands inside an item, supplementary view, or decoration view of the
// collection view. If it does, the method returns the layout attributes for
// that element.
// 
// Layouts that support inter-item gaps as drop targets must override this
// method and use it to return the layout attributes that represent that gap.
// In your implementation, calculate the index path just after the gap and
// pass that value to the [LayoutAttributesForInterItemGapBeforeIndexPath]
// class method of [NSCollectionViewLayoutAttributes]. Set the [Frame]
// property of the resulting attributes object to the rectangle that best
// represents the gap and also contains the specified point. When overriding
// this method, you can call `super` at any time to get the default behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForDropTarget(at:)
func (c NSCollectionViewLayout) LayoutAttributesForDropTargetAtPoint(pointInCollectionView corefoundation.CGPoint) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("layoutAttributesForDropTargetAtPoint:"), pointInCollectionView)
	return NSCollectionViewLayoutAttributesFromID(rv)
}
// Returns layout attributes for the inter-item gap at the specified location
// in your layout.
//
// indexPath: The index path of the item that follows the inter-item gap. For a gap that
// follows the last item in the section, set the item property to the total
// number of items in the section.
//
// # Return Value
// 
// A layout attributes object containing the layout information to apply to
// the inter-item gap, or `nil` if no attributes are available.
//
// # Discussion
// 
// The default implementation of this method returns `nil`. Subclasses can
// override this method to provide layout attributes for inter-item gaps. In
// your implementation, use the specified index path to compute the location
// of the gap in collection view’s content. If the gap represents a valid
// location, use the [LayoutAttributesForInterItemGapBeforeIndexPath] class
// method of [NSCollectionViewLayoutAttributes] to create a new layout
// attributes object and set the [Frame] property to the rectangle you
// computed.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesForInterItemGap(before:)
func (c NSCollectionViewLayout) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("layoutAttributesForInterItemGapBeforeIndexPath:"), indexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}
// Returns the offset value to use after an animated layout update or change.
//
// proposedContentOffset: The proposed point (in the collection view’s coordinate space) for the
// lower-left corner of the visible content. The collection view calculates
// this value as the most likely value to use at the end of animations.
//
// # Return Value
// 
// The offset value that you want to use for the content.
//
// # Discussion
// 
// During layout updates, or when transitioning between layouts, the
// collection view calls this method to give you the opportunity to tweak the
// position of the collection view’s content at the end of animations. The
// default implementation of this method returns the value in the
// `proposedContentOffset` parameter. Subclasses can override it and return an
// offset value that positions content in a way that is more optimal for the
// custom layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/targetContentOffset(forProposedContentOffset:)
func (c NSCollectionViewLayout) TargetContentOffsetForProposedContentOffset(proposedContentOffset corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("targetContentOffsetForProposedContentOffset:"), proposedContentOffset)
	return corefoundation.CGPoint(rv)
}
// Returns the offset value to use for the collection view’s content at the
// end of scrolling.
//
// proposedContentOffset: The proposed point (in the collection view’s coordinate space) for the
// lower-left corner of the visible content. The collection view calculates
// this value as the most likely value to use at the end of animations.
//
// velocity: The current horizontal and vertical scrolling velocities. The value is
// specified in points per second.
//
// # Return Value
// 
// The offset value that you want to use for the content.
//
// # Discussion
// 
// Use this method to position the collection view’s content appropriately
// after scrolling. This method tells the scroll view where to stop scrolling
// so that the collection view’s content is displayed optimally. For
// example, you might use this method to adjust the proposed content offset so
// that it falls on a boundary between rows of items, as opposed to stopping
// in the middle of a row.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/targetContentOffset(forProposedContentOffset:withScrollingVelocity:)
func (c NSCollectionViewLayout) TargetContentOffsetForProposedContentOffsetWithScrollingVelocity(proposedContentOffset corefoundation.CGPoint, velocity corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("targetContentOffsetForProposedContentOffset:withScrollingVelocity:"), proposedContentOffset, velocity)
	return corefoundation.CGPoint(rv)
}
// Performs needed tasks before items are inserted, deleted, or moved within
// the collection view.
//
// updateItems: An array of [NSCollectionViewUpdateItem] objects that identify the changes
// being made.
//
// # Discussion
// 
// When items are inserted, deleted, or moved, the collection view calls this
// method to report those changes to the layout object. The default
// implementation uses the provided information to plan the layout animations
// needed to respond to the changes. Subclasses can override this method and
// use it to prepare for any custom changes, but you should always call
// `super` at some point in your implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/prepare(forCollectionViewUpdates:)
func (c NSCollectionViewLayout) PrepareForCollectionViewUpdates(updateItems []NSCollectionViewUpdateItem) {
	objc.Send[objc.ID](c.ID, objc.Sel("prepareForCollectionViewUpdates:"), objectivec.IObjectSliceToNSArray(updateItems))
}
// Performs needed steps after items are inserted, deleted, or moved within a
// collection view.
//
// # Discussion
// 
// When items are inserted, deleted, or moved, the collection view calls this
// method as a final step before performing the associated animations. This
// method is called within the animation block used to configure the
// insertion, deletion, and move animations, so you can use this method to
// configure additional animations that you want to run at the same time. You
// can also use this method to perform any cleanup steps to account for the
// changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalizeCollectionViewUpdates()
func (c NSCollectionViewLayout) FinalizeCollectionViewUpdates() {
	objc.Send[objc.ID](c.ID, objc.Sel("finalizeCollectionViewUpdates"))
}
// Returns the index paths for any supplementary views that the layout object
// wants to add to the collection view.
//
// elementKind: The type of the supplementary views to add.
//
// # Return Value
// 
// The set of [NSIndexPath] objects representing the supplementary views to
// insert, or an empty array if you do not want to insert any supplementary
// views.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// When your app inserts or deletes items or sections in the collection view,
// the collection view calls this method for each of the registered
// supplementary view types. The default implementation returns an empty
// array, but you can override it and return index paths for each
// supplementary view you want to add. For example, when a section is added,
// you might want to add the supplementary views that belong in that section.
// In that case, you would add index paths to the array that contain the
// section numbers that were added.
// 
// The index paths you return should always contain a valid section number,
// but the item number is optional. The item number is necessary only if you
// support multiple supplementary views of the same type in a single section.
// If you do, your layout object can use the item numbers internally to
// differentiate the supplementary views.
// 
// Subclasses are expected to override this method, as needed, and provide any
// appropriate index paths.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/indexPathsToInsertForSupplementaryView(ofKind:)
func (c NSCollectionViewLayout) IndexPathsToInsertForSupplementaryViewOfKind(elementKind NSCollectionViewSupplementaryElementKind) foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPathsToInsertForSupplementaryViewOfKind:"), objc.String(string(elementKind)))
	return foundation.NSSetFromID(rv)
}
// Returns the index paths for any decoration views that the layout object
// wants to add to the collection view.
//
// elementKind: The type of the decoration views to add.
//
// # Return Value
// 
// The set of [NSIndexPath] objects representing the decoration views to
// insert, or an empty array if you do not want to insert any decoration
// views.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// When your app inserts or deletes items or sections in the collection view,
// the collection view calls this method for each of the layout’s registered
// decoration view types. The default implementation returns an empty array,
// but you can override it and return index paths for each decoration view you
// want to add. For example, when a section is added, you might want to add
// the decoration views that are used to adorn that section. In that case, you
// would add index paths to the array that contain the section numbers that
// were added.
// 
// The index paths you return should always contain a valid section number,
// but the item number is optional. The item number is necessary only if you
// support multiple decoration views of the same type in a single section. If
// you do, your layout object can use the item numbers internally to
// differentiate the decoration views.
// 
// Subclasses are expected to override this method, as needed, and provide any
// appropriate index paths.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/indexPathsToInsertForDecorationView(ofKind:)
func (c NSCollectionViewLayout) IndexPathsToInsertForDecorationViewOfKind(elementKind NSCollectionViewDecorationElementKind) foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPathsToInsertForDecorationViewOfKind:"), objc.String(string(elementKind)))
	return foundation.NSSetFromID(rv)
}
// Returns the starting layout information for an item being inserted into the
// collection view.
//
// itemIndexPath: The index path of the item being inserted. You can use this path to
// retrieve any relevant information from the collection view’s data source.
//
// # Return Value
// 
// The layout attributes object that describes the item’s position and
// properties at the start of animations.
//
// # Discussion
// 
// When your app inserts new items into the collection view, the collection
// view calls this method for each item you insert. Because new items are not
// yet visible in the collection view, the attributes you return represent the
// item’s starting state. For example, you might return attributes that
// position the item offscreen or set its initial alpha to `0`. The collection
// view uses the attributes you return as the starting point for any
// animations. (The end point of the animation is the item’s new location
// and attributes.) If you return `nil`, the layout uses the item’s final
// attributes for both the start point and end point of the animation.
// 
// The default implementation of this method returns `nil`. Subclasses are
// expected to override this method, as needed, and provide any initial
// attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/initialLayoutAttributesForAppearingItem(at:)
func (c NSCollectionViewLayout) InitialLayoutAttributesForAppearingItemAtIndexPath(itemIndexPath foundation.INSIndexPath) NSCollectionViewLayout {
	rv := objc.Send[NSCollectionViewLayout](c.ID, objc.Sel("initialLayoutAttributesForAppearingItemAtIndexPath:"), itemIndexPath)
	return rv
}
// Returns the starting layout information for a supplementary view being
// added to the collection view.
//
// elementKind: The type of the supplementary view being added.
//
// elementIndexPath: The index path of the supplementary view. You can use this path to retrieve
// any relevant information from the collection view’s data source.
//
// # Return Value
// 
// The layout attributes object that describes the supplementary view’s
// position and properties at the start of animations.
//
// # Discussion
// 
// When your layout object adds supplementary views in response to other
// changes in the collection view, the collection view calls this method for
// each supplementary view before animating it onscreen. Because new
// supplementary views are not yet visible in the collection view, the
// attributes you return represent the view’s starting state. For example,
// you might return attributes that position the supplementary view offscreen
// or set its initial alpha to `0`. The collection view uses the attributes
// you return as the starting point for any animations. (The end point of the
// animation is the view’s new location and attributes.) If you return
// `nil`, the layout uses the supplementary view’s final attributes for both
// the start point and end point of the animation.
// 
// The default implementation of this method returns `nil`. Subclasses are
// expected to override this method, as needed, and provide any initial
// attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/initialLayoutAttributesForAppearingSupplementaryElement(ofKind:at:)
func (c NSCollectionViewLayout) InitialLayoutAttributesForAppearingSupplementaryElementOfKindAtIndexPath(elementKind NSCollectionViewSupplementaryElementKind, elementIndexPath foundation.INSIndexPath) NSCollectionViewLayout {
	rv := objc.Send[NSCollectionViewLayout](c.ID, objc.Sel("initialLayoutAttributesForAppearingSupplementaryElementOfKind:atIndexPath:"), objc.String(string(elementKind)), elementIndexPath)
	return rv
}
// Returns the starting layout information for a decoration view being added
// to the collection view.
//
// elementKind: The type of the decoration view being added.
//
// decorationIndexPath: The index path of the decoration view. You can use this path to retrieve
// any relevant information from the collection view’s data source.
//
// # Return Value
// 
// The layout attributes object that describes the decoration view’s
// position and properties at the start of animations.
//
// # Discussion
// 
// When your layout object adds decoration views in response to other changes
// in the collection view, the collection view calls this method for each
// decoration view before animating it onscreen. Because new decoration views
// are not yet visible in the collection view, the attributes you return
// represent the view’s starting state. For example, you might return
// attributes that position the decoration view offscreen or set its initial
// alpha to `0`. The collection view uses the attributes you return as the
// starting point for any animations. (The end point of the animation is the
// view’s new location and attributes.) If you return `nil`, the layout uses
// the decoration view’s final attributes for both the start point and end
// point of the animation.
// 
// The default implementation of this method returns `nil`. Subclasses are
// expected to override this method, as needed, and provide any initial
// attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/initialLayoutAttributesForAppearingDecorationElement(ofKind:at:)
func (c NSCollectionViewLayout) InitialLayoutAttributesForAppearingDecorationElementOfKindAtIndexPath(elementKind NSCollectionViewDecorationElementKind, decorationIndexPath foundation.INSIndexPath) NSCollectionViewLayout {
	rv := objc.Send[NSCollectionViewLayout](c.ID, objc.Sel("initialLayoutAttributesForAppearingDecorationElementOfKind:atIndexPath:"), objc.String(string(elementKind)), decorationIndexPath)
	return rv
}
// Returns the index paths for any supplementary views that the layout object
// wants to remove from the collection view.
//
// elementKind: The type of the supplementary views to remove.
//
// # Return Value
// 
// The set of [NSIndexPath] objects representing the supplementary views to
// remove, or an empty array if you do not want to remove any supplementary
// views.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// When your app inserts or deletes items or sections in the collection view,
// the collection view calls this method for each of the registered
// supplementary view types. The default implementation returns an empty
// array, but you can override it and return index paths for each
// supplementary view you want to remove. For example, when a section is
// removed, you might want to remove the supplementary views associated with
// that section. In that case, you would add index paths to the array that
// contain the section numbers that were removed.
// 
// The index paths you return should always contain a valid section number,
// but the item number is optional. The item number is necessary only if you
// support multiple supplementary views of the same type in a single section.
// If you do, your layout object can use the item numbers internally to
// differentiate the supplementary views.
// 
// Subclasses are expected to override this method, as needed, and provide any
// appropriate index paths.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/indexPathsToDeleteForSupplementaryView(ofKind:)
func (c NSCollectionViewLayout) IndexPathsToDeleteForSupplementaryViewOfKind(elementKind NSCollectionViewSupplementaryElementKind) foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPathsToDeleteForSupplementaryViewOfKind:"), objc.String(string(elementKind)))
	return foundation.NSSetFromID(rv)
}
// Returns index paths for any decoration views that the layout object wants
// to remove from the collection view.
//
// elementKind: The type of the decoration views to remove.
//
// # Return Value
// 
// The set of [NSIndexPath] objects representing the decoration views to
// remove, or an empty array if you do not want to remove any decoration
// views.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// When your app inserts or deletes items or sections in the collection view,
// the collection view calls this method for each of the registered decoration
// view types. The default implementation returns an empty array, but you can
// override it and return index paths for each decoration view you want to
// remove. For example, when a section is removed, you might want to remove
// the decoration views associated with that section. In that case, you would
// add index paths to the array that contain the section numbers that were
// removed.
// 
// The index paths you return should always contain a valid section number,
// but the item number is optional. The item number is necessary only if you
// support multiple supplementary views of the same type in a single section.
// If you do, your layout object can use the item numbers internally to
// differentiate the supplementary views.
// 
// Subclasses are expected to override this method, as needed, and provide any
// appropriate index paths.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/indexPathsToDeleteForDecorationView(ofKind:)
func (c NSCollectionViewLayout) IndexPathsToDeleteForDecorationViewOfKind(elementKind NSCollectionViewDecorationElementKind) foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPathsToDeleteForDecorationViewOfKind:"), objc.String(string(elementKind)))
	return foundation.NSSetFromID(rv)
}
// Returns the ending layout information for an item being removed from the
// collection view.
//
// itemIndexPath: The index path of the item being removed. You can use this path to retrieve
// any relevant information from the collection view’s data source.
//
// # Return Value
// 
// The layout attributes object that describes the item’s position and
// properties at the end of animations.
//
// # Discussion
// 
// When your app removes items from the collection view, the collection view
// calls this method for each item you remove. Use this method to specify the
// layout attributes of the item after it has been removed. For example, you
// might return attributes that position the item offscreen or set its final
// alpha to `0`. The collection view uses the attributes you return as the end
// point for any animations. (The start point of the animation is the item’s
// current location and attributes.) If you return `nil`, the layout uses the
// item’s current attributes for both the start point and end point of the
// animation.
// 
// The default implementation of this method returns `nil`. Subclasses are
// expected to override this method, as needed, and provide any final
// attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalLayoutAttributesForDisappearingItem(at:)
func (c NSCollectionViewLayout) FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("finalLayoutAttributesForDisappearingItemAtIndexPath:"), itemIndexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}
// Returns the ending layout information for a supplementary view being
// removed from the collection view.
//
// elementKind: The type of the decoration view being removed.
//
// elementIndexPath: The index path of the supplementary view being removed. You can use this
// path to identify the view internally.
//
// # Return Value
// 
// The layout attributes object that describes the supplementary view’s
// position and properties at the end of animations.
//
// # Discussion
// 
// When your layout object removes supplementary views in response to other
// changes in the collection view, the collection view calls this method for
// each supplementary view you remove. Use this method to specify the layout
// attributes for the view after it has been removed. For example, you might
// return attributes that position the supplementary view offscreen or set its
// alpha to `0`. The collection view uses the attributes you return as the end
// point for any animations. (The start point of the animation is the view’s
// current location and attributes.) If you return `nil`, the layout uses the
// supplementary view’s current attributes for both the start point and end
// point of the animation.
// 
// The default implementation of this method returns `nil`. Subclasses are
// expected to override this method, as needed, and provide any initial
// attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalLayoutAttributesForDisappearingSupplementaryElement(ofKind:at:)
func (c NSCollectionViewLayout) FinalLayoutAttributesForDisappearingSupplementaryElementOfKindAtIndexPath(elementKind NSCollectionViewSupplementaryElementKind, elementIndexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("finalLayoutAttributesForDisappearingSupplementaryElementOfKind:atIndexPath:"), objc.String(string(elementKind)), elementIndexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}
// Returns the ending layout information for a decoration view being removed
// from the collection view.
//
// elementKind: The type of the decoration view being removed.
//
// decorationIndexPath: The index path of the decoration view being removed. You can use this path
// to identify the view internally.
//
// # Return Value
// 
// The layout attributes object that describes the decoration view’s
// position and properties at the end of animations.
//
// # Discussion
// 
// When your layout object removes decoration views in response to other
// changes in the collection view, the collection view calls this method for
// each decoration view you remove. Use this method to specify the layout
// attributes for the view after it has been removed. For example, you might
// return attributes that position the decoration view offscreen or set its
// alpha to `0`. The collection view uses the attributes you return as the end
// point for any animations. (The start point of the animation is the view’s
// current location and attributes.) If you return `nil`, the layout uses the
// decoration view’s current attributes for both the start point and end
// point of the animation.
// 
// The default implementation of this method returns `nil`. Subclasses are
// expected to override this method, as needed, and provide any initial
// attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalLayoutAttributesForDisappearingDecorationElement(ofKind:at:)
func (c NSCollectionViewLayout) FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath(elementKind NSCollectionViewDecorationElementKind, decorationIndexPath foundation.INSIndexPath) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("finalLayoutAttributesForDisappearingDecorationElementOfKind:atIndexPath:"), objc.String(string(elementKind)), decorationIndexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}
// Invalidates all layout information and triggers a layout update.
//
// # Discussion
// 
// Call this method when you make changes that require updating all of the
// current layout information. This method marks the layout as invalid and
// returns right away, so you can call this method multiple times from the
// same block of code without triggering multiple layout updates. During the
// next update cycle, the collection view requests new layout information and
// updates its contents accordingly.
// 
// If you override this method, you must call `super` at some point in your
// implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/invalidateLayout()
func (c NSCollectionViewLayout) InvalidateLayout() {
	objc.Send[objc.ID](c.ID, objc.Sel("invalidateLayout"))
}
// Invalidates specific parts of the layout using the specified context
// object.
//
// context: The context object indicating which parts of the layout need to be updated.
//
// # Discussion
// 
// Call this method when you make changes that need to be reflected by the
// collection view, but which do not require the replacement of all of the
// layout information. You use this method to minimize the work performed by
// the layout object. Instead of optimizing everything, the specified context
// object indicates which parts of the layout need to be recomputed. All other
// layout information is left alone.
// 
// When implementing a custom layout, you can override this method and use it
// to process information provided by a custom invalidation context. You are
// not required to provide a custom invalidation context but might do so if
// you are able to provide additional properties that can help optimize layout
// updates. If you override this method, you must call `super` at some point
// in your implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/invalidateLayout(with:)
func (c NSCollectionViewLayout) InvalidateLayoutWithContext(context INSCollectionViewLayoutInvalidationContext) {
	objc.Send[objc.ID](c.ID, objc.Sel("invalidateLayoutWithContext:"), context)
}
// Returns a Boolean indicating whether a bounds change triggers a layout
// update.
//
// newBounds: The new bounds of the collection view.
//
// # Return Value
// 
// [true] if a layout should be invalidated or [false] if the layout is still
// valid.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The default implementation of this method returns [false]. You can override
// this method in your custom layout classes and return a different value as
// needed. Your implementation of the method should determine if the new
// bounds would cause changes to the layout of other portions of the
// collection view.
// 
// If you return [true] from this method, the collection view invalidates the
// layout using the [InvalidateLayoutWithContext] method. The invalidation
// context passed to that method is created using the
// [InvalidationContextForBoundsChange] method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/shouldInvalidateLayout(forBoundsChange:)
func (c NSCollectionViewLayout) ShouldInvalidateLayoutForBoundsChange(newBounds corefoundation.CGRect) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("shouldInvalidateLayoutForBoundsChange:"), newBounds)
	return rv
}
// Returns a Boolean indicating whether changes to a cell’s layout
// attributes trigger a larger layout update.
//
// preferredAttributes: The preferred layout attributes of an element.
//
// originalAttributes: The attributes that the layout object originally suggested for the item.
//
// # Return Value
// 
// [true] if the layout should be invalidated or [false] if it should not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The default implementation of this method returns NO to indicate that
// layout is not needed. You can override this method in your custom layout
// classes and return a different value as needed. Your implementation of this
// method should determine if the new attributes would cause changes to the
// layout of other portions of the collection view.
// 
// If you return [true] from this method, the collection view invalidates the
// layout using the [InvalidateLayoutWithContext] method. The invalidation
// context passed to that method is created using the
// [InvalidationContextForPreferredLayoutAttributesWithOriginalAttributes]
// method.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/shouldInvalidateLayout(forPreferredLayoutAttributes:withOriginalAttributes:)
func (c NSCollectionViewLayout) ShouldInvalidateLayoutForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes INSCollectionViewLayoutAttributes, originalAttributes INSCollectionViewLayoutAttributes) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("shouldInvalidateLayoutForPreferredLayoutAttributes:withOriginalAttributes:"), preferredAttributes, originalAttributes)
	return rv
}
// Returns an invalidation context object that defines the portions of the
// layout that need to be updated.
//
// newBounds: The new bounds for the collection view.
//
// # Return Value
// 
// An invalidation context that describes the changes to be made. This value
// is never `nil`.
//
// # Discussion
// 
// The default implementation of this method creates an instance of the class
// returned by the [InvalidationContextClass] method and initializes it using
// its [init()] method. Subclasses can override this method and configure
// additional properties of the invalidation context. In your implementation,
// you must call `super` first to get the context object; you can then
// configure that object and return it.
//
// [init()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/init()
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/invalidationContext(forBoundsChange:)
func (c NSCollectionViewLayout) InvalidationContextForBoundsChange(newBounds corefoundation.CGRect) INSCollectionViewLayoutInvalidationContext {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("invalidationContextForBoundsChange:"), newBounds)
	return NSCollectionViewLayoutInvalidationContextFromID(rv)
}
// Returns an invalidation context object that defines the portions of the
// layout that need to be updated.
//
// preferredAttributes: The preferred layout attributes of an element.
//
// originalAttributes: The attributes that the layout object originally suggested for the item.
//
// # Return Value
// 
// An invalidation context that describes the changes to be made. This value
// is never `nil`.
//
// # Discussion
// 
// The default implementation of this method creates an instance of the class
// returned by the [InvalidationContextClass] method and initializes it using
// its [init()] method. Subclasses can override this method and configure
// additional properties of the invalidation context. In your implementation,
// you must call `super` first to get the context object; you can then
// configure that object and return it.
//
// [init()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/init()
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/invalidationContext(forPreferredLayoutAttributes:withOriginalAttributes:)
func (c NSCollectionViewLayout) InvalidationContextForPreferredLayoutAttributesWithOriginalAttributes(preferredAttributes INSCollectionViewLayoutAttributes, originalAttributes INSCollectionViewLayoutAttributes) INSCollectionViewLayoutInvalidationContext {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("invalidationContextForPreferredLayoutAttributes:withOriginalAttributes:"), preferredAttributes, originalAttributes)
	return NSCollectionViewLayoutInvalidationContextFromID(rv)
}
// Prepares the layout object for animated changes to the collection view’s
// bounds or for the insertion or deletion of items.
//
// oldBounds: The current bounds of the collection view.
//
// # Discussion
// 
// The default implementation of this method does nothing. The collection view
// calls this method before performing any animated changes to the collection
// view’s bounds. It also calls this method before the animated insertion or
// deletion of items. Subclasses can use this method to perform any
// calculations needed to prepare for animated changes. Specifically, you
// might use this method to calculate the initial or final positions of
// inserted or deleted items so that you can return those values when asked
// for them. You can also use this method to perform custom animations. Any
// animations you create are added to the animation block used to handle the
// insertions, deletions, and bounds changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/prepare(forAnimatedBoundsChange:)
func (c NSCollectionViewLayout) PrepareForAnimatedBoundsChange(oldBounds corefoundation.CGRect) {
	objc.Send[objc.ID](c.ID, objc.Sel("prepareForAnimatedBoundsChange:"), oldBounds)
}
// Cleans up after any animated changes to the collection view’s bounds or
// after the insertion or deletion of items.
//
// # Discussion
// 
// The default implementation of this method does nothing. The collection view
// calls this method after creating the animations for changing inserting or
// deleting items or for changing the collection view’s bounds. Subclasses
// can use this method to perform any cleanup operations related to those
// changes. You can also use this method to perform custom animations. Any
// animations you create are added to the animation block used to handle the
// insertions, deletions, and bounds changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalizeAnimatedBoundsChange()
func (c NSCollectionViewLayout) FinalizeAnimatedBoundsChange() {
	objc.Send[objc.ID](c.ID, objc.Sel("finalizeAnimatedBoundsChange"))
}
// Registers a class to use when creating the layout’s decoration views.
//
// viewClass: The class to use for the decoration view. This class must conform to the
// [NSCollectionViewElement] protocol. Specify `nil` to unregister a
// previously registered class or nib file.
//
// elementKind: The string your layout uses to identify the decoration view’s type. This
// parameter must not be `nil` and must not be an empty string.
//
// # Discussion
// 
// Call this method as part of your layout object’s initialization and use
// it to register any decoration views needed for your layout. Decoration
// views are visual adornments that you include in your layouts and use to
// present the collection view’s content. For example, you might use
// decoration views to implement a dynamic background that can expand or
// shrink to match the current number of items. The layout object defines and
// owns the decoration views it uses. Decoration views do not have any ties to
// the collection view’s data source object.
// 
// After registering your decoration views, you create decoration views by
// returning an appropriate set of layout attributes from the
// [LayoutAttributesForElementsInRect] method. When you return a
// [NSCollectionViewLayoutAttributes] object configured for a decoration view,
// the collection view uses your registered nib or class information to create
// the corresponding views.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/register(_:forDecorationViewOfKind:)-44qmc
func (c NSCollectionViewLayout) RegisterClassForDecorationViewOfKind(viewClass objc.Class, elementKind NSCollectionViewDecorationElementKind) {
	objc.Send[objc.ID](c.ID, objc.Sel("registerClass:forDecorationViewOfKind:"), viewClass, objc.String(string(elementKind)))
}
// Registers a nib file to use when creating the layout’s decoration views.
//
// nib: The nib object containing the decoration view’s definition. The nib file
// must contain exactly one [NSView] object at the top level and that view
// must conform to the [NSCollectionViewElement] protocol. Specify `nil` to
// unregister a previously registered class or nib file.
//
// elementKind: The string your layout uses to identify the decoration view’s type. This
// parameter must not be `nil` and must not be an empty string.
//
// # Discussion
// 
// Call this method as part of your layout object’s initialization and use
// it to register any decoration views needed for your layout. Decoration
// views are visual adornments that you include in your layouts and use to
// present the collection view’s content. For example, you might use
// decoration views to implement a dynamic background that can expand or
// shrink to match the current number of items. The layout object defines and
// owns the decoration views it uses. Decoration views do not have any ties to
// the collection view’s data source object.
// 
// After registering your decoration views, you create decoration views by
// returning an appropriate set of layout attributes from the
// [LayoutAttributesForElementsInRect] method. When you return a
// [NSCollectionViewLayoutAttributes] object configured for a decoration view,
// the collection view uses your registered nib or class information to create
// the corresponding views.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/register(_:forDecorationViewOfKind:)-7z7uf
func (c NSCollectionViewLayout) RegisterNibForDecorationViewOfKind(nib INSNib, elementKind NSCollectionViewDecorationElementKind) {
	objc.Send[objc.ID](c.ID, objc.Sel("registerNib:forDecorationViewOfKind:"), nib, objc.String(string(elementKind)))
}
// Prepares the layout object to be installed in the collection view.
//
// oldLayout: The layout object installed in the collection view at the beginning of the
// transition. You might use this object to retrieve the current layout
// attributes for items and views.
//
// # Discussion
// 
// Prior to transitioning to a new layout object, the collection view calls
// this method on the new layout object to give it time to perform any initial
// calculations.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/prepareForTransition(from:)
func (c NSCollectionViewLayout) PrepareForTransitionFromLayout(oldLayout INSCollectionViewLayout) {
	objc.Send[objc.ID](c.ID, objc.Sel("prepareForTransitionFromLayout:"), oldLayout)
}
// Prepares the layout object to be uninstalled from the collection view.
//
// newLayout: The layout object to install in the collection view at the end of the
// transition.
//
// # Discussion
// 
// Prior to transitioning to a new layout object, the collection view calls
// this method on the old layout object to give it time to perform any cleanup
// operations.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/prepareForTransition(to:)
func (c NSCollectionViewLayout) PrepareForTransitionToLayout(newLayout INSCollectionViewLayout) {
	objc.Send[objc.ID](c.ID, objc.Sel("prepareForTransitionToLayout:"), newLayout)
}
// Performs any final steps related to a layout transition before the
// transition animations actually occur.
//
// # Discussion
// 
// After it has gathered all of the layout attributes it needs, the collection
// view calls this method on both the new and old layout objects to give them
// time to perform any additional operations. Use this method to clean up any
// data structures or caches related to the transition that your layout object
// no longer needs.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/finalizeLayoutTransition()
func (c NSCollectionViewLayout) FinalizeLayoutTransition() {
	objc.Send[objc.ID](c.ID, objc.Sel("finalizeLayoutTransition"))
}
func (c NSCollectionViewLayout) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The collection view object currently using this layout.
//
// # Discussion
// 
// When you assign a layout object to a collection view, the collection view
// automatically updates this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/collectionView
func (c NSCollectionViewLayout) CollectionView() INSCollectionView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("collectionView"))
	return NSCollectionViewFromID(objc.ID(rv))
}
// The width and height of the collection view’s contents.
//
// # Discussion
// 
// This property contains the width and height of all of the collection
// view’s contents, not just the content that is currently visible. The
// collection view uses this information to configure its scroll view.
// 
// When creating custom layouts, you must reimplement this property and
// provide the size of the collection view’s contents. It is recommended
// that you cache the content size and adjust the value when the layout
// changes or when items are added and removed.
// 
// The default value in this property is [NSZeroSize].
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/collectionViewContentSize
func (c NSCollectionViewLayout) CollectionViewContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("collectionViewContentSize"))
	return corefoundation.CGSize(rv)
}
// The layout object used to organize the collection view’s content.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionview/collectionviewlayout
func (c NSCollectionViewLayout) CollectionViewLayout() INSCollectionViewLayout {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("collectionViewLayout"))
	return NSCollectionViewLayoutFromID(objc.ID(rv))
}
func (c NSCollectionViewLayout) SetCollectionViewLayout(value INSCollectionViewLayout) {
	objc.Send[struct{}](c.ID, objc.Sel("setCollectionViewLayout:"), value)
}

// Returns the class to use for layout attribute objects
//
// # Return Value
// 
// The class to use for layout attribute objects.
// 
// # Discussion
// 
// Override this method if you define a custom
// [NSCollectionViewLayoutAttributes] subclass for managing layout-related
// attributes. In your implementation, return the class object for your custom
// subclass.
// 
// You can call this method as needed to create new layout objects. A typical
// usage of this method is as follows:
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/layoutAttributesClass
func (_NSCollectionViewLayoutClass NSCollectionViewLayoutClass) LayoutAttributesClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSCollectionViewLayoutClass.class), objc.Sel("layoutAttributesClass"))
	return rv
}
// Returns the class to use when creating an invalidation context object for
// the layout.
//
// # Return Value
// 
// A custom class that descends from
// [NSCollectionViewLayoutInvalidationContext].
// 
// # Discussion
// 
// If you define a custom invalidation context class to store information
// related to your layout, override this method and use it to return your
// custom subclass. Methods of this class that create invalidation contexts
// automatically create instances of the class you provide, initializing those
// instances using its [init()] method.
//
// [init()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/init()
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/invalidationContextClass
func (_NSCollectionViewLayoutClass NSCollectionViewLayoutClass) InvalidationContextClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSCollectionViewLayoutClass.class), objc.Sel("invalidationContextClass"))
	return rv
}

