// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/avfoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MEEstimatedSampleLocation] class.
var (
	_MEEstimatedSampleLocationClass     MEEstimatedSampleLocationClass
	_MEEstimatedSampleLocationClassOnce sync.Once
)

func getMEEstimatedSampleLocationClass() MEEstimatedSampleLocationClass {
	_MEEstimatedSampleLocationClassOnce.Do(func() {
		_MEEstimatedSampleLocationClass = MEEstimatedSampleLocationClass{class: objc.GetClass("MEEstimatedSampleLocation")}
	})
	return _MEEstimatedSampleLocationClass
}

// GetMEEstimatedSampleLocationClass returns the class object for MEEstimatedSampleLocation.
func GetMEEstimatedSampleLocationClass() MEEstimatedSampleLocationClass {
	return getMEEstimatedSampleLocationClass()
}

type MEEstimatedSampleLocationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MEEstimatedSampleLocationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MEEstimatedSampleLocationClass) Alloc() MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about the estimated sample location
// with the media.
//
// # Creating an estimated sample location
//
//   - [MEEstimatedSampleLocation.InitWithByteSourceEstimatedSampleLocationRefinementDataLocation]: Creates an estimated sample location object with the byte source, sample location, and data location that you specify.
//
// # Inspecting an estimated sample location
//
//   - [MEEstimatedSampleLocation.ByteSource]: The byte source to use to read the data for the sample.
//   - [MEEstimatedSampleLocation.EstimatedSampleLocation]: The estimated starting file offset and size in bytes of the sample.
//   - [MEEstimatedSampleLocation.RefinementDataLocation]: The starting file offset and size in bytes of the data necessary to provide an accurate sample location.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation
type MEEstimatedSampleLocation struct {
	objectivec.Object
}

// MEEstimatedSampleLocationFromID constructs a [MEEstimatedSampleLocation] from an objc.ID.
//
// An object that provides information about the estimated sample location
// with the media.
func MEEstimatedSampleLocationFromID(id objc.ID) MEEstimatedSampleLocation {
	return MEEstimatedSampleLocation{objectivec.Object{ID: id}}
}

// NOTE: MEEstimatedSampleLocation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MEEstimatedSampleLocation] class.
//
// # Creating an estimated sample location
//
//   - [IMEEstimatedSampleLocation.InitWithByteSourceEstimatedSampleLocationRefinementDataLocation]: Creates an estimated sample location object with the byte source, sample location, and data location that you specify.
//
// # Inspecting an estimated sample location
//
//   - [IMEEstimatedSampleLocation.ByteSource]: The byte source to use to read the data for the sample.
//   - [IMEEstimatedSampleLocation.EstimatedSampleLocation]: The estimated starting file offset and size in bytes of the sample.
//   - [IMEEstimatedSampleLocation.RefinementDataLocation]: The starting file offset and size in bytes of the data necessary to provide an accurate sample location.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation
type IMEEstimatedSampleLocation interface {
	objectivec.IObject

	// Topic: Creating an estimated sample location

	// Creates an estimated sample location object with the byte source, sample location, and data location that you specify.
	InitWithByteSourceEstimatedSampleLocationRefinementDataLocation(byteSource IMEByteSource, estimatedSampleLocation avfoundation.AVSampleCursorStorageRange, refinementDataLocation avfoundation.AVSampleCursorStorageRange) MEEstimatedSampleLocation

	// Topic: Inspecting an estimated sample location

	// The byte source to use to read the data for the sample.
	ByteSource() IMEByteSource
	// The estimated starting file offset and size in bytes of the sample.
	EstimatedSampleLocation() avfoundation.AVSampleCursorStorageRange
	// The starting file offset and size in bytes of the data necessary to provide an accurate sample location.
	RefinementDataLocation() avfoundation.AVSampleCursorStorageRange
}

// Init initializes the instance.
func (m MEEstimatedSampleLocation) Init() MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MEEstimatedSampleLocation) Autorelease() MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEEstimatedSampleLocation creates a new MEEstimatedSampleLocation instance.
func NewMEEstimatedSampleLocation() MEEstimatedSampleLocation {
	class := getMEEstimatedSampleLocationClass()
	rv := objc.Send[MEEstimatedSampleLocation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an estimated sample location object with the byte source, sample
// location, and data location that you specify.
//
// byteSource: The byte source to use to read the data for the sample.
//
// estimatedSampleLocation: The estimated starting file offset and size in bytes of the sample.
//
// refinementDataLocation: The starting file offset and size in bytes of the data necessary to provide
// an accurate sample location.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/init(byteSource:estimatedSampleLocation:refinementDataLocation:)
func NewMEEstimatedSampleLocationWithByteSourceEstimatedSampleLocationRefinementDataLocation(byteSource IMEByteSource, estimatedSampleLocation avfoundation.AVSampleCursorStorageRange, refinementDataLocation avfoundation.AVSampleCursorStorageRange) MEEstimatedSampleLocation {
	instance := getMEEstimatedSampleLocationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithByteSource:estimatedSampleLocation:refinementDataLocation:"), byteSource, estimatedSampleLocation, refinementDataLocation)
	return MEEstimatedSampleLocationFromID(rv)
}

// Creates an estimated sample location object with the byte source, sample
// location, and data location that you specify.
//
// byteSource: The byte source to use to read the data for the sample.
//
// estimatedSampleLocation: The estimated starting file offset and size in bytes of the sample.
//
// refinementDataLocation: The starting file offset and size in bytes of the data necessary to provide
// an accurate sample location.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/init(byteSource:estimatedSampleLocation:refinementDataLocation:)
func (m MEEstimatedSampleLocation) InitWithByteSourceEstimatedSampleLocationRefinementDataLocation(byteSource IMEByteSource, estimatedSampleLocation avfoundation.AVSampleCursorStorageRange, refinementDataLocation avfoundation.AVSampleCursorStorageRange) MEEstimatedSampleLocation {
	rv := objc.Send[MEEstimatedSampleLocation](m.ID, objc.Sel("initWithByteSource:estimatedSampleLocation:refinementDataLocation:"), byteSource, estimatedSampleLocation, refinementDataLocation)
	return rv
}

// The byte source to use to read the data for the sample.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/byteSource
func (m MEEstimatedSampleLocation) ByteSource() IMEByteSource {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("byteSource"))
	return MEByteSourceFromID(objc.ID(rv))
}

// The estimated starting file offset and size in bytes of the sample.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/estimatedSampleLocation
func (m MEEstimatedSampleLocation) EstimatedSampleLocation() avfoundation.AVSampleCursorStorageRange {
	rv := objc.Send[avfoundation.AVSampleCursorStorageRange](m.ID, objc.Sel("estimatedSampleLocation"))
	return avfoundation.AVSampleCursorStorageRange(rv)
}

// The starting file offset and size in bytes of the data necessary to provide
// an accurate sample location.
//
// # Discussion
//
// Pass this refinement data to the
// [RefineSampleLocationRefinementDataRefinementDataLengthRefinedLocationError]
// to determine the exact sample location.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEEstimatedSampleLocation/refinementDataLocation
func (m MEEstimatedSampleLocation) RefinementDataLocation() avfoundation.AVSampleCursorStorageRange {
	rv := objc.Send[avfoundation.AVSampleCursorStorageRange](m.ID, objc.Sel("refinementDataLocation"))
	return avfoundation.AVSampleCursorStorageRange(rv)
}
