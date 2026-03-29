// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSScrubberProportionalLayout] class.
var (
	_NSScrubberProportionalLayoutClass     NSScrubberProportionalLayoutClass
	_NSScrubberProportionalLayoutClassOnce sync.Once
)

func getNSScrubberProportionalLayoutClass() NSScrubberProportionalLayoutClass {
	_NSScrubberProportionalLayoutClassOnce.Do(func() {
		_NSScrubberProportionalLayoutClass = NSScrubberProportionalLayoutClass{class: objc.GetClass("NSScrubberProportionalLayout")}
	})
	return _NSScrubberProportionalLayoutClass
}

// GetNSScrubberProportionalLayoutClass returns the class object for NSScrubberProportionalLayout.
func GetNSScrubberProportionalLayoutClass() NSScrubberProportionalLayoutClass {
	return getNSScrubberProportionalLayoutClass()
}

type NSScrubberProportionalLayoutClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSScrubberProportionalLayoutClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrubberProportionalLayoutClass) Alloc() NSScrubberProportionalLayout {
	rv := objc.Send[NSScrubberProportionalLayout](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A concrete layout object that sizes each item to some fraction of the
// scrubber’s visible size.
//
// # Initializing a proprotional layout
//
//   - [NSScrubberProportionalLayout.InitWithNumberOfVisibleItems]: Initializes and returns a newly allocated proportional layout, configured to display the given number of items.
//
// # Configuring the layout
//
//   - [NSScrubberProportionalLayout.NumberOfVisibleItems]: The number of items visible in the scrubber at once.
//   - [NSScrubberProportionalLayout.SetNumberOfVisibleItems]
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout
type NSScrubberProportionalLayout struct {
	NSScrubberLayout
}

// NSScrubberProportionalLayoutFromID constructs a [NSScrubberProportionalLayout] from an objc.ID.
//
// A concrete layout object that sizes each item to some fraction of the
// scrubber’s visible size.
func NSScrubberProportionalLayoutFromID(id objc.ID) NSScrubberProportionalLayout {
	return NSScrubberProportionalLayout{NSScrubberLayout: NSScrubberLayoutFromID(id)}
}
// NOTE: NSScrubberProportionalLayout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScrubberProportionalLayout] class.
//
// # Initializing a proprotional layout
//
//   - [INSScrubberProportionalLayout.InitWithNumberOfVisibleItems]: Initializes and returns a newly allocated proportional layout, configured to display the given number of items.
//
// # Configuring the layout
//
//   - [INSScrubberProportionalLayout.NumberOfVisibleItems]: The number of items visible in the scrubber at once.
//   - [INSScrubberProportionalLayout.SetNumberOfVisibleItems]
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout
type INSScrubberProportionalLayout interface {
	INSScrubberLayout

	// Topic: Initializing a proprotional layout

	// Initializes and returns a newly allocated proportional layout, configured to display the given number of items.
	InitWithNumberOfVisibleItems(numberOfVisibleItems int) NSScrubberProportionalLayout

	// Topic: Configuring the layout

	// The number of items visible in the scrubber at once.
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
}

// Init initializes the instance.
func (s NSScrubberProportionalLayout) Init() NSScrubberProportionalLayout {
	rv := objc.Send[NSScrubberProportionalLayout](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrubberProportionalLayout) Autorelease() NSScrubberProportionalLayout {
	rv := objc.Send[NSScrubberProportionalLayout](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrubberProportionalLayout creates a new NSScrubberProportionalLayout instance.
func NewNSScrubberProportionalLayout() NSScrubberProportionalLayout {
	class := getNSScrubberProportionalLayoutClass()
	rv := objc.Send[NSScrubberProportionalLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a newly allocated proprotional layout object from a
// storyboard or nib file.
//
// # Return Value
// 
// A properly initialized proportional layout object.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout/init(coder:)
func NewScrubberProportionalLayoutWithCoder(coder foundation.INSCoder) NSScrubberProportionalLayout {
	instance := getNSScrubberProportionalLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrubberProportionalLayoutFromID(rv)
}

// Initializes and returns a newly allocated proportional layout, configured
// to display the given number of items.
//
// numberOfVisibleItems: The number of items that should be visible in the scrubber at once.
//
// # Return Value
// 
// A properly initialized proportional layout object.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout/init(numberOfVisibleItems:)
func NewScrubberProportionalLayoutWithNumberOfVisibleItems(numberOfVisibleItems int) NSScrubberProportionalLayout {
	instance := getNSScrubberProportionalLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNumberOfVisibleItems:"), numberOfVisibleItems)
	return NSScrubberProportionalLayoutFromID(rv)
}

// Initializes and returns a newly allocated proportional layout, configured
// to display the given number of items.
//
// numberOfVisibleItems: The number of items that should be visible in the scrubber at once.
//
// # Return Value
// 
// A properly initialized proportional layout object.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout/init(numberOfVisibleItems:)
func (s NSScrubberProportionalLayout) InitWithNumberOfVisibleItems(numberOfVisibleItems int) NSScrubberProportionalLayout {
	rv := objc.Send[NSScrubberProportionalLayout](s.ID, objc.Sel("initWithNumberOfVisibleItems:"), numberOfVisibleItems)
	return rv
}

// The number of items visible in the scrubber at once.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout/numberOfVisibleItems
func (s NSScrubberProportionalLayout) NumberOfVisibleItems() int {
	rv := objc.Send[int](s.ID, objc.Sel("numberOfVisibleItems"))
	return rv
}
func (s NSScrubberProportionalLayout) SetNumberOfVisibleItems(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setNumberOfVisibleItems:"), value)
}

