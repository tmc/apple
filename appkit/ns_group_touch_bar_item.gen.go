// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSGroupTouchBarItem] class.
var (
	_NSGroupTouchBarItemClass     NSGroupTouchBarItemClass
	_NSGroupTouchBarItemClassOnce sync.Once
)

func getNSGroupTouchBarItemClass() NSGroupTouchBarItemClass {
	_NSGroupTouchBarItemClassOnce.Do(func() {
		_NSGroupTouchBarItemClass = NSGroupTouchBarItemClass{class: objc.GetClass("NSGroupTouchBarItem")}
	})
	return _NSGroupTouchBarItemClass
}

// GetNSGroupTouchBarItemClass returns the class object for NSGroupTouchBarItem.
func GetNSGroupTouchBarItemClass() NSGroupTouchBarItemClass {
	return getNSGroupTouchBarItemClass()
}

type NSGroupTouchBarItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSGroupTouchBarItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGroupTouchBarItemClass) Alloc() NSGroupTouchBarItem {
	rv := objc.Send[NSGroupTouchBarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A bar item that provides a bar to contain other items.
//
// # Configuring groups
//
//   - [NSGroupTouchBarItem.GroupUserInterfaceLayoutDirection]: The user interface direction that controls the layout order of the items.
//   - [NSGroupTouchBarItem.SetGroupUserInterfaceLayoutDirection]
//
// # Configuring item width
//
//   - [NSGroupTouchBarItem.PrefersEqualWidths]: A Boolean value that specifies that items should have equal widths when possible.
//   - [NSGroupTouchBarItem.SetPrefersEqualWidths]
//   - [NSGroupTouchBarItem.PreferredItemWidth]: The preferred width for items in the group.
//   - [NSGroupTouchBarItem.SetPreferredItemWidth]
//
// # Configuring item compression
//
//   - [NSGroupTouchBarItem.EffectiveCompressionOptions]: The compression options that are currently active on the group.
//   - [NSGroupTouchBarItem.PrioritizedCompressionOptions]: The allowed compression options, in the order they should be applied.
//   - [NSGroupTouchBarItem.SetPrioritizedCompressionOptions]
//
// See: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem
type NSGroupTouchBarItem struct {
	NSTouchBarItem
}

// NSGroupTouchBarItemFromID constructs a [NSGroupTouchBarItem] from an objc.ID.
//
// A bar item that provides a bar to contain other items.
func NSGroupTouchBarItemFromID(id objc.ID) NSGroupTouchBarItem {
	return NSGroupTouchBarItem{NSTouchBarItem: NSTouchBarItemFromID(id)}
}

// NOTE: NSGroupTouchBarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSGroupTouchBarItem] class.
//
// # Configuring groups
//
//   - [INSGroupTouchBarItem.GroupUserInterfaceLayoutDirection]: The user interface direction that controls the layout order of the items.
//   - [INSGroupTouchBarItem.SetGroupUserInterfaceLayoutDirection]
//
// # Configuring item width
//
//   - [INSGroupTouchBarItem.PrefersEqualWidths]: A Boolean value that specifies that items should have equal widths when possible.
//   - [INSGroupTouchBarItem.SetPrefersEqualWidths]
//   - [INSGroupTouchBarItem.PreferredItemWidth]: The preferred width for items in the group.
//   - [INSGroupTouchBarItem.SetPreferredItemWidth]
//
// # Configuring item compression
//
//   - [INSGroupTouchBarItem.EffectiveCompressionOptions]: The compression options that are currently active on the group.
//   - [INSGroupTouchBarItem.PrioritizedCompressionOptions]: The allowed compression options, in the order they should be applied.
//   - [INSGroupTouchBarItem.SetPrioritizedCompressionOptions]
//
// See: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem
type INSGroupTouchBarItem interface {
	INSTouchBarItem

	// Topic: Configuring groups

	// The user interface direction that controls the layout order of the items.
	GroupUserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection
	SetGroupUserInterfaceLayoutDirection(value NSUserInterfaceLayoutDirection)

	// Topic: Configuring item width

	// A Boolean value that specifies that items should have equal widths when possible.
	PrefersEqualWidths() bool
	SetPrefersEqualWidths(value bool)
	// The preferred width for items in the group.
	PreferredItemWidth() float64
	SetPreferredItemWidth(value float64)

	// Topic: Configuring item compression

	// The compression options that are currently active on the group.
	EffectiveCompressionOptions() INSUserInterfaceCompressionOptions
	// The allowed compression options, in the order they should be applied.
	PrioritizedCompressionOptions() []NSUserInterfaceCompressionOptions
	SetPrioritizedCompressionOptions(value []NSUserInterfaceCompressionOptions)
}

// Init initializes the instance.
func (g NSGroupTouchBarItem) Init() NSGroupTouchBarItem {
	rv := objc.Send[NSGroupTouchBarItem](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGroupTouchBarItem) Autorelease() NSGroupTouchBarItem {
	rv := objc.Send[NSGroupTouchBarItem](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGroupTouchBarItem creates a new NSGroupTouchBarItem instance.
func NewNSGroupTouchBarItem() NSGroupTouchBarItem {
	class := getNSGroupTouchBarItemClass()
	rv := objc.Send[NSGroupTouchBarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a group item configured to match system alerts.
//
// # Discussion
//
// You can control spacing between items, but it is recommended to use
// [fixedSpaceLarge] to maintain consistency.
//
// The [GroupUserInterfaceLayoutDirection] is set to match the application’s
// [UserInterfaceLayoutDirection].
//
// See: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(alertStyleWithIdentifier:)
//
// [fixedSpaceLarge]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Identifier-swift.struct/fixedSpaceLarge
func NewGroupTouchBarItemAlertStyleGroupItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSGroupTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(getNSGroupTouchBarItemClass().class), objc.Sel("alertStyleGroupItemWithIdentifier:"), objc.String(string(identifier)))
	return NSGroupTouchBarItemFromID(rv)
}

// Initializes and returns a group item whose bar is constructed from the
// supplied items.
//
// See: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:)
func NewGroupTouchBarItemGroupItemWithIdentifierItems(identifier NSTouchBarItemIdentifier, items []NSTouchBarItem) NSGroupTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(getNSGroupTouchBarItemClass().class), objc.Sel("groupItemWithIdentifier:items:"), objc.String(string(identifier)), objectivec.IObjectSliceToNSArray(items))
	return NSGroupTouchBarItemFromID(rv)
}

// Initializes and returns a group item whose bar is constructed from the
// supplied items, and with the specified compression options.
//
// # Discussion
//
// Use this initializer to specify which compression options are applied to
// the group item. The system applies options in the following default order:
// `breakEqualWidths`, `reduceMetrics`, `hideText`, `hideImages`.
//
// If you want to use non-standard compression options, add them by using the
// [PrioritizedCompressionOptions] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/init(identifier:items:allowedCompressionOptions:)
func NewGroupTouchBarItemGroupItemWithIdentifierItemsAllowedCompressionOptions(identifier NSTouchBarItemIdentifier, items []NSTouchBarItem, allowedCompressionOptions INSUserInterfaceCompressionOptions) NSGroupTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(getNSGroupTouchBarItemClass().class), objc.Sel("groupItemWithIdentifier:items:allowedCompressionOptions:"), objc.String(string(identifier)), objectivec.IObjectSliceToNSArray(items), allowedCompressionOptions)
	return NSGroupTouchBarItemFromID(rv)
}

// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewGroupTouchBarItemWithCoder(coder foundation.INSCoder) NSGroupTouchBarItem {
	instance := getNSGroupTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSGroupTouchBarItemFromID(rv)
}

// Creates a new item with the specified identifier.
//
// # Discussion
//
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewGroupTouchBarItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSGroupTouchBarItem {
	instance := getNSGroupTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSGroupTouchBarItemFromID(rv)
}

// The user interface direction that controls the layout order of the items.
//
// # Discussion
//
// The default value is [NSUserInterfaceLayoutDirectionLeftToRight].
//
// If you want the order of the items in the group to respect the user’s
// preferred layout, set this property to the value of
// [UserInterfaceLayoutDirection] on the [NSApplication].
//
// See: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/groupUserInterfaceLayoutDirection
func (g NSGroupTouchBarItem) GroupUserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection {
	rv := objc.Send[NSUserInterfaceLayoutDirection](g.ID, objc.Sel("groupUserInterfaceLayoutDirection"))
	return NSUserInterfaceLayoutDirection(rv)
}
func (g NSGroupTouchBarItem) SetGroupUserInterfaceLayoutDirection(value NSUserInterfaceLayoutDirection) {
	objc.Send[struct{}](g.ID, objc.Sel("setGroupUserInterfaceLayoutDirection:"), value)
}

// A Boolean value that specifies that items should have equal widths when
// possible.
//
// # Discussion
//
// When true, items in the [GroupTouchBar] are sized to have equal widths when
// possible.
//
// The default value is false.
//
// See: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/prefersEqualWidths
func (g NSGroupTouchBarItem) PrefersEqualWidths() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("prefersEqualWidths"))
	return rv
}
func (g NSGroupTouchBarItem) SetPrefersEqualWidths(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setPrefersEqualWidths:"), value)
}

// The preferred width for items in the group.
//
// # Discussion
//
// This width applies when [PrefersEqualWidths] is true.
//
// This is the width that items are set to if there is enough room, and if the
// items don’t clip.
//
// This value is ignored if it is negative. The default value is `-1`.
//
// See: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/preferredItemWidth
func (g NSGroupTouchBarItem) PreferredItemWidth() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("preferredItemWidth"))
	return rv
}
func (g NSGroupTouchBarItem) SetPreferredItemWidth(value float64) {
	objc.Send[struct{}](g.ID, objc.Sel("setPreferredItemWidth:"), value)
}

// The compression options that are currently active on the group.
//
// See: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/effectiveCompressionOptions
func (g NSGroupTouchBarItem) EffectiveCompressionOptions() INSUserInterfaceCompressionOptions {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("effectiveCompressionOptions"))
	return NSUserInterfaceCompressionOptionsFromID(objc.ID(rv))
}

// The allowed compression options, in the order they should be applied.
//
// # Discussion
//
// Use this property when you want to control the order of the system
// compression options, or if you want to use custom compression options.
//
// The default value is an array containing all standard AppKit options, in
// the AppKit-defined order.
//
// See: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem/prioritizedCompressionOptions
func (g NSGroupTouchBarItem) PrioritizedCompressionOptions() []NSUserInterfaceCompressionOptions {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("prioritizedCompressionOptions"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSUserInterfaceCompressionOptions {
		return NSUserInterfaceCompressionOptionsFromID(id)
	})
}
func (g NSGroupTouchBarItem) SetPrioritizedCompressionOptions(value []NSUserInterfaceCompressionOptions) {
	objc.Send[struct{}](g.ID, objc.Sel("setPrioritizedCompressionOptions:"), objectivec.IObjectSliceToNSArray(value))
}
