// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPDFInfo] class.
var (
	_NSPDFInfoClass     NSPDFInfoClass
	_NSPDFInfoClassOnce sync.Once
)

func getNSPDFInfoClass() NSPDFInfoClass {
	_NSPDFInfoClassOnce.Do(func() {
		_NSPDFInfoClass = NSPDFInfoClass{class: objc.GetClass("NSPDFInfo")}
	})
	return _NSPDFInfoClass
}

// GetNSPDFInfoClass returns the class object for NSPDFInfo.
func GetNSPDFInfoClass() NSPDFInfoClass {
	return getNSPDFInfoClass()
}

type NSPDFInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPDFInfoClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPDFInfoClass) Alloc() NSPDFInfo {
	rv := objc.Send[NSPDFInfo](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that stores information associated with the creation of a PDF
// file, such as its URL, tag names, page orientation, and paper size.
//
// # Overview
//
// Typically, a PDF panel—that is, a panel created by an [NSPDFPanel]
// object—displays the information supplied by an [NSPDFInfo] object when
// the user wants to export content as a PDF file. A PDF panel can also update
// a PDF info object with information it receives from the user.
//
// # Specifying PDF Information
//
//   - [NSPDFInfo.URL]: The URL identifying the location at which the PDF file will be created.
//   - [NSPDFInfo.SetURL]
//   - [NSPDFInfo.FileExtensionHidden]: A Boolean value that indicates whether the file extension should appear after the filename.
//   - [NSPDFInfo.SetFileExtensionHidden]
//   - [NSPDFInfo.TagNames]: An array of tag names that should be applied to the PDF file after it’s created.
//   - [NSPDFInfo.SetTagNames]
//   - [NSPDFInfo.Orientation]: The paper orientation to use when exporting content as a PDF file.
//   - [NSPDFInfo.SetOrientation]
//   - [NSPDFInfo.PaperSize]: The paper size to use when exporting content as a PDF file.
//   - [NSPDFInfo.SetPaperSize]
//   - [NSPDFInfo.Attributes]: A dictionary of additional attributes that describe how to export content as a PDF file.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFInfo
type NSPDFInfo struct {
	objectivec.Object
}

// NSPDFInfoFromID constructs a [NSPDFInfo] from an objc.ID.
//
// An object that stores information associated with the creation of a PDF
// file, such as its URL, tag names, page orientation, and paper size.
func NSPDFInfoFromID(id objc.ID) NSPDFInfo {
	return NSPDFInfo{objectivec.Object{ID: id}}
}

// NOTE: NSPDFInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPDFInfo] class.
//
// # Specifying PDF Information
//
//   - [INSPDFInfo.URL]: The URL identifying the location at which the PDF file will be created.
//   - [INSPDFInfo.SetURL]
//   - [INSPDFInfo.FileExtensionHidden]: A Boolean value that indicates whether the file extension should appear after the filename.
//   - [INSPDFInfo.SetFileExtensionHidden]
//   - [INSPDFInfo.TagNames]: An array of tag names that should be applied to the PDF file after it’s created.
//   - [INSPDFInfo.SetTagNames]
//   - [INSPDFInfo.Orientation]: The paper orientation to use when exporting content as a PDF file.
//   - [INSPDFInfo.SetOrientation]
//   - [INSPDFInfo.PaperSize]: The paper size to use when exporting content as a PDF file.
//   - [INSPDFInfo.SetPaperSize]
//   - [INSPDFInfo.Attributes]: A dictionary of additional attributes that describe how to export content as a PDF file.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFInfo
type INSPDFInfo interface {
	objectivec.IObject

	// Topic: Specifying PDF Information

	// The URL identifying the location at which the PDF file will be created.
	URL() foundation.INSURL
	SetURL(value foundation.INSURL)
	// A Boolean value that indicates whether the file extension should appear after the filename.
	FileExtensionHidden() bool
	SetFileExtensionHidden(value bool)
	// An array of tag names that should be applied to the PDF file after it’s created.
	TagNames() []string
	SetTagNames(value []string)
	// The paper orientation to use when exporting content as a PDF file.
	Orientation() NSPaperOrientation
	SetOrientation(value NSPaperOrientation)
	// The paper size to use when exporting content as a PDF file.
	PaperSize() corefoundation.CGSize
	SetPaperSize(value corefoundation.CGSize)
	// A dictionary of additional attributes that describe how to export content as a PDF file.
	Attributes() foundation.INSDictionary

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p NSPDFInfo) Init() NSPDFInfo {
	rv := objc.Send[NSPDFInfo](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPDFInfo) Autorelease() NSPDFInfo {
	rv := objc.Send[NSPDFInfo](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPDFInfo creates a new NSPDFInfo instance.
func NewNSPDFInfo() NSPDFInfo {
	class := getNSPDFInfoClass()
	rv := objc.Send[NSPDFInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p NSPDFInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The URL identifying the location at which the PDF file will be created.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFInfo/url
func (p NSPDFInfo) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (p NSPDFInfo) SetURL(value foundation.INSURL) {
	objc.Send[struct{}](p.ID, objc.Sel("setURL:"), value)
}

// A Boolean value that indicates whether the file extension should appear
// after the filename.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFInfo/isFileExtensionHidden
func (p NSPDFInfo) FileExtensionHidden() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isFileExtensionHidden"))
	return rv
}
func (p NSPDFInfo) SetFileExtensionHidden(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setFileExtensionHidden:"), value)
}

// An array of tag names that should be applied to the PDF file after it’s
// created.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFInfo/tagNames
func (p NSPDFInfo) TagNames() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("tagNames"))
	return objc.ConvertSliceToStrings(rv)
}
func (p NSPDFInfo) SetTagNames(value []string) {
	objc.Send[struct{}](p.ID, objc.Sel("setTagNames:"), objectivec.StringSliceToNSArray(value))
}

// The paper orientation to use when exporting content as a PDF file.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFInfo/orientation
func (p NSPDFInfo) Orientation() NSPaperOrientation {
	rv := objc.Send[NSPaperOrientation](p.ID, objc.Sel("orientation"))
	return NSPaperOrientation(rv)
}
func (p NSPDFInfo) SetOrientation(value NSPaperOrientation) {
	objc.Send[struct{}](p.ID, objc.Sel("setOrientation:"), value)
}

// The paper size to use when exporting content as a PDF file.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFInfo/paperSize
func (p NSPDFInfo) PaperSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("paperSize"))
	return corefoundation.CGSize(rv)
}
func (p NSPDFInfo) SetPaperSize(value corefoundation.CGSize) {
	objc.Send[struct{}](p.ID, objc.Sel("setPaperSize:"), value)
}

// A dictionary of additional attributes that describe how to export content
// as a PDF file.
//
// # Discussion
//
// Although `attributes` is a read-only property, you can modify the mutable
// dictionary associated with it. Typically, this dictionary contains custom
// attributes or parameters that are set by a custom accessory view in the PDF
// panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSPDFInfo/attributes
func (p NSPDFInfo) Attributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("attributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
