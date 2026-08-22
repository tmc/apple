// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VNTrackRectangleRequest] class.
var (
	_VNTrackRectangleRequestClass     VNTrackRectangleRequestClass
	_VNTrackRectangleRequestClassOnce sync.Once
)

func getVNTrackRectangleRequestClass() VNTrackRectangleRequestClass {
	_VNTrackRectangleRequestClassOnce.Do(func() {
		_VNTrackRectangleRequestClass = VNTrackRectangleRequestClass{class: objc.GetClass("VNTrackRectangleRequest")}
	})
	return _VNTrackRectangleRequestClass
}

// GetVNTrackRectangleRequestClass returns the class object for VNTrackRectangleRequest.
func GetVNTrackRectangleRequestClass() VNTrackRectangleRequestClass {
	return getVNTrackRectangleRequestClass()
}

type VNTrackRectangleRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNTrackRectangleRequestClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNTrackRectangleRequestClass) Alloc() VNTrackRectangleRequest {
	rv := objc.Send[VNTrackRectangleRequest](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An image-analysis request that tracks movement of a previously identified
// rectangular object across multiple images or video frames.
//
// # Overview
//
// Use this type of request to track the bounding boxes of rectangles
// throughout a sequence of images. Vision returns locations for rectangles
// found in all orientations and sizes.
//
// # Initializing a Rectangle Tracking Request
//
//   - [VNTrackRectangleRequest.InitWithRectangleObservation]: Creates a new rectangle tracking request with a rectangle observation.
//   - [VNTrackRectangleRequest.InitWithRectangleObservationCompletionHandler]: Creates a new rectangle tracking request with a rectangle observation.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest
type VNTrackRectangleRequest struct {
	VNTrackingRequest
}

// VNTrackRectangleRequestFromID constructs a [VNTrackRectangleRequest] from an objc.ID.
//
// An image-analysis request that tracks movement of a previously identified
// rectangular object across multiple images or video frames.
func VNTrackRectangleRequestFromID(id objc.ID) VNTrackRectangleRequest {
	return VNTrackRectangleRequest{VNTrackingRequest: VNTrackingRequestFromID(id)}
}

// NOTE: VNTrackRectangleRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNTrackRectangleRequest] class.
//
// # Initializing a Rectangle Tracking Request
//
//   - [IVNTrackRectangleRequest.InitWithRectangleObservation]: Creates a new rectangle tracking request with a rectangle observation.
//   - [IVNTrackRectangleRequest.InitWithRectangleObservationCompletionHandler]: Creates a new rectangle tracking request with a rectangle observation.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest
type IVNTrackRectangleRequest interface {
	IVNTrackingRequest

	// Topic: Initializing a Rectangle Tracking Request

	// Creates a new rectangle tracking request with a rectangle observation.
	InitWithRectangleObservation(observation IVNRectangleObservation) VNTrackRectangleRequest
	// Creates a new rectangle tracking request with a rectangle observation.
	InitWithRectangleObservationCompletionHandler(observation IVNRectangleObservation, completionHandler VNRequestErrorHandler) VNTrackRectangleRequest
}

// Init initializes the instance.
func (t VNTrackRectangleRequest) Init() VNTrackRectangleRequest {
	rv := objc.Send[VNTrackRectangleRequest](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t VNTrackRectangleRequest) Autorelease() VNTrackRectangleRequest {
	rv := objc.Send[VNTrackRectangleRequest](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNTrackRectangleRequest creates a new VNTrackRectangleRequest instance.
func NewVNTrackRectangleRequest() VNTrackRectangleRequest {
	class := getVNTrackRectangleRequestClass()
	rv := objc.Send[VNTrackRectangleRequest](objc.ID(class.class), objc.Sel("new"))
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
// [VNImageRequestHandler.PerformRequestsError].
//
// See: https://developer.apple.com/documentation/Vision/VNRequest/init(completionHandler:)
func NewTrackRectangleRequestWithCompletionHandler(completionHandler VNRequestErrorHandler) VNTrackRectangleRequest {
	_block0, _ := NewVNRequestErrorBlock(completionHandler)
	instance := getVNTrackRectangleRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompletionHandler:"), _block0)
	return VNTrackRectangleRequestFromID(rv)
}

// Creates a new rectangle tracking request with a rectangle observation.
//
// observation: A rectangle observation with bounding box and corner location information.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest/init(rectangleObservation:)
func NewTrackRectangleRequestWithRectangleObservation(observation IVNRectangleObservation) VNTrackRectangleRequest {
	instance := getVNTrackRectangleRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRectangleObservation:"), observation)
	return VNTrackRectangleRequestFromID(rv)
}

// Creates a new rectangle tracking request with a rectangle observation.
//
// observation: A rectangle observation with bounding box and corner location information.
//
// completionHandler: The block to invoke after performing the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest/init(rectangleObservation:completionHandler:)
func NewTrackRectangleRequestWithRectangleObservationCompletionHandler(observation IVNRectangleObservation, completionHandler VNRequestErrorHandler) VNTrackRectangleRequest {
	_block1, _ := NewVNRequestErrorBlock(completionHandler)
	instance := getVNTrackRectangleRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRectangleObservation:completionHandler:"), observation, _block1)
	return VNTrackRectangleRequestFromID(rv)
}

// Creates a new rectangle tracking request with a rectangle observation.
//
// observation: A rectangle observation with bounding box and corner location information.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest/init(rectangleObservation:)
func (t VNTrackRectangleRequest) InitWithRectangleObservation(observation IVNRectangleObservation) VNTrackRectangleRequest {
	rv := objc.Send[VNTrackRectangleRequest](t.ID, objc.Sel("initWithRectangleObservation:"), observation)
	return rv
}

// Creates a new rectangle tracking request with a rectangle observation.
//
// observation: A rectangle observation with bounding box and corner location information.
//
// completionHandler: The block to invoke after performing the request.
//
// See: https://developer.apple.com/documentation/Vision/VNTrackRectangleRequest/init(rectangleObservation:completionHandler:)
func (t VNTrackRectangleRequest) InitWithRectangleObservationCompletionHandler(observation IVNRectangleObservation, completionHandler VNRequestErrorHandler) VNTrackRectangleRequest {
	_block1, _ := NewVNRequestErrorBlock(completionHandler)
	rv := objc.Send[VNTrackRectangleRequest](t.ID, objc.Sel("initWithRectangleObservation:completionHandler:"), observation, _block1)
	return rv
}

// InitWithRectangleObservationSync is a synchronous wrapper around [VNTrackRectangleRequest.InitWithRectangleObservationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t VNTrackRectangleRequest) InitWithRectangleObservationSync(ctx context.Context, observation IVNRectangleObservation) (*VNRequest, error) {
	type result struct {
		val *VNRequest
		err error
	}
	done := make(chan result, 1)
	t.InitWithRectangleObservationCompletionHandler(observation, func(val *VNRequest, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
