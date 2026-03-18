// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionViewCompositionalLayoutConfiguration] class.
var (
	_NSCollectionViewCompositionalLayoutConfigurationClass     NSCollectionViewCompositionalLayoutConfigurationClass
	_NSCollectionViewCompositionalLayoutConfigurationClassOnce sync.Once
)

func getNSCollectionViewCompositionalLayoutConfigurationClass() NSCollectionViewCompositionalLayoutConfigurationClass {
	_NSCollectionViewCompositionalLayoutConfigurationClassOnce.Do(func() {
		_NSCollectionViewCompositionalLayoutConfigurationClass = NSCollectionViewCompositionalLayoutConfigurationClass{class: objc.GetClass("NSCollectionViewCompositionalLayoutConfiguration")}
	})
	return _NSCollectionViewCompositionalLayoutConfigurationClass
}

// GetNSCollectionViewCompositionalLayoutConfigurationClass returns the class object for NSCollectionViewCompositionalLayoutConfiguration.
func GetNSCollectionViewCompositionalLayoutConfigurationClass() NSCollectionViewCompositionalLayoutConfigurationClass {
	return getNSCollectionViewCompositionalLayoutConfigurationClass()
}

type NSCollectionViewCompositionalLayoutConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewCompositionalLayoutConfigurationClass) Alloc() NSCollectionViewCompositionalLayoutConfiguration {
	rv := objc.Send[NSCollectionViewCompositionalLayoutConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that defines scroll direction, section spacing, and headers or
// footers for the layout.
//
// # Overview
// 
// You use a layout configuration to modify a collection view layout’s
// default scroll direction, add extra spacing between each section of the
// layout, and add headers or footers to the entire layout.
// 
// You can pass in this configuration when creating an
// [NSCollectionViewCompositionalLayout], or you can set the [NSCollectionViewCompositionalLayoutConfiguration.Configuration]
// property on an existing layout. If you modify the configuration on an
// existing layout, the system invalidates the layout so that it will be
// updated with the new configuration.
//
// # Specifying Scroll Direction
//
//   - [NSCollectionViewCompositionalLayoutConfiguration.ScrollDirection]: The axis that the content in the collection view layout scrolls along.
//   - [NSCollectionViewCompositionalLayoutConfiguration.SetScrollDirection]
//
// # Configuring Spacing
//
//   - [NSCollectionViewCompositionalLayoutConfiguration.InterSectionSpacing]: The amount of space between the sections in the layout.
//   - [NSCollectionViewCompositionalLayoutConfiguration.SetInterSectionSpacing]
//
// # Configuring Additional Views
//
//   - [NSCollectionViewCompositionalLayoutConfiguration.BoundarySupplementaryItems]: An array of the supplementary items that are associated with the boundary edges of the entire layout, such as global headers and footers.
//   - [NSCollectionViewCompositionalLayoutConfiguration.SetBoundarySupplementaryItems]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration
type NSCollectionViewCompositionalLayoutConfiguration struct {
	objectivec.Object
}

// NSCollectionViewCompositionalLayoutConfigurationFromID constructs a [NSCollectionViewCompositionalLayoutConfiguration] from an objc.ID.
//
// An object that defines scroll direction, section spacing, and headers or
// footers for the layout.
func NSCollectionViewCompositionalLayoutConfigurationFromID(id objc.ID) NSCollectionViewCompositionalLayoutConfiguration {
	return NSCollectionViewCompositionalLayoutConfiguration{objectivec.Object{ID: id}}
}
// NOTE: NSCollectionViewCompositionalLayoutConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionViewCompositionalLayoutConfiguration] class.
//
// # Specifying Scroll Direction
//
//   - [INSCollectionViewCompositionalLayoutConfiguration.ScrollDirection]: The axis that the content in the collection view layout scrolls along.
//   - [INSCollectionViewCompositionalLayoutConfiguration.SetScrollDirection]
//
// # Configuring Spacing
//
//   - [INSCollectionViewCompositionalLayoutConfiguration.InterSectionSpacing]: The amount of space between the sections in the layout.
//   - [INSCollectionViewCompositionalLayoutConfiguration.SetInterSectionSpacing]
//
// # Configuring Additional Views
//
//   - [INSCollectionViewCompositionalLayoutConfiguration.BoundarySupplementaryItems]: An array of the supplementary items that are associated with the boundary edges of the entire layout, such as global headers and footers.
//   - [INSCollectionViewCompositionalLayoutConfiguration.SetBoundarySupplementaryItems]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration
type INSCollectionViewCompositionalLayoutConfiguration interface {
	objectivec.IObject

	// Topic: Specifying Scroll Direction

	// The axis that the content in the collection view layout scrolls along.
	ScrollDirection() NSCollectionViewScrollDirection
	SetScrollDirection(value NSCollectionViewScrollDirection)

	// Topic: Configuring Spacing

	// The amount of space between the sections in the layout.
	InterSectionSpacing() float64
	SetInterSectionSpacing(value float64)

	// Topic: Configuring Additional Views

	// An array of the supplementary items that are associated with the boundary edges of the entire layout, such as global headers and footers.
	BoundarySupplementaryItems() []NSCollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []NSCollectionLayoutBoundarySupplementaryItem)

	// The layout’s configuration, such as its scroll direction and section spacing.
	Configuration() INSCollectionViewCompositionalLayoutConfiguration
	SetConfiguration(value INSCollectionViewCompositionalLayoutConfiguration)
}

// Init initializes the instance.
func (c NSCollectionViewCompositionalLayoutConfiguration) Init() NSCollectionViewCompositionalLayoutConfiguration {
	rv := objc.Send[NSCollectionViewCompositionalLayoutConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewCompositionalLayoutConfiguration) Autorelease() NSCollectionViewCompositionalLayoutConfiguration {
	rv := objc.Send[NSCollectionViewCompositionalLayoutConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewCompositionalLayoutConfiguration creates a new NSCollectionViewCompositionalLayoutConfiguration instance.
func NewNSCollectionViewCompositionalLayoutConfiguration() NSCollectionViewCompositionalLayoutConfiguration {
	class := getNSCollectionViewCompositionalLayoutConfigurationClass()
	rv := objc.Send[NSCollectionViewCompositionalLayoutConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The axis that the content in the collection view layout scrolls along.
//
// # Discussion
// 
// The default value of this property is
// [NSCollectionView.ScrollDirection.vertical].
//
// [NSCollectionView.ScrollDirection.vertical]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollDirection/vertical
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration/scrollDirection
func (c NSCollectionViewCompositionalLayoutConfiguration) ScrollDirection() NSCollectionViewScrollDirection {
	rv := objc.Send[NSCollectionViewScrollDirection](c.ID, objc.Sel("scrollDirection"))
	return NSCollectionViewScrollDirection(rv)
}
func (c NSCollectionViewCompositionalLayoutConfiguration) SetScrollDirection(value NSCollectionViewScrollDirection) {
	objc.Send[struct{}](c.ID, objc.Sel("setScrollDirection:"), value)
}

// The amount of space between the sections in the layout.
//
// # Discussion
// 
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration/interSectionSpacing
func (c NSCollectionViewCompositionalLayoutConfiguration) InterSectionSpacing() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("interSectionSpacing"))
	return rv
}
func (c NSCollectionViewCompositionalLayoutConfiguration) SetInterSectionSpacing(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setInterSectionSpacing:"), value)
}

// An array of the supplementary items that are associated with the boundary
// edges of the entire layout, such as global headers and footers.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutConfiguration/boundarySupplementaryItems
func (c NSCollectionViewCompositionalLayoutConfiguration) BoundarySupplementaryItems() []NSCollectionLayoutBoundarySupplementaryItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("boundarySupplementaryItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSCollectionLayoutBoundarySupplementaryItem {
		return NSCollectionLayoutBoundarySupplementaryItemFromID(id)
	})
}
func (c NSCollectionViewCompositionalLayoutConfiguration) SetBoundarySupplementaryItems(value []NSCollectionLayoutBoundarySupplementaryItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setBoundarySupplementaryItems:"), objectivec.IObjectSliceToNSArray(value))
}

// The layout’s configuration, such as its scroll direction and section
// spacing.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionviewcompositionallayout/configuration
func (c NSCollectionViewCompositionalLayoutConfiguration) Configuration() INSCollectionViewCompositionalLayoutConfiguration {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("configuration"))
	return NSCollectionViewCompositionalLayoutConfigurationFromID(objc.ID(rv))
}
func (c NSCollectionViewCompositionalLayoutConfiguration) SetConfiguration(value INSCollectionViewCompositionalLayoutConfiguration) {
	objc.Send[struct{}](c.ID, objc.Sel("setConfiguration:"), value)
}

