// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTabViewController] class.
var (
	_NSTabViewControllerClass     NSTabViewControllerClass
	_NSTabViewControllerClassOnce sync.Once
)

func getNSTabViewControllerClass() NSTabViewControllerClass {
	_NSTabViewControllerClassOnce.Do(func() {
		_NSTabViewControllerClass = NSTabViewControllerClass{class: objc.GetClass("NSTabViewController")}
	})
	return _NSTabViewControllerClass
}

// GetNSTabViewControllerClass returns the class object for NSTabViewController.
func GetNSTabViewControllerClass() NSTabViewControllerClass {
	return getNSTabViewControllerClass()
}

type NSTabViewControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTabViewControllerClass) Alloc() NSTabViewController {
	rv := objc.Send[NSTabViewController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A container view controller that manages a tab view interface, which
// organizes multiple pages of content but displays only one page at a time.
//
// # Overview
// 
// Each page of content is managed by a separate child view controller.
// Navigation between child view controllers is accomplished with the help of
// an [NSTabView] object, which the tab view controller manages. When the user
// selects a new tab, the tab view controller displays the content associated
// with the associated child view controller, replacing the previous content.
// 
// Each tab is represented by an [NSTabViewItem] object, which contains the
// name of the tab and stores a pointer to the child view controller that
// manages the tab’s content. Normally, you configure the tab view items at
// design time using Interface Builder, but you can also add them
// programmatically using the methods of this class. Always assign a child
// view controller to new tab view items before adding those items to the tab
// view interface.
// 
// Another way to add tabs programmatically is to add child view controllers
// directly to the tab view controller. When you call the
// [AddChildViewController] or [InsertChildViewControllerAtIndex] method of
// this class, the tab view controller automatically creates a default
// [NSTabViewItem] object for the specified view controller. You can fetch the
// newly created item using the [NSTabViewController.TabViewItemForViewController] method and
// configure it. Removing a child view controller with the
// [RemoveChildViewControllerAtIndex] method similarly removes the
// corresponding tab view item.
// 
// The tab view controller lazily loads the views associated with each child
// view controller, creating them only after the corresponding tab is
// selected. When the tab view controller’s view is first displayed, only
// the view for the initially selected tab is loaded.
// 
// The [NSTabViewController.TabStyle] property determines the appearance of the tab controls. A
// tab view controller can display a segmented control or display tabs in the
// window’s toolbar. You can also provide your own control for displaying
// tabs. The tab view controller automatically coordinates interactions
// between designated control and the corresponding [TabView] object.
//
// # Configuring the Tab View
//
//   - [NSTabViewController.TabStyle]: The style used to display the tabs.
//   - [NSTabViewController.SetTabStyle]
//   - [NSTabViewController.TabView]: The tab view that manages the views of the interface.
//   - [NSTabViewController.SetTabView]
//   - [NSTabViewController.TransitionOptions]: The animation options to use when switching between tabs.
//   - [NSTabViewController.SetTransitionOptions]
//   - [NSTabViewController.CanPropagateSelectedChildViewControllerTitle]: A Boolean value indicating whether the tab view controller gets its title from the selected child view controller.
//   - [NSTabViewController.SetCanPropagateSelectedChildViewControllerTitle]
//
// # Managing Tab View Items
//
//   - [NSTabViewController.TabViewItems]: The array of tab view items used to manage each of the child view controllers.
//   - [NSTabViewController.SetTabViewItems]
//   - [NSTabViewController.TabViewItemForViewController]: Returns the tab view item for the specified child view controller.
//   - [NSTabViewController.AddTabViewItem]: Adds the specified tab to the end of the tab view controller’s list of tabs.
//   - [NSTabViewController.InsertTabViewItemAtIndex]: Inserts a tab view into the tab view controller’s list of tabs.
//   - [NSTabViewController.RemoveTabViewItem]: Removes the specified tab view item from the tab view controller.
//   - [NSTabViewController.SelectedTabViewItemIndex]: The index of the selected tab.
//   - [NSTabViewController.SetSelectedTabViewItemIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController
type NSTabViewController struct {
	NSViewController
}

// NSTabViewControllerFromID constructs a [NSTabViewController] from an objc.ID.
//
// A container view controller that manages a tab view interface, which
// organizes multiple pages of content but displays only one page at a time.
func NSTabViewControllerFromID(id objc.ID) NSTabViewController {
	return NSTabViewController{NSViewController: NSViewControllerFromID(id)}
}
// NOTE: NSTabViewController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTabViewController] class.
//
// # Configuring the Tab View
//
//   - [INSTabViewController.TabStyle]: The style used to display the tabs.
//   - [INSTabViewController.SetTabStyle]
//   - [INSTabViewController.TabView]: The tab view that manages the views of the interface.
//   - [INSTabViewController.SetTabView]
//   - [INSTabViewController.TransitionOptions]: The animation options to use when switching between tabs.
//   - [INSTabViewController.SetTransitionOptions]
//   - [INSTabViewController.CanPropagateSelectedChildViewControllerTitle]: A Boolean value indicating whether the tab view controller gets its title from the selected child view controller.
//   - [INSTabViewController.SetCanPropagateSelectedChildViewControllerTitle]
//
// # Managing Tab View Items
//
//   - [INSTabViewController.TabViewItems]: The array of tab view items used to manage each of the child view controllers.
//   - [INSTabViewController.SetTabViewItems]
//   - [INSTabViewController.TabViewItemForViewController]: Returns the tab view item for the specified child view controller.
//   - [INSTabViewController.AddTabViewItem]: Adds the specified tab to the end of the tab view controller’s list of tabs.
//   - [INSTabViewController.InsertTabViewItemAtIndex]: Inserts a tab view into the tab view controller’s list of tabs.
//   - [INSTabViewController.RemoveTabViewItem]: Removes the specified tab view item from the tab view controller.
//   - [INSTabViewController.SelectedTabViewItemIndex]: The index of the selected tab.
//   - [INSTabViewController.SetSelectedTabViewItemIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController
type INSTabViewController interface {
	INSViewController
	NSTabViewDelegate
	NSToolbarDelegate

	// Topic: Configuring the Tab View

	// The style used to display the tabs.
	TabStyle() NSTabViewControllerTabStyle
	SetTabStyle(value NSTabViewControllerTabStyle)
	// The tab view that manages the views of the interface.
	TabView() INSTabView
	SetTabView(value INSTabView)
	// The animation options to use when switching between tabs.
	TransitionOptions() NSViewControllerTransitionOptions
	SetTransitionOptions(value NSViewControllerTransitionOptions)
	// A Boolean value indicating whether the tab view controller gets its title from the selected child view controller.
	CanPropagateSelectedChildViewControllerTitle() bool
	SetCanPropagateSelectedChildViewControllerTitle(value bool)

	// Topic: Managing Tab View Items

	// The array of tab view items used to manage each of the child view controllers.
	TabViewItems() []NSTabViewItem
	SetTabViewItems(value []NSTabViewItem)
	// Returns the tab view item for the specified child view controller.
	TabViewItemForViewController(viewController INSViewController) INSTabViewItem
	// Adds the specified tab to the end of the tab view controller’s list of tabs.
	AddTabViewItem(tabViewItem INSTabViewItem)
	// Inserts a tab view into the tab view controller’s list of tabs.
	InsertTabViewItemAtIndex(tabViewItem INSTabViewItem, index int)
	// Removes the specified tab view item from the tab view controller.
	RemoveTabViewItem(tabViewItem INSTabViewItem)
	// The index of the selected tab.
	SelectedTabViewItemIndex() int
	SetSelectedTabViewItemIndex(value int)

	// An array of view controllers that are hierarchical children of the view controller.
	Children() INSViewController
	SetChildren(value INSViewController)
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (t NSTabViewController) Init() NSTabViewController {
	rv := objc.Send[NSTabViewController](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTabViewController) Autorelease() NSTabViewController {
	rv := objc.Send[NSTabViewController](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTabViewController creates a new NSTabViewController instance.
func NewNSTabViewController() NSTabViewController {
	class := getNSTabViewControllerClass()
	rv := objc.Send[NSTabViewController](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(coder:)
func NewTabViewControllerWithCoder(coder foundation.INSCoder) NSTabViewController {
	instance := getNSTabViewControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTabViewControllerFromID(rv)
}


// Returns a view controller object initialized to the nib file in the
// specified bundle.
//
// nibNameOrNil: The name of the nib file, without any leading path information.
//
// nibBundleOrNil: The bundle in which to search for the nib file. If you specify `nil`, this
// method looks for the nib file in the main bundle.
//
// # Return Value
// 
// The initialized [NSViewController] object.
//
// # Discussion
// 
// The [NSViewController] object looks for the nib file in the bundle’s
// language-specific project directories first, followed by the Resources
// directory.
// 
// The specified nib file should typically have the class of the file’s
// owner set to [NSViewController], or a custom subclass, with the `view`
// outlet connected to a view.
// 
// If you pass in `nil` for `nibNameOrNil`, [NibName] returns `nil` and
// [LoadView] throws an exception; in this case you must set [View] before
// [View] is invoked, or override [LoadView].
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(nibName:bundle:)
func NewTabViewControllerWithNibNameBundle(nibNameOrNil NSNibName, nibBundleOrNil *foundation.NSBundle) NSTabViewController {
	instance := getNSTabViewControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNibName:bundle:"), objc.String(string(nibNameOrNil)), nibBundleOrNil)
	return NSTabViewControllerFromID(rv)
}







// Returns the tab view item for the specified child view controller.
//
// viewController: The child view controller whose tab view item you want.
//
// # Return Value
// 
// The tab view item associated with the view controller or `nil` if the view
// controller is not managed by the tab view controller.
//
// # Discussion
// 
// This method is a convenient way to map a tab view item to a newly added
// child view controller. When you add child view controllers using the
// [AddChildViewController] method, the tab view automatically controller
// creates a new tab view item. Use this method to fetch that tab view item
// and configure it.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabViewItem(for:)
func (t NSTabViewController) TabViewItemForViewController(viewController INSViewController) INSTabViewItem {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tabViewItemForViewController:"), viewController)
	return NSTabViewItemFromID(rv)
}

// Adds the specified tab to the end of the tab view controller’s list of
// tabs.
//
// tabViewItem: The tab view item to add. The tab view item must have an associated view
// controller. If this parameter is `nil` or if the tab view item does not
// have a view controller, this method raises an exception.
//
// # Discussion
// 
// Use this method to add new tabs to a tab view controller. This method adds
// the tab’s associated view controller as a child of the tab view
// controller, so you do not need to call the [AddChildViewController] method
// directly. The view for the new view controller is not loaded until its
// corresponding tab is selected by the user.
// 
// If you override this method, you must call `super` at some point in your
// implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/addTabViewItem(_:)
func (t NSTabViewController) AddTabViewItem(tabViewItem INSTabViewItem) {
	objc.Send[objc.ID](t.ID, objc.Sel("addTabViewItem:"), tabViewItem)
}

// Inserts a tab view into the tab view controller’s list of tabs.
//
// tabViewItem: The tab view item to insert. The tab view item must have an associated view
// controller. If this parameter is `nil` or if the tab view item does not
// have a view controller, this method throws an exception.
//
// index: The zero-based index at which to insert the tab. If the index is less than
// `0` or greater than the number of items currently in the array, this method
// raises an exception.
//
// # Discussion
// 
// Use this method to insert new tabs into the existing list of tabs. This
// method adds the tab’s associated view controller as a child of the tab
// view controller, so you do not need to call the [AddChildViewController]
// method directly. The view for the new view controller is not loaded until
// its corresponding tab is selected by the user.
// 
// Inserting a new tab updates the tab view interface and adjusts the value in
// the [SelectedTabViewItemIndex] property as needed.
// 
// If you override this method, you must call `super` at some point in your
// implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/insertTabViewItem(_:at:)
func (t NSTabViewController) InsertTabViewItemAtIndex(tabViewItem INSTabViewItem, index int) {
	objc.Send[objc.ID](t.ID, objc.Sel("insertTabViewItem:atIndex:"), tabViewItem, index)
}

// Removes the specified tab view item from the tab view controller.
//
// tabViewItem: The tab view item to remove. If this parameter is `nil` or the item does
// not belong to the tab view controller, this method throws an exception.
//
// # Discussion
// 
// Use this method to remove a tab view item from the tab view interface.
// Removing the item removes the corresponding view controller from the tab
// view controller’s list of child view controllers. If the removed tab view
// item is currently selected, the tab view controller selects the next item
// (or the previous item if there is no next item). Removing the last tab view
// item sets the [SelectedTabViewItemIndex] property to `-1`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/removeTabViewItem(_:)
func (t NSTabViewController) RemoveTabViewItem(tabViewItem INSTabViewItem) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeTabViewItem:"), tabViewItem)
}

// Asks the tab view controller if the specified tab should be selected.
//
// tabView: The tab view object making the request.
//
// tabViewItem: The tab view item to select.
//
// # Return Value
// 
// [true] if the tab should be selected or [false] if it should not be
// selected.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is a delegate method called by the [NSTabView] object when
// changes occur. Use it to dynamically determine whether a tab should be
// selected.
// 
// If you override this method, you must call `super` at some point in your
// implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabView(_:shouldSelect:)
func (t NSTabViewController) TabViewShouldSelectTabViewItem(tabView INSTabView, tabViewItem INSTabViewItem) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("tabView:shouldSelectTabViewItem:"), tabView, tabViewItem)
	return rv
}

// Informs the tab view controller that the specified tab is about to be
// selected.
//
// tabView: The tab view object whose tab is about to be selected.
//
// tabViewItem: The tab view item that will be selected.
//
// # Discussion
// 
// This method is a delegate method called by the [NSTabView] object when
// changes occur. Use it to update your UI or perform any tasks before a tab
// is selected.
// 
// If you override this method, you must call `super` at some point in your
// implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabView(_:willSelect:)
func (t NSTabViewController) TabViewWillSelectTabViewItem(tabView INSTabView, tabViewItem INSTabViewItem) {
	objc.Send[objc.ID](t.ID, objc.Sel("tabView:willSelectTabViewItem:"), tabView, tabViewItem)
}

// Informs the tab view controller that the specified tab was selected.
//
// tabView: The tab view object whose tab was selected.
//
// tabViewItem: The tab view item that was selected.
//
// # Discussion
// 
// This method is a delegate method called by the [NSTabView] object when
// changes occur. Use it to perform any necessary tasks after a tab is
// selected.
// 
// If you override this method, you must call `super` at some point in your
// implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabView(_:didSelect:)
func (t NSTabViewController) TabViewDidSelectTabViewItem(tabView INSTabView, tabViewItem INSTabViewItem) {
	objc.Send[objc.ID](t.ID, objc.Sel("tabView:didSelectTabViewItem:"), tabView, tabViewItem)
}

// Returns the toolbar item for the specified identifier.
//
// toolbar: The toolbar making the request.
//
// itemIdentifier: The identifier of the toolbar item being requested.
//
// flag: A Boolean indicating whether the item is inserted immediately into the
// toolbar. A value of [true] means the item is inserted into the toolbar. A
// value of [false] means the item is added to the toolbar’s configuration
// palette. The same item may be requested more than once with different
// values for this flag.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The requested toolbar item or `nil` to indicate that the specified item is
// not supported. When the same item is requested again, you may return the
// same [NSToolbarItem] object or a different one.
//
// # Discussion
// 
// This method is called for tab view interfaces that use the
// [NSTabViewController.TabStyle.toolbar] style. Use this method to create
// toolbar items for any custom identifiers you specified in the
// [ToolbarAllowedItemIdentifiers] and [ToolbarDefaultItemIdentifiers]
// methods.
// 
// If you override this method, you must call `super` at some point in your
// implementation. The default implementation of this method returns toolbar
// items for the tabs in the tab bar interface. The identifier for each
// toolbar item is the same as the identifier for the corresponding tab view
// item. Similarly, the toolbar item’s [Label], [Image] and [ToolTip]
// properties are bound to those of the corresponding tab view item.
//
// [NSTabViewController.TabStyle.toolbar]: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum/toolbar
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/toolbar(_:itemForItemIdentifier:willBeInsertedIntoToolbar:)
func (t NSTabViewController) ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar INSToolbar, itemIdentifier NSToolbarItemIdentifier, flag bool) INSToolbarItem {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"), toolbar, objc.String(string(itemIdentifier)), flag)
	return NSToolbarItemFromID(rv)
}

// Returns the array of identifier strings for the allowed toolbar items.
//
// toolbar: The toolbar making the request.
//
// # Return Value
// 
// An array of [NSString] objects, each of which contains an identifier for an
// available toolbar item. The array must contain all of the items returned by
// the [ToolbarDefaultItemIdentifiers] method.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// # Discussion
// 
// This method is called for tab view interfaces that use the
// [NSTabViewController.TabStyle.toolbar] style. Use this method to specify
// all possible items that may be included in the toolbar. The order of the
// items in the array is used to set their position in the toolbar
// configuration palette. If you include custom identifiers in the returned
// array, you must also override the
// [ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar] method to specify
// the content for those toolbar items.
// 
// If you override this method, you must call `super` at some point in your
// implementation. The default implementation of this method returns the
// identifiers for all toolbar items that correspond to tabs in the tab bar
// interface.
//
// [NSTabViewController.TabStyle.toolbar]: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum/toolbar
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/toolbarAllowedItemIdentifiers(_:)
func (t NSTabViewController) ToolbarAllowedItemIdentifiers(toolbar INSToolbar) []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("toolbarAllowedItemIdentifiers:"), toolbar)
	return objc.ConvertSliceToStrings(rv)
}

// Returns the array of identifier strings for the default toolbar items.
//
// toolbar: The toolbar making the request.
//
// # Return Value
// 
// An array of [NSString] objects, each of which contains an identifier for a
// toolbar item that is part of the default configuration. The order of items
// in the array is used to set the order of items in the toolbar.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// # Discussion
// 
// This method is called for tab view interfaces that use the
// [NSTabViewController.TabStyle.toolbar] style. Use this method to return the
// default set of toolbar items, including any extra toolbar items you want
// included. For example, include [flexibleSpace] strings as the first and
// last elements of the array to center the remaining toolbar items. If you
// add custom identifiers, you must also override the
// [ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar] method to specify
// the content for those toolbar items.
// 
// If you override this method, you must call `super` at some point in your
// implementation. The default implementation of this method returns the
// identifiers for all toolbar items that correspond to tabs in the tab bar
// interface.
//
// [NSTabViewController.TabStyle.toolbar]: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum/toolbar
// [flexibleSpace]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier/flexibleSpace
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/toolbarDefaultItemIdentifiers(_:)
func (t NSTabViewController) ToolbarDefaultItemIdentifiers(toolbar INSToolbar) []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("toolbarDefaultItemIdentifiers:"), toolbar)
	return objc.ConvertSliceToStrings(rv)
}

// Returns the array of identifier strings for the selectable toolbar items
//
// toolbar: The toolbar making the request.
//
// # Return Value
// 
// An array of [NSString] objects, each of which contains an identifier for a
// toolbar item that may be selected.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// # Discussion
// 
// This method is called for tab view interfaces that use the
// [NSTabViewController.TabStyle.toolbar] style. Use this method to indicate
// which toolbar items are selectable. When an item is selected, the toolbar
// displays it with a visual highlight and updates the
// [SelectedTabViewItemIndex] property. Typically, the toolbar items
// associated with tabs are selectable so that the user can tell which tab is
// selected.
// 
// If you override this method, you must call `super` at some point in your
// implementation. The default implementation of this method returns the
// identifiers for all toolbar items that correspond to tabs in the tab bar
// interface.
//
// [NSTabViewController.TabStyle.toolbar]: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum/toolbar
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/toolbarSelectableItemIdentifiers(_:)
func (t NSTabViewController) ToolbarSelectableItemIdentifiers(toolbar INSToolbar) []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("toolbarSelectableItemIdentifiers:"), toolbar)
	return objc.ConvertSliceToStrings(rv)
}

// Informs the delegate that the number of tab view items in `tabView` has
// changed.
//
// tabView: The tab view that added or removed tabview items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewDelegate/tabViewDidChangeNumberOfTabViewItems(_:)
func (t NSTabViewController) TabViewDidChangeNumberOfTabViewItems(tabView INSTabView) {
	objc.Send[objc.ID](t.ID, objc.Sel("tabViewDidChangeNumberOfTabViewItems:"), tabView)
}

// Tells the delegate that the toolbar removed the specified item.
//
// notification: A notification named [didRemoveItemNotification].
// //
// [didRemoveItemNotification]: https://developer.apple.com/documentation/AppKit/NSToolbar/didRemoveItemNotification
//
// # Discussion
// 
// Use this method to update data structures related to your toolbar items.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarDelegate/toolbarDidRemoveItem(_:)
func (t NSTabViewController) ToolbarDidRemoveItem(notification foundation.NSNotification) {
	objc.Send[objc.ID](t.ID, objc.Sel("toolbarDidRemoveItem:"), notification)
}

// Asks the delegate for a Boolean value that indicates whether the toolbar
// can place the item at the specified position.
//
// toolbar: The toolbar that contains the items.
//
// itemIdentifier: The identifier of the toolbar item to insert at the specified index.
//
// index: The proposed index at which to place the item. If the toolbar is removing
// the item, this value is [NSNotFound].
// //
// [NSNotFound]: https://developer.apple.com/documentation/Foundation/NSNotFound-4qp9h
//
// # Return Value
// 
// [true] to allow the toolbar to place the item at the specified location, or
// [false] to prevent the toolbar from placing the item in that location.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Implement this method to control the placement of items in the toolbar.
// During a drag operation, the toolbar calls this method to determine if the
// specified index is an acceptable location for the item. Return a Boolean
// value that indicates whether the new posiition is acceptable.
// 
// Don’t use the `index` parameter to determine the final location of the
// toolbar item. During a drag operation, the toolbar can call this method
// multiple times, so the index value can change later.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarDelegate/toolbar(_:itemIdentifier:canBeInsertedAt:)
func (t NSTabViewController) ToolbarItemIdentifierCanBeInsertedAtIndex(toolbar INSToolbar, itemIdentifier NSToolbarItemIdentifier, index int) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("toolbar:itemIdentifier:canBeInsertedAtIndex:"), toolbar, objc.String(string(itemIdentifier)), index)
	return rv
}

// Tells the delegate that the toolbar is about to add the specified item.
//
// notification: A notification named [willAddItemNotification].
// //
// [willAddItemNotification]: https://developer.apple.com/documentation/AppKit/NSToolbar/willAddItemNotification
//
// # Discussion
// 
// Use this method to cache references to new toolbar items or perform any
// tasks related to the addition of those items.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarDelegate/toolbarWillAddItem(_:)
func (t NSTabViewController) ToolbarWillAddItem(notification foundation.NSNotification) {
	objc.Send[objc.ID](t.ID, objc.Sel("toolbarWillAddItem:"), notification)
}

// Asks the delegate to provide the items that people can’t remove from the
// toolbar or rearrange during the customization process.
//
// toolbar: The toolbar that contains the items.
//
// # Return Value
// 
// The set of item identifiers that people can’t remove from the toolbar or
// move to other locations in the toolbar. Return an empty set to let someone
// customize all toolbar items.
//
// # Discussion
// 
// Implement this method in your delegate and return any items you don’t
// want people to remove or rearrange. If you don’t implement this method,
// the toolbar lets people rearrange and remove all toolbar items.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarDelegate/toolbarImmovableItemIdentifiers(_:)
func (t NSTabViewController) ToolbarImmovableItemIdentifiers(toolbar INSToolbar) foundation.INSSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("toolbarImmovableItemIdentifiers:"), toolbar)
	return foundation.NSSetFromID(rv)
}
func (t NSTabViewController) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The style used to display the tabs.
//
// # Discussion
// 
// The default value of this property is
// [NSTabViewController.TabStyle.segmentedControlOnTop]. Changing the style at
// runtime updates the appearance of the tab view controller interface.
//
// [NSTabViewController.TabStyle.segmentedControlOnTop]: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum/segmentedControlOnTop
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabStyle-swift.property
func (t NSTabViewController) TabStyle() NSTabViewControllerTabStyle {
	rv := objc.Send[NSTabViewControllerTabStyle](t.ID, objc.Sel("tabStyle"))
	return NSTabViewControllerTabStyle(rv)
}
func (t NSTabViewController) SetTabStyle(value NSTabViewControllerTabStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setTabStyle:"), value)
}



// The tab view that manages the views of the interface.
//
// # Discussion
// 
// Use this property to access the tab view controller’s content view. The
// object in this property may not be the same as the one in the tab view
// controller’s [View] property. The tab view controller works directly with
// the [NSTabView] object, setting itself as the tab view’s delegate. You
// must not modify the items of the tab view directly or change its delegate.
// Instead, use the methods of this class to make your changes.
// 
// Accessing this property creates the tab view object if it does not already
// exist. To determine whether the tab view has been created (without creating
// it prematurely), use the [ViewLoaded] property.
// 
// You may provide your own tab view by assigning it to this property. If you
// do so, you must assign your custom object before the tab view controller
// creates one of its own. In other words, you must assign your tab view
// object to this property while the [ViewLoaded] property is still [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabView
func (t NSTabViewController) TabView() INSTabView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tabView"))
	return NSTabViewFromID(objc.ID(rv))
}
func (t NSTabViewController) SetTabView(value INSTabView) {
	objc.Send[struct{}](t.ID, objc.Sel("setTabView:"), value)
}



// The animation options to use when switching between tabs.
//
// # Discussion
// 
// By default, this property is set to the [ViewControllerTransitionCrossfade]
// and [ViewControllerTransitionAllowUserInteraction] options.
// 
// The tab view controller uses the
// [TransitionFromViewControllerToViewControllerOptionsCompletionHandler]
// method to perform transitions between tabs. For more information about how
// transitions happen, see the description of that method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/transitionOptions
func (t NSTabViewController) TransitionOptions() NSViewControllerTransitionOptions {
	rv := objc.Send[NSViewControllerTransitionOptions](t.ID, objc.Sel("transitionOptions"))
	return NSViewControllerTransitionOptions(rv)
}
func (t NSTabViewController) SetTransitionOptions(value NSViewControllerTransitionOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setTransitionOptions:"), value)
}



// A Boolean value indicating whether the tab view controller gets its title
// from the selected child view controller.
//
// # Discussion
// 
// When this property is [true] and the tab view controller’s own title is
// `nil`, the tab view controller gets its title from the [Title] property of
// the selected child view controller. When this property is [false], the tab
// view controller always provides the title, which may be `nil`. The default
// value of this property is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/canPropagateSelectedChildViewControllerTitle
func (t NSTabViewController) CanPropagateSelectedChildViewControllerTitle() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("canPropagateSelectedChildViewControllerTitle"))
	return rv
}
func (t NSTabViewController) SetCanPropagateSelectedChildViewControllerTitle(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setCanPropagateSelectedChildViewControllerTitle:"), value)
}



// The array of tab view items used to manage each of the child view
// controllers.
//
// # Discussion
// 
// This property contains an array of [NSTabViewItem] objects. Each tab view
// item contains information about a tab in the tab view interface, including
// the child view controller that manages the tab’s contents.
// 
// Assigning a new array to this property updates the set of tabs displayed by
// the tab view controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/tabViewItems
func (t NSTabViewController) TabViewItems() []NSTabViewItem {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("tabViewItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTabViewItem {
		return NSTabViewItemFromID(id)
	})
}
func (t NSTabViewController) SetTabViewItems(value []NSTabViewItem) {
	objc.Send[struct{}](t.ID, objc.Sel("setTabViewItems:"), objectivec.IObjectSliceToNSArray(value))
}



// The index of the selected tab.
//
// # Discussion
// 
// Use this property to get and set the selected tab. The property is
// key-value coding compliant and can be the target of bindings.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewController/selectedTabViewItemIndex
func (t NSTabViewController) SelectedTabViewItemIndex() int {
	rv := objc.Send[int](t.ID, objc.Sel("selectedTabViewItemIndex"))
	return rv
}
func (t NSTabViewController) SetSelectedTabViewItemIndex(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectedTabViewItemIndex:"), value)
}



// An array of view controllers that are hierarchical children of the view
// controller.
//
// See: https://developer.apple.com/documentation/appkit/nsviewcontroller/children
func (t NSTabViewController) Children() INSViewController {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("childViewControllers"))
	return NSViewControllerFromID(objc.ID(rv))
}
func (t NSTabViewController) SetChildren(value INSViewController) {
	objc.Send[struct{}](t.ID, objc.Sel("setChildViewControllers:"), value)
}



















			// Protocol methods for NSTabViewDelegate
			




			// Protocol methods for NSToolbarDelegate
			


















