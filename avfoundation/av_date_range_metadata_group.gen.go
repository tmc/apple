// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVDateRangeMetadataGroup] class.
var (
	_AVDateRangeMetadataGroupClass     AVDateRangeMetadataGroupClass
	_AVDateRangeMetadataGroupClassOnce sync.Once
)

func getAVDateRangeMetadataGroupClass() AVDateRangeMetadataGroupClass {
	_AVDateRangeMetadataGroupClassOnce.Do(func() {
		_AVDateRangeMetadataGroupClass = AVDateRangeMetadataGroupClass{class: objc.GetClass("AVDateRangeMetadataGroup")}
	})
	return _AVDateRangeMetadataGroupClass
}

// GetAVDateRangeMetadataGroupClass returns the class object for AVDateRangeMetadataGroup.
func GetAVDateRangeMetadataGroupClass() AVDateRangeMetadataGroupClass {
	return getAVDateRangeMetadataGroupClass()
}

type AVDateRangeMetadataGroupClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVDateRangeMetadataGroupClass) Alloc() AVDateRangeMetadataGroup {
	rv := objc.Send[AVDateRangeMetadataGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A collection of metadata items that are valid for use within a specific
// date range.
//
// # Creating a date range group
//
//   - [AVDateRangeMetadataGroup.InitWithItemsStartDateEndDate]: Initializes an instance of [AVDateRangeMetadataGroup] with a collection of metadata items.
//
// # Accessing the date range
//
//   - [AVDateRangeMetadataGroup.StartDate]: The start date for the metadata date range group.
//   - [AVDateRangeMetadataGroup.EndDate]: The end date for the metadata date range group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup
type AVDateRangeMetadataGroup struct {
	AVMetadataGroup
}

// AVDateRangeMetadataGroupFromID constructs a [AVDateRangeMetadataGroup] from an objc.ID.
//
// A collection of metadata items that are valid for use within a specific
// date range.
func AVDateRangeMetadataGroupFromID(id objc.ID) AVDateRangeMetadataGroup {
	return AVDateRangeMetadataGroup{AVMetadataGroup: AVMetadataGroupFromID(id)}
}
// NOTE: AVDateRangeMetadataGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVDateRangeMetadataGroup] class.
//
// # Creating a date range group
//
//   - [IAVDateRangeMetadataGroup.InitWithItemsStartDateEndDate]: Initializes an instance of [AVDateRangeMetadataGroup] with a collection of metadata items.
//
// # Accessing the date range
//
//   - [IAVDateRangeMetadataGroup.StartDate]: The start date for the metadata date range group.
//   - [IAVDateRangeMetadataGroup.EndDate]: The end date for the metadata date range group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup
type IAVDateRangeMetadataGroup interface {
	IAVMetadataGroup

	// Topic: Creating a date range group

	// Initializes an instance of [AVDateRangeMetadataGroup] with a collection of metadata items.
	InitWithItemsStartDateEndDate(items []AVMetadataItem, startDate foundation.INSDate, endDate foundation.INSDate) AVDateRangeMetadataGroup

	// Topic: Accessing the date range

	// The start date for the metadata date range group.
	StartDate() foundation.INSDate
	// The end date for the metadata date range group.
	EndDate() foundation.INSDate
}

// Init initializes the instance.
func (d AVDateRangeMetadataGroup) Init() AVDateRangeMetadataGroup {
	rv := objc.Send[AVDateRangeMetadataGroup](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d AVDateRangeMetadataGroup) Autorelease() AVDateRangeMetadataGroup {
	rv := objc.Send[AVDateRangeMetadataGroup](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVDateRangeMetadataGroup creates a new AVDateRangeMetadataGroup instance.
func NewAVDateRangeMetadataGroup() AVDateRangeMetadataGroup {
	class := getAVDateRangeMetadataGroupClass()
	rv := objc.Send[AVDateRangeMetadataGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an instance of [AVDateRangeMetadataGroup] with a collection of
// metadata items.
//
// items: The array of [AVMetadataItem] instances to associate with this group.
//
// startDate: The starting date for the group of metadata items.
//
// endDate: The ending date for the group of metadata items.
//
// # Return Value
// 
// A new instance of [AVDateRangeMetadataGroup].
//
// # Discussion
// 
// Creates a new instance of [AVDateRangeMetadataGroup] with the specified
// collection of metadata items. The `startDate` and `endDate` arguments
// define the effective time range on the timeline to which the metadata
// applies.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup/init(items:start:end:)
func NewDateRangeMetadataGroupWithItemsStartDateEndDate(items []AVMetadataItem, startDate foundation.INSDate, endDate foundation.INSDate) AVDateRangeMetadataGroup {
	instance := getAVDateRangeMetadataGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItems:startDate:endDate:"), objectivec.IObjectSliceToNSArray(items), startDate, endDate)
	return AVDateRangeMetadataGroupFromID(rv)
}

// Initializes an instance of [AVDateRangeMetadataGroup] with a collection of
// metadata items.
//
// items: The array of [AVMetadataItem] instances to associate with this group.
//
// startDate: The starting date for the group of metadata items.
//
// endDate: The ending date for the group of metadata items.
//
// # Return Value
// 
// A new instance of [AVDateRangeMetadataGroup].
//
// # Discussion
// 
// Creates a new instance of [AVDateRangeMetadataGroup] with the specified
// collection of metadata items. The `startDate` and `endDate` arguments
// define the effective time range on the timeline to which the metadata
// applies.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup/init(items:start:end:)
func (d AVDateRangeMetadataGroup) InitWithItemsStartDateEndDate(items []AVMetadataItem, startDate foundation.INSDate, endDate foundation.INSDate) AVDateRangeMetadataGroup {
	rv := objc.Send[AVDateRangeMetadataGroup](d.ID, objc.Sel("initWithItems:startDate:endDate:"), objectivec.IObjectSliceToNSArray(items), startDate, endDate)
	return rv
}

// The start date for the metadata date range group.
//
// # Discussion
// 
// The `startDate` defines the starting date on the timeline for which the
// associated metadata is valid.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup/startDate
func (d AVDateRangeMetadataGroup) StartDate() foundation.INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("startDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
// The end date for the metadata date range group.
//
// # Discussion
// 
// The `endDate` defines the ending date on the timeline for which the
// associated metadata is valid.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup/endDate
func (d AVDateRangeMetadataGroup) EndDate() foundation.INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("endDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

