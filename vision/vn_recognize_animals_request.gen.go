// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNRecognizeAnimalsRequest] class.
var (
	_VNRecognizeAnimalsRequestClass     VNRecognizeAnimalsRequestClass
	_VNRecognizeAnimalsRequestClassOnce sync.Once
)

func getVNRecognizeAnimalsRequestClass() VNRecognizeAnimalsRequestClass {
	_VNRecognizeAnimalsRequestClassOnce.Do(func() {
		_VNRecognizeAnimalsRequestClass = VNRecognizeAnimalsRequestClass{class: objc.GetClass("VNRecognizeAnimalsRequest")}
	})
	return _VNRecognizeAnimalsRequestClass
}

// GetVNRecognizeAnimalsRequestClass returns the class object for VNRecognizeAnimalsRequest.
func GetVNRecognizeAnimalsRequestClass() VNRecognizeAnimalsRequestClass {
	return getVNRecognizeAnimalsRequestClass()
}

type VNRecognizeAnimalsRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNRecognizeAnimalsRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNRecognizeAnimalsRequestClass) Alloc() VNRecognizeAnimalsRequest {
	rv := objc.Send[VNRecognizeAnimalsRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A request that recognizes animals in an image.
//
// # Overview
// 
// Use the [knownAnimalIdentifiers(forRevision:)] method to determine which
// animals the request supports.
//
// [knownAnimalIdentifiers(forRevision:)]: https://developer.apple.com/documentation/Vision/VNRecognizeAnimalsRequest/knownAnimalIdentifiers(forRevision:)
//
// # Identifying Animals
//
//   - [VNRecognizeAnimalsRequest.SupportedIdentifiersAndReturnError]: Returns the identifiers of the animals that the request detects.
//
// # Identifying Request Revisions
//
//   - [VNRecognizeAnimalsRequest.VNRecognizeAnimalsRequestRevision2]: A constant for specifying revision 2 of the animal recognition request.
//   - [VNRecognizeAnimalsRequest.VNRecognizeAnimalsRequestRevision1]: A constant for specifying revision 1 of the animal recognition request.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeAnimalsRequest
type VNRecognizeAnimalsRequest struct {
	VNImageBasedRequest
}

// VNRecognizeAnimalsRequestFromID constructs a [VNRecognizeAnimalsRequest] from an objc.ID.
//
// A request that recognizes animals in an image.
func VNRecognizeAnimalsRequestFromID(id objc.ID) VNRecognizeAnimalsRequest {
	return VNRecognizeAnimalsRequest{VNImageBasedRequest: VNImageBasedRequestFromID(id)}
}
// NOTE: VNRecognizeAnimalsRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNRecognizeAnimalsRequest] class.
//
// # Identifying Animals
//
//   - [IVNRecognizeAnimalsRequest.SupportedIdentifiersAndReturnError]: Returns the identifiers of the animals that the request detects.
//
// # Identifying Request Revisions
//
//   - [IVNRecognizeAnimalsRequest.VNRecognizeAnimalsRequestRevision2]: A constant for specifying revision 2 of the animal recognition request.
//   - [IVNRecognizeAnimalsRequest.VNRecognizeAnimalsRequestRevision1]: A constant for specifying revision 1 of the animal recognition request.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeAnimalsRequest
type IVNRecognizeAnimalsRequest interface {
	IVNImageBasedRequest

	// Topic: Identifying Animals

	// Returns the identifiers of the animals that the request detects.
	SupportedIdentifiersAndReturnError() ([]string, error)

	// Topic: Identifying Request Revisions

	// A constant for specifying revision 2 of the animal recognition request.
	VNRecognizeAnimalsRequestRevision2() int
	// A constant for specifying revision 1 of the animal recognition request.
	VNRecognizeAnimalsRequestRevision1() int
}

// Init initializes the instance.
func (r VNRecognizeAnimalsRequest) Init() VNRecognizeAnimalsRequest {
	rv := objc.Send[VNRecognizeAnimalsRequest](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r VNRecognizeAnimalsRequest) Autorelease() VNRecognizeAnimalsRequest {
	rv := objc.Send[VNRecognizeAnimalsRequest](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNRecognizeAnimalsRequest creates a new VNRecognizeAnimalsRequest instance.
func NewVNRecognizeAnimalsRequest() VNRecognizeAnimalsRequest {
	class := getVNRecognizeAnimalsRequestClass()
	rv := objc.Send[VNRecognizeAnimalsRequest](objc.ID(class.class), objc.Sel("new"))
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
func NewRecognizeAnimalsRequestWithCompletionHandler(completionHandler VNRequestCompletionHandler) VNRecognizeAnimalsRequest {
	instance := getVNRecognizeAnimalsRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return VNRecognizeAnimalsRequestFromID(rv)
}

// Returns the identifiers of the animals that the request detects.
//
// # Return Value
// 
// The animal identifiers.
//
// See: https://developer.apple.com/documentation/Vision/VNRecognizeAnimalsRequest/supportedIdentifiers()
func (r VNRecognizeAnimalsRequest) SupportedIdentifiersAndReturnError() ([]string, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("supportedIdentifiersAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSliceToStrings(rv), nil

}

// A constant for specifying revision 2 of the animal recognition request.
//
// See: https://developer.apple.com/documentation/vision/vnrecognizeanimalsrequestrevision2
func (r VNRecognizeAnimalsRequest) VNRecognizeAnimalsRequestRevision2() int {
	rv := objc.Send[int](r.ID, objc.Sel("VNRecognizeAnimalsRequestRevision2"))
	return rv
}
// A constant for specifying revision 1 of the animal recognition request.
//
// See: https://developer.apple.com/documentation/vision/vnrecognizeanimalsrequestrevision1
func (r VNRecognizeAnimalsRequest) VNRecognizeAnimalsRequestRevision1() int {
	rv := objc.Send[int](r.ID, objc.Sel("VNRecognizeAnimalsRequestRevision1"))
	return rv
}

