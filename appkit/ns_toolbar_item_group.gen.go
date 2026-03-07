// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSToolbarItemGroup] class.
var (
	_NSToolbarItemGroupClass     NSToolbarItemGroupClass
	_NSToolbarItemGroupClassOnce sync.Once
)

func getNSToolbarItemGroupClass() NSToolbarItemGroupClass {
	_NSToolbarItemGroupClassOnce.Do(func() {
		_NSToolbarItemGroupClass = NSToolbarItemGroupClass{class: objc.GetClass("NSToolbarItemGroup")}
	})
	return _NSToolbarItemGroupClass
}

// GetNSToolbarItemGroupClass returns the class object for NSToolbarItemGroup.
func GetNSToolbarItemGroupClass() NSToolbarItemGroupClass {
	return getNSToolbarItemGroupClass()
}

type NSToolbarItemGroupClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSToolbarItemGroupClass) Alloc() NSToolbarItemGroup {
	rv := objc.Send[NSToolbarItemGroup](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A group of subitems in a toolbar item.
//
// # Overview
// 
// An [NSToolbarItemGroup] represents a collection set of subitems in a
// toolbar that the system displays based on available space and settings that
// you specify. The system uses the views and labels of the subitems, but the
// parent’s attributes take precedence. This differs from other
// [NSToolbarItem] objects because they’re attached — the user drags them
// together as a single item rather than separately.
// 
// If a subitem of the group has an action set on it, the group uses that
// action instead of its own when the user clicks or taps on that item. The
// system prefers the subitem’s action if it exists, otherwise it uses the
// group’s action.
// 
// To configure an instance of [NSToolbarItemGroup], you first create the
// individual toolbar subitems:
// 
// Then, you put them in a grouped item:
// 
// In this configuration, you get two grouped items, and two labels.
// 
// If you set a label on the parent item, you get two grouped items with one
// shared label:
// 
// If instead you set a view on the parent item, you get two labels with one
// shared view:
//
// # Working with subitems
//
//   - [NSToolbarItemGroup.Subitems]: The subitems of the grouped toolbar item.
//   - [NSToolbarItemGroup.SetSubitems]
//   - [NSToolbarItemGroup.SelectedIndex]: The index value for the most recently selected subitem of a grouped toolbar item.
//   - [NSToolbarItemGroup.SetSelectedIndex]
//   - [NSToolbarItemGroup.IsSelectedAtIndex]: Indicates whether a specified index is currently selected.
//   - [NSToolbarItemGroup.SetSelectedAtIndex]: Sets the selected state of a subitem in a grouped toolbar item.
//
// # Configuring grouped toolbar items
//
//   - [NSToolbarItemGroup.ControlRepresentation]: A value that represents how a toolbar displays a grouped toolbar item.
//   - [NSToolbarItemGroup.SetControlRepresentation]
//   - [NSToolbarItemGroup.SelectionMode]: The selection mode of the grouped toolbar item.
//   - [NSToolbarItemGroup.SetSelectionMode]
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup
type NSToolbarItemGroup struct {
	NSToolbarItem
}

// NSToolbarItemGroupFromID constructs a [NSToolbarItemGroup] from an objc.ID.
//
// A group of subitems in a toolbar item.
func NSToolbarItemGroupFromID(id objc.ID) NSToolbarItemGroup {
	return NSToolbarItemGroup{NSToolbarItem: NSToolbarItemFromID(id)}
}
// NOTE: NSToolbarItemGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSToolbarItemGroup] class.
//
// # Working with subitems
//
//   - [INSToolbarItemGroup.Subitems]: The subitems of the grouped toolbar item.
//   - [INSToolbarItemGroup.SetSubitems]
//   - [INSToolbarItemGroup.SelectedIndex]: The index value for the most recently selected subitem of a grouped toolbar item.
//   - [INSToolbarItemGroup.SetSelectedIndex]
//   - [INSToolbarItemGroup.IsSelectedAtIndex]: Indicates whether a specified index is currently selected.
//   - [INSToolbarItemGroup.SetSelectedAtIndex]: Sets the selected state of a subitem in a grouped toolbar item.
//
// # Configuring grouped toolbar items
//
//   - [INSToolbarItemGroup.ControlRepresentation]: A value that represents how a toolbar displays a grouped toolbar item.
//   - [INSToolbarItemGroup.SetControlRepresentation]
//   - [INSToolbarItemGroup.SelectionMode]: The selection mode of the grouped toolbar item.
//   - [INSToolbarItemGroup.SetSelectionMode]
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup
type INSToolbarItemGroup interface {
	INSToolbarItem

	// Topic: Working with subitems

	// The subitems of the grouped toolbar item.
	Subitems() []NSToolbarItem
	SetSubitems(value []NSToolbarItem)
	// The index value for the most recently selected subitem of a grouped toolbar item.
	SelectedIndex() int
	SetSelectedIndex(value int)
	// Indicates whether a specified index is currently selected.
	IsSelectedAtIndex(index int) bool
	// Sets the selected state of a subitem in a grouped toolbar item.
	SetSelectedAtIndex(selected bool, index int)

	// Topic: Configuring grouped toolbar items

	// A value that represents how a toolbar displays a grouped toolbar item.
	ControlRepresentation() NSToolbarItemGroupControlRepresentation
	SetControlRepresentation(value NSToolbarItemGroupControlRepresentation)
	// The selection mode of the grouped toolbar item.
	SelectionMode() NSToolbarItemGroupSelectionMode
	SetSelectionMode(value NSToolbarItemGroupSelectionMode)
}





// Init initializes the instance.
func (t NSToolbarItemGroup) Init() NSToolbarItemGroup {
	rv := objc.Send[NSToolbarItemGroup](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSToolbarItemGroup) Autorelease() NSToolbarItemGroup {
	rv := objc.Send[NSToolbarItemGroup](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSToolbarItemGroup creates a new NSToolbarItemGroup instance.
func NewNSToolbarItemGroup() NSToolbarItemGroup {
	class := getNSToolbarItemGroupClass()
	rv := objc.Send[NSToolbarItemGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a toolbar item with the specified identifier.
//
// itemIdentifier: The identifier for the toolbar item. You use this value to identify the
// item within your app, so you don’t need to localize it. For example, your
// toolbar delegate uses this value to identify the specific toolbar item.
//
// # Return Value
// 
// A new toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/init(itemIdentifier:)
func NewToolbarItemGroupWithItemIdentifier(itemIdentifier NSToolbarItemIdentifier) NSToolbarItemGroup {
	instance := getNSToolbarItemGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItemIdentifier:"), objc.String(string(itemIdentifier)))
	return NSToolbarItemGroupFromID(rv)
}


// Creates a grouped toolbar item with images.
//
// itemIdentifier: The identifier for the grouped toolbar item.
//
// images: An array of images to present as subitems in the grouped toolbar item.
//
// selectionMode: A value that indicates how the grouped toolbar item presents selections.
//
// labels: Labels that correspond to the specified images.
//
// target: The target that the toolbar calls upon selection.
// 
// If target is `nil`, the toolbar attempts to invoke the specified action on
// the first responder and, failing that, passes the action up the responder
// chain.
//
// action: The selector that the toolbar invokes on the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/init(itemIdentifier:images:selectionMode:labels:target:action:)
func NewToolbarItemGroupWithItemIdentifierImagesSelectionModeLabelsTargetAction(itemIdentifier NSToolbarItemIdentifier, images []objectivec.Object, selectionMode NSToolbarItemGroupSelectionMode, labels []string, target objectivec.IObject, action objc.SEL) NSToolbarItemGroup {
	rv := objc.Send[objc.ID](objc.ID(getNSToolbarItemGroupClass().class), objc.Sel("groupWithItemIdentifier:images:selectionMode:labels:target:action:"), objc.String(string(itemIdentifier)), objectivec.IObjectSliceToNSArray(images), selectionMode, objectivec.StringSliceToNSArray(labels), target, action)
	return NSToolbarItemGroupFromID(rv)
}


// Creates a grouped toolbar item with labels.
//
// itemIdentifier: The identifier for the grouped toolbar item.
//
// titles: An array of titles to present as subitems in the grouped toolbar item.
//
// selectionMode: A Boolean value that indicates how the grouped toolbar item presents
// selections.
//
// labels: Labels that correspond to the specified titles.
//
// target: The target that the toolbar calls upon selection.
// 
// If target is `nil`, the toolbar attempts to invoke the specified action on
// the first responder and, failing that, passes the action up the responder
// chain.
//
// action: The selector that the toolbar invokes on the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/init(itemIdentifier:titles:selectionMode:labels:target:action:)
func NewToolbarItemGroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction(itemIdentifier NSToolbarItemIdentifier, titles []string, selectionMode NSToolbarItemGroupSelectionMode, labels []string, target objectivec.IObject, action objc.SEL) NSToolbarItemGroup {
	rv := objc.Send[objc.ID](objc.ID(getNSToolbarItemGroupClass().class), objc.Sel("groupWithItemIdentifier:titles:selectionMode:labels:target:action:"), objc.String(string(itemIdentifier)), objectivec.StringSliceToNSArray(titles), selectionMode, objectivec.StringSliceToNSArray(labels), target, action)
	return NSToolbarItemGroupFromID(rv)
}







// Indicates whether a specified index is currently selected.
//
// index: The index of the subitems in a grouped toolbar item.
//
// # Return Value
// 
// A Boolean value indicating whether the specified index is currently
// selected.
//
// # Discussion
// 
// Use this method when you specify the
// [NSToolbarItemGroup.SelectionMode.selectAny] selection mode for the grouped
// toolbar item to determine which subitems are currently selected.
//
// [NSToolbarItemGroup.SelectionMode.selectAny]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum/selectAny
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/isSelected(at:)
func (t NSToolbarItemGroup) IsSelectedAtIndex(index int) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSelectedAtIndex:"), index)
	return rv
}

// Sets the selected state of a subitem in a grouped toolbar item.
//
// selected: If `true`, indicates whether to select a subitem, `false` otherwise.
//
// index: The index location of the subitem.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/setSelected(_:at:)
func (t NSToolbarItemGroup) SetSelectedAtIndex(selected bool, index int) {
	objc.Send[objc.ID](t.ID, objc.Sel("setSelected:atIndex:"), selected, index)
}












// The subitems of the grouped toolbar item.
//
// # Discussion
// 
// By default, an [NSToolbarItemGroup] instance has an empty array of
// subitems.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/subitems
func (t NSToolbarItemGroup) Subitems() []NSToolbarItem {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("subitems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSToolbarItem {
		return NSToolbarItemFromID(id)
	})
}
func (t NSToolbarItemGroup) SetSubitems(value []NSToolbarItem) {
	objc.Send[struct{}](t.ID, objc.Sel("setSubitems:"), objectivec.IObjectSliceToNSArray(value))
}



// The index value for the most recently selected subitem of a grouped toolbar
// item.
//
// # Discussion
// 
// When using the [NSToolbarItemGroup.SelectionMode.selectAny] or
// [NSToolbarItemGroup.SelectionMode.momentary] selection mode, don’t assume
// that this value represents the selected subitem. This method returns the
// index of the most recently selected subitem.
// 
// To determine if a specific subitem of a grouped toolbar item is selected,
// use the [IsSelectedAtIndex] method.
//
// [NSToolbarItemGroup.SelectionMode.momentary]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum/momentary
// [NSToolbarItemGroup.SelectionMode.selectAny]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum/selectAny
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/selectedIndex
func (t NSToolbarItemGroup) SelectedIndex() int {
	rv := objc.Send[int](t.ID, objc.Sel("selectedIndex"))
	return rv
}
func (t NSToolbarItemGroup) SetSelectedIndex(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectedIndex:"), value)
}



// A value that represents how a toolbar displays a grouped toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/controlRepresentation-swift.property
func (t NSToolbarItemGroup) ControlRepresentation() NSToolbarItemGroupControlRepresentation {
	rv := objc.Send[NSToolbarItemGroupControlRepresentation](t.ID, objc.Sel("controlRepresentation"))
	return NSToolbarItemGroupControlRepresentation(rv)
}
func (t NSToolbarItemGroup) SetControlRepresentation(value NSToolbarItemGroupControlRepresentation) {
	objc.Send[struct{}](t.ID, objc.Sel("setControlRepresentation:"), value)
}



// The selection mode of the grouped toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/selectionMode-swift.property
func (t NSToolbarItemGroup) SelectionMode() NSToolbarItemGroupSelectionMode {
	rv := objc.Send[NSToolbarItemGroupSelectionMode](t.ID, objc.Sel("selectionMode"))
	return NSToolbarItemGroupSelectionMode(rv)
}
func (t NSToolbarItemGroup) SetSelectionMode(value NSToolbarItemGroupSelectionMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectionMode:"), value)
}



























