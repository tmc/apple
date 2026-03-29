// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptionGroup] class.
var (
	_AVCaptionGroupClass     AVCaptionGroupClass
	_AVCaptionGroupClassOnce sync.Once
)

func getAVCaptionGroupClass() AVCaptionGroupClass {
	_AVCaptionGroupClassOnce.Do(func() {
		_AVCaptionGroupClass = AVCaptionGroupClass{class: objc.GetClass("AVCaptionGroup")}
	})
	return _AVCaptionGroupClass
}

// GetAVCaptionGroupClass returns the class object for AVCaptionGroup.
func GetAVCaptionGroupClass() AVCaptionGroupClass {
	return getAVCaptionGroupClass()
}

type AVCaptionGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptionGroupClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionGroupClass) Alloc() AVCaptionGroup {
	rv := objc.Send[AVCaptionGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents zero or more captions that intersect in time.
//
// # Creating a caption group
//
//   - [AVCaptionGroup.InitWithTimeRange]: Creates a caption group with a time range.
//   - [AVCaptionGroup.InitWithCaptionsTimeRange]: Creates a caption group with captions and a time range.
//
// # Inspecting the caption group
//
//   - [AVCaptionGroup.Captions]: The captions associated with the caption group.
//   - [AVCaptionGroup.TimeRange]: The time range of the caption group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup
type AVCaptionGroup struct {
	objectivec.Object
}

// AVCaptionGroupFromID constructs a [AVCaptionGroup] from an objc.ID.
//
// An object that represents zero or more captions that intersect in time.
func AVCaptionGroupFromID(id objc.ID) AVCaptionGroup {
	return AVCaptionGroup{objectivec.Object{ID: id}}
}
// NOTE: AVCaptionGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptionGroup] class.
//
// # Creating a caption group
//
//   - [IAVCaptionGroup.InitWithTimeRange]: Creates a caption group with a time range.
//   - [IAVCaptionGroup.InitWithCaptionsTimeRange]: Creates a caption group with captions and a time range.
//
// # Inspecting the caption group
//
//   - [IAVCaptionGroup.Captions]: The captions associated with the caption group.
//   - [IAVCaptionGroup.TimeRange]: The time range of the caption group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup
type IAVCaptionGroup interface {
	objectivec.IObject

	// Topic: Creating a caption group

	// Creates a caption group with a time range.
	InitWithTimeRange(timeRange uintptr) AVCaptionGroup
	// Creates a caption group with captions and a time range.
	InitWithCaptionsTimeRange(captions []AVCaption, timeRange uintptr) AVCaptionGroup

	// Topic: Inspecting the caption group

	// The captions associated with the caption group.
	Captions() []AVCaption
	// The time range of the caption group.
	TimeRange() uintptr
}

// Init initializes the instance.
func (c AVCaptionGroup) Init() AVCaptionGroup {
	rv := objc.Send[AVCaptionGroup](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptionGroup) Autorelease() AVCaptionGroup {
	rv := objc.Send[AVCaptionGroup](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptionGroup creates a new AVCaptionGroup instance.
func NewAVCaptionGroup() AVCaptionGroup {
	class := getAVCaptionGroupClass()
	rv := objc.Send[AVCaptionGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a caption group with captions and a time range.
//
// captions: The captions of the caption group.
//
// timeRange: The time range of the caption group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/init(captions:timeRange:)
func NewCaptionGroupWithCaptionsTimeRange(captions []AVCaption, timeRange uintptr) AVCaptionGroup {
	instance := getAVCaptionGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCaptions:timeRange:"), objectivec.IObjectSliceToNSArray(captions), timeRange)
	return AVCaptionGroupFromID(rv)
}

// Creates a caption group with a time range.
//
// timeRange: The time range of the caption group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/init(timeRange:)
func NewCaptionGroupWithTimeRange(timeRange uintptr) AVCaptionGroup {
	instance := getAVCaptionGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTimeRange:"), timeRange)
	return AVCaptionGroupFromID(rv)
}

// Creates a caption group with a time range.
//
// timeRange: The time range of the caption group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/init(timeRange:)
func (c AVCaptionGroup) InitWithTimeRange(timeRange uintptr) AVCaptionGroup {
	rv := objc.Send[AVCaptionGroup](c.ID, objc.Sel("initWithTimeRange:"), timeRange)
	return rv
}
// Creates a caption group with captions and a time range.
//
// captions: The captions of the caption group.
//
// timeRange: The time range of the caption group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/init(captions:timeRange:)
func (c AVCaptionGroup) InitWithCaptionsTimeRange(captions []AVCaption, timeRange uintptr) AVCaptionGroup {
	rv := objc.Send[AVCaptionGroup](c.ID, objc.Sel("initWithCaptions:timeRange:"), objectivec.IObjectSliceToNSArray(captions), timeRange)
	return rv
}

// The captions associated with the caption group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/captions
func (c AVCaptionGroup) Captions() []AVCaption {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("captions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaption {
		return AVCaptionFromID(id)
	})
}
// The time range of the caption group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/timeRange
func (c AVCaptionGroup) TimeRange() uintptr {
	rv := objc.Send[uintptr](c.ID, objc.Sel("timeRange"))
	return rv
}

