// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPrintInfo] class.
var (
	_NSPrintInfoClass     NSPrintInfoClass
	_NSPrintInfoClassOnce sync.Once
)

func getNSPrintInfoClass() NSPrintInfoClass {
	_NSPrintInfoClassOnce.Do(func() {
		_NSPrintInfoClass = NSPrintInfoClass{class: objc.GetClass("NSPrintInfo")}
	})
	return _NSPrintInfoClass
}

// GetNSPrintInfoClass returns the class object for NSPrintInfo.
func GetNSPrintInfoClass() NSPrintInfoClass {
	return getNSPrintInfoClass()
}

type NSPrintInfoClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPrintInfoClass) Alloc() NSPrintInfo {
	rv := objc.Send[NSPrintInfo](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that stores information that’s used to generate printed output.
//
// # Overview
// 
// A shared [NSPrintInfo] object is automatically created for an app and is
// used by default for all printing jobs for that app. The printing
// information in an [NSPrintInfo] object is stored in a dictionary. To access
// the standard attributes in the dictionary directly, this class defines a
// set of keys and provides the [NSPrintInfo.Dictionary] method. You can also initialize
// an instance of this class using the [NSPrintInfo.InitWithDictionary] method.
// 
// You can use this dictionary to store custom information associated with a
// print job. Any non-object values should be stored as [NSNumber] or
// [NSValue] objects in the dictionary. See [NSNumber] for a list of types
// which should be stored as numbers. For other non-object values, use the
// [NSValue] class.
// 
// To store custom information that belongs in printing presets you should use
// the dictionary returned by the [NSPrintInfo.PrintSettings] method.
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
//
// # Creating the Printing Information Object
//
//   - [NSPrintInfo.InitWithDictionary]: Returns a printing information object initialized with the parameters in the specified dictionary.
//   - [NSPrintInfo.InitWithCoder]: Creates a printing information object from data in an unarchiver.
//
// # Managing the Printing Rectangle
//
//   - [NSPrintInfo.PaperSize]: The size of the paper.
//   - [NSPrintInfo.SetPaperSize]
//   - [NSPrintInfo.TopMargin]: The top margin to the specified size.
//   - [NSPrintInfo.SetTopMargin]
//   - [NSPrintInfo.BottomMargin]: The height of the bottom margin.
//   - [NSPrintInfo.SetBottomMargin]
//   - [NSPrintInfo.LeftMargin]: The width of the left margin.
//   - [NSPrintInfo.SetLeftMargin]
//   - [NSPrintInfo.RightMargin]: The width of the right margin.
//   - [NSPrintInfo.SetRightMargin]
//   - [NSPrintInfo.ImageablePageBounds]: The imageable area of a sheet of paper specified by the print info.
//   - [NSPrintInfo.Orientation]: The orientation attribute.
//   - [NSPrintInfo.SetOrientation]
//   - [NSPrintInfo.PaperName]: The name of the currently selected paper size.
//   - [NSPrintInfo.SetPaperName]
//   - [NSPrintInfo.LocalizedPaperName]: The human-readable name of the currently selected paper size, suitable for presentation in user interfaces.
//
// # Pagination
//
//   - [NSPrintInfo.HorizontalPagination]: The horizontal pagination mode.
//   - [NSPrintInfo.SetHorizontalPagination]
//   - [NSPrintInfo.VerticalPagination]: The vertical pagination to the specified mode.
//   - [NSPrintInfo.SetVerticalPagination]
//
// # Positioning the Image on the Page
//
//   - [NSPrintInfo.HorizontallyCentered]: A Boolean value that indicates whether the image is centered horizontally.
//   - [NSPrintInfo.SetHorizontallyCentered]
//   - [NSPrintInfo.VerticallyCentered]: A Boolean value that indicates whether the image is centered vertically.
//   - [NSPrintInfo.SetVerticallyCentered]
//
// # Specifying the Printer
//
//   - [NSPrintInfo.Printer]: The printer object to be used for printing.
//   - [NSPrintInfo.SetPrinter]
//
// # Controlling Printing
//
//   - [NSPrintInfo.JobDisposition]: The action specified for the job.
//   - [NSPrintInfo.SetJobDisposition]
//   - [NSPrintInfo.SetUpPrintOperationDefaultValues]: Validates the attributes encapsulated by the print info.
//
// # Accessing the Print Info Dictionary
//
//   - [NSPrintInfo.Dictionary]: Returns the print info’s dictionary that contains the printing attributes.
//
// # Print Settings Convenience Methods
//
//   - [NSPrintInfo.SelectionOnly]: A Boolean value that indicates whether only the currently selected contents should be printed.
//   - [NSPrintInfo.SetSelectionOnly]
//   - [NSPrintInfo.ScalingFactor]: The current scaling factor.
//   - [NSPrintInfo.SetScalingFactor]
//
// # Accessing Core Printing Information
//
//   - [NSPrintInfo.PrintSettings]: A mutable dictionary containing the print settings from Core Printing.
//   - [NSPrintInfo.PMPrintSession]: Returns a Core Printing object configured with the print info’s session information.
//   - [NSPrintInfo.PMPageFormat]: Returns a Core Printing object configured with the print info’s page format information.
//   - [NSPrintInfo.PMPrintSettings]: Returns a Core Printing object configured with the print info’s print settings information
//   - [NSPrintInfo.UpdateFromPMPageFormat]: Synchronizes the print info’s page format information with information from its associated page format object.
//   - [NSPrintInfo.UpdateFromPMPrintSettings]: Synchronizes the print info’s print settings information with information from its associated print settings object.
//   - [NSPrintInfo.TakeSettingsFromPDFInfo]: Updates the print info with all the settings and attributes in the specified PDF info object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo
type NSPrintInfo struct {
	objectivec.Object
}

// NSPrintInfoFromID constructs a [NSPrintInfo] from an objc.ID.
//
// An object that stores information that’s used to generate printed output.
func NSPrintInfoFromID(id objc.ID) NSPrintInfo {
	return NSPrintInfo{objectivec.Object{ID: id}}
}
// NOTE: NSPrintInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPrintInfo] class.
//
// # Creating the Printing Information Object
//
//   - [INSPrintInfo.InitWithDictionary]: Returns a printing information object initialized with the parameters in the specified dictionary.
//   - [INSPrintInfo.InitWithCoder]: Creates a printing information object from data in an unarchiver.
//
// # Managing the Printing Rectangle
//
//   - [INSPrintInfo.PaperSize]: The size of the paper.
//   - [INSPrintInfo.SetPaperSize]
//   - [INSPrintInfo.TopMargin]: The top margin to the specified size.
//   - [INSPrintInfo.SetTopMargin]
//   - [INSPrintInfo.BottomMargin]: The height of the bottom margin.
//   - [INSPrintInfo.SetBottomMargin]
//   - [INSPrintInfo.LeftMargin]: The width of the left margin.
//   - [INSPrintInfo.SetLeftMargin]
//   - [INSPrintInfo.RightMargin]: The width of the right margin.
//   - [INSPrintInfo.SetRightMargin]
//   - [INSPrintInfo.ImageablePageBounds]: The imageable area of a sheet of paper specified by the print info.
//   - [INSPrintInfo.Orientation]: The orientation attribute.
//   - [INSPrintInfo.SetOrientation]
//   - [INSPrintInfo.PaperName]: The name of the currently selected paper size.
//   - [INSPrintInfo.SetPaperName]
//   - [INSPrintInfo.LocalizedPaperName]: The human-readable name of the currently selected paper size, suitable for presentation in user interfaces.
//
// # Pagination
//
//   - [INSPrintInfo.HorizontalPagination]: The horizontal pagination mode.
//   - [INSPrintInfo.SetHorizontalPagination]
//   - [INSPrintInfo.VerticalPagination]: The vertical pagination to the specified mode.
//   - [INSPrintInfo.SetVerticalPagination]
//
// # Positioning the Image on the Page
//
//   - [INSPrintInfo.HorizontallyCentered]: A Boolean value that indicates whether the image is centered horizontally.
//   - [INSPrintInfo.SetHorizontallyCentered]
//   - [INSPrintInfo.VerticallyCentered]: A Boolean value that indicates whether the image is centered vertically.
//   - [INSPrintInfo.SetVerticallyCentered]
//
// # Specifying the Printer
//
//   - [INSPrintInfo.Printer]: The printer object to be used for printing.
//   - [INSPrintInfo.SetPrinter]
//
// # Controlling Printing
//
//   - [INSPrintInfo.JobDisposition]: The action specified for the job.
//   - [INSPrintInfo.SetJobDisposition]
//   - [INSPrintInfo.SetUpPrintOperationDefaultValues]: Validates the attributes encapsulated by the print info.
//
// # Accessing the Print Info Dictionary
//
//   - [INSPrintInfo.Dictionary]: Returns the print info’s dictionary that contains the printing attributes.
//
// # Print Settings Convenience Methods
//
//   - [INSPrintInfo.SelectionOnly]: A Boolean value that indicates whether only the currently selected contents should be printed.
//   - [INSPrintInfo.SetSelectionOnly]
//   - [INSPrintInfo.ScalingFactor]: The current scaling factor.
//   - [INSPrintInfo.SetScalingFactor]
//
// # Accessing Core Printing Information
//
//   - [INSPrintInfo.PrintSettings]: A mutable dictionary containing the print settings from Core Printing.
//   - [INSPrintInfo.PMPrintSession]: Returns a Core Printing object configured with the print info’s session information.
//   - [INSPrintInfo.PMPageFormat]: Returns a Core Printing object configured with the print info’s page format information.
//   - [INSPrintInfo.PMPrintSettings]: Returns a Core Printing object configured with the print info’s print settings information
//   - [INSPrintInfo.UpdateFromPMPageFormat]: Synchronizes the print info’s page format information with information from its associated page format object.
//   - [INSPrintInfo.UpdateFromPMPrintSettings]: Synchronizes the print info’s print settings information with information from its associated print settings object.
//   - [INSPrintInfo.TakeSettingsFromPDFInfo]: Updates the print info with all the settings and attributes in the specified PDF info object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo
type INSPrintInfo interface {
	objectivec.IObject

	// Topic: Creating the Printing Information Object

	// Returns a printing information object initialized with the parameters in the specified dictionary.
	InitWithDictionary(attributes foundation.INSDictionary) NSPrintInfo
	// Creates a printing information object from data in an unarchiver.
	InitWithCoder(coder foundation.INSCoder) NSPrintInfo

	// Topic: Managing the Printing Rectangle

	// The size of the paper.
	PaperSize() corefoundation.CGSize
	SetPaperSize(value corefoundation.CGSize)
	// The top margin to the specified size.
	TopMargin() float64
	SetTopMargin(value float64)
	// The height of the bottom margin.
	BottomMargin() float64
	SetBottomMargin(value float64)
	// The width of the left margin.
	LeftMargin() float64
	SetLeftMargin(value float64)
	// The width of the right margin.
	RightMargin() float64
	SetRightMargin(value float64)
	// The imageable area of a sheet of paper specified by the print info.
	ImageablePageBounds() corefoundation.CGRect
	// The orientation attribute.
	Orientation() NSPaperOrientation
	SetOrientation(value NSPaperOrientation)
	// The name of the currently selected paper size.
	PaperName() NSPrinterPaperName
	SetPaperName(value NSPrinterPaperName)
	// The human-readable name of the currently selected paper size, suitable for presentation in user interfaces.
	LocalizedPaperName() string

	// Topic: Pagination

	// The horizontal pagination mode.
	HorizontalPagination() NSPrintingPaginationMode
	SetHorizontalPagination(value NSPrintingPaginationMode)
	// The vertical pagination to the specified mode.
	VerticalPagination() NSPrintingPaginationMode
	SetVerticalPagination(value NSPrintingPaginationMode)

	// Topic: Positioning the Image on the Page

	// A Boolean value that indicates whether the image is centered horizontally.
	HorizontallyCentered() bool
	SetHorizontallyCentered(value bool)
	// A Boolean value that indicates whether the image is centered vertically.
	VerticallyCentered() bool
	SetVerticallyCentered(value bool)

	// Topic: Specifying the Printer

	// The printer object to be used for printing.
	Printer() INSPrinter
	SetPrinter(value INSPrinter)

	// Topic: Controlling Printing

	// The action specified for the job.
	JobDisposition() NSPrintJobDispositionValue
	SetJobDisposition(value NSPrintJobDispositionValue)
	// Validates the attributes encapsulated by the print info.
	SetUpPrintOperationDefaultValues()

	// Topic: Accessing the Print Info Dictionary

	// Returns the print info’s dictionary that contains the printing attributes.
	Dictionary() foundation.INSDictionary

	// Topic: Print Settings Convenience Methods

	// A Boolean value that indicates whether only the currently selected contents should be printed.
	SelectionOnly() bool
	SetSelectionOnly(value bool)
	// The current scaling factor.
	ScalingFactor() float64
	SetScalingFactor(value float64)

	// Topic: Accessing Core Printing Information

	// A mutable dictionary containing the print settings from Core Printing.
	PrintSettings() foundation.INSDictionary
	// Returns a Core Printing object configured with the print info’s session information.
	PMPrintSession() unsafe.Pointer
	// Returns a Core Printing object configured with the print info’s page format information.
	PMPageFormat() unsafe.Pointer
	// Returns a Core Printing object configured with the print info’s print settings information
	PMPrintSettings() unsafe.Pointer
	// Synchronizes the print info’s page format information with information from its associated page format object.
	UpdateFromPMPageFormat()
	// Synchronizes the print info’s print settings information with information from its associated print settings object.
	UpdateFromPMPrintSettings()
	// Updates the print info with all the settings and attributes in the specified PDF info object.
	TakeSettingsFromPDFInfo(inPDFInfo INSPDFInfo)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p NSPrintInfo) Init() NSPrintInfo {
	rv := objc.Send[NSPrintInfo](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPrintInfo) Autorelease() NSPrintInfo {
	rv := objc.Send[NSPrintInfo](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPrintInfo creates a new NSPrintInfo instance.
func NewNSPrintInfo() NSPrintInfo {
	class := getNSPrintInfoClass()
	rv := objc.Send[NSPrintInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a printing information object from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/init(coder:)
func NewPrintInfoWithCoder(coder foundation.INSCoder) NSPrintInfo {
	instance := getNSPrintInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPrintInfoFromID(rv)
}

// Returns a printing information object initialized with the parameters in
// the specified dictionary.
//
// attributes: The possible key-value pairs contained in `aDictionary` are described in
// Constants.
//
// # Return Value
// 
// An initialized [NSPrintInfo] object, or nil if the object could not be
// created.
//
// # Discussion
// 
// This method is the designated initializer for this class. Non-object values
// should be stored in [NSValue] objects (or an appropriate subclass like
// [NSNumber]) in the dictionary. See [NSNumber] for a list of types which
// should be stored using the [NSNumber] class; otherwise use [NSValue].
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/init(dictionary:)
func NewPrintInfoWithDictionary(attributes foundation.INSDictionary) NSPrintInfo {
	instance := getNSPrintInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionary:"), attributes)
	return NSPrintInfoFromID(rv)
}

// Returns a printing information object initialized with the parameters in
// the specified dictionary.
//
// attributes: The possible key-value pairs contained in `aDictionary` are described in
// Constants.
//
// # Return Value
// 
// An initialized [NSPrintInfo] object, or nil if the object could not be
// created.
//
// # Discussion
// 
// This method is the designated initializer for this class. Non-object values
// should be stored in [NSValue] objects (or an appropriate subclass like
// [NSNumber]) in the dictionary. See [NSNumber] for a list of types which
// should be stored using the [NSNumber] class; otherwise use [NSValue].
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/init(dictionary:)
func (p NSPrintInfo) InitWithDictionary(attributes foundation.INSDictionary) NSPrintInfo {
	rv := objc.Send[NSPrintInfo](p.ID, objc.Sel("initWithDictionary:"), attributes)
	return rv
}
// Creates a printing information object from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/init(coder:)
func (p NSPrintInfo) InitWithCoder(coder foundation.INSCoder) NSPrintInfo {
	rv := objc.Send[NSPrintInfo](p.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Validates the attributes encapsulated by the print info.
//
// # Discussion
// 
// Invoked when the print operation is about to start. Subclasses may override
// this method to set default values for any attributes that are not set.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/setUpPrintOperationDefaultValues()
func (p NSPrintInfo) SetUpPrintOperationDefaultValues() {
	objc.Send[objc.ID](p.ID, objc.Sel("setUpPrintOperationDefaultValues"))
}
// Returns the print info’s dictionary that contains the printing
// attributes.
//
// # Discussion
// 
// The key-value pairs contained in the dictionary are described in Constants.
// Modifying the returned dictionary changes the receiver’s attributes.
// 
// This dictionary is key-value observing compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/dictionary()
func (p NSPrintInfo) Dictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dictionary"))
	return foundation.NSDictionaryFromID(rv)
}
// Returns a Core Printing object configured with the print info’s session
// information.
//
// # Return Value
// 
// A pointer to a [PMPrintSession] object, an opaque type that stores
// information about a print job. You should not call [PMRelease] to release
// the returned object, except to balance calls to [PMRetain] that your code
// also issued.
//
// [PMPrintSession]: https://developer.apple.com/documentation/applicationservices/pmprintsession
//
// # Discussion
// 
// The information in the returned [PMPrintSession] object is consistent with
// the receiver’s session information at the time this method is called.
// Subsequent changes to the receiving [NSPrintInfo] object do not result in
// changes to the information in the [PMPrintSession] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/pmPrintSession()
func (p NSPrintInfo) PMPrintSession() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p.ID, objc.Sel("PMPrintSession"))
	return rv
}
// Returns a Core Printing object configured with the print info’s page
// format information.
//
// # Return Value
// 
// A pointer to a [PMPageFormat] object, an opaque data type that stores
// information such as the paper size, orientation, and scale of pages in a
// printing session. You should not call [PMRelease] to release the returned
// object, except to balance calls to [PMRetain] that your code also issued.
//
// [PMPageFormat]: https://developer.apple.com/documentation/applicationservices/pmpageformat
//
// # Discussion
// 
// The information in the returned [PMPageFormat] object is consistent with
// the receiver’s page format information at the time this method is called.
// Subsequent changes to the receiving [NSPrintInfo] object do not result in
// changes to the information in the [PMPageFormat] object.
// 
// If you make changes to the data in the [PMPageFormat] object, you should
// invoke the [UpdateFromPMPageFormat] method to synchronize those changes
// with the [NSPrintInfo] object that created the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/pmPageFormat()
func (p NSPrintInfo) PMPageFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p.ID, objc.Sel("PMPageFormat"))
	return rv
}
// Returns a Core Printing object configured with the print info’s print
// settings information
//
// # Return Value
// 
// A pointer to a [PMPrintSettings] object, an opaque data type used to store
// information such as the number of copies and the range of pages in a
// printing session. You should not call [PMRelease] to release the returned
// object, except to balance calls to [PMRetain] that your code also issued.
//
// [PMPrintSettings]: https://developer.apple.com/documentation/applicationservices/pmprintsettings
//
// # Discussion
// 
// The information in the returned [PMPrintSettings] object is consistent with
// the receiver’s print settings at the time this method is called.
// Subsequent changes to the receiving [NSPrintInfo] object do not result in
// changes to the information in the [PMPrintSettings] data type.
// 
// If you make changes to the data in the [PMPrintSettings] object, you should
// invoke the [UpdateFromPMPrintSettings] method to synchronize those changes
// with the [NSPrintInfo] object that created the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/pmPrintSettings()
func (p NSPrintInfo) PMPrintSettings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p.ID, objc.Sel("PMPrintSettings"))
	return rv
}
// Synchronizes the print info’s page format information with information
// from its associated page format object.
//
// # Discussion
// 
// You should use this method after making changes to the [PMPageFormat]
// object obtained from the receiver. Each [NSPrintInfo] object keeps track of
// the object returned from its [PMPageFormat] method and obtains any updated
// information from the object directly. You only need to synchronize the
// objects once when you have made all of the desired changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/updateFromPMPageFormat()
func (p NSPrintInfo) UpdateFromPMPageFormat() {
	objc.Send[objc.ID](p.ID, objc.Sel("updateFromPMPageFormat"))
}
// Synchronizes the print info’s print settings information with information
// from its associated print settings object.
//
// # Discussion
// 
// You should use this method after making changes to the [PMPrintSettings]
// object obtained from the receiver. Each [NSPrintInfo] object keeps track of
// the object returned from its [PMPrintSettings] method and obtains any
// updated information from the object directly. You only need to synchronize
// the objects once when you have made all of the desired changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/updateFromPMPrintSettings()
func (p NSPrintInfo) UpdateFromPMPrintSettings() {
	objc.Send[objc.ID](p.ID, objc.Sel("updateFromPMPrintSettings"))
}
// Updates the print info with all the settings and attributes in the
// specified PDF info object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/takeSettings(from:)
func (p NSPrintInfo) TakeSettingsFromPDFInfo(inPDFInfo INSPDFInfo) {
	objc.Send[objc.ID](p.ID, objc.Sel("takeSettingsFromPDFInfo:"), inPDFInfo)
}
func (p NSPrintInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The size of the paper.
//
// # Discussion
// 
// The size is measured in points in the user coordinate space.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/paperSize
func (p NSPrintInfo) PaperSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("paperSize"))
	return corefoundation.CGSize(rv)
}
func (p NSPrintInfo) SetPaperSize(value corefoundation.CGSize) {
	objc.Send[struct{}](p.ID, objc.Sel("setPaperSize:"), value)
}
// The top margin to the specified size.
//
// # Discussion
// 
// Size is measured in points in the user coordinate space.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/topMargin
func (p NSPrintInfo) TopMargin() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("topMargin"))
	return rv
}
func (p NSPrintInfo) SetTopMargin(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setTopMargin:"), value)
}
// The height of the bottom margin.
//
// # Discussion
// 
// bottomMargin is measured in points in the user coordinate space.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/bottomMargin
func (p NSPrintInfo) BottomMargin() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("bottomMargin"))
	return rv
}
func (p NSPrintInfo) SetBottomMargin(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setBottomMargin:"), value)
}
// The width of the left margin.
//
// # Discussion
// 
// This value is measured in points in the user coordinate space.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/leftMargin
func (p NSPrintInfo) LeftMargin() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("leftMargin"))
	return rv
}
func (p NSPrintInfo) SetLeftMargin(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setLeftMargin:"), value)
}
// The width of the right margin.
//
// # Discussion
// 
// Size is measured in points in the user coordinate space.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/rightMargin
func (p NSPrintInfo) RightMargin() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("rightMargin"))
	return rv
}
func (p NSPrintInfo) SetRightMargin(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setRightMargin:"), value)
}
// The imageable area of a sheet of paper specified by the print info.
//
// # Discussion
// 
// This property takes into account the current printer, paper size, and
// orientation settings, but not scaling factors. “Imageable area” is the
// maximum area that can possibly be marked on by the printer hardware, not
// the area defined by the current margin settings.
// 
// The origin (0, 0) of the rectangle is in the lower-left corner of the
// oriented sheet. The imageable bounds may extend past the edges of the sheet
// when, for example, a printer driver specifies it so that borderless
// printing can be done reliably.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/imageablePageBounds
func (p NSPrintInfo) ImageablePageBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("imageablePageBounds"))
	return corefoundation.CGRect(rv)
}
// The orientation attribute.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/orientation-swift.property
func (p NSPrintInfo) Orientation() NSPaperOrientation {
	rv := objc.Send[NSPaperOrientation](p.ID, objc.Sel("orientation"))
	return NSPaperOrientation(rv)
}
func (p NSPrintInfo) SetOrientation(value NSPaperOrientation) {
	objc.Send[struct{}](p.ID, objc.Sel("setOrientation:"), value)
}
// The name of the currently selected paper size.
//
// # Discussion
// 
// The string contains a value such as Letter or Legal. Paper names are
// implementation specific.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/paperName
func (p NSPrintInfo) PaperName() NSPrinterPaperName {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("paperName"))
	return NSPrinterPaperName(foundation.NSStringFromID(rv).String())
}
func (p NSPrintInfo) SetPaperName(value NSPrinterPaperName) {
	objc.Send[struct{}](p.ID, objc.Sel("setPaperName:"), objc.String(string(value)))
}
// The human-readable name of the currently selected paper size, suitable for
// presentation in user interfaces.
//
// # Discussion
// 
// This is typically different from the value of [PaperName], which is almost
// never suitable for presentation to the user.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/localizedPaperName
func (p NSPrintInfo) LocalizedPaperName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("localizedPaperName"))
	return foundation.NSStringFromID(rv).String()
}
// The horizontal pagination mode.
//
// # Discussion
// 
// Pagination modes described in [NSPrintInfo.PaginationMode].
//
// [NSPrintInfo.PaginationMode]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaginationMode
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/horizontalPagination
func (p NSPrintInfo) HorizontalPagination() NSPrintingPaginationMode {
	rv := objc.Send[NSPrintingPaginationMode](p.ID, objc.Sel("horizontalPagination"))
	return NSPrintingPaginationMode(rv)
}
func (p NSPrintInfo) SetHorizontalPagination(value NSPrintingPaginationMode) {
	objc.Send[struct{}](p.ID, objc.Sel("setHorizontalPagination:"), value)
}
// The vertical pagination to the specified mode.
//
// # Discussion
// 
// The pagination modes are described in [NSPrintInfo.PaginationMode]
//
// [NSPrintInfo.PaginationMode]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaginationMode
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/verticalPagination
func (p NSPrintInfo) VerticalPagination() NSPrintingPaginationMode {
	rv := objc.Send[NSPrintingPaginationMode](p.ID, objc.Sel("verticalPagination"))
	return NSPrintingPaginationMode(rv)
}
func (p NSPrintInfo) SetVerticalPagination(value NSPrintingPaginationMode) {
	objc.Send[struct{}](p.ID, objc.Sel("setVerticalPagination:"), value)
}
// A Boolean value that indicates whether the image is centered horizontally.
//
// # Discussion
// 
// [true] if the image is centered horizontally; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isHorizontallyCentered
func (p NSPrintInfo) HorizontallyCentered() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isHorizontallyCentered"))
	return rv
}
func (p NSPrintInfo) SetHorizontallyCentered(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setHorizontallyCentered:"), value)
}
// A Boolean value that indicates whether the image is centered vertically.
//
// # Discussion
// 
// [true] if the image is centered vertically; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isVerticallyCentered
func (p NSPrintInfo) VerticallyCentered() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isVerticallyCentered"))
	return rv
}
func (p NSPrintInfo) SetVerticallyCentered(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setVerticallyCentered:"), value)
}
// The printer object to be used for printing.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/printer
func (p NSPrintInfo) Printer() INSPrinter {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("printer"))
	return NSPrinterFromID(objc.ID(rv))
}
func (p NSPrintInfo) SetPrinter(value INSPrinter) {
	objc.Send[struct{}](p.ID, objc.Sel("setPrinter:"), value)
}
// The action specified for the job.
//
// # Discussion
// 
// One of the following value:
// 
// - [spool] is a normal print job. - [preview] sends the print job to the
// Preview application. - [save] saves the print job to a file. - [cancel]
// aborts the print job.
//
// [cancel]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct/cancel
// [preview]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct/preview
// [save]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct/save
// [spool]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct/spool
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/jobDisposition-swift.property
func (p NSPrintInfo) JobDisposition() NSPrintJobDispositionValue {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("jobDisposition"))
	return NSPrintJobDispositionValue(foundation.NSStringFromID(rv).String())
}
func (p NSPrintInfo) SetJobDisposition(value NSPrintJobDispositionValue) {
	objc.Send[struct{}](p.ID, objc.Sel("setJobDisposition:"), objc.String(string(value)))
}
// A Boolean value that indicates whether only the currently selected contents
// should be printed.
//
// # Discussion
// 
// [true] if only the currently selected contents should be printed, otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isSelectionOnly
func (p NSPrintInfo) SelectionOnly() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isSelectionOnly"))
	return rv
}
func (p NSPrintInfo) SetSelectionOnly(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setSelectionOnly:"), value)
}
// The current scaling factor.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/scalingFactor
func (p NSPrintInfo) ScalingFactor() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("scalingFactor"))
	return rv
}
func (p NSPrintInfo) SetScalingFactor(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setScalingFactor:"), value)
}
// A mutable dictionary containing the print settings from Core Printing.
//
// # Discussion
// 
// You can use this property to get and set values from the system print
// settings. The keys in the dictionary represent the values returned by the
// Core Printing function [PMPrintSettingsGetValue(_:_:_:)]. They correspond
// to the settings currently in the print panel and include everything from
// custom values set by your accessory panels to values provided by the
// printer driver’s print dialog extension.
// 
// Adding keys to the dictionary is equivalent to calling the Core Printing
// function [PMPrintSettingsSetValue(_:_:_:_:)]. Your new keys are added to
// the current print settings and are saved with any user preset files
// generated by the macOS printing system. Because the print settings are
// stored in a property list, any values you add to the dictionary must
// correspond to scalar types such as strings, numbers, dates, booleans, and
// data objects or collection types such as dictionaries and arrays.
// 
// Other parts of the printing system use key strings like
// `com.AppleXCUIElementTypePrint().PrintSettings.PMColorSyncProfileID` to
// identify print settings. Cocoa replaces the periods in such strings with
// underscores. Thus, the preceding key string would be
// `com_apple_print_PrintSettings_PMColorSyncProfileID` instead. If you use
// reverse-DNS style key strings for your custom attributes, you should follow
// the same convention of using underscore characters instead of periods.
//
// [PMPrintSettingsGetValue(_:_:_:)]: https://developer.apple.com/documentation/applicationservices/1460602-pmprintsettingsgetvalue
// [PMPrintSettingsSetValue(_:_:_:_:)]: https://developer.apple.com/documentation/applicationservices/1461697-pmprintsettingssetvalue
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/printSettings
func (p NSPrintInfo) PrintSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("printSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The shared printing information object.
//
// # Return Value
// 
// The shared printer information.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/shared
func (_NSPrintInfoClass NSPrintInfoClass) SharedPrintInfo() NSPrintInfo {
	rv := objc.Send[objc.ID](objc.ID(_NSPrintInfoClass.class), objc.Sel("sharedPrintInfo"))
	return NSPrintInfoFromID(objc.ID(rv))
}
func (_NSPrintInfoClass NSPrintInfoClass) SetSharedPrintInfo(value NSPrintInfo) {
	objc.Send[struct{}](objc.ID(_NSPrintInfoClass.class), objc.Sel("setSharedPrintInfo:"), value)
}
// Deprecated.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/defaultPrinter
func (_NSPrintInfoClass NSPrintInfoClass) DefaultPrinter() NSPrinter {
	rv := objc.Send[objc.ID](objc.ID(_NSPrintInfoClass.class), objc.Sel("defaultPrinter"))
	return NSPrinterFromID(objc.ID(rv))
}

