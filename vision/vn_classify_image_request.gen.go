// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNClassifyImageRequest] class.
var (
	_VNClassifyImageRequestClass     VNClassifyImageRequestClass
	_VNClassifyImageRequestClassOnce sync.Once
)

func getVNClassifyImageRequestClass() VNClassifyImageRequestClass {
	_VNClassifyImageRequestClassOnce.Do(func() {
		_VNClassifyImageRequestClass = VNClassifyImageRequestClass{class: objc.GetClass("VNClassifyImageRequest")}
	})
	return _VNClassifyImageRequestClass
}

// GetVNClassifyImageRequestClass returns the class object for VNClassifyImageRequest.
func GetVNClassifyImageRequestClass() VNClassifyImageRequestClass {
	return getVNClassifyImageRequestClass()
}

type VNClassifyImageRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNClassifyImageRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNClassifyImageRequestClass) Alloc() VNClassifyImageRequest {
	rv := objc.Send[VNClassifyImageRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A request to classify an image.
//
// # Overview
//
// This type of request produces a collection of [VNClassificationObservation]
// objects that describe an image. Access the classifications through
// [knownClassifications(forRevision:)].
//
// # Accessing Results
//
//   - [VNClassifyImageRequest.SupportedIdentifiersAndReturnError]: Returns the classification identifiers that the request supports in its current configuration.
//
// # Specifying Algorithm Revision
//
//   - [VNClassifyImageRequest.VNClassifyImageRequestRevision1]: A constant for specifying the first revision of the image-classification request.
//
// See: https://developer.apple.com/documentation/Vision/VNClassifyImageRequest
//
// [knownClassifications(forRevision:)]: https://developer.apple.com/documentation/Vision/VNClassifyImageRequest/knownClassifications(forRevision:)
type VNClassifyImageRequest struct {
	VNImageBasedRequest
}

// VNClassifyImageRequestFromID constructs a [VNClassifyImageRequest] from an objc.ID.
//
// A request to classify an image.
func VNClassifyImageRequestFromID(id objc.ID) VNClassifyImageRequest {
	return VNClassifyImageRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}

// NOTE: VNClassifyImageRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNClassifyImageRequest] class.
//
// # Accessing Results
//
//   - [IVNClassifyImageRequest.SupportedIdentifiersAndReturnError]: Returns the classification identifiers that the request supports in its current configuration.
//
// # Specifying Algorithm Revision
//
//   - [IVNClassifyImageRequest.VNClassifyImageRequestRevision1]: A constant for specifying the first revision of the image-classification request.
//
// See: https://developer.apple.com/documentation/Vision/VNClassifyImageRequest
type IVNClassifyImageRequest interface {
	IVNImageBasedRequest

	// Topic: Accessing Results

	// Returns the classification identifiers that the request supports in its current configuration.
	SupportedIdentifiersAndReturnError() ([]string, error)

	// Topic: Specifying Algorithm Revision

	// A constant for specifying the first revision of the image-classification request.
	VNClassifyImageRequestRevision1() int
}

// Init initializes the instance.
func (c VNClassifyImageRequest) Init() VNClassifyImageRequest {
	rv := objc.Send[VNClassifyImageRequest](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VNClassifyImageRequest) Autorelease() VNClassifyImageRequest {
	rv := objc.Send[VNClassifyImageRequest](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNClassifyImageRequest creates a new VNClassifyImageRequest instance.
func NewVNClassifyImageRequest() VNClassifyImageRequest {
	class := getVNClassifyImageRequestClass()
	rv := objc.Send[VNClassifyImageRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewClassifyImageRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNClassifyImageRequest {
	instance := getVNClassifyImageRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNClassifyImageRequestFromID(rv)
}

// Returns the classification identifiers that the request supports in its
// current configuration.
//
// # Return Value
//
// An array of supported identifiers.
//
// See: https://developer.apple.com/documentation/Vision/VNClassifyImageRequest/supportedIdentifiers()
func (c VNClassifyImageRequest) SupportedIdentifiersAndReturnError() ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedIdentifiersAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}

// A constant for specifying the first revision of the image-classification
// request.
//
// See: https://developer.apple.com/documentation/vision/vnclassifyimagerequestrevision1
func (c VNClassifyImageRequest) VNClassifyImageRequestRevision1() int {
	rv := objc.Send[int](c.ID, objc.Sel("VNClassifyImageRequestRevision1"))
	return rv
}
