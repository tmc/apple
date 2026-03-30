// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPrintPanel] class.
var (
	_NSPrintPanelClass     NSPrintPanelClass
	_NSPrintPanelClassOnce sync.Once
)

func getNSPrintPanelClass() NSPrintPanelClass {
	_NSPrintPanelClassOnce.Do(func() {
		_NSPrintPanelClass = NSPrintPanelClass{class: objc.GetClass("NSPrintPanel")}
	})
	return _NSPrintPanelClass
}

// GetNSPrintPanelClass returns the class object for NSPrintPanel.
func GetNSPrintPanelClass() NSPrintPanelClass {
	return getNSPrintPanelClass()
}

type NSPrintPanelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPrintPanelClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPrintPanelClass) Alloc() NSPrintPanel {
	rv := objc.Send[NSPrintPanel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The Print panel that queries the user for information about a print job.
//
// # Overview
//
// A Print panel may let the user select the range of pages to print and the
// number of copies before executing the Print command. Print panels can
// display a simplified interface when printing certain types of data. For
// example, the panel can display a list of print-setting presets, which lets
// the user enable print settings in groups as opposed to individually.
// Assigning an appropriate string to the [NSPrintPanel.JobStyleHint] property activates
// the simplified interface and identifies which presets to display.
//
// For design guidance, see [Human Interface Guidelines].
//
// # Customizing the Panel
//
//   - [NSPrintPanel.JobStyleHint]: The type of settings that the print panel displays.
//   - [NSPrintPanel.SetJobStyleHint]
//   - [NSPrintPanel.Options]: The current configuration options for the Print panel.
//   - [NSPrintPanel.SetOptions]
//   - [NSPrintPanel.DefaultButtonTitle]: Returns the title of the Print panel’s default button.
//   - [NSPrintPanel.SetDefaultButtonTitle]: Sets the title of the Print panel’s default button.
//   - [NSPrintPanel.HelpAnchor]: The HTML help anchor associated with the Print panel.
//   - [NSPrintPanel.SetHelpAnchor]
//
// # Managing Accessory Views
//
//   - [NSPrintPanel.AddAccessoryController]: Adds a custom controller to the Print panel to manage an accessory view.
//   - [NSPrintPanel.RemoveAccessoryController]: Removes the specified controller and accessory view from the Print panel.
//   - [NSPrintPanel.AccessoryControllers]: The array of controller objects that manage the Print panel’s accessory views.
//
// # Running the Panel
//
//   - [NSPrintPanel.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
//   - [NSPrintPanel.RunModal]: Displays the Print panel and begins the modal loop.
//   - [NSPrintPanel.RunModalWithPrintInfo]: Displays the Print panel and runs the modal loop using the specified printing information.
//
// # Accessing the Printing Information
//
//   - [NSPrintPanel.PrintInfo]: The information associated with the running Print panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel
//
// [Human Interface Guidelines]: https://developer.apple.com/design/human-interface-guidelines/macos/system-capabilities/printing/
type NSPrintPanel struct {
	objectivec.Object
}

// NSPrintPanelFromID constructs a [NSPrintPanel] from an objc.ID.
//
// The Print panel that queries the user for information about a print job.
func NSPrintPanelFromID(id objc.ID) NSPrintPanel {
	return NSPrintPanel{objectivec.Object{ID: id}}
}

// NOTE: NSPrintPanel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPrintPanel] class.
//
// # Customizing the Panel
//
//   - [INSPrintPanel.JobStyleHint]: The type of settings that the print panel displays.
//   - [INSPrintPanel.SetJobStyleHint]
//   - [INSPrintPanel.Options]: The current configuration options for the Print panel.
//   - [INSPrintPanel.SetOptions]
//   - [INSPrintPanel.DefaultButtonTitle]: Returns the title of the Print panel’s default button.
//   - [INSPrintPanel.SetDefaultButtonTitle]: Sets the title of the Print panel’s default button.
//   - [INSPrintPanel.HelpAnchor]: The HTML help anchor associated with the Print panel.
//   - [INSPrintPanel.SetHelpAnchor]
//
// # Managing Accessory Views
//
//   - [INSPrintPanel.AddAccessoryController]: Adds a custom controller to the Print panel to manage an accessory view.
//   - [INSPrintPanel.RemoveAccessoryController]: Removes the specified controller and accessory view from the Print panel.
//   - [INSPrintPanel.AccessoryControllers]: The array of controller objects that manage the Print panel’s accessory views.
//
// # Running the Panel
//
//   - [INSPrintPanel.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
//   - [INSPrintPanel.RunModal]: Displays the Print panel and begins the modal loop.
//   - [INSPrintPanel.RunModalWithPrintInfo]: Displays the Print panel and runs the modal loop using the specified printing information.
//
// # Accessing the Printing Information
//
//   - [INSPrintPanel.PrintInfo]: The information associated with the running Print panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel
type INSPrintPanel interface {
	objectivec.IObject

	// Topic: Customizing the Panel

	// The type of settings that the print panel displays.
	JobStyleHint() NSPrintPanelJobStyleHint
	SetJobStyleHint(value NSPrintPanelJobStyleHint)
	// The current configuration options for the Print panel.
	Options() NSPrintPanelOptions
	SetOptions(value NSPrintPanelOptions)
	// Returns the title of the Print panel’s default button.
	DefaultButtonTitle() string
	// Sets the title of the Print panel’s default button.
	SetDefaultButtonTitle(defaultButtonTitle string)
	// The HTML help anchor associated with the Print panel.
	HelpAnchor() NSHelpAnchorName
	SetHelpAnchor(value NSHelpAnchorName)

	// Topic: Managing Accessory Views

	// Adds a custom controller to the Print panel to manage an accessory view.
	AddAccessoryController(accessoryController INSViewController)
	// Removes the specified controller and accessory view from the Print panel.
	RemoveAccessoryController(accessoryController INSViewController)
	// The array of controller objects that manage the Print panel’s accessory views.
	AccessoryControllers() []NSViewController

	// Topic: Running the Panel

	BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo INSPrintInfo, parentWindow INSWindow, handler PrintPanelResultHandler)
	// Displays the Print panel and begins the modal loop.
	RunModal() int
	// Displays the Print panel and runs the modal loop using the specified printing information.
	RunModalWithPrintInfo(printInfo INSPrintInfo) int

	// Topic: Accessing the Printing Information

	// The information associated with the running Print panel.
	PrintInfo() INSPrintInfo
}

// Init initializes the instance.
func (p NSPrintPanel) Init() NSPrintPanel {
	rv := objc.Send[NSPrintPanel](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPrintPanel) Autorelease() NSPrintPanel {
	rv := objc.Send[NSPrintPanel](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPrintPanel creates a new NSPrintPanel instance.
func NewNSPrintPanel() NSPrintPanel {
	class := getNSPrintPanelClass()
	rv := objc.Send[NSPrintPanel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the title of the Print panel’s default button.
//
// # Return Value
//
// The title of the default button.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/defaultButtonTitle()
func (p NSPrintPanel) DefaultButtonTitle() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("defaultButtonTitle"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the title of the Print panel’s default button.
//
// defaultButtonTitle: The string to use for the button title.
//
// # Discussion
//
// You can use this method to change the default button title from “Print”
// to something more appropriate for your usage of the panel. For example, if
// you are using the Print panel to save a representation of the document to a
// file, you might change the title to “Save”.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/setDefaultButtonTitle(_:)
func (p NSPrintPanel) SetDefaultButtonTitle(defaultButtonTitle string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setDefaultButtonTitle:"), objc.String(defaultButtonTitle))
}

// Adds a custom controller to the Print panel to manage an accessory view.
//
// accessoryController: The view controller that manages your custom accessory views.
//
// # Discussion
//
// You can invoke this method multiple times to add multiple accessory views
// to the receiver’s Print panel.
//
// The title for the accessory view is obtained from the `title` method of the
// view controller object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/addAccessoryController(_:)
func (p NSPrintPanel) AddAccessoryController(accessoryController INSViewController) {
	objc.Send[objc.ID](p.ID, objc.Sel("addAccessoryController:"), accessoryController)
}

// Removes the specified controller and accessory view from the Print panel.
//
// accessoryController: The view controller to remove.
//
// # Discussion
//
// You use this method to remove any view controllers responsible for
// displaying accessory views you do not want to include in the Print panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/removeAccessoryController(_:)
func (p NSPrintPanel) RemoveAccessoryController(accessoryController INSViewController) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeAccessoryController:"), accessoryController)
}

// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/beginSheet(using:on:completionHandler:)
func (p NSPrintPanel) BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo INSPrintInfo, parentWindow INSWindow, handler PrintPanelResultHandler) {
	_block2, _ := NewPrintPanelResultBlock(handler)
	objc.Send[objc.ID](p.ID, objc.Sel("beginSheetUsingPrintInfo:onWindow:completionHandler:"), printInfo, parentWindow, _block2)
}

// Displays the Print panel and begins the modal loop.
//
// # Return Value
//
// [NSCancelButton] if the user clicks the Cancel button; otherwise
// [NSOKButton].
//
// # Discussion
//
// This method uses the printing information associated with the current
// printing operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/runModal()
func (p NSPrintPanel) RunModal() int {
	rv := objc.Send[int](p.ID, objc.Sel("runModal"))
	return rv
}

// Displays the Print panel and runs the modal loop using the specified
// printing information.
//
// printInfo: The printing information to use while displaying the Print panel.
//
// # Return Value
//
// [NSCancelButton] if the user clicks the Cancel button; otherwise
// [NSOKButton].
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/runModal(with:)
func (p NSPrintPanel) RunModalWithPrintInfo(printInfo INSPrintInfo) int {
	rv := objc.Send[int](p.ID, objc.Sel("runModalWithPrintInfo:"), printInfo)
	return rv
}

// Returns a new print panel object.
//
// # Return Value
//
// The print panel object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/printPanel
func (_NSPrintPanelClass NSPrintPanelClass) PrintPanel() NSPrintPanel {
	rv := objc.Send[objc.ID](objc.ID(_NSPrintPanelClass.class), objc.Sel("printPanel"))
	return NSPrintPanelFromID(rv)
}

// The type of settings that the print panel displays.
//
// # Discussion
//
// This property controls the set of items that appear in the Presets menu of
// the simplified Print panel interface. For a list of supported job style
// hints, see `Job Style Hints`. Set this property to `nil` to deactivate the
// simplified Print panel interface and use the standard interface instead
// (the equivalent of Core Printing’s `kPMPresetGraphicsTypeGeneral`).
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/jobStyleHint-swift.property
func (p NSPrintPanel) JobStyleHint() NSPrintPanelJobStyleHint {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("jobStyleHint"))
	return NSPrintPanelJobStyleHint(foundation.NSStringFromID(rv).String())
}
func (p NSPrintPanel) SetJobStyleHint(value NSPrintPanelJobStyleHint) {
	objc.Send[struct{}](p.ID, objc.Sel("setJobStyleHint:"), objc.String(string(value)))
}

// The current configuration options for the Print panel.
//
// # Discussion
//
// You can specify multiple options by adding them together. For a list of
// supported options, see [NSPrintPanel.Options].
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/options-swift.property
//
// [NSPrintPanel.Options]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct
func (p NSPrintPanel) Options() NSPrintPanelOptions {
	rv := objc.Send[NSPrintPanelOptions](p.ID, objc.Sel("options"))
	return NSPrintPanelOptions(rv)
}
func (p NSPrintPanel) SetOptions(value NSPrintPanelOptions) {
	objc.Send[struct{}](p.ID, objc.Sel("setOptions:"), value)
}

// The HTML help anchor associated with the Print panel.
//
// # Discussion
//
// Use this property to specify the anchor name in your Apple Help file. The
// string you assign should contain only the name portion of the HTML anchor
// element.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/helpAnchor
func (p NSPrintPanel) HelpAnchor() NSHelpAnchorName {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("helpAnchor"))
	return NSHelpAnchorName(foundation.NSStringFromID(rv).String())
}
func (p NSPrintPanel) SetHelpAnchor(value NSHelpAnchorName) {
	objc.Send[struct{}](p.ID, objc.Sel("setHelpAnchor:"), objc.String(string(value)))
}

// The array of controller objects that manage the Print panel’s accessory
// views.
//
// # Discussion
//
// This property contains an array of [NSViewController] objects, each of
// which represents an accessory view added using the [AddAccessoryController]
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/accessoryControllers
func (p NSPrintPanel) AccessoryControllers() []NSViewController {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("accessoryControllers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSViewController {
		return NSViewControllerFromID(id)
	})
}

// The information associated with the running Print panel.
//
// # Discussion
//
// The value in this property is `nil` if the Print panel is not currently
// running.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/printInfo
func (p NSPrintPanel) PrintInfo() INSPrintInfo {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("printInfo"))
	return NSPrintInfoFromID(objc.ID(rv))
}

// BeginSheetUsingPrintInfoOnWindow is a synchronous wrapper around [NSPrintPanel.BeginSheetUsingPrintInfoOnWindowCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p NSPrintPanel) BeginSheetUsingPrintInfoOnWindow(ctx context.Context, printInfo INSPrintInfo, parentWindow INSWindow) (NSPrintPanelResult, error) {
	done := make(chan NSPrintPanelResult, 1)
	p.BeginSheetUsingPrintInfoOnWindowCompletionHandler(printInfo, parentWindow, func(val NSPrintPanelResult) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
