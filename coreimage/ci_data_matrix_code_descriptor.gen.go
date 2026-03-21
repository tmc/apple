// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [CIDataMatrixCodeDescriptor] class.
var (
	_CIDataMatrixCodeDescriptorClass     CIDataMatrixCodeDescriptorClass
	_CIDataMatrixCodeDescriptorClassOnce sync.Once
)

func getCIDataMatrixCodeDescriptorClass() CIDataMatrixCodeDescriptorClass {
	_CIDataMatrixCodeDescriptorClassOnce.Do(func() {
		_CIDataMatrixCodeDescriptorClass = CIDataMatrixCodeDescriptorClass{class: objc.GetClass("CIDataMatrixCodeDescriptor")}
	})
	return _CIDataMatrixCodeDescriptorClass
}

// GetCIDataMatrixCodeDescriptorClass returns the class object for CIDataMatrixCodeDescriptor.
func GetCIDataMatrixCodeDescriptorClass() CIDataMatrixCodeDescriptorClass {
	return getCIDataMatrixCodeDescriptorClass()
}

type CIDataMatrixCodeDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIDataMatrixCodeDescriptorClass) Alloc() CIDataMatrixCodeDescriptor {
	rv := objc.Send[CIDataMatrixCodeDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A concrete subclass the Core Image Barcode Descriptor that represents an
// Data Matrix code symbol.
//
// # Overview
// 
// A Data Matrix code symbol is a 2D barcode format defined by the ISO/IEC
// 16022:2006(E) standard. It encodes data in square or rectangular symbol
// with solid lines on the left and bottom sides
//
// # Creating a Descriptor
//
//   - [CIDataMatrixCodeDescriptor.InitWithPayloadRowCountColumnCountEccVersion]: Initializes a Data Matrix code descriptor for the given payload and parameters.
//
// # Examining a Descriptor
//
//   - [CIDataMatrixCodeDescriptor.ErrorCorrectedPayload]: The error-corrected payload containing the data encoded in the Data Matrix code symbol.
//   - [CIDataMatrixCodeDescriptor.RowCount]: The number of rows in the Data Matrix code symbol.
//   - [CIDataMatrixCodeDescriptor.ColumnCount]: The number of columns in the Data Matrix code symbol.
//   - [CIDataMatrixCodeDescriptor.EccVersion]: The error correction version of the Data Matrix code symbol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor
type CIDataMatrixCodeDescriptor struct {
	CIBarcodeDescriptor
}

// CIDataMatrixCodeDescriptorFromID constructs a [CIDataMatrixCodeDescriptor] from an objc.ID.
//
// A concrete subclass the Core Image Barcode Descriptor that represents an
// Data Matrix code symbol.
func CIDataMatrixCodeDescriptorFromID(id objc.ID) CIDataMatrixCodeDescriptor {
	return CIDataMatrixCodeDescriptor{CIBarcodeDescriptor: CIBarcodeDescriptorFromID(id)}
}
// NOTE: CIDataMatrixCodeDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIDataMatrixCodeDescriptor] class.
//
// # Creating a Descriptor
//
//   - [ICIDataMatrixCodeDescriptor.InitWithPayloadRowCountColumnCountEccVersion]: Initializes a Data Matrix code descriptor for the given payload and parameters.
//
// # Examining a Descriptor
//
//   - [ICIDataMatrixCodeDescriptor.ErrorCorrectedPayload]: The error-corrected payload containing the data encoded in the Data Matrix code symbol.
//   - [ICIDataMatrixCodeDescriptor.RowCount]: The number of rows in the Data Matrix code symbol.
//   - [ICIDataMatrixCodeDescriptor.ColumnCount]: The number of columns in the Data Matrix code symbol.
//   - [ICIDataMatrixCodeDescriptor.EccVersion]: The error correction version of the Data Matrix code symbol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor
type ICIDataMatrixCodeDescriptor interface {
	ICIBarcodeDescriptor

	// Topic: Creating a Descriptor

	// Initializes a Data Matrix code descriptor for the given payload and parameters.
	InitWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload foundation.INSData, rowCount int, columnCount int, eccVersion CIDataMatrixCodeECCVersion) CIDataMatrixCodeDescriptor

	// Topic: Examining a Descriptor

	// The error-corrected payload containing the data encoded in the Data Matrix code symbol.
	ErrorCorrectedPayload() foundation.INSData
	// The number of rows in the Data Matrix code symbol.
	RowCount() int
	// The number of columns in the Data Matrix code symbol.
	ColumnCount() int
	// The error correction version of the Data Matrix code symbol.
	EccVersion() CIDataMatrixCodeECCVersion
}

// Init initializes the instance.
func (d CIDataMatrixCodeDescriptor) Init() CIDataMatrixCodeDescriptor {
	rv := objc.Send[CIDataMatrixCodeDescriptor](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d CIDataMatrixCodeDescriptor) Autorelease() CIDataMatrixCodeDescriptor {
	rv := objc.Send[CIDataMatrixCodeDescriptor](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIDataMatrixCodeDescriptor creates a new CIDataMatrixCodeDescriptor instance.
func NewCIDataMatrixCodeDescriptor() CIDataMatrixCodeDescriptor {
	class := getCIDataMatrixCodeDescriptorClass()
	rv := objc.Send[CIDataMatrixCodeDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a Data Matrix code descriptor for the given payload and
// parameters.
//
// errorCorrectedPayload: The data to encode in the Data Matrix code symbol.
//
// rowCount: The number of rows in the Data Matrix code symbol.
//
// columnCount: The number of columns in the Data Matrix code symbol.
//
// eccVersion: The [CIDataMatrixCodeDescriptor.ECCVersion] for the Data Matrix code
// symbol.
// //
// [CIDataMatrixCodeDescriptor.ECCVersion]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum
//
// # Return Value
// 
// An initialized [CIAztecCodeDescriptor] instance or `nil` if the parameters
// are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/init(payload:rowCount:columnCount:eccVersion:)
func NewDataMatrixCodeDescriptorWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload foundation.INSData, rowCount int, columnCount int, eccVersion CIDataMatrixCodeECCVersion) CIDataMatrixCodeDescriptor {
	instance := getCIDataMatrixCodeDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPayload:rowCount:columnCount:eccVersion:"), errorCorrectedPayload, rowCount, columnCount, eccVersion)
	return CIDataMatrixCodeDescriptorFromID(rv)
}

// Initializes a Data Matrix code descriptor for the given payload and
// parameters.
//
// errorCorrectedPayload: The data to encode in the Data Matrix code symbol.
//
// rowCount: The number of rows in the Data Matrix code symbol.
//
// columnCount: The number of columns in the Data Matrix code symbol.
//
// eccVersion: The [CIDataMatrixCodeDescriptor.ECCVersion] for the Data Matrix code
// symbol.
// //
// [CIDataMatrixCodeDescriptor.ECCVersion]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum
//
// # Return Value
// 
// An initialized [CIAztecCodeDescriptor] instance or `nil` if the parameters
// are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/init(payload:rowCount:columnCount:eccVersion:)
func (d CIDataMatrixCodeDescriptor) InitWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload foundation.INSData, rowCount int, columnCount int, eccVersion CIDataMatrixCodeECCVersion) CIDataMatrixCodeDescriptor {
	rv := objc.Send[CIDataMatrixCodeDescriptor](d.ID, objc.Sel("initWithPayload:rowCount:columnCount:eccVersion:"), errorCorrectedPayload, rowCount, columnCount, eccVersion)
	return rv
}

// Creates a Data Matrix code descriptor for the given payload and parameters.
//
// errorCorrectedPayload: The data to encode in the Data Matrix code symbol.
//
// rowCount: The number of rows in the Data Matrix code symbol.
//
// columnCount: The number of columns in the Data Matrix code symbol.
//
// eccVersion: The [CIDataMatrixCodeDescriptor.ECCVersion] for the Data Matrix code
// symbol.
// //
// [CIDataMatrixCodeDescriptor.ECCVersion]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum
//
// # Return Value
// 
// An autoreleased [CIAztecCodeDescriptor] instance or `nil` if the parameters
// are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/descriptorWithPayload:rowCount:columnCount:eccVersion:
func (_CIDataMatrixCodeDescriptorClass CIDataMatrixCodeDescriptorClass) DescriptorWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload foundation.INSData, rowCount int, columnCount int, eccVersion CIDataMatrixCodeECCVersion) CIDataMatrixCodeDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_CIDataMatrixCodeDescriptorClass.class), objc.Sel("descriptorWithPayload:rowCount:columnCount:eccVersion:"), errorCorrectedPayload, rowCount, columnCount, eccVersion)
	return CIDataMatrixCodeDescriptorFromID(rv)
}

// The error-corrected payload containing the data encoded in the Data Matrix
// code symbol.
//
// # Discussion
// 
// DataMatrix symbols are specified bn ISO/IEC 16022:2006(E). ECC 200-type
// symbols will always have an even number of rows and columns.
// 
// For ECC 200-type symbols, the phases of encoding data into a symbol are
// described in section 5.1 – Encode procedure overview. The error corrected
// payload comprises the de-interleaved bits of the message described at the
// end of Step 1: Data encodation.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/errorCorrectedPayload-swift.property
func (d CIDataMatrixCodeDescriptor) ErrorCorrectedPayload() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("errorCorrectedPayload"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// The number of rows in the Data Matrix code symbol.
//
// # Discussion
// 
// Refer to ISO/IEC 16022:2006(E) for valid module row and column count
// combinations.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/rowCount-swift.property
func (d CIDataMatrixCodeDescriptor) RowCount() int {
	rv := objc.Send[int](d.ID, objc.Sel("rowCount"))
	return rv
}
// The number of columns in the Data Matrix code symbol.
//
// # Discussion
// 
// Refer to ISO/IEC 16022:2006(E) for valid module row and column count
// combinations.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/columnCount-swift.property
func (d CIDataMatrixCodeDescriptor) ColumnCount() int {
	rv := objc.Send[int](d.ID, objc.Sel("columnCount"))
	return rv
}
// The error correction version of the Data Matrix code symbol.
//
// # Discussion
// 
// The possible error correction version are enumerated in
// [CIDataMatrixCodeDescriptor.ECCVersion]. Any symbol with an even number of
// rows and columns will be ECC 200.
//
// [CIDataMatrixCodeDescriptor.ECCVersion]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum
//
// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/eccVersion-swift.property
func (d CIDataMatrixCodeDescriptor) EccVersion() CIDataMatrixCodeECCVersion {
	rv := objc.Send[CIDataMatrixCodeECCVersion](d.ID, objc.Sel("eccVersion"))
	return CIDataMatrixCodeECCVersion(rv)
}

