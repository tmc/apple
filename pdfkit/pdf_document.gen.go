// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PDFDocument] class.
var (
	_PDFDocumentClass     PDFDocumentClass
	_PDFDocumentClassOnce sync.Once
)

func getPDFDocumentClass() PDFDocumentClass {
	_PDFDocumentClassOnce.Do(func() {
		_PDFDocumentClass = PDFDocumentClass{class: objc.GetClass("PDFDocument")}
	})
	return _PDFDocumentClass
}

// GetPDFDocumentClass returns the class object for PDFDocument.
func GetPDFDocumentClass() PDFDocumentClass {
	return getPDFDocumentClass()
}

type PDFDocumentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFDocumentClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFDocumentClass) Alloc() PDFDocument {
	rv := objc.Send[PDFDocument](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents PDF data or a PDF file and defines methods for
// writing, searching, and selecting PDF data.
//
// # Overview
//
// The other utility classes are either instantiated from methods in
// [PDFDocument], as are [PDFPage] and [PDFOutline]; or support it, as do
// [PDFSelection] and [PDFDestination].
//
// You initialize a [PDFDocument] object with PDF data or with a URL to a PDF
// file. You can then ask for the page count, add or delete pages, perform a
// find, or parse selected content into an [NSString] object.
//
// # Initializing Documents
//
//   - [PDFDocument.InitWithURL]: Initializes a [PDFDocument] object with the contents at the specified URL (if the URL is invalid, this method returns [NULL]).
//   - [PDFDocument.InitWithData]: Initializes a [PDFDocument] object with the passed-in data.
//
// # Setting the Delegate
//
//   - [PDFDocument.Delegate]: The object acting as the delegate for the [PDFDocument] object.
//   - [PDFDocument.SetDelegate]
//
// # Instance Properties
//
//   - [PDFDocument.AccessPermissions]
//
// # Instance Methods
//
//   - [PDFDocument.SelectionFromPageAtPointToPageAtPointWithGranularity]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument
type PDFDocument struct {
	objectivec.Object
}

// PDFDocumentFromID constructs a [PDFDocument] from an objc.ID.
//
// An object that represents PDF data or a PDF file and defines methods for
// writing, searching, and selecting PDF data.
func PDFDocumentFromID(id objc.ID) PDFDocument {
	return PDFDocument{objectivec.Object{ID: id}}
}

// NOTE: PDFDocument adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFDocument] class.
//
// # Initializing Documents
//
//   - [IPDFDocument.InitWithURL]: Initializes a [PDFDocument] object with the contents at the specified URL (if the URL is invalid, this method returns [NULL]).
//   - [IPDFDocument.InitWithData]: Initializes a [PDFDocument] object with the passed-in data.
//
// # Setting the Delegate
//
//   - [IPDFDocument.Delegate]: The object acting as the delegate for the [PDFDocument] object.
//   - [IPDFDocument.SetDelegate]
//
// # Instance Properties
//
//   - [IPDFDocument.AccessPermissions]
//
// # Instance Methods
//
//   - [IPDFDocument.SelectionFromPageAtPointToPageAtPointWithGranularity]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument
type IPDFDocument interface {
	objectivec.IObject

	// Topic: Initializing Documents

	// Initializes a [PDFDocument] object with the contents at the specified URL (if the URL is invalid, this method returns [NULL]).
	InitWithURL(url foundation.INSURL) PDFDocument
	// Initializes a [PDFDocument] object with the passed-in data.
	InitWithData(data foundation.INSData) PDFDocument

	// Topic: Setting the Delegate

	// The object acting as the delegate for the [PDFDocument] object.
	Delegate() PDFDocumentDelegate
	SetDelegate(value PDFDocumentDelegate)

	// Topic: Instance Properties

	AccessPermissions() PDFAccessPermissions

	// Topic: Instance Methods

	SelectionFromPageAtPointToPageAtPointWithGranularity(startPage IPDFPage, startPoint corefoundation.CGPoint, endPage IPDFPage, endPoint corefoundation.CGPoint, granularity PDFSelectionGranularity) IPDFSelection

	// A Boolean value indicating whether you can create or modify document annotations, including form field entries.
	AllowsCommenting() bool
	// A Boolean value indicating whether you can extract content from the document, but only for the purpose of accessibility.
	AllowsContentAccessibility() bool
	// A Boolean value indicating whether the document allows copying of content to the Pasteboard.
	AllowsCopying() bool
	// A Boolean value indicating whether you can manage a document by inserting, deleting, and rotating pages.
	AllowsDocumentAssembly() bool
	// A Boolean value indicating whether you can modify the document contents except for document attributes.
	AllowsDocumentChanges() bool
	// A Boolean value indicating whether you can modify form field entries even if you can’t edit document annotations.
	AllowsFormFieldEntry() bool
	// A Boolean value indicating whether the document allows printing.
	AllowsPrinting() bool
	// A dictionary of document metadata.
	DocumentAttributes() foundation.INSDictionary
	SetDocumentAttributes(value foundation.INSDictionary)
	// The [CGPDFDocument] associated with the [PDFDocument] object.
	DocumentRef() coregraphics.CGPDFDocumentRef
	// The URL for the document.
	DocumentURL() foundation.INSURL
	// A Boolean value specifying whether the document is encrypted.
	IsEncrypted() bool
	// Returns a Boolean value indicating whether an asynchronous find operation is in progress.
	IsFinding() bool
	// A Boolean value indicating whether the document is locked.
	IsLocked() bool
	// The major version of the document.
	MajorVersion() int
	// The minor version of the document.
	MinorVersion() int
	// The document’s root outline to a PDF outline object.
	OutlineRoot() IPDFOutline
	SetOutlineRoot(value IPDFOutline)
	// The class that is allocated and initialized when page objects are created for the document.
	PageClass() objc.Class
	// The number of pages in the document.
	PageCount() uint
	// The permissions status of the PDF document.
	PermissionsStatus() PDFDocumentPermissions
	// Returns a selection representing the textual content of the entire document.
	SelectionForEntireDocument() IPDFSelection
	// A string representing the textual content for the entire document.
	String() string
	// Asynchronously finds all instances of the specified string in the document.
	BeginFindStringWithOptions(string_ string, options foundation.NSStringCompareOptions)
	// Asynchronously finds all instances of the specified array of strings in the document.
	BeginFindStringsWithOptions(strings []string, options foundation.NSStringCompareOptions)
	// Cancels a search initiated with [beginFindString(_:withOptions:)](<doc://com.apple.pdfkit/documentation/PDFKit/PDFDocument/beginFindString(_:withOptions:)>).
	CancelFindString()
	// Returns a representation of the document as an [NSData] object.
	DataRepresentation() foundation.INSData
	// Returns a representation of the document as an [NSData] object with additional options applied, such as filters.
	DataRepresentationWithOptions(options foundation.INSDictionary) foundation.INSData
	// Swaps one page with another.
	ExchangePageAtIndexWithPageAtIndex(indexA uint, indexB uint)
	// Synchronously finds the next occurance of a string after the specified selection (or before the selection if you specified [NSBackwardsSearch] as a search option.
	FindStringFromSelectionWithOptions(string_ string, selection IPDFSelection, options foundation.NSStringCompareOptions) IPDFSelection
	// Synchronously finds all instances of the specified string in the document.
	FindStringWithOptions(string_ string, options foundation.NSStringCompareOptions) []PDFSelection
	// Gets the index number for the specified page.
	IndexForPage(page IPDFPage) uint
	// Inserts a page at the specified index point.
	InsertPageAtIndex(page IPDFPage, index uint)
	// Returns the most likely parent PDF outline object for the selection.
	OutlineItemForSelection(selection IPDFSelection) IPDFOutline
	// Returns the page at the specified index number.
	PageAtIndex(index uint) IPDFPage
	// Returns a print operation suitable for printing the PDF document.
	PrintOperationForPrintInfoScalingModeAutoRotate(printInfo appkit.NSPrintInfo, scaleMode PDFPrintScalingMode, doRotate bool) appkit.NSPrintOperation
	// Removes the page at the specified index point.
	RemovePageAtIndex(index uint)
	// Returns the specified selection based on starting and ending character indexes.
	SelectionFromPageAtCharacterIndexToPageAtCharacterIndex(startPage IPDFPage, startCharacter uint, endPage IPDFPage, endCharacter uint) IPDFSelection
	// Returns the specified selection based on starting and ending points.
	SelectionFromPageAtPointToPageAtPoint(startPage IPDFPage, startPoint corefoundation.CGPoint, endPage IPDFPage, endPoint corefoundation.CGPoint) IPDFSelection
	// Attempts to unlock an encrypted document.
	UnlockWithPassword(password string) bool
	// Writes the document to a file at the specified path.
	WriteToFile(path string) bool
	// Writes the document to a file at the specified path with the specified options.
	WriteToFileWithOptions(path string, options foundation.INSDictionary) bool
	// Writes the document to a location specified by the passed-in URL.
	WriteToURL(url foundation.INSURL) bool
	// Writes the document to the specified URL with the specified options.
	WriteToURLWithOptions(url foundation.INSURL, options foundation.INSDictionary) bool
}

// Init initializes the instance.
func (p PDFDocument) Init() PDFDocument {
	rv := objc.Send[PDFDocument](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFDocument) Autorelease() PDFDocument {
	rv := objc.Send[PDFDocument](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFDocument creates a new PDFDocument instance.
func NewPDFDocument() PDFDocument {
	class := getPDFDocumentClass()
	rv := objc.Send[PDFDocument](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a [PDFDocument] object with the passed-in data.
//
// # Return Value
//
// A [PDFDocument] instance initialized with the passed-in PDF data, or [NULL]
// if the object could not be initialized.
//
// # Discussion
//
// The data must be PDF data encapsulated in an [NSData] object; otherwise
// this method returns [NULL].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/init(data:)
func NewPDFDocumentWithData(data foundation.INSData) PDFDocument {
	instance := getPDFDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return PDFDocumentFromID(rv)
}

// Initializes a [PDFDocument] object with the contents at the specified URL
// (if the URL is invalid, this method returns [NULL]).
//
// # Return Value
//
// A [PDFDocument] instance initialized with the data at the passed-in URL or
// [NULL] if the object could not be initialized or if the URL is invalid.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/init(url:)
func NewPDFDocumentWithURL(url foundation.INSURL) PDFDocument {
	instance := getPDFDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return PDFDocumentFromID(rv)
}

// Initializes a [PDFDocument] object with the contents at the specified URL
// (if the URL is invalid, this method returns [NULL]).
//
// # Return Value
//
// A [PDFDocument] instance initialized with the data at the passed-in URL or
// [NULL] if the object could not be initialized or if the URL is invalid.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/init(url:)
func (p PDFDocument) InitWithURL(url foundation.INSURL) PDFDocument {
	rv := objc.Send[PDFDocument](p.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// Initializes a [PDFDocument] object with the passed-in data.
//
// # Return Value
//
// A [PDFDocument] instance initialized with the passed-in PDF data, or [NULL]
// if the object could not be initialized.
//
// # Discussion
//
// The data must be PDF data encapsulated in an [NSData] object; otherwise
// this method returns [NULL].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/init(data:)
func (p PDFDocument) InitWithData(data foundation.INSData) PDFDocument {
	rv := objc.Send[PDFDocument](p.ID, objc.Sel("initWithData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/selection(from:at:to:at:with:)
func (p PDFDocument) SelectionFromPageAtPointToPageAtPointWithGranularity(startPage IPDFPage, startPoint corefoundation.CGPoint, endPage IPDFPage, endPoint corefoundation.CGPoint, granularity PDFSelectionGranularity) IPDFSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectionFromPage:atPoint:toPage:atPoint:withGranularity:"), startPage, startPoint, endPage, endPoint, granularity)
	return PDFSelectionFromID(rv)
}

// Asynchronously finds all instances of the specified string in the document.
//
// # Discussion
//
// This method returns immediately. It causes notifications to be issued when
// searching begins and ends, on each search hit, and when the search proceeds
// to a new page. For options, refer to [Searching, Comparing, and Sorting
// Strings].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/beginFindString(_:withOptions:)
//
// [Searching, Comparing, and Sorting Strings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/SearchingStrings.html#//apple_ref/doc/uid/20000149
func (p PDFDocument) BeginFindStringWithOptions(string_ string, options foundation.NSStringCompareOptions) {
	objc.Send[objc.ID](p.ID, objc.Sel("beginFindString:withOptions:"), objc.String(string_), options)
}

// Asynchronously finds all instances of the specified array of strings in the
// document.
//
// # Discussion
//
// This method returns immediately. It causes notifications to be issued when
// searching begins and ends, on each search hit, and when the search proceeds
// to a new page. For options, refer to [Searching, Comparing, and Sorting
// Strings].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/beginFindStrings(_:withOptions:)
//
// [Searching, Comparing, and Sorting Strings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/SearchingStrings.html#//apple_ref/doc/uid/20000149
func (p PDFDocument) BeginFindStringsWithOptions(strings []string, options foundation.NSStringCompareOptions) {
	objc.Send[objc.ID](p.ID, objc.Sel("beginFindStrings:withOptions:"), objectivec.StringSliceToNSArray(strings), options)
}

// Cancels a search initiated with [BeginFindStringWithOptions].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/cancelFindString()
func (p PDFDocument) CancelFindString() {
	objc.Send[objc.ID](p.ID, objc.Sel("cancelFindString"))
}

// Returns a representation of the document as an [NSData] object.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/dataRepresentation()
func (p PDFDocument) DataRepresentation() foundation.INSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dataRepresentation"))
	return foundation.NSDataFromID(rv)
}

// Returns a representation of the document as an [NSData] object with
// additional options applied, such as filters.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/dataRepresentation(options:)
func (p PDFDocument) DataRepresentationWithOptions(options foundation.INSDictionary) foundation.INSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dataRepresentationWithOptions:"), options)
	return foundation.NSDataFromID(rv)
}

// Swaps one page with another.
//
// # Discussion
//
// This method raises an exception if either `index` value is out of bounds.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/exchangePage(at:withPageAt:)
func (p PDFDocument) ExchangePageAtIndexWithPageAtIndex(indexA uint, indexB uint) {
	objc.Send[objc.ID](p.ID, objc.Sel("exchangePageAtIndex:withPageAtIndex:"), indexA, indexB)
}

// Synchronously finds the next occurance of a string after the specified
// selection (or before the selection if you specified [NSBackwardsSearch] as
// a search option.
//
// # Discussion
//
// Matches are returned as a [PDFSelection] object. If the search reaches the
// end (or beginning) of the document without any hits, this method returns
// [NULL].
//
// If you pass [NULL] for the selection, this method begins searching from the
// beginning of the document (or the end, if you specified
// [NSBackwardsSearch]).
//
// You can use this method to implement “Find Again” behavior. For
// options, refer to [Searching, Comparing, and Sorting Strings].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/findString(_:fromSelection:withOptions:)
//
// [Searching, Comparing, and Sorting Strings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/SearchingStrings.html#//apple_ref/doc/uid/20000149
func (p PDFDocument) FindStringFromSelectionWithOptions(string_ string, selection IPDFSelection, options foundation.NSStringCompareOptions) IPDFSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("findString:fromSelection:withOptions:"), objc.String(string_), selection, options)
	return PDFSelectionFromID(rv)
}

// Synchronously finds all instances of the specified string in the document.
//
// # Discussion
//
// Each hit gets added to an [NSArray] object as a [PDFSelection] object. If
// there are no hits, this method returns an empty array.
//
// Use this method when the complete search process will be brief and when you
// don’t need the flexibility or control offered by
// [BeginFindStringWithOptions]. For options, refer to [Searching, Comparing,
// and Sorting Strings].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/findString(_:withOptions:)
//
// [Searching, Comparing, and Sorting Strings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/SearchingStrings.html#//apple_ref/doc/uid/20000149
func (p PDFDocument) FindStringWithOptions(string_ string, options foundation.NSStringCompareOptions) []PDFSelection {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("findString:withOptions:"), objc.String(string_), options)
	return objc.ConvertSlice(rv, func(id objc.ID) PDFSelection {
		return PDFSelectionFromID(id)
	})
}

// Gets the index number for the specified page.
//
// # Discussion
//
// Indexes are zero-based. This method raises an exception and returns
// [NSNotFound] if `page` is not found.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/index(for:)
func (p PDFDocument) IndexForPage(page IPDFPage) uint {
	rv := objc.Send[uint](p.ID, objc.Sel("indexForPage:"), page)
	return rv
}

// Inserts a page at the specified index point.
//
// # Discussion
//
// This method raises an exception if `index` is out of bounds.
//
// Be aware that a PDF viewing application might use the size of the first
// page in the document as representative of all page sizes when reporting the
// size of a document. If you need to get the actual size of an individual
// page, you can use [BoundsForBox] (note that the size is returned in points,
// which are typically converted to inches or centimeters by PDF viewing
// applications).
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/insert(_:at:)
func (p PDFDocument) InsertPageAtIndex(page IPDFPage, index uint) {
	objc.Send[objc.ID](p.ID, objc.Sel("insertPage:atIndex:"), page, index)
}

// Returns the most likely parent PDF outline object for the selection.
//
// selection: The area of the document currently selected by the user. A selection can
// span multiple outline items, but only the point representing the first
// character is considered.
//
// # Return Value
//
// The PDF outline object that is the most likely parent of the specified
// selection. Note that only the point representing the first character of the
// selection is considered in this method.
//
// # Discussion
//
// Typically, outlines represent structural items such as chapters. You can
// use this method to identify the chapter that a selection falls within.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/outlineItem(for:)
func (p PDFDocument) OutlineItemForSelection(selection IPDFSelection) IPDFOutline {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("outlineItemForSelection:"), selection)
	return PDFOutlineFromID(rv)
}

// Returns the page at the specified index number.
//
// # Discussion
//
// Indexes are zero based. This method raises an exception if `index` is out
// of bounds.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/page(at:)
func (p PDFDocument) PageAtIndex(index uint) IPDFPage {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pageAtIndex:"), index)
	return PDFPageFromID(rv)
}

// Returns a print operation suitable for printing the PDF document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/printOperation(for:scalingMode:autoRotate:)
func (p PDFDocument) PrintOperationForPrintInfoScalingModeAutoRotate(printInfo appkit.NSPrintInfo, scaleMode PDFPrintScalingMode, doRotate bool) appkit.NSPrintOperation {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("printOperationForPrintInfo:scalingMode:autoRotate:"), printInfo, scaleMode, doRotate)
	return appkit.NSPrintOperationFromID(rv)
}

// Removes the page at the specified index point.
//
// # Discussion
//
// This method raises an exception if `index` is out of bounds.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/removePage(at:)
func (p PDFDocument) RemovePageAtIndex(index uint) {
	objc.Send[objc.ID](p.ID, objc.Sel("removePageAtIndex:"), index)
}

// Returns the specified selection based on starting and ending character
// indexes.
//
// # Discussion
//
// The selection begins at `startChar` on `startPage` and ends at `endChar` on
// `endPage`. The starting and ending index values must be in the range of the
// number of characters (as returned by [NumberOfCharacters]) within the
// respective [PDFPage] objects.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/selection(from:atCharacterIndex:to:atCharacterIndex:)
func (p PDFDocument) SelectionFromPageAtCharacterIndexToPageAtCharacterIndex(startPage IPDFPage, startCharacter uint, endPage IPDFPage, endCharacter uint) IPDFSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectionFromPage:atCharacterIndex:toPage:atCharacterIndex:"), startPage, startCharacter, endPage, endCharacter)
	return PDFSelectionFromID(rv)
}

// Returns the specified selection based on starting and ending points.
//
// # Discussion
//
// The selection begins at `startPt` on `startPage` and ends at `endPt` on
// `endPage`. The starting and ending points should be specified in page
// space, relative to their respective pages.
//
// The starting and ending points can be on the same page. In this case,
// invoking this method is equivalent to sending the “ message to a [PDFPage]
// object.
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/selection(from:at:to:at:)
func (p PDFDocument) SelectionFromPageAtPointToPageAtPoint(startPage IPDFPage, startPoint corefoundation.CGPoint, endPage IPDFPage, endPoint corefoundation.CGPoint) IPDFSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectionFromPage:atPoint:toPage:atPoint:"), startPage, startPoint, endPage, endPoint)
	return PDFSelectionFromID(rv)
}

// Attempts to unlock an encrypted document.
//
// password: The password to unlock an encrypted document (you cannot lock an unlocked
// PDF document by using an incorrect password).
//
// # Return Value
//
// true if the specified password unlocks the document, false otherwise.
//
// # Discussion
//
// If the password is correct, this method returns true, and a
// [PDFDocumentDidUnlockNotification] notification is sent. Once unlocked, you
// cannot use this function to relock the document.
//
// If you attempt to unlock an already unlocked document, one of the following
// occurs:
//
// - If the document is unlocked with full owner permissions,
// `unlockWithPassword` does nothing and returns true. The password string is
// ignored. - If the document is unlocked with only user permissions,
// `unlockWithPassword` attempts to obtain full owner permissions with the
// password string. If the string fails, the document maintains its user
// permissions. In either case, this method returns true.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/unlock(withPassword:)
func (p PDFDocument) UnlockWithPassword(password string) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("unlockWithPassword:"), objc.String(password))
	return rv
}

// Writes the document to a file at the specified path.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/write(toFile:)
func (p PDFDocument) WriteToFile(path string) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("writeToFile:"), objc.String(path))
	return rv
}

// Writes the document to a file at the specified path with the specified
// options.
//
// # Discussion
//
// The most commonly-used options are `kCGPDFContextOwnerPassword`,
// `kCGPDFContextUserPassword`, `kCGPDFContextAllowsCopying` and
// `kCGPDFContextAllowsPrinting`. For more details about these options, see
// [Auxiliary Dictionary Keys].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/write(toFile:withOptions:)
//
// [Auxiliary Dictionary Keys]: https://developer.apple.com/documentation/CoreGraphics/auxiliary-dictionary-keys
func (p PDFDocument) WriteToFileWithOptions(path string, options foundation.INSDictionary) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("writeToFile:withOptions:"), objc.String(path), options)
	return rv
}

// Writes the document to a location specified by the passed-in URL.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/write(to:)
func (p PDFDocument) WriteToURL(url foundation.INSURL) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("writeToURL:"), url)
	return rv
}

// Writes the document to the specified URL with the specified options.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/write(to:withOptions:)
func (p PDFDocument) WriteToURLWithOptions(url foundation.INSURL, options foundation.INSDictionary) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("writeToURL:withOptions:"), url, options)
	return rv
}

// The object acting as the delegate for the [PDFDocument] object.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/delegate
func (p PDFDocument) Delegate() PDFDocumentDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return PDFDocumentDelegateObjectFromID(rv)
}
func (p PDFDocument) SetDelegate(value PDFDocumentDelegate) {
	objc.Send[struct{}](p.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/accessPermissions
func (p PDFDocument) AccessPermissions() PDFAccessPermissions {
	rv := objc.Send[PDFAccessPermissions](p.ID, objc.Sel("accessPermissions"))
	return PDFAccessPermissions(rv)
}

// A Boolean value indicating whether you can create or modify document
// annotations, including form field entries.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsCommenting
func (p PDFDocument) AllowsCommenting() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsCommenting"))
	return rv
}

// A Boolean value indicating whether you can extract content from the
// document, but only for the purpose of accessibility.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsContentAccessibility
func (p PDFDocument) AllowsContentAccessibility() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsContentAccessibility"))
	return rv
}

// A Boolean value indicating whether the document allows copying of content
// to the Pasteboard.
//
// # Discussion
//
// The ability to copy content from a PDF document is an attribute unrelated
// to whether the document is locked or unlocked. It depends on the PDF
// permissions set by the document’s author.
//
// This method only determines the desired permissions setting in the PDF
// document; it is up to the application to enforce (or ignore) the
// permissions.
//
// This method always returns true if the document is not encrypted. Note that
// in many cases an encrypted document may still be readable by all users due
// to the standard empty string password. For more details about user and
// owner passwords, see the Adobe PDF specification.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsCopying
func (p PDFDocument) AllowsCopying() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsCopying"))
	return rv
}

// A Boolean value indicating whether you can manage a document by inserting,
// deleting, and rotating pages.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsDocumentAssembly
func (p PDFDocument) AllowsDocumentAssembly() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsDocumentAssembly"))
	return rv
}

// A Boolean value indicating whether you can modify the document contents
// except for document attributes.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsDocumentChanges
func (p PDFDocument) AllowsDocumentChanges() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsDocumentChanges"))
	return rv
}

// A Boolean value indicating whether you can modify form field entries even
// if you can’t edit document annotations.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsFormFieldEntry
func (p PDFDocument) AllowsFormFieldEntry() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsFormFieldEntry"))
	return rv
}

// A Boolean value indicating whether the document allows printing.
//
// # Discussion
//
// The ability to print a PDF document is an attribute unrelated to whether
// the document is locked or unlocked. It depends on the PDF permissions set
// by the document’s author.
//
// This method only determines the desired permissions setting in the PDF
// document; it is up to the application to enforce (or ignore) the
// permissions.
//
// This method always returns true if the document is not encrypted. Note that
// in many cases an encrypted document may still be readable by all users due
// to the standard empty string password. For more details about user and
// owner passwords, see the Adobe PDF specification.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsPrinting
func (p PDFDocument) AllowsPrinting() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsPrinting"))
	return rv
}

// A dictionary of document metadata.
//
// # Return Value
//
// The dictionary of document metadata. The dictionary may be empty, or only
// some of the keys may have associated values.
//
// # Discussion
//
// Metadata is optional for PDF documents.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/documentAttributes
func (p PDFDocument) DocumentAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("documentAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p PDFDocument) SetDocumentAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setDocumentAttributes:"), value)
}

// The [CGPDFDocument] associated with the [PDFDocument] object.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/documentRef
func (p PDFDocument) DocumentRef() coregraphics.CGPDFDocumentRef {
	rv := objc.Send[coregraphics.CGPDFDocumentRef](p.ID, objc.Sel("documentRef"))
	return coregraphics.CGPDFDocumentRef(rv)
}

// The URL for the document.
//
// # Return Value
//
// The URL for the document; may return [NULL] if the document was created
// from an [NSData] object.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/documentURL
func (p PDFDocument) DocumentURL() foundation.INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("documentURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// A Boolean value specifying whether the document is encrypted.
//
// # Return Value
//
// true if the document is encrypted, whether it is locked or unlocked; false
// otherwise.
//
// # Discussion
//
// If encrypted, reading the document requires a password.
//
// Encrypted documents whose password is the empty string are unlocked
// automatically upon opening, because PDF Kit tries the empty string as a
// password if none is supplied. Use the [UnlockWithPassword] method to unlock
// a document using a password.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/isEncrypted
func (p PDFDocument) IsEncrypted() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isEncrypted"))
	return rv
}

// Returns a Boolean value indicating whether an asynchronous find operation
// is in progress.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/isFinding
func (p PDFDocument) IsFinding() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isFinding"))
	return rv
}

// A Boolean value indicating whether the document is locked.
//
// # Return Value
//
// true if the document is locked; false otherwise.
//
// # Discussion
//
// Only encrypted documents can be locked. Encrypted documents whose password
// is the empty string are unlocked automatically upon opening, because PDF
// Kit tries the empty string as a password if none is supplied. Use the
// [UnlockWithPassword] method to unlock a document using a password.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/isLocked
func (p PDFDocument) IsLocked() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isLocked"))
	return rv
}

// The major version of the document.
//
// # Return Value
//
// The major version of the document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/majorVersion
func (p PDFDocument) MajorVersion() int {
	rv := objc.Send[int](p.ID, objc.Sel("majorVersion"))
	return rv
}

// The minor version of the document.
//
// # Return Value
//
// The minor version of the document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/minorVersion
func (p PDFDocument) MinorVersion() int {
	rv := objc.Send[int](p.ID, objc.Sel("minorVersion"))
	return rv
}

// The document’s root outline to a PDF outline object.
//
// # Discussion
//
// When a PDF document is saved, the outline tree structure is written out to
// the destination PDF file.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/outlineRoot
func (p PDFDocument) OutlineRoot() IPDFOutline {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("outlineRoot"))
	return PDFOutlineFromID(objc.ID(rv))
}
func (p PDFDocument) SetOutlineRoot(value IPDFOutline) {
	objc.Send[struct{}](p.ID, objc.Sel("setOutlineRoot:"), value)
}

// The class that is allocated and initialized when page objects are created
// for the document.
//
// # Discussion
//
// If you want to supply a custom page class, subclass [PDFDocument] and
// implement this method to return your custom class. Note that your custom
// class must be a subclass of [PDFPage]; otherwise, the behavior is
// undefined.
//
// The default implementation of `pageClass` returns `[PDFPage class]`.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/pageClass
func (p PDFDocument) PageClass() objc.Class {
	rv := objc.Send[objc.Class](p.ID, objc.Sel("pageClass"))
	return rv
}

// The number of pages in the document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/pageCount
func (p PDFDocument) PageCount() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("pageCount"))
	return rv
}

// The permissions status of the PDF document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/permissionsStatus
func (p PDFDocument) PermissionsStatus() PDFDocumentPermissions {
	rv := objc.Send[PDFDocumentPermissions](p.ID, objc.Sel("permissionsStatus"))
	return PDFDocumentPermissions(rv)
}

// Returns a selection representing the textual content of the entire
// document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/selectionForEntireDocument
func (p PDFDocument) SelectionForEntireDocument() IPDFSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectionForEntireDocument"))
	return PDFSelectionFromID(objc.ID(rv))
}

// A string representing the textual content for the entire document.
//
// # Return Value
//
// A string that represents the textual content of the entire document.
//
// # Discussion
//
// Pages are delimited with linefeed characters.
//
// This is a convenience method, equivalent to creating a selection object for
// the entire document and then invoking the [PDFSelection] class’s [String]
// method.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDocument/string
func (p PDFDocument) String() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}
