// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPrintOperation] class.
var (
	_NSPrintOperationClass     NSPrintOperationClass
	_NSPrintOperationClassOnce sync.Once
)

func getNSPrintOperationClass() NSPrintOperationClass {
	_NSPrintOperationClassOnce.Do(func() {
		_NSPrintOperationClass = NSPrintOperationClass{class: objc.GetClass("NSPrintOperation")}
	})
	return _NSPrintOperationClass
}

// GetNSPrintOperationClass returns the class object for NSPrintOperation.
func GetNSPrintOperationClass() NSPrintOperationClass {
	return getNSPrintOperationClass()
}

type NSPrintOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPrintOperationClass) Alloc() NSPrintOperation {
	rv := objc.Send[NSPrintOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that controls operations that generate Encapsulated PostScript
// (EPS) code, Portable Document Format (PDF) code, or print jobs.
//
// # Overview
// 
// An [NSPrintOperation] object works in conjunction with two other objects:
// an [NSPrintInfo] object, which specifies how the code should be generated,
// and an [NSView] object, which generates the actual code.
// 
// It is important to note that the majority of methods in [NSPrintOperation]
// copy the instance of [NSPrintInfo] passed into them. Future changes to that
// print info are not reflected in the print info retained by the current
// [NSPrintOperation] object. All changes should be made to the print info
// before passing to the methods of this class. The only method in
// [NSPrintOperation] which does not copy the [NSPrintInfo] instance is
// [PrintInfo].
//
// # Determining the Type of Operation
//
//   - [NSPrintOperation.CopyingOperation]: A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.
//
// # Modifying the Printing Information
//
//   - [NSPrintOperation.PrintInfo]: The printing information associated with the print operation.
//   - [NSPrintOperation.SetPrintInfo]
//
// # Getting the View
//
//   - [NSPrintOperation.View]: The view object that generates the actual data for the print operation.
//
// # Getting the Printing Quality
//
//   - [NSPrintOperation.PreferredRenderingQuality]: The printing quality.
//
// # Running the Print Operation
//
//   - [NSPrintOperation.RunOperation]: Runs the print operation on the current thread.
//   - [NSPrintOperation.RunOperationModalForWindowDelegateDidRunSelectorContextInfo]: Runs the print operation, calling your custom delegate method upon completion.
//   - [NSPrintOperation.CleanUpOperation]: Called at the end of a print operation to remove the print operation as the current operation.
//   - [NSPrintOperation.DeliverResult]: Delivers the results of the print operation to the intended destination.
//
// # Modifying the User Interface
//
//   - [NSPrintOperation.ShowsPrintPanel]: A Boolean value that determines whether the print operation displays a print panel.
//   - [NSPrintOperation.SetShowsPrintPanel]
//   - [NSPrintOperation.ShowsProgressPanel]: A Boolean value that determines whether the print operation displays a progress panel.
//   - [NSPrintOperation.SetShowsProgressPanel]
//   - [NSPrintOperation.JobTitle]: The custom title of the print job.
//   - [NSPrintOperation.SetJobTitle]
//   - [NSPrintOperation.PrintPanel]: The print panel object to use during the operation.
//   - [NSPrintOperation.SetPrintPanel]
//   - [NSPrintOperation.PDFPanel]: The PDF panel object to use during the operation.
//   - [NSPrintOperation.SetPDFPanel]
//
// # Managing the Drawing Context
//
//   - [NSPrintOperation.Context]: The graphics context object used for generating output.
//   - [NSPrintOperation.CreateContext]: Creates the graphics context object used for drawing during the operation.
//   - [NSPrintOperation.DestroyContext]: Destroys the print operation’s graphics context.
//
// # Managing Page Information
//
//   - [NSPrintOperation.CurrentPage]: The current page number being printed.
//   - [NSPrintOperation.PageRange]: The range of pages associated with the print operation.
//   - [NSPrintOperation.PageOrder]: The print order for the pages of the operation.
//   - [NSPrintOperation.SetPageOrder]
//
// # Managing Printing Threads
//
//   - [NSPrintOperation.CanSpawnSeparateThread]: A Boolean value that determines whether the print operation is allowed to spawn a separate printing thread.
//   - [NSPrintOperation.SetCanSpawnSeparateThread]
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation
type NSPrintOperation struct {
	objectivec.Object
}

// NSPrintOperationFromID constructs a [NSPrintOperation] from an objc.ID.
//
// An object that controls operations that generate Encapsulated PostScript
// (EPS) code, Portable Document Format (PDF) code, or print jobs.
func NSPrintOperationFromID(id objc.ID) NSPrintOperation {
	return NSPrintOperation{objectivec.Object{ID: id}}
}
// NOTE: NSPrintOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPrintOperation] class.
//
// # Determining the Type of Operation
//
//   - [INSPrintOperation.CopyingOperation]: A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.
//
// # Modifying the Printing Information
//
//   - [INSPrintOperation.PrintInfo]: The printing information associated with the print operation.
//   - [INSPrintOperation.SetPrintInfo]
//
// # Getting the View
//
//   - [INSPrintOperation.View]: The view object that generates the actual data for the print operation.
//
// # Getting the Printing Quality
//
//   - [INSPrintOperation.PreferredRenderingQuality]: The printing quality.
//
// # Running the Print Operation
//
//   - [INSPrintOperation.RunOperation]: Runs the print operation on the current thread.
//   - [INSPrintOperation.RunOperationModalForWindowDelegateDidRunSelectorContextInfo]: Runs the print operation, calling your custom delegate method upon completion.
//   - [INSPrintOperation.CleanUpOperation]: Called at the end of a print operation to remove the print operation as the current operation.
//   - [INSPrintOperation.DeliverResult]: Delivers the results of the print operation to the intended destination.
//
// # Modifying the User Interface
//
//   - [INSPrintOperation.ShowsPrintPanel]: A Boolean value that determines whether the print operation displays a print panel.
//   - [INSPrintOperation.SetShowsPrintPanel]
//   - [INSPrintOperation.ShowsProgressPanel]: A Boolean value that determines whether the print operation displays a progress panel.
//   - [INSPrintOperation.SetShowsProgressPanel]
//   - [INSPrintOperation.JobTitle]: The custom title of the print job.
//   - [INSPrintOperation.SetJobTitle]
//   - [INSPrintOperation.PrintPanel]: The print panel object to use during the operation.
//   - [INSPrintOperation.SetPrintPanel]
//   - [INSPrintOperation.PDFPanel]: The PDF panel object to use during the operation.
//   - [INSPrintOperation.SetPDFPanel]
//
// # Managing the Drawing Context
//
//   - [INSPrintOperation.Context]: The graphics context object used for generating output.
//   - [INSPrintOperation.CreateContext]: Creates the graphics context object used for drawing during the operation.
//   - [INSPrintOperation.DestroyContext]: Destroys the print operation’s graphics context.
//
// # Managing Page Information
//
//   - [INSPrintOperation.CurrentPage]: The current page number being printed.
//   - [INSPrintOperation.PageRange]: The range of pages associated with the print operation.
//   - [INSPrintOperation.PageOrder]: The print order for the pages of the operation.
//   - [INSPrintOperation.SetPageOrder]
//
// # Managing Printing Threads
//
//   - [INSPrintOperation.CanSpawnSeparateThread]: A Boolean value that determines whether the print operation is allowed to spawn a separate printing thread.
//   - [INSPrintOperation.SetCanSpawnSeparateThread]
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation
type INSPrintOperation interface {
	objectivec.IObject

	// Topic: Determining the Type of Operation

	// A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.
	CopyingOperation() bool

	// Topic: Modifying the Printing Information

	// The printing information associated with the print operation.
	PrintInfo() INSPrintInfo
	SetPrintInfo(value INSPrintInfo)

	// Topic: Getting the View

	// The view object that generates the actual data for the print operation.
	View() INSView

	// Topic: Getting the Printing Quality

	// The printing quality.
	PreferredRenderingQuality() NSPrintRenderingQuality

	// Topic: Running the Print Operation

	// Runs the print operation on the current thread.
	RunOperation() bool
	// Runs the print operation, calling your custom delegate method upon completion.
	RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow INSWindow, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo uintptr)
	// Called at the end of a print operation to remove the print operation as the current operation.
	CleanUpOperation()
	// Delivers the results of the print operation to the intended destination.
	DeliverResult() bool

	// Topic: Modifying the User Interface

	// A Boolean value that determines whether the print operation displays a print panel.
	ShowsPrintPanel() bool
	SetShowsPrintPanel(value bool)
	// A Boolean value that determines whether the print operation displays a progress panel.
	ShowsProgressPanel() bool
	SetShowsProgressPanel(value bool)
	// The custom title of the print job.
	JobTitle() string
	SetJobTitle(value string)
	// The print panel object to use during the operation.
	PrintPanel() INSPrintPanel
	SetPrintPanel(value INSPrintPanel)
	// The PDF panel object to use during the operation.
	PDFPanel() INSPDFPanel
	SetPDFPanel(value INSPDFPanel)

	// Topic: Managing the Drawing Context

	// The graphics context object used for generating output.
	Context() INSGraphicsContext
	// Creates the graphics context object used for drawing during the operation.
	CreateContext() INSGraphicsContext
	// Destroys the print operation’s graphics context.
	DestroyContext()

	// Topic: Managing Page Information

	// The current page number being printed.
	CurrentPage() int
	// The range of pages associated with the print operation.
	PageRange() foundation.NSRange
	// The print order for the pages of the operation.
	PageOrder() NSPrintingPageOrder
	SetPageOrder(value NSPrintingPageOrder)

	// Topic: Managing Printing Threads

	// A Boolean value that determines whether the print operation is allowed to spawn a separate printing thread.
	CanSpawnSeparateThread() bool
	SetCanSpawnSeparateThread(value bool)
}

// Init initializes the instance.
func (p NSPrintOperation) Init() NSPrintOperation {
	rv := objc.Send[NSPrintOperation](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPrintOperation) Autorelease() NSPrintOperation {
	rv := objc.Send[NSPrintOperation](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPrintOperation creates a new NSPrintOperation instance.
func NewNSPrintOperation() NSPrintOperation {
	class := getNSPrintOperationClass()
	rv := objc.Send[NSPrintOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns an print operation object ready to control the printing
// of the specified view.
//
// view: The view whose contents you want to print.
//
// # Return Value
// 
// The new [NSPrintOperation] object. You must run the operation to print the
// view.
//
// # Discussion
// 
// The new [NSPrintOperation] object uses the settings stored in the shared
// [NSPrintInfo] object. This method raises an
// [NSPrintOperationExistsException] if there is already a print operation in
// progress; otherwise the returned object is made the current print operation
// for this thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/init(view:)
func NewPrintOperationWithView(view INSView) NSPrintOperation {
	rv := objc.Send[objc.ID](objc.ID(getNSPrintOperationClass().class), objc.Sel("printOperationWithView:"), view)
	return NSPrintOperationFromID(rv)
}

// Creates and returns an print operation object ready to control the printing
// of the specified view using custom print settings.
//
// view: The view whose contents you want to print.
//
// printInfo: The print settings to use when printing the view.
//
// # Return Value
// 
// The new [NSPrintOperation] object. You must run the operation to print the
// view.
//
// # Discussion
// 
// This method raises an [NSPrintOperationExistsException] if there is already
// a print operation in progress; otherwise the returned object is made the
// current print operation for this thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/init(view:printInfo:)
func NewPrintOperationWithViewPrintInfo(view INSView, printInfo INSPrintInfo) NSPrintOperation {
	rv := objc.Send[objc.ID](objc.ID(getNSPrintOperationClass().class), objc.Sel("printOperationWithView:printInfo:"), view, printInfo)
	return NSPrintOperationFromID(rv)
}

// Runs the print operation on the current thread.
//
// # Return Value
// 
// [true] if the operation was successful; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The operation runs to completion in the current thread, blocking the
// application. A separate thread is not spawned, even if
// [CanSpawnSeparateThread] is [true]. Use
// [RunOperationModalForWindowDelegateDidRunSelectorContextInfo] to use
// document-modal sheets and to allow a separate thread to perform the
// operation.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/run()
func (p NSPrintOperation) RunOperation() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("runOperation"))
	return rv
}

// Runs the print operation, calling your custom delegate method upon
// completion.
//
// docWindow: The document window to receive a print progress sheet.
//
// delegate: The printing delegate object. Messages are sent to this object.
//
// didRunSelector: The delegate method called after the completion of the print operation.
//
// contextInfo: A pointer to any data you want passed to the method in the `didRunSelector`
// parameter.
//
// # Discussion
// 
// The method specified by the `didRunSelector` parameter must have the
// following signature:
// 
// The value of `success` is [true] if the print operation ran to completion
// without cancellation or error, and [false] otherwise.
// 
// If you send [CanSpawnSeparateThread] to an [NSPrintOperation] object with
// an argument of [true], then the delegate specified in a subsequent
// invocation of [RunOperationModalForWindowDelegateDidRunSelectorContextInfo]
// may be messaged in that spawned, non-main thread.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/runModal(for:delegate:didRun:contextInfo:)
func (p NSPrintOperation) RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow INSWindow, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](p.ID, objc.Sel("runOperationModalForWindow:delegate:didRunSelector:contextInfo:"), docWindow, delegate, didRunSelector, contextInfo)
}

// Called at the end of a print operation to remove the print operation as the
// current operation.
//
// # Discussion
// 
// You typically do not invoke this method directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/cleanUp()
func (p NSPrintOperation) CleanUpOperation() {
	objc.Send[objc.ID](p.ID, objc.Sel("cleanUpOperation"))
}

// Delivers the results of the print operation to the intended destination.
//
// # Return Value
// 
// [true] if the results were successfully delivered; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method may be called to deliver the results to the printer spool or
// preview application. Do not invoke this method directly—it is invoked
// automatically when the print operation is done.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/deliverResult()
func (p NSPrintOperation) DeliverResult() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("deliverResult"))
	return rv
}

// Creates the graphics context object used for drawing during the operation.
//
// # Return Value
// 
// The graphics context object used for drawing. This object is created using
// the settings from the receiver’s [NSPrintInfo] object.
//
// # Discussion
// 
// Do not invoke this method directly—it is invoked before any output is
// generated.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/createContext()
func (p NSPrintOperation) CreateContext() INSGraphicsContext {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("createContext"))
	return NSGraphicsContextFromID(rv)
}

// Destroys the print operation’s graphics context.
//
// # Discussion
// 
// Do not invoke this method directly—it is invoked at the end of a print
// operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/destroyContext()
func (p NSPrintOperation) DestroyContext() {
	objc.Send[objc.ID](p.ID, objc.Sel("destroyContext"))
}

// Creates and returns a new print operation object ready to control the
// copying of EPS graphics from the specified view.
//
// view: The view containing the data to be turned into EPS data.
//
// rect: The portion of the view (specified in points in the view’s coordinate
// space) to be rendered as EPS data.
//
// data: An empty [NSMutableData] object. After the job is run, this object contains
// the EPS data.
//
// # Return Value
// 
// The new [NSPrintOperation] object. You must run the operation to generate
// the EPS data.
//
// # Discussion
// 
// The new [NSPrintOperation] object uses the default [NSPrintInfo] object.
// This method raises an [NSPrintOperationExistsException] if there is already
// a print operation in progress; otherwise the returned object is made the
// current print operation for this thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/epsOperation(with:inside:to:)
func (_NSPrintOperationClass NSPrintOperationClass) EPSOperationWithViewInsideRectToData(view INSView, rect corefoundation.CGRect, data foundation.NSMutableData) NSPrintOperation {
	rv := objc.Send[objc.ID](objc.ID(_NSPrintOperationClass.class), objc.Sel("EPSOperationWithView:insideRect:toData:"), view, rect, data)
	return NSPrintOperationFromID(rv)
}

// Creates and returns a new print operation object ready to control the
// copying of EPS graphics from the specified view using the specified print
// settings.
//
// view: The view containing the data to be turned into EPS data.
//
// rect: The portion of the view (specified in points in the view’s coordinate
// space) to be rendered as EPS data.
//
// data: An empty [NSMutableData] object. After the job is run, this object contains
// the EPS data.
//
// printInfo: The print settings to use when generating the EPS data.
//
// # Return Value
// 
// The new [NSPrintOperation] object. You must run the operation to generate
// the EPS data.
//
// # Discussion
// 
// This method raises an [NSPrintOperationExistsException] if there is already
// a print operation in progress; otherwise the returned object is made the
// current print operation for this thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/epsOperation(with:inside:to:printInfo:)
func (_NSPrintOperationClass NSPrintOperationClass) EPSOperationWithViewInsideRectToDataPrintInfo(view INSView, rect corefoundation.CGRect, data foundation.NSMutableData, printInfo INSPrintInfo) NSPrintOperation {
	rv := objc.Send[objc.ID](objc.ID(_NSPrintOperationClass.class), objc.Sel("EPSOperationWithView:insideRect:toData:printInfo:"), view, rect, data, printInfo)
	return NSPrintOperationFromID(rv)
}

// Creates and returns a new print operation object ready to control the
// copying of EPS graphics from the specified view and write the resulting
// data to the specified file.
//
// view: The view containing the data to be turned into EPS data.
//
// rect: The portion of the view (specified in points in the view’s coordinate
// space) to be rendered as EPS data.
//
// path: The path to a file. After the job is run, this file contains the EPS data.
//
// printInfo: The print settings to use when generating the EPS data.
//
// # Return Value
// 
// The new [NSPrintOperation] object. You must run the operation to generate
// the EPS data.
//
// # Discussion
// 
// This method raises an [NSPrintOperationExistsException] if there is already
// a print operation in progress; otherwise the returned object is made the
// current print operation for this thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/epsOperation(with:inside:toPath:printInfo:)
func (_NSPrintOperationClass NSPrintOperationClass) EPSOperationWithViewInsideRectToPathPrintInfo(view INSView, rect corefoundation.CGRect, path string, printInfo INSPrintInfo) NSPrintOperation {
	rv := objc.Send[objc.ID](objc.ID(_NSPrintOperationClass.class), objc.Sel("EPSOperationWithView:insideRect:toPath:printInfo:"), view, rect, objc.String(path), printInfo)
	return NSPrintOperationFromID(rv)
}

// Creates and returns a new print operation object ready to control the
// copying of PDF graphics from the specified view.
//
// view: The view containing the data to be turned into PDF data.
//
// rect: The portion of the view (specified in points in the view’s coordinate
// space) to be rendered as PDF data.
//
// data: An empty [NSMutableData] object. After the job is run, this object contains
// the PDF data.
//
// # Return Value
// 
// The new [NSPrintOperation] object. You must run the operation to generate
// the PDF data.
//
// # Discussion
// 
// The new [NSPrintOperation] object uses the default [NSPrintInfo] object.
// This method raises an [NSPrintOperationExistsException] if there is already
// a print operation in progress; otherwise the returned object is made the
// current print operation for this thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfOperation(with:inside:to:)
func (_NSPrintOperationClass NSPrintOperationClass) PDFOperationWithViewInsideRectToData(view INSView, rect corefoundation.CGRect, data foundation.NSMutableData) NSPrintOperation {
	rv := objc.Send[objc.ID](objc.ID(_NSPrintOperationClass.class), objc.Sel("PDFOperationWithView:insideRect:toData:"), view, rect, data)
	return NSPrintOperationFromID(rv)
}

// Creates and returns a new print operation object ready to control the
// copying of PDF graphics from the specified view using the specified print
// settings.
//
// view: The view containing the data to be turned into PDF data.
//
// rect: The portion of the view (specified in points in the view’s coordinate
// space) to be rendered as PDF data.
//
// data: An empty [NSMutableData] object. After the job is run, this object contains
// the PDF data.
//
// printInfo: The print settings to use when generating the PDF data.
//
// # Return Value
// 
// The new [NSPrintOperation] object. You must run the operation to generate
// the PDF data.
//
// # Discussion
// 
// This method raises an [NSPrintOperationExistsException] if there is already
// a print operation in progress; otherwise the returned object is made the
// current print operation for this thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfOperation(with:inside:to:printInfo:)
func (_NSPrintOperationClass NSPrintOperationClass) PDFOperationWithViewInsideRectToDataPrintInfo(view INSView, rect corefoundation.CGRect, data foundation.NSMutableData, printInfo INSPrintInfo) NSPrintOperation {
	rv := objc.Send[objc.ID](objc.ID(_NSPrintOperationClass.class), objc.Sel("PDFOperationWithView:insideRect:toData:printInfo:"), view, rect, data, printInfo)
	return NSPrintOperationFromID(rv)
}

// Creates and returns a new print operation object ready to control the
// copying of PDF graphics from the specified view and write the resulting
// data to the specified file.
//
// view: The view containing the data to be turned into PDF data.
//
// rect: The portion of the view (specified in points in the view’s coordinate
// space) to be rendered as PDF data.
//
// path: The path to a file. After the job is run, this file contains the PDF data.
//
// printInfo: The print settings to use when generating the PDF data.
//
// # Return Value
// 
// The new [NSPrintOperation] object. You must run the operation to generate
// the PDF data.
//
// # Discussion
// 
// This method raises an [NSPrintOperationExistsException] if there is already
// a print operation in progress; otherwise the returned object is made the
// current print operation for this thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfOperation(with:inside:toPath:printInfo:)
func (_NSPrintOperationClass NSPrintOperationClass) PDFOperationWithViewInsideRectToPathPrintInfo(view INSView, rect corefoundation.CGRect, path string, printInfo INSPrintInfo) NSPrintOperation {
	rv := objc.Send[objc.ID](objc.ID(_NSPrintOperationClass.class), objc.Sel("PDFOperationWithView:insideRect:toPath:printInfo:"), view, rect, objc.String(path), printInfo)
	return NSPrintOperationFromID(rv)
}

// A Boolean value that indicates whether the print operation is an EPS or PDF
// copy operation.
//
// # Return Value
// 
// [true] if the receiver is an EPS or PDF copy operation; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/isCopyingOperation
func (p NSPrintOperation) CopyingOperation() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isCopyingOperation"))
	return rv
}

// The printing information associated with the print operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printInfo
func (p NSPrintOperation) PrintInfo() INSPrintInfo {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("printInfo"))
	return NSPrintInfoFromID(objc.ID(rv))
}
func (p NSPrintOperation) SetPrintInfo(value INSPrintInfo) {
	objc.Send[struct{}](p.ID, objc.Sel("setPrintInfo:"), value)
}

// The view object that generates the actual data for the print operation.
//
// # Return Value
// 
// The view object that generates the data.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/view
func (p NSPrintOperation) View() INSView {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("view"))
	return NSViewFromID(objc.ID(rv))
}

// The printing quality.
//
// # Return Value
// 
// The preferred printing quality. See [NSPrintOperation.RenderingQuality] for
// the possible values.
// 
// # Discussion
// 
// If the print sheet is unresponsive or sluggish due to the time is takes to
// fully render a page, you can check this method in `` and other printing
// methods such as `beginDocument` and `` to determine if the print operation
// prefers speed over fidelity. Most applications render each page fast enough
// and do not need to call this method. Only use this method after
// establishing that best quality rendering does indeed make the user
// interface unresponsive.
// 
// The following is an example use of this method:
//
// [NSPrintOperation.RenderingQuality]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/RenderingQuality
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/preferredRenderingQuality
func (p NSPrintOperation) PreferredRenderingQuality() NSPrintRenderingQuality {
	rv := objc.Send[NSPrintRenderingQuality](p.ID, objc.Sel("preferredRenderingQuality"))
	return NSPrintRenderingQuality(rv)
}

// A Boolean value that determines whether the print operation displays a
// print panel.
//
// # Discussion
// 
// This method does not affect the display of a progress panel; that operation
// is controlled by the [ShowsProgressPanel] method.
// 
// Operations that generate EPS or PDF data do no display a progress panel,
// regardless of the value in the `flag` parameter.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/showsPrintPanel
func (p NSPrintOperation) ShowsPrintPanel() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("showsPrintPanel"))
	return rv
}
func (p NSPrintOperation) SetShowsPrintPanel(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setShowsPrintPanel:"), value)
}

// A Boolean value that determines whether the print operation displays a
// progress panel.
//
// # Discussion
// 
// This method does not affect the display of a print panel; that operation is
// controlled by the [ShowsPrintPanel] method.
// 
// Operations that generate EPS or PDF data do no display a progress panel,
// regardless of the value in the `flag` parameter.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/showsProgressPanel
func (p NSPrintOperation) ShowsProgressPanel() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("showsProgressPanel"))
	return rv
}
func (p NSPrintOperation) SetShowsProgressPanel(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setShowsProgressPanel:"), value)
}

// The custom title of the print job.
//
// # Discussion
// 
// Assigning a title with this method overrides the job title provided by the
// printing view’s [printJobTitle] method. Specifying `nil` for the
// `jobTitle` parameter causes the receiver to once again take its title from
// the printing view.
//
// [printJobTitle]: https://developer.apple.com/documentation/AppKit/NSView/printJobTitle
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/jobTitle
func (p NSPrintOperation) JobTitle() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("jobTitle"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPrintOperation) SetJobTitle(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setJobTitle:"), objc.String(value))
}

// The print panel object to use during the operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printPanel
func (p NSPrintOperation) PrintPanel() INSPrintPanel {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("printPanel"))
	return NSPrintPanelFromID(objc.ID(rv))
}
func (p NSPrintOperation) SetPrintPanel(value INSPrintPanel) {
	objc.Send[struct{}](p.ID, objc.Sel("setPrintPanel:"), value)
}

// The PDF panel object to use during the operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfPanel
func (p NSPrintOperation) PDFPanel() INSPDFPanel {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("PDFPanel"))
	return NSPDFPanelFromID(objc.ID(rv))
}
func (p NSPrintOperation) SetPDFPanel(value INSPDFPanel) {
	objc.Send[struct{}](p.ID, objc.Sel("setPDFPanel:"), value)
}

// The graphics context object used for generating output.
//
// # Return Value
// 
// The graphics context object used for drawing during the operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/context
func (p NSPrintOperation) Context() INSGraphicsContext {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("context"))
	return NSGraphicsContextFromID(objc.ID(rv))
}

// The current page number being printed.
//
// # Return Value
// 
// The current page being printed.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/currentPage
func (p NSPrintOperation) CurrentPage() int {
	rv := objc.Send[int](p.ID, objc.Sel("currentPage"))
	return rv
}

// The range of pages associated with the print operation.
//
// # Return Value
// 
// The range of page numbers. Page numbers are one-based values where the
// index of page one is 1, the index of page two is 2, and so on. Depending on
// the information returned by the printing view, the starting page number may
// not be 1. Also, if the number of pages being printed is not known, the page
// count may be set to [NSIntegerMax].
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pageRange
func (p NSPrintOperation) PageRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](p.ID, objc.Sel("pageRange"))
	return foundation.NSRange(rv)
}

// The print order for the pages of the operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pageOrder-swift.property
func (p NSPrintOperation) PageOrder() NSPrintingPageOrder {
	rv := objc.Send[NSPrintingPageOrder](p.ID, objc.Sel("pageOrder"))
	return NSPrintingPageOrder(rv)
}
func (p NSPrintOperation) SetPageOrder(value NSPrintingPageOrder) {
	objc.Send[struct{}](p.ID, objc.Sel("setPageOrder:"), value)
}

// A Boolean value that determines whether the print operation is allowed to
// spawn a separate printing thread.
//
// # Discussion
// 
// If `canSpawnSeparateThread` is [true], an [NSThread] object is detached
// when the print panel is dismissed (or immediately, if the panel is not to
// be displayed). The new thread performs the print operation, so that control
// can return to your application. A thread is detached only if the print
// operation is run using the
// [RunOperationModalForWindowDelegateDidRunSelectorContextInfo] method. If
// `canSpawnSeparateThread` is [false], the operation runs on the current
// thread, blocking the application until the operation completes.
// 
// If you send [CanSpawnSeparateThread] to an [NSPrintOperation] object with
// an argument of [true], then the delegate specified in a subsequent
// invocation of [RunOperationModalForWindowDelegateDidRunSelectorContextInfo]
// may be messaged in that spawned, non-main thread.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/canSpawnSeparateThread
func (p NSPrintOperation) CanSpawnSeparateThread() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canSpawnSeparateThread"))
	return rv
}
func (p NSPrintOperation) SetCanSpawnSeparateThread(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setCanSpawnSeparateThread:"), value)
}

// The current print operation for this thread.
//
// # Return Value
// 
// The print operation object, or `nil` if there is no current operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintOperation/current
func (_NSPrintOperationClass NSPrintOperationClass) CurrentOperation() NSPrintOperation {
	rv := objc.Send[objc.ID](objc.ID(_NSPrintOperationClass.class), objc.Sel("currentOperation"))
	return NSPrintOperationFromID(objc.ID(rv))
}
func (_NSPrintOperationClass NSPrintOperationClass) SetCurrentOperation(value NSPrintOperation) {
	objc.Send[struct{}](objc.ID(_NSPrintOperationClass.class), objc.Sel("setCurrentOperation:"), value)
}

