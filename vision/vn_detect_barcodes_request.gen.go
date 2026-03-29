// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNDetectBarcodesRequest] class.
var (
	_VNDetectBarcodesRequestClass     VNDetectBarcodesRequestClass
	_VNDetectBarcodesRequestClassOnce sync.Once
)

func getVNDetectBarcodesRequestClass() VNDetectBarcodesRequestClass {
	_VNDetectBarcodesRequestClassOnce.Do(func() {
		_VNDetectBarcodesRequestClass = VNDetectBarcodesRequestClass{class: objc.GetClass("VNDetectBarcodesRequest")}
	})
	return _VNDetectBarcodesRequestClass
}

// GetVNDetectBarcodesRequestClass returns the class object for VNDetectBarcodesRequest.
func GetVNDetectBarcodesRequestClass() VNDetectBarcodesRequestClass {
	return getVNDetectBarcodesRequestClass()
}

type VNDetectBarcodesRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNDetectBarcodesRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNDetectBarcodesRequestClass) Alloc() VNDetectBarcodesRequest {
	rv := objc.Send[VNDetectBarcodesRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A request that detects barcodes in an image.
//
// # Overview
// 
// This request returns an array of [VNBarcodeObservation] objects, one for
// each barcode it detects.
//
// # Specifying Symbologies
//
//   - [VNDetectBarcodesRequest.SupportedSymbologiesAndReturnError]: Returns the barcode symbologies that the request supports.
//   - [VNDetectBarcodesRequest.Symbologies]: The barcode symbologies that the request detects in an image.
//   - [VNDetectBarcodesRequest.SetSymbologies]
//   - [VNDetectBarcodesRequest.CoalesceCompositeSymbologies]: A Boolean value that indicates whether to coalesce multiple codes based on the symbology.
//   - [VNDetectBarcodesRequest.SetCoalesceCompositeSymbologies]
//
// # Identifying Request Revisions
//
//   - [VNDetectBarcodesRequest.VNDetectBarcodesRequestRevision3]: A constant for specifying revision 3 of the barcode detection request.
//   - [VNDetectBarcodesRequest.VNDetectBarcodesRequestRevision2]: A constant for specifying revision 2 of the barcode detection request.
//   - [VNDetectBarcodesRequest.VNDetectBarcodesRequestRevision1]: A constant for specifying revision 1 of the barcode detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest
type VNDetectBarcodesRequest struct {
	VNImageBasedRequest
}

// VNDetectBarcodesRequestFromID constructs a [VNDetectBarcodesRequest] from an objc.ID.
//
// A request that detects barcodes in an image.
func VNDetectBarcodesRequestFromID(id objc.ID) VNDetectBarcodesRequest {
	return VNDetectBarcodesRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNDetectBarcodesRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNDetectBarcodesRequest] class.
//
// # Specifying Symbologies
//
//   - [IVNDetectBarcodesRequest.SupportedSymbologiesAndReturnError]: Returns the barcode symbologies that the request supports.
//   - [IVNDetectBarcodesRequest.Symbologies]: The barcode symbologies that the request detects in an image.
//   - [IVNDetectBarcodesRequest.SetSymbologies]
//   - [IVNDetectBarcodesRequest.CoalesceCompositeSymbologies]: A Boolean value that indicates whether to coalesce multiple codes based on the symbology.
//   - [IVNDetectBarcodesRequest.SetCoalesceCompositeSymbologies]
//
// # Identifying Request Revisions
//
//   - [IVNDetectBarcodesRequest.VNDetectBarcodesRequestRevision3]: A constant for specifying revision 3 of the barcode detection request.
//   - [IVNDetectBarcodesRequest.VNDetectBarcodesRequestRevision2]: A constant for specifying revision 2 of the barcode detection request.
//   - [IVNDetectBarcodesRequest.VNDetectBarcodesRequestRevision1]: A constant for specifying revision 1 of the barcode detection request.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest
type IVNDetectBarcodesRequest interface {
	IVNImageBasedRequest

	// Topic: Specifying Symbologies

	// Returns the barcode symbologies that the request supports.
	SupportedSymbologiesAndReturnError() ([]string, error)
	// The barcode symbologies that the request detects in an image.
	Symbologies() []string
	SetSymbologies(value []string)
	// A Boolean value that indicates whether to coalesce multiple codes based on the symbology.
	CoalesceCompositeSymbologies() bool
	SetCoalesceCompositeSymbologies(value bool)

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 3 of the barcode detection request.
	VNDetectBarcodesRequestRevision3() int
	// A constant for specifying revision 2 of the barcode detection request.
	VNDetectBarcodesRequestRevision2() int
	// A constant for specifying revision 1 of the barcode detection request.
	VNDetectBarcodesRequestRevision1() int
}

// Init initializes the instance.
func (d VNDetectBarcodesRequest) Init() VNDetectBarcodesRequest {
	rv := objc.Send[VNDetectBarcodesRequest](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VNDetectBarcodesRequest) Autorelease() VNDetectBarcodesRequest {
	rv := objc.Send[VNDetectBarcodesRequest](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNDetectBarcodesRequest creates a new VNDetectBarcodesRequest instance.
func NewVNDetectBarcodesRequest() VNDetectBarcodesRequest {
	class := getVNDetectBarcodesRequestClass()
	rv := objc.Send[VNDetectBarcodesRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new Vision request with an optional completion handler.
//
// completionHandler: The block to invoke after the request finishes processing.
//
// # Discussion
// 
// Vision executes the completion handler on the same queue that it executes
// the request; however, this queue differs from the one where you called
// [PerformRequestsError].
//
// See: https://developer.apple.com/documentation/Vision/VNRequest/init(completionHandler:)
func NewDetectBarcodesRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNDetectBarcodesRequest {
	instance := getVNDetectBarcodesRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNDetectBarcodesRequestFromID(rv)
}

// Returns the barcode symbologies that the request supports.
//
// # Return Value
// 
// An array of symbologies.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest/supportedSymbologies()
func (d VNDetectBarcodesRequest) SupportedSymbologiesAndReturnError() ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("supportedSymbologiesAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}

// The barcode symbologies that the request detects in an image.
//
// # Discussion
// 
// By default, a request scans for all symbologies. Specify a subset of
// symbologies to limit the request’s detection range.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest/symbologies
func (d VNDetectBarcodesRequest) Symbologies() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("symbologies"))
	return objc.ConvertSliceToStrings(rv)
}
func (d VNDetectBarcodesRequest) SetSymbologies(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setSymbologies:"), objectivec.StringSliceToNSArray(value))
}
// A Boolean value that indicates whether to coalesce multiple codes based on
// the symbology.
//
// See: https://developer.apple.com/documentation/Vision/VNDetectBarcodesRequest/coalesceCompositeSymbologies
func (d VNDetectBarcodesRequest) CoalesceCompositeSymbologies() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("coalesceCompositeSymbologies"))
	return rv
}
func (d VNDetectBarcodesRequest) SetCoalesceCompositeSymbologies(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setCoalesceCompositeSymbologies:"), value)
}
// A constant for specifying revision 3 of the barcode detection request.
//
// See: https://developer.apple.com/documentation/vision/vndetectbarcodesrequestrevision3
func (d VNDetectBarcodesRequest) VNDetectBarcodesRequestRevision3() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectBarcodesRequestRevision3"))
	return rv
}
// A constant for specifying revision 2 of the barcode detection request.
//
// See: https://developer.apple.com/documentation/vision/vndetectbarcodesrequestrevision2
func (d VNDetectBarcodesRequest) VNDetectBarcodesRequestRevision2() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectBarcodesRequestRevision2"))
	return rv
}
// A constant for specifying revision 1 of the barcode detection request.
//
// See: https://developer.apple.com/documentation/vision/vndetectbarcodesrequestrevision1
func (d VNDetectBarcodesRequest) VNDetectBarcodesRequestRevision1() int {
	rv := objc.Send[int](d.ID, objc.Sel("VNDetectBarcodesRequestRevision1"))
	return rv
}

