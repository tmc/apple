// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPrinter] class.
var (
	_NSPrinterClass     NSPrinterClass
	_NSPrinterClassOnce sync.Once
)

func getNSPrinterClass() NSPrinterClass {
	_NSPrinterClassOnce.Do(func() {
		_NSPrinterClass = NSPrinterClass{class: objc.GetClass("NSPrinter")}
	})
	return _NSPrinterClass
}

// GetNSPrinterClass returns the class object for NSPrinter.
func GetNSPrinterClass() NSPrinterClass {
	return getNSPrinterClass()
}

type NSPrinterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPrinterClass) Alloc() NSPrinter {
	rv := objc.Send[NSPrinter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a printer’s capabilities.
//
// # Overview
// 
// [NSPrinter] provides information about a printer; it does not modify
// printer attributes or control a printing job. A printer object can be
// constructed by specifying either the printer name or the make and model of
// an available printer. Typically, Cocoa apps don’t create [NSPrinter]
// objects; instead, the printing system uses these objects to support the
// printing jobs and when it shows users a list of printers.
//
// # Getting Attributes
//
//   - [NSPrinter.Name]: The printer’s name.
//   - [NSPrinter.Type]: A description of the printer’s make and model.
//
// # Getting Page and Printer Information
//
//   - [NSPrinter.PageSizeForPaper]: Returns the size of the page for the specified paper type.
//   - [NSPrinter.LanguageLevel]: The PostScript language level recognized by the printer.
//
// # Querying Tables
//
//   - [NSPrinter.DeviceDescription]: A dictionary of keys and values that describe the device.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter
type NSPrinter struct {
	objectivec.Object
}

// NSPrinterFromID constructs a [NSPrinter] from an objc.ID.
//
// An object that describes a printer’s capabilities.
func NSPrinterFromID(id objc.ID) NSPrinter {
	return NSPrinter{objectivec.Object{ID: id}}
}
// NOTE: NSPrinter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPrinter] class.
//
// # Getting Attributes
//
//   - [INSPrinter.Name]: The printer’s name.
//   - [INSPrinter.Type]: A description of the printer’s make and model.
//
// # Getting Page and Printer Information
//
//   - [INSPrinter.PageSizeForPaper]: Returns the size of the page for the specified paper type.
//   - [INSPrinter.LanguageLevel]: The PostScript language level recognized by the printer.
//
// # Querying Tables
//
//   - [INSPrinter.DeviceDescription]: A dictionary of keys and values that describe the device.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter
type INSPrinter interface {
	objectivec.IObject

	// Topic: Getting Attributes

	// The printer’s name.
	Name() string
	// A description of the printer’s make and model.
	Type() NSPrinterTypeName

	// Topic: Getting Page and Printer Information

	// Returns the size of the page for the specified paper type.
	PageSizeForPaper(paperName NSPrinterPaperName) corefoundation.CGSize
	// The PostScript language level recognized by the printer.
	LanguageLevel() int

	// Topic: Querying Tables

	// A dictionary of keys and values that describe the device.
	DeviceDescription() foundation.INSDictionary

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p NSPrinter) Init() NSPrinter {
	rv := objc.Send[NSPrinter](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPrinter) Autorelease() NSPrinter {
	rv := objc.Send[NSPrinter](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPrinter creates a new NSPrinter instance.
func NewNSPrinter() NSPrinter {
	class := getNSPrinterClass()
	rv := objc.Send[NSPrinter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns a printer object initialized with the specified printer
// name.
//
// name: The name of the printer.
//
// # Return Value
// 
// An initialized [NSPrinter] object, or `nil` if the specified printer was
// not available.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter/init(name:)
func NewPrinterWithName(name string) NSPrinter {
	rv := objc.Send[objc.ID](objc.ID(getNSPrinterClass().class), objc.Sel("printerWithName:"), objc.String(name))
	return NSPrinterFromID(rv)
}

// Creates and returns a printer object initialized to the first available
// printer with the specified make and model information.
//
// type: A string describing the make and model information. You can get this string
// using the [PrinterTypes] method.
//
// # Return Value
// 
// An initialized [NSPrinter] object, or `nil` if the specified printer was
// not available.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter/init(type:)
func NewPrinterWithType(type_ NSPrinterTypeName) NSPrinter {
	rv := objc.Send[objc.ID](objc.ID(getNSPrinterClass().class), objc.Sel("printerWithType:"), objc.String(string(type_)))
	return NSPrinterFromID(rv)
}

// Returns the size of the page for the specified paper type.
//
// paperName: Possible values are printer-dependent and are contained in the printer’s
// PPD file. Typical values are “Letter” and “Legal”.
//
// # Return Value
// 
// The size of the page, measured in points in the user coordinate space. The
// returned size is zero if the specified paper name is not recognized or its
// entry in the PPD file cannot be parsed.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter/pageSize(forPaper:)
func (p NSPrinter) PageSizeForPaper(paperName NSPrinterPaperName) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("pageSizeForPaper:"), objc.String(string(paperName)))
	return corefoundation.CGSize(rv)
}
func (p NSPrinter) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The printer’s name.
//
// # Return Value
// 
// The printer name.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter/name
func (p NSPrinter) Name() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// A description of the printer’s make and model.
//
// # Return Value
// 
// A description of the printer’s make and model.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter/type
func (p NSPrinter) Type() NSPrinterTypeName {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("type"))
	return NSPrinterTypeName(foundation.NSStringFromID(rv).String())
}

// The PostScript language level recognized by the printer.
//
// # Return Value
// 
// The PostScript language level. The value is 0 if the receiver is not a
// PostScript printer.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter/languageLevel
func (p NSPrinter) LanguageLevel() int {
	rv := objc.Send[int](p.ID, objc.Sel("languageLevel"))
	return rv
}

// A dictionary of keys and values that describe the device.
//
// # Return Value
// 
// A dictionary of the device properties. See `NSGraphics.H()` for possible
// keys. The only key guaranteed to exist is [NSDeviceIsPrinter].
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter/deviceDescription
func (p NSPrinter) DeviceDescription() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("deviceDescription"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// Returns the names of all available printers.
//
// # Return Value
// 
// An array of [NSString] objects, each of which contains the name of an
// available printer.
// 
// # Discussion
// 
// The user constructs the list of available printers when adding a printer in
// the Print panel or setting up printers in the Print & Scan preferences
// pane.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter/printerNames
func (_NSPrinterClass NSPrinterClass) PrinterNames() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSPrinterClass.class), objc.Sel("printerNames"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns descriptions of the makes and models of all available printers.
//
// # Return Value
// 
// An array of [NSString] objects, each of which contains the make and model
// information for a supported printer.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter/printerTypes
func (_NSPrinterClass NSPrinterClass) PrinterTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSPrinterClass.class), objc.Sel("printerTypes"))
	return objc.ConvertSliceToStrings(rv)
}

