// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
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
// See: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup
type IAVTimedMetadataGroup interface {
	IAVMetadataGroup

	// Topic: Creating a timed metadata group

	// Creates a timed metadata group initialized with the given metadata items.
	InitWithItemsTimeRange(items []AVMetadataItem, timeRange objectivec.IObject) AVTimedMetadataGroup

	// Topic: Accessing group attributes

	// The time range for the timed metadata.
	TimeRange() objectivec.IObject
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
// timeRange is a [coremedia.CMTimeRange].
func NewTimedMetadataGroupWithItemsTimeRange(items []AVMetadataItem, timeRange objectivec.IObject) AVTimedMetadataGroup {
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
// timeRange is a [coremedia.CMTimeRange].
//
// # Return Value
// 
// A metadata group initialized with `items`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/init(items:timeRange:)
func (t AVTimedMetadataGroup) InitWithItemsTimeRange(items []AVMetadataItem, timeRange objectivec.IObject) AVTimedMetadataGroup {
	rv := objc.Send[AVTimedMetadataGroup](t.ID, objc.Sel("initWithItems:timeRange:"), objectivec.IObjectSliceToNSArray(items), timeRange)
	return rv
}

// The time range for the timed metadata.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/timeRange
func (t AVTimedMetadataGroup) TimeRange() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("timeRange"))
	return objectivec.Object{ID: rv}
}

