// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [CIQRCodeDescriptor] class.
var (
	_CIQRCodeDescriptorClass     CIQRCodeDescriptorClass
	_CIQRCodeDescriptorClassOnce sync.Once
)

func getCIQRCodeDescriptorClass() CIQRCodeDescriptorClass {
	_CIQRCodeDescriptorClassOnce.Do(func() {
		_CIQRCodeDescriptorClass = CIQRCodeDescriptorClass{class: objc.GetClass("CIQRCodeDescriptor")}
	})
	return _CIQRCodeDescriptorClass
}

// GetCIQRCodeDescriptorClass returns the class object for CIQRCodeDescriptor.
func GetCIQRCodeDescriptorClass() CIQRCodeDescriptorClass {
	return getCIQRCodeDescriptorClass()
}

type CIQRCodeDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIQRCodeDescriptorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIQRCodeDescriptorClass) Alloc() CIQRCodeDescriptor {
	rv := objc.Send[CIQRCodeDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A concrete subclass of the Core Image Barcode Descriptor that represents a
// square QR code symbol.
//
// # Overview
// 
// ISO/IEC 18004 defines versions from 1 to 40, where a higher symbol version
// indicates a larger data-carrying capacity. QR Codes can encode text, vCard
// contact information, or Uniform Resource Identifiers (URI).
//
// # Creating a Descriptor
//
//   - [CIQRCodeDescriptor.InitWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel]: Initializes a QR code descriptor for the given payload and parameters.
//
// # Examining a Descriptor
//
//   - [CIQRCodeDescriptor.ErrorCorrectedPayload]: The error-corrected codeword payload that comprises the QR code symbol.
//   - [CIQRCodeDescriptor.SymbolVersion]: The version of the QR code which corresponds to the size of the QR code symbol.
//   - [CIQRCodeDescriptor.MaskPattern]: The data mask pattern for the QR code symbol.
//   - [CIQRCodeDescriptor.ErrorCorrectionLevel]: The error correction level of the QR code symbol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor
type CIQRCodeDescriptor struct {
	CIBarcodeDescriptor
}

// CIQRCodeDescriptorFromID constructs a [CIQRCodeDescriptor] from an objc.ID.
//
// A concrete subclass of the Core Image Barcode Descriptor that represents a
// square QR code symbol.
func CIQRCodeDescriptorFromID(id objc.ID) CIQRCodeDescriptor {
	return CIQRCodeDescriptor{CIBarcodeDescriptor: CIBarcodeDescriptorFromID(id)}
}
// NOTE: CIQRCodeDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIQRCodeDescriptor] class.
//
// # Creating a Descriptor
//
//   - [ICIQRCodeDescriptor.InitWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel]: Initializes a QR code descriptor for the given payload and parameters.
//
// # Examining a Descriptor
//
//   - [ICIQRCodeDescriptor.ErrorCorrectedPayload]: The error-corrected codeword payload that comprises the QR code symbol.
//   - [ICIQRCodeDescriptor.SymbolVersion]: The version of the QR code which corresponds to the size of the QR code symbol.
//   - [ICIQRCodeDescriptor.MaskPattern]: The data mask pattern for the QR code symbol.
//   - [ICIQRCodeDescriptor.ErrorCorrectionLevel]: The error correction level of the QR code symbol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor
type ICIQRCodeDescriptor interface {
	ICIBarcodeDescriptor

	// Topic: Creating a Descriptor

	// Initializes a QR code descriptor for the given payload and parameters.
	InitWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload foundation.INSData, symbolVersion int, maskPattern uint8, errorCorrectionLevel CIQRCodeErrorCorrectionLevel) CIQRCodeDescriptor

	// Topic: Examining a Descriptor

	// The error-corrected codeword payload that comprises the QR code symbol.
	ErrorCorrectedPayload() foundation.INSData
	// The version of the QR code which corresponds to the size of the QR code symbol.
	SymbolVersion() int
	// The data mask pattern for the QR code symbol.
	MaskPattern() uint8
	// The error correction level of the QR code symbol.
	ErrorCorrectionLevel() CIQRCodeErrorCorrectionLevel
}

// Init initializes the instance.
func (q CIQRCodeDescriptor) Init() CIQRCodeDescriptor {
	rv := objc.Send[CIQRCodeDescriptor](q.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q CIQRCodeDescriptor) Autorelease() CIQRCodeDescriptor {
	rv := objc.Send[CIQRCodeDescriptor](q.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIQRCodeDescriptor creates a new CIQRCodeDescriptor instance.
func NewCIQRCodeDescriptor() CIQRCodeDescriptor {
	class := getCIQRCodeDescriptorClass()
	rv := objc.Send[CIQRCodeDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a QR code descriptor for the given payload and parameters.
//
// errorCorrectedPayload: The data to encode in the QR code symbol.
//
// symbolVersion: The symbol version, from 1 through 40.
//
// maskPattern: The mask pattern to use in the QR code, from 0 to 7.
//
// errorCorrectionLevel: The QR code’s error correction level: L, M, Q, or H.
//
// # Return Value
// 
// An initialized [CIAztecCodeDescriptor] instance or `nil` if the parameters
// are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/init(payload:symbolVersion:maskPattern:errorCorrectionLevel:)
func NewQRCodeDescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload foundation.INSData, symbolVersion int, maskPattern uint8, errorCorrectionLevel CIQRCodeErrorCorrectionLevel) CIQRCodeDescriptor {
	instance := getCIQRCodeDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:"), errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
	return CIQRCodeDescriptorFromID(rv)
}

// Initializes a QR code descriptor for the given payload and parameters.
//
// errorCorrectedPayload: The data to encode in the QR code symbol.
//
// symbolVersion: The symbol version, from 1 through 40.
//
// maskPattern: The mask pattern to use in the QR code, from 0 to 7.
//
// errorCorrectionLevel: The QR code’s error correction level: L, M, Q, or H.
//
// # Return Value
// 
// An initialized [CIAztecCodeDescriptor] instance or `nil` if the parameters
// are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/init(payload:symbolVersion:maskPattern:errorCorrectionLevel:)
func (q CIQRCodeDescriptor) InitWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload foundation.INSData, symbolVersion int, maskPattern uint8, errorCorrectionLevel CIQRCodeErrorCorrectionLevel) CIQRCodeDescriptor {
	rv := objc.Send[CIQRCodeDescriptor](q.ID, objc.Sel("initWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:"), errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
	return rv
}

// Creates a QR code descriptor for the given payload and parameters.
//
// errorCorrectedPayload: The data to encode in the QR code symbol.
//
// symbolVersion: The symbol version, from 1 through 40.
//
// maskPattern: The mask pattern to use in the QR code, from 0 to 7.
//
// errorCorrectionLevel: The QR code’s error correction level: L, M, Q, or H.
//
// # Return Value
// 
// An autoreleased [CIAztecCodeDescriptor] instance or `nil` if the parameters
// are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/descriptorWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:
func (_CIQRCodeDescriptorClass CIQRCodeDescriptorClass) DescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload foundation.INSData, symbolVersion int, maskPattern uint8, errorCorrectionLevel CIQRCodeErrorCorrectionLevel) CIQRCodeDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_CIQRCodeDescriptorClass.class), objc.Sel("descriptorWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:"), errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
	return CIQRCodeDescriptorFromID(rv)
}

// The error-corrected codeword payload that comprises the QR code symbol.
//
// # Discussion
// 
// QR Codes are formally specified in ISO/IEC 18004:2006(E). Section 6.4.10
// “Bitstream to codeword conversion” specifies the set of 8-bit codewords
// in the symbol immediately prior to splitting the message into blocks and
// applying error correction.
// 
// During decode, error correction is applied and if successful, the message
// is re-ordered to the state immediately following “Bitstream to codeword
// conversion.”
// 
// The `errorCorrectedPayload` corresponds to this sequence of 8-bit
// codewords.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectedPayload-swift.property
func (q CIQRCodeDescriptor) ErrorCorrectedPayload() foundation.INSData {
	rv := objc.Send[objc.ID](q.ID, objc.Sel("errorCorrectedPayload"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// The version of the QR code which corresponds to the size of the QR code
// symbol.
//
// # Discussion
// 
// ISO/IEC 18004 defines versions from 1 to 40, where a higher symbol version
// indicates a larger data-carrying capacity. This field is required in order
// to properly interpret the error corrected payload.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/symbolVersion-swift.property
func (q CIQRCodeDescriptor) SymbolVersion() int {
	rv := objc.Send[int](q.ID, objc.Sel("symbolVersion"))
	return rv
}
// The data mask pattern for the QR code symbol.
//
// # Discussion
// 
// QR Codes support eight data mask patterns, which are used to avoid large
// black or large white areas inside the symbol body. Valid values range from
// 0 to 7.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/maskPattern-swift.property
func (q CIQRCodeDescriptor) MaskPattern() uint8 {
	rv := objc.Send[uint8](q.ID, objc.Sel("maskPattern"))
	return rv
}
// The error correction level of the QR code symbol.
//
// # Discussion
// 
// QR Codes support four levels of Reed-Solomon error correction.
// 
// The possible error correction levels are enumerated in
// [CIDataMatrixCodeDescriptor.ECCVersion].
//
// [CIDataMatrixCodeDescriptor.ECCVersion]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectionLevel-swift.property
func (q CIQRCodeDescriptor) ErrorCorrectionLevel() CIQRCodeErrorCorrectionLevel {
	rv := objc.Send[CIQRCodeErrorCorrectionLevel](q.ID, objc.Sel("errorCorrectionLevel"))
	return CIQRCodeErrorCorrectionLevel(rv)
}

