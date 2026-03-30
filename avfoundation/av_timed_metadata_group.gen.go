// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVTimedMetadataGroup] class.
var (
	_AVTimedMetadataGroupClass     AVTimedMetadataGroupClass
	_AVTimedMetadataGroupClassOnce sync.Once
)

func getAVTimedMetadataGroupClass() AVTimedMetadataGroupClass {
	_AVTimedMetadataGroupClassOnce.Do(func() {
		_AVTimedMetadataGroupClass = AVTimedMetadataGroupClass{class: objc.GetClass("AVTimedMetadataGroup")}
	})
	return _AVTimedMetadataGroupClass
}

// GetAVTimedMetadataGroupClass returns the class object for AVTimedMetadataGroup.
func GetAVTimedMetadataGroupClass() AVTimedMetadataGroupClass {
	return getAVTimedMetadataGroupClass()
}

type AVTimedMetadataGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVTimedMetadataGroupClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVTimedMetadataGroupClass) Alloc() AVTimedMetadataGroup {
	rv := objc.Send[AVTimedMetadataGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A collection of metadata items that are valid for use during a specific
// time range.
//
// # Overview
//
// For example, [AVTimedMetadataGroups] are used to represent chapters,
// optionally containing metadata items for chapter titles and chapter images.
//
// # Creating a timed metadata group
//
//   - [AVTimedMetadataGroup.InitWithItemsTimeRange]: Creates a timed metadata group initialized with the given metadata items.
//
// # Accessing group attributes
//
//   - [AVTimedMetadataGroup.TimeRange]: The time range for the timed metadata.
//
// # Creating a format description
//
//   - [AVTimedMetadataGroup.CopyFormatDescription]: Creates a format description based on the receiver’s items.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup
type AVTimedMetadataGroup struct {
	AVMetadataGroup
}

// AVTimedMetadataGroupFromID constructs a [AVTimedMetadataGroup] from an objc.ID.
//
// A collection of metadata items that are valid for use during a specific
// time range.
func AVTimedMetadataGroupFromID(id objc.ID) AVTimedMetadataGroup {
	return AVTimedMetadataGroup{AVMetadataGroup: AVMetadataGroupFromID(id)}
}

// NOTE: AVTimedMetadataGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVTimedMetadataGroup] class.
//
// # Creating a timed metadata group
//
//   - [IAVTimedMetadataGroup.InitWithItemsTimeRange]: Creates a timed metadata group initialized with the given metadata items.
//
// # Accessing group attributes
//
//   - [IAVTimedMetadataGroup.TimeRange]: The time range for the timed metadata.
//
// # Creating a format description
//
//   - [IAVTimedMetadataGroup.CopyFormatDescription]: Creates a format description based on the receiver’s items.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup
type IAVTimedMetadataGroup interface {
	IAVMetadataGroup

	// Topic: Creating a timed metadata group

	// Creates a timed metadata group initialized with the given metadata items.
	InitWithItemsTimeRange(items []AVMetadataItem, timeRange coremedia.CMTimeRange) AVTimedMetadataGroup

	// Topic: Accessing group attributes

	// The time range for the timed metadata.
	TimeRange() coremedia.CMTimeRange

	// Topic: Creating a format description

	// Creates a format description based on the receiver’s items.
	CopyFormatDescription() coremedia.CMFormatDescriptionRef
}

// Init initializes the instance.
func (t AVTimedMetadataGroup) Init() AVTimedMetadataGroup {
	rv := objc.Send[AVTimedMetadataGroup](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t AVTimedMetadataGroup) Autorelease() AVTimedMetadataGroup {
	rv := objc.Send[AVTimedMetadataGroup](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVTimedMetadataGroup creates a new AVTimedMetadataGroup instance.
func NewAVTimedMetadataGroup() AVTimedMetadataGroup {
	class := getAVTimedMetadataGroupClass()
	rv := objc.Send[AVTimedMetadataGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a timed metadata group initialized with the given metadata items.
//
// items: An array of [AVMetadataItem] objects.
//
// timeRange: The time range of the metadata contained in `items`.
//
// # Return Value
//
// A metadata group initialized with `items`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/init(items:timeRange:)
func NewTimedMetadataGroupWithItemsTimeRange(items []AVMetadataItem, timeRange coremedia.CMTimeRange) AVTimedMetadataGroup {
	instance := getAVTimedMetadataGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItems:timeRange:"), objectivec.IObjectSliceToNSArray(items), timeRange)
	return AVTimedMetadataGroupFromID(rv)
}

// Creates a timed metadata group initialized with the given metadata items.
//
// items: An array of [AVMetadataItem] objects.
//
// timeRange: The time range of the metadata contained in `items`.
//
// # Return Value
//
// A metadata group initialized with `items`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/init(items:timeRange:)
func (t AVTimedMetadataGroup) InitWithItemsTimeRange(items []AVMetadataItem, timeRange coremedia.CMTimeRange) AVTimedMetadataGroup {
	rv := objc.Send[AVTimedMetadataGroup](t.ID, objc.Sel("initWithItems:timeRange:"), objectivec.IObjectSliceToNSArray(items), timeRange)
	return rv
}

// Creates a format description based on the receiver’s items.
//
// # Return Value
//
// An instance of [CMMetadataFormatDescription] sufficient to describe the
// contents of all the items referenced by the object.
//
// # Discussion
//
// The returned format description is suitable for use as the format hint
// parameter when creating an instance of [AVAssetWriterInput].
//
// Each item referenced by the receiver must carry a non-`nil` value for its
// `dataType` property. An exception will be thrown if any item does not have
// a data type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/copyFormatDescription()
//
// [CMMetadataFormatDescription]: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescription
func (t AVTimedMetadataGroup) CopyFormatDescription() coremedia.CMFormatDescriptionRef {
	rv := objc.Send[coremedia.CMFormatDescriptionRef](t.ID, objc.Sel("copyFormatDescription"))
	return coremedia.CMFormatDescriptionRef(rv)
}

// The time range for the timed metadata.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/timeRange
func (t AVTimedMetadataGroup) TimeRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](t.ID, objc.Sel("timeRange"))
	return coremedia.CMTimeRange(rv)
}
