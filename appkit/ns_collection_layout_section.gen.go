// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutSection] class.
var (
	_NSCollectionLayoutSectionClass     NSCollectionLayoutSectionClass
	_NSCollectionLayoutSectionClassOnce sync.Once
)

func getNSCollectionLayoutSectionClass() NSCollectionLayoutSectionClass {
	_NSCollectionLayoutSectionClassOnce.Do(func() {
		_NSCollectionLayoutSectionClass = NSCollectionLayoutSectionClass{class: objc.GetClass("NSCollectionLayoutSection")}
	})
	return _NSCollectionLayoutSectionClass
}

// GetNSCollectionLayoutSectionClass returns the class object for NSCollectionLayoutSection.
func GetNSCollectionLayoutSectionClass() NSCollectionLayoutSectionClass {
	return getNSCollectionLayoutSectionClass()
}

type NSCollectionLayoutSectionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutSectionClass) Alloc() NSCollectionLayoutSection {
	rv := objc.Send[NSCollectionLayoutSection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A container that combines a set of groups into distinct visual groupings.
//
// # Overview
// 
// A collection view layout has one or more sections. Sections provide a way
// to separate the layout into distinct pieces.
// 
// Each section can have the same layout or a different layout than the other
// sections in the collection view. A section’s layout is determined by the
// properties of the group ([NSCollectionLayoutGroup]) that’s used to create
// the section.
// 
// In the Photos app, each section in the Years page uses the same layout. In
// the App Store, the Apps page displays several sections with different
// content arrangements.
// 
// [media-3568661]
// 
// Each section can have its own background, header, and footer to distinguish
// it from other sections.
//
// # Specifying scrolling behavior
//
//   - [NSCollectionLayoutSection.OrthogonalScrollingBehavior]: The section’s scrolling behavior in relation to the main layout axis.
//   - [NSCollectionLayoutSection.SetOrthogonalScrollingBehavior]
//
// # Configuring section spacing
//
//   - [NSCollectionLayoutSection.InterGroupSpacing]: The amount of space between the groups in the section.
//   - [NSCollectionLayoutSection.SetInterGroupSpacing]
//   - [NSCollectionLayoutSection.ContentInsets]: The amount of space between the content of the section and its boundaries.
//   - [NSCollectionLayoutSection.SetContentInsets]
//
// # Configuring additional views
//
//   - [NSCollectionLayoutSection.BoundarySupplementaryItems]: An array of the supplementary items that are associated with the boundary edges of the section, such as headers and footers.
//   - [NSCollectionLayoutSection.SetBoundarySupplementaryItems]
//   - [NSCollectionLayoutSection.DecorationItems]: An array of the decoration items that are anchored to the section, such as background decoration views.
//   - [NSCollectionLayoutSection.SetDecorationItems]
//
// # Rendering items
//
//   - [NSCollectionLayoutSection.VisibleItemsInvalidationHandler]: A closure called before each layout cycle to allow modification of the items in the section immediately before they’re displayed.
//   - [NSCollectionLayoutSection.SetVisibleItemsInvalidationHandler]
//
// # Deprecated
//
//   - [NSCollectionLayoutSection.SupplementariesFollowContentInsets]: A Boolean value that indicates whether the section’s supplementary items follow the specified content insets for the section.
//   - [NSCollectionLayoutSection.SetSupplementariesFollowContentInsets]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection
type NSCollectionLayoutSection struct {
	objectivec.Object
}

// NSCollectionLayoutSectionFromID constructs a [NSCollectionLayoutSection] from an objc.ID.
//
// A container that combines a set of groups into distinct visual groupings.
func NSCollectionLayoutSectionFromID(id objc.ID) NSCollectionLayoutSection {
	return NSCollectionLayoutSection{objectivec.Object{ID: id}}
}
// NOTE: NSCollectionLayoutSection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionLayoutSection] class.
//
// # Specifying scrolling behavior
//
//   - [INSCollectionLayoutSection.OrthogonalScrollingBehavior]: The section’s scrolling behavior in relation to the main layout axis.
//   - [INSCollectionLayoutSection.SetOrthogonalScrollingBehavior]
//
// # Configuring section spacing
//
//   - [INSCollectionLayoutSection.InterGroupSpacing]: The amount of space between the groups in the section.
//   - [INSCollectionLayoutSection.SetInterGroupSpacing]
//   - [INSCollectionLayoutSection.ContentInsets]: The amount of space between the content of the section and its boundaries.
//   - [INSCollectionLayoutSection.SetContentInsets]
//
// # Configuring additional views
//
//   - [INSCollectionLayoutSection.BoundarySupplementaryItems]: An array of the supplementary items that are associated with the boundary edges of the section, such as headers and footers.
//   - [INSCollectionLayoutSection.SetBoundarySupplementaryItems]
//   - [INSCollectionLayoutSection.DecorationItems]: An array of the decoration items that are anchored to the section, such as background decoration views.
//   - [INSCollectionLayoutSection.SetDecorationItems]
//
// # Rendering items
//
//   - [INSCollectionLayoutSection.VisibleItemsInvalidationHandler]: A closure called before each layout cycle to allow modification of the items in the section immediately before they’re displayed.
//   - [INSCollectionLayoutSection.SetVisibleItemsInvalidationHandler]
//
// # Deprecated
//
//   - [INSCollectionLayoutSection.SupplementariesFollowContentInsets]: A Boolean value that indicates whether the section’s supplementary items follow the specified content insets for the section.
//   - [INSCollectionLayoutSection.SetSupplementariesFollowContentInsets]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection
type INSCollectionLayoutSection interface {
	objectivec.IObject

	// Topic: Specifying scrolling behavior

	// The section’s scrolling behavior in relation to the main layout axis.
	OrthogonalScrollingBehavior() NSCollectionLayoutSectionOrthogonalScrollingBehavior
	SetOrthogonalScrollingBehavior(value NSCollectionLayoutSectionOrthogonalScrollingBehavior)

	// Topic: Configuring section spacing

	// The amount of space between the groups in the section.
	InterGroupSpacing() float64
	SetInterGroupSpacing(value float64)
	// The amount of space between the content of the section and its boundaries.
	ContentInsets() NSDirectionalEdgeInsets
	SetContentInsets(value NSDirectionalEdgeInsets)

	// Topic: Configuring additional views

	// An array of the supplementary items that are associated with the boundary edges of the section, such as headers and footers.
	BoundarySupplementaryItems() []NSCollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []NSCollectionLayoutBoundarySupplementaryItem)
	// An array of the decoration items that are anchored to the section, such as background decoration views.
	DecorationItems() []NSCollectionLayoutDecorationItem
	SetDecorationItems(value []NSCollectionLayoutDecorationItem)

	// Topic: Rendering items

	// A closure called before each layout cycle to allow modification of the items in the section immediately before they’re displayed.
	VisibleItemsInvalidationHandler() NSCollectionLayoutSectionVisibleItemsInvalidationHandler
	SetVisibleItemsInvalidationHandler(value NSCollectionLayoutSectionVisibleItemsInvalidationHandler)

	// Topic: Deprecated

	// A Boolean value that indicates whether the section’s supplementary items follow the specified content insets for the section.
	SupplementariesFollowContentInsets() bool
	SetSupplementariesFollowContentInsets(value bool)
}

// Init initializes the instance.
func (c NSCollectionLayoutSection) Init() NSCollectionLayoutSection {
	rv := objc.Send[NSCollectionLayoutSection](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutSection) Autorelease() NSCollectionLayoutSection {
	rv := objc.Send[NSCollectionLayoutSection](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutSection creates a new NSCollectionLayoutSection instance.
func NewNSCollectionLayoutSection() NSCollectionLayoutSection {
	class := getNSCollectionLayoutSectionClass()
	rv := objc.Send[NSCollectionLayoutSection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a section containing the specified group.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/init(group:)
func NewCollectionLayoutSectionWithGroup(group INSCollectionLayoutGroup) NSCollectionLayoutSection {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutSectionClass().class), objc.Sel("sectionWithGroup:"), group)
	return NSCollectionLayoutSectionFromID(rv)
}

// The section’s scrolling behavior in relation to the main layout axis.
//
// # Discussion
// 
// The default value of this property is
// [UICollectionLayoutSectionOrthogonalScrollingBehavior.none], which means
// the section lays out its content along the main axis of its layout, defined
// by the layout configuration’s [scrollDirection] property. Set a different
// value for this property to get the section to lay out its content
// orthogonally to the main layout axis.
//
// [UICollectionLayoutSectionOrthogonalScrollingBehavior.none]: https://developer.apple.com/documentation/UIKit/UICollectionLayoutSectionOrthogonalScrollingBehavior/none
// [scrollDirection]: https://developer.apple.com/documentation/UIKit/UICollectionViewCompositionalLayoutConfiguration/scrollDirection
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/orthogonalScrollingBehavior
func (c NSCollectionLayoutSection) OrthogonalScrollingBehavior() NSCollectionLayoutSectionOrthogonalScrollingBehavior {
	rv := objc.Send[NSCollectionLayoutSectionOrthogonalScrollingBehavior](c.ID, objc.Sel("orthogonalScrollingBehavior"))
	return NSCollectionLayoutSectionOrthogonalScrollingBehavior(rv)
}
func (c NSCollectionLayoutSection) SetOrthogonalScrollingBehavior(value NSCollectionLayoutSectionOrthogonalScrollingBehavior) {
	objc.Send[struct{}](c.ID, objc.Sel("setOrthogonalScrollingBehavior:"), value)
}

// The amount of space between the groups in the section.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/interGroupSpacing
func (c NSCollectionLayoutSection) InterGroupSpacing() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("interGroupSpacing"))
	return rv
}
func (c NSCollectionLayoutSection) SetInterGroupSpacing(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setInterGroupSpacing:"), value)
}

// The amount of space between the content of the section and its boundaries.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/contentInsets
func (c NSCollectionLayoutSection) ContentInsets() NSDirectionalEdgeInsets {
	rv := objc.Send[NSDirectionalEdgeInsets](c.ID, objc.Sel("contentInsets"))
	return NSDirectionalEdgeInsets(rv)
}
func (c NSCollectionLayoutSection) SetContentInsets(value NSDirectionalEdgeInsets) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentInsets:"), value)
}

// An array of the supplementary items that are associated with the boundary
// edges of the section, such as headers and footers.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/boundarySupplementaryItems
func (c NSCollectionLayoutSection) BoundarySupplementaryItems() []NSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("boundarySupplementaryItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSCollectionLayoutBoundarySupplementaryItem {
		return NSCollectionLayoutBoundarySupplementaryItemFromID(id)
	})
}
func (c NSCollectionLayoutSection) SetBoundarySupplementaryItems(value []NSCollectionLayoutBoundarySupplementaryItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setBoundarySupplementaryItems:"), objectivec.IObjectSliceToNSArray(value))
}

// An array of the decoration items that are anchored to the section, such as
// background decoration views.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/decorationItems
func (c NSCollectionLayoutSection) DecorationItems() []NSCollectionLayoutDecorationItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("decorationItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSCollectionLayoutDecorationItem {
		return NSCollectionLayoutDecorationItemFromID(id)
	})
}
func (c NSCollectionLayoutSection) SetDecorationItems(value []NSCollectionLayoutDecorationItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setDecorationItems:"), objectivec.IObjectSliceToNSArray(value))
}

// A closure called before each layout cycle to allow modification of the
// items in the section immediately before they’re displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/visibleItemsInvalidationHandler
func (c NSCollectionLayoutSection) VisibleItemsInvalidationHandler() NSCollectionLayoutSectionVisibleItemsInvalidationHandler {
	rv := objc.Send[NSCollectionLayoutSectionVisibleItemsInvalidationHandler](c.ID, objc.Sel("visibleItemsInvalidationHandler"))
	return NSCollectionLayoutSectionVisibleItemsInvalidationHandler(rv)
}
func (c NSCollectionLayoutSection) SetVisibleItemsInvalidationHandler(value NSCollectionLayoutSectionVisibleItemsInvalidationHandler) {
	objc.Send[struct{}](c.ID, objc.Sel("setVisibleItemsInvalidationHandler:"), value)
}

// A Boolean value that indicates whether the section’s supplementary items
// follow the specified content insets for the section.
//
// # Discussion
// 
// The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSection/supplementariesFollowContentInsets
func (c NSCollectionLayoutSection) SupplementariesFollowContentInsets() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("supplementariesFollowContentInsets"))
	return rv
}
func (c NSCollectionLayoutSection) SetSupplementariesFollowContentInsets(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSupplementariesFollowContentInsets:"), value)
}

