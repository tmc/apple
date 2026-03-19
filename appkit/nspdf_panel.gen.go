// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPDFPanel] class.
var (
	_NSPDFPanelClass     NSPDFPanelClass
	_NSPDFPanelClassOnce sync.Once
)

func getNSPDFPanelClass() NSPDFPanelClass {
	_NSPDFPanelClassOnce.Do(func() {
		_NSPDFPanelClass = NSPDFPanelClass{class: objc.GetClass("NSPDFPanel")}
	})
	return _NSPDFPanelClass
}

// GetNSPDFPanelClass returns the class object for NSPDFPanel.
func GetNSPDFPanelClass() NSPDFPanelClass {
	return getNSPDFPanelClass()
}

type NSPDFPanelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPDFPanelClass) Alloc() NSPDFPanel {
	rv := objc.Send[NSPDFPanel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A Save or Export as PDF panel that’s consistent with the macOS user
// interface.
//
// # Overview
// 
// A PDF panel has a variety of built-in customization controls, such as page
// orientation, paper size, and tags. It also supports the use of a custom
// accessory view controller that allows an app to specify how a PDF file
// should be created.
//
// # Managing the Contents of a PDF Panel
//
//   - [NSPDFPanel.AccessoryController]: A view controller for the accessory view that the panel can present.
//   - [NSPDFPanel.SetAccessoryController]
//   - [NSPDFPanel.Options]: A set of configuration options that determine the accessory views the PDF panel should display.
//   - [NSPDFPanel.SetOptions]
//   - [NSPDFPanel.DefaultFileName]: The initial value for the user-editable filename shown in the name field of the PDF panel.
//   - [NSPDFPanel.SetDefaultFileName]
//
// # Displaying a PDF Panel
//
//   - [NSPDFPanel.BeginSheetWithPDFInfoModalForWindowCompletionHandler]: Presents a document-modal PDF panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFPanel
type NSPDFPanel struct {
	objectivec.Object
}

// NSPDFPanelFromID constructs a [NSPDFPanel] from an objc.ID.
//
// A Save or Export as PDF panel that’s consistent with the macOS user
// interface.
func NSPDFPanelFromID(id objc.ID) NSPDFPanel {
	return NSPDFPanel{objectivec.Object{ID: id}}
}
// NOTE: NSPDFPanel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPDFPanel] class.
//
// # Managing the Contents of a PDF Panel
//
//   - [INSPDFPanel.AccessoryController]: A view controller for the accessory view that the panel can present.
//   - [INSPDFPanel.SetAccessoryController]
//   - [INSPDFPanel.Options]: A set of configuration options that determine the accessory views the PDF panel should display.
//   - [INSPDFPanel.SetOptions]
//   - [INSPDFPanel.DefaultFileName]: The initial value for the user-editable filename shown in the name field of the PDF panel.
//   - [INSPDFPanel.SetDefaultFileName]
//
// # Displaying a PDF Panel
//
//   - [INSPDFPanel.BeginSheetWithPDFInfoModalForWindowCompletionHandler]: Presents a document-modal PDF panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFPanel
type INSPDFPanel interface {
	objectivec.IObject

	// Topic: Managing the Contents of a PDF Panel

	// A view controller for the accessory view that the panel can present.
	AccessoryController() INSViewController
	SetAccessoryController(value INSViewController)
	// A set of configuration options that determine the accessory views the PDF panel should display.
	Options() NSPDFPanelOptions
	SetOptions(value NSPDFPanelOptions)
	// The initial value for the user-editable filename shown in the name field of the PDF panel.
	DefaultFileName() string
	SetDefaultFileName(value string)

	// Topic: Displaying a PDF Panel

	// Presents a document-modal PDF panel.
	BeginSheetWithPDFInfoModalForWindowCompletionHandler(pdfInfo INSPDFInfo, docWindow INSWindow, completionHandler IntHandler)
}

// Init initializes the instance.
func (p NSPDFPanel) Init() NSPDFPanel {
	rv := objc.Send[NSPDFPanel](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPDFPanel) Autorelease() NSPDFPanel {
	rv := objc.Send[NSPDFPanel](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPDFPanel creates a new NSPDFPanel instance.
func NewNSPDFPanel() NSPDFPanel {
	class := getNSPDFPanelClass()
	rv := objc.Send[NSPDFPanel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Presents a document-modal PDF panel.
//
// pdfInfo: The [NSPDFInfo] object describing the parameters to be used when creating
// the PDF file.
//
// docWindow: The window in which the PDF panel will be presented.
//
// completionHandler: The block called when the user dismisses the PDF panel.
//
// # Discussion
// 
// This method presents a slightly different PDF panel depending on whether
// the [PDFPanelRequestsParentDirectory] constant is set. If the user
// dismisses the panel without canceling it, this method updates the
// [NSPDFInfo] object with any changes the user makes.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFPanel/beginSheet(with:modalFor:completionHandler:)
func (p NSPDFPanel) BeginSheetWithPDFInfoModalForWindowCompletionHandler(pdfInfo INSPDFInfo, docWindow INSWindow, completionHandler IntHandler) {
_block2, _cleanup2 := NewIntBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](p.ID, objc.Sel("beginSheetWithPDFInfo:modalForWindow:completionHandler:"), pdfInfo, docWindow, _block2)
}

// A view controller for the accessory view that the panel can present.
//
// # Discussion
// 
// The PDF panel passes an [NSPDFInfo] object to the accessory view controller
// to display the various attributes associated with the PDF file. Unlike a
// print panel (that is, an [NSPrintPanel] object), a PDF panel can have only
// one accessory view.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFPanel/accessoryController
func (p NSPDFPanel) AccessoryController() INSViewController {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessoryController"))
	return NSViewControllerFromID(objc.ID(rv))
}
func (p NSPDFPanel) SetAccessoryController(value INSViewController) {
	objc.Send[struct{}](p.ID, objc.Sel("setAccessoryController:"), value)
}
// A set of configuration options that determine the accessory views the PDF
// panel should display.
//
// # Discussion
// 
// You specify a set of options by combining the appropriate constants defined
// in [NSPDFPanel.Options].
//
// [NSPDFPanel.Options]: https://developer.apple.com/documentation/AppKit/NSPDFPanel/Options-swift.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFPanel/options-swift.property
func (p NSPDFPanel) Options() NSPDFPanelOptions {
	rv := objc.Send[NSPDFPanelOptions](p.ID, objc.Sel("options"))
	return NSPDFPanelOptions(rv)
}
func (p NSPDFPanel) SetOptions(value NSPDFPanelOptions) {
	objc.Send[struct{}](p.ID, objc.Sel("setOptions:"), value)
}
// The initial value for the user-editable filename shown in the name field of
// the PDF panel.
//
// # Discussion
// 
// The `defaultFileName` string should never include a file extension. By
// default, the string’s value is “Untitled” (or its equivalent for the
// current locale).
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFPanel/defaultFileName
func (p NSPDFPanel) DefaultFileName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("defaultFileName"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPDFPanel) SetDefaultFileName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setDefaultFileName:"), objc.String(value))
}

// BeginSheetWithPDFInfoModalForWindow is a synchronous wrapper around [NSPDFPanel.BeginSheetWithPDFInfoModalForWindowCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p NSPDFPanel) BeginSheetWithPDFInfoModalForWindow(ctx context.Context, pdfInfo INSPDFInfo, docWindow INSWindow) (int, error) {
	done := make(chan int, 1)
	p.BeginSheetWithPDFInfoModalForWindowCompletionHandler(pdfInfo, docWindow, func(val int) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

