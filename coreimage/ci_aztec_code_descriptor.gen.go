// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [CIAztecCodeDescriptor] class.
var (
	_CIAztecCodeDescriptorClass     CIAztecCodeDescriptorClass
	_CIAztecCodeDescriptorClassOnce sync.Once
)

func getCIAztecCodeDescriptorClass() CIAztecCodeDescriptorClass {
	_CIAztecCodeDescriptorClassOnce.Do(func() {
		_CIAztecCodeDescriptorClass = CIAztecCodeDescriptorClass{class: objc.GetClass("CIAztecCodeDescriptor")}
	})
	return _CIAztecCodeDescriptorClass
}

// GetCIAztecCodeDescriptorClass returns the class object for CIAztecCodeDescriptor.
func GetCIAztecCodeDescriptorClass() CIAztecCodeDescriptorClass {
	return getCIAztecCodeDescriptorClass()
}

type CIAztecCodeDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIAztecCodeDescriptorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIAztecCodeDescriptorClass) Alloc() CIAztecCodeDescriptor {
	rv := objc.Send[CIAztecCodeDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A concrete subclass the Core Image Barcode Descriptor that represents an
// Aztec code symbol.
//
// # Overview
// 
// An Aztec code symbol is a 2D barcode format defined by the ISO/IEC
// 24778:2008 standard. It encodes data in concentric square rings around a
// central bullseye pattern.
//
// # Creating a Descriptor
//
//   - [CIAztecCodeDescriptor.InitWithPayloadIsCompactLayerCountDataCodewordCount]: Initializes an Aztec code descriptor for the given payload and parameters.
//
// # Examining a Descriptor
//
//   - [CIAztecCodeDescriptor.ErrorCorrectedPayload]: The error-corrected payload that comprises the the Aztec code symbol.
//   - [CIAztecCodeDescriptor.IsCompact]: A Boolean value telling if the Aztec code is compact.
//   - [CIAztecCodeDescriptor.LayerCount]: The number of data layers in the Aztec code symbol.
//   - [CIAztecCodeDescriptor.DataCodewordCount]: The number of non-error-correction codewords carried by the Aztec code symbol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor
type CIAztecCodeDescriptor struct {
	CIBarcodeDescriptor
}

// CIAztecCodeDescriptorFromID constructs a [CIAztecCodeDescriptor] from an objc.ID.
//
// A concrete subclass the Core Image Barcode Descriptor that represents an
// Aztec code symbol.
func CIAztecCodeDescriptorFromID(id objc.ID) CIAztecCodeDescriptor {
	return CIAztecCodeDescriptor{CIBarcodeDescriptor: CIBarcodeDescriptorFromID(id)}
}
// NOTE: CIAztecCodeDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIAztecCodeDescriptor] class.
//
// # Creating a Descriptor
//
//   - [ICIAztecCodeDescriptor.InitWithPayloadIsCompactLayerCountDataCodewordCount]: Initializes an Aztec code descriptor for the given payload and parameters.
//
// # Examining a Descriptor
//
//   - [ICIAztecCodeDescriptor.ErrorCorrectedPayload]: The error-corrected payload that comprises the the Aztec code symbol.
//   - [ICIAztecCodeDescriptor.IsCompact]: A Boolean value telling if the Aztec code is compact.
//   - [ICIAztecCodeDescriptor.LayerCount]: The number of data layers in the Aztec code symbol.
//   - [ICIAztecCodeDescriptor.DataCodewordCount]: The number of non-error-correction codewords carried by the Aztec code symbol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor
type ICIAztecCodeDescriptor interface {
	ICIBarcodeDescriptor

	// Topic: Creating a Descriptor

	// Initializes an Aztec code descriptor for the given payload and parameters.
	InitWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload foundation.INSData, isCompact bool, layerCount int, dataCodewordCount int) CIAztecCodeDescriptor

	// Topic: Examining a Descriptor

	// The error-corrected payload that comprises the the Aztec code symbol.
	ErrorCorrectedPayload() foundation.INSData
	// A Boolean value telling if the Aztec code is compact.
	IsCompact() bool
	// The number of data layers in the Aztec code symbol.
	LayerCount() int
	// The number of non-error-correction codewords carried by the Aztec code symbol.
	DataCodewordCount() int
}

// Init initializes the instance.
func (a CIAztecCodeDescriptor) Init() CIAztecCodeDescriptor {
	rv := objc.Send[CIAztecCodeDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a CIAztecCodeDescriptor) Autorelease() CIAztecCodeDescriptor {
	rv := objc.Send[CIAztecCodeDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIAztecCodeDescriptor creates a new CIAztecCodeDescriptor instance.
func NewCIAztecCodeDescriptor() CIAztecCodeDescriptor {
	class := getCIAztecCodeDescriptorClass()
	rv := objc.Send[CIAztecCodeDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an Aztec code descriptor for the given payload and parameters.
//
// errorCorrectedPayload: The data to encode in the Aztec code symbol.
//
// isCompact: A Boolean indicating whether or not the Aztec code is compact.
//
// layerCount: The number of layers in the Aztec code, from 1 to 32.
//
// dataCodewordCount: The number of codewords in the Aztec code, from 1 to 2048.
//
// # Return Value
// 
// An initialized [CIAztecCodeDescriptor] instance or `nil` if the parameters
// are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/init(payload:isCompact:layerCount:dataCodewordCount:)
func NewAztecCodeDescriptorWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload foundation.INSData, isCompact bool, layerCount int, dataCodewordCount int) CIAztecCodeDescriptor {
	instance := getCIAztecCodeDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPayload:isCompact:layerCount:dataCodewordCount:"), errorCorrectedPayload, isCompact, layerCount, dataCodewordCount)
	return CIAztecCodeDescriptorFromID(rv)
}

// Initializes an Aztec code descriptor for the given payload and parameters.
//
// errorCorrectedPayload: The data to encode in the Aztec code symbol.
//
// isCompact: A Boolean indicating whether or not the Aztec code is compact.
//
// layerCount: The number of layers in the Aztec code, from 1 to 32.
//
// dataCodewordCount: The number of codewords in the Aztec code, from 1 to 2048.
//
// # Return Value
// 
// An initialized [CIAztecCodeDescriptor] instance or `nil` if the parameters
// are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/init(payload:isCompact:layerCount:dataCodewordCount:)
func (a CIAztecCodeDescriptor) InitWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload foundation.INSData, isCompact bool, layerCount int, dataCodewordCount int) CIAztecCodeDescriptor {
	rv := objc.Send[CIAztecCodeDescriptor](a.ID, objc.Sel("initWithPayload:isCompact:layerCount:dataCodewordCount:"), errorCorrectedPayload, isCompact, layerCount, dataCodewordCount)
	return rv
}

// Creates an Aztec code descriptor for the given payload and parameters.
//
// errorCorrectedPayload: The data to encode in the Aztec code symbol.
//
// isCompact: A Boolean indicating whether or not the Aztec code is compact.
//
// layerCount: The number of layers in the Aztec code, from 1 to 32.
//
// dataCodewordCount: The number of codewords in the Aztec code, from 1 to 2048.
//
// # Return Value
// 
// An autoreleased [CIAztecCodeDescriptor] instance or `nil` if the parameters
// are invalid
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/descriptorWithPayload:isCompact:layerCount:dataCodewordCount:
func (_CIAztecCodeDescriptorClass CIAztecCodeDescriptorClass) DescriptorWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload foundation.INSData, isCompact bool, layerCount int, dataCodewordCount int) CIAztecCodeDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_CIAztecCodeDescriptorClass.class), objc.Sel("descriptorWithPayload:isCompact:layerCount:dataCodewordCount:"), errorCorrectedPayload, isCompact, layerCount, dataCodewordCount)
	return CIAztecCodeDescriptorFromID(rv)
}

// The error-corrected payload that comprises the the Aztec code symbol.
//
// # Discussion
// 
// Aztec Codes are formally specified in ISO/IEC 24778:2008(E).
// 
// The error corrected payload consists of the 6-, 8-, 10-, or 12-bit message
// codewords produced at the end of the step described in section 7.3.1.2
// “Formation of data codewords”, which exists immediately prior to adding
// error correction. These codewords have dummy bits inserted to ensure that
// an entire codeword isn’t all 0’s or all 1’s. Clients will need to
// remove these extra bits as part of interpreting the payload.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/errorCorrectedPayload-swift.property
func (a CIAztecCodeDescriptor) ErrorCorrectedPayload() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("errorCorrectedPayload"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// A Boolean value telling if the Aztec code is compact.
//
// # Discussion
// 
// Compact Aztec symbols use one-fewer ring in the central finder pattern than
// full-range Aztec symbols of the same number of data layers.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/isCompact-swift.property
func (a CIAztecCodeDescriptor) IsCompact() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isCompact"))
	return rv
}
// The number of data layers in the Aztec code symbol.
//
// # Discussion
// 
// Combined with [IsCompact], the number of data layers determines the number
// of modules in the Aztec Code symbol. Valid values range from 1 to 32.
// Compact symbols can have up to 4 data layers.
// 
// The number of data layers also determines the number of bits in each data
// codeword of the message carried by the Aztec Code symbol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/layerCount-swift.property
func (a CIAztecCodeDescriptor) LayerCount() int {
	rv := objc.Send[int](a.ID, objc.Sel("layerCount"))
	return rv
}
// The number of non-error-correction codewords carried by the Aztec code
// symbol.
//
// # Discussion
// 
// Used to determine the level of error correction in conjunction with the
// number of data layers. Valid values are 1 to 2048. Compact symbols can have
// up to 64 message codewords.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/dataCodewordCount-swift.property
func (a CIAztecCodeDescriptor) DataCodewordCount() int {
	rv := objc.Send[int](a.ID, objc.Sel("dataCodewordCount"))
	return rv
}

