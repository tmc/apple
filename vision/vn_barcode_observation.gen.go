// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNBarcodeObservation] class.
var (
	_VNBarcodeObservationClass     VNBarcodeObservationClass
	_VNBarcodeObservationClassOnce sync.Once
)

func getVNBarcodeObservationClass() VNBarcodeObservationClass {
	_VNBarcodeObservationClassOnce.Do(func() {
		_VNBarcodeObservationClass = VNBarcodeObservationClass{class: objc.GetClass("VNBarcodeObservation")}
	})
	return _VNBarcodeObservationClass
}

// GetVNBarcodeObservationClass returns the class object for VNBarcodeObservation.
func GetVNBarcodeObservationClass() VNBarcodeObservationClass {
	return getVNBarcodeObservationClass()
}

type VNBarcodeObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNBarcodeObservationClass) Alloc() VNBarcodeObservation {
	rv := objc.Send[VNBarcodeObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that represents barcode information that an image analysis
// request detects.
//
// # Overview
// 
// This type of observation results from a [VNDetectBarcodesRequest]. It
// contains information about the detected barcode, including parsed payload
// data for supported symbologies.
//
// # Parsing the Payload
//
//   - [VNBarcodeObservation.PayloadStringValue]: A string value that represents the barcode payload.
//   - [VNBarcodeObservation.PayloadData]: The raw data representation of the barcode’s payload.
//   - [VNBarcodeObservation.SupplementalPayloadString]: The supplemental code decoded as a string value.
//   - [VNBarcodeObservation.SupplementalPayloadData]
//   - [VNBarcodeObservation.SupplementalCompositeType]: The supplemental composite type.
//   - [VNBarcodeObservation.IsGS1DataCarrier]: A Boolean value that indicates whether the barcode carries any global standards data.
//
// # Reading Barcode Descriptors
//
//   - [VNBarcodeObservation.BarcodeDescriptor]: An object that describes the low-level details about the barcode and its data.
//
// # Identifying Barcode Types
//
//   - [VNBarcodeObservation.Symbology]: The symbology of the observed barcode.
//
// # Identifying Barcode Colors
//
//   - [VNBarcodeObservation.IsColorInverted]: A Boolean value that indicates whether the barcode is color inverted.
//
// See: https://developer.apple.com/documentation/Vision/VNBarcodeObservation
type VNBarcodeObservation struct {
	VNRectangleObservation
}

// VNBarcodeObservationFromID constructs a [VNBarcodeObservation] from an objc.ID.
//
// An object that represents barcode information that an image analysis
// request detects.
func VNBarcodeObservationFromID(id objc.ID) VNBarcodeObservation {
	return VNBarcodeObservation{VNRectangleObservation: VNRectangleObservationFromID(id)}
}
// NOTE: VNBarcodeObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNBarcodeObservation] class.
//
// # Parsing the Payload
//
//   - [IVNBarcodeObservation.PayloadStringValue]: A string value that represents the barcode payload.
//   - [IVNBarcodeObservation.PayloadData]: The raw data representation of the barcode’s payload.
//   - [IVNBarcodeObservation.SupplementalPayloadString]: The supplemental code decoded as a string value.
//   - [IVNBarcodeObservation.SupplementalPayloadData]
//   - [IVNBarcodeObservation.SupplementalCompositeType]: The supplemental composite type.
//   - [IVNBarcodeObservation.IsGS1DataCarrier]: A Boolean value that indicates whether the barcode carries any global standards data.
//
// # Reading Barcode Descriptors
//
//   - [IVNBarcodeObservation.BarcodeDescriptor]: An object that describes the low-level details about the barcode and its data.
//
// # Identifying Barcode Types
//
//   - [IVNBarcodeObservation.Symbology]: The symbology of the observed barcode.
//
// # Identifying Barcode Colors
//
//   - [IVNBarcodeObservation.IsColorInverted]: A Boolean value that indicates whether the barcode is color inverted.
//
// See: https://developer.apple.com/documentation/Vision/VNBarcodeObservation
type IVNBarcodeObservation interface {
	IVNRectangleObservation

	// Topic: Parsing the Payload

	// A string value that represents the barcode payload.
	PayloadStringValue() string
	// The raw data representation of the barcode’s payload.
	PayloadData() foundation.INSData
	// The supplemental code decoded as a string value.
	SupplementalPayloadString() string
	SupplementalPayloadData() foundation.INSData
	// The supplemental composite type.
	SupplementalCompositeType() VNBarcodeCompositeType
	// A Boolean value that indicates whether the barcode carries any global standards data.
	IsGS1DataCarrier() bool

	// Topic: Reading Barcode Descriptors

	// An object that describes the low-level details about the barcode and its data.
	BarcodeDescriptor() objectivec.IObject

	// Topic: Identifying Barcode Types

	// The symbology of the observed barcode.
	Symbology() VNBarcodeSymbology

	// Topic: Identifying Barcode Colors

	// A Boolean value that indicates whether the barcode is color inverted.
	IsColorInverted() bool

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (b VNBarcodeObservation) Init() VNBarcodeObservation {
	rv := objc.Send[VNBarcodeObservation](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b VNBarcodeObservation) Autorelease() VNBarcodeObservation {
	rv := objc.Send[VNBarcodeObservation](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNBarcodeObservation creates a new VNBarcodeObservation instance.
func NewVNBarcodeObservation() VNBarcodeObservation {
	class := getVNBarcodeObservationClass()
	rv := objc.Send[VNBarcodeObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a rectangle observation from its corner points.
//
// requestRevision: The rectangle detector revision number. A higher revision indicates more
// recent iterations of the framework.
//
// topLeft: The upper-left corner point.
//
// bottomLeft: The lower-left corner point.
//
// bottomRight: The lower-right corner point.
//
// topRight: The upper-right corner point.
//
// See: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:bottomLeft:bottomRight:topRight:)
func NewBarcodeObservationRectangleObservationWithRequestRevisionTopLeftBottomLeftBottomRightTopRight(requestRevision uint, topLeft corefoundation.CGPoint, bottomLeft corefoundation.CGPoint, bottomRight corefoundation.CGPoint, topRight corefoundation.CGPoint) VNBarcodeObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNBarcodeObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:bottomLeft:bottomRight:topRight:"), requestRevision, topLeft, bottomLeft, bottomRight, topRight)
	return VNBarcodeObservationFromID(rv)
}


// Creates a rectangle observation from its corner points.
//
// requestRevision: The rectangle detector revision number. A higher revision indicates more
// recent iterations of the framework.
//
// topLeft: The upper-left corner point.
//
// topRight: The upper-right corner point.
//
// bottomRight: The lower-right corner point.
//
// bottomLeft: The lower-left corner point.
//
// See: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:topRight:bottomRight:bottomLeft:)
func NewBarcodeObservationRectangleObservationWithRequestRevisionTopLeftTopRightBottomRightBottomLeft(requestRevision uint, topLeft corefoundation.CGPoint, topRight corefoundation.CGPoint, bottomRight corefoundation.CGPoint, bottomLeft corefoundation.CGPoint) VNBarcodeObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNBarcodeObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:topRight:bottomRight:bottomLeft:"), requestRevision, topLeft, topRight, bottomRight, bottomLeft)
	return VNBarcodeObservationFromID(rv)
}


// Creates an observation with a bounding box.
//
// boundingBox: The observation’s bounding box, in coordinates normalized to the
// dimensions of the processed image, with its origin at the image’s
// lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(boundingBox:)
func NewBarcodeObservationWithBoundingBox(boundingBox corefoundation.CGRect) VNBarcodeObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNBarcodeObservationClass().class), objc.Sel("observationWithBoundingBox:"), boundingBox)
	return VNBarcodeObservationFromID(rv)
}


// Creates an observation with a revision number and bounding box.
//
// requestRevision: The revision of the request to use.
//
// boundingBox: The observation’s bounding box, in coordinates normalized to the
// dimensions of the processed image, with its origin at the image’s
// lower-left corner.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation/init(requestRevision:boundingBox:)
func NewBarcodeObservationWithRequestRevisionBoundingBox(requestRevision uint, boundingBox corefoundation.CGRect) VNBarcodeObservation {
	rv := objc.Send[objc.ID](objc.ID(getVNBarcodeObservationClass().class), objc.Sel("observationWithRequestRevision:boundingBox:"), requestRevision, boundingBox)
	return VNBarcodeObservationFromID(rv)
}






func (b VNBarcodeObservation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeWithCoder:"), coder)
}












// A string value that represents the barcode payload.
//
// # Discussion
// 
// Depending on the symbology or the payload data itself, a string
// representation of the payload may not be available.
//
// See: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/payloadStringValue
func (b VNBarcodeObservation) PayloadStringValue() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("payloadStringValue"))
	return foundation.NSStringFromID(rv).String()
}



// The raw data representation of the barcode’s payload.
//
// See: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/payloadData
func (b VNBarcodeObservation) PayloadData() foundation.INSData {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("payloadData"))
	return foundation.NSDataFromID(objc.ID(rv))
}



// The supplemental code decoded as a string value.
//
// See: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/supplementalPayloadString
func (b VNBarcodeObservation) SupplementalPayloadString() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("supplementalPayloadString"))
	return foundation.NSStringFromID(rv).String()
}



// See: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/supplementalPayloadData
func (b VNBarcodeObservation) SupplementalPayloadData() foundation.INSData {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("supplementalPayloadData"))
	return foundation.NSDataFromID(objc.ID(rv))
}



// The supplemental composite type.
//
// See: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/supplementalCompositeType
func (b VNBarcodeObservation) SupplementalCompositeType() VNBarcodeCompositeType {
	rv := objc.Send[VNBarcodeCompositeType](b.ID, objc.Sel("supplementalCompositeType"))
	return VNBarcodeCompositeType(rv)
}



// A Boolean value that indicates whether the barcode carries any global
// standards data.
//
// See: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/isGS1DataCarrier
func (b VNBarcodeObservation) IsGS1DataCarrier() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isGS1DataCarrier"))
	return rv
}



// An object that describes the low-level details about the barcode and its
// data.
//
// # Discussion
// 
// Use this object to have Core Image regenerate the observed barcode.
//
// See: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/barcodeDescriptor
func (b VNBarcodeObservation) BarcodeDescriptor() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("barcodeDescriptor"))
	return objectivec.Object{ID: rv}
}



// The symbology of the observed barcode.
//
// See: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/symbology
func (b VNBarcodeObservation) Symbology() VNBarcodeSymbology {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("symbology"))
	return VNBarcodeSymbology(foundation.NSStringFromID(rv).String())
}



// A Boolean value that indicates whether the barcode is color inverted.
//
// See: https://developer.apple.com/documentation/Vision/VNBarcodeObservation/isColorInverted
func (b VNBarcodeObservation) IsColorInverted() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isColorInverted"))
	return rv
}



























