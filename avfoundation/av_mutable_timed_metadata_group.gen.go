// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMutableTimedMetadataGroup] class.
var (
	_AVMutableTimedMetadataGroupClass     AVMutableTimedMetadataGroupClass
	_AVMutableTimedMetadataGroupClassOnce sync.Once
)

func getAVMutableTimedMetadataGroupClass() AVMutableTimedMetadataGroupClass {
	_AVMutableTimedMetadataGroupClassOnce.Do(func() {
		_AVMutableTimedMetadataGroupClass = AVMutableTimedMetadataGroupClass{class: objc.GetClass("AVMutableTimedMetadataGroup")}
	})
	return _AVMutableTimedMetadataGroupClass
}

// GetAVMutableTimedMetadataGroupClass returns the class object for AVMutableTimedMetadataGroup.
func GetAVMutableTimedMetadataGroupClass() AVMutableTimedMetadataGroupClass {
	return getAVMutableTimedMetadataGroupClass()
}

type AVMutableTimedMetadataGroupClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableTimedMetadataGroupClass) Alloc() AVMutableTimedMetadataGroup {
	rv := objc.Send[AVMutableTimedMetadataGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable collection of metadata items that are valid for use during a
// specific time range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableTimedMetadataGroup
type AVMutableTimedMetadataGroup struct {
	AVTimedMetadataGroup
}

// AVMutableTimedMetadataGroupFromID constructs a [AVMutableTimedMetadataGroup] from an objc.ID.
//
// A mutable collection of metadata items that are valid for use during a
// specific time range.
func AVMutableTimedMetadataGroupFromID(id objc.ID) AVMutableTimedMetadataGroup {
	return AVMutableTimedMetadataGroup{AVTimedMetadataGroup: AVTimedMetadataGroupFromID(id)}
}
// NOTE: AVMutableTimedMetadataGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableTimedMetadataGroup] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableTimedMetadataGroup
type IAVMutableTimedMetadataGroup interface {
	IAVTimedMetadataGroup
}

// Init initializes the instance.
func (m AVMutableTimedMetadataGroup) Init() AVMutableTimedMetadataGroup {
	rv := objc.Send[AVMutableTimedMetadataGroup](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableTimedMetadataGroup) Autorelease() AVMutableTimedMetadataGroup {
	rv := objc.Send[AVMutableTimedMetadataGroup](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableTimedMetadataGroup creates a new AVMutableTimedMetadataGroup instance.
func NewAVMutableTimedMetadataGroup() AVMutableTimedMetadataGroup {
	class := getAVMutableTimedMetadataGroupClass()
	rv := objc.Send[AVMutableTimedMetadataGroup](objc.ID(class.class), objc.Sel("new"))
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
func NewMutableTimedMetadataGroupWithItemsTimeRange(items []AVMetadataItem, timeRange objectivec.IObject) AVMutableTimedMetadataGroup {
	instance := getAVMutableTimedMetadataGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItems:timeRange:"), objectivec.IObjectSliceToNSArray(items), timeRange)
	return AVMutableTimedMetadataGroupFromID(rv)
}

