// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSCollectionViewCompositionalLayout] class.
var (
	_NSCollectionViewCompositionalLayoutClass     NSCollectionViewCompositionalLayoutClass
	_NSCollectionViewCompositionalLayoutClassOnce sync.Once
)

func getNSCollectionViewCompositionalLayoutClass() NSCollectionViewCompositionalLayoutClass {
	_NSCollectionViewCompositionalLayoutClassOnce.Do(func() {
		_NSCollectionViewCompositionalLayoutClass = NSCollectionViewCompositionalLayoutClass{class: objc.GetClass("NSCollectionViewCompositionalLayout")}
	})
	return _NSCollectionViewCompositionalLayoutClass
}

// GetNSCollectionViewCompositionalLayoutClass returns the class object for NSCollectionViewCompositionalLayout.
func GetNSCollectionViewCompositionalLayoutClass() NSCollectionViewCompositionalLayoutClass {
	return getNSCollectionViewCompositionalLayoutClass()
}

type NSCollectionViewCompositionalLayoutClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCollectionViewCompositionalLayoutClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewCompositionalLayoutClass) Alloc() NSCollectionViewCompositionalLayout {
	rv := objc.Send[NSCollectionViewCompositionalLayout](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A layout object that lets you combine items in highly adaptive and flexible
// visual arrangements.
//
// # Overview
//
// A compositional layout is a type of collection view layout. It’s designed
// to be composable, flexible, and fast, letting you build any kind of visual
// arrangement for your content by combining—or compositing—each smaller
// component into a full layout.
//
// A compositional layout is composed of one or more sections that break up
// the layout into distinct visual groupings. Each section is composed of
// groups of individual items, the smallest unit of data you want to present.
// A group might lay out its items in a horizontal row, a vertical column, or
// a custom arrangement.
//
// You combine the components by building up from items into a group, from
// groups into a section, and finally into a full layout, like in this example
// of a basic list layout:
//
// # Creating a Layout
//
//   - [NSCollectionViewCompositionalLayout.InitWithSection]: Creates a compositional layout object with a single section.
//   - [NSCollectionViewCompositionalLayout.InitWithSectionConfiguration]: Creates a compositional layout object with a single section and an additional configuration.
//   - [NSCollectionViewCompositionalLayout.InitWithSectionProvider]: Creates a compositional layout object with a section provider to supply the layout’s sections.
//   - [NSCollectionViewCompositionalLayout.InitWithSectionProviderConfiguration]: Creates a compositional layout object with a section provider and an additional configuration.
//
// # Configuring the Layout
//
//   - [NSCollectionViewCompositionalLayout.Configuration]: The layout’s configuration, such as its scroll direction and section spacing.
//   - [NSCollectionViewCompositionalLayout.SetConfiguration]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout
type NSCollectionViewCompositionalLayout struct {
	NSCollectionViewLayout
}

// NSCollectionViewCompositionalLayoutFromID constructs a [NSCollectionViewCompositionalLayout] from an objc.ID.
//
// A layout object that lets you combine items in highly adaptive and flexible
// visual arrangements.
func NSCollectionViewCompositionalLayoutFromID(id objc.ID) NSCollectionViewCompositionalLayout {
	return NSCollectionViewCompositionalLayout{NSCollectionViewLayout: NSCollectionViewLayoutFromID(id)}
}

// NOTE: NSCollectionViewCompositionalLayout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionViewCompositionalLayout] class.
//
// # Creating a Layout
//
//   - [INSCollectionViewCompositionalLayout.InitWithSection]: Creates a compositional layout object with a single section.
//   - [INSCollectionViewCompositionalLayout.InitWithSectionConfiguration]: Creates a compositional layout object with a single section and an additional configuration.
//   - [INSCollectionViewCompositionalLayout.InitWithSectionProvider]: Creates a compositional layout object with a section provider to supply the layout’s sections.
//   - [INSCollectionViewCompositionalLayout.InitWithSectionProviderConfiguration]: Creates a compositional layout object with a section provider and an additional configuration.
//
// # Configuring the Layout
//
//   - [INSCollectionViewCompositionalLayout.Configuration]: The layout’s configuration, such as its scroll direction and section spacing.
//   - [INSCollectionViewCompositionalLayout.SetConfiguration]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout
type INSCollectionViewCompositionalLayout interface {
	INSCollectionViewLayout

	// Topic: Creating a Layout

	// Creates a compositional layout object with a single section.
	InitWithSection(section INSCollectionLayoutSection) NSCollectionViewCompositionalLayout
	// Creates a compositional layout object with a single section and an additional configuration.
	InitWithSectionConfiguration(section INSCollectionLayoutSection, configuration INSCollectionViewCompositionalLayoutConfiguration) NSCollectionViewCompositionalLayout
	// Creates a compositional layout object with a section provider to supply the layout’s sections.
	InitWithSectionProvider(sectionProvider NSCollectionLayoutSectionInt64Handler) NSCollectionViewCompositionalLayout
	// Creates a compositional layout object with a section provider and an additional configuration.
	InitWithSectionProviderConfiguration(sectionProvider NSCollectionLayoutSectionInt64Handler, configuration INSCollectionViewCompositionalLayoutConfiguration) NSCollectionViewCompositionalLayout

	// Topic: Configuring the Layout

	// The layout’s configuration, such as its scroll direction and section spacing.
	Configuration() INSCollectionViewCompositionalLayoutConfiguration
	SetConfiguration(value INSCollectionViewCompositionalLayoutConfiguration)
}

// Init initializes the instance.
func (c NSCollectionViewCompositionalLayout) Init() NSCollectionViewCompositionalLayout {
	rv := objc.Send[NSCollectionViewCompositionalLayout](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewCompositionalLayout) Autorelease() NSCollectionViewCompositionalLayout {
	rv := objc.Send[NSCollectionViewCompositionalLayout](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewCompositionalLayout creates a new NSCollectionViewCompositionalLayout instance.
func NewNSCollectionViewCompositionalLayout() NSCollectionViewCompositionalLayout {
	class := getNSCollectionViewCompositionalLayoutClass()
	rv := objc.Send[NSCollectionViewCompositionalLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout/init(coder:)
func NewCollectionViewCompositionalLayoutWithCoder(coder foundation.INSCoder) NSCollectionViewCompositionalLayout {
	instance := getNSCollectionViewCompositionalLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCollectionViewCompositionalLayoutFromID(rv)
}

// Creates a compositional layout object with a single section.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(section:)
func NewCollectionViewCompositionalLayoutWithSection(section INSCollectionLayoutSection) NSCollectionViewCompositionalLayout {
	instance := getNSCollectionViewCompositionalLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSection:"), section)
	return NSCollectionViewCompositionalLayoutFromID(rv)
}

// Creates a compositional layout object with a single section and an
// additional configuration.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(section:configuration:)
func NewCollectionViewCompositionalLayoutWithSectionConfiguration(section INSCollectionLayoutSection, configuration INSCollectionViewCompositionalLayoutConfiguration) NSCollectionViewCompositionalLayout {
	instance := getNSCollectionViewCompositionalLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSection:configuration:"), section, configuration)
	return NSCollectionViewCompositionalLayoutFromID(rv)
}

// Creates a compositional layout object with a section provider to supply the
// layout’s sections.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(sectionProvider:)
func NewCollectionViewCompositionalLayoutWithSectionProvider(sectionProvider NSCollectionLayoutSectionInt64Handler) NSCollectionViewCompositionalLayout {
	_block0, _ := NewNSCollectionLayoutSectionInt64Block(sectionProvider)
	instance := getNSCollectionViewCompositionalLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSectionProvider:"), _block0)
	return NSCollectionViewCompositionalLayoutFromID(rv)
}

// Creates a compositional layout object with a section provider and an
// additional configuration.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(sectionProvider:configuration:)
func NewCollectionViewCompositionalLayoutWithSectionProviderConfiguration(sectionProvider NSCollectionLayoutSectionInt64Handler, configuration INSCollectionViewCompositionalLayoutConfiguration) NSCollectionViewCompositionalLayout {
	_block0, _ := NewNSCollectionLayoutSectionInt64Block(sectionProvider)
	instance := getNSCollectionViewCompositionalLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSectionProvider:configuration:"), _block0, configuration)
	return NSCollectionViewCompositionalLayoutFromID(rv)
}

// Creates a compositional layout object with a single section.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(section:)
func (c NSCollectionViewCompositionalLayout) InitWithSection(section INSCollectionLayoutSection) NSCollectionViewCompositionalLayout {
	rv := objc.Send[NSCollectionViewCompositionalLayout](c.ID, objc.Sel("initWithSection:"), section)
	return rv
}

// Creates a compositional layout object with a single section and an
// additional configuration.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(section:configuration:)
func (c NSCollectionViewCompositionalLayout) InitWithSectionConfiguration(section INSCollectionLayoutSection, configuration INSCollectionViewCompositionalLayoutConfiguration) NSCollectionViewCompositionalLayout {
	rv := objc.Send[NSCollectionViewCompositionalLayout](c.ID, objc.Sel("initWithSection:configuration:"), section, configuration)
	return rv
}

var _nscollectionviewcompositionallayout_initwithsectionprovider_p0_key byte

// Creates a compositional layout object with a section provider to supply the
// layout’s sections.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(sectionProvider:)
func (c NSCollectionViewCompositionalLayout) InitWithSectionProvider(sectionProvider NSCollectionLayoutSectionInt64Handler) NSCollectionViewCompositionalLayout {
	_block0, _ := NewNSCollectionLayoutSectionInt64Block(sectionProvider)
	rv := objc.Send[NSCollectionViewCompositionalLayout](c.ID, objc.Sel("initWithSectionProvider:"), _block0)
	return rv
}

var _nscollectionviewcompositionallayout_initwithsectionprovider_configuration_p0_key byte

// Creates a compositional layout object with a section provider and an
// additional configuration.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/init(sectionProvider:configuration:)
func (c NSCollectionViewCompositionalLayout) InitWithSectionProviderConfiguration(sectionProvider NSCollectionLayoutSectionInt64Handler, configuration INSCollectionViewCompositionalLayoutConfiguration) NSCollectionViewCompositionalLayout {
	_block0, _ := NewNSCollectionLayoutSectionInt64Block(sectionProvider)
	rv := objc.Send[NSCollectionViewCompositionalLayout](c.ID, objc.Sel("initWithSectionProvider:configuration:"), _block0, configuration)
	return rv
}

// The layout’s configuration, such as its scroll direction and section
// spacing.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayout/configuration
func (c NSCollectionViewCompositionalLayout) Configuration() INSCollectionViewCompositionalLayoutConfiguration {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("configuration"))
	return NSCollectionViewCompositionalLayoutConfigurationFromID(objc.ID(rv))
}
func (c NSCollectionViewCompositionalLayout) SetConfiguration(value INSCollectionViewCompositionalLayoutConfiguration) {
	objc.Send[struct{}](c.ID, objc.Sel("setConfiguration:"), value)
}
