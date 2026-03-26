// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMutableDateRangeMetadataGroup] class.
var (
	_AVMutableDateRangeMetadataGroupClass     AVMutableDateRangeMetadataGroupClass
	_AVMutableDateRangeMetadataGroupClassOnce sync.Once
)

func getAVMutableDateRangeMetadataGroupClass() AVMutableDateRangeMetadataGroupClass {
	_AVMutableDateRangeMetadataGroupClassOnce.Do(func() {
		_AVMutableDateRangeMetadataGroupClass = AVMutableDateRangeMetadataGroupClass{class: objc.GetClass("AVMutableDateRangeMetadataGroup")}
	})
	return _AVMutableDateRangeMetadataGroupClass
}

// GetAVMutableDateRangeMetadataGroupClass returns the class object for AVMutableDateRangeMetadataGroup.
func GetAVMutableDateRangeMetadataGroupClass() AVMutableDateRangeMetadataGroupClass {
	return getAVMutableDateRangeMetadataGroupClass()
}

type AVMutableDateRangeMetadataGroupClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableDateRangeMetadataGroupClass) Alloc() AVMutableDateRangeMetadataGroup {
	rv := objc.Send[AVMutableDateRangeMetadataGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable collection of metadata items that are valid for use within a
// specific range of dates.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup
type AVMutableDateRangeMetadataGroup struct {
	AVDateRangeMetadataGroup
}

// AVMutableDateRangeMetadataGroupFromID constructs a [AVMutableDateRangeMetadataGroup] from an objc.ID.
//
// A mutable collection of metadata items that are valid for use within a
// specific range of dates.
func AVMutableDateRangeMetadataGroupFromID(id objc.ID) AVMutableDateRangeMetadataGroup {
	return AVMutableDateRangeMetadataGroup{AVDateRangeMetadataGroup: AVDateRangeMetadataGroupFromID(id)}
}
// NOTE: AVMutableDateRangeMetadataGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableDateRangeMetadataGroup] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup
type IAVMutableDateRangeMetadataGroup interface {
	IAVDateRangeMetadataGroup
}

// Init initializes the instance.
func (m AVMutableDateRangeMetadataGroup) Init() AVMutableDateRangeMetadataGroup {
	rv := objc.Send[AVMutableDateRangeMetadataGroup](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableDateRangeMetadataGroup) Autorelease() AVMutableDateRangeMetadataGroup {
	rv := objc.Send[AVMutableDateRangeMetadataGroup](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableDateRangeMetadataGroup creates a new AVMutableDateRangeMetadataGroup instance.
func NewAVMutableDateRangeMetadataGroup() AVMutableDateRangeMetadataGroup {
	class := getAVMutableDateRangeMetadataGroupClass()
	rv := objc.Send[AVMutableDateRangeMetadataGroup](objc.ID(class.class), objc.Sel("new"))
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
func NewMutableDateRangeMetadataGroupWithItemsStartDateEndDate(items []AVMetadataItem, startDate foundation.INSDate, endDate foundation.INSDate) AVMutableDateRangeMetadataGroup {
	instance := getAVMutableDateRangeMetadataGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItems:startDate:endDate:"), objectivec.IObjectSliceToNSArray(items), startDate, endDate)
	return AVMutableDateRangeMetadataGroupFromID(rv)
}

