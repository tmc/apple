// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionViewLayoutAttributes] class.
var (
	_NSCollectionViewLayoutAttributesClass     NSCollectionViewLayoutAttributesClass
	_NSCollectionViewLayoutAttributesClassOnce sync.Once
)

func getNSCollectionViewLayoutAttributesClass() NSCollectionViewLayoutAttributesClass {
	_NSCollectionViewLayoutAttributesClassOnce.Do(func() {
		_NSCollectionViewLayoutAttributesClass = NSCollectionViewLayoutAttributesClass{class: objc.GetClass("NSCollectionViewLayoutAttributes")}
	})
	return _NSCollectionViewLayoutAttributesClass
}

// GetNSCollectionViewLayoutAttributesClass returns the class object for NSCollectionViewLayoutAttributes.
func GetNSCollectionViewLayoutAttributesClass() NSCollectionViewLayoutAttributesClass {
	return getNSCollectionViewLayoutAttributesClass()
}

type NSCollectionViewLayoutAttributesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCollectionViewLayoutAttributesClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewLayoutAttributesClass) Alloc() NSCollectionViewLayoutAttributes {
	rv := objc.Send[NSCollectionViewLayoutAttributes](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains layout-related attributes for an element in a
// collection view.
//
// # Overview
//
// During the layout, the layout object creates instances of
// [NSCollectionViewLayoutAttributes] for each element displayed in the
// collection view. The layout attributes describe the position of an element
// and other information such as its alpha and position on the z axis. The
// collection view later applies the layout attributes to the onscreen
// elements.
//
// The only time you interact with layout attribute objects is when you
// implement a custom layout, and the interactions are straightforward. When
// asked for layout attributes for a specific element, your layout object uses
// the methods of this class to create an appropriate instance of the class
// based on the type of the requested element. It then configures the
// properties of the object and returns it to the requester.
//
// # Subclassing Notes
//
// If you implement a custom layout object and your layout object requires
// additional attributes, you can subclass [NSCollectionViewLayoutAttributes]
// and add custom properties to your subclass. In your subclass, be sure to do
// the following:
//
// - Provide an [init()] method with no parameters to initialize your
// subclass. - Implement support for the [NSCopying] protocol. The collection
// view caches layout attribute objects for later use. - Override the
// inherited [isEqual(_:)] method to perform any relevant equality checks.
//
// Supporting equality checks is important because of how the collection view
// manages layout attributes. As an optimization, the collection view applies
// layout attributes only when they change. When the layout object returns a
// layout attributes object, the collection view checks to see if the new
// attributes are equal to any cached attributes. Therefore, if you want to
// include any new properties in the equality check, you must override the
// [isEqual(_:)] method.
//
// In addition to defining your [NSCollectionViewLayoutAttributes] subclass,
// override the [NSCollectionViewLayoutAttributes.LayoutAttributesClass] method of your layout object. That
// method is a funnel point for creating new layout attribute objects.
// Returning your custom class from that method ensures that the correct class
// is instantiated.
//
// # Identifying the Element
//
//   - [NSCollectionViewLayoutAttributes.RepresentedElementCategory]: The type of the element.
//   - [NSCollectionViewLayoutAttributes.IndexPath]: The index path of the element.
//   - [NSCollectionViewLayoutAttributes.SetIndexPath]
//   - [NSCollectionViewLayoutAttributes.RepresentedElementKind]: The identifier for specific elements of your collection view interface.
//
// # Accessing the Layout Attributes
//
//   - [NSCollectionViewLayoutAttributes.Frame]: The frame rectangle of the element.
//   - [NSCollectionViewLayoutAttributes.SetFrame]
//   - [NSCollectionViewLayoutAttributes.Size]: The size of the element.
//   - [NSCollectionViewLayoutAttributes.SetSize]
//   - [NSCollectionViewLayoutAttributes.Alpha]: The transparency of the element.
//   - [NSCollectionViewLayoutAttributes.SetAlpha]
//   - [NSCollectionViewLayoutAttributes.Hidden]: A Boolean value indicating whether the element is hidden.
//   - [NSCollectionViewLayoutAttributes.SetHidden]
//   - [NSCollectionViewLayoutAttributes.ZIndex]: The element’s position on the z axis.
//   - [NSCollectionViewLayoutAttributes.SetZIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes
//
// [NSCopying]: https://developer.apple.com/documentation/Foundation/NSCopying
// [init()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/init()
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
type NSCollectionViewLayoutAttributes struct {
	objectivec.Object
}

// NSCollectionViewLayoutAttributesFromID constructs a [NSCollectionViewLayoutAttributes] from an objc.ID.
//
// An object that contains layout-related attributes for an element in a
// collection view.
func NSCollectionViewLayoutAttributesFromID(id objc.ID) NSCollectionViewLayoutAttributes {
	return NSCollectionViewLayoutAttributes{objectivec.Object{ID: id}}
}

// NOTE: NSCollectionViewLayoutAttributes adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionViewLayoutAttributes] class.
//
// # Identifying the Element
//
//   - [INSCollectionViewLayoutAttributes.RepresentedElementCategory]: The type of the element.
//   - [INSCollectionViewLayoutAttributes.IndexPath]: The index path of the element.
//   - [INSCollectionViewLayoutAttributes.SetIndexPath]
//   - [INSCollectionViewLayoutAttributes.RepresentedElementKind]: The identifier for specific elements of your collection view interface.
//
// # Accessing the Layout Attributes
//
//   - [INSCollectionViewLayoutAttributes.Frame]: The frame rectangle of the element.
//   - [INSCollectionViewLayoutAttributes.SetFrame]
//   - [INSCollectionViewLayoutAttributes.Size]: The size of the element.
//   - [INSCollectionViewLayoutAttributes.SetSize]
//   - [INSCollectionViewLayoutAttributes.Alpha]: The transparency of the element.
//   - [INSCollectionViewLayoutAttributes.SetAlpha]
//   - [INSCollectionViewLayoutAttributes.Hidden]: A Boolean value indicating whether the element is hidden.
//   - [INSCollectionViewLayoutAttributes.SetHidden]
//   - [INSCollectionViewLayoutAttributes.ZIndex]: The element’s position on the z axis.
//   - [INSCollectionViewLayoutAttributes.SetZIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes
type INSCollectionViewLayoutAttributes interface {
	objectivec.IObject

	// Topic: Identifying the Element

	// The type of the element.
	RepresentedElementCategory() NSCollectionElementCategory
	// The index path of the element.
	IndexPath() objc.ID
	SetIndexPath(value objc.ID)
	// The identifier for specific elements of your collection view interface.
	RepresentedElementKind() string

	// Topic: Accessing the Layout Attributes

	// The frame rectangle of the element.
	Frame() corefoundation.CGRect
	SetFrame(value corefoundation.CGRect)
	// The size of the element.
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)
	// The transparency of the element.
	Alpha() float64
	SetAlpha(value float64)
	// A Boolean value indicating whether the element is hidden.
	Hidden() bool
	SetHidden(value bool)
	// The element’s position on the z axis.
	ZIndex() int
	SetZIndex(value int)
}

// Init initializes the instance.
func (c NSCollectionViewLayoutAttributes) Init() NSCollectionViewLayoutAttributes {
	rv := objc.Send[NSCollectionViewLayoutAttributes](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewLayoutAttributes) Autorelease() NSCollectionViewLayoutAttributes {
	rv := objc.Send[NSCollectionViewLayoutAttributes](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewLayoutAttributes creates a new NSCollectionViewLayoutAttributes instance.
func NewNSCollectionViewLayoutAttributes() NSCollectionViewLayoutAttributes {
	class := getNSCollectionViewLayoutAttributesClass()
	rv := objc.Send[NSCollectionViewLayoutAttributes](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns a layout attributes object for a decoration view based
// on the specified information.
//
// decorationViewKind: A string that identifies the type of the decoration view. Use this string
// to differentiate from among the decoration views in a given section. This
// parameter must contain a valid value.
//
// indexPath: The index path of the item. You can use this information to identify the
// item in your app’s data structures.
//
// # Return Value
//
// A new layout attributes object configured with the initial attributes for
// the decoration view.
//
// # Discussion
//
// Call this method when you need to create a layout attributes object for a
// decoration view in a collection view. Decoration views are a tertiary type
// of content that display visual adornments in your collection view
// interface. For example, decoration views might display custom backgrounds.
// This method uses the parameters to set the initial values of the
// [IndexPath] and [RepresentedElementKind] properties the returned object.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forDecorationViewOfKind:with:)
func NewCollectionViewLayoutAttributesForDecorationViewOfKindWithIndexPath(decorationViewKind NSCollectionViewDecorationElementKind, indexPath foundation.INSIndexPath) NSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForDecorationViewOfKind:withIndexPath:"), objc.String(string(decorationViewKind)), indexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}

// Creates and returns a layout attributes object for an inter-item gap view
// at the specified index path.
//
// indexPath: The index path at which to insert the gap view. The gap is placed after the
// item specified by the index path. This parameter must contain a valid
// value.
//
// # Return Value
//
// A new layout attributes object configured with the initial attributes for
// the inter-item gap view.
//
// # Discussion
//
// Call this method when you need to create a layout attributes object for an
// inter-item gap view in a collection view. Gap views are used during drag
// and drop to indicate the area where content will drop. This method uses the
// parameters to set the initial values of the [IndexPath] property of the
// returned object. The [RepresentedElementKind] property is set to
// [ElementKindInterItemGapIndicator].
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forInterItemGapBefore:)
func NewCollectionViewLayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.INSIndexPath) NSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForInterItemGapBeforeIndexPath:"), indexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}

// Creates and returns a layout attributes object for the item at the
// specified index path.
//
// indexPath: The index path of the item. You can use this information to identify the
// item in your app’s data structures. This parameter must contain a valid
// value.
//
// # Return Value
//
// A new layout attributes object containing the initial attributes for the
// item.
//
// # Discussion
//
// Call this method when you need to create a layout attributes object for an
// item in a collection view. Items are the main type of content presented by
// a collection view. Items are grouped into sections, although a collection
// view may have only one section. This method assigns the provided index path
// to the [IndexPath] property of the returned object.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forItemWith:)
func NewCollectionViewLayoutAttributesForItemWithIndexPath(indexPath foundation.INSIndexPath) NSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForItemWithIndexPath:"), indexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}

// Creates and returns a layout attributes object for a supplementary view
// based on the specified information.
//
// elementKind: A string that identifies the type of the supplementary view. Use this
// string to differentiate from among the supplementary views in a given
// section. This parameter must contain a valid value.
//
// indexPath: The index path of the item. You can use this information to identify the
// item in your app’s data structures. This parameter must contain a valid
// value.
//
// # Return Value
//
// A new layout attributes object configured with the initial attributes for
// the supplementary view.
//
// # Discussion
//
// Call this method when you need to create a layout attributes object for a
// supplementary view in a collection view. Supplementary views are a
// secondary type of content that display data related to a specific section.
// For example, header and footer views in a grid layout implemented using
// supplementary views. This method uses the parameters to set the initial
// values of the [IndexPath] and [RepresentedElementKind] properties the
// returned object.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/init(forSupplementaryViewOfKind:with:)
func NewCollectionViewLayoutAttributesForSupplementaryViewOfKindWithIndexPath(elementKind NSCollectionViewSupplementaryElementKind, indexPath foundation.INSIndexPath) NSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionViewLayoutAttributesClass().class), objc.Sel("layoutAttributesForSupplementaryViewOfKind:withIndexPath:"), objc.String(string(elementKind)), indexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}

// The type of the element.
//
// # Discussion
//
// Use this property to distinguish whether the layout attributes apply to an
// item, a supplementary view, a decoration view, or another type of element
// presented by the collection view.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/representedElementCategory
func (c NSCollectionViewLayoutAttributes) RepresentedElementCategory() NSCollectionElementCategory {
	rv := objc.Send[NSCollectionElementCategory](c.ID, objc.Sel("representedElementCategory"))
	return NSCollectionElementCategory(rv)
}

// The index path of the element.
//
// # Discussion
//
// Use the index path to locate information about the item in your app’s
// data structures. For supplementary and decoration views, you must also use
// the [RepresentedElementKind] property to identify the element.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/indexPath
func (c NSCollectionViewLayoutAttributes) IndexPath() objc.ID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPath"))
	return rv
}
func (c NSCollectionViewLayoutAttributes) SetIndexPath(value objc.ID) {
	objc.Send[struct{}](c.ID, objc.Sel("setIndexPath:"), value)
}

// The identifier for specific elements of your collection view interface.
//
// # Discussion
//
// For supplementary and decoration views, you use this string to distinguish
// between views in a given section. You also use this string to identify the
// intended purpose of the view in your collection view interface.
//
// When the value of the [RepresentedElementCategory] property is
// [NSCollectionElementCategoryItem], this property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/representedElementKind
func (c NSCollectionViewLayoutAttributes) RepresentedElementKind() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("representedElementKind"))
	return foundation.NSStringFromID(rv).String()
}

// The frame rectangle of the element.
//
// # Discussion
//
// The frame rectangle is measured in points and specified in the collection
// view’s coordinate system. Setting the value of this property also updates
// the value in the [Size] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/frame
func (c NSCollectionViewLayoutAttributes) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}
func (c NSCollectionViewLayoutAttributes) SetFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](c.ID, objc.Sel("setFrame:"), value)
}

// The size of the element.
//
// # Discussion
//
// Setting the value of this property also updates the value in the [Frame]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/size
func (c NSCollectionViewLayoutAttributes) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
func (c NSCollectionViewLayoutAttributes) SetSize(value corefoundation.CGSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setSize:"), value)
}

// The transparency of the element.
//
// # Discussion
//
// Possible values are between `0.0` (fully transparent) and `1.0` (fully
// opaque). The default value is `1.0`.
//
// Transparent items continue to participate in hit testing for the collection
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/alpha
func (c NSCollectionViewLayoutAttributes) Alpha() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("alpha"))
	return rv
}
func (c NSCollectionViewLayoutAttributes) SetAlpha(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlpha:"), value)
}

// A Boolean value indicating whether the element is hidden.
//
// # Discussion
//
// The default value of this property is false. As an optimization, the
// collection view might not create the corresponding view when the value of
// this property is true. Because there might not be a view, hidden elements
// do not participate in hit testing for the collection view.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/isHidden
func (c NSCollectionViewLayoutAttributes) Hidden() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isHidden"))
	return rv
}
func (c NSCollectionViewLayoutAttributes) SetHidden(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setHidden:"), value)
}

// The element’s position on the z axis.
//
// # Discussion
//
// Use this property to specify the front-to-back ordering of items during
// layout. Items with higher index values appear on top of those with lower
// values. Items with the same value have an undetermined order.
//
// The default value of this property is `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayoutAttributes/zIndex
func (c NSCollectionViewLayoutAttributes) ZIndex() int {
	rv := objc.Send[int](c.ID, objc.Sel("zIndex"))
	return rv
}
func (c NSCollectionViewLayoutAttributes) SetZIndex(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setZIndex:"), value)
}

// The element kind string assigned to the attributes object when it
// represents an inter-item gap.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionview/elementkindinteritemgapindicator
func (_NSCollectionViewLayoutAttributesClass NSCollectionViewLayoutAttributesClass) ElementKindInterItemGapIndicator() string {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionViewLayoutAttributesClass.class), objc.Sel("NSCollectionElementKindInterItemGapIndicator"))
	return foundation.NSStringFromID(rv).String()
}

// A supplementary view that acts as a footer for a given section.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionview/elementkindsectionfooter
func (_NSCollectionViewLayoutAttributesClass NSCollectionViewLayoutAttributesClass) ElementKindSectionFooter() string {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionViewLayoutAttributesClass.class), objc.Sel("NSCollectionElementKindSectionFooter"))
	return foundation.NSStringFromID(rv).String()
}

// A supplementary view that acts as a header for a given section.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionview/elementkindsectionheader
func (_NSCollectionViewLayoutAttributesClass NSCollectionViewLayoutAttributesClass) ElementKindSectionHeader() string {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionViewLayoutAttributesClass.class), objc.Sel("NSCollectionElementKindSectionHeader"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the class to use for layout attribute objects
//
// See: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/layoutattributesclass
func (_NSCollectionViewLayoutAttributesClass NSCollectionViewLayoutAttributesClass) LayoutAttributesClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSCollectionViewLayoutAttributesClass.class), objc.Sel("layoutAttributesClass"))
	return rv
}
func (_NSCollectionViewLayoutAttributesClass NSCollectionViewLayoutAttributesClass) SetLayoutAttributesClass(value objc.Class) {
	objc.Send[struct{}](objc.ID(_NSCollectionViewLayoutAttributesClass.class), objc.Sel("setLayoutAttributesClass:"), value)
}
