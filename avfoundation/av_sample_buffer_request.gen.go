// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSampleBufferRequest] class.
var (
	_AVSampleBufferRequestClass     AVSampleBufferRequestClass
	_AVSampleBufferRequestClassOnce sync.Once
)

func getAVSampleBufferRequestClass() AVSampleBufferRequestClass {
	_AVSampleBufferRequestClassOnce.Do(func() {
		_AVSampleBufferRequestClass = AVSampleBufferRequestClass{class: objc.GetClass("AVSampleBufferRequest")}
	})
	return _AVSampleBufferRequestClass
}

// GetAVSampleBufferRequestClass returns the class object for AVSampleBufferRequest.
func GetAVSampleBufferRequestClass() AVSampleBufferRequestClass {
	return getAVSampleBufferRequestClass()
}

type AVSampleBufferRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSampleBufferRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSampleBufferRequestClass) Alloc() AVSampleBufferRequest {
	rv := objc.Send[AVSampleBufferRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a sample buffer creation request.
//
// # Creating a request
//
//   - [AVSampleBufferRequest.InitWithStartCursor]: Creates a newly allocated sample buffer request with the specified sample cursor.
//
// # Configuring sample buffer request parameters
//
//   - [AVSampleBufferRequest.Direction]: The buffer sample direction.
//   - [AVSampleBufferRequest.SetDirection]
//   - [AVSampleBufferRequest.LimitCursor]: The limiting position for sample loading.
//   - [AVSampleBufferRequest.SetLimitCursor]
//   - [AVSampleBufferRequest.MaxSampleCount]: The maximum number of samples to load.
//   - [AVSampleBufferRequest.SetMaxSampleCount]
//   - [AVSampleBufferRequest.Mode]: The sample buffer request mode.
//   - [AVSampleBufferRequest.SetMode]
//   - [AVSampleBufferRequest.OverrideTime]: The deadline for sample data and output PTS for the sample buffer.
//   - [AVSampleBufferRequest.SetOverrideTime]
//   - [AVSampleBufferRequest.PreferredMinSampleCount]: The preferred minimum number of samples to load.
//   - [AVSampleBufferRequest.SetPreferredMinSampleCount]
//   - [AVSampleBufferRequest.StartCursor]: The starting cursor position.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest
type AVSampleBufferRequest struct {
	objectivec.Object
}

// AVSampleBufferRequestFromID constructs a [AVSampleBufferRequest] from an objc.ID.
//
// An object that describes a sample buffer creation request.
func AVSampleBufferRequestFromID(id objc.ID) AVSampleBufferRequest {
	return AVSampleBufferRequest{objectivec.Object{ID: id}}
}

// NOTE: AVSampleBufferRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSampleBufferRequest] class.
//
// # Creating a request
//
//   - [IAVSampleBufferRequest.InitWithStartCursor]: Creates a newly allocated sample buffer request with the specified sample cursor.
//
// # Configuring sample buffer request parameters
//
//   - [IAVSampleBufferRequest.Direction]: The buffer sample direction.
//   - [IAVSampleBufferRequest.SetDirection]
//   - [IAVSampleBufferRequest.LimitCursor]: The limiting position for sample loading.
//   - [IAVSampleBufferRequest.SetLimitCursor]
//   - [IAVSampleBufferRequest.MaxSampleCount]: The maximum number of samples to load.
//   - [IAVSampleBufferRequest.SetMaxSampleCount]
//   - [IAVSampleBufferRequest.Mode]: The sample buffer request mode.
//   - [IAVSampleBufferRequest.SetMode]
//   - [IAVSampleBufferRequest.OverrideTime]: The deadline for sample data and output PTS for the sample buffer.
//   - [IAVSampleBufferRequest.SetOverrideTime]
//   - [IAVSampleBufferRequest.PreferredMinSampleCount]: The preferred minimum number of samples to load.
//   - [IAVSampleBufferRequest.SetPreferredMinSampleCount]
//   - [IAVSampleBufferRequest.StartCursor]: The starting cursor position.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest
type IAVSampleBufferRequest interface {
	objectivec.IObject

	// Topic: Creating a request

	// Creates a newly allocated sample buffer request with the specified sample cursor.
	InitWithStartCursor(startCursor IAVSampleCursor) AVSampleBufferRequest

	// Topic: Configuring sample buffer request parameters

	// The buffer sample direction.
	Direction() AVSampleBufferRequestDirection
	SetDirection(value AVSampleBufferRequestDirection)
	// The limiting position for sample loading.
	LimitCursor() IAVSampleCursor
	SetLimitCursor(value IAVSampleCursor)
	// The maximum number of samples to load.
	MaxSampleCount() int
	SetMaxSampleCount(value int)
	// The sample buffer request mode.
	Mode() AVSampleBufferRequestMode
	SetMode(value AVSampleBufferRequestMode)
	// The deadline for sample data and output PTS for the sample buffer.
	OverrideTime() coremedia.CMTime
	SetOverrideTime(value coremedia.CMTime)
	// The preferred minimum number of samples to load.
	PreferredMinSampleCount() int
	SetPreferredMinSampleCount(value int)
	// The starting cursor position.
	StartCursor() IAVSampleCursor
}

// Init initializes the instance.
func (s AVSampleBufferRequest) Init() AVSampleBufferRequest {
	rv := objc.Send[AVSampleBufferRequest](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSampleBufferRequest) Autorelease() AVSampleBufferRequest {
	rv := objc.Send[AVSampleBufferRequest](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleBufferRequest creates a new AVSampleBufferRequest instance.
func NewAVSampleBufferRequest() AVSampleBufferRequest {
	class := getAVSampleBufferRequestClass()
	rv := objc.Send[AVSampleBufferRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a newly allocated sample buffer request with the specified sample
// cursor.
//
// startCursor: The starting cursor position.
//
// # Return Value
//
// An initialized [AVSampleBufferRequest] instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/init(start:)
func NewSampleBufferRequestWithStartCursor(startCursor IAVSampleCursor) AVSampleBufferRequest {
	instance := getAVSampleBufferRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStartCursor:"), startCursor)
	return AVSampleBufferRequestFromID(rv)
}

// Creates a newly allocated sample buffer request with the specified sample
// cursor.
//
// startCursor: The starting cursor position.
//
// # Return Value
//
// An initialized [AVSampleBufferRequest] instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/init(start:)
func (s AVSampleBufferRequest) InitWithStartCursor(startCursor IAVSampleCursor) AVSampleBufferRequest {
	rv := objc.Send[AVSampleBufferRequest](s.ID, objc.Sel("initWithStartCursor:"), startCursor)
	return rv
}

// The buffer sample direction.
//
// # Discussion
//
// The default value is [AVSampleBufferRequestDirectionNone].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/direction-swift.property
func (s AVSampleBufferRequest) Direction() AVSampleBufferRequestDirection {
	rv := objc.Send[AVSampleBufferRequestDirection](s.ID, objc.Sel("direction"))
	return AVSampleBufferRequestDirection(rv)
}
func (s AVSampleBufferRequest) SetDirection(value AVSampleBufferRequestDirection) {
	objc.Send[struct{}](s.ID, objc.Sel("setDirection:"), value)
}

// The limiting position for sample loading.
//
// # Discussion
//
// If the value isn’t `nil`, the sequence of samples to load may include the
// sample at this position, but no further.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/limitCursor
func (s AVSampleBufferRequest) LimitCursor() IAVSampleCursor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("limitCursor"))
	return AVSampleCursorFromID(objc.ID(rv))
}
func (s AVSampleBufferRequest) SetLimitCursor(value IAVSampleCursor) {
	objc.Send[struct{}](s.ID, objc.Sel("setLimitCursor:"), value)
}

// The maximum number of samples to load.
//
// # Discussion
//
// If the value isn’t `0`, indicates the maximum number of samples to load.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/maxSampleCount
func (s AVSampleBufferRequest) MaxSampleCount() int {
	rv := objc.Send[int](s.ID, objc.Sel("maxSampleCount"))
	return rv
}
func (s AVSampleBufferRequest) SetMaxSampleCount(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaxSampleCount:"), value)
}

// The sample buffer request mode.
//
// # Discussion
//
// Default is [AVSampleBufferRequestModeImmediate].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/mode-swift.property
func (s AVSampleBufferRequest) Mode() AVSampleBufferRequestMode {
	rv := objc.Send[AVSampleBufferRequestMode](s.ID, objc.Sel("mode"))
	return AVSampleBufferRequestMode(rv)
}
func (s AVSampleBufferRequest) SetMode(value AVSampleBufferRequestMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setMode:"), value)
}

// The deadline for sample data and output PTS for the sample buffer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/overrideTime
func (s AVSampleBufferRequest) OverrideTime() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](s.ID, objc.Sel("overrideTime"))
	return coremedia.CMTime(rv)
}
func (s AVSampleBufferRequest) SetOverrideTime(value coremedia.CMTime) {
	objc.Send[struct{}](s.ID, objc.Sel("setOverrideTime:"), value)
}

// The preferred minimum number of samples to load.
//
// # Discussion
//
// A value that isn’t `0` indicates the preferred number of samples to load.
// Fewer samples may be loaded if there is a change of format description.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/preferredMinSampleCount
func (s AVSampleBufferRequest) PreferredMinSampleCount() int {
	rv := objc.Send[int](s.ID, objc.Sel("preferredMinSampleCount"))
	return rv
}
func (s AVSampleBufferRequest) SetPreferredMinSampleCount(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreferredMinSampleCount:"), value)
}

// The starting cursor position.
//
// # Discussion
//
// The [CMSampleBuffer] created with the request must include the sample at
// this position.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/startCursor
//
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
func (s AVSampleBufferRequest) StartCursor() IAVSampleCursor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("startCursor"))
	return AVSampleCursorFromID(objc.ID(rv))
}
