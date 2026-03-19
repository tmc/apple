// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPageLayout] class.
var (
	_NSPageLayoutClass     NSPageLayoutClass
	_NSPageLayoutClassOnce sync.Once
)

func getNSPageLayoutClass() NSPageLayoutClass {
	_NSPageLayoutClassOnce.Do(func() {
		_NSPageLayoutClass = NSPageLayoutClass{class: objc.GetClass("NSPageLayout")}
	})
	return _NSPageLayoutClass
}

// GetNSPageLayoutClass returns the class object for NSPageLayout.
func GetNSPageLayoutClass() NSPageLayoutClass {
	return getNSPageLayoutClass()
}

type NSPageLayoutClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPageLayoutClass) Alloc() NSPageLayout {
	rv := objc.Send[NSPageLayout](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A panel that queries the user for information such as paper type and
// orientation.
//
// # Overview
// 
// A page layout panel is typically displayed in response to the user
// selecting the Page Setup menu item. You obtain an instance with the
// [pageLayout] class method. The pane can then be run as a sheet using
// [NSPageLayout.BeginSheetWithPrintInfoModalForWindowDelegateDidEndSelectorContextInfo] or
// modally using [NSPageLayout.RunModal] or [NSPageLayout.RunModalWithPrintInfo].
// 
// For design guidance, see [Human Interface Guidelines].
//
// [Human Interface Guidelines]: https://developer.apple.com/design/human-interface-guidelines/macos/system-capabilities/printing#page-setup-dialogs
// [pageLayout]: https://developer.apple.com/documentation/AppKit/NSPageLayout/pageLayout
//
// # Running the page setup dialog
//
//   - [NSPageLayout.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
//   - [NSPageLayout.RunModal]: Displays the page layout panel and begins the modal loop using the shared print info object.
//   - [NSPageLayout.RunModalWithPrintInfo]: Displays the page layout panel and begins the modal loop using the specified print info object.
//
// # Customizing the page setup dialog
//
//   - [NSPageLayout.AddAccessoryController]: Adds the specified controller of an accessory view to be presented in the page setup panel.
//   - [NSPageLayout.RemoveAccessoryController]: Removes the specified controller of an accessory view.
//   - [NSPageLayout.AccessoryControllers]: An array of accessory view controllers belonging to the page layout panel.
//
// # Accessing the printing information
//
//   - [NSPageLayout.PrintInfo]: The printing information object used when the page layout panel is run.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageLayout
type NSPageLayout struct {
	objectivec.Object
}

// NSPageLayoutFromID constructs a [NSPageLayout] from an objc.ID.
//
// A panel that queries the user for information such as paper type and
// orientation.
func NSPageLayoutFromID(id objc.ID) NSPageLayout {
	return NSPageLayout{objectivec.Object{ID: id}}
}
// NOTE: NSPageLayout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPageLayout] class.
//
// # Running the page setup dialog
//
//   - [INSPageLayout.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
//   - [INSPageLayout.RunModal]: Displays the page layout panel and begins the modal loop using the shared print info object.
//   - [INSPageLayout.RunModalWithPrintInfo]: Displays the page layout panel and begins the modal loop using the specified print info object.
//
// # Customizing the page setup dialog
//
//   - [INSPageLayout.AddAccessoryController]: Adds the specified controller of an accessory view to be presented in the page setup panel.
//   - [INSPageLayout.RemoveAccessoryController]: Removes the specified controller of an accessory view.
//   - [INSPageLayout.AccessoryControllers]: An array of accessory view controllers belonging to the page layout panel.
//
// # Accessing the printing information
//
//   - [INSPageLayout.PrintInfo]: The printing information object used when the page layout panel is run.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageLayout
type INSPageLayout interface {
	objectivec.IObject

	// Topic: Running the page setup dialog

	BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo INSPrintInfo, parentWindow INSWindow, handler PageLayoutResultHandler)
	// Displays the page layout panel and begins the modal loop using the shared print info object.
	RunModal() int
	// Displays the page layout panel and begins the modal loop using the specified print info object.
	RunModalWithPrintInfo(printInfo INSPrintInfo) int

	// Topic: Customizing the page setup dialog

	// Adds the specified controller of an accessory view to be presented in the page setup panel.
	AddAccessoryController(accessoryController INSViewController)
	// Removes the specified controller of an accessory view.
	RemoveAccessoryController(accessoryController INSViewController)
	// An array of accessory view controllers belonging to the page layout panel.
	AccessoryControllers() []NSViewController

	// Topic: Accessing the printing information

	// The printing information object used when the page layout panel is run.
	PrintInfo() INSPrintInfo
}

// Init initializes the instance.
func (p NSPageLayout) Init() NSPageLayout {
	rv := objc.Send[NSPageLayout](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPageLayout) Autorelease() NSPageLayout {
	rv := objc.Send[NSPageLayout](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPageLayout creates a new NSPageLayout instance.
func NewNSPageLayout() NSPageLayout {
	class := getNSPageLayoutClass()
	rv := objc.Send[NSPageLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSPageLayout/beginSheet(using:on:completionHandler:)
func (p NSPageLayout) BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo INSPrintInfo, parentWindow INSWindow, handler PageLayoutResultHandler) {
_block2, _cleanup2 := NewPageLayoutResultBlock(handler)
	defer _cleanup2()
	objc.Send[objc.ID](p.ID, objc.Sel("beginSheetUsingPrintInfo:onWindow:completionHandler:"), printInfo, parentWindow, _block2)
}
// Displays the page layout panel and begins the modal loop using the shared
// print info object.
//
// # Return Value
// 
// [NSCancelButton] if the user clicks the Cancel button; otherwise,
// [NSOKButton].
//
// # Discussion
// 
// The receiver’s values are recorded in the shared [NSPrintInfo] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageLayout/runModal()
func (p NSPageLayout) RunModal() int {
	rv := objc.Send[int](p.ID, objc.Sel("runModal"))
	return rv
}
// Displays the page layout panel and begins the modal loop using the
// specified print info object.
//
// printInfo: The [NSPrintInfo] object to use.
//
// # Return Value
// 
// [NSCancelButton] if the user clicks the Cancel button; otherwise,
// [NSOKButton].
//
// # Discussion
// 
// The receiver’s values are recorded in `printInfo`.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageLayout/runModal(with:)
func (p NSPageLayout) RunModalWithPrintInfo(printInfo INSPrintInfo) int {
	rv := objc.Send[int](p.ID, objc.Sel("runModalWithPrintInfo:"), printInfo)
	return rv
}
// Adds the specified controller of an accessory view to be presented in the
// page setup panel.
//
// accessoryController: The controller to add.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageLayout/addAccessoryController(_:)
func (p NSPageLayout) AddAccessoryController(accessoryController INSViewController) {
	objc.Send[objc.ID](p.ID, objc.Sel("addAccessoryController:"), accessoryController)
}
// Removes the specified controller of an accessory view.
//
// accessoryController: The controller to remove.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageLayout/removeAccessoryController(_:)
func (p NSPageLayout) RemoveAccessoryController(accessoryController INSViewController) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeAccessoryController:"), accessoryController)
}

// An array of accessory view controllers belonging to the page layout panel.
//
// # Discussion
// 
// The [NSViewController] instances representing the accessory view
// controllers belonging to the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageLayout/accessoryControllers
func (p NSPageLayout) AccessoryControllers() []NSViewController {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("accessoryControllers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSViewController {
		return NSViewControllerFromID(id)
	})
}
// The printing information object used when the page layout panel is run.
//
// # Discussion
// 
// The NSPrintInfo object is set using the
// [BeginSheetWithPrintInfoModalForWindowDelegateDidEndSelectorContextInfo] or
// [RunModalWithPrintInfo] method. The shared NSPrintInfo object is used if
// the receiver is run using [RunModal].
//
// See: https://developer.apple.com/documentation/AppKit/NSPageLayout/printInfo
func (p NSPageLayout) PrintInfo() INSPrintInfo {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("printInfo"))
	return NSPrintInfoFromID(objc.ID(rv))
}

// BeginSheetUsingPrintInfoOnWindow is a synchronous wrapper around [NSPageLayout.BeginSheetUsingPrintInfoOnWindowCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p NSPageLayout) BeginSheetUsingPrintInfoOnWindow(ctx context.Context, printInfo INSPrintInfo, parentWindow INSWindow) (NSPageLayoutResult, error) {
	done := make(chan NSPageLayoutResult, 1)
	p.BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo, parentWindow, func(val NSPageLayoutResult) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

