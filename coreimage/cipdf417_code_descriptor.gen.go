// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [CIPDF417CodeDescriptor] class.
var (
	_CIPDF417CodeDescriptorClass     CIPDF417CodeDescriptorClass
	_CIPDF417CodeDescriptorClassOnce sync.Once
)

func getCIPDF417CodeDescriptorClass() CIPDF417CodeDescriptorClass {
	_CIPDF417CodeDescriptorClassOnce.Do(func() {
		_CIPDF417CodeDescriptorClass = CIPDF417CodeDescriptorClass{class: objc.GetClass("CIPDF417CodeDescriptor")}
	})
	return _CIPDF417CodeDescriptorClass
}

// GetCIPDF417CodeDescriptorClass returns the class object for CIPDF417CodeDescriptor.
func GetCIPDF417CodeDescriptorClass() CIPDF417CodeDescriptorClass {
	return getCIPDF417CodeDescriptorClass()
}

type CIPDF417CodeDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIPDF417CodeDescriptorClass) Alloc() CIPDF417CodeDescriptor {
	rv := objc.Send[CIPDF417CodeDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A concrete subclass of Core Image Barcode Descriptor that represents a
// PDF417 symbol.
//
// # Overview
// 
// PDF417 is a stacked linear barcode symbol format used predominantly in
// transport, ID cards, and inventory management. Each pattern in the code
// comprises 4 bars and spaces, 17 units long.
// 
// Refer to the ISO/IEC 15438:2006(E) for the PDF417 symbol specification.
//
// # Creating a Descriptor
//
//   - [CIPDF417CodeDescriptor.InitWithPayloadIsCompactRowCountColumnCount]: Initializes an PDF417 code descriptor for the given payload and parameters.
//
// # Examining a Descriptor
//
//   - [CIPDF417CodeDescriptor.ErrorCorrectedPayload]: The error-corrected payload containing the data encoded in the PDF417 code symbol.
//   - [CIPDF417CodeDescriptor.IsCompact]: A boolean value telling if the PDF417 code is compact.
//   - [CIPDF417CodeDescriptor.RowCount]: The number of rows in the PDF417 code symbol.
//   - [CIPDF417CodeDescriptor.ColumnCount]: The number of columns in the PDF417 code symbol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor
type CIPDF417CodeDescriptor struct {
	CIBarcodeDescriptor
}

// CIPDF417CodeDescriptorFromID constructs a [CIPDF417CodeDescriptor] from an objc.ID.
//
// A concrete subclass of Core Image Barcode Descriptor that represents a
// PDF417 symbol.
func CIPDF417CodeDescriptorFromID(id objc.ID) CIPDF417CodeDescriptor {
	return CIPDF417CodeDescriptor{CIBarcodeDescriptor: CIBarcodeDescriptorFromID(id)}
}
// NOTE: CIPDF417CodeDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIPDF417CodeDescriptor] class.
//
// # Creating a Descriptor
//
//   - [ICIPDF417CodeDescriptor.InitWithPayloadIsCompactRowCountColumnCount]: Initializes an PDF417 code descriptor for the given payload and parameters.
//
// # Examining a Descriptor
//
//   - [ICIPDF417CodeDescriptor.ErrorCorrectedPayload]: The error-corrected payload containing the data encoded in the PDF417 code symbol.
//   - [ICIPDF417CodeDescriptor.IsCompact]: A boolean value telling if the PDF417 code is compact.
//   - [ICIPDF417CodeDescriptor.RowCount]: The number of rows in the PDF417 code symbol.
//   - [ICIPDF417CodeDescriptor.ColumnCount]: The number of columns in the PDF417 code symbol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor
type ICIPDF417CodeDescriptor interface {
	ICIBarcodeDescriptor

	// Topic: Creating a Descriptor

	// Initializes an PDF417 code descriptor for the given payload and parameters.
	InitWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload foundation.INSData, isCompact bool, rowCount int, columnCount int) CIPDF417CodeDescriptor

	// Topic: Examining a Descriptor

	// The error-corrected payload containing the data encoded in the PDF417 code symbol.
	ErrorCorrectedPayload() foundation.INSData
	// A boolean value telling if the PDF417 code is compact.
	IsCompact() bool
	// The number of rows in the PDF417 code symbol.
	RowCount() int
	// The number of columns in the PDF417 code symbol.
	ColumnCount() int
}

// Init initializes the instance.
func (p CIPDF417CodeDescriptor) Init() CIPDF417CodeDescriptor {
	rv := objc.Send[CIPDF417CodeDescriptor](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p CIPDF417CodeDescriptor) Autorelease() CIPDF417CodeDescriptor {
	rv := objc.Send[CIPDF417CodeDescriptor](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIPDF417CodeDescriptor creates a new CIPDF417CodeDescriptor instance.
func NewCIPDF417CodeDescriptor() CIPDF417CodeDescriptor {
	class := getCIPDF417CodeDescriptorClass()
	rv := objc.Send[CIPDF417CodeDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an PDF417 code descriptor for the given payload and parameters.
//
// errorCorrectedPayload: The data to encode in the PDF417 code symbol.
//
// isCompact: A Boolean indicating whether or not the PDF417 code is compact.
//
// rowCount: The number of rows in the PDF417 code, from 3 to 90.
//
// columnCount: The number of columns in the Aztec code, from 1 to 30.
//
// # Return Value
// 
// An initialized [CIPDF417CodeDescriptor] instance or `nil` if the parameters
// are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/init(payload:isCompact:rowCount:columnCount:)
func NewPDF417CodeDescriptorWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload foundation.INSData, isCompact bool, rowCount int, columnCount int) CIPDF417CodeDescriptor {
	instance := getCIPDF417CodeDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPayload:isCompact:rowCount:columnCount:"), errorCorrectedPayload, isCompact, rowCount, columnCount)
	return CIPDF417CodeDescriptorFromID(rv)
}

// Initializes an PDF417 code descriptor for the given payload and parameters.
//
// errorCorrectedPayload: The data to encode in the PDF417 code symbol.
//
// isCompact: A Boolean indicating whether or not the PDF417 code is compact.
//
// rowCount: The number of rows in the PDF417 code, from 3 to 90.
//
// columnCount: The number of columns in the Aztec code, from 1 to 30.
//
// # Return Value
// 
// An initialized [CIPDF417CodeDescriptor] instance or `nil` if the parameters
// are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/init(payload:isCompact:rowCount:columnCount:)
func (p CIPDF417CodeDescriptor) InitWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload foundation.INSData, isCompact bool, rowCount int, columnCount int) CIPDF417CodeDescriptor {
	rv := objc.Send[CIPDF417CodeDescriptor](p.ID, objc.Sel("initWithPayload:isCompact:rowCount:columnCount:"), errorCorrectedPayload, isCompact, rowCount, columnCount)
	return rv
}

// Creates an PDF417 code descriptor for the given payload and parameters.
//
// errorCorrectedPayload: The data to encode in the PDF417 code symbol.
//
// isCompact: A Boolean indicating whether or not the PDF417 code is compact.
//
// rowCount: The number of rows in the PDF417 code, from 3 to 90.
//
// columnCount: The number of columns in the Aztec code, from 1 to 30.
//
// # Return Value
// 
// An autoreleased [CIPDF417CodeDescriptor] instance or `nil` if the
// parameters are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/descriptorWithPayload:isCompact:rowCount:columnCount:
func (_CIPDF417CodeDescriptorClass CIPDF417CodeDescriptorClass) DescriptorWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload foundation.INSData, isCompact bool, rowCount int, columnCount int) CIPDF417CodeDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_CIPDF417CodeDescriptorClass.class), objc.Sel("descriptorWithPayload:isCompact:rowCount:columnCount:"), errorCorrectedPayload, isCompact, rowCount, columnCount)
	return CIPDF417CodeDescriptorFromID(rv)
}

// The error-corrected payload containing the data encoded in the PDF417 code
// symbol.
//
// # Discussion
// 
// The first codeword indicates the number of data codewords in the
// errorCorrectedPayload.
// 
// PDF417 codes are comprised of a start character on the left and a stop
// character on the right. Each row begins and ends with special characters
// indicating the current row as well as information about the dimensions of
// the PDF417 symbol. The errorCorrectedPayload represents the sequence of
// PDF417 codewords that make up the body of the message. The first codeword
// indicates the number of codewords in the message. This count includes the
// “count” codeword and any padding codewords, but does not include the
// error correction codewords. Each codeword is a 16-bit value in the range of
// 0…928. The sequence is to be interpreted as described in the PDF417 bar
// code symbology specification – ISO/IEC 15438:2006(E).
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/errorCorrectedPayload-swift.property
func (p CIPDF417CodeDescriptor) ErrorCorrectedPayload() foundation.INSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("errorCorrectedPayload"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// A boolean value telling if the PDF417 code is compact.
//
// # Discussion
// 
// Compact PDF417 symbols have abbreviated right-side guard bars.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/isCompact-swift.property
func (p CIPDF417CodeDescriptor) IsCompact() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isCompact"))
	return rv
}
// The number of rows in the PDF417 code symbol.
//
// # Discussion
// 
// Valid row count values are from 3 to 90.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/rowCount-swift.property
func (p CIPDF417CodeDescriptor) RowCount() int {
	rv := objc.Send[int](p.ID, objc.Sel("rowCount"))
	return rv
}
// The number of columns in the PDF417 code symbol.
//
// # Discussion
// 
// Valid column count values are from 1 to 30. This count excluded the columns
// used to indicate the symbol structure.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPDF417CodeDescriptor/columnCount-swift.property
func (p CIPDF417CodeDescriptor) ColumnCount() int {
	rv := objc.Send[int](p.ID, objc.Sel("columnCount"))
	return rv
}

