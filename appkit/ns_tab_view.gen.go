// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTabView] class.
var (
	_NSTabViewClass     NSTabViewClass
	_NSTabViewClassOnce sync.Once
)

func getNSTabViewClass() NSTabViewClass {
	_NSTabViewClassOnce.Do(func() {
		_NSTabViewClass = NSTabViewClass{class: objc.GetClass("NSTabView")}
	})
	return _NSTabViewClass
}

// GetNSTabViewClass returns the class object for NSTabView.
func GetNSTabViewClass() NSTabViewClass {
	return getNSTabViewClass()
}

type NSTabViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTabViewClass) Alloc() NSTabView {
	rv := objc.Send[NSTabView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A multipage interface that displays one page at a time.
//
// # Overview
// 
// A tab view contains a row of tabs that give the appearance of folder tabs,
// as shown in the [Figure 1]. The user selects the desired page by clicking
// the appropriate tab or using the arrow keys to move between pages. Each
// page displays a view hierarchy provided by your app.
// 
// [media-2555818]
//
// # Handling the Selection of Tabs
//
//   - [NSTabView.Delegate]: The tab view’s delegate.
//   - [NSTabView.SetDelegate]
//
// # Adding and Removing Tabs
//
//   - [NSTabView.AddTabViewItem]: Adds the specified tab item.
//   - [NSTabView.InsertTabViewItemAtIndex]: Inserts the specified item into the tab view’s array of tab view items at the specified index.
//   - [NSTabView.RemoveTabViewItem]: Removes the specified item from the tab view’s array of tab view items.
//
// # Accessing Tabs
//
//   - [NSTabView.IndexOfTabViewItem]: Returns the index of the specified item in the tab view.
//   - [NSTabView.IndexOfTabViewItemWithIdentifier]: Returns the index of the item that matches the specified identifier or [NSNotFound] if the item is not found.
//   - [NSTabView.NumberOfTabViewItems]: The number of items in the tab view’s array of tab view items.
//   - [NSTabView.TabViewItemAtIndex]: Returns the tab view item at `index` in the tab view’s array of items.
//   - [NSTabView.TabViewItems]: The tab view’s array of tab view items.
//   - [NSTabView.SetTabViewItems]
//
// # Configuring the Tab Attributes
//
//   - [NSTabView.TabViewType]: The tab type to display the tabs.
//   - [NSTabView.SetTabViewType]
//   - [NSTabView.TabPosition]
//   - [NSTabView.SetTabPosition]
//   - [NSTabView.TabViewBorderType]
//   - [NSTabView.SetTabViewBorderType]
//
// # Selecting a Tab
//
//   - [NSTabView.SelectFirstTabViewItem]: This action method selects the first tab view item.
//   - [NSTabView.SelectLastTabViewItem]: This action method selects the last tab view item.
//   - [NSTabView.SelectNextTabViewItem]: This action method selects the next tab view item in the sequence.
//   - [NSTabView.SelectPreviousTabViewItem]: This action method selects the previous tab view item in the sequence.
//   - [NSTabView.SelectTabViewItem]: Selects the specified tab view item.
//   - [NSTabView.SelectTabViewItemAtIndex]: Selects the tab view item specified by `index`.
//   - [NSTabView.SelectTabViewItemWithIdentifier]: Selects the tab view item specified by `identifier`.
//   - [NSTabView.SelectedTabViewItem]: The tab view item for the currently selected tab.
//   - [NSTabView.TakeSelectedTabViewItemFromSender]: Sets the selected tab view item to the selected item obtained from the sender.
//
// # Modifying the Font
//
//   - [NSTabView.Font]: The font used for the tab view’s label text.
//   - [NSTabView.SetFont]
//
// # Manipulating the Background
//
//   - [NSTabView.DrawsBackground]: A Boolean value that indicates if the tab view draws a background color when its type is [NSNoTabsNoBorder].
//   - [NSTabView.SetDrawsBackground]
//
// # Determining the Size
//
//   - [NSTabView.MinimumSize]: The minimum size necessary for the tab view to display tabs in a useful way.
//   - [NSTabView.ContentRect]: The rectangle describing the content area of the tab view.
//   - [NSTabView.ControlSize]: The size of the tab view.
//   - [NSTabView.SetControlSize]
//
// # Truncating Tab Labels
//
//   - [NSTabView.AllowsTruncatedLabels]: A Boolean value that indicates if the tab view allows truncating for labels that don’t fit on a tab.
//   - [NSTabView.SetAllowsTruncatedLabels]
//
// # Event Handling
//
//   - [NSTabView.TabViewItemAtPoint]: Returns the tab view item at the specified point.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView
type NSTabView struct {
	NSView
}

// NSTabViewFromID constructs a [NSTabView] from an objc.ID.
//
// A multipage interface that displays one page at a time.
func NSTabViewFromID(id objc.ID) NSTabView {
	return NSTabView{NSView: NSViewFromID(id)}
}
// NOTE: NSTabView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTabView] class.
//
// # Handling the Selection of Tabs
//
//   - [INSTabView.Delegate]: The tab view’s delegate.
//   - [INSTabView.SetDelegate]
//
// # Adding and Removing Tabs
//
//   - [INSTabView.AddTabViewItem]: Adds the specified tab item.
//   - [INSTabView.InsertTabViewItemAtIndex]: Inserts the specified item into the tab view’s array of tab view items at the specified index.
//   - [INSTabView.RemoveTabViewItem]: Removes the specified item from the tab view’s array of tab view items.
//
// # Accessing Tabs
//
//   - [INSTabView.IndexOfTabViewItem]: Returns the index of the specified item in the tab view.
//   - [INSTabView.IndexOfTabViewItemWithIdentifier]: Returns the index of the item that matches the specified identifier or [NSNotFound] if the item is not found.
//   - [INSTabView.NumberOfTabViewItems]: The number of items in the tab view’s array of tab view items.
//   - [INSTabView.TabViewItemAtIndex]: Returns the tab view item at `index` in the tab view’s array of items.
//   - [INSTabView.TabViewItems]: The tab view’s array of tab view items.
//   - [INSTabView.SetTabViewItems]
//
// # Configuring the Tab Attributes
//
//   - [INSTabView.TabViewType]: The tab type to display the tabs.
//   - [INSTabView.SetTabViewType]
//   - [INSTabView.TabPosition]
//   - [INSTabView.SetTabPosition]
//   - [INSTabView.TabViewBorderType]
//   - [INSTabView.SetTabViewBorderType]
//
// # Selecting a Tab
//
//   - [INSTabView.SelectFirstTabViewItem]: This action method selects the first tab view item.
//   - [INSTabView.SelectLastTabViewItem]: This action method selects the last tab view item.
//   - [INSTabView.SelectNextTabViewItem]: This action method selects the next tab view item in the sequence.
//   - [INSTabView.SelectPreviousTabViewItem]: This action method selects the previous tab view item in the sequence.
//   - [INSTabView.SelectTabViewItem]: Selects the specified tab view item.
//   - [INSTabView.SelectTabViewItemAtIndex]: Selects the tab view item specified by `index`.
//   - [INSTabView.SelectTabViewItemWithIdentifier]: Selects the tab view item specified by `identifier`.
//   - [INSTabView.SelectedTabViewItem]: The tab view item for the currently selected tab.
//   - [INSTabView.TakeSelectedTabViewItemFromSender]: Sets the selected tab view item to the selected item obtained from the sender.
//
// # Modifying the Font
//
//   - [INSTabView.Font]: The font used for the tab view’s label text.
//   - [INSTabView.SetFont]
//
// # Manipulating the Background
//
//   - [INSTabView.DrawsBackground]: A Boolean value that indicates if the tab view draws a background color when its type is [NSNoTabsNoBorder].
//   - [INSTabView.SetDrawsBackground]
//
// # Determining the Size
//
//   - [INSTabView.MinimumSize]: The minimum size necessary for the tab view to display tabs in a useful way.
//   - [INSTabView.ContentRect]: The rectangle describing the content area of the tab view.
//   - [INSTabView.ControlSize]: The size of the tab view.
//   - [INSTabView.SetControlSize]
//
// # Truncating Tab Labels
//
//   - [INSTabView.AllowsTruncatedLabels]: A Boolean value that indicates if the tab view allows truncating for labels that don’t fit on a tab.
//   - [INSTabView.SetAllowsTruncatedLabels]
//
// # Event Handling
//
//   - [INSTabView.TabViewItemAtPoint]: Returns the tab view item at the specified point.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView
type INSTabView interface {
	INSView

	// Topic: Handling the Selection of Tabs

	// The tab view’s delegate.
	Delegate() NSTabViewDelegate
	SetDelegate(value NSTabViewDelegate)

	// Topic: Adding and Removing Tabs

	// Adds the specified tab item.
	AddTabViewItem(tabViewItem INSTabViewItem)
	// Inserts the specified item into the tab view’s array of tab view items at the specified index.
	InsertTabViewItemAtIndex(tabViewItem INSTabViewItem, index int)
	// Removes the specified item from the tab view’s array of tab view items.
	RemoveTabViewItem(tabViewItem INSTabViewItem)

	// Topic: Accessing Tabs

	// Returns the index of the specified item in the tab view.
	IndexOfTabViewItem(tabViewItem INSTabViewItem) int
	// Returns the index of the item that matches the specified identifier or [NSNotFound] if the item is not found.
	IndexOfTabViewItemWithIdentifier(identifier objectivec.IObject) int
	// The number of items in the tab view’s array of tab view items.
	NumberOfTabViewItems() int
	// Returns the tab view item at `index` in the tab view’s array of items.
	TabViewItemAtIndex(index int) INSTabViewItem
	// The tab view’s array of tab view items.
	TabViewItems() []NSTabViewItem
	SetTabViewItems(value []NSTabViewItem)

	// Topic: Configuring the Tab Attributes

	// The tab type to display the tabs.
	TabViewType() NSTabViewType
	SetTabViewType(value NSTabViewType)
	TabPosition() NSTabPosition
	SetTabPosition(value NSTabPosition)
	TabViewBorderType() NSTabViewBorderType
	SetTabViewBorderType(value NSTabViewBorderType)

	// Topic: Selecting a Tab

	// This action method selects the first tab view item.
	SelectFirstTabViewItem(sender objectivec.IObject)
	// This action method selects the last tab view item.
	SelectLastTabViewItem(sender objectivec.IObject)
	// This action method selects the next tab view item in the sequence.
	SelectNextTabViewItem(sender objectivec.IObject)
	// This action method selects the previous tab view item in the sequence.
	SelectPreviousTabViewItem(sender objectivec.IObject)
	// Selects the specified tab view item.
	SelectTabViewItem(tabViewItem INSTabViewItem)
	// Selects the tab view item specified by `index`.
	SelectTabViewItemAtIndex(index int)
	// Selects the tab view item specified by `identifier`.
	SelectTabViewItemWithIdentifier(identifier objectivec.IObject)
	// The tab view item for the currently selected tab.
	SelectedTabViewItem() INSTabViewItem
	// Sets the selected tab view item to the selected item obtained from the sender.
	TakeSelectedTabViewItemFromSender(sender objectivec.IObject)

	// Topic: Modifying the Font

	// The font used for the tab view’s label text.
	Font() NSFont
	SetFont(value NSFont)

	// Topic: Manipulating the Background

	// A Boolean value that indicates if the tab view draws a background color when its type is [NSNoTabsNoBorder].
	DrawsBackground() bool
	SetDrawsBackground(value bool)

	// Topic: Determining the Size

	// The minimum size necessary for the tab view to display tabs in a useful way.
	MinimumSize() corefoundation.CGSize
	// The rectangle describing the content area of the tab view.
	ContentRect() corefoundation.CGRect
	// The size of the tab view.
	ControlSize() NSControlSize
	SetControlSize(value NSControlSize)

	// Topic: Truncating Tab Labels

	// A Boolean value that indicates if the tab view allows truncating for labels that don’t fit on a tab.
	AllowsTruncatedLabels() bool
	SetAllowsTruncatedLabels(value bool)

	// Topic: Event Handling

	// Returns the tab view item at the specified point.
	TabViewItemAtPoint(point corefoundation.CGPoint) INSTabViewItem
}

// Init initializes the instance.
func (t NSTabView) Init() NSTabView {
	rv := objc.Send[NSTabView](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTabView) Autorelease() NSTabView {
	rv := objc.Send[NSTabView](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTabView creates a new NSTabView instance.
func NewNSTabView() NSTabView {
	class := getNSTabViewClass()
	rv := objc.Send[NSTabView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a view using from data in the specified coder object.
//
// coder: The coder object that contains the view’s configuration details.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewTabViewWithCoder(coder foundation.INSCoder) NSTabView {
	instance := getNSTabViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTabViewFromID(rv)
}

// Initializes and returns a newly allocated [NSView] object with a specified
// frame rectangle.
//
// frameRect: The frame rectangle for the created view object.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// # Discussion
// 
// Insert the view into your window’s view hieararchy before you can do
// anything with it. This method is the designated initializer for the
// [NSView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewTabViewWithFrame(frameRect corefoundation.CGRect) NSTabView {
	instance := getNSTabViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSTabViewFromID(rv)
}

// Adds the specified tab item.
//
// tabViewItem: The tab view item to be added.
//
// # Discussion
// 
// The item is added at the end of the array of tab items, so the new tab
// appears on the right side of the view. If the delegate supports it, it
// invokes the delegate’s [TabViewDidChangeNumberOfTabViewItems] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/addTabViewItem(_:)
func (t NSTabView) AddTabViewItem(tabViewItem INSTabViewItem) {
	objc.Send[objc.ID](t.ID, objc.Sel("addTabViewItem:"), tabViewItem)
}
// Inserts the specified item into the tab view’s array of tab view items at
// the specified index.
//
// tabViewItem: The tab view item to be added.
//
// index: The index at which to insert the tab view item. The `index` parameter is
// zero-based.
//
// # Discussion
// 
// If there is a delegate and the delegate supports it, sends the delegate the
// [TabViewDidChangeNumberOfTabViewItems] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/insertTabViewItem(_:at:)
func (t NSTabView) InsertTabViewItemAtIndex(tabViewItem INSTabViewItem, index int) {
	objc.Send[objc.ID](t.ID, objc.Sel("insertTabViewItem:atIndex:"), tabViewItem, index)
}
// Removes the specified item from the tab view’s array of tab view items.
//
// tabViewItem: The tab view item to be removed.
//
// # Discussion
// 
// If there is a delegate and the delegate supports it, sends the delegate the
// [TabViewDidChangeNumberOfTabViewItems] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/removeTabViewItem(_:)
func (t NSTabView) RemoveTabViewItem(tabViewItem INSTabViewItem) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeTabViewItem:"), tabViewItem)
}
// Returns the index of the specified item in the tab view.
//
// tabViewItem: The tab view item.
//
// # Return Value
// 
// The zero-based index of `tabViewItem`, or [NSNotFound] if the item is not
// found.
//
// # Discussion
// 
// The returned index is zero-based.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/indexOfTabViewItem(_:)
func (t NSTabView) IndexOfTabViewItem(tabViewItem INSTabViewItem) int {
	rv := objc.Send[int](t.ID, objc.Sel("indexOfTabViewItem:"), tabViewItem)
	return rv
}
// Returns the index of the item that matches the specified identifier or
// [NSNotFound] if the item is not found.
//
// identifier: The identifier of a tab view item.
//
// # Return Value
// 
// The zero-based index of the tab view item corresponding to `identifier`, or
// [NSNotFound] if the item is not found.
//
// # Discussion
// 
// The returned index is zero-based.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/indexOfTabViewItem(withIdentifier:)
func (t NSTabView) IndexOfTabViewItemWithIdentifier(identifier objectivec.IObject) int {
	rv := objc.Send[int](t.ID, objc.Sel("indexOfTabViewItemWithIdentifier:"), identifier)
	return rv
}
// Returns the tab view item at `index` in the tab view’s array of items.
//
// index: The index at which to insert the tab view item. The `index` parameter is
// zero-based.
//
// # Return Value
// 
// The tab view item at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItem(at:)-7r3at
func (t NSTabView) TabViewItemAtIndex(index int) INSTabViewItem {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tabViewItemAtIndex:"), index)
	return NSTabViewItemFromID(rv)
}
// This action method selects the first tab view item.
//
// sender: Typically the object that sent the message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/selectFirstTabViewItem(_:)
func (t NSTabView) SelectFirstTabViewItem(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("selectFirstTabViewItem:"), sender)
}
// This action method selects the last tab view item.
//
// sender: Typically the object that sent invoked the message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/selectLastTabViewItem(_:)
func (t NSTabView) SelectLastTabViewItem(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("selectLastTabViewItem:"), sender)
}
// This action method selects the next tab view item in the sequence.
//
// sender: Typically the object that sent the message.
//
// # Discussion
// 
// If the currently visible item is the last item in the sequence, this method
// does nothing, and the last pane remains displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/selectNextTabViewItem(_:)
func (t NSTabView) SelectNextTabViewItem(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("selectNextTabViewItem:"), sender)
}
// This action method selects the previous tab view item in the sequence.
//
// sender: Typically the object that sent the message.
//
// # Discussion
// 
// If the currently visible item is the first item in the sequence, this
// method does nothing, and the first pane remains displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/selectPreviousTabViewItem(_:)
func (t NSTabView) SelectPreviousTabViewItem(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("selectPreviousTabViewItem:"), sender)
}
// Selects the specified tab view item.
//
// tabViewItem: The tab item to select.
//
// # Discussion
// 
// If there is a delegate and the delegate supports it, sends the delegate the
// [TabViewShouldSelectTabViewItem] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(_:)
func (t NSTabView) SelectTabViewItem(tabViewItem INSTabViewItem) {
	objc.Send[objc.ID](t.ID, objc.Sel("selectTabViewItem:"), tabViewItem)
}
// Selects the tab view item specified by `index`.
//
// index: The index of the tab item to selected.
//
// # Discussion
// 
// The `index` parameter is base 0. If there is a delegate and the delegate
// supports it, sends the delegate the [TabViewShouldSelectTabViewItem]
// message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(at:)
func (t NSTabView) SelectTabViewItemAtIndex(index int) {
	objc.Send[objc.ID](t.ID, objc.Sel("selectTabViewItemAtIndex:"), index)
}
// Selects the tab view item specified by `identifier`.
//
// identifier: The identifier of the tab item to select.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/selectTabViewItem(withIdentifier:)
func (t NSTabView) SelectTabViewItemWithIdentifier(identifier objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("selectTabViewItemWithIdentifier:"), identifier)
}
// Sets the selected tab view item to the selected item obtained from the
// sender.
//
// sender: Typically the object that sent the message.
//
// # Discussion
// 
// If `sender` responds to the `indexOfSelectedItem` method, this method
// invokes that method and selects the tab view item at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/takeSelectedTabViewItemFromSender(_:)
func (t NSTabView) TakeSelectedTabViewItemFromSender(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("takeSelectedTabViewItemFromSender:"), sender)
}
// Returns the tab view item at the specified point.
//
// point: The hit point.
//
// # Return Value
// 
// The tab view item under the hit point, or `nil` if no tab view item is
// under that location.
//
// # Discussion
// 
// You can use this method to find a tab view item based on a user’s mouse
// click.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItem(at:)-8gnqw
func (t NSTabView) TabViewItemAtPoint(point corefoundation.CGPoint) INSTabViewItem {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tabViewItemAtPoint:"), point)
	return NSTabViewItemFromID(rv)
}

// The tab view’s delegate.
//
// # Discussion
// 
// The value of this property must conform to the [NSTabViewDelegate]
// protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/delegate
func (t NSTabView) Delegate() NSTabViewDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return NSTabViewDelegateObjectFromID(rv)
}
func (t NSTabView) SetDelegate(value NSTabViewDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}
// The number of items in the tab view’s array of tab view items.
//
// # Discussion
// 
// The default value of this property is 0.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/numberOfTabViewItems
func (t NSTabView) NumberOfTabViewItems() int {
	rv := objc.Send[int](t.ID, objc.Sel("numberOfTabViewItems"))
	return rv
}
// The tab view’s array of tab view items.
//
// # Discussion
// 
// A tab view keeps an array containing one tab view item for each tab in the
// view. The default value of this property is an empty array.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewItems
func (t NSTabView) TabViewItems() []NSTabViewItem {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("tabViewItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTabViewItem {
		return NSTabViewItemFromID(id)
	})
}
func (t NSTabView) SetTabViewItems(value []NSTabViewItem) {
	objc.Send[struct{}](t.ID, objc.Sel("setTabViewItems:"), objectivec.IObjectSliceToNSArray(value))
}
// The tab type to display the tabs.
//
// # Discussion
// 
// The supported values for this property are listed in [NSTabView.TabType].
// The default value of this property is
// [NSTabView.TabType.topTabsBezelBorder].
//
// [NSTabView.TabType.topTabsBezelBorder]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType/topTabsBezelBorder
// [NSTabView.TabType]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewType
func (t NSTabView) TabViewType() NSTabViewType {
	rv := objc.Send[NSTabViewType](t.ID, objc.Sel("tabViewType"))
	return NSTabViewType(rv)
}
func (t NSTabView) SetTabViewType(value NSTabViewType) {
	objc.Send[struct{}](t.ID, objc.Sel("setTabViewType:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTabView/tabPosition-swift.property
func (t NSTabView) TabPosition() NSTabPosition {
	rv := objc.Send[NSTabPosition](t.ID, objc.Sel("tabPosition"))
	return NSTabPosition(rv)
}
func (t NSTabView) SetTabPosition(value NSTabPosition) {
	objc.Send[struct{}](t.ID, objc.Sel("setTabPosition:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTabView/tabViewBorderType-swift.property
func (t NSTabView) TabViewBorderType() NSTabViewBorderType {
	rv := objc.Send[NSTabViewBorderType](t.ID, objc.Sel("tabViewBorderType"))
	return NSTabViewBorderType(rv)
}
func (t NSTabView) SetTabViewBorderType(value NSTabViewBorderType) {
	objc.Send[struct{}](t.ID, objc.Sel("setTabViewBorderType:"), value)
}
// The tab view item for the currently selected tab.
//
// # Discussion
// 
// If no item is selected, the value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/selectedTabViewItem
func (t NSTabView) SelectedTabViewItem() INSTabViewItem {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("selectedTabViewItem"))
	return NSTabViewItemFromID(objc.ID(rv))
}
// The font used for the tab view’s label text.
//
// # Discussion
// 
// The default value of this property is the message font of default size (see
// [MessageFontOfSize]), which is equivalent to the system font of default
// size. Tab height is adjusted automatically to accommodate a new font size.
// If the view allows truncating, tab labels are truncated as needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/font
func (t NSTabView) Font() NSFont {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("font"))
	return NSFontFromID(objc.ID(rv))
}
func (t NSTabView) SetFont(value NSFont) {
	objc.Send[struct{}](t.ID, objc.Sel("setFont:"), value)
}
// A Boolean value that indicates if the tab view draws a background color
// when its type is [NSNoTabsNoBorder].
//
// # Discussion
// 
// When the value of this property is [true], the tab view draws a background
// color when the its type is [NSNoTabsNoBorder], otherwise it does not. If
// the tab view has a bezeled border or a line border, the appropriate
// background for that border is used. The default value of this property is
// [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/drawsBackground
func (t NSTabView) DrawsBackground() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("drawsBackground"))
	return rv
}
func (t NSTabView) SetDrawsBackground(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setDrawsBackground:"), value)
}
// The minimum size necessary for the tab view to display tabs in a useful
// way.
//
// # Discussion
// 
// You can use the value of this property to limit how much a user can resize
// a tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/minimumSize
func (t NSTabView) MinimumSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("minimumSize"))
	return corefoundation.CGSize(rv)
}
// The rectangle describing the content area of the tab view.
//
// # Discussion
// 
// This area does not include the space required for the tab view’s tabs or
// borders (if any).
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/contentRect
func (t NSTabView) ContentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("contentRect"))
	return corefoundation.CGRect(rv)
}
// The size of the tab view.
//
// # Discussion
// 
// The valid values for this property are described inControl Sizes in
// [NSCell]. The default value of this property is [NSRegularControlSize].
//
// [NSRegularControlSize]: https://developer.apple.com/documentation/AppKit/NSRegularControlSize
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/controlSize
func (t NSTabView) ControlSize() NSControlSize {
	rv := objc.Send[NSControlSize](t.ID, objc.Sel("controlSize"))
	return NSControlSize(rv)
}
func (t NSTabView) SetControlSize(value NSControlSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setControlSize:"), value)
}
// A Boolean value that indicates if the tab view allows truncating for labels
// that don’t fit on a tab.
//
// # Discussion
// 
// When the value of this property is [true], the tab view allows truncating
// for labels that don’t fit on a tab, otherwise it does not. The default
// value is [true]. When truncating is allowed, the tab view inserts an
// ellipsis, if necessary, to fit a label in the tab.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTabView/allowsTruncatedLabels
func (t NSTabView) AllowsTruncatedLabels() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsTruncatedLabels"))
	return rv
}
func (t NSTabView) SetAllowsTruncatedLabels(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsTruncatedLabels:"), value)
}

