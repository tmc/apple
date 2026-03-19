// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A set of optional methods you use to configure the toolbar and respond to changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarDelegate
type NSToolbarDelegate interface {
	objectivec.IObject
}

// NSToolbarDelegateObject wraps an existing Objective-C object that conforms to the NSToolbarDelegate protocol.
type NSToolbarDelegateObject struct {
	objectivec.Object
}
func (o NSToolbarDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSToolbarDelegateObjectFromID constructs a [NSToolbarDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSToolbarDelegateObjectFromID(id objc.ID) NSToolbarDelegateObject {
	return NSToolbarDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate for the toolbar item associated with the specified
// identifier.
//
// toolbar: The toolbar for which the item is being requested.
//
// itemIdentifier: The identifier for the requested item.
//
// flag: [true] if the toolbar will insert the item immediately. If this parameter
// is [false], provide a canonical representation for the item. For example,
// provide a version of the item suitable for display in the toolbar
// customization sheet.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A new [NSToolbarItem] object, or `nil` if no toolbar item is available for
// the specified identifier.
//
// # Discussion
// 
// Use this method to create new [NSToolbarItem] objects when the toolbar asks
// for them. If your toolbar item uses a custom view, make sure that view is
// fully configured before you return the item. The toolbar becomes the owner
// of the returned item, but can display the item either in the toolbar or the
// customization palette.
// 
// Don’t recycle toolbar items; always provide a new instance, even if the
// toolbar previously asked for an item with the same identifier.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarDelegate/toolbar(_:itemForItemIdentifier:willBeInsertedIntoToolbar:)
func (o NSToolbarDelegateObject) ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(toolbar INSToolbar, itemIdentifier NSToolbarItemIdentifier, flag bool) INSToolbarItem {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"), toolbar, objc.String(string(itemIdentifier)), flag)
	return NSToolbarItemFromID(rv)
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
func (o NSToolbarDelegateObject) ToolbarWillAddItem(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("toolbarWillAddItem:"), notification)
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
func (o NSToolbarDelegateObject) ToolbarDidRemoveItem(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("toolbarDidRemoveItem:"), notification)
	}
// Asks the delegate to provide the items allowed on the toolbar.
//
// toolbar: The toolbar whose allowed item identifiers are to be returned.
//
// # Return Value
// 
// An array of toolbar item identifiers, each of which represents an item that
// appears in the customization palette. Arrange the identifiers in the order
// you want them to appear in the palette, with the first item appearing on
// the palette’s leading edge.
//
// # Discussion
// 
// Include all of your toolbar’s items, including standard ones defined by
// [NSToolbarIdentifier]. The array must include all of the default menu items
// in your toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarDelegate/toolbarAllowedItemIdentifiers(_:)
func (o NSToolbarDelegateObject) ToolbarAllowedItemIdentifiers(toolbar INSToolbar) []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("toolbarAllowedItemIdentifiers:"), toolbar)
	return objc.ConvertSliceToStrings(rv)
	}
// Asks the delegate to provide the default items to display on the toolbar.
//
// toolbar: The toolbar whose default item identifiers are to be returned.
//
// # Return Value
// 
// An array of toolbar item identifiers, each of which represents an item that
// appears in the default toolbar. Arrange the identifiers in the order you
// want them to appear in the toolbar, with the first item appearing on the
// toolbar’s leading edge.
//
// # Discussion
// 
// The toolbar calls this method when user settings don’t contain any custom
// configuration data for the toolbar. The toolbar also calls it to initialize
// the customization palette’s contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarDelegate/toolbarDefaultItemIdentifiers(_:)
func (o NSToolbarDelegateObject) ToolbarDefaultItemIdentifiers(toolbar INSToolbar) []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("toolbarDefaultItemIdentifiers:"), toolbar)
	return objc.ConvertSliceToStrings(rv)
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
func (o NSToolbarDelegateObject) ToolbarImmovableItemIdentifiers(toolbar INSToolbar) foundation.INSSet {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("toolbarImmovableItemIdentifiers:"), toolbar)
	return foundation.NSSetFromID(rv)
	}
// Asks the delegate to provide the set of selectable items in the toolbar.
//
// toolbar: The toolbar that contains the items.
//
// # Return Value
// 
// An array of item identifiers, each of which corresponds to an
// [NSToolbarItem] that should display a selection indicator in the specified
// toolbar.
//
// # Discussion
// 
// Use this method to return the complete list of toolbar items that support
// selection. When someone selects one of the returned items, the toolbar
// automatically displays that item with a visual highlight. The toolbar also
// places the currently selected item in its [SelectedItemIdentifier]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarDelegate/toolbarSelectableItemIdentifiers(_:)
func (o NSToolbarDelegateObject) ToolbarSelectableItemIdentifiers(toolbar INSToolbar) []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("toolbarSelectableItemIdentifiers:"), toolbar)
	return objc.ConvertSliceToStrings(rv)
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
func (o NSToolbarDelegateObject) ToolbarItemIdentifierCanBeInsertedAtIndex(toolbar INSToolbar, itemIdentifier NSToolbarItemIdentifier, index int) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("toolbar:itemIdentifier:canBeInsertedAtIndex:"), toolbar, objc.String(string(itemIdentifier)), index)
	return rv
	}

// NSToolbarDelegateConfig holds optional typed callbacks for [NSToolbarDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstoolbardelegate
type NSToolbarDelegateConfig struct {

	// Adding and removing items
	// ItemForItemIdentifierWillBeInsertedIntoToolbar — Asks the delegate for the toolbar item associated with the specified identifier.
	ItemForItemIdentifierWillBeInsertedIntoToolbar func(toolbar NSToolbar, itemIdentifier NSToolbarItemIdentifier, flag bool) NSToolbarItem
	// WillAddItem — Tells the delegate that the toolbar is about to add the specified item.
	WillAddItem func(notification foundation.NSNotification)
	// DidRemoveItem — Tells the delegate that the toolbar removed the specified item.
	DidRemoveItem func(notification foundation.NSNotification)

	// Configuring the behavior of items
	// ImmovableItemIdentifiers — Asks the delegate to provide the items that people can’t remove from the toolbar or rearrange during the customization process.
	ImmovableItemIdentifiers func(toolbar NSToolbar) foundation.INSSet

	// Other Methods
	// ItemIdentifierCanBeInsertedAtIndex — Asks the delegate for a Boolean value that indicates whether the toolbar can place the item at the specified position.
	ItemIdentifierCanBeInsertedAtIndex func(toolbar NSToolbar, itemIdentifier NSToolbarItemIdentifier, index int) bool
}

// NewNSToolbarDelegate creates an Objective-C object implementing the [NSToolbarDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSToolbarDelegateObject] satisfies the [NSToolbarDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstoolbardelegate
func NewNSToolbarDelegate(config NSToolbarDelegateConfig) NSToolbarDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSToolbarDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ItemForItemIdentifierWillBeInsertedIntoToolbar != nil {
		fn := config.ItemForItemIdentifierWillBeInsertedIntoToolbar
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"),
			Fn: func(self objc.ID, _cmd objc.SEL, toolbarID objc.ID, itemIdentifierID objc.ID, flag bool) objc.ID {
				toolbar := NSToolbarFromID(toolbarID)
				itemIdentifier := NSToolbarItemIdentifier(objc.GoString(objc.Send[*byte](itemIdentifierID, objc.Sel("UTF8String"))))
				return fn(toolbar, itemIdentifier, flag).GetID()
			},
		})
	}

	if config.WillAddItem != nil {
		fn := config.WillAddItem
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("toolbarWillAddItem:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidRemoveItem != nil {
		fn := config.DidRemoveItem
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("toolbarDidRemoveItem:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ImmovableItemIdentifiers != nil {
		fn := config.ImmovableItemIdentifiers
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("toolbarImmovableItemIdentifiers:"),
			Fn: func(self objc.ID, _cmd objc.SEL, toolbarID objc.ID) objc.ID {
				toolbar := NSToolbarFromID(toolbarID)
				return fn(toolbar).GetID()
			},
		})
	}

	if config.ItemIdentifierCanBeInsertedAtIndex != nil {
		fn := config.ItemIdentifierCanBeInsertedAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("toolbar:itemIdentifier:canBeInsertedAtIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, toolbarID objc.ID, itemIdentifierID objc.ID, index int) bool {
				toolbar := NSToolbarFromID(toolbarID)
				itemIdentifier := NSToolbarItemIdentifier(objc.GoString(objc.Send[*byte](itemIdentifierID, objc.Sel("UTF8String"))))
				return fn(toolbar, itemIdentifier, index)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSToolbarDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSToolbarDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSToolbarDelegateObjectFromID(instance)
}

