// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/avfoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MESampleLocation] class.
var (
	_MESampleLocationClass     MESampleLocationClass
	_MESampleLocationClassOnce sync.Once
)

func getMESampleLocationClass() MESampleLocationClass {
	_MESampleLocationClassOnce.Do(func() {
		_MESampleLocationClass = MESampleLocationClass{class: objc.GetClass("MESampleLocation")}
	})
	return _MESampleLocationClass
}

// GetMESampleLocationClass returns the class object for MESampleLocation.
func GetMESampleLocationClass() MESampleLocationClass {
	return getMESampleLocationClass()
}

type MESampleLocationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MESampleLocationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MESampleLocationClass) Alloc() MESampleLocation {
	rv := objc.Send[MESampleLocation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about the sample location with the
// media.
//
// # Creating a sample location
//
//   - [MESampleLocation.InitWithByteSourceSampleLocation]: Creates a sample location object with the byte source and sample location that you specify.
//
// # Inspecting a sample location
//
//   - [MESampleLocation.ByteSource]: The byte source to use to read the data for the sample.
//   - [MESampleLocation.SampleLocation]: The starting file offset and size in bytes of the sample.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleLocation
type MESampleLocation struct {
	objectivec.Object
}

// MESampleLocationFromID constructs a [MESampleLocation] from an objc.ID.
//
// An object that provides information about the sample location with the
// media.
func MESampleLocationFromID(id objc.ID) MESampleLocation {
	return MESampleLocation{objectivec.Object{ID: id}}
}

// NOTE: MESampleLocation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MESampleLocation] class.
//
// # Creating a sample location
//
//   - [IMESampleLocation.InitWithByteSourceSampleLocation]: Creates a sample location object with the byte source and sample location that you specify.
//
// # Inspecting a sample location
//
//   - [IMESampleLocation.ByteSource]: The byte source to use to read the data for the sample.
//   - [IMESampleLocation.SampleLocation]: The starting file offset and size in bytes of the sample.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleLocation
type IMESampleLocation interface {
	objectivec.IObject

	// Topic: Creating a sample location

	// Creates a sample location object with the byte source and sample location that you specify.
	InitWithByteSourceSampleLocation(byteSource IMEByteSource, sampleLocation avfoundation.AVSampleCursorStorageRange) MESampleLocation

	// Topic: Inspecting a sample location

	// The byte source to use to read the data for the sample.
	ByteSource() IMEByteSource
	// The starting file offset and size in bytes of the sample.
	SampleLocation() avfoundation.AVSampleCursorStorageRange
}

// Init initializes the instance.
func (m MESampleLocation) Init() MESampleLocation {
	rv := objc.Send[MESampleLocation](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MESampleLocation) Autorelease() MESampleLocation {
	rv := objc.Send[MESampleLocation](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMESampleLocation creates a new MESampleLocation instance.
func NewMESampleLocation() MESampleLocation {
	class := getMESampleLocationClass()
	rv := objc.Send[MESampleLocation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a sample location object with the byte source and sample location
// that you specify.
//
// byteSource: The byte source to use to read the data for the sample.
//
// sampleLocation: The starting file offset and size in bytes of the sample.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleLocation/init(byteSource:sampleLocation:)
func NewMESampleLocationWithByteSourceSampleLocation(byteSource IMEByteSource, sampleLocation avfoundation.AVSampleCursorStorageRange) MESampleLocation {
	instance := getMESampleLocationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithByteSource:sampleLocation:"), byteSource, sampleLocation)
	return MESampleLocationFromID(rv)
}

// Creates a sample location object with the byte source and sample location
// that you specify.
//
// byteSource: The byte source to use to read the data for the sample.
//
// sampleLocation: The starting file offset and size in bytes of the sample.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleLocation/init(byteSource:sampleLocation:)
func (m MESampleLocation) InitWithByteSourceSampleLocation(byteSource IMEByteSource, sampleLocation avfoundation.AVSampleCursorStorageRange) MESampleLocation {
	rv := objc.Send[MESampleLocation](m.ID, objc.Sel("initWithByteSource:sampleLocation:"), byteSource, sampleLocation)
	return rv
}

// The byte source to use to read the data for the sample.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleLocation/byteSource
func (m MESampleLocation) ByteSource() IMEByteSource {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("byteSource"))
	return MEByteSourceFromID(objc.ID(rv))
}

// The starting file offset and size in bytes of the sample.
//
// See: https://developer.apple.com/documentation/MediaExtension/MESampleLocation/sampleLocation
func (m MESampleLocation) SampleLocation() avfoundation.AVSampleCursorStorageRange {
	rv := objc.Send[avfoundation.AVSampleCursorStorageRange](m.ID, objc.Sel("sampleLocation"))
	return avfoundation.AVSampleCursorStorageRange(rv)
}
