// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// The optional methods implemented by delegates of [NSMenu](<doc://com.apple.appkit/documentation/AppKit/NSMenu>) objects to manage menu display and handle some events.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuDelegate
type NSMenuDelegate interface {
	objectivec.IObject
}



// NSMenuDelegateObject wraps an existing Objective-C object that conforms to the NSMenuDelegate protocol.
type NSMenuDelegateObject struct {
	objectivec.Object
}
func (o NSMenuDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSMenuDelegateObjectFromID constructs a [NSMenuDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSMenuDelegateObjectFromID(id objc.ID) NSMenuDelegateObject {
	return NSMenuDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Invoked to allow the delegate to return the target and action for a
// key-down event.
//
// menu: The menu object sending the delegation message.
//
// event: An [NSEvent] object representing a key-down event.
//
// target: Return by reference the target object for the menu item that corresponds to
// the event. Specify `nil` to request the menu’s target.
//
// action: Return by reference the action selector for the menu item that corresponds
// to the event.
//
// # Return Value
// 
// If there is a valid and enabled menu item that corresponds to this key-down
// even, return [true] after specifying the target and action. Return [false]
// if there are no items with that key equivalent or if the item is disabled.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the delegate doesn’t define this method, the menu is populated to find
// out if any items have a matching key equivalent.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuDelegate/menuHasKeyEquivalent(_:for:target:action:)

func (o NSMenuDelegateObject) MenuHasKeyEquivalentForEventTargetAction(menu INSMenu, event INSEvent, target []objectivec.IObject, action objc.SEL) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("menuHasKeyEquivalent:forEvent:target:action:"), menu, event, objectivec.IObjectSliceToNSArray(target), action)
	return rv
	}

// Invoked to let the delegate update a menu item before it is displayed.
//
// menu: The menu object that owns `item`.
//
// item: The menu-item object that may be updated.
//
// index: The integer index of the menu item.
//
// shouldCancel: Set to [true] if, due to some user action, the menu no longer needs to be
// displayed before all the menu items have been updated. You can ignore this
// flag, return [true], and continue; or you can save your work (to save time
// the next time your delegate is called) and return [false] to stop the
// updating.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// [true] to continue the process. If you return [false], your
// [MenuUpdateItemAtIndexShouldCancel] is not called again. In that case,
// it’s your responsibility to trim any extra items from the menu.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If your [NumberOfItemsInMenu] delegate method returns a positive value,
// then your [MenuUpdateItemAtIndexShouldCancel] method is called for each
// item in the menu. You can then update the menu title, image, and so forth
// for each menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuDelegate/menu(_:update:at:shouldCancel:)

func (o NSMenuDelegateObject) MenuUpdateItemAtIndexShouldCancel(menu INSMenu, item INSMenuItem, index int, shouldCancel bool) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("menu:updateItem:atIndex:shouldCancel:"), menu, item, index, shouldCancel)
	return rv
	}

// Invoked to allow the delegate to specify a display location for the menu.
//
// menu: The menu object.
//
// screen: The screen the menu will open on.
//
// # Return Value
// 
// The rectangle the menu should be displayed within, in screen coordinates.
//
// # Discussion
// 
// This method is sent to the delegate when a menu is about to be opened on
// the specified screen.
// 
// If you return [NSZeroRect], or if the delegate doesn’t implement this
// method, the menu will be confined to the bounds appropriate for the given
// screen. The returned rect may not be honored in all cases, for example, if
// it would force the menu to be too small.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuDelegate/confinementRect(for:on:)

func (o NSMenuDelegateObject) ConfinementRectForMenuOnScreen(menu INSMenu, screen INSScreen) corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("confinementRectForMenu:onScreen:"), menu, screen)
	return rv
	}

// Invoked to indicate that a menu is about to highlight a given item.
//
// menu: The menu object about to highlight an item.
//
// item: The item about to be highlighted.
//
// # Discussion
// 
// Only one item per menu can be highlighted at a time. If `item` is `nil`, it
// means that all items in the menu are about to be unhighlighted.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuDelegate/menu(_:willHighlight:)

func (o NSMenuDelegateObject) MenuWillHighlightItem(menu INSMenu, item INSMenuItem) {
	
	objc.Send[struct{}](o.ID, objc.Sel("menu:willHighlightItem:"), menu, item)
	}

// Invoked when a menu is about to open.
//
// menu: The menu that is about to open.
//
// # Discussion
// 
// Don’t modify the structure of the menu or the menu items during this
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuDelegate/menuWillOpen(_:)

func (o NSMenuDelegateObject) MenuWillOpen(menu INSMenu) {
	
	objc.Send[struct{}](o.ID, objc.Sel("menuWillOpen:"), menu)
	}

// Invoked after a menu closed.
//
// menu: The menu that closed.
//
// # Discussion
// 
// Don’t modify the structure of the menu or the menu items during this
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuDelegate/menuDidClose(_:)

func (o NSMenuDelegateObject) MenuDidClose(menu INSMenu) {
	
	objc.Send[struct{}](o.ID, objc.Sel("menuDidClose:"), menu)
	}

// Invoked when a menu is about to be displayed at the start of a tracking
// session so the delegate can specify the number of items in the menu.
//
// menu: The menu object about to be displayed.
//
// # Return Value
// 
// The number of menu items in the menu.
//
// # Discussion
// 
// If you return a positive value, the menu is resized by either removing or
// adding items. Newly created items are blank. After the menu is resized,
// your [MenuUpdateItemAtIndexShouldCancel] method is called for each item. If
// you return a negative value, the number of items is left unchanged and
// [MenuUpdateItemAtIndexShouldCancel] is not called. If you can populate the
// menu quickly, you can implement [MenuNeedsUpdate] instead of
// [NumberOfItemsInMenu] and [MenuUpdateItemAtIndexShouldCancel].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuDelegate/numberOfItems(in:)

func (o NSMenuDelegateObject) NumberOfItemsInMenu(menu INSMenu) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("numberOfItemsInMenu:"), menu)
	return rv
	}

// Invoked when a menu is about to be displayed at the start of a tracking
// session.
//
// menu: The menu object that is about to be displayed.
//
// # Discussion
// 
// Using this method, the delegate can change the menu by adding, removing, or
// modifying menu items. If populating the menu will take a long time,
// implement [NumberOfItemsInMenu] and [MenuUpdateItemAtIndexShouldCancel]
// instead.
// 
// Menu item validation occurs after this method is called. If the menu is
// updated because the user pressed a command key, only the menu item with the
// matching command key is validated; if the menu is updated because the user
// opened it, then every menu item is validated.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuDelegate/menuNeedsUpdate(_:)

func (o NSMenuDelegateObject) MenuNeedsUpdate(menu INSMenu) {
	
	objc.Send[struct{}](o.ID, objc.Sel("menuNeedsUpdate:"), menu)
	}





// NSMenuDelegateConfig holds optional typed callbacks for [NSMenuDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsmenudelegate
type NSMenuDelegateConfig struct {

	// Handling Open and Close Events
	// WillOpen — Invoked when a menu is about to open.
	WillOpen func(menu NSMenu)
	// DidClose — Invoked after a menu closed.
	DidClose func(menu NSMenu)

	// Handling Tracking
	// NeedsUpdate — Invoked when a menu is about to be displayed at the start of a tracking session.
	NeedsUpdate func(menu NSMenu)

	// Other Methods
	// UpdateItemAtIndexShouldCancel — Invoked to let the delegate update a menu item before it is displayed.
	UpdateItemAtIndexShouldCancel func(menu NSMenu, item NSMenuItem, index int, shouldCancel bool) bool
	// WillHighlightItem — Invoked to indicate that a menu is about to highlight a given item.
	WillHighlightItem func(menu NSMenu, item NSMenuItem)
	// NumberOfItemsInMenu — Invoked when a menu is about to be displayed at the start of a tracking session so the delegate can specify the number of items in the menu.
	NumberOfItemsInMenu func(menu NSMenu) int
}

// NewNSMenuDelegate creates an Objective-C object implementing the [NSMenuDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSMenuDelegateObject] satisfies the [NSMenuDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsmenudelegate
func NewNSMenuDelegate(config NSMenuDelegateConfig) NSMenuDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSMenuDelegate_%d", n)

	var methods []objc.MethodDef

	if config.UpdateItemAtIndexShouldCancel != nil {
		fn := config.UpdateItemAtIndexShouldCancel
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("menu:updateItem:atIndex:shouldCancel:"),
			Fn: func(self objc.ID, _cmd objc.SEL, menuID objc.ID, itemID objc.ID, index int, shouldCancel bool) bool {
				menu := NSMenuFromID(menuID)
				item := NSMenuItemFromID(itemID)
				return fn(menu, item, index, shouldCancel)
			},
		})
	}

	if config.WillHighlightItem != nil {
		fn := config.WillHighlightItem
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("menu:willHighlightItem:"),
			Fn: func(self objc.ID, _cmd objc.SEL, menuID objc.ID, itemID objc.ID) {
				menu := NSMenuFromID(menuID)
				item := NSMenuItemFromID(itemID)
				fn(menu, item)
			},
		})
	}

	if config.WillOpen != nil {
		fn := config.WillOpen
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("menuWillOpen:"),
			Fn: func(self objc.ID, _cmd objc.SEL, menuID objc.ID) {
				menu := NSMenuFromID(menuID)
				fn(menu)
			},
		})
	}

	if config.DidClose != nil {
		fn := config.DidClose
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("menuDidClose:"),
			Fn: func(self objc.ID, _cmd objc.SEL, menuID objc.ID) {
				menu := NSMenuFromID(menuID)
				fn(menu)
			},
		})
	}

	if config.NumberOfItemsInMenu != nil {
		fn := config.NumberOfItemsInMenu
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("numberOfItemsInMenu:"),
			Fn: func(self objc.ID, _cmd objc.SEL, menuID objc.ID) int {
				menu := NSMenuFromID(menuID)
				return fn(menu)
			},
		})
	}

	if config.NeedsUpdate != nil {
		fn := config.NeedsUpdate
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("menuNeedsUpdate:"),
			Fn: func(self objc.ID, _cmd objc.SEL, menuID objc.ID) {
				menu := NSMenuFromID(menuID)
				fn(menu)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSMenuDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSMenuDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSMenuDelegateObjectFromID(instance)
}





