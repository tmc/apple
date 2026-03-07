// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSCollectionViewFlowLayout] class.
var (
	_NSCollectionViewFlowLayoutClass     NSCollectionViewFlowLayoutClass
	_NSCollectionViewFlowLayoutClassOnce sync.Once
)

func getNSCollectionViewFlowLayoutClass() NSCollectionViewFlowLayoutClass {
	_NSCollectionViewFlowLayoutClassOnce.Do(func() {
		_NSCollectionViewFlowLayoutClass = NSCollectionViewFlowLayoutClass{class: objc.GetClass("NSCollectionViewFlowLayout")}
	})
	return _NSCollectionViewFlowLayoutClass
}

// GetNSCollectionViewFlowLayoutClass returns the class object for NSCollectionViewFlowLayout.
func GetNSCollectionViewFlowLayoutClass() NSCollectionViewFlowLayoutClass {
	return getNSCollectionViewFlowLayoutClass()
}

type NSCollectionViewFlowLayoutClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewFlowLayoutClass) Alloc() NSCollectionViewFlowLayout {
	rv := objc.Send[NSCollectionViewFlowLayout](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A layout that organizes items into a flexible and configurable arrangement.
//
// # Overview
// 
// In a flow layout, the first item is positioned in the top-left corner and
// other items are laid out either horizontally or vertically based on the
// scroll direction, which is configurable. Items may be the same size or
// different sizes, and you may use the flow layout object or the collection
// view’s delegate object to specify the size of items and the spacing
// around them. The flow layout also lets you specify custom header and footer
// views for each section.
// 
// You can use an [NSCollectionViewFlowLayout] object as-is or subclass it to
// modify more aspects of the layout behavior. There are several ways to
// customize the basic layout behavior that do not require subclassing. For
// example, you can use a delegate object to change the size and spacing of
// items dynamically. Subclassing is appropriate for more advanced layout
// changes, such as adding supplementary views or decoration views, supporting
// custom layout attributes, or customizing the layout animations when
// inserting or deleting items.
// 
// # Configuring a Collection View to Use a Flow Layout
// 
// You can configure a collection view to use a flow layout object
// programmatically or at design time:
// 
// - At design time, set the Layout attribute of your collection view to Flow.
// - Create an [NSCollectionViewFlowLayout] object programmatically and assign
// it to the collection view’s [CollectionViewLayout] property.
// 
// Normally, you specify the size of items and their spacing using the
// properties of this class. If you want to customize the values for your
// items, implement the methods of the [NSCollectionViewDelegateFlowLayout]
// protocol in the object assigned as the collection view’s [NSCollectionViewFlowLayout.Delegate]. The
// delegate methods let you adjust layout information dynamically and change
// the values later as needed. If you do not provide a delegate, the flow
// layout object uses the same values for all items.
// 
// For more information about customizing the layout dynamically using a
// delegate, see [NSCollectionViewDelegateFlowLayout].
// 
// # Adding Header and Footer Views to Sections
// 
// The flow layout supports custom header and footer views for each section,
// displaying them as supplementary views in the layout. [Figure 1] shows how
// header and footer views are presented in your layout. The header view
// appears above the items in a section and the footer view appears below
// those items.
// 
// [media-1965629]
// 
// The size of the header and footer views is configurable using the
// [NSCollectionViewFlowLayout.HeaderReferenceSize] and [NSCollectionViewFlowLayout.FooterReferenceSize] properties of this class or
// using the delegate object. For a vertically scrolling flow layout, the
// header and footer views span the width of the collection view and the
// height is based on the values you specify. For a horizontal layout, the
// footers span the height of the collection view and the width is
// configurable.
// 
// To add header or footer views to your flow layout, you must do the
// following:
// 
// - Register your header and footer views using the
// [RegisterClassForSupplementaryViewOfKindWithIdentifier] or
// [RegisterNibForSupplementaryViewOfKindWithIdentifier] method. - Implement
// the [CollectionViewViewForSupplementaryElementOfKindAtIndexPath] method in
// your collection view’s data source object.
// 
// When registering your header view, specify [elementKindSectionHeader] for
// the kind string. When registering your footer view, specify the
// [elementKindSectionFooter] string. You also pass those strings to the
// [SupplementaryViewOfKindWithIdentifierForIndexPath] method when creating
// new views in your data source object.
// 
// You configure the contents of your header and footer views in the same way
// you configure items. Use your app data to configure the content of the
// supplementary view and its subviews. When you want to update the content in
// a header or footer, call the [ReloadSections] method and force the
// collection view to update the section, including its headers and footers.
// Reloading the section is safer than maintaining references to the views
// themselves because of how the collection view recycles views. If you do
// want to update views directly, use the methods of the
// [NSCollectionViewDelegate] protocol to track when your views are added and
// removed from the collection view.
// 
// # Understanding How the Flow Layout is Generated
// 
// The flow layout builds its grid dynamically, using the item size and
// spacing information to place items sequentially into rows (for a vertically
// scrolling layout) or columns (for a horizontally scrolling layout). The
// number of items displayed in each row or column is determined by the
// collection view’s size and the size of the items in that row. When items
// have a uniform size, the number of items in each row is the same, but when
// items are different sizes, some rows may contain more or fewer items.
// [Figure 2] illustrates how the flow layout places items in a row until
// there is no more space and then begins a new row and continues the layout
// process.
// 
// [media-1965630]
// 
// The flow layout follows a specific set of steps to determine the size of
// each item. Whenever possible, the flow layout uses previously calculated
// information. When that information is not available, it falls back on other
// techniques to retrieve the size of the item. Specifically, the flow layout
// takes the following steps, stopping when it acquires a valid item size:
// 
// - Get the size of the item from the already computed layout attributes. -
// Call the [CollectionViewLayoutSizeForItemAtIndexPath] method of the
// delegate to get the item size. - Use the [NSCollectionViewFlowLayout.EstimatedItemSize] property, if
// it is not set to [NSCollectionViewFlowLayout.NSZeroSize]. - Use the [NSCollectionViewFlowLayout.ItemSize] property to get the
// size.
// 
// Individual spacing between items and between different lines of items is
// controlled by the properties of this class and the delegate. For line
// spacing, the largest item in the line defines the line’s height and
// thereby controls the spacing between that line and other lines. The minimum
// inter-item spacing controls only how many items are placed on the line, and
// does not set the spacing between items.
// 
// # Subclassing Notes
// 
// When you want to do more than specify the size or spacing of items,
// subclass [NSCollectionViewFlowLayout] and override the appropriate methods.
// When subclassing, the best approach is to take advantage of the built-in
// flow layout behavior and then make tweaks as needed.
// 
// # Adding New Supplementary or Decoration Views to the Layout
// 
// The standard flow layout object supports only header and footer
// supplementary views for each section. To support additional supplementary
// views, or to add decoration views, you must override the following methods:
// 
// - [LayoutAttributesForElementsInRect] (Required) -
// [LayoutAttributesForItemAtIndexPath] (Required) -
// [LayoutAttributesForSupplementaryViewOfKindAtIndexPath] (Required only if
// you are adding supplementary views) -
// [LayoutAttributesForDecorationViewOfKindAtIndexPath] (Required only if you
// are adding decoration views)
// 
// Your implementations of these methods should adjust the position of items
// and views to accommodate your custom content. For each of your
// implementations, call `super` first so that you can modify the attributes
// returned by the default flow layout behavior before adding any custom
// layout attributes.
// 
// # Tweaking the Layout Attributes
// 
// To tweak the flow layout algorithm, override the
// [LayoutAttributesForElementsInRect] method and any other methods that
// return layout attributes that you need to modify. In each method, call
// `super` first and then modify the attributes returned by the default flow
// layout behavior.
// 
// # Changing the Initial or Final Attributes of Elements
// 
// To customize the insertion or deletion animations performed by the layout,
// override some or all of the following methods:
// 
// - [NSCollectionViewFlowLayout.InitialLayoutAttributesForAppearingItemAtIndexPath] -
// [NSCollectionViewFlowLayout.InitialLayoutAttributesForAppearingSupplementaryElementOfKindAtIndexPath]
// - [NSCollectionViewFlowLayout.InitialLayoutAttributesForAppearingDecorationElementOfKindAtIndexPath] -
// [FinalLayoutAttributesForDisappearingItemAtIndexPath] -
// [FinalLayoutAttributesForDisappearingSupplementaryElementOfKindAtIndexPath]
// - [FinalLayoutAttributesForDisappearingDecorationElementOfKindAtIndexPath]
// 
// In your custom implementations, specify the layout attributes for each item
// or view being inserted or deleted. The flow layout handles the creation of
// the appropriate animations, using the initial or final attributes that you
// provide. It is also recommended that you override the
// [PrepareForCollectionViewUpdates] and [FinalizeCollectionViewUpdates]
// methods to track the insertions and deletions.
// 
// # Supporting Custom Layout Attributes
// 
// Subclassing is required if you want the flow layout to manage attributes
// other than the size and visibility of items and views. If you subclass
// [NSCollectionViewLayoutAttributes] and define additional layout attributes,
// you must also subclass [NSCollectionViewFlowLayout] to provide values for
// those attributes . In your flow layout subclass, override the following:
// 
// - [NSCollectionViewFlowLayout.LayoutAttributesClass] - [LayoutAttributesForElementsInRect] -
// [LayoutAttributesForItemAtIndexPath] - All other methods that return layout
// attributes
// 
// The implementations of your custom methods should set the values of any
// layout attributes that you define. Call `super` first to retrieve the
// default layout attributes objects and then modify the returned objects by
// adding your custom data.
// 
// For more information about defining custom layout attributes, see
// [NSCollectionViewLayoutAttributes].
//
// [elementKindSectionFooter]: https://developer.apple.com/documentation/AppKit/NSCollectionView/elementKindSectionFooter
// [elementKindSectionHeader]: https://developer.apple.com/documentation/AppKit/NSCollectionView/elementKindSectionHeader
//
// # Configuring the Scroll Direction
//
//   - [NSCollectionViewFlowLayout.ScrollDirection]: The scroll direction of the layout.
//   - [NSCollectionViewFlowLayout.SetScrollDirection]
//
// # Configuring the Item Spacing
//
//   - [NSCollectionViewFlowLayout.MinimumLineSpacing]: The minimum spacing (in points) to use between rows or columns.
//   - [NSCollectionViewFlowLayout.SetMinimumLineSpacing]
//   - [NSCollectionViewFlowLayout.MinimumInteritemSpacing]: The minimum spacing (in points) to use between items in the same row or column.
//   - [NSCollectionViewFlowLayout.SetMinimumInteritemSpacing]
//   - [NSCollectionViewFlowLayout.EstimatedItemSize]: The estimated size of items in the collection view.
//   - [NSCollectionViewFlowLayout.SetEstimatedItemSize]
//   - [NSCollectionViewFlowLayout.ItemSize]: The default size to use for items.
//   - [NSCollectionViewFlowLayout.SetItemSize]
//   - [NSCollectionViewFlowLayout.SectionInset]: The margins used to lay out content in a section.
//   - [NSCollectionViewFlowLayout.SetSectionInset]
//
// # Configuring the Supplementary Views
//
//   - [NSCollectionViewFlowLayout.HeaderReferenceSize]: The default size to use for section headers.
//   - [NSCollectionViewFlowLayout.SetHeaderReferenceSize]
//   - [NSCollectionViewFlowLayout.FooterReferenceSize]: The default size to use for section footers.
//   - [NSCollectionViewFlowLayout.SetFooterReferenceSize]
//
// # Instance Properties
//
//   - [NSCollectionViewFlowLayout.SectionFootersPinToVisibleBounds]
//   - [NSCollectionViewFlowLayout.SetSectionFootersPinToVisibleBounds]
//   - [NSCollectionViewFlowLayout.SectionHeadersPinToVisibleBounds]
//   - [NSCollectionViewFlowLayout.SetSectionHeadersPinToVisibleBounds]
//
// # Instance Methods
//
//   - [NSCollectionViewFlowLayout.CollapseSectionAtIndex]
//   - [NSCollectionViewFlowLayout.ExpandSectionAtIndex]
//   - [NSCollectionViewFlowLayout.SectionAtIndexIsCollapsed]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout
type NSCollectionViewFlowLayout struct {
	NSCollectionViewLayout
}

// NSCollectionViewFlowLayoutFromID constructs a [NSCollectionViewFlowLayout] from an objc.ID.
//
// A layout that organizes items into a flexible and configurable arrangement.
func NSCollectionViewFlowLayoutFromID(id objc.ID) NSCollectionViewFlowLayout {
	return NSCollectionViewFlowLayout{NSCollectionViewLayout: NSCollectionViewLayoutFromID(id)}
}
// NOTE: NSCollectionViewFlowLayout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCollectionViewFlowLayout] class.
//
// # Configuring the Scroll Direction
//
//   - [INSCollectionViewFlowLayout.ScrollDirection]: The scroll direction of the layout.
//   - [INSCollectionViewFlowLayout.SetScrollDirection]
//
// # Configuring the Item Spacing
//
//   - [INSCollectionViewFlowLayout.MinimumLineSpacing]: The minimum spacing (in points) to use between rows or columns.
//   - [INSCollectionViewFlowLayout.SetMinimumLineSpacing]
//   - [INSCollectionViewFlowLayout.MinimumInteritemSpacing]: The minimum spacing (in points) to use between items in the same row or column.
//   - [INSCollectionViewFlowLayout.SetMinimumInteritemSpacing]
//   - [INSCollectionViewFlowLayout.EstimatedItemSize]: The estimated size of items in the collection view.
//   - [INSCollectionViewFlowLayout.SetEstimatedItemSize]
//   - [INSCollectionViewFlowLayout.ItemSize]: The default size to use for items.
//   - [INSCollectionViewFlowLayout.SetItemSize]
//   - [INSCollectionViewFlowLayout.SectionInset]: The margins used to lay out content in a section.
//   - [INSCollectionViewFlowLayout.SetSectionInset]
//
// # Configuring the Supplementary Views
//
//   - [INSCollectionViewFlowLayout.HeaderReferenceSize]: The default size to use for section headers.
//   - [INSCollectionViewFlowLayout.SetHeaderReferenceSize]
//   - [INSCollectionViewFlowLayout.FooterReferenceSize]: The default size to use for section footers.
//   - [INSCollectionViewFlowLayout.SetFooterReferenceSize]
//
// # Instance Properties
//
//   - [INSCollectionViewFlowLayout.SectionFootersPinToVisibleBounds]
//   - [INSCollectionViewFlowLayout.SetSectionFootersPinToVisibleBounds]
//   - [INSCollectionViewFlowLayout.SectionHeadersPinToVisibleBounds]
//   - [INSCollectionViewFlowLayout.SetSectionHeadersPinToVisibleBounds]
//
// # Instance Methods
//
//   - [INSCollectionViewFlowLayout.CollapseSectionAtIndex]
//   - [INSCollectionViewFlowLayout.ExpandSectionAtIndex]
//   - [INSCollectionViewFlowLayout.SectionAtIndexIsCollapsed]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout
type INSCollectionViewFlowLayout interface {
	INSCollectionViewLayout

	// Topic: Configuring the Scroll Direction

	// The scroll direction of the layout.
	ScrollDirection() NSCollectionViewScrollDirection
	SetScrollDirection(value NSCollectionViewScrollDirection)

	// Topic: Configuring the Item Spacing

	// The minimum spacing (in points) to use between rows or columns.
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
	// The minimum spacing (in points) to use between items in the same row or column.
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	// The estimated size of items in the collection view.
	EstimatedItemSize() corefoundation.CGSize
	SetEstimatedItemSize(value corefoundation.CGSize)
	// The default size to use for items.
	ItemSize() corefoundation.CGSize
	SetItemSize(value corefoundation.CGSize)
	// The margins used to lay out content in a section.
	SectionInset() foundation.NSEdgeInsets
	SetSectionInset(value foundation.NSEdgeInsets)

	// Topic: Configuring the Supplementary Views

	// The default size to use for section headers.
	HeaderReferenceSize() corefoundation.CGSize
	SetHeaderReferenceSize(value corefoundation.CGSize)
	// The default size to use for section footers.
	FooterReferenceSize() corefoundation.CGSize
	SetFooterReferenceSize(value corefoundation.CGSize)

	// Topic: Instance Properties

	SectionFootersPinToVisibleBounds() bool
	SetSectionFootersPinToVisibleBounds(value bool)
	SectionHeadersPinToVisibleBounds() bool
	SetSectionHeadersPinToVisibleBounds(value bool)

	// Topic: Instance Methods

	CollapseSectionAtIndex(sectionIndex uint)
	ExpandSectionAtIndex(sectionIndex uint)
	SectionAtIndexIsCollapsed(sectionIndex uint) bool

	// The collection view’s delegate object.
	Delegate() NSCollectionViewDelegate
	SetDelegate(value NSCollectionViewDelegate)
	// An `NSSize` structure set to `0` in both dimensions.
	NSZeroSize() corefoundation.CGSize
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (c NSCollectionViewFlowLayout) Init() NSCollectionViewFlowLayout {
	rv := objc.Send[NSCollectionViewFlowLayout](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewFlowLayout) Autorelease() NSCollectionViewFlowLayout {
	rv := objc.Send[NSCollectionViewFlowLayout](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewFlowLayout creates a new NSCollectionViewFlowLayout instance.
func NewNSCollectionViewFlowLayout() NSCollectionViewFlowLayout {
	class := getNSCollectionViewFlowLayoutClass()
	rv := objc.Send[NSCollectionViewFlowLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}











//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/collapseSection(at:)
func (c NSCollectionViewFlowLayout) CollapseSectionAtIndex(sectionIndex uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("collapseSectionAtIndex:"), sectionIndex)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/expandSection(at:)
func (c NSCollectionViewFlowLayout) ExpandSectionAtIndex(sectionIndex uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("expandSectionAtIndex:"), sectionIndex)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/section(atIndexIsCollapsed:)
func (c NSCollectionViewFlowLayout) SectionAtIndexIsCollapsed(sectionIndex uint) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("sectionAtIndexIsCollapsed:"), sectionIndex)
	return rv
}
func (c NSCollectionViewFlowLayout) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The scroll direction of the layout.
//
// # Discussion
// 
// The flow layout scrolls along one axis only, either horizontally or
// vertically. When the scroll direction is
// [NSCollectionView.ScrollDirection.vertical], the width of the content never
// exceeds the width of the collection view itself but the height grows as
// needed to accommodate the current items. When the scroll direction is
// [NSCollectionView.ScrollDirection.horizontal], the height never exceeds the
// height of the collection view but the width grows as needed.
// 
// The default value of this property is
// [NSCollectionView.ScrollDirection.vertical].
//
// [NSCollectionView.ScrollDirection.horizontal]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollDirection/horizontal
// [NSCollectionView.ScrollDirection.vertical]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollDirection/vertical
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/scrollDirection
func (c NSCollectionViewFlowLayout) ScrollDirection() NSCollectionViewScrollDirection {
	rv := objc.Send[NSCollectionViewScrollDirection](c.ID, objc.Sel("scrollDirection"))
	return NSCollectionViewScrollDirection(rv)
}
func (c NSCollectionViewFlowLayout) SetScrollDirection(value NSCollectionViewScrollDirection) {
	objc.Send[struct{}](c.ID, objc.Sel("setScrollDirection:"), value)
}



// The minimum spacing (in points) to use between rows or columns.
//
// # Discussion
// 
// If the delegate does not implement the
// [CollectionViewLayoutMinimumLineSpacingForSectionAtIndex] method, the flow
// layout object uses the value of this property to set the spacing between
// rows or columns.
// 
// For a vertically scrolling layout, the value represents the minimum spacing
// between successive rows. For a horizontally scrolling layout, the value
// represents the minimum spacing between successive columns. This spacing is
// not applied to the space between the header view and the first line or
// between the last line and the footer view. [Figure 1] shows how the line
// spacing is applied to rows of unevenly sized items, illustrating how the
// actual spacing between individual items may be greater than the minimum
// value.
// 
// [media-1965631]
// 
// The default value of this property is `10.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/minimumLineSpacing
func (c NSCollectionViewFlowLayout) MinimumLineSpacing() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("minimumLineSpacing"))
	return rv
}
func (c NSCollectionViewFlowLayout) SetMinimumLineSpacing(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinimumLineSpacing:"), value)
}



// The minimum spacing (in points) to use between items in the same row or
// column.
//
// # Discussion
// 
// If the delegate does not implement the
// [CollectionViewLayoutMinimumInteritemSpacingForSectionAtIndex] method, the
// flow layout object uses the value of this property to set the spacing
// between items in the same line.
// 
// For a vertically scrolling layout, the value represents the minimum spacing
// between items in the same row. For a horizontally scrolling layout, the
// value represents the minimum spacing between items in the same column. The
// layout object uses this spacing only to compute how many items can fit in a
// single row or column. The actual spacing may be increased after the number
// of items has been determined, as illustrated in [Figure 1].
// 
// [media-1965632]
// 
// The default value of this property is `10.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/minimumInteritemSpacing
func (c NSCollectionViewFlowLayout) MinimumInteritemSpacing() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("minimumInteritemSpacing"))
	return rv
}
func (c NSCollectionViewFlowLayout) SetMinimumInteritemSpacing(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinimumInteritemSpacing:"), value)
}



// The estimated size of items in the collection view.
//
// # Discussion
// 
// Providing an estimated item size lets the collection view defer some of the
// calculations needed to determine the size of its content, which can improve
// performance. Instead of explicitly computing the size of each item, the
// flow layout assumes that offscreen items have the estimated size. The
// estimated size is used only until an actual value is calculated. The
// default value of this property is [NSZeroSize].
// 
// If the value of this property is not [NSZeroSize], the flow layout uses the
// estimated size you specified. If all of your items actually have the same
// size, use the [ItemSize] property to set their size and set this property
// to [NSZeroSize]. For more information about how item sizes are determined,
// see [NSCollectionViewFlowLayout].
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/estimatedItemSize
func (c NSCollectionViewFlowLayout) EstimatedItemSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("estimatedItemSize"))
	return corefoundation.CGSize(rv)
}
func (c NSCollectionViewFlowLayout) SetEstimatedItemSize(value corefoundation.CGSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setEstimatedItemSize:"), value)
}



// The default size to use for items.
//
// # Discussion
// 
// This property contains the default size of items. If you do not provide an
// estimated size or implement the
// [CollectionViewLayoutSizeForItemAtIndexPath] method in your delegate, the
// flow layout uses this value for the size of each item. All items are set to
// the same size. This value applies only to items and not to supplementary
// views.
// 
// The default value of this property is (`50.0`, `50.0`). For more
// information about how item sizes are determined, see
// [NSCollectionViewFlowLayout].
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/itemSize
func (c NSCollectionViewFlowLayout) ItemSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("itemSize"))
	return corefoundation.CGSize(rv)
}
func (c NSCollectionViewFlowLayout) SetItemSize(value corefoundation.CGSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setItemSize:"), value)
}



// The margins used to lay out content in a section.
//
// # Discussion
// 
// If the delegate does not implement the
// [CollectionViewLayoutInsetForSectionAtIndex] method, the flow layout object
// uses the value of this property to set the margins for each section.
// 
// Section insets reflect the spacing at the outer edges of the section. The
// margins affect the positioning of the header view, the minimum space on
// either side of each line of items, and the distance from the last line to
// the footer view, as shown in [Figure 1]. The margin insets do not affect
// the size of the header and footer views in the nonscrolling direction.
// 
// [media-1965633]
// 
// The default insets are all set to `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionInset
func (c NSCollectionViewFlowLayout) SectionInset() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](c.ID, objc.Sel("sectionInset"))
	return foundation.NSEdgeInsets(rv)
}
func (c NSCollectionViewFlowLayout) SetSectionInset(value foundation.NSEdgeInsets) {
	objc.Send[struct{}](c.ID, objc.Sel("setSectionInset:"), value)
}



// The default size to use for section headers.
//
// # Discussion
// 
// If the delegate does not implement the
// [CollectionViewLayoutReferenceSizeForHeaderInSection] method, the flow
// layout object uses the value of this property as the header size.
// 
// The layout object uses only the value that is appropriate for the current
// scrolling direction. In other words, the layout object uses only the height
// value when the content scrolls vertically, setting the width of the header
// to the width of the collection view. Similarly, the layout object uses only
// the width value when the content scrolls horizontally, setting the
// header’s height to the height of the collection view. If the size value
// for the appropriate dimension is `0`, the layout object omits the header
// entirely.
// 
// The default value of this property is [NSZeroSize].
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/headerReferenceSize
func (c NSCollectionViewFlowLayout) HeaderReferenceSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("headerReferenceSize"))
	return corefoundation.CGSize(rv)
}
func (c NSCollectionViewFlowLayout) SetHeaderReferenceSize(value corefoundation.CGSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setHeaderReferenceSize:"), value)
}



// The default size to use for section footers.
//
// # Discussion
// 
// If the delegate does not implement the
// [CollectionViewLayoutReferenceSizeForFooterInSection] method, the flow
// layout object uses the value of this property as the footer size.
// 
// The layout object uses only the value that is appropriate for the current
// scrolling direction. In other words, the layout object uses only the height
// value when the content scrolls vertically, setting the width of the footer
// to the width of the collection view. Similarly, the layout object uses only
// the width value when the content scrolls horizontally, setting the
// footer’s height to the height of the collection view. If the size value
// for the appropriate dimension is `0`, the layout object omits the footer
// entirely.
// 
// The default value of this property is [NSZeroSize].
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/footerReferenceSize
func (c NSCollectionViewFlowLayout) FooterReferenceSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("footerReferenceSize"))
	return corefoundation.CGSize(rv)
}
func (c NSCollectionViewFlowLayout) SetFooterReferenceSize(value corefoundation.CGSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setFooterReferenceSize:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionFootersPinToVisibleBounds
func (c NSCollectionViewFlowLayout) SectionFootersPinToVisibleBounds() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("sectionFootersPinToVisibleBounds"))
	return rv
}
func (c NSCollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSectionFootersPinToVisibleBounds:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout/sectionHeadersPinToVisibleBounds
func (c NSCollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("sectionHeadersPinToVisibleBounds"))
	return rv
}
func (c NSCollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSectionHeadersPinToVisibleBounds:"), value)
}



// The collection view’s delegate object.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionview/delegate
func (c NSCollectionViewFlowLayout) Delegate() NSCollectionViewDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return NSCollectionViewDelegateObjectFromID(rv)
}
func (c NSCollectionViewFlowLayout) SetDelegate(value NSCollectionViewDelegate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelegate:"), value)
}



// An `NSSize` structure set to `0` in both dimensions.
//
// See: https://developer.apple.com/documentation/Foundation/NSZeroSize
func (c NSCollectionViewFlowLayout) NSZeroSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("NSZeroSize"))
	return corefoundation.CGSize(rv)
}







// A supplementary view that acts as a footer for a given section.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionview/elementkindsectionfooter
func (_NSCollectionViewFlowLayoutClass NSCollectionViewFlowLayoutClass) ElementKindSectionFooter() string {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionViewFlowLayoutClass.class), objc.Sel("NSCollectionElementKindSectionFooter"))
	return foundation.NSStringFromID(rv).String()
}



// A supplementary view that acts as a header for a given section.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionview/elementkindsectionheader
func (_NSCollectionViewFlowLayoutClass NSCollectionViewFlowLayoutClass) ElementKindSectionHeader() string {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionViewFlowLayoutClass.class), objc.Sel("NSCollectionElementKindSectionHeader"))
	return foundation.NSStringFromID(rv).String()
}






















